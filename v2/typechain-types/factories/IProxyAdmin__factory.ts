/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type { IProxyAdmin, IProxyAdminInterface } from "../IProxyAdmin";

const _abi = [
  {
    type: "function",
    name: "upgrade",
    inputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
      {
        name: "",
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
        name: "",
        type: "address",
        internalType: "address",
      },
      {
        name: "",
        type: "address",
        internalType: "address",
      },
      {
        name: "",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [],
    stateMutability: "payable",
  },
] as const;

export class IProxyAdmin__factory {
  static readonly abi = _abi;
  static createInterface(): IProxyAdminInterface {
    return new Interface(_abi) as IProxyAdminInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): IProxyAdmin {
    return new Contract(address, _abi, runner) as unknown as IProxyAdmin;
  }
}
