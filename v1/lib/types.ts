import { ZetaProtocolNetwork } from "./address.tools";

export type ParamSymbol =
  | "BNB.BSC"
  | "BTC.BTC"
  | "DAI.ETH"
  | "ETH.ETH"
  | "gETH"
  | "MATIC.AMOY"
  | "PEPE.ETH"
  | "POL.POLYGON"
  | "sETH.SEPOLIA"
  | "SHIB.ETH"
  | "tBNB"
  | "tBTC"
  | "tMATIC"
  | "USDC.BSC"
  | "USDC.ETH"
  | "USDC.SEPOLIA"
  | "USDC"
  | "USDT.BSC"
  | "USDT.ETH";

export type ParamChainName = ZetaProtocolNetwork;

export type ParamType =
  | "connector"
  | "erc20Custody"
  | "fungibleModule"
  | "pauser"
  | "systemContract"
  | "tss"
  | "tssUpdater"
  | "uniswapV2Factory"
  | "uniswapV2Router02"
  | "uniswapV3Factory"
  | "uniswapV3Router"
  | "weth9"
  | "zetaToken"
  | "zetaTokenConsumerUniV3"
  | "zrc20";
