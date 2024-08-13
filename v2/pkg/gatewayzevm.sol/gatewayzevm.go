// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayzevm

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

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Asset         common.Address
	Amount        uint64
	RevertMessage []byte
}

// RevertOptions is an auto generated low-level Go binding around an user-defined struct.
type RevertOptions struct {
	RevertAddress common.Address
	CallOnRevert  bool
	AbortAddress  common.Address
	RevertMessage []byte
}

// ZContext is an auto generated low-level Go binding around an user-defined struct.
type ZContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// GatewayZEVMMetaData contains all meta data concerning the GatewayZEVM contract.
var GatewayZEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structzContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structzContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndRevert\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structzContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotFungibleModule\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrFungible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60805161391b6100fd6000396000818161230a015281816123330152612509015261391b6000f3fe6080604052600436106101bb5760003560e01c8063717d335f116100ec578063ad3cb1cc1161008a578063d547741f11610064578063d547741f14610600578063dc9ca2e714610620578063e63ab1e914610640578063f45346dc1461067457600080fd5b8063ad3cb1cc1461056a578063bcf7f32b146105c0578063c39aca37146105e057600080fd5b806390ad3e23116100c657806390ad3e23146104b057806391d14854146104d05780639ec6048414610535578063a217fddf1461055557600080fd5b8063717d335f1461045b57806371f6af541461047b5780638456cb591461049b57600080fd5b806336568abe11610159578063485cc95511610133578063485cc955146103dc5780634f1ef286146103fc57806352d1902d1461040f5780635c975abb1461042457600080fd5b806336568abe1461037f5780633ce4a5bc1461039f5780633f4ba83a146103c757600080fd5b806321501a951161019557806321501a95146102aa57806321e093b1146102ca578063248a9ca3146103025780632f2ff15d1461035f57600080fd5b806301ffc9a7146102355780630310eb761461026a5780631e83da291461028a57600080fd5b36610230576101c8610694565b6000546001600160a01b031633148015906101f757503373735b14bb79463307aacbed86daf3322b1e6226ab14155b1561022e576040517f229930b200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b34801561024157600080fd5b50610255610250366004612d2b565b6106f2565b60405190151581526020015b60405180910390f35b34801561027657600080fd5b5061022e610285366004612d9a565b61078b565b34801561029657600080fd5b5061022e6102a5366004612f5c565b610a17565b3480156102b657600080fd5b5061022e6102c5366004613008565b610b42565b3480156102d657600080fd5b506000546102ea906001600160a01b031681565b6040516001600160a01b039091168152602001610261565b34801561030e57600080fd5b5061035161031d366004613094565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b604051908152602001610261565b34801561036b57600080fd5b5061022e61037a3660046130ad565b610d0f565b34801561038b57600080fd5b5061022e61039a3660046130ad565b610d59565b3480156103ab57600080fd5b506102ea73735b14bb79463307aacbed86daf3322b1e6226ab81565b3480156103d357600080fd5b5061022e610daa565b3480156103e857600080fd5b5061022e6103f73660046130dd565b610ddf565b61022e61040a36600461310b565b611035565b34801561041b57600080fd5b50610351611054565b34801561043057600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610255565b34801561046757600080fd5b5061022e61047636600461315b565b611083565b34801561048757600080fd5b5061022e6104963660046131a1565b61118b565b3480156104a757600080fd5b5061022e611359565b3480156104bc57600080fd5b5061022e6104cb36600461321a565b61138b565b3480156104dc57600080fd5b506102556104eb3660046130ad565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561054157600080fd5b5061022e6105503660046132da565b6114fd565b34801561056157600080fd5b50610351600081565b34801561057657600080fd5b506105b36040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161026191906133ac565b3480156105cc57600080fd5b5061022e6105db3660046133bf565b61161c565b3480156105ec57600080fd5b5061022e6105fb3660046133bf565b61174c565b34801561060c57600080fd5b5061022e61061b3660046130ad565b61195b565b34801561062c57600080fd5b5061022e61063b36600461345d565b61199f565b34801561064c57600080fd5b506103517f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b34801561068057600080fd5b5061022e61068f3660046134f9565b611c08565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16156106f0576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061078557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107d8576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107e0610694565b6001600160a01b03841615806107fd57506001600160a01b038216155b15610834576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8260000361086e576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03821673735b14bb79463307aacbed86daf3322b1e6226ab14806108a157506001600160a01b03821630145b156108d8576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152602482018590528516906347e7ef24906044016020604051808303816000875af1158015610940573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109649190613549565b61099a576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f660b9de00000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063660b9de0906109df908490600401613613565b600060405180830381600087803b1580156109f957600080fd5b505af1158015610a0d573d6000803e3d6000fd5b5050505050505050565b610a1f611e17565b610a27610694565b8551600003610a62576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84600003610a9c576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610aba8573735b14bb79463307aacbed86daf3322b1e6226ab611e98565b60008054604051869233927ff6e8f2c19d797cd515ed61d8ac909afa628f4df0bb153a3f81d0f5fb7b8207c092610b09928c926001600160a01b0316918c919081908c908c9083908d906136e4565b60405180910390a3610b3a60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610b8f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b97610694565b6001600160a01b038316610bd7576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600003610c11576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab1480610c4457506001600160a01b03831630145b15610c7b576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c858484611e98565b6000546040517fde43156e0000000000000000000000000000000000000000000000000000000081526001600160a01b038086169263de43156e92610cd6928a921690899088908890600401613757565b600060405180830381600087803b158015610cf057600080fd5b505af1158015610d04573d6000803e3d6000fd5b505050505050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610d498161208c565b610d538383612096565b50505050565b6001600160a01b0381163314610d9b576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610da58282612183565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610dd48161208c565b610ddc612247565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610e2a5750825b905060008267ffffffffffffffff166001148015610e475750303b155b905081158015610e55575080155b15610e8c576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610eed5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b0387161580610f0a57506001600160a01b038616155b15610f41576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f496122d7565b610f516122d7565b610f596122df565b610f616122ef565b610f6c600087612096565b50610f977f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87612096565b50600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038916179055831561102c5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b61103d6122ff565b611046826123cf565b61105082826123da565b5050565b600061105e6124fe565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b3373735b14bb79463307aacbed86daf3322b1e6226ab146110d0576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110d8610694565b6001600160a01b038216611118576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f660b9de00000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063660b9de09061115d908490600401613613565b600060405180830381600087803b15801561117757600080fd5b505af1158015610b3a573d6000803e3d6000fd5b611193611e17565b61119b610694565b83516000036111d6576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82600003611210576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061121c8484612560565b90506000336001600160a01b03167ff6e8f2c19d797cd515ed61d8ac909afa628f4df0bb153a3f81d0f5fb7b8207c087868886896001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561128d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112b191906137d4565b8a6001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112ef573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061131391906137d4565b8a60405161132797969594939291906137ed565b60405180910390a350610d5360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6113838161208c565b610ddc6125d5565b611393611e17565b61139b610694565b86516000036113d6576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b85600003611410576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061141d87878561264e565b90506000336001600160a01b03167ff6e8f2c19d797cd515ed61d8ac909afa628f4df0bb153a3f81d0f5fb7b8207c08a898b868c6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561148e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b291906137d4565b8c8c8c8c6040516114cb999897969594939291906136e4565b60405180910390a35061102c60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b611505611e17565b61150d610694565b8351600003611548576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82600003611582576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115a08373735b14bb79463307aacbed86daf3322b1e6226ab611e98565b60008054604051849233927ff6e8f2c19d797cd515ed61d8ac909afa628f4df0bb153a3f81d0f5fb7b8207c0926115eb928a926001600160a01b0316918a9190819081908b906137ed565b60405180910390a3610d5360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611669576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611671610694565b6001600160a01b038516158061168e57506001600160a01b038316155b156116c5576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fde43156e0000000000000000000000000000000000000000000000000000000081526001600160a01b0384169063de43156e906117129089908990899088908890600401613757565b600060405180830381600087803b15801561172c57600080fd5b505af1158015611740573d6000803e3d6000fd5b50505050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611799576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117a1610694565b6001600160a01b03851615806117be57506001600160a01b038316155b156117f5576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8360000361182f576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab148061186257506001600160a01b03831630145b15611899576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038481166004830152602482018690528616906347e7ef24906044016020604051808303816000875af1158015611901573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119259190613549565b6116c5576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546119958161208c565b610d538383612183565b6119a7611e17565b6119af610694565b85516000036119ea576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000839003611a25576040517fbb1b658000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517ffc5fecd50000000000000000000000000000000000000000000000000000000081526004810183905260009081906001600160a01b0388169063fc5fecd5906024016040805180830381865afa158015611a87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aab919061385b565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab60248201526044810182905291935091506001600160a01b038316906323b872dd906064016020604051808303816000875af1158015611b30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b549190613549565b611b8a576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b866001600160a01b0316336001600160a01b03167f885ccd61b2cc8be1daa203d4456750151d128d2da2ace4c9fa4b63c395fd99d88a89898989604051611bd5959493929190613889565b60405180910390a35050610b3a60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611c55576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c5d610694565b6001600160a01b0383161580611c7a57506001600160a01b038116155b15611cb1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600003611ceb576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03811673735b14bb79463307aacbed86daf3322b1e6226ab1480611d1e57506001600160a01b03811630145b15611d55576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038281166004830152602482018490528416906347e7ef24906044016020604051808303816000875af1158015611dbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611de19190613549565b610da5576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611e92576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6000546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490526001600160a01b03909116906323b872dd906064016020604051808303816000875af1158015611f08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f2c9190613549565b611f62576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000546040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b0390911690632e1a7d4d90602401600060405180830381600087803b158015611fc157600080fd5b505af1158015611fd5573d6000803e3d6000fd5b505050506000816001600160a01b03168360405160006040518083038185875af1925050503d8060008114612026576040519150601f19603f3d011682016040523d82523d6000602084013e61202b565b606091505b5050905080610da5576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b610ddc8133612929565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612179576000848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561212f3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610785565b6000915050610785565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612179576000848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610785565b61224f6129b6565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b6106f0612a11565b6122e7612a11565b6106f0612a78565b6122f7612a11565b6106f0612ac9565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061239857507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661238c7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b156106f0576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006110508161208c565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612452575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261244f918101906137d4565b60015b612498576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146124f4576040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004810182905260240161248f565b610da58383612ad1565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106f0576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006125ce8383846001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa1580156125a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125c991906137d4565b61264e565b9392505050565b6125dd610694565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336122b9565b6000806000846001600160a01b031663fc5fecd5856040518263ffffffff1660e01b815260040161268191815260200190565b6040805180830381865afa15801561269d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126c1919061385b565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab60248201526044810182905291935091506001600160a01b038316906323b872dd906064016020604051808303816000875af1158015612746573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061276a9190613549565b6127a0576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018790526001600160a01b038616906323b872dd906064016020604051808303816000875af115801561280c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128309190613549565b612866576040517f4dd9ee8d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f42966c68000000000000000000000000000000000000000000000000000000008152600481018790526001600160a01b038616906342966c68906024016020604051808303816000875af11580156128c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128ea9190613549565b612920576040517f2c77e05c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b95945050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16611050576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024810183905260440161248f565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166106f0576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff166106f0576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612a80612a11565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b612066612a11565b612ada82612b27565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115612b1f57610da58282612bcf565b611050612c3c565b806001600160a01b03163b600003612b76576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161248f565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051612bec91906138c9565b600060405180830381855af49150503d8060008114612c27576040519150601f19603f3d011682016040523d82523d6000602084013e612c2c565b606091505b5091509150612920858383612c74565b34156106f0576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082612c8957612c8482612ce9565b6125ce565b8151158015612ca057506001600160a01b0384163b155b15612ce2576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161248f565b50806125ce565b805115612cf95780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215612d3d57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146125ce57600080fd5b6001600160a01b0381168114610ddc57600080fd5b600060608284031215612d9457600080fd5b50919050565b60008060008060808587031215612db057600080fd5b8435612dbb81612d6d565b9350602085013592506040850135612dd281612d6d565b9150606085013567ffffffffffffffff811115612dee57600080fd5b612dfa87828801612d82565b91505092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112612e4657600080fd5b813567ffffffffffffffff811115612e6057612e60612e06565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715612ecc57612ecc612e06565b604052818152838201602001851015612ee457600080fd5b816020850160208301376000918101602001919091529392505050565b60008083601f840112612f1357600080fd5b50813567ffffffffffffffff811115612f2b57600080fd5b602083019150836020828501011115612f4357600080fd5b9250929050565b600060808284031215612d9457600080fd5b60008060008060008060a08789031215612f7557600080fd5b863567ffffffffffffffff811115612f8c57600080fd5b612f9889828a01612e35565b9650506020870135945060408701359350606087013567ffffffffffffffff811115612fc357600080fd5b612fcf89828a01612f01565b909450925050608087013567ffffffffffffffff811115612fef57600080fd5b612ffb89828a01612f4a565b9150509295509295509295565b60008060008060006080868803121561302057600080fd5b853567ffffffffffffffff81111561303757600080fd5b61304388828901612d82565b95505060208601359350604086013561305b81612d6d565b9250606086013567ffffffffffffffff81111561307757600080fd5b61308388828901612f01565b969995985093965092949392505050565b6000602082840312156130a657600080fd5b5035919050565b600080604083850312156130c057600080fd5b8235915060208301356130d281612d6d565b809150509250929050565b600080604083850312156130f057600080fd5b82356130fb81612d6d565b915060208301356130d281612d6d565b6000806040838503121561311e57600080fd5b823561312981612d6d565b9150602083013567ffffffffffffffff81111561314557600080fd5b61315185828601612e35565b9150509250929050565b6000806040838503121561316e57600080fd5b823561317981612d6d565b9150602083013567ffffffffffffffff81111561319557600080fd5b61315185828601612d82565b600080600080608085870312156131b757600080fd5b843567ffffffffffffffff8111156131ce57600080fd5b6131da87828801612e35565b9450506020850135925060408501356131f281612d6d565b9150606085013567ffffffffffffffff81111561320e57600080fd5b612dfa87828801612f4a565b600080600080600080600060c0888a03121561323557600080fd5b873567ffffffffffffffff81111561324c57600080fd5b6132588a828b01612e35565b97505060208801359550604088013561327081612d6d565b9450606088013567ffffffffffffffff81111561328c57600080fd5b6132988a828b01612f01565b9095509350506080880135915060a088013567ffffffffffffffff8111156132bf57600080fd5b6132cb8a828b01612f4a565b91505092959891949750929550565b600080600080608085870312156132f057600080fd5b843567ffffffffffffffff81111561330757600080fd5b61331387828801612e35565b9450506020850135925060408501359150606085013567ffffffffffffffff81111561320e57600080fd5b60005b83811015613359578181015183820152602001613341565b50506000910152565b6000815180845261337a81602086016020860161333e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006125ce6020830184613362565b60008060008060008060a087890312156133d857600080fd5b863567ffffffffffffffff8111156133ef57600080fd5b6133fb89828a01612d82565b965050602087013561340c81612d6d565b945060408701359350606087013561342381612d6d565b9250608087013567ffffffffffffffff81111561343f57600080fd5b61344b89828a01612f01565b979a9699509497509295939492505050565b60008060008060008060a0878903121561347657600080fd5b863567ffffffffffffffff81111561348d57600080fd5b61349989828a01612e35565b96505060208701356134aa81612d6d565b9450604087013567ffffffffffffffff8111156134c657600080fd5b6134d289828a01612f01565b90955093505060608701359150608087013567ffffffffffffffff811115612fef57600080fd5b60008060006060848603121561350e57600080fd5b833561351981612d6d565b925060208401359150604084013561353081612d6d565b809150509250925092565b8015158114610ddc57600080fd5b60006020828403121561355b57600080fd5b81516125ce8161353b565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261359b57600080fd5b830160208101925035905067ffffffffffffffff8111156135bb57600080fd5b803603821315612f4357600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b602081526000823561362481612d6d565b6001600160a01b038116602084015250602083013567ffffffffffffffff811680821461365057600080fd5b80604085015250506136656040840184613566565b6060808501526129206080850182846135ca565b6000813561368681612d6d565b6001600160a01b03168352602082013561369f8161353b565b1515602084015260408201356136b481612d6d565b6001600160a01b031660408401526136cf6060830183613566565b608060608601526129206080860182846135ca565b610100815260006136f961010083018c613362565b6001600160a01b038b16602084015289604084015288606084015287608084015282810360a084015261372d8187896135ca565b90508460c084015282810360e08401526137478185613679565b9c9b505050505050505050505050565b6080815260006137678788613566565b6060608085015261377c60e0850182846135ca565b915050602088013561378d81612d6d565b6001600160a01b0390811660a085015260408981013560c08601529088166020850152830186905282810360608401526137c88185876135ca565b98975050505050505050565b6000602082840312156137e657600080fd5b5051919050565b6101008152600061380261010083018a613362565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a0850152600082528560c08501526020810160e08501525061384d6020820185613679565b9a9950505050505050505050565b6000806040838503121561386e57600080fd5b825161387981612d6d565b6020939093015192949293505050565b60808152600061389c6080830188613362565b82810360208401526138af8187896135ca565b905084604084015282810360608401526137c88185613679565b600082516138db81846020870161333e565b919091019291505056fea264697066735822122079b4048b9e6e573d5d1463264f1ba2207e660695af25c98b4b72aeea4df7a67064736f6c634300081a0033",
}

// GatewayZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayZEVMMetaData.ABI instead.
var GatewayZEVMABI = GatewayZEVMMetaData.ABI

// GatewayZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayZEVMMetaData.Bin instead.
var GatewayZEVMBin = GatewayZEVMMetaData.Bin

// DeployGatewayZEVM deploys a new Ethereum contract, binding an instance of GatewayZEVM to it.
func DeployGatewayZEVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayZEVM, error) {
	parsed, err := GatewayZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayZEVMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayZEVM{GatewayZEVMCaller: GatewayZEVMCaller{contract: contract}, GatewayZEVMTransactor: GatewayZEVMTransactor{contract: contract}, GatewayZEVMFilterer: GatewayZEVMFilterer{contract: contract}}, nil
}

// GatewayZEVM is an auto generated Go binding around an Ethereum contract.
type GatewayZEVM struct {
	GatewayZEVMCaller     // Read-only binding to the contract
	GatewayZEVMTransactor // Write-only binding to the contract
	GatewayZEVMFilterer   // Log filterer for contract events
}

// GatewayZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayZEVMSession struct {
	Contract     *GatewayZEVM      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayZEVMCallerSession struct {
	Contract *GatewayZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GatewayZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayZEVMTransactorSession struct {
	Contract     *GatewayZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayZEVMRaw struct {
	Contract *GatewayZEVM // Generic contract binding to access the raw methods on
}

// GatewayZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayZEVMCallerRaw struct {
	Contract *GatewayZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayZEVMTransactorRaw struct {
	Contract *GatewayZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayZEVM creates a new instance of GatewayZEVM, bound to a specific deployed contract.
func NewGatewayZEVM(address common.Address, backend bind.ContractBackend) (*GatewayZEVM, error) {
	contract, err := bindGatewayZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVM{GatewayZEVMCaller: GatewayZEVMCaller{contract: contract}, GatewayZEVMTransactor: GatewayZEVMTransactor{contract: contract}, GatewayZEVMFilterer: GatewayZEVMFilterer{contract: contract}}, nil
}

// NewGatewayZEVMCaller creates a new read-only instance of GatewayZEVM, bound to a specific deployed contract.
func NewGatewayZEVMCaller(address common.Address, caller bind.ContractCaller) (*GatewayZEVMCaller, error) {
	contract, err := bindGatewayZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMCaller{contract: contract}, nil
}

// NewGatewayZEVMTransactor creates a new write-only instance of GatewayZEVM, bound to a specific deployed contract.
func NewGatewayZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayZEVMTransactor, error) {
	contract, err := bindGatewayZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMTransactor{contract: contract}, nil
}

// NewGatewayZEVMFilterer creates a new log filterer instance of GatewayZEVM, bound to a specific deployed contract.
func NewGatewayZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayZEVMFilterer, error) {
	contract, err := bindGatewayZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMFilterer{contract: contract}, nil
}

// bindGatewayZEVM binds a generic wrapper to an already deployed contract.
func bindGatewayZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVM *GatewayZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVM.Contract.GatewayZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVM *GatewayZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.GatewayZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVM *GatewayZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.GatewayZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVM *GatewayZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVM *GatewayZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVM *GatewayZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayZEVM.Contract.DEFAULTADMINROLE(&_GatewayZEVM.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayZEVM.Contract.DEFAULTADMINROLE(&_GatewayZEVM.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_GatewayZEVM *GatewayZEVMCaller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_GatewayZEVM *GatewayZEVMSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _GatewayZEVM.Contract.FUNGIBLEMODULEADDRESS(&_GatewayZEVM.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_GatewayZEVM *GatewayZEVMCallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _GatewayZEVM.Contract.FUNGIBLEMODULEADDRESS(&_GatewayZEVM.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVM.Contract.PAUSERROLE(&_GatewayZEVM.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVM.Contract.PAUSERROLE(&_GatewayZEVM.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayZEVM *GatewayZEVMCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayZEVM *GatewayZEVMSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayZEVM.Contract.UPGRADEINTERFACEVERSION(&_GatewayZEVM.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayZEVM *GatewayZEVMCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayZEVM.Contract.UPGRADEINTERFACEVERSION(&_GatewayZEVM.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayZEVM.Contract.GetRoleAdmin(&_GatewayZEVM.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayZEVM.Contract.GetRoleAdmin(&_GatewayZEVM.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayZEVM *GatewayZEVMCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayZEVM *GatewayZEVMSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayZEVM.Contract.HasRole(&_GatewayZEVM.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayZEVM *GatewayZEVMCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayZEVM.Contract.HasRole(&_GatewayZEVM.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayZEVM *GatewayZEVMCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayZEVM *GatewayZEVMSession) Paused() (bool, error) {
	return _GatewayZEVM.Contract.Paused(&_GatewayZEVM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayZEVM *GatewayZEVMCallerSession) Paused() (bool, error) {
	return _GatewayZEVM.Contract.Paused(&_GatewayZEVM.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayZEVM.Contract.ProxiableUUID(&_GatewayZEVM.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayZEVM.Contract.ProxiableUUID(&_GatewayZEVM.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayZEVM *GatewayZEVMCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayZEVM *GatewayZEVMSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayZEVM.Contract.SupportsInterface(&_GatewayZEVM.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayZEVM *GatewayZEVMCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayZEVM.Contract.SupportsInterface(&_GatewayZEVM.CallOpts, interfaceId)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayZEVM *GatewayZEVMCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayZEVM *GatewayZEVMSession) ZetaToken() (common.Address, error) {
	return _GatewayZEVM.Contract.ZetaToken(&_GatewayZEVM.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayZEVM *GatewayZEVMCallerSession) ZetaToken() (common.Address, error) {
	return _GatewayZEVM.Contract.ZetaToken(&_GatewayZEVM.CallOpts)
}

// Call is a paid mutator transaction binding the contract method 0xdc9ca2e7.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Call(opts *bind.TransactOpts, receiver []byte, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "call", receiver, zrc20, message, gasLimit, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0xdc9ca2e7.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) Call(receiver []byte, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Call(&_GatewayZEVM.TransactOpts, receiver, zrc20, message, gasLimit, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0xdc9ca2e7.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Call(receiver []byte, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Call(&_GatewayZEVM.TransactOpts, receiver, zrc20, message, gasLimit, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Deposit(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "deposit", zrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVM *GatewayZEVMSession) Deposit(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Deposit(&_GatewayZEVM.TransactOpts, zrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Deposit(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Deposit(&_GatewayZEVM.TransactOpts, zrc20, amount, target)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) DepositAndCall(opts *bind.TransactOpts, context ZContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "depositAndCall", context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMSession) DepositAndCall(context ZContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndCall(&_GatewayZEVM.TransactOpts, context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) DepositAndCall(context ZContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndCall(&_GatewayZEVM.TransactOpts, context, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) DepositAndCall0(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "depositAndCall0", context, zrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMSession) DepositAndCall0(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndCall0(&_GatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) DepositAndCall0(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndCall0(&_GatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x0310eb76.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,uint64,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) DepositAndRevert(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "depositAndRevert", zrc20, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x0310eb76.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,uint64,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMSession) DepositAndRevert(zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndRevert(&_GatewayZEVM.TransactOpts, zrc20, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x0310eb76.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,uint64,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) DepositAndRevert(zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndRevert(&_GatewayZEVM.TransactOpts, zrc20, amount, target, revertContext)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Execute(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "execute", context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMSession) Execute(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Execute(&_GatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Execute(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Execute(&_GatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x717d335f.
//
// Solidity: function executeRevert(address target, (address,uint64,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) ExecuteRevert(opts *bind.TransactOpts, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "executeRevert", target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x717d335f.
//
// Solidity: function executeRevert(address target, (address,uint64,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.ExecuteRevert(&_GatewayZEVM.TransactOpts, target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x717d335f.
//
// Solidity: function executeRevert(address target, (address,uint64,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.ExecuteRevert(&_GatewayZEVM.TransactOpts, target, revertContext)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayZEVM *GatewayZEVMSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.GrantRole(&_GatewayZEVM.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.GrantRole(&_GatewayZEVM.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address zetaToken_, address admin_) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Initialize(opts *bind.TransactOpts, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "initialize", zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address zetaToken_, address admin_) returns()
func (_GatewayZEVM *GatewayZEVMSession) Initialize(zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Initialize(&_GatewayZEVM.TransactOpts, zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address zetaToken_, address admin_) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Initialize(zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Initialize(&_GatewayZEVM.TransactOpts, zetaToken_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayZEVM *GatewayZEVMSession) Pause() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Pause(&_GatewayZEVM.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Pause() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Pause(&_GatewayZEVM.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayZEVM *GatewayZEVMSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.RenounceRole(&_GatewayZEVM.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.RenounceRole(&_GatewayZEVM.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayZEVM *GatewayZEVMSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.RevokeRole(&_GatewayZEVM.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.RevokeRole(&_GatewayZEVM.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayZEVM *GatewayZEVMSession) Unpause() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Unpause(&_GatewayZEVM.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Unpause() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Unpause(&_GatewayZEVM.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVM *GatewayZEVMTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVM *GatewayZEVMSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.UpgradeToAndCall(&_GatewayZEVM.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.UpgradeToAndCall(&_GatewayZEVM.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x71f6af54.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Withdraw(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdraw", receiver, amount, zrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x71f6af54.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x71f6af54.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x9ec60484.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Withdraw0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdraw0", receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x9ec60484.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw0(&_GatewayZEVM.TransactOpts, receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x9ec60484.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw0(&_GatewayZEVM.TransactOpts, receiver, amount, chainId, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x1e83da29.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) WithdrawAndCall(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdrawAndCall", receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x1e83da29.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) WithdrawAndCall(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.WithdrawAndCall(&_GatewayZEVM.TransactOpts, receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x1e83da29.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) WithdrawAndCall(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.WithdrawAndCall(&_GatewayZEVM.TransactOpts, receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x90ad3e23.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) WithdrawAndCall0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdrawAndCall0", receiver, amount, zrc20, message, gasLimit, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x90ad3e23.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) WithdrawAndCall0(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.WithdrawAndCall0(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, message, gasLimit, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x90ad3e23.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) WithdrawAndCall0(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.WithdrawAndCall0(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, message, gasLimit, revertOptions)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayZEVM *GatewayZEVMSession) Receive() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Receive(&_GatewayZEVM.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Receive() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Receive(&_GatewayZEVM.TransactOpts)
}

// GatewayZEVMCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayZEVM contract.
type GatewayZEVMCalledIterator struct {
	Event *GatewayZEVMCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMCalled)
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
		it.Event = new(GatewayZEVMCalled)
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
func (it *GatewayZEVMCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMCalled represents a Called event raised by the GatewayZEVM contract.
type GatewayZEVMCalled struct {
	Sender        common.Address
	Zrc20         common.Address
	Receiver      []byte
	Message       []byte
	GasLimit      *big.Int
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0x885ccd61b2cc8be1daa203d4456750151d128d2da2ace4c9fa4b63c395fd99d8.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, uint256 gasLimit, (address,bool,address,bytes) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, zrc20 []common.Address) (*GatewayZEVMCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMCalledIterator{contract: _GatewayZEVM.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x885ccd61b2cc8be1daa203d4456750151d128d2da2ace4c9fa4b63c395fd99d8.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, uint256 gasLimit, (address,bool,address,bytes) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMCalled, sender []common.Address, zrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMCalled)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Called", log); err != nil {
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

// ParseCalled is a log parse operation binding the contract event 0x885ccd61b2cc8be1daa203d4456750151d128d2da2ace4c9fa4b63c395fd99d8.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, uint256 gasLimit, (address,bool,address,bytes) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) ParseCalled(log types.Log) (*GatewayZEVMCalled, error) {
	event := new(GatewayZEVMCalled)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayZEVM contract.
type GatewayZEVMInitializedIterator struct {
	Event *GatewayZEVMInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInitialized)
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
		it.Event = new(GatewayZEVMInitialized)
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
func (it *GatewayZEVMInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInitialized represents a Initialized event raised by the GatewayZEVM contract.
type GatewayZEVMInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayZEVMInitializedIterator, error) {

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInitializedIterator{contract: _GatewayZEVM.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInitialized)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseInitialized(log types.Log) (*GatewayZEVMInitialized, error) {
	event := new(GatewayZEVMInitialized)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GatewayZEVM contract.
type GatewayZEVMPausedIterator struct {
	Event *GatewayZEVMPaused // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMPaused)
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
		it.Event = new(GatewayZEVMPaused)
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
func (it *GatewayZEVMPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMPaused represents a Paused event raised by the GatewayZEVM contract.
type GatewayZEVMPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterPaused(opts *bind.FilterOpts) (*GatewayZEVMPausedIterator, error) {

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMPausedIterator{contract: _GatewayZEVM.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GatewayZEVMPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMPaused)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParsePaused(log types.Log) (*GatewayZEVMPaused, error) {
	event := new(GatewayZEVMPaused)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GatewayZEVM contract.
type GatewayZEVMRoleAdminChangedIterator struct {
	Event *GatewayZEVMRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMRoleAdminChanged)
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
		it.Event = new(GatewayZEVMRoleAdminChanged)
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
func (it *GatewayZEVMRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMRoleAdminChanged represents a RoleAdminChanged event raised by the GatewayZEVM contract.
type GatewayZEVMRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GatewayZEVMRoleAdminChangedIterator, error) {

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

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMRoleAdminChangedIterator{contract: _GatewayZEVM.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayZEVMRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMRoleAdminChanged)
				if err := _GatewayZEVM.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseRoleAdminChanged(log types.Log) (*GatewayZEVMRoleAdminChanged, error) {
	event := new(GatewayZEVMRoleAdminChanged)
	if err := _GatewayZEVM.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GatewayZEVM contract.
type GatewayZEVMRoleGrantedIterator struct {
	Event *GatewayZEVMRoleGranted // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMRoleGranted)
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
		it.Event = new(GatewayZEVMRoleGranted)
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
func (it *GatewayZEVMRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMRoleGranted represents a RoleGranted event raised by the GatewayZEVM contract.
type GatewayZEVMRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayZEVMRoleGrantedIterator, error) {

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

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMRoleGrantedIterator{contract: _GatewayZEVM.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GatewayZEVMRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMRoleGranted)
				if err := _GatewayZEVM.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseRoleGranted(log types.Log) (*GatewayZEVMRoleGranted, error) {
	event := new(GatewayZEVMRoleGranted)
	if err := _GatewayZEVM.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GatewayZEVM contract.
type GatewayZEVMRoleRevokedIterator struct {
	Event *GatewayZEVMRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMRoleRevoked)
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
		it.Event = new(GatewayZEVMRoleRevoked)
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
func (it *GatewayZEVMRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMRoleRevoked represents a RoleRevoked event raised by the GatewayZEVM contract.
type GatewayZEVMRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayZEVMRoleRevokedIterator, error) {

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

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMRoleRevokedIterator{contract: _GatewayZEVM.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GatewayZEVMRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMRoleRevoked)
				if err := _GatewayZEVM.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseRoleRevoked(log types.Log) (*GatewayZEVMRoleRevoked, error) {
	event := new(GatewayZEVMRoleRevoked)
	if err := _GatewayZEVM.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GatewayZEVM contract.
type GatewayZEVMUnpausedIterator struct {
	Event *GatewayZEVMUnpaused // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUnpaused)
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
		it.Event = new(GatewayZEVMUnpaused)
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
func (it *GatewayZEVMUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUnpaused represents a Unpaused event raised by the GatewayZEVM contract.
type GatewayZEVMUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GatewayZEVMUnpausedIterator, error) {

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUnpausedIterator{contract: _GatewayZEVM.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUnpaused) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUnpaused)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseUnpaused(log types.Log) (*GatewayZEVMUnpaused, error) {
	event := new(GatewayZEVMUnpaused)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayZEVM contract.
type GatewayZEVMUpgradedIterator struct {
	Event *GatewayZEVMUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgraded)
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
		it.Event = new(GatewayZEVMUpgraded)
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
func (it *GatewayZEVMUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgraded represents a Upgraded event raised by the GatewayZEVM contract.
type GatewayZEVMUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayZEVMUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradedIterator{contract: _GatewayZEVM.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgraded)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseUpgraded(log types.Log) (*GatewayZEVMUpgraded, error) {
	event := new(GatewayZEVMUpgraded)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayZEVM contract.
type GatewayZEVMWithdrawnIterator struct {
	Event *GatewayZEVMWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMWithdrawn)
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
		it.Event = new(GatewayZEVMWithdrawn)
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
func (it *GatewayZEVMWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMWithdrawn represents a Withdrawn event raised by the GatewayZEVM contract.
type GatewayZEVMWithdrawn struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Zrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	GasLimit        *big.Int
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xf6e8f2c19d797cd515ed61d8ac909afa628f4df0bb153a3f81d0f5fb7b8207c0.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 gasLimit, (address,bool,address,bytes) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMWithdrawnIterator{contract: _GatewayZEVM.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xf6e8f2c19d797cd515ed61d8ac909afa628f4df0bb153a3f81d0f5fb7b8207c0.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 gasLimit, (address,bool,address,bytes) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayZEVMWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMWithdrawn)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xf6e8f2c19d797cd515ed61d8ac909afa628f4df0bb153a3f81d0f5fb7b8207c0.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 gasLimit, (address,bool,address,bytes) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) ParseWithdrawn(log types.Log) (*GatewayZEVMWithdrawn, error) {
	event := new(GatewayZEVMWithdrawn)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
