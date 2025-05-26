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
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForDifferentZRC20WithdrawAndCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForDifferentZRC20Withdrawal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForZRC20WithdrawAndCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForZRC20Withdrawal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnProtocolFeesFailsWithInsufficientAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDeposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfSenderNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractWhenPaused\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteAbortUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContractIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContractFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextData\",\"inputs\":[{\"name\":\"origin\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"msgSender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataAbort\",\"inputs\":[{\"name\":\"abortContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAbortContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outgoing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataRevert\",\"inputs\":[{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b506201352b806200003e6000396000f3fe608060405234801561001057600080fd5b50600436106104065760003560e01c80637f924c4e1161021a578063c8814d2e11610135578063eab7674e116100c8578063ef2b539411610097578063f2a913581161007c578063f2a9135814610674578063fa7626d41461067c578063fb339a1c1461068957600080fd5b8063ef2b539414610664578063f1d98f1b1461066c57600080fd5b8063eab7674e14610644578063eb78bd7d1461064c578063ec294d9f14610654578063ee0f4ea11461065c57600080fd5b8063e09bc65911610104578063e09bc659146105f7578063e20c9f71146105ff578063e36c0a2314610607578063e63ab1e91461060f57600080fd5b8063c8814d2e146105d7578063ca26929c146105df578063cf2c3d1d146105e7578063da8fd81d146105ef57600080fd5b8063996b7675116101ad578063b84a3d2f1161017c578063b84a3d2f146105af578063b936be8c14610545578063ba414fa6146105b7578063c35cb5e4146105cf57600080fd5b8063996b76751461058f5780639c9acd5d14610597578063b0464fdc1461059f578063b5508aa9146105a757600080fd5b8063890a2d67116101e9578063890a2d671461056a578063916a17c61461057257806396d9d8761461058757806397f7661f1461054557600080fd5b80637f924c4e1461053d578063828d267c1461054557806385226c811461054d578063884660a31461056257600080fd5b80633ab5b199116103255780635b4c90e1116102b857806362fe2915116102875780636efa04b51161026c5780636efa04b51461052557806375ca55871461052d5780637cec29b01461053557600080fd5b806362fe29151461050857806366d9a9a01461051057600080fd5b80635b4c90e11461046b5780635cec7db5146104f05780636163f8ef146104f857806361cfddb71461050057600080fd5b806350dd38ca116102f457806350dd38ca146104d057806351336fb0146104d8578063514fcc7f146104e057806358c9987f146104e857600080fd5b80633ab5b199146104b05780633e5e3c23146104b85780633f7286f4146104c057806348f4fd07146104c857600080fd5b8063278206251161039d5780632ed7cff61161036c5780632ed7cff6146104905780633177f38114610498578063339bd828146104a05780633626c616146104a857600080fd5b806327820625146104635780632948df411461046b5780632acb21b4146104735780632ade38801461047b57600080fd5b80631832cb6e116103d95780631832cb6e1461042d5780631c785a14146104355780631ed7831c1461043d5780632468bc0f1461045b57600080fd5b8063084fafab1461040b5780630a9254e4146104155780630c4a14b41461041d57806314b7a6da14610425575b600080fd5b610413610691565b005b610413610867565b610413611488565b61041361187e565b6104136119ce565b610413611b6f565b610445612486565b604051610452919061d28a565b60405180910390f35b6104136124e8565b61041361299b565b610413612a63565b610413612c44565b610483612e03565b604051610452919061d326565b610413612f45565b61041361308a565b610413613776565b6104136138c2565b610413613a4a565b610445613bfb565b610445613c5b565b610413613cbb565b610413613e06565b610413613f4f565b61041361409a565b6104136141e0565b610413614395565b610413614583565b6104136146d1565b610413614bcc565b61051861509b565b604051610452919061d48c565b61041361521d565b6104136155c1565b610413615dff565b610413615fea565b610413616023565b6105556161cf565b604051610452919061d52a565b61041361629f565b61041361657e565b61057a616739565b604051610452919061d5a1565b610413616834565b610413616987565b610413616adb565b61057a616c2b565b610555616d26565b610413616df6565b6105bf617547565b6040519015158152602001610452565b61041361761b565b6104136177d6565b610413617918565b610413617b9b565b610413617d59565b610413617ea0565b610445618059565b6104136180b9565b6106367f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b604051908152602001610452565b610413618184565b6104136182d2565b610413618490565b61041361863f565b610413618781565b610413618932565b610413618a80565b601f546105bf9060ff1681565b610413618df7565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561072057600080fd5b505af1158015610734573d6000803e3d6000fd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602c604051610768919061d769565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b1580156107ca57600080fd5b505af11580156107de573d6000803e3d6000fd5b50506020546024546040517f184b07930000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063184b079393506108339290911690602c9060040161d77c565b600060405180830381600087803b15801561084d57600080fd5b505af1158015610861573d6000803e3d6000fd5b50505050565b602580547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602680549091166112341790556040516108ad9061d199565b604051809103906000f0801580156108c9573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600f81527f476174657761795a45564d2e736f6c0000000000000000000000000000000000602082015260255491516024810193909352921660448201526109ad919060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f485cc95500000000000000000000000000000000000000000000000000000000179052618faa565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b039384168102919091179182905560208054919092049092167fffffffffffffffffffffffff000000000000000000000000000000000000000090921682178155604080517f2722feee0000000000000000000000000000000000000000000000000000000081529051632722feee926004808401939192918290030181865afa158015610a6f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a93919061d7ba565b602780547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055604051610ad79061d1a7565b604051809103906000f080158015610af3573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161781556027546040517f06447d5600000000000000000000000000000000000000000000000000000000815292166004830152737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d569101600060405180830381600087803b158015610b8f57600080fd5b505af1158015610ba3573d6000803e3d6000fd5b505050506000806000604051610bb89061d1b5565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610bf4573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602054604051601293600193849360009391921690610c4a9061d1c3565b610c599695949392919061d810565b604051809103906000f080158015610c75573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179091556023546040517fee2815ba0000000000000000000000000000000000000000000000000000000081526001600482015260248101929092529091169063ee2815ba90604401600060405180830381600087803b158015610d0c57600080fd5b505af1158015610d20573d6000803e3d6000fd5b50506023546040517fa7cb050700000000000000000000000000000000000000000000000000000000815260016004820181905260248201526001600160a01b03909116925063a7cb05079150604401600060405180830381600087803b158015610d8a57600080fd5b505af1158015610d9e573d6000803e3d6000fd5b50506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152633b9aca006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b158015610e1e57600080fd5b505af1158015610e32573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b158015610e8757600080fd5b505af1158015610e9b573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af1158015610f0f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f33919061d8ca565b506021546025546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a060248201529116906347e7ef24906044016020604051808303816000875af1158015610fa4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fc8919061d8ca565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561102757600080fd5b505af115801561103b573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b1580156110b157600080fd5b505af11580156110c5573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a060248201529116925063095ea7b391506044016020604051808303816000875af1158015611139573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115d919061d8ca565b50602260009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b1580156111af57600080fd5b505af11580156111c3573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af1158015611237573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061125b919061d8ca565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156112ba57600080fd5b505af11580156112ce573d6000803e3d6000fd5b5050604080516080810182526025546001600160a01b0390811682526000602080840182815260018587019081528651928301909652918152606084018190528351602c80549185167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161781559251602d8054919095169116179092559251602e55909350909150602f90611366908261d962565b50506040805160c0810190915260255460601b6bffffffffffffffffffffffff191660e082015290508060f4810160408051601f198184030181529181529082526000602083810182905260018484018190526060850183905260808501528251908101909252815260a090910152805160309081906113e6908261d962565b5060208201516001820180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039092169190911790556040820151600282015560608201516003820180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790556080820151600482015560a08201516005820190611483908261d962565b505050565b6022546027546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa1580156114d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114fd919061da21565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa15801561154f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611573919061da21565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b031660848201819052919250903190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561160a57600080fd5b505af115801561161e573d6000803e3d6000fd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602c604051611652919061d769565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156116b357600080fd5b505af11580156116c7573d6000803e3d6000fd5b50506020546024546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506345122cfd935061171c9289921690602c9060040161da3a565b600060405180830381600087803b15801561173657600080fd5b505af115801561174a573d6000803e3d6000fd5b50506022546027546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561179d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117c1919061da21565b90506117d66117d0868661da91565b82618fc9565b6022546020546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611827573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061184b919061da21565b90506118578482618fc9565b611876611864878561daa4565b6024546001600160a01b031631618fc9565b505050505050565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156118d757600080fd5b505af11580156118eb573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561195b57600080fd5b505af115801561196f573d6000803e3d6000fd5b50506020546021546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260016024820152600060448201529116925063f45346dc9150606401610833565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611a5d57600080fd5b505af1158015611a71573d6000803e3d6000fd5b505050507f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db76030604051611aa5919061db15565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611b0657600080fd5b505af1158015611b1a573d6000803e3d6000fd5b50506020546024546040517f2095dedb0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450632095dedb9350610833929091169060309060040161db28565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611bc857600080fd5b505af1158015611bdc573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611cc7919060040161db4a565b600060405180830381600087803b158015611ce157600080fd5b505af1158015611cf5573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611d4957600080fd5b505af1158015611d5d573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015611dba57600080fd5b505af1158015611dce573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611eb9919060040161db4a565b600060405180830381600087803b158015611ed357600080fd5b505af1158015611ee7573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611f3b57600080fd5b505af1158015611f4f573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015611fac57600080fd5b505af1158015611fc0573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561201457600080fd5b505af1158015612028573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd93c066500000000000000000000000000000000000000000000000000000000600482015260019250737109709ecfa91a80626ff3989d68f67f5b1dd12d915063c31eb0e090602401600060405180830381600087803b15801561209b57600080fd5b505af11580156120af573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561210c57600080fd5b505af1158015612120573d6000803e3d6000fd5b50506020546021546026546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810187905290821660448201529116925063f45346dc9150606401600060405180830381600087803b15801561219857600080fd5b505af11580156121ac573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561220957600080fd5b505af115801561221d573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561227157600080fd5b505af1158015612285573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a08231906024015b602060405180830381865afa1580156122d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122fd919061da21565b905061230a600082618fc9565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561236357600080fd5b505af1158015612377573d6000803e3d6000fd5b50506020546021546026546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810188905290821660448201529116925063f45346dc9150606401600060405180830381600087803b1580156123ef57600080fd5b505af1158015612403573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015612456573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061247a919061da21565b90506114838382618fc9565b606060168054806020026020016040519081016040528092919081815260200182805480156124de57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116124c0575b5050505050905090565b6022546027546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015612539573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061255d919061da21565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa1580156125af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125d3919061da21565b6024546040519192506001600160a01b031631906000906125f69060200161db5d565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff191660808301529150600090806094810160408051808303601f190181529181529082526027546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156126e457600080fd5b505af11580156126f8573d6000803e3d6000fd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e945061275393506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602754602054612783936001600160a01b03928316928c92169061db9a565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156127e457600080fd5b505af11580156127f8573d6000803e3d6000fd5b50506020546024546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321501a95935061284f9286928c92911690889060040161dc4b565b600060405180830381600087803b15801561286957600080fd5b505af115801561287d573d6000803e3d6000fd5b50506022546027546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156128d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128f4919061da21565b90506129036117d0888861da91565b6022546020546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015612954573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612978919061da21565b90506129848682618fc9565b612991611864898761daa4565b5050505050505050565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612a0757600080fd5b505af1158015612a1b573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024016107b0565b6000604051602001612a749061db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015612b2257600080fd5b505af1158015612b36573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015612ba657600080fd5b505af1158015612bba573d6000803e3d6000fd5b50506020546024546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063bcf7f32b9350612c1692869260009260019290911690899060040161dc85565b600060405180830381600087803b158015612c3057600080fd5b505af1158015611876573d6000803e3d6000fd5b6000604051602001612c559061db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b158015612d1f57600080fd5b505af1158015612d33573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015b600060405180830381600087803b158015612d9157600080fd5b505af1158015612da5573d6000803e3d6000fd5b50506020546021546024546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063bcf7f32b9450612c16938793811692600192911690899060040161dc85565b6060601e805480602002602001604051908101604052809291908181526020016000905b82821015612f3c57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015612f25578382906000526020600020018054612e989061d638565b80601f0160208091040260200160405190810160405280929190818152602001828054612ec49061d638565b8015612f115780601f10612ee657610100808354040283529160200191612f11565b820191906000526020600020905b815481529060010190602001808311612ef457829003601f168201915b505050505081526020019060010190612e79565b505050508152505081526020019060010190612e27565b50505050905090565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612f9e57600080fd5b505af1158015612fb2573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561302257600080fd5b505af1158015613036573d6000803e3d6000fd5b50506020546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506345122cfd915061083390600190600090602c9060040161da3a565b6026546040516001600160a01b03909116602482015260019060009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae7600000000000000000000000000000000000000000000000000000000017905260275490517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561317557600080fd5b505af1158015613189573d6000803e3d6000fd5b50506021546040517ff687d12a000000000000000000000000000000000000000000000000000000008152620186a060048201526001600160a01b03909116925063f687d12a9150602401600060405180830381600087803b1580156131ee57600080fd5b505af1158015613202573d6000803e3d6000fd5b50506021546040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152620186a06004820152600093506001600160a01b03909116915063fc5fecd5906024016040805180830381865afa15801561326c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613290919061dcda565b6021546025549193506001600160a01b0390811692506347e7ef2491166132b7848761daa4565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af115801561331a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061333e919061d8ca565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561339d57600080fd5b505af11580156133b1573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561342757600080fd5b505af115801561343b573d6000803e3d6000fd5b50506021546020546001600160a01b03918216935063095ea7b3925016613462848761daa4565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af11580156134c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134e9919061d8ca565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561354857600080fd5b505af115801561355c573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156135b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135d9919061da21565b602080546026546040519394506001600160a01b0391821693637b15118b9361361f9392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f198184030181526021548383018352620186a084526001602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815261368e9391928a926001600160a01b03909116918a919060289060040161dd63565b600060405180830381600087803b1580156136a857600080fd5b505af11580156136bc573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613715573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613739919061da21565b90506000613747848761daa4565b9050611876613756828561da91565b836040518060600160405280603a8152602001620134bc603a9139619049565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156137cf57600080fd5b505af11580156137e3573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561385357600080fd5b505af1158015613867573d6000803e3d6000fd5b50506020546021546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450639d4ba46593506108339290911690600190600090602c9060040161ddd1565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561391e57600080fd5b505af1158015613932573d6000803e3d6000fd5b5050604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156139a257600080fd5b505af11580156139b6573d6000803e3d6000fd5b50506020546021546027546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810187905290821660448201529116925063f45346dc91506064015b600060405180830381600087803b158015613a2f57600080fd5b505af1158015613a43573d6000803e3d6000fd5b5050505050565b6000604051602001613a5b9061db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015613b0957600080fd5b505af1158015613b1d573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015613b8d57600080fd5b505af1158015613ba1573d6000803e3d6000fd5b50506020546021546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063bcf7f32b9350612c169286921690600190600090899060040161dc85565b606060188054806020026020016040519081016040528092919081815260200182805480156124de576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116124c0575050505050905090565b606060178054806020026020016040519081016040528092919081815260200182805480156124de576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116124c0575050505050905090565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613d1457600080fd5b505af1158015613d28573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015613d9857600080fd5b505af1158015613dac573d6000803e3d6000fd5b50506020546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450639d4ba465935061083392600092600192911690602c9060040161ddd1565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015613e7557600080fd5b505af1158015613e89573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015613ee657600080fd5b505af1158015613efa573d6000803e3d6000fd5b50506020546027546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506345122cfd9350613a159286921690602c9060040161da3a565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015613fbb57600080fd5b505af1158015613fcf573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561402c57600080fd5b505af1158015614040573d6000803e3d6000fd5b50506020546021546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450639d4ba465935061083392909116906001908590602c9060040161ddd1565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561410957600080fd5b505af115801561411d573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561417a57600080fd5b505af115801561418e573d6000803e3d6000fd5b50506020546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506345122cfd9150613a159084908490602c9060040161da3a565b60006040516020016141f19061db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b15801561429f57600080fd5b505af11580156142b3573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561432357600080fd5b505af1158015614337573d6000803e3d6000fd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450612c16938793811692600092911690899060040161dc85565b6040516001906000906143aa9060200161db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b15801561447457600080fd5b505af1158015614488573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156144e557600080fd5b505af11580156144f9573d6000803e3d6000fd5b50506020546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506321501a95915061454c90849087908590889060040161dc4b565b600060405180830381600087803b15801561456657600080fd5b505af115801561457a573d6000803e3d6000fd5b50505050505050565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156145ef57600080fd5b505af1158015614603573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561466057600080fd5b505af1158015614674573d6000803e3d6000fd5b50506020546021546027546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061083393928316926001921690602c9060040161ddd1565b6025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561474657600080fd5b505af115801561475a573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018690529116925063095ea7b391506044016020604051808303816000875af11580156147cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906147f0919061d8ca565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561484f57600080fd5b505af1158015614863573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156148c057600080fd5b505af11580156148d4573d6000803e3d6000fd5b50506021546040517ff687d12a00000000000000000000000000000000000000000000000000000000815262030d4060048201526001600160a01b03909116925063f687d12a9150602401600060405180830381600087803b15801561493957600080fd5b505af115801561494d573d6000803e3d6000fd5b50506021546040517ffc5fecd500000000000000000000000000000000000000000000000000000000815262030d406004820152600093508392506001600160a01b039091169063fc5fecd5906024016040805180830381865afa1580156149b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906149dd919061dcda565b60208054604080516001600160a01b03868116602483015290921660448301526064808301859052815180840390910181526084909201815291810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f667011120000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152929450909250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391614ab09160040161db4a565b600060405180830381600087803b158015614aca57600080fd5b505af1158015614ade573d6000803e3d6000fd5b5050602080546026546040805160609290921b6bffffffffffffffffffffffff19168285015280518083036014018152602154600460348501818152605886018552605490950180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f19ff1d21000000000000000000000000000000000000000000000000000000001790528351808501855262030d40815260019781019790975292517f06cb89830000000000000000000000000000000000000000000000000000000081526001600160a01b0395861698506306cb8983975061454c96929590911693926028910161de08565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614c2857600080fd5b505af1158015614c3c573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b158015614c9057600080fd5b505af1158015614ca4573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd93c0665000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015614d1457600080fd5b505af1158015614d28573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015614d8557600080fd5b505af1158015614d99573d6000803e3d6000fd5b50506020546024546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506345122cfd9350614dee9286921690602c9060040161da3a565b600060405180830381600087803b158015614e0857600080fd5b505af1158015614e1c573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015614e7957600080fd5b505af1158015614e8d573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015614ee157600080fd5b505af1158015614ef5573d6000803e3d6000fd5b5050602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015614f8857600080fd5b505af1158015614f9c573d6000803e3d6000fd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602c604051614fd0919061d769565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b15801561503257600080fd5b505af1158015615046573d6000803e3d6000fd5b50506020546024546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506345122cfd9350613a159286921690602c9060040161da3a565b6060601b805480602002602001604051908101604052809291908181526020016000905b82821015612f3c57838290600052602060002090600202016040518060400160405290816000820180546150f29061d638565b80601f016020809104026020016040519081016040528092919081815260200182805461511e9061d638565b801561516b5780601f106151405761010080835404028352916020019161516b565b820191906000526020600020905b81548152906001019060200180831161514e57829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561520557602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116151b25790505b505050505081525050815260200190600101906150bf565b602154602480546040516370a0823160e01b81526001600160a01b03918216600482015260009391909116916370a082319101602060405180830381865afa15801561526d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615291919061da21565b905061529e600082618fc9565b60006040516020016152af9061db5d565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff191660808301529150600090806094810160408051808303601f190181529181529082526027546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561539d57600080fd5b505af11580156153b1573d6000803e3d6000fd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e945061540c93506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f198184030181529082905260275460205461543d936001600160a01b0392831692600192169061db9a565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561549e57600080fd5b505af11580156154b2573d6000803e3d6000fd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450615510938793811692600192911690899060040161dc85565b600060405180830381600087803b15801561552a57600080fd5b505af115801561553e573d6000803e3d6000fd5b5050602154602480546040516370a0823160e01b81526001600160a01b03918216600482015260009550921692506370a082319101602060405180830381865afa158015615590573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906155b4919061da21565b9050610861600182618fc9565b6027546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561563357600080fd5b505af1158015615647573d6000803e3d6000fd5b50506023546020546040516000945060129350600192600292620186a0926001600160a01b03928316929091169061567e9061d1c3565b61568d9695949392919061de63565b604051809103906000f0801580156156a9573d6000803e3d6000fd5b506025546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152606460248201529192508216906347e7ef24906044016020604051808303816000875af1158015615718573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061573c919061d8ca565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561579b57600080fd5b505af11580156157af573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561582557600080fd5b505af1158015615839573d6000803e3d6000fd5b50506020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260646024820152908416925063095ea7b391506044016020604051808303816000875af11580156158a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906158cd919061d8ca565b506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482015291169063095ea7b3906044016020604051808303816000875af115801561593e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615962919061d8ca565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156159c157600080fd5b505af11580156159d5573d6000803e3d6000fd5b50506026546040516001600160a01b039091166024820152600192506000915060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae76000000000000000000000000000000000000000000000000000000000179052517ffc5fecd5000000000000000000000000000000000000000000000000000000008152620186a0600482015290915060009081906001600160a01b0386169063fc5fecd5906024016040805180830381865afa158015615ab5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615ad9919061dcda565b6021549193509150615af59083906001600160a01b03166190c9565b602154604080517f18160ddd00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916318160ddd9160048083019260209291908290030181865afa158015615b58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615b7c919061da21565b90506000866001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615bbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615be2919061da21565b602080546026546040519394506001600160a01b0391821693637b15118b93615c289392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f198184030181528282018252620186a083526001602084015290517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152615c87928b918d918c9160289060040161dd63565b600060405180830381600087803b158015615ca157600080fd5b505af1158015615cb5573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615d0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615d32919061da21565b90506000886001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615d74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615d98919061da21565b9050615dc7615da7868661da91565b836040518060600160405280602b815260200162013433602b9139619049565b615df4615dd4898561da91565b826040518060600160405280603881526020016201348460389139619049565b505050505050505050565b6000604051602001615e109061db5d565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff191660808301529150600090806094810160408051808303601f190181529181529082526027546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015615efe57600080fd5b505af1158015615f12573d6000803e3d6000fd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e9450615f6d93506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602754602054615f9e936001600160a01b0392831692600192169061db9a565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401612d77565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a08231906024016122bc565b60006040516020016160349061db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b1580156160e257600080fd5b505af11580156160f6573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561616657600080fd5b505af115801561617a573d6000803e3d6000fd5b50506020546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506321501a959150612c16908490600190600090889060040161dc4b565b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015612f3c5783829060005260206000200180546162129061d638565b80601f016020809104026020016040519081016040528092919081815260200182805461623e9061d638565b801561628b5780601f106162605761010080835404028352916020019161628b565b820191906000526020600020905b81548152906001019060200180831161626e57829003601f168201915b5050505050815260200190600101906161f3565b602154602480546040516370a0823160e01b81526001600160a01b03918216600482015260009391909116916370a082319101602060405180830381865afa1580156162ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616313919061da21565b9050616320600082618fc9565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156163af57600080fd5b505af11580156163c3573d6000803e3d6000fd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602c6040516163f7919061d769565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561645857600080fd5b505af115801561646c573d6000803e3d6000fd5b50506020546021546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba46594506164c993928316926001921690602c9060040161ddd1565b600060405180830381600087803b1580156164e357600080fd5b505af11580156164f7573d6000803e3d6000fd5b5050602154602480546040516370a0823160e01b81526001600160a01b03918216600482015260009550921692506370a082319101602060405180830381865afa158015616549573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061656d919061da21565b905061657a600182618fc9565b5050565b6040516001906000906165939060200161db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b15801561665d57600080fd5b505af1158015616671573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156166ce57600080fd5b505af11580156166e2573d6000803e3d6000fd5b50506020546027546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321501a95935061454c9286928992911690889060040161dc4b565b6060601d805480602002602001604051908101604052809291908181526020016000905b82821015612f3c5760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561681c57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116167c95790505b5050505050815250508152602001906001019061675d565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561689057600080fd5b505af11580156168a4573d6000803e3d6000fd5b5050604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561691457600080fd5b505af1158015616928573d6000803e3d6000fd5b50506020546021546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101869052911660448201819052925063f45346dc9150606401613a15565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156169e057600080fd5b505af11580156169f4573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015616a6457600080fd5b505af1158015616a78573d6000803e3d6000fd5b50506020546021546026546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526000602482015290821660448201529116925063f45346dc9150606401610833565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015616b3457600080fd5b505af1158015616b48573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015616bb857600080fd5b505af1158015616bcc573d6000803e3d6000fd5b50506020546026546040517ff45346dc00000000000000000000000000000000000000000000000000000000815260006004820152600160248201526001600160a01b0391821660448201529116925063f45346dc9150606401610833565b6060601c805480602002602001604051908101604052809291908181526020016000905b82821015612f3c5760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015616d0e57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411616cbb5790505b50505050508152505081526020019060010190616c4f565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015612f3c578382906000526020600020018054616d699061d638565b80601f0160208091040260200160405190810160405280929190818152602001828054616d959061d638565b8015616de25780601f10616db757610100808354040283529160200191616de2565b820191906000526020600020905b815481529060010190602001808311616dc557829003601f168201915b505050505081526020019060010190616d4a565b6027546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015616e6857600080fd5b505af1158015616e7c573d6000803e3d6000fd5b50506023546020546040516000945060129350600192600292620186a0926001600160a01b039283169290911690616eb39061d1c3565b616ec29695949392919061de63565b604051809103906000f080158015616ede573d6000803e3d6000fd5b506025546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152606460248201529192508216906347e7ef24906044016020604051808303816000875af1158015616f4d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616f71919061d8ca565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015616fd057600080fd5b505af1158015616fe4573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561705a57600080fd5b505af115801561706e573d6000803e3d6000fd5b50506020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260646024820152908416925063095ea7b391506044016020604051808303816000875af11580156170de573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617102919061d8ca565b506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482015291169063095ea7b3906044016020604051808303816000875af1158015617173573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617197919061d8ca565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156171f657600080fd5b505af115801561720a573d6000803e3d6000fd5b50506040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152620186a06004820152600192506000915081906001600160a01b0385169063fc5fecd5906024016040805180830381865afa158015617275573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617299919061dcda565b60215491935091506172b59083906001600160a01b03166190c9565b602154604080517f18160ddd00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916318160ddd9160048083019260209291908290030181865afa158015617318573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061733c919061da21565b90506000856001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561737e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906173a2919061da21565b602080546026546040519394506001600160a01b0391821693637c0dcb5f936173e89392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b604051602081830303815290604052878960286040518563ffffffff1660e01b815260040161741a949392919061def3565b600060405180830381600087803b15801561743457600080fd5b505af1158015617448573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156174a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906174c5919061da21565b90506000876001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015617507573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061752b919061da21565b905061753a615da7868661da91565b612991615dd4888561da91565b60085460009060ff161561755f575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa1580156175f0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617614919061da21565b1415905090565b6040516001906000906176309060200161db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b1580156176fa57600080fd5b505af115801561770e573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561776b57600080fd5b505af115801561777f573d6000803e3d6000fd5b50506020546024546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321501a95935061454c9286928992911690889060040161dc4b565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561782f57600080fd5b505af1158015617843573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156178b357600080fd5b505af11580156178c7573d6000803e3d6000fd5b50506020546040517f184b07930000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063184b0793915061083390600090602c9060040161d77c565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015617969573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061798d919061da21565b905061799a600082618fc9565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156179f357600080fd5b505af1158015617a07573d6000803e3d6000fd5b5050604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015617a7757600080fd5b505af1158015617a8b573d6000803e3d6000fd5b50506020546021546026546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810188905290821660448201529116925063f45346dc9150606401600060405180830381600087803b158015617b0357600080fd5b505af1158015617b17573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015617b6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617b8e919061da21565b9050611483600082618fc9565b6000604051602001617bac9061db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b158015617c7657600080fd5b505af1158015617c8a573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015617ce757600080fd5b505af1158015617cfb573d6000803e3d6000fd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450612c16938793811692600192911690899060040161dc85565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015617db257600080fd5b505af1158015617dc6573d6000803e3d6000fd5b5050604051630618f58760e51b81527f19c08f49000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015617e3657600080fd5b505af1158015617e4a573d6000803e3d6000fd5b50506020546024546040517f45122cfd0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506345122cfd9350610833926000921690602c9060040161da3a565b6000604051602001617eb19061db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b158015617f7b57600080fd5b505af1158015617f8f573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015617fec57600080fd5b505af1158015618000573d6000803e3d6000fd5b50506020546021546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca379350612c1692869216906001908690899060040161dc85565b606060158054806020026020016040519081016040528092919081815260200182805480156124de576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116124c0575050505050905090565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561812857600080fd5b505af115801561813c573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401615018565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156181dd57600080fd5b505af11580156181f1573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561826157600080fd5b505af1158015618275573d6000803e3d6000fd5b50506020546021546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061083393928316926000921690602c9060040161ddd1565b60006040516020016182e39061db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b1580156183ad57600080fd5b505af11580156183c1573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561841e57600080fd5b505af1158015618432573d6000803e3d6000fd5b50506020546021546027546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450612c16938793811692600192911690899060040161dc85565b60006040516020016184a19061db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b15801561854f57600080fd5b505af1158015618563573d6000803e3d6000fd5b5050604051630618f58760e51b81527f19c08f49000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156185d357600080fd5b505af11580156185e7573d6000803e3d6000fd5b50506020546021546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321501a959350612c16928692600092911690889060040161dc4b565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561869857600080fd5b505af11580156186ac573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561871c57600080fd5b505af1158015618730573d6000803e3d6000fd5b50506020546040517f2095dedb0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632095dedb91506108339060009060309060040161db28565b60006040516020016187929061db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b15801561884057600080fd5b505af1158015618854573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156188c457600080fd5b505af11580156188d8573d6000803e3d6000fd5b50506020546021546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca379350612c169286921690600190600090899060040161dc85565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561899e57600080fd5b505af11580156189b2573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015618a0f57600080fd5b505af1158015618a23573d6000803e3d6000fd5b50506020546021546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061083393928316926001921690602c9060040161ddd1565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015618adc57600080fd5b505af1158015618af0573d6000803e3d6000fd5b50506021546040517ff687d12a00000000000000000000000000000000000000000000000000000000815261c35060048201526001600160a01b03909116925063f687d12a9150602401600060405180830381600087803b158015618b5457600080fd5b505af1158015618b68573d6000803e3d6000fd5b50506021546040517ffc5fecd500000000000000000000000000000000000000000000000000000000815261c3506004820152600093506001600160a01b03909116915063fc5fecd5906024016040805180830381865afa158015618bd1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618bf5919061dcda565b9150506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015618c4d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618c71919061da21565b602080546026546040519394506001600160a01b0391821693637c0dcb5f93618cb79392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352618d0f9288916001600160a01b03169060289060040161def3565b600060405180830381600087803b158015618d2957600080fd5b505af1158015618d3d573d6000803e3d6000fd5b505050506000602160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015618d96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190618dba919061da21565b90506000618dc8848661daa4565b9050613a43618dd7828561da91565b836040518060600160405280602681526020016201345e60269139619049565b6000604051602001618e089061db5d565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015618eb657600080fd5b505af1158015618eca573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015618f3a57600080fd5b505af1158015618f4e573d6000803e3d6000fd5b50506020546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca379350612c1692869260009260019290911690899060040161dc85565b6000618fb461d1d1565b618fbf84848361912a565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015b60006040518083038186803b15801561903557600080fd5b505afa158015611876573d6000803e3d6000fd5b6040517f88b44c85000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d906388b44c859061909d9086908690869060040161df2d565b60006040518083038186803b1580156190b557600080fd5b505afa15801561457a573d6000803e3d6000fd5b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f69060440161901d565b60008061913785846191a5565b905061919a6040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f7879000000815250828660405160200161918592919061df4c565b604051602081830303815290604052856191b1565b9150505b9392505050565b600061919e83836191df565b60c081015151600090156191d5576191ce84848460c001516191fa565b905061919e565b6191ce84846193a0565b60006191eb838361948b565b61919e838360200151846191b1565b600080619205619497565b90506000619213868361956a565b9050600061922a8260600151836020015185619a10565b9050600061923a83838989619c22565b905060006192478261aa9f565b602081015181519192509060030b156192ba5789826040015160405160200161927192919061df6e565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526192b19160040161db4a565b60405180910390fd5b60006192fd6040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a20000000000000000000000081525083600161ac6e565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d9061935090849060040161db4a565b602060405180830381865afa15801561936d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190619391919061d7ba565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc925906193f590879060040161db4a565b600060405180830381865afa158015619412573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261943a919081019061e0a8565b90506000619468828560405160200161945492919061e0dd565b60405160208183030381529060405261ae6e565b90506001600160a01b038116618fbf57848460405160200161927192919061e10c565b61657a8282600061ae81565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c9061951e90849060040161e1b7565b600060405180830381865afa15801561953b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619563919081019061e1fe565b9250505090565b61959c6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d90506195e76040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6195f08561af84565b602082015260006196008661b369565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa158015619642573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261966a919081019061e1fe565b86838560200151604051602001619684949392919061e247565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb11906196dc90859060040161db4a565b600060405180830381865afa1580156196f9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619721919081019061e1fe565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f69061976990849060040161e34b565b602060405180830381865afa158015619786573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906197aa919061d8ca565b6197bf5781604051602001619271919061e39d565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061980490849060040161e42f565b600060405180830381865afa158015619821573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619849919081019061e1fe565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f69061989090849060040161e481565b602060405180830381865afa1580156198ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906198d1919061d8ca565b15619966576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061991b90849060040161e481565b600060405180830381865afa158015619938573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052619960919081019061e1fe565b60408501525b846001600160a01b03166349c4fac882866000015160405160200161998b919061e4d3565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016199b792919061e53f565b600060405180830381865afa1580156199d4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526199fc919081019061e1fe565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b6060815260200190600190039081619a2c5790505090506040518060400160405280600481526020017f677265700000000000000000000000000000000000000000000000000000000081525081600081518110619a8c57619a8c61e564565b60200260200101819052506040518060400160405280600381526020017f2d726c000000000000000000000000000000000000000000000000000000000081525081600181518110619ae057619ae061e564565b602002602001018190525084604051602001619afc919061e593565b60405160208183030381529060405281600281518110619b1e57619b1e61e564565b602002602001018190525082604051602001619b3a919061e5ff565b60405160208183030381529060405281600381518110619b5c57619b5c61e564565b60200260200101819052506000619b728261aa9f565b602080820151604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000008185019081528251808401845260008082529086015282518084019093529051825292810192909252919250619c03906040805180820182526000808252602091820152815180830190925284518252808501908201529061b5ec565b619c185785604051602001619271919061e640565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015619c72565b511590565b619de657826020015115619d2e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a4016192b1565b8260c0015115619de6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a4016192b1565b6040805160ff8082526120008201909252600091816020015b6060815260200190600190039081619dff57905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280619e5a9061e6d1565b935060ff1681518110619e6f57619e6f61e564565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e3700000000000000000000000000000000000000815250604051602001619ec0919061e6f0565b604051602081830303815290604052828280619edb9061e6d1565b935060ff1681518110619ef057619ef061e564565b60200260200101819052506040518060400160405280600681526020017f6465706c6f790000000000000000000000000000000000000000000000000000815250828280619f3d9061e6d1565b935060ff1681518110619f5257619f5261e564565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d65000000000000000000000000000000000000815250828280619f9f9061e6d1565b935060ff1681518110619fb457619fb461e564565b60200260200101819052508760200151828280619fd09061e6d1565b935060ff1681518110619fe557619fe561e564565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163745061746800000000000000000000000000000000000081525082828061a0329061e6d1565b935060ff168151811061a0475761a04761e564565b60209081029190910101528751828261a05f8161e6d1565b935060ff168151811061a0745761a07461e564565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e4964000000000000000000000000000000000000000000000081525082828061a0c19061e6d1565b935060ff168151811061a0d65761a0d661e564565b602002602001018190525061a0ea4661b64d565b828261a0f58161e6d1565b935060ff168151811061a10a5761a10a61e564565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c65000000000000000000000000000000000081525082828061a1579061e6d1565b935060ff168151811061a16c5761a16c61e564565b60200260200101819052508682828061a1849061e6d1565b935060ff168151811061a1995761a19961e564565b602090810291909101015285511561a2c05760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f646500000000000000000000006020820152828261a1ea8161e6d1565b935060ff168151811061a1ff5761a1ff61e564565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d9061a24f90899060040161db4a565b600060405180830381865afa15801561a26c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a294919081019061e1fe565b828261a29f8161e6d1565b935060ff168151811061a2b45761a2b461e564565b60200260200101819052505b84602001511561a3905760408051808201909152601281527f2d2d766572696679536f75726365436f646500000000000000000000000000006020820152828261a3098161e6d1565b935060ff168151811061a31e5761a31e61e564565b60200260200101819052506040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525082828061a36b9061e6d1565b935060ff168151811061a3805761a38061e564565b602002602001018190525061a557565b61a3c8619c6d8660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b61a45b5760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261a40b8161e6d1565b935060ff168151811061a4205761a42061e564565b60200260200101819052508460a0015160405160200161a440919061e593565b60405160208183030381529060405282828061a36b9061e6d1565b8460c0015115801561a49e57506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261a49c90511590565b155b1561a5575760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261a4e28161e6d1565b935060ff168151811061a4f75761a4f761e564565b602002602001018190525061a50b8861b6ed565b60405160200161a51b919061e593565b60405160208183030381529060405282828061a5369061e6d1565b935060ff168151811061a54b5761a54b61e564565b60200260200101819052505b6040808601518151808301835260008082526020918201528251808401909352815183529081019082015261a58b90511590565b61a6205760408051808201909152600b81527f2d2d72656c6179657249640000000000000000000000000000000000000000006020820152828261a5ce8161e6d1565b935060ff168151811061a5e35761a5e361e564565b6020026020010181905250846040015182828061a5ff9061e6d1565b935060ff168151811061a6145761a61461e564565b60200260200101819052505b60608501511561a7415760408051808201909152600681527f2d2d73616c7400000000000000000000000000000000000000000000000000006020820152828261a6698161e6d1565b935060ff168151811061a67e5761a67e61e564565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa15801561a6ed573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261a715919081019061e1fe565b828261a7208161e6d1565b935060ff168151811061a7355761a73561e564565b60200260200101819052505b60e0850151511561a7e85760408051808201909152600a81527f2d2d6761734c696d6974000000000000000000000000000000000000000000006020820152828261a78b8161e6d1565b935060ff168151811061a7a05761a7a061e564565b602002602001018190525061a7bc8560e001516000015161b64d565b828261a7c78161e6d1565b935060ff168151811061a7dc5761a7dc61e564565b60200260200101819052505b60e0850151602001511561a8925760408051808201909152600a81527f2d2d6761735072696365000000000000000000000000000000000000000000006020820152828261a8358161e6d1565b935060ff168151811061a84a5761a84a61e564565b602002602001018190525061a8668560e001516020015161b64d565b828261a8718161e6d1565b935060ff168151811061a8865761a88661e564565b60200260200101819052505b60e0850151604001511561a93c5760408051808201909152600e81527f2d2d6d61784665655065724761730000000000000000000000000000000000006020820152828261a8df8161e6d1565b935060ff168151811061a8f45761a8f461e564565b602002602001018190525061a9108560e001516040015161b64d565b828261a91b8161e6d1565b935060ff168151811061a9305761a93061e564565b60200260200101819052505b60e0850151606001511561a9e65760408051808201909152601681527f2d2d6d61785072696f72697479466565506572476173000000000000000000006020820152828261a9898161e6d1565b935060ff168151811061a99e5761a99e61e564565b602002602001018190525061a9ba8560e001516060015161b64d565b828261a9c58161e6d1565b935060ff168151811061a9da5761a9da61e564565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561aa045761aa0461d8ec565b60405190808252806020026020018201604052801561aa3757816020015b606081526020019060019003908161aa225790505b50905060005b8260ff168160ff16101561aa9057838160ff168151811061aa605761aa6061e564565b6020026020010151828260ff168151811061aa7d5761aa7d61e564565b602090810291909101015260010161aa3d565b5093505050505b949350505050565b61aac66040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c9161ab4c9186910161e75b565b600060405180830381865afa15801561ab69573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261ab91919081019061e1fe565b9050600061ab9f868361c1dc565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b815260040161abcf919061d52a565b6000604051808303816000875af115801561abee573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261ac16919081019061e7a2565b805190915060030b1580159061ac2f5750602081015151155b801561ac3e5750604081015151155b15619c18578160008151811061ac565761ac5661e564565b6020026020010151604051602001619271919061e858565b6060600061aca38560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252865182528087019082015290915061acda9082905b9061c331565b1561ae3757600061ad578261ad518461ad4b61ad1d8a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061c358565b9061c3ba565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061adbb90829061c331565b1561ae2557604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ae22905b829061c43f565b90505b61ae2e8161c465565b9250505061919e565b821561ae5057848460405160200161927192919061ea44565b505060408051602081019091526000815261919e565b509392505050565b6000808251602084016000f09392505050565b8160a001511561ae9057505050565b600061ae9d84848461c4ce565b9050600061aeaa8261aa9f565b602081015181519192509060030b15801561af465750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261af469060408051808201825260008082526020918201528151808301909252845182528085019082015261acd4565b1561af5357505050505050565b6040820151511561af73578160400151604051602001619271919061eaeb565b80604051602001619271919061eb49565b6060600061afb98360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061b01e905b829061b5ec565b1561b08d57604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261919e9061b08890839061ca69565b61c465565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b0ef905b829061caf3565b60010361b1bc57604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b1559061ae1b565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261919e9061b088905b839061c43f565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b21b9061b017565b1561b35257604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061b28390839061cb8d565b90506000816001835161b296919061da91565b8151811061b2a65761b2a661e564565b6020026020010151905061b34961b08861b31c6040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925285518252808601908201529061ca69565b95945050505050565b82604051602001619271919061ebb4565b50919050565b6060600061b39e8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061b4009061b017565b1561b40e5761919e8161c465565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b46d9061b0e8565b60010361b4d757604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261919e9061b0889061b1b5565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b5369061b017565b1561b35257604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061b59e90839061cb8d565b905060018151111561b5da57806002825161b5b9919061da91565b8151811061b5c95761b5c961e564565b602002602001015192505050919050565b5082604051602001619271919061ebb4565b80518251600091111561b60157506000618fc3565b8151835160208501516000929161b6179161daa4565b61b621919061da91565b90508260200151810361b638576001915050618fc3565b82516020840151819020912014905092915050565b6060600061b65a8361cc32565b600101905060008167ffffffffffffffff81111561b67a5761b67a61d8ec565b6040519080825280601f01601f19166020018201604052801561b6a4576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461b6ae57509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e534544000000000000000000000000000000000000000000008184019081528551808701875283815284019290925284518086019095525184529083015260609161b779905b829061cd14565b1561b7b957505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b8189061b772565b1561b85857505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b8b79061b772565b1561b8f757505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b9569061b772565b8061b9bb5750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261b9bb9061b772565b1561b9fb57505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261ba5a9061b772565b8061babf5750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261babf9061b772565b1561baff57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bb5e9061b772565b8061bbc35750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bbc39061b772565b1561bc0357505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bc629061b772565b8061bcc75750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bcc79061b772565b1561bd0757505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bd669061b772565b1561bda657505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261be059061b772565b1561be4557505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bea49061b772565b1561bee457505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bf439061b772565b1561bf8357505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261bfe29061b772565b1561c02257505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c0819061b772565b8061c0e65750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c0e69061b772565b1561c12657505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261c1859061b772565b1561c1c557505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151619271929060200161ec92565b60608060005b845181101561c267578185828151811061c1fe5761c1fe61e564565b602002602001015160405160200161c21792919061e0dd565b60405160208183030381529060405291506001855161c236919061da91565b811461c25f578160405160200161c24d919061edfb565b60405160208183030381529060405291505b60010161c1e2565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161c280579050509050838160008151811061c2ab5761c2ab61e564565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061c2ff5761c2ff61e564565b6020026020010181905250818160028151811061c31e5761c31e61e564565b6020908102919091010152949350505050565b602080830151835183519284015160009361c34f929184919061cd28565b14159392505050565b6040805180820190915260008082526020820152600061c38a846000015185602001518560000151866020015161ce39565b905083602001518161c39c919061da91565b8451859061c3ab90839061da91565b90525060208401525090919050565b604080518082019091526000808252602082015281518351101561c3df575081618fc3565b602080830151908401516001911461c4065750815160208481015190840151829020919020145b801561c4375782518451859061c41d90839061da91565b905250825160208501805161c43390839061daa4565b9052505b509192915050565b604080518082019091526000808252602082015261c45e83838361cf59565b5092915050565b60606000826000015167ffffffffffffffff81111561c4865761c48661d8ec565b6040519080825280601f01601f19166020018201604052801561c4b0576020820181803683370190505b509050600060208201905061c45e818560200151866000015161d004565b6060600061c4da619497565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161c4f757905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061c5529061e6d1565b935060ff168151811061c5675761c56761e564565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161c5b8919061ee3c565b60405160208183030381529060405282828061c5d39061e6d1565b935060ff168151811061c5e85761c5e861e564565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061c6359061e6d1565b935060ff168151811061c64a5761c64a61e564565b60200260200101819052508260405160200161c666919061e5ff565b60405160208183030381529060405282828061c6819061e6d1565b935060ff168151811061c6965761c69661e564565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061c6e39061e6d1565b935060ff168151811061c6f85761c6f861e564565b602002602001018190525061c70d878461d07e565b828261c7188161e6d1565b935060ff168151811061c72d5761c72d61e564565b60209081029190910101528551511561c7d95760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261c77f8161e6d1565b935060ff168151811061c7945761c79461e564565b602002602001018190525061c7ad86600001518461d07e565b828261c7b88161e6d1565b935060ff168151811061c7cd5761c7cd61e564565b60200260200101819052505b85608001511561c8475760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261c8228161e6d1565b935060ff168151811061c8375761c83761e564565b602002602001018190525061c8ad565b841561c8ad5760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261c88c8161e6d1565b935060ff168151811061c8a15761c8a161e564565b60200260200101819052505b6040860151511561c9495760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261c8f78161e6d1565b935060ff168151811061c90c5761c90c61e564565b6020026020010181905250856040015182828061c9289061e6d1565b935060ff168151811061c93d5761c93d61e564565b60200260200101819052505b85606001511561c9b35760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261c9928161e6d1565b935060ff168151811061c9a75761c9a761e564565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561c9d15761c9d161d8ec565b60405190808252806020026020018201604052801561ca0457816020015b606081526020019060019003908161c9ef5790505b50905060005b8260ff168160ff16101561ca5d57838160ff168151811061ca2d5761ca2d61e564565b6020026020010151828260ff168151811061ca4a5761ca4a61e564565b602090810291909101015260010161ca0a565b50979650505050505050565b604080518082019091526000808252602082015281518351101561ca8e575081618fc3565b8151835160208501516000929161caa49161daa4565b61caae919061da91565b6020840151909150600190821461cacf575082516020840151819020908220145b801561caea5783518551869061cae690839061da91565b9052505b50929392505050565b600080826000015161cb17856000015186602001518660000151876020015161ce39565b61cb21919061daa4565b90505b8351602085015161cb35919061daa4565b811161c45e578161cb458161ee81565b925050826000015161cb7c85602001518361cb60919061da91565b865161cb6c919061da91565b838660000151876020015161ce39565b61cb86919061daa4565b905061cb24565b6060600061cb9b848461caf3565b61cba690600161daa4565b67ffffffffffffffff81111561cbbe5761cbbe61d8ec565b60405190808252806020026020018201604052801561cbf157816020015b606081526020019060019003908161cbdc5790505b50905060005b815181101561ae665761cc0d61b088868661c43f565b82828151811061cc1f5761cc1f61e564565b602090810291909101015260010161cbf7565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061cc7b577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061cca7576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061ccc557662386f26fc10000830492506010015b6305f5e100831061ccdd576305f5e100830492506008015b612710831061ccf157612710830492506004015b6064831061cd03576064830492506002015b600a8310618fc35760010192915050565b600061cd20838361d0be565b159392505050565b60008085841161ce2f576020841161cddb576000841561cd7357600161cd4f86602061da91565b61cd5a90600861ee9b565b61cd6590600261ef99565b61cd6f919061da91565b1990505b835181168561cd82898961daa4565b61cd8c919061da91565b805190935082165b81811461cdc65787841161cdae578794505050505061aa97565b8361cdb88161efa5565b94505082845116905061cd94565b61cdd0878561daa4565b94505050505061aa97565b83832061cde8858861da91565b61cdf2908761daa4565b91505b85821061ce2d5784822080820361ce1a5761ce10868461daa4565b935050505061aa97565b61ce2560018461da91565b92505061cdf5565b505b5092949350505050565b6000838186851161cf44576020851161cef3576000851561ce8557600161ce6187602061da91565b61ce6c90600861ee9b565b61ce7790600261ef99565b61ce81919061da91565b1990505b8451811660008761ce968b8b61daa4565b61cea0919061da91565b855190915083165b82811461cee55781861061cecd5761cec08b8b61daa4565b965050505050505061aa97565b8561ced78161ee81565b96505083865116905061cea8565b85965050505050505061aa97565b508383206000905b61cf05868961da91565b821161cf425785832080820361cf21578394505050505061aa97565b61cf2c60018561daa4565b935050818061cf3a9061ee81565b92505061cefb565b505b61cf4e878761daa4565b979650505050505050565b6040805180820190915260008082526020820152600061cf8b856000015186602001518660000151876020015161ce39565b60208087018051918601919091525190915061cfa7908261da91565b83528451602086015161cfba919061daa4565b810361cfc9576000855261cffb565b8351835161cfd7919061daa4565b8551869061cfe690839061da91565b905250835161cff5908261daa4565b60208601525b50909392505050565b6020811061d03c578151835261d01b60208461daa4565b925061d02860208361daa4565b915061d03560208261da91565b905061d004565b600019811561d06b57600161d05283602061da91565b61d05e9061010061ef99565b61d068919061da91565b90505b9151835183169219169190911790915250565b6060600061d08c848461956a565b805160208083015160405193945061d0a69390910161efbc565b60405160208183030381529060405291505092915050565b815181516000919081111561d0d1575081515b6020808501519084015160005b8381101561d18a578251825180821461d15a57600019602087101561d1395760018461d10b89602061da91565b61d115919061daa4565b61d12090600861ee9b565b61d12b90600261ef99565b61d135919061da91565b1990505b818116838216818103911461d157579750618fc39650505050505050565b50505b61d16560208661daa4565b945061d17260208561daa4565b9350505060208161d183919061daa4565b905061d0de565b5084518651619c18919061f014565b610b67806200f03583390190565b6107b6806200fb9c83390190565b61106f806201035283390190565b61207280620113c183390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161d21461d219565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161d2146040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561d2cb5783516001600160a01b031683526020938401939092019160010161d2a4565b509095945050505050565b60005b8381101561d2f157818101518382015260200161d2d9565b50506000910152565b6000815180845261d31281602086016020860161d2d6565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561d422577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561d408577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261d3f284865161d2fa565b602095860195909450929092019160010161d3b8565b50919750505060209485019492909201915060010161d34e565b50929695505050505050565b600081518084526020840193506020830160005b8281101561d4825781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161d442565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561d422577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261d4f8604088018261d2fa565b905060208201519150868103602088015261d513818361d42e565b96505050602093840193919091019060010161d4b4565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561d422577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261d58c85835161d2fa565b9450602093840193919091019060010161d552565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561d422577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261d622604087018261d42e565b955050602093840193919091019060010161d5c9565b600181811c9082168061d64c57607f821691505b60208210810361b363577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000815461d6928161d638565b80855260018216801561d6ac576001811461d6e65761d71d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083166020870152602082151560051b870101935061d71d565b84600052602060002060005b8381101561d7145781546020828a01015260018201915060208101905061d6f2565b87016020019450505b50505092915050565b6001600160a01b0381541682526001600160a01b0360018201541660208301526002810154604083015260806060830152600061919e608084016003840161d685565b60208152600061919e602083018461d726565b6001600160a01b038316815260406020820152600061aa97604083018461d726565b80516001600160a01b038116811461d7b557600080fd5b919050565b60006020828403121561d7cc57600080fd5b61919e8261d79e565b6003811061d80c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff8816604083015286606083015261d8a0608083018761d7d5565b8460a08301526001600160a01b03841660c083015261cf4e60e08301846001600160a01b03169052565b60006020828403121561d8dc57600080fd5b8151801515811461919e57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561148357806000526020600020601f840160051c8101602085101561d9425750805b601f840160051c820191505b81811015613a43576000815560010161d94e565b815167ffffffffffffffff81111561d97c5761d97c61d8ec565b61d9908161d98a845461d638565b8461d91b565b6020601f82116001811461d9c4576000831561d9ac5750848201515b600019600385901b1c1916600184901b178455613a43565b600084815260208120601f198516915b8281101561d9f4578785015182556020948501946001909201910161d9d4565b508482101561da125786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60006020828403121561da3357600080fd5b5051919050565b8381526001600160a01b038316602082015260606040820152600061b349606083018461d726565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810381811115618fc357618fc361da62565b80820180821115618fc357618fc361da62565b60c08252600061daca60c084018361d685565b6001600160a01b0360018401541660208501526002830154604085015260ff600384015416151560608501526004830154608085015283810360a0850152618fbf816005850161d685565b60208152600061919e602083018461dab7565b6001600160a01b038316815260406020820152600061aa97604083018461dab7565b60208152600061919e602083018461d2fa565b602081526000618fc360208301600581527f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260400190565b60a08152600061dbad60a083018761d2fa565b6001600160a01b03861660208401528460408401526001600160a01b0384166060840152828103608084015261cf4e81600581527f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260400190565b600081516060845261dc22606085018261d2fa565b90506001600160a01b036020840151166020850152604083015160408501528091505092915050565b60808152600061dc5e608083018761dc0d565b8560208401526001600160a01b0385166040840152828103606084015261cf4e818561d2fa565b60a08152600061dc9860a083018861dc0d565b6001600160a01b03871660208401528560408401526001600160a01b0385166060840152828103608084015261dcce818561d2fa565b98975050505050505050565b6000806040838503121561dced57600080fd5b61dcf68361d79e565b6020939093015192949293505050565b600081546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b03600183015416604084015260a0606084015261dd4f60a084016002840161d685565b600383015460808501528091505092915050565b60e08152600061dd7660e083018961d2fa565b8760208401526001600160a01b0387166040840152828103606084015261dd9d818761d2fa565b855160808501526020860151151560a0850152905082810360c084015261ddc4818561dd06565b9998505050505050505050565b6001600160a01b03851681528360208201526001600160a01b0383166040820152608060608201526000619c18608083018461d726565b60c08152600061de1b60c083018861d2fa565b6001600160a01b0387166020840152828103604084015261de3c818761d2fa565b85516060850152602086015115156080850152905082810360a084015261dcce818561dd06565b610100815260066101008201527f5345434f4e440000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f534543000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff8816604083015286606083015261d8a0608083018761d7d5565b60808152600061df06608083018761d2fa565b8560208401526001600160a01b0385166040840152828103606084015261cf4e818561dd06565b83815282602082015260606040820152600061b349606083018461d2fa565b6001600160a01b038316815260406020820152600061aa97604083018461d2fa565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161dfa681601a85016020880161d2d6565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161dfe381601c84016020880161d2d6565b01601c01949350505050565b6040516060810167ffffffffffffffff8111828210171561e0125761e01261d8ec565b60405290565b60008067ffffffffffffffff84111561e0335761e03361d8ec565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561e0625761e06261d8ec565b60405283815290508082840185101561e07a57600080fd5b61ae6684602083018561d2d6565b600082601f83011261e09957600080fd5b61919e8383516020850161e018565b60006020828403121561e0ba57600080fd5b815167ffffffffffffffff81111561e0d157600080fd5b618fbf8482850161e088565b6000835161e0ef81846020880161d2d6565b83519083019061e10381836020880161d2d6565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161e14481601a85016020880161d2d6565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161e18181603384016020880161d2d6565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f5554000000000000000000000000000000000000000000606082015260806020820152600061919e608083018461d2fa565b60006020828403121561e21057600080fd5b815167ffffffffffffffff81111561e22757600080fd5b8201601f8101841361e23857600080fd5b618fbf8482516020840161e018565b6000855161e259818460208a0161d2d6565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161e293816001840160208a0161d2d6565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161e2d181600284016020890161d2d6565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161e31381600284016020880161d2d6565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061e35e604083018461d2fa565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161e3d581601f85016020870161d2d6565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061e442604083018461d2fa565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061e494604083018461d2fa565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161e50b81601485016020870161d2d6565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061e552604083018561d2fa565b828103602084015261919a818561d2fa565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161e5cb81600185016020870161d2d6565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161e61181846020870161d2d6565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161e6c481604b85016020870161d2d6565b91909101604b0192915050565b600060ff821660ff810361e6e75761e6e761da62565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161e74e81602985016020870161d2d6565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f5041544800000000000000000000606082015260806020820152600061919e608083018461d2fa565b60006020828403121561e7b457600080fd5b815167ffffffffffffffff81111561e7cb57600080fd5b82016060818503121561e7dd57600080fd5b61e7e561dfef565b81518060030b811461e7f657600080fd5b8152602082015167ffffffffffffffff81111561e81257600080fd5b61e81e8682850161e088565b602083015250604082015167ffffffffffffffff81111561e83e57600080fd5b61e84a8682850161e088565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161e8b681602185016020870161d2d6565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161eaa281602185016020880161d2d6565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161eadf81602e84016020880161d2d6565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161e74e81602985016020870161d2d6565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161eba781602285016020870161d2d6565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161ebec81600e85016020870161d2d6565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161ecca81601885016020880161d2d6565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161ed0781601c84016020880161d2d6565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161ee0d81846020870161d2d6565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161ee7481601c85016020870161d2d6565b91909101601c0192915050565b6000600019820361ee945761ee9461da62565b5060010190565b8082028115828204841417618fc357618fc361da62565b6001815b600184111561eeed5780850481111561eed15761eed161da62565b600184161561eedf57908102905b60019390931c92800261eeb6565b935093915050565b60008261ef0457506001618fc3565b8161ef1157506000618fc3565b816001811461ef27576002811461ef315761ef4d565b6001915050618fc3565b60ff84111561ef425761ef4261da62565b50506001821b618fc3565b5060208310610133831016604e8410600b841016171561ef70575081810a618fc3565b61ef7d600019848461eeb2565b806000190482111561ef915761ef9161da62565b029392505050565b600061919e838361eef5565b60008161efb45761efb461da62565b506000190190565b6000835161efce81846020880161d2d6565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161f00881600184016020880161d2d6565b01600101949350505050565b818103600083128015838313168383128216171561c45e5761c45e61da6256fe60c0604052600d60809081526c2bb930b83832b21022ba3432b960991b60a05260009061002c9082610114565b506040805180820190915260048152630ae8aa8960e31b60208201526001906100559082610114565b506002805460ff1916601217905534801561006f57600080fd5b506101d2565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061009f57607f821691505b6020821081036100bf57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561010f57806000526020600020601f840160051c810160208510156100ec5750805b601f840160051c820191505b8181101561010c57600081556001016100f8565b50505b505050565b81516001600160401b0381111561012d5761012d610075565b6101418161013b845461008b565b846100c5565b6020601f821160018114610175576000831561015d5750848201515b600019600385901b1c1916600184901b17845561010c565b600084815260208120601f198516915b828110156101a55787850151825560209485019460019092019101610185565b50848210156101c35786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610986806101e16000396000f3fe6080604052600436106100c05760003560e01c8063313ce56711610074578063a9059cbb1161004e578063a9059cbb146101fa578063d0e30db01461021a578063dd62ed3e1461022257600080fd5b8063313ce5671461018c57806370a08231146101b857806395d89b41146101e557600080fd5b806318160ddd116100a557806318160ddd1461012f57806323b872dd1461014c5780632e1a7d4d1461016c57600080fd5b806306fdde03146100d4578063095ea7b3146100ff57600080fd5b366100cf576100cd61025a565b005b600080fd5b3480156100e057600080fd5b506100e96102b5565b6040516100f69190610745565b60405180910390f35b34801561010b57600080fd5b5061011f61011a3660046107da565b610343565b60405190151581526020016100f6565b34801561013b57600080fd5b50475b6040519081526020016100f6565b34801561015857600080fd5b5061011f610167366004610804565b6103bd565b34801561017857600080fd5b506100cd610187366004610841565b610647565b34801561019857600080fd5b506002546101a69060ff1681565b60405160ff90911681526020016100f6565b3480156101c457600080fd5b5061013e6101d336600461085a565b60036020526000908152604090205481565b3480156101f157600080fd5b506100e9610724565b34801561020657600080fd5b5061011f6102153660046107da565b610731565b6100cd61025a565b34801561022e57600080fd5b5061013e61023d366004610875565b600460209081526000928352604080842090915290825290205481565b33600090815260036020526040812080543492906102799084906108d7565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b600080546102c2906108ea565b80601f01602080910402602001604051908101604052809291908181526020018280546102ee906108ea565b801561033b5780601f106103105761010080835404028352916020019161033b565b820191906000526020600020905b81548152906001019060200180831161031e57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103ab9086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604081205482111561042b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841633148015906104a1575073ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156105605773ffffffffffffffffffffffffffffffffffffffff8416600090815260046020908152604080832033845290915290205482111561051a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b73ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091528120805484929061055a90849061093d565b90915550505b73ffffffffffffffffffffffffffffffffffffffff84166000908152600360205260408120805484929061059590849061093d565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040812080548492906105cf9084906108d7565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161063591815260200190565b60405180910390a35060019392505050565b3360009081526003602052604090205481111561069a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b33600090815260036020526040812080548392906106b990849061093d565b9091555050604051339082156108fc029083906000818181858888f193505050501580156106eb573d6000803e3d6000fd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b600180546102c2906108ea565b600061073e3384846103bd565b9392505050565b602081526000825180602084015260005b818110156107735760208186018101516040868401015201610756565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146107d557600080fd5b919050565b600080604083850312156107ed57600080fd5b6107f6836107b1565b946020939093013593505050565b60008060006060848603121561081957600080fd5b610822846107b1565b9250610830602085016107b1565b929592945050506040919091013590565b60006020828403121561085357600080fd5b5035919050565b60006020828403121561086c57600080fd5b61073e826107b1565b6000806040838503121561088857600080fd5b610891836107b1565b915061089f602084016107b1565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156103b7576103b76108a8565b600181811c908216806108fe57607f821691505b602082108103610937577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b818103818111156103b7576103b76108a856fea2646970667358221220bb5b9dcef0ba90bcdefcbd63f71b1df95b50e29550a7456c69c6b9ff9dcdd20e64736f6c634300081a00336080604052348015600f57600080fd5b506107978061001f6000396000f3fe6080604052600436106100355760003560e01c80632d4cfb7e1461003e5780635bcfd6161461005e578063c9028a361461007e57005b3661003c57005b005b34801561004a57600080fd5b5061003c610059366004610182565b61009e565b34801561006a57600080fd5b5061003c6100793660046101ed565b6100d8565b34801561008a57600080fd5b5061003c6100993660046102aa565b610153565b7f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7816040516100cd9190610399565b60405180910390a150565b606081156100ef576100ec8284018461049f565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e61011a8780610595565b61012a60408a0160208b016105fa565b8960400135338660405161014396959493929190610615565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4816040516100cd91906106d7565b60006020828403121561019457600080fd5b813567ffffffffffffffff8111156101ab57600080fd5b820160c081850312156101bd57600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146101e857600080fd5b919050565b60008060008060006080868803121561020557600080fd5b853567ffffffffffffffff81111561021c57600080fd5b86016060818903121561022e57600080fd5b945061023c602087016101c4565b935060408601359250606086013567ffffffffffffffff81111561025f57600080fd5b8601601f8101881361027057600080fd5b803567ffffffffffffffff81111561028757600080fd5b88602082840101111561029957600080fd5b959894975092955050506020019190565b6000602082840312156102bc57600080fd5b813567ffffffffffffffff8111156102d357600080fd5b8201608081850312156101bd57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261031a57600080fd5b830160208101925035905067ffffffffffffffff81111561033a57600080fd5b80360382131561034957600080fd5b9250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6020815260006103a983846102e5565b60c060208501526103be60e085018284610350565b91505073ffffffffffffffffffffffffffffffffffffffff6103e2602086016101c4565b16604084015260006040850135905080606085015250606084013580151580821461040c57600080fd5b80608086015250506000608085013590508060a08501525061043160a08501856102e5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c0860152610466838284610350565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156104b157600080fd5b813567ffffffffffffffff8111156104c857600080fd5b8201601f810184136104d957600080fd5b803567ffffffffffffffff8111156104f3576104f3610470565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561055f5761055f610470565b60405281815282820160200186101561057757600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126105ca57600080fd5b83018035915067ffffffffffffffff8211156105e557600080fd5b60200191503681900382131561034957600080fd5b60006020828403121561060c57600080fd5b6101bd826101c4565b60a08152600061062960a08301888a610350565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b8181101561069357602081870181015184830182015201610677565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff6106f9836101c4565b16602082015273ffffffffffffffffffffffffffffffffffffffff610720602084016101c4565b166040820152600080604084013590508060608401525061074460608401846102e5565b60808085015261075860a085018284610350565b9594505050505056fea2646970667358221220fc81246aa4ca7c3f65ecc98dd976d57eaff81defb66e963ecbcfe23fc4c4507b64736f6c634300081a003360c060405234801561001057600080fd5b5060405161106f38038061106f83398101604081905261002f916100db565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006357604051632b2add3d60e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0385811691909117909155828116608052811660a0526040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac590600090a150505061011e565b80516001600160a01b03811681146100d657600080fd5b919050565b6000806000606084860312156100f057600080fd5b6100f9846100bf565b9250610107602085016100bf565b9150610115604085016100bf565b90509250925092565b60805160a051610f2561014a60003960006101e50152600081816102b9015261045b0152610f256000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806397770dff11610097578063c63585cc11610066578063c63585cc14610273578063d7fd7afb14610286578063d936a012146102b4578063ee2815ba146102db57600080fd5b806397770dff1461021a578063a7cb05071461022d578063c39aca3714610240578063c62178ac1461025357600080fd5b8063513a9c05116100d3578063513a9c051461018a578063569541b9146101c0578063842da36d146101e057806391dd645f1461020757600080fd5b80630be15547146100fa5780631f0e251b1461015a5780633ce4a5bc1461016f575b600080fd5b610130610108366004610bd1565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61016d610168366004610c13565b6102ee565b005b61013073735b14bb79463307aacbed86daf3322b1e6226ab81565b610130610198366004610bd1565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101309073ffffffffffffffffffffffffffffffffffffffff1681565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d610215366004610c35565b610402565b61016d610228366004610c13565b610526565b61016d61023b366004610c61565b610633565b61016d61024e366004610c83565b6106ce565b6004546101309073ffffffffffffffffffffffffffffffffffffffff1681565b610130610281366004610d53565b6108cd565b6102a6610294366004610bd1565b60006020819052908152604090205481565b604051908152602001610151565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d6102e9366004610c35565b610a02565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461033b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610388576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461044f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354600090610497907f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff16846108cd565b60008481526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251878152918201529192507f0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e910160405180910390a1505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610573576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166105c0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e906020016103f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610680576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461071b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831673735b14bb79463307aacbed86daf3322b1e6226ab1480610768575073ffffffffffffffffffffffffffffffffffffffff831630145b1561079f576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690528616906347e7ef24906044016020604051808303816000875af1158015610814573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108389190610d96565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063de43156e906108939089908990899088908890600401610e01565b600060405180830381600087803b1580156108ad57600080fd5b505af11580156108c1573d6000803e3d6000fd5b50505050505050505050565b60008060006108dc8585610ad3565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b16603482015291935091508690604801604051602081830303815290604052805190602001206040516020016109c29291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a4f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d91016106c2565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b3b576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610b75578284610b78565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610bca576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b600060208284031215610be357600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c0e57600080fd5b919050565b600060208284031215610c2557600080fd5b610c2e82610bea565b9392505050565b60008060408385031215610c4857600080fd5b82359150610c5860208401610bea565b90509250929050565b60008060408385031215610c7457600080fd5b50508035926020909101359150565b60008060008060008060a08789031215610c9c57600080fd5b863567ffffffffffffffff811115610cb357600080fd5b87016060818a031215610cc557600080fd5b9550610cd360208801610bea565b945060408701359350610ce860608801610bea565b9250608087013567ffffffffffffffff811115610d0457600080fd5b8701601f81018913610d1557600080fd5b803567ffffffffffffffff811115610d2c57600080fd5b896020828401011115610d3e57600080fd5b60208201935080925050509295509295509295565b600080600060608486031215610d6857600080fd5b610d7184610bea565b9250610d7f60208501610bea565b9150610d8d60408501610bea565b90509250925092565b600060208284031215610da857600080fd5b81518015158114610c2e57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60808152600086357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018112610e3957600080fd5b870160208101903567ffffffffffffffff811115610e5657600080fd5b803603821315610e6557600080fd5b60606080850152610e7a60e085018284610db8565b91505073ffffffffffffffffffffffffffffffffffffffff610e9e60208a01610bea565b1660a0840152604088013560c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401528281036060840152610ee3818587610db8565b9897505050505050505056fea2646970667358221220f69feb008c6ee4b966089046dae83838fc4badacdb1c3d4d0b6408081d4aaec464736f6c634300081a003360c060405234801561001057600080fd5b5060405161207238038061207283398101604081905261002f916101f0565b6001600160a01b038216158061004c57506001600160a01b038116155b1561006a5760405163d92e233d60e01b815260040160405180910390fd5b60066100768982610342565b5060076100838882610342565b506008805460ff191660ff881617905560808590528360028111156100aa576100aa610400565b60a08160028111156100be576100be610400565b905250600192909255600080546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506104169350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261013357600080fd5b81516001600160401b0381111561014c5761014c61010c565b604051601f8201601f19908116603f011681016001600160401b038111828210171561017a5761017a61010c565b60405281815283820160200185101561019257600080fd5b60005b828110156101b157602081860181015183830182015201610195565b506000918101602001919091529392505050565b8051600381106101d457600080fd5b919050565b80516001600160a01b03811681146101d457600080fd5b600080600080600080600080610100898b03121561020d57600080fd5b88516001600160401b0381111561022357600080fd5b61022f8b828c01610122565b60208b015190995090506001600160401b0381111561024d57600080fd5b6102598b828c01610122565b975050604089015160ff8116811461027057600080fd5b60608a0151909650945061028660808a016101c5565b60a08a0151909450925061029c60c08a016101d9565b91506102aa60e08a016101d9565b90509295985092959890939650565b600181811c908216806102cd57607f821691505b6020821081036102ed57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561033d57806000526020600020601f840160051c8101602085101561031a5750805b601f840160051c820191505b8181101561033a5760008155600101610326565b50505b505050565b81516001600160401b0381111561035b5761035b61010c565b61036f8161036984546102b9565b846102f3565b6020601f8211600181146103a3576000831561038b5750848201515b600019600385901b1c1916600184901b17845561033a565b600084815260208120601f198516915b828110156103d357878501518255602094850194600190920191016103b3565b50848210156103f15786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a051611c1b61045760003960006103440152600081816102f001528181610bdc01528181610ce201528181610efe01526110040152611c1b6000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806395d89b41116100f9578063ccc7759911610097578063eddeb12311610071578063eddeb12314610461578063f2441b3214610474578063f687d12a14610494578063fc5fecd5146104a757600080fd5b8063ccc77599146103d4578063d9eeebed146103e7578063dd62ed3e1461041b57600080fd5b8063b84c8246116100d3578063b84c824614610386578063c47f00271461039b578063c7012626146103ae578063c835d7cc146103c157600080fd5b806395d89b4114610337578063a3413d031461033f578063a9059cbb1461037357600080fd5b80633ce4a5bc116101665780634d8943bb116101405780634d8943bb146102ac57806370a08231146102b557806385e1f4d0146102eb5780638b851b951461031257600080fd5b80633ce4a5bc1461024657806342966c681461028657806347e7ef241461029957600080fd5b806318160ddd1161019757806318160ddd1461021657806323b872dd1461021e578063313ce5671461023157600080fd5b806306fdde03146101be578063091d2788146101dc578063095ea7b3146101f3575b600080fd5b6101c66104ba565b6040516101d39190611648565b60405180910390f35b6101e560015481565b6040519081526020016101d3565b610206610201366004611687565b61054c565b60405190151581526020016101d3565b6005546101e5565b61020661022c3660046116b3565b610563565b60085460405160ff90911681526020016101d3565b61026173735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6102066102943660046116f4565b6105fa565b6102066102a7366004611687565b61060e565b6101e560025481565b6101e56102c336600461170d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101e57f000000000000000000000000000000000000000000000000000000000000000081565b60085461026190610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101c6610767565b6103667f000000000000000000000000000000000000000000000000000000000000000081565b6040516101d3919061172a565b610206610381366004611687565b610776565b610399610394366004611832565b610783565b005b6103996103a9366004611832565b6107e0565b6102066103bc366004611883565b610839565b6103996103cf36600461170d565b610988565b6103996103e236600461170d565b610a9c565b6103ef610bb0565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101d3565b6101e56104293660046118dc565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b61039961046f3660046116f4565b610dce565b6000546102619073ffffffffffffffffffffffffffffffffffffffff1681565b6103996104a23660046116f4565b610e50565b6103ef6104b53660046116f4565b610ed2565b6060600680546104c990611915565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590611915565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b5050505050905090565b60006105593384846110ee565b5060015b92915050565b60006105708484846111f7565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054828110156105db576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ef85336105ea8685611997565b6110ee565b506001949350505050565b600061060633836113b2565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab1480159061064c575060005473ffffffffffffffffffffffffffffffffffffffff163314155b80156106755750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b156106ac576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106b683836114f4565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107569186906119aa565b60405180910390a250600192915050565b6060600780546104c990611915565b60006105593384846111f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107d0576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107dc8282611a1b565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461082d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107dc8282611a1b565b6000806000610846610bb0565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156108d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fc9190611b34565b610932576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61093c33856113b2565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161097591899189918791611b56565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109d5576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a22576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610ae9576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b36576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c679190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610cb6576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d699190611ba2565b905080600003610da5576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610db89190611bbb565b610dc29190611bd2565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e1b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a91565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e9d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f899190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610fd8576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015611067573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108b9190611ba2565b9050806000036110c7576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000906110d78784611bbb565b6110e19190611bd2565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661113b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611188576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611244576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611291576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040902054818110156112f1576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112fb8282611997565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260036020526040808220939093559085168152908120805484929061133e908490611bd2565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113a491815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113ff576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260409020548181101561145f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114698282611997565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812091909155600580548492906114a4908490611997565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111ea565b73ffffffffffffffffffffffffffffffffffffffff8216611541576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008282546115539190611bd2565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805483929061158d908490611bd2565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561160a576020818501810151868301820152016115ee565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061165b60208301846115e4565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461168457600080fd5b50565b6000806040838503121561169a57600080fd5b82356116a581611662565b946020939093013593505050565b6000806000606084860312156116c857600080fd5b83356116d381611662565b925060208401356116e381611662565b929592945050506040919091013590565b60006020828403121561170657600080fd5b5035919050565b60006020828403121561171f57600080fd5b813561165b81611662565b6020810160038310611765577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff8411156117b5576117b561176b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156118025761180261176b565b60405283815290508082840185101561181a57600080fd5b83836020830137600060208583010152509392505050565b60006020828403121561184457600080fd5b813567ffffffffffffffff81111561185b57600080fd5b8201601f8101841361186c57600080fd5b61187b8482356020840161179a565b949350505050565b6000806040838503121561189657600080fd5b823567ffffffffffffffff8111156118ad57600080fd5b8301601f810185136118be57600080fd5b6118cd8582356020840161179a565b95602094909401359450505050565b600080604083850312156118ef57600080fd5b82356118fa81611662565b9150602083013561190a81611662565b809150509250929050565b600181811c9082168061192957607f821691505b602082108103611962577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561055d5761055d611968565b6040815260006119bd60408301856115e4565b90508260208301529392505050565b601f821115611a1657806000526020600020601f840160051c810160208510156119f35750805b601f840160051c820191505b81811015611a1357600081556001016119ff565b50505b505050565b815167ffffffffffffffff811115611a3557611a3561176b565b611a4981611a438454611915565b846119cc565b6020601f821160018114611a9b5760008315611a655750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611a13565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611ae95787850151825560209485019460019092019101611ac9565b5084821015611b2557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b600060208284031215611b4657600080fd5b8151801515811461165b57600080fd5b608081526000611b6960808301876115e4565b6020830195909552506040810192909252606090910152919050565b600060208284031215611b9757600080fd5b815161165b81611662565b600060208284031215611bb457600080fd5b5051919050565b808202811582820484141761055d5761055d611968565b8082018082111561055d5761055d61196856fea26469706673582212208307d60e253f5034856b93df00e3e2f46b06f9765d906dbd93ee947935fc608764736f6c634300081a003347617320666565206e6f74206275726e656420636f72726563746c792066726f6d20676173205a524332305a5243323020746f6b656e732077657265206e6f74206275726e656420636f72726563746c795769746864726177616c20616d6f756e74206e6f74206275726e656420636f72726563746c792066726f6d207365636f6e64205a524332305a5243323020746f6b656e732077657265206e6f74206275726e656420636f72726563746c7920666f72207769746864726177416e6443616c6ca264697066735822122001767d5527995bdad787dfde906a57cbf38f2db1258ae5cede49280f8a678f5864736f6c634300081a0033",
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
