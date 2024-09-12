// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "../contracts/zevm/SystemContract.sol";

import "./utils/TestUniversalContract.sol";

import "./utils/WZETA.sol";

import "../contracts/zevm/GatewayZEVM.sol";
import "../contracts/zevm/ZRC20.sol";
import "../contracts/zevm/interfaces/IGatewayZEVM.sol";
import "../contracts/zevm/interfaces/IZRC20.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract GatewayZEVMInboundTest is Test, IGatewayZEVMEvents, IGatewayZEVMErrors {
    address payable proxy;
    GatewayZEVM gateway;
    ZRC20 zrc20;
    WETH9 zetaToken;
    SystemContract systemContract;
    TestUniversalContract testUniversalContract;
    address owner;
    address addr1;
    address protocolAddress;
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

        protocolAddress = gateway.PROTOCOL_ADDRESS();
        testUniversalContract = new TestUniversalContract();

        vm.startPrank(protocolAddress);
        systemContract = new SystemContract(address(0), address(0), address(0));
        zrc20 = new ZRC20("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        vm.deal(protocolAddress, 1_000_000_000);
        zetaToken.deposit{ value: 10 }();
        zetaToken.approve(address(gateway), 10);
        zrc20.deposit(owner, 100_000);
        vm.stopPrank();

        vm.startPrank(owner);
        zrc20.approve(address(gateway), 100_000);
        zetaToken.deposit{ value: 10 }();
        zetaToken.approve(address(gateway), 10);
        vm.stopPrank();

        revertOptions = RevertOptions({
            revertAddress: address(0x321),
            callOnRevert: true,
            abortAddress: address(0x321),
            revertMessage: "",
            onRevertGasLimit: 0
        });
    }

    function testWithdrawZRC20() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawn(
            owner,
            0,
            abi.encodePacked(addr1),
            address(zrc20),
            amount,
            0,
            zrc20.PROTOCOL_FLAT_FEE(),
            "",
            CallOptions({ gasLimit: 0, isArbitraryCall: true }),
            revertOptions
        );
        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20), revertOptions);

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount, ownerBalanceAfter);
    }

    function testWithdrawZRC20FailsIfNoBalanceForGasFee() public {
        uint256 amount = 1;
        uint256 ownerBalance = zrc20.balanceOf(owner);
        zrc20.transfer(address(0x123), ownerBalance);

        vm.prank(protocolAddress);
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
        gateway.withdrawAndCall(abi.encodePacked(""), 1, address(zrc20), message, 1, revertOptions);
    }

    function testWithdrawAndCallZRC20FailsIfAmountIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.withdrawAndCall(abi.encodePacked(addr1), 0, address(zrc20), message, 1, revertOptions);
    }

    function testWithdrawZRC20WithMessageFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        // Remove allowance for gateway
        vm.prank(owner);
        zrc20.approve(address(gateway), 0);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert();
        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, address(zrc20), message, 1, revertOptions);

        // Check that balance didn't change
        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);
    }

    function testWithdrawZRC20WithMessage() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 expectedGasFee = 1;
        uint256 gasLimit = 1;
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawn(
            owner,
            0,
            abi.encodePacked(addr1),
            address(zrc20),
            amount,
            expectedGasFee,
            zrc20.PROTOCOL_FLAT_FEE(),
            message,
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
            revertOptions
        );
        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, address(zrc20), message, gasLimit, revertOptions);

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount - expectedGasFee, ownerBalanceAfter);
    }

    function testWithdrawZETAFailsIfAmountIsZero() public {
        vm.expectRevert(InsufficientZetaAmount.selector);
        gateway.withdraw(abi.encodePacked(addr1), 0, 1, revertOptions);
    }

    function testWithdrawZETAFailsIfReceiverIsZeroAddress() public {
        vm.expectRevert(ZeroAddress.selector);
        gateway.withdraw(abi.encodePacked(""), 0, 1, revertOptions);
    }

    function testWithdrawAndCallZETAFailsIfAmountIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(InsufficientZetaAmount.selector);
        gateway.withdrawAndCall(abi.encodePacked(addr1), 0, 1, message, revertOptions);
    }

    function testWithdrawAndCallZETAFailsIfAmountIsReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(ZeroAddress.selector);
        gateway.withdrawAndCall(abi.encodePacked(""), 1, 1, message, revertOptions);
    }

    function testWithdrawZETA() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 protocolBalanceBefore = protocolAddress.balance;
        uint256 chainId = 1;

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawn(
            owner,
            chainId,
            abi.encodePacked(addr1),
            address(zetaToken),
            amount,
            0,
            0,
            "",
            CallOptions({ gasLimit: 0, isArbitraryCall: true }),
            revertOptions
        );
        gateway.withdraw(abi.encodePacked(addr1), amount, chainId, revertOptions);

        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore - 1, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        // Verify amount is transfered to protocol address
        assertEq(protocolBalanceBefore + 1, protocolAddress.balance);
    }

    function testWithdrawZETAFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 protocolBalanceBefore = protocolAddress.balance;
        uint256 chainId = 1;

        // Remove allowance for gateway
        vm.prank(owner);
        zetaToken.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdraw(abi.encodePacked(addr1), amount, chainId, revertOptions);

        // Verify balances not changed
        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        assertEq(protocolBalanceBefore, protocolAddress.balance);
    }

    function testWithdrawZETAFailsIfNoBalance() public {
        uint256 amount = 1;
        uint256 ownerBalance = zetaToken.balanceOf(owner);
        zetaToken.transfer(address(0x123), ownerBalance);
        uint256 chainId = 1;

        vm.expectRevert();
        gateway.withdraw(abi.encodePacked(addr1), amount, chainId, revertOptions);
    }

    function testWithdrawZETAWithMessage() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 protocolBalanceBefore = protocolAddress.balance;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 chainId = 1;

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawn(
            owner,
            chainId,
            abi.encodePacked(addr1),
            address(zetaToken),
            amount,
            0,
            0,
            message,
            CallOptions({ gasLimit: 0, isArbitraryCall: true }),
            revertOptions
        );
        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, chainId, message, revertOptions);

        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore - 1, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        // Verify amount is transfered to protocol account
        assertEq(protocolBalanceBefore + 1, protocolAddress.balance);
    }

    function testWithdrawZETAWithMessageFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 protocolBalanceBefore = protocolAddress.balance;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 chainId = 1;

        // Remove allowance for gateway
        vm.prank(owner);
        zetaToken.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, chainId, message, revertOptions);

        // Verify balances not changed
        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        assertEq(protocolBalanceBefore, protocolAddress.balance);
    }

    function testCallFailsIfReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(ZeroAddress.selector);
        gateway.call(abi.encodePacked(""), address(zrc20), message, 1, revertOptions);
    }

    function testCall() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectEmit(true, true, true, true, address(gateway));

        emit Called(
            owner,
            address(zrc20),
            abi.encodePacked(addr1),
            message,
            CallOptions({ gasLimit: 1, isArbitraryCall: true }),
            revertOptions
        );
        gateway.call(abi.encodePacked(addr1), address(zrc20), message, 1, revertOptions);
    }
}

contract GatewayZEVMOutboundTest is Test, IGatewayZEVMEvents, IGatewayZEVMErrors {
    address payable proxy;
    GatewayZEVM gateway;
    ZRC20 zrc20;
    WETH9 zetaToken;
    SystemContract systemContract;
    TestUniversalContract testUniversalContract;
    address owner;
    address addr1;
    address protocolAddress;
    RevertOptions revertOptions;
    RevertContext revertContext;

    event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message);
    event ContextDataRevert(RevertContext revertContext);

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

        protocolAddress = gateway.PROTOCOL_ADDRESS();

        testUniversalContract = new TestUniversalContract();

        vm.startPrank(protocolAddress);
        systemContract = new SystemContract(address(0), address(0), address(0));
        zrc20 = new ZRC20("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        vm.deal(protocolAddress, 1_000_000_000);
        zetaToken.deposit{ value: 10 }();
        zetaToken.approve(address(gateway), 10);
        zrc20.deposit(owner, 100_000);
        vm.stopPrank();

        vm.startPrank(owner);
        zrc20.approve(address(gateway), 100_000);
        zetaToken.deposit{ value: 10 }();
        zetaToken.approve(address(gateway), 10);
        vm.stopPrank();

        revertContext = RevertContext({ asset: address(0), amount: 1, revertMessage: "" });
    }

    function testDepositFailsIfZRC20IsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.deposit(address(0), 1, addr1);
    }

    function testDepositFailsIfTargetIsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.deposit(address(zrc20), 1, address(0));
    }

    function testDepositFailsIfAmountIs0() public {
        vm.prank(protocolAddress);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.deposit(address(zrc20), 0, addr1);
    }

    function testDeposit() public {
        uint256 amount = 1;
        uint256 balanceBefore = zrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(protocolAddress);
        gateway.deposit(address(zrc20), amount, addr1);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testDepositTogglePause() public {
        vm.prank(protocolAddress);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, protocolAddress, PAUSER_ROLE));
        gateway.pause();

        vm.prank(protocolAddress);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, protocolAddress, PAUSER_ROLE));
        gateway.unpause();

        vm.prank(owner);
        gateway.pause();

        uint256 amount = 1;

        vm.expectRevert(EnforcedPause.selector);
        vm.prank(protocolAddress);
        gateway.deposit(address(zrc20), amount, addr1);

        vm.prank(owner);
        gateway.unpause();

        uint256 balanceBefore = zrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(protocolAddress);
        gateway.deposit(address(zrc20), amount, addr1);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testDepositFailsIfSenderNotProtocol() public {
        uint256 amount = 1;
        uint256 balanceBefore = zrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(owner);
        vm.expectRevert(CallerIsNotProtocol.selector);
        gateway.deposit(address(zrc20), amount, addr1);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(0, balanceAfter);
    }

    function testDepositFailsIfTargetIsGateway() public {
        uint256 amount = 1;

        vm.prank(protocolAddress);
        vm.expectRevert(InvalidTarget.selector);
        gateway.deposit(address(zrc20), amount, address(gateway));
    }

    function testDepositFailsIfTargetIsProtocol() public {
        uint256 amount = 1;
        vm.prank(protocolAddress);
        vm.expectRevert(InvalidTarget.selector);
        gateway.deposit(address(zrc20), amount, protocolAddress);
    }

    function testExecuteFailsIfZRC20IsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(context, address(0), 1, address(testUniversalContract), message);
    }

    function testExecuteFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(context, address(zrc20), 1, address(0), message);
    }

    function testExecuteUniversalContractFailsIfZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(context, address(0), 1, address(testUniversalContract), message);
    }

    function testExecuteUniversalContract() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextData(abi.encodePacked(gateway), protocolAddress, 1, address(gateway), "hello");
        vm.prank(protocolAddress);
        gateway.execute(context, address(zrc20), 1, address(testUniversalContract), message);
    }

    function testExecuteUniversalContractFailsIfSenderIsNotProtocol() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.execute(context, address(zrc20), 1, address(testUniversalContract), message);
    }

    function testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.executeRevert(address(0), revertContext);
    }

    function testExecuteRevertUniversalContract() public {
        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextDataRevert(revertContext);
        vm.prank(protocolAddress);
        gateway.executeRevert(address(testUniversalContract), revertContext);
    }

    function testExecuteRevertUniversalContractIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.executeRevert(address(testUniversalContract), revertContext);
    }

    function testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, address(0), 1, address(testUniversalContract), message);
    }

    function testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, address(zrc20), 1, address(0), message);
    }

    function testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.depositAndCall(context, address(zrc20), 0, address(testUniversalContract), message);
    }

    function testDepositZRC20AndCallUniversalContract() public {
        uint256 balanceBefore = zrc20.balanceOf(address(testUniversalContract));
        assertEq(0, balanceBefore);

        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextData(abi.encodePacked(gateway), protocolAddress, 1, address(gateway), "hello");
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, address(zrc20), 1, address(testUniversalContract), message);

        uint256 balanceAfter = zrc20.balanceOf(address(testUniversalContract));
        assertEq(1, balanceAfter);
    }

    function testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.depositAndCall(context, address(zrc20), 1, address(testUniversalContract), message);
    }

    function testDepositZRC20AndCallUniversalContractIfTargetIsProtocol() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, address(zrc20), 1, protocolAddress, message);
    }

    function testDepositZRC20AndCallUniversalContractIfTargetIsGateway() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, address(zrc20), 1, address(gateway), message);
    }

    function testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndRevert(address(0), 1, address(testUniversalContract), revertContext);
    }

    function testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndRevert(address(zrc20), 1, address(0), revertContext);
    }

    function testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() public {
        vm.prank(protocolAddress);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.depositAndRevert(address(zrc20), 0, address(testUniversalContract), revertContext);
    }

    function testDepositAndRevertZRC20AndCallUniversalContract() public {
        uint256 balanceBefore = zrc20.balanceOf(address(testUniversalContract));
        assertEq(0, balanceBefore);

        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextDataRevert(revertContext);
        vm.prank(protocolAddress);
        gateway.depositAndRevert(address(zrc20), 1, address(testUniversalContract), revertContext);

        uint256 balanceAfter = zrc20.balanceOf(address(testUniversalContract));
        assertEq(1, balanceAfter);
    }

    function testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.depositAndRevert(address(zrc20), 1, address(testUniversalContract), revertContext);
    }

    function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() public {
        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndRevert(address(zrc20), 1, protocolAddress, revertContext);
    }

    function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() public {
        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndRevert(address(zrc20), 1, address(gateway), revertContext);
    }

    function testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, 1, address(0), message);
    }

    function testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(InsufficientZetaAmount.selector);
        gateway.depositAndCall(context, 0, address(zrc20), message);
    }

    function testDepositZETAAndCallUniversalContractFailsIfZeroAddress() public {
        uint256 destinationBalanceBefore = address(testUniversalContract).balance;
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, 1, address(0), message);
    }

    function testDepositZETAAndCallUniversal() public {
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, 1, address(0), message);
    }

    function testDepositZETAAndCallUniversalContract() public {
        uint256 amount = 1;
        uint256 protocolBalanceBefore = zetaToken.balanceOf(protocolAddress);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 destinationBalanceBefore = address(testUniversalContract).balance;
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextData(abi.encodePacked(gateway), protocolAddress, amount, address(gateway), "hello");
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, amount, address(testUniversalContract), message);

        uint256 protocolBalanceAfter = zetaToken.balanceOf(protocolAddress);
        assertEq(protocolBalanceBefore - amount, protocolBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        // Verify amount is transfered to destination
        assertEq(destinationBalanceBefore + amount, address(testUniversalContract).balance);
    }

    function testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.depositAndCall(context, amount, address(testUniversalContract), message);
    }

    function testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, amount, protocolAddress, message);
    }

    function testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        zContext memory context =
            zContext({ origin: abi.encodePacked(address(gateway)), sender: protocolAddress, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, amount, address(gateway), message);
    }
}
