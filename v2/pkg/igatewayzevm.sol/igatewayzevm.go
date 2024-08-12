// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igatewayzevm

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

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Asset         common.Address
	Amount        uint64
	RevertMessage []byte
}

// RevertOptions is an auto generated low-level Go binding around an user-defined struct.
type RevertOptions struct {
	RevertAddress common.Address
	CallOnRevert  bool
	AbortAddress  common.Address
	RevertMessage []byte
}

// ZContext is an auto generated low-level Go binding around an user-defined struct.
type ZContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// IGatewayZEVMMetaData contains all meta data concerning the IGatewayZEVM contract.
var IGatewayZEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structzContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structzContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndRevert\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structzContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structzContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structzContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Call\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotFungibleModule\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrFungible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]}]",
}

// IGatewayZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use IGatewayZEVMMetaData.ABI instead.
var IGatewayZEVMABI = IGatewayZEVMMetaData.ABI

// IGatewayZEVM is an auto generated Go binding around an Ethereum contract.
type IGatewayZEVM struct {
	IGatewayZEVMCaller     // Read-only binding to the contract
	IGatewayZEVMTransactor // Write-only binding to the contract
	IGatewayZEVMFilterer   // Log filterer for contract events
}

// IGatewayZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGatewayZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGatewayZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGatewayZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGatewayZEVMSession struct {
	Contract     *IGatewayZEVM     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGatewayZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGatewayZEVMCallerSession struct {
	Contract *IGatewayZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IGatewayZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGatewayZEVMTransactorSession struct {
	Contract     *IGatewayZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IGatewayZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGatewayZEVMRaw struct {
	Contract *IGatewayZEVM // Generic contract binding to access the raw methods on
}

// IGatewayZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGatewayZEVMCallerRaw struct {
	Contract *IGatewayZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// IGatewayZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGatewayZEVMTransactorRaw struct {
	Contract *IGatewayZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGatewayZEVM creates a new instance of IGatewayZEVM, bound to a specific deployed contract.
func NewIGatewayZEVM(address common.Address, backend bind.ContractBackend) (*IGatewayZEVM, error) {
	contract, err := bindIGatewayZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVM{IGatewayZEVMCaller: IGatewayZEVMCaller{contract: contract}, IGatewayZEVMTransactor: IGatewayZEVMTransactor{contract: contract}, IGatewayZEVMFilterer: IGatewayZEVMFilterer{contract: contract}}, nil
}

// NewIGatewayZEVMCaller creates a new read-only instance of IGatewayZEVM, bound to a specific deployed contract.
func NewIGatewayZEVMCaller(address common.Address, caller bind.ContractCaller) (*IGatewayZEVMCaller, error) {
	contract, err := bindIGatewayZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMCaller{contract: contract}, nil
}

// NewIGatewayZEVMTransactor creates a new write-only instance of IGatewayZEVM, bound to a specific deployed contract.
func NewIGatewayZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*IGatewayZEVMTransactor, error) {
	contract, err := bindIGatewayZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMTransactor{contract: contract}, nil
}

// NewIGatewayZEVMFilterer creates a new log filterer instance of IGatewayZEVM, bound to a specific deployed contract.
func NewIGatewayZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*IGatewayZEVMFilterer, error) {
	contract, err := bindIGatewayZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMFilterer{contract: contract}, nil
}

// bindIGatewayZEVM binds a generic wrapper to an already deployed contract.
func bindIGatewayZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGatewayZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayZEVM *IGatewayZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayZEVM.Contract.IGatewayZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayZEVM *IGatewayZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.IGatewayZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayZEVM *IGatewayZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.IGatewayZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayZEVM *IGatewayZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayZEVM *IGatewayZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayZEVM *IGatewayZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.contract.Transact(opts, method, params...)
}

// Call is a paid mutator transaction binding the contract method 0xe192e9e5.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (address,bool,address,bytes) revertOptions, uint256 gasLimit) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) Call(opts *bind.TransactOpts, receiver []byte, zrc20 common.Address, message []byte, revertOptions RevertOptions, gasLimit *big.Int) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "call", receiver, zrc20, message, revertOptions, gasLimit)
}

// Call is a paid mutator transaction binding the contract method 0xe192e9e5.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (address,bool,address,bytes) revertOptions, uint256 gasLimit) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) Call(receiver []byte, zrc20 common.Address, message []byte, revertOptions RevertOptions, gasLimit *big.Int) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Call(&_IGatewayZEVM.TransactOpts, receiver, zrc20, message, revertOptions, gasLimit)
}

// Call is a paid mutator transaction binding the contract method 0xe192e9e5.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (address,bool,address,bytes) revertOptions, uint256 gasLimit) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) Call(receiver []byte, zrc20 common.Address, message []byte, revertOptions RevertOptions, gasLimit *big.Int) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Call(&_IGatewayZEVM.TransactOpts, receiver, zrc20, message, revertOptions, gasLimit)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) Deposit(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "deposit", zrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) Deposit(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Deposit(&_IGatewayZEVM.TransactOpts, zrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) Deposit(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Deposit(&_IGatewayZEVM.TransactOpts, zrc20, amount, target)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) DepositAndCall(opts *bind.TransactOpts, context ZContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "depositAndCall", context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) DepositAndCall(context ZContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.DepositAndCall(&_IGatewayZEVM.TransactOpts, context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) DepositAndCall(context ZContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.DepositAndCall(&_IGatewayZEVM.TransactOpts, context, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) DepositAndCall0(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "depositAndCall0", context, zrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) DepositAndCall0(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.DepositAndCall0(&_IGatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) DepositAndCall0(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.DepositAndCall0(&_IGatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0xa13b8cff.
//
// Solidity: function depositAndRevert((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message, (address,uint64,bytes) revertContext) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) DepositAndRevert(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "depositAndRevert", context, zrc20, amount, target, message, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0xa13b8cff.
//
// Solidity: function depositAndRevert((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message, (address,uint64,bytes) revertContext) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) DepositAndRevert(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.DepositAndRevert(&_IGatewayZEVM.TransactOpts, context, zrc20, amount, target, message, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0xa13b8cff.
//
// Solidity: function depositAndRevert((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message, (address,uint64,bytes) revertContext) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) DepositAndRevert(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.DepositAndRevert(&_IGatewayZEVM.TransactOpts, context, zrc20, amount, target, message, revertContext)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) Execute(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "execute", context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) Execute(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Execute(&_IGatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) Execute(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Execute(&_IGatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0xaa6585d0.
//
// Solidity: function executeRevert((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message, (address,uint64,bytes) revertContext) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) ExecuteRevert(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "executeRevert", context, zrc20, amount, target, message, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0xaa6585d0.
//
// Solidity: function executeRevert((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message, (address,uint64,bytes) revertContext) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) ExecuteRevert(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.ExecuteRevert(&_IGatewayZEVM.TransactOpts, context, zrc20, amount, target, message, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0xaa6585d0.
//
// Solidity: function executeRevert((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message, (address,uint64,bytes) revertContext) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) ExecuteRevert(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.ExecuteRevert(&_IGatewayZEVM.TransactOpts, context, zrc20, amount, target, message, revertContext)
}

// Withdraw is a paid mutator transaction binding the contract method 0x71f6af54.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes) revertOptions) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) Withdraw(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "withdraw", receiver, amount, zrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x71f6af54.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes) revertOptions) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Withdraw(&_IGatewayZEVM.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x71f6af54.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes) revertOptions) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Withdraw(&_IGatewayZEVM.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x9ec60484.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes) revertOptions) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) Withdraw0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "withdraw0", receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x9ec60484.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes) revertOptions) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Withdraw0(&_IGatewayZEVM.TransactOpts, receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x9ec60484.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes) revertOptions) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Withdraw0(&_IGatewayZEVM.TransactOpts, receiver, amount, chainId, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x1e83da29.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes) revertOptions) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) WithdrawAndCall(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "withdrawAndCall", receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x1e83da29.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes) revertOptions) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) WithdrawAndCall(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.WithdrawAndCall(&_IGatewayZEVM.TransactOpts, receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x1e83da29.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes) revertOptions) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) WithdrawAndCall(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.WithdrawAndCall(&_IGatewayZEVM.TransactOpts, receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0xc138da5d.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (address,bool,address,bytes) revertOptions, uint256 gasLimit) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) WithdrawAndCall0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, revertOptions RevertOptions, gasLimit *big.Int) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "withdrawAndCall0", receiver, amount, zrc20, message, revertOptions, gasLimit)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0xc138da5d.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (address,bool,address,bytes) revertOptions, uint256 gasLimit) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) WithdrawAndCall0(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, revertOptions RevertOptions, gasLimit *big.Int) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.WithdrawAndCall0(&_IGatewayZEVM.TransactOpts, receiver, amount, zrc20, message, revertOptions, gasLimit)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0xc138da5d.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (address,bool,address,bytes) revertOptions, uint256 gasLimit) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) WithdrawAndCall0(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, revertOptions RevertOptions, gasLimit *big.Int) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.WithdrawAndCall0(&_IGatewayZEVM.TransactOpts, receiver, amount, zrc20, message, revertOptions, gasLimit)
}

// IGatewayZEVMCallIterator is returned from FilterCall and is used to iterate over the raw logs and unpacked data for Call events raised by the IGatewayZEVM contract.
type IGatewayZEVMCallIterator struct {
	Event *IGatewayZEVMCall // Event containing the contract specifics and raw log

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
func (it *IGatewayZEVMCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayZEVMCall)
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
		it.Event = new(IGatewayZEVMCall)
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
func (it *IGatewayZEVMCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayZEVMCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayZEVMCall represents a Call event raised by the IGatewayZEVM contract.
type IGatewayZEVMCall struct {
	Sender        common.Address
	Zrc20         common.Address
	Receiver      []byte
	Message       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCall is a free log retrieval operation binding the contract event 0x00f591fbf375f41aa42e0570842d2ffe8d576fbd09b191b11effd8f58c7a4aad.
//
// Solidity: event Call(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (address,bool,address,bytes) revertOptions)
func (_IGatewayZEVM *IGatewayZEVMFilterer) FilterCall(opts *bind.FilterOpts, sender []common.Address, zrc20 []common.Address) (*IGatewayZEVMCallIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _IGatewayZEVM.contract.FilterLogs(opts, "Call", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMCallIterator{contract: _IGatewayZEVM.contract, event: "Call", logs: logs, sub: sub}, nil
}

// WatchCall is a free log subscription operation binding the contract event 0x00f591fbf375f41aa42e0570842d2ffe8d576fbd09b191b11effd8f58c7a4aad.
//
// Solidity: event Call(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (address,bool,address,bytes) revertOptions)
func (_IGatewayZEVM *IGatewayZEVMFilterer) WatchCall(opts *bind.WatchOpts, sink chan<- *IGatewayZEVMCall, sender []common.Address, zrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _IGatewayZEVM.contract.WatchLogs(opts, "Call", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayZEVMCall)
				if err := _IGatewayZEVM.contract.UnpackLog(event, "Call", log); err != nil {
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

// ParseCall is a log parse operation binding the contract event 0x00f591fbf375f41aa42e0570842d2ffe8d576fbd09b191b11effd8f58c7a4aad.
//
// Solidity: event Call(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (address,bool,address,bytes) revertOptions)
func (_IGatewayZEVM *IGatewayZEVMFilterer) ParseCall(log types.Log) (*IGatewayZEVMCall, error) {
	event := new(IGatewayZEVMCall)
	if err := _IGatewayZEVM.contract.UnpackLog(event, "Call", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayZEVMWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the IGatewayZEVM contract.
type IGatewayZEVMWithdrawalIterator struct {
	Event *IGatewayZEVMWithdrawal // Event containing the contract specifics and raw log

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
func (it *IGatewayZEVMWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayZEVMWithdrawal)
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
		it.Event = new(IGatewayZEVMWithdrawal)
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
func (it *IGatewayZEVMWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayZEVMWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayZEVMWithdrawal represents a Withdrawal event raised by the IGatewayZEVM contract.
type IGatewayZEVMWithdrawal struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Zrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x64dceb6f53e438def301bbc224b2d20ca86b2d453ac04dece1ff6cc90e10def0.
//
// Solidity: event Withdrawal(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (address,bool,address,bytes) revertOptions)
func (_IGatewayZEVM *IGatewayZEVMFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*IGatewayZEVMWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _IGatewayZEVM.contract.FilterLogs(opts, "Withdrawal", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMWithdrawalIterator{contract: _IGatewayZEVM.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x64dceb6f53e438def301bbc224b2d20ca86b2d453ac04dece1ff6cc90e10def0.
//
// Solidity: event Withdrawal(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (address,bool,address,bytes) revertOptions)
func (_IGatewayZEVM *IGatewayZEVMFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *IGatewayZEVMWithdrawal, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _IGatewayZEVM.contract.WatchLogs(opts, "Withdrawal", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayZEVMWithdrawal)
				if err := _IGatewayZEVM.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x64dceb6f53e438def301bbc224b2d20ca86b2d453ac04dece1ff6cc90e10def0.
//
// Solidity: event Withdrawal(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (address,bool,address,bytes) revertOptions)
func (_IGatewayZEVM *IGatewayZEVMFilterer) ParseWithdrawal(log types.Log) (*IGatewayZEVMWithdrawal, error) {
	event := new(IGatewayZEVMWithdrawal)
	if err := _IGatewayZEVM.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
