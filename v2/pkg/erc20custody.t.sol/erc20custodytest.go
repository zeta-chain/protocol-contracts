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

// ERC20CustodyTestMetaData contains all meta data concerning the ERC20CustodyTest contract.
var ERC20CustodyTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ASSET_HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WHITELISTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testDepositLegacy\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositLegacyFailsIfNotSupported\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositLegacyFailsIfTokenNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20PartialThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNoParamsThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testNewCustodyFailsIfAddressesAreZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgrade\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfSenderIsNotTSSUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUnwhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUnwhitelistFailsIfSenderIsNotWhitelister\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUnwhitelistFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhitelistFailsIfSenderIsNotWhitelister\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhitelistFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedCustodyTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LegacyMethodsNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5062012bb8806200003e6000396000f3fe608060405234801561001057600080fd5b50600436106103365760003560e01c806385226c81116101b2578063b5508aa9116100f9578063eb1ce7f9116100a2578063fa2a70741161007c578063fa2a7074146105b0578063fa7626d4146105b8578063fb176c12146105c5578063fe8e5f1b146105cd57600080fd5b8063eb1ce7f914610598578063f0c8e7e0146105a0578063f4221f08146105a857600080fd5b8063cbd57e2f116100d3578063cbd57e2f14610561578063e20c9f7114610569578063e63ab1e91461057157600080fd5b8063b5508aa914610539578063ba414fa614610541578063c713f8271461055957600080fd5b8063a217fddf1161015b578063a783c78911610135578063a783c78914610502578063b0464fdc14610529578063b421ca701461053157600080fd5b8063a217fddf146104ea578063a3f9d0e0146104f2578063a4943deb146104fa57600080fd5b8063916a17c61161018c578063916a17c6146104c55780639918c1c2146104da5780639fc7fd55146104e257600080fd5b806385226c811461048157806385f438c1146104965780639158c623146104bd57600080fd5b806349c783dd116102815780635d62c8601161022a5780637099d6f8116102045780637099d6f81461046157806371149c94146104695780637e91c50f1461047157806382c529921461047957600080fd5b80635d62c8601461041d57806366d9a9a0146104445780636a6218541461045957600080fd5b806351ecdf3c1161025b57806351ecdf3c146103d857806352ff5939146103e0578063570618e1146103e857600080fd5b806349c783dd146103c05780634b5838d2146103c85780634df42da1146103d057600080fd5b80632ade3880116102e35780633e73ecb4116102bd5780633e73ecb4146103a85780633ee92923146103b05780633f7286f4146103b857600080fd5b80632ade3880146103835780632be6a162146103985780633e5e3c23146103a057600080fd5b80631779672f116103145780631779672f146103555780631ed7831c1461035d578063284cb9291461037b57600080fd5b8063070f2ad01461033b5780630a9254e4146103455780630eee72a91461034d575b600080fd5b6103436105d5565b005b6103436107d6565b610343611232565b6103436114d7565b610365611628565b604051610372919061be74565b60405180910390f35b61034361168a565b61038b61196e565b604051610372919061bf10565b610343611ab0565b610365611c72565b610343611cd2565b610343612247565b610365612316565b610343612376565b6103436126b3565b6103436129d1565b610343612b29565b610343612cfa565b61040f7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b604051908152602001610372565b61040f7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b981565b61044c61352a565b604051610372919061c076565b6103436136ac565b610343613778565b610343613a25565b61034361423a565b6103436143c5565b610489614639565b604051610372919061c114565b61040f7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b610343614709565b6104cd6147d7565b604051610372919061c18b565b6103436148d2565b610343614bcc565b61040f600081565b610343614c9a565b6103436152a2565b61040f7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6104cd6154ca565b6103436155c5565b610489615afa565b610549615bca565b6040519015158152602001610372565b610343615c9e565b610343616941565b610365616a02565b61040f7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b610343616a62565b610343616b6e565b610343616d30565b610343616f9d565b601f546105499060ff1681565b61034361727b565b6103436178f9565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561064757600080fd5b505af115801561065b573d6000803e3d6000fd5b5050602854604080516001600160a01b039092166024830152600060448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250610727919060040161c222565b600060405180830381600087803b15801561074157600080fd5b505af1158015610755573d6000803e3d6000fd5b50506022546026546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063950837aa91506024015b600060405180830381600087803b1580156107bc57600080fd5b505af11580156107d0573d6000803e3d6000fd5b50505050565b602680547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602780548216611234179055602880549091166156781790556040516108289061bd83565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156108ad573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03929092169190911790556040516108f29061bd83565b604080825260049082018190527f7a6574610000000000000000000000000000000000000000000000000000000060608301526080602083018190528201527f5a4554410000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f080158015610976573d6000803e3d6000fd5b50602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260285460265492519085166024820152604481019390935292166064820152610a65919060840160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b00000000000000000000000000000000000000000000000000000000179052617aea565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff00000000000000000000000000000000000000009091168117909155602854602654604051929391821692911690610af19061bd91565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610b2d573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556020546025546028546026546040519385169492831693918316921690610b889061bd9f565b6001600160a01b039485168152928416602084015290831660408301529091166060820152608001604051809103906000f080158015610bcc573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055604051610c119061bdad565b604051809103906000f080158015610c2d573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556028546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b158015610cd957600080fd5b505af1158015610ced573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015610d6357600080fd5b505af1158015610d77573d6000803e3d6000fd5b50506020546022546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f9150602401600060405180830381600087803b158015610ddd57600080fd5b505af1158015610df1573d6000803e3d6000fd5b50506020546023546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef9150602401600060405180830381600087803b158015610e5757600080fd5b505af1158015610e6b573d6000803e3d6000fd5b5050602254602480546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639b19251a925001600060405180830381600087803b158015610ed057600080fd5b505af1158015610ee4573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610f4657600080fd5b505af1158015610f5a573d6000803e3d6000fd5b5050602480546026546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240938101939093521692506340c10f199150604401600060405180830381600087803b158015610fcb57600080fd5b505af1158015610fdf573d6000803e3d6000fd5b50506025546026546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42406024820152911692506340c10f199150604401600060405180830381600087803b15801561104e57600080fd5b505af1158015611062573d6000803e3d6000fd5b5050602480546022546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526207a1209381019390935216925063a9059cbb91506044016020604051808303816000875af11580156110d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110fc919061c235565b506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b15801561117d57600080fd5b505af1158015611191573d6000803e3d6000fd5b5050604080516080810182526026546001600160a01b039081168252602454811660208084019182526001848601908152855191820190955260008152606084018190528351602980549185167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161781559251602a8054919095169116179092559251602b55909350909150602c9061122d908261c31a565b505050565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc513169100000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b15801561131757600080fd5b505af115801561132b573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611416919060040161c222565b600060405180830381600087803b15801561143057600080fd5b505af1158015611444573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506114a19392831692909116908790879060040161c3d9565b600060405180830381600087803b1580156114bb57600080fd5b505af11580156114cf573d6000803e3d6000fd5b505050505050565b6024805460275460405160009381018490526001600160a01b03928316604482015291166064820152819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602854905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b1580156115b957600080fd5b505af11580156115cd573d6000803e3d6000fd5b5050604051630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401611416565b6060601680548060200260200160405190810160405280929190818152602001828054801561168057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611662575b5050505050905090565b6022546025546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600092919091169063d936547e90602401602060405180830381865afa1580156116f4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611718919061c235565b9050611725600082617b09565b6022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561179a57600080fd5b505af11580156117ae573d6000803e3d6000fd5b50506025546040516001600160a01b0390911692507faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549150600090a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561184357600080fd5b505af1158015611857573d6000803e3d6000fd5b50506022546025546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639b19251a9150602401600060405180830381600087803b1580156118bd57600080fd5b505af11580156118d1573d6000803e3d6000fd5b50506022546025546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063d936547e9150602401602060405180830381865afa15801561193a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061195e919061c235565b905061196b600182617b09565b50565b6060601e805480602002602001604051908101604052809291908181526020016000905b82821015611aa757600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015611a90578382906000526020600020018054611a039061c286565b80601f0160208091040260200160405190810160405280929190818152602001828054611a2f9061c286565b8015611a7c5780601f10611a5157610100808354040283529160200191611a7c565b820191906000526020600020905b815481529060010190602001808311611a5f57829003601f168201915b5050505050815260200190600101906119e4565b505050508152505081526020019060010190611992565b50505050905090565b60405163ca669fa760e01b81526101236004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611afe57600080fd5b505af1158015611b12573d6000803e3d6000fd5b50506040805161012360248201527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611bf3919060040161c222565b600060405180830381600087803b158015611c0d57600080fd5b505af1158015611c21573d6000803e3d6000fd5b50506022546025546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639b19251a91506024016107a2565b60606018805480602002602001604051908101604052809291908181526020018280548015611680576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611662575050505050905090565b602480546027546040516370a0823160e01b81526001600160a01b039182166004820152620186a09360009392909216916370a082319101602060405180830381865afa158015611d27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d4b919061c410565b9050611d58816000617b8b565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015611da8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dcc919061c410565b6027546040516001600160a01b0390911660248201526044810185905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391611eaf916001600160a01b039190911690600090869060040161c429565b600060405180830381600087803b158015611ec957600080fd5b505af1158015611edd573d6000803e3d6000fd5b50506022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015611f5657600080fd5b505af1158015611f6a573d6000803e3d6000fd5b50506024546027546040518881526001600160a01b039283169450911691507fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060200160405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561200f57600080fd5b505af1158015612023573d6000803e3d6000fd5b5050602254602754602480546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b03938416600482015290831691810191909152604481018990529116925063d9caed129150606401600060405180830381600087803b15801561209e57600080fd5b505af11580156120b2573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa158015612104573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612128919061c410565b90506121348186617b8b565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015612184573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121a8919061c410565b90506121bd816121b8888761c480565b617b8b565b602480546020546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa15801561220d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612231919061c410565b905061223e816000617b8b565b50505050505050565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024016112fd565b60606017805480602002602001604051908101604052809291908181526020018280548015611680576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611662575050505050905090565b6022546040517feab103df000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b039091169063eab103df90602401600060405180830381600087803b1580156123d557600080fd5b505af11580156123e9573d6000803e3d6000fd5b5050602480546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42409381019390935216925063095ea7b391506044016020604051808303816000875af115801561245f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612483919061c235565b506025546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240602482015291169063095ea7b3906044016020604051808303816000875af11580156124f4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612518919061c235565b50604080517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152815160058183030181526025820192839052630618f58760e51b9092527f73cba663000000000000000000000000000000000000000000000000000000006029820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906049015b600060405180830381600087803b1580156125c057600080fd5b505af11580156125d4573d6000803e3d6000fd5b505060225460275460405160609190911b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201526001600160a01b03909116925063e609055e915060340160408051601f19818403018152908290526024547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261267e926001600160a01b03909116906103e890879060040161c493565b600060405180830381600087803b15801561269857600080fd5b505af11580156126ac573d6000803e3d6000fd5b5050505050565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561271f57600080fd5b505af1158015612733573d6000803e3d6000fd5b5050602854602654604051600094508493506001600160a01b03928316929091169061275e9061bd91565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f08015801561279a573d6000803e3d6000fd5b50604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561280a57600080fd5b505af115801561281e573d6000803e3d6000fd5b50506020546026546040516001600160a01b039283169450600093509116906128469061bd91565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015612882573d6000803e3d6000fd5b50604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156128f257600080fd5b505af1158015612906573d6000803e3d6000fd5b50506020546028546040516001600160a01b0392831694509116915060009061292e9061bd91565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f08015801561296a573d6000803e3d6000fd5b506020546028546026546040519394506001600160a01b03928316939183169216906129959061bd91565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f08015801561122d573d6000803e3d6000fd5b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015612a4357600080fd5b505af1158015612a57573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015612ac757600080fd5b505af1158015612adb573d6000803e3d6000fd5b50506022546040517f950837aa000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b03909116925063950837aa91506024016107a2565b6024805460275460405160019381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602854905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612c0b57600080fd5b505af1158015612c1f573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015612c8f57600080fd5b505af1158015612ca3573d6000803e3d6000fd5b50506022546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321fc65f293506114a19260009216908790879060040161c3d9565b6022546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4600482015261432160248201819052916000916001600160a01b03909116906391d1485490604401602060405180830381865afa158015612d89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dad919061c235565b9050612db881617be3565b6022546040517f91d148540000000000000000000000000000000000000000000000000000000081527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60048201526001600160a01b03848116602483015260009216906391d1485490604401602060405180830381865afa158015612e42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e66919061c235565b9050612e7181617be3565b6022546028546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460048201526001600160a01b03918216602482015260009291909116906391d1485490604401602060405180830381865afa158015612f01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f25919061c235565b9050612f3081617c5d565b6022546028546040517f91d148540000000000000000000000000000000000000000000000000000000081527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60048201526001600160a01b03918216602482015260009291909116906391d1485490604401602060405180830381865afa158015612fc0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fe4919061c235565b9050612fef81617c5d565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561306157600080fd5b505af1158015613075573d6000803e3d6000fd5b50506022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156130ee57600080fd5b505af1158015613102573d6000803e3d6000fd5b5050602854604080516001600160a01b03928316815291891660208301527f4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6935001905060405180910390a16022546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0387811660048301529091169063950837aa90602401600060405180830381600087803b1580156131ae57600080fd5b505af11580156131c2573d6000803e3d6000fd5b5050505061324685602260009054906101000a90046001600160a01b03166001600160a01b0316635b1125916040518163ffffffff1660e01b8152600401602060405180830381865afa15801561321d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613241919061c4cd565b617caf565b6022546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460048201526001600160a01b038781166024830152909116906391d1485490604401602060405180830381865afa1580156132cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132f3919061c235565b93506132fe84617c5d565b6022546040517f91d148540000000000000000000000000000000000000000000000000000000081527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60048201526001600160a01b038781166024830152909116906391d1485490604401602060405180830381865afa158015613387573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133ab919061c235565b92506133b683617c5d565b6022546028546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460048201526001600160a01b0391821660248201529116906391d1485490604401602060405180830381865afa158015613441573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613465919061c235565b915061347082617be3565b6022546028546040517f91d148540000000000000000000000000000000000000000000000000000000081527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60048201526001600160a01b0391821660248201529116906391d1485490604401602060405180830381865afa1580156134fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061351f919061c235565b90506126ac81617be3565b6060601b805480602002602001604051908101604052809291908181526020016000905b82821015611aa757838290600052602060002090600202016040518060400160405290816000820180546135819061c286565b80601f01602080910402602001604051908101604052809291908181526020018280546135ad9061c286565b80156135fa5780601f106135cf576101008083540402835291602001916135fa565b820191906000526020600020905b8154815290600101906020018083116135dd57829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561369457602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116136415790505b5050505050815250508152602001906001019061354e565b6024805460275460405160009381018490526001600160a01b03928316604482015291166064820152819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc513169100000000000000000000000000000000000000000000000000000000179052602854905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa79060240161159f565b6022546040517feab103df000000000000000000000000000000000000000000000000000000008152600160048201526001600160a01b039091169063eab103df90602401600060405180830381600087803b1580156137d757600080fd5b505af11580156137eb573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b15801561385057600080fd5b505af1158015613864573d6000803e3d6000fd5b5050602480546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42409381019390935216925063095ea7b391506044016020604051808303816000875af11580156138da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138fe919061c235565b506025546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240602482015291169063095ea7b3906044016020604051808303816000875af115801561396f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613993919061c235565b50604080517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152815160058183030181526025820192839052630618f58760e51b9092527f584a7938000000000000000000000000000000000000000000000000000000006029820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906049016125a6565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a09060009060250160408051808303601f1901815290829052602480546021546370a0823160e01b85526001600160a01b0390811660048601529294506000939216916370a082319101602060405180830381865afa158015613ab8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613adc919061c410565b9050613ae9816000617b8b565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015613b39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b5d919061c410565b6020546040516001600160a01b0390911660248201526044810186905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391613c40916001600160a01b039190911690600090869060040161c429565b600060405180830381600087803b158015613c5a57600080fd5b505af1158015613c6e573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015613ce757600080fd5b505af1158015613cfb573d6000803e3d6000fd5b50506020546040517f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b9350613d3f92506001600160a01b039091169060299061c5d3565b60405180910390a16020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015613dbc57600080fd5b505af1158015613dd0573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507fde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e03590613e1e908990899060299061c5f5565b60405180910390a36022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015613e9b57600080fd5b505af1158015613eaf573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb972190613efd908990899060299061c5f5565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613f5e57600080fd5b505af1158015613f72573d6000803e3d6000fd5b50506022546021546024546040517f99a3c3560000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506399a3c3569450613fd29392831692909116908a908a9060299060040161c620565b600060405180830381600087803b158015613fec57600080fd5b505af1158015614000573d6000803e3d6000fd5b5050602480546021546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a0823191015b602060405180830381865afa158015614053573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614077919061c410565b90506140838187617b8b565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa1580156140d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906140f7919061c410565b9050614107816121b8898761c480565b602480546020546021546040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529082169381019390935260009291169063dd62ed3e90604401602060405180830381865afa15801561417d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141a1919061c410565b90506141ae816000617b8b565b602480546020546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa1580156141fe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614222919061c410565b905061422f816000617b8b565b505050505050505050565b6040517f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260019060009060250160408051808303601f190181529082905260285463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156142d357600080fd5b505af11580156142e7573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561435757600080fd5b505af115801561436b573d6000803e3d6000fd5b50506022546024546040517f99a3c3560000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506399a3c35693506114a19260009216908790879060299060040161c620565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561443757600080fd5b505af115801561444b573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b1580156144b057600080fd5b505af11580156144c4573d6000803e3d6000fd5b5050604051630618f58760e51b81527f584a7938000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561453457600080fd5b505af1158015614548573d6000803e3d6000fd5b5050602254602754602480546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b03938416600482015290831691810191909152600160448201529116925063d9caed129150606401600060405180830381600087803b1580156145c357600080fd5b505af11580156145d7573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156107bc57600080fd5b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015611aa757838290600052602060002001805461467c9061c286565b80601f01602080910402602001604051908101604052809291908181526020018280546146a89061c286565b80156146f55780601f106146ca576101008083540402835291602001916146f5565b820191906000526020600020905b8154815290600101906020018083116146d857829003601f168201915b50505050508152602001906001019061465d565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561477557600080fd5b505af1158015614789573d6000803e3d6000fd5b50506022546040517f9a590427000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b039091169250639a59042791506024016107a2565b6060601d805480602002602001604051908101604052809291908181526020016000905b82821015611aa75760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156148ba57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116148675790505b505050505081525050815260200190600101906147fb565b602480546027546040516001938101939093526001600160a01b03918216604484015216606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905260285490517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156149cd57600080fd5b505af11580156149e1573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b158015614a4657600080fd5b505af1158015614a5a573d6000803e3d6000fd5b5050604051630618f58760e51b81527f584a7938000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015614aca57600080fd5b505af1158015614ade573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f29450614b3c939283169290911690600190879060040161c3d9565b600060405180830381600087803b158015614b5657600080fd5b505af1158015614b6a573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561269857600080fd5b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015614c3857600080fd5b505af1158015614c4c573d6000803e3d6000fd5b50506022546040517f9b19251a000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b039091169250639b19251a91506024016107a2565b604080516004808252602480830184526020830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed7016900000000000000000000000000000000000000000000000000000000179052805460275494516370a0823160e01b81526001600160a01b0395861693810193909352620186a0946000939116916370a082319101602060405180830381865afa158015614d44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614d68919061c410565b9050614d75816000617b8b565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015614dc5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614de9919061c410565b6020546040516001600160a01b0390911660248201526044810186905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391614ecc916001600160a01b039190911690600090869060040161c429565b600060405180830381600087803b158015614ee657600080fd5b505af1158015614efa573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015614f7357600080fd5b505af1158015614f87573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0935001905060405180910390a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561504057600080fd5b505af1158015615054573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d59061509f908990899061c675565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561510057600080fd5b505af1158015615114573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506151719392831692909116908a908a9060040161c3d9565b600060405180830381600087803b15801561518b57600080fd5b505af115801561519f573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa1580156151f1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615215919061c410565b9050615222816000617b8b565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015615272573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615296919061c410565b90506141078185617b8b565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a09060009060250160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561533d57600080fd5b505af1158015615351573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061543c919060040161c222565b600060405180830381600087803b15801561545657600080fd5b505af115801561546a573d6000803e3d6000fd5b50506022546021546024546040517f99a3c3560000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506399a3c35694506114a19392831692909116908790879060299060040161c620565b6060601c805480602002602001604051908101604052809291908181526020016000905b82821015611aa75760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156155ad57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001906004019060208260030104928301926001038202915080841161555a5790505b505050505081525050815260200190600101906154ee565b6022546040517feab103df000000000000000000000000000000000000000000000000000000008152600160048201526103e8916001600160a01b03169063eab103df90602401600060405180830381600087803b15801561562657600080fd5b505af115801561563a573d6000803e3d6000fd5b5050602480546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42409381019390935216925063095ea7b391506044016020604051808303816000875af11580156156b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906156d4919061c235565b506025546022546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240602482015291169063095ea7b3906044016020604051808303816000875af1158015615745573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615769919061c235565b50602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa1580156157ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906157de919061c410565b6025546028546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa158015615830573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615854919061c410565b9050600060405160200161588b907f68656c6c6f000000000000000000000000000000000000000000000000000000815260050190565b60408051808303601f19018152908290526022546381bad6f360e01b8352600160048401819052602484018190526044840181905260648401526001600160a01b031660848301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561590e57600080fd5b505af1158015615922573d6000803e3d6000fd5b505060245460275460405160609190911b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201526001600160a01b0390911692507f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae915060340160408051601f19818403018152908290526159ab918890869061c68e565b60405180910390a26022546027546040805160609290921b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208301528051808303601401815260348301918290526024547fe609055e000000000000000000000000000000000000000000000000000000009092526001600160a01b039384169363e609055e93615a4a9391909116908990879060380161c493565b600060405180830381600087803b158015615a6457600080fd5b505af1158015615a78573d6000803e3d6000fd5b505050506107d08484615a8b919061c6b9565b602480546022546040516370a0823160e01b81526001600160a01b0391821660048201529116916370a082319101602060405180830381865afa158015615ad6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121b8919061c410565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015611aa7578382906000526020600020018054615b3d9061c286565b80601f0160208091040260200160405190810160405280929190818152602001828054615b699061c286565b8015615bb65780601f10615b8b57610100808354040283529160200191615bb6565b820191906000526020600020905b815481529060010190602001808311615b9957829003601f168201915b505050505081526020019060010190615b1e565b60085460009060ff1615615be2575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015615c73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615c97919061c410565b1415905090565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015615cf757600080fd5b505af1158015615d0b573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250615df6919060040161c222565b600060405180830381600087803b158015615e1057600080fd5b505af1158015615e24573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b158015615e7857600080fd5b505af1158015615e8c573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015615ee957600080fd5b505af1158015615efd573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250615fe8919060040161c222565b600060405180830381600087803b15801561600257600080fd5b505af1158015616016573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561606a57600080fd5b505af115801561607e573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156160db57600080fd5b505af11580156160ef573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561614357600080fd5b505af1158015616157573d6000803e3d6000fd5b505060248054602754604051620186a09381018490526001600160a01b039283166044820152911660648201529092506000915060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905251630618f58760e51b81527fd93c0665000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561625357600080fd5b505af1158015616267573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156162c457600080fd5b505af11580156162d8573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506163359392831692909116908790879060040161c3d9565b600060405180830381600087803b15801561634f57600080fd5b505af1158015616363573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156163c057600080fd5b505af11580156163d4573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561642857600080fd5b505af115801561643c573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a0823191015b602060405180830381865afa15801561648f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906164b3919061c410565b90506164c0816000617b8b565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015616510573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616534919061c410565b6020546040516001600160a01b0390911660248201526044810186905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391616617916001600160a01b039190911690600090869060040161c429565b600060405180830381600087803b15801561663157600080fd5b505af1158015616645573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156166be57600080fd5b505af11580156166d2573d6000803e3d6000fd5b505060208054602454602754604080516001600160a01b0394851681529485018c905291831684830152919091166060830152517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609350908190036080019150a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156167a857600080fd5b505af11580156167bc573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d590616807908990899061c675565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561686857600080fd5b505af115801561687c573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506168d99392831692909116908a908a9060040161c3d9565b600060405180830381600087803b1580156168f357600080fd5b505af1158015616907573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101614036565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a2000000000000000000000000000000000000000000000000000000001790526024805460275492516370a0823160e01b81526001600160a01b0393841660048201529394506000939216916370a082319101616472565b60606015805480602002602001604051908101604052809291908181526020018280548015611680576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611662575050505050905090565b600080604051602001616a98907f68656c6c6f000000000000000000000000000000000000000000000000000000815260050190565b60408051808303601f190181529082905260285463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015616aff57600080fd5b505af1158015616b13573d6000803e3d6000fd5b5050604051630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e0915060240161543c565b60405163ca669fa760e01b81526101236004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015616bbc57600080fd5b505af1158015616bd0573d6000803e3d6000fd5b50506040805161012360248201527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250616cb1919060040161c222565b600060405180830381600087803b158015616ccb57600080fd5b505af1158015616cdf573d6000803e3d6000fd5b50506022546025546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a59042791506024016107a2565b602480546027546040516001938101939093526001600160a01b03918216604484015216606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905260285490517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015616e2b57600080fd5b505af1158015616e3f573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b158015616ea457600080fd5b505af1158015616eb8573d6000803e3d6000fd5b5050604051630618f58760e51b81527f584a7938000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015616f2857600080fd5b505af1158015616f3c573d6000803e3d6000fd5b50506022546021546024546040517f99a3c3560000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506399a3c3569450614b3c939283169290911690600190879060299060040161c620565b602254602480546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600093919091169163d936547e9101602060405180830381865afa158015617006573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061702a919061c235565b9050617037600182617b09565b6022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156170ac57600080fd5b505af11580156170c0573d6000803e3d6000fd5b50506024546040516001600160a01b0390911692507f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919150600090a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561715557600080fd5b505af1158015617169573d6000803e3d6000fd5b5050602254602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a590427925001600060405180830381600087803b1580156171ce57600080fd5b505af11580156171e2573d6000803e3d6000fd5b5050602254602480546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529216935063d936547e925001602060405180830381865afa15801561724a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061726e919061c235565b905061196b600082617b09565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc5131691000000000000000000000000000000000000000000000000000000001790526024805460275492516370a0823160e01b81526001600160a01b0393841660048201529394506000939216916370a082319101602060405180830381865afa158015617354573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617378919061c410565b9050617385816000617b8b565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa1580156173d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906173f9919061c410565b6020546040516001600160a01b0390911660248201526044810186905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba3916174dc916001600160a01b039190911690600090869060040161c429565b600060405180830381600087803b1580156174f657600080fd5b505af115801561750a573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561758357600080fd5b505af1158015617597573d6000803e3d6000fd5b50506020547f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af6092506001600160a01b031690506175d560028861c6cc565b602454602754604080516001600160a01b03958616815260208101949094529184168383015292909216606082015290519081900360800190a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561768457600080fd5b505af1158015617698573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5906176e3908990899061c675565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561774457600080fd5b505af1158015617758573d6000803e3d6000fd5b50506022546021546024546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506177b59392831692909116908a908a9060040161c3d9565b600060405180830381600087803b1580156177cf57600080fd5b505af11580156177e3573d6000803e3d6000fd5b5050602480546027546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa158015617835573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617859919061c410565b905061786a816121b860028961c6cc565b602480546022546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa1580156178ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906178de919061c410565b9050614107816178ef60028a61c6cc565b6121b8908761c480565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561795757600080fd5b505af115801561796b573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250617a56919060040161c222565b600060405180830381600087803b158015617a7057600080fd5b505af1158015617a84573d6000803e3d6000fd5b5050602254602754602480546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b03938416600482015290831691810191909152604481018690529116925063d9caed12915060640161267e565b6000617af461bdbb565b617aff848483617d10565b9150505b92915050565b6040517ff7fe347700000000000000000000000000000000000000000000000000000000815282151560048201528115156024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f7fe3477906044015b60006040518083038186803b158015617b7757600080fd5b505afa1580156114cf573d6000803e3d6000fd5b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c5490604401617b5f565b6040517fa59828850000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a5982885906024015b60006040518083038186803b158015617c4957600080fd5b505afa1580156126ac573d6000803e3d6000fd5b6040517f0c9fd5810000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90630c9fd58190602401617c31565b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f690604401617b5f565b600080617d1d8584617d8b565b9050617d806040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f78790000008152508286604051602001617d6b92919061c707565b60405160208183030381529060405285617d97565b9150505b9392505050565b6000617d848383617dc5565b60c08101515160009015617dbb57617db484848460c00151617de0565b9050617d84565b617db48484617f86565b6000617dd18383618071565b617d8483836020015184617d97565b600080617deb618081565b90506000617df98683618154565b90506000617e1082606001518360200151856185fa565b90506000617e208383898961880c565b90506000617e2d82619689565b602081015181519192509060030b15617ea057898260400151604051602001617e5792919061c729565b60408051601f19818403018152908290527f08c379a0000000000000000000000000000000000000000000000000000000008252617e979160040161c222565b60405180910390fd5b6000617ee36040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001619858565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d90617f3690849060040161c222565b602060405180830381865afa158015617f53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617f77919061c4cd565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc92590617fdb90879060040161c222565b600060405180830381865afa158015617ff8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618020919081019061c863565b9050600061804e828560405160200161803a92919061c898565b604051602081830303815290604052619a58565b90506001600160a01b038116617aff578484604051602001617e5792919061c8c7565b61807d82826000619a6b565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c9061810890849060040161c972565b600060405180830381865afa158015618125573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261814d919081019061c9b9565b9250505090565b6181866040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d90506181d16040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6181da85619b6e565b602082015260006181ea86619f53565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa15801561822c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618254919081019061c9b9565b8683856020015160405160200161826e949392919061ca02565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb11906182c690859060040161c222565b600060405180830381865afa1580156182e3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261830b919081019061c9b9565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f69061835390849060040161cb06565b602060405180830381865afa158015618370573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618394919061c235565b6183a95781604051602001617e57919061cb58565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906183ee90849060040161cbea565b600060405180830381865afa15801561840b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618433919081019061c9b9565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f69061847a90849060040161cc3c565b602060405180830381865afa158015618497573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906184bb919061c235565b15618550576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061850590849060040161cc3c565b600060405180830381865afa158015618522573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261854a919081019061c9b9565b60408501525b846001600160a01b03166349c4fac8828660000151604051602001618575919061cc8e565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016185a192919061ccfa565b600060405180830381865afa1580156185be573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526185e6919081019061c9b9565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b60608152602001906001900390816186165790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250816000815181106186765761867661cd1f565b60200260200101819052506040518060400160405280600381526020017f2d726c0000000000000000000000000000000000000000000000000000000000815250816001815181106186ca576186ca61cd1f565b6020026020010181905250846040516020016186e6919061cd4e565b604051602081830303815290604052816002815181106187085761870861cd1f565b602002602001018190525082604051602001618724919061cdba565b604051602081830303815290604052816003815181106187465761874661cd1f565b6020026020010181905250600061875c82619689565b602080820151604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000081850190815282518084018452600080825290860152825180840190935290518252928101929092529192506187ed906040805180820182526000808252602091820152815180830190925284518252808501908201529061a1d6565b6188025785604051602001617e57919061cdfb565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d901561885c565b511590565b6189d057826020015115618918576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401617e97565b8260c00151156189d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401617e97565b6040805160ff8082526120008201909252600091816020015b60608152602001906001900390816189e957905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280618a449061ce8c565b935060ff1681518110618a5957618a5961cd1f565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e3700000000000000000000000000000000000000815250604051602001618aaa919061ceab565b604051602081830303815290604052828280618ac59061ce8c565b935060ff1681518110618ada57618ada61cd1f565b60200260200101819052506040518060400160405280600681526020017f6465706c6f790000000000000000000000000000000000000000000000000000815250828280618b279061ce8c565b935060ff1681518110618b3c57618b3c61cd1f565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d65000000000000000000000000000000000000815250828280618b899061ce8c565b935060ff1681518110618b9e57618b9e61cd1f565b60200260200101819052508760200151828280618bba9061ce8c565b935060ff1681518110618bcf57618bcf61cd1f565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e747261637450617468000000000000000000000000000000000000815250828280618c1c9061ce8c565b935060ff1681518110618c3157618c3161cd1f565b602090810291909101015287518282618c498161ce8c565b935060ff1681518110618c5e57618c5e61cd1f565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e49640000000000000000000000000000000000000000000000815250828280618cab9061ce8c565b935060ff1681518110618cc057618cc061cd1f565b6020026020010181905250618cd44661a237565b8282618cdf8161ce8c565b935060ff1681518110618cf457618cf461cd1f565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c650000000000000000000000000000000000815250828280618d419061ce8c565b935060ff1681518110618d5657618d5661cd1f565b602002602001018190525086828280618d6e9061ce8c565b935060ff1681518110618d8357618d8361cd1f565b6020908102919091010152855115618eaa5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f6465000000000000000000000060208201528282618dd48161ce8c565b935060ff1681518110618de957618de961cd1f565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d90618e3990899060040161c222565b600060405180830381865afa158015618e56573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618e7e919081019061c9b9565b8282618e898161ce8c565b935060ff1681518110618e9e57618e9e61cd1f565b60200260200101819052505b846020015115618f7a5760408051808201909152601281527f2d2d766572696679536f75726365436f6465000000000000000000000000000060208201528282618ef38161ce8c565b935060ff1681518110618f0857618f0861cd1f565b60200260200101819052506040518060400160405280600581526020017f66616c7365000000000000000000000000000000000000000000000000000000815250828280618f559061ce8c565b935060ff1681518110618f6a57618f6a61cd1f565b6020026020010181905250619141565b618fb26188578660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6190455760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282618ff58161ce8c565b935060ff168151811061900a5761900a61cd1f565b60200260200101819052508460a0015160405160200161902a919061cd4e565b604051602081830303815290604052828280618f559061ce8c565b8460c0015115801561908857506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261908690511590565b155b156191415760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826190cc8161ce8c565b935060ff16815181106190e1576190e161cd1f565b60200260200101819052506190f58861a2d7565b604051602001619105919061cd4e565b6040516020818303038152906040528282806191209061ce8c565b935060ff16815181106191355761913561cd1f565b60200260200101819052505b6040808601518151808301835260008082526020918201528251808401909352815183529081019082015261917590511590565b61920a5760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826191b88161ce8c565b935060ff16815181106191cd576191cd61cd1f565b602002602001018190525084604001518282806191e99061ce8c565b935060ff16815181106191fe576191fe61cd1f565b60200260200101819052505b60608501511561932b5760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826192538161ce8c565b935060ff16815181106192685761926861cd1f565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa1580156192d7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526192ff919081019061c9b9565b828261930a8161ce8c565b935060ff168151811061931f5761931f61cd1f565b60200260200101819052505b60e085015151156193d25760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826193758161ce8c565b935060ff168151811061938a5761938a61cd1f565b60200260200101819052506193a68560e001516000015161a237565b82826193b18161ce8c565b935060ff16815181106193c6576193c661cd1f565b60200260200101819052505b60e0850151602001511561947c5760408051808201909152600a81527f2d2d6761735072696365000000000000000000000000000000000000000000006020820152828261941f8161ce8c565b935060ff16815181106194345761943461cd1f565b60200260200101819052506194508560e001516020015161a237565b828261945b8161ce8c565b935060ff16815181106194705761947061cd1f565b60200260200101819052505b60e085015160400151156195265760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826194c98161ce8c565b935060ff16815181106194de576194de61cd1f565b60200260200101819052506194fa8560e001516040015161a237565b82826195058161ce8c565b935060ff168151811061951a5761951a61cd1f565b60200260200101819052505b60e085015160600151156195d05760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826195738161ce8c565b935060ff16815181106195885761958861cd1f565b60200260200101819052506195a48560e001516060015161a237565b82826195af8161ce8c565b935060ff16815181106195c4576195c461cd1f565b60200260200101819052505b60008160ff1667ffffffffffffffff8111156195ee576195ee61c257565b60405190808252806020026020018201604052801561962157816020015b606081526020019060019003908161960c5790505b50905060005b8260ff168160ff16101561967a57838160ff168151811061964a5761964a61cd1f565b6020026020010151828260ff16815181106196675761966761cd1f565b6020908102919091010152600101619627565b5093505050505b949350505050565b6196b06040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c916197369186910161cf16565b600060405180830381865afa158015619753573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261977b919081019061c9b9565b90506000619789868361adc6565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b81526004016197b9919061c114565b6000604051808303816000875af11580156197d8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619800919081019061cf5d565b805190915060030b158015906198195750602081015151155b80156198285750604081015151155b1561880257816000815181106198405761984061cd1f565b6020026020010151604051602001617e57919061d013565b6060600061988d8560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925286518252808701908201529091506198c49082905b9061af1b565b15619a215760006199418261993b846199356199078a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061af42565b9061afa4565b604080518082018252600181527f0a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506199a590829061af1b565b15619a0f57604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619a0c905b829061b029565b90505b619a188161b04f565b92505050617d84565b8215619a3a578484604051602001617e5792919061d1ff565b5050604080516020810190915260008152617d84565b509392505050565b6000808251602084016000f09392505050565b8160a0015115619a7a57505050565b6000619a8784848461b0b8565b90506000619a9482619689565b602081015181519192509060030b158015619b305750604080518082018252600781527f535543434553530000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619b30906040805180820182526000808252602091820152815180830190925284518252808501908201526198be565b15619b3d57505050505050565b60408201515115619b5d578160400151604051602001617e57919061d2a6565b80604051602001617e57919061d304565b60606000619ba38360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150619c08905b829061a1d6565b15619c7757604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617d8490619c7290839061b653565b61b04f565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619cd9905b829061b6dd565b600103619da657604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619d3f90619a05565b50604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617d8490619c72905b839061b029565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619e0590619c01565b15619f3c57604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290619e6d90839061b777565b905060008160018351619e80919061c480565b81518110619e9057619e9061cd1f565b60200260200101519050619f33619c72619f066040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925285518252808601908201529061b653565b95945050505050565b82604051602001617e57919061d36f565b50919050565b60606000619f888360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150619fea90619c01565b15619ff857617d848161b04f565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a05790619cd2565b60010361a0c157604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617d8490619c7290619d9f565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a12090619c01565b15619f3c57604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061a18890839061b777565b905060018151111561a1c457806002825161a1a3919061c480565b8151811061a1b35761a1b361cd1f565b602002602001015192505050919050565b5082604051602001617e57919061d36f565b80518251600091111561a1eb57506000617b03565b8151835160208501516000929161a2019161c6b9565b61a20b919061c480565b90508260200151810361a222576001915050617b03565b82516020840151819020912014905092915050565b6060600061a2448361b81c565b600101905060008167ffffffffffffffff81111561a2645761a26461c257565b6040519080825280601f01601f19166020018201604052801561a28e576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461a29857509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e534544000000000000000000000000000000000000000000008184019081528551808701875283815284019290925284518086019095525184529083015260609161a363905b829061b8fe565b1561a3a357505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a4029061a35c565b1561a44257505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a4a19061a35c565b1561a4e157505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a5409061a35c565b8061a5a55750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a5a59061a35c565b1561a5e557505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a6449061a35c565b8061a6a95750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a6a99061a35c565b1561a6e957505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a7489061a35c565b8061a7ad5750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a7ad9061a35c565b1561a7ed57505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a84c9061a35c565b8061a8b15750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a8b19061a35c565b1561a8f157505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a9509061a35c565b1561a99057505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a9ef9061a35c565b1561aa2f57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261aa8e9061a35c565b1561aace57505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ab2d9061a35c565b1561ab6d57505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261abcc9061a35c565b1561ac0c57505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ac6b9061a35c565b8061acd05750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261acd09061a35c565b1561ad1057505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ad6f9061a35c565b1561adaf57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151617e57929060200161d44d565b60608060005b845181101561ae51578185828151811061ade85761ade861cd1f565b602002602001015160405160200161ae0192919061c898565b60405160208183030381529060405291506001855161ae20919061c480565b811461ae49578160405160200161ae37919061d5b6565b60405160208183030381529060405291505b60010161adcc565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161ae6a579050509050838160008151811061ae955761ae9561cd1f565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061aee95761aee961cd1f565b6020026020010181905250818160028151811061af085761af0861cd1f565b6020908102919091010152949350505050565b602080830151835183519284015160009361af39929184919061b912565b14159392505050565b6040805180820190915260008082526020820152600061af74846000015185602001518560000151866020015161ba23565b905083602001518161af86919061c480565b8451859061af9590839061c480565b90525060208401525090919050565b604080518082019091526000808252602082015281518351101561afc9575081617b03565b602080830151908401516001911461aff05750815160208481015190840151829020919020145b801561b0215782518451859061b00790839061c480565b905250825160208501805161b01d90839061c6b9565b9052505b509192915050565b604080518082019091526000808252602082015261b04883838361bb43565b5092915050565b60606000826000015167ffffffffffffffff81111561b0705761b07061c257565b6040519080825280601f01601f19166020018201604052801561b09a576020820181803683370190505b509050600060208201905061b048818560200151866000015161bbee565b6060600061b0c4618081565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161b0e157905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061b13c9061ce8c565b935060ff168151811061b1515761b15161cd1f565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161b1a2919061d5f7565b60405160208183030381529060405282828061b1bd9061ce8c565b935060ff168151811061b1d25761b1d261cd1f565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061b21f9061ce8c565b935060ff168151811061b2345761b23461cd1f565b60200260200101819052508260405160200161b250919061cdba565b60405160208183030381529060405282828061b26b9061ce8c565b935060ff168151811061b2805761b28061cd1f565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061b2cd9061ce8c565b935060ff168151811061b2e25761b2e261cd1f565b602002602001018190525061b2f7878461bc68565b828261b3028161ce8c565b935060ff168151811061b3175761b31761cd1f565b60209081029190910101528551511561b3c35760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261b3698161ce8c565b935060ff168151811061b37e5761b37e61cd1f565b602002602001018190525061b39786600001518461bc68565b828261b3a28161ce8c565b935060ff168151811061b3b75761b3b761cd1f565b60200260200101819052505b85608001511561b4315760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261b40c8161ce8c565b935060ff168151811061b4215761b42161cd1f565b602002602001018190525061b497565b841561b4975760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261b4768161ce8c565b935060ff168151811061b48b5761b48b61cd1f565b60200260200101819052505b6040860151511561b5335760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261b4e18161ce8c565b935060ff168151811061b4f65761b4f661cd1f565b6020026020010181905250856040015182828061b5129061ce8c565b935060ff168151811061b5275761b52761cd1f565b60200260200101819052505b85606001511561b59d5760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261b57c8161ce8c565b935060ff168151811061b5915761b59161cd1f565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561b5bb5761b5bb61c257565b60405190808252806020026020018201604052801561b5ee57816020015b606081526020019060019003908161b5d95790505b50905060005b8260ff168160ff16101561b64757838160ff168151811061b6175761b61761cd1f565b6020026020010151828260ff168151811061b6345761b63461cd1f565b602090810291909101015260010161b5f4565b50979650505050505050565b604080518082019091526000808252602082015281518351101561b678575081617b03565b8151835160208501516000929161b68e9161c6b9565b61b698919061c480565b6020840151909150600190821461b6b9575082516020840151819020908220145b801561b6d45783518551869061b6d090839061c480565b9052505b50929392505050565b600080826000015161b701856000015186602001518660000151876020015161ba23565b61b70b919061c6b9565b90505b8351602085015161b71f919061c6b9565b811161b048578161b72f8161d63c565b925050826000015161b76685602001518361b74a919061c480565b865161b756919061c480565b838660000151876020015161ba23565b61b770919061c6b9565b905061b70e565b6060600061b785848461b6dd565b61b79090600161c6b9565b67ffffffffffffffff81111561b7a85761b7a861c257565b60405190808252806020026020018201604052801561b7db57816020015b606081526020019060019003908161b7c65790505b50905060005b8151811015619a505761b7f7619c72868661b029565b82828151811061b8095761b80961cd1f565b602090810291909101015260010161b7e1565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061b865577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061b891576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061b8af57662386f26fc10000830492506010015b6305f5e100831061b8c7576305f5e100830492506008015b612710831061b8db57612710830492506004015b6064831061b8ed576064830492506002015b600a8310617b035760010192915050565b600061b90a838361bca8565b159392505050565b60008085841161ba19576020841161b9c5576000841561b95d57600161b93986602061c480565b61b94490600861d656565b61b94f90600261d754565b61b959919061c480565b1990505b835181168561b96c898961c6b9565b61b976919061c480565b805190935082165b81811461b9b05787841161b9985787945050505050619681565b8361b9a28161d760565b94505082845116905061b97e565b61b9ba878561c6b9565b945050505050619681565b83832061b9d2858861c480565b61b9dc908761c6b9565b91505b85821061ba175784822080820361ba045761b9fa868461c6b9565b9350505050619681565b61ba0f60018461c480565b92505061b9df565b505b5092949350505050565b6000838186851161bb2e576020851161badd576000851561ba6f57600161ba4b87602061c480565b61ba5690600861d656565b61ba6190600261d754565b61ba6b919061c480565b1990505b8451811660008761ba808b8b61c6b9565b61ba8a919061c480565b855190915083165b82811461bacf5781861061bab75761baaa8b8b61c6b9565b9650505050505050619681565b8561bac18161d63c565b96505083865116905061ba92565b859650505050505050619681565b508383206000905b61baef868961c480565b821161bb2c5785832080820361bb0b5783945050505050619681565b61bb1660018561c6b9565b935050818061bb249061d63c565b92505061bae5565b505b61bb38878761c6b9565b979650505050505050565b6040805180820190915260008082526020820152600061bb75856000015186602001518660000151876020015161ba23565b60208087018051918601919091525190915061bb91908261c480565b83528451602086015161bba4919061c6b9565b810361bbb3576000855261bbe5565b8351835161bbc1919061c6b9565b8551869061bbd090839061c480565b905250835161bbdf908261c6b9565b60208601525b50909392505050565b6020811061bc26578151835261bc0560208461c6b9565b925061bc1260208361c6b9565b915061bc1f60208261c480565b905061bbee565b600019811561bc5557600161bc3c83602061c480565b61bc489061010061d754565b61bc52919061c480565b90505b9151835183169219169190911790915250565b6060600061bc768484618154565b805160208083015160405193945061bc909390910161d777565b60405160208183030381529060405291505092915050565b815181516000919081111561bcbb575081515b6020808501519084015160005b8381101561bd74578251825180821461bd4457600019602087101561bd235760018461bcf589602061c480565b61bcff919061c6b9565b61bd0a90600861d656565b61bd1590600261d754565b61bd1f919061c480565b1990505b818116838216818103911461bd41579750617b039650505050505050565b50505b61bd4f60208661c6b9565b945061bd5c60208561c6b9565b9350505060208161bd6d919061c6b9565b905061bcc8565b5084518651618802919061d7cf565b610c9f806200d7f083390190565b611e18806200e48f83390190565b6119e280620102a783390190565b610efa8062011c8983390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161bdfe61be03565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161bdfe6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561beb55783516001600160a01b031683526020938401939092019160010161be8e565b509095945050505050565b60005b8381101561bedb57818101518382015260200161bec3565b50506000910152565b6000815180845261befc81602086016020860161bec0565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561c00c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561bff2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261bfdc84865161bee4565b602095860195909450929092019160010161bfa2565b50919750505060209485019492909201915060010161bf38565b50929695505050505050565b600081518084526020840193506020830160005b8281101561c06c5781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161c02c565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561c00c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261c0e2604088018261bee4565b905060208201519150868103602088015261c0fd818361c018565b96505050602093840193919091019060010161c09e565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561c00c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261c17685835161bee4565b9450602093840193919091019060010161c13c565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561c00c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261c20c604087018261c018565b955050602093840193919091019060010161c1b3565b602081526000617d84602083018461bee4565b60006020828403121561c24757600080fd5b81518015158114617d8457600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c9082168061c29a57607f821691505b602082108103619f4d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f82111561122d57806000526020600020601f840160051c8101602085101561c2fa5750805b601f840160051c820191505b818110156126ac576000815560010161c306565b815167ffffffffffffffff81111561c3345761c33461c257565b61c3488161c342845461c286565b8461c2d3565b6020601f82116001811461c37c576000831561c3645750848201515b600019600385901b1c1916600184901b1784556126ac565b600084815260208120601f198516915b8281101561c3ac578785015182556020948501946001909201910161c38c565b508482101561c3ca5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6001600160a01b03851681526001600160a01b0384166020820152826040820152608060608201526000618802608083018461bee4565b60006020828403121561c42257600080fd5b5051919050565b6001600160a01b0384168152826020820152606060408201526000619f33606083018461bee4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810381811115617b0357617b0361c451565b60808152600061c4a6608083018761bee4565b6001600160a01b0386166020840152846040840152828103606084015261bb38818561bee4565b60006020828403121561c4df57600080fd5b81516001600160a01b0381168114617d8457600080fd5b6001600160a01b0381541682526001600160a01b03600182015416602083015260028101546040830152600060038201608060608501526000815461c53a8161c286565b806080880152600182166000811461c559576001811461c5935761c5c7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660a089015260a082151560051b890101935061c5c7565b84600052602060002060005b8381101561c5be5781548a820160a0015260019091019060200161c59f565b890160a0019450505b50919695505050505050565b6001600160a01b0383168152604060208201526000619681604083018461c4f6565b83815260606020820152600061c60e606083018561bee4565b8281036040840152618802818561c4f6565b6001600160a01b03861681526001600160a01b038516602082015283604082015260a06060820152600061c65760a083018561bee4565b828103608084015261c669818561c4f6565b98975050505050505050565b828152604060208201526000619681604083018461bee4565b60608152600061c6a1606083018661bee4565b8460208401528281036040840152618802818561bee4565b80820180821115617b0357617b0361c451565b60008261c702577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6001600160a01b0383168152604060208201526000619681604083018461bee4565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161c76181601a85016020880161bec0565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161c79e81601c84016020880161bec0565b01601c01949350505050565b6040516060810167ffffffffffffffff8111828210171561c7cd5761c7cd61c257565b60405290565b60008067ffffffffffffffff84111561c7ee5761c7ee61c257565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561c81d5761c81d61c257565b60405283815290508082840185101561c83557600080fd5b619a5084602083018561bec0565b600082601f83011261c85457600080fd5b617d848383516020850161c7d3565b60006020828403121561c87557600080fd5b815167ffffffffffffffff81111561c88c57600080fd5b617aff8482850161c843565b6000835161c8aa81846020880161bec0565b83519083019061c8be81836020880161bec0565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161c8ff81601a85016020880161bec0565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161c93c81603384016020880161bec0565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000617d84608083018461bee4565b60006020828403121561c9cb57600080fd5b815167ffffffffffffffff81111561c9e257600080fd5b8201601f8101841361c9f357600080fd5b617aff8482516020840161c7d3565b6000855161ca14818460208a0161bec0565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161ca4e816001840160208a0161bec0565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161ca8c81600284016020890161bec0565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161cace81600284016020880161bec0565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061cb19604083018461bee4565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161cb9081601f85016020870161bec0565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061cbfd604083018461bee4565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061cc4f604083018461bee4565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161ccc681601485016020870161bec0565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061cd0d604083018561bee4565b8281036020840152617d80818561bee4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161cd8681600185016020870161bec0565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161cdcc81846020870161bec0565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161ce7f81604b85016020870161bec0565b91909101604b0192915050565b600060ff821660ff810361cea25761cea261c451565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161cf0981602985016020870161bec0565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000617d84608083018461bee4565b60006020828403121561cf6f57600080fd5b815167ffffffffffffffff81111561cf8657600080fd5b82016060818503121561cf9857600080fd5b61cfa061c7aa565b81518060030b811461cfb157600080fd5b8152602082015167ffffffffffffffff81111561cfcd57600080fd5b61cfd98682850161c843565b602083015250604082015167ffffffffffffffff81111561cff957600080fd5b61d0058682850161c843565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161d07181602185016020870161bec0565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161d25d81602185016020880161bec0565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161d29a81602e84016020880161bec0565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161cf0981602985016020870161bec0565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161d36281602285016020870161bec0565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161d3a781600e85016020870161bec0565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161d48581601885016020880161bec0565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161d4c281601c84016020880161bec0565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161d5c881846020870161bec0565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161d62f81601c85016020870161bec0565b91909101601c0192915050565b6000600019820361d64f5761d64f61c451565b5060010190565b8082028115828204841417617b0357617b0361c451565b6001815b600184111561d6a85780850481111561d68c5761d68c61c451565b600184161561d69a57908102905b60019390931c92800261d671565b935093915050565b60008261d6bf57506001617b03565b8161d6cc57506000617b03565b816001811461d6e2576002811461d6ec5761d708565b6001915050617b03565b60ff84111561d6fd5761d6fd61c451565b50506001821b617b03565b5060208310610133831016604e8410600b841016171561d72b575081810a617b03565b61d738600019848461d66d565b806000190482111561d74c5761d74c61c451565b029392505050565b6000617d84838361d6b0565b60008161d76f5761d76f61c451565b506000190190565b6000835161d78981846020880161bec0565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161d7c381600184016020880161bec0565b01600101949350505050565b818103600083128015838313168383128216171561b0485761b04861c45156fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220a043c41353215fce25ecb67a8a4f6f724aaa86dea8dcb0a6975bfb1f10ff3beb64736f6c634300081a003360a060405234801561001057600080fd5b50604051611e18380380611e1883398101604081905261002f916101fd565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03838116608052600480546001600160a01b0319169184169190911790556100bb60008261014c565b506100e67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261014c565b506101117f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361014c565b5061012a600080516020611df88339815191528261014c565b50610143600080516020611df88339815191528361014c565b50505050610240565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166101d75760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016101db565b5060005b92915050565b80516001600160a01b03811681146101f857600080fd5b919050565b60008060006060848603121561021257600080fd5b61021b846101e1565b9250610229602085016101e1565b9150610237604085016101e1565b90509250925092565b608051611b81610277600039600081816101d501528181610574015281816105c9015281816109ab0152610a000152611b816000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c806385f438c1116100ee578063a217fddf11610097578063d9caed1211610071578063d9caed12146103e0578063e609055e146103f3578063e63ab1e914610406578063eab103df1461042d57600080fd5b8063a217fddf146103a2578063d547741f146103aa578063d936547e146103bd57600080fd5b806399a3c356116100c857806399a3c356146103695780639a5904271461037c5780639b19251a1461038f57600080fd5b806385f438c1146102f657806391d148541461031d578063950837aa1461035657600080fd5b806336568abe116101505780635b1125911161012a5780635b112591146102d05780635c975abb146102e35780638456cb59146102ee57600080fd5b806336568abe1461028e5780633f4ba83a146102a1578063570618e1146102a957600080fd5b8063248a9ca311610181578063248a9ca314610224578063252f07bf146102565780632f2ff15d1461027b57600080fd5b806301ffc9a7146101a8578063116191b6146101d057806321fc65f21461020f575b600080fd5b6101bb6101b6366004611573565b610440565b60405190151581526020015b60405180910390f35b6101f77f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101c7565b61022261021d366004611613565b6104d9565b005b610248610232366004611686565b6000908152600160208190526040909120015490565b6040519081526020016101c7565b6004546101bb9074010000000000000000000000000000000000000000900460ff1681565b61022261028936600461169f565b610699565b61022261029c36600461169f565b6106c5565b610222610716565b6102487f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b6004546101f7906001600160a01b031681565b60025460ff166101bb565b61022261074b565b6102487f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101bb61032b36600461169f565b60009182526001602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6102226103643660046116cf565b61077d565b6102226103773660046116ec565b610910565b61022261038a3660046116cf565b610ad5565b61022261039d3660046116cf565b610b89565b610248600081565b6102226103b836600461169f565b610c40565b6101bb6103cb3660046116cf565b60036020526000908152604090205460ff1681565b6102226103ee36600461178f565b610c66565b6102226104013660046117d0565b610d5e565b6102487f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b61022261043b36600461186f565b610f8a565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104d357507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104e1610fe0565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461050b81611023565b61051361102d565b6001600160a01b03851660009081526003602052604090205460ff16610565576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105996001600160a01b0386167f00000000000000000000000000000000000000000000000000000000000000008661106c565b6040517f5131ab590000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906106069088908a908990899089906004016118d5565b600060405180830381600087803b15801561062057600080fd5b505af1158015610634573d6000803e3d6000fd5b50505050846001600160a01b0316866001600160a01b03167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d586868660405161067f93929190611918565b60405180910390a3506106926001600055565b5050505050565b600082815260016020819052604090912001546106b581611023565b6106bf83836110e0565b50505050565b6001600160a01b0381163314610707576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107118282611173565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61074081611023565b6107486111fa565b50565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61077581611023565b61074861124c565b600061078881611023565b6001600160a01b0382166107c8576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004546107ff907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4906001600160a01b0316611173565b50600454610837907f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a906001600160a01b0316611173565b506108627f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4836110e0565b5061088d7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a836110e0565b50600454604080516001600160a01b03928316815291841660208301527f4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6910160405180910390a150600480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b610918610fe0565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461094281611023565b61094a61102d565b6001600160a01b03861660009081526003602052604090205460ff1661099c576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109d06001600160a01b0387167f00000000000000000000000000000000000000000000000000000000000000008761106c565b6040517faa0c0fc10000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063aa0c0fc190610a3f9089908b908a908a908a908a906004016119f0565b600060405180830381600087803b158015610a5957600080fd5b505af1158015610a6d573d6000803e3d6000fd5b50505050856001600160a01b0316876001600160a01b03167f7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb972187878787604051610aba9493929190611a47565b60405180910390a350610acd6001600055565b505050505050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a610aff81611023565b6001600160a01b038216610b3f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a610bb381611023565b6001600160a01b038216610bf3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b60008281526001602081905260409091200154610c5c81611023565b6106bf8383611173565b610c6e610fe0565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610c9881611023565b610ca061102d565b6001600160a01b03831660009081526003602052604090205460ff16610cf2576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d066001600160a01b038416858461106c565b826001600160a01b0316846001600160a01b03167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610d4b91815260200190565b60405180910390a3506107116001600055565b610d66610fe0565b610d6e61102d565b60045474010000000000000000000000000000000000000000900460ff16610dc2576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03841660009081526003602052604090205460ff16610e14576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610e74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e989190611a73565b9050610eaf6001600160a01b038616333087611289565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa158015610f36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5a9190611a73565b610f649190611a8c565b8787604051610f77959493929190611ac6565b60405180910390a250610acd6001600055565b6000610f9581611023565b506004805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b60026000540361101c576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61074881336112c2565b60025460ff161561106a576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6040516001600160a01b0383811660248301526044820183905261071191859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611339565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff1661116b5760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016104d3565b5060006104d3565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff161561116b5760008381526001602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104d3565b6112026113b5565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b61125461102d565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861122f3390565b6040516001600160a01b0384811660248301528381166044830152606482018390526106bf9186918216906323b872dd90608401611099565b60008281526001602090815260408083206001600160a01b038516845290915290205460ff16611335576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044015b60405180910390fd5b5050565b600061134e6001600160a01b038416836113f1565b905080516000141580156113735750808060200190518101906113719190611aff565b155b15610711576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b038416600482015260240161132c565b60025460ff1661106a576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606113ff83836000611406565b9392505050565b606081471015611444576040517fcd78605900000000000000000000000000000000000000000000000000000000815230600482015260240161132c565b600080856001600160a01b031684866040516114609190611b1c565b60006040518083038185875af1925050503d806000811461149d576040519150601f19603f3d011682016040523d82523d6000602084013e6114a2565b606091505b50915091506114b28683836114bc565b9695505050505050565b6060826114d1576114cc82611531565b6113ff565b81511580156114e857506001600160a01b0384163b155b1561152a576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161132c565b50806113ff565b8051156115415780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561158557600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146113ff57600080fd5b6001600160a01b038116811461074857600080fd5b60008083601f8401126115dc57600080fd5b50813567ffffffffffffffff8111156115f457600080fd5b60208301915083602082850101111561160c57600080fd5b9250929050565b60008060008060006080868803121561162b57600080fd5b8535611636816115b5565b94506020860135611646816115b5565b935060408601359250606086013567ffffffffffffffff81111561166957600080fd5b611675888289016115ca565b969995985093965092949392505050565b60006020828403121561169857600080fd5b5035919050565b600080604083850312156116b257600080fd5b8235915060208301356116c4816115b5565b809150509250929050565b6000602082840312156116e157600080fd5b81356113ff816115b5565b60008060008060008060a0878903121561170557600080fd5b8635611710816115b5565b95506020870135611720816115b5565b945060408701359350606087013567ffffffffffffffff81111561174357600080fd5b61174f89828a016115ca565b909450925050608087013567ffffffffffffffff81111561176f57600080fd5b87016080818a03121561178157600080fd5b809150509295509295509295565b6000806000606084860312156117a457600080fd5b83356117af816115b5565b925060208401356117bf816115b5565b929592945050506040919091013590565b600080600080600080608087890312156117e957600080fd5b863567ffffffffffffffff81111561180057600080fd5b61180c89828a016115ca565b9097509550506020870135611820816115b5565b935060408701359250606087013567ffffffffffffffff81111561184357600080fd5b61184f89828a016115ca565b979a9699509497509295939492505050565b801515811461074857600080fd5b60006020828403121561188157600080fd5b81356113ff81611861565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b03861681526001600160a01b038516602082015283604082015260806060820152600061190d60808301848661188c565b979650505050505050565b83815260406020820152600061193260408301848661188c565b95945050505050565b60008135611948816115b5565b6001600160a01b031683526020820135611961816115b5565b6001600160a01b03166020840152604082810135908401526060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126119af57600080fd5b820160208101903567ffffffffffffffff8111156119cc57600080fd5b8036038213156119db57600080fd5b6080606086015261193260808601828461188c565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a060608201526000611a2860a08301858761188c565b8281036080840152611a3a818561193b565b9998505050505050505050565b848152606060208201526000611a6160608301858761188c565b828103604084015261190d818561193b565b600060208284031215611a8557600080fd5b5051919050565b818103818111156104d3577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b606081526000611ada60608301878961188c565b8560208401528281036040840152611af381858761188c565b98975050505050505050565b600060208284031215611b1157600080fd5b81516113ff81611861565b6000825160005b81811015611b3d5760208186018101518583015201611b23565b50600092019182525091905056fea264697066735822122051b8efab451617e3639e230507d505329c61fd7005e44790e327352674e231e264736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60c060405260001960045534801561001657600080fd5b506040516119e23803806119e283398101604081905261003591610238565b60016000819055805460ff19169055838383836001600160a01b038416158061006557506001600160a01b038316155b8061007757506001600160a01b038216155b8061008957506001600160a01b038116155b156100a75760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0384811660805283811660a052600380546001600160a01b0319169184169190911790556100dd60008261016c565b506101087f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361016c565b506101337f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb8361016c565b5061015e7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261016c565b50505050505050505061028c565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166102125760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101ca3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610216565b5060005b92915050565b80516001600160a01b038116811461023357600080fd5b919050565b6000806000806080858703121561024e57600080fd5b6102578561021c565b93506102656020860161021c565b92506102736040860161021c565b91506102816060860161021c565b905092959194509250565b60805160a0516116f26102f060003960008181610220015281816106d80152818161086d015281816109e301528181610d0c0152610e2e0152600081816101d401528181610648015281816106ab015281816107dd015261084001526116f26000f3fe608060405234801561001057600080fd5b506004361061018d5760003560e01c80636f8728ad116100e3578063950837aa1161008c578063d547741f11610066578063d547741f146103cf578063d5abeb01146103e2578063e63ab1e9146103eb57600080fd5b8063950837aa1461038d578063a217fddf146103a0578063a783c789146103a857600080fd5b80638456cb59116100bd5780638456cb591461031857806385f438c11461032057806391d148541461034757600080fd5b80636f8728ad146102df5780636f8b44b0146102f2578063743e0c9b1461030557600080fd5b80632f2ff15d116101455780635b1125911161011f5780635b112591146102a15780635c975abb146102c15780635e3e9fef146102cc57600080fd5b80632f2ff15d1461027357806336568abe146102865780633f4ba83a1461029957600080fd5b8063116191b611610176578063116191b6146101cf57806321e093b11461021b578063248a9ca31461024257600080fd5b806301ffc9a714610192578063106e6290146101ba575b600080fd5b6101a56101a03660046111f0565b610412565b60405190151581526020015b60405180910390f35b6101cd6101c8366004611262565b6104ab565b005b6101f67f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b1565b6101f67f000000000000000000000000000000000000000000000000000000000000000081565b610265610250366004611295565b60009081526002602052604090206001015490565b6040519081526020016101b1565b6101cd6102813660046112ae565b610550565b6101cd6102943660046112ae565b61057b565b6101cd6105d4565b6003546101f69073ffffffffffffffffffffffffffffffffffffffff1681565b60015460ff166101a5565b6101cd6102da366004611323565b610609565b6101cd6102ed366004611385565b61079e565b6101cd610300366004611295565b610938565b6101cd610313366004611295565b6109a6565b6101cd610a50565b6102657f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101a56103553660046112ae565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6101cd61039b36600461141d565b610a82565b610265600081565b6102657f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101cd6103dd3660046112ae565b610c56565b61026560045481565b6102657f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104a557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104b3610c7b565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46104dd81610cbe565b6104e5610cc8565b6104f0848484610d07565b8373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d58460405161053891815260200190565b60405180910390a25061054b6001600055565b505050565b60008281526002602052604090206001015461056b81610cbe565b6105758383610e8f565b50505050565b73ffffffffffffffffffffffffffffffffffffffff811633146105ca576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61054b8282610f8f565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6105fe81610cbe565b61060661104e565b50565b610611610c7b565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461063b81610cbe565b610643610cc8565b61066e7f00000000000000000000000000000000000000000000000000000000000000008684610d07565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab5990610708907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a90600401611481565b600060405180830381600087803b15801561072257600080fd5b505af1158015610736573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d868686604051610784939291906114de565b60405180910390a2506107976001600055565b5050505050565b6107a6610c7b565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46107d081610cbe565b6107d8610cc8565b6108037f00000000000000000000000000000000000000000000000000000000000000008785610d07565b6040517faa0c0fc100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063aa0c0fc19061089f907f0000000000000000000000000000000000000000000000000000000000000000908b908b908b908b908a906004016115cc565b600060405180830381600087803b1580156108b957600080fd5b505af11580156108cd573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff167f5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff08787878660405161091d949392919061163d565b60405180910390a2506109306001600055565b505050505050565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61096281610cbe565b61096a610cc8565b60048290556040518281527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c9060200160405180910390a15050565b6109ae610cc8565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b158015610a3c57600080fd5b505af1158015610797573d6000803e3d6000fd5b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a7a81610cbe565b6106066110cb565b6000610a8d81610cbe565b73ffffffffffffffffffffffffffffffffffffffff8216610ada576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354610b1e907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e49073ffffffffffffffffffffffffffffffffffffffff16610f8f565b50600354610b63907f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb9073ffffffffffffffffffffffffffffffffffffffff16610f8f565b50610b8e7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610e8f565b50610bb97f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610e8f565b506003546040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301527f33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e910160405180910390a150600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600082815260026020526040902060010154610c7181610cbe565b6105758383610f8f565b600260005403610cb7576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6106068133611124565b60015460ff1615610d05576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6004547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d999190611669565b610da39084611682565b1115610ddb576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201849052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b158015610e7257600080fd5b505af1158015610e86573d6000803e3d6000fd5b50505050505050565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610f8757600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610f253390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016104a5565b5060006104a5565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610f8757600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104a5565b6110566111b4565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6110d3610cc8565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336110a1565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166111b0576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440160405180910390fd5b5050565b60015460ff16610d05576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561120257600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461123257600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461125d57600080fd5b919050565b60008060006060848603121561127757600080fd5b61128084611239565b95602085013595506040909401359392505050565b6000602082840312156112a757600080fd5b5035919050565b600080604083850312156112c157600080fd5b823591506112d160208401611239565b90509250929050565b60008083601f8401126112ec57600080fd5b50813567ffffffffffffffff81111561130457600080fd5b60208301915083602082850101111561131c57600080fd5b9250929050565b60008060008060006080868803121561133b57600080fd5b61134486611239565b945060208601359350604086013567ffffffffffffffff81111561136757600080fd5b611373888289016112da565b96999598509660600135949350505050565b60008060008060008060a0878903121561139e57600080fd5b6113a787611239565b955060208701359450604087013567ffffffffffffffff8111156113ca57600080fd5b6113d689828a016112da565b90955093505060608701359150608087013567ffffffffffffffff8111156113fd57600080fd5b87016080818a03121561140f57600080fd5b809150509295509295509295565b60006020828403121561142f57600080fd5b61123282611239565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006114d3608083018486611438565b979650505050505050565b8381526040602082015260006114f8604083018486611438565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff61151f82611239565b16825273ffffffffffffffffffffffffffffffffffffffff61154360208301611239565b1660208301526040818101359083015260006060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261158b57600080fd5b820160208101903567ffffffffffffffff8111156115a857600080fd5b8036038213156115b757600080fd5b608060608601526114f8608086018284611438565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a06060820152600061161e60a083018587611438565b82810360808401526116308185611501565b9998505050505050505050565b848152606060208201526000611657606083018587611438565b82810360408401526114d38185611501565b60006020828403121561167b57600080fd5b5051919050565b808201808211156104a5577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220fc7262a0fde87d560cbb5c9f1834fd4872192776fa2aa5b2f0c6b493cb1f637864736f6c634300081a00336080604052348015600f57600080fd5b506001600055610ed6806100246000396000f3fe60806040526004361061006e5760003560e01c8063c51316911161004b578063c5131691146100d5578063c9028a36146100f5578063e04d4f9714610115578063f05b6abf1461012857005b8063357fc5a214610077578063676cc054146100975780636ed70169146100c057005b3661007557005b005b34801561008357600080fd5b50610075610092366004610724565b610148565b6100aa6100a5366004610760565b6101de565b6040516100b7919061085b565b60405180910390f35b3480156100cc57600080fd5b50610075610211565b3480156100e157600080fd5b506100756100f0366004610724565b610246565b34801561010157600080fd5b5061007561011036600461086e565b610321565b6100756101233660046109ce565b61035d565b34801561013457600080fd5b50610075610143366004610aba565b6103a1565b6101506103d6565b61017273ffffffffffffffffffffffffffffffffffffffff8316338386610419565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101d96001600055565b505050565b6040516060907f3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee90600090a19392505050565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61024e6103d6565b600061025b600285610ba4565b905080600003610297576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102b973ffffffffffffffffffffffffffffffffffffffff8416338484610419565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101d96001600055565b7f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b3382604051610352929190610c28565b60405180910390a150565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa3334858585604051610394959493929190610d1a565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146338484846040516103949493929190610da4565b600260005403610412576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526104ae9085906104b4565b50505050565b60006104d673ffffffffffffffffffffffffffffffffffffffff84168361054f565b905080516000141580156104fb5750808060200190518101906104f99190610e67565b155b156101d9576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b606061055d83836000610564565b9392505050565b6060814710156105a2576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610546565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105cb9190610e84565b60006040518083038185875af1925050503d8060008114610608576040519150601f19603f3d011682016040523d82523d6000602084013e61060d565b606091505b509150915061061d868383610627565b9695505050505050565b60608261063c57610637826106b6565b61055d565b8151158015610660575073ffffffffffffffffffffffffffffffffffffffff84163b155b156106af576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610546565b508061055d565b8051156106c65780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff8116811461071f57600080fd5b919050565b60008060006060848603121561073957600080fd5b83359250610749602085016106fb565b9150610757604085016106fb565b90509250925092565b6000806000838503604081121561077657600080fd5b602081121561078457600080fd5b50839250602084013567ffffffffffffffff8111156107a257600080fd5b8401601f810186136107b357600080fd5b803567ffffffffffffffff8111156107ca57600080fd5b8660208284010111156107dc57600080fd5b939660209190910195509293505050565b60005b838110156108085781810151838201526020016107f0565b50506000910152565b600081518084526108298160208601602086016107ed565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061055d6020830184610811565b60006020828403121561088057600080fd5b813567ffffffffffffffff81111561089757600080fd5b82016080818503121561055d57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561091f5761091f6108a9565b604052919050565b600082601f83011261093857600080fd5b813567ffffffffffffffff811115610952576109526108a9565b61098360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016108d8565b81815284602083860101111561099857600080fd5b816020850160208301376000918101602001919091529392505050565b80151581146106f857600080fd5b803561071f816109b5565b6000806000606084860312156109e357600080fd5b833567ffffffffffffffff8111156109fa57600080fd5b610a0686828701610927565b935050602084013591506040840135610a1e816109b5565b809150509250925092565b600067ffffffffffffffff821115610a4357610a436108a9565b5060051b60200190565b600082601f830112610a5e57600080fd5b8135610a71610a6c82610a29565b6108d8565b8082825260208201915060208360051b860101925085831115610a9357600080fd5b602085015b83811015610ab0578035835260209283019201610a98565b5095945050505050565b600080600060608486031215610acf57600080fd5b833567ffffffffffffffff811115610ae657600080fd5b8401601f81018613610af757600080fd5b8035610b05610a6c82610a29565b8082825260208201915060208360051b850101925088831115610b2757600080fd5b602084015b83811015610b6957803567ffffffffffffffff811115610b4b57600080fd5b610b5a8b602083890101610927565b84525060209283019201610b2c565b509550505050602084013567ffffffffffffffff811115610b8957600080fd5b610b9586828701610a4d565b925050610757604085016109c3565b600082610bda577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610c66836106fb565b16604082015273ffffffffffffffffffffffffffffffffffffffff610c8d602084016106fb565b166060820152600080604084013590508060808401525060608301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610cd957600080fd5b830160208101903567ffffffffffffffff811115610cf657600080fd5b803603821315610d0557600080fd5b608060a085015261061d60c085018284610bdf565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610d4f60a0830186610811565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610d9a578151865260209586019590910190600101610d7c565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610e37577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610e22858351610811565b94506020938401939190910190600101610de8565b505050508281036040840152610e4d8186610d68565b915050610e5e606083018415159052565b95945050505050565b600060208284031215610e7957600080fd5b815161055d816109b5565b60008251610e968184602087016107ed565b919091019291505056fea2646970667358221220db940ec4fecaa3b11f71d6cf6ef0ac3141552c5e39d750c1e2295a761a9396ad64736f6c634300081a0033a2646970667358221220b1685e295c918cb24bb7e0a00b1912bd08f8bb8204511c28c2c1400683b36d1664736f6c634300081a0033",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.DEFAULTADMINROLE(&_ERC20CustodyTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20CustodyTest *ERC20CustodyTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20CustodyTest.Contract.DEFAULTADMINROLE(&_ERC20CustodyTest.CallOpts)
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

// ERC20CustodyTestReceivedOnCallIterator is returned from FilterReceivedOnCall and is used to iterate over the raw logs and unpacked data for ReceivedOnCall events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReceivedOnCallIterator struct {
	Event *ERC20CustodyTestReceivedOnCall // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestReceivedOnCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestReceivedOnCall)
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
		it.Event = new(ERC20CustodyTestReceivedOnCall)
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
func (it *ERC20CustodyTestReceivedOnCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestReceivedOnCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestReceivedOnCall represents a ReceivedOnCall event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestReceivedOnCall struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCall is a free log retrieval operation binding the contract event 0x3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee.
//
// Solidity: event ReceivedOnCall()
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterReceivedOnCall(opts *bind.FilterOpts) (*ERC20CustodyTestReceivedOnCallIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestReceivedOnCallIterator{contract: _ERC20CustodyTest.contract, event: "ReceivedOnCall", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCall is a free log subscription operation binding the contract event 0x3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee.
//
// Solidity: event ReceivedOnCall()
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchReceivedOnCall(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestReceivedOnCall) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestReceivedOnCall)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseReceivedOnCall(log types.Log) (*ERC20CustodyTestReceivedOnCall, error) {
	event := new(ERC20CustodyTestReceivedOnCall)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
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

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*ERC20CustodyTestReceivedRevertIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestReceivedRevertIterator{contract: _ERC20CustodyTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
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

// ParseReceivedRevert is a log parse operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
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

// FilterReverted is a free log retrieval operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
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

// WatchReverted is a free log subscription operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
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

// ParseReverted is a log parse operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
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

// ERC20CustodyTestUpdatedCustodyTSSAddressIterator is returned from FilterUpdatedCustodyTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedCustodyTSSAddress events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestUpdatedCustodyTSSAddressIterator struct {
	Event *ERC20CustodyTestUpdatedCustodyTSSAddress // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestUpdatedCustodyTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestUpdatedCustodyTSSAddress)
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
		it.Event = new(ERC20CustodyTestUpdatedCustodyTSSAddress)
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
func (it *ERC20CustodyTestUpdatedCustodyTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestUpdatedCustodyTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestUpdatedCustodyTSSAddress represents a UpdatedCustodyTSSAddress event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestUpdatedCustodyTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCustodyTSSAddress is a free log retrieval operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterUpdatedCustodyTSSAddress(opts *bind.FilterOpts) (*ERC20CustodyTestUpdatedCustodyTSSAddressIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestUpdatedCustodyTSSAddressIterator{contract: _ERC20CustodyTest.contract, event: "UpdatedCustodyTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedCustodyTSSAddress is a free log subscription operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchUpdatedCustodyTSSAddress(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestUpdatedCustodyTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestUpdatedCustodyTSSAddress)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
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

// ParseUpdatedCustodyTSSAddress is a log parse operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseUpdatedCustodyTSSAddress(log types.Log) (*ERC20CustodyTestUpdatedCustodyTSSAddress, error) {
	event := new(ERC20CustodyTestUpdatedCustodyTSSAddress)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestUpdatedGatewayTSSAddressIterator struct {
	Event *ERC20CustodyTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestUpdatedGatewayTSSAddress)
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
		it.Event = new(ERC20CustodyTestUpdatedGatewayTSSAddress)
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
func (it *ERC20CustodyTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestUpdatedGatewayTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*ERC20CustodyTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestUpdatedGatewayTSSAddressIterator{contract: _ERC20CustodyTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestUpdatedGatewayTSSAddress)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
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
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*ERC20CustodyTestUpdatedGatewayTSSAddress, error) {
	event := new(ERC20CustodyTestUpdatedGatewayTSSAddress)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
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

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
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

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
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

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
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
