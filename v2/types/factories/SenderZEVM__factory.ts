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
  "0x6080604052348015600f57600080fd5b50604051610a12380380610a12833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b6109858061008d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b61461005b5780637a34d8bb146100a4575b600080fd5b6100596100543660046105e7565b6100b7565b005b60005461007b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100596100b2366004610687565b61030c565b60008383836040516024016100ce93929190610783565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905260005490915073ffffffffffffffffffffffffffffffffffffffff8087169163095ea7b391166101758960026107ad565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af11580156101e5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020991906107ed565b61023f576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160808101825261032180825260016020808401829052838501929092528351918201845260008083526060840192909252905492517f90ad3e23000000000000000000000000000000000000000000000000000000008152919273ffffffffffffffffffffffffffffffffffffffff16916390ad3e23916102d0918c918c918c918991899060040161087a565b600060405180830381600087803b1580156102ea57600080fd5b505af11580156102fe573d6000803e3d6000fd5b505050505050505050505050565b600083838360405160240161032393929190610783565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905281516080810183526103218082526001828401819052828501919091528351928301845260008084526060830193909352915492517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff93841660048201526024810183905293945092909188169063095ea7b3906044016020604051808303816000875af1158015610447573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061046b91906107ed565b506000546040517fdc9ca2e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063dc9ca2e7906102d0908b908b908890879089906004016108e8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261050a57600080fd5b81356020830160008067ffffffffffffffff84111561052b5761052b6104ca565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715610578576105786104ca565b60405283815290508082840187101561059057600080fd5b838360208301376000602085830101528094505050505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146105d157600080fd5b919050565b80151581146105e457600080fd5b50565b60008060008060008060c0878903121561060057600080fd5b863567ffffffffffffffff81111561061757600080fd5b61062389828a016104f9565b96505060208701359450610639604088016105ad565b9350606087013567ffffffffffffffff81111561065557600080fd5b61066189828a016104f9565b9350506080870135915060a0870135610679816105d6565b809150509295509295509295565b600080600080600060a0868803121561069f57600080fd5b853567ffffffffffffffff8111156106b657600080fd5b6106c2888289016104f9565b9550506106d1602087016105ad565b9350604086013567ffffffffffffffff8111156106ed57600080fd5b6106f9888289016104f9565b935050606086013591506080860135610711816105d6565b809150509295509295909350565b6000815180845260005b8181101561074557602081850181015186830182015201610729565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b606081526000610796606083018661071f565b602083019490945250901515604090910152919050565b808201808211156107e7577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b6000602082840312156107ff57600080fd5b815161080a816105d6565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff60408201511660408301526000606082015160806060850152610872608085018261071f565b949350505050565b60c08152600061088d60c083018961071f565b87602084015273ffffffffffffffffffffffffffffffffffffffff8716604084015282810360608401526108c1818761071f565b905084608084015282810360a08401526108db8185610811565b9998505050505050505050565b60a0815260006108fb60a083018861071f565b73ffffffffffffffffffffffffffffffffffffffff871660208401528281036040840152610929818761071f565b905084606084015282810360808401526109438185610811565b9897505050505050505056fea26469706673582212204f4ea9cf773618d79d4e70bf94f359bbfd033c2b583a0249b440d641fcd9a50164736f6c634300081a0033";

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
