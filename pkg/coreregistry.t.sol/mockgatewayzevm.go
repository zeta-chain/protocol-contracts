// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coreregistry

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

// CallOptions is an auto generated low-level Go binding around an user-defined struct.
type CallOptions struct {
	GasLimit        *big.Int
	IsArbitraryCall bool
}

// RevertOptions is an auto generated low-level Go binding around an user-defined struct.
type RevertOptions struct {
	RevertAddress    common.Address
	CallOnRevert     bool
	AbortAddress     common.Address
	RevertMessage    []byte
	OnRevertGasLimit *big.Int
}

// MockGatewayZEVMMetaData contains all meta data concerning the MockGatewayZEVM contract.
var MockGatewayZEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallEmitted\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600f57600080fd5b506105298061001f6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806306cb898314610030575b600080fd5b61004361003e366004610159565b610045565b005b7ffa84a91b0ed9555afae4459021634264ec770c42989afa595d13944058e229e586868686868660405161007e9695949392919061042b565b60405180910390a1505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b803573ffffffffffffffffffffffffffffffffffffffff811681146100e157600080fd5b919050565b60008083601f8401126100f857600080fd5b50813567ffffffffffffffff81111561011057600080fd5b60208301915083602082850101111561012857600080fd5b9250929050565b60006040828403121561014157600080fd5b50919050565b600060a0828403121561014157600080fd5b60008060008060008060c0878903121561017257600080fd5b863567ffffffffffffffff81111561018957600080fd5b8701601f8101891361019a57600080fd5b803567ffffffffffffffff8111156101b4576101b461008e565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156102205761022061008e565b6040528181528282016020018b101561023857600080fd5b8160208401602083013760006020838301015280985050505061025d602088016100bd565b9450604087013567ffffffffffffffff81111561027957600080fd5b61028589828a016100e6565b90955093506102999050886060890161012f565b915060a087013567ffffffffffffffff8111156102b557600080fd5b6102c189828a01610147565b9150509295509295509295565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b803580151581146100e157600080fd5b8035825261033760208201610317565b151560208301525050565b73ffffffffffffffffffffffffffffffffffffffff610360826100bd565b16825261036f60208201610317565b1515602083015273ffffffffffffffffffffffffffffffffffffffff610397604083016100bd565b166040830152600060608201357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18336030181126103d457600080fd5b820160208101903567ffffffffffffffff8111156103f157600080fd5b80360382131561040057600080fd5b60a0606086015261041560a0860182846102ce565b6080948501359590940194909452509092915050565b60c08152600087518060c084015260005b81811015610459576020818b0181015160e086840101520161043c565b50600083820160e00152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016104ac602084018973ffffffffffffffffffffffffffffffffffffffff169052565b60e08382030160408401526104c560e0820187896102ce565b90506104d46060840186610327565b82810360a08401526104e68185610342565b999850505050505050505056fea2646970667358221220d180334cfdb237bb267ca76cafaac12773ecf5fb2bb2dd3ae6b6a77201a719da64736f6c634300081a0033",
}

// MockGatewayZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use MockGatewayZEVMMetaData.ABI instead.
var MockGatewayZEVMABI = MockGatewayZEVMMetaData.ABI

// MockGatewayZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockGatewayZEVMMetaData.Bin instead.
var MockGatewayZEVMBin = MockGatewayZEVMMetaData.Bin

// DeployMockGatewayZEVM deploys a new Ethereum contract, binding an instance of MockGatewayZEVM to it.
func DeployMockGatewayZEVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockGatewayZEVM, error) {
	parsed, err := MockGatewayZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockGatewayZEVMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockGatewayZEVM{MockGatewayZEVMCaller: MockGatewayZEVMCaller{contract: contract}, MockGatewayZEVMTransactor: MockGatewayZEVMTransactor{contract: contract}, MockGatewayZEVMFilterer: MockGatewayZEVMFilterer{contract: contract}}, nil
}

// MockGatewayZEVM is an auto generated Go binding around an Ethereum contract.
type MockGatewayZEVM struct {
	MockGatewayZEVMCaller     // Read-only binding to the contract
	MockGatewayZEVMTransactor // Write-only binding to the contract
	MockGatewayZEVMFilterer   // Log filterer for contract events
}

// MockGatewayZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockGatewayZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockGatewayZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockGatewayZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockGatewayZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockGatewayZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockGatewayZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockGatewayZEVMSession struct {
	Contract     *MockGatewayZEVM  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockGatewayZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockGatewayZEVMCallerSession struct {
	Contract *MockGatewayZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MockGatewayZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockGatewayZEVMTransactorSession struct {
	Contract     *MockGatewayZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MockGatewayZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockGatewayZEVMRaw struct {
	Contract *MockGatewayZEVM // Generic contract binding to access the raw methods on
}

// MockGatewayZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockGatewayZEVMCallerRaw struct {
	Contract *MockGatewayZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// MockGatewayZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockGatewayZEVMTransactorRaw struct {
	Contract *MockGatewayZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockGatewayZEVM creates a new instance of MockGatewayZEVM, bound to a specific deployed contract.
func NewMockGatewayZEVM(address common.Address, backend bind.ContractBackend) (*MockGatewayZEVM, error) {
	contract, err := bindMockGatewayZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockGatewayZEVM{MockGatewayZEVMCaller: MockGatewayZEVMCaller{contract: contract}, MockGatewayZEVMTransactor: MockGatewayZEVMTransactor{contract: contract}, MockGatewayZEVMFilterer: MockGatewayZEVMFilterer{contract: contract}}, nil
}

// NewMockGatewayZEVMCaller creates a new read-only instance of MockGatewayZEVM, bound to a specific deployed contract.
func NewMockGatewayZEVMCaller(address common.Address, caller bind.ContractCaller) (*MockGatewayZEVMCaller, error) {
	contract, err := bindMockGatewayZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockGatewayZEVMCaller{contract: contract}, nil
}

// NewMockGatewayZEVMTransactor creates a new write-only instance of MockGatewayZEVM, bound to a specific deployed contract.
func NewMockGatewayZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*MockGatewayZEVMTransactor, error) {
	contract, err := bindMockGatewayZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockGatewayZEVMTransactor{contract: contract}, nil
}

// NewMockGatewayZEVMFilterer creates a new log filterer instance of MockGatewayZEVM, bound to a specific deployed contract.
func NewMockGatewayZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*MockGatewayZEVMFilterer, error) {
	contract, err := bindMockGatewayZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockGatewayZEVMFilterer{contract: contract}, nil
}

// bindMockGatewayZEVM binds a generic wrapper to an already deployed contract.
func bindMockGatewayZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockGatewayZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockGatewayZEVM *MockGatewayZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockGatewayZEVM.Contract.MockGatewayZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockGatewayZEVM *MockGatewayZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockGatewayZEVM.Contract.MockGatewayZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockGatewayZEVM *MockGatewayZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockGatewayZEVM.Contract.MockGatewayZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockGatewayZEVM *MockGatewayZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockGatewayZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockGatewayZEVM *MockGatewayZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockGatewayZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockGatewayZEVM *MockGatewayZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockGatewayZEVM.Contract.contract.Transact(opts, method, params...)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_MockGatewayZEVM *MockGatewayZEVMTransactor) Call(opts *bind.TransactOpts, receiver []byte, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _MockGatewayZEVM.contract.Transact(opts, "call", receiver, zrc20, message, callOptions, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_MockGatewayZEVM *MockGatewayZEVMSession) Call(receiver []byte, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _MockGatewayZEVM.Contract.Call(&_MockGatewayZEVM.TransactOpts, receiver, zrc20, message, callOptions, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_MockGatewayZEVM *MockGatewayZEVMTransactorSession) Call(receiver []byte, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _MockGatewayZEVM.Contract.Call(&_MockGatewayZEVM.TransactOpts, receiver, zrc20, message, callOptions, revertOptions)
}

// MockGatewayZEVMCallEmittedIterator is returned from FilterCallEmitted and is used to iterate over the raw logs and unpacked data for CallEmitted events raised by the MockGatewayZEVM contract.
type MockGatewayZEVMCallEmittedIterator struct {
	Event *MockGatewayZEVMCallEmitted // Event containing the contract specifics and raw log

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
func (it *MockGatewayZEVMCallEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockGatewayZEVMCallEmitted)
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
		it.Event = new(MockGatewayZEVMCallEmitted)
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
func (it *MockGatewayZEVMCallEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockGatewayZEVMCallEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockGatewayZEVMCallEmitted represents a CallEmitted event raised by the MockGatewayZEVM contract.
type MockGatewayZEVMCallEmitted struct {
	Receiver      []byte
	Zrc20         common.Address
	Message       []byte
	CallOptions   CallOptions
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCallEmitted is a free log retrieval operation binding the contract event 0xfa84a91b0ed9555afae4459021634264ec770c42989afa595d13944058e229e5.
//
// Solidity: event CallEmitted(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_MockGatewayZEVM *MockGatewayZEVMFilterer) FilterCallEmitted(opts *bind.FilterOpts) (*MockGatewayZEVMCallEmittedIterator, error) {

	logs, sub, err := _MockGatewayZEVM.contract.FilterLogs(opts, "CallEmitted")
	if err != nil {
		return nil, err
	}
	return &MockGatewayZEVMCallEmittedIterator{contract: _MockGatewayZEVM.contract, event: "CallEmitted", logs: logs, sub: sub}, nil
}

// WatchCallEmitted is a free log subscription operation binding the contract event 0xfa84a91b0ed9555afae4459021634264ec770c42989afa595d13944058e229e5.
//
// Solidity: event CallEmitted(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_MockGatewayZEVM *MockGatewayZEVMFilterer) WatchCallEmitted(opts *bind.WatchOpts, sink chan<- *MockGatewayZEVMCallEmitted) (event.Subscription, error) {

	logs, sub, err := _MockGatewayZEVM.contract.WatchLogs(opts, "CallEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockGatewayZEVMCallEmitted)
				if err := _MockGatewayZEVM.contract.UnpackLog(event, "CallEmitted", log); err != nil {
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

// ParseCallEmitted is a log parse operation binding the contract event 0xfa84a91b0ed9555afae4459021634264ec770c42989afa595d13944058e229e5.
//
// Solidity: event CallEmitted(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_MockGatewayZEVM *MockGatewayZEVMFilterer) ParseCallEmitted(log types.Log) (*MockGatewayZEVMCallEmitted, error) {
	event := new(MockGatewayZEVMCallEmitted)
	if err := _MockGatewayZEVM.contract.UnpackLog(event, "CallEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
