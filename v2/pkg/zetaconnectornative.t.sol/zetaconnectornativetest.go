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
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testWithdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20Partial\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveNoParams\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevert\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5061ef0b8061003c6000396000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c8063b0464fdc116100e3578063dcf7d0371161008c578063e63ab1e911610066578063e63ab1e9146102ce578063fa7626d4146102f5578063fe574f841461030257600080fd5b8063dcf7d037146102b6578063de1cb76c146102be578063e20c9f71146102c657600080fd5b8063c1909972116100bd578063c19099721461029e578063ccb0e3f2146102a6578063d509b16c146102ae57600080fd5b8063b0464fdc14610276578063b5508aa91461027e578063ba414fa61461028657600080fd5b8063493465581161014557806385226c811161011f57806385226c811461021757806385f438c11461022c578063916a17c61461026157600080fd5b806349346558146101f257806366d9a9a0146101fa578063828320141461020f57600080fd5b80633cba0107116101765780633cba0107146101da5780633e5e3c23146101e25780633f7286f4146101ea57600080fd5b80630a9254e41461019d5780631ed7831c146101a75780632ade3880146101c5575b600080fd5b6101a561030a565b005b6101af610b0d565b6040516101bc9190618303565b60405180910390f35b6101cd610b6f565b6040516101bc919061839f565b6101a5610cb1565b6101af61146b565b6101af6114cb565b6101a561152b565b610202611b6e565b6040516101bc9190618505565b6101a5611cf0565b61021f611f4c565b6040516101bc91906185a3565b6102537f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6040519081526020016101bc565b61026961201c565b6040516101bc919061861a565b610269612117565b61021f612212565b61028e6122e2565b60405190151581526020016101bc565b6101a56123b6565b6101a5612621565b6101a561314d565b6101a5613189565b6101a5613830565b6101af613e87565b6102537f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b601f5461028e9060ff1681565b6101a5613ee7565b602580547fffffffffffffffffffffffff000000000000000000000000000000000000000090811630179091556026805482166112341790556027805490911661567817905560405161035c90618216565b604080825260049082018190527f7a6574610000000000000000000000000000000000000000000000000000000060608301526080602083018190528201527f5a4554410000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156103e0573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081178255604080518082018252600e81527f4761746577617945564d2e736f6c0000000000000000000000000000000000006020820152602754602554925190861694810194909452604484019290925290921660648201526104d1919060840160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b00000000000000000000000000000000000000000000000000000000179052614105565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116811790915560275460255460405192939182169291169061055d90618223565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610599573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392831617905560205460245460275460255460405193851694928316939183169216906105f490618230565b6001600160a01b039485168152928416602084015290831660408301529091166060820152608001604051809103906000f080158015610638573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039290921691909117905560405161067d9061823d565b604051809103906000f080158015610699573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556027546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b15801561074557600080fd5b505af1158015610759573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b1580156107cf57600080fd5b505af11580156107e3573d6000803e3d6000fd5b50506020546022546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f9150602401600060405180830381600087803b15801561084957600080fd5b505af115801561085d573d6000803e3d6000fd5b50506020546023546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef9150602401600060405180830381600087803b1580156108c357600080fd5b505af11580156108d7573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561093957600080fd5b505af115801561094d573d6000803e3d6000fd5b5050602480546023546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152624c4b40938101939093521692506340c10f199150604401600060405180830381600087803b1580156109be57600080fd5b505af11580156109d2573d6000803e3d6000fd5b50506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b158015610a5657600080fd5b505af1158015610a6a573d6000803e3d6000fd5b5050604080516060810182526024546001600160a01b039081168252600160208084019182528451908101855260008152938301849052825160288054925167ffffffffffffffff1674010000000000000000000000000000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090931691909316171781559093509150602990610b089082618774565b505050565b60606016805480602002602001604051908101604052809291908181526020018280548015610b6557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610b47575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020016000905b82821015610ca857600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015610c91578382906000526020600020018054610c04906186e0565b80601f0160208091040260200160405190810160405280929190818152602001828054610c30906186e0565b8015610c7d5780601f10610c5257610100808354040283529160200191610c7d565b820191906000526020600020905b815481529060010190602001808311610c6057829003601f168201915b505050505081526020019060010190610be5565b505050508152505081526020019060010190610b93565b50505050905090565b60248054602654604051620186a09381018490526001600160a01b03928316604482015291166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a2000000000000000000000000000000000000000000000000000000001790526024805460265492516370a0823160e01b81526001600160a01b0393841660048201529394506000939216916370a082319101602060405180830381865afa158015610d8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db09190618833565b9050610dbd816000614124565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015610e0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e319190618833565b6020546040516001600160a01b0390911660248201526044810187905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391610f14916001600160a01b039190911690600090869060040161884c565b600060405180830381600087803b158015610f2e57600080fd5b505af1158015610f42573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015610fd457600080fd5b505af1158015610fe8573d6000803e3d6000fd5b505060208054602454602654604080516001600160a01b0394851681529485018d905291831684830152919091166060830152517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609350908190036080019150a16023546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156110d757600080fd5b505af11580156110eb573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d91506111309089908890618874565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561119157600080fd5b505af11580156111a5573d6000803e3d6000fd5b50506023546021546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef93506111fd92909116908a9089908b9060040161888d565b600060405180830381600087803b15801561121757600080fd5b505af115801561122b573d6000803e3d6000fd5b5050602480546026546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a0823191015b602060405180830381865afa15801561127e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112a29190618833565b90506112ae8188614124565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa1580156112fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113229190618833565b9050611337816113328a876188f5565b614124565b602480546020546021546040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529082169381019390935260009291169063dd62ed3e90604401602060405180830381865afa1580156113ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113d19190618833565b90506113de816000614124565b602480546020546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa15801561142e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114529190618833565b905061145f816000614124565b50505050505050505050565b60606018805480602002602001604051908101604052809291908181526020018280548015610b65576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610b47575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015610b65576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610b47575050505050905090565b604080516004808252602480830184526020830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed70169000000000000000000000000000000000000000000000000000000001790525460265493516370a0823160e01b8152620186a0946000949385936001600160a01b03908116936370a08231936115cc93921691016001600160a01b0391909116815260200190565b602060405180830381865afa1580156115e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061160d9190618833565b905061161a816000614124565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa15801561166a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061168e9190618833565b6020546040516001600160a01b0390911660248201526044810187905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391611771916001600160a01b039190911690600090869060040161884c565b600060405180830381600087803b15801561178b57600080fd5b505af115801561179f573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561183157600080fd5b505af1158015611845573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0935001905060405180910390a16023546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561191757600080fd5b505af115801561192b573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d91506119709089908890618874565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156119d157600080fd5b505af11580156119e5573d6000803e3d6000fd5b50506023546021546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef9350611a3d92909116908a9089908b9060040161888d565b600060405180830381600087803b158015611a5757600080fd5b505af1158015611a6b573d6000803e3d6000fd5b5050602480546026546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa158015611abd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ae19190618833565b9050611aee816000614124565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015611b3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b629190618833565b90506113378185614124565b6060601b805480602002602001604051908101604052809291908181526020016000905b82821015610ca85783829060005260206000209060020201604051806040016040529081600082018054611bc5906186e0565b80601f0160208091040260200160405190810160405280929190818152602001828054611bf1906186e0565b8015611c3e5780601f10611c1357610100808354040283529160200191611c3e565b820191906000526020600020905b815481529060010190602001808311611c2157829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015611cd857602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411611c855790505b50505050508152505081526020019060010190611b92565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a090600090819060250160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611d8d57600080fd5b505af1158015611da1573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611e8c9190600401618908565b600060405180830381600087803b158015611ea657600080fd5b505af1158015611eba573d6000803e3d6000fd5b50506023546021546040517f057e0f250000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063057e0f259350611f1592909116908790869088906028906004016189f0565b600060405180830381600087803b158015611f2f57600080fd5b505af1158015611f43573d6000803e3d6000fd5b50505050505050565b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015610ca8578382906000526020600020018054611f8f906186e0565b80601f0160208091040260200160405190810160405280929190818152602001828054611fbb906186e0565b80156120085780601f10611fdd57610100808354040283529160200191612008565b820191906000526020600020905b815481529060010190602001808311611feb57829003601f168201915b505050505081526020019060010190611f70565b6060601d805480602002602001604051908101604052809291908181526020016000905b82821015610ca85760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156120ff57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116120ac5790505b50505050508152505081526020019060010190612040565b6060601c805480602002602001604051908101604052809291908181526020016000905b82821015610ca85760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156121fa57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116121a75790505b5050505050815250508152602001906001019061213b565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015610ca8578382906000526020600020018054612255906186e0565b80601f0160208091040260200160405190810160405280929190818152602001828054612281906186e0565b80156122ce5780601f106122a3576101008083540402835291602001916122ce565b820191906000526020600020905b8154815290600101906020018083116122b157829003601f168201915b505050505081526020019060010190612236565b60085460009060ff16156122fa575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa15801561238b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123af9190618833565b1415905090565b60248054602654604051620186a09381018490526001600160a01b03928316604482015291166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a200000000000000000000000000000000000000000000000000000000179052602554905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561249c57600080fd5b505af11580156124b0573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061259b9190600401618908565b600060405180830381600087803b1580156125b557600080fd5b505af11580156125c9573d6000803e3d6000fd5b50506023546021546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef9350611f15929091169087908690889060040161888d565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090600090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561268257600080fd5b505af1158015612696573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506127819190600401618908565b600060405180830381600087803b15801561279b57600080fd5b505af11580156127af573d6000803e3d6000fd5b50505050602360009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561280357600080fd5b505af1158015612817573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561287457600080fd5b505af1158015612888573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506129739190600401618908565b600060405180830381600087803b15801561298d57600080fd5b505af11580156129a1573d6000803e3d6000fd5b50505050602360009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156129f557600080fd5b505af1158015612a09573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015612a6657600080fd5b505af1158015612a7a573d6000803e3d6000fd5b50505050602360009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b158015612ace57600080fd5b505af1158015612ae2573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd93c0665000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015612b6b57600080fd5b505af1158015612b7f573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015612bdc57600080fd5b505af1158015612bf0573d6000803e3d6000fd5b50506023546026546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101879052604481018690529116925063106e62909150606401600060405180830381600087803b158015612c6457600080fd5b505af1158015612c78573d6000803e3d6000fd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015612cd557600080fd5b505af1158015612ce9573d6000803e3d6000fd5b50505050602360009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015612d3d57600080fd5b505af1158015612d51573d6000803e3d6000fd5b5050602480546026546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a0823191015b602060405180830381865afa158015612da4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dc89190618833565b9050612dd5816000614124565b6026546040516001600160a01b0390911660248201526044810184905260009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391612eb5916001600160a01b039190911690600090869060040161884c565b600060405180830381600087803b158015612ecf57600080fd5b505af1158015612ee3573d6000803e3d6000fd5b50506023546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015612f7557600080fd5b505af1158015612f89573d6000803e3d6000fd5b50506026546040518781526001600160a01b0390911692507f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5915060200160405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561302857600080fd5b505af115801561303c573d6000803e3d6000fd5b50506023546026546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101899052604481018890529116925063106e62909150606401600060405180830381600087803b1580156130b057600080fd5b505af11580156130c4573d6000803e3d6000fd5b5050602480546026546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa158015613116573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061313a9190618833565b90506131468186614124565b5050505050565b602480546026546040516370a0823160e01b81526001600160a01b039182166004820152620186a093600093849316916370a082319101612d87565b60248054602654604051620186a09381018490526001600160a01b03928316604482015291166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc5131691000000000000000000000000000000000000000000000000000000001790526024805460265492516370a0823160e01b81526001600160a01b0393841660048201529394506000939216916370a082319101602060405180830381865afa158015613264573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132889190618833565b9050613295816000614124565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa1580156132e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133099190618833565b6020546040516001600160a01b0390911660248201526044810187905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba3916133ec916001600160a01b039190911690600090869060040161884c565b600060405180830381600087803b15801561340657600080fd5b505af115801561341a573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156134ac57600080fd5b505af11580156134c0573d6000803e3d6000fd5b50506020547f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af6092506001600160a01b031690506134fe600289618a3c565b602454602654604080516001600160a01b03958616815260208101949094529184168383015292909216606082015290519081900360800190a16023546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156135c657600080fd5b505af11580156135da573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d915061361f9089908890618874565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561368057600080fd5b505af1158015613694573d6000803e3d6000fd5b50506023546021546040517f5e3e9fef0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450635e3e9fef93506136ec92909116908a9089908b9060040161888d565b600060405180830381600087803b15801561370657600080fd5b505af115801561371a573d6000803e3d6000fd5b5050602480546026546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101602060405180830381865afa15801561376c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137909190618833565b90506137a18161133260028a618a3c565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa1580156137f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138159190618833565b90506113378161382660028b618a3c565b61133290876188f5565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a090600090819060250160408051808303601f1901815290829052602480546021546370a0823160e01b85526001600160a01b0390811660048601529294506000939216916370a082319101602060405180830381865afa1580156138c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138e99190618833565b90506138f6816000614124565b602480546023546040516370a0823160e01b81526001600160a01b03918216600482015260009391909216916370a082319101602060405180830381865afa158015613946573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061396a9190618833565b6020546040516001600160a01b0390911660248201526044810187905290915060009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260245490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391613a4d916001600160a01b039190911690600090869060040161884c565b600060405180830381600087803b158015613a6757600080fd5b505af1158015613a7b573d6000803e3d6000fd5b50506021546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015613b0d57600080fd5b505af1158015613b21573d6000803e3d6000fd5b50506020546040517f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e9350613b6592506001600160a01b0390911690602890618a77565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015613bfb57600080fd5b505af1158015613c0f573d6000803e3d6000fd5b50506024546021546040516001600160a01b039283169450911691507f1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b143690613c5d908a908990602890618a99565b60405180910390a36023546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015613cf357600080fd5b505af1158015613d07573d6000803e3d6000fd5b50506021546040516001600160a01b0390911692507f52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c9150613d4f9089908890602890618a99565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613db057600080fd5b505af1158015613dc4573d6000803e3d6000fd5b50506023546021546040517f057e0f250000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063057e0f259350613e1f92909116908a9089908b906028906004016189f0565b600060405180830381600087803b158015613e3957600080fd5b505af1158015613e4d573d6000803e3d6000fd5b5050602480546021546040516370a0823160e01b81526001600160a01b03918216600482015260009550911692506370a082319101611261565b60606015805480602002602001604051908101604052809291908181526020018280548015610b65576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610b47575050505050905090565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090600090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613f4857600080fd5b505af1158015613f5c573d6000803e3d6000fd5b5050602554604080516001600160a01b0390921660248301527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506140479190600401618908565b600060405180830381600087803b15801561406157600080fd5b505af1158015614075573d6000803e3d6000fd5b50506023546026546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101879052604481018690529116925063106e62909150606401600060405180830381600087803b1580156140e957600080fd5b505af11580156140fd573d6000803e3d6000fd5b505050505050565b600061410f61824a565b61411a8484836141a3565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c549060440160006040518083038186803b15801561418f57600080fd5b505afa1580156140fd573d6000803e3d6000fd5b6000806141b0858461421e565b90506142136040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016141fe929190618ac4565b6040516020818303038152906040528561422a565b9150505b9392505050565b60006142178383614258565b60c0810151516000901561424e5761424784848460c00151614273565b9050614217565b6142478484614419565b60006142648383614504565b6142178383602001518461422a565b60008061427e614514565b9050600061428c86836145e7565b905060006142a38260600151836020015185614a8d565b905060006142b383838989614c9f565b905060006142c082615b1c565b602081015181519192509060030b15614333578982604001516040516020016142ea929190618ae6565b60408051601f19818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261432a91600401618908565b60405180910390fd5b60006143766040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001615ceb565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906143c9908490600401618908565b602060405180830381865afa1580156143e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061440a9190618b67565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc9259061446e908790600401618908565b600060405180830381865afa15801561448b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526144b39190810190618c49565b905060006144e182856040516020016144cd929190618c7e565b604051602081830303815290604052615eeb565b90506001600160a01b03811661411a5784846040516020016142ea929190618cad565b61451082826000615efe565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c9061459b908490600401618d58565b600060405180830381865afa1580156145b8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526145e09190810190618d9f565b9250505090565b6146196040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d90506146646040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b61466d85616001565b6020820152600061467d866163e6565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa1580156146bf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526146e79190810190618d9f565b868385602001516040516020016147019493929190618de8565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb1190614759908590600401618908565b600060405180830381865afa158015614776573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261479e9190810190618d9f565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906147e6908490600401618eec565b602060405180830381865afa158015614803573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906148279190618f3e565b61483c57816040516020016142ea9190618f60565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890614881908490600401618ff2565b600060405180830381865afa15801561489e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526148c69190810190618d9f565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f69061490d908490600401619044565b602060405180830381865afa15801561492a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061494e9190618f3e565b156149e3576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890614998908490600401619044565b600060405180830381865afa1580156149b5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526149dd9190810190618d9f565b60408501525b846001600160a01b03166349c4fac8828660000151604051602001614a089190619096565b6040516020818303038152906040526040518363ffffffff1660e01b8152600401614a34929190619102565b600060405180830381865afa158015614a51573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052614a799190810190618d9f565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b6060815260200190600190039081614aa95790505090506040518060400160405280600481526020017f677265700000000000000000000000000000000000000000000000000000000081525081600081518110614b0957614b09619127565b60200260200101819052506040518060400160405280600381526020017f2d726c000000000000000000000000000000000000000000000000000000000081525081600181518110614b5d57614b5d619127565b602002602001018190525084604051602001614b799190619156565b60405160208183030381529060405281600281518110614b9b57614b9b619127565b602002602001018190525082604051602001614bb791906191c2565b60405160208183030381529060405281600381518110614bd957614bd9619127565b60200260200101819052506000614bef82615b1c565b602080820151604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000008185019081528251808401845260008082529086015282518084019093529051825292810192909252919250614c809060408051808201825260008082526020918201528151808301909252845182528085019082015290616669565b614c9557856040516020016142ea9190619203565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015614cef565b511590565b614e6357826020015115614dab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a40161432a565b8260c0015115614e63576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a40161432a565b6040805160ff8082526120008201909252600091816020015b6060815260200190600190039081614e7c57905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280614ed790619294565b935060ff1681518110614eec57614eec619127565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e3700000000000000000000000000000000000000815250604051602001614f3d91906192b3565b604051602081830303815290604052828280614f5890619294565b935060ff1681518110614f6d57614f6d619127565b60200260200101819052506040518060400160405280600681526020017f6465706c6f790000000000000000000000000000000000000000000000000000815250828280614fba90619294565b935060ff1681518110614fcf57614fcf619127565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d6500000000000000000000000000000000000081525082828061501c90619294565b935060ff168151811061503157615031619127565b6020026020010181905250876020015182828061504d90619294565b935060ff168151811061506257615062619127565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e7472616374506174680000000000000000000000000000000000008152508282806150af90619294565b935060ff16815181106150c4576150c4619127565b6020908102919091010152875182826150dc81619294565b935060ff16815181106150f1576150f1619127565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e4964000000000000000000000000000000000000000000000081525082828061513e90619294565b935060ff168151811061515357615153619127565b6020026020010181905250615167466166ca565b828261517281619294565b935060ff168151811061518757615187619127565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806151d490619294565b935060ff16815181106151e9576151e9619127565b60200260200101819052508682828061520190619294565b935060ff168151811061521657615216619127565b602090810291909101015285511561533d5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f646500000000000000000000006020820152828261526781619294565b935060ff168151811061527c5761527c619127565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906152cc908990600401618908565b600060405180830381865afa1580156152e9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526153119190810190618d9f565b828261531c81619294565b935060ff168151811061533157615331619127565b60200260200101819052505b84602001511561540d5760408051808201909152601281527f2d2d766572696679536f75726365436f646500000000000000000000000000006020820152828261538681619294565b935060ff168151811061539b5761539b619127565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806153e890619294565b935060ff16815181106153fd576153fd619127565b60200260200101819052506155d4565b615445614cea8660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6154d85760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261548881619294565b935060ff168151811061549d5761549d619127565b60200260200101819052508460a001516040516020016154bd9190619156565b6040516020818303038152906040528282806153e890619294565b8460c0015115801561551b57506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261551990511590565b155b156155d45760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261555f81619294565b935060ff168151811061557457615574619127565b60200260200101819052506155888861676a565b6040516020016155989190619156565b6040516020818303038152906040528282806155b390619294565b935060ff16815181106155c8576155c8619127565b60200260200101819052505b6040808601518151808301835260008082526020918201528251808401909352815183529081019082015261560890511590565b61569d5760408051808201909152600b81527f2d2d72656c6179657249640000000000000000000000000000000000000000006020820152828261564b81619294565b935060ff168151811061566057615660619127565b6020026020010181905250846040015182828061567c90619294565b935060ff168151811061569157615691619127565b60200260200101819052505b6060850151156157be5760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826156e681619294565b935060ff16815181106156fb576156fb619127565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa15801561576a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526157929190810190618d9f565b828261579d81619294565b935060ff16815181106157b2576157b2619127565b60200260200101819052505b60e085015151156158655760408051808201909152600a81527f2d2d6761734c696d6974000000000000000000000000000000000000000000006020820152828261580881619294565b935060ff168151811061581d5761581d619127565b60200260200101819052506158398560e00151600001516166ca565b828261584481619294565b935060ff168151811061585957615859619127565b60200260200101819052505b60e0850151602001511561590f5760408051808201909152600a81527f2d2d676173507269636500000000000000000000000000000000000000000000602082015282826158b281619294565b935060ff16815181106158c7576158c7619127565b60200260200101819052506158e38560e00151602001516166ca565b82826158ee81619294565b935060ff168151811061590357615903619127565b60200260200101819052505b60e085015160400151156159b95760408051808201909152600e81527f2d2d6d61784665655065724761730000000000000000000000000000000000006020820152828261595c81619294565b935060ff168151811061597157615971619127565b602002602001018190525061598d8560e00151604001516166ca565b828261599881619294565b935060ff16815181106159ad576159ad619127565b60200260200101819052505b60e08501516060015115615a635760408051808201909152601681527f2d2d6d61785072696f726974794665655065724761730000000000000000000060208201528282615a0681619294565b935060ff1681518110615a1b57615a1b619127565b6020026020010181905250615a378560e00151606001516166ca565b8282615a4281619294565b935060ff1681518110615a5757615a57619127565b60200260200101819052505b60008160ff1667ffffffffffffffff811115615a8157615a816186b1565b604051908082528060200260200182016040528015615ab457816020015b6060815260200190600190039081615a9f5790505b50905060005b8260ff168160ff161015615b0d57838160ff1681518110615add57615add619127565b6020026020010151828260ff1681518110615afa57615afa619127565b6020908102919091010152600101615aba565b5093505050505b949350505050565b615b436040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c91615bc99186910161931e565b600060405180830381865afa158015615be6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615c0e9190810190618d9f565b90506000615c1c8683617259565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b8152600401615c4c91906185a3565b6000604051808303816000875af1158015615c6b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615c939190810190619365565b805190915060030b15801590615cac5750602081015151155b8015615cbb5750604081015151155b15614c955781600081518110615cd357615cd3619127565b60200260200101516040516020016142ea919061941b565b60606000615d208560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528651825280870190820152909150615d579082905b906173ae565b15615eb4576000615dd482615dce84615dc8615d9a8a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b906173d5565b90617437565b604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150615e389082906173ae565b15615ea257604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615e9f905b82906174bc565b90505b615eab816174e2565b92505050614217565b8215615ecd5784846040516020016142ea929190619607565b5050604080516020810190915260008152614217565b509392505050565b6000808251602084016000f09392505050565b8160a0015115615f0d57505050565b6000615f1a84848461754b565b90506000615f2782615b1c565b602081015181519192509060030b158015615fc35750604080518082018252600781527f535543434553530000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615fc390604080518082018252600080825260209182015281518083019092528451825280850190820152615d51565b15615fd057505050505050565b60408201515115615ff05781604001516040516020016142ea91906196ae565b806040516020016142ea919061970c565b606060006160368360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061609b905b8290616669565b1561610a57604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261421790616105908390617ae6565b6174e2565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261616c905b8290617b70565b60010361623957604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526161d290615e98565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261421790616105905b83906174bc565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261629890616094565b156163cf57604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290616300908390617c0a565b90506000816001835161631391906188f5565b8151811061632357616323619127565b602002602001015190506163c66161056163996040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252855182528086019082015290617ae6565b95945050505050565b826040516020016142ea9190619777565b50919050565b6060600061641b8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061647d90616094565b1561648b57614217816174e2565b604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526164ea90616165565b60010361655457604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526142179061610590616232565b604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526165b390616094565b156163cf57604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061661b908390617c0a565b905060018151111561665757806002825161663691906188f5565b8151811061664657616646619127565b602002602001015192505050919050565b50826040516020016142ea9190619777565b80518251600091111561667e5750600061411e565b8151835160208501516000929161669491619855565b61669e91906188f5565b9050826020015181036166b557600191505061411e565b82516020840151819020912014905092915050565b606060006166d783617caf565b600101905060008167ffffffffffffffff8111156166f7576166f76186b1565b6040519080825280601f01601f191660200182016040528015616721576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461672b57509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e53454400000000000000000000000000000000000000000000818401908152855180870187528381528401929092528451808601909552518452908301526060916167f6905b8290617d91565b1561683657505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e7365000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616895906167ef565b156168d557505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d4954000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616934906167ef565b1561697457505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526169d3906167ef565b80616a385750604080518082018252601081527f47504c2d322e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616a38906167ef565b15616a7857505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c79000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616ad7906167ef565b80616b3c5750604080518082018252601081527f47504c2d332e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616b3c906167ef565b15616b7c57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616bdb906167ef565b80616c405750604080518082018252601181527f4c47504c2d322e312d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616c40906167ef565b15616c8057505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616cdf906167ef565b80616d445750604080518082018252601181527f4c47504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616d44906167ef565b15616d8457505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616de3906167ef565b15616e2357505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616e82906167ef565b15616ec257505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616f21906167ef565b15616f6157505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616fc0906167ef565b1561700057505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261705f906167ef565b1561709f57505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526170fe906167ef565b806171635750604080518082018252601181527f4147504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617163906167ef565b156171a357505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e3100000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617202906167ef565b1561724257505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516142ea9290602001619868565b60608060005b84518110156172e4578185828151811061727b5761727b619127565b6020026020010151604051602001617294929190618c7e565b6040516020818303038152906040529150600185516172b391906188f5565b81146172dc57816040516020016172ca91906199d1565b60405160208183030381529060405291505b60010161725f565b5060408051600380825260808201909252600091816020015b60608152602001906001900390816172fd579050509050838160008151811061732857617328619127565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061737c5761737c619127565b6020026020010181905250818160028151811061739b5761739b619127565b6020908102919091010152949350505050565b60208083015183518351928401516000936173cc9291849190617da5565b14159392505050565b604080518082019091526000808252602082015260006174078460000151856020015185600001518660200151617eb6565b905083602001518161741991906188f5565b845185906174289083906188f5565b90525060208401525090919050565b604080518082019091526000808252602082015281518351101561745c57508161411e565b60208083015190840151600191146174835750815160208481015190840151829020919020145b80156174b45782518451859061749a9083906188f5565b90525082516020850180516174b0908390619855565b9052505b509192915050565b60408051808201909152600080825260208201526174db838383617fd6565b5092915050565b60606000826000015167ffffffffffffffff811115617503576175036186b1565b6040519080825280601f01601f19166020018201604052801561752d576020820181803683370190505b50905060006020820190506174db8185602001518660000151618081565b60606000617557614514565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161757457905050905060006040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806175cf90619294565b935060ff16815181106175e4576175e4619127565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e33000000000000000000000000000000000000000000000000008152506040516020016176359190619a12565b60405160208183030381529060405282828061765090619294565b935060ff168151811061766557617665619127565b60200260200101819052506040518060400160405280600881526020017f76616c69646174650000000000000000000000000000000000000000000000008152508282806176b290619294565b935060ff16815181106176c7576176c7619127565b6020026020010181905250826040516020016176e391906191c2565b6040516020818303038152906040528282806176fe90619294565b935060ff168151811061771357617713619127565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061776090619294565b935060ff168151811061777557617775619127565b602002602001018190525061778a87846180fb565b828261779581619294565b935060ff16815181106177aa576177aa619127565b6020908102919091010152855151156178565760408051808201909152600b81527f2d2d7265666572656e6365000000000000000000000000000000000000000000602082015282826177fc81619294565b935060ff168151811061781157617811619127565b602002602001018190525061782a8660000151846180fb565b828261783581619294565b935060ff168151811061784a5761784a619127565b60200260200101819052505b8560800151156178c45760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261789f81619294565b935060ff16815181106178b4576178b4619127565b602002602001018190525061792a565b841561792a5760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261790981619294565b935060ff168151811061791e5761791e619127565b60200260200101819052505b604086015151156179c65760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261797481619294565b935060ff168151811061798957617989619127565b602002602001018190525085604001518282806179a590619294565b935060ff16815181106179ba576179ba619127565b60200260200101819052505b856060015115617a305760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d657300000000000000000000000060208201528282617a0f81619294565b935060ff1681518110617a2457617a24619127565b60200260200101819052505b60008160ff1667ffffffffffffffff811115617a4e57617a4e6186b1565b604051908082528060200260200182016040528015617a8157816020015b6060815260200190600190039081617a6c5790505b50905060005b8260ff168160ff161015617ada57838160ff1681518110617aaa57617aaa619127565b6020026020010151828260ff1681518110617ac757617ac7619127565b6020908102919091010152600101617a87565b50979650505050505050565b6040805180820190915260008082526020820152815183511015617b0b57508161411e565b81518351602085015160009291617b2191619855565b617b2b91906188f5565b60208401519091506001908214617b4c575082516020840151819020908220145b8015617b6757835185518690617b639083906188f5565b9052505b50929392505050565b6000808260000151617b948560000151866020015186600001518760200151617eb6565b617b9e9190619855565b90505b83516020850151617bb29190619855565b81116174db5781617bc281619a57565b9250508260000151617bf9856020015183617bdd91906188f5565b8651617be991906188f5565b8386600001518760200151617eb6565b617c039190619855565b9050617ba1565b60606000617c188484617b70565b617c23906001619855565b67ffffffffffffffff811115617c3b57617c3b6186b1565b604051908082528060200260200182016040528015617c6e57816020015b6060815260200190600190039081617c595790505b50905060005b8151811015615ee357617c8a61610586866174bc565b828281518110617c9c57617c9c619127565b6020908102919091010152600101617c74565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310617cf8577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310617d24576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310617d4257662386f26fc10000830492506010015b6305f5e1008310617d5a576305f5e100830492506008015b6127108310617d6e57612710830492506004015b60648310617d80576064830492506002015b600a831061411e5760010192915050565b6000617d9d838361813b565b159392505050565b600080858411617eac5760208411617e585760008415617df0576001617dcc8660206188f5565b617dd7906008619a71565b617de2906002619b6f565b617dec91906188f5565b1990505b8351811685617dff8989619855565b617e0991906188f5565b805190935082165b818114617e4357878411617e2b5787945050505050615b14565b83617e3581619b7b565b945050828451169050617e11565b617e4d8785619855565b945050505050615b14565b838320617e6585886188f5565b617e6f9087619855565b91505b858210617eaa57848220808203617e9757617e8d8684619855565b9350505050615b14565b617ea26001846188f5565b925050617e72565b505b5092949350505050565b60008381868511617fc15760208511617f705760008515617f02576001617ede8760206188f5565b617ee9906008619a71565b617ef4906002619b6f565b617efe91906188f5565b1990505b84518116600087617f138b8b619855565b617f1d91906188f5565b855190915083165b828114617f6257818610617f4a57617f3d8b8b619855565b9650505050505050615b14565b85617f5481619a57565b965050838651169050617f25565b859650505050505050615b14565b508383206000905b617f8286896188f5565b8211617fbf57858320808203617f9e5783945050505050615b14565b617fa9600185619855565b9350508180617fb790619a57565b925050617f78565b505b617fcb8787619855565b979650505050505050565b604080518082019091526000808252602082015260006180088560000151866020015186600001518760200151617eb6565b60208087018051918601919091525190915061802490826188f5565b8352845160208601516180379190619855565b81036180465760008552618078565b835183516180549190619855565b855186906180639083906188f5565b90525083516180729082619855565b60208601525b50909392505050565b602081106180b95781518352618098602084619855565b92506180a5602083619855565b91506180b26020826188f5565b9050618081565b60001981156180e85760016180cf8360206188f5565b6180db90610100619b6f565b6180e591906188f5565b90505b9151835183169219169190911790915250565b6060600061810984846145e7565b805160208083015160405193945061812393909101619b92565b60405160208183030381529060405291505092915050565b815181516000919081111561814e575081515b6020808501519084015160005b8381101561820757825182518082146181d75760001960208710156181b6576001846181888960206188f5565b6181929190619855565b61819d906008619a71565b6181a8906002619b6f565b6181b291906188f5565b1990505b81811683821681810391146181d457975061411e9650505050505050565b50505b6181e2602086619855565b94506181ef602085619855565b935050506020816182009190619855565b905061815b565b5084518651614c959190619bea565b610c9f80619c0b83390190565b611eb98061a8aa83390190565b61188c8061c76383390190565b610ee78061dfef83390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161828d618292565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161828d6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b818110156183445783516001600160a01b031683526020938401939092019160010161831d565b509095945050505050565b60005b8381101561836a578181015183820152602001618352565b50506000910152565b6000815180845261838b81602086016020860161834f565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561849b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b81811015618481577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261846b848651618373565b6020958601959094509290920191600101618431565b5091975050506020948501949290920191506001016183c7565b50929695505050505050565b600081518084526020840193506020830160005b828110156184fb5781517fffffffff00000000000000000000000000000000000000000000000000000000168652602095860195909101906001016184bb565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561849b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051604087526185716040880182618373565b905060208201519150868103602088015261858c81836184a7565b96505050602093840193919091019060010161852d565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561849b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452618605858351618373565b945060209384019391909101906001016185cb565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561849b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261869b60408701826184a7565b9550506020938401939190910190600101618642565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c908216806186f457607f821691505b6020821081036163e0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f821115610b0857806000526020600020601f840160051c810160208510156187545750805b601f840160051c820191505b818110156131465760008155600101618760565b815167ffffffffffffffff81111561878e5761878e6186b1565b6187a28161879c84546186e0565b8461872d565b6020601f8211600181146187d657600083156187be5750848201515b600019600385901b1c1916600184901b178455613146565b600084815260208120601f198516915b8281101561880657878501518255602094850194600190920191016187e6565b50848210156188245786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60006020828403121561884557600080fd5b5051919050565b6001600160a01b03841681528260208201526060604082015260006163c66060830184618373565b828152604060208201526000615b146040830184618373565b6001600160a01b03851681528360208201526080604082015260006188b56080830185618373565b905082606083015295945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561411e5761411e6188c6565b6020815260006142176020830184618373565b600081546001600160a01b038116845267ffffffffffffffff8160a01c16602085015250600182016060604085015260008154618957816186e0565b806060880152600182166000811461897657600181146189b0576189e4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083166080890152608082151560051b89010193506189e4565b84600052602060002060005b838110156189db5781548a8201608001526001909101906020016189bc565b89016080019450505b50919695505050505050565b6001600160a01b038616815284602082015260a060408201526000618a1860a0830186618373565b8460608401528281036080840152618a30818561891b565b98975050505050505050565b600082618a72577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6001600160a01b0383168152604060208201526000615b14604083018461891b565b838152606060208201526000618ab26060830185618373565b8281036040840152614c95818561891b565b6001600160a01b0383168152604060208201526000615b146040830184618373565b7f4661696c656420746f206465706c6f7920636f6e747261637420000000000000815260008351618b1e81601a85016020880161834f565b7f3a20000000000000000000000000000000000000000000000000000000000000601a918401918201528351618b5b81601c84016020880161834f565b01601c01949350505050565b600060208284031215618b7957600080fd5b81516001600160a01b038116811461421757600080fd5b6040516060810167ffffffffffffffff81118282101715618bb357618bb36186b1565b60405290565b60008067ffffffffffffffff841115618bd457618bd46186b1565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff82111715618c0357618c036186b1565b604052838152905080828401851015618c1b57600080fd5b615ee384602083018561834f565b600082601f830112618c3a57600080fd5b61421783835160208501618bb9565b600060208284031215618c5b57600080fd5b815167ffffffffffffffff811115618c7257600080fd5b61411a84828501618c29565b60008351618c9081846020880161834f565b835190830190618ca481836020880161834f565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e747261637420000000000000815260008351618ce581601a85016020880161834f565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a918401918201528351618d2281603384016020880161834f565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f555400000000000000000000000000000000000000000060608201526080602082015260006142176080830184618373565b600060208284031215618db157600080fd5b815167ffffffffffffffff811115618dc857600080fd5b8201601f81018413618dd957600080fd5b61411a84825160208401618bb9565b60008551618dfa818460208a0161834f565b7f2f000000000000000000000000000000000000000000000000000000000000009083019081528551618e34816001840160208a0161834f565b7f2f00000000000000000000000000000000000000000000000000000000000000600192909101918201528451618e7281600284016020890161834f565b6001818301019150507f2f0000000000000000000000000000000000000000000000000000000000000060018201528351618eb481600284016020880161834f565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b604081526000618eff6040830184618373565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b600060208284031215618f5057600080fd5b8151801515811461421757600080fd5b7f436f756c64206e6f742066696e642041535420696e2061727469666163742000815260008251618f9881601f85016020870161834f565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b6040815260006190056040830184618373565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b6040815260006190576040830184618373565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b270000000000000000000000008152600082516190ce81601485016020870161834f565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b6040815260006191156040830185618373565b82810360208401526142138185618373565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161918e81600185016020870161834f565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b600082516191d481846020870161834f565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161928781604b85016020870161834f565b91909101604b0192915050565b600060ff821660ff81036192aa576192aa6188c6565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161931181602985016020870161834f565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f504154480000000000000000000060608201526080602082015260006142176080830184618373565b60006020828403121561937757600080fd5b815167ffffffffffffffff81111561938e57600080fd5b8201606081850312156193a057600080fd5b6193a8618b90565b81518060030b81146193b957600080fd5b8152602082015167ffffffffffffffff8111156193d557600080fd5b6193e186828501618c29565b602083015250604082015167ffffffffffffffff81111561940157600080fd5b61940d86828501618c29565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161947981602185016020870161834f565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161966581602185016020880161834f565b7f2720696e206f75747075743a200000000000000000000000000000000000000060219184019182015283516196a281602e84016020880161834f565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161931181602985016020870161834f565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161976a81602285016020870161834f565b9190910160220192915050565b7f436f6e7472616374206e616d65200000000000000000000000000000000000008152600082516197af81600e85016020870161834f565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b8082018082111561411e5761411e6188c6565b7f53504458206c6963656e7365206964656e7469666965722000000000000000008152600083516198a081601885016020880161834f565b7f20696e200000000000000000000000000000000000000000000000000000000060189184019182015283516198dd81601c84016020880161834f565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b600082516199e381846020870161834f565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f72654000000000815260008251619a4a81601c85016020870161834f565b91909101601c0192915050565b60006000198203619a6a57619a6a6188c6565b5060010190565b808202811582820484141761411e5761411e6188c6565b6001815b6001841115619ac357808504811115619aa757619aa76188c6565b6001841615619ab557908102905b60019390931c928002619a8c565b935093915050565b600082619ada5750600161411e565b81619ae75750600061411e565b8160018114619afd5760028114619b0757619b23565b600191505061411e565b60ff841115619b1857619b186188c6565b50506001821b61411e565b5060208310610133831016604e8410600b8410161715619b46575081810a61411e565b619b536000198484619a88565b8060001904821115619b6757619b676188c6565b029392505050565b60006142178383619acb565b600081619b8a57619b8a6188c6565b506000190190565b60008351619ba481846020880161834f565b7f3a000000000000000000000000000000000000000000000000000000000000009083019081528351619bde81600184016020880161834f565b01600101949350505050565b81810360008312801583831316838312821617156174db576174db6188c656fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220a043c41353215fce25ecb67a8a4f6f724aaa86dea8dcb0a6975bfb1f10ff3beb64736f6c634300081a003360a060405234801561001057600080fd5b50604051611eb9380380611eb983398101604081905261002f916101fd565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03838116608052600480546001600160a01b0319169184169190911790556100bb60008261014c565b506100e67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261014c565b506101117f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361014c565b5061012a600080516020611e998339815191528261014c565b50610143600080516020611e998339815191528361014c565b50505050610240565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166101d75760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016101db565b5060005b92915050565b80516001600160a01b03811681146101f857600080fd5b919050565b60008060006060848603121561021257600080fd5b61021b846101e1565b9250610229602085016101e1565b9150610237604085016101e1565b90509250925092565b608051611c22610277600039600081816101ca01528181610597015281816105f901528181610a280152610a8a0152611c226000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c806385f438c1116100e3578063d547741f1161008c578063e609055e11610066578063e609055e146103fc578063e63ab1e91461040f578063eab103df1461043657600080fd5b8063d547741f146103b3578063d936547e146103c6578063d9caed12146103e957600080fd5b80639b19251a116100bd5780639b19251a14610385578063a217fddf14610398578063c709ab6e146103a057600080fd5b806385f438c11461030557806391d148541461032c5780639a5904271461037257600080fd5b806336568abe116101455780635b1125911161011f5780635b112591146102d25780635c975abb146102f25780638456cb59146102fd57600080fd5b806336568abe146102905780633f4ba83a146102a3578063570618e1146102ab57600080fd5b8063248a9ca311610176578063248a9ca314610226578063252f07bf146102585780632f2ff15d1461027d57600080fd5b806301ffc9a71461019d578063116191b6146101c557806321fc65f214610211575b600080fd5b6101b06101ab3660046115ca565b610449565b60405190151581526020015b60405180910390f35b6101ec7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101bc565b61022461021f366004611677565b6104e2565b005b61024a6102343660046116ea565b6000908152600160208190526040909120015490565b6040519081526020016101bc565b6004546101b09074010000000000000000000000000000000000000000900460ff1681565b61022461028b366004611703565b6106e3565b61022461029e366004611703565b61070f565b61022461076d565b61024a7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b6004546101ec9073ffffffffffffffffffffffffffffffffffffffff1681565b60025460ff166101b0565b6102246107a2565b61024a7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101b061033a366004611703565b600091825260016020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610224610380366004611733565b6107d4565b610224610393366004611733565b6108a2565b61024a600081565b6102246103ae366004611750565b610973565b6102246103c1366004611703565b610b79565b6101b06103d4366004611733565b60036020526000908152604090205460ff1681565b6102246103f73660046117f3565b610b9f565b61022461040a366004611834565b610ccb565b61024a7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6102246104443660046118d3565b610f2b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104dc57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104ea610f81565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461051481610fc4565b61051c610fce565b73ffffffffffffffffffffffffffffffffffffffff851660009081526003602052604090205460ff1661057b576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105bc73ffffffffffffffffffffffffffffffffffffffff86167f00000000000000000000000000000000000000000000000000000000000000008661100d565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906106369088908a90899089908990600401611939565b600060405180830381600087803b15801561065057600080fd5b505af1158015610664573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d58686866040516106c993929190611996565b60405180910390a3506106dc6001600055565b5050505050565b600082815260016020819052604090912001546106ff81610fc4565b610709838361108e565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331461075e576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610768828261113b565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61079781610fc4565b61079f6111dc565b50565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107cc81610fc4565b61079f61123b565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6107fe81610fc4565b73ffffffffffffffffffffffffffffffffffffffff821661084b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6108cc81610fc4565b73ffffffffffffffffffffffffffffffffffffffff8216610919576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b61097b610f81565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46109a581610fc4565b6109ad610fce565b73ffffffffffffffffffffffffffffffffffffffff861660009081526003602052604090205460ff16610a0c576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a4d73ffffffffffffffffffffffffffffffffffffffff87167f00000000000000000000000000000000000000000000000000000000000000008761100d565b6040517fd0b492c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610ac99089908b908a908a908a908a90600401611a77565b600060405180830381600087803b158015610ae357600080fd5b505af1158015610af7573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b87878787604051610b5e9493929190611ae8565b60405180910390a350610b716001600055565b505050505050565b60008281526001602081905260409091200154610b9581610fc4565b610709838361113b565b610ba7610f81565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610bd181610fc4565b610bd9610fce565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604090205460ff16610c38576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c5973ffffffffffffffffffffffffffffffffffffffff8416858461100d565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610cb891815260200190565b60405180910390a3506107686001600055565b610cd3610f81565b610cdb610fce565b60045474010000000000000000000000000000000000000000900460ff16610d2f576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841660009081526003602052604090205460ff16610d8e576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015610dfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1f9190611b14565b9050610e4373ffffffffffffffffffffffffffffffffffffffff8616333087611278565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa158015610ed7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610efb9190611b14565b610f059190611b2d565b8787604051610f18959493929190611b67565b60405180910390a250610b716001600055565b6000610f3681610fc4565b506004805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b600260005403610fbd576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61079f81336112be565b60025460ff161561100b576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60405173ffffffffffffffffffffffffffffffffffffffff83811660248301526044820183905261076891859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061134f565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1661113357600083815260016020818152604080842073ffffffffffffffffffffffffffffffffffffffff87168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016104dc565b5060006104dc565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff161561113357600083815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104dc565b6111e46113e5565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b611243610fce565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586112113390565b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526107099186918216906323b872dd90608401611047565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661134b576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044015b60405180910390fd5b5050565b600061137173ffffffffffffffffffffffffffffffffffffffff841683611421565b905080516000141580156113965750808060200190518101906113949190611ba0565b155b15610768576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401611342565b60025460ff1661100b576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606061142f83836000611436565b9392505050565b606081471015611474576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611342565b6000808573ffffffffffffffffffffffffffffffffffffffff16848660405161149d9190611bbd565b60006040518083038185875af1925050503d80600081146114da576040519150601f19603f3d011682016040523d82523d6000602084013e6114df565b606091505b50915091506114ef8683836114f9565b9695505050505050565b60608261150e5761150982611588565b61142f565b8151158015611532575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611581576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401611342565b508061142f565b8051156115985780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156115dc57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461142f57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461079f57600080fd5b60008083601f84011261164057600080fd5b50813567ffffffffffffffff81111561165857600080fd5b60208301915083602082850101111561167057600080fd5b9250929050565b60008060008060006080868803121561168f57600080fd5b853561169a8161160c565b945060208601356116aa8161160c565b935060408601359250606086013567ffffffffffffffff8111156116cd57600080fd5b6116d98882890161162e565b969995985093965092949392505050565b6000602082840312156116fc57600080fd5b5035919050565b6000806040838503121561171657600080fd5b8235915060208301356117288161160c565b809150509250929050565b60006020828403121561174557600080fd5b813561142f8161160c565b60008060008060008060a0878903121561176957600080fd5b86356117748161160c565b955060208701356117848161160c565b945060408701359350606087013567ffffffffffffffff8111156117a757600080fd5b6117b389828a0161162e565b909450925050608087013567ffffffffffffffff8111156117d357600080fd5b87016060818a0312156117e557600080fd5b809150509295509295509295565b60008060006060848603121561180857600080fd5b83356118138161160c565b925060208401356118238161160c565b929592945050506040919091013590565b6000806000806000806080878903121561184d57600080fd5b863567ffffffffffffffff81111561186457600080fd5b61187089828a0161162e565b90975095505060208701356118848161160c565b935060408701359250606087013567ffffffffffffffff8111156118a757600080fd5b6118b389828a0161162e565b979a9699509497509295939492505050565b801515811461079f57600080fd5b6000602082840312156118e557600080fd5b813561142f816118c5565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015283604082015260806060820152600061198b6080830184866118f0565b979650505050505050565b8381526040602082015260006119b06040830184866118f0565b95945050505050565b600081356119c68161160c565b73ffffffffffffffffffffffffffffffffffffffff168352602082013567ffffffffffffffff81168082146119fa57600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112611a3657600080fd5b820160208101903567ffffffffffffffff811115611a5357600080fd5b803603821315611a6257600080fd5b606060408601526119b06060860182846118f0565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a060608201526000611ac960a0830185876118f0565b8281036080840152611adb81856119b9565b9998505050505050505050565b848152606060208201526000611b026060830185876118f0565b828103604084015261198b81856119b9565b600060208284031215611b2657600080fd5b5051919050565b818103818111156104dc577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b606081526000611b7b6060830187896118f0565b8560208401528281036040840152611b948185876118f0565b98975050505050505050565b600060208284031215611bb257600080fd5b815161142f816118c5565b6000825160005b81811015611bde5760208186018101518583015201611bc4565b50600092019182525091905056fea264697066735822122029d8a02c5c60bae7d404d08786de2b2dd6b37e11d10fa76c968c7809ff06f3b864736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60c060405234801561001057600080fd5b5060405161188c38038061188c83398101604081905261002f9161021a565b60016000819055805460ff19169055838383836001600160a01b038416158061005f57506001600160a01b038316155b8061007157506001600160a01b038216155b8061008357506001600160a01b038116155b156100a15760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03808516608052831660a0526100bf60008261014e565b506100ea7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361014e565b506101157f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb8361014e565b506101407f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261014e565b50505050505050505061026e565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166101f45760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101ac3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101f8565b5060005b92915050565b80516001600160a01b038116811461021557600080fd5b919050565b6000806000806080858703121561023057600080fd5b610239856101fe565b9350610247602086016101fe565b9250610255604086016101fe565b9150610263606086016101fe565b905092959194509250565b60805160a0516115b36102d9600039600081816101f7015281816104740152818161052501528181610644015281816107d80152818161088901526109710152600081816101ab01528181610496015281816104f8015281816107fa015261085c01526115b36000f3fe608060405234801561001057600080fd5b50600436106101515760003560e01c80635c975abb116100cd57806391d1485411610081578063a783c78911610066578063a783c78914610326578063d547741f1461034d578063e63ab1e91461036057600080fd5b806391d14854146102d8578063a217fddf1461031e57600080fd5b8063743e0c9b116100b2578063743e0c9b146102965780638456cb59146102a957806385f438c1146102b157600080fd5b80635c975abb146102785780635e3e9fef1461028357600080fd5b806321e093b1116101245780632f2ff15d116101095780632f2ff15d1461024a57806336568abe1461025d5780633f4ba83a1461027057600080fd5b806321e093b1146101f2578063248a9ca31461021957600080fd5b806301ffc9a714610156578063057e0f251461017e578063106e629014610193578063116191b6146101a6575b600080fd5b6101696101643660046110e4565b610387565b60405190151581526020015b60405180910390f35b61019161018c366004611198565b610420565b005b6101916101a1366004611230565b6105f0565b6101cd7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610175565b6101cd7f000000000000000000000000000000000000000000000000000000000000000081565b61023c610227366004611263565b60009081526002602052604090206001015490565b604051908152602001610175565b61019161025836600461127c565b6106cb565b61019161026b36600461127c565b6106f6565b61019161074f565b60015460ff16610169565b6101916102913660046112a8565b610784565b6101916102a4366004611263565b61094f565b610191610999565b61023c7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101696102e636600461127c565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61023c600081565b61023c7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b61019161035b36600461127c565b6109cb565b61023c7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061041a57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104286109f0565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461045281610a33565b61045a610a3d565b6104bb73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000088610a7c565b6040517fd0b492c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610557907f0000000000000000000000000000000000000000000000000000000000000000908b908b908b908b908a90600401611418565b600060405180830381600087803b15801561057157600080fd5b505af1158015610585573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff167f52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c878787866040516105d59493929190611489565b60405180910390a2506105e86001600055565b505050505050565b6105f86109f0565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461062281610a33565b61062a610a3d565b61066b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168585610a7c565b8373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5846040516106b391815260200190565b60405180910390a2506106c66001600055565b505050565b6000828152600260205260409020600101546106e681610a33565b6106f08383610afd565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610745576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106c68282610bfd565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61077981610a33565b610781610cbc565b50565b61078c6109f0565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46107b681610a33565b6107be610a3d565b61081f73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000087610a7c565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906108b9907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a906004016114c0565b600060405180830381600087803b1580156108d357600080fd5b505af11580156108e7573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d86868660405161093593929190611512565b60405180910390a2506109486001600055565b5050505050565b610957610a3d565b61078173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084610d39565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6109c381610a33565b610781610d7f565b6000828152600260205260409020600101546109e681610a33565b6106f08383610bfd565b600260005403610a2c576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6107818133610dd8565b60015460ff1615610a7a576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60405173ffffffffffffffffffffffffffffffffffffffff8381166024830152604482018390526106c691859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610e69565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610bf557600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610b933390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161041a565b50600061041a565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610bf557600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161041a565b610cc4610eff565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526106f09186918216906323b872dd90608401610ab6565b610d87610a3d565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610d0f565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610e65576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044015b60405180910390fd5b5050565b6000610e8b73ffffffffffffffffffffffffffffffffffffffff841683610f3b565b90508051600014158015610eb0575080806020019051810190610eae919061152c565b155b156106c6576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610e5c565b60015460ff16610a7a576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060610f4983836000610f50565b9392505050565b606081471015610f8e576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610e5c565b6000808573ffffffffffffffffffffffffffffffffffffffff168486604051610fb7919061154e565b60006040518083038185875af1925050503d8060008114610ff4576040519150601f19603f3d011682016040523d82523d6000602084013e610ff9565b606091505b5091509150611009868383611013565b9695505050505050565b60608261102857611023826110a2565b610f49565b815115801561104c575073ffffffffffffffffffffffffffffffffffffffff84163b155b1561109b576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610e5c565b5080610f49565b8051156110b25780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156110f657600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610f4957600080fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461114a57600080fd5b919050565b60008083601f84011261116157600080fd5b50813567ffffffffffffffff81111561117957600080fd5b60208301915083602082850101111561119157600080fd5b9250929050565b60008060008060008060a087890312156111b157600080fd5b6111ba87611126565b955060208701359450604087013567ffffffffffffffff8111156111dd57600080fd5b6111e989828a0161114f565b90955093505060608701359150608087013567ffffffffffffffff81111561121057600080fd5b87016060818a03121561122257600080fd5b809150509295509295509295565b60008060006060848603121561124557600080fd5b61124e84611126565b95602085013595506040909401359392505050565b60006020828403121561127557600080fd5b5035919050565b6000806040838503121561128f57600080fd5b8235915061129f60208401611126565b90509250929050565b6000806000806000608086880312156112c057600080fd5b6112c986611126565b945060208601359350604086013567ffffffffffffffff8111156112ec57600080fd5b6112f88882890161114f565b96999598509660600135949350505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff61137182611126565b1682526000602082013567ffffffffffffffff811680821461139257600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126113ce57600080fd5b820160208101903567ffffffffffffffff8111156113eb57600080fd5b8036038213156113fa57600080fd5b6060604086015261140f60608601828461130a565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a06060820152600061146a60a08301858761130a565b828103608084015261147c8185611353565b9998505050505050505050565b8481526060602082015260006114a360608301858761130a565b82810360408401526114b58185611353565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006114b560808301848661130a565b83815260406020820152600061140f60408301848661130a565b60006020828403121561153e57600080fd5b81518015158114610f4957600080fd5b6000825160005b8181101561156f5760208186018101518583015201611555565b50600092019182525091905056fea26469706673582212208568eb5447d9d0fee399f3e161164b36a7f43cc5bc13eb5835e3c9f436b2697664736f6c634300081a00336080604052348015600f57600080fd5b506001600055610ec3806100246000396000f3fe60806040526004361061006e5760003560e01c80636ed701691161004b5780636ed70169146100e0578063c5131691146100f5578063e04d4f9714610115578063f05b6abf1461012857005b8063357fc5a214610077578063660b9de014610097578063676cc054146100b757005b3661007557005b005b34801561008357600080fd5b50610075610092366004610724565b610148565b3480156100a357600080fd5b506100756100b2366004610760565b6101de565b6100ca6100c536600461079b565b61021a565b6040516100d79190610896565b60405180910390f35b3480156100ec57600080fd5b5061007561024d565b34801561010157600080fd5b50610075610110366004610724565b610282565b6100756101233660046109ce565b61035d565b34801561013457600080fd5b50610075610143366004610aba565b6103a1565b6101506103d6565b61017273ffffffffffffffffffffffffffffffffffffffff8316338386610419565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101d96001600055565b505050565b7f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e338260405161020f929190610bed565b60405180910390a150565b6040516060907f3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee90600090a19392505050565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61028a6103d6565b6000610297600285610ccc565b9050806000036102d3576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102f573ffffffffffffffffffffffffffffffffffffffff8416338484610419565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101d96001600055565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa3334858585604051610394959493929190610d07565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146338484846040516103949493929190610d91565b600260005403610412576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526104ae9085906104b4565b50505050565b60006104d673ffffffffffffffffffffffffffffffffffffffff84168361054f565b905080516000141580156104fb5750808060200190518101906104f99190610e54565b155b156101d9576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b606061055d83836000610564565b9392505050565b6060814710156105a2576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610546565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105cb9190610e71565b60006040518083038185875af1925050503d8060008114610608576040519150601f19603f3d011682016040523d82523d6000602084013e61060d565b606091505b509150915061061d868383610627565b9695505050505050565b60608261063c57610637826106b6565b61055d565b8151158015610660575073ffffffffffffffffffffffffffffffffffffffff84163b155b156106af576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610546565b508061055d565b8051156106c65780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff8116811461071f57600080fd5b919050565b60008060006060848603121561073957600080fd5b83359250610749602085016106fb565b9150610757604085016106fb565b90509250925092565b60006020828403121561077257600080fd5b813567ffffffffffffffff81111561078957600080fd5b82016060818503121561055d57600080fd5b600080600083850360408112156107b157600080fd5b60208112156107bf57600080fd5b50839250602084013567ffffffffffffffff8111156107dd57600080fd5b8401601f810186136107ee57600080fd5b803567ffffffffffffffff81111561080557600080fd5b86602082840101111561081757600080fd5b939660209190910195509293505050565b60005b8381101561084357818101518382015260200161082b565b50506000910152565b60008151808452610864816020860160208601610828565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061055d602083018461084c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561091f5761091f6108a9565b604052919050565b600082601f83011261093857600080fd5b813567ffffffffffffffff811115610952576109526108a9565b61098360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016108d8565b81815284602083860101111561099857600080fd5b816020850160208301376000918101602001919091529392505050565b80151581146106f857600080fd5b803561071f816109b5565b6000806000606084860312156109e357600080fd5b833567ffffffffffffffff8111156109fa57600080fd5b610a0686828701610927565b935050602084013591506040840135610a1e816109b5565b809150509250925092565b600067ffffffffffffffff821115610a4357610a436108a9565b5060051b60200190565b600082601f830112610a5e57600080fd5b8135610a71610a6c82610a29565b6108d8565b8082825260208201915060208360051b860101925085831115610a9357600080fd5b602085015b83811015610ab0578035835260209283019201610a98565b5095945050505050565b600080600060608486031215610acf57600080fd5b833567ffffffffffffffff811115610ae657600080fd5b8401601f81018613610af757600080fd5b8035610b05610a6c82610a29565b8082825260208201915060208360051b850101925088831115610b2757600080fd5b602084015b83811015610b6957803567ffffffffffffffff811115610b4b57600080fd5b610b5a8b602083890101610927565b84525060209283019201610b2c565b509550505050602084013567ffffffffffffffff811115610b8957600080fd5b610b9586828701610a4d565b925050610757604085016109c3565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610c2b836106fb565b1660408201526000602083013567ffffffffffffffff8116808214610c4f57600080fd5b6060840152506040830135368490037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112610c8b57600080fd5b830160208101903567ffffffffffffffff811115610ca857600080fd5b803603821315610cb757600080fd5b6060608085015261061d60a085018284610ba4565b600082610d02577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610d3c60a083018661084c565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610d87578151865260209586019590910190600101610d69565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610e24577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610e0f85835161084c565b94506020938401939190910190600101610dd5565b505050508281036040840152610e3a8186610d55565b915050610e4b606083018415159052565b95945050505050565b600060208284031215610e6657600080fd5b815161055d816109b5565b60008251610e83818460208701610828565b919091019291505056fea2646970667358221220d537f023478b5746f5f67760e722db6f4a51db89700ac69d6ba33759b8ac342f64736f6c634300081a0033a26469706673582212202ae40bba2840cd77ea3a08d624bc18c4c9c30c264ee29671e8fc48518796931b64736f6c634300081a0033",
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

// ZetaConnectorNativeTestReceivedOnCallIterator is returned from FilterReceivedOnCall and is used to iterate over the raw logs and unpacked data for ReceivedOnCall events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedOnCallIterator struct {
	Event *ZetaConnectorNativeTestReceivedOnCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaConnectorNativeTestReceivedOnCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReceivedOnCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaConnectorNativeTestReceivedOnCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaConnectorNativeTestReceivedOnCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestReceivedOnCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReceivedOnCall represents a ReceivedOnCall event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedOnCall struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCall is a free log retrieval operation binding the contract event 0x3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee.
//
// Solidity: event ReceivedOnCall()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReceivedOnCall(opts *bind.FilterOpts) (*ZetaConnectorNativeTestReceivedOnCallIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestReceivedOnCallIterator{contract: _ZetaConnectorNativeTest.contract, event: "ReceivedOnCall", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCall is a free log subscription operation binding the contract event 0x3658b46bab672c7672b69c2f0feda706eabdb7d2231421c96e9049b2db5e7eee.
//
// Solidity: event ReceivedOnCall()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReceivedOnCall(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReceivedOnCall) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReceivedOnCall)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReceivedOnCall(log types.Log) (*ZetaConnectorNativeTestReceivedOnCall, error) {
	event := new(ZetaConnectorNativeTestReceivedOnCall)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
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
