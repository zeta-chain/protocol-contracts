// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayv2

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

// GatewayV2MetaData contains all meta data concerning the GatewayV2 contract.
var GatewayV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ExecutedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ExecutedWithERC20V2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"custody\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeWithERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_custody\",\"type\":\"address\"}],\"name\":\"setCustody\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b8152503480156200004757600080fd5b50620000586200005e60201b60201c565b62000208565b600060019054906101000a900460ff1615620000b1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a8906200015c565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff1614620001225760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200011991906200017e565b60405180910390a15b565b6000620001336027836200019b565b91506200014082620001b9565b604082019050919050565b6200015681620001ac565b82525050565b60006020820190508181036000830152620001778162000124565b9050919050565b60006020820190506200019560008301846200014b565b92915050565b600082825260208201905092915050565b600060ff82169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60805160601c6121c062000243600039600081816102c4015281816103530152818161044d015281816104dc015261087801526121c06000f3fe60806040526004361061009c5760003560e01c8063715018a611610064578063715018a61461017e5780638129fc1c146101955780638da5cb5b146101ac578063ae7a3a6f146101d7578063dda79b7514610200578063f2fde38b1461022b5761009c565b80631cff79cd146100a15780633659cfe6146100d15780634f1ef286146100fa5780635131ab591461011657806352d1902d14610153575b600080fd5b6100bb60048036038101906100b69190611554565b610254565b6040516100c89190611a10565b60405180910390f35b3480156100dd57600080fd5b506100f860048036038101906100f3919061149f565b6102c2565b005b610114600480360381019061010f91906115b4565b61044b565b005b34801561012257600080fd5b5061013d600480360381019061013891906114cc565b610588565b60405161014a9190611a10565b60405180910390f35b34801561015f57600080fd5b50610168610874565b60405161017591906119f5565b60405180910390f35b34801561018a57600080fd5b5061019361092d565b005b3480156101a157600080fd5b506101aa610941565b005b3480156101b857600080fd5b506101c1610a87565b6040516101ce9190611988565b60405180910390f35b3480156101e357600080fd5b506101fe60048036038101906101f9919061149f565b610ab1565b005b34801561020c57600080fd5b50610215610af5565b6040516102229190611988565b60405180910390f35b34801561023757600080fd5b50610252600480360381019061024d919061149f565b610b1b565b005b60606000610263858585610b9f565b90508473ffffffffffffffffffffffffffffffffffffffff167f373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e85463486866040516102af93929190611bcf565b60405180910390a2809150509392505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610351576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034890611a8f565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610390610c56565b73ffffffffffffffffffffffffffffffffffffffff16146103e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103dd90611aaf565b60405180910390fd5b6103ef81610cad565b61044881600067ffffffffffffffff81111561040e5761040d611d90565b5b6040519080825280601f01601f1916602001820160405280156104405781602001600182028036833780820191505090505b506000610cb8565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614156104da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d190611a8f565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610519610c56565b73ffffffffffffffffffffffffffffffffffffffff161461056f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056690611aaf565b60405180910390fd5b61057882610cad565b61058482826001610cb8565b5050565b60608573ffffffffffffffffffffffffffffffffffffffff1663095ea7b386866040518363ffffffff1660e01b81526004016105c59291906119cc565b602060405180830381600087803b1580156105df57600080fd5b505af11580156105f3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106179190611610565b506000610625868585610b9f565b90508673ffffffffffffffffffffffffffffffffffffffff1663095ea7b38760006040518363ffffffff1660e01b81526004016106639291906119a3565b602060405180830381600087803b15801561067d57600080fd5b505af1158015610691573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b59190611610565b5060008773ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016106f19190611988565b60206040518083038186803b15801561070957600080fd5b505afa15801561071d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610741919061166a565b905060008111156107fd578773ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b81526004016107a99291906119cc565b602060405180830381600087803b1580156107c357600080fd5b505af11580156107d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107fb9190611610565b505b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f887e0acc3616142401641abfc50d7d7ae169b6ce55d8dc06ff5e21501ddb341b88888860405161085e93929190611bcf565b60405180910390a3819250505095945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610904576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108fb90611acf565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b610935610e35565b61093f6000610eb3565b565b60008060019054906101000a900460ff161590508080156109725750600160008054906101000a900460ff1660ff16105b8061099f575061098130610f79565b15801561099e5750600160008054906101000a900460ff1660ff16145b5b6109de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d590611b0f565b60405180910390fd5b60016000806101000a81548160ff021916908360ff1602179055508015610a1b576001600060016101000a81548160ff0219169083151502179055505b610a23610f9c565b610a2b610ff5565b8015610a845760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610a7b9190611a32565b60405180910390a15b50565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b8060c960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610b23610e35565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610b93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8a90611a6f565b60405180910390fd5b610b9c81610eb3565b50565b60606000808573ffffffffffffffffffffffffffffffffffffffff16348686604051610bcc929190611958565b60006040518083038185875af1925050503d8060008114610c09576040519150601f19603f3d011682016040523d82523d6000602084013e610c0e565b606091505b509150915081610c4a576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80925050509392505050565b6000610c847f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b611046565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610cb5610e35565b50565b610ce47f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b611050565b60000160009054906101000a900460ff1615610d0857610d038361105a565b610e30565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610d4e57600080fd5b505afa925050508015610d7f57506040513d601f19601f82011682018060405250810190610d7c919061163d565b60015b610dbe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610db590611b2f565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114610e23576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1a90611aef565b60405180910390fd5b50610e2f838383611113565b5b505050565b610e3d61113f565b73ffffffffffffffffffffffffffffffffffffffff16610e5b610a87565b73ffffffffffffffffffffffffffffffffffffffff1614610eb1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea890611b6f565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff16610feb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fe290611baf565b60405180910390fd5b610ff3611147565b565b600060019054906101000a900460ff16611044576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103b90611baf565b60405180910390fd5b565b6000819050919050565b6000819050919050565b61106381610f79565b6110a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161109990611b4f565b60405180910390fd5b806110cf7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b611046565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61111c836111a8565b6000825111806111295750805b1561113a5761113883836111f7565b505b505050565b600033905090565b600060019054906101000a900460ff16611196576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118d90611baf565b60405180910390fd5b6111a66111a161113f565b610eb3565b565b6111b18161105a565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b606061121c838360405180606001604052806027815260200161216460279139611224565b905092915050565b60606000808573ffffffffffffffffffffffffffffffffffffffff168560405161124e9190611971565b600060405180830381855af49150503d8060008114611289576040519150601f19603f3d011682016040523d82523d6000602084013e61128e565b606091505b509150915061129f868383876112aa565b925050509392505050565b6060831561130d57600083511415611305576112c585610f79565b611304576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112fb90611b8f565b60405180910390fd5b5b829050611318565b6113178383611320565b5b949350505050565b6000825111156113335781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113679190611a4d565b60405180910390fd5b600061138361137e84611c26565b611c01565b90508281526020810184848401111561139f5761139e611dce565b5b6113aa848285611d1d565b509392505050565b6000813590506113c181612107565b92915050565b6000815190506113d68161211e565b92915050565b6000815190506113eb81612135565b92915050565b60008083601f84011261140757611406611dc4565b5b8235905067ffffffffffffffff81111561142457611423611dbf565b5b6020830191508360018202830111156114405761143f611dc9565b5b9250929050565b600082601f83011261145c5761145b611dc4565b5b813561146c848260208601611370565b91505092915050565b6000813590506114848161214c565b92915050565b6000815190506114998161214c565b92915050565b6000602082840312156114b5576114b4611dd8565b5b60006114c3848285016113b2565b91505092915050565b6000806000806000608086880312156114e8576114e7611dd8565b5b60006114f6888289016113b2565b9550506020611507888289016113b2565b945050604061151888828901611475565b935050606086013567ffffffffffffffff81111561153957611538611dd3565b5b611545888289016113f1565b92509250509295509295909350565b60008060006040848603121561156d5761156c611dd8565b5b600061157b868287016113b2565b935050602084013567ffffffffffffffff81111561159c5761159b611dd3565b5b6115a8868287016113f1565b92509250509250925092565b600080604083850312156115cb576115ca611dd8565b5b60006115d9858286016113b2565b925050602083013567ffffffffffffffff8111156115fa576115f9611dd3565b5b61160685828601611447565b9150509250929050565b60006020828403121561162657611625611dd8565b5b6000611634848285016113c7565b91505092915050565b60006020828403121561165357611652611dd8565b5b6000611661848285016113dc565b91505092915050565b6000602082840312156116805761167f611dd8565b5b600061168e8482850161148a565b91505092915050565b6116a081611c9a565b82525050565b6116af81611cb8565b82525050565b60006116c18385611c6d565b93506116ce838584611d1d565b6116d783611ddd565b840190509392505050565b60006116ee8385611c7e565b93506116fb838584611d1d565b82840190509392505050565b600061171282611c57565b61171c8185611c6d565b935061172c818560208601611d2c565b61173581611ddd565b840191505092915050565b600061174b82611c57565b6117558185611c7e565b9350611765818560208601611d2c565b80840191505092915050565b61177a81611cf9565b82525050565b61178981611d0b565b82525050565b600061179a82611c62565b6117a48185611c89565b93506117b4818560208601611d2c565b6117bd81611ddd565b840191505092915050565b60006117d5602683611c89565b91506117e082611dee565b604082019050919050565b60006117f8602c83611c89565b915061180382611e3d565b604082019050919050565b600061181b602c83611c89565b915061182682611e8c565b604082019050919050565b600061183e603883611c89565b915061184982611edb565b604082019050919050565b6000611861602983611c89565b915061186c82611f2a565b604082019050919050565b6000611884602e83611c89565b915061188f82611f79565b604082019050919050565b60006118a7602e83611c89565b91506118b282611fc8565b604082019050919050565b60006118ca602d83611c89565b91506118d582612017565b604082019050919050565b60006118ed602083611c89565b91506118f882612066565b602082019050919050565b6000611910601d83611c89565b915061191b8261208f565b602082019050919050565b6000611933602b83611c89565b915061193e826120b8565b604082019050919050565b61195281611ce2565b82525050565b60006119658284866116e2565b91508190509392505050565b600061197d8284611740565b915081905092915050565b600060208201905061199d6000830184611697565b92915050565b60006040820190506119b86000830185611697565b6119c56020830184611771565b9392505050565b60006040820190506119e16000830185611697565b6119ee6020830184611949565b9392505050565b6000602082019050611a0a60008301846116a6565b92915050565b60006020820190508181036000830152611a2a8184611707565b905092915050565b6000602082019050611a476000830184611780565b92915050565b60006020820190508181036000830152611a67818461178f565b905092915050565b60006020820190508181036000830152611a88816117c8565b9050919050565b60006020820190508181036000830152611aa8816117eb565b9050919050565b60006020820190508181036000830152611ac88161180e565b9050919050565b60006020820190508181036000830152611ae881611831565b9050919050565b60006020820190508181036000830152611b0881611854565b9050919050565b60006020820190508181036000830152611b2881611877565b9050919050565b60006020820190508181036000830152611b488161189a565b9050919050565b60006020820190508181036000830152611b68816118bd565b9050919050565b60006020820190508181036000830152611b88816118e0565b9050919050565b60006020820190508181036000830152611ba881611903565b9050919050565b60006020820190508181036000830152611bc881611926565b9050919050565b6000604082019050611be46000830186611949565b8181036020830152611bf78184866116b5565b9050949350505050565b6000611c0b611c1c565b9050611c178282611d5f565b919050565b6000604051905090565b600067ffffffffffffffff821115611c4157611c40611d90565b5b611c4a82611ddd565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611ca582611cc2565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611d0482611ce2565b9050919050565b6000611d1682611cec565b9050919050565b82818337600083830152505050565b60005b83811015611d4a578082015181840152602081019050611d2f565b83811115611d59576000848401525b50505050565b611d6882611ddd565b810181811067ffffffffffffffff82111715611d8757611d86611d90565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b61211081611c9a565b811461211b57600080fd5b50565b61212781611cac565b811461213257600080fd5b50565b61213e81611cb8565b811461214957600080fd5b50565b61215581611ce2565b811461216057600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220a349306a7d1588b676f5b6e95783cf1d798266d946f8cdd77432870e89cdcd4f64736f6c63430008070033",
}

// GatewayV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayV2MetaData.ABI instead.
var GatewayV2ABI = GatewayV2MetaData.ABI

// GatewayV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayV2MetaData.Bin instead.
var GatewayV2Bin = GatewayV2MetaData.Bin

// DeployGatewayV2 deploys a new Ethereum contract, binding an instance of GatewayV2 to it.
func DeployGatewayV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayV2, error) {
	parsed, err := GatewayV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayV2{GatewayV2Caller: GatewayV2Caller{contract: contract}, GatewayV2Transactor: GatewayV2Transactor{contract: contract}, GatewayV2Filterer: GatewayV2Filterer{contract: contract}}, nil
}

// GatewayV2 is an auto generated Go binding around an Ethereum contract.
type GatewayV2 struct {
	GatewayV2Caller     // Read-only binding to the contract
	GatewayV2Transactor // Write-only binding to the contract
	GatewayV2Filterer   // Log filterer for contract events
}

// GatewayV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayV2Session struct {
	Contract     *GatewayV2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayV2CallerSession struct {
	Contract *GatewayV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GatewayV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayV2TransactorSession struct {
	Contract     *GatewayV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GatewayV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayV2Raw struct {
	Contract *GatewayV2 // Generic contract binding to access the raw methods on
}

// GatewayV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayV2CallerRaw struct {
	Contract *GatewayV2Caller // Generic read-only contract binding to access the raw methods on
}

// GatewayV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayV2TransactorRaw struct {
	Contract *GatewayV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayV2 creates a new instance of GatewayV2, bound to a specific deployed contract.
func NewGatewayV2(address common.Address, backend bind.ContractBackend) (*GatewayV2, error) {
	contract, err := bindGatewayV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayV2{GatewayV2Caller: GatewayV2Caller{contract: contract}, GatewayV2Transactor: GatewayV2Transactor{contract: contract}, GatewayV2Filterer: GatewayV2Filterer{contract: contract}}, nil
}

// NewGatewayV2Caller creates a new read-only instance of GatewayV2, bound to a specific deployed contract.
func NewGatewayV2Caller(address common.Address, caller bind.ContractCaller) (*GatewayV2Caller, error) {
	contract, err := bindGatewayV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayV2Caller{contract: contract}, nil
}

// NewGatewayV2Transactor creates a new write-only instance of GatewayV2, bound to a specific deployed contract.
func NewGatewayV2Transactor(address common.Address, transactor bind.ContractTransactor) (*GatewayV2Transactor, error) {
	contract, err := bindGatewayV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayV2Transactor{contract: contract}, nil
}

// NewGatewayV2Filterer creates a new log filterer instance of GatewayV2, bound to a specific deployed contract.
func NewGatewayV2Filterer(address common.Address, filterer bind.ContractFilterer) (*GatewayV2Filterer, error) {
	contract, err := bindGatewayV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayV2Filterer{contract: contract}, nil
}

// bindGatewayV2 binds a generic wrapper to an already deployed contract.
func bindGatewayV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayV2 *GatewayV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayV2.Contract.GatewayV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayV2 *GatewayV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayV2.Contract.GatewayV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayV2 *GatewayV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayV2.Contract.GatewayV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayV2 *GatewayV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayV2 *GatewayV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayV2 *GatewayV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayV2.Contract.contract.Transact(opts, method, params...)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayV2 *GatewayV2Caller) Custody(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayV2.contract.Call(opts, &out, "custody")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayV2 *GatewayV2Session) Custody() (common.Address, error) {
	return _GatewayV2.Contract.Custody(&_GatewayV2.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayV2 *GatewayV2CallerSession) Custody() (common.Address, error) {
	return _GatewayV2.Contract.Custody(&_GatewayV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayV2 *GatewayV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayV2 *GatewayV2Session) Owner() (common.Address, error) {
	return _GatewayV2.Contract.Owner(&_GatewayV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayV2 *GatewayV2CallerSession) Owner() (common.Address, error) {
	return _GatewayV2.Contract.Owner(&_GatewayV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayV2 *GatewayV2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayV2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayV2 *GatewayV2Session) ProxiableUUID() ([32]byte, error) {
	return _GatewayV2.Contract.ProxiableUUID(&_GatewayV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayV2 *GatewayV2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayV2.Contract.ProxiableUUID(&_GatewayV2.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayV2 *GatewayV2Transactor) Execute(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayV2.contract.Transact(opts, "execute", destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayV2 *GatewayV2Session) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayV2.Contract.Execute(&_GatewayV2.TransactOpts, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayV2 *GatewayV2TransactorSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayV2.Contract.Execute(&_GatewayV2.TransactOpts, destination, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_GatewayV2 *GatewayV2Transactor) ExecuteWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayV2.contract.Transact(opts, "executeWithERC20", token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_GatewayV2 *GatewayV2Session) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayV2.Contract.ExecuteWithERC20(&_GatewayV2.TransactOpts, token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_GatewayV2 *GatewayV2TransactorSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayV2.Contract.ExecuteWithERC20(&_GatewayV2.TransactOpts, token, to, amount, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GatewayV2 *GatewayV2Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayV2.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GatewayV2 *GatewayV2Session) Initialize() (*types.Transaction, error) {
	return _GatewayV2.Contract.Initialize(&_GatewayV2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GatewayV2 *GatewayV2TransactorSession) Initialize() (*types.Transaction, error) {
	return _GatewayV2.Contract.Initialize(&_GatewayV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayV2 *GatewayV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayV2 *GatewayV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayV2.Contract.RenounceOwnership(&_GatewayV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayV2 *GatewayV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayV2.Contract.RenounceOwnership(&_GatewayV2.TransactOpts)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_GatewayV2 *GatewayV2Transactor) SetCustody(opts *bind.TransactOpts, _custody common.Address) (*types.Transaction, error) {
	return _GatewayV2.contract.Transact(opts, "setCustody", _custody)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_GatewayV2 *GatewayV2Session) SetCustody(_custody common.Address) (*types.Transaction, error) {
	return _GatewayV2.Contract.SetCustody(&_GatewayV2.TransactOpts, _custody)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_GatewayV2 *GatewayV2TransactorSession) SetCustody(_custody common.Address) (*types.Transaction, error) {
	return _GatewayV2.Contract.SetCustody(&_GatewayV2.TransactOpts, _custody)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayV2 *GatewayV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayV2 *GatewayV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayV2.Contract.TransferOwnership(&_GatewayV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayV2 *GatewayV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayV2.Contract.TransferOwnership(&_GatewayV2.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GatewayV2 *GatewayV2Transactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _GatewayV2.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GatewayV2 *GatewayV2Session) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GatewayV2.Contract.UpgradeTo(&_GatewayV2.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GatewayV2 *GatewayV2TransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GatewayV2.Contract.UpgradeTo(&_GatewayV2.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayV2 *GatewayV2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayV2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayV2 *GatewayV2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayV2.Contract.UpgradeToAndCall(&_GatewayV2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayV2 *GatewayV2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayV2.Contract.UpgradeToAndCall(&_GatewayV2.TransactOpts, newImplementation, data)
}

// GatewayV2AdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the GatewayV2 contract.
type GatewayV2AdminChangedIterator struct {
	Event *GatewayV2AdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayV2AdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayV2AdminChanged)
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
		it.Event = new(GatewayV2AdminChanged)
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
func (it *GatewayV2AdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayV2AdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayV2AdminChanged represents a AdminChanged event raised by the GatewayV2 contract.
type GatewayV2AdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GatewayV2 *GatewayV2Filterer) FilterAdminChanged(opts *bind.FilterOpts) (*GatewayV2AdminChangedIterator, error) {

	logs, sub, err := _GatewayV2.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &GatewayV2AdminChangedIterator{contract: _GatewayV2.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GatewayV2 *GatewayV2Filterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayV2AdminChanged) (event.Subscription, error) {

	logs, sub, err := _GatewayV2.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayV2AdminChanged)
				if err := _GatewayV2.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GatewayV2 *GatewayV2Filterer) ParseAdminChanged(log types.Log) (*GatewayV2AdminChanged, error) {
	event := new(GatewayV2AdminChanged)
	if err := _GatewayV2.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayV2BeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the GatewayV2 contract.
type GatewayV2BeaconUpgradedIterator struct {
	Event *GatewayV2BeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayV2BeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayV2BeaconUpgraded)
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
		it.Event = new(GatewayV2BeaconUpgraded)
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
func (it *GatewayV2BeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayV2BeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayV2BeaconUpgraded represents a BeaconUpgraded event raised by the GatewayV2 contract.
type GatewayV2BeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GatewayV2 *GatewayV2Filterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*GatewayV2BeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GatewayV2.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &GatewayV2BeaconUpgradedIterator{contract: _GatewayV2.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GatewayV2 *GatewayV2Filterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayV2BeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GatewayV2.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayV2BeaconUpgraded)
				if err := _GatewayV2.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GatewayV2 *GatewayV2Filterer) ParseBeaconUpgraded(log types.Log) (*GatewayV2BeaconUpgraded, error) {
	event := new(GatewayV2BeaconUpgraded)
	if err := _GatewayV2.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayV2ExecutedV2Iterator is returned from FilterExecutedV2 and is used to iterate over the raw logs and unpacked data for ExecutedV2 events raised by the GatewayV2 contract.
type GatewayV2ExecutedV2Iterator struct {
	Event *GatewayV2ExecutedV2 // Event containing the contract specifics and raw log

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
func (it *GatewayV2ExecutedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayV2ExecutedV2)
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
		it.Event = new(GatewayV2ExecutedV2)
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
func (it *GatewayV2ExecutedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayV2ExecutedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayV2ExecutedV2 represents a ExecutedV2 event raised by the GatewayV2 contract.
type GatewayV2ExecutedV2 struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutedV2 is a free log retrieval operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayV2 *GatewayV2Filterer) FilterExecutedV2(opts *bind.FilterOpts, destination []common.Address) (*GatewayV2ExecutedV2Iterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayV2.contract.FilterLogs(opts, "ExecutedV2", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayV2ExecutedV2Iterator{contract: _GatewayV2.contract, event: "ExecutedV2", logs: logs, sub: sub}, nil
}

// WatchExecutedV2 is a free log subscription operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayV2 *GatewayV2Filterer) WatchExecutedV2(opts *bind.WatchOpts, sink chan<- *GatewayV2ExecutedV2, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayV2.contract.WatchLogs(opts, "ExecutedV2", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayV2ExecutedV2)
				if err := _GatewayV2.contract.UnpackLog(event, "ExecutedV2", log); err != nil {
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

// ParseExecutedV2 is a log parse operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayV2 *GatewayV2Filterer) ParseExecutedV2(log types.Log) (*GatewayV2ExecutedV2, error) {
	event := new(GatewayV2ExecutedV2)
	if err := _GatewayV2.contract.UnpackLog(event, "ExecutedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayV2ExecutedWithERC20V2Iterator is returned from FilterExecutedWithERC20V2 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20V2 events raised by the GatewayV2 contract.
type GatewayV2ExecutedWithERC20V2Iterator struct {
	Event *GatewayV2ExecutedWithERC20V2 // Event containing the contract specifics and raw log

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
func (it *GatewayV2ExecutedWithERC20V2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayV2ExecutedWithERC20V2)
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
		it.Event = new(GatewayV2ExecutedWithERC20V2)
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
func (it *GatewayV2ExecutedWithERC20V2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayV2ExecutedWithERC20V2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayV2ExecutedWithERC20V2 represents a ExecutedWithERC20V2 event raised by the GatewayV2 contract.
type GatewayV2ExecutedWithERC20V2 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20V2 is a free log retrieval operation binding the contract event 0x887e0acc3616142401641abfc50d7d7ae169b6ce55d8dc06ff5e21501ddb341b.
//
// Solidity: event ExecutedWithERC20V2(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayV2 *GatewayV2Filterer) FilterExecutedWithERC20V2(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayV2ExecutedWithERC20V2Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayV2.contract.FilterLogs(opts, "ExecutedWithERC20V2", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayV2ExecutedWithERC20V2Iterator{contract: _GatewayV2.contract, event: "ExecutedWithERC20V2", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20V2 is a free log subscription operation binding the contract event 0x887e0acc3616142401641abfc50d7d7ae169b6ce55d8dc06ff5e21501ddb341b.
//
// Solidity: event ExecutedWithERC20V2(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayV2 *GatewayV2Filterer) WatchExecutedWithERC20V2(opts *bind.WatchOpts, sink chan<- *GatewayV2ExecutedWithERC20V2, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayV2.contract.WatchLogs(opts, "ExecutedWithERC20V2", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayV2ExecutedWithERC20V2)
				if err := _GatewayV2.contract.UnpackLog(event, "ExecutedWithERC20V2", log); err != nil {
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

// ParseExecutedWithERC20V2 is a log parse operation binding the contract event 0x887e0acc3616142401641abfc50d7d7ae169b6ce55d8dc06ff5e21501ddb341b.
//
// Solidity: event ExecutedWithERC20V2(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayV2 *GatewayV2Filterer) ParseExecutedWithERC20V2(log types.Log) (*GatewayV2ExecutedWithERC20V2, error) {
	event := new(GatewayV2ExecutedWithERC20V2)
	if err := _GatewayV2.contract.UnpackLog(event, "ExecutedWithERC20V2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayV2 contract.
type GatewayV2InitializedIterator struct {
	Event *GatewayV2Initialized // Event containing the contract specifics and raw log

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
func (it *GatewayV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayV2Initialized)
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
		it.Event = new(GatewayV2Initialized)
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
func (it *GatewayV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayV2Initialized represents a Initialized event raised by the GatewayV2 contract.
type GatewayV2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayV2 *GatewayV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayV2InitializedIterator, error) {

	logs, sub, err := _GatewayV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayV2InitializedIterator{contract: _GatewayV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayV2 *GatewayV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayV2Initialized) (event.Subscription, error) {

	logs, sub, err := _GatewayV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayV2Initialized)
				if err := _GatewayV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayV2 *GatewayV2Filterer) ParseInitialized(log types.Log) (*GatewayV2Initialized, error) {
	event := new(GatewayV2Initialized)
	if err := _GatewayV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GatewayV2 contract.
type GatewayV2OwnershipTransferredIterator struct {
	Event *GatewayV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayV2OwnershipTransferred)
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
		it.Event = new(GatewayV2OwnershipTransferred)
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
func (it *GatewayV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayV2OwnershipTransferred represents a OwnershipTransferred event raised by the GatewayV2 contract.
type GatewayV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayV2 *GatewayV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayV2OwnershipTransferredIterator{contract: _GatewayV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayV2 *GatewayV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayV2OwnershipTransferred)
				if err := _GatewayV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayV2 *GatewayV2Filterer) ParseOwnershipTransferred(log types.Log) (*GatewayV2OwnershipTransferred, error) {
	event := new(GatewayV2OwnershipTransferred)
	if err := _GatewayV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayV2 contract.
type GatewayV2UpgradedIterator struct {
	Event *GatewayV2Upgraded // Event containing the contract specifics and raw log

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
func (it *GatewayV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayV2Upgraded)
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
		it.Event = new(GatewayV2Upgraded)
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
func (it *GatewayV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayV2Upgraded represents a Upgraded event raised by the GatewayV2 contract.
type GatewayV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayV2 *GatewayV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayV2UpgradedIterator{contract: _GatewayV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayV2 *GatewayV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayV2Upgraded)
				if err := _GatewayV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayV2 *GatewayV2Filterer) ParseUpgraded(log types.Log) (*GatewayV2Upgraded, error) {
	event := new(GatewayV2Upgraded)
	if err := _GatewayV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
