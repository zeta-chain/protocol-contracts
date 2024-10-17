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
  "0x6080604052348015600f57600080fd5b50604051610aa9380380610aa9833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b610a1c8061008d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b61461005b5780637a34d8bb146100a4575b600080fd5b610059610054366004610658565b6100b7565b005b60005461007b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100596100b23660046106f8565b61032a565b60008383836040516024016100ce939291906107f4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905260005490915073ffffffffffffffffffffffffffffffffffffffff8087169163095ea7b3911661017589600261081e565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af11580156101e5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610209919061085e565b61023f576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160a081018252610321808252600160208084018290528385019290925283518083018552600080825260608501919091526080840181905284518086018652918252918101829052905492517f7b15118b0000000000000000000000000000000000000000000000000000000081529192909173ffffffffffffffffffffffffffffffffffffffff90911690637b15118b906102ed908c908c908c90899088908a906004016108f7565b600060405180830381600087803b15801561030757600080fd5b505af115801561031b573d6000803e3d6000fd5b50505050505050505050505050565b6000838383604051602401610341939291906107f4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f9700000000000000000000000000000000000000000000000000000000179052815160a081018352610321808252600182840181905282850191909152835180840185526000808252606084019190915260808301819052845180860186528281529384018190525493517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9485166004820152602481019190915293945092909188169063095ea7b3906044016020604051808303816000875af115801561047c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104a0919061085e565b506000546040517f06cb898300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116906306cb8983906104ff908b908b90889087908990600401610972565b600060405180830381600087803b15801561051957600080fd5b505af115801561052d573d6000803e3d6000fd5b505050505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261057b57600080fd5b81356020830160008067ffffffffffffffff84111561059c5761059c61053b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156105e9576105e961053b565b60405283815290508082840187101561060157600080fd5b838360208301376000602085830101528094505050505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461064257600080fd5b919050565b801515811461065557600080fd5b50565b60008060008060008060c0878903121561067157600080fd5b863567ffffffffffffffff81111561068857600080fd5b61069489828a0161056a565b965050602087013594506106aa6040880161061e565b9350606087013567ffffffffffffffff8111156106c657600080fd5b6106d289828a0161056a565b9350506080870135915060a08701356106ea81610647565b809150509295509295509295565b600080600080600060a0868803121561071057600080fd5b853567ffffffffffffffff81111561072757600080fd5b6107338882890161056a565b9550506107426020870161061e565b9350604086013567ffffffffffffffff81111561075e57600080fd5b61076a8882890161056a565b93505060608601359150608086013561078281610647565b809150509295509295909350565b6000815180845260005b818110156107b65760208185018101518683018201520161079a565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6060815260006108076060830186610790565b602083019490945250901515604090910152919050565b80820180821115610858577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b60006020828403121561087057600080fd5b815161087b81610647565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff60408201511660408301526000606082015160a060608501526108e360a0850182610790565b608093840151949093019390935250919050565b60e08152600061090a60e0830189610790565b87602084015273ffffffffffffffffffffffffffffffffffffffff87166040840152828103606084015261093e8187610790565b855160808501526020860151151560a0850152905082810360c08401526109658185610882565b9998505050505050505050565b60c08152600061098560c0830188610790565b73ffffffffffffffffffffffffffffffffffffffff8716602084015282810360408401526109b38187610790565b85516060850152602086015115156080850152905082810360a08401526109da8185610882565b9897505050505050505056fea2646970667358221220b23242a00fc774320fc48fb30348eb769c2fb8424032c0459d5e8de2913fd20e64736f6c634300081a0033";

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
