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
        name: "_gateway",
        type: "address",
        internalType: "address",
      },
      {
        name: "_zetaToken",
        type: "address",
        internalType: "address",
      },
      {
        name: "_tssAddress",
        type: "address",
        internalType: "address",
      },
      {
        name: "_admin",
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
        name: "_maxSupply",
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
  "0x60c060405260001960045534801561001657600080fd5b5060405161182a38038061182a833981016040819052610035916101fb565b60016000819055805460ff19169055838383836001600160a01b038416158061006557506001600160a01b038316155b8061007757506001600160a01b038216155b156100955760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0384811660805283811660a052600380546001600160a01b0319169184169190911790556100cb60008261012f565b506100f67f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb8361012f565b506101217f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261012f565b50505050505050505061024f565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166101d55760008381526002602090815260408083206001600160a01b03861684529091529020805460ff1916600117905561018d3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101d9565b5060005b92915050565b80516001600160a01b03811681146101f657600080fd5b919050565b6000806000806080858703121561021157600080fd5b61021a856101df565b9350610228602086016101df565b9250610236604086016101df565b9150610244606086016101df565b905092959194509250565b60805160a05161155b6102cf6000396000818161021d0152818161049a015281816105dc015281816106a2015281816107a7015281816108c901528181610a7d01528181610bbf01528181610c850152610ddc0152600081816101d1015281816105a60152818161067301528181610b890152610c56015261155b6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80635c975abb116100d857806391d148541161008c578063d547741f11610066578063d547741f1461037f578063d5abeb0114610392578063e63ab1e91461039b57600080fd5b806391d148541461030a578063a217fddf14610350578063a783c7891461035857600080fd5b80636f8b44b0116100bd5780636f8b44b0146102dc578063743e0c9b146102ef5780638456cb591461030257600080fd5b80635c975abb146102be5780635e3e9fef146102c957600080fd5b8063248a9ca31161012f57806336568abe1161011457806336568abe146102835780633f4ba83a146102965780635b1125911461029e57600080fd5b8063248a9ca31461023f5780632f2ff15d1461027057600080fd5b8063106e629011610160578063106e6290146101b9578063116191b6146101cc57806321e093b11461021857600080fd5b806301ffc9a71461017c57806302d5c899146101a4575b600080fd5b61018f61018a36600461128d565b6103c2565b60405190151581526020015b60405180910390f35b6101b76101b23660046112ff565b61045b565b005b6101b76101c7366004611391565b610768565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b61026261024d3660046113c4565b60009081526002602052604090206001015490565b60405190815260200161019b565b6101b761027e3660046113dd565b610985565b6101b76102913660046113dd565b6109b0565b6101b7610a09565b6003546101f39073ffffffffffffffffffffffffffffffffffffffff1681565b60015460ff1661018f565b6101b76102d73660046112ff565b610a3e565b6101b76102ea3660046113c4565b610d31565b6101b76102fd3660046113c4565b610d9f565b6101b7610e49565b61018f6103183660046113dd565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610262600081565b6102627f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101b761038d3660046113dd565b610e7b565b61026260045481565b6102627f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061045557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b610463610ea0565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61048d81610ee3565b610495610eed565b6004547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610503573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105279190611409565b6105319087611422565b1115610569576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116600483015260248201879052604482018490527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b15801561062057600080fd5b505af1158015610634573d6000803e3d6000fd5b50506040517fb8969bd400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016925063b8969bd491506106d2907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a906004016114a5565b600060405180830381600087803b1580156106ec57600080fd5b505af1158015610700573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167fba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe86868660405161074e93929190611502565b60405180910390a2506107616001600055565b5050505050565b610770610ea0565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61079a81610ee3565b6107a2610eed565b6004547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610810573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108349190611409565b61083e9085611422565b1115610876576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015260248201859052604482018490527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b15801561090d57600080fd5b505af1158015610921573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648460405161096d91815260200190565b60405180910390a2506109806001600055565b505050565b6000828152600260205260409020600101546109a081610ee3565b6109aa8383610f2c565b50505050565b73ffffffffffffffffffffffffffffffffffffffff811633146109ff576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610980828261102c565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a3381610ee3565b610a3b6110eb565b50565b610a46610ea0565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb610a7081610ee3565b610a78610eed565b6004547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ae6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0a9190611409565b610b149087611422565b1115610b4c576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116600483015260248201879052604482018490527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b158015610c0357600080fd5b505af1158015610c17573d6000803e3d6000fd5b50506040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169250635131ab599150610cb5907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a906004016114a5565b600060405180830381600087803b158015610ccf57600080fd5b505af1158015610ce3573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced86868660405161074e93929190611502565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb610d5b81610ee3565b610d63610eed565b60048290556040518281527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c9060200160405180910390a15050565b610da7610eed565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b158015610e3557600080fd5b505af1158015610761573d6000803e3d6000fd5b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610e7381610ee3565b610a3b611168565b600082815260026020526040902060010154610e9681610ee3565b6109aa838361102c565b600260005403610edc576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b610a3b81336111c1565b60015460ff1615610f2a576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1661102457600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610fc23390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610455565b506000610455565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff161561102457600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610455565b6110f3611251565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b611170610eed565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361113e565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661124d576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440160405180910390fd5b5050565b60015460ff16610f2a576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561129f57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146112cf57600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146112fa57600080fd5b919050565b60008060008060006080868803121561131757600080fd5b611320866112d6565b945060208601359350604086013567ffffffffffffffff81111561134357600080fd5b8601601f8101881361135457600080fd5b803567ffffffffffffffff81111561136b57600080fd5b88602082840101111561137d57600080fd5b959894975060200195606001359392505050565b6000806000606084860312156113a657600080fd5b6113af846112d6565b95602085013595506040909401359392505050565b6000602082840312156113d657600080fd5b5035919050565b600080604083850312156113f057600080fd5b82359150611400602084016112d6565b90509250929050565b60006020828403121561141b57600080fd5b5051919050565b80820180821115610455577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006114f760808301848661145c565b979650505050505050565b83815260406020820152600061151c60408301848661145c565b9594505050505056fea26469706673582212209b62602791b99badf27295a4c7053b416ddaf8d056a258bb1c59e6643f5ed2d364736f6c634300081a0033";

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
    _gateway: AddressLike,
    _zetaToken: AddressLike,
    _tssAddress: AddressLike,
    _admin: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(
      _gateway,
      _zetaToken,
      _tssAddress,
      _admin,
      overrides || {}
    );
  }
  override deploy(
    _gateway: AddressLike,
    _zetaToken: AddressLike,
    _tssAddress: AddressLike,
    _admin: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(
      _gateway,
      _zetaToken,
      _tssAddress,
      _admin,
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
