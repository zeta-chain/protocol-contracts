// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package revertoptions

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

// RevertOptionsMetaData contains all meta data concerning the RevertOptions contract.
var RevertOptionsMetaData = &bind.MetaData{
	ABI: "[]",
}

// RevertOptionsABI is the input ABI used to generate the binding from.
// Deprecated: Use RevertOptionsMetaData.ABI instead.
var RevertOptionsABI = RevertOptionsMetaData.ABI

// RevertOptions is an auto generated Go binding around an Ethereum contract.
type RevertOptions struct {
	RevertOptionsCaller     // Read-only binding to the contract
	RevertOptionsTransactor // Write-only binding to the contract
	RevertOptionsFilterer   // Log filterer for contract events
}

// RevertOptionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RevertOptionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertOptionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RevertOptionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertOptionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RevertOptionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertOptionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RevertOptionsSession struct {
	Contract     *RevertOptions    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RevertOptionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RevertOptionsCallerSession struct {
	Contract *RevertOptionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RevertOptionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RevertOptionsTransactorSession struct {
	Contract     *RevertOptionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RevertOptionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RevertOptionsRaw struct {
	Contract *RevertOptions // Generic contract binding to access the raw methods on
}

// RevertOptionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RevertOptionsCallerRaw struct {
	Contract *RevertOptionsCaller // Generic read-only contract binding to access the raw methods on
}

// RevertOptionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RevertOptionsTransactorRaw struct {
	Contract *RevertOptionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRevertOptions creates a new instance of RevertOptions, bound to a specific deployed contract.
func NewRevertOptions(address common.Address, backend bind.ContractBackend) (*RevertOptions, error) {
	contract, err := bindRevertOptions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RevertOptions{RevertOptionsCaller: RevertOptionsCaller{contract: contract}, RevertOptionsTransactor: RevertOptionsTransactor{contract: contract}, RevertOptionsFilterer: RevertOptionsFilterer{contract: contract}}, nil
}

// NewRevertOptionsCaller creates a new read-only instance of RevertOptions, bound to a specific deployed contract.
func NewRevertOptionsCaller(address common.Address, caller bind.ContractCaller) (*RevertOptionsCaller, error) {
	contract, err := bindRevertOptions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RevertOptionsCaller{contract: contract}, nil
}

// NewRevertOptionsTransactor creates a new write-only instance of RevertOptions, bound to a specific deployed contract.
func NewRevertOptionsTransactor(address common.Address, transactor bind.ContractTransactor) (*RevertOptionsTransactor, error) {
	contract, err := bindRevertOptions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RevertOptionsTransactor{contract: contract}, nil
}

// NewRevertOptionsFilterer creates a new log filterer instance of RevertOptions, bound to a specific deployed contract.
func NewRevertOptionsFilterer(address common.Address, filterer bind.ContractFilterer) (*RevertOptionsFilterer, error) {
	contract, err := bindRevertOptions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RevertOptionsFilterer{contract: contract}, nil
}

// bindRevertOptions binds a generic wrapper to an already deployed contract.
func bindRevertOptions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RevertOptionsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevertOptions *RevertOptionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevertOptions.Contract.RevertOptionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevertOptions *RevertOptionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevertOptions.Contract.RevertOptionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevertOptions *RevertOptionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevertOptions.Contract.RevertOptionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevertOptions *RevertOptionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevertOptions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevertOptions *RevertOptionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevertOptions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevertOptions *RevertOptionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevertOptions.Contract.contract.Transact(opts, method, params...)
}
