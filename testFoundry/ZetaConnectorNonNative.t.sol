// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "contracts/prototypes/evm/GatewayEVM.sol";
import "contracts/prototypes/evm/ReceiverEVM.sol";
import "contracts/prototypes/evm/ERC20CustodyNew.sol";
import "contracts/prototypes/evm/ZetaConnectorNonNative.sol";
import "contracts/prototypes/evm/TestERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "contracts/prototypes/evm/IGatewayEVM.sol";
import "contracts/prototypes/evm/IReceiverEVM.sol";
import "contracts/evm/Zeta.non-eth.sol";
import {Upgrades} from "openzeppelin-foundry-upgrades/LegacyUpgrades.sol";

contract ZetaConnectorNonNativeTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IReceiverEVMEvents {
    using SafeERC20 for IERC20;

    address proxy;
    GatewayEVM gateway;
    ReceiverEVM receiver;
    ERC20CustodyNew custody;
    ZetaConnectorNonNative zetaConnector;
    TestERC20 token;
    ZetaNonEth zetaToken;
    address owner;
    address destination;
    address tssAddress;

    event Withdraw(address indexed to, uint256 amount);
    event WithdrawAndCall(address indexed to, uint256 amount, bytes data);

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);

        token = new TestERC20("test", "TTK");
        zetaToken = new ZetaNonEth(tssAddress, tssAddress);

        proxy = address(new ERC1967Proxy(
            address(new GatewayEVM()),
            abi.encodeWithSelector(GatewayEVM.initialize.selector, tssAddress, address(zetaToken))
        ));
        gateway = GatewayEVM(proxy);
        custody = new ERC20CustodyNew(address(gateway));
        zetaConnector = new ZetaConnectorNonNative(address(gateway), address(zetaToken));

        vm.prank(tssAddress);
        zetaToken.updateTssAndConnectorAddresses(tssAddress, address(zetaConnector));

        receiver = new ReceiverEVM();

        gateway.setCustody(address(custody));
        gateway.setConnector(address(zetaConnector));

        token.mint(owner, 1000000);
        token.transfer(address(custody), 500000);
    }

    function testWithdraw() public {
        uint256 amount = 100000;
        uint256 balanceBefore = zetaToken.balanceOf(destination);
        bytes32 internalSendHash = "";
        assertEq(balanceBefore, 0);

        bytes memory data = abi.encodeWithSignature("mint(address,uint256,bytes32)", destination, amount, internalSendHash);
        vm.expectCall(address(zetaToken), 0, data);
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit Withdraw(destination, amount);
        zetaConnector.withdraw(destination, amount, internalSendHash);
        uint256 balanceAfter = zetaToken.balanceOf(destination);
        assertEq(balanceAfter, amount);
    }

    function testWithdrawAndCallReceiveERC20() public {
        uint256 amount = 100000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(zetaToken), destination);
        uint256 balanceBefore = zetaToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceBeforeZetaConnector, 0);

        bytes memory mintData = abi.encodeWithSignature("mint(address,uint256,bytes32)", address(gateway), amount, internalSendHash);
        vm.expectCall(address(zetaToken), 0, mintData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount, address(zetaToken), destination);
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit WithdrawAndCall(address(receiver), amount, data);
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

    function testWithdrawAndCallReceiveNoParams() public {
        uint256 amount = 100000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");
        uint256 balanceBefore = zetaToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceBeforeZetaConnector, 0);

        bytes memory mintData = abi.encodeWithSignature("mint(address,uint256,bytes32)", address(gateway), amount, internalSendHash);
        vm.expectCall(address(zetaToken), 0, mintData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit WithdrawAndCall(address(receiver), amount, data);
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
        uint256 amount = 100000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodeWithSignature("receiveERC20Partial(uint256,address,address)", amount, address(zetaToken), destination);
        uint256 balanceBefore = zetaToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceBeforeZetaConnector, 0);

        bytes memory mintData = abi.encodeWithSignature("mint(address,uint256,bytes32)", address(gateway), amount, internalSendHash);
        vm.expectCall(address(zetaToken), 0, mintData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount / 2, address(zetaToken), destination);
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit WithdrawAndCall(address(receiver), amount, data);
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
}