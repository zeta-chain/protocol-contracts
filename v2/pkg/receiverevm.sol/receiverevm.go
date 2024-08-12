// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package receiverevm

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
	Asset         common.Address
	Amount        uint64
	RevertMessage []byte
}

// ReceiverEVMMetaData contains all meta data concerning the ReceiverEVM contract.
var ReceiverEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"onRevert\",\"inputs\":[{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveERC20\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveERC20Partial\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveNoParams\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveNonPayable\",\"inputs\":[{\"name\":\"strs\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receivePayable\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600f57600080fd5b506001600055610dbc806100246000396000f3fe6080604052600436106100635760003560e01c8063c513169111610040578063c5131691146100c1578063e04d4f97146100e1578063f05b6abf146100f457005b8063357fc5a21461006c578063660b9de01461008c5780636ed70169146100ac57005b3661006a57005b005b34801561007857600080fd5b5061006a6100873660046106bd565b610114565b34801561009857600080fd5b5061006a6100a73660046106f9565b6101aa565b3480156100b857600080fd5b5061006a6101e6565b3480156100cd57600080fd5b5061006a6100dc3660046106bd565b61021b565b61006a6100ef366004610859565b6102f6565b34801561010057600080fd5b5061006a61010f366004610945565b61033a565b61011c61036f565b61013e73ffffffffffffffffffffffffffffffffffffffff83163383866103b2565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101a56001600055565b505050565b7f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e33826040516101db929190610a78565b60405180910390a150565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61022361036f565b6000610230600285610b57565b90508060000361026c576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61028e73ffffffffffffffffffffffffffffffffffffffff84163384846103b2565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101a56001600055565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa333485858560405161032d959493929190610c00565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b1463384848460405161032d9493929190610c8a565b6002600054036103ab576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261044790859061044d565b50505050565b600061046f73ffffffffffffffffffffffffffffffffffffffff8416836104e8565b905080516000141580156104945750808060200190518101906104929190610d4d565b155b156101a5576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b60606104f6838360006104fd565b9392505050565b60608147101561053b576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016104df565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105649190610d6a565b60006040518083038185875af1925050503d80600081146105a1576040519150601f19603f3d011682016040523d82523d6000602084013e6105a6565b606091505b50915091506105b68683836105c0565b9695505050505050565b6060826105d5576105d08261064f565b6104f6565b81511580156105f9575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610648576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016104df565b50806104f6565b80511561065f5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff811681146106b857600080fd5b919050565b6000806000606084860312156106d257600080fd5b833592506106e260208501610694565b91506106f060408501610694565b90509250925092565b60006020828403121561070b57600080fd5b813567ffffffffffffffff81111561072257600080fd5b8201606081850312156104f657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156107aa576107aa610734565b604052919050565b600082601f8301126107c357600080fd5b813567ffffffffffffffff8111156107dd576107dd610734565b61080e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610763565b81815284602083860101111561082357600080fd5b816020850160208301376000918101602001919091529392505050565b801515811461069157600080fd5b80356106b881610840565b60008060006060848603121561086e57600080fd5b833567ffffffffffffffff81111561088557600080fd5b610891868287016107b2565b9350506020840135915060408401356108a981610840565b809150509250925092565b600067ffffffffffffffff8211156108ce576108ce610734565b5060051b60200190565b600082601f8301126108e957600080fd5b81356108fc6108f7826108b4565b610763565b8082825260208201915060208360051b86010192508583111561091e57600080fd5b602085015b8381101561093b578035835260209283019201610923565b5095945050505050565b60008060006060848603121561095a57600080fd5b833567ffffffffffffffff81111561097157600080fd5b8401601f8101861361098257600080fd5b80356109906108f7826108b4565b8082825260208201915060208360051b8501019250888311156109b257600080fd5b602084015b838110156109f457803567ffffffffffffffff8111156109d657600080fd5b6109e58b6020838901016107b2565b845250602092830192016109b7565b509550505050602084013567ffffffffffffffff811115610a1457600080fd5b610a20868287016108d8565b9250506106f06040850161084e565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610ab683610694565b1660408201526000602083013567ffffffffffffffff8116808214610ada57600080fd5b6060840152506040830135368490037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112610b1657600080fd5b830160208101903567ffffffffffffffff811115610b3357600080fd5b803603821315610b4257600080fd5b606060808501526105b660a085018284610a2f565b600082610b8d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60005b83811015610bad578181015183820152602001610b95565b50506000910152565b60008151808452610bce816020860160208601610b92565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610c3560a0830186610bb6565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610c80578151865260209586019590910190600101610c62565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610d1d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610d08858351610bb6565b94506020938401939190910190600101610cce565b505050508281036040840152610d338186610c4e565b915050610d44606083018415159052565b95945050505050565b600060208284031215610d5f57600080fd5b81516104f681610840565b60008251610d7c818460208701610b92565b919091019291505056fea2646970667358221220d72588ea8acbf56155273edae233d1fe939bca804f44cf5204c615539f68e9ae64736f6c634300081a0033",
}

// ReceiverEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiverEVMMetaData.ABI instead.
var ReceiverEVMABI = ReceiverEVMMetaData.ABI

// ReceiverEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReceiverEVMMetaData.Bin instead.
var ReceiverEVMBin = ReceiverEVMMetaData.Bin

// DeployReceiverEVM deploys a new Ethereum contract, binding an instance of ReceiverEVM to it.
func DeployReceiverEVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReceiverEVM, error) {
	parsed, err := ReceiverEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiverEVMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReceiverEVM{ReceiverEVMCaller: ReceiverEVMCaller{contract: contract}, ReceiverEVMTransactor: ReceiverEVMTransactor{contract: contract}, ReceiverEVMFilterer: ReceiverEVMFilterer{contract: contract}}, nil
}

// ReceiverEVM is an auto generated Go binding around an Ethereum contract.
type ReceiverEVM struct {
	ReceiverEVMCaller     // Read-only binding to the contract
	ReceiverEVMTransactor // Write-only binding to the contract
	ReceiverEVMFilterer   // Log filterer for contract events
}

// ReceiverEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiverEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiverEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiverEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiverEVMSession struct {
	Contract     *ReceiverEVM      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiverEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiverEVMCallerSession struct {
	Contract *ReceiverEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ReceiverEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiverEVMTransactorSession struct {
	Contract     *ReceiverEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ReceiverEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiverEVMRaw struct {
	Contract *ReceiverEVM // Generic contract binding to access the raw methods on
}

// ReceiverEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiverEVMCallerRaw struct {
	Contract *ReceiverEVMCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiverEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiverEVMTransactorRaw struct {
	Contract *ReceiverEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiverEVM creates a new instance of ReceiverEVM, bound to a specific deployed contract.
func NewReceiverEVM(address common.Address, backend bind.ContractBackend) (*ReceiverEVM, error) {
	contract, err := bindReceiverEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiverEVM{ReceiverEVMCaller: ReceiverEVMCaller{contract: contract}, ReceiverEVMTransactor: ReceiverEVMTransactor{contract: contract}, ReceiverEVMFilterer: ReceiverEVMFilterer{contract: contract}}, nil
}

// NewReceiverEVMCaller creates a new read-only instance of ReceiverEVM, bound to a specific deployed contract.
func NewReceiverEVMCaller(address common.Address, caller bind.ContractCaller) (*ReceiverEVMCaller, error) {
	contract, err := bindReceiverEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverEVMCaller{contract: contract}, nil
}

// NewReceiverEVMTransactor creates a new write-only instance of ReceiverEVM, bound to a specific deployed contract.
func NewReceiverEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiverEVMTransactor, error) {
	contract, err := bindReceiverEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverEVMTransactor{contract: contract}, nil
}

// NewReceiverEVMFilterer creates a new log filterer instance of ReceiverEVM, bound to a specific deployed contract.
func NewReceiverEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiverEVMFilterer, error) {
	contract, err := bindReceiverEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiverEVMFilterer{contract: contract}, nil
}

// bindReceiverEVM binds a generic wrapper to an already deployed contract.
func bindReceiverEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReceiverEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiverEVM *ReceiverEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiverEVM.Contract.ReceiverEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiverEVM *ReceiverEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.ReceiverEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiverEVM *ReceiverEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.ReceiverEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiverEVM *ReceiverEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiverEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiverEVM *ReceiverEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiverEVM *ReceiverEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.contract.Transact(opts, method, params...)
}

// OnRevert is a paid mutator transaction binding the contract method 0x660b9de0.
//
// Solidity: function onRevert((address,uint64,bytes) revertContext) returns()
func (_ReceiverEVM *ReceiverEVMTransactor) OnRevert(opts *bind.TransactOpts, revertContext RevertContext) (*types.Transaction, error) {
	return _ReceiverEVM.contract.Transact(opts, "onRevert", revertContext)
}

// OnRevert is a paid mutator transaction binding the contract method 0x660b9de0.
//
// Solidity: function onRevert((address,uint64,bytes) revertContext) returns()
func (_ReceiverEVM *ReceiverEVMSession) OnRevert(revertContext RevertContext) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.OnRevert(&_ReceiverEVM.TransactOpts, revertContext)
}

// OnRevert is a paid mutator transaction binding the contract method 0x660b9de0.
//
// Solidity: function onRevert((address,uint64,bytes) revertContext) returns()
func (_ReceiverEVM *ReceiverEVMTransactorSession) OnRevert(revertContext RevertContext) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.OnRevert(&_ReceiverEVM.TransactOpts, revertContext)
}

// ReceiveERC20 is a paid mutator transaction binding the contract method 0x357fc5a2.
//
// Solidity: function receiveERC20(uint256 amount, address token, address destination) returns()
func (_ReceiverEVM *ReceiverEVMTransactor) ReceiveERC20(opts *bind.TransactOpts, amount *big.Int, token common.Address, destination common.Address) (*types.Transaction, error) {
	return _ReceiverEVM.contract.Transact(opts, "receiveERC20", amount, token, destination)
}

// ReceiveERC20 is a paid mutator transaction binding the contract method 0x357fc5a2.
//
// Solidity: function receiveERC20(uint256 amount, address token, address destination) returns()
func (_ReceiverEVM *ReceiverEVMSession) ReceiveERC20(amount *big.Int, token common.Address, destination common.Address) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.ReceiveERC20(&_ReceiverEVM.TransactOpts, amount, token, destination)
}

// ReceiveERC20 is a paid mutator transaction binding the contract method 0x357fc5a2.
//
// Solidity: function receiveERC20(uint256 amount, address token, address destination) returns()
func (_ReceiverEVM *ReceiverEVMTransactorSession) ReceiveERC20(amount *big.Int, token common.Address, destination common.Address) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.ReceiveERC20(&_ReceiverEVM.TransactOpts, amount, token, destination)
}

// ReceiveERC20Partial is a paid mutator transaction binding the contract method 0xc5131691.
//
// Solidity: function receiveERC20Partial(uint256 amount, address token, address destination) returns()
func (_ReceiverEVM *ReceiverEVMTransactor) ReceiveERC20Partial(opts *bind.TransactOpts, amount *big.Int, token common.Address, destination common.Address) (*types.Transaction, error) {
	return _ReceiverEVM.contract.Transact(opts, "receiveERC20Partial", amount, token, destination)
}

// ReceiveERC20Partial is a paid mutator transaction binding the contract method 0xc5131691.
//
// Solidity: function receiveERC20Partial(uint256 amount, address token, address destination) returns()
func (_ReceiverEVM *ReceiverEVMSession) ReceiveERC20Partial(amount *big.Int, token common.Address, destination common.Address) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.ReceiveERC20Partial(&_ReceiverEVM.TransactOpts, amount, token, destination)
}

// ReceiveERC20Partial is a paid mutator transaction binding the contract method 0xc5131691.
//
// Solidity: function receiveERC20Partial(uint256 amount, address token, address destination) returns()
func (_ReceiverEVM *ReceiverEVMTransactorSession) ReceiveERC20Partial(amount *big.Int, token common.Address, destination common.Address) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.ReceiveERC20Partial(&_ReceiverEVM.TransactOpts, amount, token, destination)
}

// ReceiveNoParams is a paid mutator transaction binding the contract method 0x6ed70169.
//
// Solidity: function receiveNoParams() returns()
func (_ReceiverEVM *ReceiverEVMTransactor) ReceiveNoParams(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverEVM.contract.Transact(opts, "receiveNoParams")
}

// ReceiveNoParams is a paid mutator transaction binding the contract method 0x6ed70169.
//
// Solidity: function receiveNoParams() returns()
func (_ReceiverEVM *ReceiverEVMSession) ReceiveNoParams() (*types.Transaction, error) {
	return _ReceiverEVM.Contract.ReceiveNoParams(&_ReceiverEVM.TransactOpts)
}

// ReceiveNoParams is a paid mutator transaction binding the contract method 0x6ed70169.
//
// Solidity: function receiveNoParams() returns()
func (_ReceiverEVM *ReceiverEVMTransactorSession) ReceiveNoParams() (*types.Transaction, error) {
	return _ReceiverEVM.Contract.ReceiveNoParams(&_ReceiverEVM.TransactOpts)
}

// ReceiveNonPayable is a paid mutator transaction binding the contract method 0xf05b6abf.
//
// Solidity: function receiveNonPayable(string[] strs, uint256[] nums, bool flag) returns()
func (_ReceiverEVM *ReceiverEVMTransactor) ReceiveNonPayable(opts *bind.TransactOpts, strs []string, nums []*big.Int, flag bool) (*types.Transaction, error) {
	return _ReceiverEVM.contract.Transact(opts, "receiveNonPayable", strs, nums, flag)
}

// ReceiveNonPayable is a paid mutator transaction binding the contract method 0xf05b6abf.
//
// Solidity: function receiveNonPayable(string[] strs, uint256[] nums, bool flag) returns()
func (_ReceiverEVM *ReceiverEVMSession) ReceiveNonPayable(strs []string, nums []*big.Int, flag bool) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.ReceiveNonPayable(&_ReceiverEVM.TransactOpts, strs, nums, flag)
}

// ReceiveNonPayable is a paid mutator transaction binding the contract method 0xf05b6abf.
//
// Solidity: function receiveNonPayable(string[] strs, uint256[] nums, bool flag) returns()
func (_ReceiverEVM *ReceiverEVMTransactorSession) ReceiveNonPayable(strs []string, nums []*big.Int, flag bool) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.ReceiveNonPayable(&_ReceiverEVM.TransactOpts, strs, nums, flag)
}

// ReceivePayable is a paid mutator transaction binding the contract method 0xe04d4f97.
//
// Solidity: function receivePayable(string str, uint256 num, bool flag) payable returns()
func (_ReceiverEVM *ReceiverEVMTransactor) ReceivePayable(opts *bind.TransactOpts, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _ReceiverEVM.contract.Transact(opts, "receivePayable", str, num, flag)
}

// ReceivePayable is a paid mutator transaction binding the contract method 0xe04d4f97.
//
// Solidity: function receivePayable(string str, uint256 num, bool flag) payable returns()
func (_ReceiverEVM *ReceiverEVMSession) ReceivePayable(str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.ReceivePayable(&_ReceiverEVM.TransactOpts, str, num, flag)
}

// ReceivePayable is a paid mutator transaction binding the contract method 0xe04d4f97.
//
// Solidity: function receivePayable(string str, uint256 num, bool flag) payable returns()
func (_ReceiverEVM *ReceiverEVMTransactorSession) ReceivePayable(str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.ReceivePayable(&_ReceiverEVM.TransactOpts, str, num, flag)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ReceiverEVM *ReceiverEVMTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ReceiverEVM.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ReceiverEVM *ReceiverEVMSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.Fallback(&_ReceiverEVM.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ReceiverEVM *ReceiverEVMTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ReceiverEVM.Contract.Fallback(&_ReceiverEVM.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ReceiverEVM *ReceiverEVMTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverEVM.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ReceiverEVM *ReceiverEVMSession) Receive() (*types.Transaction, error) {
	return _ReceiverEVM.Contract.Receive(&_ReceiverEVM.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ReceiverEVM *ReceiverEVMTransactorSession) Receive() (*types.Transaction, error) {
	return _ReceiverEVM.Contract.Receive(&_ReceiverEVM.TransactOpts)
}

// ReceiverEVMReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the ReceiverEVM contract.
type ReceiverEVMReceivedERC20Iterator struct {
	Event *ReceiverEVMReceivedERC20 // Event containing the contract specifics and raw log

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
func (it *ReceiverEVMReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiverEVMReceivedERC20)
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
		it.Event = new(ReceiverEVMReceivedERC20)
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
func (it *ReceiverEVMReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiverEVMReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiverEVMReceivedERC20 represents a ReceivedERC20 event raised by the ReceiverEVM contract.
type ReceiverEVMReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_ReceiverEVM *ReceiverEVMFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*ReceiverEVMReceivedERC20Iterator, error) {

	logs, sub, err := _ReceiverEVM.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &ReceiverEVMReceivedERC20Iterator{contract: _ReceiverEVM.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_ReceiverEVM *ReceiverEVMFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *ReceiverEVMReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _ReceiverEVM.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiverEVMReceivedERC20)
				if err := _ReceiverEVM.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
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

// ParseReceivedERC20 is a log parse operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_ReceiverEVM *ReceiverEVMFilterer) ParseReceivedERC20(log types.Log) (*ReceiverEVMReceivedERC20, error) {
	event := new(ReceiverEVMReceivedERC20)
	if err := _ReceiverEVM.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiverEVMReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the ReceiverEVM contract.
type ReceiverEVMReceivedNoParamsIterator struct {
	Event *ReceiverEVMReceivedNoParams // Event containing the contract specifics and raw log

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
func (it *ReceiverEVMReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiverEVMReceivedNoParams)
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
		it.Event = new(ReceiverEVMReceivedNoParams)
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
func (it *ReceiverEVMReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiverEVMReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiverEVMReceivedNoParams represents a ReceivedNoParams event raised by the ReceiverEVM contract.
type ReceiverEVMReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_ReceiverEVM *ReceiverEVMFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*ReceiverEVMReceivedNoParamsIterator, error) {

	logs, sub, err := _ReceiverEVM.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &ReceiverEVMReceivedNoParamsIterator{contract: _ReceiverEVM.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_ReceiverEVM *ReceiverEVMFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *ReceiverEVMReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _ReceiverEVM.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiverEVMReceivedNoParams)
				if err := _ReceiverEVM.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
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

// ParseReceivedNoParams is a log parse operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_ReceiverEVM *ReceiverEVMFilterer) ParseReceivedNoParams(log types.Log) (*ReceiverEVMReceivedNoParams, error) {
	event := new(ReceiverEVMReceivedNoParams)
	if err := _ReceiverEVM.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiverEVMReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the ReceiverEVM contract.
type ReceiverEVMReceivedNonPayableIterator struct {
	Event *ReceiverEVMReceivedNonPayable // Event containing the contract specifics and raw log

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
func (it *ReceiverEVMReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiverEVMReceivedNonPayable)
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
		it.Event = new(ReceiverEVMReceivedNonPayable)
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
func (it *ReceiverEVMReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiverEVMReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiverEVMReceivedNonPayable represents a ReceivedNonPayable event raised by the ReceiverEVM contract.
type ReceiverEVMReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_ReceiverEVM *ReceiverEVMFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*ReceiverEVMReceivedNonPayableIterator, error) {

	logs, sub, err := _ReceiverEVM.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &ReceiverEVMReceivedNonPayableIterator{contract: _ReceiverEVM.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_ReceiverEVM *ReceiverEVMFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *ReceiverEVMReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _ReceiverEVM.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiverEVMReceivedNonPayable)
				if err := _ReceiverEVM.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
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

// ParseReceivedNonPayable is a log parse operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_ReceiverEVM *ReceiverEVMFilterer) ParseReceivedNonPayable(log types.Log) (*ReceiverEVMReceivedNonPayable, error) {
	event := new(ReceiverEVMReceivedNonPayable)
	if err := _ReceiverEVM.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiverEVMReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the ReceiverEVM contract.
type ReceiverEVMReceivedPayableIterator struct {
	Event *ReceiverEVMReceivedPayable // Event containing the contract specifics and raw log

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
func (it *ReceiverEVMReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiverEVMReceivedPayable)
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
		it.Event = new(ReceiverEVMReceivedPayable)
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
func (it *ReceiverEVMReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiverEVMReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiverEVMReceivedPayable represents a ReceivedPayable event raised by the ReceiverEVM contract.
type ReceiverEVMReceivedPayable struct {
	Sender common.Address
	Value  *big.Int
	Str    string
	Num    *big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedPayable is a free log retrieval operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_ReceiverEVM *ReceiverEVMFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*ReceiverEVMReceivedPayableIterator, error) {

	logs, sub, err := _ReceiverEVM.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &ReceiverEVMReceivedPayableIterator{contract: _ReceiverEVM.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_ReceiverEVM *ReceiverEVMFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *ReceiverEVMReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _ReceiverEVM.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiverEVMReceivedPayable)
				if err := _ReceiverEVM.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
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

// ParseReceivedPayable is a log parse operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_ReceiverEVM *ReceiverEVMFilterer) ParseReceivedPayable(log types.Log) (*ReceiverEVMReceivedPayable, error) {
	event := new(ReceiverEVMReceivedPayable)
	if err := _ReceiverEVM.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiverEVMReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the ReceiverEVM contract.
type ReceiverEVMReceivedRevertIterator struct {
	Event *ReceiverEVMReceivedRevert // Event containing the contract specifics and raw log

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
func (it *ReceiverEVMReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiverEVMReceivedRevert)
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
		it.Event = new(ReceiverEVMReceivedRevert)
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
func (it *ReceiverEVMReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiverEVMReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiverEVMReceivedRevert represents a ReceivedRevert event raised by the ReceiverEVM contract.
type ReceiverEVMReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_ReceiverEVM *ReceiverEVMFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*ReceiverEVMReceivedRevertIterator, error) {

	logs, sub, err := _ReceiverEVM.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &ReceiverEVMReceivedRevertIterator{contract: _ReceiverEVM.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_ReceiverEVM *ReceiverEVMFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *ReceiverEVMReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _ReceiverEVM.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiverEVMReceivedRevert)
				if err := _ReceiverEVM.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
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

// ParseReceivedRevert is a log parse operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_ReceiverEVM *ReceiverEVMFilterer) ParseReceivedRevert(log types.Log) (*ReceiverEVMReceivedRevert, error) {
	event := new(ReceiverEVMReceivedRevert)
	if err := _ReceiverEVM.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
