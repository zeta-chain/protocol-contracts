// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20custodynew

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

// ERC20CustodyNewMetaData contains all meta data concerning the ERC20CustodyNew contract.
var ERC20CustodyNewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610a66380380610a668339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610108565b600081519050610087816100f1565b92915050565b6000602082840312156100a3576100a26100ec565b5b60006100b184828501610078565b91505092915050565b60006100c5826100cc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6100fa816100ba565b811461010557600080fd5b50565b61094f806101176000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063116191b61461004657806321fc65f214610064578063d9caed1214610080575b600080fd5b61004e61009c565b60405161005b91906106b6565b60405180910390f35b61007e600480360381019061007991906104e7565b6100c0565b005b61009a60048036038101906100959190610494565b610297565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518363ffffffff1660e01b815260040161011b92919061068d565b602060405180830381600087803b15801561013557600080fd5b505af1158015610149573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061016d919061056f565b5060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635131ab5986868686866040518663ffffffff1660e01b81526004016101cf95949392919061063f565b600060405180830381600087803b1580156101e957600080fd5b505af11580156101fd573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610226919061059c565b508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e858585604051610288939291906106ec565b60405180910390a35050505050565b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b81526004016102d292919061068d565b602060405180830381600087803b1580156102ec57600080fd5b505af1158015610300573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610324919061056f565b508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb8360405161038291906106d1565b60405180910390a3505050565b60006103a261039d84610743565b61071e565b9050828152602081018484840111156103be576103bd6108b4565b5b6103c9848285610812565b509392505050565b6000813590506103e0816108d4565b92915050565b6000815190506103f5816108eb565b92915050565b60008083601f840112610411576104106108aa565b5b8235905067ffffffffffffffff81111561042e5761042d6108a5565b5b60208301915083600182028301111561044a576104496108af565b5b9250929050565b600082601f830112610466576104656108aa565b5b815161047684826020860161038f565b91505092915050565b60008135905061048e81610902565b92915050565b6000806000606084860312156104ad576104ac6108be565b5b60006104bb868287016103d1565b93505060206104cc868287016103d1565b92505060406104dd8682870161047f565b9150509250925092565b600080600080600060808688031215610503576105026108be565b5b6000610511888289016103d1565b9550506020610522888289016103d1565b94505060406105338882890161047f565b935050606086013567ffffffffffffffff811115610554576105536108b9565b5b610560888289016103fb565b92509250509295509295909350565b600060208284031215610585576105846108be565b5b6000610593848285016103e6565b91505092915050565b6000602082840312156105b2576105b16108be565b5b600082015167ffffffffffffffff8111156105d0576105cf6108b9565b5b6105dc84828501610451565b91505092915050565b6105ee81610785565b82525050565b60006106008385610774565b935061060d838584610803565b610616836108c3565b840190509392505050565b61062a816107cd565b82525050565b610639816107c3565b82525050565b600060808201905061065460008301886105e5565b61066160208301876105e5565b61066e6040830186610630565b81810360608301526106818184866105f4565b90509695505050505050565b60006040820190506106a260008301856105e5565b6106af6020830184610630565b9392505050565b60006020820190506106cb6000830184610621565b92915050565b60006020820190506106e66000830184610630565b92915050565b60006040820190506107016000830186610630565b81810360208301526107148184866105f4565b9050949350505050565b6000610728610739565b90506107348282610845565b919050565b6000604051905090565b600067ffffffffffffffff82111561075e5761075d610876565b5b610767826108c3565b9050602081019050919050565b600082825260208201905092915050565b6000610790826107a3565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006107d8826107df565b9050919050565b60006107ea826107f1565b9050919050565b60006107fc826107a3565b9050919050565b82818337600083830152505050565b60005b83811015610830578082015181840152602081019050610815565b8381111561083f576000848401525b50505050565b61084e826108c3565b810181811067ffffffffffffffff8211171561086d5761086c610876565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6108dd81610785565b81146108e857600080fd5b50565b6108f481610797565b81146108ff57600080fd5b50565b61090b816107c3565b811461091657600080fd5b5056fea2646970667358221220824749f88b009234130b10d7ae24afbd12608af72eb829c4db86afe703db98e364736f6c63430008070033",
}

// ERC20CustodyNewABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20CustodyNewMetaData.ABI instead.
var ERC20CustodyNewABI = ERC20CustodyNewMetaData.ABI

// ERC20CustodyNewBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20CustodyNewMetaData.Bin instead.
var ERC20CustodyNewBin = ERC20CustodyNewMetaData.Bin

// DeployERC20CustodyNew deploys a new Ethereum contract, binding an instance of ERC20CustodyNew to it.
func DeployERC20CustodyNew(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address) (common.Address, *types.Transaction, *ERC20CustodyNew, error) {
	parsed, err := ERC20CustodyNewMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20CustodyNewBin), backend, _gateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20CustodyNew{ERC20CustodyNewCaller: ERC20CustodyNewCaller{contract: contract}, ERC20CustodyNewTransactor: ERC20CustodyNewTransactor{contract: contract}, ERC20CustodyNewFilterer: ERC20CustodyNewFilterer{contract: contract}}, nil
}

// ERC20CustodyNew is an auto generated Go binding around an Ethereum contract.
type ERC20CustodyNew struct {
	ERC20CustodyNewCaller     // Read-only binding to the contract
	ERC20CustodyNewTransactor // Write-only binding to the contract
	ERC20CustodyNewFilterer   // Log filterer for contract events
}

// ERC20CustodyNewCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20CustodyNewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyNewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20CustodyNewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyNewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20CustodyNewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyNewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20CustodyNewSession struct {
	Contract     *ERC20CustodyNew  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CustodyNewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CustodyNewCallerSession struct {
	Contract *ERC20CustodyNewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ERC20CustodyNewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20CustodyNewTransactorSession struct {
	Contract     *ERC20CustodyNewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC20CustodyNewRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20CustodyNewRaw struct {
	Contract *ERC20CustodyNew // Generic contract binding to access the raw methods on
}

// ERC20CustodyNewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CustodyNewCallerRaw struct {
	Contract *ERC20CustodyNewCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20CustodyNewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20CustodyNewTransactorRaw struct {
	Contract *ERC20CustodyNewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20CustodyNew creates a new instance of ERC20CustodyNew, bound to a specific deployed contract.
func NewERC20CustodyNew(address common.Address, backend bind.ContractBackend) (*ERC20CustodyNew, error) {
	contract, err := bindERC20CustodyNew(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyNew{ERC20CustodyNewCaller: ERC20CustodyNewCaller{contract: contract}, ERC20CustodyNewTransactor: ERC20CustodyNewTransactor{contract: contract}, ERC20CustodyNewFilterer: ERC20CustodyNewFilterer{contract: contract}}, nil
}

// NewERC20CustodyNewCaller creates a new read-only instance of ERC20CustodyNew, bound to a specific deployed contract.
func NewERC20CustodyNewCaller(address common.Address, caller bind.ContractCaller) (*ERC20CustodyNewCaller, error) {
	contract, err := bindERC20CustodyNew(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyNewCaller{contract: contract}, nil
}

// NewERC20CustodyNewTransactor creates a new write-only instance of ERC20CustodyNew, bound to a specific deployed contract.
func NewERC20CustodyNewTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20CustodyNewTransactor, error) {
	contract, err := bindERC20CustodyNew(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyNewTransactor{contract: contract}, nil
}

// NewERC20CustodyNewFilterer creates a new log filterer instance of ERC20CustodyNew, bound to a specific deployed contract.
func NewERC20CustodyNewFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20CustodyNewFilterer, error) {
	contract, err := bindERC20CustodyNew(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyNewFilterer{contract: contract}, nil
}

// bindERC20CustodyNew binds a generic wrapper to an already deployed contract.
func bindERC20CustodyNew(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20CustodyNewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20CustodyNew *ERC20CustodyNewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20CustodyNew.Contract.ERC20CustodyNewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20CustodyNew *ERC20CustodyNewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyNew.Contract.ERC20CustodyNewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20CustodyNew *ERC20CustodyNewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20CustodyNew.Contract.ERC20CustodyNewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20CustodyNew *ERC20CustodyNewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20CustodyNew.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20CustodyNew *ERC20CustodyNewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyNew.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20CustodyNew *ERC20CustodyNewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20CustodyNew.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ERC20CustodyNew *ERC20CustodyNewCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20CustodyNew.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ERC20CustodyNew *ERC20CustodyNewSession) Gateway() (common.Address, error) {
	return _ERC20CustodyNew.Contract.Gateway(&_ERC20CustodyNew.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ERC20CustodyNew *ERC20CustodyNewCallerSession) Gateway() (common.Address, error) {
	return _ERC20CustodyNew.Contract.Gateway(&_ERC20CustodyNew.CallOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_ERC20CustodyNew *ERC20CustodyNewTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CustodyNew.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_ERC20CustodyNew *ERC20CustodyNewSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CustodyNew.Contract.Withdraw(&_ERC20CustodyNew.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_ERC20CustodyNew *ERC20CustodyNewTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CustodyNew.Contract.Withdraw(&_ERC20CustodyNew.TransactOpts, token, to, amount)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x21fc65f2.
//
// Solidity: function withdrawAndCall(address token, address to, uint256 amount, bytes data) returns()
func (_ERC20CustodyNew *ERC20CustodyNewTransactor) WithdrawAndCall(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20CustodyNew.contract.Transact(opts, "withdrawAndCall", token, to, amount, data)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x21fc65f2.
//
// Solidity: function withdrawAndCall(address token, address to, uint256 amount, bytes data) returns()
func (_ERC20CustodyNew *ERC20CustodyNewSession) WithdrawAndCall(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20CustodyNew.Contract.WithdrawAndCall(&_ERC20CustodyNew.TransactOpts, token, to, amount, data)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x21fc65f2.
//
// Solidity: function withdrawAndCall(address token, address to, uint256 amount, bytes data) returns()
func (_ERC20CustodyNew *ERC20CustodyNewTransactorSession) WithdrawAndCall(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20CustodyNew.Contract.WithdrawAndCall(&_ERC20CustodyNew.TransactOpts, token, to, amount, data)
}

// ERC20CustodyNewWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ERC20CustodyNew contract.
type ERC20CustodyNewWithdrawIterator struct {
	Event *ERC20CustodyNewWithdraw // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyNewWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyNewWithdraw)
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
		it.Event = new(ERC20CustodyNewWithdraw)
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
func (it *ERC20CustodyNewWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyNewWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyNewWithdraw represents a Withdraw event raised by the ERC20CustodyNew contract.
type ERC20CustodyNewWithdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_ERC20CustodyNew *ERC20CustodyNewFilterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*ERC20CustodyNewWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20CustodyNew.contract.FilterLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyNewWithdrawIterator{contract: _ERC20CustodyNew.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_ERC20CustodyNew *ERC20CustodyNewFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ERC20CustodyNewWithdraw, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20CustodyNew.contract.WatchLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyNewWithdraw)
				if err := _ERC20CustodyNew.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_ERC20CustodyNew *ERC20CustodyNewFilterer) ParseWithdraw(log types.Log) (*ERC20CustodyNewWithdraw, error) {
	event := new(ERC20CustodyNewWithdraw)
	if err := _ERC20CustodyNew.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyNewWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the ERC20CustodyNew contract.
type ERC20CustodyNewWithdrawAndCallIterator struct {
	Event *ERC20CustodyNewWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyNewWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyNewWithdrawAndCall)
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
		it.Event = new(ERC20CustodyNewWithdrawAndCall)
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
func (it *ERC20CustodyNewWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyNewWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyNewWithdrawAndCall represents a WithdrawAndCall event raised by the ERC20CustodyNew contract.
type ERC20CustodyNewWithdrawAndCall struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e.
//
// Solidity: event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ERC20CustodyNew *ERC20CustodyNewFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*ERC20CustodyNewWithdrawAndCallIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20CustodyNew.contract.FilterLogs(opts, "WithdrawAndCall", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyNewWithdrawAndCallIterator{contract: _ERC20CustodyNew.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e.
//
// Solidity: event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ERC20CustodyNew *ERC20CustodyNewFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *ERC20CustodyNewWithdrawAndCall, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20CustodyNew.contract.WatchLogs(opts, "WithdrawAndCall", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyNewWithdrawAndCall)
				if err := _ERC20CustodyNew.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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

// ParseWithdrawAndCall is a log parse operation binding the contract event 0x85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e.
//
// Solidity: event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ERC20CustodyNew *ERC20CustodyNewFilterer) ParseWithdrawAndCall(log types.Log) (*ERC20CustodyNewWithdrawAndCall, error) {
	event := new(ERC20CustodyNewWithdrawAndCall)
	if err := _ERC20CustodyNew.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}