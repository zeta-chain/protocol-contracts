/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  ERC20CustodyNew,
  ERC20CustodyNewInterface,
} from "../../../../contracts/prototypes/evm/ERC20CustodyNew";

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
        name: "_tssAddress",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
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
    name: "WithdrawAndCall",
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
    name: "withdrawAndCall",
    outputs: [],
    stateMutability: "nonpayable",
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
    name: "withdrawAndRevert",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

const _bytecode =
  "0x60806040523480156200001157600080fd5b506040516200130d3803806200130d833981810160405281019062000037919062000180565b6001600081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480620000a75750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15620000df576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506200021a565b6000815190506200017a8162000200565b92915050565b600080604083850312156200019a5762000199620001fb565b5b6000620001aa8582860162000169565b9250506020620001bd8582860162000169565b9150509250929050565b6000620001d482620001db565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200020b81620001c7565b81146200021757600080fd5b50565b6110e3806200022a6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063116191b61461005c57806321fc65f21461007a5780635b11259114610096578063c8a02362146100b4578063d9caed12146100d0575b600080fd5b6100646100ec565b6040516100719190610d41565b60405180910390f35b610094600480360381019061008f9190610a93565b610112565b005b61009e6102fb565b6040516100ab9190610caf565b60405180910390f35b6100ce60048036038101906100c99190610a93565b610321565b005b6100ea60048036038101906100e59190610a40565b61050a565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61011a610636565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101a1576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6101ee600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848773ffffffffffffffffffffffffffffffffffffffff166106869092919063ffffffff16565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635131ab5986868686866040518663ffffffff1660e01b8152600401610251959493929190610cca565b600060405180830381600087803b15801561026b57600080fd5b505af115801561027f573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e8585856040516102e493929190610e19565b60405180910390a36102f461070c565b5050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610329610636565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103b0576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103fd600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848773ffffffffffffffffffffffffffffffffffffffff166106869092919063ffffffff16565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b8969bd486868686866040518663ffffffff1660e01b8152600401610460959493929190610cca565b600060405180830381600087803b15801561047a57600080fd5b505af115801561048e573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c88585856040516104f393929190610e19565b60405180910390a361050361070c565b5050505050565b610512610636565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610599576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105c482828573ffffffffffffffffffffffffffffffffffffffff166106869092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb836040516106219190610dfe565b60405180910390a361063161070c565b505050565b6002600054141561067c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067390610dde565b60405180910390fd5b6002600081905550565b6107078363a9059cbb60e01b84846040516024016106a5929190610d18565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610716565b505050565b6001600081905550565b6000610778826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166107dd9092919063ffffffff16565b90506000815111156107d857808060200190518101906107989190610b1b565b6107d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ce90610dbe565b60405180910390fd5b5b505050565b60606107ec84846000856107f5565b90509392505050565b60608247101561083a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083190610d7e565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516108639190610c98565b60006040518083038185875af1925050503d80600081146108a0576040519150601f19603f3d011682016040523d82523d6000602084013e6108a5565b606091505b50915091506108b6878383876108c2565b92505050949350505050565b606083156109255760008351141561091d576108dd85610938565b61091c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091390610d9e565b60405180910390fd5b5b829050610930565b61092f838361095b565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60008251111561096e5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a29190610d5c565b60405180910390fd5b6000813590506109ba81611068565b92915050565b6000815190506109cf8161107f565b92915050565b60008083601f8401126109eb576109ea610f53565b5b8235905067ffffffffffffffff811115610a0857610a07610f4e565b5b602083019150836001820283011115610a2457610a23610f58565b5b9250929050565b600081359050610a3a81611096565b92915050565b600080600060608486031215610a5957610a58610f62565b5b6000610a67868287016109ab565b9350506020610a78868287016109ab565b9250506040610a8986828701610a2b565b9150509250925092565b600080600080600060808688031215610aaf57610aae610f62565b5b6000610abd888289016109ab565b9550506020610ace888289016109ab565b9450506040610adf88828901610a2b565b935050606086013567ffffffffffffffff811115610b0057610aff610f5d565b5b610b0c888289016109d5565b92509250509295509295909350565b600060208284031215610b3157610b30610f62565b5b6000610b3f848285016109c0565b91505092915050565b610b5181610e8e565b82525050565b6000610b638385610e61565b9350610b70838584610f0c565b610b7983610f67565b840190509392505050565b6000610b8f82610e4b565b610b998185610e72565b9350610ba9818560208601610f1b565b80840191505092915050565b610bbe81610ed6565b82525050565b6000610bcf82610e56565b610bd98185610e7d565b9350610be9818560208601610f1b565b610bf281610f67565b840191505092915050565b6000610c0a602683610e7d565b9150610c1582610f78565b604082019050919050565b6000610c2d601d83610e7d565b9150610c3882610fc7565b602082019050919050565b6000610c50602a83610e7d565b9150610c5b82610ff0565b604082019050919050565b6000610c73601f83610e7d565b9150610c7e8261103f565b602082019050919050565b610c9281610ecc565b82525050565b6000610ca48284610b84565b915081905092915050565b6000602082019050610cc46000830184610b48565b92915050565b6000608082019050610cdf6000830188610b48565b610cec6020830187610b48565b610cf96040830186610c89565b8181036060830152610d0c818486610b57565b90509695505050505050565b6000604082019050610d2d6000830185610b48565b610d3a6020830184610c89565b9392505050565b6000602082019050610d566000830184610bb5565b92915050565b60006020820190508181036000830152610d768184610bc4565b905092915050565b60006020820190508181036000830152610d9781610bfd565b9050919050565b60006020820190508181036000830152610db781610c20565b9050919050565b60006020820190508181036000830152610dd781610c43565b9050919050565b60006020820190508181036000830152610df781610c66565b9050919050565b6000602082019050610e136000830184610c89565b92915050565b6000604082019050610e2e6000830186610c89565b8181036020830152610e41818486610b57565b9050949350505050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610e9982610eac565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610ee182610ee8565b9050919050565b6000610ef382610efa565b9050919050565b6000610f0582610eac565b9050919050565b82818337600083830152505050565b60005b83811015610f39578082015181840152602081019050610f1e565b83811115610f48576000848401525b50505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b61107181610e8e565b811461107c57600080fd5b50565b61108881610ea0565b811461109357600080fd5b50565b61109f81610ecc565b81146110aa57600080fd5b5056fea26469706673582212205340b8d5ae0d440bae0160f1fe56f515f5471e0710e9b5f33c8b948d4d8769e364736f6c63430008070033";

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

  override deploy(
    _gateway: PromiseOrValue<string>,
    _tssAddress: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ERC20CustodyNew> {
    return super.deploy(
      _gateway,
      _tssAddress,
      overrides || {}
    ) as Promise<ERC20CustodyNew>;
  }
  override getDeployTransaction(
    _gateway: PromiseOrValue<string>,
    _tssAddress: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(_gateway, _tssAddress, overrides || {});
  }
  override attach(address: string): ERC20CustodyNew {
    return super.attach(address) as ERC20CustodyNew;
  }
  override connect(signer: Signer): ERC20CustodyNew__factory {
    return super.connect(signer) as ERC20CustodyNew__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ERC20CustodyNewInterface {
    return new utils.Interface(_abi) as ERC20CustodyNewInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ERC20CustodyNew {
    return new Contract(address, _abi, signerOrProvider) as ERC20CustodyNew;
  }
}
