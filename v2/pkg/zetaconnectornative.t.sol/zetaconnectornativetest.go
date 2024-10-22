// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornative

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

// ZetaConnectorNativeTestMetaData contains all meta data concerning the ZetaConnectorNativeTest contract.
var ZetaConnectorNativeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testTSSUpgrade\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfSenderIsNotTSSUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpgradeAndWithdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20Partial\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveNoParams\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevert\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedZetaConnectorTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5061cd678061003c6000396000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c8063a217fddf11610104578063ccb0e3f2116100a2578063e20c9f7111610071578063e20c9f7114610357578063e63ab1e91461035f578063fa7626d414610386578063fe574f841461039357600080fd5b8063ccb0e3f214610337578063d509b16c1461033f578063dcf7d03714610347578063de1cb76c1461034f57600080fd5b8063b0464fdc116100de578063b0464fdc14610307578063b5508aa91461030f578063ba414fa614610317578063c19099721461032f57600080fd5b8063a217fddf146102d0578063a783c789146102d8578063af298bb1146102ff57600080fd5b8063493465581161017c578063828320141161014b578063828320141461026957806385226c811461027157806385f438c114610286578063916a17c6146102bb57600080fd5b8063493465581461023c5780634df42da11461024457806352ff59391461024c57806366d9a9a01461025457600080fd5b80632ade3880116101b85780632ade38801461020f5780633cba0107146102245780633e5e3c231461022c5780633f7286f41461023457600080fd5b8063070f2ad0146101df5780630a9254e4146101e95780631ed7831c146101f1575b600080fd5b6101e761039b565b005b6101e761059b565b6101f9610dee565b6040516102069190619b1e565b60405180910390f35b610217610e50565b6040516102069190619bba565b6101e7610f92565b6101f9611755565b6101f96117b5565b6101e7611815565b6101e7611e59565b6101e7611fc9565b61025c612818565b6040516102069190619d20565b6101e761299a565b610279612bfa565b6040516102069190619dbe565b6102ad7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b604051908152602001610206565b6102c3612cca565b6040516102069190619e35565b6102ad600081565b6102ad7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101e7612dc5565b6102c3613211565b61027961330c565b61031f6133dc565b6040519015158152602001610206565b6101e76134b0565b6101e7613720565b6101e7614246565b6101e7614284565b6101e761492e565b6101f9614f8b565b6102ad7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b601f5461031f9060ff1681565b6101e7614feb565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561040d57600080fd5b505af1158015610421573d6000803e3d6000fd5b5050602654604080516001600160a01b039092166024830152600060448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506104ed9190600401619ecc565b600060405180830381600087803b15801561050757600080fd5b505af115801561051b573d6000803e3d6000fd5b5050602254602480546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529216935063950837aa9250015b600060405180830381600087803b15801561058157600080fd5b505af1158015610595573d6000803e3d6000fd5b50505050565b602480547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602580548216611234179055602680548216615678179055602780549091166198761790556040516105f990619a4b565b604080825260049082018190527f7a6574610000000000000000000000000000000000000000000000000000000060608301526080602083018190528201527f5a4554410000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f08015801561067d573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600e81527f4761746577617945564d2e736f6c00000000000000000000000000000000000060208201526026546024805493519186169082015260448101939093529216606482015260009161076f916084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b00000000000000000000000000000000000000000000000000000000179052615204565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0384811682029290921792839055604080518082018252601081527f4552433230437573746f64792e736f6c000000000000000000000000000000006020820152602654602480549351949096048516958401959095529383166044830152909116606482015291925061081591608401610712565b602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691909117909155604080518082018252601781527f5a657461436f6e6e6563746f724e61746976652e736f6c0000000000000000006020820152601f5460235460265460248054955161010090940487169084015290851660448301528416606482015291909216608482015291925061091a9160a40160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff8c8765e00000000000000000000000000000000000000000000000000000000179052615204565b602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03831617905560405190915061095c90619a58565b604051809103906000f080158015610978573d6000803e3d6000fd5b50602080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556026546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b158015610a2457600080fd5b505af1158015610a38573d6000803e3d6000fd5b5050602480546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d93506306447d56925001600060405180830381600087803b158015610aad57600080fd5b505af1158015610ac1573d6000803e3d6000fd5b5050601f546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015261010090920416925063ae7a3a6f9150602401600060405180830381600087803b158015610b2c57600080fd5b505af1158015610b40573d6000803e3d6000fd5b5050601f546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526101009092041692506310188aef9150602401600060405180830381600087803b158015610bab57600080fd5b505af1158015610bbf573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610c2157600080fd5b505af1158015610c35573d6000803e3d6000fd5b50506023546022546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152624c4b406024820152911692506340c10f199150604401600060405180830381600087803b158015610ca457600080fd5b505af1158015610cb8573d6000803e3d6000fd5b50506026546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b158015610d3c57600080fd5b505af1158015610d50573d6000803e3d6000fd5b5050604080516060810182526023546001600160a01b039081168252600160208084019182528451908101855260008152938301849052825160288054925167ffffffffffffffff1674010000000000000000000000000000000000000000027fffffffff00000000000000000000000000000000000000000000000000000000909316919093161717815590935091506029906105959082619fa7565b60606016805480602002602001604051908101604052809291908181526020018280548015610e4657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e28575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020016000905b82821015610f8957600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015610f72578382906000526020600020018054610ee590619f0e565b80601f0160208091040260200160405190810160405280929190818152602001828054610f1190619f0e565b8015610f5e5780601f10610f3357610100808354040283529160200191610f5e565b820191906000526020600020905b815481529060010190602001808311610f4157829003601f168201915b505050505081526020019060010190610ec6565b505050508152505081526020019060010190610e74565b50505050905090565b602354602554604051620186a0602482018190526001600160a01b039384166044830152929091166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905260235460255491516370a0823160e01b81526001600160a01b0392831660048201529293506000929116906370a0823190602401602060405180830381865afa158015611070573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611094919061a066565b90506110a1816000615223565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156110f2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611116919061a066565b601f54604080516001600160a01b036101009093048316602482015260448082018a905282518083039091018152606490910182526020810180517fa9059cbb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba3926111f192911690600090869060040161a07f565b600060405180830381600087803b15801561120b57600080fd5b505af115801561121f573d6000803e3d6000fd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156112b157600080fd5b505af11580156112c5573d6000803e3d6000fd5b5050601f54602354602554604080516101009094046001600160a01b039081168552602085018d9052928316908401521660608201527f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609250608001905060405180910390a16022546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156113b957600080fd5b505af11580156113cd573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d9150611412908990889061a0a7565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561147357600080fd5b505af1158015611487573d6000803e3d6000fd5b50506022546020546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef93506114df92909116908a9089908b9060040161a0c0565b600060405180830381600087803b1580156114f957600080fd5b505af115801561150d573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a08231906024015b602060405180830381865afa158015611561573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611585919061a066565b90506115918188615223565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156115e2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611606919061a066565b905061161b816116168a8761a128565b615223565b602354601f546020546040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081526101009092046001600160a01b0390811660048401529081166024830152600092169063dd62ed3e90604401602060405180830381865afa158015611693573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116b7919061a066565b90506116c4816000615223565b602354601f546040516370a0823160e01b81526101009091046001600160a01b03908116600483015260009216906370a0823190602401602060405180830381865afa158015611718573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061173c919061a066565b9050611749816000615223565b50505050505050505050565b60606018805480602002602001604051908101604052809291908181526020018280548015610e46576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610e28575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015610e46576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610e28575050505050905090565b6040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed701690000000000000000000000000000000000000000000000000000000017905260235460255493516370a0823160e01b8152620186a0946000949385936001600160a01b03908116936370a08231936118b793921691016001600160a01b0391909116815260200190565b602060405180830381865afa1580156118d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f8919061a066565b9050611905816000615223565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611956573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061197a919061a066565b601f54604080516001600160a01b036101009093048316602482015260448082018a905282518083039091018152606490910182526020810180517fa9059cbb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba392611a5592911690600090869060040161a07f565b600060405180830381600087803b158015611a6f57600080fd5b505af1158015611a83573d6000803e3d6000fd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015611b1557600080fd5b505af1158015611b29573d6000803e3d6000fd5b5050601f546040516101009091046001600160a01b031681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09250602001905060405180910390a16022546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611c0057600080fd5b505af1158015611c14573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d9150611c59908990889061a0a7565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611cba57600080fd5b505af1158015611cce573d6000803e3d6000fd5b50506022546020546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef9350611d2692909116908a9089908b9060040161a0c0565b600060405180830381600087803b158015611d4057600080fd5b505af1158015611d54573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015611da7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dcb919061a066565b9050611dd8816000615223565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611e29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e4d919061a066565b905061161b8185615223565b602480546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d569101600060405180830381600087803b158015611eca57600080fd5b505af1158015611ede573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015611f6757600080fd5b505af1158015611f7b573d6000803e3d6000fd5b50506022546040517f950837aa000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b03909116925063950837aa9150602401610567565b6022546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4600482015261432160248201819052916000916001600160a01b03909116906391d1485490604401602060405180830381865afa158015612058573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061207c919061a13b565b9050612087816152a3565b6022546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b03848116602483015260009216906391d1485490604401602060405180830381865afa158015612111573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612135919061a13b565b9050612140816152a3565b6022546026546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460048201526001600160a01b03918216602482015260009291909116906391d1485490604401602060405180830381865afa1580156121d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121f4919061a13b565b90506121ff8161531d565b6022546026546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b03918216602482015260009291909116906391d1485490604401602060405180830381865afa15801561228f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122b3919061a13b565b90506122be8161531d565b602480546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d569101600060405180830381600087803b15801561232f57600080fd5b505af1158015612343573d6000803e3d6000fd5b50506022546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156123d557600080fd5b505af11580156123e9573d6000803e3d6000fd5b5050602654604080516001600160a01b03928316815291891660208301527f33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e935001905060405180910390a16022546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0387811660048301529091169063950837aa90602401600060405180830381600087803b15801561249557600080fd5b505af11580156124a9573d6000803e3d6000fd5b5050505061252d85602260009054906101000a90046001600160a01b03166001600160a01b0316635b1125916040518163ffffffff1660e01b8152600401602060405180830381865afa158015612504573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612528919061a15d565b61536f565b6022546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460048201526001600160a01b038781166024830152909116906391d1485490604401602060405180830381865afa1580156125b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125da919061a13b565b93506125e58461531d565b6022546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b038781166024830152909116906391d1485490604401602060405180830381865afa15801561266e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612692919061a13b565b925061269d8361531d565b6022546026546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460048201526001600160a01b0391821660248201529116906391d1485490604401602060405180830381865afa158015612728573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061274c919061a13b565b9150612757826152a3565b6022546026546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b0391821660248201529116906391d1485490604401602060405180830381865afa1580156127e2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612806919061a13b565b9050612811816152a3565b5050505050565b6060601b805480602002602001604051908101604052809291908181526020016000905b82821015610f89578382906000526020600020906002020160405180604001604052908160008201805461286f90619f0e565b80601f016020809104026020016040519081016040528092919081815260200182805461289b90619f0e565b80156128e85780601f106128bd576101008083540402835291602001916128e8565b820191906000526020600020905b8154815290600101906020018083116128cb57829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561298257602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001906004019060208260030104928301926001038202915080841161292f5790505b5050505050815250508152602001906001019061283c565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a090600090819060250160408051808303601f19018152908290526024805463ca669fa760e01b84526001600160a01b03166004840152909250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa79101600060405180830381600087803b158015612a3757600080fd5b505af1158015612a4b573d6000803e3d6000fd5b505060248054604080516001600160a01b03909216928201929092527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250612b3a9190600401619ecc565b600060405180830381600087803b158015612b5457600080fd5b505af1158015612b68573d6000803e3d6000fd5b50506022546020546040517f057e0f250000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063057e0f259350612bc3929091169087908690889060289060040161a25b565b600060405180830381600087803b158015612bdd57600080fd5b505af1158015612bf1573d6000803e3d6000fd5b50505050505050565b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015610f89578382906000526020600020018054612c3d90619f0e565b80601f0160208091040260200160405190810160405280929190818152602001828054612c6990619f0e565b8015612cb65780601f10612c8b57610100808354040283529160200191612cb6565b820191906000526020600020905b815481529060010190602001808311612c9957829003601f168201915b505050505081526020019060010190612c1e565b6060601d805480602002602001604051908101604052809291908181526020016000905b82821015610f895760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015612dad57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411612d5a5790505b50505050508152505081526020019060010190612cee565b602280546040805160608101909152828152612e0f926001600160a01b039092169161cd1060208301396040805160208101909152600081526024546001600160a01b03166153d0565b6022546023546025546040516370a0823160e01b81526001600160a01b03918216600482015292811692620186a09260009283929116906370a0823190602401602060405180830381865afa158015612e6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e90919061a066565b9050612e9d816000615223565b6025546040516001600160a01b0390911660248201526044810184905260009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260235490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391612f7d916001600160a01b039190911690600090869060040161a07f565b600060405180830381600087803b158015612f9757600080fd5b505af1158015612fab573d6000803e3d6000fd5b50506040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b0388166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561303957600080fd5b505af115801561304d573d6000803e3d6000fd5b50506025546040518781526001600160a01b0390911692507f3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9915060200160405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156130ec57600080fd5b505af1158015613100573d6000803e3d6000fd5b50506025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526024810188905260448101879052908816925063106e62909150606401600060405180830381600087803b15801561317257600080fd5b505af1158015613186573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156131d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131fd919061a066565b90506132098186615223565b505050505050565b6060601c805480602002602001604051908101604052809291908181526020016000905b82821015610f895760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156132f457602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116132a15790505b50505050508152505081526020019060010190613235565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015610f8957838290600052602060002001805461334f90619f0e565b80601f016020809104026020016040519081016040528092919081815260200182805461337b90619f0e565b80156133c85780601f1061339d576101008083540402835291602001916133c8565b820191906000526020600020905b8154815290600101906020018083116133ab57829003601f168201915b505050505081526020019060010190613330565b60085460009060ff16156133f4575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015613485573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134a9919061a066565b1415905090565b602354602554604051620186a0602482018190526001600160a01b039384166044830152929091166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905260248054915163ca669fa760e01b81526001600160a01b039092166004830152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa79101600060405180830381600087803b15801561359757600080fd5b505af11580156135ab573d6000803e3d6000fd5b505060248054604080516001600160a01b03909216928201929092527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061369a9190600401619ecc565b600060405180830381600087803b1580156136b457600080fd5b505af11580156136c8573d6000803e3d6000fd5b50506022546020546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef9350612bc3929091169087908690889060040161a0c0565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090600090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561378157600080fd5b505af1158015613795573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506138809190600401619ecc565b600060405180830381600087803b15801561389a57600080fd5b505af11580156138ae573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561390257600080fd5b505af1158015613916573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561397357600080fd5b505af1158015613987573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250613a729190600401619ecc565b600060405180830381600087803b158015613a8c57600080fd5b505af1158015613aa0573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015613af457600080fd5b505af1158015613b08573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015613b6557600080fd5b505af1158015613b79573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b158015613bcd57600080fd5b505af1158015613be1573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd93c0665000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015613c6a57600080fd5b505af1158015613c7e573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015613cdb57600080fd5b505af1158015613cef573d6000803e3d6000fd5b50506022546025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101879052604481018690529116925063106e62909150606401600060405180830381600087803b158015613d6357600080fd5b505af1158015613d77573d6000803e3d6000fd5b50506024805460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063ca669fa7925001600060405180830381600087803b158015613dd357600080fd5b505af1158015613de7573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015613e3b57600080fd5b505af1158015613e4f573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a08231906024015b602060405180830381865afa158015613ea3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ec7919061a066565b9050613ed4816000615223565b6025546040516001600160a01b0390911660248201526044810184905260009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260235490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391613fb4916001600160a01b039190911690600090869060040161a07f565b600060405180830381600087803b158015613fce57600080fd5b505af1158015613fe2573d6000803e3d6000fd5b50506022546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561407457600080fd5b505af1158015614088573d6000803e3d6000fd5b50506025546040518781526001600160a01b0390911692507f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5915060200160405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561412757600080fd5b505af115801561413b573d6000803e3d6000fd5b50506022546025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101899052604481018890529116925063106e62909150606401600060405180830381600087803b1580156141af57600080fd5b505af11580156141c3573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015614216573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061423a919061a066565b90506128118186615223565b6023546025546040516370a0823160e01b81526001600160a01b039182166004820152620186a09260009283929116906370a0823190602401613e86565b602354602554604051620186a0602482018190526001600160a01b039384166044830152929091166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc51316910000000000000000000000000000000000000000000000000000000017905260235460255491516370a0823160e01b81526001600160a01b0392831660048201529293506000929116906370a0823190602401602060405180830381865afa158015614362573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614386919061a066565b9050614393816000615223565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156143e4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614408919061a066565b601f54604080516001600160a01b036101009093048316602482015260448082018a905282518083039091018152606490910182526020810180517fa9059cbb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba3926144e392911690600090869060040161a07f565b600060405180830381600087803b1580156144fd57600080fd5b505af1158015614511573d6000803e3d6000fd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156145a357600080fd5b505af11580156145b7573d6000803e3d6000fd5b5050601f547f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60925061010090046001600160a01b031690506145fa60028961a2a7565b602354602554604080516001600160a01b03958616815260208101949094529184168383015292909216606082015290519081900360800190a16022546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156146c257600080fd5b505af11580156146d6573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d915061471b908990889061a0a7565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561477c57600080fd5b505af1158015614790573d6000803e3d6000fd5b50506022546020546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef93506147e892909116908a9089908b9060040161a0c0565b600060405180830381600087803b15801561480257600080fd5b505af1158015614816573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015614869573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061488d919061a066565b905061489e8161161660028a61a2a7565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156148ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614913919061a066565b905061161b8161492460028b61a2a7565b611616908761a128565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a090600090819060250160408051808303601f19018152908290526023546020546370a0823160e01b84526001600160a01b0390811660048501529193506000929116906370a0823190602401602060405180830381865afa1580156149c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906149e8919061a066565b90506149f5816000615223565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015614a46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a6a919061a066565b601f54604080516001600160a01b036101009093048316602482015260448082018a905282518083039091018152606490910182526020810180517fa9059cbb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba392614b4592911690600090869060040161a07f565b600060405180830381600087803b158015614b5f57600080fd5b505af1158015614b73573d6000803e3d6000fd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015614c0557600080fd5b505af1158015614c19573d6000803e3d6000fd5b505050507f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e601f60019054906101000a90046001600160a01b03166028604051614c6492919061a2e2565b60405180910390a1601f546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015614cfe57600080fd5b505af1158015614d12573d6000803e3d6000fd5b50506023546020546040516001600160a01b039283169450911691507f1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b143690614d60908a90899060289061a304565b60405180910390a36022546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015614df657600080fd5b505af1158015614e0a573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c9150614e52908990889060289061a304565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614eb357600080fd5b505af1158015614ec7573d6000803e3d6000fd5b50506022546020546040517f057e0f250000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063057e0f259350614f2292909116908a9089908b9060289060040161a25b565b600060405180830381600087803b158015614f3c57600080fd5b505af1158015614f50573d6000803e3d6000fd5b50506023546020546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401611544565b60606015805480602002602001604051908101604052809291908181526020018280548015610e46576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610e28575050505050905090565b6024805460405163ca669fa760e01b81526001600160a01b039091166004820152620186a091600091737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa79101600060405180830381600087803b15801561504b57600080fd5b505af115801561505f573d6000803e3d6000fd5b505060248054604080516001600160a01b03909216928201929092527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061514e9190600401619ecc565b600060405180830381600087803b15801561516857600080fd5b505af115801561517c573d6000803e3d6000fd5b50506022546025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101879052604481018690529116925063106e62909150606401600060405180830381600087803b1580156151f057600080fd5b505af1158015613209573d6000803e3d6000fd5b600061520e619a65565b6152198484836153e5565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015b60006040518083038186803b15801561528f57600080fd5b505afa158015613209573d6000803e3d6000fd5b6040517fa59828850000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a5982885906024015b60006040518083038186803b15801561530957600080fd5b505afa158015612811573d6000803e3d6000fd5b6040517f0c9fd5810000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90630c9fd581906024016152f1565b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f690604401615277565b6153d8619a65565b6128118585858486615460565b6000806153f28584615560565b90506154556040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f7879000000815250828660405160200161544092919061a32f565b6040516020818303038152906040528561556c565b9150505b9392505050565b6040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201528190737109709ecfa91a80626ff3989d68f67f5b1dd12d9081906306447d5690602401600060405180830381600087803b1580156154d257600080fd5b505af19250505080156154e3575060015b6154f8576154f38787878761559a565b612bf1565b6155048787878761559a565b806001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561553f57600080fd5b505af1158015615553573d6000803e3d6000fd5b5050505050505050505050565b600061545983836155b3565b60c081015151600090156155905761558984848460c001516155ce565b9050615459565b6155898484615774565b60006155a6848361585f565b905061281185828561586b565b60006155bf8383615c35565b6154598383602001518461556c565b6000806155d9615c45565b905060006155e78683615d18565b905060006155fe82606001518360200151856161be565b9050600061560e838389896163d0565b9050600061561b8261724d565b602081015181519192509060030b1561568e5789826040015160405160200161564592919061a351565b60408051601f19818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261568591600401619ecc565b60405180910390fd5b60006156d16040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a20000000000000000000000081525083600161741c565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d90615724908490600401619ecc565b602060405180830381865afa158015615741573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615765919061a15d565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc925906157c9908790600401619ecc565b600060405180830381865afa1580156157e6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261580e919081019061a48b565b9050600061583c828560405160200161582892919061a4c0565b60405160208183030381529060405261761c565b90506001600160a01b03811661521957848460405160200161564592919061a4ef565b60006155bf838361762f565b6040517f667f9d700000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90600090829063667f9d7090604401602060405180830381865afa158015615907573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061592b919061a066565b905080615ad257600061593d8661763b565b604080518082018252600581527f352e302e30000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506159c8905b60408051808201825260008082526020918201528151808301909252845182528085019082015290617733565b806159d4575060008451115b15615a57576040517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b03871690634f1ef28690615a20908890889060040161a32f565b600060405180830381600087803b158015615a3a57600080fd5b505af1158015615a4e573d6000803e3d6000fd5b50505050615acc565b6040517f3659cfe60000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152871690633659cfe690602401600060405180830381600087803b158015615ab357600080fd5b505af1158015615ac7573d6000803e3d6000fd5b505050505b50612811565b806000615ade8261763b565b604080518082018252600581527f352e302e3000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150615b409061599b565b80615b4c575060008551115b15615bd1576040517f9623609d0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690639623609d90615b9a908a908a908a9060040161a59a565b600060405180830381600087803b158015615bb457600080fd5b505af1158015615bc8573d6000803e3d6000fd5b50505050612bf1565b6040517f99a88ec40000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015287811660248301528316906399a88ec490604401600060405180830381600087803b15801561553f57600080fd5b615c4182826000617747565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c90615ccc90849060040161a5cb565b600060405180830381865afa158015615ce9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615d11919081019061a612565b9250505090565b615d4a6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d9050615d956040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b615d9e8561784a565b60208201526000615dae86617c2f565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa158015615df0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615e18919081019061a612565b86838560200151604051602001615e32949392919061a65b565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb1190615e8a908590600401619ecc565b600060405180830381865afa158015615ea7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615ecf919081019061a612565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f690615f1790849060040161a75f565b602060405180830381865afa158015615f34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615f58919061a13b565b615f6d5781604051602001615645919061a7b1565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890615fb290849060040161a843565b600060405180830381865afa158015615fcf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615ff7919081019061a612565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f69061603e90849060040161a895565b602060405180830381865afa15801561605b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061607f919061a13b565b15616114576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906160c990849060040161a895565b600060405180830381865afa1580156160e6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261610e919081019061a612565b60408501525b846001600160a01b03166349c4fac8828660000151604051602001616139919061a8e7565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161616592919061a953565b600060405180830381865afa158015616182573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526161aa919081019061a612565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b60608152602001906001900390816161da5790505090506040518060400160405280600481526020017f67726570000000000000000000000000000000000000000000000000000000008152508160008151811061623a5761623a61a978565b60200260200101819052506040518060400160405280600381526020017f2d726c00000000000000000000000000000000000000000000000000000000008152508160018151811061628e5761628e61a978565b6020026020010181905250846040516020016162aa919061a9a7565b604051602081830303815290604052816002815181106162cc576162cc61a978565b6020026020010181905250826040516020016162e8919061aa13565b6040516020818303038152906040528160038151811061630a5761630a61a978565b602002602001018190525060006163208261724d565b602080820151604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000081850190815282518084018452600080825290860152825180840190935290518252928101929092529192506163b19060408051808201825260008082526020918201528151808301909252845182528085019082015290617eb2565b6163c65785604051602001615645919061aa54565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015616420565b511590565b616594578260200151156164dc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401615685565b8260c0015115616594576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401615685565b6040805160ff8082526120008201909252600091816020015b60608152602001906001900390816165ad57905050905060006040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806166089061aae5565b935060ff168151811061661d5761661d61a978565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e370000000000000000000000000000000000000081525060405160200161666e919061ab04565b6040516020818303038152906040528282806166899061aae5565b935060ff168151811061669e5761669e61a978565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806166eb9061aae5565b935060ff16815181106167005761670061a978565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d6500000000000000000000000000000000000081525082828061674d9061aae5565b935060ff16815181106167625761676261a978565b6020026020010181905250876020015182828061677e9061aae5565b935060ff16815181106167935761679361a978565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e7472616374506174680000000000000000000000000000000000008152508282806167e09061aae5565b935060ff16815181106167f5576167f561a978565b60209081029190910101528751828261680d8161aae5565b935060ff16815181106168225761682261a978565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e4964000000000000000000000000000000000000000000000081525082828061686f9061aae5565b935060ff16815181106168845761688461a978565b602002602001018190525061689846617f13565b82826168a38161aae5565b935060ff16815181106168b8576168b861a978565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806169059061aae5565b935060ff168151811061691a5761691a61a978565b6020026020010181905250868282806169329061aae5565b935060ff16815181106169475761694761a978565b6020908102919091010152855115616a6e5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826169988161aae5565b935060ff16815181106169ad576169ad61a978565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906169fd908990600401619ecc565b600060405180830381865afa158015616a1a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616a42919081019061a612565b8282616a4d8161aae5565b935060ff1681518110616a6257616a6261a978565b60200260200101819052505b846020015115616b3e5760408051808201909152601281527f2d2d766572696679536f75726365436f6465000000000000000000000000000060208201528282616ab78161aae5565b935060ff1681518110616acc57616acc61a978565b60200260200101819052506040518060400160405280600581526020017f66616c7365000000000000000000000000000000000000000000000000000000815250828280616b199061aae5565b935060ff1681518110616b2e57616b2e61a978565b6020026020010181905250616d05565b616b7661641b8660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b616c095760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282616bb98161aae5565b935060ff1681518110616bce57616bce61a978565b60200260200101819052508460a00151604051602001616bee919061a9a7565b604051602081830303815290604052828280616b199061aae5565b8460c00151158015616c4c575060408089015181518083018352600080825260209182015282518084019093528151835290810190820152616c4a90511590565b155b15616d055760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282616c908161aae5565b935060ff1681518110616ca557616ca561a978565b6020026020010181905250616cb988617fb3565b604051602001616cc9919061a9a7565b604051602081830303815290604052828280616ce49061aae5565b935060ff1681518110616cf957616cf961a978565b60200260200101819052505b60408086015181518083018352600080825260209182015282518084019093528151835290810190820152616d3990511590565b616dce5760408051808201909152600b81527f2d2d72656c61796572496400000000000000000000000000000000000000000060208201528282616d7c8161aae5565b935060ff1681518110616d9157616d9161a978565b60200260200101819052508460400151828280616dad9061aae5565b935060ff1681518110616dc257616dc261a978565b60200260200101819052505b606085015115616eef5760408051808201909152600681527f2d2d73616c74000000000000000000000000000000000000000000000000000060208201528282616e178161aae5565b935060ff1681518110616e2c57616e2c61a978565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015616e9b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616ec3919081019061a612565b8282616ece8161aae5565b935060ff1681518110616ee357616ee361a978565b60200260200101819052505b60e08501515115616f965760408051808201909152600a81527f2d2d6761734c696d69740000000000000000000000000000000000000000000060208201528282616f398161aae5565b935060ff1681518110616f4e57616f4e61a978565b6020026020010181905250616f6a8560e0015160000151617f13565b8282616f758161aae5565b935060ff1681518110616f8a57616f8a61a978565b60200260200101819052505b60e085015160200151156170405760408051808201909152600a81527f2d2d67617350726963650000000000000000000000000000000000000000000060208201528282616fe38161aae5565b935060ff1681518110616ff857616ff861a978565b60200260200101819052506170148560e0015160200151617f13565b828261701f8161aae5565b935060ff16815181106170345761703461a978565b60200260200101819052505b60e085015160400151156170ea5760408051808201909152600e81527f2d2d6d61784665655065724761730000000000000000000000000000000000006020820152828261708d8161aae5565b935060ff16815181106170a2576170a261a978565b60200260200101819052506170be8560e0015160400151617f13565b82826170c98161aae5565b935060ff16815181106170de576170de61a978565b60200260200101819052505b60e085015160600151156171945760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826171378161aae5565b935060ff168151811061714c5761714c61a978565b60200260200101819052506171688560e0015160600151617f13565b82826171738161aae5565b935060ff16815181106171885761718861a978565b60200260200101819052505b60008160ff1667ffffffffffffffff8111156171b2576171b2619edf565b6040519080825280602002602001820160405280156171e557816020015b60608152602001906001900390816171d05790505b50905060005b8260ff168160ff16101561723e57838160ff168151811061720e5761720e61a978565b6020026020010151828260ff168151811061722b5761722b61a978565b60209081029190910101526001016171eb565b5093505050505b949350505050565b6172746040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c916172fa9186910161ab6f565b600060405180830381865afa158015617317573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261733f919081019061a612565b9050600061734d8683618aa2565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b815260040161737d9190619dbe565b6000604051808303816000875af115801561739c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526173c4919081019061abb6565b805190915060030b158015906173dd5750602081015151155b80156173ec5750604081015151155b156163c657816000815181106174045761740461a978565b6020026020010151604051602001615645919061ac6c565b606060006174518560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925286518252808701908201529091506174889082905b90618bf7565b156175e5576000617505826174ff846174f96174cb8a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b90618c1e565b90618c80565b604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150617569908290618bf7565b156175d357604080518082018252600181527f0a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526175d0905b8290618d05565b90505b6175dc81618d2b565b92505050615459565b82156175fe57848460405160200161564592919061ae58565b5050604080516020810190915260008152615459565b509392505050565b6000808251602084016000f09392505050565b615c4182826001617747565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fad3cb1cc00000000000000000000000000000000000000000000000000000000179052905160609160009182916001600160a01b038616916176b0919061aeff565b600060405180830381855afa9150503d80600081146176eb576040519150601f19603f3d011682016040523d82523d6000602084013e6176f0565b606091505b5091509150818015617703575060208151115b1561771c5780806020019051810190617245919061a612565b505060408051602081019091526000815292915050565b600061773f8383618d94565b159392505050565b8160a001511561775657505050565b6000617763848484618e6f565b905060006177708261724d565b602081015181519192509060030b15801561780c5750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261780c90604080518082018252600080825260209182015281518083019092528451825280850190820152617482565b1561781957505050505050565b60408201515115617839578160400151604051602001615645919061af1b565b80604051602001615645919061af79565b6060600061787f8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506178e4905b8290617eb2565b1561795357604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526154599061794e90839061940a565b618d2b565b604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526179b5905b8290619494565b600103617a8257604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617a1b906175c9565b50604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526154599061794e905b8390618d05565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617ae1906178dd565b15617c1857604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290617b4990839061952e565b905060008160018351617b5c919061a128565b81518110617b6c57617b6c61a978565b60200260200101519050617c0f61794e617be26040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925285518252808601908201529061940a565b95945050505050565b82604051602001615645919061afe4565b50919050565b60606000617c648360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150617cc6906178dd565b15617cd45761545981618d2b565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617d33906179ae565b600103617d9d57604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526154599061794e90617a7b565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617dfc906178dd565b15617c1857604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290617e6490839061952e565b9050600181511115617ea0578060028251617e7f919061a128565b81518110617e8f57617e8f61a978565b602002602001015192505050919050565b5082604051602001615645919061afe4565b805182516000911115617ec75750600061521d565b81518351602085015160009291617edd9161b0c2565b617ee7919061a128565b905082602001518103617efe57600191505061521d565b82516020840151819020912014905092915050565b60606000617f20836195d3565b600101905060008167ffffffffffffffff811115617f4057617f40619edf565b6040519080825280601f01601f191660200182016040528015617f6a576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084617f7457509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e534544000000000000000000000000000000000000000000008184019081528551808701875283815284019290925284518086019095525184529083015260609161803f905b8290617733565b1561807f57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e73650000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526180de90618038565b1561811e57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261817d90618038565b156181bd57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261821c90618038565b806182815750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261828190618038565b156182c157505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261832090618038565b806183855750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261838590618038565b156183c557505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261842490618038565b806184895750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261848990618038565b156184c957505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261852890618038565b8061858d5750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261858d90618038565b156185cd57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261862c90618038565b1561866c57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c617573650000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526186cb90618038565b1561870b57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261876a90618038565b156187aa57505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261880990618038565b1561884957505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e3000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526188a890618038565b156188e857505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261894790618038565b806189ac5750604080518082018252601181527f4147504c2d332e302d6f722d6c61746572000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526189ac90618038565b156189ec57505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e3100000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618a4b90618038565b15618a8b57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151615645929060200161b0d5565b60608060005b8451811015618b2d5781858281518110618ac457618ac461a978565b6020026020010151604051602001618add92919061a4c0565b604051602081830303815290604052915060018551618afc919061a128565b8114618b255781604051602001618b13919061b23e565b60405160208183030381529060405291505b600101618aa8565b5060408051600380825260808201909252600091816020015b6060815260200190600190039081618b465790505090508381600081518110618b7157618b7161a978565b60200260200101819052506040518060400160405280600281526020017f2d6300000000000000000000000000000000000000000000000000000000000081525081600181518110618bc557618bc561a978565b60200260200101819052508181600281518110618be457618be461a978565b6020908102919091010152949350505050565b6020808301518351835192840151600093618c1592918491906196b5565b14159392505050565b60408051808201909152600080825260208201526000618c5084600001518560200151856000015186602001516197c6565b9050836020015181618c62919061a128565b84518590618c7190839061a128565b90525060208401525090919050565b6040805180820190915260008082526020820152815183511015618ca557508161521d565b6020808301519084015160019114618ccc5750815160208481015190840151829020919020145b8015618cfd57825184518590618ce390839061a128565b9052508251602085018051618cf990839061b0c2565b9052505b509192915050565b6040805180820190915260008082526020820152618d248383836198e6565b5092915050565b60606000826000015167ffffffffffffffff811115618d4c57618d4c619edf565b6040519080825280601f01601f191660200182016040528015618d76576020820181803683370190505b5090506000602082019050618d248185602001518660000151619991565b8151815160009190811115618da7575081515b6020808501519084015160005b83811015618e605782518251808214618e30576000196020871015618e0f57600184618de189602061a128565b618deb919061b0c2565b618df690600861b27f565b618e0190600261b37d565b618e0b919061a128565b1990505b8181168382168181039114618e2d57975061521d9650505050505050565b50505b618e3b60208661b0c2565b9450618e4860208561b0c2565b93505050602081618e59919061b0c2565b9050618db4565b50845186516163c6919061b389565b60606000618e7b615c45565b6040805160ff808252612000820190925291925060009190816020015b6060815260200190600190039081618e9857905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280618ef39061aae5565b935060ff1681518110618f0857618f0861a978565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e3300000000000000000000000000000000000000000000000000815250604051602001618f59919061b3a9565b604051602081830303815290604052828280618f749061aae5565b935060ff1681518110618f8957618f8961a978565b60200260200101819052506040518060400160405280600881526020017f76616c6964617465000000000000000000000000000000000000000000000000815250828280618fd69061aae5565b935060ff1681518110618feb57618feb61a978565b602002602001018190525082604051602001619007919061aa13565b6040516020818303038152906040528282806190229061aae5565b935060ff16815181106190375761903761a978565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e7472616374000000000000000000000000000000000000000000008152508282806190849061aae5565b935060ff16815181106190995761909961a978565b60200260200101819052506190ae8784619a0b565b82826190b98161aae5565b935060ff16815181106190ce576190ce61a978565b60209081029190910101528551511561917a5760408051808201909152600b81527f2d2d7265666572656e6365000000000000000000000000000000000000000000602082015282826191208161aae5565b935060ff16815181106191355761913561a978565b602002602001018190525061914e866000015184619a0b565b82826191598161aae5565b935060ff168151811061916e5761916e61a978565b60200260200101819052505b8560800151156191e85760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b0000000000000000602082015282826191c38161aae5565b935060ff16815181106191d8576191d861a978565b602002602001018190525061924e565b841561924e5760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261922d8161aae5565b935060ff16815181106192425761924261a978565b60200260200101819052505b604086015151156192ea5760408051808201909152600d81527f2d2d756e73616665416c6c6f7700000000000000000000000000000000000000602082015282826192988161aae5565b935060ff16815181106192ad576192ad61a978565b602002602001018190525085604001518282806192c99061aae5565b935060ff16815181106192de576192de61a978565b60200260200101819052505b8560600151156193545760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826193338161aae5565b935060ff16815181106193485761934861a978565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561937257619372619edf565b6040519080825280602002602001820160405280156193a557816020015b60608152602001906001900390816193905790505b50905060005b8260ff168160ff1610156193fe57838160ff16815181106193ce576193ce61a978565b6020026020010151828260ff16815181106193eb576193eb61a978565b60209081029190910101526001016193ab565b50979650505050505050565b604080518082019091526000808252602082015281518351101561942f57508161521d565b815183516020850151600092916194459161b0c2565b61944f919061a128565b60208401519091506001908214619470575082516020840151819020908220145b801561948b5783518551869061948790839061a128565b9052505b50929392505050565b60008082600001516194b885600001518660200151866000015187602001516197c6565b6194c2919061b0c2565b90505b835160208501516194d6919061b0c2565b8111618d2457816194e68161b3ee565b925050826000015161951d856020015183619501919061a128565b865161950d919061a128565b83866000015187602001516197c6565b619527919061b0c2565b90506194c5565b6060600061953c8484619494565b61954790600161b0c2565b67ffffffffffffffff81111561955f5761955f619edf565b60405190808252806020026020018201604052801561959257816020015b606081526020019060019003908161957d5790505b50905060005b8151811015617614576195ae61794e8686618d05565b8282815181106195c0576195c061a978565b6020908102919091010152600101619598565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061961c577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310619648576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061966657662386f26fc10000830492506010015b6305f5e100831061967e576305f5e100830492506008015b612710831061969257612710830492506004015b606483106196a4576064830492506002015b600a831061521d5760010192915050565b6000808584116197bc576020841161976857600084156197005760016196dc86602061a128565b6196e790600861b27f565b6196f290600261b37d565b6196fc919061a128565b1990505b835181168561970f898961b0c2565b619719919061a128565b805190935082165b8181146197535787841161973b5787945050505050617245565b836197458161b408565b945050828451169050619721565b61975d878561b0c2565b945050505050617245565b838320619775858861a128565b61977f908761b0c2565b91505b8582106197ba578482208082036197a75761979d868461b0c2565b9350505050617245565b6197b260018461a128565b925050619782565b505b5092949350505050565b600083818685116198d1576020851161988057600085156198125760016197ee87602061a128565b6197f990600861b27f565b61980490600261b37d565b61980e919061a128565b1990505b845181166000876198238b8b61b0c2565b61982d919061a128565b855190915083165b8281146198725781861061985a5761984d8b8b61b0c2565b9650505050505050617245565b856198648161b3ee565b965050838651169050619835565b859650505050505050617245565b508383206000905b619892868961a128565b82116198cf578583208082036198ae5783945050505050617245565b6198b960018561b0c2565b93505081806198c79061b3ee565b925050619888565b505b6198db878761b0c2565b979650505050505050565b6040805180820190915260008082526020820152600061991885600001518660200151866000015187602001516197c6565b602080870180519186019190915251909150619934908261a128565b835284516020860151619947919061b0c2565b81036199565760008552619988565b83518351619964919061b0c2565b8551869061997390839061a128565b9052508351619982908261b0c2565b60208601525b50909392505050565b602081106199c957815183526199a860208461b0c2565b92506199b560208361b0c2565b91506199c260208261a128565b9050619991565b60001981156199f85760016199df83602061a128565b6199eb9061010061b37d565b6199f5919061a128565b90505b9151835183169219169190911790915250565b60606000619a198484615d18565b8051602080830151604051939450619a339390910161b41f565b60405160208183030381529060405291505092915050565b610c9f8061b47883390190565b610bf98061c11783390190565b6040518060e00160405280606081526020016060815260200160608152602001600015158152602001600015158152602001600015158152602001619aa8619aad565b905290565b60405180610100016040528060001515815260200160001515815260200160608152602001600080191681526020016060815260200160608152602001600015158152602001619aa86040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b81811015619b5f5783516001600160a01b0316835260209384019390920191600101619b38565b509095945050505050565b60005b83811015619b85578181015183820152602001619b6d565b50506000910152565b60008151808452619ba6816020860160208601619b6a565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015619cb6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b81811015619c9c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8503018352619c86848651619b8e565b6020958601959094509290920191600101619c4c565b509197505050602094850194929092019150600101619be2565b50929695505050505050565b600081518084526020840193506020830160005b82811015619d165781517fffffffff0000000000000000000000000000000000000000000000000000000016865260209586019590910190600101619cd6565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015619cb6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160408752619d8c6040880182619b8e565b9050602082015191508681036020880152619da78183619cc2565b965050506020938401939190910190600101619d48565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015619cb6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452619e20858351619b8e565b94506020938401939190910190600101619de6565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015619cb6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b0381511686526020810151905060406020870152619eb66040870182619cc2565b9550506020938401939190910190600101619e5d565b6020815260006154596020830184619b8e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c90821680619f2257607f821691505b602082108103617c29577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f821115619fa257806000526020600020601f840160051c81016020851015619f825750805b601f840160051c820191505b818110156128115760008155600101619f8e565b505050565b815167ffffffffffffffff811115619fc157619fc1619edf565b619fd581619fcf8454619f0e565b84619f5b565b6020601f82116001811461a0095760008315619ff15750848201515b600019600385901b1c1916600184901b178455612811565b600084815260208120601f198516915b8281101561a039578785015182556020948501946001909201910161a019565b508482101561a0575786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60006020828403121561a07857600080fd5b5051919050565b6001600160a01b0384168152826020820152606060408201526000617c0f6060830184619b8e565b8281526040602082015260006172456040830184619b8e565b6001600160a01b038516815283602082015260806040820152600061a0e86080830185619b8e565b905082606083015295945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561521d5761521d61a0f9565b60006020828403121561a14d57600080fd5b8151801515811461545957600080fd5b60006020828403121561a16f57600080fd5b81516001600160a01b038116811461545957600080fd5b600081546001600160a01b038116845267ffffffffffffffff8160a01c1660208501525060018201606060408501526000815461a1c281619f0e565b806060880152600182166000811461a1e1576001811461a21b5761a24f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083166080890152608082151560051b890101935061a24f565b84600052602060002060005b8381101561a2465781548a82016080015260019091019060200161a227565b89016080019450505b50919695505050505050565b6001600160a01b038616815284602082015260a06040820152600061a28360a0830186619b8e565b846060840152828103608084015261a29b818561a186565b98975050505050505050565b60008261a2dd577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6001600160a01b0383168152604060208201526000617245604083018461a186565b83815260606020820152600061a31d6060830185619b8e565b82810360408401526163c6818561a186565b6001600160a01b03831681526040602082015260006172456040830184619b8e565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161a38981601a850160208801619b6a565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161a3c681601c840160208801619b6a565b01601c01949350505050565b6040516060810167ffffffffffffffff8111828210171561a3f55761a3f5619edf565b60405290565b60008067ffffffffffffffff84111561a4165761a416619edf565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561a4455761a445619edf565b60405283815290508082840185101561a45d57600080fd5b617614846020830185619b6a565b600082601f83011261a47c57600080fd5b6154598383516020850161a3fb565b60006020828403121561a49d57600080fd5b815167ffffffffffffffff81111561a4b457600080fd5b6152198482850161a46b565b6000835161a4d2818460208801619b6a565b83519083019061a4e6818360208801619b6a565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161a52781601a850160208801619b6a565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161a564816033840160208801619b6a565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b6001600160a01b03841681526001600160a01b0383166020820152606060408201526000617c0f6060830184619b8e565b60408152600b60408201527f464f554e4452595f4f555400000000000000000000000000000000000000000060608201526080602082015260006154596080830184619b8e565b60006020828403121561a62457600080fd5b815167ffffffffffffffff81111561a63b57600080fd5b8201601f8101841361a64c57600080fd5b6152198482516020840161a3fb565b6000855161a66d818460208a01619b6a565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161a6a7816001840160208a01619b6a565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161a6e5816002840160208901619b6a565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161a727816002840160208801619b6a565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061a7726040830184619b8e565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161a7e981601f850160208701619b6a565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061a8566040830184619b8e565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061a8a86040830184619b8e565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161a91f816014850160208701619b6a565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061a9666040830185619b8e565b82810360208401526154558185619b8e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161a9df816001850160208701619b6a565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161aa25818460208701619b6a565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161aad881604b850160208701619b6a565b91909101604b0192915050565b600060ff821660ff810361aafb5761aafb61a0f9565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161ab62816029850160208701619b6a565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f504154480000000000000000000060608201526080602082015260006154596080830184619b8e565b60006020828403121561abc857600080fd5b815167ffffffffffffffff81111561abdf57600080fd5b82016060818503121561abf157600080fd5b61abf961a3d2565b81518060030b811461ac0a57600080fd5b8152602082015167ffffffffffffffff81111561ac2657600080fd5b61ac328682850161a46b565b602083015250604082015167ffffffffffffffff81111561ac5257600080fd5b61ac5e8682850161a46b565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161acca816021850160208701619b6a565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161aeb6816021850160208801619b6a565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161aef381602e840160208801619b6a565b01602e01949350505050565b6000825161af11818460208701619b6a565b9190910192915050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161ab62816029850160208701619b6a565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161afd7816022850160208701619b6a565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161b01c81600e850160208701619b6a565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b8082018082111561521d5761521d61a0f9565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161b10d816018850160208801619b6a565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161b14a81601c840160208801619b6a565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161b250818460208701619b6a565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b808202811582820484141761521d5761521d61a0f9565b6001815b600184111561b2d15780850481111561b2b55761b2b561a0f9565b600184161561b2c357908102905b60019390931c92800261b29a565b935093915050565b60008261b2e85750600161521d565b8161b2f55750600061521d565b816001811461b30b576002811461b3155761b331565b600191505061521d565b60ff84111561b3265761b32661a0f9565b50506001821b61521d565b5060208310610133831016604e8410600b841016171561b354575081810a61521d565b61b361600019848461b296565b806000190482111561b3755761b37561a0f9565b029392505050565b6000615459838361b2d9565b8181036000831280158383131683831282161715618d2457618d2461a0f9565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161b3e181601c850160208701619b6a565b91909101601c0192915050565b6000600019820361b4015761b40161a0f9565b5060010190565b60008161b4175761b41761a0f9565b506000190190565b6000835161b431818460208801619b6a565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161b46b816001840160208801619b6a565b0160010194935050505056fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea264697066735822122022532604dee90faafba8f6cebfe76ce3538c9312bd5371974218df483209449464736f6c634300081a00336080604052348015600f57600080fd5b506001600055610bd5806100246000396000f3fe6080604052600436106100635760003560e01c8063c513169111610040578063c5131691146100c1578063e04d4f97146100e1578063f05b6abf146100f457005b8063357fc5a21461006c578063660b9de01461008c5780636ed70169146100ac57005b3661006a57005b005b34801561007857600080fd5b5061006a61008736600461051e565b610114565b34801561009857600080fd5b5061006a6100a736600461055a565b6101aa565b3480156100b857600080fd5b5061006a6101e6565b3480156100cd57600080fd5b5061006a6100dc36600461051e565b61021b565b61006a6100ef3660046106b8565b6102f6565b34801561010057600080fd5b5061006a61010f366004610797565b61033a565b61011c61036f565b61013e73ffffffffffffffffffffffffffffffffffffffff83163383866103b2565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101a56001600055565b505050565b7f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e33826040516101db9291906108ca565b60405180910390a150565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61022361036f565b60006102306002856109b3565b90508060000361026c576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61028e73ffffffffffffffffffffffffffffffffffffffff84163384846103b2565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101a56001600055565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa333485858560405161032d959493929190610a52565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b1463384848460405161032d9493929190610adc565b6002600054036103ab576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261044790859061044d565b50505050565b600080602060008451602086016000885af180610470576040513d6000823e3d81fd5b50506000513d915081156104885780600114156104a2565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15610447576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240160405180910390fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461051957600080fd5b919050565b60008060006060848603121561053357600080fd5b83359250610543602085016104f5565b9150610551604085016104f5565b90509250925092565b60006020828403121561056c57600080fd5b813567ffffffffffffffff81111561058357600080fd5b82016060818503121561059557600080fd5b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156106125761061261059c565b604052919050565b600082601f83011261062b57600080fd5b813567ffffffffffffffff8111156106455761064561059c565b61067660207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016105cb565b81815284602083860101111561068b57600080fd5b816020850160208301376000918101602001919091529392505050565b8035801515811461051957600080fd5b6000806000606084860312156106cd57600080fd5b833567ffffffffffffffff8111156106e457600080fd5b6106f08682870161061a565b93505060208401359150610551604085016106a8565b600067ffffffffffffffff8211156107205761072061059c565b5060051b60200190565b600082601f83011261073b57600080fd5b813561074e61074982610706565b6105cb565b8082825260208201915060208360051b86010192508583111561077057600080fd5b602085015b8381101561078d578035835260209283019201610775565b5095945050505050565b6000806000606084860312156107ac57600080fd5b833567ffffffffffffffff8111156107c357600080fd5b8401601f810186136107d457600080fd5b80356107e261074982610706565b8082825260208201915060208360051b85010192508883111561080457600080fd5b602084015b8381101561084657803567ffffffffffffffff81111561082857600080fd5b6108378b60208389010161061a565b84525060209283019201610809565b509550505050602084013567ffffffffffffffff81111561086657600080fd5b6108728682870161072a565b925050610551604085016106a8565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610908836104f5565b1660408201526000602083013567ffffffffffffffff811680821461092c57600080fd5b6060840152506040830135368490037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261096857600080fd5b830160208101903567ffffffffffffffff81111561098557600080fd5b80360382131561099457600080fd5b606060808501526109a960a085018284610881565b9695505050505050565b6000826109e9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000815180845260005b81811015610a14576020818501810151868301820152016109f8565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610a8760a08301866109ee565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610ad2578151865260209586019590910190600101610ab4565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610b6f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610b5a8583516109ee565b94506020938401939190910190600101610b20565b505050508281036040840152610b858186610aa0565b915050610b96606083018415159052565b9594505050505056fea264697066735822122019741560ffed69bfe2dceeeb3d46120ee7f3cc17870cc7eb97a63cfb8ccf845564736f6c634300081a00335a657461436f6e6e6563746f724e617469766555706772616465546573742e736f6ca2646970667358221220d814620d02d24f526e4cafb249672e0b59f465ad8a490a1658267b2d9ff1067c64736f6c634300081a0033",
}

// ZetaConnectorNativeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNativeTestMetaData.ABI instead.
var ZetaConnectorNativeTestABI = ZetaConnectorNativeTestMetaData.ABI

// ZetaConnectorNativeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNativeTestMetaData.Bin instead.
var ZetaConnectorNativeTestBin = ZetaConnectorNativeTestMetaData.Bin

// DeployZetaConnectorNativeTest deploys a new Ethereum contract, binding an instance of ZetaConnectorNativeTest to it.
func DeployZetaConnectorNativeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZetaConnectorNativeTest, error) {
	parsed, err := ZetaConnectorNativeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNativeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNativeTest{ZetaConnectorNativeTestCaller: ZetaConnectorNativeTestCaller{contract: contract}, ZetaConnectorNativeTestTransactor: ZetaConnectorNativeTestTransactor{contract: contract}, ZetaConnectorNativeTestFilterer: ZetaConnectorNativeTestFilterer{contract: contract}}, nil
}

// ZetaConnectorNativeTest is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNativeTest struct {
	ZetaConnectorNativeTestCaller     // Read-only binding to the contract
	ZetaConnectorNativeTestTransactor // Write-only binding to the contract
	ZetaConnectorNativeTestFilterer   // Log filterer for contract events
}

// ZetaConnectorNativeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNativeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNativeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNativeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNativeTestSession struct {
	Contract     *ZetaConnectorNativeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ZetaConnectorNativeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNativeTestCallerSession struct {
	Contract *ZetaConnectorNativeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ZetaConnectorNativeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNativeTestTransactorSession struct {
	Contract     *ZetaConnectorNativeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ZetaConnectorNativeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNativeTestRaw struct {
	Contract *ZetaConnectorNativeTest // Generic contract binding to access the raw methods on
}

// ZetaConnectorNativeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNativeTestCallerRaw struct {
	Contract *ZetaConnectorNativeTestCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNativeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNativeTestTransactorRaw struct {
	Contract *ZetaConnectorNativeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNativeTest creates a new instance of ZetaConnectorNativeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeTest(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNativeTest, error) {
	contract, err := bindZetaConnectorNativeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTest{ZetaConnectorNativeTestCaller: ZetaConnectorNativeTestCaller{contract: contract}, ZetaConnectorNativeTestTransactor: ZetaConnectorNativeTestTransactor{contract: contract}, ZetaConnectorNativeTestFilterer: ZetaConnectorNativeTestFilterer{contract: contract}}, nil
}

// NewZetaConnectorNativeTestCaller creates a new read-only instance of ZetaConnectorNativeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeTestCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNativeTestCaller, error) {
	contract, err := bindZetaConnectorNativeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestCaller{contract: contract}, nil
}

// NewZetaConnectorNativeTestTransactor creates a new write-only instance of ZetaConnectorNativeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNativeTestTransactor, error) {
	contract, err := bindZetaConnectorNativeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestTransactor{contract: contract}, nil
}

// NewZetaConnectorNativeTestFilterer creates a new log filterer instance of ZetaConnectorNativeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNativeTestFilterer, error) {
	contract, err := bindZetaConnectorNativeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestFilterer{contract: contract}, nil
}

// bindZetaConnectorNativeTest binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNativeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNativeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNativeTest.Contract.ZetaConnectorNativeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.ZetaConnectorNativeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.ZetaConnectorNativeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNativeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.DEFAULTADMINROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.DEFAULTADMINROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) ISTEST() (bool, error) {
	return _ZetaConnectorNativeTest.Contract.ISTEST(&_ZetaConnectorNativeTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) ISTEST() (bool, error) {
	return _ZetaConnectorNativeTest.Contract.ISTEST(&_ZetaConnectorNativeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.PAUSERROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.PAUSERROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.TSSROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.TSSROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.WITHDRAWERROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.WITHDRAWERROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) ExcludeArtifacts() ([]string, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeArtifacts(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeArtifacts(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) ExcludeContracts() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeContracts(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeContracts(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeSelectors(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeSelectors(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) ExcludeSenders() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeSenders(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeSenders(&_ZetaConnectorNativeTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) Failed() (bool, error) {
	return _ZetaConnectorNativeTest.Contract.Failed(&_ZetaConnectorNativeTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) Failed() (bool, error) {
	return _ZetaConnectorNativeTest.Contract.Failed(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _ZetaConnectorNativeTest.Contract.TargetArtifactSelectors(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _ZetaConnectorNativeTest.Contract.TargetArtifactSelectors(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TargetArtifacts() ([]string, error) {
	return _ZetaConnectorNativeTest.Contract.TargetArtifacts(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TargetArtifacts() ([]string, error) {
	return _ZetaConnectorNativeTest.Contract.TargetArtifacts(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TargetContracts() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.TargetContracts(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.TargetContracts(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _ZetaConnectorNativeTest.Contract.TargetInterfaces(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _ZetaConnectorNativeTest.Contract.TargetInterfaces(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ZetaConnectorNativeTest.Contract.TargetSelectors(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ZetaConnectorNativeTest.Contract.TargetSelectors(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TargetSenders() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.TargetSenders(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.TargetSenders(&_ZetaConnectorNativeTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) SetUp() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.SetUp(&_ZetaConnectorNativeTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.SetUp(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestTSSUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testTSSUpgrade")
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestTSSUpgrade(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestTSSUpgrade(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testTSSUpgradeFailsIfSenderIsNotTSSUpdater")
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestTSSUpgradeFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testTSSUpgradeFailsIfZeroAddress")
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestUpgradeAndWithdraw is a paid mutator transaction binding the contract method 0xaf298bb1.
//
// Solidity: function testUpgradeAndWithdraw() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestUpgradeAndWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testUpgradeAndWithdraw")
}

// TestUpgradeAndWithdraw is a paid mutator transaction binding the contract method 0xaf298bb1.
//
// Solidity: function testUpgradeAndWithdraw() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestUpgradeAndWithdraw() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestUpgradeAndWithdraw(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestUpgradeAndWithdraw is a paid mutator transaction binding the contract method 0xaf298bb1.
//
// Solidity: function testUpgradeAndWithdraw() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestUpgradeAndWithdraw() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestUpgradeAndWithdraw(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdraw is a paid mutator transaction binding the contract method 0xd509b16c.
//
// Solidity: function testWithdraw() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdraw")
}

// TestWithdraw is a paid mutator transaction binding the contract method 0xd509b16c.
//
// Solidity: function testWithdraw() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdraw() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdraw(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdraw is a paid mutator transaction binding the contract method 0xd509b16c.
//
// Solidity: function testWithdraw() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdraw() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdraw(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20 is a paid mutator transaction binding the contract method 0x3cba0107.
//
// Solidity: function testWithdrawAndCallReceiveERC20() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndCallReceiveERC20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveERC20")
}

// TestWithdrawAndCallReceiveERC20 is a paid mutator transaction binding the contract method 0x3cba0107.
//
// Solidity: function testWithdrawAndCallReceiveERC20() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndCallReceiveERC20() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveERC20(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20 is a paid mutator transaction binding the contract method 0x3cba0107.
//
// Solidity: function testWithdrawAndCallReceiveERC20() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndCallReceiveERC20() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveERC20(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xc1909972.
//
// Solidity: function testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer")
}

// TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xc1909972.
//
// Solidity: function testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xc1909972.
//
// Solidity: function testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20Partial is a paid mutator transaction binding the contract method 0xdcf7d037.
//
// Solidity: function testWithdrawAndCallReceiveERC20Partial() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndCallReceiveERC20Partial(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveERC20Partial")
}

// TestWithdrawAndCallReceiveERC20Partial is a paid mutator transaction binding the contract method 0xdcf7d037.
//
// Solidity: function testWithdrawAndCallReceiveERC20Partial() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndCallReceiveERC20Partial() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveERC20Partial(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20Partial is a paid mutator transaction binding the contract method 0xdcf7d037.
//
// Solidity: function testWithdrawAndCallReceiveERC20Partial() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndCallReceiveERC20Partial() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveERC20Partial(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveNoParams is a paid mutator transaction binding the contract method 0x49346558.
//
// Solidity: function testWithdrawAndCallReceiveNoParams() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndCallReceiveNoParams(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveNoParams")
}

// TestWithdrawAndCallReceiveNoParams is a paid mutator transaction binding the contract method 0x49346558.
//
// Solidity: function testWithdrawAndCallReceiveNoParams() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndCallReceiveNoParams() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveNoParams(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveNoParams is a paid mutator transaction binding the contract method 0x49346558.
//
// Solidity: function testWithdrawAndCallReceiveNoParams() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndCallReceiveNoParams() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveNoParams(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndRevert is a paid mutator transaction binding the contract method 0xde1cb76c.
//
// Solidity: function testWithdrawAndRevert() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndRevert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndRevert")
}

// TestWithdrawAndRevert is a paid mutator transaction binding the contract method 0xde1cb76c.
//
// Solidity: function testWithdrawAndRevert() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndRevert() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndRevert(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndRevert is a paid mutator transaction binding the contract method 0xde1cb76c.
//
// Solidity: function testWithdrawAndRevert() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndRevert() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndRevert(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x82832014.
//
// Solidity: function testWithdrawAndRevertFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndRevertFailsIfSenderIsNotWithdrawer")
}

// TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x82832014.
//
// Solidity: function testWithdrawAndRevertFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x82832014.
//
// Solidity: function testWithdrawAndRevertFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xfe574f84.
//
// Solidity: function testWithdrawFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawFailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawFailsIfSenderIsNotWithdrawer")
}

// TestWithdrawFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xfe574f84.
//
// Solidity: function testWithdrawFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawFailsIfSenderIsNotWithdrawer(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xfe574f84.
//
// Solidity: function testWithdrawFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawFailsIfSenderIsNotWithdrawer(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawTogglePause is a paid mutator transaction binding the contract method 0xccb0e3f2.
//
// Solidity: function testWithdrawTogglePause() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawTogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawTogglePause")
}

// TestWithdrawTogglePause is a paid mutator transaction binding the contract method 0xccb0e3f2.
//
// Solidity: function testWithdrawTogglePause() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawTogglePause() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawTogglePause(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawTogglePause is a paid mutator transaction binding the contract method 0xccb0e3f2.
//
// Solidity: function testWithdrawTogglePause() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawTogglePause() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawTogglePause(&_ZetaConnectorNativeTest.TransactOpts)
}

// ZetaConnectorNativeTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestCalledIterator struct {
	Event *ZetaConnectorNativeTestCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestCalled represents a Called event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ZetaConnectorNativeTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestCalledIterator{contract: _ZetaConnectorNativeTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestCalled)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Called", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseCalled(log types.Log) (*ZetaConnectorNativeTestCalled, error) {
	event := new(ZetaConnectorNativeTestCalled)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestDepositedIterator struct {
	Event *ZetaConnectorNativeTestDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestDeposited represents a Deposited event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestDeposited struct {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterDeposited(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ZetaConnectorNativeTestDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestDepositedIterator{contract: _ZetaConnectorNativeTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestDeposited, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestDeposited)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseDeposited(log types.Log) (*ZetaConnectorNativeTestDeposited, error) {
	event := new(ZetaConnectorNativeTestDeposited)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestExecutedIterator struct {
	Event *ZetaConnectorNativeTestExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestExecuted represents a Executed event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*ZetaConnectorNativeTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestExecutedIterator{contract: _ZetaConnectorNativeTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestExecuted)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseExecuted(log types.Log) (*ZetaConnectorNativeTestExecuted, error) {
	event := new(ZetaConnectorNativeTestExecuted)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestExecutedWithERC20Iterator struct {
	Event *ZetaConnectorNativeTestExecutedWithERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestExecutedWithERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestExecutedWithERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*ZetaConnectorNativeTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestExecutedWithERC20Iterator{contract: _ZetaConnectorNativeTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestExecutedWithERC20)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseExecutedWithERC20(log types.Log) (*ZetaConnectorNativeTestExecutedWithERC20, error) {
	event := new(ZetaConnectorNativeTestExecutedWithERC20)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedERC20Iterator struct {
	Event *ZetaConnectorNativeTestReceivedERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReceivedERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestReceivedERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReceivedERC20 represents a ReceivedERC20 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*ZetaConnectorNativeTestReceivedERC20Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestReceivedERC20Iterator{contract: _ZetaConnectorNativeTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReceivedERC20)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReceivedERC20(log types.Log) (*ZetaConnectorNativeTestReceivedERC20, error) {
	event := new(ZetaConnectorNativeTestReceivedERC20)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedNoParamsIterator struct {
	Event *ZetaConnectorNativeTestReceivedNoParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReceivedNoParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestReceivedNoParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReceivedNoParams represents a ReceivedNoParams event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*ZetaConnectorNativeTestReceivedNoParamsIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestReceivedNoParamsIterator{contract: _ZetaConnectorNativeTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReceivedNoParams)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReceivedNoParams(log types.Log) (*ZetaConnectorNativeTestReceivedNoParams, error) {
	event := new(ZetaConnectorNativeTestReceivedNoParams)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedNonPayableIterator struct {
	Event *ZetaConnectorNativeTestReceivedNonPayable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReceivedNonPayable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestReceivedNonPayable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReceivedNonPayable represents a ReceivedNonPayable event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*ZetaConnectorNativeTestReceivedNonPayableIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestReceivedNonPayableIterator{contract: _ZetaConnectorNativeTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReceivedNonPayable)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReceivedNonPayable(log types.Log) (*ZetaConnectorNativeTestReceivedNonPayable, error) {
	event := new(ZetaConnectorNativeTestReceivedNonPayable)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedPayableIterator struct {
	Event *ZetaConnectorNativeTestReceivedPayable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReceivedPayable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestReceivedPayable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReceivedPayable represents a ReceivedPayable event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedPayable struct {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*ZetaConnectorNativeTestReceivedPayableIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestReceivedPayableIterator{contract: _ZetaConnectorNativeTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReceivedPayable)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReceivedPayable(log types.Log) (*ZetaConnectorNativeTestReceivedPayable, error) {
	event := new(ZetaConnectorNativeTestReceivedPayable)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedRevertIterator struct {
	Event *ZetaConnectorNativeTestReceivedRevert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReceivedRevert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestReceivedRevert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReceivedRevert represents a ReceivedRevert event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*ZetaConnectorNativeTestReceivedRevertIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestReceivedRevertIterator{contract: _ZetaConnectorNativeTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReceivedRevert)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReceivedRevert(log types.Log) (*ZetaConnectorNativeTestReceivedRevert, error) {
	event := new(ZetaConnectorNativeTestReceivedRevert)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestRevertedIterator struct {
	Event *ZetaConnectorNativeTestReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReverted represents a Reverted event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReverted struct {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ZetaConnectorNativeTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestRevertedIterator{contract: _ZetaConnectorNativeTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0x1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b1436.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReverted)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Reverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReverted(log types.Log) (*ZetaConnectorNativeTestReverted, error) {
	event := new(ZetaConnectorNativeTestReverted)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator struct {
	Event *ZetaConnectorNativeTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestUpdatedGatewayTSSAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestUpdatedGatewayTSSAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestUpdatedGatewayTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator{contract: _ZetaConnectorNativeTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestUpdatedGatewayTSSAddress)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*ZetaConnectorNativeTestUpdatedGatewayTSSAddress, error) {
	event := new(ZetaConnectorNativeTestUpdatedGatewayTSSAddress)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator is returned from FilterUpdatedZetaConnectorTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedZetaConnectorTSSAddress events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator struct {
	Event *ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress represents a UpdatedZetaConnectorTSSAddress event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedZetaConnectorTSSAddress is a free log retrieval operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterUpdatedZetaConnectorTSSAddress(opts *bind.FilterOpts) (*ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator{contract: _ZetaConnectorNativeTest.contract, event: "UpdatedZetaConnectorTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedZetaConnectorTSSAddress is a free log subscription operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchUpdatedZetaConnectorTSSAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedZetaConnectorTSSAddress is a log parse operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseUpdatedZetaConnectorTSSAddress(log types.Log) (*ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress, error) {
	event := new(ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnIterator struct {
	Event *ZetaConnectorNativeTestWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestWithdrawn represents a Withdrawn event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeTestWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestWithdrawnIterator{contract: _ZetaConnectorNativeTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestWithdrawn)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseWithdrawn(log types.Log) (*ZetaConnectorNativeTestWithdrawn, error) {
	event := new(ZetaConnectorNativeTestWithdrawn)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnAndCalledIterator struct {
	Event *ZetaConnectorNativeTestWithdrawnAndCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestWithdrawnAndCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestWithdrawnAndCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnAndCalled struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeTestWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestWithdrawnAndCalledIterator{contract: _ZetaConnectorNativeTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestWithdrawnAndCalled, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestWithdrawnAndCalled)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*ZetaConnectorNativeTestWithdrawnAndCalled, error) {
	event := new(ZetaConnectorNativeTestWithdrawnAndCalled)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnAndRevertedIterator struct {
	Event *ZetaConnectorNativeTestWithdrawnAndReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestWithdrawnAndReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestWithdrawnAndReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnAndReverted struct {
	To            common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeTestWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestWithdrawnAndRevertedIterator{contract: _ZetaConnectorNativeTest.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestWithdrawnAndReverted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestWithdrawnAndReverted)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*ZetaConnectorNativeTestWithdrawnAndReverted, error) {
	event := new(ZetaConnectorNativeTestWithdrawnAndReverted)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnV2Iterator struct {
	Event *ZetaConnectorNativeTestWithdrawnV2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestWithdrawnV2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestWithdrawnV2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestWithdrawnV2 represents a WithdrawnV2 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnV2 struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnV2 is a free log retrieval operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeTestWithdrawnV2Iterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "WithdrawnV2", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestWithdrawnV2Iterator{contract: _ZetaConnectorNativeTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestWithdrawnV2, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "WithdrawnV2", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestWithdrawnV2)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnV2 is a log parse operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseWithdrawnV2(log types.Log) (*ZetaConnectorNativeTestWithdrawnV2, error) {
	event := new(ZetaConnectorNativeTestWithdrawnV2)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogIterator struct {
	Event *ZetaConnectorNativeTestLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLog represents a Log event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLog(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogIterator{contract: _ZetaConnectorNativeTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLog) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLog)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLog(log types.Log) (*ZetaConnectorNativeTestLog, error) {
	event := new(ZetaConnectorNativeTestLog)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogAddressIterator struct {
	Event *ZetaConnectorNativeTestLogAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogAddress represents a LogAddress event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogAddressIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogAddress)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogAddress(log types.Log) (*ZetaConnectorNativeTestLogAddress, error) {
	event := new(ZetaConnectorNativeTestLogAddress)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogArrayIterator struct {
	Event *ZetaConnectorNativeTestLogArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogArray represents a LogArray event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogArrayIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogArrayIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogArray) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogArray)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogArray(log types.Log) (*ZetaConnectorNativeTestLogArray, error) {
	event := new(ZetaConnectorNativeTestLogArray)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogArray0Iterator struct {
	Event *ZetaConnectorNativeTestLogArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogArray0 represents a LogArray0 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogArray0Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogArray0Iterator{contract: _ZetaConnectorNativeTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogArray0)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogArray0(log types.Log) (*ZetaConnectorNativeTestLogArray0, error) {
	event := new(ZetaConnectorNativeTestLogArray0)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogArray1Iterator struct {
	Event *ZetaConnectorNativeTestLogArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogArray1 represents a LogArray1 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogArray1Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogArray1Iterator{contract: _ZetaConnectorNativeTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogArray1)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogArray1(log types.Log) (*ZetaConnectorNativeTestLogArray1, error) {
	event := new(ZetaConnectorNativeTestLogArray1)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogBytesIterator struct {
	Event *ZetaConnectorNativeTestLogBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogBytes represents a LogBytes event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogBytesIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogBytesIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogBytes)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogBytes(log types.Log) (*ZetaConnectorNativeTestLogBytes, error) {
	event := new(ZetaConnectorNativeTestLogBytes)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogBytes32Iterator struct {
	Event *ZetaConnectorNativeTestLogBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogBytes32 represents a LogBytes32 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogBytes32Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogBytes32Iterator{contract: _ZetaConnectorNativeTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogBytes32)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogBytes32(log types.Log) (*ZetaConnectorNativeTestLogBytes32, error) {
	event := new(ZetaConnectorNativeTestLogBytes32)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogIntIterator struct {
	Event *ZetaConnectorNativeTestLogInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogInt represents a LogInt event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogIntIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogIntIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogInt) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogInt)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogInt(log types.Log) (*ZetaConnectorNativeTestLogInt, error) {
	event := new(ZetaConnectorNativeTestLogInt)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedAddressIterator struct {
	Event *ZetaConnectorNativeTestLogNamedAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogNamedAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedAddress represents a LogNamedAddress event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedAddressIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedAddress)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedAddress(log types.Log) (*ZetaConnectorNativeTestLogNamedAddress, error) {
	event := new(ZetaConnectorNativeTestLogNamedAddress)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedArrayIterator struct {
	Event *ZetaConnectorNativeTestLogNamedArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogNamedArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedArray represents a LogNamedArray event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedArrayIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedArrayIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedArray)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedArray(log types.Log) (*ZetaConnectorNativeTestLogNamedArray, error) {
	event := new(ZetaConnectorNativeTestLogNamedArray)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedArray0Iterator struct {
	Event *ZetaConnectorNativeTestLogNamedArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogNamedArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedArray0 represents a LogNamedArray0 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedArray0Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedArray0Iterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedArray0)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedArray0(log types.Log) (*ZetaConnectorNativeTestLogNamedArray0, error) {
	event := new(ZetaConnectorNativeTestLogNamedArray0)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedArray1Iterator struct {
	Event *ZetaConnectorNativeTestLogNamedArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogNamedArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedArray1 represents a LogNamedArray1 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedArray1Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedArray1Iterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedArray1)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedArray1(log types.Log) (*ZetaConnectorNativeTestLogNamedArray1, error) {
	event := new(ZetaConnectorNativeTestLogNamedArray1)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedBytesIterator struct {
	Event *ZetaConnectorNativeTestLogNamedBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogNamedBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedBytes represents a LogNamedBytes event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedBytesIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedBytesIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedBytes)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedBytes(log types.Log) (*ZetaConnectorNativeTestLogNamedBytes, error) {
	event := new(ZetaConnectorNativeTestLogNamedBytes)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedBytes32Iterator struct {
	Event *ZetaConnectorNativeTestLogNamedBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogNamedBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedBytes32Iterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedBytes32)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedBytes32(log types.Log) (*ZetaConnectorNativeTestLogNamedBytes32, error) {
	event := new(ZetaConnectorNativeTestLogNamedBytes32)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedDecimalIntIterator struct {
	Event *ZetaConnectorNativeTestLogNamedDecimalInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedDecimalInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogNamedDecimalInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedDecimalIntIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedDecimalInt)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*ZetaConnectorNativeTestLogNamedDecimalInt, error) {
	event := new(ZetaConnectorNativeTestLogNamedDecimalInt)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedDecimalUintIterator struct {
	Event *ZetaConnectorNativeTestLogNamedDecimalUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedDecimalUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogNamedDecimalUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedDecimalUintIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedDecimalUint)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*ZetaConnectorNativeTestLogNamedDecimalUint, error) {
	event := new(ZetaConnectorNativeTestLogNamedDecimalUint)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedIntIterator struct {
	Event *ZetaConnectorNativeTestLogNamedInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogNamedInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedInt represents a LogNamedInt event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedIntIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedIntIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedInt)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedInt(log types.Log) (*ZetaConnectorNativeTestLogNamedInt, error) {
	event := new(ZetaConnectorNativeTestLogNamedInt)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedStringIterator struct {
	Event *ZetaConnectorNativeTestLogNamedString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogNamedString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedString represents a LogNamedString event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedStringIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedStringIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedString)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedString(log types.Log) (*ZetaConnectorNativeTestLogNamedString, error) {
	event := new(ZetaConnectorNativeTestLogNamedString)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedUintIterator struct {
	Event *ZetaConnectorNativeTestLogNamedUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogNamedUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedUint represents a LogNamedUint event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedUintIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedUintIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedUint)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedUint(log types.Log) (*ZetaConnectorNativeTestLogNamedUint, error) {
	event := new(ZetaConnectorNativeTestLogNamedUint)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogStringIterator struct {
	Event *ZetaConnectorNativeTestLogString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogString represents a LogString event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogString(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogStringIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogStringIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogString) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogString)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogString(log types.Log) (*ZetaConnectorNativeTestLogString, error) {
	event := new(ZetaConnectorNativeTestLogString)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogUintIterator struct {
	Event *ZetaConnectorNativeTestLogUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogUint represents a LogUint event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogUintIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogUintIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogUint) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogUint)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogUint(log types.Log) (*ZetaConnectorNativeTestLogUint, error) {
	event := new(ZetaConnectorNativeTestLogUint)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogsIterator struct {
	Event *ZetaConnectorNativeTestLogs // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogs)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestLogs)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogs represents a Logs event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogs(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogsIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogsIterator{contract: _ZetaConnectorNativeTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogs) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogs)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "logs", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogs(log types.Log) (*ZetaConnectorNativeTestLogs, error) {
	event := new(ZetaConnectorNativeTestLogs)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
