// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20custody

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

// ERC20CustodyTestMetaData contains all meta data concerning the ERC20CustodyTest contract.
var ERC20CustodyTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ASSET_HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WHITELISTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testDepositLegacy\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositLegacyFailsIfNotSupported\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositLegacyFailsIfTokenNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20PartialThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNoParamsThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testNewCustodyFailsIfAddressesAreZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUnwhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUnwhitelistFailsIfSenderIsNotWhitelister\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUnwhitelistFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhitelistFailsIfSenderIsNotWhitelister\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhitelistFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LegacyMethodsNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5062011c08806200003e6000396000f3fe608060405234801561001057600080fd5b506004361061030a5760003560e01c806385f438c11161019c578063ba414fa6116100ee578063f0c8e7e011610097578063fa7626d411610071578063fa7626d41461056c578063fb176c1214610579578063fe8e5f1b1461058157600080fd5b8063f0c8e7e014610554578063f4221f081461055c578063fa2a70741461056457600080fd5b8063e20c9f71116100c8578063e20c9f711461051d578063e63ab1e914610525578063eb1ce7f91461054c57600080fd5b8063ba414fa6146104f5578063c713f8271461050d578063cbd57e2f1461051557600080fd5b8063a3f9d0e011610150578063b0464fdc1161012a578063b0464fdc146104dd578063b421ca70146104e5578063b5508aa9146104ed57600080fd5b8063a3f9d0e0146104a6578063a4943deb146104ae578063a783c789146104b657600080fd5b8063916a17c611610181578063916a17c6146104815780639918c1c2146104965780639fc7fd551461049e57600080fd5b806385f438c1146104525780639158c6231461047957600080fd5b806349c783dd116102605780636a621854116102095780637e91c50f116101e35780637e91c50f1461042d57806382c529921461043557806385226c811461043d57600080fd5b80636a621854146104155780637099d6f81461041d57806371149c941461042557600080fd5b8063570618e11161023a578063570618e1146103a45780635d62c860146103d957806366d9a9a01461040057600080fd5b806349c783dd1461038c5780634b5838d21461039457806351ecdf3c1461039c57600080fd5b80632ade3880116102c25780633e73ecb41161029c5780633e73ecb4146103745780633ee929231461037c5780633f7286f41461038457600080fd5b80632ade38801461034f5780632be6a162146103645780633e5e3c231461036c57600080fd5b80631779672f116102f35780631779672f146103215780631ed7831c14610329578063284cb9291461034757600080fd5b80630a9254e41461030f5780630eee72a914610319575b600080fd5b610317610589565b005b610317610fe7565b61031761128c565b6103316113dd565b60405161033e919061b1a4565b60405180910390f35b61031761143f565b610357611723565b60405161033e919061b240565b610317611865565b610331611a57565b610317611ab7565b61031761202c565b6103316120fb565b61031761215b565b610317612498565b6103176127b6565b6103cb7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b60405190815260200161033e565b6103cb7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b981565b610408612987565b60405161033e919061b3a6565b610317612b09565b610317612bd5565b610317612e82565b610317613697565b610317613822565b610445613a96565b60405161033e919061b444565b6103cb7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b610317613b66565b610489613c34565b60405161033e919061b4bb565b610317613d2f565b610317614029565b6103176140f7565b6103176146ff565b6103cb7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b610489614927565b610317614a22565b610445614f57565b6104fd615027565b604051901515815260200161033e565b6103176150fb565b610317615d9e565b610331615e5f565b6103cb7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b610317615ebf565b610317615fcb565b61031761618d565b6103176163fa565b601f546104fd9060ff1681565b6103176166d8565b610317616d56565b602680547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602780548216611234179055602880549091166156781790556040516105db9061b0b3565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f080158015610660573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03929092169190911790556040516106a59061b0b3565b604080825260049082018190527f7a6574610000000000000000000000000000000000000000000000000000000060608301526080602083018190528201527f5a4554410000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f080158015610729573d6000803e3d6000fd5b50602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260285460265492519085166024820152604481019390935292166064820152610818919060840160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b00000000000000000000000000000000000000000000000000000000179052616f47565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff000000000000000000000000000000000000000090911681179091556028546026546040519293918216929116906108a49061b0c1565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f0801580156108e0573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316179055602054602554602854602654604051938516949283169391831692169061093b9061b0cf565b6001600160a01b039485168152928416602084015290831660408301529091166060820152608001604051809103906000f08015801561097f573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03929092169190911790556040516109c49061b0dd565b604051809103906000f0801580156109e0573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556028546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b158015610a8c57600080fd5b505af1158015610aa0573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015610b1657600080fd5b505af1158015610b2a573d6000803e3d6000fd5b50506020546022546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f9150602401600060405180830381600087803b158015610b9057600080fd5b505af1158015610ba4573d6000803e3d6000fd5b50506020546023546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef9150602401600060405180830381600087803b158015610c0a57600080fd5b505af1158015610c1e573d6000803e3d6000fd5b5050602254602480546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639b19251a925001600060405180830381600087803b158015610c8357600080fd5b505af1158015610c97573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610cf957600080fd5b505af1158015610d0d573d6000803e3d6000fd5b5050602480546026546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240938101939093521692506340c10f199150604401600060405180830381600087803b158015610d7e57600080fd5b505af1158015610d92573d6000803e3d6000fd5b50506025546026546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42406024820152911692506340c10f199150604401600060405180830381600087803b158015610e0157600080fd5b505af1158015610e15573d6000803e3d6000fd5b5050602480546022546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526207a1209381019390935216925063a9059cbb91506044016020604051808303816000875af1158015610e8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eaf919061b552565b506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b158015610f3057600080fd5b505af1158015610f44573d6000803e3d6000fd5b5050604080516060810182526024546001600160a01b039081168252600160208084019182528451908101855260008152938301849052825160298054925167ffffffffffffffff1674010000000000000000000000000000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090931691909316171781559093509150602a90610fe2908261b637565b505050565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc513169100000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b1580156110cc57600080fd5b505af11580156110e0573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506111cb919060040161b6f6565b600060405180830381600087803b1580156111e557600080fd5b505af11580156111f9573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506112569392831692909116908790879060040161b709565b600060405180830381600087803b15801561127057600080fd5b505af1158015611284573d6000803e3d6000fd5b505050505050565b6024805460275460405160009381018490526001600160a01b03928316604482015291166064820152819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602854905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b15801561136e57600080fd5b505af1158015611382573d6000803e3d6000fd5b5050604051630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024016111cb565b6060601680548060200260200160405190810160405280929190818152602001828054801561143557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611417575b5050505050905090565b6022546025546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600092919091169063d936547e90602401602060405180830381865afa1580156114a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114cd919061b552565b90506114da600082616f66565b6022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561154f57600080fd5b505af1158015611563573d6000803e3d6000fd5b50506025546040516001600160a01b0390911692507faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549150600090a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156115f857600080fd5b505af115801561160c573d6000803e3d6000fd5b50506022546025546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639b19251a9150602401600060405180830381600087803b15801561167257600080fd5b505af1158015611686573d6000803e3d6000fd5b50506022546025546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063d936547e9150602401602060405180830381865afa1580156116ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611713919061b552565b9050611720600182616f66565b50565b6060601e805480602002602001604051908101604052809291908181526020016000905b8282101561185c57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156118455783829060005260206000200180546117b89061b5a3565b80601f01602080910402602001604051908101604052809291908181526020018280546117e49061b5a3565b80156118315780601f1061180657610100808354040283529160200191611831565b820191906000526020600020905b81548152906001019060200180831161181457829003601f168201915b505050505081526020019060010190611799565b505050508152505081526020019060010190611747565b50505050905090565b60405163ca669fa760e01b81526101236004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156118b357600080fd5b505af11580156118c7573d6000803e3d6000fd5b50506040805161012360248201527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506119a8919060040161b6f6565b600060405180830381600087803b1580156119c257600080fd5b505af11580156119d6573d6000803e3d6000fd5b50506022546025546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639b19251a91506024015b600060405180830381600087803b158015611a3d57600080fd5b505af1158015611a51573d6000803e3d6000fd5b50505050565b60606018805480602002602001604051908101604052809291908181526020018280548015611435576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611417575050505050905090565b602480546027546040516370a0823160e01b81526001600160a01b039182166004820152620186a09360009392909216916370a082319101602060405180830381865afa158015611b0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b30919061b740565b9050611b3d816000616fe8565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015611b8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bb1919061b740565b6027546040516001600160a01b0390911660248201526044810185905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391611c94916001600160a01b039190911690600090869060040161b759565b600060405180830381600087803b158015611cae57600080fd5b505af1158015611cc2573d6000803e3d6000fd5b50506022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015611d3b57600080fd5b505af1158015611d4f573d6000803e3d6000fd5b50506024546027546040518881526001600160a01b039283169450911691507fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060200160405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611df457600080fd5b505af1158015611e08573d6000803e3d6000fd5b5050602254602754602480546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b03938416600482015290831691810191909152604481018990529116925063d9caed129150606401600060405180830381600087803b158015611e8357600080fd5b505af1158015611e97573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa158015611ee9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f0d919061b740565b9050611f198186616fe8565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015611f69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f8d919061b740565b9050611fa281611f9d888761b7b0565b616fe8565b602480546020546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015611ff2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612016919061b740565b9050612023816000616fe8565b50505050505050565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024016110b2565b60606017805480602002602001604051908101604052809291908181526020018280548015611435576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611417575050505050905090565b6022546040517feab103df000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b039091169063eab103df90602401600060405180830381600087803b1580156121ba57600080fd5b505af11580156121ce573d6000803e3d6000fd5b5050602480546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42409381019390935216925063095ea7b391506044016020604051808303816000875af1158015612244573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612268919061b552565b506025546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240602482015291169063095ea7b3906044016020604051808303816000875af11580156122d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122fd919061b552565b50604080517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152815160058183030181526025820192839052630618f58760e51b9092527f73cba663000000000000000000000000000000000000000000000000000000006029820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906049015b600060405180830381600087803b1580156123a557600080fd5b505af11580156123b9573d6000803e3d6000fd5b505060225460275460405160609190911b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201526001600160a01b03909116925063e609055e915060340160408051601f19818403018152908290526024547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352612463926001600160a01b03909116906103e890879060040161b7c3565b600060405180830381600087803b15801561247d57600080fd5b505af1158015612491573d6000803e3d6000fd5b5050505050565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561250457600080fd5b505af1158015612518573d6000803e3d6000fd5b5050602854602654604051600094508493506001600160a01b0392831692909116906125439061b0c1565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f08015801561257f573d6000803e3d6000fd5b50604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156125ef57600080fd5b505af1158015612603573d6000803e3d6000fd5b50506020546026546040516001600160a01b0392831694506000935091169061262b9061b0c1565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015612667573d6000803e3d6000fd5b50604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156126d757600080fd5b505af11580156126eb573d6000803e3d6000fd5b50506020546028546040516001600160a01b039283169450911691506000906127139061b0c1565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f08015801561274f573d6000803e3d6000fd5b506020546028546026546040519394506001600160a01b039283169391831692169061277a9061b0c1565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610fe2573d6000803e3d6000fd5b6024805460275460405160019381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602854905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561289857600080fd5b505af11580156128ac573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561291c57600080fd5b505af1158015612930573d6000803e3d6000fd5b50506022546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321fc65f293506112569260009216908790879060040161b709565b6060601b805480602002602001604051908101604052809291908181526020016000905b8282101561185c57838290600052602060002090600202016040518060400160405290816000820180546129de9061b5a3565b80601f0160208091040260200160405190810160405280929190818152602001828054612a0a9061b5a3565b8015612a575780601f10612a2c57610100808354040283529160200191612a57565b820191906000526020600020905b815481529060010190602001808311612a3a57829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015612af157602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411612a9e5790505b505050505081525050815260200190600101906129ab565b6024805460275460405160009381018490526001600160a01b03928316604482015291166064820152819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc513169100000000000000000000000000000000000000000000000000000000179052602854905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401611354565b6022546040517feab103df000000000000000000000000000000000000000000000000000000008152600160048201526001600160a01b039091169063eab103df90602401600060405180830381600087803b158015612c3457600080fd5b505af1158015612c48573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b158015612cad57600080fd5b505af1158015612cc1573d6000803e3d6000fd5b5050602480546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42409381019390935216925063095ea7b391506044016020604051808303816000875af1158015612d37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d5b919061b552565b506025546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240602482015291169063095ea7b3906044016020604051808303816000875af1158015612dcc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612df0919061b552565b50604080517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152815160058183030181526025820192839052630618f58760e51b9092527f584a7938000000000000000000000000000000000000000000000000000000006029820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e09060490161238b565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a09060009060250160408051808303601f1901815290829052602480546021546370a0823160e01b85526001600160a01b0390811660048601529294506000939216916370a082319101602060405180830381865afa158015612f15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f39919061b740565b9050612f46816000616fe8565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015612f96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fba919061b740565b6020546040516001600160a01b0390911660248201526044810186905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161309d916001600160a01b039190911690600090869060040161b759565b600060405180830381600087803b1580156130b757600080fd5b505af11580156130cb573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561314457600080fd5b505af1158015613158573d6000803e3d6000fd5b50506020546040517f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e935061319c92506001600160a01b039091169060299061b8d2565b60405180910390a16020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561321957600080fd5b505af115801561322d573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b14369061327b908990899060299061b8f4565b60405180910390a36022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156132f857600080fd5b505af115801561330c573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b9061335a908990899060299061b8f4565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156133bb57600080fd5b505af11580156133cf573d6000803e3d6000fd5b50506022546021546024546040517fc709ab6e0000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c709ab6e945061342f9392831692909116908a908a9060299060040161b91f565b600060405180830381600087803b15801561344957600080fd5b505af115801561345d573d6000803e3d6000fd5b5050602480546021546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a0823191015b602060405180830381865afa1580156134b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134d4919061b740565b90506134e08187616fe8565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015613530573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613554919061b740565b905061356481611f9d898761b7b0565b602480546020546021546040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529082169381019390935260009291169063dd62ed3e90604401602060405180830381865afa1580156135da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135fe919061b740565b905061360b816000616fe8565b602480546020546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa15801561365b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061367f919061b740565b905061368c816000616fe8565b505050505050505050565b6040517f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260019060009060250160408051808303601f190181529082905260285463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561373057600080fd5b505af1158015613744573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156137b457600080fd5b505af11580156137c8573d6000803e3d6000fd5b50506022546024546040517fc709ab6e0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c709ab6e93506112569260009216908790879060299060040161b91f565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561389457600080fd5b505af11580156138a8573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b15801561390d57600080fd5b505af1158015613921573d6000803e3d6000fd5b5050604051630618f58760e51b81527f584a7938000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561399157600080fd5b505af11580156139a5573d6000803e3d6000fd5b5050602254602754602480546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b03938416600482015290831691810191909152600160448201529116925063d9caed129150606401600060405180830381600087803b158015613a2057600080fd5b505af1158015613a34573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611a3d57600080fd5b6060601a805480602002602001604051908101604052809291908181526020016000905b8282101561185c578382906000526020600020018054613ad99061b5a3565b80601f0160208091040260200160405190810160405280929190818152602001828054613b059061b5a3565b8015613b525780601f10613b2757610100808354040283529160200191613b52565b820191906000526020600020905b815481529060010190602001808311613b3557829003601f168201915b505050505081526020019060010190613aba565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015613bd257600080fd5b505af1158015613be6573d6000803e3d6000fd5b50506022546040517f9a590427000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b039091169250639a5904279150602401611a23565b6060601d805480602002602001604051908101604052809291908181526020016000905b8282101561185c5760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015613d1757602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411613cc45790505b50505050508152505081526020019060010190613c58565b602480546027546040516001938101939093526001600160a01b03918216604484015216606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905260285490517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015613e2a57600080fd5b505af1158015613e3e573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b158015613ea357600080fd5b505af1158015613eb7573d6000803e3d6000fd5b5050604051630618f58760e51b81527f584a7938000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015613f2757600080fd5b505af1158015613f3b573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f29450613f99939283169290911690600190879060040161b709565b600060405180830381600087803b158015613fb357600080fd5b505af1158015613fc7573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561247d57600080fd5b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561409557600080fd5b505af11580156140a9573d6000803e3d6000fd5b50506022546040517f9b19251a000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b039091169250639b19251a9150602401611a23565b604080516004808252602480830184526020830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed7016900000000000000000000000000000000000000000000000000000000179052805460275494516370a0823160e01b81526001600160a01b0395861693810193909352620186a0946000939116916370a082319101602060405180830381865afa1580156141a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141c5919061b740565b90506141d2816000616fe8565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015614222573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614246919061b740565b6020546040516001600160a01b0390911660248201526044810186905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391614329916001600160a01b039190911690600090869060040161b759565b600060405180830381600087803b15801561434357600080fd5b505af1158015614357573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156143d057600080fd5b505af11580156143e4573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0935001905060405180910390a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561449d57600080fd5b505af11580156144b1573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5906144fc908990899061b974565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561455d57600080fd5b505af1158015614571573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506145ce9392831692909116908a908a9060040161b709565b600060405180830381600087803b1580156145e857600080fd5b505af11580156145fc573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa15801561464e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614672919061b740565b905061467f816000616fe8565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa1580156146cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906146f3919061b740565b90506135648185616fe8565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a09060009060250160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561479a57600080fd5b505af11580156147ae573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614899919060040161b6f6565b600060405180830381600087803b1580156148b357600080fd5b505af11580156148c7573d6000803e3d6000fd5b50506022546021546024546040517fc709ab6e0000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c709ab6e94506112569392831692909116908790879060299060040161b91f565b6060601c805480602002602001604051908101604052809291908181526020016000905b8282101561185c5760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015614a0a57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116149b75790505b5050505050815250508152602001906001019061494b565b6022546040517feab103df000000000000000000000000000000000000000000000000000000008152600160048201526103e8916001600160a01b03169063eab103df90602401600060405180830381600087803b158015614a8357600080fd5b505af1158015614a97573d6000803e3d6000fd5b5050602480546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42409381019390935216925063095ea7b391506044016020604051808303816000875af1158015614b0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614b31919061b552565b506025546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240602482015291169063095ea7b3906044016020604051808303816000875af1158015614ba2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614bc6919061b552565b50602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015614c17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614c3b919061b740565b6025546028546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa158015614c8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614cb1919061b740565b90506000604051602001614ce8907f68656c6c6f000000000000000000000000000000000000000000000000000000815260050190565b60408051808303601f19018152908290526022546381bad6f360e01b8352600160048401819052602484018190526044840181905260648401526001600160a01b031660848301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015614d6b57600080fd5b505af1158015614d7f573d6000803e3d6000fd5b505060245460275460405160609190911b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201526001600160a01b0390911692507f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae915060340160408051601f1981840301815290829052614e08918890869061b98d565b60405180910390a26022546027546040805160609290921b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208301528051808303601401815260348301918290526024547fe609055e000000000000000000000000000000000000000000000000000000009092526001600160a01b039384169363e609055e93614ea79391909116908990879060380161b7c3565b600060405180830381600087803b158015614ec157600080fd5b505af1158015614ed5573d6000803e3d6000fd5b50505050611a518484614ee8919061b9b8565b602480546022546040516370a0823160e01b81526001600160a01b0391821660048201529116916370a082319101602060405180830381865afa158015614f33573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f9d919061b740565b60606019805480602002602001604051908101604052809291908181526020016000905b8282101561185c578382906000526020600020018054614f9a9061b5a3565b80601f0160208091040260200160405190810160405280929190818152602001828054614fc69061b5a3565b80156150135780601f10614fe857610100808354040283529160200191615013565b820191906000526020600020905b815481529060010190602001808311614ff657829003601f168201915b505050505081526020019060010190614f7b565b60085460009060ff161561503f575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa1580156150d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906150f4919061b740565b1415905090565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561515457600080fd5b505af1158015615168573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250615253919060040161b6f6565b600060405180830381600087803b15801561526d57600080fd5b505af1158015615281573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156152d557600080fd5b505af11580156152e9573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561534657600080fd5b505af115801561535a573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250615445919060040161b6f6565b600060405180830381600087803b15801561545f57600080fd5b505af1158015615473573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156154c757600080fd5b505af11580156154db573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561553857600080fd5b505af115801561554c573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156155a057600080fd5b505af11580156155b4573d6000803e3d6000fd5b505060248054602754604051620186a09381018490526001600160a01b039283166044820152911660648201529092506000915060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905251630618f58760e51b81527fd93c0665000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156156b057600080fd5b505af11580156156c4573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561572157600080fd5b505af1158015615735573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506157929392831692909116908790879060040161b709565b600060405180830381600087803b1580156157ac57600080fd5b505af11580156157c0573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561581d57600080fd5b505af1158015615831573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561588557600080fd5b505af1158015615899573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a0823191015b602060405180830381865afa1580156158ec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615910919061b740565b905061591d816000616fe8565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa15801561596d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615991919061b740565b6020546040516001600160a01b0390911660248201526044810186905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391615a74916001600160a01b039190911690600090869060040161b759565b600060405180830381600087803b158015615a8e57600080fd5b505af1158015615aa2573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015615b1b57600080fd5b505af1158015615b2f573d6000803e3d6000fd5b505060208054602454602754604080516001600160a01b0394851681529485018c905291831684830152919091166060830152517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609350908190036080019150a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015615c0557600080fd5b505af1158015615c19573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d590615c64908990899061b974565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015615cc557600080fd5b505af1158015615cd9573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f29450615d369392831692909116908a908a9060040161b709565b600060405180830381600087803b158015615d5057600080fd5b505af1158015615d64573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101613493565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a2000000000000000000000000000000000000000000000000000000001790526024805460275492516370a0823160e01b81526001600160a01b0393841660048201529394506000939216916370a0823191016158cf565b60606015805480602002602001604051908101604052809291908181526020018280548015611435576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611417575050505050905090565b600080604051602001615ef5907f68656c6c6f000000000000000000000000000000000000000000000000000000815260050190565b60408051808303601f190181529082905260285463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015615f5c57600080fd5b505af1158015615f70573d6000803e3d6000fd5b5050604051630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401614899565b60405163ca669fa760e01b81526101236004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561601957600080fd5b505af115801561602d573d6000803e3d6000fd5b50506040805161012360248201527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061610e919060040161b6f6565b600060405180830381600087803b15801561612857600080fd5b505af115801561613c573d6000803e3d6000fd5b50506022546025546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a5904279150602401611a23565b602480546027546040516001938101939093526001600160a01b03918216604484015216606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905260285490517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561628857600080fd5b505af115801561629c573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b15801561630157600080fd5b505af1158015616315573d6000803e3d6000fd5b5050604051630618f58760e51b81527f584a7938000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561638557600080fd5b505af1158015616399573d6000803e3d6000fd5b50506022546021546024546040517fc709ab6e0000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c709ab6e9450613f99939283169290911690600190879060299060040161b91f565b602254602480546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600093919091169163d936547e9101602060405180830381865afa158015616463573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616487919061b552565b9050616494600182616f66565b6022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561650957600080fd5b505af115801561651d573d6000803e3d6000fd5b50506024546040516001600160a01b0390911692507f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919150600090a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156165b257600080fd5b505af11580156165c6573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b15801561662b57600080fd5b505af115801561663f573d6000803e3d6000fd5b5050602254602480546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529216935063d936547e925001602060405180830381865afa1580156166a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906166cb919061b552565b9050611720600082616f66565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc5131691000000000000000000000000000000000000000000000000000000001790526024805460275492516370a0823160e01b81526001600160a01b0393841660048201529394506000939216916370a082319101602060405180830381865afa1580156167b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906167d5919061b740565b90506167e2816000616fe8565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015616832573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616856919061b740565b6020546040516001600160a01b0390911660248201526044810186905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391616939916001600160a01b039190911690600090869060040161b759565b600060405180830381600087803b15801561695357600080fd5b505af1158015616967573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156169e057600080fd5b505af11580156169f4573d6000803e3d6000fd5b50506020547f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af6092506001600160a01b03169050616a3260028861b9cb565b602454602754604080516001600160a01b03958616815260208101949094529184168383015292909216606082015290519081900360800190a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015616ae157600080fd5b505af1158015616af5573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d590616b40908990899061b974565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015616ba157600080fd5b505af1158015616bb5573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f29450616c129392831692909116908a908a9060040161b709565b600060405180830381600087803b158015616c2c57600080fd5b505af1158015616c40573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa158015616c92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616cb6919061b740565b9050616cc781611f9d60028961b9cb565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015616d17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616d3b919061b740565b905061356481616d4c60028a61b9cb565b611f9d908761b7b0565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015616db457600080fd5b505af1158015616dc8573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250616eb3919060040161b6f6565b600060405180830381600087803b158015616ecd57600080fd5b505af1158015616ee1573d6000803e3d6000fd5b5050602254602754602480546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b03938416600482015290831691810191909152604481018690529116925063d9caed129150606401612463565b6000616f5161b0eb565b616f5c848483617040565b9150505b92915050565b6040517ff7fe347700000000000000000000000000000000000000000000000000000000815282151560048201528115156024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f7fe3477906044015b60006040518083038186803b158015616fd457600080fd5b505afa158015611284573d6000803e3d6000fd5b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c5490604401616fbc565b60008061704d85846170bb565b90506170b06040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f7879000000815250828660405160200161709b92919061ba06565b604051602081830303815290604052856170c7565b9150505b9392505050565b60006170b483836170f5565b60c081015151600090156170eb576170e484848460c00151617110565b90506170b4565b6170e484846172b6565b600061710183836173a1565b6170b4838360200151846170c7565b60008061711b6173b1565b905060006171298683617484565b90506000617140826060015183602001518561792a565b9050600061715083838989617b3c565b9050600061715d826189b9565b602081015181519192509060030b156171d05789826040015160405160200161718792919061ba28565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526171c79160040161b6f6565b60405180910390fd5b60006172136040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001618b88565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d9061726690849060040161b6f6565b602060405180830381865afa158015617283573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906172a7919061baa9565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc9259061730b90879060040161b6f6565b600060405180830381865afa158015617328573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617350919081019061bb8b565b9050600061737e828560405160200161736a92919061bbc0565b604051602081830303815290604052618d88565b90506001600160a01b038116616f5c57848460405160200161718792919061bbef565b6173ad82826000618d9b565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c9061743890849060040161bc9a565b600060405180830381865afa158015617455573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261747d919081019061bce1565b9250505090565b6174b66040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d90506175016040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b61750a85618e9e565b6020820152600061751a86619283565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa15801561755c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617584919081019061bce1565b8683856020015160405160200161759e949392919061bd2a565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb11906175f690859060040161b6f6565b600060405180830381865afa158015617613573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261763b919081019061bce1565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f69061768390849060040161be2e565b602060405180830381865afa1580156176a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906176c4919061b552565b6176d95781604051602001617187919061be80565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061771e90849060040161bf12565b600060405180830381865afa15801561773b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617763919081019061bce1565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906177aa90849060040161bf64565b602060405180830381865afa1580156177c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906177eb919061b552565b15617880576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061783590849060040161bf64565b600060405180830381865afa158015617852573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261787a919081019061bce1565b60408501525b846001600160a01b03166349c4fac88286600001516040516020016178a5919061bfb6565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016178d192919061c022565b600060405180830381865afa1580156178ee573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617916919081019061bce1565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b60608152602001906001900390816179465790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250816000815181106179a6576179a661c047565b60200260200101819052506040518060400160405280600381526020017f2d726c0000000000000000000000000000000000000000000000000000000000815250816001815181106179fa576179fa61c047565b602002602001018190525084604051602001617a16919061c076565b60405160208183030381529060405281600281518110617a3857617a3861c047565b602002602001018190525082604051602001617a54919061c0e2565b60405160208183030381529060405281600381518110617a7657617a7661c047565b60200260200101819052506000617a8c826189b9565b602080820151604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000008185019081528251808401845260008082529086015282518084019093529051825292810192909252919250617b1d9060408051808201825260008082526020918201528151808301909252845182528085019082015290619506565b617b325785604051602001617187919061c123565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015617b8c565b511590565b617d0057826020015115617c48576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a4016171c7565b8260c0015115617d00576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a4016171c7565b6040805160ff8082526120008201909252600091816020015b6060815260200190600190039081617d1957905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280617d749061c1b4565b935060ff1681518110617d8957617d8961c047565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e3700000000000000000000000000000000000000815250604051602001617dda919061c1d3565b604051602081830303815290604052828280617df59061c1b4565b935060ff1681518110617e0a57617e0a61c047565b60200260200101819052506040518060400160405280600681526020017f6465706c6f790000000000000000000000000000000000000000000000000000815250828280617e579061c1b4565b935060ff1681518110617e6c57617e6c61c047565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d65000000000000000000000000000000000000815250828280617eb99061c1b4565b935060ff1681518110617ece57617ece61c047565b60200260200101819052508760200151828280617eea9061c1b4565b935060ff1681518110617eff57617eff61c047565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e747261637450617468000000000000000000000000000000000000815250828280617f4c9061c1b4565b935060ff1681518110617f6157617f6161c047565b602090810291909101015287518282617f798161c1b4565b935060ff1681518110617f8e57617f8e61c047565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e49640000000000000000000000000000000000000000000000815250828280617fdb9061c1b4565b935060ff1681518110617ff057617ff061c047565b602002602001018190525061800446619567565b828261800f8161c1b4565b935060ff16815181106180245761802461c047565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806180719061c1b4565b935060ff16815181106180865761808661c047565b60200260200101819052508682828061809e9061c1b4565b935060ff16815181106180b3576180b361c047565b60209081029190910101528551156181da5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826181048161c1b4565b935060ff16815181106181195761811961c047565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d9061816990899060040161b6f6565b600060405180830381865afa158015618186573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526181ae919081019061bce1565b82826181b98161c1b4565b935060ff16815181106181ce576181ce61c047565b60200260200101819052505b8460200151156182aa5760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826182238161c1b4565b935060ff16815181106182385761823861c047565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806182859061c1b4565b935060ff168151811061829a5761829a61c047565b6020026020010181905250618471565b6182e2617b878660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6183755760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826183258161c1b4565b935060ff168151811061833a5761833a61c047565b60200260200101819052508460a0015160405160200161835a919061c076565b6040516020818303038152906040528282806182859061c1b4565b8460c001511580156183b85750604080890151815180830183526000808252602091820152825180840190935281518352908101908201526183b690511590565b155b156184715760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826183fc8161c1b4565b935060ff16815181106184115761841161c047565b602002602001018190525061842588619607565b604051602001618435919061c076565b6040516020818303038152906040528282806184509061c1b4565b935060ff16815181106184655761846561c047565b60200260200101819052505b604080860151815180830183526000808252602091820152825180840190935281518352908101908201526184a590511590565b61853a5760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826184e88161c1b4565b935060ff16815181106184fd576184fd61c047565b602002602001018190525084604001518282806185199061c1b4565b935060ff168151811061852e5761852e61c047565b60200260200101819052505b60608501511561865b5760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826185838161c1b4565b935060ff16815181106185985761859861c047565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015618607573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261862f919081019061bce1565b828261863a8161c1b4565b935060ff168151811061864f5761864f61c047565b60200260200101819052505b60e085015151156187025760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826186a58161c1b4565b935060ff16815181106186ba576186ba61c047565b60200260200101819052506186d68560e0015160000151619567565b82826186e18161c1b4565b935060ff16815181106186f6576186f661c047565b60200260200101819052505b60e085015160200151156187ac5760408051808201909152600a81527f2d2d6761735072696365000000000000000000000000000000000000000000006020820152828261874f8161c1b4565b935060ff16815181106187645761876461c047565b60200260200101819052506187808560e0015160200151619567565b828261878b8161c1b4565b935060ff16815181106187a0576187a061c047565b60200260200101819052505b60e085015160400151156188565760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826187f98161c1b4565b935060ff168151811061880e5761880e61c047565b602002602001018190525061882a8560e0015160400151619567565b82826188358161c1b4565b935060ff168151811061884a5761884a61c047565b60200260200101819052505b60e085015160600151156189005760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826188a38161c1b4565b935060ff16815181106188b8576188b861c047565b60200260200101819052506188d48560e0015160600151619567565b82826188df8161c1b4565b935060ff16815181106188f4576188f461c047565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561891e5761891e61b574565b60405190808252806020026020018201604052801561895157816020015b606081526020019060019003908161893c5790505b50905060005b8260ff168160ff1610156189aa57838160ff168151811061897a5761897a61c047565b6020026020010151828260ff16815181106189975761899761c047565b6020908102919091010152600101618957565b5093505050505b949350505050565b6189e06040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c91618a669186910161c23e565b600060405180830381865afa158015618a83573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618aab919081019061bce1565b90506000618ab9868361a0f6565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b8152600401618ae9919061b444565b6000604051808303816000875af1158015618b08573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618b30919081019061c285565b805190915060030b15801590618b495750602081015151155b8015618b585750604081015151155b15617b325781600081518110618b7057618b7061c047565b6020026020010151604051602001617187919061c33b565b60606000618bbd8560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528651825280870190820152909150618bf49082905b9061a24b565b15618d51576000618c7182618c6b84618c65618c378a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061a272565b9061a2d4565b604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150618cd590829061a24b565b15618d3f57604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618d3c905b829061a359565b90505b618d488161a37f565b925050506170b4565b8215618d6a57848460405160200161718792919061c527565b50506040805160208101909152600081526170b4565b509392505050565b6000808251602084016000f09392505050565b8160a0015115618daa57505050565b6000618db784848461a3e8565b90506000618dc4826189b9565b602081015181519192509060030b158015618e605750604080518082018252600781527f535543434553530000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618e6090604080518082018252600080825260209182015281518083019092528451825280850190820152618bee565b15618e6d57505050505050565b60408201515115618e8d578160400151604051602001617187919061c5ce565b80604051602001617187919061c62c565b60606000618ed38360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150618f38905b8290619506565b15618fa757604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526170b490618fa290839061a983565b61a37f565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619009905b829061aa0d565b6001036190d657604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261906f90618d35565b50604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526170b490618fa2905b839061a359565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261913590618f31565b1561926c57604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061919d90839061aaa7565b9050600081600183516191b0919061b7b0565b815181106191c0576191c061c047565b60200260200101519050619263618fa26192366040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925285518252808601908201529061a983565b95945050505050565b82604051602001617187919061c697565b50919050565b606060006192b88360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061931a90618f31565b15619328576170b48161a37f565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261938790619002565b6001036193f157604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526170b490618fa2906190cf565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261945090618f31565b1561926c57604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201819052845180860190955292518452830152906194b890839061aaa7565b90506001815111156194f45780600282516194d3919061b7b0565b815181106194e3576194e361c047565b602002602001015192505050919050565b5082604051602001617187919061c697565b80518251600091111561951b57506000616f60565b815183516020850151600092916195319161b9b8565b61953b919061b7b0565b905082602001518103619552576001915050616f60565b82516020840151819020912014905092915050565b606060006195748361ab4c565b600101905060008167ffffffffffffffff8111156195945761959461b574565b6040519080825280601f01601f1916602001820160405280156195be576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846195c857509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e5345440000000000000000000000000000000000000000000081840190815285518087018752838152840192909252845180860190955251845290830152606091619693905b829061ac2e565b156196d357505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e73650000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526197329061968c565b1561977257505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d49540000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526197d19061968c565b1561981157505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526198709061968c565b806198d55750604080518082018252601081527f47504c2d322e302d6f722d6c6174657200000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526198d59061968c565b1561991557505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526199749061968c565b806199d95750604080518082018252601081527f47504c2d332e302d6f722d6c6174657200000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526199d99061968c565b15619a1957505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619a789061968c565b80619add5750604080518082018252601181527f4c47504c2d322e312d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619add9061968c565b15619b1d57505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619b7c9061968c565b80619be15750604080518082018252601181527f4c47504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619be19061968c565b15619c2157505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619c809061968c565b15619cc057505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619d1f9061968c565b15619d5f57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619dbe9061968c565b15619dfe57505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619e5d9061968c565b15619e9d57505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e300000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619efc9061968c565b15619f3c57505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619f9b9061968c565b8061a0005750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a0009061968c565b1561a04057505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a09f9061968c565b1561a0df57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151617187929060200161c775565b60608060005b845181101561a181578185828151811061a1185761a11861c047565b602002602001015160405160200161a13192919061bbc0565b60405160208183030381529060405291506001855161a150919061b7b0565b811461a179578160405160200161a167919061c8de565b60405160208183030381529060405291505b60010161a0fc565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161a19a579050509050838160008151811061a1c55761a1c561c047565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061a2195761a21961c047565b6020026020010181905250818160028151811061a2385761a23861c047565b6020908102919091010152949350505050565b602080830151835183519284015160009361a269929184919061ac42565b14159392505050565b6040805180820190915260008082526020820152600061a2a4846000015185602001518560000151866020015161ad53565b905083602001518161a2b6919061b7b0565b8451859061a2c590839061b7b0565b90525060208401525090919050565b604080518082019091526000808252602082015281518351101561a2f9575081616f60565b602080830151908401516001911461a3205750815160208481015190840151829020919020145b801561a3515782518451859061a33790839061b7b0565b905250825160208501805161a34d90839061b9b8565b9052505b509192915050565b604080518082019091526000808252602082015261a37883838361ae73565b5092915050565b60606000826000015167ffffffffffffffff81111561a3a05761a3a061b574565b6040519080825280601f01601f19166020018201604052801561a3ca576020820181803683370190505b509050600060208201905061a378818560200151866000015161af1e565b6060600061a3f46173b1565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161a41157905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061a46c9061c1b4565b935060ff168151811061a4815761a48161c047565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161a4d2919061c91f565b60405160208183030381529060405282828061a4ed9061c1b4565b935060ff168151811061a5025761a50261c047565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061a54f9061c1b4565b935060ff168151811061a5645761a56461c047565b60200260200101819052508260405160200161a580919061c0e2565b60405160208183030381529060405282828061a59b9061c1b4565b935060ff168151811061a5b05761a5b061c047565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061a5fd9061c1b4565b935060ff168151811061a6125761a61261c047565b602002602001018190525061a627878461af98565b828261a6328161c1b4565b935060ff168151811061a6475761a64761c047565b60209081029190910101528551511561a6f35760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261a6998161c1b4565b935060ff168151811061a6ae5761a6ae61c047565b602002602001018190525061a6c786600001518461af98565b828261a6d28161c1b4565b935060ff168151811061a6e75761a6e761c047565b60200260200101819052505b85608001511561a7615760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261a73c8161c1b4565b935060ff168151811061a7515761a75161c047565b602002602001018190525061a7c7565b841561a7c75760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261a7a68161c1b4565b935060ff168151811061a7bb5761a7bb61c047565b60200260200101819052505b6040860151511561a8635760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261a8118161c1b4565b935060ff168151811061a8265761a82661c047565b6020026020010181905250856040015182828061a8429061c1b4565b935060ff168151811061a8575761a85761c047565b60200260200101819052505b85606001511561a8cd5760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261a8ac8161c1b4565b935060ff168151811061a8c15761a8c161c047565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561a8eb5761a8eb61b574565b60405190808252806020026020018201604052801561a91e57816020015b606081526020019060019003908161a9095790505b50905060005b8260ff168160ff16101561a97757838160ff168151811061a9475761a94761c047565b6020026020010151828260ff168151811061a9645761a96461c047565b602090810291909101015260010161a924565b50979650505050505050565b604080518082019091526000808252602082015281518351101561a9a8575081616f60565b8151835160208501516000929161a9be9161b9b8565b61a9c8919061b7b0565b6020840151909150600190821461a9e9575082516020840151819020908220145b801561aa045783518551869061aa0090839061b7b0565b9052505b50929392505050565b600080826000015161aa31856000015186602001518660000151876020015161ad53565b61aa3b919061b9b8565b90505b8351602085015161aa4f919061b9b8565b811161a378578161aa5f8161c964565b925050826000015161aa9685602001518361aa7a919061b7b0565b865161aa86919061b7b0565b838660000151876020015161ad53565b61aaa0919061b9b8565b905061aa3e565b6060600061aab5848461aa0d565b61aac090600161b9b8565b67ffffffffffffffff81111561aad85761aad861b574565b60405190808252806020026020018201604052801561ab0b57816020015b606081526020019060019003908161aaf65790505b50905060005b8151811015618d805761ab27618fa2868661a359565b82828151811061ab395761ab3961c047565b602090810291909101015260010161ab11565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061ab95577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061abc1576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061abdf57662386f26fc10000830492506010015b6305f5e100831061abf7576305f5e100830492506008015b612710831061ac0b57612710830492506004015b6064831061ac1d576064830492506002015b600a8310616f605760010192915050565b600061ac3a838361afd8565b159392505050565b60008085841161ad49576020841161acf5576000841561ac8d57600161ac6986602061b7b0565b61ac7490600861c97e565b61ac7f90600261ca7c565b61ac89919061b7b0565b1990505b835181168561ac9c898961b9b8565b61aca6919061b7b0565b805190935082165b81811461ace05787841161acc857879450505050506189b1565b8361acd28161ca88565b94505082845116905061acae565b61acea878561b9b8565b9450505050506189b1565b83832061ad02858861b7b0565b61ad0c908761b9b8565b91505b85821061ad475784822080820361ad345761ad2a868461b9b8565b93505050506189b1565b61ad3f60018461b7b0565b92505061ad0f565b505b5092949350505050565b6000838186851161ae5e576020851161ae0d576000851561ad9f57600161ad7b87602061b7b0565b61ad8690600861c97e565b61ad9190600261ca7c565b61ad9b919061b7b0565b1990505b8451811660008761adb08b8b61b9b8565b61adba919061b7b0565b855190915083165b82811461adff5781861061ade75761adda8b8b61b9b8565b96505050505050506189b1565b8561adf18161c964565b96505083865116905061adc2565b8596505050505050506189b1565b508383206000905b61ae1f868961b7b0565b821161ae5c5785832080820361ae3b57839450505050506189b1565b61ae4660018561b9b8565b935050818061ae549061c964565b92505061ae15565b505b61ae68878761b9b8565b979650505050505050565b6040805180820190915260008082526020820152600061aea5856000015186602001518660000151876020015161ad53565b60208087018051918601919091525190915061aec1908261b7b0565b83528451602086015161aed4919061b9b8565b810361aee3576000855261af15565b8351835161aef1919061b9b8565b8551869061af0090839061b7b0565b905250835161af0f908261b9b8565b60208601525b50909392505050565b6020811061af56578151835261af3560208461b9b8565b925061af4260208361b9b8565b915061af4f60208261b7b0565b905061af1e565b600019811561af8557600161af6c83602061b7b0565b61af789061010061ca7c565b61af82919061b7b0565b90505b9151835183169219169190911790915250565b6060600061afa68484617484565b805160208083015160405193945061afc09390910161ca9f565b60405160208183030381529060405291505092915050565b815181516000919081111561afeb575081515b6020808501519084015160005b8381101561b0a4578251825180821461b07457600019602087101561b0535760018461b02589602061b7b0565b61b02f919061b9b8565b61b03a90600861c97e565b61b04590600261ca7c565b61b04f919061b7b0565b1990505b818116838216818103911461b071579750616f609650505050505050565b50505b61b07f60208661b9b8565b945061b08c60208561b9b8565b9350505060208161b09d919061b9b8565b905061aff8565b5084518651617b32919061caf7565b610c9f806200cb1883390190565b611eb9806200d7b783390190565b611783806200f67083390190565b610de08062010df383390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161b12e61b133565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161b12e6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561b1e55783516001600160a01b031683526020938401939092019160010161b1be565b509095945050505050565b60005b8381101561b20b57818101518382015260200161b1f3565b50506000910152565b6000815180845261b22c81602086016020860161b1f0565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b33c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561b322577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261b30c84865161b214565b602095860195909450929092019160010161b2d2565b50919750505060209485019492909201915060010161b268565b50929695505050505050565b600081518084526020840193506020830160005b8281101561b39c5781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161b35c565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b33c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261b412604088018261b214565b905060208201519150868103602088015261b42d818361b348565b96505050602093840193919091019060010161b3ce565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b33c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261b4a685835161b214565b9450602093840193919091019060010161b46c565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b33c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261b53c604087018261b348565b955050602093840193919091019060010161b4e3565b60006020828403121561b56457600080fd5b815180151581146170b457600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c9082168061b5b757607f821691505b60208210810361927d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f821115610fe257806000526020600020601f840160051c8101602085101561b6175750805b601f840160051c820191505b81811015612491576000815560010161b623565b815167ffffffffffffffff81111561b6515761b65161b574565b61b6658161b65f845461b5a3565b8461b5f0565b6020601f82116001811461b699576000831561b6815750848201515b600019600385901b1c1916600184901b178455612491565b600084815260208120601f198516915b8281101561b6c9578785015182556020948501946001909201910161b6a9565b508482101561b6e75786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6020815260006170b4602083018461b214565b6001600160a01b03851681526001600160a01b0384166020820152826040820152608060608201526000617b32608083018461b214565b60006020828403121561b75257600080fd5b5051919050565b6001600160a01b0384168152826020820152606060408201526000619263606083018461b214565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810381811115616f6057616f6061b781565b60808152600061b7d6608083018761b214565b6001600160a01b0386166020840152846040840152828103606084015261ae68818561b214565b600081546001600160a01b038116845267ffffffffffffffff8160a01c1660208501525060018201606060408501526000815461b8398161b5a3565b806060880152600182166000811461b858576001811461b8925761b8c6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083166080890152608082151560051b890101935061b8c6565b84600052602060002060005b8381101561b8bd5781548a82016080015260019091019060200161b89e565b89016080019450505b50919695505050505050565b6001600160a01b03831681526040602082015260006189b1604083018461b7fd565b83815260606020820152600061b90d606083018561b214565b8281036040840152617b32818561b7fd565b6001600160a01b03861681526001600160a01b038516602082015283604082015260a06060820152600061b95660a083018561b214565b828103608084015261b968818561b7fd565b98975050505050505050565b8281526040602082015260006189b1604083018461b214565b60608152600061b9a0606083018661b214565b8460208401528281036040840152617b32818561b214565b80820180821115616f6057616f6061b781565b60008261ba01577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6001600160a01b03831681526040602082015260006189b1604083018461b214565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161ba6081601a85016020880161b1f0565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161ba9d81601c84016020880161b1f0565b01601c01949350505050565b60006020828403121561babb57600080fd5b81516001600160a01b03811681146170b457600080fd5b6040516060810167ffffffffffffffff8111828210171561baf55761baf561b574565b60405290565b60008067ffffffffffffffff84111561bb165761bb1661b574565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561bb455761bb4561b574565b60405283815290508082840185101561bb5d57600080fd5b618d8084602083018561b1f0565b600082601f83011261bb7c57600080fd5b6170b48383516020850161bafb565b60006020828403121561bb9d57600080fd5b815167ffffffffffffffff81111561bbb457600080fd5b616f5c8482850161bb6b565b6000835161bbd281846020880161b1f0565b83519083019061bbe681836020880161b1f0565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161bc2781601a85016020880161b1f0565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161bc6481603384016020880161b1f0565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f555400000000000000000000000000000000000000000060608201526080602082015260006170b4608083018461b214565b60006020828403121561bcf357600080fd5b815167ffffffffffffffff81111561bd0a57600080fd5b8201601f8101841361bd1b57600080fd5b616f5c8482516020840161bafb565b6000855161bd3c818460208a0161b1f0565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161bd76816001840160208a0161b1f0565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161bdb481600284016020890161b1f0565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161bdf681600284016020880161b1f0565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061be41604083018461b214565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161beb881601f85016020870161b1f0565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061bf25604083018461b214565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061bf77604083018461b214565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161bfee81601485016020870161b1f0565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061c035604083018561b214565b82810360208401526170b0818561b214565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161c0ae81600185016020870161b1f0565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161c0f481846020870161b1f0565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161c1a781604b85016020870161b1f0565b91909101604b0192915050565b600060ff821660ff810361c1ca5761c1ca61b781565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161c23181602985016020870161b1f0565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f504154480000000000000000000060608201526080602082015260006170b4608083018461b214565b60006020828403121561c29757600080fd5b815167ffffffffffffffff81111561c2ae57600080fd5b82016060818503121561c2c057600080fd5b61c2c861bad2565b81518060030b811461c2d957600080fd5b8152602082015167ffffffffffffffff81111561c2f557600080fd5b61c3018682850161bb6b565b602083015250604082015167ffffffffffffffff81111561c32157600080fd5b61c32d8682850161bb6b565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161c39981602185016020870161b1f0565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161c58581602185016020880161b1f0565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161c5c281602e84016020880161b1f0565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161c23181602985016020870161b1f0565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161c68a81602285016020870161b1f0565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161c6cf81600e85016020870161b1f0565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161c7ad81601885016020880161b1f0565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161c7ea81601c84016020880161b1f0565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161c8f081846020870161b1f0565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161c95781601c85016020870161b1f0565b91909101601c0192915050565b6000600019820361c9775761c97761b781565b5060010190565b8082028115828204841417616f6057616f6061b781565b6001815b600184111561c9d05780850481111561c9b45761c9b461b781565b600184161561c9c257908102905b60019390931c92800261c999565b935093915050565b60008261c9e757506001616f60565b8161c9f457506000616f60565b816001811461ca0a576002811461ca145761ca30565b6001915050616f60565b60ff84111561ca255761ca2561b781565b50506001821b616f60565b5060208310610133831016604e8410600b841016171561ca53575081810a616f60565b61ca60600019848461c995565b806000190482111561ca745761ca7461b781565b029392505050565b60006170b4838361c9d8565b60008161ca975761ca9761b781565b506000190190565b6000835161cab181846020880161b1f0565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161caeb81600184016020880161b1f0565b01600101949350505050565b818103600083128015838313168383128216171561a3785761a37861b78156fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220276992b147ed0376af9183cadceed8aaeee816ba7f9608e8ca0386208ed8f31264736f6c634300081a003360a060405234801561001057600080fd5b50604051611eb9380380611eb983398101604081905261002f916101fd565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03838116608052600480546001600160a01b0319169184169190911790556100bb60008261014c565b506100e67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261014c565b506101117f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361014c565b5061012a600080516020611e998339815191528261014c565b50610143600080516020611e998339815191528361014c565b50505050610240565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166101d75760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016101db565b5060005b92915050565b80516001600160a01b03811681146101f857600080fd5b919050565b60008060006060848603121561021257600080fd5b61021b846101e1565b9250610229602085016101e1565b9150610237604085016101e1565b90509250925092565b608051611c22610277600039600081816101ca01528181610597015281816105f901528181610a280152610a8a0152611c226000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c806385f438c1116100e3578063d547741f1161008c578063e609055e11610066578063e609055e146103fc578063e63ab1e91461040f578063eab103df1461043657600080fd5b8063d547741f146103b3578063d936547e146103c6578063d9caed12146103e957600080fd5b80639b19251a116100bd5780639b19251a14610385578063a217fddf14610398578063c709ab6e146103a057600080fd5b806385f438c11461030557806391d148541461032c5780639a5904271461037257600080fd5b806336568abe116101455780635b1125911161011f5780635b112591146102d25780635c975abb146102f25780638456cb59146102fd57600080fd5b806336568abe146102905780633f4ba83a146102a3578063570618e1146102ab57600080fd5b8063248a9ca311610176578063248a9ca314610226578063252f07bf146102585780632f2ff15d1461027d57600080fd5b806301ffc9a71461019d578063116191b6146101c557806321fc65f214610211575b600080fd5b6101b06101ab3660046115ca565b610449565b60405190151581526020015b60405180910390f35b6101ec7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101bc565b61022461021f366004611677565b6104e2565b005b61024a6102343660046116ea565b6000908152600160208190526040909120015490565b6040519081526020016101bc565b6004546101b09074010000000000000000000000000000000000000000900460ff1681565b61022461028b366004611703565b6106e3565b61022461029e366004611703565b61070f565b61022461076d565b61024a7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b6004546101ec9073ffffffffffffffffffffffffffffffffffffffff1681565b60025460ff166101b0565b6102246107a2565b61024a7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101b061033a366004611703565b600091825260016020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610224610380366004611733565b6107d4565b610224610393366004611733565b6108a2565b61024a600081565b6102246103ae366004611750565b610973565b6102246103c1366004611703565b610b79565b6101b06103d4366004611733565b60036020526000908152604090205460ff1681565b6102246103f73660046117f3565b610b9f565b61022461040a366004611834565b610ccb565b61024a7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6102246104443660046118d3565b610f2b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104dc57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104ea610f81565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461051481610fc4565b61051c610fce565b73ffffffffffffffffffffffffffffffffffffffff851660009081526003602052604090205460ff1661057b576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105bc73ffffffffffffffffffffffffffffffffffffffff86167f00000000000000000000000000000000000000000000000000000000000000008661100d565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906106369088908a90899089908990600401611939565b600060405180830381600087803b15801561065057600080fd5b505af1158015610664573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d58686866040516106c993929190611996565b60405180910390a3506106dc6001600055565b5050505050565b600082815260016020819052604090912001546106ff81610fc4565b610709838361108e565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331461075e576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610768828261113b565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61079781610fc4565b61079f6111dc565b50565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107cc81610fc4565b61079f61123b565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6107fe81610fc4565b73ffffffffffffffffffffffffffffffffffffffff821661084b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6108cc81610fc4565b73ffffffffffffffffffffffffffffffffffffffff8216610919576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b61097b610f81565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46109a581610fc4565b6109ad610fce565b73ffffffffffffffffffffffffffffffffffffffff861660009081526003602052604090205460ff16610a0c576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a4d73ffffffffffffffffffffffffffffffffffffffff87167f00000000000000000000000000000000000000000000000000000000000000008761100d565b6040517fd0b492c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610ac99089908b908a908a908a908a90600401611a77565b600060405180830381600087803b158015610ae357600080fd5b505af1158015610af7573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b87878787604051610b5e9493929190611ae8565b60405180910390a350610b716001600055565b505050505050565b60008281526001602081905260409091200154610b9581610fc4565b610709838361113b565b610ba7610f81565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610bd181610fc4565b610bd9610fce565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604090205460ff16610c38576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c5973ffffffffffffffffffffffffffffffffffffffff8416858461100d565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610cb891815260200190565b60405180910390a3506107686001600055565b610cd3610f81565b610cdb610fce565b60045474010000000000000000000000000000000000000000900460ff16610d2f576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841660009081526003602052604090205460ff16610d8e576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015610dfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1f9190611b14565b9050610e4373ffffffffffffffffffffffffffffffffffffffff8616333087611278565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa158015610ed7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610efb9190611b14565b610f059190611b2d565b8787604051610f18959493929190611b67565b60405180910390a250610b716001600055565b6000610f3681610fc4565b506004805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b600260005403610fbd576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61079f81336112be565b60025460ff161561100b576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60405173ffffffffffffffffffffffffffffffffffffffff83811660248301526044820183905261076891859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061134f565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1661113357600083815260016020818152604080842073ffffffffffffffffffffffffffffffffffffffff87168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016104dc565b5060006104dc565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff161561113357600083815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104dc565b6111e46113e5565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b611243610fce565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586112113390565b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526107099186918216906323b872dd90608401611047565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661134b576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044015b60405180910390fd5b5050565b600061137173ffffffffffffffffffffffffffffffffffffffff841683611421565b905080516000141580156113965750808060200190518101906113949190611ba0565b155b15610768576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401611342565b60025460ff1661100b576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606061142f83836000611436565b9392505050565b606081471015611474576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611342565b6000808573ffffffffffffffffffffffffffffffffffffffff16848660405161149d9190611bbd565b60006040518083038185875af1925050503d80600081146114da576040519150601f19603f3d011682016040523d82523d6000602084013e6114df565b606091505b50915091506114ef8683836114f9565b9695505050505050565b60608261150e5761150982611588565b61142f565b8151158015611532575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611581576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401611342565b508061142f565b8051156115985780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156115dc57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461142f57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461079f57600080fd5b60008083601f84011261164057600080fd5b50813567ffffffffffffffff81111561165857600080fd5b60208301915083602082850101111561167057600080fd5b9250929050565b60008060008060006080868803121561168f57600080fd5b853561169a8161160c565b945060208601356116aa8161160c565b935060408601359250606086013567ffffffffffffffff8111156116cd57600080fd5b6116d98882890161162e565b969995985093965092949392505050565b6000602082840312156116fc57600080fd5b5035919050565b6000806040838503121561171657600080fd5b8235915060208301356117288161160c565b809150509250929050565b60006020828403121561174557600080fd5b813561142f8161160c565b60008060008060008060a0878903121561176957600080fd5b86356117748161160c565b955060208701356117848161160c565b945060408701359350606087013567ffffffffffffffff8111156117a757600080fd5b6117b389828a0161162e565b909450925050608087013567ffffffffffffffff8111156117d357600080fd5b87016060818a0312156117e557600080fd5b809150509295509295509295565b60008060006060848603121561180857600080fd5b83356118138161160c565b925060208401356118238161160c565b929592945050506040919091013590565b6000806000806000806080878903121561184d57600080fd5b863567ffffffffffffffff81111561186457600080fd5b61187089828a0161162e565b90975095505060208701356118848161160c565b935060408701359250606087013567ffffffffffffffff8111156118a757600080fd5b6118b389828a0161162e565b979a9699509497509295939492505050565b801515811461079f57600080fd5b6000602082840312156118e557600080fd5b813561142f816118c5565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015283604082015260806060820152600061198b6080830184866118f0565b979650505050505050565b8381526040602082015260006119b06040830184866118f0565b95945050505050565b600081356119c68161160c565b73ffffffffffffffffffffffffffffffffffffffff168352602082013567ffffffffffffffff81168082146119fa57600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112611a3657600080fd5b820160208101903567ffffffffffffffff811115611a5357600080fd5b803603821315611a6257600080fd5b606060408601526119b06060860182846118f0565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a060608201526000611ac960a0830185876118f0565b8281036080840152611adb81856119b9565b9998505050505050505050565b848152606060208201526000611b026060830185876118f0565b828103604084015261198b81856119b9565b600060208284031215611b2657600080fd5b5051919050565b818103818111156104dc577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b606081526000611b7b6060830187896118f0565b8560208401528281036040840152611b948185876118f0565b98975050505050505050565b600060208284031215611bb257600080fd5b815161142f816118c5565b6000825160005b81811015611bde5760208186018101518583015201611bc4565b50600092019182525091905056fea264697066735822122056c2173fe07c19d0ad29056cae738132a6fc66c78781b309e0333efa436428de64736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60c060405260001960035534801561001657600080fd5b5060405161178338038061178383398101604081905261003591610220565b60016000819055805460ff19169055838383836001600160a01b038416158061006557506001600160a01b038316155b8061007757506001600160a01b038216155b8061008957506001600160a01b038116155b156100a75760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03808516608052831660a0526100c5600082610154565b506100f07f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610154565b5061011b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610154565b506101467f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610154565b505050505050505050610274565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166101fa5760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101b23390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101fe565b5060005b92915050565b80516001600160a01b038116811461021b57600080fd5b919050565b6000806000806080858703121561023657600080fd5b61023f85610204565b935061024d60208601610204565b925061025b60408601610204565b915061026960608601610204565b905092959194509250565b60805160a0516114ab6102d86000396000818161021d01528181610531015281816108290152818161099a01528181610aef0152610c110152600081816101d1015281816104a1015281816105040152818161079901526107fc01526114ab6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80635e3e9fef116100d857806391d148541161008c578063d547741f11610066578063d547741f14610386578063d5abeb0114610399578063e63ab1e9146103a257600080fd5b806391d1485414610311578063a217fddf14610357578063a783c7891461035f57600080fd5b8063743e0c9b116100bd578063743e0c9b146102cf5780638456cb59146102e257806385f438c1146102ea57600080fd5b80635e3e9fef146102a95780636f8b44b0146102bc57600080fd5b8063248a9ca31161012f57806336568abe1161011457806336568abe146102835780633f4ba83a146102965780635c975abb1461029e57600080fd5b8063248a9ca31461023f5780632f2ff15d1461027057600080fd5b8063106e629011610160578063106e6290146101b9578063116191b6146101cc57806321e093b11461021857600080fd5b806301ffc9a71461017c578063057e0f25146101a4575b600080fd5b61018f61018a366004610fd3565b6103c9565b60405190151581526020015b60405180910390f35b6101b76101b236600461108e565b610462565b005b6101b76101c7366004611126565b6105fc565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b61026261024d366004611159565b60009081526002602052604090206001015490565b60405190815260200161019b565b6101b761027e366004611172565b6106a1565b6101b7610291366004611172565b6106cc565b6101b7610725565b60015460ff1661018f565b6101b76102b736600461119e565b61075a565b6101b76102ca366004611159565b6108ef565b6101b76102dd366004611159565b61095d565b6101b7610a07565b6102627f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b61018f61031f366004611172565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610262600081565b6102627f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101b7610394366004611172565b610a39565b61026260035481565b6102627f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061045c57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b61046a610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461049481610aa1565b61049c610aab565b6104c77f00000000000000000000000000000000000000000000000000000000000000008785610aea565b6040517fd0b492c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610563907f0000000000000000000000000000000000000000000000000000000000000000908b908b908b908b908a9060040161130e565b600060405180830381600087803b15801561057d57600080fd5b505af1158015610591573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff167f52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c878787866040516105e1949392919061137f565b60405180910390a2506105f46001600055565b505050505050565b610604610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461062e81610aa1565b610636610aab565b610641848484610aea565b8373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d58460405161068991815260200190565b60405180910390a25061069c6001600055565b505050565b6000828152600260205260409020600101546106bc81610aa1565b6106c68383610c72565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331461071b576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61069c8282610d72565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61074f81610aa1565b610757610e31565b50565b610762610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461078c81610aa1565b610794610aab565b6107bf7f00000000000000000000000000000000000000000000000000000000000000008684610aea565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab5990610859907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a906004016113b6565b600060405180830381600087803b15801561087357600080fd5b505af1158015610887573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d8686866040516108d593929190611408565b60405180910390a2506108e86001600055565b5050505050565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61091981610aa1565b610921610aab565b60038290556040518281527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c9060200160405180910390a15050565b610965610aab565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b1580156109f357600080fd5b505af11580156108e8573d6000803e3d6000fd5b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a3181610aa1565b610757610eae565b600082815260026020526040902060010154610a5481610aa1565b6106c68383610d72565b600260005403610a9a576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6107578133610f07565b60015460ff1615610ae8576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6003547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7c9190611422565b610b86908461143b565b1115610bbe576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201849052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b158015610c5557600080fd5b505af1158015610c69573d6000803e3d6000fd5b50505050505050565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610d6a57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610d083390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161045c565b50600061045c565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610d6a57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161045c565b610e39610f97565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b610eb6610aab565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610e84565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610f93576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440160405180910390fd5b5050565b60015460ff16610ae8576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215610fe557600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461101557600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461104057600080fd5b919050565b60008083601f84011261105757600080fd5b50813567ffffffffffffffff81111561106f57600080fd5b60208301915083602082850101111561108757600080fd5b9250929050565b60008060008060008060a087890312156110a757600080fd5b6110b08761101c565b955060208701359450604087013567ffffffffffffffff8111156110d357600080fd5b6110df89828a01611045565b90955093505060608701359150608087013567ffffffffffffffff81111561110657600080fd5b87016060818a03121561111857600080fd5b809150509295509295509295565b60008060006060848603121561113b57600080fd5b6111448461101c565b95602085013595506040909401359392505050565b60006020828403121561116b57600080fd5b5035919050565b6000806040838503121561118557600080fd5b823591506111956020840161101c565b90509250929050565b6000806000806000608086880312156111b657600080fd5b6111bf8661101c565b945060208601359350604086013567ffffffffffffffff8111156111e257600080fd5b6111ee88828901611045565b96999598509660600135949350505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff6112678261101c565b1682526000602082013567ffffffffffffffff811680821461128857600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126112c457600080fd5b820160208101903567ffffffffffffffff8111156112e157600080fd5b8036038213156112f057600080fd5b60606040860152611305606086018284611200565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a06060820152600061136060a083018587611200565b82810360808401526113728185611249565b9998505050505050505050565b848152606060208201526000611399606083018587611200565b82810360408401526113ab8185611249565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006113ab608083018486611200565b838152604060208201526000611305604083018486611200565b60006020828403121561143457600080fd5b5051919050565b8082018082111561045c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea264697066735822122075681706993fbead450c86620a3d8fbe6bf0843c672f2e25fa00b8c0cb37b10c64736f6c634300081a00336080604052348015600f57600080fd5b506001600055610dbc806100246000396000f3fe6080604052600436106100635760003560e01c8063c513169111610040578063c5131691146100c1578063e04d4f97146100e1578063f05b6abf146100f457005b8063357fc5a21461006c578063660b9de01461008c5780636ed70169146100ac57005b3661006a57005b005b34801561007857600080fd5b5061006a6100873660046106bd565b610114565b34801561009857600080fd5b5061006a6100a73660046106f9565b6101aa565b3480156100b857600080fd5b5061006a6101e6565b3480156100cd57600080fd5b5061006a6100dc3660046106bd565b61021b565b61006a6100ef366004610859565b6102f6565b34801561010057600080fd5b5061006a61010f366004610945565b61033a565b61011c61036f565b61013e73ffffffffffffffffffffffffffffffffffffffff83163383866103b2565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101a56001600055565b505050565b7f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e33826040516101db929190610a78565b60405180910390a150565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61022361036f565b6000610230600285610b57565b90508060000361026c576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61028e73ffffffffffffffffffffffffffffffffffffffff84163384846103b2565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101a56001600055565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa333485858560405161032d959493929190610c00565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b1463384848460405161032d9493929190610c8a565b6002600054036103ab576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261044790859061044d565b50505050565b600061046f73ffffffffffffffffffffffffffffffffffffffff8416836104e8565b905080516000141580156104945750808060200190518101906104929190610d4d565b155b156101a5576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b60606104f6838360006104fd565b9392505050565b60608147101561053b576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016104df565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105649190610d6a565b60006040518083038185875af1925050503d80600081146105a1576040519150601f19603f3d011682016040523d82523d6000602084013e6105a6565b606091505b50915091506105b68683836105c0565b9695505050505050565b6060826105d5576105d08261064f565b6104f6565b81511580156105f9575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610648576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016104df565b50806104f6565b80511561065f5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff811681146106b857600080fd5b919050565b6000806000606084860312156106d257600080fd5b833592506106e260208501610694565b91506106f060408501610694565b90509250925092565b60006020828403121561070b57600080fd5b813567ffffffffffffffff81111561072257600080fd5b8201606081850312156104f657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156107aa576107aa610734565b604052919050565b600082601f8301126107c357600080fd5b813567ffffffffffffffff8111156107dd576107dd610734565b61080e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610763565b81815284602083860101111561082357600080fd5b816020850160208301376000918101602001919091529392505050565b801515811461069157600080fd5b80356106b881610840565b60008060006060848603121561086e57600080fd5b833567ffffffffffffffff81111561088557600080fd5b610891868287016107b2565b9350506020840135915060408401356108a981610840565b809150509250925092565b600067ffffffffffffffff8211156108ce576108ce610734565b5060051b60200190565b600082601f8301126108e957600080fd5b81356108fc6108f7826108b4565b610763565b8082825260208201915060208360051b86010192508583111561091e57600080fd5b602085015b8381101561093b578035835260209283019201610923565b5095945050505050565b60008060006060848603121561095a57600080fd5b833567ffffffffffffffff81111561097157600080fd5b8401601f8101861361098257600080fd5b80356109906108f7826108b4565b8082825260208201915060208360051b8501019250888311156109b257600080fd5b602084015b838110156109f457803567ffffffffffffffff8111156109d657600080fd5b6109e58b6020838901016107b2565b845250602092830192016109b7565b509550505050602084013567ffffffffffffffff811115610a1457600080fd5b610a20868287016108d8565b9250506106f06040850161084e565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610ab683610694565b1660408201526000602083013567ffffffffffffffff8116808214610ada57600080fd5b6060840152506040830135368490037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112610b1657600080fd5b830160208101903567ffffffffffffffff811115610b3357600080fd5b803603821315610b4257600080fd5b606060808501526105b660a085018284610a2f565b600082610b8d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60005b83811015610bad578181015183820152602001610b95565b50506000910152565b60008151808452610bce816020860160208601610b92565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610c3560a0830186610bb6565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610c80578151865260209586019590910190600101610c62565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610d1d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610d08858351610bb6565b94506020938401939190910190600101610cce565b505050508281036040840152610d338186610c4e565b915050610d44606083018415159052565b95945050505050565b600060208284031215610d5f57600080fd5b81516104f681610840565b60008251610d7c818460208701610b92565b919091019291505056fea26469706673582212206cfe8cc1e428962f6c4bd592761feb021724aeb74be964f4dd1bfa5d83cee47364736f6c634300081a0033a2646970667358221220b47cdac57c847fe6b172f610cdb503869f1ecf0abbeef85b91a16eccd43e5db264736f6c634300081a0033",
}

// ERC20CustodyTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20CustodyTestMetaData.ABI instead.
var ERC20CustodyTestABI = ERC20CustodyTestMetaData.ABI

// ERC20CustodyTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20CustodyTestMetaData.Bin instead.
var ERC20CustodyTestBin = ERC20CustodyTestMetaData.Bin

// DeployERC20CustodyTest deploys a new Ethereum contract, binding an instance of ERC20CustodyTest to it.
func DeployERC20CustodyTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20CustodyTest, error) {
	parsed, err := ERC20CustodyTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20CustodyTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20CustodyTest{ERC20CustodyTestCaller: ERC20CustodyTestCaller{contract: contract}, ERC20CustodyTestTransactor: ERC20CustodyTestTransactor{contract: contract}, ERC20CustodyTestFilterer: ERC20CustodyTestFilterer{contract: contract}}, nil
}

// ERC20CustodyTest is an auto generated Go binding around an Ethereum contract.
type ERC20CustodyTest struct {
	ERC20CustodyTestCaller     // Read-only binding to the contract
	ERC20CustodyTestTransactor // Write-only binding to the contract
	ERC20CustodyTestFilterer   // Log filterer for contract events
}

// ERC20CustodyTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20CustodyTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20CustodyTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20CustodyTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20CustodyTestSession struct {
	Contract     *ERC20CustodyTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CustodyTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CustodyTestCallerSession struct {
	Contract *ERC20CustodyTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC20CustodyTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20CustodyTestTransactorSession struct {
	Contract     *ERC20CustodyTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC20CustodyTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20CustodyTestRaw struct {
	Contract *ERC20CustodyTest // Generic contract binding to access the raw methods on
}

// ERC20CustodyTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CustodyTestCallerRaw struct {
	Contract *ERC20CustodyTestCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20CustodyTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20CustodyTestTransactorRaw struct {
	Contract *ERC20CustodyTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20CustodyTest creates a new instance of ERC20CustodyTest, bound to a specific deployed contract.
func NewERC20CustodyTest(address common.Address, backend bind.ContractBackend) (*ERC20CustodyTest, error) {
	contract, err := bindERC20CustodyTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTest{ERC20CustodyTestCaller: ERC20CustodyTestCaller{contract: contract}, ERC20CustodyTestTransactor: ERC20CustodyTestTransactor{contract: contract}, ERC20CustodyTestFilterer: ERC20CustodyTestFilterer{contract: contract}}, nil
}

// NewERC20CustodyTestCaller creates a new read-only instance of ERC20CustodyTest, bound to a specific deployed contract.
func NewERC20CustodyTestCaller(address common.Address, caller bind.ContractCaller) (*ERC20CustodyTestCaller, error) {
	contract, err := bindERC20CustodyTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestCaller{contract: contract}, nil
}

// NewERC20CustodyTestTransactor creates a new write-only instance of ERC20CustodyTest, bound to a specific deployed contract.
func NewERC20CustodyTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20CustodyTestTransactor, error) {
	contract, err := bindERC20CustodyTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestTransactor{contract: contract}, nil
}

// NewERC20CustodyTestFilterer creates a new log filterer instance of ERC20CustodyTest, bound to a specific deployed contract.
func NewERC20CustodyTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20CustodyTestFilterer, error) {
	contract, err := bindERC20CustodyTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestFilterer{contract: contract}, nil
}

// bindERC20CustodyTest binds a generic wrapper to an already deployed contract.
func bindERC20CustodyTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20CustodyTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20CustodyTest *ERC20CustodyTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20CustodyTest.Contract.ERC20CustodyTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20CustodyTest *ERC20CustodyTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.ERC20CustodyTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20CustodyTest *ERC20CustodyTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.ERC20CustodyTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20CustodyTest *ERC20CustodyTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20CustodyTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20CustodyTest *ERC20CustodyTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20CustodyTest *ERC20CustodyTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.contract.Transact(opts, method, params...)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) ASSETHANDLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "ASSET_HANDLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.ASSETHANDLERROLE(&_ERC20CustodyTest.CallOpts)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.ASSETHANDLERROLE(&_ERC20CustodyTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_ERC20CustodyTest *ERC20CustodyTestSession) ISTEST() (bool, error) {
	return _ERC20CustodyTest.Contract.ISTEST(&_ERC20CustodyTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) ISTEST() (bool, error) {
	return _ERC20CustodyTest.Contract.ISTEST(&_ERC20CustodyTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestSession) PAUSERROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.PAUSERROLE(&_ERC20CustodyTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.PAUSERROLE(&_ERC20CustodyTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestSession) TSSROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.TSSROLE(&_ERC20CustodyTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) TSSROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.TSSROLE(&_ERC20CustodyTest.CallOpts)
}

// WHITELISTERROLE is a free data retrieval call binding the contract method 0x570618e1.
//
// Solidity: function WHITELISTER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) WHITELISTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "WHITELISTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WHITELISTERROLE is a free data retrieval call binding the contract method 0x570618e1.
//
// Solidity: function WHITELISTER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestSession) WHITELISTERROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.WHITELISTERROLE(&_ERC20CustodyTest.CallOpts)
}

// WHITELISTERROLE is a free data retrieval call binding the contract method 0x570618e1.
//
// Solidity: function WHITELISTER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) WHITELISTERROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.WHITELISTERROLE(&_ERC20CustodyTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.WITHDRAWERROLE(&_ERC20CustodyTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.WITHDRAWERROLE(&_ERC20CustodyTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_ERC20CustodyTest *ERC20CustodyTestSession) ExcludeArtifacts() ([]string, error) {
	return _ERC20CustodyTest.Contract.ExcludeArtifacts(&_ERC20CustodyTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _ERC20CustodyTest.Contract.ExcludeArtifacts(&_ERC20CustodyTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_ERC20CustodyTest *ERC20CustodyTestSession) ExcludeContracts() ([]common.Address, error) {
	return _ERC20CustodyTest.Contract.ExcludeContracts(&_ERC20CustodyTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _ERC20CustodyTest.Contract.ExcludeContracts(&_ERC20CustodyTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_ERC20CustodyTest *ERC20CustodyTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ERC20CustodyTest.Contract.ExcludeSelectors(&_ERC20CustodyTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ERC20CustodyTest.Contract.ExcludeSelectors(&_ERC20CustodyTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_ERC20CustodyTest *ERC20CustodyTestSession) ExcludeSenders() ([]common.Address, error) {
	return _ERC20CustodyTest.Contract.ExcludeSenders(&_ERC20CustodyTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _ERC20CustodyTest.Contract.ExcludeSenders(&_ERC20CustodyTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_ERC20CustodyTest *ERC20CustodyTestSession) Failed() (bool, error) {
	return _ERC20CustodyTest.Contract.Failed(&_ERC20CustodyTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) Failed() (bool, error) {
	return _ERC20CustodyTest.Contract.Failed(&_ERC20CustodyTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_ERC20CustodyTest *ERC20CustodyTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _ERC20CustodyTest.Contract.TargetArtifactSelectors(&_ERC20CustodyTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _ERC20CustodyTest.Contract.TargetArtifactSelectors(&_ERC20CustodyTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_ERC20CustodyTest *ERC20CustodyTestSession) TargetArtifacts() ([]string, error) {
	return _ERC20CustodyTest.Contract.TargetArtifacts(&_ERC20CustodyTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) TargetArtifacts() ([]string, error) {
	return _ERC20CustodyTest.Contract.TargetArtifacts(&_ERC20CustodyTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_ERC20CustodyTest *ERC20CustodyTestSession) TargetContracts() ([]common.Address, error) {
	return _ERC20CustodyTest.Contract.TargetContracts(&_ERC20CustodyTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _ERC20CustodyTest.Contract.TargetContracts(&_ERC20CustodyTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_ERC20CustodyTest *ERC20CustodyTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _ERC20CustodyTest.Contract.TargetInterfaces(&_ERC20CustodyTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _ERC20CustodyTest.Contract.TargetInterfaces(&_ERC20CustodyTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_ERC20CustodyTest *ERC20CustodyTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ERC20CustodyTest.Contract.TargetSelectors(&_ERC20CustodyTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ERC20CustodyTest.Contract.TargetSelectors(&_ERC20CustodyTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_ERC20CustodyTest *ERC20CustodyTestSession) TargetSenders() ([]common.Address, error) {
	return _ERC20CustodyTest.Contract.TargetSenders(&_ERC20CustodyTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _ERC20CustodyTest.Contract.TargetSenders(&_ERC20CustodyTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) SetUp() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.SetUp(&_ERC20CustodyTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.SetUp(&_ERC20CustodyTest.TransactOpts)
}

// TestDepositLegacy is a paid mutator transaction binding the contract method 0xb421ca70.
//
// Solidity: function testDepositLegacy() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestDepositLegacy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testDepositLegacy")
}

// TestDepositLegacy is a paid mutator transaction binding the contract method 0xb421ca70.
//
// Solidity: function testDepositLegacy() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestDepositLegacy() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestDepositLegacy(&_ERC20CustodyTest.TransactOpts)
}

// TestDepositLegacy is a paid mutator transaction binding the contract method 0xb421ca70.
//
// Solidity: function testDepositLegacy() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestDepositLegacy() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestDepositLegacy(&_ERC20CustodyTest.TransactOpts)
}

// TestDepositLegacyFailsIfNotSupported is a paid mutator transaction binding the contract method 0x49c783dd.
//
// Solidity: function testDepositLegacyFailsIfNotSupported() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestDepositLegacyFailsIfNotSupported(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testDepositLegacyFailsIfNotSupported")
}

// TestDepositLegacyFailsIfNotSupported is a paid mutator transaction binding the contract method 0x49c783dd.
//
// Solidity: function testDepositLegacyFailsIfNotSupported() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestDepositLegacyFailsIfNotSupported() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestDepositLegacyFailsIfNotSupported(&_ERC20CustodyTest.TransactOpts)
}

// TestDepositLegacyFailsIfNotSupported is a paid mutator transaction binding the contract method 0x49c783dd.
//
// Solidity: function testDepositLegacyFailsIfNotSupported() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestDepositLegacyFailsIfNotSupported() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestDepositLegacyFailsIfNotSupported(&_ERC20CustodyTest.TransactOpts)
}

// TestDepositLegacyFailsIfTokenNotWhitelisted is a paid mutator transaction binding the contract method 0x7099d6f8.
//
// Solidity: function testDepositLegacyFailsIfTokenNotWhitelisted() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestDepositLegacyFailsIfTokenNotWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testDepositLegacyFailsIfTokenNotWhitelisted")
}

// TestDepositLegacyFailsIfTokenNotWhitelisted is a paid mutator transaction binding the contract method 0x7099d6f8.
//
// Solidity: function testDepositLegacyFailsIfTokenNotWhitelisted() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestDepositLegacyFailsIfTokenNotWhitelisted() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestDepositLegacyFailsIfTokenNotWhitelisted(&_ERC20CustodyTest.TransactOpts)
}

// TestDepositLegacyFailsIfTokenNotWhitelisted is a paid mutator transaction binding the contract method 0x7099d6f8.
//
// Solidity: function testDepositLegacyFailsIfTokenNotWhitelisted() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestDepositLegacyFailsIfTokenNotWhitelisted() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestDepositLegacyFailsIfTokenNotWhitelisted(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20PartialThroughCustody is a paid mutator transaction binding the contract method 0xfb176c12.
//
// Solidity: function testForwardCallToReceiveERC20PartialThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestForwardCallToReceiveERC20PartialThroughCustody(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testForwardCallToReceiveERC20PartialThroughCustody")
}

// TestForwardCallToReceiveERC20PartialThroughCustody is a paid mutator transaction binding the contract method 0xfb176c12.
//
// Solidity: function testForwardCallToReceiveERC20PartialThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestForwardCallToReceiveERC20PartialThroughCustody() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20PartialThroughCustody(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20PartialThroughCustody is a paid mutator transaction binding the contract method 0xfb176c12.
//
// Solidity: function testForwardCallToReceiveERC20PartialThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestForwardCallToReceiveERC20PartialThroughCustody() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20PartialThroughCustody(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x6a621854.
//
// Solidity: function testForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0")
}

// TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x6a621854.
//
// Solidity: function testForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x6a621854.
//
// Solidity: function testForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x0eee72a9.
//
// Solidity: function testForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer")
}

// TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x0eee72a9.
//
// Solidity: function testForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x0eee72a9.
//
// Solidity: function testForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20ThroughCustody is a paid mutator transaction binding the contract method 0xcbd57e2f.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestForwardCallToReceiveERC20ThroughCustody(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testForwardCallToReceiveERC20ThroughCustody")
}

// TestForwardCallToReceiveERC20ThroughCustody is a paid mutator transaction binding the contract method 0xcbd57e2f.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestForwardCallToReceiveERC20ThroughCustody() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20ThroughCustody(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20ThroughCustody is a paid mutator transaction binding the contract method 0xcbd57e2f.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestForwardCallToReceiveERC20ThroughCustody() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20ThroughCustody(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x1779672f.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0")
}

// TestForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x1779672f.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x1779672f.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51ecdf3c.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress")
}

// TestForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51ecdf3c.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51ecdf3c.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x3ee92923.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer")
}

// TestForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x3ee92923.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x3ee92923.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20ThroughCustodyTogglePause is a paid mutator transaction binding the contract method 0xc713f827.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustodyTogglePause() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestForwardCallToReceiveERC20ThroughCustodyTogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testForwardCallToReceiveERC20ThroughCustodyTogglePause")
}

// TestForwardCallToReceiveERC20ThroughCustodyTogglePause is a paid mutator transaction binding the contract method 0xc713f827.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustodyTogglePause() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestForwardCallToReceiveERC20ThroughCustodyTogglePause() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20ThroughCustodyTogglePause(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveERC20ThroughCustodyTogglePause is a paid mutator transaction binding the contract method 0xc713f827.
//
// Solidity: function testForwardCallToReceiveERC20ThroughCustodyTogglePause() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestForwardCallToReceiveERC20ThroughCustodyTogglePause() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveERC20ThroughCustodyTogglePause(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveNoParamsThroughCustody is a paid mutator transaction binding the contract method 0xa3f9d0e0.
//
// Solidity: function testForwardCallToReceiveNoParamsThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestForwardCallToReceiveNoParamsThroughCustody(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testForwardCallToReceiveNoParamsThroughCustody")
}

// TestForwardCallToReceiveNoParamsThroughCustody is a paid mutator transaction binding the contract method 0xa3f9d0e0.
//
// Solidity: function testForwardCallToReceiveNoParamsThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestForwardCallToReceiveNoParamsThroughCustody() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveNoParamsThroughCustody(&_ERC20CustodyTest.TransactOpts)
}

// TestForwardCallToReceiveNoParamsThroughCustody is a paid mutator transaction binding the contract method 0xa3f9d0e0.
//
// Solidity: function testForwardCallToReceiveNoParamsThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestForwardCallToReceiveNoParamsThroughCustody() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestForwardCallToReceiveNoParamsThroughCustody(&_ERC20CustodyTest.TransactOpts)
}

// TestNewCustodyFailsIfAddressesAreZero is a paid mutator transaction binding the contract method 0x4b5838d2.
//
// Solidity: function testNewCustodyFailsIfAddressesAreZero() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestNewCustodyFailsIfAddressesAreZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testNewCustodyFailsIfAddressesAreZero")
}

// TestNewCustodyFailsIfAddressesAreZero is a paid mutator transaction binding the contract method 0x4b5838d2.
//
// Solidity: function testNewCustodyFailsIfAddressesAreZero() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestNewCustodyFailsIfAddressesAreZero() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestNewCustodyFailsIfAddressesAreZero(&_ERC20CustodyTest.TransactOpts)
}

// TestNewCustodyFailsIfAddressesAreZero is a paid mutator transaction binding the contract method 0x4b5838d2.
//
// Solidity: function testNewCustodyFailsIfAddressesAreZero() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestNewCustodyFailsIfAddressesAreZero() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestNewCustodyFailsIfAddressesAreZero(&_ERC20CustodyTest.TransactOpts)
}

// TestUnwhitelist is a paid mutator transaction binding the contract method 0xfa2a7074.
//
// Solidity: function testUnwhitelist() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestUnwhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testUnwhitelist")
}

// TestUnwhitelist is a paid mutator transaction binding the contract method 0xfa2a7074.
//
// Solidity: function testUnwhitelist() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestUnwhitelist() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestUnwhitelist(&_ERC20CustodyTest.TransactOpts)
}

// TestUnwhitelist is a paid mutator transaction binding the contract method 0xfa2a7074.
//
// Solidity: function testUnwhitelist() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestUnwhitelist() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestUnwhitelist(&_ERC20CustodyTest.TransactOpts)
}

// TestUnwhitelistFailsIfSenderIsNotWhitelister is a paid mutator transaction binding the contract method 0xf0c8e7e0.
//
// Solidity: function testUnwhitelistFailsIfSenderIsNotWhitelister() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestUnwhitelistFailsIfSenderIsNotWhitelister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testUnwhitelistFailsIfSenderIsNotWhitelister")
}

// TestUnwhitelistFailsIfSenderIsNotWhitelister is a paid mutator transaction binding the contract method 0xf0c8e7e0.
//
// Solidity: function testUnwhitelistFailsIfSenderIsNotWhitelister() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestUnwhitelistFailsIfSenderIsNotWhitelister() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestUnwhitelistFailsIfSenderIsNotWhitelister(&_ERC20CustodyTest.TransactOpts)
}

// TestUnwhitelistFailsIfSenderIsNotWhitelister is a paid mutator transaction binding the contract method 0xf0c8e7e0.
//
// Solidity: function testUnwhitelistFailsIfSenderIsNotWhitelister() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestUnwhitelistFailsIfSenderIsNotWhitelister() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestUnwhitelistFailsIfSenderIsNotWhitelister(&_ERC20CustodyTest.TransactOpts)
}

// TestUnwhitelistFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x9158c623.
//
// Solidity: function testUnwhitelistFailsIfZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestUnwhitelistFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testUnwhitelistFailsIfZeroAddress")
}

// TestUnwhitelistFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x9158c623.
//
// Solidity: function testUnwhitelistFailsIfZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestUnwhitelistFailsIfZeroAddress() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestUnwhitelistFailsIfZeroAddress(&_ERC20CustodyTest.TransactOpts)
}

// TestUnwhitelistFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x9158c623.
//
// Solidity: function testUnwhitelistFailsIfZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestUnwhitelistFailsIfZeroAddress() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestUnwhitelistFailsIfZeroAddress(&_ERC20CustodyTest.TransactOpts)
}

// TestWhitelist is a paid mutator transaction binding the contract method 0x284cb929.
//
// Solidity: function testWhitelist() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testWhitelist")
}

// TestWhitelist is a paid mutator transaction binding the contract method 0x284cb929.
//
// Solidity: function testWhitelist() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestWhitelist() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWhitelist(&_ERC20CustodyTest.TransactOpts)
}

// TestWhitelist is a paid mutator transaction binding the contract method 0x284cb929.
//
// Solidity: function testWhitelist() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestWhitelist() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWhitelist(&_ERC20CustodyTest.TransactOpts)
}

// TestWhitelistFailsIfSenderIsNotWhitelister is a paid mutator transaction binding the contract method 0x2be6a162.
//
// Solidity: function testWhitelistFailsIfSenderIsNotWhitelister() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestWhitelistFailsIfSenderIsNotWhitelister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testWhitelistFailsIfSenderIsNotWhitelister")
}

// TestWhitelistFailsIfSenderIsNotWhitelister is a paid mutator transaction binding the contract method 0x2be6a162.
//
// Solidity: function testWhitelistFailsIfSenderIsNotWhitelister() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestWhitelistFailsIfSenderIsNotWhitelister() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWhitelistFailsIfSenderIsNotWhitelister(&_ERC20CustodyTest.TransactOpts)
}

// TestWhitelistFailsIfSenderIsNotWhitelister is a paid mutator transaction binding the contract method 0x2be6a162.
//
// Solidity: function testWhitelistFailsIfSenderIsNotWhitelister() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestWhitelistFailsIfSenderIsNotWhitelister() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWhitelistFailsIfSenderIsNotWhitelister(&_ERC20CustodyTest.TransactOpts)
}

// TestWhitelistFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x9fc7fd55.
//
// Solidity: function testWhitelistFailsIfZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestWhitelistFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testWhitelistFailsIfZeroAddress")
}

// TestWhitelistFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x9fc7fd55.
//
// Solidity: function testWhitelistFailsIfZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestWhitelistFailsIfZeroAddress() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWhitelistFailsIfZeroAddress(&_ERC20CustodyTest.TransactOpts)
}

// TestWhitelistFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x9fc7fd55.
//
// Solidity: function testWhitelistFailsIfZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestWhitelistFailsIfZeroAddress() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWhitelistFailsIfZeroAddress(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawAndCallFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0x9918c1c2.
//
// Solidity: function testWithdrawAndCallFailsIfTokenIsNotWhitelisted() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestWithdrawAndCallFailsIfTokenIsNotWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testWithdrawAndCallFailsIfTokenIsNotWhitelisted")
}

// TestWithdrawAndCallFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0x9918c1c2.
//
// Solidity: function testWithdrawAndCallFailsIfTokenIsNotWhitelisted() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestWithdrawAndCallFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawAndCallFailsIfTokenIsNotWhitelisted(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawAndCallFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0x9918c1c2.
//
// Solidity: function testWithdrawAndCallFailsIfTokenIsNotWhitelisted() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestWithdrawAndCallFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawAndCallFailsIfTokenIsNotWhitelisted(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawAndRevertFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xf4221f08.
//
// Solidity: function testWithdrawAndRevertFailsIfTokenIsNotWhitelisted() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestWithdrawAndRevertFailsIfTokenIsNotWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testWithdrawAndRevertFailsIfTokenIsNotWhitelisted")
}

// TestWithdrawAndRevertFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xf4221f08.
//
// Solidity: function testWithdrawAndRevertFailsIfTokenIsNotWhitelisted() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestWithdrawAndRevertFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawAndRevertFailsIfTokenIsNotWhitelisted(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawAndRevertFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xf4221f08.
//
// Solidity: function testWithdrawAndRevertFailsIfTokenIsNotWhitelisted() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestWithdrawAndRevertFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawAndRevertFailsIfTokenIsNotWhitelisted(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawAndRevertThroughCustody is a paid mutator transaction binding the contract method 0x71149c94.
//
// Solidity: function testWithdrawAndRevertThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestWithdrawAndRevertThroughCustody(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testWithdrawAndRevertThroughCustody")
}

// TestWithdrawAndRevertThroughCustody is a paid mutator transaction binding the contract method 0x71149c94.
//
// Solidity: function testWithdrawAndRevertThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestWithdrawAndRevertThroughCustody() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawAndRevertThroughCustody(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawAndRevertThroughCustody is a paid mutator transaction binding the contract method 0x71149c94.
//
// Solidity: function testWithdrawAndRevertThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestWithdrawAndRevertThroughCustody() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawAndRevertThroughCustody(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawAndRevertThroughCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xeb1ce7f9.
//
// Solidity: function testWithdrawAndRevertThroughCustodyFailsIfAmountIs0() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestWithdrawAndRevertThroughCustodyFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testWithdrawAndRevertThroughCustodyFailsIfAmountIs0")
}

// TestWithdrawAndRevertThroughCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xeb1ce7f9.
//
// Solidity: function testWithdrawAndRevertThroughCustodyFailsIfAmountIs0() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestWithdrawAndRevertThroughCustodyFailsIfAmountIs0() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawAndRevertThroughCustodyFailsIfAmountIs0(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawAndRevertThroughCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xeb1ce7f9.
//
// Solidity: function testWithdrawAndRevertThroughCustodyFailsIfAmountIs0() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestWithdrawAndRevertThroughCustodyFailsIfAmountIs0() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawAndRevertThroughCustodyFailsIfAmountIs0(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x7e91c50f.
//
// Solidity: function testWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress")
}

// TestWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x7e91c50f.
//
// Solidity: function testWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x7e91c50f.
//
// Solidity: function testWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xa4943deb.
//
// Solidity: function testWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer")
}

// TestWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xa4943deb.
//
// Solidity: function testWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xa4943deb.
//
// Solidity: function testWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0x82c52992.
//
// Solidity: function testWithdrawFailsIfTokenIsNotWhitelisted() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestWithdrawFailsIfTokenIsNotWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testWithdrawFailsIfTokenIsNotWhitelisted")
}

// TestWithdrawFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0x82c52992.
//
// Solidity: function testWithdrawFailsIfTokenIsNotWhitelisted() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestWithdrawFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawFailsIfTokenIsNotWhitelisted(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0x82c52992.
//
// Solidity: function testWithdrawFailsIfTokenIsNotWhitelisted() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestWithdrawFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawFailsIfTokenIsNotWhitelisted(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawThroughCustody is a paid mutator transaction binding the contract method 0x3e73ecb4.
//
// Solidity: function testWithdrawThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestWithdrawThroughCustody(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testWithdrawThroughCustody")
}

// TestWithdrawThroughCustody is a paid mutator transaction binding the contract method 0x3e73ecb4.
//
// Solidity: function testWithdrawThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestWithdrawThroughCustody() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawThroughCustody(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawThroughCustody is a paid mutator transaction binding the contract method 0x3e73ecb4.
//
// Solidity: function testWithdrawThroughCustody() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestWithdrawThroughCustody() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawThroughCustody(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xfe8e5f1b.
//
// Solidity: function testWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer")
}

// TestWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xfe8e5f1b.
//
// Solidity: function testWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer(&_ERC20CustodyTest.TransactOpts)
}

// TestWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xfe8e5f1b.
//
// Solidity: function testWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer(&_ERC20CustodyTest.TransactOpts)
}

// ERC20CustodyTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestCalledIterator struct {
	Event *ERC20CustodyTestCalled // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestCalled)
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
		it.Event = new(ERC20CustodyTestCalled)
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
func (it *ERC20CustodyTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestCalled represents a Called event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ERC20CustodyTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestCalledIterator{contract: _ERC20CustodyTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestCalled)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseCalled(log types.Log) (*ERC20CustodyTestCalled, error) {
	event := new(ERC20CustodyTestCalled)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestDepositedIterator struct {
	Event *ERC20CustodyTestDeposited // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestDeposited)
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
		it.Event = new(ERC20CustodyTestDeposited)
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
func (it *ERC20CustodyTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestDeposited represents a Deposited event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestDeposited struct {
	Recipient []byte
	Asset     common.Address
	Amount    *big.Int
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterDeposited(opts *bind.FilterOpts, asset []common.Address) (*ERC20CustodyTestDepositedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestDepositedIterator{contract: _ERC20CustodyTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestDeposited, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestDeposited)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseDeposited(log types.Log) (*ERC20CustodyTestDeposited, error) {
	event := new(ERC20CustodyTestDeposited)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestDeposited0Iterator is returned from FilterDeposited0 and is used to iterate over the raw logs and unpacked data for Deposited0 events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestDeposited0Iterator struct {
	Event *ERC20CustodyTestDeposited0 // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestDeposited0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestDeposited0)
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
		it.Event = new(ERC20CustodyTestDeposited0)
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
func (it *ERC20CustodyTestDeposited0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestDeposited0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestDeposited0 represents a Deposited0 event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestDeposited0 struct {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterDeposited0(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ERC20CustodyTestDeposited0Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "Deposited0", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestDeposited0Iterator{contract: _ERC20CustodyTest.contract, event: "Deposited0", logs: logs, sub: sub}, nil
}

// WatchDeposited0 is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchDeposited0(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestDeposited0, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "Deposited0", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestDeposited0)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "Deposited0", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseDeposited0(log types.Log) (*ERC20CustodyTestDeposited0, error) {
	event := new(ERC20CustodyTestDeposited0)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "Deposited0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestExecutedIterator struct {
	Event *ERC20CustodyTestExecuted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestExecuted)
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
		it.Event = new(ERC20CustodyTestExecuted)
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
func (it *ERC20CustodyTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestExecuted represents a Executed event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*ERC20CustodyTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestExecutedIterator{contract: _ERC20CustodyTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestExecuted)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "Executed", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseExecuted(log types.Log) (*ERC20CustodyTestExecuted, error) {
	event := new(ERC20CustodyTestExecuted)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestExecutedWithERC20Iterator struct {
	Event *ERC20CustodyTestExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestExecutedWithERC20)
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
		it.Event = new(ERC20CustodyTestExecutedWithERC20)
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
func (it *ERC20CustodyTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*ERC20CustodyTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestExecutedWithERC20Iterator{contract: _ERC20CustodyTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestExecutedWithERC20)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseExecutedWithERC20(log types.Log) (*ERC20CustodyTestExecutedWithERC20, error) {
	event := new(ERC20CustodyTestExecutedWithERC20)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReceivedERC20Iterator struct {
	Event *ERC20CustodyTestReceivedERC20 // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestReceivedERC20)
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
		it.Event = new(ERC20CustodyTestReceivedERC20)
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
func (it *ERC20CustodyTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestReceivedERC20 represents a ReceivedERC20 event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*ERC20CustodyTestReceivedERC20Iterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestReceivedERC20Iterator{contract: _ERC20CustodyTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestReceivedERC20)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseReceivedERC20(log types.Log) (*ERC20CustodyTestReceivedERC20, error) {
	event := new(ERC20CustodyTestReceivedERC20)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReceivedNoParamsIterator struct {
	Event *ERC20CustodyTestReceivedNoParams // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestReceivedNoParams)
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
		it.Event = new(ERC20CustodyTestReceivedNoParams)
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
func (it *ERC20CustodyTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestReceivedNoParams represents a ReceivedNoParams event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*ERC20CustodyTestReceivedNoParamsIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestReceivedNoParamsIterator{contract: _ERC20CustodyTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestReceivedNoParams)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseReceivedNoParams(log types.Log) (*ERC20CustodyTestReceivedNoParams, error) {
	event := new(ERC20CustodyTestReceivedNoParams)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReceivedNonPayableIterator struct {
	Event *ERC20CustodyTestReceivedNonPayable // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestReceivedNonPayable)
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
		it.Event = new(ERC20CustodyTestReceivedNonPayable)
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
func (it *ERC20CustodyTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestReceivedNonPayable represents a ReceivedNonPayable event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*ERC20CustodyTestReceivedNonPayableIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestReceivedNonPayableIterator{contract: _ERC20CustodyTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestReceivedNonPayable)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseReceivedNonPayable(log types.Log) (*ERC20CustodyTestReceivedNonPayable, error) {
	event := new(ERC20CustodyTestReceivedNonPayable)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReceivedPayableIterator struct {
	Event *ERC20CustodyTestReceivedPayable // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestReceivedPayable)
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
		it.Event = new(ERC20CustodyTestReceivedPayable)
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
func (it *ERC20CustodyTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestReceivedPayable represents a ReceivedPayable event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReceivedPayable struct {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*ERC20CustodyTestReceivedPayableIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestReceivedPayableIterator{contract: _ERC20CustodyTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestReceivedPayable)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseReceivedPayable(log types.Log) (*ERC20CustodyTestReceivedPayable, error) {
	event := new(ERC20CustodyTestReceivedPayable)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReceivedRevertIterator struct {
	Event *ERC20CustodyTestReceivedRevert // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestReceivedRevert)
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
		it.Event = new(ERC20CustodyTestReceivedRevert)
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
func (it *ERC20CustodyTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestReceivedRevert represents a ReceivedRevert event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*ERC20CustodyTestReceivedRevertIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestReceivedRevertIterator{contract: _ERC20CustodyTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestReceivedRevert)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
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

// ParseReceivedRevert is a log parse operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseReceivedRevert(log types.Log) (*ERC20CustodyTestReceivedRevert, error) {
	event := new(ERC20CustodyTestReceivedRevert)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestRevertedIterator struct {
	Event *ERC20CustodyTestReverted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestReverted)
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
		it.Event = new(ERC20CustodyTestReverted)
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
func (it *ERC20CustodyTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestReverted represents a Reverted event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReverted struct {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestRevertedIterator{contract: _ERC20CustodyTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0x1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b1436.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestReverted)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "Reverted", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseReverted(log types.Log) (*ERC20CustodyTestReverted, error) {
	event := new(ERC20CustodyTestReverted)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestUnwhitelistedIterator is returned from FilterUnwhitelisted and is used to iterate over the raw logs and unpacked data for Unwhitelisted events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestUnwhitelistedIterator struct {
	Event *ERC20CustodyTestUnwhitelisted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestUnwhitelisted)
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
		it.Event = new(ERC20CustodyTestUnwhitelisted)
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
func (it *ERC20CustodyTestUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestUnwhitelisted represents a Unwhitelisted event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestUnwhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnwhitelisted is a free log retrieval operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterUnwhitelisted(opts *bind.FilterOpts, token []common.Address) (*ERC20CustodyTestUnwhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestUnwhitelistedIterator{contract: _ERC20CustodyTest.contract, event: "Unwhitelisted", logs: logs, sub: sub}, nil
}

// WatchUnwhitelisted is a free log subscription operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchUnwhitelisted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestUnwhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestUnwhitelisted)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseUnwhitelisted(log types.Log) (*ERC20CustodyTestUnwhitelisted, error) {
	event := new(ERC20CustodyTestUnwhitelisted)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestWhitelistedIterator struct {
	Event *ERC20CustodyTestWhitelisted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestWhitelisted)
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
		it.Event = new(ERC20CustodyTestWhitelisted)
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
func (it *ERC20CustodyTestWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestWhitelisted represents a Whitelisted event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestWhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterWhitelisted(opts *bind.FilterOpts, token []common.Address) (*ERC20CustodyTestWhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestWhitelistedIterator{contract: _ERC20CustodyTest.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestWhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestWhitelisted)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseWhitelisted(log types.Log) (*ERC20CustodyTestWhitelisted, error) {
	event := new(ERC20CustodyTestWhitelisted)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestWithdrawnIterator struct {
	Event *ERC20CustodyTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestWithdrawn)
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
		it.Event = new(ERC20CustodyTestWithdrawn)
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
func (it *ERC20CustodyTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestWithdrawn represents a Withdrawn event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestWithdrawn struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyTestWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestWithdrawnIterator{contract: _ERC20CustodyTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestWithdrawn, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestWithdrawn)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseWithdrawn(log types.Log) (*ERC20CustodyTestWithdrawn, error) {
	event := new(ERC20CustodyTestWithdrawn)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestWithdrawnAndCalledIterator struct {
	Event *ERC20CustodyTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestWithdrawnAndCalled)
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
		it.Event = new(ERC20CustodyTestWithdrawnAndCalled)
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
func (it *ERC20CustodyTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestWithdrawnAndCalled struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyTestWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestWithdrawnAndCalledIterator{contract: _ERC20CustodyTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestWithdrawnAndCalled, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestWithdrawnAndCalled)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*ERC20CustodyTestWithdrawnAndCalled, error) {
	event := new(ERC20CustodyTestWithdrawnAndCalled)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestWithdrawnAndRevertedIterator struct {
	Event *ERC20CustodyTestWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestWithdrawnAndReverted)
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
		it.Event = new(ERC20CustodyTestWithdrawnAndReverted)
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
func (it *ERC20CustodyTestWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestWithdrawnAndReverted struct {
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyTestWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestWithdrawnAndRevertedIterator{contract: _ERC20CustodyTest.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestWithdrawnAndReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestWithdrawnAndReverted)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*ERC20CustodyTestWithdrawnAndReverted, error) {
	event := new(ERC20CustodyTestWithdrawnAndReverted)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogIterator struct {
	Event *ERC20CustodyTestLog // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLog)
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
		it.Event = new(ERC20CustodyTestLog)
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
func (it *ERC20CustodyTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLog represents a Log event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLog(opts *bind.FilterOpts) (*ERC20CustodyTestLogIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogIterator{contract: _ERC20CustodyTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLog) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLog)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLog(log types.Log) (*ERC20CustodyTestLog, error) {
	event := new(ERC20CustodyTestLog)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogAddressIterator struct {
	Event *ERC20CustodyTestLogAddress // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogAddress)
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
		it.Event = new(ERC20CustodyTestLogAddress)
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
func (it *ERC20CustodyTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogAddress represents a LogAddress event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*ERC20CustodyTestLogAddressIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogAddressIterator{contract: _ERC20CustodyTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogAddress)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_address", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogAddress(log types.Log) (*ERC20CustodyTestLogAddress, error) {
	event := new(ERC20CustodyTestLogAddress)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogArrayIterator struct {
	Event *ERC20CustodyTestLogArray // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogArray)
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
		it.Event = new(ERC20CustodyTestLogArray)
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
func (it *ERC20CustodyTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogArray represents a LogArray event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*ERC20CustodyTestLogArrayIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogArrayIterator{contract: _ERC20CustodyTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogArray) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogArray)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_array", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogArray(log types.Log) (*ERC20CustodyTestLogArray, error) {
	event := new(ERC20CustodyTestLogArray)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogArray0Iterator struct {
	Event *ERC20CustodyTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogArray0)
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
		it.Event = new(ERC20CustodyTestLogArray0)
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
func (it *ERC20CustodyTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogArray0 represents a LogArray0 event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*ERC20CustodyTestLogArray0Iterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogArray0Iterator{contract: _ERC20CustodyTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogArray0)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogArray0(log types.Log) (*ERC20CustodyTestLogArray0, error) {
	event := new(ERC20CustodyTestLogArray0)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogArray1Iterator struct {
	Event *ERC20CustodyTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogArray1)
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
		it.Event = new(ERC20CustodyTestLogArray1)
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
func (it *ERC20CustodyTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogArray1 represents a LogArray1 event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*ERC20CustodyTestLogArray1Iterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogArray1Iterator{contract: _ERC20CustodyTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogArray1)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogArray1(log types.Log) (*ERC20CustodyTestLogArray1, error) {
	event := new(ERC20CustodyTestLogArray1)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogBytesIterator struct {
	Event *ERC20CustodyTestLogBytes // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogBytes)
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
		it.Event = new(ERC20CustodyTestLogBytes)
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
func (it *ERC20CustodyTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogBytes represents a LogBytes event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*ERC20CustodyTestLogBytesIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogBytesIterator{contract: _ERC20CustodyTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogBytes)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogBytes(log types.Log) (*ERC20CustodyTestLogBytes, error) {
	event := new(ERC20CustodyTestLogBytes)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogBytes32Iterator struct {
	Event *ERC20CustodyTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogBytes32)
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
		it.Event = new(ERC20CustodyTestLogBytes32)
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
func (it *ERC20CustodyTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogBytes32 represents a LogBytes32 event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*ERC20CustodyTestLogBytes32Iterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogBytes32Iterator{contract: _ERC20CustodyTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogBytes32)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogBytes32(log types.Log) (*ERC20CustodyTestLogBytes32, error) {
	event := new(ERC20CustodyTestLogBytes32)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogIntIterator struct {
	Event *ERC20CustodyTestLogInt // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogInt)
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
		it.Event = new(ERC20CustodyTestLogInt)
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
func (it *ERC20CustodyTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogInt represents a LogInt event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*ERC20CustodyTestLogIntIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogIntIterator{contract: _ERC20CustodyTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogInt) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogInt)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_int", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogInt(log types.Log) (*ERC20CustodyTestLogInt, error) {
	event := new(ERC20CustodyTestLogInt)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedAddressIterator struct {
	Event *ERC20CustodyTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogNamedAddress)
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
		it.Event = new(ERC20CustodyTestLogNamedAddress)
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
func (it *ERC20CustodyTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogNamedAddress represents a LogNamedAddress event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*ERC20CustodyTestLogNamedAddressIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogNamedAddressIterator{contract: _ERC20CustodyTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogNamedAddress)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogNamedAddress(log types.Log) (*ERC20CustodyTestLogNamedAddress, error) {
	event := new(ERC20CustodyTestLogNamedAddress)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedArrayIterator struct {
	Event *ERC20CustodyTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogNamedArray)
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
		it.Event = new(ERC20CustodyTestLogNamedArray)
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
func (it *ERC20CustodyTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogNamedArray represents a LogNamedArray event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*ERC20CustodyTestLogNamedArrayIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogNamedArrayIterator{contract: _ERC20CustodyTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogNamedArray)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogNamedArray(log types.Log) (*ERC20CustodyTestLogNamedArray, error) {
	event := new(ERC20CustodyTestLogNamedArray)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedArray0Iterator struct {
	Event *ERC20CustodyTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogNamedArray0)
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
		it.Event = new(ERC20CustodyTestLogNamedArray0)
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
func (it *ERC20CustodyTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogNamedArray0 represents a LogNamedArray0 event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*ERC20CustodyTestLogNamedArray0Iterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogNamedArray0Iterator{contract: _ERC20CustodyTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogNamedArray0)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogNamedArray0(log types.Log) (*ERC20CustodyTestLogNamedArray0, error) {
	event := new(ERC20CustodyTestLogNamedArray0)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedArray1Iterator struct {
	Event *ERC20CustodyTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogNamedArray1)
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
		it.Event = new(ERC20CustodyTestLogNamedArray1)
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
func (it *ERC20CustodyTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogNamedArray1 represents a LogNamedArray1 event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*ERC20CustodyTestLogNamedArray1Iterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogNamedArray1Iterator{contract: _ERC20CustodyTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogNamedArray1)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogNamedArray1(log types.Log) (*ERC20CustodyTestLogNamedArray1, error) {
	event := new(ERC20CustodyTestLogNamedArray1)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedBytesIterator struct {
	Event *ERC20CustodyTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogNamedBytes)
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
		it.Event = new(ERC20CustodyTestLogNamedBytes)
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
func (it *ERC20CustodyTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogNamedBytes represents a LogNamedBytes event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*ERC20CustodyTestLogNamedBytesIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogNamedBytesIterator{contract: _ERC20CustodyTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogNamedBytes)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogNamedBytes(log types.Log) (*ERC20CustodyTestLogNamedBytes, error) {
	event := new(ERC20CustodyTestLogNamedBytes)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedBytes32Iterator struct {
	Event *ERC20CustodyTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogNamedBytes32)
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
		it.Event = new(ERC20CustodyTestLogNamedBytes32)
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
func (it *ERC20CustodyTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*ERC20CustodyTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogNamedBytes32Iterator{contract: _ERC20CustodyTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogNamedBytes32)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogNamedBytes32(log types.Log) (*ERC20CustodyTestLogNamedBytes32, error) {
	event := new(ERC20CustodyTestLogNamedBytes32)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedDecimalIntIterator struct {
	Event *ERC20CustodyTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogNamedDecimalInt)
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
		it.Event = new(ERC20CustodyTestLogNamedDecimalInt)
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
func (it *ERC20CustodyTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*ERC20CustodyTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogNamedDecimalIntIterator{contract: _ERC20CustodyTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogNamedDecimalInt)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*ERC20CustodyTestLogNamedDecimalInt, error) {
	event := new(ERC20CustodyTestLogNamedDecimalInt)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedDecimalUintIterator struct {
	Event *ERC20CustodyTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogNamedDecimalUint)
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
		it.Event = new(ERC20CustodyTestLogNamedDecimalUint)
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
func (it *ERC20CustodyTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*ERC20CustodyTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogNamedDecimalUintIterator{contract: _ERC20CustodyTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogNamedDecimalUint)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*ERC20CustodyTestLogNamedDecimalUint, error) {
	event := new(ERC20CustodyTestLogNamedDecimalUint)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedIntIterator struct {
	Event *ERC20CustodyTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogNamedInt)
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
		it.Event = new(ERC20CustodyTestLogNamedInt)
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
func (it *ERC20CustodyTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogNamedInt represents a LogNamedInt event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*ERC20CustodyTestLogNamedIntIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogNamedIntIterator{contract: _ERC20CustodyTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogNamedInt)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogNamedInt(log types.Log) (*ERC20CustodyTestLogNamedInt, error) {
	event := new(ERC20CustodyTestLogNamedInt)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedStringIterator struct {
	Event *ERC20CustodyTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogNamedString)
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
		it.Event = new(ERC20CustodyTestLogNamedString)
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
func (it *ERC20CustodyTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogNamedString represents a LogNamedString event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*ERC20CustodyTestLogNamedStringIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogNamedStringIterator{contract: _ERC20CustodyTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogNamedString)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogNamedString(log types.Log) (*ERC20CustodyTestLogNamedString, error) {
	event := new(ERC20CustodyTestLogNamedString)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedUintIterator struct {
	Event *ERC20CustodyTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogNamedUint)
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
		it.Event = new(ERC20CustodyTestLogNamedUint)
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
func (it *ERC20CustodyTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogNamedUint represents a LogNamedUint event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*ERC20CustodyTestLogNamedUintIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogNamedUintIterator{contract: _ERC20CustodyTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogNamedUint)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogNamedUint(log types.Log) (*ERC20CustodyTestLogNamedUint, error) {
	event := new(ERC20CustodyTestLogNamedUint)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogStringIterator struct {
	Event *ERC20CustodyTestLogString // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogString)
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
		it.Event = new(ERC20CustodyTestLogString)
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
func (it *ERC20CustodyTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogString represents a LogString event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogString(opts *bind.FilterOpts) (*ERC20CustodyTestLogStringIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogStringIterator{contract: _ERC20CustodyTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogString) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogString)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_string", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogString(log types.Log) (*ERC20CustodyTestLogString, error) {
	event := new(ERC20CustodyTestLogString)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogUintIterator struct {
	Event *ERC20CustodyTestLogUint // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogUint)
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
		it.Event = new(ERC20CustodyTestLogUint)
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
func (it *ERC20CustodyTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogUint represents a LogUint event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*ERC20CustodyTestLogUintIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogUintIterator{contract: _ERC20CustodyTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogUint) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogUint)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogUint(log types.Log) (*ERC20CustodyTestLogUint, error) {
	event := new(ERC20CustodyTestLogUint)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogsIterator struct {
	Event *ERC20CustodyTestLogs // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestLogs)
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
		it.Event = new(ERC20CustodyTestLogs)
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
func (it *ERC20CustodyTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestLogs represents a Logs event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterLogs(opts *bind.FilterOpts) (*ERC20CustodyTestLogsIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestLogsIterator{contract: _ERC20CustodyTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestLogs) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestLogs)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "logs", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseLogs(log types.Log) (*ERC20CustodyTestLogs, error) {
	event := new(ERC20CustodyTestLogs)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
