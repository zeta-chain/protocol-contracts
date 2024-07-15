// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "contracts/prototypes/zevm/GatewayZEVM.sol";
import "contracts/zevm/ZRC20New.sol";
import "contracts/zevm/SystemContract.sol";
import "contracts/zevm/interfaces/IZRC20.sol";
import "contracts/prototypes/zevm/TestZContract.sol";
import "../zevm/interfaces.sol";

contract GatewayZEVMInboundTest is Test, IGatewayZEVMEvents, IGatewayZEVMErrors {
    GatewayZEVM gateway;
    ZRC20New zrc20;
    SystemContract systemContract;
    TestZContract testZContract;
    address owner;
    address addr1;

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1234);

        gateway = new GatewayZEVM();
        testZContract = new TestZContract();

        vm.startPrank(gateway.FUNGIBLE_MODULE_ADDRESS());
        systemContract = new SystemContract(address(0), address(0), address(0));
        zrc20 = new ZRC20New("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);

        vm.deal(gateway.FUNGIBLE_MODULE_ADDRESS(), 1000000000);
        zrc20.deposit(owner, 100000);
        vm.stopPrank();

        vm.prank(owner);
        zrc20.approve(address(gateway), 100000);
    }

    function testWithdrawZRC20() public {
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawal(owner, abi.encodePacked(addr1), 1, 0, zrc20.PROTOCOL_FLAT_FEE(), "");
        gateway.withdraw(abi.encodePacked(addr1), 1, address(zrc20));

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - 1, ownerBalanceAfter);
    }

    function testWithdrawZRC20WithMessage() public {
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawal(owner, abi.encodePacked(addr1), 1, 0, zrc20.PROTOCOL_FLAT_FEE(), message);
        gateway.withdrawAndCall(abi.encodePacked(addr1), 1, address(zrc20), message);

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - 1, ownerBalanceAfter);
    }

    function testCall() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Call(owner, abi.encodePacked(addr1), message);
        gateway.call(abi.encodePacked(addr1), message);
    }
}

contract GatewayZEVMOutboundTest is Test, IGatewayZEVMEvents, IGatewayZEVMErrors {
    GatewayZEVM gateway;
    ZRC20New zrc20;
    SystemContract systemContract;
    TestZContract testZContract;
    address owner;
    address addr1;
    event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message);

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1234);

        gateway = new GatewayZEVM();
        testZContract = new TestZContract();

        vm.startPrank(gateway.FUNGIBLE_MODULE_ADDRESS());
        systemContract = new SystemContract(address(0), address(0), address(0));
        zrc20 = new ZRC20New("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);

        vm.deal(gateway.FUNGIBLE_MODULE_ADDRESS(), 1000000000);
        zrc20.deposit(owner, 100000);
        vm.stopPrank();

        vm.prank(owner);
        zrc20.approve(address(gateway), 100000);
    }

    function testDeposit() public {
        uint256 balanceBefore = zrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(gateway.FUNGIBLE_MODULE_ADDRESS());
        gateway.deposit(address(zrc20), 1, addr1);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(1, balanceAfter);
    }

    function testExecuteZContract() public {
        bytes memory message = abi.encode("hello");
        zContext memory context = zContext({
            origin: abi.encodePacked(address(gateway)),
            sender: gateway.FUNGIBLE_MODULE_ADDRESS(),
            chainID: 1
        });

        vm.expectEmit(true, true, true, true, address(testZContract));
        emit ContextData(abi.encodePacked(gateway), gateway.FUNGIBLE_MODULE_ADDRESS(), 1, address(gateway), "hello");
        vm.prank(gateway.FUNGIBLE_MODULE_ADDRESS());
        gateway.execute(context, address(zrc20), 1, address(testZContract), message);
    }
   
    function testDepositAndCallZContract() public {
        uint256 balanceBefore = zrc20.balanceOf(address(testZContract));
        assertEq(0, balanceBefore);

        bytes memory message = abi.encode("hello");
        zContext memory context = zContext({
            origin: abi.encodePacked(address(gateway)),
            sender: gateway.FUNGIBLE_MODULE_ADDRESS(),
            chainID: 1
        });
    
        vm.expectEmit(true, true, true, true, address(testZContract));
        emit ContextData(abi.encodePacked(gateway), gateway.FUNGIBLE_MODULE_ADDRESS(), 1, address(gateway), "hello");
        vm.prank(gateway.FUNGIBLE_MODULE_ADDRESS());
        gateway.depositAndCall(context, address(zrc20), 1, address(testZContract), message);

        uint256 balanceAfter = zrc20.balanceOf(address(testZContract));
        assertEq(1, balanceAfter);
    }
}