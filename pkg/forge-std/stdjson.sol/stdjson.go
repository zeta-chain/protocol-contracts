// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stdjson

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

// StdJsonMetaData contains all meta data concerning the StdJson contract.
var StdJsonMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fd118c1fcd3907eb874182e9003db5246bd358195bf6b0e544e419dc1d2826d964736f6c63430008070033",
}

// StdJsonABI is the input ABI used to generate the binding from.
// Deprecated: Use StdJsonMetaData.ABI instead.
var StdJsonABI = StdJsonMetaData.ABI

// StdJsonBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StdJsonMetaData.Bin instead.
var StdJsonBin = StdJsonMetaData.Bin

// DeployStdJson deploys a new Ethereum contract, binding an instance of StdJson to it.
func DeployStdJson(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StdJson, error) {
	parsed, err := StdJsonMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StdJsonBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StdJson{StdJsonCaller: StdJsonCaller{contract: contract}, StdJsonTransactor: StdJsonTransactor{contract: contract}, StdJsonFilterer: StdJsonFilterer{contract: contract}}, nil
}

// StdJson is an auto generated Go binding around an Ethereum contract.
type StdJson struct {
	StdJsonCaller     // Read-only binding to the contract
	StdJsonTransactor // Write-only binding to the contract
	StdJsonFilterer   // Log filterer for contract events
}

// StdJsonCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdJsonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdJsonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdJsonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdJsonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdJsonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdJsonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdJsonSession struct {
	Contract     *StdJson          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdJsonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdJsonCallerSession struct {
	Contract *StdJsonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StdJsonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdJsonTransactorSession struct {
	Contract     *StdJsonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StdJsonRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdJsonRaw struct {
	Contract *StdJson // Generic contract binding to access the raw methods on
}

// StdJsonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdJsonCallerRaw struct {
	Contract *StdJsonCaller // Generic read-only contract binding to access the raw methods on
}

// StdJsonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdJsonTransactorRaw struct {
	Contract *StdJsonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdJson creates a new instance of StdJson, bound to a specific deployed contract.
func NewStdJson(address common.Address, backend bind.ContractBackend) (*StdJson, error) {
	contract, err := bindStdJson(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdJson{StdJsonCaller: StdJsonCaller{contract: contract}, StdJsonTransactor: StdJsonTransactor{contract: contract}, StdJsonFilterer: StdJsonFilterer{contract: contract}}, nil
}

// NewStdJsonCaller creates a new read-only instance of StdJson, bound to a specific deployed contract.
func NewStdJsonCaller(address common.Address, caller bind.ContractCaller) (*StdJsonCaller, error) {
	contract, err := bindStdJson(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdJsonCaller{contract: contract}, nil
}

// NewStdJsonTransactor creates a new write-only instance of StdJson, bound to a specific deployed contract.
func NewStdJsonTransactor(address common.Address, transactor bind.ContractTransactor) (*StdJsonTransactor, error) {
	contract, err := bindStdJson(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdJsonTransactor{contract: contract}, nil
}

// NewStdJsonFilterer creates a new log filterer instance of StdJson, bound to a specific deployed contract.
func NewStdJsonFilterer(address common.Address, filterer bind.ContractFilterer) (*StdJsonFilterer, error) {
	contract, err := bindStdJson(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdJsonFilterer{contract: contract}, nil
}

// bindStdJson binds a generic wrapper to an already deployed contract.
func bindStdJson(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdJsonMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdJson *StdJsonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdJson.Contract.StdJsonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdJson *StdJsonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdJson.Contract.StdJsonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdJson *StdJsonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdJson.Contract.StdJsonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdJson *StdJsonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdJson.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdJson *StdJsonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdJson.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdJson *StdJsonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdJson.Contract.contract.Transact(opts, method, params...)
}
