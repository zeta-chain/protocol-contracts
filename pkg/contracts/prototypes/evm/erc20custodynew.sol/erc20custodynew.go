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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIGatewayEVM\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200108b3803806200108b83398181016040528101906200003791906200009e565b600160008190555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000123565b600081519050620000988162000109565b92915050565b600060208284031215620000b757620000b662000104565b5b6000620000c78482850162000087565b91505092915050565b6000620000dd82620000e4565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200011481620000d0565b81146200012057600080fd5b50565b610f5880620001336000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063116191b61461004657806321fc65f214610064578063d9caed1214610080575b600080fd5b61004e61009c565b60405161005b9190610afb565b60405180910390f35b61007e6004803603810190610079919061081f565b6100c2565b005b61009a600480360381019061009591906107cc565b6102ad565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6100ca610352565b8473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518363ffffffff1660e01b8152600401610127929190610ad2565b602060405180830381600087803b15801561014157600080fd5b505af1158015610155573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061017991906108a7565b50600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635131ab5986868686866040518663ffffffff1660e01b81526004016101dd959493929190610a84565b600060405180830381600087803b1580156101f757600080fd5b505af115801561020b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061023491906108d4565b508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e85858560405161029693929190610bd3565b60405180910390a36102a66103a2565b5050505050565b6102b5610352565b6102e082828573ffffffffffffffffffffffffffffffffffffffff166103ac9092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb8360405161033d9190610bb8565b60405180910390a361034d6103a2565b505050565b60026000541415610398576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038f90610b98565b60405180910390fd5b6002600081905550565b6001600081905550565b61042d8363a9059cbb60e01b84846040516024016103cb929190610ad2565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610432565b505050565b6000610494826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166104f99092919063ffffffff16565b90506000815111156104f457808060200190518101906104b491906108a7565b6104f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ea90610b78565b60405180910390fd5b5b505050565b60606105088484600085610511565b90509392505050565b606082471015610556576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054d90610b38565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161057f9190610a6d565b60006040518083038185875af1925050503d80600081146105bc576040519150601f19603f3d011682016040523d82523d6000602084013e6105c1565b606091505b50915091506105d2878383876105de565b92505050949350505050565b6060831561064157600083511415610639576105f985610654565b610638576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062f90610b58565b60405180910390fd5b5b82905061064c565b61064b8383610677565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60008251111561068a5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106be9190610b16565b60405180910390fd5b60006106da6106d584610c2a565b610c05565b9050828152602081018484840111156106f6576106f5610dcd565b5b610701848285610d2b565b509392505050565b60008135905061071881610edd565b92915050565b60008151905061072d81610ef4565b92915050565b60008083601f84011261074957610748610dc3565b5b8235905067ffffffffffffffff81111561076657610765610dbe565b5b60208301915083600182028301111561078257610781610dc8565b5b9250929050565b600082601f83011261079e5761079d610dc3565b5b81516107ae8482602086016106c7565b91505092915050565b6000813590506107c681610f0b565b92915050565b6000806000606084860312156107e5576107e4610dd7565b5b60006107f386828701610709565b935050602061080486828701610709565b9250506040610815868287016107b7565b9150509250925092565b60008060008060006080868803121561083b5761083a610dd7565b5b600061084988828901610709565b955050602061085a88828901610709565b945050604061086b888289016107b7565b935050606086013567ffffffffffffffff81111561088c5761088b610dd2565b5b61089888828901610733565b92509250509295509295909350565b6000602082840312156108bd576108bc610dd7565b5b60006108cb8482850161071e565b91505092915050565b6000602082840312156108ea576108e9610dd7565b5b600082015167ffffffffffffffff81111561090857610907610dd2565b5b61091484828501610789565b91505092915050565b61092681610c9e565b82525050565b60006109388385610c71565b9350610945838584610d1c565b61094e83610ddc565b840190509392505050565b600061096482610c5b565b61096e8185610c82565b935061097e818560208601610d2b565b80840191505092915050565b61099381610ce6565b82525050565b60006109a482610c66565b6109ae8185610c8d565b93506109be818560208601610d2b565b6109c781610ddc565b840191505092915050565b60006109df602683610c8d565b91506109ea82610ded565b604082019050919050565b6000610a02601d83610c8d565b9150610a0d82610e3c565b602082019050919050565b6000610a25602a83610c8d565b9150610a3082610e65565b604082019050919050565b6000610a48601f83610c8d565b9150610a5382610eb4565b602082019050919050565b610a6781610cdc565b82525050565b6000610a798284610959565b915081905092915050565b6000608082019050610a99600083018861091d565b610aa6602083018761091d565b610ab36040830186610a5e565b8181036060830152610ac681848661092c565b90509695505050505050565b6000604082019050610ae7600083018561091d565b610af46020830184610a5e565b9392505050565b6000602082019050610b10600083018461098a565b92915050565b60006020820190508181036000830152610b308184610999565b905092915050565b60006020820190508181036000830152610b51816109d2565b9050919050565b60006020820190508181036000830152610b71816109f5565b9050919050565b60006020820190508181036000830152610b9181610a18565b9050919050565b60006020820190508181036000830152610bb181610a3b565b9050919050565b6000602082019050610bcd6000830184610a5e565b92915050565b6000604082019050610be86000830186610a5e565b8181036020830152610bfb81848661092c565b9050949350505050565b6000610c0f610c20565b9050610c1b8282610d5e565b919050565b6000604051905090565b600067ffffffffffffffff821115610c4557610c44610d8f565b5b610c4e82610ddc565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610ca982610cbc565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610cf182610cf8565b9050919050565b6000610d0382610d0a565b9050919050565b6000610d1582610cbc565b9050919050565b82818337600083830152505050565b60005b83811015610d49578082015181840152602081019050610d2e565b83811115610d58576000848401525b50505050565b610d6782610ddc565b810181811067ffffffffffffffff82111715610d8657610d85610d8f565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b610ee681610c9e565b8114610ef157600080fd5b50565b610efd81610cb0565b8114610f0857600080fd5b50565b610f1481610cdc565b8114610f1f57600080fd5b5056fea2646970667358221220a59bf09ff2b63c7a4f1f46f0ceb21bc24c06e7fe90fe67b94cfbac6fb195e97464736f6c63430008070033",
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
