/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Contract,
  ContractFactory,
  ContractTransactionResponse,
  Interface,
} from "ethers";
import type {
  Signer,
  AddressLike,
  ContractDeployTransaction,
  ContractRunner,
} from "ethers";
import type { NonPayableOverrides } from "../common";
import type {
  UpgradeableBeacon,
  UpgradeableBeaconInterface,
} from "../UpgradeableBeacon";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "implementation_",
        type: "address",
        internalType: "address",
      },
      {
        name: "initialOwner",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "implementation",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "owner",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "renounceOwnership",
    inputs: [],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "transferOwnership",
    inputs: [
      {
        name: "newOwner",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "upgradeTo",
    inputs: [
      {
        name: "newImplementation",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    name: "OwnershipTransferred",
    inputs: [
      {
        name: "previousOwner",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "newOwner",
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
  {
    type: "error",
    name: "BeaconInvalidImplementation",
    inputs: [
      {
        name: "implementation",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "OwnableInvalidOwner",
    inputs: [
      {
        name: "owner",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "OwnableUnauthorizedAccount",
    inputs: [
      {
        name: "account",
        type: "address",
        internalType: "address",
      },
    ],
  },
] as const;

const _bytecode =
  "0x608060405234801561001057600080fd5b5060405161054538038061054583398101604081905261002f91610165565b806001600160a01b03811661005f57604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b61006881610079565b50610072826100c9565b5050610198565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b806001600160a01b03163b6000036100ff5760405163211eb15960e21b81526001600160a01b0382166004820152602401610056565b600180546001600160a01b0319166001600160a01b0383169081179091556040517fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b80516001600160a01b038116811461016057600080fd5b919050565b6000806040838503121561017857600080fd5b61018183610149565b915061018f60208401610149565b90509250929050565b61039e806101a76000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063715018a611610050578063715018a6146100c45780638da5cb5b146100cc578063f2fde38b146100ea57600080fd5b80633659cfe61461006c5780635c60da1b14610081575b600080fd5b61007f61007a36600461032b565b6100fd565b005b60015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b61007f610111565b60005473ffffffffffffffffffffffffffffffffffffffff1661009b565b61007f6100f836600461032b565b610125565b61010561018b565b61010e816101de565b50565b61011961018b565b61012360006102b6565b565b61012d61018b565b73ffffffffffffffffffffffffffffffffffffffff8116610182576040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600060048201526024015b60405180910390fd5b61010e816102b6565b60005473ffffffffffffffffffffffffffffffffffffffff163314610123576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610179565b8073ffffffffffffffffffffffffffffffffffffffff163b600003610247576040517f847ac56400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610179565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561033d57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461036157600080fd5b939250505056fea2646970667358221220f63f0ff7fc3e813a85ced034f789b78328d46922533a9c25f261edb53a279d9164736f6c634300081a0033";

type UpgradeableBeaconConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: UpgradeableBeaconConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class UpgradeableBeacon__factory extends ContractFactory {
  constructor(...args: UpgradeableBeaconConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    implementation_: AddressLike,
    initialOwner: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(
      implementation_,
      initialOwner,
      overrides || {}
    );
  }
  override deploy(
    implementation_: AddressLike,
    initialOwner: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(
      implementation_,
      initialOwner,
      overrides || {}
    ) as Promise<
      UpgradeableBeacon & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): UpgradeableBeacon__factory {
    return super.connect(runner) as UpgradeableBeacon__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): UpgradeableBeaconInterface {
    return new Interface(_abi) as UpgradeableBeaconInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): UpgradeableBeacon {
    return new Contract(address, _abi, runner) as unknown as UpgradeableBeacon;
  }
}
