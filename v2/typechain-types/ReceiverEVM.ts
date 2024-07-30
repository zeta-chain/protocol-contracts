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

export interface ReceiverEVMInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "onRevert"
      | "receiveERC20"
      | "receiveERC20Partial"
      | "receiveNoParams"
      | "receiveNonPayable"
      | "receivePayable"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic:
      | "ReceivedERC20"
      | "ReceivedNoParams"
      | "ReceivedNonPayable"
      | "ReceivedPayable"
      | "ReceivedRevert"
  ): EventFragment;

  encodeFunctionData(functionFragment: "onRevert", values: [BytesLike]): string;
  encodeFunctionData(
    functionFragment: "receiveERC20",
    values: [BigNumberish, AddressLike, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "receiveERC20Partial",
    values: [BigNumberish, AddressLike, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "receiveNoParams",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "receiveNonPayable",
    values: [string[], BigNumberish[], boolean]
  ): string;
  encodeFunctionData(
    functionFragment: "receivePayable",
    values: [string, BigNumberish, boolean]
  ): string;

  decodeFunctionResult(functionFragment: "onRevert", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "receiveERC20",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "receiveERC20Partial",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "receiveNoParams",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "receiveNonPayable",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "receivePayable",
    data: BytesLike
  ): Result;
}

export namespace ReceivedERC20Event {
  export type InputTuple = [
    sender: AddressLike,
    amount: BigNumberish,
    token: AddressLike,
    destination: AddressLike
  ];
  export type OutputTuple = [
    sender: string,
    amount: bigint,
    token: string,
    destination: string
  ];
  export interface OutputObject {
    sender: string;
    amount: bigint;
    token: string;
    destination: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ReceivedNoParamsEvent {
  export type InputTuple = [sender: AddressLike];
  export type OutputTuple = [sender: string];
  export interface OutputObject {
    sender: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ReceivedNonPayableEvent {
  export type InputTuple = [
    sender: AddressLike,
    strs: string[],
    nums: BigNumberish[],
    flag: boolean
  ];
  export type OutputTuple = [
    sender: string,
    strs: string[],
    nums: bigint[],
    flag: boolean
  ];
  export interface OutputObject {
    sender: string;
    strs: string[];
    nums: bigint[];
    flag: boolean;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ReceivedPayableEvent {
  export type InputTuple = [
    sender: AddressLike,
    value: BigNumberish,
    str: string,
    num: BigNumberish,
    flag: boolean
  ];
  export type OutputTuple = [
    sender: string,
    value: bigint,
    str: string,
    num: bigint,
    flag: boolean
  ];
  export interface OutputObject {
    sender: string;
    value: bigint;
    str: string;
    num: bigint;
    flag: boolean;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ReceivedRevertEvent {
  export type InputTuple = [sender: AddressLike, data: BytesLike];
  export type OutputTuple = [sender: string, data: string];
  export interface OutputObject {
    sender: string;
    data: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface ReceiverEVM extends BaseContract {
  connect(runner?: ContractRunner | null): ReceiverEVM;
  waitForDeployment(): Promise<this>;

  interface: ReceiverEVMInterface;

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

  onRevert: TypedContractMethod<[data: BytesLike], [void], "nonpayable">;

  receiveERC20: TypedContractMethod<
    [amount: BigNumberish, token: AddressLike, destination: AddressLike],
    [void],
    "nonpayable"
  >;

  receiveERC20Partial: TypedContractMethod<
    [amount: BigNumberish, token: AddressLike, destination: AddressLike],
    [void],
    "nonpayable"
  >;

  receiveNoParams: TypedContractMethod<[], [void], "nonpayable">;

  receiveNonPayable: TypedContractMethod<
    [strs: string[], nums: BigNumberish[], flag: boolean],
    [void],
    "nonpayable"
  >;

  receivePayable: TypedContractMethod<
    [str: string, num: BigNumberish, flag: boolean],
    [void],
    "payable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "onRevert"
  ): TypedContractMethod<[data: BytesLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "receiveERC20"
  ): TypedContractMethod<
    [amount: BigNumberish, token: AddressLike, destination: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "receiveERC20Partial"
  ): TypedContractMethod<
    [amount: BigNumberish, token: AddressLike, destination: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "receiveNoParams"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "receiveNonPayable"
  ): TypedContractMethod<
    [strs: string[], nums: BigNumberish[], flag: boolean],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "receivePayable"
  ): TypedContractMethod<
    [str: string, num: BigNumberish, flag: boolean],
    [void],
    "payable"
  >;

  getEvent(
    key: "ReceivedERC20"
  ): TypedContractEvent<
    ReceivedERC20Event.InputTuple,
    ReceivedERC20Event.OutputTuple,
    ReceivedERC20Event.OutputObject
  >;
  getEvent(
    key: "ReceivedNoParams"
  ): TypedContractEvent<
    ReceivedNoParamsEvent.InputTuple,
    ReceivedNoParamsEvent.OutputTuple,
    ReceivedNoParamsEvent.OutputObject
  >;
  getEvent(
    key: "ReceivedNonPayable"
  ): TypedContractEvent<
    ReceivedNonPayableEvent.InputTuple,
    ReceivedNonPayableEvent.OutputTuple,
    ReceivedNonPayableEvent.OutputObject
  >;
  getEvent(
    key: "ReceivedPayable"
  ): TypedContractEvent<
    ReceivedPayableEvent.InputTuple,
    ReceivedPayableEvent.OutputTuple,
    ReceivedPayableEvent.OutputObject
  >;
  getEvent(
    key: "ReceivedRevert"
  ): TypedContractEvent<
    ReceivedRevertEvent.InputTuple,
    ReceivedRevertEvent.OutputTuple,
    ReceivedRevertEvent.OutputObject
  >;

  filters: {
    "ReceivedERC20(address,uint256,address,address)": TypedContractEvent<
      ReceivedERC20Event.InputTuple,
      ReceivedERC20Event.OutputTuple,
      ReceivedERC20Event.OutputObject
    >;
    ReceivedERC20: TypedContractEvent<
      ReceivedERC20Event.InputTuple,
      ReceivedERC20Event.OutputTuple,
      ReceivedERC20Event.OutputObject
    >;

    "ReceivedNoParams(address)": TypedContractEvent<
      ReceivedNoParamsEvent.InputTuple,
      ReceivedNoParamsEvent.OutputTuple,
      ReceivedNoParamsEvent.OutputObject
    >;
    ReceivedNoParams: TypedContractEvent<
      ReceivedNoParamsEvent.InputTuple,
      ReceivedNoParamsEvent.OutputTuple,
      ReceivedNoParamsEvent.OutputObject
    >;

    "ReceivedNonPayable(address,string[],uint256[],bool)": TypedContractEvent<
      ReceivedNonPayableEvent.InputTuple,
      ReceivedNonPayableEvent.OutputTuple,
      ReceivedNonPayableEvent.OutputObject
    >;
    ReceivedNonPayable: TypedContractEvent<
      ReceivedNonPayableEvent.InputTuple,
      ReceivedNonPayableEvent.OutputTuple,
      ReceivedNonPayableEvent.OutputObject
    >;

    "ReceivedPayable(address,uint256,string,uint256,bool)": TypedContractEvent<
      ReceivedPayableEvent.InputTuple,
      ReceivedPayableEvent.OutputTuple,
      ReceivedPayableEvent.OutputObject
    >;
    ReceivedPayable: TypedContractEvent<
      ReceivedPayableEvent.InputTuple,
      ReceivedPayableEvent.OutputTuple,
      ReceivedPayableEvent.OutputObject
    >;

    "ReceivedRevert(address,bytes)": TypedContractEvent<
      ReceivedRevertEvent.InputTuple,
      ReceivedRevertEvent.OutputTuple,
      ReceivedRevertEvent.OutputObject
    >;
    ReceivedRevert: TypedContractEvent<
      ReceivedRevertEvent.InputTuple,
      ReceivedRevertEvent.OutputTuple,
      ReceivedRevertEvent.OutputObject
    >;
  };
}
