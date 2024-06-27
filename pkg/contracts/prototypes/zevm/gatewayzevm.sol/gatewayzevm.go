// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayzevm

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

// GatewayZEVMMetaData contains all meta data concerning the GatewayZEVM contract.
var GatewayZEVMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientZRC20Amount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Call\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b8152503480156200004757600080fd5b50620000586200005e60201b60201c565b62000208565b600060019054906101000a900460ff1615620000b1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a8906200015c565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff1614620001225760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200011991906200017e565b60405180910390a15b565b6000620001336027836200019b565b91506200014082620001b9565b604082019050919050565b6200015681620001ac565b82525050565b60006020820190508181036000830152620001778162000124565b9050919050565b60006020820190506200019560008301846200014b565b92915050565b600082825260208201905092915050565b600060ff82169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60805160601c611f036200024360003960008181610346015281816103d5015281816104cf0152818161055e015261060e0152611f036000f3fe6080604052600436106100915760003560e01c8063715018a611610059578063715018a6146101585780637993c1e01461016f5780638129fc1c146101985780638da5cb5b146101af578063f2fde38b146101da57610091565b80630ac7c44c14610096578063135390f9146100bf5780633659cfe6146100e85780634f1ef2861461011157806352d1902d1461012d575b600080fd5b3480156100a257600080fd5b506100bd60048036038101906100b891906112be565b610203565b005b3480156100cb57600080fd5b506100e660048036038101906100e1919061133a565b61026e565b005b3480156100f457600080fd5b5061010f600480360381019061010a91906111db565b610344565b005b61012b60048036038101906101269190611208565b6104cd565b005b34801561013957600080fd5b5061014261060a565b60405161014f919061170c565b60405180910390f35b34801561016457600080fd5b5061016d6106c3565b005b34801561017b57600080fd5b50610196600480360381019061019191906113a9565b6106d7565b005b3480156101a457600080fd5b506101ad6107b3565b005b3480156101bb57600080fd5b506101c46108f9565b6040516101d191906116f1565b60405180910390f35b3480156101e657600080fd5b5061020160048036038101906101fc91906111db565b610923565b005b8260405161021191906116da565b60405180910390203373ffffffffffffffffffffffffffffffffffffffff167f2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f8484604051610261929190611727565b60405180910390a3505050565b60008173ffffffffffffffffffffffffffffffffffffffff1663c701262685856040518363ffffffff1660e01b81526004016102ab92919061174b565b602060405180830381600087803b1580156102c557600080fd5b505af11580156102d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102fd9190611264565b905060001515811515141561033e576040517f27fcd9d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614156103d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ca90611816565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166104126109a7565b73ffffffffffffffffffffffffffffffffffffffff1614610468576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045f90611836565b60405180910390fd5b610471816109fe565b6104ca81600067ffffffffffffffff8111156104905761048f611ad3565b5b6040519080825280601f01601f1916602001820160405280156104c25781602001600182028036833780820191505090505b506000610a09565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16141561055c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055390611816565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661059b6109a7565b73ffffffffffffffffffffffffffffffffffffffff16146105f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e890611836565b60405180910390fd5b6105fa826109fe565b61060682826001610a09565b5050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161461069a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069190611856565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b6106cb610b86565b6106d56000610c04565b565b3073ffffffffffffffffffffffffffffffffffffffff1663135390f98686866040518463ffffffff1660e01b81526004016107149392919061177b565b600060405180830381600087803b15801561072e57600080fd5b505af1158015610742573d6000803e3d6000fd5b505050508460405161075491906116da565b60405180910390203373ffffffffffffffffffffffffffffffffffffffff167f2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f84846040516107a4929190611727565b60405180910390a35050505050565b60008060019054906101000a900460ff161590508080156107e45750600160008054906101000a900460ff1660ff16105b8061081157506107f330610cca565b1580156108105750600160008054906101000a900460ff1660ff16145b5b610850576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084790611896565b60405180910390fd5b60016000806101000a81548160ff021916908360ff160217905550801561088d576001600060016101000a81548160ff0219169083151502179055505b610895610ced565b61089d610d46565b80156108f65760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516108ed91906117b9565b60405180910390a15b50565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61092b610b86565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561099b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610992906117f6565b60405180910390fd5b6109a481610c04565b50565b60006109d57f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b610d97565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610a06610b86565b50565b610a357f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b610da1565b60000160009054906101000a900460ff1615610a5957610a5483610dab565b610b81565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a9f57600080fd5b505afa925050508015610ad057506040513d601f19601f82011682018060405250810190610acd9190611291565b60015b610b0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b06906118b6565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114610b74576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6b90611876565b60405180910390fd5b50610b80838383610e64565b5b505050565b610b8e610e90565b73ffffffffffffffffffffffffffffffffffffffff16610bac6108f9565b73ffffffffffffffffffffffffffffffffffffffff1614610c02576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bf9906118f6565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff16610d3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3390611936565b60405180910390fd5b610d44610e98565b565b600060019054906101000a900460ff16610d95576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d8c90611936565b60405180910390fd5b565b6000819050919050565b6000819050919050565b610db481610cca565b610df3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dea906118d6565b60405180910390fd5b80610e207f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b610d97565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b610e6d83610ef9565b600082511180610e7a5750805b15610e8b57610e898383610f48565b505b505050565b600033905090565b600060019054906101000a900460ff16610ee7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ede90611936565b60405180910390fd5b610ef7610ef2610e90565b610c04565b565b610f0281610dab565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b6060610f6d8383604051806060016040528060278152602001611ea760279139610f75565b905092915050565b60606000808573ffffffffffffffffffffffffffffffffffffffff1685604051610f9f91906116da565b600060405180830381855af49150503d8060008114610fda576040519150601f19603f3d011682016040523d82523d6000602084013e610fdf565b606091505b5091509150610ff086838387610ffb565b925050509392505050565b6060831561105e576000835114156110565761101685610cca565b611055576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161104c90611916565b60405180910390fd5b5b829050611069565b6110688383611071565b5b949350505050565b6000825111156110845781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110b891906117d4565b60405180910390fd5b60006110d46110cf8461197b565b611956565b9050828152602081018484840111156110f0576110ef611b11565b5b6110fb848285611a60565b509392505050565b60008135905061111281611e4a565b92915050565b60008151905061112781611e61565b92915050565b60008151905061113c81611e78565b92915050565b60008083601f84011261115857611157611b07565b5b8235905067ffffffffffffffff81111561117557611174611b02565b5b60208301915083600182028301111561119157611190611b0c565b5b9250929050565b600082601f8301126111ad576111ac611b07565b5b81356111bd8482602086016110c1565b91505092915050565b6000813590506111d581611e8f565b92915050565b6000602082840312156111f1576111f0611b1b565b5b60006111ff84828501611103565b91505092915050565b6000806040838503121561121f5761121e611b1b565b5b600061122d85828601611103565b925050602083013567ffffffffffffffff81111561124e5761124d611b16565b5b61125a85828601611198565b9150509250929050565b60006020828403121561127a57611279611b1b565b5b600061128884828501611118565b91505092915050565b6000602082840312156112a7576112a6611b1b565b5b60006112b58482850161112d565b91505092915050565b6000806000604084860312156112d7576112d6611b1b565b5b600084013567ffffffffffffffff8111156112f5576112f4611b16565b5b61130186828701611198565b935050602084013567ffffffffffffffff81111561132257611321611b16565b5b61132e86828701611142565b92509250509250925092565b60008060006060848603121561135357611352611b1b565b5b600084013567ffffffffffffffff81111561137157611370611b16565b5b61137d86828701611198565b935050602061138e868287016111c6565b925050604061139f86828701611103565b9150509250925092565b6000806000806000608086880312156113c5576113c4611b1b565b5b600086013567ffffffffffffffff8111156113e3576113e2611b16565b5b6113ef88828901611198565b9550506020611400888289016111c6565b945050604061141188828901611103565b935050606086013567ffffffffffffffff81111561143257611431611b16565b5b61143e88828901611142565b92509250509295509295909350565b611456816119ef565b82525050565b61146581611a0d565b82525050565b600061147783856119c2565b9350611484838584611a60565b61148d83611b20565b840190509392505050565b60006114a3826119ac565b6114ad81856119c2565b93506114bd818560208601611a6f565b6114c681611b20565b840191505092915050565b60006114dc826119ac565b6114e681856119d3565b93506114f6818560208601611a6f565b80840191505092915050565b61150b81611a4e565b82525050565b600061151c826119b7565b61152681856119de565b9350611536818560208601611a6f565b61153f81611b20565b840191505092915050565b60006115576026836119de565b915061156282611b31565b604082019050919050565b600061157a602c836119de565b915061158582611b80565b604082019050919050565b600061159d602c836119de565b91506115a882611bcf565b604082019050919050565b60006115c06038836119de565b91506115cb82611c1e565b604082019050919050565b60006115e36029836119de565b91506115ee82611c6d565b604082019050919050565b6000611606602e836119de565b915061161182611cbc565b604082019050919050565b6000611629602e836119de565b915061163482611d0b565b604082019050919050565b600061164c602d836119de565b915061165782611d5a565b604082019050919050565b600061166f6020836119de565b915061167a82611da9565b602082019050919050565b6000611692601d836119de565b915061169d82611dd2565b602082019050919050565b60006116b5602b836119de565b91506116c082611dfb565b604082019050919050565b6116d481611a37565b82525050565b60006116e682846114d1565b915081905092915050565b6000602082019050611706600083018461144d565b92915050565b6000602082019050611721600083018461145c565b92915050565b6000602082019050818103600083015261174281848661146b565b90509392505050565b600060408201905081810360008301526117658185611498565b905061177460208301846116cb565b9392505050565b600060608201905081810360008301526117958186611498565b90506117a460208301856116cb565b6117b1604083018461144d565b949350505050565b60006020820190506117ce6000830184611502565b92915050565b600060208201905081810360008301526117ee8184611511565b905092915050565b6000602082019050818103600083015261180f8161154a565b9050919050565b6000602082019050818103600083015261182f8161156d565b9050919050565b6000602082019050818103600083015261184f81611590565b9050919050565b6000602082019050818103600083015261186f816115b3565b9050919050565b6000602082019050818103600083015261188f816115d6565b9050919050565b600060208201905081810360008301526118af816115f9565b9050919050565b600060208201905081810360008301526118cf8161161c565b9050919050565b600060208201905081810360008301526118ef8161163f565b9050919050565b6000602082019050818103600083015261190f81611662565b9050919050565b6000602082019050818103600083015261192f81611685565b9050919050565b6000602082019050818103600083015261194f816116a8565b9050919050565b6000611960611971565b905061196c8282611aa2565b919050565b6000604051905090565b600067ffffffffffffffff82111561199657611995611ad3565b5b61199f82611b20565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006119fa82611a17565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611a5982611a41565b9050919050565b82818337600083830152505050565b60005b83811015611a8d578082015181840152602081019050611a72565b83811115611a9c576000848401525b50505050565b611aab82611b20565b810181811067ffffffffffffffff82111715611aca57611ac9611ad3565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b611e53816119ef565b8114611e5e57600080fd5b50565b611e6a81611a01565b8114611e7557600080fd5b50565b611e8181611a0d565b8114611e8c57600080fd5b50565b611e9881611a37565b8114611ea357600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220bb6fa20a09f26183befe47ca247b509b8c32da9568301b8b24a02aab55dd676a64736f6c63430008070033",
}

// GatewayZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayZEVMMetaData.ABI instead.
var GatewayZEVMABI = GatewayZEVMMetaData.ABI

// GatewayZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayZEVMMetaData.Bin instead.
var GatewayZEVMBin = GatewayZEVMMetaData.Bin

// DeployGatewayZEVM deploys a new Ethereum contract, binding an instance of GatewayZEVM to it.
func DeployGatewayZEVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayZEVM, error) {
	parsed, err := GatewayZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayZEVMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayZEVM{GatewayZEVMCaller: GatewayZEVMCaller{contract: contract}, GatewayZEVMTransactor: GatewayZEVMTransactor{contract: contract}, GatewayZEVMFilterer: GatewayZEVMFilterer{contract: contract}}, nil
}

// GatewayZEVM is an auto generated Go binding around an Ethereum contract.
type GatewayZEVM struct {
	GatewayZEVMCaller     // Read-only binding to the contract
	GatewayZEVMTransactor // Write-only binding to the contract
	GatewayZEVMFilterer   // Log filterer for contract events
}

// GatewayZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayZEVMSession struct {
	Contract     *GatewayZEVM      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayZEVMCallerSession struct {
	Contract *GatewayZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GatewayZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayZEVMTransactorSession struct {
	Contract     *GatewayZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayZEVMRaw struct {
	Contract *GatewayZEVM // Generic contract binding to access the raw methods on
}

// GatewayZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayZEVMCallerRaw struct {
	Contract *GatewayZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayZEVMTransactorRaw struct {
	Contract *GatewayZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayZEVM creates a new instance of GatewayZEVM, bound to a specific deployed contract.
func NewGatewayZEVM(address common.Address, backend bind.ContractBackend) (*GatewayZEVM, error) {
	contract, err := bindGatewayZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVM{GatewayZEVMCaller: GatewayZEVMCaller{contract: contract}, GatewayZEVMTransactor: GatewayZEVMTransactor{contract: contract}, GatewayZEVMFilterer: GatewayZEVMFilterer{contract: contract}}, nil
}

// NewGatewayZEVMCaller creates a new read-only instance of GatewayZEVM, bound to a specific deployed contract.
func NewGatewayZEVMCaller(address common.Address, caller bind.ContractCaller) (*GatewayZEVMCaller, error) {
	contract, err := bindGatewayZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMCaller{contract: contract}, nil
}

// NewGatewayZEVMTransactor creates a new write-only instance of GatewayZEVM, bound to a specific deployed contract.
func NewGatewayZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayZEVMTransactor, error) {
	contract, err := bindGatewayZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMTransactor{contract: contract}, nil
}

// NewGatewayZEVMFilterer creates a new log filterer instance of GatewayZEVM, bound to a specific deployed contract.
func NewGatewayZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayZEVMFilterer, error) {
	contract, err := bindGatewayZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMFilterer{contract: contract}, nil
}

// bindGatewayZEVM binds a generic wrapper to an already deployed contract.
func bindGatewayZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVM *GatewayZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVM.Contract.GatewayZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVM *GatewayZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.GatewayZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVM *GatewayZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.GatewayZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVM *GatewayZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVM *GatewayZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVM *GatewayZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayZEVM *GatewayZEVMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayZEVM *GatewayZEVMSession) Owner() (common.Address, error) {
	return _GatewayZEVM.Contract.Owner(&_GatewayZEVM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayZEVM *GatewayZEVMCallerSession) Owner() (common.Address, error) {
	return _GatewayZEVM.Contract.Owner(&_GatewayZEVM.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayZEVM.Contract.ProxiableUUID(&_GatewayZEVM.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayZEVM.Contract.ProxiableUUID(&_GatewayZEVM.CallOpts)
}

// Call is a paid mutator transaction binding the contract method 0x0ac7c44c.
//
// Solidity: function call(bytes receiver, bytes payload) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Call(opts *bind.TransactOpts, receiver []byte, payload []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "call", receiver, payload)
}

// Call is a paid mutator transaction binding the contract method 0x0ac7c44c.
//
// Solidity: function call(bytes receiver, bytes payload) returns()
func (_GatewayZEVM *GatewayZEVMSession) Call(receiver []byte, payload []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Call(&_GatewayZEVM.TransactOpts, receiver, payload)
}

// Call is a paid mutator transaction binding the contract method 0x0ac7c44c.
//
// Solidity: function call(bytes receiver, bytes payload) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Call(receiver []byte, payload []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Call(&_GatewayZEVM.TransactOpts, receiver, payload)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GatewayZEVM *GatewayZEVMSession) Initialize() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Initialize(&_GatewayZEVM.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Initialize() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Initialize(&_GatewayZEVM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayZEVM *GatewayZEVMTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayZEVM *GatewayZEVMSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.RenounceOwnership(&_GatewayZEVM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.RenounceOwnership(&_GatewayZEVM.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayZEVM *GatewayZEVMSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.TransferOwnership(&_GatewayZEVM.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.TransferOwnership(&_GatewayZEVM.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GatewayZEVM *GatewayZEVMSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.UpgradeTo(&_GatewayZEVM.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.UpgradeTo(&_GatewayZEVM.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVM *GatewayZEVMTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVM *GatewayZEVMSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.UpgradeToAndCall(&_GatewayZEVM.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.UpgradeToAndCall(&_GatewayZEVM.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x135390f9.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Withdraw(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdraw", receiver, amount, zrc20)
}

// Withdraw is a paid mutator transaction binding the contract method 0x135390f9.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20) returns()
func (_GatewayZEVM *GatewayZEVMSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20)
}

// Withdraw is a paid mutator transaction binding the contract method 0x135390f9.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x7993c1e0.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) WithdrawAndCall(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdrawAndCall", receiver, amount, zrc20, message)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x7993c1e0.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMSession) WithdrawAndCall(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.WithdrawAndCall(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, message)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x7993c1e0.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) WithdrawAndCall(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.WithdrawAndCall(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, message)
}

// GatewayZEVMAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the GatewayZEVM contract.
type GatewayZEVMAdminChangedIterator struct {
	Event *GatewayZEVMAdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMAdminChanged)
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
		it.Event = new(GatewayZEVMAdminChanged)
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
func (it *GatewayZEVMAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMAdminChanged represents a AdminChanged event raised by the GatewayZEVM contract.
type GatewayZEVMAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*GatewayZEVMAdminChangedIterator, error) {

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMAdminChangedIterator{contract: _GatewayZEVM.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayZEVMAdminChanged) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMAdminChanged)
				if err := _GatewayZEVM.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseAdminChanged(log types.Log) (*GatewayZEVMAdminChanged, error) {
	event := new(GatewayZEVMAdminChanged)
	if err := _GatewayZEVM.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the GatewayZEVM contract.
type GatewayZEVMBeaconUpgradedIterator struct {
	Event *GatewayZEVMBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMBeaconUpgraded)
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
		it.Event = new(GatewayZEVMBeaconUpgraded)
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
func (it *GatewayZEVMBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMBeaconUpgraded represents a BeaconUpgraded event raised by the GatewayZEVM contract.
type GatewayZEVMBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*GatewayZEVMBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMBeaconUpgradedIterator{contract: _GatewayZEVM.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayZEVMBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMBeaconUpgraded)
				if err := _GatewayZEVM.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseBeaconUpgraded(log types.Log) (*GatewayZEVMBeaconUpgraded, error) {
	event := new(GatewayZEVMBeaconUpgraded)
	if err := _GatewayZEVM.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMCallIterator is returned from FilterCall and is used to iterate over the raw logs and unpacked data for Call events raised by the GatewayZEVM contract.
type GatewayZEVMCallIterator struct {
	Event *GatewayZEVMCall // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMCall)
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
		it.Event = new(GatewayZEVMCall)
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
func (it *GatewayZEVMCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMCall represents a Call event raised by the GatewayZEVM contract.
type GatewayZEVMCall struct {
	Sender   common.Address
	Receiver common.Hash
	Message  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCall is a free log retrieval operation binding the contract event 0x2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f.
//
// Solidity: event Call(address indexed sender, bytes indexed receiver, bytes message)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterCall(opts *bind.FilterOpts, sender []common.Address, receiver [][]byte) (*GatewayZEVMCallIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Call", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMCallIterator{contract: _GatewayZEVM.contract, event: "Call", logs: logs, sub: sub}, nil
}

// WatchCall is a free log subscription operation binding the contract event 0x2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f.
//
// Solidity: event Call(address indexed sender, bytes indexed receiver, bytes message)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchCall(opts *bind.WatchOpts, sink chan<- *GatewayZEVMCall, sender []common.Address, receiver [][]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Call", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMCall)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Call", log); err != nil {
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

// ParseCall is a log parse operation binding the contract event 0x2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f.
//
// Solidity: event Call(address indexed sender, bytes indexed receiver, bytes message)
func (_GatewayZEVM *GatewayZEVMFilterer) ParseCall(log types.Log) (*GatewayZEVMCall, error) {
	event := new(GatewayZEVMCall)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Call", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayZEVM contract.
type GatewayZEVMInitializedIterator struct {
	Event *GatewayZEVMInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInitialized)
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
		it.Event = new(GatewayZEVMInitialized)
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
func (it *GatewayZEVMInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInitialized represents a Initialized event raised by the GatewayZEVM contract.
type GatewayZEVMInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayZEVMInitializedIterator, error) {

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInitializedIterator{contract: _GatewayZEVM.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInitialized)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseInitialized(log types.Log) (*GatewayZEVMInitialized, error) {
	event := new(GatewayZEVMInitialized)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GatewayZEVM contract.
type GatewayZEVMOwnershipTransferredIterator struct {
	Event *GatewayZEVMOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOwnershipTransferred)
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
		it.Event = new(GatewayZEVMOwnershipTransferred)
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
func (it *GatewayZEVMOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOwnershipTransferred represents a OwnershipTransferred event raised by the GatewayZEVM contract.
type GatewayZEVMOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayZEVMOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOwnershipTransferredIterator{contract: _GatewayZEVM.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOwnershipTransferred)
				if err := _GatewayZEVM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayZEVMOwnershipTransferred, error) {
	event := new(GatewayZEVMOwnershipTransferred)
	if err := _GatewayZEVM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayZEVM contract.
type GatewayZEVMUpgradedIterator struct {
	Event *GatewayZEVMUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgraded)
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
		it.Event = new(GatewayZEVMUpgraded)
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
func (it *GatewayZEVMUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgraded represents a Upgraded event raised by the GatewayZEVM contract.
type GatewayZEVMUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayZEVMUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradedIterator{contract: _GatewayZEVM.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgraded)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseUpgraded(log types.Log) (*GatewayZEVMUpgraded, error) {
	event := new(GatewayZEVMUpgraded)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
