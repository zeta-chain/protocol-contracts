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

export type MessageContextStruct = { sender: AddressLike };

export type MessageContextStructOutput = [sender: string] & { sender: string };

export interface RegistryInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "DEFAULT_ADMIN_ROLE"
      | "PAUSER_ROLE"
      | "RELAY_ROLE"
      | "UPGRADE_INTERFACE_VERSION"
      | "_activeChains"
      | "changeChainStatus"
      | "coreRegistry"
      | "gatewayEVM"
      | "getActiveChains"
      | "getChainMetadata"
      | "getContractConfiguration"
      | "getContractInfo"
      | "getRoleAdmin"
      | "getZRC20AddressByForeignAsset"
      | "getZRC20TokenInfo"
      | "grantRole"
      | "hasRole"
      | "initialize"
      | "onCall"
      | "pause"
      | "paused"
      | "proxiableUUID"
      | "registerContract"
      | "registerZRC20Token"
      | "renounceRole"
      | "revokeRole"
      | "setContractActive"
      | "supportsInterface"
      | "unpause"
      | "updateChainMetadata"
      | "updateContractConfiguration"
      | "updateZRC20Token"
      | "upgradeToAndCall"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic:
      | "ChainStatusChanged"
      | "ContractRegistered"
      | "ContractStatusChanged"
      | "Initialized"
      | "NewChainMetadata"
      | "NewContractConfiguration"
      | "Paused"
      | "RoleAdminChanged"
      | "RoleGranted"
      | "RoleRevoked"
      | "Unpaused"
      | "Upgraded"
      | "ZRC20TokenRegistered"
      | "ZRC20TokenUpdated"
  ): EventFragment;

  encodeFunctionData(
    functionFragment: "DEFAULT_ADMIN_ROLE",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "PAUSER_ROLE",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "RELAY_ROLE",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "UPGRADE_INTERFACE_VERSION",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "_activeChains",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "changeChainStatus",
    values: [BigNumberish, boolean]
  ): string;
  encodeFunctionData(
    functionFragment: "coreRegistry",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "gatewayEVM",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getActiveChains",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getChainMetadata",
    values: [BigNumberish, string]
  ): string;
  encodeFunctionData(
    functionFragment: "getContractConfiguration",
    values: [BigNumberish, string, string]
  ): string;
  encodeFunctionData(
    functionFragment: "getContractInfo",
    values: [BigNumberish, string]
  ): string;
  encodeFunctionData(
    functionFragment: "getRoleAdmin",
    values: [BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "getZRC20AddressByForeignAsset",
    values: [BigNumberish, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "getZRC20TokenInfo",
    values: [AddressLike]
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
  encodeFunctionData(
    functionFragment: "onCall",
    values: [MessageContextStruct, BytesLike]
  ): string;
  encodeFunctionData(functionFragment: "pause", values?: undefined): string;
  encodeFunctionData(functionFragment: "paused", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "proxiableUUID",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "registerContract",
    values: [BigNumberish, AddressLike, string, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "registerZRC20Token",
    values: [AddressLike, string, BigNumberish, BytesLike, string, BigNumberish]
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
    functionFragment: "setContractActive",
    values: [BigNumberish, string, boolean]
  ): string;
  encodeFunctionData(
    functionFragment: "supportsInterface",
    values: [BytesLike]
  ): string;
  encodeFunctionData(functionFragment: "unpause", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "updateChainMetadata",
    values: [BigNumberish, string, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "updateContractConfiguration",
    values: [BigNumberish, string, string, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "updateZRC20Token",
    values: [AddressLike, boolean]
  ): string;
  encodeFunctionData(
    functionFragment: "upgradeToAndCall",
    values: [AddressLike, BytesLike]
  ): string;

  decodeFunctionResult(
    functionFragment: "DEFAULT_ADMIN_ROLE",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "PAUSER_ROLE",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "RELAY_ROLE", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "UPGRADE_INTERFACE_VERSION",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "_activeChains",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "changeChainStatus",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "coreRegistry",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "gatewayEVM", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "getActiveChains",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getChainMetadata",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getContractConfiguration",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getContractInfo",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getRoleAdmin",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getZRC20AddressByForeignAsset",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getZRC20TokenInfo",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "grantRole", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "hasRole", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "initialize", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "onCall", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "pause", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "paused", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "proxiableUUID",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "registerContract",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "registerZRC20Token",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "renounceRole",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "revokeRole", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "setContractActive",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "supportsInterface",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "unpause", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "updateChainMetadata",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "updateContractConfiguration",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "updateZRC20Token",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "upgradeToAndCall",
    data: BytesLike
  ): Result;
}

export namespace ChainStatusChangedEvent {
  export type InputTuple = [
    chainId: BigNumberish,
    previousState: boolean,
    newState: boolean
  ];
  export type OutputTuple = [
    chainId: bigint,
    previousState: boolean,
    newState: boolean
  ];
  export interface OutputObject {
    chainId: bigint;
    previousState: boolean;
    newState: boolean;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ContractRegisteredEvent {
  export type InputTuple = [
    chainId: BigNumberish,
    contractType: string,
    addressBytes: BytesLike
  ];
  export type OutputTuple = [
    chainId: bigint,
    contractType: string,
    addressBytes: string
  ];
  export interface OutputObject {
    chainId: bigint;
    contractType: string;
    addressBytes: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ContractStatusChangedEvent {
  export type InputTuple = [addressBytes: BytesLike];
  export type OutputTuple = [addressBytes: string];
  export interface OutputObject {
    addressBytes: string;
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

export namespace NewChainMetadataEvent {
  export type InputTuple = [
    chainId: BigNumberish,
    key: string,
    value: BytesLike
  ];
  export type OutputTuple = [chainId: bigint, key: string, value: string];
  export interface OutputObject {
    chainId: bigint;
    key: string;
    value: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace NewContractConfigurationEvent {
  export type InputTuple = [
    chainId: BigNumberish,
    contractType: string,
    key: string,
    value: BytesLike
  ];
  export type OutputTuple = [
    chainId: bigint,
    contractType: string,
    key: string,
    value: string
  ];
  export interface OutputObject {
    chainId: bigint;
    contractType: string;
    key: string;
    value: string;
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

export namespace ZRC20TokenRegisteredEvent {
  export type InputTuple = [
    originAddress: BytesLike,
    zrc20: AddressLike,
    decimals: BigNumberish,
    originChainId: BigNumberish,
    symbol: string
  ];
  export type OutputTuple = [
    originAddress: string,
    zrc20: string,
    decimals: bigint,
    originChainId: bigint,
    symbol: string
  ];
  export interface OutputObject {
    originAddress: string;
    zrc20: string;
    decimals: bigint;
    originChainId: bigint;
    symbol: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ZRC20TokenUpdatedEvent {
  export type InputTuple = [zrc20: AddressLike, active: boolean];
  export type OutputTuple = [zrc20: string, active: boolean];
  export interface OutputObject {
    zrc20: string;
    active: boolean;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface Registry extends BaseContract {
  connect(runner?: ContractRunner | null): Registry;
  waitForDeployment(): Promise<this>;

  interface: RegistryInterface;

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

  PAUSER_ROLE: TypedContractMethod<[], [string], "view">;

  RELAY_ROLE: TypedContractMethod<[], [string], "view">;

  UPGRADE_INTERFACE_VERSION: TypedContractMethod<[], [string], "view">;

  _activeChains: TypedContractMethod<[arg0: BigNumberish], [bigint], "view">;

  changeChainStatus: TypedContractMethod<
    [chainId: BigNumberish, active: boolean],
    [void],
    "nonpayable"
  >;

  coreRegistry: TypedContractMethod<[], [string], "view">;

  gatewayEVM: TypedContractMethod<[], [string], "view">;

  getActiveChains: TypedContractMethod<[], [bigint[]], "view">;

  getChainMetadata: TypedContractMethod<
    [chainId: BigNumberish, key: string],
    [string],
    "view"
  >;

  getContractConfiguration: TypedContractMethod<
    [chainId: BigNumberish, contractType: string, key: string],
    [string],
    "view"
  >;

  getContractInfo: TypedContractMethod<
    [chainId: BigNumberish, contractType: string],
    [[boolean, string] & { active: boolean; address_: string }],
    "view"
  >;

  getRoleAdmin: TypedContractMethod<[role: BytesLike], [string], "view">;

  getZRC20AddressByForeignAsset: TypedContractMethod<
    [originChainId: BigNumberish, originAddress: BytesLike],
    [string],
    "view"
  >;

  getZRC20TokenInfo: TypedContractMethod<
    [address_: AddressLike],
    [
      [boolean, string, bigint, string, string, bigint] & {
        active: boolean;
        symbol: string;
        originChainId: bigint;
        originAddress: string;
        coinType: string;
        decimals: bigint;
      }
    ],
    "view"
  >;

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
    [admin_: AddressLike, gatewayEVM_: AddressLike],
    [void],
    "nonpayable"
  >;

  onCall: TypedContractMethod<
    [context: MessageContextStruct, data: BytesLike],
    [string],
    "nonpayable"
  >;

  pause: TypedContractMethod<[], [void], "nonpayable">;

  paused: TypedContractMethod<[], [boolean], "view">;

  proxiableUUID: TypedContractMethod<[], [string], "view">;

  registerContract: TypedContractMethod<
    [
      chainId: BigNumberish,
      address_: AddressLike,
      contractType: string,
      addressBytes: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  registerZRC20Token: TypedContractMethod<
    [
      address_: AddressLike,
      symbol: string,
      originChainId: BigNumberish,
      originAddress: BytesLike,
      coinType: string,
      decimals: BigNumberish
    ],
    [void],
    "nonpayable"
  >;

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

  setContractActive: TypedContractMethod<
    [chainId: BigNumberish, contractType: string, active: boolean],
    [void],
    "nonpayable"
  >;

  supportsInterface: TypedContractMethod<
    [interfaceId: BytesLike],
    [boolean],
    "view"
  >;

  unpause: TypedContractMethod<[], [void], "nonpayable">;

  updateChainMetadata: TypedContractMethod<
    [chainId: BigNumberish, key: string, value: BytesLike],
    [void],
    "nonpayable"
  >;

  updateContractConfiguration: TypedContractMethod<
    [
      chainId: BigNumberish,
      contractType: string,
      key: string,
      value: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  updateZRC20Token: TypedContractMethod<
    [address_: AddressLike, active: boolean],
    [void],
    "nonpayable"
  >;

  upgradeToAndCall: TypedContractMethod<
    [newImplementation: AddressLike, data: BytesLike],
    [void],
    "payable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "DEFAULT_ADMIN_ROLE"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "PAUSER_ROLE"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "RELAY_ROLE"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "UPGRADE_INTERFACE_VERSION"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "_activeChains"
  ): TypedContractMethod<[arg0: BigNumberish], [bigint], "view">;
  getFunction(
    nameOrSignature: "changeChainStatus"
  ): TypedContractMethod<
    [chainId: BigNumberish, active: boolean],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "coreRegistry"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "gatewayEVM"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "getActiveChains"
  ): TypedContractMethod<[], [bigint[]], "view">;
  getFunction(
    nameOrSignature: "getChainMetadata"
  ): TypedContractMethod<
    [chainId: BigNumberish, key: string],
    [string],
    "view"
  >;
  getFunction(
    nameOrSignature: "getContractConfiguration"
  ): TypedContractMethod<
    [chainId: BigNumberish, contractType: string, key: string],
    [string],
    "view"
  >;
  getFunction(
    nameOrSignature: "getContractInfo"
  ): TypedContractMethod<
    [chainId: BigNumberish, contractType: string],
    [[boolean, string] & { active: boolean; address_: string }],
    "view"
  >;
  getFunction(
    nameOrSignature: "getRoleAdmin"
  ): TypedContractMethod<[role: BytesLike], [string], "view">;
  getFunction(
    nameOrSignature: "getZRC20AddressByForeignAsset"
  ): TypedContractMethod<
    [originChainId: BigNumberish, originAddress: BytesLike],
    [string],
    "view"
  >;
  getFunction(
    nameOrSignature: "getZRC20TokenInfo"
  ): TypedContractMethod<
    [address_: AddressLike],
    [
      [boolean, string, bigint, string, string, bigint] & {
        active: boolean;
        symbol: string;
        originChainId: bigint;
        originAddress: string;
        coinType: string;
        decimals: bigint;
      }
    ],
    "view"
  >;
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
    [admin_: AddressLike, gatewayEVM_: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onCall"
  ): TypedContractMethod<
    [context: MessageContextStruct, data: BytesLike],
    [string],
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
    nameOrSignature: "registerContract"
  ): TypedContractMethod<
    [
      chainId: BigNumberish,
      address_: AddressLike,
      contractType: string,
      addressBytes: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "registerZRC20Token"
  ): TypedContractMethod<
    [
      address_: AddressLike,
      symbol: string,
      originChainId: BigNumberish,
      originAddress: BytesLike,
      coinType: string,
      decimals: BigNumberish
    ],
    [void],
    "nonpayable"
  >;
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
    nameOrSignature: "setContractActive"
  ): TypedContractMethod<
    [chainId: BigNumberish, contractType: string, active: boolean],
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
    nameOrSignature: "updateChainMetadata"
  ): TypedContractMethod<
    [chainId: BigNumberish, key: string, value: BytesLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "updateContractConfiguration"
  ): TypedContractMethod<
    [
      chainId: BigNumberish,
      contractType: string,
      key: string,
      value: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "updateZRC20Token"
  ): TypedContractMethod<
    [address_: AddressLike, active: boolean],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "upgradeToAndCall"
  ): TypedContractMethod<
    [newImplementation: AddressLike, data: BytesLike],
    [void],
    "payable"
  >;

  getEvent(
    key: "ChainStatusChanged"
  ): TypedContractEvent<
    ChainStatusChangedEvent.InputTuple,
    ChainStatusChangedEvent.OutputTuple,
    ChainStatusChangedEvent.OutputObject
  >;
  getEvent(
    key: "ContractRegistered"
  ): TypedContractEvent<
    ContractRegisteredEvent.InputTuple,
    ContractRegisteredEvent.OutputTuple,
    ContractRegisteredEvent.OutputObject
  >;
  getEvent(
    key: "ContractStatusChanged"
  ): TypedContractEvent<
    ContractStatusChangedEvent.InputTuple,
    ContractStatusChangedEvent.OutputTuple,
    ContractStatusChangedEvent.OutputObject
  >;
  getEvent(
    key: "Initialized"
  ): TypedContractEvent<
    InitializedEvent.InputTuple,
    InitializedEvent.OutputTuple,
    InitializedEvent.OutputObject
  >;
  getEvent(
    key: "NewChainMetadata"
  ): TypedContractEvent<
    NewChainMetadataEvent.InputTuple,
    NewChainMetadataEvent.OutputTuple,
    NewChainMetadataEvent.OutputObject
  >;
  getEvent(
    key: "NewContractConfiguration"
  ): TypedContractEvent<
    NewContractConfigurationEvent.InputTuple,
    NewContractConfigurationEvent.OutputTuple,
    NewContractConfigurationEvent.OutputObject
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
    key: "ZRC20TokenRegistered"
  ): TypedContractEvent<
    ZRC20TokenRegisteredEvent.InputTuple,
    ZRC20TokenRegisteredEvent.OutputTuple,
    ZRC20TokenRegisteredEvent.OutputObject
  >;
  getEvent(
    key: "ZRC20TokenUpdated"
  ): TypedContractEvent<
    ZRC20TokenUpdatedEvent.InputTuple,
    ZRC20TokenUpdatedEvent.OutputTuple,
    ZRC20TokenUpdatedEvent.OutputObject
  >;

  filters: {
    "ChainStatusChanged(uint256,bool,bool)": TypedContractEvent<
      ChainStatusChangedEvent.InputTuple,
      ChainStatusChangedEvent.OutputTuple,
      ChainStatusChangedEvent.OutputObject
    >;
    ChainStatusChanged: TypedContractEvent<
      ChainStatusChangedEvent.InputTuple,
      ChainStatusChangedEvent.OutputTuple,
      ChainStatusChangedEvent.OutputObject
    >;

    "ContractRegistered(uint256,string,bytes)": TypedContractEvent<
      ContractRegisteredEvent.InputTuple,
      ContractRegisteredEvent.OutputTuple,
      ContractRegisteredEvent.OutputObject
    >;
    ContractRegistered: TypedContractEvent<
      ContractRegisteredEvent.InputTuple,
      ContractRegisteredEvent.OutputTuple,
      ContractRegisteredEvent.OutputObject
    >;

    "ContractStatusChanged(bytes)": TypedContractEvent<
      ContractStatusChangedEvent.InputTuple,
      ContractStatusChangedEvent.OutputTuple,
      ContractStatusChangedEvent.OutputObject
    >;
    ContractStatusChanged: TypedContractEvent<
      ContractStatusChangedEvent.InputTuple,
      ContractStatusChangedEvent.OutputTuple,
      ContractStatusChangedEvent.OutputObject
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

    "NewChainMetadata(uint256,string,bytes)": TypedContractEvent<
      NewChainMetadataEvent.InputTuple,
      NewChainMetadataEvent.OutputTuple,
      NewChainMetadataEvent.OutputObject
    >;
    NewChainMetadata: TypedContractEvent<
      NewChainMetadataEvent.InputTuple,
      NewChainMetadataEvent.OutputTuple,
      NewChainMetadataEvent.OutputObject
    >;

    "NewContractConfiguration(uint256,string,string,bytes)": TypedContractEvent<
      NewContractConfigurationEvent.InputTuple,
      NewContractConfigurationEvent.OutputTuple,
      NewContractConfigurationEvent.OutputObject
    >;
    NewContractConfiguration: TypedContractEvent<
      NewContractConfigurationEvent.InputTuple,
      NewContractConfigurationEvent.OutputTuple,
      NewContractConfigurationEvent.OutputObject
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

    "ZRC20TokenRegistered(bytes,address,uint8,uint256,string)": TypedContractEvent<
      ZRC20TokenRegisteredEvent.InputTuple,
      ZRC20TokenRegisteredEvent.OutputTuple,
      ZRC20TokenRegisteredEvent.OutputObject
    >;
    ZRC20TokenRegistered: TypedContractEvent<
      ZRC20TokenRegisteredEvent.InputTuple,
      ZRC20TokenRegisteredEvent.OutputTuple,
      ZRC20TokenRegisteredEvent.OutputObject
    >;

    "ZRC20TokenUpdated(address,bool)": TypedContractEvent<
      ZRC20TokenUpdatedEvent.InputTuple,
      ZRC20TokenUpdatedEvent.OutputTuple,
      ZRC20TokenUpdatedEvent.OutputObject
    >;
    ZRC20TokenUpdated: TypedContractEvent<
      ZRC20TokenUpdatedEvent.InputTuple,
      ZRC20TokenUpdatedEvent.OutputTuple,
      ZRC20TokenUpdatedEvent.OutputObject
    >;
  };
}
