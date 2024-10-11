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
import type {
  ZetaConnectorNative,
  ZetaConnectorNativeInterface,
} from "../ZetaConnectorNative";

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
        name: "zetaToken_",
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
    name: "TSS_ROLE",
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
    name: "receiveTokens",
    inputs: [
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
    name: "withdraw",
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
      {
        name: "internalSendHash",
        type: "bytes32",
        internalType: "bytes32",
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
        name: "internalSendHash",
        type: "bytes32",
        internalType: "bytes32",
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
        name: "internalSendHash",
        type: "bytes32",
        internalType: "bytes32",
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
    type: "function",
    name: "zetaToken",
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
    name: "UpdatedZetaConnectorTSSAddress",
    inputs: [
      {
        name: "oldTSSAddress",
        type: "address",
        indexed: false,
        internalType: "address",
      },
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
    name: "Withdrawn",
    inputs: [
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
  "0x60c060405234801561001057600080fd5b50604051611afb380380611afb83398101604081905261002f91610232565b60016000819055805460ff19169055838383836001600160a01b038416158061005f57506001600160a01b038316155b8061007157506001600160a01b038216155b8061008357506001600160a01b038116155b156100a15760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0384811660805283811660a052600380546001600160a01b0319169184169190911790556100d7600082610166565b506101027f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610166565b5061012d7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610166565b506101587f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610166565b505050505050505050610286565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff1661020c5760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101c43390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610210565b5060005b92915050565b80516001600160a01b038116811461022d57600080fd5b919050565b6000806000806080858703121561024857600080fd5b61025185610216565b935061025f60208601610216565b925061026d60408601610216565b915061027b60608601610216565b905092959194509250565b60805160a05161180a6102f16000396000818161020a015281816104cd01528181610661015281816107120152818161082c015281816108dd01526109ca0152600081816101be01528181610683015281816106e50152818161084e01526108b0015261180a6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80635e3e9fef116100d857806391d148541161008c578063a783c78911610066578063a783c7891461037f578063d547741f146103a6578063e63ab1e9146103b957600080fd5b806391d148541461031e578063950837aa14610364578063a217fddf1461037757600080fd5b8063743e0c9b116100bd578063743e0c9b146102dc5780638456cb59146102ef57806385f438c1146102f757600080fd5b80635e3e9fef146102b65780636f8728ad146102c957600080fd5b80632f2ff15d1161012f5780633f4ba83a116101145780633f4ba83a146102835780635b1125911461028b5780635c975abb146102ab57600080fd5b80632f2ff15d1461025d57806336568abe1461027057600080fd5b8063116191b611610160578063116191b6146101b957806321e093b114610205578063248a9ca31461022c57600080fd5b806301ffc9a71461017c578063106e6290146101a4575b600080fd5b61018f61018a366004611311565b6103e0565b60405190151581526020015b60405180910390f35b6101b76101b236600461137c565b610479565b005b6101e07f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101e07f000000000000000000000000000000000000000000000000000000000000000081565b61024f61023a3660046113af565b60009081526002602052604090206001015490565b60405190815260200161019b565b6101b761026b3660046113c8565b610554565b6101b761027e3660046113c8565b61057f565b6101b76105d8565b6003546101e09073ffffffffffffffffffffffffffffffffffffffff1681565b60015460ff1661018f565b6101b76102c436600461143d565b61060d565b6101b76102d736600461149f565b6107d8565b6101b76102ea3660046113af565b6109a8565b6101b76109f2565b61024f7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b61018f61032c3660046113c8565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6101b7610372366004611537565b610a24565b61024f600081565b61024f7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101b76103b43660046113c8565b610bf8565b61024f7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061047357507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b610481610c1d565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46104ab81610c60565b6104b3610c6a565b6104f473ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168585610ca9565b8373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d58460405161053c91815260200190565b60405180910390a25061054f6001600055565b505050565b60008281526002602052604090206001015461056f81610c60565b6105798383610d2a565b50505050565b73ffffffffffffffffffffffffffffffffffffffff811633146105ce576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61054f8282610e2a565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61060281610c60565b61060a610ee9565b50565b610615610c1d565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461063f81610c60565b610647610c6a565b6106a873ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000087610ca9565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab5990610742907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a9060040161159b565b600060405180830381600087803b15801561075c57600080fd5b505af1158015610770573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d8686866040516107be939291906115f8565b60405180910390a2506107d16001600055565b5050505050565b6107e0610c1d565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461080a81610c60565b610812610c6a565b61087373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000088610ca9565b6040517faa0c0fc100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063aa0c0fc19061090f907f0000000000000000000000000000000000000000000000000000000000000000908b908b908b908b908a906004016116e6565b600060405180830381600087803b15801561092957600080fd5b505af115801561093d573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff167f5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff08787878660405161098d9493929190611757565b60405180910390a2506109a06001600055565b505050505050565b6109b0610c6a565b61060a73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084610f66565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a1c81610c60565b61060a610fac565b6000610a2f81610c60565b73ffffffffffffffffffffffffffffffffffffffff8216610a7c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354610ac0907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e49073ffffffffffffffffffffffffffffffffffffffff16610e2a565b50600354610b05907f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb9073ffffffffffffffffffffffffffffffffffffffff16610e2a565b50610b307f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610d2a565b50610b5b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610d2a565b506003546040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301527f33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e910160405180910390a150600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600082815260026020526040902060010154610c1381610c60565b6105798383610e2a565b600260005403610c59576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61060a8133611005565b60015460ff1615610ca7576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60405173ffffffffffffffffffffffffffffffffffffffff83811660248301526044820183905261054f91859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611096565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610e2257600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610dc03390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610473565b506000610473565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610e2257600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610473565b610ef161112c565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526105799186918216906323b872dd90608401610ce3565b610fb4610c6a565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610f3c565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16611092576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044015b60405180910390fd5b5050565b60006110b873ffffffffffffffffffffffffffffffffffffffff841683611168565b905080516000141580156110dd5750808060200190518101906110db9190611783565b155b1561054f576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401611089565b60015460ff16610ca7576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606111768383600061117d565b9392505050565b6060814710156111bb576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611089565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516111e491906117a5565b60006040518083038185875af1925050503d8060008114611221576040519150601f19603f3d011682016040523d82523d6000602084013e611226565b606091505b5091509150611236868383611240565b9695505050505050565b60608261125557611250826112cf565b611176565b8151158015611279575073ffffffffffffffffffffffffffffffffffffffff84163b155b156112c8576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401611089565b5080611176565b8051156112df5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561132357600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461117657600080fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461137757600080fd5b919050565b60008060006060848603121561139157600080fd5b61139a84611353565b95602085013595506040909401359392505050565b6000602082840312156113c157600080fd5b5035919050565b600080604083850312156113db57600080fd5b823591506113eb60208401611353565b90509250929050565b60008083601f84011261140657600080fd5b50813567ffffffffffffffff81111561141e57600080fd5b60208301915083602082850101111561143657600080fd5b9250929050565b60008060008060006080868803121561145557600080fd5b61145e86611353565b945060208601359350604086013567ffffffffffffffff81111561148157600080fd5b61148d888289016113f4565b96999598509660600135949350505050565b60008060008060008060a087890312156114b857600080fd5b6114c187611353565b955060208701359450604087013567ffffffffffffffff8111156114e457600080fd5b6114f089828a016113f4565b90955093505060608701359150608087013567ffffffffffffffff81111561151757600080fd5b87016080818a03121561152957600080fd5b809150509295509295509295565b60006020828403121561154957600080fd5b61117682611353565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006115ed608083018486611552565b979650505050505050565b838152604060208201526000611612604083018486611552565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff61163982611353565b16825273ffffffffffffffffffffffffffffffffffffffff61165d60208301611353565b1660208301526040818101359083015260006060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126116a557600080fd5b820160208101903567ffffffffffffffff8111156116c257600080fd5b8036038213156116d157600080fd5b60806060860152611612608086018284611552565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a06060820152600061173860a083018587611552565b828103608084015261174a818561161b565b9998505050505050505050565b848152606060208201526000611771606083018587611552565b82810360408401526115ed818561161b565b60006020828403121561179557600080fd5b8151801515811461117657600080fd5b6000825160005b818110156117c657602081860181015185830152016117ac565b50600092019182525091905056fea26469706673582212202540296267e9877725c1683053202b4d3b35eb6bc8480c57ffa0432a0801a67c64736f6c634300081a0033";

type ZetaConnectorNativeConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: ZetaConnectorNativeConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class ZetaConnectorNative__factory extends ContractFactory {
  constructor(...args: ZetaConnectorNativeConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    gateway_: AddressLike,
    zetaToken_: AddressLike,
    tssAddress_: AddressLike,
    admin_: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(
      gateway_,
      zetaToken_,
      tssAddress_,
      admin_,
      overrides || {}
    );
  }
  override deploy(
    gateway_: AddressLike,
    zetaToken_: AddressLike,
    tssAddress_: AddressLike,
    admin_: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(
      gateway_,
      zetaToken_,
      tssAddress_,
      admin_,
      overrides || {}
    ) as Promise<
      ZetaConnectorNative & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(
    runner: ContractRunner | null
  ): ZetaConnectorNative__factory {
    return super.connect(runner) as ZetaConnectorNative__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ZetaConnectorNativeInterface {
    return new Interface(_abi) as ZetaConnectorNativeInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): ZetaConnectorNative {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as ZetaConnectorNative;
  }
}
