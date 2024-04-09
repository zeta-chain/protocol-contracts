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
	Bin: "0x60806040523480156200001157600080fd5b506040516200140b3803806200140b833981810160405281019062000037919062000095565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200011a565b6000815190506200008f8162000100565b92915050565b600060208284031215620000ae57620000ad620000fb565b5b6000620000be848285016200007e565b91505092915050565b6000620000d482620000db565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200010b81620000c7565b81146200011757600080fd5b50565b6112e1806200012a6000396000f3fe6080604052600436106100585760003560e01c8062173d46146100e957806329dd214d146101145780633ce4a5bc1461013d578063942a5e1614610168578063eb3bacbd14610191578063ec026901146101ba576100e4565b366100e45760008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100e2576040517f6e6b6de700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b3480156100f557600080fd5b506100fe6101e3565b60405161010b9190610f04565b60405180910390f35b34801561012057600080fd5b5061013b60048036038101906101369190610c05565b610207565b005b34801561014957600080fd5b5061015261041e565b60405161015f9190610f04565b60405180910390f35b34801561017457600080fd5b5061018f600480360381019061018a9190610af6565b610436565b005b34801561019d57600080fd5b506101b860048036038101906101b39190610ac9565b610641565b005b3480156101c657600080fd5b506101e160048036038101906101dc9190610cd4565b610734565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610280576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008383905011156103bc578473ffffffffffffffffffffffffffffffffffffffff16633749c51a6040518060a001604052808b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018981526020018873ffffffffffffffffffffffffffffffffffffffff16815260200187815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152506040518263ffffffff1660e01b81526004016103899190611080565b600060405180830381600087803b1580156103a357600080fd5b505af11580156103b7573d6000803e3d6000fd5b505050505b808573ffffffffffffffffffffffffffffffffffffffff16877ff1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d6988b8b89898960405161040c959493929190611037565b60405180910390a45050505050505050565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104af576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008383905011156105f1578873ffffffffffffffffffffffffffffffffffffffff16633ff0693c6040518060c001604052808c73ffffffffffffffffffffffffffffffffffffffff1681526020018b81526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200188815260200187815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152506040518263ffffffff1660e01b81526004016105be91906110a2565b600060405180830381600087803b1580156105d857600080fd5b505af11580156105ec573d6000803e3d6000fd5b505050505b80857f521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c888b8b8b8b8a8a8a60405161062e9796959493929190610fd2565b60405180910390a3505050505050505050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ba576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f7325870b05f8f3412c318a35fc6a74feca51ea15811ec7a257676ca4db9d4176816040516107299190610f04565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd333084608001356040518463ffffffff1660e01b815260040161079593929190610f1f565b602060405180830381600087803b1580156107af57600080fd5b505af11580156107c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e79190610bd8565b61081d576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d82608001356040518263ffffffff1660e01b815260040161087a91906110c4565b600060405180830381600087803b15801561089457600080fd5b505af11580156108a8573d6000803e3d6000fd5b50505050600073735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff1682608001356040516108ea90610eef565b60006040518083038185875af1925050503d8060008114610927576040519150601f19603f3d011682016040523d82523d6000602084013e61092c565b606091505b5050905080610967576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600001353373ffffffffffffffffffffffffffffffffffffffff167f7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4328580602001906109b591906110df565b876080013588604001358980606001906109cf91906110df565b8b8060a001906109df91906110df565b6040516109f499989796959493929190610f56565b60405180910390a35050565b600081359050610a0f8161124f565b92915050565b600081519050610a2481611266565b92915050565b600081359050610a398161127d565b92915050565b60008083601f840112610a5557610a54611213565b5b8235905067ffffffffffffffff811115610a7257610a7161120e565b5b602083019150836001820283011115610a8e57610a8d611227565b5b9250929050565b600060c08284031215610aab57610aaa61121d565b5b81905092915050565b600081359050610ac381611294565b92915050565b600060208284031215610adf57610ade611236565b5b6000610aed84828501610a00565b91505092915050565b600080600080600080600080600060e08a8c031215610b1857610b17611236565b5b6000610b268c828d01610a00565b9950506020610b378c828d01610ab4565b98505060408a013567ffffffffffffffff811115610b5857610b57611231565b5b610b648c828d01610a3f565b97509750506060610b778c828d01610ab4565b9550506080610b888c828d01610ab4565b94505060a08a013567ffffffffffffffff811115610ba957610ba8611231565b5b610bb58c828d01610a3f565b935093505060c0610bc88c828d01610a2a565b9150509295985092959850929598565b600060208284031215610bee57610bed611236565b5b6000610bfc84828501610a15565b91505092915050565b60008060008060008060008060c0898b031215610c2557610c24611236565b5b600089013567ffffffffffffffff811115610c4357610c42611231565b5b610c4f8b828c01610a3f565b98509850506020610c628b828c01610ab4565b9650506040610c738b828c01610a00565b9550506060610c848b828c01610ab4565b945050608089013567ffffffffffffffff811115610ca557610ca4611231565b5b610cb18b828c01610a3f565b935093505060a0610cc48b828c01610a2a565b9150509295985092959890939650565b600060208284031215610cea57610ce9611236565b5b600082013567ffffffffffffffff811115610d0857610d07611231565b5b610d1484828501610a95565b91505092915050565b610d268161117a565b82525050565b610d358161117a565b82525050565b6000610d47838561115e565b9350610d548385846111cc565b610d5d8361123b565b840190509392505050565b6000610d7382611142565b610d7d818561114d565b9350610d8d8185602086016111db565b610d968161123b565b840191505092915050565b6000610dae60008361116f565b9150610db98261124c565b600082019050919050565b600060a0830160008301518482036000860152610de18282610d68565b9150506020830151610df66020860182610ed1565b506040830151610e096040860182610d1d565b506060830151610e1c6060860182610ed1565b5060808301518482036080860152610e348282610d68565b9150508091505092915050565b600060c083016000830151610e596000860182610d1d565b506020830151610e6c6020860182610ed1565b5060408301518482036040860152610e848282610d68565b9150506060830151610e996060860182610ed1565b506080830151610eac6080860182610ed1565b5060a083015184820360a0860152610ec48282610d68565b9150508091505092915050565b610eda816111c2565b82525050565b610ee9816111c2565b82525050565b6000610efa82610da1565b9150819050919050565b6000602082019050610f196000830184610d2c565b92915050565b6000606082019050610f346000830186610d2c565b610f416020830185610d2c565b610f4e6040830184610ee0565b949350505050565b600060c082019050610f6b600083018c610d2c565b8181036020830152610f7e818a8c610d3b565b9050610f8d6040830189610ee0565b610f9a6060830188610ee0565b8181036080830152610fad818688610d3b565b905081810360a0830152610fc2818486610d3b565b90509a9950505050505050505050565b600060a082019050610fe7600083018a610d2c565b610ff46020830189610ee0565b8181036040830152611007818789610d3b565b90506110166060830186610ee0565b8181036080830152611029818486610d3b565b905098975050505050505050565b60006060820190508181036000830152611052818789610d3b565b90506110616020830186610ee0565b8181036040830152611074818486610d3b565b90509695505050505050565b6000602082019050818103600083015261109a8184610dc4565b905092915050565b600060208201905081810360008301526110bc8184610e41565b905092915050565b60006020820190506110d96000830184610ee0565b92915050565b600080833560016020038436030381126110fc576110fb611222565b5b80840192508235915067ffffffffffffffff82111561111e5761111d611218565b5b60208301925060018202360383131561113a5761113961122c565b5b509250929050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611185826111a2565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156111f95780820151818401526020810190506111de565b83811115611208576000848401525b50505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b50565b6112588161117a565b811461126357600080fd5b50565b61126f8161118c565b811461127a57600080fd5b50565b61128681611198565b811461129157600080fd5b50565b61129d816111c2565b81146112a857600080fd5b5056fea2646970667358221220b3a3931eeb28cdb7c155d8c742a593159337562cab897949bfc1d07aa0acccbb64736f6c63430008070033",
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
	parsed, err := abi.JSON(strings.NewReader(ZetaConnectorZEVMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
