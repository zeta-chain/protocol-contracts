// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectorneweth

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

// ZetaConnectorNewEthMetaData contains all meta data concerning the ZetaConnectorNewEth contract.
var ZetaConnectorNewEthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zetaToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIGatewayEVM\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"receiveTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620011f8380380620011f8833981810160405281019062000037919062000170565b81816001600081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480620000a95750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15620000e1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b81525050505050506200020a565b6000815190506200016a81620001f0565b92915050565b600080604083850312156200018a5762000189620001eb565b5b60006200019a8582860162000159565b9250506020620001ad8582860162000159565b9150509250929050565b6000620001c482620001cb565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620001fb81620001b7565b81146200020757600080fd5b50565b60805160601c60a05160601c610f996200025f6000396000818160fb015281816101c00152818161021101528181610293015261037901526000818161019c015281816101ef01526102570152610f996000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063106e62901461005c578063116191b61461007857806321e093b1146100965780635e3e9fef146100b4578063743e0c9b146100d0575b600080fd5b61007660048036038101906100719190610871565b6100ec565b005b61008061019a565b60405161008d9190610bd6565b60405180910390f35b61009e6101be565b6040516100ab9190610b0d565b60405180910390f35b6100ce60048036038101906100c991906108c4565b6101e2565b005b6100ea60048036038101906100e59190610979565b610369565b005b6100f46103c9565b61013f83837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166104199092919063ffffffff16565b8273ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364836040516101859190610c93565b60405180910390a261019561049f565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b6101ea6103c9565b6102557f0000000000000000000000000000000000000000000000000000000000000000857f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166104199092919063ffffffff16565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635131ab597f0000000000000000000000000000000000000000000000000000000000000000878787876040518663ffffffff1660e01b81526004016102d6959493929190610b5f565b600060405180830381600087803b1580156102f057600080fd5b505af1158015610304573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced85858560405161035293929190610cae565b60405180910390a261036261049f565b5050505050565b6103716103c9565b6103be3330837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166104a9909392919063ffffffff16565b6103c661049f565b50565b6002600054141561040f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040690610c73565b60405180910390fd5b6002600081905550565b61049a8363a9059cbb60e01b8484604051602401610438929190610bad565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610532565b505050565b6001600081905550565b61052c846323b872dd60e01b8585856040516024016104ca93929190610b28565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610532565b50505050565b6000610594826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166105f99092919063ffffffff16565b90506000815111156105f457808060200190518101906105b4919061094c565b6105f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ea90610c53565b60405180910390fd5b5b505050565b60606106088484600085610611565b90509392505050565b606082471015610656576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064d90610c13565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161067f9190610af6565b60006040518083038185875af1925050503d80600081146106bc576040519150601f19603f3d011682016040523d82523d6000602084013e6106c1565b606091505b50915091506106d2878383876106de565b92505050949350505050565b6060831561074157600083511415610739576106f985610754565b610738576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072f90610c33565b60405180910390fd5b5b82905061074c565b61074b8383610777565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60008251111561078a5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107be9190610bf1565b60405180910390fd5b6000813590506107d681610f07565b92915050565b6000815190506107eb81610f1e565b92915050565b60008135905061080081610f35565b92915050565b60008083601f84011261081c5761081b610df2565b5b8235905067ffffffffffffffff81111561083957610838610ded565b5b60208301915083600182028301111561085557610854610df7565b5b9250929050565b60008135905061086b81610f4c565b92915050565b60008060006060848603121561088a57610889610e01565b5b6000610898868287016107c7565b93505060206108a98682870161085c565b92505060406108ba868287016107f1565b9150509250925092565b6000806000806000608086880312156108e0576108df610e01565b5b60006108ee888289016107c7565b95505060206108ff8882890161085c565b945050604086013567ffffffffffffffff8111156109205761091f610dfc565b5b61092c88828901610806565b9350935050606061093f888289016107f1565b9150509295509295909350565b60006020828403121561096257610961610e01565b5b6000610970848285016107dc565b91505092915050565b60006020828403121561098f5761098e610e01565b5b600061099d8482850161085c565b91505092915050565b6109af81610d23565b82525050565b60006109c18385610cf6565b93506109ce838584610dab565b6109d783610e06565b840190509392505050565b60006109ed82610ce0565b6109f78185610d07565b9350610a07818560208601610dba565b80840191505092915050565b610a1c81610d75565b82525050565b6000610a2d82610ceb565b610a378185610d12565b9350610a47818560208601610dba565b610a5081610e06565b840191505092915050565b6000610a68602683610d12565b9150610a7382610e17565b604082019050919050565b6000610a8b601d83610d12565b9150610a9682610e66565b602082019050919050565b6000610aae602a83610d12565b9150610ab982610e8f565b604082019050919050565b6000610ad1601f83610d12565b9150610adc82610ede565b602082019050919050565b610af081610d6b565b82525050565b6000610b0282846109e2565b915081905092915050565b6000602082019050610b2260008301846109a6565b92915050565b6000606082019050610b3d60008301866109a6565b610b4a60208301856109a6565b610b576040830184610ae7565b949350505050565b6000608082019050610b7460008301886109a6565b610b8160208301876109a6565b610b8e6040830186610ae7565b8181036060830152610ba18184866109b5565b90509695505050505050565b6000604082019050610bc260008301856109a6565b610bcf6020830184610ae7565b9392505050565b6000602082019050610beb6000830184610a13565b92915050565b60006020820190508181036000830152610c0b8184610a22565b905092915050565b60006020820190508181036000830152610c2c81610a5b565b9050919050565b60006020820190508181036000830152610c4c81610a7e565b9050919050565b60006020820190508181036000830152610c6c81610aa1565b9050919050565b60006020820190508181036000830152610c8c81610ac4565b9050919050565b6000602082019050610ca86000830184610ae7565b92915050565b6000604082019050610cc36000830186610ae7565b8181036020830152610cd68184866109b5565b9050949350505050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610d2e82610d4b565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610d8082610d87565b9050919050565b6000610d9282610d99565b9050919050565b6000610da482610d4b565b9050919050565b82818337600083830152505050565b60005b83811015610dd8578082015181840152602081019050610dbd565b83811115610de7576000848401525b50505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b610f1081610d23565b8114610f1b57600080fd5b50565b610f2781610d35565b8114610f3257600080fd5b50565b610f3e81610d41565b8114610f4957600080fd5b50565b610f5581610d6b565b8114610f6057600080fd5b5056fea264697066735822122091075d1eeef27fe44b0eaedb3337a61ceaeefaf51461670f2fbf1bc67c152e6064736f6c63430008070033",
}

// ZetaConnectorNewEthABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNewEthMetaData.ABI instead.
var ZetaConnectorNewEthABI = ZetaConnectorNewEthMetaData.ABI

// ZetaConnectorNewEthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNewEthMetaData.Bin instead.
var ZetaConnectorNewEthBin = ZetaConnectorNewEthMetaData.Bin

// DeployZetaConnectorNewEth deploys a new Ethereum contract, binding an instance of ZetaConnectorNewEth to it.
func DeployZetaConnectorNewEth(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address, _zetaToken common.Address) (common.Address, *types.Transaction, *ZetaConnectorNewEth, error) {
	parsed, err := ZetaConnectorNewEthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNewEthBin), backend, _gateway, _zetaToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNewEth{ZetaConnectorNewEthCaller: ZetaConnectorNewEthCaller{contract: contract}, ZetaConnectorNewEthTransactor: ZetaConnectorNewEthTransactor{contract: contract}, ZetaConnectorNewEthFilterer: ZetaConnectorNewEthFilterer{contract: contract}}, nil
}

// ZetaConnectorNewEth is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNewEth struct {
	ZetaConnectorNewEthCaller     // Read-only binding to the contract
	ZetaConnectorNewEthTransactor // Write-only binding to the contract
	ZetaConnectorNewEthFilterer   // Log filterer for contract events
}

// ZetaConnectorNewEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNewEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNewEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNewEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNewEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNewEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNewEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNewEthSession struct {
	Contract     *ZetaConnectorNewEth // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZetaConnectorNewEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNewEthCallerSession struct {
	Contract *ZetaConnectorNewEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ZetaConnectorNewEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNewEthTransactorSession struct {
	Contract     *ZetaConnectorNewEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ZetaConnectorNewEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNewEthRaw struct {
	Contract *ZetaConnectorNewEth // Generic contract binding to access the raw methods on
}

// ZetaConnectorNewEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNewEthCallerRaw struct {
	Contract *ZetaConnectorNewEthCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNewEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNewEthTransactorRaw struct {
	Contract *ZetaConnectorNewEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNewEth creates a new instance of ZetaConnectorNewEth, bound to a specific deployed contract.
func NewZetaConnectorNewEth(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNewEth, error) {
	contract, err := bindZetaConnectorNewEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewEth{ZetaConnectorNewEthCaller: ZetaConnectorNewEthCaller{contract: contract}, ZetaConnectorNewEthTransactor: ZetaConnectorNewEthTransactor{contract: contract}, ZetaConnectorNewEthFilterer: ZetaConnectorNewEthFilterer{contract: contract}}, nil
}

// NewZetaConnectorNewEthCaller creates a new read-only instance of ZetaConnectorNewEth, bound to a specific deployed contract.
func NewZetaConnectorNewEthCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNewEthCaller, error) {
	contract, err := bindZetaConnectorNewEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewEthCaller{contract: contract}, nil
}

// NewZetaConnectorNewEthTransactor creates a new write-only instance of ZetaConnectorNewEth, bound to a specific deployed contract.
func NewZetaConnectorNewEthTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNewEthTransactor, error) {
	contract, err := bindZetaConnectorNewEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewEthTransactor{contract: contract}, nil
}

// NewZetaConnectorNewEthFilterer creates a new log filterer instance of ZetaConnectorNewEth, bound to a specific deployed contract.
func NewZetaConnectorNewEthFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNewEthFilterer, error) {
	contract, err := bindZetaConnectorNewEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewEthFilterer{contract: contract}, nil
}

// bindZetaConnectorNewEth binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNewEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNewEthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNewEth *ZetaConnectorNewEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNewEth.Contract.ZetaConnectorNewEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNewEth *ZetaConnectorNewEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.Contract.ZetaConnectorNewEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNewEth *ZetaConnectorNewEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.Contract.ZetaConnectorNewEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNewEth *ZetaConnectorNewEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNewEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNewEth *ZetaConnectorNewEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNewEth *ZetaConnectorNewEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNewEth *ZetaConnectorNewEthCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNewEth.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNewEth *ZetaConnectorNewEthSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNewEth.Contract.Gateway(&_ZetaConnectorNewEth.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNewEth *ZetaConnectorNewEthCallerSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNewEth.Contract.Gateway(&_ZetaConnectorNewEth.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNewEth *ZetaConnectorNewEthCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNewEth.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNewEth *ZetaConnectorNewEthSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNewEth.Contract.ZetaToken(&_ZetaConnectorNewEth.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNewEth *ZetaConnectorNewEthCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNewEth.Contract.ZetaToken(&_ZetaConnectorNewEth.CallOpts)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNewEth *ZetaConnectorNewEthTransactor) ReceiveTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.contract.Transact(opts, "receiveTokens", amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNewEth *ZetaConnectorNewEthSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.Contract.ReceiveTokens(&_ZetaConnectorNewEth.TransactOpts, amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNewEth *ZetaConnectorNewEthTransactorSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.Contract.ReceiveTokens(&_ZetaConnectorNewEth.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewEth *ZetaConnectorNewEthTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.contract.Transact(opts, "withdraw", to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewEth *ZetaConnectorNewEthSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.Contract.Withdraw(&_ZetaConnectorNewEth.TransactOpts, to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewEth *ZetaConnectorNewEthTransactorSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.Contract.Withdraw(&_ZetaConnectorNewEth.TransactOpts, to, amount, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewEth *ZetaConnectorNewEthTransactor) WithdrawAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.contract.Transact(opts, "withdrawAndCall", to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewEth *ZetaConnectorNewEthSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.Contract.WithdrawAndCall(&_ZetaConnectorNewEth.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewEth *ZetaConnectorNewEthTransactorSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewEth.Contract.WithdrawAndCall(&_ZetaConnectorNewEth.TransactOpts, to, amount, data, internalSendHash)
}

// ZetaConnectorNewEthWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ZetaConnectorNewEth contract.
type ZetaConnectorNewEthWithdrawIterator struct {
	Event *ZetaConnectorNewEthWithdraw // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNewEthWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNewEthWithdraw)
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
		it.Event = new(ZetaConnectorNewEthWithdraw)
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
func (it *ZetaConnectorNewEthWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNewEthWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNewEthWithdraw represents a Withdraw event raised by the ZetaConnectorNewEth contract.
type ZetaConnectorNewEthWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNewEth *ZetaConnectorNewEthFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNewEthWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewEth.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewEthWithdrawIterator{contract: _ZetaConnectorNewEth.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNewEth *ZetaConnectorNewEthFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNewEthWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewEth.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNewEthWithdraw)
				if err := _ZetaConnectorNewEth.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_ZetaConnectorNewEth *ZetaConnectorNewEthFilterer) ParseWithdraw(log types.Log) (*ZetaConnectorNewEthWithdraw, error) {
	event := new(ZetaConnectorNewEthWithdraw)
	if err := _ZetaConnectorNewEth.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNewEthWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the ZetaConnectorNewEth contract.
type ZetaConnectorNewEthWithdrawAndCallIterator struct {
	Event *ZetaConnectorNewEthWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNewEthWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNewEthWithdrawAndCall)
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
		it.Event = new(ZetaConnectorNewEthWithdrawAndCall)
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
func (it *ZetaConnectorNewEthWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNewEthWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNewEthWithdrawAndCall represents a WithdrawAndCall event raised by the ZetaConnectorNewEth contract.
type ZetaConnectorNewEthWithdrawAndCall struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNewEth *ZetaConnectorNewEthFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNewEthWithdrawAndCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewEth.contract.FilterLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewEthWithdrawAndCallIterator{contract: _ZetaConnectorNewEth.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNewEth *ZetaConnectorNewEthFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNewEthWithdrawAndCall, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewEth.contract.WatchLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNewEthWithdrawAndCall)
				if err := _ZetaConnectorNewEth.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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
func (_ZetaConnectorNewEth *ZetaConnectorNewEthFilterer) ParseWithdrawAndCall(log types.Log) (*ZetaConnectorNewEthWithdrawAndCall, error) {
	event := new(ZetaConnectorNewEthWithdrawAndCall)
	if err := _ZetaConnectorNewEth.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
