// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izetaregistry

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

// IZetaRegistryChainInfo is an auto generated low-level Go binding around an user-defined struct.
type IZetaRegistryChainInfo struct {
	ChainId            *big.Int
	Name               string
	Active             bool
	ChainType          uint8
	BlockConfirmations uint64
}

// IZetaRegistryContractInfo is an auto generated low-level Go binding around an user-defined struct.
type IZetaRegistryContractInfo struct {
	Active            bool
	Addr              common.Address
	Version           string
	Implementation    common.Address
	ConfigurationData []byte
}

// IZetaRegistryMetaData contains all meta data concerning the IZetaRegistry contract.
var IZetaRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addChain\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainInfo\",\"type\":\"tuple\",\"internalType\":\"structIZetaRegistry.ChainInfo\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockConfirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addContract\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"category\",\"type\":\"uint8\",\"internalType\":\"enumIZetaRegistry.ContractCategory\"},{\"name\":\"contractInfo\",\"type\":\"tuple\",\"internalType\":\"structIZetaRegistry.ContractInfo\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"configurationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addZRC20Token\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchAddContracts\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifiers\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"categories\",\"type\":\"uint8[]\",\"internalType\":\"enumIZetaRegistry.ContractCategory[]\"},{\"name\":\"contractInfos\",\"type\":\"tuple[]\",\"internalType\":\"structIZetaRegistry.ContractInfo[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"configurationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chainExists\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractExists\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createChainIdentifier\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"identifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"createContractIdentifier\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"identifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getActiveChainIdentifiers\",\"inputs\":[],\"outputs\":[{\"name\":\"activeChainIdentifiers\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllChainIdentifiers\",\"inputs\":[],\"outputs\":[{\"name\":\"chainIdentifiers\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainCount\",\"inputs\":[],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainInfo\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"chainInfo\",\"type\":\"tuple\",\"internalType\":\"structIZetaRegistry.ChainInfo\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockConfirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractAddress\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractIdentifiers\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"contractIdentifiers\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractIdentifiersByCategory\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"category\",\"type\":\"uint8\",\"internalType\":\"enumIZetaRegistry.ContractCategory\"}],\"outputs\":[{\"name\":\"contractIdentifiers\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractInfo\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"contractInfo\",\"type\":\"tuple\",\"internalType\":\"structIZetaRegistry.ContractInfo\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"configurationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSystemAddresses\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"connector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"erc20Custody\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tss\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssUpdater\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zetaToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getZRC20Address\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeChain\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeContract\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainActive\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractActive\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateChain\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainInfo\",\"type\":\"tuple\",\"internalType\":\"structIZetaRegistry.ChainInfo\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockConfirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateContract\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractInfo\",\"type\":\"tuple\",\"internalType\":\"structIZetaRegistry.ContractInfo\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"configurationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateContractAddress\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateContractConfiguration\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"configurationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainStatusUpdated\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"active\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractAddressUpdated\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"category\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractConfigurationUpdated\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"configurationData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractStatusUpdated\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"active\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenAdded\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"tokenIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ChainDoesNotExist\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ContractAlreadyExists\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ContractDoesNotExist\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidContractAddress\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZRC20TokenAlreadyExists\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZRC20TokenDoesNotExist\",\"inputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// IZetaRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IZetaRegistryMetaData.ABI instead.
var IZetaRegistryABI = IZetaRegistryMetaData.ABI

// IZetaRegistry is an auto generated Go binding around an Ethereum contract.
type IZetaRegistry struct {
	IZetaRegistryCaller     // Read-only binding to the contract
	IZetaRegistryTransactor // Write-only binding to the contract
	IZetaRegistryFilterer   // Log filterer for contract events
}

// IZetaRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZetaRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZetaRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZetaRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZetaRegistrySession struct {
	Contract     *IZetaRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZetaRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZetaRegistryCallerSession struct {
	Contract *IZetaRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IZetaRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZetaRegistryTransactorSession struct {
	Contract     *IZetaRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IZetaRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZetaRegistryRaw struct {
	Contract *IZetaRegistry // Generic contract binding to access the raw methods on
}

// IZetaRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZetaRegistryCallerRaw struct {
	Contract *IZetaRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IZetaRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZetaRegistryTransactorRaw struct {
	Contract *IZetaRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZetaRegistry creates a new instance of IZetaRegistry, bound to a specific deployed contract.
func NewIZetaRegistry(address common.Address, backend bind.ContractBackend) (*IZetaRegistry, error) {
	contract, err := bindIZetaRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistry{IZetaRegistryCaller: IZetaRegistryCaller{contract: contract}, IZetaRegistryTransactor: IZetaRegistryTransactor{contract: contract}, IZetaRegistryFilterer: IZetaRegistryFilterer{contract: contract}}, nil
}

// NewIZetaRegistryCaller creates a new read-only instance of IZetaRegistry, bound to a specific deployed contract.
func NewIZetaRegistryCaller(address common.Address, caller bind.ContractCaller) (*IZetaRegistryCaller, error) {
	contract, err := bindIZetaRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryCaller{contract: contract}, nil
}

// NewIZetaRegistryTransactor creates a new write-only instance of IZetaRegistry, bound to a specific deployed contract.
func NewIZetaRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IZetaRegistryTransactor, error) {
	contract, err := bindIZetaRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryTransactor{contract: contract}, nil
}

// NewIZetaRegistryFilterer creates a new log filterer instance of IZetaRegistry, bound to a specific deployed contract.
func NewIZetaRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IZetaRegistryFilterer, error) {
	contract, err := bindIZetaRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryFilterer{contract: contract}, nil
}

// bindIZetaRegistry binds a generic wrapper to an already deployed contract.
func bindIZetaRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZetaRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaRegistry *IZetaRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaRegistry.Contract.IZetaRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaRegistry *IZetaRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.IZetaRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaRegistry *IZetaRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.IZetaRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaRegistry *IZetaRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaRegistry *IZetaRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaRegistry *IZetaRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.contract.Transact(opts, method, params...)
}

// ChainExists is a free data retrieval call binding the contract method 0xca804ccd.
//
// Solidity: function chainExists(bytes32 chainIdentifier) view returns(bool exists)
func (_IZetaRegistry *IZetaRegistryCaller) ChainExists(opts *bind.CallOpts, chainIdentifier [32]byte) (bool, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "chainExists", chainIdentifier)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChainExists is a free data retrieval call binding the contract method 0xca804ccd.
//
// Solidity: function chainExists(bytes32 chainIdentifier) view returns(bool exists)
func (_IZetaRegistry *IZetaRegistrySession) ChainExists(chainIdentifier [32]byte) (bool, error) {
	return _IZetaRegistry.Contract.ChainExists(&_IZetaRegistry.CallOpts, chainIdentifier)
}

// ChainExists is a free data retrieval call binding the contract method 0xca804ccd.
//
// Solidity: function chainExists(bytes32 chainIdentifier) view returns(bool exists)
func (_IZetaRegistry *IZetaRegistryCallerSession) ChainExists(chainIdentifier [32]byte) (bool, error) {
	return _IZetaRegistry.Contract.ChainExists(&_IZetaRegistry.CallOpts, chainIdentifier)
}

// ContractExists is a free data retrieval call binding the contract method 0x681de405.
//
// Solidity: function contractExists(bytes32 chainIdentifier, bytes32 contractIdentifier) view returns(bool exists)
func (_IZetaRegistry *IZetaRegistryCaller) ContractExists(opts *bind.CallOpts, chainIdentifier [32]byte, contractIdentifier [32]byte) (bool, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "contractExists", chainIdentifier, contractIdentifier)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractExists is a free data retrieval call binding the contract method 0x681de405.
//
// Solidity: function contractExists(bytes32 chainIdentifier, bytes32 contractIdentifier) view returns(bool exists)
func (_IZetaRegistry *IZetaRegistrySession) ContractExists(chainIdentifier [32]byte, contractIdentifier [32]byte) (bool, error) {
	return _IZetaRegistry.Contract.ContractExists(&_IZetaRegistry.CallOpts, chainIdentifier, contractIdentifier)
}

// ContractExists is a free data retrieval call binding the contract method 0x681de405.
//
// Solidity: function contractExists(bytes32 chainIdentifier, bytes32 contractIdentifier) view returns(bool exists)
func (_IZetaRegistry *IZetaRegistryCallerSession) ContractExists(chainIdentifier [32]byte, contractIdentifier [32]byte) (bool, error) {
	return _IZetaRegistry.Contract.ContractExists(&_IZetaRegistry.CallOpts, chainIdentifier, contractIdentifier)
}

// CreateChainIdentifier is a free data retrieval call binding the contract method 0xfa228515.
//
// Solidity: function createChainIdentifier(string name) pure returns(bytes32 identifier)
func (_IZetaRegistry *IZetaRegistryCaller) CreateChainIdentifier(opts *bind.CallOpts, name string) ([32]byte, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "createChainIdentifier", name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CreateChainIdentifier is a free data retrieval call binding the contract method 0xfa228515.
//
// Solidity: function createChainIdentifier(string name) pure returns(bytes32 identifier)
func (_IZetaRegistry *IZetaRegistrySession) CreateChainIdentifier(name string) ([32]byte, error) {
	return _IZetaRegistry.Contract.CreateChainIdentifier(&_IZetaRegistry.CallOpts, name)
}

// CreateChainIdentifier is a free data retrieval call binding the contract method 0xfa228515.
//
// Solidity: function createChainIdentifier(string name) pure returns(bytes32 identifier)
func (_IZetaRegistry *IZetaRegistryCallerSession) CreateChainIdentifier(name string) ([32]byte, error) {
	return _IZetaRegistry.Contract.CreateChainIdentifier(&_IZetaRegistry.CallOpts, name)
}

// CreateContractIdentifier is a free data retrieval call binding the contract method 0x2ff3e88a.
//
// Solidity: function createContractIdentifier(string name) pure returns(bytes32 identifier)
func (_IZetaRegistry *IZetaRegistryCaller) CreateContractIdentifier(opts *bind.CallOpts, name string) ([32]byte, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "createContractIdentifier", name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CreateContractIdentifier is a free data retrieval call binding the contract method 0x2ff3e88a.
//
// Solidity: function createContractIdentifier(string name) pure returns(bytes32 identifier)
func (_IZetaRegistry *IZetaRegistrySession) CreateContractIdentifier(name string) ([32]byte, error) {
	return _IZetaRegistry.Contract.CreateContractIdentifier(&_IZetaRegistry.CallOpts, name)
}

// CreateContractIdentifier is a free data retrieval call binding the contract method 0x2ff3e88a.
//
// Solidity: function createContractIdentifier(string name) pure returns(bytes32 identifier)
func (_IZetaRegistry *IZetaRegistryCallerSession) CreateContractIdentifier(name string) ([32]byte, error) {
	return _IZetaRegistry.Contract.CreateContractIdentifier(&_IZetaRegistry.CallOpts, name)
}

// GetActiveChainIdentifiers is a free data retrieval call binding the contract method 0xd11dd58f.
//
// Solidity: function getActiveChainIdentifiers() view returns(bytes32[] activeChainIdentifiers)
func (_IZetaRegistry *IZetaRegistryCaller) GetActiveChainIdentifiers(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "getActiveChainIdentifiers")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetActiveChainIdentifiers is a free data retrieval call binding the contract method 0xd11dd58f.
//
// Solidity: function getActiveChainIdentifiers() view returns(bytes32[] activeChainIdentifiers)
func (_IZetaRegistry *IZetaRegistrySession) GetActiveChainIdentifiers() ([][32]byte, error) {
	return _IZetaRegistry.Contract.GetActiveChainIdentifiers(&_IZetaRegistry.CallOpts)
}

// GetActiveChainIdentifiers is a free data retrieval call binding the contract method 0xd11dd58f.
//
// Solidity: function getActiveChainIdentifiers() view returns(bytes32[] activeChainIdentifiers)
func (_IZetaRegistry *IZetaRegistryCallerSession) GetActiveChainIdentifiers() ([][32]byte, error) {
	return _IZetaRegistry.Contract.GetActiveChainIdentifiers(&_IZetaRegistry.CallOpts)
}

// GetAllChainIdentifiers is a free data retrieval call binding the contract method 0xf95312fa.
//
// Solidity: function getAllChainIdentifiers() view returns(bytes32[] chainIdentifiers)
func (_IZetaRegistry *IZetaRegistryCaller) GetAllChainIdentifiers(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "getAllChainIdentifiers")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllChainIdentifiers is a free data retrieval call binding the contract method 0xf95312fa.
//
// Solidity: function getAllChainIdentifiers() view returns(bytes32[] chainIdentifiers)
func (_IZetaRegistry *IZetaRegistrySession) GetAllChainIdentifiers() ([][32]byte, error) {
	return _IZetaRegistry.Contract.GetAllChainIdentifiers(&_IZetaRegistry.CallOpts)
}

// GetAllChainIdentifiers is a free data retrieval call binding the contract method 0xf95312fa.
//
// Solidity: function getAllChainIdentifiers() view returns(bytes32[] chainIdentifiers)
func (_IZetaRegistry *IZetaRegistryCallerSession) GetAllChainIdentifiers() ([][32]byte, error) {
	return _IZetaRegistry.Contract.GetAllChainIdentifiers(&_IZetaRegistry.CallOpts)
}

// GetChainCount is a free data retrieval call binding the contract method 0xd25d23bd.
//
// Solidity: function getChainCount() view returns(uint256 count)
func (_IZetaRegistry *IZetaRegistryCaller) GetChainCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "getChainCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainCount is a free data retrieval call binding the contract method 0xd25d23bd.
//
// Solidity: function getChainCount() view returns(uint256 count)
func (_IZetaRegistry *IZetaRegistrySession) GetChainCount() (*big.Int, error) {
	return _IZetaRegistry.Contract.GetChainCount(&_IZetaRegistry.CallOpts)
}

// GetChainCount is a free data retrieval call binding the contract method 0xd25d23bd.
//
// Solidity: function getChainCount() view returns(uint256 count)
func (_IZetaRegistry *IZetaRegistryCallerSession) GetChainCount() (*big.Int, error) {
	return _IZetaRegistry.Contract.GetChainCount(&_IZetaRegistry.CallOpts)
}

// GetChainInfo is a free data retrieval call binding the contract method 0x7e482f47.
//
// Solidity: function getChainInfo(bytes32 chainIdentifier) view returns((uint256,string,bool,uint8,uint64) chainInfo)
func (_IZetaRegistry *IZetaRegistryCaller) GetChainInfo(opts *bind.CallOpts, chainIdentifier [32]byte) (IZetaRegistryChainInfo, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "getChainInfo", chainIdentifier)

	if err != nil {
		return *new(IZetaRegistryChainInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IZetaRegistryChainInfo)).(*IZetaRegistryChainInfo)

	return out0, err

}

// GetChainInfo is a free data retrieval call binding the contract method 0x7e482f47.
//
// Solidity: function getChainInfo(bytes32 chainIdentifier) view returns((uint256,string,bool,uint8,uint64) chainInfo)
func (_IZetaRegistry *IZetaRegistrySession) GetChainInfo(chainIdentifier [32]byte) (IZetaRegistryChainInfo, error) {
	return _IZetaRegistry.Contract.GetChainInfo(&_IZetaRegistry.CallOpts, chainIdentifier)
}

// GetChainInfo is a free data retrieval call binding the contract method 0x7e482f47.
//
// Solidity: function getChainInfo(bytes32 chainIdentifier) view returns((uint256,string,bool,uint8,uint64) chainInfo)
func (_IZetaRegistry *IZetaRegistryCallerSession) GetChainInfo(chainIdentifier [32]byte) (IZetaRegistryChainInfo, error) {
	return _IZetaRegistry.Contract.GetChainInfo(&_IZetaRegistry.CallOpts, chainIdentifier)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x778c271d.
//
// Solidity: function getContractAddress(bytes32 chainIdentifier, bytes32 contractIdentifier) view returns(address addr)
func (_IZetaRegistry *IZetaRegistryCaller) GetContractAddress(opts *bind.CallOpts, chainIdentifier [32]byte, contractIdentifier [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "getContractAddress", chainIdentifier, contractIdentifier)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractAddress is a free data retrieval call binding the contract method 0x778c271d.
//
// Solidity: function getContractAddress(bytes32 chainIdentifier, bytes32 contractIdentifier) view returns(address addr)
func (_IZetaRegistry *IZetaRegistrySession) GetContractAddress(chainIdentifier [32]byte, contractIdentifier [32]byte) (common.Address, error) {
	return _IZetaRegistry.Contract.GetContractAddress(&_IZetaRegistry.CallOpts, chainIdentifier, contractIdentifier)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x778c271d.
//
// Solidity: function getContractAddress(bytes32 chainIdentifier, bytes32 contractIdentifier) view returns(address addr)
func (_IZetaRegistry *IZetaRegistryCallerSession) GetContractAddress(chainIdentifier [32]byte, contractIdentifier [32]byte) (common.Address, error) {
	return _IZetaRegistry.Contract.GetContractAddress(&_IZetaRegistry.CallOpts, chainIdentifier, contractIdentifier)
}

// GetContractIdentifiers is a free data retrieval call binding the contract method 0x9d9dcd8b.
//
// Solidity: function getContractIdentifiers(bytes32 chainIdentifier) view returns(bytes32[] contractIdentifiers)
func (_IZetaRegistry *IZetaRegistryCaller) GetContractIdentifiers(opts *bind.CallOpts, chainIdentifier [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "getContractIdentifiers", chainIdentifier)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetContractIdentifiers is a free data retrieval call binding the contract method 0x9d9dcd8b.
//
// Solidity: function getContractIdentifiers(bytes32 chainIdentifier) view returns(bytes32[] contractIdentifiers)
func (_IZetaRegistry *IZetaRegistrySession) GetContractIdentifiers(chainIdentifier [32]byte) ([][32]byte, error) {
	return _IZetaRegistry.Contract.GetContractIdentifiers(&_IZetaRegistry.CallOpts, chainIdentifier)
}

// GetContractIdentifiers is a free data retrieval call binding the contract method 0x9d9dcd8b.
//
// Solidity: function getContractIdentifiers(bytes32 chainIdentifier) view returns(bytes32[] contractIdentifiers)
func (_IZetaRegistry *IZetaRegistryCallerSession) GetContractIdentifiers(chainIdentifier [32]byte) ([][32]byte, error) {
	return _IZetaRegistry.Contract.GetContractIdentifiers(&_IZetaRegistry.CallOpts, chainIdentifier)
}

// GetContractIdentifiersByCategory is a free data retrieval call binding the contract method 0xb3a116ea.
//
// Solidity: function getContractIdentifiersByCategory(bytes32 chainIdentifier, uint8 category) view returns(bytes32[] contractIdentifiers)
func (_IZetaRegistry *IZetaRegistryCaller) GetContractIdentifiersByCategory(opts *bind.CallOpts, chainIdentifier [32]byte, category uint8) ([][32]byte, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "getContractIdentifiersByCategory", chainIdentifier, category)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetContractIdentifiersByCategory is a free data retrieval call binding the contract method 0xb3a116ea.
//
// Solidity: function getContractIdentifiersByCategory(bytes32 chainIdentifier, uint8 category) view returns(bytes32[] contractIdentifiers)
func (_IZetaRegistry *IZetaRegistrySession) GetContractIdentifiersByCategory(chainIdentifier [32]byte, category uint8) ([][32]byte, error) {
	return _IZetaRegistry.Contract.GetContractIdentifiersByCategory(&_IZetaRegistry.CallOpts, chainIdentifier, category)
}

// GetContractIdentifiersByCategory is a free data retrieval call binding the contract method 0xb3a116ea.
//
// Solidity: function getContractIdentifiersByCategory(bytes32 chainIdentifier, uint8 category) view returns(bytes32[] contractIdentifiers)
func (_IZetaRegistry *IZetaRegistryCallerSession) GetContractIdentifiersByCategory(chainIdentifier [32]byte, category uint8) ([][32]byte, error) {
	return _IZetaRegistry.Contract.GetContractIdentifiersByCategory(&_IZetaRegistry.CallOpts, chainIdentifier, category)
}

// GetContractInfo is a free data retrieval call binding the contract method 0xf42ed3f9.
//
// Solidity: function getContractInfo(bytes32 chainIdentifier, bytes32 contractIdentifier) view returns((bool,address,string,address,bytes) contractInfo)
func (_IZetaRegistry *IZetaRegistryCaller) GetContractInfo(opts *bind.CallOpts, chainIdentifier [32]byte, contractIdentifier [32]byte) (IZetaRegistryContractInfo, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "getContractInfo", chainIdentifier, contractIdentifier)

	if err != nil {
		return *new(IZetaRegistryContractInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IZetaRegistryContractInfo)).(*IZetaRegistryContractInfo)

	return out0, err

}

// GetContractInfo is a free data retrieval call binding the contract method 0xf42ed3f9.
//
// Solidity: function getContractInfo(bytes32 chainIdentifier, bytes32 contractIdentifier) view returns((bool,address,string,address,bytes) contractInfo)
func (_IZetaRegistry *IZetaRegistrySession) GetContractInfo(chainIdentifier [32]byte, contractIdentifier [32]byte) (IZetaRegistryContractInfo, error) {
	return _IZetaRegistry.Contract.GetContractInfo(&_IZetaRegistry.CallOpts, chainIdentifier, contractIdentifier)
}

// GetContractInfo is a free data retrieval call binding the contract method 0xf42ed3f9.
//
// Solidity: function getContractInfo(bytes32 chainIdentifier, bytes32 contractIdentifier) view returns((bool,address,string,address,bytes) contractInfo)
func (_IZetaRegistry *IZetaRegistryCallerSession) GetContractInfo(chainIdentifier [32]byte, contractIdentifier [32]byte) (IZetaRegistryContractInfo, error) {
	return _IZetaRegistry.Contract.GetContractInfo(&_IZetaRegistry.CallOpts, chainIdentifier, contractIdentifier)
}

// GetSystemAddresses is a free data retrieval call binding the contract method 0x5c5e0136.
//
// Solidity: function getSystemAddresses(bytes32 chainIdentifier) view returns(address connector, address erc20Custody, address tss, address tssUpdater, address zetaToken)
func (_IZetaRegistry *IZetaRegistryCaller) GetSystemAddresses(opts *bind.CallOpts, chainIdentifier [32]byte) (struct {
	Connector    common.Address
	Erc20Custody common.Address
	Tss          common.Address
	TssUpdater   common.Address
	ZetaToken    common.Address
}, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "getSystemAddresses", chainIdentifier)

	outstruct := new(struct {
		Connector    common.Address
		Erc20Custody common.Address
		Tss          common.Address
		TssUpdater   common.Address
		ZetaToken    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Connector = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Erc20Custody = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Tss = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TssUpdater = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ZetaToken = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetSystemAddresses is a free data retrieval call binding the contract method 0x5c5e0136.
//
// Solidity: function getSystemAddresses(bytes32 chainIdentifier) view returns(address connector, address erc20Custody, address tss, address tssUpdater, address zetaToken)
func (_IZetaRegistry *IZetaRegistrySession) GetSystemAddresses(chainIdentifier [32]byte) (struct {
	Connector    common.Address
	Erc20Custody common.Address
	Tss          common.Address
	TssUpdater   common.Address
	ZetaToken    common.Address
}, error) {
	return _IZetaRegistry.Contract.GetSystemAddresses(&_IZetaRegistry.CallOpts, chainIdentifier)
}

// GetSystemAddresses is a free data retrieval call binding the contract method 0x5c5e0136.
//
// Solidity: function getSystemAddresses(bytes32 chainIdentifier) view returns(address connector, address erc20Custody, address tss, address tssUpdater, address zetaToken)
func (_IZetaRegistry *IZetaRegistryCallerSession) GetSystemAddresses(chainIdentifier [32]byte) (struct {
	Connector    common.Address
	Erc20Custody common.Address
	Tss          common.Address
	TssUpdater   common.Address
	ZetaToken    common.Address
}, error) {
	return _IZetaRegistry.Contract.GetSystemAddresses(&_IZetaRegistry.CallOpts, chainIdentifier)
}

// GetZRC20Address is a free data retrieval call binding the contract method 0x340aba03.
//
// Solidity: function getZRC20Address(bytes32 chainIdentifier, bytes32 tokenIdentifier) view returns(address tokenAddress)
func (_IZetaRegistry *IZetaRegistryCaller) GetZRC20Address(opts *bind.CallOpts, chainIdentifier [32]byte, tokenIdentifier [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IZetaRegistry.contract.Call(opts, &out, "getZRC20Address", chainIdentifier, tokenIdentifier)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetZRC20Address is a free data retrieval call binding the contract method 0x340aba03.
//
// Solidity: function getZRC20Address(bytes32 chainIdentifier, bytes32 tokenIdentifier) view returns(address tokenAddress)
func (_IZetaRegistry *IZetaRegistrySession) GetZRC20Address(chainIdentifier [32]byte, tokenIdentifier [32]byte) (common.Address, error) {
	return _IZetaRegistry.Contract.GetZRC20Address(&_IZetaRegistry.CallOpts, chainIdentifier, tokenIdentifier)
}

// GetZRC20Address is a free data retrieval call binding the contract method 0x340aba03.
//
// Solidity: function getZRC20Address(bytes32 chainIdentifier, bytes32 tokenIdentifier) view returns(address tokenAddress)
func (_IZetaRegistry *IZetaRegistryCallerSession) GetZRC20Address(chainIdentifier [32]byte, tokenIdentifier [32]byte) (common.Address, error) {
	return _IZetaRegistry.Contract.GetZRC20Address(&_IZetaRegistry.CallOpts, chainIdentifier, tokenIdentifier)
}

// AddChain is a paid mutator transaction binding the contract method 0xd66c4f03.
//
// Solidity: function addChain(bytes32 chainIdentifier, (uint256,string,bool,uint8,uint64) chainInfo) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactor) AddChain(opts *bind.TransactOpts, chainIdentifier [32]byte, chainInfo IZetaRegistryChainInfo) (*types.Transaction, error) {
	return _IZetaRegistry.contract.Transact(opts, "addChain", chainIdentifier, chainInfo)
}

// AddChain is a paid mutator transaction binding the contract method 0xd66c4f03.
//
// Solidity: function addChain(bytes32 chainIdentifier, (uint256,string,bool,uint8,uint64) chainInfo) returns(bool success)
func (_IZetaRegistry *IZetaRegistrySession) AddChain(chainIdentifier [32]byte, chainInfo IZetaRegistryChainInfo) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.AddChain(&_IZetaRegistry.TransactOpts, chainIdentifier, chainInfo)
}

// AddChain is a paid mutator transaction binding the contract method 0xd66c4f03.
//
// Solidity: function addChain(bytes32 chainIdentifier, (uint256,string,bool,uint8,uint64) chainInfo) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactorSession) AddChain(chainIdentifier [32]byte, chainInfo IZetaRegistryChainInfo) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.AddChain(&_IZetaRegistry.TransactOpts, chainIdentifier, chainInfo)
}

// AddContract is a paid mutator transaction binding the contract method 0x2a0529a7.
//
// Solidity: function addContract(bytes32 chainIdentifier, bytes32 contractIdentifier, uint8 category, (bool,address,string,address,bytes) contractInfo) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactor) AddContract(opts *bind.TransactOpts, chainIdentifier [32]byte, contractIdentifier [32]byte, category uint8, contractInfo IZetaRegistryContractInfo) (*types.Transaction, error) {
	return _IZetaRegistry.contract.Transact(opts, "addContract", chainIdentifier, contractIdentifier, category, contractInfo)
}

// AddContract is a paid mutator transaction binding the contract method 0x2a0529a7.
//
// Solidity: function addContract(bytes32 chainIdentifier, bytes32 contractIdentifier, uint8 category, (bool,address,string,address,bytes) contractInfo) returns(bool success)
func (_IZetaRegistry *IZetaRegistrySession) AddContract(chainIdentifier [32]byte, contractIdentifier [32]byte, category uint8, contractInfo IZetaRegistryContractInfo) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.AddContract(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifier, category, contractInfo)
}

// AddContract is a paid mutator transaction binding the contract method 0x2a0529a7.
//
// Solidity: function addContract(bytes32 chainIdentifier, bytes32 contractIdentifier, uint8 category, (bool,address,string,address,bytes) contractInfo) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactorSession) AddContract(chainIdentifier [32]byte, contractIdentifier [32]byte, category uint8, contractInfo IZetaRegistryContractInfo) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.AddContract(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifier, category, contractInfo)
}

// AddZRC20Token is a paid mutator transaction binding the contract method 0x50f2e6b0.
//
// Solidity: function addZRC20Token(bytes32 chainIdentifier, bytes32 tokenIdentifier, address tokenAddress) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactor) AddZRC20Token(opts *bind.TransactOpts, chainIdentifier [32]byte, tokenIdentifier [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _IZetaRegistry.contract.Transact(opts, "addZRC20Token", chainIdentifier, tokenIdentifier, tokenAddress)
}

// AddZRC20Token is a paid mutator transaction binding the contract method 0x50f2e6b0.
//
// Solidity: function addZRC20Token(bytes32 chainIdentifier, bytes32 tokenIdentifier, address tokenAddress) returns(bool success)
func (_IZetaRegistry *IZetaRegistrySession) AddZRC20Token(chainIdentifier [32]byte, tokenIdentifier [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.AddZRC20Token(&_IZetaRegistry.TransactOpts, chainIdentifier, tokenIdentifier, tokenAddress)
}

// AddZRC20Token is a paid mutator transaction binding the contract method 0x50f2e6b0.
//
// Solidity: function addZRC20Token(bytes32 chainIdentifier, bytes32 tokenIdentifier, address tokenAddress) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactorSession) AddZRC20Token(chainIdentifier [32]byte, tokenIdentifier [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.AddZRC20Token(&_IZetaRegistry.TransactOpts, chainIdentifier, tokenIdentifier, tokenAddress)
}

// BatchAddContracts is a paid mutator transaction binding the contract method 0x6c6ccd64.
//
// Solidity: function batchAddContracts(bytes32 chainIdentifier, bytes32[] contractIdentifiers, uint8[] categories, (bool,address,string,address,bytes)[] contractInfos) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactor) BatchAddContracts(opts *bind.TransactOpts, chainIdentifier [32]byte, contractIdentifiers [][32]byte, categories []uint8, contractInfos []IZetaRegistryContractInfo) (*types.Transaction, error) {
	return _IZetaRegistry.contract.Transact(opts, "batchAddContracts", chainIdentifier, contractIdentifiers, categories, contractInfos)
}

// BatchAddContracts is a paid mutator transaction binding the contract method 0x6c6ccd64.
//
// Solidity: function batchAddContracts(bytes32 chainIdentifier, bytes32[] contractIdentifiers, uint8[] categories, (bool,address,string,address,bytes)[] contractInfos) returns(bool success)
func (_IZetaRegistry *IZetaRegistrySession) BatchAddContracts(chainIdentifier [32]byte, contractIdentifiers [][32]byte, categories []uint8, contractInfos []IZetaRegistryContractInfo) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.BatchAddContracts(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifiers, categories, contractInfos)
}

// BatchAddContracts is a paid mutator transaction binding the contract method 0x6c6ccd64.
//
// Solidity: function batchAddContracts(bytes32 chainIdentifier, bytes32[] contractIdentifiers, uint8[] categories, (bool,address,string,address,bytes)[] contractInfos) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactorSession) BatchAddContracts(chainIdentifier [32]byte, contractIdentifiers [][32]byte, categories []uint8, contractInfos []IZetaRegistryContractInfo) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.BatchAddContracts(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifiers, categories, contractInfos)
}

// RemoveChain is a paid mutator transaction binding the contract method 0x2ccde4b6.
//
// Solidity: function removeChain(bytes32 chainIdentifier) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactor) RemoveChain(opts *bind.TransactOpts, chainIdentifier [32]byte) (*types.Transaction, error) {
	return _IZetaRegistry.contract.Transact(opts, "removeChain", chainIdentifier)
}

// RemoveChain is a paid mutator transaction binding the contract method 0x2ccde4b6.
//
// Solidity: function removeChain(bytes32 chainIdentifier) returns(bool success)
func (_IZetaRegistry *IZetaRegistrySession) RemoveChain(chainIdentifier [32]byte) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.RemoveChain(&_IZetaRegistry.TransactOpts, chainIdentifier)
}

// RemoveChain is a paid mutator transaction binding the contract method 0x2ccde4b6.
//
// Solidity: function removeChain(bytes32 chainIdentifier) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactorSession) RemoveChain(chainIdentifier [32]byte) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.RemoveChain(&_IZetaRegistry.TransactOpts, chainIdentifier)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xd8f4b2bc.
//
// Solidity: function removeContract(bytes32 chainIdentifier, bytes32 contractIdentifier) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactor) RemoveContract(opts *bind.TransactOpts, chainIdentifier [32]byte, contractIdentifier [32]byte) (*types.Transaction, error) {
	return _IZetaRegistry.contract.Transact(opts, "removeContract", chainIdentifier, contractIdentifier)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xd8f4b2bc.
//
// Solidity: function removeContract(bytes32 chainIdentifier, bytes32 contractIdentifier) returns(bool success)
func (_IZetaRegistry *IZetaRegistrySession) RemoveContract(chainIdentifier [32]byte, contractIdentifier [32]byte) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.RemoveContract(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifier)
}

// RemoveContract is a paid mutator transaction binding the contract method 0xd8f4b2bc.
//
// Solidity: function removeContract(bytes32 chainIdentifier, bytes32 contractIdentifier) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactorSession) RemoveContract(chainIdentifier [32]byte, contractIdentifier [32]byte) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.RemoveContract(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifier)
}

// SetChainActive is a paid mutator transaction binding the contract method 0x634c31da.
//
// Solidity: function setChainActive(bytes32 chainIdentifier, bool active) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactor) SetChainActive(opts *bind.TransactOpts, chainIdentifier [32]byte, active bool) (*types.Transaction, error) {
	return _IZetaRegistry.contract.Transact(opts, "setChainActive", chainIdentifier, active)
}

// SetChainActive is a paid mutator transaction binding the contract method 0x634c31da.
//
// Solidity: function setChainActive(bytes32 chainIdentifier, bool active) returns(bool success)
func (_IZetaRegistry *IZetaRegistrySession) SetChainActive(chainIdentifier [32]byte, active bool) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.SetChainActive(&_IZetaRegistry.TransactOpts, chainIdentifier, active)
}

// SetChainActive is a paid mutator transaction binding the contract method 0x634c31da.
//
// Solidity: function setChainActive(bytes32 chainIdentifier, bool active) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactorSession) SetChainActive(chainIdentifier [32]byte, active bool) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.SetChainActive(&_IZetaRegistry.TransactOpts, chainIdentifier, active)
}

// SetContractActive is a paid mutator transaction binding the contract method 0x208b64af.
//
// Solidity: function setContractActive(bytes32 chainIdentifier, bytes32 contractIdentifier, bool active) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactor) SetContractActive(opts *bind.TransactOpts, chainIdentifier [32]byte, contractIdentifier [32]byte, active bool) (*types.Transaction, error) {
	return _IZetaRegistry.contract.Transact(opts, "setContractActive", chainIdentifier, contractIdentifier, active)
}

// SetContractActive is a paid mutator transaction binding the contract method 0x208b64af.
//
// Solidity: function setContractActive(bytes32 chainIdentifier, bytes32 contractIdentifier, bool active) returns(bool success)
func (_IZetaRegistry *IZetaRegistrySession) SetContractActive(chainIdentifier [32]byte, contractIdentifier [32]byte, active bool) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.SetContractActive(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifier, active)
}

// SetContractActive is a paid mutator transaction binding the contract method 0x208b64af.
//
// Solidity: function setContractActive(bytes32 chainIdentifier, bytes32 contractIdentifier, bool active) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactorSession) SetContractActive(chainIdentifier [32]byte, contractIdentifier [32]byte, active bool) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.SetContractActive(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifier, active)
}

// UpdateChain is a paid mutator transaction binding the contract method 0xf0b4a432.
//
// Solidity: function updateChain(bytes32 chainIdentifier, (uint256,string,bool,uint8,uint64) chainInfo) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactor) UpdateChain(opts *bind.TransactOpts, chainIdentifier [32]byte, chainInfo IZetaRegistryChainInfo) (*types.Transaction, error) {
	return _IZetaRegistry.contract.Transact(opts, "updateChain", chainIdentifier, chainInfo)
}

// UpdateChain is a paid mutator transaction binding the contract method 0xf0b4a432.
//
// Solidity: function updateChain(bytes32 chainIdentifier, (uint256,string,bool,uint8,uint64) chainInfo) returns(bool success)
func (_IZetaRegistry *IZetaRegistrySession) UpdateChain(chainIdentifier [32]byte, chainInfo IZetaRegistryChainInfo) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.UpdateChain(&_IZetaRegistry.TransactOpts, chainIdentifier, chainInfo)
}

// UpdateChain is a paid mutator transaction binding the contract method 0xf0b4a432.
//
// Solidity: function updateChain(bytes32 chainIdentifier, (uint256,string,bool,uint8,uint64) chainInfo) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactorSession) UpdateChain(chainIdentifier [32]byte, chainInfo IZetaRegistryChainInfo) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.UpdateChain(&_IZetaRegistry.TransactOpts, chainIdentifier, chainInfo)
}

// UpdateContract is a paid mutator transaction binding the contract method 0x40bbbef4.
//
// Solidity: function updateContract(bytes32 chainIdentifier, bytes32 contractIdentifier, (bool,address,string,address,bytes) contractInfo) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactor) UpdateContract(opts *bind.TransactOpts, chainIdentifier [32]byte, contractIdentifier [32]byte, contractInfo IZetaRegistryContractInfo) (*types.Transaction, error) {
	return _IZetaRegistry.contract.Transact(opts, "updateContract", chainIdentifier, contractIdentifier, contractInfo)
}

// UpdateContract is a paid mutator transaction binding the contract method 0x40bbbef4.
//
// Solidity: function updateContract(bytes32 chainIdentifier, bytes32 contractIdentifier, (bool,address,string,address,bytes) contractInfo) returns(bool success)
func (_IZetaRegistry *IZetaRegistrySession) UpdateContract(chainIdentifier [32]byte, contractIdentifier [32]byte, contractInfo IZetaRegistryContractInfo) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.UpdateContract(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifier, contractInfo)
}

// UpdateContract is a paid mutator transaction binding the contract method 0x40bbbef4.
//
// Solidity: function updateContract(bytes32 chainIdentifier, bytes32 contractIdentifier, (bool,address,string,address,bytes) contractInfo) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactorSession) UpdateContract(chainIdentifier [32]byte, contractIdentifier [32]byte, contractInfo IZetaRegistryContractInfo) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.UpdateContract(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifier, contractInfo)
}

// UpdateContractAddress is a paid mutator transaction binding the contract method 0x59a6152b.
//
// Solidity: function updateContractAddress(bytes32 chainIdentifier, bytes32 contractIdentifier, address newAddress) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactor) UpdateContractAddress(opts *bind.TransactOpts, chainIdentifier [32]byte, contractIdentifier [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _IZetaRegistry.contract.Transact(opts, "updateContractAddress", chainIdentifier, contractIdentifier, newAddress)
}

// UpdateContractAddress is a paid mutator transaction binding the contract method 0x59a6152b.
//
// Solidity: function updateContractAddress(bytes32 chainIdentifier, bytes32 contractIdentifier, address newAddress) returns(bool success)
func (_IZetaRegistry *IZetaRegistrySession) UpdateContractAddress(chainIdentifier [32]byte, contractIdentifier [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.UpdateContractAddress(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifier, newAddress)
}

// UpdateContractAddress is a paid mutator transaction binding the contract method 0x59a6152b.
//
// Solidity: function updateContractAddress(bytes32 chainIdentifier, bytes32 contractIdentifier, address newAddress) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactorSession) UpdateContractAddress(chainIdentifier [32]byte, contractIdentifier [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.UpdateContractAddress(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifier, newAddress)
}

// UpdateContractConfiguration is a paid mutator transaction binding the contract method 0xada6ae93.
//
// Solidity: function updateContractConfiguration(bytes32 chainIdentifier, bytes32 contractIdentifier, bytes configurationData) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactor) UpdateContractConfiguration(opts *bind.TransactOpts, chainIdentifier [32]byte, contractIdentifier [32]byte, configurationData []byte) (*types.Transaction, error) {
	return _IZetaRegistry.contract.Transact(opts, "updateContractConfiguration", chainIdentifier, contractIdentifier, configurationData)
}

// UpdateContractConfiguration is a paid mutator transaction binding the contract method 0xada6ae93.
//
// Solidity: function updateContractConfiguration(bytes32 chainIdentifier, bytes32 contractIdentifier, bytes configurationData) returns(bool success)
func (_IZetaRegistry *IZetaRegistrySession) UpdateContractConfiguration(chainIdentifier [32]byte, contractIdentifier [32]byte, configurationData []byte) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.UpdateContractConfiguration(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifier, configurationData)
}

// UpdateContractConfiguration is a paid mutator transaction binding the contract method 0xada6ae93.
//
// Solidity: function updateContractConfiguration(bytes32 chainIdentifier, bytes32 contractIdentifier, bytes configurationData) returns(bool success)
func (_IZetaRegistry *IZetaRegistryTransactorSession) UpdateContractConfiguration(chainIdentifier [32]byte, contractIdentifier [32]byte, configurationData []byte) (*types.Transaction, error) {
	return _IZetaRegistry.Contract.UpdateContractConfiguration(&_IZetaRegistry.TransactOpts, chainIdentifier, contractIdentifier, configurationData)
}

// IZetaRegistryChainAddedIterator is returned from FilterChainAdded and is used to iterate over the raw logs and unpacked data for ChainAdded events raised by the IZetaRegistry contract.
type IZetaRegistryChainAddedIterator struct {
	Event *IZetaRegistryChainAdded // Event containing the contract specifics and raw log

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
func (it *IZetaRegistryChainAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaRegistryChainAdded)
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
		it.Event = new(IZetaRegistryChainAdded)
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
func (it *IZetaRegistryChainAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaRegistryChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaRegistryChainAdded represents a ChainAdded event raised by the IZetaRegistry contract.
type IZetaRegistryChainAdded struct {
	ChainIdentifier [32]byte
	ChainId         *big.Int
	Name            string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChainAdded is a free log retrieval operation binding the contract event 0x5a8e7b221f7affd9f2b1becd6ad936d4c721d1dd0f8666d7ddf1d624664edd1a.
//
// Solidity: event ChainAdded(bytes32 indexed chainIdentifier, uint256 chainId, string name)
func (_IZetaRegistry *IZetaRegistryFilterer) FilterChainAdded(opts *bind.FilterOpts, chainIdentifier [][32]byte) (*IZetaRegistryChainAddedIterator, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}

	logs, sub, err := _IZetaRegistry.contract.FilterLogs(opts, "ChainAdded", chainIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryChainAddedIterator{contract: _IZetaRegistry.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

// WatchChainAdded is a free log subscription operation binding the contract event 0x5a8e7b221f7affd9f2b1becd6ad936d4c721d1dd0f8666d7ddf1d624664edd1a.
//
// Solidity: event ChainAdded(bytes32 indexed chainIdentifier, uint256 chainId, string name)
func (_IZetaRegistry *IZetaRegistryFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *IZetaRegistryChainAdded, chainIdentifier [][32]byte) (event.Subscription, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}

	logs, sub, err := _IZetaRegistry.contract.WatchLogs(opts, "ChainAdded", chainIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaRegistryChainAdded)
				if err := _IZetaRegistry.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

// ParseChainAdded is a log parse operation binding the contract event 0x5a8e7b221f7affd9f2b1becd6ad936d4c721d1dd0f8666d7ddf1d624664edd1a.
//
// Solidity: event ChainAdded(bytes32 indexed chainIdentifier, uint256 chainId, string name)
func (_IZetaRegistry *IZetaRegistryFilterer) ParseChainAdded(log types.Log) (*IZetaRegistryChainAdded, error) {
	event := new(IZetaRegistryChainAdded)
	if err := _IZetaRegistry.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaRegistryChainStatusUpdatedIterator is returned from FilterChainStatusUpdated and is used to iterate over the raw logs and unpacked data for ChainStatusUpdated events raised by the IZetaRegistry contract.
type IZetaRegistryChainStatusUpdatedIterator struct {
	Event *IZetaRegistryChainStatusUpdated // Event containing the contract specifics and raw log

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
func (it *IZetaRegistryChainStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaRegistryChainStatusUpdated)
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
		it.Event = new(IZetaRegistryChainStatusUpdated)
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
func (it *IZetaRegistryChainStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaRegistryChainStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaRegistryChainStatusUpdated represents a ChainStatusUpdated event raised by the IZetaRegistry contract.
type IZetaRegistryChainStatusUpdated struct {
	ChainIdentifier [32]byte
	Active          bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChainStatusUpdated is a free log retrieval operation binding the contract event 0xae28df7592823455d07510da74e0d26fa77925e92a66817e8c842056c490d4f8.
//
// Solidity: event ChainStatusUpdated(bytes32 indexed chainIdentifier, bool active)
func (_IZetaRegistry *IZetaRegistryFilterer) FilterChainStatusUpdated(opts *bind.FilterOpts, chainIdentifier [][32]byte) (*IZetaRegistryChainStatusUpdatedIterator, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}

	logs, sub, err := _IZetaRegistry.contract.FilterLogs(opts, "ChainStatusUpdated", chainIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryChainStatusUpdatedIterator{contract: _IZetaRegistry.contract, event: "ChainStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchChainStatusUpdated is a free log subscription operation binding the contract event 0xae28df7592823455d07510da74e0d26fa77925e92a66817e8c842056c490d4f8.
//
// Solidity: event ChainStatusUpdated(bytes32 indexed chainIdentifier, bool active)
func (_IZetaRegistry *IZetaRegistryFilterer) WatchChainStatusUpdated(opts *bind.WatchOpts, sink chan<- *IZetaRegistryChainStatusUpdated, chainIdentifier [][32]byte) (event.Subscription, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}

	logs, sub, err := _IZetaRegistry.contract.WatchLogs(opts, "ChainStatusUpdated", chainIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaRegistryChainStatusUpdated)
				if err := _IZetaRegistry.contract.UnpackLog(event, "ChainStatusUpdated", log); err != nil {
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

// ParseChainStatusUpdated is a log parse operation binding the contract event 0xae28df7592823455d07510da74e0d26fa77925e92a66817e8c842056c490d4f8.
//
// Solidity: event ChainStatusUpdated(bytes32 indexed chainIdentifier, bool active)
func (_IZetaRegistry *IZetaRegistryFilterer) ParseChainStatusUpdated(log types.Log) (*IZetaRegistryChainStatusUpdated, error) {
	event := new(IZetaRegistryChainStatusUpdated)
	if err := _IZetaRegistry.contract.UnpackLog(event, "ChainStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaRegistryContractAddressUpdatedIterator is returned from FilterContractAddressUpdated and is used to iterate over the raw logs and unpacked data for ContractAddressUpdated events raised by the IZetaRegistry contract.
type IZetaRegistryContractAddressUpdatedIterator struct {
	Event *IZetaRegistryContractAddressUpdated // Event containing the contract specifics and raw log

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
func (it *IZetaRegistryContractAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaRegistryContractAddressUpdated)
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
		it.Event = new(IZetaRegistryContractAddressUpdated)
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
func (it *IZetaRegistryContractAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaRegistryContractAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaRegistryContractAddressUpdated represents a ContractAddressUpdated event raised by the IZetaRegistry contract.
type IZetaRegistryContractAddressUpdated struct {
	ChainIdentifier    [32]byte
	ContractIdentifier [32]byte
	Category           uint8
	ContractAddress    common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterContractAddressUpdated is a free log retrieval operation binding the contract event 0xe8bec620648cefe1f492bdb89a04125718c351bc0c539443bbfddba38e60a652.
//
// Solidity: event ContractAddressUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, uint8 indexed category, address contractAddress)
func (_IZetaRegistry *IZetaRegistryFilterer) FilterContractAddressUpdated(opts *bind.FilterOpts, chainIdentifier [][32]byte, contractIdentifier [][32]byte, category []uint8) (*IZetaRegistryContractAddressUpdatedIterator, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}
	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _IZetaRegistry.contract.FilterLogs(opts, "ContractAddressUpdated", chainIdentifierRule, contractIdentifierRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryContractAddressUpdatedIterator{contract: _IZetaRegistry.contract, event: "ContractAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchContractAddressUpdated is a free log subscription operation binding the contract event 0xe8bec620648cefe1f492bdb89a04125718c351bc0c539443bbfddba38e60a652.
//
// Solidity: event ContractAddressUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, uint8 indexed category, address contractAddress)
func (_IZetaRegistry *IZetaRegistryFilterer) WatchContractAddressUpdated(opts *bind.WatchOpts, sink chan<- *IZetaRegistryContractAddressUpdated, chainIdentifier [][32]byte, contractIdentifier [][32]byte, category []uint8) (event.Subscription, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}
	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _IZetaRegistry.contract.WatchLogs(opts, "ContractAddressUpdated", chainIdentifierRule, contractIdentifierRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaRegistryContractAddressUpdated)
				if err := _IZetaRegistry.contract.UnpackLog(event, "ContractAddressUpdated", log); err != nil {
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

// ParseContractAddressUpdated is a log parse operation binding the contract event 0xe8bec620648cefe1f492bdb89a04125718c351bc0c539443bbfddba38e60a652.
//
// Solidity: event ContractAddressUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, uint8 indexed category, address contractAddress)
func (_IZetaRegistry *IZetaRegistryFilterer) ParseContractAddressUpdated(log types.Log) (*IZetaRegistryContractAddressUpdated, error) {
	event := new(IZetaRegistryContractAddressUpdated)
	if err := _IZetaRegistry.contract.UnpackLog(event, "ContractAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaRegistryContractConfigurationUpdatedIterator is returned from FilterContractConfigurationUpdated and is used to iterate over the raw logs and unpacked data for ContractConfigurationUpdated events raised by the IZetaRegistry contract.
type IZetaRegistryContractConfigurationUpdatedIterator struct {
	Event *IZetaRegistryContractConfigurationUpdated // Event containing the contract specifics and raw log

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
func (it *IZetaRegistryContractConfigurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaRegistryContractConfigurationUpdated)
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
		it.Event = new(IZetaRegistryContractConfigurationUpdated)
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
func (it *IZetaRegistryContractConfigurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaRegistryContractConfigurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaRegistryContractConfigurationUpdated represents a ContractConfigurationUpdated event raised by the IZetaRegistry contract.
type IZetaRegistryContractConfigurationUpdated struct {
	ChainIdentifier    [32]byte
	ContractIdentifier [32]byte
	ConfigurationData  []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterContractConfigurationUpdated is a free log retrieval operation binding the contract event 0x11af48d095e78bf8811830bb22710e171451b6a34ddbe65490b4b2e31b34c310.
//
// Solidity: event ContractConfigurationUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bytes configurationData)
func (_IZetaRegistry *IZetaRegistryFilterer) FilterContractConfigurationUpdated(opts *bind.FilterOpts, chainIdentifier [][32]byte, contractIdentifier [][32]byte) (*IZetaRegistryContractConfigurationUpdatedIterator, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}

	logs, sub, err := _IZetaRegistry.contract.FilterLogs(opts, "ContractConfigurationUpdated", chainIdentifierRule, contractIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryContractConfigurationUpdatedIterator{contract: _IZetaRegistry.contract, event: "ContractConfigurationUpdated", logs: logs, sub: sub}, nil
}

// WatchContractConfigurationUpdated is a free log subscription operation binding the contract event 0x11af48d095e78bf8811830bb22710e171451b6a34ddbe65490b4b2e31b34c310.
//
// Solidity: event ContractConfigurationUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bytes configurationData)
func (_IZetaRegistry *IZetaRegistryFilterer) WatchContractConfigurationUpdated(opts *bind.WatchOpts, sink chan<- *IZetaRegistryContractConfigurationUpdated, chainIdentifier [][32]byte, contractIdentifier [][32]byte) (event.Subscription, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}

	logs, sub, err := _IZetaRegistry.contract.WatchLogs(opts, "ContractConfigurationUpdated", chainIdentifierRule, contractIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaRegistryContractConfigurationUpdated)
				if err := _IZetaRegistry.contract.UnpackLog(event, "ContractConfigurationUpdated", log); err != nil {
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

// ParseContractConfigurationUpdated is a log parse operation binding the contract event 0x11af48d095e78bf8811830bb22710e171451b6a34ddbe65490b4b2e31b34c310.
//
// Solidity: event ContractConfigurationUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bytes configurationData)
func (_IZetaRegistry *IZetaRegistryFilterer) ParseContractConfigurationUpdated(log types.Log) (*IZetaRegistryContractConfigurationUpdated, error) {
	event := new(IZetaRegistryContractConfigurationUpdated)
	if err := _IZetaRegistry.contract.UnpackLog(event, "ContractConfigurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaRegistryContractStatusUpdatedIterator is returned from FilterContractStatusUpdated and is used to iterate over the raw logs and unpacked data for ContractStatusUpdated events raised by the IZetaRegistry contract.
type IZetaRegistryContractStatusUpdatedIterator struct {
	Event *IZetaRegistryContractStatusUpdated // Event containing the contract specifics and raw log

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
func (it *IZetaRegistryContractStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaRegistryContractStatusUpdated)
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
		it.Event = new(IZetaRegistryContractStatusUpdated)
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
func (it *IZetaRegistryContractStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaRegistryContractStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaRegistryContractStatusUpdated represents a ContractStatusUpdated event raised by the IZetaRegistry contract.
type IZetaRegistryContractStatusUpdated struct {
	ChainIdentifier    [32]byte
	ContractIdentifier [32]byte
	Active             bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterContractStatusUpdated is a free log retrieval operation binding the contract event 0x166bdf9d7b5f4849bcabfef5465f6edae79fe8f333c2019f6f5cf489084cf581.
//
// Solidity: event ContractStatusUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bool active)
func (_IZetaRegistry *IZetaRegistryFilterer) FilterContractStatusUpdated(opts *bind.FilterOpts, chainIdentifier [][32]byte, contractIdentifier [][32]byte) (*IZetaRegistryContractStatusUpdatedIterator, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}

	logs, sub, err := _IZetaRegistry.contract.FilterLogs(opts, "ContractStatusUpdated", chainIdentifierRule, contractIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryContractStatusUpdatedIterator{contract: _IZetaRegistry.contract, event: "ContractStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchContractStatusUpdated is a free log subscription operation binding the contract event 0x166bdf9d7b5f4849bcabfef5465f6edae79fe8f333c2019f6f5cf489084cf581.
//
// Solidity: event ContractStatusUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bool active)
func (_IZetaRegistry *IZetaRegistryFilterer) WatchContractStatusUpdated(opts *bind.WatchOpts, sink chan<- *IZetaRegistryContractStatusUpdated, chainIdentifier [][32]byte, contractIdentifier [][32]byte) (event.Subscription, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}

	logs, sub, err := _IZetaRegistry.contract.WatchLogs(opts, "ContractStatusUpdated", chainIdentifierRule, contractIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaRegistryContractStatusUpdated)
				if err := _IZetaRegistry.contract.UnpackLog(event, "ContractStatusUpdated", log); err != nil {
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

// ParseContractStatusUpdated is a log parse operation binding the contract event 0x166bdf9d7b5f4849bcabfef5465f6edae79fe8f333c2019f6f5cf489084cf581.
//
// Solidity: event ContractStatusUpdated(bytes32 indexed chainIdentifier, bytes32 indexed contractIdentifier, bool active)
func (_IZetaRegistry *IZetaRegistryFilterer) ParseContractStatusUpdated(log types.Log) (*IZetaRegistryContractStatusUpdated, error) {
	event := new(IZetaRegistryContractStatusUpdated)
	if err := _IZetaRegistry.contract.UnpackLog(event, "ContractStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaRegistryZRC20TokenAddedIterator is returned from FilterZRC20TokenAdded and is used to iterate over the raw logs and unpacked data for ZRC20TokenAdded events raised by the IZetaRegistry contract.
type IZetaRegistryZRC20TokenAddedIterator struct {
	Event *IZetaRegistryZRC20TokenAdded // Event containing the contract specifics and raw log

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
func (it *IZetaRegistryZRC20TokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaRegistryZRC20TokenAdded)
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
		it.Event = new(IZetaRegistryZRC20TokenAdded)
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
func (it *IZetaRegistryZRC20TokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaRegistryZRC20TokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaRegistryZRC20TokenAdded represents a ZRC20TokenAdded event raised by the IZetaRegistry contract.
type IZetaRegistryZRC20TokenAdded struct {
	ChainIdentifier [32]byte
	TokenIdentifier [32]byte
	TokenAddress    common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterZRC20TokenAdded is a free log retrieval operation binding the contract event 0xbc0f3f1af072bc04772392bb2fcb5a0f1a9ec90139271c5599034d3b8e5888b3.
//
// Solidity: event ZRC20TokenAdded(bytes32 indexed chainIdentifier, bytes32 indexed tokenIdentifier, address tokenAddress)
func (_IZetaRegistry *IZetaRegistryFilterer) FilterZRC20TokenAdded(opts *bind.FilterOpts, chainIdentifier [][32]byte, tokenIdentifier [][32]byte) (*IZetaRegistryZRC20TokenAddedIterator, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var tokenIdentifierRule []interface{}
	for _, tokenIdentifierItem := range tokenIdentifier {
		tokenIdentifierRule = append(tokenIdentifierRule, tokenIdentifierItem)
	}

	logs, sub, err := _IZetaRegistry.contract.FilterLogs(opts, "ZRC20TokenAdded", chainIdentifierRule, tokenIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IZetaRegistryZRC20TokenAddedIterator{contract: _IZetaRegistry.contract, event: "ZRC20TokenAdded", logs: logs, sub: sub}, nil
}

// WatchZRC20TokenAdded is a free log subscription operation binding the contract event 0xbc0f3f1af072bc04772392bb2fcb5a0f1a9ec90139271c5599034d3b8e5888b3.
//
// Solidity: event ZRC20TokenAdded(bytes32 indexed chainIdentifier, bytes32 indexed tokenIdentifier, address tokenAddress)
func (_IZetaRegistry *IZetaRegistryFilterer) WatchZRC20TokenAdded(opts *bind.WatchOpts, sink chan<- *IZetaRegistryZRC20TokenAdded, chainIdentifier [][32]byte, tokenIdentifier [][32]byte) (event.Subscription, error) {

	var chainIdentifierRule []interface{}
	for _, chainIdentifierItem := range chainIdentifier {
		chainIdentifierRule = append(chainIdentifierRule, chainIdentifierItem)
	}
	var tokenIdentifierRule []interface{}
	for _, tokenIdentifierItem := range tokenIdentifier {
		tokenIdentifierRule = append(tokenIdentifierRule, tokenIdentifierItem)
	}

	logs, sub, err := _IZetaRegistry.contract.WatchLogs(opts, "ZRC20TokenAdded", chainIdentifierRule, tokenIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaRegistryZRC20TokenAdded)
				if err := _IZetaRegistry.contract.UnpackLog(event, "ZRC20TokenAdded", log); err != nil {
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

// ParseZRC20TokenAdded is a log parse operation binding the contract event 0xbc0f3f1af072bc04772392bb2fcb5a0f1a9ec90139271c5599034d3b8e5888b3.
//
// Solidity: event ZRC20TokenAdded(bytes32 indexed chainIdentifier, bytes32 indexed tokenIdentifier, address tokenAddress)
func (_IZetaRegistry *IZetaRegistryFilterer) ParseZRC20TokenAdded(log types.Log) (*IZetaRegistryZRC20TokenAdded, error) {
	event := new(IZetaRegistryZRC20TokenAdded)
	if err := _IZetaRegistry.contract.UnpackLog(event, "ZRC20TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
