// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izetaregistry

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

// IZetaRegistryErrorsMetaData contains all meta data concerning the IZetaRegistryErrors contract.
var IZetaRegistryErrorsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ChainDoesNotExist\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ContractAlreadyExists\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ContractDoesNotExist\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidContractAddress\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZRC20TokenAlreadyExists\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZRC20TokenDoesNotExist\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// IZetaRegistryErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IZetaRegistryErrorsMetaData.ABI instead.
var IZetaRegistryErrorsABI = IZetaRegistryErrorsMetaData.ABI

// IZetaRegistryErrors is an auto generated Go binding around an Ethereum contract.
type IZetaRegistryErrors struct {
	IZetaRegistryErrorsCaller     // Read-only binding to the contract
	IZetaRegistryErrorsTransactor // Write-only binding to the contract
	IZetaRegistryErrorsFilterer   // Log filterer for contract events
}

// IZetaRegistryErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZetaRegistryErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaRegistryErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZetaRegistryErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaRegistryErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZetaRegistryErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaRegistryErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZetaRegistryErrorsSession struct {
	Contract     *IZetaRegistryErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IZetaRegistryErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZetaRegistryErrorsCallerSession struct {
	Contract *IZetaRegistryErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IZetaRegistryErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZetaRegistryErrorsTransactorSession struct {
	Contract     *IZetaRegistryErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IZetaRegistryErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZetaRegistryErrorsRaw struct {
	Contract *IZetaRegistryErrors // Generic contract binding to access the raw methods on
}

// IZetaRegistryErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZetaRegistryErrorsCallerRaw struct {
	Contract *IZetaRegistryErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IZetaRegistryErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZetaRegistryErrorsTransactorRaw struct {
	Contract *IZetaRegistryErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZetaRegistryErrors creates a new instance of IZetaRegistryErrors, bound to a specific deployed contract.
func NewIZetaRegistryErrors(address common.Address, backend bind.ContractBackend) (*IZetaRegistryErrors, error) {
	contract, err := bindIZetaRegistryErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryErrors{IZetaRegistryErrorsCaller: IZetaRegistryErrorsCaller{contract: contract}, IZetaRegistryErrorsTransactor: IZetaRegistryErrorsTransactor{contract: contract}, IZetaRegistryErrorsFilterer: IZetaRegistryErrorsFilterer{contract: contract}}, nil
}

// NewIZetaRegistryErrorsCaller creates a new read-only instance of IZetaRegistryErrors, bound to a specific deployed contract.
func NewIZetaRegistryErrorsCaller(address common.Address, caller bind.ContractCaller) (*IZetaRegistryErrorsCaller, error) {
	contract, err := bindIZetaRegistryErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryErrorsCaller{contract: contract}, nil
}

// NewIZetaRegistryErrorsTransactor creates a new write-only instance of IZetaRegistryErrors, bound to a specific deployed contract.
func NewIZetaRegistryErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IZetaRegistryErrorsTransactor, error) {
	contract, err := bindIZetaRegistryErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryErrorsTransactor{contract: contract}, nil
}

// NewIZetaRegistryErrorsFilterer creates a new log filterer instance of IZetaRegistryErrors, bound to a specific deployed contract.
func NewIZetaRegistryErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IZetaRegistryErrorsFilterer, error) {
	contract, err := bindIZetaRegistryErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryErrorsFilterer{contract: contract}, nil
}

// bindIZetaRegistryErrors binds a generic wrapper to an already deployed contract.
func bindIZetaRegistryErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZetaRegistryErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaRegistryErrors *IZetaRegistryErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaRegistryErrors.Contract.IZetaRegistryErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaRegistryErrors *IZetaRegistryErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaRegistryErrors.Contract.IZetaRegistryErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaRegistryErrors *IZetaRegistryErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaRegistryErrors.Contract.IZetaRegistryErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaRegistryErrors *IZetaRegistryErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaRegistryErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaRegistryErrors *IZetaRegistryErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaRegistryErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaRegistryErrors *IZetaRegistryErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaRegistryErrors.Contract.contract.Transact(opts, method, params...)
}
