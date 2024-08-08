// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "./utils/SystemContract.sol";

import "./utils/TestZContract.sol";

import "./utils/WZETA.sol";

import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";
import "src/zevm/GatewayZEVM.sol";
import "src/zevm/ZRC20.sol";
import "src/zevm/interfaces/IGatewayZEVM.sol";
import "src/zevm/interfaces/IZRC20.sol";

contract GatewayZEVMInboundTest is Test, IGatewayZEVMEvents, IGatewayZEVMErrors {
    address payable proxy;
    GatewayZEVM gateway;
    ZRC20 zrc20;
    WETH9 zetaToken;
    SystemContract systemContract;
    TestZContract testZContract;
    address owner;
    address addr1;
    address fungibleModule;
    RevertOptions revertOptions;

    error ZeroAddress();
    error LowBalance();

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1234);

        zetaToken = new WETH9();

        proxy = payable(
            Upgrades.deployUUPSProxy(
                "GatewayZEVM.sol", abi.encodeCall(GatewayZEVM.initialize, (address(zetaToken), owner))
            )
        );
        gateway = GatewayZEVM(proxy);

        fungibleModule = gateway.FUNGIBLE_MODULE_ADDRESS();
        testZContract = new TestZContract();

        vm.startPrank(fungibleModule);
        systemContract = new SystemContract(address(0), address(0), address(0));
        zrc20 = new ZRC20("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        vm.deal(fungibleModule, 1_000_000_000);
        zetaToken.deposit{ value: 10 }();
        zetaToken.approve(address(gateway), 10);
        zrc20.deposit(owner, 100_000);
        vm.stopPrank();

        vm.startPrank(owner);
        zrc20.approve(address(gateway), 100_000);
        zetaToken.deposit{ value: 10 }();
        zetaToken.approve(address(gateway), 10);
        vm.stopPrank();

        revertOptions =
            RevertOptions({ revertAddress: address(0x321), callOnRevert: true, abortAddress: address(0x321) });
    }

    function testWithdrawZRC20() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawal(
            owner, 0, abi.encodePacked(addr1), address(zrc20), amount, 0, zrc20.PROTOCOL_FLAT_FEE(), "", revertOptions
        );
        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20), revertOptions);

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount, ownerBalanceAfter);
    }

    function testWithdrawZRC20FailsIfNoBalanceForGasFee() public {
        uint256 amount = 1;
        uint256 ownerBalance = zrc20.balanceOf(owner);
        zrc20.transfer(address(0x123), ownerBalance);

        vm.prank(fungibleModule);
        zrc20.updateGasLimit(10);

        vm.expectRevert(LowBalance.selector);
        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20), revertOptions);
    }

    function testWithdrawZRC20FailsIfNoBalanceForTransfer() public {
        uint256 amount = 2;
        uint256 ownerBalance = zrc20.balanceOf(owner);
        zrc20.transfer(address(0x123), ownerBalance - 1);

        vm.expectRevert(LowBalance.selector);
        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20), revertOptions);
    }

    function testWithdrawZRC20FailsIsAmountIs0() public {
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.withdraw(abi.encodePacked(addr1), 0, address(zrc20), revertOptions);
    }

    function testWithdrawZRC20FailsIfReceiverIsZeroAddress() public {
        vm.expectRevert(ZeroAddress.selector);
        gateway.withdraw(abi.encodePacked(""), 1, address(zrc20), revertOptions);
    }

    function testWithdrawZRC20FailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        // Remove allowance for gateway
        vm.prank(owner);
        zrc20.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20), revertOptions);

        // Check that balance didn't change
        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);
    }

    function testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(ZeroAddress.selector);
        gateway.withdrawAndCall(abi.encodePacked(""), 1, address(zrc20), message, revertOptions);
    }

    function testWithdrawAndCallZRC20FailsIfAmountIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.withdrawAndCall(abi.encodePacked(addr1), 0, address(zrc20), message, revertOptions);
    }

    function testWithdrawZRC20WithMessageFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        // Remove allowance for gateway
        vm.prank(owner);
        zrc20.approve(address(gateway), 0);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert();
        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, address(zrc20), message, revertOptions);

        // Check that balance didn't change
        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);
    }

    function testWithdrawZRC20WithMessage() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawal(
            owner,
            0,
            abi.encodePacked(addr1),
            address(zrc20),
            amount,
            0,
            zrc20.PROTOCOL_FLAT_FEE(),
            message,
            revertOptions
        );
        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, address(zrc20), message, revertOptions);

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount, ownerBalanceAfter);
    }

    function testWithdrawZETAFailsIfAmountIsZero() public {
        vm.expectRevert(InsufficientZetaAmount.selector);
        gateway.withdraw(0, 1, revertOptions);
    }

    function testWithdrawAndcallZETAFailsIfAmountIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(InsufficientZetaAmount.selector);
        gateway.withdrawAndCall(0, 1, message, revertOptions);
    }

    function testWithdrawZETA() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 fungibleModuleBalanceBefore = fungibleModule.balance;
        uint256 chainId = 1;

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawal(
            owner, chainId, abi.encodePacked(fungibleModule), address(zetaToken), amount, 0, 0, "", revertOptions
        );
        gateway.withdraw(amount, chainId, revertOptions);

        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore - 1, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        // Verify amount is transfered to fungible module
        assertEq(fungibleModuleBalanceBefore + 1, fungibleModule.balance);
    }

    function testWithdrawZETAFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 fungibleModuleBalanceBefore = fungibleModule.balance;
        uint256 chainId = 1;

        // Remove allowance for gateway
        vm.prank(owner);
        zetaToken.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdraw(amount, chainId, revertOptions);

        // Verify balances not changed
        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        assertEq(fungibleModuleBalanceBefore, fungibleModule.balance);
    }

    function testWithdrawZETAFailsIfNoBalance() public {
        uint256 amount = 1;
        uint256 ownerBalance = zetaToken.balanceOf(owner);
        zetaToken.transfer(address(0x123), ownerBalance);
        uint256 chainId = 1;

        vm.expectRevert();
        gateway.withdraw(amount, chainId, revertOptions);
    }

    function testWithdrawZETAWithMessage() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 fungibleModuleBalanceBefore = fungibleModule.balance;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 chainId = 1;

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawal(
            owner, chainId, abi.encodePacked(fungibleModule), address(zetaToken), amount, 0, 0, message, revertOptions
        );
        gateway.withdrawAndCall(amount, chainId, message, revertOptions);

        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore - 1, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        // Verify amount is transfered to fungible module
        assertEq(fungibleModuleBalanceBefore + 1, fungibleModule.balance);
    }

    function testWithdrawZETAWithMessageFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 fungibleModuleBalanceBefore = fungibleModule.balance;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 chainId = 1;

        // Remove allowance for gateway
        vm.prank(owner);
        zetaToken.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdrawAndCall(amount, chainId, message, revertOptions);

        // Verify balances not changed
        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        assertEq(fungibleModuleBalanceBefore, fungibleModule.balance);
    }

    function testCallFailsIfReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(ZeroAddress.selector);
        gateway.call(abi.encodePacked(""), 1, message, revertOptions);
    }

    function testCall() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectEmit(true, true, true, true, address(gateway));
        uint256 chainId = 1;

        emit Call(owner, chainId, abi.encodePacked(addr1), message, revertOptions);
        gateway.call(abi.encodePacked(addr1), chainId, message, revertOptions);
    }
}

contract GatewayZEVMOutboundTest is Test, IGatewayZEVMEvents, IGatewayZEVMErrors {
    address payable proxy;
    GatewayZEVM gateway;
    ZRC20 zrc20;
    WETH9 zetaToken;
    SystemContract systemContract;
    TestZContract testZContract;
    address owner;
    address addr1;
    address fungibleModule;

    event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message);
    event ContextDataRevert(bytes origin, address sender, uint256 chainID, address msgSender, string message);

    error ZeroAddress();
    error EnforcedPause();
    error AccessControlUnauthorizedAccount(address account, bytes32 neededRole);

    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1234);

        zetaToken = new WETH9();

        proxy = payable(
            Upgrades.deployUUPSProxy(
                "GatewayZEVM.sol", abi.encodeCall(GatewayZEVM.initialize, (address(zetaToken), owner))
            )
        );
        gateway = GatewayZEVM(proxy);

        fungibleModule = gateway.FUNGIBLE_MODULE_ADDRESS();

        testZContract = new TestZContract();

        vm.startPrank(fungibleModule);
        systemContract = new SystemContract(address(0), address(0), address(0));
        zrc20 = new ZRC20("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        vm.deal(fungibleModule, 1_000_000_000);
        zetaToken.deposit{ value: 10 }();
        zetaToken.approve(address(gateway), 10);
        zrc20.deposit(owner, 100_000);
        vm.stopPrank();

        vm.startPrank(owner);
        zrc20.approve(address(gateway), 100_000);
        zetaToken.deposit{ value: 10 }();
        zetaToken.approve(address(gateway), 10);
        vm.stopPrank();
    }

    function testDepositFailsIfZrc20IsZeroAddress() public {
        vm.prank(fungibleModule);
        vm.expectRevert(ZeroAddress.selector);
        gateway.deposit(address(0), 1, addr1);
    }

    function testDepositFailsIfTargetIsZeroAddress() public {
        vm.prank(fungibleModule);
        vm.expectRevert(ZeroAddress.selector);
        gateway.deposit(address(zrc20), 1, address(0));
    }

    function testDepositFailsIfAmountIs0() public {
        vm.prank(fungibleModule);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.deposit(address(zrc20), 0, addr1);
    }

    function testDeposit() public {
        uint256 amount = 1;
        uint256 balanceBefore = zrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(fungibleModule);
        gateway.deposit(address(zrc20), amount, addr1);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testDepositTogglePause() public {
        vm.prank(fungibleModule);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, fungibleModule, PAUSER_ROLE));
        gateway.pause();

        vm.prank(fungibleModule);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, fungibleModule, PAUSER_ROLE));
        gateway.unpause();

        vm.prank(owner);
        gateway.pause();

        uint256 amount = 1;

        vm.expectRevert(EnforcedPause.selector);
        vm.prank(fungibleModule);
        gateway.deposit(address(zrc20), amount, addr1);

        vm.prank(owner);
        gateway.unpause();

        uint256 balanceBefore = zrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(fungibleModule);
        gateway.deposit(address(zrc20), amount, addr1);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testDepositFailsIfSenderNotFungibleModule() public {
        uint256 amount = 1;
        uint256 balanceBefore = zrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(owner);
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        gateway.deposit(address(zrc20), amount, addr1);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(0, balanceAfter);
    }

    function testDepositFailsIfTargetIsGateway() public {
        uint256 amount = 1;

        vm.prank(fungibleModule);
        vm.expectRevert(InvalidTarget.selector);
        gateway.deposit(address(zrc20), amount, address(gateway));
    }

    function testDepositFailsIfTargetIsFungibleModule() public {
        uint256 amount = 1;
        vm.prank(fungibleModule);
        vm.expectRevert(InvalidTarget.selector);
        gateway.deposit(address(zrc20), amount, fungibleModule);
    }

    function testExecuteFailsIfZrc20IsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(context, address(0), 1, address(testZContract), message);
    }

    function testExecuteFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(context, address(zrc20), 1, address(0), message);
    }

    function testExecuteFailsIfAmountIsZero() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.execute(context, address(zrc20), 0, address(testZContract), message);
    }

    function testExecuteZContract() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testZContract));
        emit ContextData(abi.encodePacked(gateway), fungibleModule, 1, address(gateway), "hello");
        vm.prank(fungibleModule);
        gateway.execute(context, address(zrc20), 1, address(testZContract), message);
    }

    function testExecuteZContractFailsIfSenderIsNotFungibleModule() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectRevert(CallerIsNotFungibleModule.selector);
        vm.prank(owner);
        gateway.execute(context, address(zrc20), 1, address(testZContract), message);
    }

    function testExecuteRevertZContractFailsIfZrc20IsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(ZeroAddress.selector);
        gateway.executeRevert(context, address(0), 1, address(testZContract), message);
    }

    function testExecuteRevertZContractFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(ZeroAddress.selector);
        gateway.executeRevert(context, address(zrc20), 1, address(0), message);
    }

    function testExecuteRevertZContractFailsIfAmountIsZero() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.executeRevert(context, address(zrc20), 0, address(testZContract), message);
    }

    function testExecuteRevertZContract() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testZContract));
        emit ContextDataRevert(abi.encodePacked(gateway), fungibleModule, 1, address(gateway), "hello");
        vm.prank(fungibleModule);
        gateway.executeRevert(context, address(zrc20), 1, address(testZContract), message);
    }

    function testExecuteRevertZContractIfSenderIsNotFungibleModule() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectRevert(CallerIsNotFungibleModule.selector);
        vm.prank(owner);
        gateway.executeRevert(context, address(zrc20), 1, address(testZContract), message);
    }

    function testDepositZRC20AndCallZContractFailsIfZrc20IsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, address(0), 1, address(testZContract), message);
    }

    function testDepositZRC20AndCallZContractFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, address(zrc20), 1, address(0), message);
    }

    function testDepositZRC20AndCallZContractFailsIfAmountIsZero() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.depositAndCall(context, address(zrc20), 0, address(testZContract), message);
    }

    function testDepositZRC20AndCallZContract() public {
        uint256 balanceBefore = zrc20.balanceOf(address(testZContract));
        assertEq(0, balanceBefore);

        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testZContract));
        emit ContextData(abi.encodePacked(gateway), fungibleModule, 1, address(gateway), "hello");
        vm.prank(fungibleModule);
        gateway.depositAndCall(context, address(zrc20), 1, address(testZContract), message);

        uint256 balanceAfter = zrc20.balanceOf(address(testZContract));
        assertEq(1, balanceAfter);
    }

    function testDepositZRC20AndCallZContractFailsIfSenderIsNotFungibleModule() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectRevert(CallerIsNotFungibleModule.selector);
        vm.prank(owner);
        gateway.depositAndCall(context, address(zrc20), 1, address(testZContract), message);
    }

    function testDepositZRC20AndCallZContractIfTargetIsFungibleModule() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(fungibleModule);
        gateway.depositAndCall(context, address(zrc20), 1, fungibleModule, message);
    }

    function testDepositZRC20AndCallZContractIfTargetIsGateway() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(fungibleModule);
        gateway.depositAndCall(context, address(zrc20), 1, address(gateway), message);
    }

    function testDepositAndRevertZRC20AndCallZContractFailsIfZrc20IsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndRevert(context, address(0), 1, address(testZContract), message);
    }

    function testDepositAndRevertZRC20AndCallZContractFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndRevert(context, address(zrc20), 1, address(0), message);
    }

    function testDepositAndRevertZRC20AndCallZContractFailsIfAmountIsZero() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.depositAndRevert(context, address(zrc20), 0, address(testZContract), message);
    }

    function testDepositAndRevertZRC20AndCallZContract() public {
        uint256 balanceBefore = zrc20.balanceOf(address(testZContract));
        assertEq(0, balanceBefore);

        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testZContract));
        emit ContextDataRevert(abi.encodePacked(gateway), fungibleModule, 1, address(gateway), "hello");
        vm.prank(fungibleModule);
        gateway.depositAndRevert(context, address(zrc20), 1, address(testZContract), message);

        uint256 balanceAfter = zrc20.balanceOf(address(testZContract));
        assertEq(1, balanceAfter);
    }

    function testDepositAndRevertZRC20AndCallZContractFailsIfSenderIsNotFungibleModule() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectRevert(CallerIsNotFungibleModule.selector);
        vm.prank(owner);
        gateway.depositAndRevert(context, address(zrc20), 1, address(testZContract), message);
    }

    function testDepositAndRevertZRC20AndCallZContractFailsITargetIsFungibleModule() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(fungibleModule);
        gateway.depositAndRevert(context, address(zrc20), 1, fungibleModule, message);
    }

    function testDepositAndRevertZRC20AndCallZContractFailsITargetIsGateway() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(fungibleModule);
        gateway.depositAndRevert(context, address(zrc20), 1, address(gateway), message);
    }

    function testDepositZETAAndCallZContractFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, 1, address(0), message);
    }

    function testDepositZETAAndCallZContractFailsIfTargetIsAmountIsZero() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.prank(fungibleModule);
        vm.expectRevert(InsufficientZetaAmount.selector);
        gateway.depositAndCall(context, 0, address(zrc20), message);
    }

    function testDepositZETAAndCallZContract() public {
        uint256 amount = 1;
        uint256 fungibleBalanceBefore = zetaToken.balanceOf(fungibleModule);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 destinationBalanceBefore = address(testZContract).balance;
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testZContract));
        emit ContextData(abi.encodePacked(gateway), fungibleModule, amount, address(gateway), "hello");
        vm.prank(fungibleModule);
        gateway.depositAndCall(context, amount, address(testZContract), message);

        uint256 fungibleBalanceAfter = zetaToken.balanceOf(fungibleModule);
        assertEq(fungibleBalanceBefore - amount, fungibleBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        // Verify amount is transfered to destination
        assertEq(destinationBalanceBefore + amount, address(testZContract).balance);
    }

    function testDepositZETAAndCallZContractFailsIfSenderIsNotFungibleModule() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectRevert(CallerIsNotFungibleModule.selector);
        vm.prank(owner);
        gateway.depositAndCall(context, amount, address(testZContract), message);
    }

    function testDepositZETAAndCallZContractFailsIfTargetIsFungibleModule() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(fungibleModule);
        gateway.depositAndCall(context, amount, fungibleModule, message);
    }

    function testDepositZETAAndCallZContractFailsIfTargetIsGateway() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: fungibleModule, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(fungibleModule);
        gateway.depositAndCall(context, amount, address(gateway), message);
    }
}
