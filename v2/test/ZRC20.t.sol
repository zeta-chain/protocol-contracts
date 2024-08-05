// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "./utils/SystemContract.sol";
import "./utils/WZETA.sol";

import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";
import "src/zevm/GatewayZEVM.sol";
import "src/zevm/ZRC20.sol";
import "src/zevm/interfaces/IGatewayZEVM.sol";
import "src/zevm/interfaces/IZRC20.sol";

contract ZRC20Test is Test, ZRC20Errors {
    ZRC20 zrc20;
    SystemContract systemContract;
    GatewayZEVM gateway;
    address payable proxy;
    WETH9 zetaToken;

    address owner;
    address addr1;
    address fungibleModule;

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

        vm.startPrank(fungibleModule);
        systemContract = new SystemContract(address(0), address(0), address(0));
        zrc20 = new ZRC20("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        vm.deal(fungibleModule, 1_000_000_000);
        vm.deal(proxy, 1_000_000_000);
        zrc20.deposit(owner, 100_000);
        vm.stopPrank();
    }

    function testZRC20BasicInfo() public {
        string memory name = zrc20.name();
        assertEq("TOKEN", name);

        string memory symbol = zrc20.symbol();
        assertEq("TKN", symbol);

        uint8 decimals = zrc20.decimals();
        assertEq(18, decimals);

        uint256 totalSupply = zrc20.totalSupply();
        assertEq(100_000, totalSupply);
    }

    function testTransferFrom() public {
        uint256 balanceStart = zrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 50_000;
        zrc20.approve(owner, amount);
        zrc20.transferFrom(owner, addr1, amount);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testTransferFromFailsIfNoAllowance() public {
        uint256 balanceStart = zrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 50_000;
        vm.expectRevert(LowAllowance.selector);
        zrc20.transferFrom(owner, addr1, amount);
    }

    function testTransferFromFailsIfNoBalance() public {
        uint256 balanceStart = zrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 500_000;
        zrc20.approve(owner, amount);
        vm.expectRevert(LowBalance.selector);
        zrc20.transferFrom(owner, addr1, amount);
    }

    function testTransferFromFailsIfSenderIsZeroAddress() public {
        uint256 balanceStart = zrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 500_000;
        zrc20.approve(owner, amount);
        vm.expectRevert(ZeroAddress.selector);
        zrc20.transferFrom(address(0), addr1, amount);
    }

    function testTransferFromFailsIfRecipientIsZeroAddress() public {
        uint256 amount = 500_000;
        zrc20.approve(owner, amount);
        vm.expectRevert(ZeroAddress.selector);
        zrc20.transferFrom(owner, address(0), amount);
    }

    function testBurn() public {
        uint256 balanceStart = zrc20.balanceOf(owner);
        assertEq(100_000, balanceStart);
        uint256 totalSupplyStart = zrc20.totalSupply();
        assertEq(100_000, totalSupplyStart);

        zrc20.burn(50_000);
        uint256 balanceAfter = zrc20.balanceOf(owner);
        assertEq(50_000, balanceAfter);
        uint256 totalSupplyAfter = zrc20.totalSupply();
        assertEq(50_000, totalSupplyAfter);
    }

    function testBurnFailsIfNoBalance() public {
        vm.expectRevert(LowBalance.selector);
        zrc20.burn(150_000);
    }

    function testDeposit() public {
        uint256 totalSupplyStart = zrc20.totalSupply();
        assertEq(100_000, totalSupplyStart);

        vm.prank(proxy);
        zrc20.deposit(owner, 100_000);

        uint256 totalSupplyEnd = zrc20.totalSupply();
        assertEq(200_000, totalSupplyEnd);
    }

    function testWithdrawGasFee() public {
        vm.prank(fungibleModule);
        uint256 gasLimit = 10;
        uint256 protocolFlatFee = 10;

        zrc20.updateGasLimit(10);

        vm.prank(fungibleModule);
        zrc20.updateProtocolFlatFee(10);

        (address gasZRC20, uint256 gasFee) = zrc20.withdrawGasFee();
        assertEq(address(zrc20), gasZRC20);
        assertEq(gasLimit + protocolFlatFee, gasFee);
    }

    function testWithdrawGasFeeFailsIfGasCoinNotSetForChainId() public {
        vm.prank(fungibleModule);
        systemContract.setGasCoinZRC20(1, address(0));

        vm.expectRevert(ZeroGasCoin.selector);
        (address gasZRC20, uint256 gasFee) = zrc20.withdrawGasFee();
    }

    function testWithdrawGasFeeFailsIfGasPriceNotSetForChainId() public {
        vm.prank(fungibleModule);
        systemContract.setGasPrice(1, 0);

        vm.expectRevert(ZeroGasPrice.selector);
        (address gasZRC20, uint256 gasFee) = zrc20.withdrawGasFee();
    }

    function testWithdraw() public {
        uint256 gasLimit = 10;
        uint256 protocolFlatFee = 10;

        vm.prank(fungibleModule);
        zrc20.updateGasLimit(gasLimit);

        vm.prank(fungibleModule);
        zrc20.updateProtocolFlatFee(protocolFlatFee);
        

        uint256 balanceStart = zrc20.balanceOf(owner);
        assertEq(100_000, balanceStart);
        uint256 totalSupplyStart = zrc20.totalSupply();
        assertEq(100_000, totalSupplyStart);

        uint256 fungibleModuleBalanceStart = zrc20.balanceOf(fungibleModule);

        zrc20.approve(address(zrc20), 50_000);
        zrc20.withdraw(abi.encodePacked(addr1), 50_000);

        uint256 fungibleModuleBalanceAfter = zrc20.balanceOf(fungibleModule);
        assertEq(fungibleModuleBalanceStart + gasLimit + protocolFlatFee, fungibleModuleBalanceAfter);

        uint256 balanceAfter = zrc20.balanceOf(owner);
        assertEq(50_000 - gasLimit - protocolFlatFee, balanceAfter);
        uint256 totalSupplyAfter = zrc20.totalSupply();
        assertEq(50_000, totalSupplyAfter);
    }

    function testDepositFailsIfSenderIsNotGateway() public {
        vm.expectRevert(InvalidSender.selector);
        zrc20.deposit(owner, 100_000);
    }

    function testUpdateSystemContractAddress() public {
        vm.prank(fungibleModule);
        zrc20.updateSystemContractAddress(address(0x3211));
        assertEq(zrc20.systemContractAddress(), address(0x3211));
    }

    function testUpdateSystemContractAddressFailsIfSenderIsNotFungible() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        zrc20.updateSystemContractAddress(address(0x3211));
    }

    function testUpdateGatewayAddress() public {
        vm.prank(fungibleModule);
        zrc20.updateGatewayAddress(address(0x3211));
        assertEq(zrc20.gatewayAddress(), address(0x3211));
    }

    function testUpdateGatewayAddressFailsIfSenderIsNotFungible() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        zrc20.updateGatewayAddress(address(0x3211));
    }

    function testUpdateGasLimit() public {
        vm.prank(fungibleModule);
        zrc20.updateGasLimit(10);
        assertEq(10, zrc20.gasLimit());
    }

    function testUpdateGasLimitFailsIfSenderIsNotFungible() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        zrc20.updateGasLimit(10);
    }

    function testUpdateProtocolFlatFee() public {
        vm.prank(fungibleModule);
        zrc20.updateProtocolFlatFee(10);
        assertEq(10, zrc20.protocolFlatFee());
    }

    function testUpdateProtocolFlatFeeFailsIfSenderIsNotFungible() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        zrc20.updateProtocolFlatFee(10);
    }
}
