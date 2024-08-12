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
        name: "_gateway",
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
        name: "chainId",
        type: "uint256",
        internalType: "uint256",
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
  "0x6080604052348015600f57600080fd5b50604051610908380380610908833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b61087b8061008d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b61461005b578063865b36f6146100a4575b600080fd5b61005961005436600461052d565b6100b7565b005b60005461007b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100596100b23660046105e8565b6102e5565b60008383836040516024016100ce939291906106dd565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905260005490517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526024810189905291925086169063095ea7b3906044016020604051808303816000875af11580156101be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101e29190610707565b610218576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160808101825261032180825260016020808401919091528284019190915282519081018352600080825260608301919091525491517f939bc896000000000000000000000000000000000000000000000000000000008152909173ffffffffffffffffffffffffffffffffffffffff169063939bc896906102a9908b908b908b9088908890600401610794565b600060405180830381600087803b1580156102c357600080fd5b505af11580156102d7573d6000803e3d6000fd5b505050505050505050505050565b60008383836040516024016102fc939291906106dd565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f97000000000000000000000000000000000000000000000000000000001790528151608081018352610321808252600182840152818401528251918201835260008083526060820192909252905491517f845b96ce0000000000000000000000000000000000000000000000000000000081529293509173ffffffffffffffffffffffffffffffffffffffff919091169063845b96ce906103fe908a908a90879087906004016107fb565b600060405180830381600087803b15801561041857600080fd5b505af115801561042c573d6000803e3d6000fd5b5050505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261047957600080fd5b81356020830160008067ffffffffffffffff84111561049a5761049a610439565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156104e7576104e7610439565b6040528381529050808284018710156104ff57600080fd5b838360208301376000602085830101528094505050505092915050565b801515811461052a57600080fd5b50565b60008060008060008060c0878903121561054657600080fd5b863567ffffffffffffffff81111561055d57600080fd5b61056989828a01610468565b96505060208701359450604087013573ffffffffffffffffffffffffffffffffffffffff8116811461059a57600080fd5b9350606087013567ffffffffffffffff8111156105b657600080fd5b6105c289828a01610468565b9350506080870135915060a08701356105da8161051c565b809150509295509295509295565b600080600080600060a0868803121561060057600080fd5b853567ffffffffffffffff81111561061757600080fd5b61062388828901610468565b95505060208601359350604086013567ffffffffffffffff81111561064757600080fd5b61065388828901610468565b93505060608601359150608086013561066b8161051c565b809150509295509295909350565b6000815180845260005b8181101561069f57602081850181015186830182015201610683565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6060815260006106f06060830186610679565b602083019490945250901515604090910152919050565b60006020828403121561071957600080fd5b81516107248161051c565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff6040820151166040830152600060608201516080606085015261078c6080850182610679565b949350505050565b60a0815260006107a760a0830188610679565b86602084015273ffffffffffffffffffffffffffffffffffffffff8616604084015282810360608401526107db8186610679565b905082810360808401526107ef818561072b565b98975050505050505050565b60808152600061080e6080830187610679565b85602084015282810360408401526108268186610679565b9050828103606084015261083a818561072b565b97965050505050505056fea2646970667358221220c9953a5a3591ba568d24416b7c5e66851cb43783dbb9c5612d4167988483236364736f6c634300081a0033";

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
    _gateway: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(_gateway, overrides || {});
  }
  override deploy(
    _gateway: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(_gateway, overrides || {}) as Promise<
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
