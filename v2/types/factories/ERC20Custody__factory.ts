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
import type { ERC20Custody, ERC20CustodyInterface } from "../ERC20Custody";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "gateway_",
        type: "address",
        internalType: "address",
      },
      {
        name: "tssAddress_",
        type: "address",
        internalType: "address",
      },
      {
        name: "admin_",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "DEFAULT_ADMIN_ROLE",
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
    name: "PAUSER_ROLE",
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
    name: "WHITELISTER_ROLE",
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
    name: "WITHDRAWER_ROLE",
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
    name: "deposit",
    inputs: [
      {
        name: "recipient",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "asset",
        type: "address",
        internalType: "contract IERC20",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
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
        internalType: "contract IGatewayEVM",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "getRoleAdmin",
    inputs: [
      {
        name: "role",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
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
    name: "grantRole",
    inputs: [
      {
        name: "role",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "account",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "hasRole",
    inputs: [
      {
        name: "role",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "account",
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
    name: "pause",
    inputs: [],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "paused",
    inputs: [],
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
    name: "renounceRole",
    inputs: [
      {
        name: "role",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "callerConfirmation",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "revokeRole",
    inputs: [
      {
        name: "role",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "account",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "setSupportsLegacy",
    inputs: [
      {
        name: "_supportsLegacy",
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
    name: "supportsLegacy",
    inputs: [],
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
    name: "tssAddress",
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
    name: "unpause",
    inputs: [],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "unwhitelist",
    inputs: [
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "updateTSSAddress",
    inputs: [
      {
        name: "newTSSAddress",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "whitelist",
    inputs: [
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "whitelisted",
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
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "withdraw",
    inputs: [
      {
        name: "to",
        type: "address",
        internalType: "address",
      },
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "withdrawAndCall",
    inputs: [
      {
        name: "to",
        type: "address",
        internalType: "address",
      },
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
      {
        name: "amount",
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
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "withdrawAndRevert",
    inputs: [
      {
        name: "to",
        type: "address",
        internalType: "address",
      },
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "data",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "revertContext",
        type: "tuple",
        internalType: "struct RevertContext",
        components: [
          {
            name: "sender",
            type: "address",
            internalType: "address",
          },
          {
            name: "asset",
            type: "address",
            internalType: "address",
          },
          {
            name: "amount",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "revertMessage",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    name: "Deposited",
    inputs: [
      {
        name: "recipient",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "asset",
        type: "address",
        indexed: true,
        internalType: "contract IERC20",
      },
      {
        name: "amount",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "message",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Paused",
    inputs: [
      {
        name: "account",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "RoleAdminChanged",
    inputs: [
      {
        name: "role",
        type: "bytes32",
        indexed: true,
        internalType: "bytes32",
      },
      {
        name: "previousAdminRole",
        type: "bytes32",
        indexed: true,
        internalType: "bytes32",
      },
      {
        name: "newAdminRole",
        type: "bytes32",
        indexed: true,
        internalType: "bytes32",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "RoleGranted",
    inputs: [
      {
        name: "role",
        type: "bytes32",
        indexed: true,
        internalType: "bytes32",
      },
      {
        name: "account",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "sender",
        type: "address",
        indexed: true,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "RoleRevoked",
    inputs: [
      {
        name: "role",
        type: "bytes32",
        indexed: true,
        internalType: "bytes32",
      },
      {
        name: "account",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "sender",
        type: "address",
        indexed: true,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Unpaused",
    inputs: [
      {
        name: "account",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Unwhitelisted",
    inputs: [
      {
        name: "token",
        type: "address",
        indexed: true,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "UpdatedCustodyTSSAddress",
    inputs: [
      {
        name: "newTSSAddress",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Whitelisted",
    inputs: [
      {
        name: "token",
        type: "address",
        indexed: true,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Withdrawn",
    inputs: [
      {
        name: "to",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "token",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "WithdrawnAndCalled",
    inputs: [
      {
        name: "to",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "token",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "data",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "WithdrawnAndReverted",
    inputs: [
      {
        name: "to",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "token",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "data",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "revertContext",
        type: "tuple",
        indexed: false,
        internalType: "struct RevertContext",
        components: [
          {
            name: "sender",
            type: "address",
            internalType: "address",
          },
          {
            name: "asset",
            type: "address",
            internalType: "address",
          },
          {
            name: "amount",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "revertMessage",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
    ],
    anonymous: false,
  },
  {
    type: "error",
    name: "AccessControlBadConfirmation",
    inputs: [],
  },
  {
    type: "error",
    name: "AccessControlUnauthorizedAccount",
    inputs: [
      {
        name: "account",
        type: "address",
        internalType: "address",
      },
      {
        name: "neededRole",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
  },
  {
    type: "error",
    name: "AddressEmptyCode",
    inputs: [
      {
        name: "target",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "AddressInsufficientBalance",
    inputs: [
      {
        name: "account",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "EnforcedPause",
    inputs: [],
  },
  {
    type: "error",
    name: "ExpectedPause",
    inputs: [],
  },
  {
    type: "error",
    name: "FailedInnerCall",
    inputs: [],
  },
  {
    type: "error",
    name: "LegacyMethodsNotSupported",
    inputs: [],
  },
  {
    type: "error",
    name: "NotWhitelisted",
    inputs: [],
  },
  {
    type: "error",
    name: "ReentrancyGuardReentrantCall",
    inputs: [],
  },
  {
    type: "error",
    name: "SafeERC20FailedOperation",
    inputs: [
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "ZeroAddress",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x60a060405234801561001057600080fd5b50604051611e03380380611e0383398101604081905261002f916101fd565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03838116608052600480546001600160a01b0319169184169190911790556100bb60008261014c565b506100e67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261014c565b506101117f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361014c565b5061012a600080516020611de38339815191528261014c565b50610143600080516020611de38339815191528361014c565b50505050610240565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166101d75760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016101db565b5060005b92915050565b80516001600160a01b03811681146101f857600080fd5b919050565b60008060006060848603121561021257600080fd5b61021b846101e1565b9250610229602085016101e1565b9150610237604085016101e1565b90509250925092565b608051611b6c610277600039600081816101d501528181610574015281816105c90152818161099601526109eb0152611b6c6000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c806385f438c1116100ee578063a217fddf11610097578063d9caed1211610071578063d9caed12146103e0578063e609055e146103f3578063e63ab1e914610406578063eab103df1461042d57600080fd5b8063a217fddf146103a2578063d547741f146103aa578063d936547e146103bd57600080fd5b806399a3c356116100c857806399a3c356146103695780639a5904271461037c5780639b19251a1461038f57600080fd5b806385f438c1146102f657806391d148541461031d578063950837aa1461035657600080fd5b806336568abe116101505780635b1125911161012a5780635b112591146102d05780635c975abb146102e35780638456cb59146102ee57600080fd5b806336568abe1461028e5780633f4ba83a146102a1578063570618e1146102a957600080fd5b8063248a9ca311610181578063248a9ca314610224578063252f07bf146102565780632f2ff15d1461027b57600080fd5b806301ffc9a7146101a8578063116191b6146101d057806321fc65f21461020f575b600080fd5b6101bb6101b636600461155e565b610440565b60405190151581526020015b60405180910390f35b6101f77f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101c7565b61022261021d3660046115fe565b6104d9565b005b610248610232366004611671565b6000908152600160208190526040909120015490565b6040519081526020016101c7565b6004546101bb9074010000000000000000000000000000000000000000900460ff1681565b61022261028936600461168a565b610699565b61022261029c36600461168a565b6106c5565b610222610716565b6102487f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b6004546101f7906001600160a01b031681565b60025460ff166101bb565b61022261074b565b6102487f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101bb61032b36600461168a565b60009182526001602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6102226103643660046116ba565b61077d565b6102226103773660046116d7565b6108fb565b61022261038a3660046116ba565b610ac0565b61022261039d3660046116ba565b610b74565b610248600081565b6102226103b836600461168a565b610c2b565b6101bb6103cb3660046116ba565b60036020526000908152604090205460ff1681565b6102226103ee36600461177a565b610c51565b6102226104013660046117bb565b610d49565b6102487f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b61022261043b36600461185a565b610f75565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104d357507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104e1610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461050b8161100e565b610513611018565b6001600160a01b03851660009081526003602052604090205460ff16610565576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105996001600160a01b0386167f000000000000000000000000000000000000000000000000000000000000000086611057565b6040517f5131ab590000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906106069088908a908990899089906004016118c0565b600060405180830381600087803b15801561062057600080fd5b505af1158015610634573d6000803e3d6000fd5b50505050846001600160a01b0316866001600160a01b03167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d586868660405161067f93929190611903565b60405180910390a3506106926001600055565b5050505050565b600082815260016020819052604090912001546106b58161100e565b6106bf83836110cb565b50505050565b6001600160a01b0381163314610707576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610711828261115e565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107408161100e565b6107486111e5565b50565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107758161100e565b610748611237565b60006107888161100e565b6001600160a01b0382166107c8576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004546107ff907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4906001600160a01b031661115e565b50600454610837907f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a906001600160a01b031661115e565b506108627f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4836110cb565b5061088d7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a836110cb565b50600480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040519081527f086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba9060200160405180910390a15050565b610903610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461092d8161100e565b610935611018565b6001600160a01b03861660009081526003602052604090205460ff16610987576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109bb6001600160a01b0387167f000000000000000000000000000000000000000000000000000000000000000087611057565b6040517faa0c0fc10000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063aa0c0fc190610a2a9089908b908a908a908a908a906004016119db565b600060405180830381600087803b158015610a4457600080fd5b505af1158015610a58573d6000803e3d6000fd5b50505050856001600160a01b0316876001600160a01b03167f7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb972187878787604051610aa59493929190611a32565b60405180910390a350610ab86001600055565b505050505050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a610aea8161100e565b6001600160a01b038216610b2a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a610b9e8161100e565b6001600160a01b038216610bde576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b60008281526001602081905260409091200154610c478161100e565b6106bf838361115e565b610c59610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610c838161100e565b610c8b611018565b6001600160a01b03831660009081526003602052604090205460ff16610cdd576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cf16001600160a01b0384168584611057565b826001600160a01b0316846001600160a01b03167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610d3691815260200190565b60405180910390a3506107116001600055565b610d51610fcb565b610d59611018565b60045474010000000000000000000000000000000000000000900460ff16610dad576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03841660009081526003602052604090205460ff16610dff576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610e5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e839190611a5e565b9050610e9a6001600160a01b038616333087611274565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa158015610f21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f459190611a5e565b610f4f9190611a77565b8787604051610f62959493929190611ab1565b60405180910390a250610ab86001600055565b6000610f808161100e565b506004805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b600260005403611007576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61074881336112ad565b60025460ff1615611055576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6040516001600160a01b0383811660248301526044820183905261071191859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611324565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166111565760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016104d3565b5060006104d3565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff16156111565760008381526001602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104d3565b6111ed6113a0565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b61123f611018565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861121a3390565b6040516001600160a01b0384811660248301528381166044830152606482018390526106bf9186918216906323b872dd90608401611084565b60008281526001602090815260408083206001600160a01b038516845290915290205460ff16611320576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044015b60405180910390fd5b5050565b60006113396001600160a01b038416836113dc565b9050805160001415801561135e57508080602001905181019061135c9190611aea565b155b15610711576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611317565b60025460ff16611055576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606113ea838360006113f1565b9392505050565b60608147101561142f576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611317565b600080856001600160a01b0316848660405161144b9190611b07565b60006040518083038185875af1925050503d8060008114611488576040519150601f19603f3d011682016040523d82523d6000602084013e61148d565b606091505b509150915061149d8683836114a7565b9695505050505050565b6060826114bc576114b78261151c565b6113ea565b81511580156114d357506001600160a01b0384163b155b15611515576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611317565b50806113ea565b80511561152c5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561157057600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146113ea57600080fd5b6001600160a01b038116811461074857600080fd5b60008083601f8401126115c757600080fd5b50813567ffffffffffffffff8111156115df57600080fd5b6020830191508360208285010111156115f757600080fd5b9250929050565b60008060008060006080868803121561161657600080fd5b8535611621816115a0565b94506020860135611631816115a0565b935060408601359250606086013567ffffffffffffffff81111561165457600080fd5b611660888289016115b5565b969995985093965092949392505050565b60006020828403121561168357600080fd5b5035919050565b6000806040838503121561169d57600080fd5b8235915060208301356116af816115a0565b809150509250929050565b6000602082840312156116cc57600080fd5b81356113ea816115a0565b60008060008060008060a087890312156116f057600080fd5b86356116fb816115a0565b9550602087013561170b816115a0565b945060408701359350606087013567ffffffffffffffff81111561172e57600080fd5b61173a89828a016115b5565b909450925050608087013567ffffffffffffffff81111561175a57600080fd5b87016080818a03121561176c57600080fd5b809150509295509295509295565b60008060006060848603121561178f57600080fd5b833561179a816115a0565b925060208401356117aa816115a0565b929592945050506040919091013590565b600080600080600080608087890312156117d457600080fd5b863567ffffffffffffffff8111156117eb57600080fd5b6117f789828a016115b5565b909750955050602087013561180b816115a0565b935060408701359250606087013567ffffffffffffffff81111561182e57600080fd5b61183a89828a016115b5565b979a9699509497509295939492505050565b801515811461074857600080fd5b60006020828403121561186c57600080fd5b81356113ea8161184c565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b03861681526001600160a01b03851660208201528360408201526080606082015260006118f8608083018486611877565b979650505050505050565b83815260406020820152600061191d604083018486611877565b95945050505050565b60008135611933816115a0565b6001600160a01b03168352602082013561194c816115a0565b6001600160a01b03166020840152604082810135908401526060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261199a57600080fd5b820160208101903567ffffffffffffffff8111156119b757600080fd5b8036038213156119c657600080fd5b6080606086015261191d608086018284611877565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a060608201526000611a1360a083018587611877565b8281036080840152611a258185611926565b9998505050505050505050565b848152606060208201526000611a4c606083018587611877565b82810360408401526118f88185611926565b600060208284031215611a7057600080fd5b5051919050565b818103818111156104d3577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b606081526000611ac5606083018789611877565b8560208401528281036040840152611ade818587611877565b98975050505050505050565b600060208284031215611afc57600080fd5b81516113ea8161184c565b6000825160005b81811015611b285760208186018101518583015201611b0e565b50600092019182525091905056fea26469706673582212208d8c335f9d1dd65279a2dcfe126916b06e449663af5f38182aa9e1d5612b9ff164736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a";

type ERC20CustodyConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: ERC20CustodyConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class ERC20Custody__factory extends ContractFactory {
  constructor(...args: ERC20CustodyConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    gateway_: AddressLike,
    tssAddress_: AddressLike,
    admin_: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(
      gateway_,
      tssAddress_,
      admin_,
      overrides || {}
    );
  }
  override deploy(
    gateway_: AddressLike,
    tssAddress_: AddressLike,
    admin_: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(
      gateway_,
      tssAddress_,
      admin_,
      overrides || {}
    ) as Promise<
      ERC20Custody & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): ERC20Custody__factory {
    return super.connect(runner) as ERC20Custody__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ERC20CustodyInterface {
    return new Interface(_abi) as ERC20CustodyInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): ERC20Custody {
    return new Contract(address, _abi, runner) as unknown as ERC20Custody;
  }
}
