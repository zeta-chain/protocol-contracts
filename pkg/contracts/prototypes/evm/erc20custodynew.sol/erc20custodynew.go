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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndRevert\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIGatewayEVM\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdrawAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200121938038062001219833981810160405281019062000037919062000106565b6001600081905550600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620000a7576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200018b565b600081519050620001008162000171565b92915050565b6000602082840312156200011f576200011e6200016c565b5b60006200012f84828501620000ef565b91505092915050565b600062000145826200014c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200017c8162000138565b81146200018857600080fd5b50565b61107e806200019b6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063116191b61461005157806321fc65f21461006f578063c8a023621461008b578063d9caed12146100a7575b600080fd5b6100596100c3565b6040516100669190610c21565b60405180910390f35b61008960048036038101906100849190610945565b6100e9565b005b6100a560048036038101906100a09190610945565b610271565b005b6100c160048036038101906100bc91906108f2565b6103d3565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6100f1610478565b61013e600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848773ffffffffffffffffffffffffffffffffffffffff166104c89092919063ffffffff16565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635131ab5986868686866040518663ffffffff1660e01b81526004016101a1959493929190610baa565b600060405180830381600087803b1580156101bb57600080fd5b505af11580156101cf573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906101f891906109fa565b508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e85858560405161025a93929190610cf9565b60405180910390a361026a61054e565b5050505050565b610279610478565b6102c6600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848773ffffffffffffffffffffffffffffffffffffffff166104c89092919063ffffffff16565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b8969bd486868686866040518663ffffffff1660e01b8152600401610329959493929190610baa565b600060405180830381600087803b15801561034357600080fd5b505af1158015610357573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c88585856040516103bc93929190610cf9565b60405180910390a36103cc61054e565b5050505050565b6103db610478565b61040682828573ffffffffffffffffffffffffffffffffffffffff166104c89092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb836040516104639190610cde565b60405180910390a361047361054e565b505050565b600260005414156104be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b590610cbe565b60405180910390fd5b6002600081905550565b6105498363a9059cbb60e01b84846040516024016104e7929190610bf8565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610558565b505050565b6001600081905550565b60006105ba826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661061f9092919063ffffffff16565b905060008151111561061a57808060200190518101906105da91906109cd565b610619576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061090610c9e565b60405180910390fd5b5b505050565b606061062e8484600085610637565b90509392505050565b60608247101561067c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067390610c5e565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516106a59190610b93565b60006040518083038185875af1925050503d80600081146106e2576040519150601f19603f3d011682016040523d82523d6000602084013e6106e7565b606091505b50915091506106f887838387610704565b92505050949350505050565b606083156107675760008351141561075f5761071f8561077a565b61075e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075590610c7e565b60405180910390fd5b5b829050610772565b610771838361079d565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156107b05781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e49190610c3c565b60405180910390fd5b60006108006107fb84610d50565b610d2b565b90508281526020810184848401111561081c5761081b610ef3565b5b610827848285610e51565b509392505050565b60008135905061083e81611003565b92915050565b6000815190506108538161101a565b92915050565b60008083601f84011261086f5761086e610ee9565b5b8235905067ffffffffffffffff81111561088c5761088b610ee4565b5b6020830191508360018202830111156108a8576108a7610eee565b5b9250929050565b600082601f8301126108c4576108c3610ee9565b5b81516108d48482602086016107ed565b91505092915050565b6000813590506108ec81611031565b92915050565b60008060006060848603121561090b5761090a610efd565b5b60006109198682870161082f565b935050602061092a8682870161082f565b925050604061093b868287016108dd565b9150509250925092565b60008060008060006080868803121561096157610960610efd565b5b600061096f8882890161082f565b95505060206109808882890161082f565b9450506040610991888289016108dd565b935050606086013567ffffffffffffffff8111156109b2576109b1610ef8565b5b6109be88828901610859565b92509250509295509295909350565b6000602082840312156109e3576109e2610efd565b5b60006109f184828501610844565b91505092915050565b600060208284031215610a1057610a0f610efd565b5b600082015167ffffffffffffffff811115610a2e57610a2d610ef8565b5b610a3a848285016108af565b91505092915050565b610a4c81610dc4565b82525050565b6000610a5e8385610d97565b9350610a6b838584610e42565b610a7483610f02565b840190509392505050565b6000610a8a82610d81565b610a948185610da8565b9350610aa4818560208601610e51565b80840191505092915050565b610ab981610e0c565b82525050565b6000610aca82610d8c565b610ad48185610db3565b9350610ae4818560208601610e51565b610aed81610f02565b840191505092915050565b6000610b05602683610db3565b9150610b1082610f13565b604082019050919050565b6000610b28601d83610db3565b9150610b3382610f62565b602082019050919050565b6000610b4b602a83610db3565b9150610b5682610f8b565b604082019050919050565b6000610b6e601f83610db3565b9150610b7982610fda565b602082019050919050565b610b8d81610e02565b82525050565b6000610b9f8284610a7f565b915081905092915050565b6000608082019050610bbf6000830188610a43565b610bcc6020830187610a43565b610bd96040830186610b84565b8181036060830152610bec818486610a52565b90509695505050505050565b6000604082019050610c0d6000830185610a43565b610c1a6020830184610b84565b9392505050565b6000602082019050610c366000830184610ab0565b92915050565b60006020820190508181036000830152610c568184610abf565b905092915050565b60006020820190508181036000830152610c7781610af8565b9050919050565b60006020820190508181036000830152610c9781610b1b565b9050919050565b60006020820190508181036000830152610cb781610b3e565b9050919050565b60006020820190508181036000830152610cd781610b61565b9050919050565b6000602082019050610cf36000830184610b84565b92915050565b6000604082019050610d0e6000830186610b84565b8181036020830152610d21818486610a52565b9050949350505050565b6000610d35610d46565b9050610d418282610e84565b919050565b6000604051905090565b600067ffffffffffffffff821115610d6b57610d6a610eb5565b5b610d7482610f02565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610dcf82610de2565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610e1782610e1e565b9050919050565b6000610e2982610e30565b9050919050565b6000610e3b82610de2565b9050919050565b82818337600083830152505050565b60005b83811015610e6f578082015181840152602081019050610e54565b83811115610e7e576000848401525b50505050565b610e8d82610f02565b810181811067ffffffffffffffff82111715610eac57610eab610eb5565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b61100c81610dc4565b811461101757600080fd5b50565b61102381610dd6565b811461102e57600080fd5b50565b61103a81610e02565b811461104557600080fd5b5056fea26469706673582212205d03adc1d4bdce0ac9310d24d3164abfa5397ef53e14dff14823a8864c09e42164736f6c63430008070033",
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
