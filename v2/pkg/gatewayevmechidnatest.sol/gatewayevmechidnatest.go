// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayevmechidnatest

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
	RevertAddress    common.Address
	CallOnRevert     bool
	AbortAddress     common.Address
	RevertMessage    []byte
	OnRevertGasLimit *big.Int
}

// GatewayEVMEchidnaTestMetaData contains all meta data concerning the GatewayEVMEchidnaTest contract.
var GatewayEVMEchidnaTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ASSET_HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"custody\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"echidnaCaller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revertWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConnector\",\"inputs\":[{\"name\":\"zetaConnector_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustody\",\"inputs\":[{\"name\":\"custody_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testERC20\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractTestERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testExecuteWithERC20\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"zetaConnector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a060405230608052600580546001600160a01b0319163317905534801561002657600080fd5b5061002f610154565b600554600180546001600160a01b039092166001600160a01b031992831617905560028054610123921691909117905560405161006b90610206565b60408082526004908201819052631d195cdd60e21b606083015260806020830181905282015263151154d560e21b60a082015260c001604051809103906000f0801580156100bd573d6000803e3d6000fd5b50600480546001600160a01b0319166001600160a01b039283161790556001546040513092919091169082906100f290610213565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f08015801561012e573d6000803e3d6000fd5b50600080546001600160a01b0319166001600160a01b0392909216919091179055610220565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156101a45760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146102035780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b610c9f80613c3e83390190565b611dff806148dd83390190565b6080516139f5610249600039600081816123710152818161239a015261282601526139f56000f3fe60806040526004361061024f5760003560e01c80636ab90f9b11610138578063ad3cb1cc116100b0578063d0b492c31161007f578063dda79b7511610064578063dda79b751461073a578063e63ab1e91461075a578063f7ad60db1461078e57600080fd5b8063d0b492c3146106fa578063d547741f1461071a57600080fd5b8063ad3cb1cc14610651578063ae7a3a6f1461069a578063c0c53b8b146106ba578063d09e3b78146106da57600080fd5b80638456cb5911610107578063950837aa116100ec578063950837aa146105e8578063a217fddf14610608578063a783c7891461061d57600080fd5b80638456cb591461056e57806391d148541461058357600080fd5b80636ab90f9b14610508578063726ac97c14610528578063744b9b8b1461053b57806381100bf01461054e57600080fd5b80633c2f05a8116101cb57806352d1902d1161019a5780635b1125911161017f5780635b1125911461047d5780635c975abb1461049d5780635d62c860146104d457600080fd5b806352d1902d1461044857806357bec62f1461045d57600080fd5b80633c2f05a8146103e05780633f4ba83a146104005780634f1ef286146104155780635131ab591461042857600080fd5b80631cff79cd11610222578063248a9ca311610207578063248a9ca3146103435780632f2ff15d146103a057806336568abe146103c057600080fd5b80631cff79cd146102eb57806321e093b11461030b57600080fd5b806301ffc9a71461025457806310188aef14610289578063102614b0146102ab5780631becceb4146102cb575b600080fd5b34801561026057600080fd5b5061027461026f366004612f9b565b6107a1565b60405190151581526020015b60405180910390f35b34801561029557600080fd5b506102a96102a4366004612ff9565b61083a565b005b3480156102b757600080fd5b506102a96102c636600461302c565b610915565b3480156102d757600080fd5b506102a96102e63660046130dd565b610a0f565b6102fe6102f9366004613144565b610a8d565b6040516102809190613205565b34801561031757600080fd5b5060035461032b906001600160a01b031681565b6040516001600160a01b039091168152602001610280565b34801561034f57600080fd5b5061039261035e366004613218565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b604051908152602001610280565b3480156103ac57600080fd5b506102a96103bb366004613231565b610b77565b3480156103cc57600080fd5b506102a96103db366004613231565b610bbb565b3480156103ec57600080fd5b5060045461032b906001600160a01b031681565b34801561040c57600080fd5b506102a9610c0c565b6102a961042336600461328c565b610c41565b34801561043457600080fd5b506102a9610443366004613393565b610c60565b34801561045457600080fd5b50610392610f60565b34801561046957600080fd5b5060025461032b906001600160a01b031681565b34801561048957600080fd5b5060015461032b906001600160a01b031681565b3480156104a957600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610274565b3480156104e057600080fd5b506103927f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b981565b34801561051457600080fd5b506102a9610523366004613402565b610f8f565b6102a961053636600461345c565b6110ba565b6102a96105493660046130dd565b611232565b34801561055a57600080fd5b5060055461032b906001600160a01b031681565b34801561057a57600080fd5b506102a96113ae565b34801561058f57600080fd5b5061027461059e366004613231565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156105f457600080fd5b506102a9610603366004612ff9565b6113e0565b34801561061457600080fd5b50610392600081565b34801561062957600080fd5b506103927f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b34801561065d57600080fd5b506102fe6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b3480156106a657600080fd5b506102a96106b5366004612ff9565b6114e2565b3480156106c657600080fd5b506102a96106d53660046134aa565b6115bd565b3480156106e657600080fd5b506102a96106f53660046134ed565b611859565b34801561070657600080fd5b506102a9610715366004613597565b611951565b34801561072657600080fd5b506102a9610735366004613231565b611afa565b34801561074657600080fd5b5060005461032b906001600160a01b031681565b34801561076657600080fd5b506103927f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6102a961079c366004613622565b611b3e565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061083457507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b600061084581611d26565b6001600160a01b03821661086c5760405163d92e233d60e01b815260040160405180910390fd5b6002546001600160a01b0316156108af576040517f0c8dc01600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108d97f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b983611d30565b5050600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b61091d611e1d565b610925611e7b565b8260000361095f576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384166109865760405163d92e233d60e01b815260040160405180910390fd5b610991338385611efc565b836001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c8585856040516109d8939291906137cd565b60405180910390a3610a0960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050565b610a17611e1d565b610a1f611e7b565b6001600160a01b038416610a465760405163d92e233d60e01b815260040160405180910390fd5b836001600160a01b0316336001600160a01b03167fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d9748585856040516109d893929190613803565b60607f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb610ab981611d26565b610ac1611e1d565b610ac9611e7b565b6001600160a01b038516610af05760405163d92e233d60e01b815260040160405180910390fd5b6000610afd86868661215f565b9050856001600160a01b03167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f348787604051610b3c93929190613829565b60405180910390a29150610b6f60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b509392505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610bb181611d26565b610a098383611d30565b6001600160a01b0381163314610bfd576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c078282612212565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610c3681611d26565b610c3e6122d6565b50565b610c49612366565b610c5282612436565b610c5c8282612441565b5050565b7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b9610c8a81611d26565b610c92611e1d565b610c9a611e7b565b83600003610cd4576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038516610cfb5760405163d92e233d60e01b815260040160405180910390fd5b610d058686612565565b610d3b576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0386811660048301526024820186905287169063095ea7b3906044016020604051808303816000875af1158015610da3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc79190613843565b610dfd576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e0885848461215f565b50610e138686612565565b610e49576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038816906370a0823190602401602060405180830381865afa158015610ea9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ecd9190613860565b90508015610edf57610edf87826125f5565b856001600160a01b0316876001600160a01b03167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382878787604051610f2693929190613829565b60405180910390a350610f5860017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b6000610f6a61281b565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b600480546040517f40c10f190000000000000000000000000000000000000000000000000000000081523092810192909252602482018590526001600160a01b0316906340c10f1990604401600060405180830381600087803b158015610ff557600080fd5b505af1158015611009573d6000803e3d6000fd5b505060045461102792506001600160a01b0316905085858585610c60565b600480546040517f70a0823100000000000000000000000000000000000000000000000000000000815230928101929092526001600160a01b0316906370a0823190602401602060405180830381865afa158015611089573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110ad9190613860565b15610a0957610a09613879565b6110c2611e1d565b6110ca611e7b565b34600003611104576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03821661112b5760405163d92e233d60e01b815260040160405180910390fd5b6001546040516000916001600160a01b03169034908381818185875af1925050503d8060008114611178576040519150601f19603f3d011682016040523d82523d6000602084013e61117d565b606091505b50509050806111b8576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c34600086604051611200939291906137cd565b60405180910390a350610c5c60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b61123a611e1d565b611242611e7b565b3460000361127c576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384166112a35760405163d92e233d60e01b815260040160405180910390fd5b6001546040516000916001600160a01b03169034908381818185875af1925050503d80600081146112f0576040519150601f19603f3d011682016040523d82523d6000602084013e6112f5565b606091505b5050905080611330576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b846001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c34600088888860405161137c9594939291906138a8565b60405180910390a350610a0960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6113d881611d26565b610c3e61287d565b60006113eb81611d26565b6001600160a01b0382166114125760405163d92e233d60e01b815260040160405180910390fd5b600154611449907f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb906001600160a01b0316612212565b506114747f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83611d30565b50600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040519081527f7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a059060200160405180910390a15050565b60006114ed81611d26565b6001600160a01b0382166115145760405163d92e233d60e01b815260040160405180910390fd5b6000546001600160a01b031615611557576040517fb337f37800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115817f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b983611d30565b5050600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156116085750825b905060008267ffffffffffffffff1660011480156116255750303b155b905081158015611633575080155b1561166a576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156116cb5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b03881615806116e857506001600160a01b038716155b156117065760405163d92e233d60e01b815260040160405180910390fd5b61170e6128f6565b6117166128fe565b61171e6128f6565b61172661290e565b611731600087611d30565b5061175c7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87611d30565b50600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038a161790556117ba7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb89611d30565b50600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038916179055831561184f5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b611861611e1d565b611869611e7b565b846000036118a3576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0386166118ca5760405163d92e233d60e01b815260040160405180910390fd5b6118d5338587611efc565b856001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c87878787876040516119209594939291906138a8565b60405180910390a3610f5860017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b961197b81611d26565b611983611e1d565b61198b611e7b565b846000036119c5576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0386166119ec5760405163d92e233d60e01b815260040160405180910390fd5b611a006001600160a01b038816878761291e565b6040517f660b9de00000000000000000000000000000000000000000000000000000000081526001600160a01b0387169063660b9de090611a45908590600401613949565b600060405180830381600087803b158015611a5f57600080fd5b505af1158015611a73573d6000803e3d6000fd5b50505050866001600160a01b0316866001600160a01b03167f1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b143687878787604051611ac0949392919061395c565b60405180910390a3611af160017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611b3481611d26565b610a098383612212565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb611b6881611d26565b611b70611e1d565b611b78611e7b565b6001600160a01b038516611b9f5760405163d92e233d60e01b815260040160405180910390fd5b6000856001600160a01b03163460405160006040518083038185875af1925050503d8060008114611bec576040519150601f19603f3d011682016040523d82523d6000602084013e611bf1565b606091505b5050905080611c2c576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f660b9de00000000000000000000000000000000000000000000000000000000081526001600160a01b0387169063660b9de090611c71908690600401613949565b600060405180830381600087803b158015611c8b57600080fd5b505af1158015611c9f573d6000803e3d6000fd5b5050505060006001600160a01b0316866001600160a01b03167f1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b143634888888604051611ced949392919061395c565b60405180910390a350611d1f60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050505050565b610c3e8133612992565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16611e13576000848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611dc93390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610834565b6000915050610834565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615611e79576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611ef6576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6003546001600160a01b039081169083160361206057611f276001600160a01b038316843084612a1f565b6002546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303816000875af1158015611f93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fb79190613843565b611fed576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546040517f743e0c9b000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063743e0c9b90602401600060405180830381600087803b15801561204c57600080fd5b505af1158015611af1573d6000803e3d6000fd5b6000546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529091169063d936547e90602401602060405180830381865afa1580156120c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120e79190613843565b61211d576040517fac2175f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610c07906001600160a01b038481169186911684612a1f565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b606061216b8383612a58565b600080856001600160a01b0316348686604051612189929190613993565b60006040518083038185875af1925050503d80600081146121c6576040519150601f19603f3d011682016040523d82523d6000602084013e6121cb565b606091505b509150915081612207576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9150505b9392505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615611e13576000848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610834565b6122de612add565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806123ff57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166123f37f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15611e79576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610c5c81611d26565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156124b9575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526124b691810190613860565b60015b6124ff576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc811461255b576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016124f6565b610c078383612b38565b6040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b038281166004830152600060248301819052919084169063095ea7b3906044016020604051808303816000875af11580156125d1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061220b9190613843565b6003546001600160a01b0390811690831603612744576002546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303816000875af1158015612677573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061269b9190613843565b6126d1576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546040517f743e0c9b000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063743e0c9b90602401600060405180830381600087803b15801561273057600080fd5b505af1158015610f58573d6000803e3d6000fd5b6000546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529091169063d936547e90602401602060405180830381865afa1580156127a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127cb9190613843565b612801576040517fac2175f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610c5c906001600160a01b0384811691168361291e565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611e79576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612885611e1d565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612348565b611e79612b8e565b612906612b8e565b611e79612bf5565b612916612b8e565b611e79612bfd565b6040516001600160a01b03838116602483015260448201839052610c0791859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612c4e565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610c5c576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044016124f6565b6040516001600160a01b038481166024830152838116604483015260648201839052610a099186918216906323b872dd9060840161294b565b60048110610c5c5781357f99f46220000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821601610c07576040517ff3459a9600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16611e79576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612b4182612cca565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115612b8657610c078282612d72565b610c5c612de8565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16611e79576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612139612b8e565b612c05612b8e565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6000612c636001600160a01b03841683612e20565b90508051600014158015612c88575080806020019051810190612c869190613843565b155b15610c07576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201526024016124f6565b806001600160a01b03163b600003612d19576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016124f6565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051612d8f91906139a3565b600060405180830381855af49150503d8060008114612dca576040519150601f19603f3d011682016040523d82523d6000602084013e612dcf565b606091505b5091509150612ddf858383612e2e565b95945050505050565b3415611e79576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606061220b83836000612ea3565b606082612e4357612e3e82612f59565b61220b565b8151158015612e5a57506001600160a01b0384163b155b15612e9c576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016124f6565b508061220b565b606081471015612ee1576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016124f6565b600080856001600160a01b03168486604051612efd91906139a3565b60006040518083038185875af1925050503d8060008114612f3a576040519150601f19603f3d011682016040523d82523d6000602084013e612f3f565b606091505b5091509150612f4f868383612e2e565b9695505050505050565b805115612f695780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215612fad57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461220b57600080fd5b80356001600160a01b0381168114612ff457600080fd5b919050565b60006020828403121561300b57600080fd5b61220b82612fdd565b600060a0828403121561302657600080fd5b50919050565b6000806000806080858703121561304257600080fd5b61304b85612fdd565b93506020850135925061306060408601612fdd565b9150606085013567ffffffffffffffff81111561307c57600080fd5b61308887828801613014565b91505092959194509250565b60008083601f8401126130a657600080fd5b50813567ffffffffffffffff8111156130be57600080fd5b6020830191508360208285010111156130d657600080fd5b9250929050565b600080600080606085870312156130f357600080fd5b6130fc85612fdd565b9350602085013567ffffffffffffffff81111561311857600080fd5b61312487828801613094565b909450925050604085013567ffffffffffffffff81111561307c57600080fd5b60008060006040848603121561315957600080fd5b61316284612fdd565b9250602084013567ffffffffffffffff81111561317e57600080fd5b61318a86828701613094565b9497909650939450505050565b60005b838110156131b257818101518382015260200161319a565b50506000910152565b600081518084526131d3816020860160208601613197565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061220b60208301846131bb565b60006020828403121561322a57600080fd5b5035919050565b6000806040838503121561324457600080fd5b8235915061325460208401612fdd565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806040838503121561329f57600080fd5b6132a883612fdd565b9150602083013567ffffffffffffffff8111156132c457600080fd5b8301601f810185136132d557600080fd5b803567ffffffffffffffff8111156132ef576132ef61325d565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561335b5761335b61325d565b60405281815282820160200187101561337357600080fd5b816020840160208301376000602083830101528093505050509250929050565b6000806000806000608086880312156133ab57600080fd5b6133b486612fdd565b94506133c260208701612fdd565b935060408601359250606086013567ffffffffffffffff8111156133e557600080fd5b6133f188828901613094565b969995985093965092949392505050565b6000806000806060858703121561341857600080fd5b61342185612fdd565b935060208501359250604085013567ffffffffffffffff81111561344457600080fd5b61345087828801613094565b95989497509550505050565b6000806040838503121561346f57600080fd5b61347883612fdd565b9150602083013567ffffffffffffffff81111561349457600080fd5b6134a085828601613014565b9150509250929050565b6000806000606084860312156134bf57600080fd5b6134c884612fdd565b92506134d660208501612fdd565b91506134e460408501612fdd565b90509250925092565b60008060008060008060a0878903121561350657600080fd5b61350f87612fdd565b95506020870135945061352460408801612fdd565b9350606087013567ffffffffffffffff81111561354057600080fd5b61354c89828a01613094565b909450925050608087013567ffffffffffffffff81111561356c57600080fd5b61357889828a01613014565b9150509295509295509295565b60006060828403121561302657600080fd5b60008060008060008060a087890312156135b057600080fd5b6135b987612fdd565b95506135c760208801612fdd565b945060408701359350606087013567ffffffffffffffff8111156135ea57600080fd5b6135f689828a01613094565b909450925050608087013567ffffffffffffffff81111561361657600080fd5b61357889828a01613585565b6000806000806060858703121561363857600080fd5b61364185612fdd565b9350602085013567ffffffffffffffff81111561365d57600080fd5b61366987828801613094565b909450925050604085013567ffffffffffffffff81111561368957600080fd5b61308887828801613585565b8015158114610c3e57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126136d857600080fd5b830160208101925035905067ffffffffffffffff8111156136f857600080fd5b8036038213156130d657600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b0361376182612fdd565b1682526000602082013561377481613695565b151560208401526001600160a01b0361378f60408401612fdd565b1660408401526137a260608301836136a3565b60a060608601526137b760a086018284613707565b6080948501359590940194909452509092915050565b8381526001600160a01b0383166020820152608060408201526000608082015260a060608201526000612ddf60a0830184613750565b604081526000613817604083018587613707565b8281036020840152612f4f8185613750565b838152604060208201526000612ddf604083018486613707565b60006020828403121561385557600080fd5b815161220b81613695565b60006020828403121561387257600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b8581526001600160a01b03851660208201526080604082015260006138d1608083018587613707565b82810360608401526138e38185613750565b98975050505050505050565b6001600160a01b0361390082612fdd565b1682526000602082013567ffffffffffffffff811680821461392157600080fd5b60208501525061393460408301836136a3565b60606040860152612ddf606086018284613707565b60208152600061220b60208301846138ef565b848152606060208201526000613976606083018587613707565b828103604084015261398881856138ef565b979650505050505050565b8183823760009101908152919050565b600082516139b5818460208701613197565b919091019291505056fea2646970667358221220eb9db78f69035a9231a94744bf13ed3516f4b1c8f6d842cc4bf3592dcc14194064736f6c634300081a0033608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220a043c41353215fce25ecb67a8a4f6f724aaa86dea8dcb0a6975bfb1f10ff3beb64736f6c634300081a003360a060405234801561001057600080fd5b50604051611dff380380611dff83398101604081905261002f916101fd565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03838116608052600480546001600160a01b0319169184169190911790556100bb60008261014c565b506100e67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261014c565b506101117f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361014c565b5061012a600080516020611ddf8339815191528261014c565b50610143600080516020611ddf8339815191528361014c565b50505050610240565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166101d75760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016101db565b5060005b92915050565b80516001600160a01b03811681146101f857600080fd5b919050565b60008060006060848603121561021257600080fd5b61021b846101e1565b9250610229602085016101e1565b9150610237604085016101e1565b90509250925092565b608051611b68610277600039600081816101d501528181610574015281816105c901528181610b010152610b560152611b686000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c806385f438c1116100ee578063c709ab6e11610097578063d9caed1211610071578063d9caed12146103e0578063e609055e146103f3578063e63ab1e914610406578063eab103df1461042d57600080fd5b8063c709ab6e14610397578063d547741f146103aa578063d936547e146103bd57600080fd5b80639a590427116100c85780639a590427146103695780639b19251a1461037c578063a217fddf1461038f57600080fd5b806385f438c1146102f657806391d148541461031d578063950837aa1461035657600080fd5b806336568abe116101505780635b1125911161012a5780635b112591146102d05780635c975abb146102e35780638456cb59146102ee57600080fd5b806336568abe1461028e5780633f4ba83a146102a1578063570618e1146102a957600080fd5b8063248a9ca311610181578063248a9ca314610224578063252f07bf146102565780632f2ff15d1461027b57600080fd5b806301ffc9a7146101a8578063116191b6146101d057806321fc65f21461020f575b600080fd5b6101bb6101b636600461155e565b610440565b60405190151581526020015b60405180910390f35b6101f77f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101c7565b61022261021d3660046115fe565b6104d9565b005b610248610232366004611671565b6000908152600160208190526040909120015490565b6040519081526020016101c7565b6004546101bb9074010000000000000000000000000000000000000000900460ff1681565b61022261028936600461168a565b610699565b61022261029c36600461168a565b6106c5565b610222610716565b6102487f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b6004546101f7906001600160a01b031681565b60025460ff166101bb565b61022261074b565b6102487f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101bb61032b36600461168a565b60009182526001602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6102226103643660046116ba565b61077d565b6102226103773660046116ba565b6108fb565b61022261038a3660046116ba565b6109af565b610248600081565b6102226103a53660046116d7565b610a66565b6102226103b836600461168a565b610c2b565b6101bb6103cb3660046116ba565b60036020526000908152604090205460ff1681565b6102226103ee36600461177a565b610c51565b6102226104013660046117bb565b610d49565b6102487f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b61022261043b36600461185a565b610f75565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104d357507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104e1610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461050b8161100e565b610513611018565b6001600160a01b03851660009081526003602052604090205460ff16610565576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105996001600160a01b0386167f000000000000000000000000000000000000000000000000000000000000000086611057565b6040517f5131ab590000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906106069088908a908990899089906004016118c0565b600060405180830381600087803b15801561062057600080fd5b505af1158015610634573d6000803e3d6000fd5b50505050846001600160a01b0316866001600160a01b03167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d586868660405161067f93929190611903565b60405180910390a3506106926001600055565b5050505050565b600082815260016020819052604090912001546106b58161100e565b6106bf83836110cb565b50505050565b6001600160a01b0381163314610707576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610711828261115e565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107408161100e565b6107486111e5565b50565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107758161100e565b610748611237565b60006107888161100e565b6001600160a01b0382166107c8576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004546107ff907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4906001600160a01b031661115e565b50600454610837907f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a906001600160a01b031661115e565b506108627f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4836110cb565b5061088d7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a836110cb565b50600480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040519081527f086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba9060200160405180910390a15050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6109258161100e565b6001600160a01b038216610965576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6109d98161100e565b6001600160a01b038216610a19576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b610a6e610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610a988161100e565b610aa0611018565b6001600160a01b03861660009081526003602052604090205460ff16610af2576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b266001600160a01b0387167f000000000000000000000000000000000000000000000000000000000000000087611057565b6040517fd0b492c30000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610b959089908b908a908a908a908a906004016119d7565b600060405180830381600087803b158015610baf57600080fd5b505af1158015610bc3573d6000803e3d6000fd5b50505050856001600160a01b0316876001600160a01b03167f2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b87878787604051610c109493929190611a2e565b60405180910390a350610c236001600055565b505050505050565b60008281526001602081905260409091200154610c478161100e565b6106bf838361115e565b610c59610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610c838161100e565b610c8b611018565b6001600160a01b03831660009081526003602052604090205460ff16610cdd576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cf16001600160a01b0384168584611057565b826001600160a01b0316846001600160a01b03167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610d3691815260200190565b60405180910390a3506107116001600055565b610d51610fcb565b610d59611018565b60045474010000000000000000000000000000000000000000900460ff16610dad576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03841660009081526003602052604090205460ff16610dff576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610e5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e839190611a5a565b9050610e9a6001600160a01b038616333087611274565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa158015610f21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f459190611a5a565b610f4f9190611a73565b8787604051610f62959493929190611aad565b60405180910390a250610c236001600055565b6000610f808161100e565b506004805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b600260005403611007576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61074881336112ad565b60025460ff1615611055576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6040516001600160a01b0383811660248301526044820183905261071191859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611324565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166111565760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016104d3565b5060006104d3565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff16156111565760008381526001602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104d3565b6111ed6113a0565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b61123f611018565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861121a3390565b6040516001600160a01b0384811660248301528381166044830152606482018390526106bf9186918216906323b872dd90608401611084565b60008281526001602090815260408083206001600160a01b038516845290915290205460ff16611320576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044015b60405180910390fd5b5050565b60006113396001600160a01b038416836113dc565b9050805160001415801561135e57508080602001905181019061135c9190611ae6565b155b15610711576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611317565b60025460ff16611055576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606113ea838360006113f1565b9392505050565b60608147101561142f576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611317565b600080856001600160a01b0316848660405161144b9190611b03565b60006040518083038185875af1925050503d8060008114611488576040519150601f19603f3d011682016040523d82523d6000602084013e61148d565b606091505b509150915061149d8683836114a7565b9695505050505050565b6060826114bc576114b78261151c565b6113ea565b81511580156114d357506001600160a01b0384163b155b15611515576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611317565b50806113ea565b80511561152c5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561157057600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146113ea57600080fd5b6001600160a01b038116811461074857600080fd5b60008083601f8401126115c757600080fd5b50813567ffffffffffffffff8111156115df57600080fd5b6020830191508360208285010111156115f757600080fd5b9250929050565b60008060008060006080868803121561161657600080fd5b8535611621816115a0565b94506020860135611631816115a0565b935060408601359250606086013567ffffffffffffffff81111561165457600080fd5b611660888289016115b5565b969995985093965092949392505050565b60006020828403121561168357600080fd5b5035919050565b6000806040838503121561169d57600080fd5b8235915060208301356116af816115a0565b809150509250929050565b6000602082840312156116cc57600080fd5b81356113ea816115a0565b60008060008060008060a087890312156116f057600080fd5b86356116fb816115a0565b9550602087013561170b816115a0565b945060408701359350606087013567ffffffffffffffff81111561172e57600080fd5b61173a89828a016115b5565b909450925050608087013567ffffffffffffffff81111561175a57600080fd5b87016060818a03121561176c57600080fd5b809150509295509295509295565b60008060006060848603121561178f57600080fd5b833561179a816115a0565b925060208401356117aa816115a0565b929592945050506040919091013590565b600080600080600080608087890312156117d457600080fd5b863567ffffffffffffffff8111156117eb57600080fd5b6117f789828a016115b5565b909750955050602087013561180b816115a0565b935060408701359250606087013567ffffffffffffffff81111561182e57600080fd5b61183a89828a016115b5565b979a9699509497509295939492505050565b801515811461074857600080fd5b60006020828403121561186c57600080fd5b81356113ea8161184c565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b03861681526001600160a01b03851660208201528360408201526080606082015260006118f8608083018486611877565b979650505050505050565b83815260406020820152600061191d604083018486611877565b95945050505050565b60008135611933816115a0565b6001600160a01b03168352602082013567ffffffffffffffff811680821461195a57600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261199657600080fd5b820160208101903567ffffffffffffffff8111156119b357600080fd5b8036038213156119c257600080fd5b6060604086015261191d606086018284611877565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a060608201526000611a0f60a083018587611877565b8281036080840152611a218185611926565b9998505050505050505050565b848152606060208201526000611a48606083018587611877565b82810360408401526118f88185611926565b600060208284031215611a6c57600080fd5b5051919050565b818103818111156104d3577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b606081526000611ac1606083018789611877565b8560208401528281036040840152611ada818587611877565b98975050505050505050565b600060208284031215611af857600080fd5b81516113ea8161184c565b6000825160005b81811015611b245760208186018101518583015201611b0a565b50600092019182525091905056fea26469706673582212209a03cf1000b6a4d1e791944a4fc981506999a411e7296ec7a4f0824d26398a3764736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a",
}

// GatewayEVMEchidnaTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMEchidnaTestMetaData.ABI instead.
var GatewayEVMEchidnaTestABI = GatewayEVMEchidnaTestMetaData.ABI

// GatewayEVMEchidnaTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMEchidnaTestMetaData.Bin instead.
var GatewayEVMEchidnaTestBin = GatewayEVMEchidnaTestMetaData.Bin

// DeployGatewayEVMEchidnaTest deploys a new Ethereum contract, binding an instance of GatewayEVMEchidnaTest to it.
func DeployGatewayEVMEchidnaTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMEchidnaTest, error) {
	parsed, err := GatewayEVMEchidnaTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMEchidnaTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMEchidnaTest{GatewayEVMEchidnaTestCaller: GatewayEVMEchidnaTestCaller{contract: contract}, GatewayEVMEchidnaTestTransactor: GatewayEVMEchidnaTestTransactor{contract: contract}, GatewayEVMEchidnaTestFilterer: GatewayEVMEchidnaTestFilterer{contract: contract}}, nil
}

// GatewayEVMEchidnaTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMEchidnaTest struct {
	GatewayEVMEchidnaTestCaller     // Read-only binding to the contract
	GatewayEVMEchidnaTestTransactor // Write-only binding to the contract
	GatewayEVMEchidnaTestFilterer   // Log filterer for contract events
}

// GatewayEVMEchidnaTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMEchidnaTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMEchidnaTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMEchidnaTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMEchidnaTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMEchidnaTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMEchidnaTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMEchidnaTestSession struct {
	Contract     *GatewayEVMEchidnaTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayEVMEchidnaTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMEchidnaTestCallerSession struct {
	Contract *GatewayEVMEchidnaTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// GatewayEVMEchidnaTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMEchidnaTestTransactorSession struct {
	Contract     *GatewayEVMEchidnaTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// GatewayEVMEchidnaTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMEchidnaTestRaw struct {
	Contract *GatewayEVMEchidnaTest // Generic contract binding to access the raw methods on
}

// GatewayEVMEchidnaTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMEchidnaTestCallerRaw struct {
	Contract *GatewayEVMEchidnaTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMEchidnaTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMEchidnaTestTransactorRaw struct {
	Contract *GatewayEVMEchidnaTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMEchidnaTest creates a new instance of GatewayEVMEchidnaTest, bound to a specific deployed contract.
func NewGatewayEVMEchidnaTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMEchidnaTest, error) {
	contract, err := bindGatewayEVMEchidnaTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTest{GatewayEVMEchidnaTestCaller: GatewayEVMEchidnaTestCaller{contract: contract}, GatewayEVMEchidnaTestTransactor: GatewayEVMEchidnaTestTransactor{contract: contract}, GatewayEVMEchidnaTestFilterer: GatewayEVMEchidnaTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMEchidnaTestCaller creates a new read-only instance of GatewayEVMEchidnaTest, bound to a specific deployed contract.
func NewGatewayEVMEchidnaTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMEchidnaTestCaller, error) {
	contract, err := bindGatewayEVMEchidnaTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestCaller{contract: contract}, nil
}

// NewGatewayEVMEchidnaTestTransactor creates a new write-only instance of GatewayEVMEchidnaTest, bound to a specific deployed contract.
func NewGatewayEVMEchidnaTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMEchidnaTestTransactor, error) {
	contract, err := bindGatewayEVMEchidnaTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestTransactor{contract: contract}, nil
}

// NewGatewayEVMEchidnaTestFilterer creates a new log filterer instance of GatewayEVMEchidnaTest, bound to a specific deployed contract.
func NewGatewayEVMEchidnaTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMEchidnaTestFilterer, error) {
	contract, err := bindGatewayEVMEchidnaTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestFilterer{contract: contract}, nil
}

// bindGatewayEVMEchidnaTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMEchidnaTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMEchidnaTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMEchidnaTest.Contract.GatewayEVMEchidnaTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.GatewayEVMEchidnaTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.GatewayEVMEchidnaTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMEchidnaTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.contract.Transact(opts, method, params...)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) ASSETHANDLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "ASSET_HANDLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVMEchidnaTest.Contract.ASSETHANDLERROLE(&_GatewayEVMEchidnaTest.CallOpts)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVMEchidnaTest.Contract.ASSETHANDLERROLE(&_GatewayEVMEchidnaTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVMEchidnaTest.Contract.DEFAULTADMINROLE(&_GatewayEVMEchidnaTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVMEchidnaTest.Contract.DEFAULTADMINROLE(&_GatewayEVMEchidnaTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVMEchidnaTest.Contract.PAUSERROLE(&_GatewayEVMEchidnaTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVMEchidnaTest.Contract.PAUSERROLE(&_GatewayEVMEchidnaTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVMEchidnaTest.Contract.TSSROLE(&_GatewayEVMEchidnaTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVMEchidnaTest.Contract.TSSROLE(&_GatewayEVMEchidnaTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayEVMEchidnaTest.Contract.UPGRADEINTERFACEVERSION(&_GatewayEVMEchidnaTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayEVMEchidnaTest.Contract.UPGRADEINTERFACEVERSION(&_GatewayEVMEchidnaTest.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) Custody(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "custody")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) Custody() (common.Address, error) {
	return _GatewayEVMEchidnaTest.Contract.Custody(&_GatewayEVMEchidnaTest.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) Custody() (common.Address, error) {
	return _GatewayEVMEchidnaTest.Contract.Custody(&_GatewayEVMEchidnaTest.CallOpts)
}

// EchidnaCaller is a free data retrieval call binding the contract method 0x81100bf0.
//
// Solidity: function echidnaCaller() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) EchidnaCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "echidnaCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EchidnaCaller is a free data retrieval call binding the contract method 0x81100bf0.
//
// Solidity: function echidnaCaller() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) EchidnaCaller() (common.Address, error) {
	return _GatewayEVMEchidnaTest.Contract.EchidnaCaller(&_GatewayEVMEchidnaTest.CallOpts)
}

// EchidnaCaller is a free data retrieval call binding the contract method 0x81100bf0.
//
// Solidity: function echidnaCaller() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) EchidnaCaller() (common.Address, error) {
	return _GatewayEVMEchidnaTest.Contract.EchidnaCaller(&_GatewayEVMEchidnaTest.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayEVMEchidnaTest.Contract.GetRoleAdmin(&_GatewayEVMEchidnaTest.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayEVMEchidnaTest.Contract.GetRoleAdmin(&_GatewayEVMEchidnaTest.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayEVMEchidnaTest.Contract.HasRole(&_GatewayEVMEchidnaTest.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayEVMEchidnaTest.Contract.HasRole(&_GatewayEVMEchidnaTest.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) Paused() (bool, error) {
	return _GatewayEVMEchidnaTest.Contract.Paused(&_GatewayEVMEchidnaTest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) Paused() (bool, error) {
	return _GatewayEVMEchidnaTest.Contract.Paused(&_GatewayEVMEchidnaTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayEVMEchidnaTest.Contract.ProxiableUUID(&_GatewayEVMEchidnaTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayEVMEchidnaTest.Contract.ProxiableUUID(&_GatewayEVMEchidnaTest.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayEVMEchidnaTest.Contract.SupportsInterface(&_GatewayEVMEchidnaTest.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayEVMEchidnaTest.Contract.SupportsInterface(&_GatewayEVMEchidnaTest.CallOpts, interfaceId)
}

// TestERC20 is a free data retrieval call binding the contract method 0x3c2f05a8.
//
// Solidity: function testERC20() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) TestERC20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "testERC20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TestERC20 is a free data retrieval call binding the contract method 0x3c2f05a8.
//
// Solidity: function testERC20() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) TestERC20() (common.Address, error) {
	return _GatewayEVMEchidnaTest.Contract.TestERC20(&_GatewayEVMEchidnaTest.CallOpts)
}

// TestERC20 is a free data retrieval call binding the contract method 0x3c2f05a8.
//
// Solidity: function testERC20() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) TestERC20() (common.Address, error) {
	return _GatewayEVMEchidnaTest.Contract.TestERC20(&_GatewayEVMEchidnaTest.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) TssAddress() (common.Address, error) {
	return _GatewayEVMEchidnaTest.Contract.TssAddress(&_GatewayEVMEchidnaTest.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) TssAddress() (common.Address, error) {
	return _GatewayEVMEchidnaTest.Contract.TssAddress(&_GatewayEVMEchidnaTest.CallOpts)
}

// ZetaConnector is a free data retrieval call binding the contract method 0x57bec62f.
//
// Solidity: function zetaConnector() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) ZetaConnector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "zetaConnector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaConnector is a free data retrieval call binding the contract method 0x57bec62f.
//
// Solidity: function zetaConnector() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) ZetaConnector() (common.Address, error) {
	return _GatewayEVMEchidnaTest.Contract.ZetaConnector(&_GatewayEVMEchidnaTest.CallOpts)
}

// ZetaConnector is a free data retrieval call binding the contract method 0x57bec62f.
//
// Solidity: function zetaConnector() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) ZetaConnector() (common.Address, error) {
	return _GatewayEVMEchidnaTest.Contract.ZetaConnector(&_GatewayEVMEchidnaTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMEchidnaTest.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) ZetaToken() (common.Address, error) {
	return _GatewayEVMEchidnaTest.Contract.ZetaToken(&_GatewayEVMEchidnaTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestCallerSession) ZetaToken() (common.Address, error) {
	return _GatewayEVMEchidnaTest.Contract.ZetaToken(&_GatewayEVMEchidnaTest.CallOpts)
}

// Call is a paid mutator transaction binding the contract method 0x1becceb4.
//
// Solidity: function call(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) Call(opts *bind.TransactOpts, receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "call", receiver, payload, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x1becceb4.
//
// Solidity: function call(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) Call(receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Call(&_GatewayEVMEchidnaTest.TransactOpts, receiver, payload, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x1becceb4.
//
// Solidity: function call(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) Call(receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Call(&_GatewayEVMEchidnaTest.TransactOpts, receiver, payload, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0x102614b0.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) Deposit(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, asset common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "deposit", receiver, amount, asset, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0x102614b0.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) Deposit(receiver common.Address, amount *big.Int, asset common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Deposit(&_GatewayEVMEchidnaTest.TransactOpts, receiver, amount, asset, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0x102614b0.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) Deposit(receiver common.Address, amount *big.Int, asset common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Deposit(&_GatewayEVMEchidnaTest.TransactOpts, receiver, amount, asset, revertOptions)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x726ac97c.
//
// Solidity: function deposit(address receiver, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) Deposit0(opts *bind.TransactOpts, receiver common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "deposit0", receiver, revertOptions)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x726ac97c.
//
// Solidity: function deposit(address receiver, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) Deposit0(receiver common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Deposit0(&_GatewayEVMEchidnaTest.TransactOpts, receiver, revertOptions)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x726ac97c.
//
// Solidity: function deposit(address receiver, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) Deposit0(receiver common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Deposit0(&_GatewayEVMEchidnaTest.TransactOpts, receiver, revertOptions)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x744b9b8b.
//
// Solidity: function depositAndCall(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) DepositAndCall(opts *bind.TransactOpts, receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "depositAndCall", receiver, payload, revertOptions)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x744b9b8b.
//
// Solidity: function depositAndCall(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) DepositAndCall(receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.DepositAndCall(&_GatewayEVMEchidnaTest.TransactOpts, receiver, payload, revertOptions)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x744b9b8b.
//
// Solidity: function depositAndCall(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) DepositAndCall(receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.DepositAndCall(&_GatewayEVMEchidnaTest.TransactOpts, receiver, payload, revertOptions)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xd09e3b78.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) DepositAndCall0(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, asset common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "depositAndCall0", receiver, amount, asset, payload, revertOptions)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xd09e3b78.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) DepositAndCall0(receiver common.Address, amount *big.Int, asset common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.DepositAndCall0(&_GatewayEVMEchidnaTest.TransactOpts, receiver, amount, asset, payload, revertOptions)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xd09e3b78.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) DepositAndCall0(receiver common.Address, amount *big.Int, asset common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.DepositAndCall0(&_GatewayEVMEchidnaTest.TransactOpts, receiver, amount, asset, payload, revertOptions)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) Execute(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "execute", destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Execute(&_GatewayEVMEchidnaTest.TransactOpts, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) payable returns(bytes)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Execute(&_GatewayEVMEchidnaTest.TransactOpts, destination, data)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0xf7ad60db.
//
// Solidity: function executeRevert(address destination, bytes data, (address,uint64,bytes) revertContext) payable returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) ExecuteRevert(opts *bind.TransactOpts, destination common.Address, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "executeRevert", destination, data, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0xf7ad60db.
//
// Solidity: function executeRevert(address destination, bytes data, (address,uint64,bytes) revertContext) payable returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) ExecuteRevert(destination common.Address, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.ExecuteRevert(&_GatewayEVMEchidnaTest.TransactOpts, destination, data, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0xf7ad60db.
//
// Solidity: function executeRevert(address destination, bytes data, (address,uint64,bytes) revertContext) payable returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) ExecuteRevert(destination common.Address, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.ExecuteRevert(&_GatewayEVMEchidnaTest.TransactOpts, destination, data, revertContext)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) ExecuteWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "executeWithERC20", token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.ExecuteWithERC20(&_GatewayEVMEchidnaTest.TransactOpts, token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x5131ab59.
//
// Solidity: function executeWithERC20(address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) ExecuteWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.ExecuteWithERC20(&_GatewayEVMEchidnaTest.TransactOpts, token, to, amount, data)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.GrantRole(&_GatewayEVMEchidnaTest.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.GrantRole(&_GatewayEVMEchidnaTest.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address tssAddress_, address zetaToken_, address admin_) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) Initialize(opts *bind.TransactOpts, tssAddress_ common.Address, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "initialize", tssAddress_, zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address tssAddress_, address zetaToken_, address admin_) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) Initialize(tssAddress_ common.Address, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Initialize(&_GatewayEVMEchidnaTest.TransactOpts, tssAddress_, zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address tssAddress_, address zetaToken_, address admin_) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) Initialize(tssAddress_ common.Address, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Initialize(&_GatewayEVMEchidnaTest.TransactOpts, tssAddress_, zetaToken_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) Pause() (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Pause(&_GatewayEVMEchidnaTest.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) Pause() (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Pause(&_GatewayEVMEchidnaTest.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.RenounceRole(&_GatewayEVMEchidnaTest.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.RenounceRole(&_GatewayEVMEchidnaTest.TransactOpts, role, callerConfirmation)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xd0b492c3.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data, (address,uint64,bytes) revertContext) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) RevertWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "revertWithERC20", token, to, amount, data, revertContext)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xd0b492c3.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data, (address,uint64,bytes) revertContext) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) RevertWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.RevertWithERC20(&_GatewayEVMEchidnaTest.TransactOpts, token, to, amount, data, revertContext)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xd0b492c3.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data, (address,uint64,bytes) revertContext) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) RevertWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.RevertWithERC20(&_GatewayEVMEchidnaTest.TransactOpts, token, to, amount, data, revertContext)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.RevokeRole(&_GatewayEVMEchidnaTest.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.RevokeRole(&_GatewayEVMEchidnaTest.TransactOpts, role, account)
}

// SetConnector is a paid mutator transaction binding the contract method 0x10188aef.
//
// Solidity: function setConnector(address zetaConnector_) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) SetConnector(opts *bind.TransactOpts, zetaConnector_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "setConnector", zetaConnector_)
}

// SetConnector is a paid mutator transaction binding the contract method 0x10188aef.
//
// Solidity: function setConnector(address zetaConnector_) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) SetConnector(zetaConnector_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.SetConnector(&_GatewayEVMEchidnaTest.TransactOpts, zetaConnector_)
}

// SetConnector is a paid mutator transaction binding the contract method 0x10188aef.
//
// Solidity: function setConnector(address zetaConnector_) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) SetConnector(zetaConnector_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.SetConnector(&_GatewayEVMEchidnaTest.TransactOpts, zetaConnector_)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address custody_) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) SetCustody(opts *bind.TransactOpts, custody_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "setCustody", custody_)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address custody_) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) SetCustody(custody_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.SetCustody(&_GatewayEVMEchidnaTest.TransactOpts, custody_)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address custody_) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) SetCustody(custody_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.SetCustody(&_GatewayEVMEchidnaTest.TransactOpts, custody_)
}

// TestExecuteWithERC20 is a paid mutator transaction binding the contract method 0x6ab90f9b.
//
// Solidity: function testExecuteWithERC20(address to, uint256 amount, bytes data) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) TestExecuteWithERC20(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "testExecuteWithERC20", to, amount, data)
}

// TestExecuteWithERC20 is a paid mutator transaction binding the contract method 0x6ab90f9b.
//
// Solidity: function testExecuteWithERC20(address to, uint256 amount, bytes data) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) TestExecuteWithERC20(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.TestExecuteWithERC20(&_GatewayEVMEchidnaTest.TransactOpts, to, amount, data)
}

// TestExecuteWithERC20 is a paid mutator transaction binding the contract method 0x6ab90f9b.
//
// Solidity: function testExecuteWithERC20(address to, uint256 amount, bytes data) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) TestExecuteWithERC20(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.TestExecuteWithERC20(&_GatewayEVMEchidnaTest.TransactOpts, to, amount, data)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) Unpause() (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Unpause(&_GatewayEVMEchidnaTest.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) Unpause() (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Unpause(&_GatewayEVMEchidnaTest.TransactOpts)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) UpdateTSSAddress(opts *bind.TransactOpts, newTSSAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "updateTSSAddress", newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.UpdateTSSAddress(&_GatewayEVMEchidnaTest.TransactOpts, newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.UpdateTSSAddress(&_GatewayEVMEchidnaTest.TransactOpts, newTSSAddress)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.UpgradeToAndCall(&_GatewayEVMEchidnaTest.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.UpgradeToAndCall(&_GatewayEVMEchidnaTest.TransactOpts, newImplementation, data)
}

// GatewayEVMEchidnaTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestCalledIterator struct {
	Event *GatewayEVMEchidnaTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestCalled)
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
		it.Event = new(GatewayEVMEchidnaTestCalled)
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
func (it *GatewayEVMEchidnaTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestCalled represents a Called event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMEchidnaTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestCalledIterator{contract: _GatewayEVMEchidnaTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestCalled)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParseCalled(log types.Log) (*GatewayEVMEchidnaTestCalled, error) {
	event := new(GatewayEVMEchidnaTestCalled)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMEchidnaTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestDepositedIterator struct {
	Event *GatewayEVMEchidnaTestDeposited // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestDeposited)
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
		it.Event = new(GatewayEVMEchidnaTestDeposited)
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
func (it *GatewayEVMEchidnaTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestDeposited represents a Deposited event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestDeposited struct {
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
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterDeposited(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMEchidnaTestDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestDepositedIterator{contract: _GatewayEVMEchidnaTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestDeposited, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestDeposited)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParseDeposited(log types.Log) (*GatewayEVMEchidnaTestDeposited, error) {
	event := new(GatewayEVMEchidnaTestDeposited)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMEchidnaTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestExecutedIterator struct {
	Event *GatewayEVMEchidnaTestExecuted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestExecuted)
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
		it.Event = new(GatewayEVMEchidnaTestExecuted)
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
func (it *GatewayEVMEchidnaTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestExecuted represents a Executed event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMEchidnaTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestExecutedIterator{contract: _GatewayEVMEchidnaTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestExecuted)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Executed", log); err != nil {
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
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParseExecuted(log types.Log) (*GatewayEVMEchidnaTestExecuted, error) {
	event := new(GatewayEVMEchidnaTestExecuted)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMEchidnaTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMEchidnaTestExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestExecutedWithERC20)
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
		it.Event = new(GatewayEVMEchidnaTestExecutedWithERC20)
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
func (it *GatewayEVMEchidnaTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMEchidnaTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestExecutedWithERC20Iterator{contract: _GatewayEVMEchidnaTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestExecutedWithERC20)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMEchidnaTestExecutedWithERC20, error) {
	event := new(GatewayEVMEchidnaTestExecutedWithERC20)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMEchidnaTestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestInitializedIterator struct {
	Event *GatewayEVMEchidnaTestInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestInitialized)
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
		it.Event = new(GatewayEVMEchidnaTestInitialized)
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
func (it *GatewayEVMEchidnaTestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestInitialized represents a Initialized event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayEVMEchidnaTestInitializedIterator, error) {

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestInitializedIterator{contract: _GatewayEVMEchidnaTest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestInitialized)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParseInitialized(log types.Log) (*GatewayEVMEchidnaTestInitialized, error) {
	event := new(GatewayEVMEchidnaTestInitialized)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMEchidnaTestPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestPausedIterator struct {
	Event *GatewayEVMEchidnaTestPaused // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestPaused)
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
		it.Event = new(GatewayEVMEchidnaTestPaused)
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
func (it *GatewayEVMEchidnaTestPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestPaused represents a Paused event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterPaused(opts *bind.FilterOpts) (*GatewayEVMEchidnaTestPausedIterator, error) {

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestPausedIterator{contract: _GatewayEVMEchidnaTest.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestPaused)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParsePaused(log types.Log) (*GatewayEVMEchidnaTestPaused, error) {
	event := new(GatewayEVMEchidnaTestPaused)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMEchidnaTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestRevertedIterator struct {
	Event *GatewayEVMEchidnaTestReverted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestReverted)
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
		it.Event = new(GatewayEVMEchidnaTestReverted)
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
func (it *GatewayEVMEchidnaTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestReverted represents a Reverted event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestReverted struct {
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0x1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b1436.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMEchidnaTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestRevertedIterator{contract: _GatewayEVMEchidnaTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0x1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b1436.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestReverted)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Reverted", log); err != nil {
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

// ParseReverted is a log parse operation binding the contract event 0x1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b1436.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParseReverted(log types.Log) (*GatewayEVMEchidnaTestReverted, error) {
	event := new(GatewayEVMEchidnaTestReverted)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMEchidnaTestRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestRoleAdminChangedIterator struct {
	Event *GatewayEVMEchidnaTestRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestRoleAdminChanged)
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
		it.Event = new(GatewayEVMEchidnaTestRoleAdminChanged)
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
func (it *GatewayEVMEchidnaTestRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestRoleAdminChanged represents a RoleAdminChanged event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GatewayEVMEchidnaTestRoleAdminChangedIterator, error) {

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

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestRoleAdminChangedIterator{contract: _GatewayEVMEchidnaTest.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestRoleAdminChanged)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParseRoleAdminChanged(log types.Log) (*GatewayEVMEchidnaTestRoleAdminChanged, error) {
	event := new(GatewayEVMEchidnaTestRoleAdminChanged)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMEchidnaTestRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestRoleGrantedIterator struct {
	Event *GatewayEVMEchidnaTestRoleGranted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestRoleGranted)
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
		it.Event = new(GatewayEVMEchidnaTestRoleGranted)
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
func (it *GatewayEVMEchidnaTestRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestRoleGranted represents a RoleGranted event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayEVMEchidnaTestRoleGrantedIterator, error) {

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

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestRoleGrantedIterator{contract: _GatewayEVMEchidnaTest.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestRoleGranted)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParseRoleGranted(log types.Log) (*GatewayEVMEchidnaTestRoleGranted, error) {
	event := new(GatewayEVMEchidnaTestRoleGranted)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMEchidnaTestRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestRoleRevokedIterator struct {
	Event *GatewayEVMEchidnaTestRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestRoleRevoked)
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
		it.Event = new(GatewayEVMEchidnaTestRoleRevoked)
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
func (it *GatewayEVMEchidnaTestRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestRoleRevoked represents a RoleRevoked event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayEVMEchidnaTestRoleRevokedIterator, error) {

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

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestRoleRevokedIterator{contract: _GatewayEVMEchidnaTest.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestRoleRevoked)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParseRoleRevoked(log types.Log) (*GatewayEVMEchidnaTestRoleRevoked, error) {
	event := new(GatewayEVMEchidnaTestRoleRevoked)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMEchidnaTestUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestUnpausedIterator struct {
	Event *GatewayEVMEchidnaTestUnpaused // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestUnpaused)
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
		it.Event = new(GatewayEVMEchidnaTestUnpaused)
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
func (it *GatewayEVMEchidnaTestUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestUnpaused represents a Unpaused event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GatewayEVMEchidnaTestUnpausedIterator, error) {

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestUnpausedIterator{contract: _GatewayEVMEchidnaTest.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestUnpaused) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestUnpaused)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParseUnpaused(log types.Log) (*GatewayEVMEchidnaTestUnpaused, error) {
	event := new(GatewayEVMEchidnaTestUnpaused)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMEchidnaTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestUpdatedGatewayTSSAddressIterator struct {
	Event *GatewayEVMEchidnaTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestUpdatedGatewayTSSAddress)
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
		it.Event = new(GatewayEVMEchidnaTestUpdatedGatewayTSSAddress)
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
func (it *GatewayEVMEchidnaTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestUpdatedGatewayTSSAddress struct {
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*GatewayEVMEchidnaTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestUpdatedGatewayTSSAddressIterator{contract: _GatewayEVMEchidnaTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestUpdatedGatewayTSSAddress)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
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

// ParseUpdatedGatewayTSSAddress is a log parse operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*GatewayEVMEchidnaTestUpdatedGatewayTSSAddress, error) {
	event := new(GatewayEVMEchidnaTestUpdatedGatewayTSSAddress)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMEchidnaTestUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestUpgradedIterator struct {
	Event *GatewayEVMEchidnaTestUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayEVMEchidnaTestUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMEchidnaTestUpgraded)
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
		it.Event = new(GatewayEVMEchidnaTestUpgraded)
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
func (it *GatewayEVMEchidnaTestUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMEchidnaTestUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMEchidnaTestUpgraded represents a Upgraded event raised by the GatewayEVMEchidnaTest contract.
type GatewayEVMEchidnaTestUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayEVMEchidnaTestUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayEVMEchidnaTest.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMEchidnaTestUpgradedIterator{contract: _GatewayEVMEchidnaTest.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayEVMEchidnaTestUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayEVMEchidnaTest.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMEchidnaTestUpgraded)
				if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestFilterer) ParseUpgraded(log types.Log) (*GatewayEVMEchidnaTestUpgraded, error) {
	event := new(GatewayEVMEchidnaTestUpgraded)
	if err := _GatewayEVMEchidnaTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
