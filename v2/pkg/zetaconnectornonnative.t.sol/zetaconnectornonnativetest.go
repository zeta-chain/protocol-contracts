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

// ZetaConnectorNonNativeTestMetaData contains all meta data concerning the ZetaConnectorNonNativeTest contract.
var ZetaConnectorNonNativeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testSexMaxSupplyFailsIfSenderIsNotTss\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpgradeAndWithdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallFailsIfMaxSupplyIsReached\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20Partial\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveNoParams\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndFailsIfMaxSupplyIsReached\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevert\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertFailsIfMaxSupplyIsReached\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedZetaConnectorTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxSupply\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5061d3628061003c6000396000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c8063aaf7419211610104578063d509b16c116100a2578063e63ab1e911610071578063e63ab1e914610357578063fa7626d41461037e578063fdca90521461038b578063fe574f841461039357600080fd5b8063d509b16c14610337578063dcf7d0371461033f578063de1cb76c14610347578063e20c9f711461034f57600080fd5b8063b5508aa9116100de578063b5508aa914610307578063ba414fa61461030f578063c190997214610327578063ccb0e3f21461032f57600080fd5b8063aaf74192146102ef578063af298bb1146102f7578063b0464fdc146102ff57600080fd5b8063493465581161017c57806385226c811161014b57806385226c811461026957806385f438c11461027e578063916a17c6146102b3578063a783c789146102c857600080fd5b8063493465581461023c57806366d9a9a0146102445780637db20efb14610259578063828320141461026157600080fd5b80632ade3880116101b85780632ade38801461020f5780633cba0107146102245780633e5e3c231461022c5780633f7286f41461023457600080fd5b80630a9254e4146101df5780631ed7831c146101e95780632506ef0314610207575b600080fd5b6101e761039b565b005b6101f1610b80565b6040516101fe9190619a95565b60405180910390f35b6101e7610be2565b610217610e96565b6040516101fe9190619b31565b6101e7610fd8565b6101f16117af565b6101f161180f565b6101e761186f565b61024c611e4f565b6040516101fe9190619c97565b6101e7611fd1565b6101e7612270565b6102716124d0565b6040516101fe9190619d35565b6102a57f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6040519081526020016101fe565b6102bb6125a0565b6040516101fe9190619dac565b6102a57f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101e761269b565b6101e7612916565b6102bb612d60565b610271612e5b565b610317612f2b565b60405190151581526020016101fe565b6101e7612fff565b6101e761326f565b6101e7613d9c565b6101e76140ef565b6101e761472b565b6101f1614e58565b6102a57f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b601f546103179060ff1681565b6101e7614eb8565b6101e76150b9565b602480547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602580548216611234179055602680548216615678908117909155602780549092166198761790915560405181906103ff906199c2565b6001600160a01b03928316815291166020820152604001604051809103906000f080158015610432573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260265460248054935191861690820152604481019390935292166064820152600091610524916084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b000000000000000000000000000000000000000000000000000000001790526152a9565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0384811682029290921792839055604080518082018252601081527f4552433230437573746f64792e736f6c00000000000000000000000000000000602082015260265460248054935194909604851695840195909552938316604483015290911660648201529192506105ca916084016104c7565b602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691909117909155604080518082018252601a81527f5a657461436f6e6e6563746f724e6f6e4e61746976652e736f6c0000000000006020820152601f546023546026546024805495516101009094048716908401529085166044830152841660648201529190921660848201529192506106cf9160a40160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff8c8765e000000000000000000000000000000000000000000000000000000001790526152a9565b602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038381169190911790915560265460405163ca669fa760e01b815291166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561075a57600080fd5b505af115801561076e573d6000803e3d6000fd5b50506023546026546022546040517f15d57fd40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529082166024820152911692506315d57fd49150604401600060405180830381600087803b1580156107df57600080fd5b505af11580156107f3573d6000803e3d6000fd5b50505050604051610803906199cf565b604051809103906000f08015801561081f573d6000803e3d6000fd5b50602080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556026546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b1580156108cb57600080fd5b505af11580156108df573d6000803e3d6000fd5b5050602480546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d93506306447d56925001600060405180830381600087803b15801561095457600080fd5b505af1158015610968573d6000803e3d6000fd5b5050601f546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015261010090920416925063ae7a3a6f9150602401600060405180830381600087803b1580156109d357600080fd5b505af11580156109e7573d6000803e3d6000fd5b5050601f546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526101009092041692506310188aef9150602401600060405180830381600087803b158015610a5257600080fd5b505af1158015610a66573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610ac857600080fd5b505af1158015610adc573d6000803e3d6000fd5b5050604080516060810182526023546001600160a01b039081168252600160208084019182528451908101855260008152938301849052825160288054925167ffffffffffffffff1674010000000000000000000000000000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090931691909316171781559093509150602990610b7a9082619f0b565b50505050565b60606016805480602002602001604051908101604052809291908181526020018280548015610bd857602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610bba575b5050505050905090565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a09060009060250160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015610c7d57600080fd5b505af1158015610c91573d6000803e3d6000fd5b50506022546040517f6f8b44b0000000000000000000000000000000000000000000000000000000008152600481018690526001600160a01b039091169250636f8b44b09150602401600060405180830381600087803b158015610cf457600080fd5b505af1158015610d08573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015610d6557600080fd5b505af1158015610d79573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fc30436e9000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015610e0257600080fd5b505af1158015610e16573d6000803e3d6000fd5b50506022546020546001600160a01b03918216935063057e0f25925016610e3e856001619ff9565b8460286040518563ffffffff1660e01b8152600401610e60949392919061a0e1565b600060405180830381600087803b158015610e7a57600080fd5b505af1158015610e8e573d6000803e3d6000fd5b505050505050565b6060601e805480602002602001604051908101604052809291908181526020016000905b82821015610fcf57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015610fb8578382906000526020600020018054610f2b90619e72565b80601f0160208091040260200160405190810160405280929190818152602001828054610f5790619e72565b8015610fa45780601f10610f7957610100808354040283529160200191610fa4565b820191906000526020600020905b815481529060010190602001808311610f8757829003601f168201915b505050505081526020019060010190610f0c565b505050508152505081526020019060010190610eba565b50505050905090565b602354602554604051620186a0602482018190526001600160a01b039384166044830152929091166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905260235460255491516370a0823160e01b81526001600160a01b0392831660048201529293506000929116906370a0823190602401602060405180830381865afa1580156110b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110da919061a122565b90506110e78160006152c8565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611138573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115c919061a122565b90506111698160006152c8565b601f546040516101009091046001600160a01b03166024820152604481018690526064810185905260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1e458bee0000000000000000000000000000000000000000000000000000000017905260235490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391611254916001600160a01b039190911690600090869060040161a13b565b600060405180830381600087803b15801561126e57600080fd5b505af1158015611282573d6000803e3d6000fd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561131457600080fd5b505af1158015611328573d6000803e3d6000fd5b5050601f54602354602554604080516101009094046001600160a01b039081168552602085018d9052928316908401521660608201527f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609250608001905060405180910390a16022546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561141c57600080fd5b505af1158015611430573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d9150611475908990889061a163565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156114d657600080fd5b505af11580156114ea573d6000803e3d6000fd5b50506022546020546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef935061154292909116908a9089908b9060040161a17c565b600060405180830381600087803b15801561155c57600080fd5b505af1158015611570573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156115c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115e7919061a122565b90506115f381886152c8565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611644573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611668919061a122565b90506116758160006152c8565b602354601f546020546040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081526101009092046001600160a01b0390811660048401529081166024830152600092169063dd62ed3e90604401602060405180830381865afa1580156116ed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611711919061a122565b905061171e8160006152c8565b602354601f546040516370a0823160e01b81526101009091046001600160a01b03908116600483015260009216906370a0823190602401602060405180830381865afa158015611772573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611796919061a122565b90506117a38160006152c8565b50505050505050505050565b60606018805480602002602001604051908101604052809291908181526020018280548015610bd8576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610bba575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015610bd8576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610bba575050505050905090565b6040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed701690000000000000000000000000000000000000000000000000000000017905260235460255493516370a0823160e01b8152620186a0946000949385936001600160a01b03908116936370a082319361191193921691016001600160a01b0391909116815260200190565b602060405180830381865afa15801561192e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611952919061a122565b905061195f8160006152c8565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156119b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119d4919061a122565b90506119e18160006152c8565b601f546040516101009091046001600160a01b03166024820152604481018690526064810185905260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1e458bee0000000000000000000000000000000000000000000000000000000017905260235490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391611acc916001600160a01b039190911690600090869060040161a13b565b600060405180830381600087803b158015611ae657600080fd5b505af1158015611afa573d6000803e3d6000fd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015611b8c57600080fd5b505af1158015611ba0573d6000803e3d6000fd5b5050601f546040516101009091046001600160a01b031681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09250602001905060405180910390a16022546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611c7757600080fd5b505af1158015611c8b573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d9150611cd0908990889061a163565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611d3157600080fd5b505af1158015611d45573d6000803e3d6000fd5b50506022546020546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef9350611d9d92909116908a9089908b9060040161a17c565b600060405180830381600087803b158015611db757600080fd5b505af1158015611dcb573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015611e1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e42919061a122565b90506115f38160006152c8565b6060601b805480602002602001604051908101604052809291908181526020016000905b82821015610fcf5783829060005260206000209060020201604051806040016040529081600082018054611ea690619e72565b80601f0160208091040260200160405190810160405280929190818152602001828054611ed290619e72565b8015611f1f5780601f10611ef457610100808354040283529160200191611f1f565b820191906000526020600020905b815481529060010190602001808311611f0257829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015611fb957602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411611f665790505b50505050508152505081526020019060010190611e73565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561202f57600080fd5b505af1158015612043573d6000803e3d6000fd5b50506022546040517f6f8b44b0000000000000000000000000000000000000000000000000000000008152600481018590526001600160a01b039091169250636f8b44b09150602401600060405180830381600087803b1580156120a657600080fd5b505af11580156120ba573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561211757600080fd5b505af115801561212b573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fc30436e9000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156121b457600080fd5b505af11580156121c8573d6000803e3d6000fd5b50506022546020546001600160a01b03918216935063106e62909250166121f0846001619ff9565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b039092166004830152602482015260006044820152606401600060405180830381600087803b15801561225557600080fd5b505af1158015612269573d6000803e3d6000fd5b5050505050565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a090600090819060250160408051808303601f19018152908290526024805463ca669fa760e01b84526001600160a01b03166004840152909250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa79101600060405180830381600087803b15801561230d57600080fd5b505af1158015612321573d6000803e3d6000fd5b505060248054604080516001600160a01b03909216928201929092527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250612410919060040161a1b5565b600060405180830381600087803b15801561242a57600080fd5b505af115801561243e573d6000803e3d6000fd5b50506022546020546040517f057e0f250000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063057e0f259350612499929091169087908690889060289060040161a1c8565b600060405180830381600087803b1580156124b357600080fd5b505af11580156124c7573d6000803e3d6000fd5b50505050505050565b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015610fcf57838290600052602060002001805461251390619e72565b80601f016020809104026020016040519081016040528092919081815260200182805461253f90619e72565b801561258c5780601f106125615761010080835404028352916020019161258c565b820191906000526020600020905b81548152906001019060200180831161256f57829003601f168201915b5050505050815260200190600101906124f4565b6060601d805480602002602001604051908101604052809291908181526020016000905b82821015610fcf5760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561268357602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116126305790505b505050505081525050815260200190600101906125c4565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a09060009060250160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561273657600080fd5b505af115801561274a573d6000803e3d6000fd5b50506022546040517f6f8b44b0000000000000000000000000000000000000000000000000000000008152600481018690526001600160a01b039091169250636f8b44b09150602401600060405180830381600087803b1580156127ad57600080fd5b505af11580156127c1573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561281e57600080fd5b505af1158015612832573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fc30436e9000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156128bb57600080fd5b505af11580156128cf573d6000803e3d6000fd5b50506022546020546001600160a01b039182169350635e3e9fef9250166128f7856001619ff9565b846040518463ffffffff1660e01b8152600401610e609392919061a214565b60225460408051606081019091526025808252612960926001600160a01b0316919061d30860208301396040805160208101909152600081526024546001600160a01b0316615347565b6022546023546025546040516370a0823160e01b81526001600160a01b03918216600482015292811692620186a09260009216906370a0823190602401602060405180830381865afa1580156129ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129de919061a122565b90506129eb8160006152c8565b6025546040516001600160a01b0390911660248201526044810183905260006064820181905290819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1e458bee0000000000000000000000000000000000000000000000000000000017905260235490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391612ad4916001600160a01b039190911690600090869060040161a13b565b600060405180830381600087803b158015612aee57600080fd5b505af1158015612b02573d6000803e3d6000fd5b50506040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b0388166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015612b9057600080fd5b505af1158015612ba4573d6000803e3d6000fd5b50506025546040518781526001600160a01b0390911692507f3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9915060200160405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612c4357600080fd5b505af1158015612c57573d6000803e3d6000fd5b50506025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526024810188905260448101869052908816925063106e62909150606401600060405180830381600087803b158015612cc957600080fd5b505af1158015612cdd573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015612d30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d54919061a122565b9050610e8e81866152c8565b6060601c805480602002602001604051908101604052809291908181526020016000905b82821015610fcf5760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015612e4357602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411612df05790505b50505050508152505081526020019060010190612d84565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015610fcf578382906000526020600020018054612e9e90619e72565b80601f0160208091040260200160405190810160405280929190818152602001828054612eca90619e72565b8015612f175780601f10612eec57610100808354040283529160200191612f17565b820191906000526020600020905b815481529060010190602001808311612efa57829003601f168201915b505050505081526020019060010190612e7f565b60085460009060ff1615612f43575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015612fd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ff8919061a122565b1415905090565b602354602554604051620186a0602482018190526001600160a01b039384166044830152929091166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905260248054915163ca669fa760e01b81526001600160a01b039092166004830152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa79101600060405180830381600087803b1580156130e657600080fd5b505af11580156130fa573d6000803e3d6000fd5b505060248054604080516001600160a01b03909216928201929092527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506131e9919060040161a1b5565b600060405180830381600087803b15801561320357600080fd5b505af1158015613217573d6000803e3d6000fd5b50506022546020546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef9350612499929091169087908690889060040161a17c565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090600090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156132d057600080fd5b505af11580156132e4573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506133cf919060040161a1b5565b600060405180830381600087803b1580156133e957600080fd5b505af11580156133fd573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561345157600080fd5b505af1158015613465573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156134c257600080fd5b505af11580156134d6573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506135c1919060040161a1b5565b600060405180830381600087803b1580156135db57600080fd5b505af11580156135ef573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561364357600080fd5b505af1158015613657573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156136b457600080fd5b505af11580156136c8573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561371c57600080fd5b505af1158015613730573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd93c0665000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b1580156137b957600080fd5b505af11580156137cd573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561382a57600080fd5b505af115801561383e573d6000803e3d6000fd5b50506022546025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101879052604481018690529116925063106e62909150606401600060405180830381600087803b1580156138b257600080fd5b505af11580156138c6573d6000803e3d6000fd5b50506024805460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063ca669fa7925001600060405180830381600087803b15801561392257600080fd5b505af1158015613936573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561398a57600080fd5b505af115801561399e573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156139f1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a15919061a122565b9050613a228160006152c8565b6025546040516001600160a01b039091166024820152604481018490526064810183905260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1e458bee0000000000000000000000000000000000000000000000000000000017905260235490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391613b09916001600160a01b039190911690600090869060040161a13b565b600060405180830381600087803b158015613b2357600080fd5b505af1158015613b37573d6000803e3d6000fd5b50506022546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015613bc957600080fd5b505af1158015613bdd573d6000803e3d6000fd5b50506025546040518781526001600160a01b0390911692507f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5915060200160405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613c7c57600080fd5b505af1158015613c90573d6000803e3d6000fd5b50506022546025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101899052604481018890529116925063106e629091506064015b600060405180830381600087803b158015613d0557600080fd5b505af1158015613d19573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015613d6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d90919061a122565b905061226981866152c8565b6023546025546040516370a0823160e01b81526001600160a01b039182166004820152620186a09260009216906370a0823190602401602060405180830381865afa158015613def573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e13919061a122565b9050613e208160006152c8565b6025546040516001600160a01b0390911660248201526044810183905260006064820181905290819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1e458bee0000000000000000000000000000000000000000000000000000000017905260235490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391613f09916001600160a01b039190911690600090869060040161a13b565b600060405180830381600087803b158015613f2357600080fd5b505af1158015613f37573d6000803e3d6000fd5b50506022546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015613fc957600080fd5b505af1158015613fdd573d6000803e3d6000fd5b50506025546040518781526001600160a01b0390911692507f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5915060200160405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561407c57600080fd5b505af1158015614090573d6000803e3d6000fd5b50506022546025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101899052604481018790529116925063106e62909150606401613ceb565b602354602554604051620186a0602482018190526001600160a01b039384166044830152929091166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc51316910000000000000000000000000000000000000000000000000000000017905260235460255491516370a0823160e01b81526001600160a01b0392831660048201529293506000929116906370a0823190602401602060405180830381865afa1580156141cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141f1919061a122565b90506141fe8160006152c8565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa15801561424f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614273919061a122565b90506142808160006152c8565b601f546040516101009091046001600160a01b03166024820152604481018690526064810185905260009060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1e458bee0000000000000000000000000000000000000000000000000000000017905260235490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161436b916001600160a01b039190911690600090869060040161a13b565b600060405180830381600087803b15801561438557600080fd5b505af1158015614399573d6000803e3d6000fd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561442b57600080fd5b505af115801561443f573d6000803e3d6000fd5b5050601f547f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60925061010090046001600160a01b0316905061448260028961a24d565b602354602554604080516001600160a01b03958616815260208101949094529184168383015292909216606082015290519081900360800190a16022546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561454a57600080fd5b505af115801561455e573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d91506145a3908990889061a163565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561460457600080fd5b505af1158015614618573d6000803e3d6000fd5b50506022546020546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef935061467092909116908a9089908b9060040161a17c565b600060405180830381600087803b15801561468a57600080fd5b505af115801561469e573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156146f1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614715919061a122565b90506115f38161472660028a61a24d565b6152c8565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a090600090819060250160408051808303601f19018152908290526023546020546370a0823160e01b84526001600160a01b0390811660048501529193506000929116906370a0823190602401602060405180830381865afa1580156147c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906147e5919061a122565b90506147f28160006152c8565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015614843573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614867919061a122565b601f54604080516001600160a01b036101009093048316602482015260448101899052606480820189905282518083039091018152608490910182526020810180517f1e458bee000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba39261494992911690600090869060040161a13b565b600060405180830381600087803b15801561496357600080fd5b505af1158015614977573d6000803e3d6000fd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015614a0957600080fd5b505af1158015614a1d573d6000803e3d6000fd5b505050507f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e601f60019054906101000a90046001600160a01b03166028604051614a6892919061a288565b60405180910390a1601f546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015614b0257600080fd5b505af1158015614b16573d6000803e3d6000fd5b50506023546020546040516001600160a01b039283169450911691507f1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b143690614b64908a90899060289061a2aa565b60405180910390a36022546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015614bfa57600080fd5b505af1158015614c0e573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c9150614c56908990889060289061a2aa565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614cb757600080fd5b505af1158015614ccb573d6000803e3d6000fd5b50506022546020546040517f057e0f250000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063057e0f259350614d2692909116908a9089908b9060289060040161a1c8565b600060405180830381600087803b158015614d4057600080fd5b505af1158015614d54573d6000803e3d6000fd5b50506023546020546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015614da7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614dcb919061a122565b9050614dd781886152c8565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015614e28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614e4c919061a122565b905061167581856152c8565b60606015805480602002602001604051908101604052809291908181526020018280548015610bd8576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610bba575050505050905090565b6024805460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa79101600060405180830381600087803b158015614f1057600080fd5b505af1158015614f24573d6000803e3d6000fd5b505060248054604080516001600160a01b03909216928201929092527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250615013919060040161a1b5565b600060405180830381600087803b15801561502d57600080fd5b505af1158015615041573d6000803e3d6000fd5b50506022546040517f6f8b44b000000000000000000000000000000000000000000000000000000000815261271060048201526001600160a01b039091169250636f8b44b09150602401600060405180830381600087803b1580156150a557600080fd5b505af1158015610b7a573d6000803e3d6000fd5b6024805460405163ca669fa760e01b81526001600160a01b039091166004820152620186a091600091737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa79101600060405180830381600087803b15801561511957600080fd5b505af115801561512d573d6000803e3d6000fd5b505060248054604080516001600160a01b03909216928201929092527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061521c919060040161a1b5565b600060405180830381600087803b15801561523657600080fd5b505af115801561524a573d6000803e3d6000fd5b50506022546025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101879052604481018690529116925063106e62909150606401610e60565b60006152b36199dc565b6152be84848361535c565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c549060440160006040518083038186803b15801561533357600080fd5b505afa158015610e8e573d6000803e3d6000fd5b61534f6199dc565b61226985858584866153d7565b60008061536985846154d7565b90506153cc6040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016153b792919061a2d5565b604051602081830303815290604052856154e3565b9150505b9392505050565b6040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201528190737109709ecfa91a80626ff3989d68f67f5b1dd12d9081906306447d5690602401600060405180830381600087803b15801561544957600080fd5b505af192505050801561545a575060015b61546f5761546a87878787615511565b6124c7565b61547b87878787615511565b806001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156154b657600080fd5b505af11580156154ca573d6000803e3d6000fd5b5050505050505050505050565b60006153d0838361552a565b60c081015151600090156155075761550084848460c00151615545565b90506153d0565b61550084846156eb565b600061551d84836157d6565b90506122698582856157e2565b60006155368383615bac565b6153d0838360200151846154e3565b600080615550615bbc565b9050600061555e8683615c8f565b905060006155758260600151836020015185616135565b9050600061558583838989616347565b90506000615592826171c4565b602081015181519192509060030b15615605578982604001516040516020016155bc92919061a2f7565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526155fc9160040161a1b5565b60405180910390fd5b60006156486040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001617393565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d9061569b90849060040161a1b5565b602060405180830381865afa1580156156b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906156dc919061a378565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc9259061574090879060040161a1b5565b600060405180830381865afa15801561575d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615785919081019061a45a565b905060006157b3828560405160200161579f92919061a48f565b604051602081830303815290604052617593565b90506001600160a01b0381166152be5784846040516020016155bc92919061a4be565b600061553683836175a6565b6040517f667f9d700000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90600090829063667f9d7090604401602060405180830381865afa15801561587e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906158a2919061a122565b905080615a495760006158b4866175b2565b604080518082018252600581527f352e302e300000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061593f905b604080518082018252600080825260209182015281518083019092528451825280850190820152906176aa565b8061594b575060008451115b156159ce576040517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b03871690634f1ef28690615997908890889060040161a2d5565b600060405180830381600087803b1580156159b157600080fd5b505af11580156159c5573d6000803e3d6000fd5b50505050615a43565b6040517f3659cfe60000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152871690633659cfe690602401600060405180830381600087803b158015615a2a57600080fd5b505af1158015615a3e573d6000803e3d6000fd5b505050505b50612269565b806000615a55826175b2565b604080518082018252600581527f352e302e3000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150615ab790615912565b80615ac3575060008551115b15615b48576040517f9623609d0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690639623609d90615b11908a908a908a9060040161a569565b600060405180830381600087803b158015615b2b57600080fd5b505af1158015615b3f573d6000803e3d6000fd5b505050506124c7565b6040517f99a88ec40000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015287811660248301528316906399a88ec490604401600060405180830381600087803b1580156154b657600080fd5b615bb8828260006176be565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c90615c4390849060040161a59a565b600060405180830381865afa158015615c60573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615c88919081019061a5e1565b9250505090565b615cc16040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d9050615d0c6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b615d15856177c1565b60208201526000615d2586617ba6565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa158015615d67573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615d8f919081019061a5e1565b86838560200151604051602001615da9949392919061a62a565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb1190615e0190859060040161a1b5565b600060405180830381865afa158015615e1e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615e46919081019061a5e1565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f690615e8e90849060040161a72e565b602060405180830381865afa158015615eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615ecf919061a780565b615ee457816040516020016155bc919061a7a2565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890615f2990849060040161a834565b600060405180830381865afa158015615f46573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615f6e919081019061a5e1565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f690615fb590849060040161a886565b602060405180830381865afa158015615fd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615ff6919061a780565b1561608b576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061604090849060040161a886565b600060405180830381865afa15801561605d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616085919081019061a5e1565b60408501525b846001600160a01b03166349c4fac88286600001516040516020016160b0919061a8d8565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016160dc92919061a944565b600060405180830381865afa1580156160f9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616121919081019061a5e1565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b60608152602001906001900390816161515790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250816000815181106161b1576161b161a969565b60200260200101819052506040518060400160405280600381526020017f2d726c0000000000000000000000000000000000000000000000000000000000815250816001815181106162055761620561a969565b602002602001018190525084604051602001616221919061a998565b604051602081830303815290604052816002815181106162435761624361a969565b60200260200101819052508260405160200161625f919061aa04565b604051602081830303815290604052816003815181106162815761628161a969565b60200260200101819052506000616297826171c4565b602080820151604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000081850190815282518084018452600080825290860152825180840190935290518252928101929092529192506163289060408051808201825260008082526020918201528151808301909252845182528085019082015290617e29565b61633d57856040516020016155bc919061aa45565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015616397565b511590565b61650b57826020015115616453576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a4016155fc565b8260c001511561650b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a4016155fc565b6040805160ff8082526120008201909252600091816020015b606081526020019060019003908161652457905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061657f9061aad6565b935060ff16815181106165945761659461a969565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e37000000000000000000000000000000000000008152506040516020016165e5919061aaf5565b6040516020818303038152906040528282806166009061aad6565b935060ff16815181106166155761661561a969565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806166629061aad6565b935060ff16815181106166775761667761a969565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806166c49061aad6565b935060ff16815181106166d9576166d961a969565b602002602001018190525087602001518282806166f59061aad6565b935060ff168151811061670a5761670a61a969565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e7472616374506174680000000000000000000000000000000000008152508282806167579061aad6565b935060ff168151811061676c5761676c61a969565b6020908102919091010152875182826167848161aad6565b935060ff16815181106167995761679961a969565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e496400000000000000000000000000000000000000000000008152508282806167e69061aad6565b935060ff16815181106167fb576167fb61a969565b602002602001018190525061680f46617e8a565b828261681a8161aad6565b935060ff168151811061682f5761682f61a969565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c65000000000000000000000000000000000081525082828061687c9061aad6565b935060ff16815181106168915761689161a969565b6020026020010181905250868282806168a99061aad6565b935060ff16815181106168be576168be61a969565b60209081029190910101528551156169e55760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f646500000000000000000000006020820152828261690f8161aad6565b935060ff16815181106169245761692461a969565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d9061697490899060040161a1b5565b600060405180830381865afa158015616991573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526169b9919081019061a5e1565b82826169c48161aad6565b935060ff16815181106169d9576169d961a969565b60200260200101819052505b846020015115616ab55760408051808201909152601281527f2d2d766572696679536f75726365436f6465000000000000000000000000000060208201528282616a2e8161aad6565b935060ff1681518110616a4357616a4361a969565b60200260200101819052506040518060400160405280600581526020017f66616c7365000000000000000000000000000000000000000000000000000000815250828280616a909061aad6565b935060ff1681518110616aa557616aa561a969565b6020026020010181905250616c7c565b616aed6163928660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b616b805760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282616b308161aad6565b935060ff1681518110616b4557616b4561a969565b60200260200101819052508460a00151604051602001616b65919061a998565b604051602081830303815290604052828280616a909061aad6565b8460c00151158015616bc3575060408089015181518083018352600080825260209182015282518084019093528151835290810190820152616bc190511590565b155b15616c7c5760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282616c078161aad6565b935060ff1681518110616c1c57616c1c61a969565b6020026020010181905250616c3088617f2a565b604051602001616c40919061a998565b604051602081830303815290604052828280616c5b9061aad6565b935060ff1681518110616c7057616c7061a969565b60200260200101819052505b60408086015181518083018352600080825260209182015282518084019093528151835290810190820152616cb090511590565b616d455760408051808201909152600b81527f2d2d72656c61796572496400000000000000000000000000000000000000000060208201528282616cf38161aad6565b935060ff1681518110616d0857616d0861a969565b60200260200101819052508460400151828280616d249061aad6565b935060ff1681518110616d3957616d3961a969565b60200260200101819052505b606085015115616e665760408051808201909152600681527f2d2d73616c74000000000000000000000000000000000000000000000000000060208201528282616d8e8161aad6565b935060ff1681518110616da357616da361a969565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015616e12573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616e3a919081019061a5e1565b8282616e458161aad6565b935060ff1681518110616e5a57616e5a61a969565b60200260200101819052505b60e08501515115616f0d5760408051808201909152600a81527f2d2d6761734c696d69740000000000000000000000000000000000000000000060208201528282616eb08161aad6565b935060ff1681518110616ec557616ec561a969565b6020026020010181905250616ee18560e0015160000151617e8a565b8282616eec8161aad6565b935060ff1681518110616f0157616f0161a969565b60200260200101819052505b60e08501516020015115616fb75760408051808201909152600a81527f2d2d67617350726963650000000000000000000000000000000000000000000060208201528282616f5a8161aad6565b935060ff1681518110616f6f57616f6f61a969565b6020026020010181905250616f8b8560e0015160200151617e8a565b8282616f968161aad6565b935060ff1681518110616fab57616fab61a969565b60200260200101819052505b60e085015160400151156170615760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826170048161aad6565b935060ff16815181106170195761701961a969565b60200260200101819052506170358560e0015160400151617e8a565b82826170408161aad6565b935060ff16815181106170555761705561a969565b60200260200101819052505b60e0850151606001511561710b5760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826170ae8161aad6565b935060ff16815181106170c3576170c361a969565b60200260200101819052506170df8560e0015160600151617e8a565b82826170ea8161aad6565b935060ff16815181106170ff576170ff61a969565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561712957617129619e43565b60405190808252806020026020018201604052801561715c57816020015b60608152602001906001900390816171475790505b50905060005b8260ff168160ff1610156171b557838160ff16815181106171855761718561a969565b6020026020010151828260ff16815181106171a2576171a261a969565b6020908102919091010152600101617162565b5093505050505b949350505050565b6171eb6040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c916172719186910161ab60565b600060405180830381865afa15801561728e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526172b6919081019061a5e1565b905060006172c48683618a19565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b81526004016172f49190619d35565b6000604051808303816000875af1158015617313573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261733b919081019061aba7565b805190915060030b158015906173545750602081015151155b80156173635750604081015151155b1561633d578160008151811061737b5761737b61a969565b60200260200101516040516020016155bc919061ac5d565b606060006173c88560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925286518252808701908201529091506173ff9082905b90618b6e565b1561755c57600061747c82617476846174706174428a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b90618b95565b90618bf7565b604080518082018252600181527f0a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506174e0908290618b6e565b1561754a57604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617547905b8290618c7c565b90505b61755381618ca2565b925050506153d0565b82156175755784846040516020016155bc92919061ae49565b50506040805160208101909152600081526153d0565b509392505050565b6000808251602084016000f09392505050565b615bb8828260016176be565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fad3cb1cc00000000000000000000000000000000000000000000000000000000179052905160609160009182916001600160a01b03861691617627919061aef0565b600060405180830381855afa9150503d8060008114617662576040519150601f19603f3d011682016040523d82523d6000602084013e617667565b606091505b509150915081801561767a575060208151115b1561769357808060200190518101906171bc919061a5e1565b505060408051602081019091526000815292915050565b60006176b68383618d0b565b159392505050565b8160a00151156176cd57505050565b60006176da848484618de6565b905060006176e7826171c4565b602081015181519192509060030b1580156177835750604080518082018252600781527f535543434553530000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617783906040805180820182526000808252602091820152815180830190925284518252808501908201526173f9565b1561779057505050505050565b604082015151156177b05781604001516040516020016155bc919061af0c565b806040516020016155bc919061af6a565b606060006177f68360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061785b905b8290617e29565b156178ca57604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526153d0906178c5908390619381565b618ca2565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261792c905b829061940b565b6001036179f957604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261799290617540565b50604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526153d0906178c5905b8390618c7c565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617a5890617854565b15617b8f57604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290617ac09083906194a5565b905060008160018351617ad3919061afd5565b81518110617ae357617ae361a969565b60200260200101519050617b866178c5617b596040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252855182528086019082015290619381565b95945050505050565b826040516020016155bc919061afe8565b50919050565b60606000617bdb8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150617c3d90617854565b15617c4b576153d081618ca2565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617caa90617925565b600103617d1457604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526153d0906178c5906179f2565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617d7390617854565b15617b8f57604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290617ddb9083906194a5565b9050600181511115617e17578060028251617df6919061afd5565b81518110617e0657617e0661a969565b602002602001015192505050919050565b50826040516020016155bc919061afe8565b805182516000911115617e3e575060006152c2565b81518351602085015160009291617e5491619ff9565b617e5e919061afd5565b905082602001518103617e755760019150506152c2565b82516020840151819020912014905092915050565b60606000617e978361954a565b600101905060008167ffffffffffffffff811115617eb757617eb7619e43565b6040519080825280601f01601f191660200182016040528015617ee1576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084617eeb57509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e5345440000000000000000000000000000000000000000000081840190815285518087018752838152840192909252845180860190955251845290830152606091617fb6905b82906176aa565b15617ff657505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261805590617faf565b1561809557505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d49540000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526180f490617faf565b1561813457505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261819390617faf565b806181f85750604080518082018252601081527f47504c2d322e302d6f722d6c6174657200000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526181f890617faf565b1561823857505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261829790617faf565b806182fc5750604080518082018252601081527f47504c2d332e302d6f722d6c6174657200000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526182fc90617faf565b1561833c57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261839b90617faf565b806184005750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261840090617faf565b1561844057505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261849f90617faf565b806185045750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261850490617faf565b1561854457505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c617573650000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526185a390617faf565b156185e357505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261864290617faf565b1561868257505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e3000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526186e190617faf565b1561872157505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261878090617faf565b156187c057505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261881f90617faf565b1561885f57505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526188be90617faf565b806189235750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261892390617faf565b1561896357505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e31000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526189c290617faf565b15618a0257505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516155bc929060200161b0c6565b60608060005b8451811015618aa45781858281518110618a3b57618a3b61a969565b6020026020010151604051602001618a5492919061a48f565b604051602081830303815290604052915060018551618a73919061afd5565b8114618a9c5781604051602001618a8a919061b22f565b60405160208183030381529060405291505b600101618a1f565b5060408051600380825260808201909252600091816020015b6060815260200190600190039081618abd5790505090508381600081518110618ae857618ae861a969565b60200260200101819052506040518060400160405280600281526020017f2d6300000000000000000000000000000000000000000000000000000000000081525081600181518110618b3c57618b3c61a969565b60200260200101819052508181600281518110618b5b57618b5b61a969565b6020908102919091010152949350505050565b6020808301518351835192840151600093618b8c929184919061962c565b14159392505050565b60408051808201909152600080825260208201526000618bc7846000015185602001518560000151866020015161973d565b9050836020015181618bd9919061afd5565b84518590618be890839061afd5565b90525060208401525090919050565b6040805180820190915260008082526020820152815183511015618c1c5750816152c2565b6020808301519084015160019114618c435750815160208481015190840151829020919020145b8015618c7457825184518590618c5a90839061afd5565b9052508251602085018051618c70908390619ff9565b9052505b509192915050565b6040805180820190915260008082526020820152618c9b83838361985d565b5092915050565b60606000826000015167ffffffffffffffff811115618cc357618cc3619e43565b6040519080825280601f01601f191660200182016040528015618ced576020820181803683370190505b5090506000602082019050618c9b8185602001518660000151619908565b8151815160009190811115618d1e575081515b6020808501519084015160005b83811015618dd75782518251808214618da7576000196020871015618d8657600184618d5889602061afd5565b618d629190619ff9565b618d6d90600861b270565b618d7890600261b36e565b618d82919061afd5565b1990505b8181168382168181039114618da45797506152c29650505050505050565b50505b618db2602086619ff9565b9450618dbf602085619ff9565b93505050602081618dd09190619ff9565b9050618d2b565b508451865161633d919061b37a565b60606000618df2615bbc565b6040805160ff808252612000820190925291925060009190816020015b6060815260200190600190039081618e0f57905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280618e6a9061aad6565b935060ff1681518110618e7f57618e7f61a969565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e3300000000000000000000000000000000000000000000000000815250604051602001618ed0919061b39a565b604051602081830303815290604052828280618eeb9061aad6565b935060ff1681518110618f0057618f0061a969565b60200260200101819052506040518060400160405280600881526020017f76616c6964617465000000000000000000000000000000000000000000000000815250828280618f4d9061aad6565b935060ff1681518110618f6257618f6261a969565b602002602001018190525082604051602001618f7e919061aa04565b604051602081830303815290604052828280618f999061aad6565b935060ff1681518110618fae57618fae61a969565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e747261637400000000000000000000000000000000000000000000815250828280618ffb9061aad6565b935060ff16815181106190105761901061a969565b60200260200101819052506190258784619982565b82826190308161aad6565b935060ff16815181106190455761904561a969565b6020908102919091010152855151156190f15760408051808201909152600b81527f2d2d7265666572656e6365000000000000000000000000000000000000000000602082015282826190978161aad6565b935060ff16815181106190ac576190ac61a969565b60200260200101819052506190c5866000015184619982565b82826190d08161aad6565b935060ff16815181106190e5576190e561a969565b60200260200101819052505b85608001511561915f5760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261913a8161aad6565b935060ff168151811061914f5761914f61a969565b60200260200101819052506191c5565b84156191c55760408051808201909152601281527f2d2d726571756972655265666572656e63650000000000000000000000000000602082015282826191a48161aad6565b935060ff16815181106191b9576191b961a969565b60200260200101819052505b604086015151156192615760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261920f8161aad6565b935060ff16815181106192245761922461a969565b602002602001018190525085604001518282806192409061aad6565b935060ff16815181106192555761925561a969565b60200260200101819052505b8560600151156192cb5760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826192aa8161aad6565b935060ff16815181106192bf576192bf61a969565b60200260200101819052505b60008160ff1667ffffffffffffffff8111156192e9576192e9619e43565b60405190808252806020026020018201604052801561931c57816020015b60608152602001906001900390816193075790505b50905060005b8260ff168160ff16101561937557838160ff16815181106193455761934561a969565b6020026020010151828260ff16815181106193625761936261a969565b6020908102919091010152600101619322565b50979650505050505050565b60408051808201909152600080825260208201528151835110156193a65750816152c2565b815183516020850151600092916193bc91619ff9565b6193c6919061afd5565b602084015190915060019082146193e7575082516020840151819020908220145b8015619402578351855186906193fe90839061afd5565b9052505b50929392505050565b600080826000015161942f856000015186602001518660000151876020015161973d565b6194399190619ff9565b90505b8351602085015161944d9190619ff9565b8111618c9b578161945d8161b3df565b9250508260000151619494856020015183619478919061afd5565b8651619484919061afd5565b838660000151876020015161973d565b61949e9190619ff9565b905061943c565b606060006194b3848461940b565b6194be906001619ff9565b67ffffffffffffffff8111156194d6576194d6619e43565b60405190808252806020026020018201604052801561950957816020015b60608152602001906001900390816194f45790505b50905060005b815181101561758b576195256178c58686618c7c565b8282815181106195375761953761a969565b602090810291909101015260010161950f565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310619593577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106195bf576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106195dd57662386f26fc10000830492506010015b6305f5e10083106195f5576305f5e100830492506008015b612710831061960957612710830492506004015b6064831061961b576064830492506002015b600a83106152c25760010192915050565b60008085841161973357602084116196df576000841561967757600161965386602061afd5565b61965e90600861b270565b61966990600261b36e565b619673919061afd5565b1990505b83518116856196868989619ff9565b619690919061afd5565b805190935082165b8181146196ca578784116196b257879450505050506171bc565b836196bc8161b3f9565b945050828451169050619698565b6196d48785619ff9565b9450505050506171bc565b8383206196ec858861afd5565b6196f69087619ff9565b91505b8582106197315784822080820361971e576197148684619ff9565b93505050506171bc565b61972960018461afd5565b9250506196f9565b505b5092949350505050565b6000838186851161984857602085116197f7576000851561978957600161976587602061afd5565b61977090600861b270565b61977b90600261b36e565b619785919061afd5565b1990505b8451811660008761979a8b8b619ff9565b6197a4919061afd5565b855190915083165b8281146197e9578186106197d1576197c48b8b619ff9565b96505050505050506171bc565b856197db8161b3df565b9650508386511690506197ac565b8596505050505050506171bc565b508383206000905b619809868961afd5565b82116198465785832080820361982557839450505050506171bc565b619830600185619ff9565b935050818061983e9061b3df565b9250506197ff565b505b6198528787619ff9565b979650505050505050565b6040805180820190915260008082526020820152600061988f856000015186602001518660000151876020015161973d565b6020808701805191860191909152519091506198ab908261afd5565b8352845160208601516198be9190619ff9565b81036198cd57600085526198ff565b835183516198db9190619ff9565b855186906198ea90839061afd5565b90525083516198f99082619ff9565b60208601525b50909392505050565b60208110619940578151835261991f602084619ff9565b925061992c602083619ff9565b915061993960208261afd5565b9050619908565b600019811561996f57600161995683602061afd5565b6199629061010061b36e565b61996c919061afd5565b90505b9151835183169219169190911790915250565b606060006199908484615c8f565b80516020808301516040519394506199aa9390910161b410565b60405160208183030381529060405291505092915050565b6112a68061b46983390190565b610bf98061c70f83390190565b6040518060e00160405280606081526020016060815260200160608152602001600015158152602001600015158152602001600015158152602001619a1f619a24565b905290565b60405180610100016040528060001515815260200160001515815260200160608152602001600080191681526020016060815260200160608152602001600015158152602001619a1f6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b81811015619ad65783516001600160a01b0316835260209384019390920191600101619aaf565b509095945050505050565b60005b83811015619afc578181015183820152602001619ae4565b50506000910152565b60008151808452619b1d816020860160208601619ae1565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015619c2d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b81811015619c13577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8503018352619bfd848651619b05565b6020958601959094509290920191600101619bc3565b509197505050602094850194929092019150600101619b59565b50929695505050505050565b600081518084526020840193506020830160005b82811015619c8d5781517fffffffff0000000000000000000000000000000000000000000000000000000016865260209586019590910190600101619c4d565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015619c2d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160408752619d036040880182619b05565b9050602082015191508681036020880152619d1e8183619c39565b965050506020938401939190910190600101619cbf565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015619c2d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452619d97858351619b05565b94506020938401939190910190600101619d5d565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015619c2d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b0381511686526020810151905060406020870152619e2d6040870182619c39565b9550506020938401939190910190600101619dd4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c90821680619e8657607f821691505b602082108103617ba0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f821115619f0657806000526020600020601f840160051c81016020851015619ee65750805b601f840160051c820191505b818110156122695760008155600101619ef2565b505050565b815167ffffffffffffffff811115619f2557619f25619e43565b619f3981619f338454619e72565b84619ebf565b6020601f821160018114619f6d5760008315619f555750848201515b600019600385901b1c1916600184901b178455612269565b600084815260208120601f198516915b82811015619f9d5787850151825560209485019460019092019101619f7d565b5084821015619fbb5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156152c2576152c2619fca565b600081546001600160a01b038116845267ffffffffffffffff8160a01c1660208501525060018201606060408501526000815461a04881619e72565b806060880152600182166000811461a067576001811461a0a15761a0d5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083166080890152608082151560051b890101935061a0d5565b84600052602060002060005b8381101561a0cc5781548a82016080015260019091019060200161a0ad565b89016080019450505b50919695505050505050565b6001600160a01b038516815283602082015260a06040820152600061a10960a0830185619b05565b600060608401528281036080840152619852818561a00c565b60006020828403121561a13457600080fd5b5051919050565b6001600160a01b0384168152826020820152606060408201526000617b866060830184619b05565b8281526040602082015260006171bc6040830184619b05565b6001600160a01b038516815283602082015260806040820152600061a1a46080830185619b05565b905082606083015295945050505050565b6020815260006153d06020830184619b05565b6001600160a01b038616815284602082015260a06040820152600061a1f060a0830186619b05565b846060840152828103608084015261a208818561a00c565b98975050505050505050565b6001600160a01b038416815282602082015260806040820152600061a23c6080830184619b05565b905060006060830152949350505050565b60008261a283577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6001600160a01b03831681526040602082015260006171bc604083018461a00c565b83815260606020820152600061a2c36060830185619b05565b828103604084015261633d818561a00c565b6001600160a01b03831681526040602082015260006171bc6040830184619b05565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161a32f81601a850160208801619ae1565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161a36c81601c840160208801619ae1565b01601c01949350505050565b60006020828403121561a38a57600080fd5b81516001600160a01b03811681146153d057600080fd5b6040516060810167ffffffffffffffff8111828210171561a3c45761a3c4619e43565b60405290565b60008067ffffffffffffffff84111561a3e55761a3e5619e43565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561a4145761a414619e43565b60405283815290508082840185101561a42c57600080fd5b61758b846020830185619ae1565b600082601f83011261a44b57600080fd5b6153d08383516020850161a3ca565b60006020828403121561a46c57600080fd5b815167ffffffffffffffff81111561a48357600080fd5b6152be8482850161a43a565b6000835161a4a1818460208801619ae1565b83519083019061a4b5818360208801619ae1565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161a4f681601a850160208801619ae1565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161a533816033840160208801619ae1565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b6001600160a01b03841681526001600160a01b0383166020820152606060408201526000617b866060830184619b05565b60408152600b60408201527f464f554e4452595f4f555400000000000000000000000000000000000000000060608201526080602082015260006153d06080830184619b05565b60006020828403121561a5f357600080fd5b815167ffffffffffffffff81111561a60a57600080fd5b8201601f8101841361a61b57600080fd5b6152be8482516020840161a3ca565b6000855161a63c818460208a01619ae1565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161a676816001840160208a01619ae1565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161a6b4816002840160208901619ae1565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161a6f6816002840160208801619ae1565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061a7416040830184619b05565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b60006020828403121561a79257600080fd5b815180151581146153d057600080fd5b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161a7da81601f850160208701619ae1565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061a8476040830184619b05565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061a8996040830184619b05565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161a910816014850160208701619ae1565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061a9576040830185619b05565b82810360208401526153cc8185619b05565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161a9d0816001850160208701619ae1565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161aa16818460208701619ae1565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161aac981604b850160208701619ae1565b91909101604b0192915050565b600060ff821660ff810361aaec5761aaec619fca565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161ab53816029850160208701619ae1565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f504154480000000000000000000060608201526080602082015260006153d06080830184619b05565b60006020828403121561abb957600080fd5b815167ffffffffffffffff81111561abd057600080fd5b82016060818503121561abe257600080fd5b61abea61a3a1565b81518060030b811461abfb57600080fd5b8152602082015167ffffffffffffffff81111561ac1757600080fd5b61ac238682850161a43a565b602083015250604082015167ffffffffffffffff81111561ac4357600080fd5b61ac4f8682850161a43a565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161acbb816021850160208701619ae1565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161aea7816021850160208801619ae1565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161aee481602e840160208801619ae1565b01602e01949350505050565b6000825161af02818460208701619ae1565b9190910192915050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161ab53816029850160208701619ae1565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161afc8816022850160208701619ae1565b9190910160220192915050565b818103818111156152c2576152c2619fca565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161b02081600e850160208701619ae1565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161b0fe816018850160208801619ae1565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161b13b81601c840160208801619ae1565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161b241818460208701619ae1565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b80820281158282048414176152c2576152c2619fca565b6001815b600184111561b2c25780850481111561b2a65761b2a6619fca565b600184161561b2b457908102905b60019390931c92800261b28b565b935093915050565b60008261b2d9575060016152c2565b8161b2e6575060006152c2565b816001811461b2fc576002811461b3065761b322565b60019150506152c2565b60ff84111561b3175761b317619fca565b50506001821b6152c2565b5060208310610133831016604e8410600b841016171561b345575081810a6152c2565b61b352600019848461b287565b806000190482111561b3665761b366619fca565b029392505050565b60006153d0838361b2ca565b8181036000831280158383131683831282161715618c9b57618c9b619fca565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161b3d281601c850160208701619ae1565b91909101601c0192915050565b6000600019820361b3f25761b3f2619fca565b5060010190565b60008161b4085761b408619fca565b506000190190565b6000835161b422818460208801619ae1565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161b45c816001840160208801619ae1565b0160010194935050505056fe608060405234801561001057600080fd5b506040516112a63803806112a683398101604081905261002f91610110565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007891906101e2565b50600461008582826101e2565b5050506001600160a01b03821615806100a557506001600160a01b038116155b156100c35760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b031991821617909155600780549290931691161790556102a0565b80516001600160a01b038116811461010b57600080fd5b919050565b6000806040838503121561012357600080fd5b61012c836100f4565b915061013a602084016100f4565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061016d57607f821691505b60208210810361018d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101dd57806000526020600020601f840160051c810160208510156101ba5750805b601f840160051c820191505b818110156101da57600081556001016101c6565b50505b505050565b81516001600160401b038111156101fb576101fb610143565b61020f816102098454610159565b84610193565b6020601f821160018114610243576000831561022b5750848201515b600019600385901b1c1916600184901b1784556101da565b600084815260208120601f198516915b828110156102735787850151825560209485019460019092019101610253565b50848210156102915786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610ff7806102af6000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806342966c68116100b257806379cc679011610081578063a9059cbb11610066578063a9059cbb1461028e578063bff9662a146102a1578063dd62ed3e146102c157600080fd5b806379cc67901461027357806395d89b411461028657600080fd5b806342966c68146102025780635b1125911461021557806370a0823114610235578063779e3b631461026b57600080fd5b80631e458bee116100ee5780631e458bee1461018857806323b872dd1461019b578063313ce567146101ae578063328a01d0146101bd57600080fd5b806306fdde0314610120578063095ea7b31461013e57806315d57fd41461016157806318160ddd14610176575b600080fd5b610128610307565b6040516101359190610d97565b60405180910390f35b61015161014c366004610e2c565b610399565b6040519015158152602001610135565b61017461016f366004610e56565b6103b3565b005b6002545b604051908152602001610135565b610174610196366004610e89565b61057e565b6101516101a9366004610ebc565b610631565b60405160128152602001610135565b6007546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610135565b610174610210366004610ef9565b610655565b6006546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a610243366004610f12565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b610174610662565b610174610281366004610e2c565b610786565b610128610837565b61015161029c366004610e2c565b610846565b6005546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a6102cf366004610e56565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60606003805461031690610f34565b80601f016020809104026020016040519081016040528092919081815260200182805461034290610f34565b801561038f5780601f106103645761010080835404028352916020019161038f565b820191906000526020600020905b81548152906001019060200180831161037257829003601f168201915b5050505050905090565b6000336103a7818585610854565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103f3575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610431576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82161580610468575073ffffffffffffffffffffffffffffffffffffffff8116155b1561049f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105d1576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6105db8383610866565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161062491815260200190565b60405180910390a3505050565b60003361063f8582856108c6565b61064a858585610995565b506001949350505050565b61065f3382610a40565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106b5576040517fe700765e000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b60065473ffffffffffffffffffffffffffffffffffffffff16610704576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107d9576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6107e38282610a9c565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161082b91815260200190565b60405180910390a25050565b60606004805461031690610f34565b6000336103a7818585610995565b6108618383836001610ab1565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108b6576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c260008383610bf9565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461098f5781811015610980576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610428565b61098f84848484036000610ab1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109e5576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8216610a35576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b610861838383610bf9565b73ffffffffffffffffffffffffffffffffffffffff8216610a90576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c282600083610bf9565b610aa78233836108c6565b6108c28282610a40565b73ffffffffffffffffffffffffffffffffffffffff8416610b01576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8316610b51576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600160209081526040808320938716835292905220829055801561098f578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610beb91815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c31578060026000828254610c269190610f87565b90915550610ce39050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610cb7576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610428565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610d0c57600280548290039055610d38565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161062491815260200190565b602081526000825180602084015260005b81811015610dc55760208186018101516040868401015201610da8565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610e2757600080fd5b919050565b60008060408385031215610e3f57600080fd5b610e4883610e03565b946020939093013593505050565b60008060408385031215610e6957600080fd5b610e7283610e03565b9150610e8060208401610e03565b90509250929050565b600080600060608486031215610e9e57600080fd5b610ea784610e03565b95602085013595506040909401359392505050565b600080600060608486031215610ed157600080fd5b610eda84610e03565b9250610ee860208501610e03565b929592945050506040919091013590565b600060208284031215610f0b57600080fd5b5035919050565b600060208284031215610f2457600080fd5b610f2d82610e03565b9392505050565b600181811c90821680610f4857607f821691505b602082108103610f81577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b808201808211156103ad577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220c088d439a8fedc78e9a5e8cc1447108cf5ecd03d1ba62702320339d3442c0aca64736f6c634300081a00336080604052348015600f57600080fd5b506001600055610bd5806100246000396000f3fe6080604052600436106100635760003560e01c8063c513169111610040578063c5131691146100c1578063e04d4f97146100e1578063f05b6abf146100f457005b8063357fc5a21461006c578063660b9de01461008c5780636ed70169146100ac57005b3661006a57005b005b34801561007857600080fd5b5061006a61008736600461051e565b610114565b34801561009857600080fd5b5061006a6100a736600461055a565b6101aa565b3480156100b857600080fd5b5061006a6101e6565b3480156100cd57600080fd5b5061006a6100dc36600461051e565b61021b565b61006a6100ef3660046106b8565b6102f6565b34801561010057600080fd5b5061006a61010f366004610797565b61033a565b61011c61036f565b61013e73ffffffffffffffffffffffffffffffffffffffff83163383866103b2565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101a56001600055565b505050565b7f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e33826040516101db9291906108ca565b60405180910390a150565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61022361036f565b60006102306002856109b3565b90508060000361026c576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61028e73ffffffffffffffffffffffffffffffffffffffff84163384846103b2565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101a56001600055565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa333485858560405161032d959493929190610a52565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b1463384848460405161032d9493929190610adc565b6002600054036103ab576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261044790859061044d565b50505050565b600080602060008451602086016000885af180610470576040513d6000823e3d81fd5b50506000513d915081156104885780600114156104a2565b73ffffffffffffffffffffffffffffffffffffffff84163b155b15610447576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240160405180910390fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461051957600080fd5b919050565b60008060006060848603121561053357600080fd5b83359250610543602085016104f5565b9150610551604085016104f5565b90509250925092565b60006020828403121561056c57600080fd5b813567ffffffffffffffff81111561058357600080fd5b82016060818503121561059557600080fd5b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156106125761061261059c565b604052919050565b600082601f83011261062b57600080fd5b813567ffffffffffffffff8111156106455761064561059c565b61067660207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016105cb565b81815284602083860101111561068b57600080fd5b816020850160208301376000918101602001919091529392505050565b8035801515811461051957600080fd5b6000806000606084860312156106cd57600080fd5b833567ffffffffffffffff8111156106e457600080fd5b6106f08682870161061a565b93505060208401359150610551604085016106a8565b600067ffffffffffffffff8211156107205761072061059c565b5060051b60200190565b600082601f83011261073b57600080fd5b813561074e61074982610706565b6105cb565b8082825260208201915060208360051b86010192508583111561077057600080fd5b602085015b8381101561078d578035835260209283019201610775565b5095945050505050565b6000806000606084860312156107ac57600080fd5b833567ffffffffffffffff8111156107c357600080fd5b8401601f810186136107d457600080fd5b80356107e261074982610706565b8082825260208201915060208360051b85010192508883111561080457600080fd5b602084015b8381101561084657803567ffffffffffffffff81111561082857600080fd5b6108378b60208389010161061a565b84525060209283019201610809565b509550505050602084013567ffffffffffffffff81111561086657600080fd5b6108728682870161072a565b925050610551604085016106a8565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610908836104f5565b1660408201526000602083013567ffffffffffffffff811680821461092c57600080fd5b6060840152506040830135368490037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261096857600080fd5b830160208101903567ffffffffffffffff81111561098557600080fd5b80360382131561099457600080fd5b606060808501526109a960a085018284610881565b9695505050505050565b6000826109e9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000815180845260005b81811015610a14576020818501810151868301820152016109f8565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610a8760a08301866109ee565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610ad2578151865260209586019590910190600101610ab4565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610b6f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610b5a8583516109ee565b94506020938401939190910190600101610b20565b505050508281036040840152610b858186610aa0565b915050610b96606083018415159052565b9594505050505056fea264697066735822122019741560ffed69bfe2dceeeb3d46120ee7f3cc17870cc7eb97a63cfb8ccf845564736f6c634300081a00335a657461436f6e6e6563746f724e6f6e4e617469766555706772616465546573742e736f6ca26469706673582212209ed1ce9f95f21fd6a23a2a935f1e6c77a442669b397de4c91fa4342c539da19564736f6c634300081a0033",
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

// TestUpgradeAndWithdraw is a paid mutator transaction binding the contract method 0xaf298bb1.
//
// Solidity: function testUpgradeAndWithdraw() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactor) TestUpgradeAndWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.contract.Transact(opts, "testUpgradeAndWithdraw")
}

// TestUpgradeAndWithdraw is a paid mutator transaction binding the contract method 0xaf298bb1.
//
// Solidity: function testUpgradeAndWithdraw() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestSession) TestUpgradeAndWithdraw() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestUpgradeAndWithdraw(&_ZetaConnectorNonNativeTest.TransactOpts)
}

// TestUpgradeAndWithdraw is a paid mutator transaction binding the contract method 0xaf298bb1.
//
// Solidity: function testUpgradeAndWithdraw() returns()
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestTransactorSession) TestUpgradeAndWithdraw() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeTest.Contract.TestUpgradeAndWithdraw(&_ZetaConnectorNonNativeTest.TransactOpts)
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

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestReceivedRevertIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestReceivedRevertIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
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

// ParseReceivedRevert is a log parse operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
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

// FilterReverted is a free log retrieval operation binding the contract event 0x1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b1436.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
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

// WatchReverted is a free log subscription operation binding the contract event 0x1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b1436.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
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

// ParseReverted is a log parse operation binding the contract event 0x1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b1436.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
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
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestUpdatedGatewayTSSAddressIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
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

// ParseUpdatedGatewayTSSAddress is a log parse operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
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
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedZetaConnectorTSSAddress is a free log retrieval operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterUpdatedZetaConnectorTSSAddress(opts *bind.FilterOpts) (*ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestUpdatedZetaConnectorTSSAddressIterator{contract: _ZetaConnectorNonNativeTest.contract, event: "UpdatedZetaConnectorTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedZetaConnectorTSSAddress is a free log subscription operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
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

// ParseUpdatedZetaConnectorTSSAddress is a log parse operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
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

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
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

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
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

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*ZetaConnectorNonNativeTestWithdrawnAndReverted, error) {
	event := new(ZetaConnectorNonNativeTestWithdrawnAndReverted)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestWithdrawnV2Iterator struct {
	Event *ZetaConnectorNonNativeTestWithdrawnV2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNonNativeTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeTestWithdrawnV2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNonNativeTestWithdrawnV2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNonNativeTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeTestWithdrawnV2 represents a WithdrawnV2 event raised by the ZetaConnectorNonNativeTest contract.
type ZetaConnectorNonNativeTestWithdrawnV2 struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnV2 is a free log retrieval operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeTestWithdrawnV2Iterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.FilterLogs(opts, "WithdrawnV2", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTestWithdrawnV2Iterator{contract: _ZetaConnectorNonNativeTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeTestWithdrawnV2, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeTest.contract.WatchLogs(opts, "WithdrawnV2", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeTestWithdrawnV2)
				if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNonNativeTest *ZetaConnectorNonNativeTestFilterer) ParseWithdrawnV2(log types.Log) (*ZetaConnectorNonNativeTestWithdrawnV2, error) {
	event := new(ZetaConnectorNonNativeTestWithdrawnV2)
	if err := _ZetaConnectorNonNativeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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
