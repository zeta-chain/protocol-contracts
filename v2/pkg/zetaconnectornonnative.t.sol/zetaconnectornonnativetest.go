// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornonnative

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

// ZetaConnectorNonNativeTestMetaData contains all meta data concerning the ZetaConnectorNonNativeTest contract.
var ZetaConnectorNonNativeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testSexMaxSupplyFailsIfSenderIsNotTss\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallFailsIfMaxSupplyIsReached\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20Partial\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveNoParams\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndFailsIfMaxSupplyIsReached\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevert\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertFailsIfMaxSupplyIsReached\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedZetaConnectorTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxSupply\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b506201035d806200003e6000396000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c8063a783c78911610104578063d509b16c116100a2578063e63ab1e911610071578063e63ab1e914610344578063fa7626d41461036b578063fdca905214610378578063fe574f841461038057600080fd5b8063d509b16c14610324578063dcf7d0371461032c578063de1cb76c14610334578063e20c9f711461033c57600080fd5b8063b5508aa9116100de578063b5508aa9146102f4578063ba414fa6146102fc578063c190997214610314578063ccb0e3f21461031c57600080fd5b8063a783c789146102bd578063aaf74192146102e4578063b0464fdc146102ec57600080fd5b80634934655811610171578063828320141161014b578063828320141461025657806385226c811461025e57806385f438c114610273578063916a17c6146102a857600080fd5b8063493465581461023157806366d9a9a0146102395780637db20efb1461024e57600080fd5b80632ade3880116101ad5780632ade3880146102045780633cba0107146102195780633e5e3c23146102215780633f7286f41461022957600080fd5b80630a9254e4146101d45780631ed7831c146101de5780632506ef03146101fc575b600080fd5b6101dc610388565b005b6101e6610b4f565b6040516101f39190618ff2565b60405180910390f35b6101dc610bb1565b61020c610e65565b6040516101f3919061908e565b6101dc610fa7565b6101e6611769565b6101e66117c9565b6101dc611829565b610241611dfd565b6040516101f391906191f4565b6101dc611f7f565b6101dc61221e565b61026661247a565b6040516101f39190619292565b61029a7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6040519081526020016101f3565b6102b061254a565b6040516101f39190619309565b61029a7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101dc612645565b6102b06128c0565b6102666129bb565b610304612a8b565b60405190151581526020016101f3565b6101dc612b5f565b6101dc612dca565b6101dc6138f6565b6101dc613c4b565b6101dc614279565b6101e661499f565b61029a7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b601f546103049060ff1681565b6101dc6149ff565b6101dc614c03565b60258054307fffffffffffffffffffffffff000000000000000000000000000000000000000091821617909155602680546112349083161790556027805461567892168217905560405181906103dd90618f01565b6001600160a01b03928316815291166020820152604001604051809103906000f080158015610410573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081178255604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260275460255492519086169481019490945260448401929092529092166064820152610501919060840160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b00000000000000000000000000000000000000000000000000000000179052614df0565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116811790915560275460255460405192939182169291169061058d90618f0f565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f0801580156105c9573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316179055602054602454602754602554604051938516949283169391831692169061062490618f1d565b6001600160a01b039485168152928416602084015290831660408301529091166060820152608001604051809103906000f080158015610668573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392831617905560275460405163ca669fa760e01b815291166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156106ed57600080fd5b505af1158015610701573d6000803e3d6000fd5b5050602480546027546023546040517f15d57fd40000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152908216938101939093521692506315d57fd49150604401600060405180830381600087803b15801561077457600080fd5b505af1158015610788573d6000803e3d6000fd5b5050505060405161079890618f2b565b604051809103906000f0801580156107b4573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556027546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b15801561086057600080fd5b505af1158015610874573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b1580156108ea57600080fd5b505af11580156108fe573d6000803e3d6000fd5b50506020546022546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f9150602401600060405180830381600087803b15801561096457600080fd5b505af1158015610978573d6000803e3d6000fd5b50506020546023546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef9150602401600060405180830381600087803b1580156109de57600080fd5b505af11580156109f2573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610a5457600080fd5b505af1158015610a68573d6000803e3d6000fd5b5050604080516080810182526025546001600160a01b039081168252602454811660208084019182526001848601908152855191820190955260008152606084018190528351602880549185167fffffffffffffffffffffffff0000000000000000000000000000000000000000909216919091178155915160298054965167ffffffffffffffff1674010000000000000000000000000000000000000000027fffffffff000000000000000000000000000000000000000000000000000000009097169190941617949094179091559093509150602a90610b4a9082619463565b505050565b60606016805480602002602001604051908101604052809291908181526020018280548015610ba757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610b89575b5050505050905090565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a09060009060250160408051808303601f190181529082905260275463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015610c4c57600080fd5b505af1158015610c60573d6000803e3d6000fd5b50506023546040517f6f8b44b0000000000000000000000000000000000000000000000000000000008152600481018690526001600160a01b039091169250636f8b44b09150602401600060405180830381600087803b158015610cc357600080fd5b505af1158015610cd7573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015610d3457600080fd5b505af1158015610d48573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fc30436e9000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015610dd157600080fd5b505af1158015610de5573d6000803e3d6000fd5b50506023546021546001600160a01b039182169350632ccb404d925016610e0d856001619551565b8460286040518563ffffffff1660e01b8152600401610e2f949392919061964c565b600060405180830381600087803b158015610e4957600080fd5b505af1158015610e5d573d6000803e3d6000fd5b505050505050565b6060601e805480602002602001604051908101604052809291908181526020016000905b82821015610f9e57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015610f87578382906000526020600020018054610efa906193cf565b80601f0160208091040260200160405190810160405280929190818152602001828054610f26906193cf565b8015610f735780601f10610f4857610100808354040283529160200191610f73565b820191906000526020600020905b815481529060010190602001808311610f5657829003601f168201915b505050505081526020019060010190610edb565b505050508152505081526020019060010190610e89565b50505050905090565b60248054602654604051620186a09381018490526001600160a01b03928316604482015291166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a2000000000000000000000000000000000000000000000000000000001790526024805460265492516370a0823160e01b81526001600160a01b0393841660048201529394506000939216916370a082319101602060405180830381865afa158015611082573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110a6919061968d565b90506110b3816000614e0f565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015611103573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611127919061968d565b9050611134816000614e0f565b6020546040516001600160a01b039091166024820152604481018690526064810185905260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1e458bee0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161121b916001600160a01b03919091169060009086906004016196a6565b600060405180830381600087803b15801561123557600080fd5b505af1158015611249573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156112db57600080fd5b505af11580156112ef573d6000803e3d6000fd5b505060208054602454602654604080516001600160a01b0394851681529485018d905291831684830152919091166060830152517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609350908190036080019150a16023546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156113de57600080fd5b505af11580156113f2573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d915061143790899088906196ce565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561149857600080fd5b505af11580156114ac573d6000803e3d6000fd5b50506023546021546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef935061150492909116908a9089908b906004016196e7565b600060405180830381600087803b15801561151e57600080fd5b505af1158015611532573d6000803e3d6000fd5b5050602480546026546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa158015611584573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115a8919061968d565b90506115b48188614e0f565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015611604573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611628919061968d565b9050611635816000614e0f565b602480546020546021546040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529082169381019390935260009291169063dd62ed3e90604401602060405180830381865afa1580156116ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116cf919061968d565b90506116dc816000614e0f565b602480546020546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa15801561172c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611750919061968d565b905061175d816000614e0f565b50505050505050505050565b60606018805480602002602001604051908101604052809291908181526020018280548015610ba7576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610b89575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015610ba7576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610b89575050505050905090565b604080516004808252602480830184526020830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed70169000000000000000000000000000000000000000000000000000000001790525460265493516370a0823160e01b8152620186a0946000949385936001600160a01b03908116936370a08231936118ca93921691016001600160a01b0391909116815260200190565b602060405180830381865afa1580156118e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061190b919061968d565b9050611918816000614e0f565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015611968573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061198c919061968d565b9050611999816000614e0f565b6020546040516001600160a01b039091166024820152604481018690526064810185905260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1e458bee0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391611a80916001600160a01b03919091169060009086906004016196a6565b600060405180830381600087803b158015611a9a57600080fd5b505af1158015611aae573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015611b4057600080fd5b505af1158015611b54573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0935001905060405180910390a16023546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611c2657600080fd5b505af1158015611c3a573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d9150611c7f90899088906196ce565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611ce057600080fd5b505af1158015611cf4573d6000803e3d6000fd5b50506023546021546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef9350611d4c92909116908a9089908b906004016196e7565b600060405180830381600087803b158015611d6657600080fd5b505af1158015611d7a573d6000803e3d6000fd5b5050602480546026546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa158015611dcc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611df0919061968d565b90506115b4816000614e0f565b6060601b805480602002602001604051908101604052809291908181526020016000905b82821015610f9e5783829060005260206000209060020201604051806040016040529081600082018054611e54906193cf565b80601f0160208091040260200160405190810160405280929190818152602001828054611e80906193cf565b8015611ecd5780601f10611ea257610100808354040283529160200191611ecd565b820191906000526020600020905b815481529060010190602001808311611eb057829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015611f6757602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411611f145790505b50505050508152505081526020019060010190611e21565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611fdd57600080fd5b505af1158015611ff1573d6000803e3d6000fd5b50506023546040517f6f8b44b0000000000000000000000000000000000000000000000000000000008152600481018590526001600160a01b039091169250636f8b44b09150602401600060405180830381600087803b15801561205457600080fd5b505af1158015612068573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156120c557600080fd5b505af11580156120d9573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fc30436e9000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561216257600080fd5b505af1158015612176573d6000803e3d6000fd5b50506023546021546001600160a01b03918216935063106e629092501661219e846001619551565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b039092166004830152602482015260006044820152606401600060405180830381600087803b15801561220357600080fd5b505af1158015612217573d6000803e3d6000fd5b5050505050565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a090600090819060250160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156122bb57600080fd5b505af11580156122cf573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506123ba9190600401619720565b600060405180830381600087803b1580156123d457600080fd5b505af11580156123e8573d6000803e3d6000fd5b50506023546021546040517f2ccb404d0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450632ccb404d93506124439290911690879086908890602890600401619733565b600060405180830381600087803b15801561245d57600080fd5b505af1158015612471573d6000803e3d6000fd5b50505050505050565b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015610f9e5783829060005260206000200180546124bd906193cf565b80601f01602080910402602001604051908101604052809291908181526020018280546124e9906193cf565b80156125365780601f1061250b57610100808354040283529160200191612536565b820191906000526020600020905b81548152906001019060200180831161251957829003601f168201915b50505050508152602001906001019061249e565b6060601d805480602002602001604051908101604052809291908181526020016000905b82821015610f9e5760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561262d57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116125da5790505b5050505050815250508152602001906001019061256e565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a09060009060250160408051808303601f190181529082905260275463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156126e057600080fd5b505af11580156126f4573d6000803e3d6000fd5b50506023546040517f6f8b44b0000000000000000000000000000000000000000000000000000000008152600481018690526001600160a01b039091169250636f8b44b09150602401600060405180830381600087803b15801561275757600080fd5b505af115801561276b573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156127c857600080fd5b505af11580156127dc573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fc30436e9000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561286557600080fd5b505af1158015612879573d6000803e3d6000fd5b50506023546021546001600160a01b039182169350635e3e9fef9250166128a1856001619551565b846040518463ffffffff1660e01b8152600401610e2f9392919061977f565b6060601c805480602002602001604051908101604052809291908181526020016000905b82821015610f9e5760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156129a357602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116129505790505b505050505081525050815260200190600101906128e4565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015610f9e5783829060005260206000200180546129fe906193cf565b80601f0160208091040260200160405190810160405280929190818152602001828054612a2a906193cf565b8015612a775780601f10612a4c57610100808354040283529160200191612a77565b820191906000526020600020905b815481529060010190602001808311612a5a57829003601f168201915b5050505050815260200190600101906129df565b60085460009060ff1615612aa3575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015612b34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b58919061968d565b1415905090565b60248054602654604051620186a09381018490526001600160a01b03928316604482015291166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602554905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612c4557600080fd5b505af1158015612c59573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250612d449190600401619720565b600060405180830381600087803b158015612d5e57600080fd5b505af1158015612d72573d6000803e3d6000fd5b50506023546021546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef935061244392909116908790869088906004016196e7565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090600090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612e2b57600080fd5b505af1158015612e3f573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250612f2a9190600401619720565b600060405180830381600087803b158015612f4457600080fd5b505af1158015612f58573d6000803e3d6000fd5b50505050602360009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b158015612fac57600080fd5b505af1158015612fc0573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561301d57600080fd5b505af1158015613031573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061311c9190600401619720565b600060405180830381600087803b15801561313657600080fd5b505af115801561314a573d6000803e3d6000fd5b50505050602360009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561319e57600080fd5b505af11580156131b2573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561320f57600080fd5b505af1158015613223573d6000803e3d6000fd5b50505050602360009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561327757600080fd5b505af115801561328b573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd93c0665000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561331457600080fd5b505af1158015613328573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561338557600080fd5b505af1158015613399573d6000803e3d6000fd5b50506023546026546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101879052604481018690529116925063106e62909150606401600060405180830381600087803b15801561340d57600080fd5b505af1158015613421573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561347e57600080fd5b505af1158015613492573d6000803e3d6000fd5b50505050602360009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156134e657600080fd5b505af11580156134fa573d6000803e3d6000fd5b5050602480546026546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa15801561354c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613570919061968d565b905061357d816000614e0f565b6026546040516001600160a01b039091166024820152604481018490526064810183905260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1e458bee0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391613664916001600160a01b03919091169060009086906004016196a6565b600060405180830381600087803b15801561367e57600080fd5b505af1158015613692573d6000803e3d6000fd5b50506023546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561372457600080fd5b505af1158015613738573d6000803e3d6000fd5b50506026546040518781526001600160a01b0390911692507f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5915060200160405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156137d757600080fd5b505af11580156137eb573d6000803e3d6000fd5b50506023546026546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101899052604481018890529116925063106e629091506064015b600060405180830381600087803b15801561386057600080fd5b505af1158015613874573d6000803e3d6000fd5b5050602480546026546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa1580156138c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138ea919061968d565b90506122178186614e0f565b602480546026546040516370a0823160e01b81526001600160a01b039182166004820152620186a09360009392909216916370a082319101602060405180830381865afa15801561394b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061396f919061968d565b905061397c816000614e0f565b6026546040516001600160a01b0390911660248201526044810183905260006064820181905290819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1e458bee0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391613a65916001600160a01b03919091169060009086906004016196a6565b600060405180830381600087803b158015613a7f57600080fd5b505af1158015613a93573d6000803e3d6000fd5b50506023546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015613b2557600080fd5b505af1158015613b39573d6000803e3d6000fd5b50506026546040518781526001600160a01b0390911692507f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5915060200160405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613bd857600080fd5b505af1158015613bec573d6000803e3d6000fd5b50506023546026546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101899052604481018790529116925063106e62909150606401613846565b60248054602654604051620186a09381018490526001600160a01b03928316604482015291166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc5131691000000000000000000000000000000000000000000000000000000001790526024805460265492516370a0823160e01b81526001600160a01b0393841660048201529394506000939216916370a082319101602060405180830381865afa158015613d26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d4a919061968d565b9050613d57816000614e0f565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015613da7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613dcb919061968d565b9050613dd8816000614e0f565b6020546040516001600160a01b039091166024820152604481018690526064810185905260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1e458bee0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391613ebf916001600160a01b03919091169060009086906004016196a6565b600060405180830381600087803b158015613ed957600080fd5b505af1158015613eed573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015613f7f57600080fd5b505af1158015613f93573d6000803e3d6000fd5b50506020547f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af6092506001600160a01b03169050613fd16002896197b8565b602454602654604080516001600160a01b03958616815260208101949094529184168383015292909216606082015290519081900360800190a16023546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561409957600080fd5b505af11580156140ad573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d91506140f290899088906196ce565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561415357600080fd5b505af1158015614167573d6000803e3d6000fd5b50506023546021546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef93506141bf92909116908a9089908b906004016196e7565b600060405180830381600087803b1580156141d957600080fd5b505af11580156141ed573d6000803e3d6000fd5b5050602480546026546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa15801561423f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614263919061968d565b90506115b48161427460028a6197b8565b614e0f565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a090600090819060250160408051808303601f1901815290829052602480546021546370a0823160e01b85526001600160a01b0390811660048601529294506000939216916370a082319101602060405180830381865afa15801561430e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614332919061968d565b905061433f816000614e0f565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa15801561438f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906143b3919061968d565b6020546040516001600160a01b039091166024820152604481018790526064810186905290915060009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1e458bee0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161449d916001600160a01b03919091169060009086906004016196a6565b600060405180830381600087803b1580156144b757600080fd5b505af11580156144cb573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561455d57600080fd5b505af1158015614571573d6000803e3d6000fd5b50506020546040517f8f39d9623d22cb24244fa4bf1219c32f0a8c7522a0978f7778ec6e376507e26493506145b592506001600160a01b03909116906028906197f3565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561464b57600080fd5b505af115801561465f573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507ff196c7207ed90b5ece8dbf4d77864a91b1ccca28fe83e005b4969d7584b2b19b906146ad908a908990602890619815565b60405180910390a36023546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561474357600080fd5b505af1158015614757573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507fa99787224bed6d592c223365ff0f91d203cd2eea00f6b1be8b0263b835378017915061479f9089908890602890619815565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561480057600080fd5b505af1158015614814573d6000803e3d6000fd5b50506023546021546040517f2ccb404d0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450632ccb404d935061486f92909116908a9089908b90602890600401619733565b600060405180830381600087803b15801561488957600080fd5b505af115801561489d573d6000803e3d6000fd5b5050602480546021546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa1580156148ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614913919061968d565b905061491f8188614e0f565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa15801561496f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614993919061968d565b90506116358185614e0f565b60606015805480602002602001604051908101604052809291908181526020018280548015610ba7576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610b89575050505050905090565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614a5857600080fd5b505af1158015614a6c573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614b579190600401619720565b600060405180830381600087803b158015614b7157600080fd5b505af1158015614b85573d6000803e3d6000fd5b50506023546040517f6f8b44b000000000000000000000000000000000000000000000000000000000815261271060048201526001600160a01b039091169250636f8b44b09150602401600060405180830381600087803b158015614be957600080fd5b505af1158015614bfd573d6000803e3d6000fd5b50505050565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090600090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614c6457600080fd5b505af1158015614c78573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614d639190600401619720565b600060405180830381600087803b158015614d7d57600080fd5b505af1158015614d91573d6000803e3d6000fd5b50506023546026546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101879052604481018690529116925063106e62909150606401610e2f565b6000614dfa618f39565b614e05848483614e8e565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c549060440160006040518083038186803b158015614e7a57600080fd5b505afa158015610e5d573d6000803e3d6000fd5b600080614e9b8584614f09565b9050614efe6040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f78790000008152508286604051602001614ee9929190619840565b60405160208183030381529060405285614f15565b9150505b9392505050565b6000614f028383614f43565b60c08101515160009015614f3957614f3284848460c00151614f5e565b9050614f02565b614f328484615104565b6000614f4f83836151ef565b614f0283836020015184614f15565b600080614f696151ff565b90506000614f7786836152d2565b90506000614f8e8260600151836020015185615778565b90506000614f9e8383898961598a565b90506000614fab82616807565b602081015181519192509060030b1561501e57898260400151604051602001614fd5929190619862565b60408051601f19818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261501591600401619720565b60405180910390fd5b60006150616040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a2000000000000000000000008152508360016169d6565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906150b4908490600401619720565b602060405180830381865afa1580156150d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906150f591906198e3565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc92590615159908790600401619720565b600060405180830381865afa158015615176573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261519e91908101906199c5565b905060006151cc82856040516020016151b89291906199fa565b604051602081830303815290604052616bd6565b90506001600160a01b038116614e05578484604051602001614fd5929190619a29565b6151fb82826000616be9565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c90615286908490600401619ad4565b600060405180830381865afa1580156152a3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526152cb9190810190619b1b565b9250505090565b6153046040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d905061534f6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b61535885616cec565b60208201526000615368866170d1565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa1580156153aa573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526153d29190810190619b1b565b868385602001516040516020016153ec9493929190619b64565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb1190615444908590600401619720565b600060405180830381865afa158015615461573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526154899190810190619b1b565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906154d1908490600401619c68565b602060405180830381865afa1580156154ee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906155129190619cba565b6155275781604051602001614fd59190619cdc565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061556c908490600401619d6e565b600060405180830381865afa158015615589573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526155b19190810190619b1b565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906155f8908490600401619dc0565b602060405180830381865afa158015615615573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906156399190619cba565b156156ce576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890615683908490600401619dc0565b600060405180830381865afa1580156156a0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526156c89190810190619b1b565b60408501525b846001600160a01b03166349c4fac88286600001516040516020016156f39190619e12565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161571f929190619e7e565b600060405180830381865afa15801561573c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526157649190810190619b1b565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b60608152602001906001900390816157945790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250816000815181106157f4576157f4619ea3565b60200260200101819052506040518060400160405280600381526020017f2d726c00000000000000000000000000000000000000000000000000000000008152508160018151811061584857615848619ea3565b6020026020010181905250846040516020016158649190619ed2565b6040516020818303038152906040528160028151811061588657615886619ea3565b6020026020010181905250826040516020016158a29190619f3e565b604051602081830303815290604052816003815181106158c4576158c4619ea3565b602002602001018190525060006158da82616807565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184526000808252908601528251808401909352905182529281019290925291925061596b9060408051808201825260008082526020918201528151808301909252845182528085019082015290617354565b6159805785604051602001614fd59190619f7f565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d90156159da565b511590565b615b4e57826020015115615a96576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401615015565b8260c0015115615b4e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401615015565b6040805160ff8082526120008201909252600091816020015b6060815260200190600190039081615b6757905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280615bc29061a010565b935060ff1681518110615bd757615bd7619ea3565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e3700000000000000000000000000000000000000815250604051602001615c28919061a02f565b604051602081830303815290604052828280615c439061a010565b935060ff1681518110615c5857615c58619ea3565b60200260200101819052506040518060400160405280600681526020017f6465706c6f790000000000000000000000000000000000000000000000000000815250828280615ca59061a010565b935060ff1681518110615cba57615cba619ea3565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d65000000000000000000000000000000000000815250828280615d079061a010565b935060ff1681518110615d1c57615d1c619ea3565b60200260200101819052508760200151828280615d389061a010565b935060ff1681518110615d4d57615d4d619ea3565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e747261637450617468000000000000000000000000000000000000815250828280615d9a9061a010565b935060ff1681518110615daf57615daf619ea3565b602090810291909101015287518282615dc78161a010565b935060ff1681518110615ddc57615ddc619ea3565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e49640000000000000000000000000000000000000000000000815250828280615e299061a010565b935060ff1681518110615e3e57615e3e619ea3565b6020026020010181905250615e52466173b5565b8282615e5d8161a010565b935060ff1681518110615e7257615e72619ea3565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c650000000000000000000000000000000000815250828280615ebf9061a010565b935060ff1681518110615ed457615ed4619ea3565b602002602001018190525086828280615eec9061a010565b935060ff1681518110615f0157615f01619ea3565b60209081029190910101528551156160285760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f6465000000000000000000000060208201528282615f528161a010565b935060ff1681518110615f6757615f67619ea3565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d90615fb7908990600401619720565b600060405180830381865afa158015615fd4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615ffc9190810190619b1b565b82826160078161a010565b935060ff168151811061601c5761601c619ea3565b60200260200101819052505b8460200151156160f85760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826160718161a010565b935060ff168151811061608657616086619ea3565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806160d39061a010565b935060ff16815181106160e8576160e8619ea3565b60200260200101819052506162bf565b6161306159d58660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6161c35760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826161738161a010565b935060ff168151811061618857616188619ea3565b60200260200101819052508460a001516040516020016161a89190619ed2565b6040516020818303038152906040528282806160d39061a010565b8460c0015115801561620657506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261620490511590565b155b156162bf5760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261624a8161a010565b935060ff168151811061625f5761625f619ea3565b602002602001018190525061627388617455565b6040516020016162839190619ed2565b60405160208183030381529060405282828061629e9061a010565b935060ff16815181106162b3576162b3619ea3565b60200260200101819052505b604080860151815180830183526000808252602091820152825180840190935281518352908101908201526162f390511590565b6163885760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826163368161a010565b935060ff168151811061634b5761634b619ea3565b602002602001018190525084604001518282806163679061a010565b935060ff168151811061637c5761637c619ea3565b60200260200101819052505b6060850151156164a95760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826163d18161a010565b935060ff16815181106163e6576163e6619ea3565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015616455573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261647d9190810190619b1b565b82826164888161a010565b935060ff168151811061649d5761649d619ea3565b60200260200101819052505b60e085015151156165505760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826164f38161a010565b935060ff168151811061650857616508619ea3565b60200260200101819052506165248560e00151600001516173b5565b828261652f8161a010565b935060ff168151811061654457616544619ea3565b60200260200101819052505b60e085015160200151156165fa5760408051808201909152600a81527f2d2d6761735072696365000000000000000000000000000000000000000000006020820152828261659d8161a010565b935060ff16815181106165b2576165b2619ea3565b60200260200101819052506165ce8560e00151602001516173b5565b82826165d98161a010565b935060ff16815181106165ee576165ee619ea3565b60200260200101819052505b60e085015160400151156166a45760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826166478161a010565b935060ff168151811061665c5761665c619ea3565b60200260200101819052506166788560e00151604001516173b5565b82826166838161a010565b935060ff168151811061669857616698619ea3565b60200260200101819052505b60e0850151606001511561674e5760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826166f18161a010565b935060ff168151811061670657616706619ea3565b60200260200101819052506167228560e00151606001516173b5565b828261672d8161a010565b935060ff168151811061674257616742619ea3565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561676c5761676c6193a0565b60405190808252806020026020018201604052801561679f57816020015b606081526020019060019003908161678a5790505b50905060005b8260ff168160ff1610156167f857838160ff16815181106167c8576167c8619ea3565b6020026020010151828260ff16815181106167e5576167e5619ea3565b60209081029190910101526001016167a5565b5093505050505b949350505050565b61682e6040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c916168b49186910161a09a565b600060405180830381865afa1580156168d1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526168f99190810190619b1b565b905060006169078683617f44565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b81526004016169379190619292565b6000604051808303816000875af1158015616956573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261697e919081019061a0e1565b805190915060030b158015906169975750602081015151155b80156169a65750604081015151155b1561598057816000815181106169be576169be619ea3565b6020026020010151604051602001614fd5919061a197565b60606000616a0b8560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528651825280870190820152909150616a429082905b90618099565b15616b9f576000616abf82616ab984616ab3616a858a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b906180c0565b90618122565b604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150616b23908290618099565b15616b8d57604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616b8a905b82906181a7565b90505b616b96816181cd565b92505050614f02565b8215616bb8578484604051602001614fd592919061a383565b5050604080516020810190915260008152614f02565b509392505050565b6000808251602084016000f09392505050565b8160a0015115616bf857505050565b6000616c05848484618236565b90506000616c1282616807565b602081015181519192509060030b158015616cae5750604080518082018252600781527f535543434553530000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616cae90604080518082018252600080825260209182015281518083019092528451825280850190820152616a3c565b15616cbb57505050505050565b60408201515115616cdb578160400151604051602001614fd5919061a42a565b80604051602001614fd5919061a488565b60606000616d218360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150616d86905b8290617354565b15616df557604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152614f0290616df09083906187d1565b6181cd565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616e57905b829061885b565b600103616f2457604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616ebd90616b83565b50604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152614f0290616df0905b83906181a7565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616f8390616d7f565b156170ba57604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290616feb9083906188f5565b905060008160018351616ffe919061a4f3565b8151811061700e5761700e619ea3565b602002602001015190506170b1616df06170846040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528551825280860190820152906187d1565b95945050505050565b82604051602001614fd5919061a506565b50919050565b606060006171068360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061716890616d7f565b1561717657614f02816181cd565b604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526171d590616e50565b60010361723f57604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152614f0290616df090616f1d565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261729e90616d7f565b156170ba57604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201819052845180860190955292518452830152906173069083906188f5565b9050600181511115617342578060028251617321919061a4f3565b8151811061733157617331619ea3565b602002602001015192505050919050565b5082604051602001614fd5919061a506565b80518251600091111561736957506000614e09565b8151835160208501516000929161737f91619551565b617389919061a4f3565b9050826020015181036173a0576001915050614e09565b82516020840151819020912014905092915050565b606060006173c28361899a565b600101905060008167ffffffffffffffff8111156173e2576173e26193a0565b6040519080825280601f01601f19166020018201604052801561740c576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461741657509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e53454400000000000000000000000000000000000000000000818401908152855180870187528381528401929092528451808601909552518452908301526060916174e1905b8290618a7c565b1561752157505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e7365000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617580906174da565b156175c057505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261761f906174da565b1561765f57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526176be906174da565b806177235750604080518082018252601081527f47504c2d322e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617723906174da565b1561776357505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526177c2906174da565b806178275750604080518082018252601081527f47504c2d332e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617827906174da565b1561786757505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526178c6906174da565b8061792b5750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261792b906174da565b1561796b57505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526179ca906174da565b80617a2f5750604080518082018252601181527f4c47504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617a2f906174da565b15617a6f57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617ace906174da565b15617b0e57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617b6d906174da565b15617bad57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617c0c906174da565b15617c4c57505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617cab906174da565b15617ceb57505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e300000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617d4a906174da565b15617d8a57505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617de9906174da565b80617e4e5750604080518082018252601181527f4147504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617e4e906174da565b15617e8e57505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e3100000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617eed906174da565b15617f2d57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151614fd5929060200161a5e4565b60608060005b8451811015617fcf5781858281518110617f6657617f66619ea3565b6020026020010151604051602001617f7f9291906199fa565b604051602081830303815290604052915060018551617f9e919061a4f3565b8114617fc75781604051602001617fb5919061a74d565b60405160208183030381529060405291505b600101617f4a565b5060408051600380825260808201909252600091816020015b6060815260200190600190039081617fe8579050509050838160008151811061801357618013619ea3565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061806757618067619ea3565b6020026020010181905250818160028151811061808657618086619ea3565b6020908102919091010152949350505050565b60208083015183518351928401516000936180b79291849190618a90565b14159392505050565b604080518082019091526000808252602082015260006180f28460000151856020015185600001518660200151618ba1565b9050836020015181618104919061a4f3565b8451859061811390839061a4f3565b90525060208401525090919050565b6040805180820190915260008082526020820152815183511015618147575081614e09565b602080830151908401516001911461816e5750815160208481015190840151829020919020145b801561819f5782518451859061818590839061a4f3565b905250825160208501805161819b908390619551565b9052505b509192915050565b60408051808201909152600080825260208201526181c6838383618cc1565b5092915050565b60606000826000015167ffffffffffffffff8111156181ee576181ee6193a0565b6040519080825280601f01601f191660200182016040528015618218576020820181803683370190505b50905060006020820190506181c68185602001518660000151618d6c565b606060006182426151ff565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161825f57905050905060006040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806182ba9061a010565b935060ff16815181106182cf576182cf619ea3565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e3300000000000000000000000000000000000000000000000000815250604051602001618320919061a78e565b60405160208183030381529060405282828061833b9061a010565b935060ff168151811061835057618350619ea3565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061839d9061a010565b935060ff16815181106183b2576183b2619ea3565b6020026020010181905250826040516020016183ce9190619f3e565b6040516020818303038152906040528282806183e99061a010565b935060ff16815181106183fe576183fe619ea3565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061844b9061a010565b935060ff168151811061846057618460619ea3565b60200260200101819052506184758784618de6565b82826184808161a010565b935060ff168151811061849557618495619ea3565b6020908102919091010152855151156185415760408051808201909152600b81527f2d2d7265666572656e6365000000000000000000000000000000000000000000602082015282826184e78161a010565b935060ff16815181106184fc576184fc619ea3565b6020026020010181905250618515866000015184618de6565b82826185208161a010565b935060ff168151811061853557618535619ea3565b60200260200101819052505b8560800151156185af5760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261858a8161a010565b935060ff168151811061859f5761859f619ea3565b6020026020010181905250618615565b84156186155760408051808201909152601281527f2d2d726571756972655265666572656e63650000000000000000000000000000602082015282826185f48161a010565b935060ff168151811061860957618609619ea3565b60200260200101819052505b604086015151156186b15760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261865f8161a010565b935060ff168151811061867457618674619ea3565b602002602001018190525085604001518282806186909061a010565b935060ff16815181106186a5576186a5619ea3565b60200260200101819052505b85606001511561871b5760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826186fa8161a010565b935060ff168151811061870f5761870f619ea3565b60200260200101819052505b60008160ff1667ffffffffffffffff811115618739576187396193a0565b60405190808252806020026020018201604052801561876c57816020015b60608152602001906001900390816187575790505b50905060005b8260ff168160ff1610156187c557838160ff168151811061879557618795619ea3565b6020026020010151828260ff16815181106187b2576187b2619ea3565b6020908102919091010152600101618772565b50979650505050505050565b60408051808201909152600080825260208201528151835110156187f6575081614e09565b8151835160208501516000929161880c91619551565b618816919061a4f3565b60208401519091506001908214618837575082516020840151819020908220145b80156188525783518551869061884e90839061a4f3565b9052505b50929392505050565b600080826000015161887f8560000151866020015186600001518760200151618ba1565b6188899190619551565b90505b8351602085015161889d9190619551565b81116181c657816188ad8161a7d3565b92505082600001516188e48560200151836188c8919061a4f3565b86516188d4919061a4f3565b8386600001518760200151618ba1565b6188ee9190619551565b905061888c565b60606000618903848461885b565b61890e906001619551565b67ffffffffffffffff811115618926576189266193a0565b60405190808252806020026020018201604052801561895957816020015b60608152602001906001900390816189445790505b50905060005b8151811015616bce57618975616df086866181a7565b82828151811061898757618987619ea3565b602090810291909101015260010161895f565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106189e3577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310618a0f576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310618a2d57662386f26fc10000830492506010015b6305f5e1008310618a45576305f5e100830492506008015b6127108310618a5957612710830492506004015b60648310618a6b576064830492506002015b600a8310614e095760010192915050565b6000618a888383618e26565b159392505050565b600080858411618b975760208411618b435760008415618adb576001618ab786602061a4f3565b618ac290600861a7ed565b618acd90600261a8eb565b618ad7919061a4f3565b1990505b8351811685618aea8989619551565b618af4919061a4f3565b805190935082165b818114618b2e57878411618b1657879450505050506167ff565b83618b208161a8f7565b945050828451169050618afc565b618b388785619551565b9450505050506167ff565b838320618b50858861a4f3565b618b5a9087619551565b91505b858210618b9557848220808203618b8257618b788684619551565b93505050506167ff565b618b8d60018461a4f3565b925050618b5d565b505b5092949350505050565b60008381868511618cac5760208511618c5b5760008515618bed576001618bc987602061a4f3565b618bd490600861a7ed565b618bdf90600261a8eb565b618be9919061a4f3565b1990505b84518116600087618bfe8b8b619551565b618c08919061a4f3565b855190915083165b828114618c4d57818610618c3557618c288b8b619551565b96505050505050506167ff565b85618c3f8161a7d3565b965050838651169050618c10565b8596505050505050506167ff565b508383206000905b618c6d868961a4f3565b8211618caa57858320808203618c8957839450505050506167ff565b618c94600185619551565b9350508180618ca29061a7d3565b925050618c63565b505b618cb68787619551565b979650505050505050565b60408051808201909152600080825260208201526000618cf38560000151866020015186600001518760200151618ba1565b602080870180519186019190915251909150618d0f908261a4f3565b835284516020860151618d229190619551565b8103618d315760008552618d63565b83518351618d3f9190619551565b85518690618d4e90839061a4f3565b9052508351618d5d9082619551565b60208601525b50909392505050565b60208110618da45781518352618d83602084619551565b9250618d90602083619551565b9150618d9d60208261a4f3565b9050618d6c565b6000198115618dd3576001618dba83602061a4f3565b618dc69061010061a8eb565b618dd0919061a4f3565b90505b9151835183169219169190911790915250565b60606000618df484846152d2565b8051602080830151604051939450618e0e9390910161a90e565b60405160208183030381529060405291505092915050565b8151815160009190811115618e39575081515b6020808501519084015160005b83811015618ef25782518251808214618ec2576000196020871015618ea157600184618e7389602061a4f3565b618e7d9190619551565b618e8890600861a7ed565b618e9390600261a8eb565b618e9d919061a4f3565b1990505b8181168382168181039114618ebf579750614e099650505050505050565b50505b618ecd602086619551565b9450618eda602085619551565b93505050602081618eeb9190619551565b9050618e46565b5084518651615980919061a966565b6112a6806200a98783390190565b611e1b806200bc2d83390190565b6119d2806200da4883390190565b610f0e806200f41a83390190565b6040518060e00160405280606081526020016060815260200160608152602001600015158152602001600015158152602001600015158152602001618f7c618f81565b905290565b60405180610100016040528060001515815260200160001515815260200160608152602001600080191681526020016060815260200160608152602001600015158152602001618f7c6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b818110156190335783516001600160a01b031683526020938401939092019160010161900c565b509095945050505050565b60005b83811015619059578181015183820152602001619041565b50506000910152565b6000815180845261907a81602086016020860161903e565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561918a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b81811015619170577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261915a848651619062565b6020958601959094509290920191600101619120565b5091975050506020948501949290920191506001016190b6565b50929695505050505050565b600081518084526020840193506020830160005b828110156191ea5781517fffffffff00000000000000000000000000000000000000000000000000000000168652602095860195909101906001016191aa565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561918a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051604087526192606040880182619062565b905060208201519150868103602088015261927b8183619196565b96505050602093840193919091019060010161921c565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561918a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184526192f4858351619062565b945060209384019391909101906001016192ba565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561918a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261938a6040870182619196565b9550506020938401939190910190600101619331565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c908216806193e357607f821691505b6020821081036170cb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f821115610b4a57806000526020600020601f840160051c810160208510156194435750805b601f840160051c820191505b81811015612217576000815560010161944f565b815167ffffffffffffffff81111561947d5761947d6193a0565b6194918161948b84546193cf565b8461941c565b6020601f8211600181146194c557600083156194ad5750848201515b600019600385901b1c1916600184901b178455612217565b600084815260208120601f198516915b828110156194f557878501518255602094850194600190920191016194d5565b50848210156195135786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115614e0957614e09619522565b6001600160a01b038154168252600060018201546001600160a01b038116602085015267ffffffffffffffff8160a01c166040850152506002820160806060850152600081546195b3816193cf565b80608088015260018216600081146195d2576001811461960c57619640565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660a089015260a082151560051b8901019350619640565b84600052602060002060005b838110156196375781548a820160a00152600190910190602001619618565b890160a0019450505b50919695505050505050565b6001600160a01b038516815283602082015260a06040820152600061967460a0830185619062565b600060608401528281036080840152618cb68185619564565b60006020828403121561969f57600080fd5b5051919050565b6001600160a01b03841681528260208201526060604082015260006170b16060830184619062565b8281526040602082015260006167ff6040830184619062565b6001600160a01b038516815283602082015260806040820152600061970f6080830185619062565b905082606083015295945050505050565b602081526000614f026020830184619062565b6001600160a01b038616815284602082015260a06040820152600061975b60a0830186619062565b84606084015282810360808401526197738185619564565b98975050505050505050565b6001600160a01b03841681528260208201526080604082015260006197a76080830184619062565b905060006060830152949350505050565b6000826197ee577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6001600160a01b03831681526040602082015260006167ff6040830184619564565b83815260606020820152600061982e6060830185619062565b82810360408401526159808185619564565b6001600160a01b03831681526040602082015260006167ff6040830184619062565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161989a81601a85016020880161903e565b7f3a20000000000000000000000000000000000000000000000000000000000000601a9184019182015283516198d781601c84016020880161903e565b01601c01949350505050565b6000602082840312156198f557600080fd5b81516001600160a01b0381168114614f0257600080fd5b6040516060810167ffffffffffffffff8111828210171561992f5761992f6193a0565b60405290565b60008067ffffffffffffffff841115619950576199506193a0565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561997f5761997f6193a0565b60405283815290508082840185101561999757600080fd5b616bce84602083018561903e565b600082601f8301126199b657600080fd5b614f0283835160208501619935565b6000602082840312156199d757600080fd5b815167ffffffffffffffff8111156199ee57600080fd5b614e05848285016199a5565b60008351619a0c81846020880161903e565b835190830190619a2081836020880161903e565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e747261637420000000000000815260008351619a6181601a85016020880161903e565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a918401918201528351619a9e81603384016020880161903e565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000614f026080830184619062565b600060208284031215619b2d57600080fd5b815167ffffffffffffffff811115619b4457600080fd5b8201601f81018413619b5557600080fd5b614e0584825160208401619935565b60008551619b76818460208a0161903e565b7f2f000000000000000000000000000000000000000000000000000000000000009083019081528551619bb0816001840160208a0161903e565b7f2f00000000000000000000000000000000000000000000000000000000000000600192909101918201528451619bee81600284016020890161903e565b6001818301019150507f2f0000000000000000000000000000000000000000000000000000000000000060018201528351619c3081600284016020880161903e565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b604081526000619c7b6040830184619062565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b600060208284031215619ccc57600080fd5b81518015158114614f0257600080fd5b7f436f756c64206e6f742066696e642041535420696e2061727469666163742000815260008251619d1481601f85016020870161903e565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b604081526000619d816040830184619062565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081526000619dd36040830184619062565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b27000000000000000000000000815260008251619e4a81601485016020870161903e565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b604081526000619e916040830185619062565b8281036020840152614efe8185619062565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f2200000000000000000000000000000000000000000000000000000000000000815260008251619f0a81600185016020870161903e565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b60008251619f5081846020870161903e565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161a00381604b85016020870161903e565b91909101604b0192915050565b600060ff821660ff810361a0265761a026619522565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161a08d81602985016020870161903e565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000614f026080830184619062565b60006020828403121561a0f357600080fd5b815167ffffffffffffffff81111561a10a57600080fd5b82016060818503121561a11c57600080fd5b61a12461990c565b81518060030b811461a13557600080fd5b8152602082015167ffffffffffffffff81111561a15157600080fd5b61a15d868285016199a5565b602083015250604082015167ffffffffffffffff81111561a17d57600080fd5b61a189868285016199a5565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161a1f581602185016020870161903e565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161a3e181602185016020880161903e565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161a41e81602e84016020880161903e565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161a08d81602985016020870161903e565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161a4e681602285016020870161903e565b9190910160220192915050565b81810381811115614e0957614e09619522565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161a53e81600e85016020870161903e565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161a61c81601885016020880161903e565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161a65981601c84016020880161903e565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161a75f81846020870161903e565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161a7c681601c85016020870161903e565b91909101601c0192915050565b6000600019820361a7e65761a7e6619522565b5060010190565b8082028115828204841417614e0957614e09619522565b6001815b600184111561a83f5780850481111561a8235761a823619522565b600184161561a83157908102905b60019390931c92800261a808565b935093915050565b60008261a85657506001614e09565b8161a86357506000614e09565b816001811461a879576002811461a8835761a89f565b6001915050614e09565b60ff84111561a8945761a894619522565b50506001821b614e09565b5060208310610133831016604e8410600b841016171561a8c2575081810a614e09565b61a8cf600019848461a804565b806000190482111561a8e35761a8e3619522565b029392505050565b6000614f02838361a847565b60008161a9065761a906619522565b506000190190565b6000835161a92081846020880161903e565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161a95a81600184016020880161903e565b01600101949350505050565b81810360008312801583831316838312821617156181c6576181c661952256fe608060405234801561001057600080fd5b506040516112a63803806112a683398101604081905261002f91610110565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007891906101e2565b50600461008582826101e2565b5050506001600160a01b03821615806100a557506001600160a01b038116155b156100c35760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b031991821617909155600780549290931691161790556102a0565b80516001600160a01b038116811461010b57600080fd5b919050565b6000806040838503121561012357600080fd5b61012c836100f4565b915061013a602084016100f4565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061016d57607f821691505b60208210810361018d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101dd57806000526020600020601f840160051c810160208510156101ba5750805b601f840160051c820191505b818110156101da57600081556001016101c6565b50505b505050565b81516001600160401b038111156101fb576101fb610143565b61020f816102098454610159565b84610193565b6020601f821160018114610243576000831561022b5750848201515b600019600385901b1c1916600184901b1784556101da565b600084815260208120601f198516915b828110156102735787850151825560209485019460019092019101610253565b50848210156102915786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610ff7806102af6000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806342966c68116100b257806379cc679011610081578063a9059cbb11610066578063a9059cbb1461028e578063bff9662a146102a1578063dd62ed3e146102c157600080fd5b806379cc67901461027357806395d89b411461028657600080fd5b806342966c68146102025780635b1125911461021557806370a0823114610235578063779e3b631461026b57600080fd5b80631e458bee116100ee5780631e458bee1461018857806323b872dd1461019b578063313ce567146101ae578063328a01d0146101bd57600080fd5b806306fdde0314610120578063095ea7b31461013e57806315d57fd41461016157806318160ddd14610176575b600080fd5b610128610307565b6040516101359190610d97565b60405180910390f35b61015161014c366004610e2c565b610399565b6040519015158152602001610135565b61017461016f366004610e56565b6103b3565b005b6002545b604051908152602001610135565b610174610196366004610e89565b61057e565b6101516101a9366004610ebc565b610631565b60405160128152602001610135565b6007546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610135565b610174610210366004610ef9565b610655565b6006546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a610243366004610f12565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b610174610662565b610174610281366004610e2c565b610786565b610128610837565b61015161029c366004610e2c565b610846565b6005546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a6102cf366004610e56565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60606003805461031690610f34565b80601f016020809104026020016040519081016040528092919081815260200182805461034290610f34565b801561038f5780601f106103645761010080835404028352916020019161038f565b820191906000526020600020905b81548152906001019060200180831161037257829003601f168201915b5050505050905090565b6000336103a7818585610854565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103f3575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610431576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82161580610468575073ffffffffffffffffffffffffffffffffffffffff8116155b1561049f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105d1576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6105db8383610866565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161062491815260200190565b60405180910390a3505050565b60003361063f8582856108c6565b61064a858585610995565b506001949350505050565b61065f3382610a40565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106b5576040517fe700765e000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b60065473ffffffffffffffffffffffffffffffffffffffff16610704576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107d9576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6107e38282610a9c565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161082b91815260200190565b60405180910390a25050565b60606004805461031690610f34565b6000336103a7818585610995565b6108618383836001610ab1565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108b6576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c260008383610bf9565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461098f5781811015610980576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610428565b61098f84848484036000610ab1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109e5576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8216610a35576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b610861838383610bf9565b73ffffffffffffffffffffffffffffffffffffffff8216610a90576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c282600083610bf9565b610aa78233836108c6565b6108c28282610a40565b73ffffffffffffffffffffffffffffffffffffffff8416610b01576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8316610b51576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600160209081526040808320938716835292905220829055801561098f578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610beb91815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c31578060026000828254610c269190610f87565b90915550610ce39050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610cb7576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610428565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610d0c57600280548290039055610d38565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161062491815260200190565b602081526000825180602084015260005b81811015610dc55760208186018101516040868401015201610da8565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610e2757600080fd5b919050565b60008060408385031215610e3f57600080fd5b610e4883610e03565b946020939093013593505050565b60008060408385031215610e6957600080fd5b610e7283610e03565b9150610e8060208401610e03565b90509250929050565b600080600060608486031215610e9e57600080fd5b610ea784610e03565b95602085013595506040909401359392505050565b600080600060608486031215610ed157600080fd5b610eda84610e03565b9250610ee860208501610e03565b929592945050506040919091013590565b600060208284031215610f0b57600080fd5b5035919050565b600060208284031215610f2457600080fd5b610f2d82610e03565b9392505050565b600181811c90821680610f4857607f821691505b602082108103610f81577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b808201808211156103ad577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea264697066735822122001ec0ce060384773f3d3389fab7bed652c6b8ee389a7471cce10d00d87a75a0c64736f6c634300081a003360a060405234801561001057600080fd5b50604051611e1b380380611e1b83398101604081905261002f916101fd565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03838116608052600480546001600160a01b0319169184169190911790556100bb60008261014c565b506100e67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261014c565b506101117f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361014c565b5061012a600080516020611dfb8339815191528261014c565b50610143600080516020611dfb8339815191528361014c565b50505050610240565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166101d75760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016101db565b5060005b92915050565b80516001600160a01b03811681146101f857600080fd5b919050565b60008060006060848603121561021257600080fd5b61021b846101e1565b9250610229602085016101e1565b9150610237604085016101e1565b90509250925092565b608051611b84610277600039600081816101d501528181610574015281816105c9015281816107e6015261083b0152611b846000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c80638456cb59116100ee578063a217fddf11610097578063d9caed1211610071578063d9caed12146103e0578063e609055e146103f3578063e63ab1e914610406578063eab103df1461042d57600080fd5b8063a217fddf146103a2578063d547741f146103aa578063d936547e146103bd57600080fd5b8063950837aa116100c8578063950837aa146103695780639a5904271461037c5780639b19251a1461038f57600080fd5b80638456cb591461030157806385f438c11461030957806391d148541461033057600080fd5b806336568abe116101505780635b1125911161012a5780635b112591146102d05780635c975abb146102e35780635eb72e15146102ee57600080fd5b806336568abe1461028e5780633f4ba83a146102a1578063570618e1146102a957600080fd5b8063248a9ca311610181578063248a9ca314610224578063252f07bf146102565780632f2ff15d1461027b57600080fd5b806301ffc9a7146101a8578063116191b6146101d057806321fc65f21461020f575b600080fd5b6101bb6101b636600461155e565b610440565b60405190151581526020015b60405180910390f35b6101f77f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101c7565b61022261021d3660046115fe565b6104d9565b005b610248610232366004611671565b6000908152600160208190526040909120015490565b6040519081526020016101c7565b6004546101bb9074010000000000000000000000000000000000000000900460ff1681565b61022261028936600461168a565b610699565b61022261029c36600461168a565b6106c5565b610222610716565b6102487f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b6004546101f7906001600160a01b031681565b60025460ff166101bb565b6102226102fc3660046116ba565b61074b565b610222610910565b6102487f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101bb61033e36600461168a565b60009182526001602090815260408084206001600160a01b0393909316845291905290205460ff1690565b61022261037736600461175d565b610942565b61022261038a36600461175d565b610ac0565b61022261039d36600461175d565b610b74565b610248600081565b6102226103b836600461168a565b610c2b565b6101bb6103cb36600461175d565b60036020526000908152604090205460ff1681565b6102226103ee36600461177a565b610c51565b6102226104013660046117bb565b610d49565b6102487f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b61022261043b36600461185a565b610f75565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104d357507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104e1610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461050b8161100e565b610513611018565b6001600160a01b03851660009081526003602052604090205460ff16610565576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105996001600160a01b0386167f000000000000000000000000000000000000000000000000000000000000000086611057565b6040517f5131ab590000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906106069088908a908990899089906004016118c0565b600060405180830381600087803b15801561062057600080fd5b505af1158015610634573d6000803e3d6000fd5b50505050846001600160a01b0316866001600160a01b03167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d586868660405161067f93929190611903565b60405180910390a3506106926001600055565b5050505050565b600082815260016020819052604090912001546106b58161100e565b6106bf83836110cb565b50505050565b6001600160a01b0381163314610707576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610711828261115e565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107408161100e565b6107486111e5565b50565b610753610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461077d8161100e565b610785611018565b6001600160a01b03861660009081526003602052604090205460ff166107d7576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61080b6001600160a01b0387167f000000000000000000000000000000000000000000000000000000000000000087611057565b6040517f36c99f3b0000000000000000000000000000000000000000000000000000000081526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906336c99f3b9061087a9089908b908a908a908a908a906004016119f3565b600060405180830381600087803b15801561089457600080fd5b505af11580156108a8573d6000803e3d6000fd5b50505050856001600160a01b0316876001600160a01b03167f84a6ffacc2d811bfa095a126654bc0ae038aac9eeed4390895bf3b3228f17b37878787876040516108f59493929190611a4a565b60405180910390a3506109086001600055565b505050505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61093a8161100e565b610748611237565b600061094d8161100e565b6001600160a01b03821661098d576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004546109c4907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4906001600160a01b031661115e565b506004546109fc907f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a906001600160a01b031661115e565b50610a277f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4836110cb565b50610a527f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a836110cb565b50600480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040519081527f086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba9060200160405180910390a15050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a610aea8161100e565b6001600160a01b038216610b2a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a610b9e8161100e565b6001600160a01b038216610bde576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b60008281526001602081905260409091200154610c478161100e565b6106bf838361115e565b610c59610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610c838161100e565b610c8b611018565b6001600160a01b03831660009081526003602052604090205460ff16610cdd576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cf16001600160a01b0384168584611057565b826001600160a01b0316846001600160a01b03167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610d3691815260200190565b60405180910390a3506107116001600055565b610d51610fcb565b610d59611018565b60045474010000000000000000000000000000000000000000900460ff16610dad576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03841660009081526003602052604090205460ff16610dff576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610e5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e839190611a76565b9050610e9a6001600160a01b038616333087611274565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa158015610f21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f459190611a76565b610f4f9190611a8f565b8787604051610f62959493929190611ac9565b60405180910390a2506109086001600055565b6000610f808161100e565b506004805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b600260005403611007576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61074881336112ad565b60025460ff1615611055576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6040516001600160a01b0383811660248301526044820183905261071191859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611324565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166111565760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016104d3565b5060006104d3565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff16156111565760008381526001602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104d3565b6111ed6113a0565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b61123f611018565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861121a3390565b6040516001600160a01b0384811660248301528381166044830152606482018390526106bf9186918216906323b872dd90608401611084565b60008281526001602090815260408083206001600160a01b038516845290915290205460ff16611320576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044015b60405180910390fd5b5050565b60006113396001600160a01b038416836113dc565b9050805160001415801561135e57508080602001905181019061135c9190611b02565b155b15610711576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611317565b60025460ff16611055576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606113ea838360006113f1565b9392505050565b60608147101561142f576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611317565b600080856001600160a01b0316848660405161144b9190611b1f565b60006040518083038185875af1925050503d8060008114611488576040519150601f19603f3d011682016040523d82523d6000602084013e61148d565b606091505b509150915061149d8683836114a7565b9695505050505050565b6060826114bc576114b78261151c565b6113ea565b81511580156114d357506001600160a01b0384163b155b15611515576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611317565b50806113ea565b80511561152c5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561157057600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146113ea57600080fd5b6001600160a01b038116811461074857600080fd5b60008083601f8401126115c757600080fd5b50813567ffffffffffffffff8111156115df57600080fd5b6020830191508360208285010111156115f757600080fd5b9250929050565b60008060008060006080868803121561161657600080fd5b8535611621816115a0565b94506020860135611631816115a0565b935060408601359250606086013567ffffffffffffffff81111561165457600080fd5b611660888289016115b5565b969995985093965092949392505050565b60006020828403121561168357600080fd5b5035919050565b6000806040838503121561169d57600080fd5b8235915060208301356116af816115a0565b809150509250929050565b60008060008060008060a087890312156116d357600080fd5b86356116de816115a0565b955060208701356116ee816115a0565b945060408701359350606087013567ffffffffffffffff81111561171157600080fd5b61171d89828a016115b5565b909450925050608087013567ffffffffffffffff81111561173d57600080fd5b87016080818a03121561174f57600080fd5b809150509295509295509295565b60006020828403121561176f57600080fd5b81356113ea816115a0565b60008060006060848603121561178f57600080fd5b833561179a816115a0565b925060208401356117aa816115a0565b929592945050506040919091013590565b600080600080600080608087890312156117d457600080fd5b863567ffffffffffffffff8111156117eb57600080fd5b6117f789828a016115b5565b909750955050602087013561180b816115a0565b935060408701359250606087013567ffffffffffffffff81111561182e57600080fd5b61183a89828a016115b5565b979a9699509497509295939492505050565b801515811461074857600080fd5b60006020828403121561186c57600080fd5b81356113ea8161184c565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b03861681526001600160a01b03851660208201528360408201526080606082015260006118f8608083018486611877565b979650505050505050565b83815260406020820152600061191d604083018486611877565b95945050505050565b60008135611933816115a0565b6001600160a01b03168352602082013561194c816115a0565b6001600160a01b03166020840152604082013567ffffffffffffffff811680821461197657600080fd5b6040850152506060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126119b257600080fd5b820160208101903567ffffffffffffffff8111156119cf57600080fd5b8036038213156119de57600080fd5b6080606086015261191d608086018284611877565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a060608201526000611a2b60a083018587611877565b8281036080840152611a3d8185611926565b9998505050505050505050565b848152606060208201526000611a64606083018587611877565b82810360408401526118f88185611926565b600060208284031215611a8857600080fd5b5051919050565b818103818111156104d3577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b606081526000611add606083018789611877565b8560208401528281036040840152611af6818587611877565b98975050505050505050565b600060208284031215611b1457600080fd5b81516113ea8161184c565b6000825160005b81811015611b405760208186018101518583015201611b26565b50600092019182525091905056fea26469706673582212201054222be85c46839f1dc80ff3a2eb3b313883279fc7395a4cd77b868c16ae8964736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60c060405260001960045534801561001657600080fd5b506040516119d23803806119d283398101604081905261003591610238565b60016000819055805460ff19169055838383836001600160a01b038416158061006557506001600160a01b038316155b8061007757506001600160a01b038216155b8061008957506001600160a01b038116155b156100a75760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0384811660805283811660a052600380546001600160a01b0319169184169190911790556100dd60008261016c565b506101087f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361016c565b506101337f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb8361016c565b5061015e7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261016c565b50505050505050505061028c565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166102125760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101ca3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610216565b5060005b92915050565b80516001600160a01b038116811461023357600080fd5b919050565b6000806000806080858703121561024e57600080fd5b6102578561021c565b93506102656020860161021c565b92506102736040860161021c565b91506102816060860161021c565b905092959194509250565b60805160a0516116e26102f0600039600081816102200152818161061f01528181610872015281816109e401528181610ce40152610e060152600081816101d40152818161058f015281816105f2015281816107e2015261084501526116e26000f3fe608060405234801561001057600080fd5b506004361061018d5760003560e01c80635e3e9fef116100e3578063950837aa1161008c578063d547741f11610066578063d547741f146103cf578063d5abeb01146103e2578063e63ab1e9146103eb57600080fd5b8063950837aa1461038d578063a217fddf146103a0578063a783c789146103a857600080fd5b80638456cb59116100bd5780638456cb591461031857806385f438c11461032057806391d148541461034757600080fd5b80635e3e9fef146102df5780636f8b44b0146102f2578063743e0c9b1461030557600080fd5b80632ccb404d116101455780633f4ba83a1161011f5780633f4ba83a146102ac5780635b112591146102b45780635c975abb146102d457600080fd5b80632ccb404d146102735780632f2ff15d1461028657806336568abe1461029957600080fd5b8063116191b611610176578063116191b6146101cf57806321e093b11461021b578063248a9ca31461024257600080fd5b806301ffc9a714610192578063106e6290146101ba575b600080fd5b6101a56101a03660046111c8565b610412565b60405190151581526020015b60405180910390f35b6101cd6101c836600461123a565b6104ab565b005b6101f67f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b1565b6101f67f000000000000000000000000000000000000000000000000000000000000000081565b61026561025036600461126d565b60009081526002602052604090206001015490565b6040519081526020016101b1565b6101cd6102813660046112cf565b610550565b6101cd610294366004611367565b6106ea565b6101cd6102a7366004611367565b610715565b6101cd61076e565b6003546101f69073ffffffffffffffffffffffffffffffffffffffff1681565b60015460ff166101a5565b6101cd6102ed366004611393565b6107a3565b6101cd61030036600461126d565b610938565b6101cd61031336600461126d565b6109a7565b6101cd610a51565b6102657f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101a5610355366004611367565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6101cd61039b3660046113f5565b610a83565b610265600081565b6102657f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101cd6103dd366004611367565b610c2e565b61026560045481565b6102657f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104a557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104b3610c53565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46104dd81610c96565b6104e5610ca0565b6104f0848484610cdf565b8373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d58460405161053891815260200190565b60405180910390a25061054b6001600055565b505050565b610558610c53565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461058281610c96565b61058a610ca0565b6105b57f00000000000000000000000000000000000000000000000000000000000000008785610cdf565b6040517f36c99f3b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906336c99f3b90610651907f0000000000000000000000000000000000000000000000000000000000000000908b908b908b908b908a90600401611545565b600060405180830381600087803b15801561066b57600080fd5b505af115801561067f573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff167fa99787224bed6d592c223365ff0f91d203cd2eea00f6b1be8b0263b835378017878787866040516106cf94939291906115b6565b60405180910390a2506106e26001600055565b505050505050565b60008281526002602052604090206001015461070581610c96565b61070f8383610e67565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610764576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61054b8282610f67565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61079881610c96565b6107a0611026565b50565b6107ab610c53565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46107d581610c96565b6107dd610ca0565b6108087f00000000000000000000000000000000000000000000000000000000000000008684610cdf565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906108a2907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a906004016115ed565b600060405180830381600087803b1580156108bc57600080fd5b505af11580156108d0573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d86868660405161091e9392919061163f565b60405180910390a2506109316001600055565b5050505050565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61096281610c96565b61096a610ca0565b60048290556040518281527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c906020015b60405180910390a15050565b6109af610ca0565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b158015610a3d57600080fd5b505af1158015610931573d6000803e3d6000fd5b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a7b81610c96565b6107a06110a3565b6000610a8e81610c96565b73ffffffffffffffffffffffffffffffffffffffff8216610adb576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354610b1f907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e49073ffffffffffffffffffffffffffffffffffffffff16610f67565b50600354610b64907f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb9073ffffffffffffffffffffffffffffffffffffffff16610f67565b50610b8f7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610e67565b50610bba7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610e67565b50600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091556040519081527fa38189053f94a2657ffb2b9fc651eddd1606a7cefc9f08d30eb72e3dbb51c1f19060200161099b565b600082815260026020526040902060010154610c4981610c96565b61070f8383610f67565b600260005403610c8f576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6107a081336110fc565b60015460ff1615610cdd576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6004547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d4d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d719190611659565b610d7b9084611672565b1115610db3576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201849052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b158015610e4a57600080fd5b505af1158015610e5e573d6000803e3d6000fd5b50505050505050565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610f5f57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610efd3390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016104a5565b5060006104a5565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610f5f57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104a5565b61102e61118c565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6110ab610ca0565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611079565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16611188576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440160405180910390fd5b5050565b60015460ff16610cdd576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156111da57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461120a57600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461123557600080fd5b919050565b60008060006060848603121561124f57600080fd5b61125884611211565b95602085013595506040909401359392505050565b60006020828403121561127f57600080fd5b5035919050565b60008083601f84011261129857600080fd5b50813567ffffffffffffffff8111156112b057600080fd5b6020830191508360208285010111156112c857600080fd5b9250929050565b60008060008060008060a087890312156112e857600080fd5b6112f187611211565b955060208701359450604087013567ffffffffffffffff81111561131457600080fd5b61132089828a01611286565b90955093505060608701359150608087013567ffffffffffffffff81111561134757600080fd5b87016080818a03121561135957600080fd5b809150509295509295509295565b6000806040838503121561137a57600080fd5b8235915061138a60208401611211565b90509250929050565b6000806000806000608086880312156113ab57600080fd5b6113b486611211565b945060208601359350604086013567ffffffffffffffff8111156113d757600080fd5b6113e388828901611286565b96999598509660600135949350505050565b60006020828403121561140757600080fd5b61120a82611211565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff61147782611211565b16825273ffffffffffffffffffffffffffffffffffffffff61149b60208301611211565b1660208301526000604082013567ffffffffffffffff81168082146114bf57600080fd5b6040850152506060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126114fb57600080fd5b820160208101903567ffffffffffffffff81111561151857600080fd5b80360382131561152757600080fd5b6080606086015261153c608086018284611410565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a06060820152600061159760a083018587611410565b82810360808401526115a98185611459565b9998505050505050505050565b8481526060602082015260006115d0606083018587611410565b82810360408401526115e28185611459565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006115e2608083018486611410565b83815260406020820152600061153c604083018486611410565b60006020828403121561166b57600080fd5b5051919050565b808201808211156104a5577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea26469706673582212202a101f8c334fa2686902c04f67a7cb3214f6f59d39a9c4f2f556e803f034e49d64736f6c634300081a00336080604052348015600f57600080fd5b506001600055610eea806100246000396000f3fe60806040526004361061006e5760003560e01c80636ed701691161004b5780636ed70169146100e0578063c5131691146100f5578063e04d4f9714610115578063f05b6abf1461012857005b8063357fc5a2146100775780635ac1e07014610097578063676cc054146100b757005b3661007557005b005b34801561008357600080fd5b50610075610092366004610724565b610148565b3480156100a357600080fd5b506100756100b2366004610760565b6101de565b6100ca6100c536600461079b565b61021a565b6040516100d79190610896565b60405180910390f35b3480156100ec57600080fd5b5061007561024d565b34801561010157600080fd5b50610075610110366004610724565b610282565b6100756101233660046109ce565b61035d565b34801561013457600080fd5b50610075610143366004610aba565b6103a1565b6101506103d6565b61017273ffffffffffffffffffffffffffffffffffffffff8316338386610419565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101d96001600055565b505050565b7f8f39d9623d22cb24244fa4bf1219c32f0a8c7522a0978f7778ec6e376507e264338260405161020f929190610bed565b60405180910390a150565b6040516060907f3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee90600090a19392505050565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61028a6103d6565b6000610297600285610cf3565b9050806000036102d3576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102f573ffffffffffffffffffffffffffffffffffffffff8416338484610419565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101d96001600055565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa3334858585604051610394959493929190610d2e565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146338484846040516103949493929190610db8565b600260005403610412576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526104ae9085906104b4565b50505050565b60006104d673ffffffffffffffffffffffffffffffffffffffff84168361054f565b905080516000141580156104fb5750808060200190518101906104f99190610e7b565b155b156101d9576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b606061055d83836000610564565b9392505050565b6060814710156105a2576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610546565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105cb9190610e98565b60006040518083038185875af1925050503d8060008114610608576040519150601f19603f3d011682016040523d82523d6000602084013e61060d565b606091505b509150915061061d868383610627565b9695505050505050565b60608261063c57610637826106b6565b61055d565b8151158015610660575073ffffffffffffffffffffffffffffffffffffffff84163b155b156106af576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610546565b508061055d565b8051156106c65780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff8116811461071f57600080fd5b919050565b60008060006060848603121561073957600080fd5b83359250610749602085016106fb565b9150610757604085016106fb565b90509250925092565b60006020828403121561077257600080fd5b813567ffffffffffffffff81111561078957600080fd5b82016080818503121561055d57600080fd5b600080600083850360408112156107b157600080fd5b60208112156107bf57600080fd5b50839250602084013567ffffffffffffffff8111156107dd57600080fd5b8401601f810186136107ee57600080fd5b803567ffffffffffffffff81111561080557600080fd5b86602082840101111561081757600080fd5b939660209190910195509293505050565b60005b8381101561084357818101518382015260200161082b565b50506000910152565b60008151808452610864816020860160208601610828565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061055d602083018461084c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561091f5761091f6108a9565b604052919050565b600082601f83011261093857600080fd5b813567ffffffffffffffff811115610952576109526108a9565b61098360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016108d8565b81815284602083860101111561099857600080fd5b816020850160208301376000918101602001919091529392505050565b80151581146106f857600080fd5b803561071f816109b5565b6000806000606084860312156109e357600080fd5b833567ffffffffffffffff8111156109fa57600080fd5b610a0686828701610927565b935050602084013591506040840135610a1e816109b5565b809150509250925092565b600067ffffffffffffffff821115610a4357610a436108a9565b5060051b60200190565b600082601f830112610a5e57600080fd5b8135610a71610a6c82610a29565b6108d8565b8082825260208201915060208360051b860101925085831115610a9357600080fd5b602085015b83811015610ab0578035835260209283019201610a98565b5095945050505050565b600080600060608486031215610acf57600080fd5b833567ffffffffffffffff811115610ae657600080fd5b8401601f81018613610af757600080fd5b8035610b05610a6c82610a29565b8082825260208201915060208360051b850101925088831115610b2757600080fd5b602084015b83811015610b6957803567ffffffffffffffff811115610b4b57600080fd5b610b5a8b602083890101610927565b84525060209283019201610b2c565b509550505050602084013567ffffffffffffffff811115610b8957600080fd5b610b9586828701610a4d565b925050610757604085016109c3565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610c2b836106fb565b16604082015273ffffffffffffffffffffffffffffffffffffffff610c52602084016106fb565b1660608201526000604083013567ffffffffffffffff8116808214610c7657600080fd5b6080840152506060830135368490037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112610cb257600080fd5b830160208101903567ffffffffffffffff811115610ccf57600080fd5b803603821315610cde57600080fd5b608060a085015261061d60c085018284610ba4565b600082610d29577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610d6360a083018661084c565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610dae578151865260209586019590910190600101610d90565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610e4b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610e3685835161084c565b94506020938401939190910190600101610dfc565b505050508281036040840152610e618186610d7c565b915050610e72606083018415159052565b95945050505050565b600060208284031215610e8d57600080fd5b815161055d816109b5565b60008251610eaa818460208701610828565b919091019291505056fea2646970667358221220a223844ed3c7f0b442329f08a7526e0cb2702943c9de69a2a8c93081f9401d8564736f6c634300081a0033a26469706673582212209da5e0d17deb1ae5659cc31cd865b3c0816b93d9755d2b2e5ccbb162caab95f864736f6c634300081a0033",
}

// ZetaConnectorNonNativeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNonNativeTestMetaData.ABI instead.
var ZetaConnectorNonNativeTestABI = ZetaConnectorNonNativeTestMetaData.ABI

// ZetaConnectorNonNativeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNonNativeTestMetaData.Bin instead.
var ZetaConnectorNonNativeTestBin = ZetaConnectorNonNativeTestMetaData.Bin

// DeployZetaConnectorNonNativeTest deploys a new Ethereum contract, binding an instance of ZetaConnectorNonNativeTest to it.
func DeployZetaConnectorNonNativeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZetaConnectorNonNativeTest, error) {
	parsed, err := ZetaConnectorNonNativeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNonNativeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNonNativeTest{ZetaConnectorNonNativeTestCaller: ZetaConnectorNonNativeTestCaller{contract: contract}, ZetaConnectorNonNativeTestTransactor: ZetaConnectorNonNativeTestTransactor{contract: contract}, ZetaConnectorNonNativeTestFilterer: ZetaConnectorNonNativeTestFilterer{contract: contract}}, nil
}

// ZetaConnectorNonNativeTest is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNonNativeTest struct {
	ZetaConnectorNonNativeTestCaller     // Read-only binding to the contract
	ZetaConnectorNonNativeTestTransactor // Write-only binding to the contract
	ZetaConnectorNonNativeTestFilterer   // Log filterer for contract events
}

// ZetaConnectorNonNativeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonNativeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonNativeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNonNativeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonNativeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNonNativeTestSession struct {
	Contract     *ZetaConnectorNonNativeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ZetaConnectorNonNativeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNonNativeTestCallerSession struct {
	Contract *ZetaConnectorNonNativeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ZetaConnectorNonNativeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNonNativeTestTransactorSession struct {
	Contract     *ZetaConnectorNonNativeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ZetaConnectorNonNativeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNonNativeTestRaw struct {
	Contract *ZetaConnectorNonNativeTest // Generic contract binding to access the raw methods on
}

// ZetaConnectorNonNativeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeTestCallerRaw struct {
	Contract *ZetaConnectorNonNativeTestCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNonNativeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeTestTransactorRaw struct {
	Contract *ZetaConnectorNonNativeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNonNativeTest creates a new instance of ZetaConnectorNonNativeTest, bound to a specific deployed contract.
func NewZetaConnectorNonNativeTest(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNonNativeTest, error) {
	contract, err := bindZetaConnectorNonNativeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTest{ZetaConnectorNonNativeTestCaller: ZetaConnectorNonNativeTestCaller{contract: contract}, ZetaConnectorNonNativeTestTransactor: ZetaConnectorNonNativeTestTransactor{contract: contract}, ZetaConnectorNonNativeTestFilterer: ZetaConnectorNonNativeTestFilterer{contract: contract}}, nil
}

// NewZetaConnectorNonNativeTestCaller creates a new read-only instance of ZetaConnectorNonNativeTest, bound to a specific deployed contract.
func NewZetaConnectorNonNativeTestCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNonNativeTestCaller, error) {
	contract, err := bindZetaConnectorNonNativeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestCaller{contract: contract}, nil
}

// NewZetaConnectorNonNativeTestTransactor creates a new write-only instance of ZetaConnectorNonNativeTest, bound to a specific deployed contract.
func NewZetaConnectorNonNativeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNonNativeTestTransactor, error) {
	contract, err := bindZetaConnectorNonNativeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestTransactor{contract: contract}, nil
}

// NewZetaConnectorNonNativeTestFilterer creates a new log filterer instance of ZetaConnectorNonNativeTest, bound to a specific deployed contract.
func NewZetaConnectorNonNativeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNonNativeTestFilterer, error) {
	contract, err := bindZetaConnectorNonNativeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestFilterer{contract: contract}, nil
}

// bindZetaConnectorNonNativeTest binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNonNativeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNonNativeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNonNativeTest.Contract.ZetaConnectorNonNativeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.ZetaConnectorNonNativeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.ZetaConnectorNonNativeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNonNativeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) ISTEST() (bool, error) {
	return _ZetaConnectorNonNativeTest.Contract.ISTEST(&_ZetaConnectorNonNativeTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) ISTEST() (bool, error) {
	return _ZetaConnectorNonNativeTest.Contract.ISTEST(&_ZetaConnectorNonNativeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeTest.Contract.PAUSERROLE(&_ZetaConnectorNonNativeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeTest.Contract.PAUSERROLE(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeTest.Contract.TSSROLE(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeTest.Contract.TSSROLE(&_ZetaConnectorNonNativeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeTest.Contract.WITHDRAWERROLE(&_ZetaConnectorNonNativeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeTest.Contract.WITHDRAWERROLE(&_ZetaConnectorNonNativeTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) ExcludeArtifacts() ([]string, error) {
	return _ZetaConnectorNonNativeTest.Contract.ExcludeArtifacts(&_ZetaConnectorNonNativeTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _ZetaConnectorNonNativeTest.Contract.ExcludeArtifacts(&_ZetaConnectorNonNativeTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) ExcludeContracts() ([]common.Address, error) {
	return _ZetaConnectorNonNativeTest.Contract.ExcludeContracts(&_ZetaConnectorNonNativeTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _ZetaConnectorNonNativeTest.Contract.ExcludeContracts(&_ZetaConnectorNonNativeTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ZetaConnectorNonNativeTest.Contract.ExcludeSelectors(&_ZetaConnectorNonNativeTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ZetaConnectorNonNativeTest.Contract.ExcludeSelectors(&_ZetaConnectorNonNativeTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) ExcludeSenders() ([]common.Address, error) {
	return _ZetaConnectorNonNativeTest.Contract.ExcludeSenders(&_ZetaConnectorNonNativeTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _ZetaConnectorNonNativeTest.Contract.ExcludeSenders(&_ZetaConnectorNonNativeTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) Failed() (bool, error) {
	return _ZetaConnectorNonNativeTest.Contract.Failed(&_ZetaConnectorNonNativeTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) Failed() (bool, error) {
	return _ZetaConnectorNonNativeTest.Contract.Failed(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _ZetaConnectorNonNativeTest.Contract.TargetArtifactSelectors(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _ZetaConnectorNonNativeTest.Contract.TargetArtifactSelectors(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TargetArtifacts() ([]string, error) {
	return _ZetaConnectorNonNativeTest.Contract.TargetArtifacts(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) TargetArtifacts() ([]string, error) {
	return _ZetaConnectorNonNativeTest.Contract.TargetArtifacts(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TargetContracts() ([]common.Address, error) {
	return _ZetaConnectorNonNativeTest.Contract.TargetContracts(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _ZetaConnectorNonNativeTest.Contract.TargetContracts(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _ZetaConnectorNonNativeTest.Contract.TargetInterfaces(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _ZetaConnectorNonNativeTest.Contract.TargetInterfaces(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ZetaConnectorNonNativeTest.Contract.TargetSelectors(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ZetaConnectorNonNativeTest.Contract.TargetSelectors(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TargetSenders() ([]common.Address, error) {
	return _ZetaConnectorNonNativeTest.Contract.TargetSenders(&_ZetaConnectorNonNativeTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _ZetaConnectorNonNativeTest.Contract.TargetSenders(&_ZetaConnectorNonNativeTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) SetUp() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.SetUp(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.SetUp(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestSexMaxSupplyFailsIfSenderIsNotTss is a paid mutator transaction binding the contract method 0xfdca9052.
//
// Solidity: function testSexMaxSupplyFailsIfSenderIsNotTss() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestSexMaxSupplyFailsIfSenderIsNotTss(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testSexMaxSupplyFailsIfSenderIsNotTss")
}

// TestSexMaxSupplyFailsIfSenderIsNotTss is a paid mutator transaction binding the contract method 0xfdca9052.
//
// Solidity: function testSexMaxSupplyFailsIfSenderIsNotTss() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestSexMaxSupplyFailsIfSenderIsNotTss() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestSexMaxSupplyFailsIfSenderIsNotTss(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestSexMaxSupplyFailsIfSenderIsNotTss is a paid mutator transaction binding the contract method 0xfdca9052.
//
// Solidity: function testSexMaxSupplyFailsIfSenderIsNotTss() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestSexMaxSupplyFailsIfSenderIsNotTss() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestSexMaxSupplyFailsIfSenderIsNotTss(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdraw is a paid mutator transaction binding the contract method 0xd509b16c.
//
// Solidity: function testWithdraw() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testWithdraw")
}

// TestWithdraw is a paid mutator transaction binding the contract method 0xd509b16c.
//
// Solidity: function testWithdraw() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestWithdraw() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdraw(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdraw is a paid mutator transaction binding the contract method 0xd509b16c.
//
// Solidity: function testWithdraw() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestWithdraw() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdraw(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndCallFailsIfMaxSupplyIsReached is a paid mutator transaction binding the contract method 0xaaf74192.
//
// Solidity: function testWithdrawAndCallFailsIfMaxSupplyIsReached() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestWithdrawAndCallFailsIfMaxSupplyIsReached(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testWithdrawAndCallFailsIfMaxSupplyIsReached")
}

// TestWithdrawAndCallFailsIfMaxSupplyIsReached is a paid mutator transaction binding the contract method 0xaaf74192.
//
// Solidity: function testWithdrawAndCallFailsIfMaxSupplyIsReached() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestWithdrawAndCallFailsIfMaxSupplyIsReached() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndCallFailsIfMaxSupplyIsReached(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndCallFailsIfMaxSupplyIsReached is a paid mutator transaction binding the contract method 0xaaf74192.
//
// Solidity: function testWithdrawAndCallFailsIfMaxSupplyIsReached() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestWithdrawAndCallFailsIfMaxSupplyIsReached() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndCallFailsIfMaxSupplyIsReached(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20 is a paid mutator transaction binding the contract method 0x3cba0107.
//
// Solidity: function testWithdrawAndCallReceiveERC20() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestWithdrawAndCallReceiveERC20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveERC20")
}

// TestWithdrawAndCallReceiveERC20 is a paid mutator transaction binding the contract method 0x3cba0107.
//
// Solidity: function testWithdrawAndCallReceiveERC20() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestWithdrawAndCallReceiveERC20() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndCallReceiveERC20(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20 is a paid mutator transaction binding the contract method 0x3cba0107.
//
// Solidity: function testWithdrawAndCallReceiveERC20() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestWithdrawAndCallReceiveERC20() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndCallReceiveERC20(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xc1909972.
//
// Solidity: function testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer")
}

// TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xc1909972.
//
// Solidity: function testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xc1909972.
//
// Solidity: function testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20Partial is a paid mutator transaction binding the contract method 0xdcf7d037.
//
// Solidity: function testWithdrawAndCallReceiveERC20Partial() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestWithdrawAndCallReceiveERC20Partial(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveERC20Partial")
}

// TestWithdrawAndCallReceiveERC20Partial is a paid mutator transaction binding the contract method 0xdcf7d037.
//
// Solidity: function testWithdrawAndCallReceiveERC20Partial() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestWithdrawAndCallReceiveERC20Partial() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndCallReceiveERC20Partial(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20Partial is a paid mutator transaction binding the contract method 0xdcf7d037.
//
// Solidity: function testWithdrawAndCallReceiveERC20Partial() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestWithdrawAndCallReceiveERC20Partial() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndCallReceiveERC20Partial(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveNoParams is a paid mutator transaction binding the contract method 0x49346558.
//
// Solidity: function testWithdrawAndCallReceiveNoParams() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestWithdrawAndCallReceiveNoParams(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveNoParams")
}

// TestWithdrawAndCallReceiveNoParams is a paid mutator transaction binding the contract method 0x49346558.
//
// Solidity: function testWithdrawAndCallReceiveNoParams() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestWithdrawAndCallReceiveNoParams() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndCallReceiveNoParams(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveNoParams is a paid mutator transaction binding the contract method 0x49346558.
//
// Solidity: function testWithdrawAndCallReceiveNoParams() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestWithdrawAndCallReceiveNoParams() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndCallReceiveNoParams(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndFailsIfMaxSupplyIsReached is a paid mutator transaction binding the contract method 0x7db20efb.
//
// Solidity: function testWithdrawAndFailsIfMaxSupplyIsReached() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestWithdrawAndFailsIfMaxSupplyIsReached(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testWithdrawAndFailsIfMaxSupplyIsReached")
}

// TestWithdrawAndFailsIfMaxSupplyIsReached is a paid mutator transaction binding the contract method 0x7db20efb.
//
// Solidity: function testWithdrawAndFailsIfMaxSupplyIsReached() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestWithdrawAndFailsIfMaxSupplyIsReached() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndFailsIfMaxSupplyIsReached(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndFailsIfMaxSupplyIsReached is a paid mutator transaction binding the contract method 0x7db20efb.
//
// Solidity: function testWithdrawAndFailsIfMaxSupplyIsReached() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestWithdrawAndFailsIfMaxSupplyIsReached() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndFailsIfMaxSupplyIsReached(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndRevert is a paid mutator transaction binding the contract method 0xde1cb76c.
//
// Solidity: function testWithdrawAndRevert() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestWithdrawAndRevert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testWithdrawAndRevert")
}

// TestWithdrawAndRevert is a paid mutator transaction binding the contract method 0xde1cb76c.
//
// Solidity: function testWithdrawAndRevert() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestWithdrawAndRevert() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndRevert(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndRevert is a paid mutator transaction binding the contract method 0xde1cb76c.
//
// Solidity: function testWithdrawAndRevert() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestWithdrawAndRevert() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndRevert(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndRevertFailsIfMaxSupplyIsReached is a paid mutator transaction binding the contract method 0x2506ef03.
//
// Solidity: function testWithdrawAndRevertFailsIfMaxSupplyIsReached() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestWithdrawAndRevertFailsIfMaxSupplyIsReached(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testWithdrawAndRevertFailsIfMaxSupplyIsReached")
}

// TestWithdrawAndRevertFailsIfMaxSupplyIsReached is a paid mutator transaction binding the contract method 0x2506ef03.
//
// Solidity: function testWithdrawAndRevertFailsIfMaxSupplyIsReached() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestWithdrawAndRevertFailsIfMaxSupplyIsReached() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndRevertFailsIfMaxSupplyIsReached(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndRevertFailsIfMaxSupplyIsReached is a paid mutator transaction binding the contract method 0x2506ef03.
//
// Solidity: function testWithdrawAndRevertFailsIfMaxSupplyIsReached() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestWithdrawAndRevertFailsIfMaxSupplyIsReached() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndRevertFailsIfMaxSupplyIsReached(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x82832014.
//
// Solidity: function testWithdrawAndRevertFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testWithdrawAndRevertFailsIfSenderIsNotWithdrawer")
}

// TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x82832014.
//
// Solidity: function testWithdrawAndRevertFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x82832014.
//
// Solidity: function testWithdrawAndRevertFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xfe574f84.
//
// Solidity: function testWithdrawFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestWithdrawFailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testWithdrawFailsIfSenderIsNotWithdrawer")
}

// TestWithdrawFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xfe574f84.
//
// Solidity: function testWithdrawFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestWithdrawFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawFailsIfSenderIsNotWithdrawer(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xfe574f84.
//
// Solidity: function testWithdrawFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestWithdrawFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawFailsIfSenderIsNotWithdrawer(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawTogglePause is a paid mutator transaction binding the contract method 0xccb0e3f2.
//
// Solidity: function testWithdrawTogglePause() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestWithdrawTogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testWithdrawTogglePause")
}

// TestWithdrawTogglePause is a paid mutator transaction binding the contract method 0xccb0e3f2.
//
// Solidity: function testWithdrawTogglePause() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestWithdrawTogglePause() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawTogglePause(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestWithdrawTogglePause is a paid mutator transaction binding the contract method 0xccb0e3f2.
//
// Solidity: function testWithdrawTogglePause() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestWithdrawTogglePause() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestWithdrawTogglePause(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// ZetaConnectorNonNativeTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestCalledIterator struct {
	Event *ZetaConnectorNonNativeTestCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestCalled represents a Called event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ZetaConnectorNonNativeTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestCalledIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestCalled)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "Called", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseCalled(log types.Log) (*ZetaConnectorNonNativeTestCalled, error) {
	event := new(ZetaConnectorNonNativeTestCalled)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestDepositedIterator struct {
	Event *ZetaConnectorNonNativeTestDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestDeposited represents a Deposited event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestDeposited struct {
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterDeposited(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ZetaConnectorNonNativeTestDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestDepositedIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestDeposited, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestDeposited)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseDeposited(log types.Log) (*ZetaConnectorNonNativeTestDeposited, error) {
	event := new(ZetaConnectorNonNativeTestDeposited)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestExecutedIterator struct {
	Event *ZetaConnectorNonNativeTestExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestExecuted represents a Executed event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*ZetaConnectorNonNativeTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestExecutedIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestExecuted)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseExecuted(log types.Log) (*ZetaConnectorNonNativeTestExecuted, error) {
	event := new(ZetaConnectorNonNativeTestExecuted)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestExecutedWithERC20Iterator struct {
	Event *ZetaConnectorNonNativeTestExecutedWithERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestExecutedWithERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestExecutedWithERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*ZetaConnectorNonNativeTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestExecutedWithERC20Iterator{contract: _ZetaConnectorNonNativeTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestExecutedWithERC20)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseExecutedWithERC20(log types.Log) (*ZetaConnectorNonNativeTestExecutedWithERC20, error) {
	event := new(ZetaConnectorNonNativeTestExecutedWithERC20)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReceivedERC20Iterator struct {
	Event *ZetaConnectorNonNativeTestReceivedERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestReceivedERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestReceivedERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestReceivedERC20 represents a ReceivedERC20 event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestReceivedERC20Iterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestReceivedERC20Iterator{contract: _ZetaConnectorNonNativeTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestReceivedERC20)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseReceivedERC20(log types.Log) (*ZetaConnectorNonNativeTestReceivedERC20, error) {
	event := new(ZetaConnectorNonNativeTestReceivedERC20)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReceivedNoParamsIterator struct {
	Event *ZetaConnectorNonNativeTestReceivedNoParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestReceivedNoParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestReceivedNoParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestReceivedNoParams represents a ReceivedNoParams event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestReceivedNoParamsIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestReceivedNoParamsIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestReceivedNoParams)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseReceivedNoParams(log types.Log) (*ZetaConnectorNonNativeTestReceivedNoParams, error) {
	event := new(ZetaConnectorNonNativeTestReceivedNoParams)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReceivedNonPayableIterator struct {
	Event *ZetaConnectorNonNativeTestReceivedNonPayable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestReceivedNonPayable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestReceivedNonPayable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestReceivedNonPayable represents a ReceivedNonPayable event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestReceivedNonPayableIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestReceivedNonPayableIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestReceivedNonPayable)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseReceivedNonPayable(log types.Log) (*ZetaConnectorNonNativeTestReceivedNonPayable, error) {
	event := new(ZetaConnectorNonNativeTestReceivedNonPayable)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestReceivedOnCallIterator is returned from FilterReceivedOnCall and is used to iterate over the raw logs and unpacked data for ReceivedOnCall events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReceivedOnCallIterator struct {
	Event *ZetaConnectorNonNativeTestReceivedOnCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestReceivedOnCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestReceivedOnCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestReceivedOnCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestReceivedOnCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestReceivedOnCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestReceivedOnCall represents a ReceivedOnCall event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReceivedOnCall struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCall is a free log retrieval operation binding the contract event 0x3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee.
//
// Solidity: event ReceivedOnCall()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterReceivedOnCall(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestReceivedOnCallIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestReceivedOnCallIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "ReceivedOnCall", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCall is a free log subscription operation binding the contract event 0x3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee.
//
// Solidity: event ReceivedOnCall()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchReceivedOnCall(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestReceivedOnCall) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestReceivedOnCall)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseReceivedOnCall(log types.Log) (*ZetaConnectorNonNativeTestReceivedOnCall, error) {
	event := new(ZetaConnectorNonNativeTestReceivedOnCall)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReceivedPayableIterator struct {
	Event *ZetaConnectorNonNativeTestReceivedPayable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestReceivedPayable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestReceivedPayable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestReceivedPayable represents a ReceivedPayable event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReceivedPayable struct {
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestReceivedPayableIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestReceivedPayableIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestReceivedPayable)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseReceivedPayable(log types.Log) (*ZetaConnectorNonNativeTestReceivedPayable, error) {
	event := new(ZetaConnectorNonNativeTestReceivedPayable)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReceivedRevertIterator struct {
	Event *ZetaConnectorNonNativeTestReceivedRevert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestReceivedRevert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestReceivedRevert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestReceivedRevert represents a ReceivedRevert event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestReceivedRevertIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestReceivedRevertIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestReceivedRevert)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseReceivedRevert(log types.Log) (*ZetaConnectorNonNativeTestReceivedRevert, error) {
	event := new(ZetaConnectorNonNativeTestReceivedRevert)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestRevertedIterator struct {
	Event *ZetaConnectorNonNativeTestReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestReverted represents a Reverted event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestReverted struct {
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ZetaConnectorNonNativeTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestRevertedIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestReverted)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "Reverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseReverted(log types.Log) (*ZetaConnectorNonNativeTestReverted, error) {
	event := new(ZetaConnectorNonNativeTestReverted)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestUpdatedGatewayTSSAddressIterator struct {
	Event *ZetaConnectorNonNativeTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestUpdatedGatewayTSSAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestUpdatedGatewayTSSAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestUpdatedGatewayTSSAddress struct {
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestUpdatedGatewayTSSAddressIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestUpdatedGatewayTSSAddress)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*ZetaConnectorNonNativeTestUpdatedGatewayTSSAddress, error) {
	event := new(ZetaConnectorNonNativeTestUpdatedGatewayTSSAddress)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddressIterator is returned from FilterUpdatedZetaConnectorTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedZetaConnectorTSSAddress events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddressIterator struct {
	Event *ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddress represents a UpdatedZetaConnectorTSSAddress event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddress struct {
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedZetaConnectorTSSAddress is a free log retrieval operation binding the contract event 0xa38189053f94a2657ffb2b9fc651eddd1606a7cefc9f08d30eb72e3dbb51c1f1.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address newTSSAddress)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterUpdatedZetaConnectorTSSAddress(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddressIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "UpdatedZetaConnectorTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedZetaConnectorTSSAddress is a free log subscription operation binding the contract event 0xa38189053f94a2657ffb2b9fc651eddd1606a7cefc9f08d30eb72e3dbb51c1f1.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address newTSSAddress)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchUpdatedZetaConnectorTSSAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddress)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedZetaConnectorTSSAddress is a log parse operation binding the contract event 0xa38189053f94a2657ffb2b9fc651eddd1606a7cefc9f08d30eb72e3dbb51c1f1.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address newTSSAddress)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseUpdatedZetaConnectorTSSAddress(log types.Log) (*ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddress, error) {
	event := new(ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddress)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestWithdrawnIterator struct {
	Event *ZetaConnectorNonNativeTestWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestWithdrawn represents a Withdrawn event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeTestWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestWithdrawnIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestWithdrawn)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseWithdrawn(log types.Log) (*ZetaConnectorNonNativeTestWithdrawn, error) {
	event := new(ZetaConnectorNonNativeTestWithdrawn)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestWithdrawnAndCalledIterator struct {
	Event *ZetaConnectorNonNativeTestWithdrawnAndCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestWithdrawnAndCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestWithdrawnAndCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestWithdrawnAndCalled struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeTestWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestWithdrawnAndCalledIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestWithdrawnAndCalled, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestWithdrawnAndCalled)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*ZetaConnectorNonNativeTestWithdrawnAndCalled, error) {
	event := new(ZetaConnectorNonNativeTestWithdrawnAndCalled)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestWithdrawnAndRevertedIterator struct {
	Event *ZetaConnectorNonNativeTestWithdrawnAndReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestWithdrawnAndReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestWithdrawnAndReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestWithdrawnAndReverted struct {
	To            common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeTestWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestWithdrawnAndRevertedIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestWithdrawnAndReverted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestWithdrawnAndReverted)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*ZetaConnectorNonNativeTestWithdrawnAndReverted, error) {
	event := new(ZetaConnectorNonNativeTestWithdrawnAndReverted)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogIterator struct {
	Event *ZetaConnectorNonNativeTestLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLog represents a Log event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLog(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLog) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLog)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLog(log types.Log) (*ZetaConnectorNonNativeTestLog, error) {
	event := new(ZetaConnectorNonNativeTestLog)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogAddressIterator struct {
	Event *ZetaConnectorNonNativeTestLogAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogAddress represents a LogAddress event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogAddressIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogAddress)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogAddress(log types.Log) (*ZetaConnectorNonNativeTestLogAddress, error) {
	event := new(ZetaConnectorNonNativeTestLogAddress)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogArrayIterator struct {
	Event *ZetaConnectorNonNativeTestLogArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogArray represents a LogArray event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogArrayIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogArrayIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogArray) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogArray)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogArray(log types.Log) (*ZetaConnectorNonNativeTestLogArray, error) {
	event := new(ZetaConnectorNonNativeTestLogArray)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogArray0Iterator struct {
	Event *ZetaConnectorNonNativeTestLogArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogArray0 represents a LogArray0 event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogArray0Iterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogArray0Iterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogArray0)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogArray0(log types.Log) (*ZetaConnectorNonNativeTestLogArray0, error) {
	event := new(ZetaConnectorNonNativeTestLogArray0)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogArray1Iterator struct {
	Event *ZetaConnectorNonNativeTestLogArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogArray1 represents a LogArray1 event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogArray1Iterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogArray1Iterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogArray1)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogArray1(log types.Log) (*ZetaConnectorNonNativeTestLogArray1, error) {
	event := new(ZetaConnectorNonNativeTestLogArray1)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogBytesIterator struct {
	Event *ZetaConnectorNonNativeTestLogBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogBytes represents a LogBytes event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogBytesIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogBytesIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogBytes)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogBytes(log types.Log) (*ZetaConnectorNonNativeTestLogBytes, error) {
	event := new(ZetaConnectorNonNativeTestLogBytes)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogBytes32Iterator struct {
	Event *ZetaConnectorNonNativeTestLogBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogBytes32 represents a LogBytes32 event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogBytes32Iterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogBytes32Iterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogBytes32)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogBytes32(log types.Log) (*ZetaConnectorNonNativeTestLogBytes32, error) {
	event := new(ZetaConnectorNonNativeTestLogBytes32)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogIntIterator struct {
	Event *ZetaConnectorNonNativeTestLogInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogInt represents a LogInt event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogIntIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogIntIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogInt) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogInt)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogInt(log types.Log) (*ZetaConnectorNonNativeTestLogInt, error) {
	event := new(ZetaConnectorNonNativeTestLogInt)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedAddressIterator struct {
	Event *ZetaConnectorNonNativeTestLogNamedAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogNamedAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogNamedAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogNamedAddress represents a LogNamedAddress event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogNamedAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogNamedAddressIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogNamedAddress)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogNamedAddress(log types.Log) (*ZetaConnectorNonNativeTestLogNamedAddress, error) {
	event := new(ZetaConnectorNonNativeTestLogNamedAddress)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedArrayIterator struct {
	Event *ZetaConnectorNonNativeTestLogNamedArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogNamedArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogNamedArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogNamedArray represents a LogNamedArray event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogNamedArrayIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogNamedArrayIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogNamedArray)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogNamedArray(log types.Log) (*ZetaConnectorNonNativeTestLogNamedArray, error) {
	event := new(ZetaConnectorNonNativeTestLogNamedArray)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedArray0Iterator struct {
	Event *ZetaConnectorNonNativeTestLogNamedArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogNamedArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogNamedArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogNamedArray0 represents a LogNamedArray0 event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogNamedArray0Iterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogNamedArray0Iterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogNamedArray0)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogNamedArray0(log types.Log) (*ZetaConnectorNonNativeTestLogNamedArray0, error) {
	event := new(ZetaConnectorNonNativeTestLogNamedArray0)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedArray1Iterator struct {
	Event *ZetaConnectorNonNativeTestLogNamedArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogNamedArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogNamedArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogNamedArray1 represents a LogNamedArray1 event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogNamedArray1Iterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogNamedArray1Iterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogNamedArray1)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogNamedArray1(log types.Log) (*ZetaConnectorNonNativeTestLogNamedArray1, error) {
	event := new(ZetaConnectorNonNativeTestLogNamedArray1)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedBytesIterator struct {
	Event *ZetaConnectorNonNativeTestLogNamedBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogNamedBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogNamedBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogNamedBytes represents a LogNamedBytes event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogNamedBytesIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogNamedBytesIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogNamedBytes)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogNamedBytes(log types.Log) (*ZetaConnectorNonNativeTestLogNamedBytes, error) {
	event := new(ZetaConnectorNonNativeTestLogNamedBytes)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedBytes32Iterator struct {
	Event *ZetaConnectorNonNativeTestLogNamedBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogNamedBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogNamedBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogNamedBytes32Iterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogNamedBytes32)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogNamedBytes32(log types.Log) (*ZetaConnectorNonNativeTestLogNamedBytes32, error) {
	event := new(ZetaConnectorNonNativeTestLogNamedBytes32)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedDecimalIntIterator struct {
	Event *ZetaConnectorNonNativeTestLogNamedDecimalInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogNamedDecimalInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogNamedDecimalInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogNamedDecimalIntIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogNamedDecimalInt)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*ZetaConnectorNonNativeTestLogNamedDecimalInt, error) {
	event := new(ZetaConnectorNonNativeTestLogNamedDecimalInt)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedDecimalUintIterator struct {
	Event *ZetaConnectorNonNativeTestLogNamedDecimalUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogNamedDecimalUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogNamedDecimalUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogNamedDecimalUintIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogNamedDecimalUint)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*ZetaConnectorNonNativeTestLogNamedDecimalUint, error) {
	event := new(ZetaConnectorNonNativeTestLogNamedDecimalUint)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedIntIterator struct {
	Event *ZetaConnectorNonNativeTestLogNamedInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogNamedInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogNamedInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogNamedInt represents a LogNamedInt event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogNamedIntIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogNamedIntIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogNamedInt)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogNamedInt(log types.Log) (*ZetaConnectorNonNativeTestLogNamedInt, error) {
	event := new(ZetaConnectorNonNativeTestLogNamedInt)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedStringIterator struct {
	Event *ZetaConnectorNonNativeTestLogNamedString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogNamedString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogNamedString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogNamedString represents a LogNamedString event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogNamedStringIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogNamedStringIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogNamedString)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogNamedString(log types.Log) (*ZetaConnectorNonNativeTestLogNamedString, error) {
	event := new(ZetaConnectorNonNativeTestLogNamedString)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedUintIterator struct {
	Event *ZetaConnectorNonNativeTestLogNamedUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogNamedUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogNamedUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogNamedUint represents a LogNamedUint event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogNamedUintIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogNamedUintIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogNamedUint)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogNamedUint(log types.Log) (*ZetaConnectorNonNativeTestLogNamedUint, error) {
	event := new(ZetaConnectorNonNativeTestLogNamedUint)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogStringIterator struct {
	Event *ZetaConnectorNonNativeTestLogString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogString represents a LogString event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogString(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogStringIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogStringIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogString) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogString)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogString(log types.Log) (*ZetaConnectorNonNativeTestLogString, error) {
	event := new(ZetaConnectorNonNativeTestLogString)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogUintIterator struct {
	Event *ZetaConnectorNonNativeTestLogUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogUint represents a LogUint event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogUintIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogUintIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogUint) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogUint)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogUint(log types.Log) (*ZetaConnectorNonNativeTestLogUint, error) {
	event := new(ZetaConnectorNonNativeTestLogUint)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogsIterator struct {
	Event *ZetaConnectorNonNativeTestLogs // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestLogs)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestLogs)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestLogs represents a Logs event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterLogs(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestLogsIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestLogsIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestLogs) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestLogs)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "logs", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseLogs(log types.Log) (*ZetaConnectorNonNativeTestLogs, error) {
	event := new(ZetaConnectorNonNativeTestLogs)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
