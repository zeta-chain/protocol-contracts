// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc20custody

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

// IERC20CustodyEventsMetaData contains all meta data concerning the IERC20CustodyEvents contract.
var IERC20CustodyEventsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawAndCall\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawAndRevert\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// IERC20CustodyEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20CustodyEventsMetaData.ABI instead.
var IERC20CustodyEventsABI = IERC20CustodyEventsMetaData.ABI

// IERC20CustodyEvents is an auto generated Go binding around an Ethereum contract.
type IERC20CustodyEvents struct {
	IERC20CustodyEventsCaller     // Read-only binding to the contract
	IERC20CustodyEventsTransactor // Write-only binding to the contract
	IERC20CustodyEventsFilterer   // Log filterer for contract events
}

// IERC20CustodyEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20CustodyEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20CustodyEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20CustodyEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20CustodyEventsSession struct {
	Contract     *IERC20CustodyEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IERC20CustodyEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CustodyEventsCallerSession struct {
	Contract *IERC20CustodyEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IERC20CustodyEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20CustodyEventsTransactorSession struct {
	Contract     *IERC20CustodyEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IERC20CustodyEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20CustodyEventsRaw struct {
	Contract *IERC20CustodyEvents // Generic contract binding to access the raw methods on
}

// IERC20CustodyEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CustodyEventsCallerRaw struct {
	Contract *IERC20CustodyEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20CustodyEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20CustodyEventsTransactorRaw struct {
	Contract *IERC20CustodyEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20CustodyEvents creates a new instance of IERC20CustodyEvents, bound to a specific deployed contract.
func NewIERC20CustodyEvents(address common.Address, backend bind.ContractBackend) (*IERC20CustodyEvents, error) {
	contract, err := bindIERC20CustodyEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEvents{IERC20CustodyEventsCaller: IERC20CustodyEventsCaller{contract: contract}, IERC20CustodyEventsTransactor: IERC20CustodyEventsTransactor{contract: contract}, IERC20CustodyEventsFilterer: IERC20CustodyEventsFilterer{contract: contract}}, nil
}

// NewIERC20CustodyEventsCaller creates a new read-only instance of IERC20CustodyEvents, bound to a specific deployed contract.
func NewIERC20CustodyEventsCaller(address common.Address, caller bind.ContractCaller) (*IERC20CustodyEventsCaller, error) {
	contract, err := bindIERC20CustodyEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsCaller{contract: contract}, nil
}

// NewIERC20CustodyEventsTransactor creates a new write-only instance of IERC20CustodyEvents, bound to a specific deployed contract.
func NewIERC20CustodyEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20CustodyEventsTransactor, error) {
	contract, err := bindIERC20CustodyEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsTransactor{contract: contract}, nil
}

// NewIERC20CustodyEventsFilterer creates a new log filterer instance of IERC20CustodyEvents, bound to a specific deployed contract.
func NewIERC20CustodyEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20CustodyEventsFilterer, error) {
	contract, err := bindIERC20CustodyEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsFilterer{contract: contract}, nil
}

// bindIERC20CustodyEvents binds a generic wrapper to an already deployed contract.
func bindIERC20CustodyEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20CustodyEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20CustodyEvents *IERC20CustodyEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20CustodyEvents.Contract.IERC20CustodyEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20CustodyEvents *IERC20CustodyEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20CustodyEvents.Contract.IERC20CustodyEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20CustodyEvents *IERC20CustodyEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20CustodyEvents.Contract.IERC20CustodyEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20CustodyEvents *IERC20CustodyEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20CustodyEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20CustodyEvents *IERC20CustodyEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20CustodyEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20CustodyEvents *IERC20CustodyEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20CustodyEvents.Contract.contract.Transact(opts, method, params...)
}

// IERC20CustodyEventsWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWithdrawIterator struct {
	Event *IERC20CustodyEventsWithdraw // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyEventsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyEventsWithdraw)
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
		it.Event = new(IERC20CustodyEventsWithdraw)
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
func (it *IERC20CustodyEventsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyEventsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyEventsWithdraw represents a Withdraw event raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWithdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*IERC20CustodyEventsWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.FilterLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsWithdrawIterator{contract: _IERC20CustodyEvents.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IERC20CustodyEventsWithdraw, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.WatchLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyEventsWithdraw)
				if err := _IERC20CustodyEvents.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) ParseWithdraw(log types.Log) (*IERC20CustodyEventsWithdraw, error) {
	event := new(IERC20CustodyEventsWithdraw)
	if err := _IERC20CustodyEvents.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyEventsWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWithdrawAndCallIterator struct {
	Event *IERC20CustodyEventsWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyEventsWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyEventsWithdrawAndCall)
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
		it.Event = new(IERC20CustodyEventsWithdrawAndCall)
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
func (it *IERC20CustodyEventsWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyEventsWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyEventsWithdrawAndCall represents a WithdrawAndCall event raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWithdrawAndCall struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e.
//
// Solidity: event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*IERC20CustodyEventsWithdrawAndCallIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.FilterLogs(opts, "WithdrawAndCall", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsWithdrawAndCallIterator{contract: _IERC20CustodyEvents.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e.
//
// Solidity: event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *IERC20CustodyEventsWithdrawAndCall, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.WatchLogs(opts, "WithdrawAndCall", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyEventsWithdrawAndCall)
				if err := _IERC20CustodyEvents.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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

// ParseWithdrawAndCall is a log parse operation binding the contract event 0x85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e.
//
// Solidity: event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) ParseWithdrawAndCall(log types.Log) (*IERC20CustodyEventsWithdrawAndCall, error) {
	event := new(IERC20CustodyEventsWithdrawAndCall)
	if err := _IERC20CustodyEvents.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyEventsWithdrawAndRevertIterator is returned from FilterWithdrawAndRevert and is used to iterate over the raw logs and unpacked data for WithdrawAndRevert events raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWithdrawAndRevertIterator struct {
	Event *IERC20CustodyEventsWithdrawAndRevert // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyEventsWithdrawAndRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyEventsWithdrawAndRevert)
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
		it.Event = new(IERC20CustodyEventsWithdrawAndRevert)
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
func (it *IERC20CustodyEventsWithdrawAndRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyEventsWithdrawAndRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyEventsWithdrawAndRevert represents a WithdrawAndRevert event raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWithdrawAndRevert struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndRevert is a free log retrieval operation binding the contract event 0xb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c8.
//
// Solidity: event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) FilterWithdrawAndRevert(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*IERC20CustodyEventsWithdrawAndRevertIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.FilterLogs(opts, "WithdrawAndRevert", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsWithdrawAndRevertIterator{contract: _IERC20CustodyEvents.contract, event: "WithdrawAndRevert", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndRevert is a free log subscription operation binding the contract event 0xb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c8.
//
// Solidity: event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) WatchWithdrawAndRevert(opts *bind.WatchOpts, sink chan<- *IERC20CustodyEventsWithdrawAndRevert, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.WatchLogs(opts, "WithdrawAndRevert", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyEventsWithdrawAndRevert)
				if err := _IERC20CustodyEvents.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
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

// ParseWithdrawAndRevert is a log parse operation binding the contract event 0xb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c8.
//
// Solidity: event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) ParseWithdrawAndRevert(log types.Log) (*IERC20CustodyEventsWithdrawAndRevert, error) {
	event := new(IERC20CustodyEventsWithdrawAndRevert)
	if err := _IERC20CustodyEvents.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
