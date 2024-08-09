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

// IGatewayEVMErrorsMetaData contains all meta data concerning the IGatewayEVMErrors contract.
var IGatewayEVMErrorsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// IGatewayEVMErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IGatewayEVMErrorsMetaData.ABI instead.
var IGatewayEVMErrorsABI = IGatewayEVMErrorsMetaData.ABI

// IGatewayEVMErrors is an auto generated Go binding around an Ethereum contract.
type IGatewayEVMErrors struct {
	IGatewayEVMErrorsCaller     // Read-only binding to the contract
	IGatewayEVMErrorsTransactor // Write-only binding to the contract
	IGatewayEVMErrorsFilterer   // Log filterer for contract events
}

// IGatewayEVMErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGatewayEVMErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayEVMErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGatewayEVMErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayEVMErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGatewayEVMErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayEVMErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGatewayEVMErrorsSession struct {
	Contract     *IGatewayEVMErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IGatewayEVMErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGatewayEVMErrorsCallerSession struct {
	Contract *IGatewayEVMErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IGatewayEVMErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGatewayEVMErrorsTransactorSession struct {
	Contract     *IGatewayEVMErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IGatewayEVMErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGatewayEVMErrorsRaw struct {
	Contract *IGatewayEVMErrors // Generic contract binding to access the raw methods on
}

// IGatewayEVMErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGatewayEVMErrorsCallerRaw struct {
	Contract *IGatewayEVMErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IGatewayEVMErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGatewayEVMErrorsTransactorRaw struct {
	Contract *IGatewayEVMErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGatewayEVMErrors creates a new instance of IGatewayEVMErrors, bound to a specific deployed contract.
func NewIGatewayEVMErrors(address common.Address, backend bind.ContractBackend) (*IGatewayEVMErrors, error) {
	contract, err := bindIGatewayEVMErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMErrors{IGatewayEVMErrorsCaller: IGatewayEVMErrorsCaller{contract: contract}, IGatewayEVMErrorsTransactor: IGatewayEVMErrorsTransactor{contract: contract}, IGatewayEVMErrorsFilterer: IGatewayEVMErrorsFilterer{contract: contract}}, nil
}

// NewIGatewayEVMErrorsCaller creates a new read-only instance of IGatewayEVMErrors, bound to a specific deployed contract.
func NewIGatewayEVMErrorsCaller(address common.Address, caller bind.ContractCaller) (*IGatewayEVMErrorsCaller, error) {
	contract, err := bindIGatewayEVMErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMErrorsCaller{contract: contract}, nil
}

// NewIGatewayEVMErrorsTransactor creates a new write-only instance of IGatewayEVMErrors, bound to a specific deployed contract.
func NewIGatewayEVMErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IGatewayEVMErrorsTransactor, error) {
	contract, err := bindIGatewayEVMErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMErrorsTransactor{contract: contract}, nil
}

// NewIGatewayEVMErrorsFilterer creates a new log filterer instance of IGatewayEVMErrors, bound to a specific deployed contract.
func NewIGatewayEVMErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IGatewayEVMErrorsFilterer, error) {
	contract, err := bindIGatewayEVMErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMErrorsFilterer{contract: contract}, nil
}

// bindIGatewayEVMErrors binds a generic wrapper to an already deployed contract.
func bindIGatewayEVMErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGatewayEVMErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayEVMErrors *IGatewayEVMErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayEVMErrors.Contract.IGatewayEVMErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayEVMErrors *IGatewayEVMErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayEVMErrors.Contract.IGatewayEVMErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayEVMErrors *IGatewayEVMErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayEVMErrors.Contract.IGatewayEVMErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayEVMErrors *IGatewayEVMErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayEVMErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayEVMErrors *IGatewayEVMErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayEVMErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayEVMErrors *IGatewayEVMErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayEVMErrors.Contract.contract.Transact(opts, method, params...)
}
