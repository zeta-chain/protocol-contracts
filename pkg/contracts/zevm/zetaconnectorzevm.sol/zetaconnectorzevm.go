// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectorzevm

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

// ZetaInterfacesSendInput is an auto generated low-level Go binding around an user-defined struct.
type ZetaInterfacesSendInput struct {
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	DestinationGasLimit *big.Int
	Message             []byte
	ZetaValueAndGas     *big.Int
	ZetaParams          []byte
}

// ZetaConnectorZEVMMetaData contains all meta data concerning the ZetaConnectorZEVM contract.
var ZetaConnectorZEVMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wzeta_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedZetaSent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyFungibleModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWZETA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WZETATransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wzeta_\",\"type\":\"address\"}],\"name\":\"SetWZETA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zetaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"ZetaReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zetaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingZetaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"ZetaReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceTxOriginAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zetaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zetaValueAndGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"zetaParams\",\"type\":\"bytes\"}],\"name\":\"ZetaSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"onReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zetaTxSenderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingZetaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"onRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"zetaValueAndGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zetaParams\",\"type\":\"bytes\"}],\"internalType\":\"structZetaInterfaces.SendInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wzeta_\",\"type\":\"address\"}],\"name\":\"setWzetaAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wzeta\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620016d7380380620016d7833981810160405281019062000037919062000095565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200011a565b6000815190506200008f8162000100565b92915050565b600060208284031215620000ae57620000ad620000fb565b5b6000620000be848285016200007e565b91505092915050565b6000620000d482620000db565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200010b81620000c7565b81146200011757600080fd5b50565b6115ad806200012a6000396000f3fe6080604052600436106100585760003560e01c8062173d46146100e957806329dd214d146101145780633ce4a5bc1461013d578063942a5e1614610168578063eb3bacbd14610191578063ec026901146101ba576100e4565b366100e45760008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100e2576040517f6e6b6de700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b3480156100f557600080fd5b506100fe6101e3565b60405161010b91906111d0565b60405180910390f35b34801561012057600080fd5b5061013b60048036038101906101369190610ed1565b610207565b005b34801561014957600080fd5b50610152610584565b60405161015f91906111d0565b60405180910390f35b34801561017457600080fd5b5061018f600480360381019061018a9190610dc2565b61059c565b005b34801561019d57600080fd5b506101b860048036038101906101b39190610d95565b61090d565b005b3480156101c657600080fd5b506101e160048036038101906101dc9190610fa0565b610a00565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610280576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b1580156102e857600080fd5b505af11580156102fc573d6000803e3d6000fd5b505050505060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3087876040518463ffffffff1660e01b815260040161035e939291906111eb565b602060405180830381600087803b15801561037857600080fd5b505af115801561038c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b09190610ea4565b6103e6576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000838390501115610522578473ffffffffffffffffffffffffffffffffffffffff16633749c51a6040518060a001604052808b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018981526020018873ffffffffffffffffffffffffffffffffffffffff16815260200187815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152506040518263ffffffff1660e01b81526004016104ef919061134c565b600060405180830381600087803b15801561050957600080fd5b505af115801561051d573d6000803e3d6000fd5b505050505b808573ffffffffffffffffffffffffffffffffffffffff16877ff1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d6988b8b898989604051610572959493929190611303565b60405180910390a45050505050505050565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610615576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b15801561067d57600080fd5b505af1158015610691573d6000803e3d6000fd5b505050505060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd308b876040518463ffffffff1660e01b81526004016106f3939291906111eb565b602060405180830381600087803b15801561070d57600080fd5b505af1158015610721573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107459190610ea4565b61077b576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008383905011156108bd578873ffffffffffffffffffffffffffffffffffffffff16633ff0693c6040518060c001604052808c73ffffffffffffffffffffffffffffffffffffffff1681526020018b81526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200188815260200187815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152506040518263ffffffff1660e01b815260040161088a919061136e565b600060405180830381600087803b1580156108a457600080fd5b505af11580156108b8573d6000803e3d6000fd5b505050505b80857f521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c888b8b8b8b8a8a8a6040516108fa979695949392919061129e565b60405180910390a3505050505050505050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610986576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f7325870b05f8f3412c318a35fc6a74feca51ea15811ec7a257676ca4db9d4176816040516109f591906111d0565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd333084608001356040518463ffffffff1660e01b8152600401610a61939291906111eb565b602060405180830381600087803b158015610a7b57600080fd5b505af1158015610a8f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab39190610ea4565b610ae9576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d82608001356040518263ffffffff1660e01b8152600401610b469190611390565b600060405180830381600087803b158015610b6057600080fd5b505af1158015610b74573d6000803e3d6000fd5b50505050600073735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168260800135604051610bb6906111bb565b60006040518083038185875af1925050503d8060008114610bf3576040519150601f19603f3d011682016040523d82523d6000602084013e610bf8565b606091505b5050905080610c33576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600001353373ffffffffffffffffffffffffffffffffffffffff167f7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e432858060200190610c8191906113ab565b87608001358860400135898060600190610c9b91906113ab565b8b8060a00190610cab91906113ab565b604051610cc099989796959493929190611222565b60405180910390a35050565b600081359050610cdb8161151b565b92915050565b600081519050610cf081611532565b92915050565b600081359050610d0581611549565b92915050565b60008083601f840112610d2157610d206114df565b5b8235905067ffffffffffffffff811115610d3e57610d3d6114da565b5b602083019150836001820283011115610d5a57610d596114f3565b5b9250929050565b600060c08284031215610d7757610d766114e9565b5b81905092915050565b600081359050610d8f81611560565b92915050565b600060208284031215610dab57610daa611502565b5b6000610db984828501610ccc565b91505092915050565b600080600080600080600080600060e08a8c031215610de457610de3611502565b5b6000610df28c828d01610ccc565b9950506020610e038c828d01610d80565b98505060408a013567ffffffffffffffff811115610e2457610e236114fd565b5b610e308c828d01610d0b565b97509750506060610e438c828d01610d80565b9550506080610e548c828d01610d80565b94505060a08a013567ffffffffffffffff811115610e7557610e746114fd565b5b610e818c828d01610d0b565b935093505060c0610e948c828d01610cf6565b9150509295985092959850929598565b600060208284031215610eba57610eb9611502565b5b6000610ec884828501610ce1565b91505092915050565b60008060008060008060008060c0898b031215610ef157610ef0611502565b5b600089013567ffffffffffffffff811115610f0f57610f0e6114fd565b5b610f1b8b828c01610d0b565b98509850506020610f2e8b828c01610d80565b9650506040610f3f8b828c01610ccc565b9550506060610f508b828c01610d80565b945050608089013567ffffffffffffffff811115610f7157610f706114fd565b5b610f7d8b828c01610d0b565b935093505060a0610f908b828c01610cf6565b9150509295985092959890939650565b600060208284031215610fb657610fb5611502565b5b600082013567ffffffffffffffff811115610fd457610fd36114fd565b5b610fe084828501610d61565b91505092915050565b610ff281611446565b82525050565b61100181611446565b82525050565b6000611013838561142a565b9350611020838584611498565b61102983611507565b840190509392505050565b600061103f8261140e565b6110498185611419565b93506110598185602086016114a7565b61106281611507565b840191505092915050565b600061107a60008361143b565b915061108582611518565b600082019050919050565b600060a08301600083015184820360008601526110ad8282611034565b91505060208301516110c2602086018261119d565b5060408301516110d56040860182610fe9565b5060608301516110e8606086018261119d565b50608083015184820360808601526111008282611034565b9150508091505092915050565b600060c0830160008301516111256000860182610fe9565b506020830151611138602086018261119d565b50604083015184820360408601526111508282611034565b9150506060830151611165606086018261119d565b506080830151611178608086018261119d565b5060a083015184820360a08601526111908282611034565b9150508091505092915050565b6111a68161148e565b82525050565b6111b58161148e565b82525050565b60006111c68261106d565b9150819050919050565b60006020820190506111e56000830184610ff8565b92915050565b60006060820190506112006000830186610ff8565b61120d6020830185610ff8565b61121a60408301846111ac565b949350505050565b600060c082019050611237600083018c610ff8565b818103602083015261124a818a8c611007565b905061125960408301896111ac565b61126660608301886111ac565b8181036080830152611279818688611007565b905081810360a083015261128e818486611007565b90509a9950505050505050505050565b600060a0820190506112b3600083018a610ff8565b6112c060208301896111ac565b81810360408301526112d3818789611007565b90506112e260608301866111ac565b81810360808301526112f5818486611007565b905098975050505050505050565b6000606082019050818103600083015261131e818789611007565b905061132d60208301866111ac565b8181036040830152611340818486611007565b90509695505050505050565b600060208201905081810360008301526113668184611090565b905092915050565b60006020820190508181036000830152611388818461110d565b905092915050565b60006020820190506113a560008301846111ac565b92915050565b600080833560016020038436030381126113c8576113c76114ee565b5b80840192508235915067ffffffffffffffff8211156113ea576113e96114e4565b5b602083019250600182023603831315611406576114056114f8565b5b509250929050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006114518261146e565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156114c55780820151818401526020810190506114aa565b838111156114d4576000848401525b50505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b50565b61152481611446565b811461152f57600080fd5b50565b61153b81611458565b811461154657600080fd5b50565b61155281611464565b811461155d57600080fd5b50565b6115698161148e565b811461157457600080fd5b5056fea2646970667358221220791a738c017d91460a2a96b3f39009822b40983215cd9948e7dd61174ca835ce64736f6c63430008070033",
}

// ZetaConnectorZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorZEVMMetaData.ABI instead.
var ZetaConnectorZEVMABI = ZetaConnectorZEVMMetaData.ABI

// ZetaConnectorZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorZEVMMetaData.Bin instead.
var ZetaConnectorZEVMBin = ZetaConnectorZEVMMetaData.Bin

// DeployZetaConnectorZEVM deploys a new Ethereum contract, binding an instance of ZetaConnectorZEVM to it.
func DeployZetaConnectorZEVM(auth *bind.TransactOpts, backend bind.ContractBackend, wzeta_ common.Address) (common.Address, *types.Transaction, *ZetaConnectorZEVM, error) {
	parsed, err := ZetaConnectorZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorZEVMBin), backend, wzeta_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorZEVM{ZetaConnectorZEVMCaller: ZetaConnectorZEVMCaller{contract: contract}, ZetaConnectorZEVMTransactor: ZetaConnectorZEVMTransactor{contract: contract}, ZetaConnectorZEVMFilterer: ZetaConnectorZEVMFilterer{contract: contract}}, nil
}

// ZetaConnectorZEVM is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorZEVM struct {
	ZetaConnectorZEVMCaller     // Read-only binding to the contract
	ZetaConnectorZEVMTransactor // Write-only binding to the contract
	ZetaConnectorZEVMFilterer   // Log filterer for contract events
}

// ZetaConnectorZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorZEVMSession struct {
	Contract     *ZetaConnectorZEVM // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZetaConnectorZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorZEVMCallerSession struct {
	Contract *ZetaConnectorZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ZetaConnectorZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorZEVMTransactorSession struct {
	Contract     *ZetaConnectorZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZetaConnectorZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorZEVMRaw struct {
	Contract *ZetaConnectorZEVM // Generic contract binding to access the raw methods on
}

// ZetaConnectorZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorZEVMCallerRaw struct {
	Contract *ZetaConnectorZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorZEVMTransactorRaw struct {
	Contract *ZetaConnectorZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorZEVM creates a new instance of ZetaConnectorZEVM, bound to a specific deployed contract.
func NewZetaConnectorZEVM(address common.Address, backend bind.ContractBackend) (*ZetaConnectorZEVM, error) {
	contract, err := bindZetaConnectorZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVM{ZetaConnectorZEVMCaller: ZetaConnectorZEVMCaller{contract: contract}, ZetaConnectorZEVMTransactor: ZetaConnectorZEVMTransactor{contract: contract}, ZetaConnectorZEVMFilterer: ZetaConnectorZEVMFilterer{contract: contract}}, nil
}

// NewZetaConnectorZEVMCaller creates a new read-only instance of ZetaConnectorZEVM, bound to a specific deployed contract.
func NewZetaConnectorZEVMCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorZEVMCaller, error) {
	contract, err := bindZetaConnectorZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMCaller{contract: contract}, nil
}

// NewZetaConnectorZEVMTransactor creates a new write-only instance of ZetaConnectorZEVM, bound to a specific deployed contract.
func NewZetaConnectorZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorZEVMTransactor, error) {
	contract, err := bindZetaConnectorZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMTransactor{contract: contract}, nil
}

// NewZetaConnectorZEVMFilterer creates a new log filterer instance of ZetaConnectorZEVM, bound to a specific deployed contract.
func NewZetaConnectorZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorZEVMFilterer, error) {
	contract, err := bindZetaConnectorZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMFilterer{contract: contract}, nil
}

// bindZetaConnectorZEVM binds a generic wrapper to an already deployed contract.
func bindZetaConnectorZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorZEVM.Contract.ZetaConnectorZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.ZetaConnectorZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.ZetaConnectorZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.contract.Transact(opts, method, params...)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCaller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorZEVM.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.FUNGIBLEMODULEADDRESS(&_ZetaConnectorZEVM.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.FUNGIBLEMODULEADDRESS(&_ZetaConnectorZEVM.CallOpts)
}

// Wzeta is a free data retrieval call binding the contract method 0x00173d46.
//
// Solidity: function wzeta() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCaller) Wzeta(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorZEVM.contract.Call(opts, &out, "wzeta")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wzeta is a free data retrieval call binding the contract method 0x00173d46.
//
// Solidity: function wzeta() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) Wzeta() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.Wzeta(&_ZetaConnectorZEVM.CallOpts)
}

// Wzeta is a free data retrieval call binding the contract method 0x00173d46.
//
// Solidity: function wzeta() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerSession) Wzeta() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.Wzeta(&_ZetaConnectorZEVM.CallOpts)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) OnReceive(opts *bind.TransactOpts, zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "onReceive", zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) OnReceive(zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.OnReceive(&_ZetaConnectorZEVM.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) OnReceive(zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.OnReceive(&_ZetaConnectorZEVM.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) OnRevert(opts *bind.TransactOpts, zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "onRevert", zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) OnRevert(zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.OnRevert(&_ZetaConnectorZEVM.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) OnRevert(zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.OnRevert(&_ZetaConnectorZEVM.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) Send(opts *bind.TransactOpts, input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) Send(input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Send(&_ZetaConnectorZEVM.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) Send(input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Send(&_ZetaConnectorZEVM.TransactOpts, input)
}

// SetWzetaAddress is a paid mutator transaction binding the contract method 0xeb3bacbd.
//
// Solidity: function setWzetaAddress(address wzeta_) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) SetWzetaAddress(opts *bind.TransactOpts, wzeta_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "setWzetaAddress", wzeta_)
}

// SetWzetaAddress is a paid mutator transaction binding the contract method 0xeb3bacbd.
//
// Solidity: function setWzetaAddress(address wzeta_) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) SetWzetaAddress(wzeta_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.SetWzetaAddress(&_ZetaConnectorZEVM.TransactOpts, wzeta_)
}

// SetWzetaAddress is a paid mutator transaction binding the contract method 0xeb3bacbd.
//
// Solidity: function setWzetaAddress(address wzeta_) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) SetWzetaAddress(wzeta_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.SetWzetaAddress(&_ZetaConnectorZEVM.TransactOpts, wzeta_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) Receive() (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Receive(&_ZetaConnectorZEVM.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) Receive() (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Receive(&_ZetaConnectorZEVM.TransactOpts)
}

// ZetaConnectorZEVMSetWZETAIterator is returned from FilterSetWZETA and is used to iterate over the raw logs and unpacked data for SetWZETA events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMSetWZETAIterator struct {
	Event *ZetaConnectorZEVMSetWZETA // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorZEVMSetWZETAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMSetWZETA)
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
		it.Event = new(ZetaConnectorZEVMSetWZETA)
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
func (it *ZetaConnectorZEVMSetWZETAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMSetWZETAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMSetWZETA represents a SetWZETA event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMSetWZETA struct {
	Wzeta common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetWZETA is a free log retrieval operation binding the contract event 0x7325870b05f8f3412c318a35fc6a74feca51ea15811ec7a257676ca4db9d4176.
//
// Solidity: event SetWZETA(address wzeta_)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterSetWZETA(opts *bind.FilterOpts) (*ZetaConnectorZEVMSetWZETAIterator, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "SetWZETA")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMSetWZETAIterator{contract: _ZetaConnectorZEVM.contract, event: "SetWZETA", logs: logs, sub: sub}, nil
}

// WatchSetWZETA is a free log subscription operation binding the contract event 0x7325870b05f8f3412c318a35fc6a74feca51ea15811ec7a257676ca4db9d4176.
//
// Solidity: event SetWZETA(address wzeta_)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchSetWZETA(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMSetWZETA) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "SetWZETA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMSetWZETA)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "SetWZETA", log); err != nil {
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

// ParseSetWZETA is a log parse operation binding the contract event 0x7325870b05f8f3412c318a35fc6a74feca51ea15811ec7a257676ca4db9d4176.
//
// Solidity: event SetWZETA(address wzeta_)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseSetWZETA(log types.Log) (*ZetaConnectorZEVMSetWZETA, error) {
	event := new(ZetaConnectorZEVMSetWZETA)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "SetWZETA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMZetaReceivedIterator is returned from FilterZetaReceived and is used to iterate over the raw logs and unpacked data for ZetaReceived events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaReceivedIterator struct {
	Event *ZetaConnectorZEVMZetaReceived // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorZEVMZetaReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMZetaReceived)
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
		it.Event = new(ZetaConnectorZEVMZetaReceived)
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
func (it *ZetaConnectorZEVMZetaReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMZetaReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMZetaReceived represents a ZetaReceived event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaReceived struct {
	ZetaTxSenderAddress []byte
	SourceChainId       *big.Int
	DestinationAddress  common.Address
	ZetaValue           *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterZetaReceived is a free log retrieval operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterZetaReceived(opts *bind.FilterOpts, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (*ZetaConnectorZEVMZetaReceivedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "ZetaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMZetaReceivedIterator{contract: _ZetaConnectorZEVM.contract, event: "ZetaReceived", logs: logs, sub: sub}, nil
}

// WatchZetaReceived is a free log subscription operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchZetaReceived(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMZetaReceived, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "ZetaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMZetaReceived)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaReceived", log); err != nil {
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

// ParseZetaReceived is a log parse operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseZetaReceived(log types.Log) (*ZetaConnectorZEVMZetaReceived, error) {
	event := new(ZetaConnectorZEVMZetaReceived)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMZetaRevertedIterator is returned from FilterZetaReverted and is used to iterate over the raw logs and unpacked data for ZetaReverted events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaRevertedIterator struct {
	Event *ZetaConnectorZEVMZetaReverted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorZEVMZetaRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMZetaReverted)
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
		it.Event = new(ZetaConnectorZEVMZetaReverted)
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
func (it *ZetaConnectorZEVMZetaRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMZetaRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMZetaReverted represents a ZetaReverted event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaReverted struct {
	ZetaTxSenderAddress common.Address
	SourceChainId       *big.Int
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	RemainingZetaValue  *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterZetaReverted is a free log retrieval operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterZetaReverted(opts *bind.FilterOpts, destinationChainId []*big.Int, internalSendHash [][32]byte) (*ZetaConnectorZEVMZetaRevertedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "ZetaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMZetaRevertedIterator{contract: _ZetaConnectorZEVM.contract, event: "ZetaReverted", logs: logs, sub: sub}, nil
}

// WatchZetaReverted is a free log subscription operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchZetaReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMZetaReverted, destinationChainId []*big.Int, internalSendHash [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "ZetaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMZetaReverted)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaReverted", log); err != nil {
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

// ParseZetaReverted is a log parse operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseZetaReverted(log types.Log) (*ZetaConnectorZEVMZetaReverted, error) {
	event := new(ZetaConnectorZEVMZetaReverted)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMZetaSentIterator is returned from FilterZetaSent and is used to iterate over the raw logs and unpacked data for ZetaSent events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaSentIterator struct {
	Event *ZetaConnectorZEVMZetaSent // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorZEVMZetaSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMZetaSent)
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
		it.Event = new(ZetaConnectorZEVMZetaSent)
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
func (it *ZetaConnectorZEVMZetaSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMZetaSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMZetaSent represents a ZetaSent event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaSent struct {
	SourceTxOriginAddress common.Address
	ZetaTxSenderAddress   common.Address
	DestinationChainId    *big.Int
	DestinationAddress    []byte
	ZetaValueAndGas       *big.Int
	DestinationGasLimit   *big.Int
	Message               []byte
	ZetaParams            []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterZetaSent is a free log retrieval operation binding the contract event 0x7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4.
//
// Solidity: event ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterZetaSent(opts *bind.FilterOpts, zetaTxSenderAddress []common.Address, destinationChainId []*big.Int) (*ZetaConnectorZEVMZetaSentIterator, error) {

	var zetaTxSenderAddressRule []interface{}
	for _, zetaTxSenderAddressItem := range zetaTxSenderAddress {
		zetaTxSenderAddressRule = append(zetaTxSenderAddressRule, zetaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "ZetaSent", zetaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMZetaSentIterator{contract: _ZetaConnectorZEVM.contract, event: "ZetaSent", logs: logs, sub: sub}, nil
}

// WatchZetaSent is a free log subscription operation binding the contract event 0x7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4.
//
// Solidity: event ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchZetaSent(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMZetaSent, zetaTxSenderAddress []common.Address, destinationChainId []*big.Int) (event.Subscription, error) {

	var zetaTxSenderAddressRule []interface{}
	for _, zetaTxSenderAddressItem := range zetaTxSenderAddress {
		zetaTxSenderAddressRule = append(zetaTxSenderAddressRule, zetaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "ZetaSent", zetaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMZetaSent)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaSent", log); err != nil {
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

// ParseZetaSent is a log parse operation binding the contract event 0x7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4.
//
// Solidity: event ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseZetaSent(log types.Log) (*ZetaConnectorZEVMZetaSent, error) {
	event := new(ZetaConnectorZEVMZetaSent)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
