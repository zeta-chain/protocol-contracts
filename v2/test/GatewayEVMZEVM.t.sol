// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "./utils/ReceiverEVM.sol";

import "./utils/TestERC20.sol";
import "src/evm/ERC20Custody.sol";
import "src/evm/GatewayEVM.sol";
import "src/evm/ZetaConnectorNonNative.sol";

import "./utils/SenderZEVM.sol";

import "./utils/SystemContractMock.sol";

import "src/zevm/GatewayZEVM.sol";
import "src/zevm/ZRC20.sol";

import "./utils/IReceiverEVM.sol";
import "src/evm/interfaces/IGatewayEVM.sol";
import "src/zevm/interfaces/IGatewayZEVM.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract GatewayEVMZEVMTest is
    Test,
    IGatewayEVMErrors,
    IGatewayEVMEvents,
    IGatewayZEVMEvents,
    IGatewayZEVMErrors,
    IReceiverEVMEvents
{
    // evm
    using SafeERC20 for IERC20;

    address proxyEVM;
    GatewayEVM gatewayEVM;
    ERC20Custody custody;
    ZetaConnectorNonNative zetaConnector;
    TestERC20 token;
    TestERC20 zeta;
    ReceiverEVM receiverEVM;
    address ownerEVM;
    address destination;
    address tssAddress;

    // zevm
    address payable proxyZEVM;
    GatewayZEVM gatewayZEVM;
    SenderZEVM senderZEVM;
    SystemContractMock systemContract;
    ZRC20 zrc20;
    address ownerZEVM;

    function setUp() public {
        // evm
        ownerEVM = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);
        ownerZEVM = address(0x4321);

        token = new TestERC20("test", "TTK");
        zeta = new TestERC20("zeta", "ZETA");

        proxyEVM = Upgrades.deployUUPSProxy(
            "GatewayEVM.sol", abi.encodeCall(GatewayEVM.initialize, (tssAddress, address(zeta), ownerEVM))
        );
        gatewayEVM = GatewayEVM(proxyEVM);
        custody = new ERC20Custody(address(gatewayEVM), tssAddress, ownerEVM);
        zetaConnector = new ZetaConnectorNonNative(address(gatewayEVM), address(zeta), tssAddress, ownerEVM);

        vm.deal(tssAddress, 1 ether);

        vm.startPrank(ownerEVM);
        gatewayEVM.setCustody(address(custody));
        gatewayEVM.setConnector(address(zetaConnector));
        vm.stopPrank();

        token.mint(ownerEVM, 1_000_000);
        token.transfer(address(custody), 500_000);

        receiverEVM = new ReceiverEVM();

        // zevm
        proxyZEVM = payable(
            Upgrades.deployUUPSProxy(
                "GatewayZEVM.sol", abi.encodeCall(GatewayZEVM.initialize, (address(zeta), ownerZEVM))
            )
        );
        gatewayZEVM = GatewayZEVM(proxyZEVM);

        senderZEVM = new SenderZEVM(address(gatewayZEVM));
        address fungibleModuleAddress = address(0x735b14BB79463307AAcBED86DAf3322B1e6226aB);
        vm.startPrank(fungibleModuleAddress);
        systemContract = new SystemContractMock(address(0), address(0), address(0));
        zrc20 = new ZRC20("TOKEN", "TKN", 18, 1, CoinType.Zeta, 0, address(systemContract), address(gatewayZEVM));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        zrc20.deposit(ownerZEVM, 1_000_000);
        zrc20.deposit(address(senderZEVM), 1_000_000);
        vm.stopPrank();

        vm.prank(ownerZEVM);
        zrc20.approve(address(gatewayZEVM), 1_000_000);

        vm.deal(tssAddress, 1 ether);
    }

    function testCallReceiverEVMFromZEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;
        uint256 chainId = 1;

        // Encode the function call data and call on zevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        vm.prank(ownerZEVM);
        vm.expectEmit(true, true, true, true, address(gatewayZEVM));
        emit Called(address(ownerZEVM), chainId, abi.encodePacked(receiverEVM), message);
        gatewayZEVM.call(abi.encodePacked(receiverEVM), chainId, message);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        vm.prank(tssAddress);
        gatewayEVM.execute{ value: value }(address(receiverEVM), message);
    }

    function testCallReceiverEVMFromSenderZEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;
        uint256 chainId = 1;

        // Encode the function call data and call on zevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        bytes memory data =
            abi.encodeWithSignature("call(bytes,uint256,bytes)", abi.encodePacked(receiverEVM), chainId, message);
        vm.expectCall(address(gatewayZEVM), 0, data);
        vm.prank(ownerZEVM);
        senderZEVM.callReceiver(abi.encodePacked(receiverEVM), chainId, str, num, flag);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(receiverEVM));
        emit ReceivedPayable(address(gatewayEVM), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        vm.prank(tssAddress);
        gatewayEVM.execute{ value: value }(address(receiverEVM), message);
    }

    function testWithdrawAndCallReceiverEVMFromZEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on zevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayZEVM));
        emit Withdrawn(
            ownerZEVM,
            0,
            abi.encodePacked(receiverEVM),
            address(zrc20),
            1_000_000,
            0,
            zrc20.PROTOCOL_FLAT_FEE(),
            message
        );
        vm.prank(ownerZEVM);
        gatewayZEVM.withdrawAndCall(abi.encodePacked(receiverEVM), 1_000_000, address(zrc20), message);

        // Check the balance after withdrawal
        uint256 balanceOfAfterWithdrawal = zrc20.balanceOf(ownerZEVM);
        assertEq(balanceOfAfterWithdrawal, 0);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(receiverEVM));
        emit ReceivedPayable(address(gatewayEVM), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        vm.prank(tssAddress);
        gatewayEVM.execute{ value: value }(address(receiverEVM), message);
    }

    function testWithdrawAndCallReceiverEVMFromSenderZEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on zevm
        uint256 senderBalanceBeforeWithdrawal = IZRC20(zrc20).balanceOf(address(senderZEVM));
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        bytes memory data = abi.encodeWithSignature(
            "withdrawAndCall(bytes,uint256,address,bytes)",
            abi.encodePacked(receiverEVM),
            1_000_000,
            address(zrc20),
            message
        );
        vm.expectCall(address(gatewayZEVM), 0, data);
        vm.prank(ownerZEVM);
        senderZEVM.withdrawAndCallReceiver(abi.encodePacked(receiverEVM), 1_000_000, address(zrc20), str, num, flag);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(receiverEVM));
        emit ReceivedPayable(address(gatewayEVM), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        vm.prank(tssAddress);
        gatewayEVM.execute{ value: value }(address(receiverEVM), message);

        // Check the balance after withdrawal
        uint256 senderBalanceAfterWithdrawal = IZRC20(zrc20).balanceOf(address(senderZEVM));
        assertEq(senderBalanceAfterWithdrawal, senderBalanceBeforeWithdrawal - 1_000_000);
    }
}
