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

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
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

// GatewayZEVMOutboundTestMetaData contains all meta data concerning the GatewayZEVMOutboundTest contract.
var GatewayZEVMOutboundTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testDeposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfSenderNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContractIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContractFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextData\",\"inputs\":[{\"name\":\"origin\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"msgSender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataRevert\",\"inputs\":[{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5061fe848061003c6000396000f3fe608060405234801561001057600080fd5b506004361061034c5760003560e01c806385226c81116101bd578063c8814d2e116100f9578063eab7674e116100a2578063ef2b53941161007c578063ef2b53941461054a578063f1d98f1b14610552578063fa7626d41461055a578063fb339a1c1461056757600080fd5b8063eab7674e14610532578063eb78bd7d1461053a578063ec294d9f1461054257600080fd5b8063e09bc659116100d3578063e09bc659146104ed578063e20c9f71146104f5578063e63ab1e9146104fd57600080fd5b8063c8814d2e146104d5578063ca26929c146104dd578063cf2c3d1d146104e557600080fd5b8063996b767511610166578063b5508aa911610140578063b5508aa9146104a5578063b936be8c146104ad578063ba414fa6146104b5578063c35cb5e4146104cd57600080fd5b8063996b76751461048d5780639c9acd5d14610495578063b0464fdc1461049d57600080fd5b8063916a17c611610197578063916a17c61461047057806396d9d8761461048557806397f7661f1461044357600080fd5b806385226c811461044b578063884660a314610460578063890a2d671461046857600080fd5b80633e5e3c231161028c5780635cec7db5116102355780636efa04b51161020f5780636efa04b51461042b5780637cec29b0146104335780637f924c4e1461043b578063828d267c1461044357600080fd5b80635cec7db5146104065780636163f8ef1461040e57806366d9a9a01461041657600080fd5b806351336fb01161026657806351336fb0146103f657806358c9987f146103fe5780635b4c90e1146103a157600080fd5b80633e5e3c23146103de5780633f7286f4146103e657806348f4fd07146103ee57600080fd5b806327820625116102f95780632ade3880116102d35780632ade3880146103b1578063339bd828146103c65780633626c616146103ce5780633ab5b199146103d657600080fd5b806327820625146103995780632948df41146103a15780632acb21b4146103a957600080fd5b80631c785a141161032a5780631c785a141461036b5780631ed7831c146103735780632468bc0f1461039157600080fd5b8063084fafab146103515780630a9254e41461035b57806314b7a6da14610363575b600080fd5b61035961056f565b005b610359610745565b61035961128f565b6103596113df565b61037b611cf6565b604051610388919061a10d565b60405180910390f35b610359611d58565b610359612223565b6103596122eb565b6103596124d4565b6103b9612693565b604051610388919061a1a9565b6103596127d5565b610359612921565b610359612aa9565b61037b612c5a565b61037b612cba565b610359612d1a565b610359612e65565b610359612fb0565b610359613165565b610359613353565b61041e6134a1565b604051610388919061a30f565b610359613623565b6103596139c7565b610359613bb2565b610359613beb565b610453613d97565b604051610388919061a3ad565b610359613e67565b610359614146565b610478614301565b604051610388919061a424565b6103596143fc565b61035961454f565b6103596146a3565b6104786147f3565b6104536148ee565b6103596149be565b6104bd614b7b565b6040519015158152602001610388565b610359614c4f565b610359614e0a565b610359614f4c565b6103596151cf565b61035961538d565b61037b615546565b6105247f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b604051908152602001610388565b6103596155a6565b6103596156f4565b6103596158b2565b610359615a61565b610359615c12565b601f546104bd9060ff1681565b610359615d60565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156105fe57600080fd5b505af1158015610612573d6000803e3d6000fd5b505050507ff5d4d6ed1657736cfd625a04e0585c5def81b280f6bc1b7ed4add114e79f7e9a602c604051610646919061a5f0565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b600060405180830381600087803b1580156106a857600080fd5b505af11580156106bc573d6000803e3d6000fd5b50506020546024546040517fe7d926b40000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063e7d926b493506107119290911690602c9060040161a603565b600060405180830381600087803b15801561072b57600080fd5b505af115801561073f573d6000803e3d6000fd5b50505050565b602580547fffffffffffffffffffffffff000000000000000000000000000000000000000090811630179091556026805490911661123417905560405161078b9061a020565b604051809103906000f0801580156107a7573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600f81527f476174657761795a45564d2e736f6c00000000000000000000000000000000006020820152602554915160248101939093529216604482015261088b919060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f485cc95500000000000000000000000000000000000000000000000000000000179052615f13565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b039384168102919091179182905560208054919092049092167fffffffffffffffffffffffff000000000000000000000000000000000000000090921682178155604080517f2722feee0000000000000000000000000000000000000000000000000000000081529051632722feee926004808401939192918290030181865afa15801561094d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610971919061a625565b602780547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03929092169190911790556040516109b59061a02d565b604051809103906000f0801580156109d1573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161781556027546040517f06447d5600000000000000000000000000000000000000000000000000000000815292166004830152737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d569101600060405180830381600087803b158015610a6d57600080fd5b505af1158015610a81573d6000803e3d6000fd5b505050506000806000604051610a969061a03a565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610ad2573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602054604051601293600193849360009391921690610b289061a047565b610b379695949392919061a64e565b604051809103906000f080158015610b53573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179091556023546040517fee2815ba0000000000000000000000000000000000000000000000000000000081526001600482015260248101929092529091169063ee2815ba90604401600060405180830381600087803b158015610bea57600080fd5b505af1158015610bfe573d6000803e3d6000fd5b50506023546040517fa7cb050700000000000000000000000000000000000000000000000000000000815260016004820181905260248201526001600160a01b03909116925063a7cb05079150604401600060405180830381600087803b158015610c6857600080fd5b505af1158015610c7c573d6000803e3d6000fd5b50506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152633b9aca006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b158015610cfc57600080fd5b505af1158015610d10573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b158015610d6557600080fd5b505af1158015610d79573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af1158015610ded573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e11919061a743565b506021546025546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a060248201529116906347e7ef24906044016020604051808303816000875af1158015610e82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea6919061a743565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610f0557600080fd5b505af1158015610f19573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015610f8f57600080fd5b505af1158015610fa3573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a060248201529116925063095ea7b391506044016020604051808303816000875af1158015611017573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103b919061a743565b50602260009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b15801561108d57600080fd5b505af11580156110a1573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af1158015611115573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611139919061a743565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561119857600080fd5b505af11580156111ac573d6000803e3d6000fd5b5050604080516080810182526025546001600160a01b0390811682526000602080840182815260018587019081528651928301909652918152606084018190528351602c80549185167fffffffffffffffffffffffff00000000000000000000000000000000000000009092169190911781559151602d8054965167ffffffffffffffff1674010000000000000000000000000000000000000000027fffffffff000000000000000000000000000000000000000000000000000000009097169190941617949094179091559093509150602e9061128a908261a7db565b505050565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156112e857600080fd5b505af11580156112fc573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561136c57600080fd5b505af1158015611380573d6000803e3d6000fd5b50506020546021546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260016024820152600060448201529116925063f45346dc9150606401610711565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561143857600080fd5b505af115801561144c573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611537919060040161a89a565b600060405180830381600087803b15801561155157600080fd5b505af1158015611565573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156115b957600080fd5b505af11580156115cd573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561162a57600080fd5b505af115801561163e573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611729919060040161a89a565b600060405180830381600087803b15801561174357600080fd5b505af1158015611757573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156117ab57600080fd5b505af11580156117bf573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561181c57600080fd5b505af1158015611830573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561188457600080fd5b505af1158015611898573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd93c066500000000000000000000000000000000000000000000000000000000600482015260019250737109709ecfa91a80626ff3989d68f67f5b1dd12d915063c31eb0e090602401600060405180830381600087803b15801561190b57600080fd5b505af115801561191f573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561197c57600080fd5b505af1158015611990573d6000803e3d6000fd5b50506020546021546026546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810187905290821660448201529116925063f45346dc9150606401600060405180830381600087803b158015611a0857600080fd5b505af1158015611a1c573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015611a7957600080fd5b505af1158015611a8d573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611ae157600080fd5b505af1158015611af5573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a08231906024015b602060405180830381865afa158015611b49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b6d919061a8ad565b9050611b7a600082615f32565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611bd357600080fd5b505af1158015611be7573d6000803e3d6000fd5b50506020546021546026546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810188905290821660448201529116925063f45346dc9150606401600060405180830381600087803b158015611c5f57600080fd5b505af1158015611c73573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015611cc6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cea919061a8ad565b905061128a8382615f32565b60606016805480602002602001604051908101604052809291908181526020018280548015611d4e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611d30575b5050505050905090565b6022546027546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015611da9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dcd919061a8ad565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa158015611e1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e43919061a8ad565b6024546040519192506001600160a01b03163190600090611e669060200161a8c6565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff191660808301529150600090806094810160408051808303601f190181529181529082526027546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611f5457600080fd5b505af1158015611f68573d6000803e3d6000fd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e9450611fc393506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602754602054611ff3936001600160a01b03928316928c92169061a903565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561205457600080fd5b505af1158015612068573d6000803e3d6000fd5b50506020546024546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321501a9593506120bf9286928c92911690889060040161a9b4565b600060405180830381600087803b1580156120d957600080fd5b505af11580156120ed573d6000803e3d6000fd5b50506022546027546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015612140573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612164919061a8ad565b9050612179612173888861aa1d565b82615f32565b6022546020546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156121ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121ee919061a8ad565b90506121fa8682615f32565b612219612207898761aa30565b6024546001600160a01b031631615f32565b5050505050505050565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561228f57600080fd5b505af11580156122a3573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa7915060240161068e565b60006040516020016122fc9061a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b1580156123aa57600080fd5b505af11580156123be573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561242e57600080fd5b505af1158015612442573d6000803e3d6000fd5b50506020546024546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063bcf7f32b935061249e92869260009260019290911690899060040161aa43565b600060405180830381600087803b1580156124b857600080fd5b505af11580156124cc573d6000803e3d6000fd5b505050505050565b60006040516020016124e59061a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b1580156125af57600080fd5b505af11580156125c3573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015b600060405180830381600087803b15801561262157600080fd5b505af1158015612635573d6000803e3d6000fd5b50506020546021546024546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063bcf7f32b945061249e938793811692600192911690899060040161aa43565b6060601e805480602002602001604051908101604052809291908181526020016000905b828210156127cc57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156127b55783829060005260206000200180546127289061a4bb565b80601f01602080910402602001604051908101604052809291908181526020018280546127549061a4bb565b80156127a15780601f10612776576101008083540402835291602001916127a1565b820191906000526020600020905b81548152906001019060200180831161278457829003601f168201915b505050505081526020019060010190612709565b5050505081525050815260200190600101906126b7565b50505050905090565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561282e57600080fd5b505af1158015612842573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156128b257600080fd5b505af11580156128c6573d6000803e3d6000fd5b50506020546021546040517fcd73d4c70000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063cd73d4c793506107119290911690600190600090602c9060040161aa98565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561297d57600080fd5b505af1158015612991573d6000803e3d6000fd5b5050604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015612a0157600080fd5b505af1158015612a15573d6000803e3d6000fd5b50506020546021546027546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810187905290821660448201529116925063f45346dc91506064015b600060405180830381600087803b158015612a8e57600080fd5b505af1158015612aa2573d6000803e3d6000fd5b5050505050565b6000604051602001612aba9061a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015612b6857600080fd5b505af1158015612b7c573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015612bec57600080fd5b505af1158015612c00573d6000803e3d6000fd5b50506020546021546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063bcf7f32b935061249e9286921690600190600090899060040161aa43565b60606018805480602002602001604051908101604052809291908181526020018280548015611d4e576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611d30575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015611d4e576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611d30575050505050905090565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612d7357600080fd5b505af1158015612d87573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015612df757600080fd5b505af1158015612e0b573d6000803e3d6000fd5b50506020546024546040517fcd73d4c70000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063cd73d4c7935061071192600092600192911690602c9060040161aa98565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612ed157600080fd5b505af1158015612ee5573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015612f4257600080fd5b505af1158015612f56573d6000803e3d6000fd5b50506020546021546040517fcd73d4c70000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063cd73d4c7935061071192909116906001908590602c9060040161aa98565b6000604051602001612fc19061a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b15801561306f57600080fd5b505af1158015613083573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156130f357600080fd5b505af1158015613107573d6000803e3d6000fd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca37945061249e938793811692600092911690899060040161aa43565b60405160019060009061317a9060200161a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b15801561324457600080fd5b505af1158015613258573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156132b557600080fd5b505af11580156132c9573d6000803e3d6000fd5b50506020546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506321501a95915061331c90849087908590889060040161a9b4565b600060405180830381600087803b15801561333657600080fd5b505af115801561334a573d6000803e3d6000fd5b50505050505050565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156133bf57600080fd5b505af11580156133d3573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561343057600080fd5b505af1158015613444573d6000803e3d6000fd5b50506020546021546027546040517fcd73d4c70000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063cd73d4c7945061071193928316926001921690602c9060040161aa98565b6060601b805480602002602001604051908101604052809291908181526020016000905b828210156127cc57838290600052602060002090600202016040518060400160405290816000820180546134f89061a4bb565b80601f01602080910402602001604051908101604052809291908181526020018280546135249061a4bb565b80156135715780601f1061354657610100808354040283529160200191613571565b820191906000526020600020905b81548152906001019060200180831161355457829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561360b57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116135b85790505b505050505081525050815260200190600101906134c5565b602154602480546040516370a0823160e01b81526001600160a01b03918216600482015260009391909116916370a082319101602060405180830381865afa158015613673573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613697919061a8ad565b90506136a4600082615f32565b60006040516020016136b59061a8c6565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff191660808301529150600090806094810160408051808303601f190181529181529082526027546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156137a357600080fd5b505af11580156137b7573d6000803e3d6000fd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e945061381293506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602754602054613843936001600160a01b0392831692600192169061a903565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156138a457600080fd5b505af11580156138b8573d6000803e3d6000fd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450613916938793811692600192911690899060040161aa43565b600060405180830381600087803b15801561393057600080fd5b505af1158015613944573d6000803e3d6000fd5b5050602154602480546040516370a0823160e01b81526001600160a01b03918216600482015260009550921692506370a082319101602060405180830381865afa158015613996573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139ba919061a8ad565b905061073f600182615f32565b60006040516020016139d89061a8c6565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff191660808301529150600090806094810160408051808303601f190181529181529082526027546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015613ac657600080fd5b505af1158015613ada573d6000803e3d6000fd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e9450613b3593506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602754602054613b66936001600160a01b0392831692600192169061a903565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401612607565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401611b2c565b6000604051602001613bfc9061a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015613caa57600080fd5b505af1158015613cbe573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015613d2e57600080fd5b505af1158015613d42573d6000803e3d6000fd5b50506020546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506321501a95915061249e908490600190600090889060040161a9b4565b6060601a805480602002602001604051908101604052809291908181526020016000905b828210156127cc578382906000526020600020018054613dda9061a4bb565b80601f0160208091040260200160405190810160405280929190818152602001828054613e069061a4bb565b8015613e535780601f10613e2857610100808354040283529160200191613e53565b820191906000526020600020905b815481529060010190602001808311613e3657829003601f168201915b505050505081526020019060010190613dbb565b602154602480546040516370a0823160e01b81526001600160a01b03918216600482015260009391909116916370a082319101602060405180830381865afa158015613eb7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613edb919061a8ad565b9050613ee8600082615f32565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015613f7757600080fd5b505af1158015613f8b573d6000803e3d6000fd5b505050507ff5d4d6ed1657736cfd625a04e0585c5def81b280f6bc1b7ed4add114e79f7e9a602c604051613fbf919061a5f0565b60405180910390a160275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561402057600080fd5b505af1158015614034573d6000803e3d6000fd5b50506020546021546024546040517fcd73d4c70000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063cd73d4c7945061409193928316926001921690602c9060040161aa98565b600060405180830381600087803b1580156140ab57600080fd5b505af11580156140bf573d6000803e3d6000fd5b5050602154602480546040516370a0823160e01b81526001600160a01b03918216600482015260009550921692506370a082319101602060405180830381865afa158015614111573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614135919061a8ad565b9050614142600182615f32565b5050565b60405160019060009061415b9060200161a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b15801561422557600080fd5b505af1158015614239573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561429657600080fd5b505af11580156142aa573d6000803e3d6000fd5b50506020546027546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321501a95935061331c9286928992911690889060040161a9b4565b6060601d805480602002602001604051908101604052809291908181526020016000905b828210156127cc5760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156143e457602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116143915790505b50505050508152505081526020019060010190614325565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561445857600080fd5b505af115801561446c573d6000803e3d6000fd5b5050604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156144dc57600080fd5b505af11580156144f0573d6000803e3d6000fd5b50506020546021546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101869052911660448201819052925063f45346dc9150606401612a74565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156145a857600080fd5b505af11580156145bc573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561462c57600080fd5b505af1158015614640573d6000803e3d6000fd5b50506020546021546026546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526000602482015290821660448201529116925063f45346dc9150606401610711565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156146fc57600080fd5b505af1158015614710573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561478057600080fd5b505af1158015614794573d6000803e3d6000fd5b50506020546026546040517ff45346dc00000000000000000000000000000000000000000000000000000000815260006004820152600160248201526001600160a01b0391821660448201529116925063f45346dc9150606401610711565b6060601c805480602002602001604051908101604052809291908181526020016000905b828210156127cc5760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156148d657602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116148835790505b50505050508152505081526020019060010190614817565b60606019805480602002602001604051908101604052809291908181526020016000905b828210156127cc5783829060005260206000200180546149319061a4bb565b80601f016020809104026020016040519081016040528092919081815260200182805461495d9061a4bb565b80156149aa5780601f1061497f576101008083540402835291602001916149aa565b820191906000526020600020905b81548152906001019060200180831161498d57829003601f168201915b505050505081526020019060010190614912565b6024546040516001600160a01b0390911631906000906149e09060200161a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015614a8e57600080fd5b505af1158015614aa2573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015614b1257600080fd5b505af1158015614b26573d6000803e3d6000fd5b50506020546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506321501a95915061331c908490600190600090889060040161a9b4565b60085460009060ff1615614b93575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015614c24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614c48919061a8ad565b1415905090565b604051600190600090614c649060200161a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b158015614d2e57600080fd5b505af1158015614d42573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015614d9f57600080fd5b505af1158015614db3573d6000803e3d6000fd5b50506020546024546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321501a95935061331c9286928992911690889060040161a9b4565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614e6357600080fd5b505af1158015614e77573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015614ee757600080fd5b505af1158015614efb573d6000803e3d6000fd5b50506020546040517fe7d926b40000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063e7d926b4915061071190600090602c9060040161a603565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015614f9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614fc1919061a8ad565b9050614fce600082615f32565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561502757600080fd5b505af115801561503b573d6000803e3d6000fd5b5050604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156150ab57600080fd5b505af11580156150bf573d6000803e3d6000fd5b50506020546021546026546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810188905290821660448201529116925063f45346dc9150606401600060405180830381600087803b15801561513757600080fd5b505af115801561514b573d6000803e3d6000fd5b50506021546026546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561519e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906151c2919061a8ad565b905061128a600082615f32565b60006040516020016151e09061a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b1580156152aa57600080fd5b505af11580156152be573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561531b57600080fd5b505af115801561532f573d6000803e3d6000fd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca37945061249e938793811692600192911690899060040161aa43565b600060405160200161539e9061a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b15801561546857600080fd5b505af115801561547c573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156154d957600080fd5b505af11580156154ed573d6000803e3d6000fd5b50506020546021546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca37935061249e92869216906001908690899060040161aa43565b60606015805480602002602001604051908101604052809291908181526020018280548015611d4e576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611d30575050505050905090565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156155ff57600080fd5b505af1158015615613573d6000803e3d6000fd5b5050604051630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561568357600080fd5b505af1158015615697573d6000803e3d6000fd5b50506020546021546024546040517fcd73d4c70000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063cd73d4c7945061071193928316926000921690602c9060040161aa98565b60006040516020016157059061a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e091602480830192600092919082900301818387803b1580156157cf57600080fd5b505af11580156157e3573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561584057600080fd5b505af1158015615854573d6000803e3d6000fd5b50506020546021546027546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca37945061249e938793811692600192911690899060040161aa43565b60006040516020016158c39061a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b15801561597157600080fd5b505af1158015615985573d6000803e3d6000fd5b5050604051630618f58760e51b81527f19c08f49000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156159f557600080fd5b505af1158015615a09573d6000803e3d6000fd5b50506020546021546040517f21501a950000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506321501a95935061249e928692600092911690889060040161a9b4565b6000604051602001615a729061a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015615b2057600080fd5b505af1158015615b34573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015615ba457600080fd5b505af1158015615bb8573d6000803e3d6000fd5b50506020546021546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca37935061249e9286921690600190600090899060040161aa43565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015615c7e57600080fd5b505af1158015615c92573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015615cef57600080fd5b505af1158015615d03573d6000803e3d6000fd5b50506020546021546024546040517fcd73d4c70000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063cd73d4c7945061071193928316926001921690602c9060040161aa98565b6000604051602001615d719061a8c6565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526027546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791602480830192600092919082900301818387803b158015615e1f57600080fd5b505af1158015615e33573d6000803e3d6000fd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015615ea357600080fd5b505af1158015615eb7573d6000803e3d6000fd5b50506020546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca37935061249e92869260009260019290911690899060040161aa43565b6000615f1d61a054565b615f28848483615fb1565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c549060440160006040518083038186803b158015615f9d57600080fd5b505afa1580156124cc573d6000803e3d6000fd5b600080615fbe858461602c565b90506160216040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f7879000000815250828660405160200161600c92919061aacf565b60405160208183030381529060405285616038565b9150505b9392505050565b60006160258383616066565b60c0810151516000901561605c5761605584848460c00151616081565b9050616025565b6160558484616227565b60006160728383616312565b61602583836020015184616038565b60008061608c61631e565b9050600061609a86836163f1565b905060006160b18260600151836020015185616897565b905060006160c183838989616aa9565b905060006160ce82617926565b602081015181519192509060030b15616141578982604001516040516020016160f892919061aaf1565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526161389160040161a89a565b60405180910390fd5b60006161846040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001617af5565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906161d790849060040161a89a565b602060405180830381865afa1580156161f4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616218919061a625565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc9259061627c90879060040161a89a565b600060405180830381865afa158015616299573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526162c1919081019061ac2b565b905060006162ef82856040516020016162db92919061ac60565b604051602081830303815290604052617cf5565b90506001600160a01b038116615f285784846040516020016160f892919061ac8f565b61414282826000617d08565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c906163a590849060040161ad3a565b600060405180830381865afa1580156163c2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526163ea919081019061ad81565b9250505090565b6164236040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d905061646e6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b61647785617e0b565b60208201526000616487866181f0565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa1580156164c9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526164f1919081019061ad81565b8683856020015160405160200161650b949392919061adca565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb119061656390859060040161a89a565b600060405180830381865afa158015616580573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526165a8919081019061ad81565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906165f090849060040161aece565b602060405180830381865afa15801561660d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616631919061a743565b61664657816040516020016160f8919061af20565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061668b90849060040161afb2565b600060405180830381865afa1580156166a8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526166d0919081019061ad81565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f69061671790849060040161b004565b602060405180830381865afa158015616734573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616758919061a743565b156167ed576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906167a290849060040161b004565b600060405180830381865afa1580156167bf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526167e7919081019061ad81565b60408501525b846001600160a01b03166349c4fac8828660000151604051602001616812919061b056565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161683e92919061b0c2565b600060405180830381865afa15801561685b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616883919081019061ad81565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b60608152602001906001900390816168b35790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250816000815181106169135761691361b0e7565b60200260200101819052506040518060400160405280600381526020017f2d726c0000000000000000000000000000000000000000000000000000000000815250816001815181106169675761696761b0e7565b602002602001018190525084604051602001616983919061b116565b604051602081830303815290604052816002815181106169a5576169a561b0e7565b6020026020010181905250826040516020016169c1919061b182565b604051602081830303815290604052816003815181106169e3576169e361b0e7565b602002602001018190525060006169f982617926565b602080820151604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000008185019081528251808401845260008082529086015282518084019093529051825292810192909252919250616a8a9060408051808201825260008082526020918201528151808301909252845182528085019082015290618473565b616a9f57856040516020016160f8919061b1c3565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015616af9565b511590565b616c6d57826020015115616bb5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401616138565b8260c0015115616c6d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401616138565b6040805160ff8082526120008201909252600091816020015b6060815260200190600190039081616c8657905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280616ce19061b254565b935060ff1681518110616cf657616cf661b0e7565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e3700000000000000000000000000000000000000815250604051602001616d47919061b273565b604051602081830303815290604052828280616d629061b254565b935060ff1681518110616d7757616d7761b0e7565b60200260200101819052506040518060400160405280600681526020017f6465706c6f790000000000000000000000000000000000000000000000000000815250828280616dc49061b254565b935060ff1681518110616dd957616dd961b0e7565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d65000000000000000000000000000000000000815250828280616e269061b254565b935060ff1681518110616e3b57616e3b61b0e7565b60200260200101819052508760200151828280616e579061b254565b935060ff1681518110616e6c57616e6c61b0e7565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e747261637450617468000000000000000000000000000000000000815250828280616eb99061b254565b935060ff1681518110616ece57616ece61b0e7565b602090810291909101015287518282616ee68161b254565b935060ff1681518110616efb57616efb61b0e7565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e49640000000000000000000000000000000000000000000000815250828280616f489061b254565b935060ff1681518110616f5d57616f5d61b0e7565b6020026020010181905250616f71466184d4565b8282616f7c8161b254565b935060ff1681518110616f9157616f9161b0e7565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c650000000000000000000000000000000000815250828280616fde9061b254565b935060ff1681518110616ff357616ff361b0e7565b60200260200101819052508682828061700b9061b254565b935060ff16815181106170205761702061b0e7565b60209081029190910101528551156171475760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826170718161b254565b935060ff16815181106170865761708661b0e7565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906170d690899060040161a89a565b600060405180830381865afa1580156170f3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261711b919081019061ad81565b82826171268161b254565b935060ff168151811061713b5761713b61b0e7565b60200260200101819052505b8460200151156172175760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826171908161b254565b935060ff16815181106171a5576171a561b0e7565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806171f29061b254565b935060ff16815181106172075761720761b0e7565b60200260200101819052506173de565b61724f616af48660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6172e25760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826172928161b254565b935060ff16815181106172a7576172a761b0e7565b60200260200101819052508460a001516040516020016172c7919061b116565b6040516020818303038152906040528282806171f29061b254565b8460c0015115801561732557506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261732390511590565b155b156173de5760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826173698161b254565b935060ff168151811061737e5761737e61b0e7565b602002602001018190525061739288618574565b6040516020016173a2919061b116565b6040516020818303038152906040528282806173bd9061b254565b935060ff16815181106173d2576173d261b0e7565b60200260200101819052505b6040808601518151808301835260008082526020918201528251808401909352815183529081019082015261741290511590565b6174a75760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826174558161b254565b935060ff168151811061746a5761746a61b0e7565b602002602001018190525084604001518282806174869061b254565b935060ff168151811061749b5761749b61b0e7565b60200260200101819052505b6060850151156175c85760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826174f08161b254565b935060ff16815181106175055761750561b0e7565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015617574573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261759c919081019061ad81565b82826175a78161b254565b935060ff16815181106175bc576175bc61b0e7565b60200260200101819052505b60e0850151511561766f5760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826176128161b254565b935060ff16815181106176275761762761b0e7565b60200260200101819052506176438560e00151600001516184d4565b828261764e8161b254565b935060ff16815181106176635761766361b0e7565b60200260200101819052505b60e085015160200151156177195760408051808201909152600a81527f2d2d676173507269636500000000000000000000000000000000000000000000602082015282826176bc8161b254565b935060ff16815181106176d1576176d161b0e7565b60200260200101819052506176ed8560e00151602001516184d4565b82826176f88161b254565b935060ff168151811061770d5761770d61b0e7565b60200260200101819052505b60e085015160400151156177c35760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826177668161b254565b935060ff168151811061777b5761777b61b0e7565b60200260200101819052506177978560e00151604001516184d4565b82826177a28161b254565b935060ff16815181106177b7576177b761b0e7565b60200260200101819052505b60e0850151606001511561786d5760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826178108161b254565b935060ff16815181106178255761782561b0e7565b60200260200101819052506178418560e00151606001516184d4565b828261784c8161b254565b935060ff16815181106178615761786161b0e7565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561788b5761788b61a765565b6040519080825280602002602001820160405280156178be57816020015b60608152602001906001900390816178a95790505b50905060005b8260ff168160ff16101561791757838160ff16815181106178e7576178e761b0e7565b6020026020010151828260ff16815181106179045761790461b0e7565b60209081029190910101526001016178c4565b5093505050505b949350505050565b61794d6040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c916179d39186910161b2de565b600060405180830381865afa1580156179f0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617a18919081019061ad81565b90506000617a268683619063565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b8152600401617a56919061a3ad565b6000604051808303816000875af1158015617a75573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617a9d919081019061b325565b805190915060030b15801590617ab65750602081015151155b8015617ac55750604081015151155b15616a9f5781600081518110617add57617add61b0e7565b60200260200101516040516020016160f8919061b3db565b60606000617b2a8560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528651825280870190820152909150617b619082905b906191b8565b15617cbe576000617bde82617bd884617bd2617ba48a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b906191df565b90619241565b604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150617c429082906191b8565b15617cac57604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617ca9905b82906192c6565b90505b617cb5816192ec565b92505050616025565b8215617cd75784846040516020016160f892919061b5c7565b5050604080516020810190915260008152616025565b509392505050565b6000808251602084016000f09392505050565b8160a0015115617d1757505050565b6000617d24848484619355565b90506000617d3182617926565b602081015181519192509060030b158015617dcd5750604080518082018252600781527f535543434553530000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617dcd90604080518082018252600080825260209182015281518083019092528451825280850190820152617b5b565b15617dda57505050505050565b60408201515115617dfa5781604001516040516020016160f8919061b66e565b806040516020016160f8919061b6cc565b60606000617e408360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150617ea5905b8290618473565b15617f1457604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261602590617f0f9083906198f0565b6192ec565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617f76905b829061997a565b60010361804357604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617fdc90617ca2565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261602590617f0f905b83906192c6565b604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526180a290617e9e565b156181d957604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061810a908390619a14565b90506000816001835161811d919061aa1d565b8151811061812d5761812d61b0e7565b602002602001015190506181d0617f0f6181a36040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528551825280860190820152906198f0565b95945050505050565b826040516020016160f8919061b737565b50919050565b606060006182258360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061828790617e9e565b1561829557616025816192ec565b604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526182f490617f6f565b60010361835e57604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261602590617f0f9061803c565b604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526183bd90617e9e565b156181d957604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290618425908390619a14565b9050600181511115618461578060028251618440919061aa1d565b815181106184505761845061b0e7565b602002602001015192505050919050565b50826040516020016160f8919061b737565b80518251600091111561848857506000615f2c565b8151835160208501516000929161849e9161aa30565b6184a8919061aa1d565b9050826020015181036184bf576001915050615f2c565b82516020840151819020912014905092915050565b606060006184e183619ab9565b600101905060008167ffffffffffffffff8111156185015761850161a765565b6040519080825280601f01601f19166020018201604052801561852b576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461853557509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e5345440000000000000000000000000000000000000000000081840190815285518087018752838152840192909252845180860190955251845290830152606091618600905b8290619b9b565b1561864057505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261869f906185f9565b156186df57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261873e906185f9565b1561877e57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526187dd906185f9565b806188425750604080518082018252601081527f47504c2d322e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618842906185f9565b1561888257505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526188e1906185f9565b806189465750604080518082018252601081527f47504c2d332e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618946906185f9565b1561898657505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526189e5906185f9565b80618a4a5750604080518082018252601181527f4c47504c2d322e312d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618a4a906185f9565b15618a8a57505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618ae9906185f9565b80618b4e5750604080518082018252601181527f4c47504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618b4e906185f9565b15618b8e57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618bed906185f9565b15618c2d57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618c8c906185f9565b15618ccc57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618d2b906185f9565b15618d6b57505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618dca906185f9565b15618e0a57505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e300000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618e69906185f9565b15618ea957505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618f08906185f9565b80618f6d5750604080518082018252601181527f4147504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618f6d906185f9565b15618fad57505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261900c906185f9565b1561904c57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516160f8929060200161b815565b60608060005b84518110156190ee57818582815181106190855761908561b0e7565b602002602001015160405160200161909e92919061ac60565b6040516020818303038152906040529150600185516190bd919061aa1d565b81146190e657816040516020016190d4919061b97e565b60405160208183030381529060405291505b600101619069565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161910757905050905083816000815181106191325761913261b0e7565b60200260200101819052506040518060400160405280600281526020017f2d63000000000000000000000000000000000000000000000000000000000000815250816001815181106191865761918661b0e7565b602002602001018190525081816002815181106191a5576191a561b0e7565b6020908102919091010152949350505050565b60208083015183518351928401516000936191d69291849190619baf565b14159392505050565b604080518082019091526000808252602082015260006192118460000151856020015185600001518660200151619cc0565b9050836020015181619223919061aa1d565b8451859061923290839061aa1d565b90525060208401525090919050565b6040805180820190915260008082526020820152815183511015619266575081615f2c565b602080830151908401516001911461928d5750815160208481015190840151829020919020145b80156192be578251845185906192a490839061aa1d565b90525082516020850180516192ba90839061aa30565b9052505b509192915050565b60408051808201909152600080825260208201526192e5838383619de0565b5092915050565b60606000826000015167ffffffffffffffff81111561930d5761930d61a765565b6040519080825280601f01601f191660200182016040528015619337576020820181803683370190505b50905060006020820190506192e58185602001518660000151619e8b565b6060600061936161631e565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161937e57905050905060006040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806193d99061b254565b935060ff16815181106193ee576193ee61b0e7565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161943f919061b9bf565b60405160208183030381529060405282828061945a9061b254565b935060ff168151811061946f5761946f61b0e7565b60200260200101819052506040518060400160405280600881526020017f76616c69646174650000000000000000000000000000000000000000000000008152508282806194bc9061b254565b935060ff16815181106194d1576194d161b0e7565b6020026020010181905250826040516020016194ed919061b182565b6040516020818303038152906040528282806195089061b254565b935060ff168151811061951d5761951d61b0e7565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061956a9061b254565b935060ff168151811061957f5761957f61b0e7565b60200260200101819052506195948784619f05565b828261959f8161b254565b935060ff16815181106195b4576195b461b0e7565b6020908102919091010152855151156196605760408051808201909152600b81527f2d2d7265666572656e6365000000000000000000000000000000000000000000602082015282826196068161b254565b935060ff168151811061961b5761961b61b0e7565b6020026020010181905250619634866000015184619f05565b828261963f8161b254565b935060ff16815181106196545761965461b0e7565b60200260200101819052505b8560800151156196ce5760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b0000000000000000602082015282826196a98161b254565b935060ff16815181106196be576196be61b0e7565b6020026020010181905250619734565b84156197345760408051808201909152601281527f2d2d726571756972655265666572656e63650000000000000000000000000000602082015282826197138161b254565b935060ff16815181106197285761972861b0e7565b60200260200101819052505b604086015151156197d05760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261977e8161b254565b935060ff16815181106197935761979361b0e7565b602002602001018190525085604001518282806197af9061b254565b935060ff16815181106197c4576197c461b0e7565b60200260200101819052505b85606001511561983a5760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826198198161b254565b935060ff168151811061982e5761982e61b0e7565b60200260200101819052505b60008160ff1667ffffffffffffffff8111156198585761985861a765565b60405190808252806020026020018201604052801561988b57816020015b60608152602001906001900390816198765790505b50905060005b8260ff168160ff1610156198e457838160ff16815181106198b4576198b461b0e7565b6020026020010151828260ff16815181106198d1576198d161b0e7565b6020908102919091010152600101619891565b50979650505050505050565b6040805180820190915260008082526020820152815183511015619915575081615f2c565b8151835160208501516000929161992b9161aa30565b619935919061aa1d565b60208401519091506001908214619956575082516020840151819020908220145b80156199715783518551869061996d90839061aa1d565b9052505b50929392505050565b600080826000015161999e8560000151866020015186600001518760200151619cc0565b6199a8919061aa30565b90505b835160208501516199bc919061aa30565b81116192e557816199cc8161ba04565b9250508260000151619a038560200151836199e7919061aa1d565b86516199f3919061aa1d565b8386600001518760200151619cc0565b619a0d919061aa30565b90506199ab565b60606000619a22848461997a565b619a2d90600161aa30565b67ffffffffffffffff811115619a4557619a4561a765565b604051908082528060200260200182016040528015619a7857816020015b6060815260200190600190039081619a635790505b50905060005b8151811015617ced57619a94617f0f86866192c6565b828281518110619aa657619aa661b0e7565b6020908102919091010152600101619a7e565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310619b02577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310619b2e576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310619b4c57662386f26fc10000830492506010015b6305f5e1008310619b64576305f5e100830492506008015b6127108310619b7857612710830492506004015b60648310619b8a576064830492506002015b600a8310615f2c5760010192915050565b6000619ba78383619f45565b159392505050565b600080858411619cb65760208411619c625760008415619bfa576001619bd686602061aa1d565b619be190600861ba1e565b619bec90600261bb1c565b619bf6919061aa1d565b1990505b8351811685619c09898961aa30565b619c13919061aa1d565b805190935082165b818114619c4d57878411619c35578794505050505061791e565b83619c3f8161bb28565b945050828451169050619c1b565b619c57878561aa30565b94505050505061791e565b838320619c6f858861aa1d565b619c79908761aa30565b91505b858210619cb457848220808203619ca157619c97868461aa30565b935050505061791e565b619cac60018461aa1d565b925050619c7c565b505b5092949350505050565b60008381868511619dcb5760208511619d7a5760008515619d0c576001619ce887602061aa1d565b619cf390600861ba1e565b619cfe90600261bb1c565b619d08919061aa1d565b1990505b84518116600087619d1d8b8b61aa30565b619d27919061aa1d565b855190915083165b828114619d6c57818610619d5457619d478b8b61aa30565b965050505050505061791e565b85619d5e8161ba04565b965050838651169050619d2f565b85965050505050505061791e565b508383206000905b619d8c868961aa1d565b8211619dc957858320808203619da8578394505050505061791e565b619db360018561aa30565b9350508180619dc19061ba04565b925050619d82565b505b619dd5878761aa30565b979650505050505050565b60408051808201909152600080825260208201526000619e128560000151866020015186600001518760200151619cc0565b602080870180519186019190915251909150619e2e908261aa1d565b835284516020860151619e41919061aa30565b8103619e505760008552619e82565b83518351619e5e919061aa30565b85518690619e6d90839061aa1d565b9052508351619e7c908261aa30565b60208601525b50909392505050565b60208110619ec35781518352619ea260208461aa30565b9250619eaf60208361aa30565b9150619ebc60208261aa1d565b9050619e8b565b6000198115619ef2576001619ed983602061aa1d565b619ee59061010061bb1c565b619eef919061aa1d565b90505b9151835183169219169190911790915250565b60606000619f1384846163f1565b8051602080830151604051939450619f2d9390910161bb3f565b60405160208183030381529060405291505092915050565b8151815160009190811115619f58575081515b6020808501519084015160005b8381101561a0115782518251808214619fe1576000196020871015619fc057600184619f9289602061aa1d565b619f9c919061aa30565b619fa790600861ba1e565b619fb290600261bb1c565b619fbc919061aa1d565b1990505b8181168382168181039114619fde579750615f2c9650505050505050565b50505b619fec60208661aa30565b9450619ff960208561aa30565b9350505060208161a00a919061aa30565b9050619f65565b5084518651616a9f919061bb97565b610b678061bbb883390190565b61064f8061c71f83390190565b61106f8061cd6e83390190565b6120728061dddd83390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161a09761a09c565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161a0976040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561a14e5783516001600160a01b031683526020938401939092019160010161a127565b509095945050505050565b60005b8381101561a17457818101518382015260200161a15c565b50506000910152565b6000815180845261a19581602086016020860161a159565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561a2a5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561a28b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261a27584865161a17d565b602095860195909450929092019160010161a23b565b50919750505060209485019492909201915060010161a1d1565b50929695505050505050565b600081518084526020840193506020830160005b8281101561a3055781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161a2c5565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561a2a5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261a37b604088018261a17d565b905060208201519150868103602088015261a396818361a2b1565b96505050602093840193919091019060010161a337565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561a2a5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261a40f85835161a17d565b9450602093840193919091019060010161a3d5565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561a2a5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261a4a5604087018261a2b1565b955050602093840193919091019060010161a44c565b600181811c9082168061a4cf57607f821691505b6020821081036181ea577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6001600160a01b038154168252600060018201546001600160a01b038116602085015267ffffffffffffffff8160a01c1660408501525060028201608060608501526000815461a5578161a4bb565b806080880152600182166000811461a576576001811461a5b05761a5e4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660a089015260a082151560051b890101935061a5e4565b84600052602060002060005b8381101561a5db5781548a820160a0015260019091019060200161a5bc565b890160a0019450505b50919695505050505050565b602081526000616025602083018461a508565b6001600160a01b038316815260406020820152600061791e604083018461a508565b60006020828403121561a63757600080fd5b81516001600160a01b038116811461602557600080fd5b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff881660408301528660608301526003861061a708577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8560808301528460a083015261a72960c08301856001600160a01b03169052565b6001600160a01b03831660e0830152979650505050505050565b60006020828403121561a75557600080fd5b8151801515811461602557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561128a57806000526020600020601f840160051c8101602085101561a7bb5750805b601f840160051c820191505b81811015612aa2576000815560010161a7c7565b815167ffffffffffffffff81111561a7f55761a7f561a765565b61a8098161a803845461a4bb565b8461a794565b6020601f82116001811461a83d576000831561a8255750848201515b600019600385901b1c1916600184901b178455612aa2565b600084815260208120601f198516915b8281101561a86d578785015182556020948501946001909201910161a84d565b508482101561a88b5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b602081526000616025602083018461a17d565b60006020828403121561a8bf57600080fd5b5051919050565b602081526000615f2c60208301600581527f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260400190565b60a08152600061a91660a083018761a17d565b6001600160a01b03861660208401528460408401526001600160a01b03841660608401528281036080840152619dd581600581527f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260400190565b600081516060845261a98b606085018261a17d565b90506001600160a01b036020840151166020850152604083015160408501528091505092915050565b60808152600061a9c7608083018761a976565b8560208401526001600160a01b03851660408401528281036060840152619dd5818561a17d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810381811115615f2c57615f2c61a9ee565b80820180821115615f2c57615f2c61a9ee565b60a08152600061aa5660a083018861a976565b6001600160a01b03871660208401528560408401526001600160a01b0385166060840152828103608084015261aa8c818561a17d565b98975050505050505050565b6001600160a01b03851681528360208201526001600160a01b0383166040820152608060608201526000616a9f608083018461a508565b6001600160a01b038316815260406020820152600061791e604083018461a17d565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161ab2981601a85016020880161a159565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161ab6681601c84016020880161a159565b01601c01949350505050565b6040516060810167ffffffffffffffff8111828210171561ab955761ab9561a765565b60405290565b60008067ffffffffffffffff84111561abb65761abb661a765565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561abe55761abe561a765565b60405283815290508082840185101561abfd57600080fd5b617ced84602083018561a159565b600082601f83011261ac1c57600080fd5b6160258383516020850161ab9b565b60006020828403121561ac3d57600080fd5b815167ffffffffffffffff81111561ac5457600080fd5b615f288482850161ac0b565b6000835161ac7281846020880161a159565b83519083019061ac8681836020880161a159565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161acc781601a85016020880161a159565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161ad0481603384016020880161a159565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000616025608083018461a17d565b60006020828403121561ad9357600080fd5b815167ffffffffffffffff81111561adaa57600080fd5b8201601f8101841361adbb57600080fd5b615f288482516020840161ab9b565b6000855161addc818460208a0161a159565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161ae16816001840160208a0161a159565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161ae5481600284016020890161a159565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161ae9681600284016020880161a159565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061aee1604083018461a17d565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161af5881601f85016020870161a159565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061afc5604083018461a17d565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061b017604083018461a17d565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161b08e81601485016020870161a159565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061b0d5604083018561a17d565b8281036020840152616021818561a17d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161b14e81600185016020870161a159565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161b19481846020870161a159565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161b24781604b85016020870161a159565b91909101604b0192915050565b600060ff821660ff810361b26a5761b26a61a9ee565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161b2d181602985016020870161a159565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000616025608083018461a17d565b60006020828403121561b33757600080fd5b815167ffffffffffffffff81111561b34e57600080fd5b82016060818503121561b36057600080fd5b61b36861ab72565b81518060030b811461b37957600080fd5b8152602082015167ffffffffffffffff81111561b39557600080fd5b61b3a18682850161ac0b565b602083015250604082015167ffffffffffffffff81111561b3c157600080fd5b61b3cd8682850161ac0b565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161b43981602185016020870161a159565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161b62581602185016020880161a159565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161b66281602e84016020880161a159565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161b2d181602985016020870161a159565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161b72a81602285016020870161a159565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161b76f81600e85016020870161a159565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161b84d81601885016020880161a159565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161b88a81601c84016020880161a159565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161b99081846020870161a159565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161b9f781601c85016020870161a159565b91909101601c0192915050565b6000600019820361ba175761ba1761a9ee565b5060010190565b8082028115828204841417615f2c57615f2c61a9ee565b6001815b600184111561ba705780850481111561ba545761ba5461a9ee565b600184161561ba6257908102905b60019390931c92800261ba39565b935093915050565b60008261ba8757506001615f2c565b8161ba9457506000615f2c565b816001811461baaa576002811461bab45761bad0565b6001915050615f2c565b60ff84111561bac55761bac561a9ee565b50506001821b615f2c565b5060208310610133831016604e8410600b841016171561baf3575081810a615f2c565b61bb00600019848461ba35565b806000190482111561bb145761bb1461a9ee565b029392505050565b6000616025838361ba78565b60008161bb375761bb3761a9ee565b506000190190565b6000835161bb5181846020880161a159565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161bb8b81600184016020880161a159565b01600101949350505050565b81810360008312801583831316838312821617156192e5576192e561a9ee56fe60c0604052600d60809081526c2bb930b83832b21022ba3432b960991b60a05260009061002c9082610114565b506040805180820190915260048152630ae8aa8960e31b60208201526001906100559082610114565b506002805460ff1916601217905534801561006f57600080fd5b506101d2565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061009f57607f821691505b6020821081036100bf57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561010f57806000526020600020601f840160051c810160208510156100ec5750805b601f840160051c820191505b8181101561010c57600081556001016100f8565b50505b505050565b81516001600160401b0381111561012d5761012d610075565b6101418161013b845461008b565b846100c5565b6020601f821160018114610175576000831561015d5750848201515b600019600385901b1c1916600184901b17845561010c565b600084815260208120601f198516915b828110156101a55787850151825560209485019460019092019101610185565b50848210156101c35786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610986806101e16000396000f3fe6080604052600436106100c05760003560e01c8063313ce56711610074578063a9059cbb1161004e578063a9059cbb146101fa578063d0e30db01461021a578063dd62ed3e1461022257600080fd5b8063313ce5671461018c57806370a08231146101b857806395d89b41146101e557600080fd5b806318160ddd116100a557806318160ddd1461012f57806323b872dd1461014c5780632e1a7d4d1461016c57600080fd5b806306fdde03146100d4578063095ea7b3146100ff57600080fd5b366100cf576100cd61025a565b005b600080fd5b3480156100e057600080fd5b506100e96102b5565b6040516100f69190610745565b60405180910390f35b34801561010b57600080fd5b5061011f61011a3660046107da565b610343565b60405190151581526020016100f6565b34801561013b57600080fd5b50475b6040519081526020016100f6565b34801561015857600080fd5b5061011f610167366004610804565b6103bd565b34801561017857600080fd5b506100cd610187366004610841565b610647565b34801561019857600080fd5b506002546101a69060ff1681565b60405160ff90911681526020016100f6565b3480156101c457600080fd5b5061013e6101d336600461085a565b60036020526000908152604090205481565b3480156101f157600080fd5b506100e9610724565b34801561020657600080fd5b5061011f6102153660046107da565b610731565b6100cd61025a565b34801561022e57600080fd5b5061013e61023d366004610875565b600460209081526000928352604080842090915290825290205481565b33600090815260036020526040812080543492906102799084906108d7565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b600080546102c2906108ea565b80601f01602080910402602001604051908101604052809291908181526020018280546102ee906108ea565b801561033b5780601f106103105761010080835404028352916020019161033b565b820191906000526020600020905b81548152906001019060200180831161031e57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103ab9086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604081205482111561042b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841633148015906104a1575073ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156105605773ffffffffffffffffffffffffffffffffffffffff8416600090815260046020908152604080832033845290915290205482111561051a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b73ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091528120805484929061055a90849061093d565b90915550505b73ffffffffffffffffffffffffffffffffffffffff84166000908152600360205260408120805484929061059590849061093d565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040812080548492906105cf9084906108d7565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161063591815260200190565b60405180910390a35060019392505050565b3360009081526003602052604090205481111561069a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b33600090815260036020526040812080548392906106b990849061093d565b9091555050604051339082156108fc029083906000818181858888f193505050501580156106eb573d6000803e3d6000fd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b600180546102c2906108ea565b600061073e3384846103bd565b9392505050565b602081526000825180602084015260005b818110156107735760208186018101516040868401015201610756565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146107d557600080fd5b919050565b600080604083850312156107ed57600080fd5b6107f6836107b1565b946020939093013593505050565b60008060006060848603121561081957600080fd5b610822846107b1565b9250610830602085016107b1565b929592945050506040919091013590565b60006020828403121561085357600080fd5b5035919050565b60006020828403121561086c57600080fd5b61073e826107b1565b6000806040838503121561088857600080fd5b610891836107b1565b915061089f602084016107b1565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156103b7576103b76108a8565b600181811c908216806108fe57607f821691505b602082108103610937577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b818103818111156103b7576103b76108a856fea2646970667358221220b09c98b42d894b1b92a74ecc691587bcc17012fff7ef3bcfa6fe755f9b6255a564736f6c634300081a00336080604052348015600f57600080fd5b506106308061001f6000396000f3fe60806040526004361061002a5760003560e01c80635ac1e07014610033578063de43156e1461005357005b3661003157005b005b34801561003f57600080fd5b5061003161004e366004610128565b610073565b34801561005f57600080fd5b5061003161006e366004610193565b6100ad565b7ff5d4d6ed1657736cfd625a04e0585c5def81b280f6bc1b7ed4add114e79f7e9a816040516100a29190610299565b60405180910390a150565b606081156100c4576100c1828401846103bb565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e6100ef87806104b1565b6100ff60408a0160208b0161051d565b8960400135338660405161011896959493929190610538565b60405180910390a1505050505050565b60006020828403121561013a57600080fd5b813567ffffffffffffffff81111561015157600080fd5b82016080818503121561016357600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461018e57600080fd5b919050565b6000806000806000608086880312156101ab57600080fd5b853567ffffffffffffffff8111156101c257600080fd5b8601606081890312156101d457600080fd5b94506101e26020870161016a565b935060408601359250606086013567ffffffffffffffff81111561020557600080fd5b8601601f8101881361021657600080fd5b803567ffffffffffffffff81111561022d57600080fd5b88602082840101111561023f57600080fd5b959894975092955050506020019190565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6020815273ffffffffffffffffffffffffffffffffffffffff6102bb8361016a565b16602082015273ffffffffffffffffffffffffffffffffffffffff6102e26020840161016a565b1660408201526000604083013567ffffffffffffffff811680821461030657600080fd5b806060850152505060608301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261034357600080fd5b830160208101903567ffffffffffffffff81111561036057600080fd5b80360382131561036f57600080fd5b60808085015261038360a085018284610250565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156103cd57600080fd5b813567ffffffffffffffff8111156103e457600080fd5b8201601f810184136103f557600080fd5b803567ffffffffffffffff81111561040f5761040f61038c565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561047b5761047b61038c565b60405281815282820160200186101561049357600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126104e657600080fd5b83018035915067ffffffffffffffff82111561050157600080fd5b60200191503681900382131561051657600080fd5b9250929050565b60006020828403121561052f57600080fd5b6101638261016a565b60a08152600061054c60a08301888a610250565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b818110156105b65760208187018101518483018201520161059a565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168301019250505097965050505050505056fea264697066735822122023cfd100c0b5ffb64d0a74f749750529178249712786f8146b8cbbfeedf3ca8264736f6c634300081a003360c060405234801561001057600080fd5b5060405161106f38038061106f83398101604081905261002f916100db565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006357604051632b2add3d60e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0385811691909117909155828116608052811660a0526040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac590600090a150505061011e565b80516001600160a01b03811681146100d657600080fd5b919050565b6000806000606084860312156100f057600080fd5b6100f9846100bf565b9250610107602085016100bf565b9150610115604085016100bf565b90509250925092565b60805160a051610f2561014a60003960006101e50152600081816102b9015261045b0152610f256000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806397770dff11610097578063c63585cc11610066578063c63585cc14610273578063d7fd7afb14610286578063d936a012146102b4578063ee2815ba146102db57600080fd5b806397770dff1461021a578063a7cb05071461022d578063c39aca3714610240578063c62178ac1461025357600080fd5b8063513a9c05116100d3578063513a9c051461018a578063569541b9146101c0578063842da36d146101e057806391dd645f1461020757600080fd5b80630be15547146100fa5780631f0e251b1461015a5780633ce4a5bc1461016f575b600080fd5b610130610108366004610bd1565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61016d610168366004610c13565b6102ee565b005b61013073735b14bb79463307aacbed86daf3322b1e6226ab81565b610130610198366004610bd1565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101309073ffffffffffffffffffffffffffffffffffffffff1681565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d610215366004610c35565b610402565b61016d610228366004610c13565b610526565b61016d61023b366004610c61565b610633565b61016d61024e366004610c83565b6106ce565b6004546101309073ffffffffffffffffffffffffffffffffffffffff1681565b610130610281366004610d53565b6108cd565b6102a6610294366004610bd1565b60006020819052908152604090205481565b604051908152602001610151565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d6102e9366004610c35565b610a02565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461033b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610388576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461044f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354600090610497907f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff16846108cd565b60008481526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251878152918201529192507f0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e910160405180910390a1505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610573576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166105c0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e906020016103f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610680576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461071b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831673735b14bb79463307aacbed86daf3322b1e6226ab1480610768575073ffffffffffffffffffffffffffffffffffffffff831630145b1561079f576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690528616906347e7ef24906044016020604051808303816000875af1158015610814573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108389190610d96565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063de43156e906108939089908990899088908890600401610e01565b600060405180830381600087803b1580156108ad57600080fd5b505af11580156108c1573d6000803e3d6000fd5b50505050505050505050565b60008060006108dc8585610ad3565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b16603482015291935091508690604801604051602081830303815290604052805190602001206040516020016109c29291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a4f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d91016106c2565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b3b576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610b75578284610b78565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610bca576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b600060208284031215610be357600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c0e57600080fd5b919050565b600060208284031215610c2557600080fd5b610c2e82610bea565b9392505050565b60008060408385031215610c4857600080fd5b82359150610c5860208401610bea565b90509250929050565b60008060408385031215610c7457600080fd5b50508035926020909101359150565b60008060008060008060a08789031215610c9c57600080fd5b863567ffffffffffffffff811115610cb357600080fd5b87016060818a031215610cc557600080fd5b9550610cd360208801610bea565b945060408701359350610ce860608801610bea565b9250608087013567ffffffffffffffff811115610d0457600080fd5b8701601f81018913610d1557600080fd5b803567ffffffffffffffff811115610d2c57600080fd5b896020828401011115610d3e57600080fd5b60208201935080925050509295509295509295565b600080600060608486031215610d6857600080fd5b610d7184610bea565b9250610d7f60208501610bea565b9150610d8d60408501610bea565b90509250925092565b600060208284031215610da857600080fd5b81518015158114610c2e57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60808152600086357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018112610e3957600080fd5b870160208101903567ffffffffffffffff811115610e5657600080fd5b803603821315610e6557600080fd5b60606080850152610e7a60e085018284610db8565b91505073ffffffffffffffffffffffffffffffffffffffff610e9e60208a01610bea565b1660a0840152604088013560c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401528281036060840152610ee3818587610db8565b9897505050505050505056fea26469706673582212201b1a136275cbee5409c60f315f9fc4e6eecd668368ddacf6f2cfdcef88de5c5a64736f6c634300081a003360c060405234801561001057600080fd5b5060405161207238038061207283398101604081905261002f916101f0565b6001600160a01b038216158061004c57506001600160a01b038116155b1561006a5760405163d92e233d60e01b815260040160405180910390fd5b60066100768982610342565b5060076100838882610342565b506008805460ff191660ff881617905560808590528360028111156100aa576100aa610400565b60a08160028111156100be576100be610400565b905250600192909255600080546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506104169350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261013357600080fd5b81516001600160401b0381111561014c5761014c61010c565b604051601f8201601f19908116603f011681016001600160401b038111828210171561017a5761017a61010c565b60405281815283820160200185101561019257600080fd5b60005b828110156101b157602081860181015183830182015201610195565b506000918101602001919091529392505050565b8051600381106101d457600080fd5b919050565b80516001600160a01b03811681146101d457600080fd5b600080600080600080600080610100898b03121561020d57600080fd5b88516001600160401b0381111561022357600080fd5b61022f8b828c01610122565b60208b015190995090506001600160401b0381111561024d57600080fd5b6102598b828c01610122565b975050604089015160ff8116811461027057600080fd5b60608a0151909650945061028660808a016101c5565b60a08a0151909450925061029c60c08a016101d9565b91506102aa60e08a016101d9565b90509295985092959890939650565b600181811c908216806102cd57607f821691505b6020821081036102ed57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561033d57806000526020600020601f840160051c8101602085101561031a5750805b601f840160051c820191505b8181101561033a5760008155600101610326565b50505b505050565b81516001600160401b0381111561035b5761035b61010c565b61036f8161036984546102b9565b846102f3565b6020601f8211600181146103a3576000831561038b5750848201515b600019600385901b1c1916600184901b17845561033a565b600084815260208120601f198516915b828110156103d357878501518255602094850194600190920191016103b3565b50848210156103f15786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a051611c1b61045760003960006103440152600081816102f001528181610bdc01528181610ce201528181610efe01526110040152611c1b6000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806395d89b41116100f9578063ccc7759911610097578063eddeb12311610071578063eddeb12314610461578063f2441b3214610474578063f687d12a14610494578063fc5fecd5146104a757600080fd5b8063ccc77599146103d4578063d9eeebed146103e7578063dd62ed3e1461041b57600080fd5b8063b84c8246116100d3578063b84c824614610386578063c47f00271461039b578063c7012626146103ae578063c835d7cc146103c157600080fd5b806395d89b4114610337578063a3413d031461033f578063a9059cbb1461037357600080fd5b80633ce4a5bc116101665780634d8943bb116101405780634d8943bb146102ac57806370a08231146102b557806385e1f4d0146102eb5780638b851b951461031257600080fd5b80633ce4a5bc1461024657806342966c681461028657806347e7ef241461029957600080fd5b806318160ddd1161019757806318160ddd1461021657806323b872dd1461021e578063313ce5671461023157600080fd5b806306fdde03146101be578063091d2788146101dc578063095ea7b3146101f3575b600080fd5b6101c66104ba565b6040516101d39190611648565b60405180910390f35b6101e560015481565b6040519081526020016101d3565b610206610201366004611687565b61054c565b60405190151581526020016101d3565b6005546101e5565b61020661022c3660046116b3565b610563565b60085460405160ff90911681526020016101d3565b61026173735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6102066102943660046116f4565b6105fa565b6102066102a7366004611687565b61060e565b6101e560025481565b6101e56102c336600461170d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101e57f000000000000000000000000000000000000000000000000000000000000000081565b60085461026190610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101c6610767565b6103667f000000000000000000000000000000000000000000000000000000000000000081565b6040516101d3919061172a565b610206610381366004611687565b610776565b610399610394366004611832565b610783565b005b6103996103a9366004611832565b6107e0565b6102066103bc366004611883565b610839565b6103996103cf36600461170d565b610988565b6103996103e236600461170d565b610a9c565b6103ef610bb0565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101d3565b6101e56104293660046118dc565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b61039961046f3660046116f4565b610dce565b6000546102619073ffffffffffffffffffffffffffffffffffffffff1681565b6103996104a23660046116f4565b610e50565b6103ef6104b53660046116f4565b610ed2565b6060600680546104c990611915565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590611915565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b5050505050905090565b60006105593384846110ee565b5060015b92915050565b60006105708484846111f7565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054828110156105db576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ef85336105ea8685611997565b6110ee565b506001949350505050565b600061060633836113b2565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab1480159061064c575060005473ffffffffffffffffffffffffffffffffffffffff163314155b80156106755750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b156106ac576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106b683836114f4565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107569186906119aa565b60405180910390a250600192915050565b6060600780546104c990611915565b60006105593384846111f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107d0576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107dc8282611a1b565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461082d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107dc8282611a1b565b6000806000610846610bb0565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156108d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fc9190611b34565b610932576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61093c33856113b2565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161097591899189918791611b56565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109d5576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a22576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610ae9576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b36576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c679190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610cb6576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d699190611ba2565b905080600003610da5576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610db89190611bbb565b610dc29190611bd2565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e1b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a91565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e9d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f899190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610fd8576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015611067573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108b9190611ba2565b9050806000036110c7576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000906110d78784611bbb565b6110e19190611bd2565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661113b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611188576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611244576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611291576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040902054818110156112f1576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112fb8282611997565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260036020526040808220939093559085168152908120805484929061133e908490611bd2565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113a491815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113ff576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260409020548181101561145f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114698282611997565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812091909155600580548492906114a4908490611997565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111ea565b73ffffffffffffffffffffffffffffffffffffffff8216611541576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008282546115539190611bd2565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805483929061158d908490611bd2565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561160a576020818501810151868301820152016115ee565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061165b60208301846115e4565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461168457600080fd5b50565b6000806040838503121561169a57600080fd5b82356116a581611662565b946020939093013593505050565b6000806000606084860312156116c857600080fd5b83356116d381611662565b925060208401356116e381611662565b929592945050506040919091013590565b60006020828403121561170657600080fd5b5035919050565b60006020828403121561171f57600080fd5b813561165b81611662565b6020810160038310611765577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff8411156117b5576117b561176b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156118025761180261176b565b60405283815290508082840185101561181a57600080fd5b83836020830137600060208583010152509392505050565b60006020828403121561184457600080fd5b813567ffffffffffffffff81111561185b57600080fd5b8201601f8101841361186c57600080fd5b61187b8482356020840161179a565b949350505050565b6000806040838503121561189657600080fd5b823567ffffffffffffffff8111156118ad57600080fd5b8301601f810185136118be57600080fd5b6118cd8582356020840161179a565b95602094909401359450505050565b600080604083850312156118ef57600080fd5b82356118fa81611662565b9150602083013561190a81611662565b809150509250929050565b600181811c9082168061192957607f821691505b602082108103611962577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561055d5761055d611968565b6040815260006119bd60408301856115e4565b90508260208301529392505050565b601f821115611a1657806000526020600020601f840160051c810160208510156119f35750805b601f840160051c820191505b81811015611a1357600081556001016119ff565b50505b505050565b815167ffffffffffffffff811115611a3557611a3561176b565b611a4981611a438454611915565b846119cc565b6020601f821160018114611a9b5760008315611a655750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611a13565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611ae95787850151825560209485019460019092019101611ac9565b5084821015611b2557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b600060208284031215611b4657600080fd5b8151801515811461165b57600080fd5b608081526000611b6960808301876115e4565b6020830195909552506040810192909252606090910152919050565b600060208284031215611b9757600080fd5b815161165b81611662565b600060208284031215611bb457600080fd5b5051919050565b808202811582820484141761055d5761055d611968565b8082018082111561055d5761055d61196856fea2646970667358221220d6ba834f25782689ed13bffb6ac9ff2c8d3b5342c94a515aea8197a76070ad3f64736f6c634300081a0033a264697066735822122038f5e6e956cf2c96bf1463494f8fe65913f394a1a4ddb9a0a0db2420321a134a64736f6c634300081a0033",
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

// FilterContextDataRevert is a free log retrieval operation binding the contract event 0xf5d4d6ed1657736cfd625a04e0585c5def81b280f6bc1b7ed4add114e79f7e9a.
//
// Solidity: event ContextDataRevert((address,address,uint64,bytes) revertContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterContextDataRevert(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestContextDataRevertIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestContextDataRevertIterator{contract: _GatewayZEVMOutboundTest.contract, event: "ContextDataRevert", logs: logs, sub: sub}, nil
}

// WatchContextDataRevert is a free log subscription operation binding the contract event 0xf5d4d6ed1657736cfd625a04e0585c5def81b280f6bc1b7ed4add114e79f7e9a.
//
// Solidity: event ContextDataRevert((address,address,uint64,bytes) revertContext)
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

// ParseContextDataRevert is a log parse operation binding the contract event 0xf5d4d6ed1657736cfd625a04e0585c5def81b280f6bc1b7ed4add114e79f7e9a.
//
// Solidity: event ContextDataRevert((address,address,uint64,bytes) revertContext)
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
