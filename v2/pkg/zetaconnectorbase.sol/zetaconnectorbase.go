// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectorbase

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

// ZetaConnectorBaseMetaData contains all meta data concerning the ZetaConnectorBase contract.
var ZetaConnectorBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveTokens\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndRevert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawAndCall\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawAndRevert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// ZetaConnectorBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorBaseMetaData.ABI instead.
var ZetaConnectorBaseABI = ZetaConnectorBaseMetaData.ABI

// ZetaConnectorBase is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorBase struct {
	ZetaConnectorBaseCaller     // Read-only binding to the contract
	ZetaConnectorBaseTransactor // Write-only binding to the contract
	ZetaConnectorBaseFilterer   // Log filterer for contract events
}

// ZetaConnectorBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorBaseSession struct {
	Contract     *ZetaConnectorBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZetaConnectorBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorBaseCallerSession struct {
	Contract *ZetaConnectorBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ZetaConnectorBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorBaseTransactorSession struct {
	Contract     *ZetaConnectorBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZetaConnectorBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorBaseRaw struct {
	Contract *ZetaConnectorBase // Generic contract binding to access the raw methods on
}

// ZetaConnectorBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorBaseCallerRaw struct {
	Contract *ZetaConnectorBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorBaseTransactorRaw struct {
	Contract *ZetaConnectorBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorBase creates a new instance of ZetaConnectorBase, bound to a specific deployed contract.
func NewZetaConnectorBase(address common.Address, backend bind.ContractBackend) (*ZetaConnectorBase, error) {
	contract, err := bindZetaConnectorBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBase{ZetaConnectorBaseCaller: ZetaConnectorBaseCaller{contract: contract}, ZetaConnectorBaseTransactor: ZetaConnectorBaseTransactor{contract: contract}, ZetaConnectorBaseFilterer: ZetaConnectorBaseFilterer{contract: contract}}, nil
}

// NewZetaConnectorBaseCaller creates a new read-only instance of ZetaConnectorBase, bound to a specific deployed contract.
func NewZetaConnectorBaseCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorBaseCaller, error) {
	contract, err := bindZetaConnectorBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseCaller{contract: contract}, nil
}

// NewZetaConnectorBaseTransactor creates a new write-only instance of ZetaConnectorBase, bound to a specific deployed contract.
func NewZetaConnectorBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorBaseTransactor, error) {
	contract, err := bindZetaConnectorBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseTransactor{contract: contract}, nil
}

// NewZetaConnectorBaseFilterer creates a new log filterer instance of ZetaConnectorBase, bound to a specific deployed contract.
func NewZetaConnectorBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorBaseFilterer, error) {
	contract, err := bindZetaConnectorBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseFilterer{contract: contract}, nil
}

// bindZetaConnectorBase binds a generic wrapper to an already deployed contract.
func bindZetaConnectorBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorBase *ZetaConnectorBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorBase.Contract.ZetaConnectorBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorBase *ZetaConnectorBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.ZetaConnectorBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorBase *ZetaConnectorBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.ZetaConnectorBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorBase *ZetaConnectorBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorBase.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseSession) Gateway() (common.Address, error) {
	return _ZetaConnectorBase.Contract.Gateway(&_ZetaConnectorBase.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCallerSession) Gateway() (common.Address, error) {
	return _ZetaConnectorBase.Contract.Gateway(&_ZetaConnectorBase.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorBase.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorBase.Contract.TssAddress(&_ZetaConnectorBase.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCallerSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorBase.Contract.TssAddress(&_ZetaConnectorBase.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorBase.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorBase.Contract.ZetaToken(&_ZetaConnectorBase.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorBase.Contract.ZetaToken(&_ZetaConnectorBase.CallOpts)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactor) ReceiveTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorBase.contract.Transact(opts, "receiveTokens", amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.ReceiveTokens(&_ZetaConnectorBase.TransactOpts, amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.ReceiveTokens(&_ZetaConnectorBase.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.contract.Transact(opts, "withdraw", to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.Withdraw(&_ZetaConnectorBase.TransactOpts, to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.Withdraw(&_ZetaConnectorBase.TransactOpts, to, amount, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactor) WithdrawAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.contract.Transact(opts, "withdrawAndCall", to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.WithdrawAndCall(&_ZetaConnectorBase.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.WithdrawAndCall(&_ZetaConnectorBase.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactor) WithdrawAndRevert(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.contract.Transact(opts, "withdrawAndRevert", to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.WithdrawAndRevert(&_ZetaConnectorBase.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.WithdrawAndRevert(&_ZetaConnectorBase.TransactOpts, to, amount, data, internalSendHash)
}

// ZetaConnectorBaseWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseWithdrawIterator struct {
	Event *ZetaConnectorBaseWithdraw // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorBaseWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorBaseWithdraw)
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
		it.Event = new(ZetaConnectorBaseWithdraw)
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
func (it *ZetaConnectorBaseWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorBaseWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorBaseWithdraw represents a Withdraw event raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorBaseWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorBase.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseWithdrawIterator{contract: _ZetaConnectorBase.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ZetaConnectorBaseWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorBase.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorBaseWithdraw)
				if err := _ZetaConnectorBase.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) ParseWithdraw(log types.Log) (*ZetaConnectorBaseWithdraw, error) {
	event := new(ZetaConnectorBaseWithdraw)
	if err := _ZetaConnectorBase.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorBaseWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseWithdrawAndCallIterator struct {
	Event *ZetaConnectorBaseWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorBaseWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorBaseWithdrawAndCall)
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
		it.Event = new(ZetaConnectorBaseWithdrawAndCall)
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
func (it *ZetaConnectorBaseWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorBaseWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorBaseWithdrawAndCall represents a WithdrawAndCall event raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseWithdrawAndCall struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorBaseWithdrawAndCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorBase.contract.FilterLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseWithdrawAndCallIterator{contract: _ZetaConnectorBase.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *ZetaConnectorBaseWithdrawAndCall, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorBase.contract.WatchLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorBaseWithdrawAndCall)
				if err := _ZetaConnectorBase.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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

// ParseWithdrawAndCall is a log parse operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) ParseWithdrawAndCall(log types.Log) (*ZetaConnectorBaseWithdrawAndCall, error) {
	event := new(ZetaConnectorBaseWithdrawAndCall)
	if err := _ZetaConnectorBase.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorBaseWithdrawAndRevertIterator is returned from FilterWithdrawAndRevert and is used to iterate over the raw logs and unpacked data for WithdrawAndRevert events raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseWithdrawAndRevertIterator struct {
	Event *ZetaConnectorBaseWithdrawAndRevert // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorBaseWithdrawAndRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorBaseWithdrawAndRevert)
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
		it.Event = new(ZetaConnectorBaseWithdrawAndRevert)
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
func (it *ZetaConnectorBaseWithdrawAndRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorBaseWithdrawAndRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorBaseWithdrawAndRevert represents a WithdrawAndRevert event raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseWithdrawAndRevert struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndRevert is a free log retrieval operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) FilterWithdrawAndRevert(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorBaseWithdrawAndRevertIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorBase.contract.FilterLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseWithdrawAndRevertIterator{contract: _ZetaConnectorBase.contract, event: "WithdrawAndRevert", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndRevert is a free log subscription operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) WatchWithdrawAndRevert(opts *bind.WatchOpts, sink chan<- *ZetaConnectorBaseWithdrawAndRevert, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorBase.contract.WatchLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorBaseWithdrawAndRevert)
				if err := _ZetaConnectorBase.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
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

// ParseWithdrawAndRevert is a log parse operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) ParseWithdrawAndRevert(log types.Log) (*ZetaConnectorBaseWithdrawAndRevert, error) {
	event := new(ZetaConnectorBaseWithdrawAndRevert)
	if err := _ZetaConnectorBase.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
