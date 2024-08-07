// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetatokenconsumertrident

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ZetaTokenConsumerTridentMetaData contains all meta data concerning the ZetaTokenConsumerTrident contract.
var ZetaTokenConsumerTridentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zetaToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV3Router_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WETH9Address_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolFactory_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorSendingETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getZetaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getZetaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasZetaLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolFactory\",\"outputs\":[{\"internalType\":\"contractConcentratedLiquidityPoolFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tridentRouter\",\"outputs\":[{\"internalType\":\"contractIPoolRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162002b5a38038062002b5a833981810160405281019062000038919062000245565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480620000a05750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b80620000d85750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b80620001105750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b1562000148576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508273ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b81525050505050506200030a565b6000815190506200023f81620002f0565b92915050565b60008060008060808587031215620002625762000261620002eb565b5b600062000272878288016200022e565b945050602062000285878288016200022e565b935050604062000298878288016200022e565b9250506060620002ab878288016200022e565b91505092959194509250565b6000620002c482620002cb565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620002fb81620002b7565b81146200030757600080fd5b50565b60805160601c60a05160601c60c05160601c60e05160601c6127466200041460003960008181610316015281816107000152818161080c01528181610b4201528181610d14015281816111e301526112cf0152600081816104620152818161068501528181610a4801528181610c5901528181610e7f01528181610f7401528181611128015261152b0152600081816102ea01528181610557015281816107dc01528181610c0f01528181610c7b01528181610cc701528181610dd9015281816110de0152818161114a0152818161119601526114b60152600081816102c9015281816106d4015281816107bb01528181610ce8015281816111b7015261129e01526127466000f3fe60806040526004361061007f5760003560e01c806354c49a2a1161004e57806354c49a2a1461014e57806364b5528a1461018b57806380801f84146101b6578063a53fb10b146101e157610086565b8063013b2ff81461008b57806321e093b1146100bb5780632405620a146100e65780634219dc401461012357610086565b3661008657005b600080fd5b6100a560048036038101906100a09190611c10565b61021e565b6040516100b2919061231a565b60405180910390f35b3480156100c757600080fd5b506100d0610555565b6040516100dd91906120ca565b60405180910390f35b3480156100f257600080fd5b5061010d60048036038101906101089190611c50565b610579565b60405161011a919061231a565b60405180910390f35b34801561012f57600080fd5b50610138610b40565b6040516101459190612205565b60405180910390f35b34801561015a57600080fd5b5061017560048036038101906101709190611cb7565b610b64565b604051610182919061231a565b60405180910390f35b34801561019757600080fd5b506101a0610f72565b6040516101ad9190612220565b60405180910390f35b3480156101c257600080fd5b506101cb610f96565b6040516101d891906121ea565b60405180910390f35b3480156101ed57600080fd5b5061020860048036038101906102039190611c50565b610f9b565b604051610215919061231a565b60405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610286576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60003414156102c1576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008061030e7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000061163d565b9150915060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128484600060016040518563ffffffff1660e01b8152600401610375949392919061210e565b60006040518083038186803b15801561038d57600080fd5b505afa1580156103a1573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103ca9190611d0a565b905060006040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020013481526020018781526020018360008151811061041657610415612532565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200160001515815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c07f5c3234846040518363ffffffff1660e01b81526004016104ba91906122ff565b6020604051808303818588803b1580156104d357600080fd5b505af11580156104e7573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061050c9190611d80565b90507f87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11348260405161053f929190612335565b60405180910390a1809550505050505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806105e15750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610618576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610653576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106803330848673ffffffffffffffffffffffffffffffffffffffff1661168d909392919063ffffffff16565b6106cb7f0000000000000000000000000000000000000000000000000000000000000000838573ffffffffffffffffffffffffffffffffffffffff166117169092919063ffffffff16565b6000806106f8857f000000000000000000000000000000000000000000000000000000000000000061163d565b9150915060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128484600060016040518563ffffffff1660e01b815260040161075f949392919061210e565b60006040518083038186803b15801561077757600080fd5b505afa15801561078b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906107b49190611d0a565b90506108007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000061163d565b809350819450505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128585600060016040518563ffffffff1660e01b815260040161086b949392919061210e565b60006040518083038186803b15801561088357600080fd5b505afa158015610897573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906108c09190611d0a565b90506000600267ffffffffffffffff8111156108df576108de612561565b5b60405190808252806020026020018201604052801561090d5781602001602082028036833780820191505090505b5090508260008151811061092457610923612532565b5b6020026020010151816000815181106109405761093f612532565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508160008151811061098e5761098d612532565b5b6020026020010151816001815181106109aa576109a9612532565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060006040518060c001604052808a73ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018b81526020018381526020018c73ffffffffffffffffffffffffffffffffffffffff16815260200160001515815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663363a9dba836040518263ffffffff1660e01b8152600401610a9f91906122dd565b602060405180830381600087803b158015610ab957600080fd5b505af1158015610acd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af19190611d80565b90507f017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f8a8a83604051610b26939291906121b3565b60405180910390a180975050505050505050949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610bcc576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610c07576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c543330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661168d909392919063ffffffff16565b610cbf7f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166117169092919063ffffffff16565b600080610d0c7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000061163d565b9150915060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128484600060016040518563ffffffff1660e01b8152600401610d73949392919061210e565b60006040518083038186803b158015610d8b57600080fd5b505afa158015610d9f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610dc89190611d0a565b905060006040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815260200187815260200188815260200183600081518110610e3357610e32612532565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff16815260200160011515815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c07f5c32836040518263ffffffff1660e01b8152600401610ed691906122ff565b602060405180830381600087803b158015610ef057600080fd5b505af1158015610f04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f289190611d80565b90507f74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae8782604051610f5b929190612335565b60405180910390a180955050505050509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b600090565b60008060009054906101000a900460ff1615610fe3576040517f29f745a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016000806101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806110645750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b1561109b576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008214156110d6576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111233330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661168d909392919063ffffffff16565b61118e7f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166117169092919063ffffffff16565b6000806111db7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000061163d565b9150915060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128484600060016040518563ffffffff1660e01b8152600401611242949392919061210e565b60006040518083038186803b15801561125a57600080fd5b505afa15801561126e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906112979190611d0a565b90506112c37f00000000000000000000000000000000000000000000000000000000000000008761163d565b809350819450505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128585600060016040518563ffffffff1660e01b815260040161132e949392919061210e565b60006040518083038186803b15801561134657600080fd5b505afa15801561135a573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906113839190611d0a565b90506000600267ffffffffffffffff8111156113a2576113a1612561565b5b6040519080825280602002602001820160405280156113d05781602001602082028036833780820191505090505b509050826000815181106113e7576113e6612532565b5b60200260200101518160008151811061140357611402612532565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508160008151811061145157611450612532565b5b60200260200101518160018151811061146d5761146c612532565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060006040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018b81526020018381526020018c73ffffffffffffffffffffffffffffffffffffffff16815260200160001515815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663363a9dba836040518263ffffffff1660e01b815260040161158291906122dd565b602060405180830381600087803b15801561159c57600080fd5b505af11580156115b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d49190611d80565b90507f0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b8a8a83604051611609939291906121b3565b60405180910390a18097505050505050505060008060006101000a81548160ff021916908315150217905550949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16101561167f57838391509150611686565b8284915091505b9250929050565b611710846323b872dd60e01b8585856040516024016116ae93929190612153565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611874565b50505050565b60008114806117af575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b815260040161175d9291906120e5565b60206040518083038186803b15801561177557600080fd5b505afa158015611789573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117ad9190611d80565b145b6117ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117e5906122bd565b60405180910390fd5b61186f8363095ea7b360e01b848460405160240161180d92919061218a565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611874565b505050565b60006118d6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661193b9092919063ffffffff16565b905060008151111561193657808060200190518101906118f69190611d53565b611935576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161192c9061229d565b60405180910390fd5b5b505050565b606061194a8484600085611953565b90509392505050565b606082471015611998576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161198f9061225d565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516119c191906120b3565b60006040518083038185875af1925050503d80600081146119fe576040519150601f19603f3d011682016040523d82523d6000602084013e611a03565b606091505b5091509150611a1487838387611a20565b92505050949350505050565b60608315611a8357600083511415611a7b57611a3b85611a96565b611a7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a719061227d565b60405180910390fd5b5b829050611a8e565b611a8d8383611ab9565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082511115611acc5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b00919061223b565b60405180910390fd5b6000611b1c611b1784612383565b61235e565b90508083825260208201905082856020860282011115611b3f57611b3e612595565b5b60005b85811015611b6f5781611b558882611b8e565b845260208401935060208301925050600181019050611b42565b5050509392505050565b600081359050611b88816126cb565b92915050565b600081519050611b9d816126cb565b92915050565b600082601f830112611bb857611bb7612590565b5b8151611bc8848260208601611b09565b91505092915050565b600081519050611be0816126e2565b92915050565b600081359050611bf5816126f9565b92915050565b600081519050611c0a816126f9565b92915050565b60008060408385031215611c2757611c2661259f565b5b6000611c3585828601611b79565b9250506020611c4685828601611be6565b9150509250929050565b60008060008060808587031215611c6a57611c6961259f565b5b6000611c7887828801611b79565b9450506020611c8987828801611be6565b9350506040611c9a87828801611b79565b9250506060611cab87828801611be6565b91505092959194509250565b600080600060608486031215611cd057611ccf61259f565b5b6000611cde86828701611b79565b9350506020611cef86828701611be6565b9250506040611d0086828701611be6565b9150509250925092565b600060208284031215611d2057611d1f61259f565b5b600082015167ffffffffffffffff811115611d3e57611d3d61259a565b5b611d4a84828501611ba3565b91505092915050565b600060208284031215611d6957611d6861259f565b5b6000611d7784828501611bd1565b91505092915050565b600060208284031215611d9657611d9561259f565b5b6000611da484828501611bfb565b91505092915050565b6000611db98383611dc5565b60208301905092915050565b611dce8161241a565b82525050565b611ddd8161241a565b82525050565b6000611dee826123bf565b611df881856123ed565b9350611e03836123af565b8060005b83811015611e34578151611e1b8882611dad565b9750611e26836123e0565b925050600181019050611e07565b5085935050505092915050565b611e4a8161242c565b82525050565b611e598161242c565b82525050565b6000611e6a826123ca565b611e7481856123fe565b9350611e848185602086016124ce565b80840191505092915050565b611e9981612462565b82525050565b611ea881612474565b82525050565b611eb781612486565b82525050565b611ec681612498565b82525050565b6000611ed7826123d5565b611ee18185612409565b9350611ef18185602086016124ce565b611efa816125a4565b840191505092915050565b6000611f12602683612409565b9150611f1d826125b5565b604082019050919050565b6000611f35601d83612409565b9150611f4082612604565b602082019050919050565b6000611f58602a83612409565b9150611f638261262d565b604082019050919050565b6000611f7b603683612409565b9150611f868261267c565b604082019050919050565b600060c083016000830151611fa96000860182611dc5565b506020830151611fbc6020860182612095565b506040830151611fcf6040860182612095565b5060608301518482036060860152611fe78282611de3565b9150506080830151611ffc6080860182611dc5565b5060a083015161200f60a0860182611e41565b508091505092915050565b60c0820160008201516120306000850182611dc5565b5060208201516120436020850182612095565b5060408201516120566040850182612095565b5060608201516120696060850182611dc5565b50608082015161207c6080850182611dc5565b5060a082015161208f60a0850182611e41565b50505050565b61209e81612458565b82525050565b6120ad81612458565b82525050565b60006120bf8284611e5f565b915081905092915050565b60006020820190506120df6000830184611dd4565b92915050565b60006040820190506120fa6000830185611dd4565b6121076020830184611dd4565b9392505050565b60006080820190506121236000830187611dd4565b6121306020830186611dd4565b61213d6040830185611eae565b61214a6060830184611ebd565b95945050505050565b60006060820190506121686000830186611dd4565b6121756020830185611dd4565b61218260408301846120a4565b949350505050565b600060408201905061219f6000830185611dd4565b6121ac60208301846120a4565b9392505050565b60006060820190506121c86000830186611dd4565b6121d560208301856120a4565b6121e260408301846120a4565b949350505050565b60006020820190506121ff6000830184611e50565b92915050565b600060208201905061221a6000830184611e90565b92915050565b60006020820190506122356000830184611e9f565b92915050565b600060208201905081810360008301526122558184611ecc565b905092915050565b6000602082019050818103600083015261227681611f05565b9050919050565b6000602082019050818103600083015261229681611f28565b9050919050565b600060208201905081810360008301526122b681611f4b565b9050919050565b600060208201905081810360008301526122d681611f6e565b9050919050565b600060208201905081810360008301526122f78184611f91565b905092915050565b600060c082019050612314600083018461201a565b92915050565b600060208201905061232f60008301846120a4565b92915050565b600060408201905061234a60008301856120a4565b61235760208301846120a4565b9392505050565b6000612368612379565b90506123748282612501565b919050565b6000604051905090565b600067ffffffffffffffff82111561239e5761239d612561565b5b602082029050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061242582612438565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061246d826124aa565b9050919050565b600061247f826124aa565b9050919050565b600061249182612458565b9050919050565b60006124a382612458565b9050919050565b60006124b5826124bc565b9050919050565b60006124c782612438565b9050919050565b60005b838110156124ec5780820151818401526020810190506124d1565b838111156124fb576000848401525b50505050565b61250a826125a4565b810181811067ffffffffffffffff8211171561252957612528612561565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b6126d48161241a565b81146126df57600080fd5b50565b6126eb8161242c565b81146126f657600080fd5b50565b61270281612458565b811461270d57600080fd5b5056fea264697066735822122052bfd21f617a098a8e9f7179f87b0511ff7b45f29043d147ad39b847dabfaf1364736f6c63430008070033",
}

// ZetaTokenConsumerTridentABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaTokenConsumerTridentMetaData.ABI instead.
var ZetaTokenConsumerTridentABI = ZetaTokenConsumerTridentMetaData.ABI

// ZetaTokenConsumerTridentBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaTokenConsumerTridentMetaData.Bin instead.
var ZetaTokenConsumerTridentBin = ZetaTokenConsumerTridentMetaData.Bin

// DeployZetaTokenConsumerTrident deploys a new Ethereum contract, binding an instance of ZetaTokenConsumerTrident to it.
func DeployZetaTokenConsumerTrident(auth *bind.TransactOpts, backend bind.ContractBackend, zetaToken_ common.Address, uniswapV3Router_ common.Address, WETH9Address_ common.Address, poolFactory_ common.Address) (common.Address, *types.Transaction, *ZetaTokenConsumerTrident, error) {
	parsed, err := ZetaTokenConsumerTridentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaTokenConsumerTridentBin), backend, zetaToken_, uniswapV3Router_, WETH9Address_, poolFactory_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaTokenConsumerTrident{ZetaTokenConsumerTridentCaller: ZetaTokenConsumerTridentCaller{contract: contract}, ZetaTokenConsumerTridentTransactor: ZetaTokenConsumerTridentTransactor{contract: contract}, ZetaTokenConsumerTridentFilterer: ZetaTokenConsumerTridentFilterer{contract: contract}}, nil
}

// ZetaTokenConsumerTrident is an auto generated Go binding around an Ethereum contract.
type ZetaTokenConsumerTrident struct {
	ZetaTokenConsumerTridentCaller     // Read-only binding to the contract
	ZetaTokenConsumerTridentTransactor // Write-only binding to the contract
	ZetaTokenConsumerTridentFilterer   // Log filterer for contract events
}

// ZetaTokenConsumerTridentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerTridentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerTridentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaTokenConsumerTridentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerTridentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaTokenConsumerTridentSession struct {
	Contract     *ZetaTokenConsumerTrident // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerTridentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaTokenConsumerTridentCallerSession struct {
	Contract *ZetaTokenConsumerTridentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ZetaTokenConsumerTridentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaTokenConsumerTridentTransactorSession struct {
	Contract     *ZetaTokenConsumerTridentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerTridentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentRaw struct {
	Contract *ZetaTokenConsumerTrident // Generic contract binding to access the raw methods on
}

// ZetaTokenConsumerTridentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentCallerRaw struct {
	Contract *ZetaTokenConsumerTridentCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaTokenConsumerTridentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentTransactorRaw struct {
	Contract *ZetaTokenConsumerTridentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaTokenConsumerTrident creates a new instance of ZetaTokenConsumerTrident, bound to a specific deployed contract.
func NewZetaTokenConsumerTrident(address common.Address, backend bind.ContractBackend) (*ZetaTokenConsumerTrident, error) {
	contract, err := bindZetaTokenConsumerTrident(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTrident{ZetaTokenConsumerTridentCaller: ZetaTokenConsumerTridentCaller{contract: contract}, ZetaTokenConsumerTridentTransactor: ZetaTokenConsumerTridentTransactor{contract: contract}, ZetaTokenConsumerTridentFilterer: ZetaTokenConsumerTridentFilterer{contract: contract}}, nil
}

// NewZetaTokenConsumerTridentCaller creates a new read-only instance of ZetaTokenConsumerTrident, bound to a specific deployed contract.
func NewZetaTokenConsumerTridentCaller(address common.Address, caller bind.ContractCaller) (*ZetaTokenConsumerTridentCaller, error) {
	contract, err := bindZetaTokenConsumerTrident(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentCaller{contract: contract}, nil
}

// NewZetaTokenConsumerTridentTransactor creates a new write-only instance of ZetaTokenConsumerTrident, bound to a specific deployed contract.
func NewZetaTokenConsumerTridentTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaTokenConsumerTridentTransactor, error) {
	contract, err := bindZetaTokenConsumerTrident(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentTransactor{contract: contract}, nil
}

// NewZetaTokenConsumerTridentFilterer creates a new log filterer instance of ZetaTokenConsumerTrident, bound to a specific deployed contract.
func NewZetaTokenConsumerTridentFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaTokenConsumerTridentFilterer, error) {
	contract, err := bindZetaTokenConsumerTrident(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentFilterer{contract: contract}, nil
}

// bindZetaTokenConsumerTrident binds a generic wrapper to an already deployed contract.
func bindZetaTokenConsumerTrident(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaTokenConsumerTridentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerTrident.Contract.ZetaTokenConsumerTridentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.ZetaTokenConsumerTridentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.ZetaTokenConsumerTridentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerTrident.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.contract.Transact(opts, method, params...)
}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCaller) HasZetaLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaTokenConsumerTrident.contract.Call(opts, &out, "hasZetaLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) HasZetaLiquidity() (bool, error) {
	return _ZetaTokenConsumerTrident.Contract.HasZetaLiquidity(&_ZetaTokenConsumerTrident.CallOpts)
}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCallerSession) HasZetaLiquidity() (bool, error) {
	return _ZetaTokenConsumerTrident.Contract.HasZetaLiquidity(&_ZetaTokenConsumerTrident.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCaller) PoolFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerTrident.contract.Call(opts, &out, "poolFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) PoolFactory() (common.Address, error) {
	return _ZetaTokenConsumerTrident.Contract.PoolFactory(&_ZetaTokenConsumerTrident.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCallerSession) PoolFactory() (common.Address, error) {
	return _ZetaTokenConsumerTrident.Contract.PoolFactory(&_ZetaTokenConsumerTrident.CallOpts)
}

// TridentRouter is a free data retrieval call binding the contract method 0x64b5528a.
//
// Solidity: function tridentRouter() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCaller) TridentRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerTrident.contract.Call(opts, &out, "tridentRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TridentRouter is a free data retrieval call binding the contract method 0x64b5528a.
//
// Solidity: function tridentRouter() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) TridentRouter() (common.Address, error) {
	return _ZetaTokenConsumerTrident.Contract.TridentRouter(&_ZetaTokenConsumerTrident.CallOpts)
}

// TridentRouter is a free data retrieval call binding the contract method 0x64b5528a.
//
// Solidity: function tridentRouter() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCallerSession) TridentRouter() (common.Address, error) {
	return _ZetaTokenConsumerTrident.Contract.TridentRouter(&_ZetaTokenConsumerTrident.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerTrident.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) ZetaToken() (common.Address, error) {
	return _ZetaTokenConsumerTrident.Contract.ZetaToken(&_ZetaTokenConsumerTrident.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaTokenConsumerTrident.Contract.ZetaToken(&_ZetaTokenConsumerTrident.CallOpts)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactor) GetEthFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.contract.Transact(opts, "getEthFromZeta", destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetEthFromZeta(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetEthFromZeta(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactor) GetTokenFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.contract.Transact(opts, "getTokenFromZeta", destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetTokenFromZeta(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetTokenFromZeta(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactor) GetZetaFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.contract.Transact(opts, "getZetaFromEth", destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetZetaFromEth(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetZetaFromEth(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactor) GetZetaFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.contract.Transact(opts, "getZetaFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetZetaFromToken(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetZetaFromToken(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) Receive() (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.Receive(&_ZetaTokenConsumerTrident.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorSession) Receive() (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.Receive(&_ZetaTokenConsumerTrident.TransactOpts)
}

// ZetaTokenConsumerTridentEthExchangedForZetaIterator is returned from FilterEthExchangedForZeta and is used to iterate over the raw logs and unpacked data for EthExchangedForZeta events raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentEthExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerTridentEthExchangedForZeta // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaTokenConsumerTridentEthExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerTridentEthExchangedForZeta)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaTokenConsumerTridentEthExchangedForZeta)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaTokenConsumerTridentEthExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerTridentEthExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerTridentEthExchangedForZeta represents a EthExchangedForZeta event raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentEthExchangedForZeta struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForZeta is a free log retrieval operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) FilterEthExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerTridentEthExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.FilterLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentEthExchangedForZetaIterator{contract: _ZetaTokenConsumerTrident.contract, event: "EthExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForZeta is a free log subscription operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) WatchEthExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerTridentEthExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.WatchLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerTridentEthExchangedForZeta)
				if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEthExchangedForZeta is a log parse operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) ParseEthExchangedForZeta(log types.Log) (*ZetaTokenConsumerTridentEthExchangedForZeta, error) {
	event := new(ZetaTokenConsumerTridentEthExchangedForZeta)
	if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerTridentTokenExchangedForZetaIterator is returned from FilterTokenExchangedForZeta and is used to iterate over the raw logs and unpacked data for TokenExchangedForZeta events raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentTokenExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerTridentTokenExchangedForZeta // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaTokenConsumerTridentTokenExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerTridentTokenExchangedForZeta)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaTokenConsumerTridentTokenExchangedForZeta)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaTokenConsumerTridentTokenExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerTridentTokenExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerTridentTokenExchangedForZeta represents a TokenExchangedForZeta event raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentTokenExchangedForZeta struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForZeta is a free log retrieval operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) FilterTokenExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerTridentTokenExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.FilterLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentTokenExchangedForZetaIterator{contract: _ZetaTokenConsumerTrident.contract, event: "TokenExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForZeta is a free log subscription operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) WatchTokenExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerTridentTokenExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.WatchLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerTridentTokenExchangedForZeta)
				if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenExchangedForZeta is a log parse operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) ParseTokenExchangedForZeta(log types.Log) (*ZetaTokenConsumerTridentTokenExchangedForZeta, error) {
	event := new(ZetaTokenConsumerTridentTokenExchangedForZeta)
	if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerTridentZetaExchangedForEthIterator is returned from FilterZetaExchangedForEth and is used to iterate over the raw logs and unpacked data for ZetaExchangedForEth events raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentZetaExchangedForEthIterator struct {
	Event *ZetaTokenConsumerTridentZetaExchangedForEth // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaTokenConsumerTridentZetaExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerTridentZetaExchangedForEth)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaTokenConsumerTridentZetaExchangedForEth)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaTokenConsumerTridentZetaExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerTridentZetaExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerTridentZetaExchangedForEth represents a ZetaExchangedForEth event raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentZetaExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForEth is a free log retrieval operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) FilterZetaExchangedForEth(opts *bind.FilterOpts) (*ZetaTokenConsumerTridentZetaExchangedForEthIterator, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.FilterLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentZetaExchangedForEthIterator{contract: _ZetaTokenConsumerTrident.contract, event: "ZetaExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForEth is a free log subscription operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) WatchZetaExchangedForEth(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerTridentZetaExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.WatchLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerTridentZetaExchangedForEth)
				if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseZetaExchangedForEth is a log parse operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) ParseZetaExchangedForEth(log types.Log) (*ZetaTokenConsumerTridentZetaExchangedForEth, error) {
	event := new(ZetaTokenConsumerTridentZetaExchangedForEth)
	if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerTridentZetaExchangedForTokenIterator is returned from FilterZetaExchangedForToken and is used to iterate over the raw logs and unpacked data for ZetaExchangedForToken events raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentZetaExchangedForTokenIterator struct {
	Event *ZetaTokenConsumerTridentZetaExchangedForToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaTokenConsumerTridentZetaExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerTridentZetaExchangedForToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaTokenConsumerTridentZetaExchangedForToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaTokenConsumerTridentZetaExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerTridentZetaExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerTridentZetaExchangedForToken represents a ZetaExchangedForToken event raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentZetaExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForToken is a free log retrieval operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) FilterZetaExchangedForToken(opts *bind.FilterOpts) (*ZetaTokenConsumerTridentZetaExchangedForTokenIterator, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.FilterLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentZetaExchangedForTokenIterator{contract: _ZetaTokenConsumerTrident.contract, event: "ZetaExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForToken is a free log subscription operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) WatchZetaExchangedForToken(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerTridentZetaExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.WatchLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerTridentZetaExchangedForToken)
				if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseZetaExchangedForToken is a log parse operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) ParseZetaExchangedForToken(log types.Log) (*ZetaTokenConsumerTridentZetaExchangedForToken, error) {
	event := new(ZetaTokenConsumerTridentZetaExchangedForToken)
	if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
