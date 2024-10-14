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
	ABI: "[{\"type\":\"function\",\"name\":\"ASSET_HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WHITELISTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testDepositLegacy\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositLegacyFailsIfNotSupported\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositLegacyFailsIfTokenNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20PartialThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveERC20ThroughCustodyTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNoParamsThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgrade\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfSenderIsNotTSSUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUnwhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUnwhitelistFailsIfSenderIsNotWhitelister\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUnwhitelistFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpgradeAndWithdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhitelistFailsIfSenderIsNotWhitelister\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhitelistFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawThroughCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedCustodyTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LegacyMethodsNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5061fa4f8061003c6000396000f3fe608060405234801561001057600080fd5b50600436106103365760003560e01c806385f438c1116101b2578063b5508aa9116100f9578063eb1ce7f9116100a2578063fa2a70741161007c578063fa2a7074146105b0578063fa7626d4146105b8578063fb176c12146105c5578063fe8e5f1b146105cd57600080fd5b8063eb1ce7f914610598578063f0c8e7e0146105a0578063f4221f08146105a857600080fd5b8063cbd57e2f116100d3578063cbd57e2f14610561578063e20c9f7114610569578063e63ab1e91461057157600080fd5b8063b5508aa914610539578063ba414fa614610541578063c713f8271461055957600080fd5b8063a3f9d0e01161015b578063af298bb111610135578063af298bb114610521578063b0464fdc14610529578063b421ca701461053157600080fd5b8063a3f9d0e0146104ea578063a4943deb146104f2578063a783c789146104fa57600080fd5b80639918c1c21161018c5780639918c1c2146104d25780639fc7fd55146104da578063a217fddf146104e257600080fd5b806385f438c11461048e5780639158c623146104b5578063916a17c6146104bd57600080fd5b806349c783dd1161028157806366d9a9a01161022a57806371149c941161020457806371149c94146104615780637e91c50f1461046957806382c529921461047157806385226c811461047957600080fd5b806366d9a9a01461043c5780636a621854146104515780637099d6f81461045957600080fd5b806352ff59391161025b57806352ff5939146103d8578063570618e1146103e05780635d62c8601461041557600080fd5b806349c783dd146103c05780634df42da1146103c857806351ecdf3c146103d057600080fd5b80632ade3880116102e35780633e73ecb4116102bd5780633e73ecb4146103a85780633ee92923146103b05780633f7286f4146103b857600080fd5b80632ade3880146103835780632be6a162146103985780633e5e3c23146103a057600080fd5b80631779672f116103145780631779672f146103555780631ed7831c1461035d578063284cb9291461037b57600080fd5b8063070f2ad01461033b5780630a9254e4146103455780630eee72a91461034d575b600080fd5b6103436105d5565b005b6103436107c1565b610343611248565b6103436114c5565b610365611603565b604051610372919061c4b3565b60405180910390f35b610343611665565b61038b611946565b604051610372919061c54f565b610343611a88565b610365611c34565b610343611c94565b6103436121e7565b6103656122a3565b610343612303565b610343612640565b610343612798565b610343612956565b6104077f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b604051908152602001610372565b6104077f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b981565b61044461317b565b604051610372919061c6b5565b6103436132e8565b6103436133a1565b61034361364f565b610343613e5c565b610343613fe7565b610481614259565b604051610372919061c753565b6104077f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b610343614329565b6104c56143f7565b604051610372919061c7ca565b6103436144dd565b6103436147c1565b610407600081565b61034361488f565b610343614e6e565b6104077f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b610343615081565b6104c5615625565b61034361570b565b610481615c41565b610549615d11565b6040519015158152602001610372565b610343615de5565b610343616a36565b610365616ad4565b6104077f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b610343616b34565b610343616c40565b610343616dec565b610343617043565b601f546105499060ff1681565b610343617324565b61034361797b565b6027546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561064757600080fd5b505af115801561065b573d6000803e3d6000fd5b5050602754604080516001600160a01b039092166024830152600060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250610712919060040161c861565b600060405180830381600087803b15801561072c57600080fd5b505af1158015610740573d6000803e3d6000fd5b50506021546025546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063950837aa91506024015b600060405180830381600087803b1580156107a757600080fd5b505af11580156107bb573d6000803e3d6000fd5b50505050565b602580547fffffffffffffffffffffffff000000000000000000000000000000000000000090811630179091556026805482166112341790556027805482166156781790556028805490911661987617905560405161081f9061c3e0565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156108a4573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03929092169190911790556040516108e99061c3e0565b604080825260049082018190527f7a6574610000000000000000000000000000000000000000000000000000000060608301526080602083018190528201527f5a4554410000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f08015801561096d573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081178255604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260275460255492519086169481019490945260448401929092529092166064820152600091610a4c916084015b60408051601f198184030181529190526020810180516001600160e01b03167fc0c53b8b00000000000000000000000000000000000000000000000000000000179052617b54565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0384811682029290921792839055604080518082018252601081527f4552433230437573746f64792e736f6c0000000000000000000000000000000060208201526027546025549251939095048416602484015293831660448301529091166064820152919250610aef91608401610a04565b602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691909117909155604080518082018252601a81527f5a657461436f6e6e6563746f724e6f6e4e61746976652e736f6c0000000000006020820152601f546024805460275460255495516101009094048716928401929092528516604483015284166064820152919092166084820152919250610be19160a40160408051601f198184030181529190526020810180516001600160e01b03167ff8c8765e00000000000000000000000000000000000000000000000000000000179052617b54565b602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038316179055604051909150610c239061c3ed565b604051809103906000f080158015610c3f573d6000803e3d6000fd5b50602080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556027546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b158015610ceb57600080fd5b505af1158015610cff573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015610d7557600080fd5b505af1158015610d89573d6000803e3d6000fd5b5050601f546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015261010090920416925063ae7a3a6f9150602401600060405180830381600087803b158015610df457600080fd5b505af1158015610e08573d6000803e3d6000fd5b5050601f546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526101009092041692506310188aef9150602401600060405180830381600087803b158015610e7357600080fd5b505af1158015610e87573d6000803e3d6000fd5b50506021546023546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639b19251a9150602401600060405180830381600087803b158015610eed57600080fd5b505af1158015610f01573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610f6357600080fd5b505af1158015610f77573d6000803e3d6000fd5b50506023546025546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42406024820152911692506340c10f199150604401600060405180830381600087803b158015610fe657600080fd5b505af1158015610ffa573d6000803e3d6000fd5b5050602480546025546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240938101939093521692506340c10f199150604401600060405180830381600087803b15801561106b57600080fd5b505af115801561107f573d6000803e3d6000fd5b50506023546021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526207a12060248201529116925063a9059cbb91506044016020604051808303816000875af11580156110f3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611117919061c874565b506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b15801561119857600080fd5b505af11580156111ac573d6000803e3d6000fd5b5050604080516080810182526025546001600160a01b039081168252602354811660208084019182526001848601908152855191820190955260008152606084018190528351602980549185167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161781559251602a8054919095169116179092559251602b55909350909150602c906107bb908261c95e565b602354602654604051620186a0602482018190526001600160a01b03938416604483015292909116606482015260009060840160408051601f198184030181529181526020820180516001600160e01b03167fc513169100000000000000000000000000000000000000000000000000000000179052602554905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b15801561131a57600080fd5b505af115801561132e573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611404919060040161c861565b600060405180830381600087803b15801561141e57600080fd5b505af1158015611432573d6000803e3d6000fd5b50506021546020546023546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f2945061148f9392831692909116908790879060040161ca1d565b600060405180830381600087803b1580156114a957600080fd5b505af11580156114bd573d6000803e3d6000fd5b505050505050565b6023546026546040516000602482018190526001600160a01b039384166044830152929091166064820152819060840160408051601f198184030181529181526020820180516001600160e01b03167f357fc5a200000000000000000000000000000000000000000000000000000000179052602754905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b15801561159457600080fd5b505af11580156115a8573d6000803e3d6000fd5b5050604051630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401611404565b6060601680548060200260200160405190810160405280929190818152602001828054801561165b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161163d575b5050505050905090565b602154602480546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600093919091169163d936547e9101602060405180830381865afa1580156116ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f2919061c874565b90506116ff600082617b73565b6021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561177457600080fd5b505af1158015611788573d6000803e3d6000fd5b50506024546040516001600160a01b0390911692507faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549150600090a260255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561181d57600080fd5b505af1158015611831573d6000803e3d6000fd5b5050602154602480546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639b19251a925001600060405180830381600087803b15801561189657600080fd5b505af11580156118aa573d6000803e3d6000fd5b5050602154602480546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529216935063d936547e925001602060405180830381865afa158015611912573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611936919061c874565b9050611943600182617b73565b50565b6060601e805480602002602001604051908101604052809291908181526020016000905b82821015611a7f57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015611a685783829060005260206000200180546119db9061c8c5565b80601f0160208091040260200160405190810160405280929190818152602001828054611a079061c8c5565b8015611a545780601f10611a2957610100808354040283529160200191611a54565b820191906000526020600020905b815481529060010190602001808311611a3757829003601f168201915b5050505050815260200190600101906119bc565b50505050815250508152602001906001019061196a565b50505050905090565b60405163ca669fa760e01b81526101236004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611ad657600080fd5b505af1158015611aea573d6000803e3d6000fd5b50506040805161012360248201527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60448083019190915282518083039091018152606490910182526020810180516001600160e01b03167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611bb6919060040161c861565b600060405180830381600087803b158015611bd057600080fd5b505af1158015611be4573d6000803e3d6000fd5b5050602154602480546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639b19251a92500161078d565b6060601880548060200260200160405190810160405280929190818152602001828054801561165b576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161163d575050505050905090565b6023546026546040516370a0823160e01b81526001600160a01b039182166004820152620186a09260009216906370a0823190602401602060405180830381865afa158015611ce7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d0b919061ca54565b9050611d18816000617bf5565b6023546021546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611d69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d8d919061ca54565b602654604080516001600160a01b039283166024820152604480820188905282518083039091018152606490910182526020810180516001600160e01b03167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba392611e4c92911690600090869060040161ca6d565b600060405180830381600087803b158015611e6657600080fd5b505af1158015611e7a573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015611ef357600080fd5b505af1158015611f07573d6000803e3d6000fd5b50506023546026546040518881526001600160a01b039283169450911691507fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9060200160405180910390a360275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611fac57600080fd5b505af1158015611fc0573d6000803e3d6000fd5b50506021546026546023546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529082166024820152604481018990529116925063d9caed129150606401600060405180830381600087803b15801561203857600080fd5b505af115801561204c573d6000803e3d6000fd5b50506023546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561209f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120c3919061ca54565b90506120cf8186617bf5565b6023546021546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015612120573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612144919061ca54565b905061215981612154888761cac4565b617bf5565b602354601f546040516370a0823160e01b81526101009091046001600160a01b03908116600483015260009216906370a0823190602401602060405180830381865afa1580156121ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121d1919061ca54565b90506121de816000617bf5565b50505050505050565b602354602654604051620186a0602482018190526001600160a01b03938416604483015292909116606482015260009060840160408051601f198184030181529181526020820180516001600160e01b03167f357fc5a200000000000000000000000000000000000000000000000000000000179052602554905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401611300565b6060601780548060200260200160405190810160405280929190818152602001828054801561165b576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161163d575050505050905090565b6021546040517feab103df000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b039091169063eab103df90602401600060405180830381600087803b15801561236257600080fd5b505af1158015612376573d6000803e3d6000fd5b50506023546021546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424060248201529116925063095ea7b391506044016020604051808303816000875af11580156123ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061240e919061c874565b50602480546021546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424093810193909352169063095ea7b3906044016020604051808303816000875af1158015612481573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124a5919061c874565b50604080517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152815160058183030181526025820192839052630618f58760e51b9092527f73cba663000000000000000000000000000000000000000000000000000000006029820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906049015b600060405180830381600087803b15801561254d57600080fd5b505af1158015612561573d6000803e3d6000fd5b505060215460265460405160609190911b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201526001600160a01b03909116925063e609055e915060340160408051601f19818403018152908290526023547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261260b926001600160a01b03909116906103e890879060040161cad7565b600060405180830381600087803b15801561262557600080fd5b505af1158015612639573d6000803e3d6000fd5b5050505050565b6025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156126b257600080fd5b505af11580156126c6573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561273657600080fd5b505af115801561274a573d6000803e3d6000fd5b50506021546040517f950837aa000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b03909116925063950837aa915060240161078d565b6023546026546040516001602482018190526001600160a01b03938416604483015292909116606482015260009060840160408051601f198184030181529181526020820180516001600160e01b03167f357fc5a200000000000000000000000000000000000000000000000000000000179052602754905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561286757600080fd5b505af115801561287b573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156128eb57600080fd5b505af11580156128ff573d6000803e3d6000fd5b50506021546023546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321fc65f2935061148f9260009216908790879060040161ca1d565b6021546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4600482015261432160248201819052916000916001600160a01b03909116906391d1485490604401602060405180830381865afa1580156129e5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a09919061c874565b9050612a1481617c4d565b6021546040517f91d148540000000000000000000000000000000000000000000000000000000081527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60048201526001600160a01b03848116602483015260009216906391d1485490604401602060405180830381865afa158015612a9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ac2919061c874565b9050612acd81617c4d565b6021546027546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460048201526001600160a01b03918216602482015260009291909116906391d1485490604401602060405180830381865afa158015612b5d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b81919061c874565b9050612b8c81617cc7565b6021546027546040517f91d148540000000000000000000000000000000000000000000000000000000081527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60048201526001600160a01b03918216602482015260009291909116906391d1485490604401602060405180830381865afa158015612c1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c40919061c874565b9050612c4b81617cc7565b6025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015612cbd57600080fd5b505af1158015612cd1573d6000803e3d6000fd5b50506021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015612d4a57600080fd5b505af1158015612d5e573d6000803e3d6000fd5b50506040516001600160a01b03881681527f086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba9250602001905060405180910390a16021546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0387811660048301529091169063950837aa90602401600060405180830381600087803b158015612dff57600080fd5b505af1158015612e13573d6000803e3d6000fd5b50505050612e9785602160009054906101000a90046001600160a01b03166001600160a01b0316635b1125916040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e92919061cb11565b617d19565b6021546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460048201526001600160a01b038781166024830152909116906391d1485490604401602060405180830381865afa158015612f20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f44919061c874565b9350612f4f84617cc7565b6021546040517f91d148540000000000000000000000000000000000000000000000000000000081527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60048201526001600160a01b038781166024830152909116906391d1485490604401602060405180830381865afa158015612fd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ffc919061c874565b925061300783617cc7565b6021546027546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460048201526001600160a01b0391821660248201529116906391d1485490604401602060405180830381865afa158015613092573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130b6919061c874565b91506130c182617c4d565b6021546027546040517f91d148540000000000000000000000000000000000000000000000000000000081527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60048201526001600160a01b0391821660248201529116906391d1485490604401602060405180830381865afa15801561314c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613170919061c874565b905061263981617c4d565b6060601b805480602002602001604051908101604052809291908181526020016000905b82821015611a7f57838290600052602060002090600202016040518060400160405290816000820180546131d29061c8c5565b80601f01602080910402602001604051908101604052809291908181526020018280546131fe9061c8c5565b801561324b5780601f106132205761010080835404028352916020019161324b565b820191906000526020600020905b81548152906001019060200180831161322e57829003601f168201915b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156132d057602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116132925790505b5050505050815250508152602001906001019061319f565b6023546026546040516000602482018190526001600160a01b039384166044830152929091166064820152819060840160408051601f198184030181529181526020820180516001600160e01b03167fc513169100000000000000000000000000000000000000000000000000000000179052602754905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa79060240161157a565b6021546040517feab103df000000000000000000000000000000000000000000000000000000008152600160048201526001600160a01b039091169063eab103df90602401600060405180830381600087803b15801561340057600080fd5b505af1158015613414573d6000803e3d6000fd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a5904279150602401600060405180830381600087803b15801561347a57600080fd5b505af115801561348e573d6000803e3d6000fd5b50506023546021546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424060248201529116925063095ea7b391506044016020604051808303816000875af1158015613502573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613526919061c874565b50602480546021546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424093810193909352169063095ea7b3906044016020604051808303816000875af1158015613599573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135bd919061c874565b50604080517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152815160058183030181526025820192839052630618f58760e51b9092527f584a7938000000000000000000000000000000000000000000000000000000006029820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090604901612533565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a09060009060250160408051808303601f19018152908290526023546020546370a0823160e01b84526001600160a01b0390811660048501529193506000929116906370a0823190602401602060405180830381865afa1580156136e3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613707919061ca54565b9050613714816000617bf5565b6023546021546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015613765573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613789919061ca54565b601f54604080516001600160a01b0361010090930483166024820152604480820189905282518083039091018152606490910182526020810180517fa9059cbb000000000000000000000000000000000000000000000000000000006001600160e01b0390911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba39261384f92911690600090869060040161ca6d565b600060405180830381600087803b15801561386957600080fd5b505af115801561387d573d6000803e3d6000fd5b50506020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156138f657600080fd5b505af115801561390a573d6000803e3d6000fd5b505050507f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b601f60019054906101000a90046001600160a01b0316602960405161395592919061cc17565b60405180910390a1601f546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156139d657600080fd5b505af11580156139ea573d6000803e3d6000fd5b50506023546020546040516001600160a01b039283169450911691507fde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e03590613a38908990899060299061cc39565b60405180910390a36021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015613ab557600080fd5b505af1158015613ac9573d6000803e3d6000fd5b50506023546020546040516001600160a01b039283169450911691507f7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb972190613b17908990899060299061cc39565b60405180910390a360275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613b7857600080fd5b505af1158015613b8c573d6000803e3d6000fd5b50506021546020546023546040517f99a3c3560000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506399a3c3569450613bec9392831692909116908a908a9060299060040161cc64565b600060405180830381600087803b158015613c0657600080fd5b505af1158015613c1a573d6000803e3d6000fd5b50506023546020546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a08231906024015b602060405180830381865afa158015613c6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c92919061ca54565b9050613c9e8187617bf5565b6023546021546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015613cef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d13919061ca54565b9050613d2381612154898761cac4565b602354601f546020546040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081526101009092046001600160a01b0390811660048401529081166024830152600092169063dd62ed3e90604401602060405180830381865afa158015613d9b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613dbf919061ca54565b9050613dcc816000617bf5565b602354601f546040516370a0823160e01b81526101009091046001600160a01b03908116600483015260009216906370a0823190602401602060405180830381865afa158015613e20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e44919061ca54565b9050613e51816000617bf5565b505050505050505050565b6040517f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260019060009060250160408051808303601f190181529082905260275463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613ef557600080fd5b505af1158015613f09573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015613f7957600080fd5b505af1158015613f8d573d6000803e3d6000fd5b50506021546023546040517f99a3c3560000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506399a3c356935061148f9260009216908790879060299060040161cc64565b6027546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561405957600080fd5b505af115801561406d573d6000803e3d6000fd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a5904279150602401600060405180830381600087803b1580156140d357600080fd5b505af11580156140e7573d6000803e3d6000fd5b5050604051630618f58760e51b81527f584a7938000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561415757600080fd5b505af115801561416b573d6000803e3d6000fd5b50506021546026546023546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529082166024820152600160448201529116925063d9caed129150606401600060405180830381600087803b1580156141e357600080fd5b505af11580156141f7573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156107a757600080fd5b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015611a7f57838290600052602060002001805461429c9061c8c5565b80601f01602080910402602001604051908101604052809291908181526020018280546142c89061c8c5565b80156143155780601f106142ea57610100808354040283529160200191614315565b820191906000526020600020905b8154815290600101906020018083116142f857829003601f168201915b50505050508152602001906001019061427d565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561439557600080fd5b505af11580156143a9573d6000803e3d6000fd5b50506021546040517f9a590427000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b039091169250639a590427915060240161078d565b6060601d805480602002602001604051908101604052809291908181526020016000905b82821015611a7f5760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156144c557602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116144875790505b5050505050815250508152602001906001019061441b565b602354602654604051600160248201526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180516001600160e01b03167f357fc5a20000000000000000000000000000000000000000000000000000000017905260275490517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156145c157600080fd5b505af11580156145d5573d6000803e3d6000fd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a5904279150602401600060405180830381600087803b15801561463b57600080fd5b505af115801561464f573d6000803e3d6000fd5b5050604051630618f58760e51b81527f584a7938000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156146bf57600080fd5b505af11580156146d3573d6000803e3d6000fd5b50506021546020546023546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f29450614731939283169290911690600190879060040161ca1d565b600060405180830381600087803b15801561474b57600080fd5b505af115801561475f573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561262557600080fd5b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561482d57600080fd5b505af1158015614841573d6000803e3d6000fd5b50506021546040517f9b19251a000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b039091169250639b19251a915060240161078d565b604080516004808252602480830184526020830180516001600160e01b03167f6ed701690000000000000000000000000000000000000000000000000000000017905260235460265494516370a0823160e01b81526001600160a01b0395861693810193909352620186a0946000939116916370a082319101602060405180830381865afa158015614925573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614949919061ca54565b9050614956816000617bf5565b6023546021546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156149a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906149cb919061ca54565b601f54604080516001600160a01b0361010090930483166024820152604480820189905282518083039091018152606490910182526020810180517fa9059cbb000000000000000000000000000000000000000000000000000000006001600160e01b0390911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba392614a9192911690600090869060040161ca6d565b600060405180830381600087803b158015614aab57600080fd5b505af1158015614abf573d6000803e3d6000fd5b50506020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015614b3857600080fd5b505af1158015614b4c573d6000803e3d6000fd5b5050601f546040516101009091046001600160a01b031681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09250602001905060405180910390a16021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015614c0a57600080fd5b505af1158015614c1e573d6000803e3d6000fd5b50506023546020546040516001600160a01b039283169450911691507f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d590614c69908990899061ccb9565b60405180910390a360275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614cca57600080fd5b505af1158015614cde573d6000803e3d6000fd5b50506021546020546023546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f29450614d3b9392831692909116908a908a9060040161ca1d565b600060405180830381600087803b158015614d5557600080fd5b505af1158015614d69573d6000803e3d6000fd5b50506023546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015614dbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614de0919061ca54565b9050614ded816000617bf5565b6023546021546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015614e3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614e62919061ca54565b9050613d238185617bf5565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a09060009060250160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614f0957600080fd5b505af1158015614f1d573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614ff3919060040161c861565b600060405180830381600087803b15801561500d57600080fd5b505af1158015615021573d6000803e3d6000fd5b50506021546020546023546040517f99a3c3560000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506399a3c356945061148f9392831692909116908790879060299060040161cc64565b602154604080518082018252601b81527f4552433230437573746f647955706772616465546573742e736f6c00000000006020808301919091528251908101909252600082526025546150e1936001600160a01b03908116939116617d7a565b6021546023546026546040516370a0823160e01b81526001600160a01b03918216600482015292811692620186a09260009216906370a0823190602401602060405180830381865afa15801561513b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061515f919061ca54565b905061516c816000617bf5565b6023546040516370a0823160e01b81526001600160a01b03858116600483015260009216906370a0823190602401602060405180830381865afa1580156151b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906151db919061ca54565b602654604080516001600160a01b039283166024820152604480820188905282518083039091018152606490910182526020810180516001600160e01b03167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba39261529a92911690600090869060040161ca6d565b600060405180830381600087803b1580156152b457600080fd5b505af11580156152c8573d6000803e3d6000fd5b50506040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b0388166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561533d57600080fd5b505af1158015615351573d6000803e3d6000fd5b50506023546026546040518881526001600160a01b039283169450911691507fd4dabfe72081670cc78f2ebda8e2eddaf3feebde6288dcb8fe673b3dc201b5a49060200160405180910390a360275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156153f657600080fd5b505af115801561540a573d6000803e3d6000fd5b50506026546023546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152908216602482015260448101889052908816925063d9caed129150606401600060405180830381600087803b15801561548057600080fd5b505af1158015615494573d6000803e3d6000fd5b50506023546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156154e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061550b919061ca54565b90506155178186617bf5565b6023546040516370a0823160e01b81526001600160a01b03888116600483015260009216906370a0823190602401602060405180830381865afa158015615562573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615586919061ca54565b905061559681612154888761cac4565b602354601f546040516370a0823160e01b81526101009091046001600160a01b03908116600483015260009216906370a0823190602401602060405180830381865afa1580156155ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061560e919061ca54565b905061561b816000617bf5565b5050505050505050565b6060601c805480602002602001604051908101604052809291908181526020016000905b82821015611a7f5760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156156f357602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116156b55790505b50505050508152505081526020019060010190615649565b6021546040517feab103df000000000000000000000000000000000000000000000000000000008152600160048201526103e8916001600160a01b03169063eab103df90602401600060405180830381600087803b15801561576c57600080fd5b505af1158015615780573d6000803e3d6000fd5b50506023546021546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424060248201529116925063095ea7b391506044016020604051808303816000875af11580156157f4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615818919061c874565b50602480546021546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424093810193909352169063095ea7b3906044016020604051808303816000875af115801561588b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906158af919061c874565b506023546021546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015615901573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615925919061ca54565b602480546027546040516370a0823160e01b81526001600160a01b0391821660048201529394506000939116916370a082319101602060405180830381865afa158015615976573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061599a919061ca54565b905060006040516020016159d1907f68656c6c6f000000000000000000000000000000000000000000000000000000815260050190565b60408051808303601f19018152908290526021546381bad6f360e01b8352600160048401819052602484018190526044840181905260648401526001600160a01b031660848301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015615a5457600080fd5b505af1158015615a68573d6000803e3d6000fd5b505060235460265460405160609190911b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201526001600160a01b0390911692507f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae915060340160408051601f1981840301815290829052615af1918890869061ccd2565b60405180910390a26021546026546040805160609290921b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208301528051808303601401815260348301918290526023547fe609055e000000000000000000000000000000000000000000000000000000009092526001600160a01b039384169363e609055e93615b909391909116908990879060380161cad7565b600060405180830381600087803b158015615baa57600080fd5b505af1158015615bbe573d6000803e3d6000fd5b505050506107bb8484615bd1919061ccfd565b6023546021546040516370a0823160e01b81526001600160a01b0391821660048201529116906370a0823190602401602060405180830381865afa158015615c1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612154919061ca54565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015611a7f578382906000526020600020018054615c849061c8c5565b80601f0160208091040260200160405190810160405280929190818152602001828054615cb09061c8c5565b8015615cfd5780601f10615cd257610100808354040283529160200191615cfd565b820191906000526020600020905b815481529060010190602001808311615ce057829003601f168201915b505050505081526020019060010190615c65565b60085460009060ff1615615d29575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015615dba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615dde919061ca54565b1415905090565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015615e3e57600080fd5b505af1158015615e52573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250615f28919060040161c861565b600060405180830381600087803b158015615f4257600080fd5b505af1158015615f56573d6000803e3d6000fd5b50505050602160009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b158015615faa57600080fd5b505af1158015615fbe573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561601b57600080fd5b505af115801561602f573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250616105919060040161c861565b600060405180830381600087803b15801561611f57600080fd5b505af1158015616133573d6000803e3d6000fd5b50505050601f60019054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561618757600080fd5b505af115801561619b573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156161f857600080fd5b505af115801561620c573d6000803e3d6000fd5b50505050602160009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561626057600080fd5b505af1158015616274573d6000803e3d6000fd5b5050602354602654604051620186a0602482018190526001600160a01b0393841660448301529290911660648201529092506000915060840160408051601f198184030181529181526020820180516001600160e01b03167f357fc5a20000000000000000000000000000000000000000000000000000000017905251630618f58760e51b81527fd93c0665000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561635d57600080fd5b505af1158015616371573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156163ce57600080fd5b505af11580156163e2573d6000803e3d6000fd5b50506021546020546023546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f2945061643f9392831692909116908790879060040161ca1d565b600060405180830381600087803b15801561645957600080fd5b505af115801561646d573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156164ca57600080fd5b505af11580156164de573d6000803e3d6000fd5b50505050602160009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561653257600080fd5b505af1158015616546573d6000803e3d6000fd5b50506023546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a08231906024015b602060405180830381865afa15801561659a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906165be919061ca54565b90506165cb816000617bf5565b6023546021546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa15801561661c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616640919061ca54565b601f54604080516001600160a01b0361010090930483166024820152604480820189905282518083039091018152606490910182526020810180517fa9059cbb000000000000000000000000000000000000000000000000000000006001600160e01b0390911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba39261670692911690600090869060040161ca6d565b600060405180830381600087803b15801561672057600080fd5b505af1158015616734573d6000803e3d6000fd5b50506020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156167ad57600080fd5b505af11580156167c1573d6000803e3d6000fd5b5050601f54602354602654604080516101009094046001600160a01b039081168552602085018c9052928316908401521660608201527f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609250608001905060405180910390a16021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561689c57600080fd5b505af11580156168b0573d6000803e3d6000fd5b50506023546020546040516001600160a01b039283169450911691507f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5906168fb908990899061ccb9565b60405180910390a360275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561695c57600080fd5b505af1158015616970573d6000803e3d6000fd5b50506021546020546023546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506169cd9392831692909116908a908a9060040161ca1d565b600060405180830381600087803b1580156169e757600080fd5b505af11580156169fb573d6000803e3d6000fd5b50506023546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401613c51565b60235460265460408051620186a060248083018290526001600160a01b039586166044840181905295909416606480840182905284518085039091018152608490930184526020830180516001600160e01b03167f357fc5a20000000000000000000000000000000000000000000000000000000017905292516370a0823160e01b815260048101939093529390926000926370a08231910161657d565b6060601580548060200260200160405190810160405280929190818152602001828054801561165b576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161163d575050505050905090565b600080604051602001616b6a907f68656c6c6f000000000000000000000000000000000000000000000000000000815260050190565b60408051808303601f190181529082905260275463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015616bd157600080fd5b505af1158015616be5573d6000803e3d6000fd5b5050604051630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401614ff3565b60405163ca669fa760e01b81526101236004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015616c8e57600080fd5b505af1158015616ca2573d6000803e3d6000fd5b50506040805161012360248201527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60448083019190915282518083039091018152606490910182526020810180516001600160e01b03167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250616d6e919060040161c861565b600060405180830381600087803b158015616d8857600080fd5b505af1158015616d9c573d6000803e3d6000fd5b5050602154602480546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292169350639a59042792500161078d565b602354602654604051600160248201526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180516001600160e01b03167f357fc5a20000000000000000000000000000000000000000000000000000000017905260275490517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015616ed057600080fd5b505af1158015616ee4573d6000803e3d6000fd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a5904279150602401600060405180830381600087803b158015616f4a57600080fd5b505af1158015616f5e573d6000803e3d6000fd5b5050604051630618f58760e51b81527f584a7938000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015616fce57600080fd5b505af1158015616fe2573d6000803e3d6000fd5b50506021546020546023546040517f99a3c3560000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506399a3c3569450614731939283169290911690600190879060299060040161cc64565b6021546023546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600092919091169063d936547e90602401602060405180830381865afa1580156170ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906170d1919061c874565b90506170de600182617b73565b6021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561715357600080fd5b505af1158015617167573d6000803e3d6000fd5b50506023546040516001600160a01b0390911692507f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919150600090a260255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156171fc57600080fd5b505af1158015617210573d6000803e3d6000fd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a5904279150602401600060405180830381600087803b15801561727657600080fd5b505af115801561728a573d6000803e3d6000fd5b50506021546023546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063d936547e9150602401602060405180830381865afa1580156172f3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617317919061c874565b9050611943600082617b73565b602354602654604051620186a0602482018190526001600160a01b03938416604483015292909116606482015260009060840160408051601f198184030181529181526020820180516001600160e01b03167fc51316910000000000000000000000000000000000000000000000000000000017905260235460265491516370a0823160e01b81526001600160a01b0392831660048201529293506000929116906370a0823190602401602060405180830381865afa1580156173eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061740f919061ca54565b905061741c816000617bf5565b6023546021546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa15801561746d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617491919061ca54565b601f54604080516001600160a01b0361010090930483166024820152604480820189905282518083039091018152606490910182526020810180517fa9059cbb000000000000000000000000000000000000000000000000000000006001600160e01b0390911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba39261755792911690600090869060040161ca6d565b600060405180830381600087803b15801561757157600080fd5b505af1158015617585573d6000803e3d6000fd5b50506020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156175fe57600080fd5b505af1158015617612573d6000803e3d6000fd5b5050601f547f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60925061010090046001600160a01b0316905061765560028861cd10565b602354602654604080516001600160a01b03958616815260208101949094529184168383015292909216606082015290519081900360800190a16021546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561770457600080fd5b505af1158015617718573d6000803e3d6000fd5b50506023546020546040516001600160a01b039283169450911691507f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d590617763908990899061ccb9565b60405180910390a360275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156177c457600080fd5b505af11580156177d8573d6000803e3d6000fd5b50506021546020546023546040517f21fc65f20000000000000000000000000000000000000000000000000000000081526001600160a01b0393841695506321fc65f294506178359392831692909116908a908a9060040161ca1d565b600060405180830381600087803b15801561784f57600080fd5b505af1158015617863573d6000803e3d6000fd5b50506023546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156178b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906178da919061ca54565b90506178eb8161215460028961cd10565b6023546021546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa15801561793c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617960919061ca54565b9050613d238161797160028a61cd10565b612154908761cac4565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156179d957600080fd5b505af11580156179ed573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250617ac3919060040161c861565b600060405180830381600087803b158015617add57600080fd5b505af1158015617af1573d6000803e3d6000fd5b50506021546026546023546040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529082166024820152604481018690529116925063d9caed12915060640161260b565b6000617b5e61c3fa565b617b69848483617d8f565b9150505b92915050565b6040517ff7fe347700000000000000000000000000000000000000000000000000000000815282151560048201528115156024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f7fe3477906044015b60006040518083038186803b158015617be157600080fd5b505afa1580156114bd573d6000803e3d6000fd5b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c5490604401617bc9565b6040517fa59828850000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a5982885906024015b60006040518083038186803b158015617cb357600080fd5b505afa158015612639573d6000803e3d6000fd5b6040517f0c9fd5810000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90630c9fd58190602401617c9b565b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f690604401617bc9565b617d8261c3fa565b6126398585858486617e0a565b600080617d9c8584617f0a565b9050617dff6040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f78790000008152508286604051602001617dea92919061cd4b565b60405160208183030381529060405285617f16565b9150505b9392505050565b6040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201528190737109709ecfa91a80626ff3989d68f67f5b1dd12d9081906306447d5690602401600060405180830381600087803b158015617e7c57600080fd5b505af1925050508015617e8d575060015b617ea257617e9d87878787617f44565b6121de565b617eae87878787617f44565b806001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015617ee957600080fd5b505af1158015617efd573d6000803e3d6000fd5b5050505050505050505050565b6000617e038383617f5d565b60c08101515160009015617f3a57617f3384848460c00151617f78565b9050617e03565b617f33848461811e565b6000617f508483618209565b9050612639858285618215565b6000617f6983836185df565b617e0383836020015184617f16565b600080617f836185ef565b90506000617f9186836186c2565b90506000617fa88260600151836020015185618b68565b90506000617fb883838989618d7a565b90506000617fc582619bf7565b602081015181519192509060030b1561803857898260400151604051602001617fef92919061cd6d565b60408051601f19818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261802f9160040161c861565b60405180910390fd5b600061807b6040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001619dc6565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906180ce90849060040161c861565b602060405180830381865afa1580156180eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061810f919061cb11565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc9259061817390879060040161c861565b600060405180830381865afa158015618190573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526181b8919081019061cea7565b905060006181e682856040516020016181d292919061cedc565b604051602081830303815290604052619fc6565b90506001600160a01b038116617b69578484604051602001617fef92919061cf0b565b6000617f698383619fd9565b6040517f667f9d700000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90600090829063667f9d7090604401602060405180830381865afa1580156182b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906182d5919061ca54565b90508061847c5760006182e786619fe5565b604080518082018252600581527f352e302e3000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150618372905b6040805180820182526000808252602091820152815180830190925284518252808501908201529061a0c8565b8061837e575060008451115b15618401576040517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b03871690634f1ef286906183ca908890889060040161cd4b565b600060405180830381600087803b1580156183e457600080fd5b505af11580156183f8573d6000803e3d6000fd5b50505050618476565b6040517f3659cfe60000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152871690633659cfe690602401600060405180830381600087803b15801561845d57600080fd5b505af1158015618471573d6000803e3d6000fd5b505050505b50612639565b80600061848882619fe5565b604080518082018252600581527f352e302e30000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506184ea90618345565b806184f6575060008551115b1561857b576040517f9623609d0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690639623609d90618544908a908a908a9060040161cfb6565b600060405180830381600087803b15801561855e57600080fd5b505af1158015618572573d6000803e3d6000fd5b505050506121de565b6040517f99a88ec40000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015287811660248301528316906399a88ec490604401600060405180830381600087803b158015617ee957600080fd5b6185eb8282600061a0dc565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c9061867690849060040161cfe7565b600060405180830381865afa158015618693573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526186bb919081019061d02e565b9250505090565b6186f46040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d905061873f6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6187488561a1df565b602082015260006187588661a5c4565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa15801561879a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526187c2919081019061d02e565b868385602001516040516020016187dc949392919061d077565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb119061883490859060040161c861565b600060405180830381865afa158015618851573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618879919081019061d02e565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906188c190849060040161d17b565b602060405180830381865afa1580156188de573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618902919061c874565b6189175781604051602001617fef919061d1cd565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061895c90849060040161d25f565b600060405180830381865afa158015618979573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526189a1919081019061d02e565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906189e890849060040161d2b1565b602060405180830381865afa158015618a05573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618a29919061c874565b15618abe576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890618a7390849060040161d2b1565b600060405180830381865afa158015618a90573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618ab8919081019061d02e565b60408501525b846001600160a01b03166349c4fac8828660000151604051602001618ae3919061d303565b6040516020818303038152906040526040518363ffffffff1660e01b8152600401618b0f92919061d36f565b600060405180830381865afa158015618b2c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618b54919081019061d02e565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b6060815260200190600190039081618b845790505090506040518060400160405280600481526020017f677265700000000000000000000000000000000000000000000000000000000081525081600081518110618be457618be461d394565b60200260200101819052506040518060400160405280600381526020017f2d726c000000000000000000000000000000000000000000000000000000000081525081600181518110618c3857618c3861d394565b602002602001018190525084604051602001618c54919061d3c3565b60405160208183030381529060405281600281518110618c7657618c7661d394565b602002602001018190525082604051602001618c92919061d42f565b60405160208183030381529060405281600381518110618cb457618cb461d394565b60200260200101819052506000618cca82619bf7565b602080820151604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000008185019081528251808401845260008082529086015282518084019093529051825292810192909252919250618d5b906040805180820182526000808252602091820152815180830190925284518252808501908201529061a847565b618d705785604051602001617fef919061d470565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015618dca565b511590565b618f3e57826020015115618e86576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a40161802f565b8260c0015115618f3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a40161802f565b6040805160ff8082526120008201909252600091816020015b6060815260200190600190039081618f5757905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280618fb29061d501565b935060ff1681518110618fc757618fc761d394565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e3700000000000000000000000000000000000000815250604051602001619018919061d520565b6040516020818303038152906040528282806190339061d501565b935060ff16815181106190485761904861d394565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806190959061d501565b935060ff16815181106190aa576190aa61d394565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806190f79061d501565b935060ff168151811061910c5761910c61d394565b602002602001018190525087602001518282806191289061d501565b935060ff168151811061913d5761913d61d394565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163745061746800000000000000000000000000000000000081525082828061918a9061d501565b935060ff168151811061919f5761919f61d394565b6020908102919091010152875182826191b78161d501565b935060ff16815181106191cc576191cc61d394565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e496400000000000000000000000000000000000000000000008152508282806192199061d501565b935060ff168151811061922e5761922e61d394565b60200260200101819052506192424661a8a8565b828261924d8161d501565b935060ff16815181106192625761926261d394565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806192af9061d501565b935060ff16815181106192c4576192c461d394565b6020026020010181905250868282806192dc9061d501565b935060ff16815181106192f1576192f161d394565b60209081029190910101528551156194185760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826193428161d501565b935060ff16815181106193575761935761d394565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906193a790899060040161c861565b600060405180830381865afa1580156193c4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526193ec919081019061d02e565b82826193f78161d501565b935060ff168151811061940c5761940c61d394565b60200260200101819052505b8460200151156194e85760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826194618161d501565b935060ff16815181106194765761947661d394565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806194c39061d501565b935060ff16815181106194d8576194d861d394565b60200260200101819052506196af565b619520618dc58660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6195b35760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826195638161d501565b935060ff16815181106195785761957861d394565b60200260200101819052508460a00151604051602001619598919061d3c3565b6040516020818303038152906040528282806194c39061d501565b8460c001511580156195f65750604080890151815180830183526000808252602091820152825180840190935281518352908101908201526195f490511590565b155b156196af5760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261963a8161d501565b935060ff168151811061964f5761964f61d394565b60200260200101819052506196638861a948565b604051602001619673919061d3c3565b60405160208183030381529060405282828061968e9061d501565b935060ff16815181106196a3576196a361d394565b60200260200101819052505b604080860151815180830183526000808252602091820152825180840190935281518352908101908201526196e390511590565b6197785760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826197268161d501565b935060ff168151811061973b5761973b61d394565b602002602001018190525084604001518282806197579061d501565b935060ff168151811061976c5761976c61d394565b60200260200101819052505b6060850151156198995760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826197c18161d501565b935060ff16815181106197d6576197d661d394565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015619845573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261986d919081019061d02e565b82826198788161d501565b935060ff168151811061988d5761988d61d394565b60200260200101819052505b60e085015151156199405760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826198e38161d501565b935060ff16815181106198f8576198f861d394565b60200260200101819052506199148560e001516000015161a8a8565b828261991f8161d501565b935060ff16815181106199345761993461d394565b60200260200101819052505b60e085015160200151156199ea5760408051808201909152600a81527f2d2d6761735072696365000000000000000000000000000000000000000000006020820152828261998d8161d501565b935060ff16815181106199a2576199a261d394565b60200260200101819052506199be8560e001516020015161a8a8565b82826199c98161d501565b935060ff16815181106199de576199de61d394565b60200260200101819052505b60e08501516040015115619a945760408051808201909152600e81527f2d2d6d617846656550657247617300000000000000000000000000000000000060208201528282619a378161d501565b935060ff1681518110619a4c57619a4c61d394565b6020026020010181905250619a688560e001516040015161a8a8565b8282619a738161d501565b935060ff1681518110619a8857619a8861d394565b60200260200101819052505b60e08501516060015115619b3e5760408051808201909152601681527f2d2d6d61785072696f726974794665655065724761730000000000000000000060208201528282619ae18161d501565b935060ff1681518110619af657619af661d394565b6020026020010181905250619b128560e001516060015161a8a8565b8282619b1d8161d501565b935060ff1681518110619b3257619b3261d394565b60200260200101819052505b60008160ff1667ffffffffffffffff811115619b5c57619b5c61c896565b604051908082528060200260200182016040528015619b8f57816020015b6060815260200190600190039081619b7a5790505b50905060005b8260ff168160ff161015619be857838160ff1681518110619bb857619bb861d394565b6020026020010151828260ff1681518110619bd557619bd561d394565b6020908102919091010152600101619b95565b5093505050505b949350505050565b619c1e6040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c91619ca49186910161d58b565b600060405180830381865afa158015619cc1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619ce9919081019061d02e565b90506000619cf7868361b437565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b8152600401619d27919061c753565b6000604051808303816000875af1158015619d46573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619d6e919081019061d5d2565b805190915060030b15801590619d875750602081015151155b8015619d965750604081015151155b15618d705781600081518110619dae57619dae61d394565b6020026020010151604051602001617fef919061d688565b60606000619dfb8560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528651825280870190820152909150619e329082905b9061b58c565b15619f8f576000619eaf82619ea984619ea3619e758a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061b5b3565b9061b615565b604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150619f1390829061b58c565b15619f7d57604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619f7a905b829061b69a565b90505b619f868161b6c0565b92505050617e03565b8215619fa8578484604051602001617fef92919061d874565b5050604080516020810190915260008152617e03565b509392505050565b6000808251602084016000f09392505050565b6185eb8282600161a0dc565b60408051600481526024810182526020810180516001600160e01b03167fad3cb1cc00000000000000000000000000000000000000000000000000000000179052905160609160009182916001600160a01b0386169161a045919061d91b565b600060405180830381855afa9150503d806000811461a080576040519150601f19603f3d011682016040523d82523d6000602084013e61a085565b606091505b509150915081801561a098575060208151115b1561a0b15780806020019051810190619bef919061d02e565b505060408051602081019091526000815292915050565b600061a0d4838361b729565b159392505050565b8160a001511561a0eb57505050565b600061a0f884848461b804565b9050600061a10582619bf7565b602081015181519192509060030b15801561a1a15750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a1a190604080518082018252600080825260209182015281518083019092528451825280850190820152619e2c565b1561a1ae57505050505050565b6040820151511561a1ce578160400151604051602001617fef919061d937565b80604051602001617fef919061d995565b6060600061a2148360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061a279905b829061a847565b1561a2e857604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617e039061a2e390839061bd9f565b61b6c0565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a34a905b829061be29565b60010361a41757604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a3b090619f73565b50604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617e039061a2e3905b839061b69a565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a4769061a272565b1561a5ad57604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061a4de90839061bec3565b90506000816001835161a4f1919061cac4565b8151811061a5015761a50161d394565b6020026020010151905061a5a461a2e361a5776040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925285518252808601908201529061bd9f565b95945050505050565b82604051602001617fef919061da00565b50919050565b6060600061a5f98360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061a65b9061a272565b1561a66957617e038161b6c0565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a6c89061a343565b60010361a73257604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617e039061a2e39061a410565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a7919061a272565b1561a5ad57604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061a7f990839061bec3565b905060018151111561a83557806002825161a814919061cac4565b8151811061a8245761a82461d394565b602002602001015192505050919050565b5082604051602001617fef919061da00565b80518251600091111561a85c57506000617b6d565b8151835160208501516000929161a8729161ccfd565b61a87c919061cac4565b90508260200151810361a893576001915050617b6d565b82516020840151819020912014905092915050565b6060600061a8b58361bf68565b600101905060008167ffffffffffffffff81111561a8d55761a8d561c896565b6040519080825280601f01601f19166020018201604052801561a8ff576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461a90957509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e534544000000000000000000000000000000000000000000008184019081528551808701875283815284019290925284518086019095525184529083015260609161a9d4905b829061a0c8565b1561aa1457505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261aa739061a9cd565b1561aab357505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ab129061a9cd565b1561ab5257505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261abb19061a9cd565b8061ac165750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ac169061a9cd565b1561ac5657505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261acb59061a9cd565b8061ad1a5750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ad1a9061a9cd565b1561ad5a57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261adb99061a9cd565b8061ae1e5750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ae1e9061a9cd565b1561ae5e57505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261aebd9061a9cd565b8061af225750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261af229061a9cd565b1561af6257505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261afc19061a9cd565b1561b00157505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b0609061a9cd565b1561b0a057505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b0ff9061a9cd565b1561b13f57505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b19e9061a9cd565b1561b1de57505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b23d9061a9cd565b1561b27d57505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b2dc9061a9cd565b8061b3415750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b3419061a9cd565b1561b38157505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b3e09061a9cd565b1561b42057505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151617fef929060200161dade565b60608060005b845181101561b4c2578185828151811061b4595761b45961d394565b602002602001015160405160200161b47292919061cedc565b60405160208183030381529060405291506001855161b491919061cac4565b811461b4ba578160405160200161b4a8919061dc47565b60405160208183030381529060405291505b60010161b43d565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161b4db579050509050838160008151811061b5065761b50661d394565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061b55a5761b55a61d394565b6020026020010181905250818160028151811061b5795761b57961d394565b6020908102919091010152949350505050565b602080830151835183519284015160009361b5aa929184919061c04a565b14159392505050565b6040805180820190915260008082526020820152600061b5e5846000015185602001518560000151866020015161c15b565b905083602001518161b5f7919061cac4565b8451859061b60690839061cac4565b90525060208401525090919050565b604080518082019091526000808252602082015281518351101561b63a575081617b6d565b602080830151908401516001911461b6615750815160208481015190840151829020919020145b801561b6925782518451859061b67890839061cac4565b905250825160208501805161b68e90839061ccfd565b9052505b509192915050565b604080518082019091526000808252602082015261b6b983838361c27b565b5092915050565b60606000826000015167ffffffffffffffff81111561b6e15761b6e161c896565b6040519080825280601f01601f19166020018201604052801561b70b576020820181803683370190505b509050600060208201905061b6b9818560200151866000015161c326565b815181516000919081111561b73c575081515b6020808501519084015160005b8381101561b7f5578251825180821461b7c557600019602087101561b7a45760018461b77689602061cac4565b61b780919061ccfd565b61b78b90600861dc88565b61b79690600261dd86565b61b7a0919061cac4565b1990505b818116838216818103911461b7c2579750617b6d9650505050505050565b50505b61b7d060208661ccfd565b945061b7dd60208561ccfd565b9350505060208161b7ee919061ccfd565b905061b749565b5084518651618d70919061dd92565b6060600061b8106185ef565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161b82d57905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061b8889061d501565b935060ff168151811061b89d5761b89d61d394565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161b8ee919061ddb2565b60405160208183030381529060405282828061b9099061d501565b935060ff168151811061b91e5761b91e61d394565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061b96b9061d501565b935060ff168151811061b9805761b98061d394565b60200260200101819052508260405160200161b99c919061d42f565b60405160208183030381529060405282828061b9b79061d501565b935060ff168151811061b9cc5761b9cc61d394565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061ba199061d501565b935060ff168151811061ba2e5761ba2e61d394565b602002602001018190525061ba43878461c3a0565b828261ba4e8161d501565b935060ff168151811061ba635761ba6361d394565b60209081029190910101528551511561bb0f5760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261bab58161d501565b935060ff168151811061baca5761baca61d394565b602002602001018190525061bae386600001518461c3a0565b828261baee8161d501565b935060ff168151811061bb035761bb0361d394565b60200260200101819052505b85608001511561bb7d5760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261bb588161d501565b935060ff168151811061bb6d5761bb6d61d394565b602002602001018190525061bbe3565b841561bbe35760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261bbc28161d501565b935060ff168151811061bbd75761bbd761d394565b60200260200101819052505b6040860151511561bc7f5760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261bc2d8161d501565b935060ff168151811061bc425761bc4261d394565b6020026020010181905250856040015182828061bc5e9061d501565b935060ff168151811061bc735761bc7361d394565b60200260200101819052505b85606001511561bce95760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261bcc88161d501565b935060ff168151811061bcdd5761bcdd61d394565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561bd075761bd0761c896565b60405190808252806020026020018201604052801561bd3a57816020015b606081526020019060019003908161bd255790505b50905060005b8260ff168160ff16101561bd9357838160ff168151811061bd635761bd6361d394565b6020026020010151828260ff168151811061bd805761bd8061d394565b602090810291909101015260010161bd40565b50979650505050505050565b604080518082019091526000808252602082015281518351101561bdc4575081617b6d565b8151835160208501516000929161bdda9161ccfd565b61bde4919061cac4565b6020840151909150600190821461be05575082516020840151819020908220145b801561be205783518551869061be1c90839061cac4565b9052505b50929392505050565b600080826000015161be4d856000015186602001518660000151876020015161c15b565b61be57919061ccfd565b90505b8351602085015161be6b919061ccfd565b811161b6b9578161be7b8161ddf7565b925050826000015161beb285602001518361be96919061cac4565b865161bea2919061cac4565b838660000151876020015161c15b565b61bebc919061ccfd565b905061be5a565b6060600061bed1848461be29565b61bedc90600161ccfd565b67ffffffffffffffff81111561bef45761bef461c896565b60405190808252806020026020018201604052801561bf2757816020015b606081526020019060019003908161bf125790505b50905060005b8151811015619fbe5761bf4361a2e3868661b69a565b82828151811061bf555761bf5561d394565b602090810291909101015260010161bf2d565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061bfb1577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061bfdd576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061bffb57662386f26fc10000830492506010015b6305f5e100831061c013576305f5e100830492506008015b612710831061c02757612710830492506004015b6064831061c039576064830492506002015b600a8310617b6d5760010192915050565b60008085841161c151576020841161c0fd576000841561c09557600161c07186602061cac4565b61c07c90600861dc88565b61c08790600261dd86565b61c091919061cac4565b1990505b835181168561c0a4898961ccfd565b61c0ae919061cac4565b805190935082165b81811461c0e85787841161c0d05787945050505050619bef565b8361c0da8161de11565b94505082845116905061c0b6565b61c0f2878561ccfd565b945050505050619bef565b83832061c10a858861cac4565b61c114908761ccfd565b91505b85821061c14f5784822080820361c13c5761c132868461ccfd565b9350505050619bef565b61c14760018461cac4565b92505061c117565b505b5092949350505050565b6000838186851161c266576020851161c215576000851561c1a757600161c18387602061cac4565b61c18e90600861dc88565b61c19990600261dd86565b61c1a3919061cac4565b1990505b8451811660008761c1b88b8b61ccfd565b61c1c2919061cac4565b855190915083165b82811461c2075781861061c1ef5761c1e28b8b61ccfd565b9650505050505050619bef565b8561c1f98161ddf7565b96505083865116905061c1ca565b859650505050505050619bef565b508383206000905b61c227868961cac4565b821161c2645785832080820361c2435783945050505050619bef565b61c24e60018561ccfd565b935050818061c25c9061ddf7565b92505061c21d565b505b61c270878761ccfd565b979650505050505050565b6040805180820190915260008082526020820152600061c2ad856000015186602001518660000151876020015161c15b565b60208087018051918601919091525190915061c2c9908261cac4565b83528451602086015161c2dc919061ccfd565b810361c2eb576000855261c31d565b8351835161c2f9919061ccfd565b8551869061c30890839061cac4565b905250835161c317908261ccfd565b60208601525b50909392505050565b6020811061c35e578151835261c33d60208461ccfd565b925061c34a60208361ccfd565b915061c35760208261cac4565b905061c326565b600019811561c38d57600161c37483602061cac4565b61c3809061010061dd86565b61c38a919061cac4565b90505b9151835183169219169190911790915250565b6060600061c3ae84846186c2565b805160208083015160405193945061c3c89390910161de28565b60405160208183030381529060405291505092915050565b610c9f8061de8183390190565b610efa8061eb2083390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161c43d61c442565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161c43d6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561c4f45783516001600160a01b031683526020938401939092019160010161c4cd565b509095945050505050565b60005b8381101561c51a57818101518382015260200161c502565b50506000910152565b6000815180845261c53b81602086016020860161c4ff565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561c64b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561c631577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261c61b84865161c523565b602095860195909450929092019160010161c5e1565b50919750505060209485019492909201915060010161c577565b50929695505050505050565b600081518084526020840193506020830160005b8281101561c6ab5781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161c66b565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561c64b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261c721604088018261c523565b905060208201519150868103602088015261c73c818361c657565b96505050602093840193919091019060010161c6dd565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561c64b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261c7b585835161c523565b9450602093840193919091019060010161c77b565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561c64b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261c84b604087018261c657565b955050602093840193919091019060010161c7f2565b602081526000617e03602083018461c523565b60006020828403121561c88657600080fd5b81518015158114617e0357600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c9082168061c8d957607f821691505b60208210810361a5be577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f82111561c95957806000526020600020601f840160051c8101602085101561c9395750805b601f840160051c820191505b81811015612639576000815560010161c945565b505050565b815167ffffffffffffffff81111561c9785761c97861c896565b61c98c8161c986845461c8c5565b8461c912565b6020601f82116001811461c9c0576000831561c9a85750848201515b600019600385901b1c1916600184901b178455612639565b600084815260208120601f198516915b8281101561c9f0578785015182556020948501946001909201910161c9d0565b508482101561ca0e5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6001600160a01b03851681526001600160a01b0384166020820152826040820152608060608201526000618d70608083018461c523565b60006020828403121561ca6657600080fd5b5051919050565b6001600160a01b038416815282602082015260606040820152600061a5a4606083018461c523565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810381811115617b6d57617b6d61ca95565b60808152600061caea608083018761c523565b6001600160a01b0386166020840152846040840152828103606084015261c270818561c523565b60006020828403121561cb2357600080fd5b81516001600160a01b0381168114617e0357600080fd5b6001600160a01b0381541682526001600160a01b03600182015416602083015260028101546040830152600060038201608060608501526000815461cb7e8161c8c5565b806080880152600182166000811461cb9d576001811461cbd75761cc0b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660a089015260a082151560051b890101935061cc0b565b84600052602060002060005b8381101561cc025781548a820160a0015260019091019060200161cbe3565b890160a0019450505b50919695505050505050565b6001600160a01b0383168152604060208201526000619bef604083018461cb3a565b83815260606020820152600061cc52606083018561c523565b8281036040840152618d70818561cb3a565b6001600160a01b03861681526001600160a01b038516602082015283604082015260a06060820152600061cc9b60a083018561c523565b828103608084015261ccad818561cb3a565b98975050505050505050565b828152604060208201526000619bef604083018461c523565b60608152600061cce5606083018661c523565b8460208401528281036040840152618d70818561c523565b80820180821115617b6d57617b6d61ca95565b60008261cd46577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6001600160a01b0383168152604060208201526000619bef604083018461c523565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161cda581601a85016020880161c4ff565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161cde281601c84016020880161c4ff565b01601c01949350505050565b6040516060810167ffffffffffffffff8111828210171561ce115761ce1161c896565b60405290565b60008067ffffffffffffffff84111561ce325761ce3261c896565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561ce615761ce6161c896565b60405283815290508082840185101561ce7957600080fd5b619fbe84602083018561c4ff565b600082601f83011261ce9857600080fd5b617e038383516020850161ce17565b60006020828403121561ceb957600080fd5b815167ffffffffffffffff81111561ced057600080fd5b617b698482850161ce87565b6000835161ceee81846020880161c4ff565b83519083019061cf0281836020880161c4ff565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161cf4381601a85016020880161c4ff565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161cf8081603384016020880161c4ff565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b6001600160a01b03841681526001600160a01b038316602082015260606040820152600061a5a4606083018461c523565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000617e03608083018461c523565b60006020828403121561d04057600080fd5b815167ffffffffffffffff81111561d05757600080fd5b8201601f8101841361d06857600080fd5b617b698482516020840161ce17565b6000855161d089818460208a0161c4ff565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161d0c3816001840160208a0161c4ff565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161d10181600284016020890161c4ff565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161d14381600284016020880161c4ff565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061d18e604083018461c523565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161d20581601f85016020870161c4ff565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061d272604083018461c523565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061d2c4604083018461c523565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161d33b81601485016020870161c4ff565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061d382604083018561c523565b8281036020840152617dff818561c523565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161d3fb81600185016020870161c4ff565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161d44181846020870161c4ff565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161d4f481604b85016020870161c4ff565b91909101604b0192915050565b600060ff821660ff810361d5175761d51761ca95565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161d57e81602985016020870161c4ff565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000617e03608083018461c523565b60006020828403121561d5e457600080fd5b815167ffffffffffffffff81111561d5fb57600080fd5b82016060818503121561d60d57600080fd5b61d61561cdee565b81518060030b811461d62657600080fd5b8152602082015167ffffffffffffffff81111561d64257600080fd5b61d64e8682850161ce87565b602083015250604082015167ffffffffffffffff81111561d66e57600080fd5b61d67a8682850161ce87565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161d6e681602185016020870161c4ff565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161d8d281602185016020880161c4ff565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161d90f81602e84016020880161c4ff565b01602e01949350505050565b6000825161d92d81846020870161c4ff565b9190910192915050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161d57e81602985016020870161c4ff565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161d9f381602285016020870161c4ff565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161da3881600e85016020870161c4ff565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161db1681601885016020880161c4ff565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161db5381601c84016020880161c4ff565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161dc5981846020870161c4ff565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b8082028115828204841417617b6d57617b6d61ca95565b6001815b600184111561dcda5780850481111561dcbe5761dcbe61ca95565b600184161561dccc57908102905b60019390931c92800261dca3565b935093915050565b60008261dcf157506001617b6d565b8161dcfe57506000617b6d565b816001811461dd14576002811461dd1e5761dd3a565b6001915050617b6d565b60ff84111561dd2f5761dd2f61ca95565b50506001821b617b6d565b5060208310610133831016604e8410600b841016171561dd5d575081810a617b6d565b61dd6a600019848461dc9f565b806000190482111561dd7e5761dd7e61ca95565b029392505050565b6000617e03838361dce2565b818103600083128015838313168383128216171561b6b95761b6b961ca95565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161ddea81601c85016020870161c4ff565b91909101601c0192915050565b6000600019820361de0a5761de0a61ca95565b5060010190565b60008161de205761de2061ca95565b506000190190565b6000835161de3a81846020880161c4ff565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161de7481600184016020880161c4ff565b0160010194935050505056fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220a043c41353215fce25ecb67a8a4f6f724aaa86dea8dcb0a6975bfb1f10ff3beb64736f6c634300081a00336080604052348015600f57600080fd5b506001600055610ed6806100246000396000f3fe60806040526004361061006e5760003560e01c8063c51316911161004b578063c5131691146100d5578063c9028a36146100f5578063e04d4f9714610115578063f05b6abf1461012857005b8063357fc5a214610077578063676cc054146100975780636ed70169146100c057005b3661007557005b005b34801561008357600080fd5b50610075610092366004610724565b610148565b6100aa6100a5366004610760565b6101de565b6040516100b7919061085b565b60405180910390f35b3480156100cc57600080fd5b50610075610211565b3480156100e157600080fd5b506100756100f0366004610724565b610246565b34801561010157600080fd5b5061007561011036600461086e565b610321565b6100756101233660046109ce565b61035d565b34801561013457600080fd5b50610075610143366004610aba565b6103a1565b6101506103d6565b61017273ffffffffffffffffffffffffffffffffffffffff8316338386610419565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101d96001600055565b505050565b6040516060907f3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee90600090a19392505050565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61024e6103d6565b600061025b600285610ba4565b905080600003610297576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102b973ffffffffffffffffffffffffffffffffffffffff8416338484610419565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101d96001600055565b7f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b3382604051610352929190610c28565b60405180910390a150565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa3334858585604051610394959493929190610d1a565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146338484846040516103949493929190610da4565b600260005403610412576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526104ae9085906104b4565b50505050565b60006104d673ffffffffffffffffffffffffffffffffffffffff84168361054f565b905080516000141580156104fb5750808060200190518101906104f99190610e67565b155b156101d9576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b606061055d83836000610564565b9392505050565b6060814710156105a2576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610546565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105cb9190610e84565b60006040518083038185875af1925050503d8060008114610608576040519150601f19603f3d011682016040523d82523d6000602084013e61060d565b606091505b509150915061061d868383610627565b9695505050505050565b60608261063c57610637826106b6565b61055d565b8151158015610660575073ffffffffffffffffffffffffffffffffffffffff84163b155b156106af576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610546565b508061055d565b8051156106c65780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff8116811461071f57600080fd5b919050565b60008060006060848603121561073957600080fd5b83359250610749602085016106fb565b9150610757604085016106fb565b90509250925092565b6000806000838503604081121561077657600080fd5b602081121561078457600080fd5b50839250602084013567ffffffffffffffff8111156107a257600080fd5b8401601f810186136107b357600080fd5b803567ffffffffffffffff8111156107ca57600080fd5b8660208284010111156107dc57600080fd5b939660209190910195509293505050565b60005b838110156108085781810151838201526020016107f0565b50506000910152565b600081518084526108298160208601602086016107ed565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061055d6020830184610811565b60006020828403121561088057600080fd5b813567ffffffffffffffff81111561089757600080fd5b82016080818503121561055d57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561091f5761091f6108a9565b604052919050565b600082601f83011261093857600080fd5b813567ffffffffffffffff811115610952576109526108a9565b61098360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016108d8565b81815284602083860101111561099857600080fd5b816020850160208301376000918101602001919091529392505050565b80151581146106f857600080fd5b803561071f816109b5565b6000806000606084860312156109e357600080fd5b833567ffffffffffffffff8111156109fa57600080fd5b610a0686828701610927565b935050602084013591506040840135610a1e816109b5565b809150509250925092565b600067ffffffffffffffff821115610a4357610a436108a9565b5060051b60200190565b600082601f830112610a5e57600080fd5b8135610a71610a6c82610a29565b6108d8565b8082825260208201915060208360051b860101925085831115610a9357600080fd5b602085015b83811015610ab0578035835260209283019201610a98565b5095945050505050565b600080600060608486031215610acf57600080fd5b833567ffffffffffffffff811115610ae657600080fd5b8401601f81018613610af757600080fd5b8035610b05610a6c82610a29565b8082825260208201915060208360051b850101925088831115610b2757600080fd5b602084015b83811015610b6957803567ffffffffffffffff811115610b4b57600080fd5b610b5a8b602083890101610927565b84525060209283019201610b2c565b509550505050602084013567ffffffffffffffff811115610b8957600080fd5b610b9586828701610a4d565b925050610757604085016109c3565b600082610bda577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610c66836106fb565b16604082015273ffffffffffffffffffffffffffffffffffffffff610c8d602084016106fb565b166060820152600080604084013590508060808401525060608301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610cd957600080fd5b830160208101903567ffffffffffffffff811115610cf657600080fd5b803603821315610d0557600080fd5b608060a085015261061d60c085018284610bdf565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610d4f60a0830186610811565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610d9a578151865260209586019590910190600101610d7c565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610e37577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610e22858351610811565b94506020938401939190910190600101610de8565b505050508281036040840152610e4d8186610d68565b915050610e5e606083018415159052565b95945050505050565b600060208284031215610e7957600080fd5b815161055d816109b5565b60008251610e968184602087016107ed565b919091019291505056fea26469706673582212202ffe1386f11164490b308722577353cf98c797014ad1ba3830d8d3f4c65fc53a64736f6c634300081a0033a264697066735822122051c5163ad68d52723c0a62be5ba5ee64124201ee34afd46810a467d604d5a7b464736f6c634300081a0033",
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

// TestUpgradeAndWithdraw is a paid mutator transaction binding the contract method 0xaf298bb1.
//
// Solidity: function testUpgradeAndWithdraw() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactor) TestUpgradeAndWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyTest.contract.Transact(opts, "testUpgradeAndWithdraw")
}

// TestUpgradeAndWithdraw is a paid mutator transaction binding the contract method 0xaf298bb1.
//
// Solidity: function testUpgradeAndWithdraw() returns()
func (_ERC20CustodyTest *ERC20CustodyTestSession) TestUpgradeAndWithdraw() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestUpgradeAndWithdraw(&_ERC20CustodyTest.TransactOpts)
}

// TestUpgradeAndWithdraw is a paid mutator transaction binding the contract method 0xaf298bb1.
//
// Solidity: function testUpgradeAndWithdraw() returns()
func (_ERC20CustodyTest *ERC20CustodyTestTransactorSession) TestUpgradeAndWithdraw() (*types.Transaction, error) {
	return _ERC20CustodyTest.Contract.TestUpgradeAndWithdraw(&_ERC20CustodyTest.TransactOpts)
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
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCustodyTSSAddress is a free log retrieval operation binding the contract event 0x086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba.
//
// Solidity: event UpdatedCustodyTSSAddress(address newTSSAddress)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterUpdatedCustodyTSSAddress(opts *bind.FilterOpts) (*ERC20CustodyTestUpdatedCustodyTSSAddressIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestUpdatedCustodyTSSAddressIterator{contract: _ERC20CustodyTest.contract, event: "UpdatedCustodyTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedCustodyTSSAddress is a free log subscription operation binding the contract event 0x086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba.
//
// Solidity: event UpdatedCustodyTSSAddress(address newTSSAddress)
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

// ParseUpdatedCustodyTSSAddress is a log parse operation binding the contract event 0x086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba.
//
// Solidity: event UpdatedCustodyTSSAddress(address newTSSAddress)
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
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*ERC20CustodyTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestUpdatedGatewayTSSAddressIterator{contract: _ERC20CustodyTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
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

// ParseUpdatedGatewayTSSAddress is a log parse operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
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

// ERC20CustodyTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the ERC20CustodyTest contract.
type ERC20CustodyTestWithdrawnV2Iterator struct {
	Event *ERC20CustodyTestWithdrawnV2 // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyTestWithdrawnV2)
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
		it.Event = new(ERC20CustodyTestWithdrawnV2)
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
func (it *ERC20CustodyTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyTestWithdrawnV2 represents a WithdrawnV2 event raised by the ERC20CustodyTest contract.
type ERC20CustodyTestWithdrawnV2 struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnV2 is a free log retrieval operation binding the contract event 0xd4dabfe72081670cc78f2ebda8e2eddaf3feebde6288dcb8fe673b3dc201b5a4.
//
// Solidity: event WithdrawnV2(address indexed to, address indexed token, uint256 amount)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyTestWithdrawnV2Iterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.FilterLogs(opts, "WithdrawnV2", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTestWithdrawnV2Iterator{contract: _ERC20CustodyTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0xd4dabfe72081670cc78f2ebda8e2eddaf3feebde6288dcb8fe673b3dc201b5a4.
//
// Solidity: event WithdrawnV2(address indexed to, address indexed token, uint256 amount)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *ERC20CustodyTestWithdrawnV2, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyTest.contract.WatchLogs(opts, "WithdrawnV2", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyTestWithdrawnV2)
				if err := _ERC20CustodyTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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

// ParseWithdrawnV2 is a log parse operation binding the contract event 0xd4dabfe72081670cc78f2ebda8e2eddaf3feebde6288dcb8fe673b3dc201b5a4.
//
// Solidity: event WithdrawnV2(address indexed to, address indexed token, uint256 amount)
func (_ERC20CustodyTest *ERC20CustodyTestFilterer) ParseWithdrawnV2(log types.Log) (*ERC20CustodyTestWithdrawnV2, error) {
	event := new(ERC20CustodyTestWithdrawnV2)
	if err := _ERC20CustodyTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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
