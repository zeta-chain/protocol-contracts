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

// GatewayEVMInboundTestMetaData contains all meta data concerning the GatewayEVMInboundTest contract.
var GatewayEVMInboundTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ADDITIONAL_ACTION_FEE_WEI\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testAdditionalActionDisabledReverts\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallSecondActionFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallSecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfCallOnRevertIsTrue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfDestinationIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20SecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20ToCustodyFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20ToCustodyFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthToTssFailsForSubsequentActions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthWithAmountSecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthWithAmountToTssFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20SecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodySecondActionFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTss\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssFailsForSubsequentActions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayloadFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountSecondActionFailsWithOnlyFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountSecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTss\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZetaToConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testFeeSystemWithUpdatedFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testMixedActionTypesInSameTransaction\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testNispe\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpdateAdditionalActionFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpdateAdditionalActionFeeOnlyAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositedAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedAdditionalActionFee\",\"inputs\":[{\"name\":\"oldFeeWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newFeeWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AdditionalActionDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallOnRevertNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExcessETHProvided\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectValueProvided\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFee\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f8054909116909117905569021e19e0c9bab2400000602c553480156039575f80fd5b50620136fe80620000495f395ff3fe608060405234801562000010575f80fd5b5060043610620005f4575f3560e01c806384c59bcf1162000317578063b284906311620001a7578063dc23a35f11620000fb578063f1c6b4d311620000ab578063f75fc9691162000083578063f75fc969146200097f578063f8d37ef21462000989578063fa7626d414620009a3575f80fd5b8063f1c6b4d31462000961578063f2036eda146200096b578063f6e53a5d1462000975575f80fd5b8063e306a97811620000df578063e306a9781462000943578063e85c5a07146200094d578063f0a635c51462000957575f80fd5b8063dc23a35f146200092f578063e20c9f711462000939575f80fd5b8063c0fab86d1162000157578063c57926c6116200013b578063c57926c61462000911578063c7a1eccb146200091b578063d86a4a0c1462000925575f80fd5b8063c0fab86d14620008fd578063c51a4cbe1462000907575f80fd5b8063b966e8fa116200018b578063b966e8fa14620008ce578063ba414fa614620008d8578063ba46ba2314620008f3575f80fd5b8063b284906314620008ba578063b5508aa914620008c4575f80fd5b806391a5c818116200026b578063a00a6fff116200021b578063b0464fdc11620001f3578063b0464fdc146200089c578063b1409f7114620008a6578063b152693414620008b0575f80fd5b8063a00a6fff146200087e578063a0d60b3a1462000888578063aa030c1c1462000892575f80fd5b80639a34d8f3116200024f5780639a34d8f314620008605780639acda9fa146200086a5780639fd1e5971462000874575f80fd5b806391a5c818146200084c57806395cd04451462000856575f80fd5b8063895c2bc611620002c75780638be96f5c11620002ab5780638be96f5c146200081f5780639073eafe1462000829578063916a17c61462000833575f80fd5b8063895c2bc6146200080b5780638aeeb7e71462000815575f80fd5b806385f5c51c11620002fb57806385f5c51c14620007ed57806386b6e42914620007f757806388343a411462000801575f80fd5b806384c59bcf14620007ca57806385226c8114620007d4575f80fd5b8063478095e51162000493578063586da26711620003e75780636aa02e8b11620003975780637bb46d46116200036f5780637bb46d4614620007ac5780637e7dfee314620007b6578063846b9f7f14620007c0575f80fd5b80636aa02e8b146200078e5780636ab1c51614620007985780636c33bacb14620007a2575f80fd5b806363d7a41811620003cb57806363d7a41814620007615780636459542a146200076b57806366d9a9a01462000775575f80fd5b8063586da267146200074d578063598b7edb1462000757575f80fd5b80634ce25c0a116200044357806351ee53cb116200042757806351ee53cb146200072f57806353c9a0a31462000739578063545c37451462000743575f80fd5b80634ce25c0a146200071b57806351da903d1462000725575f80fd5b80634845232911620004775780634845232914620006fd57806349050a4414620007075780634a7803391462000711575f80fd5b8063478095e514620006e9578063481daadb14620006f3575f80fd5b80632ade3880116200054b5780633424914f11620004fb5780633f7286f411620004df5780633f7286f414620006cb578063458085f514620006d5578063466f332e14620006df575f80fd5b80633424914f14620006b75780633e5e3c2314620006c1575f80fd5b806330f7c04f116200052f57806330f7c04f146200069957806332fc1fad14620006a357806333ed091c14620006ad575f80fd5b80632ade388014620006765780632cf9572d146200068f575f80fd5b806315ee8f3611620005a75780631ed7831c116200058b5780631ed7831c14620006405780631f4f542f14620006625780632aa5ca5d146200066c575f80fd5b806315ee8f36146200062c5780631b906ef31462000636575f80fd5b806309a263c111620005db57806309a263c1146200060e5780630a9254e414620006185780630fc37f5e1462000622575f80fd5b806301a74aee14620005f85780630724d8e31462000604575b5f80fd5b62000602620009b1565b005b6200060262000b90565b6200060262000d46565b6200060262000f6e565b6200060262001ae9565b6200060262001d8d565b6200060262001f4b565b6200064a620022e8565b6040516200065991906200f168565b60405180910390f35b620006026200234a565b6200060262002533565b6200068062002673565b6040516200065991906200f1e3565b62000602620027bb565b6200060262002a06565b6200060262002e03565b6200060262002f65565b62000602620030e3565b6200064a620033f9565b6200064a62003459565b62000602620034b9565b6200060262003719565b6200060262003a9d565b6200060262003c50565b6200060262003f8a565b620006026200427d565b6200060262004559565b620006026200467a565b6200060262004971565b6200060262004ab3565b6200060262004bec565b6200060262004d08565b6200060262004f7a565b6200060262005137565b620006026200526e565b62000602620058d9565b6200077f62005c6c565b6040516200065991906200f34c565b6200060262005ddc565b6200060262005ef3565b62000602620064e0565b62000602620066d6565b62000602620067ad565b620006026200696c565b6200060262006a2d565b620007de62006acd565b6040516200065991906200f3ee565b6200060262006ba2565b6200060262006d25565b6200060262006e7f565b6200060262006fe8565b620006026200710a565b62000602620071de565b6200060262007330565b6200083d6200748f565b6040516200065991906200f467565b6200060262007574565b6200060262007674565b6200060262007a74565b6200060262007b92565b6200060262007c68565b6200060262007e0c565b620006026200815c565b62000602620082a8565b6200083d62008437565b620006026200851c565b6200060262008861565b62000602620089c8565b620007de62008bf5565b6200060262008cca565b620008e262008f58565b604051901515815260200162000659565b620006026200902c565b6200060262009187565b6200060262009338565b6200060262009551565b620006026200962a565b62000602620097bc565b6200060262009b63565b6200064a6200a0c2565b620006026200a122565b620006026200a3b4565b620006026200a6b0565b620006026200a7e9565b620006026200a9d4565b620006026200ab30565b620006026200ac49565b6200099462030d4081565b60405190815260200162000659565b601f54620008e29060ff1681565b6026546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a0810183526103218082525f8284018190528285019190915283519283019093528282526060810191909152919250906080810162000a3e621e848060016200f52d565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162000ada916004016200f543565b5f604051808303815f87803b15801562000af2575f80fd5b505af115801562000b05573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb4935062000b5d9290911690869086906004016200f5b3565b5f604051808303815f87803b15801562000b75575f80fd5b505af115801562000b88573d5f803e3d5ffd5b505050505050565b6027546020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039182166084820152620186a092919091163190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562000c2a575f80fd5b505af115801562000c3d573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c9062000c8d9086905f906028906200f728565b60405180910390a36020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263726ac97c92869262000cea92909116906028906004016200f75e565b5f604051808303818588803b15801562000d02575f80fd5b505af115801562000d15573d5f803e3d5ffd5b50506027546001600160a01b031631925062000d41915062000d3a905084846200f52d565b826200ace9565b505050565b6020546026546040517f282906ed000000000000000000000000000000000000000000000000000000008152620186a0926001600160a01b039081169263282906ed92859262000da092169083906028906004016200f781565b5f604051808303818588803b15801562000db8575f80fd5b505af115801562000dcb573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392507f84222ac700000000000000000000000000000000000000000000000000000000915062000e1e905062030d40856200f52d565b62000e2d62030d40866200f52d565b62000e3a9060016200f52d565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b909216825262000eac916004016200f543565b5f604051808303815f87803b15801562000ec4575f80fd5b505af115801562000ed7573d5f803e3d5ffd5b50506020546001600160a01b0316915063282906ed905062000efd62030d40846200f52d565b62000f0a9060016200f52d565b6026546040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815262000f56916001600160a01b03169086906028906004016200f781565b5f604051808303818588803b15801562000b75575f80fd5b602580547fffffffffffffffffffffffff000000000000000000000000000000000000000090811630179091556026805482166112341790556027805490911661567817905560405162000fc2906200f08d565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103905ff08015801562001045573d5f803e3d5ffd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316179055602754604051911690819062001090906200f09b565b6001600160a01b03928316815291166020820152604001604051809103905ff080158015620010c1573d5f803e3d5ffd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081178255604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260275460255492519086169481019490945260448401929092529092166064820152620011a091906084015b60408051601f198184030181529190526020810180516001600160e01b03167fc0c53b8b000000000000000000000000000000000000000000000000000000001790526200ad66565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff000000000000000000000000000000000000000090911681178255604080518082018252601081527f4552433230437573746f64792e736f6c000000000000000000000000000000009381019390935260275460255491516024810193909352841660448301529092166064830152620012729160840162001157565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602180549190920483167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055604080518082018252601a81527f5a657461436f6e6e6563746f724e6f6e4e61746976652e736f6c0000000000006020808301919091525460248054602754602554955193871692840192909252851660448301528416606482015291909216608482015262001398919060a40160408051601f198184030181529190526020810180516001600160e01b03167ff8c8765e000000000000000000000000000000000000000000000000000000001790526200ad66565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602280549190920483167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790556027546040517fca669fa700000000000000000000000000000000000000000000000000000000815291166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200146e575f80fd5b505af115801562001481573d5f803e3d5ffd5b5050602480546027546022546040517f15d57fd40000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152908216938101939093521692506315d57fd491506044015f604051808303815f87803b158015620014f2575f80fd5b505af115801562001505573d5f803e3d5ffd5b50506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d91506044015f604051808303815f87803b15801562001587575f80fd5b505af11580156200159a573d5f803e3d5ffd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d5691506024015f604051808303815f87803b1580156200160e575f80fd5b505af115801562001621573d5f803e3d5ffd5b50506020546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f91506024015f604051808303815f87803b15801562001685575f80fd5b505af115801562001698573d5f803e3d5ffd5b50506020546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef91506024015f604051808303815f87803b158015620016fc575f80fd5b505af11580156200170f573d5f803e3d5ffd5b50506021546023546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639b19251a91506024015f604051808303815f87803b15801562001773575f80fd5b505af115801562001786573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015620017e5575f80fd5b505af1158015620017f8573d5f803e3d5ffd5b5050602354602554602c546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810191909152911692506340c10f1991506044015f604051808303815f87803b15801562001867575f80fd5b505af11580156200187a573d5f803e3d5ffd5b50506027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015620018ee575f80fd5b505af115801562001901573d5f803e3d5ffd5b5050602254602554602c546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101919091525f60448201529116925063106e629091506064015f604051808303815f87803b15801562001976575f80fd5b505af115801562001989573d5f803e3d5ffd5b50506040805160a0810182526103218082525f602080840182815284860193845285519182019095528181526060840181905260808401919091528251602880549551151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009096166001600160a01b0392831617959095178555915160298054919093167fffffffffffffffffffffffff00000000000000000000000000000000000000009190911617909155909350909150602a9062001a6390826200f81f565b50608091909101516003909101556020546040517f7c74425300000000000000000000000000000000000000000000000000000000815262030d4060048201526001600160a01b0390911690637c744253906024015f604051808303815f87803b15801562001ad0575f80fd5b505af115801562001ae3573d5f803e3d5ffd5b50505050565b60208054604080516328ae864d60e21b815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562001b2d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062001b5391906200f8e7565b62001b609060016200f52d565b67ffffffffffffffff81111562001b7b5762001b7b6200f7aa565b6040519080825280601f01601f19166020018201604052801562001ba6576020820181803683370190505b50602a9062001bb690826200f81f565b505f6028600201805462001bca906200f5ea565b905090505f60205f9054906101000a90046001600160a01b03166001600160a01b031663a2ba19346040518163ffffffff1660e01b8152600401602060405180830381865afa15801562001c20573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062001c4691906200f8e7565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162001cd6916004016200f543565b5f604051808303815f87803b15801562001cee575f80fd5b505af115801562001d01573d5f803e3d5ffd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c935060019262001d599216906028906004016200f75e565b5f604051808303818588803b15801562001d71575f80fd5b505af115801562001d84573d5f803e3d5ffd5b50505050505050565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482018190529261c35092169063095ea7b3906044016020604051808303815f875af115801562001e01573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062001e2791906200f8ff565b506040515f602482015260448101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f28dceb3907f8afe4db400000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529181526020820180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b909216825262001ed9916004016200f543565b5f604051808303815f87803b15801562001ef1575f80fd5b505af115801562001f04573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b09450869362001d599381169289929116906028906004016200f920565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052602754602554915163248e63e160e11b8152600160048201819052602482018190526044820181905260648201529293506001600160a01b03908116319291163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b1580156200200e575f80fd5b505af115801562002021573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f90620020739088905f9089906028906200f958565b60405180910390a36020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263397e375c928892620020d49290911690839089906028906004016200f995565b5f604051808303818588803b158015620020ec575f80fd5b505af1158015620020ff573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063491cc7c2925060840190505f604051808303815f87803b15801562002165575f80fd5b505af115801562002178573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f90620021ca9088905f9089906028906200f958565b60405180910390a36020546001600160a01b031663397e375c620021f262030d40876200f52d565b6026546040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815262002240916001600160a01b031690899089906028906004016200f995565b5f604051808303818588803b15801562002258575f80fd5b505af11580156200226b573d5f803e3d5ffd5b50506027546025546001600160a01b0391821631945016319150620022bb905062030d406200229c8860026200f9be565b620022a890876200f52d565b620022b491906200f52d565b836200ace9565b62000b8862030d40620022d08860026200f9be565b620022dc90866200f9d8565b62000d3a91906200f9d8565b606060168054806020026020016040519081016040528092919081815260200182805480156200234057602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831162002321575b5050505050905090565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303815f875af1158015620023bd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620023e391906200f8ff565b506040805160a0810182526103218082525f602080840182905283850192909252835191820190935282815260608201526080810162002428621e848060016200f52d565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620024c4916004016200f543565b5f604051808303815f87803b158015620024dc575f80fd5b505af1158015620024ef573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b0945062000b5d93928316928892169087906004016200f9ee565b6020546040515f916001600160a01b03169062002550906200f0a9565b6001600160a01b039091168152602001604051809103905ff0801580156200257a573d5f803e3d5ffd5b506026546040516001600160a01b039091166024820152909150620186a0905f9060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b17905290506001600160a01b03831663d51240c4620025e962030d4060026200f9be565b620025f590856200f52d565b6026546040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526200263e916001600160a01b03169086906004016200fa26565b5f604051808303818588803b15801562002656575f80fd5b505af115801562002669573d5f803e3d5ffd5b5050505050505050565b6060601e805480602002602001604051908101604052809291908181526020015f905b82821015620027b2575f84815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156200279a578382905f5260205f2001805462002708906200f5ea565b80601f016020809104026020016040519081016040528092919081815260200182805462002736906200f5ea565b8015620027855780601f106200275b5761010080835404028352916020019162002785565b820191905f5260205f20905b8154815290600101906020018083116200276757829003601f168201915b505050505081526020019060010190620026e8565b50505050815250508152602001906001019062002696565b50505050905090565b602654604080516001600160a01b0392831660248083018290528351808403909101815260449092018352602080830180516001600160e01b0316630427d73b60e51b1790525492517f282906ed000000000000000000000000000000000000000000000000000000008152620186a09492939092169163282906ed9185916200284e919083906028906004016200f781565b5f604051808303818588803b15801562002866575f80fd5b505af115801562002879573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392507f84222ac7000000000000000000000000000000000000000000000000000000009150620028cc905062030d40866200f52d565b620028db62030d40876200f52d565b620028e89060016200f52d565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b90921682526200295a916004016200f543565b5f604051808303815f87803b15801562002972575f80fd5b505af115801562002985573d5f803e3d5ffd5b50506020546001600160a01b0316915063397e375c9050620029ab62030d40856200f52d565b620029b89060016200f52d565b6026546040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815262001d59916001600160a01b031690879087906028906004016200f995565b6023546021546040516370a0823160e01b81526001600160a01b039182166004820152620186a0925f9216906370a0823190602401602060405180830381865afa15801562002a57573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002a7d91906200f8e7565b905062002a8b5f826200ace9565b6026546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052602354905491517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101879052929350169063095ea7b3906044016020604051808303815f875af115801562002b3d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002b6391906200f8ff565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562002bf0575f80fd5b505af115801562002c03573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f9262002c5b9289929091169087906028906200f958565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789362002cc1939082169289929091169087906028906004016200fa49565b5f604051808303815f87803b15801562002cd9575f80fd5b505af115801562002cec573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa15801562002d3d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002d6391906200f8e7565b905062002d7184826200ace9565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa15801562002dc0573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002de691906200f8e7565b905062002dfc85602c5462000d3a91906200f9d8565b5050505050565b6020546026546040517f726ac97c000000000000000000000000000000000000000000000000000000008152620186a0926001600160a01b039081169263726ac97c92859262002e5b9216906028906004016200f75e565b5f604051808303818588803b15801562002e73575f80fd5b505af115801562002e86573d5f803e3d5ffd5b5050604051630618f58760e51b81527f7671265e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063c31eb0e0925060240190505f604051808303815f87803b15801562002ef6575f80fd5b505af115801562002f09573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed935062030d409262000f569216905f906028906004016200f781565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052515f602482015261c35060448201819052919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f28dceb3907f8afe4db400000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529181526020820180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b90921682526200305f916004016200f543565b5f604051808303815f87803b15801562003077575f80fd5b505af11580156200308a573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb49350859262001d5992169087906028906004016200faa1565b60275460255460405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152620186a0926001600160a01b039081163192163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b1580156200315e575f80fd5b505af115801562003171573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c90620031c19087905f906028906200f728565b60405180910390a36020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263282906ed92879262003220929091169083906028906004016200f781565b5f604051808303818588803b15801562003238575f80fd5b505af11580156200324b573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063491cc7c2925060840190505f604051808303815f87803b158015620032b1575f80fd5b505af1158015620032c4573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c90620033149087905f906028906200f728565b60405180910390a36020546001600160a01b031663282906ed6200333c62030d40866200f52d565b6026546040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815262003388916001600160a01b03169088906028906004016200f781565b5f604051808303818588803b158015620033a0575f80fd5b505af1158015620033b3573d5f803e3d5ffd5b50506027546025546001600160a01b0391821631945016319150620033e4905062030d406200229c8760026200f9be565b62002dfc62030d40620022d08760026200f9be565b606060188054806020026020016040519081016040528092919081815260200182805480156200234057602002820191905f5260205f209081546001600160a01b0316815260019091019060200180831162002321575050505050905090565b606060178054806020026020016040519081016040528092919081815260200182805480156200234057602002820191905f5260205f209081546001600160a01b0316815260019091019060200180831162002321575050505050905090565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052602354905491517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101869052929350169063095ea7b3906044016020604051808303815f875af115801562003570573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200359691906200f8ff565b506040805160a0810182526103218082525f6020808401829052838501929092528351918201909352828152606082015260808101620035db621e848060016200f52d565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162003677916004016200f543565b5f604051808303815f87803b1580156200368f575f80fd5b505af1158015620036a2573d5f803e3d5ffd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b78945062003701939283169289921690889088906004016200fad8565b5f604051808303815f87803b15801562001d71575f80fd5b60208054604080516328ae864d60e21b81529051620186a0935f936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa1580156200376a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200379091906200f8e7565b6200379c91906200fb24565b67ffffffffffffffff811115620037b757620037b76200f7aa565b6040519080825280601f01601f191660200182016040528015620037e2576020820181803683370190505b5060208054604080516328ae864d60e21b815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa1580156200382d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200385391906200f8e7565b6200385f91906200fb24565b6200386c9060016200f52d565b67ffffffffffffffff8111156200388757620038876200f7aa565b6040519080825280601f01601f191660200182016040528015620038b2576020820181803683370190505b50602a90620038c290826200f81f565b505f60286002018054620038d6906200f5ea565b8351620038e492506200f52d565b60208054604080516328ae864d60e21b815290519394505f936001600160a01b039092169263a2ba1934926004808401938290030181865afa1580156200392d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200395391906200f8e7565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620039e3916004016200f543565b5f604051808303815f87803b158015620039fb575f80fd5b505af115801562003a0e573d5f803e3d5ffd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b9350889262003a6792169088906028906004016200faa1565b5f604051808303818588803b15801562003a7f575f80fd5b505af115801562003a92573d5f803e3d5ffd5b505050505050505050565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a0810183526103218082525f8284018190528285019190915283519283019093528282526060810191909152919250906080810162003b2f621e848060016200f52d565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162003bcb916004016200f543565b5f604051808303815f87803b15801562003be3575f80fd5b505af115801562003bf6573d5f803e3d5ffd5b50506020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063397e375c935087926200263e9216908390889088906004016200fb5d565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052602554602754915163248e63e160e11b8152600160048201819052602482018190526044820181905260648201529293506001600160a01b03908116319291163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562003d0e575f80fd5b505af115801562003d21573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d9749062003d6f9087906028906200fb9a565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb49262003dcb9291169087906028906004016200faa1565b5f604051808303815f87803b15801562003de3575f80fd5b505af115801562003df6573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562003e5a575f80fd5b505af115801562003e6d573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d9749062003ebb9087906028906200fb9a565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb49262030d409262003f1d929091169088906028906004016200faa1565b5f604051808303818588803b15801562003f35575f80fd5b505af115801562003f48573d5f803e3d5ffd5b50506025546027546001600160a01b039182163194501631915062003f779050620022b462030d40866200f9d8565b62002dfc62000d3a62030d40856200f52d565b6026546040516001600160a01b039091166024820152620186a09061c350905f9060440160408051601f19818403018152919052602080820180516001600160e01b0316630427d73b60e51b17905260235490549192506001600160a01b039081169163095ea7b39116620040018660026200f9be565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562004062573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200408891906200f8ff565b506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b7893620040e7939082169289929091169087906028906004016200fa49565b5f604051808303815f87803b158015620040ff575f80fd5b505af115801562004112573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f8afe4db400000000000000000000000000000000000000000000000000000000905062030d406200416486826200f52d565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b9092168252620041d6916004016200f543565b5f604051808303815f87803b158015620041ee575f80fd5b505af115801562004201573d5f803e3d5ffd5b50506020546001600160a01b0316915063d09e3b789050620042278462030d406200f52d565b6026546023546040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526200263e926001600160a01b03908116928a9291169088906028906004016200fa49565b620186a05f62004292600162030d406200f9d8565b6026546040516001600160a01b0390911660248201529091505f9060440160408051601f19818403018152919052602080820180516001600160e01b0316630427d73b60e51b17905260235490549192506001600160a01b039081169163095ea7b39116620043038660026200f9be565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562004364573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200438a91906200f8ff565b506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b7893620043e9939082169289929091169087906028906004016200fa49565b5f604051808303815f87803b15801562004401575f80fd5b505af115801562004414573d5f803e3d5ffd5b505060405162030d40602482015260448101859052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507fa458261b00000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529181526020820180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b9092168252620044cc916004016200f543565b5f604051808303815f87803b158015620044e4575f80fd5b505af1158015620044f7573d5f803e3d5ffd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b78945087936200263e938116928a9291169088906028906004016200fa49565b6026546040516001600160a01b0390911660248201525f90819060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f7671265e000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b1580156200460c575f80fd5b505af11580156200461f573d5f803e3d5ffd5b50506020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063397e375c9350869262001d59921690839087906028906004016200f995565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303815f875af1158015620046ed573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200471391906200f8ff565b5060208054604080516328ae864d60e21b815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562004758573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200477e91906200f8e7565b6200478b9060016200f52d565b67ffffffffffffffff811115620047a657620047a66200f7aa565b6040519080825280601f01601f191660200182016040528015620047d1576020820181803683370190505b50602a90620047e190826200f81f565b505f60286002018054620047f5906200f5ea565b905090505f60205f9054906101000a90046001600160a01b03166001600160a01b031663a2ba19346040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200484b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200487191906200f8e7565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162004901916004016200f543565b5f604051808303815f87803b15801562004919575f80fd5b505af11580156200492c573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b09450620037019392831692899216906028906004016200f920565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562004a21575f80fd5b505af115801562004a34573d5f803e3d5ffd5b50506020546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250631becceb4915062004a88905f9085906028906004016200faa1565b5f604051808303815f87803b15801562004aa0575f80fd5b505af115801562002dfc573d5f803e3d5ffd5b620186a0737109709ecfa91a80626ff3989d68f67f5b1dd12d63f28dceb37f84222ac70000000000000000000000000000000000000000000000000000000062004aff8460016200f52d565b60405160248101919091526044810185905260640160408051601f198184030181529181526020820180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b909216825262004b73916004016200f543565b5f604051808303815f87803b15801562004b8b575f80fd5b505af115801562004b9e573d5f803e3d5ffd5b50506020546026546001600160a01b03918216935063282906ed925084911662004bca8260016200f52d565b60286040518563ffffffff1660e01b815260040162000f56939291906200f781565b6026546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562004c9f575f80fd5b505af115801562004cb2573d5f803e3d5ffd5b50506020546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063744b9b8b9150849062001d59905f9086906028906004016200faa1565b60208054604080516328ae864d60e21b815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562004d4c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004d7291906200f8e7565b62004d7f9060016200f52d565b67ffffffffffffffff81111562004d9a5762004d9a6200f7aa565b6040519080825280601f01601f19166020018201604052801562004dc5576020820181803683370190505b50602a9062004dd590826200f81f565b505f6028600201805462004de9906200f5ea565b905090505f60205f9054906101000a90046001600160a01b03166001600160a01b031663a2ba19346040518163ffffffff1660e01b8152600401602060405180830381865afa15801562004e3f573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004e6591906200f8e7565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162004ef5916004016200f543565b5f604051808303815f87803b15801562004f0d575f80fd5b505af115801562004f20573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed935060019262001d5992169083906028906004016200f781565b60405163248e63e160e11b8152600160048201819052602482018190526044820181905260648201526509184e72a00090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562004fe2575f80fd5b505af115801562004ff5573d5f803e3d5ffd5b50506040805162030d408152602081018590527fe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8935001905060405180910390a16020546040517f7c744253000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b0390911690637c744253906024015f604051808303815f87803b15801562005093575f80fd5b505af1158015620050a6573d5f803e3d5ffd5b505060208054604080517fb010721400000000000000000000000000000000000000000000000000000000815290516200513495506001600160a01b03909216935063b01072149260048083019391928290030181865afa1580156200510e573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000d3a91906200f8e7565b50565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b17905290505f6200518f600162030d406200f9d8565b6020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081529293506001600160a01b0391821692631becceb492620051e592169086906028906004016200faa1565b5f604051808303815f87803b158015620051fd575f80fd5b505af115801562005210573d5f803e3d5ffd5b505060405162030d40602482015260448101849052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507fa458261b000000000000000000000000000000000000000000000000000000009060640162003000565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905260235460215491516370a0823160e01b81526001600160a01b0392831660048201529293505f929116906370a0823190602401602060405180830381865afa15801562005307573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200532d91906200f8e7565b6027546025546023546020549394506001600160a01b03928316319391831631929081169163095ea7b39116620053668860026200f9be565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af1158015620053c7573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620053ed91906200f8ff565b5060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b1580156200544e575f80fd5b505af115801562005461573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c92620054b7928b92909116906028906200f728565b60405180910390a3602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b0936200550293908216928b92909116906028906004016200f920565b5f604051808303815f87803b1580156200551a575f80fd5b505af11580156200552d573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562005591575f80fd5b505af1158015620055a4573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f92620055fc928b92909116908a906028906200f958565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789362030d40936200566693918316928c929116908b906028906004016200fa49565b5f604051808303818588803b1580156200567e575f80fd5b505af115801562005691573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063491cc7c2925060840190505f604051808303815f87803b158015620056f7575f80fd5b505af11580156200570a573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d97490620057589088906028906200fb9a565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb49262030d4092620057ba929091169089906028906004016200faa1565b5f604051808303818588803b158015620057d2575f80fd5b505af1158015620057e5573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319150602401602060405180830381865afa15801562005837573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200585d91906200f8e7565b6027546025549192506001600160a01b039081163191163162005899620058868960026200f9be565b6200589290886200f52d565b846200ace9565b620058b9620058ad62030d4060026200f9be565b620022b490876200f52d565b62002669620058cd62030d4060026200f9be565b62000d3a90866200f9d8565b6023546021546040516370a0823160e01b81526001600160a01b039182166004820152620186a0925f9216906370a0823190602401602060405180830381865afa1580156200592a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200595091906200f8e7565b90506200595e5f826200ace9565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526024810185905291169063095ea7b3906044016020604051808303815f875af1158015620059ca573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620059f091906200f8ff565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562005a7d575f80fd5b505af115801562005a90573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c9262005ae6928892909116906028906200f728565b60405180910390a3602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362005b3193908216928892909116906028906004016200f920565b5f604051808303815f87803b15801562005b49575f80fd5b505af115801562005b5c573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa15801562005bad573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005bd391906200f8e7565b905062005be183826200ace9565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa15801562005c30573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005c5691906200f8e7565b905062001ae384602c5462000d3a91906200f9d8565b6060601b805480602002602001604051908101604052809291908181526020015f905b82821015620027b2578382905f5260205f2090600202016040518060400160405290815f8201805462005cc2906200f5ea565b80601f016020809104026020016040519081016040528092919081815260200182805462005cf0906200f5ea565b801562005d3f5780601f1062005d155761010080835404028352916020019162005d3f565b820191905f5260205f20905b81548152906001019060200180831162005d2157829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801562005dc357602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841162005d845790505b5050505050815250508152602001906001019062005c8f565b620186a0737109709ecfa91a80626ff3989d68f67f5b1dd12d63f28dceb37f84222ac70000000000000000000000000000000000000000000000000000000062005e286001856200f9d8565b60405160248101919091526044810185905260640160408051601f198184030181529181526020820180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b909216825262005e9c916004016200f543565b5f604051808303815f87803b15801562005eb4575f80fd5b505af115801562005ec7573d5f803e3d5ffd5b50506020546026546001600160a01b03918216935063282906ed925084911662004bca6001836200f9d8565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905260235460215491516370a0823160e01b81526001600160a01b0392831660048201529293505f929116906370a0823190602401602060405180830381865afa15801562005f8c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005fb291906200f8e7565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa15801562006002573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200602891906200f8e7565b6027546023546020549293506001600160a01b0391821631929082169163095ea7b39116620060598860026200f9be565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af1158015620060ba573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620060e091906200f8ff565b5060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562006141575f80fd5b505af115801562006154573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f92620061ac928b92909116908a906028906200f958565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b78936200621293908216928b92909116908a906028906004016200fa49565b5f604051808303815f87803b1580156200622a575f80fd5b505af11580156200623d573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b158015620062a1575f80fd5b505af1158015620062b4573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f926200630c928b92909116908a906028906200f958565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789362030d40936200637693918316928c929116908b906028906004016200fa49565b5f604051808303818588803b1580156200638e575f80fd5b505af1158015620063a1573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319150602401602060405180830381865afa158015620063f3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200641991906200f8e7565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa15801562006469573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200648f91906200f8e7565b6027549091506001600160a01b031631620064b0620058868960026200f9be565b620064cd620064c18960026200f9be565b620022b490876200f9d8565b6200266962000d3a62030d40866200f52d565b6027546026546040516001600160a01b039182166024820152620186a0929190911631905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015620065c0575f80fd5b505af1158015620065d3573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f90620066259087905f9087906028906200f958565b60405180910390a36020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263397e375c928792620066869290911690839087906028906004016200f995565b5f604051808303818588803b1580156200669e575f80fd5b505af1158015620066b1573d5f803e3d5ffd5b50506027546001600160a01b031631925062001ae3915062000d3a905085856200f52d565b604051630618f58760e51b81527f7671265e0000000000000000000000000000000000000000000000000000000060048201525f90737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b15801562006743575f80fd5b505af115801562006756573d5f803e3d5ffd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c9350859262000f569216906028906004016200f75e565b602654604080516001600160a01b0392831660248083018290528351808403909101815260449092018352602080830180516001600160e01b0316630427d73b60e51b1790525492517f744b9b8b000000000000000000000000000000000000000000000000000000008152620186a09492939092169163744b9b8b91859162006840919086906028906004016200faa1565b5f604051808303818588803b15801562006858575f80fd5b505af11580156200686b573d5f803e3d5ffd5b50506040805160048082526024820183526020820180516001600160e01b03167f394836a400000000000000000000000000000000000000000000000000000000179052915163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d955063f28dceb39450620068e893509091016200f543565b5f604051808303815f87803b15801562006900575f80fd5b505af115801562006913573d5f803e3d5ffd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b9350869262001d5992169086906028906004016200faa1565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015620069d9575f80fd5b505af1158015620069ec573d5f803e3d5ffd5b5050602054602354604051630102614b60e41b81526001600160a01b03928316945063102614b0935062004a88925f9287929116906028906004016200f920565b6026546040516001600160a01b0390911660248201525f90819060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f7671265e000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401620068e8565b6060601a805480602002602001604051908101604052809291908181526020015f905b82821015620027b2578382905f5260205f2001805462006b10906200f5ea565b80601f016020809104026020016040519081016040528092919081815260200182805462006b3e906200f5ea565b801562006b8d5780601f1062006b635761010080835404028352916020019162006b8d565b820191905f5260205f20905b81548152906001019060200180831162006b6f57829003601f168201915b50505050508152602001906001019062006af0565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b1790529050737109709ecfa91a80626ff3989d68f67f5b1dd12d63f28dceb37f84222ac70000000000000000000000000000000000000000000000000000000062006c368560016200f52d565b60405160248101919091526044810186905260640160408051601f198184030181529181526020820180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b909216825262006caa916004016200f543565b5f604051808303815f87803b15801562006cc2575f80fd5b505af115801562006cd5573d5f803e3d5ffd5b50506020546026546001600160a01b03918216935063397e375c925085911662006d018260016200f52d565b8560286040518663ffffffff1660e01b815260040162001d5994939291906200f995565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f6024820181905292919091169063095ea7b3906044016020604051808303815f875af115801562006d95573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062006dbb91906200f8ff565b50604051630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b15801562006e27575f80fd5b505af115801562006e3a573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b0945062004a889392831692879216906028906004016200f920565b6040805160a0810182526103218082525f60208084018290528385019290925283519182019093528281526060820152620186a091906080810162006ec9621e848060016200f52d565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162006f65916004016200f543565b5f604051808303815f87803b15801562006f7d575f80fd5b505af115801562006f90573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed9350869262001d59921690839087906004016200fbc2565b6026546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156200709b575f80fd5b505af1158015620070ae573d5f803e3d5ffd5b50506020546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063d09e3b78935062000b5d925f92889291169087906028906004016200fa49565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562007177575f80fd5b505af11580156200718a573d5f803e3d5ffd5b50506020546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063726ac97c9150839062000f56905f906028906004016200f75e565b602654604080516001600160a01b0392831660248083018290528351808403909101815260449092018352602080830180516001600160e01b0316630427d73b60e51b1790525492517f282906ed000000000000000000000000000000000000000000000000000000008152620186a09492939092169163282906ed91859162007271919083906028906004016200f781565b5f604051808303818588803b15801562007289575f80fd5b505af11580156200729c573d5f803e3d5ffd5b50506040805162030d406024820152604480820188905282518083039091018152606490910182526020810180516001600160e01b03167fa458261b00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d945063f28dceb39350620045f492506004016200f543565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b1790529050737109709ecfa91a80626ff3989d68f67f5b1dd12d63f28dceb37f84222ac700000000000000000000000000000000000000000000000000000000620073c46001866200f9d8565b60405160248101919091526044810186905260640160408051601f198184030181529181526020820180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b909216825262007438916004016200f543565b5f604051808303815f87803b15801562007450575f80fd5b505af115801562007463573d5f803e3d5ffd5b50506020546026546001600160a01b03918216935063397e375c925085911662006d016001836200f9d8565b6060601d805480602002602001604051908101604052809291908181526020015f905b82821015620027b2575f8481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156200755b57602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116200751c5790505b50505050508152505081526020019060010190620074b2565b6020546026546040517f726ac97c000000000000000000000000000000000000000000000000000000008152620186a0926001600160a01b039081169263726ac97c928592620075cc9216906028906004016200f75e565b5f604051808303818588803b158015620075e4575f80fd5b505af1158015620075f7573d5f803e3d5ffd5b50506040805160048082526024820183526020820180516001600160e01b03167f394836a400000000000000000000000000000000000000000000000000000000179052915163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d955063f28dceb394506200672b93509091016200f543565b60208054604080516328ae864d60e21b81529051620186a0935f936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa158015620076c5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620076eb91906200f8e7565b620076f791906200fb24565b67ffffffffffffffff8111156200771257620077126200f7aa565b6040519080825280601f01601f1916602001820160405280156200773d576020820181803683370190505b5060208054604080516328ae864d60e21b815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562007788573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620077ae91906200f8e7565b620077ba91906200fb24565b620077c79060016200f52d565b67ffffffffffffffff811115620077e257620077e26200f7aa565b6040519080825280601f01601f1916602001820160405280156200780d576020820181803683370190505b50602a906200781d90826200f81f565b506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526024810185905291169063095ea7b3906044016020604051808303815f875af11580156200788a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620078b091906200f8ff565b505f60286002018054620078c4906200f5ea565b8351620078d292506200f52d565b60208054604080516328ae864d60e21b815290519394505f936001600160a01b039092169263a2ba1934926004808401938290030181865afa1580156200791b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200794191906200f8e7565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620079d1916004016200f543565b5f604051808303815f87803b158015620079e9575f80fd5b505af1158015620079fc573d5f803e3d5ffd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b78945062007a5c93928316928a92169089906028906004016200fa49565b5f604051808303815f87803b15801562002656575f80fd5b6026546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562007b27575f80fd5b505af115801562007b3a573d5f803e3d5ffd5b50506020546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063397e375c9150849062001d59905f90839087906028906004016200f995565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562007bff575f80fd5b505af115801562007c12573d5f803e3d5ffd5b50506020546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063282906ed9150839062000f56905f9083906028906004016200f781565b6027546026546040516001600160a01b039182166024820152620186a0929190911631905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562007d48575f80fd5b505af115801562007d5b573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f9062007dad9087905f9087906028906200f958565b60405180910390a36020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263744b9b8b92879262006686929091169086906028906004016200faa1565b60208054604080516328ae864d60e21b81529051620186a0935f936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa15801562007e5d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062007e8391906200f8e7565b62007e8f91906200fb24565b67ffffffffffffffff81111562007eaa5762007eaa6200f7aa565b6040519080825280601f01601f19166020018201604052801562007ed5576020820181803683370190505b5060208054604080516328ae864d60e21b815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562007f20573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062007f4691906200f8e7565b62007f5291906200fb24565b62007f5f9060016200f52d565b67ffffffffffffffff81111562007f7a5762007f7a6200f7aa565b6040519080825280601f01601f19166020018201604052801562007fa5576020820181803683370190505b50602a9062007fb590826200f81f565b505f6028600201805462007fc9906200f5ea565b835162007fd792506200f52d565b60208054604080516328ae864d60e21b815290519394505f936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562008020573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200804691906200f8e7565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620080d6916004016200f543565b5f604051808303815f87803b158015620080ee575f80fd5b505af115801562008101573d5f803e3d5ffd5b50506020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063397e375c9350889262003a67921690839089906028906004016200f995565b6027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201526509184e72a00090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620081d4575f80fd5b505af1158015620081e7573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562008246575f80fd5b505af115801562008259573d5f803e3d5ffd5b50506020546040517f7c744253000000000000000000000000000000000000000000000000000000008152600481018590526001600160a01b039091169250637c744253915060240162004a88565b6026546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b1580156200837a575f80fd5b505af11580156200838d573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d97490620083db9085906028906200fb9a565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb49262004a889291169085906028906004016200faa1565b6060601c805480602002602001604051908101604052809291908181526020015f905b82821015620027b2575f8481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156200850357602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b03191681526020019060040190602082600301049283019260010382029150808411620084c45790505b505050505081525050815260200190600101906200845a565b60208054604080516328ae864d60e21b815290515f936002936001600160a01b03169263a2ba193492600480830193928290030181865afa15801562008564573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200858a91906200f8e7565b6200859691906200fb24565b67ffffffffffffffff811115620085b157620085b16200f7aa565b6040519080825280601f01601f191660200182016040528015620085dc576020820181803683370190505b5060208054604080516328ae864d60e21b815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562008627573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200864d91906200f8e7565b6200865991906200fb24565b620086669060016200f52d565b67ffffffffffffffff8111156200868157620086816200f7aa565b6040519080825280601f01601f191660200182016040528015620086ac576020820181803683370190505b50602a90620086bc90826200f81f565b505f60286002018054620086d0906200f5ea565b8351620086de92506200f52d565b60208054604080516328ae864d60e21b815290519394505f936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562008727573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200874d91906200f8e7565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620087dd916004016200f543565b5f604051808303815f87803b158015620087f5575f80fd5b505af115801562008808573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb4935062003701929091169087906028906004016200faa1565b6040805160a0810182526103218082525f60208084018290528385019290925283519182019093528281526060820152620186a0919060808101620088ab621e848060016200f52d565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162008947916004016200f543565b5f604051808303815f87803b1580156200895f575f80fd5b505af115801562008972573d5f803e3d5ffd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c9350869262001d5992169086906004016200fbeb565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303815f875af115801562008a3b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008a6191906200f8ff565b506025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562008ad2575f80fd5b505af115801562008ae5573d5f803e3d5ffd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a59042791506024015f604051808303815f87803b15801562008b49575f80fd5b505af115801562008b5c573d5f803e3d5ffd5b5050602354604080516001600160a01b0390921660248084019190915281518084039091018152604490920181526020820180516001600160e01b03167f1387a349000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925062006e0f91906004016200f543565b60606019805480602002602001604051908101604052809291908181526020015f905b82821015620027b2578382905f5260205f2001805462008c38906200f5ea565b80601f016020809104026020016040519081016040528092919081815260200182805462008c66906200f5ea565b801562008cb55780601f1062008c8b5761010080835404028352916020019162008cb5565b820191905f5260205f20905b81548152906001019060200180831162008c9757829003601f168201915b50505050508152602001906001019062008c18565b602354602054620186a09161c350916001600160a01b039182169163095ea7b3911662008cf98560026200f9be565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562008d5a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008d8091906200f8ff565b50602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362008dc493908216928892909116906028906004016200f920565b5f604051808303815f87803b15801562008ddc575f80fd5b505af115801562008def573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f8afe4db400000000000000000000000000000000000000000000000000000000905062030d4062008e4185826200f52d565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b909216825262008eb3916004016200f543565b5f604051808303815f87803b15801562008ecb575f80fd5b505af115801562008ede573d5f803e3d5ffd5b50506020546001600160a01b0316915063102614b0905062008f048362030d406200f52d565b6026546023546040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815262001d59926001600160a01b039081169289929116906028906004016200f920565b6008545f9060ff161562008f70575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c6564000000000000000000000000000000000000000000000000000060248301525f9163667f9d7090604401602060405180830381865afa15801562008fff573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200902591906200f8e7565b1415905090565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052602880547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905551630618f58760e51b81527f19b4bff2000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156200911b575f80fd5b505af11580156200912e573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb4935062004a88929091169085906028906004016200faa1565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a0810183526103218082525f8284018190528285019190915283519283019093528282526060810191909152919250906080810162009219621e848060016200f52d565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620092b5916004016200f543565b5f604051808303815f87803b158015620092cd575f80fd5b505af1158015620092e0573d5f803e3d5ffd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b935087926200263e921690879087906004016200f5b3565b602654604080516001600160a01b0392831660248083018290528351808403909101815260449092018352602080830180516001600160e01b0316630427d73b60e51b1790525492517f1becceb4000000000000000000000000000000000000000000000000000000008152919361c350931691631becceb491620093c59186906028906004016200faa1565b5f604051808303815f87803b158015620093dd575f80fd5b505af1158015620093f0573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f8afe4db400000000000000000000000000000000000000000000000000000000905062030d406200944285826200f52d565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b9092168252620094b4916004016200f543565b5f604051808303815f87803b158015620094cc575f80fd5b505af1158015620094df573d5f803e3d5ffd5b50506020546001600160a01b03169150631becceb49050620095058362030d406200f52d565b6026546040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815262001d59916001600160a01b03169087906028906004016200faa1565b604051630618f58760e51b81527f7671265e0000000000000000000000000000000000000000000000000000000060048201525f90737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b158015620095be575f80fd5b505af1158015620095d1573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed9350859262000f5692169083906028906004016200f781565b620186a05f6200963f600162030d406200f9d8565b6023546020549192506001600160a01b039081169163095ea7b39116620096688560026200f9be565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af1158015620096c9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620096ef91906200f8ff565b50602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b0936200973393908216928892909116906028906004016200f920565b5f604051808303815f87803b1580156200974b575f80fd5b505af11580156200975e573d5f803e3d5ffd5b505060405162030d40602482015260448101849052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507fa458261b000000000000000000000000000000000000000000000000000000009060640162001e7a565b6020546040517f7c7442530000000000000000000000000000000000000000000000000000000081525f60048201526001600160a01b0390911690637c744253906024015f604051808303815f87803b15801562009818575f80fd5b505af11580156200982b573d5f803e3d5ffd5b50506026546040516001600160a01b039091166024820152620186a092505f915060440160408051601f19818403018152919052602080820180516001600160e01b0316630427d73b60e51b17905260235490549192506001600160a01b039081169163095ea7b39116620098a28560026200f9be565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562009903573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200992991906200f8ff565b50602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b0936200996d93908216928892909116906028906004016200f920565b5f604051808303815f87803b15801562009985575f80fd5b505af115801562009998573d5f803e3d5ffd5b5050604051630618f58760e51b81527f394836a4000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b15801562009a06575f80fd5b505af115801562009a19573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b0945062009a5e9392831692889216906028906004016200f920565b5f604051808303815f87803b15801562009a76575f80fd5b505af115801562009a89573d5f803e3d5ffd5b5050604051630618f58760e51b81527f394836a4000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b15801562009af7575f80fd5b505af115801562009b0a573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb4935062000b5d929091169085906028906004016200faa1565b6023546021546040516370a0823160e01b81526001600160a01b039182166004820152620186a0925f9216906370a0823190602401602060405180830381865afa15801562009bb4573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009bda91906200f8e7565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa15801562009c2a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009c5091906200f8e7565b6027546023546020549293506001600160a01b0391821631929082169163095ea7b3911662009c818760026200f9be565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562009ce2573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009d0891906200f8ff565b5060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562009d69575f80fd5b505af115801562009d7c573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c9262009dd2928a92909116906028906200f728565b60405180910390a3602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362009e1d93908216928a92909116906028906004016200f920565b5f604051808303815f87803b15801562009e35575f80fd5b505af115801562009e48573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562009eac575f80fd5b505af115801562009ebf573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c9262009f15928a92909116906028906200f728565b60405180910390a3602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362030d409362009f6493918316928b929116906028906004016200f920565b5f604051808303818588803b15801562009f7c575f80fd5b505af115801562009f8f573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319150602401602060405180830381865afa15801562009fe1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a00791906200f8e7565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa1580156200a057573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a07d91906200f8e7565b6027549091506001600160a01b0316316200a09e620058868860026200f9be565b6200a0af620064c18860026200f9be565b62001d8462000d3a62030d40866200f52d565b606060158054806020026020016040519081016040528092919081815260200182805480156200234057602002820191905f5260205f209081546001600160a01b0316815260019091019060200180831162002321575050505050905090565b602480546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a093810184905291169063095ea7b3906044016020604051808303815f875af11580156200a192573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a1b891906200f8ff565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b1580156200a245575f80fd5b505af11580156200a258573d5f803e3d5ffd5b50506026546025546024546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c926200a2ae928792909116906028906200f728565b60405180910390a3602054602654602454604051630102614b60e41b81526001600160a01b039384169363102614b0936200a2f993908216928792909116906028906004016200f920565b5f604051808303815f87803b1580156200a311575f80fd5b505af11580156200a324573d5f803e3d5ffd5b5050602480546025546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319101602060405180830381865afa1580156200a374573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a39a91906200f8e7565b90506200a3b082602c5462000d3a91906200f9d8565b5050565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052602354905491517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101869052929350169063095ea7b3906044016020604051808303815f875af11580156200a46b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a49191906200f8ff565b506025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200a502575f80fd5b505af11580156200a515573d5f803e3d5ffd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a59042791506024015f604051808303815f87803b1580156200a579575f80fd5b505af11580156200a58c573d5f803e3d5ffd5b5050602354604080516001600160a01b0390921660248084019190915281518084039091018152604490920181526020820180516001600160e01b03167f1387a349000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506200a62591906004016200f543565b5f604051808303815f87803b1580156200a63d575f80fd5b505af11580156200a650573d5f803e3d5ffd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b78945062000b5d93928316928892169087906028906004016200fa49565b6026546040516001600160a01b039091166024820152620186a09061c350905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052602354905491517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101879052929350169063095ea7b3906044016020604051808303815f875af11580156200a76b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a79191906200f8ff565b506040515f602482015260448101839052737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f28dceb3907f8afe4db400000000000000000000000000000000000000000000000000000000906064016200446d565b6020546040517f7c7442530000000000000000000000000000000000000000000000000000000081526509184e72a0006004820181905291620186a0916001600160a01b0390911690637c744253906024015f604051808303815f87803b1580156200a853575f80fd5b505af11580156200a866573d5f803e3d5ffd5b50506027546025546020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b039485163196509284163194509083169263726ac97c9287926200a8cd9216906028906004016200f75e565b5f604051808303818588803b1580156200a8e5575f80fd5b505af11580156200a8f8573d5f803e3d5ffd5b50506020546001600160a01b0316925063282906ed91506200a91d905086866200f52d565b6026546040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526200a969916001600160a01b03169088906028906004016200f781565b5f604051808303818588803b1580156200a981575f80fd5b505af11580156200a994573d5f803e3d5ffd5b50506027546025546001600160a01b03918216319450163191506200a9c29050866200229c8760026200f9be565b62000b8886620022d08760026200f9be565b6027546020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039182166084820152620186a092919091163190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b1580156200aa6e575f80fd5b505af11580156200aa81573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c906200aad19086905f906028906200f728565b60405180910390a36020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263282906ed92869262000cea929091169083906028906004016200f781565b6020546026546040517f282906ed000000000000000000000000000000000000000000000000000000008152620186a0926001600160a01b039081169263282906ed9285926200ab8a92169083906028906004016200f781565b5f604051808303818588803b1580156200aba2575f80fd5b505af11580156200abb5573d5f803e3d5ffd5b50506040805162030d406024820152604480820187905282518083039091018152606490910182526020810180516001600160e01b03167fa458261b00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d945063f28dceb39350620095a692506004016200f543565b6026546040516001600160a01b0390911660248201525f90819060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024016200a625565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015f6040518083038186803b1580156200ad53575f80fd5b505afa15801562000b88573d5f803e3d5ffd5b5f6200ad716200f0b7565b6200ad7e8484836200ad88565b9150505b92915050565b5f806200ad9685846200ae08565b90506200adfd6040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016200ade79291906200fa26565b604051602081830303815290604052856200ae15565b9150505b9392505050565b5f6200ae0183836200ae48565b60c0810151515f90156200ae3c576200ae3484848460c001516200ae66565b90506200ae01565b6200ae3484846200b01a565b5f6200ae5583836200b10b565b6200ae01838360200151846200ae15565b5f806200ae726200b118565b90505f6200ae8186836200b1ec565b90505f6200ae9982606001518360200151856200b6a4565b90505f6200aeaa838389896200b8c9565b90505f6200aeb8826200c80b565b602081015181519192509060030b156200af30578982604001516040516020016200aee59291906200fc25565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526200af27916004016200f543565b60405180910390fd5b5f6200af746040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a2000000000000000000000008152508360016200c9e1565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906200afc99084906004016200f543565b602060405180830381865afa1580156200afe5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b00b91906200fc8a565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081525f908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc925906200b0709087906004016200f543565b5f60405180830381865afa1580156200b08b573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b0b491908101906200fd7d565b90505f6200b0e582856040516020016200b0d09291906200fdb3565b6040516020818303038152906040526200cbe6565b90506001600160a01b0381166200ad7e5784846040516020016200aee59291906200fdcb565b6200a3b082825f6200cbf7565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c906200b1a19084906004016200fe5f565b5f60405180830381865afa1580156200b1bc573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b1e591908101906200fea7565b9250505090565b6200b21f6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b5f737109709ecfa91a80626ff3989d68f67f5b1dd12d90506200b26a6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6200b275856200cd04565b60208201525f6200b286866200d0fd565b90505f836001600160a01b031663d930a0e66040518163ffffffff1660e01b81526004015f60405180830381865afa1580156200b2c5573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b2ee91908101906200fea7565b868385602001516040516020016200b30a94939291906200fef1565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291505f906001600160a01b038616906360f9bb11906200b3639085906004016200f543565b5f60405180830381865afa1580156200b37e573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b3a791908101906200fea7565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906200b3f19084906004016200ffc9565b602060405180830381865afa1580156200b40d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b43391906200f8ff565b6200b44b57816040516020016200aee591906201001c565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906200b492908490600401620100a2565b5f60405180830381865afa1580156200b4ad573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b4d691908101906200fea7565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906200b51f908490600401620100f5565b602060405180830381865afa1580156200b53b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b56191906200f8ff565b156200b5f8576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906200b5ae908490600401620100f5565b5f60405180830381865afa1580156200b5c9573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b5f291908101906200fea7565b60408501525b846001600160a01b03166349c4fac882865f01516040516020016200b61e919062010148565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016200b64c929190620101a8565b5f60405180830381865afa1580156200b667573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b69091908101906200fea7565b606085015250608083015250949350505050565b60408051600480825260a082019092526060915f9190816020015b60608152602001906001900390816200b6bf5790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250815f815181106200b722576200b722620101d0565b60200260200101819052506040518060400160405280600381526020017f2d726c0000000000000000000000000000000000000000000000000000000000815250816001815181106200b779576200b779620101d0565b6020026020010181905250846040516020016200b7979190620101fd565b604051602081830303815290604052816002815181106200b7bc576200b7bc620101d0565b6020026020010181905250826040516020016200b7da91906201025d565b604051602081830303815290604052816003815181106200b7ff576200b7ff620101d0565b60200260200101819052505f6200b816826200c80b565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184525f80825290860152825180840190935290518252928101929092529192506200b8a7906040805180820182525f80825260209182015281518083019092528451825280850190820152906200d392565b6200b8bf57856040516020016200aee5919062010297565b9695505050505050565b60a08101516040805180820182525f80825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d90156200b919565b511590565b6200ba92578260200151156200b9d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a4016200af27565b8260c00151156200ba92576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a4016200af27565b6040805160ff80825261200082019092525f91816020015b60608152602001906001900390816200baaa5790505090505f6040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806200bb079062010316565b935060ff16815181106200bb1f576200bb1f620101d0565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e37000000000000000000000000000000000000008152506040516020016200bb72919062010337565b6040516020818303038152906040528282806200bb8f9062010316565b935060ff16815181106200bba7576200bba7620101d0565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806200bbf69062010316565b935060ff16815181106200bc0e576200bc0e620101d0565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806200bc5d9062010316565b935060ff16815181106200bc75576200bc75620101d0565b602002602001018190525087602001518282806200bc939062010316565b935060ff16815181106200bcab576200bcab620101d0565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e7472616374506174680000000000000000000000000000000000008152508282806200bcfa9062010316565b935060ff16815181106200bd12576200bd12620101d0565b6020908102919091010152875182826200bd2c8162010316565b935060ff16815181106200bd44576200bd44620101d0565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e496400000000000000000000000000000000000000000000008152508282806200bd939062010316565b935060ff16815181106200bdab576200bdab620101d0565b60200260200101819052506200bdc1466200d3f8565b82826200bdce8162010316565b935060ff16815181106200bde6576200bde6620101d0565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806200be359062010316565b935060ff16815181106200be4d576200be4d620101d0565b6020026020010181905250868282806200be679062010316565b935060ff16815181106200be7f576200be7f620101d0565b60209081029190910101528551156200bfb25760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826200bed38162010316565b935060ff16815181106200beeb576200beeb620101d0565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906200bf3d9089906004016200f543565b5f60405180830381865afa1580156200bf58573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200bf8191908101906200fea7565b82826200bf8e8162010316565b935060ff16815181106200bfa6576200bfa6620101d0565b60200260200101819052505b8460200151156200c08e5760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826200bffe8162010316565b935060ff16815181106200c016576200c016620101d0565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806200c0659062010316565b935060ff16815181106200c07d576200c07d620101d0565b60200260200101819052506200c270565b6200c0c76200b9148660a001516040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6200c1645760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826200c10d8162010316565b935060ff16815181106200c125576200c125620101d0565b60200260200101819052508460a001516040516020016200c1479190620101fd565b6040516020818303038152906040528282806200c0659062010316565b8460c001511580156200c1a85750604080890151815180830183525f808252602091820152825180840190935281518352908101908201526200c1a690511590565b155b156200c2705760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826200c1ef8162010316565b935060ff16815181106200c207576200c207620101d0565b60200260200101819052506200c21d886200d49c565b6040516020016200c22f9190620101fd565b6040516020818303038152906040528282806200c24c9062010316565b935060ff16815181106200c264576200c264620101d0565b60200260200101819052505b604080860151815180830183525f808252602091820152825180840190935281518352908101908201526200c2a490511590565b6200c3445760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826200c2ea8162010316565b935060ff16815181106200c302576200c302620101d0565b602002602001018190525084604001518282806200c3209062010316565b935060ff16815181106200c338576200c338620101d0565b60200260200101819052505b6060850151156200c46f5760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826200c3908162010316565b935060ff16815181106200c3a8576200c3a8620101d0565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e8906024015f60405180830381865afa1580156200c415573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200c43e91908101906200fea7565b82826200c44b8162010316565b935060ff16815181106200c463576200c463620101d0565b60200260200101819052505b60e085015151156200c5225760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826200c4bc8162010316565b935060ff16815181106200c4d4576200c4d4620101d0565b60200260200101819052506200c4f18560e001515f01516200d3f8565b82826200c4fe8162010316565b935060ff16815181106200c516576200c516620101d0565b60200260200101819052505b60e085015160200151156200c5d95760408051808201909152600a81527f2d2d676173507269636500000000000000000000000000000000000000000000602082015282826200c5728162010316565b935060ff16815181106200c58a576200c58a620101d0565b60200260200101819052506200c5a88560e00151602001516200d3f8565b82826200c5b58162010316565b935060ff16815181106200c5cd576200c5cd620101d0565b60200260200101819052505b60e085015160400151156200c6905760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826200c6298162010316565b935060ff16815181106200c641576200c641620101d0565b60200260200101819052506200c65f8560e00151604001516200d3f8565b82826200c66c8162010316565b935060ff16815181106200c684576200c684620101d0565b60200260200101819052505b60e085015160600151156200c7475760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826200c6e08162010316565b935060ff16815181106200c6f8576200c6f8620101d0565b60200260200101819052506200c7168560e00151606001516200d3f8565b82826200c7238162010316565b935060ff16815181106200c73b576200c73b620101d0565b60200260200101819052505b5f8160ff1667ffffffffffffffff8111156200c767576200c7676200f7aa565b6040519080825280602002602001820160405280156200c79c57816020015b60608152602001906001900390816200c7865790505b5090505f5b8260ff168160ff1610156200c7fc57838160ff16815181106200c7c8576200c7c8620101d0565b6020026020010151828260ff16815181106200c7e8576200c7e8620101d0565b60209081029190910101526001016200c7a1565b5093505050505b949350505050565b6200c83260405180606001604052805f60030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d925f91849163d145736c916200c8b99186910162010390565b5f60405180830381865afa1580156200c8d4573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200c8fd91908101906200fea7565b90505f6200c90c86836200dfb3565b90505f846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b81526004016200c93d91906200f3ee565b5f604051808303815f875af11580156200c959573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200c9829190810190620103d8565b805190915060030b158015906200c99c5750602081015151155b80156200c9ac5750604081015151155b156200b8bf57815f815181106200c9c7576200c9c7620101d0565b60200260200101516040516020016200aee5919062010493565b60605f6200ca15856040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925286518252808701908201529091506200ca4d9082905b906200e11b565b156200cbb4575f6200cace826200cac7846200cac06200ca938a6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925282518252918201519181019190915290565b906200e143565b906200e1a7565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200cb339082906200e11b565b156200cb9f57604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200cb9c905b82906200e233565b90505b6200cbaa816200e25a565b925050506200ae01565b82156200cbd05784846040516020016200aee592919062010672565b505060408051602081019091525f81526200ae01565b5f808251602084015ff09392505050565b8160a00151156200cc0757505050565b5f6200cc158484846200e2c5565b90505f6200cc23826200c80b565b602081015181519192509060030b1580156200ccc05750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200ccc0906040805180820182525f808252602091820152815180830190925284518252808501908201526200ca46565b156200ccce57505050505050565b604082015151156200ccf15781604001516040516020016200aee59190620106fd565b806040516020016200aee5919062010756565b60605f6200cd38836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200cd9e905b82906200d392565b156200ce1157604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200ae01906200ce0b9083906200e8b9565b6200e25a565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200ce74905b82906200e94b565b6001036200cf4557604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200cedc906200cb94565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200ae01906200ce0b905b83906200e233565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200cfa5906200cd96565b156200d0e457604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201819052845180860190955292518452830152906200d00f9083906200e9f1565b90505f81600183516200d02391906200f9d8565b815181106200d036576200d036620101d0565b602002602001015190506200d0db6200ce0b6200d0ae6040518060400160405280600581526020017f2e6a736f6e0000000000000000000000000000000000000000000000000000008152506040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f80825260209182015281518083019092528551825280860190820152906200e8b9565b95945050505050565b826040516020016200aee59190620107af565b50919050565b60605f6200d131836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200d194906200cd96565b156200d1a5576200ae01816200e25a565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d205906200ce6c565b6001036200d27257604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200ae01906200ce0b906200cf3d565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d2d2906200cd96565b156200d0e457604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201819052845180860190955292518452830152906200d33c9083906200e9f1565b90506001815111156200d37e5780600282516200d35a91906200f9d8565b815181106200d36d576200d36d620101d0565b602002602001015192505050919050565b50826040516020016200aee59190620107af565b805182515f9111156200d3a757505f6200ad82565b8151835160208501515f92916200d3be916200f52d565b6200d3ca91906200f9d8565b9050826020015181036200d3e35760019150506200ad82565b82516020840151819020912014905092915050565b60605f6200d406836200eaad565b60010190505f8167ffffffffffffffff8111156200d428576200d4286200f7aa565b6040519080825280601f01601f1916602001820160405280156200d453576020820181803683370190505b5090508181016020015b5f19017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846200d45d57509392505050565b604081810151815180830183525f808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e53454400000000000000000000000000000000000000000000818401908152855180870187528381528401929092528451808601909552518452908301526060916200d529905b82906200eb95565b156200d56a57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d5ca906200d521565b156200d60b57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d66b906200d521565b156200d6ac57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d70c906200d521565b806200d7735750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d773906200d521565b156200d7b457505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d814906200d521565b806200d87b5750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d87b906200d521565b156200d8bc57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d91c906200d521565b806200d9835750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d983906200d521565b156200d9c457505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200da24906200d521565b806200da8b5750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200da8b906200d521565b156200dacc57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200db2c906200d521565b156200db6d57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200dbcd906200d521565b156200dc0e57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200dc6e906200d521565b156200dcaf57505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200dd0f906200d521565b156200dd5057505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200ddb0906200d521565b156200ddf157505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200de51906200d521565b806200deb85750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200deb8906200d521565b156200def957505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200df59906200d521565b156200df9a57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516200aee5929060200162010881565b6060805f5b84518110156200e04957818582815181106200dfd8576200dfd8620101d0565b60200260200101516040516020016200dff39291906200fdb3565b6040516020818303038152906040529150600185516200e01491906200f9d8565b81146200e04057816040516020016200e02e9190620109d3565b60405160208183030381529060405291505b6001016200dfb8565b50604080516003808252608082019092525f91816020015b60608152602001906001900390816200e06157905050905083815f815181106200e08f576200e08f620101d0565b60200260200101819052506040518060400160405280600281526020017f2d63000000000000000000000000000000000000000000000000000000000000815250816001815181106200e0e6576200e0e6620101d0565b602002602001018190525081816002815181106200e108576200e108620101d0565b6020908102919091010152949350505050565b60208083015183518351928401515f936200e13a92918491906200ebaa565b14159392505050565b604080518082019091525f80825260208201525f6200e173845f01518560200151855f015186602001516200ecdd565b90508360200151816200e18791906200f9d8565b845185906200e1989083906200f9d8565b90525060208401525090919050565b604080518082019091525f80825260208201528151835110156200e1cd5750816200ad82565b60208083015190840151600191146200e1f55750815160208481015190840151829020919020145b80156200e22b578251845185906200e20f9083906200f9d8565b90525082516020850180516200e2279083906200f52d565b9052505b509192915050565b604080518082019091525f80825260208201526200e2538383836200ee1d565b5092915050565b60605f825f015167ffffffffffffffff8111156200e27c576200e27c6200f7aa565b6040519080825280601f01601f1916602001820160405280156200e2a7576020820181803683370190505b5090505f6020820190506200e253818560200151865f01516200eed1565b60605f6200e2d26200b118565b6040805160ff80825261200082019092529192505f9190816020015b60608152602001906001900390816200e2ee5790505090505f6040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806200e34b9062010316565b935060ff16815181106200e363576200e363620101d0565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e33000000000000000000000000000000000000000000000000008152506040516020016200e3b6919062010a0d565b6040516020818303038152906040528282806200e3d39062010316565b935060ff16815181106200e3eb576200e3eb620101d0565b60200260200101819052506040518060400160405280600881526020017f76616c69646174650000000000000000000000000000000000000000000000008152508282806200e43a9062010316565b935060ff16815181106200e452576200e452620101d0565b6020026020010181905250826040516020016200e47091906201025d565b6040516020818303038152906040528282806200e48d9062010316565b935060ff16815181106200e4a5576200e4a5620101d0565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e7472616374000000000000000000000000000000000000000000008152508282806200e4f49062010316565b935060ff16815181106200e50c576200e50c620101d0565b60200260200101819052506200e52387846200ef59565b82826200e5308162010316565b935060ff16815181106200e548576200e548620101d0565b6020908102919091010152855151156200e6005760408051808201909152600b81527f2d2d7265666572656e6365000000000000000000000000000000000000000000602082015282826200e59d8162010316565b935060ff16815181106200e5b5576200e5b5620101d0565b60200260200101819052506200e5cf865f0151846200ef59565b82826200e5dc8162010316565b935060ff16815181106200e5f4576200e5f4620101d0565b60200260200101819052505b8560800151156200e6755760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b0000000000000000602082015282826200e64c8162010316565b935060ff16815181106200e664576200e664620101d0565b60200260200101819052506200e6e1565b84156200e6e15760408051808201909152601281527f2d2d726571756972655265666572656e63650000000000000000000000000000602082015282826200e6bd8162010316565b935060ff16815181106200e6d5576200e6d5620101d0565b60200260200101819052505b604086015151156200e7885760408051808201909152600d81527f2d2d756e73616665416c6c6f7700000000000000000000000000000000000000602082015282826200e72e8162010316565b935060ff16815181106200e746576200e746620101d0565b602002602001018190525085604001518282806200e7649062010316565b935060ff16815181106200e77c576200e77c620101d0565b60200260200101819052505b8560600151156200e7f85760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826200e7d48162010316565b935060ff16815181106200e7ec576200e7ec620101d0565b60200260200101819052505b5f8160ff1667ffffffffffffffff8111156200e818576200e8186200f7aa565b6040519080825280602002602001820160405280156200e84d57816020015b60608152602001906001900390816200e8375790505b5090505f5b8260ff168160ff1610156200e8ad57838160ff16815181106200e879576200e879620101d0565b6020026020010151828260ff16815181106200e899576200e899620101d0565b60209081029190910101526001016200e852565b50979650505050505050565b604080518082019091525f80825260208201528151835110156200e8df5750816200ad82565b8151835160208501515f92916200e8f6916200f52d565b6200e90291906200f9d8565b602084015190915060019082146200e924575082516020840151819020908220145b80156200e942578351855186906200e93e9083906200f9d8565b9052505b50929392505050565b5f80825f01516200e96d855f01518660200151865f015187602001516200ecdd565b6200e97991906200f52d565b90505b835160208501516200e98f91906200f52d565b81116200e25357816200e9a28162010a40565b925050825f01516200e9dd8560200151836200e9bf91906200f9d8565b86516200e9cd91906200f9d8565b83865f015187602001516200ecdd565b6200e9e991906200f52d565b90506200e97c565b60605f6200ea0084846200e94b565b6200ea0d9060016200f52d565b67ffffffffffffffff8111156200ea28576200ea286200f7aa565b6040519080825280602002602001820160405280156200ea5d57816020015b60608152602001906001900390816200ea475790505b5090505f5b81518110156200eaa5576200ea7c6200ce0b86866200e233565b8282815181106200ea91576200ea91620101d0565b60209081029190910101526001016200ea62565b509392505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106200eaf6577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106200eb23576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106200eb4257662386f26fc10000830492506010015b6305f5e10083106200eb5b576305f5e100830492506008015b61271083106200eb7057612710830492506004015b606483106200eb83576064830492506002015b600a83106200ad825760010192915050565b5f6200eba283836200ef9c565b159392505050565b5f808584116200ecd357602084116200ec73575f84156200ebfe5760016200ebd48660206200f9d8565b6200ebe19060086200f9be565b6200ebee90600262010b57565b6200ebfa91906200f9d8565b1990505b83518116856200ec0f89896200f52d565b6200ec1b91906200f9d8565b805190935082165b8181146200ec5b578784116200ec4057879450505050506200c803565b836200ec4c8162010b64565b9450508284511690506200ec23565b6200ec6787856200f52d565b9450505050506200c803565b8383206200ec8285886200f9d8565b6200ec8e90876200f52d565b91505b8582106200ecd1578482208082036200ecbb576200ecb086846200f52d565b93505050506200c803565b6200ecc86001846200f9d8565b9250506200ec91565b505b5092949350505050565b5f83818685116200ee0657602085116200edac575f85156200ed325760016200ed088760206200f9d8565b6200ed159060086200f9be565b6200ed2290600262010b57565b6200ed2e91906200f9d8565b1990505b845181165f876200ed448b8b6200f52d565b6200ed5091906200f9d8565b855190915083165b8281146200ed9d578186106200ed82576200ed748b8b6200f52d565b96505050505050506200c803565b856200ed8e8162010a40565b9650508386511690506200ed58565b8596505050505050506200c803565b508383205f905b6200edbf86896200f9d8565b82116200ee04578583208082036200edde57839450505050506200c803565b6200edeb6001856200f52d565b93505081806200edfb9062010a40565b9250506200edb3565b505b6200ee1287876200f52d565b979650505050505050565b604080518082019091525f80825260208201525f6200ee4d855f01518660200151865f015187602001516200ecdd565b6020808701805191860191909152519091506200ee6b90826200f9d8565b8352845160208601516200ee8091906200f52d565b81036200ee90575f85526200eec8565b835183516200eea091906200f52d565b855186906200eeb19083906200f9d8565b90525083516200eec290826200f52d565b60208601525b50909392505050565b602081106200ef1157815183526200eeeb6020846200f52d565b92506200eefa6020836200f52d565b91506200ef096020826200f9d8565b90506200eed1565b5f1981156200ef465760016200ef298360206200f9d8565b6200ef379061010062010b57565b6200ef4391906200f9d8565b90505b9151835183169219169190911790915250565b60605f6200ef6884846200b1ec565b80516020808301516040519394506200ef849390910162010b7c565b60405160208183030381529060405291505092915050565b815181515f91908111156200efaf575081515b602080850151908401515f5b838110156200f07c57825182518082146200f045575f1960208710156200f022576001846200efec8960206200f9d8565b6200eff891906200f52d565b6200f0059060086200f9be565b6200f01290600262010b57565b6200f01e91906200f9d8565b1990505b81811683821681810391146200f0425797506200ad829650505050505050565b50505b6200f0526020866200f52d565b94506200f0616020856200f52d565b935050506020816200f07491906200f52d565b90506200efbb565b50845186516200b8bf919062010bbb565b610c328062010bde83390190565b61124c806201181083390190565b610c6d8062012a5c83390190565b6040518060e001604052806060815260200160608152602001606081526020015f151581526020015f151581526020015f151581526020016200f0f96200f0fe565b905290565b6040518061010001604052805f151581526020015f15158152602001606081526020015f801916815260200160608152602001606081526020015f151581526020016200f0f960405180608001604052805f81526020015f81526020015f81526020015f81525090565b602080825282518282018190525f918401906040840190835b818110156200f1aa5783516001600160a01b03168352602093840193909201916001016200f181565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200f2e2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b8801810191908801905f5b818110156200f2c7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a85030183526200f2b08486516200f1b5565b60209586019590945092909201916001016200f273565b5091975050506020948501949290920191506001016200f209565b50929695505050505050565b5f8151808452602084019350602083015f5b828110156200f3425781517fffffffff00000000000000000000000000000000000000000000000000000000168652602095860195909101906001016200f300565b5093949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200f2e2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051604087526200f3b960408801826200f1b5565b90506020820151915086810360208801526200f3d681836200f2ee565b9650505060209384019391909101906001016200f372565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200f2e2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184526200f4518583516200f1b5565b945060209384019391909101906001016200f414565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200f2e2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b03815116865260208101519050604060208701526200f4e960408701826200f2ee565b95505060209384019391909101906001016200f48d565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156200ad82576200ad826200f500565b602081525f6200ae0160208301846200f1b5565b6001600160a01b0381511682526020810151151560208301526001600160a01b0360408201511660408301525f606082015160a060608501526200f59f60a08501826200f1b5565b608093840151949093019390935250919050565b6001600160a01b0384168152606060208201525f6200f5d660608301856200f1b5565b82810360408401526200b8bf81856200f557565b600181811c908216806200f5ff57607f821691505b6020821081036200d0f7577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f81546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160a060608501525f81546200f681816200f5ea565b8060a0880152600182165f81146200f6a257600181146200f6dd576200f710565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660c089015260c082151560051b89010193506200f710565b845f5260205f205f5b838110156200f7075781548a820160c001526001909101906020016200f6e6565b890160c0019450505b50505060038401546080860152809250505092915050565b8381526001600160a01b0383166020820152608060408201525f608082015260a060608201525f6200d0db60a08301846200f637565b6001600160a01b0383168152604060208201525f6200c80360408301846200f637565b6001600160a01b0384168152826020820152606060408201525f6200d0db60608301846200f637565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b601f82111562000d4157805f5260205f20601f840160051c810160208510156200f7fe5750805b601f840160051c820191505b8181101562002dfc575f81556001016200f80a565b815167ffffffffffffffff8111156200f83c576200f83c6200f7aa565b6200f854816200f84d84546200f5ea565b846200f7d7565b6020601f8211600181146200f889575f83156200f8715750848201515b5f19600385901b1c1916600184901b17845562002dfc565b5f84815260208120601f198516915b828110156200f8ba57878501518255602094850194600190920191016200f898565b50848210156200f8d857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f602082840312156200f8f8575f80fd5b5051919050565b5f602082840312156200f910575f80fd5b815180151581146200ae01575f80fd5b6001600160a01b03851681528360208201526001600160a01b0383166040820152608060608201525f6200b8bf60808301846200f637565b8481526001600160a01b0384166020820152608060408201525f6200f98160808301856200f1b5565b82810360608401526200ee1281856200f637565b6001600160a01b0385168152836020820152608060408201525f6200f98160808301856200f1b5565b80820281158282048414176200ad82576200ad826200f500565b818103818111156200ad82576200ad826200f500565b6001600160a01b03851681528360208201526001600160a01b0383166040820152608060608201525f6200b8bf60808301846200f557565b6001600160a01b0383168152604060208201525f6200c80360408301846200f1b5565b6001600160a01b03861681528460208201526001600160a01b038416604082015260a060608201525f6200fa8160a08301856200f1b5565b82810360808401526200fa9581856200f637565b98975050505050505050565b6001600160a01b0384168152606060208201525f6200fac460608301856200f1b5565b82810360408401526200b8bf81856200f637565b6001600160a01b03861681528460208201526001600160a01b038416604082015260a060608201525f6200fb1060a08301856200f1b5565b82810360808401526200fa9581856200f557565b5f826200fb58577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b6001600160a01b0385168152836020820152608060408201525f6200fb8660808301856200f1b5565b82810360608401526200ee1281856200f557565b604081525f6200fbae60408301856200f1b5565b82810360208401526200adfd81856200f637565b6001600160a01b0384168152826020820152606060408201525f6200d0db60608301846200f557565b6001600160a01b0383168152604060208201525f6200c80360408301846200f557565b5f81518060208401855e5f93019283525090919050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f6200fc58601a8301856200fc0e565b7f3a2000000000000000000000000000000000000000000000000000000000000081526200adfd60028201856200fc0e565b5f602082840312156200fc9b575f80fd5b81516001600160a01b03811681146200ae01575f80fd5b6040516060810167ffffffffffffffff811182821017156200fcd8576200fcd86200f7aa565b60405290565b5f8067ffffffffffffffff8411156200fcfb576200fcfb6200f7aa565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff821117156200fd2d576200fd2d6200f7aa565b6040528381529050808284018510156200fd45575f80fd5b8383602083015e5f60208583010152509392505050565b5f82601f8301126200fd6c575f80fd5b6200ae01838351602085016200fcde565b5f602082840312156200fd8e575f80fd5b815167ffffffffffffffff8111156200fda5575f80fd5b6200ad7e848285016200fd5c565b5f6200c8036200fdc483866200fc0e565b846200fc0e565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f6200fdfe601a8301856200fc0e565b7f207573696e6720636f6e7374727563746f72206461746120220000000000000081526200fe3060198201856200fc0e565b7f2200000000000000000000000000000000000000000000000000000000000000815260010195945050505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201525f6200ae0160808301846200f1b5565b5f602082840312156200feb8575f80fd5b815167ffffffffffffffff8111156200fecf575f80fd5b8201601f810184136200fee0575f80fd5b6200ad7e848251602084016200fcde565b5f6200fefe82876200fc0e565b7f2f0000000000000000000000000000000000000000000000000000000000000081526200ff3060018201876200fc0e565b90507f2f0000000000000000000000000000000000000000000000000000000000000081526200ff6460018201866200fc0e565b90507f2f0000000000000000000000000000000000000000000000000000000000000081526200ff9860018201856200fc0e565b7f2e6a736f6e0000000000000000000000000000000000000000000000000000008152600501979650505050505050565b604081525f6200ffdd60408301846200f1b5565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081525f6201004f601f8301846200fc0e565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f81527f6d6c00000000000000000000000000000000000000000000000000000000000060208201526022019392505050565b604081525f620100b660408301846200f1b5565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081525f6201010960408301846200f1b5565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081525f6201017b60148301846200fc0e565b7f275d2e6b656363616b32353600000000000000000000000000000000000000008152600c019392505050565b604081525f620101bc60408301856200f1b5565b82810360208401526200adfd81856200f1b5565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f220000000000000000000000000000000000000000000000000000000000000081525f6201023060018301846200fc0e565b7f220000000000000000000000000000000000000000000000000000000000000081526001019392505050565b5f6201026a82846200fc0e565b7f2f6275696c642d696e666f0000000000000000000000000000000000000000008152600b019392505050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201525f6200ae01604b8301846200fc0e565b5f60ff821660ff81036201032e576201032e6200f500565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201525f6200ae0160298301846200fc0e565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201525f6200ae0160808301846200f1b5565b5f60208284031215620103e9575f80fd5b815167ffffffffffffffff81111562010400575f80fd5b82016060818503121562010412575f80fd5b6201041c6200fcb2565b81518060030b81146201042d575f80fd5b8152602082015167ffffffffffffffff81111562010449575f80fd5b62010457868285016200fd5c565b602083015250604082015167ffffffffffffffff81111562010477575f80fd5b62010485868285016200fd5c565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201525f620104ec60218301846200fc0e565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657481527f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960208201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560408201527f616c69666965642070617468206f66207468652062617368206578656375746160608201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960808201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60a08201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960c08201527f6f75722070726f6a65637420287573696e6720666f727761726420736c61736860e08201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101008201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061012082015261013b019392505050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201525f620106cb60218301856200fc0e565b7f2720696e206f75747075743a200000000000000000000000000000000000000081526200adfd600d8201856200fc0e565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201525f6200ae0160298301846200fc0e565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201525f6200ae0160228301846200fc0e565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081525f620107e2600e8301846200fc0e565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e7472616381527f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e7360208201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e747260408201527f6163742e6a736f6e00000000000000000000000000000000000000000000000060608201526068019392505050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081525f620108b460188301856200fc0e565b7f20696e20000000000000000000000000000000000000000000000000000000008152620108e660048201856200fc0e565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f727465642081527f6c6963656e736520666f7220626c6f636b206578706c6f72657220766572696660208201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f60408201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c60608201527f206f7220736574207468652060736b69704c6963656e73655479706560206f7060808201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060a082015260b70195945050505050565b5f620109e082846200fc0e565b7f200000000000000000000000000000000000000000000000000000000000000081526001019392505050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081525f6200ae01601c8301846200fc0e565b5f5f19820362010a545762010a546200f500565b5060010190565b6001815b600184111562010a9c5780850481111562010a7e5762010a7e6200f500565b600184161562010a8d57908102905b60019390931c92800262010a5f565b935093915050565b5f8262010ab4575060016200ad82565b8162010ac257505f6200ad82565b816001811462010adb576002811462010ae65762010b06565b60019150506200ad82565b60ff84111562010afa5762010afa6200f500565b50506001821b6200ad82565b5060208310610133831016604e8410600b841016171562010b2b575081810a6200ad82565b62010b395f19848462010a5b565b805f190482111562010b4f5762010b4f6200f500565b029392505050565b5f6200ae01838362010aa4565b5f8162010b755762010b756200f500565b505f190190565b5f62010b8982856200fc0e565b7f3a0000000000000000000000000000000000000000000000000000000000000081526200adfd60018201856200fc0e565b8181035f8312801583831316838312821617156200e253576200e2536200f50056fe608060405234801561000f575f80fd5b50604051610c32380380610c3283398101604081905261002e916100f0565b8181600361003c83826101d9565b50600461004982826101d9565b5050505050610293565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610076575f80fd5b81516001600160401b0381111561008f5761008f610053565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100bd576100bd610053565b6040528181528382016020018510156100d4575f80fd5b8160208501602083015e5f918101602001919091529392505050565b5f8060408385031215610101575f80fd5b82516001600160401b03811115610116575f80fd5b61012285828601610067565b602085015190935090506001600160401b0381111561013f575f80fd5b61014b85828601610067565b9150509250929050565b600181811c9082168061016957607f821691505b60208210810361018757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101d457805f5260205f20601f840160051c810160208510156101b25750805b601f840160051c820191505b818110156101d1575f81556001016101be565b50505b505050565b81516001600160401b038111156101f2576101f2610053565b610206816102008454610155565b8461018d565b6020601f821160018114610238575f83156102215750848201515b5f19600385901b1c1916600184901b1784556101d1565b5f84815260208120601f198516915b828110156102675787850151825560209485019460019092019101610247565b508482101561028457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610992806102a05f395ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c806340c10f191161007257806395d89b411161005857806395d89b411461017c578063a9059cbb14610184578063dd62ed3e14610197575f80fd5b806340c10f191461013257806370a0823114610147575f80fd5b806318160ddd116100a257806318160ddd146100fe57806323b872dd14610110578063313ce56714610123575f80fd5b806306fdde03146100bd578063095ea7b3146100db575b5f80fd5b6100c56101dc565b6040516100d291906107a5565b60405180910390f35b6100ee6100e9366004610820565b61026c565b60405190151581526020016100d2565b6002545b6040519081526020016100d2565b6100ee61011e366004610848565b610285565b604051601281526020016100d2565b610145610140366004610820565b6102a8565b005b610102610155366004610882565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b6100c56102b6565b6100ee610192366004610820565b6102c5565b6101026101a53660046108a2565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101eb906108d3565b80601f0160208091040260200160405190810160405280929190818152602001828054610217906108d3565b80156102625780601f1061023957610100808354040283529160200191610262565b820191905f5260205f20905b81548152906001019060200180831161024557829003601f168201915b5050505050905090565b5f336102798185856102d2565b60019150505b92915050565b5f336102928582856102e4565b61029d8585856103b6565b506001949350505050565b6102b2828261045f565b5050565b6060600480546101eb906108d3565b5f336102798185856103b6565b6102df83838360016104b9565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103b057818110156103a2576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103b084848484035f6104b9565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610405576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b73ffffffffffffffffffffffffffffffffffffffff8216610454576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b6102df8383836105fe565b73ffffffffffffffffffffffffffffffffffffffff82166104ae576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b6102b25f83836105fe565b73ffffffffffffffffffffffffffffffffffffffff8416610508576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b73ffffffffffffffffffffffffffffffffffffffff8316610557576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b73ffffffffffffffffffffffffffffffffffffffff8085165f90815260016020908152604080832093871683529290522082905580156103b0578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105f091815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610635578060025f82825461062a9190610924565b909155506106e59050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260208190526040902054818110156106ba576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610399565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661070e57600280548290039055610739565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161079891815260200190565b60405180910390a3505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461081b575f80fd5b919050565b5f8060408385031215610831575f80fd5b61083a836107f8565b946020939093013593505050565b5f805f6060848603121561085a575f80fd5b610863846107f8565b9250610871602085016107f8565b929592945050506040919091013590565b5f60208284031215610892575f80fd5b61089b826107f8565b9392505050565b5f80604083850312156108b3575f80fd5b6108bc836107f8565b91506108ca602084016107f8565b90509250929050565b600181811c908216806108e757607f821691505b60208210810361091e577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b8082018082111561027f577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea2646970667358221220569a538c722848c143f241fcbfd3f113f81549aa32aa89f9a6169cd25cb5399e64736f6c634300081a0033608060405234801561000f575f80fd5b5060405161124c38038061124c83398101604081905261002e9161010e565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007791906101d7565b50600461008482826101d7565b5050506001600160a01b03821615806100a457506001600160a01b038116155b156100c25760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b03199182161790915560078054929093169116179055610291565b80516001600160a01b0381168114610109575f80fd5b919050565b5f806040838503121561011f575f80fd5b610128836100f3565b9150610136602084016100f3565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061016757607f821691505b60208210810361018557634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101d257805f5260205f20601f840160051c810160208510156101b05750805b601f840160051c820191505b818110156101cf575f81556001016101bc565b50505b505050565b81516001600160401b038111156101f0576101f061013f565b610204816101fe8454610153565b8461018b565b6020601f821160018114610236575f831561021f5750848201515b5f19600385901b1c1916600184901b1784556101cf565b5f84815260208120601f198516915b828110156102655787850151825560209485019460019092019101610245565b508482101561028257868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610fae8061029e5f395ff3fe608060405234801561000f575f80fd5b5060043610610115575f3560e01c806342966c68116100ad57806379cc67901161007d578063a9059cbb11610063578063a9059cbb14610286578063bff9662a14610299578063dd62ed3e146102b9575f80fd5b806379cc67901461026b57806395d89b411461027e575f80fd5b806342966c68146101fb5780635b1125911461020e57806370a082311461022e578063779e3b6314610263575f80fd5b80631e458bee116100e85780631e458bee1461018157806323b872dd14610194578063313ce567146101a7578063328a01d0146101b6575f80fd5b806306fdde0314610119578063095ea7b31461013757806315d57fd41461015a57806318160ddd1461016f575b5f80fd5b6101216102fe565b60405161012e9190610d7a565b60405180910390f35b61014a610145366004610df5565b61038e565b604051901515815260200161012e565b61016d610168366004610e1d565b6103a7565b005b6002545b60405190815260200161012e565b61016d61018f366004610e4e565b610572565b61014a6101a2366004610e7e565b610625565b6040516012815260200161012e565b6007546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012e565b61016d610209366004610eb8565b610648565b6006546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b61017361023c366004610ecf565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b61016d610655565b61016d610279366004610df5565b610779565b61012161082a565b61014a610294366004610df5565b610839565b6005546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b6101736102c7366004610e1d565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b60606003805461030d90610eef565b80601f016020809104026020016040519081016040528092919081815260200182805461033990610eef565b80156103845780601f1061035b57610100808354040283529160200191610384565b820191905f5260205f20905b81548152906001019060200180831161036757829003601f168201915b5050505050905090565b5f3361039b818585610846565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103e7575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610425576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216158061045c575073ffffffffffffffffffffffffffffffffffffffff8116155b15610493576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105c5576040517f3fe32fba00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b6105cf8383610858565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161061891815260200190565b60405180910390a3505050565b5f336106328582856108b6565b61063d858585610983565b506001949350505050565b6106523382610a2c565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106a8576040517fe700765e00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b60065473ffffffffffffffffffffffffffffffffffffffff166106f7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107cc576040517f3fe32fba00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b6107d68282610a86565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161081e91815260200190565b60405180910390a25050565b60606004805461030d90610eef565b5f3361039b818585610983565b6108538383836001610a9b565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108a7576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b6108b25f8383610be0565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461097d578181101561096f576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602481018290526044810183905260640161041c565b61097d84848484035f610a9b565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109d2576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8216610a21576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b610853838383610be0565b73ffffffffffffffffffffffffffffffffffffffff8216610a7b576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b6108b2825f83610be0565b610a918233836108b6565b6108b28282610a2c565b73ffffffffffffffffffffffffffffffffffffffff8416610aea576040517fe602df050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8316610b39576040517f94280d620000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8085165f908152600160209081526040808320938716835292905220829055801561097d578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bd291815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c17578060025f828254610c0c9190610f40565b90915550610cc79050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020819052604090205481811015610c9c576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602481018290526044810183905260640161041c565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610cf057600280548290039055610d1b565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161061891815260200190565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610df0575f80fd5b919050565b5f8060408385031215610e06575f80fd5b610e0f83610dcd565b946020939093013593505050565b5f8060408385031215610e2e575f80fd5b610e3783610dcd565b9150610e4560208401610dcd565b90509250929050565b5f805f60608486031215610e60575f80fd5b610e6984610dcd565b95602085013595506040909401359392505050565b5f805f60608486031215610e90575f80fd5b610e9984610dcd565b9250610ea760208501610dcd565b929592945050506040919091013590565b5f60208284031215610ec8575f80fd5b5035919050565b5f60208284031215610edf575f80fd5b610ee882610dcd565b9392505050565b600181811c90821680610f0357607f821691505b602082108103610f3a577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b808201808211156103a1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea26469706673582212206a41e2cd6fbf39f12e2917f13d39a938ebbfbe1bbc40009f596c89c4dc977dca64736f6c634300081a003360a0604052348015600e575f80fd5b50604051610c6d380380610c6d833981016040819052602b91603b565b6001600160a01b03166080526066565b5f60208284031215604a575f80fd5b81516001600160a01b0381168114605f575f80fd5b9392505050565b608051610bc06100ad5f395f818160470152818160e8015281816101b80152818161028101528181610361015281816104450152818161059001526106650152610bc05ff3fe60806040526004361061002b575f3560e01c8063116191b614610036578063d51240c414610092575f80fd5b3661003257005b5f80fd5b348015610041575f80fd5b506100697f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100a56100a03660046108b6565b6100a7565b005b6100e56040518060400160405280600481526020017f48495421000000000000000000000000000000000000000000000000000000008152506106d8565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b01072146040518163ffffffff1660e01b81526004016020604051808303815f875af1158015610150573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610174919061094e565b905061017f8161076a565b5f61018b826002610992565b61019590346109af565b90506101a181346107fb565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663726ac97c6101e86004846109c2565b6040805160a0810182523381525f6020808301829052828401829052835190810184528181526060830152608082015290517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152610251918a91600401610aba565b5f604051808303818588803b158015610268575f80fd5b505af115801561027a573d5f803e3d5ffd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663726ac97c6004836102c891906109c2565b6040805160a0810182523381525f6020808301829052828401829052835190810184528181526060830152608082015290517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152610331918a91600401610aba565b5f604051808303818588803b158015610348575f80fd5b505af115801561035a573d5f803e3d5ffd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663744b9b8b6004836103a891906109c2565b6040805160a0810182523381525f6020808301829052828401829052835190810184528181526060830152608082015290517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152610415918a918a918a91600401610af0565b5f604051808303818588803b15801561042c575f80fd5b505af115801561043e573d5f803e3d5ffd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663744b9b8b60048361048c91906109c2565b6040805160a0810182523381525f6020808301829052828401829052835190810184528181526060830152608082015290517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526104f9918a918a918a91600401610af0565b5f604051808303818588803b158015610510575f80fd5b505af1158015610522573d5f803e3d5ffd5b50506040805160a0810182523381525f6020808301829052828401829052835190810184528181526060830152608082015290517f1becceb400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169450631becceb493508692506105ce918a918a918a91600401610af0565b5f604051808303818588803b1580156105e5575f80fd5b505af11580156105f7573d5f803e3d5ffd5b50506040805160a0810182523381525f6020808301829052828401829052835190810184528181526060830152608082015290517f1becceb400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169450631becceb493508692506106a3918a918a918a91600401610af0565b5f604051808303818588803b1580156106ba575f80fd5b505af11580156106cc573d5f803e3d5ffd5b50505050505050505050565b610767816040516024016106ec9190610b71565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f41304fac00000000000000000000000000000000000000000000000000000000179052610892565b50565b6107678160405160240161078091815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff5b1bba900000000000000000000000000000000000000000000000000000000179052610892565b604051602481018390526044810182905261088e90606401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6c0f698000000000000000000000000000000000000000000000000000000000179052610892565b5050565b6107678180516a636f6e736f6c652e6c6f67602083015f808483855afa5050505050565b5f805f604084860312156108c8575f80fd5b833573ffffffffffffffffffffffffffffffffffffffff811681146108eb575f80fd5b9250602084013567ffffffffffffffff811115610906575f80fd5b8401601f81018613610916575f80fd5b803567ffffffffffffffff81111561092c575f80fd5b86602082840101111561093d575f80fd5b939660209190910195509293505050565b5f6020828403121561095e575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820281158282048414176109a9576109a9610965565b92915050565b818103818111156109a9576109a9610965565b5f826109f5577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff60408201511660408301525f606082015160a06060850152610aa660a08501826109fa565b608093840151949093019390935250919050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f610ae86040830184610a46565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff8516815260606020820152826060820152828460808301375f608084830101525f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011682016080838203016040840152610b666080820185610a46565b979650505050505050565b602081525f610b8360208301846109fa565b939250505056fea26469706673582212209836c527f0b29707020bfa435fb8a3969b017c288b14121f45b26581585074f864736f6c634300081a0033a26469706673582212206c40bf91299e1741ecc5e444c16e68087894f45f2a4d89fc9a249f12c60cc81464736f6c634300081a0033",
}

// GatewayEVMInboundTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMInboundTestMetaData.ABI instead.
var GatewayEVMInboundTestABI = GatewayEVMInboundTestMetaData.ABI

// GatewayEVMInboundTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMInboundTestMetaData.Bin instead.
var GatewayEVMInboundTestBin = GatewayEVMInboundTestMetaData.Bin

// DeployGatewayEVMInboundTest deploys a new Ethereum contract, binding an instance of GatewayEVMInboundTest to it.
func DeployGatewayEVMInboundTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMInboundTest, error) {
	parsed, err := GatewayEVMInboundTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMInboundTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMInboundTest{GatewayEVMInboundTestCaller: GatewayEVMInboundTestCaller{contract: contract}, GatewayEVMInboundTestTransactor: GatewayEVMInboundTestTransactor{contract: contract}, GatewayEVMInboundTestFilterer: GatewayEVMInboundTestFilterer{contract: contract}}, nil
}

// GatewayEVMInboundTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMInboundTest struct {
	GatewayEVMInboundTestCaller     // Read-only binding to the contract
	GatewayEVMInboundTestTransactor // Write-only binding to the contract
	GatewayEVMInboundTestFilterer   // Log filterer for contract events
}

// GatewayEVMInboundTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMInboundTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMInboundTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMInboundTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMInboundTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMInboundTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMInboundTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMInboundTestSession struct {
	Contract     *GatewayEVMInboundTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayEVMInboundTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMInboundTestCallerSession struct {
	Contract *GatewayEVMInboundTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// GatewayEVMInboundTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMInboundTestTransactorSession struct {
	Contract     *GatewayEVMInboundTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// GatewayEVMInboundTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMInboundTestRaw struct {
	Contract *GatewayEVMInboundTest // Generic contract binding to access the raw methods on
}

// GatewayEVMInboundTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMInboundTestCallerRaw struct {
	Contract *GatewayEVMInboundTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMInboundTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMInboundTestTransactorRaw struct {
	Contract *GatewayEVMInboundTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMInboundTest creates a new instance of GatewayEVMInboundTest, bound to a specific deployed contract.
func NewGatewayEVMInboundTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMInboundTest, error) {
	contract, err := bindGatewayEVMInboundTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTest{GatewayEVMInboundTestCaller: GatewayEVMInboundTestCaller{contract: contract}, GatewayEVMInboundTestTransactor: GatewayEVMInboundTestTransactor{contract: contract}, GatewayEVMInboundTestFilterer: GatewayEVMInboundTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMInboundTestCaller creates a new read-only instance of GatewayEVMInboundTest, bound to a specific deployed contract.
func NewGatewayEVMInboundTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMInboundTestCaller, error) {
	contract, err := bindGatewayEVMInboundTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestCaller{contract: contract}, nil
}

// NewGatewayEVMInboundTestTransactor creates a new write-only instance of GatewayEVMInboundTest, bound to a specific deployed contract.
func NewGatewayEVMInboundTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMInboundTestTransactor, error) {
	contract, err := bindGatewayEVMInboundTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestTransactor{contract: contract}, nil
}

// NewGatewayEVMInboundTestFilterer creates a new log filterer instance of GatewayEVMInboundTest, bound to a specific deployed contract.
func NewGatewayEVMInboundTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMInboundTestFilterer, error) {
	contract, err := bindGatewayEVMInboundTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestFilterer{contract: contract}, nil
}

// bindGatewayEVMInboundTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMInboundTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMInboundTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMInboundTest.Contract.GatewayEVMInboundTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.GatewayEVMInboundTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.GatewayEVMInboundTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMInboundTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.contract.Transact(opts, method, params...)
}

// ADDITIONALACTIONFEEWEI is a free data retrieval call binding the contract method 0xf8d37ef2.
//
// Solidity: function ADDITIONAL_ACTION_FEE_WEI() view returns(uint256)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ADDITIONALACTIONFEEWEI(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "ADDITIONAL_ACTION_FEE_WEI")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ADDITIONALACTIONFEEWEI is a free data retrieval call binding the contract method 0xf8d37ef2.
//
// Solidity: function ADDITIONAL_ACTION_FEE_WEI() view returns(uint256)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ADDITIONALACTIONFEEWEI() (*big.Int, error) {
	return _GatewayEVMInboundTest.Contract.ADDITIONALACTIONFEEWEI(&_GatewayEVMInboundTest.CallOpts)
}

// ADDITIONALACTIONFEEWEI is a free data retrieval call binding the contract method 0xf8d37ef2.
//
// Solidity: function ADDITIONAL_ACTION_FEE_WEI() view returns(uint256)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ADDITIONALACTIONFEEWEI() (*big.Int, error) {
	return _GatewayEVMInboundTest.Contract.ADDITIONALACTIONFEEWEI(&_GatewayEVMInboundTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ISTEST() (bool, error) {
	return _GatewayEVMInboundTest.Contract.ISTEST(&_GatewayEVMInboundTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ISTEST() (bool, error) {
	return _GatewayEVMInboundTest.Contract.ISTEST(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeArtifacts(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeArtifacts(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeContracts(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeContracts(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeSenders(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeSenders(&_GatewayEVMInboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) Failed() (bool, error) {
	return _GatewayEVMInboundTest.Contract.Failed(&_GatewayEVMInboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) Failed() (bool, error) {
	return _GatewayEVMInboundTest.Contract.Failed(&_GatewayEVMInboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMInboundTest.Contract.TargetArtifactSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMInboundTest.Contract.TargetArtifactSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMInboundTest.Contract.TargetArtifacts(&_GatewayEVMInboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMInboundTest.Contract.TargetArtifacts(&_GatewayEVMInboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.TargetContracts(&_GatewayEVMInboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.TargetContracts(&_GatewayEVMInboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMInboundTest.Contract.TargetInterfaces(&_GatewayEVMInboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMInboundTest.Contract.TargetInterfaces(&_GatewayEVMInboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMInboundTest.Contract.TargetSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMInboundTest.Contract.TargetSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.TargetSenders(&_GatewayEVMInboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.TargetSenders(&_GatewayEVMInboundTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.SetUp(&_GatewayEVMInboundTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.SetUp(&_GatewayEVMInboundTest.TransactOpts)
}

// TestAdditionalActionDisabledReverts is a paid mutator transaction binding the contract method 0xd86a4a0c.
//
// Solidity: function testAdditionalActionDisabledReverts() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestAdditionalActionDisabledReverts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testAdditionalActionDisabledReverts")
}

// TestAdditionalActionDisabledReverts is a paid mutator transaction binding the contract method 0xd86a4a0c.
//
// Solidity: function testAdditionalActionDisabledReverts() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestAdditionalActionDisabledReverts() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestAdditionalActionDisabledReverts(&_GatewayEVMInboundTest.TransactOpts)
}

// TestAdditionalActionDisabledReverts is a paid mutator transaction binding the contract method 0xd86a4a0c.
//
// Solidity: function testAdditionalActionDisabledReverts() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestAdditionalActionDisabledReverts() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestAdditionalActionDisabledReverts(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x33ed091c.
//
// Solidity: function testCallFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallFailsIfExcessEthProvided(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallFailsIfExcessEthProvided")
}

// TestCallFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x33ed091c.
//
// Solidity: function testCallFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x33ed091c.
//
// Solidity: function testCallFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x598b7edb.
//
// Solidity: function testCallFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallFailsIfInsufficientFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallFailsIfInsufficientFee")
}

// TestCallFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x598b7edb.
//
// Solidity: function testCallFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x598b7edb.
//
// Solidity: function testCallFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallFailsIfRevertGasLimitExceeded")
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallSecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xc51a4cbe.
//
// Solidity: function testCallSecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallSecondActionFailsIfExcessEthProvided(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallSecondActionFailsIfExcessEthProvided")
}

// TestCallSecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xc51a4cbe.
//
// Solidity: function testCallSecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallSecondActionFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallSecondActionFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallSecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xc51a4cbe.
//
// Solidity: function testCallSecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallSecondActionFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallSecondActionFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x481daadb.
//
// Solidity: function testCallSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallSecondActionRequiresFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallSecondActionRequiresFee")
}

// TestCallSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x481daadb.
//
// Solidity: function testCallSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallSecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallSecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x481daadb.
//
// Solidity: function testCallSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallSecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallSecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayload is a paid mutator transaction binding the contract method 0xaa030c1c.
//
// Solidity: function testCallWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallWithPayload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallWithPayload")
}

// TestCallWithPayload is a paid mutator transaction binding the contract method 0xaa030c1c.
//
// Solidity: function testCallWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayload is a paid mutator transaction binding the contract method 0xaa030c1c.
//
// Solidity: function testCallWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfCallOnRevertIsTrue is a paid mutator transaction binding the contract method 0xba46ba23.
//
// Solidity: function testCallWithPayloadFailsIfCallOnRevertIsTrue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallWithPayloadFailsIfCallOnRevertIsTrue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallWithPayloadFailsIfCallOnRevertIsTrue")
}

// TestCallWithPayloadFailsIfCallOnRevertIsTrue is a paid mutator transaction binding the contract method 0xba46ba23.
//
// Solidity: function testCallWithPayloadFailsIfCallOnRevertIsTrue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallWithPayloadFailsIfCallOnRevertIsTrue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfCallOnRevertIsTrue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfCallOnRevertIsTrue is a paid mutator transaction binding the contract method 0xba46ba23.
//
// Solidity: function testCallWithPayloadFailsIfCallOnRevertIsTrue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallWithPayloadFailsIfCallOnRevertIsTrue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfCallOnRevertIsTrue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x51da903d.
//
// Solidity: function testCallWithPayloadFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallWithPayloadFailsIfDestinationIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallWithPayloadFailsIfDestinationIsZeroAddress")
}

// TestCallWithPayloadFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x51da903d.
//
// Solidity: function testCallWithPayloadFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallWithPayloadFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfDestinationIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x51da903d.
//
// Solidity: function testCallWithPayloadFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallWithPayloadFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfDestinationIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xb1409f71.
//
// Solidity: function testCallWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallWithPayloadFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallWithPayloadFailsIfPayloadSizeExceeded")
}

// TestCallWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xb1409f71.
//
// Solidity: function testCallWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xb1409f71.
//
// Solidity: function testCallWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20SecondActionRequiresFee is a paid mutator transaction binding the contract method 0x6ab1c516.
//
// Solidity: function testDepositAndCallERC20SecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallERC20SecondActionRequiresFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallERC20SecondActionRequiresFee")
}

// TestDepositAndCallERC20SecondActionRequiresFee is a paid mutator transaction binding the contract method 0x6ab1c516.
//
// Solidity: function testDepositAndCallERC20SecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallERC20SecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20SecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20SecondActionRequiresFee is a paid mutator transaction binding the contract method 0x6ab1c516.
//
// Solidity: function testDepositAndCallERC20SecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallERC20SecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20SecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xf0a635c5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallERC20ToCustodyFailsIfExcessEthProvided")
}

// TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xf0a635c5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xf0a635c5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x49050a44.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallERC20ToCustodyFailsIfInsufficientFee")
}

// TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x49050a44.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x49050a44.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x458085f5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded")
}

// TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x458085f5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x458085f5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x48452329.
//
// Solidity: function testDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided")
}

// TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x48452329.
//
// Solidity: function testDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x48452329.
//
// Solidity: function testDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc0fab86d.
//
// Solidity: function testDepositAndCallEthFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthFailsIfRevertGasLimitExceeded")
}

// TestDepositAndCallEthFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc0fab86d.
//
// Solidity: function testDepositAndCallEthFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc0fab86d.
//
// Solidity: function testDepositAndCallEthFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthToTssFailsForSubsequentActions is a paid mutator transaction binding the contract method 0x7e7dfee3.
//
// Solidity: function testDepositAndCallEthToTssFailsForSubsequentActions() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthToTssFailsForSubsequentActions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthToTssFailsForSubsequentActions")
}

// TestDepositAndCallEthToTssFailsForSubsequentActions is a paid mutator transaction binding the contract method 0x7e7dfee3.
//
// Solidity: function testDepositAndCallEthToTssFailsForSubsequentActions() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthToTssFailsForSubsequentActions() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthToTssFailsForSubsequentActions(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthToTssFailsForSubsequentActions is a paid mutator transaction binding the contract method 0x7e7dfee3.
//
// Solidity: function testDepositAndCallEthToTssFailsForSubsequentActions() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthToTssFailsForSubsequentActions() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthToTssFailsForSubsequentActions(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x478095e5.
//
// Solidity: function testDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded")
}

// TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x478095e5.
//
// Solidity: function testDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x478095e5.
//
// Solidity: function testDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x1b906ef3.
//
// Solidity: function testDepositAndCallEthWithAmountSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthWithAmountSecondActionRequiresFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthWithAmountSecondActionRequiresFee")
}

// TestDepositAndCallEthWithAmountSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x1b906ef3.
//
// Solidity: function testDepositAndCallEthWithAmountSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthWithAmountSecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountSecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x1b906ef3.
//
// Solidity: function testDepositAndCallEthWithAmountSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthWithAmountSecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountSecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x8be96f5c.
//
// Solidity: function testDepositAndCallEthWithAmountToTssFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthWithAmountToTssFailsIfInsufficientFee")
}

// TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x8be96f5c.
//
// Solidity: function testDepositAndCallEthWithAmountToTssFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x8be96f5c.
//
// Solidity: function testDepositAndCallEthWithAmountToTssFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue is a paid mutator transaction binding the contract method 0x2cf9572d.
//
// Solidity: function testDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue")
}

// TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue is a paid mutator transaction binding the contract method 0x2cf9572d.
//
// Solidity: function testDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue is a paid mutator transaction binding the contract method 0x2cf9572d.
//
// Solidity: function testDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20SecondActionRequiresFee is a paid mutator transaction binding the contract method 0xdc23a35f.
//
// Solidity: function testDepositERC20SecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20SecondActionRequiresFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20SecondActionRequiresFee")
}

// TestDepositERC20SecondActionRequiresFee is a paid mutator transaction binding the contract method 0xdc23a35f.
//
// Solidity: function testDepositERC20SecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20SecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20SecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20SecondActionRequiresFee is a paid mutator transaction binding the contract method 0xdc23a35f.
//
// Solidity: function testDepositERC20SecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20SecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20SecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustody is a paid mutator transaction binding the contract method 0x6459542a.
//
// Solidity: function testDepositERC20ToCustody() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustody(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustody")
}

// TestDepositERC20ToCustody is a paid mutator transaction binding the contract method 0x6459542a.
//
// Solidity: function testDepositERC20ToCustody() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustody() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustody(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustody is a paid mutator transaction binding the contract method 0x6459542a.
//
// Solidity: function testDepositERC20ToCustody() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustody() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustody(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x86b6e429.
//
// Solidity: function testDepositERC20ToCustodyFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfAmountIs0")
}

// TestDepositERC20ToCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x86b6e429.
//
// Solidity: function testDepositERC20ToCustodyFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x86b6e429.
//
// Solidity: function testDepositERC20ToCustodyFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x15ee8f36.
//
// Solidity: function testDepositERC20ToCustodyFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfExcessEthProvided(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfExcessEthProvided")
}

// TestDepositERC20ToCustodyFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x15ee8f36.
//
// Solidity: function testDepositERC20ToCustodyFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x15ee8f36.
//
// Solidity: function testDepositERC20ToCustodyFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0xc7a1eccb.
//
// Solidity: function testDepositERC20ToCustodyFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfInsufficientFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfInsufficientFee")
}

// TestDepositERC20ToCustodyFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0xc7a1eccb.
//
// Solidity: function testDepositERC20ToCustodyFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0xc7a1eccb.
//
// Solidity: function testDepositERC20ToCustodyFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x4ce25c0a.
//
// Solidity: function testDepositERC20ToCustodyFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfPayloadSizeExceeded")
}

// TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x4ce25c0a.
//
// Solidity: function testDepositERC20ToCustodyFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x4ce25c0a.
//
// Solidity: function testDepositERC20ToCustodyFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x846b9f7f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfReceiverIsZeroAddress")
}

// TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x846b9f7f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x846b9f7f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x1f4f542f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded")
}

// TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x1f4f542f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x1f4f542f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xb2849063.
//
// Solidity: function testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted")
}

// TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xb2849063.
//
// Solidity: function testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xb2849063.
//
// Solidity: function testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xb966e8fa.
//
// Solidity: function testDepositERC20ToCustodySecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodySecondActionFailsIfExcessEthProvided")
}

// TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xb966e8fa.
//
// Solidity: function testDepositERC20ToCustodySecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xb966e8fa.
//
// Solidity: function testDepositERC20ToCustodySecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayload is a paid mutator transaction binding the contract method 0x30f7c04f.
//
// Solidity: function testDepositERC20ToCustodyWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayload")
}

// TestDepositERC20ToCustodyWithPayload is a paid mutator transaction binding the contract method 0x30f7c04f.
//
// Solidity: function testDepositERC20ToCustodyWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayload is a paid mutator transaction binding the contract method 0x30f7c04f.
//
// Solidity: function testDepositERC20ToCustodyWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xf75fc969.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayloadFailsIfAmountIs0")
}

// TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xf75fc969.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xf75fc969.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x95cd0445.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded")
}

// TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x95cd0445.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x95cd0445.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x895c2bc6.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress")
}

// TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x895c2bc6.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x895c2bc6.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xe85c5a07.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted")
}

// TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xe85c5a07.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xe85c5a07.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTss is a paid mutator transaction binding the contract method 0x0724d8e3.
//
// Solidity: function testDepositEthToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTss(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTss")
}

// TestDepositEthToTss is a paid mutator transaction binding the contract method 0x0724d8e3.
//
// Solidity: function testDepositEthToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTss() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTss(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTss is a paid mutator transaction binding the contract method 0x0724d8e3.
//
// Solidity: function testDepositEthToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTss() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTss(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsForSubsequentActions is a paid mutator transaction binding the contract method 0x91a5c818.
//
// Solidity: function testDepositEthToTssFailsForSubsequentActions() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssFailsForSubsequentActions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssFailsForSubsequentActions")
}

// TestDepositEthToTssFailsForSubsequentActions is a paid mutator transaction binding the contract method 0x91a5c818.
//
// Solidity: function testDepositEthToTssFailsForSubsequentActions() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssFailsForSubsequentActions() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsForSubsequentActions(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsForSubsequentActions is a paid mutator transaction binding the contract method 0x91a5c818.
//
// Solidity: function testDepositEthToTssFailsForSubsequentActions() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssFailsForSubsequentActions() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsForSubsequentActions(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x7bb46d46.
//
// Solidity: function testDepositEthToTssFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssFailsIfAmountIs0")
}

// TestDepositEthToTssFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x7bb46d46.
//
// Solidity: function testDepositEthToTssFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x7bb46d46.
//
// Solidity: function testDepositEthToTssFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x8aeeb7e7.
//
// Solidity: function testDepositEthToTssFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssFailsIfReceiverIsZeroAddress")
}

// TestDepositEthToTssFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x8aeeb7e7.
//
// Solidity: function testDepositEthToTssFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x8aeeb7e7.
//
// Solidity: function testDepositEthToTssFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xb1526934.
//
// Solidity: function testDepositEthToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssFailsIfRevertGasLimitExceeded")
}

// TestDepositEthToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xb1526934.
//
// Solidity: function testDepositEthToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xb1526934.
//
// Solidity: function testDepositEthToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayload is a paid mutator transaction binding the contract method 0x9fd1e597.
//
// Solidity: function testDepositEthToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssWithPayload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssWithPayload")
}

// TestDepositEthToTssWithPayload is a paid mutator transaction binding the contract method 0x9fd1e597.
//
// Solidity: function testDepositEthToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayload is a paid mutator transaction binding the contract method 0x9fd1e597.
//
// Solidity: function testDepositEthToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x84c59bcf.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssWithPayloadFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssWithPayloadFailsIfAmountIs0")
}

// TestDepositEthToTssWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x84c59bcf.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssWithPayloadFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x84c59bcf.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssWithPayloadFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x466f332e.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded")
}

// TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x466f332e.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x466f332e.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x53c9a0a3.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress")
}

// TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x53c9a0a3.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x53c9a0a3.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountSecondActionFailsWithOnlyFee is a paid mutator transaction binding the contract method 0x32fc1fad.
//
// Solidity: function testDepositEthWithAmountSecondActionFailsWithOnlyFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountSecondActionFailsWithOnlyFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountSecondActionFailsWithOnlyFee")
}

// TestDepositEthWithAmountSecondActionFailsWithOnlyFee is a paid mutator transaction binding the contract method 0x32fc1fad.
//
// Solidity: function testDepositEthWithAmountSecondActionFailsWithOnlyFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountSecondActionFailsWithOnlyFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountSecondActionFailsWithOnlyFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountSecondActionFailsWithOnlyFee is a paid mutator transaction binding the contract method 0x32fc1fad.
//
// Solidity: function testDepositEthWithAmountSecondActionFailsWithOnlyFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountSecondActionFailsWithOnlyFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountSecondActionFailsWithOnlyFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x3424914f.
//
// Solidity: function testDepositEthWithAmountSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountSecondActionRequiresFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountSecondActionRequiresFee")
}

// TestDepositEthWithAmountSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x3424914f.
//
// Solidity: function testDepositEthWithAmountSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountSecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountSecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x3424914f.
//
// Solidity: function testDepositEthWithAmountSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountSecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountSecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTss is a paid mutator transaction binding the contract method 0xf2036eda.
//
// Solidity: function testDepositEthWithAmountToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTss(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTss")
}

// TestDepositEthWithAmountToTss is a paid mutator transaction binding the contract method 0xf2036eda.
//
// Solidity: function testDepositEthWithAmountToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTss() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTss(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTss is a paid mutator transaction binding the contract method 0xf2036eda.
//
// Solidity: function testDepositEthWithAmountToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTss() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTss(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xc57926c6.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfAmountIs0")
}

// TestDepositEthWithAmountToTssFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xc57926c6.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xc57926c6.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue is a paid mutator transaction binding the contract method 0x6aa02e8b.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue")
}

// TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue is a paid mutator transaction binding the contract method 0x6aa02e8b.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue is a paid mutator transaction binding the contract method 0x6aa02e8b.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue is a paid mutator transaction binding the contract method 0x51ee53cb.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue")
}

// TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue is a paid mutator transaction binding the contract method 0x51ee53cb.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue is a paid mutator transaction binding the contract method 0x51ee53cb.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0xf6e53a5d.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfInsufficientFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfInsufficientFee")
}

// TestDepositEthWithAmountToTssFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0xf6e53a5d.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0xf6e53a5d.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x545c3745.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfPayloadSizeExceeded")
}

// TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x545c3745.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x545c3745.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9acda9fa.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress")
}

// TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9acda9fa.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9acda9fa.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x88343a41.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded")
}

// TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x88343a41.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x88343a41.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue is a paid mutator transaction binding the contract method 0x09a263c1.
//
// Solidity: function testDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue")
}

// TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue is a paid mutator transaction binding the contract method 0x09a263c1.
//
// Solidity: function testDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue is a paid mutator transaction binding the contract method 0x09a263c1.
//
// Solidity: function testDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayload is a paid mutator transaction binding the contract method 0x6c33bacb.
//
// Solidity: function testDepositEthWithAmountToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssWithPayload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssWithPayload")
}

// TestDepositEthWithAmountToTssWithPayload is a paid mutator transaction binding the contract method 0x6c33bacb.
//
// Solidity: function testDepositEthWithAmountToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayload is a paid mutator transaction binding the contract method 0x6c33bacb.
//
// Solidity: function testDepositEthWithAmountToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x4a780339.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0")
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x4a780339.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x4a780339.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue is a paid mutator transaction binding the contract method 0x9073eafe.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue")
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue is a paid mutator transaction binding the contract method 0x9073eafe.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue is a paid mutator transaction binding the contract method 0x9073eafe.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue is a paid mutator transaction binding the contract method 0x85f5c51c.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue")
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue is a paid mutator transaction binding the contract method 0x85f5c51c.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue is a paid mutator transaction binding the contract method 0x85f5c51c.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xa00a6fff.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded")
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xa00a6fff.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xa00a6fff.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9a34d8f3.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress")
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9a34d8f3.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9a34d8f3.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositZetaToConnector is a paid mutator transaction binding the contract method 0xe306a978.
//
// Solidity: function testDepositZetaToConnector() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositZetaToConnector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositZetaToConnector")
}

// TestDepositZetaToConnector is a paid mutator transaction binding the contract method 0xe306a978.
//
// Solidity: function testDepositZetaToConnector() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositZetaToConnector() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositZetaToConnector(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositZetaToConnector is a paid mutator transaction binding the contract method 0xe306a978.
//
// Solidity: function testDepositZetaToConnector() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositZetaToConnector() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositZetaToConnector(&_GatewayEVMInboundTest.TransactOpts)
}

// TestFeeSystemWithUpdatedFee is a paid mutator transaction binding the contract method 0xf1c6b4d3.
//
// Solidity: function testFeeSystemWithUpdatedFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestFeeSystemWithUpdatedFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testFeeSystemWithUpdatedFee")
}

// TestFeeSystemWithUpdatedFee is a paid mutator transaction binding the contract method 0xf1c6b4d3.
//
// Solidity: function testFeeSystemWithUpdatedFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestFeeSystemWithUpdatedFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestFeeSystemWithUpdatedFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestFeeSystemWithUpdatedFee is a paid mutator transaction binding the contract method 0xf1c6b4d3.
//
// Solidity: function testFeeSystemWithUpdatedFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestFeeSystemWithUpdatedFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestFeeSystemWithUpdatedFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestMixedActionTypesInSameTransaction is a paid mutator transaction binding the contract method 0x63d7a418.
//
// Solidity: function testMixedActionTypesInSameTransaction() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestMixedActionTypesInSameTransaction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testMixedActionTypesInSameTransaction")
}

// TestMixedActionTypesInSameTransaction is a paid mutator transaction binding the contract method 0x63d7a418.
//
// Solidity: function testMixedActionTypesInSameTransaction() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestMixedActionTypesInSameTransaction() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestMixedActionTypesInSameTransaction(&_GatewayEVMInboundTest.TransactOpts)
}

// TestMixedActionTypesInSameTransaction is a paid mutator transaction binding the contract method 0x63d7a418.
//
// Solidity: function testMixedActionTypesInSameTransaction() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestMixedActionTypesInSameTransaction() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestMixedActionTypesInSameTransaction(&_GatewayEVMInboundTest.TransactOpts)
}

// TestNispe is a paid mutator transaction binding the contract method 0x2aa5ca5d.
//
// Solidity: function testNispe() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestNispe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testNispe")
}

// TestNispe is a paid mutator transaction binding the contract method 0x2aa5ca5d.
//
// Solidity: function testNispe() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestNispe() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestNispe(&_GatewayEVMInboundTest.TransactOpts)
}

// TestNispe is a paid mutator transaction binding the contract method 0x2aa5ca5d.
//
// Solidity: function testNispe() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestNispe() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestNispe(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x0fc37f5e.
//
// Solidity: function testRevertDepositEthToTssIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositEthToTssIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositEthToTssIfPayloadSizeExceeded")
}

// TestRevertDepositEthToTssIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x0fc37f5e.
//
// Solidity: function testRevertDepositEthToTssIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositEthToTssIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x0fc37f5e.
//
// Solidity: function testRevertDepositEthToTssIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositEthToTssIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestUpdateAdditionalActionFee is a paid mutator transaction binding the contract method 0x586da267.
//
// Solidity: function testUpdateAdditionalActionFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestUpdateAdditionalActionFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testUpdateAdditionalActionFee")
}

// TestUpdateAdditionalActionFee is a paid mutator transaction binding the contract method 0x586da267.
//
// Solidity: function testUpdateAdditionalActionFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestUpdateAdditionalActionFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestUpdateAdditionalActionFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestUpdateAdditionalActionFee is a paid mutator transaction binding the contract method 0x586da267.
//
// Solidity: function testUpdateAdditionalActionFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestUpdateAdditionalActionFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestUpdateAdditionalActionFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestUpdateAdditionalActionFeeOnlyAdmin is a paid mutator transaction binding the contract method 0xa0d60b3a.
//
// Solidity: function testUpdateAdditionalActionFeeOnlyAdmin() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestUpdateAdditionalActionFeeOnlyAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testUpdateAdditionalActionFeeOnlyAdmin")
}

// TestUpdateAdditionalActionFeeOnlyAdmin is a paid mutator transaction binding the contract method 0xa0d60b3a.
//
// Solidity: function testUpdateAdditionalActionFeeOnlyAdmin() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestUpdateAdditionalActionFeeOnlyAdmin() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestUpdateAdditionalActionFeeOnlyAdmin(&_GatewayEVMInboundTest.TransactOpts)
}

// TestUpdateAdditionalActionFeeOnlyAdmin is a paid mutator transaction binding the contract method 0xa0d60b3a.
//
// Solidity: function testUpdateAdditionalActionFeeOnlyAdmin() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestUpdateAdditionalActionFeeOnlyAdmin() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestUpdateAdditionalActionFeeOnlyAdmin(&_GatewayEVMInboundTest.TransactOpts)
}

// GatewayEVMInboundTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestCalledIterator struct {
	Event *GatewayEVMInboundTestCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestCalled represents a Called event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMInboundTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestCalledIterator{contract: _GatewayEVMInboundTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestCalled)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Called", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseCalled(log types.Log) (*GatewayEVMInboundTestCalled, error) {
	event := new(GatewayEVMInboundTestCalled)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestDepositedIterator struct {
	Event *GatewayEVMInboundTestDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestDeposited represents a Deposited event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestDeposited struct {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterDeposited(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMInboundTestDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestDepositedIterator{contract: _GatewayEVMInboundTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestDeposited, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestDeposited)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseDeposited(log types.Log) (*GatewayEVMInboundTestDeposited, error) {
	event := new(GatewayEVMInboundTestDeposited)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestDepositedAndCalledIterator is returned from FilterDepositedAndCalled and is used to iterate over the raw logs and unpacked data for DepositedAndCalled events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestDepositedAndCalledIterator struct {
	Event *GatewayEVMInboundTestDepositedAndCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestDepositedAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestDepositedAndCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestDepositedAndCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestDepositedAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestDepositedAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestDepositedAndCalled represents a DepositedAndCalled event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestDepositedAndCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Amount        *big.Int
	Asset         common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositedAndCalled is a free log retrieval operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterDepositedAndCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMInboundTestDepositedAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestDepositedAndCalledIterator{contract: _GatewayEVMInboundTest.contract, event: "DepositedAndCalled", logs: logs, sub: sub}, nil
}

// WatchDepositedAndCalled is a free log subscription operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchDepositedAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestDepositedAndCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestDepositedAndCalled)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositedAndCalled is a log parse operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseDepositedAndCalled(log types.Log) (*GatewayEVMInboundTestDepositedAndCalled, error) {
	event := new(GatewayEVMInboundTestDepositedAndCalled)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestExecutedIterator struct {
	Event *GatewayEVMInboundTestExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestExecuted represents a Executed event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMInboundTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestExecutedIterator{contract: _GatewayEVMInboundTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestExecuted)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseExecuted(log types.Log) (*GatewayEVMInboundTestExecuted, error) {
	event := new(GatewayEVMInboundTestExecuted)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMInboundTestExecutedWithERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestExecutedWithERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestExecutedWithERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMInboundTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestExecutedWithERC20Iterator{contract: _GatewayEVMInboundTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestExecutedWithERC20)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMInboundTestExecutedWithERC20, error) {
	event := new(GatewayEVMInboundTestExecutedWithERC20)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedERC20Iterator struct {
	Event *GatewayEVMInboundTestReceivedERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestReceivedERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedERC20 represents a ReceivedERC20 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedERC20Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedERC20Iterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedERC20)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedERC20(log types.Log) (*GatewayEVMInboundTestReceivedERC20, error) {
	event := new(GatewayEVMInboundTestReceivedERC20)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedNoParamsIterator struct {
	Event *GatewayEVMInboundTestReceivedNoParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedNoParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestReceivedNoParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedNoParams represents a ReceivedNoParams event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedNoParamsIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedNoParamsIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedNoParams)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedNoParams(log types.Log) (*GatewayEVMInboundTestReceivedNoParams, error) {
	event := new(GatewayEVMInboundTestReceivedNoParams)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedNonPayableIterator struct {
	Event *GatewayEVMInboundTestReceivedNonPayable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedNonPayable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestReceivedNonPayable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedNonPayable represents a ReceivedNonPayable event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedNonPayableIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedNonPayableIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedNonPayable)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedNonPayable(log types.Log) (*GatewayEVMInboundTestReceivedNonPayable, error) {
	event := new(GatewayEVMInboundTestReceivedNonPayable)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedOnCallIterator is returned from FilterReceivedOnCall and is used to iterate over the raw logs and unpacked data for ReceivedOnCall events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedOnCallIterator struct {
	Event *GatewayEVMInboundTestReceivedOnCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestReceivedOnCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedOnCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestReceivedOnCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestReceivedOnCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedOnCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedOnCall represents a ReceivedOnCall event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedOnCall struct {
	Sender  common.Address
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCall is a free log retrieval operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedOnCall(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedOnCallIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedOnCallIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedOnCall", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCall is a free log subscription operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedOnCall(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedOnCall) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedOnCall)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceivedOnCall is a log parse operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedOnCall(log types.Log) (*GatewayEVMInboundTestReceivedOnCall, error) {
	event := new(GatewayEVMInboundTestReceivedOnCall)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedPayableIterator struct {
	Event *GatewayEVMInboundTestReceivedPayable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedPayable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestReceivedPayable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedPayable represents a ReceivedPayable event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedPayable struct {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedPayableIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedPayableIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedPayable)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedPayable(log types.Log) (*GatewayEVMInboundTestReceivedPayable, error) {
	event := new(GatewayEVMInboundTestReceivedPayable)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedRevertIterator struct {
	Event *GatewayEVMInboundTestReceivedRevert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedRevert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestReceivedRevert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedRevert represents a ReceivedRevert event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedRevertIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedRevertIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedRevert)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedRevert(log types.Log) (*GatewayEVMInboundTestReceivedRevert, error) {
	event := new(GatewayEVMInboundTestReceivedRevert)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestRevertedIterator struct {
	Event *GatewayEVMInboundTestReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReverted represents a Reverted event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReverted struct {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMInboundTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestRevertedIterator{contract: _GatewayEVMInboundTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReverted)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Reverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReverted(log types.Log) (*GatewayEVMInboundTestReverted, error) {
	event := new(GatewayEVMInboundTestReverted)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator is returned from FilterUpdatedAdditionalActionFee and is used to iterate over the raw logs and unpacked data for UpdatedAdditionalActionFee events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator struct {
	Event *GatewayEVMInboundTestUpdatedAdditionalActionFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestUpdatedAdditionalActionFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestUpdatedAdditionalActionFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestUpdatedAdditionalActionFee represents a UpdatedAdditionalActionFee event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedAdditionalActionFee struct {
	OldFeeWei *big.Int
	NewFeeWei *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAdditionalActionFee is a free log retrieval operation binding the contract event 0xe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8.
//
// Solidity: event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterUpdatedAdditionalActionFee(opts *bind.FilterOpts) (*GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "UpdatedAdditionalActionFee")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator{contract: _GatewayEVMInboundTest.contract, event: "UpdatedAdditionalActionFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedAdditionalActionFee is a free log subscription operation binding the contract event 0xe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8.
//
// Solidity: event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchUpdatedAdditionalActionFee(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestUpdatedAdditionalActionFee) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "UpdatedAdditionalActionFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestUpdatedAdditionalActionFee)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedAdditionalActionFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedAdditionalActionFee is a log parse operation binding the contract event 0xe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8.
//
// Solidity: event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseUpdatedAdditionalActionFee(log types.Log) (*GatewayEVMInboundTestUpdatedAdditionalActionFee, error) {
	event := new(GatewayEVMInboundTestUpdatedAdditionalActionFee)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedAdditionalActionFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator struct {
	Event *GatewayEVMInboundTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestUpdatedGatewayTSSAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestUpdatedGatewayTSSAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedGatewayTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator{contract: _GatewayEVMInboundTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestUpdatedGatewayTSSAddress)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*GatewayEVMInboundTestUpdatedGatewayTSSAddress, error) {
	event := new(GatewayEVMInboundTestUpdatedGatewayTSSAddress)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogIterator struct {
	Event *GatewayEVMInboundTestLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLog represents a Log event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogIterator{contract: _GatewayEVMInboundTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLog)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLog(log types.Log) (*GatewayEVMInboundTestLog, error) {
	event := new(GatewayEVMInboundTestLog)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogAddressIterator struct {
	Event *GatewayEVMInboundTestLogAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogAddress represents a LogAddress event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogAddressIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogAddressIterator{contract: _GatewayEVMInboundTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogAddress)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogAddress(log types.Log) (*GatewayEVMInboundTestLogAddress, error) {
	event := new(GatewayEVMInboundTestLogAddress)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArrayIterator struct {
	Event *GatewayEVMInboundTestLogArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogArray represents a LogArray event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogArrayIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogArrayIterator{contract: _GatewayEVMInboundTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogArray)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogArray(log types.Log) (*GatewayEVMInboundTestLogArray, error) {
	event := new(GatewayEVMInboundTestLogArray)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray0Iterator struct {
	Event *GatewayEVMInboundTestLogArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogArray0 represents a LogArray0 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogArray0Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogArray0)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogArray0(log types.Log) (*GatewayEVMInboundTestLogArray0, error) {
	event := new(GatewayEVMInboundTestLogArray0)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray1Iterator struct {
	Event *GatewayEVMInboundTestLogArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogArray1 represents a LogArray1 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogArray1Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogArray1)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogArray1(log types.Log) (*GatewayEVMInboundTestLogArray1, error) {
	event := new(GatewayEVMInboundTestLogArray1)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogBytesIterator struct {
	Event *GatewayEVMInboundTestLogBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogBytes represents a LogBytes event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogBytesIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogBytesIterator{contract: _GatewayEVMInboundTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogBytes)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogBytes(log types.Log) (*GatewayEVMInboundTestLogBytes, error) {
	event := new(GatewayEVMInboundTestLogBytes)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogBytes32Iterator struct {
	Event *GatewayEVMInboundTestLogBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogBytes32 represents a LogBytes32 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogBytes32Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogBytes32)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogBytes32(log types.Log) (*GatewayEVMInboundTestLogBytes32, error) {
	event := new(GatewayEVMInboundTestLogBytes32)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogIntIterator struct {
	Event *GatewayEVMInboundTestLogInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogInt represents a LogInt event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogIntIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogIntIterator{contract: _GatewayEVMInboundTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogInt)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogInt(log types.Log) (*GatewayEVMInboundTestLogInt, error) {
	event := new(GatewayEVMInboundTestLogInt)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedAddressIterator struct {
	Event *GatewayEVMInboundTestLogNamedAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogNamedAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedAddressIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedAddress)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayEVMInboundTestLogNamedAddress, error) {
	event := new(GatewayEVMInboundTestLogNamedAddress)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArrayIterator struct {
	Event *GatewayEVMInboundTestLogNamedArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogNamedArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedArray represents a LogNamedArray event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedArrayIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedArray)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayEVMInboundTestLogNamedArray, error) {
	event := new(GatewayEVMInboundTestLogNamedArray)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray0Iterator struct {
	Event *GatewayEVMInboundTestLogNamedArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogNamedArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedArray0Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedArray0)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayEVMInboundTestLogNamedArray0, error) {
	event := new(GatewayEVMInboundTestLogNamedArray0)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray1Iterator struct {
	Event *GatewayEVMInboundTestLogNamedArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogNamedArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedArray1Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedArray1)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayEVMInboundTestLogNamedArray1, error) {
	event := new(GatewayEVMInboundTestLogNamedArray1)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedBytesIterator struct {
	Event *GatewayEVMInboundTestLogNamedBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogNamedBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedBytesIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedBytes)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayEVMInboundTestLogNamedBytes, error) {
	event := new(GatewayEVMInboundTestLogNamedBytes)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedBytes32Iterator struct {
	Event *GatewayEVMInboundTestLogNamedBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogNamedBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedBytes32Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedBytes32)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayEVMInboundTestLogNamedBytes32, error) {
	event := new(GatewayEVMInboundTestLogNamedBytes32)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedDecimalIntIterator struct {
	Event *GatewayEVMInboundTestLogNamedDecimalInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedDecimalInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogNamedDecimalInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedDecimalIntIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedDecimalInt)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayEVMInboundTestLogNamedDecimalInt, error) {
	event := new(GatewayEVMInboundTestLogNamedDecimalInt)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedDecimalUintIterator struct {
	Event *GatewayEVMInboundTestLogNamedDecimalUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedDecimalUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogNamedDecimalUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedDecimalUintIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedDecimalUint)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayEVMInboundTestLogNamedDecimalUint, error) {
	event := new(GatewayEVMInboundTestLogNamedDecimalUint)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedIntIterator struct {
	Event *GatewayEVMInboundTestLogNamedInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogNamedInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedInt represents a LogNamedInt event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedIntIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedInt)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayEVMInboundTestLogNamedInt, error) {
	event := new(GatewayEVMInboundTestLogNamedInt)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedStringIterator struct {
	Event *GatewayEVMInboundTestLogNamedString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogNamedString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedString represents a LogNamedString event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedStringIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedString)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedString(log types.Log) (*GatewayEVMInboundTestLogNamedString, error) {
	event := new(GatewayEVMInboundTestLogNamedString)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedUintIterator struct {
	Event *GatewayEVMInboundTestLogNamedUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogNamedUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedUint represents a LogNamedUint event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedUintIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedUint)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayEVMInboundTestLogNamedUint, error) {
	event := new(GatewayEVMInboundTestLogNamedUint)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogStringIterator struct {
	Event *GatewayEVMInboundTestLogString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogString represents a LogString event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogStringIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogStringIterator{contract: _GatewayEVMInboundTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogString)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogString(log types.Log) (*GatewayEVMInboundTestLogString, error) {
	event := new(GatewayEVMInboundTestLogString)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogUintIterator struct {
	Event *GatewayEVMInboundTestLogUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogUint represents a LogUint event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogUintIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogUintIterator{contract: _GatewayEVMInboundTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogUint)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogUint(log types.Log) (*GatewayEVMInboundTestLogUint, error) {
	event := new(GatewayEVMInboundTestLogUint)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogsIterator struct {
	Event *GatewayEVMInboundTestLogs // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMInboundTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogs)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMInboundTestLogs)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMInboundTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogs represents a Logs event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogsIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogsIterator{contract: _GatewayEVMInboundTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogs)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "logs", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogs(log types.Log) (*GatewayEVMInboundTestLogs, error) {
	event := new(GatewayEVMInboundTestLogs)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
