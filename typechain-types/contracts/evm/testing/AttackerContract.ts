/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type { FunctionFragment, Result } from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../../../common";

export interface AttackerContractInterface extends utils.Interface {
  functions: {
    "attack(bytes4,bytes)": FunctionFragment;
    "entered()": FunctionFragment;
    "set(bytes4,bytes)": FunctionFragment;
    "victimContractAddress()": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "attack"
      | "entered"
      | "set"
      | "victimContractAddress"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "attack",
    values: [PromiseOrValue<BytesLike>, PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(functionFragment: "entered", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "set",
    values: [PromiseOrValue<BytesLike>, PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "victimContractAddress",
    values?: undefined
  ): string;

  decodeFunctionResult(functionFragment: "attack", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "entered", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "set", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "victimContractAddress",
    data: BytesLike
  ): Result;

  events: {};
}

export interface AttackerContract extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: AttackerContractInterface;

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
    attack(
      _methodSignature: PromiseOrValue<BytesLike>,
      _params: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    entered(overrides?: CallOverrides): Promise<[boolean]>;

    set(
      _methodSignature: PromiseOrValue<BytesLike>,
      _params: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    victimContractAddress(overrides?: CallOverrides): Promise<[string]>;
  };

  attack(
    _methodSignature: PromiseOrValue<BytesLike>,
    _params: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  entered(overrides?: CallOverrides): Promise<boolean>;

  set(
    _methodSignature: PromiseOrValue<BytesLike>,
    _params: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  victimContractAddress(overrides?: CallOverrides): Promise<string>;

  callStatic: {
    attack(
      _methodSignature: PromiseOrValue<BytesLike>,
      _params: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    entered(overrides?: CallOverrides): Promise<boolean>;

    set(
      _methodSignature: PromiseOrValue<BytesLike>,
      _params: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    victimContractAddress(overrides?: CallOverrides): Promise<string>;
  };

  filters: {};

  estimateGas: {
    attack(
      _methodSignature: PromiseOrValue<BytesLike>,
      _params: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    entered(overrides?: CallOverrides): Promise<BigNumber>;

    set(
      _methodSignature: PromiseOrValue<BytesLike>,
      _params: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    victimContractAddress(overrides?: CallOverrides): Promise<BigNumber>;
  };

  populateTransaction: {
    attack(
      _methodSignature: PromiseOrValue<BytesLike>,
      _params: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    entered(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    set(
      _methodSignature: PromiseOrValue<BytesLike>,
      _params: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    victimContractAddress(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}