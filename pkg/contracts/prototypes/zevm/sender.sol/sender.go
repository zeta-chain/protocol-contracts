// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sender

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

// SenderMetaData contains all meta data concerning the Sender contract.
var SenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"callReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"withdrawAndCallReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610b98380380610b988339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610108565b600081519050610087816100f1565b92915050565b6000602082840312156100a3576100a26100ec565b5b60006100b184828501610078565b91505092915050565b60006100c5826100cc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6100fa816100ba565b811461010557600080fd5b50565b610a81806101176000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b614610062578063a0a1730b14610080575b600080fd5b610060600480360381019061005b91906105c8565b61009c565b005b61006a61027a565b604051610077919061072c565b60405180910390f35b61009a60048036038101906100959190610529565b61029e565b005b60008383836040516024016100b3939291906107fa565b6040516020818303038152906040527f6fa220ad000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090508473ffffffffffffffffffffffffffffffffffffffff1663095ea7b360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16886040518363ffffffff1660e01b815260040161018d929190610747565b602060405180830381600087803b1580156101a757600080fd5b505af11580156101bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101df91906104fc565b5060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637993c1e0888888856040518563ffffffff1660e01b815260040161023f94939291906107a7565b600060405180830381600087803b15801561025957600080fd5b505af115801561026d573d6000803e3d6000fd5b5050505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008383836040516024016102b5939291906107fa565b6040516020818303038152906040527f6fa220ad000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630ac7c44c86836040518363ffffffff1660e01b815260040161038f929190610770565b600060405180830381600087803b1580156103a957600080fd5b505af11580156103bd573d6000803e3d6000fd5b505050505050505050565b60006103db6103d68461085d565b610838565b9050828152602081018484840111156103f7576103f66109e6565b5b61040284828561093f565b509392505050565b600061041d6104188461088e565b610838565b905082815260208101848484011115610439576104386109e6565b5b61044484828561093f565b509392505050565b60008135905061045b81610a06565b92915050565b60008135905061047081610a1d565b92915050565b60008151905061048581610a1d565b92915050565b600082601f8301126104a05761049f6109e1565b5b81356104b08482602086016103c8565b91505092915050565b600082601f8301126104ce576104cd6109e1565b5b81356104de84826020860161040a565b91505092915050565b6000813590506104f681610a34565b92915050565b600060208284031215610512576105116109f0565b5b600061052084828501610476565b91505092915050565b60008060008060808587031215610543576105426109f0565b5b600085013567ffffffffffffffff811115610561576105606109eb565b5b61056d8782880161048b565b945050602085013567ffffffffffffffff81111561058e5761058d6109eb565b5b61059a878288016104b9565b93505060406105ab878288016104e7565b92505060606105bc87828801610461565b91505092959194509250565b60008060008060008060c087890312156105e5576105e46109f0565b5b600087013567ffffffffffffffff811115610603576106026109eb565b5b61060f89828a0161048b565b965050602061062089828a016104e7565b955050604061063189828a0161044c565b945050606087013567ffffffffffffffff811115610652576106516109eb565b5b61065e89828a016104b9565b935050608061066f89828a016104e7565b92505060a061068089828a01610461565b9150509295509295509295565b610696816108f7565b82525050565b6106a581610909565b82525050565b60006106b6826108bf565b6106c081856108d5565b93506106d081856020860161094e565b6106d9816109f5565b840191505092915050565b60006106ef826108ca565b6106f981856108e6565b935061070981856020860161094e565b610712816109f5565b840191505092915050565b61072681610935565b82525050565b6000602082019050610741600083018461068d565b92915050565b600060408201905061075c600083018561068d565b610769602083018461071d565b9392505050565b6000604082019050818103600083015261078a81856106ab565b9050818103602083015261079e81846106ab565b90509392505050565b600060808201905081810360008301526107c181876106ab565b90506107d0602083018661071d565b6107dd604083018561068d565b81810360608301526107ef81846106ab565b905095945050505050565b6000606082019050818103600083015261081481866106e4565b9050610823602083018561071d565b610830604083018461069c565b949350505050565b6000610842610853565b905061084e8282610981565b919050565b6000604051905090565b600067ffffffffffffffff821115610878576108776109b2565b5b610881826109f5565b9050602081019050919050565b600067ffffffffffffffff8211156108a9576108a86109b2565b5b6108b2826109f5565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600061090282610915565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561096c578082015181840152602081019050610951565b8381111561097b576000848401525b50505050565b61098a826109f5565b810181811067ffffffffffffffff821117156109a9576109a86109b2565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b610a0f816108f7565b8114610a1a57600080fd5b50565b610a2681610909565b8114610a3157600080fd5b50565b610a3d81610935565b8114610a4857600080fd5b5056fea26469706673582212205ca1aa6c38923357ab6a03800c93961e3f1fa3d8f0a80b31f28f02206204b7f364736f6c63430008070033",
}

// SenderABI is the input ABI used to generate the binding from.
// Deprecated: Use SenderMetaData.ABI instead.
var SenderABI = SenderMetaData.ABI

// SenderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SenderMetaData.Bin instead.
var SenderBin = SenderMetaData.Bin

// DeploySender deploys a new Ethereum contract, binding an instance of Sender to it.
func DeploySender(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address) (common.Address, *types.Transaction, *Sender, error) {
	parsed, err := SenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenderBin), backend, _gateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sender{SenderCaller: SenderCaller{contract: contract}, SenderTransactor: SenderTransactor{contract: contract}, SenderFilterer: SenderFilterer{contract: contract}}, nil
}

// Sender is an auto generated Go binding around an Ethereum contract.
type Sender struct {
	SenderCaller     // Read-only binding to the contract
	SenderTransactor // Write-only binding to the contract
	SenderFilterer   // Log filterer for contract events
}

// SenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenderSession struct {
	Contract     *Sender           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenderCallerSession struct {
	Contract *SenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenderTransactorSession struct {
	Contract     *SenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenderRaw struct {
	Contract *Sender // Generic contract binding to access the raw methods on
}

// SenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenderCallerRaw struct {
	Contract *SenderCaller // Generic read-only contract binding to access the raw methods on
}

// SenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenderTransactorRaw struct {
	Contract *SenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSender creates a new instance of Sender, bound to a specific deployed contract.
func NewSender(address common.Address, backend bind.ContractBackend) (*Sender, error) {
	contract, err := bindSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sender{SenderCaller: SenderCaller{contract: contract}, SenderTransactor: SenderTransactor{contract: contract}, SenderFilterer: SenderFilterer{contract: contract}}, nil
}

// NewSenderCaller creates a new read-only instance of Sender, bound to a specific deployed contract.
func NewSenderCaller(address common.Address, caller bind.ContractCaller) (*SenderCaller, error) {
	contract, err := bindSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenderCaller{contract: contract}, nil
}

// NewSenderTransactor creates a new write-only instance of Sender, bound to a specific deployed contract.
func NewSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*SenderTransactor, error) {
	contract, err := bindSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenderTransactor{contract: contract}, nil
}

// NewSenderFilterer creates a new log filterer instance of Sender, bound to a specific deployed contract.
func NewSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*SenderFilterer, error) {
	contract, err := bindSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenderFilterer{contract: contract}, nil
}

// bindSender binds a generic wrapper to an already deployed contract.
func bindSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SenderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sender *SenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sender.Contract.SenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sender *SenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sender.Contract.SenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sender *SenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sender.Contract.SenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sender *SenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sender *SenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sender *SenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sender.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Sender *SenderCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sender.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Sender *SenderSession) Gateway() (common.Address, error) {
	return _Sender.Contract.Gateway(&_Sender.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Sender *SenderCallerSession) Gateway() (common.Address, error) {
	return _Sender.Contract.Gateway(&_Sender.CallOpts)
}

// CallReceiver is a paid mutator transaction binding the contract method 0xa0a1730b.
//
// Solidity: function callReceiver(bytes receiver, string str, uint256 num, bool flag) returns()
func (_Sender *SenderTransactor) CallReceiver(opts *bind.TransactOpts, receiver []byte, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _Sender.contract.Transact(opts, "callReceiver", receiver, str, num, flag)
}

// CallReceiver is a paid mutator transaction binding the contract method 0xa0a1730b.
//
// Solidity: function callReceiver(bytes receiver, string str, uint256 num, bool flag) returns()
func (_Sender *SenderSession) CallReceiver(receiver []byte, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _Sender.Contract.CallReceiver(&_Sender.TransactOpts, receiver, str, num, flag)
}

// CallReceiver is a paid mutator transaction binding the contract method 0xa0a1730b.
//
// Solidity: function callReceiver(bytes receiver, string str, uint256 num, bool flag) returns()
func (_Sender *SenderTransactorSession) CallReceiver(receiver []byte, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _Sender.Contract.CallReceiver(&_Sender.TransactOpts, receiver, str, num, flag)
}

// WithdrawAndCallReceiver is a paid mutator transaction binding the contract method 0x0abd8905.
//
// Solidity: function withdrawAndCallReceiver(bytes receiver, uint256 amount, address zrc20, string str, uint256 num, bool flag) returns()
func (_Sender *SenderTransactor) WithdrawAndCallReceiver(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _Sender.contract.Transact(opts, "withdrawAndCallReceiver", receiver, amount, zrc20, str, num, flag)
}

// WithdrawAndCallReceiver is a paid mutator transaction binding the contract method 0x0abd8905.
//
// Solidity: function withdrawAndCallReceiver(bytes receiver, uint256 amount, address zrc20, string str, uint256 num, bool flag) returns()
func (_Sender *SenderSession) WithdrawAndCallReceiver(receiver []byte, amount *big.Int, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _Sender.Contract.WithdrawAndCallReceiver(&_Sender.TransactOpts, receiver, amount, zrc20, str, num, flag)
}

// WithdrawAndCallReceiver is a paid mutator transaction binding the contract method 0x0abd8905.
//
// Solidity: function withdrawAndCallReceiver(bytes receiver, uint256 amount, address zrc20, string str, uint256 num, bool flag) returns()
func (_Sender *SenderTransactorSession) WithdrawAndCallReceiver(receiver []byte, amount *big.Int, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _Sender.Contract.WithdrawAndCallReceiver(&_Sender.TransactOpts, receiver, amount, zrc20, str, num, flag)
}
