// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package console2

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

// Console2MetaData contains all meta data concerning the Console2 contract.
var Console2MetaData = &bind.MetaData{
	ABI: "[]",
}

// Console2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Console2MetaData.ABI instead.
var Console2ABI = Console2MetaData.ABI

// Console2 is an auto generated Go binding around an Ethereum contract.
type Console2 struct {
	Console2Caller     // Read-only binding to the contract
	Console2Transactor // Write-only binding to the contract
	Console2Filterer   // Log filterer for contract events
}

// Console2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Console2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Console2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Console2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Console2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Console2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Console2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Console2Session struct {
	Contract     *Console2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Console2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Console2CallerSession struct {
	Contract *Console2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Console2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Console2TransactorSession struct {
	Contract     *Console2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Console2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Console2Raw struct {
	Contract *Console2 // Generic contract binding to access the raw methods on
}

// Console2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Console2CallerRaw struct {
	Contract *Console2Caller // Generic read-only contract binding to access the raw methods on
}

// Console2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Console2TransactorRaw struct {
	Contract *Console2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewConsole2 creates a new instance of Console2, bound to a specific deployed contract.
func NewConsole2(address common.Address, backend bind.ContractBackend) (*Console2, error) {
	contract, err := bindConsole2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Console2{Console2Caller: Console2Caller{contract: contract}, Console2Transactor: Console2Transactor{contract: contract}, Console2Filterer: Console2Filterer{contract: contract}}, nil
}

// NewConsole2Caller creates a new read-only instance of Console2, bound to a specific deployed contract.
func NewConsole2Caller(address common.Address, caller bind.ContractCaller) (*Console2Caller, error) {
	contract, err := bindConsole2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Console2Caller{contract: contract}, nil
}

// NewConsole2Transactor creates a new write-only instance of Console2, bound to a specific deployed contract.
func NewConsole2Transactor(address common.Address, transactor bind.ContractTransactor) (*Console2Transactor, error) {
	contract, err := bindConsole2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Console2Transactor{contract: contract}, nil
}

// NewConsole2Filterer creates a new log filterer instance of Console2, bound to a specific deployed contract.
func NewConsole2Filterer(address common.Address, filterer bind.ContractFilterer) (*Console2Filterer, error) {
	contract, err := bindConsole2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Console2Filterer{contract: contract}, nil
}

// bindConsole2 binds a generic wrapper to an already deployed contract.
func bindConsole2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Console2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Console2 *Console2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Console2.Contract.Console2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Console2 *Console2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Console2.Contract.Console2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Console2 *Console2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Console2.Contract.Console2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Console2 *Console2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Console2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Console2 *Console2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Console2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Console2 *Console2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Console2.Contract.contract.Transact(opts, method, params...)
}
