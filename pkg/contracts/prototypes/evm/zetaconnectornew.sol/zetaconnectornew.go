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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zeta\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIGatewayEVM\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zeta\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620011d4380380620011d483398181016040528101906200003791906200016c565b6001600081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480620000a75750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15620000df576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b81525050505062000206565b6000815190506200016681620001ec565b92915050565b60008060408385031215620001865762000185620001e7565b5b6000620001968582860162000155565b9250506020620001a98582860162000155565b9150509250929050565b6000620001c082620001c7565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620001f781620001b3565b81146200020357600080fd5b50565b60805160601c60a05160601c610f81620002536000396000818160f4015281816101760152818161029701526102c801526000818160d20152818161013a01526102730152610f816000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806302a04c0d14610051578063116191b61461006d578063e8f9cb3a1461008b578063f3fef3a3146100a9575b600080fd5b61006b60048036038101906100669190610820565b6100c5565b005b610075610271565b6040516100829190610b12565b60405180910390f35b610093610295565b6040516100a09190610af7565b60405180910390f35b6100c360048036038101906100be91906107e0565b6102b9565b005b6100cd610366565b6101387f0000000000000000000000000000000000000000000000000000000000000000847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166103b69092919063ffffffff16565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635131ab597f0000000000000000000000000000000000000000000000000000000000000000868686866040518663ffffffff1660e01b81526004016101b9959493929190610a80565b600060405180830381600087803b1580156101d357600080fd5b505af11580156101e7573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061021091906108c1565b508373ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced84848460405161025b93929190610bea565b60405180910390a261026b61043c565b50505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b6102c1610366565b61030c82827f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166103b69092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364826040516103529190610bcf565b60405180910390a261036261043c565b5050565b600260005414156103ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a390610baf565b60405180910390fd5b6002600081905550565b6104378363a9059cbb60e01b84846040516024016103d5929190610ace565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610446565b505050565b6001600081905550565b60006104a8826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661050d9092919063ffffffff16565b905060008151111561050857808060200190518101906104c89190610894565b610507576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104fe90610b8f565b60405180910390fd5b5b505050565b606061051c8484600085610525565b90509392505050565b60608247101561056a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056190610b4f565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516105939190610a69565b60006040518083038185875af1925050503d80600081146105d0576040519150601f19603f3d011682016040523d82523d6000602084013e6105d5565b606091505b50915091506105e6878383876105f2565b92505050949350505050565b606083156106555760008351141561064d5761060d85610668565b61064c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064390610b6f565b60405180910390fd5b5b829050610660565b61065f838361068b565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60008251111561069e5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d29190610b2d565b60405180910390fd5b60006106ee6106e984610c41565b610c1c565b90508281526020810184848401111561070a57610709610df6565b5b610715848285610d54565b509392505050565b60008135905061072c81610f06565b92915050565b60008151905061074181610f1d565b92915050565b60008083601f84011261075d5761075c610dec565b5b8235905067ffffffffffffffff81111561077a57610779610de7565b5b60208301915083600182028301111561079657610795610df1565b5b9250929050565b600082601f8301126107b2576107b1610dec565b5b81516107c28482602086016106db565b91505092915050565b6000813590506107da81610f34565b92915050565b600080604083850312156107f7576107f6610e00565b5b60006108058582860161071d565b9250506020610816858286016107cb565b9150509250929050565b6000806000806060858703121561083a57610839610e00565b5b60006108488782880161071d565b9450506020610859878288016107cb565b935050604085013567ffffffffffffffff81111561087a57610879610dfb565b5b61088687828801610747565b925092505092959194509250565b6000602082840312156108aa576108a9610e00565b5b60006108b884828501610732565b91505092915050565b6000602082840312156108d7576108d6610e00565b5b600082015167ffffffffffffffff8111156108f5576108f4610dfb565b5b6109018482850161079d565b91505092915050565b61091381610cb5565b82525050565b60006109258385610c88565b9350610932838584610d45565b61093b83610e05565b840190509392505050565b600061095182610c72565b61095b8185610c99565b935061096b818560208601610d54565b80840191505092915050565b61098081610cfd565b82525050565b61098f81610d0f565b82525050565b60006109a082610c7d565b6109aa8185610ca4565b93506109ba818560208601610d54565b6109c381610e05565b840191505092915050565b60006109db602683610ca4565b91506109e682610e16565b604082019050919050565b60006109fe601d83610ca4565b9150610a0982610e65565b602082019050919050565b6000610a21602a83610ca4565b9150610a2c82610e8e565b604082019050919050565b6000610a44601f83610ca4565b9150610a4f82610edd565b602082019050919050565b610a6381610cf3565b82525050565b6000610a758284610946565b915081905092915050565b6000608082019050610a95600083018861090a565b610aa2602083018761090a565b610aaf6040830186610a5a565b8181036060830152610ac2818486610919565b90509695505050505050565b6000604082019050610ae3600083018561090a565b610af06020830184610a5a565b9392505050565b6000602082019050610b0c6000830184610977565b92915050565b6000602082019050610b276000830184610986565b92915050565b60006020820190508181036000830152610b478184610995565b905092915050565b60006020820190508181036000830152610b68816109ce565b9050919050565b60006020820190508181036000830152610b88816109f1565b9050919050565b60006020820190508181036000830152610ba881610a14565b9050919050565b60006020820190508181036000830152610bc881610a37565b9050919050565b6000602082019050610be46000830184610a5a565b92915050565b6000604082019050610bff6000830186610a5a565b8181036020830152610c12818486610919565b9050949350505050565b6000610c26610c37565b9050610c328282610d87565b919050565b6000604051905090565b600067ffffffffffffffff821115610c5c57610c5b610db8565b5b610c6582610e05565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610cc082610cd3565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610d0882610d21565b9050919050565b6000610d1a82610d21565b9050919050565b6000610d2c82610d33565b9050919050565b6000610d3e82610cd3565b9050919050565b82818337600083830152505050565b60005b83811015610d72578082015181840152602081019050610d57565b83811115610d81576000848401525b50505050565b610d9082610e05565b810181811067ffffffffffffffff82111715610daf57610dae610db8565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b610f0f81610cb5565b8114610f1a57600080fd5b50565b610f2681610cc7565b8114610f3157600080fd5b50565b610f3d81610cf3565b8114610f4857600080fd5b5056fea2646970667358221220a6f4adf403ae4d0e302ba8acf592a958a7d462446fca2ab7897a2575097f40d564736f6c63430008070033",
}

// ZetaConnectorNewABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNewMetaData.ABI instead.
var ZetaConnectorNewABI = ZetaConnectorNewMetaData.ABI

// ZetaConnectorNewBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNewMetaData.Bin instead.
var ZetaConnectorNewBin = ZetaConnectorNewMetaData.Bin

// DeployZetaConnectorNew deploys a new Ethereum contract, binding an instance of ZetaConnectorNew to it.
func DeployZetaConnectorNew(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address, _zeta common.Address) (common.Address, *types.Transaction, *ZetaConnectorNew, error) {
	parsed, err := ZetaConnectorNewMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNewBin), backend, _gateway, _zeta)
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

// Zeta is a free data retrieval call binding the contract method 0xe8f9cb3a.
//
// Solidity: function zeta() view returns(address)
func (_ZetaConnectorNew *ZetaConnectorNewCaller) Zeta(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNew.contract.Call(opts, &out, "zeta")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Zeta is a free data retrieval call binding the contract method 0xe8f9cb3a.
//
// Solidity: function zeta() view returns(address)
func (_ZetaConnectorNew *ZetaConnectorNewSession) Zeta() (common.Address, error) {
	return _ZetaConnectorNew.Contract.Zeta(&_ZetaConnectorNew.CallOpts)
}

// Zeta is a free data retrieval call binding the contract method 0xe8f9cb3a.
//
// Solidity: function zeta() view returns(address)
func (_ZetaConnectorNew *ZetaConnectorNewCallerSession) Zeta() (common.Address, error) {
	return _ZetaConnectorNew.Contract.Zeta(&_ZetaConnectorNew.CallOpts)
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
