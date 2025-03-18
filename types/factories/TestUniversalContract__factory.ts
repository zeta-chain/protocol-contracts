/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Contract,
  ContractFactory,
  ContractTransactionResponse,
  Interface,
} from "ethers";
import type { Signer, ContractDeployTransaction, ContractRunner } from "ethers";
import type { NonPayableOverrides } from "../common";
import type {
  TestUniversalContract,
  TestUniversalContractInterface,
} from "../TestUniversalContract";

const _abi = [
  {
    type: "fallback",
    stateMutability: "payable",
  },
  {
    type: "receive",
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "onAbort",
    inputs: [
      {
        name: "abortContext",
        type: "tuple",
        internalType: "struct AbortContext",
        components: [
          {
            name: "sender",
            type: "bytes",
            internalType: "bytes",
          },
          {
            name: "asset",
            type: "address",
            internalType: "address",
          },
          {
            name: "amount",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "outgoing",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "chainID",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "revertMessage",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "onCall",
    inputs: [
      {
        name: "context",
        type: "tuple",
        internalType: "struct MessageContext",
        components: [
          {
            name: "sender",
            type: "bytes",
            internalType: "bytes",
          },
          {
            name: "senderEVM",
            type: "address",
            internalType: "address",
          },
          {
            name: "chainID",
            type: "uint256",
            internalType: "uint256",
          },
        ],
      },
      {
        name: "",
        type: "address",
        internalType: "address",
      },
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "onRevert",
    inputs: [
      {
        name: "revertContext",
        type: "tuple",
        internalType: "struct RevertContext",
        components: [
          {
            name: "sender",
            type: "address",
            internalType: "address",
          },
          {
            name: "asset",
            type: "address",
            internalType: "address",
          },
          {
            name: "amount",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "revertMessage",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    name: "ContextData",
    inputs: [
      {
        name: "origin",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "sender",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "chainID",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "msgSender",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "message",
        type: "string",
        indexed: false,
        internalType: "string",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ContextDataAbort",
    inputs: [
      {
        name: "abortContext",
        type: "tuple",
        indexed: false,
        internalType: "struct AbortContext",
        components: [
          {
            name: "sender",
            type: "bytes",
            internalType: "bytes",
          },
          {
            name: "asset",
            type: "address",
            internalType: "address",
          },
          {
            name: "amount",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "outgoing",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "chainID",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "revertMessage",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ContextDataRevert",
    inputs: [
      {
        name: "revertContext",
        type: "tuple",
        indexed: false,
        internalType: "struct RevertContext",
        components: [
          {
            name: "sender",
            type: "address",
            internalType: "address",
          },
          {
            name: "asset",
            type: "address",
            internalType: "address",
          },
          {
            name: "amount",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "revertMessage",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
    ],
    anonymous: false,
  },
] as const;

const _bytecode =
  "0x6080604052348015600f57600080fd5b506107978061001f6000396000f3fe6080604052600436106100355760003560e01c80632d4cfb7e1461003e5780635bcfd6161461005e578063c9028a361461007e57005b3661003c57005b005b34801561004a57600080fd5b5061003c610059366004610182565b61009e565b34801561006a57600080fd5b5061003c6100793660046101ed565b6100d8565b34801561008a57600080fd5b5061003c6100993660046102aa565b610153565b7f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7816040516100cd9190610399565b60405180910390a150565b606081156100ef576100ec8284018461049f565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e61011a8780610595565b61012a60408a0160208b016105fa565b8960400135338660405161014396959493929190610615565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4816040516100cd91906106d7565b60006020828403121561019457600080fd5b813567ffffffffffffffff8111156101ab57600080fd5b820160c081850312156101bd57600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146101e857600080fd5b919050565b60008060008060006080868803121561020557600080fd5b853567ffffffffffffffff81111561021c57600080fd5b86016060818903121561022e57600080fd5b945061023c602087016101c4565b935060408601359250606086013567ffffffffffffffff81111561025f57600080fd5b8601601f8101881361027057600080fd5b803567ffffffffffffffff81111561028757600080fd5b88602082840101111561029957600080fd5b959894975092955050506020019190565b6000602082840312156102bc57600080fd5b813567ffffffffffffffff8111156102d357600080fd5b8201608081850312156101bd57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261031a57600080fd5b830160208101925035905067ffffffffffffffff81111561033a57600080fd5b80360382131561034957600080fd5b9250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6020815260006103a983846102e5565b60c060208501526103be60e085018284610350565b91505073ffffffffffffffffffffffffffffffffffffffff6103e2602086016101c4565b16604084015260006040850135905080606085015250606084013580151580821461040c57600080fd5b80608086015250506000608085013590508060a08501525061043160a08501856102e5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c0860152610466838284610350565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156104b157600080fd5b813567ffffffffffffffff8111156104c857600080fd5b8201601f810184136104d957600080fd5b803567ffffffffffffffff8111156104f3576104f3610470565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561055f5761055f610470565b60405281815282820160200186101561057757600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126105ca57600080fd5b83018035915067ffffffffffffffff8211156105e557600080fd5b60200191503681900382131561034957600080fd5b60006020828403121561060c57600080fd5b6101bd826101c4565b60a08152600061062960a08301888a610350565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b8181101561069357602081870181015184830182015201610677565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff6106f9836101c4565b16602082015273ffffffffffffffffffffffffffffffffffffffff610720602084016101c4565b166040820152600080604084013590508060608401525061074460608401846102e5565b60808085015261075860a085018284610350565b9594505050505056fea2646970667358221220121f02e9d998ca4c5df318c6995b5a93c9477c42c6e34b1b54eafd621fd74daf64736f6c634300081a0033";

type TestUniversalContractConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: TestUniversalContractConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class TestUniversalContract__factory extends ContractFactory {
  constructor(...args: TestUniversalContractConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(overrides || {});
  }
  override deploy(overrides?: NonPayableOverrides & { from?: string }) {
    return super.deploy(overrides || {}) as Promise<
      TestUniversalContract & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(
    runner: ContractRunner | null
  ): TestUniversalContract__factory {
    return super.connect(runner) as TestUniversalContract__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): TestUniversalContractInterface {
    return new Interface(_abi) as TestUniversalContractInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): TestUniversalContract {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as TestUniversalContract;
  }
}
