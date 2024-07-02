/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  GatewayEVMUpgradeTest,
  GatewayEVMUpgradeTestInterface,
} from "../../../../contracts/prototypes/evm/GatewayEVMUpgradeTest";

const _abi = [
  {
    inputs: [],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [],
    name: "ApprovalFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "ExecutionFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "InsufficientETHAmount",
    type: "error",
  },
  {
    inputs: [],
    name: "SendFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "ZeroAddress",
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
        name: "destination",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "value",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
    ],
    name: "ExecutedV2",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "token",
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
        name: "amount",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
    ],
    name: "ExecutedWithERC20",
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
        indexed: false,
        internalType: "bytes",
        name: "recipient",
        type: "bytes",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "Send",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "bytes",
        name: "recipient",
        type: "bytes",
      },
      {
        indexed: true,
        internalType: "address",
        name: "asset",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "SendERC20",
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
    inputs: [],
    name: "custody",
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
        name: "destination",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
    ],
    name: "execute",
    outputs: [
      {
        internalType: "bytes",
        name: "",
        type: "bytes",
      },
    ],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "token",
        type: "address",
      },
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
      {
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
    ],
    name: "executeWithERC20",
    outputs: [
      {
        internalType: "bytes",
        name: "",
        type: "bytes",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "_tssAddress",
        type: "address",
      },
    ],
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
        internalType: "bytes",
        name: "recipient",
        type: "bytes",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "send",
    outputs: [],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes",
        name: "recipient",
        type: "bytes",
      },
      {
        internalType: "address",
        name: "token",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "sendERC20",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "_custody",
        type: "address",
      },
    ],
    name: "setCustody",
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
    inputs: [],
    name: "tssAddress",
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
] as const;

const _bytecode =
  "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b8152503480156200004757600080fd5b50620000586200005e60201b60201c565b62000208565b600060019054906101000a900460ff1615620000b1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a8906200015c565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff1614620001225760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200011991906200017e565b60405180910390a15b565b6000620001336027836200019b565b91506200014082620001b9565b604082019050919050565b6200015681620001ac565b82525050565b60006020820190508181036000830152620001778162000124565b9050919050565b60006020820190506200019560008301846200014b565b92915050565b600082825260208201905092915050565b600060ff82169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60805160601c612c4d620002436000396000818161038701528181610416015281816105100152818161059f0152610a060152612c4d6000f3fe6080604052600436106100dd5760003560e01c80638da5cb5b1161007f578063c4d66de811610059578063c4d66de814610271578063cb0271ed1461029a578063dda79b75146102c3578063f2fde38b146102ee576100dd565b80638da5cb5b146102015780639372c4ab1461022c578063ae7a3a6f14610248576100dd565b80635131ab59116100bb5780635131ab591461015757806352d1902d146101945780635b112591146101bf578063715018a6146101ea576100dd565b80631cff79cd146100e25780633659cfe6146101125780634f1ef2861461013b575b600080fd5b6100fc60048036038101906100f79190611d45565b610317565b60405161010991906123bc565b60405180910390f35b34801561011e57600080fd5b5061013960048036038101906101349190611c90565b610385565b005b61015560048036038101906101509190611da5565b61050e565b005b34801561016357600080fd5b5061017e60048036038101906101799190611cbd565b61064b565b60405161018b91906123bc565b60405180910390f35b3480156101a057600080fd5b506101a9610a02565b6040516101b6919061236f565b60405180910390f35b3480156101cb57600080fd5b506101d4610abb565b6040516101e191906122cb565b60405180910390f35b3480156101f657600080fd5b506101ff610ae1565b005b34801561020d57600080fd5b50610216610af5565b60405161022391906122cb565b60405180910390f35b61024660048036038101906102419190611ecf565b610b1f565b005b34801561025457600080fd5b5061026f600480360381019061026a9190611c90565b610c68565b005b34801561027d57600080fd5b5061029860048036038101906102939190611c90565b610cac565b005b3480156102a657600080fd5b506102c160048036038101906102bc9190611e5b565b610e9b565b005b3480156102cf57600080fd5b506102d8610f42565b6040516102e591906122cb565b60405180910390f35b3480156102fa57600080fd5b5061031560048036038101906103109190611c90565b610f68565b005b60606000610326858585610fec565b90508473ffffffffffffffffffffffffffffffffffffffff167f373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546348686604051610372939291906125bb565b60405180910390a2809150509392505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040b9061243b565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166104536110a3565b73ffffffffffffffffffffffffffffffffffffffff16146104a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a09061245b565b60405180910390fd5b6104b2816110fa565b61050b81600067ffffffffffffffff8111156104d1576104d061277c565b5b6040519080825280601f01601f1916602001820160405280156105035781602001600182028036833780820191505090505b506000611105565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16141561059d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105949061243b565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166105dc6110a3565b73ffffffffffffffffffffffffffffffffffffffff1614610632576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106299061245b565b60405180910390fd5b61063b826110fa565b61064782826001611105565b5050565b60608573ffffffffffffffffffffffffffffffffffffffff1663095ea7b38660006040518363ffffffff1660e01b815260040161068992919061231d565b602060405180830381600087803b1580156106a357600080fd5b505af11580156106b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106db9190611e01565b610711576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff1663095ea7b386866040518363ffffffff1660e01b815260040161074c929190612346565b602060405180830381600087803b15801561076657600080fd5b505af115801561077a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079e9190611e01565b6107d4576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006107e1868585610fec565b90508673ffffffffffffffffffffffffffffffffffffffff1663095ea7b38760006040518363ffffffff1660e01b815260040161081f92919061231d565b602060405180830381600087803b15801561083957600080fd5b505af115801561084d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108719190611e01565b6108a7576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008773ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016108e291906122cb565b60206040518083038186803b1580156108fa57600080fd5b505afa15801561090e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109329190611f2f565b9050600081111561098b5761098a60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16828a73ffffffffffffffffffffffffffffffffffffffff166112829092919063ffffffff16565b5b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b73828888886040516109ec939291906125bb565b60405180910390a3819250505095945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610a92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a899061249b565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b60ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610ae9611308565b610af36000611386565b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000341415610b5a576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1634604051610ba2906122b6565b60006040518083038185875af1925050503d8060008114610bdf576040519150601f19603f3d011682016040523d82523d6000602084013e610be4565b606091505b50509050600015158115151415610c27576040517f81063e5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fa93afc57f3be4641cf20c7165d11856f3b46dd376108e5fffb06f73f2b2a6d58848434604051610c5a9392919061238a565b60405180910390a150505050565b8060c960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060019054906101000a900460ff16159050808015610cdd5750600160008054906101000a900460ff1660ff16105b80610d0a5750610cec3061144c565b158015610d095750600160008054906101000a900460ff1660ff16145b5b610d49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d40906124db565b60405180910390fd5b60016000806101000a81548160ff021916908360ff1602179055508015610d86576001600060016101000a81548160ff0219169083151502179055505b610d8e61146f565b610d966114c8565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610dfd576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160ca60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015610e975760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610e8e91906123de565b60405180910390a15b5050565b610eea3360c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16838573ffffffffffffffffffffffffffffffffffffffff16611519909392919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff167f35fb30ed1b8e81eb91001dad742b13b1491a67c722e8c593a886a18500f7d9af858584604051610f349392919061238a565b60405180910390a250505050565b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610f70611308565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610fe0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fd79061241b565b60405180910390fd5b610fe981611386565b50565b60606000808573ffffffffffffffffffffffffffffffffffffffff16348686604051611019929190612286565b60006040518083038185875af1925050503d8060008114611056576040519150601f19603f3d011682016040523d82523d6000602084013e61105b565b606091505b509150915081611097576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80925050509392505050565b60006110d17f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b6115a2565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611102611308565b50565b6111317f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b6115ac565b60000160009054906101000a900460ff161561115557611150836115b6565b61127d565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561119b57600080fd5b505afa9250505080156111cc57506040513d601f19601f820116820180604052508101906111c99190611e2e565b60015b61120b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611202906124fb565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114611270576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611267906124bb565b60405180910390fd5b5061127c83838361166f565b5b505050565b6113038363a9059cbb60e01b84846040516024016112a1929190612346565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061169b565b505050565b611310611762565b73ffffffffffffffffffffffffffffffffffffffff1661132e610af5565b73ffffffffffffffffffffffffffffffffffffffff1614611384576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161137b9061253b565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff166114be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114b59061257b565b60405180910390fd5b6114c661176a565b565b600060019054906101000a900460ff16611517576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161150e9061257b565b60405180910390fd5b565b61159c846323b872dd60e01b85858560405160240161153a939291906122e6565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061169b565b50505050565b6000819050919050565b6000819050919050565b6115bf8161144c565b6115fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115f59061251b565b60405180910390fd5b8061162b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b6115a2565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611678836117cb565b6000825111806116855750805b1561169657611694838361181a565b505b505050565b60006116fd826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166118479092919063ffffffff16565b905060008151111561175d578080602001905181019061171d9190611e01565b61175c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117539061259b565b60405180910390fd5b5b505050565b600033905090565b600060019054906101000a900460ff166117b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117b09061257b565b60405180910390fd5b6117c96117c4611762565b611386565b565b6117d4816115b6565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b606061183f8383604051806060016040528060278152602001612bf16027913961185f565b905092915050565b606061185684846000856118e5565b90509392505050565b60606000808573ffffffffffffffffffffffffffffffffffffffff1685604051611889919061229f565b600060405180830381855af49150503d80600081146118c4576040519150601f19603f3d011682016040523d82523d6000602084013e6118c9565b606091505b50915091506118da868383876119b2565b925050509392505050565b60608247101561192a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119219061247b565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611953919061229f565b60006040518083038185875af1925050503d8060008114611990576040519150601f19603f3d011682016040523d82523d6000602084013e611995565b606091505b50915091506119a687838387611a28565b92505050949350505050565b60608315611a1557600083511415611a0d576119cd8561144c565b611a0c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a039061255b565b60405180910390fd5b5b829050611a20565b611a1f8383611a9e565b5b949350505050565b60608315611a8b57600083511415611a8357611a4385611aee565b611a82576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a799061255b565b60405180910390fd5b5b829050611a96565b611a958383611b11565b5b949350505050565b600082511115611ab15781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ae591906123f9565b60405180910390fd5b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082511115611b245781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b5891906123f9565b60405180910390fd5b6000611b74611b6f84612612565b6125ed565b905082815260208101848484011115611b9057611b8f6127ba565b5b611b9b848285612709565b509392505050565b600081359050611bb281612b94565b92915050565b600081519050611bc781612bab565b92915050565b600081519050611bdc81612bc2565b92915050565b60008083601f840112611bf857611bf76127b0565b5b8235905067ffffffffffffffff811115611c1557611c146127ab565b5b602083019150836001820283011115611c3157611c306127b5565b5b9250929050565b600082601f830112611c4d57611c4c6127b0565b5b8135611c5d848260208601611b61565b91505092915050565b600081359050611c7581612bd9565b92915050565b600081519050611c8a81612bd9565b92915050565b600060208284031215611ca657611ca56127c4565b5b6000611cb484828501611ba3565b91505092915050565b600080600080600060808688031215611cd957611cd86127c4565b5b6000611ce788828901611ba3565b9550506020611cf888828901611ba3565b9450506040611d0988828901611c66565b935050606086013567ffffffffffffffff811115611d2a57611d296127bf565b5b611d3688828901611be2565b92509250509295509295909350565b600080600060408486031215611d5e57611d5d6127c4565b5b6000611d6c86828701611ba3565b935050602084013567ffffffffffffffff811115611d8d57611d8c6127bf565b5b611d9986828701611be2565b92509250509250925092565b60008060408385031215611dbc57611dbb6127c4565b5b6000611dca85828601611ba3565b925050602083013567ffffffffffffffff811115611deb57611dea6127bf565b5b611df785828601611c38565b9150509250929050565b600060208284031215611e1757611e166127c4565b5b6000611e2584828501611bb8565b91505092915050565b600060208284031215611e4457611e436127c4565b5b6000611e5284828501611bcd565b91505092915050565b60008060008060608587031215611e7557611e746127c4565b5b600085013567ffffffffffffffff811115611e9357611e926127bf565b5b611e9f87828801611be2565b94509450506020611eb287828801611ba3565b9250506040611ec387828801611c66565b91505092959194509250565b600080600060408486031215611ee857611ee76127c4565b5b600084013567ffffffffffffffff811115611f0657611f056127bf565b5b611f1286828701611be2565b93509350506020611f2586828701611c66565b9150509250925092565b600060208284031215611f4557611f446127c4565b5b6000611f5384828501611c7b565b91505092915050565b611f6581612686565b82525050565b611f74816126a4565b82525050565b6000611f868385612659565b9350611f93838584612709565b611f9c836127c9565b840190509392505050565b6000611fb3838561266a565b9350611fc0838584612709565b82840190509392505050565b6000611fd782612643565b611fe18185612659565b9350611ff1818560208601612718565b611ffa816127c9565b840191505092915050565b600061201082612643565b61201a818561266a565b935061202a818560208601612718565b80840191505092915050565b61203f816126e5565b82525050565b61204e816126f7565b82525050565b600061205f8261264e565b6120698185612675565b9350612079818560208601612718565b612082816127c9565b840191505092915050565b600061209a602683612675565b91506120a5826127da565b604082019050919050565b60006120bd602c83612675565b91506120c882612829565b604082019050919050565b60006120e0602c83612675565b91506120eb82612878565b604082019050919050565b6000612103602683612675565b915061210e826128c7565b604082019050919050565b6000612126603883612675565b915061213182612916565b604082019050919050565b6000612149602983612675565b915061215482612965565b604082019050919050565b600061216c602e83612675565b9150612177826129b4565b604082019050919050565b600061218f602e83612675565b915061219a82612a03565b604082019050919050565b60006121b2602d83612675565b91506121bd82612a52565b604082019050919050565b60006121d5602083612675565b91506121e082612aa1565b602082019050919050565b60006121f860008361266a565b915061220382612aca565b600082019050919050565b600061221b601d83612675565b915061222682612acd565b602082019050919050565b600061223e602b83612675565b915061224982612af6565b604082019050919050565b6000612261602a83612675565b915061226c82612b45565b604082019050919050565b612280816126ce565b82525050565b6000612293828486611fa7565b91508190509392505050565b60006122ab8284612005565b915081905092915050565b60006122c1826121eb565b9150819050919050565b60006020820190506122e06000830184611f5c565b92915050565b60006060820190506122fb6000830186611f5c565b6123086020830185611f5c565b6123156040830184612277565b949350505050565b60006040820190506123326000830185611f5c565b61233f6020830184612036565b9392505050565b600060408201905061235b6000830185611f5c565b6123686020830184612277565b9392505050565b60006020820190506123846000830184611f6b565b92915050565b600060408201905081810360008301526123a5818587611f7a565b90506123b46020830184612277565b949350505050565b600060208201905081810360008301526123d68184611fcc565b905092915050565b60006020820190506123f36000830184612045565b92915050565b600060208201905081810360008301526124138184612054565b905092915050565b600060208201905081810360008301526124348161208d565b9050919050565b60006020820190508181036000830152612454816120b0565b9050919050565b60006020820190508181036000830152612474816120d3565b9050919050565b60006020820190508181036000830152612494816120f6565b9050919050565b600060208201905081810360008301526124b481612119565b9050919050565b600060208201905081810360008301526124d48161213c565b9050919050565b600060208201905081810360008301526124f48161215f565b9050919050565b6000602082019050818103600083015261251481612182565b9050919050565b60006020820190508181036000830152612534816121a5565b9050919050565b60006020820190508181036000830152612554816121c8565b9050919050565b600060208201905081810360008301526125748161220e565b9050919050565b6000602082019050818103600083015261259481612231565b9050919050565b600060208201905081810360008301526125b481612254565b9050919050565b60006040820190506125d06000830186612277565b81810360208301526125e3818486611f7a565b9050949350505050565b60006125f7612608565b9050612603828261274b565b919050565b6000604051905090565b600067ffffffffffffffff82111561262d5761262c61277c565b5b612636826127c9565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000612691826126ae565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006126f0826126ce565b9050919050565b6000612702826126d8565b9050919050565b82818337600083830152505050565b60005b8381101561273657808201518184015260208101905061271b565b83811115612745576000848401525b50505050565b612754826127c9565b810181811067ffffffffffffffff821117156127735761277261277c565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b612b9d81612686565b8114612ba857600080fd5b50565b612bb481612698565b8114612bbf57600080fd5b50565b612bcb816126a4565b8114612bd657600080fd5b50565b612be2816126ce565b8114612bed57600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220bb1f10e7c7799312e46755171ace34b1c83ee097036b64b3a920c525fdadf36764736f6c63430008070033";

type GatewayEVMUpgradeTestConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: GatewayEVMUpgradeTestConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class GatewayEVMUpgradeTest__factory extends ContractFactory {
  constructor(...args: GatewayEVMUpgradeTestConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<GatewayEVMUpgradeTest> {
    return super.deploy(overrides || {}) as Promise<GatewayEVMUpgradeTest>;
  }
  override getDeployTransaction(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  override attach(address: string): GatewayEVMUpgradeTest {
    return super.attach(address) as GatewayEVMUpgradeTest;
  }
  override connect(signer: Signer): GatewayEVMUpgradeTest__factory {
    return super.connect(signer) as GatewayEVMUpgradeTest__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): GatewayEVMUpgradeTestInterface {
    return new utils.Interface(_abi) as GatewayEVMUpgradeTestInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): GatewayEVMUpgradeTest {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as GatewayEVMUpgradeTest;
  }
}
