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

export interface GatewayZEVMInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "DEFAULT_ADMIN_ROLE"
      | "MAX_GAS_LIMIT"
      | "MAX_MESSAGE_SIZE"
      | "PAUSER_ROLE"
      | "PROTOCOL_ADDRESS"
      | "UPGRADE_INTERFACE_VERSION"
      | "call"
      | "deposit"
      | "depositAndCall((bytes,address,uint256),uint256,address,bytes)"
      | "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)"
      | "depositAndRevert"
      | "execute"
      | "executeRevert"
      | "getRoleAdmin"
      | "grantRole"
      | "hasRole"
      | "initialize"
      | "pause"
      | "paused"
      | "proxiableUUID"
      | "renounceRole"
      | "revokeRole"
      | "supportsInterface"
      | "unpause"
      | "upgradeToAndCall"
      | "withdraw(bytes,uint256,address,(address,bool,address,bytes,uint256))"
      | "withdraw(bytes,uint256,uint256,(address,bool,address,bytes,uint256))"
      | "withdrawAndCall(bytes,uint256,uint256,bytes,(uint256,bool),(address,bool,address,bytes,uint256))"
      | "withdrawAndCall(bytes,uint256,address,bytes,(uint256,bool),(address,bool,address,bytes,uint256))"
      | "zetaToken"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic:
      | "Called"
      | "Initialized"
      | "Paused"
      | "RoleAdminChanged"
      | "RoleGranted"
      | "RoleRevoked"
      | "Unpaused"
      | "Upgraded"
      | "Withdrawn"
      | "WithdrawnAndCalled"
  ): EventFragment;

  encodeFunctionData(
    functionFragment: "DEFAULT_ADMIN_ROLE",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "MAX_GAS_LIMIT",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "MAX_MESSAGE_SIZE",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "PAUSER_ROLE",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "PROTOCOL_ADDRESS",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "UPGRADE_INTERFACE_VERSION",
    values?: undefined
  ): string;
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
    functionFragment: "getRoleAdmin",
    values: [BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "grantRole",
    values: [BytesLike, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "hasRole",
    values: [BytesLike, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "initialize",
    values: [AddressLike, AddressLike]
  ): string;
  encodeFunctionData(functionFragment: "pause", values?: undefined): string;
  encodeFunctionData(functionFragment: "paused", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "proxiableUUID",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "renounceRole",
    values: [BytesLike, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "revokeRole",
    values: [BytesLike, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "supportsInterface",
    values: [BytesLike]
  ): string;
  encodeFunctionData(functionFragment: "unpause", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "upgradeToAndCall",
    values: [AddressLike, BytesLike]
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
  encodeFunctionData(functionFragment: "zetaToken", values?: undefined): string;

  decodeFunctionResult(
    functionFragment: "DEFAULT_ADMIN_ROLE",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "MAX_GAS_LIMIT",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "MAX_MESSAGE_SIZE",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "PAUSER_ROLE",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "PROTOCOL_ADDRESS",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "UPGRADE_INTERFACE_VERSION",
    data: BytesLike
  ): Result;
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
    functionFragment: "getRoleAdmin",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "grantRole", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "hasRole", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "initialize", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "pause", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "paused", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "proxiableUUID",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "renounceRole",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "revokeRole", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "supportsInterface",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "unpause", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "upgradeToAndCall",
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
  decodeFunctionResult(functionFragment: "zetaToken", data: BytesLike): Result;
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

export namespace InitializedEvent {
  export type InputTuple = [version: BigNumberish];
  export type OutputTuple = [version: bigint];
  export interface OutputObject {
    version: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace PausedEvent {
  export type InputTuple = [account: AddressLike];
  export type OutputTuple = [account: string];
  export interface OutputObject {
    account: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace RoleAdminChangedEvent {
  export type InputTuple = [
    role: BytesLike,
    previousAdminRole: BytesLike,
    newAdminRole: BytesLike
  ];
  export type OutputTuple = [
    role: string,
    previousAdminRole: string,
    newAdminRole: string
  ];
  export interface OutputObject {
    role: string;
    previousAdminRole: string;
    newAdminRole: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace RoleGrantedEvent {
  export type InputTuple = [
    role: BytesLike,
    account: AddressLike,
    sender: AddressLike
  ];
  export type OutputTuple = [role: string, account: string, sender: string];
  export interface OutputObject {
    role: string;
    account: string;
    sender: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace RoleRevokedEvent {
  export type InputTuple = [
    role: BytesLike,
    account: AddressLike,
    sender: AddressLike
  ];
  export type OutputTuple = [role: string, account: string, sender: string];
  export interface OutputObject {
    role: string;
    account: string;
    sender: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace UnpausedEvent {
  export type InputTuple = [account: AddressLike];
  export type OutputTuple = [account: string];
  export interface OutputObject {
    account: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace UpgradedEvent {
  export type InputTuple = [implementation: AddressLike];
  export type OutputTuple = [implementation: string];
  export interface OutputObject {
    implementation: string;
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

export interface GatewayZEVM extends BaseContract {
  connect(runner?: ContractRunner | null): GatewayZEVM;
  waitForDeployment(): Promise<this>;

  interface: GatewayZEVMInterface;

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

  DEFAULT_ADMIN_ROLE: TypedContractMethod<[], [string], "view">;

  MAX_GAS_LIMIT: TypedContractMethod<[], [bigint], "view">;

  MAX_MESSAGE_SIZE: TypedContractMethod<[], [bigint], "view">;

  PAUSER_ROLE: TypedContractMethod<[], [string], "view">;

  PROTOCOL_ADDRESS: TypedContractMethod<[], [string], "view">;

  UPGRADE_INTERFACE_VERSION: TypedContractMethod<[], [string], "view">;

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

  getRoleAdmin: TypedContractMethod<[role: BytesLike], [string], "view">;

  grantRole: TypedContractMethod<
    [role: BytesLike, account: AddressLike],
    [void],
    "nonpayable"
  >;

  hasRole: TypedContractMethod<
    [role: BytesLike, account: AddressLike],
    [boolean],
    "view"
  >;

  initialize: TypedContractMethod<
    [zetaToken_: AddressLike, admin_: AddressLike],
    [void],
    "nonpayable"
  >;

  pause: TypedContractMethod<[], [void], "nonpayable">;

  paused: TypedContractMethod<[], [boolean], "view">;

  proxiableUUID: TypedContractMethod<[], [string], "view">;

  renounceRole: TypedContractMethod<
    [role: BytesLike, callerConfirmation: AddressLike],
    [void],
    "nonpayable"
  >;

  revokeRole: TypedContractMethod<
    [role: BytesLike, account: AddressLike],
    [void],
    "nonpayable"
  >;

  supportsInterface: TypedContractMethod<
    [interfaceId: BytesLike],
    [boolean],
    "view"
  >;

  unpause: TypedContractMethod<[], [void], "nonpayable">;

  upgradeToAndCall: TypedContractMethod<
    [newImplementation: AddressLike, data: BytesLike],
    [void],
    "payable"
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
      arg0: BytesLike,
      arg1: BigNumberish,
      arg2: BigNumberish,
      arg3: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  "withdrawAndCall(bytes,uint256,uint256,bytes,(uint256,bool),(address,bool,address,bytes,uint256))": TypedContractMethod<
    [
      arg0: BytesLike,
      arg1: BigNumberish,
      arg2: BigNumberish,
      arg3: BytesLike,
      arg4: CallOptionsStruct,
      arg5: RevertOptionsStruct
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

  zetaToken: TypedContractMethod<[], [string], "view">;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "DEFAULT_ADMIN_ROLE"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "MAX_GAS_LIMIT"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "MAX_MESSAGE_SIZE"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "PAUSER_ROLE"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "PROTOCOL_ADDRESS"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "UPGRADE_INTERFACE_VERSION"
  ): TypedContractMethod<[], [string], "view">;
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
    nameOrSignature: "getRoleAdmin"
  ): TypedContractMethod<[role: BytesLike], [string], "view">;
  getFunction(
    nameOrSignature: "grantRole"
  ): TypedContractMethod<
    [role: BytesLike, account: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "hasRole"
  ): TypedContractMethod<
    [role: BytesLike, account: AddressLike],
    [boolean],
    "view"
  >;
  getFunction(
    nameOrSignature: "initialize"
  ): TypedContractMethod<
    [zetaToken_: AddressLike, admin_: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "pause"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "paused"
  ): TypedContractMethod<[], [boolean], "view">;
  getFunction(
    nameOrSignature: "proxiableUUID"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "renounceRole"
  ): TypedContractMethod<
    [role: BytesLike, callerConfirmation: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "revokeRole"
  ): TypedContractMethod<
    [role: BytesLike, account: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "supportsInterface"
  ): TypedContractMethod<[interfaceId: BytesLike], [boolean], "view">;
  getFunction(
    nameOrSignature: "unpause"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "upgradeToAndCall"
  ): TypedContractMethod<
    [newImplementation: AddressLike, data: BytesLike],
    [void],
    "payable"
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
      arg0: BytesLike,
      arg1: BigNumberish,
      arg2: BigNumberish,
      arg3: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "withdrawAndCall(bytes,uint256,uint256,bytes,(uint256,bool),(address,bool,address,bytes,uint256))"
  ): TypedContractMethod<
    [
      arg0: BytesLike,
      arg1: BigNumberish,
      arg2: BigNumberish,
      arg3: BytesLike,
      arg4: CallOptionsStruct,
      arg5: RevertOptionsStruct
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
  getFunction(
    nameOrSignature: "zetaToken"
  ): TypedContractMethod<[], [string], "view">;

  getEvent(
    key: "Called"
  ): TypedContractEvent<
    CalledEvent.InputTuple,
    CalledEvent.OutputTuple,
    CalledEvent.OutputObject
  >;
  getEvent(
    key: "Initialized"
  ): TypedContractEvent<
    InitializedEvent.InputTuple,
    InitializedEvent.OutputTuple,
    InitializedEvent.OutputObject
  >;
  getEvent(
    key: "Paused"
  ): TypedContractEvent<
    PausedEvent.InputTuple,
    PausedEvent.OutputTuple,
    PausedEvent.OutputObject
  >;
  getEvent(
    key: "RoleAdminChanged"
  ): TypedContractEvent<
    RoleAdminChangedEvent.InputTuple,
    RoleAdminChangedEvent.OutputTuple,
    RoleAdminChangedEvent.OutputObject
  >;
  getEvent(
    key: "RoleGranted"
  ): TypedContractEvent<
    RoleGrantedEvent.InputTuple,
    RoleGrantedEvent.OutputTuple,
    RoleGrantedEvent.OutputObject
  >;
  getEvent(
    key: "RoleRevoked"
  ): TypedContractEvent<
    RoleRevokedEvent.InputTuple,
    RoleRevokedEvent.OutputTuple,
    RoleRevokedEvent.OutputObject
  >;
  getEvent(
    key: "Unpaused"
  ): TypedContractEvent<
    UnpausedEvent.InputTuple,
    UnpausedEvent.OutputTuple,
    UnpausedEvent.OutputObject
  >;
  getEvent(
    key: "Upgraded"
  ): TypedContractEvent<
    UpgradedEvent.InputTuple,
    UpgradedEvent.OutputTuple,
    UpgradedEvent.OutputObject
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

    "Initialized(uint64)": TypedContractEvent<
      InitializedEvent.InputTuple,
      InitializedEvent.OutputTuple,
      InitializedEvent.OutputObject
    >;
    Initialized: TypedContractEvent<
      InitializedEvent.InputTuple,
      InitializedEvent.OutputTuple,
      InitializedEvent.OutputObject
    >;

    "Paused(address)": TypedContractEvent<
      PausedEvent.InputTuple,
      PausedEvent.OutputTuple,
      PausedEvent.OutputObject
    >;
    Paused: TypedContractEvent<
      PausedEvent.InputTuple,
      PausedEvent.OutputTuple,
      PausedEvent.OutputObject
    >;

    "RoleAdminChanged(bytes32,bytes32,bytes32)": TypedContractEvent<
      RoleAdminChangedEvent.InputTuple,
      RoleAdminChangedEvent.OutputTuple,
      RoleAdminChangedEvent.OutputObject
    >;
    RoleAdminChanged: TypedContractEvent<
      RoleAdminChangedEvent.InputTuple,
      RoleAdminChangedEvent.OutputTuple,
      RoleAdminChangedEvent.OutputObject
    >;

    "RoleGranted(bytes32,address,address)": TypedContractEvent<
      RoleGrantedEvent.InputTuple,
      RoleGrantedEvent.OutputTuple,
      RoleGrantedEvent.OutputObject
    >;
    RoleGranted: TypedContractEvent<
      RoleGrantedEvent.InputTuple,
      RoleGrantedEvent.OutputTuple,
      RoleGrantedEvent.OutputObject
    >;

    "RoleRevoked(bytes32,address,address)": TypedContractEvent<
      RoleRevokedEvent.InputTuple,
      RoleRevokedEvent.OutputTuple,
      RoleRevokedEvent.OutputObject
    >;
    RoleRevoked: TypedContractEvent<
      RoleRevokedEvent.InputTuple,
      RoleRevokedEvent.OutputTuple,
      RoleRevokedEvent.OutputObject
    >;

    "Unpaused(address)": TypedContractEvent<
      UnpausedEvent.InputTuple,
      UnpausedEvent.OutputTuple,
      UnpausedEvent.OutputObject
    >;
    Unpaused: TypedContractEvent<
      UnpausedEvent.InputTuple,
      UnpausedEvent.OutputTuple,
      UnpausedEvent.OutputObject
    >;

    "Upgraded(address)": TypedContractEvent<
      UpgradedEvent.InputTuple,
      UpgradedEvent.OutputTuple,
      UpgradedEvent.OutputObject
    >;
    Upgraded: TypedContractEvent<
      UpgradedEvent.InputTuple,
      UpgradedEvent.OutputTuple,
      UpgradedEvent.OutputObject
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
