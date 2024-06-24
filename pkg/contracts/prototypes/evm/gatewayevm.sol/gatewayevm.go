// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayevm

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

// GatewayEVMMetaData contains all meta data concerning the GatewayEVM contract.
var GatewayEVMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientETHAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SendFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ExecutedWithERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Send\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SendERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"custody\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeWithERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tssAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_custody\",\"type\":\"address\"}],\"name\":\"setCustody\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b8152503480156200004757600080fd5b50620000586200005e60201b60201c565b62000208565b600060019054906101000a900460ff1615620000b1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a8906200015c565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff1614620001225760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200011991906200017e565b60405180910390a15b565b6000620001336027836200019b565b91506200014082620001b9565b604082019050919050565b6200015681620001ac565b82525050565b60006020820190508181036000830152620001778162000124565b9050919050565b60006020820190506200019560008301846200014b565b92915050565b600082825260208201905092915050565b600060ff82169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60805160601c6126c8620002436000396000818161038701528181610416015281816105100152818161059f015261093b01526126c86000f3fe6080604052600436106100dd5760003560e01c80638da5cb5b1161007f578063c4d66de811610059578063c4d66de814610271578063cb0271ed1461029a578063dda79b75146102c3578063f2fde38b146102ee576100dd565b80638da5cb5b146102015780639372c4ab1461022c578063ae7a3a6f14610248576100dd565b80635131ab59116100bb5780635131ab591461015757806352d1902d146101945780635b112591146101bf578063715018a6146101ea576100dd565b80631cff79cd146100e25780633659cfe6146101125780634f1ef2861461013b575b600080fd5b6100fc60048036038101906100f791906118d6565b610317565b6040516101099190611f15565b60405180910390f35b34801561011e57600080fd5b5061013960048036038101906101349190611821565b610385565b005b61015560048036038101906101509190611936565b61050e565b005b34801561016357600080fd5b5061017e6004803603810190610179919061184e565b61064b565b60405161018b9190611f15565b60405180910390f35b3480156101a057600080fd5b506101a9610937565b6040516101b69190611efa565b60405180910390f35b3480156101cb57600080fd5b506101d46109f0565b6040516101e19190611e16565b60405180910390f35b3480156101f657600080fd5b506101ff610a16565b005b34801561020d57600080fd5b50610216610a2a565b6040516102239190611e16565b60405180910390f35b61024660048036038101906102419190611a60565b610a54565b005b34801561025457600080fd5b5061026f600480360381019061026a9190611821565b610b9f565b005b34801561027d57600080fd5b5061029860048036038101906102939190611821565b610be3565b005b3480156102a657600080fd5b506102c160048036038101906102bc91906119ec565b610d6b565b005b3480156102cf57600080fd5b506102d8610e77565b6040516102e59190611e16565b60405180910390f35b3480156102fa57600080fd5b5061031560048036038101906103109190611821565b610e9d565b005b60606000610326858585610f21565b90508473ffffffffffffffffffffffffffffffffffffffff167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f348686604051610372939291906120d4565b60405180910390a2809150509392505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040b90611f94565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610453610fd8565b73ffffffffffffffffffffffffffffffffffffffff16146104a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a090611fb4565b60405180910390fd5b6104b28161102f565b61050b81600067ffffffffffffffff8111156104d1576104d0612295565b5b6040519080825280601f01601f1916602001820160405280156105035781602001600182028036833780820191505090505b50600061103a565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16141561059d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059490611f94565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166105dc610fd8565b73ffffffffffffffffffffffffffffffffffffffff1614610632576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062990611fb4565b60405180910390fd5b61063b8261102f565b6106478282600161103a565b5050565b60608573ffffffffffffffffffffffffffffffffffffffff1663095ea7b386866040518363ffffffff1660e01b8152600401610688929190611ed1565b602060405180830381600087803b1580156106a257600080fd5b505af11580156106b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106da9190611992565b5060006106e8868585610f21565b90508673ffffffffffffffffffffffffffffffffffffffff1663095ea7b38760006040518363ffffffff1660e01b8152600401610726929190611ea8565b602060405180830381600087803b15801561074057600080fd5b505af1158015610754573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107789190611992565b5060008773ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016107b49190611e16565b60206040518083038186803b1580156107cc57600080fd5b505afa1580156107e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108049190611ac0565b905060008111156108c0578773ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161086c929190611ed1565b602060405180830381600087803b15801561088657600080fd5b505af115801561089a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108be9190611992565b505b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382888888604051610921939291906120d4565b60405180910390a3819250505095945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16146109c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109be90611fd4565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b60ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610a1e6111b7565b610a286000611235565b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000341415610a8f576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1634604051610ad790611e01565b60006040518083038185875af1925050503d8060008114610b14576040519150601f19603f3d011682016040523d82523d6000602084013e610b19565b606091505b50509050600015158115151415610b5c576040517f81063e5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f84583322b159855a8990ae36b9708ebcb1d1c7fdccf3cd41e19e4f65f00bb2b633858534604051610b919493929190611e68565b60405180910390a150505050565b8060c960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060019054906101000a900460ff16159050808015610c145750600160008054906101000a900460ff1660ff16105b80610c415750610c23306112fb565b158015610c405750600160008054906101000a900460ff1660ff16145b5b610c80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7790612014565b60405180910390fd5b60016000806101000a81548160ff021916908360ff1602179055508015610cbd576001600060016101000a81548160ff0219169083151502179055505b610cc561131e565b610ccd611377565b8160ca60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015610d675760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610d5e9190611f37565b60405180910390a15b5050565b8173ffffffffffffffffffffffffffffffffffffffff166323b872dd3360c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401610dca93929190611e31565b602060405180830381600087803b158015610de457600080fd5b505af1158015610df8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1c9190611992565b508173ffffffffffffffffffffffffffffffffffffffff167ff31c5882db91f6b727bb3fe49d02ca7e0e036d99bd5f6bb64cb0c8d9f74f991133868685604051610e699493929190611e68565b60405180910390a250505050565b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610ea56111b7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610f15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f0c90611f74565b60405180910390fd5b610f1e81611235565b50565b60606000808573ffffffffffffffffffffffffffffffffffffffff16348686604051610f4e929190611dd1565b60006040518083038185875af1925050503d8060008114610f8b576040519150601f19603f3d011682016040523d82523d6000602084013e610f90565b606091505b509150915081610fcc576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80925050509392505050565b60006110067f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b6113c8565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6110376111b7565b50565b6110667f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b6113d2565b60000160009054906101000a900460ff161561108a57611085836113dc565b6111b2565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156110d057600080fd5b505afa92505050801561110157506040513d601f19601f820116820180604052508101906110fe91906119bf565b60015b611140576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161113790612034565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b81146111a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161119c90611ff4565b60405180910390fd5b506111b1838383611495565b5b505050565b6111bf6114c1565b73ffffffffffffffffffffffffffffffffffffffff166111dd610a2a565b73ffffffffffffffffffffffffffffffffffffffff1614611233576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122a90612074565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff1661136d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611364906120b4565b60405180910390fd5b6113756114c9565b565b600060019054906101000a900460ff166113c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113bd906120b4565b60405180910390fd5b565b6000819050919050565b6000819050919050565b6113e5816112fb565b611424576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141b90612054565b60405180910390fd5b806114517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b6113c8565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61149e8361152a565b6000825111806114ab5750805b156114bc576114ba8383611579565b505b505050565b600033905090565b600060019054906101000a900460ff16611518576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161150f906120b4565b60405180910390fd5b6115286115236114c1565b611235565b565b611533816113dc565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b606061159e838360405180606001604052806027815260200161266c602791396115a6565b905092915050565b60606000808573ffffffffffffffffffffffffffffffffffffffff16856040516115d09190611dea565b600060405180830381855af49150503d806000811461160b576040519150601f19603f3d011682016040523d82523d6000602084013e611610565b606091505b50915091506116218683838761162c565b925050509392505050565b6060831561168f5760008351141561168757611647856112fb565b611686576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161167d90612094565b60405180910390fd5b5b82905061169a565b61169983836116a2565b5b949350505050565b6000825111156116b55781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116e99190611f52565b60405180910390fd5b60006117056117008461212b565b612106565b905082815260208101848484011115611721576117206122d3565b5b61172c848285612222565b509392505050565b6000813590506117438161260f565b92915050565b60008151905061175881612626565b92915050565b60008151905061176d8161263d565b92915050565b60008083601f840112611789576117886122c9565b5b8235905067ffffffffffffffff8111156117a6576117a56122c4565b5b6020830191508360018202830111156117c2576117c16122ce565b5b9250929050565b600082601f8301126117de576117dd6122c9565b5b81356117ee8482602086016116f2565b91505092915050565b60008135905061180681612654565b92915050565b60008151905061181b81612654565b92915050565b600060208284031215611837576118366122dd565b5b600061184584828501611734565b91505092915050565b60008060008060006080868803121561186a576118696122dd565b5b600061187888828901611734565b955050602061188988828901611734565b945050604061189a888289016117f7565b935050606086013567ffffffffffffffff8111156118bb576118ba6122d8565b5b6118c788828901611773565b92509250509295509295909350565b6000806000604084860312156118ef576118ee6122dd565b5b60006118fd86828701611734565b935050602084013567ffffffffffffffff81111561191e5761191d6122d8565b5b61192a86828701611773565b92509250509250925092565b6000806040838503121561194d5761194c6122dd565b5b600061195b85828601611734565b925050602083013567ffffffffffffffff81111561197c5761197b6122d8565b5b611988858286016117c9565b9150509250929050565b6000602082840312156119a8576119a76122dd565b5b60006119b684828501611749565b91505092915050565b6000602082840312156119d5576119d46122dd565b5b60006119e38482850161175e565b91505092915050565b60008060008060608587031215611a0657611a056122dd565b5b600085013567ffffffffffffffff811115611a2457611a236122d8565b5b611a3087828801611773565b94509450506020611a4387828801611734565b9250506040611a54878288016117f7565b91505092959194509250565b600080600060408486031215611a7957611a786122dd565b5b600084013567ffffffffffffffff811115611a9757611a966122d8565b5b611aa386828701611773565b93509350506020611ab6868287016117f7565b9150509250925092565b600060208284031215611ad657611ad56122dd565b5b6000611ae48482850161180c565b91505092915050565b611af68161219f565b82525050565b611b05816121bd565b82525050565b6000611b178385612172565b9350611b24838584612222565b611b2d836122e2565b840190509392505050565b6000611b448385612183565b9350611b51838584612222565b82840190509392505050565b6000611b688261215c565b611b728185612172565b9350611b82818560208601612231565b611b8b816122e2565b840191505092915050565b6000611ba18261215c565b611bab8185612183565b9350611bbb818560208601612231565b80840191505092915050565b611bd0816121fe565b82525050565b611bdf81612210565b82525050565b6000611bf082612167565b611bfa818561218e565b9350611c0a818560208601612231565b611c13816122e2565b840191505092915050565b6000611c2b60268361218e565b9150611c36826122f3565b604082019050919050565b6000611c4e602c8361218e565b9150611c5982612342565b604082019050919050565b6000611c71602c8361218e565b9150611c7c82612391565b604082019050919050565b6000611c9460388361218e565b9150611c9f826123e0565b604082019050919050565b6000611cb760298361218e565b9150611cc28261242f565b604082019050919050565b6000611cda602e8361218e565b9150611ce58261247e565b604082019050919050565b6000611cfd602e8361218e565b9150611d08826124cd565b604082019050919050565b6000611d20602d8361218e565b9150611d2b8261251c565b604082019050919050565b6000611d4360208361218e565b9150611d4e8261256b565b602082019050919050565b6000611d66600083612183565b9150611d7182612594565b600082019050919050565b6000611d89601d8361218e565b9150611d9482612597565b602082019050919050565b6000611dac602b8361218e565b9150611db7826125c0565b604082019050919050565b611dcb816121e7565b82525050565b6000611dde828486611b38565b91508190509392505050565b6000611df68284611b96565b915081905092915050565b6000611e0c82611d59565b9150819050919050565b6000602082019050611e2b6000830184611aed565b92915050565b6000606082019050611e466000830186611aed565b611e536020830185611aed565b611e606040830184611dc2565b949350505050565b6000606082019050611e7d6000830187611aed565b8181036020830152611e90818587611b0b565b9050611e9f6040830184611dc2565b95945050505050565b6000604082019050611ebd6000830185611aed565b611eca6020830184611bc7565b9392505050565b6000604082019050611ee66000830185611aed565b611ef36020830184611dc2565b9392505050565b6000602082019050611f0f6000830184611afc565b92915050565b60006020820190508181036000830152611f2f8184611b5d565b905092915050565b6000602082019050611f4c6000830184611bd6565b92915050565b60006020820190508181036000830152611f6c8184611be5565b905092915050565b60006020820190508181036000830152611f8d81611c1e565b9050919050565b60006020820190508181036000830152611fad81611c41565b9050919050565b60006020820190508181036000830152611fcd81611c64565b9050919050565b60006020820190508181036000830152611fed81611c87565b9050919050565b6000602082019050818103600083015261200d81611caa565b9050919050565b6000602082019050818103600083015261202d81611ccd565b9050919050565b6000602082019050818103600083015261204d81611cf0565b9050919050565b6000602082019050818103600083015261206d81611d13565b9050919050565b6000602082019050818103600083015261208d81611d36565b9050919050565b600060208201905081810360008301526120ad81611d7c565b9050919050565b600060208201905081810360008301526120cd81611d9f565b9050919050565b60006040820190506120e96000830186611dc2565b81810360208301526120fc818486611b0b565b9050949350505050565b6000612110612121565b905061211c8282612264565b919050565b6000604051905090565b600067ffffffffffffffff82111561214657612145612295565b5b61214f826122e2565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006121aa826121c7565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000612209826121e7565b9050919050565b600061221b826121f1565b9050919050565b82818337600083830152505050565b60005b8381101561224f578082015181840152602081019050612234565b8381111561225e576000848401525b50505050565b61226d826122e2565b810181811067ffffffffffffffff8211171561228c5761228b612295565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b6126188161219f565b811461262357600080fd5b50565b61262f816121b1565b811461263a57600080fd5b50565b612646816121bd565b811461265157600080fd5b50565b61265d816121e7565b811461266857600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220ea1799872442b30dbe6690f698427118dc3ef24baf52729719a1766515b7cfca64736f6c63430008070033",
}

// GatewayEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMMetaData.ABI instead.
var GatewayEVMABI = GatewayEVMMetaData.ABI

// GatewayEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMMetaData.Bin instead.
var GatewayEVMBin = GatewayEVMMetaData.Bin

// DeployGatewayEVM deploys a new Ethereum contract, binding an instance of GatewayEVM to it.
func DeployGatewayEVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVM, error) {
	parsed, err := GatewayEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVM{GatewayEVMCaller: GatewayEVMCaller{contract: contract}, GatewayEVMTransactor: GatewayEVMTransactor{contract: contract}, GatewayEVMFilterer: GatewayEVMFilterer{contract: contract}}, nil
}

// GatewayEVM is an auto generated Go binding around an Ethereum contract.
type GatewayEVM struct {
	GatewayEVMCaller     // Read-only binding to the contract
	GatewayEVMTransactor // Write-only binding to the contract
	GatewayEVMFilterer   // Log filterer for contract events
}

// GatewayEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMSession struct {
	Contract     *GatewayEVM       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMCallerSession struct {
	Contract *GatewayEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GatewayEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMTransactorSession struct {
	Contract     *GatewayEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GatewayEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMRaw struct {
	Contract *GatewayEVM // Generic contract binding to access the raw methods on
}

// GatewayEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMCallerRaw struct {
	Contract *GatewayEVMCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMTransactorRaw struct {
	Contract *GatewayEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVM creates a new instance of GatewayEVM, bound to a specific deployed contract.
func NewGatewayEVM(address common.Address, backend bind.ContractBackend) (*GatewayEVM, error) {
	contract, err := bindGatewayEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVM{GatewayEVMCaller: GatewayEVMCaller{contract: contract}, GatewayEVMTransactor: GatewayEVMTransactor{contract: contract}, GatewayEVMFilterer: GatewayEVMFilterer{contract: contract}}, nil
}

// NewGatewayEVMCaller creates a new read-only instance of GatewayEVM, bound to a specific deployed contract.
func NewGatewayEVMCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMCaller, error) {
	contract, err := bindGatewayEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMCaller{contract: contract}, nil
}

// NewGatewayEVMTransactor creates a new write-only instance of GatewayEVM, bound to a specific deployed contract.
func NewGatewayEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMTransactor, error) {
	contract, err := bindGatewayEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTransactor{contract: contract}, nil
}

// NewGatewayEVMFilterer creates a new log filterer instance of GatewayEVM, bound to a specific deployed contract.
func NewGatewayEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMFilterer, error) {
	contract, err := bindGatewayEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMFilterer{contract: contract}, nil
}

// bindGatewayEVM binds a generic wrapper to an already deployed contract.
func bindGatewayEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVM *GatewayEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVM.Contract.GatewayEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVM *GatewayEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVM.Contract.GatewayEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVM *GatewayEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVM.Contract.GatewayEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVM *GatewayEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVM *GatewayEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVM *GatewayEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVM.Contract.contract.Transact(opts, method, params...)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVM *GatewayEVMCaller) Custody(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "custody")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVM *GatewayEVMSession) Custody() (common.Address, error) {
	return _GatewayEVM.Contract.Custody(&_GatewayEVM.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVM *GatewayEVMCallerSession) Custody() (common.Address, error) {
	return _GatewayEVM.Contract.Custody(&_GatewayEVM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayEVM *GatewayEVMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayEVM *GatewayEVMSession) Owner() (common.Address, error) {
	return _GatewayEVM.Contract.Owner(&_GatewayEVM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayEVM *GatewayEVMCallerSession) Owner() (common.Address, error) {
	return _GatewayEVM.Contract.Owner(&_GatewayEVM.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVM *GatewayEVMCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVM *GatewayEVMSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayEVM.Contract.ProxiableUUID(&_GatewayEVM.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVM *GatewayEVMCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayEVM.Contract.ProxiableUUID(&_GatewayEVM.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVM *GatewayEVMCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVM *GatewayEVMSession) TssAddress() (common.Address, error) {
	return _GatewayEVM.Contract.TssAddress(&_GatewayEVM.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVM *GatewayEVMCallerSession) TssAddress() (common.Address, error) {
	return _GatewayEVM.Contract.TssAddress(&_GatewayEVM.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVM *GatewayEVMTransactor) Execute(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "execute", destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVM *GatewayEVMSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Execute(&_GatewayEVM.TransactOpts, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVM *GatewayEVMTransactorSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Execute(&_GatewayEVM.TransactOpts, destination, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_GatewayEVM *GatewayEVMTransactor) ExecuteWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "executeWithERC20", token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_GatewayEVM *GatewayEVMSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.ExecuteWithERC20(&_GatewayEVM.TransactOpts, token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_GatewayEVM *GatewayEVMTransactorSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.ExecuteWithERC20(&_GatewayEVM.TransactOpts, token, to, amount, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tssAddress) returns()
func (_GatewayEVM *GatewayEVMTransactor) Initialize(opts *bind.TransactOpts, _tssAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "initialize", _tssAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tssAddress) returns()
func (_GatewayEVM *GatewayEVMSession) Initialize(_tssAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Initialize(&_GatewayEVM.TransactOpts, _tssAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tssAddress) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) Initialize(_tssAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Initialize(&_GatewayEVM.TransactOpts, _tssAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayEVM *GatewayEVMTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayEVM *GatewayEVMSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayEVM.Contract.RenounceOwnership(&_GatewayEVM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayEVM *GatewayEVMTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayEVM.Contract.RenounceOwnership(&_GatewayEVM.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x9372c4ab.
//
// Solidity: function send(bytes recipient, uint256 amount) payable returns()
func (_GatewayEVM *GatewayEVMTransactor) Send(opts *bind.TransactOpts, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "send", recipient, amount)
}

// Send is a paid mutator transaction binding the contract method 0x9372c4ab.
//
// Solidity: function send(bytes recipient, uint256 amount) payable returns()
func (_GatewayEVM *GatewayEVMSession) Send(recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Send(&_GatewayEVM.TransactOpts, recipient, amount)
}

// Send is a paid mutator transaction binding the contract method 0x9372c4ab.
//
// Solidity: function send(bytes recipient, uint256 amount) payable returns()
func (_GatewayEVM *GatewayEVMTransactorSession) Send(recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Send(&_GatewayEVM.TransactOpts, recipient, amount)
}

// SendERC20 is a paid mutator transaction binding the contract method 0xcb0271ed.
//
// Solidity: function sendERC20(bytes recipient, address token, uint256 amount) returns()
func (_GatewayEVM *GatewayEVMTransactor) SendERC20(opts *bind.TransactOpts, recipient []byte, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "sendERC20", recipient, token, amount)
}

// SendERC20 is a paid mutator transaction binding the contract method 0xcb0271ed.
//
// Solidity: function sendERC20(bytes recipient, address token, uint256 amount) returns()
func (_GatewayEVM *GatewayEVMSession) SendERC20(recipient []byte, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GatewayEVM.Contract.SendERC20(&_GatewayEVM.TransactOpts, recipient, token, amount)
}

// SendERC20 is a paid mutator transaction binding the contract method 0xcb0271ed.
//
// Solidity: function sendERC20(bytes recipient, address token, uint256 amount) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) SendERC20(recipient []byte, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GatewayEVM.Contract.SendERC20(&_GatewayEVM.TransactOpts, recipient, token, amount)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_GatewayEVM *GatewayEVMTransactor) SetCustody(opts *bind.TransactOpts, _custody common.Address) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "setCustody", _custody)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_GatewayEVM *GatewayEVMSession) SetCustody(_custody common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.SetCustody(&_GatewayEVM.TransactOpts, _custody)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) SetCustody(_custody common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.SetCustody(&_GatewayEVM.TransactOpts, _custody)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayEVM *GatewayEVMTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayEVM *GatewayEVMSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.TransferOwnership(&_GatewayEVM.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.TransferOwnership(&_GatewayEVM.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GatewayEVM *GatewayEVMTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GatewayEVM *GatewayEVMSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.UpgradeTo(&_GatewayEVM.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.UpgradeTo(&_GatewayEVM.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVM *GatewayEVMTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVM *GatewayEVMSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.UpgradeToAndCall(&_GatewayEVM.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVM *GatewayEVMTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.UpgradeToAndCall(&_GatewayEVM.TransactOpts, newImplementation, data)
}

// GatewayEVMAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the GatewayEVM contract.
type GatewayEVMAdminChangedIterator struct {
	Event *GatewayEVMAdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayEVMAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMAdminChanged)
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
		it.Event = new(GatewayEVMAdminChanged)
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
func (it *GatewayEVMAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMAdminChanged represents a AdminChanged event raised by the GatewayEVM contract.
type GatewayEVMAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GatewayEVM *GatewayEVMFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*GatewayEVMAdminChangedIterator, error) {

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMAdminChangedIterator{contract: _GatewayEVM.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GatewayEVM *GatewayEVMFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayEVMAdminChanged) (event.Subscription, error) {

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMAdminChanged)
				if err := _GatewayEVM.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_GatewayEVM *GatewayEVMFilterer) ParseAdminChanged(log types.Log) (*GatewayEVMAdminChanged, error) {
	event := new(GatewayEVMAdminChanged)
	if err := _GatewayEVM.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the GatewayEVM contract.
type GatewayEVMBeaconUpgradedIterator struct {
	Event *GatewayEVMBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayEVMBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMBeaconUpgraded)
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
		it.Event = new(GatewayEVMBeaconUpgraded)
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
func (it *GatewayEVMBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMBeaconUpgraded represents a BeaconUpgraded event raised by the GatewayEVM contract.
type GatewayEVMBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GatewayEVM *GatewayEVMFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*GatewayEVMBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMBeaconUpgradedIterator{contract: _GatewayEVM.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GatewayEVM *GatewayEVMFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayEVMBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMBeaconUpgraded)
				if err := _GatewayEVM.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_GatewayEVM *GatewayEVMFilterer) ParseBeaconUpgraded(log types.Log) (*GatewayEVMBeaconUpgraded, error) {
	event := new(GatewayEVMBeaconUpgraded)
	if err := _GatewayEVM.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVM contract.
type GatewayEVMExecutedIterator struct {
	Event *GatewayEVMExecuted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMExecuted)
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
		it.Event = new(GatewayEVMExecuted)
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
func (it *GatewayEVMExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMExecuted represents a Executed event raised by the GatewayEVM contract.
type GatewayEVMExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVM *GatewayEVMFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMExecutedIterator{contract: _GatewayEVM.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVM *GatewayEVMFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMExecuted)
				if err := _GatewayEVM.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVM *GatewayEVMFilterer) ParseExecuted(log types.Log) (*GatewayEVMExecuted, error) {
	event := new(GatewayEVMExecuted)
	if err := _GatewayEVM.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVM contract.
type GatewayEVMExecutedWithERC20Iterator struct {
	Event *GatewayEVMExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMExecutedWithERC20)
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
		it.Event = new(GatewayEVMExecutedWithERC20)
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
func (it *GatewayEVMExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVM contract.
type GatewayEVMExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVM *GatewayEVMFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMExecutedWithERC20Iterator{contract: _GatewayEVM.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVM *GatewayEVMFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMExecutedWithERC20)
				if err := _GatewayEVM.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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

// ParseExecutedWithERC20 is a log parse operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVM *GatewayEVMFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMExecutedWithERC20, error) {
	event := new(GatewayEVMExecutedWithERC20)
	if err := _GatewayEVM.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayEVM contract.
type GatewayEVMInitializedIterator struct {
	Event *GatewayEVMInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInitialized)
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
		it.Event = new(GatewayEVMInitialized)
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
func (it *GatewayEVMInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInitialized represents a Initialized event raised by the GatewayEVM contract.
type GatewayEVMInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayEVM *GatewayEVMFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayEVMInitializedIterator, error) {

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInitializedIterator{contract: _GatewayEVM.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayEVM *GatewayEVMFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayEVMInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInitialized)
				if err := _GatewayEVM.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayEVM *GatewayEVMFilterer) ParseInitialized(log types.Log) (*GatewayEVMInitialized, error) {
	event := new(GatewayEVMInitialized)
	if err := _GatewayEVM.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GatewayEVM contract.
type GatewayEVMOwnershipTransferredIterator struct {
	Event *GatewayEVMOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayEVMOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMOwnershipTransferred)
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
		it.Event = new(GatewayEVMOwnershipTransferred)
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
func (it *GatewayEVMOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMOwnershipTransferred represents a OwnershipTransferred event raised by the GatewayEVM contract.
type GatewayEVMOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayEVM *GatewayEVMFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayEVMOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMOwnershipTransferredIterator{contract: _GatewayEVM.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayEVM *GatewayEVMFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayEVMOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMOwnershipTransferred)
				if err := _GatewayEVM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GatewayEVM *GatewayEVMFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayEVMOwnershipTransferred, error) {
	event := new(GatewayEVMOwnershipTransferred)
	if err := _GatewayEVM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMSendIterator is returned from FilterSend and is used to iterate over the raw logs and unpacked data for Send events raised by the GatewayEVM contract.
type GatewayEVMSendIterator struct {
	Event *GatewayEVMSend // Event containing the contract specifics and raw log

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
func (it *GatewayEVMSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMSend)
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
		it.Event = new(GatewayEVMSend)
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
func (it *GatewayEVMSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMSend represents a Send event raised by the GatewayEVM contract.
type GatewayEVMSend struct {
	Sender    common.Address
	Recipient []byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSend is a free log retrieval operation binding the contract event 0x84583322b159855a8990ae36b9708ebcb1d1c7fdccf3cd41e19e4f65f00bb2b6.
//
// Solidity: event Send(address sender, bytes recipient, uint256 amount)
func (_GatewayEVM *GatewayEVMFilterer) FilterSend(opts *bind.FilterOpts) (*GatewayEVMSendIterator, error) {

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMSendIterator{contract: _GatewayEVM.contract, event: "Send", logs: logs, sub: sub}, nil
}

// WatchSend is a free log subscription operation binding the contract event 0x84583322b159855a8990ae36b9708ebcb1d1c7fdccf3cd41e19e4f65f00bb2b6.
//
// Solidity: event Send(address sender, bytes recipient, uint256 amount)
func (_GatewayEVM *GatewayEVMFilterer) WatchSend(opts *bind.WatchOpts, sink chan<- *GatewayEVMSend) (event.Subscription, error) {

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMSend)
				if err := _GatewayEVM.contract.UnpackLog(event, "Send", log); err != nil {
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

// ParseSend is a log parse operation binding the contract event 0x84583322b159855a8990ae36b9708ebcb1d1c7fdccf3cd41e19e4f65f00bb2b6.
//
// Solidity: event Send(address sender, bytes recipient, uint256 amount)
func (_GatewayEVM *GatewayEVMFilterer) ParseSend(log types.Log) (*GatewayEVMSend, error) {
	event := new(GatewayEVMSend)
	if err := _GatewayEVM.contract.UnpackLog(event, "Send", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMSendERC20Iterator is returned from FilterSendERC20 and is used to iterate over the raw logs and unpacked data for SendERC20 events raised by the GatewayEVM contract.
type GatewayEVMSendERC20Iterator struct {
	Event *GatewayEVMSendERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMSendERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMSendERC20)
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
		it.Event = new(GatewayEVMSendERC20)
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
func (it *GatewayEVMSendERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMSendERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMSendERC20 represents a SendERC20 event raised by the GatewayEVM contract.
type GatewayEVMSendERC20 struct {
	Sender    common.Address
	Recipient []byte
	Asset     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSendERC20 is a free log retrieval operation binding the contract event 0xf31c5882db91f6b727bb3fe49d02ca7e0e036d99bd5f6bb64cb0c8d9f74f9911.
//
// Solidity: event SendERC20(address sender, bytes recipient, address indexed asset, uint256 amount)
func (_GatewayEVM *GatewayEVMFilterer) FilterSendERC20(opts *bind.FilterOpts, asset []common.Address) (*GatewayEVMSendERC20Iterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "SendERC20", assetRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMSendERC20Iterator{contract: _GatewayEVM.contract, event: "SendERC20", logs: logs, sub: sub}, nil
}

// WatchSendERC20 is a free log subscription operation binding the contract event 0xf31c5882db91f6b727bb3fe49d02ca7e0e036d99bd5f6bb64cb0c8d9f74f9911.
//
// Solidity: event SendERC20(address sender, bytes recipient, address indexed asset, uint256 amount)
func (_GatewayEVM *GatewayEVMFilterer) WatchSendERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMSendERC20, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "SendERC20", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMSendERC20)
				if err := _GatewayEVM.contract.UnpackLog(event, "SendERC20", log); err != nil {
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

// ParseSendERC20 is a log parse operation binding the contract event 0xf31c5882db91f6b727bb3fe49d02ca7e0e036d99bd5f6bb64cb0c8d9f74f9911.
//
// Solidity: event SendERC20(address sender, bytes recipient, address indexed asset, uint256 amount)
func (_GatewayEVM *GatewayEVMFilterer) ParseSendERC20(log types.Log) (*GatewayEVMSendERC20, error) {
	event := new(GatewayEVMSendERC20)
	if err := _GatewayEVM.contract.UnpackLog(event, "SendERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayEVM contract.
type GatewayEVMUpgradedIterator struct {
	Event *GatewayEVMUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgraded)
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
		it.Event = new(GatewayEVMUpgraded)
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
func (it *GatewayEVMUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgraded represents a Upgraded event raised by the GatewayEVM contract.
type GatewayEVMUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVM *GatewayEVMFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayEVMUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradedIterator{contract: _GatewayEVM.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVM *GatewayEVMFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgraded)
				if err := _GatewayEVM.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GatewayEVM *GatewayEVMFilterer) ParseUpgraded(log types.Log) (*GatewayEVMUpgraded, error) {
	event := new(GatewayEVMUpgraded)
	if err := _GatewayEVM.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
