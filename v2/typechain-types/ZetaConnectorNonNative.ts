/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumberish,
  BytesLike,
  FunctionFragment,
  Result,
  Interface,
  EventFragment,
  AddressLike,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedLogDescription,
  TypedListener,
  TypedContractMethod,
} from "./common";

export interface ZetaConnectorNonNativeInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "gateway"
      | "maxSupply"
      | "receiveTokens"
      | "setMaxSupply"
      | "tssAddress"
      | "withdraw"
      | "withdrawAndCall"
      | "withdrawAndRevert"
      | "zetaToken"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic:
      | "MaxSupplyUpdated"
      | "Withdraw"
      | "WithdrawAndCall"
      | "WithdrawAndRevert"
  ): EventFragment;

  encodeFunctionData(functionFragment: "gateway", values?: undefined): string;
  encodeFunctionData(functionFragment: "maxSupply", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "receiveTokens",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "setMaxSupply",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "tssAddress",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "withdraw",
    values: [AddressLike, BigNumberish, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawAndCall",
    values: [AddressLike, BigNumberish, BytesLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawAndRevert",
    values: [AddressLike, BigNumberish, BytesLike, BytesLike]
  ): string;
  encodeFunctionData(functionFragment: "zetaToken", values?: undefined): string;

  decodeFunctionResult(functionFragment: "gateway", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "maxSupply", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "receiveTokens",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setMaxSupply",
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
}

export namespace MaxSupplyUpdatedEvent {
  export type InputTuple = [maxSupply: BigNumberish];
  export type OutputTuple = [maxSupply: bigint];
  export interface OutputObject {
    maxSupply: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace WithdrawEvent {
  export type InputTuple = [to: AddressLike, amount: BigNumberish];
  export type OutputTuple = [to: string, amount: bigint];
  export interface OutputObject {
    to: string;
    amount: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace WithdrawAndCallEvent {
  export type InputTuple = [
    to: AddressLike,
    amount: BigNumberish,
    data: BytesLike
  ];
  export type OutputTuple = [to: string, amount: bigint, data: string];
  export interface OutputObject {
    to: string;
    amount: bigint;
    data: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace WithdrawAndRevertEvent {
  export type InputTuple = [
    to: AddressLike,
    amount: BigNumberish,
    data: BytesLike
  ];
  export type OutputTuple = [to: string, amount: bigint, data: string];
  export interface OutputObject {
    to: string;
    amount: bigint;
    data: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface ZetaConnectorNonNative extends BaseContract {
  connect(runner?: ContractRunner | null): ZetaConnectorNonNative;
  waitForDeployment(): Promise<this>;

  interface: ZetaConnectorNonNativeInterface;

  queryFilter<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;
  queryFilter<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;

  on<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  on<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  once<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  once<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  listeners<TCEvent extends TypedContractEvent>(
    event: TCEvent
  ): Promise<Array<TypedListener<TCEvent>>>;
  listeners(eventName?: string): Promise<Array<Listener>>;
  removeAllListeners<TCEvent extends TypedContractEvent>(
    event?: TCEvent
  ): Promise<this>;

  gateway: TypedContractMethod<[], [string], "view">;

  maxSupply: TypedContractMethod<[], [bigint], "view">;

  receiveTokens: TypedContractMethod<
    [amount: BigNumberish],
    [void],
    "nonpayable"
  >;

  setMaxSupply: TypedContractMethod<
    [_maxSupply: BigNumberish],
    [void],
    "nonpayable"
  >;

  tssAddress: TypedContractMethod<[], [string], "view">;

  withdraw: TypedContractMethod<
    [to: AddressLike, amount: BigNumberish, internalSendHash: BytesLike],
    [void],
    "nonpayable"
  >;

  withdrawAndCall: TypedContractMethod<
    [
      to: AddressLike,
      amount: BigNumberish,
      data: BytesLike,
      internalSendHash: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  withdrawAndRevert: TypedContractMethod<
    [
      to: AddressLike,
      amount: BigNumberish,
      data: BytesLike,
      internalSendHash: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  zetaToken: TypedContractMethod<[], [string], "view">;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "gateway"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "maxSupply"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "receiveTokens"
  ): TypedContractMethod<[amount: BigNumberish], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "setMaxSupply"
  ): TypedContractMethod<[_maxSupply: BigNumberish], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "tssAddress"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "withdraw"
  ): TypedContractMethod<
    [to: AddressLike, amount: BigNumberish, internalSendHash: BytesLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdrawAndCall"
  ): TypedContractMethod<
    [
      to: AddressLike,
      amount: BigNumberish,
      data: BytesLike,
      internalSendHash: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdrawAndRevert"
  ): TypedContractMethod<
    [
      to: AddressLike,
      amount: BigNumberish,
      data: BytesLike,
      internalSendHash: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "zetaToken"
  ): TypedContractMethod<[], [string], "view">;

  getEvent(
    key: "MaxSupplyUpdated"
  ): TypedContractEvent<
    MaxSupplyUpdatedEvent.InputTuple,
    MaxSupplyUpdatedEvent.OutputTuple,
    MaxSupplyUpdatedEvent.OutputObject
  >;
  getEvent(
    key: "Withdraw"
  ): TypedContractEvent<
    WithdrawEvent.InputTuple,
    WithdrawEvent.OutputTuple,
    WithdrawEvent.OutputObject
  >;
  getEvent(
    key: "WithdrawAndCall"
  ): TypedContractEvent<
    WithdrawAndCallEvent.InputTuple,
    WithdrawAndCallEvent.OutputTuple,
    WithdrawAndCallEvent.OutputObject
  >;
  getEvent(
    key: "WithdrawAndRevert"
  ): TypedContractEvent<
    WithdrawAndRevertEvent.InputTuple,
    WithdrawAndRevertEvent.OutputTuple,
    WithdrawAndRevertEvent.OutputObject
  >;

  filters: {
    "MaxSupplyUpdated(uint256)": TypedContractEvent<
      MaxSupplyUpdatedEvent.InputTuple,
      MaxSupplyUpdatedEvent.OutputTuple,
      MaxSupplyUpdatedEvent.OutputObject
    >;
    MaxSupplyUpdated: TypedContractEvent<
      MaxSupplyUpdatedEvent.InputTuple,
      MaxSupplyUpdatedEvent.OutputTuple,
      MaxSupplyUpdatedEvent.OutputObject
    >;

    "Withdraw(address,uint256)": TypedContractEvent<
      WithdrawEvent.InputTuple,
      WithdrawEvent.OutputTuple,
      WithdrawEvent.OutputObject
    >;
    Withdraw: TypedContractEvent<
      WithdrawEvent.InputTuple,
      WithdrawEvent.OutputTuple,
      WithdrawEvent.OutputObject
    >;

    "WithdrawAndCall(address,uint256,bytes)": TypedContractEvent<
      WithdrawAndCallEvent.InputTuple,
      WithdrawAndCallEvent.OutputTuple,
      WithdrawAndCallEvent.OutputObject
    >;
    WithdrawAndCall: TypedContractEvent<
      WithdrawAndCallEvent.InputTuple,
      WithdrawAndCallEvent.OutputTuple,
      WithdrawAndCallEvent.OutputObject
    >;

    "WithdrawAndRevert(address,uint256,bytes)": TypedContractEvent<
      WithdrawAndRevertEvent.InputTuple,
      WithdrawAndRevertEvent.OutputTuple,
      WithdrawAndRevertEvent.OutputObject
    >;
    WithdrawAndRevert: TypedContractEvent<
      WithdrawAndRevertEvent.InputTuple,
      WithdrawAndRevertEvent.OutputTuple,
      WithdrawAndRevertEvent.OutputObject
    >;
  };
}
