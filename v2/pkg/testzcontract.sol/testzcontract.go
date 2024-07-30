// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testzcontract

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

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// ZContext is an auto generated low-level Go binding around an user-defined struct.
type ZContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// TestZContractMetaData contains all meta data concerning the TestZContract contract.
var TestZContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"onCrossChainCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structzContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRevert\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structrevertContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ContextData\",\"inputs\":[{\"name\":\"origin\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"msgSender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataRevert\",\"inputs\":[{\"name\":\"origin\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"msgSender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b506104c4806100206000396000f3fe60806040526004361061002a5760003560e01c806369582bee14610033578063de43156e1461005357005b3661003157005b005b34801561003f57600080fd5b5061003161004e3660046101ba565b610073565b34801561005f57600080fd5b5061003161006e3660046101ba565b6100ee565b6060811561008a576100878284018461026e565b90505b7ffdc887992b033668833927e252058e468fac0b6bd196d520f09c61b740e999486100b5878061033d565b6100c560408a0160208b016103a2565b896040013533866040516100de969594939291906103c4565b60405180910390a1505050505050565b60608115610105576101028284018461026e565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e6100b5878061033d565b60006060828403121561014257600080fd5b50919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461016c57600080fd5b919050565b60008083601f84011261018357600080fd5b50813567ffffffffffffffff81111561019b57600080fd5b6020830191508360208285010111156101b357600080fd5b9250929050565b6000806000806000608086880312156101d257600080fd5b853567ffffffffffffffff808211156101ea57600080fd5b6101f689838a01610130565b965061020460208901610148565b955060408801359450606088013591508082111561022157600080fd5b5061022e88828901610171565b969995985093965092949392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561028057600080fd5b813567ffffffffffffffff8082111561029857600080fd5b818401915084601f8301126102ac57600080fd5b8135818111156102be576102be61023f565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156103045761030461023f565b8160405282815287602084870101111561031d57600080fd5b826020860160208301376000928101602001929092525095945050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261037257600080fd5b83018035915067ffffffffffffffff82111561038d57600080fd5b6020019150368190038213156101b357600080fd5b6000602082840312156103b457600080fd5b6103bd82610148565b9392505050565b60a081528560a0820152858760c0830137600060c0878301015260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe080601f8901168301602073ffffffffffffffffffffffffffffffffffffffff808a168287015288604087015280881660608701525060c085830301608086015285518060c084015260005b818110156104685787810183015184820160e00152820161044c565b50600060e0828501015260e084601f83011684010194505050505097965050505050505056fea264697066735822122078d900ecbf3fff2425499bb2eafdc9d45545dc5fea40a7cdc33694659008318c64736f6c63430008150033",
}

// TestZContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TestZContractMetaData.ABI instead.
var TestZContractABI = TestZContractMetaData.ABI

// TestZContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestZContractMetaData.Bin instead.
var TestZContractBin = TestZContractMetaData.Bin

// DeployTestZContract deploys a new Ethereum contract, binding an instance of TestZContract to it.
func DeployTestZContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestZContract, error) {
	parsed, err := TestZContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestZContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestZContract{TestZContractCaller: TestZContractCaller{contract: contract}, TestZContractTransactor: TestZContractTransactor{contract: contract}, TestZContractFilterer: TestZContractFilterer{contract: contract}}, nil
}

// TestZContract is an auto generated Go binding around an Ethereum contract.
type TestZContract struct {
	TestZContractCaller     // Read-only binding to the contract
	TestZContractTransactor // Write-only binding to the contract
	TestZContractFilterer   // Log filterer for contract events
}

// TestZContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestZContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestZContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestZContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestZContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestZContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestZContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestZContractSession struct {
	Contract     *TestZContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestZContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestZContractCallerSession struct {
	Contract *TestZContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestZContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestZContractTransactorSession struct {
	Contract     *TestZContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestZContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestZContractRaw struct {
	Contract *TestZContract // Generic contract binding to access the raw methods on
}

// TestZContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestZContractCallerRaw struct {
	Contract *TestZContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestZContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestZContractTransactorRaw struct {
	Contract *TestZContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestZContract creates a new instance of TestZContract, bound to a specific deployed contract.
func NewTestZContract(address common.Address, backend bind.ContractBackend) (*TestZContract, error) {
	contract, err := bindTestZContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestZContract{TestZContractCaller: TestZContractCaller{contract: contract}, TestZContractTransactor: TestZContractTransactor{contract: contract}, TestZContractFilterer: TestZContractFilterer{contract: contract}}, nil
}

// NewTestZContractCaller creates a new read-only instance of TestZContract, bound to a specific deployed contract.
func NewTestZContractCaller(address common.Address, caller bind.ContractCaller) (*TestZContractCaller, error) {
	contract, err := bindTestZContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestZContractCaller{contract: contract}, nil
}

// NewTestZContractTransactor creates a new write-only instance of TestZContract, bound to a specific deployed contract.
func NewTestZContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestZContractTransactor, error) {
	contract, err := bindTestZContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestZContractTransactor{contract: contract}, nil
}

// NewTestZContractFilterer creates a new log filterer instance of TestZContract, bound to a specific deployed contract.
func NewTestZContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestZContractFilterer, error) {
	contract, err := bindTestZContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestZContractFilterer{contract: contract}, nil
}

// bindTestZContract binds a generic wrapper to an already deployed contract.
func bindTestZContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestZContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestZContract *TestZContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestZContract.Contract.TestZContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestZContract *TestZContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestZContract.Contract.TestZContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestZContract *TestZContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestZContract.Contract.TestZContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestZContract *TestZContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestZContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestZContract *TestZContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestZContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestZContract *TestZContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestZContract.Contract.contract.Transact(opts, method, params...)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractTransactor) OnCrossChainCall(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.contract.Transact(opts, "onCrossChainCall", context, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractSession) OnCrossChainCall(context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.OnCrossChainCall(&_TestZContract.TransactOpts, context, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractTransactorSession) OnCrossChainCall(context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.OnCrossChainCall(&_TestZContract.TransactOpts, context, zrc20, amount, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0x69582bee.
//
// Solidity: function onRevert((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractTransactor) OnRevert(opts *bind.TransactOpts, context RevertContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.contract.Transact(opts, "onRevert", context, zrc20, amount, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0x69582bee.
//
// Solidity: function onRevert((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractSession) OnRevert(context RevertContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.OnRevert(&_TestZContract.TransactOpts, context, zrc20, amount, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0x69582bee.
//
// Solidity: function onRevert((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractTransactorSession) OnRevert(context RevertContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.OnRevert(&_TestZContract.TransactOpts, context, zrc20, amount, message)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestZContract *TestZContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TestZContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestZContract *TestZContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.Fallback(&_TestZContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestZContract *TestZContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.Fallback(&_TestZContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestZContract *TestZContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestZContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestZContract *TestZContractSession) Receive() (*types.Transaction, error) {
	return _TestZContract.Contract.Receive(&_TestZContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestZContract *TestZContractTransactorSession) Receive() (*types.Transaction, error) {
	return _TestZContract.Contract.Receive(&_TestZContract.TransactOpts)
}

// TestZContractContextDataIterator is returned from FilterContextData and is used to iterate over the raw logs and unpacked data for ContextData events raised by the TestZContract contract.
type TestZContractContextDataIterator struct {
	Event *TestZContractContextData // Event containing the contract specifics and raw log

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
func (it *TestZContractContextDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestZContractContextData)
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
		it.Event = new(TestZContractContextData)
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
func (it *TestZContractContextDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestZContractContextDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestZContractContextData represents a ContextData event raised by the TestZContract contract.
type TestZContractContextData struct {
	Origin    []byte
	Sender    common.Address
	ChainID   *big.Int
	MsgSender common.Address
	Message   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContextData is a free log retrieval operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) FilterContextData(opts *bind.FilterOpts) (*TestZContractContextDataIterator, error) {

	logs, sub, err := _TestZContract.contract.FilterLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return &TestZContractContextDataIterator{contract: _TestZContract.contract, event: "ContextData", logs: logs, sub: sub}, nil
}

// WatchContextData is a free log subscription operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) WatchContextData(opts *bind.WatchOpts, sink chan<- *TestZContractContextData) (event.Subscription, error) {

	logs, sub, err := _TestZContract.contract.WatchLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestZContractContextData)
				if err := _TestZContract.contract.UnpackLog(event, "ContextData", log); err != nil {
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

// ParseContextData is a log parse operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) ParseContextData(log types.Log) (*TestZContractContextData, error) {
	event := new(TestZContractContextData)
	if err := _TestZContract.contract.UnpackLog(event, "ContextData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestZContractContextDataRevertIterator is returned from FilterContextDataRevert and is used to iterate over the raw logs and unpacked data for ContextDataRevert events raised by the TestZContract contract.
type TestZContractContextDataRevertIterator struct {
	Event *TestZContractContextDataRevert // Event containing the contract specifics and raw log

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
func (it *TestZContractContextDataRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestZContractContextDataRevert)
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
		it.Event = new(TestZContractContextDataRevert)
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
func (it *TestZContractContextDataRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestZContractContextDataRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestZContractContextDataRevert represents a ContextDataRevert event raised by the TestZContract contract.
type TestZContractContextDataRevert struct {
	Origin    []byte
	Sender    common.Address
	ChainID   *big.Int
	MsgSender common.Address
	Message   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContextDataRevert is a free log retrieval operation binding the contract event 0xfdc887992b033668833927e252058e468fac0b6bd196d520f09c61b740e99948.
//
// Solidity: event ContextDataRevert(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) FilterContextDataRevert(opts *bind.FilterOpts) (*TestZContractContextDataRevertIterator, error) {

	logs, sub, err := _TestZContract.contract.FilterLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return &TestZContractContextDataRevertIterator{contract: _TestZContract.contract, event: "ContextDataRevert", logs: logs, sub: sub}, nil
}

// WatchContextDataRevert is a free log subscription operation binding the contract event 0xfdc887992b033668833927e252058e468fac0b6bd196d520f09c61b740e99948.
//
// Solidity: event ContextDataRevert(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) WatchContextDataRevert(opts *bind.WatchOpts, sink chan<- *TestZContractContextDataRevert) (event.Subscription, error) {

	logs, sub, err := _TestZContract.contract.WatchLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestZContractContextDataRevert)
				if err := _TestZContract.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
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

// ParseContextDataRevert is a log parse operation binding the contract event 0xfdc887992b033668833927e252058e468fac0b6bd196d520f09c61b740e99948.
//
// Solidity: event ContextDataRevert(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) ParseContextDataRevert(log types.Log) (*TestZContractContextDataRevert, error) {
	event := new(TestZContractContextDataRevert)
	if err := _TestZContract.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
