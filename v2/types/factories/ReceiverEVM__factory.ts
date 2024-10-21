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
    name: "onRevert",
    inputs: [
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
  "0x6080604052348015600f57600080fd5b506001600055610bac806100246000396000f3fe6080604052600436106100635760003560e01c8063c513169111610040578063c5131691146100c1578063e04d4f97146100e1578063f05b6abf146100f457005b8063357fc5a21461006c5780636ed701691461008c578063a9b0a73c146100a157005b3661006a57005b005b34801561007857600080fd5b5061006a61008736600461051e565b610114565b34801561009857600080fd5b5061006a6101aa565b3480156100ad57600080fd5b5061006a6100bc36600461055a565b6101df565b3480156100cd57600080fd5b5061006a6100dc36600461051e565b61021b565b61006a6100ef3660046106b8565b6102f6565b34801561010057600080fd5b5061006a61010f366004610797565b61033a565b61011c61036f565b61013e73ffffffffffffffffffffffffffffffffffffffff83163383866103b2565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101a56001600055565b505050565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b7fcbc82ed03b989d2f4138dbfa0d14554c8be0f0cfa5a3ebec16d6658f127ba2e53382604051610210929190610881565b60405180910390a150565b61022361036f565b600061023060028561098a565b90508060000361026c576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61028e73ffffffffffffffffffffffffffffffffffffffff84163384846103b2565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101a56001600055565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa333485858560405161032d959493929190610a29565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b1463384848460405161032d9493929190610ab3565b6002600054036103ab576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261044790859061044d565b50505050565b600080602060008451602086016000885af180610470576040513d6000823e3d81fd5b50506000513d915081156104885780600114156104a2565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15610447576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240160405180910390fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461051957600080fd5b919050565b60008060006060848603121561053357600080fd5b83359250610543602085016104f5565b9150610551604085016104f5565b90509250925092565b60006020828403121561056c57600080fd5b813567ffffffffffffffff81111561058357600080fd5b82016060818503121561059557600080fd5b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156106125761061261059c565b604052919050565b600082601f83011261062b57600080fd5b813567ffffffffffffffff8111156106455761064561059c565b61067660207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016105cb565b81815284602083860101111561068b57600080fd5b816020850160208301376000918101602001919091529392505050565b8035801515811461051957600080fd5b6000806000606084860312156106cd57600080fd5b833567ffffffffffffffff8111156106e457600080fd5b6106f08682870161061a565b93505060208401359150610551604085016106a8565b600067ffffffffffffffff8211156107205761072061059c565b5060051b60200190565b600082601f83011261073b57600080fd5b813561074e61074982610706565b6105cb565b8082825260208201915060208360051b86010192508583111561077057600080fd5b602085015b8381101561078d578035835260209283019201610775565b5095945050505050565b6000806000606084860312156107ac57600080fd5b833567ffffffffffffffff8111156107c357600080fd5b8401601f810186136107d457600080fd5b80356107e261074982610706565b8082825260208201915060208360051b85010192508883111561080457600080fd5b602084015b8381101561084657803567ffffffffffffffff81111561082857600080fd5b6108378b60208389010161061a565b84525060209283019201610809565b509550505050602084013567ffffffffffffffff81111561086657600080fd5b6108728682870161072a565b925050610551604085016106a8565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff6108bf836104f5565b166040820152600080602084013590508060608401525060408301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261090b57600080fd5b830160208101903567ffffffffffffffff81111561092857600080fd5b80360382131561093757600080fd5b606060808501528060a0850152808260c0860137600060c0828601015260c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050509392505050565b6000826109c0577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000815180845260005b818110156109eb576020818501810151868301820152016109cf565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610a5e60a08301866109c5565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610aa9578151865260209586019590910190600101610a8b565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610b46577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610b318583516109c5565b94506020938401939190910190600101610af7565b505050508281036040840152610b5c8186610a77565b915050610b6d606083018415159052565b9594505050505056fea26469706673582212200d3df21e77b9f0684cf57bfd7e586aa277028da8e173c89f9ee68e07dfcf745c64736f6c634300081a0033";

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
