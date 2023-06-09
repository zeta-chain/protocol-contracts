/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  ZetaTokenConsumerTridentErrors,
  ZetaTokenConsumerTridentErrorsInterface,
} from "../../../../../contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol/ZetaTokenConsumerTridentErrors";

const _abi = [
  {
    inputs: [],
    name: "ErrorSendingETH",
    type: "error",
  },
  {
    inputs: [],
    name: "InputCantBeZero",
    type: "error",
  },
  {
    inputs: [],
    name: "ReentrancyError",
    type: "error",
  },
] as const;

export class ZetaTokenConsumerTridentErrors__factory {
  static readonly abi = _abi;
  static createInterface(): ZetaTokenConsumerTridentErrorsInterface {
    return new utils.Interface(_abi) as ZetaTokenConsumerTridentErrorsInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ZetaTokenConsumerTridentErrors {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as ZetaTokenConsumerTridentErrors;
  }
}
