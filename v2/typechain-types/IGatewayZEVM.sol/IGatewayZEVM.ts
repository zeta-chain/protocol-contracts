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
      | "withdraw(bytes,uint256,address,(address,bool,address,bytes))"
      | "withdraw(bytes,uint256,uint256,(address,bool,address,bytes))"
      | "withdrawAndCall(bytes,uint256,uint256,bytes,(address,bool,address,bytes))"
      | "withdrawAndCall(bytes,uint256,address,bytes,uint256,(address,bool,address,bytes))"
  ): FunctionFragment;

  getEvent(nameOrSignatureOrTopic: "Called" | "Withdrawn"): EventFragment;

  encodeFunctionData(
    functionFragment: "call",
    values: [
      BytesLike,
      AddressLike,
      BytesLike,
      BigNumberish,
      RevertOptionsStruct
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "deposit",
    values: [AddressLike, BigNumberish, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "depositAndCall((bytes,address,uint256),uint256,address,bytes)",
    values: [ZContextStruct, BigNumberish, AddressLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)",
    values: [ZContextStruct, AddressLike, BigNumberish, AddressLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "depositAndRevert",
    values: [
      ZContextStruct,
      AddressLike,
      BigNumberish,
      AddressLike,
      BytesLike,
      RevertContextStruct
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "execute",
    values: [ZContextStruct, AddressLike, BigNumberish, AddressLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "executeRevert",
    values: [
      ZContextStruct,
      AddressLike,
      BigNumberish,
      AddressLike,
      BytesLike,
      RevertContextStruct
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "withdraw(bytes,uint256,address,(address,bool,address,bytes))",
    values: [BytesLike, BigNumberish, AddressLike, RevertOptionsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdraw(bytes,uint256,uint256,(address,bool,address,bytes))",
    values: [BytesLike, BigNumberish, BigNumberish, RevertOptionsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawAndCall(bytes,uint256,uint256,bytes,(address,bool,address,bytes))",
    values: [
      BytesLike,
      BigNumberish,
      BigNumberish,
      BytesLike,
      RevertOptionsStruct
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawAndCall(bytes,uint256,address,bytes,uint256,(address,bool,address,bytes))",
    values: [
      BytesLike,
      BigNumberish,
      AddressLike,
      BytesLike,
      BigNumberish,
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
    functionFragment: "withdraw(bytes,uint256,address,(address,bool,address,bytes))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdraw(bytes,uint256,uint256,(address,bool,address,bytes))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawAndCall(bytes,uint256,uint256,bytes,(address,bool,address,bytes))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawAndCall(bytes,uint256,address,bytes,uint256,(address,bool,address,bytes))",
    data: BytesLike
  ): Result;
}

export namespace CalledEvent {
  export type InputTuple = [
    sender: AddressLike,
    zrc20: AddressLike,
    receiver: BytesLike,
    message: BytesLike,
    revertOptions: RevertOptionsStruct
  ];
  export type OutputTuple = [
    sender: string,
    zrc20: string,
    receiver: string,
    message: string,
    revertOptions: RevertOptionsStructOutput
  ];
  export interface OutputObject {
    sender: string;
    zrc20: string;
    receiver: string;
    message: string;
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
      gasLimit: BigNumberish,
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
      context: ZContextStruct,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)": TypedContractMethod<
    [
      context: ZContextStruct,
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
      context: ZContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike,
      revertContext: RevertContextStruct
    ],
    [void],
    "nonpayable"
  >;

  execute: TypedContractMethod<
    [
      context: ZContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  executeRevert: TypedContractMethod<
    [
      context: ZContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike,
      revertContext: RevertContextStruct
    ],
    [void],
    "nonpayable"
  >;

  "withdraw(bytes,uint256,address,(address,bool,address,bytes))": TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      zrc20: AddressLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  "withdraw(bytes,uint256,uint256,(address,bool,address,bytes))": TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      chainId: BigNumberish,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  "withdrawAndCall(bytes,uint256,uint256,bytes,(address,bool,address,bytes))": TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      chainId: BigNumberish,
      message: BytesLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  "withdrawAndCall(bytes,uint256,address,bytes,uint256,(address,bool,address,bytes))": TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      zrc20: AddressLike,
      message: BytesLike,
      gasLimit: BigNumberish,
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
      gasLimit: BigNumberish,
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
      context: ZContextStruct,
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
      context: ZContextStruct,
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
      context: ZContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike,
      revertContext: RevertContextStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "execute"
  ): TypedContractMethod<
    [
      context: ZContextStruct,
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
    [
      context: ZContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike,
      revertContext: RevertContextStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdraw(bytes,uint256,address,(address,bool,address,bytes))"
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
    nameOrSignature: "withdraw(bytes,uint256,uint256,(address,bool,address,bytes))"
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
    nameOrSignature: "withdrawAndCall(bytes,uint256,uint256,bytes,(address,bool,address,bytes))"
  ): TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      chainId: BigNumberish,
      message: BytesLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdrawAndCall(bytes,uint256,address,bytes,uint256,(address,bool,address,bytes))"
  ): TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      zrc20: AddressLike,
      message: BytesLike,
      gasLimit: BigNumberish,
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

  filters: {
    "Call(address,address,bytes,bytes,tuple)": TypedContractEvent<
      CallEvent.InputTuple,
      CallEvent.OutputTuple,
      CallEvent.OutputObject
    >;
    Called: TypedContractEvent<
      CalledEvent.InputTuple,
      CalledEvent.OutputTuple,
      CalledEvent.OutputObject
    >;

    "Withdrawal(address,uint256,bytes,address,uint256,uint256,uint256,bytes,tuple)": TypedContractEvent<
      WithdrawalEvent.InputTuple,
      WithdrawalEvent.OutputTuple,
      WithdrawalEvent.OutputObject
    >;
    Withdrawn: TypedContractEvent<
      WithdrawnEvent.InputTuple,
      WithdrawnEvent.OutputTuple,
      WithdrawnEvent.OutputObject
    >;
  };
}
