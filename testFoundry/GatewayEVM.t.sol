// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "contracts/prototypes/evm/GatewayEVM.sol";
import "contracts/prototypes/evm/ReceiverEVM.sol";
import "contracts/prototypes/evm/ERC20CustodyNew.sol";
import "contracts/prototypes/evm/TestERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "contracts/prototypes/evm/interfaces.sol";
import {Upgrades} from "openzeppelin-foundry-upgrades/LegacyUpgrades.sol";

contract GatewayEVMTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IReceiverEVMEvents {
    using SafeERC20 for IERC20;

    address proxy;
    GatewayEVM gateway;
    ReceiverEVM receiver;
    ERC20CustodyNew custody;
    TestERC20 token;
    address owner;
    address destination;
    address tssAddress;

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);

        token = new TestERC20("test", "TTK");
        proxy = address(new ERC1967Proxy(
            address(new GatewayEVM()),
            abi.encodeWithSelector(GatewayEVM.initialize.selector, (tssAddress))
        ));
        gateway = GatewayEVM(proxy);
        custody = new ERC20CustodyNew(address(gateway));
        receiver = new ReceiverEVM();

        gateway.setCustody(address(custody));

        token.mint(owner, 1000000);
        token.transfer(address(custody), 500000);
    }

    function testForwardCallToReceivePayable() public {
        string memory str = "Hello, Foundry!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        bytes memory data = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);
        
        vm.expectCall(address(receiver), value, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedPayable(address(gateway), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), value, data);
      
        gateway.execute{value: value}(address(receiver), data);
    }

    function testForwardCallToReceiveNonPayable() public {
        string[] memory str = new string[](1);
        str[0] = "Hello, Foundry!";
        uint256[] memory num = new uint256[](1);
        num[0] = 42;
        bool flag = true;

        bytes memory data = abi.encodeWithSignature("receiveNonPayable(string[],uint256[],bool)", str, num, flag);
        
        vm.expectCall(address(receiver), 0, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNonPayable(address(gateway), str, num, flag);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), 0, data);
        
        gateway.execute(address(receiver), data);
    }

    function testForwardCallToReceiveNoParams() public {
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");
        
        vm.expectCall(address(receiver), 0, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), 0, data);
        
        gateway.execute(address(receiver), data);
    }

    function testForwardCallToReceiveERC20ThroughCustody() public {
        uint256 amount = 100000;
        bytes memory data = abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount, address(token), destination);
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
    }

    function testForwardCallToReceiveNoParamsThroughCustody() public {
        uint256 amount = 100000;
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
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
    }
}

contract GatewayEVMInboundTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IReceiverEVMEvents {
    using SafeERC20 for IERC20;

    GatewayEVM gateway;
    ERC20CustodyNew custody;
    TestERC20 token;
    address owner;
    address destination;
    address tssAddress;

    uint256 ownerAmount = 1000000;

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);

        token = new TestERC20("test", "TTK");
         address proxy = address(new ERC1967Proxy(
            address(new GatewayEVM()),
            abi.encodeWithSelector(GatewayEVM.initialize.selector, (tssAddress))
        ));
        gateway = GatewayEVM(proxy);
        custody = new ERC20CustodyNew(address(gateway));

        gateway.setCustody(address(custody));

        token.mint(owner, ownerAmount);
    }

    function testDepositERC20ToCustody() public {
        uint256 amount = 100000;
        uint256 custodyBalanceBefore = token.balanceOf(address(custody));
        assertEq(0, custodyBalanceBefore);

        token.approve(address(gateway), amount);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Deposit(owner, destination, amount, address(token), "");
        gateway.deposit(destination, amount, address(token));

        uint256 custodyBalanceAfter = token.balanceOf(address(custody));
        assertEq(amount, custodyBalanceAfter);

        uint256 ownerAmountAfter = token.balanceOf(owner);
        assertEq(ownerAmount - amount, ownerAmountAfter);
    }

    function testFailDepositERC20ToCustodyIfAmountIs0() public {
        uint256 amount = 0;

        token.approve(address(gateway), amount);

        vm.expectRevert("InsufficientERC20Amount");
        gateway.deposit(destination, amount, address(token));
    }

    function testDepositEthToTss() public {
        uint256 amount = 100000;
        uint256 tssBalanceBefore = tssAddress.balance;

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Deposit(owner, destination, amount, address(0), "");
        gateway.deposit{value: amount}(destination);

        uint256 tssBalanceAfter = tssAddress.balance;
        assertEq(tssBalanceBefore + amount, tssBalanceAfter);
    }

    function testFailDepositEthToTssIfAmountIs0() public {
        uint256 amount = 0;

        vm.expectRevert("InsufficientETHAmount");
        gateway.deposit{value: amount}(destination);
    }

    function testDepositERC20ToCustodyWithPayload() public {
        uint256 amount = 100000;
        uint256 custodyBalanceBefore = token.balanceOf(address(custody));
        assertEq(0, custodyBalanceBefore);

        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        token.approve(address(gateway), amount);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Deposit(owner, destination, amount, address(token), payload);
        gateway.depositAndCall(destination, amount, address(token), payload);

        uint256 custodyBalanceAfter = token.balanceOf(address(custody));
        assertEq(amount, custodyBalanceAfter);

        uint256 ownerAmountAfter = token.balanceOf(owner);
        assertEq(ownerAmount - amount, ownerAmountAfter);
    }

    function testFailDepositERC20ToCustodyWithPayloadIfAmountIs0() public {
        uint256 amount = 0;

        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectRevert("InsufficientERC20Amount");
        gateway.depositAndCall(destination, amount, address(token), payload);
    }

    function testDepositEthToTssWithPayload() public {
        uint256 amount = 100000;
        uint256 tssBalanceBefore = tssAddress.balance;
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Deposit(owner, destination, amount, address(0), payload);
        gateway.depositAndCall{value: amount}(destination, payload);

        uint256 tssBalanceAfter = tssAddress.balance;
        assertEq(tssBalanceBefore + amount, tssBalanceAfter);
    }

    function testFailDepositEthToTssWithPayloadIfAmountIs0() public {
        uint256 amount = 0;
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectRevert("InsufficientETHAmount");
        gateway.depositAndCall{value: amount}(destination, payload);
    }

    function testCallWithPayload() public {
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Call(owner, destination, payload);
        gateway.call(destination, payload);
    }
}