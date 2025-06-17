// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coreregistry

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

// CoreRegistryMetaData contains all meta data concerning the CoreRegistry contract.
var CoreRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"CROSS_CHAIN_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY_MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeChainStatus\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasZRC20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"activation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeRegistryManager\",\"inputs\":[{\"name\":\"newRegistryManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gatewayZEVM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayZEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllChains\",\"inputs\":[],\"outputs\":[{\"name\":\"chainsInfo\",\"type\":\"tuple[]\",\"internalType\":\"structChainInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasZRC20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"contractsInfo\",\"type\":\"tuple[]\",\"internalType\":\"structContractInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllZRC20Tokens\",\"inputs\":[],\"outputs\":[{\"name\":\"tokensInfo\",\"type\":\"tuple[]\",\"internalType\":\"structZRC20Info[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainMetadata\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractConfiguration\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractInfo\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getZRC20AddressByForeignAsset\",\"inputs\":[{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getZRC20TokenInfo\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registryManager_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gatewayZEVM_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerContract\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerZRC20Token\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setZRC20TokenActive\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateChainMetadata\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateContractConfiguration\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"oldAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainMetadataUpdated\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainStatusChanged\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newStatus\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractConfigurationUpdated\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractStatusChanged\",\"inputs\":[{\"name\":\"addressBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistryManagerChanged\",\"inputs\":[{\"name\":\"oldRegistryManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRegistryManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenRegistered\",\"inputs\":[{\"name\":\"originAddress\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"address_\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenUpdated\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ChainActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainNonActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ContractAlreadyRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ContractNotFound\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContractType\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZRC20AlreadyRegistered\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZRC20SymbolAlreadyInUse\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614fa86100fd60003960008181612797015281816127c001526129730152614fa86000f3fe6080604052600436106102855760003560e01c80639060bda911610153578063bd8fde1c116100cb578063d547741f1161007f578063e9d6c5ba11610064578063e9d6c5ba146107e4578063f354b31f14610816578063f851a4401461083657600080fd5b8063d547741f14610790578063e63ab1e9146107b057600080fd5b8063c1bd469f116100b0578063c1bd469f1461072e578063cc5ad8b614610750578063d3523ea21461077057600080fd5b8063bd8fde1c146106da578063c0c53b8b1461070e57600080fd5b8063a217fddf11610122578063a8f2cb9611610107578063a8f2cb9614610651578063aa808c0614610671578063ad3cb1cc1461069157600080fd5b8063a217fddf14610625578063a3ebd14c1461063a57600080fd5b80639060bda91461055c57806391d148541461057c57806394cc8683146105e15780639ca220dd1461060357600080fd5b80633f4ba83a11610201578063631d62e4116101b55780637066b18d1161019a5780637066b18d146104fa5780638456cb59146105275780638f2839701461053c57600080fd5b8063631d62e4146104ba5780636e9e2d3f146104da57600080fd5b806352d1902d116101e657806352d1902d146104405780635c975abb146104555780635cf92c9f1461048c57600080fd5b80633f4ba83a146104185780634f1ef2861461042d57600080fd5b80632259e9e5116102585780632f2ff15d1161023d5780632f2ff15d146103b85780633500c24b146103d857806336568abe146103f857600080fd5b80632259e9e51461033b578063248a9ca31461035b57600080fd5b806301ffc9a71461028a5780630c63109e146102bf57806310d29b9e146102f757806318d3ce9614610319575b600080fd5b34801561029657600080fd5b506102aa6102a5366004613ed5565b610856565b60405190151581526020015b60405180910390f35b3480156102cb57600080fd5b506001546102df906001600160a01b031681565b6040516001600160a01b0390911681526020016102b6565b34801561030357600080fd5b50610317610312366004613f6e565b6108ef565b005b34801561032557600080fd5b5061032e6109aa565b6040516102b6919061401d565b34801561034757600080fd5b506103176103563660046140de565b610c4c565b34801561036757600080fd5b506103aa61037636600461415d565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016102b6565b3480156103c457600080fd5b506103176103d336600461418b565b610cde565b3480156103e457600080fd5b506103176103f33660046141bb565b610d28565b34801561040457600080fd5b5061031761041336600461418b565b610ebb565b34801561042457600080fd5b50610317610f0c565b61031761043b366004614207565b610f22565b34801561044c57600080fd5b506103aa610f41565b34801561046157600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166102aa565b34801561049857600080fd5b506104ac6104a73660046142d4565b610f70565b6040516102b6929190614320565b3480156104c657600080fd5b506103176104d53660046140de565b61106b565b3480156104e657600080fd5b506103176104f5366004614343565b611111565b34801561050657600080fd5b5061051a6105153660046142d4565b6111d0565b6040516102b6919061441e565b34801561053357600080fd5b50610317611299565b34801561054857600080fd5b506103176105573660046141bb565b6112cb565b34801561056857600080fd5b50610317610577366004614431565b61141f565b34801561058857600080fd5b506102aa61059736600461418b565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156105ed57600080fd5b506105f66114ad565b6040516102b6919061445f565b34801561060f57600080fd5b50610618611505565b6040516102b691906144a2565b34801561063157600080fd5b506103aa600081565b34801561064657600080fd5b506103aa6207a12081565b34801561065d57600080fd5b5061031761066c36600461454f565b6116c4565b34801561067d57600080fd5b506102df61068c3660046142d4565b611744565b34801561069d57600080fd5b5061051a6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b3480156106e657600080fd5b506103aa7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa381565b34801561071a57600080fd5b506103176107293660046145c3565b611784565b34801561073a57600080fd5b50610743611b31565b6040516102b6919061460e565b34801561075c57600080fd5b50600b546102df906001600160a01b031681565b34801561077c57600080fd5b5061051a61078b3660046140de565b611e38565b34801561079c57600080fd5b506103176107ab36600461418b565b611f20565b3480156107bc57600080fd5b506103aa7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b3480156107f057600080fd5b506108046107ff3660046141bb565b611f64565b6040516102b696959493929190614705565b34801561082257600080fd5b50610317610831366004614763565b6121b6565b34801561084257600080fd5b506000546102df906001600160a01b031681565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806108e957507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361091981612252565b61092161225c565b61092d858585856122ba565b6109398585858561240e565b7f6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c60076000878152602001908152602001600020858560405161097d929190614813565b908152602001604051809103902060010160405161099b91906148f9565b60405180910390a15050505050565b6004546060908067ffffffffffffffff8111156109c9576109c96141d8565b604051908082528060200260200182016040528015610a2757816020015b610a1460405180608001604052806000151581526020016060815260200160608152602001600081525090565b8152602001906001900390816109e75790505b50915060005b81811015610c4757600060048281548110610a4a57610a4a61490c565b906000526020600020906002020160405180604001604052908160008201548152602001600182018054610a7d90614823565b80601f0160208091040260200160405190810160405280929190818152602001828054610aa990614823565b8015610af65780601f10610acb57610100808354040283529160200191610af6565b820191906000526020600020905b815481529060010190602001808311610ad957829003601f168201915b505050505081525050905060008160000151905060008260200151905060405180608001604052806007600085815260200190815260200160002083604051610b3f919061493b565b90815260408051602092819003830190205460ff161515835260008681526007835281902090519290910191610b7690859061493b565b90815260200160405180910390206001018054610b9290614823565b80601f0160208091040260200160405190810160405280929190818152602001828054610bbe90614823565b8015610c0b5780601f10610be057610100808354040283529160200191610c0b565b820191906000526020600020905b815481529060010190602001808311610bee57829003601f168201915b5050505050815260200182815260200183815250868581518110610c3157610c3161490c565b6020908102919091010152505050600101610a2d565b505090565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3610c7681612252565b610c7e61225c565b610c8b8686868686612491565b610c988686868686612527565b857f40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee63486868686604051610cce9493929190614982565b60405180910390a2505050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610d1881612252565b610d2283836125a5565b50505050565b6000610d3381612252565b6001600160a01b038216610d73576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d9d7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3836125a5565b50610dc87f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a836125a5565b50600154610e00907ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3906001600160a01b0316612674565b50600154610e38907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b0316612674565b50600154604080516001600160a01b03928316815291841660208301527f6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279910160405180910390a150600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6001600160a01b0381163314610efd576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f078282612674565b505050565b6000610f1781612252565b610f1f61271a565b50565b610f2a61278c565b610f338261285c565b610f3d8282612867565b5050565b6000610f4b612968565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b600083815260076020526040808220905160609190610f929086908690614813565b908152604080519182900360209081018320546000898152600790925291902060ff909116935090610fc79086908690614813565b90815260200160405180910390206001018054610fe390614823565b80601f016020809104026020016040519081016040528092919081815260200182805461100f90614823565b801561105c5780601f106110315761010080835404028352916020019161105c565b820191906000526020600020905b81548152906001019060200180831161103f57829003601f168201915b50505050509050935093915050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361109581612252565b61109d61225c565b6110aa86868686866129ca565b6110b78686868686612cae565b84846040516110c7929190614813565b6040518091039020867f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce258185856040516111019291906149b4565b60405180910390a3505050505050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361113b81612252565b61114361225c565b6111548a8a8a8a8a8a8a8a8a612d2c565b6111658a8a8a8a8a8a8a8a8a613062565b896001600160a01b0316868660405161117f929190614813565b60405180910390207fa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367848a8d8d6040516111bc94939291906149c8565b60405180910390a350505050505050505050565b60606006600085815260200190815260200160002060040183836040516111f8929190614813565b9081526020016040518091039020805461121190614823565b80601f016020809104026020016040519081016040528092919081815260200182805461123d90614823565b801561128a5780601f1061125f5761010080835404028352916020019161128a565b820191906000526020600020905b81548152906001019060200180831161126d57829003601f168201915b505050505090505b9392505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6112c381612252565b610f1f6130f4565b60006112d681612252565b6001600160a01b038216611316576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113216000836125a5565b5061134c7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a836125a5565b506000805461136491906001600160a01b0316612674565b5060005461139c907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b0316612674565b50600054604080516001600160a01b03928316815291841660208301527f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f910160405180910390a150600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361144981612252565b61145161225c565b61145b838361314f565b611465838361323f565b604080516001600160a01b038516815283151560208201527f9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8910160405180910390a1505050565b606060028054806020026020016040519081016040528092919081815260200182805480156114fb57602002820191906000526020600020905b8154815260200190600101908083116114e7575b5050505050905090565b6003546060908067ffffffffffffffff811115611524576115246141d8565b60405190808252806020026020018201604052801561159357816020015b60408051608081018252600080825260208083018290529282015260608082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092019101816115425790505b50915060005b81811015610c47576000600382815481106115b6576115b661490c565b6000918252602080832090910154604080516080810182528285526006808552828620805460ff161515835282860185905260028101546001600160a01b031693830193909352948390529390925260039091018054919350606083019161161d90614823565b80601f016020809104026020016040519081016040528092919081815260200182805461164990614823565b80156116965780601f1061166b57610100808354040283529160200191611696565b820191906000526020600020905b81548152906001019060200180831161167957829003601f168201915b50505050508152508483815181106116b0576116b061490c565b602090810291909101015250600101611599565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36116ee81612252565b6116f661225c565b61170386868686866132c2565b6117108686868686613478565b857fc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e83604051610cce911515815260200190565b6000838152600a602052604080822090516117629085908590614813565b908152604051908190036020019020546001600160a01b031690509392505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156117cf5750825b905060008267ffffffffffffffff1660011480156117ec5750303b155b9050811580156117fa575080155b15611831576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156118925784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b03881615806118af57506001600160a01b038716155b806118c157506001600160a01b038616155b156118f8576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119006134f6565b6119086134f6565b6119106134fe565b61191b6000896125a5565b506119467ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3886125a5565b506119717f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a886125a5565b5061199c7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a896125a5565b50600080546001600160a01b038a81167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316178355600180548b8316908416178155600b8054928b16929093169190911790915546808352600660208181526040808620805460ff1916909517855580513060601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016818401528151808203601401815260349091019091529290945290925260030190611a609082614a3c565b50600280546001818101909255467f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9091018190556003805492830181556000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b909101558315611b275784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b6005546060908067ffffffffffffffff811115611b5057611b506141d8565b604051908082528060200260200182016040528015611bcf57816020015b611bbc6040518060e0016040528060001515815260200160006001600160a01b0316815260200160608152602001600081526020016060815260200160608152602001600060ff1681525090565b815260200190600190039081611b6e5790505b50915060005b81811015610c4757600060058281548110611bf257611bf261490c565b60009182526020808320909101546001600160a01b0390811680845260088352604093849020845160e081018652815460ff811615158252610100900490931693830193909352600183018054919550919384019190611c5190614823565b80601f0160208091040260200160405190810160405280929190818152602001828054611c7d90614823565b8015611cca5780601f10611c9f57610100808354040283529160200191611cca565b820191906000526020600020905b815481529060010190602001808311611cad57829003601f168201915b5050505050815260200160028201548152602001600382018054611ced90614823565b80601f0160208091040260200160405190810160405280929190818152602001828054611d1990614823565b8015611d665780601f10611d3b57610100808354040283529160200191611d66565b820191906000526020600020905b815481529060010190602001808311611d4957829003601f168201915b50505050508152602001600482018054611d7f90614823565b80601f0160208091040260200160405190810160405280929190818152602001828054611dab90614823565b8015611df85780601f10611dcd57610100808354040283529160200191611df8565b820191906000526020600020905b815481529060010190602001808311611ddb57829003601f168201915b50505091835250506005919091015460ff166020909101528451859084908110611e2457611e2461490c565b602090810291909101015250600101611bd5565b6060600760008781526020019081526020016000208585604051611e5d929190614813565b90815260200160405180910390206003018383604051611e7e929190614813565b90815260200160405180910390208054611e9790614823565b80601f0160208091040260200160405190810160405280929190818152602001828054611ec390614823565b8015611f105780601f10611ee557610100808354040283529160200191611f10565b820191906000526020600020905b815481529060010190602001808311611ef357829003601f168201915b5050505050905095945050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611f5a81612252565b610d228383612674565b6001600160a01b038082166000908152600860209081526040808320815160e081018352815460ff81161515825261010090049095169285019290925260018201805493946060948694869485948794859490939284019190611fc690614823565b80601f0160208091040260200160405190810160405280929190818152602001828054611ff290614823565b801561203f5780601f106120145761010080835404028352916020019161203f565b820191906000526020600020905b81548152906001019060200180831161202257829003601f168201915b505050505081526020016002820154815260200160038201805461206290614823565b80601f016020809104026020016040519081016040528092919081815260200182805461208e90614823565b80156120db5780601f106120b0576101008083540402835291602001916120db565b820191906000526020600020905b8154815290600101906020018083116120be57829003601f168201915b505050505081526020016004820180546120f490614823565b80601f016020809104026020016040519081016040528092919081815260200182805461212090614823565b801561216d5780601f106121425761010080835404028352916020019161216d565b820191906000526020600020905b81548152906001019060200180831161215057829003601f168201915b50505091835250506005919091015460ff16602090910152805160808201516060830151604084015160a085015160c090950151939d929c50909a509850919650945092505050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36121e081612252565b6121e861225c565b6121f788888888888888613531565b61220688888888888888613689565b877faea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c588888888888860405161224096959493929190614b37565b60405180910390a25050505050505050565b610f1f813361370b565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16156122b8576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60008481526006602052604090205460ff1661230a576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018590526024015b60405180910390fd5b60008290036123485782826040517ec10cfd0000000000000000000000000000000000000000000000000000000081526004016123019291906149b4565b6000848152600760205260409081902090516123679085908590614813565b9081526020016040518091039020600101805461238390614823565b90506000036123c4578383836040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161230193929190614b80565b806007600086815260200190815260200160002084846040516123e8929190614813565b908152604051908190036020019020805491151560ff1990921691909117905550505050565b6000848484846040516024016124279493929190614b9a565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f10d29b9e00000000000000000000000000000000000000000000000000000000179052905061248a81613798565b5050505050565b60008581526006602052604090205460ff166124dc576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612301565b8181600660008881526020019081526020016000206004018686604051612504929190614813565b9081526020016040518091039020918261251f929190614bc7565b505050505050565b60008585858585604051602401612542959493929190614cc3565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f2259e9e500000000000000000000000000000000000000000000000000000000179052905061251f81613798565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1661266a576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556126203390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506108e9565b60009150506108e9565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff161561266a576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506108e9565b612722613840565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061282557507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166128197f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b156122b8576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610f3d81612252565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156128c1575060408051601f3d908101601f191682019092526128be91810190614cfc565b60015b612902576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401612301565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc811461295e576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401612301565b610f07838361389b565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146122b8576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008581526006602052604090205460ff16612a15576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612301565b6000839003612a535783836040517ec10cfd0000000000000000000000000000000000000000000000000000000081526004016123019291906149b4565b6000819003612a8e576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000858152600760205260408082209051612aac9087908790614813565b90815260200160405180910390206001018054612ac890614823565b90501115612b0c5784848484846040517f2b192eab000000000000000000000000000000000000000000000000000000008152600401612301959493929190614cc3565b6001600760008781526020019081526020016000208585604051612b31929190614813565b9081526040805160209281900383018120805460ff1916941515949094179093556000888152600790925290208391839190612b709088908890614813565b90815260200160405180910390206001019182612b8e929190614bc7565b508383600760008881526020019081526020016000208686604051612bb4929190614813565b90815260200160405180910390206002019182612bd2929190614bc7565b506004604051806040016040528087815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939094525050835460018181018655948252602091829020845160029092020190815590830151929390929083019150612c529082614a3c565b5050508383604051612c65929190614813565b6040518091039020857f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce25818484604051612c9f9291906149b4565b60405180910390a35050505050565b60008585858585604051602401612cc9959493929190614cc3565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f631d62e400000000000000000000000000000000000000000000000000000000179052905061251f81613798565b6001600160a01b038916612d6c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000879003612dd6576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f53796d626f6c2063616e6e6f7420626520656d707479000000000000000000006044820152606401612301565b6001600160a01b0389811660009081526008602052604090205461010090041615612e38576040517f63f4ee1f0000000000000000000000000000000000000000000000000000000081526001600160a01b038a166004820152602401612301565b60006001600160a01b031660098989604051612e55929190614813565b908152604051908190036020019020546001600160a01b031614612ea95787876040517fb295cac90000000000000000000000000000000000000000000000000000000081526004016123019291906149b4565b6001600160a01b0389166000818152600860205260409020805460017fffffffffffffffffffffff000000000000000000000000000000000000000000909116610100909302929092178217815501612f03858783614bc7565b506001600160a01b038916600090815260086020526040902060028101879055600301612f31888a83614bc7565b506001600160a01b038916600090815260086020526040902060058101805460ff191660ff8416179055600401612f69838583614bc7565b5088600a60008881526020019081526020016000208686604051612f8e929190614813565b908152602001604051809103902060006101000a8154816001600160a01b0302191690836001600160a01b031602179055508860098989604051612fd3929190614813565b90815260405190819003602001902080546001600160a01b039283167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617909155600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180549b9092169a16999099179098555050505050505050565b600089898989898989898960405160240161308599989796959493929190614d15565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6e9e2d3f0000000000000000000000000000000000000000000000000000000017905290506130e881613798565b50505050505050505050565b6130fc61225c565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361276e565b6001600160a01b03821661318f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260086020526040902054610100900416613214576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5a52433230206e6f7420726567697374657265640000000000000000000000006044820152606401612301565b6001600160a01b03919091166000908152600860205260409020805460ff1916911515919091179055565b6040516001600160a01b0383166024820152811515604482015260009060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9060bda9000000000000000000000000000000000000000000000000000000001790529050610f0781613798565b60008581526006602052604090205460ff1680156132dd5750805b15613317576040517fa1452cb000000000000000000000000000000000000000000000000000000000815260048101869052602401612301565b60008581526006602052604090205460ff16158015613334575080155b1561336e576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612301565b6000858152600660205260409020600201546001600160a01b03161580156133965750468514155b156133d157600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018590555b6000858152600660205260409020805460ff19168215151781556002810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03871617905560030161342e838583614bc7565b50801561346f57600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0185905561248a565b61248a856138f1565b60008585858585604051602401613493959493929190614d80565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa8f2cb9600000000000000000000000000000000000000000000000000000000179052905061251f81613798565b6122b861399f565b61350661399f565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b60008781526006602052604090205460ff1661357c576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101889052602401612301565b60008590036135ba5785856040517ec10cfd0000000000000000000000000000000000000000000000000000000081526004016123019291906149b4565b6000878152600760205260409081902090516135d99088908890614813565b9081526040519081900360200190205460ff16613628578686866040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161230193929190614b80565b8181600760008a8152602001908152602001600020888860405161364d929190614813565b9081526020016040518091039020600301868660405161366e929190614813565b90815260200160405180910390209182611b27929190614bc7565b6000878787878787876040516024016136a89796959493929190614dbd565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff354b31f000000000000000000000000000000000000000000000000000000001790529050611b2781613798565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610f3d576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401612301565b60005b600254811015610f3d5746600282815481106137b9576137b961490c565b9060005260206000200154148061380d575060066000600283815481106137e2576137e261490c565b90600052602060002001548152602001908152602001600020600301805461380990614823565b1590505b61383857613838600282815481106138275761382761490c565b906000526020600020015483613a06565b60010161379b565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166122b8576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6138a482613cc8565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156138e957610f078282613d70565b610f3d613de6565b60005b600254811015610f3d5781600282815481106139125761391261490c565b906000526020600020015403613997576002805461393290600190614e0d565b815481106139425761394261490c565b9060005260206000200154600282815481106139605761396061490c565b600091825260209091200155600280548061397d5761397d614e47565b600190038181906000526020600020016000905590555050565b6001016138f4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff166122b8576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526207a120815260006020808301829052835160a0810185528281529081018290529283018190526060808401526080830152906000848152600660205260408082206002015490517ffc5fecd50000000000000000000000000000000000000000000000000000000081526207a12060048201526001600160a01b039091169190829063fc5fecd5906024016040805180830381865afa158015613ab7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613adb9190614e76565b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018290529092506001600160a01b03841691506323b872dd906064016020604051808303816000875af1158015613b4b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b6f9190614ea4565b613ba5576040517f90b8ec1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303816000875af1158015613c11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c359190614ea4565b50600b546000878152600660205260409081902090517f06cb89830000000000000000000000000000000000000000000000000000000081526001600160a01b03909216916306cb898391613c9a9160039091019086908a908a908a90600401614ec1565b600060405180830381600087803b158015613cb457600080fd5b505af11580156130e8573d6000803e3d6000fd5b806001600160a01b03163b600003613d17576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401612301565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051613d8d919061493b565b600060405180830381855af49150503d8060008114613dc8576040519150601f19603f3d011682016040523d82523d6000602084013e613dcd565b606091505b5091509150613ddd858383613e1e565b95945050505050565b34156122b8576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082613e3357613e2e82613e93565b611292565b8151158015613e4a57506001600160a01b0384163b155b15613e8c576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401612301565b5080611292565b805115613ea35780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215613ee757600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461129257600080fd5b60008083601f840112613f2957600080fd5b50813567ffffffffffffffff811115613f4157600080fd5b602083019150836020828501011115613f5957600080fd5b9250929050565b8015158114610f1f57600080fd5b60008060008060608587031215613f8457600080fd5b84359350602085013567ffffffffffffffff811115613fa257600080fd5b613fae87828801613f17565b9094509250506040850135613fc281613f60565b939692955090935050565b60005b83811015613fe8578181015183820152602001613fd0565b50506000910152565b60008151808452614009816020860160208601613fcd565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156140d2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051151586526020810151608060208801526140956080880182613ff1565b9050604082015187820360408901526140ae8282613ff1565b60609384015198909301979097525094506020938401939190910190600101614045565b50929695505050505050565b6000806000806000606086880312156140f657600080fd5b85359450602086013567ffffffffffffffff81111561411457600080fd5b61412088828901613f17565b909550935050604086013567ffffffffffffffff81111561414057600080fd5b61414c88828901613f17565b969995985093965092949392505050565b60006020828403121561416f57600080fd5b5035919050565b6001600160a01b0381168114610f1f57600080fd5b6000806040838503121561419e57600080fd5b8235915060208301356141b081614176565b809150509250929050565b6000602082840312156141cd57600080fd5b813561129281614176565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806040838503121561421a57600080fd5b823561422581614176565b9150602083013567ffffffffffffffff81111561424157600080fd5b8301601f8101851361425257600080fd5b803567ffffffffffffffff81111561426c5761426c6141d8565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff8211171561429c5761429c6141d8565b6040528181528282016020018710156142b457600080fd5b816020840160208301376000602083830101528093505050509250929050565b6000806000604084860312156142e957600080fd5b83359250602084013567ffffffffffffffff81111561430757600080fd5b61431386828701613f17565b9497909650939450505050565b821515815260406020820152600061433b6040830184613ff1565b949350505050565b600080600080600080600080600060c08a8c03121561436157600080fd5b893561436c81614176565b985060208a013567ffffffffffffffff81111561438857600080fd5b6143948c828d01613f17565b90995097505060408a0135955060608a013567ffffffffffffffff8111156143bb57600080fd5b6143c78c828d01613f17565b90965094505060808a013567ffffffffffffffff8111156143e757600080fd5b6143f38c828d01613f17565b90945092505060a08a013560ff8116811461440d57600080fd5b809150509295985092959850929598565b6020815260006112926020830184613ff1565b6000806040838503121561444457600080fd5b823561444f81614176565b915060208301356141b081613f60565b602080825282518282018190526000918401906040840190835b81811015614497578351835260209384019390920191600101614479565b509095945050505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156140d2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160208701526001600160a01b03604082015116604087015260608101519050608060608701526145396080870182613ff1565b95505060209384019391909101906001016144ca565b60008060008060006080868803121561456757600080fd5b85359450602086013561457981614176565b9350604086013567ffffffffffffffff81111561459557600080fd5b6145a188828901613f17565b90945092505060608601356145b581613f60565b809150509295509295909350565b6000806000606084860312156145d857600080fd5b83356145e381614176565b925060208401356145f381614176565b9150604084013561460381614176565b809150509250925092565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156140d2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051151586526001600160a01b036020820151166020870152604081015160e0604088015261469960e0880182613ff1565b905060608201516060880152608082015187820360808901526146bc8282613ff1565b91505060a082015187820360a08901526146d68282613ff1565b91505060c082015191506146ef60c088018360ff169052565b9550506020938401939190910190600101614636565b861515815260c06020820152600061472060c0830188613ff1565b86604084015282810360608401526147388187613ff1565b9050828103608084015261474c8186613ff1565b91505060ff831660a0830152979650505050505050565b60008060008060008060006080888a03121561477e57600080fd5b87359650602088013567ffffffffffffffff81111561479c57600080fd5b6147a88a828b01613f17565b909750955050604088013567ffffffffffffffff8111156147c857600080fd5b6147d48a828b01613f17565b909550935050606088013567ffffffffffffffff8111156147f457600080fd5b6148008a828b01613f17565b989b979a50959850939692959293505050565b8183823760009101908152919050565b600181811c9082168061483757607f821691505b602082108103614870577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b6000815461488381614823565b80855260018216801561489d57600181146148b9576148f0565b60ff1983166020870152602082151560051b87010193506148f0565b84600052602060002060005b838110156148e75781546020828a0101526001820191506020810190506148c5565b87016020019450505b50505092915050565b6020815260006112926020830184614876565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000825161494d818460208701613fcd565b9190910192915050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b604081526000614996604083018688614957565b82810360208401526149a9818587614957565b979650505050505050565b60208152600061433b602083018486614957565b60ff851681528360208201526060604082015260006149eb606083018486614957565b9695505050505050565b601f821115610f0757806000526020600020601f840160051c81016020851015614a1c5750805b601f840160051c820191505b8181101561248a5760008155600101614a28565b815167ffffffffffffffff811115614a5657614a566141d8565b614a6a81614a648454614823565b846149f5565b6020601f821160018114614abc5760008315614a865750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b17845561248a565b600084815260208120601f198516915b82811015614aec5787850151825560209485019460019092019101614acc565b5084821015614b2857868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b606081526000614b4b60608301888a614957565b8281036020840152614b5e818789614957565b90508281036040840152614b73818587614957565b9998505050505050505050565b838152604060208201526000613ddd604083018486614957565b848152606060208201526000614bb4606083018587614957565b9050821515604083015295945050505050565b67ffffffffffffffff831115614bdf57614bdf6141d8565b614bf383614bed8354614823565b836149f5565b6000601f841160018114614c455760008515614c0f5750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b17835561248a565b600083815260209020601f19861690835b82811015614c765786850135825560209485019460019092019101614c56565b5086821015614cb1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b858152606060208201526000614cdd606083018688614957565b8281036040840152614cf0818587614957565b98975050505050505050565b600060208284031215614d0e57600080fd5b5051919050565b6001600160a01b038a16815260c060208201526000614d3860c083018a8c614957565b8860408401528281036060840152614d5181888a614957565b90508281036080840152614d66818688614957565b91505060ff831660a08301529a9950505050505050505050565b8581526001600160a01b0385166020820152608060408201526000614da9608083018587614957565b905082151560608301529695505050505050565b878152608060208201526000614dd760808301888a614957565b8281036040840152614dea818789614957565b90508281036060840152614dff818587614957565b9a9950505050505050505050565b818103818111156108e9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008060408385031215614e8957600080fd5b8251614e9481614176565b6020939093015192949293505050565b600060208284031215614eb657600080fd5b815161129281613f60565b60c081526000614ed460c0830188614876565b6001600160a01b03871660208401528281036040840152614ef58187613ff1565b90508451606084015260208501511515608084015282810360a08401526001600160a01b0384511681526020840151151560208201526001600160a01b036040850151166040820152606084015160a06060830152614f5760a0830182613ff1565b9050608085015160808301528092505050969550505050505056fea2646970667358221220e10a5e884a982853f44b63ef0d0b3817d5fe3eb6137d384d35bc5978c7af242c64736f6c634300081a0033",
}

// CoreRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use CoreRegistryMetaData.ABI instead.
var CoreRegistryABI = CoreRegistryMetaData.ABI

// CoreRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CoreRegistryMetaData.Bin instead.
var CoreRegistryBin = CoreRegistryMetaData.Bin

// DeployCoreRegistry deploys a new Ethereum contract, binding an instance of CoreRegistry to it.
func DeployCoreRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CoreRegistry, error) {
	parsed, err := CoreRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CoreRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CoreRegistry{CoreRegistryCaller: CoreRegistryCaller{contract: contract}, CoreRegistryTransactor: CoreRegistryTransactor{contract: contract}, CoreRegistryFilterer: CoreRegistryFilterer{contract: contract}}, nil
}

// CoreRegistry is an auto generated Go binding around an Ethereum contract.
type CoreRegistry struct {
	CoreRegistryCaller     // Read-only binding to the contract
	CoreRegistryTransactor // Write-only binding to the contract
	CoreRegistryFilterer   // Log filterer for contract events
}

// CoreRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoreRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoreRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoreRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoreRegistrySession struct {
	Contract     *CoreRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoreRegistryCallerSession struct {
	Contract *CoreRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CoreRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoreRegistryTransactorSession struct {
	Contract     *CoreRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CoreRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoreRegistryRaw struct {
	Contract *CoreRegistry // Generic contract binding to access the raw methods on
}

// CoreRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoreRegistryCallerRaw struct {
	Contract *CoreRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CoreRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoreRegistryTransactorRaw struct {
	Contract *CoreRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoreRegistry creates a new instance of CoreRegistry, bound to a specific deployed contract.
func NewCoreRegistry(address common.Address, backend bind.ContractBackend) (*CoreRegistry, error) {
	contract, err := bindCoreRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoreRegistry{CoreRegistryCaller: CoreRegistryCaller{contract: contract}, CoreRegistryTransactor: CoreRegistryTransactor{contract: contract}, CoreRegistryFilterer: CoreRegistryFilterer{contract: contract}}, nil
}

// NewCoreRegistryCaller creates a new read-only instance of CoreRegistry, bound to a specific deployed contract.
func NewCoreRegistryCaller(address common.Address, caller bind.ContractCaller) (*CoreRegistryCaller, error) {
	contract, err := bindCoreRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryCaller{contract: contract}, nil
}

// NewCoreRegistryTransactor creates a new write-only instance of CoreRegistry, bound to a specific deployed contract.
func NewCoreRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CoreRegistryTransactor, error) {
	contract, err := bindCoreRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTransactor{contract: contract}, nil
}

// NewCoreRegistryFilterer creates a new log filterer instance of CoreRegistry, bound to a specific deployed contract.
func NewCoreRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CoreRegistryFilterer, error) {
	contract, err := bindCoreRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryFilterer{contract: contract}, nil
}

// bindCoreRegistry binds a generic wrapper to an already deployed contract.
func bindCoreRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoreRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoreRegistry *CoreRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoreRegistry.Contract.CoreRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoreRegistry *CoreRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistry.Contract.CoreRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoreRegistry *CoreRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoreRegistry.Contract.CoreRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoreRegistry *CoreRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoreRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoreRegistry *CoreRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoreRegistry *CoreRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoreRegistry.Contract.contract.Transact(opts, method, params...)
}

// CROSSCHAINGASLIMIT is a free data retrieval call binding the contract method 0xa3ebd14c.
//
// Solidity: function CROSS_CHAIN_GAS_LIMIT() view returns(uint256)
func (_CoreRegistry *CoreRegistryCaller) CROSSCHAINGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "CROSS_CHAIN_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CROSSCHAINGASLIMIT is a free data retrieval call binding the contract method 0xa3ebd14c.
//
// Solidity: function CROSS_CHAIN_GAS_LIMIT() view returns(uint256)
func (_CoreRegistry *CoreRegistrySession) CROSSCHAINGASLIMIT() (*big.Int, error) {
	return _CoreRegistry.Contract.CROSSCHAINGASLIMIT(&_CoreRegistry.CallOpts)
}

// CROSSCHAINGASLIMIT is a free data retrieval call binding the contract method 0xa3ebd14c.
//
// Solidity: function CROSS_CHAIN_GAS_LIMIT() view returns(uint256)
func (_CoreRegistry *CoreRegistryCallerSession) CROSSCHAINGASLIMIT() (*big.Int, error) {
	return _CoreRegistry.Contract.CROSSCHAINGASLIMIT(&_CoreRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CoreRegistry *CoreRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CoreRegistry *CoreRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CoreRegistry.Contract.DEFAULTADMINROLE(&_CoreRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CoreRegistry *CoreRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CoreRegistry.Contract.DEFAULTADMINROLE(&_CoreRegistry.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_CoreRegistry *CoreRegistryCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_CoreRegistry *CoreRegistrySession) PAUSERROLE() ([32]byte, error) {
	return _CoreRegistry.Contract.PAUSERROLE(&_CoreRegistry.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_CoreRegistry *CoreRegistryCallerSession) PAUSERROLE() ([32]byte, error) {
	return _CoreRegistry.Contract.PAUSERROLE(&_CoreRegistry.CallOpts)
}

// REGISTRYMANAGERROLE is a free data retrieval call binding the contract method 0xbd8fde1c.
//
// Solidity: function REGISTRY_MANAGER_ROLE() view returns(bytes32)
func (_CoreRegistry *CoreRegistryCaller) REGISTRYMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "REGISTRY_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRYMANAGERROLE is a free data retrieval call binding the contract method 0xbd8fde1c.
//
// Solidity: function REGISTRY_MANAGER_ROLE() view returns(bytes32)
func (_CoreRegistry *CoreRegistrySession) REGISTRYMANAGERROLE() ([32]byte, error) {
	return _CoreRegistry.Contract.REGISTRYMANAGERROLE(&_CoreRegistry.CallOpts)
}

// REGISTRYMANAGERROLE is a free data retrieval call binding the contract method 0xbd8fde1c.
//
// Solidity: function REGISTRY_MANAGER_ROLE() view returns(bytes32)
func (_CoreRegistry *CoreRegistryCallerSession) REGISTRYMANAGERROLE() ([32]byte, error) {
	return _CoreRegistry.Contract.REGISTRYMANAGERROLE(&_CoreRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CoreRegistry *CoreRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CoreRegistry *CoreRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CoreRegistry.Contract.UPGRADEINTERFACEVERSION(&_CoreRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CoreRegistry *CoreRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CoreRegistry.Contract.UPGRADEINTERFACEVERSION(&_CoreRegistry.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CoreRegistry *CoreRegistryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CoreRegistry *CoreRegistrySession) Admin() (common.Address, error) {
	return _CoreRegistry.Contract.Admin(&_CoreRegistry.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CoreRegistry *CoreRegistryCallerSession) Admin() (common.Address, error) {
	return _CoreRegistry.Contract.Admin(&_CoreRegistry.CallOpts)
}

// GatewayZEVM is a free data retrieval call binding the contract method 0xcc5ad8b6.
//
// Solidity: function gatewayZEVM() view returns(address)
func (_CoreRegistry *CoreRegistryCaller) GatewayZEVM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "gatewayZEVM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayZEVM is a free data retrieval call binding the contract method 0xcc5ad8b6.
//
// Solidity: function gatewayZEVM() view returns(address)
func (_CoreRegistry *CoreRegistrySession) GatewayZEVM() (common.Address, error) {
	return _CoreRegistry.Contract.GatewayZEVM(&_CoreRegistry.CallOpts)
}

// GatewayZEVM is a free data retrieval call binding the contract method 0xcc5ad8b6.
//
// Solidity: function gatewayZEVM() view returns(address)
func (_CoreRegistry *CoreRegistryCallerSession) GatewayZEVM() (common.Address, error) {
	return _CoreRegistry.Contract.GatewayZEVM(&_CoreRegistry.CallOpts)
}

// GetActiveChains is a free data retrieval call binding the contract method 0x94cc8683.
//
// Solidity: function getActiveChains() view returns(uint256[])
func (_CoreRegistry *CoreRegistryCaller) GetActiveChains(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "getActiveChains")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveChains is a free data retrieval call binding the contract method 0x94cc8683.
//
// Solidity: function getActiveChains() view returns(uint256[])
func (_CoreRegistry *CoreRegistrySession) GetActiveChains() ([]*big.Int, error) {
	return _CoreRegistry.Contract.GetActiveChains(&_CoreRegistry.CallOpts)
}

// GetActiveChains is a free data retrieval call binding the contract method 0x94cc8683.
//
// Solidity: function getActiveChains() view returns(uint256[])
func (_CoreRegistry *CoreRegistryCallerSession) GetActiveChains() ([]*big.Int, error) {
	return _CoreRegistry.Contract.GetActiveChains(&_CoreRegistry.CallOpts)
}

// GetAllChains is a free data retrieval call binding the contract method 0x9ca220dd.
//
// Solidity: function getAllChains() view returns((bool,uint256,address,bytes)[] chainsInfo)
func (_CoreRegistry *CoreRegistryCaller) GetAllChains(opts *bind.CallOpts) ([]ChainInfoDTO, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "getAllChains")

	if err != nil {
		return *new([]ChainInfoDTO), err
	}

	out0 := *abi.ConvertType(out[0], new([]ChainInfoDTO)).(*[]ChainInfoDTO)

	return out0, err

}

// GetAllChains is a free data retrieval call binding the contract method 0x9ca220dd.
//
// Solidity: function getAllChains() view returns((bool,uint256,address,bytes)[] chainsInfo)
func (_CoreRegistry *CoreRegistrySession) GetAllChains() ([]ChainInfoDTO, error) {
	return _CoreRegistry.Contract.GetAllChains(&_CoreRegistry.CallOpts)
}

// GetAllChains is a free data retrieval call binding the contract method 0x9ca220dd.
//
// Solidity: function getAllChains() view returns((bool,uint256,address,bytes)[] chainsInfo)
func (_CoreRegistry *CoreRegistryCallerSession) GetAllChains() ([]ChainInfoDTO, error) {
	return _CoreRegistry.Contract.GetAllChains(&_CoreRegistry.CallOpts)
}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns((bool,bytes,string,uint256)[] contractsInfo)
func (_CoreRegistry *CoreRegistryCaller) GetAllContracts(opts *bind.CallOpts) ([]ContractInfoDTO, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "getAllContracts")

	if err != nil {
		return *new([]ContractInfoDTO), err
	}

	out0 := *abi.ConvertType(out[0], new([]ContractInfoDTO)).(*[]ContractInfoDTO)

	return out0, err

}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns((bool,bytes,string,uint256)[] contractsInfo)
func (_CoreRegistry *CoreRegistrySession) GetAllContracts() ([]ContractInfoDTO, error) {
	return _CoreRegistry.Contract.GetAllContracts(&_CoreRegistry.CallOpts)
}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns((bool,bytes,string,uint256)[] contractsInfo)
func (_CoreRegistry *CoreRegistryCallerSession) GetAllContracts() ([]ContractInfoDTO, error) {
	return _CoreRegistry.Contract.GetAllContracts(&_CoreRegistry.CallOpts)
}

// GetAllZRC20Tokens is a free data retrieval call binding the contract method 0xc1bd469f.
//
// Solidity: function getAllZRC20Tokens() view returns((bool,address,bytes,uint256,string,string,uint8)[] tokensInfo)
func (_CoreRegistry *CoreRegistryCaller) GetAllZRC20Tokens(opts *bind.CallOpts) ([]ZRC20Info, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "getAllZRC20Tokens")

	if err != nil {
		return *new([]ZRC20Info), err
	}

	out0 := *abi.ConvertType(out[0], new([]ZRC20Info)).(*[]ZRC20Info)

	return out0, err

}

// GetAllZRC20Tokens is a free data retrieval call binding the contract method 0xc1bd469f.
//
// Solidity: function getAllZRC20Tokens() view returns((bool,address,bytes,uint256,string,string,uint8)[] tokensInfo)
func (_CoreRegistry *CoreRegistrySession) GetAllZRC20Tokens() ([]ZRC20Info, error) {
	return _CoreRegistry.Contract.GetAllZRC20Tokens(&_CoreRegistry.CallOpts)
}

// GetAllZRC20Tokens is a free data retrieval call binding the contract method 0xc1bd469f.
//
// Solidity: function getAllZRC20Tokens() view returns((bool,address,bytes,uint256,string,string,uint8)[] tokensInfo)
func (_CoreRegistry *CoreRegistryCallerSession) GetAllZRC20Tokens() ([]ZRC20Info, error) {
	return _CoreRegistry.Contract.GetAllZRC20Tokens(&_CoreRegistry.CallOpts)
}

// GetChainMetadata is a free data retrieval call binding the contract method 0x7066b18d.
//
// Solidity: function getChainMetadata(uint256 chainId, string key) view returns(bytes)
func (_CoreRegistry *CoreRegistryCaller) GetChainMetadata(opts *bind.CallOpts, chainId *big.Int, key string) ([]byte, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "getChainMetadata", chainId, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetChainMetadata is a free data retrieval call binding the contract method 0x7066b18d.
//
// Solidity: function getChainMetadata(uint256 chainId, string key) view returns(bytes)
func (_CoreRegistry *CoreRegistrySession) GetChainMetadata(chainId *big.Int, key string) ([]byte, error) {
	return _CoreRegistry.Contract.GetChainMetadata(&_CoreRegistry.CallOpts, chainId, key)
}

// GetChainMetadata is a free data retrieval call binding the contract method 0x7066b18d.
//
// Solidity: function getChainMetadata(uint256 chainId, string key) view returns(bytes)
func (_CoreRegistry *CoreRegistryCallerSession) GetChainMetadata(chainId *big.Int, key string) ([]byte, error) {
	return _CoreRegistry.Contract.GetChainMetadata(&_CoreRegistry.CallOpts, chainId, key)
}

// GetContractConfiguration is a free data retrieval call binding the contract method 0xd3523ea2.
//
// Solidity: function getContractConfiguration(uint256 chainId, string contractType, string key) view returns(bytes)
func (_CoreRegistry *CoreRegistryCaller) GetContractConfiguration(opts *bind.CallOpts, chainId *big.Int, contractType string, key string) ([]byte, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "getContractConfiguration", chainId, contractType, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetContractConfiguration is a free data retrieval call binding the contract method 0xd3523ea2.
//
// Solidity: function getContractConfiguration(uint256 chainId, string contractType, string key) view returns(bytes)
func (_CoreRegistry *CoreRegistrySession) GetContractConfiguration(chainId *big.Int, contractType string, key string) ([]byte, error) {
	return _CoreRegistry.Contract.GetContractConfiguration(&_CoreRegistry.CallOpts, chainId, contractType, key)
}

// GetContractConfiguration is a free data retrieval call binding the contract method 0xd3523ea2.
//
// Solidity: function getContractConfiguration(uint256 chainId, string contractType, string key) view returns(bytes)
func (_CoreRegistry *CoreRegistryCallerSession) GetContractConfiguration(chainId *big.Int, contractType string, key string) ([]byte, error) {
	return _CoreRegistry.Contract.GetContractConfiguration(&_CoreRegistry.CallOpts, chainId, contractType, key)
}

// GetContractInfo is a free data retrieval call binding the contract method 0x5cf92c9f.
//
// Solidity: function getContractInfo(uint256 chainId, string contractType) view returns(bool active, bytes addressBytes)
func (_CoreRegistry *CoreRegistryCaller) GetContractInfo(opts *bind.CallOpts, chainId *big.Int, contractType string) (struct {
	Active       bool
	AddressBytes []byte
}, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "getContractInfo", chainId, contractType)

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
func (_CoreRegistry *CoreRegistrySession) GetContractInfo(chainId *big.Int, contractType string) (struct {
	Active       bool
	AddressBytes []byte
}, error) {
	return _CoreRegistry.Contract.GetContractInfo(&_CoreRegistry.CallOpts, chainId, contractType)
}

// GetContractInfo is a free data retrieval call binding the contract method 0x5cf92c9f.
//
// Solidity: function getContractInfo(uint256 chainId, string contractType) view returns(bool active, bytes addressBytes)
func (_CoreRegistry *CoreRegistryCallerSession) GetContractInfo(chainId *big.Int, contractType string) (struct {
	Active       bool
	AddressBytes []byte
}, error) {
	return _CoreRegistry.Contract.GetContractInfo(&_CoreRegistry.CallOpts, chainId, contractType)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CoreRegistry *CoreRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CoreRegistry *CoreRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CoreRegistry.Contract.GetRoleAdmin(&_CoreRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CoreRegistry *CoreRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CoreRegistry.Contract.GetRoleAdmin(&_CoreRegistry.CallOpts, role)
}

// GetZRC20AddressByForeignAsset is a free data retrieval call binding the contract method 0xaa808c06.
//
// Solidity: function getZRC20AddressByForeignAsset(uint256 originChainId, bytes originAddress) view returns(address)
func (_CoreRegistry *CoreRegistryCaller) GetZRC20AddressByForeignAsset(opts *bind.CallOpts, originChainId *big.Int, originAddress []byte) (common.Address, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "getZRC20AddressByForeignAsset", originChainId, originAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetZRC20AddressByForeignAsset is a free data retrieval call binding the contract method 0xaa808c06.
//
// Solidity: function getZRC20AddressByForeignAsset(uint256 originChainId, bytes originAddress) view returns(address)
func (_CoreRegistry *CoreRegistrySession) GetZRC20AddressByForeignAsset(originChainId *big.Int, originAddress []byte) (common.Address, error) {
	return _CoreRegistry.Contract.GetZRC20AddressByForeignAsset(&_CoreRegistry.CallOpts, originChainId, originAddress)
}

// GetZRC20AddressByForeignAsset is a free data retrieval call binding the contract method 0xaa808c06.
//
// Solidity: function getZRC20AddressByForeignAsset(uint256 originChainId, bytes originAddress) view returns(address)
func (_CoreRegistry *CoreRegistryCallerSession) GetZRC20AddressByForeignAsset(originChainId *big.Int, originAddress []byte) (common.Address, error) {
	return _CoreRegistry.Contract.GetZRC20AddressByForeignAsset(&_CoreRegistry.CallOpts, originChainId, originAddress)
}

// GetZRC20TokenInfo is a free data retrieval call binding the contract method 0xe9d6c5ba.
//
// Solidity: function getZRC20TokenInfo(address address_) view returns(bool active, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals)
func (_CoreRegistry *CoreRegistryCaller) GetZRC20TokenInfo(opts *bind.CallOpts, address_ common.Address) (struct {
	Active        bool
	Symbol        string
	OriginChainId *big.Int
	OriginAddress []byte
	CoinType      string
	Decimals      uint8
}, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "getZRC20TokenInfo", address_)

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
func (_CoreRegistry *CoreRegistrySession) GetZRC20TokenInfo(address_ common.Address) (struct {
	Active        bool
	Symbol        string
	OriginChainId *big.Int
	OriginAddress []byte
	CoinType      string
	Decimals      uint8
}, error) {
	return _CoreRegistry.Contract.GetZRC20TokenInfo(&_CoreRegistry.CallOpts, address_)
}

// GetZRC20TokenInfo is a free data retrieval call binding the contract method 0xe9d6c5ba.
//
// Solidity: function getZRC20TokenInfo(address address_) view returns(bool active, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals)
func (_CoreRegistry *CoreRegistryCallerSession) GetZRC20TokenInfo(address_ common.Address) (struct {
	Active        bool
	Symbol        string
	OriginChainId *big.Int
	OriginAddress []byte
	CoinType      string
	Decimals      uint8
}, error) {
	return _CoreRegistry.Contract.GetZRC20TokenInfo(&_CoreRegistry.CallOpts, address_)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CoreRegistry *CoreRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CoreRegistry *CoreRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CoreRegistry.Contract.HasRole(&_CoreRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CoreRegistry *CoreRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CoreRegistry.Contract.HasRole(&_CoreRegistry.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CoreRegistry *CoreRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CoreRegistry *CoreRegistrySession) Paused() (bool, error) {
	return _CoreRegistry.Contract.Paused(&_CoreRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CoreRegistry *CoreRegistryCallerSession) Paused() (bool, error) {
	return _CoreRegistry.Contract.Paused(&_CoreRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CoreRegistry *CoreRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CoreRegistry *CoreRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _CoreRegistry.Contract.ProxiableUUID(&_CoreRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CoreRegistry *CoreRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CoreRegistry.Contract.ProxiableUUID(&_CoreRegistry.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_CoreRegistry *CoreRegistryCaller) RegistryManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "registryManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_CoreRegistry *CoreRegistrySession) RegistryManager() (common.Address, error) {
	return _CoreRegistry.Contract.RegistryManager(&_CoreRegistry.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_CoreRegistry *CoreRegistryCallerSession) RegistryManager() (common.Address, error) {
	return _CoreRegistry.Contract.RegistryManager(&_CoreRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoreRegistry *CoreRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CoreRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoreRegistry *CoreRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CoreRegistry.Contract.SupportsInterface(&_CoreRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoreRegistry *CoreRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CoreRegistry.Contract.SupportsInterface(&_CoreRegistry.CallOpts, interfaceId)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_CoreRegistry *CoreRegistryTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_CoreRegistry *CoreRegistrySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _CoreRegistry.Contract.ChangeAdmin(&_CoreRegistry.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _CoreRegistry.Contract.ChangeAdmin(&_CoreRegistry.TransactOpts, newAdmin)
}

// ChangeChainStatus is a paid mutator transaction binding the contract method 0xa8f2cb96.
//
// Solidity: function changeChainStatus(uint256 chainId, address gasZRC20, bytes registry, bool activation) returns()
func (_CoreRegistry *CoreRegistryTransactor) ChangeChainStatus(opts *bind.TransactOpts, chainId *big.Int, gasZRC20 common.Address, registry []byte, activation bool) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "changeChainStatus", chainId, gasZRC20, registry, activation)
}

// ChangeChainStatus is a paid mutator transaction binding the contract method 0xa8f2cb96.
//
// Solidity: function changeChainStatus(uint256 chainId, address gasZRC20, bytes registry, bool activation) returns()
func (_CoreRegistry *CoreRegistrySession) ChangeChainStatus(chainId *big.Int, gasZRC20 common.Address, registry []byte, activation bool) (*types.Transaction, error) {
	return _CoreRegistry.Contract.ChangeChainStatus(&_CoreRegistry.TransactOpts, chainId, gasZRC20, registry, activation)
}

// ChangeChainStatus is a paid mutator transaction binding the contract method 0xa8f2cb96.
//
// Solidity: function changeChainStatus(uint256 chainId, address gasZRC20, bytes registry, bool activation) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) ChangeChainStatus(chainId *big.Int, gasZRC20 common.Address, registry []byte, activation bool) (*types.Transaction, error) {
	return _CoreRegistry.Contract.ChangeChainStatus(&_CoreRegistry.TransactOpts, chainId, gasZRC20, registry, activation)
}

// ChangeRegistryManager is a paid mutator transaction binding the contract method 0x3500c24b.
//
// Solidity: function changeRegistryManager(address newRegistryManager) returns()
func (_CoreRegistry *CoreRegistryTransactor) ChangeRegistryManager(opts *bind.TransactOpts, newRegistryManager common.Address) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "changeRegistryManager", newRegistryManager)
}

// ChangeRegistryManager is a paid mutator transaction binding the contract method 0x3500c24b.
//
// Solidity: function changeRegistryManager(address newRegistryManager) returns()
func (_CoreRegistry *CoreRegistrySession) ChangeRegistryManager(newRegistryManager common.Address) (*types.Transaction, error) {
	return _CoreRegistry.Contract.ChangeRegistryManager(&_CoreRegistry.TransactOpts, newRegistryManager)
}

// ChangeRegistryManager is a paid mutator transaction binding the contract method 0x3500c24b.
//
// Solidity: function changeRegistryManager(address newRegistryManager) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) ChangeRegistryManager(newRegistryManager common.Address) (*types.Transaction, error) {
	return _CoreRegistry.Contract.ChangeRegistryManager(&_CoreRegistry.TransactOpts, newRegistryManager)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CoreRegistry *CoreRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CoreRegistry *CoreRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreRegistry.Contract.GrantRole(&_CoreRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreRegistry.Contract.GrantRole(&_CoreRegistry.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address admin_, address registryManager_, address gatewayZEVM_) returns()
func (_CoreRegistry *CoreRegistryTransactor) Initialize(opts *bind.TransactOpts, admin_ common.Address, registryManager_ common.Address, gatewayZEVM_ common.Address) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "initialize", admin_, registryManager_, gatewayZEVM_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address admin_, address registryManager_, address gatewayZEVM_) returns()
func (_CoreRegistry *CoreRegistrySession) Initialize(admin_ common.Address, registryManager_ common.Address, gatewayZEVM_ common.Address) (*types.Transaction, error) {
	return _CoreRegistry.Contract.Initialize(&_CoreRegistry.TransactOpts, admin_, registryManager_, gatewayZEVM_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address admin_, address registryManager_, address gatewayZEVM_) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) Initialize(admin_ common.Address, registryManager_ common.Address, gatewayZEVM_ common.Address) (*types.Transaction, error) {
	return _CoreRegistry.Contract.Initialize(&_CoreRegistry.TransactOpts, admin_, registryManager_, gatewayZEVM_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CoreRegistry *CoreRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CoreRegistry *CoreRegistrySession) Pause() (*types.Transaction, error) {
	return _CoreRegistry.Contract.Pause(&_CoreRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CoreRegistry *CoreRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _CoreRegistry.Contract.Pause(&_CoreRegistry.TransactOpts)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x631d62e4.
//
// Solidity: function registerContract(uint256 chainId, string contractType, bytes addressBytes) returns()
func (_CoreRegistry *CoreRegistryTransactor) RegisterContract(opts *bind.TransactOpts, chainId *big.Int, contractType string, addressBytes []byte) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "registerContract", chainId, contractType, addressBytes)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x631d62e4.
//
// Solidity: function registerContract(uint256 chainId, string contractType, bytes addressBytes) returns()
func (_CoreRegistry *CoreRegistrySession) RegisterContract(chainId *big.Int, contractType string, addressBytes []byte) (*types.Transaction, error) {
	return _CoreRegistry.Contract.RegisterContract(&_CoreRegistry.TransactOpts, chainId, contractType, addressBytes)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x631d62e4.
//
// Solidity: function registerContract(uint256 chainId, string contractType, bytes addressBytes) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) RegisterContract(chainId *big.Int, contractType string, addressBytes []byte) (*types.Transaction, error) {
	return _CoreRegistry.Contract.RegisterContract(&_CoreRegistry.TransactOpts, chainId, contractType, addressBytes)
}

// RegisterZRC20Token is a paid mutator transaction binding the contract method 0x6e9e2d3f.
//
// Solidity: function registerZRC20Token(address address_, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals) returns()
func (_CoreRegistry *CoreRegistryTransactor) RegisterZRC20Token(opts *bind.TransactOpts, address_ common.Address, symbol string, originChainId *big.Int, originAddress []byte, coinType string, decimals uint8) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "registerZRC20Token", address_, symbol, originChainId, originAddress, coinType, decimals)
}

// RegisterZRC20Token is a paid mutator transaction binding the contract method 0x6e9e2d3f.
//
// Solidity: function registerZRC20Token(address address_, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals) returns()
func (_CoreRegistry *CoreRegistrySession) RegisterZRC20Token(address_ common.Address, symbol string, originChainId *big.Int, originAddress []byte, coinType string, decimals uint8) (*types.Transaction, error) {
	return _CoreRegistry.Contract.RegisterZRC20Token(&_CoreRegistry.TransactOpts, address_, symbol, originChainId, originAddress, coinType, decimals)
}

// RegisterZRC20Token is a paid mutator transaction binding the contract method 0x6e9e2d3f.
//
// Solidity: function registerZRC20Token(address address_, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) RegisterZRC20Token(address_ common.Address, symbol string, originChainId *big.Int, originAddress []byte, coinType string, decimals uint8) (*types.Transaction, error) {
	return _CoreRegistry.Contract.RegisterZRC20Token(&_CoreRegistry.TransactOpts, address_, symbol, originChainId, originAddress, coinType, decimals)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CoreRegistry *CoreRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CoreRegistry *CoreRegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CoreRegistry.Contract.RenounceRole(&_CoreRegistry.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CoreRegistry.Contract.RenounceRole(&_CoreRegistry.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CoreRegistry *CoreRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CoreRegistry *CoreRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreRegistry.Contract.RevokeRole(&_CoreRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CoreRegistry.Contract.RevokeRole(&_CoreRegistry.TransactOpts, role, account)
}

// SetContractActive is a paid mutator transaction binding the contract method 0x10d29b9e.
//
// Solidity: function setContractActive(uint256 chainId, string contractType, bool active) returns()
func (_CoreRegistry *CoreRegistryTransactor) SetContractActive(opts *bind.TransactOpts, chainId *big.Int, contractType string, active bool) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "setContractActive", chainId, contractType, active)
}

// SetContractActive is a paid mutator transaction binding the contract method 0x10d29b9e.
//
// Solidity: function setContractActive(uint256 chainId, string contractType, bool active) returns()
func (_CoreRegistry *CoreRegistrySession) SetContractActive(chainId *big.Int, contractType string, active bool) (*types.Transaction, error) {
	return _CoreRegistry.Contract.SetContractActive(&_CoreRegistry.TransactOpts, chainId, contractType, active)
}

// SetContractActive is a paid mutator transaction binding the contract method 0x10d29b9e.
//
// Solidity: function setContractActive(uint256 chainId, string contractType, bool active) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) SetContractActive(chainId *big.Int, contractType string, active bool) (*types.Transaction, error) {
	return _CoreRegistry.Contract.SetContractActive(&_CoreRegistry.TransactOpts, chainId, contractType, active)
}

// SetZRC20TokenActive is a paid mutator transaction binding the contract method 0x9060bda9.
//
// Solidity: function setZRC20TokenActive(address address_, bool active) returns()
func (_CoreRegistry *CoreRegistryTransactor) SetZRC20TokenActive(opts *bind.TransactOpts, address_ common.Address, active bool) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "setZRC20TokenActive", address_, active)
}

// SetZRC20TokenActive is a paid mutator transaction binding the contract method 0x9060bda9.
//
// Solidity: function setZRC20TokenActive(address address_, bool active) returns()
func (_CoreRegistry *CoreRegistrySession) SetZRC20TokenActive(address_ common.Address, active bool) (*types.Transaction, error) {
	return _CoreRegistry.Contract.SetZRC20TokenActive(&_CoreRegistry.TransactOpts, address_, active)
}

// SetZRC20TokenActive is a paid mutator transaction binding the contract method 0x9060bda9.
//
// Solidity: function setZRC20TokenActive(address address_, bool active) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) SetZRC20TokenActive(address_ common.Address, active bool) (*types.Transaction, error) {
	return _CoreRegistry.Contract.SetZRC20TokenActive(&_CoreRegistry.TransactOpts, address_, active)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CoreRegistry *CoreRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CoreRegistry *CoreRegistrySession) Unpause() (*types.Transaction, error) {
	return _CoreRegistry.Contract.Unpause(&_CoreRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CoreRegistry *CoreRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _CoreRegistry.Contract.Unpause(&_CoreRegistry.TransactOpts)
}

// UpdateChainMetadata is a paid mutator transaction binding the contract method 0x2259e9e5.
//
// Solidity: function updateChainMetadata(uint256 chainId, string key, bytes value) returns()
func (_CoreRegistry *CoreRegistryTransactor) UpdateChainMetadata(opts *bind.TransactOpts, chainId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "updateChainMetadata", chainId, key, value)
}

// UpdateChainMetadata is a paid mutator transaction binding the contract method 0x2259e9e5.
//
// Solidity: function updateChainMetadata(uint256 chainId, string key, bytes value) returns()
func (_CoreRegistry *CoreRegistrySession) UpdateChainMetadata(chainId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _CoreRegistry.Contract.UpdateChainMetadata(&_CoreRegistry.TransactOpts, chainId, key, value)
}

// UpdateChainMetadata is a paid mutator transaction binding the contract method 0x2259e9e5.
//
// Solidity: function updateChainMetadata(uint256 chainId, string key, bytes value) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) UpdateChainMetadata(chainId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _CoreRegistry.Contract.UpdateChainMetadata(&_CoreRegistry.TransactOpts, chainId, key, value)
}

// UpdateContractConfiguration is a paid mutator transaction binding the contract method 0xf354b31f.
//
// Solidity: function updateContractConfiguration(uint256 chainId, string contractType, string key, bytes value) returns()
func (_CoreRegistry *CoreRegistryTransactor) UpdateContractConfiguration(opts *bind.TransactOpts, chainId *big.Int, contractType string, key string, value []byte) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "updateContractConfiguration", chainId, contractType, key, value)
}

// UpdateContractConfiguration is a paid mutator transaction binding the contract method 0xf354b31f.
//
// Solidity: function updateContractConfiguration(uint256 chainId, string contractType, string key, bytes value) returns()
func (_CoreRegistry *CoreRegistrySession) UpdateContractConfiguration(chainId *big.Int, contractType string, key string, value []byte) (*types.Transaction, error) {
	return _CoreRegistry.Contract.UpdateContractConfiguration(&_CoreRegistry.TransactOpts, chainId, contractType, key, value)
}

// UpdateContractConfiguration is a paid mutator transaction binding the contract method 0xf354b31f.
//
// Solidity: function updateContractConfiguration(uint256 chainId, string contractType, string key, bytes value) returns()
func (_CoreRegistry *CoreRegistryTransactorSession) UpdateContractConfiguration(chainId *big.Int, contractType string, key string, value []byte) (*types.Transaction, error) {
	return _CoreRegistry.Contract.UpdateContractConfiguration(&_CoreRegistry.TransactOpts, chainId, contractType, key, value)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CoreRegistry *CoreRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CoreRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CoreRegistry *CoreRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CoreRegistry.Contract.UpgradeToAndCall(&_CoreRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CoreRegistry *CoreRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CoreRegistry.Contract.UpgradeToAndCall(&_CoreRegistry.TransactOpts, newImplementation, data)
}

// CoreRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the CoreRegistry contract.
type CoreRegistryAdminChangedIterator struct {
	Event *CoreRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *CoreRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryAdminChanged)
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
		it.Event = new(CoreRegistryAdminChanged)
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
func (it *CoreRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryAdminChanged represents a AdminChanged event raised by the CoreRegistry contract.
type CoreRegistryAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_CoreRegistry *CoreRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*CoreRegistryAdminChangedIterator, error) {

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryAdminChangedIterator{contract: _CoreRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_CoreRegistry *CoreRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *CoreRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryAdminChanged)
				if err := _CoreRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_CoreRegistry *CoreRegistryFilterer) ParseAdminChanged(log types.Log) (*CoreRegistryAdminChanged, error) {
	event := new(CoreRegistryAdminChanged)
	if err := _CoreRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryChainMetadataUpdatedIterator is returned from FilterChainMetadataUpdated and is used to iterate over the raw logs and unpacked data for ChainMetadataUpdated events raised by the CoreRegistry contract.
type CoreRegistryChainMetadataUpdatedIterator struct {
	Event *CoreRegistryChainMetadataUpdated // Event containing the contract specifics and raw log

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
func (it *CoreRegistryChainMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryChainMetadataUpdated)
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
		it.Event = new(CoreRegistryChainMetadataUpdated)
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
func (it *CoreRegistryChainMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryChainMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryChainMetadataUpdated represents a ChainMetadataUpdated event raised by the CoreRegistry contract.
type CoreRegistryChainMetadataUpdated struct {
	ChainId *big.Int
	Key     string
	Value   []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainMetadataUpdated is a free log retrieval operation binding the contract event 0x40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee634.
//
// Solidity: event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value)
func (_CoreRegistry *CoreRegistryFilterer) FilterChainMetadataUpdated(opts *bind.FilterOpts, chainId []*big.Int) (*CoreRegistryChainMetadataUpdatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "ChainMetadataUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryChainMetadataUpdatedIterator{contract: _CoreRegistry.contract, event: "ChainMetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchChainMetadataUpdated is a free log subscription operation binding the contract event 0x40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee634.
//
// Solidity: event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value)
func (_CoreRegistry *CoreRegistryFilterer) WatchChainMetadataUpdated(opts *bind.WatchOpts, sink chan<- *CoreRegistryChainMetadataUpdated, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "ChainMetadataUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryChainMetadataUpdated)
				if err := _CoreRegistry.contract.UnpackLog(event, "ChainMetadataUpdated", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseChainMetadataUpdated(log types.Log) (*CoreRegistryChainMetadataUpdated, error) {
	event := new(CoreRegistryChainMetadataUpdated)
	if err := _CoreRegistry.contract.UnpackLog(event, "ChainMetadataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryChainStatusChangedIterator is returned from FilterChainStatusChanged and is used to iterate over the raw logs and unpacked data for ChainStatusChanged events raised by the CoreRegistry contract.
type CoreRegistryChainStatusChangedIterator struct {
	Event *CoreRegistryChainStatusChanged // Event containing the contract specifics and raw log

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
func (it *CoreRegistryChainStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryChainStatusChanged)
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
		it.Event = new(CoreRegistryChainStatusChanged)
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
func (it *CoreRegistryChainStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryChainStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryChainStatusChanged represents a ChainStatusChanged event raised by the CoreRegistry contract.
type CoreRegistryChainStatusChanged struct {
	ChainId   *big.Int
	NewStatus bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChainStatusChanged is a free log retrieval operation binding the contract event 0xc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId, bool newStatus)
func (_CoreRegistry *CoreRegistryFilterer) FilterChainStatusChanged(opts *bind.FilterOpts, chainId []*big.Int) (*CoreRegistryChainStatusChangedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "ChainStatusChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryChainStatusChangedIterator{contract: _CoreRegistry.contract, event: "ChainStatusChanged", logs: logs, sub: sub}, nil
}

// WatchChainStatusChanged is a free log subscription operation binding the contract event 0xc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId, bool newStatus)
func (_CoreRegistry *CoreRegistryFilterer) WatchChainStatusChanged(opts *bind.WatchOpts, sink chan<- *CoreRegistryChainStatusChanged, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "ChainStatusChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryChainStatusChanged)
				if err := _CoreRegistry.contract.UnpackLog(event, "ChainStatusChanged", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseChainStatusChanged(log types.Log) (*CoreRegistryChainStatusChanged, error) {
	event := new(CoreRegistryChainStatusChanged)
	if err := _CoreRegistry.contract.UnpackLog(event, "ChainStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryContractConfigurationUpdatedIterator is returned from FilterContractConfigurationUpdated and is used to iterate over the raw logs and unpacked data for ContractConfigurationUpdated events raised by the CoreRegistry contract.
type CoreRegistryContractConfigurationUpdatedIterator struct {
	Event *CoreRegistryContractConfigurationUpdated // Event containing the contract specifics and raw log

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
func (it *CoreRegistryContractConfigurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryContractConfigurationUpdated)
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
		it.Event = new(CoreRegistryContractConfigurationUpdated)
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
func (it *CoreRegistryContractConfigurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryContractConfigurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryContractConfigurationUpdated represents a ContractConfigurationUpdated event raised by the CoreRegistry contract.
type CoreRegistryContractConfigurationUpdated struct {
	ChainId      *big.Int
	ContractType string
	Key          string
	Value        []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractConfigurationUpdated is a free log retrieval operation binding the contract event 0xaea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c5.
//
// Solidity: event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value)
func (_CoreRegistry *CoreRegistryFilterer) FilterContractConfigurationUpdated(opts *bind.FilterOpts, chainId []*big.Int) (*CoreRegistryContractConfigurationUpdatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "ContractConfigurationUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryContractConfigurationUpdatedIterator{contract: _CoreRegistry.contract, event: "ContractConfigurationUpdated", logs: logs, sub: sub}, nil
}

// WatchContractConfigurationUpdated is a free log subscription operation binding the contract event 0xaea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c5.
//
// Solidity: event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value)
func (_CoreRegistry *CoreRegistryFilterer) WatchContractConfigurationUpdated(opts *bind.WatchOpts, sink chan<- *CoreRegistryContractConfigurationUpdated, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "ContractConfigurationUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryContractConfigurationUpdated)
				if err := _CoreRegistry.contract.UnpackLog(event, "ContractConfigurationUpdated", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseContractConfigurationUpdated(log types.Log) (*CoreRegistryContractConfigurationUpdated, error) {
	event := new(CoreRegistryContractConfigurationUpdated)
	if err := _CoreRegistry.contract.UnpackLog(event, "ContractConfigurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryContractRegisteredIterator is returned from FilterContractRegistered and is used to iterate over the raw logs and unpacked data for ContractRegistered events raised by the CoreRegistry contract.
type CoreRegistryContractRegisteredIterator struct {
	Event *CoreRegistryContractRegistered // Event containing the contract specifics and raw log

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
func (it *CoreRegistryContractRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryContractRegistered)
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
		it.Event = new(CoreRegistryContractRegistered)
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
func (it *CoreRegistryContractRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryContractRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryContractRegistered represents a ContractRegistered event raised by the CoreRegistry contract.
type CoreRegistryContractRegistered struct {
	ChainId      *big.Int
	ContractType common.Hash
	AddressBytes []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractRegistered is a free log retrieval operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_CoreRegistry *CoreRegistryFilterer) FilterContractRegistered(opts *bind.FilterOpts, chainId []*big.Int, contractType []string) (*CoreRegistryContractRegisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "ContractRegistered", chainIdRule, contractTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryContractRegisteredIterator{contract: _CoreRegistry.contract, event: "ContractRegistered", logs: logs, sub: sub}, nil
}

// WatchContractRegistered is a free log subscription operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_CoreRegistry *CoreRegistryFilterer) WatchContractRegistered(opts *bind.WatchOpts, sink chan<- *CoreRegistryContractRegistered, chainId []*big.Int, contractType []string) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "ContractRegistered", chainIdRule, contractTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryContractRegistered)
				if err := _CoreRegistry.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseContractRegistered(log types.Log) (*CoreRegistryContractRegistered, error) {
	event := new(CoreRegistryContractRegistered)
	if err := _CoreRegistry.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryContractStatusChangedIterator is returned from FilterContractStatusChanged and is used to iterate over the raw logs and unpacked data for ContractStatusChanged events raised by the CoreRegistry contract.
type CoreRegistryContractStatusChangedIterator struct {
	Event *CoreRegistryContractStatusChanged // Event containing the contract specifics and raw log

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
func (it *CoreRegistryContractStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryContractStatusChanged)
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
		it.Event = new(CoreRegistryContractStatusChanged)
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
func (it *CoreRegistryContractStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryContractStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryContractStatusChanged represents a ContractStatusChanged event raised by the CoreRegistry contract.
type CoreRegistryContractStatusChanged struct {
	AddressBytes []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractStatusChanged is a free log retrieval operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_CoreRegistry *CoreRegistryFilterer) FilterContractStatusChanged(opts *bind.FilterOpts) (*CoreRegistryContractStatusChangedIterator, error) {

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "ContractStatusChanged")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryContractStatusChangedIterator{contract: _CoreRegistry.contract, event: "ContractStatusChanged", logs: logs, sub: sub}, nil
}

// WatchContractStatusChanged is a free log subscription operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_CoreRegistry *CoreRegistryFilterer) WatchContractStatusChanged(opts *bind.WatchOpts, sink chan<- *CoreRegistryContractStatusChanged) (event.Subscription, error) {

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "ContractStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryContractStatusChanged)
				if err := _CoreRegistry.contract.UnpackLog(event, "ContractStatusChanged", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseContractStatusChanged(log types.Log) (*CoreRegistryContractStatusChanged, error) {
	event := new(CoreRegistryContractStatusChanged)
	if err := _CoreRegistry.contract.UnpackLog(event, "ContractStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CoreRegistry contract.
type CoreRegistryInitializedIterator struct {
	Event *CoreRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *CoreRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryInitialized)
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
		it.Event = new(CoreRegistryInitialized)
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
func (it *CoreRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryInitialized represents a Initialized event raised by the CoreRegistry contract.
type CoreRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CoreRegistry *CoreRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*CoreRegistryInitializedIterator, error) {

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryInitializedIterator{contract: _CoreRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CoreRegistry *CoreRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CoreRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryInitialized)
				if err := _CoreRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseInitialized(log types.Log) (*CoreRegistryInitialized, error) {
	event := new(CoreRegistryInitialized)
	if err := _CoreRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CoreRegistry contract.
type CoreRegistryPausedIterator struct {
	Event *CoreRegistryPaused // Event containing the contract specifics and raw log

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
func (it *CoreRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryPaused)
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
		it.Event = new(CoreRegistryPaused)
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
func (it *CoreRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryPaused represents a Paused event raised by the CoreRegistry contract.
type CoreRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CoreRegistry *CoreRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*CoreRegistryPausedIterator, error) {

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryPausedIterator{contract: _CoreRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CoreRegistry *CoreRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CoreRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryPaused)
				if err := _CoreRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParsePaused(log types.Log) (*CoreRegistryPaused, error) {
	event := new(CoreRegistryPaused)
	if err := _CoreRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryRegistryManagerChangedIterator is returned from FilterRegistryManagerChanged and is used to iterate over the raw logs and unpacked data for RegistryManagerChanged events raised by the CoreRegistry contract.
type CoreRegistryRegistryManagerChangedIterator struct {
	Event *CoreRegistryRegistryManagerChanged // Event containing the contract specifics and raw log

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
func (it *CoreRegistryRegistryManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryRegistryManagerChanged)
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
		it.Event = new(CoreRegistryRegistryManagerChanged)
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
func (it *CoreRegistryRegistryManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryRegistryManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryRegistryManagerChanged represents a RegistryManagerChanged event raised by the CoreRegistry contract.
type CoreRegistryRegistryManagerChanged struct {
	OldRegistryManager common.Address
	NewRegistryManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRegistryManagerChanged is a free log retrieval operation binding the contract event 0x6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279.
//
// Solidity: event RegistryManagerChanged(address oldRegistryManager, address newRegistryManager)
func (_CoreRegistry *CoreRegistryFilterer) FilterRegistryManagerChanged(opts *bind.FilterOpts) (*CoreRegistryRegistryManagerChangedIterator, error) {

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "RegistryManagerChanged")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryRegistryManagerChangedIterator{contract: _CoreRegistry.contract, event: "RegistryManagerChanged", logs: logs, sub: sub}, nil
}

// WatchRegistryManagerChanged is a free log subscription operation binding the contract event 0x6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279.
//
// Solidity: event RegistryManagerChanged(address oldRegistryManager, address newRegistryManager)
func (_CoreRegistry *CoreRegistryFilterer) WatchRegistryManagerChanged(opts *bind.WatchOpts, sink chan<- *CoreRegistryRegistryManagerChanged) (event.Subscription, error) {

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "RegistryManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryRegistryManagerChanged)
				if err := _CoreRegistry.contract.UnpackLog(event, "RegistryManagerChanged", log); err != nil {
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

// ParseRegistryManagerChanged is a log parse operation binding the contract event 0x6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279.
//
// Solidity: event RegistryManagerChanged(address oldRegistryManager, address newRegistryManager)
func (_CoreRegistry *CoreRegistryFilterer) ParseRegistryManagerChanged(log types.Log) (*CoreRegistryRegistryManagerChanged, error) {
	event := new(CoreRegistryRegistryManagerChanged)
	if err := _CoreRegistry.contract.UnpackLog(event, "RegistryManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CoreRegistry contract.
type CoreRegistryRoleAdminChangedIterator struct {
	Event *CoreRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CoreRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryRoleAdminChanged)
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
		it.Event = new(CoreRegistryRoleAdminChanged)
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
func (it *CoreRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the CoreRegistry contract.
type CoreRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CoreRegistry *CoreRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CoreRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryRoleAdminChangedIterator{contract: _CoreRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CoreRegistry *CoreRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CoreRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryRoleAdminChanged)
				if err := _CoreRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*CoreRegistryRoleAdminChanged, error) {
	event := new(CoreRegistryRoleAdminChanged)
	if err := _CoreRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CoreRegistry contract.
type CoreRegistryRoleGrantedIterator struct {
	Event *CoreRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *CoreRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryRoleGranted)
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
		it.Event = new(CoreRegistryRoleGranted)
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
func (it *CoreRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryRoleGranted represents a RoleGranted event raised by the CoreRegistry contract.
type CoreRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CoreRegistry *CoreRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CoreRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryRoleGrantedIterator{contract: _CoreRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CoreRegistry *CoreRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CoreRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryRoleGranted)
				if err := _CoreRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseRoleGranted(log types.Log) (*CoreRegistryRoleGranted, error) {
	event := new(CoreRegistryRoleGranted)
	if err := _CoreRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CoreRegistry contract.
type CoreRegistryRoleRevokedIterator struct {
	Event *CoreRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CoreRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryRoleRevoked)
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
		it.Event = new(CoreRegistryRoleRevoked)
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
func (it *CoreRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryRoleRevoked represents a RoleRevoked event raised by the CoreRegistry contract.
type CoreRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CoreRegistry *CoreRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CoreRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryRoleRevokedIterator{contract: _CoreRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CoreRegistry *CoreRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CoreRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryRoleRevoked)
				if err := _CoreRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseRoleRevoked(log types.Log) (*CoreRegistryRoleRevoked, error) {
	event := new(CoreRegistryRoleRevoked)
	if err := _CoreRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CoreRegistry contract.
type CoreRegistryUnpausedIterator struct {
	Event *CoreRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *CoreRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryUnpaused)
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
		it.Event = new(CoreRegistryUnpaused)
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
func (it *CoreRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryUnpaused represents a Unpaused event raised by the CoreRegistry contract.
type CoreRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CoreRegistry *CoreRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CoreRegistryUnpausedIterator, error) {

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryUnpausedIterator{contract: _CoreRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CoreRegistry *CoreRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CoreRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryUnpaused)
				if err := _CoreRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseUnpaused(log types.Log) (*CoreRegistryUnpaused, error) {
	event := new(CoreRegistryUnpaused)
	if err := _CoreRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CoreRegistry contract.
type CoreRegistryUpgradedIterator struct {
	Event *CoreRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *CoreRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryUpgraded)
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
		it.Event = new(CoreRegistryUpgraded)
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
func (it *CoreRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryUpgraded represents a Upgraded event raised by the CoreRegistry contract.
type CoreRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CoreRegistry *CoreRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CoreRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryUpgradedIterator{contract: _CoreRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CoreRegistry *CoreRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CoreRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryUpgraded)
				if err := _CoreRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseUpgraded(log types.Log) (*CoreRegistryUpgraded, error) {
	event := new(CoreRegistryUpgraded)
	if err := _CoreRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryZRC20TokenRegisteredIterator is returned from FilterZRC20TokenRegistered and is used to iterate over the raw logs and unpacked data for ZRC20TokenRegistered events raised by the CoreRegistry contract.
type CoreRegistryZRC20TokenRegisteredIterator struct {
	Event *CoreRegistryZRC20TokenRegistered // Event containing the contract specifics and raw log

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
func (it *CoreRegistryZRC20TokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryZRC20TokenRegistered)
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
		it.Event = new(CoreRegistryZRC20TokenRegistered)
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
func (it *CoreRegistryZRC20TokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryZRC20TokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryZRC20TokenRegistered represents a ZRC20TokenRegistered event raised by the CoreRegistry contract.
type CoreRegistryZRC20TokenRegistered struct {
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
func (_CoreRegistry *CoreRegistryFilterer) FilterZRC20TokenRegistered(opts *bind.FilterOpts, originAddress [][]byte, address_ []common.Address) (*CoreRegistryZRC20TokenRegisteredIterator, error) {

	var originAddressRule []interface{}
	for _, originAddressItem := range originAddress {
		originAddressRule = append(originAddressRule, originAddressItem)
	}
	var address_Rule []interface{}
	for _, address_Item := range address_ {
		address_Rule = append(address_Rule, address_Item)
	}

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "ZRC20TokenRegistered", originAddressRule, address_Rule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryZRC20TokenRegisteredIterator{contract: _CoreRegistry.contract, event: "ZRC20TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchZRC20TokenRegistered is a free log subscription operation binding the contract event 0xa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367.
//
// Solidity: event ZRC20TokenRegistered(bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol)
func (_CoreRegistry *CoreRegistryFilterer) WatchZRC20TokenRegistered(opts *bind.WatchOpts, sink chan<- *CoreRegistryZRC20TokenRegistered, originAddress [][]byte, address_ []common.Address) (event.Subscription, error) {

	var originAddressRule []interface{}
	for _, originAddressItem := range originAddress {
		originAddressRule = append(originAddressRule, originAddressItem)
	}
	var address_Rule []interface{}
	for _, address_Item := range address_ {
		address_Rule = append(address_Rule, address_Item)
	}

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "ZRC20TokenRegistered", originAddressRule, address_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryZRC20TokenRegistered)
				if err := _CoreRegistry.contract.UnpackLog(event, "ZRC20TokenRegistered", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseZRC20TokenRegistered(log types.Log) (*CoreRegistryZRC20TokenRegistered, error) {
	event := new(CoreRegistryZRC20TokenRegistered)
	if err := _CoreRegistry.contract.UnpackLog(event, "ZRC20TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryZRC20TokenUpdatedIterator is returned from FilterZRC20TokenUpdated and is used to iterate over the raw logs and unpacked data for ZRC20TokenUpdated events raised by the CoreRegistry contract.
type CoreRegistryZRC20TokenUpdatedIterator struct {
	Event *CoreRegistryZRC20TokenUpdated // Event containing the contract specifics and raw log

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
func (it *CoreRegistryZRC20TokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryZRC20TokenUpdated)
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
		it.Event = new(CoreRegistryZRC20TokenUpdated)
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
func (it *CoreRegistryZRC20TokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryZRC20TokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryZRC20TokenUpdated represents a ZRC20TokenUpdated event raised by the CoreRegistry contract.
type CoreRegistryZRC20TokenUpdated struct {
	Address common.Address
	Active  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterZRC20TokenUpdated is a free log retrieval operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_CoreRegistry *CoreRegistryFilterer) FilterZRC20TokenUpdated(opts *bind.FilterOpts) (*CoreRegistryZRC20TokenUpdatedIterator, error) {

	logs, sub, err := _CoreRegistry.contract.FilterLogs(opts, "ZRC20TokenUpdated")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryZRC20TokenUpdatedIterator{contract: _CoreRegistry.contract, event: "ZRC20TokenUpdated", logs: logs, sub: sub}, nil
}

// WatchZRC20TokenUpdated is a free log subscription operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_CoreRegistry *CoreRegistryFilterer) WatchZRC20TokenUpdated(opts *bind.WatchOpts, sink chan<- *CoreRegistryZRC20TokenUpdated) (event.Subscription, error) {

	logs, sub, err := _CoreRegistry.contract.WatchLogs(opts, "ZRC20TokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryZRC20TokenUpdated)
				if err := _CoreRegistry.contract.UnpackLog(event, "ZRC20TokenUpdated", log); err != nil {
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
func (_CoreRegistry *CoreRegistryFilterer) ParseZRC20TokenUpdated(log types.Log) (*CoreRegistryZRC20TokenUpdated, error) {
	event := new(CoreRegistryZRC20TokenUpdated)
	if err := _CoreRegistry.contract.UnpackLog(event, "ZRC20TokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
