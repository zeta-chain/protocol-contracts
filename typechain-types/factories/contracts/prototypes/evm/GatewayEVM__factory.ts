/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  GatewayEVM,
  GatewayEVMInterface,
} from "../../../../contracts/prototypes/evm/GatewayEVM";

const _abi = [
  {
    inputs: [],
    stateMutability: "nonpayable",
    type: "constructor",
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
] as const;

const _bytecode =
  "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b8152503480156200004757600080fd5b50620000586200005e60201b60201c565b62000208565b600060019054906101000a900460ff1615620000b1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a8906200015c565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff1614620001225760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200011991906200017e565b60405180910390a15b565b6000620001336027836200019b565b91506200014082620001b9565b604082019050919050565b6200015681620001ac565b82525050565b60006020820190508181036000830152620001778162000124565b9050919050565b60006020820190506200019560008301846200014b565b92915050565b600082825260208201905092915050565b600060ff82169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60805160601c612c9262000243600039600081816105fc0152818161068b01528181610785015281816108140152610c3f0152612c926000f3fe6080604052600436106100fe5760003560e01c8063715018a611610095578063c4d66de811610064578063c4d66de8146102e4578063dda79b751461030d578063f2fde38b14610338578063f340fa0114610361578063f45346dc1461037d576100fe565b8063715018a6146102505780638c6f037f146102675780638da5cb5b14610290578063ae7a3a6f146102bb576100fe565b80634f1ef286116100d15780634f1ef286146101a15780635131ab59146101bd57806352d1902d146101fa5780635b11259114610225576100fe565b80631b8b921d146101035780631cff79cd1461012c57806329c59b5d1461015c5780633659cfe614610178575b600080fd5b34801561010f57600080fd5b5061012a60048036038101906101259190611e16565b6103a6565b005b61014660048036038101906101419190611e16565b610412565b6040516101539190612463565b60405180910390f35b61017660048036038101906101719190611e16565b610480565b005b34801561018457600080fd5b5061019f600480360381019061019a9190611d61565b6105fa565b005b6101bb60048036038101906101b69190611e76565b610783565b005b3480156101c957600080fd5b506101e460048036038101906101df9190611d8e565b6108c0565b6040516101f19190612463565b60405180910390f35b34801561020657600080fd5b5061020f610c3b565b60405161021c9190612424565b60405180910390f35b34801561023157600080fd5b5061023a610cf4565b6040516102479190612380565b60405180910390f35b34801561025c57600080fd5b50610265610d1a565b005b34801561027357600080fd5b5061028e60048036038101906102899190611f25565b610d2e565b005b34801561029c57600080fd5b506102a5610e8d565b6040516102b29190612380565b60405180910390f35b3480156102c757600080fd5b506102e260048036038101906102dd9190611d61565b610eb7565b005b3480156102f057600080fd5b5061030b60048036038101906103069190611d61565b610efb565b005b34801561031957600080fd5b506103226110ea565b60405161032f9190612380565b60405180910390f35b34801561034457600080fd5b5061035f600480360381019061035a9190611d61565b611110565b005b61037b60048036038101906103769190611d61565b611194565b005b34801561038957600080fd5b506103a4600480360381019061039f9190611ed2565b611308565b005b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3848460405161040592919061243f565b60405180910390a3505050565b60606000610421858585611461565b90508473ffffffffffffffffffffffffffffffffffffffff167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f34868660405161046d9392919061269e565b60405180910390a2809150509392505050565b60003414156104bb576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16346040516105039061236b565b60006040518083038185875af1925050503d8060008114610540576040519150601f19603f3d011682016040523d82523d6000602084013e610545565b606091505b50509050600015158115151415610588576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a434600087876040516105ec9493929190612622565b60405180910390a350505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610689576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610680906124e2565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166106c8611518565b73ffffffffffffffffffffffffffffffffffffffff161461071e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071590612502565b60405180910390fd5b6107278161156f565b61078081600067ffffffffffffffff8111156107465761074561285f565b5b6040519080825280601f01601f1916602001820160405280156107785781602001600182028036833780820191505090505b50600061157a565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610812576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610809906124e2565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610851611518565b73ffffffffffffffffffffffffffffffffffffffff16146108a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089e90612502565b60405180910390fd5b6108b08261156f565b6108bc8282600161157a565b5050565b60608573ffffffffffffffffffffffffffffffffffffffff1663095ea7b38660006040518363ffffffff1660e01b81526004016108fe9291906123d2565b602060405180830381600087803b15801561091857600080fd5b505af115801561092c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109509190611fad565b508573ffffffffffffffffffffffffffffffffffffffff1663095ea7b386866040518363ffffffff1660e01b815260040161098c9291906123fb565b602060405180830381600087803b1580156109a657600080fd5b505af11580156109ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109de9190611fad565b5060006109ec868585611461565b90508673ffffffffffffffffffffffffffffffffffffffff1663095ea7b38760006040518363ffffffff1660e01b8152600401610a2a9291906123d2565b602060405180830381600087803b158015610a4457600080fd5b505af1158015610a58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a7c9190611fad565b5060008773ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610ab89190612380565b60206040518083038186803b158015610ad057600080fd5b505afa158015610ae4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b089190612007565b90506000811115610bc4578773ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401610b709291906123fb565b602060405180830381600087803b158015610b8a57600080fd5b505af1158015610b9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc29190611fad565b505b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382888888604051610c259392919061269e565b60405180910390a3819250505095945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610ccb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cc290612522565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b60ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610d226116f7565b610d2c6000611775565b565b6000841415610d69576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166323b872dd3360c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16876040518463ffffffff1660e01b8152600401610dc89392919061239b565b602060405180830381600087803b158015610de257600080fd5b505af1158015610df6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1a9190611fad565b508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a486868686604051610e7e9493929190612622565b60405180910390a35050505050565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b8060c960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060019054906101000a900460ff16159050808015610f2c5750600160008054906101000a900460ff1660ff16105b80610f595750610f3b3061183b565b158015610f585750600160008054906101000a900460ff1660ff16145b5b610f98576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f8f90612562565b60405180910390fd5b60016000806101000a81548160ff021916908360ff1602179055508015610fd5576001600060016101000a81548160ff0219169083151502179055505b610fdd61185e565b610fe56118b7565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561104c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160ca60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080156110e65760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516110dd9190612485565b60405180910390a15b5050565b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6111186116f7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611188576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161117f906124c2565b60405180910390fd5b61119181611775565b50565b60003414156111cf576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16346040516112179061236b565b60006040518083038185875af1925050503d8060008114611254576040519150601f19603f3d011682016040523d82523d6000602084013e611259565b606091505b5050905060001515811515141561129c576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a43460006040516112fc929190612662565b60405180910390a35050565b6000821415611343576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166323b872dd3360c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518463ffffffff1660e01b81526004016113a29392919061239b565b602060405180830381600087803b1580156113bc57600080fd5b505af11580156113d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113f49190611fad565b508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a48484604051611454929190612662565b60405180910390a3505050565b60606000808573ffffffffffffffffffffffffffffffffffffffff1634868660405161148e92919061233b565b60006040518083038185875af1925050503d80600081146114cb576040519150601f19603f3d011682016040523d82523d6000602084013e6114d0565b606091505b50915091508161150c576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80925050509392505050565b60006115467f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b611908565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6115776116f7565b50565b6115a67f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b611912565b60000160009054906101000a900460ff16156115ca576115c58361191c565b6116f2565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561161057600080fd5b505afa92505050801561164157506040513d601f19601f8201168201806040525081019061163e9190611fda565b60015b611680576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161167790612582565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b81146116e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116dc90612542565b60405180910390fd5b506116f18383836119d5565b5b505050565b6116ff611a01565b73ffffffffffffffffffffffffffffffffffffffff1661171d610e8d565b73ffffffffffffffffffffffffffffffffffffffff1614611773576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161176a906125c2565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff166118ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118a490612602565b60405180910390fd5b6118b5611a09565b565b600060019054906101000a900460ff16611906576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118fd90612602565b60405180910390fd5b565b6000819050919050565b6000819050919050565b6119258161183b565b611964576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161195b906125a2565b60405180910390fd5b806119917f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b611908565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6119de83611a6a565b6000825111806119eb5750805b156119fc576119fa8383611ab9565b505b505050565b600033905090565b600060019054906101000a900460ff16611a58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a4f90612602565b60405180910390fd5b611a68611a63611a01565b611775565b565b611a738161191c565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b6060611ade8383604051806060016040528060278152602001612c3660279139611ae6565b905092915050565b60606000808573ffffffffffffffffffffffffffffffffffffffff1685604051611b109190612354565b600060405180830381855af49150503d8060008114611b4b576040519150601f19603f3d011682016040523d82523d6000602084013e611b50565b606091505b5091509150611b6186838387611b6c565b925050509392505050565b60608315611bcf57600083511415611bc757611b878561183b565b611bc6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bbd906125e2565b60405180910390fd5b5b829050611bda565b611bd98383611be2565b5b949350505050565b600082511115611bf55781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c2991906124a0565b60405180910390fd5b6000611c45611c40846126f5565b6126d0565b905082815260208101848484011115611c6157611c6061289d565b5b611c6c8482856127ec565b509392505050565b600081359050611c8381612bd9565b92915050565b600081519050611c9881612bf0565b92915050565b600081519050611cad81612c07565b92915050565b60008083601f840112611cc957611cc8612893565b5b8235905067ffffffffffffffff811115611ce657611ce561288e565b5b602083019150836001820283011115611d0257611d01612898565b5b9250929050565b600082601f830112611d1e57611d1d612893565b5b8135611d2e848260208601611c32565b91505092915050565b600081359050611d4681612c1e565b92915050565b600081519050611d5b81612c1e565b92915050565b600060208284031215611d7757611d766128a7565b5b6000611d8584828501611c74565b91505092915050565b600080600080600060808688031215611daa57611da96128a7565b5b6000611db888828901611c74565b9550506020611dc988828901611c74565b9450506040611dda88828901611d37565b935050606086013567ffffffffffffffff811115611dfb57611dfa6128a2565b5b611e0788828901611cb3565b92509250509295509295909350565b600080600060408486031215611e2f57611e2e6128a7565b5b6000611e3d86828701611c74565b935050602084013567ffffffffffffffff811115611e5e57611e5d6128a2565b5b611e6a86828701611cb3565b92509250509250925092565b60008060408385031215611e8d57611e8c6128a7565b5b6000611e9b85828601611c74565b925050602083013567ffffffffffffffff811115611ebc57611ebb6128a2565b5b611ec885828601611d09565b9150509250929050565b600080600060608486031215611eeb57611eea6128a7565b5b6000611ef986828701611c74565b9350506020611f0a86828701611d37565b9250506040611f1b86828701611c74565b9150509250925092565b600080600080600060808688031215611f4157611f406128a7565b5b6000611f4f88828901611c74565b9550506020611f6088828901611d37565b9450506040611f7188828901611c74565b935050606086013567ffffffffffffffff811115611f9257611f916128a2565b5b611f9e88828901611cb3565b92509250509295509295909350565b600060208284031215611fc357611fc26128a7565b5b6000611fd184828501611c89565b91505092915050565b600060208284031215611ff057611fef6128a7565b5b6000611ffe84828501611c9e565b91505092915050565b60006020828403121561201d5761201c6128a7565b5b600061202b84828501611d4c565b91505092915050565b61203d81612769565b82525050565b61204c81612787565b82525050565b600061205e838561273c565b935061206b8385846127ec565b612074836128ac565b840190509392505050565b600061208b838561274d565b93506120988385846127ec565b82840190509392505050565b60006120af82612726565b6120b9818561273c565b93506120c98185602086016127fb565b6120d2816128ac565b840191505092915050565b60006120e882612726565b6120f2818561274d565b93506121028185602086016127fb565b80840191505092915050565b612117816127c8565b82525050565b612126816127da565b82525050565b600061213782612731565b6121418185612758565b93506121518185602086016127fb565b61215a816128ac565b840191505092915050565b6000612172602683612758565b915061217d826128bd565b604082019050919050565b6000612195602c83612758565b91506121a08261290c565b604082019050919050565b60006121b8602c83612758565b91506121c38261295b565b604082019050919050565b60006121db603883612758565b91506121e6826129aa565b604082019050919050565b60006121fe602983612758565b9150612209826129f9565b604082019050919050565b6000612221602e83612758565b915061222c82612a48565b604082019050919050565b6000612244602e83612758565b915061224f82612a97565b604082019050919050565b6000612267602d83612758565b915061227282612ae6565b604082019050919050565b600061228a602083612758565b915061229582612b35565b602082019050919050565b60006122ad60008361273c565b91506122b882612b5e565b600082019050919050565b60006122d060008361274d565b91506122db82612b5e565b600082019050919050565b60006122f3601d83612758565b91506122fe82612b61565b602082019050919050565b6000612316602b83612758565b915061232182612b8a565b604082019050919050565b612335816127b1565b82525050565b600061234882848661207f565b91508190509392505050565b600061236082846120dd565b915081905092915050565b6000612376826122c3565b9150819050919050565b60006020820190506123956000830184612034565b92915050565b60006060820190506123b06000830186612034565b6123bd6020830185612034565b6123ca604083018461232c565b949350505050565b60006040820190506123e76000830185612034565b6123f4602083018461210e565b9392505050565b60006040820190506124106000830185612034565b61241d602083018461232c565b9392505050565b60006020820190506124396000830184612043565b92915050565b6000602082019050818103600083015261245a818486612052565b90509392505050565b6000602082019050818103600083015261247d81846120a4565b905092915050565b600060208201905061249a600083018461211d565b92915050565b600060208201905081810360008301526124ba818461212c565b905092915050565b600060208201905081810360008301526124db81612165565b9050919050565b600060208201905081810360008301526124fb81612188565b9050919050565b6000602082019050818103600083015261251b816121ab565b9050919050565b6000602082019050818103600083015261253b816121ce565b9050919050565b6000602082019050818103600083015261255b816121f1565b9050919050565b6000602082019050818103600083015261257b81612214565b9050919050565b6000602082019050818103600083015261259b81612237565b9050919050565b600060208201905081810360008301526125bb8161225a565b9050919050565b600060208201905081810360008301526125db8161227d565b9050919050565b600060208201905081810360008301526125fb816122e6565b9050919050565b6000602082019050818103600083015261261b81612309565b9050919050565b6000606082019050612637600083018761232c565b6126446020830186612034565b8181036040830152612657818486612052565b905095945050505050565b6000606082019050612677600083018561232c565b6126846020830184612034565b8181036040830152612695816122a0565b90509392505050565b60006040820190506126b3600083018661232c565b81810360208301526126c6818486612052565b9050949350505050565b60006126da6126eb565b90506126e6828261282e565b919050565b6000604051905090565b600067ffffffffffffffff8211156127105761270f61285f565b5b612719826128ac565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061277482612791565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006127d3826127b1565b9050919050565b60006127e5826127bb565b9050919050565b82818337600083830152505050565b60005b838110156128195780820151818401526020810190506127fe565b83811115612828576000848401525b50505050565b612837826128ac565b810181811067ffffffffffffffff821117156128565761285561285f565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b612be281612769565b8114612bed57600080fd5b50565b612bf98161277b565b8114612c0457600080fd5b50565b612c1081612787565b8114612c1b57600080fd5b50565b612c27816127b1565b8114612c3257600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212209a3fc5b5d7f525a169e4b4d46f4725a2d4003761651eaad524854a78fa8041bc64736f6c63430008070033";

type GatewayEVMConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: GatewayEVMConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class GatewayEVM__factory extends ContractFactory {
  constructor(...args: GatewayEVMConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<GatewayEVM> {
    return super.deploy(overrides || {}) as Promise<GatewayEVM>;
  }
  override getDeployTransaction(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  override attach(address: string): GatewayEVM {
    return super.attach(address) as GatewayEVM;
  }
  override connect(signer: Signer): GatewayEVM__factory {
    return super.connect(signer) as GatewayEVM__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): GatewayEVMInterface {
    return new utils.Interface(_abi) as GatewayEVMInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): GatewayEVM {
    return new Contract(address, _abi, signerOrProvider) as GatewayEVM;
  }
}
