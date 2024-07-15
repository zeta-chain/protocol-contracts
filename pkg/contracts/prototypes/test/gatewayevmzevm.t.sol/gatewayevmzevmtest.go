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
	ABI: "[{\"inputs\":[],\"name\":\"ApprovalFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerIsNotFungibleModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CustodyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasFeeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientERC20Amount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientETHAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientZRC20Amount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZRC20BurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZRC20TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Call\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"Call\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ExecutedWithERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"ReceivedERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ReceivedNoParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"strs\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nums\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"ReceivedNonPayable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"ReceivedPayable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasfee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFlatFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"log_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"log_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"log_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"log_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"log_named_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"log_named_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"log_named_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"}],\"name\":\"log_named_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"log_named_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"log_named_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"log_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"logs\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"IS_TEST\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeArtifacts\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"excludedArtifacts_\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"excludedContracts_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"excludedSenders_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetArtifactSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"artifact\",\"type\":\"string\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetArtifacts\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"targetedArtifacts_\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targetedContracts_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetInterfaces\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"artifacts\",\"type\":\"string[]\"}],\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targetedSenders_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testCallReceiverEVMFromSenderZEVM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testCallReceiverEVMFromZEVM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testWithdrawAndCallReceiverEVMFromSenderZEVM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testWithdrawAndCallReceiverEVMFromZEVM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600c60006101000a81548160ff0219169083151502179055506001601f60006101000a81548160ff02191690831515021790555034801561004657600080fd5b506201343780620000586000396000f3fe60806040523480156200001157600080fd5b5060043610620001245760003560e01c806385226c8111620000b1578063b5508aa9116200007b578063b5508aa91462000269578063ba414fa6146200028b578063d7a525fc14620002ad578063e20c9f7114620002b9578063fa7626d414620002db5762000124565b806385226c8114620001f7578063916a17c614620002195780639683c695146200023b578063b0464fdc14620002475762000124565b80633f7286f411620000f35780633f7286f4146200019b5780635247441314620001bd57806366d9a9a014620001c95780636ff15ccc14620001eb5762000124565b80630a9254e414620001295780631ed7831c14620001355780632ade388014620001575780633e5e3c231462000179575b600080fd5b62000133620002fd565b005b6200013f620010fc565b6040516200014e9190620049d8565b60405180910390f35b620001616200118c565b60405162000170919062004a44565b60405180910390f35b6200018362001326565b604051620001929190620049d8565b60405180910390f35b620001a5620013b6565b604051620001b49190620049d8565b60405180910390f35b620001c762001446565b005b620001d362001e46565b604051620001e2919062004a20565b60405180910390f35b620001f562001fdd565b005b6200020162002958565b604051620002109190620049fc565b60405180910390f35b6200022362002a3b565b60405162000232919062004a68565b60405180910390f35b6200024562002b8e565b005b620002516200322d565b60405162000260919062004a68565b60405180910390f35b6200027362003380565b604051620002829190620049fc565b60405180910390f35b6200029562003463565b604051620002a4919062004a8c565b60405180910390f35b620002b762003596565b005b620002c362003d78565b604051620002d29190620049d8565b60405180910390f35b620002e562003e08565b604051620002f4919062004a8c565b60405180910390f35b30602360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611234602460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550615678602560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550614321602a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604051620004159062003eb1565b620004209062004e6e565b604051809103906000f0801580156200043d573d6000803e3d6000fd5b50602160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040516200048c9062003ebf565b604051809103906000f080158015620004a9573d6000803e3d6000fd5b50601f60016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516200051b9062003ecd565b620005279190620047ee565b604051809103906000f08015801562000544573d6000803e3d6000fd5b50602060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c4d66de8602560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401620006049190620047ee565b600060405180830381600087803b1580156200061f57600080fd5b505af115801562000634573d6000803e3d6000fd5b50505050601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ae7a3a6f602060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401620006b79190620047ee565b600060405180830381600087803b158015620006d257600080fd5b505af1158015620006e7573d6000803e3d6000fd5b50505050602160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f19602360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16620f42406040518363ffffffff1660e01b81526004016200076f929190620048ed565b600060405180830381600087803b1580156200078a57600080fd5b505af11580156200079f573d6000803e3d6000fd5b50505050602160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb602060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166207a1206040518363ffffffff1660e01b8152600401620008279291906200491a565b602060405180830381600087803b1580156200084257600080fd5b505af115801562000857573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200087d919062003fe4565b506040516200088c9062003edb565b604051809103906000f080158015620008a9573d6000803e3d6000fd5b50602260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604051620008f89062003ee9565b604051809103906000f08015801562000915573d6000803e3d6000fd5b50602660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051620009879062003ef7565b620009939190620047ee565b604051809103906000f080158015620009b0573d6000803e3d6000fd5b50602760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073735b14bb79463307aacbed86daf3322b1e6226ab90507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166306447d56826040518263ffffffff1660e01b815260040162000a689190620047ee565b600060405180830381600087803b15801562000a8357600080fd5b505af115801562000a98573d6000803e3d6000fd5b50505050600080600060405162000aaf9062003f05565b62000abd939291906200480b565b604051809103906000f08015801562000ada573d6000803e3d6000fd5b50602860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060126001600080602860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405162000b769062003f13565b62000b879695949392919062004dd6565b604051809103906000f08015801562000ba4573d6000803e3d6000fd5b50602960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ee2815ba6001602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518363ffffffff1660e01b815260040162000c6792919062004d38565b600060405180830381600087803b15801562000c8257600080fd5b505af115801562000c97573d6000803e3d6000fd5b50505050602860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a7cb05076001806040518363ffffffff1660e01b815260040162000cfb92919062004d65565b600060405180830381600087803b15801562000d1657600080fd5b505af115801562000d2b573d6000803e3d6000fd5b50505050602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166347e7ef24602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16620f42406040518363ffffffff1660e01b815260040162000db3929190620048ed565b602060405180830381600087803b15801562000dce57600080fd5b505af115801562000de3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000e09919062003fe4565b50602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166347e7ef24602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16620f42406040518363ffffffff1660e01b815260040162000e8e929190620048ed565b602060405180830381600087803b15801562000ea957600080fd5b505af115801562000ebe573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000ee4919062003fe4565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801562000f5157600080fd5b505af115801562000f66573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663ca669fa7602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040162000fea9190620047ee565b600060405180830381600087803b1580156200100557600080fd5b505af11580156200101a573d6000803e3d6000fd5b50505050602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095ea7b3602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16620f42406040518363ffffffff1660e01b8152600401620010a2929190620048ed565b602060405180830381600087803b158015620010bd57600080fd5b505af1158015620010d2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620010f8919062003fe4565b5050565b606060168054806020026020016040519081016040528092919081815260200182805480156200118257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162001137575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020016000905b828210156200131d57838290600052602060002090600202016040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805480602002602001604051908101604052809291908181526020016000905b82821015620013055783829060005260206000200180546200127190620052ff565b80601f01602080910402602001604051908101604052809291908181526020018280546200129f90620052ff565b8015620012f05780601f10620012c457610100808354040283529160200191620012f0565b820191906000526020600020905b815481529060010190602001808311620012d257829003601f168201915b5050505050815260200190600101906200124f565b505050508152505081526020019060010190620011b0565b50505050905090565b60606018805480602002602001604051908101604052809291908181526020018280548015620013ac57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162001361575b5050505050905090565b606060178054806020026020016040519081016040528092919081815260200182805480156200143c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311620013f1575b5050505050905090565b60006040518060400160405280600681526020017f48656c6c6f21000000000000000000000000000000000000000000000000000081525090506000602a90506000600190506000670de0b6b3a764000090506000602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b81526004016200151a9190620047ee565b60206040518083038186803b1580156200153357600080fd5b505afa15801562001548573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200156e919062004099565b9050600063e04d4f9760e01b868686604051602401620015919392919062004d92565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090506000602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001620016279190620047d1565b604051602081830303815290604052620f4240602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168460405160240162001673949392919062004b41565b6040516020818303038152906040527f7993c1e0000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663f30c7ba3602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000846040518463ffffffff1660e01b81526004016200177793929190620048a9565b600060405180830381600087803b1580156200179257600080fd5b505af1158015620017a7573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663ca669fa7602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b81526004016200182b9190620047ee565b600060405180830381600087803b1580156200184657600080fd5b505af11580156200185b573d6000803e3d6000fd5b50505050602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630abd8905602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001620018d29190620047d1565b604051602081830303815290604052620f4240602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168b8b8b6040518763ffffffff1660e01b81526004016200192e9695949392919062004bf7565b600060405180830381600087803b1580156200194957600080fd5b505af11580156200195e573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663c88a5e6d601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16866040518363ffffffff1660e01b8152600401620019e492919062004947565b600060405180830381600087803b158015620019ff57600080fd5b505af115801562001a14573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166381bad6f3600180600180602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b815260040162001aa295949392919062004aa9565b600060405180830381600087803b15801562001abd57600080fd5b505af115801562001ad2573d6000803e3d6000fd5b505050507f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff168589898960405162001b3195949392919062004974565b60405180910390a17f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166381bad6f3600180600180601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b815260040162001bc395949392919062004aa9565b600060405180830381600087803b15801562001bde57600080fd5b505af115801562001bf3573d6000803e3d6000fd5b50505050602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f858460405162001c6392919062004ea5565b60405180910390a2601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631cff79cd85602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518463ffffffff1660e01b815260040162001ced92919062004875565b6000604051808303818588803b15801562001d0757600080fd5b505af115801562001d1c573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f8201168201806040525081019062001d48919062004048565b506000602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040162001dca9190620047ee565b60206040518083038186803b15801562001de357600080fd5b505afa15801562001df8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062001e1e919062004099565b905062001e3c81620f42408662001e36919062005115565b62003e1b565b5050505050505050565b6060601b805480602002602001604051908101604052809291908181526020016000905b8282101562001fd4578382906000526020600020906002020160405180604001604052908160008201805462001ea090620052ff565b80601f016020809104026020016040519081016040528092919081815260200182805462001ece90620052ff565b801562001f1f5780601f1062001ef35761010080835404028352916020019162001f1f565b820191906000526020600020905b81548152906001019060200180831162001f0157829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801562001fbb57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001906004019060208260030104928301926001038202915080841162001f675790505b5050505050815250508152602001906001019062001e6a565b50505050905090565b60006040518060400160405280600681526020017f48656c6c6f21000000000000000000000000000000000000000000000000000081525090506000602a90506000600190506000670de0b6b3a76400009050600063e04d4f9760e01b858585604051602401620020519392919062004d92565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166381bad6f3600180600180602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b81526004016200213a95949392919062004aa9565b600060405180830381600087803b1580156200215557600080fd5b505af11580156200216a573d6000803e3d6000fd5b50505050602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001620021fd9190620047d1565b604051602081830303815290604052620f42406000602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634d8943bb6040518163ffffffff1660e01b815260040160206040518083038186803b1580156200227b57600080fd5b505afa15801562002290573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620022b6919062004099565b86604051620022ca95949392919062004c72565b60405180910390a27f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663ca669fa7602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401620023529190620047ee565b600060405180830381600087803b1580156200236d57600080fd5b505af115801562002382573d6000803e3d6000fd5b50505050602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637993c1e0602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001620023f99190620047d1565b604051602081830303815290604052620f4240602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518563ffffffff1660e01b815260040162002451949392919062004b9c565b600060405180830381600087803b1580156200246c57600080fd5b505af115801562002481573d6000803e3d6000fd5b505050506000602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401620025069190620047ee565b60206040518083038186803b1580156200251f57600080fd5b505afa15801562002534573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200255a919062004099565b90506200256981600062003e1b565b7f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663c88a5e6d601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518363ffffffff1660e01b8152600401620025eb92919062004947565b600060405180830381600087803b1580156200260657600080fd5b505af11580156200261b573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166381bad6f3600180600180602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b8152600401620026a995949392919062004aa9565b600060405180830381600087803b158015620026c457600080fd5b505af1158015620026d9573d6000803e3d6000fd5b505050507f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848888886040516200273895949392919062004974565b60405180910390a17f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166381bad6f3600180600180601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b8152600401620027ca95949392919062004aa9565b600060405180830381600087803b158015620027e557600080fd5b505af1158015620027fa573d6000803e3d6000fd5b50505050602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f84846040516200286a92919062004ea5565b60405180910390a2601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631cff79cd84602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518463ffffffff1660e01b8152600401620028f492919062004875565b6000604051808303818588803b1580156200290e57600080fd5b505af115801562002923573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f820116820180604052508101906200294f919062004048565b50505050505050565b6060601a805480602002602001604051908101604052809291908181526020016000905b8282101562002a325783829060005260206000200180546200299e90620052ff565b80601f0160208091040260200160405190810160405280929190818152602001828054620029cc90620052ff565b801562002a1d5780601f10620029f15761010080835404028352916020019162002a1d565b820191906000526020600020905b815481529060010190602001808311620029ff57829003601f168201915b5050505050815260200190600101906200297c565b50505050905090565b6060601d805480602002602001604051908101604052809291908181526020016000905b8282101562002b8557838290600052602060002090600202016040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820180548060200260200160405190810160405280929190818152602001828054801562002b6c57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001906004019060208260030104928301926001038202915080841162002b185790505b5050505050815250508152602001906001019062002a5f565b50505050905090565b60006040518060400160405280600681526020017f48656c6c6f21000000000000000000000000000000000000000000000000000081525090506000602a90506000600190506000670de0b6b3a76400009050600063e04d4f9760e01b85858560405160240162002c029392919062004d92565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663ca669fa7602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040162002ce19190620047ee565b600060405180830381600087803b15801562002cfc57600080fd5b505af115801562002d11573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166381bad6f3600180600180602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b815260040162002d9f95949392919062004aa9565b600060405180830381600087803b15801562002dba57600080fd5b505af115801562002dcf573d6000803e3d6000fd5b50505050602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405160200162002e629190620047d1565b6040516020818303038152906040528360405162002e8292919062004b06565b60405180910390a2602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630ac7c44c602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405160200162002efd9190620047d1565b604051602081830303815290604052836040518363ffffffff1660e01b815260040162002f2c92919062004b06565b600060405180830381600087803b15801562002f4757600080fd5b505af115801562002f5c573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663c88a5e6d601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b815260040162002fe292919062004947565b600060405180830381600087803b15801562002ffd57600080fd5b505af115801562003012573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166381bad6f3600180600180601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b8152600401620030a095949392919062004aa9565b600060405180830381600087803b158015620030bb57600080fd5b505af1158015620030d0573d6000803e3d6000fd5b50505050602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f83836040516200314092919062004ea5565b60405180910390a2601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631cff79cd83602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401620031ca92919062004875565b6000604051808303818588803b158015620031e457600080fd5b505af1158015620031f9573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f8201168201806040525081019062003225919062004048565b505050505050565b6060601c805480602002602001604051908101604052809291908181526020016000905b828210156200337757838290600052602060002090600202016040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054806020026020016040519081016040528092919081815260200182805480156200335e57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116200330a5790505b5050505050815250508152602001906001019062003251565b50505050905090565b60606019805480602002602001604051908101604052809291908181526020016000905b828210156200345a578382906000526020600020018054620033c690620052ff565b80601f0160208091040260200160405190810160405280929190818152602001828054620033f490620052ff565b8015620034455780601f10620034195761010080835404028352916020019162003445565b820191906000526020600020905b8154815290600101906020018083116200342757829003601f168201915b505050505081526020019060010190620033a4565b50505050905090565b6000600860009054906101000a900460ff16156200349357600860009054906101000a900460ff16905062003593565b6000801b7f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663667f9d707f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c7f6661696c656400000000000000000000000000000000000000000000000000006040518363ffffffff1660e01b81526004016200353a92919062004848565b60206040518083038186803b1580156200355357600080fd5b505afa15801562003568573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200358e919062004016565b141590505b90565b60006040518060400160405280600681526020017f48656c6c6f21000000000000000000000000000000000000000000000000000081525090506000602a90506000600190506000670de0b6b3a76400009050600063e04d4f9760e01b8585856040516024016200360a9392919062004d92565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090506000602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001620036a09190620047d1565b60405160208183030381529060405282604051602401620036c392919062004b06565b6040516020818303038152906040527f0ac7c44c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663f30c7ba3602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000846040518463ffffffff1660e01b8152600401620037c793929190620048a9565b600060405180830381600087803b158015620037e257600080fd5b505af1158015620037f7573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663ca669fa7602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b81526004016200387b9190620047ee565b600060405180830381600087803b1580156200389657600080fd5b505af1158015620038ab573d6000803e3d6000fd5b50505050602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a0a1730b602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001620039229190620047d1565b6040516020818303038152906040528888886040518563ffffffff1660e01b815260040162003955949392919062004cdd565b600060405180830381600087803b1580156200397057600080fd5b505af115801562003985573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff1663c88a5e6d601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518363ffffffff1660e01b815260040162003a0b92919062004947565b600060405180830381600087803b15801562003a2657600080fd5b505af115801562003a3b573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166381bad6f3600180600180602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b815260040162003ac995949392919062004aa9565b600060405180830381600087803b15801562003ae457600080fd5b505af115801562003af9573d6000803e3d6000fd5b505050507f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff168488888860405162003b5895949392919062004974565b60405180910390a17f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166381bad6f3600180600180601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b815260040162003bea95949392919062004aa9565b600060405180830381600087803b15801562003c0557600080fd5b505af115801562003c1a573d6000803e3d6000fd5b50505050602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f848460405162003c8a92919062004ea5565b60405180910390a2601f60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631cff79cd84602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518463ffffffff1660e01b815260040162003d1492919062004875565b6000604051808303818588803b15801562003d2e57600080fd5b505af115801562003d43573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f8201168201806040525081019062003d6f919062004048565b50505050505050565b6060601580548060200260200160405190810160405280929190818152602001828054801562003dfe57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162003db3575b5050505050905090565b601f60009054906101000a900460ff1681565b7f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c73ffffffffffffffffffffffffffffffffffffffff166398296c5483836040518363ffffffff1660e01b815260040162003e7b92919062004ed9565b60006040518083038186803b15801562003e9457600080fd5b505afa15801562003ea9573d6000803e3d6000fd5b505050505050565b611813806200558b83390190565b6131a48062006d9e83390190565b6110908062009f4283390190565b6110aa806200afd283390190565b612eb5806200c07c83390190565b610bcd806200ef3183390190565b6110d7806200fafe83390190565b61282d8062010bd583390190565b600062003f3862003f328462004f2f565b62004f06565b90508281526020810184848401111562003f575762003f5662005454565b5b62003f64848285620052c9565b509392505050565b60008151905062003f7d816200553c565b92915050565b60008151905062003f948162005556565b92915050565b600082601f83011262003fb25762003fb16200544f565b5b815162003fc484826020860162003f21565b91505092915050565b60008151905062003fde8162005570565b92915050565b60006020828403121562003ffd5762003ffc6200545e565b5b60006200400d8482850162003f6c565b91505092915050565b6000602082840312156200402f576200402e6200545e565b5b60006200403f8482850162003f83565b91505092915050565b6000602082840312156200406157620040606200545e565b5b600082015167ffffffffffffffff81111562004082576200408162005459565b5b620040908482850162003f9a565b91505092915050565b600060208284031215620040b257620040b16200545e565b5b6000620040c28482850162003fcd565b91505092915050565b6000620040d9838362004157565b60208301905092915050565b6000620040f38383620044f4565b60208301905092915050565b60006200410d8383620045d8565b905092915050565b6000620041238383620046f6565b905092915050565b60006200413983836200473e565b905092915050565b60006200414f83836200477f565b905092915050565b620041628162005150565b82525050565b620041738162005150565b82525050565b6000620041868262004fc5565b6200419281856200506b565b93506200419f8362004f65565b8060005b83811015620041d6578151620041ba8882620040cb565b9750620041c7836200501d565b925050600181019050620041a3565b5085935050505092915050565b6000620041f08262004fd0565b620041fc81856200507c565b9350620042098362004f75565b8060005b8381101562004240578151620042248882620040e5565b975062004231836200502a565b9250506001810190506200420d565b5085935050505092915050565b60006200425a8262004fdb565b6200426681856200508d565b9350836020820285016200427a8562004f85565b8060005b85811015620042bc57848403895281516200429a8582620040ff565b9450620042a78362005037565b925060208a019950506001810190506200427e565b50829750879550505050505092915050565b6000620042db8262004fdb565b620042e781856200509e565b935083602082028501620042fb8562004f85565b8060005b858110156200433d57848403895281516200431b8582620040ff565b9450620043288362005037565b925060208a01995050600181019050620042ff565b50829750879550505050505092915050565b60006200435c8262004fe6565b620043688185620050af565b9350836020820285016200437c8562004f95565b8060005b85811015620043be57848403895281516200439c858262004115565b9450620043a98362005044565b925060208a0199505060018101905062004380565b50829750879550505050505092915050565b6000620043dd8262004ff1565b620043e98185620050c0565b935083602082028501620043fd8562004fa5565b8060005b858110156200443f57848403895281516200441d85826200412b565b94506200442a8362005051565b925060208a0199505060018101905062004401565b50829750879550505050505092915050565b60006200445e8262004ffc565b6200446a8185620050d1565b9350836020820285016200447e8562004fb5565b8060005b85811015620044c057848403895281516200449e858262004141565b9450620044ab836200505e565b925060208a0199505060018101905062004482565b50829750879550505050505092915050565b620044dd8162005164565b82525050565b620044ee8162005170565b82525050565b620044ff816200517a565b82525050565b6000620045128262005007565b6200451e8185620050e2565b935062004530818560208601620052c9565b6200453b8162005463565b840191505092915050565b6200455b620045558262005201565b6200536b565b82525050565b6200456c8162005215565b82525050565b6200457d8162005229565b82525050565b6200458e816200523d565b82525050565b6200459f8162005251565b82525050565b620045b08162005265565b82525050565b620045c18162005279565b82525050565b620045d2816200528d565b82525050565b6000620045e58262005012565b620045f18185620050f3565b935062004603818560208601620052c9565b6200460e8162005463565b840191505092915050565b6000620046268262005012565b62004632818562005104565b935062004644818560208601620052c9565b6200464f8162005463565b840191505092915050565b60006200466960038362005104565b9150620046768262005481565b602082019050919050565b60006200469060058362005104565b91506200469d82620054aa565b602082019050919050565b6000620046b760048362005104565b9150620046c482620054d3565b602082019050919050565b6000620046de60038362005104565b9150620046eb82620054fc565b602082019050919050565b60006040830160008301518482036000860152620047158282620045d8565b91505060208301518482036020860152620047318282620041e3565b9150508091505092915050565b600060408301600083015162004758600086018262004157565b50602083015184820360208601526200477282826200424d565b9150508091505092915050565b600060408301600083015162004799600086018262004157565b5060208301518482036020860152620047b38282620041e3565b9150508091505092915050565b620047cb81620051ea565b82525050565b6000620047df828462004546565b60148201915081905092915050565b600060208201905062004805600083018462004168565b92915050565b600060608201905062004822600083018662004168565b62004831602083018562004168565b62004840604083018462004168565b949350505050565b60006040820190506200485f600083018562004168565b6200486e6020830184620044e3565b9392505050565b60006040820190506200488c600083018562004168565b8181036020830152620048a0818462004505565b90509392505050565b6000606082019050620048c0600083018662004168565b620048cf602083018562004572565b8181036040830152620048e3818462004505565b9050949350505050565b600060408201905062004904600083018562004168565b62004913602083018462004594565b9392505050565b600060408201905062004931600083018562004168565b620049406020830184620045c7565b9392505050565b60006040820190506200495e600083018562004168565b6200496d6020830184620047c0565b9392505050565b600060a0820190506200498b600083018862004168565b6200499a6020830187620047c0565b8181036040830152620049ae818662004619565b9050620049bf6060830185620047c0565b620049ce6080830184620044d2565b9695505050505050565b60006020820190508181036000830152620049f4818462004179565b905092915050565b6000602082019050818103600083015262004a188184620042ce565b905092915050565b6000602082019050818103600083015262004a3c81846200434f565b905092915050565b6000602082019050818103600083015262004a608184620043d0565b905092915050565b6000602082019050818103600083015262004a84818462004451565b905092915050565b600060208201905062004aa36000830184620044d2565b92915050565b600060a08201905062004ac06000830188620044d2565b62004acf6020830187620044d2565b62004ade6040830186620044d2565b62004aed6060830185620044d2565b62004afc608083018462004168565b9695505050505050565b6000604082019050818103600083015262004b22818562004505565b9050818103602083015262004b38818462004505565b90509392505050565b6000608082019050818103600083015262004b5d818762004505565b905062004b6e602083018662004583565b62004b7d604083018562004168565b818103606083015262004b91818462004505565b905095945050505050565b6000608082019050818103600083015262004bb8818762004505565b905062004bc9602083018662004594565b62004bd8604083018562004168565b818103606083015262004bec818462004505565b905095945050505050565b600060c082019050818103600083015262004c13818962004505565b905062004c24602083018862004594565b62004c33604083018762004168565b818103606083015262004c47818662004619565b905062004c586080830185620047c0565b62004c6760a0830184620044d2565b979650505050505050565b600060a082019050818103600083015262004c8e818862004505565b905062004c9f602083018762004594565b62004cae604083018662004572565b62004cbd6060830185620047c0565b818103608083015262004cd1818462004505565b90509695505050505050565b6000608082019050818103600083015262004cf9818762004505565b9050818103602083015262004d0f818662004619565b905062004d206040830185620047c0565b62004d2f6060830184620044d2565b95945050505050565b600060408201905062004d4f6000830185620045b6565b62004d5e602083018462004168565b9392505050565b600060408201905062004d7c6000830185620045b6565b62004d8b6020830184620045b6565b9392505050565b6000606082019050818103600083015262004dae818662004619565b905062004dbf6020830185620047c0565b62004dce6040830184620044d2565b949350505050565b600061010082019050818103600083015262004df28162004681565b9050818103602083015262004e0781620046cf565b905062004e186040830189620045a5565b62004e276060830188620045b6565b62004e36608083018762004561565b62004e4560a083018662004572565b62004e5460c083018562004168565b62004e6360e083018462004168565b979650505050505050565b6000604082019050818103600083015262004e8981620046a8565b9050818103602083015262004e9e816200465a565b9050919050565b600060408201905062004ebc6000830185620047c0565b818103602083015262004ed0818462004505565b90509392505050565b600060408201905062004ef06000830185620047c0565b62004eff6020830184620047c0565b9392505050565b600062004f1262004f25565b905062004f20828262005335565b919050565b6000604051905090565b600067ffffffffffffffff82111562004f4d5762004f4c62005420565b5b62004f588262005463565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006200512282620051ea565b91506200512f83620051ea565b92508282101562005145576200514462005393565b5b828203905092915050565b60006200515d82620051bb565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6000819050620051b68262005525565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006200520e82620052a1565b9050919050565b60006200522282620051a6565b9050919050565b60006200523682620051ea565b9050919050565b60006200524a82620051db565b9050919050565b60006200525e82620051ea565b9050919050565b60006200527282620051f4565b9050919050565b60006200528682620051ea565b9050919050565b60006200529a82620051ea565b9050919050565b6000620052ae82620052b5565b9050919050565b6000620052c282620051bb565b9050919050565b60005b83811015620052e9578082015181840152602081019050620052cc565b83811115620052f9576000848401525b50505050565b600060028204905060018216806200531857607f821691505b602082108114156200532f576200532e620053f1565b5b50919050565b620053408262005463565b810181811067ffffffffffffffff8211171562005362576200536162005420565b5b80604052505050565b600062005378826200537f565b9050919050565b60006200538c8262005474565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f54544b0000000000000000000000000000000000000000000000000000000000600082015250565b7f544f4b454e000000000000000000000000000000000000000000000000000000600082015250565b7f7465737400000000000000000000000000000000000000000000000000000000600082015250565b7f544b4e0000000000000000000000000000000000000000000000000000000000600082015250565b60038110620055395762005538620053c2565b5b50565b620055478162005164565b81146200555357600080fd5b50565b620055618162005170565b81146200556d57600080fd5b50565b6200557b81620051ea565b81146200558757600080fd5b5056fe60806040523480156200001157600080fd5b5060405162001813380380620018138339818101604052810190620000379190620001a3565b818181600390805190602001906200005192919062000075565b5080600490805190602001906200006a92919062000075565b5050505050620003ac565b8280546200008390620002bd565b90600052602060002090601f016020900481019282620000a75760008555620000f3565b82601f10620000c257805160ff1916838001178555620000f3565b82800160010185558215620000f3579182015b82811115620000f2578251825591602001919060010190620000d5565b5b50905062000102919062000106565b5090565b5b808211156200012157600081600090555060010162000107565b5090565b60006200013c620001368462000251565b62000228565b9050828152602081018484840111156200015b576200015a6200038c565b5b6200016884828562000287565b509392505050565b600082601f83011262000188576200018762000387565b5b81516200019a84826020860162000125565b91505092915050565b60008060408385031215620001bd57620001bc62000396565b5b600083015167ffffffffffffffff811115620001de57620001dd62000391565b5b620001ec8582860162000170565b925050602083015167ffffffffffffffff81111562000210576200020f62000391565b5b6200021e8582860162000170565b9150509250929050565b60006200023462000247565b9050620002428282620002f3565b919050565b6000604051905090565b600067ffffffffffffffff8211156200026f576200026e62000358565b5b6200027a826200039b565b9050602081019050919050565b60005b83811015620002a75780820151818401526020810190506200028a565b83811115620002b7576000848401525b50505050565b60006002820490506001821680620002d657607f821691505b60208210811415620002ed57620002ec62000329565b5b50919050565b620002fe826200039b565b810181811067ffffffffffffffff8211171562000320576200031f62000358565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b61145780620003bc6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806340c10f191161007157806340c10f19146101a357806370a08231146101bf57806395d89b41146101ef578063a457c2d71461020d578063a9059cbb1461023d578063dd62ed3e1461026d576100b4565b806306fdde03146100b9578063095ea7b3146100d757806318160ddd1461010757806323b872dd14610125578063313ce567146101555780633950935114610173575b600080fd5b6100c161029d565b6040516100ce9190610ecf565b60405180910390f35b6100f160048036038101906100ec9190610cf6565b61032f565b6040516100fe9190610eb4565b60405180910390f35b61010f610352565b60405161011c9190610ff1565b60405180910390f35b61013f600480360381019061013a9190610ca3565b61035c565b60405161014c9190610eb4565b60405180910390f35b61015d61038b565b60405161016a919061100c565b60405180910390f35b61018d60048036038101906101889190610cf6565b610394565b60405161019a9190610eb4565b60405180910390f35b6101bd60048036038101906101b89190610cf6565b6103cb565b005b6101d960048036038101906101d49190610c36565b6103d9565b6040516101e69190610ff1565b60405180910390f35b6101f7610421565b6040516102049190610ecf565b60405180910390f35b61022760048036038101906102229190610cf6565b6104b3565b6040516102349190610eb4565b60405180910390f35b61025760048036038101906102529190610cf6565b61052a565b6040516102649190610eb4565b60405180910390f35b61028760048036038101906102829190610c63565b61054d565b6040516102949190610ff1565b60405180910390f35b6060600380546102ac90611121565b80601f01602080910402602001604051908101604052809291908181526020018280546102d890611121565b80156103255780601f106102fa57610100808354040283529160200191610325565b820191906000526020600020905b81548152906001019060200180831161030857829003601f168201915b5050505050905090565b60008061033a6105d4565b90506103478185856105dc565b600191505092915050565b6000600254905090565b6000806103676105d4565b90506103748582856107a7565b61037f858585610833565b60019150509392505050565b60006012905090565b60008061039f6105d4565b90506103c08185856103b1858961054d565b6103bb9190611043565b6105dc565b600191505092915050565b6103d58282610aab565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606004805461043090611121565b80601f016020809104026020016040519081016040528092919081815260200182805461045c90611121565b80156104a95780601f1061047e576101008083540402835291602001916104a9565b820191906000526020600020905b81548152906001019060200180831161048c57829003601f168201915b5050505050905090565b6000806104be6105d4565b905060006104cc828661054d565b905083811015610511576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050890610fb1565b60405180910390fd5b61051e82868684036105dc565b60019250505092915050565b6000806105356105d4565b9050610542818585610833565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561064c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064390610f91565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156106bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b390610f11565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258360405161079a9190610ff1565b60405180910390a3505050565b60006107b3848461054d565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461082d578181101561081f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081690610f31565b60405180910390fd5b61082c84848484036105dc565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156108a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089a90610f71565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610913576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090a90610ef1565b60405180910390fd5b61091e838383610c02565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156109a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099b90610f51565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a929190610ff1565b60405180910390a3610aa5848484610c07565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610b1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1290610fd1565b60405180910390fd5b610b2760008383610c02565b8060026000828254610b399190611043565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610bea9190610ff1565b60405180910390a3610bfe60008383610c07565b5050565b505050565b505050565b600081359050610c1b816113f3565b92915050565b600081359050610c308161140a565b92915050565b600060208284031215610c4c57610c4b6111b1565b5b6000610c5a84828501610c0c565b91505092915050565b60008060408385031215610c7a57610c796111b1565b5b6000610c8885828601610c0c565b9250506020610c9985828601610c0c565b9150509250929050565b600080600060608486031215610cbc57610cbb6111b1565b5b6000610cca86828701610c0c565b9350506020610cdb86828701610c0c565b9250506040610cec86828701610c21565b9150509250925092565b60008060408385031215610d0d57610d0c6111b1565b5b6000610d1b85828601610c0c565b9250506020610d2c85828601610c21565b9150509250929050565b610d3f816110ab565b82525050565b6000610d5082611027565b610d5a8185611032565b9350610d6a8185602086016110ee565b610d73816111b6565b840191505092915050565b6000610d8b602383611032565b9150610d96826111c7565b604082019050919050565b6000610dae602283611032565b9150610db982611216565b604082019050919050565b6000610dd1601d83611032565b9150610ddc82611265565b602082019050919050565b6000610df4602683611032565b9150610dff8261128e565b604082019050919050565b6000610e17602583611032565b9150610e22826112dd565b604082019050919050565b6000610e3a602483611032565b9150610e458261132c565b604082019050919050565b6000610e5d602583611032565b9150610e688261137b565b604082019050919050565b6000610e80601f83611032565b9150610e8b826113ca565b602082019050919050565b610e9f816110d7565b82525050565b610eae816110e1565b82525050565b6000602082019050610ec96000830184610d36565b92915050565b60006020820190508181036000830152610ee98184610d45565b905092915050565b60006020820190508181036000830152610f0a81610d7e565b9050919050565b60006020820190508181036000830152610f2a81610da1565b9050919050565b60006020820190508181036000830152610f4a81610dc4565b9050919050565b60006020820190508181036000830152610f6a81610de7565b9050919050565b60006020820190508181036000830152610f8a81610e0a565b9050919050565b60006020820190508181036000830152610faa81610e2d565b9050919050565b60006020820190508181036000830152610fca81610e50565b9050919050565b60006020820190508181036000830152610fea81610e73565b9050919050565b60006020820190506110066000830184610e96565b92915050565b60006020820190506110216000830184610ea5565b92915050565b600081519050919050565b600082825260208201905092915050565b600061104e826110d7565b9150611059836110d7565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561108e5761108d611153565b5b828201905092915050565b60006110a4826110b7565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b8381101561110c5780820151818401526020810190506110f1565b8381111561111b576000848401525b50505050565b6000600282049050600182168061113957607f821691505b6020821081141561114d5761114c611182565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b6000601f19601f8301169050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6113fc81611099565b811461140757600080fd5b50565b611413816110d7565b811461141e57600080fd5b5056fea2646970667358221220d4c96329400732e284f624232c8d7228fc90c00ee7b1a48a85080262107815d864736f6c6343000807003360a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b81525034801561004657600080fd5b5060805160601c613123610081600039600081816105fc0152818161068b01528181610785015281816108140152610bae01526131236000f3fe6080604052600436106100fe5760003560e01c8063715018a611610095578063c4d66de811610064578063c4d66de8146102e4578063dda79b751461030d578063f2fde38b14610338578063f340fa0114610361578063f45346dc1461037d576100fe565b8063715018a6146102505780638c6f037f146102675780638da5cb5b14610290578063ae7a3a6f146102bb576100fe565b80634f1ef286116100d15780634f1ef286146101a15780635131ab59146101bd57806352d1902d146101fa5780635b11259114610225576100fe565b80631b8b921d146101035780631cff79cd1461012c57806329c59b5d1461015c5780633659cfe614610178575b600080fd5b34801561010f57600080fd5b5061012a60048036038101906101259190612183565b6103a6565b005b61014660048036038101906101419190612183565b610412565b6040516101539190612816565b60405180910390f35b61017660048036038101906101719190612183565b610480565b005b34801561018457600080fd5b5061019f600480360381019061019a91906120ce565b6105fa565b005b6101bb60048036038101906101b691906121e3565b610783565b005b3480156101c957600080fd5b506101e460048036038101906101df91906120fb565b6108c0565b6040516101f19190612816565b60405180910390f35b34801561020657600080fd5b5061020f610baa565b60405161021c91906127d7565b60405180910390f35b34801561023157600080fd5b5061023a610c63565b6040516102479190612733565b60405180910390f35b34801561025c57600080fd5b50610265610c89565b005b34801561027357600080fd5b5061028e60048036038101906102899190612292565b610c9d565b005b34801561029c57600080fd5b506102a5610d99565b6040516102b29190612733565b60405180910390f35b3480156102c757600080fd5b506102e260048036038101906102dd91906120ce565b610dc3565b005b3480156102f057600080fd5b5061030b600480360381019061030691906120ce565b610e8f565b005b34801561031957600080fd5b5061032261107e565b60405161032f9190612733565b60405180910390f35b34801561034457600080fd5b5061035f600480360381019061035a91906120ce565b6110a4565b005b61037b600480360381019061037691906120ce565b611128565b005b34801561038957600080fd5b506103a4600480360381019061039f919061223f565b61129c565b005b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde384846040516104059291906127f2565b60405180910390a3505050565b60606000610421858585611392565b90508473ffffffffffffffffffffffffffffffffffffffff167fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f34868660405161046d93929190612a91565b60405180910390a2809150509392505050565b60003414156104bb576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16346040516105039061271e565b60006040518083038185875af1925050503d8060008114610540576040519150601f19603f3d011682016040523d82523d6000602084013e610545565b606091505b50509050600015158115151415610588576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a434600087876040516105ec9493929190612a15565b60405180910390a350505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610689576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068090612895565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166106c8611449565b73ffffffffffffffffffffffffffffffffffffffff161461071e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610715906128b5565b60405180910390fd5b610727816114a0565b61078081600067ffffffffffffffff81111561074657610745612c52565b5b6040519080825280601f01601f1916602001820160405280156107785781602001600182028036833780820191505090505b5060006114ab565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610812576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080990612895565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610851611449565b73ffffffffffffffffffffffffffffffffffffffff16146108a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089e906128b5565b60405180910390fd5b6108b0826114a0565b6108bc828260016114ab565b5050565b606060008414156108fd576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109078686611628565b61093d576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff1663095ea7b386866040518363ffffffff1660e01b81526004016109789291906127ae565b602060405180830381600087803b15801561099257600080fd5b505af11580156109a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ca919061231a565b610a00576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610a0d868585611392565b9050610a198787611628565b610a4f576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008773ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610a8a9190612733565b60206040518083038186803b158015610aa257600080fd5b505afa158015610ab6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ada9190612374565b90506000811115610b3357610b3260c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16828a73ffffffffffffffffffffffffffffffffffffffff166116c09092919063ffffffff16565b5b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382888888604051610b9493929190612a91565b60405180910390a3819250505095945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610c3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c31906128f5565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b60ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610c91611746565b610c9b60006117c4565b565b6000841415610cd8576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d273360c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16868673ffffffffffffffffffffffffffffffffffffffff1661188a909392919063ffffffff16565b8473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a486868686604051610d8a9493929190612a15565b60405180910390a35050505050565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600073ffffffffffffffffffffffffffffffffffffffff1660c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610e4b576040517fb337f37800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060c960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060019054906101000a900460ff16159050808015610ec05750600160008054906101000a900460ff1660ff16105b80610eed5750610ecf30611913565b158015610eec5750600160008054906101000a900460ff1660ff16145b5b610f2c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2390612935565b60405180910390fd5b60016000806101000a81548160ff021916908360ff1602179055508015610f69576001600060016101000a81548160ff0219169083151502179055505b610f71611936565b610f7961198f565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610fe0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160ca60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550801561107a5760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516110719190612838565b60405180910390a15b5050565b60c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6110ac611746565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561111c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111390612875565b60405180910390fd5b611125816117c4565b50565b6000341415611163576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060ca60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16346040516111ab9061271e565b60006040518083038185875af1925050503d80600081146111e8576040519150601f19603f3d011682016040523d82523d6000602084013e6111ed565b606091505b50509050600015158115151415611230576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4346000604051611290929190612a55565b60405180910390a35050565b60008214156112d7576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113263360c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848473ffffffffffffffffffffffffffffffffffffffff1661188a909392919063ffffffff16565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a48484604051611385929190612a55565b60405180910390a3505050565b60606000808573ffffffffffffffffffffffffffffffffffffffff163486866040516113bf9291906126ee565b60006040518083038185875af1925050503d80600081146113fc576040519150601f19603f3d011682016040523d82523d6000602084013e611401565b606091505b50915091508161143d576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80925050509392505050565b60006114777f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b6119e0565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6114a8611746565b50565b6114d77f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b6119ea565b60000160009054906101000a900460ff16156114fb576114f6836119f4565b611623565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561154157600080fd5b505afa92505050801561157257506040513d601f19601f8201168201806040525081019061156f9190612347565b60015b6115b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115a890612955565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114611616576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161160d90612915565b60405180910390fd5b50611622838383611aad565b5b505050565b60008273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401611666929190612785565b602060405180830381600087803b15801561168057600080fd5b505af1158015611694573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116b8919061231a565b905092915050565b6117418363a9059cbb60e01b84846040516024016116df9291906127ae565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611ad9565b505050565b61174e611ba0565b73ffffffffffffffffffffffffffffffffffffffff1661176c610d99565b73ffffffffffffffffffffffffffffffffffffffff16146117c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117b990612995565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61190d846323b872dd60e01b8585856040516024016118ab9392919061274e565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611ad9565b50505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff16611985576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161197c906129d5565b60405180910390fd5b61198d611ba8565b565b600060019054906101000a900460ff166119de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119d5906129d5565b60405180910390fd5b565b6000819050919050565b6000819050919050565b6119fd81611913565b611a3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a3390612975565b60405180910390fd5b80611a697f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b6119e0565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611ab683611c09565b600082511180611ac35750805b15611ad457611ad28383611c58565b505b505050565b6000611b3b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611c859092919063ffffffff16565b9050600081511115611b9b5780806020019051810190611b5b919061231a565b611b9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b91906129f5565b60405180910390fd5b5b505050565b600033905090565b600060019054906101000a900460ff16611bf7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bee906129d5565b60405180910390fd5b611c07611c02611ba0565b6117c4565b565b611c12816119f4565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b6060611c7d83836040518060600160405280602781526020016130c760279139611c9d565b905092915050565b6060611c948484600085611d23565b90509392505050565b60606000808573ffffffffffffffffffffffffffffffffffffffff1685604051611cc79190612707565b600060405180830381855af49150503d8060008114611d02576040519150601f19603f3d011682016040523d82523d6000602084013e611d07565b606091505b5091509150611d1886838387611df0565b925050509392505050565b606082471015611d68576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d5f906128d5565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611d919190612707565b60006040518083038185875af1925050503d8060008114611dce576040519150601f19603f3d011682016040523d82523d6000602084013e611dd3565b606091505b5091509150611de487838387611e66565b92505050949350505050565b60608315611e5357600083511415611e4b57611e0b85611913565b611e4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e41906129b5565b60405180910390fd5b5b829050611e5e565b611e5d8383611edc565b5b949350505050565b60608315611ec957600083511415611ec157611e8185611f2c565b611ec0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611eb7906129b5565b60405180910390fd5b5b829050611ed4565b611ed38383611f4f565b5b949350505050565b600082511115611eef5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f239190612853565b60405180910390fd5b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082511115611f625781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f969190612853565b60405180910390fd5b6000611fb2611fad84612ae8565b612ac3565b905082815260208101848484011115611fce57611fcd612c90565b5b611fd9848285612bdf565b509392505050565b600081359050611ff08161306a565b92915050565b60008151905061200581613081565b92915050565b60008151905061201a81613098565b92915050565b60008083601f84011261203657612035612c86565b5b8235905067ffffffffffffffff81111561205357612052612c81565b5b60208301915083600182028301111561206f5761206e612c8b565b5b9250929050565b600082601f83011261208b5761208a612c86565b5b813561209b848260208601611f9f565b91505092915050565b6000813590506120b3816130af565b92915050565b6000815190506120c8816130af565b92915050565b6000602082840312156120e4576120e3612c9a565b5b60006120f284828501611fe1565b91505092915050565b60008060008060006080868803121561211757612116612c9a565b5b600061212588828901611fe1565b955050602061213688828901611fe1565b9450506040612147888289016120a4565b935050606086013567ffffffffffffffff81111561216857612167612c95565b5b61217488828901612020565b92509250509295509295909350565b60008060006040848603121561219c5761219b612c9a565b5b60006121aa86828701611fe1565b935050602084013567ffffffffffffffff8111156121cb576121ca612c95565b5b6121d786828701612020565b92509250509250925092565b600080604083850312156121fa576121f9612c9a565b5b600061220885828601611fe1565b925050602083013567ffffffffffffffff81111561222957612228612c95565b5b61223585828601612076565b9150509250929050565b60008060006060848603121561225857612257612c9a565b5b600061226686828701611fe1565b9350506020612277868287016120a4565b925050604061228886828701611fe1565b9150509250925092565b6000806000806000608086880312156122ae576122ad612c9a565b5b60006122bc88828901611fe1565b95505060206122cd888289016120a4565b94505060406122de88828901611fe1565b935050606086013567ffffffffffffffff8111156122ff576122fe612c95565b5b61230b88828901612020565b92509250509295509295909350565b6000602082840312156123305761232f612c9a565b5b600061233e84828501611ff6565b91505092915050565b60006020828403121561235d5761235c612c9a565b5b600061236b8482850161200b565b91505092915050565b60006020828403121561238a57612389612c9a565b5b6000612398848285016120b9565b91505092915050565b6123aa81612b5c565b82525050565b6123b981612b7a565b82525050565b60006123cb8385612b2f565b93506123d8838584612bdf565b6123e183612c9f565b840190509392505050565b60006123f88385612b40565b9350612405838584612bdf565b82840190509392505050565b600061241c82612b19565b6124268185612b2f565b9350612436818560208601612bee565b61243f81612c9f565b840191505092915050565b600061245582612b19565b61245f8185612b40565b935061246f818560208601612bee565b80840191505092915050565b61248481612bbb565b82525050565b61249381612bcd565b82525050565b60006124a482612b24565b6124ae8185612b4b565b93506124be818560208601612bee565b6124c781612c9f565b840191505092915050565b60006124df602683612b4b565b91506124ea82612cb0565b604082019050919050565b6000612502602c83612b4b565b915061250d82612cff565b604082019050919050565b6000612525602c83612b4b565b915061253082612d4e565b604082019050919050565b6000612548602683612b4b565b915061255382612d9d565b604082019050919050565b600061256b603883612b4b565b915061257682612dec565b604082019050919050565b600061258e602983612b4b565b915061259982612e3b565b604082019050919050565b60006125b1602e83612b4b565b91506125bc82612e8a565b604082019050919050565b60006125d4602e83612b4b565b91506125df82612ed9565b604082019050919050565b60006125f7602d83612b4b565b915061260282612f28565b604082019050919050565b600061261a602083612b4b565b915061262582612f77565b602082019050919050565b600061263d600083612b2f565b915061264882612fa0565b600082019050919050565b6000612660600083612b40565b915061266b82612fa0565b600082019050919050565b6000612683601d83612b4b565b915061268e82612fa3565b602082019050919050565b60006126a6602b83612b4b565b91506126b182612fcc565b604082019050919050565b60006126c9602a83612b4b565b91506126d48261301b565b604082019050919050565b6126e881612ba4565b82525050565b60006126fb8284866123ec565b91508190509392505050565b6000612713828461244a565b915081905092915050565b600061272982612653565b9150819050919050565b600060208201905061274860008301846123a1565b92915050565b600060608201905061276360008301866123a1565b61277060208301856123a1565b61277d60408301846126df565b949350505050565b600060408201905061279a60008301856123a1565b6127a7602083018461247b565b9392505050565b60006040820190506127c360008301856123a1565b6127d060208301846126df565b9392505050565b60006020820190506127ec60008301846123b0565b92915050565b6000602082019050818103600083015261280d8184866123bf565b90509392505050565b600060208201905081810360008301526128308184612411565b905092915050565b600060208201905061284d600083018461248a565b92915050565b6000602082019050818103600083015261286d8184612499565b905092915050565b6000602082019050818103600083015261288e816124d2565b9050919050565b600060208201905081810360008301526128ae816124f5565b9050919050565b600060208201905081810360008301526128ce81612518565b9050919050565b600060208201905081810360008301526128ee8161253b565b9050919050565b6000602082019050818103600083015261290e8161255e565b9050919050565b6000602082019050818103600083015261292e81612581565b9050919050565b6000602082019050818103600083015261294e816125a4565b9050919050565b6000602082019050818103600083015261296e816125c7565b9050919050565b6000602082019050818103600083015261298e816125ea565b9050919050565b600060208201905081810360008301526129ae8161260d565b9050919050565b600060208201905081810360008301526129ce81612676565b9050919050565b600060208201905081810360008301526129ee81612699565b9050919050565b60006020820190508181036000830152612a0e816126bc565b9050919050565b6000606082019050612a2a60008301876126df565b612a3760208301866123a1565b8181036040830152612a4a8184866123bf565b905095945050505050565b6000606082019050612a6a60008301856126df565b612a7760208301846123a1565b8181036040830152612a8881612630565b90509392505050565b6000604082019050612aa660008301866126df565b8181036020830152612ab98184866123bf565b9050949350505050565b6000612acd612ade565b9050612ad98282612c21565b919050565b6000604051905090565b600067ffffffffffffffff821115612b0357612b02612c52565b5b612b0c82612c9f565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000612b6782612b84565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000612bc682612ba4565b9050919050565b6000612bd882612bae565b9050919050565b82818337600083830152505050565b60005b83811015612c0c578082015181840152602081019050612bf1565b83811115612c1b576000848401525b50505050565b612c2a82612c9f565b810181811067ffffffffffffffff82111715612c4957612c48612c52565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b61307381612b5c565b811461307e57600080fd5b50565b61308a81612b6e565b811461309557600080fd5b50565b6130a181612b7a565b81146130ac57600080fd5b50565b6130b881612ba4565b81146130c357600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220bda16da68115d43d899e50b7ec6d5a336fd24cf187474373ee646397da0aedcc64736f6c6343000807003360806040523480156200001157600080fd5b506040516200109038038062001090833981810160405281019062000037919062000106565b6001600081905550600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620000a7576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200018b565b600081519050620001008162000171565b92915050565b6000602082840312156200011f576200011e6200016c565b5b60006200012f84828501620000ef565b91505092915050565b600062000145826200014c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200017c8162000138565b81146200018857600080fd5b50565b610ef5806200019b6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063116191b61461004657806321fc65f214610064578063d9caed1214610080575b600080fd5b61004e61009c565b60405161005b9190610a98565b60405180910390f35b61007e600480360381019061007991906107bc565b6100c2565b005b61009a60048036038101906100959190610769565b61024a565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6100ca6102ef565b610117600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848773ffffffffffffffffffffffffffffffffffffffff1661033f9092919063ffffffff16565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635131ab5986868686866040518663ffffffff1660e01b815260040161017a959493929190610a21565b600060405180830381600087803b15801561019457600080fd5b505af11580156101a8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906101d19190610871565b508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e85858560405161023393929190610b70565b60405180910390a36102436103c5565b5050505050565b6102526102ef565b61027d82828573ffffffffffffffffffffffffffffffffffffffff1661033f9092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb836040516102da9190610b55565b60405180910390a36102ea6103c5565b505050565b60026000541415610335576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032c90610b35565b60405180910390fd5b6002600081905550565b6103c08363a9059cbb60e01b848460405160240161035e929190610a6f565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506103cf565b505050565b6001600081905550565b6000610431826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166104969092919063ffffffff16565b905060008151111561049157808060200190518101906104519190610844565b610490576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048790610b15565b60405180910390fd5b5b505050565b60606104a584846000856104ae565b90509392505050565b6060824710156104f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ea90610ad5565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161051c9190610a0a565b60006040518083038185875af1925050503d8060008114610559576040519150601f19603f3d011682016040523d82523d6000602084013e61055e565b606091505b509150915061056f8783838761057b565b92505050949350505050565b606083156105de576000835114156105d657610596856105f1565b6105d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105cc90610af5565b60405180910390fd5b5b8290506105e9565b6105e88383610614565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156106275781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065b9190610ab3565b60405180910390fd5b600061067761067284610bc7565b610ba2565b90508281526020810184848401111561069357610692610d6a565b5b61069e848285610cc8565b509392505050565b6000813590506106b581610e7a565b92915050565b6000815190506106ca81610e91565b92915050565b60008083601f8401126106e6576106e5610d60565b5b8235905067ffffffffffffffff81111561070357610702610d5b565b5b60208301915083600182028301111561071f5761071e610d65565b5b9250929050565b600082601f83011261073b5761073a610d60565b5b815161074b848260208601610664565b91505092915050565b60008135905061076381610ea8565b92915050565b60008060006060848603121561078257610781610d74565b5b6000610790868287016106a6565b93505060206107a1868287016106a6565b92505060406107b286828701610754565b9150509250925092565b6000806000806000608086880312156107d8576107d7610d74565b5b60006107e6888289016106a6565b95505060206107f7888289016106a6565b945050604061080888828901610754565b935050606086013567ffffffffffffffff81111561082957610828610d6f565b5b610835888289016106d0565b92509250509295509295909350565b60006020828403121561085a57610859610d74565b5b6000610868848285016106bb565b91505092915050565b60006020828403121561088757610886610d74565b5b600082015167ffffffffffffffff8111156108a5576108a4610d6f565b5b6108b184828501610726565b91505092915050565b6108c381610c3b565b82525050565b60006108d58385610c0e565b93506108e2838584610cb9565b6108eb83610d79565b840190509392505050565b600061090182610bf8565b61090b8185610c1f565b935061091b818560208601610cc8565b80840191505092915050565b61093081610c83565b82525050565b600061094182610c03565b61094b8185610c2a565b935061095b818560208601610cc8565b61096481610d79565b840191505092915050565b600061097c602683610c2a565b915061098782610d8a565b604082019050919050565b600061099f601d83610c2a565b91506109aa82610dd9565b602082019050919050565b60006109c2602a83610c2a565b91506109cd82610e02565b604082019050919050565b60006109e5601f83610c2a565b91506109f082610e51565b602082019050919050565b610a0481610c79565b82525050565b6000610a1682846108f6565b915081905092915050565b6000608082019050610a3660008301886108ba565b610a4360208301876108ba565b610a5060408301866109fb565b8181036060830152610a638184866108c9565b90509695505050505050565b6000604082019050610a8460008301856108ba565b610a9160208301846109fb565b9392505050565b6000602082019050610aad6000830184610927565b92915050565b60006020820190508181036000830152610acd8184610936565b905092915050565b60006020820190508181036000830152610aee8161096f565b9050919050565b60006020820190508181036000830152610b0e81610992565b9050919050565b60006020820190508181036000830152610b2e816109b5565b9050919050565b60006020820190508181036000830152610b4e816109d8565b9050919050565b6000602082019050610b6a60008301846109fb565b92915050565b6000604082019050610b8560008301866109fb565b8181036020830152610b988184866108c9565b9050949350505050565b6000610bac610bbd565b9050610bb88282610cfb565b919050565b6000604051905090565b600067ffffffffffffffff821115610be257610be1610d2c565b5b610beb82610d79565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610c4682610c59565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610c8e82610c95565b9050919050565b6000610ca082610ca7565b9050919050565b6000610cb282610c59565b9050919050565b82818337600083830152505050565b60005b83811015610ce6578082015181840152602081019050610ccb565b83811115610cf5576000848401525b50505050565b610d0482610d79565b810181811067ffffffffffffffff82111715610d2357610d22610d2c565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b610e8381610c3b565b8114610e8e57600080fd5b50565b610e9a81610c4d565b8114610ea557600080fd5b50565b610eb181610c79565b8114610ebc57600080fd5b5056fea2646970667358221220c6f7c0f9935341aaaabcabf7504d01bbd38a0606132004742a97e6a200eb432564736f6c63430008070033608060405234801561001057600080fd5b5061108a806100206000396000f3fe60806040526004361061003f5760003560e01c8063357fc5a2146100445780636ed701691461006d578063e04d4f9714610084578063f05b6abf146100a0575b600080fd5b34801561005057600080fd5b5061006b6004803603810190610066919061085a565b6100c9565b005b34801561007957600080fd5b50610082610138565b005b61009e600480360381019061009991906107eb565b610171565b005b3480156100ac57600080fd5b506100c760048036038101906100c29190610733565b6101b5565b005b6100f63382858573ffffffffffffffffffffffffffffffffffffffff166101f7909392919063ffffffff16565b7f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af603384848460405161012b9493929190610bb0565b60405180910390a1505050565b7fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0336040516101679190610b0b565b60405180910390a1565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa33348585856040516101a8959493929190610bf5565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146338484846040516101ea9493929190610b5d565b60405180910390a1505050565b61027a846323b872dd60e01b85858560405160240161021893929190610b26565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610280565b50505050565b60006102e2826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166103479092919063ffffffff16565b9050600081511115610342578080602001905181019061030291906107be565b610341576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033890610cb1565b60405180910390fd5b5b505050565b6060610356848460008561035f565b90509392505050565b6060824710156103a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039b90610c71565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516103cd9190610af4565b60006040518083038185875af1925050503d806000811461040a576040519150601f19603f3d011682016040523d82523d6000602084013e61040f565b606091505b50915091506104208783838761042c565b92505050949350505050565b6060831561048f5760008351141561048757610447856104a2565b610486576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047d90610c91565b60405180910390fd5b5b82905061049a565b61049983836104c5565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156104d85781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050c9190610c4f565b60405180910390fd5b600061052861052384610cf6565b610cd1565b9050808382526020820190508285602086028201111561054b5761054a610f23565b5b60005b8581101561059957813567ffffffffffffffff81111561057157610570610f1e565b5b80860161057e89826106f0565b8552602085019450602084019350505060018101905061054e565b5050509392505050565b60006105b66105b184610d22565b610cd1565b905080838252602082019050828560208602820111156105d9576105d8610f23565b5b60005b8581101561060957816105ef888261071e565b8452602084019350602083019250506001810190506105dc565b5050509392505050565b600061062661062184610d4e565b610cd1565b90508281526020810184848401111561064257610641610f28565b5b61064d848285610e7c565b509392505050565b6000813590506106648161100f565b92915050565b600082601f83011261067f5761067e610f1e565b5b813561068f848260208601610515565b91505092915050565b600082601f8301126106ad576106ac610f1e565b5b81356106bd8482602086016105a3565b91505092915050565b6000813590506106d581611026565b92915050565b6000815190506106ea81611026565b92915050565b600082601f83011261070557610704610f1e565b5b8135610715848260208601610613565b91505092915050565b60008135905061072d8161103d565b92915050565b60008060006060848603121561074c5761074b610f32565b5b600084013567ffffffffffffffff81111561076a57610769610f2d565b5b6107768682870161066a565b935050602084013567ffffffffffffffff81111561079757610796610f2d565b5b6107a386828701610698565b92505060406107b4868287016106c6565b9150509250925092565b6000602082840312156107d4576107d3610f32565b5b60006107e2848285016106db565b91505092915050565b60008060006060848603121561080457610803610f32565b5b600084013567ffffffffffffffff81111561082257610821610f2d565b5b61082e868287016106f0565b935050602061083f8682870161071e565b9250506040610850868287016106c6565b9150509250925092565b60008060006060848603121561087357610872610f32565b5b60006108818682870161071e565b935050602061089286828701610655565b92505060406108a386828701610655565b9150509250925092565b60006108b983836109fb565b905092915050565b60006108cd8383610ad6565b60208301905092915050565b6108e281610e34565b82525050565b60006108f382610d9f565b6108fd8185610de5565b93508360208202850161090f85610d7f565b8060005b8581101561094b578484038952815161092c85826108ad565b945061093783610dcb565b925060208a01995050600181019050610913565b50829750879550505050505092915050565b600061096882610daa565b6109728185610df6565b935061097d83610d8f565b8060005b838110156109ae57815161099588826108c1565b97506109a083610dd8565b925050600181019050610981565b5085935050505092915050565b6109c481610e46565b82525050565b60006109d582610db5565b6109df8185610e07565b93506109ef818560208601610e8b565b80840191505092915050565b6000610a0682610dc0565b610a108185610e12565b9350610a20818560208601610e8b565b610a2981610f37565b840191505092915050565b6000610a3f82610dc0565b610a498185610e23565b9350610a59818560208601610e8b565b610a6281610f37565b840191505092915050565b6000610a7a602683610e23565b9150610a8582610f48565b604082019050919050565b6000610a9d601d83610e23565b9150610aa882610f97565b602082019050919050565b6000610ac0602a83610e23565b9150610acb82610fc0565b604082019050919050565b610adf81610e72565b82525050565b610aee81610e72565b82525050565b6000610b0082846109ca565b915081905092915050565b6000602082019050610b2060008301846108d9565b92915050565b6000606082019050610b3b60008301866108d9565b610b4860208301856108d9565b610b556040830184610ae5565b949350505050565b6000608082019050610b7260008301876108d9565b8181036020830152610b8481866108e8565b90508181036040830152610b98818561095d565b9050610ba760608301846109bb565b95945050505050565b6000608082019050610bc560008301876108d9565b610bd26020830186610ae5565b610bdf60408301856108d9565b610bec60608301846108d9565b95945050505050565b600060a082019050610c0a60008301886108d9565b610c176020830187610ae5565b8181036040830152610c298186610a34565b9050610c386060830185610ae5565b610c4560808301846109bb565b9695505050505050565b60006020820190508181036000830152610c698184610a34565b905092915050565b60006020820190508181036000830152610c8a81610a6d565b9050919050565b60006020820190508181036000830152610caa81610a90565b9050919050565b60006020820190508181036000830152610cca81610ab3565b9050919050565b6000610cdb610cec565b9050610ce78282610ebe565b919050565b6000604051905090565b600067ffffffffffffffff821115610d1157610d10610eef565b5b602082029050602081019050919050565b600067ffffffffffffffff821115610d3d57610d3c610eef565b5b602082029050602081019050919050565b600067ffffffffffffffff821115610d6957610d68610eef565b5b610d7282610f37565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610e3f82610e52565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610ea9578082015181840152602081019050610e8e565b83811115610eb8576000848401525b50505050565b610ec782610f37565b810181811067ffffffffffffffff82111715610ee657610ee5610eef565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b61101881610e34565b811461102357600080fd5b50565b61102f81610e46565b811461103a57600080fd5b50565b61104681610e72565b811461105157600080fd5b5056fea2646970667358221220086a1c9f56ed96506c7198d50cf431a5b03003c16cd4168d2395cdedfc1ac06164736f6c6343000807003360a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1660601b8152503480156200004757600080fd5b50620000586200005e60201b60201c565b62000208565b600060019054906101000a900460ff1615620000b1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a8906200015c565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff1614620001225760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200011991906200017e565b60405180910390a15b565b6000620001336027836200019b565b91506200014082620001b9565b604082019050919050565b6200015681620001ac565b82525050565b60006020820190508181036000830152620001778162000124565b9050919050565b60006020820190506200019560008301846200014b565b92915050565b600082825260208201905092915050565b600060ff82169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60805160601c612c726200024360003960008181610433015281816104c2015281816105d40152818161066301526107130152612c726000f3fe6080604052600436106100dd5760003560e01c80637993c1e01161007f578063bcf7f32b11610059578063bcf7f32b14610251578063c39aca371461027a578063f2fde38b146102a3578063f45346dc146102cc576100dd565b80637993c1e0146101e65780638129fc1c1461020f5780638da5cb5b14610226576100dd565b80633ce4a5bc116100bb5780633ce4a5bc1461015d5780634f1ef2861461018857806352d1902d146101a4578063715018a6146101cf576100dd565b80630ac7c44c146100e2578063135390f91461010b5780633659cfe614610134575b600080fd5b3480156100ee57600080fd5b5061010960048036038101906101049190611c80565b6102f5565b005b34801561011757600080fd5b50610132600480360381019061012d9190611cfc565b61034c565b005b34801561014057600080fd5b5061015b60048036038101906101569190611b0a565b610431565b005b34801561016957600080fd5b506101726105ba565b60405161017f919061226e565b60405180910390f35b6101a2600480360381019061019d9190611b37565b6105d2565b005b3480156101b057600080fd5b506101b961070f565b6040516101c691906122e9565b60405180910390f35b3480156101db57600080fd5b506101e46107c8565b005b3480156101f257600080fd5b5061020d60048036038101906102089190611d6b565b6107dc565b005b34801561021b57600080fd5b506102246108c7565b005b34801561023257600080fd5b5061023b610a0d565b604051610248919061226e565b60405180910390f35b34801561025d57600080fd5b5061027860048036038101906102739190611e0f565b610a37565b005b34801561028657600080fd5b506102a1600480360381019061029c9190611e0f565b610b2b565b005b3480156102af57600080fd5b506102ca60048036038101906102c59190611b0a565b610d5d565b005b3480156102d857600080fd5b506102f360048036038101906102ee9190611bd3565b610de1565b005b3373ffffffffffffffffffffffffffffffffffffffff167f2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f84848460405161033f93929190612304565b60405180910390a2505050565b60006103588383610f9d565b90503373ffffffffffffffffffffffffffffffffffffffff167f1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d8585848673ffffffffffffffffffffffffffffffffffffffff16634d8943bb6040518163ffffffff1660e01b815260040160206040518083038186803b1580156103db57600080fd5b505afa1580156103ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104139190611ec5565b60405161042394939291906123a0565b60405180910390a250505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614156104c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b79061245c565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166104ff61128d565b73ffffffffffffffffffffffffffffffffffffffff1614610555576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054c9061247c565b60405180910390fd5b61055e816112e4565b6105b781600067ffffffffffffffff81111561057d5761057c61282b565b5b6040519080825280601f01601f1916602001820160405280156105af5781602001600182028036833780820191505090505b5060006112ef565b50565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161415610661576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106589061245c565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166106a061128d565b73ffffffffffffffffffffffffffffffffffffffff16146106f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ed9061247c565b60405180910390fd5b6106ff826112e4565b61070b828260016112ef565b5050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161461079f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107969061249c565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b6107d061146c565b6107da60006114ea565b565b60006107e88585610f9d565b90503373ffffffffffffffffffffffffffffffffffffffff167f1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d8787848873ffffffffffffffffffffffffffffffffffffffff16634d8943bb6040518163ffffffff1660e01b815260040160206040518083038186803b15801561086b57600080fd5b505afa15801561087f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108a39190611ec5565b88886040516108b79695949392919061233d565b60405180910390a2505050505050565b60008060019054906101000a900460ff161590508080156108f85750600160008054906101000a900460ff1660ff16105b806109255750610907306115b0565b1580156109245750600160008054906101000a900460ff1660ff16145b5b610964576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095b906124dc565b60405180910390fd5b60016000806101000a81548160ff021916908360ff16021790555080156109a1576001600060016101000a81548160ff0219169083151502179055505b6109a96115d3565b6109b161162c565b8015610a0a5760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610a0191906123ff565b60405180910390a15b50565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ab0576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1663de43156e87878786866040518663ffffffff1660e01b8152600401610af195949392919061259c565b600060405180830381600087803b158015610b0b57600080fd5b505af1158015610b1f573d6000803e3d6000fd5b50505050505050505050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ba4576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161480610c1d57503073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610c54576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff166347e7ef2484866040518363ffffffff1660e01b8152600401610c8f9291906122c0565b602060405180830381600087803b158015610ca957600080fd5b505af1158015610cbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce19190611c26565b508273ffffffffffffffffffffffffffffffffffffffff1663de43156e87878786866040518663ffffffff1660e01b8152600401610d2395949392919061259c565b600060405180830381600087803b158015610d3d57600080fd5b505af1158015610d51573d6000803e3d6000fd5b50505050505050505050565b610d6561146c565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610dd5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dcc9061243c565b60405180910390fd5b610dde816114ea565b50565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e5a576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161480610ed357503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15610f0a576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166347e7ef2482846040518363ffffffff1660e01b8152600401610f459291906122c0565b602060405180830381600087803b158015610f5f57600080fd5b505af1158015610f73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f979190611c26565b50505050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1663d9eeebed6040518163ffffffff1660e01b8152600401604080518083038186803b158015610fe757600080fd5b505afa158015610ffb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061101f9190611b93565b915091508173ffffffffffffffffffffffffffffffffffffffff166323b872dd3373735b14bb79463307aacbed86daf3322b1e6226ab846040518463ffffffff1660e01b815260040161107493929190612289565b602060405180830381600087803b15801561108e57600080fd5b505af11580156110a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110c69190611c26565b6110fc576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330886040518463ffffffff1660e01b815260040161113993929190612289565b602060405180830381600087803b15801561115357600080fd5b505af1158015611167573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118b9190611c26565b6111c1576040517f4dd9ee8d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff166342966c68866040518263ffffffff1660e01b81526004016111fa91906125f1565b602060405180830381600087803b15801561121457600080fd5b505af1158015611228573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061124c9190611c26565b611282576040517f2c77e05c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b809250505092915050565b60006112bb7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61167d565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6112ec61146c565b50565b61131b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b611687565b60000160009054906101000a900460ff161561133f5761133a83611691565b611467565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561138557600080fd5b505afa9250505080156113b657506040513d601f19601f820116820180604052508101906113b39190611c53565b60015b6113f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113ec906124fc565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b811461145a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611451906124bc565b60405180910390fd5b5061146683838361174a565b5b505050565b611474611776565b73ffffffffffffffffffffffffffffffffffffffff16611492610a0d565b73ffffffffffffffffffffffffffffffffffffffff16146114e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114df9061253c565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff16611622576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116199061257c565b60405180910390fd5b61162a61177e565b565b600060019054906101000a900460ff1661167b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116729061257c565b60405180910390fd5b565b6000819050919050565b6000819050919050565b61169a816115b0565b6116d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116d09061251c565b60405180910390fd5b806117067f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61167d565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611753836117df565b6000825111806117605750805b156117715761176f838361182e565b505b505050565b600033905090565b600060019054906101000a900460ff166117cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117c49061257c565b60405180910390fd5b6117dd6117d8611776565b6114ea565b565b6117e881611691565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b60606118538383604051806060016040528060278152602001612c166027913961185b565b905092915050565b60606000808573ffffffffffffffffffffffffffffffffffffffff16856040516118859190612257565b600060405180830381855af49150503d80600081146118c0576040519150601f19603f3d011682016040523d82523d6000602084013e6118c5565b606091505b50915091506118d6868383876118e1565b925050509392505050565b606083156119445760008351141561193c576118fc856115b0565b61193b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119329061255c565b60405180910390fd5b5b82905061194f565b61194e8383611957565b5b949350505050565b60008251111561196a5781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161199e919061241a565b60405180910390fd5b60006119ba6119b584612631565b61260c565b9050828152602081018484840111156119d6576119d5612878565b5b6119e18482856127b8565b509392505050565b6000813590506119f881612bb9565b92915050565b600081519050611a0d81612bb9565b92915050565b600081519050611a2281612bd0565b92915050565b600081519050611a3781612be7565b92915050565b60008083601f840112611a5357611a52612864565b5b8235905067ffffffffffffffff811115611a7057611a6f61285f565b5b602083019150836001820283011115611a8c57611a8b612873565b5b9250929050565b600082601f830112611aa857611aa7612864565b5b8135611ab88482602086016119a7565b91505092915050565b600060608284031215611ad757611ad6612869565b5b81905092915050565b600081359050611aef81612bfe565b92915050565b600081519050611b0481612bfe565b92915050565b600060208284031215611b2057611b1f612887565b5b6000611b2e848285016119e9565b91505092915050565b60008060408385031215611b4e57611b4d612887565b5b6000611b5c858286016119e9565b925050602083013567ffffffffffffffff811115611b7d57611b7c61287d565b5b611b8985828601611a93565b9150509250929050565b60008060408385031215611baa57611ba9612887565b5b6000611bb8858286016119fe565b9250506020611bc985828601611af5565b9150509250929050565b600080600060608486031215611bec57611beb612887565b5b6000611bfa868287016119e9565b9350506020611c0b86828701611ae0565b9250506040611c1c868287016119e9565b9150509250925092565b600060208284031215611c3c57611c3b612887565b5b6000611c4a84828501611a13565b91505092915050565b600060208284031215611c6957611c68612887565b5b6000611c7784828501611a28565b91505092915050565b600080600060408486031215611c9957611c98612887565b5b600084013567ffffffffffffffff811115611cb757611cb661287d565b5b611cc386828701611a93565b935050602084013567ffffffffffffffff811115611ce457611ce361287d565b5b611cf086828701611a3d565b92509250509250925092565b600080600060608486031215611d1557611d14612887565b5b600084013567ffffffffffffffff811115611d3357611d3261287d565b5b611d3f86828701611a93565b9350506020611d5086828701611ae0565b9250506040611d61868287016119e9565b9150509250925092565b600080600080600060808688031215611d8757611d86612887565b5b600086013567ffffffffffffffff811115611da557611da461287d565b5b611db188828901611a93565b9550506020611dc288828901611ae0565b9450506040611dd3888289016119e9565b935050606086013567ffffffffffffffff811115611df457611df361287d565b5b611e0088828901611a3d565b92509250509295509295909350565b60008060008060008060a08789031215611e2c57611e2b612887565b5b600087013567ffffffffffffffff811115611e4a57611e4961287d565b5b611e5689828a01611ac1565b9650506020611e6789828a016119e9565b9550506040611e7889828a01611ae0565b9450506060611e8989828a016119e9565b935050608087013567ffffffffffffffff811115611eaa57611ea961287d565b5b611eb689828a01611a3d565b92509250509295509295509295565b600060208284031215611edb57611eda612887565b5b6000611ee984828501611af5565b91505092915050565b611efb81612747565b82525050565b611f0a81612747565b82525050565b611f1981612765565b82525050565b6000611f2b8385612678565b9350611f388385846127b8565b611f418361288c565b840190509392505050565b6000611f588385612689565b9350611f658385846127b8565b611f6e8361288c565b840190509392505050565b6000611f8482612662565b611f8e8185612689565b9350611f9e8185602086016127c7565b611fa78161288c565b840191505092915050565b6000611fbd82612662565b611fc7818561269a565b9350611fd78185602086016127c7565b80840191505092915050565b611fec816127a6565b82525050565b6000611ffd8261266d565b61200781856126a5565b93506120178185602086016127c7565b6120208161288c565b840191505092915050565b60006120386026836126a5565b91506120438261289d565b604082019050919050565b600061205b602c836126a5565b9150612066826128ec565b604082019050919050565b600061207e602c836126a5565b91506120898261293b565b604082019050919050565b60006120a16038836126a5565b91506120ac8261298a565b604082019050919050565b60006120c46029836126a5565b91506120cf826129d9565b604082019050919050565b60006120e7602e836126a5565b91506120f282612a28565b604082019050919050565b600061210a602e836126a5565b915061211582612a77565b604082019050919050565b600061212d602d836126a5565b915061213882612ac6565b604082019050919050565b60006121506020836126a5565b915061215b82612b15565b602082019050919050565b6000612173600083612689565b915061217e82612b3e565b600082019050919050565b6000612196601d836126a5565b91506121a182612b41565b602082019050919050565b60006121b9602b836126a5565b91506121c482612b6a565b604082019050919050565b6000606083016121e260008401846126cd565b85830360008701526121f5838284611f1f565b9250505061220660208401846126b6565b6122136020860182611ef2565b506122216040840184612730565b61222e6040860182612239565b508091505092915050565b6122428161278f565b82525050565b6122518161278f565b82525050565b60006122638284611fb2565b915081905092915050565b60006020820190506122836000830184611f01565b92915050565b600060608201905061229e6000830186611f01565b6122ab6020830185611f01565b6122b86040830184612248565b949350505050565b60006040820190506122d56000830185611f01565b6122e26020830184612248565b9392505050565b60006020820190506122fe6000830184611f10565b92915050565b6000604082019050818103600083015261231e8186611f79565b90508181036020830152612333818486611f4c565b9050949350505050565b600060a08201905081810360008301526123578189611f79565b90506123666020830188612248565b6123736040830187612248565b6123806060830186612248565b8181036080830152612393818486611f4c565b9050979650505050505050565b600060a08201905081810360008301526123ba8187611f79565b90506123c96020830186612248565b6123d66040830185612248565b6123e36060830184612248565b81810360808301526123f481612166565b905095945050505050565b60006020820190506124146000830184611fe3565b92915050565b600060208201905081810360008301526124348184611ff2565b905092915050565b600060208201905081810360008301526124558161202b565b9050919050565b600060208201905081810360008301526124758161204e565b9050919050565b6000602082019050818103600083015261249581612071565b9050919050565b600060208201905081810360008301526124b581612094565b9050919050565b600060208201905081810360008301526124d5816120b7565b9050919050565b600060208201905081810360008301526124f5816120da565b9050919050565b60006020820190508181036000830152612515816120fd565b9050919050565b6000602082019050818103600083015261253581612120565b9050919050565b6000602082019050818103600083015261255581612143565b9050919050565b6000602082019050818103600083015261257581612189565b9050919050565b60006020820190508181036000830152612595816121ac565b9050919050565b600060808201905081810360008301526125b681886121cf565b90506125c56020830187611f01565b6125d26040830186612248565b81810360608301526125e5818486611f4c565b90509695505050505050565b60006020820190506126066000830184612248565b92915050565b6000612616612627565b905061262282826127fa565b919050565b6000604051905090565b600067ffffffffffffffff82111561264c5761264b61282b565b5b6126558261288c565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006126c560208401846119e9565b905092915050565b600080833560016020038436030381126126ea576126e9612882565b5b83810192508235915060208301925067ffffffffffffffff8211156127125761271161285a565b5b6001820236038413156127285761272761286e565b5b509250929050565b600061273f6020840184611ae0565b905092915050565b60006127528261276f565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006127b182612799565b9050919050565b82818337600083830152505050565b60005b838110156127e55780820151818401526020810190506127ca565b838111156127f4576000848401525b50505050565b6128038261288c565b810181811067ffffffffffffffff821117156128225761282161282b565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b612bc281612747565b8114612bcd57600080fd5b50565b612bd981612759565b8114612be457600080fd5b50565b612bf081612765565b8114612bfb57600080fd5b50565b612c078161278f565b8114612c1257600080fd5b5056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220627eb039ecb52d09fd70ebde1f5eca61fe0f91321a96929ebe38dc8e38d999ab64736f6c63430008070033608060405234801561001057600080fd5b50604051610bcd380380610bcd8339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610108565b600081519050610087816100f1565b92915050565b6000602082840312156100a3576100a26100ec565b5b60006100b184828501610078565b91505092915050565b60006100c5826100cc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6100fa816100ba565b811461010557600080fd5b50565b610ab6806101176000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b614610062578063a0a1730b14610080575b600080fd5b610060600480360381019061005b91906105fd565b61009c565b005b61006a6102af565b6040516100779190610761565b60405180910390f35b61009a6004803603810190610095919061055e565b6102d3565b005b60008383836040516024016100b39392919061082f565b6040516020818303038152906040527fe04d4f97000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090508473ffffffffffffffffffffffffffffffffffffffff1663095ea7b360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16886040518363ffffffff1660e01b815260040161018d92919061077c565b602060405180830381600087803b1580156101a757600080fd5b505af11580156101bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101df9190610531565b610215576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637993c1e0888888856040518563ffffffff1660e01b815260040161027494939291906107dc565b600060405180830381600087803b15801561028e57600080fd5b505af11580156102a2573d6000803e3d6000fd5b5050505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008383836040516024016102ea9392919061082f565b6040516020818303038152906040527fe04d4f97000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630ac7c44c86836040518363ffffffff1660e01b81526004016103c49291906107a5565b600060405180830381600087803b1580156103de57600080fd5b505af11580156103f2573d6000803e3d6000fd5b505050505050505050565b600061041061040b84610892565b61086d565b90508281526020810184848401111561042c5761042b610a1b565b5b610437848285610974565b509392505050565b600061045261044d846108c3565b61086d565b90508281526020810184848401111561046e5761046d610a1b565b5b610479848285610974565b509392505050565b60008135905061049081610a3b565b92915050565b6000813590506104a581610a52565b92915050565b6000815190506104ba81610a52565b92915050565b600082601f8301126104d5576104d4610a16565b5b81356104e58482602086016103fd565b91505092915050565b600082601f83011261050357610502610a16565b5b813561051384826020860161043f565b91505092915050565b60008135905061052b81610a69565b92915050565b60006020828403121561054757610546610a25565b5b6000610555848285016104ab565b91505092915050565b6000806000806080858703121561057857610577610a25565b5b600085013567ffffffffffffffff81111561059657610595610a20565b5b6105a2878288016104c0565b945050602085013567ffffffffffffffff8111156105c3576105c2610a20565b5b6105cf878288016104ee565b93505060406105e08782880161051c565b92505060606105f187828801610496565b91505092959194509250565b60008060008060008060c0878903121561061a57610619610a25565b5b600087013567ffffffffffffffff81111561063857610637610a20565b5b61064489828a016104c0565b965050602061065589828a0161051c565b955050604061066689828a01610481565b945050606087013567ffffffffffffffff81111561068757610686610a20565b5b61069389828a016104ee565b93505060806106a489828a0161051c565b92505060a06106b589828a01610496565b9150509295509295509295565b6106cb8161092c565b82525050565b6106da8161093e565b82525050565b60006106eb826108f4565b6106f5818561090a565b9350610705818560208601610983565b61070e81610a2a565b840191505092915050565b6000610724826108ff565b61072e818561091b565b935061073e818560208601610983565b61074781610a2a565b840191505092915050565b61075b8161096a565b82525050565b600060208201905061077660008301846106c2565b92915050565b600060408201905061079160008301856106c2565b61079e6020830184610752565b9392505050565b600060408201905081810360008301526107bf81856106e0565b905081810360208301526107d381846106e0565b90509392505050565b600060808201905081810360008301526107f681876106e0565b90506108056020830186610752565b61081260408301856106c2565b818103606083015261082481846106e0565b905095945050505050565b600060608201905081810360008301526108498186610719565b90506108586020830185610752565b61086560408301846106d1565b949350505050565b6000610877610888565b905061088382826109b6565b919050565b6000604051905090565b600067ffffffffffffffff8211156108ad576108ac6109e7565b5b6108b682610a2a565b9050602081019050919050565b600067ffffffffffffffff8211156108de576108dd6109e7565b5b6108e782610a2a565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60006109378261094a565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156109a1578082015181840152602081019050610986565b838111156109b0576000848401525b50505050565b6109bf82610a2a565b810181811067ffffffffffffffff821117156109de576109dd6109e7565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b610a448161092c565b8114610a4f57600080fd5b50565b610a5b8161093e565b8114610a6657600080fd5b50565b610a728161096a565b8114610a7d57600080fd5b5056fea2646970667358221220192b4e23b9b6daaae9f30fb00f65433e56a5d8a6906223e4fbd169def800a8a264736f6c6343000807003360806040523480156200001157600080fd5b50604051620010d7380380620010d7833981810160405281019062000037919062000146565b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac560405160405180910390a1505050620001f5565b6000815190506200014081620001db565b92915050565b600080600060608486031215620001625762000161620001d6565b5b600062000172868287016200012f565b935050602062000185868287016200012f565b925050604062000198868287016200012f565b9150509250925092565b6000620001af82620001b6565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620001e681620001a2565b8114620001f257600080fd5b50565b610ed280620002056000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806397770dff1161007157806397770dff14610166578063a7cb050714610182578063c63585cc1461019e578063d7fd7afb146101ce578063d936a012146101fe578063ee2815ba1461021c576100a9565b80630be15547146100ae5780633c669d55146100de578063513a9c05146100fa578063569541b91461012a578063842da36d14610148575b600080fd5b6100c860048036038101906100c3919061094d565b610238565b6040516100d59190610bce565b60405180910390f35b6100f860048036038101906100f39190610898565b61026b565b005b610114600480360381019061010f919061094d565b6103b8565b6040516101219190610bce565b60405180910390f35b6101326103eb565b60405161013f9190610bce565b60405180910390f35b610150610411565b60405161015d9190610bce565b60405180910390f35b610180600480360381019061017b9190610818565b610437565b005b61019c600480360381019061019791906109ba565b6104d4565b005b6101b860048036038101906101b39190610845565b610528565b6040516101c59190610bce565b60405180910390f35b6101e860048036038101906101e3919061094d565b61059a565b6040516101f59190610c67565b60405180910390f35b6102066105b2565b6040516102139190610bce565b60405180910390f35b6102366004803603810190610231919061097a565b6105d8565b005b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060405180606001604052806040518060200160405280600081525081526020013373ffffffffffffffffffffffffffffffffffffffff1681526020014681525090508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb87866040518363ffffffff1660e01b81526004016102ea929190610be9565b602060405180830381600087803b15801561030457600080fd5b505af1158015610318573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061033c9190610920565b508573ffffffffffffffffffffffffffffffffffffffff1663de43156e82878787876040518663ffffffff1660e01b815260040161037e959493929190610c12565b600060405180830381600087803b15801561039857600080fd5b505af11580156103ac573d6000803e3d6000fd5b50505050505050505050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516104c99190610bce565b60405180910390a150565b80600080848152602001908152602001600020819055507f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d828260405161051c929190610cab565b60405180910390a15050565b60008060006105378585610667565b9150915085828260405160200161054f929190610b60565b60405160208183030381529060405280519060200120604051602001610576929190610b8c565b6040516020818303038152906040528051906020012060001c925050509392505050565b60006020528060005260406000206000915090505481565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d828260405161065b929190610c82565b60405180910390a15050565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156106d0576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161061070a57828461070d565b83835b8092508193505050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561077c576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b60008135905061079281610e57565b92915050565b6000815190506107a781610e6e565b92915050565b60008083601f8401126107c3576107c2610dd3565b5b8235905067ffffffffffffffff8111156107e0576107df610dce565b5b6020830191508360018202830111156107fc576107fb610dd8565b5b9250929050565b60008135905061081281610e85565b92915050565b60006020828403121561082e5761082d610de2565b5b600061083c84828501610783565b91505092915050565b60008060006060848603121561085e5761085d610de2565b5b600061086c86828701610783565b935050602061087d86828701610783565b925050604061088e86828701610783565b9150509250925092565b6000806000806000608086880312156108b4576108b3610de2565b5b60006108c288828901610783565b95505060206108d388828901610783565b94505060406108e488828901610803565b935050606086013567ffffffffffffffff81111561090557610904610ddd565b5b610911888289016107ad565b92509250509295509295909350565b60006020828403121561093657610935610de2565b5b600061094484828501610798565b91505092915050565b60006020828403121561096357610962610de2565b5b600061097184828501610803565b91505092915050565b6000806040838503121561099157610990610de2565b5b600061099f85828601610803565b92505060206109b085828601610783565b9150509250929050565b600080604083850312156109d1576109d0610de2565b5b60006109df85828601610803565b92505060206109f085828601610803565b9150509250929050565b610a0381610d0c565b82525050565b610a1281610d0c565b82525050565b610a29610a2482610d0c565b610da0565b82525050565b610a40610a3b82610d2a565b610db2565b82525050565b6000610a528385610cf0565b9350610a5f838584610d5e565b610a6883610de7565b840190509392505050565b6000610a7e82610cd4565b610a888185610cdf565b9350610a98818560208601610d6d565b610aa181610de7565b840191505092915050565b6000610ab9602083610d01565b9150610ac482610e05565b602082019050919050565b6000610adc600183610d01565b9150610ae782610e2e565b600182019050919050565b60006060830160008301518482036000860152610b0f8282610a73565b9150506020830151610b2460208601826109fa565b506040830151610b376040860182610b42565b508091505092915050565b610b4b81610d54565b82525050565b610b5a81610d54565b82525050565b6000610b6c8285610a18565b601482019150610b7c8284610a18565b6014820191508190509392505050565b6000610b9782610acf565b9150610ba38285610a18565b601482019150610bb38284610a2f565b602082019150610bc282610aac565b91508190509392505050565b6000602082019050610be36000830184610a09565b92915050565b6000604082019050610bfe6000830185610a09565b610c0b6020830184610b51565b9392505050565b60006080820190508181036000830152610c2c8188610af2565b9050610c3b6020830187610a09565b610c486040830186610b51565b8181036060830152610c5b818486610a46565b90509695505050505050565b6000602082019050610c7c6000830184610b51565b92915050565b6000604082019050610c976000830185610b51565b610ca46020830184610a09565b9392505050565b6000604082019050610cc06000830185610b51565b610ccd6020830184610b51565b9392505050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000610d1782610d34565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610d8b578082015181840152602081019050610d70565b83811115610d9a576000848401525b50505050565b6000610dab82610dbc565b9050919050565b6000819050919050565b6000610dc782610df8565b9050919050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f600082015250565b7fff00000000000000000000000000000000000000000000000000000000000000600082015250565b610e6081610d0c565b8114610e6b57600080fd5b50565b610e7781610d1e565b8114610e8257600080fd5b50565b610e8e81610d54565b8114610e9957600080fd5b5056fea2646970667358221220fc7ae99de1eb543fc941c22b053bacc8edc14f929260909668ffd15cbad6fbfb64736f6c6343000807003360c06040523480156200001157600080fd5b506040516200282d3803806200282d83398181016040528101906200003791906200035b565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620000b1576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8760079080519060200190620000c9929190620001d1565b508660089080519060200190620000e2929190620001d1565b5085600960006101000a81548160ff021916908360ff16021790555084608081815250508360028111156200011c576200011b620005ae565b5b60a0816002811115620001345762000133620005ae565b5b60f81b8152505082600281905550816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050505050620006bf565b828054620001df9062000542565b90600052602060002090601f0160209004810192826200020357600085556200024f565b82601f106200021e57805160ff19168380011785556200024f565b828001600101855582156200024f579182015b828111156200024e57825182559160200191906001019062000231565b5b5090506200025e919062000262565b5090565b5b808211156200027d57600081600090555060010162000263565b5090565b60006200029862000292846200048b565b62000462565b905082815260208101848484011115620002b757620002b662000640565b5b620002c48482856200050c565b509392505050565b600081519050620002dd8162000660565b92915050565b600081519050620002f4816200067a565b92915050565b600082601f8301126200031257620003116200063b565b5b81516200032484826020860162000281565b91505092915050565b6000815190506200033e816200068b565b92915050565b6000815190506200035581620006a5565b92915050565b600080600080600080600080610100898b0312156200037f576200037e6200064a565b5b600089015167ffffffffffffffff811115620003a0576200039f62000645565b5b620003ae8b828c01620002fa565b985050602089015167ffffffffffffffff811115620003d257620003d162000645565b5b620003e08b828c01620002fa565b9750506040620003f38b828c0162000344565b9650506060620004068b828c016200032d565b9550506080620004198b828c01620002e3565b94505060a06200042c8b828c016200032d565b93505060c06200043f8b828c01620002cc565b92505060e0620004528b828c01620002cc565b9150509295985092959890939650565b60006200046e62000481565b90506200047c828262000578565b919050565b6000604051905090565b600067ffffffffffffffff821115620004a957620004a86200060c565b5b620004b4826200064f565b9050602081019050919050565b6000620004ce82620004d5565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b838110156200052c5780820151818401526020810190506200050f565b838111156200053c576000848401525b50505050565b600060028204905060018216806200055b57607f821691505b60208210811415620005725762000571620005dd565b5b50919050565b62000583826200064f565b810181811067ffffffffffffffff82111715620005a557620005a46200060c565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6200066b81620004c1565b81146200067757600080fd5b50565b600381106200068857600080fd5b50565b6200069681620004f5565b8114620006a257600080fd5b50565b620006b081620004ff565b8114620006bc57600080fd5b50565b60805160a05160f81c612137620006f660003960006109580152600081816108a201528181610c250152610d5a01526121376000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806385e1f4d0116100c3578063d9eeebed1161007c578063d9eeebed146103cc578063dd62ed3e146103eb578063e2f535b81461041b578063eddeb12314610439578063f2441b3214610455578063f687d12a146104735761014d565b806385e1f4d0146102f657806395d89b4114610314578063a3413d0314610332578063a9059cbb14610350578063c701262614610380578063c835d7cc146103b05761014d565b8063313ce56711610115578063313ce5671461020c5780633ce4a5bc1461022a57806342966c681461024857806347e7ef24146102785780634d8943bb146102a857806370a08231146102c65761014d565b806306fdde0314610152578063091d278814610170578063095ea7b31461018e57806318160ddd146101be57806323b872dd146101dc575b600080fd5b61015a61048f565b6040516101679190611cad565b60405180910390f35b610178610521565b6040516101859190611ccf565b60405180910390f35b6101a860048036038101906101a3919061196e565b610527565b6040516101b59190611bfb565b60405180910390f35b6101c6610545565b6040516101d39190611ccf565b60405180910390f35b6101f660048036038101906101f1919061191b565b61054f565b6040516102039190611bfb565b60405180910390f35b610214610647565b6040516102219190611cea565b60405180910390f35b61023261065e565b60405161023f9190611b80565b60405180910390f35b610262600480360381019061025d9190611a37565b610676565b60405161026f9190611bfb565b60405180910390f35b610292600480360381019061028d919061196e565b61068b565b60405161029f9190611bfb565b60405180910390f35b6102b0610851565b6040516102bd9190611ccf565b60405180910390f35b6102e060048036038101906102db9190611881565b610857565b6040516102ed9190611ccf565b60405180910390f35b6102fe6108a0565b60405161030b9190611ccf565b60405180910390f35b61031c6108c4565b6040516103299190611cad565b60405180910390f35b61033a610956565b6040516103479190611c92565b60405180910390f35b61036a6004803603810190610365919061196e565b61097a565b6040516103779190611bfb565b60405180910390f35b61039a600480360381019061039591906119db565b610998565b6040516103a79190611bfb565b60405180910390f35b6103ca60048036038101906103c59190611881565b610aee565b005b6103d4610be1565b6040516103e2929190611bd2565b60405180910390f35b610405600480360381019061040091906118db565b610e4e565b6040516104129190611ccf565b60405180910390f35b610423610ed5565b6040516104309190611b80565b60405180910390f35b610453600480360381019061044e9190611a37565b610efb565b005b61045d610fb5565b60405161046a9190611b80565b60405180910390f35b61048d60048036038101906104889190611a37565b610fd9565b005b60606007805461049e90611f33565b80601f01602080910402602001604051908101604052809291908181526020018280546104ca90611f33565b80156105175780601f106104ec57610100808354040283529160200191610517565b820191906000526020600020905b8154815290600101906020018083116104fa57829003601f168201915b5050505050905090565b60025481565b600061053b610534611093565b848461109b565b6001905092915050565b6000600654905090565b600061055c848484611254565b6000600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006105a7611093565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508281101561061e576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61063b8561062a611093565b85846106369190611e43565b61109b565b60019150509392505050565b6000600960009054906101000a900460ff16905090565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b600061068233836114b0565b60019050919050565b600073735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610729575060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b80156107835750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b156107ba576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107c48383611668565b8273ffffffffffffffffffffffffffffffffffffffff167f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab373735b14bb79463307aacbed86daf3322b1e6226ab6040516020016108219190611b65565b6040516020818303038152906040528460405161083f929190611c16565b60405180910390a26001905092915050565b60035481565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6060600880546108d390611f33565b80601f01602080910402602001604051908101604052809291908181526020018280546108ff90611f33565b801561094c5780601f106109215761010080835404028352916020019161094c565b820191906000526020600020905b81548152906001019060200180831161092f57829003601f168201915b5050505050905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b600061098e610987611093565b8484611254565b6001905092915050565b60008060006109a5610be1565b915091508173ffffffffffffffffffffffffffffffffffffffff166323b872dd3373735b14bb79463307aacbed86daf3322b1e6226ab846040518463ffffffff1660e01b81526004016109fa93929190611b9b565b602060405180830381600087803b158015610a1457600080fd5b505af1158015610a28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4c91906119ae565b610a82576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a8c33856114b0565b3373ffffffffffffffffffffffffffffffffffffffff167f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955868684600354604051610ada9493929190611c46565b60405180910390a260019250505092915050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b67576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae81604051610bd69190611b80565b60405180910390a150565b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630be155477f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b8152600401610c609190611ccf565b60206040518083038186803b158015610c7857600080fd5b505afa158015610c8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb091906118ae565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610d19576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d7fd7afb7f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b8152600401610d959190611ccf565b60206040518083038186803b158015610dad57600080fd5b505afa158015610dc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610de59190611a64565b90506000811415610e22576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060035460025483610e359190611de9565b610e3f9190611d93565b90508281945094505050509091565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f74576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806003819055507fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f81604051610faa9190611ccf565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611052576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806002819055507fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a816040516110889190611ccf565b60405180910390a150565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611102576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611169576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516112479190611ccf565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156112bb576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611322576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156113a0576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81816113ac9190611e43565b600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461143e9190611d93565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516114a29190611ccf565b60405180910390a350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611517576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611595576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81816115a19190611e43565b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600660008282546115f69190611e43565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161165b9190611ccf565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156116cf576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600660008282546116e19190611d93565b9250508190555080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546117379190611d93565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161179c9190611ccf565b60405180910390a35050565b60006117bb6117b684611d2a565b611d05565b9050828152602081018484840111156117d7576117d661207b565b5b6117e2848285611ef1565b509392505050565b6000813590506117f9816120bc565b92915050565b60008151905061180e816120bc565b92915050565b600081519050611823816120d3565b92915050565b600082601f83011261183e5761183d612076565b5b813561184e8482602086016117a8565b91505092915050565b600081359050611866816120ea565b92915050565b60008151905061187b816120ea565b92915050565b60006020828403121561189757611896612085565b5b60006118a5848285016117ea565b91505092915050565b6000602082840312156118c4576118c3612085565b5b60006118d2848285016117ff565b91505092915050565b600080604083850312156118f2576118f1612085565b5b6000611900858286016117ea565b9250506020611911858286016117ea565b9150509250929050565b60008060006060848603121561193457611933612085565b5b6000611942868287016117ea565b9350506020611953868287016117ea565b925050604061196486828701611857565b9150509250925092565b6000806040838503121561198557611984612085565b5b6000611993858286016117ea565b92505060206119a485828601611857565b9150509250929050565b6000602082840312156119c4576119c3612085565b5b60006119d284828501611814565b91505092915050565b600080604083850312156119f2576119f1612085565b5b600083013567ffffffffffffffff811115611a1057611a0f612080565b5b611a1c85828601611829565b9250506020611a2d85828601611857565b9150509250929050565b600060208284031215611a4d57611a4c612085565b5b6000611a5b84828501611857565b91505092915050565b600060208284031215611a7a57611a79612085565b5b6000611a888482850161186c565b91505092915050565b611a9a81611e77565b82525050565b611ab1611aac82611e77565b611f96565b82525050565b611ac081611e89565b82525050565b6000611ad182611d5b565b611adb8185611d71565b9350611aeb818560208601611f00565b611af48161208a565b840191505092915050565b611b0881611edf565b82525050565b6000611b1982611d66565b611b238185611d82565b9350611b33818560208601611f00565b611b3c8161208a565b840191505092915050565b611b5081611ec8565b82525050565b611b5f81611ed2565b82525050565b6000611b718284611aa0565b60148201915081905092915050565b6000602082019050611b956000830184611a91565b92915050565b6000606082019050611bb06000830186611a91565b611bbd6020830185611a91565b611bca6040830184611b47565b949350505050565b6000604082019050611be76000830185611a91565b611bf46020830184611b47565b9392505050565b6000602082019050611c106000830184611ab7565b92915050565b60006040820190508181036000830152611c308185611ac6565b9050611c3f6020830184611b47565b9392505050565b60006080820190508181036000830152611c608187611ac6565b9050611c6f6020830186611b47565b611c7c6040830185611b47565b611c896060830184611b47565b95945050505050565b6000602082019050611ca76000830184611aff565b92915050565b60006020820190508181036000830152611cc78184611b0e565b905092915050565b6000602082019050611ce46000830184611b47565b92915050565b6000602082019050611cff6000830184611b56565b92915050565b6000611d0f611d20565b9050611d1b8282611f65565b919050565b6000604051905090565b600067ffffffffffffffff821115611d4557611d44612047565b5b611d4e8261208a565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611d9e82611ec8565b9150611da983611ec8565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611dde57611ddd611fba565b5b828201905092915050565b6000611df482611ec8565b9150611dff83611ec8565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611e3857611e37611fba565b5b828202905092915050565b6000611e4e82611ec8565b9150611e5983611ec8565b925082821015611e6c57611e6b611fba565b5b828203905092915050565b6000611e8282611ea8565b9050919050565b60008115159050919050565b6000819050611ea3826120a8565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611eea82611e95565b9050919050565b82818337600083830152505050565b60005b83811015611f1e578082015181840152602081019050611f03565b83811115611f2d576000848401525b50505050565b60006002820490506001821680611f4b57607f821691505b60208210811415611f5f57611f5e612018565b5b50919050565b611f6e8261208a565b810181811067ffffffffffffffff82111715611f8d57611f8c612047565b5b80604052505050565b6000611fa182611fa8565b9050919050565b6000611fb38261209b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b600381106120b9576120b8611fe9565b5b50565b6120c581611e77565b81146120d057600080fd5b50565b6120dc81611e89565b81146120e757600080fd5b50565b6120f381611ec8565b81146120fe57600080fd5b5056fea26469706673582212206982505c1a546edfb86a1aaaca0ad6b7e5542f2d1f24ac30a032805fc80a650e64736f6c63430008070033a2646970667358221220da501b74b5f247f506271d2056d0f490e4e35a1198496514f5bbfae138fcc94264736f6c63430008070033",
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
	To              []byte
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message)
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

// WatchWithdrawal is a free log subscription operation binding the contract event 0x1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message)
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

// ParseWithdrawal is a log parse operation binding the contract event 0x1866ad2994816c79f4103e1eddacc7b085eb7c635205243a28940be69b01536d.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message)
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
