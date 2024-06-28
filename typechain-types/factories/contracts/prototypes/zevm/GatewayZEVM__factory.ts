/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  GatewayZEVM,
  GatewayZEVMInterface,
} from "../../../../contracts/prototypes/zevm/GatewayZEVM";

const _abi = [
  {
    inputs: [],
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
    name: "InsufficientZRC20Amount",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidTarget",
    type: "error",
  },
  {
    inputs: [],
    name: "WithdrawalFailed",
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
        internalType: "bytes",
        name: "receiver",
        type: "bytes",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "message",
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
      {
        indexed: false,
        internalType: "bytes",
        name: "message",
        type: "bytes",
      },
    ],
    name: "Withdrawal",
    type: "event",
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
    inputs: [
      {
        internalType: "bytes",
        name: "receiver",
        type: "bytes",
      },
      {
        internalType: "bytes",
        name: "message",
        type: "bytes",
      },
    ],
    name: "call",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "zrc20",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "target",
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
        components: [
          {
            internalType: "bytes",
            name: "origin",
            type: "bytes",
          },
          {
            internalType: "address",
            name: "sender",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "chainID",
            type: "uint256",
          },
        ],
        internalType: "struct zContext",
        name: "context",
        type: "tuple",
      },
      {
        internalType: "address",
        name: "zrc20",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "target",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "message",
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
        components: [
          {
            internalType: "bytes",
            name: "origin",
            type: "bytes",
          },
          {
            internalType: "address",
            name: "sender",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "chainID",
            type: "uint256",
          },
        ],
        internalType: "struct zContext",
        name: "context",
        type: "tuple",
      },
      {
        internalType: "address",
        name: "zrc20",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "target",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "message",
        type: "bytes",
      },
    ],
    name: "execute",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
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
    inputs: [
      {
        internalType: "bytes",
        name: "receiver",
        type: "bytes",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "zrc20",
        type: "address",
      },
    ],
    name: "withdraw",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes",
        name: "receiver",
        type: "bytes",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "zrc20",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "message",
        type: "bytes",
      },
    ],
    name: "withdrawAndCall",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

const _bytecode =
  "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b8152503480156200004757600080fd5b50620000586200005e60201b60201c565b62000208565b600060019054906101000a900460ff1615620000b1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a8906200015c565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff1614620001225760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200011991906200017e565b60405180910390a15b565b6000620001336027836200019b565b91506200014082620001b9565b604082019050919050565b6200015681620001ac565b82525050565b60006020820190508181036000830152620001778162000124565b9050919050565b60006020820190506200019560008301846200014b565b92915050565b600082825260208201905092915050565b600060ff82169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60805160601c612c076200024360003960008181610447015281816104d6015281816105e80152818161067701526107270152612c076000f3fe6080604052600436106100dd5760003560e01c80637993c1e01161007f578063bcf7f32b11610059578063bcf7f32b14610251578063c39aca371461027a578063f2fde38b146102a3578063f45346dc146102cc576100dd565b80637993c1e0146101e65780638129fc1c1461020f5780638da5cb5b14610226576100dd565b80633ce4a5bc116100bb5780633ce4a5bc1461015d5780634f1ef2861461018857806352d1902d146101a4578063715018a6146101cf576100dd565b80630ac7c44c146100e2578063135390f91461010b5780633659cfe614610134575b600080fd5b3480156100ee57600080fd5b5061010960048036038101906101049190611c2a565b6102f5565b005b34801561011757600080fd5b50610132600480360381019061012d9190611ca6565b610360565b005b34801561014057600080fd5b5061015b60048036038101906101569190611ab4565b610445565b005b34801561016957600080fd5b506101726105ce565b60405161017f9190612218565b60405180910390f35b6101a2600480360381019061019d9190611ae1565b6105e6565b005b3480156101b057600080fd5b506101b9610723565b6040516101c69190612293565b60405180910390f35b3480156101db57600080fd5b506101e46107dc565b005b3480156101f257600080fd5b5061020d60048036038101906102089190611d15565b6107f0565b005b34801561021b57600080fd5b506102246108db565b005b34801561023257600080fd5b5061023b610a21565b6040516102489190612218565b60405180910390f35b34801561025d57600080fd5b5061027860048036038101906102739190611db9565b610a4b565b005b34801561028657600080fd5b506102a1600480360381019061029c9190611db9565b610b3f565b005b3480156102af57600080fd5b506102ca60048036038101906102c59190611ab4565b610d71565b005b3480156102d857600080fd5b506102f360048036038101906102ee9190611b7d565b610df5565b005b826040516103039190612201565b60405180910390203373ffffffffffffffffffffffffffffffffffffffff167f2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f84846040516103539291906122ae565b60405180910390a3505050565b600061036c8383610fb1565b90503373ffffffffffffffffffffffffffffffffffffffff167f1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d8585848673ffffffffffffffffffffffffffffffffffffffff16634d8943bb6040518163ffffffff1660e01b815260040160206040518083038186803b1580156103ef57600080fd5b505afa158015610403573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104279190611e6f565b6040516104379493929190612335565b60405180910390a250505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614156104d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104cb906123f1565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610513611237565b73ffffffffffffffffffffffffffffffffffffffff1614610569576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056090612411565b60405180910390fd5b6105728161128e565b6105cb81600067ffffffffffffffff811115610591576105906127c0565b5b6040519080825280601f01601f1916602001820160405280156105c35781602001600182028036833780820191505090505b506000611299565b50565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610675576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066c906123f1565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166106b4611237565b73ffffffffffffffffffffffffffffffffffffffff161461070a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070190612411565b60405180910390fd5b6107138261128e565b61071f82826001611299565b5050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16146107b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107aa90612431565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b6107e4611416565b6107ee6000611494565b565b60006107fc8585610fb1565b90503373ffffffffffffffffffffffffffffffffffffffff167f1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d8787848873ffffffffffffffffffffffffffffffffffffffff16634d8943bb6040518163ffffffff1660e01b815260040160206040518083038186803b15801561087f57600080fd5b505afa158015610893573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b79190611e6f565b88886040516108cb969594939291906122d2565b60405180910390a2505050505050565b60008060019054906101000a900460ff1615905080801561090c5750600160008054906101000a900460ff1660ff16105b80610939575061091b3061155a565b1580156109385750600160008054906101000a900460ff1660ff16145b5b610978576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096f90612471565b60405180910390fd5b60016000806101000a81548160ff021916908360ff16021790555080156109b5576001600060016101000a81548160ff0219169083151502179055505b6109bd61157d565b6109c56115d6565b8015610a1e5760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610a159190612394565b60405180910390a15b50565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac4576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1663de43156e87878786866040518663ffffffff1660e01b8152600401610b05959493929190612531565b600060405180830381600087803b158015610b1f57600080fd5b505af1158015610b33573d6000803e3d6000fd5b50505050505050505050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bb8576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161480610c3157503073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610c68576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff166347e7ef2484866040518363ffffffff1660e01b8152600401610ca392919061226a565b602060405180830381600087803b158015610cbd57600080fd5b505af1158015610cd1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf59190611bd0565b508273ffffffffffffffffffffffffffffffffffffffff1663de43156e87878786866040518663ffffffff1660e01b8152600401610d37959493929190612531565b600060405180830381600087803b158015610d5157600080fd5b505af1158015610d65573d6000803e3d6000fd5b50505050505050505050565b610d79611416565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610de9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610de0906123d1565b60405180910390fd5b610df281611494565b50565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e6e576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161480610ee757503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15610f1e576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166347e7ef2482846040518363ffffffff1660e01b8152600401610f5992919061226a565b602060405180830381600087803b158015610f7357600080fd5b505af1158015610f87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fab9190611bd0565b50505050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1663d9eeebed6040518163ffffffff1660e01b8152600401604080518083038186803b158015610ffb57600080fd5b505afa15801561100f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110339190611b3d565b915091508173ffffffffffffffffffffffffffffffffffffffff166323b872dd3373735b14bb79463307aacbed86daf3322b1e6226ab846040518463ffffffff1660e01b815260040161108893929190612233565b602060405180830381600087803b1580156110a257600080fd5b505af11580156110b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110da9190611bd0565b611110576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330886040518463ffffffff1660e01b815260040161114d93929190612233565b602060405180830381600087803b15801561116757600080fd5b505af115801561117b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061119f9190611bd0565b508373ffffffffffffffffffffffffffffffffffffffff166342966c68866040518263ffffffff1660e01b81526004016111d99190612586565b602060405180830381600087803b1580156111f357600080fd5b505af1158015611207573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061122b9190611bd0565b50809250505092915050565b60006112657f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b611627565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611296611416565b50565b6112c57f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b611631565b60000160009054906101000a900460ff16156112e9576112e48361163b565b611411565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561132f57600080fd5b505afa92505050801561136057506040513d601f19601f8201168201806040525081019061135d9190611bfd565b60015b61139f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161139690612491565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114611404576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113fb90612451565b60405180910390fd5b506114108383836116f4565b5b505050565b61141e611720565b73ffffffffffffffffffffffffffffffffffffffff1661143c610a21565b73ffffffffffffffffffffffffffffffffffffffff1614611492576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611489906124d1565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff166115cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115c390612511565b60405180910390fd5b6115d4611728565b565b600060019054906101000a900460ff16611625576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161161c90612511565b60405180910390fd5b565b6000819050919050565b6000819050919050565b6116448161155a565b611683576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161167a906124b1565b60405180910390fd5b806116b07f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b611627565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6116fd83611789565b60008251118061170a5750805b1561171b5761171983836117d8565b505b505050565b600033905090565b600060019054906101000a900460ff16611777576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161176e90612511565b60405180910390fd5b611787611782611720565b611494565b565b6117928161163b565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b60606117fd8383604051806060016040528060278152602001612bab60279139611805565b905092915050565b60606000808573ffffffffffffffffffffffffffffffffffffffff168560405161182f9190612201565b600060405180830381855af49150503d806000811461186a576040519150601f19603f3d011682016040523d82523d6000602084013e61186f565b606091505b50915091506118808683838761188b565b925050509392505050565b606083156118ee576000835114156118e6576118a68561155a565b6118e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118dc906124f1565b60405180910390fd5b5b8290506118f9565b6118f88383611901565b5b949350505050565b6000825111156119145781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161194891906123af565b60405180910390fd5b600061196461195f846125c6565b6125a1565b9050828152602081018484840111156119805761197f61280d565b5b61198b84828561274d565b509392505050565b6000813590506119a281612b4e565b92915050565b6000815190506119b781612b4e565b92915050565b6000815190506119cc81612b65565b92915050565b6000815190506119e181612b7c565b92915050565b60008083601f8401126119fd576119fc6127f9565b5b8235905067ffffffffffffffff811115611a1a57611a196127f4565b5b602083019150836001820283011115611a3657611a35612808565b5b9250929050565b600082601f830112611a5257611a516127f9565b5b8135611a62848260208601611951565b91505092915050565b600060608284031215611a8157611a806127fe565b5b81905092915050565b600081359050611a9981612b93565b92915050565b600081519050611aae81612b93565b92915050565b600060208284031215611aca57611ac961281c565b5b6000611ad884828501611993565b91505092915050565b60008060408385031215611af857611af761281c565b5b6000611b0685828601611993565b925050602083013567ffffffffffffffff811115611b2757611b26612812565b5b611b3385828601611a3d565b9150509250929050565b60008060408385031215611b5457611b5361281c565b5b6000611b62858286016119a8565b9250506020611b7385828601611a9f565b9150509250929050565b600080600060608486031215611b9657611b9561281c565b5b6000611ba486828701611993565b9350506020611bb586828701611a8a565b9250506040611bc686828701611993565b9150509250925092565b600060208284031215611be657611be561281c565b5b6000611bf4848285016119bd565b91505092915050565b600060208284031215611c1357611c1261281c565b5b6000611c21848285016119d2565b91505092915050565b600080600060408486031215611c4357611c4261281c565b5b600084013567ffffffffffffffff811115611c6157611c60612812565b5b611c6d86828701611a3d565b935050602084013567ffffffffffffffff811115611c8e57611c8d612812565b5b611c9a868287016119e7565b92509250509250925092565b600080600060608486031215611cbf57611cbe61281c565b5b600084013567ffffffffffffffff811115611cdd57611cdc612812565b5b611ce986828701611a3d565b9350506020611cfa86828701611a8a565b9250506040611d0b86828701611993565b9150509250925092565b600080600080600060808688031215611d3157611d3061281c565b5b600086013567ffffffffffffffff811115611d4f57611d4e612812565b5b611d5b88828901611a3d565b9550506020611d6c88828901611a8a565b9450506040611d7d88828901611993565b935050606086013567ffffffffffffffff811115611d9e57611d9d612812565b5b611daa888289016119e7565b92509250509295509295909350565b60008060008060008060a08789031215611dd657611dd561281c565b5b600087013567ffffffffffffffff811115611df457611df3612812565b5b611e0089828a01611a6b565b9650506020611e1189828a01611993565b9550506040611e2289828a01611a8a565b9450506060611e3389828a01611993565b935050608087013567ffffffffffffffff811115611e5457611e53612812565b5b611e6089828a016119e7565b92509250509295509295509295565b600060208284031215611e8557611e8461281c565b5b6000611e9384828501611a9f565b91505092915050565b611ea5816126dc565b82525050565b611eb4816126dc565b82525050565b611ec3816126fa565b82525050565b6000611ed5838561260d565b9350611ee283858461274d565b611eeb83612821565b840190509392505050565b6000611f02838561261e565b9350611f0f83858461274d565b611f1883612821565b840190509392505050565b6000611f2e826125f7565b611f38818561261e565b9350611f4881856020860161275c565b611f5181612821565b840191505092915050565b6000611f67826125f7565b611f71818561262f565b9350611f8181856020860161275c565b80840191505092915050565b611f968161273b565b82525050565b6000611fa782612602565b611fb1818561263a565b9350611fc181856020860161275c565b611fca81612821565b840191505092915050565b6000611fe260268361263a565b9150611fed82612832565b604082019050919050565b6000612005602c8361263a565b915061201082612881565b604082019050919050565b6000612028602c8361263a565b9150612033826128d0565b604082019050919050565b600061204b60388361263a565b91506120568261291f565b604082019050919050565b600061206e60298361263a565b91506120798261296e565b604082019050919050565b6000612091602e8361263a565b915061209c826129bd565b604082019050919050565b60006120b4602e8361263a565b91506120bf82612a0c565b604082019050919050565b60006120d7602d8361263a565b91506120e282612a5b565b604082019050919050565b60006120fa60208361263a565b915061210582612aaa565b602082019050919050565b600061211d60008361261e565b915061212882612ad3565b600082019050919050565b6000612140601d8361263a565b915061214b82612ad6565b602082019050919050565b6000612163602b8361263a565b915061216e82612aff565b604082019050919050565b60006060830161218c6000840184612662565b858303600087015261219f838284611ec9565b925050506121b0602084018461264b565b6121bd6020860182611e9c565b506121cb60408401846126c5565b6121d860408601826121e3565b508091505092915050565b6121ec81612724565b82525050565b6121fb81612724565b82525050565b600061220d8284611f5c565b915081905092915050565b600060208201905061222d6000830184611eab565b92915050565b60006060820190506122486000830186611eab565b6122556020830185611eab565b61226260408301846121f2565b949350505050565b600060408201905061227f6000830185611eab565b61228c60208301846121f2565b9392505050565b60006020820190506122a86000830184611eba565b92915050565b600060208201905081810360008301526122c9818486611ef6565b90509392505050565b600060a08201905081810360008301526122ec8189611f23565b90506122fb60208301886121f2565b61230860408301876121f2565b61231560608301866121f2565b8181036080830152612328818486611ef6565b9050979650505050505050565b600060a082019050818103600083015261234f8187611f23565b905061235e60208301866121f2565b61236b60408301856121f2565b61237860608301846121f2565b818103608083015261238981612110565b905095945050505050565b60006020820190506123a96000830184611f8d565b92915050565b600060208201905081810360008301526123c98184611f9c565b905092915050565b600060208201905081810360008301526123ea81611fd5565b9050919050565b6000602082019050818103600083015261240a81611ff8565b9050919050565b6000602082019050818103600083015261242a8161201b565b9050919050565b6000602082019050818103600083015261244a8161203e565b9050919050565b6000602082019050818103600083015261246a81612061565b9050919050565b6000602082019050818103600083015261248a81612084565b9050919050565b600060208201905081810360008301526124aa816120a7565b9050919050565b600060208201905081810360008301526124ca816120ca565b9050919050565b600060208201905081810360008301526124ea816120ed565b9050919050565b6000602082019050818103600083015261250a81612133565b9050919050565b6000602082019050818103600083015261252a81612156565b9050919050565b6000608082019050818103600083015261254b8188612179565b905061255a6020830187611eab565b61256760408301866121f2565b818103606083015261257a818486611ef6565b90509695505050505050565b600060208201905061259b60008301846121f2565b92915050565b60006125ab6125bc565b90506125b7828261278f565b919050565b6000604051905090565b600067ffffffffffffffff8211156125e1576125e06127c0565b5b6125ea82612821565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061265a6020840184611993565b905092915050565b6000808335600160200384360303811261267f5761267e612817565b5b83810192508235915060208301925067ffffffffffffffff8211156126a7576126a66127ef565b5b6001820236038413156126bd576126bc612803565b5b509250929050565b60006126d46020840184611a8a565b905092915050565b60006126e782612704565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006127468261272e565b9050919050565b82818337600083830152505050565b60005b8381101561277a57808201518184015260208101905061275f565b83811115612789576000848401525b50505050565b61279882612821565b810181811067ffffffffffffffff821117156127b7576127b66127c0565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b612b57816126dc565b8114612b6257600080fd5b50565b612b6e816126ee565b8114612b7957600080fd5b50565b612b85816126fa565b8114612b9057600080fd5b50565b612b9c81612724565b8114612ba757600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220e84bd62cddf8d507bfa43de0a9aa436cfa921f710c9c070bc941f3c0fda67be364736f6c63430008070033";

type GatewayZEVMConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: GatewayZEVMConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class GatewayZEVM__factory extends ContractFactory {
  constructor(...args: GatewayZEVMConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<GatewayZEVM> {
    return super.deploy(overrides || {}) as Promise<GatewayZEVM>;
  }
  override getDeployTransaction(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  override attach(address: string): GatewayZEVM {
    return super.attach(address) as GatewayZEVM;
  }
  override connect(signer: Signer): GatewayZEVM__factory {
    return super.connect(signer) as GatewayZEVM__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): GatewayZEVMInterface {
    return new utils.Interface(_abi) as GatewayZEVMInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): GatewayZEVM {
    return new Contract(address, _abi, signerOrProvider) as GatewayZEVM;
  }
}
