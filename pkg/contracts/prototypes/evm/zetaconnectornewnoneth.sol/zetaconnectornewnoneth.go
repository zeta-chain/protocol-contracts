// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornewnoneth

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

// ZetaConnectorNewNonEthMetaData contains all meta data concerning the ZetaConnectorNewNonEth contract.
var ZetaConnectorNewNonEthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zetaToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIGatewayEVM\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"receiveTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b50604051610c18380380610c1883398181016040528101906100329190610166565b81816001600081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806100a35750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156100da576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b81525050505050506101f4565b600081519050610160816101dd565b92915050565b6000806040838503121561017d5761017c6101d8565b5b600061018b85828601610151565b925050602061019c85828601610151565b9150509250929050565b60006101b1826101b8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6101e6816101a6565b81146101f157600080fd5b50565b60805160601c60a05160601c6109d06102486000396000818160f601528181610204015281816102300152818161031b01526103f30152600081816101e00152818161026c01526102df01526109d06000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063106e62901461005c578063116191b61461007857806321e093b1146100965780635e3e9fef146100b4578063743e0c9b146100d0575b600080fd5b61007660048036038101906100719190610570565b6100ec565b005b6100806101de565b60405161008d91906107cd565b60405180910390f35b61009e610202565b6040516100ab9190610704565b60405180910390f35b6100ce60048036038101906100c991906105c3565b610226565b005b6100ea60048036038101906100e5919061064b565b6103f1565b005b6100f4610481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631e458bee8484846040518463ffffffff1660e01b815260040161015193929190610796565b600060405180830381600087803b15801561016b57600080fd5b505af115801561017f573d6000803e3d6000fd5b505050508273ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364836040516101c99190610808565b60405180910390a26101d96104d1565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b61022e610481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631e458bee7f000000000000000000000000000000000000000000000000000000000000000086846040518463ffffffff1660e01b81526004016102ab93929190610796565b600060405180830381600087803b1580156102c557600080fd5b505af11580156102d9573d6000803e3d6000fd5b505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635131ab597f0000000000000000000000000000000000000000000000000000000000000000878787876040518663ffffffff1660e01b815260040161035e95949392919061071f565b600060405180830381600087803b15801561037857600080fd5b505af115801561038c573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced8585856040516103da93929190610823565b60405180910390a26103ea6104d1565b5050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379cc679033836040518363ffffffff1660e01b815260040161044c92919061076d565b600060405180830381600087803b15801561046657600080fd5b505af115801561047a573d6000803e3d6000fd5b5050505050565b600260005414156104c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104be906107e8565b60405180910390fd5b6002600081905550565b6001600081905550565b6000813590506104ea81610955565b92915050565b6000813590506104ff8161096c565b92915050565b60008083601f84011261051b5761051a610907565b5b8235905067ffffffffffffffff81111561053857610537610902565b5b6020830191508360018202830111156105545761055361090c565b5b9250929050565b60008135905061056a81610983565b92915050565b60008060006060848603121561058957610588610916565b5b6000610597868287016104db565b93505060206105a88682870161055b565b92505060406105b9868287016104f0565b9150509250925092565b6000806000806000608086880312156105df576105de610916565b5b60006105ed888289016104db565b95505060206105fe8882890161055b565b945050604086013567ffffffffffffffff81111561061f5761061e610911565b5b61062b88828901610505565b9350935050606061063e888289016104f0565b9150509295509295909350565b60006020828403121561066157610660610916565b5b600061066f8482850161055b565b91505092915050565b61068181610877565b82525050565b61069081610889565b82525050565b60006106a28385610855565b93506106af8385846108f3565b6106b88361091b565b840190509392505050565b6106cc816108bd565b82525050565b60006106df601f83610866565b91506106ea8261092c565b602082019050919050565b6106fe816108b3565b82525050565b60006020820190506107196000830184610678565b92915050565b60006080820190506107346000830188610678565b6107416020830187610678565b61074e60408301866106f5565b8181036060830152610761818486610696565b90509695505050505050565b60006040820190506107826000830185610678565b61078f60208301846106f5565b9392505050565b60006060820190506107ab6000830186610678565b6107b860208301856106f5565b6107c56040830184610687565b949350505050565b60006020820190506107e260008301846106c3565b92915050565b60006020820190508181036000830152610801816106d2565b9050919050565b600060208201905061081d60008301846106f5565b92915050565b600060408201905061083860008301866106f5565b818103602083015261084b818486610696565b9050949350505050565b600082825260208201905092915050565b600082825260208201905092915050565b600061088282610893565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006108c8826108cf565b9050919050565b60006108da826108e1565b9050919050565b60006108ec82610893565b9050919050565b82818337600083830152505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b61095e81610877565b811461096957600080fd5b50565b61097581610889565b811461098057600080fd5b50565b61098c816108b3565b811461099757600080fd5b5056fea2646970667358221220af55d5f61addbf319928a723001d411ec2b726404eeedec369e4d204c7f69a1164736f6c63430008070033",
}

// ZetaConnectorNewNonEthABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNewNonEthMetaData.ABI instead.
var ZetaConnectorNewNonEthABI = ZetaConnectorNewNonEthMetaData.ABI

// ZetaConnectorNewNonEthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNewNonEthMetaData.Bin instead.
var ZetaConnectorNewNonEthBin = ZetaConnectorNewNonEthMetaData.Bin

// DeployZetaConnectorNewNonEth deploys a new Ethereum contract, binding an instance of ZetaConnectorNewNonEth to it.
func DeployZetaConnectorNewNonEth(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address, _zetaToken common.Address) (common.Address, *types.Transaction, *ZetaConnectorNewNonEth, error) {
	parsed, err := ZetaConnectorNewNonEthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNewNonEthBin), backend, _gateway, _zetaToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNewNonEth{ZetaConnectorNewNonEthCaller: ZetaConnectorNewNonEthCaller{contract: contract}, ZetaConnectorNewNonEthTransactor: ZetaConnectorNewNonEthTransactor{contract: contract}, ZetaConnectorNewNonEthFilterer: ZetaConnectorNewNonEthFilterer{contract: contract}}, nil
}

// ZetaConnectorNewNonEth is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNewNonEth struct {
	ZetaConnectorNewNonEthCaller     // Read-only binding to the contract
	ZetaConnectorNewNonEthTransactor // Write-only binding to the contract
	ZetaConnectorNewNonEthFilterer   // Log filterer for contract events
}

// ZetaConnectorNewNonEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNewNonEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNewNonEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNewNonEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNewNonEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNewNonEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNewNonEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNewNonEthSession struct {
	Contract     *ZetaConnectorNewNonEth // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZetaConnectorNewNonEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNewNonEthCallerSession struct {
	Contract *ZetaConnectorNewNonEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ZetaConnectorNewNonEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNewNonEthTransactorSession struct {
	Contract     *ZetaConnectorNewNonEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ZetaConnectorNewNonEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNewNonEthRaw struct {
	Contract *ZetaConnectorNewNonEth // Generic contract binding to access the raw methods on
}

// ZetaConnectorNewNonEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNewNonEthCallerRaw struct {
	Contract *ZetaConnectorNewNonEthCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNewNonEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNewNonEthTransactorRaw struct {
	Contract *ZetaConnectorNewNonEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNewNonEth creates a new instance of ZetaConnectorNewNonEth, bound to a specific deployed contract.
func NewZetaConnectorNewNonEth(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNewNonEth, error) {
	contract, err := bindZetaConnectorNewNonEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewNonEth{ZetaConnectorNewNonEthCaller: ZetaConnectorNewNonEthCaller{contract: contract}, ZetaConnectorNewNonEthTransactor: ZetaConnectorNewNonEthTransactor{contract: contract}, ZetaConnectorNewNonEthFilterer: ZetaConnectorNewNonEthFilterer{contract: contract}}, nil
}

// NewZetaConnectorNewNonEthCaller creates a new read-only instance of ZetaConnectorNewNonEth, bound to a specific deployed contract.
func NewZetaConnectorNewNonEthCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNewNonEthCaller, error) {
	contract, err := bindZetaConnectorNewNonEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewNonEthCaller{contract: contract}, nil
}

// NewZetaConnectorNewNonEthTransactor creates a new write-only instance of ZetaConnectorNewNonEth, bound to a specific deployed contract.
func NewZetaConnectorNewNonEthTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNewNonEthTransactor, error) {
	contract, err := bindZetaConnectorNewNonEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewNonEthTransactor{contract: contract}, nil
}

// NewZetaConnectorNewNonEthFilterer creates a new log filterer instance of ZetaConnectorNewNonEth, bound to a specific deployed contract.
func NewZetaConnectorNewNonEthFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNewNonEthFilterer, error) {
	contract, err := bindZetaConnectorNewNonEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewNonEthFilterer{contract: contract}, nil
}

// bindZetaConnectorNewNonEth binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNewNonEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNewNonEthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNewNonEth.Contract.ZetaConnectorNewNonEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.Contract.ZetaConnectorNewNonEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.Contract.ZetaConnectorNewNonEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNewNonEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNewNonEth.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNewNonEth.Contract.Gateway(&_ZetaConnectorNewNonEth.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthCallerSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNewNonEth.Contract.Gateway(&_ZetaConnectorNewNonEth.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNewNonEth.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNewNonEth.Contract.ZetaToken(&_ZetaConnectorNewNonEth.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNewNonEth.Contract.ZetaToken(&_ZetaConnectorNewNonEth.CallOpts)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthTransactor) ReceiveTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.contract.Transact(opts, "receiveTokens", amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.Contract.ReceiveTokens(&_ZetaConnectorNewNonEth.TransactOpts, amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthTransactorSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.Contract.ReceiveTokens(&_ZetaConnectorNewNonEth.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.contract.Transact(opts, "withdraw", to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.Contract.Withdraw(&_ZetaConnectorNewNonEth.TransactOpts, to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthTransactorSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.Contract.Withdraw(&_ZetaConnectorNewNonEth.TransactOpts, to, amount, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthTransactor) WithdrawAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.contract.Transact(opts, "withdrawAndCall", to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.Contract.WithdrawAndCall(&_ZetaConnectorNewNonEth.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthTransactorSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewNonEth.Contract.WithdrawAndCall(&_ZetaConnectorNewNonEth.TransactOpts, to, amount, data, internalSendHash)
}

// ZetaConnectorNewNonEthWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ZetaConnectorNewNonEth contract.
type ZetaConnectorNewNonEthWithdrawIterator struct {
	Event *ZetaConnectorNewNonEthWithdraw // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNewNonEthWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNewNonEthWithdraw)
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
		it.Event = new(ZetaConnectorNewNonEthWithdraw)
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
func (it *ZetaConnectorNewNonEthWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNewNonEthWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNewNonEthWithdraw represents a Withdraw event raised by the ZetaConnectorNewNonEth contract.
type ZetaConnectorNewNonEthWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNewNonEthWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewNonEth.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewNonEthWithdrawIterator{contract: _ZetaConnectorNewNonEth.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNewNonEthWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewNonEth.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNewNonEthWithdraw)
				if err := _ZetaConnectorNewNonEth.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthFilterer) ParseWithdraw(log types.Log) (*ZetaConnectorNewNonEthWithdraw, error) {
	event := new(ZetaConnectorNewNonEthWithdraw)
	if err := _ZetaConnectorNewNonEth.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNewNonEthWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the ZetaConnectorNewNonEth contract.
type ZetaConnectorNewNonEthWithdrawAndCallIterator struct {
	Event *ZetaConnectorNewNonEthWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNewNonEthWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNewNonEthWithdrawAndCall)
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
		it.Event = new(ZetaConnectorNewNonEthWithdrawAndCall)
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
func (it *ZetaConnectorNewNonEthWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNewNonEthWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNewNonEthWithdrawAndCall represents a WithdrawAndCall event raised by the ZetaConnectorNewNonEth contract.
type ZetaConnectorNewNonEthWithdrawAndCall struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNewNonEthWithdrawAndCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewNonEth.contract.FilterLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewNonEthWithdrawAndCallIterator{contract: _ZetaConnectorNewNonEth.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNewNonEthWithdrawAndCall, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewNonEth.contract.WatchLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNewNonEthWithdrawAndCall)
				if err := _ZetaConnectorNewNonEth.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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

// ParseWithdrawAndCall is a log parse operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNewNonEth *ZetaConnectorNewNonEthFilterer) ParseWithdrawAndCall(log types.Log) (*ZetaConnectorNewNonEthWithdrawAndCall, error) {
	event := new(ZetaConnectorNewNonEthWithdrawAndCall)
	if err := _ZetaConnectorNewNonEth.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
