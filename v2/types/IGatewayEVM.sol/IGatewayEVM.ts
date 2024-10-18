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

export type MessageContextStruct = { sender: AddressLike };

export type MessageContextStructOutput = [sender: string] & { sender: string };

export type RevertContextStruct = {
  sender: AddressLike;
  asset: AddressLike;
  amount: BigNumberish;
  revertMessage: BytesLike;
};

export type RevertContextStructOutput = [
  sender: string,
  asset: string,
  amount: bigint,
  revertMessage: string
] & { sender: string; asset: string; amount: bigint; revertMessage: string };

export interface IGatewayEVMInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "call"
      | "deposit(address,uint256,address,(address,bool,address,bytes,uint256))"
      | "deposit(address,(address,bool,address,bytes,uint256))"
      | "depositAndCall(address,bytes,(address,bool,address,bytes,uint256))"
      | "depositAndCall(address,uint256,address,bytes,(address,bool,address,bytes,uint256))"
      | "execute"
      | "executeRevert"
      | "executeWithERC20"
      | "revertWithERC20"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic:
      | "Called"
      | "Deposited"
      | "DepositedAndCalled"
      | "Executed"
      | "ExecutedWithERC20"
      | "Reverted"
      | "UpdatedGatewayTSSAddress"
  ): EventFragment;

  encodeFunctionData(
    functionFragment: "call",
    values: [AddressLike, BytesLike, RevertOptionsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "deposit(address,uint256,address,(address,bool,address,bytes,uint256))",
    values: [AddressLike, BigNumberish, AddressLike, RevertOptionsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "deposit(address,(address,bool,address,bytes,uint256))",
    values: [AddressLike, RevertOptionsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "depositAndCall(address,bytes,(address,bool,address,bytes,uint256))",
    values: [AddressLike, BytesLike, RevertOptionsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "depositAndCall(address,uint256,address,bytes,(address,bool,address,bytes,uint256))",
    values: [
      AddressLike,
      BigNumberish,
      AddressLike,
      BytesLike,
      RevertOptionsStruct
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "execute",
    values: [MessageContextStruct, AddressLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "executeRevert",
    values: [AddressLike, BytesLike, RevertContextStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "executeWithERC20",
    values: [
      MessageContextStruct,
      AddressLike,
      AddressLike,
      BigNumberish,
      BytesLike
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "revertWithERC20",
    values: [
      AddressLike,
      AddressLike,
      BigNumberish,
      BytesLike,
      RevertContextStruct
    ]
  ): string;

  decodeFunctionResult(functionFragment: "call", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "deposit(address,uint256,address,(address,bool,address,bytes,uint256))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "deposit(address,(address,bool,address,bytes,uint256))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "depositAndCall(address,bytes,(address,bool,address,bytes,uint256))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "depositAndCall(address,uint256,address,bytes,(address,bool,address,bytes,uint256))",
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

export namespace CalledEvent {
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

export namespace DepositedEvent {
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

export namespace DepositedAndCalledEvent {
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
    to: AddressLike,
    token: AddressLike,
    amount: BigNumberish,
    data: BytesLike,
    revertContext: RevertContextStruct
  ];
  export type OutputTuple = [
    to: string,
    token: string,
    amount: bigint,
    data: string,
    revertContext: RevertContextStructOutput
  ];
  export interface OutputObject {
    to: string;
    token: string;
    amount: bigint;
    data: string;
    revertContext: RevertContextStructOutput;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace UpdatedGatewayTSSAddressEvent {
  export type InputTuple = [
    oldTSSAddress: AddressLike,
    newTSSAddress: AddressLike
  ];
  export type OutputTuple = [oldTSSAddress: string, newTSSAddress: string];
  export interface OutputObject {
    oldTSSAddress: string;
    newTSSAddress: string;
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

  "deposit(address,uint256,address,(address,bool,address,bytes,uint256))": TypedContractMethod<
    [
      receiver: AddressLike,
      amount: BigNumberish,
      asset: AddressLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  "deposit(address,(address,bool,address,bytes,uint256))": TypedContractMethod<
    [receiver: AddressLike, revertOptions: RevertOptionsStruct],
    [void],
    "payable"
  >;

  "depositAndCall(address,bytes,(address,bool,address,bytes,uint256))": TypedContractMethod<
    [
      receiver: AddressLike,
      payload: BytesLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "payable"
  >;

  "depositAndCall(address,uint256,address,bytes,(address,bool,address,bytes,uint256))": TypedContractMethod<
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

  execute: TypedContractMethod<
    [
      messageContext: MessageContextStruct,
      destination: AddressLike,
      data: BytesLike
    ],
    [string],
    "payable"
  >;

  executeRevert: TypedContractMethod<
    [
      destination: AddressLike,
      data: BytesLike,
      revertContext: RevertContextStruct
    ],
    [void],
    "payable"
  >;

  executeWithERC20: TypedContractMethod<
    [
      messageContext: MessageContextStruct,
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
      data: BytesLike,
      revertContext: RevertContextStruct
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
    nameOrSignature: "deposit(address,uint256,address,(address,bool,address,bytes,uint256))"
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
    nameOrSignature: "deposit(address,(address,bool,address,bytes,uint256))"
  ): TypedContractMethod<
    [receiver: AddressLike, revertOptions: RevertOptionsStruct],
    [void],
    "payable"
  >;
  getFunction(
    nameOrSignature: "depositAndCall(address,bytes,(address,bool,address,bytes,uint256))"
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
    nameOrSignature: "depositAndCall(address,uint256,address,bytes,(address,bool,address,bytes,uint256))"
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
    nameOrSignature: "execute"
  ): TypedContractMethod<
    [
      messageContext: MessageContextStruct,
      destination: AddressLike,
      data: BytesLike
    ],
    [string],
    "payable"
  >;
  getFunction(
    nameOrSignature: "executeRevert"
  ): TypedContractMethod<
    [
      destination: AddressLike,
      data: BytesLike,
      revertContext: RevertContextStruct
    ],
    [void],
    "payable"
  >;
  getFunction(
    nameOrSignature: "executeWithERC20"
  ): TypedContractMethod<
    [
      messageContext: MessageContextStruct,
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
      data: BytesLike,
      revertContext: RevertContextStruct
    ],
    [void],
    "nonpayable"
  >;

  getEvent(
    key: "Called"
  ): TypedContractEvent<
    CalledEvent.InputTuple,
    CalledEvent.OutputTuple,
    CalledEvent.OutputObject
  >;
  getEvent(
    key: "Deposited"
  ): TypedContractEvent<
    DepositedEvent.InputTuple,
    DepositedEvent.OutputTuple,
    DepositedEvent.OutputObject
  >;
  getEvent(
    key: "DepositedAndCalled"
  ): TypedContractEvent<
    DepositedAndCalledEvent.InputTuple,
    DepositedAndCalledEvent.OutputTuple,
    DepositedAndCalledEvent.OutputObject
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
    key: "UpdatedGatewayTSSAddress"
  ): TypedContractEvent<
    UpdatedGatewayTSSAddressEvent.InputTuple,
    UpdatedGatewayTSSAddressEvent.OutputTuple,
    UpdatedGatewayTSSAddressEvent.OutputObject
  >;

  filters: {
    "Called(address,address,bytes,tuple)": TypedContractEvent<
      CalledEvent.InputTuple,
      CalledEvent.OutputTuple,
      CalledEvent.OutputObject
    >;
    Called: TypedContractEvent<
      CalledEvent.InputTuple,
      CalledEvent.OutputTuple,
      CalledEvent.OutputObject
    >;

    "Deposited(address,address,uint256,address,bytes,tuple)": TypedContractEvent<
      DepositedEvent.InputTuple,
      DepositedEvent.OutputTuple,
      DepositedEvent.OutputObject
    >;
    Deposited: TypedContractEvent<
      DepositedEvent.InputTuple,
      DepositedEvent.OutputTuple,
      DepositedEvent.OutputObject
    >;

    "DepositedAndCalled(address,address,uint256,address,bytes,tuple)": TypedContractEvent<
      DepositedAndCalledEvent.InputTuple,
      DepositedAndCalledEvent.OutputTuple,
      DepositedAndCalledEvent.OutputObject
    >;
    DepositedAndCalled: TypedContractEvent<
      DepositedAndCalledEvent.InputTuple,
      DepositedAndCalledEvent.OutputTuple,
      DepositedAndCalledEvent.OutputObject
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

    "Reverted(address,address,uint256,bytes,tuple)": TypedContractEvent<
      RevertedEvent.InputTuple,
      RevertedEvent.OutputTuple,
      RevertedEvent.OutputObject
    >;
    Reverted: TypedContractEvent<
      RevertedEvent.InputTuple,
      RevertedEvent.OutputTuple,
      RevertedEvent.OutputObject
    >;

    "UpdatedGatewayTSSAddress(address,address)": TypedContractEvent<
      UpdatedGatewayTSSAddressEvent.InputTuple,
      UpdatedGatewayTSSAddressEvent.OutputTuple,
      UpdatedGatewayTSSAddressEvent.OutputObject
    >;
    UpdatedGatewayTSSAddress: TypedContractEvent<
      UpdatedGatewayTSSAddressEvent.InputTuple,
      UpdatedGatewayTSSAddressEvent.OutputTuple,
      UpdatedGatewayTSSAddressEvent.OutputObject
    >;
  };
}
