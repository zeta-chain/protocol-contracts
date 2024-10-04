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

// StdInvariantFuzzArtifactSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzArtifactSelector struct {
	Artifact  string
	Selectors [][4]byte
}

// StdInvariantFuzzInterface is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzInterface struct {
	Addr      common.Address
	Artifacts []string
}

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}

// GatewayEVMTestMetaData contains all meta data concerning the GatewayEVMTest contract.
var GatewayEVMTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ASSET_HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfDestinationIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevert\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertFailsIfSenderIsNotTSS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteWithERC20FailsIfNotCustodyOrConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteWithMsgContextFailsIfDestinationIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNoParams\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNoParamsTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNonPayable\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveOnCallFails\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveOnCallUsingAuthCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveOnRevertFails\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceivePayable\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertWithERC20FailsIfNotCustodyOrConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetConnectorFailsIfSenderIsNotAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetConnectorFailsIfSet\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetConnectorFailsIfZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetCustodyFailsIfSenderIsNotAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetCustodyFailsIfSet\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetCustodyFailsIfZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgrade\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfSenderIsNotTSSUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedCustodyTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5062011d6a806200003e6000396000f3fe608060405234801561001057600080fd5b50600436106102e95760003560e01c806385f438c111610191578063ce5871e1116100e3578063e63ab1e911610097578063fa18c09b11610071578063fa18c09b14610514578063fa7626d41461051c578063fe7bdbb21461052957600080fd5b8063e63ab1e9146104dd578063e6afc79014610504578063f68bd1c01461050c57600080fd5b8063d38b66cd116100c8578063d38b66cd146104c5578063dd51e82f146104cd578063e20c9f71146104d557600080fd5b8063ce5871e1146104b5578063cebad2a6146104bd57600080fd5b8063a783c78911610145578063b5508aa91161011f578063b5508aa91461048d578063ba414fa614610495578063ccf20616146104ad57600080fd5b8063a783c78914610456578063b0464fdc1461047d578063b124dbf51461048557600080fd5b8063a217fddf11610176578063a217fddf1461043e578063a397ffd214610446578063a56f7a4b1461044e57600080fd5b806385f438c114610402578063916a17c61461042957600080fd5b806343fd8c7d1161024a5780635d62c860116101fe5780637d7f772a116101d85780637d7f772a146103dd5780637ebf744f146103e557806385226c81146103ed57600080fd5b80635d62c8601461038b57806366d9a9a0146103c05780636bdd212b146103d557600080fd5b80634df42da11161022f5780634df42da11461037357806351010e491461037b57806352ff59391461038357600080fd5b806343fd8c7d1461036357806344671b941461036b57600080fd5b80631855c337116102a15780632ade3880116102865780632ade38801461033e5780633e5e3c23146103535780633f7286f41461035b57600080fd5b80631855c337146103185780631ed7831c1461032057600080fd5b80630a9254e4116102d25780630a9254e4146103005780631226c65514610308578063130daf591461031057600080fd5b806304b974f9146102ee578063070f2ad0146102f8575b600080fd5b6102f6610531565b005b6102f661073e565b6102f661093f565b6102f661133b565b6102f66114ac565b6102f6611673565b6103286117e8565b6040516103359190619cf4565b60405180910390f35b61034661184a565b6040516103359190619d90565b61032861198c565b6103286119ec565b6102f6611a4c565b6102f6611dbe565b6102f66120c0565b6102f6612231565b6102f661240d565b6103b27f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b981565b604051908152602001610335565b6103c861296e565b6040516103359190619ef6565b6102f6612af0565b6102f6612cc0565b6102f6612fd5565b6103f561314a565b6040516103359190619ff0565b6103b27f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b61043161321a565b604051610335919061a003565b6103b2600081565b6102f6613315565b6102f6613651565b6103b27f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6104316137a3565b6102f661389e565b6103f5613b45565b61049d613c15565b6040519015158152602001610335565b6102f6613ce9565b6102f6613e3b565b6102f6613fac565b6102f66141d0565b6102f661432b565b610328614b20565b6103b27f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6102f6614b80565b6102f6614dee565b6102f66151e3565b601f5461049d9060ff1681565b6102f6615545565b6040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed7016900000000000000000000000000000000000000000000000000000000179052602854925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa7926105d8926001600160a01b031691016001600160a01b0391909116815260200190565b600060405180830381600087803b1580156105f257600080fd5b505af1158015610606573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561068f57600080fd5b505af11580156106a3573d6000803e3d6000fd5b50506020546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250631cff79cd91506106f390600090859060040161a09a565b6000604051808303816000875af1158015610712573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261073a919081019061a1a4565b5050565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156107b057600080fd5b505af11580156107c4573d6000803e3d6000fd5b5050602854604080516001600160a01b039092166024830152600060448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250610890919060040161a1d9565b600060405180830381600087803b1580156108aa57600080fd5b505af11580156108be573d6000803e3d6000fd5b50506020546026546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063950837aa91506024015b600060405180830381600087803b15801561092557600080fd5b505af1158015610939573d6000803e3d6000fd5b50505050565b602680547fffffffffffffffffffffffff000000000000000000000000000000000000000090811630179091556027805482166112341790556028805490911661567817905560405161099190619bf5565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f080158015610a16573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556028546040519116908190610a5f90619c03565b6001600160a01b03928316815291166020820152604001604051809103906000f080158015610a92573d6000803e3d6000fd5b50602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260285460265492519085166024820152604481019390935292166064820152610b81919060840160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b000000000000000000000000000000000000000000000000000000001790526159ba565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff00000000000000000000000000000000000000009091168117909155602854602654604051929391821692911690610c0d90619c11565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610c49573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556020546025546028546026546040519385169492831693918316921690610ca490619c1f565b6001600160a01b039485168152928416602084015290831660408301529091166060820152608001604051809103906000f080158015610ce8573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392831617905560285460405163ca669fa760e01b815291166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015610d6d57600080fd5b505af1158015610d81573d6000803e3d6000fd5b50506025546028546023546040517f15d57fd40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529082166024820152911692506315d57fd49150604401600060405180830381600087803b158015610df257600080fd5b505af1158015610e06573d6000803e3d6000fd5b50505050604051610e1690619c2d565b604051809103906000f080158015610e32573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556028546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b158015610ede57600080fd5b505af1158015610ef2573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015610f6857600080fd5b505af1158015610f7c573d6000803e3d6000fd5b50506020546022546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f9150602401600060405180830381600087803b158015610fe257600080fd5b505af1158015610ff6573d6000803e3d6000fd5b50506020546023546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef9150602401600060405180830381600087803b15801561105c57600080fd5b505af1158015611070573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156110d257600080fd5b505af11580156110e6573d6000803e3d6000fd5b5050602480546026546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240938101939093521692506340c10f199150604401600060405180830381600087803b15801561115757600080fd5b505af115801561116b573d6000803e3d6000fd5b5050602480546022546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526207a1209381019390935216925063a9059cbb91506044016020604051808303816000875af11580156111e1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611205919061a1ec565b506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b15801561128657600080fd5b505af115801561129a573d6000803e3d6000fd5b5050604080516080810182526026546001600160a01b039081168252602454811660208084019182526001848601908152855191820190955260008152606084018190528351602d80549185167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161781559251602e8054919095169116179092559251602f55909350909150603090611336908261a2a2565b505050565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156113ad57600080fd5b505af11580156113c1573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561144a57600080fd5b505af115801561145e573d6000803e3d6000fd5b50506020546040517fae7a3a6f000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b03909116925063ae7a3a6f915060240161090b565b6040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f5ac1e07000000000000000000000000000000000000000000000000000000000179052602854925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa792611553926001600160a01b031691016001600160a01b0391909116815260200190565b600060405180830381600087803b15801561156d57600080fd5b505af1158015611581573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527ff3459a96000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b600060405180830381600087803b15801561160b57600080fd5b505af115801561161f573d6000803e3d6000fd5b50506020546021546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd93506106f39290911690859060040161a09a565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156116e557600080fd5b505af11580156116f9573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fb337f378000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b600060405180830381600087803b15801561178357600080fd5b505af1158015611797573d6000803e3d6000fd5b50506020546022546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f915060240161090b565b6060601680548060200260200160405190810160405280929190818152602001828054801561184057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611822575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020016000905b8282101561198357600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b8282101561196c5783829060005260206000200180546118df9061a20e565b80601f016020809104026020016040519081016040528092919081815260200182805461190b9061a20e565b80156119585780601f1061192d57610100808354040283529160200191611958565b820191906000526020600020905b81548152906001019060200180831161193b57829003601f168201915b5050505050815260200190600101906118c0565b50505050815250508152602001906001019061186e565b50505050905090565b60606018805480602002602001604051908101604052809291908181526020018280548015611840576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611822575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015611840576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611822575050505050905090565b604080516001808252818301909252600091816020015b6060815260200190600190039081611a635790505090506040518060400160405280600f81526020017f48656c6c6f2c20466f756e64727921000000000000000000000000000000000081525081600081518110611ac357611ac361a361565b6020908102919091010152604080516001808252818301909252600091816020016020820280368337019050509050602a81600081518110611b0757611b0761a361565b6020908102919091010152604051600190600090611b2d9085908590859060240161a3c2565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff05b6abf00000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611be057600080fd5b505af1158015611bf4573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611cdf919060040161a1d9565b600060405180830381600087803b158015611cf957600080fd5b505af1158015611d0d573d6000803e3d6000fd5b505060208054604080519283018152610123835260215490517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b0392831695506338e225279450611d70939290911690869060040161a3fa565b6000604051808303816000875af1158015611d8f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611db7919081019061a1a4565b5050505050565b6040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed701690000000000000000000000000000000000000000000000000000000017905260215492517ff30c7ba30000000000000000000000000000000000000000000000000000000081529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba392611e74926001600160a01b03169160009187910161a42c565b600060405180830381600087803b158015611e8e57600080fd5b505af1158015611ea2573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015611f3457600080fd5b505af1158015611f48573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0935001905060405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561201a57600080fd5b505af115801561202e573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f915061207490600090859061a454565b60405180910390a260285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024016115f1565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561213257600080fd5b505af1158015612146573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156121cf57600080fd5b505af11580156121e3573d6000803e3d6000fd5b50506020546040517f950837aa000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b03909116925063950837aa915060240161090b565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152670de0b6b3a76400009060009060250160408051808303601f190181529082905260285463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156122d157600080fd5b505af11580156122e5573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561236e57600080fd5b505af1158015612382573d6000803e3d6000fd5b50506020546040517fcb7ba8e50000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063cb7ba8e5915084906123d7906000908690602d9060040161a54a565b6000604051808303818588803b1580156123f057600080fd5b505af1158015612404573d6000803e3d6000fd5b50505050505050565b6020546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb600482015261432160248201819052916000916001600160a01b03909116906391d1485490604401602060405180830381865afa15801561249c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124c0919061a1ec565b90506124cb816159d9565b6020546028546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b03918216602482015260009291909116906391d1485490604401602060405180830381865afa15801561255b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061257f919061a1ec565b905061258a81615a53565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156125fc57600080fd5b505af1158015612610573d6000803e3d6000fd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156126a257600080fd5b505af11580156126b6573d6000803e3d6000fd5b50506040516001600160a01b03861681527f7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a059250602001905060405180910390a16020546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301529091169063950837aa90602401600060405180830381600087803b15801561275757600080fd5b505af115801561276b573d6000803e3d6000fd5b505060208054604080517f5b11259100000000000000000000000000000000000000000000000000000000815290516127fc95508894506001600160a01b0390921692635b112591926004808401938290030181865afa1580156127d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127f7919061a57e565b615aa5565b6020546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b038581166024830152909116906391d1485490604401602060405180830381865afa158015612885573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128a9919061a1ec565b91506128b482615a53565b6020546028546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b0391821660248201529116906391d1485490604401602060405180830381865afa15801561293f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612963919061a1ec565b9050611336816159d9565b6060601b805480602002602001604051908101604052809291908181526020016000905b8282101561198357838290600052602060002090600202016040518060400160405290816000820180546129c59061a20e565b80601f01602080910402602001604051908101604052809291908181526020018280546129f19061a20e565b8015612a3e5780601f10612a1357610100808354040283529160200191612a3e565b820191906000526020600020905b815481529060010190602001808311612a2157829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015612ad857602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411612a855790505b50505050508152505081526020019060010190612992565b6040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed7016900000000000000000000000000000000000000000000000000000000179052602854925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa792612b97926001600160a01b031691016001600160a01b0391909116815260200190565b600060405180830381600087803b158015612bb157600080fd5b505af1158015612bc5573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015612c4e57600080fd5b505af1158015612c62573d6000803e3d6000fd5b5050602080546040805192830181526101238352517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b0390911693506338e2252792506106f39190600090869060040161a3fa565b604080516001808252818301909252600091816020015b6060815260200190600190039081612cd75790505090506040518060400160405280600f81526020017f48656c6c6f2c20466f756e64727921000000000000000000000000000000000081525081600081518110612d3757612d3761a361565b6020908102919091010152604080516001808252818301909252600091816020016020820280368337019050509050602a81600081518110612d7b57612d7b61a361565b6020908102919091010152604051600190600090612da19085908590859060240161a3c2565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff05b6abf00000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612e5457600080fd5b505af1158015612e68573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250612f53919060040161a1d9565b600060405180830381600087803b158015612f6d57600080fd5b505af1158015612f81573d6000803e3d6000fd5b50506020546021546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd9350611d709290911690859060040161a09a565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561304757600080fd5b505af115801561305b573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527f0c8dc016000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b600060405180830381600087803b1580156130e557600080fd5b505af11580156130f9573d6000803e3d6000fd5b50506020546023546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef915060240161090b565b6060601a805480602002602001604051908101604052809291908181526020016000905b8282101561198357838290600052602060002001805461318d9061a20e565b80601f01602080910402602001604051908101604052809291908181526020018280546131b99061a20e565b80156132065780601f106131db57610100808354040283529160200191613206565b820191906000526020600020905b8154815290600101906020018083116131e957829003601f168201915b50505050508152602001906001019061316e565b6060601d805480602002602001604051908101604052809291908181526020016000905b828210156119835760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156132fd57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116132aa5790505b5050505050815250508152602001906001019061323e565b6021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156133a357600080fd5b505af11580156133b7573d6000803e3d6000fd5b50506040517f3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee925060009150a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561347257600080fd5b505af1158015613486573d6000803e3d6000fd5b5050602154604080518082018252600181527f3100000000000000000000000000000000000000000000000000000000000000602082015290516001600160a01b0390921693507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f92506134fd916000919061a454565b60405180910390a260285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561355e57600080fd5b505af1158015613572573d6000803e3d6000fd5b505060208054604080518084018252610123815260215482518084018452600181527f31000000000000000000000000000000000000000000000000000000000000009581019590955291517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b0393841696506338e2252795506136079491939092169160040161a3fa565b6000604051808303816000875af1158015613626573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261364e919081019061a1a4565b50565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156136c357600080fd5b505af11580156136d7573d6000803e3d6000fd5b5050602854604080516001600160a01b039092166024830152600060448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611769919060040161a1d9565b6060601c805480602002602001604051908101604052809291908181526020016000905b828210156119835760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561388657602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116138335790505b505050505081525050815260200190600101906137c7565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561398257600080fd5b505af1158015613996573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b960448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250613a81919060040161a1d9565b600060405180830381600087803b158015613a9b57600080fd5b505af1158015613aaf573d6000803e3d6000fd5b50506020546024546027546040517faa0c0fc10000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063aa0c0fc19450613b0f93928316929091169087908790602d9060040161a5a7565b600060405180830381600087803b158015613b2957600080fd5b505af1158015613b3d573d6000803e3d6000fd5b505050505050565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015611983578382906000526020600020018054613b889061a20e565b80601f0160208091040260200160405190810160405280929190818152602001828054613bb49061a20e565b8015613c015780601f10613bd657610100808354040283529160200191613c01565b820191906000526020600020905b815481529060010190602001808311613be457829003601f168201915b505050505081526020019060010190613b69565b60085460009060ff1615613c2d575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015613cbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ce2919061a5fc565b1415905090565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015613d5b57600080fd5b505af1158015613d6f573d6000803e3d6000fd5b5050602854604080516001600160a01b039092166024830152600060448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506130cb919060040161a1d9565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015613ead57600080fd5b505af1158015613ec1573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015613f4a57600080fd5b505af1158015613f5e573d6000803e3d6000fd5b50506020546040517f10188aef000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b0390911692506310188aef915060240161090b565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152670de0b6b3a76400009060009060250160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561404c57600080fd5b505af1158015614060573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061414b919060040161a1d9565b600060405180830381600087803b15801561416557600080fd5b505af1158015614179573d6000803e3d6000fd5b50506020546021546040517fcb7ba8e50000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063cb7ba8e5935086926123d79216908690602d9060040161a54a565b60408051602081018252600080825291516141f091607b9160240161a09a565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f676cc05400000000000000000000000000000000000000000000000000000000179052602854905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156142a357600080fd5b505af11580156142b7573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fed699775000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024016115f1565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561438457600080fd5b505af1158015614398573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614483919060040161a1d9565b600060405180830381600087803b15801561449d57600080fd5b505af11580156144b1573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561450557600080fd5b505af1158015614519573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561457657600080fd5b505af115801561458a573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614675919060040161a1d9565b600060405180830381600087803b15801561468f57600080fd5b505af11580156146a3573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156146f757600080fd5b505af115801561470b573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561476857600080fd5b505af115801561477c573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156147d057600080fd5b505af11580156147e4573d6000803e3d6000fd5b5050604080516004808252602480830184526020830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed701690000000000000000000000000000000000000000000000000000000017905292517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd93c06650000000000000000000000000000000000000000000000000000000091810191909152909350737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09101600060405180830381600087803b1580156148c657600080fd5b505af11580156148da573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561493757600080fd5b505af115801561494b573d6000803e3d6000fd5b50506020546021546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd935061499f9290911690859060040161a09a565b6000604051808303816000875af11580156149be573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526149e6919081019061a1a4565b5060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614a4057600080fd5b505af1158015614a54573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015614aa857600080fd5b505af1158015614abc573d6000803e3d6000fd5b50506021546040517ff30c7ba3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f30c7ba39250611e74916001600160a01b031690600090869060040161a42c565b60606015805480602002602001604051908101604052809291908181526020018280548015611840576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611822575050505050905090565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614c6457600080fd5b505af1158015614c78573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b960448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614d63919060040161a1d9565b600060405180830381600087803b158015614d7d57600080fd5b505af1158015614d91573d6000803e3d6000fd5b50506020546024546027546040517f5131ab590000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550635131ab599450613b0f9392831692909116908790879060040161a615565b604080516001808252818301909252600091816020015b6060815260200190600190039081614e055790505090506040518060400160405280600f81526020017f48656c6c6f2c20466f756e64727921000000000000000000000000000000000081525081600081518110614e6557614e6561a361565b6020908102919091010152604080516001808252818301909252600091816020016020820280368337019050509050602a81600081518110614ea957614ea961a361565b6020908102919091010152604051600190600090614ecf9085908590859060240161a3c2565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff05b6abf0000000000000000000000000000000000000000000000000000000017905260215490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391614f8c916001600160a01b039190911690600090869060040161a42c565b600060405180830381600087803b158015614fa657600080fd5b505af1158015614fba573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561504c57600080fd5b505af1158015615060573d6000803e3d6000fd5b50506020546040517f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b14693506150a792506001600160a01b039091169087908790879061a64c565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561513d57600080fd5b505af1158015615151573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f915061519790600090859061a454565b60405180910390a260285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401612f53565b604080517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152815160058183030181526025909101909152602154670de0b6b3a764000091906001600160a01b031631615242816000615b2e565b6021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156152d057600080fd5b505af11580156152e4573d6000803e3d6000fd5b50506020546040517f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b935061532892506001600160a01b0390911690602d9061a694565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156153be57600080fd5b505af11580156153d2573d6000803e3d6000fd5b5050602154604051600093506001600160a01b0390911691507fde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e0359061542590670de0b6b3a7640000908790602d9061a6b6565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561548657600080fd5b505af115801561549a573d6000803e3d6000fd5b50506020546021546040517fcb7ba8e50000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063cb7ba8e5935087926154f19216908790602d9060040161a54a565b6000604051808303818588803b15801561550a57600080fd5b505af115801561551e573d6000803e3d6000fd5b50506021546001600160a01b03163192506109399150829050670de0b6b3a7640000615b2e565b60408051808201909152600f81527f48656c6c6f2c20466f756e6472792100000000000000000000000000000000006020820152602154602a90600190670de0b6b3a7640000906155a2906000906001600160a01b031631615b2e565b60008484846040516024016155b99392919061a6cf565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905260215490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161567d916001600160a01b039190911690670de0b6b3a764000090869060040161a42c565b600060405180830381600087803b15801561569757600080fd5b505af11580156156ab573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561573d57600080fd5b505af1158015615751573d6000803e3d6000fd5b50506020546040517f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa935061579a92506001600160a01b0390911690859089908990899061a6f9565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561583057600080fd5b505af1158015615844573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f915061589190670de0b6b3a764000090859061a454565b60405180910390a260285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156158f257600080fd5b505af1158015615906573d6000803e3d6000fd5b50506020546021546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd9350869261595a921690869060040161a09a565b60006040518083038185885af1158015615978573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f191682016040526159a1919081019061a1a4565b50602154611db79083906001600160a01b031631615b2e565b60006159c4619c3b565b6159cf848483615b86565b9150505b92915050565b6040517fa59828850000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a5982885906024015b60006040518083038186803b158015615a3f57600080fd5b505afa158015611db7573d6000803e3d6000fd5b6040517f0c9fd5810000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90630c9fd58190602401615a27565b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f6906044015b60006040518083038186803b158015615b1a57600080fd5b505afa158015613b3d573d6000803e3d6000fd5b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c5490604401615b02565b600080615b938584615c01565b9050615bf66040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f78790000008152508286604051602001615be192919061a09a565b60405160208183030381529060405285615c0d565b9150505b9392505050565b6000615bfa8383615c3b565b60c08101515160009015615c3157615c2a84848460c00151615c56565b9050615bfa565b615c2a8484615dfc565b6000615c478383615ee7565b615bfa83836020015184615c0d565b600080615c61615ef3565b90506000615c6f8683615fc6565b90506000615c86826060015183602001518561646c565b90506000615c968383898961667e565b90506000615ca3826174fb565b602081015181519192509060030b15615d1657898260400151604051602001615ccd92919061a73a565b60408051601f19818403018152908290527f08c379a0000000000000000000000000000000000000000000000000000000008252615d0d9160040161a1d9565b60405180910390fd5b6000615d596040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a2000000000000000000000008152508360016176ca565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d90615dac90849060040161a1d9565b602060405180830381865afa158015615dc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615ded919061a57e565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc92590615e5190879060040161a1d9565b600060405180830381865afa158015615e6e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615e96919081019061a1a4565b90506000615ec48285604051602001615eb092919061a7bb565b6040516020818303038152906040526178ca565b90506001600160a01b0381166159cf578484604051602001615ccd92919061a7ea565b61073a828260006178dd565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c90615f7a90849060040161a895565b600060405180830381865afa158015615f97573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615fbf919081019061a8dc565b9250505090565b615ff86040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d90506160436040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b61604c856179e0565b6020820152600061605c86617dc5565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa15801561609e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526160c6919081019061a8dc565b868385602001516040516020016160e0949392919061a925565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb119061613890859060040161a1d9565b600060405180830381865afa158015616155573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261617d919081019061a8dc565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906161c590849060040161aa29565b602060405180830381865afa1580156161e2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616206919061a1ec565b61621b5781604051602001615ccd919061aa7b565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061626090849060040161ab0d565b600060405180830381865afa15801561627d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526162a5919081019061a8dc565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906162ec90849060040161ab5f565b602060405180830381865afa158015616309573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061632d919061a1ec565b156163c2576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061637790849060040161ab5f565b600060405180830381865afa158015616394573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526163bc919081019061a8dc565b60408501525b846001600160a01b03166349c4fac88286600001516040516020016163e7919061abb1565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161641392919061ac1d565b600060405180830381865afa158015616430573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616458919081019061a8dc565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b60608152602001906001900390816164885790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250816000815181106164e8576164e861a361565b60200260200101819052506040518060400160405280600381526020017f2d726c00000000000000000000000000000000000000000000000000000000008152508160018151811061653c5761653c61a361565b602002602001018190525084604051602001616558919061ac42565b6040516020818303038152906040528160028151811061657a5761657a61a361565b602002602001018190525082604051602001616596919061acae565b604051602081830303815290604052816003815181106165b8576165b861a361565b602002602001018190525060006165ce826174fb565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184526000808252908601528251808401909352905182529281019290925291925061665f9060408051808201825260008082526020918201528151808301909252845182528085019082015290618048565b6166745785604051602001615ccd919061acef565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d90156166ce565b511590565b6168425782602001511561678a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401615d0d565b8260c0015115616842576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401615d0d565b6040805160ff8082526120008201909252600091816020015b606081526020019060019003908161685b57905050905060006040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806168b69061adaf565b935060ff16815181106168cb576168cb61a361565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e370000000000000000000000000000000000000081525060405160200161691c919061adce565b6040516020818303038152906040528282806169379061adaf565b935060ff168151811061694c5761694c61a361565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806169999061adaf565b935060ff16815181106169ae576169ae61a361565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806169fb9061adaf565b935060ff1681518110616a1057616a1061a361565b60200260200101819052508760200151828280616a2c9061adaf565b935060ff1681518110616a4157616a4161a361565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e747261637450617468000000000000000000000000000000000000815250828280616a8e9061adaf565b935060ff1681518110616aa357616aa361a361565b602090810291909101015287518282616abb8161adaf565b935060ff1681518110616ad057616ad061a361565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e49640000000000000000000000000000000000000000000000815250828280616b1d9061adaf565b935060ff1681518110616b3257616b3261a361565b6020026020010181905250616b46466180a9565b8282616b518161adaf565b935060ff1681518110616b6657616b6661a361565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c650000000000000000000000000000000000815250828280616bb39061adaf565b935060ff1681518110616bc857616bc861a361565b602002602001018190525086828280616be09061adaf565b935060ff1681518110616bf557616bf561a361565b6020908102919091010152855115616d1c5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f6465000000000000000000000060208201528282616c468161adaf565b935060ff1681518110616c5b57616c5b61a361565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d90616cab90899060040161a1d9565b600060405180830381865afa158015616cc8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616cf0919081019061a8dc565b8282616cfb8161adaf565b935060ff1681518110616d1057616d1061a361565b60200260200101819052505b846020015115616dec5760408051808201909152601281527f2d2d766572696679536f75726365436f6465000000000000000000000000000060208201528282616d658161adaf565b935060ff1681518110616d7a57616d7a61a361565b60200260200101819052506040518060400160405280600581526020017f66616c7365000000000000000000000000000000000000000000000000000000815250828280616dc79061adaf565b935060ff1681518110616ddc57616ddc61a361565b6020026020010181905250616fb3565b616e246166c98660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b616eb75760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282616e678161adaf565b935060ff1681518110616e7c57616e7c61a361565b60200260200101819052508460a00151604051602001616e9c919061ac42565b604051602081830303815290604052828280616dc79061adaf565b8460c00151158015616efa575060408089015181518083018352600080825260209182015282518084019093528151835290810190820152616ef890511590565b155b15616fb35760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282616f3e8161adaf565b935060ff1681518110616f5357616f5361a361565b6020026020010181905250616f6788618149565b604051602001616f77919061ac42565b604051602081830303815290604052828280616f929061adaf565b935060ff1681518110616fa757616fa761a361565b60200260200101819052505b60408086015181518083018352600080825260209182015282518084019093528151835290810190820152616fe790511590565b61707c5760408051808201909152600b81527f2d2d72656c6179657249640000000000000000000000000000000000000000006020820152828261702a8161adaf565b935060ff168151811061703f5761703f61a361565b6020026020010181905250846040015182828061705b9061adaf565b935060ff16815181106170705761707061a361565b60200260200101819052505b60608501511561719d5760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826170c58161adaf565b935060ff16815181106170da576170da61a361565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015617149573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617171919081019061a8dc565b828261717c8161adaf565b935060ff16815181106171915761719161a361565b60200260200101819052505b60e085015151156172445760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826171e78161adaf565b935060ff16815181106171fc576171fc61a361565b60200260200101819052506172188560e00151600001516180a9565b82826172238161adaf565b935060ff16815181106172385761723861a361565b60200260200101819052505b60e085015160200151156172ee5760408051808201909152600a81527f2d2d676173507269636500000000000000000000000000000000000000000000602082015282826172918161adaf565b935060ff16815181106172a6576172a661a361565b60200260200101819052506172c28560e00151602001516180a9565b82826172cd8161adaf565b935060ff16815181106172e2576172e261a361565b60200260200101819052505b60e085015160400151156173985760408051808201909152600e81527f2d2d6d61784665655065724761730000000000000000000000000000000000006020820152828261733b8161adaf565b935060ff16815181106173505761735061a361565b602002602001018190525061736c8560e00151604001516180a9565b82826173778161adaf565b935060ff168151811061738c5761738c61a361565b60200260200101819052505b60e085015160600151156174425760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826173e58161adaf565b935060ff16815181106173fa576173fa61a361565b60200260200101819052506174168560e00151606001516180a9565b82826174218161adaf565b935060ff16815181106174365761743661a361565b60200260200101819052505b60008160ff1667ffffffffffffffff8111156174605761746061a0bc565b60405190808252806020026020018201604052801561749357816020015b606081526020019060019003908161747e5790505b50905060005b8260ff168160ff1610156174ec57838160ff16815181106174bc576174bc61a361565b6020026020010151828260ff16815181106174d9576174d961a361565b6020908102919091010152600101617499565b5093505050505b949350505050565b6175226040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c916175a89186910161ae39565b600060405180830381865afa1580156175c5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526175ed919081019061a8dc565b905060006175fb8683618c38565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b815260040161762b9190619ff0565b6000604051808303816000875af115801561764a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617672919081019061ae80565b805190915060030b1580159061768b5750602081015151155b801561769a5750604081015151155b1561667457816000815181106176b2576176b261a361565b6020026020010151604051602001615ccd919061af36565b606060006176ff8560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925286518252808701908201529091506177369082905b90618d8d565b156178935760006177b3826177ad846177a76177798a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b90618db4565b90618e16565b604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150617817908290618d8d565b1561788157604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261787e905b8290618e9b565b90505b61788a81618ec1565b92505050615bfa565b82156178ac578484604051602001615ccd92919061b122565b5050604080516020810190915260008152615bfa565b509392505050565b6000808251602084016000f09392505050565b8160a00151156178ec57505050565b60006178f9848484618f2a565b90506000617906826174fb565b602081015181519192509060030b1580156179a25750604080518082018252600781527f5355434345535300000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526179a290604080518082018252600080825260209182015281518083019092528451825280850190820152617730565b156179af57505050505050565b604082015151156179cf578160400151604051602001615ccd919061b1c9565b80604051602001615ccd919061b227565b60606000617a158360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150617a7a905b8290618048565b15617ae957604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615bfa90617ae49083906194c5565b618ec1565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617b4b905b829061954f565b600103617c1857604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617bb190617877565b50604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615bfa90617ae4905b8390618e9b565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617c7790617a73565b15617dae57604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290617cdf9083906195e9565b905060008160018351617cf2919061b292565b81518110617d0257617d0261a361565b60200260200101519050617da5617ae4617d786040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528551825280860190820152906194c5565b95945050505050565b82604051602001615ccd919061b2a5565b50919050565b60606000617dfa8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150617e5c90617a73565b15617e6a57615bfa81618ec1565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617ec990617b44565b600103617f3357604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615bfa90617ae490617c11565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617f9290617a73565b15617dae57604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290617ffa9083906195e9565b9050600181511115618036578060028251618015919061b292565b815181106180255761802561a361565b602002602001015192505050919050565b5082604051602001615ccd919061b2a5565b80518251600091111561805d575060006159d3565b815183516020850151600092916180739161b383565b61807d919061b292565b9050826020015181036180945760019150506159d3565b82516020840151819020912014905092915050565b606060006180b68361968e565b600101905060008167ffffffffffffffff8111156180d6576180d661a0bc565b6040519080825280601f01601f191660200182016040528015618100576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461810a57509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e53454400000000000000000000000000000000000000000000818401908152855180870187528381528401929092528451808601909552518452908301526060916181d5905b8290619770565b1561821557505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e7365000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618274906181ce565b156182b457505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d4954000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618313906181ce565b1561835357505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526183b2906181ce565b806184175750604080518082018252601081527f47504c2d322e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618417906181ce565b1561845757505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526184b6906181ce565b8061851b5750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261851b906181ce565b1561855b57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526185ba906181ce565b8061861f5750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261861f906181ce565b1561865f57505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526186be906181ce565b806187235750604080518082018252601181527f4c47504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618723906181ce565b1561876357505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c617573650000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526187c2906181ce565b1561880257505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618861906181ce565b156188a157505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618900906181ce565b1561894057505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261899f906181ce565b156189df57505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e300000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618a3e906181ce565b15618a7e57505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618add906181ce565b80618b425750604080518082018252601181527f4147504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618b42906181ce565b15618b8257505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e3100000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618be1906181ce565b15618c2157505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151615ccd929060200161b396565b60608060005b8451811015618cc35781858281518110618c5a57618c5a61a361565b6020026020010151604051602001618c7392919061a7bb565b604051602081830303815290604052915060018551618c92919061b292565b8114618cbb5781604051602001618ca9919061b4ff565b60405160208183030381529060405291505b600101618c3e565b5060408051600380825260808201909252600091816020015b6060815260200190600190039081618cdc5790505090508381600081518110618d0757618d0761a361565b60200260200101819052506040518060400160405280600281526020017f2d6300000000000000000000000000000000000000000000000000000000000081525081600181518110618d5b57618d5b61a361565b60200260200101819052508181600281518110618d7a57618d7a61a361565b6020908102919091010152949350505050565b6020808301518351835192840151600093618dab9291849190619784565b14159392505050565b60408051808201909152600080825260208201526000618de68460000151856020015185600001518660200151619895565b9050836020015181618df8919061b292565b84518590618e0790839061b292565b90525060208401525090919050565b6040805180820190915260008082526020820152815183511015618e3b5750816159d3565b6020808301519084015160019114618e625750815160208481015190840151829020919020145b8015618e9357825184518590618e7990839061b292565b9052508251602085018051618e8f90839061b383565b9052505b509192915050565b6040805180820190915260008082526020820152618eba8383836199b5565b5092915050565b60606000826000015167ffffffffffffffff811115618ee257618ee261a0bc565b6040519080825280601f01601f191660200182016040528015618f0c576020820181803683370190505b5090506000602082019050618eba8185602001518660000151619a60565b60606000618f36615ef3565b6040805160ff808252612000820190925291925060009190816020015b6060815260200190600190039081618f5357905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280618fae9061adaf565b935060ff1681518110618fc357618fc361a361565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e3300000000000000000000000000000000000000000000000000815250604051602001619014919061b540565b60405160208183030381529060405282828061902f9061adaf565b935060ff16815181106190445761904461a361565b60200260200101819052506040518060400160405280600881526020017f76616c69646174650000000000000000000000000000000000000000000000008152508282806190919061adaf565b935060ff16815181106190a6576190a661a361565b6020026020010181905250826040516020016190c2919061acae565b6040516020818303038152906040528282806190dd9061adaf565b935060ff16815181106190f2576190f261a361565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061913f9061adaf565b935060ff16815181106191545761915461a361565b60200260200101819052506191698784619ada565b82826191748161adaf565b935060ff16815181106191895761918961a361565b6020908102919091010152855151156192355760408051808201909152600b81527f2d2d7265666572656e6365000000000000000000000000000000000000000000602082015282826191db8161adaf565b935060ff16815181106191f0576191f061a361565b6020026020010181905250619209866000015184619ada565b82826192148161adaf565b935060ff16815181106192295761922961a361565b60200260200101819052505b8560800151156192a35760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261927e8161adaf565b935060ff16815181106192935761929361a361565b6020026020010181905250619309565b84156193095760408051808201909152601281527f2d2d726571756972655265666572656e63650000000000000000000000000000602082015282826192e88161adaf565b935060ff16815181106192fd576192fd61a361565b60200260200101819052505b604086015151156193a55760408051808201909152600d81527f2d2d756e73616665416c6c6f7700000000000000000000000000000000000000602082015282826193538161adaf565b935060ff16815181106193685761936861a361565b602002602001018190525085604001518282806193849061adaf565b935060ff16815181106193995761939961a361565b60200260200101819052505b85606001511561940f5760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826193ee8161adaf565b935060ff16815181106194035761940361a361565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561942d5761942d61a0bc565b60405190808252806020026020018201604052801561946057816020015b606081526020019060019003908161944b5790505b50905060005b8260ff168160ff1610156194b957838160ff16815181106194895761948961a361565b6020026020010151828260ff16815181106194a6576194a661a361565b6020908102919091010152600101619466565b50979650505050505050565b60408051808201909152600080825260208201528151835110156194ea5750816159d3565b815183516020850151600092916195009161b383565b61950a919061b292565b6020840151909150600190821461952b575082516020840151819020908220145b80156195465783518551869061954290839061b292565b9052505b50929392505050565b60008082600001516195738560000151866020015186600001518760200151619895565b61957d919061b383565b90505b83516020850151619591919061b383565b8111618eba57816195a18161b585565b92505082600001516195d88560200151836195bc919061b292565b86516195c8919061b292565b8386600001518760200151619895565b6195e2919061b383565b9050619580565b606060006195f7848461954f565b61960290600161b383565b67ffffffffffffffff81111561961a5761961a61a0bc565b60405190808252806020026020018201604052801561964d57816020015b60608152602001906001900390816196385790505b50905060005b81518110156178c257619669617ae48686618e9b565b82828151811061967b5761967b61a361565b6020908102919091010152600101619653565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106196d7577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310619703576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061972157662386f26fc10000830492506010015b6305f5e1008310619739576305f5e100830492506008015b612710831061974d57612710830492506004015b6064831061975f576064830492506002015b600a83106159d35760010192915050565b600061977c8383619b1a565b159392505050565b60008085841161988b576020841161983757600084156197cf5760016197ab86602061b292565b6197b690600861b59f565b6197c190600261b69d565b6197cb919061b292565b1990505b83518116856197de898961b383565b6197e8919061b292565b805190935082165b8181146198225787841161980a57879450505050506174f3565b836198148161b6a9565b9450508284511690506197f0565b61982c878561b383565b9450505050506174f3565b838320619844858861b292565b61984e908761b383565b91505b858210619889578482208082036198765761986c868461b383565b93505050506174f3565b61988160018461b292565b925050619851565b505b5092949350505050565b600083818685116199a0576020851161994f57600085156198e15760016198bd87602061b292565b6198c890600861b59f565b6198d390600261b69d565b6198dd919061b292565b1990505b845181166000876198f28b8b61b383565b6198fc919061b292565b855190915083165b828114619941578186106199295761991c8b8b61b383565b96505050505050506174f3565b856199338161b585565b965050838651169050619904565b8596505050505050506174f3565b508383206000905b619961868961b292565b821161999e5785832080820361997d57839450505050506174f3565b61998860018561b383565b93505081806199969061b585565b925050619957565b505b6199aa878761b383565b979650505050505050565b604080518082019091526000808252602082015260006199e78560000151866020015186600001518760200151619895565b602080870180519186019190915251909150619a03908261b292565b835284516020860151619a16919061b383565b8103619a255760008552619a57565b83518351619a33919061b383565b85518690619a4290839061b292565b9052508351619a51908261b383565b60208601525b50909392505050565b60208110619a985781518352619a7760208461b383565b9250619a8460208361b383565b9150619a9160208261b292565b9050619a60565b6000198115619ac7576001619aae83602061b292565b619aba9061010061b69d565b619ac4919061b292565b90505b9151835183169219169190911790915250565b60606000619ae88484615fc6565b8051602080830151604051939450619b029390910161b6c0565b60405160208183030381529060405291505092915050565b8151815160009190811115619b2d575081515b6020808501519084015160005b83811015619be65782518251808214619bb6576000196020871015619b9557600184619b6789602061b292565b619b71919061b383565b619b7c90600861b59f565b619b8790600261b69d565b619b91919061b292565b1990505b8181168382168181039114619bb35797506159d39650505050505050565b50505b619bc160208661b383565b9450619bce60208561b383565b93505050602081619bdf919061b383565b9050619b3a565b5084518651616674919061b718565b610c9f806200b73983390190565b6112a6806200c3d883390190565b611e03806200d67e83390190565b6119ba806200f48183390190565b610efa8062010e3b83390190565b6040518060e00160405280606081526020016060815260200160608152602001600015158152602001600015158152602001600015158152602001619c7e619c83565b905290565b60405180610100016040528060001515815260200160001515815260200160608152602001600080191681526020016060815260200160608152602001600015158152602001619c7e6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b81811015619d355783516001600160a01b0316835260209384019390920191600101619d0e565b509095945050505050565b60005b83811015619d5b578181015183820152602001619d43565b50506000910152565b60008151808452619d7c816020860160208601619d40565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015619e8c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b81811015619e72577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8503018352619e5c848651619d64565b6020958601959094509290920191600101619e22565b509197505050602094850194929092019150600101619db8565b50929695505050505050565b600081518084526020840193506020830160005b82811015619eec5781517fffffffff0000000000000000000000000000000000000000000000000000000016865260209586019590910190600101619eac565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015619e8c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160408752619f626040880182619d64565b9050602082015191508681036020880152619f7d8183619e98565b965050506020938401939190910190600101619f1e565b600082825180855260208501945060208160051b8301016020850160005b83811015619fe457601f19858403018852619fce838351619d64565b6020988901989093509190910190600101619fb2565b50909695505050505050565b602081526000615bfa6020830184619f94565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015619e8c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261a0846040870182619e98565b955050602093840193919091019060010161a02b565b6001600160a01b03831681526040602082015260006174f36040830184619d64565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561a10e5761a10e61a0bc565b60405290565b60008067ffffffffffffffff84111561a12f5761a12f61a0bc565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561a15e5761a15e61a0bc565b60405283815290508082840185101561a17657600080fd5b6178c2846020830185619d40565b600082601f83011261a19557600080fd5b615bfa8383516020850161a114565b60006020828403121561a1b657600080fd5b815167ffffffffffffffff81111561a1cd57600080fd5b6159cf8482850161a184565b602081526000615bfa6020830184619d64565b60006020828403121561a1fe57600080fd5b81518015158114615bfa57600080fd5b600181811c9082168061a22257607f821691505b602082108103617dbf577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f82111561133657806000526020600020601f840160051c8101602085101561a2825750805b601f840160051c820191505b81811015611db7576000815560010161a28e565b815167ffffffffffffffff81111561a2bc5761a2bc61a0bc565b61a2d08161a2ca845461a20e565b8461a25b565b6020601f82116001811461a304576000831561a2ec5750848201515b600019600385901b1c1916600184901b178455611db7565b600084815260208120601f198516915b8281101561a334578785015182556020948501946001909201910161a314565b508482101561a3525786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081518084526020840193506020830160005b82811015619eec57815186526020958601959091019060010161a3a4565b60608152600061a3d56060830186619f94565b828103602084015261a3e7818661a390565b9150508215156040830152949350505050565b6001600160a01b0384511681526001600160a01b0383166020820152606060408201526000617da56060830184619d64565b6001600160a01b0384168152826020820152606060408201526000617da56060830184619d64565b8281526040602082015260006174f36040830184619d64565b6001600160a01b0381541682526001600160a01b03600182015416602083015260028101546040830152600060038201608060608501526000815461a4b18161a20e565b806080880152600182166000811461a4d0576001811461a50a5761a53e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660a089015260a082151560051b890101935061a53e565b84600052602060002060005b8381101561a5355781548a820160a0015260019091019060200161a516565b890160a0019450505b50919695505050505050565b6001600160a01b038416815260606020820152600061a56c6060830185619d64565b8281036040840152616674818561a46d565b60006020828403121561a59057600080fd5b81516001600160a01b0381168114615bfa57600080fd5b6001600160a01b03861681526001600160a01b038516602082015283604082015260a06060820152600061a5de60a0830185619d64565b828103608084015261a5f0818561a46d565b98975050505050505050565b60006020828403121561a60e57600080fd5b5051919050565b6001600160a01b03851681526001600160a01b03841660208201528260408201526080606082015260006166746080830184619d64565b6001600160a01b038516815260806020820152600061a66e6080830186619f94565b828103604084015261a680818661a390565b915050821515606083015295945050505050565b6001600160a01b03831681526040602082015260006174f3604083018461a46d565b83815260606020820152600061a56c6060830185619d64565b60608152600061a6e26060830186619d64565b602083019490945250901515604090910152919050565b6001600160a01b038616815284602082015260a06040820152600061a72160a0830186619d64565b6060830194909452509015156080909101529392505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161a77281601a850160208801619d40565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161a7af81601c840160208801619d40565b01601c01949350505050565b6000835161a7cd818460208801619d40565b83519083019061a7e1818360208801619d40565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161a82281601a850160208801619d40565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161a85f816033840160208801619d40565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000615bfa6080830184619d64565b60006020828403121561a8ee57600080fd5b815167ffffffffffffffff81111561a90557600080fd5b8201601f8101841361a91657600080fd5b6159cf8482516020840161a114565b6000855161a937818460208a01619d40565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161a971816001840160208a01619d40565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161a9af816002840160208901619d40565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161a9f1816002840160208801619d40565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061aa3c6040830184619d64565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161aab381601f850160208701619d40565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061ab206040830184619d64565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061ab726040830184619d64565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161abe9816014850160208701619d40565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061ac306040830185619d64565b8281036020840152615bf68185619d64565b7f220000000000000000000000000000000000000000000000000000000000000081526000825161ac7a816001850160208701619d40565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161acc0818460208701619d40565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161ad7381604b850160208701619d40565b91909101604b0192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600060ff821660ff810361adc55761adc561ad80565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161ae2c816029850160208701619d40565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000615bfa6080830184619d64565b60006020828403121561ae9257600080fd5b815167ffffffffffffffff81111561aea957600080fd5b82016060818503121561aebb57600080fd5b61aec361a0eb565b81518060030b811461aed457600080fd5b8152602082015167ffffffffffffffff81111561aef057600080fd5b61aefc8682850161a184565b602083015250604082015167ffffffffffffffff81111561af1c57600080fd5b61af288682850161a184565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161af94816021850160208701619d40565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161b180816021850160208801619d40565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161b1bd81602e840160208801619d40565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161ae2c816029850160208701619d40565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161b285816022850160208701619d40565b9190910160220192915050565b818103818111156159d3576159d361ad80565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161b2dd81600e850160208701619d40565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b808201808211156159d3576159d361ad80565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161b3ce816018850160208801619d40565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161b40b81601c840160208801619d40565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161b511818460208701619d40565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161b57881601c850160208701619d40565b91909101601c0192915050565b6000600019820361b5985761b59861ad80565b5060010190565b80820281158282048414176159d3576159d361ad80565b6001815b600184111561b5f15780850481111561b5d55761b5d561ad80565b600184161561b5e357908102905b60019390931c92800261b5ba565b935093915050565b60008261b608575060016159d3565b8161b615575060006159d3565b816001811461b62b576002811461b6355761b651565b60019150506159d3565b60ff84111561b6465761b64661ad80565b50506001821b6159d3565b5060208310610133831016604e8410600b841016171561b674575081810a6159d3565b61b681600019848461b5b6565b806000190482111561b6955761b69561ad80565b029392505050565b6000615bfa838361b5f9565b60008161b6b85761b6b861ad80565b506000190190565b6000835161b6d2818460208801619d40565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161b70c816001840160208801619d40565b01600101949350505050565b8181036000831280158383131683831282161715618eba57618eba61ad8056fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220a043c41353215fce25ecb67a8a4f6f724aaa86dea8dcb0a6975bfb1f10ff3beb64736f6c634300081a0033608060405234801561001057600080fd5b506040516112a63803806112a683398101604081905261002f91610110565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007891906101e2565b50600461008582826101e2565b5050506001600160a01b03821615806100a557506001600160a01b038116155b156100c35760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b031991821617909155600780549290931691161790556102a0565b80516001600160a01b038116811461010b57600080fd5b919050565b6000806040838503121561012357600080fd5b61012c836100f4565b915061013a602084016100f4565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061016d57607f821691505b60208210810361018d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101dd57806000526020600020601f840160051c810160208510156101ba5750805b601f840160051c820191505b818110156101da57600081556001016101c6565b50505b505050565b81516001600160401b038111156101fb576101fb610143565b61020f816102098454610159565b84610193565b6020601f821160018114610243576000831561022b5750848201515b600019600385901b1c1916600184901b1784556101da565b600084815260208120601f198516915b828110156102735787850151825560209485019460019092019101610253565b50848210156102915786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610ff7806102af6000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806342966c68116100b257806379cc679011610081578063a9059cbb11610066578063a9059cbb1461028e578063bff9662a146102a1578063dd62ed3e146102c157600080fd5b806379cc67901461027357806395d89b411461028657600080fd5b806342966c68146102025780635b1125911461021557806370a0823114610235578063779e3b631461026b57600080fd5b80631e458bee116100ee5780631e458bee1461018857806323b872dd1461019b578063313ce567146101ae578063328a01d0146101bd57600080fd5b806306fdde0314610120578063095ea7b31461013e57806315d57fd41461016157806318160ddd14610176575b600080fd5b610128610307565b6040516101359190610d97565b60405180910390f35b61015161014c366004610e2c565b610399565b6040519015158152602001610135565b61017461016f366004610e56565b6103b3565b005b6002545b604051908152602001610135565b610174610196366004610e89565b61057e565b6101516101a9366004610ebc565b610631565b60405160128152602001610135565b6007546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610135565b610174610210366004610ef9565b610655565b6006546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a610243366004610f12565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b610174610662565b610174610281366004610e2c565b610786565b610128610837565b61015161029c366004610e2c565b610846565b6005546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a6102cf366004610e56565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60606003805461031690610f34565b80601f016020809104026020016040519081016040528092919081815260200182805461034290610f34565b801561038f5780601f106103645761010080835404028352916020019161038f565b820191906000526020600020905b81548152906001019060200180831161037257829003601f168201915b5050505050905090565b6000336103a7818585610854565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103f3575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610431576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82161580610468575073ffffffffffffffffffffffffffffffffffffffff8116155b1561049f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105d1576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6105db8383610866565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161062491815260200190565b60405180910390a3505050565b60003361063f8582856108c6565b61064a858585610995565b506001949350505050565b61065f3382610a40565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106b5576040517fe700765e000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b60065473ffffffffffffffffffffffffffffffffffffffff16610704576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107d9576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6107e38282610a9c565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161082b91815260200190565b60405180910390a25050565b60606004805461031690610f34565b6000336103a7818585610995565b6108618383836001610ab1565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108b6576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c260008383610bf9565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461098f5781811015610980576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610428565b61098f84848484036000610ab1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109e5576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8216610a35576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b610861838383610bf9565b73ffffffffffffffffffffffffffffffffffffffff8216610a90576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c282600083610bf9565b610aa78233836108c6565b6108c28282610a40565b73ffffffffffffffffffffffffffffffffffffffff8416610b01576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8316610b51576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600160209081526040808320938716835292905220829055801561098f578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610beb91815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c31578060026000828254610c269190610f87565b90915550610ce39050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610cb7576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610428565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610d0c57600280548290039055610d38565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161062491815260200190565b602081526000825180602084015260005b81811015610dc55760208186018101516040868401015201610da8565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610e2757600080fd5b919050565b60008060408385031215610e3f57600080fd5b610e4883610e03565b946020939093013593505050565b60008060408385031215610e6957600080fd5b610e7283610e03565b9150610e8060208401610e03565b90509250929050565b600080600060608486031215610e9e57600080fd5b610ea784610e03565b95602085013595506040909401359392505050565b600080600060608486031215610ed157600080fd5b610eda84610e03565b9250610ee860208501610e03565b929592945050506040919091013590565b600060208284031215610f0b57600080fd5b5035919050565b600060208284031215610f2457600080fd5b610f2d82610e03565b9392505050565b600181811c90821680610f4857607f821691505b602082108103610f81577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b808201808211156103ad577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea264697066735822122001ec0ce060384773f3d3389fab7bed652c6b8ee389a7471cce10d00d87a75a0c64736f6c634300081a003360a060405234801561001057600080fd5b50604051611e03380380611e0383398101604081905261002f916101fd565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03838116608052600480546001600160a01b0319169184169190911790556100bb60008261014c565b506100e67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261014c565b506101117f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361014c565b5061012a600080516020611de38339815191528261014c565b50610143600080516020611de38339815191528361014c565b50505050610240565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166101d75760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016101db565b5060005b92915050565b80516001600160a01b03811681146101f857600080fd5b919050565b60008060006060848603121561021257600080fd5b61021b846101e1565b9250610229602085016101e1565b9150610237604085016101e1565b90509250925092565b608051611b6c610277600039600081816101d501528181610574015281816105c90152818161099601526109eb0152611b6c6000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c806385f438c1116100ee578063a217fddf11610097578063d9caed1211610071578063d9caed12146103e0578063e609055e146103f3578063e63ab1e914610406578063eab103df1461042d57600080fd5b8063a217fddf146103a2578063d547741f146103aa578063d936547e146103bd57600080fd5b806399a3c356116100c857806399a3c356146103695780639a5904271461037c5780639b19251a1461038f57600080fd5b806385f438c1146102f657806391d148541461031d578063950837aa1461035657600080fd5b806336568abe116101505780635b1125911161012a5780635b112591146102d05780635c975abb146102e35780638456cb59146102ee57600080fd5b806336568abe1461028e5780633f4ba83a146102a1578063570618e1146102a957600080fd5b8063248a9ca311610181578063248a9ca314610224578063252f07bf146102565780632f2ff15d1461027b57600080fd5b806301ffc9a7146101a8578063116191b6146101d057806321fc65f21461020f575b600080fd5b6101bb6101b636600461155e565b610440565b60405190151581526020015b60405180910390f35b6101f77f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101c7565b61022261021d3660046115fe565b6104d9565b005b610248610232366004611671565b6000908152600160208190526040909120015490565b6040519081526020016101c7565b6004546101bb9074010000000000000000000000000000000000000000900460ff1681565b61022261028936600461168a565b610699565b61022261029c36600461168a565b6106c5565b610222610716565b6102487f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b6004546101f7906001600160a01b031681565b60025460ff166101bb565b61022261074b565b6102487f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101bb61032b36600461168a565b60009182526001602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6102226103643660046116ba565b61077d565b6102226103773660046116d7565b6108fb565b61022261038a3660046116ba565b610ac0565b61022261039d3660046116ba565b610b74565b610248600081565b6102226103b836600461168a565b610c2b565b6101bb6103cb3660046116ba565b60036020526000908152604090205460ff1681565b6102226103ee36600461177a565b610c51565b6102226104013660046117bb565b610d49565b6102487f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b61022261043b36600461185a565b610f75565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104d357507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104e1610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461050b8161100e565b610513611018565b6001600160a01b03851660009081526003602052604090205460ff16610565576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105996001600160a01b0386167f000000000000000000000000000000000000000000000000000000000000000086611057565b6040517f5131ab590000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906106069088908a908990899089906004016118c0565b600060405180830381600087803b15801561062057600080fd5b505af1158015610634573d6000803e3d6000fd5b50505050846001600160a01b0316866001600160a01b03167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d586868660405161067f93929190611903565b60405180910390a3506106926001600055565b5050505050565b600082815260016020819052604090912001546106b58161100e565b6106bf83836110cb565b50505050565b6001600160a01b0381163314610707576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610711828261115e565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107408161100e565b6107486111e5565b50565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107758161100e565b610748611237565b60006107888161100e565b6001600160a01b0382166107c8576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004546107ff907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4906001600160a01b031661115e565b50600454610837907f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a906001600160a01b031661115e565b506108627f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4836110cb565b5061088d7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a836110cb565b50600480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040519081527f086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba9060200160405180910390a15050565b610903610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461092d8161100e565b610935611018565b6001600160a01b03861660009081526003602052604090205460ff16610987576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109bb6001600160a01b0387167f000000000000000000000000000000000000000000000000000000000000000087611057565b6040517faa0c0fc10000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063aa0c0fc190610a2a9089908b908a908a908a908a906004016119db565b600060405180830381600087803b158015610a4457600080fd5b505af1158015610a58573d6000803e3d6000fd5b50505050856001600160a01b0316876001600160a01b03167f7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb972187878787604051610aa59493929190611a32565b60405180910390a350610ab86001600055565b505050505050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a610aea8161100e565b6001600160a01b038216610b2a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a610b9e8161100e565b6001600160a01b038216610bde576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b60008281526001602081905260409091200154610c478161100e565b6106bf838361115e565b610c59610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610c838161100e565b610c8b611018565b6001600160a01b03831660009081526003602052604090205460ff16610cdd576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cf16001600160a01b0384168584611057565b826001600160a01b0316846001600160a01b03167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610d3691815260200190565b60405180910390a3506107116001600055565b610d51610fcb565b610d59611018565b60045474010000000000000000000000000000000000000000900460ff16610dad576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03841660009081526003602052604090205460ff16610dff576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610e5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e839190611a5e565b9050610e9a6001600160a01b038616333087611274565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa158015610f21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f459190611a5e565b610f4f9190611a77565b8787604051610f62959493929190611ab1565b60405180910390a250610ab86001600055565b6000610f808161100e565b506004805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b600260005403611007576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61074881336112ad565b60025460ff1615611055576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6040516001600160a01b0383811660248301526044820183905261071191859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611324565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166111565760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016104d3565b5060006104d3565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff16156111565760008381526001602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104d3565b6111ed6113a0565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b61123f611018565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861121a3390565b6040516001600160a01b0384811660248301528381166044830152606482018390526106bf9186918216906323b872dd90608401611084565b60008281526001602090815260408083206001600160a01b038516845290915290205460ff16611320576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044015b60405180910390fd5b5050565b60006113396001600160a01b038416836113dc565b9050805160001415801561135e57508080602001905181019061135c9190611aea565b155b15610711576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611317565b60025460ff16611055576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606113ea838360006113f1565b9392505050565b60608147101561142f576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611317565b600080856001600160a01b0316848660405161144b9190611b07565b60006040518083038185875af1925050503d8060008114611488576040519150601f19603f3d011682016040523d82523d6000602084013e61148d565b606091505b509150915061149d8683836114a7565b9695505050505050565b6060826114bc576114b78261151c565b6113ea565b81511580156114d357506001600160a01b0384163b155b15611515576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611317565b50806113ea565b80511561152c5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561157057600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146113ea57600080fd5b6001600160a01b038116811461074857600080fd5b60008083601f8401126115c757600080fd5b50813567ffffffffffffffff8111156115df57600080fd5b6020830191508360208285010111156115f757600080fd5b9250929050565b60008060008060006080868803121561161657600080fd5b8535611621816115a0565b94506020860135611631816115a0565b935060408601359250606086013567ffffffffffffffff81111561165457600080fd5b611660888289016115b5565b969995985093965092949392505050565b60006020828403121561168357600080fd5b5035919050565b6000806040838503121561169d57600080fd5b8235915060208301356116af816115a0565b809150509250929050565b6000602082840312156116cc57600080fd5b81356113ea816115a0565b60008060008060008060a087890312156116f057600080fd5b86356116fb816115a0565b9550602087013561170b816115a0565b945060408701359350606087013567ffffffffffffffff81111561172e57600080fd5b61173a89828a016115b5565b909450925050608087013567ffffffffffffffff81111561175a57600080fd5b87016080818a03121561176c57600080fd5b809150509295509295509295565b60008060006060848603121561178f57600080fd5b833561179a816115a0565b925060208401356117aa816115a0565b929592945050506040919091013590565b600080600080600080608087890312156117d457600080fd5b863567ffffffffffffffff8111156117eb57600080fd5b6117f789828a016115b5565b909750955050602087013561180b816115a0565b935060408701359250606087013567ffffffffffffffff81111561182e57600080fd5b61183a89828a016115b5565b979a9699509497509295939492505050565b801515811461074857600080fd5b60006020828403121561186c57600080fd5b81356113ea8161184c565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b03861681526001600160a01b03851660208201528360408201526080606082015260006118f8608083018486611877565b979650505050505050565b83815260406020820152600061191d604083018486611877565b95945050505050565b60008135611933816115a0565b6001600160a01b03168352602082013561194c816115a0565b6001600160a01b03166020840152604082810135908401526060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261199a57600080fd5b820160208101903567ffffffffffffffff8111156119b757600080fd5b8036038213156119c657600080fd5b6080606086015261191d608086018284611877565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a060608201526000611a1360a083018587611877565b8281036080840152611a258185611926565b9998505050505050505050565b848152606060208201526000611a4c606083018587611877565b82810360408401526118f88185611926565b600060208284031215611a7057600080fd5b5051919050565b818103818111156104d3577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b606081526000611ac5606083018789611877565b8560208401528281036040840152611ade818587611877565b98975050505050505050565b600060208284031215611afc57600080fd5b81516113ea8161184c565b6000825160005b81811015611b285760208186018101518583015201611b0e565b50600092019182525091905056fea264697066735822122002398946b3da587e3a38b71ff523ef333036b84dae258034ccd3cb804048b29b64736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60c060405260001960045534801561001657600080fd5b506040516119ba3803806119ba83398101604081905261003591610238565b60016000819055805460ff19169055838383836001600160a01b038416158061006557506001600160a01b038316155b8061007757506001600160a01b038216155b8061008957506001600160a01b038116155b156100a75760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0384811660805283811660a052600380546001600160a01b0319169184169190911790556100dd60008261016c565b506101087f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361016c565b506101337f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb8361016c565b5061015e7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261016c565b50505050505050505061028c565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166102125760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101ca3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610216565b5060005b92915050565b80516001600160a01b038116811461023357600080fd5b919050565b6000806000806080858703121561024e57600080fd5b6102578561021c565b93506102656020860161021c565b92506102736040860161021c565b91506102816060860161021c565b905092959194509250565b60805160a0516116ca6102f060003960008181610220015281816106d80152818161086d015281816109e401528181610ce40152610e060152600081816101d401528181610648015281816106ab015281816107dd015261084001526116ca6000f3fe608060405234801561001057600080fd5b506004361061018d5760003560e01c80636f8728ad116100e3578063950837aa1161008c578063d547741f11610066578063d547741f146103cf578063d5abeb01146103e2578063e63ab1e9146103eb57600080fd5b8063950837aa1461038d578063a217fddf146103a0578063a783c789146103a857600080fd5b80638456cb59116100bd5780638456cb591461031857806385f438c11461032057806391d148541461034757600080fd5b80636f8728ad146102df5780636f8b44b0146102f2578063743e0c9b1461030557600080fd5b80632f2ff15d116101455780635b1125911161011f5780635b112591146102a15780635c975abb146102c15780635e3e9fef146102cc57600080fd5b80632f2ff15d1461027357806336568abe146102865780633f4ba83a1461029957600080fd5b8063116191b611610176578063116191b6146101cf57806321e093b11461021b578063248a9ca31461024257600080fd5b806301ffc9a714610192578063106e6290146101ba575b600080fd5b6101a56101a03660046111c8565b610412565b60405190151581526020015b60405180910390f35b6101cd6101c836600461123a565b6104ab565b005b6101f67f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b1565b6101f67f000000000000000000000000000000000000000000000000000000000000000081565b61026561025036600461126d565b60009081526002602052604090206001015490565b6040519081526020016101b1565b6101cd610281366004611286565b610550565b6101cd610294366004611286565b61057b565b6101cd6105d4565b6003546101f69073ffffffffffffffffffffffffffffffffffffffff1681565b60015460ff166101a5565b6101cd6102da3660046112fb565b610609565b6101cd6102ed36600461135d565b61079e565b6101cd61030036600461126d565b610938565b6101cd61031336600461126d565b6109a7565b6101cd610a51565b6102657f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101a5610355366004611286565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6101cd61039b3660046113f5565b610a83565b610265600081565b6102657f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101cd6103dd366004611286565b610c2e565b61026560045481565b6102657f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104a557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104b3610c53565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46104dd81610c96565b6104e5610ca0565b6104f0848484610cdf565b8373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d58460405161053891815260200190565b60405180910390a25061054b6001600055565b505050565b60008281526002602052604090206001015461056b81610c96565b6105758383610e67565b50505050565b73ffffffffffffffffffffffffffffffffffffffff811633146105ca576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61054b8282610f67565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6105fe81610c96565b610606611026565b50565b610611610c53565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461063b81610c96565b610643610ca0565b61066e7f00000000000000000000000000000000000000000000000000000000000000008684610cdf565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab5990610708907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a90600401611459565b600060405180830381600087803b15801561072257600080fd5b505af1158015610736573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d868686604051610784939291906114b6565b60405180910390a2506107976001600055565b5050505050565b6107a6610c53565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46107d081610c96565b6107d8610ca0565b6108037f00000000000000000000000000000000000000000000000000000000000000008785610cdf565b6040517faa0c0fc100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063aa0c0fc19061089f907f0000000000000000000000000000000000000000000000000000000000000000908b908b908b908b908a906004016115a4565b600060405180830381600087803b1580156108b957600080fd5b505af11580156108cd573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff167f5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff08787878660405161091d9493929190611615565b60405180910390a2506109306001600055565b505050505050565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61096281610c96565b61096a610ca0565b60048290556040518281527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c906020015b60405180910390a15050565b6109af610ca0565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b158015610a3d57600080fd5b505af1158015610797573d6000803e3d6000fd5b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a7b81610c96565b6106066110a3565b6000610a8e81610c96565b73ffffffffffffffffffffffffffffffffffffffff8216610adb576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354610b1f907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e49073ffffffffffffffffffffffffffffffffffffffff16610f67565b50600354610b64907f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb9073ffffffffffffffffffffffffffffffffffffffff16610f67565b50610b8f7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610e67565b50610bba7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610e67565b50600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091556040519081527fa38189053f94a2657ffb2b9fc651eddd1606a7cefc9f08d30eb72e3dbb51c1f19060200161099b565b600082815260026020526040902060010154610c4981610c96565b6105758383610f67565b600260005403610c8f576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61060681336110fc565b60015460ff1615610cdd576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6004547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d4d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d719190611641565b610d7b908461165a565b1115610db3576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201849052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b158015610e4a57600080fd5b505af1158015610e5e573d6000803e3d6000fd5b50505050505050565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610f5f57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610efd3390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016104a5565b5060006104a5565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610f5f57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104a5565b61102e61118c565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6110ab610ca0565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611079565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16611188576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440160405180910390fd5b5050565b60015460ff16610cdd576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156111da57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461120a57600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461123557600080fd5b919050565b60008060006060848603121561124f57600080fd5b61125884611211565b95602085013595506040909401359392505050565b60006020828403121561127f57600080fd5b5035919050565b6000806040838503121561129957600080fd5b823591506112a960208401611211565b90509250929050565b60008083601f8401126112c457600080fd5b50813567ffffffffffffffff8111156112dc57600080fd5b6020830191508360208285010111156112f457600080fd5b9250929050565b60008060008060006080868803121561131357600080fd5b61131c86611211565b945060208601359350604086013567ffffffffffffffff81111561133f57600080fd5b61134b888289016112b2565b96999598509660600135949350505050565b60008060008060008060a0878903121561137657600080fd5b61137f87611211565b955060208701359450604087013567ffffffffffffffff8111156113a257600080fd5b6113ae89828a016112b2565b90955093505060608701359150608087013567ffffffffffffffff8111156113d557600080fd5b87016080818a0312156113e757600080fd5b809150509295509295509295565b60006020828403121561140757600080fd5b61120a82611211565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006114ab608083018486611410565b979650505050505050565b8381526040602082015260006114d0604083018486611410565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff6114f782611211565b16825273ffffffffffffffffffffffffffffffffffffffff61151b60208301611211565b1660208301526040818101359083015260006060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261156357600080fd5b820160208101903567ffffffffffffffff81111561158057600080fd5b80360382131561158f57600080fd5b608060608601526114d0608086018284611410565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a0606082015260006115f660a083018587611410565b828103608084015261160881856114d9565b9998505050505050505050565b84815260606020820152600061162f606083018587611410565b82810360408401526114ab81856114d9565b60006020828403121561165357600080fd5b5051919050565b808201808211156104a5577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220f209ed67bed6b3811a3be45203488569f0270711c5284810a8d587b2d989b97a64736f6c634300081a00336080604052348015600f57600080fd5b506001600055610ed6806100246000396000f3fe60806040526004361061006e5760003560e01c8063c51316911161004b578063c5131691146100d5578063c9028a36146100f5578063e04d4f9714610115578063f05b6abf1461012857005b8063357fc5a214610077578063676cc054146100975780636ed70169146100c057005b3661007557005b005b34801561008357600080fd5b50610075610092366004610724565b610148565b6100aa6100a5366004610760565b6101de565b6040516100b7919061085b565b60405180910390f35b3480156100cc57600080fd5b50610075610211565b3480156100e157600080fd5b506100756100f0366004610724565b610246565b34801561010157600080fd5b5061007561011036600461086e565b610321565b6100756101233660046109ce565b61035d565b34801561013457600080fd5b50610075610143366004610aba565b6103a1565b6101506103d6565b61017273ffffffffffffffffffffffffffffffffffffffff8316338386610419565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101d96001600055565b505050565b6040516060907f3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee90600090a19392505050565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61024e6103d6565b600061025b600285610ba4565b905080600003610297576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102b973ffffffffffffffffffffffffffffffffffffffff8416338484610419565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101d96001600055565b7f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b3382604051610352929190610c28565b60405180910390a150565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa3334858585604051610394959493929190610d1a565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146338484846040516103949493929190610da4565b600260005403610412576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526104ae9085906104b4565b50505050565b60006104d673ffffffffffffffffffffffffffffffffffffffff84168361054f565b905080516000141580156104fb5750808060200190518101906104f99190610e67565b155b156101d9576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b606061055d83836000610564565b9392505050565b6060814710156105a2576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610546565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105cb9190610e84565b60006040518083038185875af1925050503d8060008114610608576040519150601f19603f3d011682016040523d82523d6000602084013e61060d565b606091505b509150915061061d868383610627565b9695505050505050565b60608261063c57610637826106b6565b61055d565b8151158015610660575073ffffffffffffffffffffffffffffffffffffffff84163b155b156106af576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610546565b508061055d565b8051156106c65780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff8116811461071f57600080fd5b919050565b60008060006060848603121561073957600080fd5b83359250610749602085016106fb565b9150610757604085016106fb565b90509250925092565b6000806000838503604081121561077657600080fd5b602081121561078457600080fd5b50839250602084013567ffffffffffffffff8111156107a257600080fd5b8401601f810186136107b357600080fd5b803567ffffffffffffffff8111156107ca57600080fd5b8660208284010111156107dc57600080fd5b939660209190910195509293505050565b60005b838110156108085781810151838201526020016107f0565b50506000910152565b600081518084526108298160208601602086016107ed565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061055d6020830184610811565b60006020828403121561088057600080fd5b813567ffffffffffffffff81111561089757600080fd5b82016080818503121561055d57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561091f5761091f6108a9565b604052919050565b600082601f83011261093857600080fd5b813567ffffffffffffffff811115610952576109526108a9565b61098360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016108d8565b81815284602083860101111561099857600080fd5b816020850160208301376000918101602001919091529392505050565b80151581146106f857600080fd5b803561071f816109b5565b6000806000606084860312156109e357600080fd5b833567ffffffffffffffff8111156109fa57600080fd5b610a0686828701610927565b935050602084013591506040840135610a1e816109b5565b809150509250925092565b600067ffffffffffffffff821115610a4357610a436108a9565b5060051b60200190565b600082601f830112610a5e57600080fd5b8135610a71610a6c82610a29565b6108d8565b8082825260208201915060208360051b860101925085831115610a9357600080fd5b602085015b83811015610ab0578035835260209283019201610a98565b5095945050505050565b600080600060608486031215610acf57600080fd5b833567ffffffffffffffff811115610ae657600080fd5b8401601f81018613610af757600080fd5b8035610b05610a6c82610a29565b8082825260208201915060208360051b850101925088831115610b2757600080fd5b602084015b83811015610b6957803567ffffffffffffffff811115610b4b57600080fd5b610b5a8b602083890101610927565b84525060209283019201610b2c565b509550505050602084013567ffffffffffffffff811115610b8957600080fd5b610b9586828701610a4d565b925050610757604085016109c3565b600082610bda577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610c66836106fb565b16604082015273ffffffffffffffffffffffffffffffffffffffff610c8d602084016106fb565b166060820152600080604084013590508060808401525060608301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610cd957600080fd5b830160208101903567ffffffffffffffff811115610cf657600080fd5b803603821315610d0557600080fd5b608060a085015261061d60c085018284610bdf565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610d4f60a0830186610811565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610d9a578151865260209586019590910190600101610d7c565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610e37577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610e22858351610811565b94506020938401939190910190600101610de8565b505050508281036040840152610e4d8186610d68565b915050610e5e606083018415159052565b95945050505050565b600060208284031215610e7957600080fd5b815161055d816109b5565b60008251610e968184602087016107ed565b919091019291505056fea2646970667358221220330cf118aee9a94a22a00087447b5ffda2f56eb86d305646a2a611f5fd89305464736f6c634300081a0033a2646970667358221220e515c69d1de6a2be247604d93205cf45f591bc74de0b873554a4fb5cec284c6964736f6c634300081a0033",
}

// GatewayEVMTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMTestMetaData.ABI instead.
var GatewayEVMTestABI = GatewayEVMTestMetaData.ABI

// GatewayEVMTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMTestMetaData.Bin instead.
var GatewayEVMTestBin = GatewayEVMTestMetaData.Bin

// DeployGatewayEVMTest deploys a new Ethereum contract, binding an instance of GatewayEVMTest to it.
func DeployGatewayEVMTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMTest, error) {
	parsed, err := GatewayEVMTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMTest{GatewayEVMTestCaller: GatewayEVMTestCaller{contract: contract}, GatewayEVMTestTransactor: GatewayEVMTestTransactor{contract: contract}, GatewayEVMTestFilterer: GatewayEVMTestFilterer{contract: contract}}, nil
}

// GatewayEVMTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMTest struct {
	GatewayEVMTestCaller     // Read-only binding to the contract
	GatewayEVMTestTransactor // Write-only binding to the contract
	GatewayEVMTestFilterer   // Log filterer for contract events
}

// GatewayEVMTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMTestSession struct {
	Contract     *GatewayEVMTest   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayEVMTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMTestCallerSession struct {
	Contract *GatewayEVMTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GatewayEVMTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMTestTransactorSession struct {
	Contract     *GatewayEVMTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GatewayEVMTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMTestRaw struct {
	Contract *GatewayEVMTest // Generic contract binding to access the raw methods on
}

// GatewayEVMTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMTestCallerRaw struct {
	Contract *GatewayEVMTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMTestTransactorRaw struct {
	Contract *GatewayEVMTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMTest creates a new instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMTest, error) {
	contract, err := bindGatewayEVMTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTest{GatewayEVMTestCaller: GatewayEVMTestCaller{contract: contract}, GatewayEVMTestTransactor: GatewayEVMTestTransactor{contract: contract}, GatewayEVMTestFilterer: GatewayEVMTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMTestCaller creates a new read-only instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMTestCaller, error) {
	contract, err := bindGatewayEVMTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestCaller{contract: contract}, nil
}

// NewGatewayEVMTestTransactor creates a new write-only instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMTestTransactor, error) {
	contract, err := bindGatewayEVMTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestTransactor{contract: contract}, nil
}

// NewGatewayEVMTestFilterer creates a new log filterer instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMTestFilterer, error) {
	contract, err := bindGatewayEVMTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestFilterer{contract: contract}, nil
}

// bindGatewayEVMTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMTest *GatewayEVMTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMTest.Contract.GatewayEVMTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMTest *GatewayEVMTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.GatewayEVMTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMTest *GatewayEVMTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.GatewayEVMTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMTest *GatewayEVMTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMTest *GatewayEVMTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMTest *GatewayEVMTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.contract.Transact(opts, method, params...)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) ASSETHANDLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "ASSET_HANDLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.ASSETHANDLERROLE(&_GatewayEVMTest.CallOpts)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.ASSETHANDLERROLE(&_GatewayEVMTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.DEFAULTADMINROLE(&_GatewayEVMTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.DEFAULTADMINROLE(&_GatewayEVMTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestSession) ISTEST() (bool, error) {
	return _GatewayEVMTest.Contract.ISTEST(&_GatewayEVMTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ISTEST() (bool, error) {
	return _GatewayEVMTest.Contract.ISTEST(&_GatewayEVMTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.PAUSERROLE(&_GatewayEVMTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.PAUSERROLE(&_GatewayEVMTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.TSSROLE(&_GatewayEVMTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.TSSROLE(&_GatewayEVMTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) WITHDRAWERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.WITHDRAWERROLE(&_GatewayEVMTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.WITHDRAWERROLE(&_GatewayEVMTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.ExcludeArtifacts(&_GatewayEVMTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.ExcludeArtifacts(&_GatewayEVMTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeContracts(&_GatewayEVMTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeContracts(&_GatewayEVMTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.ExcludeSelectors(&_GatewayEVMTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.ExcludeSelectors(&_GatewayEVMTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeSenders(&_GatewayEVMTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeSenders(&_GatewayEVMTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestSession) Failed() (bool, error) {
	return _GatewayEVMTest.Contract.Failed(&_GatewayEVMTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) Failed() (bool, error) {
	return _GatewayEVMTest.Contract.Failed(&_GatewayEVMTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMTest.Contract.TargetArtifactSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMTest.Contract.TargetArtifactSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.TargetArtifacts(&_GatewayEVMTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.TargetArtifacts(&_GatewayEVMTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetContracts(&_GatewayEVMTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetContracts(&_GatewayEVMTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMTest.Contract.TargetInterfaces(&_GatewayEVMTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMTest.Contract.TargetInterfaces(&_GatewayEVMTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.TargetSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.TargetSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetSenders(&_GatewayEVMTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetSenders(&_GatewayEVMTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.SetUp(&_GatewayEVMTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.SetUp(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x04b974f9.
//
// Solidity: function testExecuteFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteFailsIfDestinationIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteFailsIfDestinationIsZeroAddress")
}

// TestExecuteFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x04b974f9.
//
// Solidity: function testExecuteFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x04b974f9.
//
// Solidity: function testExecuteFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevert is a paid mutator transaction binding the contract method 0xfa18c09b.
//
// Solidity: function testExecuteRevert() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteRevert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteRevert")
}

// TestExecuteRevert is a paid mutator transaction binding the contract method 0xfa18c09b.
//
// Solidity: function testExecuteRevert() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteRevert() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevert(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevert is a paid mutator transaction binding the contract method 0xfa18c09b.
//
// Solidity: function testExecuteRevert() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteRevert() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevert(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51010e49.
//
// Solidity: function testExecuteRevertFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteRevertFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteRevertFailsIfReceiverIsZeroAddress")
}

// TestExecuteRevertFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51010e49.
//
// Solidity: function testExecuteRevertFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteRevertFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfReceiverIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51010e49.
//
// Solidity: function testExecuteRevertFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteRevertFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfReceiverIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0xcebad2a6.
//
// Solidity: function testExecuteRevertFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteRevertFailsIfSenderIsNotTSS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteRevertFailsIfSenderIsNotTSS")
}

// TestExecuteRevertFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0xcebad2a6.
//
// Solidity: function testExecuteRevertFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteRevertFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0xcebad2a6.
//
// Solidity: function testExecuteRevertFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteRevertFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xe6afc790.
//
// Solidity: function testExecuteWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteWithERC20FailsIfNotCustodyOrConnector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteWithERC20FailsIfNotCustodyOrConnector")
}

// TestExecuteWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xe6afc790.
//
// Solidity: function testExecuteWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xe6afc790.
//
// Solidity: function testExecuteWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x6bdd212b.
//
// Solidity: function testExecuteWithMsgContextFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteWithMsgContextFailsIfDestinationIsZeroAddress")
}

// TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x6bdd212b.
//
// Solidity: function testExecuteWithMsgContextFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x6bdd212b.
//
// Solidity: function testExecuteWithMsgContextFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParams is a paid mutator transaction binding the contract method 0x44671b94.
//
// Solidity: function testForwardCallToReceiveNoParams() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNoParams(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNoParams")
}

// TestForwardCallToReceiveNoParams is a paid mutator transaction binding the contract method 0x44671b94.
//
// Solidity: function testForwardCallToReceiveNoParams() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNoParams() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParams(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParams is a paid mutator transaction binding the contract method 0x44671b94.
//
// Solidity: function testForwardCallToReceiveNoParams() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNoParams() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParams(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParamsTogglePause is a paid mutator transaction binding the contract method 0xdd51e82f.
//
// Solidity: function testForwardCallToReceiveNoParamsTogglePause() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNoParamsTogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNoParamsTogglePause")
}

// TestForwardCallToReceiveNoParamsTogglePause is a paid mutator transaction binding the contract method 0xdd51e82f.
//
// Solidity: function testForwardCallToReceiveNoParamsTogglePause() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNoParamsTogglePause() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParamsTogglePause(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParamsTogglePause is a paid mutator transaction binding the contract method 0xdd51e82f.
//
// Solidity: function testForwardCallToReceiveNoParamsTogglePause() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNoParamsTogglePause() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParamsTogglePause(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayable is a paid mutator transaction binding the contract method 0xf68bd1c0.
//
// Solidity: function testForwardCallToReceiveNonPayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNonPayable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNonPayable")
}

// TestForwardCallToReceiveNonPayable is a paid mutator transaction binding the contract method 0xf68bd1c0.
//
// Solidity: function testForwardCallToReceiveNonPayable() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNonPayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayable(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayable is a paid mutator transaction binding the contract method 0xf68bd1c0.
//
// Solidity: function testForwardCallToReceiveNonPayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNonPayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayable(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x7d7f772a.
//
// Solidity: function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS")
}

// TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x7d7f772a.
//
// Solidity: function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x7d7f772a.
//
// Solidity: function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x43fd8c7d.
//
// Solidity: function testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS")
}

// TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x43fd8c7d.
//
// Solidity: function testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x43fd8c7d.
//
// Solidity: function testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnCallFails is a paid mutator transaction binding the contract method 0xd38b66cd.
//
// Solidity: function testForwardCallToReceiveOnCallFails() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveOnCallFails(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveOnCallFails")
}

// TestForwardCallToReceiveOnCallFails is a paid mutator transaction binding the contract method 0xd38b66cd.
//
// Solidity: function testForwardCallToReceiveOnCallFails() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveOnCallFails() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnCallFails(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnCallFails is a paid mutator transaction binding the contract method 0xd38b66cd.
//
// Solidity: function testForwardCallToReceiveOnCallFails() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveOnCallFails() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnCallFails(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnCallUsingAuthCall is a paid mutator transaction binding the contract method 0xa397ffd2.
//
// Solidity: function testForwardCallToReceiveOnCallUsingAuthCall() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveOnCallUsingAuthCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveOnCallUsingAuthCall")
}

// TestForwardCallToReceiveOnCallUsingAuthCall is a paid mutator transaction binding the contract method 0xa397ffd2.
//
// Solidity: function testForwardCallToReceiveOnCallUsingAuthCall() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveOnCallUsingAuthCall() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnCallUsingAuthCall(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnCallUsingAuthCall is a paid mutator transaction binding the contract method 0xa397ffd2.
//
// Solidity: function testForwardCallToReceiveOnCallUsingAuthCall() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveOnCallUsingAuthCall() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnCallUsingAuthCall(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnRevertFails is a paid mutator transaction binding the contract method 0x130daf59.
//
// Solidity: function testForwardCallToReceiveOnRevertFails() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveOnRevertFails(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveOnRevertFails")
}

// TestForwardCallToReceiveOnRevertFails is a paid mutator transaction binding the contract method 0x130daf59.
//
// Solidity: function testForwardCallToReceiveOnRevertFails() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveOnRevertFails() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnRevertFails(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnRevertFails is a paid mutator transaction binding the contract method 0x130daf59.
//
// Solidity: function testForwardCallToReceiveOnRevertFails() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveOnRevertFails() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnRevertFails(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0xfe7bdbb2.
//
// Solidity: function testForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceivePayable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceivePayable")
}

// TestForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0xfe7bdbb2.
//
// Solidity: function testForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceivePayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceivePayable(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0xfe7bdbb2.
//
// Solidity: function testForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceivePayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceivePayable(&_GatewayEVMTest.TransactOpts)
}

// TestRevertWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xb124dbf5.
//
// Solidity: function testRevertWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestRevertWithERC20FailsIfNotCustodyOrConnector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testRevertWithERC20FailsIfNotCustodyOrConnector")
}

// TestRevertWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xb124dbf5.
//
// Solidity: function testRevertWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestRevertWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestRevertWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestRevertWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xb124dbf5.
//
// Solidity: function testRevertWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestRevertWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestRevertWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xccf20616.
//
// Solidity: function testSetConnectorFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetConnectorFailsIfSenderIsNotAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetConnectorFailsIfSenderIsNotAdmin")
}

// TestSetConnectorFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xccf20616.
//
// Solidity: function testSetConnectorFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetConnectorFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xccf20616.
//
// Solidity: function testSetConnectorFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetConnectorFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSet is a paid mutator transaction binding the contract method 0x7ebf744f.
//
// Solidity: function testSetConnectorFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetConnectorFailsIfSet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetConnectorFailsIfSet")
}

// TestSetConnectorFailsIfSet is a paid mutator transaction binding the contract method 0x7ebf744f.
//
// Solidity: function testSetConnectorFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetConnectorFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSet is a paid mutator transaction binding the contract method 0x7ebf744f.
//
// Solidity: function testSetConnectorFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetConnectorFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfZero is a paid mutator transaction binding the contract method 0xce5871e1.
//
// Solidity: function testSetConnectorFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetConnectorFailsIfZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetConnectorFailsIfZero")
}

// TestSetConnectorFailsIfZero is a paid mutator transaction binding the contract method 0xce5871e1.
//
// Solidity: function testSetConnectorFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetConnectorFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfZero is a paid mutator transaction binding the contract method 0xce5871e1.
//
// Solidity: function testSetConnectorFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetConnectorFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xa56f7a4b.
//
// Solidity: function testSetCustodyFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetCustodyFailsIfSenderIsNotAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetCustodyFailsIfSenderIsNotAdmin")
}

// TestSetCustodyFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xa56f7a4b.
//
// Solidity: function testSetCustodyFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetCustodyFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xa56f7a4b.
//
// Solidity: function testSetCustodyFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetCustodyFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSet is a paid mutator transaction binding the contract method 0x1855c337.
//
// Solidity: function testSetCustodyFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetCustodyFailsIfSet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetCustodyFailsIfSet")
}

// TestSetCustodyFailsIfSet is a paid mutator transaction binding the contract method 0x1855c337.
//
// Solidity: function testSetCustodyFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetCustodyFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSet is a paid mutator transaction binding the contract method 0x1855c337.
//
// Solidity: function testSetCustodyFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetCustodyFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfZero is a paid mutator transaction binding the contract method 0x1226c655.
//
// Solidity: function testSetCustodyFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetCustodyFailsIfZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetCustodyFailsIfZero")
}

// TestSetCustodyFailsIfZero is a paid mutator transaction binding the contract method 0x1226c655.
//
// Solidity: function testSetCustodyFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetCustodyFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfZero is a paid mutator transaction binding the contract method 0x1226c655.
//
// Solidity: function testSetCustodyFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetCustodyFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestTSSUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testTSSUpgrade")
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgrade(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgrade(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testTSSUpgradeFailsIfSenderIsNotTSSUpdater")
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestTSSUpgradeFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testTSSUpgradeFailsIfZeroAddress")
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// GatewayEVMTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayEVMTest contract.
type GatewayEVMTestCalledIterator struct {
	Event *GatewayEVMTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestCalled)
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
		it.Event = new(GatewayEVMTestCalled)
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
func (it *GatewayEVMTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestCalled represents a Called event raised by the GatewayEVMTest contract.
type GatewayEVMTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestCalledIterator{contract: _GatewayEVMTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestCalled)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseCalled(log types.Log) (*GatewayEVMTestCalled, error) {
	event := new(GatewayEVMTestCalled)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayEVMTest contract.
type GatewayEVMTestDepositedIterator struct {
	Event *GatewayEVMTestDeposited // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestDeposited)
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
		it.Event = new(GatewayEVMTestDeposited)
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
func (it *GatewayEVMTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestDeposited represents a Deposited event raised by the GatewayEVMTest contract.
type GatewayEVMTestDeposited struct {
	Recipient []byte
	Asset     common.Address
	Amount    *big.Int
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterDeposited(opts *bind.FilterOpts, asset []common.Address) (*GatewayEVMTestDepositedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestDepositedIterator{contract: _GatewayEVMTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestDeposited, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestDeposited)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseDeposited(log types.Log) (*GatewayEVMTestDeposited, error) {
	event := new(GatewayEVMTestDeposited)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestDeposited0Iterator is returned from FilterDeposited0 and is used to iterate over the raw logs and unpacked data for Deposited0 events raised by the GatewayEVMTest contract.
type GatewayEVMTestDeposited0Iterator struct {
	Event *GatewayEVMTestDeposited0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestDeposited0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestDeposited0)
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
		it.Event = new(GatewayEVMTestDeposited0)
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
func (it *GatewayEVMTestDeposited0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestDeposited0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestDeposited0 represents a Deposited0 event raised by the GatewayEVMTest contract.
type GatewayEVMTestDeposited0 struct {
	Sender        common.Address
	Receiver      common.Address
	Amount        *big.Int
	Asset         common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposited0 is a free log retrieval operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterDeposited0(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMTestDeposited0Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Deposited0", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestDeposited0Iterator{contract: _GatewayEVMTest.contract, event: "Deposited0", logs: logs, sub: sub}, nil
}

// WatchDeposited0 is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchDeposited0(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestDeposited0, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Deposited0", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestDeposited0)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited0", log); err != nil {
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

// ParseDeposited0 is a log parse operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseDeposited0(log types.Log) (*GatewayEVMTestDeposited0, error) {
	event := new(GatewayEVMTestDeposited0)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedIterator struct {
	Event *GatewayEVMTestExecuted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestExecuted)
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
		it.Event = new(GatewayEVMTestExecuted)
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
func (it *GatewayEVMTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestExecuted represents a Executed event raised by the GatewayEVMTest contract.
type GatewayEVMTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestExecutedIterator{contract: _GatewayEVMTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestExecuted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Executed", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseExecuted(log types.Log) (*GatewayEVMTestExecuted, error) {
	event := new(GatewayEVMTestExecuted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMTestExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestExecutedWithERC20)
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
		it.Event = new(GatewayEVMTestExecutedWithERC20)
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
func (it *GatewayEVMTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestExecutedWithERC20Iterator{contract: _GatewayEVMTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestExecutedWithERC20)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMTestExecutedWithERC20, error) {
	event := new(GatewayEVMTestExecutedWithERC20)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedERC20Iterator struct {
	Event *GatewayEVMTestReceivedERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedERC20)
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
		it.Event = new(GatewayEVMTestReceivedERC20)
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
func (it *GatewayEVMTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedERC20 represents a ReceivedERC20 event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*GatewayEVMTestReceivedERC20Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedERC20Iterator{contract: _GatewayEVMTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedERC20)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
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

// ParseReceivedERC20 is a log parse operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedERC20(log types.Log) (*GatewayEVMTestReceivedERC20, error) {
	event := new(GatewayEVMTestReceivedERC20)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNoParamsIterator struct {
	Event *GatewayEVMTestReceivedNoParams // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedNoParams)
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
		it.Event = new(GatewayEVMTestReceivedNoParams)
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
func (it *GatewayEVMTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedNoParams represents a ReceivedNoParams event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*GatewayEVMTestReceivedNoParamsIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedNoParamsIterator{contract: _GatewayEVMTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedNoParams)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
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

// ParseReceivedNoParams is a log parse operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedNoParams(log types.Log) (*GatewayEVMTestReceivedNoParams, error) {
	event := new(GatewayEVMTestReceivedNoParams)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNonPayableIterator struct {
	Event *GatewayEVMTestReceivedNonPayable // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedNonPayable)
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
		it.Event = new(GatewayEVMTestReceivedNonPayable)
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
func (it *GatewayEVMTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedNonPayable represents a ReceivedNonPayable event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*GatewayEVMTestReceivedNonPayableIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedNonPayableIterator{contract: _GatewayEVMTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedNonPayable)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
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

// ParseReceivedNonPayable is a log parse operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedNonPayable(log types.Log) (*GatewayEVMTestReceivedNonPayable, error) {
	event := new(GatewayEVMTestReceivedNonPayable)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedOnCallIterator is returned from FilterReceivedOnCall and is used to iterate over the raw logs and unpacked data for ReceivedOnCall events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedOnCallIterator struct {
	Event *GatewayEVMTestReceivedOnCall // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestReceivedOnCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedOnCall)
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
		it.Event = new(GatewayEVMTestReceivedOnCall)
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
func (it *GatewayEVMTestReceivedOnCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedOnCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedOnCall represents a ReceivedOnCall event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedOnCall struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCall is a free log retrieval operation binding the contract event 0x3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee.
//
// Solidity: event ReceivedOnCall()
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedOnCall(opts *bind.FilterOpts) (*GatewayEVMTestReceivedOnCallIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedOnCallIterator{contract: _GatewayEVMTest.contract, event: "ReceivedOnCall", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCall is a free log subscription operation binding the contract event 0x3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee.
//
// Solidity: event ReceivedOnCall()
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedOnCall(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedOnCall) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedOnCall)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
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

// ParseReceivedOnCall is a log parse operation binding the contract event 0x3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee.
//
// Solidity: event ReceivedOnCall()
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedOnCall(log types.Log) (*GatewayEVMTestReceivedOnCall, error) {
	event := new(GatewayEVMTestReceivedOnCall)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedPayableIterator struct {
	Event *GatewayEVMTestReceivedPayable // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedPayable)
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
		it.Event = new(GatewayEVMTestReceivedPayable)
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
func (it *GatewayEVMTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedPayable represents a ReceivedPayable event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedPayable struct {
	Sender common.Address
	Value  *big.Int
	Str    string
	Num    *big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedPayable is a free log retrieval operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*GatewayEVMTestReceivedPayableIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedPayableIterator{contract: _GatewayEVMTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedPayable)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
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

// ParseReceivedPayable is a log parse operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedPayable(log types.Log) (*GatewayEVMTestReceivedPayable, error) {
	event := new(GatewayEVMTestReceivedPayable)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedRevertIterator struct {
	Event *GatewayEVMTestReceivedRevert // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedRevert)
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
		it.Event = new(GatewayEVMTestReceivedRevert)
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
func (it *GatewayEVMTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedRevert represents a ReceivedRevert event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*GatewayEVMTestReceivedRevertIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedRevertIterator{contract: _GatewayEVMTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedRevert)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
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

// ParseReceivedRevert is a log parse operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedRevert(log types.Log) (*GatewayEVMTestReceivedRevert, error) {
	event := new(GatewayEVMTestReceivedRevert)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVMTest contract.
type GatewayEVMTestRevertedIterator struct {
	Event *GatewayEVMTestReverted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReverted)
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
		it.Event = new(GatewayEVMTestReverted)
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
func (it *GatewayEVMTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReverted represents a Reverted event raised by the GatewayEVMTest contract.
type GatewayEVMTestReverted struct {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestRevertedIterator{contract: _GatewayEVMTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReverted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReverted(log types.Log) (*GatewayEVMTestReverted, error) {
	event := new(GatewayEVMTestReverted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUnwhitelistedIterator is returned from FilterUnwhitelisted and is used to iterate over the raw logs and unpacked data for Unwhitelisted events raised by the GatewayEVMTest contract.
type GatewayEVMTestUnwhitelistedIterator struct {
	Event *GatewayEVMTestUnwhitelisted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUnwhitelisted)
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
		it.Event = new(GatewayEVMTestUnwhitelisted)
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
func (it *GatewayEVMTestUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUnwhitelisted represents a Unwhitelisted event raised by the GatewayEVMTest contract.
type GatewayEVMTestUnwhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnwhitelisted is a free log retrieval operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUnwhitelisted(opts *bind.FilterOpts, token []common.Address) (*GatewayEVMTestUnwhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUnwhitelistedIterator{contract: _GatewayEVMTest.contract, event: "Unwhitelisted", logs: logs, sub: sub}, nil
}

// WatchUnwhitelisted is a free log subscription operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUnwhitelisted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUnwhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUnwhitelisted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
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

// ParseUnwhitelisted is a log parse operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUnwhitelisted(log types.Log) (*GatewayEVMTestUnwhitelisted, error) {
	event := new(GatewayEVMTestUnwhitelisted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUpdatedCustodyTSSAddressIterator is returned from FilterUpdatedCustodyTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedCustodyTSSAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedCustodyTSSAddressIterator struct {
	Event *GatewayEVMTestUpdatedCustodyTSSAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestUpdatedCustodyTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUpdatedCustodyTSSAddress)
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
		it.Event = new(GatewayEVMTestUpdatedCustodyTSSAddress)
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
func (it *GatewayEVMTestUpdatedCustodyTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUpdatedCustodyTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUpdatedCustodyTSSAddress represents a UpdatedCustodyTSSAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedCustodyTSSAddress struct {
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCustodyTSSAddress is a free log retrieval operation binding the contract event 0x086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba.
//
// Solidity: event UpdatedCustodyTSSAddress(address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUpdatedCustodyTSSAddress(opts *bind.FilterOpts) (*GatewayEVMTestUpdatedCustodyTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUpdatedCustodyTSSAddressIterator{contract: _GatewayEVMTest.contract, event: "UpdatedCustodyTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedCustodyTSSAddress is a free log subscription operation binding the contract event 0x086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba.
//
// Solidity: event UpdatedCustodyTSSAddress(address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUpdatedCustodyTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUpdatedCustodyTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUpdatedCustodyTSSAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
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

// ParseUpdatedCustodyTSSAddress is a log parse operation binding the contract event 0x086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba.
//
// Solidity: event UpdatedCustodyTSSAddress(address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUpdatedCustodyTSSAddress(log types.Log) (*GatewayEVMTestUpdatedCustodyTSSAddress, error) {
	event := new(GatewayEVMTestUpdatedCustodyTSSAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedGatewayTSSAddressIterator struct {
	Event *GatewayEVMTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUpdatedGatewayTSSAddress)
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
		it.Event = new(GatewayEVMTestUpdatedGatewayTSSAddress)
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
func (it *GatewayEVMTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedGatewayTSSAddress struct {
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*GatewayEVMTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUpdatedGatewayTSSAddressIterator{contract: _GatewayEVMTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUpdatedGatewayTSSAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*GatewayEVMTestUpdatedGatewayTSSAddress, error) {
	event := new(GatewayEVMTestUpdatedGatewayTSSAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the GatewayEVMTest contract.
type GatewayEVMTestWhitelistedIterator struct {
	Event *GatewayEVMTestWhitelisted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWhitelisted)
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
		it.Event = new(GatewayEVMTestWhitelisted)
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
func (it *GatewayEVMTestWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWhitelisted represents a Whitelisted event raised by the GatewayEVMTest contract.
type GatewayEVMTestWhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWhitelisted(opts *bind.FilterOpts, token []common.Address) (*GatewayEVMTestWhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWhitelistedIterator{contract: _GatewayEVMTest.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWhitelisted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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

// ParseWhitelisted is a log parse operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWhitelisted(log types.Log) (*GatewayEVMTestWhitelisted, error) {
	event := new(GatewayEVMTestWhitelisted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnIterator struct {
	Event *GatewayEVMTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWithdrawn)
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
		it.Event = new(GatewayEVMTestWithdrawn)
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
func (it *GatewayEVMTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWithdrawn represents a Withdrawn event raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawn struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWithdrawnIterator{contract: _GatewayEVMTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWithdrawn, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWithdrawn)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWithdrawn(log types.Log) (*GatewayEVMTestWithdrawn, error) {
	event := new(GatewayEVMTestWithdrawn)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndCalledIterator struct {
	Event *GatewayEVMTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWithdrawnAndCalled)
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
		it.Event = new(GatewayEVMTestWithdrawnAndCalled)
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
func (it *GatewayEVMTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndCalled struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWithdrawnAndCalledIterator{contract: _GatewayEVMTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWithdrawnAndCalled, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWithdrawnAndCalled)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayEVMTestWithdrawnAndCalled, error) {
	event := new(GatewayEVMTestWithdrawnAndCalled)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndRevertedIterator struct {
	Event *GatewayEVMTestWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWithdrawnAndReverted)
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
		it.Event = new(GatewayEVMTestWithdrawnAndReverted)
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
func (it *GatewayEVMTestWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndReverted struct {
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWithdrawnAndRevertedIterator{contract: _GatewayEVMTest.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWithdrawnAndReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWithdrawnAndReverted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*GatewayEVMTestWithdrawnAndReverted, error) {
	event := new(GatewayEVMTestWithdrawnAndReverted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogIterator struct {
	Event *GatewayEVMTestLog // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLog)
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
		it.Event = new(GatewayEVMTestLog)
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
func (it *GatewayEVMTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLog represents a Log event raised by the GatewayEVMTest contract.
type GatewayEVMTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayEVMTestLogIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogIterator{contract: _GatewayEVMTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLog)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLog(log types.Log) (*GatewayEVMTestLog, error) {
	event := new(GatewayEVMTestLog)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogAddressIterator struct {
	Event *GatewayEVMTestLogAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogAddress)
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
		it.Event = new(GatewayEVMTestLogAddress)
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
func (it *GatewayEVMTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogAddress represents a LogAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayEVMTestLogAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogAddressIterator{contract: _GatewayEVMTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_address", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogAddress(log types.Log) (*GatewayEVMTestLogAddress, error) {
	event := new(GatewayEVMTestLogAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArrayIterator struct {
	Event *GatewayEVMTestLogArray // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogArray)
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
		it.Event = new(GatewayEVMTestLogArray)
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
func (it *GatewayEVMTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogArray represents a LogArray event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayEVMTestLogArrayIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogArrayIterator{contract: _GatewayEVMTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogArray)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array", log); err != nil {
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

// ParseLogArray is a log parse operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogArray(log types.Log) (*GatewayEVMTestLogArray, error) {
	event := new(GatewayEVMTestLogArray)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray0Iterator struct {
	Event *GatewayEVMTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogArray0)
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
		it.Event = new(GatewayEVMTestLogArray0)
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
func (it *GatewayEVMTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogArray0 represents a LogArray0 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayEVMTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogArray0Iterator{contract: _GatewayEVMTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogArray0)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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

// ParseLogArray0 is a log parse operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogArray0(log types.Log) (*GatewayEVMTestLogArray0, error) {
	event := new(GatewayEVMTestLogArray0)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray1Iterator struct {
	Event *GatewayEVMTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogArray1)
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
		it.Event = new(GatewayEVMTestLogArray1)
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
func (it *GatewayEVMTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogArray1 represents a LogArray1 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayEVMTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogArray1Iterator{contract: _GatewayEVMTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogArray1)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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

// ParseLogArray1 is a log parse operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogArray1(log types.Log) (*GatewayEVMTestLogArray1, error) {
	event := new(GatewayEVMTestLogArray1)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytesIterator struct {
	Event *GatewayEVMTestLogBytes // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogBytes)
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
		it.Event = new(GatewayEVMTestLogBytes)
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
func (it *GatewayEVMTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogBytes represents a LogBytes event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayEVMTestLogBytesIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogBytesIterator{contract: _GatewayEVMTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogBytes)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogBytes(log types.Log) (*GatewayEVMTestLogBytes, error) {
	event := new(GatewayEVMTestLogBytes)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytes32Iterator struct {
	Event *GatewayEVMTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogBytes32)
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
		it.Event = new(GatewayEVMTestLogBytes32)
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
func (it *GatewayEVMTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogBytes32 represents a LogBytes32 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayEVMTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogBytes32Iterator{contract: _GatewayEVMTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogBytes32)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogBytes32(log types.Log) (*GatewayEVMTestLogBytes32, error) {
	event := new(GatewayEVMTestLogBytes32)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogIntIterator struct {
	Event *GatewayEVMTestLogInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogInt)
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
		it.Event = new(GatewayEVMTestLogInt)
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
func (it *GatewayEVMTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogInt represents a LogInt event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayEVMTestLogIntIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogIntIterator{contract: _GatewayEVMTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogInt)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_int", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogInt(log types.Log) (*GatewayEVMTestLogInt, error) {
	event := new(GatewayEVMTestLogInt)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedAddressIterator struct {
	Event *GatewayEVMTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedAddress)
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
		it.Event = new(GatewayEVMTestLogNamedAddress)
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
func (it *GatewayEVMTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedAddressIterator{contract: _GatewayEVMTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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

// ParseLogNamedAddress is a log parse operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayEVMTestLogNamedAddress, error) {
	event := new(GatewayEVMTestLogNamedAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArrayIterator struct {
	Event *GatewayEVMTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedArray)
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
		it.Event = new(GatewayEVMTestLogNamedArray)
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
func (it *GatewayEVMTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedArray represents a LogNamedArray event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedArrayIterator{contract: _GatewayEVMTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedArray)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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

// ParseLogNamedArray is a log parse operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayEVMTestLogNamedArray, error) {
	event := new(GatewayEVMTestLogNamedArray)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray0Iterator struct {
	Event *GatewayEVMTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedArray0)
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
		it.Event = new(GatewayEVMTestLogNamedArray0)
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
func (it *GatewayEVMTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedArray0Iterator{contract: _GatewayEVMTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedArray0)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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

// ParseLogNamedArray0 is a log parse operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayEVMTestLogNamedArray0, error) {
	event := new(GatewayEVMTestLogNamedArray0)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray1Iterator struct {
	Event *GatewayEVMTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedArray1)
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
		it.Event = new(GatewayEVMTestLogNamedArray1)
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
func (it *GatewayEVMTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedArray1Iterator{contract: _GatewayEVMTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedArray1)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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

// ParseLogNamedArray1 is a log parse operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayEVMTestLogNamedArray1, error) {
	event := new(GatewayEVMTestLogNamedArray1)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytesIterator struct {
	Event *GatewayEVMTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedBytes)
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
		it.Event = new(GatewayEVMTestLogNamedBytes)
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
func (it *GatewayEVMTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedBytesIterator{contract: _GatewayEVMTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedBytes)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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

// ParseLogNamedBytes is a log parse operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayEVMTestLogNamedBytes, error) {
	event := new(GatewayEVMTestLogNamedBytes)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytes32Iterator struct {
	Event *GatewayEVMTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedBytes32)
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
		it.Event = new(GatewayEVMTestLogNamedBytes32)
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
func (it *GatewayEVMTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedBytes32Iterator{contract: _GatewayEVMTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedBytes32)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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

// ParseLogNamedBytes32 is a log parse operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayEVMTestLogNamedBytes32, error) {
	event := new(GatewayEVMTestLogNamedBytes32)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalIntIterator struct {
	Event *GatewayEVMTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedDecimalInt)
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
		it.Event = new(GatewayEVMTestLogNamedDecimalInt)
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
func (it *GatewayEVMTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedDecimalIntIterator{contract: _GatewayEVMTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedDecimalInt)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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

// ParseLogNamedDecimalInt is a log parse operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayEVMTestLogNamedDecimalInt, error) {
	event := new(GatewayEVMTestLogNamedDecimalInt)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalUintIterator struct {
	Event *GatewayEVMTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedDecimalUint)
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
		it.Event = new(GatewayEVMTestLogNamedDecimalUint)
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
func (it *GatewayEVMTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedDecimalUintIterator{contract: _GatewayEVMTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedDecimalUint)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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

// ParseLogNamedDecimalUint is a log parse operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayEVMTestLogNamedDecimalUint, error) {
	event := new(GatewayEVMTestLogNamedDecimalUint)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedIntIterator struct {
	Event *GatewayEVMTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedInt)
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
		it.Event = new(GatewayEVMTestLogNamedInt)
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
func (it *GatewayEVMTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedInt represents a LogNamedInt event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedIntIterator{contract: _GatewayEVMTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedInt)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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

// ParseLogNamedInt is a log parse operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayEVMTestLogNamedInt, error) {
	event := new(GatewayEVMTestLogNamedInt)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedStringIterator struct {
	Event *GatewayEVMTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedString)
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
		it.Event = new(GatewayEVMTestLogNamedString)
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
func (it *GatewayEVMTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedString represents a LogNamedString event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedStringIterator{contract: _GatewayEVMTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedString)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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

// ParseLogNamedString is a log parse operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedString(log types.Log) (*GatewayEVMTestLogNamedString, error) {
	event := new(GatewayEVMTestLogNamedString)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedUintIterator struct {
	Event *GatewayEVMTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedUint)
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
		it.Event = new(GatewayEVMTestLogNamedUint)
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
func (it *GatewayEVMTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedUint represents a LogNamedUint event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedUintIterator{contract: _GatewayEVMTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedUint)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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

// ParseLogNamedUint is a log parse operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayEVMTestLogNamedUint, error) {
	event := new(GatewayEVMTestLogNamedUint)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogStringIterator struct {
	Event *GatewayEVMTestLogString // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogString)
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
		it.Event = new(GatewayEVMTestLogString)
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
func (it *GatewayEVMTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogString represents a LogString event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayEVMTestLogStringIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogStringIterator{contract: _GatewayEVMTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogString)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_string", log); err != nil {
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

// ParseLogString is a log parse operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogString(log types.Log) (*GatewayEVMTestLogString, error) {
	event := new(GatewayEVMTestLogString)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogUintIterator struct {
	Event *GatewayEVMTestLogUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogUint)
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
		it.Event = new(GatewayEVMTestLogUint)
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
func (it *GatewayEVMTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogUint represents a LogUint event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayEVMTestLogUintIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogUintIterator{contract: _GatewayEVMTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogUint)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogUint(log types.Log) (*GatewayEVMTestLogUint, error) {
	event := new(GatewayEVMTestLogUint)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogsIterator struct {
	Event *GatewayEVMTestLogs // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogs)
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
		it.Event = new(GatewayEVMTestLogs)
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
func (it *GatewayEVMTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogs represents a Logs event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayEVMTestLogsIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogsIterator{contract: _GatewayEVMTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogs)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "logs", log); err != nil {
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

// ParseLogs is a log parse operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogs(log types.Log) (*GatewayEVMTestLogs, error) {
	event := new(GatewayEVMTestLogs)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
