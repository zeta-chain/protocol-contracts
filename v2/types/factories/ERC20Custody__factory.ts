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
            name: "asset",
            type: "address",
            internalType: "address",
          },
          {
            name: "amount",
            type: "uint64",
            internalType: "uint64",
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
        name: "token",
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
        name: "token",
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
        name: "token",
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
            name: "asset",
            type: "address",
            internalType: "address",
          },
          {
            name: "amount",
            type: "uint64",
            internalType: "uint64",
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
  "0x60a060405234801561001057600080fd5b506040516119ca3803806119ca83398101604081905261002f916101e5565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0383166080526100a3600082610134565b506100ce7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610134565b506100f97f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610134565b506101126000805160206119aa83398151915282610134565b5061012b6000805160206119aa83398151915283610134565b50505050610228565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166101bf5760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016101c3565b5060005b92915050565b80516001600160a01b03811681146101e057600080fd5b919050565b6000806000606084860312156101fa57600080fd5b610203846101c9565b9250610211602085016101c9565b915061021f604085016101c9565b90509250925092565b60805161174b61025f6000396000818161019e01528181610500015281816105620152818161099101526109f3015261174b6000f3fe608060405234801561001057600080fd5b506004361061016c5760003560e01c806385f438c1116100cd578063c709ab6e11610081578063d936547e11610066578063d936547e14610355578063d9caed1214610378578063e63ab1e91461038b57600080fd5b8063c709ab6e1461032f578063d547741f1461034257600080fd5b80639a590427116100b25780639a590427146103015780639b19251a14610314578063a217fddf1461032757600080fd5b806385f438c11461029457806391d14854146102bb57600080fd5b806336568abe11610124578063570618e111610109578063570618e11461025a5780635c975abb146102815780638456cb591461028c57600080fd5b806336568abe1461023f5780633f4ba83a1461025257600080fd5b806321fc65f21161015557806321fc65f2146101e5578063248a9ca3146101fa5780632f2ff15d1461022c57600080fd5b806301ffc9a714610171578063116191b614610199575b600080fd5b61018461017f366004611243565b6103b2565b60405190151581526020015b60405180910390f35b6101c07f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610190565b6101f86101f33660046112f7565b61044b565b005b61021e610208366004611366565b6000908152600160208190526040909120015490565b604051908152602001610190565b6101f861023a36600461137f565b61064c565b6101f861024d36600461137f565b610678565b6101f86106d6565b61021e7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b60025460ff16610184565b6101f861070b565b61021e7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101846102c936600461137f565b600091825260016020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6101f861030f3660046113ab565b61073d565b6101f86103223660046113ab565b61080b565b61021e600081565b6101f861033d3660046113c6565b6108dc565b6101f861035036600461137f565b610ae2565b6101846103633660046113ab565b60036020526000908152604090205460ff1681565b6101f8610386366004611465565b610b08565b61021e7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061044557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b610453610c34565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461047d81610c77565b610485610c81565b73ffffffffffffffffffffffffffffffffffffffff851660009081526003602052604090205460ff166104e4576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61052573ffffffffffffffffffffffffffffffffffffffff86167f000000000000000000000000000000000000000000000000000000000000000086610cc0565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab599061059f9088908a908990899089906004016114eb565b600060405180830381600087803b1580156105b957600080fd5b505af11580156105cd573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d586868660405161063293929190611548565b60405180910390a3506106456001600055565b5050505050565b6000828152600160208190526040909120015461066881610c77565b6106728383610d4d565b50505050565b73ffffffffffffffffffffffffffffffffffffffff811633146106c7576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106d18282610dfa565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61070081610c77565b610708610e9b565b50565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61073581610c77565b610708610efa565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a61076781610c77565b73ffffffffffffffffffffffffffffffffffffffff82166107b4576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a61083581610c77565b73ffffffffffffffffffffffffffffffffffffffff8216610882576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b6108e4610c34565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461090e81610c77565b610916610c81565b73ffffffffffffffffffffffffffffffffffffffff861660009081526003602052604090205460ff16610975576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109b673ffffffffffffffffffffffffffffffffffffffff87167f000000000000000000000000000000000000000000000000000000000000000087610cc0565b6040517fd0b492c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610a329089908b908a908a908a908a90600401611627565b600060405180830381600087803b158015610a4c57600080fd5b505af1158015610a60573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b87878787604051610ac79493929190611698565b60405180910390a350610ada6001600055565b505050505050565b60008281526001602081905260409091200154610afe81610c77565b6106728383610dfa565b610b10610c34565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610b3a81610c77565b610b42610c81565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604090205460ff16610ba1576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610bc273ffffffffffffffffffffffffffffffffffffffff84168584610cc0565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610c2191815260200190565b60405180910390a3506106d16001600055565b600260005403610c70576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6107088133610f37565b60025460ff1615610cbe576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526106d1908490610fc8565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610df257600083815260016020818152604080842073ffffffffffffffffffffffffffffffffffffffff87168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a4506001610445565b506000610445565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610df257600083815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610445565b610ea361105e565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b610f02610c81565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610ed03390565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610fc4576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044015b60405180910390fd5b5050565b6000610fea73ffffffffffffffffffffffffffffffffffffffff84168361109a565b9050805160001415801561100f57508080602001905181019061100d91906116c4565b155b156106d1576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610fbb565b60025460ff16610cbe576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606110a8838360006110af565b9392505050565b6060814710156110ed576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610fbb565b6000808573ffffffffffffffffffffffffffffffffffffffff16848660405161111691906116e6565b60006040518083038185875af1925050503d8060008114611153576040519150601f19603f3d011682016040523d82523d6000602084013e611158565b606091505b5091509150611168868383611172565b9695505050505050565b6060826111875761118282611201565b6110a8565b81511580156111ab575073ffffffffffffffffffffffffffffffffffffffff84163b155b156111fa576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610fbb565b50806110a8565b8051156112115780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561125557600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146110a857600080fd5b803573ffffffffffffffffffffffffffffffffffffffff811681146112a957600080fd5b919050565b60008083601f8401126112c057600080fd5b50813567ffffffffffffffff8111156112d857600080fd5b6020830191508360208285010111156112f057600080fd5b9250929050565b60008060008060006080868803121561130f57600080fd5b61131886611285565b945061132660208701611285565b935060408601359250606086013567ffffffffffffffff81111561134957600080fd5b611355888289016112ae565b969995985093965092949392505050565b60006020828403121561137857600080fd5b5035919050565b6000806040838503121561139257600080fd5b823591506113a260208401611285565b90509250929050565b6000602082840312156113bd57600080fd5b6110a882611285565b60008060008060008060a087890312156113df57600080fd5b6113e887611285565b95506113f660208801611285565b945060408701359350606087013567ffffffffffffffff81111561141957600080fd5b61142589828a016112ae565b909450925050608087013567ffffffffffffffff81111561144557600080fd5b87016060818a03121561145757600080fd5b809150509295509295509295565b60008060006060848603121561147a57600080fd5b61148384611285565b925061149160208501611285565b929592945050506040919091013590565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015283604082015260806060820152600061153d6080830184866114a2565b979650505050505050565b8381526040602082015260006115626040830184866114a2565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff61158982611285565b1682526000602082013567ffffffffffffffff81168082146115aa57600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126115e657600080fd5b820160208101903567ffffffffffffffff81111561160357600080fd5b80360382131561161257600080fd5b606060408601526115626060860182846114a2565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a06060820152600061167960a0830185876114a2565b828103608084015261168b818561156b565b9998505050505050505050565b8481526060602082015260006116b26060830185876114a2565b828103604084015261153d818561156b565b6000602082840312156116d657600080fd5b815180151581146110a857600080fd5b6000825160005b8181101561170757602081860181015185830152016116ed565b50600092019182525091905056fea26469706673582212207c49d30b0605b064d35bbcebc6cb6183e67c99c35afc6eb1f1e24f43be61a79b64736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a";

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
