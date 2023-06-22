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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zetaToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV2Router_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getZetaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getZetaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620026c2380380620026c283398181016040528101906200003791906200024e565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806200009f5750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15620000d7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1663ad5c46486040518163ffffffff1660e01b815260040160206040518083038186803b1580156200018c57600080fd5b505afa158015620001a1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001c791906200021c565b73ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250505050620002e8565b6000815190506200021681620002ce565b92915050565b600060208284031215620002355762000234620002c9565b5b6000620002458482850162000205565b91505092915050565b60008060408385031215620002685762000267620002c9565b5b6000620002788582860162000205565b92505060206200028b8582860162000205565b9150509250929050565b6000620002a282620002a9565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620002d98162000295565b8114620002e557600080fd5b50565b60805160601c60a05160601c60c05160601c612300620003c26000396000818161033401528181610598015281816108e101528181610b0f01528181610ca501528181610f0701526112700152600081816102c30152818161046a015281816106f10152818161086f01528181610ac501528181610b3101528181610bc501528181610ebd01528181610f29015281816110110152611140015260008181610254015281816105e2015281816106820152818161080001528181610c3401528181610f710152818161108001526111af01526123006000f3fe60806040526004361061004a5760003560e01c8063013b2ff81461004f57806321e093b11461007f5780632405620a146100aa57806354c49a2a146100e7578063a53fb10b14610124575b600080fd5b61006960048036038101906100649190611917565b610161565b6040516100769190611de1565b60405180910390f35b34801561008b57600080fd5b50610094610468565b6040516100a19190611c64565b60405180910390f35b3480156100b657600080fd5b506100d160048036038101906100cc9190611957565b61048c565b6040516100de9190611de1565b60405180910390f35b3480156100f357600080fd5b5061010e600480360381019061010991906119be565b610a1a565b60405161011b9190611de1565b60405180910390f35b34801561013057600080fd5b5061014b60048036038101906101469190611957565b610ddb565b6040516101589190611de1565b60405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156101c9576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000341415610204576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600267ffffffffffffffff8111156102215761022061211b565b5b60405190808252806020026020018201604052801561024f5781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110610287576102866120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f0000000000000000000000000000000000000000000000000000000000000000816001815181106102f6576102f56120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16637ff36ab53486858960c84261037f9190611f87565b6040518663ffffffff1660e01b815260040161039e9493929190611dfc565b6000604051808303818588803b1580156103b757600080fd5b505af11580156103cb573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f820116820180604052508101906103f59190611a11565b9050600081600184516104089190611fdd565b81518110610419576104186120ec565b5b602002602001015190507f87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da113482604051610454929190611e48565b60405180910390a180935050505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806104f45750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b1561052b576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610566576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105933330848673ffffffffffffffffffffffffffffffffffffffff166113a9909392919063ffffffff16565b6105de7f0000000000000000000000000000000000000000000000000000000000000000838573ffffffffffffffffffffffffffffffffffffffff166114329092919063ffffffff16565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561076357600267ffffffffffffffff81111561064f5761064e61211b565b5b60405190808252806020026020018201604052801561067d5781602001602082028036833780820191505090505b5090507f0000000000000000000000000000000000000000000000000000000000000000816000815181106106b5576106b46120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000081600181518110610724576107236120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506108dd565b600367ffffffffffffffff81111561077e5761077d61211b565b5b6040519080825280602002602001820160405280156107ac5781602001602082028036833780820191505090505b50905083816000815181106107c4576107c36120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000081600181518110610833576108326120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f0000000000000000000000000000000000000000000000000000000000000000816002815181106108a2576108a16120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166338ed17398588858b60c84261092c9190611f87565b6040518663ffffffff1660e01b815260040161094c959493929190611e71565b600060405180830381600087803b15801561096657600080fd5b505af115801561097a573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906109a39190611a11565b9050600081600184516109b69190611fdd565b815181106109c7576109c66120ec565b5b602002602001015190507f017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f868683604051610a0493929190611d08565b60405180910390a1809350505050949350505050565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610a82576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610abd576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b0a3330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166113a9909392919063ffffffff16565b610b757f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166114329092919063ffffffff16565b6000600267ffffffffffffffff811115610b9257610b9161211b565b5b604051908082528060200260200182016040528015610bc05781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110610bf857610bf76120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f000000000000000000000000000000000000000000000000000000000000000081600181518110610c6757610c666120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318cbafe58587858a60c842610cf09190611f87565b6040518663ffffffff1660e01b8152600401610d10959493929190611e71565b600060405180830381600087803b158015610d2a57600080fd5b505af1158015610d3e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610d679190611a11565b905060008160018451610d7a9190611fdd565b81518110610d8b57610d8a6120ec565b5b602002602001015190507f74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae8582604051610dc6929190611e48565b60405180910390a18093505050509392505050565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161480610e435750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610e7a576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610eb5576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f023330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166113a9909392919063ffffffff16565b610f6d7f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166114329092919063ffffffff16565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156110f257600267ffffffffffffffff811115610fde57610fdd61211b565b5b60405190808252806020026020018201604052801561100c5781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110611044576110436120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f0000000000000000000000000000000000000000000000000000000000000000816001815181106110b3576110b26120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061126c565b600367ffffffffffffffff81111561110d5761110c61211b565b5b60405190808252806020026020018201604052801561113b5781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110611173576111726120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f0000000000000000000000000000000000000000000000000000000000000000816001815181106111e2576111e16120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508381600281518110611231576112306120ec565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166338ed17398588858b60c8426112bb9190611f87565b6040518663ffffffff1660e01b81526004016112db959493929190611e71565b600060405180830381600087803b1580156112f557600080fd5b505af1158015611309573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906113329190611a11565b9050600081600184516113459190611fdd565b81518110611356576113556120ec565b5b602002602001015190507f0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b86868360405161139393929190611d08565b60405180910390a1809350505050949350505050565b61142c846323b872dd60e01b8585856040516024016113ca93929190611ca8565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611590565b50505050565b60008114806114cb575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b8152600401611479929190611c7f565b60206040518083038186803b15801561149157600080fd5b505afa1580156114a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114c99190611a87565b145b61150a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161150190611dc1565b60405180910390fd5b61158b8363095ea7b360e01b8484604051602401611529929190611cdf565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611590565b505050565b60006115f2826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166116579092919063ffffffff16565b905060008151111561165257808060200190518101906116129190611a5a565b611651576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161164890611da1565b60405180910390fd5b5b505050565b6060611666848460008561166f565b90509392505050565b6060824710156116b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116ab90611d61565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516116dd9190611c4d565b60006040518083038185875af1925050503d806000811461171a576040519150601f19603f3d011682016040523d82523d6000602084013e61171f565b606091505b50915091506117308783838761173c565b92505050949350505050565b6060831561179f5760008351141561179757611757856117b2565b611796576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161178d90611d81565b60405180910390fd5b5b8290506117aa565b6117a983836117d5565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156117e85781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161181c9190611d3f565b60405180910390fd5b600061183861183384611ef0565b611ecb565b9050808382526020820190508285602086028201111561185b5761185a61214f565b5b60005b8581101561188b57816118718882611902565b84526020840193506020830192505060018101905061185e565b5050509392505050565b6000813590506118a481612285565b92915050565b600082601f8301126118bf576118be61214a565b5b81516118cf848260208601611825565b91505092915050565b6000815190506118e78161229c565b92915050565b6000813590506118fc816122b3565b92915050565b600081519050611911816122b3565b92915050565b6000806040838503121561192e5761192d612159565b5b600061193c85828601611895565b925050602061194d858286016118ed565b9150509250929050565b6000806000806080858703121561197157611970612159565b5b600061197f87828801611895565b9450506020611990878288016118ed565b93505060406119a187828801611895565b92505060606119b2878288016118ed565b91505092959194509250565b6000806000606084860312156119d7576119d6612159565b5b60006119e586828701611895565b93505060206119f6868287016118ed565b9250506040611a07868287016118ed565b9150509250925092565b600060208284031215611a2757611a26612159565b5b600082015167ffffffffffffffff811115611a4557611a44612154565b5b611a51848285016118aa565b91505092915050565b600060208284031215611a7057611a6f612159565b5b6000611a7e848285016118d8565b91505092915050565b600060208284031215611a9d57611a9c612159565b5b6000611aab84828501611902565b91505092915050565b6000611ac08383611acc565b60208301905092915050565b611ad581612011565b82525050565b611ae481612011565b82525050565b6000611af582611f2c565b611aff8185611f5a565b9350611b0a83611f1c565b8060005b83811015611b3b578151611b228882611ab4565b9750611b2d83611f4d565b925050600181019050611b0e565b5085935050505092915050565b6000611b5382611f37565b611b5d8185611f6b565b9350611b6d818560208601612059565b80840191505092915050565b6000611b8482611f42565b611b8e8185611f76565b9350611b9e818560208601612059565b611ba78161215e565b840191505092915050565b6000611bbf602683611f76565b9150611bca8261216f565b604082019050919050565b6000611be2601d83611f76565b9150611bed826121be565b602082019050919050565b6000611c05602a83611f76565b9150611c10826121e7565b604082019050919050565b6000611c28603683611f76565b9150611c3382612236565b604082019050919050565b611c478161204f565b82525050565b6000611c598284611b48565b915081905092915050565b6000602082019050611c796000830184611adb565b92915050565b6000604082019050611c946000830185611adb565b611ca16020830184611adb565b9392505050565b6000606082019050611cbd6000830186611adb565b611cca6020830185611adb565b611cd76040830184611c3e565b949350505050565b6000604082019050611cf46000830185611adb565b611d016020830184611c3e565b9392505050565b6000606082019050611d1d6000830186611adb565b611d2a6020830185611c3e565b611d376040830184611c3e565b949350505050565b60006020820190508181036000830152611d598184611b79565b905092915050565b60006020820190508181036000830152611d7a81611bb2565b9050919050565b60006020820190508181036000830152611d9a81611bd5565b9050919050565b60006020820190508181036000830152611dba81611bf8565b9050919050565b60006020820190508181036000830152611dda81611c1b565b9050919050565b6000602082019050611df66000830184611c3e565b92915050565b6000608082019050611e116000830187611c3e565b8181036020830152611e238186611aea565b9050611e326040830185611adb565b611e3f6060830184611c3e565b95945050505050565b6000604082019050611e5d6000830185611c3e565b611e6a6020830184611c3e565b9392505050565b600060a082019050611e866000830188611c3e565b611e936020830187611c3e565b8181036040830152611ea58186611aea565b9050611eb46060830185611adb565b611ec16080830184611c3e565b9695505050505050565b6000611ed5611ee6565b9050611ee1828261208c565b919050565b6000604051905090565b600067ffffffffffffffff821115611f0b57611f0a61211b565b5b602082029050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611f928261204f565b9150611f9d8361204f565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611fd257611fd16120bd565b5b828201905092915050565b6000611fe88261204f565b9150611ff38361204f565b925082821015612006576120056120bd565b5b828203905092915050565b600061201c8261202f565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b8381101561207757808201518184015260208101905061205c565b83811115612086576000848401525b50505050565b6120958261215e565b810181811067ffffffffffffffff821117156120b4576120b361211b565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b61228e81612011565b811461229957600080fd5b50565b6122a581612023565b81146122b057600080fd5b50565b6122bc8161204f565b81146122c757600080fd5b5056fea264697066735822122019df29c139259169e5e854958a8e723d9ace2306bf0a782f32d4dfca1f62240864736f6c63430008070033",
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
