/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Signer,
  utils,
  Contract,
  ContractFactory,
  BigNumberish,
  Overrides,
} from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  ZRC20,
  ZRC20Interface,
} from "../../../../contracts/zevm/ZRC20.sol/ZRC20";

const _abi = [
  {
    inputs: [
      {
        internalType: "string",
        name: "name_",
        type: "string",
      },
      {
        internalType: "string",
        name: "symbol_",
        type: "string",
      },
      {
        internalType: "uint8",
        name: "decimals_",
        type: "uint8",
      },
      {
        internalType: "uint256",
        name: "chainid_",
        type: "uint256",
      },
      {
        internalType: "enum CoinType",
        name: "coinType_",
        type: "uint8",
      },
      {
        internalType: "uint256",
        name: "gasLimit_",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "systemContractAddress_",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [],
    name: "CallerIsNotFungibleModule",
    type: "error",
  },
  {
    inputs: [],
    name: "GasFeeTransferFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidSender",
    type: "error",
  },
  {
    inputs: [],
    name: "LowAllowance",
    type: "error",
  },
  {
    inputs: [],
    name: "LowBalance",
    type: "error",
  },
  {
    inputs: [],
    name: "ZeroAddress",
    type: "error",
  },
  {
    inputs: [],
    name: "ZeroGasCoin",
    type: "error",
  },
  {
    inputs: [],
    name: "ZeroGasPrice",
    type: "error",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "owner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "spender",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "value",
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
        indexed: false,
        internalType: "bytes",
        name: "from",
        type: "bytes",
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
        name: "value",
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
        name: "from",
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
        name: "value",
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
        indexed: false,
        internalType: "uint256",
        name: "gasLimit",
        type: "uint256",
      },
    ],
    name: "UpdatedGasLimit",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint256",
        name: "protocolFlatFee",
        type: "uint256",
      },
    ],
    name: "UpdatedProtocolFlatFee",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "systemContract",
        type: "address",
      },
    ],
    name: "UpdatedSystemContract",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "from",
        type: "address",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "to",
        type: "bytes",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "value",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "gasfee",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "protocolFlatFee",
        type: "uint256",
      },
    ],
    name: "Withdrawal",
    type: "event",
  },
  {
    inputs: [],
    name: "CHAIN_ID",
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
    name: "COIN_TYPE",
    outputs: [
      {
        internalType: "enum CoinType",
        name: "",
        type: "uint8",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "FUNGIBLE_MODULE_ADDRESS",
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
    name: "GAS_LIMIT",
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
    name: "PROTOCOL_FLAT_FEE",
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
    name: "SYSTEM_CONTRACT_ADDRESS",
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
        name: "owner",
        type: "address",
      },
      {
        internalType: "address",
        name: "spender",
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
        name: "spender",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
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
        name: "account",
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
    inputs: [
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "burn",
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
    inputs: [
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
    ],
    name: "deposit",
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
        name: "recipient",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
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
        name: "sender",
        type: "address",
      },
      {
        internalType: "address",
        name: "recipient",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
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
        name: "gasLimit",
        type: "uint256",
      },
    ],
    name: "updateGasLimit",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "protocolFlatFee",
        type: "uint256",
      },
    ],
    name: "updateProtocolFlatFee",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "addr",
        type: "address",
      },
    ],
    name: "updateSystemContractAddress",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes",
        name: "to",
        type: "bytes",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "withdraw",
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
    inputs: [],
    name: "withdrawGasFee",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
] as const;

const _bytecode =
  "0x60c06040523480156200001157600080fd5b506040516200272c3803806200272c833981810160405281019062000037919062000319565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620000b1576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8660069080519060200190620000c99291906200018f565b508560079080519060200190620000e29291906200018f565b5084600860006101000a81548160ff021916908360ff16021790555083608081815250508260028111156200011c576200011b62000556565b5b60a081600281111562000134576200013362000556565b5b60f81b8152505081600181905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050505062000667565b8280546200019d90620004ea565b90600052602060002090601f016020900481019282620001c157600085556200020d565b82601f10620001dc57805160ff19168380011785556200020d565b828001600101855582156200020d579182015b828111156200020c578251825591602001919060010190620001ef565b5b5090506200021c919062000220565b5090565b5b808211156200023b57600081600090555060010162000221565b5090565b600062000256620002508462000433565b6200040a565b905082815260208101848484011115620002755762000274620005e8565b5b62000282848285620004b4565b509392505050565b6000815190506200029b8162000608565b92915050565b600081519050620002b28162000622565b92915050565b600082601f830112620002d057620002cf620005e3565b5b8151620002e28482602086016200023f565b91505092915050565b600081519050620002fc8162000633565b92915050565b60008151905062000313816200064d565b92915050565b600080600080600080600060e0888a0312156200033b576200033a620005f2565b5b600088015167ffffffffffffffff8111156200035c576200035b620005ed565b5b6200036a8a828b01620002b8565b975050602088015167ffffffffffffffff8111156200038e576200038d620005ed565b5b6200039c8a828b01620002b8565b9650506040620003af8a828b0162000302565b9550506060620003c28a828b01620002eb565b9450506080620003d58a828b01620002a1565b93505060a0620003e88a828b01620002eb565b92505060c0620003fb8a828b016200028a565b91505092959891949750929550565b60006200041662000429565b905062000424828262000520565b919050565b6000604051905090565b600067ffffffffffffffff821115620004515762000450620005b4565b5b6200045c82620005f7565b9050602081019050919050565b600062000476826200047d565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015620004d4578082015181840152602081019050620004b7565b83811115620004e4576000848401525b50505050565b600060028204905060018216806200050357607f821691505b602082108114156200051a576200051962000585565b5b50919050565b6200052b82620005f7565b810181811067ffffffffffffffff821117156200054d576200054c620005b4565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b620006138162000469565b81146200061f57600080fd5b50565b600381106200063057600080fd5b50565b6200063e816200049d565b81146200064a57600080fd5b50565b6200065881620004a7565b81146200066457600080fd5b50565b60805160a05160f81c61208e6200069e60003960006108d501526000818161081f01528181610ba20152610cd7015261208e6000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806385e1f4d0116100b8578063c835d7cc1161007c578063c835d7cc146103a5578063d9eeebed146103c1578063dd62ed3e146103e0578063eddeb12314610410578063f2441b321461042c578063f687d12a1461044a57610142565b806385e1f4d0146102eb57806395d89b4114610309578063a3413d0314610327578063a9059cbb14610345578063c70126261461037557610142565b8063313ce5671161010a578063313ce567146102015780633ce4a5bc1461021f57806342966c681461023d57806347e7ef241461026d5780634d8943bb1461029d57806370a08231146102bb57610142565b806306fdde0314610147578063091d278814610165578063095ea7b31461018357806318160ddd146101b357806323b872dd146101d1575b600080fd5b61014f610466565b60405161015c9190611c04565b60405180910390f35b61016d6104f8565b60405161017a9190611c26565b60405180910390f35b61019d600480360381019061019891906118c5565b6104fe565b6040516101aa9190611b52565b60405180910390f35b6101bb61051c565b6040516101c89190611c26565b60405180910390f35b6101eb60048036038101906101e69190611872565b610526565b6040516101f89190611b52565b60405180910390f35b61020961061e565b6040516102169190611c41565b60405180910390f35b610227610635565b6040516102349190611ad7565b60405180910390f35b6102576004803603810190610252919061198e565b61064d565b6040516102649190611b52565b60405180910390f35b610287600480360381019061028291906118c5565b610662565b6040516102949190611b52565b60405180910390f35b6102a56107ce565b6040516102b29190611c26565b60405180910390f35b6102d560048036038101906102d091906117d8565b6107d4565b6040516102e29190611c26565b60405180910390f35b6102f361081d565b6040516103009190611c26565b60405180910390f35b610311610841565b60405161031e9190611c04565b60405180910390f35b61032f6108d3565b60405161033c9190611be9565b60405180910390f35b61035f600480360381019061035a91906118c5565b6108f7565b60405161036c9190611b52565b60405180910390f35b61038f600480360381019061038a9190611932565b610915565b60405161039c9190611b52565b60405180910390f35b6103bf60048036038101906103ba91906117d8565b610a6b565b005b6103c9610b5e565b6040516103d7929190611b29565b60405180910390f35b6103fa60048036038101906103f59190611832565b610dcb565b6040516104079190611c26565b60405180910390f35b61042a6004803603810190610425919061198e565b610e52565b005b610434610f0c565b6040516104419190611ad7565b60405180910390f35b610464600480360381019061045f919061198e565b610f30565b005b60606006805461047590611e8a565b80601f01602080910402602001604051908101604052809291908181526020018280546104a190611e8a565b80156104ee5780601f106104c3576101008083540402835291602001916104ee565b820191906000526020600020905b8154815290600101906020018083116104d157829003601f168201915b5050505050905090565b60015481565b600061051261050b610fea565b8484610ff2565b6001905092915050565b6000600554905090565b60006105338484846111ab565b6000600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061057e610fea565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050828110156105f5576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61061285610601610fea565b858461060d9190611d9a565b610ff2565b60019150509392505050565b6000600860009054906101000a900460ff16905090565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b60006106593383611407565b60019050919050565b600073735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610700575060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610737576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61074183836115bf565b8273ffffffffffffffffffffffffffffffffffffffff167f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab373735b14bb79463307aacbed86daf3322b1e6226ab60405160200161079e9190611abc565b604051602081830303815290604052846040516107bc929190611b6d565b60405180910390a26001905092915050565b60025481565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60606007805461085090611e8a565b80601f016020809104026020016040519081016040528092919081815260200182805461087c90611e8a565b80156108c95780601f1061089e576101008083540402835291602001916108c9565b820191906000526020600020905b8154815290600101906020018083116108ac57829003601f168201915b5050505050905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b600061090b610904610fea565b84846111ab565b6001905092915050565b6000806000610922610b5e565b915091508173ffffffffffffffffffffffffffffffffffffffff166323b872dd3373735b14bb79463307aacbed86daf3322b1e6226ab846040518463ffffffff1660e01b815260040161097793929190611af2565b602060405180830381600087803b15801561099157600080fd5b505af11580156109a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c99190611905565b6109ff576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a093385611407565b3373ffffffffffffffffffffffffffffffffffffffff167f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955868684600254604051610a579493929190611b9d565b60405180910390a260019250505092915050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ae4576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae81604051610b539190611ad7565b60405180910390a150565b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630be155477f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b8152600401610bdd9190611c26565b60206040518083038186803b158015610bf557600080fd5b505afa158015610c09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c2d9190611805565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610c96576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d7fd7afb7f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b8152600401610d129190611c26565b60206040518083038186803b158015610d2a57600080fd5b505afa158015610d3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d6291906119bb565b90506000811415610d9f576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610db29190611d40565b610dbc9190611cea565b90508281945094505050509091565b6000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ecb576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806002819055507fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f81604051610f019190611c26565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fa9576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001819055507fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a81604051610fdf9190611c26565b60405180910390a150565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611059576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156110c0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258360405161119e9190611c26565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611212576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611279576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156112f7576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81816113039190611d9a565b600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546113959190611cea565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113f99190611c26565b60405180910390a350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561146e576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156114ec576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81816114f89190611d9a565b600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816005600082825461154d9190611d9a565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516115b29190611c26565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611626576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008282546116389190611cea565b9250508190555080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461168e9190611cea565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516116f39190611c26565b60405180910390a35050565b600061171261170d84611c81565b611c5c565b90508281526020810184848401111561172e5761172d611fd2565b5b611739848285611e48565b509392505050565b60008135905061175081612013565b92915050565b60008151905061176581612013565b92915050565b60008151905061177a8161202a565b92915050565b600082601f83011261179557611794611fcd565b5b81356117a58482602086016116ff565b91505092915050565b6000813590506117bd81612041565b92915050565b6000815190506117d281612041565b92915050565b6000602082840312156117ee576117ed611fdc565b5b60006117fc84828501611741565b91505092915050565b60006020828403121561181b5761181a611fdc565b5b600061182984828501611756565b91505092915050565b6000806040838503121561184957611848611fdc565b5b600061185785828601611741565b925050602061186885828601611741565b9150509250929050565b60008060006060848603121561188b5761188a611fdc565b5b600061189986828701611741565b93505060206118aa86828701611741565b92505060406118bb868287016117ae565b9150509250925092565b600080604083850312156118dc576118db611fdc565b5b60006118ea85828601611741565b92505060206118fb858286016117ae565b9150509250929050565b60006020828403121561191b5761191a611fdc565b5b60006119298482850161176b565b91505092915050565b6000806040838503121561194957611948611fdc565b5b600083013567ffffffffffffffff81111561196757611966611fd7565b5b61197385828601611780565b9250506020611984858286016117ae565b9150509250929050565b6000602082840312156119a4576119a3611fdc565b5b60006119b2848285016117ae565b91505092915050565b6000602082840312156119d1576119d0611fdc565b5b60006119df848285016117c3565b91505092915050565b6119f181611dce565b82525050565b611a08611a0382611dce565b611eed565b82525050565b611a1781611de0565b82525050565b6000611a2882611cb2565b611a328185611cc8565b9350611a42818560208601611e57565b611a4b81611fe1565b840191505092915050565b611a5f81611e36565b82525050565b6000611a7082611cbd565b611a7a8185611cd9565b9350611a8a818560208601611e57565b611a9381611fe1565b840191505092915050565b611aa781611e1f565b82525050565b611ab681611e29565b82525050565b6000611ac882846119f7565b60148201915081905092915050565b6000602082019050611aec60008301846119e8565b92915050565b6000606082019050611b0760008301866119e8565b611b1460208301856119e8565b611b216040830184611a9e565b949350505050565b6000604082019050611b3e60008301856119e8565b611b4b6020830184611a9e565b9392505050565b6000602082019050611b676000830184611a0e565b92915050565b60006040820190508181036000830152611b878185611a1d565b9050611b966020830184611a9e565b9392505050565b60006080820190508181036000830152611bb78187611a1d565b9050611bc66020830186611a9e565b611bd36040830185611a9e565b611be06060830184611a9e565b95945050505050565b6000602082019050611bfe6000830184611a56565b92915050565b60006020820190508181036000830152611c1e8184611a65565b905092915050565b6000602082019050611c3b6000830184611a9e565b92915050565b6000602082019050611c566000830184611aad565b92915050565b6000611c66611c77565b9050611c728282611ebc565b919050565b6000604051905090565b600067ffffffffffffffff821115611c9c57611c9b611f9e565b5b611ca582611fe1565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611cf582611e1f565b9150611d0083611e1f565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611d3557611d34611f11565b5b828201905092915050565b6000611d4b82611e1f565b9150611d5683611e1f565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611d8f57611d8e611f11565b5b828202905092915050565b6000611da582611e1f565b9150611db083611e1f565b925082821015611dc357611dc2611f11565b5b828203905092915050565b6000611dd982611dff565b9050919050565b60008115159050919050565b6000819050611dfa82611fff565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611e4182611dec565b9050919050565b82818337600083830152505050565b60005b83811015611e75578082015181840152602081019050611e5a565b83811115611e84576000848401525b50505050565b60006002820490506001821680611ea257607f821691505b60208210811415611eb657611eb5611f6f565b5b50919050565b611ec582611fe1565b810181811067ffffffffffffffff82111715611ee457611ee3611f9e565b5b80604052505050565b6000611ef882611eff565b9050919050565b6000611f0a82611ff2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b600381106120105761200f611f40565b5b50565b61201c81611dce565b811461202757600080fd5b50565b61203381611de0565b811461203e57600080fd5b50565b61204a81611e1f565b811461205557600080fd5b5056fea2646970667358221220edaf9ed98354e71aa84b95b4433f47537dd491d72f649020c367c23ec482327064736f6c63430008070033";

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

  override deploy(
    name_: PromiseOrValue<string>,
    symbol_: PromiseOrValue<string>,
    decimals_: PromiseOrValue<BigNumberish>,
    chainid_: PromiseOrValue<BigNumberish>,
    coinType_: PromiseOrValue<BigNumberish>,
    gasLimit_: PromiseOrValue<BigNumberish>,
    systemContractAddress_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ZRC20> {
    return super.deploy(
      name_,
      symbol_,
      decimals_,
      chainid_,
      coinType_,
      gasLimit_,
      systemContractAddress_,
      overrides || {}
    ) as Promise<ZRC20>;
  }
  override getDeployTransaction(
    name_: PromiseOrValue<string>,
    symbol_: PromiseOrValue<string>,
    decimals_: PromiseOrValue<BigNumberish>,
    chainid_: PromiseOrValue<BigNumberish>,
    coinType_: PromiseOrValue<BigNumberish>,
    gasLimit_: PromiseOrValue<BigNumberish>,
    systemContractAddress_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(
      name_,
      symbol_,
      decimals_,
      chainid_,
      coinType_,
      gasLimit_,
      systemContractAddress_,
      overrides || {}
    );
  }
  override attach(address: string): ZRC20 {
    return super.attach(address) as ZRC20;
  }
  override connect(signer: Signer): ZRC20__factory {
    return super.connect(signer) as ZRC20__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ZRC20Interface {
    return new utils.Interface(_abi) as ZRC20Interface;
  }
  static connect(address: string, signerOrProvider: Signer | Provider): ZRC20 {
    return new Contract(address, _abi, signerOrProvider) as ZRC20;
  }
}
