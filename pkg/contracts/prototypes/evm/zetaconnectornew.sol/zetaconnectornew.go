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
	Bin: "0x60806040523480156200001157600080fd5b50604051620011b9380380620011b9833981810160405281019062000037919062000180565b6001600081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480620000a75750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15620000df576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506200021a565b6000815190506200017a8162000200565b92915050565b600080604083850312156200019a5762000199620001fb565b5b6000620001aa8582860162000169565b9250506020620001bd8582860162000169565b9150509250929050565b6000620001d482620001db565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200020b81620001c7565b81146200021757600080fd5b50565b610f8f806200022a6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806302a04c0d14610051578063116191b61461006d578063e8f9cb3a1461008b578063f3fef3a3146100a9575b600080fd5b61006b6004803603810190610066919061082e565b6100c5565b005b610075610279565b6040516100829190610b20565b60405180910390f35b61009361029f565b6040516100a09190610b05565b60405180910390f35b6100c360048036038101906100be91906107ee565b6102c5565b005b6100cd610374565b61013c600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166103c49092919063ffffffff16565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635131ab59600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16868686866040518663ffffffff1660e01b81526004016101c1959493929190610a8e565b600060405180830381600087803b1580156101db57600080fd5b505af11580156101ef573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061021891906108cf565b508373ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced84848460405161026393929190610bf8565b60405180910390a261027361044a565b50505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6102cd610374565b61031a8282600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166103c49092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364826040516103609190610bdd565b60405180910390a261037061044a565b5050565b600260005414156103ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b190610bbd565b60405180910390fd5b6002600081905550565b6104458363a9059cbb60e01b84846040516024016103e3929190610adc565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610454565b505050565b6001600081905550565b60006104b6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661051b9092919063ffffffff16565b905060008151111561051657808060200190518101906104d691906108a2565b610515576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050c90610b9d565b60405180910390fd5b5b505050565b606061052a8484600085610533565b90509392505050565b606082471015610578576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056f90610b5d565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516105a19190610a77565b60006040518083038185875af1925050503d80600081146105de576040519150601f19603f3d011682016040523d82523d6000602084013e6105e3565b606091505b50915091506105f487838387610600565b92505050949350505050565b606083156106635760008351141561065b5761061b85610676565b61065a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065190610b7d565b60405180910390fd5b5b82905061066e565b61066d8383610699565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156106ac5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e09190610b3b565b60405180910390fd5b60006106fc6106f784610c4f565b610c2a565b90508281526020810184848401111561071857610717610e04565b5b610723848285610d62565b509392505050565b60008135905061073a81610f14565b92915050565b60008151905061074f81610f2b565b92915050565b60008083601f84011261076b5761076a610dfa565b5b8235905067ffffffffffffffff81111561078857610787610df5565b5b6020830191508360018202830111156107a4576107a3610dff565b5b9250929050565b600082601f8301126107c0576107bf610dfa565b5b81516107d08482602086016106e9565b91505092915050565b6000813590506107e881610f42565b92915050565b6000806040838503121561080557610804610e0e565b5b60006108138582860161072b565b9250506020610824858286016107d9565b9150509250929050565b6000806000806060858703121561084857610847610e0e565b5b60006108568782880161072b565b9450506020610867878288016107d9565b935050604085013567ffffffffffffffff81111561088857610887610e09565b5b61089487828801610755565b925092505092959194509250565b6000602082840312156108b8576108b7610e0e565b5b60006108c684828501610740565b91505092915050565b6000602082840312156108e5576108e4610e0e565b5b600082015167ffffffffffffffff81111561090357610902610e09565b5b61090f848285016107ab565b91505092915050565b61092181610cc3565b82525050565b60006109338385610c96565b9350610940838584610d53565b61094983610e13565b840190509392505050565b600061095f82610c80565b6109698185610ca7565b9350610979818560208601610d62565b80840191505092915050565b61098e81610d0b565b82525050565b61099d81610d1d565b82525050565b60006109ae82610c8b565b6109b88185610cb2565b93506109c8818560208601610d62565b6109d181610e13565b840191505092915050565b60006109e9602683610cb2565b91506109f482610e24565b604082019050919050565b6000610a0c601d83610cb2565b9150610a1782610e73565b602082019050919050565b6000610a2f602a83610cb2565b9150610a3a82610e9c565b604082019050919050565b6000610a52601f83610cb2565b9150610a5d82610eeb565b602082019050919050565b610a7181610d01565b82525050565b6000610a838284610954565b915081905092915050565b6000608082019050610aa36000830188610918565b610ab06020830187610918565b610abd6040830186610a68565b8181036060830152610ad0818486610927565b90509695505050505050565b6000604082019050610af16000830185610918565b610afe6020830184610a68565b9392505050565b6000602082019050610b1a6000830184610985565b92915050565b6000602082019050610b356000830184610994565b92915050565b60006020820190508181036000830152610b5581846109a3565b905092915050565b60006020820190508181036000830152610b76816109dc565b9050919050565b60006020820190508181036000830152610b96816109ff565b9050919050565b60006020820190508181036000830152610bb681610a22565b9050919050565b60006020820190508181036000830152610bd681610a45565b9050919050565b6000602082019050610bf26000830184610a68565b92915050565b6000604082019050610c0d6000830186610a68565b8181036020830152610c20818486610927565b9050949350505050565b6000610c34610c45565b9050610c408282610d95565b919050565b6000604051905090565b600067ffffffffffffffff821115610c6a57610c69610dc6565b5b610c7382610e13565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610cce82610ce1565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610d1682610d2f565b9050919050565b6000610d2882610d2f565b9050919050565b6000610d3a82610d41565b9050919050565b6000610d4c82610ce1565b9050919050565b82818337600083830152505050565b60005b83811015610d80578082015181840152602081019050610d65565b83811115610d8f576000848401525b50505050565b610d9e82610e13565b810181811067ffffffffffffffff82111715610dbd57610dbc610dc6565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b610f1d81610cc3565b8114610f2857600080fd5b50565b610f3481610cd5565b8114610f3f57600080fd5b50565b610f4b81610d01565b8114610f5657600080fd5b5056fea2646970667358221220ffaaf5577f8164f6527ef50d54a96d5637c72ba10a83df84e4f7c8118e9d835364736f6c63430008070033",
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
