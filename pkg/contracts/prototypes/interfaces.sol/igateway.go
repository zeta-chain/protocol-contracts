// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interfaces

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

// IGatewayMetaData contains all meta data concerning the IGateway contract.
var IGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeWithERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use IGatewayMetaData.ABI instead.
var IGatewayABI = IGatewayMetaData.ABI

// IGateway is an auto generated Go binding around an Ethereum contract.
type IGateway struct {
	IGatewayCaller     // Read-only binding to the contract
	IGatewayTransactor // Write-only binding to the contract
	IGatewayFilterer   // Log filterer for contract events
}

// IGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGatewaySession struct {
	Contract     *IGateway         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGatewayCallerSession struct {
	Contract *IGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGatewayTransactorSession struct {
	Contract     *IGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGatewayRaw struct {
	Contract *IGateway // Generic contract binding to access the raw methods on
}

// IGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGatewayCallerRaw struct {
	Contract *IGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// IGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGatewayTransactorRaw struct {
	Contract *IGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGateway creates a new instance of IGateway, bound to a specific deployed contract.
func NewIGateway(address common.Address, backend bind.ContractBackend) (*IGateway, error) {
	contract, err := bindIGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGateway{IGatewayCaller: IGatewayCaller{contract: contract}, IGatewayTransactor: IGatewayTransactor{contract: contract}, IGatewayFilterer: IGatewayFilterer{contract: contract}}, nil
}

// NewIGatewayCaller creates a new read-only instance of IGateway, bound to a specific deployed contract.
func NewIGatewayCaller(address common.Address, caller bind.ContractCaller) (*IGatewayCaller, error) {
	contract, err := bindIGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayCaller{contract: contract}, nil
}

// NewIGatewayTransactor creates a new write-only instance of IGateway, bound to a specific deployed contract.
func NewIGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*IGatewayTransactor, error) {
	contract, err := bindIGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayTransactor{contract: contract}, nil
}

// NewIGatewayFilterer creates a new log filterer instance of IGateway, bound to a specific deployed contract.
func NewIGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*IGatewayFilterer, error) {
	contract, err := bindIGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGatewayFilterer{contract: contract}, nil
}

// bindIGateway binds a generic wrapper to an already deployed contract.
func bindIGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGateway *IGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGateway.Contract.IGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGateway *IGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGateway.Contract.IGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGateway *IGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGateway.Contract.IGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGateway *IGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGateway *IGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGateway *IGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGateway.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_IGateway *IGatewayTransactor) Execute(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _IGateway.contract.Transact(opts, "execute", destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_IGateway *IGatewaySession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _IGateway.Contract.Execute(&_IGateway.TransactOpts, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_IGateway *IGatewayTransactorSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _IGateway.Contract.Execute(&_IGateway.TransactOpts, destination, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_IGateway *IGatewayTransactor) ExecuteWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IGateway.contract.Transact(opts, "executeWithERC20", token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_IGateway *IGatewaySession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IGateway.Contract.ExecuteWithERC20(&_IGateway.TransactOpts, token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_IGateway *IGatewayTransactorSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IGateway.Contract.ExecuteWithERC20(&_IGateway.TransactOpts, token, to, amount, data)
}
