// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway

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

// GatewayMetaData contains all meta data concerning the Gateway contract.
var GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ExecutedWithERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"custody\",\"outputs\":[{\"internalType\":\"contractERC20CustodyNew\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeWithERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_custody\",\"type\":\"address\"}],\"name\":\"setCustody\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610aee806100206000396000f3fe60806040526004361061003f5760003560e01c80631cff79cd146100445780635131ab5914610074578063ae7a3a6f146100b1578063dda79b75146100da575b600080fd5b61005e600480360381019061005991906106da565b610105565b60405161006b91906108e1565b60405180910390f35b34801561008057600080fd5b5061009b60048036038101906100969190610652565b610173565b6040516100a891906108e1565b60405180910390f35b3480156100bd57600080fd5b506100d860048036038101906100d39190610625565b61045d565b005b3480156100e657600080fd5b506100ef6104a0565b6040516100fc9190610903565b60405180910390f35b606060006101148585856104c4565b90508473ffffffffffffffffffffffffffffffffffffffff167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f3486866040516101609392919061091e565b60405180910390a2809150509392505050565b60608573ffffffffffffffffffffffffffffffffffffffff1663095ea7b386866040518363ffffffff1660e01b81526004016101b09291906108b8565b602060405180830381600087803b1580156101ca57600080fd5b505af11580156101de573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610202919061073a565b5060006102108685856104c4565b90508673ffffffffffffffffffffffffffffffffffffffff1663095ea7b38760006040518363ffffffff1660e01b815260040161024e92919061088f565b602060405180830381600087803b15801561026857600080fd5b505af115801561027c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102a0919061073a565b5060008773ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016102dc9190610874565b60206040518083038186803b1580156102f457600080fd5b505afa158015610308573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032c9190610767565b905060008111156103e6578773ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b81526004016103929291906108b8565b602060405180830381600087803b1580156103ac57600080fd5b505af11580156103c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e4919061073a565b505b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b73828888886040516104479392919061091e565b60405180910390a3819250505095945050505050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000808573ffffffffffffffffffffffffffffffffffffffff163486866040516104f192919061085b565b60006040518083038185875af1925050503d806000811461052e576040519150601f19603f3d011682016040523d82523d6000602084013e610533565b606091505b50915091508161056f576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80925050509392505050565b60008135905061058a81610a73565b92915050565b60008151905061059f81610a8a565b92915050565b60008083601f8401126105bb576105ba610a4e565b5b8235905067ffffffffffffffff8111156105d8576105d7610a49565b5b6020830191508360018202830111156105f4576105f3610a53565b5b9250929050565b60008135905061060a81610aa1565b92915050565b60008151905061061f81610aa1565b92915050565b60006020828403121561063b5761063a610a5d565b5b60006106498482850161057b565b91505092915050565b60008060008060006080868803121561066e5761066d610a5d565b5b600061067c8882890161057b565b955050602061068d8882890161057b565b945050604061069e888289016105fb565b935050606086013567ffffffffffffffff8111156106bf576106be610a58565b5b6106cb888289016105a5565b92509250509295509295909350565b6000806000604084860312156106f3576106f2610a5d565b5b60006107018682870161057b565b935050602084013567ffffffffffffffff81111561072257610721610a58565b5b61072e868287016105a5565b92509250509250925092565b6000602082840312156107505761074f610a5d565b5b600061075e84828501610590565b91505092915050565b60006020828403121561077d5761077c610a5d565b5b600061078b84828501610610565b91505092915050565b61079d81610977565b82525050565b60006107af838561095b565b93506107bc838584610a07565b6107c583610a62565b840190509392505050565b60006107dc838561096c565b93506107e9838584610a07565b82840190509392505050565b600061080082610950565b61080a818561095b565b935061081a818560208601610a16565b61082381610a62565b840191505092915050565b610837816109bf565b82525050565b610846816109d1565b82525050565b610855816109b5565b82525050565b60006108688284866107d0565b91508190509392505050565b60006020820190506108896000830184610794565b92915050565b60006040820190506108a46000830185610794565b6108b1602083018461083d565b9392505050565b60006040820190506108cd6000830185610794565b6108da602083018461084c565b9392505050565b600060208201905081810360008301526108fb81846107f5565b905092915050565b6000602082019050610918600083018461082e565b92915050565b6000604082019050610933600083018661084c565b81810360208301526109468184866107a3565b9050949350505050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600061098282610995565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006109ca826109e3565b9050919050565b60006109dc826109b5565b9050919050565b60006109ee826109f5565b9050919050565b6000610a0082610995565b9050919050565b82818337600083830152505050565b60005b83811015610a34578082015181840152602081019050610a19565b83811115610a43576000848401525b50505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b610a7c81610977565b8114610a8757600080fd5b50565b610a9381610989565b8114610a9e57600080fd5b50565b610aaa816109b5565b8114610ab557600080fd5b5056fea2646970667358221220c089eeaf89b4386d66c554855cc5118e5526411832a5f014c27c4b0daa4a560764736f6c63430008070033",
}

// GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMetaData.ABI instead.
var GatewayABI = GatewayMetaData.ABI

// GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayMetaData.Bin instead.
var GatewayBin = GatewayMetaData.Bin

// DeployGateway deploys a new Ethereum contract, binding an instance of Gateway to it.
func DeployGateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Gateway, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// Gateway is an auto generated Go binding around an Ethereum contract.
type Gateway struct {
	GatewayCaller     // Read-only binding to the contract
	GatewayTransactor // Write-only binding to the contract
	GatewayFilterer   // Log filterer for contract events
}

// GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewaySession struct {
	Contract     *Gateway          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayCallerSession struct {
	Contract *GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayTransactorSession struct {
	Contract     *GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayRaw struct {
	Contract *Gateway // Generic contract binding to access the raw methods on
}

// GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayCallerRaw struct {
	Contract *GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayTransactorRaw struct {
	Contract *GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGateway creates a new instance of Gateway, bound to a specific deployed contract.
func NewGateway(address common.Address, backend bind.ContractBackend) (*Gateway, error) {
	contract, err := bindGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// NewGatewayCaller creates a new read-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayCaller(address common.Address, caller bind.ContractCaller) (*GatewayCaller, error) {
	contract, err := bindGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayCaller{contract: contract}, nil
}

// NewGatewayTransactor creates a new write-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayTransactor, error) {
	contract, err := bindGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayTransactor{contract: contract}, nil
}

// NewGatewayFilterer creates a new log filterer instance of Gateway, bound to a specific deployed contract.
func NewGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayFilterer, error) {
	contract, err := bindGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayFilterer{contract: contract}, nil
}

// bindGateway binds a generic wrapper to an already deployed contract.
func bindGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transact(opts, method, params...)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_Gateway *GatewayCaller) Custody(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "custody")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_Gateway *GatewaySession) Custody() (common.Address, error) {
	return _Gateway.Contract.Custody(&_Gateway.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_Gateway *GatewayCallerSession) Custody() (common.Address, error) {
	return _Gateway.Contract.Custody(&_Gateway.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_Gateway *GatewayTransactor) Execute(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "execute", destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_Gateway *GatewaySession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _Gateway.Contract.Execute(&_Gateway.TransactOpts, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_Gateway *GatewayTransactorSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _Gateway.Contract.Execute(&_Gateway.TransactOpts, destination, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_Gateway *GatewayTransactor) ExecuteWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "executeWithERC20", token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_Gateway *GatewaySession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Gateway.Contract.ExecuteWithERC20(&_Gateway.TransactOpts, token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns(bytes)
func (_Gateway *GatewayTransactorSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Gateway.Contract.ExecuteWithERC20(&_Gateway.TransactOpts, token, to, amount, data)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_Gateway *GatewayTransactor) SetCustody(opts *bind.TransactOpts, _custody common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setCustody", _custody)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_Gateway *GatewaySession) SetCustody(_custody common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetCustody(&_Gateway.TransactOpts, _custody)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address _custody) returns()
func (_Gateway *GatewayTransactorSession) SetCustody(_custody common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetCustody(&_Gateway.TransactOpts, _custody)
}

// GatewayExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the Gateway contract.
type GatewayExecutedIterator struct {
	Event *GatewayExecuted // Event containing the contract specifics and raw log

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
func (it *GatewayExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayExecuted)
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
		it.Event = new(GatewayExecuted)
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
func (it *GatewayExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayExecuted represents a Executed event raised by the Gateway contract.
type GatewayExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_Gateway *GatewayFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayExecutedIterator{contract: _Gateway.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_Gateway *GatewayFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayExecuted)
				if err := _Gateway.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_Gateway *GatewayFilterer) ParseExecuted(log types.Log) (*GatewayExecuted, error) {
	event := new(GatewayExecuted)
	if err := _Gateway.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the Gateway contract.
type GatewayExecutedWithERC20Iterator struct {
	Event *GatewayExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayExecutedWithERC20)
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
		it.Event = new(GatewayExecutedWithERC20)
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
func (it *GatewayExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayExecutedWithERC20 represents a ExecutedWithERC20 event raised by the Gateway contract.
type GatewayExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_Gateway *GatewayFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayExecutedWithERC20Iterator{contract: _Gateway.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_Gateway *GatewayFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayExecutedWithERC20)
				if err := _Gateway.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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

// ParseExecutedWithERC20 is a log parse operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_Gateway *GatewayFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayExecutedWithERC20, error) {
	event := new(GatewayExecutedWithERC20)
	if err := _Gateway.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
