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
    type: "function",
    name: "withdrawGasFeeWithGasLimit",
    inputs: [
      {
        name: "gasLimit",
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
  "0x60e060405234801561001057600080fd5b50604051611e01380380611e0183398101604081905261002f9161020e565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006357604051632b2add3d60e01b815260040160405180910390fd5b6001600160a01b038216158061008057506001600160a01b038116155b1561009e5760405163d92e233d60e01b815260040160405180910390fd5b60066100aa8982610360565b5060076100b78882610360565b5060ff861660c05260808590528360028111156100d6576100d661041e565b60a08160028111156100ea576100ea61041e565b905250600192909255600080546001600160a01b039283166001600160a01b03199182161790915560088054929093169116179055506104349350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261015157600080fd5b81516001600160401b0381111561016a5761016a61012a565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101985761019861012a565b6040528181528382016020018510156101b057600080fd5b60005b828110156101cf576020818601810151838301820152016101b3565b506000918101602001919091529392505050565b8051600381106101f257600080fd5b919050565b80516001600160a01b03811681146101f257600080fd5b600080600080600080600080610100898b03121561022b57600080fd5b88516001600160401b0381111561024157600080fd5b61024d8b828c01610140565b60208b015190995090506001600160401b0381111561026b57600080fd5b6102778b828c01610140565b975050604089015160ff8116811461028e57600080fd5b60608a015190965094506102a460808a016101e3565b60a08a015190945092506102ba60c08a016101f7565b91506102c860e08a016101f7565b90509295985092959890939650565b600181811c908216806102eb57607f821691505b60208210810361030b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561035b57806000526020600020601f840160051c810160208510156103385750805b601f840160051c820191505b818110156103585760008155600101610344565b50505b505050565b81516001600160401b038111156103795761037961012a565b61038d8161038784546102d7565b84610311565b6020601f8211600181146103c157600083156103a95750848201515b600019600385901b1c1916600184901b178455610358565b600084815260208120601f198516915b828110156103f157878501518255602094850194600190920191016103d1565b508482101561040f5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a05160c05161198261047f6000396000610222015260006103450152600081816102f601528181610af501528181610bfb01528181610e170152610f1d01526119826000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c80638b851b95116100ee578063ccc7759911610097578063eddeb12311610071578063eddeb1231461043c578063f2441b321461044f578063f687d12a1461046f578063fc5fecd51461048257600080fd5b8063ccc77599146103af578063d9eeebed146103c2578063dd62ed3e146103f657600080fd5b8063a9059cbb116100c8578063a9059cbb14610374578063c701262614610387578063c835d7cc1461039a57600080fd5b80638b851b951461031857806395d89b4114610338578063a3413d031461034057600080fd5b80633ce4a5bc116101505780634d8943bb1161012a5780634d8943bb146102b257806370a08231146102bb57806385e1f4d0146102f157600080fd5b80633ce4a5bc1461024c57806342966c681461028c57806347e7ef241461029f57600080fd5b806318160ddd1161018157806318160ddd1461020057806323b872dd14610208578063313ce5671461021b57600080fd5b806306fdde03146101a8578063091d2788146101c6578063095ea7b3146101dd575b600080fd5b6101b0610495565b6040516101bd9190611561565b60405180910390f35b6101cf60015481565b6040519081526020016101bd565b6101f06101eb3660046115a0565b610527565b60405190151581526020016101bd565b6005546101cf565b6101f06102163660046115cc565b61053e565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101bd565b61026773735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101bd565b6101f061029a36600461160d565b6105d5565b6101f06102ad3660046115a0565b6105e9565b6101cf60025481565b6101cf6102c9366004611626565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101cf7f000000000000000000000000000000000000000000000000000000000000000081565b6008546102679073ffffffffffffffffffffffffffffffffffffffff1681565b6101b061073d565b6103677f000000000000000000000000000000000000000000000000000000000000000081565b6040516101bd9190611643565b6101f06103823660046115a0565b61074c565b6101f06103953660046116b3565b610759565b6103ad6103a8366004611626565b6108a8565b005b6103ad6103bd366004611626565b6109bc565b6103ca610ac9565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101bd565b6101cf6104043660046117ab565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b6103ad61044a36600461160d565b610ce7565b6000546102679073ffffffffffffffffffffffffffffffffffffffff1681565b6103ad61047d36600461160d565b610d69565b6103ca61049036600461160d565b610deb565b6060600680546104a4906117e4565b80601f01602080910402602001604051908101604052809291908181526020018280546104d0906117e4565b801561051d5780601f106104f25761010080835404028352916020019161051d565b820191906000526020600020905b81548152906001019060200180831161050057829003601f168201915b5050505050905090565b6000610534338484611007565b5060015b92915050565b600061054b848484611110565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054828110156105b6576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ca85336105c58685611866565b611007565b506001949350505050565b60006105e133836112cb565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab14801590610627575060005473ffffffffffffffffffffffffffffffffffffffff163314155b801561064b575060085473ffffffffffffffffffffffffffffffffffffffff163314155b15610682576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61068c838361140d565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261072c918690611879565b60405180910390a250600192915050565b6060600780546104a4906117e4565b6000610534338484611110565b6000806000610766610ac9565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156107f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081c919061189b565b610852576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61085c33856112cb565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d95591610895918991899187916118bd565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146108f5576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610942576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a09576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a56576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a387906020016109b1565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610b5c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8091906118ec565b905073ffffffffffffffffffffffffffffffffffffffff8116610bcf576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610c5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c829190611909565b905080600003610cbe576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610cd19190611922565b610cdb9190611939565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610d34576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f906020016109b1565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610db6576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a906020016109b1565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610e7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea291906118ec565b905073ffffffffffffffffffffffffffffffffffffffff8116610ef1576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610f80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa49190611909565b905080600003610fe0576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254600090610ff08784611922565b610ffa9190611939565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff8316611054576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166110a1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff831661115d576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166111aa576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83166000908152600360205260409020548181101561120a576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112148282611866565b73ffffffffffffffffffffffffffffffffffffffff8086166000908152600360205260408082209390935590851681529081208054849290611257908490611939565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516112bd91815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8216611318576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821660009081526003602052604090205481811015611378576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113828282611866565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812091909155600580548492906113bd908490611866565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001611103565b73ffffffffffffffffffffffffffffffffffffffff821661145a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806005600082825461146c9190611939565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260036020526040812080548392906114a6908490611939565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561152357602081850181015186830182015201611507565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061157460208301846114fd565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461159d57600080fd5b50565b600080604083850312156115b357600080fd5b82356115be8161157b565b946020939093013593505050565b6000806000606084860312156115e157600080fd5b83356115ec8161157b565b925060208401356115fc8161157b565b929592945050506040919091013590565b60006020828403121561161f57600080fd5b5035919050565b60006020828403121561163857600080fd5b81356115748161157b565b602081016003831061167e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080604083850312156116c657600080fd5b823567ffffffffffffffff8111156116dd57600080fd5b8301601f810185136116ee57600080fd5b803567ffffffffffffffff81111561170857611708611684565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561177457611774611684565b60405281815282820160200187101561178c57600080fd5b8160208401602083013760006020928201830152969401359450505050565b600080604083850312156117be57600080fd5b82356117c98161157b565b915060208301356117d98161157b565b809150509250929050565b600181811c908216806117f857607f821691505b602082108103611831577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561053857610538611837565b60408152600061188c60408301856114fd565b90508260208301529392505050565b6000602082840312156118ad57600080fd5b8151801515811461157457600080fd5b6080815260006118d060808301876114fd565b6020830195909552506040810192909252606090910152919050565b6000602082840312156118fe57600080fd5b81516115748161157b565b60006020828403121561191b57600080fd5b5051919050565b808202811582820484141761053857610538611837565b808201808211156105385761053861183756fea2646970667358221220e53ec821ebbbdda2b5f4e7943dbf9d29da74436ee8e6da3a4de1d6ce69a4627364736f6c634300081a0033";

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
