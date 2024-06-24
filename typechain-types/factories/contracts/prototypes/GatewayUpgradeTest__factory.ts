/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../common";
import type {
  GatewayUpgradeTest,
  GatewayUpgradeTestInterface,
} from "../../../contracts/prototypes/GatewayUpgradeTest";

const _abi = [
  {
    inputs: [],
    stateMutability: "nonpayable",
    type: "constructor",
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
  "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b8152503480156200004757600080fd5b50620000586200005e60201b60201c565b62000208565b600060019054906101000a900460ff1615620000b1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a8906200015c565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff1614620001225760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200011991906200017e565b60405180910390a15b565b6000620001336027836200019b565b91506200014082620001b9565b604082019050919050565b6200015681620001ac565b82525050565b60006020820190508181036000830152620001778162000124565b9050919050565b60006020820190506200019560008301846200014b565b92915050565b600082825260208201905092915050565b600060ff82169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60805160601c6126b5620002436000396000818161038701528181610416015281816105100152818161059f015261093b01526126b56000f3fe6080604052600436106100dd5760003560e01c80638da5cb5b1161007f578063c4d66de811610059578063c4d66de814610271578063cb0271ed1461029a578063dda79b75146102c3578063f2fde38b146102ee576100dd565b80638da5cb5b146102015780639372c4ab1461022c578063ae7a3a6f14610248576100dd565b80635131ab59116100bb5780635131ab591461015757806352d1902d146101945780635b112591146101bf578063715018a6146101ea576100dd565b80631cff79cd146100e25780633659cfe6146101125780634f1ef2861461013b575b600080fd5b6100fc60048036038101906100f791906118d1565b610317565b6040516101099190611f02565b60405180910390f35b34801561011e57600080fd5b506101396004803603810190610134919061181c565b610385565b005b61015560048036038101906101509190611931565b61050e565b005b34801561016357600080fd5b5061017e60048036038101906101799190611849565b61064b565b60405161018b9190611f02565b60405180910390f35b3480156101a057600080fd5b506101a9610937565b6040516101b69190611eb5565b60405180910390f35b3480156101cb57600080fd5b506101d46109f0565b6040516101e19190611e11565b60405180910390f35b3480156101f657600080fd5b506101ff610a16565b005b34801561020d57600080fd5b50610216610a2a565b6040516102239190611e11565b60405180910390f35b61024660048036038101906102419190611a5b565b610a54565b005b34801561025457600080fd5b5061026f600480360381019061026a919061181c565b610b9c565b005b34801561027d57600080fd5b506102986004803603810190610293919061181c565b610be0565b005b3480156102a657600080fd5b506102c160048036038101906102bc91906119e7565b610d68565b005b3480156102cf57600080fd5b506102d8610e72565b6040516102e59190611e11565b60405180910390f35b3480156102fa57600080fd5b506103156004803603810190610310919061181c565b610e98565b005b60606000610326858585610f1c565b90508473ffffffffffffffffffffffffffffffffffffffff167f373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546348686604051610372939291906120c1565b60405180910390a2809150509392505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040b90611f81565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610453610fd3565b73ffffffffffffffffffffffffffffffffffffffff16146104a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a090611fa1565b60405180910390fd5b6104b28161102a565b61050b81600067ffffffffffffffff8111156104d1576104d0612282565b5b6040519080825280601f01601f1916602001820160405280156105035781602001600182028036833780820191505090505b506000611035565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16141561059d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059490611f81565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166105dc610fd3565b73ffffffffffffffffffffffffffffffffffffffff1614610632576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062990611fa1565b60405180910390fd5b61063b8261102a565b61064782826001611035565b5050565b60608573ffffffffffffffffffffffffffffffffffffffff1663095ea7b386866040518363ffffffff1660e01b8152600401610688929190611e8c565b602060405180830381600087803b1580156106a257600080fd5b505af11580156106b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106da919061198d565b5060006106e8868585610f1c565b90508673ffffffffffffffffffffffffffffffffffffffff1663095ea7b38760006040518363ffffffff1660e01b8152600401610726929190611e63565b602060405180830381600087803b15801561074057600080fd5b505af1158015610754573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610778919061198d565b5060008773ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016107b49190611e11565b60206040518083038186803b1580156107cc57600080fd5b505afa1580156107e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108049190611abb565b905060008111156108c0578773ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161086c929190611e8c565b602060405180830381600087803b15801561088657600080fd5b505af115801561089a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108be919061198d565b505b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382888888604051610921939291906120c1565b60405180910390a3819250505095945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16146109c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109be90611fc1565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b60ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610a1e6111b2565b610a286000611230565b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b80341015610a8e576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1682604051610ad690611dfc565b60006040518083038185875af1925050503d8060008114610b13576040519150601f19603f3d011682016040523d82523d6000602084013e610b18565b606091505b50509050600015158115151415610b5b576040517f81063e5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fa93afc57f3be4641cf20c7165d11856f3b46dd376108e5fffb06f73f2b2a6d58848484604051610b8e93929190611ed0565b60405180910390a150505050565b8060c960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060019054906101000a900460ff16159050808015610c115750600160008054906101000a900460ff1660ff16105b80610c3e5750610c20306112f6565b158015610c3d5750600160008054906101000a900460ff1660ff16145b5b610c7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7490612001565b60405180910390fd5b60016000806101000a81548160ff021916908360ff1602179055508015610cba576001600060016101000a81548160ff0219169083151502179055505b610cc2611319565b610cca611372565b8160ca60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015610d645760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610d5b9190611f24565b60405180910390a15b5050565b8173ffffffffffffffffffffffffffffffffffffffff166323b872dd3360c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401610dc793929190611e2c565b602060405180830381600087803b158015610de157600080fd5b505af1158015610df5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e19919061198d565b508173ffffffffffffffffffffffffffffffffffffffff167f35fb30ed1b8e81eb91001dad742b13b1491a67c722e8c593a886a18500f7d9af858584604051610e6493929190611ed0565b60405180910390a250505050565b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610ea06111b2565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610f10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f0790611f61565b60405180910390fd5b610f1981611230565b50565b60606000808573ffffffffffffffffffffffffffffffffffffffff16348686604051610f49929190611dcc565b60006040518083038185875af1925050503d8060008114610f86576040519150601f19603f3d011682016040523d82523d6000602084013e610f8b565b606091505b509150915081610fc7576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80925050509392505050565b60006110017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b6113c3565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6110326111b2565b50565b6110617f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b6113cd565b60000160009054906101000a900460ff161561108557611080836113d7565b6111ad565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156110cb57600080fd5b505afa9250505080156110fc57506040513d601f19601f820116820180604052508101906110f991906119ba565b60015b61113b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161113290612021565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b81146111a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161119790611fe1565b60405180910390fd5b506111ac838383611490565b5b505050565b6111ba6114bc565b73ffffffffffffffffffffffffffffffffffffffff166111d8610a2a565b73ffffffffffffffffffffffffffffffffffffffff161461122e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122590612061565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff16611368576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161135f906120a1565b60405180910390fd5b6113706114c4565b565b600060019054906101000a900460ff166113c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113b8906120a1565b60405180910390fd5b565b6000819050919050565b6000819050919050565b6113e0816112f6565b61141f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141690612041565b60405180910390fd5b8061144c7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b6113c3565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61149983611525565b6000825111806114a65750805b156114b7576114b58383611574565b505b505050565b600033905090565b600060019054906101000a900460ff16611513576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161150a906120a1565b60405180910390fd5b61152361151e6114bc565b611230565b565b61152e816113d7565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b60606115998383604051806060016040528060278152602001612659602791396115a1565b905092915050565b60606000808573ffffffffffffffffffffffffffffffffffffffff16856040516115cb9190611de5565b600060405180830381855af49150503d8060008114611606576040519150601f19603f3d011682016040523d82523d6000602084013e61160b565b606091505b509150915061161c86838387611627565b925050509392505050565b6060831561168a5760008351141561168257611642856112f6565b611681576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161167890612081565b60405180910390fd5b5b829050611695565b611694838361169d565b5b949350505050565b6000825111156116b05781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116e49190611f3f565b60405180910390fd5b60006117006116fb84612118565b6120f3565b90508281526020810184848401111561171c5761171b6122c0565b5b61172784828561220f565b509392505050565b60008135905061173e816125fc565b92915050565b60008151905061175381612613565b92915050565b6000815190506117688161262a565b92915050565b60008083601f840112611784576117836122b6565b5b8235905067ffffffffffffffff8111156117a1576117a06122b1565b5b6020830191508360018202830111156117bd576117bc6122bb565b5b9250929050565b600082601f8301126117d9576117d86122b6565b5b81356117e98482602086016116ed565b91505092915050565b60008135905061180181612641565b92915050565b60008151905061181681612641565b92915050565b600060208284031215611832576118316122ca565b5b60006118408482850161172f565b91505092915050565b600080600080600060808688031215611865576118646122ca565b5b60006118738882890161172f565b95505060206118848882890161172f565b9450506040611895888289016117f2565b935050606086013567ffffffffffffffff8111156118b6576118b56122c5565b5b6118c28882890161176e565b92509250509295509295909350565b6000806000604084860312156118ea576118e96122ca565b5b60006118f88682870161172f565b935050602084013567ffffffffffffffff811115611919576119186122c5565b5b6119258682870161176e565b92509250509250925092565b60008060408385031215611948576119476122ca565b5b60006119568582860161172f565b925050602083013567ffffffffffffffff811115611977576119766122c5565b5b611983858286016117c4565b9150509250929050565b6000602082840312156119a3576119a26122ca565b5b60006119b184828501611744565b91505092915050565b6000602082840312156119d0576119cf6122ca565b5b60006119de84828501611759565b91505092915050565b60008060008060608587031215611a0157611a006122ca565b5b600085013567ffffffffffffffff811115611a1f57611a1e6122c5565b5b611a2b8782880161176e565b94509450506020611a3e8782880161172f565b9250506040611a4f878288016117f2565b91505092959194509250565b600080600060408486031215611a7457611a736122ca565b5b600084013567ffffffffffffffff811115611a9257611a916122c5565b5b611a9e8682870161176e565b93509350506020611ab1868287016117f2565b9150509250925092565b600060208284031215611ad157611ad06122ca565b5b6000611adf84828501611807565b91505092915050565b611af18161218c565b82525050565b611b00816121aa565b82525050565b6000611b12838561215f565b9350611b1f83858461220f565b611b28836122cf565b840190509392505050565b6000611b3f8385612170565b9350611b4c83858461220f565b82840190509392505050565b6000611b6382612149565b611b6d818561215f565b9350611b7d81856020860161221e565b611b86816122cf565b840191505092915050565b6000611b9c82612149565b611ba68185612170565b9350611bb681856020860161221e565b80840191505092915050565b611bcb816121eb565b82525050565b611bda816121fd565b82525050565b6000611beb82612154565b611bf5818561217b565b9350611c0581856020860161221e565b611c0e816122cf565b840191505092915050565b6000611c2660268361217b565b9150611c31826122e0565b604082019050919050565b6000611c49602c8361217b565b9150611c548261232f565b604082019050919050565b6000611c6c602c8361217b565b9150611c778261237e565b604082019050919050565b6000611c8f60388361217b565b9150611c9a826123cd565b604082019050919050565b6000611cb260298361217b565b9150611cbd8261241c565b604082019050919050565b6000611cd5602e8361217b565b9150611ce08261246b565b604082019050919050565b6000611cf8602e8361217b565b9150611d03826124ba565b604082019050919050565b6000611d1b602d8361217b565b9150611d2682612509565b604082019050919050565b6000611d3e60208361217b565b9150611d4982612558565b602082019050919050565b6000611d61600083612170565b9150611d6c82612581565b600082019050919050565b6000611d84601d8361217b565b9150611d8f82612584565b602082019050919050565b6000611da7602b8361217b565b9150611db2826125ad565b604082019050919050565b611dc6816121d4565b82525050565b6000611dd9828486611b33565b91508190509392505050565b6000611df18284611b91565b915081905092915050565b6000611e0782611d54565b9150819050919050565b6000602082019050611e266000830184611ae8565b92915050565b6000606082019050611e416000830186611ae8565b611e4e6020830185611ae8565b611e5b6040830184611dbd565b949350505050565b6000604082019050611e786000830185611ae8565b611e856020830184611bc2565b9392505050565b6000604082019050611ea16000830185611ae8565b611eae6020830184611dbd565b9392505050565b6000602082019050611eca6000830184611af7565b92915050565b60006040820190508181036000830152611eeb818587611b06565b9050611efa6020830184611dbd565b949350505050565b60006020820190508181036000830152611f1c8184611b58565b905092915050565b6000602082019050611f396000830184611bd1565b92915050565b60006020820190508181036000830152611f598184611be0565b905092915050565b60006020820190508181036000830152611f7a81611c19565b9050919050565b60006020820190508181036000830152611f9a81611c3c565b9050919050565b60006020820190508181036000830152611fba81611c5f565b9050919050565b60006020820190508181036000830152611fda81611c82565b9050919050565b60006020820190508181036000830152611ffa81611ca5565b9050919050565b6000602082019050818103600083015261201a81611cc8565b9050919050565b6000602082019050818103600083015261203a81611ceb565b9050919050565b6000602082019050818103600083015261205a81611d0e565b9050919050565b6000602082019050818103600083015261207a81611d31565b9050919050565b6000602082019050818103600083015261209a81611d77565b9050919050565b600060208201905081810360008301526120ba81611d9a565b9050919050565b60006040820190506120d66000830186611dbd565b81810360208301526120e9818486611b06565b9050949350505050565b60006120fd61210e565b90506121098282612251565b919050565b6000604051905090565b600067ffffffffffffffff82111561213357612132612282565b5b61213c826122cf565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000612197826121b4565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006121f6826121d4565b9050919050565b6000612208826121de565b9050919050565b82818337600083830152505050565b60005b8381101561223c578082015181840152602081019050612221565b8381111561224b576000848401525b50505050565b61225a826122cf565b810181811067ffffffffffffffff8211171561227957612278612282565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b6126058161218c565b811461261057600080fd5b50565b61261c8161219e565b811461262757600080fd5b50565b612633816121aa565b811461263e57600080fd5b50565b61264a816121d4565b811461265557600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122016c8c753bbb77e16ebf9489b2cd4a97c5b55d04d368fbf583d2a13b6f1def8bf64736f6c63430008070033";

type GatewayUpgradeTestConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: GatewayUpgradeTestConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class GatewayUpgradeTest__factory extends ContractFactory {
  constructor(...args: GatewayUpgradeTestConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<GatewayUpgradeTest> {
    return super.deploy(overrides || {}) as Promise<GatewayUpgradeTest>;
  }
  override getDeployTransaction(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  override attach(address: string): GatewayUpgradeTest {
    return super.attach(address) as GatewayUpgradeTest;
  }
  override connect(signer: Signer): GatewayUpgradeTest__factory {
    return super.connect(signer) as GatewayUpgradeTest__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): GatewayUpgradeTestInterface {
    return new utils.Interface(_abi) as GatewayUpgradeTestInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): GatewayUpgradeTest {
    return new Contract(address, _abi, signerOrProvider) as GatewayUpgradeTest;
  }
}
