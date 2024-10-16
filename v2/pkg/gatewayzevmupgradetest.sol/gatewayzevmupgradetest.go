// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayzevmupgradetest

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

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Asset         common.Address
	Amount        uint64
	RevertMessage []byte
}

// RevertOptions is an auto generated low-level Go binding around an user-defined struct.
type RevertOptions struct {
	RevertAddress    common.Address
	CallOnRevert     bool
	AbortAddress     common.Address
	RevertMessage    []byte
	OnRevertGasLimit *big.Int
}

// GatewayZEVMUpgradeTestMetaData contains all meta data concerning the GatewayZEVMUpgradeTest contract.
var GatewayZEVMUpgradeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_MESSAGE_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndRevert\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallOnRevertNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotFungibleModule\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrFungible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZETANotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051613a126100fd600039600081816126080152818161263101526128070152613a126000f3fe6080604052600436106101c65760003560e01c806352d1902d116100f757806397d340f511610095578063c39aca3711610064578063c39aca371461061c578063d547741f1461063c578063e63ab1e91461065c578063f45346dc1461069057600080fd5b806397d340f51461057b578063a217fddf14610591578063ad3cb1cc146105a6578063bcf7f32b146105fc57600080fd5b80637c0dcb5f116100d15780637c0dcb5f146104c65780638456cb59146104e657806391d14854146104fb57806397a1cef11461056057600080fd5b806352d1902d1461045a5780635c975abb1461046f578063717d335f146104a657600080fd5b80632f2ff15d116101645780633ce4a5bc1161013e5780633ce4a5bc146103ea5780633f4ba83a14610412578063485cc955146104275780634f1ef2861461044757600080fd5b80632f2ff15d1461038a57806336568abe146103aa5780633b283933146103ca57600080fd5b80631cb5ea75116101a05780631cb5ea75146102b557806321501a95146102d557806321e093b1146102f5578063248a9ca31461032d57600080fd5b806301ffc9a7146102405780630310eb7614610275578063048ae42c1461029557600080fd5b3661023b576101d36106b0565b6000546001600160a01b0316331480159061020257503373735b14bb79463307aacbed86daf3322b1e6226ab14155b15610239576040517f229930b200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b34801561024c57600080fd5b5061026061025b366004612d50565b61070e565b60405190151581526020015b60405180910390f35b34801561028157600080fd5b50610239610290366004612dbf565b6107a7565b3480156102a157600080fd5b506102396102b0366004612f81565b610a33565b3480156102c157600080fd5b506102396102d0366004613041565b610c81565b3480156102e157600080fd5b506102396102f03660046130f6565b610f8b565b34801561030157600080fd5b50600054610315906001600160a01b031681565b6040516001600160a01b03909116815260200161026c565b34801561033957600080fd5b5061037c610348366004613182565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b60405190815260200161026c565b34801561039657600080fd5b506102396103a536600461319b565b611158565b3480156103b657600080fd5b506102396103c536600461319b565b6111a2565b3480156103d657600080fd5b506102396103e53660046131cb565b6111f3565b3480156103f657600080fd5b5061031573735b14bb79463307aacbed86daf3322b1e6226ab81565b34801561041e57600080fd5b50610239611235565b34801561043357600080fd5b5061023961044236600461325e565b61126a565b61023961045536600461328c565b6114bf565b34801561046657600080fd5b5061037c6114de565b34801561047b57600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610260565b3480156104b257600080fd5b506102396104c13660046132dc565b61150d565b3480156104d257600080fd5b506102396104e1366004613322565b611615565b3480156104f257600080fd5b50610239611874565b34801561050757600080fd5b5061026061051636600461319b565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561056c57600080fd5b506102396103e536600461339b565b34801561058757600080fd5b5061037c61040081565b34801561059d57600080fd5b5061037c600081565b3480156105b257600080fd5b506105ef6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161026c919061346d565b34801561060857600080fd5b50610239610617366004613480565b6118a6565b34801561062857600080fd5b50610239610637366004613480565b6119d6565b34801561064857600080fd5b5061023961065736600461319b565b611be5565b34801561066857600080fd5b5061037c7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b34801561069c57600080fd5b506102396106ab36600461351e565b611c29565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff161561070c576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806107a157507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107f4576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107fc6106b0565b6001600160a01b038416158061081957506001600160a01b038216155b15610850576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8260000361088a576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03821673735b14bb79463307aacbed86daf3322b1e6226ab14806108bd57506001600160a01b03821630145b156108f4576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152602482018590528516906347e7ef24906044016020604051808303816000875af115801561095c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610980919061356e565b6109b6576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f660b9de00000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063660b9de0906109fb908490600401613638565b600060405180830381600087803b158015610a1557600080fd5b505af1158015610a29573d6000803e3d6000fd5b5050505050505050565b610a3b611e38565b610a436106b0565b610a53604082016020830161369e565b15610a8a576040517f19b4bff200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8651600003610ac5576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b85600003610aff576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600003610b39576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610400610b4960608301836136bb565b610b54915085613720565b10610b8b576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610b98878785611eb9565b90506000336001600160a01b03167fda1215b0949ddb309fe466fa9e70e861a16538f11b8ecdb05c217d4d8677ed2d8a898b868c6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c2d919061375a565b8c8c8c8c604051610c46999897969594939291906137f4565b60405180910390a350610c7860017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050505050565b610c89611e38565b610c916106b0565b610ca1604082016020830161369e565b15610cd8576040517f19b4bff200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8551600003610d13576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600003610d4d576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610400610d5d60608301836136bb565b610d68915085613720565b1115610da0576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517ffc5fecd50000000000000000000000000000000000000000000000000000000081526004810183905260009081906001600160a01b0388169063fc5fecd5906024016040805180830381865afa158015610e02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e269190613867565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab60248201526044810182905291935091506001600160a01b038316906323b872dd906064016020604051808303816000875af1158015610eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ecf919061356e565b610f05576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b866001600160a01b0316336001600160a01b03167f6c6abd640fc6a0ef7cf2bc54b246b42d5c2629c30be1e24fea4a58157a7728cf8a89898989604051610f50959493929190613895565b60405180910390a35050610f8360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610fd8576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fe06106b0565b6001600160a01b038316611020576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8360000361105a576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab148061108d57506001600160a01b03831630145b156110c4576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110ce84846121bc565b6000546040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b0380861692635bcfd6169261111f928a9216908990889088906004016138e1565b600060405180830381600087803b15801561113957600080fd5b505af115801561114d573d6000803e3d6000fd5b505050505050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546111928161238a565b61119c8383612394565b50505050565b6001600160a01b03811633146111e4576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111ee8282612481565b505050565b6111fb611e38565b6112036106b0565b6040517fe4dd681d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61125f8161238a565b611267612545565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156112b55750825b905060008267ffffffffffffffff1660011480156112d25750303b155b9050811580156112e0575080155b15611317576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156113785784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038716158061139557506001600160a01b038616155b156113cc576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113d46125d5565b6113dc6125d5565b6113e46125dd565b6113ec6125ed565b6113f7600087612394565b506114227f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87612394565b50600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0389161790558315610c785784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050505050565b6114c76125fd565b6114d0826126cd565b6114da82826126d8565b5050565b60006114e86127fc565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461155a576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115626106b0565b6001600160a01b0382166115a2576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f660b9de00000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063660b9de0906115e7908490600401613638565b600060405180830381600087803b15801561160157600080fd5b505af1158015610f83573d6000803e3d6000fd5b61161d611e38565b6116256106b0565b611635604082016020830161369e565b1561166c576040517f19b4bff200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83516000036116a7576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826000036116e1576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104006116f160608301836136bb565b9050111561172b576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611737848461285e565b90506000336001600160a01b03167f3adf8a8d73037d29dfb621546205d69f95d01f9613c3c811af66901fc1d3557887868886896001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117cc919061375a565b8a6001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa15801561180a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061182e919061375a565b8a6040516118429796959493929190613952565b60405180910390a35061119c60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61189e8161238a565b6112676128cc565b3373735b14bb79463307aacbed86daf3322b1e6226ab146118f3576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6118fb6106b0565b6001600160a01b038516158061191857506001600160a01b038316155b1561194f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b03841690635bcfd6169061199c90899089908990889088906004016138e1565b600060405180830381600087803b1580156119b657600080fd5b505af11580156119ca573d6000803e3d6000fd5b50505050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611a23576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611a2b6106b0565b6001600160a01b0385161580611a4857506001600160a01b038316155b15611a7f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600003611ab9576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab1480611aec57506001600160a01b03831630145b15611b23576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038481166004830152602482018690528616906347e7ef24906044016020604051808303816000875af1158015611b8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611baf919061356e565b61194f576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611c1f8161238a565b61119c8383612481565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611c76576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c7e6106b0565b6001600160a01b0383161580611c9b57506001600160a01b038116155b15611cd2576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600003611d0c576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03811673735b14bb79463307aacbed86daf3322b1e6226ab1480611d3f57506001600160a01b03811630145b15611d76576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038281166004830152602482018490528416906347e7ef24906044016020604051808303816000875af1158015611dde573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e02919061356e565b6111ee576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611eb3576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6000806000846001600160a01b031663fc5fecd5856040518263ffffffff1660e01b8152600401611eec91815260200190565b6040805180830381865afa158015611f08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f2c9190613867565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab60248201526044810182905291935091506001600160a01b038316906323b872dd906064016020604051808303816000875af1158015611fb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fd5919061356e565b61200b576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018790526001600160a01b038616906323b872dd906064016020604051808303816000875af1158015612077573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061209b919061356e565b6120d1576040517f4dd9ee8d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f42966c68000000000000000000000000000000000000000000000000000000008152600481018790526001600160a01b038616906342966c68906024016020604051808303816000875af1158015612131573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612155919061356e565b61218b576040517f2c77e05c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9150505b9392505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6000546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490526001600160a01b03909116906323b872dd906064016020604051808303816000875af115801561222c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612250919061356e565b612286576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000546040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b0390911690632e1a7d4d90602401600060405180830381600087803b1580156122e557600080fd5b505af11580156122f9573d6000803e3d6000fd5b505050506000816001600160a01b03168360405160006040518083038185875af1925050503d806000811461234a576040519150601f19603f3d011682016040523d82523d6000602084013e61234f565b606091505b50509050806111ee576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112678133612945565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612477576000848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561242d3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506107a1565b60009150506107a1565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612477576000848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506107a1565b61254d6129d2565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b61070c612a2d565b6125e5612a2d565b61070c612a94565b6125f5612a2d565b61070c612ae5565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061269657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661268a7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b1561070c576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006114da8161238a565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612750575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261274d9181019061375a565b60015b612796576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146127f2576040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004810182905260240161278d565b6111ee8383612aed565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461070c576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061218f8383846001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128c7919061375a565b611eb9565b6128d46106b0565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336125b7565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff166114da576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024810183905260440161278d565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661070c576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661070c576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612a9c612a2d565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b612196612a2d565b612af682612b43565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115612b3b576111ee8282612beb565b6114da612c61565b806001600160a01b03163b600003612b92576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161278d565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051612c0891906139c0565b600060405180830381855af49150503d8060008114612c43576040519150601f19603f3d011682016040523d82523d6000602084013e612c48565b606091505b5091509150612c58858383612c99565b95945050505050565b341561070c576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082612cae57612ca982612d0e565b61218f565b8151158015612cc557506001600160a01b0384163b155b15612d07576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161278d565b508061218f565b805115612d1e5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215612d6257600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461218f57600080fd5b6001600160a01b038116811461126757600080fd5b600060608284031215612db957600080fd5b50919050565b60008060008060808587031215612dd557600080fd5b8435612de081612d92565b9350602085013592506040850135612df781612d92565b9150606085013567ffffffffffffffff811115612e1357600080fd5b612e1f87828801612da7565b91505092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112612e6b57600080fd5b813567ffffffffffffffff811115612e8557612e85612e2b565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715612ef157612ef1612e2b565b604052818152838201602001851015612f0957600080fd5b816020850160208301376000918101602001919091529392505050565b60008083601f840112612f3857600080fd5b50813567ffffffffffffffff811115612f5057600080fd5b602083019150836020828501011115612f6857600080fd5b9250929050565b600060a08284031215612db957600080fd5b600080600080600080600060c0888a031215612f9c57600080fd5b873567ffffffffffffffff811115612fb357600080fd5b612fbf8a828b01612e5a565b975050602088013595506040880135612fd781612d92565b9450606088013567ffffffffffffffff811115612ff357600080fd5b612fff8a828b01612f26565b9095509350506080880135915060a088013567ffffffffffffffff81111561302657600080fd5b6130328a828b01612f6f565b91505092959891949750929550565b60008060008060008060a0878903121561305a57600080fd5b863567ffffffffffffffff81111561307157600080fd5b61307d89828a01612e5a565b965050602087013561308e81612d92565b9450604087013567ffffffffffffffff8111156130aa57600080fd5b6130b689828a01612f26565b90955093505060608701359150608087013567ffffffffffffffff8111156130dd57600080fd5b6130e989828a01612f6f565b9150509295509295509295565b60008060008060006080868803121561310e57600080fd5b853567ffffffffffffffff81111561312557600080fd5b61313188828901612da7565b95505060208601359350604086013561314981612d92565b9250606086013567ffffffffffffffff81111561316557600080fd5b61317188828901612f26565b969995985093965092949392505050565b60006020828403121561319457600080fd5b5035919050565b600080604083850312156131ae57600080fd5b8235915060208301356131c081612d92565b809150509250929050565b60008060008060008060a087890312156131e457600080fd5b863567ffffffffffffffff8111156131fb57600080fd5b61320789828a01612e5a565b9650506020870135945060408701359350606087013567ffffffffffffffff81111561323257600080fd5b61323e89828a01612f26565b909450925050608087013567ffffffffffffffff8111156130dd57600080fd5b6000806040838503121561327157600080fd5b823561327c81612d92565b915060208301356131c081612d92565b6000806040838503121561329f57600080fd5b82356132aa81612d92565b9150602083013567ffffffffffffffff8111156132c657600080fd5b6132d285828601612e5a565b9150509250929050565b600080604083850312156132ef57600080fd5b82356132fa81612d92565b9150602083013567ffffffffffffffff81111561331657600080fd5b6132d285828601612da7565b6000806000806080858703121561333857600080fd5b843567ffffffffffffffff81111561334f57600080fd5b61335b87828801612e5a565b94505060208501359250604085013561337381612d92565b9150606085013567ffffffffffffffff81111561338f57600080fd5b612e1f87828801612f6f565b600080600080608085870312156133b157600080fd5b843567ffffffffffffffff8111156133c857600080fd5b6133d487828801612e5a565b9450506020850135925060408501359150606085013567ffffffffffffffff81111561338f57600080fd5b60005b8381101561341a578181015183820152602001613402565b50506000910152565b6000815180845261343b8160208601602086016133ff565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061218f6020830184613423565b60008060008060008060a0878903121561349957600080fd5b863567ffffffffffffffff8111156134b057600080fd5b6134bc89828a01612da7565b96505060208701356134cd81612d92565b94506040870135935060608701356134e481612d92565b9250608087013567ffffffffffffffff81111561350057600080fd5b61350c89828a01612f26565b979a9699509497509295939492505050565b60008060006060848603121561353357600080fd5b833561353e81612d92565b925060208401359150604084013561355581612d92565b809150509250925092565b801515811461126757600080fd5b60006020828403121561358057600080fd5b815161218f81613560565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126135c057600080fd5b830160208101925035905067ffffffffffffffff8111156135e057600080fd5b803603821315612f6857600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b602081526000823561364981612d92565b6001600160a01b038116602084015250602083013567ffffffffffffffff811680821461367557600080fd5b806040850152505061368a604084018461358b565b606080850152612c586080850182846135ef565b6000602082840312156136b057600080fd5b813561218f81613560565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126136f057600080fd5b83018035915067ffffffffffffffff82111561370b57600080fd5b602001915036819003821315612f6857600080fd5b808201808211156107a1577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006020828403121561376c57600080fd5b5051919050565b6000813561378081612d92565b6001600160a01b03168352602082013561379981613560565b1515602084015260408201356137ae81612d92565b6001600160a01b031660408401526137c9606083018361358b565b60a060608601526137de60a0860182846135ef565b6080948501359590940194909452509092915050565b6101008152600061380961010083018c613423565b6001600160a01b038b16602084015289604084015288606084015287608084015282810360a084015261383d8187896135ef565b90508460c084015282810360e08401526138578185613773565b9c9b505050505050505050505050565b6000806040838503121561387a57600080fd5b825161388581612d92565b6020939093015192949293505050565b6080815260006138a86080830188613423565b82810360208401526138bb8187896135ef565b905084604084015282810360608401526138d58185613773565b98975050505050505050565b6080815260006138f1878861358b565b6060608085015261390660e0850182846135ef565b915050602088013561391781612d92565b6001600160a01b0390811660a085015260408981013560c08601529088166020850152830186905282810360608401526138d58185876135ef565b6101008152600061396761010083018a613423565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a0850152600082528560c08501526020810160e0850152506139b26020820185613773565b9a9950505050505050505050565b600082516139d28184602087016133ff565b919091019291505056fea26469706673582212204fdb58d6b9d313e5e1051a4c98078e3c6a8baff2220a2505931828e8ef6a4f4464736f6c634300081a0033",
}

// GatewayZEVMUpgradeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayZEVMUpgradeTestMetaData.ABI instead.
var GatewayZEVMUpgradeTestABI = GatewayZEVMUpgradeTestMetaData.ABI

// GatewayZEVMUpgradeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayZEVMUpgradeTestMetaData.Bin instead.
var GatewayZEVMUpgradeTestBin = GatewayZEVMUpgradeTestMetaData.Bin

// DeployGatewayZEVMUpgradeTest deploys a new Ethereum contract, binding an instance of GatewayZEVMUpgradeTest to it.
func DeployGatewayZEVMUpgradeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayZEVMUpgradeTest, error) {
	parsed, err := GatewayZEVMUpgradeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayZEVMUpgradeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayZEVMUpgradeTest{GatewayZEVMUpgradeTestCaller: GatewayZEVMUpgradeTestCaller{contract: contract}, GatewayZEVMUpgradeTestTransactor: GatewayZEVMUpgradeTestTransactor{contract: contract}, GatewayZEVMUpgradeTestFilterer: GatewayZEVMUpgradeTestFilterer{contract: contract}}, nil
}

// GatewayZEVMUpgradeTest is an auto generated Go binding around an Ethereum contract.
type GatewayZEVMUpgradeTest struct {
	GatewayZEVMUpgradeTestCaller     // Read-only binding to the contract
	GatewayZEVMUpgradeTestTransactor // Write-only binding to the contract
	GatewayZEVMUpgradeTestFilterer   // Log filterer for contract events
}

// GatewayZEVMUpgradeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayZEVMUpgradeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMUpgradeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayZEVMUpgradeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMUpgradeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayZEVMUpgradeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMUpgradeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayZEVMUpgradeTestSession struct {
	Contract     *GatewayZEVMUpgradeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GatewayZEVMUpgradeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayZEVMUpgradeTestCallerSession struct {
	Contract *GatewayZEVMUpgradeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// GatewayZEVMUpgradeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayZEVMUpgradeTestTransactorSession struct {
	Contract     *GatewayZEVMUpgradeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// GatewayZEVMUpgradeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayZEVMUpgradeTestRaw struct {
	Contract *GatewayZEVMUpgradeTest // Generic contract binding to access the raw methods on
}

// GatewayZEVMUpgradeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayZEVMUpgradeTestCallerRaw struct {
	Contract *GatewayZEVMUpgradeTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayZEVMUpgradeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayZEVMUpgradeTestTransactorRaw struct {
	Contract *GatewayZEVMUpgradeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayZEVMUpgradeTest creates a new instance of GatewayZEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayZEVMUpgradeTest(address common.Address, backend bind.ContractBackend) (*GatewayZEVMUpgradeTest, error) {
	contract, err := bindGatewayZEVMUpgradeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTest{GatewayZEVMUpgradeTestCaller: GatewayZEVMUpgradeTestCaller{contract: contract}, GatewayZEVMUpgradeTestTransactor: GatewayZEVMUpgradeTestTransactor{contract: contract}, GatewayZEVMUpgradeTestFilterer: GatewayZEVMUpgradeTestFilterer{contract: contract}}, nil
}

// NewGatewayZEVMUpgradeTestCaller creates a new read-only instance of GatewayZEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayZEVMUpgradeTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayZEVMUpgradeTestCaller, error) {
	contract, err := bindGatewayZEVMUpgradeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestCaller{contract: contract}, nil
}

// NewGatewayZEVMUpgradeTestTransactor creates a new write-only instance of GatewayZEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayZEVMUpgradeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayZEVMUpgradeTestTransactor, error) {
	contract, err := bindGatewayZEVMUpgradeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestTransactor{contract: contract}, nil
}

// NewGatewayZEVMUpgradeTestFilterer creates a new log filterer instance of GatewayZEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayZEVMUpgradeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayZEVMUpgradeTestFilterer, error) {
	contract, err := bindGatewayZEVMUpgradeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestFilterer{contract: contract}, nil
}

// bindGatewayZEVMUpgradeTest binds a generic wrapper to an already deployed contract.
func bindGatewayZEVMUpgradeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayZEVMUpgradeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMUpgradeTest.Contract.GatewayZEVMUpgradeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.GatewayZEVMUpgradeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.GatewayZEVMUpgradeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMUpgradeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.DEFAULTADMINROLE(&_GatewayZEVMUpgradeTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.DEFAULTADMINROLE(&_GatewayZEVMUpgradeTest.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _GatewayZEVMUpgradeTest.Contract.FUNGIBLEMODULEADDRESS(&_GatewayZEVMUpgradeTest.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _GatewayZEVMUpgradeTest.Contract.FUNGIBLEMODULEADDRESS(&_GatewayZEVMUpgradeTest.CallOpts)
}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) MAXMESSAGESIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "MAX_MESSAGE_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) MAXMESSAGESIZE() (*big.Int, error) {
	return _GatewayZEVMUpgradeTest.Contract.MAXMESSAGESIZE(&_GatewayZEVMUpgradeTest.CallOpts)
}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) MAXMESSAGESIZE() (*big.Int, error) {
	return _GatewayZEVMUpgradeTest.Contract.MAXMESSAGESIZE(&_GatewayZEVMUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.PAUSERROLE(&_GatewayZEVMUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.PAUSERROLE(&_GatewayZEVMUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayZEVMUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_GatewayZEVMUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayZEVMUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_GatewayZEVMUpgradeTest.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.GetRoleAdmin(&_GatewayZEVMUpgradeTest.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.GetRoleAdmin(&_GatewayZEVMUpgradeTest.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayZEVMUpgradeTest.Contract.HasRole(&_GatewayZEVMUpgradeTest.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayZEVMUpgradeTest.Contract.HasRole(&_GatewayZEVMUpgradeTest.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Paused() (bool, error) {
	return _GatewayZEVMUpgradeTest.Contract.Paused(&_GatewayZEVMUpgradeTest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) Paused() (bool, error) {
	return _GatewayZEVMUpgradeTest.Contract.Paused(&_GatewayZEVMUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.ProxiableUUID(&_GatewayZEVMUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.ProxiableUUID(&_GatewayZEVMUpgradeTest.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayZEVMUpgradeTest.Contract.SupportsInterface(&_GatewayZEVMUpgradeTest.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayZEVMUpgradeTest.Contract.SupportsInterface(&_GatewayZEVMUpgradeTest.CallOpts, interfaceId)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) ZetaToken() (common.Address, error) {
	return _GatewayZEVMUpgradeTest.Contract.ZetaToken(&_GatewayZEVMUpgradeTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) ZetaToken() (common.Address, error) {
	return _GatewayZEVMUpgradeTest.Contract.ZetaToken(&_GatewayZEVMUpgradeTest.CallOpts)
}

// Call is a paid mutator transaction binding the contract method 0x1cb5ea75.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Call(opts *bind.TransactOpts, receiver []byte, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "call", receiver, zrc20, message, gasLimit, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x1cb5ea75.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Call(receiver []byte, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Call(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, zrc20, message, gasLimit, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x1cb5ea75.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Call(receiver []byte, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Call(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, zrc20, message, gasLimit, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Deposit(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "deposit", zrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Deposit(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Deposit(&_GatewayZEVMUpgradeTest.TransactOpts, zrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Deposit(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Deposit(&_GatewayZEVMUpgradeTest.TransactOpts, zrc20, amount, target)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) DepositAndCall(opts *bind.TransactOpts, context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "depositAndCall", context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) DepositAndCall(context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) DepositAndCall(context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, context, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) DepositAndCall0(opts *bind.TransactOpts, context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "depositAndCall0", context, zrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) DepositAndCall0(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndCall0(&_GatewayZEVMUpgradeTest.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) DepositAndCall0(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndCall0(&_GatewayZEVMUpgradeTest.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x0310eb76.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,uint64,bytes) revertContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) DepositAndRevert(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "depositAndRevert", zrc20, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x0310eb76.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,uint64,bytes) revertContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) DepositAndRevert(zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndRevert(&_GatewayZEVMUpgradeTest.TransactOpts, zrc20, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x0310eb76.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,uint64,bytes) revertContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) DepositAndRevert(zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndRevert(&_GatewayZEVMUpgradeTest.TransactOpts, zrc20, amount, target, revertContext)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Execute(opts *bind.TransactOpts, context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "execute", context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Execute(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Execute(&_GatewayZEVMUpgradeTest.TransactOpts, context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Execute(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Execute(&_GatewayZEVMUpgradeTest.TransactOpts, context, zrc20, amount, target, message)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x717d335f.
//
// Solidity: function executeRevert(address target, (address,uint64,bytes) revertContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) ExecuteRevert(opts *bind.TransactOpts, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "executeRevert", target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x717d335f.
//
// Solidity: function executeRevert(address target, (address,uint64,bytes) revertContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.ExecuteRevert(&_GatewayZEVMUpgradeTest.TransactOpts, target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x717d335f.
//
// Solidity: function executeRevert(address target, (address,uint64,bytes) revertContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.ExecuteRevert(&_GatewayZEVMUpgradeTest.TransactOpts, target, revertContext)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.GrantRole(&_GatewayZEVMUpgradeTest.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.GrantRole(&_GatewayZEVMUpgradeTest.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address zetaToken_, address admin_) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Initialize(opts *bind.TransactOpts, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "initialize", zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address zetaToken_, address admin_) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Initialize(zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Initialize(&_GatewayZEVMUpgradeTest.TransactOpts, zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address zetaToken_, address admin_) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Initialize(zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Initialize(&_GatewayZEVMUpgradeTest.TransactOpts, zetaToken_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Pause() (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Pause(&_GatewayZEVMUpgradeTest.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Pause() (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Pause(&_GatewayZEVMUpgradeTest.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.RenounceRole(&_GatewayZEVMUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.RenounceRole(&_GatewayZEVMUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.RevokeRole(&_GatewayZEVMUpgradeTest.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.RevokeRole(&_GatewayZEVMUpgradeTest.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Unpause() (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Unpause(&_GatewayZEVMUpgradeTest.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Unpause() (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Unpause(&_GatewayZEVMUpgradeTest.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.UpgradeToAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.UpgradeToAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Withdraw(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdraw", receiver, amount, zrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Withdraw0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdraw0", receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, chainId, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x048ae42c.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) WithdrawAndCall(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall", receiver, amount, zrc20, message, gasLimit, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x048ae42c.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) WithdrawAndCall(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, message, gasLimit, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x048ae42c.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) WithdrawAndCall(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, message, gasLimit, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x3b283933.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) WithdrawAndCall0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall0", receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x3b283933.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) WithdrawAndCall0(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x3b283933.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) WithdrawAndCall0(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, chainId, message, revertOptions)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Receive() (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Receive(&_GatewayZEVMUpgradeTest.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Receive() (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Receive(&_GatewayZEVMUpgradeTest.TransactOpts)
}

// GatewayZEVMUpgradeTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestCalledIterator struct {
	Event *GatewayZEVMUpgradeTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestCalled)
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
		it.Event = new(GatewayZEVMUpgradeTestCalled)
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
func (it *GatewayZEVMUpgradeTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestCalled represents a Called event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestCalled struct {
	Sender        common.Address
	Zrc20         common.Address
	Receiver      []byte
	Message       []byte
	GasLimit      *big.Int
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0x6c6abd640fc6a0ef7cf2bc54b246b42d5c2629c30be1e24fea4a58157a7728cf.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, zrc20 []common.Address) (*GatewayZEVMUpgradeTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestCalledIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x6c6abd640fc6a0ef7cf2bc54b246b42d5c2629c30be1e24fea4a58157a7728cf.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestCalled, sender []common.Address, zrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestCalled)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Called", log); err != nil {
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

// ParseCalled is a log parse operation binding the contract event 0x6c6abd640fc6a0ef7cf2bc54b246b42d5c2629c30be1e24fea4a58157a7728cf.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseCalled(log types.Log) (*GatewayZEVMUpgradeTestCalled, error) {
	event := new(GatewayZEVMUpgradeTestCalled)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestInitializedIterator struct {
	Event *GatewayZEVMUpgradeTestInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestInitialized)
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
		it.Event = new(GatewayZEVMUpgradeTestInitialized)
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
func (it *GatewayZEVMUpgradeTestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestInitialized represents a Initialized event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayZEVMUpgradeTestInitializedIterator, error) {

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestInitializedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestInitialized)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseInitialized(log types.Log) (*GatewayZEVMUpgradeTestInitialized, error) {
	event := new(GatewayZEVMUpgradeTestInitialized)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestPausedIterator struct {
	Event *GatewayZEVMUpgradeTestPaused // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestPaused)
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
		it.Event = new(GatewayZEVMUpgradeTestPaused)
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
func (it *GatewayZEVMUpgradeTestPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestPaused represents a Paused event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterPaused(opts *bind.FilterOpts) (*GatewayZEVMUpgradeTestPausedIterator, error) {

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestPausedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestPaused)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParsePaused(log types.Log) (*GatewayZEVMUpgradeTestPaused, error) {
	event := new(GatewayZEVMUpgradeTestPaused)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestRoleAdminChangedIterator struct {
	Event *GatewayZEVMUpgradeTestRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestRoleAdminChanged)
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
		it.Event = new(GatewayZEVMUpgradeTestRoleAdminChanged)
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
func (it *GatewayZEVMUpgradeTestRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestRoleAdminChanged represents a RoleAdminChanged event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GatewayZEVMUpgradeTestRoleAdminChangedIterator, error) {

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

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestRoleAdminChangedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestRoleAdminChanged)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseRoleAdminChanged(log types.Log) (*GatewayZEVMUpgradeTestRoleAdminChanged, error) {
	event := new(GatewayZEVMUpgradeTestRoleAdminChanged)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestRoleGrantedIterator struct {
	Event *GatewayZEVMUpgradeTestRoleGranted // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestRoleGranted)
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
		it.Event = new(GatewayZEVMUpgradeTestRoleGranted)
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
func (it *GatewayZEVMUpgradeTestRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestRoleGranted represents a RoleGranted event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayZEVMUpgradeTestRoleGrantedIterator, error) {

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

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestRoleGrantedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestRoleGranted)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseRoleGranted(log types.Log) (*GatewayZEVMUpgradeTestRoleGranted, error) {
	event := new(GatewayZEVMUpgradeTestRoleGranted)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestRoleRevokedIterator struct {
	Event *GatewayZEVMUpgradeTestRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestRoleRevoked)
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
		it.Event = new(GatewayZEVMUpgradeTestRoleRevoked)
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
func (it *GatewayZEVMUpgradeTestRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestRoleRevoked represents a RoleRevoked event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayZEVMUpgradeTestRoleRevokedIterator, error) {

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

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestRoleRevokedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestRoleRevoked)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseRoleRevoked(log types.Log) (*GatewayZEVMUpgradeTestRoleRevoked, error) {
	event := new(GatewayZEVMUpgradeTestRoleRevoked)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestUnpausedIterator struct {
	Event *GatewayZEVMUpgradeTestUnpaused // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestUnpaused)
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
		it.Event = new(GatewayZEVMUpgradeTestUnpaused)
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
func (it *GatewayZEVMUpgradeTestUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestUnpaused represents a Unpaused event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GatewayZEVMUpgradeTestUnpausedIterator, error) {

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestUnpausedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestUnpaused) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestUnpaused)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseUnpaused(log types.Log) (*GatewayZEVMUpgradeTestUnpaused, error) {
	event := new(GatewayZEVMUpgradeTestUnpaused)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestUpgradedIterator struct {
	Event *GatewayZEVMUpgradeTestUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestUpgraded)
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
		it.Event = new(GatewayZEVMUpgradeTestUpgraded)
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
func (it *GatewayZEVMUpgradeTestUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestUpgraded represents a Upgraded event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayZEVMUpgradeTestUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestUpgradedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestUpgraded)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseUpgraded(log types.Log) (*GatewayZEVMUpgradeTestUpgraded, error) {
	event := new(GatewayZEVMUpgradeTestUpgraded)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestWithdrawnIterator struct {
	Event *GatewayZEVMUpgradeTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestWithdrawn)
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
		it.Event = new(GatewayZEVMUpgradeTestWithdrawn)
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
func (it *GatewayZEVMUpgradeTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestWithdrawn represents a Withdrawn event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestWithdrawn struct {
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

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xda1215b0949ddb309fe466fa9e70e861a16538f11b8ecdb05c217d4d8677ed2d.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMUpgradeTestWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestWithdrawnIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xda1215b0949ddb309fe466fa9e70e861a16538f11b8ecdb05c217d4d8677ed2d.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestWithdrawn)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xda1215b0949ddb309fe466fa9e70e861a16538f11b8ecdb05c217d4d8677ed2d.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseWithdrawn(log types.Log) (*GatewayZEVMUpgradeTestWithdrawn, error) {
	event := new(GatewayZEVMUpgradeTestWithdrawn)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestWithdrawnV2Iterator struct {
	Event *GatewayZEVMUpgradeTestWithdrawnV2 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestWithdrawnV2)
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
		it.Event = new(GatewayZEVMUpgradeTestWithdrawnV2)
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
func (it *GatewayZEVMUpgradeTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestWithdrawnV2 represents a WithdrawnV2 event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestWithdrawnV2 struct {
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

// FilterWithdrawnV2 is a free log retrieval operation binding the contract event 0x3adf8a8d73037d29dfb621546205d69f95d01f9613c3c811af66901fc1d35578.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMUpgradeTestWithdrawnV2Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "WithdrawnV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestWithdrawnV2Iterator{contract: _GatewayZEVMUpgradeTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0x3adf8a8d73037d29dfb621546205d69f95d01f9613c3c811af66901fc1d35578.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestWithdrawnV2, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "WithdrawnV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestWithdrawnV2)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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

// ParseWithdrawnV2 is a log parse operation binding the contract event 0x3adf8a8d73037d29dfb621546205d69f95d01f9613c3c811af66901fc1d35578.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseWithdrawnV2(log types.Log) (*GatewayZEVMUpgradeTestWithdrawnV2, error) {
	event := new(GatewayZEVMUpgradeTestWithdrawnV2)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
