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
  AddressLike,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedListener,
  TypedContractMethod,
} from "../common";

export declare namespace ZetaInterfaces {
  export type ZetaMessageStruct = {
    zetaTxSenderAddress: BytesLike;
    sourceChainId: BigNumberish;
    destinationAddress: AddressLike;
    zetaValue: BigNumberish;
    message: BytesLike;
  };

  export type ZetaMessageStructOutput = [
    zetaTxSenderAddress: string,
    sourceChainId: bigint,
    destinationAddress: string,
    zetaValue: bigint,
    message: string
  ] & {
    zetaTxSenderAddress: string;
    sourceChainId: bigint;
    destinationAddress: string;
    zetaValue: bigint;
    message: string;
  };

  export type ZetaRevertStruct = {
    zetaTxSenderAddress: AddressLike;
    sourceChainId: BigNumberish;
    destinationAddress: BytesLike;
    destinationChainId: BigNumberish;
    remainingZetaValue: BigNumberish;
    message: BytesLike;
  };

  export type ZetaRevertStructOutput = [
    zetaTxSenderAddress: string,
    sourceChainId: bigint,
    destinationAddress: string,
    destinationChainId: bigint,
    remainingZetaValue: bigint,
    message: string
  ] & {
    zetaTxSenderAddress: string;
    sourceChainId: bigint;
    destinationAddress: string;
    destinationChainId: bigint;
    remainingZetaValue: bigint;
    message: string;
  };
}

export interface ZetaReceiverInterface extends Interface {
  getFunction(
    nameOrSignature: "onZetaMessage" | "onZetaRevert"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "onZetaMessage",
    values: [ZetaInterfaces.ZetaMessageStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "onZetaRevert",
    values: [ZetaInterfaces.ZetaRevertStruct]
  ): string;

  decodeFunctionResult(
    functionFragment: "onZetaMessage",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onZetaRevert",
    data: BytesLike
  ): Result;
}

export interface ZetaReceiver extends BaseContract {
  connect(runner?: ContractRunner | null): ZetaReceiver;
  waitForDeployment(): Promise<this>;

  interface: ZetaReceiverInterface;

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

  onZetaMessage: TypedContractMethod<
    [zetaMessage: ZetaInterfaces.ZetaMessageStruct],
    [void],
    "nonpayable"
  >;

  onZetaRevert: TypedContractMethod<
    [zetaRevert: ZetaInterfaces.ZetaRevertStruct],
    [void],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "onZetaMessage"
  ): TypedContractMethod<
    [zetaMessage: ZetaInterfaces.ZetaMessageStruct],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onZetaRevert"
  ): TypedContractMethod<
    [zetaRevert: ZetaInterfaces.ZetaRevertStruct],
    [void],
    "nonpayable"
  >;

  filters: {};
}