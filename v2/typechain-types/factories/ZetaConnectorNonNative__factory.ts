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
  ZetaConnectorNonNative,
  ZetaConnectorNonNativeInterface,
} from "../ZetaConnectorNonNative";

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
    name: "maxSupply",
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
    name: "setMaxSupply",
    inputs: [
      {
        name: "maxSupply_",
        type: "uint256",
        internalType: "uint256",
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
    name: "MaxSupplyUpdated",
    inputs: [
      {
        name: "maxSupply",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
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
    name: "Withdraw",
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
    name: "WithdrawAndCall",
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
    name: "WithdrawAndRevert",
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
    name: "EnforcedPause",
    inputs: [],
  },
  {
    type: "error",
    name: "ExceedsMaxSupply",
    inputs: [],
  },
  {
    type: "error",
    name: "ExpectedPause",
    inputs: [],
  },
  {
    type: "error",
    name: "ReentrancyGuardReentrantCall",
    inputs: [],
  },
  {
    type: "error",
    name: "ZeroAddress",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x60c060405260001960035534801561001657600080fd5b5060405161155a38038061155a83398101604081905261003591610220565b60016000819055805460ff19169055838383836001600160a01b038416158061006557506001600160a01b038316155b8061007757506001600160a01b038216155b8061008957506001600160a01b038116155b156100a75760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03808516608052831660a0526100c5600082610154565b506100f07f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610154565b5061011b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610154565b506101467f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610154565b505050505050505050610274565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166101fa5760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101b23390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101fe565b5060005b92915050565b80516001600160a01b038116811461021b57600080fd5b919050565b6000806000806080858703121561023657600080fd5b61023f85610204565b935061024d60208601610204565b925061025b60408601610204565b915061026960608601610204565b905092959194509250565b60805160a0516112826102d86000396000818161021d01528181610531015281816108240152818161097b01528181610ad00152610bf20152600081816101d1015281816104a1015281816105040152818161079401526107f701526112826000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80635e3e9fef116100d857806391d148541161008c578063d547741f11610066578063d547741f14610386578063d5abeb0114610399578063e63ab1e9146103a257600080fd5b806391d1485414610311578063a217fddf14610357578063a783c7891461035f57600080fd5b8063743e0c9b116100bd578063743e0c9b146102cf5780638456cb59146102e257806385f438c1146102ea57600080fd5b80635e3e9fef146102a95780636f8b44b0146102bc57600080fd5b8063248a9ca31161012f57806336568abe1161011457806336568abe146102835780633f4ba83a146102965780635c975abb1461029e57600080fd5b8063248a9ca31461023f5780632f2ff15d1461027057600080fd5b8063106e629011610160578063106e6290146101b9578063116191b6146101cc57806321e093b11461021857600080fd5b806301ffc9a71461017c57806302d5c899146101a4575b600080fd5b61018f61018a366004610fb4565b6103c9565b60405190151581526020015b60405180910390f35b6101b76101b2366004611026565b610462565b005b6101b76101c73660046110b8565b6105f7565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b61026261024d3660046110eb565b60009081526002602052604090206001015490565b60405190815260200161019b565b6101b761027e366004611104565b61069c565b6101b7610291366004611104565b6106c7565b6101b7610720565b60015460ff1661018f565b6101b76102b7366004611026565b610755565b6101b76102ca3660046110eb565b6108d0565b6101b76102dd3660046110eb565b61093e565b6101b76109e8565b6102627f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b61018f61031f366004611104565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610262600081565b6102627f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101b7610394366004611104565b610a1a565b61026260035481565b6102627f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061045c57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b61046a610a3f565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461049481610a82565b61049c610a8c565b6104c77f00000000000000000000000000000000000000000000000000000000000000008684610acb565b6040517fb8969bd400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063b8969bd490610561907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a90600401611179565b600060405180830381600087803b15801561057b57600080fd5b505af115801561058f573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167fba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe8686866040516105dd939291906111d6565b60405180910390a2506105f06001600055565b5050505050565b6105ff610a3f565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461062981610a82565b610631610a8c565b61063c848484610acb565b8373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648460405161068491815260200190565b60405180910390a2506106976001600055565b505050565b6000828152600260205260409020600101546106b781610a82565b6106c18383610c53565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610716576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106978282610d53565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61074a81610a82565b610752610e12565b50565b61075d610a3f565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461078781610a82565b61078f610a8c565b6107ba7f00000000000000000000000000000000000000000000000000000000000000008684610acb565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab5990610854907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a90600401611179565b600060405180830381600087803b15801561086e57600080fd5b505af1158015610882573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced8686866040516105dd939291906111d6565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb6108fa81610a82565b610902610a8c565b60038290556040518281527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c9060200160405180910390a15050565b610946610a8c565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b1580156109d457600080fd5b505af11580156105f0573d6000803e3d6000fd5b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a1281610a82565b610752610e8f565b600082815260026020526040902060010154610a3581610a82565b6106c18383610d53565b600260005403610a7b576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6107528133610ee8565b60015460ff1615610ac9576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6003547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5d91906111f9565b610b679084611212565b1115610b9f576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201849052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b158015610c3657600080fd5b505af1158015610c4a573d6000803e3d6000fd5b50505050505050565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610d4b57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610ce93390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161045c565b50600061045c565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610d4b57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161045c565b610e1a610f78565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b610e97610a8c565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610e65565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610f74576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440160405180910390fd5b5050565b60015460ff16610ac9576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215610fc657600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610ff657600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461102157600080fd5b919050565b60008060008060006080868803121561103e57600080fd5b61104786610ffd565b945060208601359350604086013567ffffffffffffffff81111561106a57600080fd5b8601601f8101881361107b57600080fd5b803567ffffffffffffffff81111561109257600080fd5b8860208284010111156110a457600080fd5b959894975060200195606001359392505050565b6000806000606084860312156110cd57600080fd5b6110d684610ffd565b95602085013595506040909401359392505050565b6000602082840312156110fd57600080fd5b5035919050565b6000806040838503121561111757600080fd5b8235915061112760208401610ffd565b90509250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006111cb608083018486611130565b979650505050505050565b8381526040602082015260006111f0604083018486611130565b95945050505050565b60006020828403121561120b57600080fd5b5051919050565b8082018082111561045c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220de5fa55e0ca79ea2f08903481d01ea00bbbf346b9cfba44e12a2dd6f3a6f2e8a64736f6c634300081a0033";

type ZetaConnectorNonNativeConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: ZetaConnectorNonNativeConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class ZetaConnectorNonNative__factory extends ContractFactory {
  constructor(...args: ZetaConnectorNonNativeConstructorParams) {
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
      ZetaConnectorNonNative & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(
    runner: ContractRunner | null
  ): ZetaConnectorNonNative__factory {
    return super.connect(runner) as ZetaConnectorNonNative__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ZetaConnectorNonNativeInterface {
    return new Interface(_abi) as ZetaConnectorNonNativeInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): ZetaConnectorNonNative {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as ZetaConnectorNonNative;
  }
}
