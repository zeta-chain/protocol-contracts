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
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testCallReceiverEVMFromSenderZEVM\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallReceiverEVMFromZEVM\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiverEVMFromSenderZEVM\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiverEVMFromZEVM\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotFungibleModule\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrFungible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5062011b42806200003e6000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806385226c81116100b2578063b5508aa911610081578063d7a525fc11610066578063d7a525fc146101ec578063e20c9f71146101f4578063fa7626d4146101fc57600080fd5b8063b5508aa9146101cc578063ba414fa6146101d457600080fd5b806385226c8114610192578063916a17c6146101a75780639683c695146101bc578063b0464fdc146101c457600080fd5b80633f7286f4116100ee5780633f7286f414610165578063524744131461016d57806366d9a9a0146101755780636ff15ccc1461018a57600080fd5b80630a9254e4146101205780631ed7831c1461012a5780632ade3880146101485780633e5e3c231461015d575b600080fd5b610128610209565b005b61013261125d565b60405161013f9190617ae3565b60405180910390f35b6101506112bf565b60405161013f9190617b7f565b610132611401565b610132611461565b6101286114c1565b61017d611d30565b60405161013f9190617ce5565b610128611eb2565b61019a6126d5565b60405161013f9190617d83565b6101af6127a5565b60405161013f9190617dfa565b6101286128a0565b6101af612e94565b61019a612f8f565b6101dc61305f565b604051901515815260200161013f565b610128613133565b610132613857565b601f546101dc9060ff1681565b602680547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602780548216611234179055602880548216615678179055602e8054909116614321179055604051610267906179c8565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156102ec573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055604051610331906179c8565b604080825260049082018190527f7a6574610000000000000000000000000000000000000000000000000000000060608301526080602083018190528201527f5a4554410000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156103b5573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081178255604080518082018252600e81527f4761746577617945564d2e736f6c0000000000000000000000000000000000006020820152602854602654925190861694810194909452604484019290925290921660648201526104a6919060840160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b000000000000000000000000000000000000000000000000000000001790526138b7565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff00000000000000000000000000000000000000009091168117909155602854602654604051929391821692911690610532906179d6565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f08015801561056e573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392831617905560205460245460285460265460405193851694928316939183169216906105c9906179e4565b6001600160a01b039485168152928416602084015290831660408301529091166060820152608001604051809103906000f08015801561060d573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556028546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b1580156106b957600080fd5b505af11580156106cd573d6000803e3d6000fd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561074357600080fd5b505af1158015610757573d6000803e3d6000fd5b50506020546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f9150602401600060405180830381600087803b1580156107bd57600080fd5b505af11580156107d1573d6000803e3d6000fd5b50506020546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef9150602401600060405180830381600087803b15801561083757600080fd5b505af115801561084b573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156108ad57600080fd5b505af11580156108c1573d6000803e3d6000fd5b50506023546026546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42406024820152911692506340c10f199150604401600060405180830381600087803b15801561093057600080fd5b505af1158015610944573d6000803e3d6000fd5b50506023546021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526207a12060248201529116925063a9059cbb91506044016020604051808303816000875af11580156109b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109dc9190617e91565b506040516109e9906179f2565b604051809103906000f080158015610a05573d6000803e3d6000fd5b50602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316179055604080518082018252600f81527f476174657761795a45564d2e736f6c0000000000000000000000000000000000602082015260248054602e54935190851691810191909152919092166044820152610aee919060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f485cc955000000000000000000000000000000000000000000000000000000001790526138b7565b602980546001600160a01b03929092167fffffffffffffffffffffffff00000000000000000000000000000000000000009283168117909155602a80549092168117909155604051610b3f90617a00565b6001600160a01b039091168152602001604051809103906000f080158015610b6b573d6000803e3d6000fd5b50602b80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03929092169190911790556040517f06447d5600000000000000000000000000000000000000000000000000000000815273735b14bb79463307aacbed86daf3322b1e6226ab6004820181905290737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015610c2057600080fd5b505af1158015610c34573d6000803e3d6000fd5b505050506000806000604051610c4990617a0e565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610c85573d6000803e3d6000fd5b50602c80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602a54604051601293600193600093849391921690610cdb90617a1c565b610cea96959493929190617eb3565b604051809103906000f080158015610d06573d6000803e3d6000fd5b50602d80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602c546040517fee2815ba0000000000000000000000000000000000000000000000000000000081526001600482015260248101929092529091169063ee2815ba90604401600060405180830381600087803b158015610d9d57600080fd5b505af1158015610db1573d6000803e3d6000fd5b5050602c546040517fa7cb050700000000000000000000000000000000000000000000000000000000815260016004820181905260248201526001600160a01b03909116925063a7cb05079150604401600060405180830381600087803b158015610e1b57600080fd5b505af1158015610e2f573d6000803e3d6000fd5b5050602d54602e546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42406024820152911692506347e7ef2491506044016020604051808303816000875af1158015610ea3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec79190617e91565b50602d54602b546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424060248201529116906347e7ef24906044016020604051808303816000875af1158015610f38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5c9190617e91565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610fbb57600080fd5b505af1158015610fcf573d6000803e3d6000fd5b5050602e546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561104557600080fd5b505af1158015611059573d6000803e3d6000fd5b5050602d54602a546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424060248201529116925063095ea7b391506044016020604051808303816000875af11580156110cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110f19190617e91565b506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b15801561117257600080fd5b505af1158015611186573d6000803e3d6000fd5b50506040805160808101825261032180825260016020808401918252838501928352845190810190945260008452606083018490528251602f80549251151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009093166001600160a01b0392831617929092178255915160308054919093167fffffffffffffffffffffffff0000000000000000000000000000000000000000919091161790915590935091506031906112579082618073565b50505050565b606060168054806020026020016040519081016040528092919081815260200182805480156112b557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611297575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020016000905b828210156113f857600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156113e157838290600052602060002001805461135490617fd7565b80601f016020809104026020016040519081016040528092919081815260200182805461138090617fd7565b80156113cd5780601f106113a2576101008083540402835291602001916113cd565b820191906000526020600020905b8154815290600101906020018083116113b057829003601f168201915b505050505081526020019060010190611335565b5050505081525050815260200190600101906112e3565b50505050905090565b606060188054806020026020016040519081016040528092919081815260200182805480156112b5576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611297575050505050905090565b606060178054806020026020016040519081016040528092919081815260200182805480156112b5576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611297575050505050905090565b604080518082018252600681527f48656c6c6f2100000000000000000000000000000000000000000000000000006020820152602d54602b5492517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0393841660048201529192602a92600192670de0b6b3a7640000926000929116906370a0823190602401602060405180830381865afa15801561156d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115919190618132565b6040519091506000907fe04d4f9700000000000000000000000000000000000000000000000000000000906115ce9088908890889060240161814b565b60408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909516949094179093526025549051919350600092611667926001600160a01b03909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602d5461169e926207a120916001600160a01b0316908690600190602f90602401618258565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f90ad3e2300000000000000000000000000000000000000000000000000000000179052602a5490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161175b916001600160a01b03919091169060009086906004016182c1565b600060405180830381600087803b15801561177557600080fd5b505af1158015611789573d6000803e3d6000fd5b5050602e546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156117ff57600080fd5b505af1158015611813573d6000803e3d6000fd5b5050602b5460255460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039091169250630abd8905915060340160408051601f1981840301815290829052602d547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526118ad926207a120916001600160a01b0316908d908d908d906004016182e9565b600060405180830381600087803b1580156118c757600080fd5b505af11580156118db573d6000803e3d6000fd5b50506020546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260248101879052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b15801561195857600080fd5b505af115801561196c573d6000803e3d6000fd5b50506025546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156119fe57600080fd5b505af1158015611a12573d6000803e3d6000fd5b50506020546040517f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa9350611a5b92506001600160a01b039091169087908b908b908b9061833e565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611af157600080fd5b505af1158015611b05573d6000803e3d6000fd5b50506025546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f9150611b4a908790869061837f565b60405180910390a26028546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611bc457600080fd5b505af1158015611bd8573d6000803e3d6000fd5b50506020546025546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd93508892611c2c9216908790600401618398565b60006040518083038185885af1158015611c4a573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052611c739190810190618473565b50602d54602b546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611cde573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d029190618132565b9050611d26816001611d176207a120886184d7565b611d2191906184d7565b6138d6565b5050505050505050565b6060601b805480602002602001604051908101604052809291908181526020016000905b828210156113f85783829060005260206000209060020201604051806040016040529081600082018054611d8790617fd7565b80601f0160208091040260200160405190810160405280929190818152602001828054611db390617fd7565b8015611e005780601f10611dd557610100808354040283529160200191611e00565b820191906000526020600020905b815481529060010190602001808311611de357829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015611e9a57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411611e475790505b50505050508152505081526020019060010190611d54565b604080518082018252600681527f48656c6c6f21000000000000000000000000000000000000000000000000000060208201529051602a90600190670de0b6b3a7640000906000907fe04d4f970000000000000000000000000000000000000000000000000000000090611f2e9087908790879060240161814b565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009490941693909317909252602a5491517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190526024820181905260448201819052606482018190526001600160a01b039093166084820152909250737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561202057600080fd5b505af1158015612034573d6000803e3d6000fd5b5050602e5460255460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507ff6e8f2c19d797cd515ed61d8ac909afa628f4df0bb153a3f81d0f5fb7b8207c09060340160408051601f19818403018152828252602d547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b03909216916207a1209188918491634d8943bb916004808201926020929091908290030181865afa158015612109573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061212d9190618132565b896001602f6040516121469897969594939291906184ea565b60405180910390a3602e546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156121c057600080fd5b505af11580156121d4573d6000803e3d6000fd5b5050602a5460255460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b0390911692506390ad3e23915060340160408051601f1981840301815290829052602d547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352612270926207a120916001600160a01b0316908890600190602f9060040161854c565b600060405180830381600087803b15801561228a57600080fd5b505af115801561229e573d6000803e3d6000fd5b5050602d54602e546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561230a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061232e9190618132565b905061234181611d21846207a1206184d7565b6020546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260248101859052737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b1580156123ba57600080fd5b505af11580156123ce573d6000803e3d6000fd5b50506025546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561246057600080fd5b505af1158015612474573d6000803e3d6000fd5b50506020546040517f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa93506124bd92506001600160a01b039091169087908b908b908b9061833e565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561255357600080fd5b505af1158015612567573d6000803e3d6000fd5b50506025546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f91506125ac908790879061837f565b60405180910390a26028546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561262657600080fd5b505af115801561263a573d6000803e3d6000fd5b50506020546025546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd9350889261268e9216908890600401618398565b60006040518083038185885af11580156126ac573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052611d269190810190618473565b6060601a805480602002602001604051908101604052809291908181526020016000905b828210156113f857838290600052602060002001805461271890617fd7565b80601f016020809104026020016040519081016040528092919081815260200182805461274490617fd7565b80156127915780601f1061276657610100808354040283529160200191612791565b820191906000526020600020905b81548152906001019060200180831161277457829003601f168201915b5050505050815260200190600101906126f9565b6060601d805480602002602001604051908101604052809291908181526020016000905b828210156113f85760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561288857602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116128355790505b505050505081525050815260200190600101906127c9565b604080518082018252600681527f48656c6c6f21000000000000000000000000000000000000000000000000000060208201529051602a90600190670de0b6b3a7640000906000907fe04d4f97000000000000000000000000000000000000000000000000000000009061291c9087908790879060240161814b565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009490941693909317909252602e5491517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b0390921660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156129ef57600080fd5b505af1158015612a03573d6000803e3d6000fd5b5050602a546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015612a9557600080fd5b505af1158015612aa9573d6000803e3d6000fd5b5050602d54602e5460255460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039283169450911691507f885ccd61b2cc8be1daa203d4456750151d128d2da2ace4c9fa4b63c395fd99d89060340160408051601f1981840301815290829052612b29918690600190602f906185a0565b60405180910390a3602a546025546040805160609290921b6bffffffffffffffffffffffff19166020830152805180830360140181526034830191829052602d547fdc9ca2e7000000000000000000000000000000000000000000000000000000009092526001600160a01b039384169363dc9ca2e793612bb99391909116908690600190602f906038016185df565b600060405180830381600087803b158015612bd357600080fd5b505af1158015612be7573d6000803e3d6000fd5b50506020546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260248101859052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b158015612c6457600080fd5b505af1158015612c78573d6000803e3d6000fd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015612d0a57600080fd5b505af1158015612d1e573d6000803e3d6000fd5b50506025546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f9150612d63908590859061837f565b60405180910390a26028546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612ddd57600080fd5b505af1158015612df1573d6000803e3d6000fd5b50506020546025546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd93508692612e459216908690600401618398565b60006040518083038185885af1158015612e63573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052612e8c9190810190618473565b505050505050565b6060601c805480602002602001604051908101604052809291908181526020016000905b828210156113f85760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015612f7757602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411612f245790505b50505050508152505081526020019060010190612eb8565b60606019805480602002602001604051908101604052809291908181526020016000905b828210156113f8578382906000526020600020018054612fd290617fd7565b80601f0160208091040260200160405190810160405280929190818152602001828054612ffe90617fd7565b801561304b5780601f106130205761010080835404028352916020019161304b565b820191906000526020600020905b81548152906001019060200180831161302e57829003601f168201915b505050505081526020019060010190612fb3565b60085460009060ff1615613077575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015613108573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061312c9190618132565b1415905090565b604080518082018252600681527f48656c6c6f21000000000000000000000000000000000000000000000000000060208201529051602a90600190670de0b6b3a7640000906000907fe04d4f9700000000000000000000000000000000000000000000000000000000906131af9087908790879060240161814b565b60408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909516949094179093526025549051919350600092613248926001600160a01b03909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602d5461327c926001600160a01b03909116908590600190602f90602401618639565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fdc9ca2e700000000000000000000000000000000000000000000000000000000179052602a5490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391613339916001600160a01b03919091169060009086906004016182c1565b600060405180830381600087803b15801561335357600080fd5b505af1158015613367573d6000803e3d6000fd5b5050602e546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156133dd57600080fd5b505af11580156133f1573d6000803e3d6000fd5b5050602b5460255460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039091169250637a34d8bb915060340160408051601f1981840301815290829052602d547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352613488926001600160a01b03909116908b908b908b9060040161868a565b600060405180830381600087803b1580156134a257600080fd5b505af11580156134b6573d6000803e3d6000fd5b50506020546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260248101869052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b15801561353357600080fd5b505af1158015613547573d6000803e3d6000fd5b50506025546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156135d957600080fd5b505af11580156135ed573d6000803e3d6000fd5b50506020546040517f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa935061363692506001600160a01b039091169086908a908a908a9061833e565b60405180910390a16020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156136cc57600080fd5b505af11580156136e0573d6000803e3d6000fd5b50506025546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f9150613725908690869061837f565b60405180910390a26028546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561379f57600080fd5b505af11580156137b3573d6000803e3d6000fd5b50506020546025546040517f1cff79cd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631cff79cd935087926138079216908790600401618398565b60006040518083038185885af1158015613825573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f1916820160405261384e9190810190618473565b50505050505050565b606060158054806020026020016040519081016040528092919081815260200182805480156112b5576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611297575050505050905090565b60006138c1617a2a565b6138cc848483613955565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c549060440160006040518083038186803b15801561394157600080fd5b505afa158015612e8c573d6000803e3d6000fd5b60008061396285846139d0565b90506139c56040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016139b0929190618398565b604051602081830303815290604052856139dc565b9150505b9392505050565b60006139c98383613a0a565b60c08101515160009015613a00576139f984848460c00151613a25565b90506139c9565b6139f98484613bcb565b6000613a168383613cb6565b6139c9838360200151846139dc565b600080613a30613cc6565b90506000613a3e8683613d99565b90506000613a55826060015183602001518561423f565b90506000613a6583838989614451565b90506000613a72826152ce565b602081015181519192509060030b15613ae557898260400151604051602001613a9c9291906186d8565b60408051601f19818403018152908290527f08c379a0000000000000000000000000000000000000000000000000000000008252613adc91600401618759565b60405180910390fd5b6000613b286040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a20000000000000000000000081525083600161549d565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d90613b7b908490600401618759565b602060405180830381865afa158015613b98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613bbc919061876c565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc92590613c20908790600401618759565b600060405180830381865afa158015613c3d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613c659190810190618473565b90506000613c938285604051602001613c7f929190618795565b60405160208183030381529060405261569d565b90506001600160a01b0381166138cc578484604051602001613a9c9291906187c4565b613cc2828260006156b0565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c90613d4d90849060040161886f565b600060405180830381865afa158015613d6a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613d9291908101906188b6565b9250505090565b613dcb6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d9050613e166040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b613e1f856157b3565b60208201526000613e2f86615b98565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa158015613e71573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613e9991908101906188b6565b86838560200151604051602001613eb394939291906188ff565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb1190613f0b908590600401618759565b600060405180830381865afa158015613f28573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613f5091908101906188b6565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f690613f98908490600401618a03565b602060405180830381865afa158015613fb5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613fd99190617e91565b613fee5781604051602001613a9c9190618a55565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890614033908490600401618ae7565b600060405180830381865afa158015614050573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261407891908101906188b6565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906140bf908490600401618b39565b602060405180830381865afa1580156140dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141009190617e91565b15614195576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061414a908490600401618b39565b600060405180830381865afa158015614167573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261418f91908101906188b6565b60408501525b846001600160a01b03166349c4fac88286600001516040516020016141ba9190618b8b565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016141e6929190618bf7565b600060405180830381865afa158015614203573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261422b91908101906188b6565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b606081526020019060019003908161425b5790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250816000815181106142bb576142bb618c1c565b60200260200101819052506040518060400160405280600381526020017f2d726c00000000000000000000000000000000000000000000000000000000008152508160018151811061430f5761430f618c1c565b60200260200101819052508460405160200161432b9190618c4b565b6040516020818303038152906040528160028151811061434d5761434d618c1c565b6020026020010181905250826040516020016143699190618cb7565b6040516020818303038152906040528160038151811061438b5761438b618c1c565b602002602001018190525060006143a1826152ce565b602080820151604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000081850190815282518084018452600080825290860152825180840190935290518252928101929092529192506144329060408051808201825260008082526020918201528151808301909252845182528085019082015290615e1b565b6144475785604051602001613a9c9190618cf8565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d90156144a1565b511590565b6146155782602001511561455d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401613adc565b8260c0015115614615576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401613adc565b6040805160ff8082526120008201909252600091816020015b606081526020019060019003908161462e57905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061468990618d89565b935060ff168151811061469e5761469e618c1c565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e37000000000000000000000000000000000000008152506040516020016146ef9190618da8565b60405160208183030381529060405282828061470a90618d89565b935060ff168151811061471f5761471f618c1c565b60200260200101819052506040518060400160405280600681526020017f6465706c6f79000000000000000000000000000000000000000000000000000081525082828061476c90618d89565b935060ff168151811061478157614781618c1c565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806147ce90618d89565b935060ff16815181106147e3576147e3618c1c565b602002602001018190525087602001518282806147ff90618d89565b935060ff168151811061481457614814618c1c565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163745061746800000000000000000000000000000000000081525082828061486190618d89565b935060ff168151811061487657614876618c1c565b60209081029190910101528751828261488e81618d89565b935060ff16815181106148a3576148a3618c1c565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e496400000000000000000000000000000000000000000000008152508282806148f090618d89565b935060ff168151811061490557614905618c1c565b602002602001018190525061491946615e7c565b828261492481618d89565b935060ff168151811061493957614939618c1c565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c65000000000000000000000000000000000081525082828061498690618d89565b935060ff168151811061499b5761499b618c1c565b6020026020010181905250868282806149b390618d89565b935060ff16815181106149c8576149c8618c1c565b6020908102919091010152855115614aef5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f6465000000000000000000000060208201528282614a1981618d89565b935060ff1681518110614a2e57614a2e618c1c565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d90614a7e908990600401618759565b600060405180830381865afa158015614a9b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052614ac391908101906188b6565b8282614ace81618d89565b935060ff1681518110614ae357614ae3618c1c565b60200260200101819052505b846020015115614bbf5760408051808201909152601281527f2d2d766572696679536f75726365436f6465000000000000000000000000000060208201528282614b3881618d89565b935060ff1681518110614b4d57614b4d618c1c565b60200260200101819052506040518060400160405280600581526020017f66616c7365000000000000000000000000000000000000000000000000000000815250828280614b9a90618d89565b935060ff1681518110614baf57614baf618c1c565b6020026020010181905250614d86565b614bf761449c8660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b614c8a5760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282614c3a81618d89565b935060ff1681518110614c4f57614c4f618c1c565b60200260200101819052508460a00151604051602001614c6f9190618c4b565b604051602081830303815290604052828280614b9a90618d89565b8460c00151158015614ccd575060408089015181518083018352600080825260209182015282518084019093528151835290810190820152614ccb90511590565b155b15614d865760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282614d1181618d89565b935060ff1681518110614d2657614d26618c1c565b6020026020010181905250614d3a88615f1c565b604051602001614d4a9190618c4b565b604051602081830303815290604052828280614d6590618d89565b935060ff1681518110614d7a57614d7a618c1c565b60200260200101819052505b60408086015181518083018352600080825260209182015282518084019093528151835290810190820152614dba90511590565b614e4f5760408051808201909152600b81527f2d2d72656c61796572496400000000000000000000000000000000000000000060208201528282614dfd81618d89565b935060ff1681518110614e1257614e12618c1c565b60200260200101819052508460400151828280614e2e90618d89565b935060ff1681518110614e4357614e43618c1c565b60200260200101819052505b606085015115614f705760408051808201909152600681527f2d2d73616c74000000000000000000000000000000000000000000000000000060208201528282614e9881618d89565b935060ff1681518110614ead57614ead618c1c565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015614f1c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052614f4491908101906188b6565b8282614f4f81618d89565b935060ff1681518110614f6457614f64618c1c565b60200260200101819052505b60e085015151156150175760408051808201909152600a81527f2d2d6761734c696d69740000000000000000000000000000000000000000000060208201528282614fba81618d89565b935060ff1681518110614fcf57614fcf618c1c565b6020026020010181905250614feb8560e0015160000151615e7c565b8282614ff681618d89565b935060ff168151811061500b5761500b618c1c565b60200260200101819052505b60e085015160200151156150c15760408051808201909152600a81527f2d2d6761735072696365000000000000000000000000000000000000000000006020820152828261506481618d89565b935060ff168151811061507957615079618c1c565b60200260200101819052506150958560e0015160200151615e7c565b82826150a081618d89565b935060ff16815181106150b5576150b5618c1c565b60200260200101819052505b60e0850151604001511561516b5760408051808201909152600e81527f2d2d6d61784665655065724761730000000000000000000000000000000000006020820152828261510e81618d89565b935060ff168151811061512357615123618c1c565b602002602001018190525061513f8560e0015160400151615e7c565b828261514a81618d89565b935060ff168151811061515f5761515f618c1c565b60200260200101819052505b60e085015160600151156152155760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826151b881618d89565b935060ff16815181106151cd576151cd618c1c565b60200260200101819052506151e98560e0015160600151615e7c565b82826151f481618d89565b935060ff168151811061520957615209618c1c565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561523357615233617fa8565b60405190808252806020026020018201604052801561526657816020015b60608152602001906001900390816152515790505b50905060005b8260ff168160ff1610156152bf57838160ff168151811061528f5761528f618c1c565b6020026020010151828260ff16815181106152ac576152ac618c1c565b602090810291909101015260010161526c565b5093505050505b949350505050565b6152f56040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c9161537b91869101618e13565b600060405180830381865afa158015615398573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526153c091908101906188b6565b905060006153ce8683616a0b565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b81526004016153fe9190617d83565b6000604051808303816000875af115801561541d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526154459190810190618e5a565b805190915060030b1580159061545e5750602081015151155b801561546d5750604081015151155b15614447578160008151811061548557615485618c1c565b6020026020010151604051602001613a9c9190618f10565b606060006154d28560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925286518252808701908201529091506155099082905b90616b60565b15615666576000615586826155808461557a61554c8a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b90616b87565b90616be9565b604080518082018252600181527f0a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506155ea908290616b60565b1561565457604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615651905b8290616c6e565b90505b61565d81616c94565b925050506139c9565b821561567f578484604051602001613a9c9291906190fc565b50506040805160208101909152600081526139c9565b509392505050565b6000808251602084016000f09392505050565b8160a00151156156bf57505050565b60006156cc848484616cfd565b905060006156d9826152ce565b602081015181519192509060030b1580156157755750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261577590604080518082018252600080825260209182015281518083019092528451825280850190820152615503565b1561578257505050505050565b604082015151156157a2578160400151604051602001613a9c91906191a3565b80604051602001613a9c9190619201565b606060006157e88360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061584d905b8290615e1b565b156158bc57604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526139c9906158b7908390617298565b616c94565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261591e905b8290617322565b6001036159eb57604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526159849061564a565b50604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526139c9906158b7905b8390616c6e565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615a4a90615846565b15615b8157604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290615ab29083906173bc565b905060008160018351615ac591906184d7565b81518110615ad557615ad5618c1c565b60200260200101519050615b786158b7615b4b6040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252855182528086019082015290617298565b95945050505050565b82604051602001613a9c919061926c565b50919050565b60606000615bcd8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150615c2f90615846565b15615c3d576139c981616c94565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615c9c90615917565b600103615d0657604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526139c9906158b7906159e4565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615d6590615846565b15615b8157604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290615dcd9083906173bc565b9050600181511115615e09578060028251615de891906184d7565b81518110615df857615df8618c1c565b602002602001015192505050919050565b5082604051602001613a9c919061926c565b805182516000911115615e30575060006138d0565b81518351602085015160009291615e469161934a565b615e5091906184d7565b905082602001518103615e675760019150506138d0565b82516020840151819020912014905092915050565b60606000615e8983617461565b600101905060008167ffffffffffffffff811115615ea957615ea9617fa8565b6040519080825280601f01601f191660200182016040528015615ed3576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084615edd57509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e5345440000000000000000000000000000000000000000000081840190815285518087018752838152840192909252845180860190955251845290830152606091615fa8905b8290617543565b15615fe857505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261604790615fa1565b1561608757505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d49540000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526160e690615fa1565b1561612657505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261618590615fa1565b806161ea5750604080518082018252601081527f47504c2d322e302d6f722d6c6174657200000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526161ea90615fa1565b1561622a57505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261628990615fa1565b806162ee5750604080518082018252601081527f47504c2d332e302d6f722d6c6174657200000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526162ee90615fa1565b1561632e57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261638d90615fa1565b806163f25750604080518082018252601181527f4c47504c2d322e312d6f722d6c61746572000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526163f290615fa1565b1561643257505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261649190615fa1565b806164f65750604080518082018252601181527f4c47504c2d332e302d6f722d6c61746572000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526164f690615fa1565b1561653657505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261659590615fa1565b156165d557505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261663490615fa1565b1561667457505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e3000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526166d390615fa1565b1561671357505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261677290615fa1565b156167b257505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261681190615fa1565b1561685157505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526168b090615fa1565b806169155750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261691590615fa1565b1561695557505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e31000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526169b490615fa1565b156169f457505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151613a9c929060200161935d565b60608060005b8451811015616a965781858281518110616a2d57616a2d618c1c565b6020026020010151604051602001616a46929190618795565b604051602081830303815290604052915060018551616a6591906184d7565b8114616a8e5781604051602001616a7c91906194c6565b60405160208183030381529060405291505b600101616a11565b5060408051600380825260808201909252600091816020015b6060815260200190600190039081616aaf5790505090508381600081518110616ada57616ada618c1c565b60200260200101819052506040518060400160405280600281526020017f2d6300000000000000000000000000000000000000000000000000000000000081525081600181518110616b2e57616b2e618c1c565b60200260200101819052508181600281518110616b4d57616b4d618c1c565b6020908102919091010152949350505050565b6020808301518351835192840151600093616b7e9291849190617557565b14159392505050565b60408051808201909152600080825260208201526000616bb98460000151856020015185600001518660200151617668565b9050836020015181616bcb91906184d7565b84518590616bda9083906184d7565b90525060208401525090919050565b6040805180820190915260008082526020820152815183511015616c0e5750816138d0565b6020808301519084015160019114616c355750815160208481015190840151829020919020145b8015616c6657825184518590616c4c9083906184d7565b9052508251602085018051616c6290839061934a565b9052505b509192915050565b6040805180820190915260008082526020820152616c8d838383617788565b5092915050565b60606000826000015167ffffffffffffffff811115616cb557616cb5617fa8565b6040519080825280601f01601f191660200182016040528015616cdf576020820181803683370190505b5090506000602082019050616c8d8185602001518660000151617833565b60606000616d09613cc6565b6040805160ff808252612000820190925291925060009190816020015b6060815260200190600190039081616d2657905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280616d8190618d89565b935060ff1681518110616d9657616d96618c1c565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e3300000000000000000000000000000000000000000000000000815250604051602001616de79190619507565b604051602081830303815290604052828280616e0290618d89565b935060ff1681518110616e1757616e17618c1c565b60200260200101819052506040518060400160405280600881526020017f76616c6964617465000000000000000000000000000000000000000000000000815250828280616e6490618d89565b935060ff1681518110616e7957616e79618c1c565b602002602001018190525082604051602001616e959190618cb7565b604051602081830303815290604052828280616eb090618d89565b935060ff1681518110616ec557616ec5618c1c565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e747261637400000000000000000000000000000000000000000000815250828280616f1290618d89565b935060ff1681518110616f2757616f27618c1c565b6020026020010181905250616f3c87846178ad565b8282616f4781618d89565b935060ff1681518110616f5c57616f5c618c1c565b6020908102919091010152855151156170085760408051808201909152600b81527f2d2d7265666572656e636500000000000000000000000000000000000000000060208201528282616fae81618d89565b935060ff1681518110616fc357616fc3618c1c565b6020026020010181905250616fdc8660000151846178ad565b8282616fe781618d89565b935060ff1681518110616ffc57616ffc618c1c565b60200260200101819052505b8560800151156170765760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261705181618d89565b935060ff168151811061706657617066618c1c565b60200260200101819052506170dc565b84156170dc5760408051808201909152601281527f2d2d726571756972655265666572656e63650000000000000000000000000000602082015282826170bb81618d89565b935060ff16815181106170d0576170d0618c1c565b60200260200101819052505b604086015151156171785760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261712681618d89565b935060ff168151811061713b5761713b618c1c565b6020026020010181905250856040015182828061715790618d89565b935060ff168151811061716c5761716c618c1c565b60200260200101819052505b8560600151156171e25760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826171c181618d89565b935060ff16815181106171d6576171d6618c1c565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561720057617200617fa8565b60405190808252806020026020018201604052801561723357816020015b606081526020019060019003908161721e5790505b50905060005b8260ff168160ff16101561728c57838160ff168151811061725c5761725c618c1c565b6020026020010151828260ff168151811061727957617279618c1c565b6020908102919091010152600101617239565b50979650505050505050565b60408051808201909152600080825260208201528151835110156172bd5750816138d0565b815183516020850151600092916172d39161934a565b6172dd91906184d7565b602084015190915060019082146172fe575082516020840151819020908220145b8015617319578351855186906173159083906184d7565b9052505b50929392505050565b60008082600001516173468560000151866020015186600001518760200151617668565b617350919061934a565b90505b83516020850151617364919061934a565b8111616c8d57816173748161954c565b92505082600001516173ab85602001518361738f91906184d7565b865161739b91906184d7565b8386600001518760200151617668565b6173b5919061934a565b9050617353565b606060006173ca8484617322565b6173d590600161934a565b67ffffffffffffffff8111156173ed576173ed617fa8565b60405190808252806020026020018201604052801561742057816020015b606081526020019060019003908161740b5790505b50905060005b81518110156156955761743c6158b78686616c6e565b82828151811061744e5761744e618c1c565b6020908102919091010152600101617426565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106174aa577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106174d6576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106174f457662386f26fc10000830492506010015b6305f5e100831061750c576305f5e100830492506008015b612710831061752057612710830492506004015b60648310617532576064830492506002015b600a83106138d05760010192915050565b600061754f83836178ed565b159392505050565b60008085841161765e576020841161760a57600084156175a257600161757e8660206184d7565b617589906008619566565b617594906002619664565b61759e91906184d7565b1990505b83518116856175b1898961934a565b6175bb91906184d7565b805190935082165b8181146175f5578784116175dd57879450505050506152c6565b836175e781619670565b9450508284511690506175c3565b6175ff878561934a565b9450505050506152c6565b83832061761785886184d7565b617621908761934a565b91505b85821061765c578482208082036176495761763f868461934a565b93505050506152c6565b6176546001846184d7565b925050617624565b505b5092949350505050565b60008381868511617773576020851161772257600085156176b45760016176908760206184d7565b61769b906008619566565b6176a6906002619664565b6176b091906184d7565b1990505b845181166000876176c58b8b61934a565b6176cf91906184d7565b855190915083165b828114617714578186106176fc576176ef8b8b61934a565b96505050505050506152c6565b856177068161954c565b9650508386511690506176d7565b8596505050505050506152c6565b508383206000905b61773486896184d7565b82116177715785832080820361775057839450505050506152c6565b61775b60018561934a565b93505081806177699061954c565b92505061772a565b505b61777d878761934a565b979650505050505050565b604080518082019091526000808252602082015260006177ba8560000151866020015186600001518760200151617668565b6020808701805191860191909152519091506177d690826184d7565b8352845160208601516177e9919061934a565b81036177f8576000855261782a565b83518351617806919061934a565b855186906178159083906184d7565b9052508351617824908261934a565b60208601525b50909392505050565b6020811061786b578151835261784a60208461934a565b925061785760208361934a565b91506178646020826184d7565b9050617833565b600019811561789a5760016178818360206184d7565b61788d90610100619664565b61789791906184d7565b90505b9151835183169219169190911790915250565b606060006178bb8484613d99565b80516020808301516040519394506178d593909101619687565b60405160208183030381529060405291505092915050565b8151815160009190811115617900575081515b6020808501519084015160005b838110156179b957825182518082146179895760001960208710156179685760018461793a8960206184d7565b617944919061934a565b61794f906008619566565b61795a906002619664565b61796491906184d7565b1990505b81811683821681810391146179865797506138d09650505050505050565b50505b61799460208661934a565b94506179a160208561934a565b935050506020816179b2919061934a565b905061790d565b508451865161444791906196df565b610c9f806200970083390190565b611eb9806200a39f83390190565b611783806200c25883390190565b610de0806200d9db83390190565b610a12806200e7bb83390190565b610b3f806200f1cd83390190565b611e01806200fd0c83390190565b6040518060e00160405280606081526020016060815260200160608152602001600015158152602001600015158152602001600015158152602001617a6d617a72565b905290565b60405180610100016040528060001515815260200160001515815260200160608152602001600080191681526020016060815260200160608152602001600015158152602001617a6d6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b81811015617b245783516001600160a01b0316835260209384019390920191600101617afd565b509095945050505050565b60005b83811015617b4a578181015183820152602001617b32565b50506000910152565b60008151808452617b6b816020860160208601617b2f565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015617c7b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b81811015617c61577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8503018352617c4b848651617b53565b6020958601959094509290920191600101617c11565b509197505050602094850194929092019150600101617ba7565b50929695505050505050565b600081518084526020840193506020830160005b82811015617cdb5781517fffffffff0000000000000000000000000000000000000000000000000000000016865260209586019590910190600101617c9b565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015617c7b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160408752617d516040880182617b53565b9050602082015191508681036020880152617d6c8183617c87565b965050506020938401939190910190600101617d0d565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015617c7b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452617de5858351617b53565b94506020938401939190910190600101617dab565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015617c7b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b0381511686526020810151905060406020870152617e7b6040870182617c87565b9550506020938401939190910190600101617e22565b600060208284031215617ea357600080fd5b815180151581146139c957600080fd5b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff8816604083015286606083015260038610617f6d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8560808301528460a0830152617f8e60c08301856001600160a01b03169052565b6001600160a01b03831660e0830152979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c90821680617feb57607f821691505b602082108103615b92577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f82111561806e57806000526020600020601f840160051c8101602085101561804b5750805b601f840160051c820191505b8181101561806b5760008155600101618057565b50505b505050565b815167ffffffffffffffff81111561808d5761808d617fa8565b6180a18161809b8454617fd7565b84618024565b6020601f8211600181146180d557600083156180bd5750848201515b600019600385901b1c1916600184901b17845561806b565b600084815260208120601f198516915b8281101561810557878501518255602094850194600190920191016180e5565b50848210156181235786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60006020828403121561814457600080fd5b5051919050565b60608152600061815e6060830186617b53565b602083019490945250901515604090910152919050565b600081546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160806060850152600081546181bf81617fd7565b80608088015260018216600081146181de57600181146182185761824c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660a089015260a082151560051b890101935061824c565b84600052602060002060005b838110156182435781548a820160a00152600190910190602001618224565b890160a0019450505b50919695505050505050565b60c08152600061826b60c0830189617b53565b62ffffff881660208401526001600160a01b038716604084015282810360608401526182978187617b53565b905060ff8516608084015282810360a08401526182b48185618175565b9998505050505050505050565b6001600160a01b0384168152826020820152606060408201526000615b786060830184617b53565b60c0815260006182fc60c0830189617b53565b8760208401526001600160a01b038716604084015282810360608401526183238187617b53565b6080840195909552505090151560a090910152949350505050565b6001600160a01b038616815284602082015260a06040820152600061836660a0830186617b53565b6060830194909452509015156080909101529392505050565b8281526040602082015260006152c66040830184617b53565b6001600160a01b03831681526040602082015260006152c66040830184617b53565b6040516060810167ffffffffffffffff811182821017156183dd576183dd617fa8565b60405290565b60008067ffffffffffffffff8411156183fe576183fe617fa8565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561842d5761842d617fa8565b60405283815290508082840185101561844557600080fd5b615695846020830185617b2f565b600082601f83011261846457600080fd5b6139c9838351602085016183e3565b60006020828403121561848557600080fd5b815167ffffffffffffffff81111561849c57600080fd5b6138cc84828501618453565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156138d0576138d06184a8565b610100815260006184ff61010083018b617b53565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a08401526185328187617b53565b90508460c084015282810360e0840152613bbc8185618175565b60c08152600061855f60c0830189617b53565b8760208401526001600160a01b038716604084015282810360608401526185868187617b53565b905084608084015282810360a08401526182b48185618175565b6080815260006185b36080830187617b53565b82810360208401526185c58187617b53565b9050846040840152828103606084015261777d8185618175565b60a0815260006185f260a0830188617b53565b6001600160a01b038716602084015282810360408401526186138187617b53565b9050846060840152828103608084015261862d8185618175565b98975050505050505050565b60a08152600061864c60a0830188617b53565b6001600160a01b0387166020840152828103604084015261866d8187617b53565b905060ff85166060840152828103608084015261862d8185618175565b60a08152600061869d60a0830188617b53565b6001600160a01b038716602084015282810360408401526186be8187617b53565b606084019590955250509015156080909101529392505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161871081601a850160208801617b2f565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161874d81601c840160208801617b2f565b01601c01949350505050565b6020815260006139c96020830184617b53565b60006020828403121561877e57600080fd5b81516001600160a01b03811681146139c957600080fd5b600083516187a7818460208801617b2f565b8351908301906187bb818360208801617b2f565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e7472616374200000000000008152600083516187fc81601a850160208801617b2f565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a918401918201528351618839816033840160208801617b2f565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f555400000000000000000000000000000000000000000060608201526080602082015260006139c96080830184617b53565b6000602082840312156188c857600080fd5b815167ffffffffffffffff8111156188df57600080fd5b8201601f810184136188f057600080fd5b6138cc848251602084016183e3565b60008551618911818460208a01617b2f565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161894b816001840160208a01617b2f565b7f2f00000000000000000000000000000000000000000000000000000000000000600192909101918201528451618989816002840160208901617b2f565b6001818301019150507f2f00000000000000000000000000000000000000000000000000000000000000600182015283516189cb816002840160208801617b2f565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b604081526000618a166040830184617b53565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e2061727469666163742000815260008251618a8d81601f850160208701617b2f565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b604081526000618afa6040830184617b53565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081526000618b4c6040830184617b53565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b27000000000000000000000000815260008251618bc3816014850160208701617b2f565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b604081526000618c0a6040830185617b53565b82810360208401526139c58185617b53565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f2200000000000000000000000000000000000000000000000000000000000000815260008251618c83816001850160208701617b2f565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b60008251618cc9818460208701617b2f565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e747261637420000000000000000000000000000000000000000000604082015260008251618d7c81604b850160208701617b2f565b91909101604b0192915050565b600060ff821660ff8103618d9f57618d9f6184a8565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c69400000000000000000000000000000000000000000000000602082015260008251618e06816029850160208701617b2f565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f504154480000000000000000000060608201526080602082015260006139c96080830184617b53565b600060208284031215618e6c57600080fd5b815167ffffffffffffffff811115618e8357600080fd5b820160608185031215618e9557600080fd5b618e9d6183ba565b81518060030b8114618eae57600080fd5b8152602082015167ffffffffffffffff811115618eca57600080fd5b618ed686828501618453565b602083015250604082015167ffffffffffffffff811115618ef657600080fd5b618f0286828501618453565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f2200000000000000000000000000000000000000000000000000000000000000602082015260008251618f6e816021850160208701617b2f565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161915a816021850160208801617b2f565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161919781602e840160208801617b2f565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a200000000000000000000000000000000000000000000000602082015260008251618e06816029850160208701617b2f565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161925f816022850160208701617b2f565b9190910160220192915050565b7f436f6e7472616374206e616d65200000000000000000000000000000000000008152600082516192a481600e850160208701617b2f565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b808201808211156138d0576138d06184a8565b7f53504458206c6963656e7365206964656e746966696572200000000000000000815260008351619395816018850160208801617b2f565b7f20696e200000000000000000000000000000000000000000000000000000000060189184019182015283516193d281601c840160208801617b2f565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b600082516194d8818460208701617b2f565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161953f81601c850160208701617b2f565b91909101601c0192915050565b6000600019820361955f5761955f6184a8565b5060010190565b80820281158282048414176138d0576138d06184a8565b6001815b60018411156195b85780850481111561959c5761959c6184a8565b60018416156195aa57908102905b60019390931c928002619581565b935093915050565b6000826195cf575060016138d0565b816195dc575060006138d0565b81600181146195f257600281146195fc57619618565b60019150506138d0565b60ff84111561960d5761960d6184a8565b50506001821b6138d0565b5060208310610133831016604e8410600b841016171561963b575081810a6138d0565b619648600019848461957d565b806000190482111561965c5761965c6184a8565b029392505050565b60006139c983836195c0565b60008161967f5761967f6184a8565b506000190190565b60008351619699818460208801617b2f565b7f3a0000000000000000000000000000000000000000000000000000000000000090830190815283516196d3816001840160208801617b2f565b01600101949350505050565b8181036000831280158383131683831282161715616c8d57616c8d6184a856fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea26469706673582212207e804ca539d49155d2b6bc19268ce22f9f857027c75247d69fb0d56a089c93d464736f6c634300081a003360a060405234801561001057600080fd5b50604051611eb9380380611eb983398101604081905261002f916101fd565b60016000556002805460ff191690556001600160a01b038316158061005b57506001600160a01b038216155b8061006d57506001600160a01b038116155b1561008b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03838116608052600480546001600160a01b0319169184169190911790556100bb60008261014c565b506100e67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8261014c565b506101117f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361014c565b5061012a600080516020611e998339815191528261014c565b50610143600080516020611e998339815191528361014c565b50505050610240565b60008281526001602090815260408083206001600160a01b038516845290915281205460ff166101d75760008381526001602081815260408084206001600160a01b0387168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016101db565b5060005b92915050565b80516001600160a01b03811681146101f857600080fd5b919050565b60008060006060848603121561021257600080fd5b61021b846101e1565b9250610229602085016101e1565b9150610237604085016101e1565b90509250925092565b608051611c22610277600039600081816101ca01528181610597015281816105f901528181610a280152610a8a0152611c226000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c806385f438c1116100e3578063d547741f1161008c578063e609055e11610066578063e609055e146103fc578063e63ab1e91461040f578063eab103df1461043657600080fd5b8063d547741f146103b3578063d936547e146103c6578063d9caed12146103e957600080fd5b80639b19251a116100bd5780639b19251a14610385578063a217fddf14610398578063c709ab6e146103a057600080fd5b806385f438c11461030557806391d148541461032c5780639a5904271461037257600080fd5b806336568abe116101455780635b1125911161011f5780635b112591146102d25780635c975abb146102f25780638456cb59146102fd57600080fd5b806336568abe146102905780633f4ba83a146102a3578063570618e1146102ab57600080fd5b8063248a9ca311610176578063248a9ca314610226578063252f07bf146102585780632f2ff15d1461027d57600080fd5b806301ffc9a71461019d578063116191b6146101c557806321fc65f214610211575b600080fd5b6101b06101ab3660046115ca565b610449565b60405190151581526020015b60405180910390f35b6101ec7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101bc565b61022461021f366004611677565b6104e2565b005b61024a6102343660046116ea565b6000908152600160208190526040909120015490565b6040519081526020016101bc565b6004546101b09074010000000000000000000000000000000000000000900460ff1681565b61022461028b366004611703565b6106e3565b61022461029e366004611703565b61070f565b61022461076d565b61024a7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b6004546101ec9073ffffffffffffffffffffffffffffffffffffffff1681565b60025460ff166101b0565b6102246107a2565b61024a7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b6101b061033a366004611703565b600091825260016020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610224610380366004611733565b6107d4565b610224610393366004611733565b6108a2565b61024a600081565b6102246103ae366004611750565b610973565b6102246103c1366004611703565b610b79565b6101b06103d4366004611733565b60036020526000908152604090205460ff1681565b6102246103f73660046117f3565b610b9f565b61022461040a366004611834565b610ccb565b61024a7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6102246104443660046118d3565b610f2b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104dc57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104ea610f81565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461051481610fc4565b61051c610fce565b73ffffffffffffffffffffffffffffffffffffffff851660009081526003602052604090205460ff1661057b576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105bc73ffffffffffffffffffffffffffffffffffffffff86167f00000000000000000000000000000000000000000000000000000000000000008661100d565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab59906106369088908a90899089908990600401611939565b600060405180830381600087803b15801561065057600080fd5b505af1158015610664573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d58686866040516106c993929190611996565b60405180910390a3506106dc6001600055565b5050505050565b600082815260016020819052604090912001546106ff81610fc4565b610709838361108e565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331461075e576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610768828261113b565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61079781610fc4565b61079f6111dc565b50565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107cc81610fc4565b61079f61123b565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6107fe81610fc4565b73ffffffffffffffffffffffffffffffffffffffff821661084b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216600081815260036020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a6108cc81610fc4565b73ffffffffffffffffffffffffffffffffffffffff8216610919576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216600081815260036020526040808220805460ff19166001179055517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b61097b610f81565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46109a581610fc4565b6109ad610fce565b73ffffffffffffffffffffffffffffffffffffffff861660009081526003602052604090205460ff16610a0c576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a4d73ffffffffffffffffffffffffffffffffffffffff87167f00000000000000000000000000000000000000000000000000000000000000008761100d565b6040517fd0b492c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610ac99089908b908a908a908a908a90600401611a77565b600060405180830381600087803b158015610ae357600080fd5b505af1158015610af7573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f2032883a139c935aa5ecfcba7233f50f723279d7418d69424daa39a5af76d13b87878787604051610b5e9493929190611ae8565b60405180910390a350610b716001600055565b505050505050565b60008281526001602081905260409091200154610b9581610fc4565b610709838361113b565b610ba7610f81565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610bd181610fc4565b610bd9610fce565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604090205460ff16610c38576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c5973ffffffffffffffffffffffffffffffffffffffff8416858461100d565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb84604051610cb891815260200190565b60405180910390a3506107686001600055565b610cd3610f81565b610cdb610fce565b60045474010000000000000000000000000000000000000000900460ff16610d2f576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841660009081526003602052604090205460ff16610d8e576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015610dfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1f9190611b14565b9050610e4373ffffffffffffffffffffffffffffffffffffffff8616333087611278565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa158015610ed7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610efb9190611b14565b610f059190611b2d565b8787604051610f18959493929190611b67565b60405180910390a250610b716001600055565b6000610f3681610fc4565b506004805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b600260005403610fbd576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b61079f81336112be565b60025460ff161561100b576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60405173ffffffffffffffffffffffffffffffffffffffff83811660248301526044820183905261076891859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061134f565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1661113357600083815260016020818152604080842073ffffffffffffffffffffffffffffffffffffffff87168086529252808420805460ff19169093179092559051339286917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45060016104dc565b5060006104dc565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff161561113357600083815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104dc565b6111e46113e5565b6002805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b611243610fce565b6002805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586112113390565b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526107099186918216906323b872dd90608401611047565b600082815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661134b576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044015b60405180910390fd5b5050565b600061137173ffffffffffffffffffffffffffffffffffffffff841683611421565b905080516000141580156113965750808060200190518101906113949190611ba0565b155b15610768576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401611342565b60025460ff1661100b576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606061142f83836000611436565b9392505050565b606081471015611474576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611342565b6000808573ffffffffffffffffffffffffffffffffffffffff16848660405161149d9190611bbd565b60006040518083038185875af1925050503d80600081146114da576040519150601f19603f3d011682016040523d82523d6000602084013e6114df565b606091505b50915091506114ef8683836114f9565b9695505050505050565b60608261150e5761150982611588565b61142f565b8151158015611532575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611581576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401611342565b508061142f565b8051156115985780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156115dc57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461142f57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461079f57600080fd5b60008083601f84011261164057600080fd5b50813567ffffffffffffffff81111561165857600080fd5b60208301915083602082850101111561167057600080fd5b9250929050565b60008060008060006080868803121561168f57600080fd5b853561169a8161160c565b945060208601356116aa8161160c565b935060408601359250606086013567ffffffffffffffff8111156116cd57600080fd5b6116d98882890161162e565b969995985093965092949392505050565b6000602082840312156116fc57600080fd5b5035919050565b6000806040838503121561171657600080fd5b8235915060208301356117288161160c565b809150509250929050565b60006020828403121561174557600080fd5b813561142f8161160c565b60008060008060008060a0878903121561176957600080fd5b86356117748161160c565b955060208701356117848161160c565b945060408701359350606087013567ffffffffffffffff8111156117a757600080fd5b6117b389828a0161162e565b909450925050608087013567ffffffffffffffff8111156117d357600080fd5b87016060818a0312156117e557600080fd5b809150509295509295509295565b60008060006060848603121561180857600080fd5b83356118138161160c565b925060208401356118238161160c565b929592945050506040919091013590565b6000806000806000806080878903121561184d57600080fd5b863567ffffffffffffffff81111561186457600080fd5b61187089828a0161162e565b90975095505060208701356118848161160c565b935060408701359250606087013567ffffffffffffffff8111156118a757600080fd5b6118b389828a0161162e565b979a9699509497509295939492505050565b801515811461079f57600080fd5b6000602082840312156118e557600080fd5b813561142f816118c5565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015283604082015260806060820152600061198b6080830184866118f0565b979650505050505050565b8381526040602082015260006119b06040830184866118f0565b95945050505050565b600081356119c68161160c565b73ffffffffffffffffffffffffffffffffffffffff168352602082013567ffffffffffffffff81168082146119fa57600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112611a3657600080fd5b820160208101903567ffffffffffffffff811115611a5357600080fd5b803603821315611a6257600080fd5b606060408601526119b06060860182846118f0565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a060608201526000611ac960a0830185876118f0565b8281036080840152611adb81856119b9565b9998505050505050505050565b848152606060208201526000611b026060830185876118f0565b828103604084015261198b81856119b9565b600060208284031215611b2657600080fd5b5051919050565b818103818111156104dc577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b606081526000611b7b6060830187896118f0565b8560208401528281036040840152611b948185876118f0565b98975050505050505050565b600060208284031215611bb257600080fd5b815161142f816118c5565b6000825160005b81811015611bde5760208186018101518583015201611bc4565b50600092019182525091905056fea26469706673582212201d2df9f6a624b84cf0d820aaf8f90e4c1ceeccfc5455d7e360a6c47b3859439864736f6c634300081a00338619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a60c060405260001960035534801561001657600080fd5b5060405161178338038061178383398101604081905261003591610220565b60016000819055805460ff19169055838383836001600160a01b038416158061006557506001600160a01b038316155b8061007757506001600160a01b038216155b8061008957506001600160a01b038116155b156100a75760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03808516608052831660a0526100c5600082610154565b506100f07f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610154565b5061011b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610154565b506101467f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610154565b505050505050505050610274565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166101fa5760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101b23390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101fe565b5060005b92915050565b80516001600160a01b038116811461021b57600080fd5b919050565b6000806000806080858703121561023657600080fd5b61023f85610204565b935061024d60208601610204565b925061025b60408601610204565b915061026960608601610204565b905092959194509250565b60805160a0516114ab6102d86000396000818161021d01528181610531015281816108290152818161099a01528181610aef0152610c110152600081816101d1015281816104a1015281816105040152818161079901526107fc01526114ab6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80635e3e9fef116100d857806391d148541161008c578063d547741f11610066578063d547741f14610386578063d5abeb0114610399578063e63ab1e9146103a257600080fd5b806391d1485414610311578063a217fddf14610357578063a783c7891461035f57600080fd5b8063743e0c9b116100bd578063743e0c9b146102cf5780638456cb59146102e257806385f438c1146102ea57600080fd5b80635e3e9fef146102a95780636f8b44b0146102bc57600080fd5b8063248a9ca31161012f57806336568abe1161011457806336568abe146102835780633f4ba83a146102965780635c975abb1461029e57600080fd5b8063248a9ca31461023f5780632f2ff15d1461027057600080fd5b8063106e629011610160578063106e6290146101b9578063116191b6146101cc57806321e093b11461021857600080fd5b806301ffc9a71461017c578063057e0f25146101a4575b600080fd5b61018f61018a366004610fd3565b6103c9565b60405190151581526020015b60405180910390f35b6101b76101b236600461108e565b610462565b005b6101b76101c7366004611126565b6105fc565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f37f000000000000000000000000000000000000000000000000000000000000000081565b61026261024d366004611159565b60009081526002602052604090206001015490565b60405190815260200161019b565b6101b761027e366004611172565b6106a1565b6101b7610291366004611172565b6106cc565b6101b7610725565b60015460ff1661018f565b6101b76102b736600461119e565b61075a565b6101b76102ca366004611159565b6108ef565b6101b76102dd366004611159565b61095d565b6101b7610a07565b6102627f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b61018f61031f366004611172565b600091825260026020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610262600081565b6102627f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101b7610394366004611172565b610a39565b61026260035481565b6102627f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061045c57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b61046a610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461049481610aa1565b61049c610aab565b6104c77f00000000000000000000000000000000000000000000000000000000000000008785610aea565b6040517fd0b492c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d0b492c390610563907f0000000000000000000000000000000000000000000000000000000000000000908b908b908b908b908a9060040161130e565b600060405180830381600087803b15801561057d57600080fd5b505af1158015610591573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff167f52d8cccccf212da1f2b87140143958eb3bbf8a92e3833c50a8bf8a719a0da44c878787866040516105e1949392919061137f565b60405180910390a2506105f46001600055565b505050505050565b610604610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461062e81610aa1565b610636610aab565b610641848484610aea565b8373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d58460405161068991815260200190565b60405180910390a25061069c6001600055565b505050565b6000828152600260205260409020600101546106bc81610aa1565b6106c68383610c72565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331461071b576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61069c8282610d72565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61074f81610aa1565b610757610e31565b50565b610762610a5e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461078c81610aa1565b610794610aab565b6107bf7f00000000000000000000000000000000000000000000000000000000000000008684610aea565b6040517f5131ab5900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690635131ab5990610859907f0000000000000000000000000000000000000000000000000000000000000000908a908a908a908a906004016113b6565b600060405180830381600087803b15801561087357600080fd5b505af1158015610887573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d8686866040516108d593929190611408565b60405180910390a2506108e86001600055565b5050505050565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb61091981610aa1565b610921610aab565b60038290556040518281527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c9060200160405180910390a15050565b610965610aab565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b1580156109f357600080fd5b505af11580156108e8573d6000803e3d6000fd5b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a3181610aa1565b610757610eae565b600082815260026020526040902060010154610a5481610aa1565b6106c68383610d72565b600260005403610a9a576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6107578133610f07565b60015460ff1615610ae8576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6003547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7c9190611422565b610b86908461143b565b1115610bbe576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201849052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b158015610c5557600080fd5b505af1158015610c69573d6000803e3d6000fd5b50505050505050565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610d6a57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610d083390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161045c565b50600061045c565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610d6a57600083815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161045c565b610e39610f97565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b610eb6610aab565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610e84565b600082815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610f93576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440160405180910390fd5b5050565b60015460ff16610ae8576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215610fe557600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461101557600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461104057600080fd5b919050565b60008083601f84011261105757600080fd5b50813567ffffffffffffffff81111561106f57600080fd5b60208301915083602082850101111561108757600080fd5b9250929050565b60008060008060008060a087890312156110a757600080fd5b6110b08761101c565b955060208701359450604087013567ffffffffffffffff8111156110d357600080fd5b6110df89828a01611045565b90955093505060608701359150608087013567ffffffffffffffff81111561110657600080fd5b87016060818a03121561111857600080fd5b809150509295509295509295565b60008060006060848603121561113b57600080fd5b6111448461101c565b95602085013595506040909401359392505050565b60006020828403121561116b57600080fd5b5035919050565b6000806040838503121561118557600080fd5b823591506111956020840161101c565b90509250929050565b6000806000806000608086880312156111b657600080fd5b6111bf8661101c565b945060208601359350604086013567ffffffffffffffff8111156111e257600080fd5b6111ee88828901611045565b96999598509660600135949350505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff6112678261101c565b1682526000602082013567ffffffffffffffff811680821461128857600080fd5b6020850152506040820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126112c457600080fd5b820160208101903567ffffffffffffffff8111156112e157600080fd5b8036038213156112f057600080fd5b60606040860152611305606086018284611200565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015284604082015260a06060820152600061136060a083018587611200565b82810360808401526113728185611249565b9998505050505050505050565b848152606060208201526000611399606083018587611200565b82810360408401526113ab8185611249565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff851660208201528360408201526080606082015260006113ab608083018486611200565b838152604060208201526000611305604083018486611200565b60006020828403121561143457600080fd5b5051919050565b8082018082111561045c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea264697066735822122088ba4caf813dddb87f7b16ee6c2b7d47544e7ce3683f60d6984a2dd9c9279b0d64736f6c634300081a00336080604052348015600f57600080fd5b506001600055610dbc806100246000396000f3fe6080604052600436106100635760003560e01c8063c513169111610040578063c5131691146100c1578063e04d4f97146100e1578063f05b6abf146100f457005b8063357fc5a21461006c578063660b9de01461008c5780636ed70169146100ac57005b3661006a57005b005b34801561007857600080fd5b5061006a6100873660046106bd565b610114565b34801561009857600080fd5b5061006a6100a73660046106f9565b6101aa565b3480156100b857600080fd5b5061006a6101e6565b3480156100cd57600080fd5b5061006a6100dc3660046106bd565b61021b565b61006a6100ef366004610859565b6102f6565b34801561010057600080fd5b5061006a61010f366004610945565b61033a565b61011c61036f565b61013e73ffffffffffffffffffffffffffffffffffffffff83163383866103b2565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101a56001600055565b505050565b7f024f3ba167dca37cfa8409bf4c2f05a4d9c2d9b39a94d7a6c970352fabaa320e33826040516101db929190610a78565b60405180910390a150565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61022361036f565b6000610230600285610b57565b90508060000361026c576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61028e73ffffffffffffffffffffffffffffffffffffffff84163384846103b2565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101a56001600055565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa333485858560405161032d959493929190610c00565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b1463384848460405161032d9493929190610c8a565b6002600054036103ab576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261044790859061044d565b50505050565b600061046f73ffffffffffffffffffffffffffffffffffffffff8416836104e8565b905080516000141580156104945750808060200190518101906104929190610d4d565b155b156101a5576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b60606104f6838360006104fd565b9392505050565b60608147101561053b576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016104df565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105649190610d6a565b60006040518083038185875af1925050503d80600081146105a1576040519150601f19603f3d011682016040523d82523d6000602084013e6105a6565b606091505b50915091506105b68683836105c0565b9695505050505050565b6060826105d5576105d08261064f565b6104f6565b81511580156105f9575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610648576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016104df565b50806104f6565b80511561065f5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff811681146106b857600080fd5b919050565b6000806000606084860312156106d257600080fd5b833592506106e260208501610694565b91506106f060408501610694565b90509250925092565b60006020828403121561070b57600080fd5b813567ffffffffffffffff81111561072257600080fd5b8201606081850312156104f657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156107aa576107aa610734565b604052919050565b600082601f8301126107c357600080fd5b813567ffffffffffffffff8111156107dd576107dd610734565b61080e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610763565b81815284602083860101111561082357600080fd5b816020850160208301376000918101602001919091529392505050565b801515811461069157600080fd5b80356106b881610840565b60008060006060848603121561086e57600080fd5b833567ffffffffffffffff81111561088557600080fd5b610891868287016107b2565b9350506020840135915060408401356108a981610840565b809150509250925092565b600067ffffffffffffffff8211156108ce576108ce610734565b5060051b60200190565b600082601f8301126108e957600080fd5b81356108fc6108f7826108b4565b610763565b8082825260208201915060208360051b86010192508583111561091e57600080fd5b602085015b8381101561093b578035835260209283019201610923565b5095945050505050565b60008060006060848603121561095a57600080fd5b833567ffffffffffffffff81111561097157600080fd5b8401601f8101861361098257600080fd5b80356109906108f7826108b4565b8082825260208201915060208360051b8501019250888311156109b257600080fd5b602084015b838110156109f457803567ffffffffffffffff8111156109d657600080fd5b6109e58b6020838901016107b2565b845250602092830192016109b7565b509550505050602084013567ffffffffffffffff811115610a1457600080fd5b610a20868287016108d8565b9250506106f06040850161084e565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610ab683610694565b1660408201526000602083013567ffffffffffffffff8116808214610ada57600080fd5b6060840152506040830135368490037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112610b1657600080fd5b830160208101903567ffffffffffffffff811115610b3357600080fd5b803603821315610b4257600080fd5b606060808501526105b660a085018284610a2f565b600082610b8d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60005b83811015610bad578181015183820152602001610b95565b50506000910152565b60008151808452610bce816020860160208601610b92565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610c3560a0830186610bb6565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610c80578151865260209586019590910190600101610c62565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610d1d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610d08858351610bb6565b94506020938401939190910190600101610cce565b505050508281036040840152610d338186610c4e565b915050610d44606083018415159052565b95945050505050565b600060208284031215610d5f57600080fd5b81516104f681610840565b60008251610d7c818460208701610b92565b919091019291505056fea2646970667358221220d72588ea8acbf56155273edae233d1fe939bca804f44cf5204c615539f68e9ae64736f6c634300081a00336080604052348015600f57600080fd5b50604051610a12380380610a12833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b6109858061008d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b61461005b5780637a34d8bb146100a4575b600080fd5b6100596100543660046105e7565b6100b7565b005b60005461007b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100596100b2366004610687565b61030c565b60008383836040516024016100ce93929190610783565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905260005490915073ffffffffffffffffffffffffffffffffffffffff8087169163095ea7b391166101758960026107ad565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af11580156101e5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020991906107ed565b61023f576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160808101825261032180825260016020808401829052838501929092528351918201845260008083526060840192909252905492517f90ad3e23000000000000000000000000000000000000000000000000000000008152919273ffffffffffffffffffffffffffffffffffffffff16916390ad3e23916102d0918c918c918c918991899060040161087a565b600060405180830381600087803b1580156102ea57600080fd5b505af11580156102fe573d6000803e3d6000fd5b505050505050505050505050565b600083838360405160240161032393929190610783565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905281516080810183526103218082526001828401819052828501919091528351928301845260008084526060830193909352915492517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff93841660048201526024810183905293945092909188169063095ea7b3906044016020604051808303816000875af1158015610447573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061046b91906107ed565b506000546040517fdc9ca2e700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063dc9ca2e7906102d0908b908b908890879089906004016108e8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261050a57600080fd5b81356020830160008067ffffffffffffffff84111561052b5761052b6104ca565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715610578576105786104ca565b60405283815290508082840187101561059057600080fd5b838360208301376000602085830101528094505050505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146105d157600080fd5b919050565b80151581146105e457600080fd5b50565b60008060008060008060c0878903121561060057600080fd5b863567ffffffffffffffff81111561061757600080fd5b61062389828a016104f9565b96505060208701359450610639604088016105ad565b9350606087013567ffffffffffffffff81111561065557600080fd5b61066189828a016104f9565b9350506080870135915060a0870135610679816105d6565b809150509295509295509295565b600080600080600060a0868803121561069f57600080fd5b853567ffffffffffffffff8111156106b657600080fd5b6106c2888289016104f9565b9550506106d1602087016105ad565b9350604086013567ffffffffffffffff8111156106ed57600080fd5b6106f9888289016104f9565b935050606086013591506080860135610711816105d6565b809150509295509295909350565b6000815180845260005b8181101561074557602081850181015186830182015201610729565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b606081526000610796606083018661071f565b602083019490945250901515604090910152919050565b808201808211156107e7577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b6000602082840312156107ff57600080fd5b815161080a816105d6565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff60408201511660408301526000606082015160806060850152610872608085018261071f565b949350505050565b60c08152600061088d60c083018961071f565b87602084015273ffffffffffffffffffffffffffffffffffffffff8716604084015282810360608401526108c1818761071f565b905084608084015282810360a08401526108db8185610811565b9998505050505050505050565b60a0815260006108fb60a083018861071f565b73ffffffffffffffffffffffffffffffffffffffff871660208401528281036040840152610929818761071f565b905084606084015282810360808401526109438185610811565b9897505050505050505056fea26469706673582212207447884a874c8537c3ef840e4ddffbc65b7e07105578a7028d806e21fe26e52164736f6c634300081a0033608060405234801561001057600080fd5b50604051610b3f380380610b3f83398101604081905261002f916100b9565b600380546001600160a01b038086166001600160a01b0319928316179092556004805485841690831617905560058054928416929091169190911790556040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac590600090a15050506100fc565b80516001600160a01b03811681146100b457600080fd5b919050565b6000806000606084860312156100ce57600080fd5b6100d78461009d565b92506100e56020850161009d565b91506100f36040850161009d565b90509250925092565b610a348061010b6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806397770dff11610081578063d7fd7afb1161005b578063d7fd7afb146101f2578063d936a01214610220578063ee2815ba1461024057600080fd5b806397770dff146101b9578063a7cb0507146101cc578063c63585cc146101df57600080fd5b8063513a9c05116100b2578063513a9c0514610143578063569541b914610179578063842da36d1461019957600080fd5b80630be15547146100ce5780633c669d551461012e575b600080fd5b6101046100dc36600461071e565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61014161013c366004610760565b610253565b005b61010461015136600461071e565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101049073ffffffffffffffffffffffffffffffffffffffff1681565b6005546101049073ffffffffffffffffffffffffffffffffffffffff1681565b6101416101c73660046107fd565b6103a0565b6101416101da36600461081f565b610419565b6101046101ed366004610841565b610467565b61021261020036600461071e565b60006020819052908152604090205481565b604051908152602001610125565b6004546101049073ffffffffffffffffffffffffffffffffffffffff1681565b61014161024e366004610884565b61059c565b604080516080810182526000606082019081528152336020820152468183015290517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820186905286169063a9059cbb906044016020604051808303816000875af11580156102e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030b91906108b0565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87169063de43156e90610366908490899089908990899060040161091b565b600060405180830381600087803b15801561038057600080fd5b505af1158015610394573d6000803e3d6000fd5b50505050505050505050565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e9060200160405180910390a150565b6000828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b60008060006104768585610620565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b166034820152919350915086906048016040516020818303038152906040528051906020012060405160200161055c9291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b60008281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d910161045b565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610688576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16106106c25782846106c5565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610717576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b60006020828403121561073057600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461075b57600080fd5b919050565b60008060008060006080868803121561077857600080fd5b61078186610737565b945061078f60208701610737565b935060408601359250606086013567ffffffffffffffff8111156107b257600080fd5b8601601f810188136107c357600080fd5b803567ffffffffffffffff8111156107da57600080fd5b8860208284010111156107ec57600080fd5b959894975092955050506020019190565b60006020828403121561080f57600080fd5b61081882610737565b9392505050565b6000806040838503121561083257600080fd5b50508035926020909101359150565b60008060006060848603121561085657600080fd5b61085f84610737565b925061086d60208501610737565b915061087b60408501610737565b90509250925092565b6000806040838503121561089757600080fd5b823591506108a760208401610737565b90509250929050565b6000602082840312156108c257600080fd5b8151801515811461081857600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60808152600086516060608084015280518060e085015260005b81811015610953576020818401810151610100878401015201610935565b5060008482016101000152602089015173ffffffffffffffffffffffffffffffffffffffff811660a0860152601f9091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168401915050604088015160c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401526101008382030160608401526109f2610100820185876108d2565b9897505050505050505056fea2646970667358221220fecb52d63ec1b967fad2855018329377964bac04ed54a998543272d32f4338fd64736f6c634300081a003360e060405234801561001057600080fd5b50604051611e01380380611e0183398101604081905261002f9161020e565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006357604051632b2add3d60e01b815260040160405180910390fd5b6001600160a01b038216158061008057506001600160a01b038116155b1561009e5760405163d92e233d60e01b815260040160405180910390fd5b60066100aa8982610360565b5060076100b78882610360565b5060ff861660c05260808590528360028111156100d6576100d661041e565b60a08160028111156100ea576100ea61041e565b905250600192909255600080546001600160a01b039283166001600160a01b03199182161790915560088054929093169116179055506104349350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261015157600080fd5b81516001600160401b0381111561016a5761016a61012a565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101985761019861012a565b6040528181528382016020018510156101b057600080fd5b60005b828110156101cf576020818601810151838301820152016101b3565b506000918101602001919091529392505050565b8051600381106101f257600080fd5b919050565b80516001600160a01b03811681146101f257600080fd5b600080600080600080600080610100898b03121561022b57600080fd5b88516001600160401b0381111561024157600080fd5b61024d8b828c01610140565b60208b015190995090506001600160401b0381111561026b57600080fd5b6102778b828c01610140565b975050604089015160ff8116811461028e57600080fd5b60608a015190965094506102a460808a016101e3565b60a08a015190945092506102ba60c08a016101f7565b91506102c860e08a016101f7565b90509295985092959890939650565b600181811c908216806102eb57607f821691505b60208210810361030b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561035b57806000526020600020601f840160051c810160208510156103385750805b601f840160051c820191505b818110156103585760008155600101610344565b50505b505050565b81516001600160401b038111156103795761037961012a565b61038d8161038784546102d7565b84610311565b6020601f8211600181146103c157600083156103a95750848201515b600019600385901b1c1916600184901b178455610358565b600084815260208120601f198516915b828110156103f157878501518255602094850194600190920191016103d1565b508482101561040f5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a05160c05161198261047f6000396000610222015260006103450152600081816102f601528181610af501528181610bfb01528181610e170152610f1d01526119826000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c80638b851b95116100ee578063ccc7759911610097578063eddeb12311610071578063eddeb1231461043c578063f2441b321461044f578063f687d12a1461046f578063fc5fecd51461048257600080fd5b8063ccc77599146103af578063d9eeebed146103c2578063dd62ed3e146103f657600080fd5b8063a9059cbb116100c8578063a9059cbb14610374578063c701262614610387578063c835d7cc1461039a57600080fd5b80638b851b951461031857806395d89b4114610338578063a3413d031461034057600080fd5b80633ce4a5bc116101505780634d8943bb1161012a5780634d8943bb146102b257806370a08231146102bb57806385e1f4d0146102f157600080fd5b80633ce4a5bc1461024c57806342966c681461028c57806347e7ef241461029f57600080fd5b806318160ddd1161018157806318160ddd1461020057806323b872dd14610208578063313ce5671461021b57600080fd5b806306fdde03146101a8578063091d2788146101c6578063095ea7b3146101dd575b600080fd5b6101b0610495565b6040516101bd9190611561565b60405180910390f35b6101cf60015481565b6040519081526020016101bd565b6101f06101eb3660046115a0565b610527565b60405190151581526020016101bd565b6005546101cf565b6101f06102163660046115cc565b61053e565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101bd565b61026773735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101bd565b6101f061029a36600461160d565b6105d5565b6101f06102ad3660046115a0565b6105e9565b6101cf60025481565b6101cf6102c9366004611626565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101cf7f000000000000000000000000000000000000000000000000000000000000000081565b6008546102679073ffffffffffffffffffffffffffffffffffffffff1681565b6101b061073d565b6103677f000000000000000000000000000000000000000000000000000000000000000081565b6040516101bd9190611643565b6101f06103823660046115a0565b61074c565b6101f06103953660046116b3565b610759565b6103ad6103a8366004611626565b6108a8565b005b6103ad6103bd366004611626565b6109bc565b6103ca610ac9565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101bd565b6101cf6104043660046117ab565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b6103ad61044a36600461160d565b610ce7565b6000546102679073ffffffffffffffffffffffffffffffffffffffff1681565b6103ad61047d36600461160d565b610d69565b6103ca61049036600461160d565b610deb565b6060600680546104a4906117e4565b80601f01602080910402602001604051908101604052809291908181526020018280546104d0906117e4565b801561051d5780601f106104f25761010080835404028352916020019161051d565b820191906000526020600020905b81548152906001019060200180831161050057829003601f168201915b5050505050905090565b6000610534338484611007565b5060015b92915050565b600061054b848484611110565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054828110156105b6576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ca85336105c58685611866565b611007565b506001949350505050565b60006105e133836112cb565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab14801590610627575060005473ffffffffffffffffffffffffffffffffffffffff163314155b801561064b575060085473ffffffffffffffffffffffffffffffffffffffff163314155b15610682576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61068c838361140d565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261072c918690611879565b60405180910390a250600192915050565b6060600780546104a4906117e4565b6000610534338484611110565b6000806000610766610ac9565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156107f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081c919061189b565b610852576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61085c33856112cb565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d95591610895918991899187916118bd565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146108f5576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610942576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a09576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a56576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a387906020016109b1565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610b5c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8091906118ec565b905073ffffffffffffffffffffffffffffffffffffffff8116610bcf576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610c5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c829190611909565b905080600003610cbe576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610cd19190611922565b610cdb9190611939565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610d34576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f906020016109b1565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610db6576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a906020016109b1565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610e7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea291906118ec565b905073ffffffffffffffffffffffffffffffffffffffff8116610ef1576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610f80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa49190611909565b905080600003610fe0576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254600090610ff08784611922565b610ffa9190611939565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff8316611054576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166110a1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff831661115d576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166111aa576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83166000908152600360205260409020548181101561120a576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112148282611866565b73ffffffffffffffffffffffffffffffffffffffff8086166000908152600360205260408082209390935590851681529081208054849290611257908490611939565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516112bd91815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8216611318576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821660009081526003602052604090205481811015611378576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113828282611866565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812091909155600580548492906113bd908490611866565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001611103565b73ffffffffffffffffffffffffffffffffffffffff821661145a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806005600082825461146c9190611939565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260036020526040812080548392906114a6908490611939565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561152357602081850181015186830182015201611507565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061157460208301846114fd565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461159d57600080fd5b50565b600080604083850312156115b357600080fd5b82356115be8161157b565b946020939093013593505050565b6000806000606084860312156115e157600080fd5b83356115ec8161157b565b925060208401356115fc8161157b565b929592945050506040919091013590565b60006020828403121561161f57600080fd5b5035919050565b60006020828403121561163857600080fd5b81356115748161157b565b602081016003831061167e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080604083850312156116c657600080fd5b823567ffffffffffffffff8111156116dd57600080fd5b8301601f810185136116ee57600080fd5b803567ffffffffffffffff81111561170857611708611684565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561177457611774611684565b60405281815282820160200187101561178c57600080fd5b8160208401602083013760006020928201830152969401359450505050565b600080604083850312156117be57600080fd5b82356117c98161157b565b915060208301356117d98161157b565b809150509250929050565b600181811c908216806117f857607f821691505b602082108103611831577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561053857610538611837565b60408152600061188c60408301856114fd565b90508260208301529392505050565b6000602082840312156118ad57600080fd5b8151801515811461157457600080fd5b6080815260006118d060808301876114fd565b6020830195909552506040810192909252606090910152919050565b6000602082840312156118fe57600080fd5b81516115748161157b565b60006020828403121561191b57600080fd5b5051919050565b808202811582820484141761053857610538611837565b808201808211156105385761053861183756fea2646970667358221220af0e942a7477c448c9068e05c443b6b8aa66c25114141f907c19366098ed4ca864736f6c634300081a0033a2646970667358221220320b110941c4776c4ab165958e7a721faffc85dae3de68a74e2448aeb36f566e64736f6c634300081a0033",
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
