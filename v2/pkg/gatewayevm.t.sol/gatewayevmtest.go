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

// GatewayEVMTestMetaData contains all meta data concerning the GatewayEVMTest contract.
var GatewayEVMTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ASSET_HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_UPDATER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfDestinationIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevert\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertFailsIfSenderIsNotTSS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteWithERC20FailsIfNotCustodyOrConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNoParams\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNoParamsTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNonPayable\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceivePayable\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertWithERC20FailsIfNotCustodyOrConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetConnectorFailsIfSenderIsNotAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetConnectorFailsIfSet\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetConnectorFailsIfZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetCustodyFailsIfSenderIsNotAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetCustodyFailsIfSet\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetCustodyFailsIfZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgrade\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfSenderIsNotTSSUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedCustodyTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5062010afe806200003e6000396000f3fe608060405234801561001057600080fd5b50600436106102ad5760003560e01c806385f438c11161017b578063ce5871e1116100d8578063e6afc7901161008c578063fa18c09b11610071578063fa18c09b146104d3578063fa7626d4146104db578063fe7bdbb2146104e857600080fd5b8063e6afc790146104c3578063f68bd1c0146104cb57600080fd5b8063dd51e82f116100bd578063dd51e82f1461048c578063e20c9f7114610494578063e63ab1e91461049c57600080fd5b8063ce5871e11461047c578063cebad2a61461048457600080fd5b8063b0464fdc1161012f578063b5508aa911610114578063b5508aa914610454578063ba414fa61461045c578063ccf206161461047457600080fd5b8063b0464fdc14610444578063b124dbf51461044c57600080fd5b8063a217fddf11610160578063a217fddf1461040d578063a56f7a4b14610415578063a783c7891461041d57600080fd5b806385f438c1146103d1578063916a17c6146103f857600080fd5b80633f7286f4116102295780635d62c860116101dd5780637d7f772a116101c25780637d7f772a146103ac5780637ebf744f146103b457806385226c81146103bc57600080fd5b80635d62c8601461037057806366d9a9a01461039757600080fd5b80634df42da11161020e5780634df42da11461035857806351010e491461036057806352ff59391461036857600080fd5b80633f7286f41461034857806344671b941461035057600080fd5b80631226c655116102805780631ed7831c116102655780631ed7831c146103165780632ade38801461032b5780633e5e3c231461034057600080fd5b80631226c655146103065780631855c3371461030e57600080fd5b806304b974f9146102b2578063070f2ad0146102bc5780630a9254e4146102c45780630d391704146102cc575b600080fd5b6102ba6104f0565b005b6102ba6106fd565b6102ba61091d565b6102f37f1ddc5e4a2f140581e5d9d5f758b2c0fbdd814c4017a02afc4b55cabcf2c1f10b81565b6040519081526020015b60405180910390f35b6102ba61131b565b6102ba61148c565b61031e611601565b6040516102fd9190618e61565b610333611663565b6040516102fd9190618efd565b61031e6117a5565b61031e611805565b6102ba611865565b6102ba611be4565b6102ba611d55565b6102ba611f31565b6102f37f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b981565b61039f6121c0565b6040516102fd9190619063565b6102ba612342565b6102ba6126a5565b6103c461281a565b6040516102fd919061915d565b6102f37f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6104006128ea565b6040516102fd9190619170565b6102f3600081565b6102ba6129e5565b6102f37f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b610400612b37565b6102ba612c32565b6103c4612ed9565b610464612fa9565b60405190151581526020016102fd565b6102ba61307d565b6102ba6131cf565b6102ba613340565b6102ba613564565b61031e613d59565b6102f37f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6102ba613db9565b6102ba614027565b6102ba61441c565b601f546104649060ff1681565b6102ba61477e565b6040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed7016900000000000000000000000000000000000000000000000000000000179052602854925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa792610597926001600160a01b031691016001600160a01b0391909116815260200190565b600060405180830381600087803b1580156105b157600080fd5b505af11580156105c5573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561064e57600080fd5b505af1158015610662573d6000803e3d6000fd5b50506020546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250631cff79cd91506106b2906000908590600401619207565b6000604051808303816000875af11580156106d1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106f99190810190619311565b5050565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561076f57600080fd5b505af1158015610783573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f1ddc5e4a2f140581e5d9d5f758b2c0fbdd814c4017a02afc4b55cabcf2c1f10b60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061086e9190600401619346565b600060405180830381600087803b15801561088857600080fd5b505af115801561089c573d6000803e3d6000fd5b50506020546026546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063950837aa91506024015b600060405180830381600087803b15801561090357600080fd5b505af1158015610917573d6000803e3d6000fd5b50505050565b602680547fffffffffffffffffffffffff000000000000000000000000000000000000000090811630179091556027805482166112341790556028805490911661567817905560405161096f90618d62565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156109f4573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556028546040519116908190610a3d90618d70565b6001600160a01b03928316815291166020820152604001604051809103906000f080158015610a70573d6000803e3d6000fd5b50602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260285460265492519085166024820152604481019390935292166064820152610b5f919060840160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b00000000000000000000000000000000000000000000000000000000179052614bf3565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff00000000000000000000000000000000000000009091168117909155602854602654604051929391821692911690610beb90618d7e565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610c27573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556020546025546028546026546040519385169492831693918316921690610c8290618d8c565b6001600160a01b039485168152928416602084015290831660408301529091166060820152608001604051809103906000f080158015610cc6573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392831617905560285460405163ca669fa760e01b815291166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015610d4b57600080fd5b505af1158015610d5f573d6000803e3d6000fd5b50506025546028546023546040517f15d57fd40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529082166024820152911692506315d57fd49150604401600060405180830381600087803b158015610dd057600080fd5b505af1158015610de4573d6000803e3d6000fd5b50505050604051610df490618d9a565b604051809103906000f080158015610e10573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556028546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b158015610ebc57600080fd5b505af1158015610ed0573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015610f4657600080fd5b505af1158015610f5a573d6000803e3d6000fd5b50506020546022546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f9150602401600060405180830381600087803b158015610fc057600080fd5b505af1158015610fd4573d6000803e3d6000fd5b50506020546023546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef9150602401600060405180830381600087803b15801561103a57600080fd5b505af115801561104e573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156110b057600080fd5b505af11580156110c4573d6000803e3d6000fd5b5050602480546026546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f4240938101939093521692506340c10f199150604401600060405180830381600087803b15801561113557600080fd5b505af1158015611149573d6000803e3d6000fd5b5050602480546022546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526207a1209381019390935216925063a9059cbb91506044016020604051808303816000875af11580156111bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111e39190619359565b506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b15801561126457600080fd5b505af1158015611278573d6000803e3d6000fd5b5050604080516060810182526024546001600160a01b0390811682526001602080840191825284519081018552600081529383018490528251602d8054925167ffffffffffffffff1674010000000000000000000000000000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090931691909316171781559093509150602e90611316908261940f565b505050565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561138d57600080fd5b505af11580156113a1573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561142a57600080fd5b505af115801561143e573d6000803e3d6000fd5b50506020546040517fae7a3a6f000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b03909116925063ae7a3a6f91506024016108e9565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156114fe57600080fd5b505af1158015611512573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fb337f378000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b600060405180830381600087803b15801561159c57600080fd5b505af11580156115b0573d6000803e3d6000fd5b50506020546022546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f91506024016108e9565b6060601680548060200260200160405190810160405280929190818152602001828054801561165957602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161163b575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020016000905b8282101561179c57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156117855783829060005260206000200180546116f89061937b565b80601f01602080910402602001604051908101604052809291908181526020018280546117249061937b565b80156117715780601f1061174657610100808354040283529160200191611771565b820191906000526020600020905b81548152906001019060200180831161175457829003601f168201915b5050505050815260200190600101906116d9565b505050508152505081526020019060010190611687565b50505050905090565b60606018805480602002602001604051908101604052809291908181526020018280548015611659576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161163b575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015611659576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161163b575050505050905090565b6040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed701690000000000000000000000000000000000000000000000000000000017905260215492517ff30c7ba30000000000000000000000000000000000000000000000000000000081529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba39261191b926001600160a01b0316916000918791016194ce565b600060405180830381600087803b15801561193557600080fd5b505af1158015611949573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156119db57600080fd5b505af11580156119ef573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0935001905060405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611ac157600080fd5b505af1158015611ad5573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f9150611b1b9060009085906194f6565b60405180910390a260285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611b7c57600080fd5b505af1158015611b90573d6000803e3d6000fd5b50506020546021546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd93506106b292909116908590600401619207565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015611c5657600080fd5b505af1158015611c6a573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015611cf357600080fd5b505af1158015611d07573d6000803e3d6000fd5b50506020546040517f950837aa000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b03909116925063950837aa91506024016108e9565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152670de0b6b3a76400009060009060250160408051808303601f190181529082905260285463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611df557600080fd5b505af1158015611e09573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015611e9257600080fd5b505af1158015611ea6573d6000803e3d6000fd5b50506020546040517ff7ad60db0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063f7ad60db91508490611efb906000908690602d906004016195e4565b6000604051808303818588803b158015611f1457600080fd5b505af1158015611f28573d6000803e3d6000fd5b50505050505050565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015611fa357600080fd5b505af1158015611fb7573d6000803e3d6000fd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561204957600080fd5b505af115801561205d573d6000803e3d6000fd5b50506028546040516001600160a01b0390911681527f7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a059250602001905060405180910390a16020546028546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169063950837aa90602401600060405180830381600087803b15801561210457600080fd5b505af1158015612118573d6000803e3d6000fd5b505060208054604080517f5b1125910000000000000000000000000000000000000000000000000000000081529051600095506001600160a01b039092169350635b1125919260048083019391928290030181865afa15801561217f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121a39190619618565b6028549091506121bd9082906001600160a01b0316614c12565b50565b6060601b805480602002602001604051908101604052809291908181526020016000905b8282101561179c57838290600052602060002090600202016040518060400160405290816000820180546122179061937b565b80601f01602080910402602001604051908101604052809291908181526020018280546122439061937b565b80156122905780601f1061226557610100808354040283529160200191612290565b820191906000526020600020905b81548152906001019060200180831161227357829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561232a57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116122d75790505b505050505081525050815260200190600101906121e4565b604080516001808252818301909252600091816020015b60608152602001906001900390816123595790505090506040518060400160405280600f81526020017f48656c6c6f2c20466f756e647279210000000000000000000000000000000000815250816000815181106123b9576123b9619641565b6020908102919091010152604080516001808252818301909252600091816020016020820280368337019050509050602a816000815181106123fd576123fd619641565b6020908102919091010152604051600190600090612423908590859085906024016196a2565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff05b6abf00000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156124d657600080fd5b505af11580156124ea573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506125d59190600401619346565b600060405180830381600087803b1580156125ef57600080fd5b505af1158015612603573d6000803e3d6000fd5b50506020546021546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd935061265792909116908590600401619207565b6000604051808303816000875af1158015612676573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261269e9190810190619311565b5050505050565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561271757600080fd5b505af115801561272b573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527f0c8dc016000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b600060405180830381600087803b1580156127b557600080fd5b505af11580156127c9573d6000803e3d6000fd5b50506020546023546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef91506024016108e9565b6060601a805480602002602001604051908101604052809291908181526020016000905b8282101561179c57838290600052602060002001805461285d9061937b565b80601f01602080910402602001604051908101604052809291908181526020018280546128899061937b565b80156128d65780601f106128ab576101008083540402835291602001916128d6565b820191906000526020600020905b8154815290600101906020018083116128b957829003601f168201915b50505050508152602001906001019061283e565b6060601d805480602002602001604051908101604052809291908181526020016000905b8282101561179c5760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156129cd57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001906004019060208260030104928301926001038202915080841161297a5790505b5050505050815250508152602001906001019061290e565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015612a5757600080fd5b505af1158015612a6b573d6000803e3d6000fd5b5050602854604080516001600160a01b039092166024830152600060448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506115829190600401619346565b6060601c805480602002602001604051908101604052809291908181526020016000905b8282101561179c5760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015612c1a57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411612bc75790505b50505050508152505081526020019060010190612b5b565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612d1657600080fd5b505af1158015612d2a573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b960448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250612e159190600401619346565b600060405180830381600087803b158015612e2f57600080fd5b505af1158015612e43573d6000803e3d6000fd5b50506020546024546027546040517fd0b492c30000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d0b492c39450612ea393928316929091169087908790602d906004016196da565b600060405180830381600087803b158015612ebd57600080fd5b505af1158015612ed1573d6000803e3d6000fd5b505050505050565b60606019805480602002602001604051908101604052809291908181526020016000905b8282101561179c578382906000526020600020018054612f1c9061937b565b80601f0160208091040260200160405190810160405280929190818152602001828054612f489061937b565b8015612f955780601f10612f6a57610100808354040283529160200191612f95565b820191906000526020600020905b815481529060010190602001808311612f7857829003601f168201915b505050505081526020019060010190612efd565b60085460009060ff1615612fc1575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015613052573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613076919061972f565b1415905090565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b1580156130ef57600080fd5b505af1158015613103573d6000803e3d6000fd5b5050602854604080516001600160a01b039092166024830152600060448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061279b9190600401619346565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561324157600080fd5b505af1158015613255573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156132de57600080fd5b505af11580156132f2573d6000803e3d6000fd5b50506020546040517f10188aef000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b0390911692506310188aef91506024016108e9565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152670de0b6b3a76400009060009060250160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156133e057600080fd5b505af11580156133f4573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506134df9190600401619346565b600060405180830381600087803b1580156134f957600080fd5b505af115801561350d573d6000803e3d6000fd5b50506020546021546040517ff7ad60db0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063f7ad60db93508692611efb9216908690602d906004016195e4565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156135bd57600080fd5b505af11580156135d1573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506136bc9190600401619346565b600060405180830381600087803b1580156136d657600080fd5b505af11580156136ea573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561373e57600080fd5b505af1158015613752573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156137af57600080fd5b505af11580156137c3573d6000803e3d6000fd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506138ae9190600401619346565b600060405180830381600087803b1580156138c857600080fd5b505af11580156138dc573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561393057600080fd5b505af1158015613944573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156139a157600080fd5b505af11580156139b5573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b158015613a0957600080fd5b505af1158015613a1d573d6000803e3d6000fd5b5050604080516004808252602480830184526020830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed701690000000000000000000000000000000000000000000000000000000017905292517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd93c06650000000000000000000000000000000000000000000000000000000091810191909152909350737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09101600060405180830381600087803b158015613aff57600080fd5b505af1158015613b13573d6000803e3d6000fd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015613b7057600080fd5b505af1158015613b84573d6000803e3d6000fd5b50506020546021546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd9350613bd892909116908590600401619207565b6000604051808303816000875af1158015613bf7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613c1f9190810190619311565b5060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613c7957600080fd5b505af1158015613c8d573d6000803e3d6000fd5b50505050602060009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015613ce157600080fd5b505af1158015613cf5573d6000803e3d6000fd5b50506021546040517ff30c7ba3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f30c7ba3925061191b916001600160a01b03169060009086906004016194ce565b60606015805480602002602001604051908101604052809291908181526020018280548015611659576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161163b575050505050905090565b60248054602754604051620186a09381018490526001600160a01b0392831660448201529116606482015260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613e9d57600080fd5b505af1158015613eb1573d6000803e3d6000fd5b5050602654604080516001600160a01b0390921660248301527f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b960448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250613f9c9190600401619346565b600060405180830381600087803b158015613fb657600080fd5b505af1158015613fca573d6000803e3d6000fd5b50506020546024546027546040517f5131ab590000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550635131ab599450612ea393928316929091169087908790600401619748565b604080516001808252818301909252600091816020015b606081526020019060019003908161403e5790505090506040518060400160405280600f81526020017f48656c6c6f2c20466f756e6472792100000000000000000000000000000000008152508160008151811061409e5761409e619641565b6020908102919091010152604080516001808252818301909252600091816020016020820280368337019050509050602a816000815181106140e2576140e2619641565b6020908102919091010152604051600190600090614108908590859085906024016196a2565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff05b6abf0000000000000000000000000000000000000000000000000000000017905260215490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba3916141c5916001600160a01b03919091169060009086906004016194ce565b600060405180830381600087803b1580156141df57600080fd5b505af11580156141f3573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561428557600080fd5b505af1158015614299573d6000803e3d6000fd5b50506020546040517f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b14693506142e092506001600160a01b039091169087908790879061977f565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561437657600080fd5b505af115801561438a573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f91506143d09060009085906194f6565b60405180910390a260285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024016125d5565b604080517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152815160058183030181526025909101909152602154670de0b6b3a764000091906001600160a01b03163161447b816000614c9b565b6021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561450957600080fd5b505af115801561451d573d6000803e3d6000fd5b50506020546040517f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e935061456192506001600160a01b0390911690602d906197c7565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156145f757600080fd5b505af115801561460b573d6000803e3d6000fd5b5050602154604051600093506001600160a01b0390911691507f1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b14369061465e90670de0b6b3a7640000908790602d906197e9565b60405180910390a360285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156146bf57600080fd5b505af11580156146d3573d6000803e3d6000fd5b50506020546021546040517ff7ad60db0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063f7ad60db9350879261472a9216908790602d906004016195e4565b6000604051808303818588803b15801561474357600080fd5b505af1158015614757573d6000803e3d6000fd5b50506021546001600160a01b03163192506109179150829050670de0b6b3a7640000614c9b565b60408051808201909152600f81527f48656c6c6f2c20466f756e6472792100000000000000000000000000000000006020820152602154602a90600190670de0b6b3a7640000906147db906000906001600160a01b031631614c9b565b60008484846040516024016147f293929190619802565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905260215490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba3916148b6916001600160a01b039190911690670de0b6b3a76400009086906004016194ce565b600060405180830381600087803b1580156148d057600080fd5b505af11580156148e4573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561497657600080fd5b505af115801561498a573d6000803e3d6000fd5b50506020546040517f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa93506149d392506001600160a01b0390911690859089908990899061982c565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015614a6957600080fd5b505af1158015614a7d573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f9150614aca90670de0b6b3a76400009085906194f6565b60405180910390a260285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614b2b57600080fd5b505af1158015614b3f573d6000803e3d6000fd5b50506020546021546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd93508692614b939216908690600401619207565b60006040518083038185885af1158015614bb1573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052614bda9190810190619311565b5060215461269e9083906001600160a01b031631614c9b565b6000614bfd618da8565b614c08848483614cf3565b9150505b92915050565b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f6906044015b60006040518083038186803b158015614c8757600080fd5b505afa158015612ed1573d6000803e3d6000fd5b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c5490604401614c6f565b600080614d008584614d6e565b9050614d636040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f78790000008152508286604051602001614d4e929190619207565b60405160208183030381529060405285614d7a565b9150505b9392505050565b6000614d678383614da8565b60c08101515160009015614d9e57614d9784848460c00151614dc3565b9050614d67565b614d978484614f69565b6000614db48383615054565b614d6783836020015184614d7a565b600080614dce615060565b90506000614ddc8683615133565b90506000614df382606001518360200151856155d9565b90506000614e03838389896157eb565b90506000614e1082616668565b602081015181519192509060030b15614e8357898260400151604051602001614e3a92919061986d565b60408051601f19818403018152908290527f08c379a0000000000000000000000000000000000000000000000000000000008252614e7a91600401619346565b60405180910390fd5b6000614ec66040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001616837565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d90614f19908490600401619346565b602060405180830381865afa158015614f36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614f5a9190619618565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc92590614fbe908790600401619346565b600060405180830381865afa158015614fdb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526150039190810190619311565b90506000615031828560405160200161501d9291906198ee565b604051602081830303815290604052616a37565b90506001600160a01b038116614c08578484604051602001614e3a92919061991d565b6106f982826000616a4a565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c906150e79084906004016199c8565b600060405180830381865afa158015615104573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261512c9190810190619a0f565b9250505090565b6151656040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d90506151b06040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6151b985616b4d565b602082015260006151c986616f32565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa15801561520b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526152339190810190619a0f565b8683856020015160405160200161524d9493929190619a58565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb11906152a5908590600401619346565b600060405180830381865afa1580156152c2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526152ea9190810190619a0f565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f690615332908490600401619b5c565b602060405180830381865afa15801561534f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906153739190619359565b6153885781604051602001614e3a9190619bae565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906153cd908490600401619c40565b600060405180830381865afa1580156153ea573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526154129190810190619a0f565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f690615459908490600401619c92565b602060405180830381865afa158015615476573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061549a9190619359565b1561552f576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906154e4908490600401619c92565b600060405180830381865afa158015615501573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526155299190810190619a0f565b60408501525b846001600160a01b03166349c4fac88286600001516040516020016155549190619ce4565b6040516020818303038152906040526040518363ffffffff1660e01b8152600401615580929190619d50565b600060405180830381865afa15801561559d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526155c59190810190619a0f565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b60608152602001906001900390816155f55790505090506040518060400160405280600481526020017f67726570000000000000000000000000000000000000000000000000000000008152508160008151811061565557615655619641565b60200260200101819052506040518060400160405280600381526020017f2d726c0000000000000000000000000000000000000000000000000000000000815250816001815181106156a9576156a9619641565b6020026020010181905250846040516020016156c59190619d75565b604051602081830303815290604052816002815181106156e7576156e7619641565b6020026020010181905250826040516020016157039190619de1565b6040516020818303038152906040528160038151811061572557615725619641565b6020026020010181905250600061573b82616668565b602080820151604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000081850190815282518084018452600080825290860152825180840190935290518252928101929092529192506157cc90604080518082018252600080825260209182015281518083019092528451825280850190820152906171b5565b6157e15785604051602001614e3a9190619e22565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d901561583b565b511590565b6159af578260200151156158f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401614e7a565b8260c00151156159af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401614e7a565b6040805160ff8082526120008201909252600091816020015b60608152602001906001900390816159c857905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280615a2390619ee2565b935060ff1681518110615a3857615a38619641565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e3700000000000000000000000000000000000000815250604051602001615a899190619f01565b604051602081830303815290604052828280615aa490619ee2565b935060ff1681518110615ab957615ab9619641565b60200260200101819052506040518060400160405280600681526020017f6465706c6f790000000000000000000000000000000000000000000000000000815250828280615b0690619ee2565b935060ff1681518110615b1b57615b1b619641565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d65000000000000000000000000000000000000815250828280615b6890619ee2565b935060ff1681518110615b7d57615b7d619641565b60200260200101819052508760200151828280615b9990619ee2565b935060ff1681518110615bae57615bae619641565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e747261637450617468000000000000000000000000000000000000815250828280615bfb90619ee2565b935060ff1681518110615c1057615c10619641565b602090810291909101015287518282615c2881619ee2565b935060ff1681518110615c3d57615c3d619641565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e49640000000000000000000000000000000000000000000000815250828280615c8a90619ee2565b935060ff1681518110615c9f57615c9f619641565b6020026020010181905250615cb346617216565b8282615cbe81619ee2565b935060ff1681518110615cd357615cd3619641565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c650000000000000000000000000000000000815250828280615d2090619ee2565b935060ff1681518110615d3557615d35619641565b602002602001018190525086828280615d4d90619ee2565b935060ff1681518110615d6257615d62619641565b6020908102919091010152855115615e895760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f6465000000000000000000000060208201528282615db381619ee2565b935060ff1681518110615dc857615dc8619641565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d90615e18908990600401619346565b600060405180830381865afa158015615e35573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615e5d9190810190619a0f565b8282615e6881619ee2565b935060ff1681518110615e7d57615e7d619641565b60200260200101819052505b846020015115615f595760408051808201909152601281527f2d2d766572696679536f75726365436f6465000000000000000000000000000060208201528282615ed281619ee2565b935060ff1681518110615ee757615ee7619641565b60200260200101819052506040518060400160405280600581526020017f66616c7365000000000000000000000000000000000000000000000000000000815250828280615f3490619ee2565b935060ff1681518110615f4957615f49619641565b6020026020010181905250616120565b615f916158368660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6160245760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282615fd481619ee2565b935060ff1681518110615fe957615fe9619641565b60200260200101819052508460a001516040516020016160099190619d75565b604051602081830303815290604052828280615f3490619ee2565b8460c0015115801561606757506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261606590511590565b155b156161205760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826160ab81619ee2565b935060ff16815181106160c0576160c0619641565b60200260200101819052506160d4886172b6565b6040516020016160e49190619d75565b6040516020818303038152906040528282806160ff90619ee2565b935060ff168151811061611457616114619641565b60200260200101819052505b6040808601518151808301835260008082526020918201528251808401909352815183529081019082015261615490511590565b6161e95760408051808201909152600b81527f2d2d72656c6179657249640000000000000000000000000000000000000000006020820152828261619781619ee2565b935060ff16815181106161ac576161ac619641565b602002602001018190525084604001518282806161c890619ee2565b935060ff16815181106161dd576161dd619641565b60200260200101819052505b60608501511561630a5760408051808201909152600681527f2d2d73616c7400000000000000000000000000000000000000000000000000006020820152828261623281619ee2565b935060ff168151811061624757616247619641565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa1580156162b6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526162de9190810190619a0f565b82826162e981619ee2565b935060ff16815181106162fe576162fe619641565b60200260200101819052505b60e085015151156163b15760408051808201909152600a81527f2d2d6761734c696d6974000000000000000000000000000000000000000000006020820152828261635481619ee2565b935060ff168151811061636957616369619641565b60200260200101819052506163858560e0015160000151617216565b828261639081619ee2565b935060ff16815181106163a5576163a5619641565b60200260200101819052505b60e0850151602001511561645b5760408051808201909152600a81527f2d2d676173507269636500000000000000000000000000000000000000000000602082015282826163fe81619ee2565b935060ff168151811061641357616413619641565b602002602001018190525061642f8560e0015160200151617216565b828261643a81619ee2565b935060ff168151811061644f5761644f619641565b60200260200101819052505b60e085015160400151156165055760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826164a881619ee2565b935060ff16815181106164bd576164bd619641565b60200260200101819052506164d98560e0015160400151617216565b82826164e481619ee2565b935060ff16815181106164f9576164f9619641565b60200260200101819052505b60e085015160600151156165af5760408051808201909152601681527f2d2d6d61785072696f72697479466565506572476173000000000000000000006020820152828261655281619ee2565b935060ff168151811061656757616567619641565b60200260200101819052506165838560e0015160600151617216565b828261658e81619ee2565b935060ff16815181106165a3576165a3619641565b60200260200101819052505b60008160ff1667ffffffffffffffff8111156165cd576165cd619229565b60405190808252806020026020018201604052801561660057816020015b60608152602001906001900390816165eb5790505b50905060005b8260ff168160ff16101561665957838160ff168151811061662957616629619641565b6020026020010151828260ff168151811061664657616646619641565b6020908102919091010152600101616606565b5093505050505b949350505050565b61668f6040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c9161671591869101619f6c565b600060405180830381865afa158015616732573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261675a9190810190619a0f565b905060006167688683617da5565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b8152600401616798919061915d565b6000604051808303816000875af11580156167b7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526167df9190810190619fb3565b805190915060030b158015906167f85750602081015151155b80156168075750604081015151155b156157e1578160008151811061681f5761681f619641565b6020026020010151604051602001614e3a919061a069565b6060600061686c8560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925286518252808701908201529091506168a39082905b90617efa565b15616a005760006169208261691a846169146168e68a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b90617f21565b90617f83565b604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150616984908290617efa565b156169ee57604080518082018252600181527f0a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526169eb905b8290618008565b90505b6169f78161802e565b92505050614d67565b8215616a19578484604051602001614e3a92919061a255565b5050604080516020810190915260008152614d67565b509392505050565b6000808251602084016000f09392505050565b8160a0015115616a5957505050565b6000616a66848484618097565b90506000616a7382616668565b602081015181519192509060030b158015616b0f5750604080518082018252600781527f535543434553530000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616b0f9060408051808201825260008082526020918201528151808301909252845182528085019082015261689d565b15616b1c57505050505050565b60408201515115616b3c578160400151604051602001614e3a919061a2fc565b80604051602001614e3a919061a35a565b60606000616b828360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150616be7905b82906171b5565b15616c5657604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152614d6790616c51908390618632565b61802e565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616cb8905b82906186bc565b600103616d8557604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616d1e906169e4565b50604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152614d6790616c51905b8390618008565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616de490616be0565b15616f1b57604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290616e4c908390618756565b905060008160018351616e5f919061a3c5565b81518110616e6f57616e6f619641565b60200260200101519050616f12616c51616ee56040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252855182528086019082015290618632565b95945050505050565b82604051602001614e3a919061a3d8565b50919050565b60606000616f678360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150616fc990616be0565b15616fd757614d678161802e565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261703690616cb1565b6001036170a057604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152614d6790616c5190616d7e565b604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526170ff90616be0565b15616f1b57604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290617167908390618756565b90506001815111156171a3578060028251617182919061a3c5565b8151811061719257617192619641565b602002602001015192505050919050565b5082604051602001614e3a919061a3d8565b8051825160009111156171ca57506000614c0c565b815183516020850151600092916171e09161a4b6565b6171ea919061a3c5565b905082602001518103617201576001915050614c0c565b82516020840151819020912014905092915050565b60606000617223836187fb565b600101905060008167ffffffffffffffff81111561724357617243619229565b6040519080825280601f01601f19166020018201604052801561726d576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461727757509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e5345440000000000000000000000000000000000000000000081840190815285518087018752838152840192909252845180860190955251845290830152606091617342905b82906188dd565b1561738257505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e73650000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526173e19061733b565b1561742157505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d49540000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526174809061733b565b156174c057505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261751f9061733b565b806175845750604080518082018252601081527f47504c2d322e302d6f722d6c6174657200000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526175849061733b565b156175c457505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526176239061733b565b806176885750604080518082018252601081527f47504c2d332e302d6f722d6c6174657200000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526176889061733b565b156176c857505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526177279061733b565b8061778c5750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261778c9061733b565b156177cc57505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261782b9061733b565b806178905750604080518082018252601181527f4c47504c2d332e302d6f722d6c61746572000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526178909061733b565b156178d057505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261792f9061733b565b1561796f57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c617573650000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526179ce9061733b565b15617a0e57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617a6d9061733b565b15617aad57505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617b0c9061733b565b15617b4c57505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e300000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617bab9061733b565b15617beb57505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617c4a9061733b565b80617caf5750604080518082018252601181527f4147504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617caf9061733b565b15617cef57505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e3100000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617d4e9061733b565b15617d8e57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151614e3a929060200161a4c9565b60608060005b8451811015617e305781858281518110617dc757617dc7619641565b6020026020010151604051602001617de09291906198ee565b604051602081830303815290604052915060018551617dff919061a3c5565b8114617e285781604051602001617e16919061a632565b60405160208183030381529060405291505b600101617dab565b5060408051600380825260808201909252600091816020015b6060815260200190600190039081617e495790505090508381600081518110617e7457617e74619641565b60200260200101819052506040518060400160405280600281526020017f2d6300000000000000000000000000000000000000000000000000000000000081525081600181518110617ec857617ec8619641565b60200260200101819052508181600281518110617ee757617ee7619641565b6020908102919091010152949350505050565b6020808301518351835192840151600093617f1892918491906188f1565b14159392505050565b60408051808201909152600080825260208201526000617f538460000151856020015185600001518660200151618a02565b9050836020015181617f65919061a3c5565b84518590617f7490839061a3c5565b90525060208401525090919050565b6040805180820190915260008082526020820152815183511015617fa8575081614c0c565b6020808301519084015160019114617fcf5750815160208481015190840151829020919020145b801561800057825184518590617fe690839061a3c5565b9052508251602085018051617ffc90839061a4b6565b9052505b509192915050565b6040805180820190915260008082526020820152618027838383618b22565b5092915050565b60606000826000015167ffffffffffffffff81111561804f5761804f619229565b6040519080825280601f01601f191660200182016040528015618079576020820181803683370190505b50905060006020820190506180278185602001518660000151618bcd565b606060006180a3615060565b6040805160ff808252612000820190925291925060009190816020015b60608152602001906001900390816180c057905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061811b90619ee2565b935060ff168151811061813057618130619641565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e3300000000000000000000000000000000000000000000000000815250604051602001618181919061a673565b60405160208183030381529060405282828061819c90619ee2565b935060ff16815181106181b1576181b1619641565b60200260200101819052506040518060400160405280600881526020017f76616c69646174650000000000000000000000000000000000000000000000008152508282806181fe90619ee2565b935060ff168151811061821357618213619641565b60200260200101819052508260405160200161822f9190619de1565b60405160208183030381529060405282828061824a90619ee2565b935060ff168151811061825f5761825f619641565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e7472616374000000000000000000000000000000000000000000008152508282806182ac90619ee2565b935060ff16815181106182c1576182c1619641565b60200260200101819052506182d68784618c47565b82826182e181619ee2565b935060ff16815181106182f6576182f6619641565b6020908102919091010152855151156183a25760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261834881619ee2565b935060ff168151811061835d5761835d619641565b6020026020010181905250618376866000015184618c47565b828261838181619ee2565b935060ff168151811061839657618396619641565b60200260200101819052505b8560800151156184105760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b0000000000000000602082015282826183eb81619ee2565b935060ff168151811061840057618400619641565b6020026020010181905250618476565b84156184765760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261845581619ee2565b935060ff168151811061846a5761846a619641565b60200260200101819052505b604086015151156185125760408051808201909152600d81527f2d2d756e73616665416c6c6f7700000000000000000000000000000000000000602082015282826184c081619ee2565b935060ff16815181106184d5576184d5619641565b602002602001018190525085604001518282806184f190619ee2565b935060ff168151811061850657618506619641565b60200260200101819052505b85606001511561857c5760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261855b81619ee2565b935060ff168151811061857057618570619641565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561859a5761859a619229565b6040519080825280602002602001820160405280156185cd57816020015b60608152602001906001900390816185b85790505b50905060005b8260ff168160ff16101561862657838160ff16815181106185f6576185f6619641565b6020026020010151828260ff168151811061861357618613619641565b60209081029190910101526001016185d3565b50979650505050505050565b6040805180820190915260008082526020820152815183511015618657575081614c0c565b8151835160208501516000929161866d9161a4b6565b618677919061a3c5565b60208401519091506001908214618698575082516020840151819020908220145b80156186b3578351855186906186af90839061a3c5565b9052505b50929392505050565b60008082600001516186e08560000151866020015186600001518760200151618a02565b6186ea919061a4b6565b90505b835160208501516186fe919061a4b6565b8111618027578161870e8161a6b8565b9250508260000151618745856020015183618729919061a3c5565b8651618735919061a3c5565b8386600001518760200151618a02565b61874f919061a4b6565b90506186ed565b6060600061876484846186bc565b61876f90600161a4b6565b67ffffffffffffffff81111561878757618787619229565b6040519080825280602002602001820160405280156187ba57816020015b60608152602001906001900390816187a55790505b50905060005b8151811015616a2f576187d6616c518686618008565b8282815181106187e8576187e8619641565b60209081029190910101526001016187c0565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310618844577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310618870576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061888e57662386f26fc10000830492506010015b6305f5e10083106188a6576305f5e100830492506008015b61271083106188ba57612710830492506004015b606483106188cc576064830492506002015b600a8310614c0c5760010192915050565b60006188e98383618c87565b159392505050565b6000808584116189f857602084116189a4576000841561893c57600161891886602061a3c5565b61892390600861a6d2565b61892e90600261a7d0565b618938919061a3c5565b1990505b835181168561894b898961a4b6565b618955919061a3c5565b805190935082165b81811461898f578784116189775787945050505050616660565b836189818161a7dc565b94505082845116905061895d565b618999878561a4b6565b945050505050616660565b8383206189b1858861a3c5565b6189bb908761a4b6565b91505b8582106189f6578482208082036189e3576189d9868461a4b6565b9350505050616660565b6189ee60018461a3c5565b9250506189be565b505b5092949350505050565b60008381868511618b0d5760208511618abc5760008515618a4e576001618a2a87602061a3c5565b618a3590600861a6d2565b618a4090600261a7d0565b618a4a919061a3c5565b1990505b84518116600087618a5f8b8b61a4b6565b618a69919061a3c5565b855190915083165b828114618aae57818610618a9657618a898b8b61a4b6565b9650505050505050616660565b85618aa08161a6b8565b965050838651169050618a71565b859650505050505050616660565b508383206000905b618ace868961a3c5565b8211618b0b57858320808203618aea5783945050505050616660565b618af560018561a4b6565b9350508180618b039061a6b8565b925050618ac4565b505b618b17878761a4b6565b979650505050505050565b60408051808201909152600080825260208201526000618b548560000151866020015186600001518760200151618a02565b602080870180519186019190915251909150618b70908261a3c5565b835284516020860151618b83919061a4b6565b8103618b925760008552618bc4565b83518351618ba0919061a4b6565b85518690618baf90839061a3c5565b9052508351618bbe908261a4b6565b60208601525b50909392505050565b60208110618c055781518352618be460208461a4b6565b9250618bf160208361a4b6565b9150618bfe60208261a3c5565b9050618bcd565b6000198115618c34576001618c1b83602061a3c5565b618c279061010061a7d0565b618c31919061a3c5565b90505b9151835183169219169190911790915250565b60606000618c558484615133565b8051602080830151604051939450618c6f9390910161a7f3565b60405160208183030381529060405291505092915050565b8151815160009190811115618c9a575081515b6020808501519084015160005b83811015618d535782518251808214618d23576000196020871015618d0257600184618cd489602061a3c5565b618cde919061a4b6565b618ce990600861a6d2565b618cf490600261a7d0565b618cfe919061a3c5565b1990505b8181168382168181039114618d20579750614c0c9650505050505050565b50505b618d2e60208661a4b6565b9450618d3b60208561a4b6565b93505050602081618d4c919061a4b6565b9050618ca7565b50845186516157e1919061a84b565b610c9f806200a86c83390190565b6112a6806200b50b83390190565b611db5806200c7b183390190565b611783806200e56683390190565b610de0806200fce983390190565b6040518060e00160405280606081526020016060815260200160608152602001600015158152602001600015158152602001600015158152602001618deb618df0565b905290565b60405180610100016040528060001515815260200160001515815260200160608152602001600080191681526020016060815260200160608152602001600015158152602001618deb6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b81811015618ea25783516001600160a01b0316835260209384019390920191600101618e7b565b509095945050505050565b60005b83811015618ec8578181015183820152602001618eb0565b50506000910152565b60008151808452618ee9816020860160208601618ead565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015618ff9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b81811015618fdf577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8503018352618fc9848651618ed1565b6020958601959094509290920191600101618f8f565b509197505050602094850194929092019150600101618f25565b50929695505050505050565b600081518084526020840193506020830160005b828110156190595781517fffffffff0000000000000000000000000000000000000000000000000000000016865260209586019590910190600101619019565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015618ff9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051604087526190cf6040880182618ed1565b90506020820151915086810360208801526190ea8183619005565b96505050602093840193919091019060010161908b565b600082825180855260208501945060208160051b8301016020850160005b8381101561915157601f1985840301885261913b838351618ed1565b602098890198909350919091019060010161911f565b50909695505050505050565b602081526000614d676020830184619101565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015618ff9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b03815116865260208101519050604060208701526191f16040870182619005565b9550506020938401939190910190600101619198565b6001600160a01b03831681526040602082015260006166606040830184618ed1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561927b5761927b619229565b60405290565b60008067ffffffffffffffff84111561929c5761929c619229565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff821117156192cb576192cb619229565b6040528381529050808284018510156192e357600080fd5b616a2f846020830185618ead565b600082601f83011261930257600080fd5b614d6783835160208501619281565b60006020828403121561932357600080fd5b815167ffffffffffffffff81111561933a57600080fd5b614c08848285016192f1565b602081526000614d676020830184618ed1565b60006020828403121561936b57600080fd5b81518015158114614d6757600080fd5b600181811c9082168061938f57607f821691505b602082108103616f2c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f82111561131657806000526020600020601f840160051c810160208510156193ef5750805b601f840160051c820191505b8181101561269e57600081556001016193fb565b815167ffffffffffffffff81111561942957619429619229565b61943d81619437845461937b565b846193c8565b6020601f82116001811461947157600083156194595750848201515b600019600385901b1c1916600184901b17845561269e565b600084815260208120601f198516915b828110156194a15787850151825560209485019460019092019101619481565b50848210156194bf5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6001600160a01b0384168152826020820152606060408201526000616f126060830184618ed1565b8281526040602082015260006166606040830184618ed1565b600081546001600160a01b038116845267ffffffffffffffff8160a01c1660208501525060018201606060408501526000815461954b8161937b565b806060880152600182166000811461956a57600181146195a4576195d8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083166080890152608082151560051b89010193506195d8565b84600052602060002060005b838110156195cf5781548a8201608001526001909101906020016195b0565b89016080019450505b50919695505050505050565b6001600160a01b03841681526060602082015260006196066060830185618ed1565b82810360408401526157e1818561950f565b60006020828403121561962a57600080fd5b81516001600160a01b0381168114614d6757600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081518084526020840193506020830160005b82811015619059578151865260209586019590910190600101619684565b6060815260006196b56060830186619101565b82810360208401526196c78186619670565b9150508215156040830152949350505050565b6001600160a01b03861681526001600160a01b038516602082015283604082015260a06060820152600061971160a0830185618ed1565b8281036080840152619723818561950f565b98975050505050505050565b60006020828403121561974157600080fd5b5051919050565b6001600160a01b03851681526001600160a01b03841660208201528260408201526080606082015260006157e16080830184618ed1565b6001600160a01b03851681526080602082015260006197a16080830186619101565b82810360408401526197b38186619670565b915050821515606083015295945050505050565b6001600160a01b0383168152604060208201526000616660604083018461950f565b8381526060602082015260006196066060830185618ed1565b6060815260006198156060830186618ed1565b602083019490945250901515604090910152919050565b6001600160a01b038616815284602082015260a06040820152600061985460a0830186618ed1565b6060830194909452509015156080909101529392505050565b7f4661696c656420746f206465706c6f7920636f6e7472616374200000000000008152600083516198a581601a850160208801618ead565b7f3a20000000000000000000000000000000000000000000000000000000000000601a9184019182015283516198e281601c840160208801618ead565b01601c01949350505050565b60008351619900818460208801618ead565b835190830190619914818360208801618ead565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161995581601a850160208801618ead565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a918401918201528351619992816033840160208801618ead565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000614d676080830184618ed1565b600060208284031215619a2157600080fd5b815167ffffffffffffffff811115619a3857600080fd5b8201601f81018413619a4957600080fd5b614c0884825160208401619281565b60008551619a6a818460208a01618ead565b7f2f000000000000000000000000000000000000000000000000000000000000009083019081528551619aa4816001840160208a01618ead565b7f2f00000000000000000000000000000000000000000000000000000000000000600192909101918201528451619ae2816002840160208901618ead565b6001818301019150507f2f0000000000000000000000000000000000000000000000000000000000000060018201528351619b24816002840160208801618ead565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b604081526000619b6f6040830184618ed1565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e2061727469666163742000815260008251619be681601f850160208701618ead565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b604081526000619c536040830184618ed1565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081526000619ca56040830184618ed1565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b27000000000000000000000000815260008251619d1c816014850160208701618ead565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b604081526000619d636040830185618ed1565b8281036020840152614d638185618ed1565b7f2200000000000000000000000000000000000000000000000000000000000000815260008251619dad816001850160208701618ead565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b60008251619df3818460208701618ead565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e747261637420000000000000000000000000000000000000000000604082015260008251619ea681604b850160208701618ead565b91909101604b0192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600060ff821660ff8103619ef857619ef8619eb3565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c69400000000000000000000000000000000000000000000000602082015260008251619f5f816029850160208701618ead565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000614d676080830184618ed1565b600060208284031215619fc557600080fd5b815167ffffffffffffffff811115619fdc57600080fd5b820160608185031215619fee57600080fd5b619ff6619258565b81518060030b811461a00757600080fd5b8152602082015167ffffffffffffffff81111561a02357600080fd5b61a02f868285016192f1565b602083015250604082015167ffffffffffffffff81111561a04f57600080fd5b61a05b868285016192f1565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161a0c7816021850160208701618ead565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161a2b3816021850160208801618ead565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161a2f081602e840160208801618ead565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a200000000000000000000000000000000000000000000000602082015260008251619f5f816029850160208701618ead565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161a3b8816022850160208701618ead565b9190910160220192915050565b81810381811115614c0c57614c0c619eb3565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161a41081600e850160208701618ead565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b80820180821115614c0c57614c0c619eb3565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161a501816018850160208801618ead565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161a53e81601c840160208801618ead565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161a644818460208701618ead565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161a6ab81601c850160208701618ead565b91909101601c0192915050565b6000600019820361a6cb5761a6cb619eb3565b5060010190565b8082028115828204841417614c0c57614c0c619eb3565b6001815b600184111561a7245780850481111561a7085761a708619eb3565b600184161561a71657908102905b60019390931c92800261a6ed565b935093915050565b60008261a73b57506001614c0c565b8161a74857506000614c0c565b816001811461a75e576002811461a7685761a784565b6001915050614c0c565b60ff84111561a7795761a779619eb3565b50506001821b614c0c565b5060208310610133831016604e8410600b841016171561a7a7575081810a614c0c565b61a7b4600019848461a6e9565b806000190482111561a7c85761a7c8619eb3565b029392505050565b6000614d67838361a72c565b60008161a7eb5761a7eb619eb3565b506000190190565b6000835161a805818460208801618ead565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161a83f816001840160208801618ead565b01600101949350505050565b818103600083128015838313168383128216171561802757618027619eb356fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220a043c41353215fce25ecb67a8a4f6f724aaa86dea8dcb0a6975bfb1f10ff3beb64736f6c634300081a0033608060405234801561001057600080fd5b506040516112a63803806112a683398101604081905261002f91610110565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007891906101e2565b50600461008582826101e2565b5050506001600160a01b03821615806100a557506001600160a01b038116155b156100c35760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b031991821617909155600780549290931691161790556102a0565b80516001600160a01b038116811461010b57600080fd5b919050565b6000806040838503121561012357600080fd5b61012c836100f4565b915061013a602084016100f4565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061016d57607f821691505b60208210810361018d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101dd57806000526020600020601f840160051c810160208510156101ba5750805b601f840160051c820191505b818110156101da57600081556001016101c6565b50505b505050565b81516001600160401b038111156101fb576101fb610143565b61020f816102098454610159565b84610193565b6020601f821160018114610243576000831561022b5750848201515b600019600385901b1c1916600184901b1784556101da565b600084815260208120601f198516915b828110156102735787850151825560209485019460019092019101610253565b50848210156102915786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610ff7806102af6000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806342966c68116100b257806379cc679011610081578063a9059cbb11610066578063a9059cbb1461028e578063bff9662a146102a1578063dd62ed3e146102c157600080fd5b806379cc67901461027357806395d89b411461028657600080fd5b806342966c68146102025780635b1125911461021557806370a0823114610235578063779e3b631461026b57600080fd5b80631e458bee116100ee5780631e458bee1461018857806323b872dd1461019b578063313ce567146101ae578063328a01d0146101bd57600080fd5b806306fdde0314610120578063095ea7b31461013e57806315d57fd41461016157806318160ddd14610176575b600080fd5b610128610307565b6040516101359190610d97565b60405180910390f35b61015161014c366004610e2c565b610399565b6040519015158152602001610135565b61017461016f366004610e56565b6103b3565b005b6002545b604051908152602001610135565b610174610196366004610e89565b61057e565b6101516101a9366004610ebc565b610631565b60405160128152602001610135565b6007546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610135565b610174610210366004610ef9565b610655565b6006546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a610243366004610f12565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b610174610662565b610174610281366004610e2c565b610786565b610128610837565b61015161029c366004610e2c565b610846565b6005546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a6102cf366004610e56565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60606003805461031690610f34565b80601f016020809104026020016040519081016040528092919081815260200182805461034290610f34565b801561038f5780601f106103645761010080835404028352916020019161038f565b820191906000526020600020905b81548152906001019060200180831161037257829003601f168201915b5050505050905090565b6000336103a7818585610854565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103f3575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610431576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82161580610468575073ffffffffffffffffffffffffffffffffffffffff8116155b1561049f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105d1576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6105db8383610866565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161062491815260200190565b60405180910390a3505050565b60003361063f8582856108c6565b61064a858585610995565b506001949350505050565b61065f3382610a40565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106b5576040517fe700765e000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b60065473ffffffffffffffffffffffffffffffffffffffff16610704576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107d9576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6107e38282610a9c565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161082b91815260200190565b60405180910390a25050565b60606004805461031690610f34565b6000336103a7818585610995565b6108618383836001610ab1565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108b6576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c260008383610bf9565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461098f5781811015610980576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610428565b61098f84848484036000610ab1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109e5576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8216610a35576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b610861838383610bf9565b73ffffffffffffffffffffffffffffffffffffffff8216610a90576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c282600083610bf9565b610aa78233836108c6565b6108c28282610a40565b73ffffffffffffffffffffffffffffffffffffffff8416610b01576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8316610b51576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600160209081526040808320938716835292905220829055801561098f578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610beb91815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c31578060026000828254610c269190610f87565b90915550610ce39050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610cb7576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610428565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610d0c57600280548290039055610d38565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161062491815260200190565b602081526000825180602084015260005b81811015610dc55760208186018101516040868401015201610da8565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610e2757600080fd5b919050565b60008060408385031215610e3f57600080fd5b610e4883610e03565b946020939093013593505050565b60008060408385031215610e6957600080fd5b610e7283610e03565b9150610e8060208401610e03565b90509250929050565b600080600060608486031215610e9e57600080fd5b610ea784610e03565b95602085013595506040909401359392505050565b600080600060608486031215610ed157600080fd5b610eda84610e03565b9250610ee860208501610e03565b929592945050506040919091013590565b600060208284031215610f0b57600080fd5b5035919050565b600060208284031215610f2457600080fd5b610f2d82610e03565b9392505050565b600181811c90821680610f4857607f821691505b602082108103610f81577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b808201808211156103ad577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea264697066735822122001ec0ce060384773f3d3389fab7bed652c6b8ee389a7471cce10d00d87a75a0c64736f6c634300081a003360a060405234801561001057600080fd5b50604051611db5380380611db583398101604081905261002f91610228565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03838116608052600480546001600160a01b0319169184169190911790556100bb600082610177565b506100e67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610177565b506101117f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610177565b5061012a600080516020611d9583398151915282610177565b50610143600080516020611d9583398151915283610177565b5061016e7f1ddc5e4a2f140581e5d9d5f758b2c0fbdd814c4017a02afc4b55cabcf2c1f10b82610177565b5050505061026b565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166102025760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a4506001610206565b5060005b92915050565b80516001600160a01b038116811461022357600080fd5b919050565b60008060006060848603121561023d57600080fd5b6102468461020c565b92506102546020850161020c565b91506102626040850161020c565b90509250925092565b608051611af36102a260003960008181610215015281816105a6015281816105fb01528181610a8c0152610ae10152611af36000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c806385f438c1116100ee578063c709ab6e11610097578063d9caed1211610071578063d9caed1214610412578063e609055e14610425578063e63ab1e914610438578063eab103df1461045f57600080fd5b8063c709ab6e146103c9578063d547741f146103dc578063d936547e146103ef57600080fd5b80639a590427116100c85780639a5904271461039b5780639b19251a146103ae578063a217fddf146103c157600080fd5b806385f438c11461032857806391d148541461034f578063950837aa1461038857600080fd5b80632f2ff15d1161015b578063570618e111610135578063570618e1146102db5780635b112591146103025780635c975abb146103155780638456cb591461032057600080fd5b80632f2ff15d146102ad57806336568abe146102c05780633f4ba83a146102d357600080fd5b806321fc65f21161018c57806321fc65f21461024f578063248a9ca314610264578063252f07bf1461028857600080fd5b806301ffc9a7146101b35780630d391704146101db578063116191b614610210575b600080fd5b6101c66101c13660046114e9565b610472565b60405190151581526020015b60405180910390f35b6102027f1ddc5e4a2f140581e5d9d5f758b2c0fbdd814c4017a02afc4b55cabcf2c1f10b81565b6040519081526020016101d2565b6102377f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101d2565b61026261025d366004611589565b61050b565b005b6102026102723660046115fc565b6000908152600160208190526040909120015490565b6004546101c69074010000000000000000000000000000000000000000900460ff1681565b6102626102bb366004611615565b6106cb565b6102626102ce366004611615565b6106f7565b610262610748565b6102027f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b600454610237906001600160a01b031681565b60025460ff166101c6565b61026261077d565b6102027f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101c661035d366004611615565b60009182526001602090815260408084206001600160a01b0393909316845291905290205460ff1690565b610262610396366004611645565b6107af565b6102626103a9366004611645565b610886565b6102626103bc366004611645565b61093a565b610202600081565b6102626103d7366004611662565b6109f1565b6102626103ea366004611615565b610bb6565b6101c66103fd366004611645565b60036020526000908152604090205460ff1681565b610262610420366004611705565b610bdc565b610262610433366004611746565b610cd4565b6102027f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b61026261046d3660046117e5565b610f00565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061050557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b610513610f56565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461053d81610f99565b610545610fa3565b6001600160a01b03851660009081526003602052604090205460ff16610597576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105cb6001600160a01b0386167f000000000000000000000000000000000000000000000000000000000000000086610fe2565b6040517f5131ab590000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906106389088908a9089908990899060040161184b565b600060405180830381600087803b15801561065257600080fd5b505af1158015610666573d6000803e3d6000fd5b50505050846001600160a01b0316866001600160a01b03167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d58686866040516106b19392919061188e565b60405180910390a3506106c46001600055565b5050505050565b600082815260016020819052604090912001546106e781610f99565b6106f18383611056565b50505050565b6001600160a01b0381163314610739576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61074382826110e9565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61077281610f99565b61077a611170565b50565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107a781610f99565b61077a6111c2565b7f1ddc5e4a2f140581e5d9d5f758b2c0fbdd814c4017a02afc4b55cabcf2c1f10b6107d981610f99565b6001600160a01b038216610819576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040519081527f086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba9060200160405180910390a15050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6108b081610f99565b6001600160a01b0382166108f0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a61096481610f99565b6001600160a01b0382166109a4576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b6109f9610f56565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610a2381610f99565b610a2b610fa3565b6001600160a01b03861660009081526003602052604090205460ff16610a7d576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ab16001600160a01b0387167f000000000000000000000000000000000000000000000000000000000000000087610fe2565b6040517fd0b492c30000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610b209089908b908a908a908a908a90600401611962565b600060405180830381600087803b158015610b3a57600080fd5b505af1158015610b4e573d6000803e3d6000fd5b50505050856001600160a01b0316876001600160a01b03167f2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b87878787604051610b9b94939291906119b9565b60405180910390a350610bae6001600055565b505050505050565b60008281526001602081905260409091200154610bd281610f99565b6106f183836110e9565b610be4610f56565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610c0e81610f99565b610c16610fa3565b6001600160a01b03831660009081526003602052604090205460ff16610c68576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c7c6001600160a01b0384168584610fe2565b826001600160a01b0316846001600160a01b03167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610cc191815260200190565b60405180910390a3506107436001600055565b610cdc610f56565b610ce4610fa3565b60045474010000000000000000000000000000000000000000900460ff16610d38576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03841660009081526003602052604090205460ff16610d8a576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610dea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e0e91906119e5565b9050610e256001600160a01b0386163330876111ff565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa158015610eac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed091906119e5565b610eda91906119fe565b8787604051610eed959493929190611a38565b60405180910390a250610bae6001600055565b6000610f0b81610f99565b506004805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b600260005403610f92576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61077a8133611238565b60025460ff1615610fe0576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6040516001600160a01b0383811660248301526044820183905261074391859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506112af565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166110e15760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a4506001610505565b506000610505565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff16156110e15760008381526001602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610505565b61117861132b565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6111ca610fa3565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586111a53390565b6040516001600160a01b0384811660248301528381166044830152606482018390526106f19186918216906323b872dd9060840161100f565b60008281526001602090815260408083206001600160a01b038516845290915290205460ff166112ab576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044015b60405180910390fd5b5050565b60006112c46001600160a01b03841683611367565b905080516000141580156112e95750808060200190518101906112e79190611a71565b155b15610743576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201526024016112a2565b60025460ff16610fe0576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606113758383600061137c565b9392505050565b6060814710156113ba576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016112a2565b600080856001600160a01b031684866040516113d69190611a8e565b60006040518083038185875af1925050503d8060008114611413576040519150601f19603f3d011682016040523d82523d6000602084013e611418565b606091505b5091509150611428868383611432565b9695505050505050565b60608261144757611442826114a7565b611375565b815115801561145e57506001600160a01b0384163b155b156114a0576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016112a2565b5080611375565b8051156114b75780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156114fb57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461137557600080fd5b6001600160a01b038116811461077a57600080fd5b60008083601f84011261155257600080fd5b50813567ffffffffffffffff81111561156a57600080fd5b60208301915083602082850101111561158257600080fd5b9250929050565b6000806000806000608086880312156115a157600080fd5b85356115ac8161152b565b945060208601356115bc8161152b565b935060408601359250606086013567ffffffffffffffff8111156115df57600080fd5b6115eb88828901611540565b969995985093965092949392505050565b60006020828403121561160e57600080fd5b5035919050565b6000806040838503121561162857600080fd5b82359150602083013561163a8161152b565b809150509250929050565b60006020828403121561165757600080fd5b81356113758161152b565b60008060008060008060a0878903121561167b57600080fd5b86356116868161152b565b955060208701356116968161152b565b945060408701359350606087013567ffffffffffffffff8111156116b957600080fd5b6116c589828a01611540565b909450925050608087013567ffffffffffffffff8111156116e557600080fd5b87016060818a0312156116f757600080fd5b809150509295509295509295565b60008060006060848603121561171a57600080fd5b83356117258161152b565b925060208401356117358161152b565b929592945050506040919091013590565b6000806000806000806080878903121561175f57600080fd5b863567ffffffffffffffff81111561177657600080fd5b61178289828a01611540565b90975095505060208701356117968161152b565b935060408701359250606087013567ffffffffffffffff8111156117b957600080fd5b6117c589828a01611540565b979a9699509497509295939492505050565b801515811461077a57600080fd5b6000602082840312156117f757600080fd5b8135611375816117d7565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b03861681526001600160a01b0385166020820152836040820152608060608201526000611883608083018486611802565b979650505050505050565b8381526040602082015260006118a8604083018486611802565b95945050505050565b600081356118be8161152b565b6001600160a01b03168352602082013567ffffffffffffffff81168082146118e557600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261192157600080fd5b820160208101903567ffffffffffffffff81111561193e57600080fd5b80360382131561194d57600080fd5b606060408601526118a8606086018284611802565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a06060820152600061199a60a083018587611802565b82810360808401526119ac81856118b1565b9998505050505050505050565b8481526060602082015260006119d3606083018587611802565b828103604084015261188381856118b1565b6000602082840312156119f757600080fd5b5051919050565b81810381811115610505577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b606081526000611a4c606083018789611802565b8560208401528281036040840152611a65818587611802565b98975050505050505050565b600060208284031215611a8357600080fd5b8151611375816117d7565b6000825160005b81811015611aaf5760208186018101518583015201611a95565b50600092019182525091905056fea26469706673582212208a3bf91c6c460ac36647ec155bee4b51ba280bf5bd55656af4321c281774a7bd64736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60c060405260001960035534801561001657600080fd5b5060405161178338038061178383398101604081905261003591610220565b60016000819055805460ff19169055838383836001600160a01b038416158061006557506001600160a01b038316155b8061007757506001600160a01b038216155b8061008957506001600160a01b038116155b156100a75760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03808516608052831660a0526100c5600082610154565b506100f07f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610154565b5061011b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610154565b506101467f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610154565b505050505050505050610274565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166101fa5760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101b23390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101fe565b5060005b92915050565b80516001600160a01b038116811461021b57600080fd5b919050565b6000806000806080858703121561023657600080fd5b61023f85610204565b935061024d60208601610204565b925061025b60408601610204565b915061026960608601610204565b905092959194509250565b60805160a0516114ab6102d86000396000818161021d01528181610531015281816108290152818161099a01528181610aef0152610c110152600081816101d1015281816104a1015281816105040152818161079901526107fc01526114ab6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80635e3e9fef116100d857806391d148541161008c578063d547741f11610066578063d547741f14610386578063d5abeb0114610399578063e63ab1e9146103a257600080fd5b806391d1485414610311578063a217fddf14610357578063a783c7891461035f57600080fd5b8063743e0c9b116100bd578063743e0c9b146102cf5780638456cb59146102e257806385f438c1146102ea57600080fd5b80635e3e9fef146102a95780636f8b44b0146102bc57600080fd5b8063248a9ca31161012f57806336568abe1161011457806336568abe146102835780633f4ba83a146102965780635c975abb1461029e57600080fd5b8063248a9ca31461023f5780632f2ff15d1461027057600080fd5b8063106e629011610160578063106e6290146101b9578063116191b6146101cc57806321e093b11461021857600080fd5b806301ffc9a71461017c578063057e0f25146101a4575b600080fd5b61018f61018a366004610fd3565b6103c9565b60405190151581526020015b60405180910390f35b6101b76101b236600461108e565b610462565b005b6101b76101c7366004611126565b6105fc565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b61026261024d366004611159565b60009081526002602052604090206001015490565b60405190815260200161019b565b6101b761027e366004611172565b6106a1565b6101b7610291366004611172565b6106cc565b6101b7610725565b60015460ff1661018f565b6101b76102b736600461119e565b61075a565b6101b76102ca366004611159565b6108ef565b6101b76102dd366004611159565b61095d565b6101b7610a07565b6102627f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b61018f61031f366004611172565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610262600081565b6102627f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101b7610394366004611172565b610a39565b61026260035481565b6102627f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061045c57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b61046a610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461049481610aa1565b61049c610aab565b6104c77f00000000000000000000000000000000000000000000000000000000000000008785610aea565b6040517fd0b492c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610563907f0000000000000000000000000000000000000000000000000000000000000000908b908b908b908b908a9060040161130e565b600060405180830381600087803b15801561057d57600080fd5b505af1158015610591573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff167f52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c878787866040516105e1949392919061137f565b60405180910390a2506105f46001600055565b505050505050565b610604610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461062e81610aa1565b610636610aab565b610641848484610aea565b8373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d58460405161068991815260200190565b60405180910390a25061069c6001600055565b505050565b6000828152600260205260409020600101546106bc81610aa1565b6106c68383610c72565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331461071b576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61069c8282610d72565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61074f81610aa1565b610757610e31565b50565b610762610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461078c81610aa1565b610794610aab565b6107bf7f00000000000000000000000000000000000000000000000000000000000000008684610aea565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab5990610859907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a906004016113b6565b600060405180830381600087803b15801561087357600080fd5b505af1158015610887573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d8686866040516108d593929190611408565b60405180910390a2506108e86001600055565b5050505050565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61091981610aa1565b610921610aab565b60038290556040518281527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c9060200160405180910390a15050565b610965610aab565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b1580156109f357600080fd5b505af11580156108e8573d6000803e3d6000fd5b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a3181610aa1565b610757610eae565b600082815260026020526040902060010154610a5481610aa1565b6106c68383610d72565b600260005403610a9a576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6107578133610f07565b60015460ff1615610ae8576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6003547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7c9190611422565b610b86908461143b565b1115610bbe576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201849052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b158015610c5557600080fd5b505af1158015610c69573d6000803e3d6000fd5b50505050505050565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610d6a57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610d083390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161045c565b50600061045c565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610d6a57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161045c565b610e39610f97565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b610eb6610aab565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610e84565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610f93576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440160405180910390fd5b5050565b60015460ff16610ae8576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215610fe557600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461101557600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461104057600080fd5b919050565b60008083601f84011261105757600080fd5b50813567ffffffffffffffff81111561106f57600080fd5b60208301915083602082850101111561108757600080fd5b9250929050565b60008060008060008060a087890312156110a757600080fd5b6110b08761101c565b955060208701359450604087013567ffffffffffffffff8111156110d357600080fd5b6110df89828a01611045565b90955093505060608701359150608087013567ffffffffffffffff81111561110657600080fd5b87016060818a03121561111857600080fd5b809150509295509295509295565b60008060006060848603121561113b57600080fd5b6111448461101c565b95602085013595506040909401359392505050565b60006020828403121561116b57600080fd5b5035919050565b6000806040838503121561118557600080fd5b823591506111956020840161101c565b90509250929050565b6000806000806000608086880312156111b657600080fd5b6111bf8661101c565b945060208601359350604086013567ffffffffffffffff8111156111e257600080fd5b6111ee88828901611045565b96999598509660600135949350505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff6112678261101c565b1682526000602082013567ffffffffffffffff811680821461128857600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126112c457600080fd5b820160208101903567ffffffffffffffff8111156112e157600080fd5b8036038213156112f057600080fd5b60606040860152611305606086018284611200565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a06060820152600061136060a083018587611200565b82810360808401526113728185611249565b9998505050505050505050565b848152606060208201526000611399606083018587611200565b82810360408401526113ab8185611249565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006113ab608083018486611200565b838152604060208201526000611305604083018486611200565b60006020828403121561143457600080fd5b5051919050565b8082018082111561045c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea26469706673582212201a525a06b9312846603a9a514a3c7da5c267e70352788b0c22c2ab173492c28964736f6c634300081a00336080604052348015600f57600080fd5b506001600055610dbc806100246000396000f3fe6080604052600436106100635760003560e01c8063c513169111610040578063c5131691146100c1578063e04d4f97146100e1578063f05b6abf146100f457005b8063357fc5a21461006c578063660b9de01461008c5780636ed70169146100ac57005b3661006a57005b005b34801561007857600080fd5b5061006a6100873660046106bd565b610114565b34801561009857600080fd5b5061006a6100a73660046106f9565b6101aa565b3480156100b857600080fd5b5061006a6101e6565b3480156100cd57600080fd5b5061006a6100dc3660046106bd565b61021b565b61006a6100ef366004610859565b6102f6565b34801561010057600080fd5b5061006a61010f366004610945565b61033a565b61011c61036f565b61013e73ffffffffffffffffffffffffffffffffffffffff83163383866103b2565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101a56001600055565b505050565b7f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e33826040516101db929190610a78565b60405180910390a150565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61022361036f565b6000610230600285610b57565b90508060000361026c576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61028e73ffffffffffffffffffffffffffffffffffffffff84163384846103b2565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101a56001600055565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa333485858560405161032d959493929190610c00565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b1463384848460405161032d9493929190610c8a565b6002600054036103ab576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261044790859061044d565b50505050565b600061046f73ffffffffffffffffffffffffffffffffffffffff8416836104e8565b905080516000141580156104945750808060200190518101906104929190610d4d565b155b156101a5576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b60606104f6838360006104fd565b9392505050565b60608147101561053b576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016104df565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105649190610d6a565b60006040518083038185875af1925050503d80600081146105a1576040519150601f19603f3d011682016040523d82523d6000602084013e6105a6565b606091505b50915091506105b68683836105c0565b9695505050505050565b6060826105d5576105d08261064f565b6104f6565b81511580156105f9575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610648576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016104df565b50806104f6565b80511561065f5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff811681146106b857600080fd5b919050565b6000806000606084860312156106d257600080fd5b833592506106e260208501610694565b91506106f060408501610694565b90509250925092565b60006020828403121561070b57600080fd5b813567ffffffffffffffff81111561072257600080fd5b8201606081850312156104f657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156107aa576107aa610734565b604052919050565b600082601f8301126107c357600080fd5b813567ffffffffffffffff8111156107dd576107dd610734565b61080e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610763565b81815284602083860101111561082357600080fd5b816020850160208301376000918101602001919091529392505050565b801515811461069157600080fd5b80356106b881610840565b60008060006060848603121561086e57600080fd5b833567ffffffffffffffff81111561088557600080fd5b610891868287016107b2565b9350506020840135915060408401356108a981610840565b809150509250925092565b600067ffffffffffffffff8211156108ce576108ce610734565b5060051b60200190565b600082601f8301126108e957600080fd5b81356108fc6108f7826108b4565b610763565b8082825260208201915060208360051b86010192508583111561091e57600080fd5b602085015b8381101561093b578035835260209283019201610923565b5095945050505050565b60008060006060848603121561095a57600080fd5b833567ffffffffffffffff81111561097157600080fd5b8401601f8101861361098257600080fd5b80356109906108f7826108b4565b8082825260208201915060208360051b8501019250888311156109b257600080fd5b602084015b838110156109f457803567ffffffffffffffff8111156109d657600080fd5b6109e58b6020838901016107b2565b845250602092830192016109b7565b509550505050602084013567ffffffffffffffff811115610a1457600080fd5b610a20868287016108d8565b9250506106f06040850161084e565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610ab683610694565b1660408201526000602083013567ffffffffffffffff8116808214610ada57600080fd5b6060840152506040830135368490037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112610b1657600080fd5b830160208101903567ffffffffffffffff811115610b3357600080fd5b803603821315610b4257600080fd5b606060808501526105b660a085018284610a2f565b600082610b8d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60005b83811015610bad578181015183820152602001610b95565b50506000910152565b60008151808452610bce816020860160208601610b92565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610c3560a0830186610bb6565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610c80578151865260209586019590910190600101610c62565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610d1d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610d08858351610bb6565b94506020938401939190910190600101610cce565b505050508281036040840152610d338186610c4e565b915050610d44606083018415159052565b95945050505050565b600060208284031215610d5f57600080fd5b81516104f681610840565b60008251610d7c818460208701610b92565b919091019291505056fea26469706673582212202c10ab94780c92f0dd656897a07025a325b3c0d38d265113cb76037b152226de64736f6c634300081a0033a2646970667358221220428585db7dfc8514ed7915cd2490266f8b527d34391c95e587be330cd12dbae964736f6c634300081a0033",
}

// GatewayEVMTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMTestMetaData.ABI instead.
var GatewayEVMTestABI = GatewayEVMTestMetaData.ABI

// GatewayEVMTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMTestMetaData.Bin instead.
var GatewayEVMTestBin = GatewayEVMTestMetaData.Bin

// DeployGatewayEVMTest deploys a new Ethereum contract, binding an instance of GatewayEVMTest to it.
func DeployGatewayEVMTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMTest, error) {
	parsed, err := GatewayEVMTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMTest{GatewayEVMTestCaller: GatewayEVMTestCaller{contract: contract}, GatewayEVMTestTransactor: GatewayEVMTestTransactor{contract: contract}, GatewayEVMTestFilterer: GatewayEVMTestFilterer{contract: contract}}, nil
}

// GatewayEVMTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMTest struct {
	GatewayEVMTestCaller     // Read-only binding to the contract
	GatewayEVMTestTransactor // Write-only binding to the contract
	GatewayEVMTestFilterer   // Log filterer for contract events
}

// GatewayEVMTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMTestSession struct {
	Contract     *GatewayEVMTest   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayEVMTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMTestCallerSession struct {
	Contract *GatewayEVMTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GatewayEVMTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMTestTransactorSession struct {
	Contract     *GatewayEVMTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GatewayEVMTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMTestRaw struct {
	Contract *GatewayEVMTest // Generic contract binding to access the raw methods on
}

// GatewayEVMTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMTestCallerRaw struct {
	Contract *GatewayEVMTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMTestTransactorRaw struct {
	Contract *GatewayEVMTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMTest creates a new instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMTest, error) {
	contract, err := bindGatewayEVMTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTest{GatewayEVMTestCaller: GatewayEVMTestCaller{contract: contract}, GatewayEVMTestTransactor: GatewayEVMTestTransactor{contract: contract}, GatewayEVMTestFilterer: GatewayEVMTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMTestCaller creates a new read-only instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMTestCaller, error) {
	contract, err := bindGatewayEVMTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestCaller{contract: contract}, nil
}

// NewGatewayEVMTestTransactor creates a new write-only instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMTestTransactor, error) {
	contract, err := bindGatewayEVMTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestTransactor{contract: contract}, nil
}

// NewGatewayEVMTestFilterer creates a new log filterer instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMTestFilterer, error) {
	contract, err := bindGatewayEVMTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestFilterer{contract: contract}, nil
}

// bindGatewayEVMTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMTest *GatewayEVMTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMTest.Contract.GatewayEVMTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMTest *GatewayEVMTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.GatewayEVMTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMTest *GatewayEVMTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.GatewayEVMTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMTest *GatewayEVMTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMTest *GatewayEVMTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMTest *GatewayEVMTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.contract.Transact(opts, method, params...)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) ASSETHANDLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "ASSET_HANDLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.ASSETHANDLERROLE(&_GatewayEVMTest.CallOpts)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.ASSETHANDLERROLE(&_GatewayEVMTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.DEFAULTADMINROLE(&_GatewayEVMTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.DEFAULTADMINROLE(&_GatewayEVMTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestSession) ISTEST() (bool, error) {
	return _GatewayEVMTest.Contract.ISTEST(&_GatewayEVMTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ISTEST() (bool, error) {
	return _GatewayEVMTest.Contract.ISTEST(&_GatewayEVMTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.PAUSERROLE(&_GatewayEVMTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.PAUSERROLE(&_GatewayEVMTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.TSSROLE(&_GatewayEVMTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.TSSROLE(&_GatewayEVMTest.CallOpts)
}

// TSSUPDATERROLE is a free data retrieval call binding the contract method 0x0d391704.
//
// Solidity: function TSS_UPDATER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) TSSUPDATERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "TSS_UPDATER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSUPDATERROLE is a free data retrieval call binding the contract method 0x0d391704.
//
// Solidity: function TSS_UPDATER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) TSSUPDATERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.TSSUPDATERROLE(&_GatewayEVMTest.CallOpts)
}

// TSSUPDATERROLE is a free data retrieval call binding the contract method 0x0d391704.
//
// Solidity: function TSS_UPDATER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TSSUPDATERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.TSSUPDATERROLE(&_GatewayEVMTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) WITHDRAWERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.WITHDRAWERROLE(&_GatewayEVMTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.WITHDRAWERROLE(&_GatewayEVMTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.ExcludeArtifacts(&_GatewayEVMTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.ExcludeArtifacts(&_GatewayEVMTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeContracts(&_GatewayEVMTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeContracts(&_GatewayEVMTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.ExcludeSelectors(&_GatewayEVMTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.ExcludeSelectors(&_GatewayEVMTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeSenders(&_GatewayEVMTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeSenders(&_GatewayEVMTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestSession) Failed() (bool, error) {
	return _GatewayEVMTest.Contract.Failed(&_GatewayEVMTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) Failed() (bool, error) {
	return _GatewayEVMTest.Contract.Failed(&_GatewayEVMTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMTest.Contract.TargetArtifactSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMTest.Contract.TargetArtifactSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.TargetArtifacts(&_GatewayEVMTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.TargetArtifacts(&_GatewayEVMTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetContracts(&_GatewayEVMTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetContracts(&_GatewayEVMTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMTest.Contract.TargetInterfaces(&_GatewayEVMTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMTest.Contract.TargetInterfaces(&_GatewayEVMTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.TargetSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.TargetSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetSenders(&_GatewayEVMTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetSenders(&_GatewayEVMTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.SetUp(&_GatewayEVMTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.SetUp(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x04b974f9.
//
// Solidity: function testExecuteFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteFailsIfDestinationIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteFailsIfDestinationIsZeroAddress")
}

// TestExecuteFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x04b974f9.
//
// Solidity: function testExecuteFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x04b974f9.
//
// Solidity: function testExecuteFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevert is a paid mutator transaction binding the contract method 0xfa18c09b.
//
// Solidity: function testExecuteRevert() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteRevert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteRevert")
}

// TestExecuteRevert is a paid mutator transaction binding the contract method 0xfa18c09b.
//
// Solidity: function testExecuteRevert() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteRevert() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevert(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevert is a paid mutator transaction binding the contract method 0xfa18c09b.
//
// Solidity: function testExecuteRevert() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteRevert() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevert(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51010e49.
//
// Solidity: function testExecuteRevertFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteRevertFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteRevertFailsIfReceiverIsZeroAddress")
}

// TestExecuteRevertFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51010e49.
//
// Solidity: function testExecuteRevertFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteRevertFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfReceiverIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51010e49.
//
// Solidity: function testExecuteRevertFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteRevertFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfReceiverIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0xcebad2a6.
//
// Solidity: function testExecuteRevertFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteRevertFailsIfSenderIsNotTSS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteRevertFailsIfSenderIsNotTSS")
}

// TestExecuteRevertFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0xcebad2a6.
//
// Solidity: function testExecuteRevertFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteRevertFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0xcebad2a6.
//
// Solidity: function testExecuteRevertFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteRevertFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xe6afc790.
//
// Solidity: function testExecuteWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteWithERC20FailsIfNotCustodyOrConnector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteWithERC20FailsIfNotCustodyOrConnector")
}

// TestExecuteWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xe6afc790.
//
// Solidity: function testExecuteWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xe6afc790.
//
// Solidity: function testExecuteWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParams is a paid mutator transaction binding the contract method 0x44671b94.
//
// Solidity: function testForwardCallToReceiveNoParams() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNoParams(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNoParams")
}

// TestForwardCallToReceiveNoParams is a paid mutator transaction binding the contract method 0x44671b94.
//
// Solidity: function testForwardCallToReceiveNoParams() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNoParams() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParams(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParams is a paid mutator transaction binding the contract method 0x44671b94.
//
// Solidity: function testForwardCallToReceiveNoParams() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNoParams() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParams(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParamsTogglePause is a paid mutator transaction binding the contract method 0xdd51e82f.
//
// Solidity: function testForwardCallToReceiveNoParamsTogglePause() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNoParamsTogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNoParamsTogglePause")
}

// TestForwardCallToReceiveNoParamsTogglePause is a paid mutator transaction binding the contract method 0xdd51e82f.
//
// Solidity: function testForwardCallToReceiveNoParamsTogglePause() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNoParamsTogglePause() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParamsTogglePause(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParamsTogglePause is a paid mutator transaction binding the contract method 0xdd51e82f.
//
// Solidity: function testForwardCallToReceiveNoParamsTogglePause() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNoParamsTogglePause() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParamsTogglePause(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayable is a paid mutator transaction binding the contract method 0xf68bd1c0.
//
// Solidity: function testForwardCallToReceiveNonPayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNonPayable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNonPayable")
}

// TestForwardCallToReceiveNonPayable is a paid mutator transaction binding the contract method 0xf68bd1c0.
//
// Solidity: function testForwardCallToReceiveNonPayable() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNonPayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayable(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayable is a paid mutator transaction binding the contract method 0xf68bd1c0.
//
// Solidity: function testForwardCallToReceiveNonPayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNonPayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayable(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x7d7f772a.
//
// Solidity: function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS")
}

// TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x7d7f772a.
//
// Solidity: function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x7d7f772a.
//
// Solidity: function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0xfe7bdbb2.
//
// Solidity: function testForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceivePayable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceivePayable")
}

// TestForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0xfe7bdbb2.
//
// Solidity: function testForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceivePayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceivePayable(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0xfe7bdbb2.
//
// Solidity: function testForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceivePayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceivePayable(&_GatewayEVMTest.TransactOpts)
}

// TestRevertWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xb124dbf5.
//
// Solidity: function testRevertWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestRevertWithERC20FailsIfNotCustodyOrConnector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testRevertWithERC20FailsIfNotCustodyOrConnector")
}

// TestRevertWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xb124dbf5.
//
// Solidity: function testRevertWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestRevertWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestRevertWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestRevertWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xb124dbf5.
//
// Solidity: function testRevertWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestRevertWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestRevertWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xccf20616.
//
// Solidity: function testSetConnectorFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetConnectorFailsIfSenderIsNotAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetConnectorFailsIfSenderIsNotAdmin")
}

// TestSetConnectorFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xccf20616.
//
// Solidity: function testSetConnectorFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetConnectorFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xccf20616.
//
// Solidity: function testSetConnectorFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetConnectorFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSet is a paid mutator transaction binding the contract method 0x7ebf744f.
//
// Solidity: function testSetConnectorFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetConnectorFailsIfSet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetConnectorFailsIfSet")
}

// TestSetConnectorFailsIfSet is a paid mutator transaction binding the contract method 0x7ebf744f.
//
// Solidity: function testSetConnectorFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetConnectorFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSet is a paid mutator transaction binding the contract method 0x7ebf744f.
//
// Solidity: function testSetConnectorFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetConnectorFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfZero is a paid mutator transaction binding the contract method 0xce5871e1.
//
// Solidity: function testSetConnectorFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetConnectorFailsIfZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetConnectorFailsIfZero")
}

// TestSetConnectorFailsIfZero is a paid mutator transaction binding the contract method 0xce5871e1.
//
// Solidity: function testSetConnectorFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetConnectorFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfZero is a paid mutator transaction binding the contract method 0xce5871e1.
//
// Solidity: function testSetConnectorFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetConnectorFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xa56f7a4b.
//
// Solidity: function testSetCustodyFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetCustodyFailsIfSenderIsNotAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetCustodyFailsIfSenderIsNotAdmin")
}

// TestSetCustodyFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xa56f7a4b.
//
// Solidity: function testSetCustodyFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetCustodyFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xa56f7a4b.
//
// Solidity: function testSetCustodyFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetCustodyFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSet is a paid mutator transaction binding the contract method 0x1855c337.
//
// Solidity: function testSetCustodyFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetCustodyFailsIfSet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetCustodyFailsIfSet")
}

// TestSetCustodyFailsIfSet is a paid mutator transaction binding the contract method 0x1855c337.
//
// Solidity: function testSetCustodyFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetCustodyFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSet is a paid mutator transaction binding the contract method 0x1855c337.
//
// Solidity: function testSetCustodyFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetCustodyFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfZero is a paid mutator transaction binding the contract method 0x1226c655.
//
// Solidity: function testSetCustodyFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetCustodyFailsIfZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetCustodyFailsIfZero")
}

// TestSetCustodyFailsIfZero is a paid mutator transaction binding the contract method 0x1226c655.
//
// Solidity: function testSetCustodyFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetCustodyFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfZero is a paid mutator transaction binding the contract method 0x1226c655.
//
// Solidity: function testSetCustodyFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetCustodyFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestTSSUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testTSSUpgrade")
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgrade(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgrade(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testTSSUpgradeFailsIfSenderIsNotTSSUpdater")
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestTSSUpgradeFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testTSSUpgradeFailsIfZeroAddress")
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// GatewayEVMTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayEVMTest contract.
type GatewayEVMTestCalledIterator struct {
	Event *GatewayEVMTestCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestCalled represents a Called event raised by the GatewayEVMTest contract.
type GatewayEVMTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestCalledIterator{contract: _GatewayEVMTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestCalled)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Called", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseCalled(log types.Log) (*GatewayEVMTestCalled, error) {
	event := new(GatewayEVMTestCalled)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayEVMTest contract.
type GatewayEVMTestDepositedIterator struct {
	Event *GatewayEVMTestDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestDeposited represents a Deposited event raised by the GatewayEVMTest contract.
type GatewayEVMTestDeposited struct {
	Recipient []byte
	Asset     common.Address
	Amount    *big.Int
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterDeposited(opts *bind.FilterOpts, asset []common.Address) (*GatewayEVMTestDepositedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestDepositedIterator{contract: _GatewayEVMTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestDeposited, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestDeposited)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseDeposited(log types.Log) (*GatewayEVMTestDeposited, error) {
	event := new(GatewayEVMTestDeposited)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestDeposited0Iterator is returned from FilterDeposited0 and is used to iterate over the raw logs and unpacked data for Deposited0 events raised by the GatewayEVMTest contract.
type GatewayEVMTestDeposited0Iterator struct {
	Event *GatewayEVMTestDeposited0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestDeposited0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestDeposited0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestDeposited0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestDeposited0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestDeposited0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestDeposited0 represents a Deposited0 event raised by the GatewayEVMTest contract.
type GatewayEVMTestDeposited0 struct {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterDeposited0(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMTestDeposited0Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Deposited0", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestDeposited0Iterator{contract: _GatewayEVMTest.contract, event: "Deposited0", logs: logs, sub: sub}, nil
}

// WatchDeposited0 is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchDeposited0(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestDeposited0, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Deposited0", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestDeposited0)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseDeposited0(log types.Log) (*GatewayEVMTestDeposited0, error) {
	event := new(GatewayEVMTestDeposited0)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedIterator struct {
	Event *GatewayEVMTestExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestExecuted represents a Executed event raised by the GatewayEVMTest contract.
type GatewayEVMTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestExecutedIterator{contract: _GatewayEVMTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestExecuted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseExecuted(log types.Log) (*GatewayEVMTestExecuted, error) {
	event := new(GatewayEVMTestExecuted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMTestExecutedWithERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestExecutedWithERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestExecutedWithERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestExecutedWithERC20Iterator{contract: _GatewayEVMTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestExecutedWithERC20)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMTestExecutedWithERC20, error) {
	event := new(GatewayEVMTestExecutedWithERC20)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedERC20Iterator struct {
	Event *GatewayEVMTestReceivedERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReceivedERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedERC20 represents a ReceivedERC20 event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*GatewayEVMTestReceivedERC20Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedERC20Iterator{contract: _GatewayEVMTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedERC20)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedERC20(log types.Log) (*GatewayEVMTestReceivedERC20, error) {
	event := new(GatewayEVMTestReceivedERC20)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNoParamsIterator struct {
	Event *GatewayEVMTestReceivedNoParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedNoParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReceivedNoParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedNoParams represents a ReceivedNoParams event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*GatewayEVMTestReceivedNoParamsIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedNoParamsIterator{contract: _GatewayEVMTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedNoParams)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedNoParams(log types.Log) (*GatewayEVMTestReceivedNoParams, error) {
	event := new(GatewayEVMTestReceivedNoParams)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNonPayableIterator struct {
	Event *GatewayEVMTestReceivedNonPayable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedNonPayable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReceivedNonPayable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedNonPayable represents a ReceivedNonPayable event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*GatewayEVMTestReceivedNonPayableIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedNonPayableIterator{contract: _GatewayEVMTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedNonPayable)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedNonPayable(log types.Log) (*GatewayEVMTestReceivedNonPayable, error) {
	event := new(GatewayEVMTestReceivedNonPayable)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedPayableIterator struct {
	Event *GatewayEVMTestReceivedPayable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedPayable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReceivedPayable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedPayable represents a ReceivedPayable event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedPayable struct {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*GatewayEVMTestReceivedPayableIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedPayableIterator{contract: _GatewayEVMTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedPayable)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedPayable(log types.Log) (*GatewayEVMTestReceivedPayable, error) {
	event := new(GatewayEVMTestReceivedPayable)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedRevertIterator struct {
	Event *GatewayEVMTestReceivedRevert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedRevert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReceivedRevert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedRevert represents a ReceivedRevert event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*GatewayEVMTestReceivedRevertIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedRevertIterator{contract: _GatewayEVMTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedRevert)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedRevert(log types.Log) (*GatewayEVMTestReceivedRevert, error) {
	event := new(GatewayEVMTestReceivedRevert)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVMTest contract.
type GatewayEVMTestRevertedIterator struct {
	Event *GatewayEVMTestReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReverted represents a Reverted event raised by the GatewayEVMTest contract.
type GatewayEVMTestReverted struct {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestRevertedIterator{contract: _GatewayEVMTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0x1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b1436.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReverted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReverted(log types.Log) (*GatewayEVMTestReverted, error) {
	event := new(GatewayEVMTestReverted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUnwhitelistedIterator is returned from FilterUnwhitelisted and is used to iterate over the raw logs and unpacked data for Unwhitelisted events raised by the GatewayEVMTest contract.
type GatewayEVMTestUnwhitelistedIterator struct {
	Event *GatewayEVMTestUnwhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUnwhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestUnwhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUnwhitelisted represents a Unwhitelisted event raised by the GatewayEVMTest contract.
type GatewayEVMTestUnwhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnwhitelisted is a free log retrieval operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUnwhitelisted(opts *bind.FilterOpts, token []common.Address) (*GatewayEVMTestUnwhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUnwhitelistedIterator{contract: _GatewayEVMTest.contract, event: "Unwhitelisted", logs: logs, sub: sub}, nil
}

// WatchUnwhitelisted is a free log subscription operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUnwhitelisted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUnwhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUnwhitelisted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUnwhitelisted(log types.Log) (*GatewayEVMTestUnwhitelisted, error) {
	event := new(GatewayEVMTestUnwhitelisted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUpdatedCustodyTSSAddressIterator is returned from FilterUpdatedCustodyTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedCustodyTSSAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedCustodyTSSAddressIterator struct {
	Event *GatewayEVMTestUpdatedCustodyTSSAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestUpdatedCustodyTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUpdatedCustodyTSSAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestUpdatedCustodyTSSAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestUpdatedCustodyTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUpdatedCustodyTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUpdatedCustodyTSSAddress represents a UpdatedCustodyTSSAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedCustodyTSSAddress struct {
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCustodyTSSAddress is a free log retrieval operation binding the contract event 0x086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba.
//
// Solidity: event UpdatedCustodyTSSAddress(address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUpdatedCustodyTSSAddress(opts *bind.FilterOpts) (*GatewayEVMTestUpdatedCustodyTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUpdatedCustodyTSSAddressIterator{contract: _GatewayEVMTest.contract, event: "UpdatedCustodyTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedCustodyTSSAddress is a free log subscription operation binding the contract event 0x086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba.
//
// Solidity: event UpdatedCustodyTSSAddress(address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUpdatedCustodyTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUpdatedCustodyTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUpdatedCustodyTSSAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUpdatedCustodyTSSAddress(log types.Log) (*GatewayEVMTestUpdatedCustodyTSSAddress, error) {
	event := new(GatewayEVMTestUpdatedCustodyTSSAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedGatewayTSSAddressIterator struct {
	Event *GatewayEVMTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUpdatedGatewayTSSAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestUpdatedGatewayTSSAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedGatewayTSSAddress struct {
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*GatewayEVMTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUpdatedGatewayTSSAddressIterator{contract: _GatewayEVMTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUpdatedGatewayTSSAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*GatewayEVMTestUpdatedGatewayTSSAddress, error) {
	event := new(GatewayEVMTestUpdatedGatewayTSSAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the GatewayEVMTest contract.
type GatewayEVMTestWhitelistedIterator struct {
	Event *GatewayEVMTestWhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestWhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWhitelisted represents a Whitelisted event raised by the GatewayEVMTest contract.
type GatewayEVMTestWhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWhitelisted(opts *bind.FilterOpts, token []common.Address) (*GatewayEVMTestWhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWhitelistedIterator{contract: _GatewayEVMTest.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWhitelisted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Whitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWhitelisted(log types.Log) (*GatewayEVMTestWhitelisted, error) {
	event := new(GatewayEVMTestWhitelisted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnIterator struct {
	Event *GatewayEVMTestWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWithdrawn represents a Withdrawn event raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawn struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWithdrawnIterator{contract: _GatewayEVMTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWithdrawn, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWithdrawn)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWithdrawn(log types.Log) (*GatewayEVMTestWithdrawn, error) {
	event := new(GatewayEVMTestWithdrawn)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndCalledIterator struct {
	Event *GatewayEVMTestWithdrawnAndCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWithdrawnAndCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestWithdrawnAndCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndCalled struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWithdrawnAndCalledIterator{contract: _GatewayEVMTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWithdrawnAndCalled, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWithdrawnAndCalled)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayEVMTestWithdrawnAndCalled, error) {
	event := new(GatewayEVMTestWithdrawnAndCalled)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndRevertedIterator struct {
	Event *GatewayEVMTestWithdrawnAndReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWithdrawnAndReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestWithdrawnAndReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndReverted struct {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWithdrawnAndRevertedIterator{contract: _GatewayEVMTest.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWithdrawnAndReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWithdrawnAndReverted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*GatewayEVMTestWithdrawnAndReverted, error) {
	event := new(GatewayEVMTestWithdrawnAndReverted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogIterator struct {
	Event *GatewayEVMTestLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLog represents a Log event raised by the GatewayEVMTest contract.
type GatewayEVMTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayEVMTestLogIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogIterator{contract: _GatewayEVMTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLog)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLog(log types.Log) (*GatewayEVMTestLog, error) {
	event := new(GatewayEVMTestLog)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogAddressIterator struct {
	Event *GatewayEVMTestLogAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogAddress represents a LogAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayEVMTestLogAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogAddressIterator{contract: _GatewayEVMTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogAddress(log types.Log) (*GatewayEVMTestLogAddress, error) {
	event := new(GatewayEVMTestLogAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArrayIterator struct {
	Event *GatewayEVMTestLogArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogArray represents a LogArray event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayEVMTestLogArrayIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogArrayIterator{contract: _GatewayEVMTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogArray)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogArray(log types.Log) (*GatewayEVMTestLogArray, error) {
	event := new(GatewayEVMTestLogArray)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray0Iterator struct {
	Event *GatewayEVMTestLogArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogArray0 represents a LogArray0 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayEVMTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogArray0Iterator{contract: _GatewayEVMTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogArray0)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogArray0(log types.Log) (*GatewayEVMTestLogArray0, error) {
	event := new(GatewayEVMTestLogArray0)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray1Iterator struct {
	Event *GatewayEVMTestLogArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogArray1 represents a LogArray1 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayEVMTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogArray1Iterator{contract: _GatewayEVMTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogArray1)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogArray1(log types.Log) (*GatewayEVMTestLogArray1, error) {
	event := new(GatewayEVMTestLogArray1)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytesIterator struct {
	Event *GatewayEVMTestLogBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogBytes represents a LogBytes event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayEVMTestLogBytesIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogBytesIterator{contract: _GatewayEVMTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogBytes)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogBytes(log types.Log) (*GatewayEVMTestLogBytes, error) {
	event := new(GatewayEVMTestLogBytes)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytes32Iterator struct {
	Event *GatewayEVMTestLogBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogBytes32 represents a LogBytes32 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayEVMTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogBytes32Iterator{contract: _GatewayEVMTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogBytes32)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogBytes32(log types.Log) (*GatewayEVMTestLogBytes32, error) {
	event := new(GatewayEVMTestLogBytes32)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogIntIterator struct {
	Event *GatewayEVMTestLogInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogInt represents a LogInt event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayEVMTestLogIntIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogIntIterator{contract: _GatewayEVMTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogInt)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogInt(log types.Log) (*GatewayEVMTestLogInt, error) {
	event := new(GatewayEVMTestLogInt)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedAddressIterator struct {
	Event *GatewayEVMTestLogNamedAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedAddressIterator{contract: _GatewayEVMTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayEVMTestLogNamedAddress, error) {
	event := new(GatewayEVMTestLogNamedAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArrayIterator struct {
	Event *GatewayEVMTestLogNamedArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedArray represents a LogNamedArray event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedArrayIterator{contract: _GatewayEVMTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedArray)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayEVMTestLogNamedArray, error) {
	event := new(GatewayEVMTestLogNamedArray)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray0Iterator struct {
	Event *GatewayEVMTestLogNamedArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedArray0Iterator{contract: _GatewayEVMTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedArray0)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayEVMTestLogNamedArray0, error) {
	event := new(GatewayEVMTestLogNamedArray0)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray1Iterator struct {
	Event *GatewayEVMTestLogNamedArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedArray1Iterator{contract: _GatewayEVMTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedArray1)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayEVMTestLogNamedArray1, error) {
	event := new(GatewayEVMTestLogNamedArray1)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytesIterator struct {
	Event *GatewayEVMTestLogNamedBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedBytesIterator{contract: _GatewayEVMTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedBytes)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayEVMTestLogNamedBytes, error) {
	event := new(GatewayEVMTestLogNamedBytes)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytes32Iterator struct {
	Event *GatewayEVMTestLogNamedBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedBytes32Iterator{contract: _GatewayEVMTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedBytes32)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayEVMTestLogNamedBytes32, error) {
	event := new(GatewayEVMTestLogNamedBytes32)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalIntIterator struct {
	Event *GatewayEVMTestLogNamedDecimalInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedDecimalInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedDecimalInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedDecimalIntIterator{contract: _GatewayEVMTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedDecimalInt)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayEVMTestLogNamedDecimalInt, error) {
	event := new(GatewayEVMTestLogNamedDecimalInt)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalUintIterator struct {
	Event *GatewayEVMTestLogNamedDecimalUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedDecimalUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedDecimalUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedDecimalUintIterator{contract: _GatewayEVMTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedDecimalUint)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayEVMTestLogNamedDecimalUint, error) {
	event := new(GatewayEVMTestLogNamedDecimalUint)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedIntIterator struct {
	Event *GatewayEVMTestLogNamedInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedInt represents a LogNamedInt event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedIntIterator{contract: _GatewayEVMTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedInt)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayEVMTestLogNamedInt, error) {
	event := new(GatewayEVMTestLogNamedInt)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedStringIterator struct {
	Event *GatewayEVMTestLogNamedString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedString represents a LogNamedString event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedStringIterator{contract: _GatewayEVMTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedString)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedString(log types.Log) (*GatewayEVMTestLogNamedString, error) {
	event := new(GatewayEVMTestLogNamedString)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedUintIterator struct {
	Event *GatewayEVMTestLogNamedUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogNamedUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedUint represents a LogNamedUint event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedUintIterator{contract: _GatewayEVMTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedUint)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayEVMTestLogNamedUint, error) {
	event := new(GatewayEVMTestLogNamedUint)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogStringIterator struct {
	Event *GatewayEVMTestLogString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogString represents a LogString event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayEVMTestLogStringIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogStringIterator{contract: _GatewayEVMTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogString)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogString(log types.Log) (*GatewayEVMTestLogString, error) {
	event := new(GatewayEVMTestLogString)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogUintIterator struct {
	Event *GatewayEVMTestLogUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogUint represents a LogUint event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayEVMTestLogUintIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogUintIterator{contract: _GatewayEVMTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogUint)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogUint(log types.Log) (*GatewayEVMTestLogUint, error) {
	event := new(GatewayEVMTestLogUint)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogsIterator struct {
	Event *GatewayEVMTestLogs // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogs)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMTestLogs)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogs represents a Logs event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayEVMTestLogsIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogsIterator{contract: _GatewayEVMTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogs)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "logs", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogs(log types.Log) (*GatewayEVMTestLogs, error) {
	event := new(GatewayEVMTestLogs)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
