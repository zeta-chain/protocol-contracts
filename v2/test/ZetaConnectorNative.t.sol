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
import "src/evm/ZetaConnectorNative.sol";
import "src/evm/interfaces/IGatewayEVM.sol";
import "src/evm/interfaces/IZetaConnector.sol";

contract ZetaConnectorNativeTest is
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
    ZetaConnectorNative zetaConnector;
    TestERC20 zetaToken;
    address owner;
    address destination;
    address tssAddress;

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);

        zetaToken = new TestERC20("zeta", "ZETA");

        proxy = Upgrades.deployUUPSProxy(
            "GatewayEVM.sol", abi.encodeCall(GatewayEVM.initialize, (tssAddress, address(zetaToken)))
        );
        gateway = GatewayEVM(proxy);
        custody = new ERC20Custody(address(gateway), tssAddress);
        zetaConnector = new ZetaConnectorNative(address(gateway), address(zetaToken), tssAddress);

        receiver = new ReceiverEVM();

        vm.deal(tssAddress, 1 ether);

        vm.startPrank(tssAddress);
        gateway.setCustody(address(custody));
        gateway.setConnector(address(zetaConnector));
        vm.stopPrank();

        zetaToken.mint(address(zetaConnector), 5_000_000);

        vm.deal(tssAddress, 1 ether);
    }

    function testWithdraw() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        uint256 balanceBefore = zetaToken.balanceOf(destination);
        assertEq(balanceBefore, 0);

        bytes memory data = abi.encodeWithSignature("transfer(address,uint256)", destination, amount);
        vm.expectCall(address(zetaToken), 0, data);
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit Withdraw(destination, amount);
        vm.prank(tssAddress);
        zetaConnector.withdraw(destination, amount, internalSendHash);
        uint256 balanceAfter = zetaToken.balanceOf(destination);
        assertEq(balanceAfter, amount);
    }

    function testWithdrawFailsIfSenderIsNotTSS() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";

        vm.prank(owner);
        vm.expectRevert(InvalidSender.selector);
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

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(zetaToken), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount, address(zetaToken), destination);
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit WithdrawAndCall(address(receiver), amount, data);
        vm.prank(tssAddress);
        zetaConnector.withdrawAndCall(address(receiver), amount, data, internalSendHash);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = zetaToken.balanceOf(destination);
        assertEq(balanceAfter, amount);

        // Verify that zeta connector doesn't get more tokens
        uint256 balanceAfterZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceAfterZetaConnector, balanceBeforeZetaConnector - amount);

        // Verify that the approval was reset
        uint256 allowance = zetaToken.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = zetaToken.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndCallReceiveERC20FailsIfSenderIsNotTSS() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(zetaToken), destination);

        vm.prank(owner);
        vm.expectRevert(InvalidSender.selector);
        zetaConnector.withdrawAndCall(address(receiver), amount, data, internalSendHash);
    }

    function testWithdrawAndCallReceiveNoParams() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");
        uint256 balanceBefore = zetaToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeZetaConnector = zetaToken.balanceOf(address(zetaConnector));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(zetaToken), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit WithdrawAndCall(address(receiver), amount, data);
        vm.prank(tssAddress);
        zetaConnector.withdrawAndCall(address(receiver), amount, data, internalSendHash);

        // Verify that the no tokens were transferred to the destination address
        uint256 balanceAfter = zetaToken.balanceOf(destination);
        assertEq(balanceAfter, 0);

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

    function testWithdrawAndCallReceiveERC20Partial() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodeWithSignature(
            "receiveERC20Partial(uint256,address,address)", amount, address(zetaToken), destination
        );
        uint256 balanceBefore = zetaToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeZetaConnector = zetaToken.balanceOf(address(zetaConnector));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(zetaToken), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount / 2, address(zetaToken), destination);
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit WithdrawAndCall(address(receiver), amount, data);
        vm.prank(tssAddress);
        zetaConnector.withdrawAndCall(address(receiver), amount, data, internalSendHash);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = zetaToken.balanceOf(destination);
        assertEq(balanceAfter, amount / 2);

        // Verify that zeta connector doesn't get more tokens
        uint256 balanceAfterZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceAfterZetaConnector, balanceBeforeZetaConnector - amount / 2);

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

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(zetaToken), 0, transferData);
        // Verify that onRevert callback was called
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedRevert(address(gateway), data);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit RevertedWithERC20(address(zetaToken), address(receiver), amount, data);
        vm.expectEmit(true, true, true, true, address(zetaConnector));
        emit WithdrawAndRevert(address(receiver), amount, data);
        vm.prank(tssAddress);
        zetaConnector.withdrawAndRevert(address(receiver), amount, data, internalSendHash);

        // Verify that the tokens were transferred to the receiver address
        uint256 balanceAfter = zetaToken.balanceOf(address(receiver));
        assertEq(balanceAfter, amount);

        // Verify that zeta connector doesn't get more tokens
        uint256 balanceAfterZetaConnector = zetaToken.balanceOf(address(zetaConnector));
        assertEq(balanceAfterZetaConnector, balanceBeforeZetaConnector - amount);

        // Verify that the approval was reset
        uint256 allowance = zetaToken.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = zetaToken.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndRevertFailsIfSenderIsNotTSS() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodePacked("hello");

        vm.prank(owner);
        vm.expectRevert(InvalidSender.selector);
        zetaConnector.withdrawAndRevert(address(receiver), amount, data, internalSendHash);
    }
}
