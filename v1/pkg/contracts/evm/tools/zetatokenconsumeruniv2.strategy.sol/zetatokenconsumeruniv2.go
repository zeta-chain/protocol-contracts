// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetatokenconsumeruniv2

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

// ZetaTokenConsumerUniV2MetaData contains all meta data concerning the ZetaTokenConsumerUniV2 contract.
var ZetaTokenConsumerUniV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zetaToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV2Router_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getZetaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getZetaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasZetaLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620029a0380380620029a083398181016040528101906200003791906200024e565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806200009f5750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15620000d7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1663ad5c46486040518163ffffffff1660e01b815260040160206040518083038186803b1580156200018c57600080fd5b505afa158015620001a1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001c791906200021c565b73ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250505050620002e8565b6000815190506200021681620002ce565b92915050565b600060208284031215620002355762000234620002c9565b5b6000620002458482850162000205565b91505092915050565b60008060408385031215620002685762000267620002c9565b5b6000620002788582860162000205565b92505060206200028b8582860162000205565b9150509250929050565b6000620002a282620002a9565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620002d98162000295565b8114620002e557600080fd5b50565b60805160601c60a05160601c60c05160601c6125c9620003d76000396000818161036a015281816105ce0152818161091701528181610b4501528181610cdb01528181610f400152818161115501526114be0152600081816102f9015281816104a001528181610727015281816108a501528181610afb01528181610b6701528181610bfb01528181610ed10152818161110b015281816111770152818161125f015261138e01526000818161028a01528181610618015281816106b80152818161083601528181610c6a01528181610e62015281816111bf015281816112ce01526113fd01526125c96000f3fe6080604052600436106100555760003560e01c8063013b2ff81461005a57806321e093b11461008a5780632405620a146100b557806354c49a2a146100f257806380801f841461012f578063a53fb10b1461015a575b600080fd5b610074600480360381019061006f9190611b65565b610197565b6040516100819190612098565b60405180910390f35b34801561009657600080fd5b5061009f61049e565b6040516100ac9190611ed0565b60405180910390f35b3480156100c157600080fd5b506100dc60048036038101906100d79190611ba5565b6104c2565b6040516100e99190612098565b60405180910390f35b3480156100fe57600080fd5b5061011960048036038101906101149190611c0c565b610a50565b6040516101269190612098565b60405180910390f35b34801561013b57600080fd5b50610144610e11565b6040516101519190611fab565b60405180910390f35b34801561016657600080fd5b50610181600480360381019061017c9190611ba5565b611029565b60405161018e9190612098565b60405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156101ff576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600034141561023a576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600267ffffffffffffffff811115610257576102566123e4565b5b6040519080825280602002602001820160405280156102855781602001602082028036833780820191505090505b5090507f0000000000000000000000000000000000000000000000000000000000000000816000815181106102bd576102bc6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f00000000000000000000000000000000000000000000000000000000000000008160018151811061032c5761032b6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16637ff36ab53486858960c8426103b5919061223e565b6040518663ffffffff1660e01b81526004016103d494939291906120b3565b6000604051808303818588803b1580156103ed57600080fd5b505af1158015610401573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f8201168201806040525081019061042b9190611c5f565b90506000816001845161043e9190612294565b8151811061044f5761044e6123b5565b5b602002602001015190507f87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11348260405161048a9291906120ff565b60405180910390a180935050505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16148061052a5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610561576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082141561059c576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105c93330848673ffffffffffffffffffffffffffffffffffffffff166115f7909392919063ffffffff16565b6106147f0000000000000000000000000000000000000000000000000000000000000000838573ffffffffffffffffffffffffffffffffffffffff166116809092919063ffffffff16565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561079957600267ffffffffffffffff811115610685576106846123e4565b5b6040519080825280602002602001820160405280156106b35781602001602082028036833780820191505090505b5090507f0000000000000000000000000000000000000000000000000000000000000000816000815181106106eb576106ea6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f00000000000000000000000000000000000000000000000000000000000000008160018151811061075a576107596123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050610913565b600367ffffffffffffffff8111156107b4576107b36123e4565b5b6040519080825280602002602001820160405280156107e25781602001602082028036833780820191505090505b50905083816000815181106107fa576107f96123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000081600181518110610869576108686123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f0000000000000000000000000000000000000000000000000000000000000000816002815181106108d8576108d76123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166338ed17398588858b60c842610962919061223e565b6040518663ffffffff1660e01b8152600401610982959493929190612128565b600060405180830381600087803b15801561099c57600080fd5b505af11580156109b0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906109d99190611c5f565b9050600081600184516109ec9190612294565b815181106109fd576109fc6123b5565b5b602002602001015190507f017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f868683604051610a3a93929190611f74565b60405180910390a1809350505050949350505050565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610ab8576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610af3576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b403330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166115f7909392919063ffffffff16565b610bab7f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166116809092919063ffffffff16565b6000600267ffffffffffffffff811115610bc857610bc76123e4565b5b604051908082528060200260200182016040528015610bf65781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110610c2e57610c2d6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000081600181518110610c9d57610c9c6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318cbafe58587858a60c842610d26919061223e565b6040518663ffffffff1660e01b8152600401610d46959493929190612128565b600060405180830381600087803b158015610d6057600080fd5b505af1158015610d74573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610d9d9190611c5f565b905060008160018451610db09190612294565b81518110610dc157610dc06123b5565b5b602002602001015190507f74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae8582604051610dfc9291906120ff565b60405180910390a18093505050509392505050565b600080600267ffffffffffffffff811115610e2f57610e2e6123e4565b5b604051908082528060200260200182016040528015610e5d5781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110610e9557610e946123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000081600181518110610f0457610f036123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d06ca61f6001836040518363ffffffff1660e01b8152600401610f9a929190611fc6565b60006040518083038186803b158015610fb257600080fd5b505afa925050508015610fe857506040513d6000823e3d601f19601f82011682018060405250810190610fe59190611c5f565b60015b610ff6576000915050611026565b600081600184516110079190612294565b81518110611018576110176123b5565b5b602002602001015111925050505b90565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806110915750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b156110c8576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415611103576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111503330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166115f7909392919063ffffffff16565b6111bb7f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166116809092919063ffffffff16565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561134057600267ffffffffffffffff81111561122c5761122b6123e4565b5b60405190808252806020026020018201604052801561125a5781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110611292576112916123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000081600181518110611301576113006123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506114ba565b600367ffffffffffffffff81111561135b5761135a6123e4565b5b6040519080825280602002602001820160405280156113895781602001602082028036833780820191505090505b5090507f0000000000000000000000000000000000000000000000000000000000000000816000815181106113c1576113c06123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f0000000000000000000000000000000000000000000000000000000000000000816001815181106114305761142f6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050838160028151811061147f5761147e6123b5565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166338ed17398588858b60c842611509919061223e565b6040518663ffffffff1660e01b8152600401611529959493929190612128565b600060405180830381600087803b15801561154357600080fd5b505af1158015611557573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906115809190611c5f565b9050600081600184516115939190612294565b815181106115a4576115a36123b5565b5b602002602001015190507f0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b8686836040516115e193929190611f74565b60405180910390a1809350505050949350505050565b61167a846323b872dd60e01b85858560405160240161161893929190611f14565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506117de565b50505050565b6000811480611719575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b81526004016116c7929190611eeb565b60206040518083038186803b1580156116df57600080fd5b505afa1580156116f3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117179190611cd5565b145b611758576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161174f90612078565b60405180910390fd5b6117d98363095ea7b360e01b8484604051602401611777929190611f4b565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506117de565b505050565b6000611840826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166118a59092919063ffffffff16565b90506000815111156118a057808060200190518101906118609190611ca8565b61189f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161189690612058565b60405180910390fd5b5b505050565b60606118b484846000856118bd565b90509392505050565b606082471015611902576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118f990612018565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161192b9190611eb9565b60006040518083038185875af1925050503d8060008114611968576040519150601f19603f3d011682016040523d82523d6000602084013e61196d565b606091505b509150915061197e8783838761198a565b92505050949350505050565b606083156119ed576000835114156119e5576119a585611a00565b6119e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119db90612038565b60405180910390fd5b5b8290506119f8565b6119f78383611a23565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082511115611a365781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a6a9190611ff6565b60405180910390fd5b6000611a86611a81846121a7565b612182565b90508083825260208201905082856020860282011115611aa957611aa8612418565b5b60005b85811015611ad95781611abf8882611b50565b845260208401935060208301925050600181019050611aac565b5050509392505050565b600081359050611af28161254e565b92915050565b600082601f830112611b0d57611b0c612413565b5b8151611b1d848260208601611a73565b91505092915050565b600081519050611b3581612565565b92915050565b600081359050611b4a8161257c565b92915050565b600081519050611b5f8161257c565b92915050565b60008060408385031215611b7c57611b7b612422565b5b6000611b8a85828601611ae3565b9250506020611b9b85828601611b3b565b9150509250929050565b60008060008060808587031215611bbf57611bbe612422565b5b6000611bcd87828801611ae3565b9450506020611bde87828801611b3b565b9350506040611bef87828801611ae3565b9250506060611c0087828801611b3b565b91505092959194509250565b600080600060608486031215611c2557611c24612422565b5b6000611c3386828701611ae3565b9350506020611c4486828701611b3b565b9250506040611c5586828701611b3b565b9150509250925092565b600060208284031215611c7557611c74612422565b5b600082015167ffffffffffffffff811115611c9357611c9261241d565b5b611c9f84828501611af8565b91505092915050565b600060208284031215611cbe57611cbd612422565b5b6000611ccc84828501611b26565b91505092915050565b600060208284031215611ceb57611cea612422565b5b6000611cf984828501611b50565b91505092915050565b6000611d0e8383611d1a565b60208301905092915050565b611d23816122c8565b82525050565b611d32816122c8565b82525050565b6000611d43826121e3565b611d4d8185612211565b9350611d58836121d3565b8060005b83811015611d89578151611d708882611d02565b9750611d7b83612204565b925050600181019050611d5c565b5085935050505092915050565b611d9f816122da565b82525050565b6000611db0826121ee565b611dba8185612222565b9350611dca818560208601612322565b80840191505092915050565b611ddf81612310565b82525050565b6000611df0826121f9565b611dfa818561222d565b9350611e0a818560208601612322565b611e1381612427565b840191505092915050565b6000611e2b60268361222d565b9150611e3682612438565b604082019050919050565b6000611e4e601d8361222d565b9150611e5982612487565b602082019050919050565b6000611e71602a8361222d565b9150611e7c826124b0565b604082019050919050565b6000611e9460368361222d565b9150611e9f826124ff565b604082019050919050565b611eb381612306565b82525050565b6000611ec58284611da5565b915081905092915050565b6000602082019050611ee56000830184611d29565b92915050565b6000604082019050611f006000830185611d29565b611f0d6020830184611d29565b9392505050565b6000606082019050611f296000830186611d29565b611f366020830185611d29565b611f436040830184611eaa565b949350505050565b6000604082019050611f606000830185611d29565b611f6d6020830184611eaa565b9392505050565b6000606082019050611f896000830186611d29565b611f966020830185611eaa565b611fa36040830184611eaa565b949350505050565b6000602082019050611fc06000830184611d96565b92915050565b6000604082019050611fdb6000830185611dd6565b8181036020830152611fed8184611d38565b90509392505050565b600060208201905081810360008301526120108184611de5565b905092915050565b6000602082019050818103600083015261203181611e1e565b9050919050565b6000602082019050818103600083015261205181611e41565b9050919050565b6000602082019050818103600083015261207181611e64565b9050919050565b6000602082019050818103600083015261209181611e87565b9050919050565b60006020820190506120ad6000830184611eaa565b92915050565b60006080820190506120c86000830187611eaa565b81810360208301526120da8186611d38565b90506120e96040830185611d29565b6120f66060830184611eaa565b95945050505050565b60006040820190506121146000830185611eaa565b6121216020830184611eaa565b9392505050565b600060a08201905061213d6000830188611eaa565b61214a6020830187611eaa565b818103604083015261215c8186611d38565b905061216b6060830185611d29565b6121786080830184611eaa565b9695505050505050565b600061218c61219d565b90506121988282612355565b919050565b6000604051905090565b600067ffffffffffffffff8211156121c2576121c16123e4565b5b602082029050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061224982612306565b915061225483612306565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561228957612288612386565b5b828201905092915050565b600061229f82612306565b91506122aa83612306565b9250828210156122bd576122bc612386565b5b828203905092915050565b60006122d3826122e6565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061231b82612306565b9050919050565b60005b83811015612340578082015181840152602081019050612325565b8381111561234f576000848401525b50505050565b61235e82612427565b810181811067ffffffffffffffff8211171561237d5761237c6123e4565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b612557816122c8565b811461256257600080fd5b50565b61256e816122da565b811461257957600080fd5b50565b61258581612306565b811461259057600080fd5b5056fea264697066735822122043c2ef06b0d58aa4604eef6078aa0a569cff03f21dd48c315ad501590326c57264736f6c63430008070033",
}

// ZetaTokenConsumerUniV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaTokenConsumerUniV2MetaData.ABI instead.
var ZetaTokenConsumerUniV2ABI = ZetaTokenConsumerUniV2MetaData.ABI

// ZetaTokenConsumerUniV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaTokenConsumerUniV2MetaData.Bin instead.
var ZetaTokenConsumerUniV2Bin = ZetaTokenConsumerUniV2MetaData.Bin

// DeployZetaTokenConsumerUniV2 deploys a new Ethereum contract, binding an instance of ZetaTokenConsumerUniV2 to it.
func DeployZetaTokenConsumerUniV2(auth *bind.TransactOpts, backend bind.ContractBackend, zetaToken_ common.Address, uniswapV2Router_ common.Address) (common.Address, *types.Transaction, *ZetaTokenConsumerUniV2, error) {
	parsed, err := ZetaTokenConsumerUniV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaTokenConsumerUniV2Bin), backend, zetaToken_, uniswapV2Router_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaTokenConsumerUniV2{ZetaTokenConsumerUniV2Caller: ZetaTokenConsumerUniV2Caller{contract: contract}, ZetaTokenConsumerUniV2Transactor: ZetaTokenConsumerUniV2Transactor{contract: contract}, ZetaTokenConsumerUniV2Filterer: ZetaTokenConsumerUniV2Filterer{contract: contract}}, nil
}

// ZetaTokenConsumerUniV2 is an auto generated Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV2 struct {
	ZetaTokenConsumerUniV2Caller     // Read-only binding to the contract
	ZetaTokenConsumerUniV2Transactor // Write-only binding to the contract
	ZetaTokenConsumerUniV2Filterer   // Log filterer for contract events
}

// ZetaTokenConsumerUniV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerUniV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerUniV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaTokenConsumerUniV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerUniV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaTokenConsumerUniV2Session struct {
	Contract     *ZetaTokenConsumerUniV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerUniV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaTokenConsumerUniV2CallerSession struct {
	Contract *ZetaTokenConsumerUniV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ZetaTokenConsumerUniV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaTokenConsumerUniV2TransactorSession struct {
	Contract     *ZetaTokenConsumerUniV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerUniV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV2Raw struct {
	Contract *ZetaTokenConsumerUniV2 // Generic contract binding to access the raw methods on
}

// ZetaTokenConsumerUniV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV2CallerRaw struct {
	Contract *ZetaTokenConsumerUniV2Caller // Generic read-only contract binding to access the raw methods on
}

// ZetaTokenConsumerUniV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV2TransactorRaw struct {
	Contract *ZetaTokenConsumerUniV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaTokenConsumerUniV2 creates a new instance of ZetaTokenConsumerUniV2, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV2(address common.Address, backend bind.ContractBackend) (*ZetaTokenConsumerUniV2, error) {
	contract, err := bindZetaTokenConsumerUniV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV2{ZetaTokenConsumerUniV2Caller: ZetaTokenConsumerUniV2Caller{contract: contract}, ZetaTokenConsumerUniV2Transactor: ZetaTokenConsumerUniV2Transactor{contract: contract}, ZetaTokenConsumerUniV2Filterer: ZetaTokenConsumerUniV2Filterer{contract: contract}}, nil
}

// NewZetaTokenConsumerUniV2Caller creates a new read-only instance of ZetaTokenConsumerUniV2, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV2Caller(address common.Address, caller bind.ContractCaller) (*ZetaTokenConsumerUniV2Caller, error) {
	contract, err := bindZetaTokenConsumerUniV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV2Caller{contract: contract}, nil
}

// NewZetaTokenConsumerUniV2Transactor creates a new write-only instance of ZetaTokenConsumerUniV2, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ZetaTokenConsumerUniV2Transactor, error) {
	contract, err := bindZetaTokenConsumerUniV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV2Transactor{contract: contract}, nil
}

// NewZetaTokenConsumerUniV2Filterer creates a new log filterer instance of ZetaTokenConsumerUniV2, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ZetaTokenConsumerUniV2Filterer, error) {
	contract, err := bindZetaTokenConsumerUniV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV2Filterer{contract: contract}, nil
}

// bindZetaTokenConsumerUniV2 binds a generic wrapper to an already deployed contract.
func bindZetaTokenConsumerUniV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaTokenConsumerUniV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerUniV2.Contract.ZetaTokenConsumerUniV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.Contract.ZetaTokenConsumerUniV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.Contract.ZetaTokenConsumerUniV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerUniV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.Contract.contract.Transact(opts, method, params...)
}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Caller) HasZetaLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV2.contract.Call(opts, &out, "hasZetaLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Session) HasZetaLiquidity() (bool, error) {
	return _ZetaTokenConsumerUniV2.Contract.HasZetaLiquidity(&_ZetaTokenConsumerUniV2.CallOpts)
}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2CallerSession) HasZetaLiquidity() (bool, error) {
	return _ZetaTokenConsumerUniV2.Contract.HasZetaLiquidity(&_ZetaTokenConsumerUniV2.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Caller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV2.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Session) ZetaToken() (common.Address, error) {
	return _ZetaTokenConsumerUniV2.Contract.ZetaToken(&_ZetaTokenConsumerUniV2.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2CallerSession) ZetaToken() (common.Address, error) {
	return _ZetaTokenConsumerUniV2.Contract.ZetaToken(&_ZetaTokenConsumerUniV2.CallOpts)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Transactor) GetEthFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.contract.Transact(opts, "getEthFromZeta", destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Session) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.Contract.GetEthFromZeta(&_ZetaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2TransactorSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.Contract.GetEthFromZeta(&_ZetaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Transactor) GetTokenFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.contract.Transact(opts, "getTokenFromZeta", destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Session) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.Contract.GetTokenFromZeta(&_ZetaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2TransactorSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.Contract.GetTokenFromZeta(&_ZetaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Transactor) GetZetaFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.contract.Transact(opts, "getZetaFromEth", destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Session) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.Contract.GetZetaFromEth(&_ZetaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2TransactorSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.Contract.GetZetaFromEth(&_ZetaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Transactor) GetZetaFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.contract.Transact(opts, "getZetaFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Session) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.Contract.GetZetaFromToken(&_ZetaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2TransactorSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2.Contract.GetZetaFromToken(&_ZetaTokenConsumerUniV2.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// ZetaTokenConsumerUniV2EthExchangedForZetaIterator is returned from FilterEthExchangedForZeta and is used to iterate over the raw logs and unpacked data for EthExchangedForZeta events raised by the ZetaTokenConsumerUniV2 contract.
type ZetaTokenConsumerUniV2EthExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerUniV2EthExchangedForZeta // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerUniV2EthExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerUniV2EthExchangedForZeta)
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
		it.Event = new(ZetaTokenConsumerUniV2EthExchangedForZeta)
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
func (it *ZetaTokenConsumerUniV2EthExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerUniV2EthExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerUniV2EthExchangedForZeta represents a EthExchangedForZeta event raised by the ZetaTokenConsumerUniV2 contract.
type ZetaTokenConsumerUniV2EthExchangedForZeta struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForZeta is a free log retrieval operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Filterer) FilterEthExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerUniV2EthExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerUniV2.contract.FilterLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV2EthExchangedForZetaIterator{contract: _ZetaTokenConsumerUniV2.contract, event: "EthExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForZeta is a free log subscription operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Filterer) WatchEthExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerUniV2EthExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerUniV2.contract.WatchLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerUniV2EthExchangedForZeta)
				if err := _ZetaTokenConsumerUniV2.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
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
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Filterer) ParseEthExchangedForZeta(log types.Log) (*ZetaTokenConsumerUniV2EthExchangedForZeta, error) {
	event := new(ZetaTokenConsumerUniV2EthExchangedForZeta)
	if err := _ZetaTokenConsumerUniV2.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerUniV2TokenExchangedForZetaIterator is returned from FilterTokenExchangedForZeta and is used to iterate over the raw logs and unpacked data for TokenExchangedForZeta events raised by the ZetaTokenConsumerUniV2 contract.
type ZetaTokenConsumerUniV2TokenExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerUniV2TokenExchangedForZeta // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerUniV2TokenExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerUniV2TokenExchangedForZeta)
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
		it.Event = new(ZetaTokenConsumerUniV2TokenExchangedForZeta)
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
func (it *ZetaTokenConsumerUniV2TokenExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerUniV2TokenExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerUniV2TokenExchangedForZeta represents a TokenExchangedForZeta event raised by the ZetaTokenConsumerUniV2 contract.
type ZetaTokenConsumerUniV2TokenExchangedForZeta struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForZeta is a free log retrieval operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Filterer) FilterTokenExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerUniV2TokenExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerUniV2.contract.FilterLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV2TokenExchangedForZetaIterator{contract: _ZetaTokenConsumerUniV2.contract, event: "TokenExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForZeta is a free log subscription operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Filterer) WatchTokenExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerUniV2TokenExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerUniV2.contract.WatchLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerUniV2TokenExchangedForZeta)
				if err := _ZetaTokenConsumerUniV2.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
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
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Filterer) ParseTokenExchangedForZeta(log types.Log) (*ZetaTokenConsumerUniV2TokenExchangedForZeta, error) {
	event := new(ZetaTokenConsumerUniV2TokenExchangedForZeta)
	if err := _ZetaTokenConsumerUniV2.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerUniV2ZetaExchangedForEthIterator is returned from FilterZetaExchangedForEth and is used to iterate over the raw logs and unpacked data for ZetaExchangedForEth events raised by the ZetaTokenConsumerUniV2 contract.
type ZetaTokenConsumerUniV2ZetaExchangedForEthIterator struct {
	Event *ZetaTokenConsumerUniV2ZetaExchangedForEth // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerUniV2ZetaExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerUniV2ZetaExchangedForEth)
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
		it.Event = new(ZetaTokenConsumerUniV2ZetaExchangedForEth)
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
func (it *ZetaTokenConsumerUniV2ZetaExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerUniV2ZetaExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerUniV2ZetaExchangedForEth represents a ZetaExchangedForEth event raised by the ZetaTokenConsumerUniV2 contract.
type ZetaTokenConsumerUniV2ZetaExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForEth is a free log retrieval operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Filterer) FilterZetaExchangedForEth(opts *bind.FilterOpts) (*ZetaTokenConsumerUniV2ZetaExchangedForEthIterator, error) {

	logs, sub, err := _ZetaTokenConsumerUniV2.contract.FilterLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV2ZetaExchangedForEthIterator{contract: _ZetaTokenConsumerUniV2.contract, event: "ZetaExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForEth is a free log subscription operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Filterer) WatchZetaExchangedForEth(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerUniV2ZetaExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerUniV2.contract.WatchLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerUniV2ZetaExchangedForEth)
				if err := _ZetaTokenConsumerUniV2.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
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
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Filterer) ParseZetaExchangedForEth(log types.Log) (*ZetaTokenConsumerUniV2ZetaExchangedForEth, error) {
	event := new(ZetaTokenConsumerUniV2ZetaExchangedForEth)
	if err := _ZetaTokenConsumerUniV2.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerUniV2ZetaExchangedForTokenIterator is returned from FilterZetaExchangedForToken and is used to iterate over the raw logs and unpacked data for ZetaExchangedForToken events raised by the ZetaTokenConsumerUniV2 contract.
type ZetaTokenConsumerUniV2ZetaExchangedForTokenIterator struct {
	Event *ZetaTokenConsumerUniV2ZetaExchangedForToken // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerUniV2ZetaExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerUniV2ZetaExchangedForToken)
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
		it.Event = new(ZetaTokenConsumerUniV2ZetaExchangedForToken)
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
func (it *ZetaTokenConsumerUniV2ZetaExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerUniV2ZetaExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerUniV2ZetaExchangedForToken represents a ZetaExchangedForToken event raised by the ZetaTokenConsumerUniV2 contract.
type ZetaTokenConsumerUniV2ZetaExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForToken is a free log retrieval operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Filterer) FilterZetaExchangedForToken(opts *bind.FilterOpts) (*ZetaTokenConsumerUniV2ZetaExchangedForTokenIterator, error) {

	logs, sub, err := _ZetaTokenConsumerUniV2.contract.FilterLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV2ZetaExchangedForTokenIterator{contract: _ZetaTokenConsumerUniV2.contract, event: "ZetaExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForToken is a free log subscription operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Filterer) WatchZetaExchangedForToken(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerUniV2ZetaExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerUniV2.contract.WatchLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerUniV2ZetaExchangedForToken)
				if err := _ZetaTokenConsumerUniV2.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
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
func (_ZetaTokenConsumerUniV2 *ZetaTokenConsumerUniV2Filterer) ParseZetaExchangedForToken(log types.Log) (*ZetaTokenConsumerUniV2ZetaExchangedForToken, error) {
	event := new(ZetaTokenConsumerUniV2ZetaExchangedForToken)
	if err := _ZetaTokenConsumerUniV2.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
