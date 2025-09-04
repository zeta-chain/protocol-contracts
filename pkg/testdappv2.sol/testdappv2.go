// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testdappv2

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

// TestDAppV2MetaData contains all meta data concerning the TestDAppV2 contract.
var TestDAppV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"gateway_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gatewayTwoDeposits\",\"inputs\":[{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"}]",
	Bin: "0x60a0604052348015600e575f80fd5b50604051610c6d380380610c6d833981016040819052602b91603b565b6001600160a01b03166080526066565b5f60208284031215604a575f80fd5b81516001600160a01b0381168114605f575f80fd5b9392505050565b608051610bc06100ad5f395f818160470152818160e8015281816101b80152818161028101528181610361015281816104450152818161059001526106650152610bc05ff3fe60806040526004361061002b575f3560e01c8063116191b614610036578063d51240c414610092575f80fd5b3661003257005b5f80fd5b348015610041575f80fd5b506100697f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100a56100a03660046108b6565b6100a7565b005b6100e56040518060400160405280600481526020017f48495421000000000000000000000000000000000000000000000000000000008152506106d8565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b01072146040518163ffffffff1660e01b81526004016020604051808303815f875af1158015610150573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610174919061094e565b905061017f8161076a565b5f61018b826002610992565b61019590346109af565b90506101a181346107fb565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663726ac97c6101e86004846109c2565b6040805160a0810182523381525f6020808301829052828401829052835190810184528181526060830152608082015290517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152610251918a91600401610aba565b5f604051808303818588803b158015610268575f80fd5b505af115801561027a573d5f803e3d5ffd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663726ac97c6004836102c891906109c2565b6040805160a0810182523381525f6020808301829052828401829052835190810184528181526060830152608082015290517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152610331918a91600401610aba565b5f604051808303818588803b158015610348575f80fd5b505af115801561035a573d5f803e3d5ffd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663744b9b8b6004836103a891906109c2565b6040805160a0810182523381525f6020808301829052828401829052835190810184528181526060830152608082015290517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152610415918a918a918a91600401610af0565b5f604051808303818588803b15801561042c575f80fd5b505af115801561043e573d5f803e3d5ffd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663744b9b8b60048361048c91906109c2565b6040805160a0810182523381525f6020808301829052828401829052835190810184528181526060830152608082015290517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526104f9918a918a918a91600401610af0565b5f604051808303818588803b158015610510575f80fd5b505af1158015610522573d5f803e3d5ffd5b50506040805160a0810182523381525f6020808301829052828401829052835190810184528181526060830152608082015290517f1becceb400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169450631becceb493508692506105ce918a918a918a91600401610af0565b5f604051808303818588803b1580156105e5575f80fd5b505af11580156105f7573d5f803e3d5ffd5b50506040805160a0810182523381525f6020808301829052828401829052835190810184528181526060830152608082015290517f1becceb400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169450631becceb493508692506106a3918a918a918a91600401610af0565b5f604051808303818588803b1580156106ba575f80fd5b505af11580156106cc573d5f803e3d5ffd5b50505050505050505050565b610767816040516024016106ec9190610b71565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f41304fac00000000000000000000000000000000000000000000000000000000179052610892565b50565b6107678160405160240161078091815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff5b1bba900000000000000000000000000000000000000000000000000000000179052610892565b604051602481018390526044810182905261088e90606401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6c0f698000000000000000000000000000000000000000000000000000000000179052610892565b5050565b6107678180516a636f6e736f6c652e6c6f67602083015f808483855afa5050505050565b5f805f604084860312156108c8575f80fd5b833573ffffffffffffffffffffffffffffffffffffffff811681146108eb575f80fd5b9250602084013567ffffffffffffffff811115610906575f80fd5b8401601f81018613610916575f80fd5b803567ffffffffffffffff81111561092c575f80fd5b86602082840101111561093d575f80fd5b939660209190910195509293505050565b5f6020828403121561095e575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820281158282048414176109a9576109a9610965565b92915050565b818103818111156109a9576109a9610965565b5f826109f5577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff60408201511660408301525f606082015160a06060850152610aa660a08501826109fa565b608093840151949093019390935250919050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f610ae86040830184610a46565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff8516815260606020820152826060820152828460808301375f608084830101525f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011682016080838203016040840152610b666080820185610a46565b979650505050505050565b602081525f610b8360208301846109fa565b939250505056fea26469706673582212209836c527f0b29707020bfa435fb8a3969b017c288b14121f45b26581585074f864736f6c634300081a0033",
}

// TestDAppV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use TestDAppV2MetaData.ABI instead.
var TestDAppV2ABI = TestDAppV2MetaData.ABI

// TestDAppV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestDAppV2MetaData.Bin instead.
var TestDAppV2Bin = TestDAppV2MetaData.Bin

// DeployTestDAppV2 deploys a new Ethereum contract, binding an instance of TestDAppV2 to it.
func DeployTestDAppV2(auth *bind.TransactOpts, backend bind.ContractBackend, gateway_ common.Address) (common.Address, *types.Transaction, *TestDAppV2, error) {
	parsed, err := TestDAppV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestDAppV2Bin), backend, gateway_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestDAppV2{TestDAppV2Caller: TestDAppV2Caller{contract: contract}, TestDAppV2Transactor: TestDAppV2Transactor{contract: contract}, TestDAppV2Filterer: TestDAppV2Filterer{contract: contract}}, nil
}

// TestDAppV2 is an auto generated Go binding around an Ethereum contract.
type TestDAppV2 struct {
	TestDAppV2Caller     // Read-only binding to the contract
	TestDAppV2Transactor // Write-only binding to the contract
	TestDAppV2Filterer   // Log filterer for contract events
}

// TestDAppV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestDAppV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDAppV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestDAppV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDAppV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestDAppV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDAppV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestDAppV2Session struct {
	Contract     *TestDAppV2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestDAppV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestDAppV2CallerSession struct {
	Contract *TestDAppV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestDAppV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestDAppV2TransactorSession struct {
	Contract     *TestDAppV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestDAppV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestDAppV2Raw struct {
	Contract *TestDAppV2 // Generic contract binding to access the raw methods on
}

// TestDAppV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestDAppV2CallerRaw struct {
	Contract *TestDAppV2Caller // Generic read-only contract binding to access the raw methods on
}

// TestDAppV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestDAppV2TransactorRaw struct {
	Contract *TestDAppV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestDAppV2 creates a new instance of TestDAppV2, bound to a specific deployed contract.
func NewTestDAppV2(address common.Address, backend bind.ContractBackend) (*TestDAppV2, error) {
	contract, err := bindTestDAppV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestDAppV2{TestDAppV2Caller: TestDAppV2Caller{contract: contract}, TestDAppV2Transactor: TestDAppV2Transactor{contract: contract}, TestDAppV2Filterer: TestDAppV2Filterer{contract: contract}}, nil
}

// NewTestDAppV2Caller creates a new read-only instance of TestDAppV2, bound to a specific deployed contract.
func NewTestDAppV2Caller(address common.Address, caller bind.ContractCaller) (*TestDAppV2Caller, error) {
	contract, err := bindTestDAppV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestDAppV2Caller{contract: contract}, nil
}

// NewTestDAppV2Transactor creates a new write-only instance of TestDAppV2, bound to a specific deployed contract.
func NewTestDAppV2Transactor(address common.Address, transactor bind.ContractTransactor) (*TestDAppV2Transactor, error) {
	contract, err := bindTestDAppV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestDAppV2Transactor{contract: contract}, nil
}

// NewTestDAppV2Filterer creates a new log filterer instance of TestDAppV2, bound to a specific deployed contract.
func NewTestDAppV2Filterer(address common.Address, filterer bind.ContractFilterer) (*TestDAppV2Filterer, error) {
	contract, err := bindTestDAppV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestDAppV2Filterer{contract: contract}, nil
}

// bindTestDAppV2 binds a generic wrapper to an already deployed contract.
func bindTestDAppV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestDAppV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestDAppV2 *TestDAppV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestDAppV2.Contract.TestDAppV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestDAppV2 *TestDAppV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestDAppV2.Contract.TestDAppV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestDAppV2 *TestDAppV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestDAppV2.Contract.TestDAppV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestDAppV2 *TestDAppV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestDAppV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestDAppV2 *TestDAppV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestDAppV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestDAppV2 *TestDAppV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestDAppV2.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_TestDAppV2 *TestDAppV2Caller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestDAppV2.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_TestDAppV2 *TestDAppV2Session) Gateway() (common.Address, error) {
	return _TestDAppV2.Contract.Gateway(&_TestDAppV2.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_TestDAppV2 *TestDAppV2CallerSession) Gateway() (common.Address, error) {
	return _TestDAppV2.Contract.Gateway(&_TestDAppV2.CallOpts)
}

// GatewayTwoDeposits is a paid mutator transaction binding the contract method 0xd51240c4.
//
// Solidity: function gatewayTwoDeposits(address dst, bytes payload) payable returns()
func (_TestDAppV2 *TestDAppV2Transactor) GatewayTwoDeposits(opts *bind.TransactOpts, dst common.Address, payload []byte) (*types.Transaction, error) {
	return _TestDAppV2.contract.Transact(opts, "gatewayTwoDeposits", dst, payload)
}

// GatewayTwoDeposits is a paid mutator transaction binding the contract method 0xd51240c4.
//
// Solidity: function gatewayTwoDeposits(address dst, bytes payload) payable returns()
func (_TestDAppV2 *TestDAppV2Session) GatewayTwoDeposits(dst common.Address, payload []byte) (*types.Transaction, error) {
	return _TestDAppV2.Contract.GatewayTwoDeposits(&_TestDAppV2.TransactOpts, dst, payload)
}

// GatewayTwoDeposits is a paid mutator transaction binding the contract method 0xd51240c4.
//
// Solidity: function gatewayTwoDeposits(address dst, bytes payload) payable returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) GatewayTwoDeposits(dst common.Address, payload []byte) (*types.Transaction, error) {
	return _TestDAppV2.Contract.GatewayTwoDeposits(&_TestDAppV2.TransactOpts, dst, payload)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestDAppV2 *TestDAppV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestDAppV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestDAppV2 *TestDAppV2Session) Receive() (*types.Transaction, error) {
	return _TestDAppV2.Contract.Receive(&_TestDAppV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestDAppV2 *TestDAppV2TransactorSession) Receive() (*types.Transaction, error) {
	return _TestDAppV2.Contract.Receive(&_TestDAppV2.TransactOpts)
}
