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

// CallOptions is an auto generated low-level Go binding around an user-defined struct.
type CallOptions struct {
	GasLimit        *big.Int
	IsArbitraryCall bool
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

// GatewayZEVMInboundTestMetaData contains all meta data concerning the GatewayZEVMInboundTest contract.
var GatewayZEVMInboundTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testCallFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOpts\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfGasLimitIsBelowMin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpgradeAndWithdrawZRC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETA\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfNoBalance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithBothDefaultValuesFromRegistry\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsWithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithDefaultGasLimit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithDefaultProtocolFlatFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfNoBalanceForGasFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfNoBalanceForTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIsAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCallOptsWithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithMessageFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZETANotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroGasPrice\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5062019163806200003e6000396000f3fe608060405234801561001057600080fd5b50600436106103995760003560e01c80636dfcbc50116101e9578063c3bb79571161010f578063e51c6388116100ad578063ecf26b301161007c578063ecf26b30146105a2578063fa7626d4146105aa578063fbc611c8146105b7578063fdad0ad0146105bf57600080fd5b8063e51c63881461058a578063e804a406146103e0578063ea37902f14610592578063eb7a2fac1461059a57600080fd5b8063d597a27a116100e9578063d597a27a1461056a578063dc749dd714610572578063dde7e9671461057a578063e20c9f711461058257600080fd5b8063c3bb795714610552578063c946d7c01461055a578063ceccfab31461056257600080fd5b8063b0464fdc11610187578063ba800c9111610156578063ba800c9114610532578063ba9adeef1461053a578063bed3e81314610542578063c20049f41461054a57600080fd5b8063b0464fdc14610502578063b51ac0711461050a578063b5508aa914610512578063ba414fa61461051a57600080fd5b806385226c81116101c357806385226c81146104d0578063916a17c6146104e557806391dc32e7146104fa578063a721b2d31461046b57600080fd5b80636dfcbc50146104b85780637ae69730146104c057806383ababa9146104c857600080fd5b80632ade3880116102ce5780634ffab9de1161026c5780636221b5091161023b5780636221b5091461048b57806364002a1f1461049357806366d9a9a01461049b5780636d6ce0d0146104b057600080fd5b80634ffab9de1461046b5780635d72228f146104735780635efe72a91461047b5780636198fb191461048357600080fd5b80633e5e3c23116102a85780633e5e3c231461044b5780633f7286f41461045357806342752d411461045b578063431814371461046357600080fd5b80632ade38801461042657806336431b3f1461043b5780633828ce8c1461044357600080fd5b80631b9641bf1161033b57806320dee15f1161031557806320dee15f1461040657806321aeb18c1461040e578063264b524c14610416578063291f1d821461041e57600080fd5b80631b9641bf146103d85780631e63d2b9146103e05780631ed7831c146103e857600080fd5b80630b5ad28d116103775780630b5ad28d146103b85780631238212c146103c057806314759766146103c857806316c09ef7146103d057600080fd5b806301a74aee1461039e57806305b9f046146103a85780630a9254e4146103b0575b600080fd5b6103a66105c7565b005b6103a6610808565b6103a66109b2565b6103a661180b565b6103a6611973565b6103a6611d13565b6103a66123ef565b6103a661296b565b6103a6612acd565b6103f0612ec6565b6040516103fd919061d9d7565b60405180910390f35b6103a6612f28565b6103a6613316565b6103a6613421565b6103a66136f9565b61042e613d3c565b6040516103fd919061da73565b6103a6613e7e565b6103a6613fbf565b6103f0614490565b6103f06144f0565b6103a6614550565b6103a66146a7565b6103a6614794565b6103a6614ba3565b6103a6614edc565b6103a6614f7f565b6103a6615130565b6103a6615355565b6104a361549b565b6040516103fd919061dbd9565b6103a6615608565b6103a66158df565b6103a6615cc9565b6103a6615f95565b6104d86160a7565b6040516103fd919061dc77565b6104ed616177565b6040516103fd919061dcee565b6103a661625d565b6104ed61647f565b6103a6616565565b6104d861664f565b61052261671f565b60405190151581526020016103fd565b6103a66167f3565b6103a6616909565b6103a6616a7c565b6103a6616be0565b6103a6616d62565b6103a6616f1f565b6103a66172d2565b6103a66176e6565b6103a661787e565b6103a6617924565b6103f0617d4a565b6103a6617daa565b6103a661821d565b6103a6618892565b6103a661897c565b601f546105229060ff1681565b6103a6618b88565b6103a6618f15565b6027546040516001600160a01b03909116602482015260009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a0810183526103218082526001828401819052828501919091528351928301909352600080835260608201929092529293509190608082019061065890621e84809061ddb4565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39161070b9160040161ddc7565b600060405180830381600087803b15801561072557600080fd5b505af1158015610739573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506306cb8983915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526107d2926001600160a01b03909116908790602d90889060040161de35565b600060405180830381600087803b1580156107ec57600080fd5b505af1158015610800573d6000803e3d6000fd5b505050505050565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b179052905061085e6001620186a061de9f565b602d55604051630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b600060405180830381600087803b1580156108ce57600080fd5b505af11580156108e2573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506306cb898391506034015b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261097d926001600160a01b03909116908690602d9060299060040161dfee565b600060405180830381600087803b15801561099757600080fd5b505af11580156109ab573d6000803e3d6000fd5b5050505050565b602680547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602780549091166112341790556040516109f89061d8d8565b604051809103906000f080158015610a14573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600f81527f476174657761795a45564d2e736f6c000000000000000000000000000000000060208201526026549151602481019390935292166044820152610ae3919060640160408051601f198184030181529190526020810180516001600160e01b03167f485cc955000000000000000000000000000000000000000000000000000000001790526191e1565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b039384168102919091179182905560208054919092049092167fffffffffffffffffffffffff000000000000000000000000000000000000000090921682178155604080517f06433b1b0000000000000000000000000000000000000000000000000000000081529051600093926306433b1b92600480820193918290030181865afa158015610ba7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bcb919061e068565b9050600060405180602001610bdf9061d8e6565b601f1982820381018352601f909101166040819052610c01919060200161e083565b60405160208183030381529060405290506000808251602084016000f590507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663b4d6c78284836001600160a01b0316803b806020016040519081016040528181526000908060200190933c6040518363ffffffff1660e01b8152600401610c9692919061e09f565b600060405180830381600087803b158015610cb057600080fd5b505af1158015610cc4573d6000803e3d6000fd5b50506026546020546040517fc0c53b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820181905260248201529082166044820152908616925063c0c53b8b9150606401600060405180830381600087803b158015610d3a57600080fd5b505af1158015610d4e573d6000803e3d6000fd5b5050602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038781169190911790915560208054604080517f2722feee0000000000000000000000000000000000000000000000000000000081529051919093169450632722feee93506004808401938290030181865afa158015610de1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e05919061e068565b602880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055604051610e499061d8f4565b604051809103906000f080158015610e65573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161781556028546040517f06447d5600000000000000000000000000000000000000000000000000000000815292166004830152737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d569101600060405180830381600087803b158015610f0157600080fd5b505af1158015610f15573d6000803e3d6000fd5b505050506000806000604051610f2a9061d902565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610f66573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602054604051601293600193849360009391921690610fbc9061d910565b610fcb9695949392919061e0c1565b604051809103906000f080158015610fe7573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179091556023546040517fee2815ba0000000000000000000000000000000000000000000000000000000081526001600482015260248101929092529091169063ee2815ba90604401600060405180830381600087803b15801561107e57600080fd5b505af1158015611092573d6000803e3d6000fd5b50506023546040517fa7cb050700000000000000000000000000000000000000000000000000000000815260016004820181905260248201526001600160a01b03909116925063a7cb05079150604401600060405180830381600087803b1580156110fc57600080fd5b505af1158015611110573d6000803e3d6000fd5b50506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152633b9aca006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b15801561119057600080fd5b505af11580156111a4573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b1580156111f957600080fd5b505af115801561120d573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af1158015611281573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112a5919061e1b6565b506021546026546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116906347e7ef24906044016020604051808303816000875af1158015611317573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061133b919061e1b6565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561139a57600080fd5b505af11580156113ae573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561142457600080fd5b505af1158015611438573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116925063095ea7b391506044016020604051808303816000875af11580156114ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114d1919061e1b6565b506021546025546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e100602482015291169063095ea7b3906044016020604051808303816000875af1158015611543573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611567919061e1b6565b50602260009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b1580156115b957600080fd5b505af11580156115cd573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af1158015611641573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611665919061e1b6565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156116c457600080fd5b505af11580156116d8573d6000803e3d6000fd5b50506040805160a08101825261032180825260016020808401918252838501928352845190810190945260008085526060840185905260808401528251602980549251151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009093166001600160a01b03928316179290921782559151602a8054919093167fffffffffffffffffffffffff000000000000000000000000000000000000000091909116179091559093509150602b906117af908261e24e565b5060809190910151600390910155505060408051808201909152620186a080825260016020909201829052602d55602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909117905550565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b1790526000602d5551630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156118c357600080fd5b505af11580156118d7573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b91506034015b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261097d926001916001600160a01b0316908790602d9060299060040161e30d565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa1580156119c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119e8919061e37f565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611a4457600080fd5b505af1158015611a58573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015611aca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aee919061e1b6565b506027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905280517ff48448140000000000000000000000000000000000000000000000000000000081529051919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f48448149160048082019260009290919082900301818387803b158015611b9c57600080fd5b505af1158015611bb0573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a084526000602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b168152611c5c93919289926001600160a01b039091169188919060299060040161e398565b600060405180830381600087803b158015611c7657600080fd5b505af1158015611c8a573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015611cdd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d01919061e37f565b9050611d0d8382619200565b50505050565b6027546040516001600160a01b039091166024820152600190819060009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b17905290516d98760000000000000000000000009181019190915290915060009060340160408051808303601f190181528282018252600883527f6761734c696d69740000000000000000000000000000000000000000000000006020840152602654915163ca669fa760e01b81526001600160a01b0390921660048301529250737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611e1657600080fd5b505af1158015611e2a573d6000803e3d6000fd5b50506025546021546040517fa8f2cb960000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063a8f2cb969350611e819289921690879060019060040161e3ec565b600060405180830381600087803b158015611e9b57600080fd5b505af1158015611eaf573d6000803e3d6000fd5b505050506040518060400160405280600f81526020017f70726f746f636f6c466c61744665650000000000000000000000000000000000815250905060006003604051602001611f08919060ff91909116815260200190565b6040516020818303038152906040529050600081806020019051810190611f2f919061e37f565b6023546040517fd7fd7afb000000000000000000000000000000000000000000000000000000008152600481018990526001600160a01b039091169063d7fd7afb90602401602060405180830381865afa158015611f91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fb5919061e37f565b602d54611fc2919061e427565b611fcc919061ddb4565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561202857600080fd5b505af115801561203c573d6000803e3d6000fd5b50506025546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632259e9e5915061208d9089908790879060040161e43e565b600060405180830381600087803b1580156120a757600080fd5b505af11580156120bb573d6000803e3d6000fd5b50506026546021546040516370a0823160e01b81526001600160a01b0392831660048201819052319450600093509116906370a0823190602401602060405180830381865afa158015612112573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612136919061e37f565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156121c757600080fd5b505af11580156121db573d6000803e3d6000fd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201528b93506001600160a01b0390911691507fd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a90603401604051602081830303815290604052602260009054906101000a90046001600160a01b03168d888a806020019051810190612270919061e37f565b8e602d602960405161228998979695949392919061e469565b60405180910390a3602080546027546040516001600160a01b039283169363c34da122938e936122d4939116910160609190911b6bffffffffffffffffffffffff1916815260140190565b6040516020818303038152906040528b8b602d60296040518763ffffffff1660e01b815260040161230995949392919061e4dd565b6000604051808303818588803b15801561232257600080fd5b505af1158015612336573d6000803e3d6000fd5b50506026546001600160a01b031631925061235e915061235890508b8561de9f565b82619200565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156123af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123d3919061e37f565b90506123e2612358868561de9f565b5050505050505050505050565b6040516d98760000000000000000000000006020820152600190819060009060340160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561247857600080fd5b505af115801561248c573d6000803e3d6000fd5b50506025546021546040517fa8f2cb960000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063a8f2cb9693506124e39287921690869060019060040161e3ec565b600060405180830381600087803b1580156124fd57600080fd5b505af1158015612511573d6000803e3d6000fd5b5050604080518082018252600f81527f70726f746f636f6c466c617446656500000000000000000000000000000000006020808301919091528251600591810191909152909350600092500160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156125c457600080fd5b505af11580156125d8573d6000803e3d6000fd5b50506025546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632259e9e591506126299087908690869060040161e43e565b600060405180830381600087803b15801561264357600080fd5b505af1158015612657573d6000803e3d6000fd5b505050506000620186a09050600082806020019051810190612679919061e37f565b6023546040517fd7fd7afb000000000000000000000000000000000000000000000000000000008152600481018990526001600160a01b039091169063d7fd7afb90602401602060405180830381865afa1580156126db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126ff919061e37f565b612709908461e427565b612713919061ddb4565b6026546021546040516370a0823160e01b81526001600160a01b0392831660048201819052939450923192600092909116906370a0823190602401602060405180830381865afa15801561276b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061278f919061e37f565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561282057600080fd5b505af1158015612834573d6000803e3d6000fd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201528b93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c90603401604051602081830303815290604052602260009054906101000a90046001600160a01b03168d888b8060200190518101906128c9919061e37f565b6040805180820182528d81526001602082015290516128f09695949392919060299061e508565b60405180910390a3602080546027546040516001600160a01b03928316936322f0ce89938e9361293b939116910160609190911b6bffffffffffffffffffffffff1916815260140190565b6040516020818303038152906040528b60296040518563ffffffff1660e01b81526004016123099392919061e58a565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612a1e57600080fd5b505af1158015612a32573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261097d926000916001600160a01b0316908790602d9060299060040161e30d565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015612b1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b42919061e37f565b6027546040516001600160a01b03909116602482015290915060009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150620186a0908190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015612c2157600080fd5b505af1158015612c35573d6000803e3d6000fd5b505060265460275460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507fd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a9060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b03909216918b9189918491634d8943bb916004808201926020929091908290030181865afa158015612d07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d2b919061e37f565b6040805180820182528a8152600160208201529051612d539695949392918d9160299061e5b5565b60405180910390a3602080546027546040516001600160a01b0392831693637b15118b93612d9b9316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815260215483830183528684526001602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b168152612e079391928b926001600160a01b03909116918a919060299060040161e398565b600060405180830381600087803b158015612e2157600080fd5b505af1158015612e35573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015612e88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612eac919061e37f565b905061080083612ebc888861de9f565b612358919061de9f565b60606016805480602002602001604051908101604052809291908181526020018280548015612f1e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612f00575b5050505050905090565b601f54604080518082018252601a81527f476174657761795a45564d55706772616465546573742e736f6c000000000000602080830191909152825190810190925260008252602654612f8d936001600160a01b03610100909104811693911661927f565b601f546021546026546040516370a0823160e01b81526001600160a01b03918216600482015261010090930481169260019260009216906370a0823190602401602060405180830381865afa158015612fea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061300e919061e37f565b6040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b0385166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561309b57600080fd5b505af11580156130af573d6000803e3d6000fd5b505060265460275460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507f5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d979060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b039092169188916000918491634d8943bb916004808201926020929091908290030181865afa158015613182573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131a6919061e37f565b604080518082018252600081526001602082015290516131ce9695949392919060299061e508565b60405180910390a360275460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b03841690637c0dcb5f9060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526132629287916001600160a01b03169060299060040161e617565b600060405180830381600087803b15801561327c57600080fd5b505af1158015613290573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156132e3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613307919061e37f565b9050611d0d612358848461de9f565b604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561338257600080fd5b505af1158015613396573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506322f0ce8991506000906034015b604051602081830303815290604052600160296040518563ffffffff1660e01b81526004016134089392919061e58a565b6000604051808303818588803b15801561099757600080fd5b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169263bcbe9365926004808401938290030181865afa15801561347f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134a3919061e37f565b6134ae90600161ddb4565b67ffffffffffffffff8111156134c6576134c661e1d8565b6040519080825280601f01601f1916602001820160405280156134f0576020820181803683370190505b50602b906134fe908261e24e565b506000602960020180546135119061deb2565b905090506000602060009054906101000a90046001600160a01b03166001600160a01b031663bcbe93656040518163ffffffff1660e01b8152600401602060405180830381865afa15801561356a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061358e919061e37f565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916136359160040161ddc7565b600060405180830381600087803b15801561364f57600080fd5b505af1158015613663573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526107d2926002916001600160a01b03169060299060040161e617565b6040516d98760000000000000000000000006020820152600190819060009060340160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561378257600080fd5b505af1158015613796573d6000803e3d6000fd5b50506025546021546040517fa8f2cb960000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063a8f2cb9693506137ed9287921690869060019060040161e3ec565b600060405180830381600087803b15801561380757600080fd5b505af115801561381b573d6000803e3d6000fd5b5050604080518082018252600881527f6761734c696d69740000000000000000000000000000000000000000000000006020808301919091528251620249f091810191909152909350600092500160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156138d057600080fd5b505af11580156138e4573d6000803e3d6000fd5b50506025546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632259e9e591506139359087908690869060040161e43e565b600060405180830381600087803b15801561394f57600080fd5b505af1158015613963573d6000803e3d6000fd5b50506023546040517fd7fd7afb00000000000000000000000000000000000000000000000000000000815260048101889052600093506001600160a01b03909116915063d7fd7afb90602401602060405180830381865afa1580156139cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139f0919061e37f565b82806020019051810190613a04919061e37f565b613a0e919061e427565b6026546021546040516370a0823160e01b81526001600160a01b03928316600482018190529394506000933192849216906370a0823190602401602060405180830381865afa158015613a65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a89919061e37f565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015613b1a57600080fd5b505af1158015613b2e573d6000803e3d6000fd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201528b93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c90603401604051602081830303815290604052602260009054906101000a90046001600160a01b03168d898960405180604001604052808e806020019051810190613bcf919061e37f565b81526001602090910152604051613bee9695949392919060299061e508565b60405180910390a3602080546027546040516001600160a01b03928316936322f0ce89938e93613c39939116910160609190911b6bffffffffffffffffffffffff1916815260140190565b6040516020818303038152906040528b60296040518563ffffffff1660e01b8152600401613c699392919061e58a565b6000604051808303818588803b158015613c8257600080fd5b505af1158015613c96573d6000803e3d6000fd5b50506026546001600160a01b0316319250613cb8915061235890508b8561de9f565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015613d09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d2d919061e37f565b90506123e2612358878561de9f565b6060601e805480602002602001604051908101604052809291908181526020016000905b82821015613e7557600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015613e5e578382906000526020600020018054613dd19061deb2565b80601f0160208091040260200160405190810160405280929190818152602001828054613dfd9061deb2565b8015613e4a5780601f10613e1f57610100808354040283529160200191613e4a565b820191906000526020600020905b815481529060010190602001808311613e2d57829003601f168201915b505050505081526020019060010190613db2565b505050508152505081526020019060010190613d60565b50505050905090565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015613f3157600080fd5b505af1158015613f45573d6000803e3d6000fd5b505060208054604080516000808252602154606083018452620186a09583019586528284019190915291517f7b15118b0000000000000000000000000000000000000000000000000000000081526001600160a01b039384169650637b15118b955061097d9491936001931691889160299060040161e398565b6040516d98760000000000000000000000006020820152600190819060009060340160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561404857600080fd5b505af115801561405c573d6000803e3d6000fd5b50506025546021546040517fa8f2cb960000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063a8f2cb9693506140b39287921690869060019060040161e3ec565b600060405180830381600087803b1580156140cd57600080fd5b505af11580156140e1573d6000803e3d6000fd5b50506023546040517fd7fd7afb00000000000000000000000000000000000000000000000000000000815260048101869052620186a09350600092506001600160a01b039091169063d7fd7afb90602401602060405180830381865afa15801561414f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614173919061e37f565b61417d908361e427565b6026546021546040516370a0823160e01b81526001600160a01b03928316600482018190529394506000933192849216906370a0823190602401602060405180830381865afa1580156141d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141f8919061e37f565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561428957600080fd5b505af115801561429d573d6000803e3d6000fd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201528a93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f1981840301815260225483830183528b84526001602085015291516143369391926001600160a01b0316918f918c918c9160299061e508565b60405180910390a3602080546027546040516001600160a01b03928316936322f0ce89938d93614381939116910160609190911b6bffffffffffffffffffffffff1916815260140190565b6040516020818303038152906040528a60296040518563ffffffff1660e01b81526004016143b19392919061e58a565b6000604051808303818588803b1580156143ca57600080fd5b505af11580156143de573d6000803e3d6000fd5b50506026546001600160a01b0316319250614400915061235890508a8561de9f565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015614451573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614475919061e37f565b9050614484612358878561de9f565b50505050505050505050565b60606018805480602002602001604051908101604052809291908181526020018280548015612f1e576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311612f00575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015612f1e576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311612f00575050505050905090565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561460357600080fd5b505af1158015614617573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c34da12291506000906034015b604051602081830303815290604052600185602d60296040518763ffffffff1660e01b815260040161468e95949392919061e4dd565b6000604051808303818588803b1580156107ec57600080fd5b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561475a57600080fd5b505af115801561476e573d6000803e3d6000fd5b5050602080546040516001600160a01b03909116935063c34da122925060019101614658565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290516000936002936001600160a01b03169263bcbe936592600480830193928290030181865afa1580156147f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061481b919061e37f565b614825919061e651565b67ffffffffffffffff81111561483d5761483d61e1d8565b6040519080825280601f01601f191660200182016040528015614867576020820181803683370190505b5060208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263bcbe9365926004808401938290030181865afa1580156148cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906148f0919061e37f565b6148fa919061e651565b61490590600161ddb4565b67ffffffffffffffff81111561491d5761491d61e1d8565b6040519080825280601f01601f191660200182016040528015614947576020820181803683370190505b50602b90614955908261e24e565b506000602960020180546149689061deb2565b8351614974925061ddb4565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290519394506000936001600160a01b039092169263bcbe9365926004808401938290030181865afa1580156149d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906149fc919061e37f565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391614aa39160040161ddc7565b600060405180830381600087803b158015614abd57600080fd5b505af1158015614ad1573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352614b6c926001916001600160a01b0316908990602d9060299060040161e30d565b600060405180830381600087803b158015614b8657600080fd5b505af1158015614b9a573d6000803e3d6000fd5b50505050505050565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015614bf4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614c18919061e37f565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614c7457600080fd5b505af1158015614c88573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015614cfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614d1e919061e1b6565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b158015614d7d57600080fd5b505af1158015614d91573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352614e269287916001600160a01b03169060299060040161e617565b600060405180830381600087803b158015614e4057600080fd5b505af1158015614e54573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015614ea7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ecb919061e37f565b9050614ed78282619200565b505050565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b1790526000602d5551630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024016108b4565b6027546040516001600160a01b03909116602482015260009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561505457600080fd5b505af1158015615068573d6000803e3d6000fd5b505060215460265460275460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039283169450911691507f306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e49060340160408051601f19818403018152908290526150e8918690602d9060299061e68c565b60405180910390a3602080546027546040516001600160a01b03928316936306cb8983936109209316910160609190911b6bffffffffffffffffffffffff1916815260140190565b6022546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015615181573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906151a5919061e37f565b6022546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526101236004820152602481018390529192506001600160a01b03169063a9059cbb906044016020604051808303816000875af1158015615212573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615236919061e1b6565b506000600190507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561529b57600080fd5b505af11580156152af573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506322f0ce89915085906034016040516020818303038152906040528460296040518563ffffffff1660e01b815260040161531e9392919061e58a565b6000604051808303818588803b15801561533757600080fd5b505af115801561534b573d6000803e3d6000fd5b5050505050505050565b6027546040516001600160a01b03909116602482015260019060009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b1790526000602d5551630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561541357600080fd5b505af1158015615427573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c34da122915085906034016040516020818303038152906040528486602d60296040518763ffffffff1660e01b815260040161531e95949392919061e4dd565b6060601b805480602002602001604051908101604052809291908181526020016000905b82821015613e7557838290600052602060002090600202016040518060400160405290816000820180546154f29061deb2565b80601f016020809104026020016040519081016040528092919081815260200182805461551e9061deb2565b801561556b5780601f106155405761010080835404028352916020019161556b565b820191906000526020600020905b81548152906001019060200180831161554e57829003601f168201915b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156155f057602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116155b25790505b505050505081525050815260200190600101906154bf565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015615659573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061567d919061e37f565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156156d957600080fd5b505af11580156156ed573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af115801561575f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615783919061e1b6565b506027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905280517ff48448140000000000000000000000000000000000000000000000000000000081529051919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f48448149160048082019260009290919082900301818387803b15801561583157600080fd5b505af1158015615845573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352611c5c9288916001600160a01b0316908790602d9060299060040161e30d565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290516000936002936001600160a01b03169263bcbe936592600480830193928290030181865afa158015615942573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615966919061e37f565b615970919061e651565b67ffffffffffffffff8111156159885761598861e1d8565b6040519080825280601f01601f1916602001820160405280156159b2576020820181803683370190505b5060208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263bcbe9365926004808401938290030181865afa158015615a17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615a3b919061e37f565b615a45919061e651565b615a5090600161ddb4565b67ffffffffffffffff811115615a6857615a6861e1d8565b6040519080825280601f01601f191660200182016040528015615a92576020820181803683370190505b50602b90615aa0908261e24e565b50600060296002018054615ab39061deb2565b8351615abf925061ddb4565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290519394506000936001600160a01b039092169263bcbe9365926004808401938290030181865afa158015615b23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615b47919061e37f565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391615bee9160040161ddc7565b600060405180830381600087803b158015615c0857600080fd5b505af1158015615c1c573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a084526000602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b168152614b6c9391926001926001600160a01b03909116918a919060299060040161e398565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169263bcbe9365926004808401938290030181865afa158015615d27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615d4b919061e37f565b615d5690600161ddb4565b67ffffffffffffffff811115615d6e57615d6e61e1d8565b6040519080825280601f01601f191660200182016040528015615d98576020820181803683370190505b50602b90615da6908261e24e565b50600060296002018054615db99061deb2565b905090506000602060009054906101000a90046001600160a01b03166001600160a01b031663bcbe93656040518163ffffffff1660e01b8152600401602060405180830381865afa158015615e12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615e36919061e37f565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391615edd9160040161ddc7565b600060405180830381600087803b158015615ef757600080fd5b505af1158015615f0b573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506322f0ce899150600190603401604051602081830303815290604052600160296040518563ffffffff1660e01b8152600401615f7c9392919061e58a565b6000604051808303818588803b158015614b8657600080fd5b604051630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561600157600080fd5b505af1158015616015573d6000803e3d6000fd5b5050602080546040805160008152928301908190526021547f7c0dcb5f000000000000000000000000000000000000000000000000000000009091526001600160a01b039182169450637c0dcb5f935061607992916001911660296024840161e617565b600060405180830381600087803b15801561609357600080fd5b505af1158015611d0d573d6000803e3d6000fd5b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015613e755783829060005260206000200180546160ea9061deb2565b80601f01602080910402602001604051908101604052809291908181526020018280546161169061deb2565b80156161635780601f1061613857610100808354040283529160200191616163565b820191906000526020600020905b81548152906001019060200180831161614657829003601f168201915b5050505050815260200190600101906160cb565b6060601d805480602002602001604051908101604052809291908181526020016000905b82821015613e755760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561624557602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116162075790505b5050505050815250508152602001906001019061619b565b6027546040516001600160a01b039091166024820152600190819060009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a081018352610321808252600182840181905282850191909152835192830190935260008083526060820192909252929350919060808201906162f390621e84809061ddb4565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916163a69160040161ddc7565b600060405180830381600087803b1580156163c057600080fd5b505af11580156163d4573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c34da122915086906034016040516020818303038152906040528686602d876040518763ffffffff1660e01b815260040161644795949392919061e6db565b6000604051808303818588803b15801561646057600080fd5b505af1158015616474573d6000803e3d6000fd5b505050505050505050565b6060601c805480602002602001604051908101604052809291908181526020016000905b82821015613e755760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561654d57602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841161650f5790505b505050505081525050815260200190600101906164a3565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561661857600080fd5b505af115801561662c573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911693506306cb8983925001610920565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015613e755783829060005260206000200180546166929061deb2565b80601f01602080910402602001604051908101604052809291908181526020018280546166be9061deb2565b801561670b5780601f106166e05761010080835404028352916020019161670b565b820191906000526020600020905b8154815290600101906020018083116166ee57829003601f168201915b505050505081526020019060010190616673565b60085460009060ff1615616737575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa1580156167c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906167ec919061e37f565b1415905090565b604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561685f57600080fd5b505af1158015616873573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352616079926000916001600160a01b03169060299060040161e617565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156169bc57600080fd5b505af11580156169d0573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a0845260006020850181905292517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b16815261097d949293926001600160a01b0390921691889160299060040161e398565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015616b2f57600080fd5b505af1158015616b43573d6000803e3d6000fd5b5050602080546027546040805160609290921b6bffffffffffffffffffffffff1916938201939093528251808203601401815260215460748301855260006034840181815260549094015293517f7b15118b0000000000000000000000000000000000000000000000000000000081526001600160a01b039384169650637b15118b955061097d949193600193921691889160299060040161e398565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015616c9357600080fd5b505af1158015616ca7573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152602154838301909252916001916001600160a01b031690869080616d1885620186a061de9f565b815260006020909101526040517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b16815261097d95949392919060299060040161e398565b6040805160a08101825261032180825260016020808401829052838501929092528351918201909352600080825260608301919091529060808101616daa621e84808561ddb4565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391616e5d9160040161ddc7565b600060405180830381600087803b158015616e7757600080fd5b505af1158015616e8b573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526107d29287916001600160a01b031690879060040161e706565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290516000936002936001600160a01b03169263bcbe936592600480830193928290030181865afa158015616f82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616fa6919061e37f565b616fb0919061e651565b67ffffffffffffffff811115616fc857616fc861e1d8565b6040519080825280601f01601f191660200182016040528015616ff2576020820181803683370190505b5060208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263bcbe9365926004808401938290030181865afa158015617057573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061707b919061e37f565b617085919061e651565b61709090600161ddb4565b67ffffffffffffffff8111156170a8576170a861e1d8565b6040519080825280601f01601f1916602001820160405280156170d2576020820181803683370190505b50602b906170e0908261e24e565b506000602960020180546170f39061deb2565b83516170ff925061ddb4565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290519394506000936001600160a01b039092169263bcbe9365926004808401938290030181865afa158015617163573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617187919061e37f565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39161722e9160040161ddc7565b600060405180830381600087803b15801561724857600080fd5b505af115801561725c573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c34da1229150600190603401604051602081830303815290604052600187602d60296040518763ffffffff1660e01b815260040161531e95949392919061e4dd565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015617323573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617347919061e37f565b6021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526101236004820152602481018390529192506001600160a01b03169063a9059cbb906044016020604051808303816000875af11580156173b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906173d8919061e1b6565b5060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561743257600080fd5b505af1158015617446573d6000803e3d6000fd5b50506021546040517ff687d12a000000000000000000000000000000000000000000000000000000008152600a60048201526001600160a01b03909116925063f687d12a9150602401600060405180830381600087803b1580156174a957600080fd5b505af11580156174bd573d6000803e3d6000fd5b50506021546040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152600a6004820152600093508392506001600160a01b039091169063fc5fecd5906024016040805180830381865afa158015617527573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061754b919061e740565b60208054604080516001600160a01b03868116602483015290921660448301526064808301859052815180840390910181526084909201815291810180516001600160e01b03167f667011120000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152929450909250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916176099160040161ddc7565b600060405180830381600087803b15801561762357600080fd5b505af1158015617637573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526176cc9289916001600160a01b03169060299060040161e617565b600060405180830381600087803b15801561533757600080fd5b6040805160a081018252610321808252600160208084018290528385019290925283519182019093526000808252606083019190915282916080810161772f621e84808561ddb4565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916177e29160040161ddc7565b600060405180830381600087803b1580156177fc57600080fd5b505af1158015617810573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506322f0ce899150859060340160405160208183030381529060405285856040518563ffffffff1660e01b815260040161531e9392919061e76c565b604051630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156178ea57600080fd5b505af11580156178fe573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911693506322f0ce899250600091016133d7565b6022546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015617975573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617999919061e37f565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa1580156179eb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617a0f919061e37f565b60285460265460405163ca669fa760e01b81526001600160a01b039182166004820152929350163190600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015617a7457600080fd5b505af1158015617a88573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015617afa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617b1e919061e1b6565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b158015617b7d57600080fd5b505af1158015617b91573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506322f0ce89915087906034016040516020818303038152906040528460296040518563ffffffff1660e01b8152600401617c009392919061e58a565b6000604051808303818588803b158015617c1957600080fd5b505af1158015617c2d573d6000803e3d6000fd5b50506022546026546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319150602401602060405180830381865afa158015617c81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617ca5919061e37f565b9050617cb18582619200565b6022546020546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015617d02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617d26919061e37f565b9050617d328582619200565b602854614b9a9085906001600160a01b031631619200565b60606015805480602002602001604051908101604052809291908181526020018280548015612f1e576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311612f00575050505050905090565b6022546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015617dfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617e1f919061e37f565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa158015617e71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617e95919061e37f565b6028546027546040516001600160a01b03918216602482015292935016319060009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015617f4257600080fd5b505af1158015617f56573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015617fc8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617fec919061e1b6565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561804b57600080fd5b505af115801561805f573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c34da122915088906034016040516020818303038152906040528486602d60296040518763ffffffff1660e01b81526004016180d395949392919061e4dd565b6000604051808303818588803b1580156180ec57600080fd5b505af1158015618100573d6000803e3d6000fd5b50506022546026546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319150602401602060405180830381865afa158015618154573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618178919061e37f565b90506181848682619200565b6022546020546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156181d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906181f9919061e37f565b90506182058682619200565b60285461534b9086906001600160a01b031631619200565b6040516d98760000000000000000000000006020820152600a9060019060009060340160408051808303601f190181528282018252600883527f6761734c696d697400000000000000000000000000000000000000000000000060208085019190915282516002918101919091529093506000910160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156182f957600080fd5b505af115801561830d573d6000803e3d6000fd5b50506025546021546040517fa8f2cb960000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063a8f2cb9693506183649289921690889060019060040161e3ec565b600060405180830381600087803b15801561837e57600080fd5b505af1158015618392573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156183ef57600080fd5b505af1158015618403573d6000803e3d6000fd5b50506025546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632259e9e591506184549087908690869060040161e43e565b600060405180830381600087803b15801561846e57600080fd5b505af1158015618482573d6000803e3d6000fd5b505050506040518060400160405280600f81526020017f70726f746f636f6c466c617446656500000000000000000000000000000000008152509150600060036040516020016184db919060ff91909116815260200190565b6040516020818303038152906040529050600081806020019051810190618502919061e37f565b6023546040517fd7fd7afb000000000000000000000000000000000000000000000000000000008152600481018990526001600160a01b039091169063d7fd7afb90602401602060405180830381865afa158015618564573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618588919061e37f565b8480602001905181019061859c919061e37f565b6185a6919061e427565b6185b0919061ddb4565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561860c57600080fd5b505af1158015618620573d6000803e3d6000fd5b50506025546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632259e9e591506186719089908890879060040161e43e565b600060405180830381600087803b15801561868b57600080fd5b505af115801561869f573d6000803e3d6000fd5b50506026546021546040516370a0823160e01b81526001600160a01b0392831660048201819052319450600093509116906370a0823190602401602060405180830381865afa1580156186f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061871a919061e37f565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156187ab57600080fd5b505af11580156187bf573d6000803e3d6000fd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201528b93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c90603401604051602081830303815290604052602260009054906101000a90046001600160a01b03168d888a806020019051810190618854919061e37f565b60405180604001604052808e806020019051810190618873919061e37f565b815260016020909101526040516128f09695949392919060299061e508565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561894557600080fd5b505af1158015618959573d6000803e3d6000fd5b5050602080546040516001600160a01b039091169350637b15118b925001611915565b6027546040516001600160a01b03909116602482015260009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a08101835261032180825260018284018190528285019190915283519283019093526000808352606082019290925292935091906080820190618a0d90621e84809061ddb4565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391618ac09160040161ddc7565b600060405180830381600087803b158015618ada57600080fd5b505af1158015618aee573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526107d2926001916001600160a01b0316908890602d90899060040161e797565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015618bd9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618bfd919061e37f565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015618c8e57600080fd5b505af1158015618ca2573d6000803e3d6000fd5b505060265460275460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b039092169188916000918491634d8943bb916004808201926020929091908290030181865afa158015618d75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618d99919061e37f565b60408051808201825260008152600160208201529051618dc19695949392919060299061e508565b60405180910390a3602080546027546040516001600160a01b0392831693637c0dcb5f93618e099316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352618e619287916001600160a01b03169060299060040161e617565b600060405180830381600087803b158015618e7b57600080fd5b505af1158015618e8f573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015618ee2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618f06919061e37f565b9050614ed7612358848461de9f565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260029260009216906370a0823190602401602060405180830381865afa158015618f66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618f8a919061e37f565b6021549091506001600160a01b031663a9059cbb610123618fac60018561de9f565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af115801561900f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190619033919061e1b6565b5060215460265460208054604080516001600160a01b039586166024820181905294861660448201819052959092166064830181905260848084018990528251808503909101815260a4909301825292820180516001600160e01b03167f489ca9b700000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152929392737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391619104919060040161ddc7565b600060405180830381600087803b15801561911e57600080fd5b505af1158015619132573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526191c7928a916001600160a01b03169060299060040161e617565b600060405180830381600087803b15801561646057600080fd5b60006191eb61d91e565b6191f6848483619294565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c549060440160006040518083038186803b15801561926b57600080fd5b505afa158015610800573d6000803e3d6000fd5b61928761d91e565b6109ab858585848661930f565b6000806192a18584619402565b90506193046040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016192ef92919061e09f565b6040516020818303038152906040528561940e565b9150505b9392505050565b6040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201528190737109709ecfa91a80626ff3989d68f67f5b1dd12d9081906306447d5690602401600060405180830381600087803b15801561938157600080fd5b505af1925050508015619392575060015b6193a7576193a28787878761943c565b614b9a565b6193b38787878761943c565b806001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156193ee57600080fd5b505af11580156123e2573d6000803e3d6000fd5b60006193088383619455565b60c081015151600090156194325761942b84848460c00151619470565b9050619308565b61942b8484619616565b60006194488483619701565b90506109ab85828561970d565b60006194618383619ad7565b6193088383602001518461940e565b60008061947b619ae7565b905060006194898683619bba565b905060006194a0826060015183602001518561a060565b905060006194b08383898961a272565b905060006194bd8261b0ef565b602081015181519192509060030b15619530578982604001516040516020016194e792919061e7fb565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526195279160040161ddc7565b60405180910390fd5b60006195736040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a20000000000000000000000081525083600161b2be565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906195c690849060040161ddc7565b602060405180830381865afa1580156195e3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190619607919061e068565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc9259061966b90879060040161ddc7565b600060405180830381865afa158015619688573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526196b0919081019061e935565b905060006196de82856040516020016196ca92919061e96a565b60405160208183030381529060405261b4be565b90506001600160a01b0381166191f65784846040516020016194e792919061e999565b6000619461838361b4d1565b6040517f667f9d700000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90600090829063667f9d7090604401602060405180830381865afa1580156197a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906197cd919061e37f565b9050806199745760006197df8661b4dd565b604080518082018252600581527f352e302e300000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061986a905b6040805180820182526000808252602091820152815180830190925284518252808501908201529061b5c0565b80619876575060008451115b156198f9576040517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b03871690634f1ef286906198c2908890889060040161e09f565b600060405180830381600087803b1580156198dc57600080fd5b505af11580156198f0573d6000803e3d6000fd5b5050505061996e565b6040517f3659cfe60000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152871690633659cfe690602401600060405180830381600087803b15801561995557600080fd5b505af1158015619969573d6000803e3d6000fd5b505050505b506109ab565b8060006199808261b4dd565b604080518082018252600581527f352e302e30000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506199e29061983d565b806199ee575060008551115b15619a73576040517f9623609d0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690639623609d90619a3c908a908a908a9060040161ea44565b600060405180830381600087803b158015619a5657600080fd5b505af1158015619a6a573d6000803e3d6000fd5b50505050614b9a565b6040517f99a88ec40000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015287811660248301528316906399a88ec490604401600060405180830381600087803b1580156193ee57600080fd5b619ae38282600061b5d4565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c90619b6e90849060040161ea75565b600060405180830381865afa158015619b8b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619bb3919081019061eabc565b9250505090565b619bec6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d9050619c376040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b619c408561b6d7565b60208201526000619c508661babc565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa158015619c92573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619cba919081019061eabc565b86838560200151604051602001619cd4949392919061eb05565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb1190619d2c90859060040161ddc7565b600060405180830381865afa158015619d49573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619d71919081019061eabc565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f690619db990849060040161ec09565b602060405180830381865afa158015619dd6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190619dfa919061e1b6565b619e0f57816040516020016194e7919061ec5b565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890619e5490849060040161eced565b600060405180830381865afa158015619e71573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619e99919081019061eabc565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f690619ee090849060040161ed3f565b602060405180830381865afa158015619efd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190619f21919061e1b6565b15619fb6576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890619f6b90849060040161ed3f565b600060405180830381865afa158015619f88573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619fb0919081019061eabc565b60408501525b846001600160a01b03166349c4fac8828660000151604051602001619fdb919061ed91565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161a00792919061edfd565b600060405180830381865afa15801561a024573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a04c919081019061eabc565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b606081526020019060019003908161a07c5790505090506040518060400160405280600481526020017f67726570000000000000000000000000000000000000000000000000000000008152508160008151811061a0dc5761a0dc61ee22565b60200260200101819052506040518060400160405280600381526020017f2d726c00000000000000000000000000000000000000000000000000000000008152508160018151811061a1305761a13061ee22565b60200260200101819052508460405160200161a14c919061ee51565b6040516020818303038152906040528160028151811061a16e5761a16e61ee22565b60200260200101819052508260405160200161a18a919061eebd565b6040516020818303038152906040528160038151811061a1ac5761a1ac61ee22565b6020026020010181905250600061a1c28261b0ef565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184526000808252908601528251808401909352905182529281019290925291925061a253906040805180820182526000808252602091820152815180830190925284518252808501908201529061bd3f565b61a26857856040516020016194e7919061eefe565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d901561a2c2565b511590565b61a4365782602001511561a37e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401619527565b8260c001511561a436576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401619527565b6040805160ff8082526120008201909252600091816020015b606081526020019060019003908161a44f57905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061a4aa9061ef8f565b935060ff168151811061a4bf5761a4bf61ee22565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e370000000000000000000000000000000000000081525060405160200161a510919061efae565b60405160208183030381529060405282828061a52b9061ef8f565b935060ff168151811061a5405761a54061ee22565b60200260200101819052506040518060400160405280600681526020017f6465706c6f79000000000000000000000000000000000000000000000000000081525082828061a58d9061ef8f565b935060ff168151811061a5a25761a5a261ee22565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d6500000000000000000000000000000000000081525082828061a5ef9061ef8f565b935060ff168151811061a6045761a60461ee22565b6020026020010181905250876020015182828061a6209061ef8f565b935060ff168151811061a6355761a63561ee22565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163745061746800000000000000000000000000000000000081525082828061a6829061ef8f565b935060ff168151811061a6975761a69761ee22565b60209081029190910101528751828261a6af8161ef8f565b935060ff168151811061a6c45761a6c461ee22565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e4964000000000000000000000000000000000000000000000081525082828061a7119061ef8f565b935060ff168151811061a7265761a72661ee22565b602002602001018190525061a73a4661bda0565b828261a7458161ef8f565b935060ff168151811061a75a5761a75a61ee22565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c65000000000000000000000000000000000081525082828061a7a79061ef8f565b935060ff168151811061a7bc5761a7bc61ee22565b60200260200101819052508682828061a7d49061ef8f565b935060ff168151811061a7e95761a7e961ee22565b602090810291909101015285511561a9105760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f646500000000000000000000006020820152828261a83a8161ef8f565b935060ff168151811061a84f5761a84f61ee22565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d9061a89f90899060040161ddc7565b600060405180830381865afa15801561a8bc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a8e4919081019061eabc565b828261a8ef8161ef8f565b935060ff168151811061a9045761a90461ee22565b60200260200101819052505b84602001511561a9e05760408051808201909152601281527f2d2d766572696679536f75726365436f646500000000000000000000000000006020820152828261a9598161ef8f565b935060ff168151811061a96e5761a96e61ee22565b60200260200101819052506040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525082828061a9bb9061ef8f565b935060ff168151811061a9d05761a9d061ee22565b602002602001018190525061aba7565b61aa1861a2bd8660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b61aaab5760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261aa5b8161ef8f565b935060ff168151811061aa705761aa7061ee22565b60200260200101819052508460a0015160405160200161aa90919061ee51565b60405160208183030381529060405282828061a9bb9061ef8f565b8460c0015115801561aaee57506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261aaec90511590565b155b1561aba75760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261ab328161ef8f565b935060ff168151811061ab475761ab4761ee22565b602002602001018190525061ab5b8861be40565b60405160200161ab6b919061ee51565b60405160208183030381529060405282828061ab869061ef8f565b935060ff168151811061ab9b5761ab9b61ee22565b60200260200101819052505b6040808601518151808301835260008082526020918201528251808401909352815183529081019082015261abdb90511590565b61ac705760408051808201909152600b81527f2d2d72656c6179657249640000000000000000000000000000000000000000006020820152828261ac1e8161ef8f565b935060ff168151811061ac335761ac3361ee22565b6020026020010181905250846040015182828061ac4f9061ef8f565b935060ff168151811061ac645761ac6461ee22565b60200260200101819052505b60608501511561ad915760408051808201909152600681527f2d2d73616c7400000000000000000000000000000000000000000000000000006020820152828261acb98161ef8f565b935060ff168151811061acce5761acce61ee22565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa15801561ad3d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261ad65919081019061eabc565b828261ad708161ef8f565b935060ff168151811061ad855761ad8561ee22565b60200260200101819052505b60e0850151511561ae385760408051808201909152600a81527f2d2d6761734c696d6974000000000000000000000000000000000000000000006020820152828261addb8161ef8f565b935060ff168151811061adf05761adf061ee22565b602002602001018190525061ae0c8560e001516000015161bda0565b828261ae178161ef8f565b935060ff168151811061ae2c5761ae2c61ee22565b60200260200101819052505b60e0850151602001511561aee25760408051808201909152600a81527f2d2d6761735072696365000000000000000000000000000000000000000000006020820152828261ae858161ef8f565b935060ff168151811061ae9a5761ae9a61ee22565b602002602001018190525061aeb68560e001516020015161bda0565b828261aec18161ef8f565b935060ff168151811061aed65761aed661ee22565b60200260200101819052505b60e0850151604001511561af8c5760408051808201909152600e81527f2d2d6d61784665655065724761730000000000000000000000000000000000006020820152828261af2f8161ef8f565b935060ff168151811061af445761af4461ee22565b602002602001018190525061af608560e001516040015161bda0565b828261af6b8161ef8f565b935060ff168151811061af805761af8061ee22565b60200260200101819052505b60e0850151606001511561b0365760408051808201909152601681527f2d2d6d61785072696f72697479466565506572476173000000000000000000006020820152828261afd98161ef8f565b935060ff168151811061afee5761afee61ee22565b602002602001018190525061b00a8560e001516060015161bda0565b828261b0158161ef8f565b935060ff168151811061b02a5761b02a61ee22565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561b0545761b05461e1d8565b60405190808252806020026020018201604052801561b08757816020015b606081526020019060019003908161b0725790505b50905060005b8260ff168160ff16101561b0e057838160ff168151811061b0b05761b0b061ee22565b6020026020010151828260ff168151811061b0cd5761b0cd61ee22565b602090810291909101015260010161b08d565b5093505050505b949350505050565b61b1166040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c9161b19c9186910161f019565b600060405180830381865afa15801561b1b9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261b1e1919081019061eabc565b9050600061b1ef868361c92f565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b815260040161b21f919061dc77565b6000604051808303816000875af115801561b23e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261b266919081019061f060565b805190915060030b1580159061b27f5750602081015151155b801561b28e5750604081015151155b1561a268578160008151811061b2a65761b2a661ee22565b60200260200101516040516020016194e7919061f116565b6060600061b2f38560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252865182528087019082015290915061b32a9082905b9061ca84565b1561b48757600061b3a78261b3a18461b39b61b36d8a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061caab565b9061cb0d565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061b40b90829061ca84565b1561b47557604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b472905b829061cb92565b90505b61b47e8161cbb8565b92505050619308565b821561b4a05784846040516020016194e792919061f302565b5050604080516020810190915260008152619308565b509392505050565b6000808251602084016000f09392505050565b619ae38282600161b5d4565b60408051600481526024810182526020810180516001600160e01b03167fad3cb1cc00000000000000000000000000000000000000000000000000000000179052905160609160009182916001600160a01b0386169161b53d919061e083565b600060405180830381855afa9150503d806000811461b578576040519150601f19603f3d011682016040523d82523d6000602084013e61b57d565b606091505b509150915081801561b590575060208151115b1561b5a9578080602001905181019061b0e7919061eabc565b505060408051602081019091526000815292915050565b600061b5cc838361cc21565b159392505050565b8160a001511561b5e357505050565b600061b5f084848461ccfc565b9050600061b5fd8261b0ef565b602081015181519192509060030b15801561b6995750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b6999060408051808201825260008082526020918201528151808301909252845182528085019082015261b324565b1561b6a657505050505050565b6040820151511561b6c65781604001516040516020016194e7919061f3a9565b806040516020016194e7919061f407565b6060600061b70c8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061b771905b829061bd3f565b1561b7e057604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526193089061b7db90839061d297565b61cbb8565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b842905b829061d321565b60010361b90f57604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b8a89061b46b565b50604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526193089061b7db905b839061cb92565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b96e9061b76a565b1561baa557604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061b9d690839061d3bb565b90506000816001835161b9e9919061de9f565b8151811061b9f95761b9f961ee22565b6020026020010151905061ba9c61b7db61ba6f6040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925285518252808601908201529061d297565b95945050505050565b826040516020016194e7919061f472565b50919050565b6060600061baf18360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061bb539061b76a565b1561bb61576193088161cbb8565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bbc09061b83b565b60010361bc2a57604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526193089061b7db9061b908565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bc899061b76a565b1561baa557604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061bcf190839061d3bb565b905060018151111561bd2d57806002825161bd0c919061de9f565b8151811061bd1c5761bd1c61ee22565b602002602001015192505050919050565b50826040516020016194e7919061f472565b80518251600091111561bd54575060006191fa565b8151835160208501516000929161bd6a9161ddb4565b61bd74919061de9f565b90508260200151810361bd8b5760019150506191fa565b82516020840151819020912014905092915050565b6060600061bdad8361d460565b600101905060008167ffffffffffffffff81111561bdcd5761bdcd61e1d8565b6040519080825280601f01601f19166020018201604052801561bdf7576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461be0157509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e534544000000000000000000000000000000000000000000008184019081528551808701875283815284019290925284518086019095525184529083015260609161becc905b829061b5c0565b1561bf0c57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bf6b9061bec5565b1561bfab57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c00a9061bec5565b1561c04a57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c0a99061bec5565b8061c10e5750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c10e9061bec5565b1561c14e57505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c1ad9061bec5565b8061c2125750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c2129061bec5565b1561c25257505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c2b19061bec5565b8061c3165750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c3169061bec5565b1561c35657505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c3b59061bec5565b8061c41a5750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c41a9061bec5565b1561c45a57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c4b99061bec5565b1561c4f957505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c5589061bec5565b1561c59857505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c5f79061bec5565b1561c63757505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c6969061bec5565b1561c6d657505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c7359061bec5565b1561c77557505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c7d49061bec5565b8061c8395750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c8399061bec5565b1561c87957505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c8d89061bec5565b1561c91857505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516194e7929060200161f550565b60608060005b845181101561c9ba578185828151811061c9515761c95161ee22565b602002602001015160405160200161c96a92919061e96a565b60405160208183030381529060405291506001855161c989919061de9f565b811461c9b2578160405160200161c9a0919061f6b9565b60405160208183030381529060405291505b60010161c935565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161c9d3579050509050838160008151811061c9fe5761c9fe61ee22565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061ca525761ca5261ee22565b6020026020010181905250818160028151811061ca715761ca7161ee22565b6020908102919091010152949350505050565b602080830151835183519284015160009361caa2929184919061d542565b14159392505050565b6040805180820190915260008082526020820152600061cadd846000015185602001518560000151866020015161d653565b905083602001518161caef919061de9f565b8451859061cafe90839061de9f565b90525060208401525090919050565b604080518082019091526000808252602082015281518351101561cb325750816191fa565b602080830151908401516001911461cb595750815160208481015190840151829020919020145b801561cb8a5782518451859061cb7090839061de9f565b905250825160208501805161cb8690839061ddb4565b9052505b509192915050565b604080518082019091526000808252602082015261cbb183838361d773565b5092915050565b60606000826000015167ffffffffffffffff81111561cbd95761cbd961e1d8565b6040519080825280601f01601f19166020018201604052801561cc03576020820181803683370190505b509050600060208201905061cbb1818560200151866000015161d81e565b815181516000919081111561cc34575081515b6020808501519084015160005b8381101561cced578251825180821461ccbd57600019602087101561cc9c5760018461cc6e89602061de9f565b61cc78919061ddb4565b61cc8390600861e427565b61cc8e90600261f7e1565b61cc98919061de9f565b1990505b818116838216818103911461ccba5797506191fa9650505050505050565b50505b61ccc860208661ddb4565b945061ccd560208561ddb4565b9350505060208161cce6919061ddb4565b905061cc41565b508451865161a268919061f7ed565b6060600061cd08619ae7565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161cd2557905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061cd809061ef8f565b935060ff168151811061cd955761cd9561ee22565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161cde6919061f80d565b60405160208183030381529060405282828061ce019061ef8f565b935060ff168151811061ce165761ce1661ee22565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061ce639061ef8f565b935060ff168151811061ce785761ce7861ee22565b60200260200101819052508260405160200161ce94919061eebd565b60405160208183030381529060405282828061ceaf9061ef8f565b935060ff168151811061cec45761cec461ee22565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061cf119061ef8f565b935060ff168151811061cf265761cf2661ee22565b602002602001018190525061cf3b878461d898565b828261cf468161ef8f565b935060ff168151811061cf5b5761cf5b61ee22565b60209081029190910101528551511561d0075760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261cfad8161ef8f565b935060ff168151811061cfc25761cfc261ee22565b602002602001018190525061cfdb86600001518461d898565b828261cfe68161ef8f565b935060ff168151811061cffb5761cffb61ee22565b60200260200101819052505b85608001511561d0755760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261d0508161ef8f565b935060ff168151811061d0655761d06561ee22565b602002602001018190525061d0db565b841561d0db5760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261d0ba8161ef8f565b935060ff168151811061d0cf5761d0cf61ee22565b60200260200101819052505b6040860151511561d1775760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261d1258161ef8f565b935060ff168151811061d13a5761d13a61ee22565b6020026020010181905250856040015182828061d1569061ef8f565b935060ff168151811061d16b5761d16b61ee22565b60200260200101819052505b85606001511561d1e15760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261d1c08161ef8f565b935060ff168151811061d1d55761d1d561ee22565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561d1ff5761d1ff61e1d8565b60405190808252806020026020018201604052801561d23257816020015b606081526020019060019003908161d21d5790505b50905060005b8260ff168160ff16101561d28b57838160ff168151811061d25b5761d25b61ee22565b6020026020010151828260ff168151811061d2785761d27861ee22565b602090810291909101015260010161d238565b50979650505050505050565b604080518082019091526000808252602082015281518351101561d2bc5750816191fa565b8151835160208501516000929161d2d29161ddb4565b61d2dc919061de9f565b6020840151909150600190821461d2fd575082516020840151819020908220145b801561d3185783518551869061d31490839061de9f565b9052505b50929392505050565b600080826000015161d345856000015186602001518660000151876020015161d653565b61d34f919061ddb4565b90505b8351602085015161d363919061ddb4565b811161cbb1578161d3738161f852565b925050826000015161d3aa85602001518361d38e919061de9f565b865161d39a919061de9f565b838660000151876020015161d653565b61d3b4919061ddb4565b905061d352565b6060600061d3c9848461d321565b61d3d490600161ddb4565b67ffffffffffffffff81111561d3ec5761d3ec61e1d8565b60405190808252806020026020018201604052801561d41f57816020015b606081526020019060019003908161d40a5790505b50905060005b815181101561b4b65761d43b61b7db868661cb92565b82828151811061d44d5761d44d61ee22565b602090810291909101015260010161d425565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061d4a9577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061d4d5576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061d4f357662386f26fc10000830492506010015b6305f5e100831061d50b576305f5e100830492506008015b612710831061d51f57612710830492506004015b6064831061d531576064830492506002015b600a83106191fa5760010192915050565b60008085841161d649576020841161d5f5576000841561d58d57600161d56986602061de9f565b61d57490600861e427565b61d57f90600261f7e1565b61d589919061de9f565b1990505b835181168561d59c898961ddb4565b61d5a6919061de9f565b805190935082165b81811461d5e05787841161d5c8578794505050505061b0e7565b8361d5d28161f86c565b94505082845116905061d5ae565b61d5ea878561ddb4565b94505050505061b0e7565b83832061d602858861de9f565b61d60c908761ddb4565b91505b85821061d6475784822080820361d6345761d62a868461ddb4565b935050505061b0e7565b61d63f60018461de9f565b92505061d60f565b505b5092949350505050565b6000838186851161d75e576020851161d70d576000851561d69f57600161d67b87602061de9f565b61d68690600861e427565b61d69190600261f7e1565b61d69b919061de9f565b1990505b8451811660008761d6b08b8b61ddb4565b61d6ba919061de9f565b855190915083165b82811461d6ff5781861061d6e75761d6da8b8b61ddb4565b965050505050505061b0e7565b8561d6f18161f852565b96505083865116905061d6c2565b85965050505050505061b0e7565b508383206000905b61d71f868961de9f565b821161d75c5785832080820361d73b578394505050505061b0e7565b61d74660018561ddb4565b935050818061d7549061f852565b92505061d715565b505b61d768878761ddb4565b979650505050505050565b6040805180820190915260008082526020820152600061d7a5856000015186602001518660000151876020015161d653565b60208087018051918601919091525190915061d7c1908261de9f565b83528451602086015161d7d4919061ddb4565b810361d7e3576000855261d815565b8351835161d7f1919061ddb4565b8551869061d80090839061de9f565b905250835161d80f908261ddb4565b60208601525b50909392505050565b6020811061d856578151835261d83560208461ddb4565b925061d84260208361ddb4565b915061d84f60208261de9f565b905061d81e565b600019811561d88557600161d86c83602061de9f565b61d8789061010061f7e1565b61d882919061de9f565b90505b9151835183169219169190911790915250565b6060600061d8a68484619bba565b805160208083015160405193945061d8c09390910161f883565b60405160208183030381529060405291505092915050565b610b67806200f8dc83390190565b6151c8806201044383390190565b610a42806201560b83390190565b61106f806201604d83390190565b61207280620170bc83390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161d96161d966565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161d9616040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561da185783516001600160a01b031683526020938401939092019160010161d9f1565b509095945050505050565b60005b8381101561da3e57818101518382015260200161da26565b50506000910152565b6000815180845261da5f81602086016020860161da23565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561db6f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561db55577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261db3f84865161da47565b602095860195909450929092019160010161db05565b50919750505060209485019492909201915060010161da9b565b50929695505050505050565b600081518084526020840193506020830160005b8281101561dbcf5781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161db8f565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561db6f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261dc45604088018261da47565b905060208201519150868103602088015261dc60818361db7b565b96505050602093840193919091019060010161dc01565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561db6f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261dcd985835161da47565b9450602093840193919091019060010161dc9f565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561db6f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261dd6f604087018261db7b565b955050602093840193919091019060010161dd16565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156191fa576191fa61dd85565b602081526000619308602083018461da47565b6001600160a01b0381511682526020810151151560208301526001600160a01b0360408201511660408301526000606082015160a0606085015261de2160a085018261da47565b608093840151949093019390935250919050565b60c08152600061de4860c083018861da47565b6001600160a01b0387166020840152828103604084015261de69818761da47565b85546060850152600186015460ff1615156080850152905082810360a084015261de93818561ddda565b98975050505050505050565b818103818111156191fa576191fa61dd85565b600181811c9082168061dec657607f821691505b60208210810361bab6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600081546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160a060608501526000815461df498161deb2565b8060a0880152600182166000811461df68576001811461dfa25761dfd6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660c089015260c082151560051b890101935061dfd6565b84600052602060002060005b8381101561dfcd5781548a820160c0015260019091019060200161dfae565b890160c0019450505b50505060038401546080860152809250505092915050565b60c08152600061e00160c083018861da47565b6001600160a01b0387166020840152828103604084015261e022818761da47565b85546060850152600186015460ff1615156080850152905082810360a084015261de93818561deff565b80516001600160a01b038116811461e06357600080fd5b919050565b60006020828403121561e07a57600080fd5b6193088261e04c565b6000825161e09581846020870161da23565b9190910192915050565b6001600160a01b038316815260406020820152600061b0e7604083018461da47565b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff881660408301528660608301526003861061e17b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8560808301528460a083015261e19c60c08301856001600160a01b03169052565b6001600160a01b03831660e0830152979650505050505050565b60006020828403121561e1c857600080fd5b8151801515811461930857600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f821115614ed757806000526020600020601f840160051c8101602085101561e22e5750805b601f840160051c820191505b818110156109ab576000815560010161e23a565b815167ffffffffffffffff81111561e2685761e26861e1d8565b61e27c8161e276845461deb2565b8461e207565b6020601f82116001811461e2b0576000831561e2985750848201515b600019600385901b1c1916600184901b1784556109ab565b600084815260208120601f198516915b8281101561e2e0578785015182556020948501946001909201910161e2c0565b508482101561e2fe5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60e08152600061e32060e083018961da47565b8760208401526001600160a01b0387166040840152828103606084015261e347818761da47565b85546080850152600186015460ff16151560a085015290505b82810360c084015261e372818561deff565b9998505050505050505050565b60006020828403121561e39157600080fd5b5051919050565b60e08152600061e3ab60e083018961da47565b8760208401526001600160a01b0387166040840152828103606084015261e3d2818761da47565b855160808501526020860151151560a0850152905061e360565b8481526001600160a01b038416602082015260806040820152600061e414608083018561da47565b9050821515606083015295945050505050565b80820281158282048414176191fa576191fa61dd85565b83815260606020820152600061e457606083018561da47565b828103604084015261a268818561da47565b6101208152600061e47e61012083018b61da47565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a084015261e4b1818761da47565b855460c0850152600186015460ff16151560e085015290505b828103610100840152619607818561deff565b60c08152600061e4f060c083018861da47565b866020840152828103604084015261e022818761da47565b6101208152600061e51d61012083018a61da47565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a08501526000825261e56460c0850187805182526020908101511515910152565b602081016101008501525061e57c602082018561deff565b9a9950505050505050505050565b60608152600061e59d606083018661da47565b846020840152828103604084015261a268818561deff565b6101208152600061e5ca61012083018b61da47565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a084015261e5fd818761da47565b855160c08501526020860151151560e0850152905061e4ca565b60808152600061e62a608083018761da47565b8560208401526001600160a01b0385166040840152828103606084015261d768818561deff565b60008261e687577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60a08152600061e69f60a083018761da47565b828103602084015261e6b1818761da47565b85546040850152600186015460ff16151560608501529050828103608084015261d768818561deff565b60c08152600061e6ee60c083018861da47565b866020840152828103604084015261de69818761da47565b60808152600061e719608083018761da47565b8560208401526001600160a01b0385166040840152828103606084015261d768818561ddda565b6000806040838503121561e75357600080fd5b61e75c8361e04c565b6020939093015192949293505050565b60608152600061e77f606083018661da47565b846020840152828103604084015261a268818561ddda565b60e08152600061e7aa60e083018961da47565b8760208401526001600160a01b0387166040840152828103606084015261e7d1818761da47565b85546080850152600186015460ff16151560a0850152905082810360c084015261e372818561ddda565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161e83381601a85016020880161da23565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161e87081601c84016020880161da23565b01601c01949350505050565b6040516060810167ffffffffffffffff8111828210171561e89f5761e89f61e1d8565b60405290565b60008067ffffffffffffffff84111561e8c05761e8c061e1d8565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561e8ef5761e8ef61e1d8565b60405283815290508082840185101561e90757600080fd5b61b4b684602083018561da23565b600082601f83011261e92657600080fd5b6193088383516020850161e8a5565b60006020828403121561e94757600080fd5b815167ffffffffffffffff81111561e95e57600080fd5b6191f68482850161e915565b6000835161e97c81846020880161da23565b83519083019061e99081836020880161da23565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161e9d181601a85016020880161da23565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161ea0e81603384016020880161da23565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b6001600160a01b03841681526001600160a01b038316602082015260606040820152600061ba9c606083018461da47565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000619308608083018461da47565b60006020828403121561eace57600080fd5b815167ffffffffffffffff81111561eae557600080fd5b8201601f8101841361eaf657600080fd5b6191f68482516020840161e8a5565b6000855161eb17818460208a0161da23565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161eb51816001840160208a0161da23565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161eb8f81600284016020890161da23565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161ebd181600284016020880161da23565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061ec1c604083018461da47565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161ec9381601f85016020870161da23565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061ed00604083018461da47565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061ed52604083018461da47565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161edc981601485016020870161da23565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061ee10604083018561da47565b8281036020840152619304818561da47565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161ee8981600185016020870161da23565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161eecf81846020870161da23565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161ef8281604b85016020870161da23565b91909101604b0192915050565b600060ff821660ff810361efa55761efa561dd85565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161f00c81602985016020870161da23565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000619308608083018461da47565b60006020828403121561f07257600080fd5b815167ffffffffffffffff81111561f08957600080fd5b82016060818503121561f09b57600080fd5b61f0a361e87c565b81518060030b811461f0b457600080fd5b8152602082015167ffffffffffffffff81111561f0d057600080fd5b61f0dc8682850161e915565b602083015250604082015167ffffffffffffffff81111561f0fc57600080fd5b61f1088682850161e915565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161f17481602185016020870161da23565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161f36081602185016020880161da23565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161f39d81602e84016020880161da23565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161f00c81602985016020870161da23565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161f46581602285016020870161da23565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161f4aa81600e85016020870161da23565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161f58881601885016020880161da23565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161f5c581601c84016020880161da23565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161f6cb81846020870161da23565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b6001815b600184111561f7355780850481111561f7195761f71961dd85565b600184161561f72757908102905b60019390931c92800261f6fe565b935093915050565b60008261f74c575060016191fa565b8161f759575060006191fa565b816001811461f76f576002811461f7795761f795565b60019150506191fa565b60ff84111561f78a5761f78a61dd85565b50506001821b6191fa565b5060208310610133831016604e8410600b841016171561f7b8575081810a6191fa565b61f7c5600019848461f6fa565b806000190482111561f7d95761f7d961dd85565b029392505050565b6000619308838361f73d565b818103600083128015838313168383128216171561cbb15761cbb161dd85565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161f84581601c85016020870161da23565b91909101601c0192915050565b6000600019820361f8655761f86561dd85565b5060010190565b60008161f87b5761f87b61dd85565b506000190190565b6000835161f89581846020880161da23565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161f8cf81600184016020880161da23565b0160010194935050505056fe60c0604052600d60809081526c2bb930b83832b21022ba3432b960991b60a05260009061002c9082610114565b506040805180820190915260048152630ae8aa8960e31b60208201526001906100559082610114565b506002805460ff1916601217905534801561006f57600080fd5b506101d2565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061009f57607f821691505b6020821081036100bf57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561010f57806000526020600020601f840160051c810160208510156100ec5750805b601f840160051c820191505b8181101561010c57600081556001016100f8565b50505b505050565b81516001600160401b0381111561012d5761012d610075565b6101418161013b845461008b565b846100c5565b6020601f821160018114610175576000831561015d5750848201515b600019600385901b1c1916600184901b17845561010c565b600084815260208120601f198516915b828110156101a55787850151825560209485019460019092019101610185565b50848210156101c35786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610986806101e16000396000f3fe6080604052600436106100c05760003560e01c8063313ce56711610074578063a9059cbb1161004e578063a9059cbb146101fa578063d0e30db01461021a578063dd62ed3e1461022257600080fd5b8063313ce5671461018c57806370a08231146101b857806395d89b41146101e557600080fd5b806318160ddd116100a557806318160ddd1461012f57806323b872dd1461014c5780632e1a7d4d1461016c57600080fd5b806306fdde03146100d4578063095ea7b3146100ff57600080fd5b366100cf576100cd61025a565b005b600080fd5b3480156100e057600080fd5b506100e96102b5565b6040516100f69190610745565b60405180910390f35b34801561010b57600080fd5b5061011f61011a3660046107da565b610343565b60405190151581526020016100f6565b34801561013b57600080fd5b50475b6040519081526020016100f6565b34801561015857600080fd5b5061011f610167366004610804565b6103bd565b34801561017857600080fd5b506100cd610187366004610841565b610647565b34801561019857600080fd5b506002546101a69060ff1681565b60405160ff90911681526020016100f6565b3480156101c457600080fd5b5061013e6101d336600461085a565b60036020526000908152604090205481565b3480156101f157600080fd5b506100e9610724565b34801561020657600080fd5b5061011f6102153660046107da565b610731565b6100cd61025a565b34801561022e57600080fd5b5061013e61023d366004610875565b600460209081526000928352604080842090915290825290205481565b33600090815260036020526040812080543492906102799084906108d7565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b600080546102c2906108ea565b80601f01602080910402602001604051908101604052809291908181526020018280546102ee906108ea565b801561033b5780601f106103105761010080835404028352916020019161033b565b820191906000526020600020905b81548152906001019060200180831161031e57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103ab9086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604081205482111561042b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841633148015906104a1575073ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156105605773ffffffffffffffffffffffffffffffffffffffff8416600090815260046020908152604080832033845290915290205482111561051a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b73ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091528120805484929061055a90849061093d565b90915550505b73ffffffffffffffffffffffffffffffffffffffff84166000908152600360205260408120805484929061059590849061093d565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040812080548492906105cf9084906108d7565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161063591815260200190565b60405180910390a35060019392505050565b3360009081526003602052604090205481111561069a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b33600090815260036020526040812080548392906106b990849061093d565b9091555050604051339082156108fc029083906000818181858888f193505050501580156106eb573d6000803e3d6000fd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b600180546102c2906108ea565b600061073e3384846103bd565b9392505050565b602081526000825180602084015260005b818110156107735760208186018101516040868401015201610756565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146107d557600080fd5b919050565b600080604083850312156107ed57600080fd5b6107f6836107b1565b946020939093013593505050565b60008060006060848603121561081957600080fd5b610822846107b1565b9250610830602085016107b1565b929592945050506040919091013590565b60006020828403121561085357600080fd5b5035919050565b60006020828403121561086c57600080fd5b61073e826107b1565b6000806040838503121561088857600080fd5b610891836107b1565b915061089f602084016107b1565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156103b7576103b76108a8565b600181811c908216806108fe57607f821691505b602082108103610937577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b818103818111156103b7576103b76108a856fea2646970667358221220bb5b9dcef0ba90bcdefcbd63f71b1df95b50e29550a7456c69c6b9ff9dcdd20e64736f6c634300081a003360a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516150cb6100fd60003960008181612898015281816128c10152612a7401526150cb6000f3fe6080604052600436106102a05760003560e01c80638f2839701161016e578063bd8fde1c116100cb578063d547741f1161007f578063e9d6c5ba11610064578063e9d6c5ba1461082d578063f354b31f1461085f578063f851a4401461087f57600080fd5b8063d547741f146107d9578063e63ab1e9146107f957600080fd5b8063c1bd469f116100b0578063c1bd469f14610777578063cc5ad8b614610799578063d3523ea2146107b957600080fd5b8063bd8fde1c14610723578063c0c53b8b1461075757600080fd5b8063a217fddf11610122578063a8f2cb9611610107578063a8f2cb961461069a578063aa808c06146106ba578063ad3cb1cc146106da57600080fd5b8063a217fddf1461066e578063a3ebd14c1461068357600080fd5b806391d148541161015357806391d14854146105c557806394cc86831461062a5780639ca220dd1461064c57600080fd5b80638f283970146105855780639060bda9146105a557600080fd5b80633f4ba83a1161021c578063631d62e4116101d05780637066b18d116101b55780637066b18d14610515578063804ea334146105425780638456cb591461057057600080fd5b8063631d62e4146104d55780636e9e2d3f146104f557600080fd5b806352d1902d1161020157806352d1902d1461045b5780635c975abb146104705780635cf92c9f146104a757600080fd5b80633f4ba83a146104335780634f1ef2861461044857600080fd5b80632259e9e5116102735780632f2ff15d116102585780632f2ff15d146103d35780633500c24b146103f357806336568abe1461041357600080fd5b80632259e9e514610356578063248a9ca31461037657600080fd5b806301ffc9a7146102a55780630c63109e146102da57806310d29b9e1461031257806318d3ce9614610334575b600080fd5b3480156102b157600080fd5b506102c56102c0366004613fd6565b61089f565b60405190151581526020015b60405180910390f35b3480156102e657600080fd5b506001546102fa906001600160a01b031681565b6040516001600160a01b0390911681526020016102d1565b34801561031e57600080fd5b5061033261032d36600461406f565b610938565b005b34801561034057600080fd5b506103496109f3565b6040516102d1919061411e565b34801561036257600080fd5b506103326103713660046141df565b610c95565b34801561038257600080fd5b506103c561039136600461425e565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016102d1565b3480156103df57600080fd5b506103326103ee36600461428c565b610d27565b3480156103ff57600080fd5b5061033261040e3660046142bc565b610d71565b34801561041f57600080fd5b5061033261042e36600461428c565b610f04565b34801561043f57600080fd5b50610332610f55565b610332610456366004614308565b610f6b565b34801561046757600080fd5b506103c5610f8a565b34801561047c57600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166102c5565b3480156104b357600080fd5b506104c76104c23660046143d5565b610fb9565b6040516102d1929190614421565b3480156104e157600080fd5b506103326104f03660046141df565b6110b4565b34801561050157600080fd5b50610332610510366004614444565b61115a565b34801561052157600080fd5b506105356105303660046143d5565b611219565b6040516102d1919061451f565b34801561054e57600080fd5b5061056261055d36600461425e565b6112e2565b6040516102d1929190614532565b34801561057c57600080fd5b5061033261139a565b34801561059157600080fd5b506103326105a03660046142bc565b6113cc565b3480156105b157600080fd5b506103326105c0366004614554565b611520565b3480156105d157600080fd5b506102c56105e036600461428c565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561063657600080fd5b5061063f6115ae565b6040516102d19190614582565b34801561065857600080fd5b50610661611606565b6040516102d191906145c5565b34801561067a57600080fd5b506103c5600081565b34801561068f57600080fd5b506103c56207a12081565b3480156106a657600080fd5b506103326106b5366004614672565b6117c5565b3480156106c657600080fd5b506102fa6106d53660046143d5565b611845565b3480156106e657600080fd5b506105356040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b34801561072f57600080fd5b506103c57ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa381565b34801561076357600080fd5b506103326107723660046146e6565b611885565b34801561078357600080fd5b5061078c611c32565b6040516102d19190614731565b3480156107a557600080fd5b50600b546102fa906001600160a01b031681565b3480156107c557600080fd5b506105356107d43660046141df565b611f39565b3480156107e557600080fd5b506103326107f436600461428c565b612021565b34801561080557600080fd5b506103c57f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b34801561083957600080fd5b5061084d6108483660046142bc565b612065565b6040516102d196959493929190614828565b34801561086b57600080fd5b5061033261087a366004614886565b6122b7565b34801561088b57600080fd5b506000546102fa906001600160a01b031681565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061093257507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361096281612353565b61096a61235d565b610976858585856123bb565b6109828585858561250f565b7f6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c6007600087815260200190815260200160002085856040516109c6929190614936565b90815260200160405180910390206001016040516109e49190614a1c565b60405180910390a15050505050565b6004546060908067ffffffffffffffff811115610a1257610a126142d9565b604051908082528060200260200182016040528015610a7057816020015b610a5d60405180608001604052806000151581526020016060815260200160608152602001600081525090565b815260200190600190039081610a305790505b50915060005b81811015610c9057600060048281548110610a9357610a93614a2f565b906000526020600020906002020160405180604001604052908160008201548152602001600182018054610ac690614946565b80601f0160208091040260200160405190810160405280929190818152602001828054610af290614946565b8015610b3f5780601f10610b1457610100808354040283529160200191610b3f565b820191906000526020600020905b815481529060010190602001808311610b2257829003601f168201915b505050505081525050905060008160000151905060008260200151905060405180608001604052806007600085815260200190815260200160002083604051610b889190614a5e565b90815260408051602092819003830190205460ff161515835260008681526007835281902090519290910191610bbf908590614a5e565b90815260200160405180910390206001018054610bdb90614946565b80601f0160208091040260200160405190810160405280929190818152602001828054610c0790614946565b8015610c545780601f10610c2957610100808354040283529160200191610c54565b820191906000526020600020905b815481529060010190602001808311610c3757829003601f168201915b5050505050815260200182815260200183815250868581518110610c7a57610c7a614a2f565b6020908102919091010152505050600101610a76565b505090565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3610cbf81612353565b610cc761235d565b610cd48686868686612592565b610ce18686868686612628565b857f40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee63486868686604051610d179493929190614aa5565b60405180910390a2505050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610d6181612353565b610d6b83836126a6565b50505050565b6000610d7c81612353565b6001600160a01b038216610dbc576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610de67ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3836126a6565b50610e117f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a836126a6565b50600154610e49907ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3906001600160a01b0316612775565b50600154610e81907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b0316612775565b50600154604080516001600160a01b03928316815291841660208301527f6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279910160405180910390a150600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6001600160a01b0381163314610f46576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f508282612775565b505050565b6000610f6081612353565b610f6861281b565b50565b610f7361288d565b610f7c8261295d565b610f868282612968565b5050565b6000610f94612a69565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b600083815260076020526040808220905160609190610fdb9086908690614936565b908152604080519182900360209081018320546000898152600790925291902060ff9091169350906110109086908690614936565b9081526020016040518091039020600101805461102c90614946565b80601f016020809104026020016040519081016040528092919081815260200182805461105890614946565b80156110a55780601f1061107a576101008083540402835291602001916110a5565b820191906000526020600020905b81548152906001019060200180831161108857829003601f168201915b50505050509050935093915050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36110de81612353565b6110e661235d565b6110f38686868686612acb565b6111008686868686612daf565b8484604051611110929190614936565b6040518091039020867f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581858560405161114a929190614ad7565b60405180910390a3505050505050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361118481612353565b61118c61235d565b61119d8a8a8a8a8a8a8a8a8a612e2d565b6111ae8a8a8a8a8a8a8a8a8a613163565b896001600160a01b031686866040516111c8929190614936565b60405180910390207fa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367848a8d8d6040516112059493929190614aeb565b60405180910390a350505050505050505050565b6060600660008581526020019081526020016000206004018383604051611241929190614936565b9081526020016040518091039020805461125a90614946565b80601f016020809104026020016040519081016040528092919081815260200182805461128690614946565b80156112d35780601f106112a8576101008083540402835291602001916112d3565b820191906000526020600020905b8154815290600101906020018083116112b657829003601f168201915b505050505090505b9392505050565b60008181526006602052604090206002810154600390910180546001600160a01b03909216916060919061131590614946565b80601f016020809104026020016040519081016040528092919081815260200182805461134190614946565b801561138e5780601f106113635761010080835404028352916020019161138e565b820191906000526020600020905b81548152906001019060200180831161137157829003601f168201915b50505050509050915091565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6113c481612353565b610f686131f5565b60006113d781612353565b6001600160a01b038216611417576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114226000836126a6565b5061144d7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a836126a6565b506000805461146591906001600160a01b0316612775565b5060005461149d907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b0316612775565b50600054604080516001600160a01b03928316815291841660208301527f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f910160405180910390a150600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361154a81612353565b61155261235d565b61155c8383613250565b6115668383613340565b604080516001600160a01b038516815283151560208201527f9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8910160405180910390a1505050565b606060028054806020026020016040519081016040528092919081815260200182805480156115fc57602002820191906000526020600020905b8154815260200190600101908083116115e8575b5050505050905090565b6003546060908067ffffffffffffffff811115611625576116256142d9565b60405190808252806020026020018201604052801561169457816020015b60408051608081018252600080825260208083018290529282015260608082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092019101816116435790505b50915060005b81811015610c90576000600382815481106116b7576116b7614a2f565b6000918252602080832090910154604080516080810182528285526006808552828620805460ff161515835282860185905260028101546001600160a01b031693830193909352948390529390925260039091018054919350606083019161171e90614946565b80601f016020809104026020016040519081016040528092919081815260200182805461174a90614946565b80156117975780601f1061176c57610100808354040283529160200191611797565b820191906000526020600020905b81548152906001019060200180831161177a57829003601f168201915b50505050508152508483815181106117b1576117b1614a2f565b60209081029190910101525060010161169a565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36117ef81612353565b6117f761235d565b61180486868686866133c3565b6118118686868686613579565b857fc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e83604051610d17911515815260200190565b6000838152600a602052604080822090516118639085908590614936565b908152604051908190036020019020546001600160a01b031690509392505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156118d05750825b905060008267ffffffffffffffff1660011480156118ed5750303b155b9050811580156118fb575080155b15611932576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156119935784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b03881615806119b057506001600160a01b038716155b806119c257506001600160a01b038616155b156119f9576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611a016135f7565b611a096135f7565b611a116135ff565b611a1c6000896126a6565b50611a477ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3886126a6565b50611a727f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a886126a6565b50611a9d7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a896126a6565b50600080546001600160a01b038a81167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316178355600180548b8316908416178155600b8054928b16929093169190911790915546808352600660208181526040808620805460ff1916909517855580513060601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016818401528151808203601401815260349091019091529290945290925260030190611b619082614b5f565b50600280546001818101909255467f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9091018190556003805492830181556000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b909101558315611c285784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b6005546060908067ffffffffffffffff811115611c5157611c516142d9565b604051908082528060200260200182016040528015611cd057816020015b611cbd6040518060e0016040528060001515815260200160006001600160a01b0316815260200160608152602001600081526020016060815260200160608152602001600060ff1681525090565b815260200190600190039081611c6f5790505b50915060005b81811015610c9057600060058281548110611cf357611cf3614a2f565b60009182526020808320909101546001600160a01b0390811680845260088352604093849020845160e081018652815460ff811615158252610100900490931693830193909352600183018054919550919384019190611d5290614946565b80601f0160208091040260200160405190810160405280929190818152602001828054611d7e90614946565b8015611dcb5780601f10611da057610100808354040283529160200191611dcb565b820191906000526020600020905b815481529060010190602001808311611dae57829003601f168201915b5050505050815260200160028201548152602001600382018054611dee90614946565b80601f0160208091040260200160405190810160405280929190818152602001828054611e1a90614946565b8015611e675780601f10611e3c57610100808354040283529160200191611e67565b820191906000526020600020905b815481529060010190602001808311611e4a57829003601f168201915b50505050508152602001600482018054611e8090614946565b80601f0160208091040260200160405190810160405280929190818152602001828054611eac90614946565b8015611ef95780601f10611ece57610100808354040283529160200191611ef9565b820191906000526020600020905b815481529060010190602001808311611edc57829003601f168201915b50505091835250506005919091015460ff166020909101528451859084908110611f2557611f25614a2f565b602090810291909101015250600101611cd6565b6060600760008781526020019081526020016000208585604051611f5e929190614936565b90815260200160405180910390206003018383604051611f7f929190614936565b90815260200160405180910390208054611f9890614946565b80601f0160208091040260200160405190810160405280929190818152602001828054611fc490614946565b80156120115780601f10611fe657610100808354040283529160200191612011565b820191906000526020600020905b815481529060010190602001808311611ff457829003601f168201915b5050505050905095945050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461205b81612353565b610d6b8383612775565b6001600160a01b038082166000908152600860209081526040808320815160e081018352815460ff811615158252610100900490951692850192909252600182018054939460609486948694859487948594909392840191906120c790614946565b80601f01602080910402602001604051908101604052809291908181526020018280546120f390614946565b80156121405780601f1061211557610100808354040283529160200191612140565b820191906000526020600020905b81548152906001019060200180831161212357829003601f168201915b505050505081526020016002820154815260200160038201805461216390614946565b80601f016020809104026020016040519081016040528092919081815260200182805461218f90614946565b80156121dc5780601f106121b1576101008083540402835291602001916121dc565b820191906000526020600020905b8154815290600101906020018083116121bf57829003601f168201915b505050505081526020016004820180546121f590614946565b80601f016020809104026020016040519081016040528092919081815260200182805461222190614946565b801561226e5780601f106122435761010080835404028352916020019161226e565b820191906000526020600020905b81548152906001019060200180831161225157829003601f168201915b50505091835250506005919091015460ff16602090910152805160808201516060830151604084015160a085015160c090950151939d929c50909a509850919650945092505050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36122e181612353565b6122e961235d565b6122f888888888888888613632565b6123078888888888888861378a565b877faea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c588888888888860405161234196959493929190614c5a565b60405180910390a25050505050505050565b610f68813361380c565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16156123b9576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60008481526006602052604090205460ff1661240b576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018590526024015b60405180910390fd5b60008290036124495782826040517ec10cfd000000000000000000000000000000000000000000000000000000008152600401612402929190614ad7565b6000848152600760205260409081902090516124689085908590614936565b9081526020016040518091039020600101805461248490614946565b90506000036124c5578383836040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161240293929190614ca3565b806007600086815260200190815260200160002084846040516124e9929190614936565b908152604051908190036020019020805491151560ff1990921691909117905550505050565b6000848484846040516024016125289493929190614cbd565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f10d29b9e00000000000000000000000000000000000000000000000000000000179052905061258b81613899565b5050505050565b60008581526006602052604090205460ff166125dd576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612402565b8181600660008881526020019081526020016000206004018686604051612605929190614936565b90815260200160405180910390209182612620929190614cea565b505050505050565b60008585858585604051602401612643959493929190614de6565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f2259e9e500000000000000000000000000000000000000000000000000000000179052905061262081613899565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1661276b576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556127213390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610932565b6000915050610932565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff161561276b576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610932565b612823613941565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061292657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661291a7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b156123b9576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610f8681612353565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156129c2575060408051601f3d908101601f191682019092526129bf91810190614e1f565b60015b612a03576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401612402565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612a5f576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401612402565b610f50838361399c565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146123b9576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008581526006602052604090205460ff16612b16576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612402565b6000839003612b545783836040517ec10cfd000000000000000000000000000000000000000000000000000000008152600401612402929190614ad7565b6000819003612b8f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000858152600760205260408082209051612bad9087908790614936565b90815260200160405180910390206001018054612bc990614946565b90501115612c0d5784848484846040517f2b192eab000000000000000000000000000000000000000000000000000000008152600401612402959493929190614de6565b6001600760008781526020019081526020016000208585604051612c32929190614936565b9081526040805160209281900383018120805460ff1916941515949094179093556000888152600790925290208391839190612c719088908890614936565b90815260200160405180910390206001019182612c8f929190614cea565b508383600760008881526020019081526020016000208686604051612cb5929190614936565b90815260200160405180910390206002019182612cd3929190614cea565b506004604051806040016040528087815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939094525050835460018181018655948252602091829020845160029092020190815590830151929390929083019150612d539082614b5f565b5050508383604051612d66929190614936565b6040518091039020857f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce25818484604051612da0929190614ad7565b60405180910390a35050505050565b60008585858585604051602401612dca959493929190614de6565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f631d62e400000000000000000000000000000000000000000000000000000000179052905061262081613899565b6001600160a01b038916612e6d576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000879003612ed7576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f53796d626f6c2063616e6e6f7420626520656d707479000000000000000000006044820152606401612402565b6001600160a01b0389811660009081526008602052604090205461010090041615612f39576040517f63f4ee1f0000000000000000000000000000000000000000000000000000000081526001600160a01b038a166004820152602401612402565b60006001600160a01b031660098989604051612f56929190614936565b908152604051908190036020019020546001600160a01b031614612faa5787876040517fb295cac9000000000000000000000000000000000000000000000000000000008152600401612402929190614ad7565b6001600160a01b0389166000818152600860205260409020805460017fffffffffffffffffffffff000000000000000000000000000000000000000000909116610100909302929092178217815501613004858783614cea565b506001600160a01b038916600090815260086020526040902060028101879055600301613032888a83614cea565b506001600160a01b038916600090815260086020526040902060058101805460ff191660ff841617905560040161306a838583614cea565b5088600a6000888152602001908152602001600020868660405161308f929190614936565b908152602001604051809103902060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555088600989896040516130d4929190614936565b90815260405190819003602001902080546001600160a01b039283167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617909155600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180549b9092169a16999099179098555050505050505050565b600089898989898989898960405160240161318699989796959493929190614e38565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6e9e2d3f0000000000000000000000000000000000000000000000000000000017905290506131e981613899565b50505050505050505050565b6131fd61235d565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361286f565b6001600160a01b038216613290576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260086020526040902054610100900416613315576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5a52433230206e6f7420726567697374657265640000000000000000000000006044820152606401612402565b6001600160a01b03919091166000908152600860205260409020805460ff1916911515919091179055565b6040516001600160a01b0383166024820152811515604482015260009060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9060bda9000000000000000000000000000000000000000000000000000000001790529050610f5081613899565b60008581526006602052604090205460ff1680156133de5750805b15613418576040517fa1452cb000000000000000000000000000000000000000000000000000000000815260048101869052602401612402565b60008581526006602052604090205460ff16158015613435575080155b1561346f576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612402565b6000858152600660205260409020600201546001600160a01b03161580156134975750468514155b156134d257600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018590555b6000858152600660205260409020805460ff19168215151781556002810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03871617905560030161352f838583614cea565b50801561357057600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0185905561258b565b61258b856139f2565b60008585858585604051602401613594959493929190614ea3565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa8f2cb9600000000000000000000000000000000000000000000000000000000179052905061262081613899565b6123b9613aa0565b613607613aa0565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b60008781526006602052604090205460ff1661367d576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101889052602401612402565b60008590036136bb5785856040517ec10cfd000000000000000000000000000000000000000000000000000000008152600401612402929190614ad7565b6000878152600760205260409081902090516136da9088908890614936565b9081526040519081900360200190205460ff16613729578686866040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161240293929190614ca3565b8181600760008a8152602001908152602001600020888860405161374e929190614936565b9081526020016040518091039020600301868660405161376f929190614936565b90815260200160405180910390209182611c28929190614cea565b6000878787878787876040516024016137a99796959493929190614ee0565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff354b31f000000000000000000000000000000000000000000000000000000001790529050611c2881613899565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610f86576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401612402565b60005b600254811015610f865746600282815481106138ba576138ba614a2f565b9060005260206000200154148061390e575060066000600283815481106138e3576138e3614a2f565b90600052602060002001548152602001908152602001600020600301805461390a90614946565b1590505b613939576139396002828154811061392857613928614a2f565b906000526020600020015483613b07565b60010161389c565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166123b9576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6139a582613dc9565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156139ea57610f508282613e71565b610f86613ee7565b60005b600254811015610f86578160028281548110613a1357613a13614a2f565b906000526020600020015403613a985760028054613a3390600190614f30565b81548110613a4357613a43614a2f565b906000526020600020015460028281548110613a6157613a61614a2f565b6000918252602090912001556002805480613a7e57613a7e614f6a565b600190038181906000526020600020016000905590555050565b6001016139f5565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff166123b9576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526207a120815260006020808301829052835160a0810185528281529081018290529283018190526060808401526080830152906000848152600660205260408082206002015490517ffc5fecd50000000000000000000000000000000000000000000000000000000081526207a12060048201526001600160a01b039091169190829063fc5fecd5906024016040805180830381865afa158015613bb8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613bdc9190614f99565b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018290529092506001600160a01b03841691506323b872dd906064016020604051808303816000875af1158015613c4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c709190614fc7565b613ca6576040517f90b8ec1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303816000875af1158015613d12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d369190614fc7565b50600b546000878152600660205260409081902090517f06cb89830000000000000000000000000000000000000000000000000000000081526001600160a01b03909216916306cb898391613d9b9160039091019086908a908a908a90600401614fe4565b600060405180830381600087803b158015613db557600080fd5b505af11580156131e9573d6000803e3d6000fd5b806001600160a01b03163b600003613e18576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401612402565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051613e8e9190614a5e565b600060405180830381855af49150503d8060008114613ec9576040519150601f19603f3d011682016040523d82523d6000602084013e613ece565b606091505b5091509150613ede858383613f1f565b95945050505050565b34156123b9576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082613f3457613f2f82613f94565b6112db565b8151158015613f4b57506001600160a01b0384163b155b15613f8d576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401612402565b50806112db565b805115613fa45780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215613fe857600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146112db57600080fd5b60008083601f84011261402a57600080fd5b50813567ffffffffffffffff81111561404257600080fd5b60208301915083602082850101111561405a57600080fd5b9250929050565b8015158114610f6857600080fd5b6000806000806060858703121561408557600080fd5b84359350602085013567ffffffffffffffff8111156140a357600080fd5b6140af87828801614018565b90945092505060408501356140c381614061565b939692955090935050565b60005b838110156140e95781810151838201526020016140d1565b50506000910152565b6000815180845261410a8160208601602086016140ce565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156141d3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180511515865260208101516080602088015261419660808801826140f2565b9050604082015187820360408901526141af82826140f2565b60609384015198909301979097525094506020938401939190910190600101614146565b50929695505050505050565b6000806000806000606086880312156141f757600080fd5b85359450602086013567ffffffffffffffff81111561421557600080fd5b61422188828901614018565b909550935050604086013567ffffffffffffffff81111561424157600080fd5b61424d88828901614018565b969995985093965092949392505050565b60006020828403121561427057600080fd5b5035919050565b6001600160a01b0381168114610f6857600080fd5b6000806040838503121561429f57600080fd5b8235915060208301356142b181614277565b809150509250929050565b6000602082840312156142ce57600080fd5b81356112db81614277565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806040838503121561431b57600080fd5b823561432681614277565b9150602083013567ffffffffffffffff81111561434257600080fd5b8301601f8101851361435357600080fd5b803567ffffffffffffffff81111561436d5761436d6142d9565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff8211171561439d5761439d6142d9565b6040528181528282016020018710156143b557600080fd5b816020840160208301376000602083830101528093505050509250929050565b6000806000604084860312156143ea57600080fd5b83359250602084013567ffffffffffffffff81111561440857600080fd5b61441486828701614018565b9497909650939450505050565b821515815260406020820152600061443c60408301846140f2565b949350505050565b600080600080600080600080600060c08a8c03121561446257600080fd5b893561446d81614277565b985060208a013567ffffffffffffffff81111561448957600080fd5b6144958c828d01614018565b90995097505060408a0135955060608a013567ffffffffffffffff8111156144bc57600080fd5b6144c88c828d01614018565b90965094505060808a013567ffffffffffffffff8111156144e857600080fd5b6144f48c828d01614018565b90945092505060a08a013560ff8116811461450e57600080fd5b809150509295985092959850929598565b6020815260006112db60208301846140f2565b6001600160a01b038316815260406020820152600061443c60408301846140f2565b6000806040838503121561456757600080fd5b823561457281614277565b915060208301356142b181614061565b602080825282518282018190526000918401906040840190835b818110156145ba57835183526020938401939092019160010161459c565b509095945050505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156141d3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160208701526001600160a01b036040820151166040870152606081015190506080606087015261465c60808701826140f2565b95505060209384019391909101906001016145ed565b60008060008060006080868803121561468a57600080fd5b85359450602086013561469c81614277565b9350604086013567ffffffffffffffff8111156146b857600080fd5b6146c488828901614018565b90945092505060608601356146d881614061565b809150509295509295909350565b6000806000606084860312156146fb57600080fd5b833561470681614277565b9250602084013561471681614277565b9150604084013561472681614277565b809150509250925092565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156141d3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051151586526001600160a01b036020820151166020870152604081015160e060408801526147bc60e08801826140f2565b905060608201516060880152608082015187820360808901526147df82826140f2565b91505060a082015187820360a08901526147f982826140f2565b91505060c0820151915061481260c088018360ff169052565b9550506020938401939190910190600101614759565b861515815260c06020820152600061484360c08301886140f2565b866040840152828103606084015261485b81876140f2565b9050828103608084015261486f81866140f2565b91505060ff831660a0830152979650505050505050565b60008060008060008060006080888a0312156148a157600080fd5b87359650602088013567ffffffffffffffff8111156148bf57600080fd5b6148cb8a828b01614018565b909750955050604088013567ffffffffffffffff8111156148eb57600080fd5b6148f78a828b01614018565b909550935050606088013567ffffffffffffffff81111561491757600080fd5b6149238a828b01614018565b989b979a50959850939692959293505050565b8183823760009101908152919050565b600181811c9082168061495a57607f821691505b602082108103614993577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b600081546149a681614946565b8085526001821680156149c057600181146149dc57614a13565b60ff1983166020870152602082151560051b8701019350614a13565b84600052602060002060005b83811015614a0a5781546020828a0101526001820191506020810190506149e8565b87016020019450505b50505092915050565b6020815260006112db6020830184614999565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008251614a708184602087016140ce565b9190910192915050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b604081526000614ab9604083018688614a7a565b8281036020840152614acc818587614a7a565b979650505050505050565b60208152600061443c602083018486614a7a565b60ff85168152836020820152606060408201526000614b0e606083018486614a7a565b9695505050505050565b601f821115610f5057806000526020600020601f840160051c81016020851015614b3f5750805b601f840160051c820191505b8181101561258b5760008155600101614b4b565b815167ffffffffffffffff811115614b7957614b796142d9565b614b8d81614b878454614946565b84614b18565b6020601f821160018114614bdf5760008315614ba95750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b17845561258b565b600084815260208120601f198516915b82811015614c0f5787850151825560209485019460019092019101614bef565b5084821015614c4b57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b606081526000614c6e60608301888a614a7a565b8281036020840152614c81818789614a7a565b90508281036040840152614c96818587614a7a565b9998505050505050505050565b838152604060208201526000613ede604083018486614a7a565b848152606060208201526000614cd7606083018587614a7a565b9050821515604083015295945050505050565b67ffffffffffffffff831115614d0257614d026142d9565b614d1683614d108354614946565b83614b18565b6000601f841160018114614d685760008515614d325750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b17835561258b565b600083815260209020601f19861690835b82811015614d995786850135825560209485019460019092019101614d79565b5086821015614dd4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b858152606060208201526000614e00606083018688614a7a565b8281036040840152614e13818587614a7a565b98975050505050505050565b600060208284031215614e3157600080fd5b5051919050565b6001600160a01b038a16815260c060208201526000614e5b60c083018a8c614a7a565b8860408401528281036060840152614e7481888a614a7a565b90508281036080840152614e89818688614a7a565b91505060ff831660a08301529a9950505050505050505050565b8581526001600160a01b0385166020820152608060408201526000614ecc608083018587614a7a565b905082151560608301529695505050505050565b878152608060208201526000614efa60808301888a614a7a565b8281036040840152614f0d818789614a7a565b90508281036060840152614f22818587614a7a565b9a9950505050505050505050565b81810381811115610932577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008060408385031215614fac57600080fd5b8251614fb781614277565b6020939093015192949293505050565b600060208284031215614fd957600080fd5b81516112db81614061565b60c081526000614ff760c0830188614999565b6001600160a01b0387166020840152828103604084015261501881876140f2565b90508451606084015260208501511515608084015282810360a08401526001600160a01b0384511681526020840151151560208201526001600160a01b036040850151166040820152606084015160a0606083015261507a60a08301826140f2565b9050608085015160808301528092505050969550505050505056fea2646970667358221220185320aaa2ff1b87c3db2fc5c5c150ce1b0fb71a2ace4fd7700966ba6b36ecda64736f6c634300081a003360a060405234801561001057600080fd5b50737cce3eb018bf23e1fe2a32692f2c77592d1103946001600160a01b031663cc5ad8b66040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610065573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610089919061009a565b6001600160a01b03166080526100ca565b6000602082840312156100ac57600080fd5b81516001600160a01b03811681146100c357600080fd5b9392505050565b60805161095e6100e46000396000607e015261095e6000f3fe6080604052600436106100635760003560e01c80637b103999116100405780637b10399914610109578063c9028a3614610131578063ebf9b2aa1461014457005b8063116191b61461006c5780632d4cfb7e146100c95780635bcfd616146100e957005b3661006a57005b005b34801561007857600080fd5b506100a07f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100d557600080fd5b5061006a6100e43660046102b4565b610157565b3480156100f557600080fd5b5061006a610104366004610380565b610191565b34801561011557600080fd5b506100a0737cce3eb018bf23e1fe2a32692f2c77592d11039481565b61006a61013f36600461040a565b61020c565b61006a610152366004610445565b61023b565b7f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7816040516101869190610560565b60405180910390a150565b606081156101a8576101a582840184610666565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e6101d3878061075c565b6101e360408a0160208b016107c1565b896040013533866040516101fc969594939291906107dc565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c481604051610186919061089e565b606081156102525761024f82840184610666565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e61027d858061075c565b61028d60408801602089016107c1565b876040013533866040516102a6969594939291906107dc565b60405180910390a150505050565b6000602082840312156102c657600080fd5b813567ffffffffffffffff8111156102dd57600080fd5b820160c081850312156102ef57600080fd5b9392505050565b60006060828403121561030857600080fd5b50919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461033257600080fd5b919050565b60008083601f84011261034957600080fd5b50813567ffffffffffffffff81111561036157600080fd5b60208301915083602082850101111561037957600080fd5b9250929050565b60008060008060006080868803121561039857600080fd5b853567ffffffffffffffff8111156103af57600080fd5b6103bb888289016102f6565b9550506103ca6020870161030e565b935060408601359250606086013567ffffffffffffffff8111156103ed57600080fd5b6103f988828901610337565b969995985093965092949392505050565b60006020828403121561041c57600080fd5b813567ffffffffffffffff81111561043357600080fd5b8201608081850312156102ef57600080fd5b60008060006040848603121561045a57600080fd5b833567ffffffffffffffff81111561047157600080fd5b61047d868287016102f6565b935050602084013567ffffffffffffffff81111561049a57600080fd5b6104a686828701610337565b9497909650939450505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126104e857600080fd5b830160208101925035905067ffffffffffffffff81111561050857600080fd5b80360382131561037957600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60208152600061057083846104b3565b60c0602085015261058560e085018284610517565b91505073ffffffffffffffffffffffffffffffffffffffff6105a96020860161030e565b1660408401526000604085013590508060608501525060608401358015158082146105d357600080fd5b80608086015250506000608085013590508060a0850152506105f860a08501856104b3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c086015261062d838284610517565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561067857600080fd5b813567ffffffffffffffff81111561068f57600080fd5b8201601f810184136106a057600080fd5b803567ffffffffffffffff8111156106ba576106ba610637565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561072657610726610637565b60405281815282820160200186101561073e57600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261079157600080fd5b83018035915067ffffffffffffffff8211156107ac57600080fd5b60200191503681900382131561037957600080fd5b6000602082840312156107d357600080fd5b6102ef8261030e565b60a0815260006107f060a08301888a610517565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b8181101561085a5760208187018101518483018201520161083e565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff6108c08361030e565b16602082015273ffffffffffffffffffffffffffffffffffffffff6108e76020840161030e565b166040820152600080604084013590508060608401525061090b60608401846104b3565b60808085015261091f60a085018284610517565b9594505050505056fea26469706673582212206f7565d8db8f7480b60517069080a5ccd0f8b46abb36bd8b344abd6930ec4d0664736f6c634300081a003360c060405234801561001057600080fd5b5060405161106f38038061106f83398101604081905261002f916100db565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006357604051632b2add3d60e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0385811691909117909155828116608052811660a0526040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac590600090a150505061011e565b80516001600160a01b03811681146100d657600080fd5b919050565b6000806000606084860312156100f057600080fd5b6100f9846100bf565b9250610107602085016100bf565b9150610115604085016100bf565b90509250925092565b60805160a051610f2561014a60003960006101e50152600081816102b9015261045b0152610f256000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806397770dff11610097578063c63585cc11610066578063c63585cc14610273578063d7fd7afb14610286578063d936a012146102b4578063ee2815ba146102db57600080fd5b806397770dff1461021a578063a7cb05071461022d578063c39aca3714610240578063c62178ac1461025357600080fd5b8063513a9c05116100d3578063513a9c051461018a578063569541b9146101c0578063842da36d146101e057806391dd645f1461020757600080fd5b80630be15547146100fa5780631f0e251b1461015a5780633ce4a5bc1461016f575b600080fd5b610130610108366004610bd1565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61016d610168366004610c13565b6102ee565b005b61013073735b14bb79463307aacbed86daf3322b1e6226ab81565b610130610198366004610bd1565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101309073ffffffffffffffffffffffffffffffffffffffff1681565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d610215366004610c35565b610402565b61016d610228366004610c13565b610526565b61016d61023b366004610c61565b610633565b61016d61024e366004610c83565b6106ce565b6004546101309073ffffffffffffffffffffffffffffffffffffffff1681565b610130610281366004610d53565b6108cd565b6102a6610294366004610bd1565b60006020819052908152604090205481565b604051908152602001610151565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d6102e9366004610c35565b610a02565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461033b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610388576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461044f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354600090610497907f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff16846108cd565b60008481526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251878152918201529192507f0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e910160405180910390a1505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610573576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166105c0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e906020016103f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610680576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461071b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831673735b14bb79463307aacbed86daf3322b1e6226ab1480610768575073ffffffffffffffffffffffffffffffffffffffff831630145b1561079f576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690528616906347e7ef24906044016020604051808303816000875af1158015610814573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108389190610d96565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063de43156e906108939089908990899088908890600401610e01565b600060405180830381600087803b1580156108ad57600080fd5b505af11580156108c1573d6000803e3d6000fd5b50505050505050505050565b60008060006108dc8585610ad3565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b16603482015291935091508690604801604051602081830303815290604052805190602001206040516020016109c29291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a4f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d91016106c2565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b3b576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610b75578284610b78565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610bca576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b600060208284031215610be357600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c0e57600080fd5b919050565b600060208284031215610c2557600080fd5b610c2e82610bea565b9392505050565b60008060408385031215610c4857600080fd5b82359150610c5860208401610bea565b90509250929050565b60008060408385031215610c7457600080fd5b50508035926020909101359150565b60008060008060008060a08789031215610c9c57600080fd5b863567ffffffffffffffff811115610cb357600080fd5b87016060818a031215610cc557600080fd5b9550610cd360208801610bea565b945060408701359350610ce860608801610bea565b9250608087013567ffffffffffffffff811115610d0457600080fd5b8701601f81018913610d1557600080fd5b803567ffffffffffffffff811115610d2c57600080fd5b896020828401011115610d3e57600080fd5b60208201935080925050509295509295509295565b600080600060608486031215610d6857600080fd5b610d7184610bea565b9250610d7f60208501610bea565b9150610d8d60408501610bea565b90509250925092565b600060208284031215610da857600080fd5b81518015158114610c2e57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60808152600086357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018112610e3957600080fd5b870160208101903567ffffffffffffffff811115610e5657600080fd5b803603821315610e6557600080fd5b60606080850152610e7a60e085018284610db8565b91505073ffffffffffffffffffffffffffffffffffffffff610e9e60208a01610bea565b1660a0840152604088013560c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401528281036060840152610ee3818587610db8565b9897505050505050505056fea26469706673582212204635fa482dcaad895b3b4ec621fd47a9c4c1d3aa3f8abdb3203f0116e91c4ab964736f6c634300081a003360c060405234801561001057600080fd5b5060405161207238038061207283398101604081905261002f916101f0565b6001600160a01b038216158061004c57506001600160a01b038116155b1561006a5760405163d92e233d60e01b815260040160405180910390fd5b60066100768982610342565b5060076100838882610342565b506008805460ff191660ff881617905560808590528360028111156100aa576100aa610400565b60a08160028111156100be576100be610400565b905250600192909255600080546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506104169350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261013357600080fd5b81516001600160401b0381111561014c5761014c61010c565b604051601f8201601f19908116603f011681016001600160401b038111828210171561017a5761017a61010c565b60405281815283820160200185101561019257600080fd5b60005b828110156101b157602081860181015183830182015201610195565b506000918101602001919091529392505050565b8051600381106101d457600080fd5b919050565b80516001600160a01b03811681146101d457600080fd5b600080600080600080600080610100898b03121561020d57600080fd5b88516001600160401b0381111561022357600080fd5b61022f8b828c01610122565b60208b015190995090506001600160401b0381111561024d57600080fd5b6102598b828c01610122565b975050604089015160ff8116811461027057600080fd5b60608a0151909650945061028660808a016101c5565b60a08a0151909450925061029c60c08a016101d9565b91506102aa60e08a016101d9565b90509295985092959890939650565b600181811c908216806102cd57607f821691505b6020821081036102ed57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561033d57806000526020600020601f840160051c8101602085101561031a5750805b601f840160051c820191505b8181101561033a5760008155600101610326565b50505b505050565b81516001600160401b0381111561035b5761035b61010c565b61036f8161036984546102b9565b846102f3565b6020601f8211600181146103a3576000831561038b5750848201515b600019600385901b1c1916600184901b17845561033a565b600084815260208120601f198516915b828110156103d357878501518255602094850194600190920191016103b3565b50848210156103f15786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a051611c1b61045760003960006103440152600081816102f001528181610bdc01528181610ce201528181610efe01526110040152611c1b6000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806395d89b41116100f9578063ccc7759911610097578063eddeb12311610071578063eddeb12314610461578063f2441b3214610474578063f687d12a14610494578063fc5fecd5146104a757600080fd5b8063ccc77599146103d4578063d9eeebed146103e7578063dd62ed3e1461041b57600080fd5b8063b84c8246116100d3578063b84c824614610386578063c47f00271461039b578063c7012626146103ae578063c835d7cc146103c157600080fd5b806395d89b4114610337578063a3413d031461033f578063a9059cbb1461037357600080fd5b80633ce4a5bc116101665780634d8943bb116101405780634d8943bb146102ac57806370a08231146102b557806385e1f4d0146102eb5780638b851b951461031257600080fd5b80633ce4a5bc1461024657806342966c681461028657806347e7ef241461029957600080fd5b806318160ddd1161019757806318160ddd1461021657806323b872dd1461021e578063313ce5671461023157600080fd5b806306fdde03146101be578063091d2788146101dc578063095ea7b3146101f3575b600080fd5b6101c66104ba565b6040516101d39190611648565b60405180910390f35b6101e560015481565b6040519081526020016101d3565b610206610201366004611687565b61054c565b60405190151581526020016101d3565b6005546101e5565b61020661022c3660046116b3565b610563565b60085460405160ff90911681526020016101d3565b61026173735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6102066102943660046116f4565b6105fa565b6102066102a7366004611687565b61060e565b6101e560025481565b6101e56102c336600461170d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101e57f000000000000000000000000000000000000000000000000000000000000000081565b60085461026190610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101c6610767565b6103667f000000000000000000000000000000000000000000000000000000000000000081565b6040516101d3919061172a565b610206610381366004611687565b610776565b610399610394366004611832565b610783565b005b6103996103a9366004611832565b6107e0565b6102066103bc366004611883565b610839565b6103996103cf36600461170d565b610988565b6103996103e236600461170d565b610a9c565b6103ef610bb0565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101d3565b6101e56104293660046118dc565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b61039961046f3660046116f4565b610dce565b6000546102619073ffffffffffffffffffffffffffffffffffffffff1681565b6103996104a23660046116f4565b610e50565b6103ef6104b53660046116f4565b610ed2565b6060600680546104c990611915565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590611915565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b5050505050905090565b60006105593384846110ee565b5060015b92915050565b60006105708484846111f7565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054828110156105db576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ef85336105ea8685611997565b6110ee565b506001949350505050565b600061060633836113b2565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab1480159061064c575060005473ffffffffffffffffffffffffffffffffffffffff163314155b80156106755750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b156106ac576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106b683836114f4565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107569186906119aa565b60405180910390a250600192915050565b6060600780546104c990611915565b60006105593384846111f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107d0576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107dc8282611a1b565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461082d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107dc8282611a1b565b6000806000610846610bb0565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156108d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fc9190611b34565b610932576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61093c33856113b2565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161097591899189918791611b56565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109d5576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a22576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610ae9576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b36576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c679190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610cb6576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d699190611ba2565b905080600003610da5576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610db89190611bbb565b610dc29190611bd2565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e1b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a91565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e9d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f899190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610fd8576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015611067573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108b9190611ba2565b9050806000036110c7576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000906110d78784611bbb565b6110e19190611bd2565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661113b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611188576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611244576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611291576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040902054818110156112f1576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112fb8282611997565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260036020526040808220939093559085168152908120805484929061133e908490611bd2565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113a491815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113ff576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260409020548181101561145f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114698282611997565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812091909155600580548492906114a4908490611997565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111ea565b73ffffffffffffffffffffffffffffffffffffffff8216611541576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008282546115539190611bd2565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805483929061158d908490611bd2565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561160a576020818501810151868301820152016115ee565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061165b60208301846115e4565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461168457600080fd5b50565b6000806040838503121561169a57600080fd5b82356116a581611662565b946020939093013593505050565b6000806000606084860312156116c857600080fd5b83356116d381611662565b925060208401356116e381611662565b929592945050506040919091013590565b60006020828403121561170657600080fd5b5035919050565b60006020828403121561171f57600080fd5b813561165b81611662565b6020810160038310611765577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff8411156117b5576117b561176b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156118025761180261176b565b60405283815290508082840185101561181a57600080fd5b83836020830137600060208583010152509392505050565b60006020828403121561184457600080fd5b813567ffffffffffffffff81111561185b57600080fd5b8201601f8101841361186c57600080fd5b61187b8482356020840161179a565b949350505050565b6000806040838503121561189657600080fd5b823567ffffffffffffffff8111156118ad57600080fd5b8301601f810185136118be57600080fd5b6118cd8582356020840161179a565b95602094909401359450505050565b600080604083850312156118ef57600080fd5b82356118fa81611662565b9150602083013561190a81611662565b809150509250929050565b600181811c9082168061192957607f821691505b602082108103611962577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561055d5761055d611968565b6040815260006119bd60408301856115e4565b90508260208301529392505050565b601f821115611a1657806000526020600020601f840160051c810160208510156119f35750805b601f840160051c820191505b81811015611a1357600081556001016119ff565b50505b505050565b815167ffffffffffffffff811115611a3557611a3561176b565b611a4981611a438454611915565b846119cc565b6020601f821160018114611a9b5760008315611a655750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611a13565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611ae95787850151825560209485019460019092019101611ac9565b5084821015611b2557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b600060208284031215611b4657600080fd5b8151801515811461165b57600080fd5b608081526000611b6960808301876115e4565b6020830195909552506040810192909252606090910152919050565b600060208284031215611b9757600080fd5b815161165b81611662565b600060208284031215611bb457600080fd5b5051919050565b808202811582820484141761055d5761055d611968565b8082018082111561055d5761055d61196856fea26469706673582212205c01f2682d8e2240fabc45c195d1ebe0c11d19758f25f18c2d70d3c6c00973bc64736f6c634300081a0033a264697066735822122031091a7f80fc59f4e94675bdfdf2c773f9a7c3947b8679fad1fad2dbaabdf35264736f6c634300081a0033",
}

// GatewayZEVMInboundTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayZEVMInboundTestMetaData.ABI instead.
var GatewayZEVMInboundTestABI = GatewayZEVMInboundTestMetaData.ABI

// GatewayZEVMInboundTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayZEVMInboundTestMetaData.Bin instead.
var GatewayZEVMInboundTestBin = GatewayZEVMInboundTestMetaData.Bin

// DeployGatewayZEVMInboundTest deploys a new Ethereum contract, binding an instance of GatewayZEVMInboundTest to it.
func DeployGatewayZEVMInboundTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayZEVMInboundTest, error) {
	parsed, err := GatewayZEVMInboundTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayZEVMInboundTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayZEVMInboundTest{GatewayZEVMInboundTestCaller: GatewayZEVMInboundTestCaller{contract: contract}, GatewayZEVMInboundTestTransactor: GatewayZEVMInboundTestTransactor{contract: contract}, GatewayZEVMInboundTestFilterer: GatewayZEVMInboundTestFilterer{contract: contract}}, nil
}

// GatewayZEVMInboundTest is an auto generated Go binding around an Ethereum contract.
type GatewayZEVMInboundTest struct {
	GatewayZEVMInboundTestCaller     // Read-only binding to the contract
	GatewayZEVMInboundTestTransactor // Write-only binding to the contract
	GatewayZEVMInboundTestFilterer   // Log filterer for contract events
}

// GatewayZEVMInboundTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayZEVMInboundTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMInboundTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayZEVMInboundTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMInboundTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayZEVMInboundTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMInboundTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayZEVMInboundTestSession struct {
	Contract     *GatewayZEVMInboundTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GatewayZEVMInboundTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayZEVMInboundTestCallerSession struct {
	Contract *GatewayZEVMInboundTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// GatewayZEVMInboundTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayZEVMInboundTestTransactorSession struct {
	Contract     *GatewayZEVMInboundTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// GatewayZEVMInboundTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayZEVMInboundTestRaw struct {
	Contract *GatewayZEVMInboundTest // Generic contract binding to access the raw methods on
}

// GatewayZEVMInboundTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayZEVMInboundTestCallerRaw struct {
	Contract *GatewayZEVMInboundTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayZEVMInboundTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayZEVMInboundTestTransactorRaw struct {
	Contract *GatewayZEVMInboundTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayZEVMInboundTest creates a new instance of GatewayZEVMInboundTest, bound to a specific deployed contract.
func NewGatewayZEVMInboundTest(address common.Address, backend bind.ContractBackend) (*GatewayZEVMInboundTest, error) {
	contract, err := bindGatewayZEVMInboundTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTest{GatewayZEVMInboundTestCaller: GatewayZEVMInboundTestCaller{contract: contract}, GatewayZEVMInboundTestTransactor: GatewayZEVMInboundTestTransactor{contract: contract}, GatewayZEVMInboundTestFilterer: GatewayZEVMInboundTestFilterer{contract: contract}}, nil
}

// NewGatewayZEVMInboundTestCaller creates a new read-only instance of GatewayZEVMInboundTest, bound to a specific deployed contract.
func NewGatewayZEVMInboundTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayZEVMInboundTestCaller, error) {
	contract, err := bindGatewayZEVMInboundTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestCaller{contract: contract}, nil
}

// NewGatewayZEVMInboundTestTransactor creates a new write-only instance of GatewayZEVMInboundTest, bound to a specific deployed contract.
func NewGatewayZEVMInboundTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayZEVMInboundTestTransactor, error) {
	contract, err := bindGatewayZEVMInboundTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestTransactor{contract: contract}, nil
}

// NewGatewayZEVMInboundTestFilterer creates a new log filterer instance of GatewayZEVMInboundTest, bound to a specific deployed contract.
func NewGatewayZEVMInboundTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayZEVMInboundTestFilterer, error) {
	contract, err := bindGatewayZEVMInboundTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestFilterer{contract: contract}, nil
}

// bindGatewayZEVMInboundTest binds a generic wrapper to an already deployed contract.
func bindGatewayZEVMInboundTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayZEVMInboundTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMInboundTest.Contract.GatewayZEVMInboundTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.GatewayZEVMInboundTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.GatewayZEVMInboundTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMInboundTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) ISTEST() (bool, error) {
	return _GatewayZEVMInboundTest.Contract.ISTEST(&_GatewayZEVMInboundTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) ISTEST() (bool, error) {
	return _GatewayZEVMInboundTest.Contract.ISTEST(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeArtifacts(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeArtifacts(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeContracts(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeContracts(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeSelectors(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeSelectors(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeSenders(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeSenders(&_GatewayZEVMInboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) Failed() (bool, error) {
	return _GatewayZEVMInboundTest.Contract.Failed(&_GatewayZEVMInboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) Failed() (bool, error) {
	return _GatewayZEVMInboundTest.Contract.Failed(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayZEVMInboundTest.Contract.TargetArtifactSelectors(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayZEVMInboundTest.Contract.TargetArtifactSelectors(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayZEVMInboundTest.Contract.TargetArtifacts(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayZEVMInboundTest.Contract.TargetArtifacts(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.TargetContracts(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.TargetContracts(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayZEVMInboundTest.Contract.TargetInterfaces(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayZEVMInboundTest.Contract.TargetInterfaces(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMInboundTest.Contract.TargetSelectors(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMInboundTest.Contract.TargetSelectors(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.TargetSenders(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.TargetSenders(&_GatewayZEVMInboundTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.SetUp(&_GatewayZEVMInboundTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.SetUp(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallFailsIfRevertGasLimitExceeded")
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOpts is a paid mutator transaction binding the contract method 0x6198fb19.
//
// Solidity: function testCallWithCallOpts() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallWithCallOpts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallWithCallOpts")
}

// TestCallWithCallOpts is a paid mutator transaction binding the contract method 0x6198fb19.
//
// Solidity: function testCallWithCallOpts() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallWithCallOpts() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOpts(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOpts is a paid mutator transaction binding the contract method 0x6198fb19.
//
// Solidity: function testCallWithCallOpts() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallWithCallOpts() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOpts(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0x05b9f046.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsBelowMin() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallWithCallOptsFailsIfGasLimitIsBelowMin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfGasLimitIsBelowMin")
}

// TestCallWithCallOptsFailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0x05b9f046.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsBelowMin() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallWithCallOptsFailsIfGasLimitIsBelowMin() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitIsBelowMin(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0x05b9f046.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsBelowMin() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfGasLimitIsBelowMin() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitIsBelowMin(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x5efe72a9.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallWithCallOptsFailsIfGasLimitIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfGasLimitIsZero")
}

// TestCallWithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x5efe72a9.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallWithCallOptsFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x5efe72a9.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x6dfcbc50.
//
// Solidity: function testCallWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallWithCallOptsFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfMessageSizeExceeded")
}

// TestCallWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x6dfcbc50.
//
// Solidity: function testCallWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallWithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x6dfcbc50.
//
// Solidity: function testCallWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xb51ac071.
//
// Solidity: function testCallWithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallWithCallOptsFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfReceiverIsZeroAddress")
}

// TestCallWithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xb51ac071.
//
// Solidity: function testCallWithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallWithCallOptsFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xb51ac071.
//
// Solidity: function testCallWithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestUpgradeAndWithdrawZRC20 is a paid mutator transaction binding the contract method 0x20dee15f.
//
// Solidity: function testUpgradeAndWithdrawZRC20() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestUpgradeAndWithdrawZRC20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testUpgradeAndWithdrawZRC20")
}

// TestUpgradeAndWithdrawZRC20 is a paid mutator transaction binding the contract method 0x20dee15f.
//
// Solidity: function testUpgradeAndWithdrawZRC20() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestUpgradeAndWithdrawZRC20() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestUpgradeAndWithdrawZRC20(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestUpgradeAndWithdrawZRC20 is a paid mutator transaction binding the contract method 0x20dee15f.
//
// Solidity: function testUpgradeAndWithdrawZRC20() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestUpgradeAndWithdrawZRC20() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestUpgradeAndWithdrawZRC20(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x91dc32e7.
//
// Solidity: function testWithdrawAndCallZETAFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZETAFailsIfRevertGasLimitExceeded")
}

// TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x91dc32e7.
//
// Solidity: function testWithdrawAndCallZETAFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x91dc32e7.
//
// Solidity: function testWithdrawAndCallZETAFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x43181437.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress")
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x43181437.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x43181437.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x42752d41.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero")
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x42752d41.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x42752d41.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xc946d7c0.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded")
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xc946d7c0.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xc946d7c0.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xba9adeef.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfAmountIsZero")
}

// TestWithdrawAndCallZRC20FailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xba9adeef.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xba9adeef.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0xc20049f4.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin")
}

// TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0xc20049f4.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0xc20049f4.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0xbed3e813.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfGasLimitIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfGasLimitIsZero")
}

// TestWithdrawAndCallZRC20FailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0xbed3e813.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0xbed3e813.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x4ffab9de.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfMessageSizeExceeded")
}

// TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x4ffab9de.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x4ffab9de.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x36431b3f.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress")
}

// TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x36431b3f.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x36431b3f.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xecf26b30.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded")
}

// TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xecf26b30.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xecf26b30.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x1b9641bf.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero")
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x1b9641bf.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x1b9641bf.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x0b5ad28d.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero")
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x0b5ad28d.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x0b5ad28d.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xa721b2d3.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded")
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xa721b2d3.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xa721b2d3.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xeb7a2fac.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress")
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xeb7a2fac.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xeb7a2fac.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETA is a paid mutator transaction binding the contract method 0xea37902f.
//
// Solidity: function testWithdrawZETA() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETA")
}

// TestWithdrawZETA is a paid mutator transaction binding the contract method 0xea37902f.
//
// Solidity: function testWithdrawZETA() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETA() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETA(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETA is a paid mutator transaction binding the contract method 0xea37902f.
//
// Solidity: function testWithdrawZETA() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETA() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETA(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x21aeb18c.
//
// Solidity: function testWithdrawZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAFailsIfAmountIsZero")
}

// TestWithdrawZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x21aeb18c.
//
// Solidity: function testWithdrawZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x21aeb18c.
//
// Solidity: function testWithdrawZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x7ae69730.
//
// Solidity: function testWithdrawZETAFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAFailsIfMessageSizeExceeded")
}

// TestWithdrawZETAFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x7ae69730.
//
// Solidity: function testWithdrawZETAFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x7ae69730.
//
// Solidity: function testWithdrawZETAFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xdde7e967.
//
// Solidity: function testWithdrawZETAFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAFailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAFailsIfNoAllowance")
}

// TestWithdrawZETAFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xdde7e967.
//
// Solidity: function testWithdrawZETAFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xdde7e967.
//
// Solidity: function testWithdrawZETAFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfNoBalance is a paid mutator transaction binding the contract method 0x6221b509.
//
// Solidity: function testWithdrawZETAFailsIfNoBalance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAFailsIfNoBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAFailsIfNoBalance")
}

// TestWithdrawZETAFailsIfNoBalance is a paid mutator transaction binding the contract method 0x6221b509.
//
// Solidity: function testWithdrawZETAFailsIfNoBalance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAFailsIfNoBalance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfNoBalance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfNoBalance is a paid mutator transaction binding the contract method 0x6221b509.
//
// Solidity: function testWithdrawZETAFailsIfNoBalance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAFailsIfNoBalance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfNoBalance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xdc749dd7.
//
// Solidity: function testWithdrawZETAFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAFailsIfReceiverIsZeroAddress")
}

// TestWithdrawZETAFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xdc749dd7.
//
// Solidity: function testWithdrawZETAFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xdc749dd7.
//
// Solidity: function testWithdrawZETAFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xd597a27a.
//
// Solidity: function testWithdrawZETAFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAFailsIfRevertGasLimitExceeded")
}

// TestWithdrawZETAFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xd597a27a.
//
// Solidity: function testWithdrawZETAFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xd597a27a.
//
// Solidity: function testWithdrawZETAFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithBothDefaultValuesFromRegistry is a paid mutator transaction binding the contract method 0x3828ce8c.
//
// Solidity: function testWithdrawZETAWithBothDefaultValuesFromRegistry() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithBothDefaultValuesFromRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithBothDefaultValuesFromRegistry")
}

// TestWithdrawZETAWithBothDefaultValuesFromRegistry is a paid mutator transaction binding the contract method 0x3828ce8c.
//
// Solidity: function testWithdrawZETAWithBothDefaultValuesFromRegistry() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithBothDefaultValuesFromRegistry() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithBothDefaultValuesFromRegistry(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithBothDefaultValuesFromRegistry is a paid mutator transaction binding the contract method 0x3828ce8c.
//
// Solidity: function testWithdrawZETAWithBothDefaultValuesFromRegistry() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithBothDefaultValuesFromRegistry() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithBothDefaultValuesFromRegistry(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsWithMessage is a paid mutator transaction binding the contract method 0x14759766.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithCallOptsWithMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithCallOptsWithMessage")
}

// TestWithdrawZETAWithCallOptsWithMessage is a paid mutator transaction binding the contract method 0x14759766.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithCallOptsWithMessage() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsWithMessage(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsWithMessage is a paid mutator transaction binding the contract method 0x14759766.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithCallOptsWithMessage() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsWithMessage(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x64002a1f.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero")
}

// TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x64002a1f.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x64002a1f.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xe51c6388.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance")
}

// TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xe51c6388.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xe51c6388.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithDefaultGasLimit is a paid mutator transaction binding the contract method 0x16c09ef7.
//
// Solidity: function testWithdrawZETAWithDefaultGasLimit() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithDefaultGasLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithDefaultGasLimit")
}

// TestWithdrawZETAWithDefaultGasLimit is a paid mutator transaction binding the contract method 0x16c09ef7.
//
// Solidity: function testWithdrawZETAWithDefaultGasLimit() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithDefaultGasLimit() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithDefaultGasLimit(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithDefaultGasLimit is a paid mutator transaction binding the contract method 0x16c09ef7.
//
// Solidity: function testWithdrawZETAWithDefaultGasLimit() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithDefaultGasLimit() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithDefaultGasLimit(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithDefaultProtocolFlatFee is a paid mutator transaction binding the contract method 0x291f1d82.
//
// Solidity: function testWithdrawZETAWithDefaultProtocolFlatFee() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithDefaultProtocolFlatFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithDefaultProtocolFlatFee")
}

// TestWithdrawZETAWithDefaultProtocolFlatFee is a paid mutator transaction binding the contract method 0x291f1d82.
//
// Solidity: function testWithdrawZETAWithDefaultProtocolFlatFee() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithDefaultProtocolFlatFee() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithDefaultProtocolFlatFee(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithDefaultProtocolFlatFee is a paid mutator transaction binding the contract method 0x291f1d82.
//
// Solidity: function testWithdrawZETAWithDefaultProtocolFlatFee() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithDefaultProtocolFlatFee() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithDefaultProtocolFlatFee(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20 is a paid mutator transaction binding the contract method 0xfbc611c8.
//
// Solidity: function testWithdrawZRC20() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20")
}

// TestWithdrawZRC20 is a paid mutator transaction binding the contract method 0xfbc611c8.
//
// Solidity: function testWithdrawZRC20() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20 is a paid mutator transaction binding the contract method 0xfbc611c8.
//
// Solidity: function testWithdrawZRC20() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x264b524c.
//
// Solidity: function testWithdrawZRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIfMessageSizeExceeded")
}

// TestWithdrawZRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x264b524c.
//
// Solidity: function testWithdrawZRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x264b524c.
//
// Solidity: function testWithdrawZRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfNoAllowance is a paid mutator transaction binding the contract method 0x5d72228f.
//
// Solidity: function testWithdrawZRC20FailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIfNoAllowance")
}

// TestWithdrawZRC20FailsIfNoAllowance is a paid mutator transaction binding the contract method 0x5d72228f.
//
// Solidity: function testWithdrawZRC20FailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfNoAllowance is a paid mutator transaction binding the contract method 0x5d72228f.
//
// Solidity: function testWithdrawZRC20FailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfNoBalanceForGasFee is a paid mutator transaction binding the contract method 0xceccfab3.
//
// Solidity: function testWithdrawZRC20FailsIfNoBalanceForGasFee() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIfNoBalanceForGasFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIfNoBalanceForGasFee")
}

// TestWithdrawZRC20FailsIfNoBalanceForGasFee is a paid mutator transaction binding the contract method 0xceccfab3.
//
// Solidity: function testWithdrawZRC20FailsIfNoBalanceForGasFee() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIfNoBalanceForGasFee() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfNoBalanceForGasFee(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfNoBalanceForGasFee is a paid mutator transaction binding the contract method 0xceccfab3.
//
// Solidity: function testWithdrawZRC20FailsIfNoBalanceForGasFee() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIfNoBalanceForGasFee() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfNoBalanceForGasFee(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfNoBalanceForTransfer is a paid mutator transaction binding the contract method 0xfdad0ad0.
//
// Solidity: function testWithdrawZRC20FailsIfNoBalanceForTransfer() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIfNoBalanceForTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIfNoBalanceForTransfer")
}

// TestWithdrawZRC20FailsIfNoBalanceForTransfer is a paid mutator transaction binding the contract method 0xfdad0ad0.
//
// Solidity: function testWithdrawZRC20FailsIfNoBalanceForTransfer() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIfNoBalanceForTransfer() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfNoBalanceForTransfer(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfNoBalanceForTransfer is a paid mutator transaction binding the contract method 0xfdad0ad0.
//
// Solidity: function testWithdrawZRC20FailsIfNoBalanceForTransfer() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIfNoBalanceForTransfer() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfNoBalanceForTransfer(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x83ababa9.
//
// Solidity: function testWithdrawZRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIfReceiverIsZeroAddress")
}

// TestWithdrawZRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x83ababa9.
//
// Solidity: function testWithdrawZRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x83ababa9.
//
// Solidity: function testWithdrawZRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc3bb7957.
//
// Solidity: function testWithdrawZRC20FailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIfRevertGasLimitExceeded")
}

// TestWithdrawZRC20FailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc3bb7957.
//
// Solidity: function testWithdrawZRC20FailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc3bb7957.
//
// Solidity: function testWithdrawZRC20FailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIsAmountIs0 is a paid mutator transaction binding the contract method 0xba800c91.
//
// Solidity: function testWithdrawZRC20FailsIsAmountIs0() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIsAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIsAmountIs0")
}

// TestWithdrawZRC20FailsIsAmountIs0 is a paid mutator transaction binding the contract method 0xba800c91.
//
// Solidity: function testWithdrawZRC20FailsIsAmountIs0() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIsAmountIs0() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIsAmountIs0(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIsAmountIs0 is a paid mutator transaction binding the contract method 0xba800c91.
//
// Solidity: function testWithdrawZRC20FailsIsAmountIs0() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIsAmountIs0() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIsAmountIs0(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCallOptsWithMessage is a paid mutator transaction binding the contract method 0xe804a406.
//
// Solidity: function testWithdrawZRC20WithCallOptsWithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithCallOptsWithMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithCallOptsWithMessage")
}

// TestWithdrawZRC20WithCallOptsWithMessage is a paid mutator transaction binding the contract method 0xe804a406.
//
// Solidity: function testWithdrawZRC20WithCallOptsWithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithCallOptsWithMessage() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCallOptsWithMessage(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCallOptsWithMessage is a paid mutator transaction binding the contract method 0xe804a406.
//
// Solidity: function testWithdrawZRC20WithCallOptsWithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithCallOptsWithMessage() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCallOptsWithMessage(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithMessage is a paid mutator transaction binding the contract method 0x1e63d2b9.
//
// Solidity: function testWithdrawZRC20WithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithMessage")
}

// TestWithdrawZRC20WithMessage is a paid mutator transaction binding the contract method 0x1e63d2b9.
//
// Solidity: function testWithdrawZRC20WithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithMessage() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithMessage(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithMessage is a paid mutator transaction binding the contract method 0x1e63d2b9.
//
// Solidity: function testWithdrawZRC20WithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithMessage() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithMessage(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x1238212c.
//
// Solidity: function testWithdrawZRC20WithMessageFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithMessageFailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithMessageFailsIfNoAllowance")
}

// TestWithdrawZRC20WithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x1238212c.
//
// Solidity: function testWithdrawZRC20WithMessageFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithMessageFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithMessageFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x1238212c.
//
// Solidity: function testWithdrawZRC20WithMessageFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithMessageFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithMessageFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x6d6ce0d0.
//
// Solidity: function testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance")
}

// TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x6d6ce0d0.
//
// Solidity: function testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x6d6ce0d0.
//
// Solidity: function testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// GatewayZEVMInboundTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestCalledIterator struct {
	Event *GatewayZEVMInboundTestCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestCalled represents a Called event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestCalled struct {
	Sender        common.Address
	Zrc20         common.Address
	Receiver      []byte
	Message       []byte
	CallOptions   CallOptions
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, zrc20 []common.Address) (*GatewayZEVMInboundTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestCalledIterator{contract: _GatewayZEVMInboundTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestCalled, sender []common.Address, zrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestCalled)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "Called", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCalled is a log parse operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseCalled(log types.Log) (*GatewayZEVMInboundTestCalled, error) {
	event := new(GatewayZEVMInboundTestCalled)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestWithdrawnIterator struct {
	Event *GatewayZEVMInboundTestWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestWithdrawn represents a Withdrawn event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestWithdrawn struct {
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

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMInboundTestWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestWithdrawnIterator{contract: _GatewayZEVMInboundTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestWithdrawn)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseWithdrawn(log types.Log) (*GatewayZEVMInboundTestWithdrawn, error) {
	event := new(GatewayZEVMInboundTestWithdrawn)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestWithdrawnAndCalledIterator struct {
	Event *GatewayZEVMInboundTestWithdrawnAndCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestWithdrawnAndCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestWithdrawnAndCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestWithdrawnAndCalled struct {
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

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMInboundTestWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestWithdrawnAndCalledIterator{contract: _GatewayZEVMInboundTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestWithdrawnAndCalled)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayZEVMInboundTestWithdrawnAndCalled, error) {
	event := new(GatewayZEVMInboundTestWithdrawnAndCalled)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestWithdrawnV2Iterator struct {
	Event *GatewayZEVMInboundTestWithdrawnV2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestWithdrawnV2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestWithdrawnV2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestWithdrawnV2 represents a WithdrawnV2 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestWithdrawnV2 struct {
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

// FilterWithdrawnV2 is a free log retrieval operation binding the contract event 0x5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d97.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMInboundTestWithdrawnV2Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "WithdrawnV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestWithdrawnV2Iterator{contract: _GatewayZEVMInboundTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0x5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d97.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestWithdrawnV2, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "WithdrawnV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestWithdrawnV2)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnV2 is a log parse operation binding the contract event 0x5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d97.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseWithdrawnV2(log types.Log) (*GatewayZEVMInboundTestWithdrawnV2, error) {
	event := new(GatewayZEVMInboundTestWithdrawnV2)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogIterator struct {
	Event *GatewayZEVMInboundTestLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLog represents a Log event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogIterator{contract: _GatewayZEVMInboundTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLog)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLog(log types.Log) (*GatewayZEVMInboundTestLog, error) {
	event := new(GatewayZEVMInboundTestLog)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogAddressIterator struct {
	Event *GatewayZEVMInboundTestLogAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogAddress represents a LogAddress event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogAddressIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogAddressIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogAddress)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogAddress(log types.Log) (*GatewayZEVMInboundTestLogAddress, error) {
	event := new(GatewayZEVMInboundTestLogAddress)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogArrayIterator struct {
	Event *GatewayZEVMInboundTestLogArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogArray represents a LogArray event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogArrayIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogArrayIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogArray)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogArray(log types.Log) (*GatewayZEVMInboundTestLogArray, error) {
	event := new(GatewayZEVMInboundTestLogArray)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogArray0Iterator struct {
	Event *GatewayZEVMInboundTestLogArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogArray0 represents a LogArray0 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogArray0Iterator{contract: _GatewayZEVMInboundTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogArray0)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogArray0(log types.Log) (*GatewayZEVMInboundTestLogArray0, error) {
	event := new(GatewayZEVMInboundTestLogArray0)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogArray1Iterator struct {
	Event *GatewayZEVMInboundTestLogArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogArray1 represents a LogArray1 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogArray1Iterator{contract: _GatewayZEVMInboundTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogArray1)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogArray1(log types.Log) (*GatewayZEVMInboundTestLogArray1, error) {
	event := new(GatewayZEVMInboundTestLogArray1)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogBytesIterator struct {
	Event *GatewayZEVMInboundTestLogBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogBytes represents a LogBytes event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogBytesIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogBytesIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogBytes)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogBytes(log types.Log) (*GatewayZEVMInboundTestLogBytes, error) {
	event := new(GatewayZEVMInboundTestLogBytes)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogBytes32Iterator struct {
	Event *GatewayZEVMInboundTestLogBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogBytes32 represents a LogBytes32 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogBytes32Iterator{contract: _GatewayZEVMInboundTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogBytes32)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogBytes32(log types.Log) (*GatewayZEVMInboundTestLogBytes32, error) {
	event := new(GatewayZEVMInboundTestLogBytes32)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogIntIterator struct {
	Event *GatewayZEVMInboundTestLogInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogInt represents a LogInt event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogIntIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogIntIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogInt)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogInt(log types.Log) (*GatewayZEVMInboundTestLogInt, error) {
	event := new(GatewayZEVMInboundTestLogInt)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedAddressIterator struct {
	Event *GatewayZEVMInboundTestLogNamedAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogNamedAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedAddressIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedAddress)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayZEVMInboundTestLogNamedAddress, error) {
	event := new(GatewayZEVMInboundTestLogNamedAddress)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedArrayIterator struct {
	Event *GatewayZEVMInboundTestLogNamedArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogNamedArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedArray represents a LogNamedArray event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedArrayIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedArray)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayZEVMInboundTestLogNamedArray, error) {
	event := new(GatewayZEVMInboundTestLogNamedArray)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedArray0Iterator struct {
	Event *GatewayZEVMInboundTestLogNamedArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogNamedArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedArray0Iterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedArray0)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayZEVMInboundTestLogNamedArray0, error) {
	event := new(GatewayZEVMInboundTestLogNamedArray0)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedArray1Iterator struct {
	Event *GatewayZEVMInboundTestLogNamedArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogNamedArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedArray1Iterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedArray1)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayZEVMInboundTestLogNamedArray1, error) {
	event := new(GatewayZEVMInboundTestLogNamedArray1)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedBytesIterator struct {
	Event *GatewayZEVMInboundTestLogNamedBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogNamedBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedBytesIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedBytes)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayZEVMInboundTestLogNamedBytes, error) {
	event := new(GatewayZEVMInboundTestLogNamedBytes)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedBytes32Iterator struct {
	Event *GatewayZEVMInboundTestLogNamedBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogNamedBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedBytes32Iterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedBytes32)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayZEVMInboundTestLogNamedBytes32, error) {
	event := new(GatewayZEVMInboundTestLogNamedBytes32)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedDecimalIntIterator struct {
	Event *GatewayZEVMInboundTestLogNamedDecimalInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedDecimalInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogNamedDecimalInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedDecimalIntIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedDecimalInt)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayZEVMInboundTestLogNamedDecimalInt, error) {
	event := new(GatewayZEVMInboundTestLogNamedDecimalInt)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedDecimalUintIterator struct {
	Event *GatewayZEVMInboundTestLogNamedDecimalUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedDecimalUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogNamedDecimalUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedDecimalUintIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedDecimalUint)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayZEVMInboundTestLogNamedDecimalUint, error) {
	event := new(GatewayZEVMInboundTestLogNamedDecimalUint)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedIntIterator struct {
	Event *GatewayZEVMInboundTestLogNamedInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogNamedInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedInt represents a LogNamedInt event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedIntIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedInt)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayZEVMInboundTestLogNamedInt, error) {
	event := new(GatewayZEVMInboundTestLogNamedInt)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedStringIterator struct {
	Event *GatewayZEVMInboundTestLogNamedString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogNamedString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedString represents a LogNamedString event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedStringIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedString)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedString(log types.Log) (*GatewayZEVMInboundTestLogNamedString, error) {
	event := new(GatewayZEVMInboundTestLogNamedString)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedUintIterator struct {
	Event *GatewayZEVMInboundTestLogNamedUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogNamedUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedUint represents a LogNamedUint event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedUintIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedUint)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayZEVMInboundTestLogNamedUint, error) {
	event := new(GatewayZEVMInboundTestLogNamedUint)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogStringIterator struct {
	Event *GatewayZEVMInboundTestLogString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogString represents a LogString event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogStringIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogStringIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogString)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogString(log types.Log) (*GatewayZEVMInboundTestLogString, error) {
	event := new(GatewayZEVMInboundTestLogString)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogUintIterator struct {
	Event *GatewayZEVMInboundTestLogUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogUint represents a LogUint event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogUintIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogUintIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogUint)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogUint(log types.Log) (*GatewayZEVMInboundTestLogUint, error) {
	event := new(GatewayZEVMInboundTestLogUint)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogsIterator struct {
	Event *GatewayZEVMInboundTestLogs // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMInboundTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogs)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMInboundTestLogs)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMInboundTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogs represents a Logs event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogsIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogsIterator{contract: _GatewayZEVMInboundTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogs)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "logs", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogs(log types.Log) (*GatewayZEVMInboundTestLogs, error) {
	event := new(GatewayZEVMInboundTestLogs)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
