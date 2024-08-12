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
} from "../common";

export type RevertOptionsStruct = {
  revertAddress: AddressLike;
  callOnRevert: boolean;
  abortAddress: AddressLike;
  revertMessage: BytesLike;
};

export type RevertOptionsStructOutput = [
  revertAddress: string,
  callOnRevert: boolean,
  abortAddress: string,
  revertMessage: string
] & {
  revertAddress: string;
  callOnRevert: boolean;
  abortAddress: string;
  revertMessage: string;
};

export interface IGatewayEVMInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "call"
      | "deposit(address,uint256,address,(address,bool,address,bytes))"
      | "deposit(address,(address,bool,address,bytes))"
      | "depositAndCall(address,uint256,address,bytes,(address,bool,address,bytes))"
      | "depositAndCall(address,bytes,(address,bool,address,bytes))"
      | "execute"
      | "executeRevert"
      | "executeWithERC20"
      | "revertWithERC20"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic:
      | "Call"
      | "Deposit"
      | "Executed"
      | "ExecutedWithERC20"
      | "Reverted"
      | "RevertedWithERC20"
  ): EventFragment;

  encodeFunctionData(
    functionFragment: "call",
    values: [AddressLike, BytesLike, RevertOptionsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "deposit(address,uint256,address,(address,bool,address,bytes))",
    values: [AddressLike, BigNumberish, AddressLike, RevertOptionsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "deposit(address,(address,bool,address,bytes))",
    values: [AddressLike, RevertOptionsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "depositAndCall(address,uint256,address,bytes,(address,bool,address,bytes))",
    values: [
      AddressLike,
      BigNumberish,
      AddressLike,
      BytesLike,
      RevertOptionsStruct
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "depositAndCall(address,bytes,(address,bool,address,bytes))",
    values: [AddressLike, BytesLike, RevertOptionsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "execute",
    values: [AddressLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "executeRevert",
    values: [AddressLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "executeWithERC20",
    values: [AddressLike, AddressLike, BigNumberish, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "revertWithERC20",
    values: [AddressLike, AddressLike, BigNumberish, BytesLike]
  ): string;

  decodeFunctionResult(functionFragment: "call", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "deposit(address,uint256,address,(address,bool,address,bytes))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "deposit(address,(address,bool,address,bytes))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "depositAndCall(address,uint256,address,bytes,(address,bool,address,bytes))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "depositAndCall(address,bytes,(address,bool,address,bytes))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "execute", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "executeRevert",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "executeWithERC20",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "revertWithERC20",
    data: BytesLike
  ): Result;
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

export interface IGatewayEVM extends BaseContract {
  connect(runner?: ContractRunner | null): IGatewayEVM;
  waitForDeployment(): Promise<this>;

  interface: IGatewayEVMInterface;

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

  call: TypedContractMethod<
    [
      receiver: AddressLike,
      payload: BytesLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  "deposit(address,uint256,address,(address,bool,address,bytes))": TypedContractMethod<
    [
      receiver: AddressLike,
      amount: BigNumberish,
      asset: AddressLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  "deposit(address,(address,bool,address,bytes))": TypedContractMethod<
    [receiver: AddressLike, revertOptions: RevertOptionsStruct],
    [void],
    "payable"
  >;

  "depositAndCall(address,uint256,address,bytes,(address,bool,address,bytes))": TypedContractMethod<
    [
      receiver: AddressLike,
      amount: BigNumberish,
      asset: AddressLike,
      payload: BytesLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  "depositAndCall(address,bytes,(address,bool,address,bytes))": TypedContractMethod<
    [
      receiver: AddressLike,
      payload: BytesLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "payable"
  >;

  execute: TypedContractMethod<
    [destination: AddressLike, data: BytesLike],
    [string],
    "payable"
  >;

  executeRevert: TypedContractMethod<
    [destination: AddressLike, data: BytesLike],
    [void],
    "payable"
  >;

  executeWithERC20: TypedContractMethod<
    [
      token: AddressLike,
      to: AddressLike,
      amount: BigNumberish,
      data: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  revertWithERC20: TypedContractMethod<
    [
      token: AddressLike,
      to: AddressLike,
      amount: BigNumberish,
      data: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "call"
  ): TypedContractMethod<
    [
      receiver: AddressLike,
      payload: BytesLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "deposit(address,uint256,address,(address,bool,address,bytes))"
  ): TypedContractMethod<
    [
      receiver: AddressLike,
      amount: BigNumberish,
      asset: AddressLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "deposit(address,(address,bool,address,bytes))"
  ): TypedContractMethod<
    [receiver: AddressLike, revertOptions: RevertOptionsStruct],
    [void],
    "payable"
  >;
  getFunction(
    nameOrSignature: "depositAndCall(address,uint256,address,bytes,(address,bool,address,bytes))"
  ): TypedContractMethod<
    [
      receiver: AddressLike,
      amount: BigNumberish,
      asset: AddressLike,
      payload: BytesLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "depositAndCall(address,bytes,(address,bool,address,bytes))"
  ): TypedContractMethod<
    [
      receiver: AddressLike,
      payload: BytesLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "payable"
  >;
  getFunction(
    nameOrSignature: "execute"
  ): TypedContractMethod<
    [destination: AddressLike, data: BytesLike],
    [string],
    "payable"
  >;
  getFunction(
    nameOrSignature: "executeRevert"
  ): TypedContractMethod<
    [destination: AddressLike, data: BytesLike],
    [void],
    "payable"
  >;
  getFunction(
    nameOrSignature: "executeWithERC20"
  ): TypedContractMethod<
    [
      token: AddressLike,
      to: AddressLike,
      amount: BigNumberish,
      data: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "revertWithERC20"
  ): TypedContractMethod<
    [
      token: AddressLike,
      to: AddressLike,
      amount: BigNumberish,
      data: BytesLike
    ],
    [void],
    "nonpayable"
  >;

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
