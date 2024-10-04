/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  Revertable,
  RevertableInterface,
} from "../../Revert.sol/Revertable";

const _abi = [
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
] as const;

export class Revertable__factory {
  static readonly abi = _abi;
  static createInterface(): RevertableInterface {
    return new Interface(_abi) as RevertableInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): Revertable {
    return new Contract(address, _abi, runner) as unknown as Revertable;
  }
}
