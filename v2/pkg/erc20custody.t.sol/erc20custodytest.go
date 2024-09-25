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
	ABI: "[{\"type\":\"function\",\"name\":\"ASSET_HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_UPDATER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WHITELISTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testDepositLegacy\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositLegacyFailsIfNotSupported\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositLegacyFailsIfTokenNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20PartialThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNoParamsThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testNewCustodyFailsIfAddressesAreZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgrade\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfSenderIsNotTSSUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUnwhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUnwhitelistFailsIfSenderIsNotWhitelister\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUnwhitelistFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhitelistFailsIfSenderIsNotWhitelister\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhitelistFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LegacyMethodsNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5062012079806200003e6000396000f3fe608060405234801561001057600080fd5b50600436106103365760003560e01c806382c52992116101b2578063b5508aa9116100f9578063eb1ce7f9116100a2578063fa2a70741161007c578063fa2a7074146105cb578063fa7626d4146105d3578063fb176c12146105e0578063fe8e5f1b146105e857600080fd5b8063eb1ce7f9146105b3578063f0c8e7e0146105bb578063f4221f08146105c357600080fd5b8063cbd57e2f116100d3578063cbd57e2f1461057c578063e20c9f7114610584578063e63ab1e91461058c57600080fd5b8063b5508aa914610554578063ba414fa61461055c578063c713f8271461057457600080fd5b80639fc7fd551161015b578063a783c78911610135578063a783c7891461051d578063b0464fdc14610544578063b421ca701461054c57600080fd5b80639fc7fd5514610505578063a3f9d0e01461050d578063a4943deb1461051557600080fd5b80639158c6231161018c5780639158c623146104e0578063916a17c6146104e85780639918c1c2146104fd57600080fd5b806382c529921461049c57806385226c81146104a457806385f438c1146104b957600080fd5b80633f7286f411610281578063570618e11161022a5780636a621854116102045780636a6218541461047c5780637099d6f81461048457806371149c941461048c5780637e91c50f1461049457600080fd5b8063570618e1146104195780635d62c8601461044057806366d9a9a01461046757600080fd5b80634df42da11161025b5780634df42da11461040157806351ecdf3c1461040957806352ff59391461041157600080fd5b80633f7286f4146103e957806349c783dd146103f15780634b5838d2146103f957600080fd5b8063284cb929116102e35780633e5e3c23116102bd5780633e5e3c23146103d15780633e73ecb4146103d95780633ee92923146103e157600080fd5b8063284cb929146103ac5780632ade3880146103b45780632be6a162146103c957600080fd5b80630eee72a9116103145780630eee72a9146103875780631779672f1461038f5780631ed7831c1461039757600080fd5b8063070f2ad01461033b5780630a9254e4146103455780630d3917041461034d575b600080fd5b6103436105f0565b005b610343610810565b6103747f1ddc5e4a2f140581e5d9d5f758b2c0fbdd814c4017a02afc4b55cabcf2c1f10b81565b6040519081526020015b60405180910390f35b61034361126e565b610343611513565b61039f611664565b60405161037e919061b74b565b6103436116c6565b6103bc6119aa565b60405161037e919061b7e7565b610343611aec565b61039f611cae565b610343611d0e565b610343612283565b61039f612352565b6103436123b2565b6103436126ef565b610343612a0d565b610343612b65565b610343612d36565b6103747f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b6103747f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b981565b61046f612ecd565b60405161037e919061b94d565b61034361304f565b61034361311b565b6103436133c8565b610343613bdd565b610343613d68565b6104ac613fdc565b60405161037e919061b9eb565b6103747f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6103436140ac565b6104f061417a565b60405161037e919061ba62565b610343614275565b61034361456f565b61034361463d565b610343614c45565b6103747f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6104f0614e6d565b610343614f68565b6104ac61549d565b61056461556d565b604051901515815260200161037e565b610343615641565b6103436162e4565b61039f6163a5565b6103747f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b610343616405565b610343616511565b6103436166d3565b610343616940565b601f546105649060ff1681565b610343616c1e565b61034361729c565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561066257600080fd5b505af1158015610676573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f1ddc5e4a2f140581e5d9d5f758b2c0fbdd814c4017a02afc4b55cabcf2c1f10b60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250610761919060040161baf9565b600060405180830381600087803b15801561077b57600080fd5b505af115801561078f573d6000803e3d6000fd5b50506022546026546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063950837aa91506024015b600060405180830381600087803b1580156107f657600080fd5b505af115801561080a573d6000803e3d6000fd5b50505050565b602680547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602780548216611234179055602880549091166156781790556040516108629061b65a565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156108e7573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039290921691909117905560405161092c9061b65a565b604080825260049082018190527f7a6574610000000000000000000000000000000000000000000000000000000060608301526080602083018190528201527f5a4554410000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156109b0573d6000803e3d6000fd5b50602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260285460265492519085166024820152604481019390935292166064820152610a9f919060840160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b0000000000000000000000000000000000000000000000000000000017905261748d565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff00000000000000000000000000000000000000009091168117909155602854602654604051929391821692911690610b2b9061b668565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610b67573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556020546025546028546026546040519385169492831693918316921690610bc29061b676565b6001600160a01b039485168152928416602084015290831660408301529091166060820152608001604051809103906000f080158015610c06573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055604051610c4b9061b684565b604051809103906000f080158015610c67573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556028546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b158015610d1357600080fd5b505af1158015610d27573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015610d9d57600080fd5b505af1158015610db1573d6000803e3d6000fd5b50506020546022546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f9150602401600060405180830381600087803b158015610e1757600080fd5b505af1158015610e2b573d6000803e3d6000fd5b50506020546023546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef9150602401600060405180830381600087803b158015610e9157600080fd5b505af1158015610ea5573d6000803e3d6000fd5b5050602254602480546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639b19251a925001600060405180830381600087803b158015610f0a57600080fd5b505af1158015610f1e573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610f8057600080fd5b505af1158015610f94573d6000803e3d6000fd5b5050602480546026546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240938101939093521692506340c10f199150604401600060405180830381600087803b15801561100557600080fd5b505af1158015611019573d6000803e3d6000fd5b50506025546026546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42406024820152911692506340c10f199150604401600060405180830381600087803b15801561108857600080fd5b505af115801561109c573d6000803e3d6000fd5b5050602480546022546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526207a1209381019390935216925063a9059cbb91506044016020604051808303816000875af1158015611112573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611136919061bb0c565b506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b1580156111b757600080fd5b505af11580156111cb573d6000803e3d6000fd5b5050604080516060810182526024546001600160a01b039081168252600160208084019182528451908101855260008152938301849052825160298054925167ffffffffffffffff1674010000000000000000000000000000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090931691909316171781559093509150602a90611269908261bbf1565b505050565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc513169100000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b15801561135357600080fd5b505af1158015611367573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611452919060040161baf9565b600060405180830381600087803b15801561146c57600080fd5b505af1158015611480573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506114dd9392831692909116908790879060040161bcb0565b600060405180830381600087803b1580156114f757600080fd5b505af115801561150b573d6000803e3d6000fd5b505050505050565b6024805460275460405160009381018490526001600160a01b03928316604482015291166064820152819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602854905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b1580156115f557600080fd5b505af1158015611609573d6000803e3d6000fd5b5050604051630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401611452565b606060168054806020026020016040519081016040528092919081815260200182805480156116bc57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161169e575b5050505050905090565b6022546025546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600092919091169063d936547e90602401602060405180830381865afa158015611730573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611754919061bb0c565b90506117616000826174ac565b6022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156117d657600080fd5b505af11580156117ea573d6000803e3d6000fd5b50506025546040516001600160a01b0390911692507faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549150600090a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561187f57600080fd5b505af1158015611893573d6000803e3d6000fd5b50506022546025546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639b19251a9150602401600060405180830381600087803b1580156118f957600080fd5b505af115801561190d573d6000803e3d6000fd5b50506022546025546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063d936547e9150602401602060405180830381865afa158015611976573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061199a919061bb0c565b90506119a76001826174ac565b50565b6060601e805480602002602001604051908101604052809291908181526020016000905b82821015611ae357600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015611acc578382906000526020600020018054611a3f9061bb5d565b80601f0160208091040260200160405190810160405280929190818152602001828054611a6b9061bb5d565b8015611ab85780601f10611a8d57610100808354040283529160200191611ab8565b820191906000526020600020905b815481529060010190602001808311611a9b57829003601f168201915b505050505081526020019060010190611a20565b5050505081525050815260200190600101906119ce565b50505050905090565b60405163ca669fa760e01b81526101236004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611b3a57600080fd5b505af1158015611b4e573d6000803e3d6000fd5b50506040805161012360248201527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611c2f919060040161baf9565b600060405180830381600087803b158015611c4957600080fd5b505af1158015611c5d573d6000803e3d6000fd5b50506022546025546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639b19251a91506024016107dc565b606060188054806020026020016040519081016040528092919081815260200182805480156116bc576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161169e575050505050905090565b602480546027546040516370a0823160e01b81526001600160a01b039182166004820152620186a09360009392909216916370a082319101602060405180830381865afa158015611d63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d87919061bce7565b9050611d9481600061752e565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015611de4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e08919061bce7565b6027546040516001600160a01b0390911660248201526044810185905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391611eeb916001600160a01b039190911690600090869060040161bd00565b600060405180830381600087803b158015611f0557600080fd5b505af1158015611f19573d6000803e3d6000fd5b50506022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015611f9257600080fd5b505af1158015611fa6573d6000803e3d6000fd5b50506024546027546040518881526001600160a01b039283169450911691507fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060200160405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561204b57600080fd5b505af115801561205f573d6000803e3d6000fd5b5050602254602754602480546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b03938416600482015290831691810191909152604481018990529116925063d9caed129150606401600060405180830381600087803b1580156120da57600080fd5b505af11580156120ee573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa158015612140573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612164919061bce7565b9050612170818661752e565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa1580156121c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121e4919061bce7565b90506121f9816121f4888761bd57565b61752e565b602480546020546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015612249573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061226d919061bce7565b905061227a81600061752e565b50505050505050565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401611339565b606060178054806020026020016040519081016040528092919081815260200182805480156116bc576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161169e575050505050905090565b6022546040517feab103df000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b039091169063eab103df90602401600060405180830381600087803b15801561241157600080fd5b505af1158015612425573d6000803e3d6000fd5b5050602480546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42409381019390935216925063095ea7b391506044016020604051808303816000875af115801561249b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124bf919061bb0c565b506025546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240602482015291169063095ea7b3906044016020604051808303816000875af1158015612530573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612554919061bb0c565b50604080517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152815160058183030181526025820192839052630618f58760e51b9092527f73cba663000000000000000000000000000000000000000000000000000000006029820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906049015b600060405180830381600087803b1580156125fc57600080fd5b505af1158015612610573d6000803e3d6000fd5b505060225460275460405160609190911b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201526001600160a01b03909116925063e609055e915060340160408051601f19818403018152908290526024547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526126ba926001600160a01b03909116906103e890879060040161bd6a565b600060405180830381600087803b1580156126d457600080fd5b505af11580156126e8573d6000803e3d6000fd5b5050505050565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561275b57600080fd5b505af115801561276f573d6000803e3d6000fd5b5050602854602654604051600094508493506001600160a01b03928316929091169061279a9061b668565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f0801580156127d6573d6000803e3d6000fd5b50604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561284657600080fd5b505af115801561285a573d6000803e3d6000fd5b50506020546026546040516001600160a01b039283169450600093509116906128829061b668565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f0801580156128be573d6000803e3d6000fd5b50604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561292e57600080fd5b505af1158015612942573d6000803e3d6000fd5b50506020546028546040516001600160a01b0392831694509116915060009061296a9061b668565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f0801580156129a6573d6000803e3d6000fd5b506020546028546026546040519394506001600160a01b03928316939183169216906129d19061b668565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015611269573d6000803e3d6000fd5b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015612a7f57600080fd5b505af1158015612a93573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015612b0357600080fd5b505af1158015612b17573d6000803e3d6000fd5b50506022546040517f950837aa000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b03909116925063950837aa91506024016107dc565b6024805460275460405160019381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602854905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612c4757600080fd5b505af1158015612c5b573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015612ccb57600080fd5b505af1158015612cdf573d6000803e3d6000fd5b50506022546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321fc65f293506114dd9260009216908790879060040161bcb0565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015612da857600080fd5b505af1158015612dbc573d6000803e3d6000fd5b50506022546026546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063950837aa9150602401600060405180830381600087803b158015612e2257600080fd5b505af1158015612e36573d6000803e3d6000fd5b505050506000602260009054906101000a90046001600160a01b03166001600160a01b0316635b1125916040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e8f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612eb3919061bda4565b6026549091506119a79082906001600160a01b0316617586565b6060601b805480602002602001604051908101604052809291908181526020016000905b82821015611ae35783829060005260206000209060020201604051806040016040529081600082018054612f249061bb5d565b80601f0160208091040260200160405190810160405280929190818152602001828054612f509061bb5d565b8015612f9d5780601f10612f7257610100808354040283529160200191612f9d565b820191906000526020600020905b815481529060010190602001808311612f8057829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561303757602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411612fe45790505b50505050508152505081526020019060010190612ef1565b6024805460275460405160009381018490526001600160a01b03928316604482015291166064820152819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc513169100000000000000000000000000000000000000000000000000000000179052602854905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024016115db565b6022546040517feab103df000000000000000000000000000000000000000000000000000000008152600160048201526001600160a01b039091169063eab103df90602401600060405180830381600087803b15801561317a57600080fd5b505af115801561318e573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b1580156131f357600080fd5b505af1158015613207573d6000803e3d6000fd5b5050602480546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42409381019390935216925063095ea7b391506044016020604051808303816000875af115801561327d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132a1919061bb0c565b506025546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240602482015291169063095ea7b3906044016020604051808303816000875af1158015613312573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613336919061bb0c565b50604080517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152815160058183030181526025820192839052630618f58760e51b9092527f584a7938000000000000000000000000000000000000000000000000000000006029820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906049016125e2565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a09060009060250160408051808303601f1901815290829052602480546021546370a0823160e01b85526001600160a01b0390811660048601529294506000939216916370a082319101602060405180830381865afa15801561345b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061347f919061bce7565b905061348c81600061752e565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa1580156134dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613500919061bce7565b6020546040516001600160a01b0390911660248201526044810186905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba3916135e3916001600160a01b039190911690600090869060040161bd00565b600060405180830381600087803b1580156135fd57600080fd5b505af1158015613611573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561368a57600080fd5b505af115801561369e573d6000803e3d6000fd5b50506020546040517f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e93506136e292506001600160a01b039091169060299061bea2565b60405180910390a16020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561375f57600080fd5b505af1158015613773573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b1436906137c1908990899060299061bec4565b60405180910390a36022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561383e57600080fd5b505af1158015613852573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b906138a0908990899060299061bec4565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561390157600080fd5b505af1158015613915573d6000803e3d6000fd5b50506022546021546024546040517fc709ab6e0000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c709ab6e94506139759392831692909116908a908a9060299060040161beef565b600060405180830381600087803b15801561398f57600080fd5b505af11580156139a3573d6000803e3d6000fd5b5050602480546021546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a0823191015b602060405180830381865afa1580156139f6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a1a919061bce7565b9050613a26818761752e565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015613a76573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a9a919061bce7565b9050613aaa816121f4898761bd57565b602480546020546021546040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529082169381019390935260009291169063dd62ed3e90604401602060405180830381865afa158015613b20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b44919061bce7565b9050613b5181600061752e565b602480546020546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015613ba1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613bc5919061bce7565b9050613bd281600061752e565b505050505050505050565b6040517f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260019060009060250160408051808303601f190181529082905260285463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613c7657600080fd5b505af1158015613c8a573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015613cfa57600080fd5b505af1158015613d0e573d6000803e3d6000fd5b50506022546024546040517fc709ab6e0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c709ab6e93506114dd9260009216908790879060299060040161beef565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015613dda57600080fd5b505af1158015613dee573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b158015613e5357600080fd5b505af1158015613e67573d6000803e3d6000fd5b5050604051630618f58760e51b81527f584a7938000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015613ed757600080fd5b505af1158015613eeb573d6000803e3d6000fd5b5050602254602754602480546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b03938416600482015290831691810191909152600160448201529116925063d9caed129150606401600060405180830381600087803b158015613f6657600080fd5b505af1158015613f7a573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156107f657600080fd5b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015611ae357838290600052602060002001805461401f9061bb5d565b80601f016020809104026020016040519081016040528092919081815260200182805461404b9061bb5d565b80156140985780601f1061406d57610100808354040283529160200191614098565b820191906000526020600020905b81548152906001019060200180831161407b57829003601f168201915b505050505081526020019060010190614000565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561411857600080fd5b505af115801561412c573d6000803e3d6000fd5b50506022546040517f9a590427000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b039091169250639a59042791506024016107dc565b6060601d805480602002602001604051908101604052809291908181526020016000905b82821015611ae35760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561425d57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001906004019060208260030104928301926001038202915080841161420a5790505b5050505050815250508152602001906001019061419e565b602480546027546040516001938101939093526001600160a01b03918216604484015216606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905260285490517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561437057600080fd5b505af1158015614384573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b1580156143e957600080fd5b505af11580156143fd573d6000803e3d6000fd5b5050604051630618f58760e51b81527f584a7938000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561446d57600080fd5b505af1158015614481573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506144df939283169290911690600190879060040161bcb0565b600060405180830381600087803b1580156144f957600080fd5b505af115801561450d573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156126d457600080fd5b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156145db57600080fd5b505af11580156145ef573d6000803e3d6000fd5b50506022546040517f9b19251a000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b039091169250639b19251a91506024016107dc565b604080516004808252602480830184526020830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed7016900000000000000000000000000000000000000000000000000000000179052805460275494516370a0823160e01b81526001600160a01b0395861693810193909352620186a0946000939116916370a082319101602060405180830381865afa1580156146e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061470b919061bce7565b905061471881600061752e565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015614768573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061478c919061bce7565b6020546040516001600160a01b0390911660248201526044810186905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161486f916001600160a01b039190911690600090869060040161bd00565b600060405180830381600087803b15801561488957600080fd5b505af115801561489d573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561491657600080fd5b505af115801561492a573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0935001905060405180910390a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156149e357600080fd5b505af11580156149f7573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d590614a42908990899061bf44565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614aa357600080fd5b505af1158015614ab7573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f29450614b149392831692909116908a908a9060040161bcb0565b600060405180830381600087803b158015614b2e57600080fd5b505af1158015614b42573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa158015614b94573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614bb8919061bce7565b9050614bc581600061752e565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015614c15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614c39919061bce7565b9050613aaa818561752e565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a09060009060250160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614ce057600080fd5b505af1158015614cf4573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614ddf919060040161baf9565b600060405180830381600087803b158015614df957600080fd5b505af1158015614e0d573d6000803e3d6000fd5b50506022546021546024546040517fc709ab6e0000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c709ab6e94506114dd9392831692909116908790879060299060040161beef565b6060601c805480602002602001604051908101604052809291908181526020016000905b82821015611ae35760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015614f5057602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411614efd5790505b50505050508152505081526020019060010190614e91565b6022546040517feab103df000000000000000000000000000000000000000000000000000000008152600160048201526103e8916001600160a01b03169063eab103df90602401600060405180830381600087803b158015614fc957600080fd5b505af1158015614fdd573d6000803e3d6000fd5b5050602480546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42409381019390935216925063095ea7b391506044016020604051808303816000875af1158015615053573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615077919061bb0c565b506025546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240602482015291169063095ea7b3906044016020604051808303816000875af11580156150e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061510c919061bb0c565b50602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa15801561515d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615181919061bce7565b6025546028546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa1580156151d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906151f7919061bce7565b9050600060405160200161522e907f68656c6c6f000000000000000000000000000000000000000000000000000000815260050190565b60408051808303601f19018152908290526022546381bad6f360e01b8352600160048401819052602484018190526044840181905260648401526001600160a01b031660848301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156152b157600080fd5b505af11580156152c5573d6000803e3d6000fd5b505060245460275460405160609190911b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201526001600160a01b0390911692507f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae915060340160408051601f198184030181529082905261534e918890869061bf5d565b60405180910390a26022546027546040805160609290921b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208301528051808303601401815260348301918290526024547fe609055e000000000000000000000000000000000000000000000000000000009092526001600160a01b039384169363e609055e936153ed9391909116908990879060380161bd6a565b600060405180830381600087803b15801561540757600080fd5b505af115801561541b573d6000803e3d6000fd5b5050505061080a848461542e919061bf88565b602480546022546040516370a0823160e01b81526001600160a01b0391821660048201529116916370a082319101602060405180830381865afa158015615479573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121f4919061bce7565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015611ae35783829060005260206000200180546154e09061bb5d565b80601f016020809104026020016040519081016040528092919081815260200182805461550c9061bb5d565b80156155595780601f1061552e57610100808354040283529160200191615559565b820191906000526020600020905b81548152906001019060200180831161553c57829003601f168201915b5050505050815260200190600101906154c1565b60085460009060ff1615615585575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015615616573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061563a919061bce7565b1415905090565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561569a57600080fd5b505af11580156156ae573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250615799919060040161baf9565b600060405180830381600087803b1580156157b357600080fd5b505af11580156157c7573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561581b57600080fd5b505af115801561582f573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561588c57600080fd5b505af11580156158a0573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061598b919060040161baf9565b600060405180830381600087803b1580156159a557600080fd5b505af11580156159b9573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015615a0d57600080fd5b505af1158015615a21573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015615a7e57600080fd5b505af1158015615a92573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b158015615ae657600080fd5b505af1158015615afa573d6000803e3d6000fd5b505060248054602754604051620186a09381018490526001600160a01b039283166044820152911660648201529092506000915060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905251630618f58760e51b81527fd93c0665000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015615bf657600080fd5b505af1158015615c0a573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015615c6757600080fd5b505af1158015615c7b573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f29450615cd89392831692909116908790879060040161bcb0565b600060405180830381600087803b158015615cf257600080fd5b505af1158015615d06573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015615d6357600080fd5b505af1158015615d77573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015615dcb57600080fd5b505af1158015615ddf573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a0823191015b602060405180830381865afa158015615e32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615e56919061bce7565b9050615e6381600061752e565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015615eb3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615ed7919061bce7565b6020546040516001600160a01b0390911660248201526044810186905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391615fba916001600160a01b039190911690600090869060040161bd00565b600060405180830381600087803b158015615fd457600080fd5b505af1158015615fe8573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561606157600080fd5b505af1158015616075573d6000803e3d6000fd5b505060208054602454602754604080516001600160a01b0394851681529485018c905291831684830152919091166060830152517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609350908190036080019150a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561614b57600080fd5b505af115801561615f573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5906161aa908990899061bf44565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561620b57600080fd5b505af115801561621f573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f2945061627c9392831692909116908a908a9060040161bcb0565b600060405180830381600087803b15801561629657600080fd5b505af11580156162aa573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a0823191016139d9565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a2000000000000000000000000000000000000000000000000000000001790526024805460275492516370a0823160e01b81526001600160a01b0393841660048201529394506000939216916370a082319101615e15565b606060158054806020026020016040519081016040528092919081815260200182805480156116bc576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161169e575050505050905090565b60008060405160200161643b907f68656c6c6f000000000000000000000000000000000000000000000000000000815260050190565b60408051808303601f190181529082905260285463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156164a257600080fd5b505af11580156164b6573d6000803e3d6000fd5b5050604051630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401614ddf565b60405163ca669fa760e01b81526101236004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561655f57600080fd5b505af1158015616573573d6000803e3d6000fd5b50506040805161012360248201527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250616654919060040161baf9565b600060405180830381600087803b15801561666e57600080fd5b505af1158015616682573d6000803e3d6000fd5b50506022546025546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a59042791506024016107dc565b602480546027546040516001938101939093526001600160a01b03918216604484015216606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905260285490517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156167ce57600080fd5b505af11580156167e2573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b15801561684757600080fd5b505af115801561685b573d6000803e3d6000fd5b5050604051630618f58760e51b81527f584a7938000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156168cb57600080fd5b505af11580156168df573d6000803e3d6000fd5b50506022546021546024546040517fc709ab6e0000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c709ab6e94506144df939283169290911690600190879060299060040161beef565b602254602480546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600093919091169163d936547e9101602060405180830381865afa1580156169a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906169cd919061bb0c565b90506169da6001826174ac565b6022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015616a4f57600080fd5b505af1158015616a63573d6000803e3d6000fd5b50506024546040516001600160a01b0390911692507f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919150600090a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015616af857600080fd5b505af1158015616b0c573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b158015616b7157600080fd5b505af1158015616b85573d6000803e3d6000fd5b5050602254602480546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529216935063d936547e925001602060405180830381865afa158015616bed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616c11919061bb0c565b90506119a76000826174ac565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc5131691000000000000000000000000000000000000000000000000000000001790526024805460275492516370a0823160e01b81526001600160a01b0393841660048201529394506000939216916370a082319101602060405180830381865afa158015616cf7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616d1b919061bce7565b9050616d2881600061752e565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015616d78573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616d9c919061bce7565b6020546040516001600160a01b0390911660248201526044810186905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391616e7f916001600160a01b039190911690600090869060040161bd00565b600060405180830381600087803b158015616e9957600080fd5b505af1158015616ead573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015616f2657600080fd5b505af1158015616f3a573d6000803e3d6000fd5b50506020547f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af6092506001600160a01b03169050616f7860028861bf9b565b602454602754604080516001600160a01b03958616815260208101949094529184168383015292909216606082015290519081900360800190a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561702757600080fd5b505af115801561703b573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d590617086908990899061bf44565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156170e757600080fd5b505af11580156170fb573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506171589392831692909116908a908a9060040161bcb0565b600060405180830381600087803b15801561717257600080fd5b505af1158015617186573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa1580156171d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906171fc919061bce7565b905061720d816121f460028961bf9b565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa15801561725d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617281919061bce7565b9050613aaa8161729260028a61bf9b565b6121f4908761bd57565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156172fa57600080fd5b505af115801561730e573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506173f9919060040161baf9565b600060405180830381600087803b15801561741357600080fd5b505af1158015617427573d6000803e3d6000fd5b5050602254602754602480546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b03938416600482015290831691810191909152604481018690529116925063d9caed1291506064016126ba565b600061749761b692565b6174a28484836175e7565b9150505b92915050565b6040517ff7fe347700000000000000000000000000000000000000000000000000000000815282151560048201528115156024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f7fe3477906044015b60006040518083038186803b15801561751a57600080fd5b505afa15801561150b573d6000803e3d6000fd5b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c5490604401617502565b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f690604401617502565b6000806175f48584617662565b90506176576040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f7879000000815250828660405160200161764292919061bfd6565b6040516020818303038152906040528561766e565b9150505b9392505050565b600061765b838361769c565b60c081015151600090156176925761768b84848460c001516176b7565b905061765b565b61768b848461785d565b60006176a88383617948565b61765b8383602001518461766e565b6000806176c2617958565b905060006176d08683617a2b565b905060006176e78260600151836020015185617ed1565b905060006176f7838389896180e3565b9050600061770482618f60565b602081015181519192509060030b156177775789826040015160405160200161772e92919061bff8565b60408051601f19818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261776e9160040161baf9565b60405180910390fd5b60006177ba6040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a20000000000000000000000081525083600161912f565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d9061780d90849060040161baf9565b602060405180830381865afa15801561782a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061784e919061bda4565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc925906178b290879060040161baf9565b600060405180830381865afa1580156178cf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526178f7919081019061c132565b90506000617925828560405160200161791192919061c167565b60405160208183030381529060405261932f565b90506001600160a01b0381166174a257848460405160200161772e92919061c196565b61795482826000619342565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c906179df90849060040161c241565b600060405180830381865afa1580156179fc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617a24919081019061c288565b9250505090565b617a5d6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d9050617aa86040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b617ab185619445565b60208201526000617ac18661982a565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa158015617b03573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617b2b919081019061c288565b86838560200151604051602001617b45949392919061c2d1565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb1190617b9d90859060040161baf9565b600060405180830381865afa158015617bba573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617be2919081019061c288565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f690617c2a90849060040161c3d5565b602060405180830381865afa158015617c47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617c6b919061bb0c565b617c80578160405160200161772e919061c427565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890617cc590849060040161c4b9565b600060405180830381865afa158015617ce2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617d0a919081019061c288565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f690617d5190849060040161c50b565b602060405180830381865afa158015617d6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617d92919061bb0c565b15617e27576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890617ddc90849060040161c50b565b600060405180830381865afa158015617df9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617e21919081019061c288565b60408501525b846001600160a01b03166349c4fac8828660000151604051602001617e4c919061c55d565b6040516020818303038152906040526040518363ffffffff1660e01b8152600401617e7892919061c5c9565b600060405180830381865afa158015617e95573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617ebd919081019061c288565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b6060815260200190600190039081617eed5790505090506040518060400160405280600481526020017f677265700000000000000000000000000000000000000000000000000000000081525081600081518110617f4d57617f4d61c5ee565b60200260200101819052506040518060400160405280600381526020017f2d726c000000000000000000000000000000000000000000000000000000000081525081600181518110617fa157617fa161c5ee565b602002602001018190525084604051602001617fbd919061c61d565b60405160208183030381529060405281600281518110617fdf57617fdf61c5ee565b602002602001018190525082604051602001617ffb919061c689565b6040516020818303038152906040528160038151811061801d5761801d61c5ee565b6020026020010181905250600061803382618f60565b602080820151604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000081850190815282518084018452600080825290860152825180840190935290518252928101929092529192506180c49060408051808201825260008082526020918201528151808301909252845182528085019082015290619aad565b6180d9578560405160200161772e919061c6ca565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015618133565b511590565b6182a7578260200151156181ef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a40161776e565b8260c00151156182a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a40161776e565b6040805160ff8082526120008201909252600091816020015b60608152602001906001900390816182c057905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061831b9061c75b565b935060ff16815181106183305761833061c5ee565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e3700000000000000000000000000000000000000815250604051602001618381919061c77a565b60405160208183030381529060405282828061839c9061c75b565b935060ff16815181106183b1576183b161c5ee565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806183fe9061c75b565b935060ff16815181106184135761841361c5ee565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806184609061c75b565b935060ff16815181106184755761847561c5ee565b602002602001018190525087602001518282806184919061c75b565b935060ff16815181106184a6576184a661c5ee565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e7472616374506174680000000000000000000000000000000000008152508282806184f39061c75b565b935060ff16815181106185085761850861c5ee565b6020908102919091010152875182826185208161c75b565b935060ff16815181106185355761853561c5ee565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e496400000000000000000000000000000000000000000000008152508282806185829061c75b565b935060ff16815181106185975761859761c5ee565b60200260200101819052506185ab46619b0e565b82826185b68161c75b565b935060ff16815181106185cb576185cb61c5ee565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806186189061c75b565b935060ff168151811061862d5761862d61c5ee565b6020026020010181905250868282806186459061c75b565b935060ff168151811061865a5761865a61c5ee565b60209081029190910101528551156187815760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826186ab8161c75b565b935060ff16815181106186c0576186c061c5ee565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d9061871090899060040161baf9565b600060405180830381865afa15801561872d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618755919081019061c288565b82826187608161c75b565b935060ff16815181106187755761877561c5ee565b60200260200101819052505b8460200151156188515760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826187ca8161c75b565b935060ff16815181106187df576187df61c5ee565b60200260200101819052506040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525082828061882c9061c75b565b935060ff16815181106188415761884161c5ee565b6020026020010181905250618a18565b61888961812e8660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b61891c5760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826188cc8161c75b565b935060ff16815181106188e1576188e161c5ee565b60200260200101819052508460a00151604051602001618901919061c61d565b60405160208183030381529060405282828061882c9061c75b565b8460c0015115801561895f57506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261895d90511590565b155b15618a185760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826189a38161c75b565b935060ff16815181106189b8576189b861c5ee565b60200260200101819052506189cc88619bae565b6040516020016189dc919061c61d565b6040516020818303038152906040528282806189f79061c75b565b935060ff1681518110618a0c57618a0c61c5ee565b60200260200101819052505b60408086015181518083018352600080825260209182015282518084019093528151835290810190820152618a4c90511590565b618ae15760408051808201909152600b81527f2d2d72656c61796572496400000000000000000000000000000000000000000060208201528282618a8f8161c75b565b935060ff1681518110618aa457618aa461c5ee565b60200260200101819052508460400151828280618ac09061c75b565b935060ff1681518110618ad557618ad561c5ee565b60200260200101819052505b606085015115618c025760408051808201909152600681527f2d2d73616c74000000000000000000000000000000000000000000000000000060208201528282618b2a8161c75b565b935060ff1681518110618b3f57618b3f61c5ee565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015618bae573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618bd6919081019061c288565b8282618be18161c75b565b935060ff1681518110618bf657618bf661c5ee565b60200260200101819052505b60e08501515115618ca95760408051808201909152600a81527f2d2d6761734c696d69740000000000000000000000000000000000000000000060208201528282618c4c8161c75b565b935060ff1681518110618c6157618c6161c5ee565b6020026020010181905250618c7d8560e0015160000151619b0e565b8282618c888161c75b565b935060ff1681518110618c9d57618c9d61c5ee565b60200260200101819052505b60e08501516020015115618d535760408051808201909152600a81527f2d2d67617350726963650000000000000000000000000000000000000000000060208201528282618cf68161c75b565b935060ff1681518110618d0b57618d0b61c5ee565b6020026020010181905250618d278560e0015160200151619b0e565b8282618d328161c75b565b935060ff1681518110618d4757618d4761c5ee565b60200260200101819052505b60e08501516040015115618dfd5760408051808201909152600e81527f2d2d6d617846656550657247617300000000000000000000000000000000000060208201528282618da08161c75b565b935060ff1681518110618db557618db561c5ee565b6020026020010181905250618dd18560e0015160400151619b0e565b8282618ddc8161c75b565b935060ff1681518110618df157618df161c5ee565b60200260200101819052505b60e08501516060015115618ea75760408051808201909152601681527f2d2d6d61785072696f726974794665655065724761730000000000000000000060208201528282618e4a8161c75b565b935060ff1681518110618e5f57618e5f61c5ee565b6020026020010181905250618e7b8560e0015160600151619b0e565b8282618e868161c75b565b935060ff1681518110618e9b57618e9b61c5ee565b60200260200101819052505b60008160ff1667ffffffffffffffff811115618ec557618ec561bb2e565b604051908082528060200260200182016040528015618ef857816020015b6060815260200190600190039081618ee35790505b50905060005b8260ff168160ff161015618f5157838160ff1681518110618f2157618f2161c5ee565b6020026020010151828260ff1681518110618f3e57618f3e61c5ee565b6020908102919091010152600101618efe565b5093505050505b949350505050565b618f876040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c9161900d9186910161c7e5565b600060405180830381865afa15801561902a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619052919081019061c288565b90506000619060868361a69d565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b8152600401619090919061b9eb565b6000604051808303816000875af11580156190af573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526190d7919081019061c82c565b805190915060030b158015906190f05750602081015151155b80156190ff5750604081015151155b156180d957816000815181106191175761911761c5ee565b602002602001015160405160200161772e919061c8e2565b606060006191648560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252865182528087019082015290915061919b9082905b9061a7f2565b156192f8576000619218826192128461920c6191de8a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061a819565b9061a87b565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061927c90829061a7f2565b156192e657604080518082018252600181527f0a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526192e3905b829061a900565b90505b6192ef8161a926565b9250505061765b565b821561931157848460405160200161772e92919061cace565b505060408051602081019091526000815261765b565b509392505050565b6000808251602084016000f09392505050565b8160a001511561935157505050565b600061935e84848461a98f565b9050600061936b82618f60565b602081015181519192509060030b1580156194075750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261940790604080518082018252600080825260209182015281518083019092528451825280850190820152619195565b1561941457505050505050565b6040820151511561943457816040015160405160200161772e919061cb75565b8060405160200161772e919061cbd3565b6060600061947a8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506194df905b8290619aad565b1561954e57604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261765b9061954990839061af2a565b61a926565b604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526195b0905b829061afb4565b60010361967d57604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619616906192dc565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261765b90619549905b839061a900565b604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526196dc906194d8565b1561981357604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061974490839061b04e565b905060008160018351619757919061bd57565b815181106197675761976761c5ee565b6020026020010151905061980a6195496197dd6040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925285518252808601908201529061af2a565b95945050505050565b8260405160200161772e919061cc3e565b50919050565b6060600061985f8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506198c1906194d8565b156198cf5761765b8161a926565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261992e906195a9565b60010361999857604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261765b9061954990619676565b604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526199f7906194d8565b1561981357604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290619a5f90839061b04e565b9050600181511115619a9b578060028251619a7a919061bd57565b81518110619a8a57619a8a61c5ee565b602002602001015192505050919050565b508260405160200161772e919061cc3e565b805182516000911115619ac2575060006174a6565b81518351602085015160009291619ad89161bf88565b619ae2919061bd57565b905082602001518103619af95760019150506174a6565b82516020840151819020912014905092915050565b60606000619b1b8361b0f3565b600101905060008167ffffffffffffffff811115619b3b57619b3b61bb2e565b6040519080825280601f01601f191660200182016040528015619b65576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084619b6f57509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e5345440000000000000000000000000000000000000000000081840190815285518087018752838152840192909252845180860190955251845290830152606091619c3a905b829061b1d5565b15619c7a57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e7365000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619cd990619c33565b15619d1957505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d4954000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619d7890619c33565b15619db857505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c79000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619e1790619c33565b80619e7c5750604080518082018252601081527f47504c2d322e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619e7c90619c33565b15619ebc57505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c79000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619f1b90619c33565b80619f805750604080518082018252601081527f47504c2d332e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619f8090619c33565b15619fc057505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a01f90619c33565b8061a0845750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a08490619c33565b1561a0c457505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a12390619c33565b8061a1885750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a18890619c33565b1561a1c857505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a22790619c33565b1561a26757505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a2c690619c33565b1561a30657505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a36590619c33565b1561a3a557505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a40490619c33565b1561a44457505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a4a390619c33565b1561a4e357505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a54290619c33565b8061a5a75750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a5a790619c33565b1561a5e757505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a64690619c33565b1561a68657505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b6040808401518451915161772e929060200161cd1c565b60608060005b845181101561a728578185828151811061a6bf5761a6bf61c5ee565b602002602001015160405160200161a6d892919061c167565b60405160208183030381529060405291506001855161a6f7919061bd57565b811461a720578160405160200161a70e919061ce85565b60405160208183030381529060405291505b60010161a6a3565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161a741579050509050838160008151811061a76c5761a76c61c5ee565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061a7c05761a7c061c5ee565b6020026020010181905250818160028151811061a7df5761a7df61c5ee565b6020908102919091010152949350505050565b602080830151835183519284015160009361a810929184919061b1e9565b14159392505050565b6040805180820190915260008082526020820152600061a84b846000015185602001518560000151866020015161b2fa565b905083602001518161a85d919061bd57565b8451859061a86c90839061bd57565b90525060208401525090919050565b604080518082019091526000808252602082015281518351101561a8a05750816174a6565b602080830151908401516001911461a8c75750815160208481015190840151829020919020145b801561a8f85782518451859061a8de90839061bd57565b905250825160208501805161a8f490839061bf88565b9052505b509192915050565b604080518082019091526000808252602082015261a91f83838361b41a565b5092915050565b60606000826000015167ffffffffffffffff81111561a9475761a94761bb2e565b6040519080825280601f01601f19166020018201604052801561a971576020820181803683370190505b509050600060208201905061a91f818560200151866000015161b4c5565b6060600061a99b617958565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161a9b857905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061aa139061c75b565b935060ff168151811061aa285761aa2861c5ee565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161aa79919061cec6565b60405160208183030381529060405282828061aa949061c75b565b935060ff168151811061aaa95761aaa961c5ee565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061aaf69061c75b565b935060ff168151811061ab0b5761ab0b61c5ee565b60200260200101819052508260405160200161ab27919061c689565b60405160208183030381529060405282828061ab429061c75b565b935060ff168151811061ab575761ab5761c5ee565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061aba49061c75b565b935060ff168151811061abb95761abb961c5ee565b602002602001018190525061abce878461b53f565b828261abd98161c75b565b935060ff168151811061abee5761abee61c5ee565b60209081029190910101528551511561ac9a5760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261ac408161c75b565b935060ff168151811061ac555761ac5561c5ee565b602002602001018190525061ac6e86600001518461b53f565b828261ac798161c75b565b935060ff168151811061ac8e5761ac8e61c5ee565b60200260200101819052505b85608001511561ad085760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261ace38161c75b565b935060ff168151811061acf85761acf861c5ee565b602002602001018190525061ad6e565b841561ad6e5760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261ad4d8161c75b565b935060ff168151811061ad625761ad6261c5ee565b60200260200101819052505b6040860151511561ae0a5760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261adb88161c75b565b935060ff168151811061adcd5761adcd61c5ee565b6020026020010181905250856040015182828061ade99061c75b565b935060ff168151811061adfe5761adfe61c5ee565b60200260200101819052505b85606001511561ae745760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261ae538161c75b565b935060ff168151811061ae685761ae6861c5ee565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561ae925761ae9261bb2e565b60405190808252806020026020018201604052801561aec557816020015b606081526020019060019003908161aeb05790505b50905060005b8260ff168160ff16101561af1e57838160ff168151811061aeee5761aeee61c5ee565b6020026020010151828260ff168151811061af0b5761af0b61c5ee565b602090810291909101015260010161aecb565b50979650505050505050565b604080518082019091526000808252602082015281518351101561af4f5750816174a6565b8151835160208501516000929161af659161bf88565b61af6f919061bd57565b6020840151909150600190821461af90575082516020840151819020908220145b801561afab5783518551869061afa790839061bd57565b9052505b50929392505050565b600080826000015161afd8856000015186602001518660000151876020015161b2fa565b61afe2919061bf88565b90505b8351602085015161aff6919061bf88565b811161a91f578161b0068161cf0b565b925050826000015161b03d85602001518361b021919061bd57565b865161b02d919061bd57565b838660000151876020015161b2fa565b61b047919061bf88565b905061afe5565b6060600061b05c848461afb4565b61b06790600161bf88565b67ffffffffffffffff81111561b07f5761b07f61bb2e565b60405190808252806020026020018201604052801561b0b257816020015b606081526020019060019003908161b09d5790505b50905060005b81518110156193275761b0ce619549868661a900565b82828151811061b0e05761b0e061c5ee565b602090810291909101015260010161b0b8565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061b13c577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061b168576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061b18657662386f26fc10000830492506010015b6305f5e100831061b19e576305f5e100830492506008015b612710831061b1b257612710830492506004015b6064831061b1c4576064830492506002015b600a83106174a65760010192915050565b600061b1e1838361b57f565b159392505050565b60008085841161b2f0576020841161b29c576000841561b23457600161b21086602061bd57565b61b21b90600861cf25565b61b22690600261d023565b61b230919061bd57565b1990505b835181168561b243898961bf88565b61b24d919061bd57565b805190935082165b81811461b2875787841161b26f5787945050505050618f58565b8361b2798161d02f565b94505082845116905061b255565b61b291878561bf88565b945050505050618f58565b83832061b2a9858861bd57565b61b2b3908761bf88565b91505b85821061b2ee5784822080820361b2db5761b2d1868461bf88565b9350505050618f58565b61b2e660018461bd57565b92505061b2b6565b505b5092949350505050565b6000838186851161b405576020851161b3b4576000851561b34657600161b32287602061bd57565b61b32d90600861cf25565b61b33890600261d023565b61b342919061bd57565b1990505b8451811660008761b3578b8b61bf88565b61b361919061bd57565b855190915083165b82811461b3a65781861061b38e5761b3818b8b61bf88565b9650505050505050618f58565b8561b3988161cf0b565b96505083865116905061b369565b859650505050505050618f58565b508383206000905b61b3c6868961bd57565b821161b4035785832080820361b3e25783945050505050618f58565b61b3ed60018561bf88565b935050818061b3fb9061cf0b565b92505061b3bc565b505b61b40f878761bf88565b979650505050505050565b6040805180820190915260008082526020820152600061b44c856000015186602001518660000151876020015161b2fa565b60208087018051918601919091525190915061b468908261bd57565b83528451602086015161b47b919061bf88565b810361b48a576000855261b4bc565b8351835161b498919061bf88565b8551869061b4a790839061bd57565b905250835161b4b6908261bf88565b60208601525b50909392505050565b6020811061b4fd578151835261b4dc60208461bf88565b925061b4e960208361bf88565b915061b4f660208261bd57565b905061b4c5565b600019811561b52c57600161b51383602061bd57565b61b51f9061010061d023565b61b529919061bd57565b90505b9151835183169219169190911790915250565b6060600061b54d8484617a2b565b805160208083015160405193945061b5679390910161d046565b60405160208183030381529060405291505092915050565b815181516000919081111561b592575081515b6020808501519084015160005b8381101561b64b578251825180821461b61b57600019602087101561b5fa5760018461b5cc89602061bd57565b61b5d6919061bf88565b61b5e190600861cf25565b61b5ec90600261d023565b61b5f6919061bd57565b1990505b818116838216818103911461b6185797506174a69650505050505050565b50505b61b62660208661bf88565b945061b63360208561bf88565b9350505060208161b644919061bf88565b905061b59f565b50845186516180d9919061d09e565b610c9f806200d0bf83390190565b611d83806200dd5e83390190565b611783806200fae183390190565b610de0806201126483390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161b6d561b6da565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161b6d56040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561b78c5783516001600160a01b031683526020938401939092019160010161b765565b509095945050505050565b60005b8381101561b7b257818101518382015260200161b79a565b50506000910152565b6000815180845261b7d381602086016020860161b797565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b8e3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561b8c9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261b8b384865161b7bb565b602095860195909450929092019160010161b879565b50919750505060209485019492909201915060010161b80f565b50929695505050505050565b600081518084526020840193506020830160005b8281101561b9435781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161b903565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b8e3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261b9b9604088018261b7bb565b905060208201519150868103602088015261b9d4818361b8ef565b96505050602093840193919091019060010161b975565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b8e3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261ba4d85835161b7bb565b9450602093840193919091019060010161ba13565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b8e3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261bae3604087018261b8ef565b955050602093840193919091019060010161ba8a565b60208152600061765b602083018461b7bb565b60006020828403121561bb1e57600080fd5b8151801515811461765b57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c9082168061bb7157607f821691505b602082108103619824577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f82111561126957806000526020600020601f840160051c8101602085101561bbd15750805b601f840160051c820191505b818110156126e8576000815560010161bbdd565b815167ffffffffffffffff81111561bc0b5761bc0b61bb2e565b61bc1f8161bc19845461bb5d565b8461bbaa565b6020601f82116001811461bc53576000831561bc3b5750848201515b600019600385901b1c1916600184901b1784556126e8565b600084815260208120601f198516915b8281101561bc83578785015182556020948501946001909201910161bc63565b508482101561bca15786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6001600160a01b03851681526001600160a01b03841660208201528260408201526080606082015260006180d9608083018461b7bb565b60006020828403121561bcf957600080fd5b5051919050565b6001600160a01b038416815282602082015260606040820152600061980a606083018461b7bb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156174a6576174a661bd28565b60808152600061bd7d608083018761b7bb565b6001600160a01b0386166020840152846040840152828103606084015261b40f818561b7bb565b60006020828403121561bdb657600080fd5b81516001600160a01b038116811461765b57600080fd5b600081546001600160a01b038116845267ffffffffffffffff8160a01c1660208501525060018201606060408501526000815461be098161bb5d565b806060880152600182166000811461be28576001811461be625761be96565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083166080890152608082151560051b890101935061be96565b84600052602060002060005b8381101561be8d5781548a82016080015260019091019060200161be6e565b89016080019450505b50919695505050505050565b6001600160a01b0383168152604060208201526000618f58604083018461bdcd565b83815260606020820152600061bedd606083018561b7bb565b82810360408401526180d9818561bdcd565b6001600160a01b03861681526001600160a01b038516602082015283604082015260a06060820152600061bf2660a083018561b7bb565b828103608084015261bf38818561bdcd565b98975050505050505050565b828152604060208201526000618f58604083018461b7bb565b60608152600061bf70606083018661b7bb565b84602084015282810360408401526180d9818561b7bb565b808201808211156174a6576174a661bd28565b60008261bfd1577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6001600160a01b0383168152604060208201526000618f58604083018461b7bb565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161c03081601a85016020880161b797565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161c06d81601c84016020880161b797565b01601c01949350505050565b6040516060810167ffffffffffffffff8111828210171561c09c5761c09c61bb2e565b60405290565b60008067ffffffffffffffff84111561c0bd5761c0bd61bb2e565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561c0ec5761c0ec61bb2e565b60405283815290508082840185101561c10457600080fd5b61932784602083018561b797565b600082601f83011261c12357600080fd5b61765b8383516020850161c0a2565b60006020828403121561c14457600080fd5b815167ffffffffffffffff81111561c15b57600080fd5b6174a28482850161c112565b6000835161c17981846020880161b797565b83519083019061c18d81836020880161b797565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161c1ce81601a85016020880161b797565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161c20b81603384016020880161b797565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f5554000000000000000000000000000000000000000000606082015260806020820152600061765b608083018461b7bb565b60006020828403121561c29a57600080fd5b815167ffffffffffffffff81111561c2b157600080fd5b8201601f8101841361c2c257600080fd5b6174a28482516020840161c0a2565b6000855161c2e3818460208a0161b797565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161c31d816001840160208a0161b797565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161c35b81600284016020890161b797565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161c39d81600284016020880161b797565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061c3e8604083018461b7bb565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161c45f81601f85016020870161b797565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061c4cc604083018461b7bb565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061c51e604083018461b7bb565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161c59581601485016020870161b797565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061c5dc604083018561b7bb565b8281036020840152617657818561b7bb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161c65581600185016020870161b797565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161c69b81846020870161b797565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161c74e81604b85016020870161b797565b91909101604b0192915050565b600060ff821660ff810361c7715761c77161bd28565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161c7d881602985016020870161b797565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f5041544800000000000000000000606082015260806020820152600061765b608083018461b7bb565b60006020828403121561c83e57600080fd5b815167ffffffffffffffff81111561c85557600080fd5b82016060818503121561c86757600080fd5b61c86f61c079565b81518060030b811461c88057600080fd5b8152602082015167ffffffffffffffff81111561c89c57600080fd5b61c8a88682850161c112565b602083015250604082015167ffffffffffffffff81111561c8c857600080fd5b61c8d48682850161c112565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161c94081602185016020870161b797565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161cb2c81602185016020880161b797565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161cb6981602e84016020880161b797565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161c7d881602985016020870161b797565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161cc3181602285016020870161b797565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161cc7681600e85016020870161b797565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161cd5481601885016020880161b797565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161cd9181601c84016020880161b797565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161ce9781846020870161b797565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161cefe81601c85016020870161b797565b91909101601c0192915050565b6000600019820361cf1e5761cf1e61bd28565b5060010190565b80820281158282048414176174a6576174a661bd28565b6001815b600184111561cf775780850481111561cf5b5761cf5b61bd28565b600184161561cf6957908102905b60019390931c92800261cf40565b935093915050565b60008261cf8e575060016174a6565b8161cf9b575060006174a6565b816001811461cfb1576002811461cfbb5761cfd7565b60019150506174a6565b60ff84111561cfcc5761cfcc61bd28565b50506001821b6174a6565b5060208310610133831016604e8410600b841016171561cffa575081810a6174a6565b61d007600019848461cf3c565b806000190482111561d01b5761d01b61bd28565b029392505050565b600061765b838361cf7f565b60008161d03e5761d03e61bd28565b506000190190565b6000835161d05881846020880161b797565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161d09281600184016020880161b797565b01600101949350505050565b818103600083128015838313168383128216171561a91f5761a91f61bd2856fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220a043c41353215fce25ecb67a8a4f6f724aaa86dea8dcb0a6975bfb1f10ff3beb64736f6c634300081a003360a060405234801561001057600080fd5b50604051611d83380380611d8383398101604081905261002f91610228565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03838116608052600480546001600160a01b0319169184169190911790556100bb600082610177565b506100e67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610177565b506101117f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610177565b5061012a600080516020611d6383398151915282610177565b50610143600080516020611d6383398151915283610177565b5061016e7f1ddc5e4a2f140581e5d9d5f758b2c0fbdd814c4017a02afc4b55cabcf2c1f10b83610177565b5050505061026b565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166102025760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a4506001610206565b5060005b92915050565b80516001600160a01b038116811461022357600080fd5b919050565b60008060006060848603121561023d57600080fd5b6102468461020c565b92506102546020850161020c565b91506102626040850161020c565b90509250925092565b608051611ac16102a260003960008181610215015281816105a6015281816105fb01528181610a5a0152610aaf0152611ac16000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c806385f438c1116100ee578063c709ab6e11610097578063d9caed1211610071578063d9caed1214610412578063e609055e14610425578063e63ab1e914610438578063eab103df1461045f57600080fd5b8063c709ab6e146103c9578063d547741f146103dc578063d936547e146103ef57600080fd5b80639a590427116100c85780639a5904271461039b5780639b19251a146103ae578063a217fddf146103c157600080fd5b806385f438c11461032857806391d148541461034f578063950837aa1461038857600080fd5b80632f2ff15d1161015b578063570618e111610135578063570618e1146102db5780635b112591146103025780635c975abb146103155780638456cb591461032057600080fd5b80632f2ff15d146102ad57806336568abe146102c05780633f4ba83a146102d357600080fd5b806321fc65f21161018c57806321fc65f21461024f578063248a9ca314610264578063252f07bf1461028857600080fd5b806301ffc9a7146101b35780630d391704146101db578063116191b614610210575b600080fd5b6101c66101c13660046114b7565b610472565b60405190151581526020015b60405180910390f35b6102027f1ddc5e4a2f140581e5d9d5f758b2c0fbdd814c4017a02afc4b55cabcf2c1f10b81565b6040519081526020016101d2565b6102377f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101d2565b61026261025d366004611557565b61050b565b005b6102026102723660046115ca565b6000908152600160208190526040909120015490565b6004546101c69074010000000000000000000000000000000000000000900460ff1681565b6102626102bb3660046115e3565b6106cb565b6102626102ce3660046115e3565b6106f7565b610262610748565b6102027f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b600454610237906001600160a01b031681565b60025460ff166101c6565b61026261077d565b6102027f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101c661035d3660046115e3565b60009182526001602090815260408084206001600160a01b0393909316845291905290205460ff1690565b610262610396366004611613565b6107af565b6102626103a9366004611613565b610854565b6102626103bc366004611613565b610908565b610202600081565b6102626103d7366004611630565b6109bf565b6102626103ea3660046115e3565b610b84565b6101c66103fd366004611613565b60036020526000908152604090205460ff1681565b6102626104203660046116d3565b610baa565b610262610433366004611714565b610ca2565b6102027f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b61026261046d3660046117b3565b610ece565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061050557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b610513610f24565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461053d81610f67565b610545610f71565b6001600160a01b03851660009081526003602052604090205460ff16610597576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105cb6001600160a01b0386167f000000000000000000000000000000000000000000000000000000000000000086610fb0565b6040517f5131ab590000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906106389088908a90899089908990600401611819565b600060405180830381600087803b15801561065257600080fd5b505af1158015610666573d6000803e3d6000fd5b50505050846001600160a01b0316866001600160a01b03167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d58686866040516106b19392919061185c565b60405180910390a3506106c46001600055565b5050505050565b600082815260016020819052604090912001546106e781610f67565b6106f18383611024565b50505050565b6001600160a01b0381163314610739576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61074382826110b7565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61077281610f67565b61077a61113e565b50565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107a781610f67565b61077a611190565b7f1ddc5e4a2f140581e5d9d5f758b2c0fbdd814c4017a02afc4b55cabcf2c1f10b6107d981610f67565b6001600160a01b038216610819576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50600480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a61087e81610f67565b6001600160a01b0382166108be576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a61093281610f67565b6001600160a01b038216610972576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b6109c7610f24565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46109f181610f67565b6109f9610f71565b6001600160a01b03861660009081526003602052604090205460ff16610a4b576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a7f6001600160a01b0387167f000000000000000000000000000000000000000000000000000000000000000087610fb0565b6040517fd0b492c30000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610aee9089908b908a908a908a908a90600401611930565b600060405180830381600087803b158015610b0857600080fd5b505af1158015610b1c573d6000803e3d6000fd5b50505050856001600160a01b0316876001600160a01b03167f2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b87878787604051610b699493929190611987565b60405180910390a350610b7c6001600055565b505050505050565b60008281526001602081905260409091200154610ba081610f67565b6106f183836110b7565b610bb2610f24565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610bdc81610f67565b610be4610f71565b6001600160a01b03831660009081526003602052604090205460ff16610c36576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c4a6001600160a01b0384168584610fb0565b826001600160a01b0316846001600160a01b03167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610c8f91815260200190565b60405180910390a3506107436001600055565b610caa610f24565b610cb2610f71565b60045474010000000000000000000000000000000000000000900460ff16610d06576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03841660009081526003602052604090205460ff16610d58576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610db8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ddc91906119b3565b9050610df36001600160a01b0386163330876111cd565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa158015610e7a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e9e91906119b3565b610ea891906119cc565b8787604051610ebb959493929190611a06565b60405180910390a250610b7c6001600055565b6000610ed981610f67565b506004805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b600260005403610f60576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61077a8133611206565b60025460ff1615610fae576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6040516001600160a01b0383811660248301526044820183905261074391859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061127d565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166110af5760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a4506001610505565b506000610505565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff16156110af5760008381526001602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610505565b6111466112f9565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b611198610f71565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586111733390565b6040516001600160a01b0384811660248301528381166044830152606482018390526106f19186918216906323b872dd90608401610fdd565b60008281526001602090815260408083206001600160a01b038516845290915290205460ff16611279576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044015b60405180910390fd5b5050565b60006112926001600160a01b03841683611335565b905080516000141580156112b75750808060200190518101906112b59190611a3f565b155b15610743576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611270565b60025460ff16610fae576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606113438383600061134a565b9392505050565b606081471015611388576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611270565b600080856001600160a01b031684866040516113a49190611a5c565b60006040518083038185875af1925050503d80600081146113e1576040519150601f19603f3d011682016040523d82523d6000602084013e6113e6565b606091505b50915091506113f6868383611400565b9695505050505050565b6060826114155761141082611475565b611343565b815115801561142c57506001600160a01b0384163b155b1561146e576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611270565b5080611343565b8051156114855780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156114c957600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461134357600080fd5b6001600160a01b038116811461077a57600080fd5b60008083601f84011261152057600080fd5b50813567ffffffffffffffff81111561153857600080fd5b60208301915083602082850101111561155057600080fd5b9250929050565b60008060008060006080868803121561156f57600080fd5b853561157a816114f9565b9450602086013561158a816114f9565b935060408601359250606086013567ffffffffffffffff8111156115ad57600080fd5b6115b98882890161150e565b969995985093965092949392505050565b6000602082840312156115dc57600080fd5b5035919050565b600080604083850312156115f657600080fd5b823591506020830135611608816114f9565b809150509250929050565b60006020828403121561162557600080fd5b8135611343816114f9565b60008060008060008060a0878903121561164957600080fd5b8635611654816114f9565b95506020870135611664816114f9565b945060408701359350606087013567ffffffffffffffff81111561168757600080fd5b61169389828a0161150e565b909450925050608087013567ffffffffffffffff8111156116b357600080fd5b87016060818a0312156116c557600080fd5b809150509295509295509295565b6000806000606084860312156116e857600080fd5b83356116f3816114f9565b92506020840135611703816114f9565b929592945050506040919091013590565b6000806000806000806080878903121561172d57600080fd5b863567ffffffffffffffff81111561174457600080fd5b61175089828a0161150e565b9097509550506020870135611764816114f9565b935060408701359250606087013567ffffffffffffffff81111561178757600080fd5b61179389828a0161150e565b979a9699509497509295939492505050565b801515811461077a57600080fd5b6000602082840312156117c557600080fd5b8135611343816117a5565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b03861681526001600160a01b03851660208201528360408201526080606082015260006118516080830184866117d0565b979650505050505050565b8381526040602082015260006118766040830184866117d0565b95945050505050565b6000813561188c816114f9565b6001600160a01b03168352602082013567ffffffffffffffff81168082146118b357600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126118ef57600080fd5b820160208101903567ffffffffffffffff81111561190c57600080fd5b80360382131561191b57600080fd5b606060408601526118766060860182846117d0565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a06060820152600061196860a0830185876117d0565b828103608084015261197a818561187f565b9998505050505050505050565b8481526060602082015260006119a16060830185876117d0565b8281036040840152611851818561187f565b6000602082840312156119c557600080fd5b5051919050565b81810381811115610505577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b606081526000611a1a6060830187896117d0565b8560208401528281036040840152611a338185876117d0565b98975050505050505050565b600060208284031215611a5157600080fd5b8151611343816117a5565b6000825160005b81811015611a7d5760208186018101518583015201611a63565b50600092019182525091905056fea26469706673582212201b2d9a074c8ea15918073e2cb2cb7dea6f0377907c07527efa37e3d88a13013464736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60c060405260001960035534801561001657600080fd5b5060405161178338038061178383398101604081905261003591610220565b60016000819055805460ff19169055838383836001600160a01b038416158061006557506001600160a01b038316155b8061007757506001600160a01b038216155b8061008957506001600160a01b038116155b156100a75760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03808516608052831660a0526100c5600082610154565b506100f07f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610154565b5061011b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610154565b506101467f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610154565b505050505050505050610274565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166101fa5760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101b23390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101fe565b5060005b92915050565b80516001600160a01b038116811461021b57600080fd5b919050565b6000806000806080858703121561023657600080fd5b61023f85610204565b935061024d60208601610204565b925061025b60408601610204565b915061026960608601610204565b905092959194509250565b60805160a0516114ab6102d86000396000818161021d01528181610531015281816108290152818161099a01528181610aef0152610c110152600081816101d1015281816104a1015281816105040152818161079901526107fc01526114ab6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80635e3e9fef116100d857806391d148541161008c578063d547741f11610066578063d547741f14610386578063d5abeb0114610399578063e63ab1e9146103a257600080fd5b806391d1485414610311578063a217fddf14610357578063a783c7891461035f57600080fd5b8063743e0c9b116100bd578063743e0c9b146102cf5780638456cb59146102e257806385f438c1146102ea57600080fd5b80635e3e9fef146102a95780636f8b44b0146102bc57600080fd5b8063248a9ca31161012f57806336568abe1161011457806336568abe146102835780633f4ba83a146102965780635c975abb1461029e57600080fd5b8063248a9ca31461023f5780632f2ff15d1461027057600080fd5b8063106e629011610160578063106e6290146101b9578063116191b6146101cc57806321e093b11461021857600080fd5b806301ffc9a71461017c578063057e0f25146101a4575b600080fd5b61018f61018a366004610fd3565b6103c9565b60405190151581526020015b60405180910390f35b6101b76101b236600461108e565b610462565b005b6101b76101c7366004611126565b6105fc565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b61026261024d366004611159565b60009081526002602052604090206001015490565b60405190815260200161019b565b6101b761027e366004611172565b6106a1565b6101b7610291366004611172565b6106cc565b6101b7610725565b60015460ff1661018f565b6101b76102b736600461119e565b61075a565b6101b76102ca366004611159565b6108ef565b6101b76102dd366004611159565b61095d565b6101b7610a07565b6102627f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b61018f61031f366004611172565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610262600081565b6102627f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101b7610394366004611172565b610a39565b61026260035481565b6102627f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061045c57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b61046a610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461049481610aa1565b61049c610aab565b6104c77f00000000000000000000000000000000000000000000000000000000000000008785610aea565b6040517fd0b492c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610563907f0000000000000000000000000000000000000000000000000000000000000000908b908b908b908b908a9060040161130e565b600060405180830381600087803b15801561057d57600080fd5b505af1158015610591573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff167f52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c878787866040516105e1949392919061137f565b60405180910390a2506105f46001600055565b505050505050565b610604610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461062e81610aa1565b610636610aab565b610641848484610aea565b8373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d58460405161068991815260200190565b60405180910390a25061069c6001600055565b505050565b6000828152600260205260409020600101546106bc81610aa1565b6106c68383610c72565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331461071b576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61069c8282610d72565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61074f81610aa1565b610757610e31565b50565b610762610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461078c81610aa1565b610794610aab565b6107bf7f00000000000000000000000000000000000000000000000000000000000000008684610aea565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab5990610859907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a906004016113b6565b600060405180830381600087803b15801561087357600080fd5b505af1158015610887573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d8686866040516108d593929190611408565b60405180910390a2506108e86001600055565b5050505050565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61091981610aa1565b610921610aab565b60038290556040518281527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c9060200160405180910390a15050565b610965610aab565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b1580156109f357600080fd5b505af11580156108e8573d6000803e3d6000fd5b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a3181610aa1565b610757610eae565b600082815260026020526040902060010154610a5481610aa1565b6106c68383610d72565b600260005403610a9a576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6107578133610f07565b60015460ff1615610ae8576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6003547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7c9190611422565b610b86908461143b565b1115610bbe576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201849052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b158015610c5557600080fd5b505af1158015610c69573d6000803e3d6000fd5b50505050505050565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610d6a57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610d083390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161045c565b50600061045c565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610d6a57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161045c565b610e39610f97565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b610eb6610aab565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610e84565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610f93576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440160405180910390fd5b5050565b60015460ff16610ae8576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215610fe557600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461101557600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461104057600080fd5b919050565b60008083601f84011261105757600080fd5b50813567ffffffffffffffff81111561106f57600080fd5b60208301915083602082850101111561108757600080fd5b9250929050565b60008060008060008060a087890312156110a757600080fd5b6110b08761101c565b955060208701359450604087013567ffffffffffffffff8111156110d357600080fd5b6110df89828a01611045565b90955093505060608701359150608087013567ffffffffffffffff81111561110657600080fd5b87016060818a03121561111857600080fd5b809150509295509295509295565b60008060006060848603121561113b57600080fd5b6111448461101c565b95602085013595506040909401359392505050565b60006020828403121561116b57600080fd5b5035919050565b6000806040838503121561118557600080fd5b823591506111956020840161101c565b90509250929050565b6000806000806000608086880312156111b657600080fd5b6111bf8661101c565b945060208601359350604086013567ffffffffffffffff8111156111e257600080fd5b6111ee88828901611045565b96999598509660600135949350505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff6112678261101c565b1682526000602082013567ffffffffffffffff811680821461128857600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126112c457600080fd5b820160208101903567ffffffffffffffff8111156112e157600080fd5b8036038213156112f057600080fd5b60606040860152611305606086018284611200565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a06060820152600061136060a083018587611200565b82810360808401526113728185611249565b9998505050505050505050565b848152606060208201526000611399606083018587611200565b82810360408401526113ab8185611249565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006113ab608083018486611200565b838152604060208201526000611305604083018486611200565b60006020828403121561143457600080fd5b5051919050565b8082018082111561045c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220cc1a57f74d7b05f5e8e999ef3dbaef86860a39bb638e05a5b55e143df03ffad964736f6c634300081a00336080604052348015600f57600080fd5b506001600055610dbc806100246000396000f3fe6080604052600436106100635760003560e01c8063c513169111610040578063c5131691146100c1578063e04d4f97146100e1578063f05b6abf146100f457005b8063357fc5a21461006c578063660b9de01461008c5780636ed70169146100ac57005b3661006a57005b005b34801561007857600080fd5b5061006a6100873660046106bd565b610114565b34801561009857600080fd5b5061006a6100a73660046106f9565b6101aa565b3480156100b857600080fd5b5061006a6101e6565b3480156100cd57600080fd5b5061006a6100dc3660046106bd565b61021b565b61006a6100ef366004610859565b6102f6565b34801561010057600080fd5b5061006a61010f366004610945565b61033a565b61011c61036f565b61013e73ffffffffffffffffffffffffffffffffffffffff83163383866103b2565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101a56001600055565b505050565b7f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e33826040516101db929190610a78565b60405180910390a150565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61022361036f565b6000610230600285610b57565b90508060000361026c576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61028e73ffffffffffffffffffffffffffffffffffffffff84163384846103b2565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101a56001600055565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa333485858560405161032d959493929190610c00565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b1463384848460405161032d9493929190610c8a565b6002600054036103ab576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261044790859061044d565b50505050565b600061046f73ffffffffffffffffffffffffffffffffffffffff8416836104e8565b905080516000141580156104945750808060200190518101906104929190610d4d565b155b156101a5576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b60606104f6838360006104fd565b9392505050565b60608147101561053b576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016104df565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105649190610d6a565b60006040518083038185875af1925050503d80600081146105a1576040519150601f19603f3d011682016040523d82523d6000602084013e6105a6565b606091505b50915091506105b68683836105c0565b9695505050505050565b6060826105d5576105d08261064f565b6104f6565b81511580156105f9575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610648576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016104df565b50806104f6565b80511561065f5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff811681146106b857600080fd5b919050565b6000806000606084860312156106d257600080fd5b833592506106e260208501610694565b91506106f060408501610694565b90509250925092565b60006020828403121561070b57600080fd5b813567ffffffffffffffff81111561072257600080fd5b8201606081850312156104f657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156107aa576107aa610734565b604052919050565b600082601f8301126107c357600080fd5b813567ffffffffffffffff8111156107dd576107dd610734565b61080e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610763565b81815284602083860101111561082357600080fd5b816020850160208301376000918101602001919091529392505050565b801515811461069157600080fd5b80356106b881610840565b60008060006060848603121561086e57600080fd5b833567ffffffffffffffff81111561088557600080fd5b610891868287016107b2565b9350506020840135915060408401356108a981610840565b809150509250925092565b600067ffffffffffffffff8211156108ce576108ce610734565b5060051b60200190565b600082601f8301126108e957600080fd5b81356108fc6108f7826108b4565b610763565b8082825260208201915060208360051b86010192508583111561091e57600080fd5b602085015b8381101561093b578035835260209283019201610923565b5095945050505050565b60008060006060848603121561095a57600080fd5b833567ffffffffffffffff81111561097157600080fd5b8401601f8101861361098257600080fd5b80356109906108f7826108b4565b8082825260208201915060208360051b8501019250888311156109b257600080fd5b602084015b838110156109f457803567ffffffffffffffff8111156109d657600080fd5b6109e58b6020838901016107b2565b845250602092830192016109b7565b509550505050602084013567ffffffffffffffff811115610a1457600080fd5b610a20868287016108d8565b9250506106f06040850161084e565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610ab683610694565b1660408201526000602083013567ffffffffffffffff8116808214610ada57600080fd5b6060840152506040830135368490037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112610b1657600080fd5b830160208101903567ffffffffffffffff811115610b3357600080fd5b803603821315610b4257600080fd5b606060808501526105b660a085018284610a2f565b600082610b8d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60005b83811015610bad578181015183820152602001610b95565b50506000910152565b60008151808452610bce816020860160208601610b92565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610c3560a0830186610bb6565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610c80578151865260209586019590910190600101610c62565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610d1d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610d08858351610bb6565b94506020938401939190910190600101610cce565b505050508281036040840152610d338186610c4e565b915050610d44606083018415159052565b95945050505050565b600060208284031215610d5f57600080fd5b81516104f681610840565b60008251610d7c818460208701610b92565b919091019291505056fea26469706673582212202c10ab94780c92f0dd656897a07025a325b3c0d38d265113cb76037b152226de64736f6c634300081a0033a2646970667358221220f8f43b660558c52473ee1300723aec1e586603b3efc026ce0c6cd6188357e12564736f6c634300081a0033",
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

// TSSUPDATERROLE is a free data retrieval call binding the contract method 0x0d391704.
//
// Solidity: function TSS_UPDATER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) TSSUPDATERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "TSS_UPDATER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSUPDATERROLE is a free data retrieval call binding the contract method 0x0d391704.
//
// Solidity: function TSS_UPDATER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestSession) TSSUPDATERROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.TSSUPDATERROLE(&_ERC20CustodyTest.CallOpts)
}

// TSSUPDATERROLE is a free data retrieval call binding the contract method 0x0d391704.
//
// Solidity: function TSS_UPDATER_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) TSSUPDATERROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.TSSUPDATERROLE(&_ERC20CustodyTest.CallOpts)
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

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestTSSUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testTSSUpgrade")
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestTSSUpgrade(&_ERC20CustodyTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestTSSUpgrade(&_ERC20CustodyTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testTSSUpgradeFailsIfSenderIsNotTSSUpdater")
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_ERC20CustodyTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_ERC20CustodyTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestTSSUpgradeFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testTSSUpgradeFailsIfZeroAddress")
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_ERC20CustodyTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_ERC20CustodyTest.TransactOpts)
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
