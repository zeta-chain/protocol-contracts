// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayevmupgradetest

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

// GatewayEVMUpgradeTestMetaData contains all meta data concerning the GatewayEVMUpgradeTest contract.
var GatewayEVMUpgradeTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientETHAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SendFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ExecutedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ExecutedWithERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Send\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SendERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"custody\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeWithERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tssAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_custody\",\"type\":\"address\"}],\"name\":\"setCustody\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b8152503480156200004757600080fd5b50620000586200005e60201b60201c565b62000208565b600060019054906101000a900460ff1615620000b1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a8906200015c565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff1614620001225760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200011991906200017e565b60405180910390a15b565b6000620001336027836200019b565b91506200014082620001b9565b604082019050919050565b6200015681620001ac565b82525050565b60006020820190508181036000830152620001778162000124565b9050919050565b60006020820190506200019560008301846200014b565b92915050565b600082825260208201905092915050565b600060ff82169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60805160601c6127ac620002436000396000818161038701528181610416015281816105100152818161059f01526109ca01526127ac6000f3fe6080604052600436106100dd5760003560e01c80638da5cb5b1161007f578063c4d66de811610059578063c4d66de814610271578063cb0271ed1461029a578063dda79b75146102c3578063f2fde38b146102ee576100dd565b80638da5cb5b146102015780639372c4ab1461022c578063ae7a3a6f14610248576100dd565b80635131ab59116100bb5780635131ab591461015757806352d1902d146101945780635b112591146101bf578063715018a6146101ea576100dd565b80631cff79cd146100e25780633659cfe6146101125780634f1ef2861461013b575b600080fd5b6100fc60048036038101906100f791906119c8565b610317565b6040516101099190611ff9565b60405180910390f35b34801561011e57600080fd5b5061013960048036038101906101349190611913565b610385565b005b61015560048036038101906101509190611a28565b61050e565b005b34801561016357600080fd5b5061017e60048036038101906101799190611940565b61064b565b60405161018b9190611ff9565b60405180910390f35b3480156101a057600080fd5b506101a96109c6565b6040516101b69190611fac565b60405180910390f35b3480156101cb57600080fd5b506101d4610a7f565b6040516101e19190611f08565b60405180910390f35b3480156101f657600080fd5b506101ff610aa5565b005b34801561020d57600080fd5b50610216610ab9565b6040516102239190611f08565b60405180910390f35b61024660048036038101906102419190611b52565b610ae3565b005b34801561025457600080fd5b5061026f600480360381019061026a9190611913565b610c2c565b005b34801561027d57600080fd5b5061029860048036038101906102939190611913565b610c70565b005b3480156102a657600080fd5b506102c160048036038101906102bc9190611ade565b610e5f565b005b3480156102cf57600080fd5b506102d8610f69565b6040516102e59190611f08565b60405180910390f35b3480156102fa57600080fd5b5061031560048036038101906103109190611913565b610f8f565b005b60606000610326858585611013565b90508473ffffffffffffffffffffffffffffffffffffffff167f373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546348686604051610372939291906121b8565b60405180910390a2809150509392505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040b90612078565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166104536110ca565b73ffffffffffffffffffffffffffffffffffffffff16146104a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a090612098565b60405180910390fd5b6104b281611121565b61050b81600067ffffffffffffffff8111156104d1576104d0612379565b5b6040519080825280601f01601f1916602001820160405280156105035781602001600182028036833780820191505090505b50600061112c565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16141561059d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059490612078565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166105dc6110ca565b73ffffffffffffffffffffffffffffffffffffffff1614610632576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062990612098565b60405180910390fd5b61063b82611121565b6106478282600161112c565b5050565b60608573ffffffffffffffffffffffffffffffffffffffff1663095ea7b38660006040518363ffffffff1660e01b8152600401610689929190611f5a565b602060405180830381600087803b1580156106a357600080fd5b505af11580156106b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106db9190611a84565b508573ffffffffffffffffffffffffffffffffffffffff1663095ea7b386866040518363ffffffff1660e01b8152600401610717929190611f83565b602060405180830381600087803b15801561073157600080fd5b505af1158015610745573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107699190611a84565b506000610777868585611013565b90508673ffffffffffffffffffffffffffffffffffffffff1663095ea7b38760006040518363ffffffff1660e01b81526004016107b5929190611f5a565b602060405180830381600087803b1580156107cf57600080fd5b505af11580156107e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108079190611a84565b5060008773ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016108439190611f08565b60206040518083038186803b15801561085b57600080fd5b505afa15801561086f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108939190611bb2565b9050600081111561094f578773ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b81526004016108fb929190611f83565b602060405180830381600087803b15801561091557600080fd5b505af1158015610929573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094d9190611a84565b505b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b73828888886040516109b0939291906121b8565b60405180910390a3819250505095945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610a56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4d906120b8565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b60ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610aad6112a9565b610ab76000611327565b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000341415610b1e576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1634604051610b6690611ef3565b60006040518083038185875af1925050503d8060008114610ba3576040519150601f19603f3d011682016040523d82523d6000602084013e610ba8565b606091505b50509050600015158115151415610beb576040517f81063e5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fa93afc57f3be4641cf20c7165d11856f3b46dd376108e5fffb06f73f2b2a6d58848434604051610c1e93929190611fc7565b60405180910390a150505050565b8060c960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060019054906101000a900460ff16159050808015610ca15750600160008054906101000a900460ff1660ff16105b80610cce5750610cb0306113ed565b158015610ccd5750600160008054906101000a900460ff1660ff16145b5b610d0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d04906120f8565b60405180910390fd5b60016000806101000a81548160ff021916908360ff1602179055508015610d4a576001600060016101000a81548160ff0219169083151502179055505b610d52611410565b610d5a611469565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610dc1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160ca60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015610e5b5760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610e52919061201b565b60405180910390a15b5050565b8173ffffffffffffffffffffffffffffffffffffffff166323b872dd3360c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401610ebe93929190611f23565b602060405180830381600087803b158015610ed857600080fd5b505af1158015610eec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f109190611a84565b508173ffffffffffffffffffffffffffffffffffffffff167f35fb30ed1b8e81eb91001dad742b13b1491a67c722e8c593a886a18500f7d9af858584604051610f5b93929190611fc7565b60405180910390a250505050565b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610f976112a9565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611007576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ffe90612058565b60405180910390fd5b61101081611327565b50565b60606000808573ffffffffffffffffffffffffffffffffffffffff16348686604051611040929190611ec3565b60006040518083038185875af1925050503d806000811461107d576040519150601f19603f3d011682016040523d82523d6000602084013e611082565b606091505b5091509150816110be576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80925050509392505050565b60006110f87f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b6114ba565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6111296112a9565b50565b6111587f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b6114c4565b60000160009054906101000a900460ff161561117c57611177836114ce565b6112a4565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156111c257600080fd5b505afa9250505080156111f357506040513d601f19601f820116820180604052508101906111f09190611ab1565b60015b611232576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122990612118565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114611297576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161128e906120d8565b60405180910390fd5b506112a3838383611587565b5b505050565b6112b16115b3565b73ffffffffffffffffffffffffffffffffffffffff166112cf610ab9565b73ffffffffffffffffffffffffffffffffffffffff1614611325576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161131c90612158565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff1661145f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161145690612198565b60405180910390fd5b6114676115bb565b565b600060019054906101000a900460ff166114b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114af90612198565b60405180910390fd5b565b6000819050919050565b6000819050919050565b6114d7816113ed565b611516576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161150d90612138565b60405180910390fd5b806115437f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b6114ba565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6115908361161c565b60008251118061159d5750805b156115ae576115ac838361166b565b505b505050565b600033905090565b600060019054906101000a900460ff1661160a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161160190612198565b60405180910390fd5b61161a6116156115b3565b611327565b565b611625816114ce565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b6060611690838360405180606001604052806027815260200161275060279139611698565b905092915050565b60606000808573ffffffffffffffffffffffffffffffffffffffff16856040516116c29190611edc565b600060405180830381855af49150503d80600081146116fd576040519150601f19603f3d011682016040523d82523d6000602084013e611702565b606091505b50915091506117138683838761171e565b925050509392505050565b606083156117815760008351141561177957611739856113ed565b611778576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161176f90612178565b60405180910390fd5b5b82905061178c565b61178b8383611794565b5b949350505050565b6000825111156117a75781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117db9190612036565b60405180910390fd5b60006117f76117f28461220f565b6121ea565b905082815260208101848484011115611813576118126123b7565b5b61181e848285612306565b509392505050565b600081359050611835816126f3565b92915050565b60008151905061184a8161270a565b92915050565b60008151905061185f81612721565b92915050565b60008083601f84011261187b5761187a6123ad565b5b8235905067ffffffffffffffff811115611898576118976123a8565b5b6020830191508360018202830111156118b4576118b36123b2565b5b9250929050565b600082601f8301126118d0576118cf6123ad565b5b81356118e08482602086016117e4565b91505092915050565b6000813590506118f881612738565b92915050565b60008151905061190d81612738565b92915050565b600060208284031215611929576119286123c1565b5b600061193784828501611826565b91505092915050565b60008060008060006080868803121561195c5761195b6123c1565b5b600061196a88828901611826565b955050602061197b88828901611826565b945050604061198c888289016118e9565b935050606086013567ffffffffffffffff8111156119ad576119ac6123bc565b5b6119b988828901611865565b92509250509295509295909350565b6000806000604084860312156119e1576119e06123c1565b5b60006119ef86828701611826565b935050602084013567ffffffffffffffff811115611a1057611a0f6123bc565b5b611a1c86828701611865565b92509250509250925092565b60008060408385031215611a3f57611a3e6123c1565b5b6000611a4d85828601611826565b925050602083013567ffffffffffffffff811115611a6e57611a6d6123bc565b5b611a7a858286016118bb565b9150509250929050565b600060208284031215611a9a57611a996123c1565b5b6000611aa88482850161183b565b91505092915050565b600060208284031215611ac757611ac66123c1565b5b6000611ad584828501611850565b91505092915050565b60008060008060608587031215611af857611af76123c1565b5b600085013567ffffffffffffffff811115611b1657611b156123bc565b5b611b2287828801611865565b94509450506020611b3587828801611826565b9250506040611b46878288016118e9565b91505092959194509250565b600080600060408486031215611b6b57611b6a6123c1565b5b600084013567ffffffffffffffff811115611b8957611b886123bc565b5b611b9586828701611865565b93509350506020611ba8868287016118e9565b9150509250925092565b600060208284031215611bc857611bc76123c1565b5b6000611bd6848285016118fe565b91505092915050565b611be881612283565b82525050565b611bf7816122a1565b82525050565b6000611c098385612256565b9350611c16838584612306565b611c1f836123c6565b840190509392505050565b6000611c368385612267565b9350611c43838584612306565b82840190509392505050565b6000611c5a82612240565b611c648185612256565b9350611c74818560208601612315565b611c7d816123c6565b840191505092915050565b6000611c9382612240565b611c9d8185612267565b9350611cad818560208601612315565b80840191505092915050565b611cc2816122e2565b82525050565b611cd1816122f4565b82525050565b6000611ce28261224b565b611cec8185612272565b9350611cfc818560208601612315565b611d05816123c6565b840191505092915050565b6000611d1d602683612272565b9150611d28826123d7565b604082019050919050565b6000611d40602c83612272565b9150611d4b82612426565b604082019050919050565b6000611d63602c83612272565b9150611d6e82612475565b604082019050919050565b6000611d86603883612272565b9150611d91826124c4565b604082019050919050565b6000611da9602983612272565b9150611db482612513565b604082019050919050565b6000611dcc602e83612272565b9150611dd782612562565b604082019050919050565b6000611def602e83612272565b9150611dfa826125b1565b604082019050919050565b6000611e12602d83612272565b9150611e1d82612600565b604082019050919050565b6000611e35602083612272565b9150611e408261264f565b602082019050919050565b6000611e58600083612267565b9150611e6382612678565b600082019050919050565b6000611e7b601d83612272565b9150611e868261267b565b602082019050919050565b6000611e9e602b83612272565b9150611ea9826126a4565b604082019050919050565b611ebd816122cb565b82525050565b6000611ed0828486611c2a565b91508190509392505050565b6000611ee88284611c88565b915081905092915050565b6000611efe82611e4b565b9150819050919050565b6000602082019050611f1d6000830184611bdf565b92915050565b6000606082019050611f386000830186611bdf565b611f456020830185611bdf565b611f526040830184611eb4565b949350505050565b6000604082019050611f6f6000830185611bdf565b611f7c6020830184611cb9565b9392505050565b6000604082019050611f986000830185611bdf565b611fa56020830184611eb4565b9392505050565b6000602082019050611fc16000830184611bee565b92915050565b60006040820190508181036000830152611fe2818587611bfd565b9050611ff16020830184611eb4565b949350505050565b600060208201905081810360008301526120138184611c4f565b905092915050565b60006020820190506120306000830184611cc8565b92915050565b600060208201905081810360008301526120508184611cd7565b905092915050565b6000602082019050818103600083015261207181611d10565b9050919050565b6000602082019050818103600083015261209181611d33565b9050919050565b600060208201905081810360008301526120b181611d56565b9050919050565b600060208201905081810360008301526120d181611d79565b9050919050565b600060208201905081810360008301526120f181611d9c565b9050919050565b6000602082019050818103600083015261211181611dbf565b9050919050565b6000602082019050818103600083015261213181611de2565b9050919050565b6000602082019050818103600083015261215181611e05565b9050919050565b6000602082019050818103600083015261217181611e28565b9050919050565b6000602082019050818103600083015261219181611e6e565b9050919050565b600060208201905081810360008301526121b181611e91565b9050919050565b60006040820190506121cd6000830186611eb4565b81810360208301526121e0818486611bfd565b9050949350505050565b60006121f4612205565b90506122008282612348565b919050565b6000604051905090565b600067ffffffffffffffff82111561222a57612229612379565b5b612233826123c6565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061228e826122ab565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006122ed826122cb565b9050919050565b60006122ff826122d5565b9050919050565b82818337600083830152505050565b60005b83811015612333578082015181840152602081019050612318565b83811115612342576000848401525b50505050565b612351826123c6565b810181811067ffffffffffffffff821117156123705761236f612379565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b6126fc81612283565b811461270757600080fd5b50565b61271381612295565b811461271e57600080fd5b50565b61272a816122a1565b811461273557600080fd5b50565b612741816122cb565b811461274c57600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220090cdf549c1ee64dd921c98a57bcd2f4bdf4ddcba5c2dcfc478208b0a2bd11f964736f6c63430008070033",
}

// GatewayEVMUpgradeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMUpgradeTestMetaData.ABI instead.
var GatewayEVMUpgradeTestABI = GatewayEVMUpgradeTestMetaData.ABI

// GatewayEVMUpgradeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMUpgradeTestMetaData.Bin instead.
var GatewayEVMUpgradeTestBin = GatewayEVMUpgradeTestMetaData.Bin

// DeployGatewayEVMUpgradeTest deploys a new Ethereum contract, binding an instance of GatewayEVMUpgradeTest to it.
func DeployGatewayEVMUpgradeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMUpgradeTest, error) {
	parsed, err := GatewayEVMUpgradeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMUpgradeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMUpgradeTest{GatewayEVMUpgradeTestCaller: GatewayEVMUpgradeTestCaller{contract: contract}, GatewayEVMUpgradeTestTransactor: GatewayEVMUpgradeTestTransactor{contract: contract}, GatewayEVMUpgradeTestFilterer: GatewayEVMUpgradeTestFilterer{contract: contract}}, nil
}

// GatewayEVMUpgradeTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMUpgradeTest struct {
	GatewayEVMUpgradeTestCaller     // Read-only binding to the contract
	GatewayEVMUpgradeTestTransactor // Write-only binding to the contract
	GatewayEVMUpgradeTestFilterer   // Log filterer for contract events
}

// GatewayEVMUpgradeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMUpgradeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMUpgradeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMUpgradeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMUpgradeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMUpgradeTestSession struct {
	Contract     *GatewayEVMUpgradeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayEVMUpgradeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMUpgradeTestCallerSession struct {
	Contract *GatewayEVMUpgradeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// GatewayEVMUpgradeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMUpgradeTestTransactorSession struct {
	Contract     *GatewayEVMUpgradeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// GatewayEVMUpgradeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestRaw struct {
	Contract *GatewayEVMUpgradeTest // Generic contract binding to access the raw methods on
}

// GatewayEVMUpgradeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestCallerRaw struct {
	Contract *GatewayEVMUpgradeTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMUpgradeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestTransactorRaw struct {
	Contract *GatewayEVMUpgradeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMUpgradeTest creates a new instance of GatewayEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayEVMUpgradeTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMUpgradeTest, error) {
	contract, err := bindGatewayEVMUpgradeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTest{GatewayEVMUpgradeTestCaller: GatewayEVMUpgradeTestCaller{contract: contract}, GatewayEVMUpgradeTestTransactor: GatewayEVMUpgradeTestTransactor{contract: contract}, GatewayEVMUpgradeTestFilterer: GatewayEVMUpgradeTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMUpgradeTestCaller creates a new read-only instance of GatewayEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayEVMUpgradeTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMUpgradeTestCaller, error) {
	contract, err := bindGatewayEVMUpgradeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestCaller{contract: contract}, nil
}

// NewGatewayEVMUpgradeTestTransactor creates a new write-only instance of GatewayEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayEVMUpgradeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMUpgradeTestTransactor, error) {
	contract, err := bindGatewayEVMUpgradeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestTransactor{contract: contract}, nil
}

// NewGatewayEVMUpgradeTestFilterer creates a new log filterer instance of GatewayEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayEVMUpgradeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMUpgradeTestFilterer, error) {
	contract, err := bindGatewayEVMUpgradeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestFilterer{contract: contract}, nil
}

// bindGatewayEVMUpgradeTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMUpgradeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMUpgradeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMUpgradeTest.Contract.GatewayEVMUpgradeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.GatewayEVMUpgradeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.GatewayEVMUpgradeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMUpgradeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.contract.Transact(opts, method, params...)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) Custody(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "custody")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Custody() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.Custody(&_GatewayEVMUpgradeTest.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) Custody() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.Custody(&_GatewayEVMUpgradeTest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Owner() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.Owner(&_GatewayEVMUpgradeTest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) Owner() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.Owner(&_GatewayEVMUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.ProxiableUUID(&_GatewayEVMUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.ProxiableUUID(&_GatewayEVMUpgradeTest.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) TssAddress() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.TssAddress(&_GatewayEVMUpgradeTest.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) TssAddress() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.TssAddress(&_GatewayEVMUpgradeTest.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Execute(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "execute", destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Execute(&_GatewayEVMUpgradeTest.TransactOpts, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Execute(&_GatewayEVMUpgradeTest.TransactOpts, destination, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) ExecuteWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "executeWithERC20", token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.ExecuteWithERC20(&_GatewayEVMUpgradeTest.TransactOpts, token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.ExecuteWithERC20(&_GatewayEVMUpgradeTest.TransactOpts, token, to, amount, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tssAddress) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Initialize(opts *bind.TransactOpts, _tssAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "initialize", _tssAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tssAddress) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Initialize(_tssAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Initialize(&_GatewayEVMUpgradeTest.TransactOpts, _tssAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tssAddress) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Initialize(_tssAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Initialize(&_GatewayEVMUpgradeTest.TransactOpts, _tssAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.RenounceOwnership(&_GatewayEVMUpgradeTest.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.RenounceOwnership(&_GatewayEVMUpgradeTest.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x9372c4ab.
//
// Solidity: function send(bytes recipient, uint256 amount) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Send(opts *bind.TransactOpts, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "send", recipient, amount)
}

// Send is a paid mutator transaction binding the contract method 0x9372c4ab.
//
// Solidity: function send(bytes recipient, uint256 amount) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Send(recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Send(&_GatewayEVMUpgradeTest.TransactOpts, recipient, amount)
}

// Send is a paid mutator transaction binding the contract method 0x9372c4ab.
//
// Solidity: function send(bytes recipient, uint256 amount) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Send(recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Send(&_GatewayEVMUpgradeTest.TransactOpts, recipient, amount)
}

// SendERC20 is a paid mutator transaction binding the contract method 0xcb0271ed.
//
// Solidity: function sendERC20(bytes recipient, address token, uint256 amount) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) SendERC20(opts *bind.TransactOpts, recipient []byte, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "sendERC20", recipient, token, amount)
}

// SendERC20 is a paid mutator transaction binding the contract method 0xcb0271ed.
//
// Solidity: function sendERC20(bytes recipient, address token, uint256 amount) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) SendERC20(recipient []byte, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.SendERC20(&_GatewayEVMUpgradeTest.TransactOpts, recipient, token, amount)
}

// SendERC20 is a paid mutator transaction binding the contract method 0xcb0271ed.
//
// Solidity: function sendERC20(bytes recipient, address token, uint256 amount) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) SendERC20(recipient []byte, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.SendERC20(&_GatewayEVMUpgradeTest.TransactOpts, recipient, token, amount)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) SetCustody(opts *bind.TransactOpts, _custody common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "setCustody", _custody)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) SetCustody(_custody common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.SetCustody(&_GatewayEVMUpgradeTest.TransactOpts, _custody)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) SetCustody(_custody common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.SetCustody(&_GatewayEVMUpgradeTest.TransactOpts, _custody)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.TransferOwnership(&_GatewayEVMUpgradeTest.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.TransferOwnership(&_GatewayEVMUpgradeTest.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.UpgradeTo(&_GatewayEVMUpgradeTest.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.UpgradeTo(&_GatewayEVMUpgradeTest.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.UpgradeToAndCall(&_GatewayEVMUpgradeTest.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.UpgradeToAndCall(&_GatewayEVMUpgradeTest.TransactOpts, newImplementation, data)
}

// GatewayEVMUpgradeTestAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestAdminChangedIterator struct {
	Event *GatewayEVMUpgradeTestAdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestAdminChanged)
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
		it.Event = new(GatewayEVMUpgradeTestAdminChanged)
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
func (it *GatewayEVMUpgradeTestAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestAdminChanged represents a AdminChanged event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*GatewayEVMUpgradeTestAdminChangedIterator, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestAdminChangedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestAdminChanged) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestAdminChanged)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseAdminChanged(log types.Log) (*GatewayEVMUpgradeTestAdminChanged, error) {
	event := new(GatewayEVMUpgradeTestAdminChanged)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestBeaconUpgradedIterator struct {
	Event *GatewayEVMUpgradeTestBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestBeaconUpgraded)
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
		it.Event = new(GatewayEVMUpgradeTestBeaconUpgraded)
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
func (it *GatewayEVMUpgradeTestBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestBeaconUpgraded represents a BeaconUpgraded event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*GatewayEVMUpgradeTestBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestBeaconUpgradedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestBeaconUpgraded)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseBeaconUpgraded(log types.Log) (*GatewayEVMUpgradeTestBeaconUpgraded, error) {
	event := new(GatewayEVMUpgradeTestBeaconUpgraded)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestExecutedV2Iterator is returned from FilterExecutedV2 and is used to iterate over the raw logs and unpacked data for ExecutedV2 events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedV2Iterator struct {
	Event *GatewayEVMUpgradeTestExecutedV2 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestExecutedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestExecutedV2)
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
		it.Event = new(GatewayEVMUpgradeTestExecutedV2)
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
func (it *GatewayEVMUpgradeTestExecutedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestExecutedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestExecutedV2 represents a ExecutedV2 event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedV2 struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutedV2 is a free log retrieval operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterExecutedV2(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMUpgradeTestExecutedV2Iterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "ExecutedV2", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestExecutedV2Iterator{contract: _GatewayEVMUpgradeTest.contract, event: "ExecutedV2", logs: logs, sub: sub}, nil
}

// WatchExecutedV2 is a free log subscription operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchExecutedV2(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestExecutedV2, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "ExecutedV2", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestExecutedV2)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "ExecutedV2", log); err != nil {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseExecutedV2(log types.Log) (*GatewayEVMUpgradeTestExecutedV2, error) {
	event := new(GatewayEVMUpgradeTestExecutedV2)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "ExecutedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMUpgradeTestExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestExecutedWithERC20)
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
		it.Event = new(GatewayEVMUpgradeTestExecutedWithERC20)
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
func (it *GatewayEVMUpgradeTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMUpgradeTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestExecutedWithERC20Iterator{contract: _GatewayEVMUpgradeTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestExecutedWithERC20)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMUpgradeTestExecutedWithERC20, error) {
	event := new(GatewayEVMUpgradeTestExecutedWithERC20)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestInitializedIterator struct {
	Event *GatewayEVMUpgradeTestInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestInitialized)
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
		it.Event = new(GatewayEVMUpgradeTestInitialized)
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
func (it *GatewayEVMUpgradeTestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestInitialized represents a Initialized event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayEVMUpgradeTestInitializedIterator, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestInitializedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestInitialized)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseInitialized(log types.Log) (*GatewayEVMUpgradeTestInitialized, error) {
	event := new(GatewayEVMUpgradeTestInitialized)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestOwnershipTransferredIterator struct {
	Event *GatewayEVMUpgradeTestOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestOwnershipTransferred)
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
		it.Event = new(GatewayEVMUpgradeTestOwnershipTransferred)
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
func (it *GatewayEVMUpgradeTestOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestOwnershipTransferred represents a OwnershipTransferred event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayEVMUpgradeTestOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestOwnershipTransferredIterator{contract: _GatewayEVMUpgradeTest.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestOwnershipTransferred)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayEVMUpgradeTestOwnershipTransferred, error) {
	event := new(GatewayEVMUpgradeTestOwnershipTransferred)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestSendIterator is returned from FilterSend and is used to iterate over the raw logs and unpacked data for Send events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestSendIterator struct {
	Event *GatewayEVMUpgradeTestSend // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestSend)
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
		it.Event = new(GatewayEVMUpgradeTestSend)
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
func (it *GatewayEVMUpgradeTestSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestSend represents a Send event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestSend struct {
	Recipient []byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSend is a free log retrieval operation binding the contract event 0xa93afc57f3be4641cf20c7165d11856f3b46dd376108e5fffb06f73f2b2a6d58.
//
// Solidity: event Send(bytes recipient, uint256 amount)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterSend(opts *bind.FilterOpts) (*GatewayEVMUpgradeTestSendIterator, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestSendIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Send", logs: logs, sub: sub}, nil
}

// WatchSend is a free log subscription operation binding the contract event 0xa93afc57f3be4641cf20c7165d11856f3b46dd376108e5fffb06f73f2b2a6d58.
//
// Solidity: event Send(bytes recipient, uint256 amount)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchSend(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestSend) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestSend)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Send", log); err != nil {
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

// ParseSend is a log parse operation binding the contract event 0xa93afc57f3be4641cf20c7165d11856f3b46dd376108e5fffb06f73f2b2a6d58.
//
// Solidity: event Send(bytes recipient, uint256 amount)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseSend(log types.Log) (*GatewayEVMUpgradeTestSend, error) {
	event := new(GatewayEVMUpgradeTestSend)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Send", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestSendERC20Iterator is returned from FilterSendERC20 and is used to iterate over the raw logs and unpacked data for SendERC20 events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestSendERC20Iterator struct {
	Event *GatewayEVMUpgradeTestSendERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestSendERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestSendERC20)
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
		it.Event = new(GatewayEVMUpgradeTestSendERC20)
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
func (it *GatewayEVMUpgradeTestSendERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestSendERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestSendERC20 represents a SendERC20 event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestSendERC20 struct {
	Recipient []byte
	Asset     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSendERC20 is a free log retrieval operation binding the contract event 0x35fb30ed1b8e81eb91001dad742b13b1491a67c722e8c593a886a18500f7d9af.
//
// Solidity: event SendERC20(bytes recipient, address indexed asset, uint256 amount)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterSendERC20(opts *bind.FilterOpts, asset []common.Address) (*GatewayEVMUpgradeTestSendERC20Iterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "SendERC20", assetRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestSendERC20Iterator{contract: _GatewayEVMUpgradeTest.contract, event: "SendERC20", logs: logs, sub: sub}, nil
}

// WatchSendERC20 is a free log subscription operation binding the contract event 0x35fb30ed1b8e81eb91001dad742b13b1491a67c722e8c593a886a18500f7d9af.
//
// Solidity: event SendERC20(bytes recipient, address indexed asset, uint256 amount)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchSendERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestSendERC20, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "SendERC20", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestSendERC20)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "SendERC20", log); err != nil {
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

// ParseSendERC20 is a log parse operation binding the contract event 0x35fb30ed1b8e81eb91001dad742b13b1491a67c722e8c593a886a18500f7d9af.
//
// Solidity: event SendERC20(bytes recipient, address indexed asset, uint256 amount)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseSendERC20(log types.Log) (*GatewayEVMUpgradeTestSendERC20, error) {
	event := new(GatewayEVMUpgradeTestSendERC20)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "SendERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestUpgradedIterator struct {
	Event *GatewayEVMUpgradeTestUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradeTestUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestUpgraded)
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
		it.Event = new(GatewayEVMUpgradeTestUpgraded)
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
func (it *GatewayEVMUpgradeTestUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestUpgraded represents a Upgraded event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayEVMUpgradeTestUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestUpgradedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestUpgraded)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseUpgraded(log types.Log) (*GatewayEVMUpgradeTestUpgraded, error) {
	event := new(GatewayEVMUpgradeTestUpgraded)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
