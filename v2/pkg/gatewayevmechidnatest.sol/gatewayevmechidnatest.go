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

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender common.Address
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

// GatewayEVMEchidnaTestMetaData contains all meta data concerning the GatewayEVMEchidnaTest contract.
var GatewayEVMEchidnaTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ASSET_HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"custody\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"echidnaCaller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"messageContext\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revertWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConnector\",\"inputs\":[{\"name\":\"zetaConnector_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustody\",\"inputs\":[{\"name\":\"custody_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testERC20\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractTestERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testExecuteWithERC20\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"zetaConnector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a060405230608052600580546001600160a01b0319163317905534801561002657600080fd5b5061002f610154565b600554600180546001600160a01b039092166001600160a01b031992831617905560028054610123921691909117905560405161006b90610206565b60408082526004908201819052631d195cdd60e21b606083015260806020830181905282015263151154d560e21b60a082015260c001604051809103906000f0801580156100bd573d6000803e3d6000fd5b50600480546001600160a01b0319166001600160a01b039283161790556001546040513092919091169082906100f290610213565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f08015801561012e573d6000803e3d6000fd5b50600080546001600160a01b0319166001600160a01b0392909216919091179055610220565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156101a45760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146102035780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b610c9f80613d5183390190565b611eb9806149f083390190565b608051613b086102496000396000818161239f015281816123c801526128360152613b086000f3fe60806040526004361061024f5760003560e01c80635d62c86011610138578063ad3cb1cc116100b0578063d0b492c31161007f578063dda79b7511610064578063dda79b751461072d578063e63ab1e91461074d578063f7ad60db1461078157600080fd5b8063d0b492c3146106ed578063d547741f1461070d57600080fd5b8063ad3cb1cc14610644578063ae7a3a6f1461068d578063c0c53b8b146106ad578063d09e3b78146106cd57600080fd5b806381100bf01161010757806391d14854116100ec57806391d1485414610596578063a217fddf146105fb578063a783c7891461061057600080fd5b806381100bf0146105615780638456cb591461058157600080fd5b80635d62c860146104e75780636ab90f9b1461051b578063726ac97c1461053b578063744b9b8b1461054e57600080fd5b806338e22527116101cb5780635131ab591161019a57806357bec62f1161017f57806357bec62f146104705780635b112591146104905780635c975abb146104b057600080fd5b80635131ab591461043b57806352d1902d1461045b57600080fd5b806338e22527146103e05780633c2f05a8146103f35780633f4ba83a146104135780634f1ef2861461042857600080fd5b80631cff79cd11610222578063248a9ca311610207578063248a9ca3146103435780632f2ff15d146103a057806336568abe146103c057600080fd5b80631cff79cd146102eb57806321e093b11461030b57600080fd5b806301ffc9a71461025457806310188aef14610289578063102614b0146102ab5780631becceb4146102cb575b600080fd5b34801561026057600080fd5b5061027461026f36600461301d565b610794565b60405190151581526020015b60405180910390f35b34801561029557600080fd5b506102a96102a436600461307b565b61082d565b005b3480156102b757600080fd5b506102a96102c63660046130ae565b610908565b3480156102d757600080fd5b506102a96102e636600461315f565b610a02565b6102fe6102f93660046131c6565b610a80565b6040516102809190613269565b34801561031757600080fd5b5060035461032b906001600160a01b031681565b6040516001600160a01b039091168152602001610280565b34801561034f57600080fd5b5061039261035e36600461327c565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b604051908152602001610280565b3480156103ac57600080fd5b506102a96103bb366004613295565b610b38565b3480156103cc57600080fd5b506102a96103db366004613295565b610b7c565b6102fe6103ee3660046132c1565b610bcd565b3480156103ff57600080fd5b5060045461032b906001600160a01b031681565b34801561041f57600080fd5b506102a9610cb9565b6102a96104363660046133b2565b610cee565b34801561044757600080fd5b506102a9610456366004613443565b610d0d565b34801561046757600080fd5b5061039261100d565b34801561047c57600080fd5b5060025461032b906001600160a01b031681565b34801561049c57600080fd5b5060015461032b906001600160a01b031681565b3480156104bc57600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610274565b3480156104f357600080fd5b506103927f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b981565b34801561052757600080fd5b506102a96105363660046134b2565b61103c565b6102a96105493660046134f4565b611167565b6102a961055c36600461315f565b6112df565b34801561056d57600080fd5b5060055461032b906001600160a01b031681565b34801561058d57600080fd5b506102a961145b565b3480156105a257600080fd5b506102746105b1366004613295565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561060757600080fd5b50610392600081565b34801561061c57600080fd5b506103927f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b34801561065057600080fd5b506102fe6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b34801561069957600080fd5b506102a96106a836600461307b565b61148d565b3480156106b957600080fd5b506102a96106c8366004613542565b611568565b3480156106d957600080fd5b506102a96106e8366004613585565b611804565b3480156106f957600080fd5b506102a961070836600461362f565b6118fc565b34801561071957600080fd5b506102a9610728366004613295565b611aa5565b34801561073957600080fd5b5060005461032b906001600160a01b031681565b34801561075957600080fd5b506103927f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6102a961078f3660046136ba565b611ae9565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061082757507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b600061083881611cd1565b6001600160a01b03821661085f5760405163d92e233d60e01b815260040160405180910390fd5b6002546001600160a01b0316156108a2576040517f0c8dc01600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108cc7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b983611cdb565b5050600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b610910611dc8565b610918611e26565b82600003610952576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384166109795760405163d92e233d60e01b815260040160405180910390fd5b610984338385611ea7565b836001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c8585856040516109cb93929190613847565b60405180910390a36109fc60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050565b610a0a611dc8565b610a12611e26565b6001600160a01b038416610a395760405163d92e233d60e01b815260040160405180910390fd5b836001600160a01b0316336001600160a01b03167fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d9748585856040516109cb9392919061387d565b60607f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb610aac81611cd1565b610ab4611dc8565b6001600160a01b038516610adb5760405163d92e233d60e01b815260040160405180910390fd5b6000610ae886868661210a565b9050856001600160a01b03167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f348787604051610b27939291906138a3565b60405180910390a295945050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610b7281611cd1565b6109fc8383611cdb565b6001600160a01b0381163314610bbe576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610bc882826121bd565b505050565b60607f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb610bf981611cd1565b610c01611dc8565b610c09611e26565b6001600160a01b038516610c305760405163d92e233d60e01b815260040160405180910390fd5b6060610c3e87878787612281565b9050856001600160a01b03167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f348787604051610c7d939291906138a3565b60405180910390a29150610cb060017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50949350505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610ce381611cd1565b610ceb612304565b50565b610cf6612394565b610cff82612464565b610d09828261246f565b5050565b7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b9610d3781611cd1565b610d3f611dc8565b610d47611e26565b83600003610d81576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038516610da85760405163d92e233d60e01b815260040160405180910390fd5b610db28686612575565b610de8576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0386811660048301526024820186905287169063095ea7b3906044016020604051808303816000875af1158015610e50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e7491906138bd565b610eaa576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610eb585848461210a565b50610ec08686612575565b610ef6576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038816906370a0823190602401602060405180830381865afa158015610f56573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f7a91906138da565b90508015610f8c57610f8c8782612605565b856001600160a01b0316876001600160a01b03167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382878787604051610fd3939291906138a3565b60405180910390a35061100560017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b600061101761282b565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b600480546040517f40c10f190000000000000000000000000000000000000000000000000000000081523092810192909252602482018590526001600160a01b0316906340c10f1990604401600060405180830381600087803b1580156110a257600080fd5b505af11580156110b6573d6000803e3d6000fd5b50506004546110d492506001600160a01b0316905085858585610d0d565b600480546040517f70a0823100000000000000000000000000000000000000000000000000000000815230928101929092526001600160a01b0316906370a0823190602401602060405180830381865afa158015611136573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115a91906138da565b156109fc576109fc6138f3565b61116f611dc8565b611177611e26565b346000036111b1576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166111d85760405163d92e233d60e01b815260040160405180910390fd5b6001546040516000916001600160a01b03169034908381818185875af1925050503d8060008114611225576040519150601f19603f3d011682016040523d82523d6000602084013e61122a565b606091505b5050905080611265576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c346000866040516112ad93929190613847565b60405180910390a350610d0960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6112e7611dc8565b6112ef611e26565b34600003611329576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384166113505760405163d92e233d60e01b815260040160405180910390fd5b6001546040516000916001600160a01b03169034908381818185875af1925050503d806000811461139d576040519150601f19603f3d011682016040523d82523d6000602084013e6113a2565b606091505b50509050806113dd576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b846001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c346000888888604051611429959493929190613922565b60405180910390a3506109fc60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61148581611cd1565b610ceb61288d565b600061149881611cd1565b6001600160a01b0382166114bf5760405163d92e233d60e01b815260040160405180910390fd5b6000546001600160a01b031615611502576040517fb337f37800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61152c7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b983611cdb565b5050600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156115b35750825b905060008267ffffffffffffffff1660011480156115d05750303b155b9050811580156115de575080155b15611615576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156116765784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816158061169357506001600160a01b038716155b156116b15760405163d92e233d60e01b815260040160405180910390fd5b6116b9612906565b6116c161290e565b6116c9612906565b6116d161291e565b6116dc600087611cdb565b506117077f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87611cdb565b50600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038a161790556117657f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb89611cdb565b50600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03891617905583156117fa5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b61180c611dc8565b611814611e26565b8460000361184e576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0386166118755760405163d92e233d60e01b815260040160405180910390fd5b611880338587611ea7565b856001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c87878787876040516118cb959493929190613922565b60405180910390a361100560017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b961192681611cd1565b61192e611dc8565b611936611e26565b84600003611970576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0386166119975760405163d92e233d60e01b815260040160405180910390fd5b6119ab6001600160a01b038816878761292e565b6040517f660b9de00000000000000000000000000000000000000000000000000000000081526001600160a01b0387169063660b9de0906119f09085906004016139c3565b600060405180830381600087803b158015611a0a57600080fd5b505af1158015611a1e573d6000803e3d6000fd5b50505050866001600160a01b0316866001600160a01b03167f1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b143687878787604051611a6b94939291906139d6565b60405180910390a3611a9c60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611adf81611cd1565b6109fc83836121bd565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb611b1381611cd1565b611b1b611dc8565b611b23611e26565b6001600160a01b038516611b4a5760405163d92e233d60e01b815260040160405180910390fd5b6000856001600160a01b03163460405160006040518083038185875af1925050503d8060008114611b97576040519150601f19603f3d011682016040523d82523d6000602084013e611b9c565b606091505b5050905080611bd7576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f660b9de00000000000000000000000000000000000000000000000000000000081526001600160a01b0387169063660b9de090611c1c9086906004016139c3565b600060405180830381600087803b158015611c3657600080fd5b505af1158015611c4a573d6000803e3d6000fd5b5050505060006001600160a01b0316866001600160a01b03167f1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b143634888888604051611c9894939291906139d6565b60405180910390a350611cca60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050505050565b610ceb81336129a2565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16611dbe576000848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611d743390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610827565b6000915050610827565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615611e24576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611ea1576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6003546001600160a01b039081169083160361200b57611ed26001600160a01b038316843084612a2f565b6002546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303816000875af1158015611f3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f6291906138bd565b611f98576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546040517f743e0c9b000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063743e0c9b90602401600060405180830381600087803b158015611ff757600080fd5b505af1158015611a9c573d6000803e3d6000fd5b6000546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529091169063d936547e90602401602060405180830381865afa15801561206e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061209291906138bd565b6120c8576040517fac2175f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610bc8906001600160a01b038481169186911684612a2f565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b60606121168383612a68565b600080856001600160a01b0316348686604051612134929190613a0d565b60006040518083038185875af1925050503d8060008114612171576040519150601f19603f3d011682016040523d82523d6000602084013e612176565b606091505b5091509150816121b2576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9150505b9392505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615611dbe576000848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610827565b6060836001600160a01b031663676cc054348786866040518563ffffffff1660e01b81526004016122b493929190613a1d565b60006040518083038185885af11580156122d2573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f191682016040526122fb9190810190613a48565b95945050505050565b61230c612b68565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061242d57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166124217f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15611e24576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610d0981611cd1565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156124c9575060408051601f3d908101601f191682019092526124c6918101906138da565b60015b61250f576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc811461256b576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401612506565b610bc88383612bc3565b6040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b038281166004830152600060248301819052919084169063095ea7b3906044016020604051808303816000875af11580156125e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121b691906138bd565b6003546001600160a01b0390811690831603612754576002546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303816000875af1158015612687573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126ab91906138bd565b6126e1576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546040517f743e0c9b000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063743e0c9b90602401600060405180830381600087803b15801561274057600080fd5b505af1158015611005573d6000803e3d6000fd5b6000546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529091169063d936547e90602401602060405180830381865afa1580156127b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127db91906138bd565b612811576040517fac2175f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610d09906001600160a01b0384811691168361292e565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611e24576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612895611dc8565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612376565b611e24612c19565b612916612c19565b611e24612c80565b612926612c19565b611e24612c88565b6040516001600160a01b03838116602483015260448201839052610bc891859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612cd9565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610d09576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401612506565b6040516001600160a01b0384811660248301528381166044830152606482018390526109fc9186918216906323b872dd9060840161295b565b60048110610d095781357f98933fac000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821601612aed576040517fed69977500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f99f46220000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821601610bc8576040517ff3459a9600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16611e24576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612bcc82612d55565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115612c1157610bc88282612dfd565b610d09612e6a565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16611e24576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6120e4612c19565b612c90612c19565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6000612cee6001600160a01b03841683612ea2565b90508051600014158015612d13575080806020019051810190612d1191906138bd565b155b15610bc8576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401612506565b806001600160a01b03163b600003612da4576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401612506565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051612e1a9190613ab6565b600060405180830381855af49150503d8060008114612e55576040519150601f19603f3d011682016040523d82523d6000602084013e612e5a565b606091505b50915091506122fb858383612eb0565b3415611e24576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606121b683836000612f25565b606082612ec557612ec082612fdb565b6121b6565b8151158015612edc57506001600160a01b0384163b155b15612f1e576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401612506565b50806121b6565b606081471015612f63576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401612506565b600080856001600160a01b03168486604051612f7f9190613ab6565b60006040518083038185875af1925050503d8060008114612fbc576040519150601f19603f3d011682016040523d82523d6000602084013e612fc1565b606091505b5091509150612fd1868383612eb0565b9695505050505050565b805115612feb5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561302f57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146121b657600080fd5b80356001600160a01b038116811461307657600080fd5b919050565b60006020828403121561308d57600080fd5b6121b68261305f565b600060a082840312156130a857600080fd5b50919050565b600080600080608085870312156130c457600080fd5b6130cd8561305f565b9350602085013592506130e26040860161305f565b9150606085013567ffffffffffffffff8111156130fe57600080fd5b61310a87828801613096565b91505092959194509250565b60008083601f84011261312857600080fd5b50813567ffffffffffffffff81111561314057600080fd5b60208301915083602082850101111561315857600080fd5b9250929050565b6000806000806060858703121561317557600080fd5b61317e8561305f565b9350602085013567ffffffffffffffff81111561319a57600080fd5b6131a687828801613116565b909450925050604085013567ffffffffffffffff8111156130fe57600080fd5b6000806000604084860312156131db57600080fd5b6131e48461305f565b9250602084013567ffffffffffffffff81111561320057600080fd5b61320c86828701613116565b9497909650939450505050565b60005b8381101561323457818101518382015260200161321c565b50506000910152565b60008151808452613255816020860160208601613219565b601f01601f19169290920160200192915050565b6020815260006121b6602083018461323d565b60006020828403121561328e57600080fd5b5035919050565b600080604083850312156132a857600080fd5b823591506132b86020840161305f565b90509250929050565b60008060008084860360608112156132d857600080fd5b60208112156132e657600080fd5b508493506132f66020860161305f565b9250604085013567ffffffffffffffff81111561331257600080fd5b61331e87828801613116565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156133825761338261332a565b604052919050565b600067ffffffffffffffff8211156133a4576133a461332a565b50601f01601f191660200190565b600080604083850312156133c557600080fd5b6133ce8361305f565b9150602083013567ffffffffffffffff8111156133ea57600080fd5b8301601f810185136133fb57600080fd5b803561340e6134098261338a565b613359565b81815286602083850101111561342357600080fd5b816020840160208301376000602083830101528093505050509250929050565b60008060008060006080868803121561345b57600080fd5b6134648661305f565b94506134726020870161305f565b935060408601359250606086013567ffffffffffffffff81111561349557600080fd5b6134a188828901613116565b969995985093965092949392505050565b600080600080606085870312156134c857600080fd5b6134d18561305f565b935060208501359250604085013567ffffffffffffffff81111561331257600080fd5b6000806040838503121561350757600080fd5b6135108361305f565b9150602083013567ffffffffffffffff81111561352c57600080fd5b61353885828601613096565b9150509250929050565b60008060006060848603121561355757600080fd5b6135608461305f565b925061356e6020850161305f565b915061357c6040850161305f565b90509250925092565b60008060008060008060a0878903121561359e57600080fd5b6135a78761305f565b9550602087013594506135bc6040880161305f565b9350606087013567ffffffffffffffff8111156135d857600080fd5b6135e489828a01613116565b909450925050608087013567ffffffffffffffff81111561360457600080fd5b61361089828a01613096565b9150509295509295509295565b6000606082840312156130a857600080fd5b60008060008060008060a0878903121561364857600080fd5b6136518761305f565b955061365f6020880161305f565b945060408701359350606087013567ffffffffffffffff81111561368257600080fd5b61368e89828a01613116565b909450925050608087013567ffffffffffffffff8111156136ae57600080fd5b61361089828a0161361d565b600080600080606085870312156136d057600080fd5b6136d98561305f565b9350602085013567ffffffffffffffff8111156136f557600080fd5b61370187828801613116565b909450925050604085013567ffffffffffffffff81111561372157600080fd5b61310a8782880161361d565b8015158114610ceb57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261377057600080fd5b830160208101925035905067ffffffffffffffff81111561379057600080fd5b80360382131561315857600080fd5b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b6001600160a01b036137db8261305f565b168252600060208201356137ee8161372d565b151560208401526001600160a01b036138096040840161305f565b16604084015261381c606083018361373b565b60a0606086015261383160a08601828461379f565b6080948501359590940194909452509092915050565b8381526001600160a01b0383166020820152608060408201526000608082015260a0606082015260006122fb60a08301846137ca565b60408152600061389160408301858761379f565b8281036020840152612fd181856137ca565b8381526040602082015260006122fb60408301848661379f565b6000602082840312156138cf57600080fd5b81516121b68161372d565b6000602082840312156138ec57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b8581526001600160a01b038516602082015260806040820152600061394b60808301858761379f565b828103606084015261395d81856137ca565b98975050505050505050565b6001600160a01b0361397a8261305f565b1682526000602082013567ffffffffffffffff811680821461399b57600080fd5b6020850152506139ae604083018361373b565b606060408601526122fb60608601828461379f565b6020815260006121b66020830184613969565b8481526060602082015260006139f060608301858761379f565b8281036040840152613a028185613969565b979650505050505050565b8183823760009101908152919050565b6001600160a01b03613a2e8561305f565b1681526040602082015260006122fb60408301848661379f565b600060208284031215613a5a57600080fd5b815167ffffffffffffffff811115613a7157600080fd5b8201601f81018413613a8257600080fd5b8051613a906134098261338a565b818152856020838501011115613aa557600080fd5b6122fb826020830160208601613219565b60008251613ac8818460208701613219565b919091019291505056fea2646970667358221220d2ecd7afd029240f00a84ca6100cd5407cb157ffe7b0c38cbda9db3a9e296c4c64736f6c634300081a0033608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220a043c41353215fce25ecb67a8a4f6f724aaa86dea8dcb0a6975bfb1f10ff3beb64736f6c634300081a003360a060405234801561001057600080fd5b50604051611eb9380380611eb983398101604081905261002f916101fd565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03838116608052600480546001600160a01b0319169184169190911790556100bb60008261014c565b506100e67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261014c565b506101117f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361014c565b5061012a600080516020611e998339815191528261014c565b50610143600080516020611e998339815191528361014c565b50505050610240565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166101d75760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016101db565b5060005b92915050565b80516001600160a01b03811681146101f857600080fd5b919050565b60008060006060848603121561021257600080fd5b61021b846101e1565b9250610229602085016101e1565b9150610237604085016101e1565b90509250925092565b608051611c22610277600039600081816101ca01528181610597015281816105f901528181610a280152610a8a0152611c226000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c806385f438c1116100e3578063d547741f1161008c578063e609055e11610066578063e609055e146103fc578063e63ab1e91461040f578063eab103df1461043657600080fd5b8063d547741f146103b3578063d936547e146103c6578063d9caed12146103e957600080fd5b80639b19251a116100bd5780639b19251a14610385578063a217fddf14610398578063c709ab6e146103a057600080fd5b806385f438c11461030557806391d148541461032c5780639a5904271461037257600080fd5b806336568abe116101455780635b1125911161011f5780635b112591146102d25780635c975abb146102f25780638456cb59146102fd57600080fd5b806336568abe146102905780633f4ba83a146102a3578063570618e1146102ab57600080fd5b8063248a9ca311610176578063248a9ca314610226578063252f07bf146102585780632f2ff15d1461027d57600080fd5b806301ffc9a71461019d578063116191b6146101c557806321fc65f214610211575b600080fd5b6101b06101ab3660046115ca565b610449565b60405190151581526020015b60405180910390f35b6101ec7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101bc565b61022461021f366004611677565b6104e2565b005b61024a6102343660046116ea565b6000908152600160208190526040909120015490565b6040519081526020016101bc565b6004546101b09074010000000000000000000000000000000000000000900460ff1681565b61022461028b366004611703565b6106e3565b61022461029e366004611703565b61070f565b61022461076d565b61024a7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b6004546101ec9073ffffffffffffffffffffffffffffffffffffffff1681565b60025460ff166101b0565b6102246107a2565b61024a7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101b061033a366004611703565b600091825260016020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610224610380366004611733565b6107d4565b610224610393366004611733565b6108a2565b61024a600081565b6102246103ae366004611750565b610973565b6102246103c1366004611703565b610b79565b6101b06103d4366004611733565b60036020526000908152604090205460ff1681565b6102246103f73660046117f3565b610b9f565b61022461040a366004611834565b610ccb565b61024a7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6102246104443660046118d3565b610f2b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104dc57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104ea610f81565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461051481610fc4565b61051c610fce565b73ffffffffffffffffffffffffffffffffffffffff851660009081526003602052604090205460ff1661057b576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105bc73ffffffffffffffffffffffffffffffffffffffff86167f00000000000000000000000000000000000000000000000000000000000000008661100d565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906106369088908a90899089908990600401611939565b600060405180830381600087803b15801561065057600080fd5b505af1158015610664573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d58686866040516106c993929190611996565b60405180910390a3506106dc6001600055565b5050505050565b600082815260016020819052604090912001546106ff81610fc4565b610709838361108e565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331461075e576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610768828261113b565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61079781610fc4565b61079f6111dc565b50565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107cc81610fc4565b61079f61123b565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6107fe81610fc4565b73ffffffffffffffffffffffffffffffffffffffff821661084b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6108cc81610fc4565b73ffffffffffffffffffffffffffffffffffffffff8216610919576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b61097b610f81565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46109a581610fc4565b6109ad610fce565b73ffffffffffffffffffffffffffffffffffffffff861660009081526003602052604090205460ff16610a0c576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a4d73ffffffffffffffffffffffffffffffffffffffff87167f00000000000000000000000000000000000000000000000000000000000000008761100d565b6040517fd0b492c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610ac99089908b908a908a908a908a90600401611a77565b600060405180830381600087803b158015610ae357600080fd5b505af1158015610af7573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b87878787604051610b5e9493929190611ae8565b60405180910390a350610b716001600055565b505050505050565b60008281526001602081905260409091200154610b9581610fc4565b610709838361113b565b610ba7610f81565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610bd181610fc4565b610bd9610fce565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604090205460ff16610c38576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c5973ffffffffffffffffffffffffffffffffffffffff8416858461100d565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610cb891815260200190565b60405180910390a3506107686001600055565b610cd3610f81565b610cdb610fce565b60045474010000000000000000000000000000000000000000900460ff16610d2f576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841660009081526003602052604090205460ff16610d8e576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015610dfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1f9190611b14565b9050610e4373ffffffffffffffffffffffffffffffffffffffff8616333087611278565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa158015610ed7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610efb9190611b14565b610f059190611b2d565b8787604051610f18959493929190611b67565b60405180910390a250610b716001600055565b6000610f3681610fc4565b506004805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b600260005403610fbd576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61079f81336112be565b60025460ff161561100b576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60405173ffffffffffffffffffffffffffffffffffffffff83811660248301526044820183905261076891859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061134f565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1661113357600083815260016020818152604080842073ffffffffffffffffffffffffffffffffffffffff87168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016104dc565b5060006104dc565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff161561113357600083815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104dc565b6111e46113e5565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b611243610fce565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586112113390565b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526107099186918216906323b872dd90608401611047565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661134b576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044015b60405180910390fd5b5050565b600061137173ffffffffffffffffffffffffffffffffffffffff841683611421565b905080516000141580156113965750808060200190518101906113949190611ba0565b155b15610768576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401611342565b60025460ff1661100b576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606061142f83836000611436565b9392505050565b606081471015611474576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611342565b6000808573ffffffffffffffffffffffffffffffffffffffff16848660405161149d9190611bbd565b60006040518083038185875af1925050503d80600081146114da576040519150601f19603f3d011682016040523d82523d6000602084013e6114df565b606091505b50915091506114ef8683836114f9565b9695505050505050565b60608261150e5761150982611588565b61142f565b8151158015611532575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611581576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401611342565b508061142f565b8051156115985780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156115dc57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461142f57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461079f57600080fd5b60008083601f84011261164057600080fd5b50813567ffffffffffffffff81111561165857600080fd5b60208301915083602082850101111561167057600080fd5b9250929050565b60008060008060006080868803121561168f57600080fd5b853561169a8161160c565b945060208601356116aa8161160c565b935060408601359250606086013567ffffffffffffffff8111156116cd57600080fd5b6116d98882890161162e565b969995985093965092949392505050565b6000602082840312156116fc57600080fd5b5035919050565b6000806040838503121561171657600080fd5b8235915060208301356117288161160c565b809150509250929050565b60006020828403121561174557600080fd5b813561142f8161160c565b60008060008060008060a0878903121561176957600080fd5b86356117748161160c565b955060208701356117848161160c565b945060408701359350606087013567ffffffffffffffff8111156117a757600080fd5b6117b389828a0161162e565b909450925050608087013567ffffffffffffffff8111156117d357600080fd5b87016060818a0312156117e557600080fd5b809150509295509295509295565b60008060006060848603121561180857600080fd5b83356118138161160c565b925060208401356118238161160c565b929592945050506040919091013590565b6000806000806000806080878903121561184d57600080fd5b863567ffffffffffffffff81111561186457600080fd5b61187089828a0161162e565b90975095505060208701356118848161160c565b935060408701359250606087013567ffffffffffffffff8111156118a757600080fd5b6118b389828a0161162e565b979a9699509497509295939492505050565b801515811461079f57600080fd5b6000602082840312156118e557600080fd5b813561142f816118c5565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015283604082015260806060820152600061198b6080830184866118f0565b979650505050505050565b8381526040602082015260006119b06040830184866118f0565b95945050505050565b600081356119c68161160c565b73ffffffffffffffffffffffffffffffffffffffff168352602082013567ffffffffffffffff81168082146119fa57600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112611a3657600080fd5b820160208101903567ffffffffffffffff811115611a5357600080fd5b803603821315611a6257600080fd5b606060408601526119b06060860182846118f0565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a060608201526000611ac960a0830185876118f0565b8281036080840152611adb81856119b9565b9998505050505050505050565b848152606060208201526000611b026060830185876118f0565b828103604084015261198b81856119b9565b600060208284031215611b2657600080fd5b5051919050565b818103818111156104dc577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b606081526000611b7b6060830187896118f0565b8560208401528281036040840152611b948185876118f0565b98975050505050505050565b600060208284031215611bb257600080fd5b815161142f816118c5565b6000825160005b81811015611bde5760208186018101518583015201611bc4565b50600092019182525091905056fea2646970667358221220bd79bcd58fba8e98853080b63f1f2957dbc1b4f08d5077194dba51b74316756764736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a",
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

// Execute0 is a paid mutator transaction binding the contract method 0x38e22527.
//
// Solidity: function execute((address) messageContext, address destination, bytes data) payable returns(bytes)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactor) Execute0(opts *bind.TransactOpts, messageContext MessageContext, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.contract.Transact(opts, "execute0", messageContext, destination, data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x38e22527.
//
// Solidity: function execute((address) messageContext, address destination, bytes data) payable returns(bytes)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestSession) Execute0(messageContext MessageContext, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Execute0(&_GatewayEVMEchidnaTest.TransactOpts, messageContext, destination, data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x38e22527.
//
// Solidity: function execute((address) messageContext, address destination, bytes data) payable returns(bytes)
func (_GatewayEVMEchidnaTest *GatewayEVMEchidnaTestTransactorSession) Execute0(messageContext MessageContext, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMEchidnaTest.Contract.Execute0(&_GatewayEVMEchidnaTest.TransactOpts, messageContext, destination, data)
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
