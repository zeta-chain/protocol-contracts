/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
  PayableOverrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type {
  FunctionFragment,
  Result,
  EventFragment,
} from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../../../common";

export type ZContextStruct = {
  origin: PromiseOrValue<BytesLike>;
  sender: PromiseOrValue<string>;
  chainID: PromiseOrValue<BigNumberish>;
};

export type ZContextStructOutput = [string, string, BigNumber] & {
  origin: string;
  sender: string;
  chainID: BigNumber;
};

export interface GatewayZEVMInterface extends utils.Interface {
  functions: {
    "FUNGIBLE_MODULE_ADDRESS()": FunctionFragment;
    "call(bytes,bytes)": FunctionFragment;
    "deposit(address,uint256,address)": FunctionFragment;
    "depositAndCall((bytes,address,uint256),uint256,address,bytes)": FunctionFragment;
    "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)": FunctionFragment;
    "execute((bytes,address,uint256),address,uint256,address,bytes)": FunctionFragment;
    "initialize(address)": FunctionFragment;
    "owner()": FunctionFragment;
    "proxiableUUID()": FunctionFragment;
    "renounceOwnership()": FunctionFragment;
    "transferOwnership(address)": FunctionFragment;
    "upgradeTo(address)": FunctionFragment;
    "upgradeToAndCall(address,bytes)": FunctionFragment;
    "withdraw(bytes,uint256,address)": FunctionFragment;
    "withdraw(uint256)": FunctionFragment;
    "withdrawAndCall(uint256,bytes)": FunctionFragment;
    "withdrawAndCall(bytes,uint256,address,bytes)": FunctionFragment;
    "wzeta()": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "FUNGIBLE_MODULE_ADDRESS"
      | "call"
      | "deposit"
      | "depositAndCall((bytes,address,uint256),uint256,address,bytes)"
      | "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)"
      | "execute"
      | "initialize"
      | "owner"
      | "proxiableUUID"
      | "renounceOwnership"
      | "transferOwnership"
      | "upgradeTo"
      | "upgradeToAndCall"
      | "withdraw(bytes,uint256,address)"
      | "withdraw(uint256)"
      | "withdrawAndCall(uint256,bytes)"
      | "withdrawAndCall(bytes,uint256,address,bytes)"
      | "wzeta"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "FUNGIBLE_MODULE_ADDRESS",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "call",
    values: [PromiseOrValue<BytesLike>, PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "deposit",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "depositAndCall((bytes,address,uint256),uint256,address,bytes)",
    values: [
      ZContextStruct,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)",
    values: [
      ZContextStruct,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "execute",
    values: [
      ZContextStruct,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "initialize",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(functionFragment: "owner", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "proxiableUUID",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "renounceOwnership",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "transferOwnership",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "upgradeTo",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "upgradeToAndCall",
    values: [PromiseOrValue<string>, PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "withdraw(bytes,uint256,address)",
    values: [
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "withdraw(uint256)",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawAndCall(uint256,bytes)",
    values: [PromiseOrValue<BigNumberish>, PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawAndCall(bytes,uint256,address,bytes)",
    values: [
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(functionFragment: "wzeta", values?: undefined): string;

  decodeFunctionResult(
    functionFragment: "FUNGIBLE_MODULE_ADDRESS",
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
  decodeFunctionResult(functionFragment: "execute", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "initialize", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "owner", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "proxiableUUID",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "renounceOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "transferOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "upgradeTo", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "upgradeToAndCall",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdraw(bytes,uint256,address)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdraw(uint256)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawAndCall(uint256,bytes)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawAndCall(bytes,uint256,address,bytes)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "wzeta", data: BytesLike): Result;

  events: {
    "AdminChanged(address,address)": EventFragment;
    "BeaconUpgraded(address)": EventFragment;
    "Call(address,bytes,bytes)": EventFragment;
    "Initialized(uint8)": EventFragment;
    "OwnershipTransferred(address,address)": EventFragment;
    "Upgraded(address)": EventFragment;
    "Withdrawal(address,address,bytes,uint256,uint256,uint256,bytes)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "AdminChanged"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "BeaconUpgraded"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Call"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Initialized"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "OwnershipTransferred"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Upgraded"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Withdrawal"): EventFragment;
}

export interface AdminChangedEventObject {
  previousAdmin: string;
  newAdmin: string;
}
export type AdminChangedEvent = TypedEvent<
  [string, string],
  AdminChangedEventObject
>;

export type AdminChangedEventFilter = TypedEventFilter<AdminChangedEvent>;

export interface BeaconUpgradedEventObject {
  beacon: string;
}
export type BeaconUpgradedEvent = TypedEvent<
  [string],
  BeaconUpgradedEventObject
>;

export type BeaconUpgradedEventFilter = TypedEventFilter<BeaconUpgradedEvent>;

export interface CallEventObject {
  sender: string;
  receiver: string;
  message: string;
}
export type CallEvent = TypedEvent<[string, string, string], CallEventObject>;

export type CallEventFilter = TypedEventFilter<CallEvent>;

export interface InitializedEventObject {
  version: number;
}
export type InitializedEvent = TypedEvent<[number], InitializedEventObject>;

export type InitializedEventFilter = TypedEventFilter<InitializedEvent>;

export interface OwnershipTransferredEventObject {
  previousOwner: string;
  newOwner: string;
}
export type OwnershipTransferredEvent = TypedEvent<
  [string, string],
  OwnershipTransferredEventObject
>;

export type OwnershipTransferredEventFilter =
  TypedEventFilter<OwnershipTransferredEvent>;

export interface UpgradedEventObject {
  implementation: string;
}
export type UpgradedEvent = TypedEvent<[string], UpgradedEventObject>;

export type UpgradedEventFilter = TypedEventFilter<UpgradedEvent>;

export interface WithdrawalEventObject {
  from: string;
  zrc20: string;
  to: string;
  value: BigNumber;
  gasfee: BigNumber;
  protocolFlatFee: BigNumber;
  message: string;
}
export type WithdrawalEvent = TypedEvent<
  [string, string, string, BigNumber, BigNumber, BigNumber, string],
  WithdrawalEventObject
>;

export type WithdrawalEventFilter = TypedEventFilter<WithdrawalEvent>;

export interface GatewayZEVM extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: GatewayZEVMInterface;

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>;

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>;
  listeners(eventName?: string): Array<Listener>;
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this;
  removeAllListeners(eventName?: string): this;
  off: OnEvent<this>;
  on: OnEvent<this>;
  once: OnEvent<this>;
  removeListener: OnEvent<this>;

  functions: {
    FUNGIBLE_MODULE_ADDRESS(overrides?: CallOverrides): Promise<[string]>;

    call(
      receiver: PromiseOrValue<BytesLike>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    deposit(
      zrc20: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "depositAndCall((bytes,address,uint256),uint256,address,bytes)"(
      context: ZContextStruct,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)"(
      context: ZContextStruct,
      zrc20: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    execute(
      context: ZContextStruct,
      zrc20: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    initialize(
      _wzeta: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    owner(overrides?: CallOverrides): Promise<[string]>;

    proxiableUUID(overrides?: CallOverrides): Promise<[string]>;

    renounceOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    upgradeTo(
      newImplementation: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    upgradeToAndCall(
      newImplementation: PromiseOrValue<string>,
      data: PromiseOrValue<BytesLike>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "withdraw(bytes,uint256,address)"(
      receiver: PromiseOrValue<BytesLike>,
      amount: PromiseOrValue<BigNumberish>,
      zrc20: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "withdraw(uint256)"(
      amount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "withdrawAndCall(uint256,bytes)"(
      amount: PromiseOrValue<BigNumberish>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "withdrawAndCall(bytes,uint256,address,bytes)"(
      receiver: PromiseOrValue<BytesLike>,
      amount: PromiseOrValue<BigNumberish>,
      zrc20: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    wzeta(overrides?: CallOverrides): Promise<[string]>;
  };

  FUNGIBLE_MODULE_ADDRESS(overrides?: CallOverrides): Promise<string>;

  call(
    receiver: PromiseOrValue<BytesLike>,
    message: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  deposit(
    zrc20: PromiseOrValue<string>,
    amount: PromiseOrValue<BigNumberish>,
    target: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "depositAndCall((bytes,address,uint256),uint256,address,bytes)"(
    context: ZContextStruct,
    amount: PromiseOrValue<BigNumberish>,
    target: PromiseOrValue<string>,
    message: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)"(
    context: ZContextStruct,
    zrc20: PromiseOrValue<string>,
    amount: PromiseOrValue<BigNumberish>,
    target: PromiseOrValue<string>,
    message: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  execute(
    context: ZContextStruct,
    zrc20: PromiseOrValue<string>,
    amount: PromiseOrValue<BigNumberish>,
    target: PromiseOrValue<string>,
    message: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  initialize(
    _wzeta: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  owner(overrides?: CallOverrides): Promise<string>;

  proxiableUUID(overrides?: CallOverrides): Promise<string>;

  renounceOwnership(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  transferOwnership(
    newOwner: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  upgradeTo(
    newImplementation: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  upgradeToAndCall(
    newImplementation: PromiseOrValue<string>,
    data: PromiseOrValue<BytesLike>,
    overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "withdraw(bytes,uint256,address)"(
    receiver: PromiseOrValue<BytesLike>,
    amount: PromiseOrValue<BigNumberish>,
    zrc20: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "withdraw(uint256)"(
    amount: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "withdrawAndCall(uint256,bytes)"(
    amount: PromiseOrValue<BigNumberish>,
    message: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "withdrawAndCall(bytes,uint256,address,bytes)"(
    receiver: PromiseOrValue<BytesLike>,
    amount: PromiseOrValue<BigNumberish>,
    zrc20: PromiseOrValue<string>,
    message: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  wzeta(overrides?: CallOverrides): Promise<string>;

  callStatic: {
    FUNGIBLE_MODULE_ADDRESS(overrides?: CallOverrides): Promise<string>;

    call(
      receiver: PromiseOrValue<BytesLike>,
      message: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    deposit(
      zrc20: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    "depositAndCall((bytes,address,uint256),uint256,address,bytes)"(
      context: ZContextStruct,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)"(
      context: ZContextStruct,
      zrc20: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    execute(
      context: ZContextStruct,
      zrc20: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    initialize(
      _wzeta: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    owner(overrides?: CallOverrides): Promise<string>;

    proxiableUUID(overrides?: CallOverrides): Promise<string>;

    renounceOwnership(overrides?: CallOverrides): Promise<void>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    upgradeTo(
      newImplementation: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    upgradeToAndCall(
      newImplementation: PromiseOrValue<string>,
      data: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    "withdraw(bytes,uint256,address)"(
      receiver: PromiseOrValue<BytesLike>,
      amount: PromiseOrValue<BigNumberish>,
      zrc20: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    "withdraw(uint256)"(
      amount: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    "withdrawAndCall(uint256,bytes)"(
      amount: PromiseOrValue<BigNumberish>,
      message: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    "withdrawAndCall(bytes,uint256,address,bytes)"(
      receiver: PromiseOrValue<BytesLike>,
      amount: PromiseOrValue<BigNumberish>,
      zrc20: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    wzeta(overrides?: CallOverrides): Promise<string>;
  };

  filters: {
    "AdminChanged(address,address)"(
      previousAdmin?: null,
      newAdmin?: null
    ): AdminChangedEventFilter;
    AdminChanged(
      previousAdmin?: null,
      newAdmin?: null
    ): AdminChangedEventFilter;

    "BeaconUpgraded(address)"(
      beacon?: PromiseOrValue<string> | null
    ): BeaconUpgradedEventFilter;
    BeaconUpgraded(
      beacon?: PromiseOrValue<string> | null
    ): BeaconUpgradedEventFilter;

    "Call(address,bytes,bytes)"(
      sender?: PromiseOrValue<string> | null,
      receiver?: null,
      message?: null
    ): CallEventFilter;
    Call(
      sender?: PromiseOrValue<string> | null,
      receiver?: null,
      message?: null
    ): CallEventFilter;

    "Initialized(uint8)"(version?: null): InitializedEventFilter;
    Initialized(version?: null): InitializedEventFilter;

    "OwnershipTransferred(address,address)"(
      previousOwner?: PromiseOrValue<string> | null,
      newOwner?: PromiseOrValue<string> | null
    ): OwnershipTransferredEventFilter;
    OwnershipTransferred(
      previousOwner?: PromiseOrValue<string> | null,
      newOwner?: PromiseOrValue<string> | null
    ): OwnershipTransferredEventFilter;

    "Upgraded(address)"(
      implementation?: PromiseOrValue<string> | null
    ): UpgradedEventFilter;
    Upgraded(
      implementation?: PromiseOrValue<string> | null
    ): UpgradedEventFilter;

    "Withdrawal(address,address,bytes,uint256,uint256,uint256,bytes)"(
      from?: PromiseOrValue<string> | null,
      zrc20?: null,
      to?: null,
      value?: null,
      gasfee?: null,
      protocolFlatFee?: null,
      message?: null
    ): WithdrawalEventFilter;
    Withdrawal(
      from?: PromiseOrValue<string> | null,
      zrc20?: null,
      to?: null,
      value?: null,
      gasfee?: null,
      protocolFlatFee?: null,
      message?: null
    ): WithdrawalEventFilter;
  };

  estimateGas: {
    FUNGIBLE_MODULE_ADDRESS(overrides?: CallOverrides): Promise<BigNumber>;

    call(
      receiver: PromiseOrValue<BytesLike>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    deposit(
      zrc20: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "depositAndCall((bytes,address,uint256),uint256,address,bytes)"(
      context: ZContextStruct,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)"(
      context: ZContextStruct,
      zrc20: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    execute(
      context: ZContextStruct,
      zrc20: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    initialize(
      _wzeta: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    owner(overrides?: CallOverrides): Promise<BigNumber>;

    proxiableUUID(overrides?: CallOverrides): Promise<BigNumber>;

    renounceOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    upgradeTo(
      newImplementation: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    upgradeToAndCall(
      newImplementation: PromiseOrValue<string>,
      data: PromiseOrValue<BytesLike>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "withdraw(bytes,uint256,address)"(
      receiver: PromiseOrValue<BytesLike>,
      amount: PromiseOrValue<BigNumberish>,
      zrc20: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "withdraw(uint256)"(
      amount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "withdrawAndCall(uint256,bytes)"(
      amount: PromiseOrValue<BigNumberish>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "withdrawAndCall(bytes,uint256,address,bytes)"(
      receiver: PromiseOrValue<BytesLike>,
      amount: PromiseOrValue<BigNumberish>,
      zrc20: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    wzeta(overrides?: CallOverrides): Promise<BigNumber>;
  };

  populateTransaction: {
    FUNGIBLE_MODULE_ADDRESS(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    call(
      receiver: PromiseOrValue<BytesLike>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    deposit(
      zrc20: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "depositAndCall((bytes,address,uint256),uint256,address,bytes)"(
      context: ZContextStruct,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "depositAndCall((bytes,address,uint256),address,uint256,address,bytes)"(
      context: ZContextStruct,
      zrc20: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    execute(
      context: ZContextStruct,
      zrc20: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      target: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    initialize(
      _wzeta: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    owner(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    proxiableUUID(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    renounceOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    upgradeTo(
      newImplementation: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    upgradeToAndCall(
      newImplementation: PromiseOrValue<string>,
      data: PromiseOrValue<BytesLike>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "withdraw(bytes,uint256,address)"(
      receiver: PromiseOrValue<BytesLike>,
      amount: PromiseOrValue<BigNumberish>,
      zrc20: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "withdraw(uint256)"(
      amount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "withdrawAndCall(uint256,bytes)"(
      amount: PromiseOrValue<BigNumberish>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "withdrawAndCall(bytes,uint256,address,bytes)"(
      receiver: PromiseOrValue<BytesLike>,
      amount: PromiseOrValue<BigNumberish>,
      zrc20: PromiseOrValue<string>,
      message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    wzeta(overrides?: CallOverrides): Promise<PopulatedTransaction>;
  };
}
