// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// ChainInfoDTO is an auto generated low-level Go binding around an user-defined struct.
type ChainInfoDTO struct {
	Active   bool
	ChainId  *big.Int
	GasZRC20 common.Address
	Registry []byte
}

// ContractInfoDTO is an auto generated low-level Go binding around an user-defined struct.
type ContractInfoDTO struct {
	Active       bool
	AddressBytes []byte
	ContractType string
	ChainId      *big.Int
}

// IRegistryChainMetadataEntry is an auto generated low-level Go binding around an user-defined struct.
type IRegistryChainMetadataEntry struct {
	ChainId *big.Int
	Key     string
	Value   []byte
}

// IRegistryContractConfigEntry is an auto generated low-level Go binding around an user-defined struct.
type IRegistryContractConfigEntry struct {
	ChainId      *big.Int
	ContractType string
	Key          string
	Value        []byte
}

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender common.Address
}

// ZRC20Info is an auto generated low-level Go binding around an user-defined struct.
type ZRC20Info struct {
	Active        bool
	Address       common.Address
	OriginAddress []byte
	OriginChainId *big.Int
	Symbol        string
	CoinType      string
	Decimals      uint8
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GATEWAY_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY_MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bootstrapChains\",\"inputs\":[{\"name\":\"chains\",\"type\":\"tuple[]\",\"internalType\":\"structChainInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasZRC20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"metadataEntries\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistry.ChainMetadataEntry[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bootstrapContracts\",\"inputs\":[{\"name\":\"contracts\",\"type\":\"tuple[]\",\"internalType\":\"structContractInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"configEntries\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistry.ContractConfigEntry[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bootstrapZRC20Tokens\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structZRC20Info[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeChainStatus\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasZRC20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"activation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"coreRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gatewayEVM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllChains\",\"inputs\":[],\"outputs\":[{\"name\":\"chainsInfo\",\"type\":\"tuple[]\",\"internalType\":\"structChainInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasZRC20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"contractsInfo\",\"type\":\"tuple[]\",\"internalType\":\"structContractInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllZRC20Tokens\",\"inputs\":[],\"outputs\":[{\"name\":\"tokensInfo\",\"type\":\"tuple[]\",\"internalType\":\"structZRC20Info[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainMetadata\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractConfiguration\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractInfo\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getZRC20AddressByForeignAsset\",\"inputs\":[{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getZRC20TokenInfo\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauserAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registryManager_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gatewayEVM_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"coreRegistry_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerContract\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerZRC20Token\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setZRC20TokenActive\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateChainMetadata\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateContractConfiguration\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"ChainMetadataUpdated\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainStatusChanged\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newStatus\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractConfigurationUpdated\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractStatusChanged\",\"inputs\":[{\"name\":\"addressBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenRegistered\",\"inputs\":[{\"name\":\"originAddress\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"address_\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenUpdated\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ChainActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainNonActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ContractAlreadyRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ContractNotFound\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContractType\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZRC20AlreadyRegistered\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZRC20SymbolAlreadyInUse\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614a016100fd6000396000818161277e015281816127a7015261295a0152614a016000f3fe6080604052600436106102a05760003560e01c80638456cb591161016e578063ad3cb1cc116100cb578063d547741f1161007f578063e63ab1e911610064578063e63ab1e914610828578063e9d6c5ba1461085c578063f354b31f1461088e57600080fd5b8063d547741f146107e8578063dc07f8661461080857600080fd5b8063c1bd469f116100b0578063c1bd469f14610772578063c50d6dd914610794578063d3523ea2146107c857600080fd5b8063ad3cb1cc146106f5578063bd8fde1c1461073e57600080fd5b80639ca220dd11610122578063a4b7e5f911610107578063a4b7e5f914610695578063a8f2cb96146106b5578063aa808c06146106d557600080fd5b80639ca220dd1461065e578063a217fddf1461068057600080fd5b8063909b6a0311610153578063909b6a03146105b757806391d14854146105d757806394cc86831461063c57600080fd5b80638456cb59146105825780639060bda91461059757600080fd5b80633f4ba83a1161021c578063631d62e4116101d05780636bf3d05a116101b55780636bf3d05a146105225780636e9e2d3f146105425780637066b18d1461056257600080fd5b8063631d62e4146104d5578063676cc054146104f557600080fd5b806352d1902d1161020157806352d1902d1461045b5780635c975abb146104705780635cf92c9f146104a757600080fd5b80633f4ba83a146104335780634f1ef2861461044857600080fd5b80632259e9e5116102735780632c78f74c116102585780632c78f74c146103bb5780632f2ff15d146103f357806336568abe1461041357600080fd5b80632259e9e51461033e578063248a9ca31461035e57600080fd5b806301ffc9a7146102a557806310d29b9e146102da5780631459457a146102fc57806318d3ce961461031c575b600080fd5b3480156102b157600080fd5b506102c56102c03660046138e7565b6108ae565b60405190151581526020015b60405180910390f35b3480156102e657600080fd5b506102fa6102f5366004613987565b610947565b005b34801561030857600080fd5b506102fa6103173660046139fc565b6109fc565b34801561032857600080fd5b50610331610d19565b6040516102d19190613ab1565b34801561034a57600080fd5b506102fa610359366004613b72565b610fbb565b34801561036a57600080fd5b506103ad610379366004613bf1565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016102d1565b3480156103c757600080fd5b506009546103db906001600160a01b031681565b6040516001600160a01b0390911681526020016102d1565b3480156103ff57600080fd5b506102fa61040e366004613c0a565b61104e565b34801561041f57600080fd5b506102fa61042e366004613c0a565b611098565b34801561043f57600080fd5b506102fa6110e9565b6102fa610456366004613c65565b6110ff565b34801561046757600080fd5b506103ad61111e565b34801561047c57600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166102c5565b3480156104b357600080fd5b506104c76104c2366004613d30565b61114d565b6040516102d1929190613d7c565b3480156104e157600080fd5b506102fa6104f0366004613b72565b611248565b34801561050157600080fd5b50610515610510366004613d9f565b6112ef565b6040516102d19190613de1565b34801561052e57600080fd5b50600a546103db906001600160a01b031681565b34801561054e57600080fd5b506102fa61055d366004613e05565b6113f7565b34801561056e57600080fd5b5061051561057d366004613d30565b6114b3565b34801561058e57600080fd5b506102fa61157c565b3480156105a357600080fd5b506102fa6105b2366004613ed5565b6115ae565b3480156105c357600080fd5b506102fa6105d2366004613f44565b611640565b3480156105e357600080fd5b506102c56105f2366004613c0a565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561064857600080fd5b50610651611701565b6040516102d19190613f86565b34801561066a57600080fd5b50610673611759565b6040516102d19190613fc9565b34801561068c57600080fd5b506103ad600081565b3480156106a157600080fd5b506102fa6106b0366004614076565b611918565b3480156106c157600080fd5b506102fa6106d03660046140e7565b611a21565b3480156106e157600080fd5b506103db6106f0366004613d30565b611aa3565b34801561070157600080fd5b506105156040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b34801561074a57600080fd5b506103ad7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa381565b34801561077e57600080fd5b50610787611ae3565b6040516102d1919061414a565b3480156107a057600080fd5b506103ad7fb90e9995c6170fff8ea03e9ad6919878e483770c237f1a6f330ceaa7112b344a81565b3480156107d457600080fd5b506105156107e3366004613b72565b611dea565b3480156107f457600080fd5b506102fa610803366004613c0a565b611ed2565b34801561081457600080fd5b506102fa610823366004614076565b611f16565b34801561083457600080fd5b506103ad7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b34801561086857600080fd5b5061087c610877366004614241565b61201a565b6040516102d19695949392919061425c565b34801561089a57600080fd5b506102fa6108a93660046142ba565b61226c565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061094157507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b333014610980576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61098c84848484612307565b7f6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c6005600086815260200190815260200160002084846040516109d092919061436a565b90815260200160405180910390206001016040516109ee91906143cd565b60405180910390a150505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610a475750825b905060008267ffffffffffffffff166001148015610a645750303b155b905081158015610a72575080155b15610aa9576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610b0a5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038a161580610b2757506001600160a01b038716155b80610b3957506001600160a01b038616155b80610b4b57506001600160a01b038916155b80610b5d57506001600160a01b038816155b15610b94576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b9c61245b565b610ba461245b565b610bac612465565b610bb760008b612498565b50610be27f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8b612498565b50610c0d7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8a612498565b50610c387ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa389612498565b50610c637fb90e9995c6170fff8ea03e9ad6919878e483770c237f1a6f330ceaa7112b344a88612498565b50600980546001600160a01b03808a167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255600a8054928916929091169190911790558315610d0d5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b6002546060908067ffffffffffffffff811115610d3857610d38613c36565b604051908082528060200260200182016040528015610d9657816020015b610d8360405180608001604052806000151581526020016060815260200160608152602001600081525090565b815260200190600190039081610d565790505b50915060005b81811015610fb657600060028281548110610db957610db9614459565b906000526020600020906002020160405180604001604052908160008201548152602001600182018054610dec9061437a565b80601f0160208091040260200160405190810160405280929190818152602001828054610e189061437a565b8015610e655780601f10610e3a57610100808354040283529160200191610e65565b820191906000526020600020905b815481529060010190602001808311610e4857829003601f168201915b505050505081525050905060008160000151905060008260200151905060405180608001604052806005600085815260200190815260200160002083604051610eae9190614488565b90815260408051602092819003830190205460ff161515835260008681526005835281902090519290910191610ee5908590614488565b90815260200160405180910390206001018054610f019061437a565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2d9061437a565b8015610f7a5780601f10610f4f57610100808354040283529160200191610f7a565b820191906000526020600020905b815481529060010190602001808311610f5d57829003601f168201915b5050505050815260200182815260200183815250868581518110610fa057610fa0614459565b6020908102919091010152505050600101610d9c565b505090565b333014610ff4576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ffc612567565b61100985858585856125c3565b847f40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee6348585858560405161103f94939291906144cf565b60405180910390a25050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461108881612651565b6110928383612498565b50505050565b6001600160a01b03811633146110da576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110e4828261265b565b505050565b60006110f481612651565b6110fc612701565b50565b611107612773565b61111082612843565b61111a828261284e565b5050565b600061112861294f565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b60008381526005602052604080822090516060919061116f908690869061436a565b908152604080519182900360209081018320546000898152600590925291902060ff9091169350906111a4908690869061436a565b908152602001604051809103902060010180546111c09061437a565b80601f01602080910402602001604051908101604052809291908181526020018280546111ec9061437a565b80156112395780601f1061120e57610100808354040283529160200191611239565b820191906000526020600020905b81548152906001019060200180831161121c57829003601f168201915b50505050509050935093915050565b333014611281576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611289612567565b61129685858585856129b1565b83836040516112a692919061436a565b6040518091039020857f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce258184846040516112e0929190614501565b60405180910390a35050505050565b60607fb90e9995c6170fff8ea03e9ad6919878e483770c237f1a6f330ceaa7112b344a61131b81612651565b611323612567565b600a546001600160a01b031661133c6020870187614241565b6001600160a01b03161461137c576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080306001600160a01b0316868660405161139992919061436a565b6000604051808303816000865af19150503d80600081146113d6576040519150601f19603f3d011682016040523d82523d6000602084013e6113db565b606091505b5091509150816113ed57805160208201fd5b9695505050505050565b333014611430576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611438612567565b611449898989898989898989612c4c565b886001600160a01b0316858560405161146392919061436a565b60405180910390207fa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e36783898c8c6040516114a09493929190614515565b60405180910390a3505050505050505050565b60606004600085815260200190815260200160002060040183836040516114db92919061436a565b908152602001604051809103902080546114f49061437a565b80601f01602080910402602001604051908101604052809291908181526020018280546115209061437a565b801561156d5780601f106115425761010080835404028352916020019161156d565b820191906000526020600020905b81548152906001019060200180831161155057829003601f168201915b505050505090505b9392505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6115a681612651565b6110fc612fec565b3330146115e7576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115ef612567565b6115f98282613047565b604080516001600160a01b038416815282151560208201527f9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8910160405180910390a15050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361166a81612651565b611672612567565b60005b82811015611092573684848381811061169057611690614459565b90506020028101906116a29190614538565b90506116f86116b76040830160208401614241565b6116c4608084018461456c565b60608501356116d6604087018761456c565b6116e360a089018961456c565b6116f360e08b0160c08c016145d1565b612c4c565b50600101611675565b6060600080548060200260200160405190810160405280929190818152602001828054801561174f57602002820191906000526020600020905b81548152602001906001019080831161173b575b5050505050905090565b6001546060908067ffffffffffffffff81111561177857611778613c36565b6040519080825280602002602001820160405280156117e757816020015b60408051608081018252600080825260208083018290529282015260608082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092019101816117965790505b50915060005b81811015610fb65760006001828154811061180a5761180a614459565b6000918252602080832090910154604080516080810182528285526004808552828620805460ff161515835282860185905260028101546001600160a01b03169383019390935294839052939092526003909101805491935060608301916118719061437a565b80601f016020809104026020016040519081016040528092919081815260200182805461189d9061437a565b80156118ea5780601f106118bf576101008083540402835291602001916118ea565b820191906000526020600020905b8154815290600101906020018083116118cd57829003601f168201915b505050505081525084838151811061190457611904614459565b6020908102919091010152506001016117ed565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361194281612651565b61194a612567565b60005b848110156119ac573686868381811061196857611968614459565b905060200281019061197a91906145ec565b90506119a36060820135611991604084018461456c565b61199e602086018661456c565b6129b1565b5060010161194d565b5060005b82811015611a1957368484838181106119cb576119cb614459565b90506020028101906119dd91906145ec565b9050611a1081356119f1602084018461456c565b6119fe604086018661456c565b611a0b606088018861456c565b613137565b506001016119b0565b505050505050565b333014611a5a576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611a62612567565b611a6f8585858585613299565b847fc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e8260405161103f911515815260200190565b6000838152600860205260408082209051611ac1908590859061436a565b908152604051908190036020019020546001600160a01b031690509392505050565b6003546060908067ffffffffffffffff811115611b0257611b02613c36565b604051908082528060200260200182016040528015611b8157816020015b611b6e6040518060e0016040528060001515815260200160006001600160a01b0316815260200160608152602001600081526020016060815260200160608152602001600060ff1681525090565b815260200190600190039081611b205790505b50915060005b81811015610fb657600060038281548110611ba457611ba4614459565b60009182526020808320909101546001600160a01b0390811680845260068352604093849020845160e081018652815460ff811615158252610100900490931693830193909352600183018054919550919384019190611c039061437a565b80601f0160208091040260200160405190810160405280929190818152602001828054611c2f9061437a565b8015611c7c5780601f10611c5157610100808354040283529160200191611c7c565b820191906000526020600020905b815481529060010190602001808311611c5f57829003601f168201915b5050505050815260200160028201548152602001600382018054611c9f9061437a565b80601f0160208091040260200160405190810160405280929190818152602001828054611ccb9061437a565b8015611d185780601f10611ced57610100808354040283529160200191611d18565b820191906000526020600020905b815481529060010190602001808311611cfb57829003601f168201915b50505050508152602001600482018054611d319061437a565b80601f0160208091040260200160405190810160405280929190818152602001828054611d5d9061437a565b8015611daa5780601f10611d7f57610100808354040283529160200191611daa565b820191906000526020600020905b815481529060010190602001808311611d8d57829003601f168201915b50505091835250506005919091015460ff166020909101528451859084908110611dd657611dd6614459565b602090810291909101015250600101611b87565b6060600560008781526020019081526020016000208585604051611e0f92919061436a565b90815260200160405180910390206003018383604051611e3092919061436a565b90815260200160405180910390208054611e499061437a565b80601f0160208091040260200160405190810160405280929190818152602001828054611e759061437a565b8015611ec25780601f10611e9757610100808354040283529160200191611ec2565b820191906000526020600020905b815481529060010190602001808311611ea557829003601f168201915b5050505050905095945050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611f0c81612651565b611092838361265b565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3611f4081612651565b611f48612567565b60005b84811015611fba5736868683818110611f6657611f66614459565b9050602002810190611f7891906145ec565b9050611fb16020820135611f926060840160408501614241565b611f9f606085018561456c565b611fac6020870187614620565b613299565b50600101611f4b565b5060005b82811015611a195736848483818110611fd957611fd9614459565b9050602002810190611feb919061463b565b90506120118135611fff602084018461456c565b61200c604086018661456c565b6125c3565b50600101611fbe565b6001600160a01b038082166000908152600660209081526040808320815160e081018352815460ff8116151582526101009004909516928501929092526001820180549394606094869486948594879485949093928401919061207c9061437a565b80601f01602080910402602001604051908101604052809291908181526020018280546120a89061437a565b80156120f55780601f106120ca576101008083540402835291602001916120f5565b820191906000526020600020905b8154815290600101906020018083116120d857829003601f168201915b50505050508152602001600282015481526020016003820180546121189061437a565b80601f01602080910402602001604051908101604052809291908181526020018280546121449061437a565b80156121915780601f1061216657610100808354040283529160200191612191565b820191906000526020600020905b81548152906001019060200180831161217457829003601f168201915b505050505081526020016004820180546121aa9061437a565b80601f01602080910402602001604051908101604052809291908181526020018280546121d69061437a565b80156122235780601f106121f857610100808354040283529160200191612223565b820191906000526020600020905b81548152906001019060200180831161220657829003601f168201915b50505091835250506005919091015460ff16602090910152805160808201516060830151604084015160a085015160c090950151939d929c50909a509850919650945092505050565b3330146122a5576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6122ad612567565b6122bc87878787878787613137565b867faea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c58787878787876040516122f69695949392919061466f565b60405180910390a250505050505050565b60008481526004602052604090205460ff16612357576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018590526024015b60405180910390fd5b60008290036123955782826040517ec10cfd00000000000000000000000000000000000000000000000000000000815260040161234e929190614501565b6000848152600560205260409081902090516123b4908590859061436a565b908152602001604051809103902060010180546123d09061437a565b9050600003612411578383836040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161234e939291906146b8565b8060056000868152602001908152602001600020848460405161243592919061436a565b908152604051908190036020019020805491151560ff1990921691909117905550505050565b612463613485565b565b61246d613485565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1661255d576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556125133390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610941565b6000915050610941565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615612463576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008581526004602052604090205460ff1661260e576040517f8e6feba50000000000000000000000000000000000000000000000000000000081526004810186905260240161234e565b818160046000888152602001908152602001600020600401868660405161263692919061436a565b90815260200160405180910390209182611a19929190614719565b6110fc81336134ec565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff161561255d576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610941565b612709613579565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061280c57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166128007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15612463576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061111a81612651565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156128a8575060408051601f3d908101601f191682019092526128a591810190614815565b60015b6128e9576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161234e565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612945576040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004810182905260240161234e565b6110e483836135d4565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612463576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008581526004602052604090205460ff166129fc576040517f8e6feba50000000000000000000000000000000000000000000000000000000081526004810186905260240161234e565b6000839003612a3a5783836040517ec10cfd00000000000000000000000000000000000000000000000000000000815260040161234e929190614501565b6000819003612a75576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000858152600560205260408082209051612a93908790879061436a565b90815260200160405180910390206001018054612aaf9061437a565b90501115612af35784848484846040517f2b192eab00000000000000000000000000000000000000000000000000000000815260040161234e95949392919061482e565b6001600560008781526020019081526020016000208585604051612b1892919061436a565b9081526040805160209281900383018120805460ff1916941515949094179093556000888152600590925290208391839190612b57908890889061436a565b90815260200160405180910390206001019182612b75929190614719565b508383600560008881526020019081526020016000208686604051612b9b92919061436a565b90815260200160405180910390206002019182612bb9929190614719565b506002604051806040016040528087815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939094525050835460018181018655948252602091829020845160029092020190815590830151929390929083019150612c399082614867565b50505083836040516112a692919061436a565b6001600160a01b038916612c8c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000879003612cf6576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f53796d626f6c2063616e6e6f7420626520656d70747900000000000000000000604482015260640161234e565b6000849003612d60576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4f726967696e20616464726573732063616e6e6f7420626520656d7074790000604482015260640161234e565b6001600160a01b0389811660009081526006602052604090205461010090041615612dc2576040517f63f4ee1f0000000000000000000000000000000000000000000000000000000081526001600160a01b038a16600482015260240161234e565b60006001600160a01b031660078989604051612ddf92919061436a565b908152604051908190036020019020546001600160a01b031614612e335787876040517fb295cac900000000000000000000000000000000000000000000000000000000815260040161234e929190614501565b6001600160a01b0389166000818152600660205260409020805460017fffffffffffffffffffffff000000000000000000000000000000000000000000909116610100909302929092178217815501612e8d858783614719565b506001600160a01b038916600090815260066020526040902060028101879055600301612ebb888a83614719565b506001600160a01b038916600090815260066020526040902060058101805460ff191660ff8416179055600401612ef3838583614719565b5088600860008881526020019081526020016000208686604051612f1892919061436a565b908152602001604051809103902060006101000a8154816001600160a01b0302191690836001600160a01b031602179055508860078989604051612f5d92919061436a565b90815260405190819003602001902080546001600160a01b039283167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617909155600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180549b9092169a16999099179098555050505050505050565b612ff4612567565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612755565b6001600160a01b038216613087576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382811660009081526006602052604090205461010090041661310c576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5a52433230206e6f742072656769737465726564000000000000000000000000604482015260640161234e565b6001600160a01b03919091166000908152600660205260409020805460ff1916911515919091179055565b60008781526004602052604090205460ff16613182576040517f8e6feba50000000000000000000000000000000000000000000000000000000081526004810188905260240161234e565b60008590036131c05785856040517ec10cfd00000000000000000000000000000000000000000000000000000000815260040161234e929190614501565b6000878152600560205260409081902090516131df908890889061436a565b9081526040519081900360200190205460ff1661322e578686866040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161234e939291906146b8565b8181600560008a8152602001908152602001600020888860405161325392919061436a565b9081526020016040518091039020600301868660405161327492919061436a565b9081526020016040518091039020918261328f929190614719565b5050505050505050565b60008290036132d4576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008581526004602052604090205460ff1680156132ef5750805b15613329576040517fa1452cb00000000000000000000000000000000000000000000000000000000081526004810186905260240161234e565b60008581526004602052604090205460ff16158015613346575080155b15613380576040517f8e6feba50000000000000000000000000000000000000000000000000000000081526004810186905260240161234e565b6000858152600460205260409020600301805461339c9061437a565b90506000036133da576001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6018590555b6000858152600460205260409020805460ff19168215151781556002810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038716179055600301613437838583614719565b50801561347557600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630185905561347e565b61347e8561362a565b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16612463576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff1661111a576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024810183905260440161234e565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16612463576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6135dd826136da565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115613622576110e48282613782565b61111a6137f8565b60005b60005481101561111a57816000828154811061364b5761364b614459565b9060005260206000200154036136d2576000805461366b90600190614962565b8154811061367b5761367b614459565b90600052602060002001546000828154811061369957613699614459565b906000526020600020018190555060008054806136b8576136b861499c565b600190038181906000526020600020016000905590555050565b60010161362d565b806001600160a01b03163b600003613729576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161234e565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b03168460405161379f9190614488565b600060405180830381855af49150503d80600081146137da576040519150601f19603f3d011682016040523d82523d6000602084013e6137df565b606091505b50915091506137ef858383613830565b95945050505050565b3415612463576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608261384557613840826138a5565b611575565b815115801561385c57506001600160a01b0384163b155b1561389e576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161234e565b5080611575565b8051156138b55780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156138f957600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461157557600080fd5b60008083601f84011261393b57600080fd5b50813567ffffffffffffffff81111561395357600080fd5b60208301915083602082850101111561396b57600080fd5b9250929050565b8035801515811461398257600080fd5b919050565b6000806000806060858703121561399d57600080fd5b84359350602085013567ffffffffffffffff8111156139bb57600080fd5b6139c787828801613929565b90945092506139da905060408601613972565b905092959194509250565b80356001600160a01b038116811461398257600080fd5b600080600080600060a08688031215613a1457600080fd5b613a1d866139e5565b9450613a2b602087016139e5565b9350613a39604087016139e5565b9250613a47606087016139e5565b9150613a55608087016139e5565b90509295509295909350565b60005b83811015613a7c578181015183820152602001613a64565b50506000910152565b60008151808452613a9d816020860160208601613a61565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613b66577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160806020880152613b296080880182613a85565b905060408201518782036040890152613b428282613a85565b60609384015198909301979097525094506020938401939190910190600101613ad9565b50929695505050505050565b600080600080600060608688031215613b8a57600080fd5b85359450602086013567ffffffffffffffff811115613ba857600080fd5b613bb488828901613929565b909550935050604086013567ffffffffffffffff811115613bd457600080fd5b613be088828901613929565b969995985093965092949392505050565b600060208284031215613c0357600080fd5b5035919050565b60008060408385031215613c1d57600080fd5b82359150613c2d602084016139e5565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060408385031215613c7857600080fd5b613c81836139e5565b9150602083013567ffffffffffffffff811115613c9d57600080fd5b8301601f81018513613cae57600080fd5b803567ffffffffffffffff811115613cc857613cc8613c36565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff82111715613cf857613cf8613c36565b604052818152828201602001871015613d1057600080fd5b816020840160208301376000602083830101528093505050509250929050565b600080600060408486031215613d4557600080fd5b83359250602084013567ffffffffffffffff811115613d6357600080fd5b613d6f86828701613929565b9497909650939450505050565b8215158152604060208201526000613d976040830184613a85565b949350505050565b60008060008385036040811215613db557600080fd5b6020811215613dc357600080fd5b50839250602084013567ffffffffffffffff811115613d6357600080fd5b6020815260006115756020830184613a85565b803560ff8116811461398257600080fd5b600080600080600080600080600060c08a8c031215613e2357600080fd5b613e2c8a6139e5565b985060208a013567ffffffffffffffff811115613e4857600080fd5b613e548c828d01613929565b90995097505060408a0135955060608a013567ffffffffffffffff811115613e7b57600080fd5b613e878c828d01613929565b90965094505060808a013567ffffffffffffffff811115613ea757600080fd5b613eb38c828d01613929565b9094509250613ec6905060a08b01613df4565b90509295985092959850929598565b60008060408385031215613ee857600080fd5b613ef1836139e5565b9150613c2d60208401613972565b60008083601f840112613f1157600080fd5b50813567ffffffffffffffff811115613f2957600080fd5b6020830191508360208260051b850101111561396b57600080fd5b60008060208385031215613f5757600080fd5b823567ffffffffffffffff811115613f6e57600080fd5b613f7a85828601613eff565b90969095509350505050565b602080825282518282018190526000918401906040840190835b81811015613fbe578351835260209384019390920191600101613fa0565b509095945050505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613b66577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160208701526001600160a01b03604082015116604087015260608101519050608060608701526140606080870182613a85565b9550506020938401939190910190600101613ff1565b6000806000806040858703121561408c57600080fd5b843567ffffffffffffffff8111156140a357600080fd5b6140af87828801613eff565b909550935050602085013567ffffffffffffffff8111156140cf57600080fd5b6140db87828801613eff565b95989497509550505050565b6000806000806000608086880312156140ff57600080fd5b8535945061410f602087016139e5565b9350604086013567ffffffffffffffff81111561412b57600080fd5b61413788828901613929565b9094509250613a55905060608701613972565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613b66577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051151586526001600160a01b036020820151166020870152604081015160e060408801526141d560e0880182613a85565b905060608201516060880152608082015187820360808901526141f88282613a85565b91505060a082015187820360a08901526142128282613a85565b91505060c0820151915061422b60c088018360ff169052565b9550506020938401939190910190600101614172565b60006020828403121561425357600080fd5b611575826139e5565b861515815260c06020820152600061427760c0830188613a85565b866040840152828103606084015261428f8187613a85565b905082810360808401526142a38186613a85565b91505060ff831660a0830152979650505050505050565b60008060008060008060006080888a0312156142d557600080fd5b87359650602088013567ffffffffffffffff8111156142f357600080fd5b6142ff8a828b01613929565b909750955050604088013567ffffffffffffffff81111561431f57600080fd5b61432b8a828b01613929565b909550935050606088013567ffffffffffffffff81111561434b57600080fd5b6143578a828b01613929565b989b979a50959850939692959293505050565b8183823760009101908152919050565b600181811c9082168061438e57607f821691505b6020821081036143c7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b6020815260008083546143df8161437a565b80602086015260018216600081146143fe576001811461441a5761444e565b60ff1983166040870152604082151560051b870101935061444e565b86600052602060002060005b8381101561444557815488820160400152600190910190602001614426565b87016040019450505b509195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000825161449a818460208701613a61565b9190910192915050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b6040815260006144e36040830186886144a4565b82810360208401526144f68185876144a4565b979650505050505050565b602081526000613d976020830184866144a4565b60ff851681528360208201526060604082015260006113ed6060830184866144a4565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2183360301811261449a57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126145a157600080fd5b83018035915067ffffffffffffffff8211156145bc57600080fd5b60200191503681900382131561396b57600080fd5b6000602082840312156145e357600080fd5b61157582613df4565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8183360301811261449a57600080fd5b60006020828403121561463257600080fd5b61157582613972565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa183360301811261449a57600080fd5b60608152600061468360608301888a6144a4565b82810360208401526146968187896144a4565b905082810360408401526146ab8185876144a4565b9998505050505050505050565b8381526040602082015260006137ef6040830184866144a4565b601f8211156110e457806000526020600020601f840160051c810160208510156146f95750805b601f840160051c820191505b8181101561347e5760008155600101614705565b67ffffffffffffffff83111561473157614731613c36565b6147458361473f835461437a565b836146d2565b6000601f84116001811461479757600085156147615750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b17835561347e565b600083815260209020601f19861690835b828110156147c857868501358255602094850194600190920191016147a8565b5086821015614803577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b60006020828403121561482757600080fd5b5051919050565b8581526060602082015260006148486060830186886144a4565b828103604084015261485b8185876144a4565b98975050505050505050565b815167ffffffffffffffff81111561488157614881613c36565b6148958161488f845461437a565b846146d2565b6020601f8211600181146148e757600083156148b15750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b17845561347e565b600084815260208120601f198516915b8281101561491757878501518255602094850194600190920191016148f7565b508482101561495357868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b81810381811115610941577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212209c5d653a556939c3e92a83f9af1b52ded581cfa461637dafe0157d306b6ee78064736f6c634300081a0033",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// RegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegistryMetaData.Bin instead.
var RegistryBin = RegistryMetaData.Bin

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registry *RegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registry *RegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Registry.Contract.DEFAULTADMINROLE(&_Registry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registry *RegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Registry.Contract.DEFAULTADMINROLE(&_Registry.CallOpts)
}

// GATEWAYROLE is a free data retrieval call binding the contract method 0xc50d6dd9.
//
// Solidity: function GATEWAY_ROLE() view returns(bytes32)
func (_Registry *RegistryCaller) GATEWAYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "GATEWAY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GATEWAYROLE is a free data retrieval call binding the contract method 0xc50d6dd9.
//
// Solidity: function GATEWAY_ROLE() view returns(bytes32)
func (_Registry *RegistrySession) GATEWAYROLE() ([32]byte, error) {
	return _Registry.Contract.GATEWAYROLE(&_Registry.CallOpts)
}

// GATEWAYROLE is a free data retrieval call binding the contract method 0xc50d6dd9.
//
// Solidity: function GATEWAY_ROLE() view returns(bytes32)
func (_Registry *RegistryCallerSession) GATEWAYROLE() ([32]byte, error) {
	return _Registry.Contract.GATEWAYROLE(&_Registry.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Registry *RegistryCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Registry *RegistrySession) PAUSERROLE() ([32]byte, error) {
	return _Registry.Contract.PAUSERROLE(&_Registry.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Registry *RegistryCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Registry.Contract.PAUSERROLE(&_Registry.CallOpts)
}

// REGISTRYMANAGERROLE is a free data retrieval call binding the contract method 0xbd8fde1c.
//
// Solidity: function REGISTRY_MANAGER_ROLE() view returns(bytes32)
func (_Registry *RegistryCaller) REGISTRYMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "REGISTRY_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRYMANAGERROLE is a free data retrieval call binding the contract method 0xbd8fde1c.
//
// Solidity: function REGISTRY_MANAGER_ROLE() view returns(bytes32)
func (_Registry *RegistrySession) REGISTRYMANAGERROLE() ([32]byte, error) {
	return _Registry.Contract.REGISTRYMANAGERROLE(&_Registry.CallOpts)
}

// REGISTRYMANAGERROLE is a free data retrieval call binding the contract method 0xbd8fde1c.
//
// Solidity: function REGISTRY_MANAGER_ROLE() view returns(bytes32)
func (_Registry *RegistryCallerSession) REGISTRYMANAGERROLE() ([32]byte, error) {
	return _Registry.Contract.REGISTRYMANAGERROLE(&_Registry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Registry *RegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Registry *RegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Registry.Contract.UPGRADEINTERFACEVERSION(&_Registry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Registry *RegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Registry.Contract.UPGRADEINTERFACEVERSION(&_Registry.CallOpts)
}

// CoreRegistry is a free data retrieval call binding the contract method 0x6bf3d05a.
//
// Solidity: function coreRegistry() view returns(address)
func (_Registry *RegistryCaller) CoreRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "coreRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CoreRegistry is a free data retrieval call binding the contract method 0x6bf3d05a.
//
// Solidity: function coreRegistry() view returns(address)
func (_Registry *RegistrySession) CoreRegistry() (common.Address, error) {
	return _Registry.Contract.CoreRegistry(&_Registry.CallOpts)
}

// CoreRegistry is a free data retrieval call binding the contract method 0x6bf3d05a.
//
// Solidity: function coreRegistry() view returns(address)
func (_Registry *RegistryCallerSession) CoreRegistry() (common.Address, error) {
	return _Registry.Contract.CoreRegistry(&_Registry.CallOpts)
}

// GatewayEVM is a free data retrieval call binding the contract method 0x2c78f74c.
//
// Solidity: function gatewayEVM() view returns(address)
func (_Registry *RegistryCaller) GatewayEVM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "gatewayEVM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayEVM is a free data retrieval call binding the contract method 0x2c78f74c.
//
// Solidity: function gatewayEVM() view returns(address)
func (_Registry *RegistrySession) GatewayEVM() (common.Address, error) {
	return _Registry.Contract.GatewayEVM(&_Registry.CallOpts)
}

// GatewayEVM is a free data retrieval call binding the contract method 0x2c78f74c.
//
// Solidity: function gatewayEVM() view returns(address)
func (_Registry *RegistryCallerSession) GatewayEVM() (common.Address, error) {
	return _Registry.Contract.GatewayEVM(&_Registry.CallOpts)
}

// GetActiveChains is a free data retrieval call binding the contract method 0x94cc8683.
//
// Solidity: function getActiveChains() view returns(uint256[])
func (_Registry *RegistryCaller) GetActiveChains(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getActiveChains")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveChains is a free data retrieval call binding the contract method 0x94cc8683.
//
// Solidity: function getActiveChains() view returns(uint256[])
func (_Registry *RegistrySession) GetActiveChains() ([]*big.Int, error) {
	return _Registry.Contract.GetActiveChains(&_Registry.CallOpts)
}

// GetActiveChains is a free data retrieval call binding the contract method 0x94cc8683.
//
// Solidity: function getActiveChains() view returns(uint256[])
func (_Registry *RegistryCallerSession) GetActiveChains() ([]*big.Int, error) {
	return _Registry.Contract.GetActiveChains(&_Registry.CallOpts)
}

// GetAllChains is a free data retrieval call binding the contract method 0x9ca220dd.
//
// Solidity: function getAllChains() view returns((bool,uint256,address,bytes)[] chainsInfo)
func (_Registry *RegistryCaller) GetAllChains(opts *bind.CallOpts) ([]ChainInfoDTO, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAllChains")

	if err != nil {
		return *new([]ChainInfoDTO), err
	}

	out0 := *abi.ConvertType(out[0], new([]ChainInfoDTO)).(*[]ChainInfoDTO)

	return out0, err

}

// GetAllChains is a free data retrieval call binding the contract method 0x9ca220dd.
//
// Solidity: function getAllChains() view returns((bool,uint256,address,bytes)[] chainsInfo)
func (_Registry *RegistrySession) GetAllChains() ([]ChainInfoDTO, error) {
	return _Registry.Contract.GetAllChains(&_Registry.CallOpts)
}

// GetAllChains is a free data retrieval call binding the contract method 0x9ca220dd.
//
// Solidity: function getAllChains() view returns((bool,uint256,address,bytes)[] chainsInfo)
func (_Registry *RegistryCallerSession) GetAllChains() ([]ChainInfoDTO, error) {
	return _Registry.Contract.GetAllChains(&_Registry.CallOpts)
}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns((bool,bytes,string,uint256)[] contractsInfo)
func (_Registry *RegistryCaller) GetAllContracts(opts *bind.CallOpts) ([]ContractInfoDTO, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAllContracts")

	if err != nil {
		return *new([]ContractInfoDTO), err
	}

	out0 := *abi.ConvertType(out[0], new([]ContractInfoDTO)).(*[]ContractInfoDTO)

	return out0, err

}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns((bool,bytes,string,uint256)[] contractsInfo)
func (_Registry *RegistrySession) GetAllContracts() ([]ContractInfoDTO, error) {
	return _Registry.Contract.GetAllContracts(&_Registry.CallOpts)
}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns((bool,bytes,string,uint256)[] contractsInfo)
func (_Registry *RegistryCallerSession) GetAllContracts() ([]ContractInfoDTO, error) {
	return _Registry.Contract.GetAllContracts(&_Registry.CallOpts)
}

// GetAllZRC20Tokens is a free data retrieval call binding the contract method 0xc1bd469f.
//
// Solidity: function getAllZRC20Tokens() view returns((bool,address,bytes,uint256,string,string,uint8)[] tokensInfo)
func (_Registry *RegistryCaller) GetAllZRC20Tokens(opts *bind.CallOpts) ([]ZRC20Info, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAllZRC20Tokens")

	if err != nil {
		return *new([]ZRC20Info), err
	}

	out0 := *abi.ConvertType(out[0], new([]ZRC20Info)).(*[]ZRC20Info)

	return out0, err

}

// GetAllZRC20Tokens is a free data retrieval call binding the contract method 0xc1bd469f.
//
// Solidity: function getAllZRC20Tokens() view returns((bool,address,bytes,uint256,string,string,uint8)[] tokensInfo)
func (_Registry *RegistrySession) GetAllZRC20Tokens() ([]ZRC20Info, error) {
	return _Registry.Contract.GetAllZRC20Tokens(&_Registry.CallOpts)
}

// GetAllZRC20Tokens is a free data retrieval call binding the contract method 0xc1bd469f.
//
// Solidity: function getAllZRC20Tokens() view returns((bool,address,bytes,uint256,string,string,uint8)[] tokensInfo)
func (_Registry *RegistryCallerSession) GetAllZRC20Tokens() ([]ZRC20Info, error) {
	return _Registry.Contract.GetAllZRC20Tokens(&_Registry.CallOpts)
}

// GetChainMetadata is a free data retrieval call binding the contract method 0x7066b18d.
//
// Solidity: function getChainMetadata(uint256 chainId, string key) view returns(bytes)
func (_Registry *RegistryCaller) GetChainMetadata(opts *bind.CallOpts, chainId *big.Int, key string) ([]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getChainMetadata", chainId, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetChainMetadata is a free data retrieval call binding the contract method 0x7066b18d.
//
// Solidity: function getChainMetadata(uint256 chainId, string key) view returns(bytes)
func (_Registry *RegistrySession) GetChainMetadata(chainId *big.Int, key string) ([]byte, error) {
	return _Registry.Contract.GetChainMetadata(&_Registry.CallOpts, chainId, key)
}

// GetChainMetadata is a free data retrieval call binding the contract method 0x7066b18d.
//
// Solidity: function getChainMetadata(uint256 chainId, string key) view returns(bytes)
func (_Registry *RegistryCallerSession) GetChainMetadata(chainId *big.Int, key string) ([]byte, error) {
	return _Registry.Contract.GetChainMetadata(&_Registry.CallOpts, chainId, key)
}

// GetContractConfiguration is a free data retrieval call binding the contract method 0xd3523ea2.
//
// Solidity: function getContractConfiguration(uint256 chainId, string contractType, string key) view returns(bytes)
func (_Registry *RegistryCaller) GetContractConfiguration(opts *bind.CallOpts, chainId *big.Int, contractType string, key string) ([]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getContractConfiguration", chainId, contractType, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetContractConfiguration is a free data retrieval call binding the contract method 0xd3523ea2.
//
// Solidity: function getContractConfiguration(uint256 chainId, string contractType, string key) view returns(bytes)
func (_Registry *RegistrySession) GetContractConfiguration(chainId *big.Int, contractType string, key string) ([]byte, error) {
	return _Registry.Contract.GetContractConfiguration(&_Registry.CallOpts, chainId, contractType, key)
}

// GetContractConfiguration is a free data retrieval call binding the contract method 0xd3523ea2.
//
// Solidity: function getContractConfiguration(uint256 chainId, string contractType, string key) view returns(bytes)
func (_Registry *RegistryCallerSession) GetContractConfiguration(chainId *big.Int, contractType string, key string) ([]byte, error) {
	return _Registry.Contract.GetContractConfiguration(&_Registry.CallOpts, chainId, contractType, key)
}

// GetContractInfo is a free data retrieval call binding the contract method 0x5cf92c9f.
//
// Solidity: function getContractInfo(uint256 chainId, string contractType) view returns(bool active, bytes addressBytes)
func (_Registry *RegistryCaller) GetContractInfo(opts *bind.CallOpts, chainId *big.Int, contractType string) (struct {
	Active       bool
	AddressBytes []byte
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getContractInfo", chainId, contractType)

	outstruct := new(struct {
		Active       bool
		AddressBytes []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.AddressBytes = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetContractInfo is a free data retrieval call binding the contract method 0x5cf92c9f.
//
// Solidity: function getContractInfo(uint256 chainId, string contractType) view returns(bool active, bytes addressBytes)
func (_Registry *RegistrySession) GetContractInfo(chainId *big.Int, contractType string) (struct {
	Active       bool
	AddressBytes []byte
}, error) {
	return _Registry.Contract.GetContractInfo(&_Registry.CallOpts, chainId, contractType)
}

// GetContractInfo is a free data retrieval call binding the contract method 0x5cf92c9f.
//
// Solidity: function getContractInfo(uint256 chainId, string contractType) view returns(bool active, bytes addressBytes)
func (_Registry *RegistryCallerSession) GetContractInfo(chainId *big.Int, contractType string) (struct {
	Active       bool
	AddressBytes []byte
}, error) {
	return _Registry.Contract.GetContractInfo(&_Registry.CallOpts, chainId, contractType)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registry *RegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registry *RegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Registry.Contract.GetRoleAdmin(&_Registry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registry *RegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Registry.Contract.GetRoleAdmin(&_Registry.CallOpts, role)
}

// GetZRC20AddressByForeignAsset is a free data retrieval call binding the contract method 0xaa808c06.
//
// Solidity: function getZRC20AddressByForeignAsset(uint256 originChainId, bytes originAddress) view returns(address)
func (_Registry *RegistryCaller) GetZRC20AddressByForeignAsset(opts *bind.CallOpts, originChainId *big.Int, originAddress []byte) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getZRC20AddressByForeignAsset", originChainId, originAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetZRC20AddressByForeignAsset is a free data retrieval call binding the contract method 0xaa808c06.
//
// Solidity: function getZRC20AddressByForeignAsset(uint256 originChainId, bytes originAddress) view returns(address)
func (_Registry *RegistrySession) GetZRC20AddressByForeignAsset(originChainId *big.Int, originAddress []byte) (common.Address, error) {
	return _Registry.Contract.GetZRC20AddressByForeignAsset(&_Registry.CallOpts, originChainId, originAddress)
}

// GetZRC20AddressByForeignAsset is a free data retrieval call binding the contract method 0xaa808c06.
//
// Solidity: function getZRC20AddressByForeignAsset(uint256 originChainId, bytes originAddress) view returns(address)
func (_Registry *RegistryCallerSession) GetZRC20AddressByForeignAsset(originChainId *big.Int, originAddress []byte) (common.Address, error) {
	return _Registry.Contract.GetZRC20AddressByForeignAsset(&_Registry.CallOpts, originChainId, originAddress)
}

// GetZRC20TokenInfo is a free data retrieval call binding the contract method 0xe9d6c5ba.
//
// Solidity: function getZRC20TokenInfo(address address_) view returns(bool active, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals)
func (_Registry *RegistryCaller) GetZRC20TokenInfo(opts *bind.CallOpts, address_ common.Address) (struct {
	Active        bool
	Symbol        string
	OriginChainId *big.Int
	OriginAddress []byte
	CoinType      string
	Decimals      uint8
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getZRC20TokenInfo", address_)

	outstruct := new(struct {
		Active        bool
		Symbol        string
		OriginChainId *big.Int
		OriginAddress []byte
		CoinType      string
		Decimals      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.OriginChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OriginAddress = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.CoinType = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Decimals = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetZRC20TokenInfo is a free data retrieval call binding the contract method 0xe9d6c5ba.
//
// Solidity: function getZRC20TokenInfo(address address_) view returns(bool active, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals)
func (_Registry *RegistrySession) GetZRC20TokenInfo(address_ common.Address) (struct {
	Active        bool
	Symbol        string
	OriginChainId *big.Int
	OriginAddress []byte
	CoinType      string
	Decimals      uint8
}, error) {
	return _Registry.Contract.GetZRC20TokenInfo(&_Registry.CallOpts, address_)
}

// GetZRC20TokenInfo is a free data retrieval call binding the contract method 0xe9d6c5ba.
//
// Solidity: function getZRC20TokenInfo(address address_) view returns(bool active, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals)
func (_Registry *RegistryCallerSession) GetZRC20TokenInfo(address_ common.Address) (struct {
	Active        bool
	Symbol        string
	OriginChainId *big.Int
	OriginAddress []byte
	CoinType      string
	Decimals      uint8
}, error) {
	return _Registry.Contract.GetZRC20TokenInfo(&_Registry.CallOpts, address_)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Registry.Contract.HasRole(&_Registry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Registry.Contract.HasRole(&_Registry.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registry *RegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registry *RegistrySession) Paused() (bool, error) {
	return _Registry.Contract.Paused(&_Registry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registry *RegistryCallerSession) Paused() (bool, error) {
	return _Registry.Contract.Paused(&_Registry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Registry *RegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Registry *RegistrySession) ProxiableUUID() ([32]byte, error) {
	return _Registry.Contract.ProxiableUUID(&_Registry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Registry *RegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Registry.Contract.ProxiableUUID(&_Registry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registry.Contract.SupportsInterface(&_Registry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registry.Contract.SupportsInterface(&_Registry.CallOpts, interfaceId)
}

// BootstrapChains is a paid mutator transaction binding the contract method 0xdc07f866.
//
// Solidity: function bootstrapChains((bool,uint256,address,bytes)[] chains, (uint256,string,bytes)[] metadataEntries) returns()
func (_Registry *RegistryTransactor) BootstrapChains(opts *bind.TransactOpts, chains []ChainInfoDTO, metadataEntries []IRegistryChainMetadataEntry) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "bootstrapChains", chains, metadataEntries)
}

// BootstrapChains is a paid mutator transaction binding the contract method 0xdc07f866.
//
// Solidity: function bootstrapChains((bool,uint256,address,bytes)[] chains, (uint256,string,bytes)[] metadataEntries) returns()
func (_Registry *RegistrySession) BootstrapChains(chains []ChainInfoDTO, metadataEntries []IRegistryChainMetadataEntry) (*types.Transaction, error) {
	return _Registry.Contract.BootstrapChains(&_Registry.TransactOpts, chains, metadataEntries)
}

// BootstrapChains is a paid mutator transaction binding the contract method 0xdc07f866.
//
// Solidity: function bootstrapChains((bool,uint256,address,bytes)[] chains, (uint256,string,bytes)[] metadataEntries) returns()
func (_Registry *RegistryTransactorSession) BootstrapChains(chains []ChainInfoDTO, metadataEntries []IRegistryChainMetadataEntry) (*types.Transaction, error) {
	return _Registry.Contract.BootstrapChains(&_Registry.TransactOpts, chains, metadataEntries)
}

// BootstrapContracts is a paid mutator transaction binding the contract method 0xa4b7e5f9.
//
// Solidity: function bootstrapContracts((bool,bytes,string,uint256)[] contracts, (uint256,string,string,bytes)[] configEntries) returns()
func (_Registry *RegistryTransactor) BootstrapContracts(opts *bind.TransactOpts, contracts []ContractInfoDTO, configEntries []IRegistryContractConfigEntry) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "bootstrapContracts", contracts, configEntries)
}

// BootstrapContracts is a paid mutator transaction binding the contract method 0xa4b7e5f9.
//
// Solidity: function bootstrapContracts((bool,bytes,string,uint256)[] contracts, (uint256,string,string,bytes)[] configEntries) returns()
func (_Registry *RegistrySession) BootstrapContracts(contracts []ContractInfoDTO, configEntries []IRegistryContractConfigEntry) (*types.Transaction, error) {
	return _Registry.Contract.BootstrapContracts(&_Registry.TransactOpts, contracts, configEntries)
}

// BootstrapContracts is a paid mutator transaction binding the contract method 0xa4b7e5f9.
//
// Solidity: function bootstrapContracts((bool,bytes,string,uint256)[] contracts, (uint256,string,string,bytes)[] configEntries) returns()
func (_Registry *RegistryTransactorSession) BootstrapContracts(contracts []ContractInfoDTO, configEntries []IRegistryContractConfigEntry) (*types.Transaction, error) {
	return _Registry.Contract.BootstrapContracts(&_Registry.TransactOpts, contracts, configEntries)
}

// BootstrapZRC20Tokens is a paid mutator transaction binding the contract method 0x909b6a03.
//
// Solidity: function bootstrapZRC20Tokens((bool,address,bytes,uint256,string,string,uint8)[] tokens) returns()
func (_Registry *RegistryTransactor) BootstrapZRC20Tokens(opts *bind.TransactOpts, tokens []ZRC20Info) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "bootstrapZRC20Tokens", tokens)
}

// BootstrapZRC20Tokens is a paid mutator transaction binding the contract method 0x909b6a03.
//
// Solidity: function bootstrapZRC20Tokens((bool,address,bytes,uint256,string,string,uint8)[] tokens) returns()
func (_Registry *RegistrySession) BootstrapZRC20Tokens(tokens []ZRC20Info) (*types.Transaction, error) {
	return _Registry.Contract.BootstrapZRC20Tokens(&_Registry.TransactOpts, tokens)
}

// BootstrapZRC20Tokens is a paid mutator transaction binding the contract method 0x909b6a03.
//
// Solidity: function bootstrapZRC20Tokens((bool,address,bytes,uint256,string,string,uint8)[] tokens) returns()
func (_Registry *RegistryTransactorSession) BootstrapZRC20Tokens(tokens []ZRC20Info) (*types.Transaction, error) {
	return _Registry.Contract.BootstrapZRC20Tokens(&_Registry.TransactOpts, tokens)
}

// ChangeChainStatus is a paid mutator transaction binding the contract method 0xa8f2cb96.
//
// Solidity: function changeChainStatus(uint256 chainId, address gasZRC20, bytes registry, bool activation) returns()
func (_Registry *RegistryTransactor) ChangeChainStatus(opts *bind.TransactOpts, chainId *big.Int, gasZRC20 common.Address, registry []byte, activation bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "changeChainStatus", chainId, gasZRC20, registry, activation)
}

// ChangeChainStatus is a paid mutator transaction binding the contract method 0xa8f2cb96.
//
// Solidity: function changeChainStatus(uint256 chainId, address gasZRC20, bytes registry, bool activation) returns()
func (_Registry *RegistrySession) ChangeChainStatus(chainId *big.Int, gasZRC20 common.Address, registry []byte, activation bool) (*types.Transaction, error) {
	return _Registry.Contract.ChangeChainStatus(&_Registry.TransactOpts, chainId, gasZRC20, registry, activation)
}

// ChangeChainStatus is a paid mutator transaction binding the contract method 0xa8f2cb96.
//
// Solidity: function changeChainStatus(uint256 chainId, address gasZRC20, bytes registry, bool activation) returns()
func (_Registry *RegistryTransactorSession) ChangeChainStatus(chainId *big.Int, gasZRC20 common.Address, registry []byte, activation bool) (*types.Transaction, error) {
	return _Registry.Contract.ChangeChainStatus(&_Registry.TransactOpts, chainId, gasZRC20, registry, activation)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.GrantRole(&_Registry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.GrantRole(&_Registry.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address admin_, address pauserAddress_, address registryManager_, address gatewayEVM_, address coreRegistry_) returns()
func (_Registry *RegistryTransactor) Initialize(opts *bind.TransactOpts, admin_ common.Address, pauserAddress_ common.Address, registryManager_ common.Address, gatewayEVM_ common.Address, coreRegistry_ common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "initialize", admin_, pauserAddress_, registryManager_, gatewayEVM_, coreRegistry_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address admin_, address pauserAddress_, address registryManager_, address gatewayEVM_, address coreRegistry_) returns()
func (_Registry *RegistrySession) Initialize(admin_ common.Address, pauserAddress_ common.Address, registryManager_ common.Address, gatewayEVM_ common.Address, coreRegistry_ common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, admin_, pauserAddress_, registryManager_, gatewayEVM_, coreRegistry_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address admin_, address pauserAddress_, address registryManager_, address gatewayEVM_, address coreRegistry_) returns()
func (_Registry *RegistryTransactorSession) Initialize(admin_ common.Address, pauserAddress_ common.Address, registryManager_ common.Address, gatewayEVM_ common.Address, coreRegistry_ common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, admin_, pauserAddress_, registryManager_, gatewayEVM_, coreRegistry_)
}

// OnCall is a paid mutator transaction binding the contract method 0x676cc054.
//
// Solidity: function onCall((address) context, bytes data) returns(bytes)
func (_Registry *RegistryTransactor) OnCall(opts *bind.TransactOpts, context MessageContext, data []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "onCall", context, data)
}

// OnCall is a paid mutator transaction binding the contract method 0x676cc054.
//
// Solidity: function onCall((address) context, bytes data) returns(bytes)
func (_Registry *RegistrySession) OnCall(context MessageContext, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.OnCall(&_Registry.TransactOpts, context, data)
}

// OnCall is a paid mutator transaction binding the contract method 0x676cc054.
//
// Solidity: function onCall((address) context, bytes data) returns(bytes)
func (_Registry *RegistryTransactorSession) OnCall(context MessageContext, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.OnCall(&_Registry.TransactOpts, context, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registry *RegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registry *RegistrySession) Pause() (*types.Transaction, error) {
	return _Registry.Contract.Pause(&_Registry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registry *RegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _Registry.Contract.Pause(&_Registry.TransactOpts)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x631d62e4.
//
// Solidity: function registerContract(uint256 chainId, string contractType, bytes addressBytes) returns()
func (_Registry *RegistryTransactor) RegisterContract(opts *bind.TransactOpts, chainId *big.Int, contractType string, addressBytes []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerContract", chainId, contractType, addressBytes)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x631d62e4.
//
// Solidity: function registerContract(uint256 chainId, string contractType, bytes addressBytes) returns()
func (_Registry *RegistrySession) RegisterContract(chainId *big.Int, contractType string, addressBytes []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterContract(&_Registry.TransactOpts, chainId, contractType, addressBytes)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x631d62e4.
//
// Solidity: function registerContract(uint256 chainId, string contractType, bytes addressBytes) returns()
func (_Registry *RegistryTransactorSession) RegisterContract(chainId *big.Int, contractType string, addressBytes []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterContract(&_Registry.TransactOpts, chainId, contractType, addressBytes)
}

// RegisterZRC20Token is a paid mutator transaction binding the contract method 0x6e9e2d3f.
//
// Solidity: function registerZRC20Token(address address_, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals) returns()
func (_Registry *RegistryTransactor) RegisterZRC20Token(opts *bind.TransactOpts, address_ common.Address, symbol string, originChainId *big.Int, originAddress []byte, coinType string, decimals uint8) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerZRC20Token", address_, symbol, originChainId, originAddress, coinType, decimals)
}

// RegisterZRC20Token is a paid mutator transaction binding the contract method 0x6e9e2d3f.
//
// Solidity: function registerZRC20Token(address address_, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals) returns()
func (_Registry *RegistrySession) RegisterZRC20Token(address_ common.Address, symbol string, originChainId *big.Int, originAddress []byte, coinType string, decimals uint8) (*types.Transaction, error) {
	return _Registry.Contract.RegisterZRC20Token(&_Registry.TransactOpts, address_, symbol, originChainId, originAddress, coinType, decimals)
}

// RegisterZRC20Token is a paid mutator transaction binding the contract method 0x6e9e2d3f.
//
// Solidity: function registerZRC20Token(address address_, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals) returns()
func (_Registry *RegistryTransactorSession) RegisterZRC20Token(address_ common.Address, symbol string, originChainId *big.Int, originAddress []byte, coinType string, decimals uint8) (*types.Transaction, error) {
	return _Registry.Contract.RegisterZRC20Token(&_Registry.TransactOpts, address_, symbol, originChainId, originAddress, coinType, decimals)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Registry *RegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Registry *RegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RenounceRole(&_Registry.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Registry *RegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RenounceRole(&_Registry.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RevokeRole(&_Registry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RevokeRole(&_Registry.TransactOpts, role, account)
}

// SetContractActive is a paid mutator transaction binding the contract method 0x10d29b9e.
//
// Solidity: function setContractActive(uint256 chainId, string contractType, bool active) returns()
func (_Registry *RegistryTransactor) SetContractActive(opts *bind.TransactOpts, chainId *big.Int, contractType string, active bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setContractActive", chainId, contractType, active)
}

// SetContractActive is a paid mutator transaction binding the contract method 0x10d29b9e.
//
// Solidity: function setContractActive(uint256 chainId, string contractType, bool active) returns()
func (_Registry *RegistrySession) SetContractActive(chainId *big.Int, contractType string, active bool) (*types.Transaction, error) {
	return _Registry.Contract.SetContractActive(&_Registry.TransactOpts, chainId, contractType, active)
}

// SetContractActive is a paid mutator transaction binding the contract method 0x10d29b9e.
//
// Solidity: function setContractActive(uint256 chainId, string contractType, bool active) returns()
func (_Registry *RegistryTransactorSession) SetContractActive(chainId *big.Int, contractType string, active bool) (*types.Transaction, error) {
	return _Registry.Contract.SetContractActive(&_Registry.TransactOpts, chainId, contractType, active)
}

// SetZRC20TokenActive is a paid mutator transaction binding the contract method 0x9060bda9.
//
// Solidity: function setZRC20TokenActive(address address_, bool active) returns()
func (_Registry *RegistryTransactor) SetZRC20TokenActive(opts *bind.TransactOpts, address_ common.Address, active bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setZRC20TokenActive", address_, active)
}

// SetZRC20TokenActive is a paid mutator transaction binding the contract method 0x9060bda9.
//
// Solidity: function setZRC20TokenActive(address address_, bool active) returns()
func (_Registry *RegistrySession) SetZRC20TokenActive(address_ common.Address, active bool) (*types.Transaction, error) {
	return _Registry.Contract.SetZRC20TokenActive(&_Registry.TransactOpts, address_, active)
}

// SetZRC20TokenActive is a paid mutator transaction binding the contract method 0x9060bda9.
//
// Solidity: function setZRC20TokenActive(address address_, bool active) returns()
func (_Registry *RegistryTransactorSession) SetZRC20TokenActive(address_ common.Address, active bool) (*types.Transaction, error) {
	return _Registry.Contract.SetZRC20TokenActive(&_Registry.TransactOpts, address_, active)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registry *RegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registry *RegistrySession) Unpause() (*types.Transaction, error) {
	return _Registry.Contract.Unpause(&_Registry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registry *RegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _Registry.Contract.Unpause(&_Registry.TransactOpts)
}

// UpdateChainMetadata is a paid mutator transaction binding the contract method 0x2259e9e5.
//
// Solidity: function updateChainMetadata(uint256 chainId, string key, bytes value) returns()
func (_Registry *RegistryTransactor) UpdateChainMetadata(opts *bind.TransactOpts, chainId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateChainMetadata", chainId, key, value)
}

// UpdateChainMetadata is a paid mutator transaction binding the contract method 0x2259e9e5.
//
// Solidity: function updateChainMetadata(uint256 chainId, string key, bytes value) returns()
func (_Registry *RegistrySession) UpdateChainMetadata(chainId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateChainMetadata(&_Registry.TransactOpts, chainId, key, value)
}

// UpdateChainMetadata is a paid mutator transaction binding the contract method 0x2259e9e5.
//
// Solidity: function updateChainMetadata(uint256 chainId, string key, bytes value) returns()
func (_Registry *RegistryTransactorSession) UpdateChainMetadata(chainId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateChainMetadata(&_Registry.TransactOpts, chainId, key, value)
}

// UpdateContractConfiguration is a paid mutator transaction binding the contract method 0xf354b31f.
//
// Solidity: function updateContractConfiguration(uint256 chainId, string contractType, string key, bytes value) returns()
func (_Registry *RegistryTransactor) UpdateContractConfiguration(opts *bind.TransactOpts, chainId *big.Int, contractType string, key string, value []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateContractConfiguration", chainId, contractType, key, value)
}

// UpdateContractConfiguration is a paid mutator transaction binding the contract method 0xf354b31f.
//
// Solidity: function updateContractConfiguration(uint256 chainId, string contractType, string key, bytes value) returns()
func (_Registry *RegistrySession) UpdateContractConfiguration(chainId *big.Int, contractType string, key string, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateContractConfiguration(&_Registry.TransactOpts, chainId, contractType, key, value)
}

// UpdateContractConfiguration is a paid mutator transaction binding the contract method 0xf354b31f.
//
// Solidity: function updateContractConfiguration(uint256 chainId, string contractType, string key, bytes value) returns()
func (_Registry *RegistryTransactorSession) UpdateContractConfiguration(chainId *big.Int, contractType string, key string, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateContractConfiguration(&_Registry.TransactOpts, chainId, contractType, key, value)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Registry *RegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Registry *RegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpgradeToAndCall(&_Registry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Registry *RegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpgradeToAndCall(&_Registry.TransactOpts, newImplementation, data)
}

// RegistryChainMetadataUpdatedIterator is returned from FilterChainMetadataUpdated and is used to iterate over the raw logs and unpacked data for ChainMetadataUpdated events raised by the Registry contract.
type RegistryChainMetadataUpdatedIterator struct {
	Event *RegistryChainMetadataUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryChainMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryChainMetadataUpdated)
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
		it.Event = new(RegistryChainMetadataUpdated)
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
func (it *RegistryChainMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryChainMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryChainMetadataUpdated represents a ChainMetadataUpdated event raised by the Registry contract.
type RegistryChainMetadataUpdated struct {
	ChainId *big.Int
	Key     string
	Value   []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainMetadataUpdated is a free log retrieval operation binding the contract event 0x40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee634.
//
// Solidity: event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value)
func (_Registry *RegistryFilterer) FilterChainMetadataUpdated(opts *bind.FilterOpts, chainId []*big.Int) (*RegistryChainMetadataUpdatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ChainMetadataUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryChainMetadataUpdatedIterator{contract: _Registry.contract, event: "ChainMetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchChainMetadataUpdated is a free log subscription operation binding the contract event 0x40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee634.
//
// Solidity: event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value)
func (_Registry *RegistryFilterer) WatchChainMetadataUpdated(opts *bind.WatchOpts, sink chan<- *RegistryChainMetadataUpdated, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ChainMetadataUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryChainMetadataUpdated)
				if err := _Registry.contract.UnpackLog(event, "ChainMetadataUpdated", log); err != nil {
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

// ParseChainMetadataUpdated is a log parse operation binding the contract event 0x40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee634.
//
// Solidity: event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value)
func (_Registry *RegistryFilterer) ParseChainMetadataUpdated(log types.Log) (*RegistryChainMetadataUpdated, error) {
	event := new(RegistryChainMetadataUpdated)
	if err := _Registry.contract.UnpackLog(event, "ChainMetadataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryChainStatusChangedIterator is returned from FilterChainStatusChanged and is used to iterate over the raw logs and unpacked data for ChainStatusChanged events raised by the Registry contract.
type RegistryChainStatusChangedIterator struct {
	Event *RegistryChainStatusChanged // Event containing the contract specifics and raw log

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
func (it *RegistryChainStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryChainStatusChanged)
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
		it.Event = new(RegistryChainStatusChanged)
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
func (it *RegistryChainStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryChainStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryChainStatusChanged represents a ChainStatusChanged event raised by the Registry contract.
type RegistryChainStatusChanged struct {
	ChainId   *big.Int
	NewStatus bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChainStatusChanged is a free log retrieval operation binding the contract event 0xc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId, bool newStatus)
func (_Registry *RegistryFilterer) FilterChainStatusChanged(opts *bind.FilterOpts, chainId []*big.Int) (*RegistryChainStatusChangedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ChainStatusChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryChainStatusChangedIterator{contract: _Registry.contract, event: "ChainStatusChanged", logs: logs, sub: sub}, nil
}

// WatchChainStatusChanged is a free log subscription operation binding the contract event 0xc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId, bool newStatus)
func (_Registry *RegistryFilterer) WatchChainStatusChanged(opts *bind.WatchOpts, sink chan<- *RegistryChainStatusChanged, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ChainStatusChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryChainStatusChanged)
				if err := _Registry.contract.UnpackLog(event, "ChainStatusChanged", log); err != nil {
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

// ParseChainStatusChanged is a log parse operation binding the contract event 0xc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId, bool newStatus)
func (_Registry *RegistryFilterer) ParseChainStatusChanged(log types.Log) (*RegistryChainStatusChanged, error) {
	event := new(RegistryChainStatusChanged)
	if err := _Registry.contract.UnpackLog(event, "ChainStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryContractConfigurationUpdatedIterator is returned from FilterContractConfigurationUpdated and is used to iterate over the raw logs and unpacked data for ContractConfigurationUpdated events raised by the Registry contract.
type RegistryContractConfigurationUpdatedIterator struct {
	Event *RegistryContractConfigurationUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryContractConfigurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractConfigurationUpdated)
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
		it.Event = new(RegistryContractConfigurationUpdated)
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
func (it *RegistryContractConfigurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractConfigurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractConfigurationUpdated represents a ContractConfigurationUpdated event raised by the Registry contract.
type RegistryContractConfigurationUpdated struct {
	ChainId      *big.Int
	ContractType string
	Key          string
	Value        []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractConfigurationUpdated is a free log retrieval operation binding the contract event 0xaea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c5.
//
// Solidity: event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value)
func (_Registry *RegistryFilterer) FilterContractConfigurationUpdated(opts *bind.FilterOpts, chainId []*big.Int) (*RegistryContractConfigurationUpdatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ContractConfigurationUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryContractConfigurationUpdatedIterator{contract: _Registry.contract, event: "ContractConfigurationUpdated", logs: logs, sub: sub}, nil
}

// WatchContractConfigurationUpdated is a free log subscription operation binding the contract event 0xaea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c5.
//
// Solidity: event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value)
func (_Registry *RegistryFilterer) WatchContractConfigurationUpdated(opts *bind.WatchOpts, sink chan<- *RegistryContractConfigurationUpdated, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ContractConfigurationUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractConfigurationUpdated)
				if err := _Registry.contract.UnpackLog(event, "ContractConfigurationUpdated", log); err != nil {
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

// ParseContractConfigurationUpdated is a log parse operation binding the contract event 0xaea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c5.
//
// Solidity: event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value)
func (_Registry *RegistryFilterer) ParseContractConfigurationUpdated(log types.Log) (*RegistryContractConfigurationUpdated, error) {
	event := new(RegistryContractConfigurationUpdated)
	if err := _Registry.contract.UnpackLog(event, "ContractConfigurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryContractRegisteredIterator is returned from FilterContractRegistered and is used to iterate over the raw logs and unpacked data for ContractRegistered events raised by the Registry contract.
type RegistryContractRegisteredIterator struct {
	Event *RegistryContractRegistered // Event containing the contract specifics and raw log

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
func (it *RegistryContractRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractRegistered)
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
		it.Event = new(RegistryContractRegistered)
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
func (it *RegistryContractRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractRegistered represents a ContractRegistered event raised by the Registry contract.
type RegistryContractRegistered struct {
	ChainId      *big.Int
	ContractType common.Hash
	AddressBytes []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractRegistered is a free log retrieval operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_Registry *RegistryFilterer) FilterContractRegistered(opts *bind.FilterOpts, chainId []*big.Int, contractType []string) (*RegistryContractRegisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ContractRegistered", chainIdRule, contractTypeRule)
	if err != nil {
		return nil, err
	}
	return &RegistryContractRegisteredIterator{contract: _Registry.contract, event: "ContractRegistered", logs: logs, sub: sub}, nil
}

// WatchContractRegistered is a free log subscription operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_Registry *RegistryFilterer) WatchContractRegistered(opts *bind.WatchOpts, sink chan<- *RegistryContractRegistered, chainId []*big.Int, contractType []string) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ContractRegistered", chainIdRule, contractTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractRegistered)
				if err := _Registry.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
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

// ParseContractRegistered is a log parse operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_Registry *RegistryFilterer) ParseContractRegistered(log types.Log) (*RegistryContractRegistered, error) {
	event := new(RegistryContractRegistered)
	if err := _Registry.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryContractStatusChangedIterator is returned from FilterContractStatusChanged and is used to iterate over the raw logs and unpacked data for ContractStatusChanged events raised by the Registry contract.
type RegistryContractStatusChangedIterator struct {
	Event *RegistryContractStatusChanged // Event containing the contract specifics and raw log

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
func (it *RegistryContractStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractStatusChanged)
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
		it.Event = new(RegistryContractStatusChanged)
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
func (it *RegistryContractStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractStatusChanged represents a ContractStatusChanged event raised by the Registry contract.
type RegistryContractStatusChanged struct {
	AddressBytes []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractStatusChanged is a free log retrieval operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_Registry *RegistryFilterer) FilterContractStatusChanged(opts *bind.FilterOpts) (*RegistryContractStatusChangedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ContractStatusChanged")
	if err != nil {
		return nil, err
	}
	return &RegistryContractStatusChangedIterator{contract: _Registry.contract, event: "ContractStatusChanged", logs: logs, sub: sub}, nil
}

// WatchContractStatusChanged is a free log subscription operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_Registry *RegistryFilterer) WatchContractStatusChanged(opts *bind.WatchOpts, sink chan<- *RegistryContractStatusChanged) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ContractStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractStatusChanged)
				if err := _Registry.contract.UnpackLog(event, "ContractStatusChanged", log); err != nil {
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

// ParseContractStatusChanged is a log parse operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_Registry *RegistryFilterer) ParseContractStatusChanged(log types.Log) (*RegistryContractStatusChanged, error) {
	event := new(RegistryContractStatusChanged)
	if err := _Registry.contract.UnpackLog(event, "ContractStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Registry contract.
type RegistryInitializedIterator struct {
	Event *RegistryInitialized // Event containing the contract specifics and raw log

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
func (it *RegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryInitialized)
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
		it.Event = new(RegistryInitialized)
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
func (it *RegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryInitialized represents a Initialized event raised by the Registry contract.
type RegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Registry *RegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*RegistryInitializedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RegistryInitializedIterator{contract: _Registry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Registry *RegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryInitialized)
				if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Registry *RegistryFilterer) ParseInitialized(log types.Log) (*RegistryInitialized, error) {
	event := new(RegistryInitialized)
	if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Registry contract.
type RegistryPausedIterator struct {
	Event *RegistryPaused // Event containing the contract specifics and raw log

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
func (it *RegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryPaused)
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
		it.Event = new(RegistryPaused)
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
func (it *RegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryPaused represents a Paused event raised by the Registry contract.
type RegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registry *RegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*RegistryPausedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RegistryPausedIterator{contract: _Registry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registry *RegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RegistryPaused) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryPaused)
				if err := _Registry.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registry *RegistryFilterer) ParsePaused(log types.Log) (*RegistryPaused, error) {
	event := new(RegistryPaused)
	if err := _Registry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Registry contract.
type RegistryRoleAdminChangedIterator struct {
	Event *RegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleAdminChanged)
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
		it.Event = new(RegistryRoleAdminChanged)
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
func (it *RegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleAdminChanged represents a RoleAdminChanged event raised by the Registry contract.
type RegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Registry *RegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleAdminChangedIterator{contract: _Registry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Registry *RegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleAdminChanged)
				if err := _Registry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Registry *RegistryFilterer) ParseRoleAdminChanged(log types.Log) (*RegistryRoleAdminChanged, error) {
	event := new(RegistryRoleAdminChanged)
	if err := _Registry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Registry contract.
type RegistryRoleGrantedIterator struct {
	Event *RegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *RegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleGranted)
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
		it.Event = new(RegistryRoleGranted)
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
func (it *RegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleGranted represents a RoleGranted event raised by the Registry contract.
type RegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleGrantedIterator{contract: _Registry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleGranted)
				if err := _Registry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) ParseRoleGranted(log types.Log) (*RegistryRoleGranted, error) {
	event := new(RegistryRoleGranted)
	if err := _Registry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Registry contract.
type RegistryRoleRevokedIterator struct {
	Event *RegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleRevoked)
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
		it.Event = new(RegistryRoleRevoked)
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
func (it *RegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleRevoked represents a RoleRevoked event raised by the Registry contract.
type RegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleRevokedIterator{contract: _Registry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleRevoked)
				if err := _Registry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) ParseRoleRevoked(log types.Log) (*RegistryRoleRevoked, error) {
	event := new(RegistryRoleRevoked)
	if err := _Registry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Registry contract.
type RegistryUnpausedIterator struct {
	Event *RegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *RegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryUnpaused)
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
		it.Event = new(RegistryUnpaused)
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
func (it *RegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryUnpaused represents a Unpaused event raised by the Registry contract.
type RegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registry *RegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RegistryUnpausedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RegistryUnpausedIterator{contract: _Registry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registry *RegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryUnpaused)
				if err := _Registry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registry *RegistryFilterer) ParseUnpaused(log types.Log) (*RegistryUnpaused, error) {
	event := new(RegistryUnpaused)
	if err := _Registry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Registry contract.
type RegistryUpgradedIterator struct {
	Event *RegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *RegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryUpgraded)
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
		it.Event = new(RegistryUpgraded)
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
func (it *RegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryUpgraded represents a Upgraded event raised by the Registry contract.
type RegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Registry *RegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RegistryUpgradedIterator{contract: _Registry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Registry *RegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryUpgraded)
				if err := _Registry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Registry *RegistryFilterer) ParseUpgraded(log types.Log) (*RegistryUpgraded, error) {
	event := new(RegistryUpgraded)
	if err := _Registry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryZRC20TokenRegisteredIterator is returned from FilterZRC20TokenRegistered and is used to iterate over the raw logs and unpacked data for ZRC20TokenRegistered events raised by the Registry contract.
type RegistryZRC20TokenRegisteredIterator struct {
	Event *RegistryZRC20TokenRegistered // Event containing the contract specifics and raw log

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
func (it *RegistryZRC20TokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryZRC20TokenRegistered)
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
		it.Event = new(RegistryZRC20TokenRegistered)
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
func (it *RegistryZRC20TokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryZRC20TokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryZRC20TokenRegistered represents a ZRC20TokenRegistered event raised by the Registry contract.
type RegistryZRC20TokenRegistered struct {
	OriginAddress common.Hash
	Address       common.Address
	Decimals      uint8
	OriginChainId *big.Int
	Symbol        string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterZRC20TokenRegistered is a free log retrieval operation binding the contract event 0xa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367.
//
// Solidity: event ZRC20TokenRegistered(bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol)
func (_Registry *RegistryFilterer) FilterZRC20TokenRegistered(opts *bind.FilterOpts, originAddress [][]byte, address_ []common.Address) (*RegistryZRC20TokenRegisteredIterator, error) {

	var originAddressRule []interface{}
	for _, originAddressItem := range originAddress {
		originAddressRule = append(originAddressRule, originAddressItem)
	}
	var address_Rule []interface{}
	for _, address_Item := range address_ {
		address_Rule = append(address_Rule, address_Item)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ZRC20TokenRegistered", originAddressRule, address_Rule)
	if err != nil {
		return nil, err
	}
	return &RegistryZRC20TokenRegisteredIterator{contract: _Registry.contract, event: "ZRC20TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchZRC20TokenRegistered is a free log subscription operation binding the contract event 0xa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367.
//
// Solidity: event ZRC20TokenRegistered(bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol)
func (_Registry *RegistryFilterer) WatchZRC20TokenRegistered(opts *bind.WatchOpts, sink chan<- *RegistryZRC20TokenRegistered, originAddress [][]byte, address_ []common.Address) (event.Subscription, error) {

	var originAddressRule []interface{}
	for _, originAddressItem := range originAddress {
		originAddressRule = append(originAddressRule, originAddressItem)
	}
	var address_Rule []interface{}
	for _, address_Item := range address_ {
		address_Rule = append(address_Rule, address_Item)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ZRC20TokenRegistered", originAddressRule, address_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryZRC20TokenRegistered)
				if err := _Registry.contract.UnpackLog(event, "ZRC20TokenRegistered", log); err != nil {
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

// ParseZRC20TokenRegistered is a log parse operation binding the contract event 0xa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367.
//
// Solidity: event ZRC20TokenRegistered(bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol)
func (_Registry *RegistryFilterer) ParseZRC20TokenRegistered(log types.Log) (*RegistryZRC20TokenRegistered, error) {
	event := new(RegistryZRC20TokenRegistered)
	if err := _Registry.contract.UnpackLog(event, "ZRC20TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryZRC20TokenUpdatedIterator is returned from FilterZRC20TokenUpdated and is used to iterate over the raw logs and unpacked data for ZRC20TokenUpdated events raised by the Registry contract.
type RegistryZRC20TokenUpdatedIterator struct {
	Event *RegistryZRC20TokenUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryZRC20TokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryZRC20TokenUpdated)
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
		it.Event = new(RegistryZRC20TokenUpdated)
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
func (it *RegistryZRC20TokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryZRC20TokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryZRC20TokenUpdated represents a ZRC20TokenUpdated event raised by the Registry contract.
type RegistryZRC20TokenUpdated struct {
	Address common.Address
	Active  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterZRC20TokenUpdated is a free log retrieval operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_Registry *RegistryFilterer) FilterZRC20TokenUpdated(opts *bind.FilterOpts) (*RegistryZRC20TokenUpdatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ZRC20TokenUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistryZRC20TokenUpdatedIterator{contract: _Registry.contract, event: "ZRC20TokenUpdated", logs: logs, sub: sub}, nil
}

// WatchZRC20TokenUpdated is a free log subscription operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_Registry *RegistryFilterer) WatchZRC20TokenUpdated(opts *bind.WatchOpts, sink chan<- *RegistryZRC20TokenUpdated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ZRC20TokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryZRC20TokenUpdated)
				if err := _Registry.contract.UnpackLog(event, "ZRC20TokenUpdated", log); err != nil {
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

// ParseZRC20TokenUpdated is a log parse operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_Registry *RegistryFilterer) ParseZRC20TokenUpdated(log types.Log) (*RegistryZRC20TokenUpdated, error) {
	event := new(RegistryZRC20TokenUpdated)
	if err := _Registry.contract.UnpackLog(event, "ZRC20TokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
