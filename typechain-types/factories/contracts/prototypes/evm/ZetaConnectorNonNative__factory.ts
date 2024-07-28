/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  ZetaConnectorNonNative,
  ZetaConnectorNonNativeInterface,
} from "../../../../contracts/prototypes/evm/ZetaConnectorNonNative";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "_gateway",
        type: "address",
      },
      {
        internalType: "address",
        name: "_zetaToken",
        type: "address",
      },
      {
        internalType: "address",
        name: "_tssAddress",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [],
    name: "ExceedsMaxSupply",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidSender",
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
        internalType: "uint256",
        name: "maxSupply",
        type: "uint256",
      },
    ],
    name: "MaxSupplyUpdated",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
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
    ],
    name: "Withdraw",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
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
    name: "WithdrawAndCall",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
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
    name: "WithdrawAndRevert",
    type: "event",
  },
  {
    inputs: [],
    name: "gateway",
    outputs: [
      {
        internalType: "contract IGatewayEVM",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "maxSupply",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "receiveTokens",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "_maxSupply",
        type: "uint256",
      },
    ],
    name: "setMaxSupply",
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
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        internalType: "bytes32",
        name: "internalSendHash",
        type: "bytes32",
      },
    ],
    name: "withdraw",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
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
      {
        internalType: "bytes32",
        name: "internalSendHash",
        type: "bytes32",
      },
    ],
    name: "withdrawAndCall",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
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
      {
        internalType: "bytes32",
        name: "internalSendHash",
        type: "bytes32",
      },
    ],
    name: "withdrawAndRevert",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "zetaToken",
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
] as const;

const _bytecode =
  "0x60c060405234801561001057600080fd5b50604051610c25380380610c25833981810160405281019061003291906101ae565b81816001600081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806100a35750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156100da576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1681525050505050506101ee565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061017b82610150565b9050919050565b61018b81610170565b811461019657600080fd5b50565b6000815190506101a881610182565b92915050565b600080604083850312156101c5576101c461014b565b5b60006101d385828601610199565b92505060206101e485828601610199565b9150509250929050565b60805160a0516109e961023c6000396000818160f601528181610204015281816102300152818161031b01526103f30152600081816101e00152818161026c01526102df01526109e96000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063106e62901461005c578063116191b61461007857806321e093b1146100965780635e3e9fef146100b4578063743e0c9b146100d0575b600080fd5b610076600480360381019061007191906105ae565b6100ec565b005b6100806101de565b60405161008d9190610660565b60405180910390f35b61009e610202565b6040516100ab919061068a565b60405180910390f35b6100ce60048036038101906100c9919061070a565b610226565b005b6100ea60048036038101906100e59190610792565b6103f1565b005b6100f4610481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631e458bee8484846040518463ffffffff1660e01b8152600401610151939291906107dd565b600060405180830381600087803b15801561016b57600080fd5b505af115801561017f573d6000803e3d6000fd5b505050508273ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364836040516101c99190610814565b60405180910390a26101d96104d0565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b61022e610481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631e458bee7f000000000000000000000000000000000000000000000000000000000000000086846040518463ffffffff1660e01b81526004016102ab939291906107dd565b600060405180830381600087803b1580156102c557600080fd5b505af11580156102d9573d6000803e3d6000fd5b505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635131ab597f0000000000000000000000000000000000000000000000000000000000000000878787876040518663ffffffff1660e01b815260040161035e95949392919061088d565b600060405180830381600087803b15801561037857600080fd5b505af115801561038c573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced8585856040516103da939291906108db565b60405180910390a26103ea6104d0565b5050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379cc679033836040518363ffffffff1660e01b815260040161044c92919061090d565b600060405180830381600087803b15801561046657600080fd5b505af115801561047a573d6000803e3d6000fd5b5050505050565b6002600054036104c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104bd90610993565b60405180910390fd5b6002600081905550565b6001600081905550565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061050f826104e4565b9050919050565b61051f81610504565b811461052a57600080fd5b50565b60008135905061053c81610516565b92915050565b6000819050919050565b61055581610542565b811461056057600080fd5b50565b6000813590506105728161054c565b92915050565b6000819050919050565b61058b81610578565b811461059657600080fd5b50565b6000813590506105a881610582565b92915050565b6000806000606084860312156105c7576105c66104da565b5b60006105d58682870161052d565b93505060206105e686828701610563565b92505060406105f786828701610599565b9150509250925092565b6000819050919050565b600061062661062161061c846104e4565b610601565b6104e4565b9050919050565b60006106388261060b565b9050919050565b600061064a8261062d565b9050919050565b61065a8161063f565b82525050565b60006020820190506106756000830184610651565b92915050565b61068481610504565b82525050565b600060208201905061069f600083018461067b565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126106ca576106c96106a5565b5b8235905067ffffffffffffffff8111156106e7576106e66106aa565b5b602083019150836001820283011115610703576107026106af565b5b9250929050565b600080600080600060808688031215610726576107256104da565b5b60006107348882890161052d565b955050602061074588828901610563565b945050604086013567ffffffffffffffff811115610766576107656104df565b5b610772888289016106b4565b9350935050606061078588828901610599565b9150509295509295909350565b6000602082840312156107a8576107a76104da565b5b60006107b684828501610563565b91505092915050565b6107c881610542565b82525050565b6107d781610578565b82525050565b60006060820190506107f2600083018661067b565b6107ff60208301856107bf565b61080c60408301846107ce565b949350505050565b600060208201905061082960008301846107bf565b92915050565b600082825260208201905092915050565b82818337600083830152505050565b6000601f19601f8301169050919050565b600061086c838561082f565b9350610879838584610840565b6108828361084f565b840190509392505050565b60006080820190506108a2600083018861067b565b6108af602083018761067b565b6108bc60408301866107bf565b81810360608301526108cf818486610860565b90509695505050505050565b60006040820190506108f060008301866107bf565b8181036020830152610903818486610860565b9050949350505050565b6000604082019050610922600083018561067b565b61092f60208301846107bf565b9392505050565b600082825260208201905092915050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b600061097d601f83610936565b915061098882610947565b602082019050919050565b600060208201905081810360008301526109ac81610970565b905091905056fea2646970667358221220f157bf02f039a41c102fa373536ef4f37d367e333284c4997a08e2c01abab3f164736f6c63430008120033";


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

  override deploy(
    _gateway: PromiseOrValue<string>,
    _zetaToken: PromiseOrValue<string>,
    _tssAddress: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ZetaConnectorNonNative> {
    return super.deploy(
      _gateway,
      _zetaToken,
      _tssAddress,
      overrides || {}
    ) as Promise<ZetaConnectorNonNative>;
  }
  override getDeployTransaction(
    _gateway: PromiseOrValue<string>,
    _zetaToken: PromiseOrValue<string>,
    _tssAddress: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(
      _gateway,
      _zetaToken,
      _tssAddress,
      overrides || {}
    );
  }
  override attach(address: string): ZetaConnectorNonNative {
    return super.attach(address) as ZetaConnectorNonNative;
  }
  override connect(signer: Signer): ZetaConnectorNonNative__factory {
    return super.connect(signer) as ZetaConnectorNonNative__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ZetaConnectorNonNativeInterface {
    return new utils.Interface(_abi) as ZetaConnectorNonNativeInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ZetaConnectorNonNative {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as ZetaConnectorNonNative;
  }
}
