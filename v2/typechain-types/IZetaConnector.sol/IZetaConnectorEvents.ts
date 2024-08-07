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

export interface IZetaConnectorEventsInterface extends Interface {
  getEvent(
    nameOrSignatureOrTopic: "Withdraw" | "WithdrawAndCall" | "WithdrawAndRevert"
  ): EventFragment;
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

export interface IZetaConnectorEvents extends BaseContract {
  connect(runner?: ContractRunner | null): IZetaConnectorEvents;
  waitForDeployment(): Promise<this>;

  interface: IZetaConnectorEventsInterface;

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
