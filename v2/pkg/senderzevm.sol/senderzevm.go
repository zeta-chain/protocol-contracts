// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package senderzevm

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

// SenderZEVMMetaData contains all meta data concerning the SenderZEVM contract.
var SenderZEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"gateway_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"callReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawAndCallReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600f57600080fd5b50604051610a12380380610a12833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b6109858061008d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b61461005b5780637a34d8bb146100a4575b600080fd5b6100596100543660046105e7565b6100b7565b005b60005461007b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100596100b2366004610687565b61030c565b60008383836040516024016100ce93929190610783565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905260005490915073ffffffffffffffffffffffffffffffffffffffff8087169163095ea7b391166101758960026107ad565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af11580156101e5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020991906107ed565b61023f576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160808101825261032180825260016020808401829052838501929092528351918201845260008083526060840192909252905492517f90ad3e23000000000000000000000000000000000000000000000000000000008152919273ffffffffffffffffffffffffffffffffffffffff16916390ad3e23916102d0918c918c918c918991899060040161087a565b600060405180830381600087803b1580156102ea57600080fd5b505af11580156102fe573d6000803e3d6000fd5b505050505050505050505050565b600083838360405160240161032393929190610783565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905281516080810183526103218082526001828401819052828501919091528351928301845260008084526060830193909352915492517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff93841660048201526024810183905293945092909188169063095ea7b3906044016020604051808303816000875af1158015610447573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061046b91906107ed565b506000546040517fdc9ca2e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063dc9ca2e7906102d0908b908b908890879089906004016108e8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261050a57600080fd5b81356020830160008067ffffffffffffffff84111561052b5761052b6104ca565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715610578576105786104ca565b60405283815290508082840187101561059057600080fd5b838360208301376000602085830101528094505050505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146105d157600080fd5b919050565b80151581146105e457600080fd5b50565b60008060008060008060c0878903121561060057600080fd5b863567ffffffffffffffff81111561061757600080fd5b61062389828a016104f9565b96505060208701359450610639604088016105ad565b9350606087013567ffffffffffffffff81111561065557600080fd5b61066189828a016104f9565b9350506080870135915060a0870135610679816105d6565b809150509295509295509295565b600080600080600060a0868803121561069f57600080fd5b853567ffffffffffffffff8111156106b657600080fd5b6106c2888289016104f9565b9550506106d1602087016105ad565b9350604086013567ffffffffffffffff8111156106ed57600080fd5b6106f9888289016104f9565b935050606086013591506080860135610711816105d6565b809150509295509295909350565b6000815180845260005b8181101561074557602081850181015186830182015201610729565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b606081526000610796606083018661071f565b602083019490945250901515604090910152919050565b808201808211156107e7577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b6000602082840312156107ff57600080fd5b815161080a816105d6565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff60408201511660408301526000606082015160806060850152610872608085018261071f565b949350505050565b60c08152600061088d60c083018961071f565b87602084015273ffffffffffffffffffffffffffffffffffffffff8716604084015282810360608401526108c1818761071f565b905084608084015282810360a08401526108db8185610811565b9998505050505050505050565b60a0815260006108fb60a083018861071f565b73ffffffffffffffffffffffffffffffffffffffff871660208401528281036040840152610929818761071f565b905084606084015282810360808401526109438185610811565b9897505050505050505056fea264697066735822122081a08bbd8f6b6c1e3e73e92df5ac001d9e383817e2faa175b98d3c6cd01edab664736f6c634300081a0033",
}

// SenderZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use SenderZEVMMetaData.ABI instead.
var SenderZEVMABI = SenderZEVMMetaData.ABI

// SenderZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SenderZEVMMetaData.Bin instead.
var SenderZEVMBin = SenderZEVMMetaData.Bin

// DeploySenderZEVM deploys a new Ethereum contract, binding an instance of SenderZEVM to it.
func DeploySenderZEVM(auth *bind.TransactOpts, backend bind.ContractBackend, gateway_ common.Address) (common.Address, *types.Transaction, *SenderZEVM, error) {
	parsed, err := SenderZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenderZEVMBin), backend, gateway_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SenderZEVM{SenderZEVMCaller: SenderZEVMCaller{contract: contract}, SenderZEVMTransactor: SenderZEVMTransactor{contract: contract}, SenderZEVMFilterer: SenderZEVMFilterer{contract: contract}}, nil
}

// SenderZEVM is an auto generated Go binding around an Ethereum contract.
type SenderZEVM struct {
	SenderZEVMCaller     // Read-only binding to the contract
	SenderZEVMTransactor // Write-only binding to the contract
	SenderZEVMFilterer   // Log filterer for contract events
}

// SenderZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenderZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenderZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenderZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenderZEVMSession struct {
	Contract     *SenderZEVM       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenderZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenderZEVMCallerSession struct {
	Contract *SenderZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SenderZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenderZEVMTransactorSession struct {
	Contract     *SenderZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SenderZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenderZEVMRaw struct {
	Contract *SenderZEVM // Generic contract binding to access the raw methods on
}

// SenderZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenderZEVMCallerRaw struct {
	Contract *SenderZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// SenderZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenderZEVMTransactorRaw struct {
	Contract *SenderZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSenderZEVM creates a new instance of SenderZEVM, bound to a specific deployed contract.
func NewSenderZEVM(address common.Address, backend bind.ContractBackend) (*SenderZEVM, error) {
	contract, err := bindSenderZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SenderZEVM{SenderZEVMCaller: SenderZEVMCaller{contract: contract}, SenderZEVMTransactor: SenderZEVMTransactor{contract: contract}, SenderZEVMFilterer: SenderZEVMFilterer{contract: contract}}, nil
}

// NewSenderZEVMCaller creates a new read-only instance of SenderZEVM, bound to a specific deployed contract.
func NewSenderZEVMCaller(address common.Address, caller bind.ContractCaller) (*SenderZEVMCaller, error) {
	contract, err := bindSenderZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenderZEVMCaller{contract: contract}, nil
}

// NewSenderZEVMTransactor creates a new write-only instance of SenderZEVM, bound to a specific deployed contract.
func NewSenderZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*SenderZEVMTransactor, error) {
	contract, err := bindSenderZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenderZEVMTransactor{contract: contract}, nil
}

// NewSenderZEVMFilterer creates a new log filterer instance of SenderZEVM, bound to a specific deployed contract.
func NewSenderZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*SenderZEVMFilterer, error) {
	contract, err := bindSenderZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenderZEVMFilterer{contract: contract}, nil
}

// bindSenderZEVM binds a generic wrapper to an already deployed contract.
func bindSenderZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SenderZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenderZEVM *SenderZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenderZEVM.Contract.SenderZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenderZEVM *SenderZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenderZEVM.Contract.SenderZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenderZEVM *SenderZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenderZEVM.Contract.SenderZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenderZEVM *SenderZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenderZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenderZEVM *SenderZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenderZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenderZEVM *SenderZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenderZEVM.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_SenderZEVM *SenderZEVMCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SenderZEVM.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_SenderZEVM *SenderZEVMSession) Gateway() (common.Address, error) {
	return _SenderZEVM.Contract.Gateway(&_SenderZEVM.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_SenderZEVM *SenderZEVMCallerSession) Gateway() (common.Address, error) {
	return _SenderZEVM.Contract.Gateway(&_SenderZEVM.CallOpts)
}

// CallReceiver is a paid mutator transaction binding the contract method 0x7a34d8bb.
//
// Solidity: function callReceiver(bytes receiver, address zrc20, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMTransactor) CallReceiver(opts *bind.TransactOpts, receiver []byte, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.contract.Transact(opts, "callReceiver", receiver, zrc20, str, num, flag)
}

// CallReceiver is a paid mutator transaction binding the contract method 0x7a34d8bb.
//
// Solidity: function callReceiver(bytes receiver, address zrc20, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMSession) CallReceiver(receiver []byte, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.Contract.CallReceiver(&_SenderZEVM.TransactOpts, receiver, zrc20, str, num, flag)
}

// CallReceiver is a paid mutator transaction binding the contract method 0x7a34d8bb.
//
// Solidity: function callReceiver(bytes receiver, address zrc20, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMTransactorSession) CallReceiver(receiver []byte, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.Contract.CallReceiver(&_SenderZEVM.TransactOpts, receiver, zrc20, str, num, flag)
}

// WithdrawAndCallReceiver is a paid mutator transaction binding the contract method 0x0abd8905.
//
// Solidity: function withdrawAndCallReceiver(bytes receiver, uint256 amount, address zrc20, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMTransactor) WithdrawAndCallReceiver(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.contract.Transact(opts, "withdrawAndCallReceiver", receiver, amount, zrc20, str, num, flag)
}

// WithdrawAndCallReceiver is a paid mutator transaction binding the contract method 0x0abd8905.
//
// Solidity: function withdrawAndCallReceiver(bytes receiver, uint256 amount, address zrc20, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMSession) WithdrawAndCallReceiver(receiver []byte, amount *big.Int, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.Contract.WithdrawAndCallReceiver(&_SenderZEVM.TransactOpts, receiver, amount, zrc20, str, num, flag)
}

// WithdrawAndCallReceiver is a paid mutator transaction binding the contract method 0x0abd8905.
//
// Solidity: function withdrawAndCallReceiver(bytes receiver, uint256 amount, address zrc20, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMTransactorSession) WithdrawAndCallReceiver(receiver []byte, amount *big.Int, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.Contract.WithdrawAndCallReceiver(&_SenderZEVM.TransactOpts, receiver, amount, zrc20, str, num, flag)
}
