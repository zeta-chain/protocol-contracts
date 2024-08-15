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

export interface GatewayEVMEchidnaTestInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "ASSET_HANDLER_ROLE"
      | "DEFAULT_ADMIN_ROLE"
      | "PAUSER_ROLE"
      | "TSS_ROLE"
      | "UPGRADE_INTERFACE_VERSION"
      | "call"
      | "custody"
      | "deposit(address,uint256,address,(address,bool,address,bytes,uint256))"
      | "deposit(address,(address,bool,address,bytes,uint256))"
      | "depositAndCall(address,bytes,(address,bool,address,bytes,uint256))"
      | "depositAndCall(address,uint256,address,bytes,(address,bool,address,bytes,uint256))"
      | "echidnaCaller"
      | "execute"
      | "executeRevert"
      | "executeWithERC20"
      | "getRoleAdmin"
      | "grantRole"
      | "hasRole"
      | "initialize"
      | "pause"
      | "paused"
      | "proxiableUUID"
      | "renounceRole"
      | "revertWithERC20"
      | "revokeRole"
      | "setConnector"
      | "setCustody"
      | "supportsInterface"
      | "testERC20"
      | "testExecuteWithERC20"
      | "tssAddress"
      | "unpause"
      | "upgradeToAndCall"
      | "zetaConnector"
      | "zetaToken"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic:
      | "Called"
      | "Deposited"
      | "Executed"
      | "ExecutedWithERC20"
      | "Initialized"
      | "Paused"
      | "Reverted"
      | "RoleAdminChanged"
      | "RoleGranted"
      | "RoleRevoked"
      | "Unpaused"
      | "Upgraded"
  ): EventFragment;

  encodeFunctionData(
    functionFragment: "ASSET_HANDLER_ROLE",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "DEFAULT_ADMIN_ROLE",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "PAUSER_ROLE",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "TSS_ROLE", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "UPGRADE_INTERFACE_VERSION",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "call",
    values: [AddressLike, BytesLike, RevertOptionsStruct]
  ): string;
  encodeFunctionData(functionFragment: "custody", values?: undefined): string;
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
    functionFragment: "echidnaCaller",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "execute",
    values: [AddressLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "executeRevert",
    values: [AddressLike, BytesLike, RevertContextStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "executeWithERC20",
    values: [AddressLike, AddressLike, BigNumberish, BytesLike]
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
    values: [AddressLike, AddressLike, AddressLike]
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
    functionFragment: "revertWithERC20",
    values: [
      AddressLike,
      AddressLike,
      BigNumberish,
      BytesLike,
      RevertContextStruct
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "revokeRole",
    values: [BytesLike, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "setConnector",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "setCustody",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "supportsInterface",
    values: [BytesLike]
  ): string;
  encodeFunctionData(functionFragment: "testERC20", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "testExecuteWithERC20",
    values: [AddressLike, BigNumberish, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "tssAddress",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "unpause", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "upgradeToAndCall",
    values: [AddressLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "zetaConnector",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "zetaToken", values?: undefined): string;

  decodeFunctionResult(
    functionFragment: "ASSET_HANDLER_ROLE",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "DEFAULT_ADMIN_ROLE",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "PAUSER_ROLE",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "TSS_ROLE", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "UPGRADE_INTERFACE_VERSION",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "call", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "custody", data: BytesLike): Result;
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
  decodeFunctionResult(
    functionFragment: "echidnaCaller",
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
  decodeFunctionResult(
    functionFragment: "revertWithERC20",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "revokeRole", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "setConnector",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "setCustody", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "supportsInterface",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "testERC20", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "testExecuteWithERC20",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "tssAddress", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "unpause", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "upgradeToAndCall",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "zetaConnector",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "zetaToken", data: BytesLike): Result;
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

export interface GatewayEVMEchidnaTest extends BaseContract {
  connect(runner?: ContractRunner | null): GatewayEVMEchidnaTest;
  waitForDeployment(): Promise<this>;

  interface: GatewayEVMEchidnaTestInterface;

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

  ASSET_HANDLER_ROLE: TypedContractMethod<[], [string], "view">;

  DEFAULT_ADMIN_ROLE: TypedContractMethod<[], [string], "view">;

  PAUSER_ROLE: TypedContractMethod<[], [string], "view">;

  TSS_ROLE: TypedContractMethod<[], [string], "view">;

  UPGRADE_INTERFACE_VERSION: TypedContractMethod<[], [string], "view">;

  call: TypedContractMethod<
    [
      receiver: AddressLike,
      payload: BytesLike,
      revertOptions: RevertOptionsStruct
    ],
    [void],
    "nonpayable"
  >;

  custody: TypedContractMethod<[], [string], "view">;

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

  echidnaCaller: TypedContractMethod<[], [string], "view">;

  execute: TypedContractMethod<
    [destination: AddressLike, data: BytesLike],
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
      token: AddressLike,
      to: AddressLike,
      amount: BigNumberish,
      data: BytesLike
    ],
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
    [tssAddress_: AddressLike, zetaToken_: AddressLike, admin_: AddressLike],
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

  revokeRole: TypedContractMethod<
    [role: BytesLike, account: AddressLike],
    [void],
    "nonpayable"
  >;

  setConnector: TypedContractMethod<
    [zetaConnector_: AddressLike],
    [void],
    "nonpayable"
  >;

  setCustody: TypedContractMethod<
    [custody_: AddressLike],
    [void],
    "nonpayable"
  >;

  supportsInterface: TypedContractMethod<
    [interfaceId: BytesLike],
    [boolean],
    "view"
  >;

  testERC20: TypedContractMethod<[], [string], "view">;

  testExecuteWithERC20: TypedContractMethod<
    [to: AddressLike, amount: BigNumberish, data: BytesLike],
    [void],
    "nonpayable"
  >;

  tssAddress: TypedContractMethod<[], [string], "view">;

  unpause: TypedContractMethod<[], [void], "nonpayable">;

  upgradeToAndCall: TypedContractMethod<
    [newImplementation: AddressLike, data: BytesLike],
    [void],
    "payable"
  >;

  zetaConnector: TypedContractMethod<[], [string], "view">;

  zetaToken: TypedContractMethod<[], [string], "view">;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "ASSET_HANDLER_ROLE"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "DEFAULT_ADMIN_ROLE"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "PAUSER_ROLE"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "TSS_ROLE"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "UPGRADE_INTERFACE_VERSION"
  ): TypedContractMethod<[], [string], "view">;
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
    nameOrSignature: "custody"
  ): TypedContractMethod<[], [string], "view">;
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
    nameOrSignature: "echidnaCaller"
  ): TypedContractMethod<[], [string], "view">;
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
      token: AddressLike,
      to: AddressLike,
      amount: BigNumberish,
      data: BytesLike
    ],
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
    [tssAddress_: AddressLike, zetaToken_: AddressLike, admin_: AddressLike],
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
  getFunction(
    nameOrSignature: "revokeRole"
  ): TypedContractMethod<
    [role: BytesLike, account: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "setConnector"
  ): TypedContractMethod<[zetaConnector_: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "setCustody"
  ): TypedContractMethod<[custody_: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "supportsInterface"
  ): TypedContractMethod<[interfaceId: BytesLike], [boolean], "view">;
  getFunction(
    nameOrSignature: "testERC20"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "testExecuteWithERC20"
  ): TypedContractMethod<
    [to: AddressLike, amount: BigNumberish, data: BytesLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "tssAddress"
  ): TypedContractMethod<[], [string], "view">;
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
    nameOrSignature: "zetaConnector"
  ): TypedContractMethod<[], [string], "view">;
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
    key: "Deposited"
  ): TypedContractEvent<
    DepositedEvent.InputTuple,
    DepositedEvent.OutputTuple,
    DepositedEvent.OutputObject
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
    key: "Reverted"
  ): TypedContractEvent<
    RevertedEvent.InputTuple,
    RevertedEvent.OutputTuple,
    RevertedEvent.OutputObject
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
  };
}
