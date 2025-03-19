// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izetaregistry

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

// IZetaRegistryEventsMetaData contains all meta data concerning the IZetaRegistryEvents contract.
var IZetaRegistryEventsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainStatusUpdated\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"active\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractAddressUpdated\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"category\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractConfigurationUpdated\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"configurationData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractStatusUpdated\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"active\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenAdded\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"tokenIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// IZetaRegistryEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IZetaRegistryEventsMetaData.ABI instead.
var IZetaRegistryEventsABI = IZetaRegistryEventsMetaData.ABI

// IZetaRegistryEvents is an auto generated Go binding around an Ethereum contract.
type IZetaRegistryEvents struct {
	IZetaRegistryEventsCaller     // Read-only binding to the contract
	IZetaRegistryEventsTransactor // Write-only binding to the contract
	IZetaRegistryEventsFilterer   // Log filterer for contract events
}

// IZetaRegistryEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZetaRegistryEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaRegistryEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZetaRegistryEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaRegistryEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZetaRegistryEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaRegistryEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZetaRegistryEventsSession struct {
	Contract     *IZetaRegistryEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IZetaRegistryEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZetaRegistryEventsCallerSession struct {
	Contract *IZetaRegistryEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IZetaRegistryEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZetaRegistryEventsTransactorSession struct {
	Contract     *IZetaRegistryEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IZetaRegistryEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZetaRegistryEventsRaw struct {
	Contract *IZetaRegistryEvents // Generic contract binding to access the raw methods on
}

// IZetaRegistryEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZetaRegistryEventsCallerRaw struct {
	Contract *IZetaRegistryEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IZetaRegistryEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZetaRegistryEventsTransactorRaw struct {
	Contract *IZetaRegistryEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZetaRegistryEvents creates a new instance of IZetaRegistryEvents, bound to a specific deployed contract.
func NewIZetaRegistryEvents(address common.Address, backend bind.ContractBackend) (*IZetaRegistryEvents, error) {
	contract, err := bindIZetaRegistryEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryEvents{IZetaRegistryEventsCaller: IZetaRegistryEventsCaller{contract: contract}, IZetaRegistryEventsTransactor: IZetaRegistryEventsTransactor{contract: contract}, IZetaRegistryEventsFilterer: IZetaRegistryEventsFilterer{contract: contract}}, nil
}

// NewIZetaRegistryEventsCaller creates a new read-only instance of IZetaRegistryEvents, bound to a specific deployed contract.
func NewIZetaRegistryEventsCaller(address common.Address, caller bind.ContractCaller) (*IZetaRegistryEventsCaller, error) {
	contract, err := bindIZetaRegistryEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryEventsCaller{contract: contract}, nil
}

// NewIZetaRegistryEventsTransactor creates a new write-only instance of IZetaRegistryEvents, bound to a specific deployed contract.
func NewIZetaRegistryEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IZetaRegistryEventsTransactor, error) {
	contract, err := bindIZetaRegistryEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryEventsTransactor{contract: contract}, nil
}

// NewIZetaRegistryEventsFilterer creates a new log filterer instance of IZetaRegistryEvents, bound to a specific deployed contract.
func NewIZetaRegistryEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IZetaRegistryEventsFilterer, error) {
	contract, err := bindIZetaRegistryEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryEventsFilterer{contract: contract}, nil
}

// bindIZetaRegistryEvents binds a generic wrapper to an already deployed contract.
func bindIZetaRegistryEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZetaRegistryEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaRegistryEvents *IZetaRegistryEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaRegistryEvents.Contract.IZetaRegistryEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaRegistryEvents *IZetaRegistryEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaRegistryEvents.Contract.IZetaRegistryEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaRegistryEvents *IZetaRegistryEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaRegistryEvents.Contract.IZetaRegistryEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaRegistryEvents *IZetaRegistryEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaRegistryEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaRegistryEvents *IZetaRegistryEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaRegistryEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaRegistryEvents *IZetaRegistryEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaRegistryEvents.Contract.contract.Transact(opts, method, params...)
}

// IZetaRegistryEventsChainAddedIterator is returned from FilterChainAdded and is used to iterate over the raw logs and unpacked data for ChainAdded events raised by the IZetaRegistryEvents contract.
type IZetaRegistryEventsChainAddedIterator struct {
	Event *IZetaRegistryEventsChainAdded // Event containing the contract specifics and raw log

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
func (it *IZetaRegistryEventsChainAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaRegistryEventsChainAdded)
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
		it.Event = new(IZetaRegistryEventsChainAdded)
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
func (it *IZetaRegistryEventsChainAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaRegistryEventsChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaRegistryEventsChainAdded represents a ChainAdded event raised by the IZetaRegistryEvents contract.
type IZetaRegistryEventsChainAdded struct {
	ChainIdentifier [32]byte
	ChainId         *big.Int
	Name            string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChainAdded is a free log retrieval operation binding the contract event 0x5a8e7b221f7affd9f2b1becd6ad936d4c721d1dd0f8666d7ddf1d624664edd1a.
//
// Solidity: event ChainAdded(bytes32 indexed chainIdentifier, uint256 chainId, string name)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) FilterChainAdded(opts *bind.FilterOpts, chainIdentifier [][32]byte) (*IZetaRegistryEventsChainAddedIterator, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}

	logs, sub, err := _IZetaRegistryEvents.contract.FilterLogs(opts, "ChainAdded", chainIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryEventsChainAddedIterator{contract: _IZetaRegistryEvents.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

// WatchChainAdded is a free log subscription operation binding the contract event 0x5a8e7b221f7affd9f2b1becd6ad936d4c721d1dd0f8666d7ddf1d624664edd1a.
//
// Solidity: event ChainAdded(bytes32 indexed chainIdentifier, uint256 chainId, string name)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *IZetaRegistryEventsChainAdded, chainIdentifier [][32]byte) (event.Subscription, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}

	logs, sub, err := _IZetaRegistryEvents.contract.WatchLogs(opts, "ChainAdded", chainIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaRegistryEventsChainAdded)
				if err := _IZetaRegistryEvents.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

// ParseChainAdded is a log parse operation binding the contract event 0x5a8e7b221f7affd9f2b1becd6ad936d4c721d1dd0f8666d7ddf1d624664edd1a.
//
// Solidity: event ChainAdded(bytes32 indexed chainIdentifier, uint256 chainId, string name)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) ParseChainAdded(log types.Log) (*IZetaRegistryEventsChainAdded, error) {
	event := new(IZetaRegistryEventsChainAdded)
	if err := _IZetaRegistryEvents.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaRegistryEventsChainStatusUpdatedIterator is returned from FilterChainStatusUpdated and is used to iterate over the raw logs and unpacked data for ChainStatusUpdated events raised by the IZetaRegistryEvents contract.
type IZetaRegistryEventsChainStatusUpdatedIterator struct {
	Event *IZetaRegistryEventsChainStatusUpdated // Event containing the contract specifics and raw log

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
func (it *IZetaRegistryEventsChainStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaRegistryEventsChainStatusUpdated)
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
		it.Event = new(IZetaRegistryEventsChainStatusUpdated)
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
func (it *IZetaRegistryEventsChainStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaRegistryEventsChainStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaRegistryEventsChainStatusUpdated represents a ChainStatusUpdated event raised by the IZetaRegistryEvents contract.
type IZetaRegistryEventsChainStatusUpdated struct {
	ChainIdentifier [32]byte
	Active          bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChainStatusUpdated is a free log retrieval operation binding the contract event 0xae28df7592823455d07510da74e0d26fa77925e92a66817e8c842056c490d4f8.
//
// Solidity: event ChainStatusUpdated(bytes32 indexed chainIdentifier, bool active)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) FilterChainStatusUpdated(opts *bind.FilterOpts, chainIdentifier [][32]byte) (*IZetaRegistryEventsChainStatusUpdatedIterator, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}

	logs, sub, err := _IZetaRegistryEvents.contract.FilterLogs(opts, "ChainStatusUpdated", chainIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryEventsChainStatusUpdatedIterator{contract: _IZetaRegistryEvents.contract, event: "ChainStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchChainStatusUpdated is a free log subscription operation binding the contract event 0xae28df7592823455d07510da74e0d26fa77925e92a66817e8c842056c490d4f8.
//
// Solidity: event ChainStatusUpdated(bytes32 indexed chainIdentifier, bool active)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) WatchChainStatusUpdated(opts *bind.WatchOpts, sink chan<- *IZetaRegistryEventsChainStatusUpdated, chainIdentifier [][32]byte) (event.Subscription, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}

	logs, sub, err := _IZetaRegistryEvents.contract.WatchLogs(opts, "ChainStatusUpdated", chainIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaRegistryEventsChainStatusUpdated)
				if err := _IZetaRegistryEvents.contract.UnpackLog(event, "ChainStatusUpdated", log); err != nil {
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

// ParseChainStatusUpdated is a log parse operation binding the contract event 0xae28df7592823455d07510da74e0d26fa77925e92a66817e8c842056c490d4f8.
//
// Solidity: event ChainStatusUpdated(bytes32 indexed chainIdentifier, bool active)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) ParseChainStatusUpdated(log types.Log) (*IZetaRegistryEventsChainStatusUpdated, error) {
	event := new(IZetaRegistryEventsChainStatusUpdated)
	if err := _IZetaRegistryEvents.contract.UnpackLog(event, "ChainStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaRegistryEventsContractAddressUpdatedIterator is returned from FilterContractAddressUpdated and is used to iterate over the raw logs and unpacked data for ContractAddressUpdated events raised by the IZetaRegistryEvents contract.
type IZetaRegistryEventsContractAddressUpdatedIterator struct {
	Event *IZetaRegistryEventsContractAddressUpdated // Event containing the contract specifics and raw log

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
func (it *IZetaRegistryEventsContractAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaRegistryEventsContractAddressUpdated)
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
		it.Event = new(IZetaRegistryEventsContractAddressUpdated)
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
func (it *IZetaRegistryEventsContractAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaRegistryEventsContractAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaRegistryEventsContractAddressUpdated represents a ContractAddressUpdated event raised by the IZetaRegistryEvents contract.
type IZetaRegistryEventsContractAddressUpdated struct {
	ChainIdentifier    [32]byte
	ContractIdentifier [32]byte
	Category           uint8
	ContractAddress    common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterContractAddressUpdated is a free log retrieval operation binding the contract event 0xe8bec620648cefe1f492bdb89a04125718c351bc0c539443bbfddba38e60a652.
//
// Solidity: event ContractAddressUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, uint8 indexed category, address contractAddress)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) FilterContractAddressUpdated(opts *bind.FilterOpts, chainIdentifier [][32]byte, contractIdentifier [][32]byte, category []uint8) (*IZetaRegistryEventsContractAddressUpdatedIterator, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}
	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _IZetaRegistryEvents.contract.FilterLogs(opts, "ContractAddressUpdated", chainIdentifierRule, contractIdentifierRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryEventsContractAddressUpdatedIterator{contract: _IZetaRegistryEvents.contract, event: "ContractAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchContractAddressUpdated is a free log subscription operation binding the contract event 0xe8bec620648cefe1f492bdb89a04125718c351bc0c539443bbfddba38e60a652.
//
// Solidity: event ContractAddressUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, uint8 indexed category, address contractAddress)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) WatchContractAddressUpdated(opts *bind.WatchOpts, sink chan<- *IZetaRegistryEventsContractAddressUpdated, chainIdentifier [][32]byte, contractIdentifier [][32]byte, category []uint8) (event.Subscription, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}
	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _IZetaRegistryEvents.contract.WatchLogs(opts, "ContractAddressUpdated", chainIdentifierRule, contractIdentifierRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaRegistryEventsContractAddressUpdated)
				if err := _IZetaRegistryEvents.contract.UnpackLog(event, "ContractAddressUpdated", log); err != nil {
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

// ParseContractAddressUpdated is a log parse operation binding the contract event 0xe8bec620648cefe1f492bdb89a04125718c351bc0c539443bbfddba38e60a652.
//
// Solidity: event ContractAddressUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, uint8 indexed category, address contractAddress)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) ParseContractAddressUpdated(log types.Log) (*IZetaRegistryEventsContractAddressUpdated, error) {
	event := new(IZetaRegistryEventsContractAddressUpdated)
	if err := _IZetaRegistryEvents.contract.UnpackLog(event, "ContractAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaRegistryEventsContractConfigurationUpdatedIterator is returned from FilterContractConfigurationUpdated and is used to iterate over the raw logs and unpacked data for ContractConfigurationUpdated events raised by the IZetaRegistryEvents contract.
type IZetaRegistryEventsContractConfigurationUpdatedIterator struct {
	Event *IZetaRegistryEventsContractConfigurationUpdated // Event containing the contract specifics and raw log

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
func (it *IZetaRegistryEventsContractConfigurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaRegistryEventsContractConfigurationUpdated)
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
		it.Event = new(IZetaRegistryEventsContractConfigurationUpdated)
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
func (it *IZetaRegistryEventsContractConfigurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaRegistryEventsContractConfigurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaRegistryEventsContractConfigurationUpdated represents a ContractConfigurationUpdated event raised by the IZetaRegistryEvents contract.
type IZetaRegistryEventsContractConfigurationUpdated struct {
	ChainIdentifier    [32]byte
	ContractIdentifier [32]byte
	ConfigurationData  []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterContractConfigurationUpdated is a free log retrieval operation binding the contract event 0x11af48d095e78bf8811830bb22710e171451b6a34ddbe65490b4b2e31b34c310.
//
// Solidity: event ContractConfigurationUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bytes configurationData)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) FilterContractConfigurationUpdated(opts *bind.FilterOpts, chainIdentifier [][32]byte, contractIdentifier [][32]byte) (*IZetaRegistryEventsContractConfigurationUpdatedIterator, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}

	logs, sub, err := _IZetaRegistryEvents.contract.FilterLogs(opts, "ContractConfigurationUpdated", chainIdentifierRule, contractIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryEventsContractConfigurationUpdatedIterator{contract: _IZetaRegistryEvents.contract, event: "ContractConfigurationUpdated", logs: logs, sub: sub}, nil
}

// WatchContractConfigurationUpdated is a free log subscription operation binding the contract event 0x11af48d095e78bf8811830bb22710e171451b6a34ddbe65490b4b2e31b34c310.
//
// Solidity: event ContractConfigurationUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bytes configurationData)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) WatchContractConfigurationUpdated(opts *bind.WatchOpts, sink chan<- *IZetaRegistryEventsContractConfigurationUpdated, chainIdentifier [][32]byte, contractIdentifier [][32]byte) (event.Subscription, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}

	logs, sub, err := _IZetaRegistryEvents.contract.WatchLogs(opts, "ContractConfigurationUpdated", chainIdentifierRule, contractIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaRegistryEventsContractConfigurationUpdated)
				if err := _IZetaRegistryEvents.contract.UnpackLog(event, "ContractConfigurationUpdated", log); err != nil {
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

// ParseContractConfigurationUpdated is a log parse operation binding the contract event 0x11af48d095e78bf8811830bb22710e171451b6a34ddbe65490b4b2e31b34c310.
//
// Solidity: event ContractConfigurationUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bytes configurationData)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) ParseContractConfigurationUpdated(log types.Log) (*IZetaRegistryEventsContractConfigurationUpdated, error) {
	event := new(IZetaRegistryEventsContractConfigurationUpdated)
	if err := _IZetaRegistryEvents.contract.UnpackLog(event, "ContractConfigurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaRegistryEventsContractStatusUpdatedIterator is returned from FilterContractStatusUpdated and is used to iterate over the raw logs and unpacked data for ContractStatusUpdated events raised by the IZetaRegistryEvents contract.
type IZetaRegistryEventsContractStatusUpdatedIterator struct {
	Event *IZetaRegistryEventsContractStatusUpdated // Event containing the contract specifics and raw log

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
func (it *IZetaRegistryEventsContractStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaRegistryEventsContractStatusUpdated)
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
		it.Event = new(IZetaRegistryEventsContractStatusUpdated)
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
func (it *IZetaRegistryEventsContractStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaRegistryEventsContractStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaRegistryEventsContractStatusUpdated represents a ContractStatusUpdated event raised by the IZetaRegistryEvents contract.
type IZetaRegistryEventsContractStatusUpdated struct {
	ChainIdentifier    [32]byte
	ContractIdentifier [32]byte
	Active             bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterContractStatusUpdated is a free log retrieval operation binding the contract event 0x166bdf9d7b5f4849bcabfef5465f6edae79fe8f333c2019f6f5cf489084cf581.
//
// Solidity: event ContractStatusUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bool active)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) FilterContractStatusUpdated(opts *bind.FilterOpts, chainIdentifier [][32]byte, contractIdentifier [][32]byte) (*IZetaRegistryEventsContractStatusUpdatedIterator, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}

	logs, sub, err := _IZetaRegistryEvents.contract.FilterLogs(opts, "ContractStatusUpdated", chainIdentifierRule, contractIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryEventsContractStatusUpdatedIterator{contract: _IZetaRegistryEvents.contract, event: "ContractStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchContractStatusUpdated is a free log subscription operation binding the contract event 0x166bdf9d7b5f4849bcabfef5465f6edae79fe8f333c2019f6f5cf489084cf581.
//
// Solidity: event ContractStatusUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bool active)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) WatchContractStatusUpdated(opts *bind.WatchOpts, sink chan<- *IZetaRegistryEventsContractStatusUpdated, chainIdentifier [][32]byte, contractIdentifier [][32]byte) (event.Subscription, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}

	logs, sub, err := _IZetaRegistryEvents.contract.WatchLogs(opts, "ContractStatusUpdated", chainIdentifierRule, contractIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaRegistryEventsContractStatusUpdated)
				if err := _IZetaRegistryEvents.contract.UnpackLog(event, "ContractStatusUpdated", log); err != nil {
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

// ParseContractStatusUpdated is a log parse operation binding the contract event 0x166bdf9d7b5f4849bcabfef5465f6edae79fe8f333c2019f6f5cf489084cf581.
//
// Solidity: event ContractStatusUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bool active)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) ParseContractStatusUpdated(log types.Log) (*IZetaRegistryEventsContractStatusUpdated, error) {
	event := new(IZetaRegistryEventsContractStatusUpdated)
	if err := _IZetaRegistryEvents.contract.UnpackLog(event, "ContractStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaRegistryEventsZRC20TokenAddedIterator is returned from FilterZRC20TokenAdded and is used to iterate over the raw logs and unpacked data for ZRC20TokenAdded events raised by the IZetaRegistryEvents contract.
type IZetaRegistryEventsZRC20TokenAddedIterator struct {
	Event *IZetaRegistryEventsZRC20TokenAdded // Event containing the contract specifics and raw log

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
func (it *IZetaRegistryEventsZRC20TokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaRegistryEventsZRC20TokenAdded)
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
		it.Event = new(IZetaRegistryEventsZRC20TokenAdded)
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
func (it *IZetaRegistryEventsZRC20TokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaRegistryEventsZRC20TokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaRegistryEventsZRC20TokenAdded represents a ZRC20TokenAdded event raised by the IZetaRegistryEvents contract.
type IZetaRegistryEventsZRC20TokenAdded struct {
	ChainIdentifier [32]byte
	TokenIdentifier [32]byte
	TokenAddress    common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterZRC20TokenAdded is a free log retrieval operation binding the contract event 0xbc0f3f1af072bc04772392bb2fcb5a0f1a9ec90139271c5599034d3b8e5888b3.
//
// Solidity: event ZRC20TokenAdded(bytes32 indexed chainIdentifier, bytes32 indexed tokenIdentifier, address tokenAddress)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) FilterZRC20TokenAdded(opts *bind.FilterOpts, chainIdentifier [][32]byte, tokenIdentifier [][32]byte) (*IZetaRegistryEventsZRC20TokenAddedIterator, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var tokenIdentifierRule []interface{}
	for _, tokenIdentifierItem := range tokenIdentifier {
		tokenIdentifierRule = append(tokenIdentifierRule, tokenIdentifierItem)
	}

	logs, sub, err := _IZetaRegistryEvents.contract.FilterLogs(opts, "ZRC20TokenAdded", chainIdentifierRule, tokenIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryEventsZRC20TokenAddedIterator{contract: _IZetaRegistryEvents.contract, event: "ZRC20TokenAdded", logs: logs, sub: sub}, nil
}

// WatchZRC20TokenAdded is a free log subscription operation binding the contract event 0xbc0f3f1af072bc04772392bb2fcb5a0f1a9ec90139271c5599034d3b8e5888b3.
//
// Solidity: event ZRC20TokenAdded(bytes32 indexed chainIdentifier, bytes32 indexed tokenIdentifier, address tokenAddress)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) WatchZRC20TokenAdded(opts *bind.WatchOpts, sink chan<- *IZetaRegistryEventsZRC20TokenAdded, chainIdentifier [][32]byte, tokenIdentifier [][32]byte) (event.Subscription, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var tokenIdentifierRule []interface{}
	for _, tokenIdentifierItem := range tokenIdentifier {
		tokenIdentifierRule = append(tokenIdentifierRule, tokenIdentifierItem)
	}

	logs, sub, err := _IZetaRegistryEvents.contract.WatchLogs(opts, "ZRC20TokenAdded", chainIdentifierRule, tokenIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaRegistryEventsZRC20TokenAdded)
				if err := _IZetaRegistryEvents.contract.UnpackLog(event, "ZRC20TokenAdded", log); err != nil {
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

// ParseZRC20TokenAdded is a log parse operation binding the contract event 0xbc0f3f1af072bc04772392bb2fcb5a0f1a9ec90139271c5599034d3b8e5888b3.
//
// Solidity: event ZRC20TokenAdded(bytes32 indexed chainIdentifier, bytes32 indexed tokenIdentifier, address tokenAddress)
func (_IZetaRegistryEvents *IZetaRegistryEventsFilterer) ParseZRC20TokenAdded(log types.Log) (*IZetaRegistryEventsZRC20TokenAdded, error) {
	event := new(IZetaRegistryEventsZRC20TokenAdded)
	if err := _IZetaRegistryEvents.contract.UnpackLog(event, "ZRC20TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
