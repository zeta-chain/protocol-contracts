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
import type { NonPayableOverrides } from "../common";
import type { MockERC20, MockERC20Interface } from "../MockERC20";

const _abi = [
  {
    type: "function",
    name: "DOMAIN_SEPARATOR",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "allowance",
    inputs: [
      {
        name: "owner",
        type: "address",
        internalType: "address",
      },
      {
        name: "spender",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "approve",
    inputs: [
      {
        name: "spender",
        type: "address",
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "balanceOf",
    inputs: [
      {
        name: "owner",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "decimals",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "uint8",
        internalType: "uint8",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "initialize",
    inputs: [
      {
        name: "name_",
        type: "string",
        internalType: "string",
      },
      {
        name: "symbol_",
        type: "string",
        internalType: "string",
      },
      {
        name: "decimals_",
        type: "uint8",
        internalType: "uint8",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "name",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "string",
        internalType: "string",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "nonces",
    inputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "permit",
    inputs: [
      {
        name: "owner",
        type: "address",
        internalType: "address",
      },
      {
        name: "spender",
        type: "address",
        internalType: "address",
      },
      {
        name: "value",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "deadline",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "v",
        type: "uint8",
        internalType: "uint8",
      },
      {
        name: "r",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "s",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "symbol",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "string",
        internalType: "string",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "totalSupply",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "transfer",
    inputs: [
      {
        name: "to",
        type: "address",
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "transferFrom",
    inputs: [
      {
        name: "from",
        type: "address",
        internalType: "address",
      },
      {
        name: "to",
        type: "address",
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    name: "Approval",
    inputs: [
      {
        name: "owner",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "spender",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "value",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Transfer",
    inputs: [
      {
        name: "from",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "to",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "value",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
] as const;

const _bytecode =
  "0x6080604052348015600f57600080fd5b5061113e8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80633644e5151161008c57806395d89b411161006657806395d89b41146101d2578063a9059cbb146101da578063d505accf146101ed578063dd62ed3e1461020057600080fd5b80633644e5151461017457806370a082311461017c5780637ecebe00146101b257600080fd5b806318160ddd116100bd57806318160ddd1461013a57806323b872dd1461014c578063313ce5671461015f57600080fd5b806306fdde03146100e4578063095ea7b3146101025780631624f6c614610125575b600080fd5b6100ec610246565b6040516100f99190610b6b565b60405180910390f35b610115610110366004610be2565b6102d8565b60405190151581526020016100f9565b610138610133366004610cdc565b610352565b005b6003545b6040519081526020016100f9565b61011561015a366004610d55565b610451565b60025460405160ff90911681526020016100f9565b61013e6105c5565b61013e61018a366004610d92565b73ffffffffffffffffffffffffffffffffffffffff1660009081526004602052604090205490565b61013e6101c0366004610d92565b60086020526000908152604090205481565b6100ec6105eb565b6101156101e8366004610be2565b6105fa565b6101386101fb366004610dad565b6106ab565b61013e61020e366004610e18565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260056020908152604080832093909416825291909152205490565b60606000805461025590610e4b565b80601f016020809104026020016040519081016040528092919081815260200182805461028190610e4b565b80156102ce5780601f106102a3576101008083540402835291602001916102ce565b820191906000526020600020905b8154815290600101906020018083116102b157829003601f168201915b5050505050905090565b33600081815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103409086815260200190565b60405180910390a35060015b92915050565b60095460ff16156103c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f414c52454144595f494e495449414c495a45440000000000000000000000000060448201526064015b60405180910390fd5b60006103d08482610eed565b5060016103dd8382610eed565b50600280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff83161790556104136109b5565b60065561041e6109ce565b6007555050600980547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905550565b73ffffffffffffffffffffffffffffffffffffffff831660009081526005602090815260408083203384529091528120547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146104e5576104b38184610a71565b73ffffffffffffffffffffffffffffffffffffffff861660009081526005602090815260408083203384529091529020555b73ffffffffffffffffffffffffffffffffffffffff85166000908152600460205260409020546105159084610a71565b73ffffffffffffffffffffffffffffffffffffffff80871660009081526004602052604080822093909355908616815220546105519084610aee565b73ffffffffffffffffffffffffffffffffffffffff80861660008181526004602052604090819020939093559151908716907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906105b29087815260200190565b60405180910390a3506001949350505050565b60006006546105d26109b5565b146105e4576105df6109ce565b905090565b5060075490565b60606001805461025590610e4b565b336000908152600460205260408120546106149083610a71565b336000908152600460205260408082209290925573ffffffffffffffffffffffffffffffffffffffff85168152205461064d9083610aee565b73ffffffffffffffffffffffffffffffffffffffff84166000818152600460205260409081902092909255905133907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906103409086815260200190565b42841015610715576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f5045524d49545f444541444c494e455f4558504952454400000000000000000060448201526064016103bb565b600060016107216105c5565b73ffffffffffffffffffffffffffffffffffffffff8a16600090815260086020526040812080547f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9928d928d928d9290919061077c83611017565b9091555060408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810188905260e0016040516020818303038152906040528051906020012060405160200161081d9291907f190100000000000000000000000000000000000000000000000000000000000081526002810192909252602282015260420190565b60408051601f198184030181528282528051602091820120600084529083018083525260ff871690820152606081018590526080810184905260a0016020604051602081039080840390855afa15801561087b573d6000803e3d6000fd5b5050604051601f19015191505073ffffffffffffffffffffffffffffffffffffffff8116158015906108d857508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b61093e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f494e56414c49445f5349474e455200000000000000000000000000000000000060448201526064016103bb565b73ffffffffffffffffffffffffffffffffffffffff81811660009081526005602090815260408083208b8516808552908352928190208a90555189815291928b16917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a35050505050505050565b6000610b67806109c763ffffffff8216565b9250505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6000604051610a00919061104f565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6610a316109b5565b604080516020810195909552840192909252606083015260808201523060a082015260c00160405160208183030381529060405280519060200120905090565b600081831015610add576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f45524332303a207375627472616374696f6e20756e646572666c6f770000000060448201526064016103bb565b610ae782846110e2565b9392505050565b600080610afb83856110f5565b905083811015610ae7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45524332303a206164646974696f6e206f766572666c6f77000000000000000060448201526064016103bb565b4690565b602081526000825180602084015260005b81811015610b995760208186018101516040868401015201610b7c565b506000604082850101526040601f19601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610bdd57600080fd5b919050565b60008060408385031215610bf557600080fd5b610bfe83610bb9565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112610c4c57600080fd5b813567ffffffffffffffff811115610c6657610c66610c0c565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff82111715610c9657610c96610c0c565b604052818152838201602001851015610cae57600080fd5b816020850160208301376000918101602001919091529392505050565b803560ff81168114610bdd57600080fd5b600080600060608486031215610cf157600080fd5b833567ffffffffffffffff811115610d0857600080fd5b610d1486828701610c3b565b935050602084013567ffffffffffffffff811115610d3157600080fd5b610d3d86828701610c3b565b925050610d4c60408501610ccb565b90509250925092565b600080600060608486031215610d6a57600080fd5b610d7384610bb9565b9250610d8160208501610bb9565b929592945050506040919091013590565b600060208284031215610da457600080fd5b610ae782610bb9565b600080600080600080600060e0888a031215610dc857600080fd5b610dd188610bb9565b9650610ddf60208901610bb9565b95506040880135945060608801359350610dfb60808901610ccb565b9699959850939692959460a0840135945060c09093013592915050565b60008060408385031215610e2b57600080fd5b610e3483610bb9565b9150610e4260208401610bb9565b90509250929050565b600181811c90821680610e5f57607f821691505b602082108103610e98577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f821115610ee857806000526020600020601f840160051c81016020851015610ec55750805b601f840160051c820191505b81811015610ee55760008155600101610ed1565b50505b505050565b815167ffffffffffffffff811115610f0757610f07610c0c565b610f1b81610f158454610e4b565b84610e9e565b6020601f821160018114610f6d5760008315610f375750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455610ee5565b600084815260208120601f198516915b82811015610f9d5787850151825560209485019460019092019101610f7d565b5084821015610fd957868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361104857611048610fe8565b5060010190565b600080835461105d81610e4b565b60018216801561107457600181146110a7576110d7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831686528115158202860193506110d7565b86600052602060002060005b838110156110cf578154888201526001909101906020016110b3565b505081860193505b509195945050505050565b8181038181111561034c5761034c610fe8565b8082018082111561034c5761034c610fe856fea2646970667358221220e406cb0d3c8b8183614613515dc389660b1e2fd6eddd2dbdb8fa3cc5e030af5c64736f6c634300081a0033";

type MockERC20ConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: MockERC20ConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class MockERC20__factory extends ContractFactory {
  constructor(...args: MockERC20ConstructorParams) {
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
      MockERC20 & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): MockERC20__factory {
    return super.connect(runner) as MockERC20__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): MockERC20Interface {
    return new Interface(_abi) as MockERC20Interface;
  }
  static connect(address: string, runner?: ContractRunner | null): MockERC20 {
    return new Contract(address, _abi, runner) as unknown as MockERC20;
  }
}
