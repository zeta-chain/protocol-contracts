/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  IGatewayEVM,
  IGatewayEVMInterface,
} from "../../IGatewayEVM.sol/IGatewayEVM";

const _abi = [
  {
    type: "function",
    name: "call",
    inputs: [
      {
        name: "receiver",
        type: "address",
        internalType: "address",
      },
      {
        name: "payload",
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
          {
            name: "onRevertGasLimit",
            type: "uint256",
            internalType: "uint256",
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
        name: "receiver",
        type: "address",
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "asset",
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
          {
            name: "onRevertGasLimit",
            type: "uint256",
            internalType: "uint256",
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
        name: "receiver",
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
          {
            name: "onRevertGasLimit",
            type: "uint256",
            internalType: "uint256",
          },
        ],
      },
    ],
    outputs: [],
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "depositAndCall",
    inputs: [
      {
        name: "receiver",
        type: "address",
        internalType: "address",
      },
      {
        name: "payload",
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
          {
            name: "onRevertGasLimit",
            type: "uint256",
            internalType: "uint256",
          },
        ],
      },
    ],
    outputs: [],
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "depositAndCall",
    inputs: [
      {
        name: "receiver",
        type: "address",
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "asset",
        type: "address",
        internalType: "address",
      },
      {
        name: "payload",
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
          {
            name: "onRevertGasLimit",
            type: "uint256",
            internalType: "uint256",
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
        name: "destination",
        type: "address",
        internalType: "address",
      },
      {
        name: "data",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "executeRevert",
    inputs: [
      {
        name: "destination",
        type: "address",
        internalType: "address",
      },
      {
        name: "data",
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
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "executeWithERC20",
    inputs: [
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
      {
        name: "to",
        type: "address",
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "data",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "revertWithERC20",
    inputs: [
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
      {
        name: "to",
        type: "address",
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "data",
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
        name: "receiver",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "payload",
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
          {
            name: "onRevertGasLimit",
            type: "uint256",
            internalType: "uint256",
          },
        ],
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Deposited",
    inputs: [
      {
        name: "sender",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "receiver",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "asset",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "payload",
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
          {
            name: "onRevertGasLimit",
            type: "uint256",
            internalType: "uint256",
          },
        ],
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Executed",
    inputs: [
      {
        name: "destination",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "value",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "data",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ExecutedWithERC20",
    inputs: [
      {
        name: "token",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "to",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "data",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Reverted",
    inputs: [
      {
        name: "to",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "token",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "data",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
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
  {
    type: "event",
    name: "UpdatedGatewayTSSAddress",
    inputs: [
      {
        name: "oldTSSAddress",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "newTSSAddress",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "error",
    name: "ApprovalFailed",
    inputs: [],
  },
  {
    type: "error",
    name: "ConnectorInitialized",
    inputs: [],
  },
  {
    type: "error",
    name: "CustodyInitialized",
    inputs: [],
  },
  {
    type: "error",
    name: "DepositFailed",
    inputs: [],
  },
  {
    type: "error",
    name: "ExecutionFailed",
    inputs: [],
  },
  {
    type: "error",
    name: "InsufficientERC20Amount",
    inputs: [],
  },
  {
    type: "error",
    name: "InsufficientETHAmount",
    inputs: [],
  },
  {
    type: "error",
    name: "NotAllowedToCallOnRevert",
    inputs: [],
  },
  {
    type: "error",
    name: "NotWhitelistedInCustody",
    inputs: [],
  },
  {
    type: "error",
    name: "PayloadSizeExceeded",
    inputs: [],
  },
  {
    type: "error",
    name: "ZeroAddress",
    inputs: [],
  },
] as const;

export class IGatewayEVM__factory {
  static readonly abi = _abi;
  static createInterface(): IGatewayEVMInterface {
    return new Interface(_abi) as IGatewayEVMInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): IGatewayEVM {
    return new Contract(address, _abi, runner) as unknown as IGatewayEVM;
  }
}
