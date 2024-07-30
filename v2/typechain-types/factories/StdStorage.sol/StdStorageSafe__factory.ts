/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Contract,
  ContractFactory,
  ContractTransactionResponse,
  Interface,
} from "ethers";
import type { Signer, ContractDeployTransaction, ContractRunner } from "ethers";
import type { NonPayableOverrides } from "../../common";
import type {
  StdStorageSafe,
  StdStorageSafeInterface,
} from "../../StdStorage.sol/StdStorageSafe";

const _abi = [
  {
    type: "event",
    name: "SlotFound",
    inputs: [
      {
        name: "who",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "fsig",
        type: "bytes4",
        indexed: false,
        internalType: "bytes4",
      },
      {
        name: "keysHash",
        type: "bytes32",
        indexed: false,
        internalType: "bytes32",
      },
      {
        name: "slot",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "WARNING_UninitedSlot",
    inputs: [
      {
        name: "who",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "slot",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
] as const;

const _bytecode =
  "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d6473a3e455f11b00943e19d62d02cf957dec4d27a7a7ddb023bf40264d580b064736f6c634300081a0033";

type StdStorageSafeConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: StdStorageSafeConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class StdStorageSafe__factory extends ContractFactory {
  constructor(...args: StdStorageSafeConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(overrides || {});
  }
  override deploy(overrides?: NonPayableOverrides & { from?: string }) {
    return super.deploy(overrides || {}) as Promise<
      StdStorageSafe & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): StdStorageSafe__factory {
    return super.connect(runner) as StdStorageSafe__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): StdStorageSafeInterface {
    return new Interface(_abi) as StdStorageSafeInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): StdStorageSafe {
    return new Contract(address, _abi, runner) as unknown as StdStorageSafe;
  }
}
