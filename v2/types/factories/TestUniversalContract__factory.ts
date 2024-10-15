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
    name: "onCall",
    inputs: [
      {
        name: "context",
        type: "tuple",
        internalType: "struct MessageContext",
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
  "0x6080604052348015600f57600080fd5b5061061e8061001f6000396000f3fe60806040526004361061002a5760003560e01c80635bcfd61614610033578063660b9de01461005357005b3661003157005b005b34801561003f57600080fd5b5061003161004e366004610169565b610073565b34801561005f57600080fd5b5061003161006e366004610221565b6100ee565b6060811561008a576100878284018461028d565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e6100b58780610383565b6100c560408a0160208b016103ef565b896040013533866040516100de9695949392919061045a565b60405180910390a1505050505050565b7f35a9324413457251c1059312318f6f1cec6bd0da4105d01315f3151b1e3a2c768160405161011d919061051c565b60405180910390a150565b60006060828403121561013a57600080fd5b50919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461016457600080fd5b919050565b60008060008060006080868803121561018157600080fd5b853567ffffffffffffffff81111561019857600080fd5b6101a488828901610128565b9550506101b360208701610140565b935060408601359250606086013567ffffffffffffffff8111156101d657600080fd5b8601601f810188136101e757600080fd5b803567ffffffffffffffff8111156101fe57600080fd5b88602082840101111561021057600080fd5b959894975092955050506020019190565b60006020828403121561023357600080fd5b813567ffffffffffffffff81111561024a57600080fd5b61025684828501610128565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561029f57600080fd5b813567ffffffffffffffff8111156102b657600080fd5b8201601f810184136102c757600080fd5b803567ffffffffffffffff8111156102e1576102e161025e565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561034d5761034d61025e565b60405281815282820160200186101561036557600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126103b857600080fd5b83018035915067ffffffffffffffff8211156103d357600080fd5b6020019150368190038213156103e857600080fd5b9250929050565b60006020828403121561040157600080fd5b61040a82610140565b9392505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60a08152600061046e60a08301888a610411565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b818110156104d8576020818701810151848301820152016104bc565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff61053e83610140565b1660208201526000602083013567ffffffffffffffff811680821461056257600080fd5b806040850152505060408301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261059f57600080fd5b830160208101903567ffffffffffffffff8111156105bc57600080fd5b8036038213156105cb57600080fd5b6060808501526105df608085018284610411565b9594505050505056fea2646970667358221220078ed23f1e4df13b41aee8ed02311eda755ac8a9fe6e1a9a5736af6396fa30f964736f6c634300081a0033";

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
