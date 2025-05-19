// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iregistry

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

// IRegistryErrorsMetaData contains all meta data concerning the IRegistryErrors contract.
var IRegistryErrorsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"ChainActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainNonActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ContractAlreadyRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ContractNotFound\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidContractType\",\"inputs\":[{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotGateway\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZRC20AlreadyRegistered\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZRC20SymbolAlreadyInUse\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// IRegistryErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IRegistryErrorsMetaData.ABI instead.
var IRegistryErrorsABI = IRegistryErrorsMetaData.ABI

// IRegistryErrors is an auto generated Go binding around an Ethereum contract.
type IRegistryErrors struct {
	IRegistryErrorsCaller     // Read-only binding to the contract
	IRegistryErrorsTransactor // Write-only binding to the contract
	IRegistryErrorsFilterer   // Log filterer for contract events
}

// IRegistryErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRegistryErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistryErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRegistryErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistryErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRegistryErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistryErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRegistryErrorsSession struct {
	Contract     *IRegistryErrors  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRegistryErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRegistryErrorsCallerSession struct {
	Contract *IRegistryErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IRegistryErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRegistryErrorsTransactorSession struct {
	Contract     *IRegistryErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IRegistryErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRegistryErrorsRaw struct {
	Contract *IRegistryErrors // Generic contract binding to access the raw methods on
}

// IRegistryErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRegistryErrorsCallerRaw struct {
	Contract *IRegistryErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IRegistryErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRegistryErrorsTransactorRaw struct {
	Contract *IRegistryErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRegistryErrors creates a new instance of IRegistryErrors, bound to a specific deployed contract.
func NewIRegistryErrors(address common.Address, backend bind.ContractBackend) (*IRegistryErrors, error) {
	contract, err := bindIRegistryErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRegistryErrors{IRegistryErrorsCaller: IRegistryErrorsCaller{contract: contract}, IRegistryErrorsTransactor: IRegistryErrorsTransactor{contract: contract}, IRegistryErrorsFilterer: IRegistryErrorsFilterer{contract: contract}}, nil
}

// NewIRegistryErrorsCaller creates a new read-only instance of IRegistryErrors, bound to a specific deployed contract.
func NewIRegistryErrorsCaller(address common.Address, caller bind.ContractCaller) (*IRegistryErrorsCaller, error) {
	contract, err := bindIRegistryErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRegistryErrorsCaller{contract: contract}, nil
}

// NewIRegistryErrorsTransactor creates a new write-only instance of IRegistryErrors, bound to a specific deployed contract.
func NewIRegistryErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IRegistryErrorsTransactor, error) {
	contract, err := bindIRegistryErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRegistryErrorsTransactor{contract: contract}, nil
}

// NewIRegistryErrorsFilterer creates a new log filterer instance of IRegistryErrors, bound to a specific deployed contract.
func NewIRegistryErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IRegistryErrorsFilterer, error) {
	contract, err := bindIRegistryErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRegistryErrorsFilterer{contract: contract}, nil
}

// bindIRegistryErrors binds a generic wrapper to an already deployed contract.
func bindIRegistryErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRegistryErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRegistryErrors *IRegistryErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRegistryErrors.Contract.IRegistryErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRegistryErrors *IRegistryErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRegistryErrors.Contract.IRegistryErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRegistryErrors *IRegistryErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRegistryErrors.Contract.IRegistryErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRegistryErrors *IRegistryErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRegistryErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRegistryErrors *IRegistryErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRegistryErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRegistryErrors *IRegistryErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRegistryErrors.Contract.contract.Transact(opts, method, params...)
}
