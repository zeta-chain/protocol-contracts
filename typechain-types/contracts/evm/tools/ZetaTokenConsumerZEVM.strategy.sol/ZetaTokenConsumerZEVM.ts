/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
  PayableOverrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type {
  FunctionFragment,
  Result,
  EventFragment,
} from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../../../../common";

export interface ZetaTokenConsumerZEVMInterface extends utils.Interface {
  functions: {
    "WETH9Address()": FunctionFragment;
    "getEthFromZeta(address,uint256,uint256)": FunctionFragment;
    "getTokenFromZeta(address,uint256,address,uint256)": FunctionFragment;
    "getZetaFromEth(address,uint256)": FunctionFragment;
    "getZetaFromToken(address,uint256,address,uint256)": FunctionFragment;
    "hasZetaLiquidity()": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "WETH9Address"
      | "getEthFromZeta"
      | "getTokenFromZeta"
      | "getZetaFromEth"
      | "getZetaFromToken"
      | "hasZetaLiquidity"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "WETH9Address",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getEthFromZeta",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "getTokenFromZeta",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "getZetaFromEth",
    values: [PromiseOrValue<string>, PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "getZetaFromToken",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "hasZetaLiquidity",
    values?: undefined
  ): string;

  decodeFunctionResult(
    functionFragment: "WETH9Address",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getEthFromZeta",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getTokenFromZeta",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getZetaFromEth",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getZetaFromToken",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "hasZetaLiquidity",
    data: BytesLike
  ): Result;

  events: {
    "EthExchangedForZeta(uint256,uint256)": EventFragment;
    "TokenExchangedForZeta(address,uint256,uint256)": EventFragment;
    "ZetaExchangedForEth(uint256,uint256)": EventFragment;
    "ZetaExchangedForToken(address,uint256,uint256)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "EthExchangedForZeta"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "TokenExchangedForZeta"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "ZetaExchangedForEth"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "ZetaExchangedForToken"): EventFragment;
}

export interface EthExchangedForZetaEventObject {
  amountIn: BigNumber;
  amountOut: BigNumber;
}
export type EthExchangedForZetaEvent = TypedEvent<
  [BigNumber, BigNumber],
  EthExchangedForZetaEventObject
>;

export type EthExchangedForZetaEventFilter =
  TypedEventFilter<EthExchangedForZetaEvent>;

export interface TokenExchangedForZetaEventObject {
  token: string;
  amountIn: BigNumber;
  amountOut: BigNumber;
}
export type TokenExchangedForZetaEvent = TypedEvent<
  [string, BigNumber, BigNumber],
  TokenExchangedForZetaEventObject
>;

export type TokenExchangedForZetaEventFilter =
  TypedEventFilter<TokenExchangedForZetaEvent>;

export interface ZetaExchangedForEthEventObject {
  amountIn: BigNumber;
  amountOut: BigNumber;
}
export type ZetaExchangedForEthEvent = TypedEvent<
  [BigNumber, BigNumber],
  ZetaExchangedForEthEventObject
>;

export type ZetaExchangedForEthEventFilter =
  TypedEventFilter<ZetaExchangedForEthEvent>;

export interface ZetaExchangedForTokenEventObject {
  token: string;
  amountIn: BigNumber;
  amountOut: BigNumber;
}
export type ZetaExchangedForTokenEvent = TypedEvent<
  [string, BigNumber, BigNumber],
  ZetaExchangedForTokenEventObject
>;

export type ZetaExchangedForTokenEventFilter =
  TypedEventFilter<ZetaExchangedForTokenEvent>;

export interface ZetaTokenConsumerZEVM extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: ZetaTokenConsumerZEVMInterface;

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>;

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>;
  listeners(eventName?: string): Array<Listener>;
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this;
  removeAllListeners(eventName?: string): this;
  off: OnEvent<this>;
  on: OnEvent<this>;
  once: OnEvent<this>;
  removeListener: OnEvent<this>;

  functions: {
    WETH9Address(overrides?: CallOverrides): Promise<[string]>;

    getEthFromZeta(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      zetaTokenAmount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    getTokenFromZeta(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      outputToken: PromiseOrValue<string>,
      zetaTokenAmount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    getZetaFromEth(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    getZetaFromToken(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      inputToken: PromiseOrValue<string>,
      inputTokenAmount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    hasZetaLiquidity(overrides?: CallOverrides): Promise<[boolean]>;
  };

  WETH9Address(overrides?: CallOverrides): Promise<string>;

  getEthFromZeta(
    destinationAddress: PromiseOrValue<string>,
    minAmountOut: PromiseOrValue<BigNumberish>,
    zetaTokenAmount: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  getTokenFromZeta(
    destinationAddress: PromiseOrValue<string>,
    minAmountOut: PromiseOrValue<BigNumberish>,
    outputToken: PromiseOrValue<string>,
    zetaTokenAmount: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  getZetaFromEth(
    destinationAddress: PromiseOrValue<string>,
    minAmountOut: PromiseOrValue<BigNumberish>,
    overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  getZetaFromToken(
    destinationAddress: PromiseOrValue<string>,
    minAmountOut: PromiseOrValue<BigNumberish>,
    inputToken: PromiseOrValue<string>,
    inputTokenAmount: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  hasZetaLiquidity(overrides?: CallOverrides): Promise<boolean>;

  callStatic: {
    WETH9Address(overrides?: CallOverrides): Promise<string>;

    getEthFromZeta(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      zetaTokenAmount: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getTokenFromZeta(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      outputToken: PromiseOrValue<string>,
      zetaTokenAmount: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getZetaFromEth(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getZetaFromToken(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      inputToken: PromiseOrValue<string>,
      inputTokenAmount: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    hasZetaLiquidity(overrides?: CallOverrides): Promise<boolean>;
  };

  filters: {
    "EthExchangedForZeta(uint256,uint256)"(
      amountIn?: null,
      amountOut?: null
    ): EthExchangedForZetaEventFilter;
    EthExchangedForZeta(
      amountIn?: null,
      amountOut?: null
    ): EthExchangedForZetaEventFilter;

    "TokenExchangedForZeta(address,uint256,uint256)"(
      token?: null,
      amountIn?: null,
      amountOut?: null
    ): TokenExchangedForZetaEventFilter;
    TokenExchangedForZeta(
      token?: null,
      amountIn?: null,
      amountOut?: null
    ): TokenExchangedForZetaEventFilter;

    "ZetaExchangedForEth(uint256,uint256)"(
      amountIn?: null,
      amountOut?: null
    ): ZetaExchangedForEthEventFilter;
    ZetaExchangedForEth(
      amountIn?: null,
      amountOut?: null
    ): ZetaExchangedForEthEventFilter;

    "ZetaExchangedForToken(address,uint256,uint256)"(
      token?: null,
      amountIn?: null,
      amountOut?: null
    ): ZetaExchangedForTokenEventFilter;
    ZetaExchangedForToken(
      token?: null,
      amountIn?: null,
      amountOut?: null
    ): ZetaExchangedForTokenEventFilter;
  };

  estimateGas: {
    WETH9Address(overrides?: CallOverrides): Promise<BigNumber>;

    getEthFromZeta(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      zetaTokenAmount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    getTokenFromZeta(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      outputToken: PromiseOrValue<string>,
      zetaTokenAmount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    getZetaFromEth(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    getZetaFromToken(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      inputToken: PromiseOrValue<string>,
      inputTokenAmount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    hasZetaLiquidity(overrides?: CallOverrides): Promise<BigNumber>;
  };

  populateTransaction: {
    WETH9Address(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    getEthFromZeta(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      zetaTokenAmount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    getTokenFromZeta(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      outputToken: PromiseOrValue<string>,
      zetaTokenAmount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    getZetaFromEth(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    getZetaFromToken(
      destinationAddress: PromiseOrValue<string>,
      minAmountOut: PromiseOrValue<BigNumberish>,
      inputToken: PromiseOrValue<string>,
      inputTokenAmount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    hasZetaLiquidity(overrides?: CallOverrides): Promise<PopulatedTransaction>;
  };
}
