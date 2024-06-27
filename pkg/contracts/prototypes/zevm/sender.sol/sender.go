// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sender

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

// SenderMetaData contains all meta data concerning the Sender contract.
var SenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Call\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"sendToReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516107e43803806107e48339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610108565b600081519050610087816100f1565b92915050565b6000602082840312156100a3576100a26100ec565b5b60006100b184828501610078565b91505092915050565b60006100c5826100cc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6100fa816100ba565b811461010557600080fd5b50565b6106cd806101176000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063116191b61461003b578063f948d35914610059575b600080fd5b610043610075565b604051610050919061040b565b60405180910390f35b610073600480360381019061006e91906102cd565b610099565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008383836040516024016100b09392919061045d565b6040516020818303038152906040527f6fa220ad000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630ac7c44c86836040518363ffffffff1660e01b815260040161018a929190610426565b600060405180830381600087803b1580156101a457600080fd5b505af11580156101b8573d6000803e3d6000fd5b505050505050505050565b60006101d66101d1846104c0565b61049b565b9050828152602081018484840111156101f2576101f1610649565b5b6101fd8482856105a2565b509392505050565b6000610218610213846104f1565b61049b565b90508281526020810184848401111561023457610233610649565b5b61023f8482856105a2565b509392505050565b60008135905061025681610669565b92915050565b600082601f83011261027157610270610644565b5b81356102818482602086016101c3565b91505092915050565b600082601f83011261029f5761029e610644565b5b81356102af848260208601610205565b91505092915050565b6000813590506102c781610680565b92915050565b600080600080608085870312156102e7576102e6610653565b5b600085013567ffffffffffffffff8111156103055761030461064e565b5b6103118782880161025c565b945050602085013567ffffffffffffffff8111156103325761033161064e565b5b61033e8782880161028a565b935050604061034f878288016102b8565b925050606061036087828801610247565b91505092959194509250565b6103758161055a565b82525050565b6103848161056c565b82525050565b600061039582610522565b61039f8185610538565b93506103af8185602086016105b1565b6103b881610658565b840191505092915050565b60006103ce8261052d565b6103d88185610549565b93506103e88185602086016105b1565b6103f181610658565b840191505092915050565b61040581610598565b82525050565b6000602082019050610420600083018461036c565b92915050565b60006040820190508181036000830152610440818561038a565b90508181036020830152610454818461038a565b90509392505050565b6000606082019050818103600083015261047781866103c3565b905061048660208301856103fc565b610493604083018461037b565b949350505050565b60006104a56104b6565b90506104b182826105e4565b919050565b6000604051905090565b600067ffffffffffffffff8211156104db576104da610615565b5b6104e482610658565b9050602081019050919050565b600067ffffffffffffffff82111561050c5761050b610615565b5b61051582610658565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600061056582610578565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156105cf5780820151818401526020810190506105b4565b838111156105de576000848401525b50505050565b6105ed82610658565b810181811067ffffffffffffffff8211171561060c5761060b610615565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6106728161056c565b811461067d57600080fd5b50565b61068981610598565b811461069457600080fd5b5056fea2646970667358221220cdaccd861067daccfac02009db402ee2d2aca72fb7ae4d1786005c3b291c2e3b64736f6c63430008070033",
}

// SenderABI is the input ABI used to generate the binding from.
// Deprecated: Use SenderMetaData.ABI instead.
var SenderABI = SenderMetaData.ABI

// SenderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SenderMetaData.Bin instead.
var SenderBin = SenderMetaData.Bin

// DeploySender deploys a new Ethereum contract, binding an instance of Sender to it.
func DeploySender(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address) (common.Address, *types.Transaction, *Sender, error) {
	parsed, err := SenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenderBin), backend, _gateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sender{SenderCaller: SenderCaller{contract: contract}, SenderTransactor: SenderTransactor{contract: contract}, SenderFilterer: SenderFilterer{contract: contract}}, nil
}

// Sender is an auto generated Go binding around an Ethereum contract.
type Sender struct {
	SenderCaller     // Read-only binding to the contract
	SenderTransactor // Write-only binding to the contract
	SenderFilterer   // Log filterer for contract events
}

// SenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenderSession struct {
	Contract     *Sender           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenderCallerSession struct {
	Contract *SenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenderTransactorSession struct {
	Contract     *SenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenderRaw struct {
	Contract *Sender // Generic contract binding to access the raw methods on
}

// SenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenderCallerRaw struct {
	Contract *SenderCaller // Generic read-only contract binding to access the raw methods on
}

// SenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenderTransactorRaw struct {
	Contract *SenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSender creates a new instance of Sender, bound to a specific deployed contract.
func NewSender(address common.Address, backend bind.ContractBackend) (*Sender, error) {
	contract, err := bindSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sender{SenderCaller: SenderCaller{contract: contract}, SenderTransactor: SenderTransactor{contract: contract}, SenderFilterer: SenderFilterer{contract: contract}}, nil
}

// NewSenderCaller creates a new read-only instance of Sender, bound to a specific deployed contract.
func NewSenderCaller(address common.Address, caller bind.ContractCaller) (*SenderCaller, error) {
	contract, err := bindSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenderCaller{contract: contract}, nil
}

// NewSenderTransactor creates a new write-only instance of Sender, bound to a specific deployed contract.
func NewSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*SenderTransactor, error) {
	contract, err := bindSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenderTransactor{contract: contract}, nil
}

// NewSenderFilterer creates a new log filterer instance of Sender, bound to a specific deployed contract.
func NewSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*SenderFilterer, error) {
	contract, err := bindSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenderFilterer{contract: contract}, nil
}

// bindSender binds a generic wrapper to an already deployed contract.
func bindSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SenderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sender *SenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sender.Contract.SenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sender *SenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sender.Contract.SenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sender *SenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sender.Contract.SenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sender *SenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sender *SenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sender *SenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sender.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Sender *SenderCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sender.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Sender *SenderSession) Gateway() (common.Address, error) {
	return _Sender.Contract.Gateway(&_Sender.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Sender *SenderCallerSession) Gateway() (common.Address, error) {
	return _Sender.Contract.Gateway(&_Sender.CallOpts)
}

// SendToReceiver is a paid mutator transaction binding the contract method 0xf948d359.
//
// Solidity: function sendToReceiver(bytes receiver, string str, uint256 num, bool flag) returns()
func (_Sender *SenderTransactor) SendToReceiver(opts *bind.TransactOpts, receiver []byte, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _Sender.contract.Transact(opts, "sendToReceiver", receiver, str, num, flag)
}

// SendToReceiver is a paid mutator transaction binding the contract method 0xf948d359.
//
// Solidity: function sendToReceiver(bytes receiver, string str, uint256 num, bool flag) returns()
func (_Sender *SenderSession) SendToReceiver(receiver []byte, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _Sender.Contract.SendToReceiver(&_Sender.TransactOpts, receiver, str, num, flag)
}

// SendToReceiver is a paid mutator transaction binding the contract method 0xf948d359.
//
// Solidity: function sendToReceiver(bytes receiver, string str, uint256 num, bool flag) returns()
func (_Sender *SenderTransactorSession) SendToReceiver(receiver []byte, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _Sender.Contract.SendToReceiver(&_Sender.TransactOpts, receiver, str, num, flag)
}

// SenderCallIterator is returned from FilterCall and is used to iterate over the raw logs and unpacked data for Call events raised by the Sender contract.
type SenderCallIterator struct {
	Event *SenderCall // Event containing the contract specifics and raw log

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
func (it *SenderCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SenderCall)
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
		it.Event = new(SenderCall)
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
func (it *SenderCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SenderCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SenderCall represents a Call event raised by the Sender contract.
type SenderCall struct {
	Sender   common.Address
	Receiver common.Hash
	Message  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCall is a free log retrieval operation binding the contract event 0x2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f.
//
// Solidity: event Call(address indexed sender, bytes indexed receiver, bytes message)
func (_Sender *SenderFilterer) FilterCall(opts *bind.FilterOpts, sender []common.Address, receiver [][]byte) (*SenderCallIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Sender.contract.FilterLogs(opts, "Call", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &SenderCallIterator{contract: _Sender.contract, event: "Call", logs: logs, sub: sub}, nil
}

// WatchCall is a free log subscription operation binding the contract event 0x2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f.
//
// Solidity: event Call(address indexed sender, bytes indexed receiver, bytes message)
func (_Sender *SenderFilterer) WatchCall(opts *bind.WatchOpts, sink chan<- *SenderCall, sender []common.Address, receiver [][]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Sender.contract.WatchLogs(opts, "Call", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SenderCall)
				if err := _Sender.contract.UnpackLog(event, "Call", log); err != nil {
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

// ParseCall is a log parse operation binding the contract event 0x2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f.
//
// Solidity: event Call(address indexed sender, bytes indexed receiver, bytes message)
func (_Sender *SenderFilterer) ParseCall(log types.Log) (*SenderCall, error) {
	event := new(SenderCall)
	if err := _Sender.contract.UnpackLog(event, "Call", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
