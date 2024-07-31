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
  "0x60c060405234801561001057600080fd5b5060405161163738038061163783398101604081905261002f916101f5565b60016000819055805460ff19169055838383836001600160a01b038416158061005f57506001600160a01b038316155b8061007157506001600160a01b038216155b1561008f5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0384811660805283811660a052600380546001600160a01b0319169184169190911790556100c5600082610129565b506100f07f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610129565b5061011b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610129565b505050505050505050610249565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166101cf5760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101873390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101d3565b5060005b92915050565b80516001600160a01b03811681146101f057600080fd5b919050565b6000806000806080858703121561020b57600080fd5b610214856101d9565b9350610222602086016101d9565b9250610230604086016101d9565b915061023e606086016101d9565b905092959194509250565b60805160a0516113836102b4600039600081816101f70152818161046d0152818161051e01528181610638015281816107cc0152818161087d015261094b0152600081816101ab0152818161048f015281816104f1015281816107ee015261085001526113836000f3fe608060405234801561001057600080fd5b50600436106101515760003560e01c80635b112591116100cd57806391d1485411610081578063a783c78911610066578063a783c7891461031f578063d547741f14610346578063e63ab1e91461035957600080fd5b806391d14854146102d1578063a217fddf1461031757600080fd5b80635e3e9fef116100b25780635e3e9fef146102a3578063743e0c9b146102b65780638456cb59146102c957600080fd5b80635b112591146102785780635c975abb1461029857600080fd5b806321e093b1116101245780632f2ff15d116101095780632f2ff15d1461024a57806336568abe1461025d5780633f4ba83a1461027057600080fd5b806321e093b1146101f2578063248a9ca31461021957600080fd5b806301ffc9a71461015657806302d5c8991461017e578063106e629014610193578063116191b6146101a6575b600080fd5b6101696101643660046110be565b610380565b60405190151581526020015b60405180910390f35b61019161018c366004611129565b610419565b005b6101916101a13660046111bb565b6105e4565b6101cd7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610175565b6101cd7f000000000000000000000000000000000000000000000000000000000000000081565b61023c6102273660046111ee565b60009081526002602052604090206001015490565b604051908152602001610175565b610191610258366004611207565b6106bf565b61019161026b366004611207565b6106ea565b610191610743565b6003546101cd9073ffffffffffffffffffffffffffffffffffffffff1681565b60015460ff16610169565b6101916102b1366004611129565b610778565b6101916102c43660046111ee565b610929565b610191610973565b6101696102df366004611207565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61023c600081565b61023c7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b610191610354366004611207565b6109a5565b61023c7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061041357507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104216109ca565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61044b81610a0d565b610453610a17565b6104b473ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000087610a56565b6040517fb8969bd400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063b8969bd49061054e907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a9060040161127c565b600060405180830381600087803b15801561056857600080fd5b505af115801561057c573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167fba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe8686866040516105ca939291906112d9565b60405180910390a2506105dd6001600055565b5050505050565b6105ec6109ca565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61061681610a0d565b61061e610a17565b61065f73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168585610a56565b8373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364846040516106a791815260200190565b60405180910390a2506106ba6001600055565b505050565b6000828152600260205260409020600101546106da81610a0d565b6106e48383610ad7565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610739576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106ba8282610bd7565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61076d81610a0d565b610775610c96565b50565b6107806109ca565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb6107aa81610a0d565b6107b2610a17565b61081373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000087610a56565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906108ad907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a9060040161127c565b600060405180830381600087803b1580156108c757600080fd5b505af11580156108db573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced8686866040516105ca939291906112d9565b610931610a17565b61077573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084610d13565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61099d81610a0d565b610775610d59565b6000828152600260205260409020600101546109c081610a0d565b6106e48383610bd7565b600260005403610a06576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6107758133610db2565b60015460ff1615610a54576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60405173ffffffffffffffffffffffffffffffffffffffff8381166024830152604482018390526106ba91859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610e43565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610bcf57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610b6d3390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610413565b506000610413565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610bcf57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610413565b610c9e610ed9565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526106e49186918216906323b872dd90608401610a90565b610d61610a17565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610ce9565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610e3f576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044015b60405180910390fd5b5050565b6000610e6573ffffffffffffffffffffffffffffffffffffffff841683610f15565b90508051600014158015610e8a575080806020019051810190610e8891906112fc565b155b156106ba576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610e36565b60015460ff16610a54576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060610f2383836000610f2a565b9392505050565b606081471015610f68576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610e36565b6000808573ffffffffffffffffffffffffffffffffffffffff168486604051610f91919061131e565b60006040518083038185875af1925050503d8060008114610fce576040519150601f19603f3d011682016040523d82523d6000602084013e610fd3565b606091505b5091509150610fe3868383610fed565b9695505050505050565b60608261100257610ffd8261107c565b610f23565b8151158015611026575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611075576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610e36565b5080610f23565b80511561108c5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156110d057600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610f2357600080fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461112457600080fd5b919050565b60008060008060006080868803121561114157600080fd5b61114a86611100565b945060208601359350604086013567ffffffffffffffff81111561116d57600080fd5b8601601f8101881361117e57600080fd5b803567ffffffffffffffff81111561119557600080fd5b8860208284010111156111a757600080fd5b959894975060200195606001359392505050565b6000806000606084860312156111d057600080fd5b6111d984611100565b95602085013595506040909401359392505050565b60006020828403121561120057600080fd5b5035919050565b6000806040838503121561121a57600080fd5b8235915061122a60208401611100565b90509250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006112ce608083018486611233565b979650505050505050565b8381526040602082015260006112f3604083018486611233565b95945050505050565b60006020828403121561130e57600080fd5b81518015158114610f2357600080fd5b6000825160005b8181101561133f5760208186018101518583015201611325565b50600092019182525091905056fea264697066735822122049fc4f7005945c94be87b64761390ef0a7a9420599ed8603d7ad523eef4e31d764736f6c634300081a0033";

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
