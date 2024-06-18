// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package receiver

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

// ReceiverMetaData contains all meta data concerning the Receiver contract.
var ReceiverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"ReceivedA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"strs\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nums\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"ReceivedB\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"receiveA\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"strs\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nums\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"receiveB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610906806100206000396000f3fe6080604052600436106100295760003560e01c80636fa220ad1461002e578063f5db6b391461004a575b600080fd5b61004860048036038101906100439190610378565b610073565b005b34801561005657600080fd5b50610071600480360381019061006c91906102ed565b6100b7565b005b7f87d67858b5cc03bdd16a3fcc56773bd59410e946cff7193cf374402c6e8fb6ee33348585856040516100aa9594939291906105e7565b60405180910390a1505050565b7f463b3e5d6969d17b19e10c4ca95f5f1b6fdef8678f980af98071bd38f0d885c2338484846040516100ec9493929190610594565b60405180910390a1505050565b600061010c61010784610666565b610641565b9050808382526020820190508285602086028201111561012f5761012e61087d565b5b60005b8581101561017d57813567ffffffffffffffff81111561015557610154610878565b5b80860161016289826102aa565b85526020850194506020840193505050600181019050610132565b5050509392505050565b600061019a61019584610692565b610641565b905080838252602082019050828560208602820111156101bd576101bc61087d565b5b60005b858110156101ed57816101d388826102d8565b8452602084019350602083019250506001810190506101c0565b5050509392505050565b600061020a610205846106be565b610641565b90508281526020810184848401111561022657610225610882565b5b6102318482856107d6565b509392505050565b600082601f83011261024e5761024d610878565b5b813561025e8482602086016100f9565b91505092915050565b600082601f83011261027c5761027b610878565b5b813561028c848260208601610187565b91505092915050565b6000813590506102a4816108a2565b92915050565b600082601f8301126102bf576102be610878565b5b81356102cf8482602086016101f7565b91505092915050565b6000813590506102e7816108b9565b92915050565b6000806000606084860312156103065761030561088c565b5b600084013567ffffffffffffffff81111561032457610323610887565b5b61033086828701610239565b935050602084013567ffffffffffffffff81111561035157610350610887565b5b61035d86828701610267565b925050604061036e86828701610295565b9150509250925092565b6000806000606084860312156103915761039061088c565b5b600084013567ffffffffffffffff8111156103af576103ae610887565b5b6103bb868287016102aa565b93505060206103cc868287016102d8565b92505060406103dd86828701610295565b9150509250925092565b60006103f38383610504565b905092915050565b60006104078383610576565b60208301905092915050565b61041c8161078e565b82525050565b600061042d8261070f565b610437818561074a565b935083602082028501610449856106ef565b8060005b85811015610485578484038952815161046685826103e7565b945061047183610730565b925060208a0199505060018101905061044d565b50829750879550505050505092915050565b60006104a28261071a565b6104ac818561075b565b93506104b7836106ff565b8060005b838110156104e85781516104cf88826103fb565b97506104da8361073d565b9250506001810190506104bb565b5085935050505092915050565b6104fe816107a0565b82525050565b600061050f82610725565b610519818561076c565b93506105298185602086016107e5565b61053281610891565b840191505092915050565b600061054882610725565b610552818561077d565b93506105628185602086016107e5565b61056b81610891565b840191505092915050565b61057f816107cc565b82525050565b61058e816107cc565b82525050565b60006080820190506105a96000830187610413565b81810360208301526105bb8186610422565b905081810360408301526105cf8185610497565b90506105de60608301846104f5565b95945050505050565b600060a0820190506105fc6000830188610413565b6106096020830187610585565b818103604083015261061b818661053d565b905061062a6060830185610585565b61063760808301846104f5565b9695505050505050565b600061064b61065c565b90506106578282610818565b919050565b6000604051905090565b600067ffffffffffffffff82111561068157610680610849565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156106ad576106ac610849565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156106d9576106d8610849565b5b6106e282610891565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610799826107ac565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156108035780820151818401526020810190506107e8565b83811115610812576000848401525b50505050565b61082182610891565b810181811067ffffffffffffffff821117156108405761083f610849565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6108ab816107a0565b81146108b657600080fd5b50565b6108c2816107cc565b81146108cd57600080fd5b5056fea2646970667358221220a471fde0547689c02df444b3d3f058d18c9dd15c58325e79133a63a1cc3608dd64736f6c63430008070033",
}

// ReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiverMetaData.ABI instead.
var ReceiverABI = ReceiverMetaData.ABI

// ReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReceiverMetaData.Bin instead.
var ReceiverBin = ReceiverMetaData.Bin

// DeployReceiver deploys a new Ethereum contract, binding an instance of Receiver to it.
func DeployReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Receiver, error) {
	parsed, err := ReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Receiver{ReceiverCaller: ReceiverCaller{contract: contract}, ReceiverTransactor: ReceiverTransactor{contract: contract}, ReceiverFilterer: ReceiverFilterer{contract: contract}}, nil
}

// Receiver is an auto generated Go binding around an Ethereum contract.
type Receiver struct {
	ReceiverCaller     // Read-only binding to the contract
	ReceiverTransactor // Write-only binding to the contract
	ReceiverFilterer   // Log filterer for contract events
}

// ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiverSession struct {
	Contract     *Receiver         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiverCallerSession struct {
	Contract *ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiverTransactorSession struct {
	Contract     *ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiverRaw struct {
	Contract *Receiver // Generic contract binding to access the raw methods on
}

// ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiverCallerRaw struct {
	Contract *ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiverTransactorRaw struct {
	Contract *ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiver creates a new instance of Receiver, bound to a specific deployed contract.
func NewReceiver(address common.Address, backend bind.ContractBackend) (*Receiver, error) {
	contract, err := bindReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Receiver{ReceiverCaller: ReceiverCaller{contract: contract}, ReceiverTransactor: ReceiverTransactor{contract: contract}, ReceiverFilterer: ReceiverFilterer{contract: contract}}, nil
}

// NewReceiverCaller creates a new read-only instance of Receiver, bound to a specific deployed contract.
func NewReceiverCaller(address common.Address, caller bind.ContractCaller) (*ReceiverCaller, error) {
	contract, err := bindReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverCaller{contract: contract}, nil
}

// NewReceiverTransactor creates a new write-only instance of Receiver, bound to a specific deployed contract.
func NewReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiverTransactor, error) {
	contract, err := bindReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverTransactor{contract: contract}, nil
}

// NewReceiverFilterer creates a new log filterer instance of Receiver, bound to a specific deployed contract.
func NewReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiverFilterer, error) {
	contract, err := bindReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiverFilterer{contract: contract}, nil
}

// bindReceiver binds a generic wrapper to an already deployed contract.
func bindReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Receiver *ReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Receiver.Contract.ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Receiver *ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Receiver *ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Receiver *ReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Receiver *ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Receiver *ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Receiver.Contract.contract.Transact(opts, method, params...)
}

// ReceiveA is a paid mutator transaction binding the contract method 0x6fa220ad.
//
// Solidity: function receiveA(string str, uint256 num, bool flag) payable returns()
func (_Receiver *ReceiverTransactor) ReceiveA(opts *bind.TransactOpts, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _Receiver.contract.Transact(opts, "receiveA", str, num, flag)
}

// ReceiveA is a paid mutator transaction binding the contract method 0x6fa220ad.
//
// Solidity: function receiveA(string str, uint256 num, bool flag) payable returns()
func (_Receiver *ReceiverSession) ReceiveA(str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiveA(&_Receiver.TransactOpts, str, num, flag)
}

// ReceiveA is a paid mutator transaction binding the contract method 0x6fa220ad.
//
// Solidity: function receiveA(string str, uint256 num, bool flag) payable returns()
func (_Receiver *ReceiverTransactorSession) ReceiveA(str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiveA(&_Receiver.TransactOpts, str, num, flag)
}

// ReceiveB is a paid mutator transaction binding the contract method 0xf5db6b39.
//
// Solidity: function receiveB(string[] strs, uint256[] nums, bool flag) returns()
func (_Receiver *ReceiverTransactor) ReceiveB(opts *bind.TransactOpts, strs []string, nums []*big.Int, flag bool) (*types.Transaction, error) {
	return _Receiver.contract.Transact(opts, "receiveB", strs, nums, flag)
}

// ReceiveB is a paid mutator transaction binding the contract method 0xf5db6b39.
//
// Solidity: function receiveB(string[] strs, uint256[] nums, bool flag) returns()
func (_Receiver *ReceiverSession) ReceiveB(strs []string, nums []*big.Int, flag bool) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiveB(&_Receiver.TransactOpts, strs, nums, flag)
}

// ReceiveB is a paid mutator transaction binding the contract method 0xf5db6b39.
//
// Solidity: function receiveB(string[] strs, uint256[] nums, bool flag) returns()
func (_Receiver *ReceiverTransactorSession) ReceiveB(strs []string, nums []*big.Int, flag bool) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiveB(&_Receiver.TransactOpts, strs, nums, flag)
}

// ReceiverReceivedAIterator is returned from FilterReceivedA and is used to iterate over the raw logs and unpacked data for ReceivedA events raised by the Receiver contract.
type ReceiverReceivedAIterator struct {
	Event *ReceiverReceivedA // Event containing the contract specifics and raw log

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
func (it *ReceiverReceivedAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiverReceivedA)
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
		it.Event = new(ReceiverReceivedA)
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
func (it *ReceiverReceivedAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiverReceivedAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiverReceivedA represents a ReceivedA event raised by the Receiver contract.
type ReceiverReceivedA struct {
	Sender common.Address
	Value  *big.Int
	Str    string
	Num    *big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedA is a free log retrieval operation binding the contract event 0x87d67858b5cc03bdd16a3fcc56773bd59410e946cff7193cf374402c6e8fb6ee.
//
// Solidity: event ReceivedA(address sender, uint256 value, string str, uint256 num, bool flag)
func (_Receiver *ReceiverFilterer) FilterReceivedA(opts *bind.FilterOpts) (*ReceiverReceivedAIterator, error) {

	logs, sub, err := _Receiver.contract.FilterLogs(opts, "ReceivedA")
	if err != nil {
		return nil, err
	}
	return &ReceiverReceivedAIterator{contract: _Receiver.contract, event: "ReceivedA", logs: logs, sub: sub}, nil
}

// WatchReceivedA is a free log subscription operation binding the contract event 0x87d67858b5cc03bdd16a3fcc56773bd59410e946cff7193cf374402c6e8fb6ee.
//
// Solidity: event ReceivedA(address sender, uint256 value, string str, uint256 num, bool flag)
func (_Receiver *ReceiverFilterer) WatchReceivedA(opts *bind.WatchOpts, sink chan<- *ReceiverReceivedA) (event.Subscription, error) {

	logs, sub, err := _Receiver.contract.WatchLogs(opts, "ReceivedA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiverReceivedA)
				if err := _Receiver.contract.UnpackLog(event, "ReceivedA", log); err != nil {
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

// ParseReceivedA is a log parse operation binding the contract event 0x87d67858b5cc03bdd16a3fcc56773bd59410e946cff7193cf374402c6e8fb6ee.
//
// Solidity: event ReceivedA(address sender, uint256 value, string str, uint256 num, bool flag)
func (_Receiver *ReceiverFilterer) ParseReceivedA(log types.Log) (*ReceiverReceivedA, error) {
	event := new(ReceiverReceivedA)
	if err := _Receiver.contract.UnpackLog(event, "ReceivedA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiverReceivedBIterator is returned from FilterReceivedB and is used to iterate over the raw logs and unpacked data for ReceivedB events raised by the Receiver contract.
type ReceiverReceivedBIterator struct {
	Event *ReceiverReceivedB // Event containing the contract specifics and raw log

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
func (it *ReceiverReceivedBIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiverReceivedB)
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
		it.Event = new(ReceiverReceivedB)
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
func (it *ReceiverReceivedBIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiverReceivedBIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiverReceivedB represents a ReceivedB event raised by the Receiver contract.
type ReceiverReceivedB struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedB is a free log retrieval operation binding the contract event 0x463b3e5d6969d17b19e10c4ca95f5f1b6fdef8678f980af98071bd38f0d885c2.
//
// Solidity: event ReceivedB(address sender, string[] strs, uint256[] nums, bool flag)
func (_Receiver *ReceiverFilterer) FilterReceivedB(opts *bind.FilterOpts) (*ReceiverReceivedBIterator, error) {

	logs, sub, err := _Receiver.contract.FilterLogs(opts, "ReceivedB")
	if err != nil {
		return nil, err
	}
	return &ReceiverReceivedBIterator{contract: _Receiver.contract, event: "ReceivedB", logs: logs, sub: sub}, nil
}

// WatchReceivedB is a free log subscription operation binding the contract event 0x463b3e5d6969d17b19e10c4ca95f5f1b6fdef8678f980af98071bd38f0d885c2.
//
// Solidity: event ReceivedB(address sender, string[] strs, uint256[] nums, bool flag)
func (_Receiver *ReceiverFilterer) WatchReceivedB(opts *bind.WatchOpts, sink chan<- *ReceiverReceivedB) (event.Subscription, error) {

	logs, sub, err := _Receiver.contract.WatchLogs(opts, "ReceivedB")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiverReceivedB)
				if err := _Receiver.contract.UnpackLog(event, "ReceivedB", log); err != nil {
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

// ParseReceivedB is a log parse operation binding the contract event 0x463b3e5d6969d17b19e10c4ca95f5f1b6fdef8678f980af98071bd38f0d885c2.
//
// Solidity: event ReceivedB(address sender, string[] strs, uint256[] nums, bool flag)
func (_Receiver *ReceiverFilterer) ParseReceivedB(log types.Log) (*ReceiverReceivedB, error) {
	event := new(ReceiverReceivedB)
	if err := _Receiver.contract.UnpackLog(event, "ReceivedB", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
