/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { ethers } from "ethers";
import {
  FactoryOptions,
  HardhatEthersHelpers as HardhatEthersHelpersBase,
} from "@nomiclabs/hardhat-ethers/types";

import * as Contracts from ".";

declare module "hardhat/types/runtime" {
  interface HardhatEthersHelpers extends HardhatEthersHelpersBase {
    getContractFactory(
      name: "OwnableUpgradeable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.OwnableUpgradeable__factory>;
    getContractFactory(
      name: "IERC1822ProxiableUpgradeable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IERC1822ProxiableUpgradeable__factory>;
    getContractFactory(
      name: "IERC1967Upgradeable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IERC1967Upgradeable__factory>;
    getContractFactory(
      name: "IBeaconUpgradeable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IBeaconUpgradeable__factory>;
    getContractFactory(
      name: "ERC1967UpgradeUpgradeable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC1967UpgradeUpgradeable__factory>;
    getContractFactory(
      name: "Initializable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.Initializable__factory>;
    getContractFactory(
      name: "UUPSUpgradeable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.UUPSUpgradeable__factory>;
    getContractFactory(
      name: "ContextUpgradeable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ContextUpgradeable__factory>;
    getContractFactory(
      name: "Ownable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.Ownable__factory>;
    getContractFactory(
      name: "Ownable2Step",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.Ownable2Step__factory>;
    getContractFactory(
      name: "Pausable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.Pausable__factory>;
    getContractFactory(
      name: "ERC20",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC20__factory>;
    getContractFactory(
      name: "IERC20Permit",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IERC20Permit__factory>;
    getContractFactory(
      name: "ERC20Burnable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC20Burnable__factory>;
    getContractFactory(
      name: "IERC20Metadata",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IERC20Metadata__factory>;
    getContractFactory(
      name: "IERC20",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IERC20__factory>;
    getContractFactory(
      name: "IERC20",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IERC20__factory>;
    getContractFactory(
      name: "IUniswapV2Callee",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV2Callee__factory>;
    getContractFactory(
      name: "IUniswapV2ERC20",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV2ERC20__factory>;
    getContractFactory(
      name: "IUniswapV2Factory",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV2Factory__factory>;
    getContractFactory(
      name: "IUniswapV2Pair",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV2Pair__factory>;
    getContractFactory(
      name: "UniswapV2ERC20",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.UniswapV2ERC20__factory>;
    getContractFactory(
      name: "UniswapV2Factory",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.UniswapV2Factory__factory>;
    getContractFactory(
      name: "UniswapV2Pair",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.UniswapV2Pair__factory>;
    getContractFactory(
      name: "IERC20",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IERC20__factory>;
    getContractFactory(
      name: "IUniswapV2Router01",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV2Router01__factory>;
    getContractFactory(
      name: "IUniswapV2Router02",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV2Router02__factory>;
    getContractFactory(
      name: "IWETH",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IWETH__factory>;
    getContractFactory(
      name: "UniswapV2Router02",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.UniswapV2Router02__factory>;
    getContractFactory(
      name: "IUniswapV3SwapCallback",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV3SwapCallback__factory>;
    getContractFactory(
      name: "IUniswapV3Factory",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV3Factory__factory>;
    getContractFactory(
      name: "IUniswapV3Pool",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV3Pool__factory>;
    getContractFactory(
      name: "IUniswapV3PoolActions",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV3PoolActions__factory>;
    getContractFactory(
      name: "IUniswapV3PoolDerivedState",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV3PoolDerivedState__factory>;
    getContractFactory(
      name: "IUniswapV3PoolEvents",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV3PoolEvents__factory>;
    getContractFactory(
      name: "IUniswapV3PoolImmutables",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV3PoolImmutables__factory>;
    getContractFactory(
      name: "IUniswapV3PoolOwnerActions",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV3PoolOwnerActions__factory>;
    getContractFactory(
      name: "IUniswapV3PoolState",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV3PoolState__factory>;
    getContractFactory(
      name: "IQuoter",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IQuoter__factory>;
    getContractFactory(
      name: "ISwapRouter",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ISwapRouter__factory>;
    getContractFactory(
      name: "ERC20Custody",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC20Custody__factory>;
    getContractFactory(
      name: "ConnectorErrors",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ConnectorErrors__factory>;
    getContractFactory(
      name: "ZetaErrors",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaErrors__factory>;
    getContractFactory(
      name: "ZetaInteractorErrors",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaInteractorErrors__factory>;
    getContractFactory(
      name: "ZetaCommonErrors",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaCommonErrors__factory>;
    getContractFactory(
      name: "ZetaConnector",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaConnector__factory>;
    getContractFactory(
      name: "ZetaReceiver",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaReceiver__factory>;
    getContractFactory(
      name: "ZetaTokenConsumer",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaTokenConsumer__factory>;
    getContractFactory(
      name: "ZetaNonEthInterface",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaNonEthInterface__factory>;
    getContractFactory(
      name: "AttackerContract",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.AttackerContract__factory>;
    getContractFactory(
      name: "Victim",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.Victim__factory>;
    getContractFactory(
      name: "ERC20Mock",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC20Mock__factory>;
    getContractFactory(
      name: "INonfungiblePositionManager",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.INonfungiblePositionManager__factory>;
    getContractFactory(
      name: "IPoolInitializer",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IPoolInitializer__factory>;
    getContractFactory(
      name: "ZetaInteractorMock",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaInteractorMock__factory>;
    getContractFactory(
      name: "ZetaReceiverMock",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaReceiverMock__factory>;
    getContractFactory(
      name: "ImmutableCreate2Factory",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ImmutableCreate2Factory__factory>;
    getContractFactory(
      name: "Ownable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.Ownable__factory>;
    getContractFactory(
      name: "ConcentratedLiquidityPoolFactory",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ConcentratedLiquidityPoolFactory__factory>;
    getContractFactory(
      name: "IPoolRouter",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IPoolRouter__factory>;
    getContractFactory(
      name: "ZetaInteractor",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaInteractor__factory>;
    getContractFactory(
      name: "ISwapRouterPancake",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ISwapRouterPancake__factory>;
    getContractFactory(
      name: "WETH9",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.WETH9__factory>;
    getContractFactory(
      name: "ZetaTokenConsumerPancakeV3",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaTokenConsumerPancakeV3__factory>;
    getContractFactory(
      name: "ZetaTokenConsumerUniV3Errors",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaTokenConsumerUniV3Errors__factory>;
    getContractFactory(
      name: "WETH9",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.WETH9__factory>;
    getContractFactory(
      name: "ZetaTokenConsumerTrident",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaTokenConsumerTrident__factory>;
    getContractFactory(
      name: "ZetaTokenConsumerTridentErrors",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaTokenConsumerTridentErrors__factory>;
    getContractFactory(
      name: "ZetaTokenConsumerUniV2",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaTokenConsumerUniV2__factory>;
    getContractFactory(
      name: "ZetaTokenConsumerUniV2Errors",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaTokenConsumerUniV2Errors__factory>;
    getContractFactory(
      name: "WETH9",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.WETH9__factory>;
    getContractFactory(
      name: "ZetaTokenConsumerUniV3",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaTokenConsumerUniV3__factory>;
    getContractFactory(
      name: "ZetaTokenConsumerUniV3Errors",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaTokenConsumerUniV3Errors__factory>;
    getContractFactory(
      name: "ZetaTokenConsumerZEVM",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaTokenConsumerZEVM__factory>;
    getContractFactory(
      name: "ZetaTokenConsumerZEVMErrors",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaTokenConsumerZEVMErrors__factory>;
    getContractFactory(
      name: "ZetaEth",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaEth__factory>;
    getContractFactory(
      name: "ZetaNonEth",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaNonEth__factory>;
    getContractFactory(
      name: "ZetaConnectorBase",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaConnectorBase__factory>;
    getContractFactory(
      name: "ZetaConnectorEth",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaConnectorEth__factory>;
    getContractFactory(
      name: "ZetaConnectorNonEth",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaConnectorNonEth__factory>;
    getContractFactory(
      name: "ERC20CustodyNew",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC20CustodyNew__factory>;
    getContractFactory(
      name: "GatewayEVM",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.GatewayEVM__factory>;
    getContractFactory(
      name: "GatewayEVMUpgradeTest",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.GatewayEVMUpgradeTest__factory>;
    getContractFactory(
      name: "IGatewayEVM",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IGatewayEVM__factory>;
    getContractFactory(
      name: "GatewayUpgradeTest",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.GatewayUpgradeTest__factory>;
    getContractFactory(
      name: "IGateway",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IGateway__factory>;
    getContractFactory(
      name: "Receiver",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.Receiver__factory>;
    getContractFactory(
      name: "TestERC20",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.TestERC20__factory>;
    getContractFactory(
      name: "ISystem",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ISystem__factory>;
    getContractFactory(
      name: "IZRC20",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IZRC20__factory>;
    getContractFactory(
      name: "IZRC20Metadata",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IZRC20Metadata__factory>;
    getContractFactory(
      name: "IUniswapV2Router01",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV2Router01__factory>;
    getContractFactory(
      name: "IUniswapV2Router02",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IUniswapV2Router02__factory>;
    getContractFactory(
      name: "IWETH9",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IWETH9__factory>;
    getContractFactory(
      name: "IZRC20",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IZRC20__factory>;
    getContractFactory(
      name: "ZContract",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZContract__factory>;
    getContractFactory(
      name: "SystemContract",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.SystemContract__factory>;
    getContractFactory(
      name: "SystemContractErrors",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.SystemContractErrors__factory>;
    getContractFactory(
      name: "SystemContractErrors",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.SystemContractErrors__factory>;
    getContractFactory(
      name: "SystemContractMock",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.SystemContractMock__factory>;
    getContractFactory(
      name: "WETH9",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.WETH9__factory>;
    getContractFactory(
      name: "ZetaConnectorZEVM",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaConnectorZEVM__factory>;
    getContractFactory(
      name: "ZetaReceiver",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZetaReceiver__factory>;
    getContractFactory(
      name: "ZRC20",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZRC20__factory>;
    getContractFactory(
      name: "ZRC20Errors",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ZRC20Errors__factory>;

    getContractAt(
      name: "OwnableUpgradeable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.OwnableUpgradeable>;
    getContractAt(
      name: "IERC1822ProxiableUpgradeable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IERC1822ProxiableUpgradeable>;
    getContractAt(
      name: "IERC1967Upgradeable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IERC1967Upgradeable>;
    getContractAt(
      name: "IBeaconUpgradeable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IBeaconUpgradeable>;
    getContractAt(
      name: "ERC1967UpgradeUpgradeable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC1967UpgradeUpgradeable>;
    getContractAt(
      name: "Initializable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.Initializable>;
    getContractAt(
      name: "UUPSUpgradeable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.UUPSUpgradeable>;
    getContractAt(
      name: "ContextUpgradeable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ContextUpgradeable>;
    getContractAt(
      name: "Ownable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.Ownable>;
    getContractAt(
      name: "Ownable2Step",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.Ownable2Step>;
    getContractAt(
      name: "Pausable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.Pausable>;
    getContractAt(
      name: "ERC20",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC20>;
    getContractAt(
      name: "IERC20Permit",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IERC20Permit>;
    getContractAt(
      name: "ERC20Burnable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC20Burnable>;
    getContractAt(
      name: "IERC20Metadata",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IERC20Metadata>;
    getContractAt(
      name: "IERC20",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IERC20>;
    getContractAt(
      name: "IERC20",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IERC20>;
    getContractAt(
      name: "IUniswapV2Callee",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV2Callee>;
    getContractAt(
      name: "IUniswapV2ERC20",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV2ERC20>;
    getContractAt(
      name: "IUniswapV2Factory",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV2Factory>;
    getContractAt(
      name: "IUniswapV2Pair",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV2Pair>;
    getContractAt(
      name: "UniswapV2ERC20",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.UniswapV2ERC20>;
    getContractAt(
      name: "UniswapV2Factory",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.UniswapV2Factory>;
    getContractAt(
      name: "UniswapV2Pair",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.UniswapV2Pair>;
    getContractAt(
      name: "IERC20",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IERC20>;
    getContractAt(
      name: "IUniswapV2Router01",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV2Router01>;
    getContractAt(
      name: "IUniswapV2Router02",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV2Router02>;
    getContractAt(
      name: "IWETH",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IWETH>;
    getContractAt(
      name: "UniswapV2Router02",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.UniswapV2Router02>;
    getContractAt(
      name: "IUniswapV3SwapCallback",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV3SwapCallback>;
    getContractAt(
      name: "IUniswapV3Factory",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV3Factory>;
    getContractAt(
      name: "IUniswapV3Pool",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV3Pool>;
    getContractAt(
      name: "IUniswapV3PoolActions",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV3PoolActions>;
    getContractAt(
      name: "IUniswapV3PoolDerivedState",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV3PoolDerivedState>;
    getContractAt(
      name: "IUniswapV3PoolEvents",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV3PoolEvents>;
    getContractAt(
      name: "IUniswapV3PoolImmutables",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV3PoolImmutables>;
    getContractAt(
      name: "IUniswapV3PoolOwnerActions",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV3PoolOwnerActions>;
    getContractAt(
      name: "IUniswapV3PoolState",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV3PoolState>;
    getContractAt(
      name: "IQuoter",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IQuoter>;
    getContractAt(
      name: "ISwapRouter",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ISwapRouter>;
    getContractAt(
      name: "ERC20Custody",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC20Custody>;
    getContractAt(
      name: "ConnectorErrors",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ConnectorErrors>;
    getContractAt(
      name: "ZetaErrors",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaErrors>;
    getContractAt(
      name: "ZetaInteractorErrors",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaInteractorErrors>;
    getContractAt(
      name: "ZetaCommonErrors",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaCommonErrors>;
    getContractAt(
      name: "ZetaConnector",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaConnector>;
    getContractAt(
      name: "ZetaReceiver",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaReceiver>;
    getContractAt(
      name: "ZetaTokenConsumer",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaTokenConsumer>;
    getContractAt(
      name: "ZetaNonEthInterface",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaNonEthInterface>;
    getContractAt(
      name: "AttackerContract",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.AttackerContract>;
    getContractAt(
      name: "Victim",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.Victim>;
    getContractAt(
      name: "ERC20Mock",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC20Mock>;
    getContractAt(
      name: "INonfungiblePositionManager",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.INonfungiblePositionManager>;
    getContractAt(
      name: "IPoolInitializer",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IPoolInitializer>;
    getContractAt(
      name: "ZetaInteractorMock",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaInteractorMock>;
    getContractAt(
      name: "ZetaReceiverMock",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaReceiverMock>;
    getContractAt(
      name: "ImmutableCreate2Factory",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ImmutableCreate2Factory>;
    getContractAt(
      name: "Ownable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.Ownable>;
    getContractAt(
      name: "ConcentratedLiquidityPoolFactory",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ConcentratedLiquidityPoolFactory>;
    getContractAt(
      name: "IPoolRouter",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IPoolRouter>;
    getContractAt(
      name: "ZetaInteractor",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaInteractor>;
    getContractAt(
      name: "ISwapRouterPancake",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ISwapRouterPancake>;
    getContractAt(
      name: "WETH9",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.WETH9>;
    getContractAt(
      name: "ZetaTokenConsumerPancakeV3",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaTokenConsumerPancakeV3>;
    getContractAt(
      name: "ZetaTokenConsumerUniV3Errors",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaTokenConsumerUniV3Errors>;
    getContractAt(
      name: "WETH9",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.WETH9>;
    getContractAt(
      name: "ZetaTokenConsumerTrident",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaTokenConsumerTrident>;
    getContractAt(
      name: "ZetaTokenConsumerTridentErrors",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaTokenConsumerTridentErrors>;
    getContractAt(
      name: "ZetaTokenConsumerUniV2",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaTokenConsumerUniV2>;
    getContractAt(
      name: "ZetaTokenConsumerUniV2Errors",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaTokenConsumerUniV2Errors>;
    getContractAt(
      name: "WETH9",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.WETH9>;
    getContractAt(
      name: "ZetaTokenConsumerUniV3",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaTokenConsumerUniV3>;
    getContractAt(
      name: "ZetaTokenConsumerUniV3Errors",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaTokenConsumerUniV3Errors>;
    getContractAt(
      name: "ZetaTokenConsumerZEVM",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaTokenConsumerZEVM>;
    getContractAt(
      name: "ZetaTokenConsumerZEVMErrors",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaTokenConsumerZEVMErrors>;
    getContractAt(
      name: "ZetaEth",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaEth>;
    getContractAt(
      name: "ZetaNonEth",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaNonEth>;
    getContractAt(
      name: "ZetaConnectorBase",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaConnectorBase>;
    getContractAt(
      name: "ZetaConnectorEth",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaConnectorEth>;
    getContractAt(
      name: "ZetaConnectorNonEth",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaConnectorNonEth>;
    getContractAt(
      name: "ERC20CustodyNew",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC20CustodyNew>;
    getContractAt(
      name: "GatewayEVM",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.GatewayEVM>;
    getContractAt(
      name: "GatewayEVMUpgradeTest",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.GatewayEVMUpgradeTest>;
    getContractAt(
      name: "IGatewayEVM",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IGatewayEVM>;
    getContractAt(
      name: "GatewayUpgradeTest",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.GatewayUpgradeTest>;
    getContractAt(
      name: "IGateway",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IGateway>;
    getContractAt(
      name: "Receiver",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.Receiver>;
    getContractAt(
      name: "TestERC20",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.TestERC20>;
    getContractAt(
      name: "ISystem",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ISystem>;
    getContractAt(
      name: "IZRC20",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IZRC20>;
    getContractAt(
      name: "IZRC20Metadata",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IZRC20Metadata>;
    getContractAt(
      name: "IUniswapV2Router01",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV2Router01>;
    getContractAt(
      name: "IUniswapV2Router02",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IUniswapV2Router02>;
    getContractAt(
      name: "IWETH9",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IWETH9>;
    getContractAt(
      name: "IZRC20",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IZRC20>;
    getContractAt(
      name: "ZContract",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZContract>;
    getContractAt(
      name: "SystemContract",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.SystemContract>;
    getContractAt(
      name: "SystemContractErrors",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.SystemContractErrors>;
    getContractAt(
      name: "SystemContractErrors",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.SystemContractErrors>;
    getContractAt(
      name: "SystemContractMock",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.SystemContractMock>;
    getContractAt(
      name: "WETH9",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.WETH9>;
    getContractAt(
      name: "ZetaConnectorZEVM",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaConnectorZEVM>;
    getContractAt(
      name: "ZetaReceiver",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZetaReceiver>;
    getContractAt(
      name: "ZRC20",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZRC20>;
    getContractAt(
      name: "ZRC20Errors",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ZRC20Errors>;

    // default types
    getContractFactory(
      name: string,
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<ethers.ContractFactory>;
    getContractFactory(
      abi: any[],
      bytecode: ethers.utils.BytesLike,
      signer?: ethers.Signer
    ): Promise<ethers.ContractFactory>;
    getContractAt(
      nameOrAbi: string | any[],
      address: string,
      signer?: ethers.Signer
    ): Promise<ethers.Contract>;
  }
}
