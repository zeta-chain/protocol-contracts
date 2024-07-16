// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "contracts/prototypes/evm/GatewayEVM.sol";
import "contracts/prototypes/evm/ReceiverEVM.sol";
import "contracts/prototypes/evm/ERC20CustodyNew.sol";
import "contracts/prototypes/evm/TestERC20.sol";
import "contracts/prototypes/evm/ReceiverEVM.sol";

import "contracts/prototypes/zevm/GatewayZEVM.sol";
import "contracts/prototypes/zevm/SenderZEVM.sol";
import "contracts/zevm/ZRC20New.sol";
import "contracts/zevm/testing/SystemContractMock.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "contracts/prototypes/evm/interfaces.sol";
import "contracts/prototypes/zevm/interfaces.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Upgrades} from "openzeppelin-foundry-upgrades/LegacyUpgrades.sol";

contract GatewayEVMZEVMTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IGatewayZEVMEvents, IGatewayZEVMErrors, IReceiverEVMEvents {
    // evm
    using SafeERC20 for IERC20;

    address proxyEVM;
    GatewayEVM gatewayEVM;
    ERC20CustodyNew custody;
    TestERC20 token;
    ReceiverEVM receiverEVM;
    address ownerEVM;
    address destination;
    address tssAddress;

    // zevm
    address proxyZEVM;
    GatewayZEVM gatewayZEVM;
    SenderZEVM senderZEVM;
    SystemContractMock systemContract;
    ZRC20New zrc20;
    address ownerZEVM;
    
    function setUp() public {
        // evm
        ownerEVM = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);
        ownerZEVM = address(0x4321);

        token = new TestERC20("test", "TTK");
        proxyEVM = address(new ERC1967Proxy(
            address(new GatewayEVM()),
            abi.encodeWithSelector(GatewayEVM.initialize.selector, (tssAddress))
        ));
        gatewayEVM = GatewayEVM(proxyEVM);
        custody = new ERC20CustodyNew(address(gatewayEVM));

        gatewayEVM.setCustody(address(custody));

        token.mint(ownerEVM, 1000000);
        token.transfer(address(custody), 500000);

        receiverEVM = new ReceiverEVM();

        // zevm
        proxyZEVM = address(new ERC1967Proxy(
            address(new GatewayZEVM()),
            abi.encodeWithSelector(GatewayZEVM.initialize.selector, "")
        ));
        gatewayZEVM = GatewayZEVM(proxyZEVM);
        senderZEVM = new SenderZEVM(address(gatewayZEVM));
        address fungibleModuleAddress = address(0x735b14BB79463307AAcBED86DAf3322B1e6226aB);
        vm.startPrank(fungibleModuleAddress);
        systemContract = new SystemContractMock(address(0), address(0), address(0));
        zrc20 = new ZRC20New("TOKEN", "TKN", 18, 1, CoinType.Zeta, 0, address(systemContract), address(gatewayZEVM));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        zrc20.deposit(ownerZEVM, 1000000);
        zrc20.deposit(address(senderZEVM), 1000000);
        vm.stopPrank();

        vm.prank(ownerZEVM);
        zrc20.approve(address(gatewayZEVM), 1000000);
    }

    function testCallReceiverEVMFromZEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on zevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        vm.prank(ownerZEVM);
        vm.expectEmit(true, true, true, true, address(gatewayZEVM));
        emit Call(address(ownerZEVM), abi.encodePacked(receiverEVM), message);
        gatewayZEVM.call(abi.encodePacked(receiverEVM), message);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        gatewayEVM.execute{value: value}(address(receiverEVM), message);
    }

    function testCallReceiverEVMFromSenderZEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on zevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        bytes memory data = abi.encodeWithSignature("call(bytes,bytes)", abi.encodePacked(receiverEVM), message);
        vm.expectCall(address(gatewayZEVM), 0, data);
        vm.prank(ownerZEVM);
        senderZEVM.callReceiver(abi.encodePacked(receiverEVM), str, num, flag);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(receiverEVM));
        emit ReceivedPayable(address(gatewayEVM), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        gatewayEVM.execute{value: value}(address(receiverEVM), message);
    }

    function testWithdrawAndCallReceiverEVMFromZEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on zevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayZEVM));
        emit Withdrawal(
            ownerZEVM,
            abi.encodePacked(receiverEVM),
            1000000,
            0,
            zrc20.PROTOCOL_FLAT_FEE(),
            message
        );
        vm.prank(ownerZEVM);
        gatewayZEVM.withdrawAndCall(
            abi.encodePacked(receiverEVM),
            1000000,
            address(zrc20),
            message
        );

        // Check the balance after withdrawal
        uint256 balanceOfAfterWithdrawal = zrc20.balanceOf(ownerZEVM);
        assertEq(balanceOfAfterWithdrawal, 0);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(receiverEVM));
        emit ReceivedPayable(address(gatewayEVM), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        gatewayEVM.execute{value: value}(address(receiverEVM), message);
    }

    function testWithdrawAndCallReceiverEVMFromSenderZEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on zevm
        uint256 senderBalanceBeforeWithdrawal = IZRC20(zrc20).balanceOf(address(senderZEVM));
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        bytes memory data = abi.encodeWithSignature("withdrawAndCall(bytes,uint256,address,bytes)", abi.encodePacked(receiverEVM), 1000000,  address(zrc20), message);
        vm.expectCall(address(gatewayZEVM), 0, data);
        vm.prank(ownerZEVM);
        senderZEVM.withdrawAndCallReceiver(abi.encodePacked(receiverEVM), 1000000, address(zrc20), str, num, flag);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(receiverEVM));
        emit ReceivedPayable(address(gatewayEVM), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        gatewayEVM.execute{value: value}(address(receiverEVM), message);

        // Check the balance after withdrawal
        uint256 senderBalanceAfterWithdrawal = IZRC20(zrc20).balanceOf(address(senderZEVM));
        assertEq(senderBalanceAfterWithdrawal, senderBalanceBeforeWithdrawal - 1000000);
    }
}