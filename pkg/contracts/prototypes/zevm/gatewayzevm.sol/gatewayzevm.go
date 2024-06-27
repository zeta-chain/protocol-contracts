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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GasFeeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientZRC20Amount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Call\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasfee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFlatFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b8152503480156200004757600080fd5b50620000586200005e60201b60201c565b62000208565b600060019054906101000a900460ff1615620000b1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a8906200015c565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff1614620001225760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200011991906200017e565b60405180910390a15b565b6000620001336027836200019b565b91506200014082620001b9565b604082019050919050565b6200015681620001ac565b82525050565b60006020820190508181036000830152620001778162000124565b9050919050565b60006020820190506200019560008301846200014b565b92915050565b600082825260208201905092915050565b600060ff82169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60805160601c612358620002436000396000818161038b0152818161041a0152818161052c015281816105bb015261066b01526123586000f3fe60806040526004361061009c5760003560e01c806352d1902d1161006457806352d1902d14610163578063715018a61461018e5780637993c1e0146101a55780638129fc1c146101ce5780638da5cb5b146101e5578063f2fde38b146102105761009c565b80630ac7c44c146100a1578063135390f9146100ca5780633659cfe6146100f35780633ce4a5bc1461011c5780634f1ef28614610147575b600080fd5b3480156100ad57600080fd5b506100c860048036038101906100c3919061161a565b610239565b005b3480156100d657600080fd5b506100f160048036038101906100ec9190611696565b6102a4565b005b3480156100ff57600080fd5b5061011a600480360381019061011591906114f7565b610389565b005b34801561012857600080fd5b50610131610512565b60405161013e9190611a9d565b60405180910390f35b610161600480360381019061015c9190611524565b61052a565b005b34801561016f57600080fd5b50610178610667565b6040516101859190611aef565b60405180910390f35b34801561019a57600080fd5b506101a3610720565b005b3480156101b157600080fd5b506101cc60048036038101906101c79190611705565b610734565b005b3480156101da57600080fd5b506101e361081f565b005b3480156101f157600080fd5b506101fa610965565b6040516102079190611a9d565b60405180910390f35b34801561021c57600080fd5b50610237600480360381019061023291906114f7565b61098f565b005b826040516102479190611a86565b60405180910390203373ffffffffffffffffffffffffffffffffffffffff167f2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f8484604051610297929190611b0a565b60405180910390a3505050565b60006102b08383610a13565b90503373ffffffffffffffffffffffffffffffffffffffff167f1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d8585848673ffffffffffffffffffffffffffffffffffffffff16634d8943bb6040518163ffffffff1660e01b815260040160206040518083038186803b15801561033357600080fd5b505afa158015610347573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036b91906117a9565b60405161037b9493929190611b91565b60405180910390a250505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610418576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040f90611c4d565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610457610c99565b73ffffffffffffffffffffffffffffffffffffffff16146104ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a490611c6d565b60405180910390fd5b6104b681610cf0565b61050f81600067ffffffffffffffff8111156104d5576104d4611f25565b5b6040519080825280601f01601f1916602001820160405280156105075781602001600182028036833780820191505090505b506000610cfb565b50565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614156105b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b090611c4d565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166105f8610c99565b73ffffffffffffffffffffffffffffffffffffffff161461064e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064590611c6d565b60405180910390fd5b61065782610cf0565b61066382826001610cfb565b5050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16146106f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ee90611c8d565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b610728610e78565b6107326000610ef6565b565b60006107408585610a13565b90503373ffffffffffffffffffffffffffffffffffffffff167f1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d8787848873ffffffffffffffffffffffffffffffffffffffff16634d8943bb6040518163ffffffff1660e01b815260040160206040518083038186803b1580156107c357600080fd5b505afa1580156107d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107fb91906117a9565b888860405161080f96959493929190611b2e565b60405180910390a2505050505050565b60008060019054906101000a900460ff161590508080156108505750600160008054906101000a900460ff1660ff16105b8061087d575061085f30610fbc565b15801561087c5750600160008054906101000a900460ff1660ff16145b5b6108bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b390611ccd565b60405180910390fd5b60016000806101000a81548160ff021916908360ff16021790555080156108f9576001600060016101000a81548160ff0219169083151502179055505b610901610fdf565b610909611038565b80156109625760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516109599190611bf0565b60405180910390a15b50565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610997610e78565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a07576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109fe90611c2d565b60405180910390fd5b610a1081610ef6565b50565b60008060008373ffffffffffffffffffffffffffffffffffffffff1663d9eeebed6040518163ffffffff1660e01b8152600401604080518083038186803b158015610a5d57600080fd5b505afa158015610a71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a959190611580565b915091508173ffffffffffffffffffffffffffffffffffffffff166323b872dd3373735b14bb79463307aacbed86daf3322b1e6226ab846040518463ffffffff1660e01b8152600401610aea93929190611ab8565b602060405180830381600087803b158015610b0457600080fd5b505af1158015610b18573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3c91906115c0565b610b72576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330886040518463ffffffff1660e01b8152600401610baf93929190611ab8565b602060405180830381600087803b158015610bc957600080fd5b505af1158015610bdd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0191906115c0565b508373ffffffffffffffffffffffffffffffffffffffff166342966c68866040518263ffffffff1660e01b8152600401610c3b9190611d8d565b602060405180830381600087803b158015610c5557600080fd5b505af1158015610c69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c8d91906115c0565b50809250505092915050565b6000610cc77f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b611089565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610cf8610e78565b50565b610d277f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b611093565b60000160009054906101000a900460ff1615610d4b57610d468361109d565b610e73565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610d9157600080fd5b505afa925050508015610dc257506040513d601f19601f82011682018060405250810190610dbf91906115ed565b60015b610e01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610df890611ced565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114610e66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e5d90611cad565b60405180910390fd5b50610e72838383611156565b5b505050565b610e80611182565b73ffffffffffffffffffffffffffffffffffffffff16610e9e610965565b73ffffffffffffffffffffffffffffffffffffffff1614610ef4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eeb90611d2d565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff1661102e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102590611d6d565b60405180910390fd5b61103661118a565b565b600060019054906101000a900460ff16611087576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107e90611d6d565b60405180910390fd5b565b6000819050919050565b6000819050919050565b6110a681610fbc565b6110e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110dc90611d0d565b60405180910390fd5b806111127f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b611089565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61115f836111eb565b60008251118061116c5750805b1561117d5761117b838361123a565b505b505050565b600033905090565b600060019054906101000a900460ff166111d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111d090611d6d565b60405180910390fd5b6111e96111e4611182565b610ef6565b565b6111f48161109d565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b606061125f83836040518060600160405280602781526020016122fc60279139611267565b905092915050565b60606000808573ffffffffffffffffffffffffffffffffffffffff16856040516112919190611a86565b600060405180830381855af49150503d80600081146112cc576040519150601f19603f3d011682016040523d82523d6000602084013e6112d1565b606091505b50915091506112e2868383876112ed565b925050509392505050565b60608315611350576000835114156113485761130885610fbc565b611347576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133e90611d4d565b60405180910390fd5b5b82905061135b565b61135a8383611363565b5b949350505050565b6000825111156113765781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113aa9190611c0b565b60405180910390fd5b60006113c66113c184611dcd565b611da8565b9050828152602081018484840111156113e2576113e1611f63565b5b6113ed848285611eb2565b509392505050565b6000813590506114048161229f565b92915050565b6000815190506114198161229f565b92915050565b60008151905061142e816122b6565b92915050565b600081519050611443816122cd565b92915050565b60008083601f84011261145f5761145e611f59565b5b8235905067ffffffffffffffff81111561147c5761147b611f54565b5b60208301915083600182028301111561149857611497611f5e565b5b9250929050565b600082601f8301126114b4576114b3611f59565b5b81356114c48482602086016113b3565b91505092915050565b6000813590506114dc816122e4565b92915050565b6000815190506114f1816122e4565b92915050565b60006020828403121561150d5761150c611f6d565b5b600061151b848285016113f5565b91505092915050565b6000806040838503121561153b5761153a611f6d565b5b6000611549858286016113f5565b925050602083013567ffffffffffffffff81111561156a57611569611f68565b5b6115768582860161149f565b9150509250929050565b6000806040838503121561159757611596611f6d565b5b60006115a58582860161140a565b92505060206115b6858286016114e2565b9150509250929050565b6000602082840312156115d6576115d5611f6d565b5b60006115e48482850161141f565b91505092915050565b60006020828403121561160357611602611f6d565b5b600061161184828501611434565b91505092915050565b60008060006040848603121561163357611632611f6d565b5b600084013567ffffffffffffffff81111561165157611650611f68565b5b61165d8682870161149f565b935050602084013567ffffffffffffffff81111561167e5761167d611f68565b5b61168a86828701611449565b92509250509250925092565b6000806000606084860312156116af576116ae611f6d565b5b600084013567ffffffffffffffff8111156116cd576116cc611f68565b5b6116d98682870161149f565b93505060206116ea868287016114cd565b92505060406116fb868287016113f5565b9150509250925092565b60008060008060006080868803121561172157611720611f6d565b5b600086013567ffffffffffffffff81111561173f5761173e611f68565b5b61174b8882890161149f565b955050602061175c888289016114cd565b945050604061176d888289016113f5565b935050606086013567ffffffffffffffff81111561178e5761178d611f68565b5b61179a88828901611449565b92509250509295509295909350565b6000602082840312156117bf576117be611f6d565b5b60006117cd848285016114e2565b91505092915050565b6117df81611e41565b82525050565b6117ee81611e5f565b82525050565b60006118008385611e14565b935061180d838584611eb2565b61181683611f72565b840190509392505050565b600061182c82611dfe565b6118368185611e14565b9350611846818560208601611ec1565b61184f81611f72565b840191505092915050565b600061186582611dfe565b61186f8185611e25565b935061187f818560208601611ec1565b80840191505092915050565b61189481611ea0565b82525050565b60006118a582611e09565b6118af8185611e30565b93506118bf818560208601611ec1565b6118c881611f72565b840191505092915050565b60006118e0602683611e30565b91506118eb82611f83565b604082019050919050565b6000611903602c83611e30565b915061190e82611fd2565b604082019050919050565b6000611926602c83611e30565b915061193182612021565b604082019050919050565b6000611949603883611e30565b915061195482612070565b604082019050919050565b600061196c602983611e30565b9150611977826120bf565b604082019050919050565b600061198f602e83611e30565b915061199a8261210e565b604082019050919050565b60006119b2602e83611e30565b91506119bd8261215d565b604082019050919050565b60006119d5602d83611e30565b91506119e0826121ac565b604082019050919050565b60006119f8602083611e30565b9150611a03826121fb565b602082019050919050565b6000611a1b600083611e14565b9150611a2682612224565b600082019050919050565b6000611a3e601d83611e30565b9150611a4982612227565b602082019050919050565b6000611a61602b83611e30565b9150611a6c82612250565b604082019050919050565b611a8081611e89565b82525050565b6000611a92828461185a565b915081905092915050565b6000602082019050611ab260008301846117d6565b92915050565b6000606082019050611acd60008301866117d6565b611ada60208301856117d6565b611ae76040830184611a77565b949350505050565b6000602082019050611b0460008301846117e5565b92915050565b60006020820190508181036000830152611b258184866117f4565b90509392505050565b600060a0820190508181036000830152611b488189611821565b9050611b576020830188611a77565b611b646040830187611a77565b611b716060830186611a77565b8181036080830152611b848184866117f4565b9050979650505050505050565b600060a0820190508181036000830152611bab8187611821565b9050611bba6020830186611a77565b611bc76040830185611a77565b611bd46060830184611a77565b8181036080830152611be581611a0e565b905095945050505050565b6000602082019050611c05600083018461188b565b92915050565b60006020820190508181036000830152611c25818461189a565b905092915050565b60006020820190508181036000830152611c46816118d3565b9050919050565b60006020820190508181036000830152611c66816118f6565b9050919050565b60006020820190508181036000830152611c8681611919565b9050919050565b60006020820190508181036000830152611ca68161193c565b9050919050565b60006020820190508181036000830152611cc68161195f565b9050919050565b60006020820190508181036000830152611ce681611982565b9050919050565b60006020820190508181036000830152611d06816119a5565b9050919050565b60006020820190508181036000830152611d26816119c8565b9050919050565b60006020820190508181036000830152611d46816119eb565b9050919050565b60006020820190508181036000830152611d6681611a31565b9050919050565b60006020820190508181036000830152611d8681611a54565b9050919050565b6000602082019050611da26000830184611a77565b92915050565b6000611db2611dc3565b9050611dbe8282611ef4565b919050565b6000604051905090565b600067ffffffffffffffff821115611de857611de7611f25565b5b611df182611f72565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611e4c82611e69565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611eab82611e93565b9050919050565b82818337600083830152505050565b60005b83811015611edf578082015181840152602081019050611ec4565b83811115611eee576000848401525b50505050565b611efd82611f72565b810181811067ffffffffffffffff82111715611f1c57611f1b611f25565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b6122a881611e41565b81146122b357600080fd5b50565b6122bf81611e53565b81146122ca57600080fd5b50565b6122d681611e5f565b81146122e157600080fd5b50565b6122ed81611e89565b81146122f857600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212203ffa97d0596ba6bc875b89937b96de47c3c194233a65e8d986c22e28cfee482464736f6c63430008070033",
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

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_GatewayZEVM *GatewayZEVMCaller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_GatewayZEVM *GatewayZEVMSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _GatewayZEVM.Contract.FUNGIBLEMODULEADDRESS(&_GatewayZEVM.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_GatewayZEVM *GatewayZEVMCallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _GatewayZEVM.Contract.FUNGIBLEMODULEADDRESS(&_GatewayZEVM.CallOpts)
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

// GatewayZEVMWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the GatewayZEVM contract.
type GatewayZEVMWithdrawalIterator struct {
	Event *GatewayZEVMWithdrawal // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMWithdrawal)
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
		it.Event = new(GatewayZEVMWithdrawal)
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
func (it *GatewayZEVMWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMWithdrawal represents a Withdrawal event raised by the GatewayZEVM contract.
type GatewayZEVMWithdrawal struct {
	From            common.Address
	To              []byte
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterWithdrawal(opts *bind.FilterOpts, from []common.Address) (*GatewayZEVMWithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMWithdrawalIterator{contract: _GatewayZEVM.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *GatewayZEVMWithdrawal, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMWithdrawal)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message)
func (_GatewayZEVM *GatewayZEVMFilterer) ParseWithdrawal(log types.Log) (*GatewayZEVMWithdrawal, error) {
	event := new(GatewayZEVMWithdrawal)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
