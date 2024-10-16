// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayevm

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
	Sender common.Address
}

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
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

// GatewayEVMMetaData contains all meta data concerning the GatewayEVM contract.
var GatewayEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ASSET_HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PAYLOAD_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"custody\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"messageContext\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revertWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConnector\",\"inputs\":[{\"name\":\"zetaConnector_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustody\",\"inputs\":[{\"name\":\"custody_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"zetaConnector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZETANotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051613b7b6100fd600039600081816123e20152818161240b01526128790152613b7b6000f3fe6080604052600436106102345760003560e01c80635d62c86011610138578063aa0c0fc1116100b0578063cb7ba8e51161007f578063d547741f11610064578063d547741f146106db578063dda79b75146106fb578063e63ab1e91461071b57600080fd5b8063cb7ba8e5146106a8578063d09e3b78146106bb57600080fd5b8063aa0c0fc1146105ff578063ad3cb1cc1461061f578063ae7a3a6f14610668578063c0c53b8b1461068857600080fd5b806391d1485411610107578063a217fddf116100ec578063a217fddf146105a0578063a2ba1934146105b5578063a783c789146105cb57600080fd5b806391d148541461051b578063950837aa1461058057600080fd5b80635d62c860146104ac578063726ac97c146104e0578063744b9b8b146104f35780638456cb591461050657600080fd5b806336568abe116101cb5780635131ab591161019a57806357bec62f1161017f57806357bec62f146104355780635b112591146104555780635c975abb1461047557600080fd5b80635131ab591461040057806352d1902d1461042057600080fd5b806336568abe146103a557806338e22527146103c55780633f4ba83a146103d85780634f1ef286146103ed57600080fd5b80631cff79cd116102075780631cff79cd146102d057806321e093b1146102f0578063248a9ca3146103285780632f2ff15d1461038557600080fd5b806301ffc9a71461023957806310188aef1461026e578063102614b0146102905780631becceb4146102b0575b600080fd5b34801561024557600080fd5b50610259610254366004613060565b61074f565b60405190151581526020015b60405180910390f35b34801561027a57600080fd5b5061028e6102893660046130be565b6107e8565b005b34801561029c57600080fd5b5061028e6102ab3660046130f1565b6108c3565b3480156102bc57600080fd5b5061028e6102cb3660046131a2565b610a07565b6102e36102de366004613209565b610ad8565b60405161026591906132ac565b3480156102fc57600080fd5b50600354610310906001600160a01b031681565b6040516001600160a01b039091168152602001610265565b34801561033457600080fd5b506103776103433660046132bf565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b604051908152602001610265565b34801561039157600080fd5b5061028e6103a03660046132d8565b610b90565b3480156103b157600080fd5b5061028e6103c03660046132d8565b610bd4565b6102e36103d3366004613304565b610c25565b3480156103e457600080fd5b5061028e610d11565b61028e6103fb3660046133f5565b610d46565b34801561040c57600080fd5b5061028e61041b366004613486565b610d65565b34801561042c57600080fd5b50610377611065565b34801561044157600080fd5b50600254610310906001600160a01b031681565b34801561046157600080fd5b50600154610310906001600160a01b031681565b34801561048157600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610259565b3480156104b857600080fd5b506103777f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b981565b61028e6104ee3660046134f5565b611094565b61028e6105013660046131a2565b611256565b34801561051257600080fd5b5061028e611425565b34801561052757600080fd5b506102596105363660046132d8565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561058c57600080fd5b5061028e61059b3660046130be565b611457565b3480156105ac57600080fd5b50610377600081565b3480156105c157600080fd5b5061037761040081565b3480156105d757600080fd5b506103777f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b34801561060b57600080fd5b5061028e61061a366004613555565b61156e565b34801561062b57600080fd5b506102e36040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b34801561067457600080fd5b5061028e6106833660046130be565b611717565b34801561069457600080fd5b5061028e6106a33660046135ed565b6117f2565b61028e6106b6366004613630565b611ab9565b3480156106c757600080fd5b5061028e6106d63660046136a3565b611ca1565b3480156106e757600080fd5b5061028e6106f63660046132d8565b611dec565b34801561070757600080fd5b50600054610310906001600160a01b031681565b34801561072757600080fd5b506103777f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806107e257507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b60006107f381611e30565b6001600160a01b03821661081a5760405163d92e233d60e01b815260040160405180910390fd5b6002546001600160a01b03161561085d576040517f0c8dc01600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108877f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b983611e3a565b5050600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6108cb611f27565b6108d3611f85565b8260000361090d576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384166109345760405163d92e233d60e01b815260040160405180910390fd5b610400610944606083018361372e565b9050111561097e576040517f386691aa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610989338385612006565b836001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c8585856040516109d0939291906138ad565b60405180910390a3610a0160017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050565b610a0f611f27565b610a17611f85565b6001600160a01b038416610a3e5760405163d92e233d60e01b815260040160405180910390fd5b610400610a4e606083018361372e565b610a599150846138e3565b1115610a91576040517f386691aa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836001600160a01b0316336001600160a01b03167fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d9748585856040516109d09392919061391d565b60607f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb610b0481611e30565b610b0c611f27565b6001600160a01b038516610b335760405163d92e233d60e01b815260040160405180910390fd5b6000610b4086868661214d565b9050856001600160a01b03167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f348787604051610b7f93929190613943565b60405180910390a295945050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610bca81611e30565b610a018383611e3a565b6001600160a01b0381163314610c16576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c208282612200565b505050565b60607f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb610c5181611e30565b610c59611f27565b610c61611f85565b6001600160a01b038516610c885760405163d92e233d60e01b815260040160405180910390fd5b6060610c96878787876122c4565b9050856001600160a01b03167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f348787604051610cd593929190613943565b60405180910390a29150610d0860017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50949350505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610d3b81611e30565b610d43612347565b50565b610d4e6123d7565b610d57826124a7565b610d6182826124b2565b5050565b7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b9610d8f81611e30565b610d97611f27565b610d9f611f85565b83600003610dd9576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038516610e005760405163d92e233d60e01b815260040160405180910390fd5b610e0a86866125b8565b610e40576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0386811660048301526024820186905287169063095ea7b3906044016020604051808303816000875af1158015610ea8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ecc919061395d565b610f02576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f0d85848461214d565b50610f1886866125b8565b610f4e576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038816906370a0823190602401602060405180830381865afa158015610fae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd2919061397a565b90508015610fe457610fe48782612648565b856001600160a01b0316876001600160a01b03167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b738287878760405161102b93929190613943565b60405180910390a35061105d60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b600061106f61286e565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b61109c611f27565b6110a4611f85565b346000036110de576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166111055760405163d92e233d60e01b815260040160405180910390fd5b610400611115606083018361372e565b9050111561114f576040517f386691aa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040516000916001600160a01b03169034908381818185875af1925050503d806000811461119c576040519150601f19603f3d011682016040523d82523d6000602084013e6111a1565b606091505b50509050806111dc576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c34600086604051611224939291906138ad565b60405180910390a350610d6160017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b61125e611f27565b611266611f85565b346000036112a0576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384166112c75760405163d92e233d60e01b815260040160405180910390fd5b6104006112d7606083018361372e565b6112e29150846138e3565b111561131a576040517f386691aa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040516000916001600160a01b03169034908381818185875af1925050503d8060008114611367576040519150601f19603f3d011682016040523d82523d6000602084013e61136c565b606091505b50509050806113a7576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b846001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c3460008888886040516113f3959493929190613993565b60405180910390a350610a0160017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61144f81611e30565b610d436128d0565b600061146281611e30565b6001600160a01b0382166114895760405163d92e233d60e01b815260040160405180910390fd5b6001546114c0907f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb906001600160a01b0316612200565b506114eb7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83611e3a565b50600154604080516001600160a01b03928316815291841660208301527f3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579910160405180910390a150600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b961159881611e30565b6115a0611f27565b6115a8611f85565b846000036115e2576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0386166116095760405163d92e233d60e01b815260040160405180910390fd5b61161d6001600160a01b0388168787612949565b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0387169063c9028a3690611662908590600401613a36565b600060405180830381600087803b15801561167c57600080fd5b505af1158015611690573d6000803e3d6000fd5b50505050866001600160a01b0316866001600160a01b03167fde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035878787876040516116dd9493929190613a49565b60405180910390a361170e60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050505050565b600061172281611e30565b6001600160a01b0382166117495760405163d92e233d60e01b815260040160405180910390fd5b6000546001600160a01b03161561178c576040517fb337f37800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117b67f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b983611e3a565b5050600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff1660008115801561183d5750825b905060008267ffffffffffffffff16600114801561185a5750303b155b905081158015611868575080155b1561189f576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156119005784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816158061191d57506001600160a01b038716155b1561193b5760405163d92e233d60e01b815260040160405180910390fd5b6119436129bd565b61194b6129c5565b6119536129bd565b61195b6129d5565b611966600087611e3a565b506119917f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87611e3a565b506119bc7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a89611e3a565b50600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038a16179055611a1a7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb89611e3a565b50600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0389161790558315611aaf5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb611ae381611e30565b611aeb611f27565b611af3611f85565b6001600160a01b038516611b1a5760405163d92e233d60e01b815260040160405180910390fd5b6000856001600160a01b03163460405160006040518083038185875af1925050503d8060008114611b67576040519150601f19603f3d011682016040523d82523d6000602084013e611b6c565b606091505b5050905080611ba7576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0387169063c9028a3690611bec908690600401613a36565b600060405180830381600087803b158015611c0657600080fd5b505af1158015611c1a573d6000803e3d6000fd5b5050505060006001600160a01b0316866001600160a01b03167fde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e03534888888604051611c689493929190613a49565b60405180910390a350611c9a60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050505050565b611ca9611f27565b611cb1611f85565b84600003611ceb576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038616611d125760405163d92e233d60e01b815260040160405180910390fd5b610400611d22606083018361372e565b611d2d9150846138e3565b1115611d65576040517f386691aa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d70338587612006565b856001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c8787878787604051611dbb959493929190613993565b60405180910390a361105d60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611e2681611e30565b610a018383612200565b610d4381336129e5565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16611f1d576000848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611ed33390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506107e2565b60009150506107e2565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615611f83576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01612000576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6003546001600160a01b039081169083160361204e576040517fe4dd681d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529091169063d936547e90602401602060405180830381865afa1580156120b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120d5919061395d565b61210b576040517fac2175f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610c20906001600160a01b038481169186911684612a72565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b60606121598383612aab565b600080856001600160a01b0316348686604051612177929190613a80565b60006040518083038185875af1925050503d80600081146121b4576040519150601f19603f3d011682016040523d82523d6000602084013e6121b9565b606091505b5091509150816121f5576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9150505b9392505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615611f1d576000848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506107e2565b6060836001600160a01b031663676cc054348786866040518563ffffffff1660e01b81526004016122f793929190613a90565b60006040518083038185885af1158015612315573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f1916820160405261233e9190810190613abb565b95945050505050565b61234f612bab565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061247057507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166124647f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15611f83576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610d6181611e30565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561250c575060408051601f3d908101601f191682019092526125099181019061397a565b60015b612552576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146125ae576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401612549565b610c208383612c06565b6040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b038281166004830152600060248301819052919084169063095ea7b3906044016020604051808303816000875af1158015612624573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121f9919061395d565b6003546001600160a01b0390811690831603612797576002546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303816000875af11580156126ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126ee919061395d565b612724576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546040517f743e0c9b000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063743e0c9b90602401600060405180830381600087803b15801561278357600080fd5b505af115801561105d573d6000803e3d6000fd5b6000546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529091169063d936547e90602401602060405180830381865afa1580156127fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061281e919061395d565b612854576040517fac2175f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610d61906001600160a01b03848116911683612949565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611f83576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6128d8611f27565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336123b9565b6040516001600160a01b03838116602483015260448201839052610c2091859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612c5c565b611f83612cd8565b6129cd612cd8565b611f83612d3f565b6129dd612cd8565b611f83612d47565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610d61576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401612549565b6040516001600160a01b038481166024830152838116604483015260648201839052610a019186918216906323b872dd90608401612976565b60048110610d615781357f98933fac000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821601612b30576040517fed69977500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f36fd75ca000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821601610c20576040517ff3459a9600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16611f83576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612c0f82612d98565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115612c5457610c208282612e40565b610d61612ead565b6000612c716001600160a01b03841683612ee5565b90508051600014158015612c96575080806020019051810190612c94919061395d565b155b15610c20576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401612549565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16611f83576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612127612cd8565b612d4f612cd8565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b806001600160a01b03163b600003612de7576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401612549565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051612e5d9190613b29565b600060405180830381855af49150503d8060008114612e98576040519150601f19603f3d011682016040523d82523d6000602084013e612e9d565b606091505b509150915061233e858383612ef3565b3415611f83576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606121f983836000612f68565b606082612f0857612f038261301e565b6121f9565b8151158015612f1f57506001600160a01b0384163b155b15612f61576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401612549565b50806121f9565b606081471015612fa6576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401612549565b600080856001600160a01b03168486604051612fc29190613b29565b60006040518083038185875af1925050503d8060008114612fff576040519150601f19603f3d011682016040523d82523d6000602084013e613004565b606091505b5091509150613014868383612ef3565b9695505050505050565b80511561302e5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561307257600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146121f957600080fd5b80356001600160a01b03811681146130b957600080fd5b919050565b6000602082840312156130d057600080fd5b6121f9826130a2565b600060a082840312156130eb57600080fd5b50919050565b6000806000806080858703121561310757600080fd5b613110856130a2565b935060208501359250613125604086016130a2565b9150606085013567ffffffffffffffff81111561314157600080fd5b61314d878288016130d9565b91505092959194509250565b60008083601f84011261316b57600080fd5b50813567ffffffffffffffff81111561318357600080fd5b60208301915083602082850101111561319b57600080fd5b9250929050565b600080600080606085870312156131b857600080fd5b6131c1856130a2565b9350602085013567ffffffffffffffff8111156131dd57600080fd5b6131e987828801613159565b909450925050604085013567ffffffffffffffff81111561314157600080fd5b60008060006040848603121561321e57600080fd5b613227846130a2565b9250602084013567ffffffffffffffff81111561324357600080fd5b61324f86828701613159565b9497909650939450505050565b60005b8381101561327757818101518382015260200161325f565b50506000910152565b6000815180845261329881602086016020860161325c565b601f01601f19169290920160200192915050565b6020815260006121f96020830184613280565b6000602082840312156132d157600080fd5b5035919050565b600080604083850312156132eb57600080fd5b823591506132fb602084016130a2565b90509250929050565b600080600080848603606081121561331b57600080fd5b602081121561332957600080fd5b50849350613339602086016130a2565b9250604085013567ffffffffffffffff81111561335557600080fd5b61336187828801613159565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156133c5576133c561336d565b604052919050565b600067ffffffffffffffff8211156133e7576133e761336d565b50601f01601f191660200190565b6000806040838503121561340857600080fd5b613411836130a2565b9150602083013567ffffffffffffffff81111561342d57600080fd5b8301601f8101851361343e57600080fd5b803561345161344c826133cd565b61339c565b81815286602083850101111561346657600080fd5b816020840160208301376000602083830101528093505050509250929050565b60008060008060006080868803121561349e57600080fd5b6134a7866130a2565b94506134b5602087016130a2565b935060408601359250606086013567ffffffffffffffff8111156134d857600080fd5b6134e488828901613159565b969995985093965092949392505050565b6000806040838503121561350857600080fd5b613511836130a2565b9150602083013567ffffffffffffffff81111561352d57600080fd5b613539858286016130d9565b9150509250929050565b6000608082840312156130eb57600080fd5b60008060008060008060a0878903121561356e57600080fd5b613577876130a2565b9550613585602088016130a2565b945060408701359350606087013567ffffffffffffffff8111156135a857600080fd5b6135b489828a01613159565b909450925050608087013567ffffffffffffffff8111156135d457600080fd5b6135e089828a01613543565b9150509295509295509295565b60008060006060848603121561360257600080fd5b61360b846130a2565b9250613619602085016130a2565b9150613627604085016130a2565b90509250925092565b6000806000806060858703121561364657600080fd5b61364f856130a2565b9350602085013567ffffffffffffffff81111561366b57600080fd5b61367787828801613159565b909450925050604085013567ffffffffffffffff81111561369757600080fd5b61314d87828801613543565b60008060008060008060a087890312156136bc57600080fd5b6136c5876130a2565b9550602087013594506136da604088016130a2565b9350606087013567ffffffffffffffff8111156136f657600080fd5b61370289828a01613159565b909450925050608087013567ffffffffffffffff81111561372257600080fd5b6135e089828a016130d9565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261376357600080fd5b83018035915067ffffffffffffffff82111561377e57600080fd5b60200191503681900382131561319b57600080fd5b8015158114610d4357600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126137d657600080fd5b830160208101925035905067ffffffffffffffff8111156137f657600080fd5b80360382131561319b57600080fd5b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b6001600160a01b03613841826130a2565b1682526000602082013561385481613793565b151560208401526001600160a01b0361386f604084016130a2565b16604084015261388260608301836137a1565b60a0606086015261389760a086018284613805565b6080948501359590940194909452509092915050565b8381526001600160a01b0383166020820152608060408201526000608082015260a06060820152600061233e60a0830184613830565b808201808211156107e2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b604081526000613931604083018587613805565b82810360208401526130148185613830565b83815260406020820152600061233e604083018486613805565b60006020828403121561396f57600080fd5b81516121f981613793565b60006020828403121561398c57600080fd5b5051919050565b8581526001600160a01b03851660208201526080604082015260006139bc608083018587613805565b82810360608401526139ce8185613830565b98975050505050505050565b6001600160a01b036139eb826130a2565b1682526001600160a01b03613a02602083016130a2565b166020830152604081810135908301526000613a2160608301836137a1565b6080606086015261233e608086018284613805565b6020815260006121f960208301846139da565b848152606060208201526000613a63606083018587613805565b8281036040840152613a7581856139da565b979650505050505050565b8183823760009101908152919050565b6001600160a01b03613aa1856130a2565b16815260406020820152600061233e604083018486613805565b600060208284031215613acd57600080fd5b815167ffffffffffffffff811115613ae457600080fd5b8201601f81018413613af557600080fd5b8051613b0361344c826133cd565b818152856020838501011115613b1857600080fd5b61233e82602083016020860161325c565b60008251613b3b81846020870161325c565b919091019291505056fea26469706673582212201229b7b25024e7ca203b587177014870919846d3750c5162ab31d92af37e129664736f6c634300081a0033",
}

// GatewayEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMMetaData.ABI instead.
var GatewayEVMABI = GatewayEVMMetaData.ABI

// GatewayEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMMetaData.Bin instead.
var GatewayEVMBin = GatewayEVMMetaData.Bin

// DeployGatewayEVM deploys a new Ethereum contract, binding an instance of GatewayEVM to it.
func DeployGatewayEVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVM, error) {
	parsed, err := GatewayEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVM{GatewayEVMCaller: GatewayEVMCaller{contract: contract}, GatewayEVMTransactor: GatewayEVMTransactor{contract: contract}, GatewayEVMFilterer: GatewayEVMFilterer{contract: contract}}, nil
}

// GatewayEVM is an auto generated Go binding around an Ethereum contract.
type GatewayEVM struct {
	GatewayEVMCaller     // Read-only binding to the contract
	GatewayEVMTransactor // Write-only binding to the contract
	GatewayEVMFilterer   // Log filterer for contract events
}

// GatewayEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMSession struct {
	Contract     *GatewayEVM       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMCallerSession struct {
	Contract *GatewayEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GatewayEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMTransactorSession struct {
	Contract     *GatewayEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GatewayEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMRaw struct {
	Contract *GatewayEVM // Generic contract binding to access the raw methods on
}

// GatewayEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMCallerRaw struct {
	Contract *GatewayEVMCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMTransactorRaw struct {
	Contract *GatewayEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVM creates a new instance of GatewayEVM, bound to a specific deployed contract.
func NewGatewayEVM(address common.Address, backend bind.ContractBackend) (*GatewayEVM, error) {
	contract, err := bindGatewayEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVM{GatewayEVMCaller: GatewayEVMCaller{contract: contract}, GatewayEVMTransactor: GatewayEVMTransactor{contract: contract}, GatewayEVMFilterer: GatewayEVMFilterer{contract: contract}}, nil
}

// NewGatewayEVMCaller creates a new read-only instance of GatewayEVM, bound to a specific deployed contract.
func NewGatewayEVMCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMCaller, error) {
	contract, err := bindGatewayEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMCaller{contract: contract}, nil
}

// NewGatewayEVMTransactor creates a new write-only instance of GatewayEVM, bound to a specific deployed contract.
func NewGatewayEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMTransactor, error) {
	contract, err := bindGatewayEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTransactor{contract: contract}, nil
}

// NewGatewayEVMFilterer creates a new log filterer instance of GatewayEVM, bound to a specific deployed contract.
func NewGatewayEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMFilterer, error) {
	contract, err := bindGatewayEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMFilterer{contract: contract}, nil
}

// bindGatewayEVM binds a generic wrapper to an already deployed contract.
func bindGatewayEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVM *GatewayEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVM.Contract.GatewayEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVM *GatewayEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVM.Contract.GatewayEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVM *GatewayEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVM.Contract.GatewayEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVM *GatewayEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVM *GatewayEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVM *GatewayEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVM.Contract.contract.Transact(opts, method, params...)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVM *GatewayEVMCaller) ASSETHANDLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "ASSET_HANDLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVM *GatewayEVMSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVM.Contract.ASSETHANDLERROLE(&_GatewayEVM.CallOpts)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVM *GatewayEVMCallerSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVM.Contract.ASSETHANDLERROLE(&_GatewayEVM.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVM *GatewayEVMCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVM *GatewayEVMSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVM.Contract.DEFAULTADMINROLE(&_GatewayEVM.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVM *GatewayEVMCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVM.Contract.DEFAULTADMINROLE(&_GatewayEVM.CallOpts)
}

// MAXPAYLOADSIZE is a free data retrieval call binding the contract method 0xa2ba1934.
//
// Solidity: function MAX_PAYLOAD_SIZE() view returns(uint256)
func (_GatewayEVM *GatewayEVMCaller) MAXPAYLOADSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "MAX_PAYLOAD_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPAYLOADSIZE is a free data retrieval call binding the contract method 0xa2ba1934.
//
// Solidity: function MAX_PAYLOAD_SIZE() view returns(uint256)
func (_GatewayEVM *GatewayEVMSession) MAXPAYLOADSIZE() (*big.Int, error) {
	return _GatewayEVM.Contract.MAXPAYLOADSIZE(&_GatewayEVM.CallOpts)
}

// MAXPAYLOADSIZE is a free data retrieval call binding the contract method 0xa2ba1934.
//
// Solidity: function MAX_PAYLOAD_SIZE() view returns(uint256)
func (_GatewayEVM *GatewayEVMCallerSession) MAXPAYLOADSIZE() (*big.Int, error) {
	return _GatewayEVM.Contract.MAXPAYLOADSIZE(&_GatewayEVM.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVM *GatewayEVMCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVM *GatewayEVMSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVM.Contract.PAUSERROLE(&_GatewayEVM.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVM *GatewayEVMCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVM.Contract.PAUSERROLE(&_GatewayEVM.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVM *GatewayEVMCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVM *GatewayEVMSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVM.Contract.TSSROLE(&_GatewayEVM.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVM *GatewayEVMCallerSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVM.Contract.TSSROLE(&_GatewayEVM.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayEVM *GatewayEVMCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayEVM *GatewayEVMSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayEVM.Contract.UPGRADEINTERFACEVERSION(&_GatewayEVM.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayEVM *GatewayEVMCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayEVM.Contract.UPGRADEINTERFACEVERSION(&_GatewayEVM.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVM *GatewayEVMCaller) Custody(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "custody")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVM *GatewayEVMSession) Custody() (common.Address, error) {
	return _GatewayEVM.Contract.Custody(&_GatewayEVM.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVM *GatewayEVMCallerSession) Custody() (common.Address, error) {
	return _GatewayEVM.Contract.Custody(&_GatewayEVM.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayEVM *GatewayEVMCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayEVM *GatewayEVMSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayEVM.Contract.GetRoleAdmin(&_GatewayEVM.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayEVM *GatewayEVMCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayEVM.Contract.GetRoleAdmin(&_GatewayEVM.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayEVM *GatewayEVMCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayEVM *GatewayEVMSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayEVM.Contract.HasRole(&_GatewayEVM.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayEVM *GatewayEVMCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayEVM.Contract.HasRole(&_GatewayEVM.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayEVM *GatewayEVMCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayEVM *GatewayEVMSession) Paused() (bool, error) {
	return _GatewayEVM.Contract.Paused(&_GatewayEVM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayEVM *GatewayEVMCallerSession) Paused() (bool, error) {
	return _GatewayEVM.Contract.Paused(&_GatewayEVM.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVM *GatewayEVMCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVM *GatewayEVMSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayEVM.Contract.ProxiableUUID(&_GatewayEVM.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVM *GatewayEVMCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayEVM.Contract.ProxiableUUID(&_GatewayEVM.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayEVM *GatewayEVMCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayEVM *GatewayEVMSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayEVM.Contract.SupportsInterface(&_GatewayEVM.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayEVM *GatewayEVMCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayEVM.Contract.SupportsInterface(&_GatewayEVM.CallOpts, interfaceId)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVM *GatewayEVMCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVM *GatewayEVMSession) TssAddress() (common.Address, error) {
	return _GatewayEVM.Contract.TssAddress(&_GatewayEVM.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVM *GatewayEVMCallerSession) TssAddress() (common.Address, error) {
	return _GatewayEVM.Contract.TssAddress(&_GatewayEVM.CallOpts)
}

// ZetaConnector is a free data retrieval call binding the contract method 0x57bec62f.
//
// Solidity: function zetaConnector() view returns(address)
func (_GatewayEVM *GatewayEVMCaller) ZetaConnector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "zetaConnector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaConnector is a free data retrieval call binding the contract method 0x57bec62f.
//
// Solidity: function zetaConnector() view returns(address)
func (_GatewayEVM *GatewayEVMSession) ZetaConnector() (common.Address, error) {
	return _GatewayEVM.Contract.ZetaConnector(&_GatewayEVM.CallOpts)
}

// ZetaConnector is a free data retrieval call binding the contract method 0x57bec62f.
//
// Solidity: function zetaConnector() view returns(address)
func (_GatewayEVM *GatewayEVMCallerSession) ZetaConnector() (common.Address, error) {
	return _GatewayEVM.Contract.ZetaConnector(&_GatewayEVM.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayEVM *GatewayEVMCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVM.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayEVM *GatewayEVMSession) ZetaToken() (common.Address, error) {
	return _GatewayEVM.Contract.ZetaToken(&_GatewayEVM.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayEVM *GatewayEVMCallerSession) ZetaToken() (common.Address, error) {
	return _GatewayEVM.Contract.ZetaToken(&_GatewayEVM.CallOpts)
}

// Call is a paid mutator transaction binding the contract method 0x1becceb4.
//
// Solidity: function call(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVM *GatewayEVMTransactor) Call(opts *bind.TransactOpts, receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "call", receiver, payload, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x1becceb4.
//
// Solidity: function call(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVM *GatewayEVMSession) Call(receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Call(&_GatewayEVM.TransactOpts, receiver, payload, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x1becceb4.
//
// Solidity: function call(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) Call(receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Call(&_GatewayEVM.TransactOpts, receiver, payload, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0x102614b0.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVM *GatewayEVMTransactor) Deposit(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, asset common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "deposit", receiver, amount, asset, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0x102614b0.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVM *GatewayEVMSession) Deposit(receiver common.Address, amount *big.Int, asset common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Deposit(&_GatewayEVM.TransactOpts, receiver, amount, asset, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0x102614b0.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) Deposit(receiver common.Address, amount *big.Int, asset common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Deposit(&_GatewayEVM.TransactOpts, receiver, amount, asset, revertOptions)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x726ac97c.
//
// Solidity: function deposit(address receiver, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVM *GatewayEVMTransactor) Deposit0(opts *bind.TransactOpts, receiver common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "deposit0", receiver, revertOptions)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x726ac97c.
//
// Solidity: function deposit(address receiver, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVM *GatewayEVMSession) Deposit0(receiver common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Deposit0(&_GatewayEVM.TransactOpts, receiver, revertOptions)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x726ac97c.
//
// Solidity: function deposit(address receiver, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVM *GatewayEVMTransactorSession) Deposit0(receiver common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Deposit0(&_GatewayEVM.TransactOpts, receiver, revertOptions)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x744b9b8b.
//
// Solidity: function depositAndCall(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVM *GatewayEVMTransactor) DepositAndCall(opts *bind.TransactOpts, receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "depositAndCall", receiver, payload, revertOptions)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x744b9b8b.
//
// Solidity: function depositAndCall(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVM *GatewayEVMSession) DepositAndCall(receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.Contract.DepositAndCall(&_GatewayEVM.TransactOpts, receiver, payload, revertOptions)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x744b9b8b.
//
// Solidity: function depositAndCall(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVM *GatewayEVMTransactorSession) DepositAndCall(receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.Contract.DepositAndCall(&_GatewayEVM.TransactOpts, receiver, payload, revertOptions)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xd09e3b78.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVM *GatewayEVMTransactor) DepositAndCall0(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, asset common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "depositAndCall0", receiver, amount, asset, payload, revertOptions)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xd09e3b78.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVM *GatewayEVMSession) DepositAndCall0(receiver common.Address, amount *big.Int, asset common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.Contract.DepositAndCall0(&_GatewayEVM.TransactOpts, receiver, amount, asset, payload, revertOptions)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xd09e3b78.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) DepositAndCall0(receiver common.Address, amount *big.Int, asset common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVM.Contract.DepositAndCall0(&_GatewayEVM.TransactOpts, receiver, amount, asset, payload, revertOptions)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVM *GatewayEVMTransactor) Execute(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "execute", destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVM *GatewayEVMSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Execute(&_GatewayEVM.TransactOpts, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVM *GatewayEVMTransactorSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Execute(&_GatewayEVM.TransactOpts, destination, data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x38e22527.
//
// Solidity: function execute((address) messageContext, address destination, bytes data) payable returns(bytes)
func (_GatewayEVM *GatewayEVMTransactor) Execute0(opts *bind.TransactOpts, messageContext MessageContext, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "execute0", messageContext, destination, data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x38e22527.
//
// Solidity: function execute((address) messageContext, address destination, bytes data) payable returns(bytes)
func (_GatewayEVM *GatewayEVMSession) Execute0(messageContext MessageContext, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Execute0(&_GatewayEVM.TransactOpts, messageContext, destination, data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x38e22527.
//
// Solidity: function execute((address) messageContext, address destination, bytes data) payable returns(bytes)
func (_GatewayEVM *GatewayEVMTransactorSession) Execute0(messageContext MessageContext, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Execute0(&_GatewayEVM.TransactOpts, messageContext, destination, data)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0xcb7ba8e5.
//
// Solidity: function executeRevert(address destination, bytes data, (address,address,uint256,bytes) revertContext) payable returns()
func (_GatewayEVM *GatewayEVMTransactor) ExecuteRevert(opts *bind.TransactOpts, destination common.Address, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "executeRevert", destination, data, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0xcb7ba8e5.
//
// Solidity: function executeRevert(address destination, bytes data, (address,address,uint256,bytes) revertContext) payable returns()
func (_GatewayEVM *GatewayEVMSession) ExecuteRevert(destination common.Address, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVM.Contract.ExecuteRevert(&_GatewayEVM.TransactOpts, destination, data, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0xcb7ba8e5.
//
// Solidity: function executeRevert(address destination, bytes data, (address,address,uint256,bytes) revertContext) payable returns()
func (_GatewayEVM *GatewayEVMTransactorSession) ExecuteRevert(destination common.Address, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVM.Contract.ExecuteRevert(&_GatewayEVM.TransactOpts, destination, data, revertContext)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVM *GatewayEVMTransactor) ExecuteWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "executeWithERC20", token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVM *GatewayEVMSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.ExecuteWithERC20(&_GatewayEVM.TransactOpts, token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.ExecuteWithERC20(&_GatewayEVM.TransactOpts, token, to, amount, data)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayEVM *GatewayEVMTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayEVM *GatewayEVMSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.GrantRole(&_GatewayEVM.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.GrantRole(&_GatewayEVM.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address tssAddress_, address zetaToken_, address admin_) returns()
func (_GatewayEVM *GatewayEVMTransactor) Initialize(opts *bind.TransactOpts, tssAddress_ common.Address, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "initialize", tssAddress_, zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address tssAddress_, address zetaToken_, address admin_) returns()
func (_GatewayEVM *GatewayEVMSession) Initialize(tssAddress_ common.Address, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Initialize(&_GatewayEVM.TransactOpts, tssAddress_, zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address tssAddress_, address zetaToken_, address admin_) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) Initialize(tssAddress_ common.Address, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.Initialize(&_GatewayEVM.TransactOpts, tssAddress_, zetaToken_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayEVM *GatewayEVMTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayEVM *GatewayEVMSession) Pause() (*types.Transaction, error) {
	return _GatewayEVM.Contract.Pause(&_GatewayEVM.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayEVM *GatewayEVMTransactorSession) Pause() (*types.Transaction, error) {
	return _GatewayEVM.Contract.Pause(&_GatewayEVM.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayEVM *GatewayEVMTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayEVM *GatewayEVMSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.RenounceRole(&_GatewayEVM.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.RenounceRole(&_GatewayEVM.TransactOpts, role, callerConfirmation)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xaa0c0fc1.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayEVM *GatewayEVMTransactor) RevertWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "revertWithERC20", token, to, amount, data, revertContext)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xaa0c0fc1.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayEVM *GatewayEVMSession) RevertWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVM.Contract.RevertWithERC20(&_GatewayEVM.TransactOpts, token, to, amount, data, revertContext)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xaa0c0fc1.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) RevertWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVM.Contract.RevertWithERC20(&_GatewayEVM.TransactOpts, token, to, amount, data, revertContext)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayEVM *GatewayEVMTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayEVM *GatewayEVMSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.RevokeRole(&_GatewayEVM.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.RevokeRole(&_GatewayEVM.TransactOpts, role, account)
}

// SetConnector is a paid mutator transaction binding the contract method 0x10188aef.
//
// Solidity: function setConnector(address zetaConnector_) returns()
func (_GatewayEVM *GatewayEVMTransactor) SetConnector(opts *bind.TransactOpts, zetaConnector_ common.Address) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "setConnector", zetaConnector_)
}

// SetConnector is a paid mutator transaction binding the contract method 0x10188aef.
//
// Solidity: function setConnector(address zetaConnector_) returns()
func (_GatewayEVM *GatewayEVMSession) SetConnector(zetaConnector_ common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.SetConnector(&_GatewayEVM.TransactOpts, zetaConnector_)
}

// SetConnector is a paid mutator transaction binding the contract method 0x10188aef.
//
// Solidity: function setConnector(address zetaConnector_) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) SetConnector(zetaConnector_ common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.SetConnector(&_GatewayEVM.TransactOpts, zetaConnector_)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address custody_) returns()
func (_GatewayEVM *GatewayEVMTransactor) SetCustody(opts *bind.TransactOpts, custody_ common.Address) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "setCustody", custody_)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address custody_) returns()
func (_GatewayEVM *GatewayEVMSession) SetCustody(custody_ common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.SetCustody(&_GatewayEVM.TransactOpts, custody_)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address custody_) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) SetCustody(custody_ common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.SetCustody(&_GatewayEVM.TransactOpts, custody_)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayEVM *GatewayEVMTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayEVM *GatewayEVMSession) Unpause() (*types.Transaction, error) {
	return _GatewayEVM.Contract.Unpause(&_GatewayEVM.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayEVM *GatewayEVMTransactorSession) Unpause() (*types.Transaction, error) {
	return _GatewayEVM.Contract.Unpause(&_GatewayEVM.TransactOpts)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_GatewayEVM *GatewayEVMTransactor) UpdateTSSAddress(opts *bind.TransactOpts, newTSSAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "updateTSSAddress", newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_GatewayEVM *GatewayEVMSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.UpdateTSSAddress(&_GatewayEVM.TransactOpts, newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_GatewayEVM *GatewayEVMTransactorSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVM.Contract.UpdateTSSAddress(&_GatewayEVM.TransactOpts, newTSSAddress)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVM *GatewayEVMTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVM *GatewayEVMSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.UpgradeToAndCall(&_GatewayEVM.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVM *GatewayEVMTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVM.Contract.UpgradeToAndCall(&_GatewayEVM.TransactOpts, newImplementation, data)
}

// GatewayEVMCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayEVM contract.
type GatewayEVMCalledIterator struct {
	Event *GatewayEVMCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMCalled)
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
		it.Event = new(GatewayEVMCalled)
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
func (it *GatewayEVMCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMCalled represents a Called event raised by the GatewayEVM contract.
type GatewayEVMCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVM *GatewayEVMFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMCalledIterator{contract: _GatewayEVM.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVM *GatewayEVMFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMCalled)
				if err := _GatewayEVM.contract.UnpackLog(event, "Called", log); err != nil {
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

// ParseCalled is a log parse operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVM *GatewayEVMFilterer) ParseCalled(log types.Log) (*GatewayEVMCalled, error) {
	event := new(GatewayEVMCalled)
	if err := _GatewayEVM.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayEVM contract.
type GatewayEVMDepositedIterator struct {
	Event *GatewayEVMDeposited // Event containing the contract specifics and raw log

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
func (it *GatewayEVMDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMDeposited)
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
		it.Event = new(GatewayEVMDeposited)
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
func (it *GatewayEVMDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMDeposited represents a Deposited event raised by the GatewayEVM contract.
type GatewayEVMDeposited struct {
	Sender        common.Address
	Receiver      common.Address
	Amount        *big.Int
	Asset         common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVM *GatewayEVMFilterer) FilterDeposited(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMDepositedIterator{contract: _GatewayEVM.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVM *GatewayEVMFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayEVMDeposited, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMDeposited)
				if err := _GatewayEVM.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVM *GatewayEVMFilterer) ParseDeposited(log types.Log) (*GatewayEVMDeposited, error) {
	event := new(GatewayEVMDeposited)
	if err := _GatewayEVM.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVM contract.
type GatewayEVMExecutedIterator struct {
	Event *GatewayEVMExecuted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMExecuted)
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
		it.Event = new(GatewayEVMExecuted)
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
func (it *GatewayEVMExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMExecuted represents a Executed event raised by the GatewayEVM contract.
type GatewayEVMExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVM *GatewayEVMFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMExecutedIterator{contract: _GatewayEVM.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVM *GatewayEVMFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMExecuted)
				if err := _GatewayEVM.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVM *GatewayEVMFilterer) ParseExecuted(log types.Log) (*GatewayEVMExecuted, error) {
	event := new(GatewayEVMExecuted)
	if err := _GatewayEVM.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVM contract.
type GatewayEVMExecutedWithERC20Iterator struct {
	Event *GatewayEVMExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMExecutedWithERC20)
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
		it.Event = new(GatewayEVMExecutedWithERC20)
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
func (it *GatewayEVMExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVM contract.
type GatewayEVMExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVM *GatewayEVMFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMExecutedWithERC20Iterator{contract: _GatewayEVM.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVM *GatewayEVMFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMExecutedWithERC20)
				if err := _GatewayEVM.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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

// ParseExecutedWithERC20 is a log parse operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVM *GatewayEVMFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMExecutedWithERC20, error) {
	event := new(GatewayEVMExecutedWithERC20)
	if err := _GatewayEVM.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayEVM contract.
type GatewayEVMInitializedIterator struct {
	Event *GatewayEVMInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInitialized)
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
		it.Event = new(GatewayEVMInitialized)
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
func (it *GatewayEVMInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInitialized represents a Initialized event raised by the GatewayEVM contract.
type GatewayEVMInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayEVM *GatewayEVMFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayEVMInitializedIterator, error) {

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInitializedIterator{contract: _GatewayEVM.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayEVM *GatewayEVMFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayEVMInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInitialized)
				if err := _GatewayEVM.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayEVM *GatewayEVMFilterer) ParseInitialized(log types.Log) (*GatewayEVMInitialized, error) {
	event := new(GatewayEVMInitialized)
	if err := _GatewayEVM.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GatewayEVM contract.
type GatewayEVMPausedIterator struct {
	Event *GatewayEVMPaused // Event containing the contract specifics and raw log

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
func (it *GatewayEVMPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMPaused)
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
		it.Event = new(GatewayEVMPaused)
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
func (it *GatewayEVMPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMPaused represents a Paused event raised by the GatewayEVM contract.
type GatewayEVMPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayEVM *GatewayEVMFilterer) FilterPaused(opts *bind.FilterOpts) (*GatewayEVMPausedIterator, error) {

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMPausedIterator{contract: _GatewayEVM.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayEVM *GatewayEVMFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GatewayEVMPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMPaused)
				if err := _GatewayEVM.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_GatewayEVM *GatewayEVMFilterer) ParsePaused(log types.Log) (*GatewayEVMPaused, error) {
	event := new(GatewayEVMPaused)
	if err := _GatewayEVM.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVM contract.
type GatewayEVMRevertedIterator struct {
	Event *GatewayEVMReverted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMReverted)
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
		it.Event = new(GatewayEVMReverted)
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
func (it *GatewayEVMRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMReverted represents a Reverted event raised by the GatewayEVM contract.
type GatewayEVMReverted struct {
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVM *GatewayEVMFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMRevertedIterator{contract: _GatewayEVM.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVM *GatewayEVMFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMReverted)
				if err := _GatewayEVM.contract.UnpackLog(event, "Reverted", log); err != nil {
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

// ParseReverted is a log parse operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVM *GatewayEVMFilterer) ParseReverted(log types.Log) (*GatewayEVMReverted, error) {
	event := new(GatewayEVMReverted)
	if err := _GatewayEVM.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GatewayEVM contract.
type GatewayEVMRoleAdminChangedIterator struct {
	Event *GatewayEVMRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayEVMRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMRoleAdminChanged)
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
		it.Event = new(GatewayEVMRoleAdminChanged)
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
func (it *GatewayEVMRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMRoleAdminChanged represents a RoleAdminChanged event raised by the GatewayEVM contract.
type GatewayEVMRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayEVM *GatewayEVMFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GatewayEVMRoleAdminChangedIterator, error) {

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

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMRoleAdminChangedIterator{contract: _GatewayEVM.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayEVM *GatewayEVMFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayEVMRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMRoleAdminChanged)
				if err := _GatewayEVM.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_GatewayEVM *GatewayEVMFilterer) ParseRoleAdminChanged(log types.Log) (*GatewayEVMRoleAdminChanged, error) {
	event := new(GatewayEVMRoleAdminChanged)
	if err := _GatewayEVM.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GatewayEVM contract.
type GatewayEVMRoleGrantedIterator struct {
	Event *GatewayEVMRoleGranted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMRoleGranted)
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
		it.Event = new(GatewayEVMRoleGranted)
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
func (it *GatewayEVMRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMRoleGranted represents a RoleGranted event raised by the GatewayEVM contract.
type GatewayEVMRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVM *GatewayEVMFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayEVMRoleGrantedIterator, error) {

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

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMRoleGrantedIterator{contract: _GatewayEVM.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVM *GatewayEVMFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GatewayEVMRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMRoleGranted)
				if err := _GatewayEVM.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_GatewayEVM *GatewayEVMFilterer) ParseRoleGranted(log types.Log) (*GatewayEVMRoleGranted, error) {
	event := new(GatewayEVMRoleGranted)
	if err := _GatewayEVM.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GatewayEVM contract.
type GatewayEVMRoleRevokedIterator struct {
	Event *GatewayEVMRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GatewayEVMRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMRoleRevoked)
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
		it.Event = new(GatewayEVMRoleRevoked)
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
func (it *GatewayEVMRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMRoleRevoked represents a RoleRevoked event raised by the GatewayEVM contract.
type GatewayEVMRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVM *GatewayEVMFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayEVMRoleRevokedIterator, error) {

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

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMRoleRevokedIterator{contract: _GatewayEVM.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVM *GatewayEVMFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GatewayEVMRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMRoleRevoked)
				if err := _GatewayEVM.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_GatewayEVM *GatewayEVMFilterer) ParseRoleRevoked(log types.Log) (*GatewayEVMRoleRevoked, error) {
	event := new(GatewayEVMRoleRevoked)
	if err := _GatewayEVM.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GatewayEVM contract.
type GatewayEVMUnpausedIterator struct {
	Event *GatewayEVMUnpaused // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUnpaused)
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
		it.Event = new(GatewayEVMUnpaused)
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
func (it *GatewayEVMUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUnpaused represents a Unpaused event raised by the GatewayEVM contract.
type GatewayEVMUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayEVM *GatewayEVMFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GatewayEVMUnpausedIterator, error) {

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUnpausedIterator{contract: _GatewayEVM.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayEVM *GatewayEVMFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GatewayEVMUnpaused) (event.Subscription, error) {

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUnpaused)
				if err := _GatewayEVM.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_GatewayEVM *GatewayEVMFilterer) ParseUnpaused(log types.Log) (*GatewayEVMUnpaused, error) {
	event := new(GatewayEVMUnpaused)
	if err := _GatewayEVM.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the GatewayEVM contract.
type GatewayEVMUpdatedGatewayTSSAddressIterator struct {
	Event *GatewayEVMUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpdatedGatewayTSSAddress)
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
		it.Event = new(GatewayEVMUpdatedGatewayTSSAddress)
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
func (it *GatewayEVMUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the GatewayEVM contract.
type GatewayEVMUpdatedGatewayTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVM *GatewayEVMFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*GatewayEVMUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpdatedGatewayTSSAddressIterator{contract: _GatewayEVM.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVM *GatewayEVMFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpdatedGatewayTSSAddress)
				if err := _GatewayEVM.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
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

// ParseUpdatedGatewayTSSAddress is a log parse operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVM *GatewayEVMFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*GatewayEVMUpdatedGatewayTSSAddress, error) {
	event := new(GatewayEVMUpdatedGatewayTSSAddress)
	if err := _GatewayEVM.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayEVM contract.
type GatewayEVMUpgradedIterator struct {
	Event *GatewayEVMUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayEVMUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgraded)
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
		it.Event = new(GatewayEVMUpgraded)
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
func (it *GatewayEVMUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgraded represents a Upgraded event raised by the GatewayEVM contract.
type GatewayEVMUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVM *GatewayEVMFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayEVMUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayEVM.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradedIterator{contract: _GatewayEVM.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVM *GatewayEVMFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayEVM.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgraded)
				if err := _GatewayEVM.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GatewayEVM *GatewayEVMFilterer) ParseUpgraded(log types.Log) (*GatewayEVMUpgraded, error) {
	event := new(GatewayEVMUpgraded)
	if err := _GatewayEVM.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
