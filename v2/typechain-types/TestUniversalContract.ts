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

export type ZContextStruct = {
  origin: BytesLike;
  sender: AddressLike;
  chainID: BigNumberish;
};

export type ZContextStructOutput = [
  origin: string,
  sender: string,
  chainID: bigint
] & { origin: string; sender: string; chainID: bigint };

export interface TestUniversalContractInterface extends Interface {
  getFunction(
    nameOrSignature: "onCrossChainCall" | "onRevert"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic: "ContextData" | "ContextDataRevert"
  ): EventFragment;

  encodeFunctionData(
    functionFragment: "onCrossChainCall",
    values: [ZContextStruct, AddressLike, BigNumberish, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "onRevert",
    values: [ZContextStruct, AddressLike, BigNumberish, BytesLike]
  ): string;

  decodeFunctionResult(
    functionFragment: "onCrossChainCall",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "onRevert", data: BytesLike): Result;
}

export namespace ContextDataEvent {
  export type InputTuple = [
    origin: BytesLike,
    sender: AddressLike,
    chainID: BigNumberish,
    msgSender: AddressLike,
    message: string
  ];
  export type OutputTuple = [
    origin: string,
    sender: string,
    chainID: bigint,
    msgSender: string,
    message: string
  ];
  export interface OutputObject {
    origin: string;
    sender: string;
    chainID: bigint;
    msgSender: string;
    message: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ContextDataRevertEvent {
  export type InputTuple = [
    origin: BytesLike,
    sender: AddressLike,
    chainID: BigNumberish,
    msgSender: AddressLike,
    message: string
  ];
  export type OutputTuple = [
    origin: string,
    sender: string,
    chainID: bigint,
    msgSender: string,
    message: string
  ];
  export interface OutputObject {
    origin: string;
    sender: string;
    chainID: bigint;
    msgSender: string;
    message: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface TestUniversalContract extends BaseContract {
  connect(runner?: ContractRunner | null): TestUniversalContract;
  waitForDeployment(): Promise<this>;

  interface: TestUniversalContractInterface;

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

  onCrossChainCall: TypedContractMethod<
    [
      context: ZContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  onRevert: TypedContractMethod<
    [
      context: ZContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "onCrossChainCall"
  ): TypedContractMethod<
    [
      context: ZContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onRevert"
  ): TypedContractMethod<
    [
      context: ZContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  getEvent(
    key: "ContextData"
  ): TypedContractEvent<
    ContextDataEvent.InputTuple,
    ContextDataEvent.OutputTuple,
    ContextDataEvent.OutputObject
  >;
  getEvent(
    key: "ContextDataRevert"
  ): TypedContractEvent<
    ContextDataRevertEvent.InputTuple,
    ContextDataRevertEvent.OutputTuple,
    ContextDataRevertEvent.OutputObject
  >;

  filters: {
    "ContextData(bytes,address,uint256,address,string)": TypedContractEvent<
      ContextDataEvent.InputTuple,
      ContextDataEvent.OutputTuple,
      ContextDataEvent.OutputObject
    >;
    ContextData: TypedContractEvent<
      ContextDataEvent.InputTuple,
      ContextDataEvent.OutputTuple,
      ContextDataEvent.OutputObject
    >;

    "ContextDataRevert(bytes,address,uint256,address,string)": TypedContractEvent<
      ContextDataRevertEvent.InputTuple,
      ContextDataRevertEvent.OutputTuple,
      ContextDataRevertEvent.OutputObject
    >;
    ContextDataRevert: TypedContractEvent<
      ContextDataRevertEvent.InputTuple,
      ContextDataRevertEvent.OutputTuple,
      ContextDataRevertEvent.OutputObject
    >;
  };
}
