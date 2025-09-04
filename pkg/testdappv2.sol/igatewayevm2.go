// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testdappv2

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

// RevertOptions2 is an auto generated low-level Go binding around an user-defined struct.
type RevertOptions2 struct {
	RevertAddress    common.Address
	CallOnRevert     bool
	AbortAddress     common.Address
	RevertMessage    []byte
	OnRevertGasLimit *big.Int
}

// IGatewayEVM2MetaData contains all meta data concerning the IGatewayEVM2 contract.
var IGatewayEVM2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"additionalActionFeeWei\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions2\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions2\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions2\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"}]",
}

// IGatewayEVM2ABI is the input ABI used to generate the binding from.
// Deprecated: Use IGatewayEVM2MetaData.ABI instead.
var IGatewayEVM2ABI = IGatewayEVM2MetaData.ABI

// IGatewayEVM2 is an auto generated Go binding around an Ethereum contract.
type IGatewayEVM2 struct {
	IGatewayEVM2Caller     // Read-only binding to the contract
	IGatewayEVM2Transactor // Write-only binding to the contract
	IGatewayEVM2Filterer   // Log filterer for contract events
}

// IGatewayEVM2Caller is an auto generated read-only Go binding around an Ethereum contract.
type IGatewayEVM2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayEVM2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IGatewayEVM2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayEVM2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGatewayEVM2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayEVM2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGatewayEVM2Session struct {
	Contract     *IGatewayEVM2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGatewayEVM2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGatewayEVM2CallerSession struct {
	Contract *IGatewayEVM2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IGatewayEVM2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGatewayEVM2TransactorSession struct {
	Contract     *IGatewayEVM2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IGatewayEVM2Raw is an auto generated low-level Go binding around an Ethereum contract.
type IGatewayEVM2Raw struct {
	Contract *IGatewayEVM2 // Generic contract binding to access the raw methods on
}

// IGatewayEVM2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGatewayEVM2CallerRaw struct {
	Contract *IGatewayEVM2Caller // Generic read-only contract binding to access the raw methods on
}

// IGatewayEVM2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGatewayEVM2TransactorRaw struct {
	Contract *IGatewayEVM2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIGatewayEVM2 creates a new instance of IGatewayEVM2, bound to a specific deployed contract.
func NewIGatewayEVM2(address common.Address, backend bind.ContractBackend) (*IGatewayEVM2, error) {
	contract, err := bindIGatewayEVM2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVM2{IGatewayEVM2Caller: IGatewayEVM2Caller{contract: contract}, IGatewayEVM2Transactor: IGatewayEVM2Transactor{contract: contract}, IGatewayEVM2Filterer: IGatewayEVM2Filterer{contract: contract}}, nil
}

// NewIGatewayEVM2Caller creates a new read-only instance of IGatewayEVM2, bound to a specific deployed contract.
func NewIGatewayEVM2Caller(address common.Address, caller bind.ContractCaller) (*IGatewayEVM2Caller, error) {
	contract, err := bindIGatewayEVM2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVM2Caller{contract: contract}, nil
}

// NewIGatewayEVM2Transactor creates a new write-only instance of IGatewayEVM2, bound to a specific deployed contract.
func NewIGatewayEVM2Transactor(address common.Address, transactor bind.ContractTransactor) (*IGatewayEVM2Transactor, error) {
	contract, err := bindIGatewayEVM2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVM2Transactor{contract: contract}, nil
}

// NewIGatewayEVM2Filterer creates a new log filterer instance of IGatewayEVM2, bound to a specific deployed contract.
func NewIGatewayEVM2Filterer(address common.Address, filterer bind.ContractFilterer) (*IGatewayEVM2Filterer, error) {
	contract, err := bindIGatewayEVM2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVM2Filterer{contract: contract}, nil
}

// bindIGatewayEVM2 binds a generic wrapper to an already deployed contract.
func bindIGatewayEVM2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGatewayEVM2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayEVM2 *IGatewayEVM2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayEVM2.Contract.IGatewayEVM2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayEVM2 *IGatewayEVM2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayEVM2.Contract.IGatewayEVM2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayEVM2 *IGatewayEVM2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayEVM2.Contract.IGatewayEVM2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayEVM2 *IGatewayEVM2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayEVM2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayEVM2 *IGatewayEVM2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayEVM2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayEVM2 *IGatewayEVM2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayEVM2.Contract.contract.Transact(opts, method, params...)
}

// AdditionalActionFeeWei is a paid mutator transaction binding the contract method 0xb0107214.
//
// Solidity: function additionalActionFeeWei() returns(uint256)
func (_IGatewayEVM2 *IGatewayEVM2Transactor) AdditionalActionFeeWei(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayEVM2.contract.Transact(opts, "additionalActionFeeWei")
}

// AdditionalActionFeeWei is a paid mutator transaction binding the contract method 0xb0107214.
//
// Solidity: function additionalActionFeeWei() returns(uint256)
func (_IGatewayEVM2 *IGatewayEVM2Session) AdditionalActionFeeWei() (*types.Transaction, error) {
	return _IGatewayEVM2.Contract.AdditionalActionFeeWei(&_IGatewayEVM2.TransactOpts)
}

// AdditionalActionFeeWei is a paid mutator transaction binding the contract method 0xb0107214.
//
// Solidity: function additionalActionFeeWei() returns(uint256)
func (_IGatewayEVM2 *IGatewayEVM2TransactorSession) AdditionalActionFeeWei() (*types.Transaction, error) {
	return _IGatewayEVM2.Contract.AdditionalActionFeeWei(&_IGatewayEVM2.TransactOpts)
}

// Call is a paid mutator transaction binding the contract method 0x1becceb4.
//
// Solidity: function call(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_IGatewayEVM2 *IGatewayEVM2Transactor) Call(opts *bind.TransactOpts, receiver common.Address, payload []byte, revertOptions RevertOptions2) (*types.Transaction, error) {
	return _IGatewayEVM2.contract.Transact(opts, "call", receiver, payload, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x1becceb4.
//
// Solidity: function call(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_IGatewayEVM2 *IGatewayEVM2Session) Call(receiver common.Address, payload []byte, revertOptions RevertOptions2) (*types.Transaction, error) {
	return _IGatewayEVM2.Contract.Call(&_IGatewayEVM2.TransactOpts, receiver, payload, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x1becceb4.
//
// Solidity: function call(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_IGatewayEVM2 *IGatewayEVM2TransactorSession) Call(receiver common.Address, payload []byte, revertOptions RevertOptions2) (*types.Transaction, error) {
	return _IGatewayEVM2.Contract.Call(&_IGatewayEVM2.TransactOpts, receiver, payload, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0x726ac97c.
//
// Solidity: function deposit(address receiver, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_IGatewayEVM2 *IGatewayEVM2Transactor) Deposit(opts *bind.TransactOpts, receiver common.Address, revertOptions RevertOptions2) (*types.Transaction, error) {
	return _IGatewayEVM2.contract.Transact(opts, "deposit", receiver, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0x726ac97c.
//
// Solidity: function deposit(address receiver, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_IGatewayEVM2 *IGatewayEVM2Session) Deposit(receiver common.Address, revertOptions RevertOptions2) (*types.Transaction, error) {
	return _IGatewayEVM2.Contract.Deposit(&_IGatewayEVM2.TransactOpts, receiver, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0x726ac97c.
//
// Solidity: function deposit(address receiver, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_IGatewayEVM2 *IGatewayEVM2TransactorSession) Deposit(receiver common.Address, revertOptions RevertOptions2) (*types.Transaction, error) {
	return _IGatewayEVM2.Contract.Deposit(&_IGatewayEVM2.TransactOpts, receiver, revertOptions)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x744b9b8b.
//
// Solidity: function depositAndCall(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_IGatewayEVM2 *IGatewayEVM2Transactor) DepositAndCall(opts *bind.TransactOpts, receiver common.Address, payload []byte, revertOptions RevertOptions2) (*types.Transaction, error) {
	return _IGatewayEVM2.contract.Transact(opts, "depositAndCall", receiver, payload, revertOptions)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x744b9b8b.
//
// Solidity: function depositAndCall(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_IGatewayEVM2 *IGatewayEVM2Session) DepositAndCall(receiver common.Address, payload []byte, revertOptions RevertOptions2) (*types.Transaction, error) {
	return _IGatewayEVM2.Contract.DepositAndCall(&_IGatewayEVM2.TransactOpts, receiver, payload, revertOptions)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x744b9b8b.
//
// Solidity: function depositAndCall(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_IGatewayEVM2 *IGatewayEVM2TransactorSession) DepositAndCall(receiver common.Address, payload []byte, revertOptions RevertOptions2) (*types.Transaction, error) {
	return _IGatewayEVM2.Contract.DepositAndCall(&_IGatewayEVM2.TransactOpts, receiver, payload, revertOptions)
}
