/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  WETH9,
  WETH9Interface,
} from "../../../../contracts/zevm/WZETA.sol/WETH9";

const _abi = [
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "src",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "guy",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "wad",
        type: "uint256",
      },
    ],
    name: "Approval",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "dst",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "wad",
        type: "uint256",
      },
    ],
    name: "Deposit",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "src",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "dst",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "wad",
        type: "uint256",
      },
    ],
    name: "Transfer",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "src",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "wad",
        type: "uint256",
      },
    ],
    name: "Withdrawal",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    name: "allowance",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "guy",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "wad",
        type: "uint256",
      },
    ],
    name: "approve",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    name: "balanceOf",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "decimals",
    outputs: [
      {
        internalType: "uint8",
        name: "",
        type: "uint8",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "deposit",
    outputs: [],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [],
    name: "name",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "symbol",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "totalSupply",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "dst",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "wad",
        type: "uint256",
      },
    ],
    name: "transfer",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "src",
        type: "address",
      },
      {
        internalType: "address",
        name: "dst",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "wad",
        type: "uint256",
      },
    ],
    name: "transferFrom",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "wad",
        type: "uint256",
      },
    ],
    name: "withdraw",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    stateMutability: "payable",
    type: "receive",
  },
] as const;

const _bytecode =
  "0x60806040526040518060400160405280600d81526020017f57726170706564204574686572000000000000000000000000000000000000008152506000908051906020019062000051929190620000d0565b506040518060400160405280600481526020017f5745544800000000000000000000000000000000000000000000000000000000815250600190805190602001906200009f929190620000d0565b506012600260006101000a81548160ff021916908360ff160217905550348015620000c957600080fd5b50620001e5565b828054620000de9062000180565b90600052602060002090601f0160209004810192826200010257600085556200014e565b82601f106200011d57805160ff19168380011785556200014e565b828001600101855582156200014e579182015b828111156200014d57825182559160200191906001019062000130565b5b5090506200015d919062000161565b5090565b5b808211156200017c57600081600090555060010162000162565b5090565b600060028204905060018216806200019957607f821691505b60208210811415620001b057620001af620001b6565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b610fd380620001f56000396000f3fe6080604052600436106100a05760003560e01c8063313ce56711610064578063313ce567146101ad57806370a08231146101d857806395d89b4114610215578063a9059cbb14610240578063d0e30db01461027d578063dd62ed3e14610287576100af565b806306fdde03146100b4578063095ea7b3146100df57806318160ddd1461011c57806323b872dd146101475780632e1a7d4d14610184576100af565b366100af576100ad6102c4565b005b600080fd5b3480156100c057600080fd5b506100c961036a565b6040516100d69190610d20565b60405180910390f35b3480156100eb57600080fd5b5061010660048036038101906101019190610c0f565b6103f8565b6040516101139190610d05565b60405180910390f35b34801561012857600080fd5b506101316104ea565b60405161013e9190610d62565b60405180910390f35b34801561015357600080fd5b5061016e60048036038101906101699190610bbc565b6104f2565b60405161017b9190610d05565b60405180910390f35b34801561019057600080fd5b506101ab60048036038101906101a69190610c4f565b6108c2565b005b3480156101b957600080fd5b506101c2610a32565b6040516101cf9190610d7d565b60405180910390f35b3480156101e457600080fd5b506101ff60048036038101906101fa9190610b4f565b610a45565b60405161020c9190610d62565b60405180910390f35b34801561022157600080fd5b5061022a610a5d565b6040516102379190610d20565b60405180910390f35b34801561024c57600080fd5b5061026760048036038101906102629190610c0f565b610aeb565b6040516102749190610d05565b60405180910390f35b6102856102c4565b005b34801561029357600080fd5b506102ae60048036038101906102a99190610b7c565b610b00565b6040516102bb9190610d62565b60405180910390f35b34600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546103139190610db4565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c346040516103609190610d62565b60405180910390a2565b6000805461037790610ec6565b80601f01602080910402602001604051908101604052809291908181526020018280546103a390610ec6565b80156103f05780601f106103c5576101008083540402835291602001916103f0565b820191906000526020600020905b8154815290600101906020018083116103d357829003601f168201915b505050505081565b600081600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104d89190610d62565b60405180910390a36001905092915050565b600047905090565b600081600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610576576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056d90610d42565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415801561064e57507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414155b156107a65781600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610712576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070990610d42565b60405180910390fd5b81600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461079e9190610e0a565b925050819055505b81600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107f59190610e0a565b9250508190555081600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461084b9190610db4565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516108af9190610d62565b60405180910390a3600190509392505050565b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610944576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093b90610d42565b60405180910390fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109939190610e0a565b925050819055503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156109e0573d6000803e3d6000fd5b503373ffffffffffffffffffffffffffffffffffffffff167f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b6582604051610a279190610d62565b60405180910390a250565b600260009054906101000a900460ff1681565b60036020528060005260406000206000915090505481565b60018054610a6a90610ec6565b80601f0160208091040260200160405190810160405280929190818152602001828054610a9690610ec6565b8015610ae35780601f10610ab857610100808354040283529160200191610ae3565b820191906000526020600020905b815481529060010190602001808311610ac657829003601f168201915b505050505081565b6000610af83384846104f2565b905092915050565b6004602052816000526040600020602052806000526040600020600091509150505481565b600081359050610b3481610f6f565b92915050565b600081359050610b4981610f86565b92915050565b600060208284031215610b6557610b64610f56565b5b6000610b7384828501610b25565b91505092915050565b60008060408385031215610b9357610b92610f56565b5b6000610ba185828601610b25565b9250506020610bb285828601610b25565b9150509250929050565b600080600060608486031215610bd557610bd4610f56565b5b6000610be386828701610b25565b9350506020610bf486828701610b25565b9250506040610c0586828701610b3a565b9150509250925092565b60008060408385031215610c2657610c25610f56565b5b6000610c3485828601610b25565b9250506020610c4585828601610b3a565b9150509250929050565b600060208284031215610c6557610c64610f56565b5b6000610c7384828501610b3a565b91505092915050565b610c8581610e50565b82525050565b6000610c9682610d98565b610ca08185610da3565b9350610cb0818560208601610e93565b610cb981610f5b565b840191505092915050565b6000610cd1600083610da3565b9150610cdc82610f6c565b600082019050919050565b610cf081610e7c565b82525050565b610cff81610e86565b82525050565b6000602082019050610d1a6000830184610c7c565b92915050565b60006020820190508181036000830152610d3a8184610c8b565b905092915050565b60006020820190508181036000830152610d5b81610cc4565b9050919050565b6000602082019050610d776000830184610ce7565b92915050565b6000602082019050610d926000830184610cf6565b92915050565b600081519050919050565b600082825260208201905092915050565b6000610dbf82610e7c565b9150610dca83610e7c565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610dff57610dfe610ef8565b5b828201905092915050565b6000610e1582610e7c565b9150610e2083610e7c565b925082821015610e3357610e32610ef8565b5b828203905092915050565b6000610e4982610e5c565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015610eb1578082015181840152602081019050610e96565b83811115610ec0576000848401525b50505050565b60006002820490506001821680610ede57607f821691505b60208210811415610ef257610ef1610f27565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b6000601f19601f8301169050919050565b50565b610f7881610e3e565b8114610f8357600080fd5b50565b610f8f81610e7c565b8114610f9a57600080fd5b5056fea2646970667358221220ed2297470e8d6c8e387b5cdc1c81dd38decdf0b011f3c15df9f52f6da3dcc17664736f6c63430008070033";

type WETH9ConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: WETH9ConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class WETH9__factory extends ContractFactory {
  constructor(...args: WETH9ConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<WETH9> {
    return super.deploy(overrides || {}) as Promise<WETH9>;
  }
  override getDeployTransaction(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  override attach(address: string): WETH9 {
    return super.attach(address) as WETH9;
  }
  override connect(signer: Signer): WETH9__factory {
    return super.connect(signer) as WETH9__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): WETH9Interface {
    return new utils.Interface(_abi) as WETH9Interface;
  }
  static connect(address: string, signerOrProvider: Signer | Provider): WETH9 {
    return new Contract(address, _abi, signerOrProvider) as WETH9;
  }
}
