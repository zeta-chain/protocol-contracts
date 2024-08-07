// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "./utils/ReceiverEVM.sol";

import "./utils/TestERC20.sol";

import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

import "./utils/IReceiverEVM.sol";
import "src/evm/ERC20Custody.sol";
import "src/evm/GatewayEVM.sol";
import "src/evm/ZetaConnectorNonNative.sol";
import "src/evm/interfaces/IERC20Custody.sol";
import "src/evm/interfaces/IGatewayEVM.sol";

contract ERC20CustodyTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IReceiverEVMEvents, IERC20CustodyEvents {
    using SafeERC20 for IERC20;

    address proxy;
    GatewayEVM gateway;
    ReceiverEVM receiver;
    ERC20Custody custody;
    ZetaConnectorNonNative zetaConnector;
    TestERC20 token;
    TestERC20 zeta;
    address owner;
    address destination;
    address tssAddress;

    error EnforcedPause();
    error AccessControlUnauthorizedAccount(address account, bytes32 neededRole);

    bytes32 public constant TSS_ROLE = keccak256("TSS_ROLE");
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
    bytes32 public constant ASSET_HANDLER_ROLE = keccak256("ASSET_HANDLER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);

        token = new TestERC20("test", "TTK");
        zeta = new TestERC20("zeta", "ZETA");

        proxy = Upgrades.deployUUPSProxy(
            "GatewayEVM.sol", abi.encodeCall(GatewayEVM.initialize, (tssAddress, address(zeta), owner))
        );
        gateway = GatewayEVM(proxy);
        custody = new ERC20Custody(address(gateway), tssAddress, owner);
        zetaConnector = new ZetaConnectorNonNative(address(gateway), address(zeta), tssAddress, owner);
        receiver = new ReceiverEVM();

        vm.deal(tssAddress, 1 ether);

        vm.startPrank(owner);
        gateway.setCustody(address(custody));
        gateway.setConnector(address(zetaConnector));
        vm.stopPrank();

        token.mint(owner, 1_000_000);
        token.transfer(address(custody), 500_000);

        vm.deal(tssAddress, 1 ether);
    }

    function testNewCustodyFailsIfAddressesAreZero() public {
        vm.expectRevert(ZeroAddress.selector);
        ERC20Custody newCustody = new ERC20Custody(address(0), tssAddress, owner);

        vm.expectRevert(ZeroAddress.selector);
        newCustody = new ERC20Custody(address(gateway), address(0), owner);

        vm.expectRevert(ZeroAddress.selector);
        newCustody = new ERC20Custody(address(gateway), tssAddress, address(0));

        newCustody = new ERC20Custody(address(gateway), tssAddress, owner);
    }

    function testForwardCallToReceiveERC20ThroughCustody() public {
        uint256 amount = 100_000;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount, address(token), destination);
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawAndCall(address(token), address(receiver), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, amount);

        // Verify that the remaining tokens were refunded to the Custody contract
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testForwardCallToReceiveERC20ThroughCustodyTogglePause() public {
        vm.prank(tssAddress);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, tssAddress, PAUSER_ROLE));
        custody.pause();

        vm.prank(tssAddress);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, tssAddress, PAUSER_ROLE));
        gateway.unpause();

        vm.prank(owner);
        custody.pause();

        uint256 amount = 100_000;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);

        vm.expectRevert(EnforcedPause.selector);
        vm.prank(tssAddress);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);

        vm.prank(owner);
        custody.unpause();

        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount, address(token), destination);
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawAndCall(address(token), address(receiver), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, amount);

        // Verify that the remaining tokens were refunded to the Custody contract
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        custody.withdrawAndCall(address(token), address(receiver), amount, data);
    }

    function testForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0() public {
        uint256 amount = 0;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);

        vm.prank(tssAddress);
        vm.expectRevert(InsufficientERC20Amount.selector);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);
    }

    function testForwardCallToReceiveERC20PartialThroughCustody() public {
        uint256 amount = 100_000;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20Partial(uint256,address,address)", amount, address(token), destination);
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount / 2, address(token), destination);
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawAndCall(address(token), address(receiver), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, amount / 2);

        // Verify that the remaining tokens were refunded to the Custody contract
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount / 2);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20Partial(uint256,address,address)", amount, address(token), destination);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        custody.withdrawAndCall(address(token), address(receiver), amount, data);
    }

    function testForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0() public {
        uint256 amount = 0;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20Partial(uint256,address,address)", amount, address(token), destination);

        vm.prank(tssAddress);
        vm.expectRevert(InsufficientERC20Amount.selector);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);
    }

    function testForwardCallToReceiveNoParamsThroughCustody() public {
        uint256 amount = 100_000;
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawAndCall(address(token), address(receiver), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);

        // Verify that the tokens were not transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, 0);

        // Verify that the remaining tokens were refunded to the Custody contract
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawThroughCustody() public {
        uint256 amount = 100_000;
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(destination), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(custody));
        emit Withdraw(address(token), destination, amount);
        vm.prank(tssAddress);
        custody.withdraw(address(token), destination, amount);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, amount);

        // Verify that the tokens were substracted from custody
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        custody.withdraw(address(token), destination, amount);
    }

    function testWithdrawAndRevertThroughCustody() public {
        uint256 amount = 100_000;
        bytes memory data = abi.encodePacked("hello");
        uint256 balanceBefore = token.balanceOf(address(receiver));
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        // Verify that onRevert callback was called
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedRevert(address(gateway), data);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit RevertedWithERC20(address(token), address(receiver), amount, data);
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawAndRevert(address(token), address(receiver), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndRevert(address(token), address(receiver), amount, data);

        // Verify that the tokens were transferred to the receiver address
        uint256 balanceAfter = token.balanceOf(address(receiver));
        assertEq(balanceAfter, amount);

        // Verify that zeta connector doesn't get more tokens
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;
        bytes memory data = abi.encodePacked("hello");

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        custody.withdrawAndRevert(address(token), address(receiver), amount, data);
    }

    function testWithdrawAndRevertThroughCustodyFailsIfAmountIs0() public {
        uint256 amount = 0;
        bytes memory data = abi.encodePacked("hello");

        vm.prank(tssAddress);
        vm.expectRevert(InsufficientERC20Amount.selector);
        custody.withdrawAndRevert(address(token), address(receiver), amount, data);
    }
}