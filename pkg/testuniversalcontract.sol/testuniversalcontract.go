// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testuniversalcontract

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

// AbortContext is an auto generated low-level Go binding around an user-defined struct.
type AbortContext struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender    []byte
	SenderEVM common.Address
	ChainID   *big.Int
}

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// TestUniversalContractMetaData contains all meta data concerning the TestUniversalContract contract.
var TestUniversalContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"gatewayAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayZEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onAbort\",\"inputs\":[{\"name\":\"abortContext\",\"type\":\"tuple\",\"internalType\":\"structAbortContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outgoing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRevert\",\"inputs\":[{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ContextData\",\"inputs\":[{\"name\":\"origin\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"msgSender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataAbort\",\"inputs\":[{\"name\":\"abortContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAbortContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outgoing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataRevert\",\"inputs\":[{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60a0604052348015600f57600080fd5b50604051610883380380610883833981016040819052602c91603c565b6001600160a01b0316608052606a565b600060208284031215604d57600080fd5b81516001600160a01b0381168114606357600080fd5b9392505050565b6080516107ff6100846000396000605b01526107ff6000f3fe6080604052600436106100405760003560e01c8063116191b6146100495780632d4cfb7e146100a65780635bcfd616146100c6578063c9028a36146100e657005b3661004757005b005b34801561005557600080fd5b5061007d7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100b257600080fd5b506100476100c13660046101ea565b610106565b3480156100d257600080fd5b506100476100e1366004610255565b610140565b3480156100f257600080fd5b50610047610101366004610312565b6101bb565b7f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7816040516101359190610401565b60405180910390a150565b606081156101575761015482840184610507565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e61018287806105fd565b61019260408a0160208b01610662565b896040013533866040516101ab9695949392919061067d565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c481604051610135919061073f565b6000602082840312156101fc57600080fd5b813567ffffffffffffffff81111561021357600080fd5b820160c0818503121561022557600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461025057600080fd5b919050565b60008060008060006080868803121561026d57600080fd5b853567ffffffffffffffff81111561028457600080fd5b86016060818903121561029657600080fd5b94506102a46020870161022c565b935060408601359250606086013567ffffffffffffffff8111156102c757600080fd5b8601601f810188136102d857600080fd5b803567ffffffffffffffff8111156102ef57600080fd5b88602082840101111561030157600080fd5b959894975092955050506020019190565b60006020828403121561032457600080fd5b813567ffffffffffffffff81111561033b57600080fd5b82016080818503121561022557600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261038257600080fd5b830160208101925035905067ffffffffffffffff8111156103a257600080fd5b8036038213156103b157600080fd5b9250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b602081526000610411838461034d565b60c0602085015261042660e0850182846103b8565b91505073ffffffffffffffffffffffffffffffffffffffff61044a6020860161022c565b16604084015260006040850135905080606085015250606084013580151580821461047457600080fd5b80608086015250506000608085013590508060a08501525061049960a085018561034d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c08601526104ce8382846103b8565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561051957600080fd5b813567ffffffffffffffff81111561053057600080fd5b8201601f8101841361054157600080fd5b803567ffffffffffffffff81111561055b5761055b6104d8565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156105c7576105c76104d8565b6040528181528282016020018610156105df57600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261063257600080fd5b83018035915067ffffffffffffffff82111561064d57600080fd5b6020019150368190038213156103b157600080fd5b60006020828403121561067457600080fd5b6102258261022c565b60a08152600061069160a08301888a6103b8565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b818110156106fb576020818701810151848301820152016106df565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff6107618361022c565b16602082015273ffffffffffffffffffffffffffffffffffffffff6107886020840161022c565b16604082015260008060408401359050806060840152506107ac606084018461034d565b6080808501526107c060a0850182846103b8565b9594505050505056fea2646970667358221220e8e6d0b153e58c66eb709da168b4b14caac0ddd832536453e3fecca9f73ba9cb64736f6c634300081a0033",
}

// TestUniversalContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TestUniversalContractMetaData.ABI instead.
var TestUniversalContractABI = TestUniversalContractMetaData.ABI

// TestUniversalContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestUniversalContractMetaData.Bin instead.
var TestUniversalContractBin = TestUniversalContractMetaData.Bin

// DeployTestUniversalContract deploys a new Ethereum contract, binding an instance of TestUniversalContract to it.
func DeployTestUniversalContract(auth *bind.TransactOpts, backend bind.ContractBackend, gatewayAddress common.Address) (common.Address, *types.Transaction, *TestUniversalContract, error) {
	parsed, err := TestUniversalContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestUniversalContractBin), backend, gatewayAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestUniversalContract{TestUniversalContractCaller: TestUniversalContractCaller{contract: contract}, TestUniversalContractTransactor: TestUniversalContractTransactor{contract: contract}, TestUniversalContractFilterer: TestUniversalContractFilterer{contract: contract}}, nil
}

// TestUniversalContract is an auto generated Go binding around an Ethereum contract.
type TestUniversalContract struct {
	TestUniversalContractCaller     // Read-only binding to the contract
	TestUniversalContractTransactor // Write-only binding to the contract
	TestUniversalContractFilterer   // Log filterer for contract events
}

// TestUniversalContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestUniversalContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUniversalContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestUniversalContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUniversalContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestUniversalContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUniversalContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestUniversalContractSession struct {
	Contract     *TestUniversalContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestUniversalContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestUniversalContractCallerSession struct {
	Contract *TestUniversalContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TestUniversalContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestUniversalContractTransactorSession struct {
	Contract     *TestUniversalContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TestUniversalContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestUniversalContractRaw struct {
	Contract *TestUniversalContract // Generic contract binding to access the raw methods on
}

// TestUniversalContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestUniversalContractCallerRaw struct {
	Contract *TestUniversalContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestUniversalContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestUniversalContractTransactorRaw struct {
	Contract *TestUniversalContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestUniversalContract creates a new instance of TestUniversalContract, bound to a specific deployed contract.
func NewTestUniversalContract(address common.Address, backend bind.ContractBackend) (*TestUniversalContract, error) {
	contract, err := bindTestUniversalContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestUniversalContract{TestUniversalContractCaller: TestUniversalContractCaller{contract: contract}, TestUniversalContractTransactor: TestUniversalContractTransactor{contract: contract}, TestUniversalContractFilterer: TestUniversalContractFilterer{contract: contract}}, nil
}

// NewTestUniversalContractCaller creates a new read-only instance of TestUniversalContract, bound to a specific deployed contract.
func NewTestUniversalContractCaller(address common.Address, caller bind.ContractCaller) (*TestUniversalContractCaller, error) {
	contract, err := bindTestUniversalContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestUniversalContractCaller{contract: contract}, nil
}

// NewTestUniversalContractTransactor creates a new write-only instance of TestUniversalContract, bound to a specific deployed contract.
func NewTestUniversalContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestUniversalContractTransactor, error) {
	contract, err := bindTestUniversalContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestUniversalContractTransactor{contract: contract}, nil
}

// NewTestUniversalContractFilterer creates a new log filterer instance of TestUniversalContract, bound to a specific deployed contract.
func NewTestUniversalContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestUniversalContractFilterer, error) {
	contract, err := bindTestUniversalContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestUniversalContractFilterer{contract: contract}, nil
}

// bindTestUniversalContract binds a generic wrapper to an already deployed contract.
func bindTestUniversalContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestUniversalContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestUniversalContract *TestUniversalContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestUniversalContract.Contract.TestUniversalContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestUniversalContract *TestUniversalContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.TestUniversalContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestUniversalContract *TestUniversalContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.TestUniversalContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestUniversalContract *TestUniversalContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestUniversalContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestUniversalContract *TestUniversalContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestUniversalContract *TestUniversalContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_TestUniversalContract *TestUniversalContractCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestUniversalContract.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_TestUniversalContract *TestUniversalContractSession) Gateway() (common.Address, error) {
	return _TestUniversalContract.Contract.Gateway(&_TestUniversalContract.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_TestUniversalContract *TestUniversalContractCallerSession) Gateway() (common.Address, error) {
	return _TestUniversalContract.Contract.Gateway(&_TestUniversalContract.CallOpts)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_TestUniversalContract *TestUniversalContractTransactor) OnAbort(opts *bind.TransactOpts, abortContext AbortContext) (*types.Transaction, error) {
	return _TestUniversalContract.contract.Transact(opts, "onAbort", abortContext)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_TestUniversalContract *TestUniversalContractSession) OnAbort(abortContext AbortContext) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnAbort(&_TestUniversalContract.TransactOpts, abortContext)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_TestUniversalContract *TestUniversalContractTransactorSession) OnAbort(abortContext AbortContext) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnAbort(&_TestUniversalContract.TransactOpts, abortContext)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) context, address , uint256 , bytes message) returns()
func (_TestUniversalContract *TestUniversalContractTransactor) OnCall(opts *bind.TransactOpts, context MessageContext, arg1 common.Address, arg2 *big.Int, message []byte) (*types.Transaction, error) {
	return _TestUniversalContract.contract.Transact(opts, "onCall", context, arg1, arg2, message)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) context, address , uint256 , bytes message) returns()
func (_TestUniversalContract *TestUniversalContractSession) OnCall(context MessageContext, arg1 common.Address, arg2 *big.Int, message []byte) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnCall(&_TestUniversalContract.TransactOpts, context, arg1, arg2, message)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) context, address , uint256 , bytes message) returns()
func (_TestUniversalContract *TestUniversalContractTransactorSession) OnCall(context MessageContext, arg1 common.Address, arg2 *big.Int, message []byte) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnCall(&_TestUniversalContract.TransactOpts, context, arg1, arg2, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0xc9028a36.
//
// Solidity: function onRevert((address,address,uint256,bytes) revertContext) returns()
func (_TestUniversalContract *TestUniversalContractTransactor) OnRevert(opts *bind.TransactOpts, revertContext RevertContext) (*types.Transaction, error) {
	return _TestUniversalContract.contract.Transact(opts, "onRevert", revertContext)
}

// OnRevert is a paid mutator transaction binding the contract method 0xc9028a36.
//
// Solidity: function onRevert((address,address,uint256,bytes) revertContext) returns()
func (_TestUniversalContract *TestUniversalContractSession) OnRevert(revertContext RevertContext) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnRevert(&_TestUniversalContract.TransactOpts, revertContext)
}

// OnRevert is a paid mutator transaction binding the contract method 0xc9028a36.
//
// Solidity: function onRevert((address,address,uint256,bytes) revertContext) returns()
func (_TestUniversalContract *TestUniversalContractTransactorSession) OnRevert(revertContext RevertContext) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnRevert(&_TestUniversalContract.TransactOpts, revertContext)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestUniversalContract *TestUniversalContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TestUniversalContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestUniversalContract *TestUniversalContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.Fallback(&_TestUniversalContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestUniversalContract *TestUniversalContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.Fallback(&_TestUniversalContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestUniversalContract *TestUniversalContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUniversalContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestUniversalContract *TestUniversalContractSession) Receive() (*types.Transaction, error) {
	return _TestUniversalContract.Contract.Receive(&_TestUniversalContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestUniversalContract *TestUniversalContractTransactorSession) Receive() (*types.Transaction, error) {
	return _TestUniversalContract.Contract.Receive(&_TestUniversalContract.TransactOpts)
}

// TestUniversalContractContextDataIterator is returned from FilterContextData and is used to iterate over the raw logs and unpacked data for ContextData events raised by the TestUniversalContract contract.
type TestUniversalContractContextDataIterator struct {
	Event *TestUniversalContractContextData // Event containing the contract specifics and raw log

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
func (it *TestUniversalContractContextDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestUniversalContractContextData)
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
		it.Event = new(TestUniversalContractContextData)
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
func (it *TestUniversalContractContextDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestUniversalContractContextDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestUniversalContractContextData represents a ContextData event raised by the TestUniversalContract contract.
type TestUniversalContractContextData struct {
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
func (_TestUniversalContract *TestUniversalContractFilterer) FilterContextData(opts *bind.FilterOpts) (*TestUniversalContractContextDataIterator, error) {

	logs, sub, err := _TestUniversalContract.contract.FilterLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return &TestUniversalContractContextDataIterator{contract: _TestUniversalContract.contract, event: "ContextData", logs: logs, sub: sub}, nil
}

// WatchContextData is a free log subscription operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestUniversalContract *TestUniversalContractFilterer) WatchContextData(opts *bind.WatchOpts, sink chan<- *TestUniversalContractContextData) (event.Subscription, error) {

	logs, sub, err := _TestUniversalContract.contract.WatchLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestUniversalContractContextData)
				if err := _TestUniversalContract.contract.UnpackLog(event, "ContextData", log); err != nil {
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
func (_TestUniversalContract *TestUniversalContractFilterer) ParseContextData(log types.Log) (*TestUniversalContractContextData, error) {
	event := new(TestUniversalContractContextData)
	if err := _TestUniversalContract.contract.UnpackLog(event, "ContextData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestUniversalContractContextDataAbortIterator is returned from FilterContextDataAbort and is used to iterate over the raw logs and unpacked data for ContextDataAbort events raised by the TestUniversalContract contract.
type TestUniversalContractContextDataAbortIterator struct {
	Event *TestUniversalContractContextDataAbort // Event containing the contract specifics and raw log

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
func (it *TestUniversalContractContextDataAbortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestUniversalContractContextDataAbort)
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
		it.Event = new(TestUniversalContractContextDataAbort)
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
func (it *TestUniversalContractContextDataAbortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestUniversalContractContextDataAbortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestUniversalContractContextDataAbort represents a ContextDataAbort event raised by the TestUniversalContract contract.
type TestUniversalContractContextDataAbort struct {
	AbortContext AbortContext
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContextDataAbort is a free log retrieval operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_TestUniversalContract *TestUniversalContractFilterer) FilterContextDataAbort(opts *bind.FilterOpts) (*TestUniversalContractContextDataAbortIterator, error) {

	logs, sub, err := _TestUniversalContract.contract.FilterLogs(opts, "ContextDataAbort")
	if err != nil {
		return nil, err
	}
	return &TestUniversalContractContextDataAbortIterator{contract: _TestUniversalContract.contract, event: "ContextDataAbort", logs: logs, sub: sub}, nil
}

// WatchContextDataAbort is a free log subscription operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_TestUniversalContract *TestUniversalContractFilterer) WatchContextDataAbort(opts *bind.WatchOpts, sink chan<- *TestUniversalContractContextDataAbort) (event.Subscription, error) {

	logs, sub, err := _TestUniversalContract.contract.WatchLogs(opts, "ContextDataAbort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestUniversalContractContextDataAbort)
				if err := _TestUniversalContract.contract.UnpackLog(event, "ContextDataAbort", log); err != nil {
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

// ParseContextDataAbort is a log parse operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_TestUniversalContract *TestUniversalContractFilterer) ParseContextDataAbort(log types.Log) (*TestUniversalContractContextDataAbort, error) {
	event := new(TestUniversalContractContextDataAbort)
	if err := _TestUniversalContract.contract.UnpackLog(event, "ContextDataAbort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestUniversalContractContextDataRevertIterator is returned from FilterContextDataRevert and is used to iterate over the raw logs and unpacked data for ContextDataRevert events raised by the TestUniversalContract contract.
type TestUniversalContractContextDataRevertIterator struct {
	Event *TestUniversalContractContextDataRevert // Event containing the contract specifics and raw log

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
func (it *TestUniversalContractContextDataRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestUniversalContractContextDataRevert)
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
		it.Event = new(TestUniversalContractContextDataRevert)
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
func (it *TestUniversalContractContextDataRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestUniversalContractContextDataRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestUniversalContractContextDataRevert represents a ContextDataRevert event raised by the TestUniversalContract contract.
type TestUniversalContractContextDataRevert struct {
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContextDataRevert is a free log retrieval operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_TestUniversalContract *TestUniversalContractFilterer) FilterContextDataRevert(opts *bind.FilterOpts) (*TestUniversalContractContextDataRevertIterator, error) {

	logs, sub, err := _TestUniversalContract.contract.FilterLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return &TestUniversalContractContextDataRevertIterator{contract: _TestUniversalContract.contract, event: "ContextDataRevert", logs: logs, sub: sub}, nil
}

// WatchContextDataRevert is a free log subscription operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_TestUniversalContract *TestUniversalContractFilterer) WatchContextDataRevert(opts *bind.WatchOpts, sink chan<- *TestUniversalContractContextDataRevert) (event.Subscription, error) {

	logs, sub, err := _TestUniversalContract.contract.WatchLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestUniversalContractContextDataRevert)
				if err := _TestUniversalContract.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
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

// ParseContextDataRevert is a log parse operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_TestUniversalContract *TestUniversalContractFilterer) ParseContextDataRevert(log types.Log) (*TestUniversalContractContextDataRevert, error) {
	event := new(TestUniversalContractContextDataRevert)
	if err := _TestUniversalContract.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
