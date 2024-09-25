// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayevmzevm

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

// GatewayEVMZEVMTestMetaData contains all meta data concerning the GatewayEVMZEVMTest contract.
var GatewayEVMZEVMTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testCallReceiverEVMFromSenderZEVM\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallReceiverEVMFromZEVM\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiverEVMFromSenderZEVM\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiverEVMFromZEVM\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5062011d30806200003e6000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806385226c81116100b2578063b5508aa911610081578063d7a525fc11610066578063d7a525fc146101ec578063e20c9f71146101f4578063fa7626d4146101fc57600080fd5b8063b5508aa9146101cc578063ba414fa6146101d457600080fd5b806385226c8114610192578063916a17c6146101a75780639683c695146101bc578063b0464fdc146101c457600080fd5b80633f7286f4116100ee5780633f7286f414610165578063524744131461016d57806366d9a9a0146101755780636ff15ccc1461018a57600080fd5b80630a9254e4146101205780631ed7831c1461012a5780632ade3880146101485780633e5e3c231461015d575b600080fd5b610128610209565b005b61013261126e565b60405161013f9190617af4565b60405180910390f35b6101506112d0565b60405161013f9190617b90565b610132611412565b610132611472565b6101286114d2565b61017d611d41565b60405161013f9190617cf6565b610128611ec3565b61019a6126e6565b60405161013f9190617d94565b6101af6127b6565b60405161013f9190617e0b565b6101286128b1565b6101af612ea5565b61019a612fa0565b6101dc613070565b604051901515815260200161013f565b610128613144565b610132613868565b601f546101dc9060ff1681565b602680547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602780548216611234179055602880548216615678179055602e8054909116614321179055604051610267906179d9565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156102ec573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055604051610331906179d9565b604080825260049082018190527f7a6574610000000000000000000000000000000000000000000000000000000060608301526080602083018190528201527f5a4554410000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156103b5573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081178255604080518082018252600e81527f4761746577617945564d2e736f6c0000000000000000000000000000000000006020820152602854602654925190861694810194909452604484019290925290921660648201526104a6919060840160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b000000000000000000000000000000000000000000000000000000001790526138c8565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff00000000000000000000000000000000000000009091168117909155602854602654604051929391821692911690610532906179e7565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f08015801561056e573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392831617905560205460245460285460265460405193851694928316939183169216906105c9906179f5565b6001600160a01b039485168152928416602084015290831660408301529091166060820152608001604051809103906000f08015801561060d573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556028546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b1580156106b957600080fd5b505af11580156106cd573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561074357600080fd5b505af1158015610757573d6000803e3d6000fd5b50506020546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f9150602401600060405180830381600087803b1580156107bd57600080fd5b505af11580156107d1573d6000803e3d6000fd5b50506020546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef9150602401600060405180830381600087803b15801561083757600080fd5b505af115801561084b573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156108ad57600080fd5b505af11580156108c1573d6000803e3d6000fd5b50506023546026546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42406024820152911692506340c10f199150604401600060405180830381600087803b15801561093057600080fd5b505af1158015610944573d6000803e3d6000fd5b50506023546021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526207a12060248201529116925063a9059cbb91506044016020604051808303816000875af11580156109b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109dc9190617ea2565b506040516109e990617a03565b604051809103906000f080158015610a05573d6000803e3d6000fd5b50602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316179055604080518082018252600f81527f476174657761795a45564d2e736f6c0000000000000000000000000000000000602082015260248054602e54935190851691810191909152919092166044820152610aee919060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f485cc955000000000000000000000000000000000000000000000000000000001790526138c8565b602980546001600160a01b03929092167fffffffffffffffffffffffff00000000000000000000000000000000000000009283168117909155602a80549092168117909155604051610b3f90617a11565b6001600160a01b039091168152602001604051809103906000f080158015610b6b573d6000803e3d6000fd5b50602b80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03929092169190911790556040517f06447d5600000000000000000000000000000000000000000000000000000000815273735b14bb79463307aacbed86daf3322b1e6226ab6004820181905290737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015610c2057600080fd5b505af1158015610c34573d6000803e3d6000fd5b505050506000806000604051610c4990617a1f565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610c85573d6000803e3d6000fd5b50602c80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602a54604051601293600193600093849391921690610cdb90617a2d565b610cea96959493929190617ec4565b604051809103906000f080158015610d06573d6000803e3d6000fd5b50602d80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602c546040517fee2815ba0000000000000000000000000000000000000000000000000000000081526001600482015260248101929092529091169063ee2815ba90604401600060405180830381600087803b158015610d9d57600080fd5b505af1158015610db1573d6000803e3d6000fd5b5050602c546040517fa7cb050700000000000000000000000000000000000000000000000000000000815260016004820181905260248201526001600160a01b03909116925063a7cb05079150604401600060405180830381600087803b158015610e1b57600080fd5b505af1158015610e2f573d6000803e3d6000fd5b5050602d54602e546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42406024820152911692506347e7ef2491506044016020604051808303816000875af1158015610ea3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec79190617ea2565b50602d54602b546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424060248201529116906347e7ef24906044016020604051808303816000875af1158015610f38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5c9190617ea2565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610fbb57600080fd5b505af1158015610fcf573d6000803e3d6000fd5b5050602e546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561104557600080fd5b505af1158015611059573d6000803e3d6000fd5b5050602d54602a546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424060248201529116925063095ea7b391506044016020604051808303816000875af11580156110cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110f19190617ea2565b506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b15801561117257600080fd5b505af1158015611186573d6000803e3d6000fd5b50506040805160a08101825261032180825260016020808401918252838501928352845190810190945260008085526060840185905260808401528251602f80549251151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009093166001600160a01b0392831617929092178255915160308054919093167fffffffffffffffffffffffff00000000000000000000000000000000000000009190911617909155909350915060319061125d9082618084565b506080820151816003015590505050565b606060168054806020026020016040519081016040528092919081815260200182805480156112c657602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116112a8575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020016000905b8282101561140957600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156113f257838290600052602060002001805461136590617fe8565b80601f016020809104026020016040519081016040528092919081815260200182805461139190617fe8565b80156113de5780601f106113b3576101008083540402835291602001916113de565b820191906000526020600020905b8154815290600101906020018083116113c157829003601f168201915b505050505081526020019060010190611346565b5050505081525050815260200190600101906112f4565b50505050905090565b606060188054806020026020016040519081016040528092919081815260200182805480156112c6576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116112a8575050505050905090565b606060178054806020026020016040519081016040528092919081815260200182805480156112c6576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116112a8575050505050905090565b604080518082018252600681527f48656c6c6f2100000000000000000000000000000000000000000000000000006020820152602d54602b5492517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0393841660048201529192602a92600192670de0b6b3a7640000926000929116906370a0823190602401602060405180830381865afa15801561157e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115a29190618143565b6040519091506000907fe04d4f9700000000000000000000000000000000000000000000000000000000906115df9088908890889060240161815c565b60408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909516949094179093526025549051919350600092611678926001600160a01b03909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602d546116af926207a120916001600160a01b0316908690600190602f90602401618275565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f048ae42c00000000000000000000000000000000000000000000000000000000179052602a5490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161176c916001600160a01b03919091169060009086906004016182de565b600060405180830381600087803b15801561178657600080fd5b505af115801561179a573d6000803e3d6000fd5b5050602e546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561181057600080fd5b505af1158015611824573d6000803e3d6000fd5b5050602b5460255460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039091169250630abd8905915060340160408051601f1981840301815290829052602d547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526118be926207a120916001600160a01b0316908d908d908d90600401618306565b600060405180830381600087803b1580156118d857600080fd5b505af11580156118ec573d6000803e3d6000fd5b50506020546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260248101879052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b15801561196957600080fd5b505af115801561197d573d6000803e3d6000fd5b50506025546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015611a0f57600080fd5b505af1158015611a23573d6000803e3d6000fd5b50506020546040517f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa9350611a6c92506001600160a01b039091169087908b908b908b9061835b565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611b0257600080fd5b505af1158015611b16573d6000803e3d6000fd5b50506025546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f9150611b5b908790869061839c565b60405180910390a26028546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611bd557600080fd5b505af1158015611be9573d6000803e3d6000fd5b50506020546025546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd93508892611c3d92169087906004016183b5565b60006040518083038185885af1158015611c5b573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052611c849190810190618490565b50602d54602b546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611cef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d139190618143565b9050611d37816001611d286207a120886184f4565b611d3291906184f4565b6138e7565b5050505050505050565b6060601b805480602002602001604051908101604052809291908181526020016000905b828210156114095783829060005260206000209060020201604051806040016040529081600082018054611d9890617fe8565b80601f0160208091040260200160405190810160405280929190818152602001828054611dc490617fe8565b8015611e115780601f10611de657610100808354040283529160200191611e11565b820191906000526020600020905b815481529060010190602001808311611df457829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015611eab57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411611e585790505b50505050508152505081526020019060010190611d65565b604080518082018252600681527f48656c6c6f21000000000000000000000000000000000000000000000000000060208201529051602a90600190670de0b6b3a7640000906000907fe04d4f970000000000000000000000000000000000000000000000000000000090611f3f9087908790879060240161815c565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009490941693909317909252602a5491517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190526024820181905260448201819052606482018190526001600160a01b039093166084820152909250737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561203157600080fd5b505af1158015612045573d6000803e3d6000fd5b5050602e5460255460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507fda1215b0949ddb309fe466fa9e70e861a16538f11b8ecdb05c217d4d8677ed2d9060340160408051601f19818403018152828252602d547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b03909216916207a1209188918491634d8943bb916004808201926020929091908290030181865afa15801561211a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061213e9190618143565b896001602f604051612157989796959493929190618507565b60405180910390a3602e546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156121d157600080fd5b505af11580156121e5573d6000803e3d6000fd5b5050602a5460255460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b03909116925063048ae42c915060340160408051601f1981840301815290829052602d547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352612281926207a120916001600160a01b0316908890600190602f90600401618569565b600060405180830381600087803b15801561229b57600080fd5b505af11580156122af573d6000803e3d6000fd5b5050602d54602e546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561231b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061233f9190618143565b905061235281611d32846207a1206184f4565b6020546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260248101859052737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b1580156123cb57600080fd5b505af11580156123df573d6000803e3d6000fd5b50506025546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561247157600080fd5b505af1158015612485573d6000803e3d6000fd5b50506020546040517f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa93506124ce92506001600160a01b039091169087908b908b908b9061835b565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561256457600080fd5b505af1158015612578573d6000803e3d6000fd5b50506025546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f91506125bd908790879061839c565b60405180910390a26028546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561263757600080fd5b505af115801561264b573d6000803e3d6000fd5b50506020546025546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd9350889261269f92169088906004016183b5565b60006040518083038185885af11580156126bd573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052611d379190810190618490565b6060601a805480602002602001604051908101604052809291908181526020016000905b8282101561140957838290600052602060002001805461272990617fe8565b80601f016020809104026020016040519081016040528092919081815260200182805461275590617fe8565b80156127a25780601f10612777576101008083540402835291602001916127a2565b820191906000526020600020905b81548152906001019060200180831161278557829003601f168201915b50505050508152602001906001019061270a565b6060601d805480602002602001604051908101604052809291908181526020016000905b828210156114095760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561289957602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116128465790505b505050505081525050815260200190600101906127da565b604080518082018252600681527f48656c6c6f21000000000000000000000000000000000000000000000000000060208201529051602a90600190670de0b6b3a7640000906000907fe04d4f97000000000000000000000000000000000000000000000000000000009061292d9087908790879060240161815c565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009490941693909317909252602e5491517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b0390921660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612a0057600080fd5b505af1158015612a14573d6000803e3d6000fd5b5050602a546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015612aa657600080fd5b505af1158015612aba573d6000803e3d6000fd5b5050602d54602e5460255460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039283169450911691507f6c6abd640fc6a0ef7cf2bc54b246b42d5c2629c30be1e24fea4a58157a7728cf9060340160408051601f1981840301815290829052612b3a918690600190602f906185bd565b60405180910390a3602a546025546040805160609290921b6bffffffffffffffffffffffff19166020830152805180830360140181526034830191829052602d547f1cb5ea75000000000000000000000000000000000000000000000000000000009092526001600160a01b0393841693631cb5ea7593612bca9391909116908690600190602f906038016185fc565b600060405180830381600087803b158015612be457600080fd5b505af1158015612bf8573d6000803e3d6000fd5b50506020546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260248101859052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b158015612c7557600080fd5b505af1158015612c89573d6000803e3d6000fd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015612d1b57600080fd5b505af1158015612d2f573d6000803e3d6000fd5b50506025546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f9150612d74908590859061839c565b60405180910390a26028546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612dee57600080fd5b505af1158015612e02573d6000803e3d6000fd5b50506020546025546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd93508692612e5692169086906004016183b5565b60006040518083038185885af1158015612e74573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052612e9d9190810190618490565b505050505050565b6060601c805480602002602001604051908101604052809291908181526020016000905b828210156114095760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015612f8857602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411612f355790505b50505050508152505081526020019060010190612ec9565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015611409578382906000526020600020018054612fe390617fe8565b80601f016020809104026020016040519081016040528092919081815260200182805461300f90617fe8565b801561305c5780601f106130315761010080835404028352916020019161305c565b820191906000526020600020905b81548152906001019060200180831161303f57829003601f168201915b505050505081526020019060010190612fc4565b60085460009060ff1615613088575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015613119573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061313d9190618143565b1415905090565b604080518082018252600681527f48656c6c6f21000000000000000000000000000000000000000000000000000060208201529051602a90600190670de0b6b3a7640000906000907fe04d4f9700000000000000000000000000000000000000000000000000000000906131c09087908790879060240161815c565b60408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909516949094179093526025549051919350600092613259926001600160a01b03909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602d5461328d926001600160a01b03909116908590600190602f90602401618656565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1cb5ea7500000000000000000000000000000000000000000000000000000000179052602a5490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161334a916001600160a01b03919091169060009086906004016182de565b600060405180830381600087803b15801561336457600080fd5b505af1158015613378573d6000803e3d6000fd5b5050602e546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156133ee57600080fd5b505af1158015613402573d6000803e3d6000fd5b5050602b5460255460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039091169250637a34d8bb915060340160408051601f1981840301815290829052602d547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352613499926001600160a01b03909116908b908b908b906004016186a7565b600060405180830381600087803b1580156134b357600080fd5b505af11580156134c7573d6000803e3d6000fd5b50506020546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260248101869052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b15801561354457600080fd5b505af1158015613558573d6000803e3d6000fd5b50506025546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156135ea57600080fd5b505af11580156135fe573d6000803e3d6000fd5b50506020546040517f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa935061364792506001600160a01b039091169086908a908a908a9061835b565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156136dd57600080fd5b505af11580156136f1573d6000803e3d6000fd5b50506025546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f9150613736908690869061839c565b60405180910390a26028546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156137b057600080fd5b505af11580156137c4573d6000803e3d6000fd5b50506020546025546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd9350879261381892169087906004016183b5565b60006040518083038185885af1158015613836573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f1916820160405261385f9190810190618490565b50505050505050565b606060158054806020026020016040519081016040528092919081815260200182805480156112c6576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116112a8575050505050905090565b60006138d2617a3b565b6138dd848483613966565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c549060440160006040518083038186803b15801561395257600080fd5b505afa158015612e9d573d6000803e3d6000fd5b60008061397385846139e1565b90506139d66040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016139c19291906183b5565b604051602081830303815290604052856139ed565b9150505b9392505050565b60006139da8383613a1b565b60c08101515160009015613a1157613a0a84848460c00151613a36565b90506139da565b613a0a8484613bdc565b6000613a278383613cc7565b6139da838360200151846139ed565b600080613a41613cd7565b90506000613a4f8683613daa565b90506000613a668260600151836020015185614250565b90506000613a7683838989614462565b90506000613a83826152df565b602081015181519192509060030b15613af657898260400151604051602001613aad9291906186f5565b60408051601f19818403018152908290527f08c379a0000000000000000000000000000000000000000000000000000000008252613aed91600401618776565b60405180910390fd5b6000613b396040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a2000000000000000000000008152508360016154ae565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d90613b8c908490600401618776565b602060405180830381865afa158015613ba9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613bcd9190618789565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc92590613c31908790600401618776565b600060405180830381865afa158015613c4e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613c769190810190618490565b90506000613ca48285604051602001613c909291906187b2565b6040516020818303038152906040526156ae565b90506001600160a01b0381166138dd578484604051602001613aad9291906187e1565b613cd3828260006156c1565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c90613d5e90849060040161888c565b600060405180830381865afa158015613d7b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613da391908101906188d3565b9250505090565b613ddc6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d9050613e276040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b613e30856157c4565b60208201526000613e4086615ba9565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa158015613e82573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613eaa91908101906188d3565b86838560200151604051602001613ec4949392919061891c565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb1190613f1c908590600401618776565b600060405180830381865afa158015613f39573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613f6191908101906188d3565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f690613fa9908490600401618a20565b602060405180830381865afa158015613fc6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613fea9190617ea2565b613fff5781604051602001613aad9190618a72565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890614044908490600401618b04565b600060405180830381865afa158015614061573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261408991908101906188d3565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906140d0908490600401618b56565b602060405180830381865afa1580156140ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141119190617ea2565b156141a6576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061415b908490600401618b56565b600060405180830381865afa158015614178573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526141a091908101906188d3565b60408501525b846001600160a01b03166349c4fac88286600001516040516020016141cb9190618ba8565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016141f7929190618c14565b600060405180830381865afa158015614214573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261423c91908101906188d3565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b606081526020019060019003908161426c5790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250816000815181106142cc576142cc618c39565b60200260200101819052506040518060400160405280600381526020017f2d726c00000000000000000000000000000000000000000000000000000000008152508160018151811061432057614320618c39565b60200260200101819052508460405160200161433c9190618c68565b6040516020818303038152906040528160028151811061435e5761435e618c39565b60200260200101819052508260405160200161437a9190618cd4565b6040516020818303038152906040528160038151811061439c5761439c618c39565b602002602001018190525060006143b2826152df565b602080820151604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000081850190815282518084018452600080825290860152825180840190935290518252928101929092529192506144439060408051808201825260008082526020918201528151808301909252845182528085019082015290615e2c565b6144585785604051602001613aad9190618d15565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d90156144b2565b511590565b6146265782602001511561456e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401613aed565b8260c0015115614626576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401613aed565b6040805160ff8082526120008201909252600091816020015b606081526020019060019003908161463f57905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061469a90618da6565b935060ff16815181106146af576146af618c39565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e37000000000000000000000000000000000000008152506040516020016147009190618dc5565b60405160208183030381529060405282828061471b90618da6565b935060ff168151811061473057614730618c39565b60200260200101819052506040518060400160405280600681526020017f6465706c6f79000000000000000000000000000000000000000000000000000081525082828061477d90618da6565b935060ff168151811061479257614792618c39565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806147df90618da6565b935060ff16815181106147f4576147f4618c39565b6020026020010181905250876020015182828061481090618da6565b935060ff168151811061482557614825618c39565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163745061746800000000000000000000000000000000000081525082828061487290618da6565b935060ff168151811061488757614887618c39565b60209081029190910101528751828261489f81618da6565b935060ff16815181106148b4576148b4618c39565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e4964000000000000000000000000000000000000000000000081525082828061490190618da6565b935060ff168151811061491657614916618c39565b602002602001018190525061492a46615e8d565b828261493581618da6565b935060ff168151811061494a5761494a618c39565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c65000000000000000000000000000000000081525082828061499790618da6565b935060ff16815181106149ac576149ac618c39565b6020026020010181905250868282806149c490618da6565b935060ff16815181106149d9576149d9618c39565b6020908102919091010152855115614b005760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f6465000000000000000000000060208201528282614a2a81618da6565b935060ff1681518110614a3f57614a3f618c39565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d90614a8f908990600401618776565b600060405180830381865afa158015614aac573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052614ad491908101906188d3565b8282614adf81618da6565b935060ff1681518110614af457614af4618c39565b60200260200101819052505b846020015115614bd05760408051808201909152601281527f2d2d766572696679536f75726365436f6465000000000000000000000000000060208201528282614b4981618da6565b935060ff1681518110614b5e57614b5e618c39565b60200260200101819052506040518060400160405280600581526020017f66616c7365000000000000000000000000000000000000000000000000000000815250828280614bab90618da6565b935060ff1681518110614bc057614bc0618c39565b6020026020010181905250614d97565b614c086144ad8660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b614c9b5760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282614c4b81618da6565b935060ff1681518110614c6057614c60618c39565b60200260200101819052508460a00151604051602001614c809190618c68565b604051602081830303815290604052828280614bab90618da6565b8460c00151158015614cde575060408089015181518083018352600080825260209182015282518084019093528151835290810190820152614cdc90511590565b155b15614d975760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282614d2281618da6565b935060ff1681518110614d3757614d37618c39565b6020026020010181905250614d4b88615f2d565b604051602001614d5b9190618c68565b604051602081830303815290604052828280614d7690618da6565b935060ff1681518110614d8b57614d8b618c39565b60200260200101819052505b60408086015181518083018352600080825260209182015282518084019093528151835290810190820152614dcb90511590565b614e605760408051808201909152600b81527f2d2d72656c61796572496400000000000000000000000000000000000000000060208201528282614e0e81618da6565b935060ff1681518110614e2357614e23618c39565b60200260200101819052508460400151828280614e3f90618da6565b935060ff1681518110614e5457614e54618c39565b60200260200101819052505b606085015115614f815760408051808201909152600681527f2d2d73616c74000000000000000000000000000000000000000000000000000060208201528282614ea981618da6565b935060ff1681518110614ebe57614ebe618c39565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015614f2d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052614f5591908101906188d3565b8282614f6081618da6565b935060ff1681518110614f7557614f75618c39565b60200260200101819052505b60e085015151156150285760408051808201909152600a81527f2d2d6761734c696d69740000000000000000000000000000000000000000000060208201528282614fcb81618da6565b935060ff1681518110614fe057614fe0618c39565b6020026020010181905250614ffc8560e0015160000151615e8d565b828261500781618da6565b935060ff168151811061501c5761501c618c39565b60200260200101819052505b60e085015160200151156150d25760408051808201909152600a81527f2d2d6761735072696365000000000000000000000000000000000000000000006020820152828261507581618da6565b935060ff168151811061508a5761508a618c39565b60200260200101819052506150a68560e0015160200151615e8d565b82826150b181618da6565b935060ff16815181106150c6576150c6618c39565b60200260200101819052505b60e0850151604001511561517c5760408051808201909152600e81527f2d2d6d61784665655065724761730000000000000000000000000000000000006020820152828261511f81618da6565b935060ff168151811061513457615134618c39565b60200260200101819052506151508560e0015160400151615e8d565b828261515b81618da6565b935060ff168151811061517057615170618c39565b60200260200101819052505b60e085015160600151156152265760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826151c981618da6565b935060ff16815181106151de576151de618c39565b60200260200101819052506151fa8560e0015160600151615e8d565b828261520581618da6565b935060ff168151811061521a5761521a618c39565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561524457615244617fb9565b60405190808252806020026020018201604052801561527757816020015b60608152602001906001900390816152625790505b50905060005b8260ff168160ff1610156152d057838160ff16815181106152a0576152a0618c39565b6020026020010151828260ff16815181106152bd576152bd618c39565b602090810291909101015260010161527d565b5093505050505b949350505050565b6153066040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c9161538c91869101618e30565b600060405180830381865afa1580156153a9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526153d191908101906188d3565b905060006153df8683616a1c565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b815260040161540f9190617d94565b6000604051808303816000875af115801561542e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526154569190810190618e77565b805190915060030b1580159061546f5750602081015151155b801561547e5750604081015151155b15614458578160008151811061549657615496618c39565b6020026020010151604051602001613aad9190618f2d565b606060006154e38560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252865182528087019082015290915061551a9082905b90616b71565b15615677576000615597826155918461558b61555d8a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b90616b98565b90616bfa565b604080518082018252600181527f0a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506155fb908290616b71565b1561566557604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615662905b8290616c7f565b90505b61566e81616ca5565b925050506139da565b8215615690578484604051602001613aad929190619119565b50506040805160208101909152600081526139da565b509392505050565b6000808251602084016000f09392505050565b8160a00151156156d057505050565b60006156dd848484616d0e565b905060006156ea826152df565b602081015181519192509060030b1580156157865750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261578690604080518082018252600080825260209182015281518083019092528451825280850190820152615514565b1561579357505050505050565b604082015151156157b3578160400151604051602001613aad91906191c0565b80604051602001613aad919061921e565b606060006157f98360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061585e905b8290615e2c565b156158cd57604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526139da906158c89083906172a9565b616ca5565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261592f905b8290617333565b6001036159fc57604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526159959061565b565b50604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526139da906158c8905b8390616c7f565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615a5b90615857565b15615b9257604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290615ac39083906173cd565b905060008160018351615ad691906184f4565b81518110615ae657615ae6618c39565b60200260200101519050615b896158c8615b5c6040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528551825280860190820152906172a9565b95945050505050565b82604051602001613aad9190619289565b50919050565b60606000615bde8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150615c4090615857565b15615c4e576139da81616ca5565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615cad90615928565b600103615d1757604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526139da906158c8906159f5565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615d7690615857565b15615b9257604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290615dde9083906173cd565b9050600181511115615e1a578060028251615df991906184f4565b81518110615e0957615e09618c39565b602002602001015192505050919050565b5082604051602001613aad9190619289565b805182516000911115615e41575060006138e1565b81518351602085015160009291615e5791619367565b615e6191906184f4565b905082602001518103615e785760019150506138e1565b82516020840151819020912014905092915050565b60606000615e9a83617472565b600101905060008167ffffffffffffffff811115615eba57615eba617fb9565b6040519080825280601f01601f191660200182016040528015615ee4576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084615eee57509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e5345440000000000000000000000000000000000000000000081840190815285518087018752838152840192909252845180860190955251845290830152606091615fb9905b8290617554565b15615ff957505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261605890615fb2565b1561609857505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d49540000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526160f790615fb2565b1561613757505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261619690615fb2565b806161fb5750604080518082018252601081527f47504c2d322e302d6f722d6c6174657200000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526161fb90615fb2565b1561623b57505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261629a90615fb2565b806162ff5750604080518082018252601081527f47504c2d332e302d6f722d6c6174657200000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526162ff90615fb2565b1561633f57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261639e90615fb2565b806164035750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261640390615fb2565b1561644357505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526164a290615fb2565b806165075750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261650790615fb2565b1561654757505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c617573650000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526165a690615fb2565b156165e657505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261664590615fb2565b1561668557505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e3000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526166e490615fb2565b1561672457505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261678390615fb2565b156167c357505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261682290615fb2565b1561686257505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526168c190615fb2565b806169265750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261692690615fb2565b1561696657505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e31000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526169c590615fb2565b15616a0557505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151613aad929060200161937a565b60608060005b8451811015616aa75781858281518110616a3e57616a3e618c39565b6020026020010151604051602001616a579291906187b2565b604051602081830303815290604052915060018551616a7691906184f4565b8114616a9f5781604051602001616a8d91906194e3565b60405160208183030381529060405291505b600101616a22565b5060408051600380825260808201909252600091816020015b6060815260200190600190039081616ac05790505090508381600081518110616aeb57616aeb618c39565b60200260200101819052506040518060400160405280600281526020017f2d6300000000000000000000000000000000000000000000000000000000000081525081600181518110616b3f57616b3f618c39565b60200260200101819052508181600281518110616b5e57616b5e618c39565b6020908102919091010152949350505050565b6020808301518351835192840151600093616b8f9291849190617568565b14159392505050565b60408051808201909152600080825260208201526000616bca8460000151856020015185600001518660200151617679565b9050836020015181616bdc91906184f4565b84518590616beb9083906184f4565b90525060208401525090919050565b6040805180820190915260008082526020820152815183511015616c1f5750816138e1565b6020808301519084015160019114616c465750815160208481015190840151829020919020145b8015616c7757825184518590616c5d9083906184f4565b9052508251602085018051616c73908390619367565b9052505b509192915050565b6040805180820190915260008082526020820152616c9e838383617799565b5092915050565b60606000826000015167ffffffffffffffff811115616cc657616cc6617fb9565b6040519080825280601f01601f191660200182016040528015616cf0576020820181803683370190505b5090506000602082019050616c9e8185602001518660000151617844565b60606000616d1a613cd7565b6040805160ff808252612000820190925291925060009190816020015b6060815260200190600190039081616d3757905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280616d9290618da6565b935060ff1681518110616da757616da7618c39565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e3300000000000000000000000000000000000000000000000000815250604051602001616df89190619524565b604051602081830303815290604052828280616e1390618da6565b935060ff1681518110616e2857616e28618c39565b60200260200101819052506040518060400160405280600881526020017f76616c6964617465000000000000000000000000000000000000000000000000815250828280616e7590618da6565b935060ff1681518110616e8a57616e8a618c39565b602002602001018190525082604051602001616ea69190618cd4565b604051602081830303815290604052828280616ec190618da6565b935060ff1681518110616ed657616ed6618c39565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e747261637400000000000000000000000000000000000000000000815250828280616f2390618da6565b935060ff1681518110616f3857616f38618c39565b6020026020010181905250616f4d87846178be565b8282616f5881618da6565b935060ff1681518110616f6d57616f6d618c39565b6020908102919091010152855151156170195760408051808201909152600b81527f2d2d7265666572656e636500000000000000000000000000000000000000000060208201528282616fbf81618da6565b935060ff1681518110616fd457616fd4618c39565b6020026020010181905250616fed8660000151846178be565b8282616ff881618da6565b935060ff168151811061700d5761700d618c39565b60200260200101819052505b8560800151156170875760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261706281618da6565b935060ff168151811061707757617077618c39565b60200260200101819052506170ed565b84156170ed5760408051808201909152601281527f2d2d726571756972655265666572656e63650000000000000000000000000000602082015282826170cc81618da6565b935060ff16815181106170e1576170e1618c39565b60200260200101819052505b604086015151156171895760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261713781618da6565b935060ff168151811061714c5761714c618c39565b6020026020010181905250856040015182828061716890618da6565b935060ff168151811061717d5761717d618c39565b60200260200101819052505b8560600151156171f35760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826171d281618da6565b935060ff16815181106171e7576171e7618c39565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561721157617211617fb9565b60405190808252806020026020018201604052801561724457816020015b606081526020019060019003908161722f5790505b50905060005b8260ff168160ff16101561729d57838160ff168151811061726d5761726d618c39565b6020026020010151828260ff168151811061728a5761728a618c39565b602090810291909101015260010161724a565b50979650505050505050565b60408051808201909152600080825260208201528151835110156172ce5750816138e1565b815183516020850151600092916172e491619367565b6172ee91906184f4565b6020840151909150600190821461730f575082516020840151819020908220145b801561732a578351855186906173269083906184f4565b9052505b50929392505050565b60008082600001516173578560000151866020015186600001518760200151617679565b6173619190619367565b90505b835160208501516173759190619367565b8111616c9e578161738581619569565b92505082600001516173bc8560200151836173a091906184f4565b86516173ac91906184f4565b8386600001518760200151617679565b6173c69190619367565b9050617364565b606060006173db8484617333565b6173e6906001619367565b67ffffffffffffffff8111156173fe576173fe617fb9565b60405190808252806020026020018201604052801561743157816020015b606081526020019060019003908161741c5790505b50905060005b81518110156156a65761744d6158c88686616c7f565b82828151811061745f5761745f618c39565b6020908102919091010152600101617437565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106174bb577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106174e7576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061750557662386f26fc10000830492506010015b6305f5e100831061751d576305f5e100830492506008015b612710831061753157612710830492506004015b60648310617543576064830492506002015b600a83106138e15760010192915050565b600061756083836178fe565b159392505050565b60008085841161766f576020841161761b57600084156175b357600161758f8660206184f4565b61759a906008619583565b6175a5906002619681565b6175af91906184f4565b1990505b83518116856175c28989619367565b6175cc91906184f4565b805190935082165b818114617606578784116175ee57879450505050506152d7565b836175f88161968d565b9450508284511690506175d4565b6176108785619367565b9450505050506152d7565b83832061762885886184f4565b6176329087619367565b91505b85821061766d5784822080820361765a576176508684619367565b93505050506152d7565b6176656001846184f4565b925050617635565b505b5092949350505050565b60008381868511617784576020851161773357600085156176c55760016176a18760206184f4565b6176ac906008619583565b6176b7906002619681565b6176c191906184f4565b1990505b845181166000876176d68b8b619367565b6176e091906184f4565b855190915083165b8281146177255781861061770d576177008b8b619367565b96505050505050506152d7565b8561771781619569565b9650508386511690506176e8565b8596505050505050506152d7565b508383206000905b61774586896184f4565b82116177825785832080820361776157839450505050506152d7565b61776c600185619367565b935050818061777a90619569565b92505061773b565b505b61778e8787619367565b979650505050505050565b604080518082019091526000808252602082015260006177cb8560000151866020015186600001518760200151617679565b6020808701805191860191909152519091506177e790826184f4565b8352845160208601516177fa9190619367565b8103617809576000855261783b565b835183516178179190619367565b855186906178269083906184f4565b90525083516178359082619367565b60208601525b50909392505050565b6020811061787c578151835261785b602084619367565b9250617868602083619367565b91506178756020826184f4565b9050617844565b60001981156178ab5760016178928360206184f4565b61789e90610100619681565b6178a891906184f4565b90505b9151835183169219169190911790915250565b606060006178cc8484613daa565b80516020808301516040519394506178e6939091016196a4565b60405160208183030381529060405291505092915050565b8151815160009190811115617911575081515b6020808501519084015160005b838110156179ca578251825180821461799a5760001960208710156179795760018461794b8960206184f4565b6179559190619367565b617960906008619583565b61796b906002619681565b61797591906184f4565b1990505b81811683821681810391146179975797506138e19650505050505050565b50505b6179a5602086619367565b94506179b2602085619367565b935050506020816179c39190619367565b905061791e565b508451865161445891906196fc565b610c9f806200971d83390190565b611dff806200a3bc83390190565b611783806200c1bb83390190565b610de0806200d93e83390190565b610a2c806200e71e83390190565b610b3f806200f14a83390190565b612072806200fc8983390190565b6040518060e00160405280606081526020016060815260200160608152602001600015158152602001600015158152602001600015158152602001617a7e617a83565b905290565b60405180610100016040528060001515815260200160001515815260200160608152602001600080191681526020016060815260200160608152602001600015158152602001617a7e6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b81811015617b355783516001600160a01b0316835260209384019390920191600101617b0e565b509095945050505050565b60005b83811015617b5b578181015183820152602001617b43565b50506000910152565b60008151808452617b7c816020860160208601617b40565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015617c8c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b81811015617c72577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8503018352617c5c848651617b64565b6020958601959094509290920191600101617c22565b509197505050602094850194929092019150600101617bb8565b50929695505050505050565b600081518084526020840193506020830160005b82811015617cec5781517fffffffff0000000000000000000000000000000000000000000000000000000016865260209586019590910190600101617cac565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015617c8c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160408752617d626040880182617b64565b9050602082015191508681036020880152617d7d8183617c98565b965050506020938401939190910190600101617d1e565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015617c8c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452617df6858351617b64565b94506020938401939190910190600101617dbc565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015617c8c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b0381511686526020810151905060406020870152617e8c6040870182617c98565b9550506020938401939190910190600101617e33565b600060208284031215617eb457600080fd5b815180151581146139da57600080fd5b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff8816604083015286606083015260038610617f7e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8560808301528460a0830152617f9f60c08301856001600160a01b03169052565b6001600160a01b03831660e0830152979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c90821680617ffc57607f821691505b602082108103615ba3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f82111561807f57806000526020600020601f840160051c8101602085101561805c5750805b601f840160051c820191505b8181101561807c5760008155600101618068565b50505b505050565b815167ffffffffffffffff81111561809e5761809e617fb9565b6180b2816180ac8454617fe8565b84618035565b6020601f8211600181146180e657600083156180ce5750848201515b600019600385901b1c1916600184901b17845561807c565b600084815260208120601f198516915b8281101561811657878501518255602094850194600190920191016180f6565b50848210156181345786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60006020828403121561815557600080fd5b5051919050565b60608152600061816f6060830186617b64565b602083019490945250901515604090910152919050565b600081546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160a06060850152600081546181d081617fe8565b8060a088015260018216600081146181ef57600181146182295761825d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660c089015260c082151560051b890101935061825d565b84600052602060002060005b838110156182545781548a820160c00152600190910190602001618235565b890160c0019450505b50505060038401546080860152809250505092915050565b60c08152600061828860c0830189617b64565b62ffffff881660208401526001600160a01b038716604084015282810360608401526182b48187617b64565b905060ff8516608084015282810360a08401526182d18185618186565b9998505050505050505050565b6001600160a01b0384168152826020820152606060408201526000615b896060830184617b64565b60c08152600061831960c0830189617b64565b8760208401526001600160a01b038716604084015282810360608401526183408187617b64565b6080840195909552505090151560a090910152949350505050565b6001600160a01b038616815284602082015260a06040820152600061838360a0830186617b64565b6060830194909452509015156080909101529392505050565b8281526040602082015260006152d76040830184617b64565b6001600160a01b03831681526040602082015260006152d76040830184617b64565b6040516060810167ffffffffffffffff811182821017156183fa576183fa617fb9565b60405290565b60008067ffffffffffffffff84111561841b5761841b617fb9565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561844a5761844a617fb9565b60405283815290508082840185101561846257600080fd5b6156a6846020830185617b40565b600082601f83011261848157600080fd5b6139da83835160208501618400565b6000602082840312156184a257600080fd5b815167ffffffffffffffff8111156184b957600080fd5b6138dd84828501618470565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156138e1576138e16184c5565b6101008152600061851c61010083018b617b64565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a084015261854f8187617b64565b90508460c084015282810360e0840152613bcd8185618186565b60c08152600061857c60c0830189617b64565b8760208401526001600160a01b038716604084015282810360608401526185a38187617b64565b905084608084015282810360a08401526182d18185618186565b6080815260006185d06080830187617b64565b82810360208401526185e28187617b64565b9050846040840152828103606084015261778e8185618186565b60a08152600061860f60a0830188617b64565b6001600160a01b038716602084015282810360408401526186308187617b64565b9050846060840152828103608084015261864a8185618186565b98975050505050505050565b60a08152600061866960a0830188617b64565b6001600160a01b0387166020840152828103604084015261868a8187617b64565b905060ff85166060840152828103608084015261864a8185618186565b60a0815260006186ba60a0830188617b64565b6001600160a01b038716602084015282810360408401526186db8187617b64565b606084019590955250509015156080909101529392505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161872d81601a850160208801617b40565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161876a81601c840160208801617b40565b01601c01949350505050565b6020815260006139da6020830184617b64565b60006020828403121561879b57600080fd5b81516001600160a01b03811681146139da57600080fd5b600083516187c4818460208801617b40565b8351908301906187d8818360208801617b40565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161881981601a850160208801617b40565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a918401918201528351618856816033840160208801617b40565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f555400000000000000000000000000000000000000000060608201526080602082015260006139da6080830184617b64565b6000602082840312156188e557600080fd5b815167ffffffffffffffff8111156188fc57600080fd5b8201601f8101841361890d57600080fd5b6138dd84825160208401618400565b6000855161892e818460208a01617b40565b7f2f000000000000000000000000000000000000000000000000000000000000009083019081528551618968816001840160208a01617b40565b7f2f000000000000000000000000000000000000000000000000000000000000006001929091019182015284516189a6816002840160208901617b40565b6001818301019150507f2f00000000000000000000000000000000000000000000000000000000000000600182015283516189e8816002840160208801617b40565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b604081526000618a336040830184617b64565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e2061727469666163742000815260008251618aaa81601f850160208701617b40565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b604081526000618b176040830184617b64565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081526000618b696040830184617b64565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b27000000000000000000000000815260008251618be0816014850160208701617b40565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b604081526000618c276040830185617b64565b82810360208401526139d68185617b64565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f2200000000000000000000000000000000000000000000000000000000000000815260008251618ca0816001850160208701617b40565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b60008251618ce6818460208701617b40565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e747261637420000000000000000000000000000000000000000000604082015260008251618d9981604b850160208701617b40565b91909101604b0192915050565b600060ff821660ff8103618dbc57618dbc6184c5565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c69400000000000000000000000000000000000000000000000602082015260008251618e23816029850160208701617b40565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f504154480000000000000000000060608201526080602082015260006139da6080830184617b64565b600060208284031215618e8957600080fd5b815167ffffffffffffffff811115618ea057600080fd5b820160608185031215618eb257600080fd5b618eba6183d7565b81518060030b8114618ecb57600080fd5b8152602082015167ffffffffffffffff811115618ee757600080fd5b618ef386828501618470565b602083015250604082015167ffffffffffffffff811115618f1357600080fd5b618f1f86828501618470565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f2200000000000000000000000000000000000000000000000000000000000000602082015260008251618f8b816021850160208701617b40565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f2700000000000000000000000000000000000000000000000000000000000000602082015260008351619177816021850160208801617b40565b7f2720696e206f75747075743a200000000000000000000000000000000000000060219184019182015283516191b481602e840160208801617b40565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a200000000000000000000000000000000000000000000000602082015260008251618e23816029850160208701617b40565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161927c816022850160208701617b40565b9190910160220192915050565b7f436f6e7472616374206e616d65200000000000000000000000000000000000008152600082516192c181600e850160208701617b40565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b808201808211156138e1576138e16184c5565b7f53504458206c6963656e7365206964656e7469666965722000000000000000008152600083516193b2816018850160208801617b40565b7f20696e200000000000000000000000000000000000000000000000000000000060189184019182015283516193ef81601c840160208801617b40565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b600082516194f5818460208701617b40565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161955c81601c850160208701617b40565b91909101601c0192915050565b6000600019820361957c5761957c6184c5565b5060010190565b80820281158282048414176138e1576138e16184c5565b6001815b60018411156195d5578085048111156195b9576195b96184c5565b60018416156195c757908102905b60019390931c92800261959e565b935093915050565b6000826195ec575060016138e1565b816195f9575060006138e1565b816001811461960f576002811461961957619635565b60019150506138e1565b60ff84111561962a5761962a6184c5565b50506001821b6138e1565b5060208310610133831016604e8410600b8410161715619658575081810a6138e1565b619665600019848461959a565b8060001904821115619679576196796184c5565b029392505050565b60006139da83836195dd565b60008161969c5761969c6184c5565b506000190190565b600083516196b6818460208801617b40565b7f3a0000000000000000000000000000000000000000000000000000000000000090830190815283516196f0816001840160208801617b40565b01600101949350505050565b8181036000831280158383131683831282161715616c9e57616c9e6184c556fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220a043c41353215fce25ecb67a8a4f6f724aaa86dea8dcb0a6975bfb1f10ff3beb64736f6c634300081a003360a060405234801561001057600080fd5b50604051611dff380380611dff83398101604081905261002f916101fd565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03838116608052600480546001600160a01b0319169184169190911790556100bb60008261014c565b506100e67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261014c565b506101117f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361014c565b5061012a600080516020611ddf8339815191528261014c565b50610143600080516020611ddf8339815191528361014c565b50505050610240565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166101d75760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016101db565b5060005b92915050565b80516001600160a01b03811681146101f857600080fd5b919050565b60008060006060848603121561021257600080fd5b61021b846101e1565b9250610229602085016101e1565b9150610237604085016101e1565b90509250925092565b608051611b68610277600039600081816101d501528181610574015281816105c901528181610b010152610b560152611b686000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c806385f438c1116100ee578063c709ab6e11610097578063d9caed1211610071578063d9caed12146103e0578063e609055e146103f3578063e63ab1e914610406578063eab103df1461042d57600080fd5b8063c709ab6e14610397578063d547741f146103aa578063d936547e146103bd57600080fd5b80639a590427116100c85780639a590427146103695780639b19251a1461037c578063a217fddf1461038f57600080fd5b806385f438c1146102f657806391d148541461031d578063950837aa1461035657600080fd5b806336568abe116101505780635b1125911161012a5780635b112591146102d05780635c975abb146102e35780638456cb59146102ee57600080fd5b806336568abe1461028e5780633f4ba83a146102a1578063570618e1146102a957600080fd5b8063248a9ca311610181578063248a9ca314610224578063252f07bf146102565780632f2ff15d1461027b57600080fd5b806301ffc9a7146101a8578063116191b6146101d057806321fc65f21461020f575b600080fd5b6101bb6101b636600461155e565b610440565b60405190151581526020015b60405180910390f35b6101f77f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101c7565b61022261021d3660046115fe565b6104d9565b005b610248610232366004611671565b6000908152600160208190526040909120015490565b6040519081526020016101c7565b6004546101bb9074010000000000000000000000000000000000000000900460ff1681565b61022261028936600461168a565b610699565b61022261029c36600461168a565b6106c5565b610222610716565b6102487f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b6004546101f7906001600160a01b031681565b60025460ff166101bb565b61022261074b565b6102487f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101bb61032b36600461168a565b60009182526001602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6102226103643660046116ba565b61077d565b6102226103773660046116ba565b6108fb565b61022261038a3660046116ba565b6109af565b610248600081565b6102226103a53660046116d7565b610a66565b6102226103b836600461168a565b610c2b565b6101bb6103cb3660046116ba565b60036020526000908152604090205460ff1681565b6102226103ee36600461177a565b610c51565b6102226104013660046117bb565b610d49565b6102487f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b61022261043b36600461185a565b610f75565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104d357507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104e1610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461050b8161100e565b610513611018565b6001600160a01b03851660009081526003602052604090205460ff16610565576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105996001600160a01b0386167f000000000000000000000000000000000000000000000000000000000000000086611057565b6040517f5131ab590000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906106069088908a908990899089906004016118c0565b600060405180830381600087803b15801561062057600080fd5b505af1158015610634573d6000803e3d6000fd5b50505050846001600160a01b0316866001600160a01b03167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d586868660405161067f93929190611903565b60405180910390a3506106926001600055565b5050505050565b600082815260016020819052604090912001546106b58161100e565b6106bf83836110cb565b50505050565b6001600160a01b0381163314610707576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610711828261115e565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107408161100e565b6107486111e5565b50565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107758161100e565b610748611237565b60006107888161100e565b6001600160a01b0382166107c8576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004546107ff907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4906001600160a01b031661115e565b50600454610837907f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a906001600160a01b031661115e565b506108627f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4836110cb565b5061088d7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a836110cb565b50600480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040519081527f086480ac96b6cbd744062a9994d7b954673bf500d6f362180ecd9cb5828e07ba9060200160405180910390a15050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6109258161100e565b6001600160a01b038216610965576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6109d98161100e565b6001600160a01b038216610a19576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b610a6e610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610a988161100e565b610aa0611018565b6001600160a01b03861660009081526003602052604090205460ff16610af2576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b266001600160a01b0387167f000000000000000000000000000000000000000000000000000000000000000087611057565b6040517fd0b492c30000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610b959089908b908a908a908a908a906004016119d7565b600060405180830381600087803b158015610baf57600080fd5b505af1158015610bc3573d6000803e3d6000fd5b50505050856001600160a01b0316876001600160a01b03167f2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b87878787604051610c109493929190611a2e565b60405180910390a350610c236001600055565b505050505050565b60008281526001602081905260409091200154610c478161100e565b6106bf838361115e565b610c59610fcb565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610c838161100e565b610c8b611018565b6001600160a01b03831660009081526003602052604090205460ff16610cdd576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cf16001600160a01b0384168584611057565b826001600160a01b0316846001600160a01b03167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610d3691815260200190565b60405180910390a3506107116001600055565b610d51610fcb565b610d59611018565b60045474010000000000000000000000000000000000000000900460ff16610dad576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03841660009081526003602052604090205460ff16610dff576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610e5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e839190611a5a565b9050610e9a6001600160a01b038616333087611274565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa158015610f21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f459190611a5a565b610f4f9190611a73565b8787604051610f62959493929190611aad565b60405180910390a250610c236001600055565b6000610f808161100e565b506004805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b600260005403611007576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61074881336112ad565b60025460ff1615611055576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6040516001600160a01b0383811660248301526044820183905261071191859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611324565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166111565760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016104d3565b5060006104d3565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff16156111565760008381526001602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104d3565b6111ed6113a0565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b61123f611018565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861121a3390565b6040516001600160a01b0384811660248301528381166044830152606482018390526106bf9186918216906323b872dd90608401611084565b60008281526001602090815260408083206001600160a01b038516845290915290205460ff16611320576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044015b60405180910390fd5b5050565b60006113396001600160a01b038416836113dc565b9050805160001415801561135e57508080602001905181019061135c9190611ae6565b155b15610711576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611317565b60025460ff16611055576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606113ea838360006113f1565b9392505050565b60608147101561142f576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611317565b600080856001600160a01b0316848660405161144b9190611b03565b60006040518083038185875af1925050503d8060008114611488576040519150601f19603f3d011682016040523d82523d6000602084013e61148d565b606091505b509150915061149d8683836114a7565b9695505050505050565b6060826114bc576114b78261151c565b6113ea565b81511580156114d357506001600160a01b0384163b155b15611515576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611317565b50806113ea565b80511561152c5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561157057600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146113ea57600080fd5b6001600160a01b038116811461074857600080fd5b60008083601f8401126115c757600080fd5b50813567ffffffffffffffff8111156115df57600080fd5b6020830191508360208285010111156115f757600080fd5b9250929050565b60008060008060006080868803121561161657600080fd5b8535611621816115a0565b94506020860135611631816115a0565b935060408601359250606086013567ffffffffffffffff81111561165457600080fd5b611660888289016115b5565b969995985093965092949392505050565b60006020828403121561168357600080fd5b5035919050565b6000806040838503121561169d57600080fd5b8235915060208301356116af816115a0565b809150509250929050565b6000602082840312156116cc57600080fd5b81356113ea816115a0565b60008060008060008060a087890312156116f057600080fd5b86356116fb816115a0565b9550602087013561170b816115a0565b945060408701359350606087013567ffffffffffffffff81111561172e57600080fd5b61173a89828a016115b5565b909450925050608087013567ffffffffffffffff81111561175a57600080fd5b87016060818a03121561176c57600080fd5b809150509295509295509295565b60008060006060848603121561178f57600080fd5b833561179a816115a0565b925060208401356117aa816115a0565b929592945050506040919091013590565b600080600080600080608087890312156117d457600080fd5b863567ffffffffffffffff8111156117eb57600080fd5b6117f789828a016115b5565b909750955050602087013561180b816115a0565b935060408701359250606087013567ffffffffffffffff81111561182e57600080fd5b61183a89828a016115b5565b979a9699509497509295939492505050565b801515811461074857600080fd5b60006020828403121561186c57600080fd5b81356113ea8161184c565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b03861681526001600160a01b03851660208201528360408201526080606082015260006118f8608083018486611877565b979650505050505050565b83815260406020820152600061191d604083018486611877565b95945050505050565b60008135611933816115a0565b6001600160a01b03168352602082013567ffffffffffffffff811680821461195a57600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261199657600080fd5b820160208101903567ffffffffffffffff8111156119b357600080fd5b8036038213156119c257600080fd5b6060604086015261191d606086018284611877565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a060608201526000611a0f60a083018587611877565b8281036080840152611a218185611926565b9998505050505050505050565b848152606060208201526000611a48606083018587611877565b82810360408401526118f88185611926565b600060208284031215611a6c57600080fd5b5051919050565b818103818111156104d3577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b606081526000611ac1606083018789611877565b8560208401528281036040840152611ada818587611877565b98975050505050505050565b600060208284031215611af857600080fd5b81516113ea8161184c565b6000825160005b81811015611b245760208186018101518583015201611b0a565b50600092019182525091905056fea264697066735822122025ce6ff52331c185a420578d5cf53c21053d0854e1a0b16899e18daa5a47faa464736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60c060405260001960035534801561001657600080fd5b5060405161178338038061178383398101604081905261003591610220565b60016000819055805460ff19169055838383836001600160a01b038416158061006557506001600160a01b038316155b8061007757506001600160a01b038216155b8061008957506001600160a01b038116155b156100a75760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03808516608052831660a0526100c5600082610154565b506100f07f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610154565b5061011b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610154565b506101467f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610154565b505050505050505050610274565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166101fa5760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101b23390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101fe565b5060005b92915050565b80516001600160a01b038116811461021b57600080fd5b919050565b6000806000806080858703121561023657600080fd5b61023f85610204565b935061024d60208601610204565b925061025b60408601610204565b915061026960608601610204565b905092959194509250565b60805160a0516114ab6102d86000396000818161021d01528181610531015281816108290152818161099a01528181610aef0152610c110152600081816101d1015281816104a1015281816105040152818161079901526107fc01526114ab6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80635e3e9fef116100d857806391d148541161008c578063d547741f11610066578063d547741f14610386578063d5abeb0114610399578063e63ab1e9146103a257600080fd5b806391d1485414610311578063a217fddf14610357578063a783c7891461035f57600080fd5b8063743e0c9b116100bd578063743e0c9b146102cf5780638456cb59146102e257806385f438c1146102ea57600080fd5b80635e3e9fef146102a95780636f8b44b0146102bc57600080fd5b8063248a9ca31161012f57806336568abe1161011457806336568abe146102835780633f4ba83a146102965780635c975abb1461029e57600080fd5b8063248a9ca31461023f5780632f2ff15d1461027057600080fd5b8063106e629011610160578063106e6290146101b9578063116191b6146101cc57806321e093b11461021857600080fd5b806301ffc9a71461017c578063057e0f25146101a4575b600080fd5b61018f61018a366004610fd3565b6103c9565b60405190151581526020015b60405180910390f35b6101b76101b236600461108e565b610462565b005b6101b76101c7366004611126565b6105fc565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b61026261024d366004611159565b60009081526002602052604090206001015490565b60405190815260200161019b565b6101b761027e366004611172565b6106a1565b6101b7610291366004611172565b6106cc565b6101b7610725565b60015460ff1661018f565b6101b76102b736600461119e565b61075a565b6101b76102ca366004611159565b6108ef565b6101b76102dd366004611159565b61095d565b6101b7610a07565b6102627f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b61018f61031f366004611172565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610262600081565b6102627f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101b7610394366004611172565b610a39565b61026260035481565b6102627f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061045c57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b61046a610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461049481610aa1565b61049c610aab565b6104c77f00000000000000000000000000000000000000000000000000000000000000008785610aea565b6040517fd0b492c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610563907f0000000000000000000000000000000000000000000000000000000000000000908b908b908b908b908a9060040161130e565b600060405180830381600087803b15801561057d57600080fd5b505af1158015610591573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff167f52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c878787866040516105e1949392919061137f565b60405180910390a2506105f46001600055565b505050505050565b610604610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461062e81610aa1565b610636610aab565b610641848484610aea565b8373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d58460405161068991815260200190565b60405180910390a25061069c6001600055565b505050565b6000828152600260205260409020600101546106bc81610aa1565b6106c68383610c72565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331461071b576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61069c8282610d72565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61074f81610aa1565b610757610e31565b50565b610762610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461078c81610aa1565b610794610aab565b6107bf7f00000000000000000000000000000000000000000000000000000000000000008684610aea565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab5990610859907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a906004016113b6565b600060405180830381600087803b15801561087357600080fd5b505af1158015610887573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d8686866040516108d593929190611408565b60405180910390a2506108e86001600055565b5050505050565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61091981610aa1565b610921610aab565b60038290556040518281527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c9060200160405180910390a15050565b610965610aab565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b1580156109f357600080fd5b505af11580156108e8573d6000803e3d6000fd5b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a3181610aa1565b610757610eae565b600082815260026020526040902060010154610a5481610aa1565b6106c68383610d72565b600260005403610a9a576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6107578133610f07565b60015460ff1615610ae8576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6003547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7c9190611422565b610b86908461143b565b1115610bbe576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201849052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b158015610c5557600080fd5b505af1158015610c69573d6000803e3d6000fd5b50505050505050565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610d6a57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610d083390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161045c565b50600061045c565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610d6a57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161045c565b610e39610f97565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b610eb6610aab565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610e84565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610f93576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440160405180910390fd5b5050565b60015460ff16610ae8576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215610fe557600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461101557600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461104057600080fd5b919050565b60008083601f84011261105757600080fd5b50813567ffffffffffffffff81111561106f57600080fd5b60208301915083602082850101111561108757600080fd5b9250929050565b60008060008060008060a087890312156110a757600080fd5b6110b08761101c565b955060208701359450604087013567ffffffffffffffff8111156110d357600080fd5b6110df89828a01611045565b90955093505060608701359150608087013567ffffffffffffffff81111561110657600080fd5b87016060818a03121561111857600080fd5b809150509295509295509295565b60008060006060848603121561113b57600080fd5b6111448461101c565b95602085013595506040909401359392505050565b60006020828403121561116b57600080fd5b5035919050565b6000806040838503121561118557600080fd5b823591506111956020840161101c565b90509250929050565b6000806000806000608086880312156111b657600080fd5b6111bf8661101c565b945060208601359350604086013567ffffffffffffffff8111156111e257600080fd5b6111ee88828901611045565b96999598509660600135949350505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff6112678261101c565b1682526000602082013567ffffffffffffffff811680821461128857600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126112c457600080fd5b820160208101903567ffffffffffffffff8111156112e157600080fd5b8036038213156112f057600080fd5b60606040860152611305606086018284611200565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a06060820152600061136060a083018587611200565b82810360808401526113728185611249565b9998505050505050505050565b848152606060208201526000611399606083018587611200565b82810360408401526113ab8185611249565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006113ab608083018486611200565b838152604060208201526000611305604083018486611200565b60006020828403121561143457600080fd5b5051919050565b8082018082111561045c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea26469706673582212201a525a06b9312846603a9a514a3c7da5c267e70352788b0c22c2ab173492c28964736f6c634300081a00336080604052348015600f57600080fd5b506001600055610dbc806100246000396000f3fe6080604052600436106100635760003560e01c8063c513169111610040578063c5131691146100c1578063e04d4f97146100e1578063f05b6abf146100f457005b8063357fc5a21461006c578063660b9de01461008c5780636ed70169146100ac57005b3661006a57005b005b34801561007857600080fd5b5061006a6100873660046106bd565b610114565b34801561009857600080fd5b5061006a6100a73660046106f9565b6101aa565b3480156100b857600080fd5b5061006a6101e6565b3480156100cd57600080fd5b5061006a6100dc3660046106bd565b61021b565b61006a6100ef366004610859565b6102f6565b34801561010057600080fd5b5061006a61010f366004610945565b61033a565b61011c61036f565b61013e73ffffffffffffffffffffffffffffffffffffffff83163383866103b2565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101a56001600055565b505050565b7f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e33826040516101db929190610a78565b60405180910390a150565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61022361036f565b6000610230600285610b57565b90508060000361026c576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61028e73ffffffffffffffffffffffffffffffffffffffff84163384846103b2565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101a56001600055565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa333485858560405161032d959493929190610c00565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b1463384848460405161032d9493929190610c8a565b6002600054036103ab576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261044790859061044d565b50505050565b600061046f73ffffffffffffffffffffffffffffffffffffffff8416836104e8565b905080516000141580156104945750808060200190518101906104929190610d4d565b155b156101a5576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b60606104f6838360006104fd565b9392505050565b60608147101561053b576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016104df565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105649190610d6a565b60006040518083038185875af1925050503d80600081146105a1576040519150601f19603f3d011682016040523d82523d6000602084013e6105a6565b606091505b50915091506105b68683836105c0565b9695505050505050565b6060826105d5576105d08261064f565b6104f6565b81511580156105f9575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610648576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016104df565b50806104f6565b80511561065f5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff811681146106b857600080fd5b919050565b6000806000606084860312156106d257600080fd5b833592506106e260208501610694565b91506106f060408501610694565b90509250925092565b60006020828403121561070b57600080fd5b813567ffffffffffffffff81111561072257600080fd5b8201606081850312156104f657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156107aa576107aa610734565b604052919050565b600082601f8301126107c357600080fd5b813567ffffffffffffffff8111156107dd576107dd610734565b61080e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610763565b81815284602083860101111561082357600080fd5b816020850160208301376000918101602001919091529392505050565b801515811461069157600080fd5b80356106b881610840565b60008060006060848603121561086e57600080fd5b833567ffffffffffffffff81111561088557600080fd5b610891868287016107b2565b9350506020840135915060408401356108a981610840565b809150509250925092565b600067ffffffffffffffff8211156108ce576108ce610734565b5060051b60200190565b600082601f8301126108e957600080fd5b81356108fc6108f7826108b4565b610763565b8082825260208201915060208360051b86010192508583111561091e57600080fd5b602085015b8381101561093b578035835260209283019201610923565b5095945050505050565b60008060006060848603121561095a57600080fd5b833567ffffffffffffffff81111561097157600080fd5b8401601f8101861361098257600080fd5b80356109906108f7826108b4565b8082825260208201915060208360051b8501019250888311156109b257600080fd5b602084015b838110156109f457803567ffffffffffffffff8111156109d657600080fd5b6109e58b6020838901016107b2565b845250602092830192016109b7565b509550505050602084013567ffffffffffffffff811115610a1457600080fd5b610a20868287016108d8565b9250506106f06040850161084e565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610ab683610694565b1660408201526000602083013567ffffffffffffffff8116808214610ada57600080fd5b6060840152506040830135368490037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112610b1657600080fd5b830160208101903567ffffffffffffffff811115610b3357600080fd5b803603821315610b4257600080fd5b606060808501526105b660a085018284610a2f565b600082610b8d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60005b83811015610bad578181015183820152602001610b95565b50506000910152565b60008151808452610bce816020860160208601610b92565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610c3560a0830186610bb6565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610c80578151865260209586019590910190600101610c62565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610d1d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610d08858351610bb6565b94506020938401939190910190600101610cce565b505050508281036040840152610d338186610c4e565b915050610d44606083018415159052565b95945050505050565b600060208284031215610d5f57600080fd5b81516104f681610840565b60008251610d7c818460208701610b92565b919091019291505056fea26469706673582212202c10ab94780c92f0dd656897a07025a325b3c0d38d265113cb76037b152226de64736f6c634300081a00336080604052348015600f57600080fd5b50604051610a2c380380610a2c833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b61099f8061008d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b61461005b5780637a34d8bb146100a4575b600080fd5b6100596100543660046105f5565b6100b7565b005b60005461007b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100596100b2366004610695565b610313565b60008383836040516024016100ce93929190610791565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905260005490915073ffffffffffffffffffffffffffffffffffffffff8087169163095ea7b391166101758960026107bb565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af11580156101e5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020991906107fb565b61023f576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160a0810182526103218082526001602080840182905283850192909252835191820184526000808352606084019290925260808301829052905492517f048ae42c000000000000000000000000000000000000000000000000000000008152919273ffffffffffffffffffffffffffffffffffffffff169163048ae42c916102d7918c918c918c9189918990600401610894565b600060405180830381600087803b1580156102f157600080fd5b505af1158015610305573d6000803e3d6000fd5b505050505050505050505050565b600083838360405160240161032a93929190610791565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f9700000000000000000000000000000000000000000000000000000000179052815160a081018352610321808252600182840181905282850191909152835192830184526000808452606083019390935260808201839052915492517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff93841660048201526024810183905293945092909188169063095ea7b3906044016020604051808303816000875af1158015610455573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047991906107fb565b506000546040517f1cb5ea7500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690631cb5ea75906102d7908b908b90889087908990600401610902565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261051857600080fd5b81356020830160008067ffffffffffffffff841115610539576105396104d8565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715610586576105866104d8565b60405283815290508082840187101561059e57600080fd5b838360208301376000602085830101528094505050505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146105df57600080fd5b919050565b80151581146105f257600080fd5b50565b60008060008060008060c0878903121561060e57600080fd5b863567ffffffffffffffff81111561062557600080fd5b61063189828a01610507565b96505060208701359450610647604088016105bb565b9350606087013567ffffffffffffffff81111561066357600080fd5b61066f89828a01610507565b9350506080870135915060a0870135610687816105e4565b809150509295509295509295565b600080600080600060a086880312156106ad57600080fd5b853567ffffffffffffffff8111156106c457600080fd5b6106d088828901610507565b9550506106df602087016105bb565b9350604086013567ffffffffffffffff8111156106fb57600080fd5b61070788828901610507565b93505060608601359150608086013561071f816105e4565b809150509295509295909350565b6000815180845260005b8181101561075357602081850181015186830182015201610737565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6060815260006107a4606083018661072d565b602083019490945250901515604090910152919050565b808201808211156107f5577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b60006020828403121561080d57600080fd5b8151610818816105e4565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff60408201511660408301526000606082015160a0606085015261088060a085018261072d565b608093840151949093019390935250919050565b60c0815260006108a760c083018961072d565b87602084015273ffffffffffffffffffffffffffffffffffffffff8716604084015282810360608401526108db818761072d565b905084608084015282810360a08401526108f5818561081f565b9998505050505050505050565b60a08152600061091560a083018861072d565b73ffffffffffffffffffffffffffffffffffffffff871660208401528281036040840152610943818761072d565b9050846060840152828103608084015261095d818561081f565b9897505050505050505056fea2646970667358221220fc80e9f3a051dcf274f22b51503d752bb5b6227e29a3a916c124600f26a0fcd764736f6c634300081a0033608060405234801561001057600080fd5b50604051610b3f380380610b3f83398101604081905261002f916100b9565b600380546001600160a01b038086166001600160a01b0319928316179092556004805485841690831617905560058054928416929091169190911790556040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac590600090a15050506100fc565b80516001600160a01b03811681146100b457600080fd5b919050565b6000806000606084860312156100ce57600080fd5b6100d78461009d565b92506100e56020850161009d565b91506100f36040850161009d565b90509250925092565b610a348061010b6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806397770dff11610081578063d7fd7afb1161005b578063d7fd7afb146101f2578063d936a01214610220578063ee2815ba1461024057600080fd5b806397770dff146101b9578063a7cb0507146101cc578063c63585cc146101df57600080fd5b8063513a9c05116100b2578063513a9c0514610143578063569541b914610179578063842da36d1461019957600080fd5b80630be15547146100ce5780633c669d551461012e575b600080fd5b6101046100dc36600461071e565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61014161013c366004610760565b610253565b005b61010461015136600461071e565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101049073ffffffffffffffffffffffffffffffffffffffff1681565b6005546101049073ffffffffffffffffffffffffffffffffffffffff1681565b6101416101c73660046107fd565b6103a0565b6101416101da36600461081f565b610419565b6101046101ed366004610841565b610467565b61021261020036600461071e565b60006020819052908152604090205481565b604051908152602001610125565b6004546101049073ffffffffffffffffffffffffffffffffffffffff1681565b61014161024e366004610884565b61059c565b604080516080810182526000606082019081528152336020820152468183015290517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820186905286169063a9059cbb906044016020604051808303816000875af11580156102e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030b91906108b0565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87169063de43156e90610366908490899089908990899060040161091b565b600060405180830381600087803b15801561038057600080fd5b505af1158015610394573d6000803e3d6000fd5b50505050505050505050565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e9060200160405180910390a150565b6000828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b60008060006104768585610620565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b166034820152919350915086906048016040516020818303038152906040528051906020012060405160200161055c9291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b60008281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d910161045b565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610688576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16106106c25782846106c5565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610717576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b60006020828403121561073057600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461075b57600080fd5b919050565b60008060008060006080868803121561077857600080fd5b61078186610737565b945061078f60208701610737565b935060408601359250606086013567ffffffffffffffff8111156107b257600080fd5b8601601f810188136107c357600080fd5b803567ffffffffffffffff8111156107da57600080fd5b8860208284010111156107ec57600080fd5b959894975092955050506020019190565b60006020828403121561080f57600080fd5b61081882610737565b9392505050565b6000806040838503121561083257600080fd5b50508035926020909101359150565b60008060006060848603121561085657600080fd5b61085f84610737565b925061086d60208501610737565b915061087b60408501610737565b90509250925092565b6000806040838503121561089757600080fd5b823591506108a760208401610737565b90509250929050565b6000602082840312156108c257600080fd5b8151801515811461081857600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60808152600086516060608084015280518060e085015260005b81811015610953576020818401810151610100878401015201610935565b5060008482016101000152602089015173ffffffffffffffffffffffffffffffffffffffff811660a0860152601f9091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168401915050604088015160c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401526101008382030160608401526109f2610100820185876108d2565b9897505050505050505056fea2646970667358221220f5c4cae5cc505bc5e795bca446028bb75a72680a148baa4a180dc1a1714ebe3764736f6c634300081a003360c060405234801561001057600080fd5b5060405161207238038061207283398101604081905261002f916101f0565b6001600160a01b038216158061004c57506001600160a01b038116155b1561006a5760405163d92e233d60e01b815260040160405180910390fd5b60066100768982610342565b5060076100838882610342565b506008805460ff191660ff881617905560808590528360028111156100aa576100aa610400565b60a08160028111156100be576100be610400565b905250600192909255600080546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506104169350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261013357600080fd5b81516001600160401b0381111561014c5761014c61010c565b604051601f8201601f19908116603f011681016001600160401b038111828210171561017a5761017a61010c565b60405281815283820160200185101561019257600080fd5b60005b828110156101b157602081860181015183830182015201610195565b506000918101602001919091529392505050565b8051600381106101d457600080fd5b919050565b80516001600160a01b03811681146101d457600080fd5b600080600080600080600080610100898b03121561020d57600080fd5b88516001600160401b0381111561022357600080fd5b61022f8b828c01610122565b60208b015190995090506001600160401b0381111561024d57600080fd5b6102598b828c01610122565b975050604089015160ff8116811461027057600080fd5b60608a0151909650945061028660808a016101c5565b60a08a0151909450925061029c60c08a016101d9565b91506102aa60e08a016101d9565b90509295985092959890939650565b600181811c908216806102cd57607f821691505b6020821081036102ed57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561033d57806000526020600020601f840160051c8101602085101561031a5750805b601f840160051c820191505b8181101561033a5760008155600101610326565b50505b505050565b81516001600160401b0381111561035b5761035b61010c565b61036f8161036984546102b9565b846102f3565b6020601f8211600181146103a3576000831561038b5750848201515b600019600385901b1c1916600184901b17845561033a565b600084815260208120601f198516915b828110156103d357878501518255602094850194600190920191016103b3565b50848210156103f15786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a051611c1b61045760003960006103440152600081816102f001528181610bdc01528181610ce201528181610efe01526110040152611c1b6000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806395d89b41116100f9578063ccc7759911610097578063eddeb12311610071578063eddeb12314610461578063f2441b3214610474578063f687d12a14610494578063fc5fecd5146104a757600080fd5b8063ccc77599146103d4578063d9eeebed146103e7578063dd62ed3e1461041b57600080fd5b8063b84c8246116100d3578063b84c824614610386578063c47f00271461039b578063c7012626146103ae578063c835d7cc146103c157600080fd5b806395d89b4114610337578063a3413d031461033f578063a9059cbb1461037357600080fd5b80633ce4a5bc116101665780634d8943bb116101405780634d8943bb146102ac57806370a08231146102b557806385e1f4d0146102eb5780638b851b951461031257600080fd5b80633ce4a5bc1461024657806342966c681461028657806347e7ef241461029957600080fd5b806318160ddd1161019757806318160ddd1461021657806323b872dd1461021e578063313ce5671461023157600080fd5b806306fdde03146101be578063091d2788146101dc578063095ea7b3146101f3575b600080fd5b6101c66104ba565b6040516101d39190611648565b60405180910390f35b6101e560015481565b6040519081526020016101d3565b610206610201366004611687565b61054c565b60405190151581526020016101d3565b6005546101e5565b61020661022c3660046116b3565b610563565b60085460405160ff90911681526020016101d3565b61026173735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6102066102943660046116f4565b6105fa565b6102066102a7366004611687565b61060e565b6101e560025481565b6101e56102c336600461170d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101e57f000000000000000000000000000000000000000000000000000000000000000081565b60085461026190610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101c6610767565b6103667f000000000000000000000000000000000000000000000000000000000000000081565b6040516101d3919061172a565b610206610381366004611687565b610776565b610399610394366004611832565b610783565b005b6103996103a9366004611832565b6107e0565b6102066103bc366004611883565b610839565b6103996103cf36600461170d565b610988565b6103996103e236600461170d565b610a9c565b6103ef610bb0565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101d3565b6101e56104293660046118dc565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b61039961046f3660046116f4565b610dce565b6000546102619073ffffffffffffffffffffffffffffffffffffffff1681565b6103996104a23660046116f4565b610e50565b6103ef6104b53660046116f4565b610ed2565b6060600680546104c990611915565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590611915565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b5050505050905090565b60006105593384846110ee565b5060015b92915050565b60006105708484846111f7565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054828110156105db576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ef85336105ea8685611997565b6110ee565b506001949350505050565b600061060633836113b2565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab1480159061064c575060005473ffffffffffffffffffffffffffffffffffffffff163314155b80156106755750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b156106ac576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106b683836114f4565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107569186906119aa565b60405180910390a250600192915050565b6060600780546104c990611915565b60006105593384846111f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107d0576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107dc8282611a1b565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461082d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107dc8282611a1b565b6000806000610846610bb0565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156108d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fc9190611b34565b610932576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61093c33856113b2565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161097591899189918791611b56565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109d5576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a22576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610ae9576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b36576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c679190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610cb6576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d699190611ba2565b905080600003610da5576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610db89190611bbb565b610dc29190611bd2565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e1b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a91565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e9d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f899190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610fd8576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015611067573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108b9190611ba2565b9050806000036110c7576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000906110d78784611bbb565b6110e19190611bd2565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661113b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611188576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611244576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611291576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040902054818110156112f1576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112fb8282611997565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260036020526040808220939093559085168152908120805484929061133e908490611bd2565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113a491815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113ff576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260409020548181101561145f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114698282611997565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812091909155600580548492906114a4908490611997565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111ea565b73ffffffffffffffffffffffffffffffffffffffff8216611541576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008282546115539190611bd2565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805483929061158d908490611bd2565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561160a576020818501810151868301820152016115ee565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061165b60208301846115e4565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461168457600080fd5b50565b6000806040838503121561169a57600080fd5b82356116a581611662565b946020939093013593505050565b6000806000606084860312156116c857600080fd5b83356116d381611662565b925060208401356116e381611662565b929592945050506040919091013590565b60006020828403121561170657600080fd5b5035919050565b60006020828403121561171f57600080fd5b813561165b81611662565b6020810160038310611765577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff8411156117b5576117b561176b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156118025761180261176b565b60405283815290508082840185101561181a57600080fd5b83836020830137600060208583010152509392505050565b60006020828403121561184457600080fd5b813567ffffffffffffffff81111561185b57600080fd5b8201601f8101841361186c57600080fd5b61187b8482356020840161179a565b949350505050565b6000806040838503121561189657600080fd5b823567ffffffffffffffff8111156118ad57600080fd5b8301601f810185136118be57600080fd5b6118cd8582356020840161179a565b95602094909401359450505050565b600080604083850312156118ef57600080fd5b82356118fa81611662565b9150602083013561190a81611662565b809150509250929050565b600181811c9082168061192957607f821691505b602082108103611962577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561055d5761055d611968565b6040815260006119bd60408301856115e4565b90508260208301529392505050565b601f821115611a1657806000526020600020601f840160051c810160208510156119f35750805b601f840160051c820191505b81811015611a1357600081556001016119ff565b50505b505050565b815167ffffffffffffffff811115611a3557611a3561176b565b611a4981611a438454611915565b846119cc565b6020601f821160018114611a9b5760008315611a655750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611a13565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611ae95787850151825560209485019460019092019101611ac9565b5084821015611b2557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b600060208284031215611b4657600080fd5b8151801515811461165b57600080fd5b608081526000611b6960808301876115e4565b6020830195909552506040810192909252606090910152919050565b600060208284031215611b9757600080fd5b815161165b81611662565b600060208284031215611bb457600080fd5b5051919050565b808202811582820484141761055d5761055d611968565b8082018082111561055d5761055d61196856fea2646970667358221220d6ba834f25782689ed13bffb6ac9ff2c8d3b5342c94a515aea8197a76070ad3f64736f6c634300081a0033a2646970667358221220770429cd9e1fb35df041de48926a5f5a7a26b6aa368aafba70ee0582f2a7fb9d64736f6c634300081a0033",
}

// GatewayEVMZEVMTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMZEVMTestMetaData.ABI instead.
var GatewayEVMZEVMTestABI = GatewayEVMZEVMTestMetaData.ABI

// GatewayEVMZEVMTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMZEVMTestMetaData.Bin instead.
var GatewayEVMZEVMTestBin = GatewayEVMZEVMTestMetaData.Bin

// DeployGatewayEVMZEVMTest deploys a new Ethereum contract, binding an instance of GatewayEVMZEVMTest to it.
func DeployGatewayEVMZEVMTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMZEVMTest, error) {
	parsed, err := GatewayEVMZEVMTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMZEVMTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMZEVMTest{GatewayEVMZEVMTestCaller: GatewayEVMZEVMTestCaller{contract: contract}, GatewayEVMZEVMTestTransactor: GatewayEVMZEVMTestTransactor{contract: contract}, GatewayEVMZEVMTestFilterer: GatewayEVMZEVMTestFilterer{contract: contract}}, nil
}

// GatewayEVMZEVMTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMZEVMTest struct {
	GatewayEVMZEVMTestCaller     // Read-only binding to the contract
	GatewayEVMZEVMTestTransactor // Write-only binding to the contract
	GatewayEVMZEVMTestFilterer   // Log filterer for contract events
}

// GatewayEVMZEVMTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMZEVMTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMZEVMTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMZEVMTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMZEVMTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMZEVMTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMZEVMTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMZEVMTestSession struct {
	Contract     *GatewayEVMZEVMTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GatewayEVMZEVMTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMZEVMTestCallerSession struct {
	Contract *GatewayEVMZEVMTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GatewayEVMZEVMTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMZEVMTestTransactorSession struct {
	Contract     *GatewayEVMZEVMTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GatewayEVMZEVMTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMZEVMTestRaw struct {
	Contract *GatewayEVMZEVMTest // Generic contract binding to access the raw methods on
}

// GatewayEVMZEVMTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMZEVMTestCallerRaw struct {
	Contract *GatewayEVMZEVMTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMZEVMTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMZEVMTestTransactorRaw struct {
	Contract *GatewayEVMZEVMTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMZEVMTest creates a new instance of GatewayEVMZEVMTest, bound to a specific deployed contract.
func NewGatewayEVMZEVMTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMZEVMTest, error) {
	contract, err := bindGatewayEVMZEVMTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTest{GatewayEVMZEVMTestCaller: GatewayEVMZEVMTestCaller{contract: contract}, GatewayEVMZEVMTestTransactor: GatewayEVMZEVMTestTransactor{contract: contract}, GatewayEVMZEVMTestFilterer: GatewayEVMZEVMTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMZEVMTestCaller creates a new read-only instance of GatewayEVMZEVMTest, bound to a specific deployed contract.
func NewGatewayEVMZEVMTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMZEVMTestCaller, error) {
	contract, err := bindGatewayEVMZEVMTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestCaller{contract: contract}, nil
}

// NewGatewayEVMZEVMTestTransactor creates a new write-only instance of GatewayEVMZEVMTest, bound to a specific deployed contract.
func NewGatewayEVMZEVMTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMZEVMTestTransactor, error) {
	contract, err := bindGatewayEVMZEVMTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestTransactor{contract: contract}, nil
}

// NewGatewayEVMZEVMTestFilterer creates a new log filterer instance of GatewayEVMZEVMTest, bound to a specific deployed contract.
func NewGatewayEVMZEVMTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMZEVMTestFilterer, error) {
	contract, err := bindGatewayEVMZEVMTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestFilterer{contract: contract}, nil
}

// bindGatewayEVMZEVMTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMZEVMTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMZEVMTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMZEVMTest.Contract.GatewayEVMZEVMTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.GatewayEVMZEVMTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.GatewayEVMZEVMTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMZEVMTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) ISTEST() (bool, error) {
	return _GatewayEVMZEVMTest.Contract.ISTEST(&_GatewayEVMZEVMTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) ISTEST() (bool, error) {
	return _GatewayEVMZEVMTest.Contract.ISTEST(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeArtifacts(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeArtifacts(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeContracts(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeContracts(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeSelectors(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeSelectors(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeSenders(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeSenders(&_GatewayEVMZEVMTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) Failed() (bool, error) {
	return _GatewayEVMZEVMTest.Contract.Failed(&_GatewayEVMZEVMTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) Failed() (bool, error) {
	return _GatewayEVMZEVMTest.Contract.Failed(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMZEVMTest.Contract.TargetArtifactSelectors(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMZEVMTest.Contract.TargetArtifactSelectors(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMZEVMTest.Contract.TargetArtifacts(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMZEVMTest.Contract.TargetArtifacts(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.TargetContracts(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.TargetContracts(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMZEVMTest.Contract.TargetInterfaces(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMZEVMTest.Contract.TargetInterfaces(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMZEVMTest.Contract.TargetSelectors(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMZEVMTest.Contract.TargetSelectors(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.TargetSenders(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.TargetSenders(&_GatewayEVMZEVMTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.SetUp(&_GatewayEVMZEVMTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.SetUp(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestCallReceiverEVMFromSenderZEVM is a paid mutator transaction binding the contract method 0xd7a525fc.
//
// Solidity: function testCallReceiverEVMFromSenderZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactor) TestCallReceiverEVMFromSenderZEVM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.contract.Transact(opts, "testCallReceiverEVMFromSenderZEVM")
}

// TestCallReceiverEVMFromSenderZEVM is a paid mutator transaction binding the contract method 0xd7a525fc.
//
// Solidity: function testCallReceiverEVMFromSenderZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TestCallReceiverEVMFromSenderZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestCallReceiverEVMFromSenderZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestCallReceiverEVMFromSenderZEVM is a paid mutator transaction binding the contract method 0xd7a525fc.
//
// Solidity: function testCallReceiverEVMFromSenderZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorSession) TestCallReceiverEVMFromSenderZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestCallReceiverEVMFromSenderZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestCallReceiverEVMFromZEVM is a paid mutator transaction binding the contract method 0x9683c695.
//
// Solidity: function testCallReceiverEVMFromZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactor) TestCallReceiverEVMFromZEVM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.contract.Transact(opts, "testCallReceiverEVMFromZEVM")
}

// TestCallReceiverEVMFromZEVM is a paid mutator transaction binding the contract method 0x9683c695.
//
// Solidity: function testCallReceiverEVMFromZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TestCallReceiverEVMFromZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestCallReceiverEVMFromZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestCallReceiverEVMFromZEVM is a paid mutator transaction binding the contract method 0x9683c695.
//
// Solidity: function testCallReceiverEVMFromZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorSession) TestCallReceiverEVMFromZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestCallReceiverEVMFromZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestWithdrawAndCallReceiverEVMFromSenderZEVM is a paid mutator transaction binding the contract method 0x52474413.
//
// Solidity: function testWithdrawAndCallReceiverEVMFromSenderZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactor) TestWithdrawAndCallReceiverEVMFromSenderZEVM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.contract.Transact(opts, "testWithdrawAndCallReceiverEVMFromSenderZEVM")
}

// TestWithdrawAndCallReceiverEVMFromSenderZEVM is a paid mutator transaction binding the contract method 0x52474413.
//
// Solidity: function testWithdrawAndCallReceiverEVMFromSenderZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TestWithdrawAndCallReceiverEVMFromSenderZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestWithdrawAndCallReceiverEVMFromSenderZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestWithdrawAndCallReceiverEVMFromSenderZEVM is a paid mutator transaction binding the contract method 0x52474413.
//
// Solidity: function testWithdrawAndCallReceiverEVMFromSenderZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorSession) TestWithdrawAndCallReceiverEVMFromSenderZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestWithdrawAndCallReceiverEVMFromSenderZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestWithdrawAndCallReceiverEVMFromZEVM is a paid mutator transaction binding the contract method 0x6ff15ccc.
//
// Solidity: function testWithdrawAndCallReceiverEVMFromZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactor) TestWithdrawAndCallReceiverEVMFromZEVM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.contract.Transact(opts, "testWithdrawAndCallReceiverEVMFromZEVM")
}

// TestWithdrawAndCallReceiverEVMFromZEVM is a paid mutator transaction binding the contract method 0x6ff15ccc.
//
// Solidity: function testWithdrawAndCallReceiverEVMFromZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TestWithdrawAndCallReceiverEVMFromZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestWithdrawAndCallReceiverEVMFromZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestWithdrawAndCallReceiverEVMFromZEVM is a paid mutator transaction binding the contract method 0x6ff15ccc.
//
// Solidity: function testWithdrawAndCallReceiverEVMFromZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorSession) TestWithdrawAndCallReceiverEVMFromZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestWithdrawAndCallReceiverEVMFromZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// GatewayEVMZEVMTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestCalledIterator struct {
	Event *GatewayEVMZEVMTestCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestCalled represents a Called event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMZEVMTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestCalledIterator{contract: _GatewayEVMZEVMTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestCalled)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Called", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseCalled(log types.Log) (*GatewayEVMZEVMTestCalled, error) {
	event := new(GatewayEVMZEVMTestCalled)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestCalled0Iterator is returned from FilterCalled0 and is used to iterate over the raw logs and unpacked data for Called0 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestCalled0Iterator struct {
	Event *GatewayEVMZEVMTestCalled0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestCalled0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestCalled0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestCalled0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestCalled0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestCalled0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestCalled0 represents a Called0 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestCalled0 struct {
	Sender        common.Address
	Zrc20         common.Address
	Receiver      []byte
	Message       []byte
	GasLimit      *big.Int
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled0 is a free log retrieval operation binding the contract event 0x6c6abd640fc6a0ef7cf2bc54b246b42d5c2629c30be1e24fea4a58157a7728cf.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterCalled0(opts *bind.FilterOpts, sender []common.Address, zrc20 []common.Address) (*GatewayEVMZEVMTestCalled0Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Called0", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestCalled0Iterator{contract: _GatewayEVMZEVMTest.contract, event: "Called0", logs: logs, sub: sub}, nil
}

// WatchCalled0 is a free log subscription operation binding the contract event 0x6c6abd640fc6a0ef7cf2bc54b246b42d5c2629c30be1e24fea4a58157a7728cf.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchCalled0(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestCalled0, sender []common.Address, zrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Called0", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestCalled0)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Called0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCalled0 is a log parse operation binding the contract event 0x6c6abd640fc6a0ef7cf2bc54b246b42d5c2629c30be1e24fea4a58157a7728cf.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseCalled0(log types.Log) (*GatewayEVMZEVMTestCalled0, error) {
	event := new(GatewayEVMZEVMTestCalled0)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Called0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestDepositedIterator struct {
	Event *GatewayEVMZEVMTestDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestDeposited represents a Deposited event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestDeposited struct {
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterDeposited(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMZEVMTestDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestDepositedIterator{contract: _GatewayEVMZEVMTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestDeposited, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestDeposited)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseDeposited(log types.Log) (*GatewayEVMZEVMTestDeposited, error) {
	event := new(GatewayEVMZEVMTestDeposited)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestExecutedIterator struct {
	Event *GatewayEVMZEVMTestExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestExecuted represents a Executed event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMZEVMTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestExecutedIterator{contract: _GatewayEVMZEVMTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestExecuted)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseExecuted(log types.Log) (*GatewayEVMZEVMTestExecuted, error) {
	event := new(GatewayEVMZEVMTestExecuted)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMZEVMTestExecutedWithERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestExecutedWithERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestExecutedWithERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMZEVMTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestExecutedWithERC20Iterator{contract: _GatewayEVMZEVMTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestExecutedWithERC20)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMZEVMTestExecutedWithERC20, error) {
	event := new(GatewayEVMZEVMTestExecutedWithERC20)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedERC20Iterator struct {
	Event *GatewayEVMZEVMTestReceivedERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReceivedERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestReceivedERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReceivedERC20 represents a ReceivedERC20 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*GatewayEVMZEVMTestReceivedERC20Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestReceivedERC20Iterator{contract: _GatewayEVMZEVMTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReceivedERC20)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReceivedERC20(log types.Log) (*GatewayEVMZEVMTestReceivedERC20, error) {
	event := new(GatewayEVMZEVMTestReceivedERC20)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedNoParamsIterator struct {
	Event *GatewayEVMZEVMTestReceivedNoParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReceivedNoParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestReceivedNoParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReceivedNoParams represents a ReceivedNoParams event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*GatewayEVMZEVMTestReceivedNoParamsIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestReceivedNoParamsIterator{contract: _GatewayEVMZEVMTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReceivedNoParams)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReceivedNoParams(log types.Log) (*GatewayEVMZEVMTestReceivedNoParams, error) {
	event := new(GatewayEVMZEVMTestReceivedNoParams)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedNonPayableIterator struct {
	Event *GatewayEVMZEVMTestReceivedNonPayable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReceivedNonPayable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestReceivedNonPayable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReceivedNonPayable represents a ReceivedNonPayable event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*GatewayEVMZEVMTestReceivedNonPayableIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestReceivedNonPayableIterator{contract: _GatewayEVMZEVMTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReceivedNonPayable)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReceivedNonPayable(log types.Log) (*GatewayEVMZEVMTestReceivedNonPayable, error) {
	event := new(GatewayEVMZEVMTestReceivedNonPayable)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedPayableIterator struct {
	Event *GatewayEVMZEVMTestReceivedPayable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReceivedPayable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestReceivedPayable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReceivedPayable represents a ReceivedPayable event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedPayable struct {
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*GatewayEVMZEVMTestReceivedPayableIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestReceivedPayableIterator{contract: _GatewayEVMZEVMTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReceivedPayable)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReceivedPayable(log types.Log) (*GatewayEVMZEVMTestReceivedPayable, error) {
	event := new(GatewayEVMZEVMTestReceivedPayable)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedRevertIterator struct {
	Event *GatewayEVMZEVMTestReceivedRevert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReceivedRevert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestReceivedRevert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReceivedRevert represents a ReceivedRevert event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*GatewayEVMZEVMTestReceivedRevertIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestReceivedRevertIterator{contract: _GatewayEVMZEVMTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e.
//
// Solidity: event ReceivedRevert(address sender, (address,uint64,bytes) revertContext)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReceivedRevert)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReceivedRevert(log types.Log) (*GatewayEVMZEVMTestReceivedRevert, error) {
	event := new(GatewayEVMZEVMTestReceivedRevert)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestRevertedIterator struct {
	Event *GatewayEVMZEVMTestReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReverted represents a Reverted event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReverted struct {
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMZEVMTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestRevertedIterator{contract: _GatewayEVMZEVMTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0x1647880d5adf27692c774debf932b888ba15381ed4c4115a9eafeb68080b1436.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,uint64,bytes) revertContext)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReverted)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReverted(log types.Log) (*GatewayEVMZEVMTestReverted, error) {
	event := new(GatewayEVMZEVMTestReverted)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator struct {
	Event *GatewayEVMZEVMTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestUpdatedGatewayTSSAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestUpdatedGatewayTSSAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestUpdatedGatewayTSSAddress struct {
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator{contract: _GatewayEVMZEVMTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x7598d084f3a8d9a71847119f6fdb694046bc0aaab0dee775c33c1df9be089a05.
//
// Solidity: event UpdatedGatewayTSSAddress(address newTSSAddress)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestUpdatedGatewayTSSAddress)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*GatewayEVMZEVMTestUpdatedGatewayTSSAddress, error) {
	event := new(GatewayEVMZEVMTestUpdatedGatewayTSSAddress)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestWithdrawnIterator struct {
	Event *GatewayEVMZEVMTestWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestWithdrawn represents a Withdrawn event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestWithdrawn struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Zrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	GasLimit        *big.Int
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xda1215b0949ddb309fe466fa9e70e861a16538f11b8ecdb05c217d4d8677ed2d.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayEVMZEVMTestWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestWithdrawnIterator{contract: _GatewayEVMZEVMTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xda1215b0949ddb309fe466fa9e70e861a16538f11b8ecdb05c217d4d8677ed2d.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestWithdrawn)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0xda1215b0949ddb309fe466fa9e70e861a16538f11b8ecdb05c217d4d8677ed2d.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseWithdrawn(log types.Log) (*GatewayEVMZEVMTestWithdrawn, error) {
	event := new(GatewayEVMZEVMTestWithdrawn)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogIterator struct {
	Event *GatewayEVMZEVMTestLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLog represents a Log event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogIterator{contract: _GatewayEVMZEVMTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLog)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLog(log types.Log) (*GatewayEVMZEVMTestLog, error) {
	event := new(GatewayEVMZEVMTestLog)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogAddressIterator struct {
	Event *GatewayEVMZEVMTestLogAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogAddress represents a LogAddress event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogAddressIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogAddressIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogAddress)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogAddress(log types.Log) (*GatewayEVMZEVMTestLogAddress, error) {
	event := new(GatewayEVMZEVMTestLogAddress)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogArrayIterator struct {
	Event *GatewayEVMZEVMTestLogArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogArray represents a LogArray event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogArrayIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogArrayIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogArray)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogArray(log types.Log) (*GatewayEVMZEVMTestLogArray, error) {
	event := new(GatewayEVMZEVMTestLogArray)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogArray0Iterator struct {
	Event *GatewayEVMZEVMTestLogArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogArray0 represents a LogArray0 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogArray0Iterator{contract: _GatewayEVMZEVMTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogArray0)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogArray0(log types.Log) (*GatewayEVMZEVMTestLogArray0, error) {
	event := new(GatewayEVMZEVMTestLogArray0)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogArray1Iterator struct {
	Event *GatewayEVMZEVMTestLogArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogArray1 represents a LogArray1 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogArray1Iterator{contract: _GatewayEVMZEVMTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogArray1)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogArray1(log types.Log) (*GatewayEVMZEVMTestLogArray1, error) {
	event := new(GatewayEVMZEVMTestLogArray1)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogBytesIterator struct {
	Event *GatewayEVMZEVMTestLogBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogBytes represents a LogBytes event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogBytesIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogBytesIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogBytes)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogBytes(log types.Log) (*GatewayEVMZEVMTestLogBytes, error) {
	event := new(GatewayEVMZEVMTestLogBytes)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogBytes32Iterator struct {
	Event *GatewayEVMZEVMTestLogBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogBytes32 represents a LogBytes32 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogBytes32Iterator{contract: _GatewayEVMZEVMTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogBytes32)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogBytes32(log types.Log) (*GatewayEVMZEVMTestLogBytes32, error) {
	event := new(GatewayEVMZEVMTestLogBytes32)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogIntIterator struct {
	Event *GatewayEVMZEVMTestLogInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogInt represents a LogInt event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogIntIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogIntIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogInt)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogInt(log types.Log) (*GatewayEVMZEVMTestLogInt, error) {
	event := new(GatewayEVMZEVMTestLogInt)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedAddressIterator struct {
	Event *GatewayEVMZEVMTestLogNamedAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogNamedAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedAddressIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedAddress)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayEVMZEVMTestLogNamedAddress, error) {
	event := new(GatewayEVMZEVMTestLogNamedAddress)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedArrayIterator struct {
	Event *GatewayEVMZEVMTestLogNamedArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogNamedArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedArray represents a LogNamedArray event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedArrayIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedArray)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayEVMZEVMTestLogNamedArray, error) {
	event := new(GatewayEVMZEVMTestLogNamedArray)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedArray0Iterator struct {
	Event *GatewayEVMZEVMTestLogNamedArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogNamedArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedArray0Iterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedArray0)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayEVMZEVMTestLogNamedArray0, error) {
	event := new(GatewayEVMZEVMTestLogNamedArray0)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedArray1Iterator struct {
	Event *GatewayEVMZEVMTestLogNamedArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogNamedArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedArray1Iterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedArray1)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayEVMZEVMTestLogNamedArray1, error) {
	event := new(GatewayEVMZEVMTestLogNamedArray1)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedBytesIterator struct {
	Event *GatewayEVMZEVMTestLogNamedBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogNamedBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedBytesIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedBytes)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayEVMZEVMTestLogNamedBytes, error) {
	event := new(GatewayEVMZEVMTestLogNamedBytes)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedBytes32Iterator struct {
	Event *GatewayEVMZEVMTestLogNamedBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogNamedBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedBytes32Iterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedBytes32)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayEVMZEVMTestLogNamedBytes32, error) {
	event := new(GatewayEVMZEVMTestLogNamedBytes32)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedDecimalIntIterator struct {
	Event *GatewayEVMZEVMTestLogNamedDecimalInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedDecimalInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogNamedDecimalInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedDecimalIntIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedDecimalInt)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayEVMZEVMTestLogNamedDecimalInt, error) {
	event := new(GatewayEVMZEVMTestLogNamedDecimalInt)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedDecimalUintIterator struct {
	Event *GatewayEVMZEVMTestLogNamedDecimalUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedDecimalUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogNamedDecimalUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedDecimalUintIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedDecimalUint)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayEVMZEVMTestLogNamedDecimalUint, error) {
	event := new(GatewayEVMZEVMTestLogNamedDecimalUint)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedIntIterator struct {
	Event *GatewayEVMZEVMTestLogNamedInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogNamedInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedInt represents a LogNamedInt event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedIntIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedInt)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayEVMZEVMTestLogNamedInt, error) {
	event := new(GatewayEVMZEVMTestLogNamedInt)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedStringIterator struct {
	Event *GatewayEVMZEVMTestLogNamedString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogNamedString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedString represents a LogNamedString event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedStringIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedString)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedString(log types.Log) (*GatewayEVMZEVMTestLogNamedString, error) {
	event := new(GatewayEVMZEVMTestLogNamedString)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedUintIterator struct {
	Event *GatewayEVMZEVMTestLogNamedUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogNamedUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedUint represents a LogNamedUint event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedUintIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedUint)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayEVMZEVMTestLogNamedUint, error) {
	event := new(GatewayEVMZEVMTestLogNamedUint)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogStringIterator struct {
	Event *GatewayEVMZEVMTestLogString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogString represents a LogString event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogStringIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogStringIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogString)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogString(log types.Log) (*GatewayEVMZEVMTestLogString, error) {
	event := new(GatewayEVMZEVMTestLogString)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogUintIterator struct {
	Event *GatewayEVMZEVMTestLogUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogUint represents a LogUint event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogUintIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogUintIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogUint)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogUint(log types.Log) (*GatewayEVMZEVMTestLogUint, error) {
	event := new(GatewayEVMZEVMTestLogUint)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogsIterator struct {
	Event *GatewayEVMZEVMTestLogs // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogs)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestLogs)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogs represents a Logs event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogsIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogsIterator{contract: _GatewayEVMZEVMTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogs)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "logs", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogs(log types.Log) (*GatewayEVMZEVMTestLogs, error) {
	event := new(GatewayEVMZEVMTestLogs)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
