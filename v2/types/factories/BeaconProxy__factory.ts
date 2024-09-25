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
  BytesLike,
  AddressLike,
  ContractDeployTransaction,
  ContractRunner,
} from "ethers";
import type { PayableOverrides } from "../common";
import type { BeaconProxy, BeaconProxyInterface } from "../BeaconProxy";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "beacon",
        type: "address",
        internalType: "address",
      },
      {
        name: "data",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "payable",
  },
  {
    type: "fallback",
    stateMutability: "payable",
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
    type: "error",
    name: "AddressEmptyCode",
    inputs: [
      {
        name: "target",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "ERC1967InvalidBeacon",
    inputs: [
      {
        name: "beacon",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "ERC1967InvalidImplementation",
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
    name: "ERC1967NonPayable",
    inputs: [],
  },
  {
    type: "error",
    name: "FailedInnerCall",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x60a06040526040516105eb3803806105eb83398101604081905261002291610387565b61002c828261003e565b506001600160a01b0316608052610484565b610047826100fe565b6040516001600160a01b038316907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e90600090a28051156100f2576100ed826001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e7919061044d565b82610211565b505050565b6100fa610288565b5050565b806001600160a01b03163b60000361013957604051631933b43b60e21b81526001600160a01b03821660048201526024015b60405180910390fd5b807fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5080546001600160a01b0319166001600160a01b0392831617905560408051635c60da1b60e01b81529051600092841691635c60da1b9160048083019260209291908290030181865afa1580156101b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d9919061044d565b9050806001600160a01b03163b6000036100fa57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610130565b6060600080846001600160a01b03168460405161022e9190610468565b600060405180830381855af49150503d8060008114610269576040519150601f19603f3d011682016040523d82523d6000602084013e61026e565b606091505b50909250905061027f8583836102a9565b95945050505050565b34156102a75760405163b398979f60e01b815260040160405180910390fd5b565b6060826102be576102b982610308565b610301565b81511580156102d557506001600160a01b0384163b155b156102fe57604051639996b31560e01b81526001600160a01b0385166004820152602401610130565b50805b9392505050565b8051156103185780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461034857600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561037e578181015183820152602001610366565b50506000910152565b6000806040838503121561039a57600080fd5b6103a383610331565b60208401519092506001600160401b038111156103bf57600080fd5b8301601f810185136103d057600080fd5b80516001600160401b038111156103e9576103e961034d565b604051601f8201601f19908116603f011681016001600160401b03811182821017156104175761041761034d565b60405281815282820160200187101561042f57600080fd5b610440826020830160208601610363565b8093505050509250929050565b60006020828403121561045f57600080fd5b61030182610331565b6000825161047a818460208701610363565b9190910192915050565b60805161014d61049e60003960006024015261014d6000f3fe608060405261000c61000e565b005b61001e610019610020565b6100b6565b565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561008d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100b191906100da565b905090565b3660008037600080366000845af43d6000803e8080156100d5573d6000f35b3d6000fd5b6000602082840312156100ec57600080fd5b815173ffffffffffffffffffffffffffffffffffffffff8116811461011057600080fd5b939250505056fea264697066735822122007d5a38c7d9b9a96ddfdd1b5e6c30f4ad5976d823e12ec481ebec20e06f430af64736f6c634300081a0033";

type BeaconProxyConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: BeaconProxyConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class BeaconProxy__factory extends ContractFactory {
  constructor(...args: BeaconProxyConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    beacon: AddressLike,
    data: BytesLike,
    overrides?: PayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(beacon, data, overrides || {});
  }
  override deploy(
    beacon: AddressLike,
    data: BytesLike,
    overrides?: PayableOverrides & { from?: string }
  ) {
    return super.deploy(beacon, data, overrides || {}) as Promise<
      BeaconProxy & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): BeaconProxy__factory {
    return super.connect(runner) as BeaconProxy__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): BeaconProxyInterface {
    return new Interface(_abi) as BeaconProxyInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): BeaconProxy {
    return new Contract(address, _abi, runner) as unknown as BeaconProxy;
  }
}
