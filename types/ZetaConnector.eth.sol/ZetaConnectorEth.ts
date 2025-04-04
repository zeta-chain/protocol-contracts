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

export declare namespace ZetaInterfaces {
  export type SendInputStruct = {
    destinationChainId: BigNumberish;
    destinationAddress: BytesLike;
    destinationGasLimit: BigNumberish;
    message: BytesLike;
    zetaValueAndGas: BigNumberish;
    zetaParams: BytesLike;
  };

  export type SendInputStructOutput = [
    destinationChainId: bigint,
    destinationAddress: string,
    destinationGasLimit: bigint,
    message: string,
    zetaValueAndGas: bigint,
    zetaParams: string
  ] & {
    destinationChainId: bigint;
    destinationAddress: string;
    destinationGasLimit: bigint;
    message: string;
    zetaValueAndGas: bigint;
    zetaParams: string;
  };
}

export interface ZetaConnectorEthInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "getLockedAmount"
      | "onReceive"
      | "onRevert"
      | "pause"
      | "paused"
      | "pauserAddress"
      | "renounceTssAddressUpdater"
      | "send"
      | "tssAddress"
      | "tssAddressUpdater"
      | "unpause"
      | "updatePauserAddress"
      | "updateTssAddress"
      | "zetaToken"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic:
      | "Paused"
      | "PauserAddressUpdated"
      | "TSSAddressUpdated"
      | "TSSAddressUpdaterUpdated"
      | "Unpaused"
      | "ZetaReceived"
      | "ZetaReverted"
      | "ZetaSent"
  ): EventFragment;

  encodeFunctionData(
    functionFragment: "getLockedAmount",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "onReceive",
    values: [
      BytesLike,
      BigNumberish,
      AddressLike,
      BigNumberish,
      BytesLike,
      BytesLike
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "onRevert",
    values: [
      AddressLike,
      BigNumberish,
      BytesLike,
      BigNumberish,
      BigNumberish,
      BytesLike,
      BytesLike
    ]
  ): string;
  encodeFunctionData(functionFragment: "pause", values?: undefined): string;
  encodeFunctionData(functionFragment: "paused", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "pauserAddress",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "renounceTssAddressUpdater",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "send",
    values: [ZetaInterfaces.SendInputStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "tssAddress",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "tssAddressUpdater",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "unpause", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "updatePauserAddress",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "updateTssAddress",
    values: [AddressLike]
  ): string;
  encodeFunctionData(functionFragment: "zetaToken", values?: undefined): string;

  decodeFunctionResult(
    functionFragment: "getLockedAmount",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "onReceive", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "onRevert", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "pause", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "paused", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "pauserAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "renounceTssAddressUpdater",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "send", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "tssAddress", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "tssAddressUpdater",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "unpause", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "updatePauserAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "updateTssAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "zetaToken", data: BytesLike): Result;
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

export namespace PauserAddressUpdatedEvent {
  export type InputTuple = [
    callerAddress: AddressLike,
    newTssAddress: AddressLike
  ];
  export type OutputTuple = [callerAddress: string, newTssAddress: string];
  export interface OutputObject {
    callerAddress: string;
    newTssAddress: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace TSSAddressUpdatedEvent {
  export type InputTuple = [
    callerAddress: AddressLike,
    newTssAddress: AddressLike
  ];
  export type OutputTuple = [callerAddress: string, newTssAddress: string];
  export interface OutputObject {
    callerAddress: string;
    newTssAddress: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace TSSAddressUpdaterUpdatedEvent {
  export type InputTuple = [
    callerAddress: AddressLike,
    newTssUpdaterAddress: AddressLike
  ];
  export type OutputTuple = [
    callerAddress: string,
    newTssUpdaterAddress: string
  ];
  export interface OutputObject {
    callerAddress: string;
    newTssUpdaterAddress: string;
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

export namespace ZetaReceivedEvent {
  export type InputTuple = [
    zetaTxSenderAddress: BytesLike,
    sourceChainId: BigNumberish,
    destinationAddress: AddressLike,
    zetaValue: BigNumberish,
    message: BytesLike,
    internalSendHash: BytesLike
  ];
  export type OutputTuple = [
    zetaTxSenderAddress: string,
    sourceChainId: bigint,
    destinationAddress: string,
    zetaValue: bigint,
    message: string,
    internalSendHash: string
  ];
  export interface OutputObject {
    zetaTxSenderAddress: string;
    sourceChainId: bigint;
    destinationAddress: string;
    zetaValue: bigint;
    message: string;
    internalSendHash: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ZetaRevertedEvent {
  export type InputTuple = [
    zetaTxSenderAddress: AddressLike,
    sourceChainId: BigNumberish,
    destinationChainId: BigNumberish,
    destinationAddress: BytesLike,
    remainingZetaValue: BigNumberish,
    message: BytesLike,
    internalSendHash: BytesLike
  ];
  export type OutputTuple = [
    zetaTxSenderAddress: string,
    sourceChainId: bigint,
    destinationChainId: bigint,
    destinationAddress: string,
    remainingZetaValue: bigint,
    message: string,
    internalSendHash: string
  ];
  export interface OutputObject {
    zetaTxSenderAddress: string;
    sourceChainId: bigint;
    destinationChainId: bigint;
    destinationAddress: string;
    remainingZetaValue: bigint;
    message: string;
    internalSendHash: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ZetaSentEvent {
  export type InputTuple = [
    sourceTxOriginAddress: AddressLike,
    zetaTxSenderAddress: AddressLike,
    destinationChainId: BigNumberish,
    destinationAddress: BytesLike,
    zetaValueAndGas: BigNumberish,
    destinationGasLimit: BigNumberish,
    message: BytesLike,
    zetaParams: BytesLike
  ];
  export type OutputTuple = [
    sourceTxOriginAddress: string,
    zetaTxSenderAddress: string,
    destinationChainId: bigint,
    destinationAddress: string,
    zetaValueAndGas: bigint,
    destinationGasLimit: bigint,
    message: string,
    zetaParams: string
  ];
  export interface OutputObject {
    sourceTxOriginAddress: string;
    zetaTxSenderAddress: string;
    destinationChainId: bigint;
    destinationAddress: string;
    zetaValueAndGas: bigint;
    destinationGasLimit: bigint;
    message: string;
    zetaParams: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface ZetaConnectorEth extends BaseContract {
  connect(runner?: ContractRunner | null): ZetaConnectorEth;
  waitForDeployment(): Promise<this>;

  interface: ZetaConnectorEthInterface;

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

  getLockedAmount: TypedContractMethod<[], [bigint], "view">;

  onReceive: TypedContractMethod<
    [
      zetaTxSenderAddress: BytesLike,
      sourceChainId: BigNumberish,
      destinationAddress: AddressLike,
      zetaValue: BigNumberish,
      message: BytesLike,
      internalSendHash: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  onRevert: TypedContractMethod<
    [
      zetaTxSenderAddress: AddressLike,
      sourceChainId: BigNumberish,
      destinationAddress: BytesLike,
      destinationChainId: BigNumberish,
      remainingZetaValue: BigNumberish,
      message: BytesLike,
      internalSendHash: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  pause: TypedContractMethod<[], [void], "nonpayable">;

  paused: TypedContractMethod<[], [boolean], "view">;

  pauserAddress: TypedContractMethod<[], [string], "view">;

  renounceTssAddressUpdater: TypedContractMethod<[], [void], "nonpayable">;

  send: TypedContractMethod<
    [input: ZetaInterfaces.SendInputStruct],
    [void],
    "nonpayable"
  >;

  tssAddress: TypedContractMethod<[], [string], "view">;

  tssAddressUpdater: TypedContractMethod<[], [string], "view">;

  unpause: TypedContractMethod<[], [void], "nonpayable">;

  updatePauserAddress: TypedContractMethod<
    [pauserAddress_: AddressLike],
    [void],
    "nonpayable"
  >;

  updateTssAddress: TypedContractMethod<
    [tssAddress_: AddressLike],
    [void],
    "nonpayable"
  >;

  zetaToken: TypedContractMethod<[], [string], "view">;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "getLockedAmount"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "onReceive"
  ): TypedContractMethod<
    [
      zetaTxSenderAddress: BytesLike,
      sourceChainId: BigNumberish,
      destinationAddress: AddressLike,
      zetaValue: BigNumberish,
      message: BytesLike,
      internalSendHash: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onRevert"
  ): TypedContractMethod<
    [
      zetaTxSenderAddress: AddressLike,
      sourceChainId: BigNumberish,
      destinationAddress: BytesLike,
      destinationChainId: BigNumberish,
      remainingZetaValue: BigNumberish,
      message: BytesLike,
      internalSendHash: BytesLike
    ],
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
    nameOrSignature: "pauserAddress"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "renounceTssAddressUpdater"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "send"
  ): TypedContractMethod<
    [input: ZetaInterfaces.SendInputStruct],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "tssAddress"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "tssAddressUpdater"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "unpause"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "updatePauserAddress"
  ): TypedContractMethod<[pauserAddress_: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "updateTssAddress"
  ): TypedContractMethod<[tssAddress_: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "zetaToken"
  ): TypedContractMethod<[], [string], "view">;

  getEvent(
    key: "Paused"
  ): TypedContractEvent<
    PausedEvent.InputTuple,
    PausedEvent.OutputTuple,
    PausedEvent.OutputObject
  >;
  getEvent(
    key: "PauserAddressUpdated"
  ): TypedContractEvent<
    PauserAddressUpdatedEvent.InputTuple,
    PauserAddressUpdatedEvent.OutputTuple,
    PauserAddressUpdatedEvent.OutputObject
  >;
  getEvent(
    key: "TSSAddressUpdated"
  ): TypedContractEvent<
    TSSAddressUpdatedEvent.InputTuple,
    TSSAddressUpdatedEvent.OutputTuple,
    TSSAddressUpdatedEvent.OutputObject
  >;
  getEvent(
    key: "TSSAddressUpdaterUpdated"
  ): TypedContractEvent<
    TSSAddressUpdaterUpdatedEvent.InputTuple,
    TSSAddressUpdaterUpdatedEvent.OutputTuple,
    TSSAddressUpdaterUpdatedEvent.OutputObject
  >;
  getEvent(
    key: "Unpaused"
  ): TypedContractEvent<
    UnpausedEvent.InputTuple,
    UnpausedEvent.OutputTuple,
    UnpausedEvent.OutputObject
  >;
  getEvent(
    key: "ZetaReceived"
  ): TypedContractEvent<
    ZetaReceivedEvent.InputTuple,
    ZetaReceivedEvent.OutputTuple,
    ZetaReceivedEvent.OutputObject
  >;
  getEvent(
    key: "ZetaReverted"
  ): TypedContractEvent<
    ZetaRevertedEvent.InputTuple,
    ZetaRevertedEvent.OutputTuple,
    ZetaRevertedEvent.OutputObject
  >;
  getEvent(
    key: "ZetaSent"
  ): TypedContractEvent<
    ZetaSentEvent.InputTuple,
    ZetaSentEvent.OutputTuple,
    ZetaSentEvent.OutputObject
  >;

  filters: {
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

    "PauserAddressUpdated(address,address)": TypedContractEvent<
      PauserAddressUpdatedEvent.InputTuple,
      PauserAddressUpdatedEvent.OutputTuple,
      PauserAddressUpdatedEvent.OutputObject
    >;
    PauserAddressUpdated: TypedContractEvent<
      PauserAddressUpdatedEvent.InputTuple,
      PauserAddressUpdatedEvent.OutputTuple,
      PauserAddressUpdatedEvent.OutputObject
    >;

    "TSSAddressUpdated(address,address)": TypedContractEvent<
      TSSAddressUpdatedEvent.InputTuple,
      TSSAddressUpdatedEvent.OutputTuple,
      TSSAddressUpdatedEvent.OutputObject
    >;
    TSSAddressUpdated: TypedContractEvent<
      TSSAddressUpdatedEvent.InputTuple,
      TSSAddressUpdatedEvent.OutputTuple,
      TSSAddressUpdatedEvent.OutputObject
    >;

    "TSSAddressUpdaterUpdated(address,address)": TypedContractEvent<
      TSSAddressUpdaterUpdatedEvent.InputTuple,
      TSSAddressUpdaterUpdatedEvent.OutputTuple,
      TSSAddressUpdaterUpdatedEvent.OutputObject
    >;
    TSSAddressUpdaterUpdated: TypedContractEvent<
      TSSAddressUpdaterUpdatedEvent.InputTuple,
      TSSAddressUpdaterUpdatedEvent.OutputTuple,
      TSSAddressUpdaterUpdatedEvent.OutputObject
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

    "ZetaReceived(bytes,uint256,address,uint256,bytes,bytes32)": TypedContractEvent<
      ZetaReceivedEvent.InputTuple,
      ZetaReceivedEvent.OutputTuple,
      ZetaReceivedEvent.OutputObject
    >;
    ZetaReceived: TypedContractEvent<
      ZetaReceivedEvent.InputTuple,
      ZetaReceivedEvent.OutputTuple,
      ZetaReceivedEvent.OutputObject
    >;

    "ZetaReverted(address,uint256,uint256,bytes,uint256,bytes,bytes32)": TypedContractEvent<
      ZetaRevertedEvent.InputTuple,
      ZetaRevertedEvent.OutputTuple,
      ZetaRevertedEvent.OutputObject
    >;
    ZetaReverted: TypedContractEvent<
      ZetaRevertedEvent.InputTuple,
      ZetaRevertedEvent.OutputTuple,
      ZetaRevertedEvent.OutputObject
    >;

    "ZetaSent(address,address,uint256,bytes,uint256,uint256,bytes,bytes)": TypedContractEvent<
      ZetaSentEvent.InputTuple,
      ZetaSentEvent.OutputTuple,
      ZetaSentEvent.OutputObject
    >;
    ZetaSent: TypedContractEvent<
      ZetaSentEvent.InputTuple,
      ZetaSentEvent.OutputTuple,
      ZetaSentEvent.OutputObject
    >;
  };
}
