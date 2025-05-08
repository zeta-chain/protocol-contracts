// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izetanonnative

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

// IZetaNonNativeMetaData contains all meta data concerning the IZetaNonNative contract.
var IZetaNonNativeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"mintee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// IZetaNonNativeABI is the input ABI used to generate the binding from.
// Deprecated: Use IZetaNonNativeMetaData.ABI instead.
var IZetaNonNativeABI = IZetaNonNativeMetaData.ABI

// IZetaNonNative is an auto generated Go binding around an Ethereum contract.
type IZetaNonNative struct {
	IZetaNonNativeCaller     // Read-only binding to the contract
	IZetaNonNativeTransactor // Write-only binding to the contract
	IZetaNonNativeFilterer   // Log filterer for contract events
}

// IZetaNonNativeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZetaNonNativeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaNonNativeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZetaNonNativeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaNonNativeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZetaNonNativeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaNonNativeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZetaNonNativeSession struct {
	Contract     *IZetaNonNative   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZetaNonNativeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZetaNonNativeCallerSession struct {
	Contract *IZetaNonNativeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IZetaNonNativeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZetaNonNativeTransactorSession struct {
	Contract     *IZetaNonNativeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IZetaNonNativeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZetaNonNativeRaw struct {
	Contract *IZetaNonNative // Generic contract binding to access the raw methods on
}

// IZetaNonNativeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZetaNonNativeCallerRaw struct {
	Contract *IZetaNonNativeCaller // Generic read-only contract binding to access the raw methods on
}

// IZetaNonNativeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZetaNonNativeTransactorRaw struct {
	Contract *IZetaNonNativeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZetaNonNative creates a new instance of IZetaNonNative, bound to a specific deployed contract.
func NewIZetaNonNative(address common.Address, backend bind.ContractBackend) (*IZetaNonNative, error) {
	contract, err := bindIZetaNonNative(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZetaNonNative{IZetaNonNativeCaller: IZetaNonNativeCaller{contract: contract}, IZetaNonNativeTransactor: IZetaNonNativeTransactor{contract: contract}, IZetaNonNativeFilterer: IZetaNonNativeFilterer{contract: contract}}, nil
}

// NewIZetaNonNativeCaller creates a new read-only instance of IZetaNonNative, bound to a specific deployed contract.
func NewIZetaNonNativeCaller(address common.Address, caller bind.ContractCaller) (*IZetaNonNativeCaller, error) {
	contract, err := bindIZetaNonNative(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaNonNativeCaller{contract: contract}, nil
}

// NewIZetaNonNativeTransactor creates a new write-only instance of IZetaNonNative, bound to a specific deployed contract.
func NewIZetaNonNativeTransactor(address common.Address, transactor bind.ContractTransactor) (*IZetaNonNativeTransactor, error) {
	contract, err := bindIZetaNonNative(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaNonNativeTransactor{contract: contract}, nil
}

// NewIZetaNonNativeFilterer creates a new log filterer instance of IZetaNonNative, bound to a specific deployed contract.
func NewIZetaNonNativeFilterer(address common.Address, filterer bind.ContractFilterer) (*IZetaNonNativeFilterer, error) {
	contract, err := bindIZetaNonNative(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZetaNonNativeFilterer{contract: contract}, nil
}

// bindIZetaNonNative binds a generic wrapper to an already deployed contract.
func bindIZetaNonNative(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZetaNonNativeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaNonNative *IZetaNonNativeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaNonNative.Contract.IZetaNonNativeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaNonNative *IZetaNonNativeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.IZetaNonNativeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaNonNative *IZetaNonNativeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.IZetaNonNativeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaNonNative *IZetaNonNativeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaNonNative.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaNonNative *IZetaNonNativeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaNonNative *IZetaNonNativeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZetaNonNative *IZetaNonNativeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IZetaNonNative.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZetaNonNative *IZetaNonNativeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IZetaNonNative.Contract.Allowance(&_IZetaNonNative.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZetaNonNative *IZetaNonNativeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IZetaNonNative.Contract.Allowance(&_IZetaNonNative.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZetaNonNative *IZetaNonNativeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IZetaNonNative.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZetaNonNative *IZetaNonNativeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IZetaNonNative.Contract.BalanceOf(&_IZetaNonNative.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZetaNonNative *IZetaNonNativeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IZetaNonNative.Contract.BalanceOf(&_IZetaNonNative.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZetaNonNative *IZetaNonNativeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZetaNonNative.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZetaNonNative *IZetaNonNativeSession) TotalSupply() (*big.Int, error) {
	return _IZetaNonNative.Contract.TotalSupply(&_IZetaNonNative.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZetaNonNative *IZetaNonNativeCallerSession) TotalSupply() (*big.Int, error) {
	return _IZetaNonNative.Contract.TotalSupply(&_IZetaNonNative.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IZetaNonNative *IZetaNonNativeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonNative.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IZetaNonNative *IZetaNonNativeSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.Approve(&_IZetaNonNative.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IZetaNonNative *IZetaNonNativeTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.Approve(&_IZetaNonNative.TransactOpts, spender, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IZetaNonNative *IZetaNonNativeTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZetaNonNative.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IZetaNonNative *IZetaNonNativeSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.BurnFrom(&_IZetaNonNative.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IZetaNonNative *IZetaNonNativeTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.BurnFrom(&_IZetaNonNative.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_IZetaNonNative *IZetaNonNativeTransactor) Mint(opts *bind.TransactOpts, mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _IZetaNonNative.contract.Transact(opts, "mint", mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_IZetaNonNative *IZetaNonNativeSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.Mint(&_IZetaNonNative.TransactOpts, mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_IZetaNonNative *IZetaNonNativeTransactorSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.Mint(&_IZetaNonNative.TransactOpts, mintee, value, internalSendHash)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IZetaNonNative *IZetaNonNativeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonNative.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IZetaNonNative *IZetaNonNativeSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.Transfer(&_IZetaNonNative.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IZetaNonNative *IZetaNonNativeTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.Transfer(&_IZetaNonNative.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IZetaNonNative *IZetaNonNativeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonNative.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IZetaNonNative *IZetaNonNativeSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.TransferFrom(&_IZetaNonNative.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IZetaNonNative *IZetaNonNativeTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonNative.Contract.TransferFrom(&_IZetaNonNative.TransactOpts, from, to, value)
}

// IZetaNonNativeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IZetaNonNative contract.
type IZetaNonNativeApprovalIterator struct {
	Event *IZetaNonNativeApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IZetaNonNativeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaNonNativeApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IZetaNonNativeApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IZetaNonNativeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaNonNativeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaNonNativeApproval represents a Approval event raised by the IZetaNonNative contract.
type IZetaNonNativeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IZetaNonNative *IZetaNonNativeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IZetaNonNativeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IZetaNonNative.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IZetaNonNativeApprovalIterator{contract: _IZetaNonNative.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IZetaNonNative *IZetaNonNativeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IZetaNonNativeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IZetaNonNative.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaNonNativeApproval)
				if err := _IZetaNonNative.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IZetaNonNative *IZetaNonNativeFilterer) ParseApproval(log types.Log) (*IZetaNonNativeApproval, error) {
	event := new(IZetaNonNativeApproval)
	if err := _IZetaNonNative.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaNonNativeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IZetaNonNative contract.
type IZetaNonNativeTransferIterator struct {
	Event *IZetaNonNativeTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IZetaNonNativeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaNonNativeTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IZetaNonNativeTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IZetaNonNativeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaNonNativeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaNonNativeTransfer represents a Transfer event raised by the IZetaNonNative contract.
type IZetaNonNativeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IZetaNonNative *IZetaNonNativeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IZetaNonNativeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaNonNative.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IZetaNonNativeTransferIterator{contract: _IZetaNonNative.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IZetaNonNative *IZetaNonNativeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IZetaNonNativeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaNonNative.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaNonNativeTransfer)
				if err := _IZetaNonNative.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IZetaNonNative *IZetaNonNativeFilterer) ParseTransfer(log types.Log) (*IZetaNonNativeTransfer, error) {
	event := new(IZetaNonNativeTransfer)
	if err := _IZetaNonNative.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
