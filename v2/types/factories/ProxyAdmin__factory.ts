/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Contract,
  ContractFactory,
  ContractTransactionResponse,
  Interface,
} from "ethers";
import type {
  Signer,
  AddressLike,
  ContractDeployTransaction,
  ContractRunner,
} from "ethers";
import type { NonPayableOverrides } from "../common";
import type { ProxyAdmin, ProxyAdminInterface } from "../ProxyAdmin";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "initialOwner",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "UPGRADE_INTERFACE_VERSION",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "string",
        internalType: "string",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "owner",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "renounceOwnership",
    inputs: [],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "transferOwnership",
    inputs: [
      {
        name: "newOwner",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "upgradeAndCall",
    inputs: [
      {
        name: "proxy",
        type: "address",
        internalType: "contract ITransparentUpgradeableProxy",
      },
      {
        name: "implementation",
        type: "address",
        internalType: "address",
      },
      {
        name: "data",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [],
    stateMutability: "payable",
  },
  {
    type: "event",
    name: "OwnershipTransferred",
    inputs: [
      {
        name: "previousOwner",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "newOwner",
        type: "address",
        indexed: true,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "error",
    name: "OwnableInvalidOwner",
    inputs: [
      {
        name: "owner",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "OwnableUnauthorizedAccount",
    inputs: [
      {
        name: "account",
        type: "address",
        internalType: "address",
      },
    ],
  },
] as const;

const _bytecode =
  "0x608060405234801561001057600080fd5b5060405161068438038061068483398101604081905261002f916100be565b806001600160a01b03811661005e57604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b6100678161006e565b50506100ee565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100d057600080fd5b81516001600160a01b03811681146100e757600080fd5b9392505050565b610587806100fd6000396000f3fe60806040526004361061005a5760003560e01c80639623609d116100435780639623609d146100b0578063ad3cb1cc146100c3578063f2fde38b1461011957600080fd5b8063715018a61461005f5780638da5cb5b14610076575b600080fd5b34801561006b57600080fd5b50610074610139565b005b34801561008257600080fd5b5060005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100746100be366004610364565b61014d565b3480156100cf57600080fd5b5061010c6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516100a791906104e3565b34801561012557600080fd5b506100746101343660046104fd565b6101e2565b61014161024b565b61014b600061029e565b565b61015561024b565b6040517f4f1ef28600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841690634f1ef2869034906101ab908690869060040161051a565b6000604051808303818588803b1580156101c457600080fd5b505af11580156101d8573d6000803e3d6000fd5b5050505050505050565b6101ea61024b565b73ffffffffffffffffffffffffffffffffffffffff811661023f576040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600060048201526024015b60405180910390fd5b6102488161029e565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461014b576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610236565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b73ffffffffffffffffffffffffffffffffffffffff8116811461024857600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060006060848603121561037957600080fd5b833561038481610313565b9250602084013561039481610313565b9150604084013567ffffffffffffffff8111156103b057600080fd5b8401601f810186136103c157600080fd5b803567ffffffffffffffff8111156103db576103db610335565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561044757610447610335565b60405281815282820160200188101561045f57600080fd5b816020840160208301376000602083830101528093505050509250925092565b6000815180845260005b818110156104a557602081850181015186830182015201610489565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006104f6602083018461047f565b9392505050565b60006020828403121561050f57600080fd5b81356104f681610313565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000610549604083018461047f565b94935050505056fea264697066735822122041909b73d43bb8b3f4b5beaf6ad1718632d58cdb8dbaffd76f7a2077399a4da564736f6c634300081a0033";

type ProxyAdminConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: ProxyAdminConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class ProxyAdmin__factory extends ContractFactory {
  constructor(...args: ProxyAdminConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    initialOwner: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(initialOwner, overrides || {});
  }
  override deploy(
    initialOwner: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(initialOwner, overrides || {}) as Promise<
      ProxyAdmin & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): ProxyAdmin__factory {
    return super.connect(runner) as ProxyAdmin__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ProxyAdminInterface {
    return new Interface(_abi) as ProxyAdminInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): ProxyAdmin {
    return new Contract(address, _abi, runner) as unknown as ProxyAdmin;
  }
}
