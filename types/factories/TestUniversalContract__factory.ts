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
<<<<<<< HEAD
  "0x6080604052348015600f57600080fd5b5061061b8061001f6000396000f3fe60806040526004361061002a5760003560e01c80635bcfd61614610033578063c9028a361461005357005b3661003157005b005b34801561003f57600080fd5b5061003161004e366004610151565b610073565b34801561005f57600080fd5b5061003161006e36600461020e565b6100ee565b6060811561008a576100878284018461027f565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e6100b58780610375565b6100c560408a0160208b016103e1565b896040013533866040516100de96959493929190610445565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c48160405161011d9190610507565b60405180910390a150565b803573ffffffffffffffffffffffffffffffffffffffff8116811461014c57600080fd5b919050565b60008060008060006080868803121561016957600080fd5b853567ffffffffffffffff81111561018057600080fd5b86016060818903121561019257600080fd5b94506101a060208701610128565b935060408601359250606086013567ffffffffffffffff8111156101c357600080fd5b8601601f810188136101d457600080fd5b803567ffffffffffffffff8111156101eb57600080fd5b8860208284010111156101fd57600080fd5b959894975092955050506020019190565b60006020828403121561022057600080fd5b813567ffffffffffffffff81111561023757600080fd5b82016080818503121561024957600080fd5b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561029157600080fd5b813567ffffffffffffffff8111156102a857600080fd5b8201601f810184136102b957600080fd5b803567ffffffffffffffff8111156102d3576102d3610250565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561033f5761033f610250565b60405281815282820160200186101561035757600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126103aa57600080fd5b83018035915067ffffffffffffffff8211156103c557600080fd5b6020019150368190038213156103da57600080fd5b9250929050565b6000602082840312156103f357600080fd5b61024982610128565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60a08152600061045960a08301888a6103fc565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b818110156104c3576020818701810151848301820152016104a7565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff61052983610128565b16602082015273ffffffffffffffffffffffffffffffffffffffff61055060208401610128565b166040820152600080604084013590508060608401525060608301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261059c57600080fd5b830160208101903567ffffffffffffffff8111156105b957600080fd5b8036038213156105c857600080fd5b6080808501526105dc60a0850182846103fc565b9594505050505056fea2646970667358221220d7ad1e086ca92690817a447fa5d15e6f90c25918a56022bc0d74b5f2c1c50c9a64736f6c634300081a0033";
=======
  "0x6080604052348015600f57600080fd5b5061061b8061001f6000396000f3fe60806040526004361061002a5760003560e01c80635bcfd61614610033578063c9028a361461005357005b3661003157005b005b34801561003f57600080fd5b5061003161004e366004610151565b610073565b34801561005f57600080fd5b5061003161006e36600461020e565b6100ee565b6060811561008a576100878284018461027f565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e6100b58780610375565b6100c560408a0160208b016103e1565b896040013533866040516100de96959493929190610445565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c48160405161011d9190610507565b60405180910390a150565b803573ffffffffffffffffffffffffffffffffffffffff8116811461014c57600080fd5b919050565b60008060008060006080868803121561016957600080fd5b853567ffffffffffffffff81111561018057600080fd5b86016060818903121561019257600080fd5b94506101a060208701610128565b935060408601359250606086013567ffffffffffffffff8111156101c357600080fd5b8601601f810188136101d457600080fd5b803567ffffffffffffffff8111156101eb57600080fd5b8860208284010111156101fd57600080fd5b959894975092955050506020019190565b60006020828403121561022057600080fd5b813567ffffffffffffffff81111561023757600080fd5b82016080818503121561024957600080fd5b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561029157600080fd5b813567ffffffffffffffff8111156102a857600080fd5b8201601f810184136102b957600080fd5b803567ffffffffffffffff8111156102d3576102d3610250565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561033f5761033f610250565b60405281815282820160200186101561035757600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126103aa57600080fd5b83018035915067ffffffffffffffff8211156103c557600080fd5b6020019150368190038213156103da57600080fd5b9250929050565b6000602082840312156103f357600080fd5b61024982610128565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60a08152600061045960a08301888a6103fc565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b818110156104c3576020818701810151848301820152016104a7565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff61052983610128565b16602082015273ffffffffffffffffffffffffffffffffffffffff61055060208401610128565b166040820152600080604084013590508060608401525060608301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261059c57600080fd5b830160208101903567ffffffffffffffff8111156105b957600080fd5b8036038213156105c857600080fd5b6080808501526105dc60a0850182846103fc565b9594505050505056fea264697066735822122007ea8d901f533f9926890d22eee0d3ddb82dde38edc6131f7a85024fe5d8d02864736f6c634300081a0033";
>>>>>>> main

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
