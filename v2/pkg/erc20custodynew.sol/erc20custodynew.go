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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_gateway\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tssAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndRevert\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawAndCall\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawAndRevert\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610b4a380380610b4a83398101604081905261002f916100bc565b60016000556001600160a01b038216158061005157506001600160a01b038116155b1561006f5760405163d92e233d60e01b815260040160405180910390fd5b600180546001600160a01b039384166001600160a01b031991821617909155600280549290931691161790556100ef565b80516001600160a01b03811681146100b757600080fd5b919050565b600080604083850312156100cf57600080fd5b6100d8836100a0565b91506100e6602084016100a0565b90509250929050565b610a4c806100fe6000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80635b112591116100505780635b112591146100ca578063c8a02362146100ea578063d9caed12146100fd57600080fd5b8063116191b61461006c57806321fc65f2146100b5575b600080fd5b60015461008c9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100c86100c3366004610822565b610110565b005b60025461008c9073ffffffffffffffffffffffffffffffffffffffff1681565b6100c86100f8366004610822565b61029a565b6100c861010b3660046108bf565b61040b565b6101186104fb565b60025473ffffffffffffffffffffffffffffffffffffffff163314610169576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546101909073ffffffffffffffffffffffffffffffffffffffff87811691168561053e565b6001546040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690635131ab59906101ee9088908890889088908890600401610945565b600060405180830381600087803b15801561020857600080fd5b505af115801561021c573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e858585604051610281939291906109a2565b60405180910390a36102936001600055565b5050505050565b6102a26104fb565b60025473ffffffffffffffffffffffffffffffffffffffff1633146102f3576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015461031a9073ffffffffffffffffffffffffffffffffffffffff87811691168561053e565b6001546040517fb8969bd400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063b8969bd4906103789088908890889088908890600401610945565b600060405180830381600087803b15801561039257600080fd5b505af11580156103a6573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c8858585604051610281939291906109a2565b6104136104fb565b60025473ffffffffffffffffffffffffffffffffffffffff163314610464576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61048573ffffffffffffffffffffffffffffffffffffffff8416838361053e565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb836040516104e491815260200190565b60405180910390a36104f66001600055565b505050565b600260005403610537576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff848116602483015260448083018590528351808403909101815260649092019092526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526104f6918591906000906105d790841683610650565b905080516000141580156105fc5750808060200190518101906105fa91906109c5565b155b156104f6576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b606061065e83836000610665565b9392505050565b6060814710156106a3576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610647565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516106cc91906109e7565b60006040518083038185875af1925050503d8060008114610709576040519150601f19603f3d011682016040523d82523d6000602084013e61070e565b606091505b509150915061071e868383610728565b9695505050505050565b60608261073d57610738826107b7565b61065e565b8151158015610761575073ffffffffffffffffffffffffffffffffffffffff84163b155b156107b0576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610647565b508061065e565b8051156107c75780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461081d57600080fd5b919050565b60008060008060006080868803121561083a57600080fd5b610843866107f9565b9450610851602087016107f9565b935060408601359250606086013567ffffffffffffffff81111561087457600080fd5b8601601f8101881361088557600080fd5b803567ffffffffffffffff81111561089c57600080fd5b8860208284010111156108ae57600080fd5b959894975092955050506020019190565b6000806000606084860312156108d457600080fd5b6108dd846107f9565b92506108eb602085016107f9565b929592945050506040919091013590565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006109976080830184866108fc565b979650505050505050565b8381526040602082015260006109bc6040830184866108fc565b95945050505050565b6000602082840312156109d757600080fd5b8151801515811461065e57600080fd5b6000825160005b81811015610a0857602081860181015185830152016109ee565b50600092019182525091905056fea264697066735822122041336b1568f49be1aa3dcd208e804b373772741e9b7e1726a8869e99fced60c764736f6c634300081a0033",
}

// ERC20CustodyNewABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20CustodyNewMetaData.ABI instead.
var ERC20CustodyNewABI = ERC20CustodyNewMetaData.ABI

// ERC20CustodyNewBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20CustodyNewMetaData.Bin instead.
var ERC20CustodyNewBin = ERC20CustodyNewMetaData.Bin

// DeployERC20CustodyNew deploys a new Ethereum contract, binding an instance of ERC20CustodyNew to it.
func DeployERC20CustodyNew(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address, _tssAddress common.Address) (common.Address, *types.Transaction, *ERC20CustodyNew, error) {
	parsed, err := ERC20CustodyNewMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20CustodyNewBin), backend, _gateway, _tssAddress)
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

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ERC20CustodyNew *ERC20CustodyNewCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20CustodyNew.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ERC20CustodyNew *ERC20CustodyNewSession) TssAddress() (common.Address, error) {
	return _ERC20CustodyNew.Contract.TssAddress(&_ERC20CustodyNew.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ERC20CustodyNew *ERC20CustodyNewCallerSession) TssAddress() (common.Address, error) {
	return _ERC20CustodyNew.Contract.TssAddress(&_ERC20CustodyNew.CallOpts)
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

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0xc8a02362.
//
// Solidity: function withdrawAndRevert(address token, address to, uint256 amount, bytes data) returns()
func (_ERC20CustodyNew *ERC20CustodyNewTransactor) WithdrawAndRevert(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20CustodyNew.contract.Transact(opts, "withdrawAndRevert", token, to, amount, data)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0xc8a02362.
//
// Solidity: function withdrawAndRevert(address token, address to, uint256 amount, bytes data) returns()
func (_ERC20CustodyNew *ERC20CustodyNewSession) WithdrawAndRevert(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20CustodyNew.Contract.WithdrawAndRevert(&_ERC20CustodyNew.TransactOpts, token, to, amount, data)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0xc8a02362.
//
// Solidity: function withdrawAndRevert(address token, address to, uint256 amount, bytes data) returns()
func (_ERC20CustodyNew *ERC20CustodyNewTransactorSession) WithdrawAndRevert(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20CustodyNew.Contract.WithdrawAndRevert(&_ERC20CustodyNew.TransactOpts, token, to, amount, data)
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

// ERC20CustodyNewWithdrawAndRevertIterator is returned from FilterWithdrawAndRevert and is used to iterate over the raw logs and unpacked data for WithdrawAndRevert events raised by the ERC20CustodyNew contract.
type ERC20CustodyNewWithdrawAndRevertIterator struct {
	Event *ERC20CustodyNewWithdrawAndRevert // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyNewWithdrawAndRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyNewWithdrawAndRevert)
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
		it.Event = new(ERC20CustodyNewWithdrawAndRevert)
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
func (it *ERC20CustodyNewWithdrawAndRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyNewWithdrawAndRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyNewWithdrawAndRevert represents a WithdrawAndRevert event raised by the ERC20CustodyNew contract.
type ERC20CustodyNewWithdrawAndRevert struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndRevert is a free log retrieval operation binding the contract event 0xb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c8.
//
// Solidity: event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ERC20CustodyNew *ERC20CustodyNewFilterer) FilterWithdrawAndRevert(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*ERC20CustodyNewWithdrawAndRevertIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20CustodyNew.contract.FilterLogs(opts, "WithdrawAndRevert", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyNewWithdrawAndRevertIterator{contract: _ERC20CustodyNew.contract, event: "WithdrawAndRevert", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndRevert is a free log subscription operation binding the contract event 0xb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c8.
//
// Solidity: event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ERC20CustodyNew *ERC20CustodyNewFilterer) WatchWithdrawAndRevert(opts *bind.WatchOpts, sink chan<- *ERC20CustodyNewWithdrawAndRevert, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20CustodyNew.contract.WatchLogs(opts, "WithdrawAndRevert", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyNewWithdrawAndRevert)
				if err := _ERC20CustodyNew.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
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

// ParseWithdrawAndRevert is a log parse operation binding the contract event 0xb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c8.
//
// Solidity: event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ERC20CustodyNew *ERC20CustodyNewFilterer) ParseWithdrawAndRevert(log types.Log) (*ERC20CustodyNewWithdrawAndRevert, error) {
	event := new(ERC20CustodyNewWithdrawAndRevert)
	if err := _ERC20CustodyNew.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
