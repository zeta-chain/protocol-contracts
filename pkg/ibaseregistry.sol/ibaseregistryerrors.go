// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibaseregistry

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

// IBaseRegistryErrorsMetaData contains all meta data concerning the IBaseRegistryErrors contract.
var IBaseRegistryErrorsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"oldAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistryManagerChanged\",\"inputs\":[{\"name\":\"oldRegistryManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRegistryManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ChainActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainNonActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ContractAlreadyRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ContractNotFound\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidContractType\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20AlreadyRegistered\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZRC20SymbolAlreadyInUse\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// IBaseRegistryErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IBaseRegistryErrorsMetaData.ABI instead.
var IBaseRegistryErrorsABI = IBaseRegistryErrorsMetaData.ABI

// IBaseRegistryErrors is an auto generated Go binding around an Ethereum contract.
type IBaseRegistryErrors struct {
	IBaseRegistryErrorsCaller     // Read-only binding to the contract
	IBaseRegistryErrorsTransactor // Write-only binding to the contract
	IBaseRegistryErrorsFilterer   // Log filterer for contract events
}

// IBaseRegistryErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBaseRegistryErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBaseRegistryErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBaseRegistryErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBaseRegistryErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBaseRegistryErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBaseRegistryErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBaseRegistryErrorsSession struct {
	Contract     *IBaseRegistryErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IBaseRegistryErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBaseRegistryErrorsCallerSession struct {
	Contract *IBaseRegistryErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IBaseRegistryErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBaseRegistryErrorsTransactorSession struct {
	Contract     *IBaseRegistryErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IBaseRegistryErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBaseRegistryErrorsRaw struct {
	Contract *IBaseRegistryErrors // Generic contract binding to access the raw methods on
}

// IBaseRegistryErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBaseRegistryErrorsCallerRaw struct {
	Contract *IBaseRegistryErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IBaseRegistryErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBaseRegistryErrorsTransactorRaw struct {
	Contract *IBaseRegistryErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBaseRegistryErrors creates a new instance of IBaseRegistryErrors, bound to a specific deployed contract.
func NewIBaseRegistryErrors(address common.Address, backend bind.ContractBackend) (*IBaseRegistryErrors, error) {
	contract, err := bindIBaseRegistryErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBaseRegistryErrors{IBaseRegistryErrorsCaller: IBaseRegistryErrorsCaller{contract: contract}, IBaseRegistryErrorsTransactor: IBaseRegistryErrorsTransactor{contract: contract}, IBaseRegistryErrorsFilterer: IBaseRegistryErrorsFilterer{contract: contract}}, nil
}

// NewIBaseRegistryErrorsCaller creates a new read-only instance of IBaseRegistryErrors, bound to a specific deployed contract.
func NewIBaseRegistryErrorsCaller(address common.Address, caller bind.ContractCaller) (*IBaseRegistryErrorsCaller, error) {
	contract, err := bindIBaseRegistryErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBaseRegistryErrorsCaller{contract: contract}, nil
}

// NewIBaseRegistryErrorsTransactor creates a new write-only instance of IBaseRegistryErrors, bound to a specific deployed contract.
func NewIBaseRegistryErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IBaseRegistryErrorsTransactor, error) {
	contract, err := bindIBaseRegistryErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBaseRegistryErrorsTransactor{contract: contract}, nil
}

// NewIBaseRegistryErrorsFilterer creates a new log filterer instance of IBaseRegistryErrors, bound to a specific deployed contract.
func NewIBaseRegistryErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IBaseRegistryErrorsFilterer, error) {
	contract, err := bindIBaseRegistryErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBaseRegistryErrorsFilterer{contract: contract}, nil
}

// bindIBaseRegistryErrors binds a generic wrapper to an already deployed contract.
func bindIBaseRegistryErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBaseRegistryErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBaseRegistryErrors *IBaseRegistryErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBaseRegistryErrors.Contract.IBaseRegistryErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBaseRegistryErrors *IBaseRegistryErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBaseRegistryErrors.Contract.IBaseRegistryErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBaseRegistryErrors *IBaseRegistryErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBaseRegistryErrors.Contract.IBaseRegistryErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBaseRegistryErrors *IBaseRegistryErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBaseRegistryErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBaseRegistryErrors *IBaseRegistryErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBaseRegistryErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBaseRegistryErrors *IBaseRegistryErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBaseRegistryErrors.Contract.contract.Transact(opts, method, params...)
}

// IBaseRegistryErrorsAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the IBaseRegistryErrors contract.
type IBaseRegistryErrorsAdminChangedIterator struct {
	Event *IBaseRegistryErrorsAdminChanged // Event containing the contract specifics and raw log

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
func (it *IBaseRegistryErrorsAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBaseRegistryErrorsAdminChanged)
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
		it.Event = new(IBaseRegistryErrorsAdminChanged)
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
func (it *IBaseRegistryErrorsAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBaseRegistryErrorsAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBaseRegistryErrorsAdminChanged represents a AdminChanged event raised by the IBaseRegistryErrors contract.
type IBaseRegistryErrorsAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_IBaseRegistryErrors *IBaseRegistryErrorsFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*IBaseRegistryErrorsAdminChangedIterator, error) {

	logs, sub, err := _IBaseRegistryErrors.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &IBaseRegistryErrorsAdminChangedIterator{contract: _IBaseRegistryErrors.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_IBaseRegistryErrors *IBaseRegistryErrorsFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *IBaseRegistryErrorsAdminChanged) (event.Subscription, error) {

	logs, sub, err := _IBaseRegistryErrors.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBaseRegistryErrorsAdminChanged)
				if err := _IBaseRegistryErrors.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_IBaseRegistryErrors *IBaseRegistryErrorsFilterer) ParseAdminChanged(log types.Log) (*IBaseRegistryErrorsAdminChanged, error) {
	event := new(IBaseRegistryErrorsAdminChanged)
	if err := _IBaseRegistryErrors.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBaseRegistryErrorsRegistryManagerChangedIterator is returned from FilterRegistryManagerChanged and is used to iterate over the raw logs and unpacked data for RegistryManagerChanged events raised by the IBaseRegistryErrors contract.
type IBaseRegistryErrorsRegistryManagerChangedIterator struct {
	Event *IBaseRegistryErrorsRegistryManagerChanged // Event containing the contract specifics and raw log

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
func (it *IBaseRegistryErrorsRegistryManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBaseRegistryErrorsRegistryManagerChanged)
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
		it.Event = new(IBaseRegistryErrorsRegistryManagerChanged)
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
func (it *IBaseRegistryErrorsRegistryManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBaseRegistryErrorsRegistryManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBaseRegistryErrorsRegistryManagerChanged represents a RegistryManagerChanged event raised by the IBaseRegistryErrors contract.
type IBaseRegistryErrorsRegistryManagerChanged struct {
	OldRegistryManager common.Address
	NewRegistryManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRegistryManagerChanged is a free log retrieval operation binding the contract event 0x6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279.
//
// Solidity: event RegistryManagerChanged(address oldRegistryManager, address newRegistryManager)
func (_IBaseRegistryErrors *IBaseRegistryErrorsFilterer) FilterRegistryManagerChanged(opts *bind.FilterOpts) (*IBaseRegistryErrorsRegistryManagerChangedIterator, error) {

	logs, sub, err := _IBaseRegistryErrors.contract.FilterLogs(opts, "RegistryManagerChanged")
	if err != nil {
		return nil, err
	}
	return &IBaseRegistryErrorsRegistryManagerChangedIterator{contract: _IBaseRegistryErrors.contract, event: "RegistryManagerChanged", logs: logs, sub: sub}, nil
}

// WatchRegistryManagerChanged is a free log subscription operation binding the contract event 0x6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279.
//
// Solidity: event RegistryManagerChanged(address oldRegistryManager, address newRegistryManager)
func (_IBaseRegistryErrors *IBaseRegistryErrorsFilterer) WatchRegistryManagerChanged(opts *bind.WatchOpts, sink chan<- *IBaseRegistryErrorsRegistryManagerChanged) (event.Subscription, error) {

	logs, sub, err := _IBaseRegistryErrors.contract.WatchLogs(opts, "RegistryManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBaseRegistryErrorsRegistryManagerChanged)
				if err := _IBaseRegistryErrors.contract.UnpackLog(event, "RegistryManagerChanged", log); err != nil {
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

// ParseRegistryManagerChanged is a log parse operation binding the contract event 0x6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279.
//
// Solidity: event RegistryManagerChanged(address oldRegistryManager, address newRegistryManager)
func (_IBaseRegistryErrors *IBaseRegistryErrorsFilterer) ParseRegistryManagerChanged(log types.Log) (*IBaseRegistryErrorsRegistryManagerChanged, error) {
	event := new(IBaseRegistryErrorsRegistryManagerChanged)
	if err := _IBaseRegistryErrors.contract.UnpackLog(event, "RegistryManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
