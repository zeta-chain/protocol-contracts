// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package icoreregistry

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

// ICoreRegistryEventsMetaData contains all meta data concerning the ICoreRegistryEvents contract.
var ICoreRegistryEventsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"ChainStatusChanged\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractStatusChanged\",\"inputs\":[{\"name\":\"addressBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewChainMetadata\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewContractConfiguration\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenRegistered\",\"inputs\":[{\"name\":\"originAddress\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"address_\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenUpdated\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
}

// ICoreRegistryEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use ICoreRegistryEventsMetaData.ABI instead.
var ICoreRegistryEventsABI = ICoreRegistryEventsMetaData.ABI

// ICoreRegistryEvents is an auto generated Go binding around an Ethereum contract.
type ICoreRegistryEvents struct {
	ICoreRegistryEventsCaller     // Read-only binding to the contract
	ICoreRegistryEventsTransactor // Write-only binding to the contract
	ICoreRegistryEventsFilterer   // Log filterer for contract events
}

// ICoreRegistryEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICoreRegistryEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICoreRegistryEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICoreRegistryEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICoreRegistryEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICoreRegistryEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICoreRegistryEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICoreRegistryEventsSession struct {
	Contract     *ICoreRegistryEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ICoreRegistryEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICoreRegistryEventsCallerSession struct {
	Contract *ICoreRegistryEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ICoreRegistryEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICoreRegistryEventsTransactorSession struct {
	Contract     *ICoreRegistryEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ICoreRegistryEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICoreRegistryEventsRaw struct {
	Contract *ICoreRegistryEvents // Generic contract binding to access the raw methods on
}

// ICoreRegistryEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICoreRegistryEventsCallerRaw struct {
	Contract *ICoreRegistryEventsCaller // Generic read-only contract binding to access the raw methods on
}

// ICoreRegistryEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICoreRegistryEventsTransactorRaw struct {
	Contract *ICoreRegistryEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICoreRegistryEvents creates a new instance of ICoreRegistryEvents, bound to a specific deployed contract.
func NewICoreRegistryEvents(address common.Address, backend bind.ContractBackend) (*ICoreRegistryEvents, error) {
	contract, err := bindICoreRegistryEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICoreRegistryEvents{ICoreRegistryEventsCaller: ICoreRegistryEventsCaller{contract: contract}, ICoreRegistryEventsTransactor: ICoreRegistryEventsTransactor{contract: contract}, ICoreRegistryEventsFilterer: ICoreRegistryEventsFilterer{contract: contract}}, nil
}

// NewICoreRegistryEventsCaller creates a new read-only instance of ICoreRegistryEvents, bound to a specific deployed contract.
func NewICoreRegistryEventsCaller(address common.Address, caller bind.ContractCaller) (*ICoreRegistryEventsCaller, error) {
	contract, err := bindICoreRegistryEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICoreRegistryEventsCaller{contract: contract}, nil
}

// NewICoreRegistryEventsTransactor creates a new write-only instance of ICoreRegistryEvents, bound to a specific deployed contract.
func NewICoreRegistryEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*ICoreRegistryEventsTransactor, error) {
	contract, err := bindICoreRegistryEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICoreRegistryEventsTransactor{contract: contract}, nil
}

// NewICoreRegistryEventsFilterer creates a new log filterer instance of ICoreRegistryEvents, bound to a specific deployed contract.
func NewICoreRegistryEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*ICoreRegistryEventsFilterer, error) {
	contract, err := bindICoreRegistryEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICoreRegistryEventsFilterer{contract: contract}, nil
}

// bindICoreRegistryEvents binds a generic wrapper to an already deployed contract.
func bindICoreRegistryEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICoreRegistryEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICoreRegistryEvents *ICoreRegistryEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICoreRegistryEvents.Contract.ICoreRegistryEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICoreRegistryEvents *ICoreRegistryEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICoreRegistryEvents.Contract.ICoreRegistryEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICoreRegistryEvents *ICoreRegistryEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICoreRegistryEvents.Contract.ICoreRegistryEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICoreRegistryEvents *ICoreRegistryEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICoreRegistryEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICoreRegistryEvents *ICoreRegistryEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICoreRegistryEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICoreRegistryEvents *ICoreRegistryEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICoreRegistryEvents.Contract.contract.Transact(opts, method, params...)
}

// ICoreRegistryEventsChainStatusChangedIterator is returned from FilterChainStatusChanged and is used to iterate over the raw logs and unpacked data for ChainStatusChanged events raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsChainStatusChangedIterator struct {
	Event *ICoreRegistryEventsChainStatusChanged // Event containing the contract specifics and raw log

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
func (it *ICoreRegistryEventsChainStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICoreRegistryEventsChainStatusChanged)
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
		it.Event = new(ICoreRegistryEventsChainStatusChanged)
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
func (it *ICoreRegistryEventsChainStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICoreRegistryEventsChainStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICoreRegistryEventsChainStatusChanged represents a ChainStatusChanged event raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsChainStatusChanged struct {
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainStatusChanged is a free log retrieval operation binding the contract event 0x804e2ce81861d12bcd678a38e1f76d4d6c93c35f58ef03a1e26ef0d7c48da707.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) FilterChainStatusChanged(opts *bind.FilterOpts, chainId []*big.Int) (*ICoreRegistryEventsChainStatusChangedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ICoreRegistryEvents.contract.FilterLogs(opts, "ChainStatusChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &ICoreRegistryEventsChainStatusChangedIterator{contract: _ICoreRegistryEvents.contract, event: "ChainStatusChanged", logs: logs, sub: sub}, nil
}

// WatchChainStatusChanged is a free log subscription operation binding the contract event 0x804e2ce81861d12bcd678a38e1f76d4d6c93c35f58ef03a1e26ef0d7c48da707.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) WatchChainStatusChanged(opts *bind.WatchOpts, sink chan<- *ICoreRegistryEventsChainStatusChanged, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ICoreRegistryEvents.contract.WatchLogs(opts, "ChainStatusChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICoreRegistryEventsChainStatusChanged)
				if err := _ICoreRegistryEvents.contract.UnpackLog(event, "ChainStatusChanged", log); err != nil {
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

// ParseChainStatusChanged is a log parse operation binding the contract event 0x804e2ce81861d12bcd678a38e1f76d4d6c93c35f58ef03a1e26ef0d7c48da707.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) ParseChainStatusChanged(log types.Log) (*ICoreRegistryEventsChainStatusChanged, error) {
	event := new(ICoreRegistryEventsChainStatusChanged)
	if err := _ICoreRegistryEvents.contract.UnpackLog(event, "ChainStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICoreRegistryEventsContractRegisteredIterator is returned from FilterContractRegistered and is used to iterate over the raw logs and unpacked data for ContractRegistered events raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsContractRegisteredIterator struct {
	Event *ICoreRegistryEventsContractRegistered // Event containing the contract specifics and raw log

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
func (it *ICoreRegistryEventsContractRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICoreRegistryEventsContractRegistered)
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
		it.Event = new(ICoreRegistryEventsContractRegistered)
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
func (it *ICoreRegistryEventsContractRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICoreRegistryEventsContractRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICoreRegistryEventsContractRegistered represents a ContractRegistered event raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsContractRegistered struct {
	ChainId      *big.Int
	ContractType common.Hash
	AddressBytes []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractRegistered is a free log retrieval operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) FilterContractRegistered(opts *bind.FilterOpts, chainId []*big.Int, contractType []string) (*ICoreRegistryEventsContractRegisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}

	logs, sub, err := _ICoreRegistryEvents.contract.FilterLogs(opts, "ContractRegistered", chainIdRule, contractTypeRule)
	if err != nil {
		return nil, err
	}
	return &ICoreRegistryEventsContractRegisteredIterator{contract: _ICoreRegistryEvents.contract, event: "ContractRegistered", logs: logs, sub: sub}, nil
}

// WatchContractRegistered is a free log subscription operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) WatchContractRegistered(opts *bind.WatchOpts, sink chan<- *ICoreRegistryEventsContractRegistered, chainId []*big.Int, contractType []string) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}

	logs, sub, err := _ICoreRegistryEvents.contract.WatchLogs(opts, "ContractRegistered", chainIdRule, contractTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICoreRegistryEventsContractRegistered)
				if err := _ICoreRegistryEvents.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
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

// ParseContractRegistered is a log parse operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) ParseContractRegistered(log types.Log) (*ICoreRegistryEventsContractRegistered, error) {
	event := new(ICoreRegistryEventsContractRegistered)
	if err := _ICoreRegistryEvents.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICoreRegistryEventsContractStatusChangedIterator is returned from FilterContractStatusChanged and is used to iterate over the raw logs and unpacked data for ContractStatusChanged events raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsContractStatusChangedIterator struct {
	Event *ICoreRegistryEventsContractStatusChanged // Event containing the contract specifics and raw log

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
func (it *ICoreRegistryEventsContractStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICoreRegistryEventsContractStatusChanged)
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
		it.Event = new(ICoreRegistryEventsContractStatusChanged)
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
func (it *ICoreRegistryEventsContractStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICoreRegistryEventsContractStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICoreRegistryEventsContractStatusChanged represents a ContractStatusChanged event raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsContractStatusChanged struct {
	AddressBytes []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractStatusChanged is a free log retrieval operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) FilterContractStatusChanged(opts *bind.FilterOpts) (*ICoreRegistryEventsContractStatusChangedIterator, error) {

	logs, sub, err := _ICoreRegistryEvents.contract.FilterLogs(opts, "ContractStatusChanged")
	if err != nil {
		return nil, err
	}
	return &ICoreRegistryEventsContractStatusChangedIterator{contract: _ICoreRegistryEvents.contract, event: "ContractStatusChanged", logs: logs, sub: sub}, nil
}

// WatchContractStatusChanged is a free log subscription operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) WatchContractStatusChanged(opts *bind.WatchOpts, sink chan<- *ICoreRegistryEventsContractStatusChanged) (event.Subscription, error) {

	logs, sub, err := _ICoreRegistryEvents.contract.WatchLogs(opts, "ContractStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICoreRegistryEventsContractStatusChanged)
				if err := _ICoreRegistryEvents.contract.UnpackLog(event, "ContractStatusChanged", log); err != nil {
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

// ParseContractStatusChanged is a log parse operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) ParseContractStatusChanged(log types.Log) (*ICoreRegistryEventsContractStatusChanged, error) {
	event := new(ICoreRegistryEventsContractStatusChanged)
	if err := _ICoreRegistryEvents.contract.UnpackLog(event, "ContractStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICoreRegistryEventsNewChainMetadataIterator is returned from FilterNewChainMetadata and is used to iterate over the raw logs and unpacked data for NewChainMetadata events raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsNewChainMetadataIterator struct {
	Event *ICoreRegistryEventsNewChainMetadata // Event containing the contract specifics and raw log

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
func (it *ICoreRegistryEventsNewChainMetadataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICoreRegistryEventsNewChainMetadata)
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
		it.Event = new(ICoreRegistryEventsNewChainMetadata)
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
func (it *ICoreRegistryEventsNewChainMetadataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICoreRegistryEventsNewChainMetadataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICoreRegistryEventsNewChainMetadata represents a NewChainMetadata event raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsNewChainMetadata struct {
	ChainId *big.Int
	Key     string
	Value   []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewChainMetadata is a free log retrieval operation binding the contract event 0xa274ec62326b7df2ece1add0babcb0f90b87b6d4a7e144f8d0b9b9057c039dc1.
//
// Solidity: event NewChainMetadata(uint256 indexed chainId, string key, bytes value)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) FilterNewChainMetadata(opts *bind.FilterOpts, chainId []*big.Int) (*ICoreRegistryEventsNewChainMetadataIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ICoreRegistryEvents.contract.FilterLogs(opts, "NewChainMetadata", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &ICoreRegistryEventsNewChainMetadataIterator{contract: _ICoreRegistryEvents.contract, event: "NewChainMetadata", logs: logs, sub: sub}, nil
}

// WatchNewChainMetadata is a free log subscription operation binding the contract event 0xa274ec62326b7df2ece1add0babcb0f90b87b6d4a7e144f8d0b9b9057c039dc1.
//
// Solidity: event NewChainMetadata(uint256 indexed chainId, string key, bytes value)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) WatchNewChainMetadata(opts *bind.WatchOpts, sink chan<- *ICoreRegistryEventsNewChainMetadata, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ICoreRegistryEvents.contract.WatchLogs(opts, "NewChainMetadata", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICoreRegistryEventsNewChainMetadata)
				if err := _ICoreRegistryEvents.contract.UnpackLog(event, "NewChainMetadata", log); err != nil {
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

// ParseNewChainMetadata is a log parse operation binding the contract event 0xa274ec62326b7df2ece1add0babcb0f90b87b6d4a7e144f8d0b9b9057c039dc1.
//
// Solidity: event NewChainMetadata(uint256 indexed chainId, string key, bytes value)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) ParseNewChainMetadata(log types.Log) (*ICoreRegistryEventsNewChainMetadata, error) {
	event := new(ICoreRegistryEventsNewChainMetadata)
	if err := _ICoreRegistryEvents.contract.UnpackLog(event, "NewChainMetadata", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICoreRegistryEventsNewContractConfigurationIterator is returned from FilterNewContractConfiguration and is used to iterate over the raw logs and unpacked data for NewContractConfiguration events raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsNewContractConfigurationIterator struct {
	Event *ICoreRegistryEventsNewContractConfiguration // Event containing the contract specifics and raw log

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
func (it *ICoreRegistryEventsNewContractConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICoreRegistryEventsNewContractConfiguration)
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
		it.Event = new(ICoreRegistryEventsNewContractConfiguration)
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
func (it *ICoreRegistryEventsNewContractConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICoreRegistryEventsNewContractConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICoreRegistryEventsNewContractConfiguration represents a NewContractConfiguration event raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsNewContractConfiguration struct {
	ChainId      *big.Int
	ContractType string
	Key          string
	Value        []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewContractConfiguration is a free log retrieval operation binding the contract event 0x9453d408be9167c029f60b137dddf03dd0d2ff9f51d834c1b2238f335bd96f9b.
//
// Solidity: event NewContractConfiguration(uint256 indexed chainId, string contractType, string key, bytes value)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) FilterNewContractConfiguration(opts *bind.FilterOpts, chainId []*big.Int) (*ICoreRegistryEventsNewContractConfigurationIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ICoreRegistryEvents.contract.FilterLogs(opts, "NewContractConfiguration", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &ICoreRegistryEventsNewContractConfigurationIterator{contract: _ICoreRegistryEvents.contract, event: "NewContractConfiguration", logs: logs, sub: sub}, nil
}

// WatchNewContractConfiguration is a free log subscription operation binding the contract event 0x9453d408be9167c029f60b137dddf03dd0d2ff9f51d834c1b2238f335bd96f9b.
//
// Solidity: event NewContractConfiguration(uint256 indexed chainId, string contractType, string key, bytes value)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) WatchNewContractConfiguration(opts *bind.WatchOpts, sink chan<- *ICoreRegistryEventsNewContractConfiguration, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ICoreRegistryEvents.contract.WatchLogs(opts, "NewContractConfiguration", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICoreRegistryEventsNewContractConfiguration)
				if err := _ICoreRegistryEvents.contract.UnpackLog(event, "NewContractConfiguration", log); err != nil {
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

// ParseNewContractConfiguration is a log parse operation binding the contract event 0x9453d408be9167c029f60b137dddf03dd0d2ff9f51d834c1b2238f335bd96f9b.
//
// Solidity: event NewContractConfiguration(uint256 indexed chainId, string contractType, string key, bytes value)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) ParseNewContractConfiguration(log types.Log) (*ICoreRegistryEventsNewContractConfiguration, error) {
	event := new(ICoreRegistryEventsNewContractConfiguration)
	if err := _ICoreRegistryEvents.contract.UnpackLog(event, "NewContractConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICoreRegistryEventsZRC20TokenRegisteredIterator is returned from FilterZRC20TokenRegistered and is used to iterate over the raw logs and unpacked data for ZRC20TokenRegistered events raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsZRC20TokenRegisteredIterator struct {
	Event *ICoreRegistryEventsZRC20TokenRegistered // Event containing the contract specifics and raw log

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
func (it *ICoreRegistryEventsZRC20TokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICoreRegistryEventsZRC20TokenRegistered)
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
		it.Event = new(ICoreRegistryEventsZRC20TokenRegistered)
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
func (it *ICoreRegistryEventsZRC20TokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICoreRegistryEventsZRC20TokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICoreRegistryEventsZRC20TokenRegistered represents a ZRC20TokenRegistered event raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsZRC20TokenRegistered struct {
	OriginAddress common.Hash
	Address       common.Address
	Decimals      uint8
	OriginChainId *big.Int
	Symbol        string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterZRC20TokenRegistered is a free log retrieval operation binding the contract event 0xa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367.
//
// Solidity: event ZRC20TokenRegistered(bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) FilterZRC20TokenRegistered(opts *bind.FilterOpts, originAddress [][]byte, address_ []common.Address) (*ICoreRegistryEventsZRC20TokenRegisteredIterator, error) {

	var originAddressRule []interface{}
	for _, originAddressItem := range originAddress {
		originAddressRule = append(originAddressRule, originAddressItem)
	}
	var address_Rule []interface{}
	for _, address_Item := range address_ {
		address_Rule = append(address_Rule, address_Item)
	}

	logs, sub, err := _ICoreRegistryEvents.contract.FilterLogs(opts, "ZRC20TokenRegistered", originAddressRule, address_Rule)
	if err != nil {
		return nil, err
	}
	return &ICoreRegistryEventsZRC20TokenRegisteredIterator{contract: _ICoreRegistryEvents.contract, event: "ZRC20TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchZRC20TokenRegistered is a free log subscription operation binding the contract event 0xa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367.
//
// Solidity: event ZRC20TokenRegistered(bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) WatchZRC20TokenRegistered(opts *bind.WatchOpts, sink chan<- *ICoreRegistryEventsZRC20TokenRegistered, originAddress [][]byte, address_ []common.Address) (event.Subscription, error) {

	var originAddressRule []interface{}
	for _, originAddressItem := range originAddress {
		originAddressRule = append(originAddressRule, originAddressItem)
	}
	var address_Rule []interface{}
	for _, address_Item := range address_ {
		address_Rule = append(address_Rule, address_Item)
	}

	logs, sub, err := _ICoreRegistryEvents.contract.WatchLogs(opts, "ZRC20TokenRegistered", originAddressRule, address_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICoreRegistryEventsZRC20TokenRegistered)
				if err := _ICoreRegistryEvents.contract.UnpackLog(event, "ZRC20TokenRegistered", log); err != nil {
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

// ParseZRC20TokenRegistered is a log parse operation binding the contract event 0xa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367.
//
// Solidity: event ZRC20TokenRegistered(bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) ParseZRC20TokenRegistered(log types.Log) (*ICoreRegistryEventsZRC20TokenRegistered, error) {
	event := new(ICoreRegistryEventsZRC20TokenRegistered)
	if err := _ICoreRegistryEvents.contract.UnpackLog(event, "ZRC20TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICoreRegistryEventsZRC20TokenUpdatedIterator is returned from FilterZRC20TokenUpdated and is used to iterate over the raw logs and unpacked data for ZRC20TokenUpdated events raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsZRC20TokenUpdatedIterator struct {
	Event *ICoreRegistryEventsZRC20TokenUpdated // Event containing the contract specifics and raw log

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
func (it *ICoreRegistryEventsZRC20TokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICoreRegistryEventsZRC20TokenUpdated)
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
		it.Event = new(ICoreRegistryEventsZRC20TokenUpdated)
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
func (it *ICoreRegistryEventsZRC20TokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICoreRegistryEventsZRC20TokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICoreRegistryEventsZRC20TokenUpdated represents a ZRC20TokenUpdated event raised by the ICoreRegistryEvents contract.
type ICoreRegistryEventsZRC20TokenUpdated struct {
	Address common.Address
	Active  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterZRC20TokenUpdated is a free log retrieval operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) FilterZRC20TokenUpdated(opts *bind.FilterOpts) (*ICoreRegistryEventsZRC20TokenUpdatedIterator, error) {

	logs, sub, err := _ICoreRegistryEvents.contract.FilterLogs(opts, "ZRC20TokenUpdated")
	if err != nil {
		return nil, err
	}
	return &ICoreRegistryEventsZRC20TokenUpdatedIterator{contract: _ICoreRegistryEvents.contract, event: "ZRC20TokenUpdated", logs: logs, sub: sub}, nil
}

// WatchZRC20TokenUpdated is a free log subscription operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) WatchZRC20TokenUpdated(opts *bind.WatchOpts, sink chan<- *ICoreRegistryEventsZRC20TokenUpdated) (event.Subscription, error) {

	logs, sub, err := _ICoreRegistryEvents.contract.WatchLogs(opts, "ZRC20TokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICoreRegistryEventsZRC20TokenUpdated)
				if err := _ICoreRegistryEvents.contract.UnpackLog(event, "ZRC20TokenUpdated", log); err != nil {
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

// ParseZRC20TokenUpdated is a log parse operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_ICoreRegistryEvents *ICoreRegistryEventsFilterer) ParseZRC20TokenUpdated(log types.Log) (*ICoreRegistryEventsZRC20TokenUpdated, error) {
	event := new(ICoreRegistryEventsZRC20TokenUpdated)
	if err := _ICoreRegistryEvents.contract.UnpackLog(event, "ZRC20TokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
