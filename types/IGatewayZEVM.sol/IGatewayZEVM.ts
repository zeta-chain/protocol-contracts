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

export type MessageContextStruct = {
  origin: BytesLike;
  sender: AddressLike;
  chainID: BigNumberish;
};

export type MessageContextStructOutput = [
  origin: string,
  sender: string,
  chainID: bigint
] & { origin: string; sender: string; chainID: bigint };

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

export interface IGatewayZEVMInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "call"
      | "deposit"
      | "depositAndCall((bytes,address,uint256),uint256,address,bytes)"
      | "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)"
      | "depositAndRevert"
      | "execute"
      | "executeRevert"
      | "withdraw(bytes,uint256,address,(address,bool,address,bytes,uint256))"
      | "withdraw(bytes,uint256,uint256,(address,bool,address,bytes,uint256))"
      | "withdrawAndCall(bytes,uint256,uint256,bytes,(uint256,bool),(address,bool,address,bytes,uint256))"
      | "withdrawAndCall(bytes,uint256,address,bytes,(uint256,bool),(address,bool,address,bytes,uint256))"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic: "Called" | "Withdrawn" | "WithdrawnAndCalled"
  ): EventFragment;

  encodeFunctionData(
    functionFragment: "call",
    values: [
      BytesLike,
      AddressLike,
      BytesLike,
      CallOptionsStruct,
      RevertOptionsStruct
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "deposit",
    values: [AddressLike, BigNumberish, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "depositAndCall((bytes,address,uint256),uint256,address,bytes)",
    values: [MessageContextStruct, BigNumberish, AddressLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)",
    values: [
      MessageContextStruct,
      AddressLike,
      BigNumberish,
      AddressLike,
      BytesLike
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "depositAndRevert",
    values: [AddressLike, BigNumberish, AddressLike, RevertContextStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "execute",
    values: [
      MessageContextStruct,
      AddressLike,
      BigNumberish,
      AddressLike,
      BytesLike
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "executeRevert",
    values: [AddressLike, RevertContextStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdraw(bytes,uint256,address,(address,bool,address,bytes,uint256))",
    values: [BytesLike, BigNumberish, AddressLike, RevertOptionsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdraw(bytes,uint256,uint256,(address,bool,address,bytes,uint256))",
    values: [BytesLike, BigNumberish, BigNumberish, RevertOptionsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawAndCall(bytes,uint256,uint256,bytes,(uint256,bool),(address,bool,address,bytes,uint256))",
    values: [
      BytesLike,
      BigNumberish,
      BigNumberish,
      BytesLike,
      CallOptionsStruct,
      RevertOptionsStruct
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawAndCall(bytes,uint256,address,bytes,(uint256,bool),(address,bool,address,bytes,uint256))",
    values: [
      BytesLike,
      BigNumberish,
      AddressLike,
      BytesLike,
      CallOptionsStruct,
      RevertOptionsStruct
    ]
  ): string;

  decodeFunctionResult(functionFragment: "call", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "deposit", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "depositAndCall((bytes,address,uint256),uint256,address,bytes)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "depositAndRevert",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "execute", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "executeRevert",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdraw(bytes,uint256,address,(address,bool,address,bytes,uint256))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdraw(bytes,uint256,uint256,(address,bool,address,bytes,uint256))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawAndCall(bytes,uint256,uint256,bytes,(uint256,bool),(address,bool,address,bytes,uint256))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawAndCall(bytes,uint256,address,bytes,(uint256,bool),(address,bool,address,bytes,uint256))",
    data: BytesLike
  ): Result;
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

export interface IGatewayZEVM extends BaseContract {
  connect(runner?: ContractRunner | null): IGatewayZEVM;
  waitForDeployment(): Promise<this>;

  interface: IGatewayZEVMInterface;

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
      receiver: BytesLike,
      zrc20: AddressLike,
      message: BytesLike,
      callOptions: CallOptionsStruct,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  deposit: TypedContractMethod<
    [zrc20: AddressLike, amount: BigNumberish, target: AddressLike],
    [void],
    "nonpayable"
  >;

  "depositAndCall((bytes,address,uint256),uint256,address,bytes)": TypedContractMethod<
    [
      context: MessageContextStruct,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)": TypedContractMethod<
    [
      context: MessageContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  depositAndRevert: TypedContractMethod<
    [
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      revertContext: RevertContextStruct
    ],
    [void],
    "nonpayable"
  >;

  execute: TypedContractMethod<
    [
      context: MessageContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  executeRevert: TypedContractMethod<
    [target: AddressLike, revertContext: RevertContextStruct],
    [void],
    "nonpayable"
  >;

  "withdraw(bytes,uint256,address,(address,bool,address,bytes,uint256))": TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      zrc20: AddressLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  "withdraw(bytes,uint256,uint256,(address,bool,address,bytes,uint256))": TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      chainId: BigNumberish,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  "withdrawAndCall(bytes,uint256,uint256,bytes,(uint256,bool),(address,bool,address,bytes,uint256))": TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      chainId: BigNumberish,
      message: BytesLike,
      callOptions: CallOptionsStruct,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  "withdrawAndCall(bytes,uint256,address,bytes,(uint256,bool),(address,bool,address,bytes,uint256))": TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      zrc20: AddressLike,
      message: BytesLike,
      callOptions: CallOptionsStruct,
      revertOptions: RevertOptionsStruct
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
      receiver: BytesLike,
      zrc20: AddressLike,
      message: BytesLike,
      callOptions: CallOptionsStruct,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "deposit"
  ): TypedContractMethod<
    [zrc20: AddressLike, amount: BigNumberish, target: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "depositAndCall((bytes,address,uint256),uint256,address,bytes)"
  ): TypedContractMethod<
    [
      context: MessageContextStruct,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)"
  ): TypedContractMethod<
    [
      context: MessageContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "depositAndRevert"
  ): TypedContractMethod<
    [
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      revertContext: RevertContextStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "execute"
  ): TypedContractMethod<
    [
      context: MessageContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "executeRevert"
  ): TypedContractMethod<
    [target: AddressLike, revertContext: RevertContextStruct],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdraw(bytes,uint256,address,(address,bool,address,bytes,uint256))"
  ): TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      zrc20: AddressLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdraw(bytes,uint256,uint256,(address,bool,address,bytes,uint256))"
  ): TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      chainId: BigNumberish,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdrawAndCall(bytes,uint256,uint256,bytes,(uint256,bool),(address,bool,address,bytes,uint256))"
  ): TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      chainId: BigNumberish,
      message: BytesLike,
      callOptions: CallOptionsStruct,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdrawAndCall(bytes,uint256,address,bytes,(uint256,bool),(address,bool,address,bytes,uint256))"
  ): TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      zrc20: AddressLike,
      message: BytesLike,
      callOptions: CallOptionsStruct,
      revertOptions: RevertOptionsStruct
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