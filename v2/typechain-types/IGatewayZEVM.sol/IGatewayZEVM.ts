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
  origin: BytesLike;
  sender: AddressLike;
  chainID: BigNumberish;
};

export type RevertContextStructOutput = [
  origin: string,
  sender: string,
  chainID: bigint
] & { origin: string; sender: string; chainID: bigint };

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
      | "withdraw(bytes,uint256,address)"
      | "withdraw(bytes,uint256,uint256)"
      | "withdrawAndCall(bytes,uint256,address,bytes)"
      | "withdrawAndCall(bytes,uint256,uint256,bytes)"
  ): FunctionFragment;

  getEvent(nameOrSignatureOrTopic: "Call" | "Withdrawal"): EventFragment;

  encodeFunctionData(
    functionFragment: "call",
    values: [BytesLike, BigNumberish, BytesLike]
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
      RevertContextStruct,
      AddressLike,
      BigNumberish,
      AddressLike,
      BytesLike
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "execute",
    values: [ZContextStruct, AddressLike, BigNumberish, AddressLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "executeRevert",
    values: [
      RevertContextStruct,
      AddressLike,
      BigNumberish,
      AddressLike,
      BytesLike
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "withdraw(bytes,uint256,address)",
    values: [BytesLike, BigNumberish, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "withdraw(bytes,uint256,uint256)",
    values: [BytesLike, BigNumberish, BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawAndCall(bytes,uint256,address,bytes)",
    values: [BytesLike, BigNumberish, AddressLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawAndCall(bytes,uint256,uint256,bytes)",
    values: [BytesLike, BigNumberish, BigNumberish, BytesLike]
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
    functionFragment: "withdraw(bytes,uint256,address)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdraw(bytes,uint256,uint256)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawAndCall(bytes,uint256,address,bytes)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawAndCall(bytes,uint256,uint256,bytes)",
    data: BytesLike
  ): Result;
}

export namespace CallEvent {
  export type InputTuple = [
    sender: AddressLike,
    chainId: BigNumberish,
    receiver: BytesLike,
    message: BytesLike
  ];
  export type OutputTuple = [
    sender: string,
    chainId: bigint,
    receiver: string,
    message: string
  ];
  export interface OutputObject {
    sender: string;
    chainId: bigint;
    receiver: string;
    message: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace WithdrawalEvent {
  export type InputTuple = [
    sender: AddressLike,
    chainId: BigNumberish,
    receiver: BytesLike,
    zrc20: AddressLike,
    value: BigNumberish,
    gasfee: BigNumberish,
    protocolFlatFee: BigNumberish,
    message: BytesLike
  ];
  export type OutputTuple = [
    sender: string,
    chainId: bigint,
    receiver: string,
    zrc20: string,
    value: bigint,
    gasfee: bigint,
    protocolFlatFee: bigint,
    message: string
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
    [receiver: BytesLike, chainId: BigNumberish, message: BytesLike],
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
      context: RevertContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike
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
      context: RevertContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  "withdraw(bytes,uint256,address)": TypedContractMethod<
    [receiver: BytesLike, amount: BigNumberish, zrc20: AddressLike],
    [void],
    "nonpayable"
  >;

  "withdraw(bytes,uint256,uint256)": TypedContractMethod<
    [receiver: BytesLike, amount: BigNumberish, chainId: BigNumberish],
    [void],
    "nonpayable"
  >;

  "withdrawAndCall(bytes,uint256,address,bytes)": TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      zrc20: AddressLike,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  "withdrawAndCall(bytes,uint256,uint256,bytes)": TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      chainId: BigNumberish,
      message: BytesLike
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
    [receiver: BytesLike, chainId: BigNumberish, message: BytesLike],
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
      context: RevertContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike
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
      context: RevertContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      target: AddressLike,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdraw(bytes,uint256,address)"
  ): TypedContractMethod<
    [receiver: BytesLike, amount: BigNumberish, zrc20: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdraw(bytes,uint256,uint256)"
  ): TypedContractMethod<
    [receiver: BytesLike, amount: BigNumberish, chainId: BigNumberish],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdrawAndCall(bytes,uint256,address,bytes)"
  ): TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      zrc20: AddressLike,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdrawAndCall(bytes,uint256,uint256,bytes)"
  ): TypedContractMethod<
    [
      receiver: BytesLike,
      amount: BigNumberish,
      chainId: BigNumberish,
      message: BytesLike
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
    key: "Withdrawal"
  ): TypedContractEvent<
    WithdrawalEvent.InputTuple,
    WithdrawalEvent.OutputTuple,
    WithdrawalEvent.OutputObject
  >;

  filters: {
    "Call(address,uint256,bytes,bytes)": TypedContractEvent<
      CallEvent.InputTuple,
      CallEvent.OutputTuple,
      CallEvent.OutputObject
    >;
    Call: TypedContractEvent<
      CallEvent.InputTuple,
      CallEvent.OutputTuple,
      CallEvent.OutputObject
    >;

    "Withdrawal(address,uint256,bytes,address,uint256,uint256,uint256,bytes)": TypedContractEvent<
      WithdrawalEvent.InputTuple,
      WithdrawalEvent.OutputTuple,
      WithdrawalEvent.OutputObject
    >;
    Withdrawal: TypedContractEvent<
      WithdrawalEvent.InputTuple,
      WithdrawalEvent.OutputTuple,
      WithdrawalEvent.OutputObject
    >;
  };
}
