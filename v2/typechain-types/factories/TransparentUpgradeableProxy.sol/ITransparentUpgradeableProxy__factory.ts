/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  ITransparentUpgradeableProxy,
  ITransparentUpgradeableProxyInterface,
} from "../../TransparentUpgradeableProxy.sol/ITransparentUpgradeableProxy";

const _abi = [
  {
    type: "function",
    name: "upgradeToAndCall",
    inputs: [
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
  {
    type: "event",
    name: "AdminChanged",
    inputs: [
      {
        name: "previousAdmin",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "newAdmin",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "BeaconUpgraded",
    inputs: [
      {
        name: "beacon",
        type: "address",
        indexed: true,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Upgraded",
    inputs: [
      {
        name: "implementation",
        type: "address",
        indexed: true,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
] as const;

export class ITransparentUpgradeableProxy__factory {
  static readonly abi = _abi;
  static createInterface(): ITransparentUpgradeableProxyInterface {
    return new Interface(_abi) as ITransparentUpgradeableProxyInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): ITransparentUpgradeableProxy {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as ITransparentUpgradeableProxy;
  }
}
