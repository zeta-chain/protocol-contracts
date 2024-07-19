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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIGatewayEVM\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610ee2380380610ee2833981810160405281019061003291906100fd565b6001600081905550600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156100a1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610178565b6000815190506100f781610161565b92915050565b6000602082840312156101135761011261015c565b5b6000610121848285016100e8565b91505092915050565b60006101358261013c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b61016a8161012a565b811461017557600080fd5b50565b610d5b806101876000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063116191b61461004657806321fc65f214610064578063d9caed1214610080575b600080fd5b61004e61009c565b60405161005b91906109b9565b60405180910390f35b61007e60048036038101906100799190610726565b6100c2565b005b61009a600480360381019061009591906106d3565b610224565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6100ca6102c9565b610117600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848773ffffffffffffffffffffffffffffffffffffffff166103199092919063ffffffff16565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635131ab5986868686866040518663ffffffff1660e01b815260040161017a959493929190610942565b600060405180830381600087803b15801561019457600080fd5b505af11580156101a8573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e85858560405161020d93929190610a91565b60405180910390a361021d61039f565b5050505050565b61022c6102c9565b61025782828573ffffffffffffffffffffffffffffffffffffffff166103199092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb836040516102b49190610a76565b60405180910390a36102c461039f565b505050565b6002600054141561030f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030690610a56565b60405180910390fd5b6002600081905550565b61039a8363a9059cbb60e01b8484604051602401610338929190610990565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506103a9565b505050565b6001600081905550565b600061040b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166104709092919063ffffffff16565b905060008151111561046b578080602001905181019061042b91906107ae565b61046a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046190610a36565b60405180910390fd5b5b505050565b606061047f8484600085610488565b90509392505050565b6060824710156104cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c4906109f6565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516104f6919061092b565b60006040518083038185875af1925050503d8060008114610533576040519150601f19603f3d011682016040523d82523d6000602084013e610538565b606091505b509150915061054987838387610555565b92505050949350505050565b606083156105b8576000835114156105b057610570856105cb565b6105af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a690610a16565b60405180910390fd5b5b8290506105c3565b6105c283836105ee565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156106015781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063591906109d4565b60405180910390fd5b60008135905061064d81610ce0565b92915050565b60008151905061066281610cf7565b92915050565b60008083601f84011261067e5761067d610bcb565b5b8235905067ffffffffffffffff81111561069b5761069a610bc6565b5b6020830191508360018202830111156106b7576106b6610bd0565b5b9250929050565b6000813590506106cd81610d0e565b92915050565b6000806000606084860312156106ec576106eb610bda565b5b60006106fa8682870161063e565b935050602061070b8682870161063e565b925050604061071c868287016106be565b9150509250925092565b60008060008060006080868803121561074257610741610bda565b5b60006107508882890161063e565b95505060206107618882890161063e565b9450506040610772888289016106be565b935050606086013567ffffffffffffffff81111561079357610792610bd5565b5b61079f88828901610668565b92509250509295509295909350565b6000602082840312156107c4576107c3610bda565b5b60006107d284828501610653565b91505092915050565b6107e481610b06565b82525050565b60006107f68385610ad9565b9350610803838584610b84565b61080c83610bdf565b840190509392505050565b600061082282610ac3565b61082c8185610aea565b935061083c818560208601610b93565b80840191505092915050565b61085181610b4e565b82525050565b600061086282610ace565b61086c8185610af5565b935061087c818560208601610b93565b61088581610bdf565b840191505092915050565b600061089d602683610af5565b91506108a882610bf0565b604082019050919050565b60006108c0601d83610af5565b91506108cb82610c3f565b602082019050919050565b60006108e3602a83610af5565b91506108ee82610c68565b604082019050919050565b6000610906601f83610af5565b915061091182610cb7565b602082019050919050565b61092581610b44565b82525050565b60006109378284610817565b915081905092915050565b600060808201905061095760008301886107db565b61096460208301876107db565b610971604083018661091c565b81810360608301526109848184866107ea565b90509695505050505050565b60006040820190506109a560008301856107db565b6109b2602083018461091c565b9392505050565b60006020820190506109ce6000830184610848565b92915050565b600060208201905081810360008301526109ee8184610857565b905092915050565b60006020820190508181036000830152610a0f81610890565b9050919050565b60006020820190508181036000830152610a2f816108b3565b9050919050565b60006020820190508181036000830152610a4f816108d6565b9050919050565b60006020820190508181036000830152610a6f816108f9565b9050919050565b6000602082019050610a8b600083018461091c565b92915050565b6000604082019050610aa6600083018661091c565b8181036020830152610ab98184866107ea565b9050949350505050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610b1182610b24565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610b5982610b60565b9050919050565b6000610b6b82610b72565b9050919050565b6000610b7d82610b24565b9050919050565b82818337600083830152505050565b60005b83811015610bb1578082015181840152602081019050610b96565b83811115610bc0576000848401525b50505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b610ce981610b06565b8114610cf457600080fd5b50565b610d0081610b18565b8114610d0b57600080fd5b50565b610d1781610b44565b8114610d2257600080fd5b5056fea2646970667358221220e0308fb4ed96edbeb2c69e14a5ca1b29a3ef7df2a96cb8bc20f2642e1f32fae964736f6c63430008070033",
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
