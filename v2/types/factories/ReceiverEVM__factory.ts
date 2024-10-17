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
import type { ReceiverEVM, ReceiverEVMInterface } from "../ReceiverEVM";

const _abi = [
  {
    type: "fallback",
    stateMutability: "payable",
  },
  {
    type: "receive",
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "onCall",
    inputs: [
      {
        name: "",
        type: "tuple",
        internalType: "struct MessageContext",
        components: [
          {
            name: "sender",
            type: "address",
            internalType: "address",
          },
        ],
      },
      {
        name: "",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "onRevert",
    inputs: [
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
    name: "receiveERC20",
    inputs: [
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
      {
        name: "destination",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "receiveERC20Partial",
    inputs: [
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "token",
        type: "address",
        internalType: "address",
      },
      {
        name: "destination",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "receiveNoParams",
    inputs: [],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "receiveNonPayable",
    inputs: [
      {
        name: "strs",
        type: "string[]",
        internalType: "string[]",
      },
      {
        name: "nums",
        type: "uint256[]",
        internalType: "uint256[]",
      },
      {
        name: "flag",
        type: "bool",
        internalType: "bool",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "receivePayable",
    inputs: [
      {
        name: "str",
        type: "string",
        internalType: "string",
      },
      {
        name: "num",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "flag",
        type: "bool",
        internalType: "bool",
      },
    ],
    outputs: [],
    stateMutability: "payable",
  },
  {
    type: "event",
    name: "ReceivedERC20",
    inputs: [
      {
        name: "sender",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "token",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "destination",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ReceivedNoParams",
    inputs: [
      {
        name: "sender",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ReceivedNonPayable",
    inputs: [
      {
        name: "sender",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "strs",
        type: "string[]",
        indexed: false,
        internalType: "string[]",
      },
      {
        name: "nums",
        type: "uint256[]",
        indexed: false,
        internalType: "uint256[]",
      },
      {
        name: "flag",
        type: "bool",
        indexed: false,
        internalType: "bool",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ReceivedOnCall",
    inputs: [],
    anonymous: false,
  },
  {
    type: "event",
    name: "ReceivedPayable",
    inputs: [
      {
        name: "sender",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "value",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "str",
        type: "string",
        indexed: false,
        internalType: "string",
      },
      {
        name: "num",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "flag",
        type: "bool",
        indexed: false,
        internalType: "bool",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ReceivedRevert",
    inputs: [
      {
        name: "sender",
        type: "address",
        indexed: false,
        internalType: "address",
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
    name: "ZeroAmount",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x6080604052348015600f57600080fd5b506001600055610e9e806100246000396000f3fe60806040526004361061006e5760003560e01c8063c51316911161004b578063c5131691146100d5578063c9028a36146100f5578063e04d4f9714610115578063f05b6abf1461012857005b8063357fc5a214610077578063676cc054146100975780636ed70169146100c057005b3661007557005b005b34801561008357600080fd5b506100756100923660046106ec565b610148565b6100aa6100a5366004610728565b6101de565b6040516100b79190610823565b60405180910390f35b3480156100cc57600080fd5b50610075610222565b3480156100e157600080fd5b506100756100f03660046106ec565b610257565b34801561010157600080fd5b50610075610110366004610836565b610332565b610075610123366004610996565b61036e565b34801561013457600080fd5b50610075610143366004610a82565b6103b2565b6101506103e7565b61017273ffffffffffffffffffffffffffffffffffffffff831633838661042a565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101d96001600055565b505050565b6040516060907f3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee90600090a1506040805160208101909152600081525b9392505050565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61025f6103e7565b600061026c600285610b6c565b9050806000036102a8576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102ca73ffffffffffffffffffffffffffffffffffffffff841633848461042a565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101d96001600055565b7f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b3382604051610363929190610bf0565b60405180910390a150565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa33348585856040516103a5959493929190610ce2565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146338484846040516103a59493929190610d6c565b600260005403610423576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526104bf9085906104c5565b50505050565b60006104e773ffffffffffffffffffffffffffffffffffffffff841683610560565b9050805160001415801561050c57508080602001905181019061050a9190610e2f565b155b156101d9576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b606061021b83836000846000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105939190610e4c565b60006040518083038185875af1925050503d80600081146105d0576040519150601f19603f3d011682016040523d82523d6000602084013e6105d5565b606091505b50915091506105e58683836105ef565b9695505050505050565b606082610604576105ff8261067e565b61021b565b8151158015610628575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610677576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610557565b508061021b565b80511561068e5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff811681146106e757600080fd5b919050565b60008060006060848603121561070157600080fd5b83359250610711602085016106c3565b915061071f604085016106c3565b90509250925092565b6000806000838503604081121561073e57600080fd5b602081121561074c57600080fd5b50839250602084013567ffffffffffffffff81111561076a57600080fd5b8401601f8101861361077b57600080fd5b803567ffffffffffffffff81111561079257600080fd5b8660208284010111156107a457600080fd5b939660209190910195509293505050565b60005b838110156107d05781810151838201526020016107b8565b50506000910152565b600081518084526107f18160208601602086016107b5565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061021b60208301846107d9565b60006020828403121561084857600080fd5b813567ffffffffffffffff81111561085f57600080fd5b82016080818503121561021b57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156108e7576108e7610871565b604052919050565b600082601f83011261090057600080fd5b813567ffffffffffffffff81111561091a5761091a610871565b61094b60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016108a0565b81815284602083860101111561096057600080fd5b816020850160208301376000918101602001919091529392505050565b80151581146106c057600080fd5b80356106e78161097d565b6000806000606084860312156109ab57600080fd5b833567ffffffffffffffff8111156109c257600080fd5b6109ce868287016108ef565b9350506020840135915060408401356109e68161097d565b809150509250925092565b600067ffffffffffffffff821115610a0b57610a0b610871565b5060051b60200190565b600082601f830112610a2657600080fd5b8135610a39610a34826109f1565b6108a0565b8082825260208201915060208360051b860101925085831115610a5b57600080fd5b602085015b83811015610a78578035835260209283019201610a60565b5095945050505050565b600080600060608486031215610a9757600080fd5b833567ffffffffffffffff811115610aae57600080fd5b8401601f81018613610abf57600080fd5b8035610acd610a34826109f1565b8082825260208201915060208360051b850101925088831115610aef57600080fd5b602084015b83811015610b3157803567ffffffffffffffff811115610b1357600080fd5b610b228b6020838901016108ef565b84525060209283019201610af4565b509550505050602084013567ffffffffffffffff811115610b5157600080fd5b610b5d86828701610a15565b92505061071f6040850161098b565b600082610ba2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610c2e836106c3565b16604082015273ffffffffffffffffffffffffffffffffffffffff610c55602084016106c3565b166060820152600080604084013590508060808401525060608301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610ca157600080fd5b830160208101903567ffffffffffffffff811115610cbe57600080fd5b803603821315610ccd57600080fd5b608060a08501526105e560c085018284610ba7565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610d1760a08301866107d9565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610d62578151865260209586019590910190600101610d44565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610dff577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610dea8583516107d9565b94506020938401939190910190600101610db0565b505050508281036040840152610e158186610d30565b915050610e26606083018415159052565b95945050505050565b600060208284031215610e4157600080fd5b815161021b8161097d565b60008251610e5e8184602087016107b5565b919091019291505056fea26469706673582212200121b4788f0b35a4b6df21db4f43708a9b20e8b16f616e6f3300230e75bbdd1364736f6c634300081a0033";

type ReceiverEVMConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: ReceiverEVMConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class ReceiverEVM__factory extends ContractFactory {
  constructor(...args: ReceiverEVMConstructorParams) {
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
      ReceiverEVM & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): ReceiverEVM__factory {
    return super.connect(runner) as ReceiverEVM__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ReceiverEVMInterface {
    return new Interface(_abi) as ReceiverEVMInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): ReceiverEVM {
    return new Contract(address, _abi, runner) as unknown as ReceiverEVM;
  }
}
