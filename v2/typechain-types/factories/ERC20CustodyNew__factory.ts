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
  ERC20CustodyNew,
  ERC20CustodyNewInterface,
} from "../ERC20CustodyNew";

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
        name: "_tssAddress",
        type: "address",
        internalType: "address",
      },
    ],
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
    name: "withdraw",
    inputs: [
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
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
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "withdrawAndCall",
    inputs: [
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
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
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "withdrawAndRevert",
    inputs: [
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
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
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    name: "Withdraw",
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
    name: "WithdrawAndCall",
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
    name: "WithdrawAndRevert",
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
    name: "FailedInnerCall",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidSender",
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
  "0x608060405234801561001057600080fd5b50604051610b4a380380610b4a83398101604081905261002f916100bc565b60016000556001600160a01b038216158061005157506001600160a01b038116155b1561006f5760405163d92e233d60e01b815260040160405180910390fd5b600180546001600160a01b039384166001600160a01b031991821617909155600280549290931691161790556100ef565b80516001600160a01b03811681146100b757600080fd5b919050565b600080604083850312156100cf57600080fd5b6100d8836100a0565b91506100e6602084016100a0565b90509250929050565b610a4c806100fe6000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80635b112591116100505780635b112591146100ca578063c8a02362146100ea578063d9caed12146100fd57600080fd5b8063116191b61461006c57806321fc65f2146100b5575b600080fd5b60015461008c9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100c86100c3366004610822565b610110565b005b60025461008c9073ffffffffffffffffffffffffffffffffffffffff1681565b6100c86100f8366004610822565b61029a565b6100c861010b3660046108bf565b61040b565b6101186104fb565b60025473ffffffffffffffffffffffffffffffffffffffff163314610169576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546101909073ffffffffffffffffffffffffffffffffffffffff87811691168561053e565b6001546040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690635131ab59906101ee9088908890889088908890600401610945565b600060405180830381600087803b15801561020857600080fd5b505af115801561021c573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e858585604051610281939291906109a2565b60405180910390a36102936001600055565b5050505050565b6102a26104fb565b60025473ffffffffffffffffffffffffffffffffffffffff1633146102f3576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015461031a9073ffffffffffffffffffffffffffffffffffffffff87811691168561053e565b6001546040517fb8969bd400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063b8969bd4906103789088908890889088908890600401610945565b600060405180830381600087803b15801561039257600080fd5b505af11580156103a6573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c8858585604051610281939291906109a2565b6104136104fb565b60025473ffffffffffffffffffffffffffffffffffffffff163314610464576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61048573ffffffffffffffffffffffffffffffffffffffff8416838361053e565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb836040516104e491815260200190565b60405180910390a36104f66001600055565b505050565b600260005403610537576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff848116602483015260448083018590528351808403909101815260649092019092526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526104f6918591906000906105d790841683610650565b905080516000141580156105fc5750808060200190518101906105fa91906109c5565b155b156104f6576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b606061065e83836000610665565b9392505050565b6060814710156106a3576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610647565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516106cc91906109e7565b60006040518083038185875af1925050503d8060008114610709576040519150601f19603f3d011682016040523d82523d6000602084013e61070e565b606091505b509150915061071e868383610728565b9695505050505050565b60608261073d57610738826107b7565b61065e565b8151158015610761575073ffffffffffffffffffffffffffffffffffffffff84163b155b156107b0576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610647565b508061065e565b8051156107c75780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461081d57600080fd5b919050565b60008060008060006080868803121561083a57600080fd5b610843866107f9565b9450610851602087016107f9565b935060408601359250606086013567ffffffffffffffff81111561087457600080fd5b8601601f8101881361088557600080fd5b803567ffffffffffffffff81111561089c57600080fd5b8860208284010111156108ae57600080fd5b959894975092955050506020019190565b6000806000606084860312156108d457600080fd5b6108dd846107f9565b92506108eb602085016107f9565b929592945050506040919091013590565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006109976080830184866108fc565b979650505050505050565b8381526040602082015260006109bc6040830184866108fc565b95945050505050565b6000602082840312156109d757600080fd5b8151801515811461065e57600080fd5b6000825160005b81811015610a0857602081860181015185830152016109ee565b50600092019182525091905056fea264697066735822122041336b1568f49be1aa3dcd208e804b373772741e9b7e1726a8869e99fced60c764736f6c634300081a0033";

type ERC20CustodyNewConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: ERC20CustodyNewConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class ERC20CustodyNew__factory extends ContractFactory {
  constructor(...args: ERC20CustodyNewConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    _gateway: AddressLike,
    _tssAddress: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(_gateway, _tssAddress, overrides || {});
  }
  override deploy(
    _gateway: AddressLike,
    _tssAddress: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(_gateway, _tssAddress, overrides || {}) as Promise<
      ERC20CustodyNew & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): ERC20CustodyNew__factory {
    return super.connect(runner) as ERC20CustodyNew__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ERC20CustodyNewInterface {
    return new Interface(_abi) as ERC20CustodyNewInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): ERC20CustodyNew {
    return new Contract(address, _abi, runner) as unknown as ERC20CustodyNew;
  }
}
