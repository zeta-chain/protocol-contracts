/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumberish,
  BytesLike,
  FunctionFragment,
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
} from "../common";

export type RevertOptionsStruct = {
  revertAddress: AddressLike;
  callOnRevert: boolean;
  abortAddress: AddressLike;
};

export type RevertOptionsStructOutput = [
  revertAddress: string,
  callOnRevert: boolean,
  abortAddress: string
] & { revertAddress: string; callOnRevert: boolean; abortAddress: string };

export interface IGatewayEVMEventsInterface extends Interface {
  getEvent(
    nameOrSignatureOrTopic:
      | "Call"
      | "Deposit"
      | "Executed"
      | "ExecutedWithERC20"
      | "Reverted"
      | "RevertedWithERC20"
  ): EventFragment;
}

export namespace CallEvent {
  export type InputTuple = [
    sender: AddressLike,
    receiver: AddressLike,
    payload: BytesLike,
    revertOptions: RevertOptionsStruct
  ];
  export type OutputTuple = [
    sender: string,
    receiver: string,
    payload: string,
    revertOptions: RevertOptionsStructOutput
  ];
  export interface OutputObject {
    sender: string;
    receiver: string;
    payload: string;
    revertOptions: RevertOptionsStructOutput;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace DepositEvent {
  export type InputTuple = [
    sender: AddressLike,
    receiver: AddressLike,
    amount: BigNumberish,
    asset: AddressLike,
    payload: BytesLike,
    revertOptions: RevertOptionsStruct
  ];
  export type OutputTuple = [
    sender: string,
    receiver: string,
    amount: bigint,
    asset: string,
    payload: string,
    revertOptions: RevertOptionsStructOutput
  ];
  export interface OutputObject {
    sender: string;
    receiver: string;
    amount: bigint;
    asset: string;
    payload: string;
    revertOptions: RevertOptionsStructOutput;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ExecutedEvent {
  export type InputTuple = [
    destination: AddressLike,
    value: BigNumberish,
    data: BytesLike
  ];
  export type OutputTuple = [destination: string, value: bigint, data: string];
  export interface OutputObject {
    destination: string;
    value: bigint;
    data: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ExecutedWithERC20Event {
  export type InputTuple = [
    token: AddressLike,
    to: AddressLike,
    amount: BigNumberish,
    data: BytesLike
  ];
  export type OutputTuple = [
    token: string,
    to: string,
    amount: bigint,
    data: string
  ];
  export interface OutputObject {
    token: string;
    to: string;
    amount: bigint;
    data: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace RevertedEvent {
  export type InputTuple = [
    destination: AddressLike,
    value: BigNumberish,
    data: BytesLike
  ];
  export type OutputTuple = [destination: string, value: bigint, data: string];
  export interface OutputObject {
    destination: string;
    value: bigint;
    data: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace RevertedWithERC20Event {
  export type InputTuple = [
    token: AddressLike,
    to: AddressLike,
    amount: BigNumberish,
    data: BytesLike
  ];
  export type OutputTuple = [
    token: string,
    to: string,
    amount: bigint,
    data: string
  ];
  export interface OutputObject {
    token: string;
    to: string;
    amount: bigint;
    data: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface IGatewayEVMEvents extends BaseContract {
  connect(runner?: ContractRunner | null): IGatewayEVMEvents;
  waitForDeployment(): Promise<this>;

  interface: IGatewayEVMEventsInterface;

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

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getEvent(
    key: "Call"
  ): TypedContractEvent<
    CallEvent.InputTuple,
    CallEvent.OutputTuple,
    CallEvent.OutputObject
  >;
  getEvent(
    key: "Deposit"
  ): TypedContractEvent<
    DepositEvent.InputTuple,
    DepositEvent.OutputTuple,
    DepositEvent.OutputObject
  >;
  getEvent(
    key: "Executed"
  ): TypedContractEvent<
    ExecutedEvent.InputTuple,
    ExecutedEvent.OutputTuple,
    ExecutedEvent.OutputObject
  >;
  getEvent(
    key: "ExecutedWithERC20"
  ): TypedContractEvent<
    ExecutedWithERC20Event.InputTuple,
    ExecutedWithERC20Event.OutputTuple,
    ExecutedWithERC20Event.OutputObject
  >;
  getEvent(
    key: "Reverted"
  ): TypedContractEvent<
    RevertedEvent.InputTuple,
    RevertedEvent.OutputTuple,
    RevertedEvent.OutputObject
  >;
  getEvent(
    key: "RevertedWithERC20"
  ): TypedContractEvent<
    RevertedWithERC20Event.InputTuple,
    RevertedWithERC20Event.OutputTuple,
    RevertedWithERC20Event.OutputObject
  >;

  filters: {
    "Call(address,address,bytes,tuple)": TypedContractEvent<
      CallEvent.InputTuple,
      CallEvent.OutputTuple,
      CallEvent.OutputObject
    >;
    Call: TypedContractEvent<
      CallEvent.InputTuple,
      CallEvent.OutputTuple,
      CallEvent.OutputObject
    >;

    "Deposit(address,address,uint256,address,bytes,tuple)": TypedContractEvent<
      DepositEvent.InputTuple,
      DepositEvent.OutputTuple,
      DepositEvent.OutputObject
    >;
    Deposit: TypedContractEvent<
      DepositEvent.InputTuple,
      DepositEvent.OutputTuple,
      DepositEvent.OutputObject
    >;

    "Executed(address,uint256,bytes)": TypedContractEvent<
      ExecutedEvent.InputTuple,
      ExecutedEvent.OutputTuple,
      ExecutedEvent.OutputObject
    >;
    Executed: TypedContractEvent<
      ExecutedEvent.InputTuple,
      ExecutedEvent.OutputTuple,
      ExecutedEvent.OutputObject
    >;

    "ExecutedWithERC20(address,address,uint256,bytes)": TypedContractEvent<
      ExecutedWithERC20Event.InputTuple,
      ExecutedWithERC20Event.OutputTuple,
      ExecutedWithERC20Event.OutputObject
    >;
    ExecutedWithERC20: TypedContractEvent<
      ExecutedWithERC20Event.InputTuple,
      ExecutedWithERC20Event.OutputTuple,
      ExecutedWithERC20Event.OutputObject
    >;

    "Reverted(address,uint256,bytes)": TypedContractEvent<
      RevertedEvent.InputTuple,
      RevertedEvent.OutputTuple,
      RevertedEvent.OutputObject
    >;
    Reverted: TypedContractEvent<
      RevertedEvent.InputTuple,
      RevertedEvent.OutputTuple,
      RevertedEvent.OutputObject
    >;

    "RevertedWithERC20(address,address,uint256,bytes)": TypedContractEvent<
      RevertedWithERC20Event.InputTuple,
      RevertedWithERC20Event.OutputTuple,
      RevertedWithERC20Event.OutputObject
    >;
    RevertedWithERC20: TypedContractEvent<
      RevertedWithERC20Event.InputTuple,
      RevertedWithERC20Event.OutputTuple,
      RevertedWithERC20Event.OutputObject
    >;
  };
}
