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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"ReceivedA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"strs\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nums\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"ReceivedB\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"ReceivedC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ReceivedD\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"receiveA\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"strs\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nums\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"receiveB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"receiveC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610bbf806100206000396000f3fe60806040526004361061003f5760003560e01c806352df08b6146100445780636fa220ad1461006d57806386c4519214610089578063f5db6b39146100a0575b600080fd5b34801561005057600080fd5b5061006b6004803603810190610066919061059f565b6100c9565b005b61008760048036038101906100829190610530565b61019b565b005b34801561009557600080fd5b5061009e6101df565b005b3480156100ac57600080fd5b506100c760048036038101906100c29190610478565b610218565b005b8173ffffffffffffffffffffffffffffffffffffffff166323b872dd3383866040518463ffffffff1660e01b8152600401610106939291906107ba565b602060405180830381600087803b15801561012057600080fd5b505af1158015610134573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101589190610503565b507fe8077791e8d0f63b9e4c0b6d386cbbf6c65e1acb2a103019edba9d1cc0b329003384848460405161018e9493929190610844565b60405180910390a1505050565b7f87d67858b5cc03bdd16a3fcc56773bd59410e946cff7193cf374402c6e8fb6ee33348585856040516101d2959493929190610889565b60405180910390a1505050565b7fcf0e6f18d967cb5a3ca7781d74b1d66411d1b8984e2dd2a066709c204a66d8623360405161020e919061079f565b60405180910390a1565b7f463b3e5d6969d17b19e10c4ca95f5f1b6fdef8678f980af98071bd38f0d885c23384848460405161024d94939291906107f1565b60405180910390a1505050565b600061026d61026884610908565b6108e3565b905080838252602082019050828560208602820111156102905761028f610b1f565b5b60005b858110156102de57813567ffffffffffffffff8111156102b6576102b5610b1a565b5b8086016102c38982610435565b85526020850194506020840193505050600181019050610293565b5050509392505050565b60006102fb6102f684610934565b6108e3565b9050808382526020820190508285602086028201111561031e5761031d610b1f565b5b60005b8581101561034e57816103348882610463565b845260208401935060208301925050600181019050610321565b5050509392505050565b600061036b61036684610960565b6108e3565b90508281526020810184848401111561038757610386610b24565b5b610392848285610a78565b509392505050565b6000813590506103a981610b44565b92915050565b600082601f8301126103c4576103c3610b1a565b5b81356103d484826020860161025a565b91505092915050565b600082601f8301126103f2576103f1610b1a565b5b81356104028482602086016102e8565b91505092915050565b60008135905061041a81610b5b565b92915050565b60008151905061042f81610b5b565b92915050565b600082601f83011261044a57610449610b1a565b5b813561045a848260208601610358565b91505092915050565b60008135905061047281610b72565b92915050565b60008060006060848603121561049157610490610b2e565b5b600084013567ffffffffffffffff8111156104af576104ae610b29565b5b6104bb868287016103af565b935050602084013567ffffffffffffffff8111156104dc576104db610b29565b5b6104e8868287016103dd565b92505060406104f98682870161040b565b9150509250925092565b60006020828403121561051957610518610b2e565b5b600061052784828501610420565b91505092915050565b60008060006060848603121561054957610548610b2e565b5b600084013567ffffffffffffffff81111561056757610566610b29565b5b61057386828701610435565b935050602061058486828701610463565b92505060406105958682870161040b565b9150509250925092565b6000806000606084860312156105b8576105b7610b2e565b5b60006105c686828701610463565b93505060206105d78682870161039a565b92505060406105e88682870161039a565b9150509250925092565b60006105fe838361070f565b905092915050565b60006106128383610781565b60208301905092915050565b61062781610a30565b82525050565b6000610638826109b1565b61064281856109ec565b93508360208202850161065485610991565b8060005b85811015610690578484038952815161067185826105f2565b945061067c836109d2565b925060208a01995050600181019050610658565b50829750879550505050505092915050565b60006106ad826109bc565b6106b781856109fd565b93506106c2836109a1565b8060005b838110156106f35781516106da8882610606565b97506106e5836109df565b9250506001810190506106c6565b5085935050505092915050565b61070981610a42565b82525050565b600061071a826109c7565b6107248185610a0e565b9350610734818560208601610a87565b61073d81610b33565b840191505092915050565b6000610753826109c7565b61075d8185610a1f565b935061076d818560208601610a87565b61077681610b33565b840191505092915050565b61078a81610a6e565b82525050565b61079981610a6e565b82525050565b60006020820190506107b4600083018461061e565b92915050565b60006060820190506107cf600083018661061e565b6107dc602083018561061e565b6107e96040830184610790565b949350505050565b6000608082019050610806600083018761061e565b8181036020830152610818818661062d565b9050818103604083015261082c81856106a2565b905061083b6060830184610700565b95945050505050565b6000608082019050610859600083018761061e565b6108666020830186610790565b610873604083018561061e565b610880606083018461061e565b95945050505050565b600060a08201905061089e600083018861061e565b6108ab6020830187610790565b81810360408301526108bd8186610748565b90506108cc6060830185610790565b6108d96080830184610700565b9695505050505050565b60006108ed6108fe565b90506108f98282610aba565b919050565b6000604051905090565b600067ffffffffffffffff82111561092357610922610aeb565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561094f5761094e610aeb565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561097b5761097a610aeb565b5b61098482610b33565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610a3b82610a4e565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610aa5578082015181840152602081019050610a8a565b83811115610ab4576000848401525b50505050565b610ac382610b33565b810181811067ffffffffffffffff82111715610ae257610ae1610aeb565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b610b4d81610a30565b8114610b5857600080fd5b50565b610b6481610a42565b8114610b6f57600080fd5b50565b610b7b81610a6e565b8114610b8657600080fd5b5056fea264697066735822122071e33ff7a639cea7ba0e3fa2cb4106a312f1b8080d32e972b0ef4823ad974cf264736f6c63430008070033",
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

// ReceiveC is a paid mutator transaction binding the contract method 0x52df08b6.
//
// Solidity: function receiveC(uint256 amount, address token, address destination) returns()
func (_Receiver *ReceiverTransactor) ReceiveC(opts *bind.TransactOpts, amount *big.Int, token common.Address, destination common.Address) (*types.Transaction, error) {
	return _Receiver.contract.Transact(opts, "receiveC", amount, token, destination)
}

// ReceiveC is a paid mutator transaction binding the contract method 0x52df08b6.
//
// Solidity: function receiveC(uint256 amount, address token, address destination) returns()
func (_Receiver *ReceiverSession) ReceiveC(amount *big.Int, token common.Address, destination common.Address) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiveC(&_Receiver.TransactOpts, amount, token, destination)
}

// ReceiveC is a paid mutator transaction binding the contract method 0x52df08b6.
//
// Solidity: function receiveC(uint256 amount, address token, address destination) returns()
func (_Receiver *ReceiverTransactorSession) ReceiveC(amount *big.Int, token common.Address, destination common.Address) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiveC(&_Receiver.TransactOpts, amount, token, destination)
}

// ReceiveD is a paid mutator transaction binding the contract method 0x86c45192.
//
// Solidity: function receiveD() returns()
func (_Receiver *ReceiverTransactor) ReceiveD(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Receiver.contract.Transact(opts, "receiveD")
}

// ReceiveD is a paid mutator transaction binding the contract method 0x86c45192.
//
// Solidity: function receiveD() returns()
func (_Receiver *ReceiverSession) ReceiveD() (*types.Transaction, error) {
	return _Receiver.Contract.ReceiveD(&_Receiver.TransactOpts)
}

// ReceiveD is a paid mutator transaction binding the contract method 0x86c45192.
//
// Solidity: function receiveD() returns()
func (_Receiver *ReceiverTransactorSession) ReceiveD() (*types.Transaction, error) {
	return _Receiver.Contract.ReceiveD(&_Receiver.TransactOpts)
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

// ReceiverReceivedCIterator is returned from FilterReceivedC and is used to iterate over the raw logs and unpacked data for ReceivedC events raised by the Receiver contract.
type ReceiverReceivedCIterator struct {
	Event *ReceiverReceivedC // Event containing the contract specifics and raw log

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
func (it *ReceiverReceivedCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiverReceivedC)
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
		it.Event = new(ReceiverReceivedC)
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
func (it *ReceiverReceivedCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiverReceivedCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiverReceivedC represents a ReceivedC event raised by the Receiver contract.
type ReceiverReceivedC struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedC is a free log retrieval operation binding the contract event 0xe8077791e8d0f63b9e4c0b6d386cbbf6c65e1acb2a103019edba9d1cc0b32900.
//
// Solidity: event ReceivedC(address sender, uint256 amount, address token, address destination)
func (_Receiver *ReceiverFilterer) FilterReceivedC(opts *bind.FilterOpts) (*ReceiverReceivedCIterator, error) {

	logs, sub, err := _Receiver.contract.FilterLogs(opts, "ReceivedC")
	if err != nil {
		return nil, err
	}
	return &ReceiverReceivedCIterator{contract: _Receiver.contract, event: "ReceivedC", logs: logs, sub: sub}, nil
}

// WatchReceivedC is a free log subscription operation binding the contract event 0xe8077791e8d0f63b9e4c0b6d386cbbf6c65e1acb2a103019edba9d1cc0b32900.
//
// Solidity: event ReceivedC(address sender, uint256 amount, address token, address destination)
func (_Receiver *ReceiverFilterer) WatchReceivedC(opts *bind.WatchOpts, sink chan<- *ReceiverReceivedC) (event.Subscription, error) {

	logs, sub, err := _Receiver.contract.WatchLogs(opts, "ReceivedC")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiverReceivedC)
				if err := _Receiver.contract.UnpackLog(event, "ReceivedC", log); err != nil {
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

// ParseReceivedC is a log parse operation binding the contract event 0xe8077791e8d0f63b9e4c0b6d386cbbf6c65e1acb2a103019edba9d1cc0b32900.
//
// Solidity: event ReceivedC(address sender, uint256 amount, address token, address destination)
func (_Receiver *ReceiverFilterer) ParseReceivedC(log types.Log) (*ReceiverReceivedC, error) {
	event := new(ReceiverReceivedC)
	if err := _Receiver.contract.UnpackLog(event, "ReceivedC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiverReceivedDIterator is returned from FilterReceivedD and is used to iterate over the raw logs and unpacked data for ReceivedD events raised by the Receiver contract.
type ReceiverReceivedDIterator struct {
	Event *ReceiverReceivedD // Event containing the contract specifics and raw log

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
func (it *ReceiverReceivedDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiverReceivedD)
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
		it.Event = new(ReceiverReceivedD)
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
func (it *ReceiverReceivedDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiverReceivedDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiverReceivedD represents a ReceivedD event raised by the Receiver contract.
type ReceiverReceivedD struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedD is a free log retrieval operation binding the contract event 0xcf0e6f18d967cb5a3ca7781d74b1d66411d1b8984e2dd2a066709c204a66d862.
//
// Solidity: event ReceivedD(address sender)
func (_Receiver *ReceiverFilterer) FilterReceivedD(opts *bind.FilterOpts) (*ReceiverReceivedDIterator, error) {

	logs, sub, err := _Receiver.contract.FilterLogs(opts, "ReceivedD")
	if err != nil {
		return nil, err
	}
	return &ReceiverReceivedDIterator{contract: _Receiver.contract, event: "ReceivedD", logs: logs, sub: sub}, nil
}

// WatchReceivedD is a free log subscription operation binding the contract event 0xcf0e6f18d967cb5a3ca7781d74b1d66411d1b8984e2dd2a066709c204a66d862.
//
// Solidity: event ReceivedD(address sender)
func (_Receiver *ReceiverFilterer) WatchReceivedD(opts *bind.WatchOpts, sink chan<- *ReceiverReceivedD) (event.Subscription, error) {

	logs, sub, err := _Receiver.contract.WatchLogs(opts, "ReceivedD")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiverReceivedD)
				if err := _Receiver.contract.UnpackLog(event, "ReceivedD", log); err != nil {
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

// ParseReceivedD is a log parse operation binding the contract event 0xcf0e6f18d967cb5a3ca7781d74b1d66411d1b8984e2dd2a066709c204a66d862.
//
// Solidity: event ReceivedD(address sender)
func (_Receiver *ReceiverFilterer) ParseReceivedD(log types.Log) (*ReceiverReceivedD, error) {
	event := new(ReceiverReceivedD)
	if err := _Receiver.contract.UnpackLog(event, "ReceivedD", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
