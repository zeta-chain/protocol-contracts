/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  IRegistry,
  IRegistryInterface,
} from "../../IRegistry.sol/IRegistry";

const _abi = [
  {
    type: "event",
    name: "ChainMetadataUpdated",
    inputs: [
      {
        name: "chainId",
        type: "uint256",
        indexed: true,
        internalType: "uint256",
      },
      {
        name: "key",
        type: "string",
        indexed: false,
        internalType: "string",
      },
      {
        name: "value",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ChainStatusChanged",
    inputs: [
      {
        name: "chainId",
        type: "uint256",
        indexed: true,
        internalType: "uint256",
      },
      {
        name: "newState",
        type: "bool",
        indexed: false,
        internalType: "bool",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ContractConfigurationUpdated",
    inputs: [
      {
        name: "chainId",
        type: "uint256",
        indexed: true,
        internalType: "uint256",
      },
      {
        name: "contractType",
        type: "string",
        indexed: false,
        internalType: "string",
      },
      {
        name: "key",
        type: "string",
        indexed: false,
        internalType: "string",
      },
      {
        name: "value",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ContractRegistered",
    inputs: [
      {
        name: "chainId",
        type: "uint256",
        indexed: true,
        internalType: "uint256",
      },
      {
        name: "contractType",
        type: "string",
        indexed: false,
        internalType: "string",
      },
      {
        name: "addressBytes",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ContractStatusChanged",
    inputs: [
      {
        name: "addressBytes",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ZRC20TokenRegistered",
    inputs: [
      {
        name: "originAddress",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "zrc20",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "decimals",
        type: "uint8",
        indexed: false,
        internalType: "uint8",
      },
      {
        name: "originChainId",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "symbol",
        type: "string",
        indexed: false,
        internalType: "string",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ZRC20TokenUpdated",
    inputs: [
      {
        name: "zrc20",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "active",
        type: "bool",
        indexed: false,
        internalType: "bool",
      },
    ],
    anonymous: false,
  },
  {
    type: "error",
    name: "ChainActive",
    inputs: [
      {
        name: "chainId",
        type: "uint256",
        internalType: "uint256",
      },
    ],
  },
  {
    type: "error",
    name: "ChainNonActive",
    inputs: [
      {
        name: "chainId",
        type: "uint256",
        internalType: "uint256",
      },
    ],
  },
  {
    type: "error",
    name: "ContractAlreadyRegistered",
    inputs: [
      {
        name: "chainId",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "contractType",
        type: "string",
        internalType: "string",
      },
      {
        name: "addressBytes",
        type: "bytes",
        internalType: "bytes",
      },
    ],
  },
  {
    type: "error",
    name: "ContractNotFound",
    inputs: [
      {
        name: "chainId",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "contractType",
        type: "string",
        internalType: "string",
      },
    ],
  },
  {
    type: "error",
    name: "InvalidContractType",
    inputs: [
      {
        name: "contractType",
        type: "string",
        internalType: "string",
      },
    ],
  },
  {
    type: "error",
    name: "InvalidSender",
    inputs: [],
  },
  {
    type: "error",
    name: "NotAuthorized",
    inputs: [
      {
        name: "caller",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "NotGateway",
    inputs: [
      {
        name: "caller",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "ZRC20AlreadyRegistered",
    inputs: [
      {
        name: "zrc20",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "ZRC20SymbolAlreadyInUse",
    inputs: [
      {
        name: "symbol",
        type: "string",
        internalType: "string",
      },
    ],
  },
  {
    type: "error",
    name: "ZeroAddress",
    inputs: [],
  },
] as const;

export class IRegistry__factory {
  static readonly abi = _abi;
  static createInterface(): IRegistryInterface {
    return new Interface(_abi) as IRegistryInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): IRegistry {
    return new Contract(address, _abi, runner) as unknown as IRegistry;
  }
}
