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

export type CallOptionsStruct = {
  gasLimit: BigNumberish;
  isArbitraryCall: boolean;
};

export type CallOptionsStructOutput = [
  gasLimit: bigint,
  isArbitraryCall: boolean
] & { gasLimit: bigint; isArbitraryCall: boolean };

export type RevertOptionsStruct = {
  revertAddress: AddressLike;
  callOnRevert: boolean;
  abortAddress: AddressLike;
  revertMessage: BytesLike;
  onRevertGasLimit: BigNumberish;
};

export type RevertOptionsStructOutput = [
  revertAddress: string,
  callOnRevert: boolean,
  abortAddress: string,
  revertMessage: string,
  onRevertGasLimit: bigint
] & {
  revertAddress: string;
  callOnRevert: boolean;
  abortAddress: string;
  revertMessage: string;
  onRevertGasLimit: bigint;
};

export interface IGatewayZEVMEventsInterface extends Interface {
  getEvent(
    nameOrSignatureOrTopic: "Called" | "Withdrawn" | "WithdrawnAndCalled"
  ): EventFragment;
}

export namespace CalledEvent {
  export type InputTuple = [
    sender: AddressLike,
    zrc20: AddressLike,
    receiver: BytesLike,
    message: BytesLike,
    callOptions: CallOptionsStruct,
    revertOptions: RevertOptionsStruct
  ];
  export type OutputTuple = [
    sender: string,
    zrc20: string,
    receiver: string,
    message: string,
    callOptions: CallOptionsStructOutput,
    revertOptions: RevertOptionsStructOutput
  ];
  export interface OutputObject {
    sender: string;
    zrc20: string;
    receiver: string;
    message: string;
    callOptions: CallOptionsStructOutput;
    revertOptions: RevertOptionsStructOutput;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace WithdrawnEvent {
  export type InputTuple = [
    sender: AddressLike,
    chainId: BigNumberish,
    receiver: BytesLike,
    zrc20: AddressLike,
    value: BigNumberish,
    gasfee: BigNumberish,
    protocolFlatFee: BigNumberish,
    message: BytesLike,
    callOptions: CallOptionsStruct,
    revertOptions: RevertOptionsStruct
  ];
  export type OutputTuple = [
    sender: string,
    chainId: bigint,
    receiver: string,
    zrc20: string,
    value: bigint,
    gasfee: bigint,
    protocolFlatFee: bigint,
    message: string,
    callOptions: CallOptionsStructOutput,
    revertOptions: RevertOptionsStructOutput
  ];
  export interface OutputObject {
    sender: string;
    chainId: bigint;
    receiver: string;
    zrc20: string;
    value: bigint;
    gasfee: bigint;
    protocolFlatFee: bigint;
    message: string;
    callOptions: CallOptionsStructOutput;
    revertOptions: RevertOptionsStructOutput;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace WithdrawnAndCalledEvent {
  export type InputTuple = [
    sender: AddressLike,
    chainId: BigNumberish,
    receiver: BytesLike,
    zrc20: AddressLike,
    value: BigNumberish,
    gasfee: BigNumberish,
    protocolFlatFee: BigNumberish,
    message: BytesLike,
    callOptions: CallOptionsStruct,
    revertOptions: RevertOptionsStruct
  ];
  export type OutputTuple = [
    sender: string,
    chainId: bigint,
    receiver: string,
    zrc20: string,
    value: bigint,
    gasfee: bigint,
    protocolFlatFee: bigint,
    message: string,
    callOptions: CallOptionsStructOutput,
    revertOptions: RevertOptionsStructOutput
  ];
  export interface OutputObject {
    sender: string;
    chainId: bigint;
    receiver: string;
    zrc20: string;
    value: bigint;
    gasfee: bigint;
    protocolFlatFee: bigint;
    message: string;
    callOptions: CallOptionsStructOutput;
    revertOptions: RevertOptionsStructOutput;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface IGatewayZEVMEvents extends BaseContract {
  connect(runner?: ContractRunner | null): IGatewayZEVMEvents;
  waitForDeployment(): Promise<this>;

  interface: IGatewayZEVMEventsInterface;

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
    key: "Called"
  ): TypedContractEvent<
    CalledEvent.InputTuple,
    CalledEvent.OutputTuple,
    CalledEvent.OutputObject
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

  filters: {
    "Called(address,address,bytes,bytes,tuple,tuple)": TypedContractEvent<
      CalledEvent.InputTuple,
      CalledEvent.OutputTuple,
      CalledEvent.OutputObject
    >;
    Called: TypedContractEvent<
      CalledEvent.InputTuple,
      CalledEvent.OutputTuple,
      CalledEvent.OutputObject
    >;

    "Withdrawn(address,uint256,bytes,address,uint256,uint256,uint256,bytes,tuple,tuple)": TypedContractEvent<
      WithdrawnEvent.InputTuple,
      WithdrawnEvent.OutputTuple,
      WithdrawnEvent.OutputObject
    >;
    Withdrawn: TypedContractEvent<
      WithdrawnEvent.InputTuple,
      WithdrawnEvent.OutputTuple,
      WithdrawnEvent.OutputObject
    >;

    "WithdrawnAndCalled(address,uint256,bytes,address,uint256,uint256,uint256,bytes,tuple,tuple)": TypedContractEvent<
      WithdrawnAndCalledEvent.InputTuple,
      WithdrawnAndCalledEvent.OutputTuple,
      WithdrawnAndCalledEvent.OutputObject
    >;
    WithdrawnAndCalled: TypedContractEvent<
      WithdrawnAndCalledEvent.InputTuple,
      WithdrawnAndCalledEvent.OutputTuple,
      WithdrawnAndCalledEvent.OutputObject
    >;
  };
}
