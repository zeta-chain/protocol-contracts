/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  GatewayEVMUpgradeTest,
  GatewayEVMUpgradeTestInterface,
} from "../../../../contracts/prototypes/evm/GatewayEVMUpgradeTest";

const _abi = [
  {
    inputs: [],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [],
    name: "ApprovalFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "CustodyInitialized",
    type: "error",
  },
  {
    inputs: [],
    name: "DepositFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "ExecutionFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "InsufficientERC20Amount",
    type: "error",
  },
  {
    inputs: [],
    name: "InsufficientETHAmount",
    type: "error",
  },
  {
    inputs: [],
    name: "ZeroAddress",
    type: "error",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "previousAdmin",
        type: "address",
      },
      {
        indexed: false,
        internalType: "address",
        name: "newAdmin",
        type: "address",
      },
    ],
    name: "AdminChanged",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "beacon",
        type: "address",
      },
    ],
    name: "BeaconUpgraded",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "sender",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "receiver",
        type: "address",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "payload",
        type: "bytes",
      },
    ],
    name: "Call",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "sender",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "receiver",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "address",
        name: "asset",
        type: "address",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "payload",
        type: "bytes",
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
        name: "destination",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "value",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
    ],
    name: "Executed",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "destination",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "value",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
    ],
    name: "ExecutedV2",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "token",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
    ],
    name: "ExecutedWithERC20",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint8",
        name: "version",
        type: "uint8",
      },
    ],
    name: "Initialized",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "previousOwner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "OwnershipTransferred",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "implementation",
        type: "address",
      },
    ],
    name: "Upgraded",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "receiver",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "payload",
        type: "bytes",
      },
    ],
    name: "call",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "custody",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "receiver",
        type: "address",
      },
    ],
    name: "deposit",
    outputs: [],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "receiver",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "asset",
        type: "address",
      },
    ],
    name: "deposit",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "receiver",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "payload",
        type: "bytes",
      },
    ],
    name: "depositAndCall",
    outputs: [],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "receiver",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "asset",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "payload",
        type: "bytes",
      },
    ],
    name: "depositAndCall",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "destination",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
    ],
    name: "execute",
    outputs: [
      {
        internalType: "bytes",
        name: "",
        type: "bytes",
      },
    ],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "token",
        type: "address",
      },
      {
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
    ],
    name: "executeWithERC20",
    outputs: [
      {
        internalType: "bytes",
        name: "",
        type: "bytes",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "_tssAddress",
        type: "address",
      },
      {
        internalType: "address",
        name: "_zetaToken",
        type: "address",
      },
    ],
    name: "initialize",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "owner",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "proxiableUUID",
    outputs: [
      {
        internalType: "bytes32",
        name: "",
        type: "bytes32",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "renounceOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "_zetaConnector",
        type: "address",
      },
    ],
    name: "setConnector",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "_custody",
        type: "address",
      },
    ],
    name: "setCustody",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "transferOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "tssAddress",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newImplementation",
        type: "address",
      },
    ],
    name: "upgradeTo",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newImplementation",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
    ],
    name: "upgradeToAndCall",
    outputs: [],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [],
    name: "zetaConnector",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "zetaToken",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
] as const;

const _bytecode =
  "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b81525034801561004657600080fd5b5060805160601c6134e36100816000396000818161078e0152818161081d01528181610b4801528181610bd70152610ff301526134e36000f3fe60806040526004361061011f5760003560e01c806357bec62f116100a0578063ae7a3a6f11610064578063ae7a3a6f14610384578063dda79b75146103ad578063f2fde38b146103d8578063f340fa0114610401578063f45346dc1461041d5761011f565b806357bec62f146102c35780635b112591146102ee578063715018a6146103195780638c6f037f146103305780638da5cb5b146103595761011f565b80633659cfe6116100e75780633659cfe6146101ed578063485cc955146102165780634f1ef2861461023f5780635131ab591461025b57806352d1902d146102985761011f565b806310188aef146101245780631b8b921d1461014d5780631cff79cd1461017657806321e093b1146101a657806329c59b5d146101d1575b600080fd5b34801561013057600080fd5b5061014b6004803603810190610146919061244e565b610446565b005b34801561015957600080fd5b50610174600480360381019061016f9190612543565b610512565b005b610190600480360381019061018b9190612543565b61057e565b60405161019d9190612bd6565b60405180910390f35b3480156101b257600080fd5b506101bb6105ec565b6040516101c89190612af3565b60405180910390f35b6101eb60048036038101906101e69190612543565b610612565b005b3480156101f957600080fd5b50610214600480360381019061020f919061244e565b61078c565b005b34801561022257600080fd5b5061023d6004803603810190610238919061247b565b610915565b005b610259600480360381019061025491906125a3565b610b46565b005b34801561026757600080fd5b50610282600480360381019061027d91906124bb565b610c83565b60405161028f9190612bd6565b60405180910390f35b3480156102a457600080fd5b506102ad610fef565b6040516102ba9190612b97565b60405180910390f35b3480156102cf57600080fd5b506102d86110a8565b6040516102e59190612af3565b60405180910390f35b3480156102fa57600080fd5b506103036110ce565b6040516103109190612af3565b60405180910390f35b34801561032557600080fd5b5061032e6110f4565b005b34801561033c57600080fd5b5061035760048036038101906103529190612652565b611108565b005b34801561036557600080fd5b5061036e611286565b60405161037b9190612af3565b60405180910390f35b34801561039057600080fd5b506103ab60048036038101906103a6919061244e565b6112b0565b005b3480156103b957600080fd5b506103c261137c565b6040516103cf9190612af3565b60405180910390f35b3480156103e457600080fd5b506103ff60048036038101906103fa919061244e565b6113a2565b005b61041b6004803603810190610416919061244e565b611426565b005b34801561042957600080fd5b50610444600480360381019061043f91906125ff565b61159a565b005b600073ffffffffffffffffffffffffffffffffffffffff1660cb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146104ce576040517fb337f37800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060cb60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde38484604051610571929190612bb2565b60405180910390a3505050565b6060600061058d858585611712565b90508473ffffffffffffffffffffffffffffffffffffffff167f373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e85463486866040516105d993929190612e51565b60405180910390a2809150509392505050565b60cc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600034141561064d576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163460405161069590612ade565b60006040518083038185875af1925050503d80600081146106d2576040519150601f19603f3d011682016040523d82523d6000602084013e6106d7565b606091505b5050905060001515811515141561071a576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4346000878760405161077e9493929190612dd5565b60405180910390a350505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16141561081b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081290612c55565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661085a6117c9565b73ffffffffffffffffffffffffffffffffffffffff16146108b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a790612c75565b60405180910390fd5b6108b981611820565b61091281600067ffffffffffffffff8111156108d8576108d7613012565b5b6040519080825280601f01601f19166020018201604052801561090a5781602001600182028036833780820191505090505b50600061182b565b50565b60008060019054906101000a900460ff161590508080156109465750600160008054906101000a900460ff1660ff16105b806109735750610955306119a8565b1580156109725750600160008054906101000a900460ff1660ff16145b5b6109b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a990612cf5565b60405180910390fd5b60016000806101000a81548160ff021916908360ff16021790555080156109ef576001600060016101000a81548160ff0219169083151502179055505b6109f76119cb565b6109ff611a24565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610a66576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8260ca60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160cc60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015610b415760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610b389190612bf8565b60405180910390a15b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610bd5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bcc90612c55565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610c146117c9565b73ffffffffffffffffffffffffffffffffffffffff1614610c6a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6190612c75565b60405180910390fd5b610c7382611820565b610c7f8282600161182b565b5050565b60606000841415610cc0576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cca8686611a75565b610d00576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff1663095ea7b386866040518363ffffffff1660e01b8152600401610d3b929190612b6e565b602060405180830381600087803b158015610d5557600080fd5b505af1158015610d69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d8d91906126da565b610dc3576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610dd0868585611712565b9050610ddc8787611a75565b610e12576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008773ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610e4d9190612af3565b60206040518083038186803b158015610e6557600080fd5b505afa158015610e79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e9d9190612734565b90506000811115610f7857600060c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060cc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff161415610f4b5760cb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b610f7681838b73ffffffffffffffffffffffffffffffffffffffff16611b0d9092919063ffffffff16565b505b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382888888604051610fd993929190612e51565b60405180910390a3819250505095945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161461107f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107690612cb5565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b60cb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6110fc611b93565b6111066000611c11565b565b6000841415611143576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060cc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156111e65760cb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b6112133382878773ffffffffffffffffffffffffffffffffffffffff16611cd7909392919063ffffffff16565b8573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4878787876040516112769493929190612dd5565b60405180910390a3505050505050565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600073ffffffffffffffffffffffffffffffffffffffff1660c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611338576040517fb337f37800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060c960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6113aa611b93565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561141a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141190612c35565b60405180910390fd5b61142381611c11565b50565b6000341415611461576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16346040516114a990612ade565b60006040518083038185875af1925050503d80600081146114e6576040519150601f19603f3d011682016040523d82523d6000602084013e6114eb565b606091505b5050905060001515811515141561152e576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a434600060405161158e929190612e15565b60405180910390a35050565b60008214156115d5576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060cc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156116785760cb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b6116a53382858573ffffffffffffffffffffffffffffffffffffffff16611cd7909392919063ffffffff16565b8373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a48585604051611704929190612e15565b60405180910390a350505050565b60606000808573ffffffffffffffffffffffffffffffffffffffff1634868660405161173f929190612aae565b60006040518083038185875af1925050503d806000811461177c576040519150601f19603f3d011682016040523d82523d6000602084013e611781565b606091505b5091509150816117bd576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80925050509392505050565b60006117f77f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b611d60565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611828611b93565b50565b6118577f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b611d6a565b60000160009054906101000a900460ff161561187b5761187683611d74565b6119a3565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156118c157600080fd5b505afa9250505080156118f257506040513d601f19601f820116820180604052508101906118ef9190612707565b60015b611931576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161192890612d15565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114611996576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161198d90612cd5565b60405180910390fd5b506119a2838383611e2d565b5b505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff16611a1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a1190612d95565b60405180910390fd5b611a22611e59565b565b600060019054906101000a900460ff16611a73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a6a90612d95565b60405180910390fd5b565b60008273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401611ab3929190612b45565b602060405180830381600087803b158015611acd57600080fd5b505af1158015611ae1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b0591906126da565b905092915050565b611b8e8363a9059cbb60e01b8484604051602401611b2c929190612b6e565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611eba565b505050565b611b9b611f81565b73ffffffffffffffffffffffffffffffffffffffff16611bb9611286565b73ffffffffffffffffffffffffffffffffffffffff1614611c0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c0690612d55565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611d5a846323b872dd60e01b858585604051602401611cf893929190612b0e565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611eba565b50505050565b6000819050919050565b6000819050919050565b611d7d816119a8565b611dbc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611db390612d35565b60405180910390fd5b80611de97f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b611d60565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611e3683611f89565b600082511180611e435750805b15611e5457611e528383611fd8565b505b505050565b600060019054906101000a900460ff16611ea8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e9f90612d95565b60405180910390fd5b611eb8611eb3611f81565b611c11565b565b6000611f1c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166120059092919063ffffffff16565b9050600081511115611f7c5780806020019051810190611f3c91906126da565b611f7b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f7290612db5565b60405180910390fd5b5b505050565b600033905090565b611f9281611d74565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b6060611ffd83836040518060600160405280602781526020016134876027913961201d565b905092915050565b606061201484846000856120a3565b90509392505050565b60606000808573ffffffffffffffffffffffffffffffffffffffff16856040516120479190612ac7565b600060405180830381855af49150503d8060008114612082576040519150601f19603f3d011682016040523d82523d6000602084013e612087565b606091505b509150915061209886838387612170565b925050509392505050565b6060824710156120e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120df90612c95565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516121119190612ac7565b60006040518083038185875af1925050503d806000811461214e576040519150601f19603f3d011682016040523d82523d6000602084013e612153565b606091505b5091509150612164878383876121e6565b92505050949350505050565b606083156121d3576000835114156121cb5761218b856119a8565b6121ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121c190612d75565b60405180910390fd5b5b8290506121de565b6121dd838361225c565b5b949350505050565b606083156122495760008351141561224157612201856122ac565b612240576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161223790612d75565b60405180910390fd5b5b829050612254565b61225383836122cf565b5b949350505050565b60008251111561226f5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122a39190612c13565b60405180910390fd5b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156122e25781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123169190612c13565b60405180910390fd5b600061233261232d84612ea8565b612e83565b90508281526020810184848401111561234e5761234d613050565b5b612359848285612f9f565b509392505050565b6000813590506123708161342a565b92915050565b60008151905061238581613441565b92915050565b60008151905061239a81613458565b92915050565b60008083601f8401126123b6576123b5613046565b5b8235905067ffffffffffffffff8111156123d3576123d2613041565b5b6020830191508360018202830111156123ef576123ee61304b565b5b9250929050565b600082601f83011261240b5761240a613046565b5b813561241b84826020860161231f565b91505092915050565b6000813590506124338161346f565b92915050565b6000815190506124488161346f565b92915050565b6000602082840312156124645761246361305a565b5b600061247284828501612361565b91505092915050565b600080604083850312156124925761249161305a565b5b60006124a085828601612361565b92505060206124b185828601612361565b9150509250929050565b6000806000806000608086880312156124d7576124d661305a565b5b60006124e588828901612361565b95505060206124f688828901612361565b945050604061250788828901612424565b935050606086013567ffffffffffffffff81111561252857612527613055565b5b612534888289016123a0565b92509250509295509295909350565b60008060006040848603121561255c5761255b61305a565b5b600061256a86828701612361565b935050602084013567ffffffffffffffff81111561258b5761258a613055565b5b612597868287016123a0565b92509250509250925092565b600080604083850312156125ba576125b961305a565b5b60006125c885828601612361565b925050602083013567ffffffffffffffff8111156125e9576125e8613055565b5b6125f5858286016123f6565b9150509250929050565b6000806000606084860312156126185761261761305a565b5b600061262686828701612361565b935050602061263786828701612424565b925050604061264886828701612361565b9150509250925092565b60008060008060006080868803121561266e5761266d61305a565b5b600061267c88828901612361565b955050602061268d88828901612424565b945050604061269e88828901612361565b935050606086013567ffffffffffffffff8111156126bf576126be613055565b5b6126cb888289016123a0565b92509250509295509295909350565b6000602082840312156126f0576126ef61305a565b5b60006126fe84828501612376565b91505092915050565b60006020828403121561271d5761271c61305a565b5b600061272b8482850161238b565b91505092915050565b60006020828403121561274a5761274961305a565b5b600061275884828501612439565b91505092915050565b61276a81612f1c565b82525050565b61277981612f3a565b82525050565b600061278b8385612eef565b9350612798838584612f9f565b6127a18361305f565b840190509392505050565b60006127b88385612f00565b93506127c5838584612f9f565b82840190509392505050565b60006127dc82612ed9565b6127e68185612eef565b93506127f6818560208601612fae565b6127ff8161305f565b840191505092915050565b600061281582612ed9565b61281f8185612f00565b935061282f818560208601612fae565b80840191505092915050565b61284481612f7b565b82525050565b61285381612f8d565b82525050565b600061286482612ee4565b61286e8185612f0b565b935061287e818560208601612fae565b6128878161305f565b840191505092915050565b600061289f602683612f0b565b91506128aa82613070565b604082019050919050565b60006128c2602c83612f0b565b91506128cd826130bf565b604082019050919050565b60006128e5602c83612f0b565b91506128f08261310e565b604082019050919050565b6000612908602683612f0b565b91506129138261315d565b604082019050919050565b600061292b603883612f0b565b9150612936826131ac565b604082019050919050565b600061294e602983612f0b565b9150612959826131fb565b604082019050919050565b6000612971602e83612f0b565b915061297c8261324a565b604082019050919050565b6000612994602e83612f0b565b915061299f82613299565b604082019050919050565b60006129b7602d83612f0b565b91506129c2826132e8565b604082019050919050565b60006129da602083612f0b565b91506129e582613337565b602082019050919050565b60006129fd600083612eef565b9150612a0882613360565b600082019050919050565b6000612a20600083612f00565b9150612a2b82613360565b600082019050919050565b6000612a43601d83612f0b565b9150612a4e82613363565b602082019050919050565b6000612a66602b83612f0b565b9150612a718261338c565b604082019050919050565b6000612a89602a83612f0b565b9150612a94826133db565b604082019050919050565b612aa881612f64565b82525050565b6000612abb8284866127ac565b91508190509392505050565b6000612ad3828461280a565b915081905092915050565b6000612ae982612a13565b9150819050919050565b6000602082019050612b086000830184612761565b92915050565b6000606082019050612b236000830186612761565b612b306020830185612761565b612b3d6040830184612a9f565b949350505050565b6000604082019050612b5a6000830185612761565b612b67602083018461283b565b9392505050565b6000604082019050612b836000830185612761565b612b906020830184612a9f565b9392505050565b6000602082019050612bac6000830184612770565b92915050565b60006020820190508181036000830152612bcd81848661277f565b90509392505050565b60006020820190508181036000830152612bf081846127d1565b905092915050565b6000602082019050612c0d600083018461284a565b92915050565b60006020820190508181036000830152612c2d8184612859565b905092915050565b60006020820190508181036000830152612c4e81612892565b9050919050565b60006020820190508181036000830152612c6e816128b5565b9050919050565b60006020820190508181036000830152612c8e816128d8565b9050919050565b60006020820190508181036000830152612cae816128fb565b9050919050565b60006020820190508181036000830152612cce8161291e565b9050919050565b60006020820190508181036000830152612cee81612941565b9050919050565b60006020820190508181036000830152612d0e81612964565b9050919050565b60006020820190508181036000830152612d2e81612987565b9050919050565b60006020820190508181036000830152612d4e816129aa565b9050919050565b60006020820190508181036000830152612d6e816129cd565b9050919050565b60006020820190508181036000830152612d8e81612a36565b9050919050565b60006020820190508181036000830152612dae81612a59565b9050919050565b60006020820190508181036000830152612dce81612a7c565b9050919050565b6000606082019050612dea6000830187612a9f565b612df76020830186612761565b8181036040830152612e0a81848661277f565b905095945050505050565b6000606082019050612e2a6000830185612a9f565b612e376020830184612761565b8181036040830152612e48816129f0565b90509392505050565b6000604082019050612e666000830186612a9f565b8181036020830152612e7981848661277f565b9050949350505050565b6000612e8d612e9e565b9050612e998282612fe1565b919050565b6000604051905090565b600067ffffffffffffffff821115612ec357612ec2613012565b5b612ecc8261305f565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000612f2782612f44565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000612f8682612f64565b9050919050565b6000612f9882612f6e565b9050919050565b82818337600083830152505050565b60005b83811015612fcc578082015181840152602081019050612fb1565b83811115612fdb576000848401525b50505050565b612fea8261305f565b810181811067ffffffffffffffff8211171561300957613008613012565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b61343381612f1c565b811461343e57600080fd5b50565b61344a81612f2e565b811461345557600080fd5b50565b61346181612f3a565b811461346c57600080fd5b50565b61347881612f64565b811461348357600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220444fb305a137c40a662d88f29c977e44ac9dc88e970228b6374e6e4c9c8ec39064736f6c63430008070033";

type GatewayEVMUpgradeTestConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: GatewayEVMUpgradeTestConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class GatewayEVMUpgradeTest__factory extends ContractFactory {
  constructor(...args: GatewayEVMUpgradeTestConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<GatewayEVMUpgradeTest> {
    return super.deploy(overrides || {}) as Promise<GatewayEVMUpgradeTest>;
  }
  override getDeployTransaction(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  override attach(address: string): GatewayEVMUpgradeTest {
    return super.attach(address) as GatewayEVMUpgradeTest;
  }
  override connect(signer: Signer): GatewayEVMUpgradeTest__factory {
    return super.connect(signer) as GatewayEVMUpgradeTest__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): GatewayEVMUpgradeTestInterface {
    return new utils.Interface(_abi) as GatewayEVMUpgradeTestInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): GatewayEVMUpgradeTest {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as GatewayEVMUpgradeTest;
  }
}
