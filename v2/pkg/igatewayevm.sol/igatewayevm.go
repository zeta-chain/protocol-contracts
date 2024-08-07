// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igatewayevm

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

// IGatewayEVMMetaData contains all meta data concerning the IGatewayEVM contract.
var IGatewayEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revertWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Call\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevertedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// IGatewayEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use IGatewayEVMMetaData.ABI instead.
var IGatewayEVMABI = IGatewayEVMMetaData.ABI

// IGatewayEVM is an auto generated Go binding around an Ethereum contract.
type IGatewayEVM struct {
	IGatewayEVMCaller     // Read-only binding to the contract
	IGatewayEVMTransactor // Write-only binding to the contract
	IGatewayEVMFilterer   // Log filterer for contract events
}

// IGatewayEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGatewayEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGatewayEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGatewayEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGatewayEVMSession struct {
	Contract     *IGatewayEVM      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGatewayEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGatewayEVMCallerSession struct {
	Contract *IGatewayEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IGatewayEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGatewayEVMTransactorSession struct {
	Contract     *IGatewayEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IGatewayEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGatewayEVMRaw struct {
	Contract *IGatewayEVM // Generic contract binding to access the raw methods on
}

// IGatewayEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGatewayEVMCallerRaw struct {
	Contract *IGatewayEVMCaller // Generic read-only contract binding to access the raw methods on
}

// IGatewayEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGatewayEVMTransactorRaw struct {
	Contract *IGatewayEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGatewayEVM creates a new instance of IGatewayEVM, bound to a specific deployed contract.
func NewIGatewayEVM(address common.Address, backend bind.ContractBackend) (*IGatewayEVM, error) {
	contract, err := bindIGatewayEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVM{IGatewayEVMCaller: IGatewayEVMCaller{contract: contract}, IGatewayEVMTransactor: IGatewayEVMTransactor{contract: contract}, IGatewayEVMFilterer: IGatewayEVMFilterer{contract: contract}}, nil
}

// NewIGatewayEVMCaller creates a new read-only instance of IGatewayEVM, bound to a specific deployed contract.
func NewIGatewayEVMCaller(address common.Address, caller bind.ContractCaller) (*IGatewayEVMCaller, error) {
	contract, err := bindIGatewayEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMCaller{contract: contract}, nil
}

// NewIGatewayEVMTransactor creates a new write-only instance of IGatewayEVM, bound to a specific deployed contract.
func NewIGatewayEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*IGatewayEVMTransactor, error) {
	contract, err := bindIGatewayEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMTransactor{contract: contract}, nil
}

// NewIGatewayEVMFilterer creates a new log filterer instance of IGatewayEVM, bound to a specific deployed contract.
func NewIGatewayEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*IGatewayEVMFilterer, error) {
	contract, err := bindIGatewayEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMFilterer{contract: contract}, nil
}

// bindIGatewayEVM binds a generic wrapper to an already deployed contract.
func bindIGatewayEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGatewayEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayEVM *IGatewayEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayEVM.Contract.IGatewayEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayEVM *IGatewayEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.IGatewayEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayEVM *IGatewayEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.IGatewayEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayEVM *IGatewayEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayEVM *IGatewayEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayEVM *IGatewayEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.contract.Transact(opts, method, params...)
}

// Call is a paid mutator transaction binding the contract method 0x1b8b921d.
//
// Solidity: function call(address receiver, bytes payload) returns()
func (_IGatewayEVM *IGatewayEVMTransactor) Call(opts *bind.TransactOpts, receiver common.Address, payload []byte) (*types.Transaction, error) {
	return _IGatewayEVM.contract.Transact(opts, "call", receiver, payload)
}

// Call is a paid mutator transaction binding the contract method 0x1b8b921d.
//
// Solidity: function call(address receiver, bytes payload) returns()
func (_IGatewayEVM *IGatewayEVMSession) Call(receiver common.Address, payload []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.Call(&_IGatewayEVM.TransactOpts, receiver, payload)
}

// Call is a paid mutator transaction binding the contract method 0x1b8b921d.
//
// Solidity: function call(address receiver, bytes payload) returns()
func (_IGatewayEVM *IGatewayEVMTransactorSession) Call(receiver common.Address, payload []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.Call(&_IGatewayEVM.TransactOpts, receiver, payload)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns()
func (_IGatewayEVM *IGatewayEVMTransactor) Deposit(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _IGatewayEVM.contract.Transact(opts, "deposit", receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns()
func (_IGatewayEVM *IGatewayEVMSession) Deposit(receiver common.Address) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.Deposit(&_IGatewayEVM.TransactOpts, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns()
func (_IGatewayEVM *IGatewayEVMTransactorSession) Deposit(receiver common.Address) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.Deposit(&_IGatewayEVM.TransactOpts, receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset) returns()
func (_IGatewayEVM *IGatewayEVMTransactor) Deposit0(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, asset common.Address) (*types.Transaction, error) {
	return _IGatewayEVM.contract.Transact(opts, "deposit0", receiver, amount, asset)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset) returns()
func (_IGatewayEVM *IGatewayEVMSession) Deposit0(receiver common.Address, amount *big.Int, asset common.Address) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.Deposit0(&_IGatewayEVM.TransactOpts, receiver, amount, asset)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset) returns()
func (_IGatewayEVM *IGatewayEVMTransactorSession) Deposit0(receiver common.Address, amount *big.Int, asset common.Address) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.Deposit0(&_IGatewayEVM.TransactOpts, receiver, amount, asset)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x29c59b5d.
//
// Solidity: function depositAndCall(address receiver, bytes payload) payable returns()
func (_IGatewayEVM *IGatewayEVMTransactor) DepositAndCall(opts *bind.TransactOpts, receiver common.Address, payload []byte) (*types.Transaction, error) {
	return _IGatewayEVM.contract.Transact(opts, "depositAndCall", receiver, payload)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x29c59b5d.
//
// Solidity: function depositAndCall(address receiver, bytes payload) payable returns()
func (_IGatewayEVM *IGatewayEVMSession) DepositAndCall(receiver common.Address, payload []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.DepositAndCall(&_IGatewayEVM.TransactOpts, receiver, payload)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x29c59b5d.
//
// Solidity: function depositAndCall(address receiver, bytes payload) payable returns()
func (_IGatewayEVM *IGatewayEVMTransactorSession) DepositAndCall(receiver common.Address, payload []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.DepositAndCall(&_IGatewayEVM.TransactOpts, receiver, payload)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0x8c6f037f.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload) returns()
func (_IGatewayEVM *IGatewayEVMTransactor) DepositAndCall0(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, asset common.Address, payload []byte) (*types.Transaction, error) {
	return _IGatewayEVM.contract.Transact(opts, "depositAndCall0", receiver, amount, asset, payload)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0x8c6f037f.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload) returns()
func (_IGatewayEVM *IGatewayEVMSession) DepositAndCall0(receiver common.Address, amount *big.Int, asset common.Address, payload []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.DepositAndCall0(&_IGatewayEVM.TransactOpts, receiver, amount, asset, payload)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0x8c6f037f.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload) returns()
func (_IGatewayEVM *IGatewayEVMTransactorSession) DepositAndCall0(receiver common.Address, amount *big.Int, asset common.Address, payload []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.DepositAndCall0(&_IGatewayEVM.TransactOpts, receiver, amount, asset, payload)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_IGatewayEVM *IGatewayEVMTransactor) Execute(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _IGatewayEVM.contract.Transact(opts, "execute", destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_IGatewayEVM *IGatewayEVMSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.Execute(&_IGatewayEVM.TransactOpts, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_IGatewayEVM *IGatewayEVMTransactorSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.Execute(&_IGatewayEVM.TransactOpts, destination, data)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x35c018db.
//
// Solidity: function executeRevert(address destination, bytes data) payable returns()
func (_IGatewayEVM *IGatewayEVMTransactor) ExecuteRevert(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _IGatewayEVM.contract.Transact(opts, "executeRevert", destination, data)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x35c018db.
//
// Solidity: function executeRevert(address destination, bytes data) payable returns()
func (_IGatewayEVM *IGatewayEVMSession) ExecuteRevert(destination common.Address, data []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.ExecuteRevert(&_IGatewayEVM.TransactOpts, destination, data)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x35c018db.
//
// Solidity: function executeRevert(address destination, bytes data) payable returns()
func (_IGatewayEVM *IGatewayEVMTransactorSession) ExecuteRevert(destination common.Address, data []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.ExecuteRevert(&_IGatewayEVM.TransactOpts, destination, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_IGatewayEVM *IGatewayEVMTransactor) ExecuteWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IGatewayEVM.contract.Transact(opts, "executeWithERC20", token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_IGatewayEVM *IGatewayEVMSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.ExecuteWithERC20(&_IGatewayEVM.TransactOpts, token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_IGatewayEVM *IGatewayEVMTransactorSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.ExecuteWithERC20(&_IGatewayEVM.TransactOpts, token, to, amount, data)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xb8969bd4.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_IGatewayEVM *IGatewayEVMTransactor) RevertWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IGatewayEVM.contract.Transact(opts, "revertWithERC20", token, to, amount, data)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xb8969bd4.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_IGatewayEVM *IGatewayEVMSession) RevertWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.RevertWithERC20(&_IGatewayEVM.TransactOpts, token, to, amount, data)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xb8969bd4.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_IGatewayEVM *IGatewayEVMTransactorSession) RevertWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IGatewayEVM.Contract.RevertWithERC20(&_IGatewayEVM.TransactOpts, token, to, amount, data)
}

// IGatewayEVMCallIterator is returned from FilterCall and is used to iterate over the raw logs and unpacked data for Call events raised by the IGatewayEVM contract.
type IGatewayEVMCallIterator struct {
	Event *IGatewayEVMCall // Event containing the contract specifics and raw log

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
func (it *IGatewayEVMCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayEVMCall)
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
		it.Event = new(IGatewayEVMCall)
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
func (it *IGatewayEVMCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayEVMCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayEVMCall represents a Call event raised by the IGatewayEVM contract.
type IGatewayEVMCall struct {
	Sender   common.Address
	Receiver common.Address
	Payload  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCall is a free log retrieval operation binding the contract event 0x2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3.
//
// Solidity: event Call(address indexed sender, address indexed receiver, bytes payload)
func (_IGatewayEVM *IGatewayEVMFilterer) FilterCall(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*IGatewayEVMCallIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IGatewayEVM.contract.FilterLogs(opts, "Call", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMCallIterator{contract: _IGatewayEVM.contract, event: "Call", logs: logs, sub: sub}, nil
}

// WatchCall is a free log subscription operation binding the contract event 0x2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3.
//
// Solidity: event Call(address indexed sender, address indexed receiver, bytes payload)
func (_IGatewayEVM *IGatewayEVMFilterer) WatchCall(opts *bind.WatchOpts, sink chan<- *IGatewayEVMCall, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IGatewayEVM.contract.WatchLogs(opts, "Call", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayEVMCall)
				if err := _IGatewayEVM.contract.UnpackLog(event, "Call", log); err != nil {
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

// ParseCall is a log parse operation binding the contract event 0x2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3.
//
// Solidity: event Call(address indexed sender, address indexed receiver, bytes payload)
func (_IGatewayEVM *IGatewayEVMFilterer) ParseCall(log types.Log) (*IGatewayEVMCall, error) {
	event := new(IGatewayEVMCall)
	if err := _IGatewayEVM.contract.UnpackLog(event, "Call", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayEVMDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the IGatewayEVM contract.
type IGatewayEVMDepositIterator struct {
	Event *IGatewayEVMDeposit // Event containing the contract specifics and raw log

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
func (it *IGatewayEVMDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayEVMDeposit)
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
		it.Event = new(IGatewayEVMDeposit)
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
func (it *IGatewayEVMDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayEVMDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayEVMDeposit represents a Deposit event raised by the IGatewayEVM contract.
type IGatewayEVMDeposit struct {
	Sender   common.Address
	Receiver common.Address
	Amount   *big.Int
	Asset    common.Address
	Payload  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4.
//
// Solidity: event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload)
func (_IGatewayEVM *IGatewayEVMFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*IGatewayEVMDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IGatewayEVM.contract.FilterLogs(opts, "Deposit", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMDepositIterator{contract: _IGatewayEVM.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4.
//
// Solidity: event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload)
func (_IGatewayEVM *IGatewayEVMFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IGatewayEVMDeposit, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IGatewayEVM.contract.WatchLogs(opts, "Deposit", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayEVMDeposit)
				if err := _IGatewayEVM.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4.
//
// Solidity: event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload)
func (_IGatewayEVM *IGatewayEVMFilterer) ParseDeposit(log types.Log) (*IGatewayEVMDeposit, error) {
	event := new(IGatewayEVMDeposit)
	if err := _IGatewayEVM.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayEVMExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the IGatewayEVM contract.
type IGatewayEVMExecutedIterator struct {
	Event *IGatewayEVMExecuted // Event containing the contract specifics and raw log

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
func (it *IGatewayEVMExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayEVMExecuted)
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
		it.Event = new(IGatewayEVMExecuted)
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
func (it *IGatewayEVMExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayEVMExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayEVMExecuted represents a Executed event raised by the IGatewayEVM contract.
type IGatewayEVMExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_IGatewayEVM *IGatewayEVMFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*IGatewayEVMExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IGatewayEVM.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMExecutedIterator{contract: _IGatewayEVM.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_IGatewayEVM *IGatewayEVMFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *IGatewayEVMExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IGatewayEVM.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayEVMExecuted)
				if err := _IGatewayEVM.contract.UnpackLog(event, "Executed", log); err != nil {
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
func (_IGatewayEVM *IGatewayEVMFilterer) ParseExecuted(log types.Log) (*IGatewayEVMExecuted, error) {
	event := new(IGatewayEVMExecuted)
	if err := _IGatewayEVM.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayEVMExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the IGatewayEVM contract.
type IGatewayEVMExecutedWithERC20Iterator struct {
	Event *IGatewayEVMExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *IGatewayEVMExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayEVMExecutedWithERC20)
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
		it.Event = new(IGatewayEVMExecutedWithERC20)
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
func (it *IGatewayEVMExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayEVMExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayEVMExecutedWithERC20 represents a ExecutedWithERC20 event raised by the IGatewayEVM contract.
type IGatewayEVMExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IGatewayEVM *IGatewayEVMFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*IGatewayEVMExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IGatewayEVM.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMExecutedWithERC20Iterator{contract: _IGatewayEVM.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IGatewayEVM *IGatewayEVMFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *IGatewayEVMExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IGatewayEVM.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayEVMExecutedWithERC20)
				if err := _IGatewayEVM.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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
func (_IGatewayEVM *IGatewayEVMFilterer) ParseExecutedWithERC20(log types.Log) (*IGatewayEVMExecutedWithERC20, error) {
	event := new(IGatewayEVMExecutedWithERC20)
	if err := _IGatewayEVM.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayEVMRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the IGatewayEVM contract.
type IGatewayEVMRevertedIterator struct {
	Event *IGatewayEVMReverted // Event containing the contract specifics and raw log

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
func (it *IGatewayEVMRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayEVMReverted)
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
		it.Event = new(IGatewayEVMReverted)
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
func (it *IGatewayEVMRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayEVMRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayEVMReverted represents a Reverted event raised by the IGatewayEVM contract.
type IGatewayEVMReverted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0xd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c.
//
// Solidity: event Reverted(address indexed destination, uint256 value, bytes data)
func (_IGatewayEVM *IGatewayEVMFilterer) FilterReverted(opts *bind.FilterOpts, destination []common.Address) (*IGatewayEVMRevertedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IGatewayEVM.contract.FilterLogs(opts, "Reverted", destinationRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMRevertedIterator{contract: _IGatewayEVM.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c.
//
// Solidity: event Reverted(address indexed destination, uint256 value, bytes data)
func (_IGatewayEVM *IGatewayEVMFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *IGatewayEVMReverted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IGatewayEVM.contract.WatchLogs(opts, "Reverted", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayEVMReverted)
				if err := _IGatewayEVM.contract.UnpackLog(event, "Reverted", log); err != nil {
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

// ParseReverted is a log parse operation binding the contract event 0xd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c.
//
// Solidity: event Reverted(address indexed destination, uint256 value, bytes data)
func (_IGatewayEVM *IGatewayEVMFilterer) ParseReverted(log types.Log) (*IGatewayEVMReverted, error) {
	event := new(IGatewayEVMReverted)
	if err := _IGatewayEVM.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayEVMRevertedWithERC20Iterator is returned from FilterRevertedWithERC20 and is used to iterate over the raw logs and unpacked data for RevertedWithERC20 events raised by the IGatewayEVM contract.
type IGatewayEVMRevertedWithERC20Iterator struct {
	Event *IGatewayEVMRevertedWithERC20 // Event containing the contract specifics and raw log

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
func (it *IGatewayEVMRevertedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayEVMRevertedWithERC20)
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
		it.Event = new(IGatewayEVMRevertedWithERC20)
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
func (it *IGatewayEVMRevertedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayEVMRevertedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayEVMRevertedWithERC20 represents a RevertedWithERC20 event raised by the IGatewayEVM contract.
type IGatewayEVMRevertedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevertedWithERC20 is a free log retrieval operation binding the contract event 0x723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda7.
//
// Solidity: event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IGatewayEVM *IGatewayEVMFilterer) FilterRevertedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*IGatewayEVMRevertedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IGatewayEVM.contract.FilterLogs(opts, "RevertedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMRevertedWithERC20Iterator{contract: _IGatewayEVM.contract, event: "RevertedWithERC20", logs: logs, sub: sub}, nil
}

// WatchRevertedWithERC20 is a free log subscription operation binding the contract event 0x723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda7.
//
// Solidity: event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IGatewayEVM *IGatewayEVMFilterer) WatchRevertedWithERC20(opts *bind.WatchOpts, sink chan<- *IGatewayEVMRevertedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IGatewayEVM.contract.WatchLogs(opts, "RevertedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayEVMRevertedWithERC20)
				if err := _IGatewayEVM.contract.UnpackLog(event, "RevertedWithERC20", log); err != nil {
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

// ParseRevertedWithERC20 is a log parse operation binding the contract event 0x723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda7.
//
// Solidity: event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IGatewayEVM *IGatewayEVMFilterer) ParseRevertedWithERC20(log types.Log) (*IGatewayEVMRevertedWithERC20, error) {
	event := new(IGatewayEVMRevertedWithERC20)
	if err := _IGatewayEVM.contract.UnpackLog(event, "RevertedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
