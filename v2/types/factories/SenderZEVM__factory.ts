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
import type { SenderZEVM, SenderZEVMInterface } from "../SenderZEVM";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "gateway_",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "callReceiver",
    inputs: [
      {
        name: "receiver",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "zrc20",
        type: "address",
        internalType: "address",
      },
      {
        name: "str",
        type: "string",
        internalType: "string",
      },
      {
        name: "num",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "flag",
        type: "bool",
        internalType: "bool",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "gateway",
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
    name: "withdrawAndCallReceiver",
    inputs: [
      {
        name: "receiver",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "zrc20",
        type: "address",
        internalType: "address",
      },
      {
        name: "str",
        type: "string",
        internalType: "string",
      },
      {
        name: "num",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "flag",
        type: "bool",
        internalType: "bool",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "error",
    name: "ApprovalFailed",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x6080604052348015600f57600080fd5b50604051610a2c380380610a2c833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b61099f8061008d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b61461005b5780637a34d8bb146100a4575b600080fd5b6100596100543660046105f5565b6100b7565b005b60005461007b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100596100b2366004610695565b610313565b60008383836040516024016100ce93929190610791565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905260005490915073ffffffffffffffffffffffffffffffffffffffff8087169163095ea7b391166101758960026107bb565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af11580156101e5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020991906107fb565b61023f576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160a0810182526103218082526001602080840182905283850192909252835191820184526000808352606084019290925260808301829052905492517f048ae42c000000000000000000000000000000000000000000000000000000008152919273ffffffffffffffffffffffffffffffffffffffff169163048ae42c916102d7918c918c918c9189918990600401610894565b600060405180830381600087803b1580156102f157600080fd5b505af1158015610305573d6000803e3d6000fd5b505050505050505050505050565b600083838360405160240161032a93929190610791565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f9700000000000000000000000000000000000000000000000000000000179052815160a081018352610321808252600182840181905282850191909152835192830184526000808452606083019390935260808201839052915492517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff93841660048201526024810183905293945092909188169063095ea7b3906044016020604051808303816000875af1158015610455573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047991906107fb565b506000546040517f1cb5ea7500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690631cb5ea75906102d7908b908b90889087908990600401610902565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261051857600080fd5b81356020830160008067ffffffffffffffff841115610539576105396104d8565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715610586576105866104d8565b60405283815290508082840187101561059e57600080fd5b838360208301376000602085830101528094505050505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146105df57600080fd5b919050565b80151581146105f257600080fd5b50565b60008060008060008060c0878903121561060e57600080fd5b863567ffffffffffffffff81111561062557600080fd5b61063189828a01610507565b96505060208701359450610647604088016105bb565b9350606087013567ffffffffffffffff81111561066357600080fd5b61066f89828a01610507565b9350506080870135915060a0870135610687816105e4565b809150509295509295509295565b600080600080600060a086880312156106ad57600080fd5b853567ffffffffffffffff8111156106c457600080fd5b6106d088828901610507565b9550506106df602087016105bb565b9350604086013567ffffffffffffffff8111156106fb57600080fd5b61070788828901610507565b93505060608601359150608086013561071f816105e4565b809150509295509295909350565b6000815180845260005b8181101561075357602081850181015186830182015201610737565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6060815260006107a4606083018661072d565b602083019490945250901515604090910152919050565b808201808211156107f5577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b60006020828403121561080d57600080fd5b8151610818816105e4565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff60408201511660408301526000606082015160a0606085015261088060a085018261072d565b608093840151949093019390935250919050565b60c0815260006108a760c083018961072d565b87602084015273ffffffffffffffffffffffffffffffffffffffff8716604084015282810360608401526108db818761072d565b905084608084015282810360a08401526108f5818561081f565b9998505050505050505050565b60a08152600061091560a083018861072d565b73ffffffffffffffffffffffffffffffffffffffff871660208401528281036040840152610943818761072d565b9050846060840152828103608084015261095d818561081f565b9897505050505050505056fea2646970667358221220313fe154e3326b0e70648c4c222266282e43bb43bc4ca556102b82635ee8d87464736f6c634300081a0033";

type SenderZEVMConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: SenderZEVMConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class SenderZEVM__factory extends ContractFactory {
  constructor(...args: SenderZEVMConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    gateway_: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(gateway_, overrides || {});
  }
  override deploy(
    gateway_: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(gateway_, overrides || {}) as Promise<
      SenderZEVM & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): SenderZEVM__factory {
    return super.connect(runner) as SenderZEVM__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): SenderZEVMInterface {
    return new Interface(_abi) as SenderZEVMInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): SenderZEVM {
    return new Contract(address, _abi, runner) as unknown as SenderZEVM;
  }
}
