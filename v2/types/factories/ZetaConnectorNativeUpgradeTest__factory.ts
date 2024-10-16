/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Contract,
  ContractFactory,
  ContractTransactionResponse,
  Interface,
} from "ethers";
import type { Signer, ContractDeployTransaction, ContractRunner } from "ethers";
import type { NonPayableOverrides } from "../common";
import type {
  ZetaConnectorNativeUpgradeTest,
  ZetaConnectorNativeUpgradeTestInterface,
} from "../ZetaConnectorNativeUpgradeTest";

const _abi = [
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
    name: "UPGRADE_INTERFACE_VERSION",
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
    name: "initialize",
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
    outputs: [],
    stateMutability: "nonpayable",
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
    name: "proxiableUUID",
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
    name: "upgradeToAndCall",
    inputs: [
      {
        name: "newImplementation",
        type: "address",
        internalType: "address",
      },
      {
        name: "data",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [],
    stateMutability: "payable",
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
    name: "Initialized",
    inputs: [
      {
        name: "version",
        type: "uint64",
        indexed: false,
        internalType: "uint64",
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
    name: "Upgraded",
    inputs: [
      {
        name: "implementation",
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
    type: "event",
    name: "WithdrawnV2",
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
    name: "ERC1967InvalidImplementation",
    inputs: [
      {
        name: "implementation",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "ERC1967NonPayable",
    inputs: [],
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
    name: "InvalidInitialization",
    inputs: [],
  },
  {
    type: "error",
    name: "NotInitializing",
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
    name: "UUPSUnauthorizedCallContext",
    inputs: [],
  },
  {
    type: "error",
    name: "UUPSUnsupportedProxiableUUID",
    inputs: [
      {
        name: "slot",
        type: "bytes32",
        internalType: "bytes32",
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
  "0x60a060405230608052348015601357600080fd5b5060805161248461003d6000396000818161125d01528181611286015261145c01526124846000f3fe6080604052600436106101965760003560e01c80635e3e9fef116100e1578063950837aa1161008a578063ad3cb1cc11610064578063ad3cb1cc146104f2578063d547741f14610548578063e63ab1e914610568578063f8c8765e1461059c57600080fd5b8063950837aa14610489578063a217fddf146104a9578063a783c789146104be57600080fd5b80638456cb59116100bb5780638456cb59146103db57806385f438c1146103f057806391d148541461042457600080fd5b80635e3e9fef1461037b5780636f8728ad1461039b578063743e0c9b146103bb57600080fd5b806336568abe1161014357806352d1902d1161011d57806352d1902d1461030f5780635b112591146103245780635c975abb1461034457600080fd5b806336568abe146102c75780633f4ba83a146102e75780634f1ef286146102fc57600080fd5b806321e093b11161017457806321e093b11461022a578063248a9ca31461024a5780632f2ff15d146102a757600080fd5b806301ffc9a71461019b578063106e6290146101d0578063116191b6146101f2575b600080fd5b3480156101a757600080fd5b506101bb6101b6366004611dea565b6105bc565b60405190151581526020015b60405180910390f35b3480156101dc57600080fd5b506101f06101eb366004611e48565b610655565b005b3480156101fe57600080fd5b50600054610212906001600160a01b031681565b6040516001600160a01b0390911681526020016101c7565b34801561023657600080fd5b50600154610212906001600160a01b031681565b34801561025657600080fd5b50610299610265366004611e7b565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016101c7565b3480156102b357600080fd5b506101f06102c2366004611e94565b610718565b3480156102d357600080fd5b506101f06102e2366004611e94565b610762565b3480156102f357600080fd5b506101f06107ae565b6101f061030a366004611eef565b6107e3565b34801561031b57600080fd5b50610299610802565b34801561033057600080fd5b50600254610212906001600160a01b031681565b34801561035057600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166101bb565b34801561038757600080fd5b506101f061039636600461203f565b610831565b3480156103a757600080fd5b506101f06103b63660046120a1565b610985565b3480156103c757600080fd5b506101f06103d6366004611e7b565b610ade565b3480156103e757600080fd5b506101f0610afe565b3480156103fc57600080fd5b506102997f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b34801561043057600080fd5b506101bb61043f366004611e94565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561049557600080fd5b506101f06104a4366004612139565b610b30565b3480156104b557600080fd5b50610299600081565b3480156104ca57600080fd5b506102997f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b3480156104fe57600080fd5b5061053b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101c79190612178565b34801561055457600080fd5b506101f0610563366004611e94565b610cc3565b34801561057457600080fd5b506102997f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b3480156105a857600080fd5b506101f06105b73660046121c9565b610d07565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061064f57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b61065d610e8e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461068781610f0f565b61068f610f19565b6001546106a6906001600160a01b03168585610f77565b836001600160a01b03167f3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9846040516106e191815260200190565b60405180910390a25061071360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461075281610f0f565b61075c8383611011565b50505050565b6001600160a01b03811633146107a4576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61071382826110fe565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107d881610f0f565b6107e06111c2565b50565b6107eb611252565b6107f482611322565b6107fe828261132d565b5050565b600061080c611451565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b610839610e8e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461086381610f0f565b61086b610f19565b600054600154610888916001600160a01b03918216911687610f77565b6000546001546040517f5131ab590000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692635131ab59926108dd929116908a908a908a908a90600401612266565b600060405180830381600087803b1580156108f757600080fd5b505af115801561090b573d6000803e3d6000fd5b50505050856001600160a01b03167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d86868660405161094c939291906122a9565b60405180910390a25061097e60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050505050565b61098d610e8e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46109b781610f0f565b6109bf610f19565b6000546001546109dc916001600160a01b03918216911688610f77565b6000546001546040517faa0c0fc10000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263aa0c0fc192610a33929116908b908b908b908b908a90600401612374565b600060405180830381600087803b158015610a4d57600080fd5b505af1158015610a61573d6000803e3d6000fd5b50505050866001600160a01b03167f5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff087878786604051610aa494939291906123cb565b60405180910390a250610ad660017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b610ae6610f19565b6001546107e0906001600160a01b03163330846114b3565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610b2881610f0f565b6107e06114ec565b6000610b3b81610f0f565b6001600160a01b038216610b7b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254610bb2907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4906001600160a01b03166110fe565b50600254610bea907f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb906001600160a01b03166110fe565b50610c157f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483611011565b50610c407f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83611011565b50600254604080516001600160a01b03928316815291841660208301527f33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e910160405180910390a150600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610cfd81610f0f565b61075c83836110fe565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610d525750825b905060008267ffffffffffffffff166001148015610d6f5750303b155b905081158015610d7d575080155b15610db4576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610e155784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610e2189898989611565565b8315610e835784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2906020015b60405180910390a15b505050505050505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01610f09576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6107e08133611870565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615610f75576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6040516001600160a01b0383811660248301526044820183905261071391859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506118fd565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166110f4576000848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556110aa3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061064f565b600091505061064f565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16156110f4576000848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061064f565b6111ca611979565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806112eb57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166112df7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15610f75576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006107fe81610f0f565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156113a5575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526113a2918101906123f7565b60015b6113eb576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611447576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016113e2565b61071383836119d4565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610f75576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040516001600160a01b03848116602483015283811660448301526064820183905261075c9186918216906323b872dd90608401610fa4565b6114f4610f19565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611234565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156115b05750825b905060008267ffffffffffffffff1660011480156115cd5750303b155b9050811580156115db575080155b15611612576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156116735784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038916158061169057506001600160a01b038816155b806116a257506001600160a01b038716155b806116b457506001600160a01b038616155b156116eb576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6116f3611a2a565b6116fb611a32565b611703611a2a565b61170b611a42565b600080546001600160a01b03808c167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316178355600180548c831690841617905560028054918b16919092161790556117669087611011565b506117917f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e488611011565b506117bc7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb88611011565b506117e77f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87611011565b506118127f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a88611011565b508315610e835784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602001610e7a565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff166107fe576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044016113e2565b60006119126001600160a01b03841683611a52565b905080516000141580156119375750808060200190518101906119359190612410565b155b15610713576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201526024016113e2565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610f75576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119dd82611a67565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115611a22576107138282611b0f565b6107fe611b85565b610f75611bbd565b611a3a611bbd565b610f75611c24565b611a4a611bbd565b610f75611c2c565b6060611a6083836000611c7d565b9392505050565b806001600160a01b03163b600003611ab6576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016113e2565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051611b2c9190612432565b600060405180830381855af49150503d8060008114611b67576040519150601f19603f3d011682016040523d82523d6000602084013e611b6c565b606091505b5091509150611b7c858383611d33565b95945050505050565b3415610f75576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610f75576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610feb611bbd565b611c34611bbd565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b606081471015611cbb576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016113e2565b600080856001600160a01b03168486604051611cd79190612432565b60006040518083038185875af1925050503d8060008114611d14576040519150601f19603f3d011682016040523d82523d6000602084013e611d19565b606091505b5091509150611d29868383611d33565b9695505050505050565b606082611d4857611d4382611da8565b611a60565b8151158015611d5f57506001600160a01b0384163b155b15611da1576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016113e2565b5080611a60565b805115611db85780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215611dfc57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611a6057600080fd5b80356001600160a01b0381168114611e4357600080fd5b919050565b600080600060608486031215611e5d57600080fd5b611e6684611e2c565b95602085013595506040909401359392505050565b600060208284031215611e8d57600080fd5b5035919050565b60008060408385031215611ea757600080fd5b82359150611eb760208401611e2c565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060408385031215611f0257600080fd5b611f0b83611e2c565b9150602083013567ffffffffffffffff811115611f2757600080fd5b8301601f81018513611f3857600080fd5b803567ffffffffffffffff811115611f5257611f52611ec0565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715611fbe57611fbe611ec0565b604052818152828201602001871015611fd657600080fd5b816020840160208301376000602083830101528093505050509250929050565b60008083601f84011261200857600080fd5b50813567ffffffffffffffff81111561202057600080fd5b60208301915083602082850101111561203857600080fd5b9250929050565b60008060008060006080868803121561205757600080fd5b61206086611e2c565b945060208601359350604086013567ffffffffffffffff81111561208357600080fd5b61208f88828901611ff6565b96999598509660600135949350505050565b60008060008060008060a087890312156120ba57600080fd5b6120c387611e2c565b955060208701359450604087013567ffffffffffffffff8111156120e657600080fd5b6120f289828a01611ff6565b90955093505060608701359150608087013567ffffffffffffffff81111561211957600080fd5b87016080818a03121561212b57600080fd5b809150509295509295509295565b60006020828403121561214b57600080fd5b611a6082611e2c565b60005b8381101561216f578181015183820152602001612157565b50506000910152565b6020815260008251806020840152612197816040850160208701612154565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600080600080608085870312156121df57600080fd5b6121e885611e2c565b93506121f660208601611e2c565b925061220460408601611e2c565b915061221260608601611e2c565b905092959194509250565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b03861681526001600160a01b038516602082015283604082015260806060820152600061229e60808301848661221d565b979650505050505050565b838152604060208201526000611b7c60408301848661221d565b6001600160a01b036122d482611e2c565b1682526001600160a01b036122eb60208301611e2c565b1660208301526040818101359083015260006060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261233357600080fd5b820160208101903567ffffffffffffffff81111561235057600080fd5b80360382131561235f57600080fd5b60806060860152611b7c60808601828461221d565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a0606082015260006123ac60a08301858761221d565b82810360808401526123be81856122c3565b9998505050505050505050565b8481526060602082015260006123e560608301858761221d565b828103604084015261229e81856122c3565b60006020828403121561240957600080fd5b5051919050565b60006020828403121561242257600080fd5b81518015158114611a6057600080fd5b60008251612444818460208701612154565b919091019291505056fea2646970667358221220d38e7983570f30666968791527c9f048458e7c85728f9d5bced7a2ac8e7e7df364736f6c634300081a0033";

type ZetaConnectorNativeUpgradeTestConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: ZetaConnectorNativeUpgradeTestConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class ZetaConnectorNativeUpgradeTest__factory extends ContractFactory {
  constructor(...args: ZetaConnectorNativeUpgradeTestConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(overrides || {});
  }
  override deploy(overrides?: NonPayableOverrides & { from?: string }) {
    return super.deploy(overrides || {}) as Promise<
      ZetaConnectorNativeUpgradeTest & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(
    runner: ContractRunner | null
  ): ZetaConnectorNativeUpgradeTest__factory {
    return super.connect(runner) as ZetaConnectorNativeUpgradeTest__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ZetaConnectorNativeUpgradeTestInterface {
    return new Interface(_abi) as ZetaConnectorNativeUpgradeTestInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): ZetaConnectorNativeUpgradeTest {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as ZetaConnectorNativeUpgradeTest;
  }
}