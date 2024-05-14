/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  ZetaTokenConsumerZEVMErrors,
  ZetaTokenConsumerZEVMErrorsInterface,
} from "../../../../../contracts/evm/tools/ZetaTokenConsumerZEVM.strategy.sol/ZetaTokenConsumerZEVMErrors";

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
    name: "InputCantBeZeta",
    type: "error",
  },
  {
    inputs: [],
    name: "NotEnoughValue",
    type: "error",
  },
  {
    inputs: [],
    name: "OnlyWZETAAllowed",
    type: "error",
  },
  {
    inputs: [],
    name: "OutputCantBeZeta",
    type: "error",
  },
  {
    inputs: [],
    name: "ReentrancyError",
    type: "error",
  },
] as const;

export class ZetaTokenConsumerZEVMErrors__factory {
  static readonly abi = _abi;
  static createInterface(): ZetaTokenConsumerZEVMErrorsInterface {
    return new utils.Interface(_abi) as ZetaTokenConsumerZEVMErrorsInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ZetaTokenConsumerZEVMErrors {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as ZetaTokenConsumerZEVMErrors;
  }
}