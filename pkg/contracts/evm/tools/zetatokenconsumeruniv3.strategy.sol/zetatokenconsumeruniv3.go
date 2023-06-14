// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetatokenconsumeruniv3

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

// ZetaTokenConsumerUniV3MetaData contains all meta data concerning the ZetaTokenConsumerUniV3 contract.
var ZetaTokenConsumerUniV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zetaToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV3Router_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WETH9Address_\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"zetaPoolFee_\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"tokenPoolFee_\",\"type\":\"uint24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorSendingETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getZetaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getZetaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoter\",\"outputs\":[{\"internalType\":\"contractIQuoter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPoolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV3Router\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaPoolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b50604051620025c7380380620025c783398181016040528101906200003891906200028a565b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161480620000a05750600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16145b80620000d85750600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16145b80620001105750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b1562000148576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508473ffffffffffffffffffffffffffffffffffffffff166101008173ffffffffffffffffffffffffffffffffffffffff1660601b815250508373ffffffffffffffffffffffffffffffffffffffff166101208173ffffffffffffffffffffffffffffffffffffffff1660601b815250508273ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508162ffffff1660808162ffffff1660e81b815250508062ffffff1660a08162ffffff1660e81b81525050505050505050620003a2565b6000815190506200026d816200036e565b92915050565b600081519050620002848162000388565b92915050565b60008060008060008060c08789031215620002aa57620002a962000369565b5b6000620002ba89828a016200025c565b9650506020620002cd89828a016200025c565b9550506040620002e089828a016200025c565b9450506060620002f389828a016200025c565b93505060806200030689828a0162000273565b92505060a06200031989828a0162000273565b9150509295509295509295565b600062000333826200033a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b600080fd5b620003798162000326565b81146200038557600080fd5b50565b62000393816200035a565b81146200039f57600080fd5b50565b60805160e81c60a05160e81c60c05160601c60e05160601c6101005160601c6101205160601c612114620004b360003960006111400152600081816104030152818161062301528181610761015281816108560152818161099101528181610b0301528181610ed30152611031015260008181610343015281816104f5015281816106dc01528181610947015281816109b301528181610a0701528181610e8901528181610ef50152610f480152600081816103070152818161069a01528181610a4301528181610bb00152610f8a01526000818161067901528181610d240152610fab01526000818161037f015281816106bb0152818161087a01528181610a7f0152610f6901526121146000f3fe60806040526004361061008a5760003560e01c80633cbd7005116100595780633cbd70051461015957806354c49a2a146101845780635d9dfdde146101c1578063a53fb10b146101ec578063c6bbd5a71461022957610091565b8063013b2ff81461009657806321e093b1146100c65780632405620a146100f15780632c76d7a61461012e57610091565b3661009157005b600080fd5b6100b060048036038101906100ab9190611632565b610254565b6040516100bd9190611d3c565b60405180910390f35b3480156100d257600080fd5b506100db6104f3565b6040516100e89190611b30565b60405180910390f35b3480156100fd57600080fd5b5061011860048036038101906101139190611672565b610517565b6040516101259190611d3c565b60405180910390f35b34801561013a57600080fd5b50610143610854565b6040516101509190611c26565b60405180910390f35b34801561016557600080fd5b5061016e610878565b60405161017b9190611d21565b60405180910390f35b34801561019057600080fd5b506101ab60048036038101906101a691906116d9565b61089c565b6040516101b89190611d3c565b60405180910390f35b3480156101cd57600080fd5b506101d6610d22565b6040516101e39190611d21565b60405180910390f35b3480156101f857600080fd5b50610213600480360381019061020e9190611672565b610d46565b6040516102209190611d3c565b60405180910390f35b34801561023557600080fd5b5061023e61113e565b60405161024b9190611c0b565b60405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156102bc576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60003414156102f7576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006040518061010001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000062ffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200160c8426103d19190611dc3565b8152602001348152602001848152602001600073ffffffffffffffffffffffffffffffffffffffff16815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663414bf38934846040518363ffffffff1660e01b815260040161045b9190611d05565b6020604051808303818588803b15801561047457600080fd5b505af1158015610488573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906104ad9190611759565b90507f87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da1134826040516104e0929190611d57565b60405180910390a1809250505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16148061057f5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b156105b6576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008214156105f1576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61061e3330848673ffffffffffffffffffffffffffffffffffffffff16611162909392919063ffffffff16565b6106697f0000000000000000000000000000000000000000000000000000000000000000838573ffffffffffffffffffffffffffffffffffffffff166111eb9092919063ffffffff16565b60006040518060a00160405280857f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060405160200161070f959493929190611aa5565b60405160208183030381529060405281526020018773ffffffffffffffffffffffffffffffffffffffff16815260200160c84261074c9190611dc3565b815260200184815260200186815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c04b8d59836040518263ffffffff1660e01b81526004016107b89190611ce3565b602060405180830381600087803b1580156107d257600080fd5b505af11580156107e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080a9190611759565b90507f017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f85858360405161083f93929190611bd4565b60405180910390a18092505050949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610904576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082141561093f576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61098c3330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611162909392919063ffffffff16565b6109f77f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166111eb9092919063ffffffff16565b60006040518061010001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000062ffffff1681526020013073ffffffffffffffffffffffffffffffffffffffff16815260200160c842610ad19190611dc3565b8152602001848152602001858152602001600073ffffffffffffffffffffffffffffffffffffffff16815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663414bf389836040518263ffffffff1660e01b8152600401610b5a9190611d05565b602060405180830381600087803b158015610b7457600080fd5b505af1158015610b88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bac9190611759565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d826040518263ffffffff1660e01b8152600401610c079190611d3c565b600060405180830381600087803b158015610c2157600080fd5b505af1158015610c35573d6000803e3d6000fd5b505050507f74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae8482604051610c6a929190611d57565b60405180910390a160008673ffffffffffffffffffffffffffffffffffffffff1682604051610c9890611b1b565b60006040518083038185875af1925050503d8060008114610cd5576040519150601f19603f3d011682016040523d82523d6000602084013e610cda565b606091505b5050905080610d15576040517f3794aeaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8193505050509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008060009054906101000a900460ff1615610d8e576040517f29f745a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016000806101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161480610e0f5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610e46576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610e81576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ece3330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611162909392919063ffffffff16565b610f397f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166111eb9092919063ffffffff16565b60006040518060a001604052807f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000089604051602001610fdf959493929190611aa5565b60405160208183030381529060405281526020018773ffffffffffffffffffffffffffffffffffffffff16815260200160c84261101c9190611dc3565b815260200184815260200186815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c04b8d59836040518263ffffffff1660e01b81526004016110889190611ce3565b602060405180830381600087803b1580156110a257600080fd5b505af11580156110b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110da9190611759565b90507f0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b85858360405161110f93929190611bd4565b60405180910390a1809250505060008060006101000a81548160ff021916908315150217905550949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6111e5846323b872dd60e01b85858560405160240161118393929190611b74565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611349565b50505050565b6000811480611284575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b8152600401611232929190611b4b565b60206040518083038186803b15801561124a57600080fd5b505afa15801561125e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112829190611759565b145b6112c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112ba90611cc3565b60405180910390fd5b6113448363095ea7b360e01b84846040516024016112e2929190611bab565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611349565b505050565b60006113ab826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166114109092919063ffffffff16565b905060008151111561140b57808060200190518101906113cb919061172c565b61140a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161140190611ca3565b60405180910390fd5b5b505050565b606061141f8484600085611428565b90509392505050565b60608247101561146d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161146490611c63565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516114969190611b04565b60006040518083038185875af1925050503d80600081146114d3576040519150601f19603f3d011682016040523d82523d6000602084013e6114d8565b606091505b50915091506114e9878383876114f5565b92505050949350505050565b6060831561155857600083511415611550576115108561156b565b61154f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161154690611c83565b60405180910390fd5b5b829050611563565b611562838361158e565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156115a15781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115d59190611c41565b60405180910390fd5b6000813590506115ed81612099565b92915050565b600081519050611602816120b0565b92915050565b600081359050611617816120c7565b92915050565b60008151905061162c816120c7565b92915050565b6000806040838503121561164957611648611f50565b5b6000611657858286016115de565b925050602061166885828601611608565b9150509250929050565b6000806000806080858703121561168c5761168b611f50565b5b600061169a878288016115de565b94505060206116ab87828801611608565b93505060406116bc878288016115de565b92505060606116cd87828801611608565b91505092959194509250565b6000806000606084860312156116f2576116f1611f50565b5b6000611700868287016115de565b935050602061171186828701611608565b925050604061172286828701611608565b9150509250925092565b60006020828403121561174257611741611f50565b5b6000611750848285016115f3565b91505092915050565b60006020828403121561176f5761176e611f50565b5b600061177d8482850161161d565b91505092915050565b61178f81611e19565b82525050565b61179e81611e19565b82525050565b6117b56117b082611e19565b611eeb565b82525050565b60006117c682611d80565b6117d08185611d96565b93506117e0818560208601611eb8565b6117e981611f55565b840191505092915050565b60006117ff82611d80565b6118098185611da7565b9350611819818560208601611eb8565b80840191505092915050565b61182e81611e70565b82525050565b61183d81611e82565b82525050565b600061184e82611d8b565b6118588185611db2565b9350611868818560208601611eb8565b61187181611f55565b840191505092915050565b6000611889602683611db2565b915061189482611f80565b604082019050919050565b60006118ac600083611da7565b91506118b782611fcf565b600082019050919050565b60006118cf601d83611db2565b91506118da82611fd2565b602082019050919050565b60006118f2602a83611db2565b91506118fd82611ffb565b604082019050919050565b6000611915603683611db2565b91506119208261204a565b604082019050919050565b600060a083016000830151848203600086015261194882826117bb565b915050602083015161195d6020860182611786565b5060408301516119706040860182611a87565b5060608301516119836060860182611a87565b5060808301516119966080860182611a87565b508091505092915050565b610100820160008201516119b86000850182611786565b5060208201516119cb6020850182611786565b5060408201516119de6040850182611a52565b5060608201516119f16060850182611786565b506080820151611a046080850182611a87565b5060a0820151611a1760a0850182611a87565b5060c0820151611a2a60c0850182611a87565b5060e0820151611a3d60e0850182611a43565b50505050565b611a4c81611e37565b82525050565b611a5b81611e57565b82525050565b611a6a81611e57565b82525050565b611a81611a7c82611e57565b611f0f565b82525050565b611a9081611e66565b82525050565b611a9f81611e66565b82525050565b6000611ab182886117a4565b601482019150611ac18287611a70565b600382019150611ad182866117a4565b601482019150611ae18285611a70565b600382019150611af182846117a4565b6014820191508190509695505050505050565b6000611b1082846117f4565b915081905092915050565b6000611b268261189f565b9150819050919050565b6000602082019050611b456000830184611795565b92915050565b6000604082019050611b606000830185611795565b611b6d6020830184611795565b9392505050565b6000606082019050611b896000830186611795565b611b966020830185611795565b611ba36040830184611a96565b949350505050565b6000604082019050611bc06000830185611795565b611bcd6020830184611a96565b9392505050565b6000606082019050611be96000830186611795565b611bf66020830185611a96565b611c036040830184611a96565b949350505050565b6000602082019050611c206000830184611825565b92915050565b6000602082019050611c3b6000830184611834565b92915050565b60006020820190508181036000830152611c5b8184611843565b905092915050565b60006020820190508181036000830152611c7c8161187c565b9050919050565b60006020820190508181036000830152611c9c816118c2565b9050919050565b60006020820190508181036000830152611cbc816118e5565b9050919050565b60006020820190508181036000830152611cdc81611908565b9050919050565b60006020820190508181036000830152611cfd818461192b565b905092915050565b600061010082019050611d1b60008301846119a1565b92915050565b6000602082019050611d366000830184611a61565b92915050565b6000602082019050611d516000830184611a96565b92915050565b6000604082019050611d6c6000830185611a96565b611d796020830184611a96565b9392505050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611dce82611e66565b9150611dd983611e66565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611e0e57611e0d611f21565b5b828201905092915050565b6000611e2482611e37565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b6000819050919050565b6000611e7b82611e94565b9050919050565b6000611e8d82611e94565b9050919050565b6000611e9f82611ea6565b9050919050565b6000611eb182611e37565b9050919050565b60005b83811015611ed6578082015181840152602081019050611ebb565b83811115611ee5576000848401525b50505050565b6000611ef682611efd565b9050919050565b6000611f0882611f73565b9050919050565b6000611f1a82611f66565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6000601f19601f8301169050919050565b60008160e81b9050919050565b60008160601b9050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b6120a281611e19565b81146120ad57600080fd5b50565b6120b981611e2b565b81146120c457600080fd5b50565b6120d081611e66565b81146120db57600080fd5b5056fea2646970667358221220b86064610cc7341a7ebda7eb469b35d8c06d4a196518062087b0933eb1902a1b64736f6c63430008070033",
}

// ZetaTokenConsumerUniV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaTokenConsumerUniV3MetaData.ABI instead.
var ZetaTokenConsumerUniV3ABI = ZetaTokenConsumerUniV3MetaData.ABI

// ZetaTokenConsumerUniV3Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaTokenConsumerUniV3MetaData.Bin instead.
var ZetaTokenConsumerUniV3Bin = ZetaTokenConsumerUniV3MetaData.Bin

// DeployZetaTokenConsumerUniV3 deploys a new Ethereum contract, binding an instance of ZetaTokenConsumerUniV3 to it.
func DeployZetaTokenConsumerUniV3(auth *bind.TransactOpts, backend bind.ContractBackend, zetaToken_ common.Address, uniswapV3Router_ common.Address, quoter_ common.Address, WETH9Address_ common.Address, zetaPoolFee_ *big.Int, tokenPoolFee_ *big.Int) (common.Address, *types.Transaction, *ZetaTokenConsumerUniV3, error) {
	parsed, err := ZetaTokenConsumerUniV3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaTokenConsumerUniV3Bin), backend, zetaToken_, uniswapV3Router_, quoter_, WETH9Address_, zetaPoolFee_, tokenPoolFee_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaTokenConsumerUniV3{ZetaTokenConsumerUniV3Caller: ZetaTokenConsumerUniV3Caller{contract: contract}, ZetaTokenConsumerUniV3Transactor: ZetaTokenConsumerUniV3Transactor{contract: contract}, ZetaTokenConsumerUniV3Filterer: ZetaTokenConsumerUniV3Filterer{contract: contract}}, nil
}

// ZetaTokenConsumerUniV3 is an auto generated Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV3 struct {
	ZetaTokenConsumerUniV3Caller     // Read-only binding to the contract
	ZetaTokenConsumerUniV3Transactor // Write-only binding to the contract
	ZetaTokenConsumerUniV3Filterer   // Log filterer for contract events
}

// ZetaTokenConsumerUniV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerUniV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerUniV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaTokenConsumerUniV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerUniV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaTokenConsumerUniV3Session struct {
	Contract     *ZetaTokenConsumerUniV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerUniV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaTokenConsumerUniV3CallerSession struct {
	Contract *ZetaTokenConsumerUniV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ZetaTokenConsumerUniV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaTokenConsumerUniV3TransactorSession struct {
	Contract     *ZetaTokenConsumerUniV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerUniV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV3Raw struct {
	Contract *ZetaTokenConsumerUniV3 // Generic contract binding to access the raw methods on
}

// ZetaTokenConsumerUniV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV3CallerRaw struct {
	Contract *ZetaTokenConsumerUniV3Caller // Generic read-only contract binding to access the raw methods on
}

// ZetaTokenConsumerUniV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV3TransactorRaw struct {
	Contract *ZetaTokenConsumerUniV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaTokenConsumerUniV3 creates a new instance of ZetaTokenConsumerUniV3, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV3(address common.Address, backend bind.ContractBackend) (*ZetaTokenConsumerUniV3, error) {
	contract, err := bindZetaTokenConsumerUniV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3{ZetaTokenConsumerUniV3Caller: ZetaTokenConsumerUniV3Caller{contract: contract}, ZetaTokenConsumerUniV3Transactor: ZetaTokenConsumerUniV3Transactor{contract: contract}, ZetaTokenConsumerUniV3Filterer: ZetaTokenConsumerUniV3Filterer{contract: contract}}, nil
}

// NewZetaTokenConsumerUniV3Caller creates a new read-only instance of ZetaTokenConsumerUniV3, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV3Caller(address common.Address, caller bind.ContractCaller) (*ZetaTokenConsumerUniV3Caller, error) {
	contract, err := bindZetaTokenConsumerUniV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3Caller{contract: contract}, nil
}

// NewZetaTokenConsumerUniV3Transactor creates a new write-only instance of ZetaTokenConsumerUniV3, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV3Transactor(address common.Address, transactor bind.ContractTransactor) (*ZetaTokenConsumerUniV3Transactor, error) {
	contract, err := bindZetaTokenConsumerUniV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3Transactor{contract: contract}, nil
}

// NewZetaTokenConsumerUniV3Filterer creates a new log filterer instance of ZetaTokenConsumerUniV3, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV3Filterer(address common.Address, filterer bind.ContractFilterer) (*ZetaTokenConsumerUniV3Filterer, error) {
	contract, err := bindZetaTokenConsumerUniV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3Filterer{contract: contract}, nil
}

// bindZetaTokenConsumerUniV3 binds a generic wrapper to an already deployed contract.
func bindZetaTokenConsumerUniV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaTokenConsumerUniV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerUniV3.Contract.ZetaTokenConsumerUniV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.ZetaTokenConsumerUniV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.ZetaTokenConsumerUniV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerUniV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.contract.Transact(opts, method, params...)
}

// Quoter is a free data retrieval call binding the contract method 0xc6bbd5a7.
//
// Solidity: function quoter() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Caller) Quoter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV3.contract.Call(opts, &out, "quoter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Quoter is a free data retrieval call binding the contract method 0xc6bbd5a7.
//
// Solidity: function quoter() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) Quoter() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.Quoter(&_ZetaTokenConsumerUniV3.CallOpts)
}

// Quoter is a free data retrieval call binding the contract method 0xc6bbd5a7.
//
// Solidity: function quoter() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerSession) Quoter() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.Quoter(&_ZetaTokenConsumerUniV3.CallOpts)
}

// TokenPoolFee is a free data retrieval call binding the contract method 0x5d9dfdde.
//
// Solidity: function tokenPoolFee() view returns(uint24)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Caller) TokenPoolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV3.contract.Call(opts, &out, "tokenPoolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenPoolFee is a free data retrieval call binding the contract method 0x5d9dfdde.
//
// Solidity: function tokenPoolFee() view returns(uint24)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) TokenPoolFee() (*big.Int, error) {
	return _ZetaTokenConsumerUniV3.Contract.TokenPoolFee(&_ZetaTokenConsumerUniV3.CallOpts)
}

// TokenPoolFee is a free data retrieval call binding the contract method 0x5d9dfdde.
//
// Solidity: function tokenPoolFee() view returns(uint24)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerSession) TokenPoolFee() (*big.Int, error) {
	return _ZetaTokenConsumerUniV3.Contract.TokenPoolFee(&_ZetaTokenConsumerUniV3.CallOpts)
}

// UniswapV3Router is a free data retrieval call binding the contract method 0x2c76d7a6.
//
// Solidity: function uniswapV3Router() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Caller) UniswapV3Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV3.contract.Call(opts, &out, "uniswapV3Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV3Router is a free data retrieval call binding the contract method 0x2c76d7a6.
//
// Solidity: function uniswapV3Router() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) UniswapV3Router() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.UniswapV3Router(&_ZetaTokenConsumerUniV3.CallOpts)
}

// UniswapV3Router is a free data retrieval call binding the contract method 0x2c76d7a6.
//
// Solidity: function uniswapV3Router() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerSession) UniswapV3Router() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.UniswapV3Router(&_ZetaTokenConsumerUniV3.CallOpts)
}

// ZetaPoolFee is a free data retrieval call binding the contract method 0x3cbd7005.
//
// Solidity: function zetaPoolFee() view returns(uint24)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Caller) ZetaPoolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV3.contract.Call(opts, &out, "zetaPoolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZetaPoolFee is a free data retrieval call binding the contract method 0x3cbd7005.
//
// Solidity: function zetaPoolFee() view returns(uint24)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) ZetaPoolFee() (*big.Int, error) {
	return _ZetaTokenConsumerUniV3.Contract.ZetaPoolFee(&_ZetaTokenConsumerUniV3.CallOpts)
}

// ZetaPoolFee is a free data retrieval call binding the contract method 0x3cbd7005.
//
// Solidity: function zetaPoolFee() view returns(uint24)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerSession) ZetaPoolFee() (*big.Int, error) {
	return _ZetaTokenConsumerUniV3.Contract.ZetaPoolFee(&_ZetaTokenConsumerUniV3.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Caller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV3.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) ZetaToken() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.ZetaToken(&_ZetaTokenConsumerUniV3.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerSession) ZetaToken() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.ZetaToken(&_ZetaTokenConsumerUniV3.CallOpts)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Transactor) GetEthFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.contract.Transact(opts, "getEthFromZeta", destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetEthFromZeta(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetEthFromZeta(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Transactor) GetTokenFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.contract.Transact(opts, "getTokenFromZeta", destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetTokenFromZeta(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetTokenFromZeta(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Transactor) GetZetaFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.contract.Transact(opts, "getZetaFromEth", destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetZetaFromEth(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetZetaFromEth(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Transactor) GetZetaFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.contract.Transact(opts, "getZetaFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetZetaFromToken(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetZetaFromToken(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) Receive() (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.Receive(&_ZetaTokenConsumerUniV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorSession) Receive() (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.Receive(&_ZetaTokenConsumerUniV3.TransactOpts)
}

// ZetaTokenConsumerUniV3EthExchangedForZetaIterator is returned from FilterEthExchangedForZeta and is used to iterate over the raw logs and unpacked data for EthExchangedForZeta events raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3EthExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerUniV3EthExchangedForZeta // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerUniV3EthExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerUniV3EthExchangedForZeta)
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
		it.Event = new(ZetaTokenConsumerUniV3EthExchangedForZeta)
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
func (it *ZetaTokenConsumerUniV3EthExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerUniV3EthExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerUniV3EthExchangedForZeta represents a EthExchangedForZeta event raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3EthExchangedForZeta struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForZeta is a free log retrieval operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) FilterEthExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerUniV3EthExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.FilterLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3EthExchangedForZetaIterator{contract: _ZetaTokenConsumerUniV3.contract, event: "EthExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForZeta is a free log subscription operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) WatchEthExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerUniV3EthExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.WatchLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerUniV3EthExchangedForZeta)
				if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
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
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) ParseEthExchangedForZeta(log types.Log) (*ZetaTokenConsumerUniV3EthExchangedForZeta, error) {
	event := new(ZetaTokenConsumerUniV3EthExchangedForZeta)
	if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerUniV3TokenExchangedForZetaIterator is returned from FilterTokenExchangedForZeta and is used to iterate over the raw logs and unpacked data for TokenExchangedForZeta events raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3TokenExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerUniV3TokenExchangedForZeta // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerUniV3TokenExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerUniV3TokenExchangedForZeta)
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
		it.Event = new(ZetaTokenConsumerUniV3TokenExchangedForZeta)
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
func (it *ZetaTokenConsumerUniV3TokenExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerUniV3TokenExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerUniV3TokenExchangedForZeta represents a TokenExchangedForZeta event raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3TokenExchangedForZeta struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForZeta is a free log retrieval operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) FilterTokenExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerUniV3TokenExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.FilterLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3TokenExchangedForZetaIterator{contract: _ZetaTokenConsumerUniV3.contract, event: "TokenExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForZeta is a free log subscription operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) WatchTokenExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerUniV3TokenExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.WatchLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerUniV3TokenExchangedForZeta)
				if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
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
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) ParseTokenExchangedForZeta(log types.Log) (*ZetaTokenConsumerUniV3TokenExchangedForZeta, error) {
	event := new(ZetaTokenConsumerUniV3TokenExchangedForZeta)
	if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerUniV3ZetaExchangedForEthIterator is returned from FilterZetaExchangedForEth and is used to iterate over the raw logs and unpacked data for ZetaExchangedForEth events raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3ZetaExchangedForEthIterator struct {
	Event *ZetaTokenConsumerUniV3ZetaExchangedForEth // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerUniV3ZetaExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerUniV3ZetaExchangedForEth)
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
		it.Event = new(ZetaTokenConsumerUniV3ZetaExchangedForEth)
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
func (it *ZetaTokenConsumerUniV3ZetaExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerUniV3ZetaExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerUniV3ZetaExchangedForEth represents a ZetaExchangedForEth event raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3ZetaExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForEth is a free log retrieval operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) FilterZetaExchangedForEth(opts *bind.FilterOpts) (*ZetaTokenConsumerUniV3ZetaExchangedForEthIterator, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.FilterLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3ZetaExchangedForEthIterator{contract: _ZetaTokenConsumerUniV3.contract, event: "ZetaExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForEth is a free log subscription operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) WatchZetaExchangedForEth(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerUniV3ZetaExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.WatchLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerUniV3ZetaExchangedForEth)
				if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
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
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) ParseZetaExchangedForEth(log types.Log) (*ZetaTokenConsumerUniV3ZetaExchangedForEth, error) {
	event := new(ZetaTokenConsumerUniV3ZetaExchangedForEth)
	if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator is returned from FilterZetaExchangedForToken and is used to iterate over the raw logs and unpacked data for ZetaExchangedForToken events raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator struct {
	Event *ZetaTokenConsumerUniV3ZetaExchangedForToken // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerUniV3ZetaExchangedForToken)
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
		it.Event = new(ZetaTokenConsumerUniV3ZetaExchangedForToken)
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
func (it *ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerUniV3ZetaExchangedForToken represents a ZetaExchangedForToken event raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3ZetaExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForToken is a free log retrieval operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) FilterZetaExchangedForToken(opts *bind.FilterOpts) (*ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.FilterLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator{contract: _ZetaTokenConsumerUniV3.contract, event: "ZetaExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForToken is a free log subscription operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) WatchZetaExchangedForToken(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerUniV3ZetaExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.WatchLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerUniV3ZetaExchangedForToken)
				if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
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
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) ParseZetaExchangedForToken(log types.Log) (*ZetaTokenConsumerUniV3ZetaExchangedForToken, error) {
	event := new(ZetaTokenConsumerUniV3ZetaExchangedForToken)
	if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
