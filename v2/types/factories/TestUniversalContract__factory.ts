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
    name: "onCrossChainCall",
    inputs: [
      {
        name: "context",
        type: "tuple",
        internalType: "struct zContext",
        components: [
          {
            name: "origin",
            type: "bytes",
            internalType: "bytes",
          },
          {
            name: "sender",
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
        name: "zrc20",
        type: "address",
        internalType: "address",
      },
      {
        name: "amount",
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
            name: "asset",
            type: "address",
            internalType: "address",
          },
          {
            name: "amount",
            type: "uint64",
            internalType: "uint64",
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
    name: "ContextDataRevert",
    inputs: [
      {
        name: "revertContext",
        type: "tuple",
        indexed: false,
        internalType: "struct RevertContext",
        components: [
          {
            name: "asset",
            type: "address",
            internalType: "address",
          },
          {
            name: "amount",
            type: "uint64",
            internalType: "uint64",
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
  "0x6080604052348015600f57600080fd5b5061061e8061001f6000396000f3fe60806040526004361061002a5760003560e01c8063660b9de014610033578063de43156e1461005357005b3661003157005b005b34801561003f57600080fd5b5061003161004e366004610140565b610073565b34801561005f57600080fd5b5061003161006e3660046101a6565b6100ad565b7f35a9324413457251c1059312318f6f1cec6bd0da4105d01315f3151b1e3a2c76816040516100a291906102a7565b60405180910390a150565b606081156100c4576100c1828401846103a2565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e6100ef8780610498565b6100ff60408a0160208b01610504565b8960400135338660405161011896959493929190610526565b60405180910390a1505050505050565b60006060828403121561013a57600080fd5b50919050565b60006020828403121561015257600080fd5b813567ffffffffffffffff81111561016957600080fd5b61017584828501610128565b949350505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146101a157600080fd5b919050565b6000806000806000608086880312156101be57600080fd5b853567ffffffffffffffff8111156101d557600080fd5b6101e188828901610128565b9550506101f06020870161017d565b935060408601359250606086013567ffffffffffffffff81111561021357600080fd5b8601601f8101881361022457600080fd5b803567ffffffffffffffff81111561023b57600080fd5b88602082840101111561024d57600080fd5b959894975092955050506020019190565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6020815273ffffffffffffffffffffffffffffffffffffffff6102c98361017d565b1660208201526000602083013567ffffffffffffffff81168082146102ed57600080fd5b806040850152505060408301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261032a57600080fd5b830160208101903567ffffffffffffffff81111561034757600080fd5b80360382131561035657600080fd5b60608085015261036a60808501828461025e565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156103b457600080fd5b813567ffffffffffffffff8111156103cb57600080fd5b8201601f810184136103dc57600080fd5b803567ffffffffffffffff8111156103f6576103f6610373565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561046257610462610373565b60405281815282820160200186101561047a57600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126104cd57600080fd5b83018035915067ffffffffffffffff8211156104e857600080fd5b6020019150368190038213156104fd57600080fd5b9250929050565b60006020828403121561051657600080fd5b61051f8261017d565b9392505050565b60a08152600061053a60a08301888a61025e565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b818110156105a457602081870181015184830182015201610588565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168301019250505097965050505050505056fea2646970667358221220b870bc54e4f04a95cdf27ec6eaa16fae7263fbc0ff9617ef6a46abf0b42d747664736f6c634300081a0033";

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
