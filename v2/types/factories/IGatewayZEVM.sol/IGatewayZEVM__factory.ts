/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  IGatewayZEVM,
  IGatewayZEVMInterface,
} from "../../IGatewayZEVM.sol/IGatewayZEVM";

const _abi = [
  {
    type: "function",
    name: "call",
    inputs: [
      {
        name: "receiver",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "zrc20",
        type: "address",
        internalType: "address",
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "gasLimit",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "revertOptions",
        type: "tuple",
        internalType: "struct RevertOptions",
        components: [
          {
            name: "revertAddress",
            type: "address",
            internalType: "address",
          },
          {
            name: "callOnRevert",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "abortAddress",
            type: "address",
            internalType: "address",
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
    name: "deposit",
    inputs: [
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
        name: "target",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "depositAndCall",
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
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "target",
        type: "address",
        internalType: "address",
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
    name: "depositAndCall",
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
        name: "target",
        type: "address",
        internalType: "address",
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
    name: "depositAndRevert",
    inputs: [
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
        name: "target",
        type: "address",
        internalType: "address",
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
      },
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
    type: "function",
    name: "execute",
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
        name: "target",
        type: "address",
        internalType: "address",
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
    name: "executeRevert",
    inputs: [
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
        name: "target",
        type: "address",
        internalType: "address",
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
      },
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
    type: "function",
    name: "withdraw",
    inputs: [
      {
        name: "receiver",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "zrc20",
        type: "address",
        internalType: "address",
      },
      {
        name: "revertOptions",
        type: "tuple",
        internalType: "struct RevertOptions",
        components: [
          {
            name: "revertAddress",
            type: "address",
            internalType: "address",
          },
          {
            name: "callOnRevert",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "abortAddress",
            type: "address",
            internalType: "address",
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
    name: "withdraw",
    inputs: [
      {
        name: "receiver",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "chainId",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "revertOptions",
        type: "tuple",
        internalType: "struct RevertOptions",
        components: [
          {
            name: "revertAddress",
            type: "address",
            internalType: "address",
          },
          {
            name: "callOnRevert",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "abortAddress",
            type: "address",
            internalType: "address",
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
    name: "withdrawAndCall",
    inputs: [
      {
        name: "receiver",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "chainId",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "revertOptions",
        type: "tuple",
        internalType: "struct RevertOptions",
        components: [
          {
            name: "revertAddress",
            type: "address",
            internalType: "address",
          },
          {
            name: "callOnRevert",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "abortAddress",
            type: "address",
            internalType: "address",
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
    name: "withdrawAndCall",
    inputs: [
      {
        name: "receiver",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "zrc20",
        type: "address",
        internalType: "address",
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "gasLimit",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "revertOptions",
        type: "tuple",
        internalType: "struct RevertOptions",
        components: [
          {
            name: "revertAddress",
            type: "address",
            internalType: "address",
          },
          {
            name: "callOnRevert",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "abortAddress",
            type: "address",
            internalType: "address",
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
    name: "Called",
    inputs: [
      {
        name: "sender",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "zrc20",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "receiver",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "message",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "revertOptions",
        type: "tuple",
        indexed: false,
        internalType: "struct RevertOptions",
        components: [
          {
            name: "revertAddress",
            type: "address",
            internalType: "address",
          },
          {
            name: "callOnRevert",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "abortAddress",
            type: "address",
            internalType: "address",
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
    name: "Withdrawn",
    inputs: [
      {
        name: "sender",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "chainId",
        type: "uint256",
        indexed: true,
        internalType: "uint256",
      },
      {
        name: "receiver",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "zrc20",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "value",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "gasfee",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "protocolFlatFee",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "message",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "revertOptions",
        type: "tuple",
        indexed: false,
        internalType: "struct RevertOptions",
        components: [
          {
            name: "revertAddress",
            type: "address",
            internalType: "address",
          },
          {
            name: "callOnRevert",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "abortAddress",
            type: "address",
            internalType: "address",
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
    type: "error",
    name: "CallerIsNotFungibleModule",
    inputs: [],
  },
  {
    type: "error",
    name: "EmptyMessage",
    inputs: [],
  },
  {
    type: "error",
    name: "FailedZetaSent",
    inputs: [],
  },
  {
    type: "error",
    name: "GasFeeTransferFailed",
    inputs: [],
  },
  {
    type: "error",
    name: "InsufficientZRC20Amount",
    inputs: [],
  },
  {
    type: "error",
    name: "InsufficientZetaAmount",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidTarget",
    inputs: [],
  },
  {
    type: "error",
    name: "OnlyWZETAOrFungible",
    inputs: [],
  },
  {
    type: "error",
    name: "WithdrawalFailed",
    inputs: [],
  },
  {
    type: "error",
    name: "ZRC20BurnFailed",
    inputs: [],
  },
  {
    type: "error",
    name: "ZRC20DepositFailed",
    inputs: [],
  },
  {
    type: "error",
    name: "ZRC20TransferFailed",
    inputs: [],
  },
] as const;

export class IGatewayZEVM__factory {
  static readonly abi = _abi;
  static createInterface(): IGatewayZEVMInterface {
    return new Interface(_abi) as IGatewayZEVMInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): IGatewayZEVM {
    return new Contract(address, _abi, runner) as unknown as IGatewayZEVM;
  }
}
