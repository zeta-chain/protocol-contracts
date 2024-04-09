// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectorzevm

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
)

// ZetaInterfacesMetaData contains all meta data concerning the ZetaInterfaces contract.
var ZetaInterfacesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zetaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"ZetaReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zetaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingZetaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"ZetaReverted\",\"type\":\"event\"}]",
}

// ZetaInterfacesABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaInterfacesMetaData.ABI instead.
var ZetaInterfacesABI = ZetaInterfacesMetaData.ABI

// ZetaInterfaces is an auto generated Go binding around an Ethereum contract.
type ZetaInterfaces struct {
	ZetaInterfacesCaller     // Read-only binding to the contract
	ZetaInterfacesTransactor // Write-only binding to the contract
	ZetaInterfacesFilterer   // Log filterer for contract events
}

// ZetaInterfacesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaInterfacesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInterfacesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaInterfacesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInterfacesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaInterfacesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInterfacesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaInterfacesSession struct {
	Contract     *ZetaInterfaces   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaInterfacesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaInterfacesCallerSession struct {
	Contract *ZetaInterfacesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ZetaInterfacesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaInterfacesTransactorSession struct {
	Contract     *ZetaInterfacesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ZetaInterfacesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaInterfacesRaw struct {
	Contract *ZetaInterfaces // Generic contract binding to access the raw methods on
}

// ZetaInterfacesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaInterfacesCallerRaw struct {
	Contract *ZetaInterfacesCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaInterfacesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaInterfacesTransactorRaw struct {
	Contract *ZetaInterfacesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaInterfaces creates a new instance of ZetaInterfaces, bound to a specific deployed contract.
func NewZetaInterfaces(address common.Address, backend bind.ContractBackend) (*ZetaInterfaces, error) {
	contract, err := bindZetaInterfaces(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaInterfaces{ZetaInterfacesCaller: ZetaInterfacesCaller{contract: contract}, ZetaInterfacesTransactor: ZetaInterfacesTransactor{contract: contract}, ZetaInterfacesFilterer: ZetaInterfacesFilterer{contract: contract}}, nil
}

// NewZetaInterfacesCaller creates a new read-only instance of ZetaInterfaces, bound to a specific deployed contract.
func NewZetaInterfacesCaller(address common.Address, caller bind.ContractCaller) (*ZetaInterfacesCaller, error) {
	contract, err := bindZetaInterfaces(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaInterfacesCaller{contract: contract}, nil
}

// NewZetaInterfacesTransactor creates a new write-only instance of ZetaInterfaces, bound to a specific deployed contract.
func NewZetaInterfacesTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaInterfacesTransactor, error) {
	contract, err := bindZetaInterfaces(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaInterfacesTransactor{contract: contract}, nil
}

// NewZetaInterfacesFilterer creates a new log filterer instance of ZetaInterfaces, bound to a specific deployed contract.
func NewZetaInterfacesFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaInterfacesFilterer, error) {
	contract, err := bindZetaInterfaces(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaInterfacesFilterer{contract: contract}, nil
}

// bindZetaInterfaces binds a generic wrapper to an already deployed contract.
func bindZetaInterfaces(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZetaInterfacesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaInterfaces *ZetaInterfacesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaInterfaces.Contract.ZetaInterfacesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaInterfaces *ZetaInterfacesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInterfaces.Contract.ZetaInterfacesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaInterfaces *ZetaInterfacesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaInterfaces.Contract.ZetaInterfacesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaInterfaces *ZetaInterfacesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaInterfaces.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaInterfaces *ZetaInterfacesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInterfaces.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaInterfaces *ZetaInterfacesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaInterfaces.Contract.contract.Transact(opts, method, params...)
}

// ZetaInterfacesZetaReceivedIterator is returned from FilterZetaReceived and is used to iterate over the raw logs and unpacked data for ZetaReceived events raised by the ZetaInterfaces contract.
type ZetaInterfacesZetaReceivedIterator struct {
	Event *ZetaInterfacesZetaReceived // Event containing the contract specifics and raw log

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
func (it *ZetaInterfacesZetaReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaInterfacesZetaReceived)
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
		it.Event = new(ZetaInterfacesZetaReceived)
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
func (it *ZetaInterfacesZetaReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaInterfacesZetaReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaInterfacesZetaReceived represents a ZetaReceived event raised by the ZetaInterfaces contract.
type ZetaInterfacesZetaReceived struct {
	ZetaTxSenderAddress []byte
	SourceChainId       *big.Int
	DestinationAddress  common.Address
	ZetaValue           *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterZetaReceived is a free log retrieval operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaInterfaces *ZetaInterfacesFilterer) FilterZetaReceived(opts *bind.FilterOpts, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (*ZetaInterfacesZetaReceivedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaInterfaces.contract.FilterLogs(opts, "ZetaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaInterfacesZetaReceivedIterator{contract: _ZetaInterfaces.contract, event: "ZetaReceived", logs: logs, sub: sub}, nil
}

// WatchZetaReceived is a free log subscription operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaInterfaces *ZetaInterfacesFilterer) WatchZetaReceived(opts *bind.WatchOpts, sink chan<- *ZetaInterfacesZetaReceived, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaInterfaces.contract.WatchLogs(opts, "ZetaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaInterfacesZetaReceived)
				if err := _ZetaInterfaces.contract.UnpackLog(event, "ZetaReceived", log); err != nil {
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

// ParseZetaReceived is a log parse operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaInterfaces *ZetaInterfacesFilterer) ParseZetaReceived(log types.Log) (*ZetaInterfacesZetaReceived, error) {
	event := new(ZetaInterfacesZetaReceived)
	if err := _ZetaInterfaces.contract.UnpackLog(event, "ZetaReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaInterfacesZetaRevertedIterator is returned from FilterZetaReverted and is used to iterate over the raw logs and unpacked data for ZetaReverted events raised by the ZetaInterfaces contract.
type ZetaInterfacesZetaRevertedIterator struct {
	Event *ZetaInterfacesZetaReverted // Event containing the contract specifics and raw log

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
func (it *ZetaInterfacesZetaRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaInterfacesZetaReverted)
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
		it.Event = new(ZetaInterfacesZetaReverted)
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
func (it *ZetaInterfacesZetaRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaInterfacesZetaRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaInterfacesZetaReverted represents a ZetaReverted event raised by the ZetaInterfaces contract.
type ZetaInterfacesZetaReverted struct {
	ZetaTxSenderAddress common.Address
	SourceChainId       *big.Int
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	RemainingZetaValue  *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterZetaReverted is a free log retrieval operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaInterfaces *ZetaInterfacesFilterer) FilterZetaReverted(opts *bind.FilterOpts, destinationChainId []*big.Int, internalSendHash [][32]byte) (*ZetaInterfacesZetaRevertedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaInterfaces.contract.FilterLogs(opts, "ZetaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaInterfacesZetaRevertedIterator{contract: _ZetaInterfaces.contract, event: "ZetaReverted", logs: logs, sub: sub}, nil
}

// WatchZetaReverted is a free log subscription operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaInterfaces *ZetaInterfacesFilterer) WatchZetaReverted(opts *bind.WatchOpts, sink chan<- *ZetaInterfacesZetaReverted, destinationChainId []*big.Int, internalSendHash [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaInterfaces.contract.WatchLogs(opts, "ZetaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaInterfacesZetaReverted)
				if err := _ZetaInterfaces.contract.UnpackLog(event, "ZetaReverted", log); err != nil {
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

// ParseZetaReverted is a log parse operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaInterfaces *ZetaInterfacesFilterer) ParseZetaReverted(log types.Log) (*ZetaInterfacesZetaReverted, error) {
	event := new(ZetaInterfacesZetaReverted)
	if err := _ZetaInterfaces.contract.UnpackLog(event, "ZetaReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
