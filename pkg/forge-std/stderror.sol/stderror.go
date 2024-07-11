// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stderror

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

// StdErrorMetaData contains all meta data concerning the StdError contract.
var StdErrorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"arithmeticError\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assertionError\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"divisionError\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"encodeStorageError\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enumConversionError\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexOOBError\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"memOverflowError\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"popError\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zeroVarError\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6109ec610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061009d5760003560e01c8063986c5f6811610070578063986c5f681461011a578063b22dc54d14610138578063b67689da14610156578063d160e4de14610174578063fa784a44146101925761009d565b806305ee8612146100a257806310332977146100c05780631de45560146100de5780638995290f146100fc575b600080fd5b6100aa6101b0565b6040516100b79190610792565b60405180910390f35b6100c8610242565b6040516100d59190610792565b60405180910390f35b6100e66102d4565b6040516100f39190610792565b60405180910390f35b610104610366565b6040516101119190610792565b60405180910390f35b6101226103f8565b60405161012f9190610792565b60405180910390f35b61014061048a565b60405161014d9190610792565b60405180910390f35b61015e61051c565b60405161016b9190610792565b60405180910390f35b61017c6105ae565b6040516101899190610792565b60405180910390f35b61019a610640565b6040516101a79190610792565b60405180910390f35b60326040516024016101c29190610856565b6040516020818303038152906040527f4e487b71000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505081565b600160405160240161025491906107ea565b6040516020818303038152906040527f4e487b71000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505081565b60216040516024016102e69190610805565b6040516020818303038152906040527f4e487b71000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505081565b601160405160240161037891906107b4565b6040516020818303038152906040527f4e487b71000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505081565b604160405160240161040a9190610871565b6040516020818303038152906040527f4e487b71000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505081565b603160405160240161049c919061083b565b6040516020818303038152906040527f4e487b71000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505081565b605160405160240161052e919061088c565b6040516020818303038152906040527f4e487b71000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505081565b60226040516024016105c09190610820565b6040516020818303038152906040527f4e487b71000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505081565b601260405160240161065291906107cf565b6040516020818303038152906040527f4e487b71000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505081565b60006106dd826108a7565b6106e781856108b2565b93506106f7818560208601610972565b610700816109a5565b840191505092915050565b610714816108d0565b82525050565b610723816108e2565b82525050565b610732816108f4565b82525050565b61074181610906565b82525050565b61075081610918565b82525050565b61075f8161092a565b82525050565b61076e8161093c565b82525050565b61077d8161094e565b82525050565b61078c81610960565b82525050565b600060208201905081810360008301526107ac81846106d2565b905092915050565b60006020820190506107c9600083018461070b565b92915050565b60006020820190506107e4600083018461071a565b92915050565b60006020820190506107ff6000830184610729565b92915050565b600060208201905061081a6000830184610738565b92915050565b60006020820190506108356000830184610747565b92915050565b60006020820190506108506000830184610756565b92915050565b600060208201905061086b6000830184610765565b92915050565b60006020820190506108866000830184610774565b92915050565b60006020820190506108a16000830184610783565b92915050565b600081519050919050565b600082825260208201905092915050565b600060ff82169050919050565b60006108db826108c3565b9050919050565b60006108ed826108c3565b9050919050565b60006108ff826108c3565b9050919050565b6000610911826108c3565b9050919050565b6000610923826108c3565b9050919050565b6000610935826108c3565b9050919050565b6000610947826108c3565b9050919050565b6000610959826108c3565b9050919050565b600061096b826108c3565b9050919050565b60005b83811015610990578082015181840152602081019050610975565b8381111561099f576000848401525b50505050565b6000601f19601f830116905091905056fea264697066735822122086a066b9b91d74494e9ce1d74ef0b42568574cbc8d6ec51dc4e6feec0c5260d764736f6c63430008070033",
}

// StdErrorABI is the input ABI used to generate the binding from.
// Deprecated: Use StdErrorMetaData.ABI instead.
var StdErrorABI = StdErrorMetaData.ABI

// StdErrorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StdErrorMetaData.Bin instead.
var StdErrorBin = StdErrorMetaData.Bin

// DeployStdError deploys a new Ethereum contract, binding an instance of StdError to it.
func DeployStdError(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StdError, error) {
	parsed, err := StdErrorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StdErrorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StdError{StdErrorCaller: StdErrorCaller{contract: contract}, StdErrorTransactor: StdErrorTransactor{contract: contract}, StdErrorFilterer: StdErrorFilterer{contract: contract}}, nil
}

// StdError is an auto generated Go binding around an Ethereum contract.
type StdError struct {
	StdErrorCaller     // Read-only binding to the contract
	StdErrorTransactor // Write-only binding to the contract
	StdErrorFilterer   // Log filterer for contract events
}

// StdErrorCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdErrorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdErrorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdErrorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdErrorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdErrorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdErrorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdErrorSession struct {
	Contract     *StdError         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdErrorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdErrorCallerSession struct {
	Contract *StdErrorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StdErrorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdErrorTransactorSession struct {
	Contract     *StdErrorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StdErrorRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdErrorRaw struct {
	Contract *StdError // Generic contract binding to access the raw methods on
}

// StdErrorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdErrorCallerRaw struct {
	Contract *StdErrorCaller // Generic read-only contract binding to access the raw methods on
}

// StdErrorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdErrorTransactorRaw struct {
	Contract *StdErrorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdError creates a new instance of StdError, bound to a specific deployed contract.
func NewStdError(address common.Address, backend bind.ContractBackend) (*StdError, error) {
	contract, err := bindStdError(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdError{StdErrorCaller: StdErrorCaller{contract: contract}, StdErrorTransactor: StdErrorTransactor{contract: contract}, StdErrorFilterer: StdErrorFilterer{contract: contract}}, nil
}

// NewStdErrorCaller creates a new read-only instance of StdError, bound to a specific deployed contract.
func NewStdErrorCaller(address common.Address, caller bind.ContractCaller) (*StdErrorCaller, error) {
	contract, err := bindStdError(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdErrorCaller{contract: contract}, nil
}

// NewStdErrorTransactor creates a new write-only instance of StdError, bound to a specific deployed contract.
func NewStdErrorTransactor(address common.Address, transactor bind.ContractTransactor) (*StdErrorTransactor, error) {
	contract, err := bindStdError(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdErrorTransactor{contract: contract}, nil
}

// NewStdErrorFilterer creates a new log filterer instance of StdError, bound to a specific deployed contract.
func NewStdErrorFilterer(address common.Address, filterer bind.ContractFilterer) (*StdErrorFilterer, error) {
	contract, err := bindStdError(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdErrorFilterer{contract: contract}, nil
}

// bindStdError binds a generic wrapper to an already deployed contract.
func bindStdError(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdErrorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdError *StdErrorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdError.Contract.StdErrorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdError *StdErrorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdError.Contract.StdErrorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdError *StdErrorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdError.Contract.StdErrorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdError *StdErrorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdError.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdError *StdErrorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdError.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdError *StdErrorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdError.Contract.contract.Transact(opts, method, params...)
}

// ArithmeticError is a free data retrieval call binding the contract method 0x8995290f.
//
// Solidity: function arithmeticError() view returns(bytes)
func (_StdError *StdErrorCaller) ArithmeticError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "arithmeticError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ArithmeticError is a free data retrieval call binding the contract method 0x8995290f.
//
// Solidity: function arithmeticError() view returns(bytes)
func (_StdError *StdErrorSession) ArithmeticError() ([]byte, error) {
	return _StdError.Contract.ArithmeticError(&_StdError.CallOpts)
}

// ArithmeticError is a free data retrieval call binding the contract method 0x8995290f.
//
// Solidity: function arithmeticError() view returns(bytes)
func (_StdError *StdErrorCallerSession) ArithmeticError() ([]byte, error) {
	return _StdError.Contract.ArithmeticError(&_StdError.CallOpts)
}

// AssertionError is a free data retrieval call binding the contract method 0x10332977.
//
// Solidity: function assertionError() view returns(bytes)
func (_StdError *StdErrorCaller) AssertionError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "assertionError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AssertionError is a free data retrieval call binding the contract method 0x10332977.
//
// Solidity: function assertionError() view returns(bytes)
func (_StdError *StdErrorSession) AssertionError() ([]byte, error) {
	return _StdError.Contract.AssertionError(&_StdError.CallOpts)
}

// AssertionError is a free data retrieval call binding the contract method 0x10332977.
//
// Solidity: function assertionError() view returns(bytes)
func (_StdError *StdErrorCallerSession) AssertionError() ([]byte, error) {
	return _StdError.Contract.AssertionError(&_StdError.CallOpts)
}

// DivisionError is a free data retrieval call binding the contract method 0xfa784a44.
//
// Solidity: function divisionError() view returns(bytes)
func (_StdError *StdErrorCaller) DivisionError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "divisionError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DivisionError is a free data retrieval call binding the contract method 0xfa784a44.
//
// Solidity: function divisionError() view returns(bytes)
func (_StdError *StdErrorSession) DivisionError() ([]byte, error) {
	return _StdError.Contract.DivisionError(&_StdError.CallOpts)
}

// DivisionError is a free data retrieval call binding the contract method 0xfa784a44.
//
// Solidity: function divisionError() view returns(bytes)
func (_StdError *StdErrorCallerSession) DivisionError() ([]byte, error) {
	return _StdError.Contract.DivisionError(&_StdError.CallOpts)
}

// EncodeStorageError is a free data retrieval call binding the contract method 0xd160e4de.
//
// Solidity: function encodeStorageError() view returns(bytes)
func (_StdError *StdErrorCaller) EncodeStorageError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "encodeStorageError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeStorageError is a free data retrieval call binding the contract method 0xd160e4de.
//
// Solidity: function encodeStorageError() view returns(bytes)
func (_StdError *StdErrorSession) EncodeStorageError() ([]byte, error) {
	return _StdError.Contract.EncodeStorageError(&_StdError.CallOpts)
}

// EncodeStorageError is a free data retrieval call binding the contract method 0xd160e4de.
//
// Solidity: function encodeStorageError() view returns(bytes)
func (_StdError *StdErrorCallerSession) EncodeStorageError() ([]byte, error) {
	return _StdError.Contract.EncodeStorageError(&_StdError.CallOpts)
}

// EnumConversionError is a free data retrieval call binding the contract method 0x1de45560.
//
// Solidity: function enumConversionError() view returns(bytes)
func (_StdError *StdErrorCaller) EnumConversionError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "enumConversionError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnumConversionError is a free data retrieval call binding the contract method 0x1de45560.
//
// Solidity: function enumConversionError() view returns(bytes)
func (_StdError *StdErrorSession) EnumConversionError() ([]byte, error) {
	return _StdError.Contract.EnumConversionError(&_StdError.CallOpts)
}

// EnumConversionError is a free data retrieval call binding the contract method 0x1de45560.
//
// Solidity: function enumConversionError() view returns(bytes)
func (_StdError *StdErrorCallerSession) EnumConversionError() ([]byte, error) {
	return _StdError.Contract.EnumConversionError(&_StdError.CallOpts)
}

// IndexOOBError is a free data retrieval call binding the contract method 0x05ee8612.
//
// Solidity: function indexOOBError() view returns(bytes)
func (_StdError *StdErrorCaller) IndexOOBError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "indexOOBError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// IndexOOBError is a free data retrieval call binding the contract method 0x05ee8612.
//
// Solidity: function indexOOBError() view returns(bytes)
func (_StdError *StdErrorSession) IndexOOBError() ([]byte, error) {
	return _StdError.Contract.IndexOOBError(&_StdError.CallOpts)
}

// IndexOOBError is a free data retrieval call binding the contract method 0x05ee8612.
//
// Solidity: function indexOOBError() view returns(bytes)
func (_StdError *StdErrorCallerSession) IndexOOBError() ([]byte, error) {
	return _StdError.Contract.IndexOOBError(&_StdError.CallOpts)
}

// MemOverflowError is a free data retrieval call binding the contract method 0x986c5f68.
//
// Solidity: function memOverflowError() view returns(bytes)
func (_StdError *StdErrorCaller) MemOverflowError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "memOverflowError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MemOverflowError is a free data retrieval call binding the contract method 0x986c5f68.
//
// Solidity: function memOverflowError() view returns(bytes)
func (_StdError *StdErrorSession) MemOverflowError() ([]byte, error) {
	return _StdError.Contract.MemOverflowError(&_StdError.CallOpts)
}

// MemOverflowError is a free data retrieval call binding the contract method 0x986c5f68.
//
// Solidity: function memOverflowError() view returns(bytes)
func (_StdError *StdErrorCallerSession) MemOverflowError() ([]byte, error) {
	return _StdError.Contract.MemOverflowError(&_StdError.CallOpts)
}

// PopError is a free data retrieval call binding the contract method 0xb22dc54d.
//
// Solidity: function popError() view returns(bytes)
func (_StdError *StdErrorCaller) PopError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "popError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PopError is a free data retrieval call binding the contract method 0xb22dc54d.
//
// Solidity: function popError() view returns(bytes)
func (_StdError *StdErrorSession) PopError() ([]byte, error) {
	return _StdError.Contract.PopError(&_StdError.CallOpts)
}

// PopError is a free data retrieval call binding the contract method 0xb22dc54d.
//
// Solidity: function popError() view returns(bytes)
func (_StdError *StdErrorCallerSession) PopError() ([]byte, error) {
	return _StdError.Contract.PopError(&_StdError.CallOpts)
}

// ZeroVarError is a free data retrieval call binding the contract method 0xb67689da.
//
// Solidity: function zeroVarError() view returns(bytes)
func (_StdError *StdErrorCaller) ZeroVarError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "zeroVarError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ZeroVarError is a free data retrieval call binding the contract method 0xb67689da.
//
// Solidity: function zeroVarError() view returns(bytes)
func (_StdError *StdErrorSession) ZeroVarError() ([]byte, error) {
	return _StdError.Contract.ZeroVarError(&_StdError.CallOpts)
}

// ZeroVarError is a free data retrieval call binding the contract method 0xb67689da.
//
// Solidity: function zeroVarError() view returns(bytes)
func (_StdError *StdErrorCallerSession) ZeroVarError() ([]byte, error) {
	return _StdError.Contract.ZeroVarError(&_StdError.CallOpts)
}
