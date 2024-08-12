/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  IGatewayZEVMEvents,
  IGatewayZEVMEventsInterface,
} from "../../IGatewayZEVM.sol/IGatewayZEVMEvents";

const _abi = [
  {
    type: "event",
    name: "Call",
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
    name: "Withdrawal",
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
] as const;

export class IGatewayZEVMEvents__factory {
  static readonly abi = _abi;
  static createInterface(): IGatewayZEVMEventsInterface {
    return new Interface(_abi) as IGatewayZEVMEventsInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): IGatewayZEVMEvents {
    return new Contract(address, _abi, runner) as unknown as IGatewayZEVMEvents;
  }
}
