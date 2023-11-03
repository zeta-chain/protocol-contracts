/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  SystemContractErrors,
  SystemContractErrorsInterface,
} from "../../../../../contracts/zevm/testing/MockSystemContract.sol/SystemContractErrors";

const _abi = [
  {
    inputs: [],
    name: "CallerIsNotFungibleModule",
    type: "error",
  },
  {
    inputs: [],
    name: "CantBeIdenticalAddresses",
    type: "error",
  },
  {
    inputs: [],
    name: "CantBeZeroAddress",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidTarget",
    type: "error",
  },
] as const;

export class SystemContractErrors__factory {
  static readonly abi = _abi;
  static createInterface(): SystemContractErrorsInterface {
    return new utils.Interface(_abi) as SystemContractErrorsInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): SystemContractErrors {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as SystemContractErrors;
  }
}