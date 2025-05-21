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

// AbortContext is an auto generated low-level Go binding around an user-defined struct.
type AbortContext struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}

// CallOptions is an auto generated low-level Go binding around an user-defined struct.
type CallOptions struct {
	GasLimit        *big.Int
	IsArbitraryCall bool
}

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender    []byte
	SenderEVM common.Address
	ChainID   *big.Int
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
	OnRevertGasLimit uint64
}

// GatewayZEVMMetaData contains all meta data concerning the GatewayZEVM contract.
var GatewayZEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_MESSAGE_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndRevert\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndRevert\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeAbort\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"abortContext\",\"type\":\"tuple\",\"internalType\":\"structAbortContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outgoing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallOnRevertNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051613f8d6100fd600039600081816123fb0152818161242401526125d70152613f8d6000f3fe6080604052600436106101e75760003560e01c806366123ee411610102578063ad3cb1cc11610095578063d547741f11610064578063d547741f14610699578063e63ab1e9146106b9578063f3357c1f146106ed578063f45346dc1461070d57600080fd5b8063ad3cb1cc146105e3578063bb4159e014610639578063bcf7f32b14610659578063c39aca371461067957600080fd5b806391d14854116100d157806391d148541461053357806397d340f5146105985780639d4ba465146105ae578063a217fddf146105ce57600080fd5b806366123ee4146104c75780637ce1ffeb146104e75780638456cb59146104fe5780638d096b471461051357600080fd5b806336568abe1161017a5780634f1ef286116101495780634f1ef2861461044857806352d1902d1461045b5780635c975abb146104705780635cf12463146104a757600080fd5b806336568abe146103d35780633f4ba83a146103f357806345122cfd14610408578063485cc9551461042857600080fd5b806321e093b1116101b657806321e093b1146102f6578063248a9ca31461032e5780632722feee1461038b5780632f2ff15d146103b357600080fd5b806301ffc9a714610261578063184b0793146102965780632095dedb146102b657806321501a95146102d657600080fd5b3661025c576101f461072d565b6000546001600160a01b0316331480159061022357503373735b14bb79463307aacbed86daf3322b1e6226ab14155b1561025a576040517fb3af013700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b34801561026d57600080fd5b5061028161027c366004613022565b61078b565b60405190151581526020015b60405180910390f35b3480156102a257600080fd5b5061025a6102b1366004613091565b610824565b3480156102c257600080fd5b5061025a6102d13660046130e1565b61094c565b3480156102e257600080fd5b5061025a6102f1366004613193565b610a15565b34801561030257600080fd5b50600054610316906001600160a01b031681565b6040516001600160a01b03909116815260200161028d565b34801561033a57600080fd5b5061037d61034936600461321f565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b60405190815260200161028d565b34801561039757600080fd5b5061031673735b14bb79463307aacbed86daf3322b1e6226ab81565b3480156103bf57600080fd5b5061025a6103ce366004613238565b610bfa565b3480156103df57600080fd5b5061025a6103ee366004613238565b610c44565b3480156103ff57600080fd5b5061025a610c95565b34801561041457600080fd5b5061025a61042336600461325d565b610cca565b34801561043457600080fd5b5061025a6104433660046132b6565b610e9c565b61025a6104563660046133a2565b6110d9565b34801561046757600080fd5b5061037d6110f4565b34801561047c57600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610281565b3480156104b357600080fd5b5061025a6104c236600461340c565b611123565b3480156104d357600080fd5b5061025a6104e23660046134ca565b611211565b3480156104f357600080fd5b5061037d620186a081565b34801561050a57600080fd5b5061025a611415565b34801561051f57600080fd5b5061025a61052e36600461354f565b611447565b34801561053f57600080fd5b5061028161054e366004613238565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156105a457600080fd5b5061037d61080081565b3480156105ba57600080fd5b5061025a6105c9366004613618565b6115e5565b3480156105da57600080fd5b5061037d600081565b3480156105ef57600080fd5b5061062c6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161028d91906136c8565b34801561064557600080fd5b5061025a6106543660046136db565b61181c565b34801561066557600080fd5b5061025a61067436600461373f565b611930565b34801561068557600080fd5b5061025a61069436600461373f565b611a70565b3480156106a557600080fd5b5061025a6106b4366004613238565b611c07565b3480156106c557600080fd5b5061037d7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b3480156106f957600080fd5b5061025a6107083660046137dd565b611c4b565b34801561071957600080fd5b5061025a610728366004613846565b611d99565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615610789576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061081e57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b61082c611f28565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610879576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61088161072d565b6001600160a01b0382166108a85760405163d92e233d60e01b815260040160405180910390fd5b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a36906108ed908490600401613917565b600060405180830381600087803b15801561090757600080fd5b505af115801561091b573d6000803e3d6000fd5b5050505061094860017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050565b610954611f28565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109a1576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109a961072d565b6001600160a01b0382166109d05760405163d92e233d60e01b815260040160405180910390fd5b6040517f2d4cfb7e0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690632d4cfb7e906108ed908490600401613995565b610a1d611f28565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a6a576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a7261072d565b6001600160a01b038316610a995760405163d92e233d60e01b815260040160405180910390fd5b83600003610ad3576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab1480610b0657506001600160a01b03831630145b15610b3d576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b478484611fcf565b6000546040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b0380861692635bcfd61692610b98928a921690899088908890600401613a3c565b600060405180830381600087803b158015610bb257600080fd5b505af1158015610bc6573d6000803e3d6000fd5b50505050610bf360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610c348161217d565b610c3e8383612187565b50505050565b6001600160a01b0381163314610c86576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c908282612274565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610cbf8161217d565b610cc7612338565b50565b610cd2611f28565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610d1f576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d2761072d565b6001600160a01b038216610d4e5760405163d92e233d60e01b815260040160405180910390fd5b82600003610d88576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03821673735b14bb79463307aacbed86daf3322b1e6226ab1480610dbb57506001600160a01b03821630145b15610df2576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610dfc8383611fcf565b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a3690610e41908490600401613917565b600060405180830381600087803b158015610e5b57600080fd5b505af1158015610e6f573d6000803e3d6000fd5b50505050610c9060017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610ee75750825b905060008267ffffffffffffffff166001148015610f045750303b155b905081158015610f12575080155b15610f49576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610faa5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b0387161580610fc757506001600160a01b038616155b15610fe55760405163d92e233d60e01b815260040160405180910390fd5b610fed6123c8565b610ff56123c8565b610ffd6123d0565b6110056123e0565b611010600087612187565b5061103b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87612187565b50600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03891617905583156110d05784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b6110e16123f0565b6110ea826124c0565b61094882826124cb565b60006110fe6125cc565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b61112b61072d565b620186a082351015611169576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108006111796060830183613ab9565b611184915085613b1e565b11156111e5576111976060820182613ab9565b6111a2915084613b1e565b6040517fcd6f4e6d000000000000000000000000000000000000000000000000000000008152600481019190915261080060248201526044015b60405180910390fd5b611209868686866111fb36889003880188613b58565b61120487613bcd565b61262e565b505050505050565b61121961072d565b835160000361123b5760405163d92e233d60e01b815260040160405180910390fd5b82600003611275576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108006112856060830183613ab9565b905011156112da5761129a6060820182613ab9565b6040517fcd6f4e6d0000000000000000000000000000000000000000000000000000000081526111dc925061080090600401918252602082015260400190565b60006112e684846127aa565b90506000336001600160a01b03167fd2cb4a520b9303966315ef36d0910a0391a7ed5af8e512e9609874f68cf59a6887868886896001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611357573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061137b9190613c7b565b60405180604001604052808c6001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113e89190613c7b565b81526001602090910152604051611406969594939291908c90613d27565b60405180910390a35050505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61143f8161217d565b610cc761281f565b61144f61072d565b86516000036114715760405163d92e233d60e01b815260040160405180910390fd5b856000036114ab576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620186a0823510156114e9576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108006114f96060830183613ab9565b611504915085613b1e565b1115611517576111976060820182613ab9565b600061152587878535612898565b90506000336001600160a01b03167ff094c41f4246a1a5f62ffecf62e8aed17a29bf0f86dfd2183f8e6dc676ebc2278a898b868c6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611596573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ba9190613c7b565b8c8c8c8c6040516115d399989796959493929190613dc8565b60405180910390a35050505050505050565b6115ed611f28565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461163a576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61164261072d565b6001600160a01b038416158061165f57506001600160a01b038216155b1561167d5760405163d92e233d60e01b815260040160405180910390fd5b826000036116b7576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03821673735b14bb79463307aacbed86daf3322b1e6226ab14806116ea57506001600160a01b03821630145b15611721576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61172c848385612a4c565b61177c576040517f40a143ba0000000000000000000000000000000000000000000000000000000081526001600160a01b03808616600483015283166024820152604481018490526064016111dc565b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a36906117c1908490600401613917565b600060405180830381600087803b1580156117db57600080fd5b505af11580156117ef573d6000803e3d6000fd5b50505050610c3e60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b61182461072d565b83516000036118465760405163d92e233d60e01b815260040160405180910390fd5b82600003611880576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108006118906060830183613ab9565b905011156118a55761129a6060820182613ab9565b6118c38373735b14bb79463307aacbed86daf3322b1e6226ab611fcf565b60008054604080518082018252838152600160208201529051859333937fd2cb4a520b9303966315ef36d0910a0391a7ed5af8e512e9609874f68cf59a6893611922938b936001600160a01b03909316928b9290918291908b90613d27565b60405180910390a350505050565b611938611f28565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611985576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61198d61072d565b6001600160a01b03851615806119aa57506001600160a01b038316155b156119c85760405163d92e233d60e01b815260040160405180910390fd5b6040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b03841690635bcfd61690611a159089908990899088908890600401613a3c565b600060405180830381600087803b158015611a2f57600080fd5b505af1158015611a43573d6000803e3d6000fd5b5050505061120960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b611a78611f28565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611ac5576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611acd61072d565b6001600160a01b0385161580611aea57506001600160a01b038316155b15611b085760405163d92e233d60e01b815260040160405180910390fd5b83600003611b42576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab1480611b7557506001600160a01b03831630145b15611bac576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bb7858486612a4c565b6119c8576040517f40a143ba0000000000000000000000000000000000000000000000000000000081526001600160a01b03808716600483015284166024820152604481018590526064016111dc565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611c418161217d565b610c3e8383612274565b611c5361072d565b8651600003611c755760405163d92e233d60e01b815260040160405180910390fd5b85600003611caf576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620186a082351015611ced576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610800611cfd6060830183613ab9565b611d08915085613b1e565b1115611d1b576111976060820182613ab9565b611d398673735b14bb79463307aacbed86daf3322b1e6226ab611fcf565b60008054604051879233927ff094c41f4246a1a5f62ffecf62e8aed17a29bf0f86dfd2183f8e6dc676ebc22792611d88928d926001600160a01b0316918d919081908d908d908d908d90613dc8565b60405180910390a350505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611de6576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611dee61072d565b6001600160a01b0383161580611e0b57506001600160a01b038116155b15611e295760405163d92e233d60e01b815260040160405180910390fd5b81600003611e63576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03811673735b14bb79463307aacbed86daf3322b1e6226ab1480611e9657506001600160a01b03811630145b15611ecd576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ed8838284612a4c565b610c90576040517f40a143ba0000000000000000000000000000000000000000000000000000000081526001600160a01b03808516600483015282166024820152604481018390526064016111dc565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611fa3576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b600054611fe7906001600160a01b0316333085612ae7565b612026576040517ff279af7e000000000000000000000000000000000000000000000000000000008152306004820152602481018390526044016111dc565b6000546040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b0390911690632e1a7d4d90602401600060405180830381600087803b15801561208557600080fd5b505af1925050508015612096575060015b6120de576040517ff279af7e0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044016111dc565b6000816001600160a01b03168360405160006040518083038185875af1925050503d806000811461212b576040519150601f19603f3d011682016040523d82523d6000602084013e612130565b606091505b5050905080610c90576040517ff279af7e0000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602481018490526044016111dc565b610cc78133612b8e565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1661226a576000848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556122203390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061081e565b600091505061081e565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff161561226a576000848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061081e565b612340612c1b565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b610789612c76565b6123d8612c76565b610789612cdd565b6123e8612c76565b610789612d2e565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061248957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661247d7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15610789576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006109488161217d565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612525575060408051601f3d908101601f1916820190925261252291810190613c7b565b60015b612566576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016111dc565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146125c2576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016111dc565b610c908383612d36565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610789576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b85516000036126505760405163d92e233d60e01b815260040160405180910390fd5b81516040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152600481019190915260009081906001600160a01b0388169063fc5fecd5906024016040805180830381865afa1580156126b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126d99190613e43565b915091506126fd823373735b14bb79463307aacbed86daf3322b1e6226ab84612ae7565b61275f576040517f667011120000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015273735b14bb79463307aacbed86daf3322b1e6226ab6024820152604481018290526064016111dc565b866001600160a01b0316336001600160a01b03167fdc8a693f0b4e37ea67f2aea353fe747ffa7cb831b3379f4f026c3d0ff7f325338a898989896040516115d3959493929190613e71565b60006128188383846001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa1580156127ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128139190613c7b565b612898565b9392505050565b61282761072d565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336123aa565b6000806000846001600160a01b031663fc5fecd5856040518263ffffffff1660e01b81526004016128cb91815260200190565b6040805180830381865afa1580156128e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061290b9190613e43565b9150915061292f823373735b14bb79463307aacbed86daf3322b1e6226ab84612ae7565b612991576040517f667011120000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015273735b14bb79463307aacbed86daf3322b1e6226ab6024820152604481018290526064016111dc565b61299d85333089612ae7565b6129f1576040517f489ca9b70000000000000000000000000000000000000000000000000000000081526001600160a01b0386166004820152336024820152306044820152606481018790526084016111dc565b6129fb8587612d8c565b612a43576040517f7112ae770000000000000000000000000000000000000000000000000000000081526001600160a01b0386166004820152602481018790526044016111dc565b95945050505050565b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b03838116600483015260248201839052600091908516906347e7ef24906044016020604051808303816000875af1925050508015612ad4575060408051601f3d908101601f19168201909252612ad191810190613f1e565b60015b612ae057506000612818565b9050612818565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081526001600160a01b038481166004830152838116602483015260448201839052600091908616906323b872dd906064016020604051808303816000875af1925050508015612b77575060408051601f3d908101601f19168201909252612b7491810190613f1e565b60015b612b8357506000612b86565b90505b949350505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610948576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044016111dc565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610789576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610789576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612ce5612c76565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b611fa9612c76565b612d3f82612e1e565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115612d8457610c908282612ec6565b610948612f33565b6040517f42966c68000000000000000000000000000000000000000000000000000000008152600481018290526000906001600160a01b038416906342966c68906024016020604051808303816000875af1925050508015612e0b575060408051601f3d908101601f19168201909252612e0891810190613f1e565b60015b612e175750600061081e565b905061081e565b806001600160a01b03163b600003612e6d576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016111dc565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051612ee39190613f3b565b600060405180830381855af49150503d8060008114612f1e576040519150601f19603f3d011682016040523d82523d6000602084013e612f23565b606091505b5091509150612a43858383612f6b565b3415610789576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082612f8057612f7b82612fe0565b612818565b8151158015612f9757506001600160a01b0384163b155b15612fd9576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016111dc565b5080612818565b805115612ff05780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561303457600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461281857600080fd5b6001600160a01b0381168114610cc757600080fd5b60006080828403121561308b57600080fd5b50919050565b600080604083850312156130a457600080fd5b82356130af81613064565b9150602083013567ffffffffffffffff8111156130cb57600080fd5b6130d785828601613079565b9150509250929050565b600080604083850312156130f457600080fd5b82356130ff81613064565b9150602083013567ffffffffffffffff81111561311b57600080fd5b830160c0818603121561312d57600080fd5b809150509250929050565b60006060828403121561308b57600080fd5b60008083601f84011261315c57600080fd5b50813567ffffffffffffffff81111561317457600080fd5b60208301915083602082850101111561318c57600080fd5b9250929050565b6000806000806000608086880312156131ab57600080fd5b853567ffffffffffffffff8111156131c257600080fd5b6131ce88828901613138565b9550506020860135935060408601356131e681613064565b9250606086013567ffffffffffffffff81111561320257600080fd5b61320e8882890161314a565b969995985093965092949392505050565b60006020828403121561323157600080fd5b5035919050565b6000806040838503121561324b57600080fd5b82359150602083013561312d81613064565b60008060006060848603121561327257600080fd5b83359250602084013561328481613064565b9150604084013567ffffffffffffffff8111156132a057600080fd5b6132ac86828701613079565b9150509250925092565b600080604083850312156132c957600080fd5b82356132d481613064565b9150602083013561312d81613064565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261332457600080fd5b813567ffffffffffffffff81111561333e5761333e6132e4565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561336d5761336d6132e4565b60405281815283820160200185101561338557600080fd5b816020850160208301376000918101602001919091529392505050565b600080604083850312156133b557600080fd5b82356133c081613064565b9150602083013567ffffffffffffffff8111156133dc57600080fd5b6130d785828601613313565b60006040828403121561308b57600080fd5b600060a0828403121561308b57600080fd5b60008060008060008060c0878903121561342557600080fd5b863567ffffffffffffffff81111561343c57600080fd5b61344889828a01613313565b965050602087013561345981613064565b9450604087013567ffffffffffffffff81111561347557600080fd5b61348189828a0161314a565b9095509350613495905088606089016133e8565b915060a087013567ffffffffffffffff8111156134b157600080fd5b6134bd89828a016133fa565b9150509295509295509295565b600080600080608085870312156134e057600080fd5b843567ffffffffffffffff8111156134f757600080fd5b61350387828801613313565b94505060208501359250604085013561351b81613064565b9150606085013567ffffffffffffffff81111561353757600080fd5b613543878288016133fa565b91505092959194509250565b600080600080600080600060e0888a03121561356a57600080fd5b873567ffffffffffffffff81111561358157600080fd5b61358d8a828b01613313565b9750506020880135955060408801356135a581613064565b9450606088013567ffffffffffffffff8111156135c157600080fd5b6135cd8a828b0161314a565b90955093506135e190508960808a016133e8565b915060c088013567ffffffffffffffff8111156135fd57600080fd5b6136098a828b016133fa565b91505092959891949750929550565b6000806000806080858703121561362e57600080fd5b843561363981613064565b935060208501359250604085013561365081613064565b9150606085013567ffffffffffffffff81111561366c57600080fd5b61354387828801613079565b60005b8381101561369357818101518382015260200161367b565b50506000910152565b600081518084526136b4816020860160208601613678565b601f01601f19169290920160200192915050565b602081526000612818602083018461369c565b600080600080608085870312156136f157600080fd5b843567ffffffffffffffff81111561370857600080fd5b61371487828801613313565b9450506020850135925060408501359150606085013567ffffffffffffffff81111561353757600080fd5b60008060008060008060a0878903121561375857600080fd5b863567ffffffffffffffff81111561376f57600080fd5b61377b89828a01613138565b965050602087013561378c81613064565b94506040870135935060608701356137a381613064565b9250608087013567ffffffffffffffff8111156137bf57600080fd5b6137cb89828a0161314a565b979a9699509497509295939492505050565b600080600080600080600060e0888a0312156137f857600080fd5b873567ffffffffffffffff81111561380f57600080fd5b61381b8a828b01613313565b9750506020880135955060408801359450606088013567ffffffffffffffff8111156135c157600080fd5b60008060006060848603121561385b57600080fd5b833561386681613064565b925060208401359150604084013561387d81613064565b809150509250925092565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126138bd57600080fd5b830160208101925035905067ffffffffffffffff8111156138dd57600080fd5b80360382131561318c57600080fd5b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b602081526000823561392881613064565b6001600160a01b038116602084015250602083013561394681613064565b6001600160a01b038116604084015250600060408401359050806060840152506139736060840184613888565b608080850152612a4360a0850182846138ec565b8015158114610cc757600080fd5b6020815260006139a58384613888565b60c060208501526139ba60e0850182846138ec565b91505060208401356139cb81613064565b6001600160a01b03166040848101919091528401356060808501919091528401356139f581613987565b8015156080850152506000608085013590508060a085015250613a1b60a0850185613888565b601f198584030160c0860152613a328382846138ec565b9695505050505050565b608081526000613a4c8788613888565b60606080850152613a6160e0850182846138ec565b9150506020880135613a7281613064565b6001600160a01b0390811660a085015260408981013560c0860152908816602085015283018690528281036060840152613aad8185876138ec565b98975050505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613aee57600080fd5b83018035915067ffffffffffffffff821115613b0957600080fd5b60200191503681900382131561318c57600080fd5b8082018082111561081e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006040828403128015613b6b57600080fd5b506040805190810167ffffffffffffffff81118282101715613b8f57613b8f6132e4565b604052823581526020830135613ba481613987565b60208201529392505050565b803567ffffffffffffffff81168114613bc857600080fd5b919050565b600060a08236031215613bdf57600080fd5b60405160a0810167ffffffffffffffff81118282101715613c0257613c026132e4565b6040528235613c1081613064565b81526020830135613c2081613987565b60208201526040830135613c3381613064565b6040820152606083013567ffffffffffffffff811115613c5257600080fd5b613c5e36828601613313565b606083015250613c7060808401613bb0565b608082015292915050565b600060208284031215613c8d57600080fd5b5051919050565b60008135613ca181613064565b6001600160a01b031683526020820135613cba81613987565b151560208401526040820135613ccf81613064565b6001600160a01b03166040840152613cea6060830183613888565b60a06060860152613cff60a0860182846138ec565b91505067ffffffffffffffff613d1760808501613bb0565b1660808501528091505092915050565b61012081526000613d3c61012083018a61369c565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a085015260008252613d8360c0850187805182526020908101511515910152565b6020810161010085015250613d9b6020820185613c94565b9a9950505050505050505050565b803582526020810135613dbb81613987565b8015156020840152505050565b61012081526000613ddd61012083018c61369c565b6001600160a01b038b16602084015289604084015288606084015287608084015282810360a0840152613e118187896138ec565b9050613e2060c0840186613da9565b828103610100840152613e338185613c94565b9c9b505050505050505050505050565b60008060408385031215613e5657600080fd5b8251613e6181613064565b6020939093015192949293505050565b60a081526000613e8460a083018861369c565b8281036020840152613e978187896138ec565b85516040850152602086015115156060850152905082810360808401526001600160a01b0384511681526020840151151560208201526001600160a01b036040850151166040820152606084015160a06060830152613ef960a083018261369c565b905067ffffffffffffffff608086015116608083015280925050509695505050505050565b600060208284031215613f3057600080fd5b815161281881613987565b60008251613f4d818460208701613678565b919091019291505056fea2646970667358221220981ac0d71eb001ed464aff5648579116bbd11e0c5c63a85d72b46b58e6eda26764736f6c634300081a0033",
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

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVM *GatewayZEVMCaller) MAXMESSAGESIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "MAX_MESSAGE_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVM *GatewayZEVMSession) MAXMESSAGESIZE() (*big.Int, error) {
	return _GatewayZEVM.Contract.MAXMESSAGESIZE(&_GatewayZEVM.CallOpts)
}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVM *GatewayZEVMCallerSession) MAXMESSAGESIZE() (*big.Int, error) {
	return _GatewayZEVM.Contract.MAXMESSAGESIZE(&_GatewayZEVM.CallOpts)
}

// MINGASLIMIT is a free data retrieval call binding the contract method 0x7ce1ffeb.
//
// Solidity: function MIN_GAS_LIMIT() view returns(uint256)
func (_GatewayZEVM *GatewayZEVMCaller) MINGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "MIN_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINGASLIMIT is a free data retrieval call binding the contract method 0x7ce1ffeb.
//
// Solidity: function MIN_GAS_LIMIT() view returns(uint256)
func (_GatewayZEVM *GatewayZEVMSession) MINGASLIMIT() (*big.Int, error) {
	return _GatewayZEVM.Contract.MINGASLIMIT(&_GatewayZEVM.CallOpts)
}

// MINGASLIMIT is a free data retrieval call binding the contract method 0x7ce1ffeb.
//
// Solidity: function MIN_GAS_LIMIT() view returns(uint256)
func (_GatewayZEVM *GatewayZEVMCallerSession) MINGASLIMIT() (*big.Int, error) {
	return _GatewayZEVM.Contract.MINGASLIMIT(&_GatewayZEVM.CallOpts)
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

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_GatewayZEVM *GatewayZEVMCaller) PROTOCOLADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "PROTOCOL_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_GatewayZEVM *GatewayZEVMSession) PROTOCOLADDRESS() (common.Address, error) {
	return _GatewayZEVM.Contract.PROTOCOLADDRESS(&_GatewayZEVM.CallOpts)
}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_GatewayZEVM *GatewayZEVMCallerSession) PROTOCOLADDRESS() (common.Address, error) {
	return _GatewayZEVM.Contract.PROTOCOLADDRESS(&_GatewayZEVM.CallOpts)
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

// Call is a paid mutator transaction binding the contract method 0x5cf12463.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Call(opts *bind.TransactOpts, receiver []byte, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "call", receiver, zrc20, message, callOptions, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x5cf12463.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) Call(receiver []byte, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Call(&_GatewayZEVM.TransactOpts, receiver, zrc20, message, callOptions, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x5cf12463.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Call(receiver []byte, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Call(&_GatewayZEVM.TransactOpts, receiver, zrc20, message, callOptions, revertOptions)
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
func (_GatewayZEVM *GatewayZEVMTransactor) DepositAndCall(opts *bind.TransactOpts, context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "depositAndCall", context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMSession) DepositAndCall(context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndCall(&_GatewayZEVM.TransactOpts, context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) DepositAndCall(context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndCall(&_GatewayZEVM.TransactOpts, context, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) DepositAndCall0(opts *bind.TransactOpts, context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "depositAndCall0", context, zrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMSession) DepositAndCall0(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndCall0(&_GatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) DepositAndCall0(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndCall0(&_GatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x45122cfd.
//
// Solidity: function depositAndRevert(uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) DepositAndRevert(opts *bind.TransactOpts, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "depositAndRevert", amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x45122cfd.
//
// Solidity: function depositAndRevert(uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMSession) DepositAndRevert(amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndRevert(&_GatewayZEVM.TransactOpts, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x45122cfd.
//
// Solidity: function depositAndRevert(uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) DepositAndRevert(amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndRevert(&_GatewayZEVM.TransactOpts, amount, target, revertContext)
}

// DepositAndRevert0 is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) DepositAndRevert0(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "depositAndRevert0", zrc20, amount, target, revertContext)
}

// DepositAndRevert0 is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMSession) DepositAndRevert0(zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndRevert0(&_GatewayZEVM.TransactOpts, zrc20, amount, target, revertContext)
}

// DepositAndRevert0 is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) DepositAndRevert0(zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndRevert0(&_GatewayZEVM.TransactOpts, zrc20, amount, target, revertContext)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Execute(opts *bind.TransactOpts, context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "execute", context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMSession) Execute(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Execute(&_GatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Execute(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Execute(&_GatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// ExecuteAbort is a paid mutator transaction binding the contract method 0x2095dedb.
//
// Solidity: function executeAbort(address target, (bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) ExecuteAbort(opts *bind.TransactOpts, target common.Address, abortContext AbortContext) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "executeAbort", target, abortContext)
}

// ExecuteAbort is a paid mutator transaction binding the contract method 0x2095dedb.
//
// Solidity: function executeAbort(address target, (bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_GatewayZEVM *GatewayZEVMSession) ExecuteAbort(target common.Address, abortContext AbortContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.ExecuteAbort(&_GatewayZEVM.TransactOpts, target, abortContext)
}

// ExecuteAbort is a paid mutator transaction binding the contract method 0x2095dedb.
//
// Solidity: function executeAbort(address target, (bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) ExecuteAbort(target common.Address, abortContext AbortContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.ExecuteAbort(&_GatewayZEVM.TransactOpts, target, abortContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) ExecuteRevert(opts *bind.TransactOpts, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "executeRevert", target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.ExecuteRevert(&_GatewayZEVM.TransactOpts, target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
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

// Withdraw is a paid mutator transaction binding the contract method 0x66123ee4.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Withdraw(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdraw", receiver, amount, zrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x66123ee4.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x66123ee4.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xbb4159e0.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Withdraw0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdraw0", receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xbb4159e0.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw0(&_GatewayZEVM.TransactOpts, receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xbb4159e0.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw0(&_GatewayZEVM.TransactOpts, receiver, amount, chainId, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x8d096b47.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) WithdrawAndCall(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdrawAndCall", receiver, amount, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x8d096b47.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) WithdrawAndCall(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.WithdrawAndCall(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x8d096b47.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) WithdrawAndCall(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.WithdrawAndCall(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0xf3357c1f.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) WithdrawAndCall0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdrawAndCall0", receiver, amount, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0xf3357c1f.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) WithdrawAndCall0(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.WithdrawAndCall0(&_GatewayZEVM.TransactOpts, receiver, amount, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0xf3357c1f.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) WithdrawAndCall0(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.WithdrawAndCall0(&_GatewayZEVM.TransactOpts, receiver, amount, chainId, message, callOptions, revertOptions)
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
	CallOptions   CallOptions
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xdc8a693f0b4e37ea67f2aea353fe747ffa7cb831b3379f4f026c3d0ff7f32533.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions)
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

// WatchCalled is a free log subscription operation binding the contract event 0xdc8a693f0b4e37ea67f2aea353fe747ffa7cb831b3379f4f026c3d0ff7f32533.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions)
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

// ParseCalled is a log parse operation binding the contract event 0xdc8a693f0b4e37ea67f2aea353fe747ffa7cb831b3379f4f026c3d0ff7f32533.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions)
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
	CallOptions     CallOptions
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd2cb4a520b9303966315ef36d0910a0391a7ed5af8e512e9609874f68cf59a68.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions)
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

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd2cb4a520b9303966315ef36d0910a0391a7ed5af8e512e9609874f68cf59a68.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions)
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd2cb4a520b9303966315ef36d0910a0391a7ed5af8e512e9609874f68cf59a68.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) ParseWithdrawn(log types.Log) (*GatewayZEVMWithdrawn, error) {
	event := new(GatewayZEVMWithdrawn)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayZEVM contract.
type GatewayZEVMWithdrawnAndCalledIterator struct {
	Event *GatewayZEVMWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMWithdrawnAndCalled)
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
		it.Event = new(GatewayZEVMWithdrawnAndCalled)
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
func (it *GatewayZEVMWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayZEVM contract.
type GatewayZEVMWithdrawnAndCalled struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Zrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	CallOptions     CallOptions
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0xf094c41f4246a1a5f62ffecf62e8aed17a29bf0f86dfd2183f8e6dc676ebc227.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMWithdrawnAndCalledIterator{contract: _GatewayZEVM.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xf094c41f4246a1a5f62ffecf62e8aed17a29bf0f86dfd2183f8e6dc676ebc227.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMWithdrawnAndCalled)
				if err := _GatewayZEVM.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0xf094c41f4246a1a5f62ffecf62e8aed17a29bf0f86dfd2183f8e6dc676ebc227.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint64) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayZEVMWithdrawnAndCalled, error) {
	event := new(GatewayZEVMWithdrawnAndCalled)
	if err := _GatewayZEVM.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
