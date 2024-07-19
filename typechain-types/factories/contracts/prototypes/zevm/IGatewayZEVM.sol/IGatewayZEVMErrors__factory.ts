/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  IGatewayZEVMErrors,
  IGatewayZEVMErrorsInterface,
} from "../../../../../contracts/prototypes/zevm/IGatewayZEVM.sol/IGatewayZEVMErrors";

const _abi = [
  {
    inputs: [],
    name: "CallerIsNotFungibleModule",
    type: "error",
  },
  {
    inputs: [],
    name: "FailedZetaSent",
    type: "error",
  },
  {
    inputs: [],
    name: "GasFeeTransferFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "InsufficientZRC20Amount",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidTarget",
    type: "error",
  },
  {
    inputs: [],
    name: "WithdrawalFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "ZRC20BurnFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "ZRC20TransferFailed",
    type: "error",
  },
] as const;

export class IGatewayZEVMErrors__factory {
  static readonly abi = _abi;
  static createInterface(): IGatewayZEVMErrorsInterface {
    return new utils.Interface(_abi) as IGatewayZEVMErrorsInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IGatewayZEVMErrors {
    return new Contract(address, _abi, signerOrProvider) as IGatewayZEVMErrors;
  }
}
