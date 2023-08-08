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

export declare namespace ZetaInterfaces {
  export type SendInputStruct = {
    destinationChainId: PromiseOrValue<BigNumberish>;
    destinationAddress: PromiseOrValue<BytesLike>;
    destinationGasLimit: PromiseOrValue<BigNumberish>;
    message: PromiseOrValue<BytesLike>;
    zetaValueAndGas: PromiseOrValue<BigNumberish>;
    zetaParams: PromiseOrValue<BytesLike>;
  };

  export type SendInputStructOutput = [
    BigNumber,
    string,
    BigNumber,
    string,
    BigNumber,
    string
  ] & {
    destinationChainId: BigNumber;
    destinationAddress: string;
    destinationGasLimit: BigNumber;
    message: string;
    zetaValueAndGas: BigNumber;
    zetaParams: string;
  };
}

export interface ZetaConnectorEthInterface extends utils.Interface {
  functions: {
    "getLockedAmount()": FunctionFragment;
    "onReceive(bytes,uint256,address,uint256,bytes,bytes32)": FunctionFragment;
    "onRevert(address,uint256,bytes,uint256,uint256,bytes,bytes32)": FunctionFragment;
    "pause()": FunctionFragment;
    "paused()": FunctionFragment;
    "pauserAddress()": FunctionFragment;
    "renounceTssAddressUpdater()": FunctionFragment;
    "send((uint256,bytes,uint256,bytes,uint256,bytes))": FunctionFragment;
    "tssAddress()": FunctionFragment;
    "tssAddressUpdater()": FunctionFragment;
    "unpause()": FunctionFragment;
    "updatePauserAddress(address)": FunctionFragment;
    "updateTssAddress(address)": FunctionFragment;
    "zetaToken()": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "getLockedAmount"
      | "onReceive"
      | "onRevert"
      | "pause"
      | "paused"
      | "pauserAddress"
      | "renounceTssAddressUpdater"
      | "send"
      | "tssAddress"
      | "tssAddressUpdater"
      | "unpause"
      | "updatePauserAddress"
      | "updateTssAddress"
      | "zetaToken"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "getLockedAmount",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "onReceive",
    values: [
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "onRevert",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(functionFragment: "pause", values?: undefined): string;
  encodeFunctionData(functionFragment: "paused", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "pauserAddress",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "renounceTssAddressUpdater",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "send",
    values: [ZetaInterfaces.SendInputStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "tssAddress",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "tssAddressUpdater",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "unpause", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "updatePauserAddress",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "updateTssAddress",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(functionFragment: "zetaToken", values?: undefined): string;

  decodeFunctionResult(
    functionFragment: "getLockedAmount",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "onReceive", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "onRevert", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "pause", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "paused", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "pauserAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "renounceTssAddressUpdater",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "send", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "tssAddress", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "tssAddressUpdater",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "unpause", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "updatePauserAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "updateTssAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "zetaToken", data: BytesLike): Result;

  events: {
    "Paused(address)": EventFragment;
    "PauserAddressUpdated(address,address)": EventFragment;
    "TSSAddressUpdated(address,address)": EventFragment;
    "Unpaused(address)": EventFragment;
    "ZetaReceived(bytes,uint256,address,uint256,bytes,bytes32)": EventFragment;
    "ZetaReverted(address,uint256,uint256,bytes,uint256,bytes,bytes32)": EventFragment;
    "ZetaSent(address,address,uint256,bytes,uint256,uint256,bytes,bytes)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "Paused"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "PauserAddressUpdated"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "TSSAddressUpdated"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Unpaused"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "ZetaReceived"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "ZetaReverted"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "ZetaSent"): EventFragment;
}

export interface PausedEventObject {
  account: string;
}
export type PausedEvent = TypedEvent<[string], PausedEventObject>;

export type PausedEventFilter = TypedEventFilter<PausedEvent>;

export interface PauserAddressUpdatedEventObject {
  updaterAddress: string;
  newTssAddress: string;
}
export type PauserAddressUpdatedEvent = TypedEvent<
  [string, string],
  PauserAddressUpdatedEventObject
>;

export type PauserAddressUpdatedEventFilter =
  TypedEventFilter<PauserAddressUpdatedEvent>;

export interface TSSAddressUpdatedEventObject {
  zetaTxSenderAddress: string;
  newTssAddress: string;
}
export type TSSAddressUpdatedEvent = TypedEvent<
  [string, string],
  TSSAddressUpdatedEventObject
>;

export type TSSAddressUpdatedEventFilter =
  TypedEventFilter<TSSAddressUpdatedEvent>;

export interface UnpausedEventObject {
  account: string;
}
export type UnpausedEvent = TypedEvent<[string], UnpausedEventObject>;

export type UnpausedEventFilter = TypedEventFilter<UnpausedEvent>;

export interface ZetaReceivedEventObject {
  zetaTxSenderAddress: string;
  sourceChainId: BigNumber;
  destinationAddress: string;
  zetaValue: BigNumber;
  message: string;
  internalSendHash: string;
}
export type ZetaReceivedEvent = TypedEvent<
  [string, BigNumber, string, BigNumber, string, string],
  ZetaReceivedEventObject
>;

export type ZetaReceivedEventFilter = TypedEventFilter<ZetaReceivedEvent>;

export interface ZetaRevertedEventObject {
  zetaTxSenderAddress: string;
  sourceChainId: BigNumber;
  destinationChainId: BigNumber;
  destinationAddress: string;
  remainingZetaValue: BigNumber;
  message: string;
  internalSendHash: string;
}
export type ZetaRevertedEvent = TypedEvent<
  [string, BigNumber, BigNumber, string, BigNumber, string, string],
  ZetaRevertedEventObject
>;

export type ZetaRevertedEventFilter = TypedEventFilter<ZetaRevertedEvent>;

export interface ZetaSentEventObject {
  sourceTxOriginAddress: string;
  zetaTxSenderAddress: string;
  destinationChainId: BigNumber;
  destinationAddress: string;
  zetaValueAndGas: BigNumber;
  destinationGasLimit: BigNumber;
  message: string;
  zetaParams: string;
}
export type ZetaSentEvent = TypedEvent<
  [string, string, BigNumber, string, BigNumber, BigNumber, string, string],
  ZetaSentEventObject
>;

export type ZetaSentEventFilter = TypedEventFilter<ZetaSentEvent>;

export interface ZetaConnectorEth extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: ZetaConnectorEthInterface;

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
    getLockedAmount(overrides?: CallOverrides): Promise<[BigNumber]>;

    onReceive(
      zetaTxSenderAddress: PromiseOrValue<BytesLike>,
      sourceChainId: PromiseOrValue<BigNumberish>,
      destinationAddress: PromiseOrValue<string>,
      zetaValue: PromiseOrValue<BigNumberish>,
      message: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    onRevert(
      zetaTxSenderAddress: PromiseOrValue<string>,
      sourceChainId: PromiseOrValue<BigNumberish>,
      destinationAddress: PromiseOrValue<BytesLike>,
      destinationChainId: PromiseOrValue<BigNumberish>,
      remainingZetaValue: PromiseOrValue<BigNumberish>,
      message: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    pause(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    paused(overrides?: CallOverrides): Promise<[boolean]>;

    pauserAddress(overrides?: CallOverrides): Promise<[string]>;

    renounceTssAddressUpdater(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    send(
      input: ZetaInterfaces.SendInputStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    tssAddress(overrides?: CallOverrides): Promise<[string]>;

    tssAddressUpdater(overrides?: CallOverrides): Promise<[string]>;

    unpause(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    updatePauserAddress(
      pauserAddress_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    updateTssAddress(
      tssAddress_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    zetaToken(overrides?: CallOverrides): Promise<[string]>;
  };

  getLockedAmount(overrides?: CallOverrides): Promise<BigNumber>;

  onReceive(
    zetaTxSenderAddress: PromiseOrValue<BytesLike>,
    sourceChainId: PromiseOrValue<BigNumberish>,
    destinationAddress: PromiseOrValue<string>,
    zetaValue: PromiseOrValue<BigNumberish>,
    message: PromiseOrValue<BytesLike>,
    internalSendHash: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  onRevert(
    zetaTxSenderAddress: PromiseOrValue<string>,
    sourceChainId: PromiseOrValue<BigNumberish>,
    destinationAddress: PromiseOrValue<BytesLike>,
    destinationChainId: PromiseOrValue<BigNumberish>,
    remainingZetaValue: PromiseOrValue<BigNumberish>,
    message: PromiseOrValue<BytesLike>,
    internalSendHash: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  pause(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  paused(overrides?: CallOverrides): Promise<boolean>;

  pauserAddress(overrides?: CallOverrides): Promise<string>;

  renounceTssAddressUpdater(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  send(
    input: ZetaInterfaces.SendInputStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  tssAddress(overrides?: CallOverrides): Promise<string>;

  tssAddressUpdater(overrides?: CallOverrides): Promise<string>;

  unpause(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  updatePauserAddress(
    pauserAddress_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  updateTssAddress(
    tssAddress_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  zetaToken(overrides?: CallOverrides): Promise<string>;

  callStatic: {
    getLockedAmount(overrides?: CallOverrides): Promise<BigNumber>;

    onReceive(
      zetaTxSenderAddress: PromiseOrValue<BytesLike>,
      sourceChainId: PromiseOrValue<BigNumberish>,
      destinationAddress: PromiseOrValue<string>,
      zetaValue: PromiseOrValue<BigNumberish>,
      message: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    onRevert(
      zetaTxSenderAddress: PromiseOrValue<string>,
      sourceChainId: PromiseOrValue<BigNumberish>,
      destinationAddress: PromiseOrValue<BytesLike>,
      destinationChainId: PromiseOrValue<BigNumberish>,
      remainingZetaValue: PromiseOrValue<BigNumberish>,
      message: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    pause(overrides?: CallOverrides): Promise<void>;

    paused(overrides?: CallOverrides): Promise<boolean>;

    pauserAddress(overrides?: CallOverrides): Promise<string>;

    renounceTssAddressUpdater(overrides?: CallOverrides): Promise<void>;

    send(
      input: ZetaInterfaces.SendInputStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    tssAddress(overrides?: CallOverrides): Promise<string>;

    tssAddressUpdater(overrides?: CallOverrides): Promise<string>;

    unpause(overrides?: CallOverrides): Promise<void>;

    updatePauserAddress(
      pauserAddress_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    updateTssAddress(
      tssAddress_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    zetaToken(overrides?: CallOverrides): Promise<string>;
  };

  filters: {
    "Paused(address)"(account?: null): PausedEventFilter;
    Paused(account?: null): PausedEventFilter;

    "PauserAddressUpdated(address,address)"(
      updaterAddress?: null,
      newTssAddress?: null
    ): PauserAddressUpdatedEventFilter;
    PauserAddressUpdated(
      updaterAddress?: null,
      newTssAddress?: null
    ): PauserAddressUpdatedEventFilter;

    "TSSAddressUpdated(address,address)"(
      zetaTxSenderAddress?: null,
      newTssAddress?: null
    ): TSSAddressUpdatedEventFilter;
    TSSAddressUpdated(
      zetaTxSenderAddress?: null,
      newTssAddress?: null
    ): TSSAddressUpdatedEventFilter;

    "Unpaused(address)"(account?: null): UnpausedEventFilter;
    Unpaused(account?: null): UnpausedEventFilter;

    "ZetaReceived(bytes,uint256,address,uint256,bytes,bytes32)"(
      zetaTxSenderAddress?: null,
      sourceChainId?: PromiseOrValue<BigNumberish> | null,
      destinationAddress?: PromiseOrValue<string> | null,
      zetaValue?: null,
      message?: null,
      internalSendHash?: PromiseOrValue<BytesLike> | null
    ): ZetaReceivedEventFilter;
    ZetaReceived(
      zetaTxSenderAddress?: null,
      sourceChainId?: PromiseOrValue<BigNumberish> | null,
      destinationAddress?: PromiseOrValue<string> | null,
      zetaValue?: null,
      message?: null,
      internalSendHash?: PromiseOrValue<BytesLike> | null
    ): ZetaReceivedEventFilter;

    "ZetaReverted(address,uint256,uint256,bytes,uint256,bytes,bytes32)"(
      zetaTxSenderAddress?: null,
      sourceChainId?: null,
      destinationChainId?: PromiseOrValue<BigNumberish> | null,
      destinationAddress?: null,
      remainingZetaValue?: null,
      message?: null,
      internalSendHash?: PromiseOrValue<BytesLike> | null
    ): ZetaRevertedEventFilter;
    ZetaReverted(
      zetaTxSenderAddress?: null,
      sourceChainId?: null,
      destinationChainId?: PromiseOrValue<BigNumberish> | null,
      destinationAddress?: null,
      remainingZetaValue?: null,
      message?: null,
      internalSendHash?: PromiseOrValue<BytesLike> | null
    ): ZetaRevertedEventFilter;

    "ZetaSent(address,address,uint256,bytes,uint256,uint256,bytes,bytes)"(
      sourceTxOriginAddress?: null,
      zetaTxSenderAddress?: PromiseOrValue<string> | null,
      destinationChainId?: PromiseOrValue<BigNumberish> | null,
      destinationAddress?: null,
      zetaValueAndGas?: null,
      destinationGasLimit?: null,
      message?: null,
      zetaParams?: null
    ): ZetaSentEventFilter;
    ZetaSent(
      sourceTxOriginAddress?: null,
      zetaTxSenderAddress?: PromiseOrValue<string> | null,
      destinationChainId?: PromiseOrValue<BigNumberish> | null,
      destinationAddress?: null,
      zetaValueAndGas?: null,
      destinationGasLimit?: null,
      message?: null,
      zetaParams?: null
    ): ZetaSentEventFilter;
  };

  estimateGas: {
    getLockedAmount(overrides?: CallOverrides): Promise<BigNumber>;

    onReceive(
      zetaTxSenderAddress: PromiseOrValue<BytesLike>,
      sourceChainId: PromiseOrValue<BigNumberish>,
      destinationAddress: PromiseOrValue<string>,
      zetaValue: PromiseOrValue<BigNumberish>,
      message: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    onRevert(
      zetaTxSenderAddress: PromiseOrValue<string>,
      sourceChainId: PromiseOrValue<BigNumberish>,
      destinationAddress: PromiseOrValue<BytesLike>,
      destinationChainId: PromiseOrValue<BigNumberish>,
      remainingZetaValue: PromiseOrValue<BigNumberish>,
      message: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    pause(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    paused(overrides?: CallOverrides): Promise<BigNumber>;

    pauserAddress(overrides?: CallOverrides): Promise<BigNumber>;

    renounceTssAddressUpdater(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    send(
      input: ZetaInterfaces.SendInputStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    tssAddress(overrides?: CallOverrides): Promise<BigNumber>;

    tssAddressUpdater(overrides?: CallOverrides): Promise<BigNumber>;

    unpause(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    updatePauserAddress(
      pauserAddress_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    updateTssAddress(
      tssAddress_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    zetaToken(overrides?: CallOverrides): Promise<BigNumber>;
  };

  populateTransaction: {
    getLockedAmount(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    onReceive(
      zetaTxSenderAddress: PromiseOrValue<BytesLike>,
      sourceChainId: PromiseOrValue<BigNumberish>,
      destinationAddress: PromiseOrValue<string>,
      zetaValue: PromiseOrValue<BigNumberish>,
      message: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    onRevert(
      zetaTxSenderAddress: PromiseOrValue<string>,
      sourceChainId: PromiseOrValue<BigNumberish>,
      destinationAddress: PromiseOrValue<BytesLike>,
      destinationChainId: PromiseOrValue<BigNumberish>,
      remainingZetaValue: PromiseOrValue<BigNumberish>,
      message: PromiseOrValue<BytesLike>,
      internalSendHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    pause(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    paused(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    pauserAddress(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    renounceTssAddressUpdater(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    send(
      input: ZetaInterfaces.SendInputStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    tssAddress(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    tssAddressUpdater(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    unpause(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    updatePauserAddress(
      pauserAddress_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    updateTssAddress(
      tssAddress_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    zetaToken(overrides?: CallOverrides): Promise<PopulatedTransaction>;
  };
}
