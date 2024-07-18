// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornew

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

// ZetaConnectorNewMetaData contains all meta data concerning the ZetaConnectorNew contract.
var ZetaConnectorNewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zetaToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIGatewayEVM\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200103a3803806200103a83398181016040528101906200003791906200016c565b6001600081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480620000a75750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15620000df576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b81525050505062000206565b6000815190506200016681620001ec565b92915050565b60008060408385031215620001865762000185620001e7565b5b6000620001968582860162000155565b9250506020620001a98582860162000155565b9150509250929050565b6000620001c082620001c7565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620001f781620001b3565b81146200020357600080fd5b50565b60805160601c60a05160601c610de7620002536000396000818160f4015281816101760152818161027101526102a201526000818160d20152818161013a015261024d0152610de76000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806302a04c0d14610051578063116191b61461006d57806321e093b11461008b578063f3fef3a3146100a9575b600080fd5b61006b6004803603810190610066919061078a565b6100c5565b005b61007561024b565b6040516100829190610a33565b60405180910390f35b61009361026f565b6040516100a09190610a18565b60405180910390f35b6100c360048036038101906100be919061074a565b610293565b005b6100cd610340565b6101387f0000000000000000000000000000000000000000000000000000000000000000847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166103909092919063ffffffff16565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635131ab597f0000000000000000000000000000000000000000000000000000000000000000868686866040518663ffffffff1660e01b81526004016101b99594939291906109a1565b600060405180830381600087803b1580156101d357600080fd5b505af11580156101e7573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced84848460405161023593929190610b0b565b60405180910390a2610245610416565b50505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b61029b610340565b6102e682827f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166103909092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648260405161032c9190610af0565b60405180910390a261033c610416565b5050565b60026000541415610386576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037d90610ad0565b60405180910390fd5b6002600081905550565b6104118363a9059cbb60e01b84846040516024016103af9291906109ef565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610420565b505050565b6001600081905550565b6000610482826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166104e79092919063ffffffff16565b90506000815111156104e257808060200190518101906104a291906107fe565b6104e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d890610ab0565b60405180910390fd5b5b505050565b60606104f684846000856104ff565b90509392505050565b606082471015610544576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053b90610a70565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161056d919061098a565b60006040518083038185875af1925050503d80600081146105aa576040519150601f19603f3d011682016040523d82523d6000602084013e6105af565b606091505b50915091506105c0878383876105cc565b92505050949350505050565b6060831561062f57600083511415610627576105e785610642565b610626576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061d90610a90565b60405180910390fd5b5b82905061063a565b6106398383610665565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156106785781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ac9190610a4e565b60405180910390fd5b6000813590506106c481610d6c565b92915050565b6000815190506106d981610d83565b92915050565b60008083601f8401126106f5576106f4610c57565b5b8235905067ffffffffffffffff81111561071257610711610c52565b5b60208301915083600182028301111561072e5761072d610c5c565b5b9250929050565b60008135905061074481610d9a565b92915050565b6000806040838503121561076157610760610c66565b5b600061076f858286016106b5565b925050602061078085828601610735565b9150509250929050565b600080600080606085870312156107a4576107a3610c66565b5b60006107b2878288016106b5565b94505060206107c387828801610735565b935050604085013567ffffffffffffffff8111156107e4576107e3610c61565b5b6107f0878288016106df565b925092505092959194509250565b60006020828403121561081457610813610c66565b5b6000610822848285016106ca565b91505092915050565b61083481610b80565b82525050565b60006108468385610b53565b9350610853838584610c10565b61085c83610c6b565b840190509392505050565b600061087282610b3d565b61087c8185610b64565b935061088c818560208601610c1f565b80840191505092915050565b6108a181610bc8565b82525050565b6108b081610bda565b82525050565b60006108c182610b48565b6108cb8185610b6f565b93506108db818560208601610c1f565b6108e481610c6b565b840191505092915050565b60006108fc602683610b6f565b915061090782610c7c565b604082019050919050565b600061091f601d83610b6f565b915061092a82610ccb565b602082019050919050565b6000610942602a83610b6f565b915061094d82610cf4565b604082019050919050565b6000610965601f83610b6f565b915061097082610d43565b602082019050919050565b61098481610bbe565b82525050565b60006109968284610867565b915081905092915050565b60006080820190506109b6600083018861082b565b6109c3602083018761082b565b6109d0604083018661097b565b81810360608301526109e381848661083a565b90509695505050505050565b6000604082019050610a04600083018561082b565b610a11602083018461097b565b9392505050565b6000602082019050610a2d6000830184610898565b92915050565b6000602082019050610a4860008301846108a7565b92915050565b60006020820190508181036000830152610a6881846108b6565b905092915050565b60006020820190508181036000830152610a89816108ef565b9050919050565b60006020820190508181036000830152610aa981610912565b9050919050565b60006020820190508181036000830152610ac981610935565b9050919050565b60006020820190508181036000830152610ae981610958565b9050919050565b6000602082019050610b05600083018461097b565b92915050565b6000604082019050610b20600083018661097b565b8181036020830152610b3381848661083a565b9050949350505050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610b8b82610b9e565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610bd382610bec565b9050919050565b6000610be582610bec565b9050919050565b6000610bf782610bfe565b9050919050565b6000610c0982610b9e565b9050919050565b82818337600083830152505050565b60005b83811015610c3d578082015181840152602081019050610c22565b83811115610c4c576000848401525b50505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b610d7581610b80565b8114610d8057600080fd5b50565b610d8c81610b92565b8114610d9757600080fd5b50565b610da381610bbe565b8114610dae57600080fd5b5056fea2646970667358221220d14dddafe100dbbc372627ee1d188fb3a3858e5b2f46489e67f1a557d54b3c3764736f6c63430008070033",
}

// ZetaConnectorNewABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNewMetaData.ABI instead.
var ZetaConnectorNewABI = ZetaConnectorNewMetaData.ABI

// ZetaConnectorNewBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNewMetaData.Bin instead.
var ZetaConnectorNewBin = ZetaConnectorNewMetaData.Bin

// DeployZetaConnectorNew deploys a new Ethereum contract, binding an instance of ZetaConnectorNew to it.
func DeployZetaConnectorNew(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address, _zetaToken common.Address) (common.Address, *types.Transaction, *ZetaConnectorNew, error) {
	parsed, err := ZetaConnectorNewMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNewBin), backend, _gateway, _zetaToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNew{ZetaConnectorNewCaller: ZetaConnectorNewCaller{contract: contract}, ZetaConnectorNewTransactor: ZetaConnectorNewTransactor{contract: contract}, ZetaConnectorNewFilterer: ZetaConnectorNewFilterer{contract: contract}}, nil
}

// ZetaConnectorNew is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNew struct {
	ZetaConnectorNewCaller     // Read-only binding to the contract
	ZetaConnectorNewTransactor // Write-only binding to the contract
	ZetaConnectorNewFilterer   // Log filterer for contract events
}

// ZetaConnectorNewCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNewSession struct {
	Contract     *ZetaConnectorNew // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaConnectorNewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNewCallerSession struct {
	Contract *ZetaConnectorNewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ZetaConnectorNewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNewTransactorSession struct {
	Contract     *ZetaConnectorNewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ZetaConnectorNewRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNewRaw struct {
	Contract *ZetaConnectorNew // Generic contract binding to access the raw methods on
}

// ZetaConnectorNewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNewCallerRaw struct {
	Contract *ZetaConnectorNewCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNewTransactorRaw struct {
	Contract *ZetaConnectorNewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNew creates a new instance of ZetaConnectorNew, bound to a specific deployed contract.
func NewZetaConnectorNew(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNew, error) {
	contract, err := bindZetaConnectorNew(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNew{ZetaConnectorNewCaller: ZetaConnectorNewCaller{contract: contract}, ZetaConnectorNewTransactor: ZetaConnectorNewTransactor{contract: contract}, ZetaConnectorNewFilterer: ZetaConnectorNewFilterer{contract: contract}}, nil
}

// NewZetaConnectorNewCaller creates a new read-only instance of ZetaConnectorNew, bound to a specific deployed contract.
func NewZetaConnectorNewCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNewCaller, error) {
	contract, err := bindZetaConnectorNew(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewCaller{contract: contract}, nil
}

// NewZetaConnectorNewTransactor creates a new write-only instance of ZetaConnectorNew, bound to a specific deployed contract.
func NewZetaConnectorNewTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNewTransactor, error) {
	contract, err := bindZetaConnectorNew(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewTransactor{contract: contract}, nil
}

// NewZetaConnectorNewFilterer creates a new log filterer instance of ZetaConnectorNew, bound to a specific deployed contract.
func NewZetaConnectorNewFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNewFilterer, error) {
	contract, err := bindZetaConnectorNew(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewFilterer{contract: contract}, nil
}

// bindZetaConnectorNew binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNew(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNew *ZetaConnectorNewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNew.Contract.ZetaConnectorNewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNew *ZetaConnectorNewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNew.Contract.ZetaConnectorNewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNew *ZetaConnectorNewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNew.Contract.ZetaConnectorNewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNew *ZetaConnectorNewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNew.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNew *ZetaConnectorNewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNew.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNew *ZetaConnectorNewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNew.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNew *ZetaConnectorNewCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNew.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNew *ZetaConnectorNewSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNew.Contract.Gateway(&_ZetaConnectorNew.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNew *ZetaConnectorNewCallerSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNew.Contract.Gateway(&_ZetaConnectorNew.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNew *ZetaConnectorNewCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNew.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNew *ZetaConnectorNewSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNew.Contract.ZetaToken(&_ZetaConnectorNew.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNew *ZetaConnectorNewCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNew.Contract.ZetaToken(&_ZetaConnectorNew.CallOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_ZetaConnectorNew *ZetaConnectorNewTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNew.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_ZetaConnectorNew *ZetaConnectorNewSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNew.Contract.Withdraw(&_ZetaConnectorNew.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_ZetaConnectorNew *ZetaConnectorNewTransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNew.Contract.Withdraw(&_ZetaConnectorNew.TransactOpts, to, amount)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x02a04c0d.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data) returns()
func (_ZetaConnectorNew *ZetaConnectorNewTransactor) WithdrawAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNew.contract.Transact(opts, "withdrawAndCall", to, amount, data)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x02a04c0d.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data) returns()
func (_ZetaConnectorNew *ZetaConnectorNewSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNew.Contract.WithdrawAndCall(&_ZetaConnectorNew.TransactOpts, to, amount, data)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x02a04c0d.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data) returns()
func (_ZetaConnectorNew *ZetaConnectorNewTransactorSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNew.Contract.WithdrawAndCall(&_ZetaConnectorNew.TransactOpts, to, amount, data)
}

// ZetaConnectorNewWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ZetaConnectorNew contract.
type ZetaConnectorNewWithdrawIterator struct {
	Event *ZetaConnectorNewWithdraw // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNewWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNewWithdraw)
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
		it.Event = new(ZetaConnectorNewWithdraw)
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
func (it *ZetaConnectorNewWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNewWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNewWithdraw represents a Withdraw event raised by the ZetaConnectorNew contract.
type ZetaConnectorNewWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNew *ZetaConnectorNewFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNewWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNew.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewWithdrawIterator{contract: _ZetaConnectorNew.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNew *ZetaConnectorNewFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNewWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNew.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNewWithdraw)
				if err := _ZetaConnectorNew.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNew *ZetaConnectorNewFilterer) ParseWithdraw(log types.Log) (*ZetaConnectorNewWithdraw, error) {
	event := new(ZetaConnectorNewWithdraw)
	if err := _ZetaConnectorNew.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNewWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the ZetaConnectorNew contract.
type ZetaConnectorNewWithdrawAndCallIterator struct {
	Event *ZetaConnectorNewWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNewWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNewWithdrawAndCall)
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
		it.Event = new(ZetaConnectorNewWithdrawAndCall)
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
func (it *ZetaConnectorNewWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNewWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNewWithdrawAndCall represents a WithdrawAndCall event raised by the ZetaConnectorNew contract.
type ZetaConnectorNewWithdrawAndCall struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNew *ZetaConnectorNewFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNewWithdrawAndCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNew.contract.FilterLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewWithdrawAndCallIterator{contract: _ZetaConnectorNew.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNew *ZetaConnectorNewFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNewWithdrawAndCall, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNew.contract.WatchLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNewWithdrawAndCall)
				if err := _ZetaConnectorNew.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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

// ParseWithdrawAndCall is a log parse operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNew *ZetaConnectorNewFilterer) ParseWithdrawAndCall(log types.Log) (*ZetaConnectorNewWithdrawAndCall, error) {
	event := new(ZetaConnectorNewWithdrawAndCall)
	if err := _ZetaConnectorNew.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
