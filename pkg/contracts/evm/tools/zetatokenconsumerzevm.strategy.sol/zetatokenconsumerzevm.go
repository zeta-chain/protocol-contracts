// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetatokenconsumerzevm

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

// ZetaTokenConsumerZEVMMetaData contains all meta data concerning the ZetaTokenConsumerZEVM contract.
var ZetaTokenConsumerZEVMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"WETH9Address_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV2Router_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorSendingETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZeta\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWZETAAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutputCantBeZeta\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH9Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getZetaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getZetaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasZetaLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200223d3803806200223d833981810160405281019062000037919062000164565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806200009f5750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15620000d7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250505050620001fe565b6000815190506200015e81620001e4565b92915050565b600080604083850312156200017e576200017d620001df565b5b60006200018e858286016200014d565b9250506020620001a1858286016200014d565b9150509250929050565b6000620001b882620001bf565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620001ef81620001ab565b8114620001fb57600080fd5b50565b60805160601c60a05160601c611fb26200028b6000396000818161051e015281816106fa01528181610ce90152610e5f0152600081816060015281816103060152818161038c015281816105660152818161068901528181610912015281816109c201528181610c1301528181610c9f01528181610d0b01528181610d9f0152610f9a0152611fb26000f3fe6080604052600436106100595760003560e01c8063013b2ff8146100ea5780632405620a1461011a57806354c49a2a1461015757806380801f8414610194578063a53fb10b146101bf578063c469cf14146101fc576100e5565b366100e5577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100e3576040517f290ee5a500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b61010460048036038101906100ff91906115b0565b610227565b6040516101119190611adc565b60405180910390f35b34801561012657600080fd5b50610141600480360381019061013c91906115f0565b610412565b60405161014e9190611adc565b60405180910390f35b34801561016357600080fd5b5061017e60048036038101906101799190611657565b610833565b60405161018b9190611adc565b60405180910390f35b3480156101a057600080fd5b506101a9610b32565b6040516101b69190611a1f565b60405180910390f35b3480156101cb57600080fd5b506101e660048036038101906101e191906115f0565b610b37565b6040516101f39190611adc565b60405180910390f35b34801561020857600080fd5b50610211610f98565b60405161021e9190611944565b60405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561028f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60003414156102ca576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81341015610304576040517fe2f844a000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561036c57600080fd5b505af1158015610380573d6000803e3d6000fd5b50505050506103d083347f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610fbc9092919063ffffffff16565b7f87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da113434604051610401929190611af7565b60405180910390a134905092915050565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16148061047a5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b156104b1576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008214156104ec576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105193330848673ffffffffffffffffffffffffffffffffffffffff16611042909392919063ffffffff16565b6105647f0000000000000000000000000000000000000000000000000000000000000000838573ffffffffffffffffffffffffffffffffffffffff166110cb9092919063ffffffff16565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156105ea576040517f6edfe50500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600267ffffffffffffffff81111561060757610606611dca565b5b6040519080825280602002602001820160405280156106355781602001602082028036833780820191505090505b509050838160008151811061064d5761064c611d9b565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f0000000000000000000000000000000000000000000000000000000000000000816001815181106106bc576106bb611d9b565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166338ed17398588858b60c8426107459190611c36565b6040518663ffffffff1660e01b8152600401610765959493929190611b20565b600060405180830381600087803b15801561077f57600080fd5b505af1158015610793573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906107bc91906116aa565b9050600081600184516107cf9190611c8c565b815181106107e0576107df611d9b565b5b602002602001015190507f017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f86868360405161081d939291906119e8565b60405180910390a1809350505050949350505050565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561089b576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008214156108d6576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82821015610910576040517fe2f844a000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161096d93929190611988565b602060405180830381600087803b15801561098757600080fd5b505af115801561099b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109bf91906116f3565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d836040518263ffffffff1660e01b8152600401610a199190611adc565b600060405180830381600087803b158015610a3357600080fd5b505af1158015610a47573d6000803e3d6000fd5b505050507f74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae8283604051610a7c929190611af7565b60405180910390a160008473ffffffffffffffffffffffffffffffffffffffff1683604051610aaa9061192f565b60006040518083038185875af1925050503d8060008114610ae7576040519150601f19603f3d011682016040523d82523d6000602084013e610aec565b606091505b5050905080610b27576040517f3794aeaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b829150509392505050565b600090565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161480610b9f5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610bd6576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610c11576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610c97576040517f8c51927900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ce43330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611042909392919063ffffffff16565b610d4f7f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166110cb9092919063ffffffff16565b6000600267ffffffffffffffff811115610d6c57610d6b611dca565b5b604051908082528060200260200182016040528015610d9a5781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110610dd257610dd1611d9b565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508381600181518110610e2157610e20611d9b565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166338ed17398588858b60c842610eaa9190611c36565b6040518663ffffffff1660e01b8152600401610eca959493929190611b20565b600060405180830381600087803b158015610ee457600080fd5b505af1158015610ef8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610f2191906116aa565b905060008160018451610f349190611c8c565b81518110610f4557610f44611d9b565b5b602002602001015190507f0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b868683604051610f82939291906119e8565b60405180910390a1809350505050949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b61103d8363a9059cbb60e01b8484604051602401610fdb9291906119bf565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611229565b505050565b6110c5846323b872dd60e01b85858560405160240161106393929190611988565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611229565b50505050565b6000811480611164575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b815260040161111292919061195f565b60206040518083038186803b15801561112a57600080fd5b505afa15801561113e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111629190611720565b145b6111a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161119a90611abc565b60405180910390fd5b6112248363095ea7b360e01b84846040516024016111c29291906119bf565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611229565b505050565b600061128b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166112f09092919063ffffffff16565b90506000815111156112eb57808060200190518101906112ab91906116f3565b6112ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112e190611a9c565b60405180910390fd5b5b505050565b60606112ff8484600085611308565b90509392505050565b60608247101561134d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161134490611a5c565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516113769190611918565b60006040518083038185875af1925050503d80600081146113b3576040519150601f19603f3d011682016040523d82523d6000602084013e6113b8565b606091505b50915091506113c9878383876113d5565b92505050949350505050565b6060831561143857600083511415611430576113f08561144b565b61142f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161142690611a7c565b60405180910390fd5b5b829050611443565b611442838361146e565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156114815781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114b59190611a3a565b60405180910390fd5b60006114d16114cc84611b9f565b611b7a565b905080838252602082019050828560208602820111156114f4576114f3611dfe565b5b60005b85811015611524578161150a888261159b565b8452602084019350602083019250506001810190506114f7565b5050509392505050565b60008135905061153d81611f37565b92915050565b600082601f83011261155857611557611df9565b5b81516115688482602086016114be565b91505092915050565b60008151905061158081611f4e565b92915050565b60008135905061159581611f65565b92915050565b6000815190506115aa81611f65565b92915050565b600080604083850312156115c7576115c6611e08565b5b60006115d58582860161152e565b92505060206115e685828601611586565b9150509250929050565b6000806000806080858703121561160a57611609611e08565b5b60006116188782880161152e565b945050602061162987828801611586565b935050604061163a8782880161152e565b925050606061164b87828801611586565b91505092959194509250565b6000806000606084860312156116705761166f611e08565b5b600061167e8682870161152e565b935050602061168f86828701611586565b92505060406116a086828701611586565b9150509250925092565b6000602082840312156116c0576116bf611e08565b5b600082015167ffffffffffffffff8111156116de576116dd611e03565b5b6116ea84828501611543565b91505092915050565b60006020828403121561170957611708611e08565b5b600061171784828501611571565b91505092915050565b60006020828403121561173657611735611e08565b5b60006117448482850161159b565b91505092915050565b60006117598383611765565b60208301905092915050565b61176e81611cc0565b82525050565b61177d81611cc0565b82525050565b600061178e82611bdb565b6117988185611c09565b93506117a383611bcb565b8060005b838110156117d45781516117bb888261174d565b97506117c683611bfc565b9250506001810190506117a7565b5085935050505092915050565b6117ea81611cd2565b82525050565b60006117fb82611be6565b6118058185611c1a565b9350611815818560208601611d08565b80840191505092915050565b600061182c82611bf1565b6118368185611c25565b9350611846818560208601611d08565b61184f81611e0d565b840191505092915050565b6000611867602683611c25565b915061187282611e1e565b604082019050919050565b600061188a600083611c1a565b915061189582611e6d565b600082019050919050565b60006118ad601d83611c25565b91506118b882611e70565b602082019050919050565b60006118d0602a83611c25565b91506118db82611e99565b604082019050919050565b60006118f3603683611c25565b91506118fe82611ee8565b604082019050919050565b61191281611cfe565b82525050565b600061192482846117f0565b915081905092915050565b600061193a8261187d565b9150819050919050565b60006020820190506119596000830184611774565b92915050565b60006040820190506119746000830185611774565b6119816020830184611774565b9392505050565b600060608201905061199d6000830186611774565b6119aa6020830185611774565b6119b76040830184611909565b949350505050565b60006040820190506119d46000830185611774565b6119e16020830184611909565b9392505050565b60006060820190506119fd6000830186611774565b611a0a6020830185611909565b611a176040830184611909565b949350505050565b6000602082019050611a3460008301846117e1565b92915050565b60006020820190508181036000830152611a548184611821565b905092915050565b60006020820190508181036000830152611a758161185a565b9050919050565b60006020820190508181036000830152611a95816118a0565b9050919050565b60006020820190508181036000830152611ab5816118c3565b9050919050565b60006020820190508181036000830152611ad5816118e6565b9050919050565b6000602082019050611af16000830184611909565b92915050565b6000604082019050611b0c6000830185611909565b611b196020830184611909565b9392505050565b600060a082019050611b356000830188611909565b611b426020830187611909565b8181036040830152611b548186611783565b9050611b636060830185611774565b611b706080830184611909565b9695505050505050565b6000611b84611b95565b9050611b908282611d3b565b919050565b6000604051905090565b600067ffffffffffffffff821115611bba57611bb9611dca565b5b602082029050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611c4182611cfe565b9150611c4c83611cfe565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611c8157611c80611d6c565b5b828201905092915050565b6000611c9782611cfe565b9150611ca283611cfe565b925082821015611cb557611cb4611d6c565b5b828203905092915050565b6000611ccb82611cde565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b83811015611d26578082015181840152602081019050611d0b565b83811115611d35576000848401525b50505050565b611d4482611e0d565b810181811067ffffffffffffffff82111715611d6357611d62611dca565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b611f4081611cc0565b8114611f4b57600080fd5b50565b611f5781611cd2565b8114611f6257600080fd5b50565b611f6e81611cfe565b8114611f7957600080fd5b5056fea2646970667358221220cd682461082c00bcc9510445b7e8accdf628761ccff0b4240c8af552d144358664736f6c63430008070033",
}

// ZetaTokenConsumerZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaTokenConsumerZEVMMetaData.ABI instead.
var ZetaTokenConsumerZEVMABI = ZetaTokenConsumerZEVMMetaData.ABI

// ZetaTokenConsumerZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaTokenConsumerZEVMMetaData.Bin instead.
var ZetaTokenConsumerZEVMBin = ZetaTokenConsumerZEVMMetaData.Bin

// DeployZetaTokenConsumerZEVM deploys a new Ethereum contract, binding an instance of ZetaTokenConsumerZEVM to it.
func DeployZetaTokenConsumerZEVM(auth *bind.TransactOpts, backend bind.ContractBackend, WETH9Address_ common.Address, uniswapV2Router_ common.Address) (common.Address, *types.Transaction, *ZetaTokenConsumerZEVM, error) {
	parsed, err := ZetaTokenConsumerZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaTokenConsumerZEVMBin), backend, WETH9Address_, uniswapV2Router_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaTokenConsumerZEVM{ZetaTokenConsumerZEVMCaller: ZetaTokenConsumerZEVMCaller{contract: contract}, ZetaTokenConsumerZEVMTransactor: ZetaTokenConsumerZEVMTransactor{contract: contract}, ZetaTokenConsumerZEVMFilterer: ZetaTokenConsumerZEVMFilterer{contract: contract}}, nil
}

// ZetaTokenConsumerZEVM is an auto generated Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVM struct {
	ZetaTokenConsumerZEVMCaller     // Read-only binding to the contract
	ZetaTokenConsumerZEVMTransactor // Write-only binding to the contract
	ZetaTokenConsumerZEVMFilterer   // Log filterer for contract events
}

// ZetaTokenConsumerZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaTokenConsumerZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaTokenConsumerZEVMSession struct {
	Contract     *ZetaTokenConsumerZEVM // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaTokenConsumerZEVMCallerSession struct {
	Contract *ZetaTokenConsumerZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ZetaTokenConsumerZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaTokenConsumerZEVMTransactorSession struct {
	Contract     *ZetaTokenConsumerZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMRaw struct {
	Contract *ZetaTokenConsumerZEVM // Generic contract binding to access the raw methods on
}

// ZetaTokenConsumerZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMCallerRaw struct {
	Contract *ZetaTokenConsumerZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaTokenConsumerZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMTransactorRaw struct {
	Contract *ZetaTokenConsumerZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaTokenConsumerZEVM creates a new instance of ZetaTokenConsumerZEVM, bound to a specific deployed contract.
func NewZetaTokenConsumerZEVM(address common.Address, backend bind.ContractBackend) (*ZetaTokenConsumerZEVM, error) {
	contract, err := bindZetaTokenConsumerZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVM{ZetaTokenConsumerZEVMCaller: ZetaTokenConsumerZEVMCaller{contract: contract}, ZetaTokenConsumerZEVMTransactor: ZetaTokenConsumerZEVMTransactor{contract: contract}, ZetaTokenConsumerZEVMFilterer: ZetaTokenConsumerZEVMFilterer{contract: contract}}, nil
}

// NewZetaTokenConsumerZEVMCaller creates a new read-only instance of ZetaTokenConsumerZEVM, bound to a specific deployed contract.
func NewZetaTokenConsumerZEVMCaller(address common.Address, caller bind.ContractCaller) (*ZetaTokenConsumerZEVMCaller, error) {
	contract, err := bindZetaTokenConsumerZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMCaller{contract: contract}, nil
}

// NewZetaTokenConsumerZEVMTransactor creates a new write-only instance of ZetaTokenConsumerZEVM, bound to a specific deployed contract.
func NewZetaTokenConsumerZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaTokenConsumerZEVMTransactor, error) {
	contract, err := bindZetaTokenConsumerZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMTransactor{contract: contract}, nil
}

// NewZetaTokenConsumerZEVMFilterer creates a new log filterer instance of ZetaTokenConsumerZEVM, bound to a specific deployed contract.
func NewZetaTokenConsumerZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaTokenConsumerZEVMFilterer, error) {
	contract, err := bindZetaTokenConsumerZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMFilterer{contract: contract}, nil
}

// bindZetaTokenConsumerZEVM binds a generic wrapper to an already deployed contract.
func bindZetaTokenConsumerZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaTokenConsumerZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerZEVM.Contract.ZetaTokenConsumerZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.ZetaTokenConsumerZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.ZetaTokenConsumerZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.contract.Transact(opts, method, params...)
}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMCaller) WETH9Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerZEVM.contract.Call(opts, &out, "WETH9Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) WETH9Address() (common.Address, error) {
	return _ZetaTokenConsumerZEVM.Contract.WETH9Address(&_ZetaTokenConsumerZEVM.CallOpts)
}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMCallerSession) WETH9Address() (common.Address, error) {
	return _ZetaTokenConsumerZEVM.Contract.WETH9Address(&_ZetaTokenConsumerZEVM.CallOpts)
}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMCaller) HasZetaLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaTokenConsumerZEVM.contract.Call(opts, &out, "hasZetaLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) HasZetaLiquidity() (bool, error) {
	return _ZetaTokenConsumerZEVM.Contract.HasZetaLiquidity(&_ZetaTokenConsumerZEVM.CallOpts)
}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMCallerSession) HasZetaLiquidity() (bool, error) {
	return _ZetaTokenConsumerZEVM.Contract.HasZetaLiquidity(&_ZetaTokenConsumerZEVM.CallOpts)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactor) GetEthFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.contract.Transact(opts, "getEthFromZeta", destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetEthFromZeta(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetEthFromZeta(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactor) GetTokenFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.contract.Transact(opts, "getTokenFromZeta", destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetTokenFromZeta(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetTokenFromZeta(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactor) GetZetaFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.contract.Transact(opts, "getZetaFromEth", destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetZetaFromEth(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetZetaFromEth(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactor) GetZetaFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.contract.Transact(opts, "getZetaFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetZetaFromToken(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetZetaFromToken(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) Receive() (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.Receive(&_ZetaTokenConsumerZEVM.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorSession) Receive() (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.Receive(&_ZetaTokenConsumerZEVM.TransactOpts)
}

// ZetaTokenConsumerZEVMEthExchangedForZetaIterator is returned from FilterEthExchangedForZeta and is used to iterate over the raw logs and unpacked data for EthExchangedForZeta events raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMEthExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerZEVMEthExchangedForZeta // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerZEVMEthExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerZEVMEthExchangedForZeta)
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
		it.Event = new(ZetaTokenConsumerZEVMEthExchangedForZeta)
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
func (it *ZetaTokenConsumerZEVMEthExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerZEVMEthExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerZEVMEthExchangedForZeta represents a EthExchangedForZeta event raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMEthExchangedForZeta struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForZeta is a free log retrieval operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) FilterEthExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerZEVMEthExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.FilterLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMEthExchangedForZetaIterator{contract: _ZetaTokenConsumerZEVM.contract, event: "EthExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForZeta is a free log subscription operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) WatchEthExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerZEVMEthExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.WatchLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerZEVMEthExchangedForZeta)
				if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
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
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) ParseEthExchangedForZeta(log types.Log) (*ZetaTokenConsumerZEVMEthExchangedForZeta, error) {
	event := new(ZetaTokenConsumerZEVMEthExchangedForZeta)
	if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerZEVMTokenExchangedForZetaIterator is returned from FilterTokenExchangedForZeta and is used to iterate over the raw logs and unpacked data for TokenExchangedForZeta events raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMTokenExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerZEVMTokenExchangedForZeta // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerZEVMTokenExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerZEVMTokenExchangedForZeta)
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
		it.Event = new(ZetaTokenConsumerZEVMTokenExchangedForZeta)
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
func (it *ZetaTokenConsumerZEVMTokenExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerZEVMTokenExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerZEVMTokenExchangedForZeta represents a TokenExchangedForZeta event raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMTokenExchangedForZeta struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForZeta is a free log retrieval operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) FilterTokenExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerZEVMTokenExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.FilterLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMTokenExchangedForZetaIterator{contract: _ZetaTokenConsumerZEVM.contract, event: "TokenExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForZeta is a free log subscription operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) WatchTokenExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerZEVMTokenExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.WatchLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerZEVMTokenExchangedForZeta)
				if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
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
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) ParseTokenExchangedForZeta(log types.Log) (*ZetaTokenConsumerZEVMTokenExchangedForZeta, error) {
	event := new(ZetaTokenConsumerZEVMTokenExchangedForZeta)
	if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerZEVMZetaExchangedForEthIterator is returned from FilterZetaExchangedForEth and is used to iterate over the raw logs and unpacked data for ZetaExchangedForEth events raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMZetaExchangedForEthIterator struct {
	Event *ZetaTokenConsumerZEVMZetaExchangedForEth // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerZEVMZetaExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerZEVMZetaExchangedForEth)
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
		it.Event = new(ZetaTokenConsumerZEVMZetaExchangedForEth)
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
func (it *ZetaTokenConsumerZEVMZetaExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerZEVMZetaExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerZEVMZetaExchangedForEth represents a ZetaExchangedForEth event raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMZetaExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForEth is a free log retrieval operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) FilterZetaExchangedForEth(opts *bind.FilterOpts) (*ZetaTokenConsumerZEVMZetaExchangedForEthIterator, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.FilterLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMZetaExchangedForEthIterator{contract: _ZetaTokenConsumerZEVM.contract, event: "ZetaExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForEth is a free log subscription operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) WatchZetaExchangedForEth(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerZEVMZetaExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.WatchLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerZEVMZetaExchangedForEth)
				if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
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
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) ParseZetaExchangedForEth(log types.Log) (*ZetaTokenConsumerZEVMZetaExchangedForEth, error) {
	event := new(ZetaTokenConsumerZEVMZetaExchangedForEth)
	if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerZEVMZetaExchangedForTokenIterator is returned from FilterZetaExchangedForToken and is used to iterate over the raw logs and unpacked data for ZetaExchangedForToken events raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMZetaExchangedForTokenIterator struct {
	Event *ZetaTokenConsumerZEVMZetaExchangedForToken // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerZEVMZetaExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerZEVMZetaExchangedForToken)
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
		it.Event = new(ZetaTokenConsumerZEVMZetaExchangedForToken)
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
func (it *ZetaTokenConsumerZEVMZetaExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerZEVMZetaExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerZEVMZetaExchangedForToken represents a ZetaExchangedForToken event raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMZetaExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForToken is a free log retrieval operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) FilterZetaExchangedForToken(opts *bind.FilterOpts) (*ZetaTokenConsumerZEVMZetaExchangedForTokenIterator, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.FilterLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMZetaExchangedForTokenIterator{contract: _ZetaTokenConsumerZEVM.contract, event: "ZetaExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForToken is a free log subscription operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) WatchZetaExchangedForToken(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerZEVMZetaExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.WatchLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerZEVMZetaExchangedForToken)
				if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
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
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) ParseZetaExchangedForToken(log types.Log) (*ZetaTokenConsumerZEVMZetaExchangedForToken, error) {
	event := new(ZetaTokenConsumerZEVMZetaExchangedForToken)
	if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
