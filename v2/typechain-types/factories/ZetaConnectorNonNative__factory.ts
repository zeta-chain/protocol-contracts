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
    name: "ExceedsMaxSupply",
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
    name: "ZeroAddress",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x60c060405260001960025534801561001657600080fd5b50604051610f9a380380610f9a833981016040819052610035916100d8565b60016000558282826001600160a01b038316158061005a57506001600160a01b038216155b8061006c57506001600160a01b038116155b1561008a5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0392831660805290821660a052600180546001600160a01b031916919092161790555061011b915050565b80516001600160a01b03811681146100d357600080fd5b919050565b6000806000606084860312156100ed57600080fd5b6100f6846100bc565b9250610104602085016100bc565b9150610112604085016100bc565b90509250925092565b60805160a051610e0061019a6000396000818161012601528181610216015281816103580152818161041e01528181610541015281816106630152818161077c015281816108be015281816109840152610af101526000818160d501528181610322015281816103ef0152818161088801526109550152610e006000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80635b112591116100765780636f8b44b01161005b5780636f8b44b01461017b578063743e0c9b1461018e578063d5abeb01146101a157600080fd5b80635b112591146101485780635e3e9fef1461016857600080fd5b806302d5c899146100a8578063106e6290146100bd578063116191b6146100d057806321e093b114610121575b600080fd5b6100bb6100b6366004610bca565b6101b8565b005b6100bb6100cb366004610c5c565b6104e3565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b6001546100f79073ffffffffffffffffffffffffffffffffffffffff1681565b6100bb610176366004610bca565b61071e565b6100bb610189366004610c8f565b610a30565b6100bb61019c366004610c8f565b610abc565b6101aa60025481565b604051908152602001610118565b6101c0610b5e565b60015473ffffffffffffffffffffffffffffffffffffffff163314610211576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561027f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102a39190610ca8565b6102ad9086610cc1565b11156102e5576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116600483015260248201869052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b15801561039c57600080fd5b505af11580156103b0573d6000803e3d6000fd5b50506040517fb8969bd400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016925063b8969bd4915061044e907f0000000000000000000000000000000000000000000000000000000000000000908990899089908990600401610d4a565b600060405180830381600087803b15801561046857600080fd5b505af115801561047c573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff167fba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe8585856040516104ca93929190610da7565b60405180910390a26104dc6001600055565b5050505050565b6104eb610b5e565b60015473ffffffffffffffffffffffffffffffffffffffff16331461053c576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ce9190610ca8565b6105d89084610cc1565b1115610610576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201849052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b1580156106a757600080fd5b505af11580156106bb573d6000803e3d6000fd5b505050508273ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648360405161070791815260200190565b60405180910390a26107196001600055565b505050565b610726610b5e565b60015473ffffffffffffffffffffffffffffffffffffffff163314610777576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108099190610ca8565b6108139086610cc1565b111561084b576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116600483015260248201869052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b15801561090257600080fd5b505af1158015610916573d6000803e3d6000fd5b50506040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169250635131ab5991506109b4907f0000000000000000000000000000000000000000000000000000000000000000908990899089908990600401610d4a565b600060405180830381600087803b1580156109ce57600080fd5b505af11580156109e2573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced8585856040516104ca93929190610da7565b60015473ffffffffffffffffffffffffffffffffffffffff163314610a81576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c9060200160405180910390a150565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b158015610b4a57600080fd5b505af11580156104dc573d6000803e3d6000fd5b600260005403610b9a576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b803573ffffffffffffffffffffffffffffffffffffffff81168114610bc557600080fd5b919050565b600080600080600060808688031215610be257600080fd5b610beb86610ba1565b945060208601359350604086013567ffffffffffffffff811115610c0e57600080fd5b8601601f81018813610c1f57600080fd5b803567ffffffffffffffff811115610c3657600080fd5b886020828401011115610c4857600080fd5b959894975060200195606001359392505050565b600080600060608486031215610c7157600080fd5b610c7a84610ba1565b95602085013595506040909401359392505050565b600060208284031215610ca157600080fd5b5035919050565b600060208284031215610cba57600080fd5b5051919050565b80820180821115610cfb577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff85166020820152836040820152608060608201526000610d9c608083018486610d01565b979650505050505050565b838152604060208201526000610dc1604083018486610d01565b9594505050505056fea2646970667358221220eef04aea35a49145971cb42500074f7c44d89db08a5df2b7d071494c05cc1a2a64736f6c634300081a0033";

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
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(
      _gateway,
      _zetaToken,
      _tssAddress,
      overrides || {}
    );
  }
  override deploy(
    _gateway: AddressLike,
    _zetaToken: AddressLike,
    _tssAddress: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(
      _gateway,
      _zetaToken,
      _tssAddress,
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
