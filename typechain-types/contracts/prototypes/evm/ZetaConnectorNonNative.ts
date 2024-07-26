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
} from "../../../common";

export interface ZetaConnectorNonNativeInterface extends utils.Interface {
  functions: {
    "gateway()": FunctionFragment;
    "receiveTokens(uint256)": FunctionFragment;
    "tssAddress()": FunctionFragment;
    "withdraw(address,uint256,bytes32)": FunctionFragment;
    "withdrawAndCall(address,uint256,bytes,bytes32)": FunctionFragment;
    "withdrawAndRevert(address,uint256,bytes,bytes32)": FunctionFragment;
    "zetaToken()": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "gateway"
      | "receiveTokens"
      | "tssAddress"
      | "withdraw"
      | "withdrawAndCall"
      | "withdrawAndRevert"
      | "zetaToken"
  ): FunctionFragment;

  encodeFunctionData(functionFragment: "gateway", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "receiveTokens",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "tssAddress",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "withdraw",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawAndCall",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawAndRevert",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(functionFragment: "zetaToken", values?: undefined): string;

  decodeFunctionResult(functionFragment: "gateway", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "receiveTokens",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "tssAddress", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "withdraw", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "withdrawAndCall",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawAndRevert",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "zetaToken", data: BytesLike): Result;

  events: {
    "Withdraw(address,uint256)": EventFragment;
    "WithdrawAndCall(address,uint256,bytes)": EventFragment;
    "WithdrawAndRevert(address,uint256,bytes)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "Withdraw"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "WithdrawAndCall"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "WithdrawAndRevert"): EventFragment;
}

export interface WithdrawEventObject {
  to: string;
  amount: BigNumber;
}
export type WithdrawEvent = TypedEvent<
  [string, BigNumber],
  WithdrawEventObject
>;

export type WithdrawEventFilter = TypedEventFilter<WithdrawEvent>;

export interface WithdrawAndCallEventObject {
  to: string;
  amount: BigNumber;
  data: string;
}
export type WithdrawAndCallEvent = TypedEvent<
  [string, BigNumber, string],
  WithdrawAndCallEventObject
>;

export type WithdrawAndCallEventFilter = TypedEventFilter<WithdrawAndCallEvent>;

export interface WithdrawAndRevertEventObject {
  to: string;
  amount: BigNumber;
  data: string;
}
export type WithdrawAndRevertEvent = TypedEvent<
  [string, BigNumber, string],
  WithdrawAndRevertEventObject
>;

export type WithdrawAndRevertEventFilter =
  TypedEventFilter<WithdrawAndRevertEvent>;

export interface ZetaConnectorNonNative extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: ZetaConnectorNonNativeInterface;

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
    gateway(overrides?: CallOverrides): Promise<[string]>;

    receiveTokens(
      amount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    tssAddress(overrides?: CallOverrides): Promise<[string]>;

    withdraw(
      to: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    withdrawAndCall(
      to: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      data: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    withdrawAndRevert(
      to: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      data: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    zetaToken(overrides?: CallOverrides): Promise<[string]>;
  };

  gateway(overrides?: CallOverrides): Promise<string>;

  receiveTokens(
    amount: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  tssAddress(overrides?: CallOverrides): Promise<string>;

  withdraw(
    to: PromiseOrValue<string>,
    amount: PromiseOrValue<BigNumberish>,
    internalSendHash: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  withdrawAndCall(
    to: PromiseOrValue<string>,
    amount: PromiseOrValue<BigNumberish>,
    data: PromiseOrValue<BytesLike>,
    internalSendHash: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  withdrawAndRevert(
    to: PromiseOrValue<string>,
    amount: PromiseOrValue<BigNumberish>,
    data: PromiseOrValue<BytesLike>,
    internalSendHash: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  zetaToken(overrides?: CallOverrides): Promise<string>;

  callStatic: {
    gateway(overrides?: CallOverrides): Promise<string>;

    receiveTokens(
      amount: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    tssAddress(overrides?: CallOverrides): Promise<string>;

    withdraw(
      to: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    withdrawAndCall(
      to: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      data: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    withdrawAndRevert(
      to: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      data: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    zetaToken(overrides?: CallOverrides): Promise<string>;
  };

  filters: {
    "Withdraw(address,uint256)"(
      to?: PromiseOrValue<string> | null,
      amount?: null
    ): WithdrawEventFilter;
    Withdraw(
      to?: PromiseOrValue<string> | null,
      amount?: null
    ): WithdrawEventFilter;

    "WithdrawAndCall(address,uint256,bytes)"(
      to?: PromiseOrValue<string> | null,
      amount?: null,
      data?: null
    ): WithdrawAndCallEventFilter;
    WithdrawAndCall(
      to?: PromiseOrValue<string> | null,
      amount?: null,
      data?: null
    ): WithdrawAndCallEventFilter;

    "WithdrawAndRevert(address,uint256,bytes)"(
      to?: PromiseOrValue<string> | null,
      amount?: null,
      data?: null
    ): WithdrawAndRevertEventFilter;
    WithdrawAndRevert(
      to?: PromiseOrValue<string> | null,
      amount?: null,
      data?: null
    ): WithdrawAndRevertEventFilter;
  };

  estimateGas: {
    gateway(overrides?: CallOverrides): Promise<BigNumber>;

    receiveTokens(
      amount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    tssAddress(overrides?: CallOverrides): Promise<BigNumber>;

    withdraw(
      to: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    withdrawAndCall(
      to: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      data: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    withdrawAndRevert(
      to: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      data: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    zetaToken(overrides?: CallOverrides): Promise<BigNumber>;
  };

  populateTransaction: {
    gateway(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    receiveTokens(
      amount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    tssAddress(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    withdraw(
      to: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    withdrawAndCall(
      to: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      data: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    withdrawAndRevert(
      to: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      data: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    zetaToken(overrides?: CallOverrides): Promise<PopulatedTransaction>;
  };
}