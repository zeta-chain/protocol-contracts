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
  BigNumberish,
  AddressLike,
  ContractDeployTransaction,
  ContractRunner,
} from "ethers";
import type { NonPayableOverrides } from "../../common";
import type { ZRC20, ZRC20Interface } from "../../ZRC20.sol/ZRC20";

const _abi = [
  {
    type: "constructor",
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
      {
        name: "chainid_",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "coinType_",
        type: "uint8",
        internalType: "enum CoinType",
      },
      {
        name: "gasLimit_",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "systemContractAddress_",
        type: "address",
        internalType: "address",
      },
      {
        name: "gatewayAddress_",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "CHAIN_ID",
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
    name: "COIN_TYPE",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "uint8",
        internalType: "enum CoinType",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "FUNGIBLE_MODULE_ADDRESS",
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
    name: "GAS_LIMIT",
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
    name: "PROTOCOL_FLAT_FEE",
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
    name: "SYSTEM_CONTRACT_ADDRESS",
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
        name: "account",
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
    name: "burn",
    inputs: [
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
    name: "deposit",
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
    name: "gatewayAddress",
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
        name: "recipient",
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
        name: "sender",
        type: "address",
        internalType: "address",
      },
      {
        name: "recipient",
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
    name: "updateGasLimit",
    inputs: [
      {
        name: "gasLimit_",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "updateGatewayAddress",
    inputs: [
      {
        name: "addr",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "updateProtocolFlatFee",
    inputs: [
      {
        name: "protocolFlatFee_",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "updateSystemContractAddress",
    inputs: [
      {
        name: "addr",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "withdraw",
    inputs: [
      {
        name: "to",
        type: "bytes",
        internalType: "bytes",
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
    name: "withdrawGasFee",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "view",
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
    name: "Deposit",
    inputs: [
      {
        name: "from",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
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
  {
    type: "event",
    name: "UpdatedGasLimit",
    inputs: [
      {
        name: "gasLimit",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "UpdatedGateway",
    inputs: [
      {
        name: "gateway",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "UpdatedProtocolFlatFee",
    inputs: [
      {
        name: "protocolFlatFee",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "UpdatedSystemContract",
    inputs: [
      {
        name: "systemContract",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Withdrawal",
    inputs: [
      {
        name: "from",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "to",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "value",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "gasFee",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "protocolFlatFee",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
  {
    type: "error",
    name: "CallerIsNotFungibleModule",
    inputs: [],
  },
  {
    type: "error",
    name: "GasFeeTransferFailed",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidSender",
    inputs: [],
  },
  {
    type: "error",
    name: "LowAllowance",
    inputs: [],
  },
  {
    type: "error",
    name: "LowBalance",
    inputs: [],
  },
  {
    type: "error",
    name: "ZeroAddress",
    inputs: [],
  },
  {
    type: "error",
    name: "ZeroGasCoin",
    inputs: [],
  },
  {
    type: "error",
    name: "ZeroGasPrice",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x60e060405234801561001057600080fd5b50604051611bb9380380611bb983398101604081905261002f9161020e565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006357604051632b2add3d60e01b815260040160405180910390fd5b6001600160a01b038216158061008057506001600160a01b038116155b1561009e5760405163d92e233d60e01b815260040160405180910390fd5b60066100aa8982610360565b5060076100b78882610360565b5060ff861660c05260808590528360028111156100d6576100d661041e565b60a08160028111156100ea576100ea61041e565b905250600192909255600080546001600160a01b039283166001600160a01b03199182161790915560088054929093169116179055506104349350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261015157600080fd5b81516001600160401b0381111561016a5761016a61012a565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101985761019861012a565b6040528181528382016020018510156101b057600080fd5b60005b828110156101cf576020818601810151838301820152016101b3565b506000918101602001919091529392505050565b8051600381106101f257600080fd5b919050565b80516001600160a01b03811681146101f257600080fd5b600080600080600080600080610100898b03121561022b57600080fd5b88516001600160401b0381111561024157600080fd5b61024d8b828c01610140565b60208b015190995090506001600160401b0381111561026b57600080fd5b6102778b828c01610140565b975050604089015160ff8116811461028e57600080fd5b60608a015190965094506102a460808a016101e3565b60a08a015190945092506102ba60c08a016101f7565b91506102c860e08a016101f7565b90509295985092959890939650565b600181811c908216806102eb57607f821691505b60208210810361030b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561035b57806000526020600020601f840160051c810160208510156103385750805b601f840160051c820191505b818110156103585760008155600101610344565b50505b505050565b81516001600160401b038111156103795761037961012a565b61038d8161038784546102d7565b84610311565b6020601f8211600181146103c157600083156103a95750848201515b600019600385901b1c1916600184901b178455610358565b600084815260208120601f198516915b828110156103f157878501518255602094850194600190920191016103d1565b508482101561040f5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a05160c05161174861047160003960006102170152600061033a0152600081816102eb01528181610ad70152610bdd01526117486000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c80638b851b95116100e3578063ccc775991161008c578063eddeb12311610066578063eddeb12314610431578063f2441b3214610444578063f687d12a1461046457600080fd5b8063ccc77599146103a4578063d9eeebed146103b7578063dd62ed3e146103eb57600080fd5b8063a9059cbb116100bd578063a9059cbb14610369578063c70126261461037c578063c835d7cc1461038f57600080fd5b80638b851b951461030d57806395d89b411461032d578063a3413d031461033557600080fd5b80633ce4a5bc116101455780634d8943bb1161011f5780634d8943bb146102a757806370a08231146102b057806385e1f4d0146102e657600080fd5b80633ce4a5bc1461024157806342966c681461028157806347e7ef241461029457600080fd5b806318160ddd1161017657806318160ddd146101f557806323b872dd146101fd578063313ce5671461021057600080fd5b806306fdde031461019d578063091d2788146101bb578063095ea7b3146101d2575b600080fd5b6101a5610477565b6040516101b29190611327565b60405180910390f35b6101c460015481565b6040519081526020016101b2565b6101e56101e0366004611366565b610509565b60405190151581526020016101b2565b6005546101c4565b6101e561020b366004611392565b610520565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101b2565b61025c73735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b2565b6101e561028f3660046113d3565b6105b7565b6101e56102a2366004611366565b6105cb565b6101c460025481565b6101c46102be3660046113ec565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101c47f000000000000000000000000000000000000000000000000000000000000000081565b60085461025c9073ffffffffffffffffffffffffffffffffffffffff1681565b6101a561071f565b61035c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516101b29190611409565b6101e5610377366004611366565b61072e565b6101e561038a366004611479565b61073b565b6103a261039d3660046113ec565b61088a565b005b6103a26103b23660046113ec565b61099e565b6103bf610aab565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101b2565b6101c46103f9366004611571565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b6103a261043f3660046113d3565b610cc9565b60005461025c9073ffffffffffffffffffffffffffffffffffffffff1681565b6103a26104723660046113d3565b610d4b565b606060068054610486906115aa565b80601f01602080910402602001604051908101604052809291908181526020018280546104b2906115aa565b80156104ff5780601f106104d4576101008083540402835291602001916104ff565b820191906000526020600020905b8154815290600101906020018083116104e257829003601f168201915b5050505050905090565b6000610516338484610dcd565b5060015b92915050565b600061052d848484610ed6565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260046020908152604080832033845290915290205482811015610598576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ac85336105a7868561162c565b610dcd565b506001949350505050565b60006105c33383611091565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab14801590610609575060005473ffffffffffffffffffffffffffffffffffffffff163314155b801561062d575060085473ffffffffffffffffffffffffffffffffffffffff163314155b15610664576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61066e83836111d3565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261070e91869061163f565b60405180910390a250600192915050565b606060078054610486906115aa565b6000610516338484610ed6565b6000806000610748610aab565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156107da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107fe9190611661565b610834576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61083e3385611091565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161087791899189918791611683565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146108d7576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610924576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109eb576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a38576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610993565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610b3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b6291906116b2565b905073ffffffffffffffffffffffffffffffffffffffff8116610bb1576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610c40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c6491906116cf565b905080600003610ca0576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610cb391906116e8565b610cbd91906116ff565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610d16576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610993565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610d98576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610993565b73ffffffffffffffffffffffffffffffffffffffff8316610e1a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610e67576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316610f23576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610f70576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604090205481811015610fd0576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fda828261162c565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260036020526040808220939093559085168152908120805484929061101d9084906116ff565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161108391815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166110de576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260409020548181101561113e576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611148828261162c565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600360205260408120919091556005805484929061118390849061162c565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610ec9565b73ffffffffffffffffffffffffffffffffffffffff8216611220576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806005600082825461123291906116ff565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805483929061126c9084906116ff565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b818110156112e9576020818501810151868301820152016112cd565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061133a60208301846112c3565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461136357600080fd5b50565b6000806040838503121561137957600080fd5b823561138481611341565b946020939093013593505050565b6000806000606084860312156113a757600080fd5b83356113b281611341565b925060208401356113c281611341565b929592945050506040919091013590565b6000602082840312156113e557600080fd5b5035919050565b6000602082840312156113fe57600080fd5b813561133a81611341565b6020810160038310611444577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806040838503121561148c57600080fd5b823567ffffffffffffffff8111156114a357600080fd5b8301601f810185136114b457600080fd5b803567ffffffffffffffff8111156114ce576114ce61144a565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561153a5761153a61144a565b60405281815282820160200187101561155257600080fd5b8160208401602083013760006020928201830152969401359450505050565b6000806040838503121561158457600080fd5b823561158f81611341565b9150602083013561159f81611341565b809150509250929050565b600181811c908216806115be57607f821691505b6020821081036115f7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561051a5761051a6115fd565b60408152600061165260408301856112c3565b90508260208301529392505050565b60006020828403121561167357600080fd5b8151801515811461133a57600080fd5b60808152600061169660808301876112c3565b6020830195909552506040810192909252606090910152919050565b6000602082840312156116c457600080fd5b815161133a81611341565b6000602082840312156116e157600080fd5b5051919050565b808202811582820484141761051a5761051a6115fd565b8082018082111561051a5761051a6115fd56fea26469706673582212204013e39523742118b3f2227e145ee7e48176ba6aebead33c596dca2acb59a74c64736f6c634300081a0033";

type ZRC20ConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: ZRC20ConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class ZRC20__factory extends ContractFactory {
  constructor(...args: ZRC20ConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    name_: string,
    symbol_: string,
    decimals_: BigNumberish,
    chainid_: BigNumberish,
    coinType_: BigNumberish,
    gasLimit_: BigNumberish,
    systemContractAddress_: AddressLike,
    gatewayAddress_: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(
      name_,
      symbol_,
      decimals_,
      chainid_,
      coinType_,
      gasLimit_,
      systemContractAddress_,
      gatewayAddress_,
      overrides || {}
    );
  }
  override deploy(
    name_: string,
    symbol_: string,
    decimals_: BigNumberish,
    chainid_: BigNumberish,
    coinType_: BigNumberish,
    gasLimit_: BigNumberish,
    systemContractAddress_: AddressLike,
    gatewayAddress_: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(
      name_,
      symbol_,
      decimals_,
      chainid_,
      coinType_,
      gasLimit_,
      systemContractAddress_,
      gatewayAddress_,
      overrides || {}
    ) as Promise<
      ZRC20 & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): ZRC20__factory {
    return super.connect(runner) as ZRC20__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ZRC20Interface {
    return new Interface(_abi) as ZRC20Interface;
  }
  static connect(address: string, runner?: ContractRunner | null): ZRC20 {
    return new Contract(address, _abi, runner) as unknown as ZRC20;
  }
}
