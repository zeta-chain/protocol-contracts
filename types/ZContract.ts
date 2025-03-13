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

export interface ZContractInterface extends Interface {
  getFunction(nameOrSignature: "onCrossChainCall"): FunctionFragment;

  encodeFunctionData(
    functionFragment: "onCrossChainCall",
    values: [ZContextStruct, AddressLike, BigNumberish, BytesLike]
  ): string;

  decodeFunctionResult(
    functionFragment: "onCrossChainCall",
    data: BytesLike
  ): Result;
}

export interface ZContract extends BaseContract {
  connect(runner?: ContractRunner | null): ZContract;
  waitForDeployment(): Promise<this>;

  interface: ZContractInterface;

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

  onCrossChainCall: TypedContractMethod<
    [
      context: ZContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "onCrossChainCall"
  ): TypedContractMethod<
    [
      context: ZContextStruct,
      zrc20: AddressLike,
      amount: BigNumberish,
      message: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  filters: {};
}
