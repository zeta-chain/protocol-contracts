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

export type AbortContextStruct = {
  sender: BytesLike;
  asset: AddressLike;
  amount: BigNumberish;
  outgoing: boolean;
  chainID: BigNumberish;
  revertMessage: BytesLike;
};

export type AbortContextStructOutput = [
  sender: string,
  asset: string,
  amount: bigint,
  outgoing: boolean,
  chainID: bigint,
  revertMessage: string
] & {
  sender: string;
  asset: string;
  amount: bigint;
  outgoing: boolean;
  chainID: bigint;
  revertMessage: string;
};

export interface AbortableInterface extends Interface {
  getFunction(nameOrSignature: "onAbort"): FunctionFragment;

  encodeFunctionData(
    functionFragment: "onAbort",
    values: [AbortContextStruct]
  ): string;

  decodeFunctionResult(functionFragment: "onAbort", data: BytesLike): Result;
}

export interface Abortable extends BaseContract {
  connect(runner?: ContractRunner | null): Abortable;
  waitForDeployment(): Promise<this>;

  interface: AbortableInterface;

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

  onAbort: TypedContractMethod<
    [abortContext: AbortContextStruct],
    [void],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "onAbort"
  ): TypedContractMethod<
    [abortContext: AbortContextStruct],
    [void],
    "nonpayable"
  >;

  filters: {};
}
