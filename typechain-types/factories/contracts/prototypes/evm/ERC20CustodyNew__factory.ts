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
    ],
    stateMutability: "nonpayable",
    type: "constructor",
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
] as const;

const _bytecode =
  "0x60806040523480156200001157600080fd5b506040516200109038038062001090833981810160405281019062000037919062000106565b6001600081905550600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620000a7576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200018b565b600081519050620001008162000171565b92915050565b6000602082840312156200011f576200011e6200016c565b5b60006200012f84828501620000ef565b91505092915050565b600062000145826200014c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200017c8162000138565b81146200018857600080fd5b50565b610ef5806200019b6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063116191b61461004657806321fc65f214610064578063d9caed1214610080575b600080fd5b61004e61009c565b60405161005b9190610a98565b60405180910390f35b61007e600480360381019061007991906107bc565b6100c2565b005b61009a60048036038101906100959190610769565b61024a565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6100ca6102ef565b610117600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848773ffffffffffffffffffffffffffffffffffffffff1661033f9092919063ffffffff16565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635131ab5986868686866040518663ffffffff1660e01b815260040161017a959493929190610a21565b600060405180830381600087803b15801561019457600080fd5b505af11580156101a8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906101d19190610871565b508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e85858560405161023393929190610b70565b60405180910390a36102436103c5565b5050505050565b6102526102ef565b61027d82828573ffffffffffffffffffffffffffffffffffffffff1661033f9092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb836040516102da9190610b55565b60405180910390a36102ea6103c5565b505050565b60026000541415610335576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032c90610b35565b60405180910390fd5b6002600081905550565b6103c08363a9059cbb60e01b848460405160240161035e929190610a6f565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506103cf565b505050565b6001600081905550565b6000610431826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166104969092919063ffffffff16565b905060008151111561049157808060200190518101906104519190610844565b610490576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048790610b15565b60405180910390fd5b5b505050565b60606104a584846000856104ae565b90509392505050565b6060824710156104f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ea90610ad5565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161051c9190610a0a565b60006040518083038185875af1925050503d8060008114610559576040519150601f19603f3d011682016040523d82523d6000602084013e61055e565b606091505b509150915061056f8783838761057b565b92505050949350505050565b606083156105de576000835114156105d657610596856105f1565b6105d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105cc90610af5565b60405180910390fd5b5b8290506105e9565b6105e88383610614565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156106275781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065b9190610ab3565b60405180910390fd5b600061067761067284610bc7565b610ba2565b90508281526020810184848401111561069357610692610d6a565b5b61069e848285610cc8565b509392505050565b6000813590506106b581610e7a565b92915050565b6000815190506106ca81610e91565b92915050565b60008083601f8401126106e6576106e5610d60565b5b8235905067ffffffffffffffff81111561070357610702610d5b565b5b60208301915083600182028301111561071f5761071e610d65565b5b9250929050565b600082601f83011261073b5761073a610d60565b5b815161074b848260208601610664565b91505092915050565b60008135905061076381610ea8565b92915050565b60008060006060848603121561078257610781610d74565b5b6000610790868287016106a6565b93505060206107a1868287016106a6565b92505060406107b286828701610754565b9150509250925092565b6000806000806000608086880312156107d8576107d7610d74565b5b60006107e6888289016106a6565b95505060206107f7888289016106a6565b945050604061080888828901610754565b935050606086013567ffffffffffffffff81111561082957610828610d6f565b5b610835888289016106d0565b92509250509295509295909350565b60006020828403121561085a57610859610d74565b5b6000610868848285016106bb565b91505092915050565b60006020828403121561088757610886610d74565b5b600082015167ffffffffffffffff8111156108a5576108a4610d6f565b5b6108b184828501610726565b91505092915050565b6108c381610c3b565b82525050565b60006108d58385610c0e565b93506108e2838584610cb9565b6108eb83610d79565b840190509392505050565b600061090182610bf8565b61090b8185610c1f565b935061091b818560208601610cc8565b80840191505092915050565b61093081610c83565b82525050565b600061094182610c03565b61094b8185610c2a565b935061095b818560208601610cc8565b61096481610d79565b840191505092915050565b600061097c602683610c2a565b915061098782610d8a565b604082019050919050565b600061099f601d83610c2a565b91506109aa82610dd9565b602082019050919050565b60006109c2602a83610c2a565b91506109cd82610e02565b604082019050919050565b60006109e5601f83610c2a565b91506109f082610e51565b602082019050919050565b610a0481610c79565b82525050565b6000610a1682846108f6565b915081905092915050565b6000608082019050610a3660008301886108ba565b610a4360208301876108ba565b610a5060408301866109fb565b8181036060830152610a638184866108c9565b90509695505050505050565b6000604082019050610a8460008301856108ba565b610a9160208301846109fb565b9392505050565b6000602082019050610aad6000830184610927565b92915050565b60006020820190508181036000830152610acd8184610936565b905092915050565b60006020820190508181036000830152610aee8161096f565b9050919050565b60006020820190508181036000830152610b0e81610992565b9050919050565b60006020820190508181036000830152610b2e816109b5565b9050919050565b60006020820190508181036000830152610b4e816109d8565b9050919050565b6000602082019050610b6a60008301846109fb565b92915050565b6000604082019050610b8560008301866109fb565b8181036020830152610b988184866108c9565b9050949350505050565b6000610bac610bbd565b9050610bb88282610cfb565b919050565b6000604051905090565b600067ffffffffffffffff821115610be257610be1610d2c565b5b610beb82610d79565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610c4682610c59565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610c8e82610c95565b9050919050565b6000610ca082610ca7565b9050919050565b6000610cb282610c59565b9050919050565b82818337600083830152505050565b60005b83811015610ce6578082015181840152602081019050610ccb565b83811115610cf5576000848401525b50505050565b610d0482610d79565b810181811067ffffffffffffffff82111715610d2357610d22610d2c565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b610e8381610c3b565b8114610e8e57600080fd5b50565b610e9a81610c4d565b8114610ea557600080fd5b50565b610eb181610c79565b8114610ebc57600080fd5b5056fea2646970667358221220046fb4aa2d281b0818a76af1349cb058a259f6e9a48634a5fc3313b1b0fdddf764736f6c63430008070033";

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
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ERC20CustodyNew> {
    return super.deploy(_gateway, overrides || {}) as Promise<ERC20CustodyNew>;
  }
  override getDeployTransaction(
    _gateway: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(_gateway, overrides || {});
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
