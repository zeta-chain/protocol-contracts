/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  IERC20CustodyErrors,
  IERC20CustodyErrorsInterface,
} from "../../IERC20Custody.sol/IERC20CustodyErrors";

const _abi = [
  {
    type: "error",
    name: "LegacyMethodsNotSupported",
    inputs: [],
  },
  {
    type: "error",
    name: "NotWhitelisted",
    inputs: [],
  },
  {
    type: "error",
    name: "ZeroAddress",
    inputs: [],
  },
] as const;

export class IERC20CustodyErrors__factory {
  static readonly abi = _abi;
  static createInterface(): IERC20CustodyErrorsInterface {
    return new Interface(_abi) as IERC20CustodyErrorsInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): IERC20CustodyErrors {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as IERC20CustodyErrors;
  }
}
