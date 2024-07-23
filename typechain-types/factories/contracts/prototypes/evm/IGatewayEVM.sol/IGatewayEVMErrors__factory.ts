/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  IGatewayEVMErrors,
  IGatewayEVMErrorsInterface,
} from "../../../../../contracts/prototypes/evm/IGatewayEVM.sol/IGatewayEVMErrors";

const _abi = [
  {
    inputs: [],
    name: "ApprovalFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "CustodyInitialized",
    type: "error",
  },
  {
    inputs: [],
    name: "DepositFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "ExecutionFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "InsufficientERC20Amount",
    type: "error",
  },
  {
    inputs: [],
    name: "InsufficientETHAmount",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidSender",
    type: "error",
  },
  {
    inputs: [],
    name: "ZeroAddress",
    type: "error",
  },
] as const;

export class IGatewayEVMErrors__factory {
  static readonly abi = _abi;
  static createInterface(): IGatewayEVMErrorsInterface {
    return new utils.Interface(_abi) as IGatewayEVMErrorsInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IGatewayEVMErrors {
    return new Contract(address, _abi, signerOrProvider) as IGatewayEVMErrors;
  }
}
