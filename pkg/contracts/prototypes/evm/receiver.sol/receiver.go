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
	Bin: "0x608060405234801561001057600080fd5b5061108a806100206000396000f3fe60806040526004361061003f5760003560e01c806352df08b6146100445780636fa220ad1461006d57806386c4519214610089578063f5db6b39146100a0575b600080fd5b34801561005057600080fd5b5061006b6004803603810190610066919061085a565b6100c9565b005b610087600480360381019061008291906107eb565b610138565b005b34801561009557600080fd5b5061009e61017c565b005b3480156100ac57600080fd5b506100c760048036038101906100c29190610733565b6101b5565b005b6100f63382858573ffffffffffffffffffffffffffffffffffffffff166101f7909392919063ffffffff16565b7fe8077791e8d0f63b9e4c0b6d386cbbf6c65e1acb2a103019edba9d1cc0b329003384848460405161012b9493929190610bb0565b60405180910390a1505050565b7f87d67858b5cc03bdd16a3fcc56773bd59410e946cff7193cf374402c6e8fb6ee333485858560405161016f959493929190610bf5565b60405180910390a1505050565b7fcf0e6f18d967cb5a3ca7781d74b1d66411d1b8984e2dd2a066709c204a66d862336040516101ab9190610b0b565b60405180910390a1565b7f463b3e5d6969d17b19e10c4ca95f5f1b6fdef8678f980af98071bd38f0d885c2338484846040516101ea9493929190610b5d565b60405180910390a1505050565b61027a846323b872dd60e01b85858560405160240161021893929190610b26565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610280565b50505050565b60006102e2826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166103479092919063ffffffff16565b9050600081511115610342578080602001905181019061030291906107be565b610341576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033890610cb1565b60405180910390fd5b5b505050565b6060610356848460008561035f565b90509392505050565b6060824710156103a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039b90610c71565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516103cd9190610af4565b60006040518083038185875af1925050503d806000811461040a576040519150601f19603f3d011682016040523d82523d6000602084013e61040f565b606091505b50915091506104208783838761042c565b92505050949350505050565b6060831561048f5760008351141561048757610447856104a2565b610486576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047d90610c91565b60405180910390fd5b5b82905061049a565b61049983836104c5565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156104d85781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050c9190610c4f565b60405180910390fd5b600061052861052384610cf6565b610cd1565b9050808382526020820190508285602086028201111561054b5761054a610f23565b5b60005b8581101561059957813567ffffffffffffffff81111561057157610570610f1e565b5b80860161057e89826106f0565b8552602085019450602084019350505060018101905061054e565b5050509392505050565b60006105b66105b184610d22565b610cd1565b905080838252602082019050828560208602820111156105d9576105d8610f23565b5b60005b8581101561060957816105ef888261071e565b8452602084019350602083019250506001810190506105dc565b5050509392505050565b600061062661062184610d4e565b610cd1565b90508281526020810184848401111561064257610641610f28565b5b61064d848285610e7c565b509392505050565b6000813590506106648161100f565b92915050565b600082601f83011261067f5761067e610f1e565b5b813561068f848260208601610515565b91505092915050565b600082601f8301126106ad576106ac610f1e565b5b81356106bd8482602086016105a3565b91505092915050565b6000813590506106d581611026565b92915050565b6000815190506106ea81611026565b92915050565b600082601f83011261070557610704610f1e565b5b8135610715848260208601610613565b91505092915050565b60008135905061072d8161103d565b92915050565b60008060006060848603121561074c5761074b610f32565b5b600084013567ffffffffffffffff81111561076a57610769610f2d565b5b6107768682870161066a565b935050602084013567ffffffffffffffff81111561079757610796610f2d565b5b6107a386828701610698565b92505060406107b4868287016106c6565b9150509250925092565b6000602082840312156107d4576107d3610f32565b5b60006107e2848285016106db565b91505092915050565b60008060006060848603121561080457610803610f32565b5b600084013567ffffffffffffffff81111561082257610821610f2d565b5b61082e868287016106f0565b935050602061083f8682870161071e565b9250506040610850868287016106c6565b9150509250925092565b60008060006060848603121561087357610872610f32565b5b60006108818682870161071e565b935050602061089286828701610655565b92505060406108a386828701610655565b9150509250925092565b60006108b983836109fb565b905092915050565b60006108cd8383610ad6565b60208301905092915050565b6108e281610e34565b82525050565b60006108f382610d9f565b6108fd8185610de5565b93508360208202850161090f85610d7f565b8060005b8581101561094b578484038952815161092c85826108ad565b945061093783610dcb565b925060208a01995050600181019050610913565b50829750879550505050505092915050565b600061096882610daa565b6109728185610df6565b935061097d83610d8f565b8060005b838110156109ae57815161099588826108c1565b97506109a083610dd8565b925050600181019050610981565b5085935050505092915050565b6109c481610e46565b82525050565b60006109d582610db5565b6109df8185610e07565b93506109ef818560208601610e8b565b80840191505092915050565b6000610a0682610dc0565b610a108185610e12565b9350610a20818560208601610e8b565b610a2981610f37565b840191505092915050565b6000610a3f82610dc0565b610a498185610e23565b9350610a59818560208601610e8b565b610a6281610f37565b840191505092915050565b6000610a7a602683610e23565b9150610a8582610f48565b604082019050919050565b6000610a9d601d83610e23565b9150610aa882610f97565b602082019050919050565b6000610ac0602a83610e23565b9150610acb82610fc0565b604082019050919050565b610adf81610e72565b82525050565b610aee81610e72565b82525050565b6000610b0082846109ca565b915081905092915050565b6000602082019050610b2060008301846108d9565b92915050565b6000606082019050610b3b60008301866108d9565b610b4860208301856108d9565b610b556040830184610ae5565b949350505050565b6000608082019050610b7260008301876108d9565b8181036020830152610b8481866108e8565b90508181036040830152610b98818561095d565b9050610ba760608301846109bb565b95945050505050565b6000608082019050610bc560008301876108d9565b610bd26020830186610ae5565b610bdf60408301856108d9565b610bec60608301846108d9565b95945050505050565b600060a082019050610c0a60008301886108d9565b610c176020830187610ae5565b8181036040830152610c298186610a34565b9050610c386060830185610ae5565b610c4560808301846109bb565b9695505050505050565b60006020820190508181036000830152610c698184610a34565b905092915050565b60006020820190508181036000830152610c8a81610a6d565b9050919050565b60006020820190508181036000830152610caa81610a90565b9050919050565b60006020820190508181036000830152610cca81610ab3565b9050919050565b6000610cdb610cec565b9050610ce78282610ebe565b919050565b6000604051905090565b600067ffffffffffffffff821115610d1157610d10610eef565b5b602082029050602081019050919050565b600067ffffffffffffffff821115610d3d57610d3c610eef565b5b602082029050602081019050919050565b600067ffffffffffffffff821115610d6957610d68610eef565b5b610d7282610f37565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610e3f82610e52565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610ea9578082015181840152602081019050610e8e565b83811115610eb8576000848401525b50505050565b610ec782610f37565b810181811067ffffffffffffffff82111715610ee657610ee5610eef565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b61101881610e34565b811461102357600080fd5b50565b61102f81610e46565b811461103a57600080fd5b50565b61104681610e72565b811461105157600080fd5b5056fea2646970667358221220e384ec89cbb2f2f1ac377f3aa6bce74ab1369ca50a5280a0f885097a981f043e64736f6c63430008070033",
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
