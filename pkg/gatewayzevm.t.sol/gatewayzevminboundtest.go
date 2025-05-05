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
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testCallWithCallOpts\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfGasLimitIsBelowMin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpgradeAndWithdrawZRC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20SolanaFails\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETA\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfNoBalance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsWithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfNoBalanceForGasFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfNoBalanceForTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIsAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20Solana\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCallOptsWithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithMessageFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZETANotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5062011972806200003e6000396000f3fe608060405234801561001057600080fd5b50600436106103575760003560e01c806373a3db40116101c8578063bed3e81311610104578063e51c6388116100a2578063eb7a2fac1161007c578063eb7a2fac14610538578063fa7626d414610540578063fbc611c81461054d578063fdad0ad01461055557600080fd5b8063e51c638814610528578063e804a4061461038e578063ea37902f1461053057600080fd5b8063ceccfab3116100de578063ceccfab314610508578063dc749dd714610510578063dde7e96714610518578063e20c9f711461052057600080fd5b8063bed3e813146104f0578063c20049f4146104f8578063c946d7c01461050057600080fd5b8063a721b2d311610171578063b5508aa91161014b578063b5508aa9146104c0578063ba414fa6146104c8578063ba800c91146104e0578063ba9adeef146104e857600080fd5b8063a721b2d3146104a8578063b0464fdc146104b0578063b51ac071146104b857600080fd5b806383ababa9116101a257806383ababa91461047657806385226c811461047e578063916a17c61461049357600080fd5b806373a3db401461045e57806377def36e146104665780637ae697301461046e57600080fd5b80633e5e3c23116102975780635efe72a91161024057806364002a1f1161021a57806364002a1f1461043157806366d9a9a0146104395780636d6ce0d01461044e5780636dfcbc501461045657600080fd5b80635efe72a9146104195780636198fb19146104215780636221b5091461042957600080fd5b8063431814371161027157806343181437146104015780634ffab9de146104095780635d72228f1461041157600080fd5b80633e5e3c23146103e95780633f7286f4146103f157806342752d41146103f957600080fd5b80631e63d2b91161030457806321aeb18c116102de57806321aeb18c146103bc578063264b524c146103c45780632ade3880146103cc57806336431b3f146103e157600080fd5b80631e63d2b91461038e5780631ed7831c1461039657806320dee15f146103b457600080fd5b80631238212c116103355780631238212c14610376578063147597661461037e5780631b9641bf1461038657600080fd5b806305b9f0461461035c5780630a9254e4146103665780630b5ad28d1461036e575b600080fd5b61036461055d565b005b61036461071c565b61036461156b565b6103646116e9565b610364611ab7565b610364611ce3565b610364611e5a565b61039e612276565b6040516103ab919061b7eb565b60405180910390f35b6103646122d8565b6103646126c6565b6103646127e4565b6103d4612a06565b6040516103ab919061b887565b610364612b48565b61039e612c9e565b61039e612cfe565b610364612d5e565b610364612edc565b61036461300f565b6103646132fa565b61036461364c565b610364613704565b6103646138ca565b610364613ab8565b610441613c94565b6040516103ab919061b9ed565b610364613e16565b61036461411b565b610364614334565b6103646146c2565b6103646148f7565b610364614ac5565b610486614ba9565b6040516103ab919061ba8b565b61049b614c79565b6040516103ab919061bb02565b610364614d74565b61049b614f8d565b610364615088565b610486615187565b6104d0615257565b60405190151581526020016103ab565b61036461532b565b610364615441565b6103646155c9565b610364615742565b6103646158d9565b610364615b5d565b610364615ea9565b610364615f4c565b61039e616389565b6103646163e9565b610364616892565b610364616a51565b601f546104d09060ff1681565b610364616b50565b610364616e92565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905290506105c86001620186a061bbc8565b602d55604051630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b600060405180830381600087803b15801561063857600080fd5b505af115801561064c573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506306cb898391506034015b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526106e7926001600160a01b03909116908690602d9060299060040161bd17565b600060405180830381600087803b15801561070157600080fd5b505af1158015610715573d6000803e3d6000fd5b5050505050565b602680547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602780549091166112341790556040516107629061b6fa565b604051809103906000f08015801561077e573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600f81527f476174657761795a45564d2e736f6c000000000000000000000000000000000060208201526026549151602481019390935292166044820152610862919060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f485cc95500000000000000000000000000000000000000000000000000000000179052617008565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b039384168102919091179182905560208054919092049092167fffffffffffffffffffffffff000000000000000000000000000000000000000090921682178155604080517f2722feee0000000000000000000000000000000000000000000000000000000081529051632722feee926004808401939192918290030181865afa158015610924573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610948919061bd81565b602880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039290921691909117905560405161098c9061b708565b604051809103906000f0801580156109a8573d6000803e3d6000fd5b50602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556028546040517f06447d5600000000000000000000000000000000000000000000000000000000815291166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015610a4657600080fd5b505af1158015610a5a573d6000803e3d6000fd5b505050506000806000604051610a6f9061b716565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610aab573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602054604051601293600193849360009391921690610b019061b724565b610b109695949392919061bde5565b604051809103906000f080158015610b2c573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602480546040517fee2815ba000000000000000000000000000000000000000000000000000000008152600160048201529182019290925291169063ee2815ba90604401600060405180830381600087803b158015610bc257600080fd5b505af1158015610bd6573d6000803e3d6000fd5b5050602480546040517fa7cb0507000000000000000000000000000000000000000000000000000000008152600160048201819052928101929092526001600160a01b0316925063a7cb05079150604401600060405180830381600087803b158015610c4157600080fd5b505af1158015610c55573d6000803e3d6000fd5b50506024546020546040516009945061038493506001926000926001600160a01b0391821692911690610c879061b724565b610c969695949392919061be9f565b604051809103906000f080158015610cb2573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602480546040517fee2815ba00000000000000000000000000000000000000000000000000000000815261038460048201529182019290925291169063ee2815ba90604401600060405180830381600087803b158015610d4957600080fd5b505af1158015610d5d573d6000803e3d6000fd5b5050602480546040517fa7cb050700000000000000000000000000000000000000000000000000000000815261038460048201526001928101929092526001600160a01b0316925063a7cb05079150604401600060405180830381600087803b158015610dc957600080fd5b505af1158015610ddd573d6000803e3d6000fd5b50506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152633b9aca006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b158015610e5d57600080fd5b505af1158015610e71573d6000803e3d6000fd5b50505050602360009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b158015610ec657600080fd5b505af1158015610eda573d6000803e3d6000fd5b50506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af1158015610f4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f72919061bf2f565b506021546026546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116906347e7ef24906044016020604051808303816000875af1158015610fe4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611008919061bf2f565b506022546026546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116906347e7ef24906044016020604051808303816000875af115801561107a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061109e919061bf2f565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156110fd57600080fd5b505af1158015611111573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561118757600080fd5b505af115801561119b573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116925063095ea7b391506044016020604051808303816000875af1158015611210573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611234919061bf2f565b506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e100602482015291169063095ea7b3906044016020604051808303816000875af11580156112a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ca919061bf2f565b50602360009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b15801561131c57600080fd5b505af1158015611330573d6000803e3d6000fd5b50506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af11580156113a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c8919061bf2f565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561142757600080fd5b505af115801561143b573d6000803e3d6000fd5b50506040805160a08101825261032180825260016020808401918252838501928352845190810190945260008085526060840185905260808401528251602980549251151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009093166001600160a01b03928316179290921782559151602a8054919093167fffffffffffffffffffffffff000000000000000000000000000000000000000091909116179091559093509150602b90611512908261bfc7565b506080919091015160039091015560408051808201909152620186a080825260016020909201829052602d55602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169091179055565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b1790526000602d5551630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b600060405180830381600087803b15801561163957600080fd5b505af115801561164d573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b91506034015b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526106e7926001916001600160a01b0316908790602d9060299060040161c086565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa15801561173a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061175e919061c0f8565b6026546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156117d357600080fd5b505af11580156117e7573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015611859573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061187d919061bf2f565b506027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905280517ff48448140000000000000000000000000000000000000000000000000000000081529051919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f48448149160048082019260009290919082900301818387803b15801561194057600080fd5b505af1158015611954573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a084526000602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b168152611a0093919289926001600160a01b039091169188919060299060040161c111565b600060405180830381600087803b158015611a1a57600080fd5b505af1158015611a2e573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015611a81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aa5919061c0f8565b9050611ab18382617027565b50505050565b6027546040516001600160a01b03909116602482015260019060009060440160408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b179052602654602754925193945060019384936001600160a01b03928316937f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c93611b79939216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602354611bb1926001600160a01b0390911690899060009081908b90602d9060299061c165565b60405180910390a3604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b600060405180830381600087803b158015611c2657600080fd5b505af1158015611c3a573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250632810ae639150603401604051602081830303815290604052858486602d60296040518763ffffffff1660e01b8152600401611cae9695949392919061c1d9565b60006040518083038186803b158015611cc657600080fd5b505afa158015611cda573d6000803e3d6000fd5b50505050505050565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015611dab57600080fd5b505af1158015611dbf573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526106e7926000916001600160a01b0316908790602d9060299060040161c086565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015611eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ecf919061c0f8565b6027546040516001600160a01b03909116602482015290915060009060440160408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150620186a0908190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611fc357600080fd5b505af1158015611fd7573d6000803e3d6000fd5b505060265460275460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507fd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a9060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b03909216918b9189918491634d8943bb916004808201926020929091908290030181865afa1580156120a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120cd919061c0f8565b6040805180820182528a81526001602082015290516120f59695949392918d9160299061c20a565b60405180910390a3602080546027546040516001600160a01b0392831693637b15118b9361213d9316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815260215483830183528684526001602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526121a99391928b926001600160a01b03909116918a919060299060040161c111565b600060405180830381600087803b1580156121c357600080fd5b505af11580156121d7573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561222a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061224e919061c0f8565b905061226e8361225e888861bbc8565b612268919061bbc8565b82617027565b505050505050565b606060168054806020026020016040519081016040528092919081815260200182805480156122ce57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116122b0575b5050505050905090565b601f54604080518082018252601a81527f476174657761795a45564d55706772616465546573742e736f6c00000000000060208083019190915282519081019092526000825260265461233d936001600160a01b03610100909104811693911661707f565b601f546021546026546040516370a0823160e01b81526001600160a01b03918216600482015261010090930481169260019260009216906370a0823190602401602060405180830381865afa15801561239a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123be919061c0f8565b6040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b0385166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561244b57600080fd5b505af115801561245f573d6000803e3d6000fd5b505060265460275460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507f5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d979060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b039092169188916000918491634d8943bb916004808201926020929091908290030181865afa158015612532573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612556919061c0f8565b6040805180820182526000815260016020820152905161257e9695949392919060299061c26c565b60405180910390a360275460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b03841690637c0dcb5f9060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526126129287916001600160a01b03169060299060040161c2ee565b600060405180830381600087803b15801561262c57600080fd5b505af1158015612640573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015612693573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126b7919061c0f8565b9050611ab1612268848461bbc8565b604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561273257600080fd5b505af1158015612746573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef191506034015b6040516020818303038152906040526000600160296040518563ffffffff1660e01b81526004016127b8949392919061c328565b60006040518083038186803b1580156127d057600080fd5b505afa158015611ab1573d6000803e3d6000fd5b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516001600160a01b03909216926397d340f5926004808401938290030181865afa158015612842573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612866919061c0f8565b61287190600161c359565b67ffffffffffffffff8111156128895761288961bf51565b6040519080825280601f01601f1916602001820160405280156128b3576020820181803683370190505b50602b906128c1908261bfc7565b50604051630618f58760e51b81527f9507fb3d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561292e57600080fd5b505af1158015612942573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526129d8926002916001600160a01b03169060299060040161c2ee565b600060405180830381600087803b1580156129f257600080fd5b505af1158015611ab1573d6000803e3d6000fd5b6060601e805480602002602001604051908101604052809291908181526020016000905b82821015612b3f57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015612b28578382906000526020600020018054612a9b9061bbdb565b80601f0160208091040260200160405190810160405280929190818152602001828054612ac79061bbdb565b8015612b145780601f10612ae957610100808354040283529160200191612b14565b820191906000526020600020905b815481529060010190602001808311612af757829003601f168201915b505050505081526020019060010190612a7c565b505050508152505081526020019060010190612a2a565b50505050905090565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612c1057600080fd5b505af1158015612c24573d6000803e3d6000fd5b505060208054604080516000808252602154606083018452620186a09583019586528284019190915291517f7b15118b0000000000000000000000000000000000000000000000000000000081526001600160a01b039384169650637b15118b95506106e79491936001931691889160299060040161c111565b606060188054806020026020016040519081016040528092919081815260200182805480156122ce576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116122b0575050505050905090565b606060178054806020026020016040519081016040528092919081815260200182805480156122ce576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116122b0575050505050905090565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612e2657600080fd5b505af1158015612e3a573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250632810ae6391506034016040516020818303038152906040526000600185602d60296040518763ffffffff1660e01b8152600401612eb09695949392919061c1d9565b60006040518083038186803b158015612ec857600080fd5b505afa158015610715573d6000803e3d6000fd5b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612fa457600080fd5b505af1158015612fb8573d6000803e3d6000fd5b5050602080546040516001600160a01b039091169350632810ae639250015b60405160208183030381529060405260018085602d60296040518763ffffffff1660e01b8152600401612eb09695949392919061c1d9565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516000936002936001600160a01b0316926397d340f592600480830193928290030181865afa158015613072573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613096919061c0f8565b6130a0919061c36c565b67ffffffffffffffff8111156130b8576130b861bf51565b6040519080825280601f01601f1916602001820160405280156130e2576020820181803683370190505b5060208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b03909216926397d340f5926004808401938290030181865afa158015613147573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061316b919061c0f8565b613175919061c36c565b61318090600161c359565b67ffffffffffffffff8111156131985761319861bf51565b6040519080825280601f01601f1916602001820160405280156131c2576020820181803683370190505b50602b906131d0908261bfc7565b50604051630618f58760e51b81527f9507fb3d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561323d57600080fd5b505af1158015613251573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152602154838301835260018085526000602086015292517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b1681526106e7949293926001600160a01b0390921691889160299060040161c111565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa15801561334b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061336f919061c0f8565b6026546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156133e457600080fd5b505af11580156133f8573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af115801561346a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061348e919061bf2f565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156134ed57600080fd5b505af1158015613501573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526135969287916001600160a01b03169060299060040161c2ee565b600060405180830381600087803b1580156135b057600080fd5b505af11580156135c4573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015613617573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061363b919061c0f8565b90506136478282617027565b505050565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b1790526000602d5551630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e09060240161061e565b6027546040516001600160a01b03909116602482015260009060440160408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156137ee57600080fd5b505af1158015613802573d6000803e3d6000fd5b505060215460265460275460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039283169450911691507f306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e49060340160408051601f1981840301815290829052613882918690602d9060299061c3a7565b60405180910390a3602080546027546040516001600160a01b03928316936306cb89839361068a9316910160609190911b6bffffffffffffffffffffffff1916815260140190565b6023546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa15801561391b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061393f919061c0f8565b6023546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526101236004820152602481018390529192506001600160a01b03169063a9059cbb906044016020604051808303816000875af11580156139ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139d0919061bf2f565b506000600190507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b158015613a3557600080fd5b505af1158015613a49573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef19150603401604051602081830303815290604052858460296040518563ffffffff1660e01b8152600401611cae949392919061c328565b6027546040516001600160a01b03909116602482015260019060009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905260265490517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613b8d57600080fd5b505af1158015613ba1573d6000803e3d6000fd5b50506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015613c13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c37919061bf2f565b506000602d55604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401611c0c565b6060601b805480602002602001604051908101604052809291908181526020016000905b82821015612b3f5783829060005260206000209060020201604051806040016040529081600082018054613ceb9061bbdb565b80601f0160208091040260200160405190810160405280929190818152602001828054613d179061bbdb565b8015613d645780601f10613d3957610100808354040283529160200191613d64565b820191906000526020600020905b815481529060010190602001808311613d4757829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015613dfe57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411613dab5790505b50505050508152505081526020019060010190613cb8565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015613e67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e8b919061c0f8565b6026546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613f0057600080fd5b505af1158015613f14573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015613f86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613faa919061bf2f565b506027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905280517ff48448140000000000000000000000000000000000000000000000000000000081529051919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f48448149160048082019260009290919082900301818387803b15801561406d57600080fd5b505af1158015614081573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352611a009288916001600160a01b0316908790602d9060299060040161c086565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516000936002936001600160a01b0316926397d340f592600480830193928290030181865afa15801561417e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141a2919061c0f8565b6141ac919061c36c565b67ffffffffffffffff8111156141c4576141c461bf51565b6040519080825280601f01601f1916602001820160405280156141ee576020820181803683370190505b5060208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b03909216926397d340f5926004808401938290030181865afa158015614253573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614277919061c0f8565b614281919061c36c565b61428c90600161c359565b67ffffffffffffffff8111156142a4576142a461bf51565b6040519080825280601f01601f1916602001820160405280156142ce576020820181803683370190505b50602b906142dc908261bfc7565b50604051630618f58760e51b81527f9507fb3d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e09060240161061e565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015614385573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906143a9919061c0f8565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561443a57600080fd5b505af115801561444e573d6000803e3d6000fd5b505060265460275460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f198184030181528282526022547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b039092169188916000918491634d8943bb916004808201926020929091908290030181865afa158015614521573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614545919061c0f8565b6040805180820182526000815260016020820152905161456d9695949392919060299061c26c565b60405180910390a3602080546027546040516001600160a01b0392831693637c0dcb5f936145b59316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526022547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261460d9287916001600160a01b03169060299060040161c2ee565b600060405180830381600087803b15801561462757600080fd5b505af115801561463b573d6000803e3d6000fd5b50506022546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a08231906024015b602060405180830381865afa15801561468f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906146b3919061c0f8565b9050613647612268848461bbc8565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015614713573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614737919061c0f8565b6027546040516001600160a01b03909116602482015290915060009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905280517ff484481400000000000000000000000000000000000000000000000000000000815290519192506001918291737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f48448149160048082019260009290919082900301818387803b15801561480157600080fd5b505af1158015614815573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f1981840301815260225483830183528684526001602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526148be9391928b926001600160a01b03909116918a919060299060040161c111565b600060405180830381600087803b1580156148d857600080fd5b505af11580156148ec573d6000803e3d6000fd5b505050505050505050565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516001600160a01b03909216926397d340f5926004808401938290030181865afa158015614955573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614979919061c0f8565b61498490600161c359565b67ffffffffffffffff81111561499c5761499c61bf51565b6040519080825280601f01601f1916602001820160405280156149c6576020820181803683370190505b50602b906149d4908261bfc7565b50604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015614a4157600080fd5b505af1158015614a55573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef1915060340160405160208183030381529060405260018060296040518563ffffffff1660e01b81526004016127b8949392919061c328565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015614b3157600080fd5b505af1158015614b45573d6000803e3d6000fd5b5050602080546040805160008152928301908190526021547f7c0dcb5f000000000000000000000000000000000000000000000000000000009091526001600160a01b039182169450637c0dcb5f93506129d892916001911660296024840161c2ee565b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015612b3f578382906000526020600020018054614bec9061bbdb565b80601f0160208091040260200160405190810160405280929190818152602001828054614c189061bbdb565b8015614c655780601f10614c3a57610100808354040283529160200191614c65565b820191906000526020600020905b815481529060010190602001808311614c4857829003601f168201915b505050505081526020019060010190614bcd565b6060601d805480602002602001604051908101604052809291908181526020016000905b82821015612b3f5760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015614d5c57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411614d095790505b50505050508152505081526020019060010190614c9d565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516000936002936001600160a01b0316926397d340f592600480830193928290030181865afa158015614dd7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614dfb919061c0f8565b614e05919061c36c565b67ffffffffffffffff811115614e1d57614e1d61bf51565b6040519080825280601f01601f191660200182016040528015614e47576020820181803683370190505b5060208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b03909216926397d340f5926004808401938290030181865afa158015614eac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ed0919061c0f8565b614eda919061c36c565b614ee590600161c359565b67ffffffffffffffff811115614efd57614efd61bf51565b6040519080825280601f01601f191660200182016040528015614f27576020820181803683370190505b50602b90614f35908261bfc7565b50604051630618f58760e51b81527f9507fb3d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e09060240161161f565b6060601c805480602002602001604051908101604052809291908181526020016000905b82821015612b3f5760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561507057602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001906004019060208260030104928301926001038202915080841161501d5790505b50505050508152505081526020019060010190614fb1565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561515057600080fd5b505af1158015615164573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911693506306cb898392500161068a565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015612b3f5783829060005260206000200180546151ca9061bbdb565b80601f01602080910402602001604051908101604052809291908181526020018280546151f69061bbdb565b80156152435780601f1061521857610100808354040283529160200191615243565b820191906000526020600020905b81548152906001019060200180831161522657829003601f168201915b5050505050815260200190600101906151ab565b60085460009060ff161561526f575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015615300573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615324919061c0f8565b1415905090565b604051630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561539757600080fd5b505af11580156153ab573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526129d8926000916001600160a01b03169060299060040161c2ee565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561550957600080fd5b505af115801561551d573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a0845260006020850181905292517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b1681526106e7949293926001600160a01b0390921691889160299060040161c111565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561569157600080fd5b505af11580156156a5573d6000803e3d6000fd5b5050602080546027546040805160609290921b6bffffffffffffffffffffffff1916938201939093528251808203601401815260215460748301855260006034840181815260549094015293517f7b15118b0000000000000000000000000000000000000000000000000000000081526001600160a01b039384169650637b15118b95506106e7949193600193921691889160299060040161c111565b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561580a57600080fd5b505af115801561581e573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152602154838301909252916001916001600160a01b03169086908061588f85620186a061bbc8565b815260006020909101526040517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b1681526106e795949392919060299060040161c111565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516000936002936001600160a01b0316926397d340f592600480830193928290030181865afa15801561593c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615960919061c0f8565b61596a919061c36c565b67ffffffffffffffff8111156159825761598261bf51565b6040519080825280601f01601f1916602001820160405280156159ac576020820181803683370190505b5060208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b03909216926397d340f5926004808401938290030181865afa158015615a11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615a35919061c0f8565b615a3f919061c36c565b615a4a90600161c359565b67ffffffffffffffff811115615a6257615a6261bf51565b6040519080825280601f01601f191660200182016040528015615a8c576020820181803683370190505b50602b90615a9a908261bfc7565b50604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015615b0757600080fd5b505af1158015615b1b573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250632810ae639150603401612fd7565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015615bae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615bd2919061c0f8565b6021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526101236004820152602481018390529192506001600160a01b03169063a9059cbb906044016020604051808303816000875af1158015615c3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615c63919061bf2f565b506028546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015615cd657600080fd5b505af1158015615cea573d6000803e3d6000fd5b50506021546040517ff687d12a000000000000000000000000000000000000000000000000000000008152600a60048201526001600160a01b03909116925063f687d12a9150602401600060405180830381600087803b158015615d4d57600080fd5b505af1158015615d61573d6000803e3d6000fd5b5050604051630618f58760e51b81527ffe382aa7000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b600060405180830381600087803b158015615dd257600080fd5b505af1158015615de6573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352615e7b9287916001600160a01b03169060299060040161c2ee565b600060405180830381600087803b158015615e9557600080fd5b505af115801561226e573d6000803e3d6000fd5b604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015615f1557600080fd5b505af1158015615f29573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911693506397a1cef1925001612784565b6023546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015615f9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615fc1919061c0f8565b6023546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa158015616013573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616037919061c0f8565b6028546026546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152929350163190600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156160b557600080fd5b505af11580156160c9573d6000803e3d6000fd5b50506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af115801561613b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061615f919061bf2f565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156161be57600080fd5b505af11580156161d2573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef19150603401604051602081830303815290604052878460296040518563ffffffff1660e01b8152600401616241949392919061c328565b60006040518083038186803b15801561625957600080fd5b505afa15801561626d573d6000803e3d6000fd5b50506023546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156162c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906162e4919061c0f8565b90506162f08582617027565b6023546020546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015616341573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616365919061c0f8565b90506163718582617027565b602854611cda9085906001600160a01b031631617027565b606060158054806020026020016040519081016040528092919081815260200182805480156122ce576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116122b0575050505050905090565b6023546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa15801561643a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061645e919061c0f8565b6023546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa1580156164b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906164d4919061c0f8565b6028546027546040516001600160a01b03918216602482015292935016319060009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905260265490517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156165af57600080fd5b505af11580156165c3573d6000803e3d6000fd5b50506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015616635573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616659919061bf2f565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156166b857600080fd5b505af11580156166cc573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250632810ae639150603401604051602081830303815290604052888486602d60296040518763ffffffff1660e01b81526004016167409695949392919061c1d9565b60006040518083038186803b15801561675857600080fd5b505afa15801561676c573d6000803e3d6000fd5b50506023546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156167bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906167e3919061c0f8565b90506167ef8682617027565b6023546020546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015616840573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616864919061c0f8565b90506168708682617027565b6028546168889086906001600160a01b031631617027565b5050505050505050565b60265460275460405160609190911b6bffffffffffffffffffffffff19166020820152600191829182916001600160a01b0316907f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f198184030181526023548383018352600080855260016020860152925161692e9492936001600160a01b03909216928a929182919060299061c26c565b60405180910390a3604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156169a257600080fd5b505af11580156169b6573d6000803e3d6000fd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef19150603401604051602081830303815290604052848460296040518563ffffffff1660e01b8152600401616a25949392919061c328565b60006040518083038186803b158015616a3d57600080fd5b505afa15801561226e573d6000803e3d6000fd5b6027546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015616b1957600080fd5b505af1158015616b2d573d6000803e3d6000fd5b5050602080546040516001600160a01b039091169350637b15118b92500161168b565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015616ba1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616bc5919061c0f8565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015616c5657600080fd5b505af1158015616c6a573d6000803e3d6000fd5b505060265460275460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b039092169188916000918491634d8943bb916004808201926020929091908290030181865afa158015616d3d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616d61919061c0f8565b60408051808201825260008152600160208201529051616d899695949392919060299061c26c565b60405180910390a3602080546027546040516001600160a01b0392831693637c0dcb5f93616dd19316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352616e299287916001600160a01b03169060299060040161c2ee565b600060405180830381600087803b158015616e4357600080fd5b505af1158015616e57573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401614672565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260029260009216906370a0823190602401602060405180830381865afa158015616ee3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616f07919061c0f8565b6021549091506001600160a01b031663a9059cbb610123616f2960018561bbc8565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015616f8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616fb0919061bf2f565b50604051630618f58760e51b81527ffe382aa7000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401615db8565b600061701261b732565b61701d848483617094565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c5490604401616a25565b61708761b732565b610715858585848661710f565b6000806170a1858461720f565b90506171046040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016170ef92919061c3f6565b6040516020818303038152906040528561721b565b9150505b9392505050565b6040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201528190737109709ecfa91a80626ff3989d68f67f5b1dd12d9081906306447d5690602401600060405180830381600087803b15801561718157600080fd5b505af1925050508015617192575060015b6171a7576171a287878787617249565b611cda565b6171b387878787617249565b806001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156171ee57600080fd5b505af1158015617202573d6000803e3d6000fd5b5050505050505050505050565b60006171088383617262565b60c0810151516000901561723f5761723884848460c0015161727d565b9050617108565b6172388484617423565b6000617255848361750e565b905061071585828561751a565b600061726e83836178e4565b6171088383602001518461721b565b6000806172886178f4565b9050600061729686836179c7565b905060006172ad8260600151836020015185617e6d565b905060006172bd8383898961807f565b905060006172ca82618efc565b602081015181519192509060030b1561733d578982604001516040516020016172f492919061c418565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526173349160040161c499565b60405180910390fd5b60006173806040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a2000000000000000000000008152508360016190cb565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906173d390849060040161c499565b602060405180830381865afa1580156173f0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617414919061bd81565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc9259061747890879060040161c499565b600060405180830381865afa158015617495573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526174bd919081019061c565565b905060006174eb82856040516020016174d792919061c59a565b6040516020818303038152906040526192cb565b90506001600160a01b03811661701d5784846040516020016172f492919061c5c9565b600061726e83836192de565b6040517f667f9d700000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90600090829063667f9d7090604401602060405180830381865afa1580156175b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906175da919061c0f8565b9050806177815760006175ec866192ea565b604080518082018252600581527f352e302e3000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150617677905b604080518082018252600080825260209182015281518083019092528451825280850190820152906193e2565b80617683575060008451115b15617706576040517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b03871690634f1ef286906176cf908890889060040161c3f6565b600060405180830381600087803b1580156176e957600080fd5b505af11580156176fd573d6000803e3d6000fd5b5050505061777b565b6040517f3659cfe60000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152871690633659cfe690602401600060405180830381600087803b15801561776257600080fd5b505af1158015617776573d6000803e3d6000fd5b505050505b50610715565b80600061778d826192ea565b604080518082018252600581527f352e302e30000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506177ef9061764a565b806177fb575060008551115b15617880576040517f9623609d0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690639623609d90617849908a908a908a9060040161c674565b600060405180830381600087803b15801561786357600080fd5b505af1158015617877573d6000803e3d6000fd5b50505050611cda565b6040517f99a88ec40000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015287811660248301528316906399a88ec490604401600060405180830381600087803b1580156171ee57600080fd5b6178f0828260006193f6565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c9061797b90849060040161c6a5565b600060405180830381865afa158015617998573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526179c0919081019061c6ec565b9250505090565b6179f96040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d9050617a446040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b617a4d856194f9565b60208201526000617a5d866198de565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa158015617a9f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617ac7919081019061c6ec565b86838560200151604051602001617ae1949392919061c735565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb1190617b3990859060040161c499565b600060405180830381865afa158015617b56573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617b7e919081019061c6ec565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f690617bc690849060040161c839565b602060405180830381865afa158015617be3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617c07919061bf2f565b617c1c57816040516020016172f4919061c88b565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890617c6190849060040161c91d565b600060405180830381865afa158015617c7e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617ca6919081019061c6ec565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f690617ced90849060040161c96f565b602060405180830381865afa158015617d0a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617d2e919061bf2f565b15617dc3576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890617d7890849060040161c96f565b600060405180830381865afa158015617d95573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617dbd919081019061c6ec565b60408501525b846001600160a01b03166349c4fac8828660000151604051602001617de8919061c9c1565b6040516020818303038152906040526040518363ffffffff1660e01b8152600401617e1492919061ca2d565b600060405180830381865afa158015617e31573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617e59919081019061c6ec565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b6060815260200190600190039081617e895790505090506040518060400160405280600481526020017f677265700000000000000000000000000000000000000000000000000000000081525081600081518110617ee957617ee961ca52565b60200260200101819052506040518060400160405280600381526020017f2d726c000000000000000000000000000000000000000000000000000000000081525081600181518110617f3d57617f3d61ca52565b602002602001018190525084604051602001617f59919061ca81565b60405160208183030381529060405281600281518110617f7b57617f7b61ca52565b602002602001018190525082604051602001617f97919061caed565b60405160208183030381529060405281600381518110617fb957617fb961ca52565b60200260200101819052506000617fcf82618efc565b602080820151604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000081850190815282518084018452600080825290860152825180840190935290518252928101929092529192506180609060408051808201825260008082526020918201528151808301909252845182528085019082015290619b61565b61807557856040516020016172f4919061cb2e565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d90156180cf565b511590565b6182435782602001511561818b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401617334565b8260c0015115618243576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401617334565b6040805160ff8082526120008201909252600091816020015b606081526020019060019003908161825c57905050905060006040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806182b79061cbbf565b935060ff16815181106182cc576182cc61ca52565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e370000000000000000000000000000000000000081525060405160200161831d919061cbde565b6040516020818303038152906040528282806183389061cbbf565b935060ff168151811061834d5761834d61ca52565b60200260200101819052506040518060400160405280600681526020017f6465706c6f79000000000000000000000000000000000000000000000000000081525082828061839a9061cbbf565b935060ff16815181106183af576183af61ca52565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806183fc9061cbbf565b935060ff16815181106184115761841161ca52565b6020026020010181905250876020015182828061842d9061cbbf565b935060ff16815181106184425761844261ca52565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163745061746800000000000000000000000000000000000081525082828061848f9061cbbf565b935060ff16815181106184a4576184a461ca52565b6020908102919091010152875182826184bc8161cbbf565b935060ff16815181106184d1576184d161ca52565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e4964000000000000000000000000000000000000000000000081525082828061851e9061cbbf565b935060ff16815181106185335761853361ca52565b602002602001018190525061854746619bc2565b82826185528161cbbf565b935060ff16815181106185675761856761ca52565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806185b49061cbbf565b935060ff16815181106185c9576185c961ca52565b6020026020010181905250868282806185e19061cbbf565b935060ff16815181106185f6576185f661ca52565b602090810291909101015285511561871d5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826186478161cbbf565b935060ff168151811061865c5761865c61ca52565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906186ac90899060040161c499565b600060405180830381865afa1580156186c9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526186f1919081019061c6ec565b82826186fc8161cbbf565b935060ff16815181106187115761871161ca52565b60200260200101819052505b8460200151156187ed5760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826187668161cbbf565b935060ff168151811061877b5761877b61ca52565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806187c89061cbbf565b935060ff16815181106187dd576187dd61ca52565b60200260200101819052506189b4565b6188256180ca8660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6188b85760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826188688161cbbf565b935060ff168151811061887d5761887d61ca52565b60200260200101819052508460a0015160405160200161889d919061ca81565b6040516020818303038152906040528282806187c89061cbbf565b8460c001511580156188fb5750604080890151815180830183526000808252602091820152825180840190935281518352908101908201526188f990511590565b155b156189b45760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261893f8161cbbf565b935060ff16815181106189545761895461ca52565b602002602001018190525061896888619c62565b604051602001618978919061ca81565b6040516020818303038152906040528282806189939061cbbf565b935060ff16815181106189a8576189a861ca52565b60200260200101819052505b604080860151815180830183526000808252602091820152825180840190935281518352908101908201526189e890511590565b618a7d5760408051808201909152600b81527f2d2d72656c61796572496400000000000000000000000000000000000000000060208201528282618a2b8161cbbf565b935060ff1681518110618a4057618a4061ca52565b60200260200101819052508460400151828280618a5c9061cbbf565b935060ff1681518110618a7157618a7161ca52565b60200260200101819052505b606085015115618b9e5760408051808201909152600681527f2d2d73616c74000000000000000000000000000000000000000000000000000060208201528282618ac68161cbbf565b935060ff1681518110618adb57618adb61ca52565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015618b4a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618b72919081019061c6ec565b8282618b7d8161cbbf565b935060ff1681518110618b9257618b9261ca52565b60200260200101819052505b60e08501515115618c455760408051808201909152600a81527f2d2d6761734c696d69740000000000000000000000000000000000000000000060208201528282618be88161cbbf565b935060ff1681518110618bfd57618bfd61ca52565b6020026020010181905250618c198560e0015160000151619bc2565b8282618c248161cbbf565b935060ff1681518110618c3957618c3961ca52565b60200260200101819052505b60e08501516020015115618cef5760408051808201909152600a81527f2d2d67617350726963650000000000000000000000000000000000000000000060208201528282618c928161cbbf565b935060ff1681518110618ca757618ca761ca52565b6020026020010181905250618cc38560e0015160200151619bc2565b8282618cce8161cbbf565b935060ff1681518110618ce357618ce361ca52565b60200260200101819052505b60e08501516040015115618d995760408051808201909152600e81527f2d2d6d617846656550657247617300000000000000000000000000000000000060208201528282618d3c8161cbbf565b935060ff1681518110618d5157618d5161ca52565b6020026020010181905250618d6d8560e0015160400151619bc2565b8282618d788161cbbf565b935060ff1681518110618d8d57618d8d61ca52565b60200260200101819052505b60e08501516060015115618e435760408051808201909152601681527f2d2d6d61785072696f726974794665655065724761730000000000000000000060208201528282618de68161cbbf565b935060ff1681518110618dfb57618dfb61ca52565b6020026020010181905250618e178560e0015160600151619bc2565b8282618e228161cbbf565b935060ff1681518110618e3757618e3761ca52565b60200260200101819052505b60008160ff1667ffffffffffffffff811115618e6157618e6161bf51565b604051908082528060200260200182016040528015618e9457816020015b6060815260200190600190039081618e7f5790505b50905060005b8260ff168160ff161015618eed57838160ff1681518110618ebd57618ebd61ca52565b6020026020010151828260ff1681518110618eda57618eda61ca52565b6020908102919091010152600101618e9a565b5093505050505b949350505050565b618f236040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c91618fa99186910161cc49565b600060405180830381865afa158015618fc6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618fee919081019061c6ec565b90506000618ffc868361a751565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b815260040161902c919061ba8b565b6000604051808303816000875af115801561904b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619073919081019061cc90565b805190915060030b1580159061908c5750602081015151155b801561909b5750604081015151155b1561807557816000815181106190b3576190b361ca52565b60200260200101516040516020016172f4919061cd46565b606060006191008560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925286518252808701908201529091506191379082905b9061a8a6565b156192945760006191b4826191ae846191a861917a8a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061a8cd565b9061a92f565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061921890829061a8a6565b1561928257604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261927f905b829061a9b4565b90505b61928b8161a9da565b92505050617108565b82156192ad5784846040516020016172f492919061cf32565b5050604080516020810190915260008152617108565b509392505050565b6000808251602084016000f09392505050565b6178f0828260016193f6565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fad3cb1cc00000000000000000000000000000000000000000000000000000000179052905160609160009182916001600160a01b0386169161935f919061cfd9565b600060405180830381855afa9150503d806000811461939a576040519150601f19603f3d011682016040523d82523d6000602084013e61939f565b606091505b50915091508180156193b2575060208151115b156193cb5780806020019051810190618ef4919061c6ec565b505060408051602081019091526000815292915050565b60006193ee838361aa43565b159392505050565b8160a001511561940557505050565b600061941284848461ab1e565b9050600061941f82618efc565b602081015181519192509060030b1580156194bb5750604080518082018252600781527f5355434345535300000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526194bb90604080518082018252600080825260209182015281518083019092528451825280850190820152619131565b156194c857505050505050565b604082015151156194e85781604001516040516020016172f4919061cff5565b806040516020016172f4919061d053565b6060600061952e8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150619593905b8290619b61565b1561960257604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617108906195fd90839061b0b9565b61a9da565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619664905b829061b143565b60010361973157604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526196ca90619278565b50604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617108906195fd905b839061a9b4565b604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526197909061958c565b156198c757604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201819052845180860190955292518452830152906197f890839061b1dd565b90506000816001835161980b919061bbc8565b8151811061981b5761981b61ca52565b602002602001015190506198be6195fd6198916040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925285518252808601908201529061b0b9565b95945050505050565b826040516020016172f4919061d0be565b50919050565b606060006199138360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506199759061958c565b15619983576171088161a9da565b604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526199e29061965d565b600103619a4c57604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617108906195fd9061972a565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619aab9061958c565b156198c757604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290619b1390839061b1dd565b9050600181511115619b4f578060028251619b2e919061bbc8565b81518110619b3e57619b3e61ca52565b602002602001015192505050919050565b50826040516020016172f4919061d0be565b805182516000911115619b7657506000617021565b81518351602085015160009291619b8c9161c359565b619b96919061bbc8565b905082602001518103619bad576001915050617021565b82516020840151819020912014905092915050565b60606000619bcf8361b282565b600101905060008167ffffffffffffffff811115619bef57619bef61bf51565b6040519080825280601f01601f191660200182016040528015619c19576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084619c2357509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e5345440000000000000000000000000000000000000000000081840190815285518087018752838152840192909252845180860190955251845290830152606091619cee905b82906193e2565b15619d2e57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e7365000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619d8d90619ce7565b15619dcd57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d4954000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619e2c90619ce7565b15619e6c57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c79000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619ecb90619ce7565b80619f305750604080518082018252601081527f47504c2d322e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619f3090619ce7565b15619f7057505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c79000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619fcf90619ce7565b8061a0345750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a03490619ce7565b1561a07457505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a0d390619ce7565b8061a1385750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a13890619ce7565b1561a17857505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a1d790619ce7565b8061a23c5750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a23c90619ce7565b1561a27c57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a2db90619ce7565b1561a31b57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a37a90619ce7565b1561a3ba57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a41990619ce7565b1561a45957505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a4b890619ce7565b1561a4f857505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a55790619ce7565b1561a59757505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a5f690619ce7565b8061a65b5750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a65b90619ce7565b1561a69b57505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a6fa90619ce7565b1561a73a57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516172f4929060200161d19c565b60608060005b845181101561a7dc578185828151811061a7735761a77361ca52565b602002602001015160405160200161a78c92919061c59a565b60405160208183030381529060405291506001855161a7ab919061bbc8565b811461a7d4578160405160200161a7c2919061d305565b60405160208183030381529060405291505b60010161a757565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161a7f5579050509050838160008151811061a8205761a82061ca52565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061a8745761a87461ca52565b6020026020010181905250818160028151811061a8935761a89361ca52565b6020908102919091010152949350505050565b602080830151835183519284015160009361a8c4929184919061b364565b14159392505050565b6040805180820190915260008082526020820152600061a8ff846000015185602001518560000151866020015161b475565b905083602001518161a911919061bbc8565b8451859061a92090839061bbc8565b90525060208401525090919050565b604080518082019091526000808252602082015281518351101561a954575081617021565b602080830151908401516001911461a97b5750815160208481015190840151829020919020145b801561a9ac5782518451859061a99290839061bbc8565b905250825160208501805161a9a890839061c359565b9052505b509192915050565b604080518082019091526000808252602082015261a9d383838361b595565b5092915050565b60606000826000015167ffffffffffffffff81111561a9fb5761a9fb61bf51565b6040519080825280601f01601f19166020018201604052801561aa25576020820181803683370190505b509050600060208201905061a9d3818560200151866000015161b640565b815181516000919081111561aa56575081515b6020808501519084015160005b8381101561ab0f578251825180821461aadf57600019602087101561aabe5760018461aa9089602061bbc8565b61aa9a919061c359565b61aaa590600861d346565b61aab090600261d444565b61aaba919061bbc8565b1990505b818116838216818103911461aadc5797506170219650505050505050565b50505b61aaea60208661c359565b945061aaf760208561c359565b9350505060208161ab08919061c359565b905061aa63565b5084518651618075919061d450565b6060600061ab2a6178f4565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161ab4757905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061aba29061cbbf565b935060ff168151811061abb75761abb761ca52565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161ac08919061d470565b60405160208183030381529060405282828061ac239061cbbf565b935060ff168151811061ac385761ac3861ca52565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061ac859061cbbf565b935060ff168151811061ac9a5761ac9a61ca52565b60200260200101819052508260405160200161acb6919061caed565b60405160208183030381529060405282828061acd19061cbbf565b935060ff168151811061ace65761ace661ca52565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061ad339061cbbf565b935060ff168151811061ad485761ad4861ca52565b602002602001018190525061ad5d878461b6ba565b828261ad688161cbbf565b935060ff168151811061ad7d5761ad7d61ca52565b60209081029190910101528551511561ae295760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261adcf8161cbbf565b935060ff168151811061ade45761ade461ca52565b602002602001018190525061adfd86600001518461b6ba565b828261ae088161cbbf565b935060ff168151811061ae1d5761ae1d61ca52565b60200260200101819052505b85608001511561ae975760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261ae728161cbbf565b935060ff168151811061ae875761ae8761ca52565b602002602001018190525061aefd565b841561aefd5760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261aedc8161cbbf565b935060ff168151811061aef15761aef161ca52565b60200260200101819052505b6040860151511561af995760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261af478161cbbf565b935060ff168151811061af5c5761af5c61ca52565b6020026020010181905250856040015182828061af789061cbbf565b935060ff168151811061af8d5761af8d61ca52565b60200260200101819052505b85606001511561b0035760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261afe28161cbbf565b935060ff168151811061aff75761aff761ca52565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561b0215761b02161bf51565b60405190808252806020026020018201604052801561b05457816020015b606081526020019060019003908161b03f5790505b50905060005b8260ff168160ff16101561b0ad57838160ff168151811061b07d5761b07d61ca52565b6020026020010151828260ff168151811061b09a5761b09a61ca52565b602090810291909101015260010161b05a565b50979650505050505050565b604080518082019091526000808252602082015281518351101561b0de575081617021565b8151835160208501516000929161b0f49161c359565b61b0fe919061bbc8565b6020840151909150600190821461b11f575082516020840151819020908220145b801561b13a5783518551869061b13690839061bbc8565b9052505b50929392505050565b600080826000015161b167856000015186602001518660000151876020015161b475565b61b171919061c359565b90505b8351602085015161b185919061c359565b811161a9d3578161b1958161d4b5565b925050826000015161b1cc85602001518361b1b0919061bbc8565b865161b1bc919061bbc8565b838660000151876020015161b475565b61b1d6919061c359565b905061b174565b6060600061b1eb848461b143565b61b1f690600161c359565b67ffffffffffffffff81111561b20e5761b20e61bf51565b60405190808252806020026020018201604052801561b24157816020015b606081526020019060019003908161b22c5790505b50905060005b81518110156192c35761b25d6195fd868661a9b4565b82828151811061b26f5761b26f61ca52565b602090810291909101015260010161b247565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061b2cb577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061b2f7576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061b31557662386f26fc10000830492506010015b6305f5e100831061b32d576305f5e100830492506008015b612710831061b34157612710830492506004015b6064831061b353576064830492506002015b600a83106170215760010192915050565b60008085841161b46b576020841161b417576000841561b3af57600161b38b86602061bbc8565b61b39690600861d346565b61b3a190600261d444565b61b3ab919061bbc8565b1990505b835181168561b3be898961c359565b61b3c8919061bbc8565b805190935082165b81811461b4025787841161b3ea5787945050505050618ef4565b8361b3f48161d4cf565b94505082845116905061b3d0565b61b40c878561c359565b945050505050618ef4565b83832061b424858861bbc8565b61b42e908761c359565b91505b85821061b4695784822080820361b4565761b44c868461c359565b9350505050618ef4565b61b46160018461bbc8565b92505061b431565b505b5092949350505050565b6000838186851161b580576020851161b52f576000851561b4c157600161b49d87602061bbc8565b61b4a890600861d346565b61b4b390600261d444565b61b4bd919061bbc8565b1990505b8451811660008761b4d28b8b61c359565b61b4dc919061bbc8565b855190915083165b82811461b5215781861061b5095761b4fc8b8b61c359565b9650505050505050618ef4565b8561b5138161d4b5565b96505083865116905061b4e4565b859650505050505050618ef4565b508383206000905b61b541868961bbc8565b821161b57e5785832080820361b55d5783945050505050618ef4565b61b56860018561c359565b935050818061b5769061d4b5565b92505061b537565b505b61b58a878761c359565b979650505050505050565b6040805180820190915260008082526020820152600061b5c7856000015186602001518660000151876020015161b475565b60208087018051918601919091525190915061b5e3908261bbc8565b83528451602086015161b5f6919061c359565b810361b605576000855261b637565b8351835161b613919061c359565b8551869061b62290839061bbc8565b905250835161b631908261c359565b60208601525b50909392505050565b6020811061b678578151835261b65760208461c359565b925061b66460208361c359565b915061b67160208261bbc8565b905061b640565b600019811561b6a757600161b68e83602061bbc8565b61b69a9061010061d444565b61b6a4919061bbc8565b90505b9151835183169219169190911790915250565b6060600061b6c884846179c7565b805160208083015160405193945061b6e29390910161d4e6565b60405160208183030381529060405291505092915050565b610b67806200d53f83390190565b6107b6806200e0a683390190565b61106f806200e85c83390190565b612072806200f8cb83390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161b77561b77a565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161b7756040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561b82c5783516001600160a01b031683526020938401939092019160010161b805565b509095945050505050565b60005b8381101561b85257818101518382015260200161b83a565b50506000910152565b6000815180845261b87381602086016020860161b837565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b983577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561b969577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261b95384865161b85b565b602095860195909450929092019160010161b919565b50919750505060209485019492909201915060010161b8af565b50929695505050505050565b600081518084526020840193506020830160005b8281101561b9e35781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161b9a3565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b983577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261ba59604088018261b85b565b905060208201519150868103602088015261ba74818361b98f565b96505050602093840193919091019060010161ba15565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b983577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261baed85835161b85b565b9450602093840193919091019060010161bab3565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b983577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261bb83604087018261b98f565b955050602093840193919091019060010161bb2a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156170215761702161bb99565b600181811c9082168061bbef57607f821691505b6020821081036198d8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600081546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160a060608501526000815461bc728161bbdb565b8060a0880152600182166000811461bc91576001811461bccb5761bcff565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660c089015260c082151560051b890101935061bcff565b84600052602060002060005b8381101561bcf65781548a820160c0015260019091019060200161bcd7565b890160c0019450505b50505060038401546080860152809250505092915050565b60c08152600061bd2a60c083018861b85b565b6001600160a01b0387166020840152828103604084015261bd4b818761b85b565b85546060850152600186015460ff1615156080850152905082810360a084015261bd75818561bc28565b98975050505050505050565b60006020828403121561bd9357600080fd5b81516001600160a01b038116811461710857600080fd5b6003811061bde1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff8816604083015286606083015261be75608083018761bdaa565b8460a08301526001600160a01b03841660c083015261b58a60e08301846001600160a01b03169052565b610100815260066101008201527f534f4c414e410000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f534f4c000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff8816604083015286606083015261be75608083018761bdaa565b60006020828403121561bf4157600080fd5b8151801515811461710857600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561364757806000526020600020601f840160051c8101602085101561bfa75750805b601f840160051c820191505b81811015610715576000815560010161bfb3565b815167ffffffffffffffff81111561bfe15761bfe161bf51565b61bff58161bfef845461bbdb565b8461bf80565b6020601f82116001811461c029576000831561c0115750848201515b600019600385901b1c1916600184901b178455610715565b600084815260208120601f198516915b8281101561c059578785015182556020948501946001909201910161c039565b508482101561c0775786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60e08152600061c09960e083018961b85b565b8760208401526001600160a01b0387166040840152828103606084015261c0c0818761b85b565b85546080850152600186015460ff16151560a085015290505b82810360c084015261c0eb818561bc28565b9998505050505050505050565b60006020828403121561c10a57600080fd5b5051919050565b60e08152600061c12460e083018961b85b565b8760208401526001600160a01b0387166040840152828103606084015261c14b818761b85b565b855160808501526020860151151560a0850152905061c0d9565b6101208152600061c17a61012083018b61b85b565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a084015261c1ad818761b85b565b855460c0850152600186015460ff16151560e085015290505b828103610100840152617414818561bc28565b60e08152600061c1ec60e083018961b85b565b876020840152866040840152828103606084015261c0c0818761b85b565b6101208152600061c21f61012083018b61b85b565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a084015261c252818761b85b565b855160c08501526020860151151560e0850152905061c1c6565b6101208152600061c28161012083018a61b85b565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a08501526000825261c2c860c0850187805182526020908101511515910152565b602081016101008501525061c2e0602082018561bc28565b9a9950505050505050505050565b60808152600061c301608083018761b85b565b8560208401526001600160a01b0385166040840152828103606084015261b58a818561bc28565b60808152600061c33b608083018761b85b565b856020840152846040840152828103606084015261b58a818561bc28565b808201808211156170215761702161bb99565b60008261c3a2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60a08152600061c3ba60a083018761b85b565b828103602084015261c3cc818761b85b565b85546040850152600186015460ff16151560608501529050828103608084015261b58a818561bc28565b6001600160a01b0383168152604060208201526000618ef4604083018461b85b565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161c45081601a85016020880161b837565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161c48d81601c84016020880161b837565b01601c01949350505050565b602081526000617108602083018461b85b565b6040516060810167ffffffffffffffff8111828210171561c4cf5761c4cf61bf51565b60405290565b60008067ffffffffffffffff84111561c4f05761c4f061bf51565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561c51f5761c51f61bf51565b60405283815290508082840185101561c53757600080fd5b6192c384602083018561b837565b600082601f83011261c55657600080fd5b6171088383516020850161c4d5565b60006020828403121561c57757600080fd5b815167ffffffffffffffff81111561c58e57600080fd5b61701d8482850161c545565b6000835161c5ac81846020880161b837565b83519083019061c5c081836020880161b837565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161c60181601a85016020880161b837565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161c63e81603384016020880161b837565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b6001600160a01b03841681526001600160a01b03831660208201526060604082015260006198be606083018461b85b565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000617108608083018461b85b565b60006020828403121561c6fe57600080fd5b815167ffffffffffffffff81111561c71557600080fd5b8201601f8101841361c72657600080fd5b61701d8482516020840161c4d5565b6000855161c747818460208a0161b837565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161c781816001840160208a0161b837565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161c7bf81600284016020890161b837565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161c80181600284016020880161b837565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061c84c604083018461b85b565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161c8c381601f85016020870161b837565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061c930604083018461b85b565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061c982604083018461b85b565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161c9f981601485016020870161b837565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061ca40604083018561b85b565b8281036020840152617104818561b85b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161cab981600185016020870161b837565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161caff81846020870161b837565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161cbb281604b85016020870161b837565b91909101604b0192915050565b600060ff821660ff810361cbd55761cbd561bb99565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161cc3c81602985016020870161b837565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000617108608083018461b85b565b60006020828403121561cca257600080fd5b815167ffffffffffffffff81111561ccb957600080fd5b82016060818503121561cccb57600080fd5b61ccd361c4ac565b81518060030b811461cce457600080fd5b8152602082015167ffffffffffffffff81111561cd0057600080fd5b61cd0c8682850161c545565b602083015250604082015167ffffffffffffffff81111561cd2c57600080fd5b61cd388682850161c545565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161cda481602185016020870161b837565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161cf9081602185016020880161b837565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161cfcd81602e84016020880161b837565b01602e01949350505050565b6000825161cfeb81846020870161b837565b9190910192915050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161cc3c81602985016020870161b837565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161d0b181602285016020870161b837565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161d0f681600e85016020870161b837565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161d1d481601885016020880161b837565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161d21181601c84016020880161b837565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161d31781846020870161b837565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b80820281158282048414176170215761702161bb99565b6001815b600184111561d3985780850481111561d37c5761d37c61bb99565b600184161561d38a57908102905b60019390931c92800261d361565b935093915050565b60008261d3af57506001617021565b8161d3bc57506000617021565b816001811461d3d2576002811461d3dc5761d3f8565b6001915050617021565b60ff84111561d3ed5761d3ed61bb99565b50506001821b617021565b5060208310610133831016604e8410600b841016171561d41b575081810a617021565b61d428600019848461d35d565b806000190482111561d43c5761d43c61bb99565b029392505050565b6000617108838361d3a0565b818103600083128015838313168383128216171561a9d35761a9d361bb99565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161d4a881601c85016020870161b837565b91909101601c0192915050565b6000600019820361d4c85761d4c861bb99565b5060010190565b60008161d4de5761d4de61bb99565b506000190190565b6000835161d4f881846020880161b837565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161d53281600184016020880161b837565b0160010194935050505056fe60c0604052600d60809081526c2bb930b83832b21022ba3432b960991b60a05260009061002c9082610114565b506040805180820190915260048152630ae8aa8960e31b60208201526001906100559082610114565b506002805460ff1916601217905534801561006f57600080fd5b506101d2565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061009f57607f821691505b6020821081036100bf57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561010f57806000526020600020601f840160051c810160208510156100ec5750805b601f840160051c820191505b8181101561010c57600081556001016100f8565b50505b505050565b81516001600160401b0381111561012d5761012d610075565b6101418161013b845461008b565b846100c5565b6020601f821160018114610175576000831561015d5750848201515b600019600385901b1c1916600184901b17845561010c565b600084815260208120601f198516915b828110156101a55787850151825560209485019460019092019101610185565b50848210156101c35786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610986806101e16000396000f3fe6080604052600436106100c05760003560e01c8063313ce56711610074578063a9059cbb1161004e578063a9059cbb146101fa578063d0e30db01461021a578063dd62ed3e1461022257600080fd5b8063313ce5671461018c57806370a08231146101b857806395d89b41146101e557600080fd5b806318160ddd116100a557806318160ddd1461012f57806323b872dd1461014c5780632e1a7d4d1461016c57600080fd5b806306fdde03146100d4578063095ea7b3146100ff57600080fd5b366100cf576100cd61025a565b005b600080fd5b3480156100e057600080fd5b506100e96102b5565b6040516100f69190610745565b60405180910390f35b34801561010b57600080fd5b5061011f61011a3660046107da565b610343565b60405190151581526020016100f6565b34801561013b57600080fd5b50475b6040519081526020016100f6565b34801561015857600080fd5b5061011f610167366004610804565b6103bd565b34801561017857600080fd5b506100cd610187366004610841565b610647565b34801561019857600080fd5b506002546101a69060ff1681565b60405160ff90911681526020016100f6565b3480156101c457600080fd5b5061013e6101d336600461085a565b60036020526000908152604090205481565b3480156101f157600080fd5b506100e9610724565b34801561020657600080fd5b5061011f6102153660046107da565b610731565b6100cd61025a565b34801561022e57600080fd5b5061013e61023d366004610875565b600460209081526000928352604080842090915290825290205481565b33600090815260036020526040812080543492906102799084906108d7565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b600080546102c2906108ea565b80601f01602080910402602001604051908101604052809291908181526020018280546102ee906108ea565b801561033b5780601f106103105761010080835404028352916020019161033b565b820191906000526020600020905b81548152906001019060200180831161031e57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103ab9086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604081205482111561042b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841633148015906104a1575073ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156105605773ffffffffffffffffffffffffffffffffffffffff8416600090815260046020908152604080832033845290915290205482111561051a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b73ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091528120805484929061055a90849061093d565b90915550505b73ffffffffffffffffffffffffffffffffffffffff84166000908152600360205260408120805484929061059590849061093d565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040812080548492906105cf9084906108d7565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161063591815260200190565b60405180910390a35060019392505050565b3360009081526003602052604090205481111561069a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b33600090815260036020526040812080548392906106b990849061093d565b9091555050604051339082156108fc029083906000818181858888f193505050501580156106eb573d6000803e3d6000fd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b600180546102c2906108ea565b600061073e3384846103bd565b9392505050565b602081526000825180602084015260005b818110156107735760208186018101516040868401015201610756565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146107d557600080fd5b919050565b600080604083850312156107ed57600080fd5b6107f6836107b1565b946020939093013593505050565b60008060006060848603121561081957600080fd5b610822846107b1565b9250610830602085016107b1565b929592945050506040919091013590565b60006020828403121561085357600080fd5b5035919050565b60006020828403121561086c57600080fd5b61073e826107b1565b6000806040838503121561088857600080fd5b610891836107b1565b915061089f602084016107b1565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156103b7576103b76108a8565b600181811c908216806108fe57607f821691505b602082108103610937577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b818103818111156103b7576103b76108a856fea2646970667358221220bb5b9dcef0ba90bcdefcbd63f71b1df95b50e29550a7456c69c6b9ff9dcdd20e64736f6c634300081a00336080604052348015600f57600080fd5b506107978061001f6000396000f3fe6080604052600436106100355760003560e01c80632d4cfb7e1461003e5780635bcfd6161461005e578063c9028a361461007e57005b3661003c57005b005b34801561004a57600080fd5b5061003c610059366004610182565b61009e565b34801561006a57600080fd5b5061003c6100793660046101ed565b6100d8565b34801561008a57600080fd5b5061003c6100993660046102aa565b610153565b7f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7816040516100cd9190610399565b60405180910390a150565b606081156100ef576100ec8284018461049f565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e61011a8780610595565b61012a60408a0160208b016105fa565b8960400135338660405161014396959493929190610615565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4816040516100cd91906106d7565b60006020828403121561019457600080fd5b813567ffffffffffffffff8111156101ab57600080fd5b820160c081850312156101bd57600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146101e857600080fd5b919050565b60008060008060006080868803121561020557600080fd5b853567ffffffffffffffff81111561021c57600080fd5b86016060818903121561022e57600080fd5b945061023c602087016101c4565b935060408601359250606086013567ffffffffffffffff81111561025f57600080fd5b8601601f8101881361027057600080fd5b803567ffffffffffffffff81111561028757600080fd5b88602082840101111561029957600080fd5b959894975092955050506020019190565b6000602082840312156102bc57600080fd5b813567ffffffffffffffff8111156102d357600080fd5b8201608081850312156101bd57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261031a57600080fd5b830160208101925035905067ffffffffffffffff81111561033a57600080fd5b80360382131561034957600080fd5b9250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6020815260006103a983846102e5565b60c060208501526103be60e085018284610350565b91505073ffffffffffffffffffffffffffffffffffffffff6103e2602086016101c4565b16604084015260006040850135905080606085015250606084013580151580821461040c57600080fd5b80608086015250506000608085013590508060a08501525061043160a08501856102e5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c0860152610466838284610350565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156104b157600080fd5b813567ffffffffffffffff8111156104c857600080fd5b8201601f810184136104d957600080fd5b803567ffffffffffffffff8111156104f3576104f3610470565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561055f5761055f610470565b60405281815282820160200186101561057757600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126105ca57600080fd5b83018035915067ffffffffffffffff8211156105e557600080fd5b60200191503681900382131561034957600080fd5b60006020828403121561060c57600080fd5b6101bd826101c4565b60a08152600061062960a08301888a610350565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b8181101561069357602081870181015184830182015201610677565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff6106f9836101c4565b16602082015273ffffffffffffffffffffffffffffffffffffffff610720602084016101c4565b166040820152600080604084013590508060608401525061074460608401846102e5565b60808085015261075860a085018284610350565b9594505050505056fea2646970667358221220a0b5f72856c3077829eb97a64d61e91f41a64976665417c3e8e8cf442695953064736f6c634300081a003360c060405234801561001057600080fd5b5060405161106f38038061106f83398101604081905261002f916100db565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006357604051632b2add3d60e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0385811691909117909155828116608052811660a0526040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac590600090a150505061011e565b80516001600160a01b03811681146100d657600080fd5b919050565b6000806000606084860312156100f057600080fd5b6100f9846100bf565b9250610107602085016100bf565b9150610115604085016100bf565b90509250925092565b60805160a051610f2561014a60003960006101e50152600081816102b9015261045b0152610f256000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806397770dff11610097578063c63585cc11610066578063c63585cc14610273578063d7fd7afb14610286578063d936a012146102b4578063ee2815ba146102db57600080fd5b806397770dff1461021a578063a7cb05071461022d578063c39aca3714610240578063c62178ac1461025357600080fd5b8063513a9c05116100d3578063513a9c051461018a578063569541b9146101c0578063842da36d146101e057806391dd645f1461020757600080fd5b80630be15547146100fa5780631f0e251b1461015a5780633ce4a5bc1461016f575b600080fd5b610130610108366004610bd1565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61016d610168366004610c13565b6102ee565b005b61013073735b14bb79463307aacbed86daf3322b1e6226ab81565b610130610198366004610bd1565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101309073ffffffffffffffffffffffffffffffffffffffff1681565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d610215366004610c35565b610402565b61016d610228366004610c13565b610526565b61016d61023b366004610c61565b610633565b61016d61024e366004610c83565b6106ce565b6004546101309073ffffffffffffffffffffffffffffffffffffffff1681565b610130610281366004610d53565b6108cd565b6102a6610294366004610bd1565b60006020819052908152604090205481565b604051908152602001610151565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d6102e9366004610c35565b610a02565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461033b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610388576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461044f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354600090610497907f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff16846108cd565b60008481526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251878152918201529192507f0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e910160405180910390a1505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610573576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166105c0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e906020016103f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610680576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461071b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831673735b14bb79463307aacbed86daf3322b1e6226ab1480610768575073ffffffffffffffffffffffffffffffffffffffff831630145b1561079f576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690528616906347e7ef24906044016020604051808303816000875af1158015610814573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108389190610d96565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063de43156e906108939089908990899088908890600401610e01565b600060405180830381600087803b1580156108ad57600080fd5b505af11580156108c1573d6000803e3d6000fd5b50505050505050505050565b60008060006108dc8585610ad3565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b16603482015291935091508690604801604051602081830303815290604052805190602001206040516020016109c29291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a4f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d91016106c2565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b3b576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610b75578284610b78565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610bca576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b600060208284031215610be357600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c0e57600080fd5b919050565b600060208284031215610c2557600080fd5b610c2e82610bea565b9392505050565b60008060408385031215610c4857600080fd5b82359150610c5860208401610bea565b90509250929050565b60008060408385031215610c7457600080fd5b50508035926020909101359150565b60008060008060008060a08789031215610c9c57600080fd5b863567ffffffffffffffff811115610cb357600080fd5b87016060818a031215610cc557600080fd5b9550610cd360208801610bea565b945060408701359350610ce860608801610bea565b9250608087013567ffffffffffffffff811115610d0457600080fd5b8701601f81018913610d1557600080fd5b803567ffffffffffffffff811115610d2c57600080fd5b896020828401011115610d3e57600080fd5b60208201935080925050509295509295509295565b600080600060608486031215610d6857600080fd5b610d7184610bea565b9250610d7f60208501610bea565b9150610d8d60408501610bea565b90509250925092565b600060208284031215610da857600080fd5b81518015158114610c2e57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60808152600086357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018112610e3957600080fd5b870160208101903567ffffffffffffffff811115610e5657600080fd5b803603821315610e6557600080fd5b60606080850152610e7a60e085018284610db8565b91505073ffffffffffffffffffffffffffffffffffffffff610e9e60208a01610bea565b1660a0840152604088013560c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401528281036060840152610ee3818587610db8565b9897505050505050505056fea26469706673582212207dafbdd8854199435f8bb1e1f1029e7cbe3dd928388cb4e978d79018e894847164736f6c634300081a003360c060405234801561001057600080fd5b5060405161207238038061207283398101604081905261002f916101f0565b6001600160a01b038216158061004c57506001600160a01b038116155b1561006a5760405163d92e233d60e01b815260040160405180910390fd5b60066100768982610342565b5060076100838882610342565b506008805460ff191660ff881617905560808590528360028111156100aa576100aa610400565b60a08160028111156100be576100be610400565b905250600192909255600080546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506104169350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261013357600080fd5b81516001600160401b0381111561014c5761014c61010c565b604051601f8201601f19908116603f011681016001600160401b038111828210171561017a5761017a61010c565b60405281815283820160200185101561019257600080fd5b60005b828110156101b157602081860181015183830182015201610195565b506000918101602001919091529392505050565b8051600381106101d457600080fd5b919050565b80516001600160a01b03811681146101d457600080fd5b600080600080600080600080610100898b03121561020d57600080fd5b88516001600160401b0381111561022357600080fd5b61022f8b828c01610122565b60208b015190995090506001600160401b0381111561024d57600080fd5b6102598b828c01610122565b975050604089015160ff8116811461027057600080fd5b60608a0151909650945061028660808a016101c5565b60a08a0151909450925061029c60c08a016101d9565b91506102aa60e08a016101d9565b90509295985092959890939650565b600181811c908216806102cd57607f821691505b6020821081036102ed57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561033d57806000526020600020601f840160051c8101602085101561031a5750805b601f840160051c820191505b8181101561033a5760008155600101610326565b50505b505050565b81516001600160401b0381111561035b5761035b61010c565b61036f8161036984546102b9565b846102f3565b6020601f8211600181146103a3576000831561038b5750848201515b600019600385901b1c1916600184901b17845561033a565b600084815260208120601f198516915b828110156103d357878501518255602094850194600190920191016103b3565b50848210156103f15786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a051611c1b61045760003960006103440152600081816102f001528181610bdc01528181610ce201528181610efe01526110040152611c1b6000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806395d89b41116100f9578063ccc7759911610097578063eddeb12311610071578063eddeb12314610461578063f2441b3214610474578063f687d12a14610494578063fc5fecd5146104a757600080fd5b8063ccc77599146103d4578063d9eeebed146103e7578063dd62ed3e1461041b57600080fd5b8063b84c8246116100d3578063b84c824614610386578063c47f00271461039b578063c7012626146103ae578063c835d7cc146103c157600080fd5b806395d89b4114610337578063a3413d031461033f578063a9059cbb1461037357600080fd5b80633ce4a5bc116101665780634d8943bb116101405780634d8943bb146102ac57806370a08231146102b557806385e1f4d0146102eb5780638b851b951461031257600080fd5b80633ce4a5bc1461024657806342966c681461028657806347e7ef241461029957600080fd5b806318160ddd1161019757806318160ddd1461021657806323b872dd1461021e578063313ce5671461023157600080fd5b806306fdde03146101be578063091d2788146101dc578063095ea7b3146101f3575b600080fd5b6101c66104ba565b6040516101d39190611648565b60405180910390f35b6101e560015481565b6040519081526020016101d3565b610206610201366004611687565b61054c565b60405190151581526020016101d3565b6005546101e5565b61020661022c3660046116b3565b610563565b60085460405160ff90911681526020016101d3565b61026173735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6102066102943660046116f4565b6105fa565b6102066102a7366004611687565b61060e565b6101e560025481565b6101e56102c336600461170d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101e57f000000000000000000000000000000000000000000000000000000000000000081565b60085461026190610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101c6610767565b6103667f000000000000000000000000000000000000000000000000000000000000000081565b6040516101d3919061172a565b610206610381366004611687565b610776565b610399610394366004611832565b610783565b005b6103996103a9366004611832565b6107e0565b6102066103bc366004611883565b610839565b6103996103cf36600461170d565b610988565b6103996103e236600461170d565b610a9c565b6103ef610bb0565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101d3565b6101e56104293660046118dc565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b61039961046f3660046116f4565b610dce565b6000546102619073ffffffffffffffffffffffffffffffffffffffff1681565b6103996104a23660046116f4565b610e50565b6103ef6104b53660046116f4565b610ed2565b6060600680546104c990611915565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590611915565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b5050505050905090565b60006105593384846110ee565b5060015b92915050565b60006105708484846111f7565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054828110156105db576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ef85336105ea8685611997565b6110ee565b506001949350505050565b600061060633836113b2565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab1480159061064c575060005473ffffffffffffffffffffffffffffffffffffffff163314155b80156106755750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b156106ac576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106b683836114f4565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107569186906119aa565b60405180910390a250600192915050565b6060600780546104c990611915565b60006105593384846111f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107d0576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107dc8282611a1b565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461082d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107dc8282611a1b565b6000806000610846610bb0565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156108d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fc9190611b34565b610932576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61093c33856113b2565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161097591899189918791611b56565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109d5576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a22576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610ae9576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b36576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c679190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610cb6576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d699190611ba2565b905080600003610da5576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610db89190611bbb565b610dc29190611bd2565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e1b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a91565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e9d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f899190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610fd8576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015611067573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108b9190611ba2565b9050806000036110c7576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000906110d78784611bbb565b6110e19190611bd2565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661113b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611188576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611244576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611291576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040902054818110156112f1576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112fb8282611997565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260036020526040808220939093559085168152908120805484929061133e908490611bd2565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113a491815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113ff576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260409020548181101561145f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114698282611997565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812091909155600580548492906114a4908490611997565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111ea565b73ffffffffffffffffffffffffffffffffffffffff8216611541576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008282546115539190611bd2565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805483929061158d908490611bd2565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561160a576020818501810151868301820152016115ee565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061165b60208301846115e4565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461168457600080fd5b50565b6000806040838503121561169a57600080fd5b82356116a581611662565b946020939093013593505050565b6000806000606084860312156116c857600080fd5b83356116d381611662565b925060208401356116e381611662565b929592945050506040919091013590565b60006020828403121561170657600080fd5b5035919050565b60006020828403121561171f57600080fd5b813561165b81611662565b6020810160038310611765577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff8411156117b5576117b561176b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156118025761180261176b565b60405283815290508082840185101561181a57600080fd5b83836020830137600060208583010152509392505050565b60006020828403121561184457600080fd5b813567ffffffffffffffff81111561185b57600080fd5b8201601f8101841361186c57600080fd5b61187b8482356020840161179a565b949350505050565b6000806040838503121561189657600080fd5b823567ffffffffffffffff8111156118ad57600080fd5b8301601f810185136118be57600080fd5b6118cd8582356020840161179a565b95602094909401359450505050565b600080604083850312156118ef57600080fd5b82356118fa81611662565b9150602083013561190a81611662565b809150509250929050565b600181811c9082168061192957607f821691505b602082108103611962577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561055d5761055d611968565b6040815260006119bd60408301856115e4565b90508260208301529392505050565b601f821115611a1657806000526020600020601f840160051c810160208510156119f35750805b601f840160051c820191505b81811015611a1357600081556001016119ff565b50505b505050565b815167ffffffffffffffff811115611a3557611a3561176b565b611a4981611a438454611915565b846119cc565b6020601f821160018114611a9b5760008315611a655750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611a13565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611ae95787850151825560209485019460019092019101611ac9565b5084821015611b2557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b600060208284031215611b4657600080fd5b8151801515811461165b57600080fd5b608081526000611b6960808301876115e4565b6020830195909552506040810192909252606090910152919050565b600060208284031215611b9757600080fd5b815161165b81611662565b600060208284031215611bb457600080fd5b5051919050565b808202811582820484141761055d5761055d611968565b8082018082111561055d5761055d61196856fea2646970667358221220e7be4f6cbd9165fe86191f383968618f320a1c54b4d694652e1d738c851a10e064736f6c634300081a0033a264697066735822122023fb2314901fca0b8fcef422896aa5a1c4532e21705af07324eb0260bdf15df064736f6c634300081a0033",
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

// TestWithdrawAndCallZRC20SolanaFails is a paid mutator transaction binding the contract method 0x77def36e.
//
// Solidity: function testWithdrawAndCallZRC20SolanaFails() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20SolanaFails(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20SolanaFails")
}

// TestWithdrawAndCallZRC20SolanaFails is a paid mutator transaction binding the contract method 0x77def36e.
//
// Solidity: function testWithdrawAndCallZRC20SolanaFails() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20SolanaFails() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20SolanaFails(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20SolanaFails is a paid mutator transaction binding the contract method 0x77def36e.
//
// Solidity: function testWithdrawAndCallZRC20SolanaFails() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20SolanaFails() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20SolanaFails(&_GatewayZEVMInboundTest.TransactOpts)
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

// TestWithdrawZRC20Solana is a paid mutator transaction binding the contract method 0x73a3db40.
//
// Solidity: function testWithdrawZRC20Solana() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20Solana(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20Solana")
}

// TestWithdrawZRC20Solana is a paid mutator transaction binding the contract method 0x73a3db40.
//
// Solidity: function testWithdrawZRC20Solana() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20Solana() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20Solana(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20Solana is a paid mutator transaction binding the contract method 0x73a3db40.
//
// Solidity: function testWithdrawZRC20Solana() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20Solana() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20Solana(&_GatewayZEVMInboundTest.TransactOpts)
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
