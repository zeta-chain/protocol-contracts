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
import type { MockERC721, MockERC721Interface } from "../MockERC721";

const _abi = [
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
        name: "id",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [],
    stateMutability: "payable",
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
    name: "getApproved",
    inputs: [
      {
        name: "id",
        type: "uint256",
        internalType: "uint256",
      },
    ],
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
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "isApprovedForAll",
    inputs: [
      {
        name: "owner",
        type: "address",
        internalType: "address",
      },
      {
        name: "operator",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "view",
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
    name: "ownerOf",
    inputs: [
      {
        name: "id",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "owner",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "safeTransferFrom",
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
        name: "id",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [],
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "safeTransferFrom",
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
        name: "id",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "data",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [],
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "setApprovalForAll",
    inputs: [
      {
        name: "operator",
        type: "address",
        internalType: "address",
      },
      {
        name: "approved",
        type: "bool",
        internalType: "bool",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "supportsInterface",
    inputs: [
      {
        name: "interfaceId",
        type: "bytes4",
        internalType: "bytes4",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "view",
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
    name: "tokenURI",
    inputs: [
      {
        name: "id",
        type: "uint256",
        internalType: "uint256",
      },
    ],
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
        name: "id",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [],
    stateMutability: "payable",
  },
  {
    type: "event",
    name: "Approval",
    inputs: [
      {
        name: "_owner",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "_approved",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "_tokenId",
        type: "uint256",
        indexed: true,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ApprovalForAll",
    inputs: [
      {
        name: "_owner",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "_operator",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "_approved",
        type: "bool",
        indexed: false,
        internalType: "bool",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Transfer",
    inputs: [
      {
        name: "_from",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "_to",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "_tokenId",
        type: "uint256",
        indexed: true,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
] as const;

const _bytecode =
  "0x6080604052348015600f57600080fd5b506114b88061001f6000396000f3fe6080604052600436106100dd5760003560e01c80636352211e1161007f578063a22cb46511610059578063a22cb4651461025f578063b88d4fde1461027f578063c87b56dd14610292578063e985e9c5146102b357600080fd5b80636352211e146101fc57806370a082311461021c57806395d89b411461024a57600080fd5b8063095ea7b3116100bb578063095ea7b3146101a157806323b872dd146101b657806342842e0e146101c95780634cd88b76146101dc57600080fd5b806301ffc9a7146100e257806306fdde0314610117578063081812fc14610139575b600080fd5b3480156100ee57600080fd5b506101026100fd366004610e1f565b610309565b60405190151581526020015b60405180910390f35b34801561012357600080fd5b5061012c6103ee565b60405161010e9190610ea7565b34801561014557600080fd5b5061017c610154366004610eba565b60009081526004602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010e565b6101b46101af366004610ef7565b610480565b005b6101b46101c4366004610f21565b6105cf565b6101b46101d7366004610f21565b6108c4565b3480156101e857600080fd5b506101b46101f7366004611045565b610a18565b34801561020857600080fd5b5061017c610217366004610eba565b610ace565b34801561022857600080fd5b5061023c6102373660046110ae565b610b5f565b60405190815260200161010e565b34801561025657600080fd5b5061012c610c07565b34801561026b57600080fd5b506101b461027a3660046110c9565b610c16565b6101b461028d366004611105565b610cad565b34801561029e57600080fd5b5061012c6102ad366004610eba565b50606090565b3480156102bf57600080fd5b506101026102ce366004611181565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260056020908152604080832093909416825291909152205460ff1690565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316148061039c57507f80ac58cd000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b806103e857507f5b5e139f000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6060600080546103fd906111b4565b80601f0160208091040260200160405190810160405280929190818152602001828054610429906111b4565b80156104765780601f1061044b57610100808354040283529160200191610476565b820191906000526020600020905b81548152906001019060200180831161045957829003601f168201915b5050505050905090565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff16338114806104e3575073ffffffffffffffffffffffffffffffffffffffff8116600090815260056020908152604080832033845290915290205460ff165b61054e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e4f545f415554484f52495a454400000000000000000000000000000000000060448201526064015b60405180910390fd5b60008281526004602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff87811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff84811691161461065f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f57524f4e475f46524f4d000000000000000000000000000000000000000000006044820152606401610545565b73ffffffffffffffffffffffffffffffffffffffff82166106dc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f494e56414c49445f524543495049454e540000000000000000000000000000006044820152606401610545565b3373ffffffffffffffffffffffffffffffffffffffff84161480610730575073ffffffffffffffffffffffffffffffffffffffff8316600090815260056020908152604080832033845290915290205460ff165b8061075e575060008181526004602052604090205473ffffffffffffffffffffffffffffffffffffffff1633145b6107c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e4f545f415554484f52495a45440000000000000000000000000000000000006044820152606401610545565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604081208054916107f583611236565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260036020526040812080549161082b8361126b565b90915550506000818152600260209081526040808320805473ffffffffffffffffffffffffffffffffffffffff8088167fffffffffffffffffffffffff000000000000000000000000000000000000000092831681179093556004909452828520805490911690559051849391928716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6108cf8383836105cf565b813b15806109ad57506040517f150b7a020000000000000000000000000000000000000000000000000000000080825233600483015273ffffffffffffffffffffffffffffffffffffffff858116602484015260448301849052608060648401526000608484015290919084169063150b7a029060a4016020604051808303816000875af1158015610965573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098991906112a3565b7fffffffff0000000000000000000000000000000000000000000000000000000016145b610a13576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f554e534146455f524543495049454e54000000000000000000000000000000006044820152606401610545565b505050565b60065460ff1615610a85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f414c52454144595f494e495449414c495a4544000000000000000000000000006044820152606401610545565b6000610a91838261130e565b506001610a9e828261130e565b5050600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905550565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff1680610b5a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e4f545f4d494e544544000000000000000000000000000000000000000000006044820152606401610545565b919050565b600073ffffffffffffffffffffffffffffffffffffffff8216610bde576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f5a45524f5f4144445245535300000000000000000000000000000000000000006044820152606401610545565b5073ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6060600180546103fd906111b4565b33600081815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168085529083529281902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b610cb88484846105cf565b823b1580610d8257506040517f150b7a02000000000000000000000000000000000000000000000000000000008082529073ffffffffffffffffffffffffffffffffffffffff85169063150b7a0290610d1b903390899088908890600401611427565b6020604051808303816000875af1158015610d3a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d5e91906112a3565b7fffffffff0000000000000000000000000000000000000000000000000000000016145b610de8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f554e534146455f524543495049454e54000000000000000000000000000000006044820152606401610545565b50505050565b7fffffffff0000000000000000000000000000000000000000000000000000000081168114610e1c57600080fd5b50565b600060208284031215610e3157600080fd5b8135610e3c81610dee565b9392505050565b6000815180845260005b81811015610e6957602081850181015186830182015201610e4d565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610e3c6020830184610e43565b600060208284031215610ecc57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610b5a57600080fd5b60008060408385031215610f0a57600080fd5b610f1383610ed3565b946020939093013593505050565b600080600060608486031215610f3657600080fd5b610f3f84610ed3565b9250610f4d60208501610ed3565b929592945050506040919091013590565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff841115610fa857610fa8610f5e565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715610ff557610ff5610f5e565b60405283815290508082840185101561100d57600080fd5b83836020830137600060208583010152509392505050565b600082601f83011261103657600080fd5b610e3c83833560208501610f8d565b6000806040838503121561105857600080fd5b823567ffffffffffffffff81111561106f57600080fd5b61107b85828601611025565b925050602083013567ffffffffffffffff81111561109857600080fd5b6110a485828601611025565b9150509250929050565b6000602082840312156110c057600080fd5b610e3c82610ed3565b600080604083850312156110dc57600080fd5b6110e583610ed3565b9150602083013580151581146110fa57600080fd5b809150509250929050565b6000806000806080858703121561111b57600080fd5b61112485610ed3565b935061113260208601610ed3565b925060408501359150606085013567ffffffffffffffff81111561115557600080fd5b8501601f8101871361116657600080fd5b61117587823560208401610f8d565b91505092959194509250565b6000806040838503121561119457600080fd5b61119d83610ed3565b91506111ab60208401610ed3565b90509250929050565b600181811c908216806111c857607f821691505b602082108103611201577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008161124557611245611207565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361129c5761129c611207565b5060010190565b6000602082840312156112b557600080fd5b8151610e3c81610dee565b601f821115610a1357806000526020600020601f840160051c810160208510156112e75750805b601f840160051c820191505b8181101561130757600081556001016112f3565b5050505050565b815167ffffffffffffffff81111561132857611328610f5e565b61133c8161133684546111b4565b846112c0565b6020601f82116001811461138e57600083156113585750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611307565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156113dc57878501518255602094850194600190920191016113bc565b508482101561141857868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b73ffffffffffffffffffffffffffffffffffffffff8516815273ffffffffffffffffffffffffffffffffffffffff841660208201528260408201526080606082015260006114786080830184610e43565b969550505050505056fea2646970667358221220bab234e314a458456f34dd6b1a1a49d22657ed9ac8d9a8a9cdab7365f18c4ee764736f6c634300081a0033";

type MockERC721ConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: MockERC721ConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class MockERC721__factory extends ContractFactory {
  constructor(...args: MockERC721ConstructorParams) {
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
      MockERC721 & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): MockERC721__factory {
    return super.connect(runner) as MockERC721__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): MockERC721Interface {
    return new Interface(_abi) as MockERC721Interface;
  }
  static connect(address: string, runner?: ContractRunner | null): MockERC721 {
    return new Contract(address, _abi, runner) as unknown as MockERC721;
  }
}
