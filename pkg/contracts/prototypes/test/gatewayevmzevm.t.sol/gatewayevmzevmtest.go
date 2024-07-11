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
	ABI: "[{\"inputs\":[],\"name\":\"ApprovalFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerIsNotFungibleModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CustodyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedZetaSent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasFeeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientERC20Amount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientETHAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientZRC20Amount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WZETATransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZRC20BurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZRC20TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Call\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"Call\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ExecutedWithERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"ReceivedERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ReceivedNoParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"strs\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nums\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"ReceivedNonPayable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"ReceivedPayable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"RevertedWithERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasfee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFlatFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"log_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"log_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"log_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"log_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"log_named_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"log_named_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"log_named_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"}],\"name\":\"log_named_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"log_named_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"log_named_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"log_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"logs\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"IS_TEST\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeArtifacts\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"excludedArtifacts_\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"excludedContracts_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"excludedSenders_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetArtifactSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"artifact\",\"type\":\"string\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetArtifacts\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"targetedArtifacts_\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targetedContracts_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetInterfaces\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"artifacts\",\"type\":\"string[]\"}],\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targetedSenders_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testCallReceiverEVMFromZEVM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600c60006101000a81548160ff0219169083151502179055506001601f60006101000a81548160ff02191690831515021790555034801561004657600080fd5b50620142a180620000586000396000f3fe60806040523480156200001157600080fd5b5060043610620001005760003560e01c8063916a17c61162000099578063b5508aa9116200006f578063b5508aa9146200022d578063ba414fa6146200024f578063e20c9f711462000271578063fa7626d414620002935762000100565b8063916a17c614620001dd5780639683c69514620001ff578063b0464fdc146200020b5762000100565b80633e5e3c2311620000db5780633e5e3c2314620001555780633f7286f4146200017757806366d9a9a0146200019957806385226c8114620001bb5762000100565b80630a9254e414620001055780631ed7831c14620001115780632ade38801462000133575b600080fd5b6200010f620002b5565b005b6200011b6200135c565b6040516200012a919062002fcc565b60405180910390f35b6200013d620013ec565b6040516200014c919062003038565b60405180910390f35b6200015f62001586565b6040516200016e919062002fcc565b60405180910390f35b6200018162001616565b60405162000190919062002fcc565b60405180910390f35b620001a3620016a6565b604051620001b2919062003014565b60405180910390f35b620001c56200183d565b604051620001d4919062002ff0565b60405180910390f35b620001e762001920565b604051620001f691906200305c565b60405180910390f35b6200020962001a73565b005b6200021562002112565b6040516200022491906200305c565b60405180910390f35b6200023762002265565b60405162000246919062002ff0565b60405180910390f35b6200025962002348565b60405162000268919062003080565b60405180910390f35b6200027b6200247b565b6040516200028a919062002fcc565b60405180910390f35b6200029d6200250b565b604051620002ac919062003080565b60405180910390f35b30602560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611234602660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550615678602760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550614321602c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604051620003cd906200251e565b620003d890620032a2565b604051809103906000f080158015620003f5573d6000803e3d6000fd5b50602360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405162000444906200251e565b6200044f90620031d3565b604051809103906000f0801580156200046c573d6000803e3d6000fd5b50602160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604051620004bb906200252c565b604051809103906000f080158015620004d8573d6000803e3d6000fd5b50601f60016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516200054a906200253a565b62000556919062002e5d565b604051809103906000f08015801562000573573d6000803e3d6000fd5b50602260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16602160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051620006089062002548565b6200061592919062002e7a565b604051809103906000f08015801562000632573d6000803e3d6000fd5b50602060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663485cc955602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16602160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518363ffffffff1660e01b81526004016200071692919062002e7a565b600060405180830381600087803b1580156200073157600080fd5b505af115801562000746573d6000803e3d6000fd5b50505050601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516200077b906200253a565b62000787919062002e5d565b604051809103906000f080158015620007a4573d6000803e3d6000fd5b50602260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ae7a3a6f602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040162000864919062002e5d565b600060405180830381600087803b1580156200087f57600080fd5b505af115801562000894573d6000803e3d6000fd5b50505050601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166310188aef602060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040162000917919062002e5d565b600060405180830381600087803b1580156200093257600080fd5b505af115801562000947573d6000803e3d6000fd5b50505050602360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f19602560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16620f42406040518363ffffffff1660e01b8152600401620009cf92919062002f45565b600060405180830381600087803b158015620009ea57600080fd5b505af1158015620009ff573d6000803e3d6000fd5b50505050602360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166207a1206040518363ffffffff1660e01b815260040162000a8792919062002f72565b602060405180830381600087803b15801562000aa257600080fd5b505af115801562000ab7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000add919062002648565b5060405162000aec9062002556565b604051809103906000f08015801562000b09573d6000803e3d6000fd5b50602460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405162000b589062002564565b604051809103906000f08015801562000b75573d6000803e3d6000fd5b50602860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405162000be79062002572565b62000bf3919062002e5d565b604051809103906000f08015801562000c10573d6000803e3d6000fd5b50602960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073735b14bb79463307aacbed86daf3322b1e6226ab90507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166306447d56826040518263ffffffff1660e01b815260040162000cc8919062002e5d565b600060405180830381600087803b15801562000ce357600080fd5b505af115801562000cf8573d6000803e3d6000fd5b50505050600080600060405162000d0f9062002580565b62000d1d9392919062002ea7565b604051809103906000f08015801562000d3a573d6000803e3d6000fd5b50602a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060126001600080602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16602860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405162000dd6906200258e565b62000de7969594939291906200320a565b604051809103906000f08015801562000e04573d6000803e3d6000fd5b50602b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ee2815ba6001602b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518363ffffffff1660e01b815260040162000ec792919062003135565b600060405180830381600087803b15801562000ee257600080fd5b505af115801562000ef7573d6000803e3d6000fd5b50505050602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a7cb05076001806040518363ffffffff1660e01b815260040162000f5b92919062003162565b600060405180830381600087803b15801562000f7657600080fd5b505af115801562000f8b573d6000803e3d6000fd5b50505050602b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166347e7ef24602c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16620f42406040518363ffffffff1660e01b81526004016200101392919062002f45565b602060405180830381600087803b1580156200102e57600080fd5b505af115801562001043573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062001069919062002648565b50602b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166347e7ef24602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16620f42406040518363ffffffff1660e01b8152600401620010ee92919062002f45565b602060405180830381600087803b1580156200110957600080fd5b505af11580156200111e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062001144919062002648565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015620011b157600080fd5b505af1158015620011c6573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663ca669fa7602c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b81526004016200124a919062002e5d565b600060405180830381600087803b1580156200126557600080fd5b505af11580156200127a573d6000803e3d6000fd5b50505050602b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095ea7b3602860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16620f42406040518363ffffffff1660e01b81526004016200130292919062002f45565b602060405180830381600087803b1580156200131d57600080fd5b505af115801562001332573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062001358919062002648565b5050565b60606016805480602002602001604051908101604052809291908181526020018280548015620013e257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162001397575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020016000905b828210156200157d57838290600052602060002090600202016040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805480602002602001604051908101604052809291908181526020016000905b8282101562001565578382906000526020600020018054620014d190620036a8565b80601f0160208091040260200160405190810160405280929190818152602001828054620014ff90620036a8565b8015620015505780601f10620015245761010080835404028352916020019162001550565b820191906000526020600020905b8154815290600101906020018083116200153257829003601f168201915b505050505081526020019060010190620014af565b50505050815250508152602001906001019062001410565b50505050905090565b606060188054806020026020016040519081016040528092919081815260200182805480156200160c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311620015c1575b5050505050905090565b606060178054806020026020016040519081016040528092919081815260200182805480156200169c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162001651575b5050505050905090565b6060601b805480602002602001604051908101604052809291908181526020016000905b828210156200183457838290600052602060002090600202016040518060400160405290816000820180546200170090620036a8565b80601f01602080910402602001604051908101604052809291908181526020018280546200172e90620036a8565b80156200177f5780601f1062001753576101008083540402835291602001916200177f565b820191906000526020600020905b8154815290600101906020018083116200176157829003601f168201915b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156200181b57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411620017c75790505b50505050508152505081526020019060010190620016ca565b50505050905090565b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015620019175783829060005260206000200180546200188390620036a8565b80601f0160208091040260200160405190810160405280929190818152602001828054620018b190620036a8565b8015620019025780601f10620018d65761010080835404028352916020019162001902565b820191906000526020600020905b815481529060010190602001808311620018e457829003601f168201915b50505050508152602001906001019062001861565b50505050905090565b6060601d805480602002602001604051908101604052809291908181526020016000905b8282101562001a6a57838290600052602060002090600202016040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820180548060200260200160405190810160405280929190818152602001828054801562001a5157602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411620019fd5790505b5050505050815250508152602001906001019062001944565b50505050905090565b60006040518060400160405280600f81526020017f48656c6c6f2c204861726468617421000000000000000000000000000000000081525090506000602a90506000600190506000670de0b6b3a76400009050600063e04d4f9760e01b85858560405160240162001ae7939291906200318f565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663ca669fa7602c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040162001bc6919062002e5d565b600060405180830381600087803b15801562001be157600080fd5b505af115801562001bf6573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166381bad6f3600180600180602860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b815260040162001c849594939291906200309d565b600060405180830381600087803b15801562001c9f57600080fd5b505af115801562001cb4573d6000803e3d6000fd5b50505050602c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f602460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405160200162001d47919062002e40565b6040516020818303038152906040528360405162001d67929190620030fa565b60405180910390a2602860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630ac7c44c602460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405160200162001de2919062002e40565b604051602081830303815290604052836040518363ffffffff1660e01b815260040162001e11929190620030fa565b600060405180830381600087803b15801562001e2c57600080fd5b505af115801562001e41573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663c88a5e6d601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b815260040162001ec792919062002f9f565b600060405180830381600087803b15801562001ee257600080fd5b505af115801562001ef7573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166381bad6f3600180600180601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b815260040162001f859594939291906200309d565b600060405180830381600087803b15801562001fa057600080fd5b505af115801562001fb5573d6000803e3d6000fd5b50505050602460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f838360405162002025929190620032d9565b60405180910390a2601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631cff79cd83602460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401620020af92919062002f11565b6000604051808303818588803b158015620020c957600080fd5b505af1158015620020de573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f820116820180604052508101906200210a9190620026ac565b505050505050565b6060601c805480602002602001604051908101604052809291908181526020016000905b828210156200225c57838290600052602060002090600202016040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054806020026020016040519081016040528092919081815260200182805480156200224357602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411620021ef5790505b5050505050815250508152602001906001019062002136565b50505050905090565b60606019805480602002602001604051908101604052809291908181526020016000905b828210156200233f578382906000526020600020018054620022ab90620036a8565b80601f0160208091040260200160405190810160405280929190818152602001828054620022d990620036a8565b80156200232a5780601f10620022fe576101008083540402835291602001916200232a565b820191906000526020600020905b8154815290600101906020018083116200230c57829003601f168201915b50505050508152602001906001019062002289565b50505050905090565b6000600860009054906101000a900460ff16156200237857600860009054906101000a900460ff16905062002478565b6000801b7f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663667f9d707f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c7f6661696c656400000000000000000000000000000000000000000000000000006040518363ffffffff1660e01b81526004016200241f92919062002ee4565b60206040518083038186803b1580156200243857600080fd5b505afa1580156200244d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200247391906200267a565b141590505b90565b606060158054806020026020016040519081016040528092919081815260200182805480156200250157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311620024b6575b5050505050905090565b601f60009054906101000a900460ff1681565b611813806200393d83390190565b613871806200515083390190565b61123f80620089c183390190565b6111b98062009c0083390190565b6110aa806200adb983390190565b613f38806200be6383390190565b610bcd806200fd9b83390190565b6110d7806201096883390190565b61282d8062011a3f83390190565b6000620025b3620025ad8462003336565b6200330d565b905082815260208101848484011115620025d257620025d1620037ce565b5b620025df84828562003672565b509392505050565b600081519050620025f88162003908565b92915050565b6000815190506200260f8162003922565b92915050565b600082601f8301126200262d576200262c620037c9565b5b81516200263f8482602086016200259c565b91505092915050565b600060208284031215620026615762002660620037d8565b5b60006200267184828501620025e7565b91505092915050565b600060208284031215620026935762002692620037d8565b5b6000620026a384828501620025fe565b91505092915050565b600060208284031215620026c557620026c4620037d8565b5b600082015167ffffffffffffffff811115620026e657620026e5620037d3565b5b620026f48482850162002615565b91505092915050565b60006200270b838362002789565b60208301905092915050565b600062002725838362002b26565b60208301905092915050565b60006200273f838362002bf9565b905092915050565b600062002755838362002d65565b905092915050565b60006200276b838362002dad565b905092915050565b600062002781838362002dee565b905092915050565b62002794816200351c565b82525050565b620027a5816200351c565b82525050565b6000620027b882620033cc565b620027c4818562003472565b9350620027d1836200336c565b8060005b8381101562002808578151620027ec8882620026fd565b9750620027f98362003424565b925050600181019050620027d5565b5085935050505092915050565b60006200282282620033d7565b6200282e818562003483565b93506200283b836200337c565b8060005b838110156200287257815162002856888262002717565b9750620028638362003431565b9250506001810190506200283f565b5085935050505092915050565b60006200288c82620033e2565b62002898818562003494565b935083602082028501620028ac856200338c565b8060005b85811015620028ee5784840389528151620028cc858262002731565b9450620028d9836200343e565b925060208a01995050600181019050620028b0565b50829750879550505050505092915050565b60006200290d82620033e2565b620029198185620034a5565b9350836020820285016200292d856200338c565b8060005b858110156200296f57848403895281516200294d858262002731565b94506200295a836200343e565b925060208a0199505060018101905062002931565b50829750879550505050505092915050565b60006200298e82620033ed565b6200299a8185620034b6565b935083602082028501620029ae856200339c565b8060005b85811015620029f05784840389528151620029ce858262002747565b9450620029db836200344b565b925060208a01995050600181019050620029b2565b50829750879550505050505092915050565b600062002a0f82620033f8565b62002a1b8185620034c7565b93508360208202850162002a2f85620033ac565b8060005b8581101562002a71578484038952815162002a4f85826200275d565b945062002a5c8362003458565b925060208a0199505060018101905062002a33565b50829750879550505050505092915050565b600062002a908262003403565b62002a9c8185620034d8565b93508360208202850162002ab085620033bc565b8060005b8581101562002af2578484038952815162002ad0858262002773565b945062002add8362003465565b925060208a0199505060018101905062002ab4565b50829750879550505050505092915050565b62002b0f8162003530565b82525050565b62002b20816200353c565b82525050565b62002b318162003546565b82525050565b600062002b44826200340e565b62002b508185620034e9565b935062002b6281856020860162003672565b62002b6d81620037dd565b840191505092915050565b62002b8d62002b8782620035be565b62003714565b82525050565b62002b9e81620035d2565b82525050565b62002baf81620035e6565b82525050565b62002bc081620035fa565b82525050565b62002bd1816200360e565b82525050565b62002be28162003622565b82525050565b62002bf38162003636565b82525050565b600062002c068262003419565b62002c128185620034fa565b935062002c2481856020860162003672565b62002c2f81620037dd565b840191505092915050565b600062002c478262003419565b62002c5381856200350b565b935062002c6581856020860162003672565b62002c7081620037dd565b840191505092915050565b600062002c8a6003836200350b565b915062002c9782620037fb565b602082019050919050565b600062002cb16004836200350b565b915062002cbe8262003824565b602082019050919050565b600062002cd86004836200350b565b915062002ce5826200384d565b602082019050919050565b600062002cff6005836200350b565b915062002d0c8262003876565b602082019050919050565b600062002d266004836200350b565b915062002d33826200389f565b602082019050919050565b600062002d4d6003836200350b565b915062002d5a82620038c8565b602082019050919050565b6000604083016000830151848203600086015262002d84828262002bf9565b9150506020830151848203602086015262002da0828262002815565b9150508091505092915050565b600060408301600083015162002dc7600086018262002789565b506020830151848203602086015262002de182826200287f565b9150508091505092915050565b600060408301600083015162002e08600086018262002789565b506020830151848203602086015262002e22828262002815565b9150508091505092915050565b62002e3a81620035a7565b82525050565b600062002e4e828462002b78565b60148201915081905092915050565b600060208201905062002e7460008301846200279a565b92915050565b600060408201905062002e9160008301856200279a565b62002ea060208301846200279a565b9392505050565b600060608201905062002ebe60008301866200279a565b62002ecd60208301856200279a565b62002edc60408301846200279a565b949350505050565b600060408201905062002efb60008301856200279a565b62002f0a602083018462002b15565b9392505050565b600060408201905062002f2860008301856200279a565b818103602083015262002f3c818462002b37565b90509392505050565b600060408201905062002f5c60008301856200279a565b62002f6b602083018462002bb5565b9392505050565b600060408201905062002f8960008301856200279a565b62002f98602083018462002be8565b9392505050565b600060408201905062002fb660008301856200279a565b62002fc5602083018462002e2f565b9392505050565b6000602082019050818103600083015262002fe88184620027ab565b905092915050565b600060208201905081810360008301526200300c818462002900565b905092915050565b6000602082019050818103600083015262003030818462002981565b905092915050565b6000602082019050818103600083015262003054818462002a02565b905092915050565b6000602082019050818103600083015262003078818462002a83565b905092915050565b600060208201905062003097600083018462002b04565b92915050565b600060a082019050620030b4600083018862002b04565b620030c3602083018762002b04565b620030d2604083018662002b04565b620030e1606083018562002b04565b620030f060808301846200279a565b9695505050505050565b6000604082019050818103600083015262003116818562002b37565b905081810360208301526200312c818462002b37565b90509392505050565b60006040820190506200314c600083018562002bd7565b6200315b60208301846200279a565b9392505050565b600060408201905062003179600083018562002bd7565b62003188602083018462002bd7565b9392505050565b60006060820190508181036000830152620031ab818662002c3a565b9050620031bc602083018562002e2f565b620031cb604083018462002b04565b949350505050565b60006040820190508181036000830152620031ee8162002ca2565b90508181036020830152620032038162002cc9565b9050919050565b6000610100820190508181036000830152620032268162002cf0565b905081810360208301526200323b8162002d3e565b90506200324c604083018962002bc6565b6200325b606083018862002bd7565b6200326a608083018762002b93565b6200327960a083018662002ba4565b6200328860c08301856200279a565b6200329760e08301846200279a565b979650505050505050565b60006040820190508181036000830152620032bd8162002d17565b90508181036020830152620032d28162002c7b565b9050919050565b6000604082019050620032f0600083018562002e2f565b818103602083015262003304818462002b37565b90509392505050565b6000620033196200332c565b9050620033278282620036de565b919050565b6000604051905090565b600067ffffffffffffffff8211156200335457620033536200379a565b5b6200335f82620037dd565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000620035298262003587565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b60008190506200358282620038f1565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000620035cb826200364a565b9050919050565b6000620035df8262003572565b9050919050565b6000620035f382620035a7565b9050919050565b60006200360782620035a7565b9050919050565b60006200361b82620035b1565b9050919050565b60006200362f82620035a7565b9050919050565b60006200364382620035a7565b9050919050565b600062003657826200365e565b9050919050565b60006200366b8262003587565b9050919050565b60005b838110156200369257808201518184015260208101905062003675565b83811115620036a2576000848401525b50505050565b60006002820490506001821680620036c157607f821691505b60208210811415620036d857620036d76200376b565b5b50919050565b620036e982620037dd565b810181811067ffffffffffffffff821117156200370b576200370a6200379a565b5b80604052505050565b6000620037218262003728565b9050919050565b60006200373582620037ee565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f54544b0000000000000000000000000000000000000000000000000000000000600082015250565b7f7a65746100000000000000000000000000000000000000000000000000000000600082015250565b7f5a45544100000000000000000000000000000000000000000000000000000000600082015250565b7f544f4b454e000000000000000000000000000000000000000000000000000000600082015250565b7f7465737400000000000000000000000000000000000000000000000000000000600082015250565b7f544b4e0000000000000000000000000000000000000000000000000000000000600082015250565b600381106200390557620039046200373c565b5b50565b620039138162003530565b81146200391f57600080fd5b50565b6200392d816200353c565b81146200393957600080fd5b5056fe60806040523480156200001157600080fd5b5060405162001813380380620018138339818101604052810190620000379190620001a3565b818181600390805190602001906200005192919062000075565b5080600490805190602001906200006a92919062000075565b5050505050620003ac565b8280546200008390620002bd565b90600052602060002090601f016020900481019282620000a75760008555620000f3565b82601f10620000c257805160ff1916838001178555620000f3565b82800160010185558215620000f3579182015b82811115620000f2578251825591602001919060010190620000d5565b5b50905062000102919062000106565b5090565b5b808211156200012157600081600090555060010162000107565b5090565b60006200013c620001368462000251565b62000228565b9050828152602081018484840111156200015b576200015a6200038c565b5b6200016884828562000287565b509392505050565b600082601f83011262000188576200018762000387565b5b81516200019a84826020860162000125565b91505092915050565b60008060408385031215620001bd57620001bc62000396565b5b600083015167ffffffffffffffff811115620001de57620001dd62000391565b5b620001ec8582860162000170565b925050602083015167ffffffffffffffff81111562000210576200020f62000391565b5b6200021e8582860162000170565b9150509250929050565b60006200023462000247565b9050620002428282620002f3565b919050565b6000604051905090565b600067ffffffffffffffff8211156200026f576200026e62000358565b5b6200027a826200039b565b9050602081019050919050565b60005b83811015620002a75780820151818401526020810190506200028a565b83811115620002b7576000848401525b50505050565b60006002820490506001821680620002d657607f821691505b60208210811415620002ed57620002ec62000329565b5b50919050565b620002fe826200039b565b810181811067ffffffffffffffff8211171562000320576200031f62000358565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b61145780620003bc6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806340c10f191161007157806340c10f19146101a357806370a08231146101bf57806395d89b41146101ef578063a457c2d71461020d578063a9059cbb1461023d578063dd62ed3e1461026d576100b4565b806306fdde03146100b9578063095ea7b3146100d757806318160ddd1461010757806323b872dd14610125578063313ce567146101555780633950935114610173575b600080fd5b6100c161029d565b6040516100ce9190610ecf565b60405180910390f35b6100f160048036038101906100ec9190610cf6565b61032f565b6040516100fe9190610eb4565b60405180910390f35b61010f610352565b60405161011c9190610ff1565b60405180910390f35b61013f600480360381019061013a9190610ca3565b61035c565b60405161014c9190610eb4565b60405180910390f35b61015d61038b565b60405161016a919061100c565b60405180910390f35b61018d60048036038101906101889190610cf6565b610394565b60405161019a9190610eb4565b60405180910390f35b6101bd60048036038101906101b89190610cf6565b6103cb565b005b6101d960048036038101906101d49190610c36565b6103d9565b6040516101e69190610ff1565b60405180910390f35b6101f7610421565b6040516102049190610ecf565b60405180910390f35b61022760048036038101906102229190610cf6565b6104b3565b6040516102349190610eb4565b60405180910390f35b61025760048036038101906102529190610cf6565b61052a565b6040516102649190610eb4565b60405180910390f35b61028760048036038101906102829190610c63565b61054d565b6040516102949190610ff1565b60405180910390f35b6060600380546102ac90611121565b80601f01602080910402602001604051908101604052809291908181526020018280546102d890611121565b80156103255780601f106102fa57610100808354040283529160200191610325565b820191906000526020600020905b81548152906001019060200180831161030857829003601f168201915b5050505050905090565b60008061033a6105d4565b90506103478185856105dc565b600191505092915050565b6000600254905090565b6000806103676105d4565b90506103748582856107a7565b61037f858585610833565b60019150509392505050565b60006012905090565b60008061039f6105d4565b90506103c08185856103b1858961054d565b6103bb9190611043565b6105dc565b600191505092915050565b6103d58282610aab565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606004805461043090611121565b80601f016020809104026020016040519081016040528092919081815260200182805461045c90611121565b80156104a95780601f1061047e576101008083540402835291602001916104a9565b820191906000526020600020905b81548152906001019060200180831161048c57829003601f168201915b5050505050905090565b6000806104be6105d4565b905060006104cc828661054d565b905083811015610511576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050890610fb1565b60405180910390fd5b61051e82868684036105dc565b60019250505092915050565b6000806105356105d4565b9050610542818585610833565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561064c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064390610f91565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156106bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b390610f11565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258360405161079a9190610ff1565b60405180910390a3505050565b60006107b3848461054d565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461082d578181101561081f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081690610f31565b60405180910390fd5b61082c84848484036105dc565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156108a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089a90610f71565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610913576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090a90610ef1565b60405180910390fd5b61091e838383610c02565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156109a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099b90610f51565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a929190610ff1565b60405180910390a3610aa5848484610c07565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610b1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1290610fd1565b60405180910390fd5b610b2760008383610c02565b8060026000828254610b399190611043565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610bea9190610ff1565b60405180910390a3610bfe60008383610c07565b5050565b505050565b505050565b600081359050610c1b816113f3565b92915050565b600081359050610c308161140a565b92915050565b600060208284031215610c4c57610c4b6111b1565b5b6000610c5a84828501610c0c565b91505092915050565b60008060408385031215610c7a57610c796111b1565b5b6000610c8885828601610c0c565b9250506020610c9985828601610c0c565b9150509250929050565b600080600060608486031215610cbc57610cbb6111b1565b5b6000610cca86828701610c0c565b9350506020610cdb86828701610c0c565b9250506040610cec86828701610c21565b9150509250925092565b60008060408385031215610d0d57610d0c6111b1565b5b6000610d1b85828601610c0c565b9250506020610d2c85828601610c21565b9150509250929050565b610d3f816110ab565b82525050565b6000610d5082611027565b610d5a8185611032565b9350610d6a8185602086016110ee565b610d73816111b6565b840191505092915050565b6000610d8b602383611032565b9150610d96826111c7565b604082019050919050565b6000610dae602283611032565b9150610db982611216565b604082019050919050565b6000610dd1601d83611032565b9150610ddc82611265565b602082019050919050565b6000610df4602683611032565b9150610dff8261128e565b604082019050919050565b6000610e17602583611032565b9150610e22826112dd565b604082019050919050565b6000610e3a602483611032565b9150610e458261132c565b604082019050919050565b6000610e5d602583611032565b9150610e688261137b565b604082019050919050565b6000610e80601f83611032565b9150610e8b826113ca565b602082019050919050565b610e9f816110d7565b82525050565b610eae816110e1565b82525050565b6000602082019050610ec96000830184610d36565b92915050565b60006020820190508181036000830152610ee98184610d45565b905092915050565b60006020820190508181036000830152610f0a81610d7e565b9050919050565b60006020820190508181036000830152610f2a81610da1565b9050919050565b60006020820190508181036000830152610f4a81610dc4565b9050919050565b60006020820190508181036000830152610f6a81610de7565b9050919050565b60006020820190508181036000830152610f8a81610e0a565b9050919050565b60006020820190508181036000830152610faa81610e2d565b9050919050565b60006020820190508181036000830152610fca81610e50565b9050919050565b60006020820190508181036000830152610fea81610e73565b9050919050565b60006020820190506110066000830184610e96565b92915050565b60006020820190506110216000830184610ea5565b92915050565b600081519050919050565b600082825260208201905092915050565b600061104e826110d7565b9150611059836110d7565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561108e5761108d611153565b5b828201905092915050565b60006110a4826110b7565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b8381101561110c5780820151818401526020810190506110f1565b8381111561111b576000848401525b50505050565b6000600282049050600182168061113957607f821691505b6020821081141561114d5761114c611182565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b6000601f19601f8301169050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6113fc81611099565b811461140757600080fd5b50565b611413816110d7565b811461141e57600080fd5b5056fea2646970667358221220d4c96329400732e284f624232c8d7228fc90c00ee7b1a48a85080262107815d864736f6c6343000807003360a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b81525034801561004657600080fd5b5060805160601c6137f0610081600039600081816107c30152818161085201528181610b7d01528181610c0c015261102801526137f06000f3fe6080604052600436106101355760003560e01c8063715018a6116100ab578063b8969bd41161006f578063b8969bd4146103b4578063dda79b75146103dd578063f2fde38b14610408578063f31e62bf14610431578063f340fa011461045c578063f45346dc1461047857610135565b8063715018a6146103045780638c6f037f1461031b5780638da5cb5b146103445780638ee0f9f21461036f578063ae7a3a6f1461038b57610135565b8063485cc955116100fd578063485cc955146102015780634f1ef2861461022a5780635131ab591461024657806352d1902d1461028357806357bec62f146102ae5780635b112591146102d957610135565b806310188aef1461013a5780631b8b921d146101635780631cff79cd1461018c57806329c59b5d146101bc5780633659cfe6146101d8575b600080fd5b34801561014657600080fd5b50610161600480360381019061015c919061275b565b6104a1565b005b34801561016f57600080fd5b5061018a60048036038101906101859190612850565b61056d565b005b6101a660048036038101906101a19190612850565b6105d9565b6040516101b39190612ee3565b60405180910390f35b6101d660048036038101906101d19190612850565b610647565b005b3480156101e457600080fd5b506101ff60048036038101906101fa919061275b565b6107c1565b005b34801561020d57600080fd5b5061022860048036038101906102239190612788565b61094a565b005b610244600480360381019061023f91906128b0565b610b7b565b005b34801561025257600080fd5b5061026d600480360381019061026891906127c8565b610cb8565b60405161027a9190612ee3565b60405180910390f35b34801561028f57600080fd5b50610298611024565b6040516102a59190612ea4565b60405180910390f35b3480156102ba57600080fd5b506102c36110dd565b6040516102d09190612e00565b60405180910390f35b3480156102e557600080fd5b506102ee611103565b6040516102fb9190612e00565b60405180910390f35b34801561031057600080fd5b50610319611129565b005b34801561032757600080fd5b50610342600480360381019061033d919061295f565b61113d565b005b34801561035057600080fd5b506103596112bb565b6040516103669190612e00565b60405180910390f35b61038960048036038101906103849190612850565b6112e5565b005b34801561039757600080fd5b506103b260048036038101906103ad919061275b565b611454565b005b3480156103c057600080fd5b506103db60048036038101906103d691906127c8565b611520565b005b3480156103e957600080fd5b506103f2611663565b6040516103ff9190612e00565b60405180910390f35b34801561041457600080fd5b5061042f600480360381019061042a919061275b565b611689565b005b34801561043d57600080fd5b5061044661170d565b6040516104539190612e00565b60405180910390f35b6104766004803603810190610471919061275b565b611733565b005b34801561048457600080fd5b5061049f600480360381019061049a919061290c565b6118a7565b005b600073ffffffffffffffffffffffffffffffffffffffff1660cb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610529576040517fb337f37800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060cb60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde384846040516105cc929190612ebf565b60405180910390a3505050565b606060006105e8858585611a1f565b90508473ffffffffffffffffffffffffffffffffffffffff167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f3486866040516106349392919061315e565b60405180910390a2809150509392505050565b6000341415610682576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16346040516106ca90612deb565b60006040518083038185875af1925050503d8060008114610707576040519150601f19603f3d011682016040523d82523d6000602084013e61070c565b606091505b5050905060001515811515141561074f576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a434600087876040516107b394939291906130e2565b60405180910390a350505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610850576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084790612f62565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661088f611ad6565b73ffffffffffffffffffffffffffffffffffffffff16146108e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108dc90612f82565b60405180910390fd5b6108ee81611b2d565b61094781600067ffffffffffffffff81111561090d5761090c61331f565b5b6040519080825280601f01601f19166020018201604052801561093f5781602001600182028036833780820191505090505b506000611b38565b50565b60008060019054906101000a900460ff1615905080801561097b5750600160008054906101000a900460ff1660ff16105b806109a8575061098a30611cb5565b1580156109a75750600160008054906101000a900460ff1660ff16145b5b6109e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109de90613002565b60405180910390fd5b60016000806101000a81548160ff021916908360ff1602179055508015610a24576001600060016101000a81548160ff0219169083151502179055505b610a2c611cd8565b610a34611d31565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610a9b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8260ca60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160cc60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015610b765760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610b6d9190612f05565b60405180910390a15b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610c0a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0190612f62565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610c49611ad6565b73ffffffffffffffffffffffffffffffffffffffff1614610c9f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9690612f82565b60405180910390fd5b610ca882611b2d565b610cb482826001611b38565b5050565b60606000841415610cf5576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cff8686611d82565b610d35576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff1663095ea7b386866040518363ffffffff1660e01b8152600401610d70929190612e7b565b602060405180830381600087803b158015610d8a57600080fd5b505af1158015610d9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc291906129e7565b610df8576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610e05868585611a1f565b9050610e118787611d82565b610e47576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008773ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610e829190612e00565b60206040518083038186803b158015610e9a57600080fd5b505afa158015610eae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed29190612a41565b90506000811115610fad57600060c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060cc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff161415610f805760cb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b610fab81838b73ffffffffffffffffffffffffffffffffffffffff16611e1a9092919063ffffffff16565b505b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b738288888860405161100e9392919061315e565b60405180910390a3819250505095945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16146110b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110ab90612fc2565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b60cb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611131611ea0565b61113b6000611f1e565b565b6000841415611178576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060cc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561121b5760cb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b6112483382878773ffffffffffffffffffffffffffffffffffffffff16611fe4909392919063ffffffff16565b8573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4878787876040516112ab94939291906130e2565b60405180910390a3505050505050565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000808473ffffffffffffffffffffffffffffffffffffffff16348585604051611310929190612dbb565b60006040518083038185875af1925050503d806000811461134d576040519150601f19603f3d011682016040523d82523d6000602084013e611352565b606091505b50915091508161138e576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff16638fcaa0b585856040518363ffffffff1660e01b81526004016113c9929190612ebf565b600060405180830381600087803b1580156113e357600080fd5b505af11580156113f7573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff167fd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c3486866040516114459392919061315e565b60405180910390a25050505050565b600073ffffffffffffffffffffffffffffffffffffffff1660c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146114dc576040517fb337f37800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060c960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600083141561155b576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61158684848773ffffffffffffffffffffffffffffffffffffffff16611e1a9092919063ffffffff16565b8373ffffffffffffffffffffffffffffffffffffffff16638fcaa0b583836040518363ffffffff1660e01b81526004016115c1929190612ebf565b600060405180830381600087803b1580156115db57600080fd5b505af11580156115ef573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda78585856040516116549392919061315e565b60405180910390a35050505050565b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611691611ea0565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611701576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116f890612f42565b60405180910390fd5b61170a81611f1e565b50565b60cc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600034141561176e576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16346040516117b690612deb565b60006040518083038185875af1925050503d80600081146117f3576040519150601f19603f3d011682016040523d82523d6000602084013e6117f8565b606091505b5050905060001515811515141561183b576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a434600060405161189b929190613122565b60405180910390a35050565b60008214156118e2576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060cc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156119855760cb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b6119b23382858573ffffffffffffffffffffffffffffffffffffffff16611fe4909392919063ffffffff16565b8373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a48585604051611a11929190613122565b60405180910390a350505050565b60606000808573ffffffffffffffffffffffffffffffffffffffff16348686604051611a4c929190612dbb565b60006040518083038185875af1925050503d8060008114611a89576040519150601f19603f3d011682016040523d82523d6000602084013e611a8e565b606091505b509150915081611aca576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80925050509392505050565b6000611b047f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61206d565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611b35611ea0565b50565b611b647f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b612077565b60000160009054906101000a900460ff1615611b8857611b8383612081565b611cb0565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611bce57600080fd5b505afa925050508015611bff57506040513d601f19601f82011682018060405250810190611bfc9190612a14565b60015b611c3e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c3590613022565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114611ca3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c9a90612fe2565b60405180910390fd5b50611caf83838361213a565b5b505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff16611d27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d1e906130a2565b60405180910390fd5b611d2f612166565b565b600060019054906101000a900460ff16611d80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d77906130a2565b60405180910390fd5b565b60008273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401611dc0929190612e52565b602060405180830381600087803b158015611dda57600080fd5b505af1158015611dee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e1291906129e7565b905092915050565b611e9b8363a9059cbb60e01b8484604051602401611e39929190612e7b565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506121c7565b505050565b611ea861228e565b73ffffffffffffffffffffffffffffffffffffffff16611ec66112bb565b73ffffffffffffffffffffffffffffffffffffffff1614611f1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f1390613062565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b612067846323b872dd60e01b85858560405160240161200593929190612e1b565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506121c7565b50505050565b6000819050919050565b6000819050919050565b61208a81611cb5565b6120c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120c090613042565b60405180910390fd5b806120f67f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61206d565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61214383612296565b6000825111806121505750805b156121615761215f83836122e5565b505b505050565b600060019054906101000a900460ff166121b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121ac906130a2565b60405180910390fd5b6121c56121c061228e565b611f1e565b565b6000612229826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166123129092919063ffffffff16565b9050600081511115612289578080602001905181019061224991906129e7565b612288576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161227f906130c2565b60405180910390fd5b5b505050565b600033905090565b61229f81612081565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b606061230a83836040518060600160405280602781526020016137946027913961232a565b905092915050565b606061232184846000856123b0565b90509392505050565b60606000808573ffffffffffffffffffffffffffffffffffffffff16856040516123549190612dd4565b600060405180830381855af49150503d806000811461238f576040519150601f19603f3d011682016040523d82523d6000602084013e612394565b606091505b50915091506123a58683838761247d565b925050509392505050565b6060824710156123f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123ec90612fa2565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161241e9190612dd4565b60006040518083038185875af1925050503d806000811461245b576040519150601f19603f3d011682016040523d82523d6000602084013e612460565b606091505b5091509150612471878383876124f3565b92505050949350505050565b606083156124e0576000835114156124d85761249885611cb5565b6124d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124ce90613082565b60405180910390fd5b5b8290506124eb565b6124ea8383612569565b5b949350505050565b606083156125565760008351141561254e5761250e856125b9565b61254d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161254490613082565b60405180910390fd5b5b829050612561565b61256083836125dc565b5b949350505050565b60008251111561257c5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125b09190612f20565b60405180910390fd5b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156125ef5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126239190612f20565b60405180910390fd5b600061263f61263a846131b5565b613190565b90508281526020810184848401111561265b5761265a61335d565b5b6126668482856132ac565b509392505050565b60008135905061267d81613737565b92915050565b6000815190506126928161374e565b92915050565b6000815190506126a781613765565b92915050565b60008083601f8401126126c3576126c2613353565b5b8235905067ffffffffffffffff8111156126e0576126df61334e565b5b6020830191508360018202830111156126fc576126fb613358565b5b9250929050565b600082601f83011261271857612717613353565b5b813561272884826020860161262c565b91505092915050565b6000813590506127408161377c565b92915050565b6000815190506127558161377c565b92915050565b60006020828403121561277157612770613367565b5b600061277f8482850161266e565b91505092915050565b6000806040838503121561279f5761279e613367565b5b60006127ad8582860161266e565b92505060206127be8582860161266e565b9150509250929050565b6000806000806000608086880312156127e4576127e3613367565b5b60006127f28882890161266e565b95505060206128038882890161266e565b945050604061281488828901612731565b935050606086013567ffffffffffffffff81111561283557612834613362565b5b612841888289016126ad565b92509250509295509295909350565b60008060006040848603121561286957612868613367565b5b60006128778682870161266e565b935050602084013567ffffffffffffffff81111561289857612897613362565b5b6128a4868287016126ad565b92509250509250925092565b600080604083850312156128c7576128c6613367565b5b60006128d58582860161266e565b925050602083013567ffffffffffffffff8111156128f6576128f5613362565b5b61290285828601612703565b9150509250929050565b60008060006060848603121561292557612924613367565b5b60006129338682870161266e565b935050602061294486828701612731565b92505060406129558682870161266e565b9150509250925092565b60008060008060006080868803121561297b5761297a613367565b5b60006129898882890161266e565b955050602061299a88828901612731565b94505060406129ab8882890161266e565b935050606086013567ffffffffffffffff8111156129cc576129cb613362565b5b6129d8888289016126ad565b92509250509295509295909350565b6000602082840312156129fd576129fc613367565b5b6000612a0b84828501612683565b91505092915050565b600060208284031215612a2a57612a29613367565b5b6000612a3884828501612698565b91505092915050565b600060208284031215612a5757612a56613367565b5b6000612a6584828501612746565b91505092915050565b612a7781613229565b82525050565b612a8681613247565b82525050565b6000612a9883856131fc565b9350612aa58385846132ac565b612aae8361336c565b840190509392505050565b6000612ac5838561320d565b9350612ad28385846132ac565b82840190509392505050565b6000612ae9826131e6565b612af381856131fc565b9350612b038185602086016132bb565b612b0c8161336c565b840191505092915050565b6000612b22826131e6565b612b2c818561320d565b9350612b3c8185602086016132bb565b80840191505092915050565b612b5181613288565b82525050565b612b608161329a565b82525050565b6000612b71826131f1565b612b7b8185613218565b9350612b8b8185602086016132bb565b612b948161336c565b840191505092915050565b6000612bac602683613218565b9150612bb78261337d565b604082019050919050565b6000612bcf602c83613218565b9150612bda826133cc565b604082019050919050565b6000612bf2602c83613218565b9150612bfd8261341b565b604082019050919050565b6000612c15602683613218565b9150612c208261346a565b604082019050919050565b6000612c38603883613218565b9150612c43826134b9565b604082019050919050565b6000612c5b602983613218565b9150612c6682613508565b604082019050919050565b6000612c7e602e83613218565b9150612c8982613557565b604082019050919050565b6000612ca1602e83613218565b9150612cac826135a6565b604082019050919050565b6000612cc4602d83613218565b9150612ccf826135f5565b604082019050919050565b6000612ce7602083613218565b9150612cf282613644565b602082019050919050565b6000612d0a6000836131fc565b9150612d158261366d565b600082019050919050565b6000612d2d60008361320d565b9150612d388261366d565b600082019050919050565b6000612d50601d83613218565b9150612d5b82613670565b602082019050919050565b6000612d73602b83613218565b9150612d7e82613699565b604082019050919050565b6000612d96602a83613218565b9150612da1826136e8565b604082019050919050565b612db581613271565b82525050565b6000612dc8828486612ab9565b91508190509392505050565b6000612de08284612b17565b915081905092915050565b6000612df682612d20565b9150819050919050565b6000602082019050612e156000830184612a6e565b92915050565b6000606082019050612e306000830186612a6e565b612e3d6020830185612a6e565b612e4a6040830184612dac565b949350505050565b6000604082019050612e676000830185612a6e565b612e746020830184612b48565b9392505050565b6000604082019050612e906000830185612a6e565b612e9d6020830184612dac565b9392505050565b6000602082019050612eb96000830184612a7d565b92915050565b60006020820190508181036000830152612eda818486612a8c565b90509392505050565b60006020820190508181036000830152612efd8184612ade565b905092915050565b6000602082019050612f1a6000830184612b57565b92915050565b60006020820190508181036000830152612f3a8184612b66565b905092915050565b60006020820190508181036000830152612f5b81612b9f565b9050919050565b60006020820190508181036000830152612f7b81612bc2565b9050919050565b60006020820190508181036000830152612f9b81612be5565b9050919050565b60006020820190508181036000830152612fbb81612c08565b9050919050565b60006020820190508181036000830152612fdb81612c2b565b9050919050565b60006020820190508181036000830152612ffb81612c4e565b9050919050565b6000602082019050818103600083015261301b81612c71565b9050919050565b6000602082019050818103600083015261303b81612c94565b9050919050565b6000602082019050818103600083015261305b81612cb7565b9050919050565b6000602082019050818103600083015261307b81612cda565b9050919050565b6000602082019050818103600083015261309b81612d43565b9050919050565b600060208201905081810360008301526130bb81612d66565b9050919050565b600060208201905081810360008301526130db81612d89565b9050919050565b60006060820190506130f76000830187612dac565b6131046020830186612a6e565b8181036040830152613117818486612a8c565b905095945050505050565b60006060820190506131376000830185612dac565b6131446020830184612a6e565b818103604083015261315581612cfd565b90509392505050565b60006040820190506131736000830186612dac565b8181036020830152613186818486612a8c565b9050949350505050565b600061319a6131ab565b90506131a682826132ee565b919050565b6000604051905090565b600067ffffffffffffffff8211156131d0576131cf61331f565b5b6131d98261336c565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061323482613251565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b600061329382613271565b9050919050565b60006132a58261327b565b9050919050565b82818337600083830152505050565b60005b838110156132d95780820151818401526020810190506132be565b838111156132e8576000848401525b50505050565b6132f78261336c565b810181811067ffffffffffffffff821117156133165761331561331f565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b61374081613229565b811461374b57600080fd5b50565b6137578161323b565b811461376257600080fd5b50565b61376e81613247565b811461377957600080fd5b50565b61378581613271565b811461379057600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220dd33fe9a4572916096cea820dc23949deb4f9b84ebfbd486b54cd5b793168ab164736f6c6343000807003360806040523480156200001157600080fd5b506040516200123f3803806200123f833981810160405281019062000037919062000106565b6001600081905550600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620000a7576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200018b565b600081519050620001008162000171565b92915050565b6000602082840312156200011f576200011e6200016c565b5b60006200012f84828501620000ef565b91505092915050565b600062000145826200014c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200017c8162000138565b81146200018857600080fd5b50565b6110a4806200019b6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063116191b61461005157806321fc65f21461006f578063c8a023621461008b578063d9caed12146100a7575b600080fd5b6100596100c3565b6040516100669190610c47565b60405180910390f35b6100896004803603810190610084919061096b565b6100e9565b005b6100a560048036038101906100a0919061096b565b610271565b005b6100c160048036038101906100bc9190610918565b6103f9565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6100f161049e565b61013e600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848773ffffffffffffffffffffffffffffffffffffffff166104ee9092919063ffffffff16565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635131ab5986868686866040518663ffffffff1660e01b81526004016101a1959493929190610bd0565b600060405180830381600087803b1580156101bb57600080fd5b505af11580156101cf573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906101f89190610a20565b508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e85858560405161025a93929190610d1f565b60405180910390a361026a610574565b5050505050565b61027961049e565b6102c6600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848773ffffffffffffffffffffffffffffffffffffffff166104ee9092919063ffffffff16565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635131ab5986868686866040518663ffffffff1660e01b8152600401610329959493929190610bd0565b600060405180830381600087803b15801561034357600080fd5b505af1158015610357573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103809190610a20565b508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c88585856040516103e293929190610d1f565b60405180910390a36103f2610574565b5050505050565b61040161049e565b61042c82828573ffffffffffffffffffffffffffffffffffffffff166104ee9092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb836040516104899190610d04565b60405180910390a3610499610574565b505050565b600260005414156104e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104db90610ce4565b60405180910390fd5b6002600081905550565b61056f8363a9059cbb60e01b848460405160240161050d929190610c1e565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061057e565b505050565b6001600081905550565b60006105e0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166106459092919063ffffffff16565b9050600081511115610640578080602001905181019061060091906109f3565b61063f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063690610cc4565b60405180910390fd5b5b505050565b6060610654848460008561065d565b90509392505050565b6060824710156106a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069990610c84565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516106cb9190610bb9565b60006040518083038185875af1925050503d8060008114610708576040519150601f19603f3d011682016040523d82523d6000602084013e61070d565b606091505b509150915061071e8783838761072a565b92505050949350505050565b6060831561078d5760008351141561078557610745856107a0565b610784576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077b90610ca4565b60405180910390fd5b5b829050610798565b61079783836107c3565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156107d65781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080a9190610c62565b60405180910390fd5b600061082661082184610d76565b610d51565b90508281526020810184848401111561084257610841610f19565b5b61084d848285610e77565b509392505050565b60008135905061086481611029565b92915050565b60008151905061087981611040565b92915050565b60008083601f84011261089557610894610f0f565b5b8235905067ffffffffffffffff8111156108b2576108b1610f0a565b5b6020830191508360018202830111156108ce576108cd610f14565b5b9250929050565b600082601f8301126108ea576108e9610f0f565b5b81516108fa848260208601610813565b91505092915050565b60008135905061091281611057565b92915050565b60008060006060848603121561093157610930610f23565b5b600061093f86828701610855565b935050602061095086828701610855565b925050604061096186828701610903565b9150509250925092565b60008060008060006080868803121561098757610986610f23565b5b600061099588828901610855565b95505060206109a688828901610855565b94505060406109b788828901610903565b935050606086013567ffffffffffffffff8111156109d8576109d7610f1e565b5b6109e48882890161087f565b92509250509295509295909350565b600060208284031215610a0957610a08610f23565b5b6000610a178482850161086a565b91505092915050565b600060208284031215610a3657610a35610f23565b5b600082015167ffffffffffffffff811115610a5457610a53610f1e565b5b610a60848285016108d5565b91505092915050565b610a7281610dea565b82525050565b6000610a848385610dbd565b9350610a91838584610e68565b610a9a83610f28565b840190509392505050565b6000610ab082610da7565b610aba8185610dce565b9350610aca818560208601610e77565b80840191505092915050565b610adf81610e32565b82525050565b6000610af082610db2565b610afa8185610dd9565b9350610b0a818560208601610e77565b610b1381610f28565b840191505092915050565b6000610b2b602683610dd9565b9150610b3682610f39565b604082019050919050565b6000610b4e601d83610dd9565b9150610b5982610f88565b602082019050919050565b6000610b71602a83610dd9565b9150610b7c82610fb1565b604082019050919050565b6000610b94601f83610dd9565b9150610b9f82611000565b602082019050919050565b610bb381610e28565b82525050565b6000610bc58284610aa5565b915081905092915050565b6000608082019050610be56000830188610a69565b610bf26020830187610a69565b610bff6040830186610baa565b8181036060830152610c12818486610a78565b90509695505050505050565b6000604082019050610c336000830185610a69565b610c406020830184610baa565b9392505050565b6000602082019050610c5c6000830184610ad6565b92915050565b60006020820190508181036000830152610c7c8184610ae5565b905092915050565b60006020820190508181036000830152610c9d81610b1e565b9050919050565b60006020820190508181036000830152610cbd81610b41565b9050919050565b60006020820190508181036000830152610cdd81610b64565b9050919050565b60006020820190508181036000830152610cfd81610b87565b9050919050565b6000602082019050610d196000830184610baa565b92915050565b6000604082019050610d346000830186610baa565b8181036020830152610d47818486610a78565b9050949350505050565b6000610d5b610d6c565b9050610d678282610eaa565b919050565b6000604051905090565b600067ffffffffffffffff821115610d9157610d90610edb565b5b610d9a82610f28565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610df582610e08565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610e3d82610e44565b9050919050565b6000610e4f82610e56565b9050919050565b6000610e6182610e08565b9050919050565b82818337600083830152505050565b60005b83811015610e95578082015181840152602081019050610e7a565b83811115610ea4576000848401525b50505050565b610eb382610f28565b810181811067ffffffffffffffff82111715610ed257610ed1610edb565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b61103281610dea565b811461103d57600080fd5b50565b61104981610dfc565b811461105457600080fd5b50565b61106081610e28565b811461106b57600080fd5b5056fea264697066735822122034916fcf763546c650529df8ac1ce35e3480b605125b539d009b0bc21e5a4c0d64736f6c6343000807003360806040523480156200001157600080fd5b50604051620011b9380380620011b9833981810160405281019062000037919062000180565b6001600081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480620000a75750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15620000df576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506200021a565b6000815190506200017a8162000200565b92915050565b600080604083850312156200019a5762000199620001fb565b5b6000620001aa8582860162000169565b9250506020620001bd8582860162000169565b9150509250929050565b6000620001d482620001db565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200020b81620001c7565b81146200021757600080fd5b50565b610f8f806200022a6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806302a04c0d14610051578063116191b61461006d578063e8f9cb3a1461008b578063f3fef3a3146100a9575b600080fd5b61006b6004803603810190610066919061082e565b6100c5565b005b610075610279565b6040516100829190610b20565b60405180910390f35b61009361029f565b6040516100a09190610b05565b60405180910390f35b6100c360048036038101906100be91906107ee565b6102c5565b005b6100cd610374565b61013c600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166103c49092919063ffffffff16565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635131ab59600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16868686866040518663ffffffff1660e01b81526004016101c1959493929190610a8e565b600060405180830381600087803b1580156101db57600080fd5b505af11580156101ef573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061021891906108cf565b508373ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced84848460405161026393929190610bf8565b60405180910390a261027361044a565b50505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6102cd610374565b61031a8282600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166103c49092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364826040516103609190610bdd565b60405180910390a261037061044a565b5050565b600260005414156103ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b190610bbd565b60405180910390fd5b6002600081905550565b6104458363a9059cbb60e01b84846040516024016103e3929190610adc565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610454565b505050565b6001600081905550565b60006104b6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661051b9092919063ffffffff16565b905060008151111561051657808060200190518101906104d691906108a2565b610515576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050c90610b9d565b60405180910390fd5b5b505050565b606061052a8484600085610533565b90509392505050565b606082471015610578576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056f90610b5d565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516105a19190610a77565b60006040518083038185875af1925050503d80600081146105de576040519150601f19603f3d011682016040523d82523d6000602084013e6105e3565b606091505b50915091506105f487838387610600565b92505050949350505050565b606083156106635760008351141561065b5761061b85610676565b61065a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065190610b7d565b60405180910390fd5b5b82905061066e565b61066d8383610699565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156106ac5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e09190610b3b565b60405180910390fd5b60006106fc6106f784610c4f565b610c2a565b90508281526020810184848401111561071857610717610e04565b5b610723848285610d62565b509392505050565b60008135905061073a81610f14565b92915050565b60008151905061074f81610f2b565b92915050565b60008083601f84011261076b5761076a610dfa565b5b8235905067ffffffffffffffff81111561078857610787610df5565b5b6020830191508360018202830111156107a4576107a3610dff565b5b9250929050565b600082601f8301126107c0576107bf610dfa565b5b81516107d08482602086016106e9565b91505092915050565b6000813590506107e881610f42565b92915050565b6000806040838503121561080557610804610e0e565b5b60006108138582860161072b565b9250506020610824858286016107d9565b9150509250929050565b6000806000806060858703121561084857610847610e0e565b5b60006108568782880161072b565b9450506020610867878288016107d9565b935050604085013567ffffffffffffffff81111561088857610887610e09565b5b61089487828801610755565b925092505092959194509250565b6000602082840312156108b8576108b7610e0e565b5b60006108c684828501610740565b91505092915050565b6000602082840312156108e5576108e4610e0e565b5b600082015167ffffffffffffffff81111561090357610902610e09565b5b61090f848285016107ab565b91505092915050565b61092181610cc3565b82525050565b60006109338385610c96565b9350610940838584610d53565b61094983610e13565b840190509392505050565b600061095f82610c80565b6109698185610ca7565b9350610979818560208601610d62565b80840191505092915050565b61098e81610d0b565b82525050565b61099d81610d1d565b82525050565b60006109ae82610c8b565b6109b88185610cb2565b93506109c8818560208601610d62565b6109d181610e13565b840191505092915050565b60006109e9602683610cb2565b91506109f482610e24565b604082019050919050565b6000610a0c601d83610cb2565b9150610a1782610e73565b602082019050919050565b6000610a2f602a83610cb2565b9150610a3a82610e9c565b604082019050919050565b6000610a52601f83610cb2565b9150610a5d82610eeb565b602082019050919050565b610a7181610d01565b82525050565b6000610a838284610954565b915081905092915050565b6000608082019050610aa36000830188610918565b610ab06020830187610918565b610abd6040830186610a68565b8181036060830152610ad0818486610927565b90509695505050505050565b6000604082019050610af16000830185610918565b610afe6020830184610a68565b9392505050565b6000602082019050610b1a6000830184610985565b92915050565b6000602082019050610b356000830184610994565b92915050565b60006020820190508181036000830152610b5581846109a3565b905092915050565b60006020820190508181036000830152610b76816109dc565b9050919050565b60006020820190508181036000830152610b96816109ff565b9050919050565b60006020820190508181036000830152610bb681610a22565b9050919050565b60006020820190508181036000830152610bd681610a45565b9050919050565b6000602082019050610bf26000830184610a68565b92915050565b6000604082019050610c0d6000830186610a68565b8181036020830152610c20818486610927565b9050949350505050565b6000610c34610c45565b9050610c408282610d95565b919050565b6000604051905090565b600067ffffffffffffffff821115610c6a57610c69610dc6565b5b610c7382610e13565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610cce82610ce1565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610d1682610d2f565b9050919050565b6000610d2882610d2f565b9050919050565b6000610d3a82610d41565b9050919050565b6000610d4c82610ce1565b9050919050565b82818337600083830152505050565b60005b83811015610d80578082015181840152602081019050610d65565b83811115610d8f576000848401525b50505050565b610d9e82610e13565b810181811067ffffffffffffffff82111715610dbd57610dbc610dc6565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b610f1d81610cc3565b8114610f2857600080fd5b50565b610f3481610cd5565b8114610f3f57600080fd5b50565b610f4b81610d01565b8114610f5657600080fd5b5056fea2646970667358221220ffaaf5577f8164f6527ef50d54a96d5637c72ba10a83df84e4f7c8118e9d835364736f6c63430008070033608060405234801561001057600080fd5b5061108a806100206000396000f3fe60806040526004361061003f5760003560e01c8063357fc5a2146100445780636ed701691461006d578063e04d4f9714610084578063f05b6abf146100a0575b600080fd5b34801561005057600080fd5b5061006b6004803603810190610066919061085a565b6100c9565b005b34801561007957600080fd5b50610082610138565b005b61009e600480360381019061009991906107eb565b610171565b005b3480156100ac57600080fd5b506100c760048036038101906100c29190610733565b6101b5565b005b6100f63382858573ffffffffffffffffffffffffffffffffffffffff166101f7909392919063ffffffff16565b7f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af603384848460405161012b9493929190610bb0565b60405180910390a1505050565b7fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0336040516101679190610b0b565b60405180910390a1565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa33348585856040516101a8959493929190610bf5565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146338484846040516101ea9493929190610b5d565b60405180910390a1505050565b61027a846323b872dd60e01b85858560405160240161021893929190610b26565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610280565b50505050565b60006102e2826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166103479092919063ffffffff16565b9050600081511115610342578080602001905181019061030291906107be565b610341576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033890610cb1565b60405180910390fd5b5b505050565b6060610356848460008561035f565b90509392505050565b6060824710156103a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039b90610c71565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516103cd9190610af4565b60006040518083038185875af1925050503d806000811461040a576040519150601f19603f3d011682016040523d82523d6000602084013e61040f565b606091505b50915091506104208783838761042c565b92505050949350505050565b6060831561048f5760008351141561048757610447856104a2565b610486576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047d90610c91565b60405180910390fd5b5b82905061049a565b61049983836104c5565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156104d85781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050c9190610c4f565b60405180910390fd5b600061052861052384610cf6565b610cd1565b9050808382526020820190508285602086028201111561054b5761054a610f23565b5b60005b8581101561059957813567ffffffffffffffff81111561057157610570610f1e565b5b80860161057e89826106f0565b8552602085019450602084019350505060018101905061054e565b5050509392505050565b60006105b66105b184610d22565b610cd1565b905080838252602082019050828560208602820111156105d9576105d8610f23565b5b60005b8581101561060957816105ef888261071e565b8452602084019350602083019250506001810190506105dc565b5050509392505050565b600061062661062184610d4e565b610cd1565b90508281526020810184848401111561064257610641610f28565b5b61064d848285610e7c565b509392505050565b6000813590506106648161100f565b92915050565b600082601f83011261067f5761067e610f1e565b5b813561068f848260208601610515565b91505092915050565b600082601f8301126106ad576106ac610f1e565b5b81356106bd8482602086016105a3565b91505092915050565b6000813590506106d581611026565b92915050565b6000815190506106ea81611026565b92915050565b600082601f83011261070557610704610f1e565b5b8135610715848260208601610613565b91505092915050565b60008135905061072d8161103d565b92915050565b60008060006060848603121561074c5761074b610f32565b5b600084013567ffffffffffffffff81111561076a57610769610f2d565b5b6107768682870161066a565b935050602084013567ffffffffffffffff81111561079757610796610f2d565b5b6107a386828701610698565b92505060406107b4868287016106c6565b9150509250925092565b6000602082840312156107d4576107d3610f32565b5b60006107e2848285016106db565b91505092915050565b60008060006060848603121561080457610803610f32565b5b600084013567ffffffffffffffff81111561082257610821610f2d565b5b61082e868287016106f0565b935050602061083f8682870161071e565b9250506040610850868287016106c6565b9150509250925092565b60008060006060848603121561087357610872610f32565b5b60006108818682870161071e565b935050602061089286828701610655565b92505060406108a386828701610655565b9150509250925092565b60006108b983836109fb565b905092915050565b60006108cd8383610ad6565b60208301905092915050565b6108e281610e34565b82525050565b60006108f382610d9f565b6108fd8185610de5565b93508360208202850161090f85610d7f565b8060005b8581101561094b578484038952815161092c85826108ad565b945061093783610dcb565b925060208a01995050600181019050610913565b50829750879550505050505092915050565b600061096882610daa565b6109728185610df6565b935061097d83610d8f565b8060005b838110156109ae57815161099588826108c1565b97506109a083610dd8565b925050600181019050610981565b5085935050505092915050565b6109c481610e46565b82525050565b60006109d582610db5565b6109df8185610e07565b93506109ef818560208601610e8b565b80840191505092915050565b6000610a0682610dc0565b610a108185610e12565b9350610a20818560208601610e8b565b610a2981610f37565b840191505092915050565b6000610a3f82610dc0565b610a498185610e23565b9350610a59818560208601610e8b565b610a6281610f37565b840191505092915050565b6000610a7a602683610e23565b9150610a8582610f48565b604082019050919050565b6000610a9d601d83610e23565b9150610aa882610f97565b602082019050919050565b6000610ac0602a83610e23565b9150610acb82610fc0565b604082019050919050565b610adf81610e72565b82525050565b610aee81610e72565b82525050565b6000610b0082846109ca565b915081905092915050565b6000602082019050610b2060008301846108d9565b92915050565b6000606082019050610b3b60008301866108d9565b610b4860208301856108d9565b610b556040830184610ae5565b949350505050565b6000608082019050610b7260008301876108d9565b8181036020830152610b8481866108e8565b90508181036040830152610b98818561095d565b9050610ba760608301846109bb565b95945050505050565b6000608082019050610bc560008301876108d9565b610bd26020830186610ae5565b610bdf60408301856108d9565b610bec60608301846108d9565b95945050505050565b600060a082019050610c0a60008301886108d9565b610c176020830187610ae5565b8181036040830152610c298186610a34565b9050610c386060830185610ae5565b610c4560808301846109bb565b9695505050505050565b60006020820190508181036000830152610c698184610a34565b905092915050565b60006020820190508181036000830152610c8a81610a6d565b9050919050565b60006020820190508181036000830152610caa81610a90565b9050919050565b60006020820190508181036000830152610cca81610ab3565b9050919050565b6000610cdb610cec565b9050610ce78282610ebe565b919050565b6000604051905090565b600067ffffffffffffffff821115610d1157610d10610eef565b5b602082029050602081019050919050565b600067ffffffffffffffff821115610d3d57610d3c610eef565b5b602082029050602081019050919050565b600067ffffffffffffffff821115610d6957610d68610eef565b5b610d7282610f37565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610e3f82610e52565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610ea9578082015181840152602081019050610e8e565b83811115610eb8576000848401525b50505050565b610ec782610f37565b810181811067ffffffffffffffff82111715610ee657610ee5610eef565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b61101881610e34565b811461102357600080fd5b50565b61102f81610e46565b811461103a57600080fd5b50565b61104681610e72565b811461105157600080fd5b5056fea2646970667358221220086a1c9f56ed96506c7198d50cf431a5b03003c16cd4168d2395cdedfc1ac06164736f6c6343000807003360a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b8152503480156200004757600080fd5b50620000586200005e60201b60201c565b62000208565b600060019054906101000a900460ff1615620000b1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a8906200015c565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff1614620001225760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200011991906200017e565b60405180910390a15b565b6000620001336027836200019b565b91506200014082620001b9565b604082019050919050565b6200015681620001ac565b82525050565b60006020820190508181036000830152620001778162000124565b9050919050565b60006020820190506200019560008301846200014b565b92915050565b600082825260208201905092915050565b600060ff82169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60805160601c613cf56200024360003960008181610a8001528181610b0f01528181610d1501528181610da40152610e540152613cf56000f3fe60806040526004361061011e5760003560e01c806352d1902d116100a0578063bcf7f32b11610064578063bcf7f32b14610373578063c39aca371461039c578063c4d66de8146103c5578063f2fde38b146103ee578063f45346dc146104175761011e565b806352d1902d146102b4578063715018a6146102df5780637993c1e0146102f65780638da5cb5b1461031f578063a613f3bb1461034a5761011e565b80632e1a7d4d116100e75780632e1a7d4d146101f25780633659cfe61461021b5780633ce4a5bc14610244578063430e8e321461026f5780634f1ef286146102985761011e565b8062173d46146101235780630ac7c44c1461014e578063135390f91461017757806321501a95146101a0578063267e75a0146101c9575b600080fd5b34801561012f57600080fd5b50610138610440565b604051610145919061315f565b60405180910390f35b34801561015a57600080fd5b50610175600480360381019061017091906128a7565b610466565b005b34801561018357600080fd5b5061019e60048036038101906101999190612923565b6104bd565b005b3480156101ac57600080fd5b506101c760048036038101906101c29190612ba2565b6105a4565b005b3480156101d557600080fd5b506101f060048036038101906101eb9190612ca0565b61094a565b005b3480156101fe57600080fd5b5061021960048036038101906102149190612c46565b6109e7565b005b34801561022757600080fd5b50610242600480360381019061023d9190612731565b610a7e565b005b34801561025057600080fd5b50610259610c07565b604051610266919061315f565b60405180910390f35b34801561027b57600080fd5b5061029660048036038101906102919190612a36565b610c1f565b005b6102b260048036038101906102ad919061275e565b610d13565b005b3480156102c057600080fd5b506102c9610e50565b6040516102d69190613396565b60405180910390f35b3480156102eb57600080fd5b506102f4610f09565b005b34801561030257600080fd5b5061031d60048036038101906103189190612992565b610f1d565b005b34801561032b57600080fd5b5061033461100a565b604051610341919061315f565b60405180910390f35b34801561035657600080fd5b50610371600480360381019061036c9190612a36565b611034565b005b34801561037f57600080fd5b5061039a60048036038101906103959190612aec565b611266565b005b3480156103a857600080fd5b506103c360048036038101906103be9190612aec565b61135a565b005b3480156103d157600080fd5b506103ec60048036038101906103e79190612731565b61158c565b005b3480156103fa57600080fd5b5061041560048036038101906104109190612731565b611736565b005b34801561042357600080fd5b5061043e600480360381019061043991906127fa565b6117ba565b005b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff167f2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f8484846040516104b0939291906133b1565b60405180910390a2505050565b60006104c98383611976565b90503373ffffffffffffffffffffffffffffffffffffffff167f2265ce9ec38ea098a1143406678482665a6e1ccd82ab22d37eea3a78abc57716838686858773ffffffffffffffffffffffffffffffffffffffff16634d8943bb6040518163ffffffff1660e01b815260040160206040518083038186803b15801561054d57600080fd5b505afa158015610561573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105859190612c73565b604051610596959493929190613300565b60405180910390a250505050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461061d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16148061069657503073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b156106cd576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330876040518463ffffffff1660e01b815260040161072c9392919061317a565b602060405180830381600087803b15801561074657600080fd5b505af115801561075a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077e919061284d565b6107b4576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d856040518263ffffffff1660e01b815260040161080f9190613631565b600060405180830381600087803b15801561082957600080fd5b505af115801561083d573d6000803e3d6000fd5b5050505060008373ffffffffffffffffffffffffffffffffffffffff16856040516108679061314a565b60006040518083038185875af1925050503d80600081146108a4576040519150601f19603f3d011682016040523d82523d6000602084013e6108a9565b606091505b505090508373ffffffffffffffffffffffffffffffffffffffff1663de43156e8760c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168887876040518663ffffffff1660e01b81526004016109109594939291906135dc565b600060405180830381600087803b15801561092a57600080fd5b505af115801561093e573d6000803e3d6000fd5b50505050505050505050565b61095383611c66565b3373ffffffffffffffffffffffffffffffffffffffff167f2265ce9ec38ea098a1143406678482665a6e1ccd82ab22d37eea3a78abc57716600073735b14bb79463307aacbed86daf3322b1e6226ab6040516020016109b29190613118565b6040516020818303038152906040528660008088886040516109da97969594939291906131b1565b60405180910390a2505050565b6109f081611c66565b3373ffffffffffffffffffffffffffffffffffffffff167f2265ce9ec38ea098a1143406678482665a6e1ccd82ab22d37eea3a78abc57716600073735b14bb79463307aacbed86daf3322b1e6226ab604051602001610a4f9190613118565b60405160208183030381529060405284600080604051610a73959493929190613222565b60405180910390a250565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610b0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0490613447565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610b4c611e95565b73ffffffffffffffffffffffffffffffffffffffff1614610ba2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9990613467565b60405180910390fd5b610bab81611eec565b610c0481600067ffffffffffffffff811115610bca57610bc96138a1565b5b6040519080825280601f01601f191660200182016040528015610bfc5781602001600182028036833780820191505090505b506000611ef7565b50565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c98576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166369582bee87878786866040518663ffffffff1660e01b8152600401610cd9959493929190613587565b600060405180830381600087803b158015610cf357600080fd5b505af1158015610d07573d6000803e3d6000fd5b50505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610da2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9990613447565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610de1611e95565b73ffffffffffffffffffffffffffffffffffffffff1614610e37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e2e90613467565b60405180910390fd5b610e4082611eec565b610e4c82826001611ef7565b5050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610ee0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed790613487565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b610f11612074565b610f1b60006120f2565b565b6000610f298585611976565b90503373ffffffffffffffffffffffffffffffffffffffff167f2265ce9ec38ea098a1143406678482665a6e1ccd82ab22d37eea3a78abc57716858888858973ffffffffffffffffffffffffffffffffffffffff16634d8943bb6040518163ffffffff1660e01b815260040160206040518083038186803b158015610fad57600080fd5b505afa158015610fc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fe59190612c73565b8989604051610ffa979695949392919061328f565b60405180910390a2505050505050565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146110ad576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16148061112657503073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b1561115d576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff166347e7ef2484866040518363ffffffff1660e01b815260040161119892919061336d565b602060405180830381600087803b1580156111b257600080fd5b505af11580156111c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ea919061284d565b508273ffffffffffffffffffffffffffffffffffffffff166369582bee87878786866040518663ffffffff1660e01b815260040161122c959493929190613587565b600060405180830381600087803b15801561124657600080fd5b505af115801561125a573d6000803e3d6000fd5b50505050505050505050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146112df576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1663de43156e87878786866040518663ffffffff1660e01b81526004016113209594939291906135dc565b600060405180830381600087803b15801561133a57600080fd5b505af115801561134e573d6000803e3d6000fd5b50505050505050505050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146113d3576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16148061144c57503073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15611483576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff166347e7ef2484866040518363ffffffff1660e01b81526004016114be92919061336d565b602060405180830381600087803b1580156114d857600080fd5b505af11580156114ec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611510919061284d565b508273ffffffffffffffffffffffffffffffffffffffff1663de43156e87878786866040518663ffffffff1660e01b81526004016115529594939291906135dc565b600060405180830381600087803b15801561156c57600080fd5b505af1158015611580573d6000803e3d6000fd5b50505050505050505050565b60008060019054906101000a900460ff161590508080156115bd5750600160008054906101000a900460ff1660ff16105b806115ea57506115cc306121b8565b1580156115e95750600160008054906101000a900460ff1660ff16145b5b611629576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611620906134c7565b60405180910390fd5b60016000806101000a81548160ff021916908360ff1602179055508015611666576001600060016101000a81548160ff0219169083151502179055505b61166e6121db565b611676612234565b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660c960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080156117325760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405161172991906133ea565b60405180910390a15b5050565b61173e612074565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156117ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117a590613427565b60405180910390fd5b6117b7816120f2565b50565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611833576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614806118ac57503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156118e3576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166347e7ef2482846040518363ffffffff1660e01b815260040161191e92919061336d565b602060405180830381600087803b15801561193857600080fd5b505af115801561194c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611970919061284d565b50505050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1663d9eeebed6040518163ffffffff1660e01b8152600401604080518083038186803b1580156119c057600080fd5b505afa1580156119d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119f891906127ba565b915091508173ffffffffffffffffffffffffffffffffffffffff166323b872dd3373735b14bb79463307aacbed86daf3322b1e6226ab846040518463ffffffff1660e01b8152600401611a4d9392919061317a565b602060405180830381600087803b158015611a6757600080fd5b505af1158015611a7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a9f919061284d565b611ad5576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330886040518463ffffffff1660e01b8152600401611b129392919061317a565b602060405180830381600087803b158015611b2c57600080fd5b505af1158015611b40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b64919061284d565b611b9a576040517f4dd9ee8d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff166342966c68866040518263ffffffff1660e01b8152600401611bd39190613631565b602060405180830381600087803b158015611bed57600080fd5b505af1158015611c01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c25919061284d565b611c5b576040517f2c77e05c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b809250505092915050565b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401611cc59392919061317a565b602060405180830381600087803b158015611cdf57600080fd5b505af1158015611cf3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d17919061284d565b611d4d576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d826040518263ffffffff1660e01b8152600401611da89190613631565b600060405180830381600087803b158015611dc257600080fd5b505af1158015611dd6573d6000803e3d6000fd5b50505050600073735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff1682604051611e149061314a565b60006040518083038185875af1925050503d8060008114611e51576040519150601f19603f3d011682016040523d82523d6000602084013e611e56565b606091505b5050905080611e91576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b6000611ec37f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b612285565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611ef4612074565b50565b611f237f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b61228f565b60000160009054906101000a900460ff1615611f4757611f4283612299565b61206f565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611f8d57600080fd5b505afa925050508015611fbe57506040513d601f19601f82011682018060405250810190611fbb919061287a565b60015b611ffd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ff4906134e7565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114612062576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612059906134a7565b60405180910390fd5b5061206e838383612352565b5b505050565b61207c61237e565b73ffffffffffffffffffffffffffffffffffffffff1661209a61100a565b73ffffffffffffffffffffffffffffffffffffffff16146120f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120e790613527565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff1661222a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161222190613567565b60405180910390fd5b612232612386565b565b600060019054906101000a900460ff16612283576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161227a90613567565b60405180910390fd5b565b6000819050919050565b6000819050919050565b6122a2816121b8565b6122e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122d890613507565b60405180910390fd5b8061230e7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b612285565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61235b836123e7565b6000825111806123685750805b15612379576123778383612436565b505b505050565b600033905090565b600060019054906101000a900460ff166123d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123cc90613567565b60405180910390fd5b6123e56123e061237e565b6120f2565b565b6123f081612299565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b606061245b8383604051806060016040528060278152602001613c9960279139612463565b905092915050565b60606000808573ffffffffffffffffffffffffffffffffffffffff168560405161248d9190613133565b600060405180830381855af49150503d80600081146124c8576040519150601f19603f3d011682016040523d82523d6000602084013e6124cd565b606091505b50915091506124de868383876124e9565b925050509392505050565b6060831561254c5760008351141561254457612504856121b8565b612543576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161253a90613547565b60405180910390fd5b5b829050612557565b612556838361255f565b5b949350505050565b6000825111156125725781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125a69190613405565b60405180910390fd5b60006125c26125bd84613671565b61364c565b9050828152602081018484840111156125de576125dd6138ee565b5b6125e984828561380a565b509392505050565b60008135905061260081613c3c565b92915050565b60008151905061261581613c3c565b92915050565b60008151905061262a81613c53565b92915050565b60008151905061263f81613c6a565b92915050565b60008083601f84011261265b5761265a6138da565b5b8235905067ffffffffffffffff811115612678576126776138d5565b5b602083019150836001820283011115612694576126936138e9565b5b9250929050565b600082601f8301126126b0576126af6138da565b5b81356126c08482602086016125af565b91505092915050565b6000606082840312156126df576126de6138df565b5b81905092915050565b6000606082840312156126fe576126fd6138df565b5b81905092915050565b60008135905061271681613c81565b92915050565b60008151905061272b81613c81565b92915050565b600060208284031215612747576127466138fd565b5b6000612755848285016125f1565b91505092915050565b60008060408385031215612775576127746138fd565b5b6000612783858286016125f1565b925050602083013567ffffffffffffffff8111156127a4576127a36138f3565b5b6127b08582860161269b565b9150509250929050565b600080604083850312156127d1576127d06138fd565b5b60006127df85828601612606565b92505060206127f08582860161271c565b9150509250929050565b600080600060608486031215612813576128126138fd565b5b6000612821868287016125f1565b935050602061283286828701612707565b9250506040612843868287016125f1565b9150509250925092565b600060208284031215612863576128626138fd565b5b60006128718482850161261b565b91505092915050565b6000602082840312156128905761288f6138fd565b5b600061289e84828501612630565b91505092915050565b6000806000604084860312156128c0576128bf6138fd565b5b600084013567ffffffffffffffff8111156128de576128dd6138f3565b5b6128ea8682870161269b565b935050602084013567ffffffffffffffff81111561290b5761290a6138f3565b5b61291786828701612645565b92509250509250925092565b60008060006060848603121561293c5761293b6138fd565b5b600084013567ffffffffffffffff81111561295a576129596138f3565b5b6129668682870161269b565b935050602061297786828701612707565b9250506040612988868287016125f1565b9150509250925092565b6000806000806000608086880312156129ae576129ad6138fd565b5b600086013567ffffffffffffffff8111156129cc576129cb6138f3565b5b6129d88882890161269b565b95505060206129e988828901612707565b94505060406129fa888289016125f1565b935050606086013567ffffffffffffffff811115612a1b57612a1a6138f3565b5b612a2788828901612645565b92509250509295509295909350565b60008060008060008060a08789031215612a5357612a526138fd565b5b600087013567ffffffffffffffff811115612a7157612a706138f3565b5b612a7d89828a016126c9565b9650506020612a8e89828a016125f1565b9550506040612a9f89828a01612707565b9450506060612ab089828a016125f1565b935050608087013567ffffffffffffffff811115612ad157612ad06138f3565b5b612add89828a01612645565b92509250509295509295509295565b60008060008060008060a08789031215612b0957612b086138fd565b5b600087013567ffffffffffffffff811115612b2757612b266138f3565b5b612b3389828a016126e8565b9650506020612b4489828a016125f1565b9550506040612b5589828a01612707565b9450506060612b6689828a016125f1565b935050608087013567ffffffffffffffff811115612b8757612b866138f3565b5b612b9389828a01612645565b92509250509295509295509295565b600080600080600060808688031215612bbe57612bbd6138fd565b5b600086013567ffffffffffffffff811115612bdc57612bdb6138f3565b5b612be8888289016126e8565b9550506020612bf988828901612707565b9450506040612c0a888289016125f1565b935050606086013567ffffffffffffffff811115612c2b57612c2a6138f3565b5b612c3788828901612645565b92509250509295509295909350565b600060208284031215612c5c57612c5b6138fd565b5b6000612c6a84828501612707565b91505092915050565b600060208284031215612c8957612c886138fd565b5b6000612c978482850161271c565b91505092915050565b600080600060408486031215612cb957612cb86138fd565b5b6000612cc786828701612707565b935050602084013567ffffffffffffffff811115612ce857612ce76138f3565b5b612cf486828701612645565b92509250509250925092565b612d0981613787565b82525050565b612d1881613787565b82525050565b612d2f612d2a82613787565b61387d565b82525050565b612d3e816137a5565b82525050565b6000612d5083856136b8565b9350612d5d83858461380a565b612d6683613902565b840190509392505050565b6000612d7d83856136c9565b9350612d8a83858461380a565b612d9383613902565b840190509392505050565b6000612da9826136a2565b612db381856136c9565b9350612dc3818560208601613819565b612dcc81613902565b840191505092915050565b6000612de2826136a2565b612dec81856136da565b9350612dfc818560208601613819565b80840191505092915050565b612e11816137e6565b82525050565b612e20816137f8565b82525050565b6000612e31826136ad565b612e3b81856136e5565b9350612e4b818560208601613819565b612e5481613902565b840191505092915050565b6000612e6c6026836136e5565b9150612e7782613920565b604082019050919050565b6000612e8f602c836136e5565b9150612e9a8261396f565b604082019050919050565b6000612eb2602c836136e5565b9150612ebd826139be565b604082019050919050565b6000612ed56038836136e5565b9150612ee082613a0d565b604082019050919050565b6000612ef86029836136e5565b9150612f0382613a5c565b604082019050919050565b6000612f1b602e836136e5565b9150612f2682613aab565b604082019050919050565b6000612f3e602e836136e5565b9150612f4982613afa565b604082019050919050565b6000612f61602d836136e5565b9150612f6c82613b49565b604082019050919050565b6000612f846020836136e5565b9150612f8f82613b98565b602082019050919050565b6000612fa76000836136c9565b9150612fb282613bc1565b600082019050919050565b6000612fca6000836136da565b9150612fd582613bc1565b600082019050919050565b6000612fed601d836136e5565b9150612ff882613bc4565b602082019050919050565b6000613010602b836136e5565b915061301b82613bed565b604082019050919050565b600060608301613039600084018461370d565b858303600087015261304c838284612d44565b9250505061305d60208401846136f6565b61306a6020860182612d00565b506130786040840184613770565b61308560408601826130fa565b508091505092915050565b6000606083016130a3600084018461370d565b85830360008701526130b6838284612d44565b925050506130c760208401846136f6565b6130d46020860182612d00565b506130e26040840184613770565b6130ef60408601826130fa565b508091505092915050565b613103816137cf565b82525050565b613112816137cf565b82525050565b60006131248284612d1e565b60148201915081905092915050565b600061313f8284612dd7565b915081905092915050565b600061315582612fbd565b9150819050919050565b60006020820190506131746000830184612d0f565b92915050565b600060608201905061318f6000830186612d0f565b61319c6020830185612d0f565b6131a96040830184613109565b949350505050565b600060c0820190506131c6600083018a612d0f565b81810360208301526131d88189612d9e565b90506131e76040830188613109565b6131f46060830187612e08565b6132016080830186612e08565b81810360a0830152613214818486612d71565b905098975050505050505050565b600060c0820190506132376000830188612d0f565b81810360208301526132498187612d9e565b90506132586040830186613109565b6132656060830185612e08565b6132726080830184612e08565b81810360a083015261328381612f9a565b90509695505050505050565b600060c0820190506132a4600083018a612d0f565b81810360208301526132b68189612d9e565b90506132c56040830188613109565b6132d26060830187613109565b6132df6080830186613109565b81810360a08301526132f2818486612d71565b905098975050505050505050565b600060c0820190506133156000830188612d0f565b81810360208301526133278187612d9e565b90506133366040830186613109565b6133436060830185613109565b6133506080830184613109565b81810360a083015261336181612f9a565b90509695505050505050565b60006040820190506133826000830185612d0f565b61338f6020830184613109565b9392505050565b60006020820190506133ab6000830184612d35565b92915050565b600060408201905081810360008301526133cb8186612d9e565b905081810360208301526133e0818486612d71565b9050949350505050565b60006020820190506133ff6000830184612e17565b92915050565b6000602082019050818103600083015261341f8184612e26565b905092915050565b6000602082019050818103600083015261344081612e5f565b9050919050565b6000602082019050818103600083015261346081612e82565b9050919050565b6000602082019050818103600083015261348081612ea5565b9050919050565b600060208201905081810360008301526134a081612ec8565b9050919050565b600060208201905081810360008301526134c081612eeb565b9050919050565b600060208201905081810360008301526134e081612f0e565b9050919050565b6000602082019050818103600083015261350081612f31565b9050919050565b6000602082019050818103600083015261352081612f54565b9050919050565b6000602082019050818103600083015261354081612f77565b9050919050565b6000602082019050818103600083015261356081612fe0565b9050919050565b6000602082019050818103600083015261358081613003565b9050919050565b600060808201905081810360008301526135a18188613026565b90506135b06020830187612d0f565b6135bd6040830186613109565b81810360608301526135d0818486612d71565b90509695505050505050565b600060808201905081810360008301526135f68188613090565b90506136056020830187612d0f565b6136126040830186613109565b8181036060830152613625818486612d71565b90509695505050505050565b60006020820190506136466000830184613109565b92915050565b6000613656613667565b9050613662828261384c565b919050565b6000604051905090565b600067ffffffffffffffff82111561368c5761368b6138a1565b5b61369582613902565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061370560208401846125f1565b905092915050565b6000808335600160200384360303811261372a576137296138f8565b5b83810192508235915060208301925067ffffffffffffffff821115613752576137516138d0565b5b600182023603841315613768576137676138e4565b5b509250929050565b600061377f6020840184612707565b905092915050565b6000613792826137af565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006137f1826137cf565b9050919050565b6000613803826137d9565b9050919050565b82818337600083830152505050565b60005b8381101561383757808201518184015260208101905061381c565b83811115613846576000848401525b50505050565b61385582613902565b810181811067ffffffffffffffff82111715613874576138736138a1565b5b80604052505050565b60006138888261388f565b9050919050565b600061389a82613913565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b613c4581613787565b8114613c5057600080fd5b50565b613c5c81613799565b8114613c6757600080fd5b50565b613c73816137a5565b8114613c7e57600080fd5b50565b613c8a816137cf565b8114613c9557600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212208488558616000861895afc52cc0d318ef6747a1009f23a98893e448920be93df64736f6c63430008070033608060405234801561001057600080fd5b50604051610bcd380380610bcd8339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610108565b600081519050610087816100f1565b92915050565b6000602082840312156100a3576100a26100ec565b5b60006100b184828501610078565b91505092915050565b60006100c5826100cc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6100fa816100ba565b811461010557600080fd5b50565b610ab6806101176000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b614610062578063a0a1730b14610080575b600080fd5b610060600480360381019061005b91906105fd565b61009c565b005b61006a6102af565b6040516100779190610761565b60405180910390f35b61009a6004803603810190610095919061055e565b6102d3565b005b60008383836040516024016100b39392919061082f565b6040516020818303038152906040527fe04d4f97000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090508473ffffffffffffffffffffffffffffffffffffffff1663095ea7b360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16886040518363ffffffff1660e01b815260040161018d92919061077c565b602060405180830381600087803b1580156101a757600080fd5b505af11580156101bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101df9190610531565b610215576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637993c1e0888888856040518563ffffffff1660e01b815260040161027494939291906107dc565b600060405180830381600087803b15801561028e57600080fd5b505af11580156102a2573d6000803e3d6000fd5b5050505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008383836040516024016102ea9392919061082f565b6040516020818303038152906040527fe04d4f97000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630ac7c44c86836040518363ffffffff1660e01b81526004016103c49291906107a5565b600060405180830381600087803b1580156103de57600080fd5b505af11580156103f2573d6000803e3d6000fd5b505050505050505050565b600061041061040b84610892565b61086d565b90508281526020810184848401111561042c5761042b610a1b565b5b610437848285610974565b509392505050565b600061045261044d846108c3565b61086d565b90508281526020810184848401111561046e5761046d610a1b565b5b610479848285610974565b509392505050565b60008135905061049081610a3b565b92915050565b6000813590506104a581610a52565b92915050565b6000815190506104ba81610a52565b92915050565b600082601f8301126104d5576104d4610a16565b5b81356104e58482602086016103fd565b91505092915050565b600082601f83011261050357610502610a16565b5b813561051384826020860161043f565b91505092915050565b60008135905061052b81610a69565b92915050565b60006020828403121561054757610546610a25565b5b6000610555848285016104ab565b91505092915050565b6000806000806080858703121561057857610577610a25565b5b600085013567ffffffffffffffff81111561059657610595610a20565b5b6105a2878288016104c0565b945050602085013567ffffffffffffffff8111156105c3576105c2610a20565b5b6105cf878288016104ee565b93505060406105e08782880161051c565b92505060606105f187828801610496565b91505092959194509250565b60008060008060008060c0878903121561061a57610619610a25565b5b600087013567ffffffffffffffff81111561063857610637610a20565b5b61064489828a016104c0565b965050602061065589828a0161051c565b955050604061066689828a01610481565b945050606087013567ffffffffffffffff81111561068757610686610a20565b5b61069389828a016104ee565b93505060806106a489828a0161051c565b92505060a06106b589828a01610496565b9150509295509295509295565b6106cb8161092c565b82525050565b6106da8161093e565b82525050565b60006106eb826108f4565b6106f5818561090a565b9350610705818560208601610983565b61070e81610a2a565b840191505092915050565b6000610724826108ff565b61072e818561091b565b935061073e818560208601610983565b61074781610a2a565b840191505092915050565b61075b8161096a565b82525050565b600060208201905061077660008301846106c2565b92915050565b600060408201905061079160008301856106c2565b61079e6020830184610752565b9392505050565b600060408201905081810360008301526107bf81856106e0565b905081810360208301526107d381846106e0565b90509392505050565b600060808201905081810360008301526107f681876106e0565b90506108056020830186610752565b61081260408301856106c2565b818103606083015261082481846106e0565b905095945050505050565b600060608201905081810360008301526108498186610719565b90506108586020830185610752565b61086560408301846106d1565b949350505050565b6000610877610888565b905061088382826109b6565b919050565b6000604051905090565b600067ffffffffffffffff8211156108ad576108ac6109e7565b5b6108b682610a2a565b9050602081019050919050565b600067ffffffffffffffff8211156108de576108dd6109e7565b5b6108e782610a2a565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60006109378261094a565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156109a1578082015181840152602081019050610986565b838111156109b0576000848401525b50505050565b6109bf82610a2a565b810181811067ffffffffffffffff821117156109de576109dd6109e7565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b610a448161092c565b8114610a4f57600080fd5b50565b610a5b8161093e565b8114610a6657600080fd5b50565b610a728161096a565b8114610a7d57600080fd5b5056fea264697066735822122071a41b7e8232f98e8257a24fd602c7cc7b7cd2eab6d352fb5c5fbe1412f3398364736f6c6343000807003360806040523480156200001157600080fd5b50604051620010d7380380620010d7833981810160405281019062000037919062000146565b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac560405160405180910390a1505050620001f5565b6000815190506200014081620001db565b92915050565b600080600060608486031215620001625762000161620001d6565b5b600062000172868287016200012f565b935050602062000185868287016200012f565b925050604062000198868287016200012f565b9150509250925092565b6000620001af82620001b6565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620001e681620001a2565b8114620001f257600080fd5b50565b610ed280620002056000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806397770dff1161007157806397770dff14610166578063a7cb050714610182578063c63585cc1461019e578063d7fd7afb146101ce578063d936a012146101fe578063ee2815ba1461021c576100a9565b80630be15547146100ae5780633c669d55146100de578063513a9c05146100fa578063569541b91461012a578063842da36d14610148575b600080fd5b6100c860048036038101906100c3919061094d565b610238565b6040516100d59190610bce565b60405180910390f35b6100f860048036038101906100f39190610898565b61026b565b005b610114600480360381019061010f919061094d565b6103b8565b6040516101219190610bce565b60405180910390f35b6101326103eb565b60405161013f9190610bce565b60405180910390f35b610150610411565b60405161015d9190610bce565b60405180910390f35b610180600480360381019061017b9190610818565b610437565b005b61019c600480360381019061019791906109ba565b6104d4565b005b6101b860048036038101906101b39190610845565b610528565b6040516101c59190610bce565b60405180910390f35b6101e860048036038101906101e3919061094d565b61059a565b6040516101f59190610c67565b60405180910390f35b6102066105b2565b6040516102139190610bce565b60405180910390f35b6102366004803603810190610231919061097a565b6105d8565b005b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060405180606001604052806040518060200160405280600081525081526020013373ffffffffffffffffffffffffffffffffffffffff1681526020014681525090508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb87866040518363ffffffff1660e01b81526004016102ea929190610be9565b602060405180830381600087803b15801561030457600080fd5b505af1158015610318573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061033c9190610920565b508573ffffffffffffffffffffffffffffffffffffffff1663de43156e82878787876040518663ffffffff1660e01b815260040161037e959493929190610c12565b600060405180830381600087803b15801561039857600080fd5b505af11580156103ac573d6000803e3d6000fd5b50505050505050505050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516104c99190610bce565b60405180910390a150565b80600080848152602001908152602001600020819055507f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d828260405161051c929190610cab565b60405180910390a15050565b60008060006105378585610667565b9150915085828260405160200161054f929190610b60565b60405160208183030381529060405280519060200120604051602001610576929190610b8c565b6040516020818303038152906040528051906020012060001c925050509392505050565b60006020528060005260406000206000915090505481565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d828260405161065b929190610c82565b60405180910390a15050565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156106d0576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161061070a57828461070d565b83835b8092508193505050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561077c576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b60008135905061079281610e57565b92915050565b6000815190506107a781610e6e565b92915050565b60008083601f8401126107c3576107c2610dd3565b5b8235905067ffffffffffffffff8111156107e0576107df610dce565b5b6020830191508360018202830111156107fc576107fb610dd8565b5b9250929050565b60008135905061081281610e85565b92915050565b60006020828403121561082e5761082d610de2565b5b600061083c84828501610783565b91505092915050565b60008060006060848603121561085e5761085d610de2565b5b600061086c86828701610783565b935050602061087d86828701610783565b925050604061088e86828701610783565b9150509250925092565b6000806000806000608086880312156108b4576108b3610de2565b5b60006108c288828901610783565b95505060206108d388828901610783565b94505060406108e488828901610803565b935050606086013567ffffffffffffffff81111561090557610904610ddd565b5b610911888289016107ad565b92509250509295509295909350565b60006020828403121561093657610935610de2565b5b600061094484828501610798565b91505092915050565b60006020828403121561096357610962610de2565b5b600061097184828501610803565b91505092915050565b6000806040838503121561099157610990610de2565b5b600061099f85828601610803565b92505060206109b085828601610783565b9150509250929050565b600080604083850312156109d1576109d0610de2565b5b60006109df85828601610803565b92505060206109f085828601610803565b9150509250929050565b610a0381610d0c565b82525050565b610a1281610d0c565b82525050565b610a29610a2482610d0c565b610da0565b82525050565b610a40610a3b82610d2a565b610db2565b82525050565b6000610a528385610cf0565b9350610a5f838584610d5e565b610a6883610de7565b840190509392505050565b6000610a7e82610cd4565b610a888185610cdf565b9350610a98818560208601610d6d565b610aa181610de7565b840191505092915050565b6000610ab9602083610d01565b9150610ac482610e05565b602082019050919050565b6000610adc600183610d01565b9150610ae782610e2e565b600182019050919050565b60006060830160008301518482036000860152610b0f8282610a73565b9150506020830151610b2460208601826109fa565b506040830151610b376040860182610b42565b508091505092915050565b610b4b81610d54565b82525050565b610b5a81610d54565b82525050565b6000610b6c8285610a18565b601482019150610b7c8284610a18565b6014820191508190509392505050565b6000610b9782610acf565b9150610ba38285610a18565b601482019150610bb38284610a2f565b602082019150610bc282610aac565b91508190509392505050565b6000602082019050610be36000830184610a09565b92915050565b6000604082019050610bfe6000830185610a09565b610c0b6020830184610b51565b9392505050565b60006080820190508181036000830152610c2c8188610af2565b9050610c3b6020830187610a09565b610c486040830186610b51565b8181036060830152610c5b818486610a46565b90509695505050505050565b6000602082019050610c7c6000830184610b51565b92915050565b6000604082019050610c976000830185610b51565b610ca46020830184610a09565b9392505050565b6000604082019050610cc06000830185610b51565b610ccd6020830184610b51565b9392505050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000610d1782610d34565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610d8b578082015181840152602081019050610d70565b83811115610d9a576000848401525b50505050565b6000610dab82610dbc565b9050919050565b6000819050919050565b6000610dc782610df8565b9050919050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f600082015250565b7fff00000000000000000000000000000000000000000000000000000000000000600082015250565b610e6081610d0c565b8114610e6b57600080fd5b50565b610e7781610d1e565b8114610e8257600080fd5b50565b610e8e81610d54565b8114610e9957600080fd5b5056fea26469706673582212200835245fdc75aa593111cf5a43e919740b2258e1512040bc79f39a56cedf119364736f6c6343000807003360c06040523480156200001157600080fd5b506040516200282d3803806200282d83398181016040528101906200003791906200035b565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620000b1576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8760079080519060200190620000c9929190620001d1565b508660089080519060200190620000e2929190620001d1565b5085600960006101000a81548160ff021916908360ff16021790555084608081815250508360028111156200011c576200011b620005ae565b5b60a0816002811115620001345762000133620005ae565b5b60f81b8152505082600281905550816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050505050620006bf565b828054620001df9062000542565b90600052602060002090601f0160209004810192826200020357600085556200024f565b82601f106200021e57805160ff19168380011785556200024f565b828001600101855582156200024f579182015b828111156200024e57825182559160200191906001019062000231565b5b5090506200025e919062000262565b5090565b5b808211156200027d57600081600090555060010162000263565b5090565b60006200029862000292846200048b565b62000462565b905082815260208101848484011115620002b757620002b662000640565b5b620002c48482856200050c565b509392505050565b600081519050620002dd8162000660565b92915050565b600081519050620002f4816200067a565b92915050565b600082601f8301126200031257620003116200063b565b5b81516200032484826020860162000281565b91505092915050565b6000815190506200033e816200068b565b92915050565b6000815190506200035581620006a5565b92915050565b600080600080600080600080610100898b0312156200037f576200037e6200064a565b5b600089015167ffffffffffffffff811115620003a0576200039f62000645565b5b620003ae8b828c01620002fa565b985050602089015167ffffffffffffffff811115620003d257620003d162000645565b5b620003e08b828c01620002fa565b9750506040620003f38b828c0162000344565b9650506060620004068b828c016200032d565b9550506080620004198b828c01620002e3565b94505060a06200042c8b828c016200032d565b93505060c06200043f8b828c01620002cc565b92505060e0620004528b828c01620002cc565b9150509295985092959890939650565b60006200046e62000481565b90506200047c828262000578565b919050565b6000604051905090565b600067ffffffffffffffff821115620004a957620004a86200060c565b5b620004b4826200064f565b9050602081019050919050565b6000620004ce82620004d5565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b838110156200052c5780820151818401526020810190506200050f565b838111156200053c576000848401525b50505050565b600060028204905060018216806200055b57607f821691505b60208210811415620005725762000571620005dd565b5b50919050565b62000583826200064f565b810181811067ffffffffffffffff82111715620005a557620005a46200060c565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6200066b81620004c1565b81146200067757600080fd5b50565b600381106200068857600080fd5b50565b6200069681620004f5565b8114620006a257600080fd5b50565b620006b081620004ff565b8114620006bc57600080fd5b50565b60805160a05160f81c612137620006f660003960006109580152600081816108a201528181610c250152610d5a01526121376000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806385e1f4d0116100c3578063d9eeebed1161007c578063d9eeebed146103cc578063dd62ed3e146103eb578063e2f535b81461041b578063eddeb12314610439578063f2441b3214610455578063f687d12a146104735761014d565b806385e1f4d0146102f657806395d89b4114610314578063a3413d0314610332578063a9059cbb14610350578063c701262614610380578063c835d7cc146103b05761014d565b8063313ce56711610115578063313ce5671461020c5780633ce4a5bc1461022a57806342966c681461024857806347e7ef24146102785780634d8943bb146102a857806370a08231146102c65761014d565b806306fdde0314610152578063091d278814610170578063095ea7b31461018e57806318160ddd146101be57806323b872dd146101dc575b600080fd5b61015a61048f565b6040516101679190611cad565b60405180910390f35b610178610521565b6040516101859190611ccf565b60405180910390f35b6101a860048036038101906101a3919061196e565b610527565b6040516101b59190611bfb565b60405180910390f35b6101c6610545565b6040516101d39190611ccf565b60405180910390f35b6101f660048036038101906101f1919061191b565b61054f565b6040516102039190611bfb565b60405180910390f35b610214610647565b6040516102219190611cea565b60405180910390f35b61023261065e565b60405161023f9190611b80565b60405180910390f35b610262600480360381019061025d9190611a37565b610676565b60405161026f9190611bfb565b60405180910390f35b610292600480360381019061028d919061196e565b61068b565b60405161029f9190611bfb565b60405180910390f35b6102b0610851565b6040516102bd9190611ccf565b60405180910390f35b6102e060048036038101906102db9190611881565b610857565b6040516102ed9190611ccf565b60405180910390f35b6102fe6108a0565b60405161030b9190611ccf565b60405180910390f35b61031c6108c4565b6040516103299190611cad565b60405180910390f35b61033a610956565b6040516103479190611c92565b60405180910390f35b61036a6004803603810190610365919061196e565b61097a565b6040516103779190611bfb565b60405180910390f35b61039a600480360381019061039591906119db565b610998565b6040516103a79190611bfb565b60405180910390f35b6103ca60048036038101906103c59190611881565b610aee565b005b6103d4610be1565b6040516103e2929190611bd2565b60405180910390f35b610405600480360381019061040091906118db565b610e4e565b6040516104129190611ccf565b60405180910390f35b610423610ed5565b6040516104309190611b80565b60405180910390f35b610453600480360381019061044e9190611a37565b610efb565b005b61045d610fb5565b60405161046a9190611b80565b60405180910390f35b61048d60048036038101906104889190611a37565b610fd9565b005b60606007805461049e90611f33565b80601f01602080910402602001604051908101604052809291908181526020018280546104ca90611f33565b80156105175780601f106104ec57610100808354040283529160200191610517565b820191906000526020600020905b8154815290600101906020018083116104fa57829003601f168201915b5050505050905090565b60025481565b600061053b610534611093565b848461109b565b6001905092915050565b6000600654905090565b600061055c848484611254565b6000600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006105a7611093565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508281101561061e576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61063b8561062a611093565b85846106369190611e43565b61109b565b60019150509392505050565b6000600960009054906101000a900460ff16905090565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b600061068233836114b0565b60019050919050565b600073735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610729575060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b80156107835750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b156107ba576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107c48383611668565b8273ffffffffffffffffffffffffffffffffffffffff167f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab373735b14bb79463307aacbed86daf3322b1e6226ab6040516020016108219190611b65565b6040516020818303038152906040528460405161083f929190611c16565b60405180910390a26001905092915050565b60035481565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6060600880546108d390611f33565b80601f01602080910402602001604051908101604052809291908181526020018280546108ff90611f33565b801561094c5780601f106109215761010080835404028352916020019161094c565b820191906000526020600020905b81548152906001019060200180831161092f57829003601f168201915b5050505050905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b600061098e610987611093565b8484611254565b6001905092915050565b60008060006109a5610be1565b915091508173ffffffffffffffffffffffffffffffffffffffff166323b872dd3373735b14bb79463307aacbed86daf3322b1e6226ab846040518463ffffffff1660e01b81526004016109fa93929190611b9b565b602060405180830381600087803b158015610a1457600080fd5b505af1158015610a28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4c91906119ae565b610a82576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a8c33856114b0565b3373ffffffffffffffffffffffffffffffffffffffff167f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955868684600354604051610ada9493929190611c46565b60405180910390a260019250505092915050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b67576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae81604051610bd69190611b80565b60405180910390a150565b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630be155477f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b8152600401610c609190611ccf565b60206040518083038186803b158015610c7857600080fd5b505afa158015610c8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb091906118ae565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610d19576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d7fd7afb7f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b8152600401610d959190611ccf565b60206040518083038186803b158015610dad57600080fd5b505afa158015610dc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610de59190611a64565b90506000811415610e22576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060035460025483610e359190611de9565b610e3f9190611d93565b90508281945094505050509091565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f74576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806003819055507fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f81604051610faa9190611ccf565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611052576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806002819055507fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a816040516110889190611ccf565b60405180910390a150565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611102576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611169576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516112479190611ccf565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156112bb576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611322576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156113a0576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81816113ac9190611e43565b600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461143e9190611d93565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516114a29190611ccf565b60405180910390a350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611517576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611595576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81816115a19190611e43565b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600660008282546115f69190611e43565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161165b9190611ccf565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156116cf576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600660008282546116e19190611d93565b9250508190555080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546117379190611d93565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161179c9190611ccf565b60405180910390a35050565b60006117bb6117b684611d2a565b611d05565b9050828152602081018484840111156117d7576117d661207b565b5b6117e2848285611ef1565b509392505050565b6000813590506117f9816120bc565b92915050565b60008151905061180e816120bc565b92915050565b600081519050611823816120d3565b92915050565b600082601f83011261183e5761183d612076565b5b813561184e8482602086016117a8565b91505092915050565b600081359050611866816120ea565b92915050565b60008151905061187b816120ea565b92915050565b60006020828403121561189757611896612085565b5b60006118a5848285016117ea565b91505092915050565b6000602082840312156118c4576118c3612085565b5b60006118d2848285016117ff565b91505092915050565b600080604083850312156118f2576118f1612085565b5b6000611900858286016117ea565b9250506020611911858286016117ea565b9150509250929050565b60008060006060848603121561193457611933612085565b5b6000611942868287016117ea565b9350506020611953868287016117ea565b925050604061196486828701611857565b9150509250925092565b6000806040838503121561198557611984612085565b5b6000611993858286016117ea565b92505060206119a485828601611857565b9150509250929050565b6000602082840312156119c4576119c3612085565b5b60006119d284828501611814565b91505092915050565b600080604083850312156119f2576119f1612085565b5b600083013567ffffffffffffffff811115611a1057611a0f612080565b5b611a1c85828601611829565b9250506020611a2d85828601611857565b9150509250929050565b600060208284031215611a4d57611a4c612085565b5b6000611a5b84828501611857565b91505092915050565b600060208284031215611a7a57611a79612085565b5b6000611a888482850161186c565b91505092915050565b611a9a81611e77565b82525050565b611ab1611aac82611e77565b611f96565b82525050565b611ac081611e89565b82525050565b6000611ad182611d5b565b611adb8185611d71565b9350611aeb818560208601611f00565b611af48161208a565b840191505092915050565b611b0881611edf565b82525050565b6000611b1982611d66565b611b238185611d82565b9350611b33818560208601611f00565b611b3c8161208a565b840191505092915050565b611b5081611ec8565b82525050565b611b5f81611ed2565b82525050565b6000611b718284611aa0565b60148201915081905092915050565b6000602082019050611b956000830184611a91565b92915050565b6000606082019050611bb06000830186611a91565b611bbd6020830185611a91565b611bca6040830184611b47565b949350505050565b6000604082019050611be76000830185611a91565b611bf46020830184611b47565b9392505050565b6000602082019050611c106000830184611ab7565b92915050565b60006040820190508181036000830152611c308185611ac6565b9050611c3f6020830184611b47565b9392505050565b60006080820190508181036000830152611c608187611ac6565b9050611c6f6020830186611b47565b611c7c6040830185611b47565b611c896060830184611b47565b95945050505050565b6000602082019050611ca76000830184611aff565b92915050565b60006020820190508181036000830152611cc78184611b0e565b905092915050565b6000602082019050611ce46000830184611b47565b92915050565b6000602082019050611cff6000830184611b56565b92915050565b6000611d0f611d20565b9050611d1b8282611f65565b919050565b6000604051905090565b600067ffffffffffffffff821115611d4557611d44612047565b5b611d4e8261208a565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611d9e82611ec8565b9150611da983611ec8565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611dde57611ddd611fba565b5b828201905092915050565b6000611df482611ec8565b9150611dff83611ec8565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611e3857611e37611fba565b5b828202905092915050565b6000611e4e82611ec8565b9150611e5983611ec8565b925082821015611e6c57611e6b611fba565b5b828203905092915050565b6000611e8282611ea8565b9050919050565b60008115159050919050565b6000819050611ea3826120a8565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611eea82611e95565b9050919050565b82818337600083830152505050565b60005b83811015611f1e578082015181840152602081019050611f03565b83811115611f2d576000848401525b50505050565b60006002820490506001821680611f4b57607f821691505b60208210811415611f5f57611f5e612018565b5b50919050565b611f6e8261208a565b810181811067ffffffffffffffff82111715611f8d57611f8c612047565b5b80604052505050565b6000611fa182611fa8565b9050919050565b6000611fb38261209b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b600381106120b9576120b8611fe9565b5b50565b6120c581611e77565b81146120d057600080fd5b50565b6120dc81611e89565b81146120e757600080fd5b50565b6120f381611ec8565b81146120fe57600080fd5b5056fea26469706673582212206982505c1a546edfb86a1aaaca0ad6b7e5542f2d1f24ac30a032805fc80a650e64736f6c63430008070033a2646970667358221220e93c95294a467be485feaaf121cb3adc54ad9c200d11fca8c43f1f820754cec364736f6c63430008070033",
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

// GatewayEVMZEVMTestCallIterator is returned from FilterCall and is used to iterate over the raw logs and unpacked data for Call events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestCallIterator struct {
	Event *GatewayEVMZEVMTestCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestCall represents a Call event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestCall struct {
	Sender   common.Address
	Receiver []byte
	Message  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCall is a free log retrieval operation binding the contract event 0x2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f.
//
// Solidity: event Call(address indexed sender, bytes receiver, bytes message)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterCall(opts *bind.FilterOpts, sender []common.Address) (*GatewayEVMZEVMTestCallIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Call", senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestCallIterator{contract: _GatewayEVMZEVMTest.contract, event: "Call", logs: logs, sub: sub}, nil
}

// WatchCall is a free log subscription operation binding the contract event 0x2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f.
//
// Solidity: event Call(address indexed sender, bytes receiver, bytes message)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchCall(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestCall, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Call", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestCall)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Call", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCall is a log parse operation binding the contract event 0x2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f.
//
// Solidity: event Call(address indexed sender, bytes receiver, bytes message)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseCall(log types.Log) (*GatewayEVMZEVMTestCall, error) {
	event := new(GatewayEVMZEVMTestCall)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Call", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestCall0Iterator is returned from FilterCall0 and is used to iterate over the raw logs and unpacked data for Call0 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestCall0Iterator struct {
	Event *GatewayEVMZEVMTestCall0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestCall0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestCall0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestCall0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestCall0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestCall0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestCall0 represents a Call0 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestCall0 struct {
	Sender   common.Address
	Receiver common.Address
	Payload  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCall0 is a free log retrieval operation binding the contract event 0x2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3.
//
// Solidity: event Call(address indexed sender, address indexed receiver, bytes payload)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterCall0(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMZEVMTestCall0Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Call0", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestCall0Iterator{contract: _GatewayEVMZEVMTest.contract, event: "Call0", logs: logs, sub: sub}, nil
}

// WatchCall0 is a free log subscription operation binding the contract event 0x2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3.
//
// Solidity: event Call(address indexed sender, address indexed receiver, bytes payload)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchCall0(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestCall0, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Call0", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestCall0)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Call0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCall0 is a log parse operation binding the contract event 0x2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3.
//
// Solidity: event Call(address indexed sender, address indexed receiver, bytes payload)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseCall0(log types.Log) (*GatewayEVMZEVMTestCall0, error) {
	event := new(GatewayEVMZEVMTestCall0)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Call0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestDepositIterator struct {
	Event *GatewayEVMZEVMTestDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestDeposit represents a Deposit event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestDeposit struct {
	Sender   common.Address
	Receiver common.Address
	Amount   *big.Int
	Asset    common.Address
	Payload  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4.
//
// Solidity: event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMZEVMTestDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Deposit", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestDepositIterator{contract: _GatewayEVMZEVMTest.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4.
//
// Solidity: event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestDeposit, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Deposit", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestDeposit)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4.
//
// Solidity: event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseDeposit(log types.Log) (*GatewayEVMZEVMTestDeposit, error) {
	event := new(GatewayEVMZEVMTestDeposit)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Deposit", log); err != nil {
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
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0xd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c.
//
// Solidity: event Reverted(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReverted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMZEVMTestRevertedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Reverted", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestRevertedIterator{contract: _GatewayEVMZEVMTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c.
//
// Solidity: event Reverted(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReverted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Reverted", destinationRule)
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

// ParseReverted is a log parse operation binding the contract event 0xd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c.
//
// Solidity: event Reverted(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReverted(log types.Log) (*GatewayEVMZEVMTestReverted, error) {
	event := new(GatewayEVMZEVMTestReverted)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestRevertedWithERC20Iterator is returned from FilterRevertedWithERC20 and is used to iterate over the raw logs and unpacked data for RevertedWithERC20 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestRevertedWithERC20Iterator struct {
	Event *GatewayEVMZEVMTestRevertedWithERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestRevertedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestRevertedWithERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestRevertedWithERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestRevertedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestRevertedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestRevertedWithERC20 represents a RevertedWithERC20 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestRevertedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevertedWithERC20 is a free log retrieval operation binding the contract event 0x723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda7.
//
// Solidity: event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterRevertedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMZEVMTestRevertedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "RevertedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestRevertedWithERC20Iterator{contract: _GatewayEVMZEVMTest.contract, event: "RevertedWithERC20", logs: logs, sub: sub}, nil
}

// WatchRevertedWithERC20 is a free log subscription operation binding the contract event 0x723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda7.
//
// Solidity: event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchRevertedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestRevertedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "RevertedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestRevertedWithERC20)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "RevertedWithERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevertedWithERC20 is a log parse operation binding the contract event 0x723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda7.
//
// Solidity: event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseRevertedWithERC20(log types.Log) (*GatewayEVMZEVMTestRevertedWithERC20, error) {
	event := new(GatewayEVMZEVMTestRevertedWithERC20)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "RevertedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestWithdrawalIterator struct {
	Event *GatewayEVMZEVMTestWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMZEVMTestWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMZEVMTestWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMZEVMTestWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestWithdrawal represents a Withdrawal event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestWithdrawal struct {
	From            common.Address
	Zrc20           common.Address
	To              []byte
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x2265ce9ec38ea098a1143406678482665a6e1ccd82ab22d37eea3a78abc57716.
//
// Solidity: event Withdrawal(address indexed from, address zrc20, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterWithdrawal(opts *bind.FilterOpts, from []common.Address) (*GatewayEVMZEVMTestWithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestWithdrawalIterator{contract: _GatewayEVMZEVMTest.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x2265ce9ec38ea098a1143406678482665a6e1ccd82ab22d37eea3a78abc57716.
//
// Solidity: event Withdrawal(address indexed from, address zrc20, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestWithdrawal, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestWithdrawal)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Withdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawal is a log parse operation binding the contract event 0x2265ce9ec38ea098a1143406678482665a6e1ccd82ab22d37eea3a78abc57716.
//
// Solidity: event Withdrawal(address indexed from, address zrc20, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseWithdrawal(log types.Log) (*GatewayEVMZEVMTestWithdrawal, error) {
	event := new(GatewayEVMZEVMTestWithdrawal)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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
