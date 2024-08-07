/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  SystemContractErrors,
  SystemContractErrorsInterface,
} from "../../SystemContractMock.sol/SystemContractErrors";

const _abi = [
  {
    type: "error",
    name: "CallerIsNotFungibleModule",
    inputs: [],
  },
  {
    type: "error",
    name: "CantBeIdenticalAddresses",
    inputs: [],
  },
  {
    type: "error",
    name: "CantBeZeroAddress",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidTarget",
    inputs: [],
  },
] as const;

export class SystemContractErrors__factory {
  static readonly abi = _abi;
  static createInterface(): SystemContractErrorsInterface {
    return new Interface(_abi) as SystemContractErrorsInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): SystemContractErrors {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as SystemContractErrors;
  }
}
