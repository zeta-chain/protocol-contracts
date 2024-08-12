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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_gateway\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"callReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawAndCallReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600f57600080fd5b50604051610908380380610908833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b61087b8061008d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b61461005b578063865b36f6146100a4575b600080fd5b61005961005436600461052d565b6100b7565b005b60005461007b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100596100b23660046105e8565b6102e5565b60008383836040516024016100ce939291906106dd565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905260005490517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526024810189905291925086169063095ea7b3906044016020604051808303816000875af11580156101be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101e29190610707565b610218576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160808101825261032180825260016020808401919091528284019190915282519081018352600080825260608301919091525491517f939bc896000000000000000000000000000000000000000000000000000000008152909173ffffffffffffffffffffffffffffffffffffffff169063939bc896906102a9908b908b908b9088908890600401610794565b600060405180830381600087803b1580156102c357600080fd5b505af11580156102d7573d6000803e3d6000fd5b505050505050505050505050565b60008383836040516024016102fc939291906106dd565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f97000000000000000000000000000000000000000000000000000000001790528151608081018352610321808252600182840152818401528251918201835260008083526060820192909252905491517f845b96ce0000000000000000000000000000000000000000000000000000000081529293509173ffffffffffffffffffffffffffffffffffffffff919091169063845b96ce906103fe908a908a90879087906004016107fb565b600060405180830381600087803b15801561041857600080fd5b505af115801561042c573d6000803e3d6000fd5b5050505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261047957600080fd5b81356020830160008067ffffffffffffffff84111561049a5761049a610439565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156104e7576104e7610439565b6040528381529050808284018710156104ff57600080fd5b838360208301376000602085830101528094505050505092915050565b801515811461052a57600080fd5b50565b60008060008060008060c0878903121561054657600080fd5b863567ffffffffffffffff81111561055d57600080fd5b61056989828a01610468565b96505060208701359450604087013573ffffffffffffffffffffffffffffffffffffffff8116811461059a57600080fd5b9350606087013567ffffffffffffffff8111156105b657600080fd5b6105c289828a01610468565b9350506080870135915060a08701356105da8161051c565b809150509295509295509295565b600080600080600060a0868803121561060057600080fd5b853567ffffffffffffffff81111561061757600080fd5b61062388828901610468565b95505060208601359350604086013567ffffffffffffffff81111561064757600080fd5b61065388828901610468565b93505060608601359150608086013561066b8161051c565b809150509295509295909350565b6000815180845260005b8181101561069f57602081850181015186830182015201610683565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6060815260006106f06060830186610679565b602083019490945250901515604090910152919050565b60006020828403121561071957600080fd5b81516107248161051c565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff6040820151166040830152600060608201516080606085015261078c6080850182610679565b949350505050565b60a0815260006107a760a0830188610679565b86602084015273ffffffffffffffffffffffffffffffffffffffff8616604084015282810360608401526107db8186610679565b905082810360808401526107ef818561072b565b98975050505050505050565b60808152600061080e6080830187610679565b85602084015282810360408401526108268186610679565b9050828103606084015261083a818561072b565b97965050505050505056fea26469706673582212207d8d7ec2c603d0efed6510211e9e75124bf0708276bf3c867c0f7fe804766a5a64736f6c634300081a0033",
}

// SenderZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use SenderZEVMMetaData.ABI instead.
var SenderZEVMABI = SenderZEVMMetaData.ABI

// SenderZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SenderZEVMMetaData.Bin instead.
var SenderZEVMBin = SenderZEVMMetaData.Bin

// DeploySenderZEVM deploys a new Ethereum contract, binding an instance of SenderZEVM to it.
func DeploySenderZEVM(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address) (common.Address, *types.Transaction, *SenderZEVM, error) {
	parsed, err := SenderZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenderZEVMBin), backend, _gateway)
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

// CallReceiver is a paid mutator transaction binding the contract method 0x865b36f6.
//
// Solidity: function callReceiver(bytes receiver, uint256 chainId, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMTransactor) CallReceiver(opts *bind.TransactOpts, receiver []byte, chainId *big.Int, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.contract.Transact(opts, "callReceiver", receiver, chainId, str, num, flag)
}

// CallReceiver is a paid mutator transaction binding the contract method 0x865b36f6.
//
// Solidity: function callReceiver(bytes receiver, uint256 chainId, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMSession) CallReceiver(receiver []byte, chainId *big.Int, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.Contract.CallReceiver(&_SenderZEVM.TransactOpts, receiver, chainId, str, num, flag)
}

// CallReceiver is a paid mutator transaction binding the contract method 0x865b36f6.
//
// Solidity: function callReceiver(bytes receiver, uint256 chainId, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMTransactorSession) CallReceiver(receiver []byte, chainId *big.Int, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.Contract.CallReceiver(&_SenderZEVM.TransactOpts, receiver, chainId, str, num, flag)
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
