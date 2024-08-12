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

export type RevertContextStruct = {
  asset: AddressLike;
  amount: BigNumberish;
  revertMessage: BytesLike;
};

export type RevertContextStructOutput = [
  asset: string,
  amount: bigint,
  revertMessage: string
] & { asset: string; amount: bigint; revertMessage: string };

export interface IERC20CustodyEventsInterface extends Interface {
  getEvent(
    nameOrSignatureOrTopic:
      | "Unwhitelisted"
      | "Whitelisted"
      | "Withdrawn"
      | "WithdrawnAndCalled"
      | "WithdrawnAndReverted"
  ): EventFragment;
}

export namespace UnwhitelistedEvent {
  export type InputTuple = [token: AddressLike];
  export type OutputTuple = [token: string];
  export interface OutputObject {
    token: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace WhitelistedEvent {
  export type InputTuple = [token: AddressLike];
  export type OutputTuple = [token: string];
  export interface OutputObject {
    token: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace WithdrawnEvent {
  export type InputTuple = [
    token: AddressLike,
    to: AddressLike,
    amount: BigNumberish
  ];
  export type OutputTuple = [token: string, to: string, amount: bigint];
  export interface OutputObject {
    token: string;
    to: string;
    amount: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace WithdrawnAndCalledEvent {
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

export namespace WithdrawnAndRevertedEvent {
  export type InputTuple = [
    token: AddressLike,
    to: AddressLike,
    amount: BigNumberish,
    data: BytesLike,
    revertContext: RevertContextStruct
  ];
  export type OutputTuple = [
    token: string,
    to: string,
    amount: bigint,
    data: string,
    revertContext: RevertContextStructOutput
  ];
  export interface OutputObject {
    token: string;
    to: string;
    amount: bigint;
    data: string;
    revertContext: RevertContextStructOutput;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface IERC20CustodyEvents extends BaseContract {
  connect(runner?: ContractRunner | null): IERC20CustodyEvents;
  waitForDeployment(): Promise<this>;

  interface: IERC20CustodyEventsInterface;

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
    key: "Unwhitelisted"
  ): TypedContractEvent<
    UnwhitelistedEvent.InputTuple,
    UnwhitelistedEvent.OutputTuple,
    UnwhitelistedEvent.OutputObject
  >;
  getEvent(
    key: "Whitelisted"
  ): TypedContractEvent<
    WhitelistedEvent.InputTuple,
    WhitelistedEvent.OutputTuple,
    WhitelistedEvent.OutputObject
  >;
  getEvent(
    key: "Withdrawn"
  ): TypedContractEvent<
    WithdrawnEvent.InputTuple,
    WithdrawnEvent.OutputTuple,
    WithdrawnEvent.OutputObject
  >;
  getEvent(
    key: "WithdrawnAndCalled"
  ): TypedContractEvent<
    WithdrawnAndCalledEvent.InputTuple,
    WithdrawnAndCalledEvent.OutputTuple,
    WithdrawnAndCalledEvent.OutputObject
  >;
  getEvent(
    key: "WithdrawnAndReverted"
  ): TypedContractEvent<
    WithdrawnAndRevertedEvent.InputTuple,
    WithdrawnAndRevertedEvent.OutputTuple,
    WithdrawnAndRevertedEvent.OutputObject
  >;

  filters: {
    "Unwhitelisted(address)": TypedContractEvent<
      UnwhitelistedEvent.InputTuple,
      UnwhitelistedEvent.OutputTuple,
      UnwhitelistedEvent.OutputObject
    >;
    Unwhitelisted: TypedContractEvent<
      UnwhitelistedEvent.InputTuple,
      UnwhitelistedEvent.OutputTuple,
      UnwhitelistedEvent.OutputObject
    >;

    "Whitelisted(address)": TypedContractEvent<
      WhitelistedEvent.InputTuple,
      WhitelistedEvent.OutputTuple,
      WhitelistedEvent.OutputObject
    >;
    Whitelisted: TypedContractEvent<
      WhitelistedEvent.InputTuple,
      WhitelistedEvent.OutputTuple,
      WhitelistedEvent.OutputObject
    >;

    "Withdrawn(address,address,uint256)": TypedContractEvent<
      WithdrawnEvent.InputTuple,
      WithdrawnEvent.OutputTuple,
      WithdrawnEvent.OutputObject
    >;
    Withdrawn: TypedContractEvent<
      WithdrawnEvent.InputTuple,
      WithdrawnEvent.OutputTuple,
      WithdrawnEvent.OutputObject
    >;

    "WithdrawnAndCalled(address,address,uint256,bytes)": TypedContractEvent<
      WithdrawnAndCalledEvent.InputTuple,
      WithdrawnAndCalledEvent.OutputTuple,
      WithdrawnAndCalledEvent.OutputObject
    >;
    WithdrawnAndCalled: TypedContractEvent<
      WithdrawnAndCalledEvent.InputTuple,
      WithdrawnAndCalledEvent.OutputTuple,
      WithdrawnAndCalledEvent.OutputObject
    >;

    "WithdrawAndRevert(address,address,uint256,bytes,tuple)": TypedContractEvent<
      WithdrawAndRevertEvent.InputTuple,
      WithdrawAndRevertEvent.OutputTuple,
      WithdrawAndRevertEvent.OutputObject
    >;
    WithdrawnAndReverted: TypedContractEvent<
      WithdrawnAndRevertedEvent.InputTuple,
      WithdrawnAndRevertedEvent.OutputTuple,
      WithdrawnAndRevertedEvent.OutputObject
    >;
  };
}
