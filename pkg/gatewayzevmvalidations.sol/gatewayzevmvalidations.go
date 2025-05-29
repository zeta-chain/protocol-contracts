// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayzevmvalidations

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

// GatewayZEVMValidationsMetaData contains all meta data concerning the GatewayZEVMValidations contract.
var GatewayZEVMValidationsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MAX_MESSAGE_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"EmptyAddress\",\"inputs\":[]}]",
	Bin: "0x609b6037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80637ce1ffeb14604257806397d340f514605d575b600080fd5b604b620186a081565b60405190815260200160405180910390f35b604b6108008156fea2646970667358221220e96b63b3395af1fcdb1da412bdd2dee128b09af1a935c94cca1eafa3cd39ef9364736f6c634300081a0033",
}

// GatewayZEVMValidationsABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayZEVMValidationsMetaData.ABI instead.
var GatewayZEVMValidationsABI = GatewayZEVMValidationsMetaData.ABI

// GatewayZEVMValidationsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayZEVMValidationsMetaData.Bin instead.
var GatewayZEVMValidationsBin = GatewayZEVMValidationsMetaData.Bin

// DeployGatewayZEVMValidations deploys a new Ethereum contract, binding an instance of GatewayZEVMValidations to it.
func DeployGatewayZEVMValidations(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayZEVMValidations, error) {
	parsed, err := GatewayZEVMValidationsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayZEVMValidationsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayZEVMValidations{GatewayZEVMValidationsCaller: GatewayZEVMValidationsCaller{contract: contract}, GatewayZEVMValidationsTransactor: GatewayZEVMValidationsTransactor{contract: contract}, GatewayZEVMValidationsFilterer: GatewayZEVMValidationsFilterer{contract: contract}}, nil
}

// GatewayZEVMValidations is an auto generated Go binding around an Ethereum contract.
type GatewayZEVMValidations struct {
	GatewayZEVMValidationsCaller     // Read-only binding to the contract
	GatewayZEVMValidationsTransactor // Write-only binding to the contract
	GatewayZEVMValidationsFilterer   // Log filterer for contract events
}

// GatewayZEVMValidationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayZEVMValidationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMValidationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayZEVMValidationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMValidationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayZEVMValidationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMValidationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayZEVMValidationsSession struct {
	Contract     *GatewayZEVMValidations // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GatewayZEVMValidationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayZEVMValidationsCallerSession struct {
	Contract *GatewayZEVMValidationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// GatewayZEVMValidationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayZEVMValidationsTransactorSession struct {
	Contract     *GatewayZEVMValidationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// GatewayZEVMValidationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayZEVMValidationsRaw struct {
	Contract *GatewayZEVMValidations // Generic contract binding to access the raw methods on
}

// GatewayZEVMValidationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayZEVMValidationsCallerRaw struct {
	Contract *GatewayZEVMValidationsCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayZEVMValidationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayZEVMValidationsTransactorRaw struct {
	Contract *GatewayZEVMValidationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayZEVMValidations creates a new instance of GatewayZEVMValidations, bound to a specific deployed contract.
func NewGatewayZEVMValidations(address common.Address, backend bind.ContractBackend) (*GatewayZEVMValidations, error) {
	contract, err := bindGatewayZEVMValidations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMValidations{GatewayZEVMValidationsCaller: GatewayZEVMValidationsCaller{contract: contract}, GatewayZEVMValidationsTransactor: GatewayZEVMValidationsTransactor{contract: contract}, GatewayZEVMValidationsFilterer: GatewayZEVMValidationsFilterer{contract: contract}}, nil
}

// NewGatewayZEVMValidationsCaller creates a new read-only instance of GatewayZEVMValidations, bound to a specific deployed contract.
func NewGatewayZEVMValidationsCaller(address common.Address, caller bind.ContractCaller) (*GatewayZEVMValidationsCaller, error) {
	contract, err := bindGatewayZEVMValidations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMValidationsCaller{contract: contract}, nil
}

// NewGatewayZEVMValidationsTransactor creates a new write-only instance of GatewayZEVMValidations, bound to a specific deployed contract.
func NewGatewayZEVMValidationsTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayZEVMValidationsTransactor, error) {
	contract, err := bindGatewayZEVMValidations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMValidationsTransactor{contract: contract}, nil
}

// NewGatewayZEVMValidationsFilterer creates a new log filterer instance of GatewayZEVMValidations, bound to a specific deployed contract.
func NewGatewayZEVMValidationsFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayZEVMValidationsFilterer, error) {
	contract, err := bindGatewayZEVMValidations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMValidationsFilterer{contract: contract}, nil
}

// bindGatewayZEVMValidations binds a generic wrapper to an already deployed contract.
func bindGatewayZEVMValidations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayZEVMValidationsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMValidations *GatewayZEVMValidationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMValidations.Contract.GatewayZEVMValidationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMValidations *GatewayZEVMValidationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMValidations.Contract.GatewayZEVMValidationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMValidations *GatewayZEVMValidationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMValidations.Contract.GatewayZEVMValidationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMValidations *GatewayZEVMValidationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMValidations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMValidations *GatewayZEVMValidationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMValidations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMValidations *GatewayZEVMValidationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMValidations.Contract.contract.Transact(opts, method, params...)
}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVMValidations *GatewayZEVMValidationsCaller) MAXMESSAGESIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayZEVMValidations.contract.Call(opts, &out, "MAX_MESSAGE_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVMValidations *GatewayZEVMValidationsSession) MAXMESSAGESIZE() (*big.Int, error) {
	return _GatewayZEVMValidations.Contract.MAXMESSAGESIZE(&_GatewayZEVMValidations.CallOpts)
}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVMValidations *GatewayZEVMValidationsCallerSession) MAXMESSAGESIZE() (*big.Int, error) {
	return _GatewayZEVMValidations.Contract.MAXMESSAGESIZE(&_GatewayZEVMValidations.CallOpts)
}

// MINGASLIMIT is a free data retrieval call binding the contract method 0x7ce1ffeb.
//
// Solidity: function MIN_GAS_LIMIT() view returns(uint256)
func (_GatewayZEVMValidations *GatewayZEVMValidationsCaller) MINGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayZEVMValidations.contract.Call(opts, &out, "MIN_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINGASLIMIT is a free data retrieval call binding the contract method 0x7ce1ffeb.
//
// Solidity: function MIN_GAS_LIMIT() view returns(uint256)
func (_GatewayZEVMValidations *GatewayZEVMValidationsSession) MINGASLIMIT() (*big.Int, error) {
	return _GatewayZEVMValidations.Contract.MINGASLIMIT(&_GatewayZEVMValidations.CallOpts)
}

// MINGASLIMIT is a free data retrieval call binding the contract method 0x7ce1ffeb.
//
// Solidity: function MIN_GAS_LIMIT() view returns(uint256)
func (_GatewayZEVMValidations *GatewayZEVMValidationsCallerSession) MINGASLIMIT() (*big.Int, error) {
	return _GatewayZEVMValidations.Contract.MINGASLIMIT(&_GatewayZEVMValidations.CallOpts)
}
