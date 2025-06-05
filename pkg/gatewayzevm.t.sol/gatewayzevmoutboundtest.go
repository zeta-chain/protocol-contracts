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

// GatewayZEVMOutboundTestMetaData contains all meta data concerning the GatewayZEVMOutboundTest contract.
var GatewayZEVMOutboundTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForDifferentZRC20WithdrawAndCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForDifferentZRC20Withdrawal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForZRC20WithdrawAndCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForZRC20Withdrawal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnProtocolFeesFailsWithInsufficientAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDeposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfSenderNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETA\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractWhenPaused\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfInsufficientAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfInsufficientProtocolBalance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsWhenPaused\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteAbortUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContractIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContractFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextData\",\"inputs\":[{\"name\":\"origin\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"msgSender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataAbort\",\"inputs\":[{\"name\":\"abortContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAbortContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outgoing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataRevert\",\"inputs\":[{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b50620144c1806200003e6000396000f3fe608060405234801561001057600080fd5b50600436106104f95760003560e01c806375ca5587116102a1578063c35cb5e41161016b578063eab7674e116100e3578063f1d98f1b11610097578063f589655d1161007c578063f589655d146107af578063fa7626d4146107b7578063fb339a1c146107c457600080fd5b8063f1d98f1b1461079f578063f2a91358146107a757600080fd5b8063ec294d9f116100c8578063ec294d9f14610787578063ee0f4ea11461078f578063ef2b53941461079757600080fd5b8063eab7674e14610777578063eb78bd7d1461077f57600080fd5b8063da8fd81d1161013a578063e20c9f711161011f578063e20c9f7114610732578063e36c0a231461073a578063e63ab1e91461074257600080fd5b8063da8fd81d14610722578063e09bc6591461072a57600080fd5b8063c35cb5e414610702578063c8814d2e1461070a578063ca26929c14610712578063cf2c3d1d1461071a57600080fd5b806397f7661f11610219578063b5508aa9116101cd578063b84a3d2f116101b2578063b84a3d2f146106e2578063b936be8c14610668578063ba414fa6146106ea57600080fd5b8063b5508aa9146106d2578063b82ce921146106da57600080fd5b80639c9acd5d116101fe5780639c9acd5d146106ba578063ae4662ec146106c2578063b0464fdc146106ca57600080fd5b806397f7661f14610668578063996b7675146106b257600080fd5b806385226c8111610270578063890a2d6711610255578063890a2d671461068d578063916a17c61461069557806396d9d876146106aa57600080fd5b806385226c8114610670578063884660a31461068557600080fd5b806375ca5587146106505780637cec29b0146106585780637f924c4e14610660578063828d267c1461066857600080fd5b8063339bd828116103e257806351336fb01161035a5780636163f8ef1161030e57806362fe2915116102f357806362fe29151461062b57806366d9a9a0146106335780636efa04b51461064857600080fd5b80636163f8ef1461061b57806361cfddb71461062357600080fd5b806358c9987f1161033f57806358c9987f1461060b5780635b4c90e1146105765780635cec7db51461061357600080fd5b806351336fb0146105fb578063514fcc7f1461060357600080fd5b80633e5e3c23116103b1578063445bbf4711610396578063445bbf47146105e357806348f4fd07146105eb57806350dd38ca146105f357600080fd5b80633e5e3c23146105d35780633f7286f4146105db57600080fd5b8063339bd828146105b357806333f7c81a146105bb5780633626c616146105c35780633ab5b199146105cb57600080fd5b806321d92591116104755780632acb21b4116104445780632ed7cff6116104295780632ed7cff61461059b5780632fb780ba146105a35780633177f381146105ab57600080fd5b80632acb21b41461057e5780632ade38801461058657600080fd5b806321d925911461055e5780632468bc0f14610566578063278206251461056e5780632948df411461057657600080fd5b80631832cb6e116104cc5780631c785a14116104b15780631c785a14146105305780631ed7831c146105385780631ff5bdd71461055657600080fd5b80631832cb6e14610520578063195ac1431461052857600080fd5b8063084fafab146104fe5780630a9254e4146105085780630c4a14b41461051057806314b7a6da14610518575b600080fd5b6105066107cc565b005b6105066109a2565b6105066116d5565b610506611acb565b610506611c02565b610506611da3565b610506611eec565b610540612808565b60405161054d919061e07f565b60405180910390f35b61050661286a565b610506612997565b610506612b14565b610506612fc7565b61050661308f565b610506613257565b61058e613416565b60405161054d919061e11b565b610506613558565b610506613684565b6105066138c7565b610506613fb3565b6105066140e6565b61050661422e565b610506614385565b61054061451d565b61054061457d565b6105066145dd565b6105066148ef565b610506614a21565b610506614b6a565b610506614cb5565b610506614dfb565b610506614fb0565b610506615195565b6105066152e3565b6105066157de565b61063b615cad565b60405161054d919061e281565b610506615e2f565b6105066161d3565b610506616a11565b610506616bfc565b610506616c35565b610678616dc8565b60405161054d919061e31f565b610506616e98565b610506617177565b61069d617332565b60405161054d919061e396565b61050661742d565b610506617580565b6105066176d4565b61050661780b565b61069d6178d6565b6106786179d1565b610506617aa1565b610506617bea565b6106f261833b565b604051901515815260200161054d565b61050661840f565b6105066185ca565b6105066186f3565b610506618976565b610506618b34565b610506618c7b565b610540618e34565b610506618e94565b6107697f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60405190815260200161054d565b610506618f5f565b6105066190ad565b61050661926b565b61050661941a565b610506619543565b6105066196db565b610506619829565b610506619ba0565b601f546106f29060ff1681565b610506619c05565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561085b57600080fd5b505af115801561086f573d6000803e3d6000fd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602c6040516108a3919061e55e565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b15801561090557600080fd5b505af1158015610919573d6000803e3d6000fd5b50506020546024546040517f184b07930000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063184b0793935061096e9290911690602c9060040161e571565b600060405180830381600087803b15801561098857600080fd5b505af115801561099c573d6000803e3d6000fd5b50505050565b602580547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602680549091166112341790556040516109e89061df8e565b604051809103906000f080158015610a04573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600f81527f476174657761795a45564d2e736f6c000000000000000000000000000000000060208201526025549151602481019390935292166044820152610ae8919060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f485cc95500000000000000000000000000000000000000000000000000000000179052619d9f565b601f80546001600160a01b039283166101009081027fffffffffffffffffffffff0000000000000000000000000000000000000000ff9092169190911791829055602080547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092049092169182178155604080516004815260248101825280830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcc5ad8b60000000000000000000000000000000000000000000000000000000017905281518084019490945281518085039093018352838201918290527fb96213e400000000000000000000000000000000000000000000000000000000909152737c591652f159496b14e15616f0948a6d63b585e892737109709ecfa91a80626ff3989d68f67f5b1dd12d9263b96213e492610c2f9286929060440161e593565b600060405180830381600087803b158015610c4957600080fd5b505af1158015610c5d573d6000803e3d6000fd5b505060208054604080517f2722feee00000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169450632722feee935060048082019392918290030181865afa158015610cc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce5919061e5e3565b602780547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055604051610d299061df9c565b604051809103906000f080158015610d45573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161781556027546040517f06447d5600000000000000000000000000000000000000000000000000000000815292166004830152737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d569101600060405180830381600087803b158015610de157600080fd5b505af1158015610df5573d6000803e3d6000fd5b505050506000806000604051610e0a9061dfaa565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610e46573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602054604051601293600193849360009391921690610e9c9061dfb8565b610eab9695949392919061e639565b604051809103906000f080158015610ec7573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179091556023546040517fee2815ba0000000000000000000000000000000000000000000000000000000081526001600482015260248101929092529091169063ee2815ba90604401600060405180830381600087803b158015610f5e57600080fd5b505af1158015610f72573d6000803e3d6000fd5b50506023546040517fa7cb050700000000000000000000000000000000000000000000000000000000815260016004820181905260248201526001600160a01b03909116925063a7cb05079150604401600060405180830381600087803b158015610fdc57600080fd5b505af1158015610ff0573d6000803e3d6000fd5b50506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152633b9aca006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b15801561107057600080fd5b505af1158015611084573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b1580156110d957600080fd5b505af11580156110ed573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af1158015611161573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611185919061e6f3565b506021546025546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a060248201529116906347e7ef24906044016020604051808303816000875af11580156111f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061121a919061e6f3565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561127957600080fd5b505af115801561128d573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561130357600080fd5b505af1158015611317573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a060248201529116925063095ea7b391506044016020604051808303816000875af115801561138b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113af919061e6f3565b50602260009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b15801561140157600080fd5b505af1158015611415573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af1158015611489573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ad919061e6f3565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561150c57600080fd5b505af1158015611520573d6000803e3d6000fd5b5050604080516080810182526025546001600160a01b0390811682526000602080840182815260018587019081528651928301909652918152606084018190528351602c80549185167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161781559251602d8054919095169116179092559251602e55909350909150602f906115b8908261e78b565b50506040805160c0810190915260255460601b6bffffffffffffffffffffffff191660e082015290508060f4810160408051601f198184030181529181529082526000602083810182905260018484018190526060850183905260808501528251908101909252815260a09091015280516030908190611638908261e78b565b5060208201516001820180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039092169190911790556040820151600282015560608201516003820180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790556080820151600482015560a0820151600582019061099c908261e78b565b6022546027546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015611726573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061174a919061e84a565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa15801561179c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117c0919061e84a565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b031660848201819052919250903190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561185757600080fd5b505af115801561186b573d6000803e3d6000fd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602c60405161189f919061e55e565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561190057600080fd5b505af1158015611914573d6000803e3d6000fd5b50506020546024546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506345122cfd93506119699289921690602c9060040161e863565b600060405180830381600087803b15801561198357600080fd5b505af1158015611997573d6000803e3d6000fd5b50506022546027546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156119ea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a0e919061e84a565b9050611a23611a1d868661e8ba565b82619dbe565b6022546020546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611a74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a98919061e84a565b9050611aa48482619dbe565b611ac3611ab1878561e8cd565b6024546001600160a01b031631619dbe565b505050505050565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611b2457600080fd5b505af1158015611b38573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015611b8f57600080fd5b505af1158015611ba3573d6000803e3d6000fd5b50506020546021546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260016024820152600060448201529116925063f45346dc915060640161096e565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611c9157600080fd5b505af1158015611ca5573d6000803e3d6000fd5b505050507f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db76030604051611cd9919061e93e565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611d3a57600080fd5b505af1158015611d4e573d6000803e3d6000fd5b50506020546024546040517f2095dedb0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450632095dedb935061096e929091169060309060040161e951565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611dfc57600080fd5b505af1158015611e10573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015611e8057600080fd5b505af1158015611e94573d6000803e3d6000fd5b50506020546026546040517f6e553f65000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b03918216602482015291169250636e553f65915060440161096e565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611f4557600080fd5b505af1158015611f59573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250612044919060040161e973565b600060405180830381600087803b15801561205e57600080fd5b505af1158015612072573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156120c657600080fd5b505af11580156120da573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561213757600080fd5b505af115801561214b573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250612236919060040161e973565b600060405180830381600087803b15801561225057600080fd5b505af1158015612264573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156122b857600080fd5b505af11580156122cc573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561232957600080fd5b505af115801561233d573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561239157600080fd5b505af11580156123a5573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd93c066500000000000000000000000000000000000000000000000000000000600482015260019250737109709ecfa91a80626ff3989d68f67f5b1dd12d915063c31eb0e090602401600060405180830381600087803b15801561241857600080fd5b505af115801561242c573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561248957600080fd5b505af115801561249d573d6000803e3d6000fd5b50506020546021546026546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810187905290821660448201529116925063f45346dc9150606401600060405180830381600087803b15801561251557600080fd5b505af1158015612529573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561258657600080fd5b505af115801561259a573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156125ee57600080fd5b505af1158015612602573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a08231906024015b602060405180830381865afa158015612656573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061267a919061e84a565b9050612687600082619dbe565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156126e057600080fd5b505af11580156126f4573d6000803e3d6000fd5b50506020546021546026546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810188905290821660448201529116925063f45346dc9150606401600060405180830381600087803b15801561276c57600080fd5b505af1158015612780573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156127d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127f7919061e84a565b90506128038382619dbe565b505050565b6060601680548060200260200160405190810160405280929190818152602001828054801561286057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612842575b5050505050905090565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156128c357600080fd5b505af11580156128d7573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561292e57600080fd5b505af1158015612942573d6000803e3d6000fd5b50506020546040517f6e553f6500000000000000000000000000000000000000000000000000000000815260016004820152600060248201526001600160a01b039091169250636e553f65915060440161096e565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612a0657600080fd5b505af1158015612a1a573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015612a7757600080fd5b505af1158015612a8b573d6000803e3d6000fd5b50506020546027546040517f6e553f65000000000000000000000000000000000000000000000000000000008152600481018690526001600160a01b03918216602482015291169250636e553f6591506044015b600060405180830381600087803b158015612af957600080fd5b505af1158015612b0d573d6000803e3d6000fd5b5050505050565b6022546027546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015612b65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b89919061e84a565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa158015612bdb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bff919061e84a565b6024546040519192506001600160a01b03163190600090612c229060200161e986565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff191660808301529150600090806094810160408051808303601f190181529181529082526027546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015612d1057600080fd5b505af1158015612d24573d6000803e3d6000fd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e9450612d7f93506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602754602054612daf936001600160a01b03928316928c92169061e9c3565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612e1057600080fd5b505af1158015612e24573d6000803e3d6000fd5b50506020546024546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321501a959350612e7b9286928c92911690889060040161ea74565b600060405180830381600087803b158015612e9557600080fd5b505af1158015612ea9573d6000803e3d6000fd5b50506022546027546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015612efc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f20919061e84a565b9050612f2f611a1d888861e8ba565b6022546020546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015612f80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fa4919061e84a565b9050612fb08682619dbe565b612fbd611ab1898761e8cd565b5050505050505050565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561303357600080fd5b505af1158015613047573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024016108eb565b60006040516020016130a09061e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b15801561314e57600080fd5b505af1158015613162573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156131b957600080fd5b505af11580156131cd573d6000803e3d6000fd5b50506020546024546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063bcf7f32b935061322992869260009260019290911690899060040161eaae565b600060405180830381600087803b15801561324357600080fd5b505af1158015611ac3573d6000803e3d6000fd5b60006040516020016132689061e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b15801561333257600080fd5b505af1158015613346573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015b600060405180830381600087803b1580156133a457600080fd5b505af11580156133b8573d6000803e3d6000fd5b50506020546021546024546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063bcf7f32b9450613229938793811692600192911690899060040161eaae565b6060601e805480602002602001604051908101604052809291908181526020016000905b8282101561354f57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156135385783829060005260206000200180546134ab9061e42d565b80601f01602080910402602001604051908101604052809291908181526020018280546134d79061e42d565b80156135245780601f106134f957610100808354040283529160200191613524565b820191906000526020600020905b81548152906001019060200180831161350757829003601f168201915b50505050508152602001906001019061348c565b50505050815250508152602001906001019061343a565b50505050905090565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156135b157600080fd5b505af11580156135c5573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561361c57600080fd5b505af1158015613630573d6000803e3d6000fd5b50506020546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506345122cfd915061096e90600190600090602c9060040161e863565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156136e057600080fd5b505af11580156136f4573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015613766573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061378a919061e6f3565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156137e957600080fd5b505af11580156137fd573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015b600060405180830381600087803b15801561385b57600080fd5b505af115801561386f573d6000803e3d6000fd5b50506020546026546040517f6e553f65000000000000000000000000000000000000000000000000000000008152600481018690526001600160a01b03918216602482015291169250636e553f659150604401612adf565b6026546040516001600160a01b03909116602482015260019060009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae7600000000000000000000000000000000000000000000000000000000017905260275490517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156139b257600080fd5b505af11580156139c6573d6000803e3d6000fd5b50506021546040517ff687d12a000000000000000000000000000000000000000000000000000000008152620186a060048201526001600160a01b03909116925063f687d12a9150602401600060405180830381600087803b158015613a2b57600080fd5b505af1158015613a3f573d6000803e3d6000fd5b50506021546040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152620186a06004820152600093506001600160a01b03909116915063fc5fecd5906024016040805180830381865afa158015613aa9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613acd919061eb03565b6021546025549193506001600160a01b0390811692506347e7ef249116613af4848761e8cd565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015613b57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b7b919061e6f3565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015613bda57600080fd5b505af1158015613bee573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015613c6457600080fd5b505af1158015613c78573d6000803e3d6000fd5b50506021546020546001600160a01b03918216935063095ea7b3925016613c9f848761e8cd565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015613d02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d26919061e6f3565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015613d8557600080fd5b505af1158015613d99573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613df2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e16919061e84a565b602080546026546040519394506001600160a01b0391821693637b15118b93613e5c9392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f198184030181526021548383018352620186a084526001602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b168152613ecb9391928a926001600160a01b03909116918a919060289060040161eb8c565b600060405180830381600087803b158015613ee557600080fd5b505af1158015613ef9573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613f52573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f76919061e84a565b90506000613f84848761e8cd565b9050611ac3613f93828561e8ba565b836040518060600160405280603a815260200162014452603a9139619e3e565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561400c57600080fd5b505af1158015614020573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561407757600080fd5b505af115801561408b573d6000803e3d6000fd5b50506020546021546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450639d4ba465935061096e9290911690600190600090602c9060040161ebfa565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561414257600080fd5b505af1158015614156573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156141aa57600080fd5b505af11580156141be573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd93c0665000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156137e957600080fd5b60275460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561428a57600080fd5b505af115801561429e573d6000803e3d6000fd5b5050604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561430e57600080fd5b505af1158015614322573d6000803e3d6000fd5b50506020546021546027546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810187905290821660448201529116925063f45346dc9150606401612adf565b60006040516020016143969061e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b15801561444457600080fd5b505af1158015614458573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156144af57600080fd5b505af11580156144c3573d6000803e3d6000fd5b50506020546021546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063bcf7f32b93506132299286921690600190600090899060040161eaae565b60606018805480602002602001604051908101604052809291908181526020018280548015612860576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311612842575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015612860576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311612842575050505050905090565b6022546027546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa15801561462e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614652919061e84a565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa1580156146a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906146c8919061e84a565b60265460275460405163ca669fa760e01b81526001600160a01b039182166004820152929350163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561472a57600080fd5b505af115801561473e573d6000803e3d6000fd5b50506020546026546040517f6e553f65000000000000000000000000000000000000000000000000000000008152600481018990526001600160a01b03918216602482015291169250636e553f659150604401600060405180830381600087803b1580156147ab57600080fd5b505af11580156147bf573d6000803e3d6000fd5b50506022546027546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015614812573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614836919061e84a565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa158015614888573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906148ac919061e84a565b6026549091506001600160a01b0316316148cf6148c9888861e8ba565b84619dbe565b6148d98583619dbe565b6148e6611a1d888661e8cd565b50505050505050565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561494857600080fd5b505af115801561495c573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156149b357600080fd5b505af11580156149c7573d6000803e3d6000fd5b50506020546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450639d4ba465935061096e92600092600192911690602c9060040161ebfa565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015614a9057600080fd5b505af1158015614aa4573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015614b0157600080fd5b505af1158015614b15573d6000803e3d6000fd5b50506020546027546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506345122cfd9350612adf9286921690602c9060040161e863565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015614bd657600080fd5b505af1158015614bea573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015614c4757600080fd5b505af1158015614c5b573d6000803e3d6000fd5b50506020546021546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450639d4ba465935061096e92909116906001908590602c9060040161ebfa565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015614d2457600080fd5b505af1158015614d38573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015614d9557600080fd5b505af1158015614da9573d6000803e3d6000fd5b50506020546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506345122cfd9150612adf9084908490602c9060040161e863565b6000604051602001614e0c9061e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015614eba57600080fd5b505af1158015614ece573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015614f3e57600080fd5b505af1158015614f52573d6000803e3d6000fd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450613229938793811692600092911690899060040161eaae565b604051600190600090614fc59060200161e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b15801561508f57600080fd5b505af11580156150a3573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561510057600080fd5b505af1158015615114573d6000803e3d6000fd5b50506020546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506321501a95915061516790849087908590889060040161ea74565b600060405180830381600087803b15801561518157600080fd5b505af11580156148e6573d6000803e3d6000fd5b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561520157600080fd5b505af1158015615215573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561527257600080fd5b505af1158015615286573d6000803e3d6000fd5b50506020546021546027546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061096e93928316926001921690602c9060040161ebfa565b6025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561535857600080fd5b505af115801561536c573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018690529116925063095ea7b391506044016020604051808303816000875af11580156153de573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615402919061e6f3565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561546157600080fd5b505af1158015615475573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156154d257600080fd5b505af11580156154e6573d6000803e3d6000fd5b50506021546040517ff687d12a00000000000000000000000000000000000000000000000000000000815262030d4060048201526001600160a01b03909116925063f687d12a9150602401600060405180830381600087803b15801561554b57600080fd5b505af115801561555f573d6000803e3d6000fd5b50506021546040517ffc5fecd500000000000000000000000000000000000000000000000000000000815262030d406004820152600093508392506001600160a01b039091169063fc5fecd5906024016040805180830381865afa1580156155cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906155ef919061eb03565b60208054604080516001600160a01b03868116602483015290921660448301526064808301859052815180840390910181526084909201815291810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f667011120000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152929450909250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916156c29160040161e973565b600060405180830381600087803b1580156156dc57600080fd5b505af11580156156f0573d6000803e3d6000fd5b5050602080546026546040805160609290921b6bffffffffffffffffffffffff19168285015280518083036014018152602154600460348501818152605886018552605490950180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f19ff1d21000000000000000000000000000000000000000000000000000000001790528351808501855262030d40815260019781019790975292517f06cb89830000000000000000000000000000000000000000000000000000000081526001600160a01b0395861698506306cb8983975061516796929590911693926028910161ec31565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561583a57600080fd5b505af115801561584e573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156158a257600080fd5b505af11580156158b6573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd93c0665000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561592657600080fd5b505af115801561593a573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561599757600080fd5b505af11580156159ab573d6000803e3d6000fd5b50506020546024546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506345122cfd9350615a009286921690602c9060040161e863565b600060405180830381600087803b158015615a1a57600080fd5b505af1158015615a2e573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015615a8b57600080fd5b505af1158015615a9f573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015615af357600080fd5b505af1158015615b07573d6000803e3d6000fd5b5050602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015615b9a57600080fd5b505af1158015615bae573d6000803e3d6000fd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602c604051615be2919061e55e565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b158015615c4457600080fd5b505af1158015615c58573d6000803e3d6000fd5b50506020546024546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506345122cfd9350612adf9286921690602c9060040161e863565b6060601b805480602002602001604051908101604052809291908181526020016000905b8282101561354f5783829060005260206000209060020201604051806040016040529081600082018054615d049061e42d565b80601f0160208091040260200160405190810160405280929190818152602001828054615d309061e42d565b8015615d7d5780601f10615d5257610100808354040283529160200191615d7d565b820191906000526020600020905b815481529060010190602001808311615d6057829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015615e1757602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411615dc45790505b50505050508152505081526020019060010190615cd1565b602154602480546040516370a0823160e01b81526001600160a01b03918216600482015260009391909116916370a082319101602060405180830381865afa158015615e7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615ea3919061e84a565b9050615eb0600082619dbe565b6000604051602001615ec19061e986565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff191660808301529150600090806094810160408051808303601f190181529181529082526027546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015615faf57600080fd5b505af1158015615fc3573d6000803e3d6000fd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e945061601e93506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f198184030181529082905260275460205461604f936001600160a01b0392831692600192169061e9c3565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156160b057600080fd5b505af11580156160c4573d6000803e3d6000fd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450616122938793811692600192911690899060040161eaae565b600060405180830381600087803b15801561613c57600080fd5b505af1158015616150573d6000803e3d6000fd5b5050602154602480546040516370a0823160e01b81526001600160a01b03918216600482015260009550921692506370a082319101602060405180830381865afa1580156161a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906161c6919061e84a565b905061099c600182619dbe565b6027546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561624557600080fd5b505af1158015616259573d6000803e3d6000fd5b50506023546020546040516000945060129350600192600292620186a0926001600160a01b0392831692909116906162909061dfb8565b61629f9695949392919061ec8c565b604051809103906000f0801580156162bb573d6000803e3d6000fd5b506025546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152606460248201529192508216906347e7ef24906044016020604051808303816000875af115801561632a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061634e919061e6f3565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156163ad57600080fd5b505af11580156163c1573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561643757600080fd5b505af115801561644b573d6000803e3d6000fd5b50506020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260646024820152908416925063095ea7b391506044016020604051808303816000875af11580156164bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906164df919061e6f3565b506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482015291169063095ea7b3906044016020604051808303816000875af1158015616550573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616574919061e6f3565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156165d357600080fd5b505af11580156165e7573d6000803e3d6000fd5b50506026546040516001600160a01b039091166024820152600192506000915060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae76000000000000000000000000000000000000000000000000000000000179052517ffc5fecd5000000000000000000000000000000000000000000000000000000008152620186a0600482015290915060009081906001600160a01b0386169063fc5fecd5906024016040805180830381865afa1580156166c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906166eb919061eb03565b60215491935091506167079083906001600160a01b0316619ebe565b602154604080517f18160ddd00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916318160ddd9160048083019260209291908290030181865afa15801561676a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061678e919061e84a565b90506000866001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156167d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906167f4919061e84a565b602080546026546040519394506001600160a01b0391821693637b15118b9361683a9392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f198184030181528282018252620186a083526001602084015290517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152616899928b918d918c9160289060040161eb8c565b600060405180830381600087803b1580156168b357600080fd5b505af11580156168c7573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015616920573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616944919061e84a565b90506000886001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015616986573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906169aa919061e84a565b90506169d96169b9868661e8ba565b836040518060600160405280602b8152602001620143c9602b9139619e3e565b616a066169e6898561e8ba565b826040518060600160405280603881526020016201441a60389139619e3e565b505050505050505050565b6000604051602001616a229061e986565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff191660808301529150600090806094810160408051808303601f190181529181529082526027546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015616b1057600080fd5b505af1158015616b24573d6000803e3d6000fd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e9450616b7f93506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602754602054616bb0936001600160a01b0392831692600192169061e9c3565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa79060240161338a565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401612639565b6000604051602001616c469061e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015616cf457600080fd5b505af1158015616d08573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015616d5f57600080fd5b505af1158015616d73573d6000803e3d6000fd5b50506020546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506321501a959150613229908490600190600090889060040161ea74565b6060601a805480602002602001604051908101604052809291908181526020016000905b8282101561354f578382906000526020600020018054616e0b9061e42d565b80601f0160208091040260200160405190810160405280929190818152602001828054616e379061e42d565b8015616e845780601f10616e5957610100808354040283529160200191616e84565b820191906000526020600020905b815481529060010190602001808311616e6757829003601f168201915b505050505081526020019060010190616dec565b602154602480546040516370a0823160e01b81526001600160a01b03918216600482015260009391909116916370a082319101602060405180830381865afa158015616ee8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616f0c919061e84a565b9050616f19600082619dbe565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015616fa857600080fd5b505af1158015616fbc573d6000803e3d6000fd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602c604051616ff0919061e55e565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561705157600080fd5b505af1158015617065573d6000803e3d6000fd5b50506020546021546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba46594506170c293928316926001921690602c9060040161ebfa565b600060405180830381600087803b1580156170dc57600080fd5b505af11580156170f0573d6000803e3d6000fd5b5050602154602480546040516370a0823160e01b81526001600160a01b03918216600482015260009550921692506370a082319101602060405180830381865afa158015617142573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617166919061e84a565b9050617173600182619dbe565b5050565b60405160019060009061718c9060200161e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b15801561725657600080fd5b505af115801561726a573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156172c757600080fd5b505af11580156172db573d6000803e3d6000fd5b50506020546027546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321501a9593506151679286928992911690889060040161ea74565b6060601d805480602002602001604051908101604052809291908181526020016000905b8282101561354f5760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561741557602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116173c25790505b50505050508152505081526020019060010190617356565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561748957600080fd5b505af115801561749d573d6000803e3d6000fd5b5050604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561750d57600080fd5b505af1158015617521573d6000803e3d6000fd5b50506020546021546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101869052911660448201819052925063f45346dc9150606401612adf565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156175d957600080fd5b505af11580156175ed573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561765d57600080fd5b505af1158015617671573d6000803e3d6000fd5b50506020546021546026546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526000602482015290821660448201529116925063f45346dc915060640161096e565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561772d57600080fd5b505af1158015617741573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561779857600080fd5b505af11580156177ac573d6000803e3d6000fd5b50506020546026546040517ff45346dc00000000000000000000000000000000000000000000000000000000815260006004820152600160248201526001600160a01b0391821660448201529116925063f45346dc915060640161096e565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561787a57600080fd5b505af115801561788e573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401613841565b6060601c805480602002602001604051908101604052809291908181526020016000905b8282101561354f5760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156179b957602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116179665790505b505050505081525050815260200190600101906178fa565b60606019805480602002602001604051908101604052809291908181526020016000905b8282101561354f578382906000526020600020018054617a149061e42d565b80601f0160208091040260200160405190810160405280929190818152602001828054617a409061e42d565b8015617a8d5780601f10617a6257610100808354040283529160200191617a8d565b820191906000526020600020905b815481529060010190602001808311617a7057829003601f168201915b5050505050815260200190600101906179f5565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015617b1057600080fd5b505af1158015617b24573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015617b8157600080fd5b505af1158015617b95573d6000803e3d6000fd5b50506020546040517f6e553f65000000000000000000000000000000000000000000000000000000008152600481018590526001600160a01b03909116602482018190529250636e553f659150604401612adf565b6027546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015617c5c57600080fd5b505af1158015617c70573d6000803e3d6000fd5b50506023546020546040516000945060129350600192600292620186a0926001600160a01b039283169290911690617ca79061dfb8565b617cb69695949392919061ec8c565b604051809103906000f080158015617cd2573d6000803e3d6000fd5b506025546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152606460248201529192508216906347e7ef24906044016020604051808303816000875af1158015617d41573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617d65919061e6f3565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015617dc457600080fd5b505af1158015617dd8573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015617e4e57600080fd5b505af1158015617e62573d6000803e3d6000fd5b50506020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260646024820152908416925063095ea7b391506044016020604051808303816000875af1158015617ed2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617ef6919061e6f3565b506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482015291169063095ea7b3906044016020604051808303816000875af1158015617f67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617f8b919061e6f3565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015617fea57600080fd5b505af1158015617ffe573d6000803e3d6000fd5b50506040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152620186a06004820152600192506000915081906001600160a01b0385169063fc5fecd5906024016040805180830381865afa158015618069573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061808d919061eb03565b60215491935091506180a99083906001600160a01b0316619ebe565b602154604080517f18160ddd00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916318160ddd9160048083019260209291908290030181865afa15801561810c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618130919061e84a565b90506000856001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015618172573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618196919061e84a565b602080546026546040519394506001600160a01b0391821693637c0dcb5f936181dc9392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b604051602081830303815290604052878960286040518563ffffffff1660e01b815260040161820e949392919061ed1c565b600060405180830381600087803b15801561822857600080fd5b505af115801561823c573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015618295573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906182b9919061e84a565b90506000876001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156182fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061831f919061e84a565b905061832e6169b9868661e8ba565b612fbd6169e6888561e8ba565b60085460009060ff1615618353575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa1580156183e4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618408919061e84a565b1415905090565b6040516001906000906184249060200161e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b1580156184ee57600080fd5b505af1158015618502573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561855f57600080fd5b505af1158015618573573d6000803e3d6000fd5b50506020546024546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321501a9593506151679286928992911690889060040161ea74565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561862357600080fd5b505af1158015618637573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561868e57600080fd5b505af11580156186a2573d6000803e3d6000fd5b50506020546040517f184b07930000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063184b0793915061096e90600090602c9060040161e571565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015618744573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618768919061e84a565b9050618775600082619dbe565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156187ce57600080fd5b505af11580156187e2573d6000803e3d6000fd5b5050604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561885257600080fd5b505af1158015618866573d6000803e3d6000fd5b50506020546021546026546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810188905290821660448201529116925063f45346dc9150606401600060405180830381600087803b1580156188de57600080fd5b505af11580156188f2573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015618945573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618969919061e84a565b9050612803600082619dbe565b60006040516020016189879061e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b158015618a5157600080fd5b505af1158015618a65573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015618ac257600080fd5b505af1158015618ad6573d6000803e3d6000fd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450613229938793811692600192911690899060040161eaae565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015618b8d57600080fd5b505af1158015618ba1573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015618c1157600080fd5b505af1158015618c25573d6000803e3d6000fd5b50506020546024546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506345122cfd935061096e926000921690602c9060040161e863565b6000604051602001618c8c9061e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b158015618d5657600080fd5b505af1158015618d6a573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015618dc757600080fd5b505af1158015618ddb573d6000803e3d6000fd5b50506020546021546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca37935061322992869216906001908690899060040161eaae565b60606015805480602002602001604051908101604052809291908181526020018280548015612860576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311612842575050505050905090565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015618f0357600080fd5b505af1158015618f17573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401615c2a565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015618fb857600080fd5b505af1158015618fcc573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561903c57600080fd5b505af1158015619050573d6000803e3d6000fd5b50506020546021546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061096e93928316926000921690602c9060040161ebfa565b60006040516020016190be9061e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b15801561918857600080fd5b505af115801561919c573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156191f957600080fd5b505af115801561920d573d6000803e3d6000fd5b50506020546021546027546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450613229938793811692600192911690899060040161eaae565b600060405160200161927c9061e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b15801561932a57600080fd5b505af115801561933e573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156193ae57600080fd5b505af11580156193c2573d6000803e3d6000fd5b50506020546021546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321501a959350613229928692600092911690889060040161ea74565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561947357600080fd5b505af1158015619487573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156194de57600080fd5b505af11580156194f2573d6000803e3d6000fd5b50506020546040517f2095dedb0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632095dedb915061096e9060009060309060040161e951565b60006040516020016195549061e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b15801561960257600080fd5b505af1158015619616573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561966d57600080fd5b505af1158015619681573d6000803e3d6000fd5b50506020546021546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca3793506132299286921690600190600090899060040161eaae565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561974757600080fd5b505af115801561975b573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156197b857600080fd5b505af11580156197cc573d6000803e3d6000fd5b50506020546021546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061096e93928316926001921690602c9060040161ebfa565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561988557600080fd5b505af1158015619899573d6000803e3d6000fd5b50506021546040517ff687d12a00000000000000000000000000000000000000000000000000000000815261c35060048201526001600160a01b03909116925063f687d12a9150602401600060405180830381600087803b1580156198fd57600080fd5b505af1158015619911573d6000803e3d6000fd5b50506021546040517ffc5fecd500000000000000000000000000000000000000000000000000000000815261c3506004820152600093506001600160a01b03909116915063fc5fecd5906024016040805180830381865afa15801561997a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061999e919061eb03565b9150506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156199f6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190619a1a919061e84a565b602080546026546040519394506001600160a01b0391821693637c0dcb5f93619a609392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352619ab89288916001600160a01b03169060289060040161ed1c565b600060405180830381600087803b158015619ad257600080fd5b505af1158015619ae6573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015619b3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190619b63919061e84a565b90506000619b71848661e8cd565b9050612b0d619b80828561e8ba565b83604051806060016040528060268152602001620143f460269139619e3e565b60006103e890507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156137e957600080fd5b6000604051602001619c169061e986565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015619cc457600080fd5b505af1158015619cd8573d6000803e3d6000fd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015619d2f57600080fd5b505af1158015619d43573d6000803e3d6000fd5b50506020546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca37935061322992869260009260019290911690899060040161eaae565b6000619da961dfc6565b619db4848483619f1f565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015b60006040518083038186803b158015619e2a57600080fd5b505afa158015611ac3573d6000803e3d6000fd5b6040517f88b44c85000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d906388b44c8590619e929086908690869060040161ed56565b60006040518083038186803b158015619eaa57600080fd5b505afa1580156148e6573d6000803e3d6000fd5b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f690604401619e12565b600080619f2c8584619f9a565b9050619f8f6040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f78790000008152508286604051602001619f7a92919061ed75565b60405160208183030381529060405285619fa6565b9150505b9392505050565b6000619f938383619fd4565b60c08101515160009015619fca57619fc384848460c00151619fef565b9050619f93565b619fc3848461a195565b6000619fe0838361a280565b619f9383836020015184619fa6565b600080619ffa61a28c565b9050600061a008868361a35f565b9050600061a01f826060015183602001518561a805565b9050600061a02f8383898961aa17565b9050600061a03c8261b894565b602081015181519192509060030b1561a0af5789826040015160405160200161a06692919061ed97565b60408051601f19818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261a0a69160040161e973565b60405180910390fd5b600061a0f26040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a20000000000000000000000081525083600161ba63565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d9061a14590849060040161e973565b602060405180830381865afa15801561a162573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061a186919061e5e3565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc9259061a1ea90879060040161e973565b600060405180830381865afa15801561a207573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a22f919081019061eed1565b9050600061a25d828560405160200161a24992919061ef06565b60405160208183030381529060405261bc63565b90506001600160a01b038116619db457848460405160200161a06692919061ef35565b6171738282600061bc76565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c9061a31390849060040161efe0565b600060405180830381865afa15801561a330573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a358919081019061f027565b9250505090565b61a3916040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d905061a3dc6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b61a3e58561bd79565b6020820152600061a3f58661c15e565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa15801561a437573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a45f919081019061f027565b8683856020015160405160200161a479949392919061f070565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb119061a4d190859060040161e973565b600060405180830381865afa15801561a4ee573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a516919081019061f027565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f69061a55e90849060040161f174565b602060405180830381865afa15801561a57b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061a59f919061e6f3565b61a5b4578160405160200161a066919061f1c6565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061a5f990849060040161f258565b600060405180830381865afa15801561a616573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a63e919081019061f027565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f69061a68590849060040161f2aa565b602060405180830381865afa15801561a6a2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061a6c6919061e6f3565b1561a75b576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061a71090849060040161f2aa565b600060405180830381865afa15801561a72d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a755919081019061f027565b60408501525b846001600160a01b03166349c4fac882866000015160405160200161a780919061f2fc565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161a7ac92919061f368565b600060405180830381865afa15801561a7c9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a7f1919081019061f027565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b606081526020019060019003908161a8215790505090506040518060400160405280600481526020017f67726570000000000000000000000000000000000000000000000000000000008152508160008151811061a8815761a88161f38d565b60200260200101819052506040518060400160405280600381526020017f2d726c00000000000000000000000000000000000000000000000000000000008152508160018151811061a8d55761a8d561f38d565b60200260200101819052508460405160200161a8f1919061f3bc565b6040516020818303038152906040528160028151811061a9135761a91361f38d565b60200260200101819052508260405160200161a92f919061f428565b6040516020818303038152906040528160038151811061a9515761a95161f38d565b6020026020010181905250600061a9678261b894565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184526000808252908601528251808401909352905182529281019290925291925061a9f8906040805180820182526000808252602091820152815180830190925284518252808501908201529061c3e1565b61aa0d578560405160200161a066919061f469565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d901561aa67565b511590565b61abdb5782602001511561ab23576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a40161a0a6565b8260c001511561abdb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a40161a0a6565b6040805160ff8082526120008201909252600091816020015b606081526020019060019003908161abf457905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061ac4f9061f4fa565b935060ff168151811061ac645761ac6461f38d565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e370000000000000000000000000000000000000081525060405160200161acb5919061f519565b60405160208183030381529060405282828061acd09061f4fa565b935060ff168151811061ace55761ace561f38d565b60200260200101819052506040518060400160405280600681526020017f6465706c6f79000000000000000000000000000000000000000000000000000081525082828061ad329061f4fa565b935060ff168151811061ad475761ad4761f38d565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d6500000000000000000000000000000000000081525082828061ad949061f4fa565b935060ff168151811061ada95761ada961f38d565b6020026020010181905250876020015182828061adc59061f4fa565b935060ff168151811061adda5761adda61f38d565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163745061746800000000000000000000000000000000000081525082828061ae279061f4fa565b935060ff168151811061ae3c5761ae3c61f38d565b60209081029190910101528751828261ae548161f4fa565b935060ff168151811061ae695761ae6961f38d565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e4964000000000000000000000000000000000000000000000081525082828061aeb69061f4fa565b935060ff168151811061aecb5761aecb61f38d565b602002602001018190525061aedf4661c442565b828261aeea8161f4fa565b935060ff168151811061aeff5761aeff61f38d565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c65000000000000000000000000000000000081525082828061af4c9061f4fa565b935060ff168151811061af615761af6161f38d565b60200260200101819052508682828061af799061f4fa565b935060ff168151811061af8e5761af8e61f38d565b602090810291909101015285511561b0b55760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f646500000000000000000000006020820152828261afdf8161f4fa565b935060ff168151811061aff45761aff461f38d565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d9061b04490899060040161e973565b600060405180830381865afa15801561b061573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261b089919081019061f027565b828261b0948161f4fa565b935060ff168151811061b0a95761b0a961f38d565b60200260200101819052505b84602001511561b1855760408051808201909152601281527f2d2d766572696679536f75726365436f646500000000000000000000000000006020820152828261b0fe8161f4fa565b935060ff168151811061b1135761b11361f38d565b60200260200101819052506040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525082828061b1609061f4fa565b935060ff168151811061b1755761b17561f38d565b602002602001018190525061b34c565b61b1bd61aa628660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b61b2505760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261b2008161f4fa565b935060ff168151811061b2155761b21561f38d565b60200260200101819052508460a0015160405160200161b235919061f3bc565b60405160208183030381529060405282828061b1609061f4fa565b8460c0015115801561b29357506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261b29190511590565b155b1561b34c5760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261b2d78161f4fa565b935060ff168151811061b2ec5761b2ec61f38d565b602002602001018190525061b3008861c4e2565b60405160200161b310919061f3bc565b60405160208183030381529060405282828061b32b9061f4fa565b935060ff168151811061b3405761b34061f38d565b60200260200101819052505b6040808601518151808301835260008082526020918201528251808401909352815183529081019082015261b38090511590565b61b4155760408051808201909152600b81527f2d2d72656c6179657249640000000000000000000000000000000000000000006020820152828261b3c38161f4fa565b935060ff168151811061b3d85761b3d861f38d565b6020026020010181905250846040015182828061b3f49061f4fa565b935060ff168151811061b4095761b40961f38d565b60200260200101819052505b60608501511561b5365760408051808201909152600681527f2d2d73616c7400000000000000000000000000000000000000000000000000006020820152828261b45e8161f4fa565b935060ff168151811061b4735761b47361f38d565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa15801561b4e2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261b50a919081019061f027565b828261b5158161f4fa565b935060ff168151811061b52a5761b52a61f38d565b60200260200101819052505b60e0850151511561b5dd5760408051808201909152600a81527f2d2d6761734c696d6974000000000000000000000000000000000000000000006020820152828261b5808161f4fa565b935060ff168151811061b5955761b59561f38d565b602002602001018190525061b5b18560e001516000015161c442565b828261b5bc8161f4fa565b935060ff168151811061b5d15761b5d161f38d565b60200260200101819052505b60e0850151602001511561b6875760408051808201909152600a81527f2d2d6761735072696365000000000000000000000000000000000000000000006020820152828261b62a8161f4fa565b935060ff168151811061b63f5761b63f61f38d565b602002602001018190525061b65b8560e001516020015161c442565b828261b6668161f4fa565b935060ff168151811061b67b5761b67b61f38d565b60200260200101819052505b60e0850151604001511561b7315760408051808201909152600e81527f2d2d6d61784665655065724761730000000000000000000000000000000000006020820152828261b6d48161f4fa565b935060ff168151811061b6e95761b6e961f38d565b602002602001018190525061b7058560e001516040015161c442565b828261b7108161f4fa565b935060ff168151811061b7255761b72561f38d565b60200260200101819052505b60e0850151606001511561b7db5760408051808201909152601681527f2d2d6d61785072696f72697479466565506572476173000000000000000000006020820152828261b77e8161f4fa565b935060ff168151811061b7935761b79361f38d565b602002602001018190525061b7af8560e001516060015161c442565b828261b7ba8161f4fa565b935060ff168151811061b7cf5761b7cf61f38d565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561b7f95761b7f961e715565b60405190808252806020026020018201604052801561b82c57816020015b606081526020019060019003908161b8175790505b50905060005b8260ff168160ff16101561b88557838160ff168151811061b8555761b85561f38d565b6020026020010151828260ff168151811061b8725761b87261f38d565b602090810291909101015260010161b832565b5093505050505b949350505050565b61b8bb6040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c9161b9419186910161f584565b600060405180830381865afa15801561b95e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261b986919081019061f027565b9050600061b994868361cfd1565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b815260040161b9c4919061e31f565b6000604051808303816000875af115801561b9e3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261ba0b919081019061f5cb565b805190915060030b1580159061ba245750602081015151155b801561ba335750604081015151155b1561aa0d578160008151811061ba4b5761ba4b61f38d565b602002602001015160405160200161a066919061f681565b6060600061ba988560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252865182528087019082015290915061bacf9082905b9061d126565b1561bc2c57600061bb4c8261bb468461bb4061bb128a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061d14d565b9061d1af565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061bbb090829061d126565b1561bc1a57604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bc17905b829061d234565b90505b61bc238161d25a565b92505050619f93565b821561bc4557848460405160200161a06692919061f86d565b5050604080516020810190915260008152619f93565b509392505050565b6000808251602084016000f09392505050565b8160a001511561bc8557505050565b600061bc9284848461d2c3565b9050600061bc9f8261b894565b602081015181519192509060030b15801561bd3b5750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bd3b9060408051808201825260008082526020918201528151808301909252845182528085019082015261bac9565b1561bd4857505050505050565b6040820151511561bd6857816040015160405160200161a066919061f914565b8060405160200161a066919061f972565b6060600061bdae8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061be13905b829061c3e1565b1561be8257604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619f939061be7d90839061d85e565b61d25a565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bee4905b829061d8e8565b60010361bfb157604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bf4a9061bc10565b50604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619f939061be7d905b839061d234565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c0109061be0c565b1561c14757604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061c07890839061d982565b90506000816001835161c08b919061e8ba565b8151811061c09b5761c09b61f38d565b6020026020010151905061c13e61be7d61c1116040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925285518252808601908201529061d85e565b95945050505050565b8260405160200161a066919061f9dd565b50919050565b6060600061c1938360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061c1f59061be0c565b1561c20357619f938161d25a565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c2629061bedd565b60010361c2cc57604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619f939061be7d9061bfaa565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c32b9061be0c565b1561c14757604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061c39390839061d982565b905060018151111561c3cf57806002825161c3ae919061e8ba565b8151811061c3be5761c3be61f38d565b602002602001015192505050919050565b508260405160200161a066919061f9dd565b80518251600091111561c3f657506000619db8565b8151835160208501516000929161c40c9161e8cd565b61c416919061e8ba565b90508260200151810361c42d576001915050619db8565b82516020840151819020912014905092915050565b6060600061c44f8361da27565b600101905060008167ffffffffffffffff81111561c46f5761c46f61e715565b6040519080825280601f01601f19166020018201604052801561c499576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461c4a357509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e534544000000000000000000000000000000000000000000008184019081528551808701875283815284019290925284518086019095525184529083015260609161c56e905b829061db09565b1561c5ae57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c60d9061c567565b1561c64d57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c6ac9061c567565b1561c6ec57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c74b9061c567565b8061c7b05750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c7b09061c567565b1561c7f057505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c84f9061c567565b8061c8b45750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c8b49061c567565b1561c8f457505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c9539061c567565b8061c9b85750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c9b89061c567565b1561c9f857505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ca579061c567565b8061cabc5750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261cabc9061c567565b1561cafc57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261cb5b9061c567565b1561cb9b57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261cbfa9061c567565b1561cc3a57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261cc999061c567565b1561ccd957505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261cd389061c567565b1561cd7857505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261cdd79061c567565b1561ce1757505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ce769061c567565b8061cedb5750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261cedb9061c567565b1561cf1b57505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261cf7a9061c567565b1561cfba57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b6040808401518451915161a066929060200161fabb565b60608060005b845181101561d05c578185828151811061cff35761cff361f38d565b602002602001015160405160200161d00c92919061ef06565b60405160208183030381529060405291506001855161d02b919061e8ba565b811461d054578160405160200161d042919061fc24565b60405160208183030381529060405291505b60010161cfd7565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161d075579050509050838160008151811061d0a05761d0a061f38d565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061d0f45761d0f461f38d565b6020026020010181905250818160028151811061d1135761d11361f38d565b6020908102919091010152949350505050565b602080830151835183519284015160009361d144929184919061db1d565b14159392505050565b6040805180820190915260008082526020820152600061d17f846000015185602001518560000151866020015161dc2e565b905083602001518161d191919061e8ba565b8451859061d1a090839061e8ba565b90525060208401525090919050565b604080518082019091526000808252602082015281518351101561d1d4575081619db8565b602080830151908401516001911461d1fb5750815160208481015190840151829020919020145b801561d22c5782518451859061d21290839061e8ba565b905250825160208501805161d22890839061e8cd565b9052505b509192915050565b604080518082019091526000808252602082015261d25383838361dd4e565b5092915050565b60606000826000015167ffffffffffffffff81111561d27b5761d27b61e715565b6040519080825280601f01601f19166020018201604052801561d2a5576020820181803683370190505b509050600060208201905061d253818560200151866000015161ddf9565b6060600061d2cf61a28c565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161d2ec57905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061d3479061f4fa565b935060ff168151811061d35c5761d35c61f38d565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161d3ad919061fc65565b60405160208183030381529060405282828061d3c89061f4fa565b935060ff168151811061d3dd5761d3dd61f38d565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061d42a9061f4fa565b935060ff168151811061d43f5761d43f61f38d565b60200260200101819052508260405160200161d45b919061f428565b60405160208183030381529060405282828061d4769061f4fa565b935060ff168151811061d48b5761d48b61f38d565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061d4d89061f4fa565b935060ff168151811061d4ed5761d4ed61f38d565b602002602001018190525061d502878461de73565b828261d50d8161f4fa565b935060ff168151811061d5225761d52261f38d565b60209081029190910101528551511561d5ce5760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261d5748161f4fa565b935060ff168151811061d5895761d58961f38d565b602002602001018190525061d5a286600001518461de73565b828261d5ad8161f4fa565b935060ff168151811061d5c25761d5c261f38d565b60200260200101819052505b85608001511561d63c5760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261d6178161f4fa565b935060ff168151811061d62c5761d62c61f38d565b602002602001018190525061d6a2565b841561d6a25760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261d6818161f4fa565b935060ff168151811061d6965761d69661f38d565b60200260200101819052505b6040860151511561d73e5760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261d6ec8161f4fa565b935060ff168151811061d7015761d70161f38d565b6020026020010181905250856040015182828061d71d9061f4fa565b935060ff168151811061d7325761d73261f38d565b60200260200101819052505b85606001511561d7a85760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261d7878161f4fa565b935060ff168151811061d79c5761d79c61f38d565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561d7c65761d7c661e715565b60405190808252806020026020018201604052801561d7f957816020015b606081526020019060019003908161d7e45790505b50905060005b8260ff168160ff16101561d85257838160ff168151811061d8225761d82261f38d565b6020026020010151828260ff168151811061d83f5761d83f61f38d565b602090810291909101015260010161d7ff565b50979650505050505050565b604080518082019091526000808252602082015281518351101561d883575081619db8565b8151835160208501516000929161d8999161e8cd565b61d8a3919061e8ba565b6020840151909150600190821461d8c4575082516020840151819020908220145b801561d8df5783518551869061d8db90839061e8ba565b9052505b50929392505050565b600080826000015161d90c856000015186602001518660000151876020015161dc2e565b61d916919061e8cd565b90505b8351602085015161d92a919061e8cd565b811161d253578161d93a8161fcaa565b925050826000015161d97185602001518361d955919061e8ba565b865161d961919061e8ba565b838660000151876020015161dc2e565b61d97b919061e8cd565b905061d919565b6060600061d990848461d8e8565b61d99b90600161e8cd565b67ffffffffffffffff81111561d9b35761d9b361e715565b60405190808252806020026020018201604052801561d9e657816020015b606081526020019060019003908161d9d15790505b50905060005b815181101561bc5b5761da0261be7d868661d234565b82828151811061da145761da1461f38d565b602090810291909101015260010161d9ec565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061da70577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061da9c576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061daba57662386f26fc10000830492506010015b6305f5e100831061dad2576305f5e100830492506008015b612710831061dae657612710830492506004015b6064831061daf8576064830492506002015b600a8310619db85760010192915050565b600061db15838361deb3565b159392505050565b60008085841161dc24576020841161dbd0576000841561db6857600161db4486602061e8ba565b61db4f90600861fcc4565b61db5a90600261fdc2565b61db64919061e8ba565b1990505b835181168561db77898961e8cd565b61db81919061e8ba565b805190935082165b81811461dbbb5787841161dba3578794505050505061b88c565b8361dbad8161fdce565b94505082845116905061db89565b61dbc5878561e8cd565b94505050505061b88c565b83832061dbdd858861e8ba565b61dbe7908761e8cd565b91505b85821061dc225784822080820361dc0f5761dc05868461e8cd565b935050505061b88c565b61dc1a60018461e8ba565b92505061dbea565b505b5092949350505050565b6000838186851161dd39576020851161dce8576000851561dc7a57600161dc5687602061e8ba565b61dc6190600861fcc4565b61dc6c90600261fdc2565b61dc76919061e8ba565b1990505b8451811660008761dc8b8b8b61e8cd565b61dc95919061e8ba565b855190915083165b82811461dcda5781861061dcc25761dcb58b8b61e8cd565b965050505050505061b88c565b8561dccc8161fcaa565b96505083865116905061dc9d565b85965050505050505061b88c565b508383206000905b61dcfa868961e8ba565b821161dd375785832080820361dd16578394505050505061b88c565b61dd2160018561e8cd565b935050818061dd2f9061fcaa565b92505061dcf0565b505b61dd43878761e8cd565b979650505050505050565b6040805180820190915260008082526020820152600061dd80856000015186602001518660000151876020015161dc2e565b60208087018051918601919091525190915061dd9c908261e8ba565b83528451602086015161ddaf919061e8cd565b810361ddbe576000855261ddf0565b8351835161ddcc919061e8cd565b8551869061dddb90839061e8ba565b905250835161ddea908261e8cd565b60208601525b50909392505050565b6020811061de31578151835261de1060208461e8cd565b925061de1d60208361e8cd565b915061de2a60208261e8ba565b905061ddf9565b600019811561de6057600161de4783602061e8ba565b61de539061010061fdc2565b61de5d919061e8ba565b90505b9151835183169219169190911790915250565b6060600061de81848461a35f565b805160208083015160405193945061de9b9390910161fde5565b60405160208183030381529060405291505092915050565b815181516000919081111561dec6575081515b6020808501519084015160005b8381101561df7f578251825180821461df4f57600019602087101561df2e5760018461df0089602061e8ba565b61df0a919061e8cd565b61df1590600861fcc4565b61df2090600261fdc2565b61df2a919061e8ba565b1990505b818116838216818103911461df4c579750619db89650505050505050565b50505b61df5a60208661e8cd565b945061df6760208561e8cd565b9350505060208161df78919061e8cd565b905061ded3565b508451865161aa0d919061fe3d565b610b67806200fe5e83390190565b61092380620109c583390190565b61106f80620112e883390190565b612072806201235783390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161e00961e00e565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161e0096040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561e0c05783516001600160a01b031683526020938401939092019160010161e099565b509095945050505050565b60005b8381101561e0e657818101518382015260200161e0ce565b50506000910152565b6000815180845261e10781602086016020860161e0cb565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561e217577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561e1fd577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261e1e784865161e0ef565b602095860195909450929092019160010161e1ad565b50919750505060209485019492909201915060010161e143565b50929695505050505050565b600081518084526020840193506020830160005b8281101561e2775781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161e237565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561e217577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261e2ed604088018261e0ef565b905060208201519150868103602088015261e308818361e223565b96505050602093840193919091019060010161e2a9565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561e217577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261e38185835161e0ef565b9450602093840193919091019060010161e347565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561e217577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261e417604087018261e223565b955050602093840193919091019060010161e3be565b600181811c9082168061e44157607f821691505b60208210810361c158577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000815461e4878161e42d565b80855260018216801561e4a1576001811461e4db5761e512565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083166020870152602082151560051b870101935061e512565b84600052602060002060005b8381101561e5095781546020828a01015260018201915060208101905061e4e7565b87016020019450505b50505092915050565b6001600160a01b0381541682526001600160a01b03600182015416602083015260028101546040830152608060608301526000619f93608084016003840161e47a565b602081526000619f93602083018461e51b565b6001600160a01b038316815260406020820152600061b88c604083018461e51b565b6001600160a01b038416815260606020820152600061e5b5606083018561e0ef565b828103604084015261aa0d818561e0ef565b80516001600160a01b038116811461e5de57600080fd5b919050565b60006020828403121561e5f557600080fd5b619f938261e5c7565b6003811061e635577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff8816604083015286606083015261e6c9608083018761e5fe565b8460a08301526001600160a01b03841660c083015261dd4360e08301846001600160a01b03169052565b60006020828403121561e70557600080fd5b81518015158114619f9357600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561280357806000526020600020601f840160051c8101602085101561e76b5750805b601f840160051c820191505b81811015612b0d576000815560010161e777565b815167ffffffffffffffff81111561e7a55761e7a561e715565b61e7b98161e7b3845461e42d565b8461e744565b6020601f82116001811461e7ed576000831561e7d55750848201515b600019600385901b1c1916600184901b178455612b0d565b600084815260208120601f198516915b8281101561e81d578785015182556020948501946001909201910161e7fd565b508482101561e83b5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60006020828403121561e85c57600080fd5b5051919050565b8381526001600160a01b038316602082015260606040820152600061c13e606083018461e51b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810381811115619db857619db861e88b565b80820180821115619db857619db861e88b565b60c08252600061e8f360c084018361e47a565b6001600160a01b0360018401541660208501526002830154604085015260ff600384015416151560608501526004830154608085015283810360a0850152619db4816005850161e47a565b602081526000619f93602083018461e8e0565b6001600160a01b038316815260406020820152600061b88c604083018461e8e0565b602081526000619f93602083018461e0ef565b602081526000619db860208301600581527f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260400190565b60a08152600061e9d660a083018761e0ef565b6001600160a01b03861660208401528460408401526001600160a01b0384166060840152828103608084015261dd4381600581527f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260400190565b600081516060845261ea4b606085018261e0ef565b90506001600160a01b036020840151166020850152604083015160408501528091505092915050565b60808152600061ea87608083018761ea36565b8560208401526001600160a01b0385166040840152828103606084015261dd43818561e0ef565b60a08152600061eac160a083018861ea36565b6001600160a01b03871660208401528560408401526001600160a01b0385166060840152828103608084015261eaf7818561e0ef565b98975050505050505050565b6000806040838503121561eb1657600080fd5b61eb1f8361e5c7565b6020939093015192949293505050565b600081546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b03600183015416604084015260a0606084015261eb7860a084016002840161e47a565b600383015460808501528091505092915050565b60e08152600061eb9f60e083018961e0ef565b8760208401526001600160a01b0387166040840152828103606084015261ebc6818761e0ef565b855160808501526020860151151560a0850152905082810360c084015261ebed818561eb2f565b9998505050505050505050565b6001600160a01b03851681528360208201526001600160a01b038316604082015260806060820152600061aa0d608083018461e51b565b60c08152600061ec4460c083018861e0ef565b6001600160a01b0387166020840152828103604084015261ec65818761e0ef565b85516060850152602086015115156080850152905082810360a084015261eaf7818561eb2f565b610100815260066101008201527f5345434f4e440000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f534543000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff8816604083015286606083015261e6c9608083018761e5fe565b60808152600061ed2f608083018761e0ef565b8560208401526001600160a01b0385166040840152828103606084015261dd43818561eb2f565b83815282602082015260606040820152600061c13e606083018461e0ef565b6001600160a01b038316815260406020820152600061b88c604083018461e0ef565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161edcf81601a85016020880161e0cb565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161ee0c81601c84016020880161e0cb565b01601c01949350505050565b6040516060810167ffffffffffffffff8111828210171561ee3b5761ee3b61e715565b60405290565b60008067ffffffffffffffff84111561ee5c5761ee5c61e715565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561ee8b5761ee8b61e715565b60405283815290508082840185101561eea357600080fd5b61bc5b84602083018561e0cb565b600082601f83011261eec257600080fd5b619f938383516020850161ee41565b60006020828403121561eee357600080fd5b815167ffffffffffffffff81111561eefa57600080fd5b619db48482850161eeb1565b6000835161ef1881846020880161e0cb565b83519083019061ef2c81836020880161e0cb565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161ef6d81601a85016020880161e0cb565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161efaa81603384016020880161e0cb565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000619f93608083018461e0ef565b60006020828403121561f03957600080fd5b815167ffffffffffffffff81111561f05057600080fd5b8201601f8101841361f06157600080fd5b619db48482516020840161ee41565b6000855161f082818460208a0161e0cb565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161f0bc816001840160208a0161e0cb565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161f0fa81600284016020890161e0cb565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161f13c81600284016020880161e0cb565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061f187604083018461e0ef565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161f1fe81601f85016020870161e0cb565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061f26b604083018461e0ef565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061f2bd604083018461e0ef565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161f33481601485016020870161e0cb565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061f37b604083018561e0ef565b8281036020840152619f8f818561e0ef565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161f3f481600185016020870161e0cb565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161f43a81846020870161e0cb565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161f4ed81604b85016020870161e0cb565b91909101604b0192915050565b600060ff821660ff810361f5105761f51061e88b565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161f57781602985016020870161e0cb565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000619f93608083018461e0ef565b60006020828403121561f5dd57600080fd5b815167ffffffffffffffff81111561f5f457600080fd5b82016060818503121561f60657600080fd5b61f60e61ee18565b81518060030b811461f61f57600080fd5b8152602082015167ffffffffffffffff81111561f63b57600080fd5b61f6478682850161eeb1565b602083015250604082015167ffffffffffffffff81111561f66757600080fd5b61f6738682850161eeb1565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161f6df81602185016020870161e0cb565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161f8cb81602185016020880161e0cb565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161f90881602e84016020880161e0cb565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161f57781602985016020870161e0cb565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161f9d081602285016020870161e0cb565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161fa1581600e85016020870161e0cb565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161faf381601885016020880161e0cb565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161fb3081601c84016020880161e0cb565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161fc3681846020870161e0cb565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161fc9d81601c85016020870161e0cb565b91909101601c0192915050565b6000600019820361fcbd5761fcbd61e88b565b5060010190565b8082028115828204841417619db857619db861e88b565b6001815b600184111561fd165780850481111561fcfa5761fcfa61e88b565b600184161561fd0857908102905b60019390931c92800261fcdf565b935093915050565b60008261fd2d57506001619db8565b8161fd3a57506000619db8565b816001811461fd50576002811461fd5a5761fd76565b6001915050619db8565b60ff84111561fd6b5761fd6b61e88b565b50506001821b619db8565b5060208310610133831016604e8410600b841016171561fd99575081810a619db8565b61fda6600019848461fcdb565b806000190482111561fdba5761fdba61e88b565b029392505050565b6000619f93838361fd1e565b60008161fddd5761fddd61e88b565b506000190190565b6000835161fdf781846020880161e0cb565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161fe3181600184016020880161e0cb565b01600101949350505050565b818103600083128015838313168383128216171561d2535761d25361e88b56fe60c0604052600d60809081526c2bb930b83832b21022ba3432b960991b60a05260009061002c9082610114565b506040805180820190915260048152630ae8aa8960e31b60208201526001906100559082610114565b506002805460ff1916601217905534801561006f57600080fd5b506101d2565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061009f57607f821691505b6020821081036100bf57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561010f57806000526020600020601f840160051c810160208510156100ec5750805b601f840160051c820191505b8181101561010c57600081556001016100f8565b50505b505050565b81516001600160401b0381111561012d5761012d610075565b6101418161013b845461008b565b846100c5565b6020601f821160018114610175576000831561015d5750848201515b600019600385901b1c1916600184901b17845561010c565b600084815260208120601f198516915b828110156101a55787850151825560209485019460019092019101610185565b50848210156101c35786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610986806101e16000396000f3fe6080604052600436106100c05760003560e01c8063313ce56711610074578063a9059cbb1161004e578063a9059cbb146101fa578063d0e30db01461021a578063dd62ed3e1461022257600080fd5b8063313ce5671461018c57806370a08231146101b857806395d89b41146101e557600080fd5b806318160ddd116100a557806318160ddd1461012f57806323b872dd1461014c5780632e1a7d4d1461016c57600080fd5b806306fdde03146100d4578063095ea7b3146100ff57600080fd5b366100cf576100cd61025a565b005b600080fd5b3480156100e057600080fd5b506100e96102b5565b6040516100f69190610745565b60405180910390f35b34801561010b57600080fd5b5061011f61011a3660046107da565b610343565b60405190151581526020016100f6565b34801561013b57600080fd5b50475b6040519081526020016100f6565b34801561015857600080fd5b5061011f610167366004610804565b6103bd565b34801561017857600080fd5b506100cd610187366004610841565b610647565b34801561019857600080fd5b506002546101a69060ff1681565b60405160ff90911681526020016100f6565b3480156101c457600080fd5b5061013e6101d336600461085a565b60036020526000908152604090205481565b3480156101f157600080fd5b506100e9610724565b34801561020657600080fd5b5061011f6102153660046107da565b610731565b6100cd61025a565b34801561022e57600080fd5b5061013e61023d366004610875565b600460209081526000928352604080842090915290825290205481565b33600090815260036020526040812080543492906102799084906108d7565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b600080546102c2906108ea565b80601f01602080910402602001604051908101604052809291908181526020018280546102ee906108ea565b801561033b5780601f106103105761010080835404028352916020019161033b565b820191906000526020600020905b81548152906001019060200180831161031e57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103ab9086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604081205482111561042b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841633148015906104a1575073ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156105605773ffffffffffffffffffffffffffffffffffffffff8416600090815260046020908152604080832033845290915290205482111561051a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b73ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091528120805484929061055a90849061093d565b90915550505b73ffffffffffffffffffffffffffffffffffffffff84166000908152600360205260408120805484929061059590849061093d565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040812080548492906105cf9084906108d7565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161063591815260200190565b60405180910390a35060019392505050565b3360009081526003602052604090205481111561069a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b33600090815260036020526040812080548392906106b990849061093d565b9091555050604051339082156108fc029083906000818181858888f193505050501580156106eb573d6000803e3d6000fd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b600180546102c2906108ea565b600061073e3384846103bd565b9392505050565b602081526000825180602084015260005b818110156107735760208186018101516040868401015201610756565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146107d557600080fd5b919050565b600080604083850312156107ed57600080fd5b6107f6836107b1565b946020939093013593505050565b60008060006060848603121561081957600080fd5b610822846107b1565b9250610830602085016107b1565b929592945050506040919091013590565b60006020828403121561085357600080fd5b5035919050565b60006020828403121561086c57600080fd5b61073e826107b1565b6000806040838503121561088857600080fd5b610891836107b1565b915061089f602084016107b1565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156103b7576103b76108a8565b600181811c908216806108fe57607f821691505b602082108103610937577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b818103818111156103b7576103b76108a856fea2646970667358221220bb5b9dcef0ba90bcdefcbd63f71b1df95b50e29550a7456c69c6b9ff9dcdd20e64736f6c634300081a003360a060405234801561001057600080fd5b50737c591652f159496b14e15616f0948a6d63b585e86001600160a01b031663cc5ad8b66040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610065573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610089919061009a565b6001600160a01b03166080526100ca565b6000602082840312156100ac57600080fd5b81516001600160a01b03811681146100c357600080fd5b9392505050565b60805161083f6100e460003960006073015261083f6000f3fe6080604052600436106100585760003560e01c80635bcfd616116100405780635bcfd616146100de5780637b103999146100fe578063c9028a361461012657005b8063116191b6146100615780632d4cfb7e146100be57005b3661005f57005b005b34801561006d57600080fd5b506100957f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100ca57600080fd5b5061005f6100d936600461022a565b610146565b3480156100ea57600080fd5b5061005f6100f9366004610295565b610180565b34801561010a57600080fd5b50610095737c591652f159496b14e15616f0948a6d63b585e881565b34801561013257600080fd5b5061005f610141366004610352565b6101fb565b7f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7816040516101759190610441565b60405180910390a150565b606081156101975761019482840184610547565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e6101c2878061063d565b6101d260408a0160208b016106a2565b896040013533866040516101eb969594939291906106bd565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c481604051610175919061077f565b60006020828403121561023c57600080fd5b813567ffffffffffffffff81111561025357600080fd5b820160c0818503121561026557600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461029057600080fd5b919050565b6000806000806000608086880312156102ad57600080fd5b853567ffffffffffffffff8111156102c457600080fd5b8601606081890312156102d657600080fd5b94506102e46020870161026c565b935060408601359250606086013567ffffffffffffffff81111561030757600080fd5b8601601f8101881361031857600080fd5b803567ffffffffffffffff81111561032f57600080fd5b88602082840101111561034157600080fd5b959894975092955050506020019190565b60006020828403121561036457600080fd5b813567ffffffffffffffff81111561037b57600080fd5b82016080818503121561026557600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126103c257600080fd5b830160208101925035905067ffffffffffffffff8111156103e257600080fd5b8036038213156103f157600080fd5b9250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b602081526000610451838461038d565b60c0602085015261046660e0850182846103f8565b91505073ffffffffffffffffffffffffffffffffffffffff61048a6020860161026c565b1660408401526000604085013590508060608501525060608401358015158082146104b457600080fd5b80608086015250506000608085013590508060a0850152506104d960a085018561038d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c086015261050e8382846103f8565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561055957600080fd5b813567ffffffffffffffff81111561057057600080fd5b8201601f8101841361058157600080fd5b803567ffffffffffffffff81111561059b5761059b610518565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561060757610607610518565b60405281815282820160200186101561061f57600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261067257600080fd5b83018035915067ffffffffffffffff82111561068d57600080fd5b6020019150368190038213156103f157600080fd5b6000602082840312156106b457600080fd5b6102658261026c565b60a0815260006106d160a08301888a6103f8565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b8181101561073b5760208187018101518483018201520161071f565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff6107a18361026c565b16602082015273ffffffffffffffffffffffffffffffffffffffff6107c86020840161026c565b16604082015260008060408401359050806060840152506107ec606084018461038d565b60808085015261080060a0850182846103f8565b9594505050505056fea2646970667358221220c50bb9baa823469eb1747a377ba052ff0a9912b1d8555ea50ed5f46a062dd81564736f6c634300081a003360c060405234801561001057600080fd5b5060405161106f38038061106f83398101604081905261002f916100db565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006357604051632b2add3d60e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0385811691909117909155828116608052811660a0526040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac590600090a150505061011e565b80516001600160a01b03811681146100d657600080fd5b919050565b6000806000606084860312156100f057600080fd5b6100f9846100bf565b9250610107602085016100bf565b9150610115604085016100bf565b90509250925092565b60805160a051610f2561014a60003960006101e50152600081816102b9015261045b0152610f256000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806397770dff11610097578063c63585cc11610066578063c63585cc14610273578063d7fd7afb14610286578063d936a012146102b4578063ee2815ba146102db57600080fd5b806397770dff1461021a578063a7cb05071461022d578063c39aca3714610240578063c62178ac1461025357600080fd5b8063513a9c05116100d3578063513a9c051461018a578063569541b9146101c0578063842da36d146101e057806391dd645f1461020757600080fd5b80630be15547146100fa5780631f0e251b1461015a5780633ce4a5bc1461016f575b600080fd5b610130610108366004610bd1565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61016d610168366004610c13565b6102ee565b005b61013073735b14bb79463307aacbed86daf3322b1e6226ab81565b610130610198366004610bd1565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101309073ffffffffffffffffffffffffffffffffffffffff1681565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d610215366004610c35565b610402565b61016d610228366004610c13565b610526565b61016d61023b366004610c61565b610633565b61016d61024e366004610c83565b6106ce565b6004546101309073ffffffffffffffffffffffffffffffffffffffff1681565b610130610281366004610d53565b6108cd565b6102a6610294366004610bd1565b60006020819052908152604090205481565b604051908152602001610151565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d6102e9366004610c35565b610a02565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461033b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610388576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461044f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354600090610497907f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff16846108cd565b60008481526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251878152918201529192507f0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e910160405180910390a1505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610573576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166105c0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e906020016103f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610680576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461071b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831673735b14bb79463307aacbed86daf3322b1e6226ab1480610768575073ffffffffffffffffffffffffffffffffffffffff831630145b1561079f576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690528616906347e7ef24906044016020604051808303816000875af1158015610814573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108389190610d96565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063de43156e906108939089908990899088908890600401610e01565b600060405180830381600087803b1580156108ad57600080fd5b505af11580156108c1573d6000803e3d6000fd5b50505050505050505050565b60008060006108dc8585610ad3565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b16603482015291935091508690604801604051602081830303815290604052805190602001206040516020016109c29291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a4f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d91016106c2565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b3b576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610b75578284610b78565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610bca576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b600060208284031215610be357600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c0e57600080fd5b919050565b600060208284031215610c2557600080fd5b610c2e82610bea565b9392505050565b60008060408385031215610c4857600080fd5b82359150610c5860208401610bea565b90509250929050565b60008060408385031215610c7457600080fd5b50508035926020909101359150565b60008060008060008060a08789031215610c9c57600080fd5b863567ffffffffffffffff811115610cb357600080fd5b87016060818a031215610cc557600080fd5b9550610cd360208801610bea565b945060408701359350610ce860608801610bea565b9250608087013567ffffffffffffffff811115610d0457600080fd5b8701601f81018913610d1557600080fd5b803567ffffffffffffffff811115610d2c57600080fd5b896020828401011115610d3e57600080fd5b60208201935080925050509295509295509295565b600080600060608486031215610d6857600080fd5b610d7184610bea565b9250610d7f60208501610bea565b9150610d8d60408501610bea565b90509250925092565b600060208284031215610da857600080fd5b81518015158114610c2e57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60808152600086357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018112610e3957600080fd5b870160208101903567ffffffffffffffff811115610e5657600080fd5b803603821315610e6557600080fd5b60606080850152610e7a60e085018284610db8565b91505073ffffffffffffffffffffffffffffffffffffffff610e9e60208a01610bea565b1660a0840152604088013560c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401528281036060840152610ee3818587610db8565b9897505050505050505056fea26469706673582212203ba16321030dc78e78d6c63054e6c38a60adfb895f33238a0ee70b193dd60fbb64736f6c634300081a003360c060405234801561001057600080fd5b5060405161207238038061207283398101604081905261002f916101f0565b6001600160a01b038216158061004c57506001600160a01b038116155b1561006a5760405163d92e233d60e01b815260040160405180910390fd5b60066100768982610342565b5060076100838882610342565b506008805460ff191660ff881617905560808590528360028111156100aa576100aa610400565b60a08160028111156100be576100be610400565b905250600192909255600080546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506104169350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261013357600080fd5b81516001600160401b0381111561014c5761014c61010c565b604051601f8201601f19908116603f011681016001600160401b038111828210171561017a5761017a61010c565b60405281815283820160200185101561019257600080fd5b60005b828110156101b157602081860181015183830182015201610195565b506000918101602001919091529392505050565b8051600381106101d457600080fd5b919050565b80516001600160a01b03811681146101d457600080fd5b600080600080600080600080610100898b03121561020d57600080fd5b88516001600160401b0381111561022357600080fd5b61022f8b828c01610122565b60208b015190995090506001600160401b0381111561024d57600080fd5b6102598b828c01610122565b975050604089015160ff8116811461027057600080fd5b60608a0151909650945061028660808a016101c5565b60a08a0151909450925061029c60c08a016101d9565b91506102aa60e08a016101d9565b90509295985092959890939650565b600181811c908216806102cd57607f821691505b6020821081036102ed57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561033d57806000526020600020601f840160051c8101602085101561031a5750805b601f840160051c820191505b8181101561033a5760008155600101610326565b50505b505050565b81516001600160401b0381111561035b5761035b61010c565b61036f8161036984546102b9565b846102f3565b6020601f8211600181146103a3576000831561038b5750848201515b600019600385901b1c1916600184901b17845561033a565b600084815260208120601f198516915b828110156103d357878501518255602094850194600190920191016103b3565b50848210156103f15786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a051611c1b61045760003960006103440152600081816102f001528181610bdc01528181610ce201528181610efe01526110040152611c1b6000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806395d89b41116100f9578063ccc7759911610097578063eddeb12311610071578063eddeb12314610461578063f2441b3214610474578063f687d12a14610494578063fc5fecd5146104a757600080fd5b8063ccc77599146103d4578063d9eeebed146103e7578063dd62ed3e1461041b57600080fd5b8063b84c8246116100d3578063b84c824614610386578063c47f00271461039b578063c7012626146103ae578063c835d7cc146103c157600080fd5b806395d89b4114610337578063a3413d031461033f578063a9059cbb1461037357600080fd5b80633ce4a5bc116101665780634d8943bb116101405780634d8943bb146102ac57806370a08231146102b557806385e1f4d0146102eb5780638b851b951461031257600080fd5b80633ce4a5bc1461024657806342966c681461028657806347e7ef241461029957600080fd5b806318160ddd1161019757806318160ddd1461021657806323b872dd1461021e578063313ce5671461023157600080fd5b806306fdde03146101be578063091d2788146101dc578063095ea7b3146101f3575b600080fd5b6101c66104ba565b6040516101d39190611648565b60405180910390f35b6101e560015481565b6040519081526020016101d3565b610206610201366004611687565b61054c565b60405190151581526020016101d3565b6005546101e5565b61020661022c3660046116b3565b610563565b60085460405160ff90911681526020016101d3565b61026173735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6102066102943660046116f4565b6105fa565b6102066102a7366004611687565b61060e565b6101e560025481565b6101e56102c336600461170d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101e57f000000000000000000000000000000000000000000000000000000000000000081565b60085461026190610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101c6610767565b6103667f000000000000000000000000000000000000000000000000000000000000000081565b6040516101d3919061172a565b610206610381366004611687565b610776565b610399610394366004611832565b610783565b005b6103996103a9366004611832565b6107e0565b6102066103bc366004611883565b610839565b6103996103cf36600461170d565b610988565b6103996103e236600461170d565b610a9c565b6103ef610bb0565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101d3565b6101e56104293660046118dc565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b61039961046f3660046116f4565b610dce565b6000546102619073ffffffffffffffffffffffffffffffffffffffff1681565b6103996104a23660046116f4565b610e50565b6103ef6104b53660046116f4565b610ed2565b6060600680546104c990611915565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590611915565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b5050505050905090565b60006105593384846110ee565b5060015b92915050565b60006105708484846111f7565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054828110156105db576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ef85336105ea8685611997565b6110ee565b506001949350505050565b600061060633836113b2565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab1480159061064c575060005473ffffffffffffffffffffffffffffffffffffffff163314155b80156106755750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b156106ac576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106b683836114f4565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107569186906119aa565b60405180910390a250600192915050565b6060600780546104c990611915565b60006105593384846111f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107d0576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107dc8282611a1b565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461082d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107dc8282611a1b565b6000806000610846610bb0565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156108d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fc9190611b34565b610932576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61093c33856113b2565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161097591899189918791611b56565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109d5576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a22576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610ae9576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b36576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c679190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610cb6576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d699190611ba2565b905080600003610da5576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610db89190611bbb565b610dc29190611bd2565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e1b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a91565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e9d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f899190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610fd8576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015611067573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108b9190611ba2565b9050806000036110c7576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000906110d78784611bbb565b6110e19190611bd2565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661113b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611188576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611244576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611291576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040902054818110156112f1576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112fb8282611997565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260036020526040808220939093559085168152908120805484929061133e908490611bd2565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113a491815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113ff576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260409020548181101561145f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114698282611997565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812091909155600580548492906114a4908490611997565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111ea565b73ffffffffffffffffffffffffffffffffffffffff8216611541576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008282546115539190611bd2565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805483929061158d908490611bd2565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561160a576020818501810151868301820152016115ee565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061165b60208301846115e4565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461168457600080fd5b50565b6000806040838503121561169a57600080fd5b82356116a581611662565b946020939093013593505050565b6000806000606084860312156116c857600080fd5b83356116d381611662565b925060208401356116e381611662565b929592945050506040919091013590565b60006020828403121561170657600080fd5b5035919050565b60006020828403121561171f57600080fd5b813561165b81611662565b6020810160038310611765577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff8411156117b5576117b561176b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156118025761180261176b565b60405283815290508082840185101561181a57600080fd5b83836020830137600060208583010152509392505050565b60006020828403121561184457600080fd5b813567ffffffffffffffff81111561185b57600080fd5b8201601f8101841361186c57600080fd5b61187b8482356020840161179a565b949350505050565b6000806040838503121561189657600080fd5b823567ffffffffffffffff8111156118ad57600080fd5b8301601f810185136118be57600080fd5b6118cd8582356020840161179a565b95602094909401359450505050565b600080604083850312156118ef57600080fd5b82356118fa81611662565b9150602083013561190a81611662565b809150509250929050565b600181811c9082168061192957607f821691505b602082108103611962577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561055d5761055d611968565b6040815260006119bd60408301856115e4565b90508260208301529392505050565b601f821115611a1657806000526020600020601f840160051c810160208510156119f35750805b601f840160051c820191505b81811015611a1357600081556001016119ff565b50505b505050565b815167ffffffffffffffff811115611a3557611a3561176b565b611a4981611a438454611915565b846119cc565b6020601f821160018114611a9b5760008315611a655750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611a13565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611ae95787850151825560209485019460019092019101611ac9565b5084821015611b2557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b600060208284031215611b4657600080fd5b8151801515811461165b57600080fd5b608081526000611b6960808301876115e4565b6020830195909552506040810192909252606090910152919050565b600060208284031215611b9757600080fd5b815161165b81611662565b600060208284031215611bb457600080fd5b5051919050565b808202811582820484141761055d5761055d611968565b8082018082111561055d5761055d61196856fea26469706673582212208307d60e253f5034856b93df00e3e2f46b06f9765d906dbd93ee947935fc608764736f6c634300081a003347617320666565206e6f74206275726e656420636f72726563746c792066726f6d20676173205a524332305a5243323020746f6b656e732077657265206e6f74206275726e656420636f72726563746c795769746864726177616c20616d6f756e74206e6f74206275726e656420636f72726563746c792066726f6d207365636f6e64205a524332305a5243323020746f6b656e732077657265206e6f74206275726e656420636f72726563746c7920666f72207769746864726177416e6443616c6ca26469706673582212204237b78314ef3207219461e1c70b3e48b706804293b5993b78f55523872b333064736f6c634300081a0033",
}

// GatewayZEVMOutboundTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayZEVMOutboundTestMetaData.ABI instead.
var GatewayZEVMOutboundTestABI = GatewayZEVMOutboundTestMetaData.ABI

// GatewayZEVMOutboundTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayZEVMOutboundTestMetaData.Bin instead.
var GatewayZEVMOutboundTestBin = GatewayZEVMOutboundTestMetaData.Bin

// DeployGatewayZEVMOutboundTest deploys a new Ethereum contract, binding an instance of GatewayZEVMOutboundTest to it.
func DeployGatewayZEVMOutboundTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayZEVMOutboundTest, error) {
	parsed, err := GatewayZEVMOutboundTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayZEVMOutboundTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayZEVMOutboundTest{GatewayZEVMOutboundTestCaller: GatewayZEVMOutboundTestCaller{contract: contract}, GatewayZEVMOutboundTestTransactor: GatewayZEVMOutboundTestTransactor{contract: contract}, GatewayZEVMOutboundTestFilterer: GatewayZEVMOutboundTestFilterer{contract: contract}}, nil
}

// GatewayZEVMOutboundTest is an auto generated Go binding around an Ethereum contract.
type GatewayZEVMOutboundTest struct {
	GatewayZEVMOutboundTestCaller     // Read-only binding to the contract
	GatewayZEVMOutboundTestTransactor // Write-only binding to the contract
	GatewayZEVMOutboundTestFilterer   // Log filterer for contract events
}

// GatewayZEVMOutboundTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMOutboundTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMOutboundTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayZEVMOutboundTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMOutboundTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayZEVMOutboundTestSession struct {
	Contract     *GatewayZEVMOutboundTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GatewayZEVMOutboundTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayZEVMOutboundTestCallerSession struct {
	Contract *GatewayZEVMOutboundTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// GatewayZEVMOutboundTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayZEVMOutboundTestTransactorSession struct {
	Contract     *GatewayZEVMOutboundTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// GatewayZEVMOutboundTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestRaw struct {
	Contract *GatewayZEVMOutboundTest // Generic contract binding to access the raw methods on
}

// GatewayZEVMOutboundTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestCallerRaw struct {
	Contract *GatewayZEVMOutboundTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayZEVMOutboundTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestTransactorRaw struct {
	Contract *GatewayZEVMOutboundTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayZEVMOutboundTest creates a new instance of GatewayZEVMOutboundTest, bound to a specific deployed contract.
func NewGatewayZEVMOutboundTest(address common.Address, backend bind.ContractBackend) (*GatewayZEVMOutboundTest, error) {
	contract, err := bindGatewayZEVMOutboundTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTest{GatewayZEVMOutboundTestCaller: GatewayZEVMOutboundTestCaller{contract: contract}, GatewayZEVMOutboundTestTransactor: GatewayZEVMOutboundTestTransactor{contract: contract}, GatewayZEVMOutboundTestFilterer: GatewayZEVMOutboundTestFilterer{contract: contract}}, nil
}

// NewGatewayZEVMOutboundTestCaller creates a new read-only instance of GatewayZEVMOutboundTest, bound to a specific deployed contract.
func NewGatewayZEVMOutboundTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayZEVMOutboundTestCaller, error) {
	contract, err := bindGatewayZEVMOutboundTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestCaller{contract: contract}, nil
}

// NewGatewayZEVMOutboundTestTransactor creates a new write-only instance of GatewayZEVMOutboundTest, bound to a specific deployed contract.
func NewGatewayZEVMOutboundTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayZEVMOutboundTestTransactor, error) {
	contract, err := bindGatewayZEVMOutboundTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestTransactor{contract: contract}, nil
}

// NewGatewayZEVMOutboundTestFilterer creates a new log filterer instance of GatewayZEVMOutboundTest, bound to a specific deployed contract.
func NewGatewayZEVMOutboundTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayZEVMOutboundTestFilterer, error) {
	contract, err := bindGatewayZEVMOutboundTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestFilterer{contract: contract}, nil
}

// bindGatewayZEVMOutboundTest binds a generic wrapper to an already deployed contract.
func bindGatewayZEVMOutboundTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayZEVMOutboundTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMOutboundTest.Contract.GatewayZEVMOutboundTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.GatewayZEVMOutboundTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.GatewayZEVMOutboundTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMOutboundTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ISTEST() (bool, error) {
	return _GatewayZEVMOutboundTest.Contract.ISTEST(&_GatewayZEVMOutboundTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ISTEST() (bool, error) {
	return _GatewayZEVMOutboundTest.Contract.ISTEST(&_GatewayZEVMOutboundTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVMOutboundTest.Contract.PAUSERROLE(&_GatewayZEVMOutboundTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVMOutboundTest.Contract.PAUSERROLE(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeArtifacts(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeArtifacts(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeContracts(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeContracts(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeSenders(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeSenders(&_GatewayZEVMOutboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) Failed() (bool, error) {
	return _GatewayZEVMOutboundTest.Contract.Failed(&_GatewayZEVMOutboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) Failed() (bool, error) {
	return _GatewayZEVMOutboundTest.Contract.Failed(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetArtifactSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetArtifactSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetArtifacts(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetArtifacts(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetContracts(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetContracts(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetInterfaces(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetInterfaces(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetSenders(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetSenders(&_GatewayZEVMOutboundTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.SetUp(&_GatewayZEVMOutboundTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.SetUp(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForDifferentZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x75ca5587.
//
// Solidity: function testBurnGasFeeForDifferentZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnGasFeeForDifferentZRC20WithdrawAndCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnGasFeeForDifferentZRC20WithdrawAndCall")
}

// TestBurnGasFeeForDifferentZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x75ca5587.
//
// Solidity: function testBurnGasFeeForDifferentZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnGasFeeForDifferentZRC20WithdrawAndCall() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForDifferentZRC20WithdrawAndCall(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForDifferentZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x75ca5587.
//
// Solidity: function testBurnGasFeeForDifferentZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnGasFeeForDifferentZRC20WithdrawAndCall() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForDifferentZRC20WithdrawAndCall(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForDifferentZRC20Withdrawal is a paid mutator transaction binding the contract method 0xb84a3d2f.
//
// Solidity: function testBurnGasFeeForDifferentZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnGasFeeForDifferentZRC20Withdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnGasFeeForDifferentZRC20Withdrawal")
}

// TestBurnGasFeeForDifferentZRC20Withdrawal is a paid mutator transaction binding the contract method 0xb84a3d2f.
//
// Solidity: function testBurnGasFeeForDifferentZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnGasFeeForDifferentZRC20Withdrawal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForDifferentZRC20Withdrawal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForDifferentZRC20Withdrawal is a paid mutator transaction binding the contract method 0xb84a3d2f.
//
// Solidity: function testBurnGasFeeForDifferentZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnGasFeeForDifferentZRC20Withdrawal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForDifferentZRC20Withdrawal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x3177f381.
//
// Solidity: function testBurnGasFeeForZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnGasFeeForZRC20WithdrawAndCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnGasFeeForZRC20WithdrawAndCall")
}

// TestBurnGasFeeForZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x3177f381.
//
// Solidity: function testBurnGasFeeForZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnGasFeeForZRC20WithdrawAndCall() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForZRC20WithdrawAndCall(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x3177f381.
//
// Solidity: function testBurnGasFeeForZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnGasFeeForZRC20WithdrawAndCall() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForZRC20WithdrawAndCall(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForZRC20Withdrawal is a paid mutator transaction binding the contract method 0xf2a91358.
//
// Solidity: function testBurnGasFeeForZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnGasFeeForZRC20Withdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnGasFeeForZRC20Withdrawal")
}

// TestBurnGasFeeForZRC20Withdrawal is a paid mutator transaction binding the contract method 0xf2a91358.
//
// Solidity: function testBurnGasFeeForZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnGasFeeForZRC20Withdrawal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForZRC20Withdrawal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForZRC20Withdrawal is a paid mutator transaction binding the contract method 0xf2a91358.
//
// Solidity: function testBurnGasFeeForZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnGasFeeForZRC20Withdrawal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForZRC20Withdrawal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnProtocolFeesFailsWithInsufficientAllowance is a paid mutator transaction binding the contract method 0x61cfddb7.
//
// Solidity: function testBurnProtocolFeesFailsWithInsufficientAllowance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnProtocolFeesFailsWithInsufficientAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnProtocolFeesFailsWithInsufficientAllowance")
}

// TestBurnProtocolFeesFailsWithInsufficientAllowance is a paid mutator transaction binding the contract method 0x61cfddb7.
//
// Solidity: function testBurnProtocolFeesFailsWithInsufficientAllowance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnProtocolFeesFailsWithInsufficientAllowance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnProtocolFeesFailsWithInsufficientAllowance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnProtocolFeesFailsWithInsufficientAllowance is a paid mutator transaction binding the contract method 0x61cfddb7.
//
// Solidity: function testBurnProtocolFeesFailsWithInsufficientAllowance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnProtocolFeesFailsWithInsufficientAllowance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnProtocolFeesFailsWithInsufficientAllowance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDeposit is a paid mutator transaction binding the contract method 0x7f924c4e.
//
// Solidity: function testDeposit() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDeposit")
}

// TestDeposit is a paid mutator transaction binding the contract method 0x7f924c4e.
//
// Solidity: function testDeposit() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDeposit() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDeposit(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDeposit is a paid mutator transaction binding the contract method 0x7f924c4e.
//
// Solidity: function testDeposit() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDeposit() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDeposit(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x884660a3.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContract")
}

// TestDepositAndRevertZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x884660a3.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x884660a3.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway is a paid mutator transaction binding the contract method 0x51336fb0.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway is a paid mutator transaction binding the contract method 0x51336fb0.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway is a paid mutator transaction binding the contract method 0x51336fb0.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol is a paid mutator transaction binding the contract method 0x6163f8ef.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol is a paid mutator transaction binding the contract method 0x6163f8ef.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol is a paid mutator transaction binding the contract method 0x6163f8ef.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xeab7674e.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xeab7674e.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xeab7674e.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xf1d98f1b.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xf1d98f1b.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xf1d98f1b.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x339bd828.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x339bd828.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x339bd828.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x48f4fd07.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x48f4fd07.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x48f4fd07.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x996b7675.
//
// Solidity: function testDepositFailsIfAmountIs0() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfAmountIs0")
}

// TestDepositFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x996b7675.
//
// Solidity: function testDepositFailsIfAmountIs0() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfAmountIs0(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x996b7675.
//
// Solidity: function testDepositFailsIfAmountIs0() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfAmountIs0(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfSenderNotProtocol is a paid mutator transaction binding the contract method 0xca26929c.
//
// Solidity: function testDepositFailsIfSenderNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfSenderNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfSenderNotProtocol")
}

// TestDepositFailsIfSenderNotProtocol is a paid mutator transaction binding the contract method 0xca26929c.
//
// Solidity: function testDepositFailsIfSenderNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfSenderNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfSenderNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfSenderNotProtocol is a paid mutator transaction binding the contract method 0xca26929c.
//
// Solidity: function testDepositFailsIfSenderNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfSenderNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfSenderNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x96d9d876.
//
// Solidity: function testDepositFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfTargetIsGateway")
}

// TestDepositFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x96d9d876.
//
// Solidity: function testDepositFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x96d9d876.
//
// Solidity: function testDepositFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x3626c616.
//
// Solidity: function testDepositFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfTargetIsProtocol")
}

// TestDepositFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x3626c616.
//
// Solidity: function testDepositFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x3626c616.
//
// Solidity: function testDepositFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x14b7a6da.
//
// Solidity: function testDepositFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfTargetIsZeroAddress")
}

// TestDepositFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x14b7a6da.
//
// Solidity: function testDepositFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x14b7a6da.
//
// Solidity: function testDepositFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x9c9acd5d.
//
// Solidity: function testDepositFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfZRC20IsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfZRC20IsZeroAddress")
}

// TestDepositFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x9c9acd5d.
//
// Solidity: function testDepositFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x9c9acd5d.
//
// Solidity: function testDepositFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositTogglePause is a paid mutator transaction binding the contract method 0x1c785a14.
//
// Solidity: function testDepositTogglePause() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositTogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositTogglePause")
}

// TestDepositTogglePause is a paid mutator transaction binding the contract method 0x1c785a14.
//
// Solidity: function testDepositTogglePause() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositTogglePause() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositTogglePause(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositTogglePause is a paid mutator transaction binding the contract method 0x1c785a14.
//
// Solidity: function testDepositTogglePause() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositTogglePause() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositTogglePause(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETA is a paid mutator transaction binding the contract method 0x445bbf47.
//
// Solidity: function testDepositZETA() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETA")
}

// TestDepositZETA is a paid mutator transaction binding the contract method 0x445bbf47.
//
// Solidity: function testDepositZETA() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETA() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETA(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETA is a paid mutator transaction binding the contract method 0x445bbf47.
//
// Solidity: function testDepositZETA() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETA() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETA(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversal is a paid mutator transaction binding the contract method 0x828d267c.
//
// Solidity: function testDepositZETAAndCallUniversal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversal")
}

// TestDepositZETAAndCallUniversal is a paid mutator transaction binding the contract method 0x828d267c.
//
// Solidity: function testDepositZETAAndCallUniversal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversal is a paid mutator transaction binding the contract method 0x828d267c.
//
// Solidity: function testDepositZETAAndCallUniversal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContract is a paid mutator transaction binding the contract method 0x2468bc0f.
//
// Solidity: function testDepositZETAAndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContract")
}

// TestDepositZETAAndCallUniversalContract is a paid mutator transaction binding the contract method 0x2468bc0f.
//
// Solidity: function testDepositZETAAndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContract is a paid mutator transaction binding the contract method 0x2468bc0f.
//
// Solidity: function testDepositZETAAndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xc35cb5e4.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol")
}

// TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xc35cb5e4.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xc35cb5e4.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero is a paid mutator transaction binding the contract method 0xec294d9f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero")
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero is a paid mutator transaction binding the contract method 0xec294d9f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero is a paid mutator transaction binding the contract method 0xec294d9f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x5cec7db5.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway")
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x5cec7db5.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x5cec7db5.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x890a2d67.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol")
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x890a2d67.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x890a2d67.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x97f7661f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress")
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x97f7661f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x97f7661f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0xb936be8c.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfZeroAddress")
}

// TestDepositZETAAndCallUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0xb936be8c.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0xb936be8c.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContract is a paid mutator transaction binding the contract method 0x0c4a14b4.
//
// Solidity: function testDepositZETAAndRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContract")
}

// TestDepositZETAAndRevertUniversalContract is a paid mutator transaction binding the contract method 0x0c4a14b4.
//
// Solidity: function testDepositZETAAndRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContract is a paid mutator transaction binding the contract method 0x0c4a14b4.
//
// Solidity: function testDepositZETAAndRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xda8fd81d.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero")
}

// TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xda8fd81d.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xda8fd81d.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xe36c0a23.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol")
}

// TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xe36c0a23.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xe36c0a23.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x514fcc7f.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway")
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x514fcc7f.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x514fcc7f.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x50dd38ca.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol")
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x50dd38ca.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x50dd38ca.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x2ed7cff6.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress")
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x2ed7cff6.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x2ed7cff6.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractWhenPaused is a paid mutator transaction binding the contract method 0x62fe2915.
//
// Solidity: function testDepositZETAAndRevertUniversalContractWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractWhenPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractWhenPaused")
}

// TestDepositZETAAndRevertUniversalContractWhenPaused is a paid mutator transaction binding the contract method 0x62fe2915.
//
// Solidity: function testDepositZETAAndRevertUniversalContractWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractWhenPaused() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractWhenPaused(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractWhenPaused is a paid mutator transaction binding the contract method 0x62fe2915.
//
// Solidity: function testDepositZETAAndRevertUniversalContractWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractWhenPaused() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractWhenPaused(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x195ac143.
//
// Solidity: function testDepositZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfAmountIsZero")
}

// TestDepositZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x195ac143.
//
// Solidity: function testDepositZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x195ac143.
//
// Solidity: function testDepositZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfInsufficientAllowance is a paid mutator transaction binding the contract method 0x2fb780ba.
//
// Solidity: function testDepositZETAFailsIfInsufficientAllowance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfInsufficientAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfInsufficientAllowance")
}

// TestDepositZETAFailsIfInsufficientAllowance is a paid mutator transaction binding the contract method 0x2fb780ba.
//
// Solidity: function testDepositZETAFailsIfInsufficientAllowance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfInsufficientAllowance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfInsufficientAllowance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfInsufficientAllowance is a paid mutator transaction binding the contract method 0x2fb780ba.
//
// Solidity: function testDepositZETAFailsIfInsufficientAllowance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfInsufficientAllowance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfInsufficientAllowance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfInsufficientProtocolBalance is a paid mutator transaction binding the contract method 0xf589655d.
//
// Solidity: function testDepositZETAFailsIfInsufficientProtocolBalance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfInsufficientProtocolBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfInsufficientProtocolBalance")
}

// TestDepositZETAFailsIfInsufficientProtocolBalance is a paid mutator transaction binding the contract method 0xf589655d.
//
// Solidity: function testDepositZETAFailsIfInsufficientProtocolBalance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfInsufficientProtocolBalance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfInsufficientProtocolBalance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfInsufficientProtocolBalance is a paid mutator transaction binding the contract method 0xf589655d.
//
// Solidity: function testDepositZETAFailsIfInsufficientProtocolBalance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfInsufficientProtocolBalance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfInsufficientProtocolBalance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xae4662ec.
//
// Solidity: function testDepositZETAFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfSenderIsNotProtocol")
}

// TestDepositZETAFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xae4662ec.
//
// Solidity: function testDepositZETAFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xae4662ec.
//
// Solidity: function testDepositZETAFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0xb82ce921.
//
// Solidity: function testDepositZETAFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfTargetIsGateway")
}

// TestDepositZETAFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0xb82ce921.
//
// Solidity: function testDepositZETAFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0xb82ce921.
//
// Solidity: function testDepositZETAFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x21d92591.
//
// Solidity: function testDepositZETAFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfTargetIsProtocol")
}

// TestDepositZETAFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x21d92591.
//
// Solidity: function testDepositZETAFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x21d92591.
//
// Solidity: function testDepositZETAFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x1ff5bdd7.
//
// Solidity: function testDepositZETAFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfTargetIsZeroAddress")
}

// TestDepositZETAFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x1ff5bdd7.
//
// Solidity: function testDepositZETAFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x1ff5bdd7.
//
// Solidity: function testDepositZETAFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsWhenPaused is a paid mutator transaction binding the contract method 0x33f7c81a.
//
// Solidity: function testDepositZETAFailsWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsWhenPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsWhenPaused")
}

// TestDepositZETAFailsWhenPaused is a paid mutator transaction binding the contract method 0x33f7c81a.
//
// Solidity: function testDepositZETAFailsWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsWhenPaused() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsWhenPaused(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsWhenPaused is a paid mutator transaction binding the contract method 0x33f7c81a.
//
// Solidity: function testDepositZETAFailsWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsWhenPaused() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsWhenPaused(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x6efa04b5.
//
// Solidity: function testDepositZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContract")
}

// TestDepositZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x6efa04b5.
//
// Solidity: function testDepositZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x6efa04b5.
//
// Solidity: function testDepositZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x58c9987f.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero")
}

// TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x58c9987f.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x58c9987f.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xcf2c3d1d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol")
}

// TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xcf2c3d1d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xcf2c3d1d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xef2b5394.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress")
}

// TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xef2b5394.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xef2b5394.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0xfb339a1c.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress")
}

// TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0xfb339a1c.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0xfb339a1c.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsGateway is a paid mutator transaction binding the contract method 0xe09bc659.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractIfTargetIsGateway")
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsGateway is a paid mutator transaction binding the contract method 0xe09bc659.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsGateway is a paid mutator transaction binding the contract method 0xe09bc659.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol is a paid mutator transaction binding the contract method 0xeb78bd7d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractIfTargetIsProtocol")
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol is a paid mutator transaction binding the contract method 0xeb78bd7d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol is a paid mutator transaction binding the contract method 0xeb78bd7d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteAbortUniversalContract is a paid mutator transaction binding the contract method 0x1832cb6e.
//
// Solidity: function testExecuteAbortUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteAbortUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteAbortUniversalContract")
}

// TestExecuteAbortUniversalContract is a paid mutator transaction binding the contract method 0x1832cb6e.
//
// Solidity: function testExecuteAbortUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteAbortUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteAbortUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteAbortUniversalContract is a paid mutator transaction binding the contract method 0x1832cb6e.
//
// Solidity: function testExecuteAbortUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteAbortUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteAbortUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xee0f4ea1.
//
// Solidity: function testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress")
}

// TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xee0f4ea1.
//
// Solidity: function testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xee0f4ea1.
//
// Solidity: function testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x3ab5b199.
//
// Solidity: function testExecuteFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteFailsIfTargetIsZeroAddress")
}

// TestExecuteFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x3ab5b199.
//
// Solidity: function testExecuteFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x3ab5b199.
//
// Solidity: function testExecuteFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x2948df41.
//
// Solidity: function testExecuteFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteFailsIfZRC20IsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteFailsIfZRC20IsZeroAddress")
}

// TestExecuteFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x2948df41.
//
// Solidity: function testExecuteFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x2948df41.
//
// Solidity: function testExecuteFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContract is a paid mutator transaction binding the contract method 0x084fafab.
//
// Solidity: function testExecuteRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteRevertUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteRevertUniversalContract")
}

// TestExecuteRevertUniversalContract is a paid mutator transaction binding the contract method 0x084fafab.
//
// Solidity: function testExecuteRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteRevertUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContract is a paid mutator transaction binding the contract method 0x084fafab.
//
// Solidity: function testExecuteRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteRevertUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xc8814d2e.
//
// Solidity: function testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress")
}

// TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xc8814d2e.
//
// Solidity: function testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xc8814d2e.
//
// Solidity: function testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContractIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x27820625.
//
// Solidity: function testExecuteRevertUniversalContractIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteRevertUniversalContractIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteRevertUniversalContractIfSenderIsNotProtocol")
}

// TestExecuteRevertUniversalContractIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x27820625.
//
// Solidity: function testExecuteRevertUniversalContractIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteRevertUniversalContractIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContractIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContractIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x27820625.
//
// Solidity: function testExecuteRevertUniversalContractIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteRevertUniversalContractIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContractIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContract is a paid mutator transaction binding the contract method 0x7cec29b0.
//
// Solidity: function testExecuteUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteUniversalContract")
}

// TestExecuteUniversalContract is a paid mutator transaction binding the contract method 0x7cec29b0.
//
// Solidity: function testExecuteUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContract is a paid mutator transaction binding the contract method 0x7cec29b0.
//
// Solidity: function testExecuteUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x2acb21b4.
//
// Solidity: function testExecuteUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteUniversalContractFailsIfSenderIsNotProtocol")
}

// TestExecuteUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x2acb21b4.
//
// Solidity: function testExecuteUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x2acb21b4.
//
// Solidity: function testExecuteUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x5b4c90e1.
//
// Solidity: function testExecuteUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteUniversalContractFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteUniversalContractFailsIfZeroAddress")
}

// TestExecuteUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x5b4c90e1.
//
// Solidity: function testExecuteUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteUniversalContractFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContractFailsIfZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x5b4c90e1.
//
// Solidity: function testExecuteUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteUniversalContractFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContractFailsIfZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// GatewayZEVMOutboundTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestCalledIterator struct {
	Event *GatewayZEVMOutboundTestCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestCalled represents a Called event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestCalled struct {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, zrc20 []common.Address) (*GatewayZEVMOutboundTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestCalledIterator{contract: _GatewayZEVMOutboundTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestCalled, sender []common.Address, zrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestCalled)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "Called", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseCalled(log types.Log) (*GatewayZEVMOutboundTestCalled, error) {
	event := new(GatewayZEVMOutboundTestCalled)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestContextDataIterator is returned from FilterContextData and is used to iterate over the raw logs and unpacked data for ContextData events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataIterator struct {
	Event *GatewayZEVMOutboundTestContextData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestContextDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestContextData)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestContextData)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestContextDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestContextDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestContextData represents a ContextData event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextData struct {
	Origin    []byte
	Sender    common.Address
	ChainID   *big.Int
	MsgSender common.Address
	Message   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContextData is a free log retrieval operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterContextData(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestContextDataIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestContextDataIterator{contract: _GatewayZEVMOutboundTest.contract, event: "ContextData", logs: logs, sub: sub}, nil
}

// WatchContextData is a free log subscription operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchContextData(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestContextData) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestContextData)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextData", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContextData is a log parse operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseContextData(log types.Log) (*GatewayZEVMOutboundTestContextData, error) {
	event := new(GatewayZEVMOutboundTestContextData)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestContextDataAbortIterator is returned from FilterContextDataAbort and is used to iterate over the raw logs and unpacked data for ContextDataAbort events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataAbortIterator struct {
	Event *GatewayZEVMOutboundTestContextDataAbort // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestContextDataAbortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestContextDataAbort)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestContextDataAbort)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestContextDataAbortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestContextDataAbortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestContextDataAbort represents a ContextDataAbort event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataAbort struct {
	AbortContext AbortContext
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContextDataAbort is a free log retrieval operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterContextDataAbort(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestContextDataAbortIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "ContextDataAbort")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestContextDataAbortIterator{contract: _GatewayZEVMOutboundTest.contract, event: "ContextDataAbort", logs: logs, sub: sub}, nil
}

// WatchContextDataAbort is a free log subscription operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchContextDataAbort(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestContextDataAbort) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "ContextDataAbort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestContextDataAbort)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextDataAbort", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContextDataAbort is a log parse operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseContextDataAbort(log types.Log) (*GatewayZEVMOutboundTestContextDataAbort, error) {
	event := new(GatewayZEVMOutboundTestContextDataAbort)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextDataAbort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestContextDataRevertIterator is returned from FilterContextDataRevert and is used to iterate over the raw logs and unpacked data for ContextDataRevert events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataRevertIterator struct {
	Event *GatewayZEVMOutboundTestContextDataRevert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestContextDataRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestContextDataRevert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestContextDataRevert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestContextDataRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestContextDataRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestContextDataRevert represents a ContextDataRevert event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataRevert struct {
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContextDataRevert is a free log retrieval operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterContextDataRevert(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestContextDataRevertIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestContextDataRevertIterator{contract: _GatewayZEVMOutboundTest.contract, event: "ContextDataRevert", logs: logs, sub: sub}, nil
}

// WatchContextDataRevert is a free log subscription operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchContextDataRevert(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestContextDataRevert) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestContextDataRevert)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContextDataRevert is a log parse operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseContextDataRevert(log types.Log) (*GatewayZEVMOutboundTestContextDataRevert, error) {
	event := new(GatewayZEVMOutboundTestContextDataRevert)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestWithdrawnIterator struct {
	Event *GatewayZEVMOutboundTestWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestWithdrawn represents a Withdrawn event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestWithdrawn struct {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMOutboundTestWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestWithdrawnIterator{contract: _GatewayZEVMOutboundTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestWithdrawn)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseWithdrawn(log types.Log) (*GatewayZEVMOutboundTestWithdrawn, error) {
	event := new(GatewayZEVMOutboundTestWithdrawn)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestWithdrawnAndCalledIterator struct {
	Event *GatewayZEVMOutboundTestWithdrawnAndCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestWithdrawnAndCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestWithdrawnAndCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestWithdrawnAndCalled struct {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMOutboundTestWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestWithdrawnAndCalledIterator{contract: _GatewayZEVMOutboundTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestWithdrawnAndCalled)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayZEVMOutboundTestWithdrawnAndCalled, error) {
	event := new(GatewayZEVMOutboundTestWithdrawnAndCalled)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogIterator struct {
	Event *GatewayZEVMOutboundTestLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLog represents a Log event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLog)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLog(log types.Log) (*GatewayZEVMOutboundTestLog, error) {
	event := new(GatewayZEVMOutboundTestLog)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogAddressIterator struct {
	Event *GatewayZEVMOutboundTestLogAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogAddress represents a LogAddress event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogAddressIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogAddressIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogAddress)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogAddress(log types.Log) (*GatewayZEVMOutboundTestLogAddress, error) {
	event := new(GatewayZEVMOutboundTestLogAddress)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArrayIterator struct {
	Event *GatewayZEVMOutboundTestLogArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogArray represents a LogArray event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogArrayIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogArrayIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogArray)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogArray(log types.Log) (*GatewayZEVMOutboundTestLogArray, error) {
	event := new(GatewayZEVMOutboundTestLogArray)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray0Iterator struct {
	Event *GatewayZEVMOutboundTestLogArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogArray0 represents a LogArray0 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogArray0Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogArray0)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogArray0(log types.Log) (*GatewayZEVMOutboundTestLogArray0, error) {
	event := new(GatewayZEVMOutboundTestLogArray0)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray1Iterator struct {
	Event *GatewayZEVMOutboundTestLogArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogArray1 represents a LogArray1 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogArray1Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogArray1)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogArray1(log types.Log) (*GatewayZEVMOutboundTestLogArray1, error) {
	event := new(GatewayZEVMOutboundTestLogArray1)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogBytesIterator struct {
	Event *GatewayZEVMOutboundTestLogBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogBytes represents a LogBytes event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogBytesIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogBytesIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogBytes)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogBytes(log types.Log) (*GatewayZEVMOutboundTestLogBytes, error) {
	event := new(GatewayZEVMOutboundTestLogBytes)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogBytes32Iterator struct {
	Event *GatewayZEVMOutboundTestLogBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogBytes32 represents a LogBytes32 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogBytes32Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogBytes32)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogBytes32(log types.Log) (*GatewayZEVMOutboundTestLogBytes32, error) {
	event := new(GatewayZEVMOutboundTestLogBytes32)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogIntIterator struct {
	Event *GatewayZEVMOutboundTestLogInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogInt represents a LogInt event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogIntIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogIntIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogInt)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogInt(log types.Log) (*GatewayZEVMOutboundTestLogInt, error) {
	event := new(GatewayZEVMOutboundTestLogInt)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedAddressIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogNamedAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedAddressIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedAddress)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayZEVMOutboundTestLogNamedAddress, error) {
	event := new(GatewayZEVMOutboundTestLogNamedAddress)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArrayIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogNamedArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedArray represents a LogNamedArray event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedArrayIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedArray)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayZEVMOutboundTestLogNamedArray, error) {
	event := new(GatewayZEVMOutboundTestLogNamedArray)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray0Iterator struct {
	Event *GatewayZEVMOutboundTestLogNamedArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogNamedArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedArray0Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedArray0)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayZEVMOutboundTestLogNamedArray0, error) {
	event := new(GatewayZEVMOutboundTestLogNamedArray0)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray1Iterator struct {
	Event *GatewayZEVMOutboundTestLogNamedArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogNamedArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedArray1Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedArray1)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayZEVMOutboundTestLogNamedArray1, error) {
	event := new(GatewayZEVMOutboundTestLogNamedArray1)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedBytesIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogNamedBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedBytesIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedBytes)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayZEVMOutboundTestLogNamedBytes, error) {
	event := new(GatewayZEVMOutboundTestLogNamedBytes)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedBytes32Iterator struct {
	Event *GatewayZEVMOutboundTestLogNamedBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogNamedBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedBytes32Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedBytes32)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayZEVMOutboundTestLogNamedBytes32, error) {
	event := new(GatewayZEVMOutboundTestLogNamedBytes32)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedDecimalIntIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedDecimalInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedDecimalInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogNamedDecimalInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedDecimalIntIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedDecimalInt)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayZEVMOutboundTestLogNamedDecimalInt, error) {
	event := new(GatewayZEVMOutboundTestLogNamedDecimalInt)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedDecimalUintIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedDecimalUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedDecimalUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogNamedDecimalUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedDecimalUintIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedDecimalUint)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayZEVMOutboundTestLogNamedDecimalUint, error) {
	event := new(GatewayZEVMOutboundTestLogNamedDecimalUint)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedIntIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogNamedInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedInt represents a LogNamedInt event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedIntIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedInt)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayZEVMOutboundTestLogNamedInt, error) {
	event := new(GatewayZEVMOutboundTestLogNamedInt)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedStringIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogNamedString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedString represents a LogNamedString event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedStringIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedString)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedString(log types.Log) (*GatewayZEVMOutboundTestLogNamedString, error) {
	event := new(GatewayZEVMOutboundTestLogNamedString)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedUintIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogNamedUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedUint represents a LogNamedUint event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedUintIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedUint)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayZEVMOutboundTestLogNamedUint, error) {
	event := new(GatewayZEVMOutboundTestLogNamedUint)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogStringIterator struct {
	Event *GatewayZEVMOutboundTestLogString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogString represents a LogString event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogStringIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogStringIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogString)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogString(log types.Log) (*GatewayZEVMOutboundTestLogString, error) {
	event := new(GatewayZEVMOutboundTestLogString)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogUintIterator struct {
	Event *GatewayZEVMOutboundTestLogUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogUint represents a LogUint event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogUintIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogUintIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogUint)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogUint(log types.Log) (*GatewayZEVMOutboundTestLogUint, error) {
	event := new(GatewayZEVMOutboundTestLogUint)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogsIterator struct {
	Event *GatewayZEVMOutboundTestLogs // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayZEVMOutboundTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogs)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayZEVMOutboundTestLogs)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayZEVMOutboundTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogs represents a Logs event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogsIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogsIterator{contract: _GatewayZEVMOutboundTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogs)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "logs", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogs(log types.Log) (*GatewayZEVMOutboundTestLogs, error) {
	event := new(GatewayZEVMOutboundTestLogs)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
