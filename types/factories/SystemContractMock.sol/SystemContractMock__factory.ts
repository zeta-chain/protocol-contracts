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
import type { NonPayableOverrides } from "../../common";
import type {
  SystemContractMock,
  SystemContractMockInterface,
} from "../../SystemContractMock.sol/SystemContractMock";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "wzeta_",
        type: "address",
        internalType: "address",
      },
      {
        name: "uniswapv2Factory_",
        type: "address",
        internalType: "address",
      },
      {
        name: "uniswapv2Router02_",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "gasCoinZRC20ByChainId",
    inputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
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
    name: "gasPriceByChainId",
    inputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
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
    name: "gasZetaPoolByChainId",
    inputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
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
    name: "onCrossChainCall",
    inputs: [
      {
        name: "target",
        type: "address",
        internalType: "address",
      },
      {
        name: "zrc20",
        type: "address",
        internalType: "address",
      },
      {
        name: "amount",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "setGasCoinZRC20",
    inputs: [
      {
        name: "chainID",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "zrc20",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "setGasPrice",
    inputs: [
      {
        name: "chainID",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "price",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "setWZETAContractAddress",
    inputs: [
      {
        name: "addr",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "uniswapv2FactoryAddress",
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
    name: "uniswapv2PairFor",
    inputs: [
      {
        name: "factory",
        type: "address",
        internalType: "address",
      },
      {
        name: "tokenA",
        type: "address",
        internalType: "address",
      },
      {
        name: "tokenB",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [
      {
        name: "pair",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "uniswapv2Router02Address",
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
    name: "wZetaContractAddress",
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
    name: "SetGasCoin",
    inputs: [
      {
        name: "",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "SetGasPrice",
    inputs: [
      {
        name: "",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "SetGasZetaPool",
    inputs: [
      {
        name: "",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "SetWZeta",
    inputs: [
      {
        name: "",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "SystemContractDeployed",
    inputs: [],
    anonymous: false,
  },
  {
    type: "error",
    name: "CallerIsNotFungibleModule",
    inputs: [],
  },
  {
    type: "error",
    name: "CantBeIdenticalAddresses",
    inputs: [],
  },
  {
    type: "error",
    name: "CantBeZeroAddress",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidTarget",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x608060405234801561001057600080fd5b50604051610b3f380380610b3f83398101604081905261002f916100b9565b600380546001600160a01b038086166001600160a01b0319928316179092556004805485841690831617905560058054928416929091169190911790556040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac590600090a15050506100fc565b80516001600160a01b03811681146100b457600080fd5b919050565b6000806000606084860312156100ce57600080fd5b6100d78461009d565b92506100e56020850161009d565b91506100f36040850161009d565b90509250925092565b610a348061010b6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806397770dff11610081578063d7fd7afb1161005b578063d7fd7afb146101f2578063d936a01214610220578063ee2815ba1461024057600080fd5b806397770dff146101b9578063a7cb0507146101cc578063c63585cc146101df57600080fd5b8063513a9c05116100b2578063513a9c0514610143578063569541b914610179578063842da36d1461019957600080fd5b80630be15547146100ce5780633c669d551461012e575b600080fd5b6101046100dc36600461071e565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61014161013c366004610760565b610253565b005b61010461015136600461071e565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101049073ffffffffffffffffffffffffffffffffffffffff1681565b6005546101049073ffffffffffffffffffffffffffffffffffffffff1681565b6101416101c73660046107fd565b6103a0565b6101416101da36600461081f565b610419565b6101046101ed366004610841565b610467565b61021261020036600461071e565b60006020819052908152604090205481565b604051908152602001610125565b6004546101049073ffffffffffffffffffffffffffffffffffffffff1681565b61014161024e366004610884565b61059c565b604080516080810182526000606082019081528152336020820152468183015290517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820186905286169063a9059cbb906044016020604051808303816000875af11580156102e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030b91906108b0565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87169063de43156e90610366908490899089908990899060040161091b565b600060405180830381600087803b15801561038057600080fd5b505af1158015610394573d6000803e3d6000fd5b50505050505050505050565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e9060200160405180910390a150565b6000828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b60008060006104768585610620565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b166034820152919350915086906048016040516020818303038152906040528051906020012060405160200161055c9291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b60008281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d910161045b565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610688576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16106106c25782846106c5565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610717576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b60006020828403121561073057600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461075b57600080fd5b919050565b60008060008060006080868803121561077857600080fd5b61078186610737565b945061078f60208701610737565b935060408601359250606086013567ffffffffffffffff8111156107b257600080fd5b8601601f810188136107c357600080fd5b803567ffffffffffffffff8111156107da57600080fd5b8860208284010111156107ec57600080fd5b959894975092955050506020019190565b60006020828403121561080f57600080fd5b61081882610737565b9392505050565b6000806040838503121561083257600080fd5b50508035926020909101359150565b60008060006060848603121561085657600080fd5b61085f84610737565b925061086d60208501610737565b915061087b60408501610737565b90509250925092565b6000806040838503121561089757600080fd5b823591506108a760208401610737565b90509250929050565b6000602082840312156108c257600080fd5b8151801515811461081857600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60808152600086516060608084015280518060e085015260005b81811015610953576020818401810151610100878401015201610935565b5060008482016101000152602089015173ffffffffffffffffffffffffffffffffffffffff811660a0860152601f9091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168401915050604088015160c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401526101008382030160608401526109f2610100820185876108d2565b9897505050505050505056fea26469706673582212205fde1ebacb4d4652487714e60fafa80a7ff13f41f31f2319e87b5afbaeff4d0264736f6c634300081a0033";

type SystemContractMockConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: SystemContractMockConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class SystemContractMock__factory extends ContractFactory {
  constructor(...args: SystemContractMockConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    wzeta_: AddressLike,
    uniswapv2Factory_: AddressLike,
    uniswapv2Router02_: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(
      wzeta_,
      uniswapv2Factory_,
      uniswapv2Router02_,
      overrides || {}
    );
  }
  override deploy(
    wzeta_: AddressLike,
    uniswapv2Factory_: AddressLike,
    uniswapv2Router02_: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(
      wzeta_,
      uniswapv2Factory_,
      uniswapv2Router02_,
      overrides || {}
    ) as Promise<
      SystemContractMock & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): SystemContractMock__factory {
    return super.connect(runner) as SystemContractMock__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): SystemContractMockInterface {
    return new Interface(_abi) as SystemContractMockInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): SystemContractMock {
    return new Contract(address, _abi, runner) as unknown as SystemContractMock;
  }
}
