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
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GATEWAY_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY_MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bootstrapChains\",\"inputs\":[{\"name\":\"chains\",\"type\":\"tuple[]\",\"internalType\":\"structChainInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasZRC20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"metadataEntries\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistry.ChainMetadataEntry[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bootstrapContracts\",\"inputs\":[{\"name\":\"contracts\",\"type\":\"tuple[]\",\"internalType\":\"structContractInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"configEntries\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistry.ContractConfigEntry[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bootstrapZRC20Tokens\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structZRC20Info[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeChainStatus\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasZRC20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"activation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"coreRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gatewayEVM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllChains\",\"inputs\":[],\"outputs\":[{\"name\":\"chainsInfo\",\"type\":\"tuple[]\",\"internalType\":\"structChainInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasZRC20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"contractsInfo\",\"type\":\"tuple[]\",\"internalType\":\"structContractInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllZRC20Tokens\",\"inputs\":[],\"outputs\":[{\"name\":\"tokensInfo\",\"type\":\"tuple[]\",\"internalType\":\"structZRC20Info[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainMetadata\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractConfiguration\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractInfo\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getZRC20AddressByForeignAsset\",\"inputs\":[{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getZRC20TokenInfo\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registryManager_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gatewayEVM_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"coreRegistry_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerContract\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerZRC20Token\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setZRC20TokenActive\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateChainMetadata\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateContractConfiguration\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"ChainMetadataUpdated\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainStatusChanged\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newStatus\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractConfigurationUpdated\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractStatusChanged\",\"inputs\":[{\"name\":\"addressBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenRegistered\",\"inputs\":[{\"name\":\"originAddress\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"address_\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenUpdated\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ChainActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainNonActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ContractAlreadyRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ContractNotFound\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContractType\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZRC20AlreadyRegistered\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZRC20SymbolAlreadyInUse\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516149f66100fd600039600081816127b0015281816127d9015261298c01526149f66000f3fe6080604052600436106102c65760003560e01c80639060bda911610179578063bd8fde1c116100d6578063dc07f8661161008a578063f354b31f11610064578063f354b31f146108b4578063f851a440146108d4578063f8c8765e146108f457600080fd5b8063dc07f8661461082e578063e63ab1e91461084e578063e9d6c5ba1461088257600080fd5b8063c50d6dd9116100bb578063c50d6dd9146107ba578063d3523ea2146107ee578063d547741f1461080e57600080fd5b8063bd8fde1c14610764578063c1bd469f1461079857600080fd5b8063a217fddf1161012d578063a8f2cb9611610112578063a8f2cb96146106db578063aa808c06146106fb578063ad3cb1cc1461071b57600080fd5b8063a217fddf146106a6578063a4b7e5f9146106bb57600080fd5b806391d148541161015e57806391d14854146105fd57806394cc8683146106625780639ca220dd1461068457600080fd5b80639060bda9146105bd578063909b6a03146105dd57600080fd5b80634f1ef28611610227578063676cc054116101db5780636e9e2d3f116101c05780636e9e2d3f146105685780637066b18d146105885780638456cb59146105a857600080fd5b8063676cc0541461051b5780636bf3d05a1461054857600080fd5b80635c975abb1161020c5780635c975abb146104965780635cf92c9f146104cd578063631d62e4146104fb57600080fd5b80634f1ef2861461046e57806352d1902d1461048157600080fd5b8063248a9ca31161027e5780632f2ff15d116102635780632f2ff15d1461041957806336568abe146104395780633f4ba83a1461045957600080fd5b8063248a9ca31461039c5780632c78f74c146103f957600080fd5b806310d29b9e116102af57806310d29b9e1461033857806318d3ce961461035a5780632259e9e51461037c57600080fd5b806301ffc9a7146102cb5780630c63109e14610300575b600080fd5b3480156102d757600080fd5b506102eb6102e63660046138ec565b610914565b60405190151581526020015b60405180910390f35b34801561030c57600080fd5b50600154610320906001600160a01b031681565b6040516001600160a01b0390911681526020016102f7565b34801561034457600080fd5b5061035861035336600461398c565b6109ad565b005b34801561036657600080fd5b5061036f610a62565b6040516102f79190613a3a565b34801561038857600080fd5b50610358610397366004613afb565b610d04565b3480156103a857600080fd5b506103eb6103b7366004613b7a565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016102f7565b34801561040557600080fd5b50600b54610320906001600160a01b031681565b34801561042557600080fd5b50610358610434366004613baa565b610d97565b34801561044557600080fd5b50610358610454366004613baa565b610de1565b34801561046557600080fd5b50610358610e32565b61035861047c366004613c05565b610e48565b34801561048d57600080fd5b506103eb610e67565b3480156104a257600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166102eb565b3480156104d957600080fd5b506104ed6104e8366004613cd0565b610e96565b6040516102f7929190613d1c565b34801561050757600080fd5b50610358610516366004613afb565b610f91565b34801561052757600080fd5b5061053b610536366004613d3f565b611038565b6040516102f79190613d81565b34801561055457600080fd5b50600c54610320906001600160a01b031681565b34801561057457600080fd5b50610358610583366004613da5565b611140565b34801561059457600080fd5b5061053b6105a3366004613cd0565b6111fc565b3480156105b457600080fd5b506103586112c5565b3480156105c957600080fd5b506103586105d8366004613e75565b6112f7565b3480156105e957600080fd5b506103586105f8366004613ee4565b611389565b34801561060957600080fd5b506102eb610618366004613baa565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561066e57600080fd5b5061067761144a565b6040516102f79190613f26565b34801561069057600080fd5b506106996114a2565b6040516102f79190613f69565b3480156106b257600080fd5b506103eb600081565b3480156106c757600080fd5b506103586106d6366004614016565b611661565b3480156106e757600080fd5b506103586106f6366004614087565b61176a565b34801561070757600080fd5b50610320610716366004613cd0565b6117ec565b34801561072757600080fd5b5061053b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b34801561077057600080fd5b506103eb7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa381565b3480156107a457600080fd5b506107ad61182c565b6040516102f791906140f6565b3480156107c657600080fd5b506103eb7fb90e9995c6170fff8ea03e9ad6919878e483770c237f1a6f330ceaa7112b344a81565b3480156107fa57600080fd5b5061053b610809366004613afb565b611b33565b34801561081a57600080fd5b50610358610829366004613baa565b611c1b565b34801561083a57600080fd5b50610358610849366004614016565b611c5f565b34801561085a57600080fd5b506103eb7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b34801561088e57600080fd5b506108a261089d3660046141ed565b611d63565b6040516102f796959493929190614208565b3480156108c057600080fd5b506103586108cf366004614266565b611fb5565b3480156108e057600080fd5b50600054610320906001600160a01b031681565b34801561090057600080fd5b5061035861090f366004614316565b612050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806109a757507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b3330146109e6576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109f284848484612374565b7f6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c600760008681526020019081526020016000208484604051610a3692919061435f565b9081526020016040518091039020600101604051610a5491906143c2565b60405180910390a150505050565b6004546060908067ffffffffffffffff811115610a8157610a81613bd6565b604051908082528060200260200182016040528015610adf57816020015b610acc60405180608001604052806000151581526020016060815260200160608152602001600081525090565b815260200190600190039081610a9f5790505b50915060005b81811015610cff57600060048281548110610b0257610b0261444e565b906000526020600020906002020160405180604001604052908160008201548152602001600182018054610b359061436f565b80601f0160208091040260200160405190810160405280929190818152602001828054610b619061436f565b8015610bae5780601f10610b8357610100808354040283529160200191610bae565b820191906000526020600020905b815481529060010190602001808311610b9157829003601f168201915b505050505081525050905060008160000151905060008260200151905060405180608001604052806007600085815260200190815260200160002083604051610bf7919061447d565b90815260408051602092819003830190205460ff161515835260008681526007835281902090519290910191610c2e90859061447d565b90815260200160405180910390206001018054610c4a9061436f565b80601f0160208091040260200160405190810160405280929190818152602001828054610c769061436f565b8015610cc35780601f10610c9857610100808354040283529160200191610cc3565b820191906000526020600020905b815481529060010190602001808311610ca657829003601f168201915b5050505050815260200182815260200183815250868581518110610ce957610ce961444e565b6020908102919091010152505050600101610ae5565b505090565b333014610d3d576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d456124c8565b610d528585858585612526565b847f40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee63485858585604051610d8894939291906144c4565b60405180910390a25050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610dd1816125b4565b610ddb83836125be565b50505050565b6001600160a01b0381163314610e23576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e2d828261268d565b505050565b6000610e3d816125b4565b610e45612733565b50565b610e506127a5565b610e5982612875565b610e638282612880565b5050565b6000610e71612981565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b600083815260076020526040808220905160609190610eb8908690869061435f565b908152604080519182900360209081018320546000898152600790925291902060ff909116935090610eed908690869061435f565b90815260200160405180910390206001018054610f099061436f565b80601f0160208091040260200160405190810160405280929190818152602001828054610f359061436f565b8015610f825780601f10610f5757610100808354040283529160200191610f82565b820191906000526020600020905b815481529060010190602001808311610f6557829003601f168201915b50505050509050935093915050565b333014610fca576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fd26124c8565b610fdf85858585856129e3565b8383604051610fef92919061435f565b6040518091039020857f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce258184846040516110299291906144f6565b60405180910390a35050505050565b60607fb90e9995c6170fff8ea03e9ad6919878e483770c237f1a6f330ceaa7112b344a611064816125b4565b61106c6124c8565b600c546001600160a01b031661108560208701876141ed565b6001600160a01b0316146110c5576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080306001600160a01b031686866040516110e292919061435f565b6000604051808303816000865af19150503d806000811461111f576040519150601f19603f3d011682016040523d82523d6000602084013e611124565b606091505b50915091508161113657805160208201fd5b9695505050505050565b333014611179576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111816124c8565b611192898989898989898989612c7e565b886001600160a01b031685856040516111ac92919061435f565b60405180910390207fa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e36783898c8c6040516111e9949392919061450a565b60405180910390a3505050505050505050565b606060066000858152602001908152602001600020600401838360405161122492919061435f565b9081526020016040518091039020805461123d9061436f565b80601f01602080910402602001604051908101604052809291908181526020018280546112699061436f565b80156112b65780601f1061128b576101008083540402835291602001916112b6565b820191906000526020600020905b81548152906001019060200180831161129957829003601f168201915b505050505090505b9392505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6112ef816125b4565b610e45612fb4565b333014611330576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113386124c8565b611342828261300f565b604080516001600160a01b038416815282151560208201527f9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8910160405180910390a15050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36113b3816125b4565b6113bb6124c8565b60005b82811015610ddb57368484838181106113d9576113d961444e565b90506020028101906113eb919061452d565b905061144161140060408301602084016141ed565b61140d6080840184614561565b606085013561141f6040870187614561565b61142c60a0890189614561565b61143c60e08b0160c08c016145c6565b612c7e565b506001016113be565b6060600280548060200260200160405190810160405280929190818152602001828054801561149857602002820191906000526020600020905b815481526020019060010190808311611484575b5050505050905090565b6003546060908067ffffffffffffffff8111156114c1576114c1613bd6565b60405190808252806020026020018201604052801561153057816020015b60408051608081018252600080825260208083018290529282015260608082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092019101816114df5790505b50915060005b81811015610cff576000600382815481106115535761155361444e565b6000918252602080832090910154604080516080810182528285526006808552828620805460ff161515835282860185905260028101546001600160a01b03169383019390935294839052939092526003909101805491935060608301916115ba9061436f565b80601f01602080910402602001604051908101604052809291908181526020018280546115e69061436f565b80156116335780601f1061160857610100808354040283529160200191611633565b820191906000526020600020905b81548152906001019060200180831161161657829003601f168201915b505050505081525084838151811061164d5761164d61444e565b602090810291909101015250600101611536565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361168b816125b4565b6116936124c8565b60005b848110156116f557368686838181106116b1576116b161444e565b90506020028101906116c391906145e1565b90506116ec60608201356116da6040840184614561565b6116e76020860186614561565b6129e3565b50600101611696565b5060005b8281101561176257368484838181106117145761171461444e565b905060200281019061172691906145e1565b9050611759813561173a6020840184614561565b6117476040860186614561565b6117546060880188614561565b6130ff565b506001016116f9565b505050505050565b3330146117a3576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117ab6124c8565b6117b88585858585613261565b847fc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e82604051610d88911515815260200190565b6000838152600a6020526040808220905161180a908590859061435f565b908152604051908190036020019020546001600160a01b031690509392505050565b6005546060908067ffffffffffffffff81111561184b5761184b613bd6565b6040519080825280602002602001820160405280156118ca57816020015b6118b76040518060e0016040528060001515815260200160006001600160a01b0316815260200160608152602001600081526020016060815260200160608152602001600060ff1681525090565b8152602001906001900390816118695790505b50915060005b81811015610cff576000600582815481106118ed576118ed61444e565b60009182526020808320909101546001600160a01b0390811680845260088352604093849020845160e081018652815460ff81161515825261010090049093169383019390935260018301805491955091938401919061194c9061436f565b80601f01602080910402602001604051908101604052809291908181526020018280546119789061436f565b80156119c55780601f1061199a576101008083540402835291602001916119c5565b820191906000526020600020905b8154815290600101906020018083116119a857829003601f168201915b50505050508152602001600282015481526020016003820180546119e89061436f565b80601f0160208091040260200160405190810160405280929190818152602001828054611a149061436f565b8015611a615780601f10611a3657610100808354040283529160200191611a61565b820191906000526020600020905b815481529060010190602001808311611a4457829003601f168201915b50505050508152602001600482018054611a7a9061436f565b80601f0160208091040260200160405190810160405280929190818152602001828054611aa69061436f565b8015611af35780601f10611ac857610100808354040283529160200191611af3565b820191906000526020600020905b815481529060010190602001808311611ad657829003601f168201915b50505091835250506005919091015460ff166020909101528451859084908110611b1f57611b1f61444e565b6020908102919091010152506001016118d0565b6060600760008781526020019081526020016000208585604051611b5892919061435f565b90815260200160405180910390206003018383604051611b7992919061435f565b90815260200160405180910390208054611b929061436f565b80601f0160208091040260200160405190810160405280929190818152602001828054611bbe9061436f565b8015611c0b5780601f10611be057610100808354040283529160200191611c0b565b820191906000526020600020905b815481529060010190602001808311611bee57829003601f168201915b5050505050905095945050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611c55816125b4565b610ddb838361268d565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3611c89816125b4565b611c916124c8565b60005b84811015611d035736868683818110611caf57611caf61444e565b9050602002810190611cc191906145e1565b9050611cfa6020820135611cdb60608401604085016141ed565b611ce86060850185614561565b611cf56020870187614615565b613261565b50600101611c94565b5060005b828110156117625736848483818110611d2257611d2261444e565b9050602002810190611d349190614630565b9050611d5a8135611d486020840184614561565b611d556040860186614561565b612526565b50600101611d07565b6001600160a01b038082166000908152600860209081526040808320815160e081018352815460ff81161515825261010090049095169285019290925260018201805493946060948694869485948794859490939284019190611dc59061436f565b80601f0160208091040260200160405190810160405280929190818152602001828054611df19061436f565b8015611e3e5780601f10611e1357610100808354040283529160200191611e3e565b820191906000526020600020905b815481529060010190602001808311611e2157829003601f168201915b5050505050815260200160028201548152602001600382018054611e619061436f565b80601f0160208091040260200160405190810160405280929190818152602001828054611e8d9061436f565b8015611eda5780601f10611eaf57610100808354040283529160200191611eda565b820191906000526020600020905b815481529060010190602001808311611ebd57829003601f168201915b50505050508152602001600482018054611ef39061436f565b80601f0160208091040260200160405190810160405280929190818152602001828054611f1f9061436f565b8015611f6c5780601f10611f4157610100808354040283529160200191611f6c565b820191906000526020600020905b815481529060010190602001808311611f4f57829003601f168201915b50505091835250506005919091015460ff16602090910152805160808201516060830151604084015160a085015160c090950151939d929c50909a509850919650945092505050565b333014611fee576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ff66124c8565b612005878787878787876130ff565b867faea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c587878787878760405161203f96959493929190614664565b60405180910390a250505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff1660008115801561209b5750825b905060008267ffffffffffffffff1660011480156120b85750303b155b9050811580156120c6575080155b156120fd576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000166001178555831561215e5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038916158061217b57506001600160a01b038716155b8061218d57506001600160a01b038616155b8061219f57506001600160a01b038816155b156121d6576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6121de613451565b6121e6613451565b6121ee613459565b6121f960008a6125be565b506122247ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3896125be565b5061224f7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a896125be565b5061227a7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8a6125be565b506122a57fb90e9995c6170fff8ea03e9ad6919878e483770c237f1a6f330ceaa7112b344a886125be565b50600080546001600160a01b03808c167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255600180548b8416908316179055600b80548a8416908316179055600c80549289169290911691909117905583156123695784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b60008481526006602052604090205460ff166123c4576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018590526024015b60405180910390fd5b60008290036124025782826040517ec10cfd0000000000000000000000000000000000000000000000000000000081526004016123bb9291906144f6565b600084815260076020526040908190209051612421908590859061435f565b9081526020016040518091039020600101805461243d9061436f565b905060000361247e578383836040517f2b4f9c860000000000000000000000000000000000000000000000000000000081526004016123bb939291906146ad565b806007600086815260200190815260200160002084846040516124a292919061435f565b908152604051908190036020019020805491151560ff1990921691909117905550505050565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615612524576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60008581526006602052604090205460ff16612571576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018690526024016123bb565b818160066000888152602001908152602001600020600401868660405161259992919061435f565b9081526020016040518091039020918261176292919061470e565b610e45813361348c565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612683576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556126393390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506109a7565b60009150506109a7565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612683576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506109a7565b61273b613519565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061283e57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166128327f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15612524576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610e63816125b4565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156128da575060408051601f3d908101601f191682019092526128d79181019061480a565b60015b61291b576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016123bb565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612977576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016123bb565b610e2d8383613574565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612524576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008581526006602052604090205460ff16612a2e576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018690526024016123bb565b6000839003612a6c5783836040517ec10cfd0000000000000000000000000000000000000000000000000000000081526004016123bb9291906144f6565b6000819003612aa7576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000858152600760205260408082209051612ac5908790879061435f565b90815260200160405180910390206001018054612ae19061436f565b90501115612b255784848484846040517f2b192eab0000000000000000000000000000000000000000000000000000000081526004016123bb959493929190614823565b6001600760008781526020019081526020016000208585604051612b4a92919061435f565b9081526040805160209281900383018120805460ff1916941515949094179093556000888152600790925290208391839190612b89908890889061435f565b90815260200160405180910390206001019182612ba792919061470e565b508383600760008881526020019081526020016000208686604051612bcd92919061435f565b90815260200160405180910390206002019182612beb92919061470e565b506004604051806040016040528087815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939094525050835460018181018655948252602091829020845160029092020190815590830151929390929083019150612c6b908261485c565b5050508383604051610fef92919061435f565b6001600160a01b038916612cbe576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000879003612d28576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f53796d626f6c2063616e6e6f7420626520656d7074790000000000000000000060448201526064016123bb565b6001600160a01b0389811660009081526008602052604090205461010090041615612d8a576040517f63f4ee1f0000000000000000000000000000000000000000000000000000000081526001600160a01b038a1660048201526024016123bb565b60006001600160a01b031660098989604051612da792919061435f565b908152604051908190036020019020546001600160a01b031614612dfb5787876040517fb295cac90000000000000000000000000000000000000000000000000000000081526004016123bb9291906144f6565b6001600160a01b0389166000818152600860205260409020805460017fffffffffffffffffffffff000000000000000000000000000000000000000000909116610100909302929092178217815501612e5585878361470e565b506001600160a01b038916600090815260086020526040902060028101879055600301612e83888a8361470e565b506001600160a01b038916600090815260086020526040902060058101805460ff191660ff8416179055600401612ebb83858361470e565b5088600a60008881526020019081526020016000208686604051612ee092919061435f565b908152602001604051809103902060006101000a8154816001600160a01b0302191690836001600160a01b031602179055508860098989604051612f2592919061435f565b90815260405190819003602001902080546001600160a01b039283167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617909155600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180549b9092169a16999099179098555050505050505050565b612fbc6124c8565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612787565b6001600160a01b03821661304f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038281166000908152600860205260409020546101009004166130d4576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5a52433230206e6f74207265676973746572656400000000000000000000000060448201526064016123bb565b6001600160a01b03919091166000908152600860205260409020805460ff1916911515919091179055565b60008781526006602052604090205460ff1661314a576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018890526024016123bb565b60008590036131885785856040517ec10cfd0000000000000000000000000000000000000000000000000000000081526004016123bb9291906144f6565b6000878152600760205260409081902090516131a7908890889061435f565b9081526040519081900360200190205460ff166131f6578686866040517f2b4f9c860000000000000000000000000000000000000000000000000000000081526004016123bb939291906146ad565b8181600760008a8152602001908152602001600020888860405161321b92919061435f565b9081526020016040518091039020600301868660405161323c92919061435f565b9081526020016040518091039020918261325792919061470e565b5050505050505050565b600082900361329c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008581526006602052604090205460ff1680156132b75750805b156132f1576040517fa1452cb0000000000000000000000000000000000000000000000000000000008152600481018690526024016123bb565b60008581526006602052604090205460ff1615801561330e575080155b15613348576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018690526024016123bb565b600085815260066020526040902060030180546133649061436f565b90506000036133a357600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018590555b6000858152600660205260409020805460ff19168215151781556002810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03871617905560030161340083858361470e565b50801561344157600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0185905561344a565b61344a856135ca565b5050505050565b612524613678565b613461613678565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610e63576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044016123bb565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16612524576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61357d826136df565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156135c257610e2d8282613787565b610e636137fd565b60005b600254811015610e635781600282815481106135eb576135eb61444e565b906000526020600020015403613670576002805461360b90600190614957565b8154811061361b5761361b61444e565b9060005260206000200154600282815481106136395761363961444e565b600091825260209091200155600280548061365657613656614991565b600190038181906000526020600020016000905590555050565b6001016135cd565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16612524576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001600160a01b03163b60000361372e576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016123bb565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516137a4919061447d565b600060405180830381855af49150503d80600081146137df576040519150601f19603f3d011682016040523d82523d6000602084013e6137e4565b606091505b50915091506137f4858383613835565b95945050505050565b3415612524576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608261384a57613845826138aa565b6112be565b815115801561386157506001600160a01b0384163b155b156138a3576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016123bb565b50806112be565b8051156138ba5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156138fe57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146112be57600080fd5b60008083601f84011261394057600080fd5b50813567ffffffffffffffff81111561395857600080fd5b60208301915083602082850101111561397057600080fd5b9250929050565b8035801515811461398757600080fd5b919050565b600080600080606085870312156139a257600080fd5b84359350602085013567ffffffffffffffff8111156139c057600080fd5b6139cc8782880161392e565b90945092506139df905060408601613977565b905092959194509250565b60005b83811015613a055781810151838201526020016139ed565b50506000910152565b60008151808452613a268160208601602086016139ea565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613aef577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160806020880152613ab26080880182613a0e565b905060408201518782036040890152613acb8282613a0e565b60609384015198909301979097525094506020938401939190910190600101613a62565b50929695505050505050565b600080600080600060608688031215613b1357600080fd5b85359450602086013567ffffffffffffffff811115613b3157600080fd5b613b3d8882890161392e565b909550935050604086013567ffffffffffffffff811115613b5d57600080fd5b613b698882890161392e565b969995985093965092949392505050565b600060208284031215613b8c57600080fd5b5035919050565b80356001600160a01b038116811461398757600080fd5b60008060408385031215613bbd57600080fd5b82359150613bcd60208401613b93565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060408385031215613c1857600080fd5b613c2183613b93565b9150602083013567ffffffffffffffff811115613c3d57600080fd5b8301601f81018513613c4e57600080fd5b803567ffffffffffffffff811115613c6857613c68613bd6565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff82111715613c9857613c98613bd6565b604052818152828201602001871015613cb057600080fd5b816020840160208301376000602083830101528093505050509250929050565b600080600060408486031215613ce557600080fd5b83359250602084013567ffffffffffffffff811115613d0357600080fd5b613d0f8682870161392e565b9497909650939450505050565b8215158152604060208201526000613d376040830184613a0e565b949350505050565b60008060008385036040811215613d5557600080fd5b6020811215613d6357600080fd5b50839250602084013567ffffffffffffffff811115613d0357600080fd5b6020815260006112be6020830184613a0e565b803560ff8116811461398757600080fd5b600080600080600080600080600060c08a8c031215613dc357600080fd5b613dcc8a613b93565b985060208a013567ffffffffffffffff811115613de857600080fd5b613df48c828d0161392e565b90995097505060408a0135955060608a013567ffffffffffffffff811115613e1b57600080fd5b613e278c828d0161392e565b90965094505060808a013567ffffffffffffffff811115613e4757600080fd5b613e538c828d0161392e565b9094509250613e66905060a08b01613d94565b90509295985092959850929598565b60008060408385031215613e8857600080fd5b613e9183613b93565b9150613bcd60208401613977565b60008083601f840112613eb157600080fd5b50813567ffffffffffffffff811115613ec957600080fd5b6020830191508360208260051b850101111561397057600080fd5b60008060208385031215613ef757600080fd5b823567ffffffffffffffff811115613f0e57600080fd5b613f1a85828601613e9f565b90969095509350505050565b602080825282518282018190526000918401906040840190835b81811015613f5e578351835260209384019390920191600101613f40565b509095945050505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613aef577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160208701526001600160a01b03604082015116604087015260608101519050608060608701526140006080870182613a0e565b9550506020938401939190910190600101613f91565b6000806000806040858703121561402c57600080fd5b843567ffffffffffffffff81111561404357600080fd5b61404f87828801613e9f565b909550935050602085013567ffffffffffffffff81111561406f57600080fd5b61407b87828801613e9f565b95989497509550505050565b60008060008060006080868803121561409f57600080fd5b853594506140af60208701613b93565b9350604086013567ffffffffffffffff8111156140cb57600080fd5b6140d78882890161392e565b90945092506140ea905060608701613977565b90509295509295909350565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613aef577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051151586526001600160a01b036020820151166020870152604081015160e0604088015261418160e0880182613a0e565b905060608201516060880152608082015187820360808901526141a48282613a0e565b91505060a082015187820360a08901526141be8282613a0e565b91505060c082015191506141d760c088018360ff169052565b955050602093840193919091019060010161411e565b6000602082840312156141ff57600080fd5b6112be82613b93565b861515815260c06020820152600061422360c0830188613a0e565b866040840152828103606084015261423b8187613a0e565b9050828103608084015261424f8186613a0e565b91505060ff831660a0830152979650505050505050565b60008060008060008060006080888a03121561428157600080fd5b87359650602088013567ffffffffffffffff81111561429f57600080fd5b6142ab8a828b0161392e565b909750955050604088013567ffffffffffffffff8111156142cb57600080fd5b6142d78a828b0161392e565b909550935050606088013567ffffffffffffffff8111156142f757600080fd5b6143038a828b0161392e565b989b979a50959850939692959293505050565b6000806000806080858703121561432c57600080fd5b61433585613b93565b935061434360208601613b93565b925061435160408601613b93565b91506139df60608601613b93565b8183823760009101908152919050565b600181811c9082168061438357607f821691505b6020821081036143bc577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b6020815260008083546143d48161436f565b80602086015260018216600081146143f3576001811461440f57614443565b60ff1983166040870152604082151560051b8701019350614443565b86600052602060002060005b8381101561443a5781548882016040015260019091019060200161441b565b87016040019450505b509195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000825161448f8184602087016139ea565b9190910192915050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b6040815260006144d8604083018688614499565b82810360208401526144eb818587614499565b979650505050505050565b602081526000613d37602083018486614499565b60ff85168152836020820152606060408201526000611136606083018486614499565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2183360301811261448f57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261459657600080fd5b83018035915067ffffffffffffffff8211156145b157600080fd5b60200191503681900382131561397057600080fd5b6000602082840312156145d857600080fd5b6112be82613d94565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8183360301811261448f57600080fd5b60006020828403121561462757600080fd5b6112be82613977565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa183360301811261448f57600080fd5b60608152600061467860608301888a614499565b828103602084015261468b818789614499565b905082810360408401526146a0818587614499565b9998505050505050505050565b8381526040602082015260006137f4604083018486614499565b601f821115610e2d57806000526020600020601f840160051c810160208510156146ee5750805b601f840160051c820191505b8181101561344a57600081556001016146fa565b67ffffffffffffffff83111561472657614726613bd6565b61473a83614734835461436f565b836146c7565b6000601f84116001811461478c57600085156147565750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b17835561344a565b600083815260209020601f19861690835b828110156147bd578685013582556020948501946001909201910161479d565b50868210156147f8577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b60006020828403121561481c57600080fd5b5051919050565b85815260606020820152600061483d606083018688614499565b8281036040840152614850818587614499565b98975050505050505050565b815167ffffffffffffffff81111561487657614876613bd6565b61488a81614884845461436f565b846146c7565b6020601f8211600181146148dc57600083156148a65750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b17845561344a565b600084815260208120601f198516915b8281101561490c57878501518255602094850194600190920191016148ec565b508482101561494857868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b818103818111156109a7577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220097274093dc633636c332b05c250b7789381af5b68809260018fa3f1b9ab3e6864736f6c634300081a0033",
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

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Registry *RegistryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Registry *RegistrySession) Admin() (common.Address, error) {
	return _Registry.Contract.Admin(&_Registry.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Registry *RegistryCallerSession) Admin() (common.Address, error) {
	return _Registry.Contract.Admin(&_Registry.CallOpts)
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

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_Registry *RegistryCaller) RegistryManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "registryManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_Registry *RegistrySession) RegistryManager() (common.Address, error) {
	return _Registry.Contract.RegistryManager(&_Registry.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_Registry *RegistryCallerSession) RegistryManager() (common.Address, error) {
	return _Registry.Contract.RegistryManager(&_Registry.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address admin_, address registryManager_, address gatewayEVM_, address coreRegistry_) returns()
func (_Registry *RegistryTransactor) Initialize(opts *bind.TransactOpts, admin_ common.Address, registryManager_ common.Address, gatewayEVM_ common.Address, coreRegistry_ common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "initialize", admin_, registryManager_, gatewayEVM_, coreRegistry_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address admin_, address registryManager_, address gatewayEVM_, address coreRegistry_) returns()
func (_Registry *RegistrySession) Initialize(admin_ common.Address, registryManager_ common.Address, gatewayEVM_ common.Address, coreRegistry_ common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, admin_, registryManager_, gatewayEVM_, coreRegistry_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address admin_, address registryManager_, address gatewayEVM_, address coreRegistry_) returns()
func (_Registry *RegistryTransactorSession) Initialize(admin_ common.Address, registryManager_ common.Address, gatewayEVM_ common.Address, coreRegistry_ common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, admin_, registryManager_, gatewayEVM_, coreRegistry_)
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
