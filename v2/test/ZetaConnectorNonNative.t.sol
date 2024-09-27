// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "./utils/ReceiverEVM.sol";

import "./utils/TestERC20.sol";

import "../contracts/evm/ERC20Custody.sol";
import "../contracts/evm/GatewayEVM.sol";
import "../contracts/evm/ZetaConnectorNonNative.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

import { IGatewayEVMErrors, IGatewayEVMEvents } from "../contracts/evm/interfaces/IGatewayEVM.sol";
import "../contracts/evm/interfaces/IZetaConnector.sol";
import "./utils/IReceiverEVM.sol";
import "./utils/Zeta.non-eth.sol";

contract ZetaConnectorNonNativeTest is
    Test,
    IGatewayEVMErrors,
    IGatewayEVMEvents,
    IReceiverEVMEvents,
    IZetaConnectorEvents
{
    using SafeERC20 for IERC20;

    address proxy;
    GatewayEVM gateway;
    ReceiverEVM receiver;
    ERC20Custody custody;
    ZetaConnectorNonNative zetaConnector;
    ZetaNonEth zetaToken;
    address owner;
    address destination;
    address tssAddress;
    RevertContext revertContext;

    error AccessControlUnauthorizedAccount(address account, bytes32 neededRole);
    error ExceedsMaxSupply();
    error EnforcedPause();

    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
    bytes32 public constant TSS_ROLE = keccak256("TSS_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);

        zetaToken = new ZetaNonEth(tssAddress, tssAddress);

        proxy = Upgrades.deployUUPSProxy(
            "GatewayEVM.sol", abi.encodeCall(GatewayEVM.initialize, (tssAddress, address(zetaToken), owner))
        );
        gateway = GatewayEVM(proxy);

        custody = new ERC20Custody(address(gateway), tssAddress, owner);
        zetaConnector = new ZetaConnectorNonNative(address(gateway), address(zetaToken), tssAddress, owner);

        vm.prank(tssAddress);
        zetaToken.updateTssAndConnectorAddresses(tssAddress, address(zetaConnector));

        receiver = new ReceiverEVM();

        vm.deal(tssAddress, 1 ether);

        vm.startPrank(owner);
        gateway.setCustody(address(custody));
        gateway.setConnector(address(zetaConnector));
        vm.stopPrank();

        revertContext = RevertContext({ sender: owner, asset: address(zetaToken), amount: 1, revertMessage: "" });
    }

    function testWithdraw() public {
        uint256 amount = 100_000;
        uint256 balanceBefore = zetaToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        bytes32 internalSendHash = "";

        bytes memory data =
            abi.encodeWithSignature("mint(address,uint256,bytes32)", destination, amount, internalSendHash);
        vm.expectCall(address(zetaToken), 0, data);
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit Withdrawn(destination, amount);
        vm.prank(tssAddress);
        zetaConnector.withdraw(destination, amount, internalSendHash);
        uint256 balanceAfter = zetaToken.balanceOf(destination);
        assertEq(balanceAfter, amount);
    }

    function testWithdrawTogglePause() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";

        vm.prank(tssAddress);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, tssAddress, PAUSER_ROLE));
        zetaConnector.pause();

        vm.prank(tssAddress);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, tssAddress, PAUSER_ROLE));
        zetaConnector.unpause();

        vm.prank(owner);
        zetaConnector.pause();

        vm.expectRevert(EnforcedPause.selector);
        vm.prank(tssAddress);
        zetaConnector.withdraw(destination, amount, internalSendHash);

        vm.prank(owner);
        zetaConnector.unpause();

        uint256 balanceBefore = zetaToken.balanceOf(destination);
        assertEq(balanceBefore, 0);

        bytes memory data =
            abi.encodeWithSignature("mint(address,uint256,bytes32)", destination, amount, internalSendHash);
        vm.expectCall(address(zetaToken), 0, data);
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit Withdrawn(destination, amount);
        vm.prank(tssAddress);
        zetaConnector.withdraw(destination, amount, internalSendHash);
        uint256 balanceAfter = zetaToken.balanceOf(destination);
        assertEq(balanceAfter, amount);
    }

    function testWithdrawFailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        zetaConnector.withdraw(destination, amount, internalSendHash);
    }

    function testWithdrawAndCallReceiveERC20() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(zetaToken), destination);
        uint256 balanceBefore = zetaToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceBeforeZetaConnector, 0);

        bytes memory mintData =
            abi.encodeWithSignature("mint(address,uint256,bytes32)", address(gateway), amount, internalSendHash);
        vm.expectCall(address(zetaToken), 0, mintData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount, address(zetaToken), destination);
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit WithdrawnAndCalled(address(receiver), amount, data);
        vm.prank(tssAddress);
        zetaConnector.withdrawAndCall(address(receiver), amount, data, internalSendHash);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = zetaToken.balanceOf(destination);
        assertEq(balanceAfter, amount);

        // Verify that zeta connector doesn't hold any tokens
        uint256 balanceAfterZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceAfterZetaConnector, 0);

        // Verify that the approval was reset
        uint256 allowance = zetaToken.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = zetaToken.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(zetaToken), destination);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        zetaConnector.withdrawAndCall(address(receiver), amount, data, internalSendHash);
    }

    function testWithdrawAndCallReceiveNoParams() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");
        uint256 balanceBefore = zetaToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceBeforeZetaConnector, 0);

        bytes memory mintData =
            abi.encodeWithSignature("mint(address,uint256,bytes32)", address(gateway), amount, internalSendHash);
        vm.expectCall(address(zetaToken), 0, mintData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit WithdrawnAndCalled(address(receiver), amount, data);
        vm.prank(tssAddress);
        zetaConnector.withdrawAndCall(address(receiver), amount, data, internalSendHash);

        // Verify that the no tokens were transferred to the destination address
        uint256 balanceAfter = zetaToken.balanceOf(destination);
        assertEq(balanceAfter, 0);

        // Verify that zeta connector doesn't hold any tokens
        uint256 balanceAfterZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceAfterZetaConnector, 0);

        // Verify that the approval was reset
        uint256 allowance = zetaToken.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = zetaToken.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndCallReceiveERC20Partial() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodeWithSignature(
            "receiveERC20Partial(uint256,address,address)", amount, address(zetaToken), destination
        );
        uint256 balanceBefore = zetaToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceBeforeZetaConnector, 0);

        bytes memory mintData =
            abi.encodeWithSignature("mint(address,uint256,bytes32)", address(gateway), amount, internalSendHash);
        vm.expectCall(address(zetaToken), 0, mintData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount / 2, address(zetaToken), destination);
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit WithdrawnAndCalled(address(receiver), amount, data);
        vm.prank(tssAddress);
        zetaConnector.withdrawAndCall(address(receiver), amount, data, internalSendHash);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = zetaToken.balanceOf(destination);
        assertEq(balanceAfter, amount / 2);

        // Verify that zeta connector doesn't hold any tokens
        uint256 balanceAfterZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceAfterZetaConnector, 0);

        // Verify that the approval was reset
        uint256 allowance = zetaToken.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = zetaToken.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndRevert() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodePacked("hello");
        uint256 balanceBefore = zetaToken.balanceOf(address(receiver));
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeZetaConnector = zetaToken.balanceOf(address(zetaConnector));

        bytes memory mintData =
            abi.encodeWithSignature("mint(address,uint256,bytes32)", address(gateway), amount, internalSendHash);
        vm.expectCall(address(zetaToken), 0, mintData);
        // Verify that onRevert callback was called
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedRevert(address(gateway), revertContext);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Reverted(address(receiver), address(zetaToken), amount, data, revertContext);
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit WithdrawnAndReverted(address(receiver), amount, data, revertContext);
        vm.prank(tssAddress);
        zetaConnector.withdrawAndRevert(address(receiver), amount, data, internalSendHash, revertContext);

        // Verify that the tokens were transferred to the receiver address
        uint256 balanceAfter = zetaToken.balanceOf(address(receiver));
        assertEq(balanceAfter, amount);

        // Verify that zeta connector doesn't get more tokens
        uint256 balanceAfterZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceAfterZetaConnector, balanceBeforeZetaConnector);

        // Verify that the approval was reset
        uint256 allowance = zetaToken.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = zetaToken.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndRevertFailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodePacked("hello");

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        zetaConnector.withdrawAndRevert(address(receiver), amount, data, internalSendHash, revertContext);
    }

    function testWithdrawAndFailsIfMaxSupplyIsReached() public {
        uint256 amount = 100_000;

        vm.prank(tssAddress);
        zetaConnector.setMaxSupply(amount);

        vm.prank(tssAddress);
        vm.expectRevert(ExceedsMaxSupply.selector);
        zetaConnector.withdraw(address(receiver), amount + 1, "");
    }

    function testWithdrawAndCallFailsIfMaxSupplyIsReached() public {
        uint256 amount = 100_000;
        bytes memory data = abi.encodePacked("hello");

        vm.prank(tssAddress);
        zetaConnector.setMaxSupply(amount);

        vm.prank(tssAddress);
        vm.expectRevert(ExceedsMaxSupply.selector);
        zetaConnector.withdrawAndCall(address(receiver), amount + 1, data, "");
    }

    function testWithdrawAndRevertFailsIfMaxSupplyIsReached() public {
        uint256 amount = 100_000;
        bytes memory data = abi.encodePacked("hello");

        vm.prank(tssAddress);
        zetaConnector.setMaxSupply(amount);

        vm.prank(tssAddress);
        vm.expectRevert(ExceedsMaxSupply.selector);
        zetaConnector.withdrawAndRevert(address(receiver), amount + 1, data, "", revertContext);
    }

    function testSexMaxSupplyFailsIfSenderIsNotTss() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, TSS_ROLE));
        zetaConnector.setMaxSupply(10_000);
    }
}
